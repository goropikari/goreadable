# Graph Report - goreadable (2026-08-16)

## Corpus Check

- 33 files · ~20,757 words
- Verdict: corpus is large enough that graph structure adds value.

## Summary

- 170 nodes · 359 edges · 12 communities (10 shown, 2 thin omitted)
- Extraction: 93% EXTRACTED · 7% INFERRED · 0% AMBIGUOUS · INFERRED: 25 edges (avg confidence: 0.8)
- Token cost: 0 input · 0 output

## Graph Freshness

- Built from commit: `ee5c1ba0`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)

- collectFunctionCandidate
- risks
- Command
- Testing Guidelines
- goreadable/main.go
- README.md
- Thresholds
- Analyze
- github.com/goropikari/goreadable
- analysis.go
- Result

## God Nodes (most connected - your core abstractions)

1. `Thresholds` - 23 edges
2. `functionMetrics()` - 19 edges
3. `collectFunctionCandidate()` - 16 edges
4. `collectTypeCandidate()` - 13 edges
5. `runCommand()` - 10 edges
6. `Options` - 10 edges
7. `AnalyzeWithOptions()` - 9 edges
8. `countNodes()` - 9 edges
9. `Result` - 9 edges
10. `TestGoreadableCLI()` - 9 edges

## Surprising Connections (you probably didn't know these)

- `Testing Guidelines` --semantically_similar_to--> `Repository Testing Guidelines` [INFERRED] [semantically similar]
  TESTING.md → docs/testing-guidelines.md
- `newCommand()` --calls--> `Defaults()` [INFERRED]
  cmd/goreadable/main.go → internal/config/config.go
- `configuredThresholds()` --calls--> `LoadFile()` [INFERRED]
  cmd/goreadable/main.go → internal/config/config.go
- `writeResult()` --calls--> `AnalyzeWithOptions()` [INFERRED]
  cmd/goreadable/main.go → internal/analysis/analysis.go
- `writeResult()` --calls--> `WriteJSON()` [INFERRED]
  cmd/goreadable/main.go → internal/report/report.go

## Import Cycles

- None detected.

## Hyperedges (group relationships)

- **Graphify Documentation Workflow** — _codex_skills_graphify_skill_graphify_workflow, _codex_skills_graphify_references_extraction_spec_graphify_extraction_spec, _codex_skills_graphify_references_query_graphify_query, _codex_skills_graphify_references_update_graphify_update [EXTRACTED 1.00]
- **Repository Testing Contract** — agents_repository_guidance, testing_testing_guidelines, docs_testing_guidelines_testing_guidelines [EXTRACTED 1.00]

## Communities (12 total, 2 thin omitted)

### Community 0 - "collectFunctionCandidate"

Cohesion: 0.13
Nodes (27): Options, commandOptions(), CommentGroup, File, FileSet, FuncDecl, AnalyzeWithOptions(), candidate() (+19 more)

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

Cohesion: 0.26
Nodes (13): changedFiles(), configuredThresholds(), execute(), flagOverrides(), Writer, inputFiles(), main(), newCommand() (+5 more)

### Community 6 - "Thresholds"

Cohesion: 0.25
Nodes (16): fileConfig, Thresholds, applyBasicFunctionFlags(), applyDetailedFunctionFlags(), applyFunctionFlags(), applyTypeFlags(), Defaults(), LoadFile() (+8 more)

### Community 7 - "Analyze"

Cohesion: 0.25
Nodes (8): Analyze(), directoryFiles(), Files(), isGenerated(), isSourceFile(), recursiveFiles(), T, TestAnalyze()

### Community 10 - "analysis.go"

Cohesion: 0.16
Nodes (30): AssignStmt, BlockStmt, DeclStmt, Expr, FuncType, Ident, argumentCount(), booleanOperators() (+22 more)

### Community 11 - "Result"

Cohesion: 0.50
Nodes (7): Writer, metricNames(), writeEmptyTextResult(), WriteJSON(), WriteText(), Candidate, Result

## Knowledge Gaps

- **23 isolated node(s):** `version`, `kind`, `build_command`, `consumer_command`, `external-artifact` (+18 more)
  These have ≤1 connection - possible missing edges or undocumented components.
- **2 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions

_Questions this graph is uniquely positioned to answer:_

- **Why does `Thresholds` connect `Thresholds` to `collectFunctionCandidate`, `goreadable/main.go`, `Analyze`?**
  _High betweenness centrality (0.206) - this node is a cross-community bridge._
- **Why does `runCommand()` connect `goreadable/main.go` to `collectFunctionCandidate`, `Command`, `Thresholds`?**
  _High betweenness centrality (0.089) - this node is a cross-community bridge._
- **Why does `collectFunctionCandidate()` connect `collectFunctionCandidate` to `analysis.go`, `Result`, `Thresholds`?**
  _High betweenness centrality (0.063) - this node is a cross-community bridge._
- **What connects `version`, `kind`, `build_command` to the rest of the system?**
  _23 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `collectFunctionCandidate` be split into smaller, more focused modules?**
  _Cohesion score 0.12561576354679804 - nodes in this community are weakly interconnected._
- **Should `risks` be split into smaller, more focused modules?**
  _Cohesion score 0.125 - nodes in this community are weakly interconnected._
- **Should `Testing Guidelines` be split into smaller, more focused modules?**
  _Cohesion score 0.10476190476190476 - nodes in this community are weakly interconnected._
