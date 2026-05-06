# Review Journal

The cases below are the review handles I would use before changing the implementation.

The local checks classify each case as `ship`, `watch`, or `hold`. That gives the project a small review vocabulary that matches its mobile workflows focus without claiming live deployment or external usage.

## Cases

- `baseline`: `form pressure`, score 178, lane `ship`
- `stress`: `sync drift`, score 144, lane `ship`
- `edge`: `local state`, score 237, lane `ship`
- `recovery`: `conflict cost`, score 210, lane `ship`
- `stale`: `form pressure`, score 217, lane `ship`

## Note

The useful failure mode here is a wrong decision on a named case, not a vague style disagreement.
