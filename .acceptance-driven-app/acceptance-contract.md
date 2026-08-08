# Acceptance Contract

The frozen black-box contract is maintained in `.acceptance-harness/manifest.json` and
`tests/acceptance/goreadable_test.go`. Every `AC-*` criterion is required and is
validated through the built CLI, not internal packages.

## Constraints

- Go 1.26; no network or AI service dependency.
- Candidate findings are informational; candidate presence exits zero.
- Configuration precedence is CLI flag, then `goreadable.json`, then defaults.
- JSON schema version is 1 and includes source context for AI review.

## Risks

- RISK-001: static thresholds can produce false positives; findings must include measurements and reasons.
- RISK-002: Git diff selection can omit or over-include a surrounding declaration; changed-code behavior is covered by AC-006.
- RISK-003: source snippets may contain repository-sensitive content; output is local and network-free.

## Validation matrix

| Criterion | Command                                                        | Observable result                                 |
| --------- | -------------------------------------------------------------- | ------------------------------------------------- |
| AC-001    | `go test ./tests/acceptance -run TestGoreadableCLI`            | Function and struct candidates are emitted        |
| AC-002    | `go test ./tests/acceptance -run TestGoreadableModuleIdentity` | Module and build identity are valid               |
| AC-003    | JSON acceptance subtest                                        | Versioned candidate schema and source context     |
| AC-004    | Text acceptance subtest                                        | Candidate, location, and reason are readable      |
| AC-005    | Configuration acceptance subtest                               | CLI threshold overrides file threshold            |
| AC-006    | Git diff acceptance subtest                                    | Only changed candidate is emitted                 |
| AC-007    | Exit/error acceptance subtests                                 | Findings exit zero; invalid options exit non-zero |
