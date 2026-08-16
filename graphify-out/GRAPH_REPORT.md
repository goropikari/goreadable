# Graph Report - goreadable (2026-08-16)

## Corpus Check

- 35 files · ~21,333 words
- Verdict: corpus is large enough that graph structure adds value.

## Summary

- 306 nodes · 533 edges · 16 communities (13 shown, 3 thin omitted)
- Extraction: 95% EXTRACTED · 5% INFERRED · 0% AMBIGUOUS · INFERRED: 25 edges (avg confidence: 0.8)
- Token cost: 0 input · 0 output

## Graph Freshness

- Built from commit: `17516c14`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)

- properties
- risks
- Command
- Testing Guidelines
- goreadable/main.go
- README.md
- Thresholds
- $defs
- github.com/goropikari/goreadable
- analysis.go
- Result
- properties
- required
- goreadable-output.schema.json
- TestJSONSchema

## God Nodes (most connected - your core abstractions)

1. `Thresholds` - 23 edges
2. `functionMetrics()` - 19 edges
3. `required` - 17 edges
4. `required` - 17 edges
5. `collectFunctionCandidate()` - 16 edges
6. `collectTypeCandidate()` - 13 edges
7. `runCommand()` - 10 edges
8. `Options` - 10 edges
9. `AnalyzeWithOptions()` - 9 edges
10. `countNodes()` - 9 edges

## Surprising Connections (you probably didn't know these)

- `Testing Guidelines` --semantically_similar_to--> `Repository Testing Guidelines` [INFERRED] [semantically similar]
  TESTING.md → docs/testing-guidelines.md
- `newCommand()` --calls--> `Defaults()` [INFERRED]
  cmd/goreadable/main.go → internal/config/config.go
- `configuredThresholds()` --calls--> `LoadFile()` [INFERRED]
  cmd/goreadable/main.go → internal/config/config.go
- `writeResult()` --calls--> `WriteJSON()` [INFERRED]
  cmd/goreadable/main.go → internal/report/report.go
- `writeResult()` --calls--> `WriteText()` [INFERRED]
  cmd/goreadable/main.go → internal/report/report.go

## Import Cycles

- None detected.

## Hyperedges (group relationships)

- **Graphify Documentation Workflow** — _codex_skills_graphify_skill_graphify_workflow, _codex_skills_graphify_references_extraction_spec_graphify_extraction_spec, _codex_skills_graphify_references_query_graphify_query, _codex_skills_graphify_references_update_graphify_update [EXTRACTED 1.00]
- **Repository Testing Contract** — agents_repository_guidance, testing_testing_guidelines, docs_testing_guidelines_testing_guidelines [EXTRACTED 1.00]

## Communities (16 total, 3 thin omitted)

### Community 0 - "properties"

Cohesion: 0.09
Nodes (32): $ref, $ref, $ref, $ref, $ref, $ref, $ref, $ref (+24 more)

### Community 1 - "risks"

Cohesion: 0.12
Nodes (14): artifact, build_command, consumer_command, kind, criteria, immutable_paths, risks, version (+6 more)

### Community 2 - "Command"

Cohesion: 0.29
Nodes (15): Command, addUntrackedFiles(), ChangedFiles(), changedTrackedFiles(), buildGoreadable(), candidateByName(), candidateNames(), candidates() (+7 more)

### Community 3 - "Testing Guidelines"

Cohesion: 0.10
Nodes (21): Graphify Add and Watch, Graphify Exports, Graphify Extraction Specification, Graphify GitHub and Merge, Graphify Hooks, Graphify Query, Graphify Transcription, Graphify Incremental Update (+13 more)

### Community 4 - "goreadable/main.go"

Cohesion: 0.19
Nodes (17): Options, changedFiles(), commandOptions(), configuredThresholds(), execute(), flagOverrides(), Writer, inputFiles() (+9 more)

### Community 6 - "Thresholds"

Cohesion: 0.19
Nodes (19): fileConfig, Thresholds, Analyze(), T, TestAnalyze(), applyBasicFunctionFlags(), applyDetailedFunctionFlags(), applyFunctionFlags() (+11 more)

### Community 7 - "$defs"

Cohesion: 0.07
Nodes (31): $ref, oneOf, unevaluatedProperties, $defs, candidate, functionMetrics, functionThresholds, metricMap (+23 more)

### Community 10 - "analysis.go"

Cohesion: 0.08
Nodes (58): AssignStmt, BlockStmt, CommentGroup, DeclStmt, Expr, File, FileSet, FuncDecl (+50 more)

### Community 11 - "Result"

Cohesion: 0.50
Nodes (7): Writer, metricNames(), writeEmptyTextResult(), WriteJSON(), WriteText(), Candidate, Result

### Community 12 - "properties"

Cohesion: 0.07
Nodes (30): code_kind, end_line, name, path, production, reasons, start_line, test (+22 more)

### Community 13 - "required"

Cohesion: 0.16
Nodes (22): boolean_operators, closure_count, comment_lines, control_blocks, cyclomatic_complexity, exported_members, function_arguments, function_calls (+14 more)

### Community 14 - "goreadable-output.schema.json"

Cohesion: 0.11
Nodes (17): candidates, version, additionalProperties, items, type, description, $id, $ref (+9 more)

## Knowledge Gaps

- **86 isolated node(s):** `version`, `kind`, `build_command`, `consumer_command`, `external-artifact` (+81 more)
  These have ≤1 connection - possible missing edges or undocumented components.
- **3 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions

_Questions this graph is uniquely positioned to answer:_

- **Why does `$defs` connect `$defs` to `properties`, `goreadable-output.schema.json`?**
  _High betweenness centrality (0.122) - this node is a cross-community bridge._
- **Why does `candidateCommon` connect `properties` to `$defs`?**
  _High betweenness centrality (0.068) - this node is a cross-community bridge._
- **Why does `Thresholds` connect `Thresholds` to `analysis.go`, `goreadable/main.go`?**
  _High betweenness centrality (0.063) - this node is a cross-community bridge._
- **What connects `version`, `kind`, `build_command` to the rest of the system?**
  _86 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `properties` be split into smaller, more focused modules?**
  _Cohesion score 0.0907258064516129 - nodes in this community are weakly interconnected._
- **Should `risks` be split into smaller, more focused modules?**
  _Cohesion score 0.125 - nodes in this community are weakly interconnected._
- **Should `Testing Guidelines` be split into smaller, more focused modules?**
  _Cohesion score 0.10476190476190476 - nodes in this community are weakly interconnected._
