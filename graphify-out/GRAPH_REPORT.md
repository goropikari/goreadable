# Graph Report - goreadable  (2026-08-16)

## Corpus Check
- 34 files · ~21,563 words
- Verdict: corpus is large enough that graph structure adds value.

## Summary
- 307 nodes · 536 edges · 17 communities (14 shown, 3 thin omitted)
- Extraction: 96% EXTRACTED · 4% INFERRED · 0% AMBIGUOUS · INFERRED: 24 edges (avg confidence: 0.8)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `1dbedf85`
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
- collectFunctionCandidate
- properties
- required
- goreadable-output.schema.json
- TestJSONSchema
- typeMetrics

## God Nodes (most connected - your core abstractions)
1. `Thresholds` - 23 edges
2. `functionMetrics()` - 19 edges
3. `collectFunctionCandidate()` - 17 edges
4. `required` - 17 edges
5. `required` - 17 edges
6. `collectTypeCandidate()` - 14 edges
7. `Options` - 12 edges
8. `runCommand()` - 10 edges
9. `AnalyzeWithOptions()` - 9 edges
10. `collectCandidates()` - 9 edges

## Surprising Connections (you probably didn't know these)
- `newCommand()` --calls--> `Defaults()`  [INFERRED]
  cmd/goreadable/main.go → internal/config/config.go
- `configuredThresholds()` --calls--> `LoadFile()`  [INFERRED]
  cmd/goreadable/main.go → internal/config/config.go
- `writeResult()` --calls--> `AnalyzeWithOptions()`  [INFERRED]
  cmd/goreadable/main.go → internal/analysis/analysis.go
- `writeResult()` --calls--> `WriteJSON()`  [INFERRED]
  cmd/goreadable/main.go → internal/report/report.go
- `writeResult()` --calls--> `WriteText()`  [INFERRED]
  cmd/goreadable/main.go → internal/report/report.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Graphify Documentation Workflow** — _codex_skills_graphify_skill_graphify_workflow, _codex_skills_graphify_references_extraction_spec_graphify_extraction_spec, _codex_skills_graphify_references_query_graphify_query, _codex_skills_graphify_references_update_graphify_update [EXTRACTED 1.00]

## Communities (17 total, 3 thin omitted)

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
Cohesion: 0.11
Nodes (20): Graphify Add and Watch, Graphify Exports, Graphify Extraction Specification, Graphify GitHub and Merge, Graphify Hooks, Graphify Query, Graphify Transcription, Graphify Incremental Update (+12 more)

### Community 4 - "goreadable/main.go"
Cohesion: 0.27
Nodes (12): changedFiles(), configuredThresholds(), execute(), flagOverrides(), Writer, main(), newCommand(), pathsRoot() (+4 more)

### Community 6 - "Thresholds"
Cohesion: 0.19
Nodes (19): fileConfig, Thresholds, Analyze(), T, TestAnalyze(), applyBasicFunctionFlags(), applyDetailedFunctionFlags(), applyFunctionFlags() (+11 more)

### Community 7 - "$defs"
Cohesion: 0.08
Nodes (26): code_kind, end_line, name, path, reasons, start_line, $ref, oneOf (+18 more)

### Community 10 - "analysis.go"
Cohesion: 0.10
Nodes (46): AssignStmt, BlockStmt, inputFiles(), DeclStmt, Expr, File, FileSet, FuncType (+38 more)

### Community 11 - "collectFunctionCandidate"
Cohesion: 0.14
Nodes (24): Options, commandOptions(), CommentGroup, FuncDecl, AnalyzeWithOptions(), candidate(), collectCandidates(), collectFunctionCandidate() (+16 more)

### Community 12 - "properties"
Cohesion: 0.10
Nodes (21): production, test, properties, enum, minimum, type, type, minLength (+13 more)

### Community 13 - "required"
Cohesion: 0.22
Nodes (17): boolean_operators, closure_count, comment_lines, control_blocks, cyclomatic_complexity, function_arguments, function_calls, function_lines (+9 more)

### Community 14 - "goreadable-output.schema.json"
Cohesion: 0.11
Nodes (17): candidates, version, additionalProperties, items, type, description, $id, $ref (+9 more)

### Community 16 - "typeMetrics"
Cohesion: 0.13
Nodes (19): exported_members, struct_fields, type_methods, typeMetrics, typeThresholds, $ref, exported_members, struct_fields (+11 more)

## Knowledge Gaps
- **87 isolated node(s):** `version`, `kind`, `build_command`, `consumer_command`, `external-artifact` (+82 more)
  These have ≤1 connection - possible missing edges or undocumented components.
- **3 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `$defs` connect `$defs` to `typeMetrics`, `goreadable-output.schema.json`?**
  _High betweenness centrality (0.121) - this node is a cross-community bridge._
- **Why does `candidateCommon` connect `$defs` to `properties`?**
  _High betweenness centrality (0.068) - this node is a cross-community bridge._
- **Why does `Thresholds` connect `Thresholds` to `collectFunctionCandidate`, `goreadable/main.go`?**
  _High betweenness centrality (0.064) - this node is a cross-community bridge._
- **What connects `version`, `kind`, `build_command` to the rest of the system?**
  _87 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `properties` be split into smaller, more focused modules?**
  _Cohesion score 0.0907258064516129 - nodes in this community are weakly interconnected._
- **Should `risks` be split into smaller, more focused modules?**
  _Cohesion score 0.125 - nodes in this community are weakly interconnected._
- **Should `Testing Guidelines` be split into smaller, more focused modules?**
  _Cohesion score 0.10526315789473684 - nodes in this community are weakly interconnected._