package analysis

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/goropikari/goreadable/internal/config"
	"github.com/goropikari/goreadable/internal/report"
)

type Options struct {
	FilterByThresholds bool
	FunctionSelectors  map[string]struct{}
}

func NewOptions(filterByThresholds bool, functionSelectors []string) (Options, error) {
	options := Options{
		FilterByThresholds: filterByThresholds,
		FunctionSelectors:  make(map[string]struct{}, len(functionSelectors)),
	}

	for _, selector := range functionSelectors {
		selector = strings.TrimSpace(selector)
		if selector == "" {
			return Options{}, fmt.Errorf("--function must be a package-qualified function name")
		}

		options.FunctionSelectors[selector] = struct{}{}
	}

	return options, nil
}

func (options Options) MetricsOnly() bool {
	return !options.FilterByThresholds
}

func (options Options) IncludesFunction(packageName string, declaration *ast.FuncDecl) bool {
	if len(options.FunctionSelectors) == 0 {
		return true
	}

	_, ok := options.FunctionSelectors[functionSelector(packageName, declaration)]

	return ok
}

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
	return AnalyzeWithOptions(files, thresholds, changed, Options{FilterByThresholds: true})
}

//nolint:cyclop,gocognit,wsl_v5 // This is the single declaration-to-candidate boundary.
func AnalyzeWithOptions(files []string, thresholds config.Thresholds, changed map[string][][2]int, options Options) (report.Result, error) {
	result := report.Result{
		Version:     1,
		MetricsOnly: options.MetricsOnly(),
	}
	packageMethods := map[string]int{}
	packageExportedMethods := map[string]int{}
	for _, path := range files {
		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			return report.Result{}, fmt.Errorf("parse %s: %w", path, err)
		}
		for name, count := range methodCounts(file) {
			packageMethods[name] += count
		}
		for name, count := range exportedMethodCounts(file) {
			packageExportedMethods[name] += count
		}
	}

	for _, path := range files {
		fset := token.NewFileSet()

		file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			return report.Result{}, fmt.Errorf("parse %s: %w", path, err)
		}

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

				metrics := functionMetrics(declaration, file, fset, start, end)
				thresholdMap := functionThresholds(thresholds)

				if options.MetricsOnly() {
					if options.IncludesFunction(file.Name.Name, declaration) {
						result.Candidates = append(result.Candidates, candidate("function", declaration.Name.Name, path, start, end, codeKind, metrics, thresholdMap, nil))
					}

					return true
				}

				reasons := reasons(metrics, thresholdMap)
				if len(reasons) > 0 {
					result.Candidates = append(result.Candidates, candidate("function", declaration.Name.Name, path, start, end, codeKind, metrics, thresholdMap, reasons))
				}
			case *ast.TypeSpec:
				if options.MetricsOnly() {
					return true
				}

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

				metrics := map[string]int{"struct_fields": structure.Fields.NumFields(), "type_methods": packageMethods[declaration.Name.Name], "exported_members": exportedFields(structure) + packageExportedMethods[declaration.Name.Name]}
				thresholdMap := map[string]int{"struct_fields": thresholds.StructFields, "type_methods": thresholds.TypeMethods, "exported_members": thresholds.ExportedMembers}

				reasons := reasons(metrics, thresholdMap)
				if len(reasons) > 0 {
					result.Candidates = append(result.Candidates, candidate("type", declaration.Name.Name, path, start, end, codeKind, metrics, thresholdMap, reasons))
				}
			}

			return true
		})
	}

	return result, nil
}

func functionSelector(packageName string, declaration *ast.FuncDecl) string {
	if declaration.Recv == nil || len(declaration.Recv.List) == 0 {
		return packageName + "." + declaration.Name.Name
	}

	return packageName + "." + receiverName(declaration.Recv.List[0].Type) + "." + declaration.Name.Name
}

func receiverName(expression ast.Expr) string {
	switch expression := expression.(type) {
	case *ast.Ident:
		return expression.Name
	case *ast.StarExpr:
		return receiverName(expression.X)
	case *ast.IndexExpr:
		return receiverName(expression.X)
	case *ast.IndexListExpr:
		return receiverName(expression.X)
	default:
		return ""
	}
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

func functionMetrics(declaration *ast.FuncDecl, file *ast.File, fset *token.FileSet, start, end int) map[string]int {
	body := declaration.Body

	return map[string]int{
		"function_lines":        end - start + 1,
		"nesting_depth":         nesting(body),
		"cyclomatic_complexity": complexity(body),
		"function_arguments":    argumentCount(declaration.Type),
		"local_variables":       localVariables(body),
		"control_blocks":        controlBlocks(body),
		"return_points":         returnPoints(body),
		"boolean_operators":     booleanOperators(body),
		"max_condition_terms":   maxConditionTerms(body),
		"function_calls":        functionCalls(body),
		"literal_values":        literalValues(body),
		"closure_count":         closureCount(body),
		"comment_lines":         commentLines(file, fset, start, end),
		"statement_count":       statementCount(body),
		"type_dependencies":     typeDependencies(declaration),
		"exported_members":      0,
	}
}

func functionThresholds(thresholds config.Thresholds) map[string]int {
	return map[string]int{
		"function_lines":        thresholds.FunctionLines,
		"nesting_depth":         thresholds.NestingDepth,
		"cyclomatic_complexity": thresholds.CyclomaticComplexity,
		"function_arguments":    thresholds.FunctionArguments,
		"local_variables":       thresholds.LocalVariables,
		"control_blocks":        thresholds.ControlBlocks,
		"return_points":         thresholds.ReturnPoints,
		"boolean_operators":     thresholds.BooleanOperators,
		"max_condition_terms":   thresholds.MaxConditionTerms,
		"function_calls":        thresholds.FunctionCalls,
		"literal_values":        thresholds.LiteralValues,
		"closure_count":         thresholds.ClosureCount,
		"comment_lines":         thresholds.CommentLines,
		"statement_count":       thresholds.StatementCount,
		"type_dependencies":     thresholds.TypeDependencies,
		"exported_members":      thresholds.ExportedMembers,
	}
}

func candidate(kind, name, path string, start, end int, codeKind string, metrics, thresholds map[string]int, reasons []string) report.Candidate {
	return report.Candidate{
		Kind:       kind,
		Name:       name,
		Path:       path,
		StartLine:  start,
		EndLine:    end,
		CodeKind:   codeKind,
		Metrics:    metrics,
		Thresholds: thresholds,
		Reasons:    reasons,
	}
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

//nolint:cyclop // The AST cases enumerate supported local binding forms.
func localVariables(body *ast.BlockStmt) int {
	if body == nil {
		return 0
	}

	count := 0

	ast.Inspect(body, func(node ast.Node) bool {
		switch declaration := node.(type) {
		case *ast.FuncLit:
			return false
		case *ast.DeclStmt:
			genDecl, ok := declaration.Decl.(*ast.GenDecl)
			if !ok || genDecl.Tok != token.VAR {
				return true
			}

			for _, specification := range genDecl.Specs {
				value, ok := specification.(*ast.ValueSpec)
				if !ok {
					continue
				}

				count += namedIdentifiers(value.Names)
			}
		case *ast.AssignStmt:
			if declaration.Tok == token.DEFINE {
				count += newLocalVariables(declaration)
			}
		case *ast.RangeStmt:
			if declaration.Tok == token.DEFINE {
				count += namedExpressions([]ast.Expr{declaration.Key, declaration.Value})
			}
		}

		return true
	})

	return count
}

func namedIdentifiers(identifiers []*ast.Ident) int {
	count := 0

	for _, identifier := range identifiers {
		if identifier != nil && identifier.Name != "_" {
			count++
		}
	}

	return count
}

func newLocalVariables(declaration *ast.AssignStmt) int {
	count := 0

	for _, expression := range declaration.Lhs {
		identifier, ok := expression.(*ast.Ident)
		if ok && identifier.Name != "_" && identifier.Obj != nil && identifier.Obj.Decl == declaration {
			count++
		}
	}

	return count
}

func namedExpressions(expressions []ast.Expr) int {
	count := 0

	for _, expression := range expressions {
		identifier, ok := expression.(*ast.Ident)
		if ok && identifier.Name != "_" {
			count++
		}
	}

	return count
}

func controlBlocks(body *ast.BlockStmt) int {
	if body == nil {
		return 0
	}

	count := 0

	ast.Inspect(body, func(node ast.Node) bool {
		if _, ok := node.(*ast.FuncLit); ok {
			return false
		}

		switch node.(type) {
		case *ast.IfStmt, *ast.ForStmt, *ast.RangeStmt, *ast.SwitchStmt, *ast.TypeSwitchStmt, *ast.SelectStmt:
			count++
		}

		return true
	})

	return count
}

func returnPoints(body *ast.BlockStmt) int {
	return countNodes(body, func(node ast.Node) bool {
		_, ok := node.(*ast.ReturnStmt)

		return ok
	})
}

func booleanOperators(body *ast.BlockStmt) int {
	return countNodes(body, func(node ast.Node) bool {
		expression, ok := node.(*ast.BinaryExpr)

		return ok && (expression.Op == token.LAND || expression.Op == token.LOR)
	})
}

func maxConditionTerms(body *ast.BlockStmt) int {
	max := 0

	inspectBody(body, func(node ast.Node) bool {
		var condition ast.Expr

		switch statement := node.(type) {
		case *ast.IfStmt:
			condition = statement.Cond
		case *ast.ForStmt:
			condition = statement.Cond
		}

		if terms := conditionTerms(condition); terms > max {
			max = terms
		}

		return true
	})

	return max
}

func conditionTerms(expression ast.Expr) int {
	if expression == nil {
		return 0
	}

	binary, ok := expression.(*ast.BinaryExpr)
	if !ok || (binary.Op != token.LAND && binary.Op != token.LOR) {
		return 1
	}

	return conditionTerms(binary.X) + conditionTerms(binary.Y)
}

func functionCalls(body *ast.BlockStmt) int {
	return countNodes(body, func(node ast.Node) bool {
		_, ok := node.(*ast.CallExpr)

		return ok
	})
}

func literalValues(body *ast.BlockStmt) int {
	return countNodes(body, func(node ast.Node) bool {
		literal, ok := node.(*ast.BasicLit)

		return ok && literal.Value != "0" && literal.Value != "1" && literal.Value != `""`
	})
}

func closureCount(body *ast.BlockStmt) int {
	if body == nil {
		return 0
	}

	count := 0

	ast.Inspect(body, func(node ast.Node) bool {
		if _, ok := node.(*ast.FuncLit); ok {
			count++
			return false
		}

		return true
	})

	return count
}

func commentLines(file *ast.File, fset *token.FileSet, start, end int) int {
	lines := map[int]struct{}{}

	for _, group := range file.Comments {
		for _, comment := range group.List {
			commentStart := fset.Position(comment.Pos()).Line

			commentEnd := fset.Position(comment.End()).Line
			for line := max(start, commentStart); line <= min(end, commentEnd); line++ {
				lines[line] = struct{}{}
			}
		}
	}

	return len(lines)
}

func statementCount(body *ast.BlockStmt) int {
	return countNodes(body, func(node ast.Node) bool {
		statement, ok := node.(ast.Stmt)
		if !ok {
			return false
		}

		switch statement.(type) {
		case *ast.BlockStmt, *ast.EmptyStmt:
			return false
		default:
			return true
		}
	})
}

func typeDependencies(declaration *ast.FuncDecl) int {
	types := map[string]struct{}{}

	for _, fieldList := range []*ast.FieldList{declaration.Recv, declaration.Type.Params, declaration.Type.Results} {
		if fieldList == nil {
			continue
		}

		for _, field := range fieldList.List {
			if typeName := typeExpression(field.Type); typeName != "" {
				types[typeName] = struct{}{}
			}
		}
	}

	return len(types)
}

func typeExpression(expression ast.Expr) string {
	var buffer bytes.Buffer
	if err := format.Node(&buffer, token.NewFileSet(), expression); err != nil {
		return ""
	}

	return buffer.String()
}

func countNodes(body *ast.BlockStmt, matches func(ast.Node) bool) int {
	count := 0

	inspectBody(body, func(node ast.Node) bool {
		if matches(node) {
			count++
		}

		return true
	})

	return count
}

func inspectBody(body *ast.BlockStmt, visit func(ast.Node) bool) {
	if body == nil {
		return
	}

	ast.Inspect(body, func(node ast.Node) bool {
		if _, ok := node.(*ast.FuncLit); ok {
			return false
		}

		return visit(node)
	})
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

func exportedMethodCounts(file *ast.File) map[string]int {
	counts := map[string]int{}

	for _, declaration := range file.Decls {
		function, ok := declaration.(*ast.FuncDecl)
		if !ok || !function.Name.IsExported() || function.Recv == nil || len(function.Recv.List) == 0 {
			continue
		}

		if receiver := receiverName(function.Recv.List[0].Type); receiver != "" {
			counts[receiver]++
		}
	}

	return counts
}

func exportedFields(structure *ast.StructType) int {
	count := 0

	for _, field := range structure.Fields.List {
		if len(field.Names) == 0 {
			if ast.IsExported(receiverName(field.Type)) {
				count++
			}

			continue
		}

		for _, name := range field.Names {
			if name.IsExported() {
				count++
			}
		}
	}

	return count
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
