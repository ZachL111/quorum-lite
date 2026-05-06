# Quorum Lite Walkthrough

The fixture is intentionally compact, so the review starts with the cases that pull farthest apart.

| Case | Focus | Score | Lane |
| --- | --- | ---: | --- |
| baseline | quorum health | 200 | ship |
| stress | lease drift | 252 | ship |
| edge | replica lag | 144 | ship |
| recovery | membership churn | 212 | ship |
| stale | quorum health | 237 | ship |

Start with `stress` and `edge`. They create the widest contrast in this repository's fixture set, which makes them better review anchors than the middle cases.

The useful comparison is `lease drift` against `replica lag`, not the raw score alone.
