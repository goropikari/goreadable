package report

import (
	"encoding/json"
	"fmt"
	"io"
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
	Source     string         `json:"source"`
}

type Result struct {
	Version    int         `json:"version"`
	Candidates []Candidate `json:"candidates"`
}

func WriteJSON(w io.Writer, result Result) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")

	return encoder.Encode(result)
}

func WriteText(w io.Writer, result Result) error {
	if len(result.Candidates) == 0 {
		_, err := fmt.Fprintln(w, "No readability review candidates found.")
		return err
	}

	for _, candidate := range result.Candidates {
		if _, err := fmt.Fprintf(w, "%s %s (%s:%d-%d, %s)\n", candidate.Kind, candidate.Name, candidate.Path, candidate.StartLine, candidate.EndLine, candidate.CodeKind); err != nil {
			return err
		}

		for _, reason := range candidate.Reasons {
			if _, err := fmt.Fprintf(w, "  - %s\n", reason); err != nil {
				return err
			}
		}
	}

	return nil
}
