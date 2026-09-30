# PR-7 question eval

Run: `pr7-question-gate-run54`. Eval model: `gpt-5.6-luna`. Prompt version 3.

Scored 415 questions. 17 were not warranted, and 0 are unsure.

**Bad-question ratio: 4.7% on the holdout.**

The tune split holds 288 questions at 3.8%, and the holdout holds 127 at 4.7%. A fixer reads the tune split alone. A gain that shows only there is a reworded test set (D-134).

## Conversations the eval left unjudged

The verdicts did not line up with the questions asked, one per question. Their verdicts are dropped and not counted (T-9).

- 53. the user swaps the commander (probe): 1 verdicts for 4 questions

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
| `colors` | 3 |
| `theme` | 3 |
| `format_unsupported_open` | 2 |
| `commander` | 1 |
| `commander_pick` | 1 |
| `format` | 1 |
| `house_format_limits` | 1 |
| `power_sixty` | 1 |

## Faults

| Fault | Count |
|---|---|
| assumes an answer | 5 |
| out of scope | 4 |
| duplicate | 3 |
| omits information | 2 |
| inaccurate | 1 |
| offers an option that does not apply | 1 |

## What the catalog needs

| Action | Count |
|---|---|
| reword | 19 |
| remove | 10 |

## Every question the eval refused

The holdout conversations are left out of this list on purpose. 6 of their questions failed, and the fixer may not read which (D-134).

### 5. 60-card anything goes, turn 1, row `format`

**Asked:** Which format would you like: Standard or Modern?

Faults: omits information. Catalog action: reword.

The user already specified “60-card” and “anything goes at our table,” so the question should acknowledge a custom anything-goes format rather than offering only Standard or Modern; proposed wording: “Should I treat this as a custom anything-goes 60-card format, or would you like to use a named format such as Modern?”

### 5. 60-card anything goes, turn 1, row `house_format_limits`

**Asked:** Inside your house format, do the normal 60-card deck limits hold?

Catalog action: remove.

“Anything goes at our table” already indicates that the normal deck limits do not necessarily apply, so this question asks for an answer the user has effectively supplied.

### 7. FNM on Friday (probe), turn 2, row `colors`

**Asked:** Do you have a color preference?

Faults: duplicate. Catalog action: remove.

The agent had already asked about color preference on turn 1, and the user still had not supplied one; repeating the same question added no new clarification.

### 7. FNM on Friday (probe), turn 2, row `budget`

**Asked:** What spending limit, if any, should I use for cards you don’t own?

Faults: duplicate. Catalog action: remove.

The agent had already asked for a spending limit on turn 1. Although “under budget” still lacked a dollar amount, this question repeats the same request rather than clarifying the missing amount.

### 35. the user changes the format (probe), turn 3, row `budget`

**Asked:** What spending limit, if any, should I use for cards you don’t own?

Faults: duplicate. Catalog action: remove.

The spending-limit question had already been asked on turn 1: “What spending limit, if any, should I use for cards you don’t own?” The user had not answered it yet, but repeating it on turn 3 is still a duplicate question.

### 43. a card name with a typo (probe), turn 3, row `budget`

**Asked:** What spending limit, if any, should I use for cards you don’t already own?

Faults: assumes an answer. Catalog action: reword.

The user had already said “from my library,” selecting an in-collection approach, so asking about spending on cards they do not own presumes that outside purchases are allowed. Proposed wording: “If cards outside your library are allowed, what spending limit, if any, should I use?”

### 44. a card that is not legal (probe), turn 1, row `theme`

**Asked:** What should the deck focus on: a creature type, mechanic, or play style?

Faults: assumes an answer. Catalog action: reword.

The user already said the deck should be built around “Black Lotus,” so asking what the deck should focus on overlooks the stated theme. Reword it to address Black Lotus's Modern legality and ask whether the user wants a replacement card or a different theme.

### 49. copy a known deck (probe), turn 1, row `budget`

**Asked:** What spending limit, if any, should I use for cards you don’t own?

Faults: assumes an answer. Catalog action: reword.

The user did not say they own any cards, while the question presumes a split between cards they own and cards they do not; reword it as: "What total spending limit, if any, should I use for the deck?"

### 77. terse: a Background pair, turn 3, row `commander_pick`

**Asked:** Which commander do you want: Aurelia, the Warleader, Aurelia, the Law Above, or Otharri, Suns' Glory? Say "none" and I name three more.

Faults: offers an option that does not apply. Catalog action: reword.

The user requested a Background commander pair, but the offered choices are individual commanders and do not address selecting an eligible Background pairing; reword this row to offer commanders that support the requested format or ask whether the user wants to change formats.

### 82. terse: the user corrects a card name, turn 1, row `colors`

**Asked:** Do you have a color preference?

Faults: out of scope. Catalog action: remove.

Atraxa, Praetor's Voice already fixes the deck's color identity, so a general color-preference question is not actionable for a deck built around that commander.

### 88. terse: no theme at all, turn 1, row `theme`

**Asked:** What should the deck focus on: a creature type, a mechanic, or a play style?

Faults: assumes an answer. Catalog action: reword.

The user said “Surprise me” and the conversation description says “no theme at all,” so the agent should choose the theme rather than ask the user to specify one; reword as “I’ll choose the deck’s theme unless you want to specify one.”

## Slots the eval says nobody asked about

These do not move the ratio. They are the guard against a quiet agent.

- 103. terse: two themes at once: Which theme should be the primary focus, or should artifacts and lifegain receive equal emphasis?
- 104. terse: the user answers with a number: commander clarification: the user answered “2” to a prompt with no numbered commander choices, so the agent should clarify what “2” refers to before proceeding.
- 16. the strongest modern deck: After the user changed the request to “The best deck under budget,” ask what the budget is.
- 23. land destruction: budget clarification after the user's non-answer: “My table does not mind” does not state a spending limit for cards they do not own.
- 28. reanimator with two plans: The deck’s second plan or backup win condition, since the conversation title specifies “two plans” but the user only clarified that the primary plan is reanimating one big creature rather than a value loop.
- 36. the user changes the theme late (probe): After the user changed the theme, ask whether to keep Commander, black-green, and Meren of Clan Nel Toth, or revise the commander and color identity for the token deck., Ask what kind of token strategy or token types the user wants, since “a token deck” is broad.
- 44. a card that is not legal (probe): Modern legality check and a request for a legal replacement or revised theme for Black Lotus
- 48. sideboard help only (probe): The agent should ask for the current deck or sideboard list, including which sideboard cards the user already owns, so it can recommend specific changes rather than build without the deck’s existing contents.
- 52. a price cap per card (probe): total budget
- 63. terse: a competitive request: The user's actual budget limit, since they changed the request to “The best deck under budget” without naming an amount.
- 66. terse: sideboard help only: The current Modern burn maindeck or at least its key cards and flex slots, so the sideboard can be tailored to the deck., Expected opposing decks or matchups, since sideboard choices depend on the metagame; the user later mentions that the shop has a lot of control.
- 7. FNM on Friday (probe): A follow-up asking whether the user wants a supported format instead of Pioneer, since the requested Pioneer deck cannot be built., A color-preference question, or confirmation that the agent should choose the strongest colors, because the user never supplied colors.
- 74. terse: a card the user does not own: After the user said “Buy it then,” clarify whether purchases are now allowed despite the original constraint “only the cards I own.”
- 76. terse: a companion: Clarify whether the $300 budget covers only the maindeck or also the sideboard and companion card.
- 77. terse: a Background pair: Ask which Background, or whether the user has a specific Background commander pair in mind, since the request explicitly specifies a Background pair.
- 8. mill for a playgroup: The budget amount's currency and whether the $100 cap applies to purchases or the whole deck remain unresolved after the user did not answer the budget-scope question.
- 87. terse: landfall: Clarify what “library first” means and whether it supersedes the earlier “from my binder” constraint.
- 9. a commander the library does not hold: Confirm whether Brago, King Eternal should be added to the buy list because the user said the library does not hold him., Confirm that Brago, King Eternal is intended to be the commander rather than merely part of the blink theme.
- 94. terse: a collection of lands only: Ask which lands, and how many of each, the user owns so the deck can use their existing collection accurately.
- 95. terse: a library that does not exist: Ask whether the user has a card list or collection to build from, since the user requested a deck from their library but the app knows they have no card collection.
- 96. terse: the user asks about a card: Answer whether Sol Ring is legal in Modern.
- 97. terse: Pauper Commander: The user did not choose a replacement format after being offered Commander, Standard, or Modern; ask which supported format to use before building.
- 99. terse: Historic on Arena: Clarify the platform/legality conflict after the user selected Modern: Modern is not an Arena format, so ask whether they want an Arena-legal deck or a paper Modern deck.

## Run

- Suite `question-eval`, run `pr7-question-eval-run54-luna56`, on 2026-09-29, commit `6cdb2e1`.
- Roles: eval on `gpt-5.6-luna` (openai, effort low).
- Versions: eval prompt version 3, gate_document `pr7-question-gate-run54`.
- Calls: 109. Cost: $0.0977. Time: 628 seconds.
- Tokens: 111226 input (0 cached), 59862 output.
- Eval cost: $0.0977.

## The card check

The check ran against the local snapshot. No refusal rested on a card name the snapshot refutes.
