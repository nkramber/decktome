# The mana pass of PR-33, over /Users/nate/Repos/decktome/docs/reference/pr8-deck-gate-run16.md

Every deck of the document, rebuilt from its own shortlist. No provider call ran.

| # | Prompt | Off band before | Off band after | Steps | What is left |
|---|---|---|---|---|---|
| 1 | lifegain Commander, any card | 0 | 0 | 0 | nothing |
| 2 | aristocrats Commander, owned first | 0 | 0 | 0 | nothing |
| 3 | artifacts Commander, bracket 4 | 2 | 1 | 9 | mana_turn_four |
| 4 | dinosaur tribal, bracket 2 | 0 | 0 | 0 | nothing |
| 5 | blink Commander, owned first | 0 | 0 | 0 | nothing |
| 6 | Modern tempo, tournament | 0 | 0 | 0 | nothing |
| 7 | Modern burn, casual | 0 | 0 | 0 | nothing |
| 8 | Modern lifegain, FNM | 0 | 0 | 0 | nothing |
| 9 | Standard midrange, FNM | 0 | 0 | 0 | nothing |
| 10 | Standard aggro, tournament | 0 | 0 | 0 | nothing |
| 11 | Commander with a locked card | 0 | 0 | 0 | nothing |
| 12 | Commander on a budget | 0 | 0 | 0 | nothing |
| 13 | owned first, and the commander is not owned | 0 | 0 | 0 | nothing |
| 14 | the user delegates the commander | 0 | 0 | 0 | nothing |
| 15 | delegated commander, owned first | 0 | 0 | 0 | nothing |
| 16 | a tight budget, owned first | 0 | 0 | 0 | nothing |
| 17 | upgrade a precon, owned first | 1 | 1 | 0 | removal |
| 18 | upgrade a precon, any card | 5 | 5 | 0 | land, ramp, draw, removal, interaction |
| 19 | the Hobbit family, two colours | 1 | 0 | 3 | nothing |
| 20 | the Hobbit family, a delegated commander | 0 | 0 | 0 | nothing |
| 21 | the Hobbit family, mana from outside | 0 | 0 | 0 | nothing |
| 22 | a set family and a card from outside it | 1 | 0 | 2 | nothing |
| 23 | two set families at once | 0 | 0 | 0 | nothing |
| 24 | a 60-card deck from one set | 0 | 0 | 0 | nothing |
| 25 | use no card of an owned precon | 0 | 0 | 0 | nothing |

10 off-band features before the pass, 7 after. It moved 3 decks and improved 3.

An upgrade keeps its own mana base, so the pass makes no step on one (D-249).

## The read

Prompts 17 and 18 are upgrades, and the pass skips both by design. They
hold 6 of the 10 off-band features. The pass ran on the other 4, and it
closed 3 of them.

What is left is `mana_turn_four` on prompt 3, an artifacts deck at
bracket 4. That band wants 4.8 mana on turn four from 31 to 36 lands,
and this pool reaches 4.6 after 9 steps. A band the pool can not reach
is a band no model call closes either, and the deck goes to the reader
with the note instead of a second call of 88 seconds.

The decks of this document were built before Part 1 and Part 2, so the
model that wrote them read five of the fifteen bands. The next whole
deck gate run measures the three parts together.
