package report

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"
)

type Candidate struct {
	Kind       string         `json:"kind"`
	Name       string         `json:"name"`
	Path       string         `json:"path"`
	StartLine  int            `json:"start_line"`
	EndLine    int            `json:"end_line"`
	CodeKind   string         `json:"code_kind"`
	Metrics    map[string]int `json:"metrics"`
	Thresholds map[string]int `json:"thresholds"`
	Reasons    []string       `json:"reasons"`
}

type Result struct {
	Version     int         `json:"version"`
	Candidates  []Candidate `json:"candidates"`
	MetricsOnly bool        `json:"-"`
}

func WriteJSON(w io.Writer, result Result) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")

	return encoder.Encode(result)
}

func WriteText(w io.Writer, result Result) error {
	if len(result.Candidates) == 0 {
		return writeEmptyTextResult(w, result.MetricsOnly)
	}

	for _, candidate := range result.Candidates {
		if _, err := fmt.Fprintf(w, "%s %s (%s:%d-%d, %s)\n", candidate.Kind, candidate.Name, candidate.Path, candidate.StartLine, candidate.EndLine, candidate.CodeKind); err != nil {
			return err
		}

		if result.MetricsOnly {
			for _, name := range metricNames(candidate.Metrics) {
				if _, err := fmt.Fprintf(w, "  - %s=%d\n", name, candidate.Metrics[name]); err != nil {
					return err
				}
			}
		}

		for _, reason := range candidate.Reasons {
			if _, err := fmt.Fprintf(w, "  - %s\n", reason); err != nil {
				return err
			}
		}
	}

	return nil
}

func writeEmptyTextResult(w io.Writer, metricsOnly bool) error {
	message := "No readability review candidates found."
	if metricsOnly {
		message = "No functions found."
	}

	_, err := fmt.Fprintln(w, message)

	return err
}

func metricNames(metrics map[string]int) []string {
	names := make([]string, 0, len(metrics))
	for name := range metrics {
		names = append(names, name)
	}

	sort.Strings(names)

	return names
}
