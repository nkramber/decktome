# M-8, the audit of the precon bar, 2026-09-10

The roadmap gate of M-8 reads:

> The document names, for every failing pair, the precon, its copy, and the three reads. It says whether the label holds.

**Verdict: the label holds, and no model meets the bar.** The broken copy is not the better deck. The break is real. On 54 of the 389 Commander pairs it moves nothing a model can read, because those precons had no synergy to break.

Every number here is free. The audit runs inside the fit, and `-audit-out` writes every own-copy pair as JSON.

## The question

The bar asks the ladder to score a precon above its own broken copy. The synergy break replaces half a precon's spells with cards the real lists play (D-488). PR-29 asked whether that leaves a weak precon's copy as the better deck. The bar then asks for a false ordering (D-649).

## The answer, in three reads

The reads come from the 389 Commander own-copy pairs of the synergy axis, on the code of `main`. The bar reads 0.70 of them.

**The failures are the weak precons.** Of the 118 pairs the bar loses, 102 sit in the weaker half of the 44 holdout precons, which is 86 percent. Of the 271 it wins, 126 sit there, which is 46 percent and the share chance gives. The median rank of a loser is 10 and of a winner 23.

**The copy is not the better deck.** Of the 118 lost pairs, the copy holds cards the top lists play more in 59, which is exactly half. It pairs them better in 19, which is 16 percent. So the break moves the synergy feature the right way in 84 percent of the pairs the bar loses. The label holds.

**The break reads as nothing at the floor.** The mean signed move of `synergy` is the copy less the precon. It reads -1.643 over the precons that score above 0.30. It reads **-0.072** over the precons that score at or under 0.05, which is 23 times less.

## Why

A weak precon's cards do not pair in the corpus to begin with. Its `synergy` feature already reads near nothing, so a break that removes half its pairs removes half of nothing.

The bands say it plainly. Each row reads every own-copy pair of that band, and not the losses alone.

| Precon score | Pairs | Share won | Mean `synergy` move | Precon and copy grade apart |
|---|---|---|---|---|
| at or under 0.05 | 54 | **0.43** | -0.072 | **0 percent** |
| 0.05 to 0.30 | 135 | 0.58 | - | 2 percent |
| above 0.30 | 200 | 0.85 | -1.643 | 58 percent |

At the floor the precon and its copy grade the same tier in every pair, and the median score gap is 0.0133. The bar there is a tie the model breaks by chance, and it reads 0.43, which is under a coin flip.

The bar climbs with the room the precon has to fall:

| The bar over precons scoring above | Pairs | Share won |
|---|---|---|
| 0.00, every pair | 389 | 0.70 |
| 0.05 | 335 | 0.74 |
| 0.10 | 262 | 0.79 |
| 0.30 | 200 | 0.85 |

## What this means for the bar of 0.95

54 of 389 pairs, 14 percent, carry a break the model can not read. A model that wins every one of the other 335 and splits the 54 by chance reads **0.93**. So **0.95 is out of reach while those pairs are in the bar**, whatever the model does.

That is not a reason to move the bar (D-486). It is a reason to look at the break.

## The repository already holds the rule this breaks

The title of D-485 holds the phrase "a break must be material", and its words are: "A break must be material, or it is no defect and the label lies."

It gives two materiality checks. A copies break needs 12 copies in playsets, and a colors break four fixing lands. A list that fails either one **breaks on synergy instead**.

So synergy is the fallback, and it is the one break with no materiality check of its own. Every list that fails the other two lands there, and so does every list too weak for synergy to mean anything. F-95 records it.

## What M-8 recommends, for the owner to decide

**A materiality check on the synergy break**, in the shape D-485 already uses for the other two. A list whose cards hold too few lifting pairs has no synergy to break. A break on it labels a deck that nothing broke.

Three questions come with it, and none is the session's to answer.

- What is "too few"? The other two checks name a count: 12 copies, four lands. The natural one here is the count of the deck's card pairs that lift at all in the corpus.
- What happens to a list that fails every check? Today synergy catches them. With a check on synergy they fall out of the synthetic set, and the bar reads fewer pairs.
- Does the bar of 0.95 stand once the unbreakable pairs leave it? The bar reads 0.74 over the pairs above the floor today, and 0.85 over the pairs with real room. Neither is 0.95.

## What this says about PR-29

PR-29 lifted the axis from 0.70 to 0.81 and both reader-facing numbers fell. This audit does not change either number. It changes what the 0.95 target means: no corpus and no fit reaches it while a seventh of the pairs carry no signal.

So the case for PR-29 rests on its own reader-facing numbers, and those read worse. The owner holds it on that basis, and this audit gives no reason to change that.

## The artifact

`-audit-out` writes every own-copy pair as JSON. Each row holds the precon and the axis. It holds both scores, both grades, the precon's rank among the holdout precons, and the signed move of three corpus features. The run of 2026-09-10 wrote 2620 pairs over the three formats.

The gate document gains one table per format, and it reads the mean rank of the winners and the losers side by side. A break that hurts every precon alike moves the two together.
