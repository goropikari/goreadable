package acceptance

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGoreadableModuleIdentity(t *testing.T) {
	t.Run("when the module and command are inspected: exposes the documented consumer identity", func(t *testing.T) {
		// Arrange
		repositoryRoot := repositoryRoot(t)

		// Act
		module := runCommand(t, repositoryRoot, "go", "list", "-m", "-f", "{{.Path}}")
		build := exec.Command("go", "build", "./cmd/goreadable")
		build.Dir = repositoryRoot
		buildOutput, buildErr := build.CombinedOutput()

		// Assert
		require.NoError(t, buildErr, string(buildOutput))
		assert.Equal(t, "github.com/goropikari/goreadable", strings.TrimSpace(module))
	})
}

func TestGoreadableCLI(t *testing.T) {
	t.Run("when a package has over-threshold function and type: reports review candidates", func(t *testing.T) {
		// Arrange
		binary := buildGoreadable(t)
		fixture := writeFixture(t)

		// Act
		output, errOutput, err := runGoreadable(binary, fixture, "--format", "json", "--max-function-lines", "3", "--max-struct-fields", "3", ".")

		// Assert
		require.NoError(t, err, errOutput)
		candidates := candidates(t, output)
		assert.Contains(t, candidateNames(candidates), "TooLong")
		assert.Contains(t, candidateNames(candidates), "TooWide")
	})

	t.Run("when JSON is requested: includes version metrics and reasons without source", func(t *testing.T) {
		// Arrange
		binary := buildGoreadable(t)
		fixture := writeFixture(t)

		// Act
		output, errOutput, err := runGoreadable(binary, fixture, "--format", "json", "--max-function-lines", "3", ".")

		// Assert
		require.NoError(t, err, errOutput)
		var report map[string]any
		require.NoError(t, json.Unmarshal([]byte(output), &report))
		assert.Equal(t, float64(1), report["version"])

		candidate := candidateByName(t, candidates(t, output), "TooLong")
		for _, field := range []string{"kind", "name", "path", "start_line", "end_line", "code_kind", "metrics", "thresholds", "reasons"} {
			assert.NotEmpty(t, candidate[field], field)
		}
		assert.NotContains(t, candidate, "source")
	})

	t.Run("when text output is requested: identifies each candidate and reason", func(t *testing.T) {
		// Arrange
		binary := buildGoreadable(t)
		fixture := writeFixture(t)

		// Act
		output, errOutput, err := runGoreadable(binary, fixture, "--max-function-lines", "3", ".")

		// Assert
		require.NoError(t, err, errOutput)
		assert.Contains(t, output, "TooLong")
		assert.Contains(t, output, "function_lines")
		assert.Contains(t, output, "fixture.go")
	})

	t.Run("when configuration and a flag set the same threshold: the flag wins", func(t *testing.T) {
		// Arrange
		binary := buildGoreadable(t)
		fixture := writeFixture(t)
		require.NoError(t, os.WriteFile(filepath.Join(fixture, "goreadable.json"), []byte(`{"thresholds":{"function_lines":3}}`), 0o600))

		// Act
		configuredOutput, configuredErrOutput, configuredErr := runGoreadable(binary, fixture, "--format", "json", ".")
		flagOutput, flagErrOutput, flagErr := runGoreadable(binary, fixture, "--format", "json", "--max-function-lines", "100", ".")

		// Assert
		require.NoError(t, configuredErr, configuredErrOutput)
		require.NoError(t, flagErr, flagErrOutput)
		assert.Contains(t, candidateNames(candidates(t, configuredOutput)), "TooLong")
		assert.NotContains(t, candidateNames(candidates(t, flagOutput)), "TooLong")
	})

	t.Run("when diff mode is given a Git ref: reports candidates in changed code only", func(t *testing.T) {
		// Arrange
		binary := buildGoreadable(t)
		fixture := writeFixture(t)
		runCommand(t, fixture, "git", "init", "--quiet")
		runCommand(t, fixture, "git", "add", ".")
		runCommand(t, fixture, "git", "-c", "user.name=Acceptance Test", "-c", "user.email=acceptance@example.test", "commit", "--quiet", "-m", "baseline")
		require.NoError(t, os.WriteFile(filepath.Join(fixture, "changed.go"), []byte("package fixture\n\nfunc Changed() int {\n\tvalue := 0\n\tvalue++\n\tvalue++\n\treturn value\n}\n"), 0o600))

		// Act
		output, errOutput, err := runGoreadable(binary, fixture, "--format", "json", "--diff", "HEAD", "--max-function-lines", "3", ".")

		// Assert
		require.NoError(t, err, errOutput)
		assert.Contains(t, candidateNames(candidates(t, output)), "Changed")
		assert.NotContains(t, candidateNames(candidates(t, output)), "TooLong")
	})

	t.Run("when candidates exist: exits zero", func(t *testing.T) {
		// Arrange
		binary := buildGoreadable(t)
		fixture := writeFixture(t)

		// Act
		_, errOutput, err := runGoreadable(binary, fixture, "--max-function-lines", "3", ".")

		// Assert
		require.NoError(t, err, errOutput)
	})

	t.Run("when an option is invalid: exits non-zero with a recoverable error", func(t *testing.T) {
		// Arrange
		binary := buildGoreadable(t)
		fixture := writeFixture(t)

		// Act
		_, errOutput, err := runGoreadable(binary, fixture, "--format", "unknown", ".")

		// Assert
		require.Error(t, err)
		assert.Contains(t, errOutput, "format")
	})
}

func buildGoreadable(t *testing.T) string {
	t.Helper()

	binary := filepath.Join(t.TempDir(), "goreadable")
	command := exec.Command("go", "build", "-o", binary, "./cmd/goreadable")
	command.Dir = repositoryRoot(t)
	output, err := command.CombinedOutput()
	require.NoError(t, err, string(output))
	return binary
}

func writeFixture(t *testing.T) string {
	t.Helper()

	fixture := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(fixture, "go.mod"), []byte("module example.com/fixture\n\ngo 1.26\n"), 0o600))
	require.NoError(t, os.WriteFile(filepath.Join(fixture, "fixture.go"), []byte(`package fixture

type TooWide struct {
	First  int
	Second int
	Third  int
	Fourth int
}

func TooLong(first, second, third, fourth int) int {
	total := first + second
	total += third
	total += fourth
	return total
}
`), 0o600))
	return fixture
}

func runGoreadable(binary, directory string, arguments ...string) (string, string, error) {
	command := exec.Command(binary, arguments...)
	command.Dir = directory
	var stdout strings.Builder
	var stderr strings.Builder
	command.Stdout = &stdout
	command.Stderr = &stderr
	err := command.Run()
	return stdout.String(), stderr.String(), err
}

func candidates(t *testing.T, output string) []map[string]any {
	t.Helper()
	var report struct {
		Candidates []map[string]any `json:"candidates"`
	}
	require.NoError(t, json.Unmarshal([]byte(output), &report))
	return report.Candidates
}

func candidateNames(candidates []map[string]any) []string {
	names := make([]string, 0, len(candidates))
	for _, candidate := range candidates {
		name, _ := candidate["name"].(string)
		names = append(names, name)
	}
	return names
}

func candidateByName(t *testing.T, candidates []map[string]any, name string) map[string]any {
	t.Helper()
	for _, candidate := range candidates {
		if candidateName, _ := candidate["name"].(string); candidateName == name {
			return candidate
		}
	}
	require.Failf(t, "candidate not found", "candidate %q was not present", name)
	return nil
}

func repositoryRoot(t *testing.T) string {
	t.Helper()
	return filepath.Clean(filepath.Join("..", ".."))
}

func runCommand(t *testing.T, directory, name string, arguments ...string) string {
	t.Helper()
	command := exec.Command(name, arguments...)
	command.Dir = directory
	output, err := command.CombinedOutput()
	require.NoError(t, err, string(output))
	return string(output)
}
