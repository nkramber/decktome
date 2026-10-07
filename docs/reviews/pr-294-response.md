# The author response to the review of #294

Author provider: Claude Code.

## P2-1: The baseline replays use a different bracket

Result: full merit.

Evidence: the base replays of the record ran at `7079e60` with no pin. The classifier lost "1 Exhibition" in both, so both built at bracket 3. The fix replays ran with the pin of D-1201 and built at bracket 1. So the two sides read different shortlists, and the comparison did not hold the condition of D-1199.

Correction: ten new replays read the base `ea792c3`, the merge base of the pull request. They read the same pin, the same snapshot of 2026-10-06, and the same four turns. Six built a deck, each one at bracket 1. Four built no deck (OQ-98).

| Side | Bracket | Themed nonland played | Themed left out |
|---|---|---|---|
| Base `ea792c3`, six pinned replays | 1 | 17, 19, 20, 22, 23, 23 | 14, 12, 11, 9, 8, 8 |
| Live deck | 1 | 26 | 5 |
| Fix `76a081c`, five pinned replays | 1 | 31 each | 0 each |

The worst fix deck played 31 and left out none. The best base deck played 23 and left out 8, and the live deck played 26 and left out 5. So the bar of D-1199 holds with each side at bracket 1. The record of D-1201, the PR-136 gate, the hand-off, and the body name the pinned base replays.

Regression check: each base and fix deck of the table reads bracket 1 in its `profile`. The ten base replays cost $0.3288, measured. The replays of the pull request cost $1.7986 over 64 runs, of the cap of $3.00.

## A correction to the PR comments of the record

The record says that the archive now contains the checkpoint. It does not. The answer on the Gitar thread refuted that part. The archive holds the resume section of the base. The checkpoint was a draft of this pull request. The hand-off restored the probe steps alone.
