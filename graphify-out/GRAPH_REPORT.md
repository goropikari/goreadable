# Graph Report - goreadable  (2026-08-14)

## Corpus Check
- 29 files · ~16,067 words
- Verdict: corpus is large enough that graph structure adds value.

## Summary
- 109 nodes · 176 edges · 10 communities (9 shown, 1 thin omitted)
- Extraction: 88% EXTRACTED · 12% INFERRED · 0% AMBIGUOUS · INFERRED: 21 edges (avg confidence: 0.8)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `e7e12cc6`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- analysis.go
- risks
- goreadable_test.go
- Graphify Incremental Update
- newCommand
- Testing Guidelines
- Thresholds
- Result
- github.com/goropikari/goreadable

## God Nodes (most connected - your core abstractions)
1. `AnalyzeWithOptions()` - 17 edges
2. `newCommand()` - 14 edges
3. `TestGoreadableCLI()` - 9 edges
4. `Thresholds` - 8 edges
5. `risks` - 6 edges
6. `Result` - 6 edges
7. `Testing Guidelines` - 6 edges
8. `execute()` - 5 edges
9. `Options` - 5 edges
10. `Analyze()` - 5 edges

## Surprising Connections (you probably didn't know these)
- `Testing Guidelines` --semantically_similar_to--> `Repository Testing Guidelines`  [INFERRED] [semantically similar]
  TESTING.md → docs/testing-guidelines.md
- `newCommand()` --calls--> `AnalyzeWithOptions()`  [INFERRED]
  cmd/goreadable/main.go → internal/analysis/analysis.go
- `newCommand()` --calls--> `NewOptions()`  [INFERRED]
  cmd/goreadable/main.go → internal/analysis/analysis.go
- `newCommand()` --calls--> `Defaults()`  [INFERRED]
  cmd/goreadable/main.go → internal/config/config.go
- `newCommand()` --calls--> `LoadFile()`  [INFERRED]
  cmd/goreadable/main.go → internal/config/config.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Graphify Documentation Workflow** — _codex_skills_graphify_skill_graphify_workflow, _codex_skills_graphify_references_extraction_spec_graphify_extraction_spec, _codex_skills_graphify_references_query_graphify_query, _codex_skills_graphify_references_update_graphify_update [EXTRACTED 1.00]
- **Repository Testing Contract** — agents_repository_guidance, testing_testing_guidelines, docs_testing_guidelines_testing_guidelines [EXTRACTED 1.00]

## Communities (10 total, 1 thin omitted)

### Community 0 - "analysis.go"
Cohesion: 0.14
Nodes (23): Options, BlockStmt, CommentGroup, Expr, File, FuncDecl, FuncType, Analyze() (+15 more)

### Community 1 - "risks"
Cohesion: 0.12
Nodes (14): artifact, build_command, consumer_command, kind, criteria, immutable_paths, risks, version (+6 more)

### Community 2 - "goreadable_test.go"
Cohesion: 0.41
Nodes (12): Command, buildGoreadable(), candidateByName(), candidateNames(), candidates(), T, repositoryRoot(), runCommand() (+4 more)

### Community 3 - "Graphify Incremental Update"
Cohesion: 0.17
Nodes (12): Graphify Add and Watch, Graphify Exports, Graphify Extraction Specification, Graphify GitHub and Merge, Graphify Hooks, Graphify Query, Graphify Transcription, Graphify Incremental Update (+4 more)

### Community 4 - "newCommand"
Cohesion: 0.23
Nodes (10): execute(), flagOverrides(), Writer, inputFiles(), main(), newCommand(), pathsRoot(), T (+2 more)

### Community 5 - "Testing Guidelines"
Cohesion: 0.28
Nodes (9): Pull Request Template, Continuous Integration Pipeline, Golangci Lint Configuration, Repository Agent Guidance, Behavior-level Testing, Continuous Integration, Repository Testing Guidelines, Goreadable CLI (+1 more)

### Community 6 - "Thresholds"
Cohesion: 0.31
Nodes (7): fileConfig, Thresholds, T, TestAnalyze(), Defaults(), LoadFile(), merge()

### Community 7 - "Result"
Cohesion: 0.52
Nodes (6): Writer, metricNames(), WriteJSON(), WriteText(), Candidate, Result

## Knowledge Gaps
- **22 isolated node(s):** `version`, `kind`, `build_command`, `consumer_command`, `external-artifact` (+17 more)
  These have ≤1 connection - possible missing edges or undocumented components.
- **1 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `newCommand()` connect `newCommand` to `analysis.go`, `goreadable_test.go`, `Thresholds`, `Result`?**
  _High betweenness centrality (0.225) - this node is a cross-community bridge._
- **Why does `AnalyzeWithOptions()` connect `analysis.go` to `newCommand`, `Thresholds`, `Result`?**
  _High betweenness centrality (0.161) - this node is a cross-community bridge._
- **Are the 7 inferred relationships involving `newCommand()` (e.g. with `AnalyzeWithOptions()` and `NewOptions()`) actually correct?**
  _`newCommand()` has 7 INFERRED edges - model-reasoned connections that need verification._
- **What connects `version`, `kind`, `build_command` to the rest of the system?**
  _22 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `analysis.go` be split into smaller, more focused modules?**
  _Cohesion score 0.14153846153846153 - nodes in this community are weakly interconnected._
- **Should `risks` be split into smaller, more focused modules?**
  _Cohesion score 0.125 - nodes in this community are weakly interconnected._