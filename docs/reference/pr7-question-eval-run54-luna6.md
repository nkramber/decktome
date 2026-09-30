# PR-7 question eval

Run: `pr7-question-gate-run54`. Eval model: `gpt-6-luna`. Prompt version 3.

Scored 410 questions. 15 were not warranted, and 0 are unsure.

**Bad-question ratio: 4.7% on the holdout.**

The tune split holds 283 questions at 3.2%, and the holdout holds 127 at 4.7%. A fixer reads the tune split alone. A gain that shows only there is a reworded test set (D-134).

## Conversations the eval left unjudged

The verdicts did not line up with the questions asked, one per question. Their verdicts are dropped and not counted (T-9).

- 40. the user asks a question back (probe): 4 verdicts for 5 questions
- 53. the user swaps the commander (probe): 2 verdicts for 4 questions

## The counters the ratio can not see

A run that asks less scores better and serves the user worse. A run can pass both bars with most of its conversations unanswered, so read these counters with the ratio.

| Counter | Value |
|---|---|
| Questions asked | 290 |
| Questions that closed a slot | 230 |
| Invented by the model | 0 |
| Catalog-only conversations | 75 |
| Premature sessions | 0 |
| Linter findings | 0 |

## Where the bad questions came from

| Row | Bad questions |
|---|---|
| `budget` | 4 |
| `colors` | 4 |
| `theme` | 3 |
| `commander` | 1 |
| `commander_pick` | 1 |
| `power_commander` | 1 |
| `power_sixty` | 1 |

## Faults

| Fault | Count |
|---|---|
| assumes an answer | 5 |
| out of scope | 5 |
| duplicate | 4 |

## What the catalog needs

| Action | Count |
|---|---|
| reword | 22 |
| remove | 5 |

## Every question the eval refused

The holdout conversations are left out of this list on purpose. 6 of their questions failed, and the fixer may not read which (D-134).

### 1. lifegain with a collection, turn 3, row `commander_pick`

**Asked:** Which commander do you want: Astarion, the Decadent, Karlov of the Ghost Council, or Sorin of House Markov // Sorin, Ravenous Neonate? Say "none" and I name three more.

Faults: assumes an answer. Catalog action: reword.

The user said, “Suggest a commander from my library,” so asking them to choose presumes they want to make the selection; reword this to offer a commander suggestion from their library.

### 7. FNM on Friday (probe), turn 2, row `colors`

**Asked:** Do you have a color preference?

Faults: duplicate. Catalog action: none.

The agent had already asked “Any color preference?” on turn 1, and the user had not answered it; repeating it on turn 2 was a duplicate.

### 7. FNM on Friday (probe), turn 2, row `budget`

**Asked:** What spending limit, if any, should I use for cards you don’t own?

Faults: duplicate. Catalog action: none.

The agent had already asked about a spending limit on turn 1, and the user had not answered it; repeating it on turn 2 was a duplicate.

### 35. the user changes the format (probe), turn 3, row `budget`

**Asked:** What spending limit, if any, should I use for cards you don’t own?

Faults: duplicate. Catalog action: none.

The agent had already asked for a spending limit on turn 1, so repeating that question was redundant.

### 43. a card name with a typo (probe), turn 3, row `budget`

**Asked:** What spending limit, if any, should I use for cards you don’t already own?

Faults: out of scope. Catalog action: reword.

The user said “from my library,” making a spending limit for cards they do not own irrelevant; reword this row to ask whether any cards outside the library may be included before asking about their spending limit.

### 44. a card that is not legal (probe), turn 1, row `theme`

**Asked:** What should the deck focus on: a creature type, mechanic, or play style?

Catalog action: reword.

The user already named the focus with “built around Black Lotus”; reword this to ask what play style or other theme they want beyond that named card.

### 82. terse: the user corrects a card name, turn 1, row `colors`

**Asked:** Do you have a color preference?

Faults: out of scope. Catalog action: remove.

The user already named the commander, whose color identity constrains the deck, so a general color preference cannot guide this deck-building request.

### 88. terse: no theme at all, turn 1, row `theme`

**Asked:** What should the deck focus on: a creature type, a mechanic, or a play style?

Faults: assumes an answer. Catalog action: reword.

“Surprise me” delegates the choice of focus to the agent; reword to make choosing a theme the default and invite preferences only optionally.

### 100. terse: a deck for a team event, turn 2, row `budget`

**Asked:** What spending limit, if any, should I use for cards you don’t own?

Faults: assumes an answer. Catalog action: reword.

The known field says the user has no card collection, but “for cards you don’t own” presumes some cards are owned; reword to: “What total spending limit, if any, should I use for the deck?”

## Slots the eval says nobody asked about

These do not move the ratio. They are the guard against a quiet agent.

- 103. terse: two themes at once: Which theme should be the primary focus, or should artifacts and lifegain be balanced equally?
- 104. terse: the user answers with a number: Clarify what the user's “2” means for the commander choice; no commander or suggestion was selected.
- 107. a set name that names two base sets (probe): Clarify whether “Tarkir: Dragonstorm” is a set restriction and, since it does not name a commander, ask whether the user wants a commander suggested.
- 108. a set beside a theme (probe): Whether the deck should use only Bloomburrow cards or may include cards from any set.
- 16. the strongest modern deck: After turn 2, ask what budget the user means by “under budget.”
- 18. pauper burn on a budget (probe): power level
- 20. dinosaur tribal for a child: Ask whether the child needs age-appropriate simplification or accessibility considerations, since the deck is for the user's kid.
- 23. land destruction: Clarify the spending limit for cards not in the user's library; “My table does not mind” does not specify a budget.
- 26. legacy for an event (probe): Target event competitiveness or tournament level.
- 28. reanimator with two plans: Ask what the deck’s second plan should be; the user specified reanimating one big creature, not a value loop, but did not name another plan.
- 37. the user contradicts the budget (probe): Color preference remains unanswered; the user did not respond to the colors question.
- 44. a card that is not legal (probe): Turn 1: Clarify that Black Lotus is not Modern-legal and ask whether the user wants a Modern-legal replacement; the user later says, “Then pick something else.”
- 48. sideboard help only (probe): The user's current sideboard list, if they want help improving or revising it.
- 50. a format we do not support (probe): commander
- 52. a price cap per card (probe): Ask for a total deck budget; the user set a $5-per-card cap but did not initially specify an overall budget.
- 55. a returning player (probe): commander preference: the user did not answer whether they wanted to name a commander or have one suggested, so follow up on that choice
- 63. terse: a competitive request: budget amount, because the user later asked for the best deck “under budget” without specifying an amount
- 66. terse: sideboard help only: Ask which opposing decks or matchups the sideboard should prioritize, including how heavily to target control.
- 69. terse: a Pauper deck named once: Preferred power level or how competitive the deck should be.
- 7. FNM on Friday (probe): Turn 3: Ask which supported format the user wants instead of Pioneer; the format is still unresolved., Turn 3: Ask for a color preference; the user has not answered that question.
- 72. terse: a five-color deck: budget clarification after the user did not answer the spending-limit question, power-bracket clarification after the user said “whatever is strongest”
- 74. terse: a card the user does not own: Clarify whether “Buy it then” changes the requirement to use only cards the user owns.
- 76. terse: a companion: color identity
- 77. terse: a Background pair: Background partner selection
- 89. terse: teaching a new player: Ask how simple or beginner-friendly the deck should be to pilot, since the user said it is for someone who has never played before.
- 9. a commander the library does not hold: Ask whether Brago should be added to the buy list, since the requested commander may not be in the user's library.
- 95. terse: a library that does not exist: Ask the user to provide their card list, since the app has no collection to build from; otherwise clarify that it cannot build from a library it cannot access.
- 96. terse: the user asks about a card: Answer the user's direct question about whether Sol Ring is legal in Modern.

## Run

- Suite `question-eval`, run `pr7-question-eval-run54-luna6`, on 2026-09-29, commit `6cdb2e1`.
- Roles: eval on `gpt-6-luna` (openai, effort low).
- Versions: eval prompt version 3, gate_document `pr7-question-gate-run54`.
- Calls: 105. Cost: $0.0424. Time: 640 seconds.
- Tokens: 111226 input (0 cached), 58906 output.
- Eval cost: $0.0424.

## The card check

The check ran against the local snapshot. No refusal rested on a card name the snapshot refutes.
