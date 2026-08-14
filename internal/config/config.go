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
	LocalVariables       int `json:"local_variables"`
	ControlBlocks        int `json:"control_blocks"`
	ReturnPoints         int `json:"return_points"`
	BooleanOperators     int `json:"boolean_operators"`
	MaxConditionTerms    int `json:"max_condition_terms"`
	FunctionCalls        int `json:"function_calls"`
	LiteralValues        int `json:"literal_values"`
	ClosureCount         int `json:"closure_count"`
	CommentLines         int `json:"comment_lines"`
	StatementCount       int `json:"statement_count"`
	TypeDependencies     int `json:"type_dependencies"`
	StructFields         int `json:"struct_fields"`
	TypeMethods          int `json:"type_methods"`
	ExportedMembers      int `json:"exported_members"`
}

func Defaults() Thresholds {
	return Thresholds{
		FunctionLines:        80,
		NestingDepth:         4,
		CyclomaticComplexity: 10,
		FunctionArguments:    5,
		LocalVariables:       15,
		ControlBlocks:        8,
		ReturnPoints:         5,
		BooleanOperators:     8,
		MaxConditionTerms:    4,
		FunctionCalls:        15,
		LiteralValues:        10,
		ClosureCount:         2,
		CommentLines:         10,
		StatementCount:       40,
		TypeDependencies:     5,
		StructFields:         8,
		TypeMethods:          10,
		ExportedMembers:      10,
	}
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
	applyFunctionFlags(t, values)
	applyTypeFlags(t, values)
}

func applyFunctionFlags(t *Thresholds, values map[string]int) {
	applyBasicFunctionFlags(t, values)
	applyDetailedFunctionFlags(t, values)
}

func applyBasicFunctionFlags(t *Thresholds, values map[string]int) {
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

	if value, ok := values["local_variables"]; ok {
		t.LocalVariables = value
	}

	if value, ok := values["control_blocks"]; ok {
		t.ControlBlocks = value
	}

	if value, ok := values["return_points"]; ok {
		t.ReturnPoints = value
	}

	if value, ok := values["boolean_operators"]; ok {
		t.BooleanOperators = value
	}
}

func applyDetailedFunctionFlags(t *Thresholds, values map[string]int) {
	if value, ok := values["max_condition_terms"]; ok {
		t.MaxConditionTerms = value
	}

	if value, ok := values["function_calls"]; ok {
		t.FunctionCalls = value
	}

	if value, ok := values["literal_values"]; ok {
		t.LiteralValues = value
	}

	if value, ok := values["closure_count"]; ok {
		t.ClosureCount = value
	}

	if value, ok := values["comment_lines"]; ok {
		t.CommentLines = value
	}

	if value, ok := values["statement_count"]; ok {
		t.StatementCount = value
	}
}

func applyTypeFlags(t *Thresholds, values map[string]int) {
	if value, ok := values["type_dependencies"]; ok {
		t.TypeDependencies = value
	}

	if value, ok := values["struct_fields"]; ok {
		t.StructFields = value
	}

	if value, ok := values["type_methods"]; ok {
		t.TypeMethods = value
	}

	if value, ok := values["exported_members"]; ok {
		t.ExportedMembers = value
	}
}

func merge(base *Thresholds, override Thresholds) {
	mergeFunctionThresholds(base, override)
	mergeTypeThresholds(base, override)
}

func mergeFunctionThresholds(base *Thresholds, override Thresholds) {
	mergeBasicFunctionThresholds(base, override)
	mergeDetailedFunctionThresholds(base, override)
}

func mergeBasicFunctionThresholds(base *Thresholds, override Thresholds) {
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

	if override.LocalVariables != 0 {
		base.LocalVariables = override.LocalVariables
	}

	if override.ControlBlocks != 0 {
		base.ControlBlocks = override.ControlBlocks
	}

	if override.ReturnPoints != 0 {
		base.ReturnPoints = override.ReturnPoints
	}

	if override.BooleanOperators != 0 {
		base.BooleanOperators = override.BooleanOperators
	}
}

func mergeDetailedFunctionThresholds(base *Thresholds, override Thresholds) {
	if override.MaxConditionTerms != 0 {
		base.MaxConditionTerms = override.MaxConditionTerms
	}

	if override.FunctionCalls != 0 {
		base.FunctionCalls = override.FunctionCalls
	}

	if override.LiteralValues != 0 {
		base.LiteralValues = override.LiteralValues
	}

	if override.ClosureCount != 0 {
		base.ClosureCount = override.ClosureCount
	}

	if override.CommentLines != 0 {
		base.CommentLines = override.CommentLines
	}

	if override.StatementCount != 0 {
		base.StatementCount = override.StatementCount
	}
}

func mergeTypeThresholds(base *Thresholds, override Thresholds) {
	if override.TypeDependencies != 0 {
		base.TypeDependencies = override.TypeDependencies
	}

	if override.StructFields != 0 {
		base.StructFields = override.StructFields
	}

	if override.TypeMethods != 0 {
		base.TypeMethods = override.TypeMethods
	}

	if override.ExportedMembers != 0 {
		base.ExportedMembers = override.ExportedMembers
	}
}
