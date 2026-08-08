# Bootstrap

- Artifact: Go CLI `cmd/goreadable`.
- Consumer contract: `.acceptance-harness/manifest.json` (AC-001–AC-007).
- Scope: static extraction of readability-review candidates; no AI API calls.
- Existing checks: `go test ./...`, `go test -race ./...`, `make fmt`, `make lint`.
- Frozen harness base: `27db120`.
