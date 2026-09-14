# PR-45b: the mana pass lane for F-133

Run date: 2026-09-14. Fixed step: commit `8ca9e58`. Old step: commit `f45791a`. Decks: deck gate run 19. Card snapshot: `20260904T210157`. Quality model: `20260914T154223Z`.

## What this reads

`make manapass-check` rebuilds each deck of a gate document from its own shortlist and runs the mana pass, with no model call. The lane ran twice over the 25 decks of deck gate run 19: once on the old step, and once on the step that keeps the job (F-133, D-713). Both runs read the shortlist of PR-45b.

## Read

- **24 of 25 decks read the same on both steps.** No deck at brackets 1 to 3 and no 60-card deck moves.
- **Deck 3 keeps one finding on the fixed step.** The artifacts deck at bracket 4 closed both of its off-band features on the old step. On the fixed step it keeps the ramp finding. Both steps moved 3 cards, and the lane does not name the cards each step added.
- **The lane reads 9 off-band features after the pass on the fixed step, against 8 on the old step.** Both steps start from 14.

The report of the fixed step follows. The report of the old step differs in the row of deck 3 and in the last line alone.

## The fixed step over `docs/reference/pr8-deck-gate-run19.md`

Every deck of the document, rebuilt from its own shortlist. No provider call ran.

| # | Prompt | Off band before | Off band after | Worst color, before to after | Basics, before to after | Steps | What is left |
|---|---|---|---|---|---|---|---|
| 1 | lifegain Commander, any card | 1 | 0 | 0.97 to 0.99 | W 11, B 16 to W 10, B 17 | 2 | nothing |
| 2 | aristocrats Commander, owned first | 0 | 0 | 1.29 to 1.29 | W 16, B 16 to W 16, B 16 | 0 | nothing |
| 3 | artifacts Commander, bracket 4 | 2 | 1 | 1.31 to 1.29 | U 26 to U 24 | 3 | ramp |
| 4 | dinosaur tribal, bracket 2 | 2 | 0 | 1.02 to 1.03 | W 9, R 5, G 7 to W 9, R 5, G 7 | 4 | nothing |
| 5 | blink Commander, owned first | 0 | 0 | 1.76 to 1.76 | W 28 to W 28 | 0 | nothing |
| 6 | Modern tempo, tournament | 0 | 0 | 1.43 to 1.43 | U 4, R 4 to U 4, R 4 | 0 | nothing |
| 7 | Modern burn, casual | 0 | 0 | 1.14 to 1.14 | R 20 to R 20 | 0 | nothing |
| 8 | Modern lifegain, FNM | 0 | 0 | 1.00 to 1.00 | W 8, B 8 to W 8, B 8 | 0 | nothing |
| 9 | Standard midrange, FNM | 0 | 0 | 1.06 to 1.06 | B 6, G 6 to B 6, G 6 | 0 | nothing |
| 10 | Standard aggro, tournament | 0 | 0 | 0.94 to 0.94 | W 9, R 11 to W 9, R 11 | 0 | nothing |
| 11 | Commander with a locked card | 0 | 0 | 1.26 to 1.26 | W 10, B 9 to W 10, B 9 | 0 | nothing |
| 12 | Commander on a budget | 1 | 0 | 1.41 to 1.41 | W 28 to W 28 | 1 | nothing |
| 13 | owned first, and the commander is not owned | 0 | 0 | 1.14 to 1.14 | W 18, B 12 to W 18, B 12 | 0 | nothing |
| 14 | the user delegates the commander | 0 | 0 | 1.20 to 1.20 | R 10, G 10 to R 10, G 10 | 0 | nothing |
| 15 | delegated commander, owned first | 0 | 0 | 1.16 to 1.16 | W 15, B 14 to W 15, B 14 | 0 | nothing |
| 16 | a tight budget, owned first | 0 | 0 | 1.03 to 1.03 | W 19, B 18 to W 19, B 18 | 0 | nothing |
| 17 | upgrade a precon, owned first | 1 | 1 | 1.87 to 1.87 | R 28 to R 28 | 0 | interaction |
| 18 | upgrade a precon, any card | 2 | 2 | 0.94 to 0.94 | W 7, U 3, B 4, R 4, G 5 to W 7, U 3, B 4, R 4, G 5 | 0 | ramp, draw |
| 19 | the Hobbit family, two colours | 0 | 0 | 0.88 to 0.88 | U 9, B 9, G 12 to U 9, B 9, G 12 | 0 | nothing |
| 20 | the Hobbit family, a delegated commander | 1 | 1 | 2.03 to 2.03 | R 36 to R 36 | 0 | draw |
| 21 | the Hobbit family, mana from outside | 0 | 0 | 1.19 to 1.19 | B 5, R 6 to B 5, R 6 | 0 | nothing |
| 22 | a set family and a card from outside it | 1 | 1 | 0.76 to 0.76 | U 8, B 13, G 10 to U 8, B 13, G 10 | 0 | color_sources |
| 23 | two set families at once | 0 | 0 | 1.73 to 1.73 | W 21 to W 21 | 0 | nothing |
| 24 | a 60-card deck from one set | 0 | 0 | 1.82 to 1.82 | R 24 to R 24 | 0 | nothing |
| 25 | use no card of an owned precon | 3 | 3 | 1.84 to 1.84 | B 36 to B 36 | 0 | ramp, draw, removal |

14 off-band features before the pass, 9 after. It moved 4 decks and improved 4.

The worst color's share of the sources it needs rose on 2 decks and fell on 1.

The sources of requirement of every deck under its need, or whose basics moved:

- 1. lifegain Commander, any card: before, sources of requirement: W 23 of 23, B 27.25 of 28. After, sources of requirement: W 22.75 of 23, B 28.25 of 28.
- 3. artifacts Commander, bracket 4: before, sources of requirement: U 34 of 26. After, sources of requirement: U 33.5 of 26.
- 10. Standard aggro, tournament: before, sources of requirement: W 13 of 13, R 15 of 16. After, sources of requirement: W 13 of 13, R 15 of 16.
- 18. upgrade a precon, any card: before, sources of requirement: W 19 of 20, U 17 of 18, B 18 of 19, R 18 of 19, G 22.25 of 23. After, sources of requirement: W 19 of 20, U 17 of 18, B 18 of 19, R 18 of 19, G 22.25 of 23.
- 19. the Hobbit family, two colours: before, sources of requirement: U 16.75 of 19, B 16.75 of 19, G 17.5 of 19. After, sources of requirement: U 16.75 of 19, B 16.75 of 19, G 17.5 of 19.
- 22. a set family and a card from outside it: before, sources of requirement: U 14.75 of 19, B 20.75 of 26, G 15.25 of 20. After, sources of requirement: U 14.75 of 19, B 20.75 of 26, G 15.25 of 20.

An upgrade keeps its own mana base, so the pass makes no step on one (D-249).
