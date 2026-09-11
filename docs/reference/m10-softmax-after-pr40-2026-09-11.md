# M-10 again, the softmax scorer after PR-40, 2026-09-11

The roadmap gate of M-10 reads:

> Gate: the three numbers of D-648 beside gate run 18. If graded-bad falls or the judge agreement rises, and no bar falls, PR-38 builds the scorer. If not, the owner decides again (D-656).

D-665 set this measurement for after the detector fix, and PR-40 merged as #144. So this run reads the gate beside gate run 20.

**Verdict: not met.** The Commander precon bar falls from 752 to 748 of 788, one pair under the 749 the bar asks. The judge agreement rises from 10 to 12, and the built decks graded bad rise from 9 to 14. The owner's grades match on 2 decks, against 5.

Every number here is free. No run called a provider.

## The runs

Each fit is the code of `38ddde6`, which carries PR-40, over the meta store of gate run 20 and the card snapshot of 2026-09-04. A throwaway patch adds the softmax ladder of M-10. It fits one weight set and one bias per tier over the same standardized features, every rung at the same weight (D-488). It keeps the iterations, the rate, and the L2 weight of the ordinal fit.

The grade, the Commander tier of PR-40, and the explain ladder read the softmax probabilities. The reasons keep the ordinal weights. The control, with the patch off, reproduces gate run 20 in every number.

## The fits

| Number | Control | Softmax |
|---|---|---|
| Commander precon over own copy | 752 of 788 (189, 189, 178, 196) | 748 of 788 (186, 189, 172, 201) |
| Commander great over precon | 0.96 | 0.97 |
| Standard precon over own copy | 1.00 of 40 | 0.97 of 40 |
| Modern precon over own copy | 0.99 of 1635 | 0.99 of 1635 |
| Commander accuracy | 0.55 | 0.63 |
| Standard accuracy | 0.45 | 0.49 |
| Modern accuracy | 0.55 | 0.61 |
| Built decks graded bad | 9 of 25 | 14 of 25 |
| Judge agreement | 10 of 25 | 12 of 25 |
| Owner matches | 5 of 10, mean error 0.8 | 2 of 10, mean error 1.2 |
| Broken copies graded bad | 4,141 of 5,905 | 4,937 of 5,905 |
| Precons graded bad | 58 of 189 | 21 of 189 |
| Average decks graded bad | 71 of 1,040 | 107 of 1,040 |
| Good lists graded bad | 247 of 4,000 | 302 of 4,000 |
| Great lists graded bad | 239 of 4,000 | 278 of 4,000 |

The only bar that fails is the Commander precon bar. The lands axis loses 3 pairs and the colors axis 6, and the synergy axis gains 5.

## The built decks

- **Decks 3, 4, 15, and 21 fall from baseline to bad.** The owner grades decks 3, 15, and 21 baseline, and deck 4 typical.
- **Decks 11, 13, 20, and 25 rise from baseline to typical.** The owner grades deck 11 typical and deck 13 baseline. The judge reads decks 11, 13, and 25 as typical, and deck 20 as baseline.
- **Deck 10 of Standard rises to typical, and deck 24 of Modern falls to bad.** The judge reads deck 10 as baseline and deck 24 as bad.

So the agreement gains decks 11, 13, 24, and 25, and it loses decks 10 and 20.

## What the numbers say

- **The scorer grades the holdout better.** It holds more broken copies at bad and fewer precons there, so the accuracy rises in every format.
- **It grades the built decks worse by the owner's grades.** Three decks the owner grades baseline fall to bad, and the matches fall from 5 to 2.
- **The Commander precon bar fails by one pair,** as it failed on 2026-09-10 at 748.

## A gap in the gate

The gate passes an item when graded bad falls or the agreement rises, and no bar falls. Every bar reads the score. So a variant that lets the softmax set the tier and keeps the ordinal score holds every bar by construction. Its tiers match this fit by construction. So it passes on an agreement of 12, while it grades 14 built decks bad and matches 2 owner grades. No fit measured that variant.

## The decisions (D-680, D-681)

**The owner closed PR-38** (D-680). The scorer failed the Commander precon bar twice, and after PR-40 it grades the built decks further from the owner's grades.

**A quality item now passes only when neither reader number gets worse** (D-681). An item passes when one of the two numbers improves, the other holds or improves, and no bar falls. Under that rule the split variant above fails, because its graded bad rises from 9 to 14.
