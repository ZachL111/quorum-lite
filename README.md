# quorum-lite

`quorum-lite` is a focused Go codebase around model quorum-backed rate limiting with lease expiry and local consensus traces. It is meant to be easy to inspect, run, and extend without a hosted service.

## Quorum Lite Walkthrough

I would read the project from the outside in: command, fixture, model, then roadmap. That keeps the distributed systems idea grounded in files that can be checked locally.

## Capabilities

- Includes extended examples for lease timing, including `recovery` and `degraded`.
- Documents message ordering tradeoffs in `docs/operations.md`.
- Runs locally with a single verification command and no external credentials.
- Stores project constants and verification metadata in `metadata/project.json`.
- Adds a repository audit script that checks structure before running the language verifier.

## Reason For The Project

The repository exists to keep a technical idea small enough to reason about. The implementation avoids external dependencies where possible, then uses fixtures to make changes easy to review.

## Where Things Live

- `policy`: Go package with the core model
- `cmd`: small command entry point
- `fixtures`: compact golden scenarios
- `examples`: expanded scenario set
- `metadata`: project constants and verification metadata
- `docs`: operations and extension notes
- `scripts`: local verification and audit commands
- `go.mod`: Go module metadata

## How It Is Put Together

The design is intentionally direct: parse or construct a signal, score it, classify it, and verify the expected branch. This makes the repository useful for studying distributed systems behavior without needing a service or database unless the language project itself is SQL. The Go layout uses small packages and table-oriented tests so the behavior stays easy to follow.

## Command Examples

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/verify.ps1
```

This runs the language-level build or test path against the compact fixture set.

## Data Notes

The examples are meant to be readable before they are exhaustive. They cover enough variation to show how latency and risk can pull a decision below the threshold.

## Check The Work

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/audit.ps1
```

The audit command checks repository structure and README constraints before it delegates to the verifier.

## Tradeoffs

The scoring model is simple by design. More domain-specific behavior should be added through explicit adapters or extra fixture classes rather than hidden constants.

## Possible Extensions

- Add a comparison mode that shows how decisions change when one signal is adjusted.
- Add a loader for `examples/extended_cases.csv` and promote selected cases into the language test suite.
- Add a short report command that prints the score breakdown for a single scenario.
- Add one more distributed systems fixture that focuses on a malformed or borderline input.

## Getting It Running

The only required setup is the local Go toolchain. After cloning, stay in the repo root so fixture paths resolve correctly.
