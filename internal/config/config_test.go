package config_test

import (
	"testing"

	"github.com/goropikari/goreadable/internal/config"
	"github.com/stretchr/testify/assert"
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
