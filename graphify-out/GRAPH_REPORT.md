# Graph Report - . (2026-08-09)

## Corpus Check

- Corpus is ~15,382 words - fits in a single context window. You may not need a graph.

## Summary

- 94 nodes · 150 edges · 9 communities (8 shown, 1 thin omitted)
- Extraction: 87% EXTRACTED · 13% INFERRED · 0% AMBIGUOUS · INFERRED: 19 edges (avg confidence: 0.81)
- Token cost: 0 input · 0 output

## Community Hubs (Navigation)

- Analysis Engine
- Acceptance Contract
- Acceptance Tests
- Graphify Documentation
- CLI Command Flow
- Repository Standards
- Configuration
- Report Output
- Module Identity

## God Nodes (most connected - your core abstractions)

1. `Analyze()` - 14 edges
2. `newCommand()` - 13 edges
3. `TestGoreadableCLI()` - 9 edges
4. `Thresholds` - 7 edges
5. `risks` - 6 edges
6. `Testing Guidelines` - 6 edges
7. `TestAnalyze()` - 5 edges
8. `Result` - 5 edges
9. `TestGoreadableModuleIdentity()` - 5 edges
10. `buildGoreadable()` - 5 edges

## Surprising Connections (you probably didn't know these)

- `Testing Guidelines` --semantically_similar_to--> `Repository Testing Guidelines` [INFERRED] [semantically similar]
  TESTING.md → docs/testing-guidelines.md
- `newCommand()` --calls--> `Analyze()` [INFERRED]
  cmd/goreadable/main.go → internal/analysis/analysis.go
- `newCommand()` --calls--> `Defaults()` [INFERRED]
  cmd/goreadable/main.go → internal/config/config.go
- `newCommand()` --calls--> `LoadFile()` [INFERRED]
  cmd/goreadable/main.go → internal/config/config.go
- `newCommand()` --calls--> `WriteJSON()` [INFERRED]
  cmd/goreadable/main.go → internal/report/report.go

## Import Cycles

- None detected.

## Hyperedges (group relationships)

- **Graphify Documentation Workflow** — _codex_skills_graphify_skill_graphify_workflow, _codex_skills_graphify_references_extraction_spec_graphify_extraction_spec, _codex_skills_graphify_references_query_graphify_query, _codex_skills_graphify_references_update_graphify_update [EXTRACTED 1.00]
- **Repository Testing Contract** — agents_repository_guidance, testing_testing_guidelines, docs_testing_guidelines_testing_guidelines [EXTRACTED 1.00]

## Communities (9 total, 1 thin omitted)

### Community 0 - "Analysis Engine"

Cohesion: 0.17
Nodes (18): BlockStmt, CommentGroup, File, FuncType, Analyze(), argumentCount(), candidate(), complexity() (+10 more)

### Community 1 - "Acceptance Contract"

Cohesion: 0.12
Nodes (14): artifact, build_command, consumer_command, kind, criteria, immutable_paths, risks, version (+6 more)

### Community 2 - "Acceptance Tests"

Cohesion: 0.41
Nodes (12): Command, buildGoreadable(), candidateByName(), candidateNames(), candidates(), T, repositoryRoot(), runCommand() (+4 more)

### Community 3 - "Graphify Documentation"

Cohesion: 0.17
Nodes (12): Graphify Add and Watch, Graphify Exports, Graphify Extraction Specification, Graphify GitHub and Merge, Graphify Hooks, Graphify Query, Graphify Transcription, Graphify Incremental Update (+4 more)

### Community 4 - "CLI Command Flow"

Cohesion: 0.33
Nodes (8): execute(), flagOverrides(), Writer, inputFiles(), main(), newCommand(), pathsRoot(), ChangedFiles()

### Community 5 - "Repository Standards"

Cohesion: 0.28
Nodes (9): Pull Request Template, Continuous Integration Pipeline, Golangci Lint Configuration, Repository Agent Guidance, Behavior-level Testing, Continuous Integration, Repository Testing Guidelines, Goreadable CLI (+1 more)

### Community 6 - "Configuration"

Cohesion: 0.52
Nodes (5): fileConfig, Thresholds, Defaults(), LoadFile(), merge()

### Community 7 - "Report Output"

Cohesion: 0.60
Nodes (5): Writer, WriteJSON(), WriteText(), Candidate, Result

## Knowledge Gaps

- **22 isolated node(s):** `version`, `kind`, `build_command`, `consumer_command`, `external-artifact` (+17 more)
  These have ≤1 connection - possible missing edges or undocumented components.
- **1 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions

_Questions this graph is uniquely positioned to answer:_

- **Why does `newCommand()` connect `CLI Command Flow` to `Analysis Engine`, `Acceptance Tests`, `Configuration`, `Report Output`?**
  _High betweenness centrality (0.202) - this node is a cross-community bridge._
- **Why does `Analyze()` connect `Analysis Engine` to `CLI Command Flow`, `Configuration`, `Report Output`?**
  _High betweenness centrality (0.161) - this node is a cross-community bridge._
- **Why does `Thresholds` connect `Configuration` to `Analysis Engine`?**
  _High betweenness centrality (0.034) - this node is a cross-community bridge._
- **Are the 2 inferred relationships involving `Analyze()` (e.g. with `newCommand()` and `TestAnalyze()`) actually correct?**
  _`Analyze()` has 2 INFERRED edges - model-reasoned connections that need verification._
- **Are the 6 inferred relationships involving `newCommand()` (e.g. with `Analyze()` and `Defaults()`) actually correct?**
  _`newCommand()` has 6 INFERRED edges - model-reasoned connections that need verification._
- **Are the 5 inferred relationships involving `Command` (e.g. with `ChangedFiles()` and `buildGoreadable()`) actually correct?**
  _`Command` has 5 INFERRED edges - model-reasoned connections that need verification._
- **What connects `version`, `kind`, `build_command` to the rest of the system?**
  _22 weakly-connected nodes found - possible documentation gaps or missing edges._
