# quorum-lite

`quorum-lite` keeps a focused Go implementation around distributed systems. The project goal is to model quorum-backed rate limiting with lease expiry and local consensus traces.

## Reason For The Project

This is intentionally local and self-contained so it can be inspected without credentials, services, or seeded history.

## Quorum Lite Review Notes

Start with `lease drift` and `replica lag`. Those cases create the widest score spread in this repo, so they are the best quick check when the model changes.

## What It Does

- `fixtures/domain_review.csv` adds cases for quorum health and lease drift.
- `metadata/domain-review.json` records the same cases in structured form.
- `config/review-profile.json` captures the read order and the two review questions.
- `examples/quorum-lite-walkthrough.md` walks through the case spread.
- The Go code includes a review path for `lease drift` and `replica lag`.
- `docs/field-notes.md` explains the strongest and weakest cases.

## How It Is Put Together

The core code exposes a scoring path and the added review layer uses `signal`, `slack`, `drag`, and `confidence`. The domain terms are `quorum health`, `lease drift`, `replica lag`, and `membership churn`.

The Go code keeps the review rule close to the tests.

## Run It

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/verify.ps1
```

## Check It

That command is also the regression path. It verifies the domain cases and catches mismatches between the CSV, metadata, and code.

## Boundaries

This remains a local project with deterministic fixtures. It does not depend on credentials, hosted services, or live data. Future work should add richer malformed inputs before widening the public API.
