package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExecute(t *testing.T) {
	t.Run("when help is requested: explains the primary workflows", func(t *testing.T) {
		// Arrange
		var stdout, stderr bytes.Buffer

		// Act
		err := execute([]string{"--help"}, &stdout, &stderr)

		// Assert
		require.NoError(t, err, stderr.String())

		for _, expected := range []string{
			"With no path, it analyzes the current directory.",
			"Use a path ending in /... to analyze that directory recursively.",
			"Review candidates prioritize human or AI review; they do not make the command fail.",
			"CLI flags override goreadable.yaml, which overrides defaults.",
			"goreadable --format json ./...",
			"goreadable --thresholds-only ./...",
			"goreadable --function package.Function ./...",
			"goreadable --diff HEAD ./...",
		} {
			assert.Contains(t, stdout.String(), expected)
		}
	})

	t.Run("when no filter is requested: reports every package function regardless of thresholds", func(t *testing.T) {
		// Arrange
		root := t.TempDir()
		path := filepath.Join(root, "sample.go")
		require.NoError(t, os.WriteFile(path, []byte(`package alpha

func Small() {}

	func Other(value int) int { return value }
`), 0o600))

		var stdout, stderr bytes.Buffer

		// Act
		err := execute([]string{"--format", "json", "--max-function-lines", "100", root}, &stdout, &stderr)

		// Assert
		require.NoError(t, err, stderr.String())

		var result struct {
			Candidates []struct {
				Name    string         `json:"name"`
				Metrics map[string]int `json:"metrics"`
			}
		}
		require.NoError(t, json.Unmarshal(stdout.Bytes(), &result))
		require.Len(t, result.Candidates, 2)
		assert.Equal(t, "Small", result.Candidates[0].Name)
		assert.Equal(t, map[string]int{"function_lines": 1, "nesting_depth": 0, "cyclomatic_complexity": 1, "function_arguments": 0, "local_variables": 0, "control_blocks": 0, "return_points": 0, "boolean_operators": 0, "max_condition_terms": 0, "function_calls": 0, "literal_values": 0, "closure_count": 0, "comment_lines": 0, "statement_count": 0, "type_dependencies": 0, "exported_members": 0}, result.Candidates[0].Metrics)
		assert.Equal(t, "Other", result.Candidates[1].Name)
	})

	t.Run("when threshold filtering is requested: reports only declarations exceeding thresholds", func(t *testing.T) {
		// Arrange
		root := t.TempDir()
		require.NoError(t, os.WriteFile(filepath.Join(root, "sample.go"), []byte(`package alpha

type Exported struct {
	First  int
	Second int
}

func Small() {}

func Large() {
	first := 1
	second := 2
	_ = first + second
}
`), 0o600))

		var stdout, stderr bytes.Buffer

		// Act
		err := execute([]string{"--format", "json", "--thresholds-only", "--max-function-lines", "1", "--max-struct-fields", "1", root}, &stdout, &stderr)

		// Assert
		require.NoError(t, err, stderr.String())

		var result struct {
			Candidates []struct {
				Kind    string   `json:"kind"`
				Name    string   `json:"name"`
				Reasons []string `json:"reasons"`
			}
		}
		require.NoError(t, json.Unmarshal(stdout.Bytes(), &result))
		require.Len(t, result.Candidates, 2)
		assert.Equal(t, "type", result.Candidates[0].Kind)
		assert.Equal(t, "Exported", result.Candidates[0].Name)
		assert.NotEmpty(t, result.Candidates[0].Reasons)
		assert.Equal(t, "function", result.Candidates[1].Kind)
		assert.Equal(t, "Large", result.Candidates[1].Name)
		assert.NotEmpty(t, result.Candidates[1].Reasons)
	})

	t.Run("when the removed all-functions option is requested: returns a recoverable error", func(t *testing.T) {
		// Arrange
		root := t.TempDir()
		require.NoError(t, os.WriteFile(filepath.Join(root, "sample.go"), []byte("package alpha\n\nfunc Small() {}\n"), 0o600))

		var stdout, stderr bytes.Buffer

		// Act
		err := execute([]string{"--format", "json", "--all-functions", root}, &stdout, &stderr)

		// Assert
		require.Error(t, err)
		assert.Contains(t, err.Error(), "--all-functions")
	})

	t.Run("when a package-qualified function is requested: excludes same-named functions from another package", func(t *testing.T) {
		// Arrange
		root := t.TempDir()
		require.NoError(t, os.WriteFile(filepath.Join(root, "alpha.go"), []byte("package alpha\n\nfunc Target() {}\n"), 0o600))
		require.NoError(t, os.WriteFile(filepath.Join(root, "beta.go"), []byte("package beta\n\nfunc Target() {}\n\nfunc Other() {}\n"), 0o600))

		var stdout, stderr bytes.Buffer

		// Act
		err := execute([]string{"--format", "json", "--function", "alpha.Target", root}, &stdout, &stderr)

		// Assert
		require.NoError(t, err, stderr.String())

		var result struct {
			Candidates []struct {
				Name string `json:"name"`
			}
		}
		require.NoError(t, json.Unmarshal(stdout.Bytes(), &result))
		require.Len(t, result.Candidates, 1)
		assert.Equal(t, "Target", result.Candidates[0].Name)
	})

	t.Run("when a package-qualified method is requested: reports the method", func(t *testing.T) {
		// Arrange
		root := t.TempDir()
		require.NoError(t, os.WriteFile(filepath.Join(root, "sample.go"), []byte("package alpha\n\ntype Thing struct{}\n\nfunc (Thing) Do() {}\n"), 0o600))

		var stdout, stderr bytes.Buffer

		// Act
		err := execute([]string{"--format", "json", "--function", "alpha.Thing.Do", root}, &stdout, &stderr)

		// Assert
		require.NoError(t, err, stderr.String())

		var result struct {
			Candidates []struct {
				Name string `json:"name"`
			}
		}
		require.NoError(t, json.Unmarshal(stdout.Bytes(), &result))
		require.Len(t, result.Candidates, 1)
		assert.Equal(t, "Do", result.Candidates[0].Name)
	})

	t.Run("when metrics are requested as text: includes every function metric", func(t *testing.T) {
		// Arrange
		root := t.TempDir()
		require.NoError(t, os.WriteFile(filepath.Join(root, "sample.go"), []byte("package alpha\n\nfunc Small() {}\n"), 0o600))

		var stdout, stderr bytes.Buffer

		// Act
		err := execute([]string{root}, &stdout, &stderr)

		// Assert
		require.NoError(t, err, stderr.String())
		assert.Contains(t, stdout.String(), "cyclomatic_complexity=1")
		assert.Contains(t, stdout.String(), "function_arguments=0")
		assert.Contains(t, stdout.String(), "function_lines=1")
		assert.Contains(t, stdout.String(), "nesting_depth=0")
		assert.Contains(t, stdout.String(), "local_variables=0")
		assert.Contains(t, stdout.String(), "control_blocks=0")
		assert.Contains(t, stdout.String(), "return_points=0")
		assert.Contains(t, stdout.String(), "boolean_operators=0")
		assert.Contains(t, stdout.String(), "max_condition_terms=0")
		assert.Contains(t, stdout.String(), "function_calls=0")
		assert.Contains(t, stdout.String(), "literal_values=0")
		assert.Contains(t, stdout.String(), "closure_count=0")
		assert.Contains(t, stdout.String(), "comment_lines=0")
		assert.Contains(t, stdout.String(), "statement_count=0")
		assert.Contains(t, stdout.String(), "type_dependencies=0")
		assert.Contains(t, stdout.String(), "exported_members=0")
	})

	t.Run("when local variable and control block thresholds are configured: flag values override the configuration file", func(t *testing.T) {
		// Arrange
		root := t.TempDir()
		require.NoError(t, os.WriteFile(filepath.Join(root, "sample.go"), []byte(`package alpha

func Target() {
	first := 1
	second := 2
	if first < second {
		return
	}
}
`), 0o600))
		require.NoError(t, os.WriteFile(filepath.Join(root, "goreadable.yaml"), []byte("thresholds:\n  local_variables: 1\n  control_blocks: 1\n"), 0o600))

		var stdout, stderr bytes.Buffer

		// Act
		err := execute([]string{"--format", "json", "--max-local-variables", "2", "--max-control-blocks", "2", root}, &stdout, &stderr)

		// Assert
		require.NoError(t, err, stderr.String())

		var result struct {
			Candidates []struct {
				Metrics    map[string]int `json:"metrics"`
				Thresholds map[string]int `json:"thresholds"`
			}
		}
		require.NoError(t, json.Unmarshal(stdout.Bytes(), &result))
		require.Len(t, result.Candidates, 1)
		assert.Equal(t, 2, result.Candidates[0].Metrics["local_variables"])
		assert.Equal(t, 1, result.Candidates[0].Metrics["control_blocks"])
		assert.Equal(t, 2, result.Candidates[0].Thresholds["local_variables"])
		assert.Equal(t, 2, result.Candidates[0].Thresholds["control_blocks"])
	})

	t.Run("when an empty function selector is requested: returns a recoverable error", func(t *testing.T) {
		// Arrange
		var stdout, stderr bytes.Buffer

		// Act
		err := execute([]string{"--function", ""}, &stdout, &stderr)

		// Assert
		require.Error(t, err)
		assert.Contains(t, err.Error(), "--function")
	})

	t.Run("when YAML and legacy JSON both configure a threshold: YAML takes precedence", func(t *testing.T) {
		// Arrange
		root := t.TempDir()
		require.NoError(t, os.WriteFile(filepath.Join(root, "sample.go"), []byte("package alpha\n\nfunc Target() {\n}\n"), 0o600))
		require.NoError(t, os.WriteFile(filepath.Join(root, "goreadable.yaml"), []byte("thresholds:\n  function_lines: 1\n"), 0o600))
		require.NoError(t, os.WriteFile(filepath.Join(root, "goreadable.json"), []byte(`{"thresholds":{"function_lines":100}}`), 0o600))

		var stdout, stderr bytes.Buffer

		// Act
		err := execute([]string{"--format", "json", "--thresholds-only", root}, &stdout, &stderr)

		// Assert
		require.NoError(t, err, stderr.String())
		assert.Contains(t, stdout.String(), `"function_lines": 1`)
	})
}
