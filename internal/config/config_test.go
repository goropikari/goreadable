package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/goropikari/goreadable/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaults(t *testing.T) {
	t.Run("when defaults are requested: includes local variable and control block thresholds", func(t *testing.T) {
		// Arrange

		// Act
		thresholds := config.Defaults()

		// Assert
		assert.Equal(t, 15, thresholds.LocalVariables)
		assert.Equal(t, 8, thresholds.ControlBlocks)
		assert.Equal(t, 5, thresholds.ReturnPoints)
		assert.Equal(t, 8, thresholds.BooleanOperators)
		assert.Equal(t, 4, thresholds.MaxConditionTerms)
		assert.Equal(t, 15, thresholds.FunctionCalls)
		assert.Equal(t, 10, thresholds.LiteralValues)
		assert.Equal(t, 2, thresholds.ClosureCount)
		assert.Equal(t, 10, thresholds.CommentLines)
		assert.Equal(t, 40, thresholds.StatementCount)
		assert.Equal(t, 5, thresholds.TypeDependencies)
		assert.Equal(t, 10, thresholds.ExportedMembers)
	})
}

func TestLoadFile(t *testing.T) {
	t.Run("when a YAML file defines one threshold: merges it with defaults", func(t *testing.T) {
		// Arrange
		path := filepath.Join(t.TempDir(), "goreadable.yaml")
		require.NoError(t, os.WriteFile(path, []byte("thresholds:\n  function_lines: 60\n"), 0o600))

		defaults := config.Defaults()

		// Act
		got, err := config.LoadFile(path, defaults)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, 60, got.FunctionLines)
		assert.Equal(t, defaults.NestingDepth, got.NestingDepth)
	})

	t.Run("when the YAML file is malformed: returns a path-aware error", func(t *testing.T) {
		// Arrange
		path := filepath.Join(t.TempDir(), "goreadable.yaml")
		require.NoError(t, os.WriteFile(path, []byte("thresholds: ["), 0o600))

		// Act
		_, err := config.LoadFile(path, config.Defaults())

		// Assert
		require.Error(t, err)
		assert.Contains(t, err.Error(), path)
	})
}
