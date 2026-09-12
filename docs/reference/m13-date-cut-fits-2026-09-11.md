# M-13, the date cut of PR-41, 2026-09-11

The roadmap gate of M-13 reads:

> Gate: the control reproduces gate run 20 in every number. A cut passes when one reader number improves, the other holds or improves, and no bar falls (guardrail 15, D-681). The owner reads the fits, and then PR-41 builds the cut or closes.

**Verdict: no cut meets the gate.** The control reproduces gate run 20 in every number. No cut moves a built deck to another tier. So every fit grades 9 built decks bad, with 10 in agreement and 5 owner matches. Every bar still passes, but each cut makes one number a little worse.

Every number here is free. No run called a provider.

## The premise (F-118)

F-108 says the top lists hold few cards of The Hobbit, because the set released on 2026-08-14. A count over the local store checked that cause before the fits. The count reads the lists the fit keeps: the newest 4,000 great Commander lists, from 2026-07-04, and the newest 4,000 good lists, from 2026-08-16. The newest list is dated 2026-09-06. The count resolves each card by name, reads the first paper set of the card, and skips the holdout fold.

| Set code | Released | Nonland cards first printed in the set | In 1 list or more | In 5 lists or more |
|---|---|---|---|---|
| mkm | 2024-02-09 | 255 | 23.1% | 8.2% |
| otj | 2024-04-19 | 244 | 29.9% | 13.1% |
| big | 2024-04-19 | 23 | 65.2% | 39.1% |
| blb | 2024-08-02 | 250 | 50.8% | 16.8% |
| dsk | 2024-09-27 | 249 | 38.2% | 15.7% |
| dft | 2025-02-14 | 242 | 34.3% | 8.7% |
| tdm | 2025-04-11 | 246 | 30.5% | 9.3% |
| fin | 2025-06-13 | 278 | 59.7% | 18.7% |
| eoe | 2025-08-01 | 245 | 40.4% | 13.9% |
| spm | 2025-09-26 | 175 | 53.1% | 15.4% |
| tla | 2025-11-21 | 252 | 45.2% | 16.7% |
| ecl | 2026-01-23 | 251 | 42.6% | 15.5% |
| tmt | 2026-03-06 | 178 | 52.8% | 14.0% |
| sos | 2026-04-24 | 247 | 45.7% | 17.8% |
| msh | 2026-06-26 | 248 | 50.0% | 15.3% |
| hob | 2026-08-14 | 177 | 54.2% | 20.9% |

**The Hobbit ranks third of these 16 expansions on the first share, and second on the second.** So the age of the set does not explain F-108.

The built decks of deck gate run 16 read their unseen nonland cards against the same lists:

| Deck | Nonland cards | Cards of The Hobbit sets | Unseen | Unseen of The Hobbit sets | Unseen of other sets |
|---|---|---|---|---|---|
| 19 | 64 | 33 | 36% | 19 | 4 |
| 20 | 64 | 39 | 17% | 10 | 1 |
| 21 | 64 | 39 | 20% | 11 | 2 |
| 22 | 64 | 33 | 39% | 17 | 8 |
| 23 | 64 | 12 | 12% | 3 | 5 |

The 14 other Commander decks hold two cards of The Hobbit sets at most. They read 0 to 32 percent unseen, with a median of 18. Decks 20 and 21 sit at that median. Decks 19 and 22 read higher, and the colors check grades both bad at 0.71 and 0.70 (D-678).

## The runs

Each fit is the code of `3bb9cd1`, over the meta store of gate run 20 and the card snapshot of 2026-09-04. One throwaway patch makes the cut, and it does not merge. With no cut, the patched code reads every card, so the control runs on the same build. The fits ran from 00:38 to 00:45 UTC on 2026-09-12, at about 107 seconds each.

A card is new when its first paper release falls within the cut of the newest great or good list of the format. The cut removes each new card from `card_rate`, `unseen_share`, and `synergy`, in the fit and in the grade. Every other feature reads every card.

- **control**: no cut.
- **cut30-cmd**: Commander alone. The cut is 2026-08-07, 30 days before the newest list of 2026-09-06, and it reads 401 cards as new.
- **cut90-cmd**: Commander alone. The cut is 2026-06-08, and it reads 1,039 cards as new. Marvel Super Heroes falls inside it.
- **cut30-all**: every format, as the text of PR-41 reads. Commander reads as cut30-cmd. Standard and Modern cut at 2026-08-09, 30 days before their newest list of 2026-09-08.

A count over the snapshot splits the 401 cards of the 30-day cut. 202 come from The Hobbit and The Hobbit Eternal. 197 come from three sets that release after the snapshot: Reality Fracture, Mystery Booster Commander Edition, and Star Trek. The count reads 399 cards against 401, because it filters the printings a little differently from the card index.

## The fits

| Fit | Precon bar | Synergy axis | Great over precon | Graded bad | Agreement | Owner | Commander accuracy | Broken copies graded bad | Precons graded right | Real lists graded bad |
|---|---|---|---|---|---|---|---|---|---|---|
| control | 752 of 788 | 196 of 232 | 0.96 | 9 | 10 | 5 | 0.55 | 4,141 of 5,905 | 73 of 189 | 615 |
| cut30-cmd | 751 of 788 | 195 of 232 | 0.96 | 9 | 10 | 5 | 0.55 | 4,145 of 5,908 | 72 of 189 | 617 |
| cut90-cmd | 752 of 790 | 196 of 234 | 0.96 | 9 | 10 | 5 | 0.55 | 4,158 of 5,917 | 66 of 189 | 618 |
| cut30-all | 751 of 788 | 195 of 232 | 0.96 | 9 | 10 | 5 | 0.55 | 4,145 of 5,908 | 72 of 189 | 617 |

The lands, curve, and colors axes read 189, 189, and 178 in every fit. The synergy check drops 240 copies in the control, 237 at 30 days, and 228 at 90 days. Standard reads 1.00 of 40 and Modern 0.99 of 1,635 on the precon bar in every fit. In cut30-all, the Standard great-over-precon bar reads 0.91 against 0.92, over its bar of 0.90.

**No built deck changes its tier in any fit.** The Commander decks graded bad are 1, 12, 14, 16, 19, and 22 in every fit. Decks 20 and 21 stay baseline, and deck 23 stays typical. The gate document prints no score for a built deck, so this record does not say how far deck 21 moved.

**What the fits say.**

- No cut moves a reader number, so no cut meets guardrail 15.
- Each cut costs a little. The cut of 30 days loses one synergy pair of the precon bar. The cut of 90 days grades 7 fewer precons right. The cut in every format lowers the Standard great-over-precon bar by 0.01.
- The fits do not measure the first weeks after a release. The good lists of the fit cover about three weeks. So in the first days after a release, few of those lists come after the release. No built deck of the gate comes from a set that new.

## The owner decision (D-683)

The owner closed PR-41 on this evidence. No cut moves a reader number, and F-118 refutes the premise of F-108. The owner chose a close over a park until a set is days old at a fit. The owner also chose no new item for the reason text of a deck limited to one set.

The first days after a release stay unmeasured. The deployed meta job refits the model every day. So for a set a few days old, most good lists of the fit come before its release.
