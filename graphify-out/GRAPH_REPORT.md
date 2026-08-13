# Graph Report - goreadable (2026-08-14)

## Corpus Check

- 31 files · ~17,641 words
- Verdict: corpus is large enough that graph structure adds value.

## Summary

- 133 nodes · 199 edges · 10 communities (9 shown, 1 thin omitted)
- Extraction: 89% EXTRACTED · 11% INFERRED · 0% AMBIGUOUS · INFERRED: 21 edges (avg confidence: 0.8)
- Token cost: 0 input · 0 output

## Graph Freshness

- Built from commit: `998919b4`
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
- blackbox-risk-based-test.md
- github.com/goropikari/goreadable
- Implementation Plan

## God Nodes (most connected - your core abstractions)

1. `AnalyzeWithOptions()` - 17 edges
2. `newCommand()` - 14 edges
3. `TestGoreadableCLI()` - 9 edges
4. `Implementation Plan` - 9 edges
5. `Thresholds` - 8 edges
6. `risks` - 6 edges
7. `Result` - 6 edges
8. `Prioritized Test Conditions` - 6 edges
9. `Testing Guidelines` - 6 edges
10. `execute()` - 5 edges

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

## Communities (10 total, 1 thin omitted)

### Community 0 - "analysis.go"

Cohesion: 0.15
Nodes (22): Options, BlockStmt, CommentGroup, Expr, File, FuncDecl, FuncType, AnalyzeWithOptions() (+14 more)

### Community 1 - "risks"

Cohesion: 0.12
Nodes (14): artifact, build_command, consumer_command, kind, criteria, immutable_paths, risks, version (+6 more)

### Community 2 - "goreadable_test.go"

Cohesion: 0.32
Nodes (13): Command, ChangedFiles(), buildGoreadable(), candidateByName(), candidateNames(), candidates(), T, repositoryRoot() (+5 more)

### Community 3 - "Graphify Incremental Update"

Cohesion: 0.17
Nodes (12): Graphify Add and Watch, Graphify Exports, Graphify Extraction Specification, Graphify GitHub and Merge, Graphify Hooks, Graphify Query, Graphify Transcription, Graphify Incremental Update (+4 more)

### Community 4 - "newCommand"

Cohesion: 0.19
Nodes (15): execute(), flagOverrides(), Writer, inputFiles(), main(), newCommand(), pathsRoot(), T (+7 more)

### Community 5 - "Testing Guidelines"

Cohesion: 0.28
Nodes (9): Pull Request Template, Continuous Integration Pipeline, Golangci Lint Configuration, Repository Agent Guidance, Behavior-level Testing, Continuous Integration, Repository Testing Guidelines, Goreadable CLI (+1 more)

### Community 6 - "Thresholds"

Cohesion: 0.29
Nodes (8): fileConfig, Thresholds, Analyze(), T, TestAnalyze(), Defaults(), LoadFile(), merge()

### Community 7 - "blackbox-risk-based-test.md"

Cohesion: 0.12
Nodes (15): Acceptance Criteria Improvements, Not Run / Deferred Checks, Prioritized Test Conditions, Regression and Exploratory Coverage, Residual Risks and Assumptions, RISK-001: A selector still filters out under-threshold functions, RISK-002: Text output omits the requested metric values, RISK-003: New flags alter the default review-candidate workflow (+7 more)

### Community 9 - "Implementation Plan"

Cohesion: 0.20
Nodes (9): Acceptance and validation, Change size and staging, Change summary, Completion evidence, Decision gates, Implementation Plan, Implementation steps, Risk and decision analysis (+1 more)

## Knowledge Gaps

- **43 isolated node(s):** `version`, `kind`, `build_command`, `consumer_command`, `external-artifact` (+38 more)
  These have ≤1 connection - possible missing edges or undocumented components.
- **1 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions

_Questions this graph is uniquely positioned to answer:_

- **Why does `newCommand()` connect `newCommand` to `analysis.go`, `goreadable_test.go`, `Thresholds`?**
  _High betweenness centrality (0.150) - this node is a cross-community bridge._
- **Why does `AnalyzeWithOptions()` connect `analysis.go` to `newCommand`, `Thresholds`?**
  _High betweenness centrality (0.108) - this node is a cross-community bridge._
- **Are the 7 inferred relationships involving `newCommand()` (e.g. with `AnalyzeWithOptions()` and `NewOptions()`) actually correct?**
  _`newCommand()` has 7 INFERRED edges - model-reasoned connections that need verification._
- **What connects `version`, `kind`, `build_command` to the rest of the system?**
  _43 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `analysis.go` be split into smaller, more focused modules?**
  _Cohesion score 0.14666666666666667 - nodes in this community are weakly interconnected._
- **Should `risks` be split into smaller, more focused modules?**
  _Cohesion score 0.125 - nodes in this community are weakly interconnected._
- **Should `blackbox-risk-based-test.md` be split into smaller, more focused modules?**
  _Cohesion score 0.125 - nodes in this community are weakly interconnected._
