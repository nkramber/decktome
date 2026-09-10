# The mana pass over /Users/nate/Repos/decktome/docs/reference/pr8-deck-gate-run18.md

Every deck of the document, rebuilt from its own shortlist. No provider call ran.

| # | Prompt | Off band before | Off band after | Worst color, before to after | Basics, before to after | Steps | What is left |
|---|---|---|---|---|---|---|---|
| 1 | lifegain Commander, any card | 0 | 0 | 1.17 to 1.17 | W 12, B 12 to W 12, B 12 | 0 | nothing |
| 2 | aristocrats Commander, owned first | 0 | 0 | 0.92 to 1.00 | W 18, B 18 to W 20, B 16 | 2 | nothing |
| 3 | artifacts Commander, bracket 4 | 1 | 1 | 1.31 to 1.31 | U 15 to U 15 | 0 | mana_turn_four |
| 4 | dinosaur tribal, bracket 2 | 0 | 0 | 1.16 to 1.16 | W 5, R 5, G 8 to W 5, R 5, G 8 | 0 | nothing |
| 5 | blink Commander, owned first | 0 | 0 | 1.56 to 1.56 | W 28 to W 28 | 0 | nothing |
| 6 | Modern tempo, tournament | 0 | 0 | 1.00 to 1.00 | U 10, R 6 to U 10, R 6 | 0 | nothing |
| 7 | Modern burn, casual | 0 | 0 | 1.71 to 1.71 | R 24 to R 24 | 0 | nothing |
| 8 | Modern lifegain, FNM | 0 | 0 | 1.22 to 1.22 | W 6, B 6 to W 6, B 6 | 0 | nothing |
| 9 | Standard midrange, FNM | 0 | 0 | 1.00 to 1.00 | B 6, G 6 to B 6, G 6 | 0 | nothing |
| 10 | Standard aggro, tournament | 0 | 0 | 1.25 to 1.25 | W 4, R 4 to W 4, R 4 | 0 | nothing |
| 11 | Commander with a locked card | 0 | 0 | 1.09 to 1.09 | W 14, B 12 to W 14, B 12 | 0 | nothing |
| 12 | Commander on a budget | 0 | 0 | 1.45 to 1.45 | W 28 to W 28 | 0 | nothing |
| 13 | owned first, and the commander is not owned | 0 | 0 | 1.18 to 1.18 | W 14, B 12 to W 14, B 12 | 0 | nothing |
| 14 | the user delegates the commander | 0 | 0 | 1.08 to 1.08 | R 13, G 12 to R 13, G 12 | 0 | nothing |
| 15 | delegated commander, owned first | 0 | 0 | 1.16 to 1.16 | W 17, B 12 to W 17, B 12 | 0 | nothing |
| 16 | a tight budget, owned first | 0 | 0 | 1.18 to 1.18 | W 19, B 18 to W 19, B 18 | 0 | nothing |
| 17 | upgrade a precon, owned first | 1 | 0 | 1.71 to 1.76 | R 25 to R 26 | 1 | nothing |
| 18 | upgrade a precon, any card | 7 | 7 | 1.10 to 1.10 | W 2, U 2, B 2, R 2, G 3 to W 2, U 2, B 2, R 2, G 3 | 0 | land, avg_mana_value, ramp, draw, removal, wipe, interaction |
| 19 | the Hobbit family, two colours | 1 | 1 | 0.66 to 0.82 | U 8, B 14, G 6 to U 7, B 13, G 10 | 4 | color_sources |
| 20 | the Hobbit family, a delegated commander | 1 | 1 | 1.91 to 1.91 | R 36 to R 36 | 0 | draw |
| 21 | the Hobbit family, mana from outside | 0 | 0 | 1.38 to 1.38 | B 9, R 10 to B 9, R 10 | 0 | nothing |
| 22 | a set family and a card from outside it | 1 | 0 | 0.78 to 0.87 | U 9, B 8, G 9 to U 9, B 7, G 11 | 3 | nothing |
| 23 | two set families at once | 0 | 0 | 1.66 to 1.66 | W 21 to W 21 | 0 | nothing |
| 24 | a 60-card deck from one set | 0 | 0 | 1.82 to 1.82 | R 24 to R 24 | 0 | nothing |
| 25 | use no card of an owned precon | 0 | 0 | 1.60 to 1.60 | B 21 to B 21 | 0 | nothing |

12 off-band features before the pass, 10 after. It moved 4 decks and improved 2.

The worst color's share of the sources it needs rose on 4 decks and fell on 0.

The sources of requirement of every deck under its need, or whose basics moved:

- 2. aristocrats Commander, owned first: before, sources of requirement: W 24 of 26, B 24 of 19. After, sources of requirement: W 26 of 26, B 22 of 19.
- 17. upgrade a precon, owned first: before, sources of requirement: R 32.5 of 19. After, sources of requirement: R 33.5 of 19.
- 19. the Hobbit family, two colours: before, sources of requirement: U 16.5 of 19, B 22.5 of 26, G 13.25 of 20. After, sources of requirement: U 15.5 of 19, B 21.5 of 26, G 17.25 of 20.
- 22. a set family and a card from outside it: before, sources of requirement: U 16.75 of 20, B 16.75 of 19, G 15.5 of 20. After, sources of requirement: U 17.5 of 20, B 16.5 of 19, G 18.25 of 20.

An upgrade keeps its own mana base, so the pass makes no step on one (D-249).
