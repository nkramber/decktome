# M-12, the rules detector fits, 2026-09-11

The roadmap gate of M-12 reads:

> Gate: each fit reports the three numbers of D-648 beside gate run 19, the broken copies graded bad (D-674), and every bar. The owner picks no fit that lowers a bar. The owner picks the design and the thresholds on the numbers (D-676, D-677), and PR-40 builds the pick.

**Verdict: design A passes in all eight combinations, and design B fails in all eight.** Design A keeps every bar as gate run 19 reads it. It lowers the built decks graded bad from 16 to 9 and raises the judge agreement from 9 to 10. It matches 5 of the owner's 10 grades, against 1. Design B lowers the precon bar to 659 to 662 of 788, and it grades the built decks as gate run 19 does.

Every number here is free. No run called a provider.

## The runs

Each fit is the code of `4b56949`, over the meta store of gate run 19 and the card snapshot of 2026-09-04. Two throwaway patches make the fits: the axis detectors of M-11 and the rules of M-12. Neither merges. The control reproduces gate run 19 in every number.

The checks read Commander alone (D-667):

- **Lands**: Karsten's need less the lands sits at 8 or more, or at 10 or more. The cheap count follows Karsten's text rules, and a test matches all 33 cards his article names.
- **Curve**: the average mana value of the nonland cards sits over 4.2, or over 4.4.
- **Colors**: the worst color's sources sit under 0.75 or under 0.8 of their need. A deck with no color row reads no colors check.

The two designs:

- **Design A**: a flagged deck grades bad, and every other deck takes the tier of the ladder alone. The score stays the score of today.
- **Design B**: the rules or the synergy detector of the axis mode of M-11 flag a deck. A flagged deck grades bad and scores under the floor, lower with its distance past the cut or with the probability of the detector. An unflagged deck takes the tier and the score of the ladder.

## Design A

| Fit | Precon bar | Great over precon | Graded bad | Agreement | Owner | Commander accuracy | Broken copies graded bad |
|---|---|---|---|---|---|---|---|
| control | 752 of 788 | 0.96 | 16 | 9 | 1, 1.4 | 0.65 | 5,730 of 5,905 |
| shortfall 8, curve 4.2, colors 0.75 | 752 of 788 | 0.96 | 9 | 10 | 5, 0.8 | 0.52 | 4,194 |
| shortfall 8, curve 4.2, colors 0.8 | 752 of 788 | 0.96 | 9 | 10 | 5, 0.8 | 0.52 | 4,236 |
| shortfall 8, curve 4.4, colors 0.75 | 752 of 788 | 0.96 | 9 | 10 | 5, 0.8 | 0.52 | 4,181 |
| shortfall 8, curve 4.4, colors 0.8 | 752 of 788 | 0.96 | 9 | 10 | 5, 0.8 | 0.52 | 4,223 |
| shortfall 10, curve 4.2, colors 0.75 | 752 of 788 | 0.96 | 9 | 10 | 5, 0.8 | 0.55 | 4,157 |
| shortfall 10, curve 4.2, colors 0.8 | 752 of 788 | 0.96 | 9 | 10 | 5, 0.8 | 0.54 | 4,200 |
| shortfall 10, curve 4.4, colors 0.75 | 752 of 788 | 0.96 | 9 | 10 | 5, 0.8 | 0.55 | 4,141 |
| shortfall 10, curve 4.4, colors 0.8 | 752 of 788 | 0.96 | 9 | 10 | 5, 0.8 | 0.54 | 4,184 |

Every axis of the precon bar reads as gate run 19: 189, 189, 178, and 196 of 232. Standard reads 1.00 of 40 and Modern 0.99 of 1635 in every fit.

**Every A fit grades the built decks the same.** The rules flag decks 19 and 22 alone. The Commander decks graded bad are 1, 12, 14, 16, 19, and 22, and decks 7, 8, and 9 of Modern and Standard stay bad. The owner's grades match on decks 3, 13, 15, 21, and 22. Deck 22, the one real defect, stays bad. The read of the plan, 9 graded bad, 10 in agreement, and 5 owner matches, holds in every fit.

**The thresholds differ in the real lists they grade bad.** The Commander holdout holds 189 precons, 1,040 average decks, and 4,000 lists in each top tier.

| Fit | Precons graded bad | Average decks graded bad | Good lists graded bad | Great lists graded bad | Precons graded right |
|---|---|---|---|---|---|
| control | 119 | 419 | 257 | 208 | 13 |
| shortfall 8, curve 4.2, colors 0.75 | 63 | 83 | 687 | 701 | 69 |
| shortfall 8, curve 4.2, colors 0.8 | 66 | 89 | 854 | 839 | 67 |
| shortfall 8, curve 4.4, colors 0.75 | 58 | 73 | 680 | 694 | 73 |
| shortfall 8, curve 4.4, colors 0.8 | 61 | 79 | 848 | 832 | 71 |
| shortfall 10, curve 4.2, colors 0.75 | 63 | 82 | 257 | 246 | 69 |
| shortfall 10, curve 4.2, colors 0.8 | 66 | 88 | 455 | 400 | 67 |
| shortfall 10, curve 4.4, colors 0.75 | 58 | 71 | 247 | 239 | 73 |
| shortfall 10, curve 4.4, colors 0.8 | 61 | 77 | 446 | 393 | 71 |

- **A shortfall of 8 grades about three times as many top lists bad as the control.** The land formula misreads the cEDH lists (F-116).
- **Colors under 0.8 grades about 200 more lists bad in each top tier than colors under 0.75.**
- **Shortfall 10, curve 4.4, and colors 0.75 grades the fewest real lists bad**: 615, against 1,003 in the control. It grades 73 of 189 precons right, against 13.
- **The cost sits in the broken copies.** That fit grades 4,141 of 5,905 bad, against 5,730 in the control and 4,236 at most in design A. The tier of A reads no synergy break, so the ladder alone catches a synergy copy.

## Design B

| Fit | Precon bar | Lands | Curve | Colors | Synergy | Great over precon | Graded bad | Agreement |
|---|---|---|---|---|---|---|---|---|
| shortfall 8, curve 4.2, colors 0.75 | 659 of 788 | 187 | 187 | 69 | 216 | 0.92 | 16 | 9 |
| shortfall 8, curve 4.2, colors 0.8 | 660 of 788 | 187 | 187 | 70 | 216 | 0.91 | 16 | 9 |
| shortfall 8, curve 4.4, colors 0.75 | 660 of 788 | 187 | 187 | 70 | 216 | 0.92 | 16 | 9 |
| shortfall 8, curve 4.4, colors 0.8 | 662 of 788 | 187 | 187 | 71 | 217 | 0.91 | 16 | 9 |
| shortfall 10, curve 4.2, colors 0.75 | 659 of 788 | 187 | 187 | 69 | 216 | 0.96 | 16 | 9 |
| shortfall 10, curve 4.2, colors 0.8 | 660 of 788 | 187 | 187 | 70 | 216 | 0.95 | 16 | 9 |
| shortfall 10, curve 4.4, colors 0.75 | 660 of 788 | 187 | 187 | 70 | 216 | 0.96 | 16 | 9 |
| shortfall 10, curve 4.4, colors 0.8 | 662 of 788 | 187 | 187 | 71 | 217 | 0.95 | 16 | 9 |

- **The precon bar fails in every fit.** The colors axis falls from 178 pairs to 69 to 71. In the fit of shortfall 10, curve 4.4, and colors 0.75, both decks carry a flag in 106 of the 108 lost colors pairs. The rules flag about 10 precons at those cuts, so the synergy detector flags most of the other precons, at its cut of 0.40. The flagged band then orders the pair by the probability of the detector against the severity of the rules. The two scales do not compare.
- **The great-over-precon bar holds at a shortfall of 10**, at 0.95 to 0.96, and falls to 0.91 to 0.92 at 8.
- **Design B grades the built decks as gate run 19 does**: 16 graded bad, 9 in agreement, and 1 owner match. Decks 3, 4, 11, 13, 15, 20, and 21 grade bad under B and not under A, so the synergy detector flags them. It reads the features the synergy break moves, `unseen_share` among them, so it brings back the fault of D-659 (F-106, F-108).

## The owner decision (D-678)

The owner picked design A at a shortfall of 10, curve 4.4, and colors 0.75. It grades the fewest real lists bad, and it grades the most precons right. PR-40 builds it for Commander alone. The fitted detector keeps the score, so every bar stays. The reasons of a flagged deck name the check.

The costs of D-676 go with it. Two detectors live in the code, and the bars guard a score no reader sees. The tier reads no synergy break, so the ladder alone catches a synergy copy. The broken copies graded bad fall from 5,730 to 4,141 of 5,905, and PR-40 reports that count (D-674).
