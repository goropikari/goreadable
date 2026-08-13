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
	t.Run("when all functions are requested: reports every package function regardless of thresholds", func(t *testing.T) {
		// Arrange
		root := t.TempDir()
		path := filepath.Join(root, "sample.go")
		require.NoError(t, os.WriteFile(path, []byte(`package alpha

func Small() {}

	func Other(value int) int { return value }
`), 0o600))

		var stdout, stderr bytes.Buffer

		// Act
		err := execute([]string{"--format", "json", "--all-functions", "--max-function-lines", "100", root}, &stdout, &stderr)

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
		assert.Equal(t, map[string]int{"function_lines": 1, "nesting_depth": 0, "cyclomatic_complexity": 1, "function_arguments": 0}, result.Candidates[0].Metrics)
		assert.Equal(t, "Other", result.Candidates[1].Name)
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
		err := execute([]string{"--all-functions", root}, &stdout, &stderr)

		// Assert
		require.NoError(t, err, stderr.String())
		assert.Contains(t, stdout.String(), "cyclomatic_complexity=1")
		assert.Contains(t, stdout.String(), "function_arguments=0")
		assert.Contains(t, stdout.String(), "function_lines=1")
		assert.Contains(t, stdout.String(), "nesting_depth=0")
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
}
