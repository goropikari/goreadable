package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Thresholds struct {
	FunctionLines        int `json:"function_lines"`
	NestingDepth         int `json:"nesting_depth"`
	CyclomaticComplexity int `json:"cyclomatic_complexity"`
	FunctionArguments    int `json:"function_arguments"`
	StructFields         int `json:"struct_fields"`
	TypeMethods          int `json:"type_methods"`
}

func Defaults() Thresholds {
	return Thresholds{FunctionLines: 80, NestingDepth: 4, CyclomaticComplexity: 10, FunctionArguments: 5, StructFields: 8, TypeMethods: 10}
}

type fileConfig struct {
	Thresholds Thresholds `json:"thresholds"`
}

func LoadFile(path string, defaults Thresholds) (Thresholds, error) {
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return defaults, nil
	}

	if err != nil {
		return Thresholds{}, err
	}

	var parsed fileConfig
	if err := json.Unmarshal(data, &parsed); err != nil {
		return Thresholds{}, fmt.Errorf("parse %s: %w", path, err)
	}

	merge(&defaults, parsed.Thresholds)

	return defaults, nil
}

func (t *Thresholds) ApplyFlags(values map[string]int) {
	if value, ok := values["function_lines"]; ok {
		t.FunctionLines = value
	}

	if value, ok := values["nesting_depth"]; ok {
		t.NestingDepth = value
	}

	if value, ok := values["cyclomatic_complexity"]; ok {
		t.CyclomaticComplexity = value
	}

	if value, ok := values["function_arguments"]; ok {
		t.FunctionArguments = value
	}

	if value, ok := values["struct_fields"]; ok {
		t.StructFields = value
	}

	if value, ok := values["type_methods"]; ok {
		t.TypeMethods = value
	}
}

func merge(base *Thresholds, override Thresholds) {
	if override.FunctionLines != 0 {
		base.FunctionLines = override.FunctionLines
	}

	if override.NestingDepth != 0 {
		base.NestingDepth = override.NestingDepth
	}

	if override.CyclomaticComplexity != 0 {
		base.CyclomaticComplexity = override.CyclomaticComplexity
	}

	if override.FunctionArguments != 0 {
		base.FunctionArguments = override.FunctionArguments
	}

	if override.StructFields != 0 {
		base.StructFields = override.StructFields
	}

	if override.TypeMethods != 0 {
		base.TypeMethods = override.TypeMethods
	}
}
