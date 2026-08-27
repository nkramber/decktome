# PR-7 question eval

Run: `pr7-question-gate-run20`. Eval model: `gpt-5.6-luna`. Prompt version 3.

Scored 430 questions. 18 were not warranted, and 0 are unsure.

**Bad-question ratio: 9.3% on the holdout.**

The tune split holds 290 questions at 1.7%, and the holdout holds 140 at 9.3%. A fixer reads the tune split alone. A gain that shows only there is a reworded test set (D-134).

## The counters the ratio can not see

A run that asks less scores better and serves the user worse. Gate run 7 of 2026-08-25 passed both bars with 26 of 30 conversations unanswered.

| Counter | Value |
|---|---|
| Questions asked | 137 |
| Questions that closed a slot | 107 |
| Invented by the model | 1 |
| Catalog-only conversations | 29 |
| Premature sessions | 0 |
| Linter findings | 0 |

## Where the bad questions came from

| Row | Bad questions |
|---|---|
| `colors` | 2 |
| `format_unsupported_open` | 2 |
| `pool_thin` | 2 |
| `power_commander` | 2 |
| `power_sixty` | 2 |
| `budget` | 1 |
| `commander` | 1 |
| `commander_pick` | 1 |
| `format` | 1 |
| `locked` | 1 |
| `pool` | 1 |
| `pool_precon` | 1 |
| `theme_card_named` | 1 |

## Faults

| Fault | Count |
|---|---|
| out of scope | 5 |
| assumes an answer | 4 |
| inaccurate | 4 |
| duplicate | 2 |

## What the catalog needs

| Action | Count |
|---|---|
| reword | 24 |
| remove | 9 |
| no | 2 |

## Every question the eval refused

The holdout conversations are left out of this list on purpose. 13 of their questions failed, and the fixer may not read which (D-134).

### 10. no strong commander in the library, turn 1, row `budget`

**Asked:** Do you have a budget for cards to buy?

Faults: assumes an answer. Catalog action: no.

The question presumes that the user may need to buy cards and has a budget, although the user had only said “Lifegain from my collection.” It should instead ask whether purchases are allowed, followed by a budget question only if needed.

### 20. dinosaur tribal for a child, turn 3, row `pool_thin`

**Asked:** You have four dinosaur cards but need at least 30. Should I prioritize the cards you own and add a buy list, or use all available cards?

Faults: assumes an answer, inaccurate. Catalog action: reword.

The transcript never says the user has exactly “four dinosaur cards,” and a Commander deck does not generally “need at least 30” dinosaur cards. Reword as: “Should I use cards from your library first and list any missing cards to buy, or build from the full available card pool?”

### 41. a request we can not serve (probe), turn 3, row `commander_pick`

**Asked:** Which one do you want: Atarka, World Render, Thrakkus the Butcher, or Dragonlord Atarka? Say "none" and I name three more.

Faults: duplicate. Catalog action: reword.

The user had not answered the commander question, but this repeats the earlier commander-selection row rather than explicitly asking whether they want to choose from the original list or see alternatives; reword it as: “Which commander do you want, or should I suggest three more?”

### 49. copy a known deck (probe), turn 1, row `colors`

**Asked:** Do you have a color preference?

Faults: assumes an answer. Catalog action: reword.

The user said “whatever is winning,” which indicates no color preference; reword as “If you have a color preference, what colors should I consider?”

### 62. terse: a precon upgrade (probe), turn 1, row `pool_precon`

**Asked:** Should I upgrade the precon, restrict the deck to cards from it, or build without using it?

Catalog action: reword.

The user already said “Upgrade my Atraxa, Praetors' Voice precon,” so asking whether to start with the precon, use only its cards, or ignore it was unnecessary. Proposed wording: “I’ll upgrade the precon; should I keep the deck limited to its original cards, or may I add cards from your collection and new purchases?”

## Slots the eval says nobody asked about

These do not move the ratio. They are the guard against a quiet agent.

- 18. pauper burn on a budget (probe): power level
- 19. a big library, no theme: Whether the deck must be built entirely from the user's collection or whether cards may be suggested for purchase after using the collection first.
- 30. another version, same plan: Ask which existing deck or deck list the user wants to reproduce, since “the same deck” and “the same plan” refer to an unspecified prior deck.
- 36. the user changes the theme late (probe): After the user changed the theme, ask whether the deck should remain Commander with Meren and black-green identity, and clarify what kind of token strategy they want.
- 48. sideboard help only (probe): Ask for the existing decklist and current sideboard, or at least the cards and slots that need changing, so the recommendations can fit the user's actual Modern burn deck.
- 52. a price cap per card (probe): A question asking for the total deck budget, since the user gave a per-card price cap but not an overall budget.
- 55. a returning player (probe): A follow-up confirming whether to suggest a commander, since the user did not answer the commander question and only said “Green.”
- 56. the user answers a different question (probe): Whether the user still wants a mill deck despite saying their friends hate mill, or would prefer a different strategy for their playgroup.
- 6. the strongest deck, no collection: Whether the $300 budget includes the sideboard and any other required cards or accessories.
- 63. terse: a competitive request (probe): Clarify what budget limit the user means by “under budget,” since this conflicts with “money is no object.”
- 71. terse: one-word answers (probe): Clarify which control decks or archetypes the user expects to face, since “Control” does not identify specific matchups.
- 74. terse: a card the user does not own (probe): Clarify whether the original “only the cards I own” restriction still applies after the user says “Buy it then.”
- 76. terse: a companion (probe): The agent should have addressed that Lurrus of the Dream-Den cannot currently be used as a Modern companion, asking whether the user means a historical Modern deck or wants to choose a legal companion.
- 80. terse: all caps, no punctuation (probe): Ask which specific control decks the shop plays so the sideboard can be tuned to the actual expected matchups.
- 95. terse: a library that does not exist (probe): The user's library or card-collection list/export, since no cards were supplied and user_has_a_card_collection is false.
- 96. terse: the user asks about a card (probe): Answer whether Sol Ring is legal in Modern before proceeding with deck-building questions.

## Run

- Calls: 102. Time: 656.5 seconds.
- Tokens: 111414 input (0 cached), 62508 output.
- Cost: $0.0973.

## The card check

The check ran against the local snapshot. No refusal rested on a card name the snapshot refutes.
