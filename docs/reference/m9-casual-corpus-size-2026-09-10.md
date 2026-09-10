# M-9, the casual corpus of PR-29 at three sizes, 2026-09-10

The roadmap gate of M-9 reads:

> The three numbers of D-648 at each size, beside gate run 18. If graded-bad and the judge agreement improve as the corpus grows, PR-35 builds with the two features. If not, the owner decides again (D-654). The two features F-94 names stay out.

**Verdict: M-9 refutes the premise of D-650.** A thicker casual corpus does not rescue the casual features of PR-29. The Commander casual lists grow from 279 to 1,040 over the three sizes. Over that growth, the built decks graded bad go from 16 to 17, and the judge agreement from 7 to 6. The features cost the reader more as the corpus grows, and not less.

Every number here is free. No run called a provider.

## The runs

Each run is the code of `main` at `302858a`, with the synergy check of D-653. It reads the meta store of gate run 18 and the card snapshot of 2026-09-04. Two throwaway patches made the fits, and neither one merges.

- **The casual features.** The diff of `corpus.go`, `features.go`, `model.go`, and `score.go` from the closed branch of PR-29 applied to `main` with no conflict, and its tests pass. It holds `casual_rate` and `casual_synergy` alone. The two features F-94 names never reached that branch.
- **The size.** A hash keeps every second or every fourth EDHREC average deck of the Commander fit. The precons and the tournament lists stay whole, and so do Standard and Modern.

The built decks are the 25 of deck gate run 16. The judge's tiers come from judge lane run 5, so every fit reads the same judge answers.

## The six fits

The casual lists count every Commander EDHREC average deck of the fit, and the train column counts fold 0.

| Code | Casual lists | Train | Synergy axis | Precon bar | Graded bad | Judge agreement |
|---|---|---|---|---|---|---|
| `main` | 1,040 | 854 | 0.84 of 232 | 0.95 | 15 of 25 | 8 of 25 |
| `main` | 523 | 425 | 0.79 of 229 | 0.94 | 15 of 25 | 8 of 25 |
| `main` | 279 | 213 | 0.80 of 223 | 0.94 | 16 of 25 | 7 of 25 |
| PR-29 | 1,040 | 854 | 0.89 of 232 | 0.97 | 17 of 25 | 6 of 25 |
| PR-29 | 523 | 425 | 0.85 of 229 | 0.96 | 17 of 25 | 6 of 25 |
| PR-29 | 279 | 213 | 0.86 of 223 | 0.96 | 16 of 25 | 7 of 25 |

The Commander weights of the two features, with `card_rate` beside them:

| Casual lists | `casual_rate` | `casual_synergy` | `card_rate` |
|---|---|---|---|
| 1,040 | +0.217 | +0.070 | +0.990 |
| 523 | +0.254 | +0.017 | +0.983 |
| 279 | +0.257 | -0.028 | +0.918 |

## What the numbers say

**The features cost the reader more as the corpus grows.** At 279 lists the features and `main` read the same: 16 graded bad and 7 agreement. At 523 and at 1,040 lists the features read 17 and 6, and `main` reads 15 and 8. So the gap between the two codes opens with the size of the casual corpus.

**The fit leans on `casual_synergy` as the corpus grows.** Its weight goes from -0.028 to +0.070. That is the change the premise of D-650 expected, and it moves the reader-facing numbers the wrong way.

**The bar reads the other way, as it did for PR-29.** The features lift the synergy axis by 0.05 to 0.07 at every size, and the precon bar by 0.01 to 0.02. The proxy and the reader disagree about the casual features at every size (D-648).

## The decks that move

The runs at half the lists grade every deck as the runs at all the lists do.

| Deck | Judge | `main`, all | `main`, quarter | PR-29, all | PR-29, quarter |
|---|---|---|---|---|---|
| 10. Standard aggro, tournament | baseline | baseline | baseline | typical | typical |
| 13. owned first, and the commander is not owned | typical | bad | bad | bad | typical |
| 18. upgrade a precon, any card | baseline | baseline | bad | bad | bad |
| 25. use no card of an owned precon | typical | baseline | baseline | bad | bad |

Deck 13 is the one deck the smallest corpus grades as the judge does, and the larger corpora grade bad. The features drop decks 18 and 25 to bad at every size. Deck 10 is a Standard deck: the features reach the 60-card formats too, and they lift it one rung past the judge.

## What this item does not measure

- **Growth past 1,040 lists.** A shrink test predicts growth by its trend alone. From 279 to 1,040 lists the trend is flat or worse, and nothing here says it turns.
- **The judge.** One judge run gives every tier. A second run moves some tiers (D-230), and it moves them for both codes alike.
- **The 60-card casual lists.** The size patch touched Commander alone.

## What it means for PR-35

PR-35 has no case with the casual features, and F-97 says it has none without them. The owner parked PR-35 and chose PR-30, the commander reference, as the next item (D-655).
