package diff

import (
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var hunk = regexp.MustCompile(`@@ -\d+(?:,\d+)? \+(\d+)(?:,(\d+))? @@`)

func ChangedFiles(root, ref string) (map[string][][2]int, error) {
	result, err := changedTrackedFiles(root, ref)
	if err != nil {
		return nil, err
	}

	return addUntrackedFiles(root, result)
}

func changedTrackedFiles(root, ref string) (map[string][][2]int, error) {
	command := exec.Command("git", "diff", "--unified=0", ref, "--", "*.go")
	command.Dir = root

	output, err := command.Output()
	if err != nil {
		return nil, err
	}

	result := map[string][][2]int{}
	current := ""

	for line := range strings.SplitSeq(string(output), "\n") {
		if after, ok := strings.CutPrefix(line, "+++ b/"); ok {
			current = filepath.Join(root, after)
			continue
		}

		match := hunk.FindStringSubmatch(line)
		if match == nil || current == "" {
			continue
		}

		start, _ := strconv.Atoi(match[1])

		count := 1
		if match[2] != "" {
			count, _ = strconv.Atoi(match[2])
			if count == 0 {
				count = 1
			}
		}

		result[current] = append(result[current], [2]int{start, start + count - 1})
	}

	return result, nil
}

func addUntrackedFiles(root string, result map[string][][2]int) (map[string][][2]int, error) {
	status := exec.Command("git", "status", "--porcelain", "--untracked-files=all")
	status.Dir = root

	statusOutput, err := status.Output()
	if err != nil {
		return nil, err
	}

	for line := range strings.SplitSeq(string(statusOutput), "\n") {
		if !strings.HasPrefix(line, "?? ") || !strings.HasSuffix(line, ".go") {
			continue
		}

		path := filepath.Join(root, strings.TrimPrefix(line, "?? "))

		data, readErr := os.ReadFile(path)
		if readErr != nil {
			return nil, readErr
		}

		result[path] = append(result[path], [2]int{1, strings.Count(string(data), "\n") + 1})
	}

	return result, nil
}
