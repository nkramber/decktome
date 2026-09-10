# M-10, a scorer with one weight set per tier, 2026-09-10

The roadmap gate of M-10 reads:

> Gate: the three numbers of D-648 beside gate run 18. If graded-bad falls or the judge agreement rises, and no bar falls, PR-38 builds the scorer. If not, the owner decides again (D-656).

**Verdict: not met.** The judge agreement rises from 8 to 9 of 25, and the Commander precon bar ends one pair under 0.95. The scorer grades real lists far better. It moves the built decks little, because the defect detector holds most of the disputed decks at bad (F-100).

Every number here is free. No run called a provider.

## The runs

Each fit is the code of `main` at `d599f06`, over the meta store of gate run 18 and the card snapshot of 2026-09-04. Throwaway patches made the fits, and none of them merges.

- **The softmax scorer.** One weight set per tier over the same standardized features, every rung at the same weight (D-488). The score stays the expected rung. The defect detector and the synergy check stay as they are.
- **The detector out of the grade.** The detector still sets the score, so every bar reads as before. The tier reads the ladder alone.

## The four fits

| Fit | Graded bad | Judge agreement | Commander precon bar | Commander accuracy |
|---|---|---|---|---|
| Gate run 18, ordinal | 15 of 25 | 8 of 25 | 751 of 788, PASS | 0.65 |
| M-10, softmax | 16 of 25 | 9 of 25 | 748 of 788, FAIL | 0.67 |
| Ordinal, detector out of the grade | 5 of 25 | 7 of 25 | 751 of 788, PASS | 0.52 |
| Softmax, detector out of the grade | 11 of 25 | 9 of 25 | 748 of 788, FAIL | 0.63 |

## The scorer grades real lists far better

The holdout of every fold, in Commander:

| Rung | Graded right, run 18 | Graded right, M-10 |
|---|---|---|
| baseline, the precons | 12 of 189 | 110 of 189 |
| typical, the average decks | 544 of 1,040 | 716 of 1,040 |

Accuracy rises in every format: Commander from 0.65 to 0.67, Standard from 0.46 to 0.49, and Modern from 0.55 to 0.61. The great-over-precon bar rises in every format, and the Commander synergy axis rises from 0.84 to 0.87.

## The cost: a broken copy can rise to a middle rung

The Commander precon bar reads 748 of 788, against 751 in run 18. The synergy axis gains 6 pairs, and the lands axis loses 3 and the colors axis 6. In each new miss the detector passes both decks, and the scorer ranks the copy at or above its precon. One weight set per tier lets a broken copy gain weight on a middle rung, and one ordinal ladder never let it.

## The built decks barely move

Under M-10, 4 of the 25 built decks change grade, and one of them is a Commander deck. The ten Commander decks that run 18 grades bad, and that the judge reads as typical, all stay bad.

## The defect detector holds the disputed decks

A throwaway diagnostic read each built deck under M-10. The detector flags 8 of the 10 disputed decks as broken, at 0.53 to 0.87. Its probability tips the other two, decks 11 and 15, at 0.40. For decks 11 and 13, the softmax scorer alone prefers typical over bad.

The strongest pushes toward bad, counted as the decks where each feature is in the top three:

- `card_rate`: all ten decks
- `unseen_share`: six decks
- `mana_turn_four`: four decks
- `ramp`: four decks

## The detector does two jobs

With the detector out of the grade, graded bad falls to 5 with the ordinal ladder, but the judge agreement falls to 7. Seven disputed decks rise to baseline and none to typical. Decks 8 and 9, which the judge reads as bad, rise to baseline too.

With the softmax scorer and the detector out of the grade, two disputed decks reach typical. Decks 7 and 9, which the judge reads as bad, rise to typical past the judge.

So the detector holds the disputed decks down, and it also grades right the decks the judge reads as bad.

## What waits

The owner asked for the ten disputed decks as a review sheet, to grade them by hand (D-657). The sheet sits outside the repository, in `.local/review/`, as a Markdown file and a CSV. The owner's grades decide the next fix: a fix of the detector, or a fix of the judge bar.
