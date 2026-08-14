package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/goropikari/goreadable/internal/analysis"
	"github.com/goropikari/goreadable/internal/config"
	"github.com/goropikari/goreadable/internal/diff"
	"github.com/goropikari/goreadable/internal/report"
	"github.com/spf13/cobra"
)

func main() {
	if err := execute(os.Args[1:], os.Stdout, os.Stderr); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func execute(arguments []string, stdout, stderr io.Writer) error {
	command := newCommand(stdout, stderr)
	command.SetArgs(arguments)

	return command.Execute()
}

//nolint:cyclop // Command setup keeps all user-visible options at one boundary.
func newCommand(stdout, stderr io.Writer) *cobra.Command {
	var (
		format, diffRef   string
		thresholdsOnly    bool
		functionSelectors []string
	)

	thresholds := config.Defaults()

	command := &cobra.Command{
		Use:   "goreadable [paths...]",
		Short: "find Go declarations that deserve a readability review",
		Long: `goreadable reports structural metrics for Go functions and can filter declarations whose metrics exceed readability thresholds.

With no path, it analyzes the current directory. Use a path ending in /... to analyze that directory recursively. Review candidates prioritize human or AI review; they do not make the command fail.

Thresholds resolve in this order: CLI flags override goreadable.json, which overrides defaults. Use --format json when another tool or AI needs the metrics, thresholds, and selection reasons.`,
		Example: `  # Review the current directory or a directory tree.
  goreadable
  goreadable ./...

  # Send an explainable JSON report to an AI or another tool.
  goreadable --format json ./...

  # Review only declarations whose metrics exceed their thresholds.
  goreadable --thresholds-only ./...

  # Report one function (or package.Type.Method).
  goreadable --function package.Function ./...

  # Review only declarations changed since a Git reference.
  goreadable --diff HEAD ./...

  # Configure thresholds in goreadable.json at the analysis root.
  # CLI flags override goreadable.json, which overrides defaults.`,
		SilenceErrors: true,
		SilenceUsage:  true,
		Args:          cobra.ArbitraryArgs,
		RunE: func(command *cobra.Command, paths []string) error {
			if format != "text" && format != "json" {
				return fmt.Errorf("invalid format %q: use text or json", format)
			}

			if len(paths) == 0 {
				paths = []string{"."}
			}

			if len(functionSelectors) > 0 && thresholdsOnly {
				return fmt.Errorf("--function cannot be used with --thresholds-only")
			}

			options, err := analysis.NewOptions(thresholdsOnly, functionSelectors)
			if err != nil {
				return err
			}

			resolved := thresholds

			fileThresholds, err := config.LoadFile(filepath.Join(pathsRoot(paths), "goreadable.json"), resolved)
			if err != nil {
				return err
			}

			resolved = fileThresholds
			resolved.ApplyFlags(flagOverrides(command))

			changed := map[string][][2]int(nil)
			if diffRef != "" {
				changed, err = diff.ChangedFiles(pathsRoot(paths), diffRef)
				if err != nil {
					return fmt.Errorf("git diff: %w", err)
				}
			}

			files, err := inputFiles(paths)
			if err != nil {
				return err
			}

			result, err := analysis.AnalyzeWithOptions(files, resolved, changed, options)
			if err != nil {
				return err
			}

			if format == "json" {
				return report.WriteJSON(stdout, result)
			}

			return report.WriteText(stdout, result)
		},
	}
	command.SetOut(stdout)
	command.SetErr(stderr)
	command.Flags().StringVar(&format, "format", "text", "output format: text or json")
	command.Flags().StringVar(&diffRef, "diff", "", "analyze declarations changed since this Git ref")
	command.Flags().BoolVar(&thresholdsOnly, "thresholds-only", false, "report only declarations whose metrics exceed thresholds")
	command.Flags().StringArrayVar(&functionSelectors, "function", nil, "report metrics for package.Function or package.Type.Method (repeatable)")
	command.Flags().Int("max-function-lines", thresholds.FunctionLines, "maximum function lines")
	command.Flags().Int("max-nesting-depth", thresholds.NestingDepth, "maximum nesting depth")
	command.Flags().Int("max-cyclomatic-complexity", thresholds.CyclomaticComplexity, "maximum cyclomatic complexity")
	command.Flags().Int("max-function-args", thresholds.FunctionArguments, "maximum function arguments")
	command.Flags().Int("max-local-variables", thresholds.LocalVariables, "maximum local variables")
	command.Flags().Int("max-control-blocks", thresholds.ControlBlocks, "maximum control blocks")
	command.Flags().Int("max-return-points", thresholds.ReturnPoints, "maximum return points")
	command.Flags().Int("max-boolean-operators", thresholds.BooleanOperators, "maximum boolean operators")
	command.Flags().Int("max-condition-terms", thresholds.MaxConditionTerms, "maximum terms in one condition")
	command.Flags().Int("max-function-calls", thresholds.FunctionCalls, "maximum function calls")
	command.Flags().Int("max-literal-values", thresholds.LiteralValues, "maximum meaningful literal values")
	command.Flags().Int("max-closures", thresholds.ClosureCount, "maximum function literals")
	command.Flags().Int("max-comment-lines", thresholds.CommentLines, "maximum comment lines")
	command.Flags().Int("max-statements", thresholds.StatementCount, "maximum statements")
	command.Flags().Int("max-type-dependencies", thresholds.TypeDependencies, "maximum distinct signature type dependencies")
	command.Flags().Int("max-struct-fields", thresholds.StructFields, "maximum struct fields")
	command.Flags().Int("max-type-methods", thresholds.TypeMethods, "maximum methods on a type")
	command.Flags().Int("max-exported-members", thresholds.ExportedMembers, "maximum exported fields and methods on a type")

	return command
}

func flagOverrides(command *cobra.Command) map[string]int {
	overrides := map[string]int{}

	flags := map[string]string{
		"max-function-lines":        "function_lines",
		"max-nesting-depth":         "nesting_depth",
		"max-cyclomatic-complexity": "cyclomatic_complexity",
		"max-function-args":         "function_arguments",
		"max-local-variables":       "local_variables",
		"max-control-blocks":        "control_blocks",
		"max-return-points":         "return_points",
		"max-boolean-operators":     "boolean_operators",
		"max-condition-terms":       "max_condition_terms",
		"max-function-calls":        "function_calls",
		"max-literal-values":        "literal_values",
		"max-closures":              "closure_count",
		"max-comment-lines":         "comment_lines",
		"max-statements":            "statement_count",
		"max-type-dependencies":     "type_dependencies",
		"max-struct-fields":         "struct_fields",
		"max-type-methods":          "type_methods",
		"max-exported-members":      "exported_members",
	}
	for flagName, thresholdName := range flags {
		if !command.Flags().Changed(flagName) {
			continue
		}

		value, err := command.Flags().GetInt(flagName)
		if err == nil {
			overrides[thresholdName] = value
		}
	}

	return overrides
}

func inputFiles(paths []string) ([]string, error) {
	var files []string

	for _, path := range paths {
		recursive := strings.HasSuffix(path, "/...") || path == "./..."

		root := strings.TrimSuffix(strings.TrimSuffix(path, "/..."), "...")
		if root == "" {
			root = "."
		}

		found, err := analysis.Files(root, recursive)
		if err != nil {
			return nil, err
		}

		files = append(files, found...)
	}

	return files, nil
}

func pathsRoot(paths []string) string {
	if len(paths) == 0 {
		return "."
	}

	root := strings.TrimSuffix(paths[0], "/...")
	if root == "" {
		return "."
	}

	return root
}
