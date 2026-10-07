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

## Round 2 of P2-1: The replay evidence is not available for verification

Result: full merit.

Evidence: the raw outputs stay under `.local`, and the checkout of a review does not hold that folder. So the review read the reported numbers alone.

Correction: `docs/reference/pr294-replays-2026-10-07.md` holds the evidence. A script read each raw output and wrote the file, so no number came from a person. The file holds:

- The input, the snapshot, and the rule of the pin.
- The 31 themed nonland cards, and a check that each fix shortlist holds the same cards.
- One row for each of the 64 replays. A row names the commit, the Go tree, the count of pins, the bracket, and the recorded power. It also gives the played count, the left-out count, the "mark" count, and the measured spend.
- One row for each deck at bracket 1. A row names the size, the block findings, a hash of the decks file, and each themed card left out.

Regression check: the record gives the five p6 decks 31 played and 0 left out each, at bracket 1. It gives the six pb decks 17 to 23 played and 8 to 14 left out, at bracket 1. Each of the eleven decks holds 100 cards and no block finding.

## Round 3 of P2-1: The record holds no deck profile

Result: full merit. The record proved the counts and not the bracket of each deck.

Evidence: round 3 stopped the loop at three heads (D-826). The owner chose to commit the deck files (D-1202).

Correction: `docs/reference/pr294-replays-2026-10-07` holds the eleven deck files of the bar, unchanged: `p6-3`, `p6-4`, `p6-6`, `p6-7`, `p6-8`, and `pb-1`, `pb-6` to `pb-10`. Each file is one JSON line. The field `deck.profile.bracket` gives the bracket, `deck.power` gives the recorded answer, and `deck.cards` gives each card.

Regression check: the SHA-256 of each file starts with the hash in the record. Each of the eleven files reads `profile.bracket` 1 and `power.bracket` 1.
