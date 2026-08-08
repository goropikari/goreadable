# Design

The implementation is split into independently testable boundaries:

1. `internal/analysis`: parse packages and compute measurements for functions and types.
2. `internal/config`: resolve defaults, `goreadable.json`, and CLI overrides.
3. `internal/report`: stable candidate model plus text and JSON encoders.
4. `internal/diff`: select declarations related to changed Git lines.
5. `cmd/goreadable`: parse arguments, compose boundaries, and map errors to exit status.

The candidate model is independent of output format and AI integration. Git access is
isolated behind a small boundary so package analysis remains deterministic in tests.
