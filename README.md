# packet-mob-contacts-core

`packet-mob-contacts-core` keeps a focused Go implementation around mobile workflows. The project goal is to create a Go reference implementation for contacts workflows, centered on security rule linting, safe and unsafe fixtures, and remediation hints.

## Why This Exists

The point is to make a small domain rule concrete enough that a reader can change it and immediately see what broke.

## Packet Mob Contacts Core Review Notes

For a quick review, compare `local state` with `sync drift` before reading the middle cases.

## Capabilities

- `fixtures/domain_review.csv` adds cases for form pressure and sync drift.
- `metadata/domain-review.json` records the same cases in structured form.
- `config/review-profile.json` captures the read order and the two review questions.
- `examples/packet-mob-contacts-walkthrough.md` walks through the case spread.
- The Go code includes a review path for `local state` and `sync drift`.
- `docs/field-notes.md` explains the strongest and weakest cases.

## Implementation Shape

The fixture data drives the tests. The code stays thin, while `metadata/domain-review.json` and `config/review-profile.json` explain what each case is meant to protect.

The Go code keeps the review rule close to the tests.

## Local Usage

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/verify.ps1
```

## Verification

The same command runs the local verification path. The highest-scoring domain case is `edge` at 237, which lands in `ship`. The most cautious case is `stress` at 144, which lands in `ship`.

## Roadmap

The fixture set is small enough to audit by hand. The next useful expansion is malformed input coverage, not extra surface area.
