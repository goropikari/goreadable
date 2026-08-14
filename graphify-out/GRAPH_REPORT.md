# Graph Report - goreadable (2026-08-14)

## Corpus Check

- 33 files · ~20,553 words
- Verdict: corpus is large enough that graph structure adds value.

## Summary

- 169 nodes · 355 edges · 11 communities (9 shown, 2 thin omitted)
- Extraction: 94% EXTRACTED · 6% INFERRED · 0% AMBIGUOUS · INFERRED: 23 edges (avg confidence: 0.8)
- Token cost: 0 input · 0 output

## Graph Freshness

- Built from commit: `641c64c9`
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
- Files
- github.com/goropikari/goreadable
- analysis.go

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
- `commandOptions()` --calls--> `NewOptions()` [INFERRED]
  cmd/goreadable/main.go → internal/analysis/analysis.go
- `configuredThresholds()` --calls--> `LoadFile()` [INFERRED]
  cmd/goreadable/main.go → internal/config/config.go
- `writeResult()` --calls--> `AnalyzeWithOptions()` [INFERRED]
  cmd/goreadable/main.go → internal/analysis/analysis.go

## Import Cycles

- None detected.

## Hyperedges (group relationships)

- **Graphify Documentation Workflow** — _codex_skills_graphify_skill_graphify_workflow, _codex_skills_graphify_references_extraction_spec_graphify_extraction_spec, _codex_skills_graphify_references_query_graphify_query, _codex_skills_graphify_references_update_graphify_update [EXTRACTED 1.00]
- **Repository Testing Contract** — agents_repository_guidance, testing_testing_guidelines, docs_testing_guidelines_testing_guidelines [EXTRACTED 1.00]

## Communities (11 total, 2 thin omitted)

### Community 0 - "collectFunctionCandidate"

Cohesion: 0.15
Nodes (24): Options, CommentGroup, FuncDecl, Analyze(), AnalyzeWithOptions(), candidate(), collectCandidates(), collectFunctionCandidate() (+16 more)

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

Cohesion: 0.24
Nodes (14): changedFiles(), commandOptions(), configuredThresholds(), execute(), flagOverrides(), Writer, inputFiles(), main() (+6 more)

### Community 6 - "Thresholds"

Cohesion: 0.25
Nodes (15): fileConfig, Thresholds, applyBasicFunctionFlags(), applyDetailedFunctionFlags(), applyFunctionFlags(), applyTypeFlags(), Defaults(), LoadFile() (+7 more)

### Community 7 - "Files"

Cohesion: 0.29
Nodes (7): directoryFiles(), Files(), isGenerated(), isSourceFile(), recursiveFiles(), T, TestAnalyze()

### Community 10 - "analysis.go"

Cohesion: 0.12
Nodes (40): AssignStmt, BlockStmt, DeclStmt, Expr, File, FileSet, FuncType, Ident (+32 more)

## Knowledge Gaps

- **23 isolated node(s):** `version`, `kind`, `build_command`, `consumer_command`, `external-artifact` (+18 more)
  These have ≤1 connection - possible missing edges or undocumented components.
- **2 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions

_Questions this graph is uniquely positioned to answer:_

- **Why does `Thresholds` connect `Thresholds` to `collectFunctionCandidate`, `goreadable/main.go`?**
  _High betweenness centrality (0.203) - this node is a cross-community bridge._
- **Why does `runCommand()` connect `goreadable/main.go` to `Command`, `Thresholds`?**
  _High betweenness centrality (0.090) - this node is a cross-community bridge._
- **Why does `collectFunctionCandidate()` connect `collectFunctionCandidate` to `analysis.go`, `Thresholds`?**
  _High betweenness centrality (0.062) - this node is a cross-community bridge._
- **What connects `version`, `kind`, `build_command` to the rest of the system?**
  _23 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `collectFunctionCandidate` be split into smaller, more focused modules?**
  _Cohesion score 0.14814814814814814 - nodes in this community are weakly interconnected._
- **Should `risks` be split into smaller, more focused modules?**
  _Cohesion score 0.125 - nodes in this community are weakly interconnected._
- **Should `Testing Guidelines` be split into smaller, more focused modules?**
  _Cohesion score 0.10476190476190476 - nodes in this community are weakly interconnected._
