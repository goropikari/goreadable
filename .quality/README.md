# Repository quality system

- `baseline.json` is the versioned quality contract. It separates the quality
  model (characteristics, factors, measures, instruments, and limitations)
  from the repository's required gate rules.
- `state.json` records the evidence from the latest completed assessment. Its
  `baseline_version` identifies the contract it was assessed against; it does
  not become current merely because `baseline.json` changes.

The model deliberately treats goreadable's source metrics as review-priority
signals, not as a universal quality score or an automatic release decision.
Thresholds used by the CLI belong to the consumer's configuration. Gate policy,
including required commands and compatibility decisions, belongs in the
baseline.

After implementation, run `$coding-quality-gate` (or `$quality-assessment`) to
produce new evidence against the current baseline.
