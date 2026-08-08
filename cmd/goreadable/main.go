package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/goropikari/goreadable/internal/analysis"
	"github.com/goropikari/goreadable/internal/config"
	"github.com/goropikari/goreadable/internal/diff"
	"github.com/goropikari/goreadable/internal/report"
)

func main() {
	if err := run(os.Args[1:], os.Stdout, os.Stderr); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

//nolint:cyclop // The CLI boundary intentionally coordinates all user-facing options.
func run(arguments []string, stdout, stderr *os.File) error {
	flags := flag.NewFlagSet("goreadable", flag.ContinueOnError)
	flags.SetOutput(stderr)
	format := flags.String("format", "text", "output format: text or json")
	diffRef := flags.String("diff", "", "analyze declarations changed since this Git ref")

	flagValues := map[string]*int{"function_lines": flags.Int("max-function-lines", 0, "maximum function lines"), "nesting_depth": flags.Int("max-nesting-depth", 0, "maximum nesting depth"), "cyclomatic_complexity": flags.Int("max-cyclomatic-complexity", 0, "maximum cyclomatic complexity"), "function_arguments": flags.Int("max-function-args", 0, "maximum function arguments"), "struct_fields": flags.Int("max-struct-fields", 0, "maximum struct fields"), "type_methods": flags.Int("max-type-methods", 0, "maximum methods on a type")}
	if err := flags.Parse(arguments); err != nil {
		return err
	}

	if *format != "text" && *format != "json" {
		return fmt.Errorf("invalid format %q: use text or json", *format)
	}

	paths := flags.Args()
	if len(paths) == 0 {
		paths = []string{"."}
	}

	thresholds := config.Defaults()
	if fileThresholds, err := config.LoadFile(filepath.Join(pathsRoot(paths), "goreadable.json"), thresholds); err != nil {
		return err
	} else {
		thresholds = fileThresholds
	}

	overrides := map[string]int{}

	for key, value := range flagValues {
		if *value != 0 {
			overrides[key] = *value
		}
	}

	thresholds.ApplyFlags(overrides)

	changed := map[string][][2]int(nil)

	if *diffRef != "" {
		var err error

		changed, err = diff.ChangedFiles(pathsRoot(paths), *diffRef)
		if err != nil {
			return fmt.Errorf("git diff: %w", err)
		}
	}

	var files []string

	for _, path := range paths {
		recursive := strings.HasSuffix(path, "/...") || path == "./..."

		root := strings.TrimSuffix(strings.TrimSuffix(path, "/..."), "...")
		if root == "" {
			root = "."
		}

		found, err := analysis.Files(root, recursive)
		if err != nil {
			return err
		}

		files = append(files, found...)
	}

	result, err := analysis.Analyze(files, thresholds, changed)
	if err != nil {
		return err
	}

	if *format == "json" {
		return report.WriteJSON(stdout, result)
	}

	return report.WriteText(stdout, result)
}

func pathsRoot(paths []string) string {
	if len(paths) == 0 {
		return "."
	}

	root := paths[0]
	root = strings.TrimSuffix(root, "/...")

	if root == "" {
		return "."
	}

	return root
}
