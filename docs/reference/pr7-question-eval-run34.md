# PR-7 question eval

Run: `pr7-question-gate-run34`. Eval model: `gpt-5.6-luna`. Prompt version 3.

Scored 392 questions. 18 were not warranted, and 0 are unsure.

**Bad-question ratio: 6.8% on the holdout.**

The tune split holds 274 questions at 3.6%, and the holdout holds 118 at 6.8%. A fixer reads the tune split alone. A gain that shows only there is a reworded test set (D-134).

## Conversations the eval left unjudged

The verdicts did not line up with the questions asked, one per question. Their verdicts are dropped and not counted (T-9).

- 22. extra turns commander: 4 verdicts for 7 questions
- 39. the user stays vague (probe): 4 verdicts for 5 questions
- 51. the user declines everything (probe): a verdict on turn 1 row power, which asked no such question
- 53. the user swaps the commander (probe): 1 verdicts for 4 questions

## The counters the ratio can not see

A run that asks less scores better and serves the user worse. A run can pass both bars with most of its conversations unanswered, so read these counters with the ratio.

| Counter | Value |
|---|---|
| Questions asked | 117 |
| Questions that closed a slot | 90 |
| Invented by the model | 1 |
| Catalog-only conversations | 26 |
| Premature sessions | 0 |
| Linter findings | 0 |

## Where the bad questions came from

| Row | Bad questions |
|---|---|
| `colors` | 4 |
| `format_unsupported_open` | 3 |
| `theme` | 3 |
| `budget` | 2 |
| `commander` | 1 |
| `commander_pick` | 1 |
| `format` | 1 |
| `pool_precon` | 1 |
| `pool_thin` | 1 |
| `power_sixty` | 1 |

## Faults

| Fault | Count |
|---|---|
| duplicate | 9 |
| out of scope | 7 |
| assumes an answer | 3 |
| inaccurate | 2 |
| omits information | 1 |

## What the catalog needs

| Action | Count |
|---|---|
| remove | 13 |
| reword | 13 |

## Every question the eval refused

The holdout conversations are left out of this list on purpose. 8 of their questions failed, and the fixer may not read which (D-134).

### 5. 60-card anything goes, turn 1, row `format`

**Asked:** Which format should the 60-card deck use?

Faults: duplicate, omits information. Catalog action: reword.

The user already said, “anything goes at our table,” so asking which format to use is redundant; the offered Standard and Modern choices also omit the stated anything-goes house format. Reword to: “Should I treat this as an anything-goes house format, or follow a named sanctioned format?”

### 7. FNM on Friday (probe), turn 2, row `colors`

**Asked:** Do you have a color preference?

Faults: duplicate. Catalog action: remove.

The agent had already asked “Do you have a color preference?” on turn 1, and the user did not provide an answer that required re-asking it.

### 10. no strong commander in the library, turn 2, row `commander`

**Asked:** Which commander do you want? Name one, or I suggest three.

Faults: assumes an answer. Catalog action: remove.

The user had already answered this by saying, “Buy the best lifegain commander,” so asking “Which commander do you want?” wrongly presumes the user must choose one or request suggestions instead of having the agent select the best commander.

### 26. legacy for an event (probe), turn 1, row `format_unsupported_open`

**Asked:** I do not build Legacy. Which format should I build instead: Commander, Standard, or Modern?

Faults: out of scope. Catalog action: reword.

The user asked for “A Legacy deck for an event,” so offering Commander, Standard, or Modern as replacements does not address the requested format; reword to explain the Legacy limitation and ask whether the user wants an alternative format.

### 35. the user changes the format (probe), turn 3, row `budget`

**Asked:** What spending limit, if any, should I use for cards you do not own?

Faults: duplicate. Catalog action: remove.

The budget was already requested on turn 1 and had not been superseded; asking again repeats the same unanswered slot rather than adding a new requirement.

### 43. a card name with a typo (probe), turn 3, row `budget`

**Asked:** What spending limit, if any, should I use for cards you do not own?

Faults: out of scope. Catalog action: reword.

The user specified “from my library,” so asking for a spending limit for “cards you do not own” is not actionable for this build. Proposed wording: “If you want to include cards outside your library, what spending limit should I use?”

### 44. a card that is not legal (probe), turn 1, row `theme`

**Asked:** What should the deck do: focus on a creature type, a mechanic, or a play style?

Faults: duplicate. Catalog action: remove.

The user already supplied the deck's central theme by asking for a deck built around Black Lotus.

### 49. copy a known deck (probe), turn 1, row `colors`

**Asked:** Do you have a color preference?

Faults: duplicate. Catalog action: remove.

The user already said “whatever is winning,” which answers that they have no color preference and want the strongest deck regardless of colors.

### 62. terse: a precon upgrade (probe), turn 1, row `pool_precon`

**Asked:** Should I start with your precon, limit the deck to cards from it, or set it aside for a fully optimized build?

Faults: duplicate. Catalog action: remove.

The user already asked to upgrade their “precon,” which establishes that the precon should be the starting point, so the question repeats an answered choice.

### 88. terse: no theme at all (probe), turn 1, row `theme`

**Asked:** What should the deck do: use a creature type, a mechanic, or a play style? You can also leave it entirely to me.

Catalog action: reword.

The user already delegated the choice by saying “Surprise me,” so asking what the deck should do is unnecessary; the row should instead let the agent proceed when the user has already chosen delegation.

## Slots the eval says nobody asked about

These do not move the ratio. They are the guard against a quiet agent.

- 100. terse: a deck for a team event (probe): Team-event rules or deck-construction constraints, such as whether a unified-deck restriction applies, Preferred strategy or archetype, such as control, aggro, or combo
- 102. terse: stax with no table named (probe): Clarify what the user's “None of those” refers to before proceeding.
- 103. terse: two themes at once (probe): Whether artifacts and lifegain should have equal priority or whether one theme should be primary.
- 104. terse: the user answers with a number (probe): Clarify which commander the user wants after the reply “2,” since that answer does not identify a commander or clearly select “Suggest one.”
- 105. a theme the library can not lead (probe): After the user said "Colorless," ask for or select a commander whose color identity is colorless, or explain that the available commander candidates do not support a colorless deck.
- 16. the strongest modern deck: budget amount after the user said “The best deck under budget.”
- 18. pauper burn on a budget (probe): Power level or intended play environment
- 23. land destruction: Clarify the spending limit for cards not in the user's library after “My table does not mind,” since that answer does not specify a budget.
- 28. reanimator with two plans: After the user clarified “Reanimate one big creature, not a value loop,” ask whether they have a particular reanimation target, creature theme, or preferred way to repeatedly or protect that threat.
- 36. the user changes the theme late (probe): After the user changed the theme, ask whether the Commander format, black-green colors, and Meren of Clan Nel Toth should remain, or whether those choices should also change., Ask what kind of token strategy the user wants, such as creature type, go-wide combat, sacrifice, or another token-focused plan.
- 46. two partner commanders (probe): Whether Thrasios, Triton Hero and Tymna the Weaver should both be used as partner commanders.
- 48. sideboard help only (probe): Current sideboard list or existing sideboard constraints
- 52. a price cap per card (probe): A total deck budget, since a per-card cap does not establish how much the complete Modern deck may cost.
- 55. a returning player (probe): Follow up on the unanswered commander choice; the user supplied colors, power, and budget but did not name or request a commander., Confirm that Commander is the intended format, since the user said “Commander, I think.”
- 56. the user answers a different question (probe): A spending limit for cards the user does not own remained unanswered after the user specified Commander, blue and black, and Bracket 3., Whether to proceed with a mill strategy despite the user's statement that their friends hate mill, or adapt the deck to be more acceptable to that playgroup.
- 57. terse: the format is named once (probe): Whether the land-destruction plan should emphasize targeted land destruction, mass land destruction, or a mix.
- 60. terse: a format we do not build (probe): After the user answered “Bracket 2,” follow up on the still-unanswered budget and commander questions.
- 61. terse: a proxy table (probe): desired power level
- 63. terse: a competitive request (probe): The agent should have asked about the target metagame or expected field, since the strongest Modern deck depends on the local or tournament environment.
- 66. terse: sideboard help only (probe): Ask for the current sideboard and, if relevant, the deck list or key cards, since sideboard recommendations depend on what is already included., Ask which matchups or archetypes the sideboard should target; the later note that the shop plays a lot of control supplies this information, but no agent question elicited it.
- 77. terse: a Background pair (probe): A question asking which Background the user wants paired with the commander, or whether the agent should choose one.
- 79. terse: everything buried in one long message (probe): Which black and green cards from your collection should the deck prioritize, or can you provide your card list?, Are there any sacrifice themes or cards you especially want to include or avoid?, What kinds of decks do your Thursday opponents usually play, so the deck can stay appropriate for that casual environment?
- 80. terse: all caps, no punctuation (probe): budget follow-up
- 85. terse: chaos (probe): Which commander should lead the deck? The user gave a theme but no commander.
- 95. terse: a library that does not exist (probe): Follow up on the unanswered budget question for cards outside the user's library., Follow up on the unanswered commander choice.
- 96. terse: the user asks about a card (probe): Answer whether Sol Ring is legal in Modern.
- 97. terse: Pauper Commander (probe): A commander or commander-selection preference, if the user accepts regular Commander instead of Pauper Commander.

## Run

- Suite `question-eval`, run `pr7-question-eval-run34`, on 2026-09-04, commit `ec058b4`.
- Roles: eval on `gpt-5.6-luna` (openai, effort low).
- Versions: eval prompt version 3, gate_document `pr7-question-gate-run34`.
- Calls: 103. Cost: $0.0923. Time: 713 seconds.
- Tokens: 112823 input (0 cached), 58110 output.
- Eval cost: $0.0923.

## The card check

The check ran against the local snapshot. No refusal rested on a card name the snapshot refutes.
