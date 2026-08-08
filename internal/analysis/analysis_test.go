package analysis_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/goropikari/goreadable/internal/analysis"
	"github.com/goropikari/goreadable/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAnalyze(t *testing.T) {
	t.Run("when a function has goreadable ignore comment: excludes that function", func(t *testing.T) {
		// Arrange
		root := t.TempDir()
		source := `package sample

// goreadable:ignore
func Ignored() int {
		value := 0
		value++
		value++
		return value
}

func Reported() int {
		value := 0
		value++
		value++
		return value
}
`
		path := filepath.Join(root, "sample.go")
		require.NoError(t, os.WriteFile(path, []byte(source), 0o600))

		files, err := analysis.Files(root, false)
		require.NoError(t, err)

		thresholds := config.Defaults()
		thresholds.FunctionLines = 3

		// Act
		result, err := analysis.Analyze(files, thresholds, nil)

		// Assert
		require.NoError(t, err)
		require.Len(t, result.Candidates, 1)
		assert.Equal(t, "Reported", result.Candidates[0].Name)
	})

	t.Run("when a type declaration has block ignore comment: excludes that type", func(t *testing.T) {
		// Arrange
		root := t.TempDir()
		source := `package sample

/* goreadable:ignore */
type Ignored struct {
		A int
		B int
		C int
		D int
}

type Reported struct {
		A int
		B int
		C int
		D int
}
`
		path := filepath.Join(root, "sample.go")
		require.NoError(t, os.WriteFile(path, []byte(source), 0o600))

		files, err := analysis.Files(root, false)
		require.NoError(t, err)

		thresholds := config.Defaults()
		thresholds.StructFields = 3

		// Act
		result, err := analysis.Analyze(files, thresholds, nil)

		// Assert
		require.NoError(t, err)
		require.Len(t, result.Candidates, 1)
		assert.Equal(t, "Reported", result.Candidates[0].Name)
	})
}
