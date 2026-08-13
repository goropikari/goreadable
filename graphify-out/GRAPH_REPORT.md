# Graph Report - goreadable (2026-08-14)

## Corpus Check

- 33 files · ~20,108 words
- Verdict: corpus is large enough that graph structure adds value.

## Summary

- 141 nodes · 259 edges · 11 communities (9 shown, 2 thin omitted)
- Extraction: 92% EXTRACTED · 8% INFERRED · 0% AMBIGUOUS · INFERRED: 22 edges (avg confidence: 0.8)
- Token cost: 0 input · 0 output

## Graph Freshness

- Built from commit: `5229c253`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)

- AnalyzeWithOptions
- risks
- goreadable_test.go
- Testing Guidelines
- newCommand
- README.md
- Thresholds
- Result
- github.com/goropikari/goreadable
- analysis.go

## God Nodes (most connected - your core abstractions)

1. `functionMetrics()` - 19 edges
2. `AnalyzeWithOptions()` - 18 edges
3. `newCommand()` - 14 edges
4. `countNodes()` - 9 edges
5. `Thresholds` - 9 edges
6. `TestGoreadableCLI()` - 9 edges
7. `risks` - 6 edges
8. `localVariables()` - 6 edges
9. `Result` - 6 edges
10. `Testing Guidelines` - 6 edges

## Surprising Connections (you probably didn't know these)

- `Testing Guidelines` --semantically_similar_to--> `Repository Testing Guidelines` [INFERRED] [semantically similar]
  TESTING.md → docs/testing-guidelines.md
- `newCommand()` --calls--> `AnalyzeWithOptions()` [INFERRED]
  cmd/goreadable/main.go → internal/analysis/analysis.go
- `newCommand()` --calls--> `NewOptions()` [INFERRED]
  cmd/goreadable/main.go → internal/analysis/analysis.go
- `newCommand()` --calls--> `Defaults()` [INFERRED]
  cmd/goreadable/main.go → internal/config/config.go
- `newCommand()` --calls--> `LoadFile()` [INFERRED]
  cmd/goreadable/main.go → internal/config/config.go

## Import Cycles

- None detected.

## Hyperedges (group relationships)

- **Graphify Documentation Workflow** — _codex_skills_graphify_skill_graphify_workflow, _codex_skills_graphify_references_extraction_spec_graphify_extraction_spec, _codex_skills_graphify_references_query_graphify_query, _codex_skills_graphify_references_update_graphify_update [EXTRACTED 1.00]
- **Repository Testing Contract** — agents_repository_guidance, testing_testing_guidelines, docs_testing_guidelines_testing_guidelines [EXTRACTED 1.00]

## Communities (11 total, 2 thin omitted)

### Community 0 - "AnalyzeWithOptions"

Cohesion: 0.15
Nodes (17): Options, CommentGroup, File, FuncDecl, AnalyzeWithOptions(), exportedFields(), exportedMethodCounts(), functionSelector() (+9 more)

### Community 1 - "risks"

Cohesion: 0.12
Nodes (14): artifact, build_command, consumer_command, kind, criteria, immutable_paths, risks, version (+6 more)

### Community 2 - "goreadable_test.go"

Cohesion: 0.41
Nodes (12): Command, buildGoreadable(), candidateByName(), candidateNames(), candidates(), T, repositoryRoot(), runCommand() (+4 more)

### Community 3 - "Testing Guidelines"

Cohesion: 0.10
Nodes (21): Graphify Add and Watch, Graphify Exports, Graphify Extraction Specification, Graphify GitHub and Merge, Graphify Hooks, Graphify Query, Graphify Transcription, Graphify Incremental Update (+13 more)

### Community 4 - "newCommand"

Cohesion: 0.23
Nodes (10): execute(), flagOverrides(), Writer, inputFiles(), main(), newCommand(), pathsRoot(), T (+2 more)

### Community 6 - "Thresholds"

Cohesion: 0.17
Nodes (12): fileConfig, Thresholds, Analyze(), Files(), isGenerated(), T, TestAnalyze(), Defaults() (+4 more)

### Community 7 - "Result"

Cohesion: 0.43
Nodes (7): candidate(), Writer, metricNames(), WriteJSON(), WriteText(), Candidate, Result

### Community 10 - "analysis.go"

Cohesion: 0.17
Nodes (29): AssignStmt, BlockStmt, Expr, FileSet, FuncType, Ident, argumentCount(), booleanOperators() (+21 more)

## Knowledge Gaps

- **23 isolated node(s):** `version`, `kind`, `build_command`, `consumer_command`, `external-artifact` (+18 more)
  These have ≤1 connection - possible missing edges or undocumented components.
- **2 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions

_Questions this graph is uniquely positioned to answer:_

- **Why does `newCommand()` connect `newCommand` to `AnalyzeWithOptions`, `goreadable_test.go`, `Thresholds`, `Result`?**
  _High betweenness centrality (0.222) - this node is a cross-community bridge._
- **Why does `AnalyzeWithOptions()` connect `AnalyzeWithOptions` to `analysis.go`, `newCommand`, `Thresholds`, `Result`?**
  _High betweenness centrality (0.168) - this node is a cross-community bridge._
- **Why does `functionMetrics()` connect `analysis.go` to `AnalyzeWithOptions`?**
  _High betweenness centrality (0.044) - this node is a cross-community bridge._
- **Are the 7 inferred relationships involving `newCommand()` (e.g. with `AnalyzeWithOptions()` and `NewOptions()`) actually correct?**
  _`newCommand()` has 7 INFERRED edges - model-reasoned connections that need verification._
- **What connects `version`, `kind`, `build_command` to the rest of the system?**
  _23 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `risks` be split into smaller, more focused modules?**
  _Cohesion score 0.125 - nodes in this community are weakly interconnected._
- **Should `Testing Guidelines` be split into smaller, more focused modules?**
  _Cohesion score 0.10476190476190476 - nodes in this community are weakly interconnected._
