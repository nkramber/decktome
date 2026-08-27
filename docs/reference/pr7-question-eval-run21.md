# PR-7 question eval

Run: `pr7-question-gate-run21`. Eval model: `gpt-5.6-luna`. Prompt version 3.

Scored 429 questions. 27 were not warranted, and 0 are unsure.

**Bad-question ratio: 10.8% on the holdout.**

The tune split holds 290 questions at 4.1%, and the holdout holds 139 at 10.8%. A fixer reads the tune split alone. A gain that shows only there is a reworded test set (D-134).

## The counters the ratio can not see

A run that asks less scores better and serves the user worse. Gate run 7 of 2026-08-25 passed both bars with 26 of 30 conversations unanswered.

| Counter | Value |
|---|---|
| Questions asked | 135 |
| Questions that closed a slot | 101 |
| Invented by the model | 1 |
| Catalog-only conversations | 29 |
| Premature sessions | 0 |
| Linter findings | 0 |

## Where the bad questions came from

| Row | Bad questions |
|---|---|
| `format_unsupported_open` | 3 |
| `meta` | 3 |
| `theme_card_named` | 3 |
| `budget` | 2 |
| `format` | 2 |
| `plan_choice` | 2 |
| `pool_precon` | 2 |
| `pool_thin` | 2 |
| `colors` | 1 |
| `commander` | 1 |
| `commander_pick` | 1 |
| `house_format_limits` | 1 |
| `house_rules` | 1 |
| `locked` | 1 |
| `pool` | 1 |
| `power_sixty` | 1 |

## Faults

| Fault | Count |
|---|---|
| assumes an answer | 17 |
| duplicate | 9 |
| inaccurate | 3 |
| out of scope | 2 |
| vague | 1 |

## What the catalog needs

| Action | Count |
|---|---|
| reword | 36 |
| remove | 7 |

## Every question the eval refused

The holdout conversations are left out of this list on purpose. 15 of their questions failed, and the fixer may not read which (D-134).

### 5. 60-card anything goes, turn 1, row `format`

**Asked:** Which format would you like: Commander, Standard, or Modern?

Faults: duplicate, assumes an answer. Catalog action: reword.

The user already said, “I want a 60-card deck, and anything goes at our table,” but the question forces a choice among Commander, Standard, and Modern instead of acknowledging that request.

### 5. 60-card anything goes, turn 1, row `house_rules`

**Asked:** Should I follow a no-ban-list rule, or did you mean something else?

Faults: duplicate. Catalog action: reword.

“Anything goes at our table” already answers the house-rules question by indicating that the user wants unrestricted card use.

### 11. another version after a build, turn 2, row `plan_choice`

**Asked:** Which of the two common plans do you want for the deck?

Faults: assumes an answer, vague. Catalog action: reword.

The question presumes that there are two established common plans, but the transcript names neither; reword it as: "What alternative game plan would you like this version to use?"

### 17. precon upgrade at bracket 2, turn 1, row `pool_precon`

**Asked:** Should I start with your Atraxa, Praetors' Voice precon, restrict the deck to its cards, or ignore it for a fully optimized build?

Faults: duplicate. Catalog action: reword.

The user already said “Upgrade my Atraxa, Praetors' Voice precon,” which establishes that the precon should be the starting point; reword the row to ask only about the degree of change, such as “Do you want to preserve the precon’s core strategy while upgrading it, or make substantial changes using it as the starting point?”

### 20. dinosaur tribal for a child, turn 1, row `commander_pick`

**Asked:** Which one do you want: Ghalta, Primal Hunger, Moon Girl and Devil Dinosaur, or Pantlaza, Sun-Favored? Say "none" and I name three more.

Faults: assumes an answer. Catalog action: reword.

The user only said "A dinosaur deck for my kid. Keep it casual" and had not said they wanted Commander; ask for the format before offering commanders.

### 20. dinosaur tribal for a child, turn 3, row `pool_thin`

**Asked:** You have four dinosaur cards but want at least 30. Should I start with the cards you own and add a buy list, or use all available cards?

Faults: assumes an answer. Catalog action: reword.

The question asserts "You have four dinosaur cards" and that the user wants "at least 30," although neither fact appears in the preceding turns; ask only whether to prioritize the collection or include a buy list.

### 28. reanimator with two plans, turn 4, row `plan_choice`

**Asked:** Which approach do you want for the deck?

Faults: duplicate. Catalog action: remove.

The user had already answered this by saying, “Reanimate one big creature, not a value loop,” so asking them to choose between “Reanimate a single powerful creature” and “Build a repeatable value loop” was redundant.

### 37. the user contradicts the budget (probe), turn 1, row `meta`

**Asked:** Should I keep the sideboard general, or would you like to name the decks you expect to face?

Faults: assumes an answer. Catalog action: reword.

The question presumes the user has an expected field of opposing decks, a fact not established at this point; reword as: “Would you prefer a general sideboard, or should I tailor it to a metagame you specify?”

### 56. the user answers a different question (probe), turn 1, row `budget`

**Asked:** Do you have a budget for cards to buy?

Faults: out of scope. Catalog action: remove.

The user explicitly requested a deck from their collection, so asking about a budget for cards to buy is outside the stated build scope.

### 61. terse: a proxy table (probe), turn 2, row `house_format_limits`

**Asked:** Inside your house format, do the normal 60-card deck limits hold?

Faults: assumes an answer. Catalog action: reword.

The user said “Any card, no ban list” but did not select one of the offered formats or establish that the deck would use a house format; reword as: “Since you did not choose one of those formats, should I use a custom 60-card format with your table’s normal deck-construction limits?”

### 62. terse: a precon upgrade (probe), turn 1, row `pool_precon`

**Asked:** Should I start with your Atraxa, Praetors' Voice precon, limit the deck to cards from it, or set it aside for a fully optimized build?

Faults: duplicate. Catalog action: reword.

The user already asked to “Upgrade my Atraxa, Praetors' Voice precon,” which establishes that the build should start from the precon; reword this row to ask only about any unusual card-pool restriction not already implied.

### 73. terse: a colorless deck (probe), turn 1, row `budget`

**Asked:** Do you have a budget for cards to buy?

Faults: assumes an answer. Catalog action: reword.

The user gave no budget, but the question presumes they have a budget. Reword as: "Should I prioritize budget-friendly cards, or may I include expensive staples?"

## Slots the eval says nobody asked about

These do not move the ratio. They are the guard against a quiet agent.

- 100. terse: a deck for a team event (probe): Ask which team-event rules apply, especially whether the event uses Unified/Team Constructed deckbuilding restrictions.
- 105. probe: a theme the library can not lead (probe): Ask which legal commander from the user's library should lead the colorless Eldrazi Commander deck, or whether the agent may choose one.
- 20. dinosaur tribal for a child: format
- 23. land destruction: Clarify the budget, because “My table does not mind” does not answer whether there is a spending limit for cards to buy.
- 26. legacy for an event (probe): After the user reaffirmed “Legacy,” ask whether they want to switch to one of the supported formats before proceeding.
- 30. another version, same plan: Ask which existing deck or decklist this should be a version of, since the transcript contains no original deck to modify., Ask which cards the user wants replaced or excluded, if “other cards” means substitutions from an existing list.
- 36. the user changes the theme late (probe): After the user changed the theme, clarify whether to retain Commander, black-green color identity, and Meren of Clan Nel Toth, and what token strategy is intended.
- 44. a card that is not legal (probe): Turn 2: Ask whether the user wants to remain in Modern and choose a legal replacement for Black Lotus, since Black Lotus is not legal in Modern and the user only said, "Then pick something else."
- 46. two partner commanders (probe): Whether Thrasios, Triton Hero and Tymna the Weaver should be used together as partner commanders
- 48. sideboard help only (probe): The agent did not ask for the current decklist or existing sideboard, which is needed to tailor sideboard recommendations to the user's actual build.
- 52. a price cap per card (probe): The user's total deck budget, since a per-card price cap does not establish a total budget.
- 56. the user answers a different question (probe): After “My friends hate mill,” ask whether the user still wants a mill deck for that playgroup or would prefer a different strategy.
- 6. the strongest deck, no collection: Follow up on the unanswered color-preference question, or explicitly confirm that the deck may use any colors., Clarify whether the $300 budget includes only the maindeck or also the sideboard, since the requested Modern deck includes both.
- 61. terse: a proxy table (probe): Clarify whether the user wants a custom Vintage-like 60-card format, rather than one of the listed supported formats, after they did not choose Commander, Standard, or Modern.
- 66. terse: sideboard help only (probe): Which specific control decks or colors do you expect to face?, Are there any cards already in your collection that you want to use in the sideboard?, What budget do you have for buying sideboard cards?, Do you have any sideboard preferences or constraints, such as avoiding narrow hate cards?
- 69. terse: a Pauper deck named once (probe): power level
- 7. FNM on Friday (probe): Ask for the user's color preference or whether they are open to any colors after the user did not answer the turn 1 colors question., Ask for the exact budget amount after the user said only “under budget.”
- 70. terse: a mill deck for a playgroup (probe): Clarify the conflict between the requested blue-and-black colors and Bruvac the Grandiloquent as commander, whose color identity permits only blue; ask whether to keep Bruvac or abandon black., Clarify what currency the stated budget of 100 uses, if the app requires a currency for card pricing.
- 76. terse: a companion (probe): Ask whether the user wants a currently Modern-legal deck despite naming Lurrus of the Dream-Den as a companion, since that request conflicts with tournament-level Modern legality.
- 79. terse: everything buried in one long message (probe): A question asking the user to provide an inventory or list of their relevant black and green cards would be needed to build primarily from their existing collection.
- 9. a commander the library does not hold: Format confirmation: whether the deck should be built for Commander., The user's available-library card pool, if the app cannot directly inspect the library.
- 95. terse: a library that does not exist (probe): Ask what cards are in the user's library or otherwise explain that a card collection is needed, since user_has_a_card_collection is false., Follow up on the unanswered commander choice: whether to suggest a commander or use one the user has in mind., Follow up on the unanswered budget question if cards may need to be bought.
- 96. terse: the user asks about a card (probe): Answer whether Sol Ring is legal in Modern: it is not legal in the Modern format., Confirm whether the user wants deck-building help rather than only a legality answer.
- 97. terse: Pauper Commander (probe): The user did not choose an alternative supported format after being told Pauper Commander was unsupported; ask them to select Commander, Standard, or Modern before proceeding., For a Commander build, ask which legendary creature or commander they want to use, unless the agent is meant to recommend one.

## Run

- Calls: 103. Time: 809.8 seconds.
- Tokens: 112545 input (0 cached), 66666 output.
- Cost: $0.1025.

## The card check

The check ran against the local snapshot. No refusal rested on a card name the snapshot refutes.
