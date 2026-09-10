# PR-30, the commander reference, 2026-09-10

The roadmap gate of PR-30 reads:

> Gate: the free quality gate with no axis under its PR-29 read, and the tier judge lane, $0.36. The count graded bad falls again, or the document says why not.

A correction of 2026-09-10 made the judge lane free: the gate reads the three numbers of D-648 against gate run 18.

**Verdict: dropped** (D-656). `commander_rate` fails in both forms a fit can give it. As built, it moves no built deck, and it drops the precon bar to 0.94. With every average deck read against itself, it raises the built decks graded bad from 15 to 20, and it cuts the judge agreement from 8 to 5.

Every number here is free. No run called a provider.

## The feature

`commander_rate` is the share of a deck's distinct nonland cards that the EDHREC average deck of its own commander holds (D-568). The fit builds the references from the EDHREC lists of the training split, and fold 0 reads 854 commanders.

An average deck matches itself. So the code of PR-30 reads no reference for an EDHREC list, and none for a broken copy of one. A deck with no reference reads the norm of the format.

**The plan held a false premise.** It said a deck with no average deck "reads the mean, as an absent value does". In the code of `main` an absent feature reads zero. PR-30 added the rule for its own feature alone.

## The three fits

Each fit reads the meta store of gate run 18 and the card snapshot of 2026-09-04, with the synergy check of D-653. The code of PR-30 sits on the local branch `pr30-commander-reference` at `a0b3de8`, and it never merged. A throwaway switch made the third fit.

| Fit | `commander_rate` weight | Synergy axis | Precon bar | Graded bad | Judge agreement |
|---|---|---|---|---|---|
| Gate run 18, no reference | - | 0.84 of 232 | 0.95, PASS | 15 of 25 | 8 of 25 |
| PR-30, no self-match | -0.351 | 0.81 of 232 | 0.94, FAIL | 15 of 25 | 8 of 25 |
| PR-30, every average deck reads itself | +0.191 | 0.91 of 232 | 0.97, PASS | 20 of 25 | 5 of 25 |

The PR-30 fit with no self-match grades every built deck as gate run 18 does.

## Why it fails

**As built, the fit reads the feature on every rung but typical.** Tournament lists overlap the average deck of their commander little, and a precon overlaps it much. A broken copy of a precon falls between. So the feature falls as the rungs climb, and the weight comes out negative. The typical rung is the one rung the feature exists to lift, and the self-match rule hides it.

**With the self-match, the typical rung reads an overlap near 1.** The fit learns that a deck is typical when it copies its average deck. No built deck does, so five built decks fall to bad:

| Deck | Judge | Gate run 18 | Every average deck reads itself |
|---|---|---|---|
| 2. aristocrats Commander, owned first | good | typical | bad |
| 5. blink Commander, owned first | typical | typical | bad |
| 18. upgrade a precon, any card | baseline | baseline | bad |
| 23. two set families at once | typical | typical | bad |
| 25. use no card of an owned precon | typical | baseline | bad |

**No honest typical row exists in the store.** Each commander holds one average deck, so the reference of a typical list is that list. An honest row needs a second deck of the same commander that is not the average, such as one community deck. The store holds none.

## The judge and the model read one ladder

The tier judge reads the ladder of D-414 (`go/internal/generate/judge.go`). It names typical as "the average deck of its commander or archetype as the community builds it". It reads card quality against the best lists of the format. So the gap between the judge and the model is no gap of frame. The model can not place a built deck on the typical rung.

## What it means

Three items aimed at that gap moved no reader-facing number, or moved it the wrong way: PR-29, M-9, and PR-30 (F-94, F-98, F-99). Each one added a feature to one ordinal ladder. That ladder gives every feature one direction up the rungs, so a feature that marks the middle rung alone has no weight that works.

The owner dropped PR-30 and chose a structural change to the model (D-656). M-10 measures a scorer with one weight set per tier first.
