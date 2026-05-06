# Packet Mob Contacts Core Walkthrough

This note is the quickest way to read the extra review model in `packet-mob-contacts-core`.

| Case | Focus | Score | Lane |
| --- | --- | ---: | --- |
| baseline | form pressure | 178 | ship |
| stress | sync drift | 144 | ship |
| edge | local state | 237 | ship |
| recovery | conflict cost | 210 | ship |
| stale | form pressure | 217 | ship |

Start with `edge` and `stress`. They create the widest contrast in this repository's fixture set, which makes them better review anchors than the middle cases.

The next useful expansion would be a malformed fixture around sync drift and conflict cost.
