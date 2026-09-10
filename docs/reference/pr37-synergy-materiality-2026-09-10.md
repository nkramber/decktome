# PR-37, the materiality check on the synergy break, 2026-09-10

The roadmap gate of PR-37 reads:

> The three numbers of D-648. The Commander synergy axis over the pairs that remain, the count of built decks graded bad, and the judge agreement. The document names how many pairs the check dropped, per format, and the `bad` tier count before and after.

**Verdict: met.** The check drops 240 of the 2958 Commander synergy copies, and 157 of the 389 own pairs of the synergy axis. The Commander synergy axis reads **0.84 of 232**, against 0.70 of 389. The built decks graded bad stay at **15 of 25**, and the judge agreement stays at **8 of 25**. Quality gate run 18 reads **PASS**.

**The PASS comes from the population of the bar, and not from a better model.** An estimate from the audit of the control predicts the pairs each floor keeps, to within two pairs. So the refit reads the kept pairs as the control did. The bar now reads pairs that carry a defect, and the model is the same model.

Every number here is free. No run called a provider.

## The check

The synergy break replaces half the spells of a list with cards the real lists play (D-488). D-485 says a break must be material. It checks the copies axis and the colors axis, and a list that fails either check breaks on synergy. Synergy checked nothing (F-95).

The check measures the move after the break (D-652). In each fold, the fit reads the synergy feature of every synergy copy in the corpus of that fold. It compares the copy with the list it came from, and it drops the copy when the fall sits under the floor.

- **The unit** is the standard deviation of `synergy` over the real training lists of the fold. No copy takes part in it, so the unit stays the same at every floor. Commander reads 0.5841 on fold 0.
- **The floor** is 0.10 of that unit (D-653). A break that lowered nothing leaves at any floor.
- **The formats:** Commander alone (D-653). Standard and Modern keep every copy (F-96).
- **A request that passes no check makes no copy.** A copies or a colors request that falls to synergy adds no row when its synergy copy then fails. The lands and the curve copies stay, because each one carries a material defect.

The check reads the corpus of the fold, so no holdout list decides a training row. A copy is holdout in the fold of the list it came from, so the count reads each dropped copy once.

## The control reproduces run 17

The control is the code of `main` at `223a28f`, over the meta store of run 17 and the card snapshot of 2026-09-04. Its document matches gate run 17 in every line but the date, the model version, and the time. So the store did not move, and every number here compares with run 17.

The code of this item with the check off reproduces the control in every bar. It also reads 15 of 25 built decks graded bad and 8 of 25 judge agreement. The PR-29 document measured those two numbers by hand for its control, and the new section of the gate reads the same numbers.

## The six floors

Each row is the code of this item over the same store, with the check on in all three formats. The Commander columns read the Commander fit, and the last three columns read all 25 built decks of deck gate run 16.

| Floor | Commander copies dropped | Commander synergy axis | Commander precon bar | Modern synergy copies dropped | Graded bad | Judge agreement | Verdict |
|---|---|---|---|---|---|---|---|
| off | 0 of 2958 | 0.70 of 389 | 0.875 of 945 | 0 of 1109 | 15 of 25 | 8 of 25 | FAIL |
| 0.05 | 113 | 0.78 of 301 | 0.924 of 857 | 851 | 14 | 7 | FAIL |
| 0.10 | 240 | 0.84 of 232 | 0.953 of 788 | 867 | 14 | 7 | PASS |
| 0.25 | 687 | 0.93 of 141 | 0.986 of 697 | 878 | 14 | 7 | FAIL |
| 0.50 | 1596 | 0.96 of 116 | 0.993 of 672 | 894 | 14 | 7 | FAIL |
| 1.00 | 2494 | 0.96 of 83 | 0.995 of 639 | 926 | 13 | 8 | FAIL |

Every run but the one at 0.10 fails the Standard great-over-precon bar. It reads 0.899 at 0.05, 0.896 at 0.25, 0.893 at 0.50, and 0.895 at 1.00, against a bar of 0.90. The check off reads 0.906, and the run at 0.10 reads 0.902. So the check in Standard moves a bar it was never meant to touch. The run at 0.05 also fails the Commander precon bar.

**In Commander, no built deck changes its grade at any floor up to 0.50.** Every change in the last two columns up to 0.50 comes from Modern. At 1.00 one Commander deck moves: deck 11, from bad to typical, where the judge reads typical.

## Why 0.10

The audit of the control holds every Commander own pair with the move of `synergy` in the unit of the fit's scaler. Scaled to the unit of the check, it predicts the kept pairs of each run: 301, 231, 141, 114, and 83. The runs kept 301, 232, 141, 116, and 83. So the audit says which pairs each floor drops, and how often the model won them before the check.

| Floor | Pairs dropped | Won before | Weak, middle, strong | The step from the floor below |
|---|---|---|---|---|
| 0.05 | 88 | 36 percent | 31, 40, 17 | - |
| 0.10 | 158 | 47 percent | 44, 76, 38 | 70 pairs, won 60 percent |
| 0.25 | 248 | 56 percent | 52, 119, 77 | 90 pairs, won 72 percent |
| 0.50 | 275 | 59 percent | 52, 131, 92 | 27 pairs, won 81 percent |
| 1.00 | 306 | 62 percent | 54, 132, 120 | 31 pairs, won 90 percent |

The bands are the ones of M-8: a precon score at or under 0.05, from 0.05 to 0.30, and above 0.30.

**At 0.10 the dropped pairs are a coin flip.** The model won them 47 percent of the time, so they carry no defect the model reads. **Each step above 0.10 drops pairs the model does read**, at 72 to 90 percent. A higher floor lifts the bar with pairs that carry a defect, and that is no measurement.

At 0.05 the model won the dropped pairs 36 percent of the time. It read most of those copies as the better deck.

The owner chose 0.10 over 0.25, 0.05, and 0.50 (D-653).

## The premise of D-652 was half right

D-652 expected the check to drop the weak precons M-8 named, 54 of 389 pairs. **The check drops 157 pairs, from every band.** The synergy break moves the synergy feature little for most precons. The median fall reads 0.13 of the unit of the check. The weak precons are the far end of that spread, and not a group apart.

## Modern and Standard keep every copy

In Modern the break does not lower the synergy feature in 825 of 845 pairs. In Standard it does not lower it in 13 of 17 (F-96). The bar still reads those pairs at 0.98 and 1.00, through `playset_share` in Modern and `unseen_share` in Standard.

A check on the synergy move there drops 867 of 1109 Modern synergy copies at 0.10, and the Modern synergy axis leaves the bar. It also moves three Modern decks the judge reads as bad one rung up:

| Deck | Judge | Check off | Check at 0.10 |
|---|---|---|---|
| 7. Modern burn, casual | bad | baseline | typical |
| 8. Modern lifegain, FNM | bad | bad | baseline |
| 24. a 60-card deck from one set | bad | baseline | typical |

Deck 8 leaves agreement with the judge, which is the point of agreement the rows above lose. So the check reads Commander alone. The owner chose that over every format (D-653).

## The three numbers, in gate run 18

Gate run 18 is the code of this item at the default floor, with the check in Commander alone.

| Number | Before the check | Gate run 18 |
|---|---|---|
| Commander synergy axis, precon over own copy | 0.70 of 389 | **0.84 of 232** |
| Commander precon over own copy, the bar | 0.88 of 945 | **0.95 of 788** |
| Built decks graded bad | 15 of 25 | **15 of 25** |
| Judge agreement | 8 of 25 | **8 of 25** |

Run 17 reports the first two numbers, and the control reads the last two over the same code.

The `bad` rung before and after the check:

| Format | Copies made | Dropped | Kept | Fold 0 train | Fold 0 holdout |
|---|---|---|---|---|---|
| commander | 6145 | 240 | 5905 | 4995 to 4804 | 1150 to 1115 |
| standard | 370 | 0 | 370 | 320 | 50 |
| modern | 2745 | 0 | 2745 | 2205 | 540 |

Commander drops them by the axis the fit asked for: copies 116, synergy 107, and colors 17. The Commander great-over-precon bar reads 0.961 against 0.963, and the accuracy stays at 0.65. Standard and Modern read as run 17 in every bar.

## What this says about the model

**The model did not improve, and the bar stopped counting pairs with no signal.** 37 of the 232 pairs still lose. The model can read the break of each of those, so they are the true gap of the synergy axis.

The item moves the bar, and it moves neither reader-facing number. A change to the population of the bar should give that result. PR-35 now measures against a bar that reads real defects (D-650).

## What this item does not measure

- **The pairs between 0.10 and 0.25.** The model reads them at 72 percent, and no audit asked whether a reader would feel their defect.
- **Whether a Modern synergy copy is a worse deck.** The bar reads those copies at 0.98 through the shape features, and F-96 names one reading of why that is unverified.

## The runs

- Gate run 18: `docs/reference/pr14b-quality-gate-run18.md`, with its run file in `docs/reference/eval/`.
- The control and the six floor runs wrote scratch files, and those files are not in the repository. The code of this item reads the check in Commander alone. So `QUALITY_GATE_ARGS="-synergy-floor=<floor>"` reproduces the Commander columns of each floor, and not the Modern column or the last three columns.
