# Review Journal

This journal records the domain cases that matter before widening the public API.

The local checks classify each case as `ship`, `watch`, or `hold`. That gives the project a small review vocabulary that matches its distributed systems focus without claiming live deployment or external usage.

## Cases

- `baseline`: `quorum health`, score 200, lane `ship`
- `stress`: `lease drift`, score 252, lane `ship`
- `edge`: `replica lag`, score 144, lane `ship`
- `recovery`: `membership churn`, score 212, lane `ship`
- `stale`: `quorum health`, score 237, lane `ship`

## Note

The useful failure mode here is a wrong decision on a named case, not a vague style disagreement.
