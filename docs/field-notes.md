# Field Notes

The fixture is small on purpose, which makes each domain case carry real weight.

The domain cases cover `quorum health`, `lease drift`, `replica lag`, and `membership churn`. They sit beside the smaller starter fixture so the project has both a compact scoring check and a domain-flavored review check.

`stress` is the strongest case at 252 on `lease drift`. `edge` is the cautious anchor at 144 on `replica lag`.

The local verifier covers this data so the notes stay tied to code.
