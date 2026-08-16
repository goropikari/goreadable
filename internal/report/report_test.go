package report_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJSONSchema(t *testing.T) {
	t.Run("when the checked-in output schema is loaded: defines the versioned report contract", func(t *testing.T) {
		// Arrange
		path := filepath.Join("..", "..", "schemas", "goreadable-output.schema.json")
		data, err := os.ReadFile(path)
		require.NoError(t, err)

		var schema map[string]any

		// Act
		require.NoError(t, json.Unmarshal(data, &schema))

		// Assert
		assert.Equal(t, "https://json-schema.org/draft/2020-12/schema", schema["$schema"])
		assert.Equal(t, "goreadable JSON output", schema["title"])
		assert.Equal(t, []any{"version", "candidates"}, schema["required"])

		definitions, ok := schema["$defs"].(map[string]any)
		require.True(t, ok)
		assert.Contains(t, definitions, "functionMetrics")
		assert.Contains(t, definitions, "functionThresholds")
		assert.Contains(t, definitions, "typeMetrics")
		assert.Contains(t, definitions, "typeThresholds")
	})
}
