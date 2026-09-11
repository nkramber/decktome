# M-11, the detector variants and the tier rule, 2026-09-11

The roadmap gate of M-11 reads:

> Gate: the three numbers of D-648 beside gate run 19, which reads 16 of 25 graded bad and 9 of 25 judge agreement. If graded bad falls or the judge agreement rises, and no bar falls, PR-40 builds the best variant. If no variant meets the gate, a rules detector comes next (D-666).

**Verdict: no detector variant meets the gate.** Two of the fifteen detector fits pass every bar, and neither one lowers graded bad or raises the agreement. **A change of the tier rule meets the gate, and it is not one of the six variants.** Every bar reads the score, and a change of the tier alone keeps the score. Over the control detector, `tier` lowers graded bad from 16 to 14. `notier` lowers it to 8 and raises the agreement to 10. Both lower the Commander holdout accuracy, and the gate does not read it.

**The owner chose the rules detector of D-666 first, with no change of the tier (D-675).**

Every number here is free. No run called a provider.

## The runs

Each fit is the code of `33d6691`, over the meta store of gate run 19 and the card snapshot of 2026-09-04. Throwaway patches made the fits, and none of them merges. The control reproduces gate run 19 in every number.

The first patch adds five flags through the variable `M11`, for the Commander detector alone (D-667):

- `moved`: the detector reads a feature only when a break moves it by 0.5 standard deviations or more.
- `base`: the top lists leave the negatives of the detector.
- `colors`: the detector reads no `source_spread` and no `tapped_share`.
- `nounseen`: the detector reads no `unseen_share`.
- `axis`: one detector per axis, over the features its own break moves. The detector flags a deck when one of them flags it.

Fifteen combinations cover the six variants of the roadmap. A second patch adds the two tier flags of the section "The tier rule".

## The fifteen detector fits

Own is the Commander precon bar, the precon over its own copy, and the bar asks for 749 of 788. Disputed counts the ten decks of D-659 that the detector flags. Owner reads the tier matches and the mean error in rungs against the grades of the owner.

| Fit | Verdict | Own | Colors axis | Great over precon | Graded bad | Agreement | Disputed | Owner |
|---|---|---|---|---|---|---|---|---|
| control | PASS | 752 | 178 of 178 | 0.96 | 16 | 9 | 8 | 1, 1.4 |
| moved | PASS | 751 | 178 of 178 | 0.96 | 17 | 9 | 9 | 1, 1.4 |
| base | FAIL | 752 | 178 of 178 | 0.88 | 16 | 9 | 8 | 1, 1.4 |
| moved, base | FAIL | 752 | 178 of 178 | 0.51 | 17 | 9 | 9 | 1, 1.4 |
| moved, colors | FAIL | 720 | 143 of 178 | 0.97 | 18 | 8 | 7 | 1, 1.4 |
| base, colors | FAIL | 757 | 169 of 178 | 0.83 | 16 | 9 | 5 | 1, 1.4 |
| moved, base, colors | FAIL | 746 | 158 of 178 | 0.41 | 18 | 8 | 6 | 1, 1.4 |
| axis | FAIL | 741 | 163 of 178 | 0.91 | 19 | 7 | 10 | 1, 1.4 |
| axis, base | FAIL | 750 | 174 of 178 | 0.31 | 19 | 7 | 8 | 1, 1.4 |
| axis, colors | FAIL | 714 | 140 of 178 | 0.90 | 19 | 7 | 10 | 1, 1.4 |
| axis, base, colors | FAIL | 739 | 167 of 178 | 0.34 | 20 | 6 | 8 | 1, 1.4 |
| moved, colors, nounseen | FAIL | 710 | 137 of 178 | 0.97 | 18 | 8 | 5 | 1, 1.4 |
| moved, base, colors, nounseen | FAIL | 747 | 158 of 178 | 0.38 | 18 | 8 | 5 | 1, 1.4 |
| axis, colors, nounseen | FAIL | 735 | 159 of 178 | 0.87 | 19 | 7 | 7 | 1, 1.4 |
| axis, base, colors, nounseen | FAIL | 750 | 168 of 178 | 0.31 | 20 | 6 | 8 | 1, 1.4 |

Standard reads 1.00 of 40 and Modern 0.99 of 1635 in every fit.

- **Moved features** pass every bar and grade one more deck bad. With `card_rate`, `fast_mana`, and `game_changer` out, `mana_turn_four` takes the gap between the top lists and the precons, at -1.62.
- **Base negatives** fail the great-over-precon bar in every combination, at 0.31 to 0.88. That is the risk of gate run 2 (D-485). A detector with no top list among its negatives flags the top lists.
- **Colors by need** lowers the colors axis in every fit, to 137 to 169 of 178. Without `base` the precon bar fails, and with `base` the great-over-precon bar fails.
- **One detector per axis** picks low cuts: 0.20 for lands and 0.22 for colors. It flags all ten disputed decks in two fits, and it grades 19 or 20 decks bad.
- **No unseen share** moves no reader-facing number. Graded bad and the agreement read the same with and without it, in all four pairs of fits.

Every fit flags deck 22, the one real defect of the ten. In eight of the nine fits with `colors`, `color_sources` joins its three strongest pushes, at +0.44 to +0.52. The exception is `axis, colors`.

## Why no detector variant moves a disputed deck (F-115)

Every disputed deck grades bad in all fifteen fits, so the owner column never moves. Three fits pass three more disputed decks under the cut, and each of those decks still grades bad.

**The cause is the tier rule, and not the flag.** The grade moves the detector probability onto the bottom rung for every deck, and the tier is the rung with the most mass (`grade` in `go/internal/quality/fit.go`). The flag decides the score band alone.

The ladder spreads its mass over five rungs, so a small probability wins the bottom rung. In the fit `base, colors`, deck 15 grades bad at 0.27 and deck 16 at 0.28. The cut of that fit is 0.52.

The mix came with PR-14B (#58, 2026-09-03), and no decision records it. D-485 names the score floor of a flagged deck, and the comment of `grade` cites D-485. F-100 saw the effect on decks 11 and 15, and the plan of M-11 read the flag alone.

## The tier rule

Every bar reads the score. A change of the tier alone keeps the score, so every bar reads as its base fit. It moves graded bad, the agreement, and the holdout accuracy. The second patch adds two flags, for Commander alone:

- `tier`: an unflagged deck takes the tier of the ladder alone. A flagged deck keeps the tier it has today.
- `notier`: no deck takes the detector into its tier.

| Fit | Verdict | Own | Great over precon | Graded bad | Agreement | Owner | Commander accuracy |
|---|---|---|---|---|---|---|---|
| control | PASS | 752 | 0.96 | 16 | 9 | 1, 1.4 | 0.65 |
| tier | PASS | 752 | 0.96 | 14 | 9 | 2, 1.2 | 0.62 |
| moved, tier | PASS | 751 | 0.96 | 15 | 9 | 1, 1.3 | 0.61 |
| notier | PASS | 752 | 0.96 | 8 | 10 | 4, 0.9 | 0.52 |
| base, colors | FAIL | 757 | 0.83 | 16 | 9 | 1, 1.4 | 0.58 |
| base, colors, tier | FAIL | 757 | 0.83 | 12 | 10 | 2, 1.1 | 0.58 |

**`tier` meets the gate.** Decks 11 and 15 rise from bad to baseline. The owner grades deck 15 baseline and deck 11 typical. The ladder alone reads deck 11 at 0.38 baseline against 0.34 typical, and deck 15 at 0.40 baseline.

`moved, tier` grades one more deck bad than `tier`, because `moved` flags deck 15 at 0.50.

**`notier` meets the gate too, and it costs more.** Decks 3, 4, 11, 13, 15, 20, 21, and 22 rise to baseline. The agreement gains deck 20, which the judge reads as baseline. The owner grades decks 3, 13, 15, and 21 baseline. **Deck 22, the one real defect, rises to baseline too.**

**The ladder alone still holds five Commander decks at bad.** Decks 1, 12, 14, 16, and 19 read bad on the ladder alone, at 0.50 to 0.59. The owner grades decks 1, 14, and 16 typical. So a perfect flag still leaves three of the ten disputed decks at bad, and that gap sits in the ladder (F-106, PR-38).

Under `tier`, a detector that flags deck 22 alone of the built decks grades every other deck as `notier` does, and deck 22 bad. That derived read is 9 graded bad, an agreement of 10, and 5 owner matches. No fit measured it.

`base, colors, tier` grades 12 decks bad at an agreement of 10. It fails the great-over-precon bar at 0.83, as its base fit does.

The review sheet of D-659 notes that the ladder alone grades decks 11 and 15 bad. That note read the fit of 2026-09-07. On the fit of gate run 19, the ladder alone grades both decks baseline.

## The cost the gate does not read

The Commander holdout accuracy falls from 0.65 to 0.62 under `tier`, and to 0.52 under `notier`. The whole fall sits in the broken copies, and every real rung grades better.

| Holdout rung | Bad, control | Bad, tier | Bad, notier | Right, control | Right, tier | Right, notier |
|---|---|---|---|---|---|---|
| bad, 5,905 broken copies | 5,730 | 5,135 | 3,669 | 5,730 | 5,135 | 3,669 |
| baseline, 189 precons | 119 | 65 | 54 | 13 | 64 | 75 |
| typical, 1,040 average decks | 419 | 115 | 46 | 543 | 629 | 632 |
| good, 4,000 lists | 257 | 128 | 41 | 1,403 | 1,408 | 1,408 |
| great, 4,000 lists | 208 | 91 | 27 | 2,140 | 2,140 | 2,140 |

Under `tier`, 604 fewer real lists grade bad, and 595 more broken copies grade above bad. Of the 770 copies above bad, 417 grade baseline, 334 typical, and 19 good.

Under `notier`, 835 fewer real lists grade bad, and 2,061 more broken copies grade above bad. Of the 2,236 copies above bad, 1,722 grade baseline, 494 typical, and 20 good.

The audit of the bar moves too. The bar loses 36 synergy pairs, and in 7 of them the control grades the copy a higher tier than its precon. That count reads 18 under `tier` and 19 under `notier`. The bar reads the score, so it holds.

A reader meets the cost as a broken build that the detector passes, or under `notier` as any broken build. That build takes the tier of the ladder, where today the probability can tip it to bad.

## The decisions (D-674, D-675)

The plan of M-11 read the flag alone, so its gate did not foresee a change of the tier. A reader sees the tier and three reasons, and never the score. The reasons name the signals of the detector for a flagged deck alone. The owner weighed three answers:

- **`tier`, then the rules detector.** No bar can move, and deck 22 stays bad. A later fix of a flag then moves the tier directly. Two built decks move now, and 595 more broken copies grade above bad.
- **`notier`.** Graded bad falls to 8 and the agreement rises to 10, and no bar can move. Deck 22 rises to baseline, and a flagged deck shows a middle tier with the reasons of the detector. A later detector reaches the score alone.
- **The rules detector first.** It aims at the root cause of D-659: seven wrong flags, and a colors fault the fitted detector can not see. It changes the score, so every bar is at risk, and the precon bar has three pairs of room. A detector probability still tips an unflagged deck to bad, so its fits read with and without `tier`.

**The owner chose the rules detector first, with no change of the tier (D-675).** Every change that moves a tier reports the broken copies graded bad beside the three numbers of D-648, as information (D-674).
