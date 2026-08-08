package analysis

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/goropikari/goreadable/internal/config"
	"github.com/goropikari/goreadable/internal/report"
)

//nolint:cyclop // File discovery keeps exclusion rules in one observable boundary.
func Files(root string, recursive bool) ([]string, error) {
	var files []string

	if recursive {
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				if info.Name() == "vendor" {
					return filepath.SkipDir
				}

				return nil
			}

			if filepath.Ext(path) == ".go" && !isGenerated(path) {
				files = append(files, path)
			}

			return nil
		})
		if err != nil {
			return nil, err
		}
	} else {
		entries, err := os.ReadDir(root)
		if err != nil {
			return nil, err
		}

		for _, entry := range entries {
			if !entry.IsDir() && filepath.Ext(entry.Name()) == ".go" && !isGenerated(filepath.Join(root, entry.Name())) {
				files = append(files, filepath.Join(root, entry.Name()))
			}
		}
	}

	sort.Strings(files)

	return files, nil
}

//nolint:cyclop,gocognit,wsl_v5 // This is the single declaration-to-candidate boundary.
func Analyze(files []string, thresholds config.Thresholds, changed map[string][][2]int) (report.Result, error) {
	result := report.Result{Version: 1}
	packageMethods := map[string]int{}
	for _, path := range files {
		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			return report.Result{}, fmt.Errorf("parse %s: %w", path, err)
		}
		for name, count := range methodCounts(file) {
			packageMethods[name] += count
		}
	}

	for _, path := range files {
		fset := token.NewFileSet()

		file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			return report.Result{}, fmt.Errorf("parse %s: %w", path, err)
		}

		sourceBytes, err := os.ReadFile(path)
		if err != nil {
			return report.Result{}, err
		}

		lines := strings.Split(string(sourceBytes), "\n")

		codeKind := "production"
		if strings.HasSuffix(path, "_test.go") {
			codeKind = "test"
		}

		ignoredTypes := ignoredTypeNames(file)
		ast.Inspect(file, func(node ast.Node) bool {
			switch declaration := node.(type) {
			case *ast.FuncDecl:
				if hasIgnoreDirective(declaration.Doc) {
					return true
				}
				start, end := fset.Position(declaration.Pos()).Line, fset.Position(declaration.End()).Line
				if !overlaps(changed, path, start, end) {
					return true
				}

				metrics := map[string]int{"function_lines": end - start + 1, "nesting_depth": nesting(declaration.Body), "cyclomatic_complexity": complexity(declaration.Body), "function_arguments": argumentCount(declaration.Type)}
				thresholdMap := map[string]int{"function_lines": thresholds.FunctionLines, "nesting_depth": thresholds.NestingDepth, "cyclomatic_complexity": thresholds.CyclomaticComplexity, "function_arguments": thresholds.FunctionArguments}

				reasons := reasons(metrics, thresholdMap)
				if len(reasons) > 0 {
					result.Candidates = append(result.Candidates, candidate("function", declaration.Name.Name, path, start, end, codeKind, metrics, thresholdMap, reasons, lines))
				}
			case *ast.TypeSpec:
				if ignoredTypes[declaration.Name.Name] || hasIgnoreDirective(declaration.Doc) {
					return true
				}
				structure, ok := declaration.Type.(*ast.StructType)
				if !ok {
					return true
				}

				start, end := fset.Position(declaration.Pos()).Line, fset.Position(declaration.End()).Line
				if !overlaps(changed, path, start, end) {
					return true
				}

				metrics := map[string]int{"struct_fields": structure.Fields.NumFields(), "type_methods": packageMethods[declaration.Name.Name]}
				thresholdMap := map[string]int{"struct_fields": thresholds.StructFields, "type_methods": thresholds.TypeMethods}

				reasons := reasons(metrics, thresholdMap)
				if len(reasons) > 0 {
					result.Candidates = append(result.Candidates, candidate("type", declaration.Name.Name, path, start, end, codeKind, metrics, thresholdMap, reasons, lines))
				}
			}

			return true
		})
	}

	return result, nil
}

func ignoredTypeNames(file *ast.File) map[string]bool {
	ignored := map[string]bool{}

	for _, declaration := range file.Decls {
		group, ok := declaration.(*ast.GenDecl)
		if !ok || !hasIgnoreDirective(group.Doc) {
			continue
		}

		for _, specification := range group.Specs {
			typeSpec, ok := specification.(*ast.TypeSpec)
			if ok {
				ignored[typeSpec.Name.Name] = true
			}
		}
	}

	return ignored
}

func hasIgnoreDirective(group *ast.CommentGroup) bool {
	if group == nil {
		return false
	}

	for _, comment := range group.List {
		if strings.Contains(comment.Text, "goreadable:ignore") {
			return true
		}
	}

	return false
}

func candidate(kind, name, path string, start, end int, codeKind string, metrics, thresholds map[string]int, reasons []string, lines []string) report.Candidate {
	if start < 1 {
		start = 1
	}

	if end > len(lines) {
		end = len(lines)
	}

	return report.Candidate{Kind: kind, Name: name, Path: path, StartLine: start, EndLine: end, CodeKind: codeKind, Metrics: metrics, Thresholds: thresholds, Reasons: reasons, Source: strings.Join(lines[start-1:end], "\n")}
}

func reasons(metrics, thresholds map[string]int) []string {
	var result []string

	for key, value := range metrics {
		if value > thresholds[key] {
			result = append(result, fmt.Sprintf("%s=%d exceeds threshold %d", key, value, thresholds[key]))
		}
	}

	sort.Strings(result)

	return result
}

func argumentCount(function *ast.FuncType) int {
	if function.Params == nil {
		return 0
	}

	count := 0

	for _, field := range function.Params.List {
		if len(field.Names) == 0 {
			count++
		} else {
			count += len(field.Names)
		}
	}

	return count
}

//nolint:wsl_v5 // The traversal maintains explicit enter/exit state.
func nesting(body *ast.BlockStmt) int {
	if body == nil {
		return 0
	}

	max, depth := 0, 0
	branches := []bool{}
	ast.Inspect(body, func(node ast.Node) bool {
		if node == nil {
			if len(branches) > 0 {
				if branches[len(branches)-1] {
					depth--
				}
				branches = branches[:len(branches)-1]
			}
			return true
		}
		branch := false
		switch node.(type) {
		case *ast.IfStmt, *ast.ForStmt, *ast.RangeStmt, *ast.SwitchStmt, *ast.TypeSwitchStmt, *ast.SelectStmt:
			branch = true
		}
		branches = append(branches, branch)
		if branch {
			depth++
			if depth > max {
				max = depth
			}
		}
		return true
	})

	return max
}

func complexity(body *ast.BlockStmt) int {
	total := 1
	if body == nil {
		return total
	}

	ast.Inspect(body, func(node ast.Node) bool {
		switch node.(type) {
		case *ast.IfStmt, *ast.ForStmt, *ast.RangeStmt, *ast.CaseClause, *ast.CommClause:
			total++
		}

		return true
	})

	return total
}

func methodCounts(file *ast.File) map[string]int {
	counts := map[string]int{}

	for _, declaration := range file.Decls {
		function, ok := declaration.(*ast.FuncDecl)
		if !ok || function.Recv == nil || len(function.Recv.List) == 0 {
			continue
		}

		receiver := function.Recv.List[0].Type
		if pointer, ok := receiver.(*ast.StarExpr); ok {
			receiver = pointer.X
		}

		if identifier, ok := receiver.(*ast.Ident); ok {
			counts[identifier.Name]++
		}
	}

	return counts
}

func overlaps(changed map[string][][2]int, path string, start, end int) bool {
	ranges, ok := changed[path]
	if !ok {
		return changed == nil
	}

	for _, interval := range ranges {
		if start <= interval[1] && end >= interval[0] {
			return true
		}
	}

	return false
}

func isGenerated(path string) bool {
	data, err := os.ReadFile(path)
	return err == nil && strings.Contains(string(data), "Code generated") && strings.Contains(string(data), "DO NOT EDIT")
}
