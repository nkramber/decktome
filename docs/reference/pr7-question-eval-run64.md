# PR-7 question eval

Run: `pr7-question-gate-run64`. Eval model: `gpt-6-luna`. Prompt version 3.

Scored 431 questions. 20 were not warranted, and 0 are unsure.

**Bad-question ratio: 8.3% on the holdout.**

The tune split holds 299 questions at 3.0%, and the holdout holds 132 at 8.3%. A fixer reads the tune split alone. A gain that shows only there is a reworded test set (D-134).

## The counters the ratio can not see

A run that asks less scores better and serves the user worse. A run can pass both bars with most of its conversations unanswered, so read these counters with the ratio.

| Counter | Value |
|---|---|
| Questions asked | 295 |
| Questions that closed a slot | 235 |
| Invented by the model | 1 |
| Catalog-only conversations | 74 |
| Premature sessions | 0 |
| Linter findings | 0 |

## Where the bad questions came from

| Row | Bad questions |
|---|---|
| `budget` | 3 |
| `theme` | 3 |
| `colors` | 2 |
| `commander` | 2 |
| `commander_pick` | 2 |
| `commander_excludes_card` | 1 |
| `format` | 1 |
| `format_unsupported_open` | 1 |
| `named_card_role` | 1 |
| `pool` | 1 |
| `power_commander` | 1 |
| `power_sixty` | 1 |
| `theme_unmatched` | 1 |

## Faults

| Fault | Count |
|---|---|
| assumes an answer | 5 |
| out of scope | 5 |
| duplicate | 4 |
| inaccurate | 3 |
| omits information | 2 |

## What the catalog needs

| Action | Count |
|---|---|
| reword | 20 |
| remove | 2 |
| add | 1 |

## Every question the eval refused

The holdout conversations are left out of this list on purpose. 11 of their questions failed, and the fixer may not read which (D-134).

### 5. 60-card anything goes, turn 1, row `format`

**Asked:** Which format would you like: Standard or Modern?

Faults: omits information. Catalog action: reword.

The user already said “anything goes at our table,” so asking them to choose only Standard or Modern misses their stated option; reword to ask whether they want a named format or their anything-goes table rules.

### 7. FNM on Friday (probe), turn 2, row `colors`

**Asked:** Do you have a color preference?

Faults: duplicate. Catalog action: remove.

The agent already asked “Do you have a color preference?” on turn 1, and the user did not answer it; repeating the same question is still a duplicate.

### 35. the user changes the format (probe), turn 3, row `budget`

**Asked:** What spending limit, if any, should I use for cards you don’t own?

Faults: duplicate. Catalog action: none.

The agent had already asked the same spending-limit question on turn 1; the format change did not make that unanswered question new.

### 44. a card that is not legal (probe), turn 1, row `theme`

**Asked:** What should the deck focus on: a creature type, mechanic, or play style?

Catalog action: reword.

The user already said they want the deck built around “Black Lotus”; reword this to acknowledge that focus rather than ask what the deck should focus on.

### 53. the user swaps the commander (probe), turn 4, row `commander_pick`

**Asked:** Which commander do you want: Astarion, the Decadent, Karlov of the Ghost Council, or Sorin of House Markov // Sorin, Ravenous Neonate? Say "none" and I name three more.

Faults: out of scope. Catalog action: reword.

The user asked the agent to “suggest one,” so asking them to choose among commanders defers the requested recommendation; instead, recommend a commander and briefly explain why.

### 71. terse: one-word answers, turn 1, row `budget`

**Asked:** What spending limit, if any, should I use for cards you don’t own?

Faults: assumes an answer. Catalog action: reword.

The question presumes a spending limit; reword it as “Do you have a spending limit for cards you don’t own?”

### 77. terse: a Background pair, turn 3, row `commander_pick`

**Asked:** Which commander do you want: Aurelia, the Warleader, Aurelia, the Law Above, or Otharri, Suns' Glory? Say "none" and I name three more.

Faults: omits information. Catalog action: reword.

The user asked for “a Background commander pair,” but the choices list only individual commanders and do not establish a Background pairing; reword to ask which Background commander and Background they prefer, or whether the agent should choose the pair.

### 82. terse: the user corrects a card name, turn 3, row `named_card_role`

**Asked:** Do you want Atraxa, Praetor's Voice as your commander or in the 99?

Faults: assumes an answer. Catalog action: remove.

The user had already said “Build around Atraxa” and “Commander,” which indicates Atraxa is to be the commander; asking whether to put Atraxa in the 99 ignores that context.

### 100. terse: a deck for a team event, turn 2, row `budget`

**Asked:** What spending limit, if any, should I use for cards you don’t own?

Faults: assumes an answer. Catalog action: reword.

The user had not said they owned any cards, and user_has_a_card_collection is false; reword to “What spending limit, if any, should I use for the deck?”

## Slots the eval says nobody asked about

These do not move the ratio. They are the guard against a quiet agent.

- 100. terse: a deck for a team event: Ask whether the team event has shared-card or other team-format deck-building restrictions.
- 102. terse: stax with no table named: After the user said “None of those,” ask what Commander bracket or power level they want instead.
- 103. terse: two themes at once: Ask how the deck should balance its artifact and lifegain themes.
- 16. the strongest modern deck: budget amount
- 2. blink with a thin library: Ask the user to provide or connect their library inventory so the deck can be built owned-first.
- 28. reanimator with two plans: Ask what the second plan should be; the user specified “Reanimate one big creature, not a value loop” but did not identify the other plan.
- 36. the user changes the theme late (probe): Whether to keep Meren of Clan Nel Toth as commander for the token deck.
- 40. the user asks a question back (probe): Turn 2: Explain what a Commander bracket is, since the user asked, “What is a bracket?”
- 44. a card that is not legal (probe): Clarify that Black Lotus is not legal in Modern and ask what legal card or theme should replace it.
- 47. a deck as a gift (probe): After the user answered the bracket question and gave a budget but did not answer the commander question, ask whether they have a commander in mind or want three suggestions.
- 48. sideboard help only (probe): Ask for the current sideboard or decklist so the recommendations can account for cards already included., Ask which control decks or other matchups the user wants the sideboard to address; the user said their shop plays a lot of control.
- 53. the user swaps the commander (probe): After the user said “Actually use a different commander, suggest one,” recommend a replacement commander rather than asking the user to choose one.
- 54. two decks at once (probe): After finishing the Commander deck, continue intake for the requested Modern deck.
- 56. the user answers a different question (probe): Ask whether to keep the mill theme despite the user saying their friends hate mill.
- 60. terse: a format we do not build: budget, commander
- 63. terse: a competitive request: Turn 2: clarify the user's budget, since “under budget” does not specify an amount.
- 66. terse: sideboard help only: Current Modern Burn decklist, so sideboard recommendations can fit the cards the user already plays.
- 69. terse: a Pauper deck named once: Which supported format should the deck use? The user gave a power level and budget but did not choose a format.
- 76. terse: a companion: deck colors
- 77. terse: a Background pair: Ask whether the user has a specific Background commander and Background pairing in mind, or wants the agent to choose one.
- 80. terse: all caps, no punctuation: Ask how common control is in the shop's metagame, or what control decks are typically played, to tailor the tournament burn list.
- 85. terse: chaos: How should the deck make play as random as possible—for example, by using random effects, unpredictable choices, or both? The user’s chaos preference is clear, but the specific approach is not.
- 90. terse: a deck for a spouse: Whether the deck should be cat-themed.
- 95. terse: a library that does not exist: Ask the user to provide their library/card list so the deck can be built from cards they own.

## Run

- Suite `question-eval`, run `pr7-question-eval-run64`, on 2026-10-09, commit `de24541`.
- Roles: eval on `gpt-6-luna` (openai, effort low).
- Versions: eval prompt version 3, gate_document `pr7-question-gate-run64`.
- Calls: 106. Cost: $0.0414. Time: 536 seconds.
- Tokens: 112647 input (0 cached), 56548 output.
- Eval cost: $0.0414.

## The card check

The check ran against the local snapshot. No refusal rested on a card name the snapshot refutes.
