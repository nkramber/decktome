# PR-7 question eval

Run: `pr7-question-gate-run33`. Eval model: `gpt-5.6-luna`. Prompt version 3.

Scored 404 questions. 22 were not warranted, and 0 are unsure.

**Bad-question ratio: 10.3% on the holdout.**

The tune split holds 287 questions at 3.5%, and the holdout holds 117 at 10.3%. A fixer reads the tune split alone. A gain that shows only there is a reworded test set (D-134).

## Conversations the eval left unjudged

The verdicts did not line up with the questions asked, one per question. Their verdicts are dropped and not counted (T-9).

- 27. a card for the 99: 3 verdicts for 6 questions
- 3. a named card, role unknown: 3 verdicts for 5 questions

## The counters the ratio can not see

A run that asks less scores better and serves the user worse. A run can pass both bars with most of its conversations unanswered, so read these counters with the ratio.

| Counter | Value |
|---|---|
| Questions asked | 119 |
| Questions that closed a slot | 98 |
| Invented by the model | 0 |
| Catalog-only conversations | 27 |
| Premature sessions | 0 |
| Linter findings | 0 |

## Where the bad questions came from

| Row | Bad questions |
|---|---|
| `colors` | 5 |
| `budget` | 3 |
| `format_unsupported_open` | 3 |
| `commander` | 2 |
| `pool_thin` | 2 |
| `theme` | 2 |
| `format` | 1 |
| `pool` | 1 |
| `pool_precon` | 1 |
| `power_commander` | 1 |
| `power_sixty` | 1 |

## Faults

| Fault | Count |
|---|---|
| out of scope | 7 |
| duplicate | 4 |
| assumes an answer | 3 |
| inaccurate | 1 |

## What the catalog needs

| Action | Count |
|---|---|
| reword | 25 |
| remove | 10 |
| duplicate | 1 |

## Every question the eval refused

The holdout conversations are left out of this list on purpose. 12 of their questions failed, and the fixer may not read which (D-134).

### 5. 60-card anything goes, turn 1, row `format`

**Asked:** Which format should the 60-card deck use: Standard or Modern?

Faults: out of scope. Catalog action: reword.

The user already said the deck is “anything goes at our table,” so forcing a choice between Standard and Modern does not fit the stated house format. Reword as: “Should I treat this as your anything-goes house format, or apply an official format’s rules?”

### 7. FNM on Friday (probe), turn 2, row `colors`

**Asked:** Do you have a color preference for the deck?

Faults: duplicate. Catalog action: duplicate.

The agent had already asked “Do you have a color preference?” on turn 1, and the turn 2 wording asks the same unanswered question again.

### 7. FNM on Friday (probe), turn 2, row `budget`

**Asked:** What spending limit, if any, should I use for cards you do not own?

Faults: duplicate. Catalog action: remove.

The agent had already asked “What spending limit, if any” on turn 1; repeating it on turn 2 is unnecessary, even though the user answered it on turn 3.

### 17. precon upgrade at bracket 2, turn 1, row `pool_precon`

**Asked:** Should I start with the precon, limit the deck to cards from it, or build without using it?

Catalog action: reword.

The user already asked to “Upgrade my Atraxa, Praetors' Voice precon,” which establishes using the precon as the starting pool; if a restriction is needed, ask only whether cards outside the precon are allowed.

### 26. legacy for an event (probe), turn 1, row `format_unsupported_open`

**Asked:** I do not build Legacy. Which format should I build instead: Commander, Standard, or Modern?

Faults: assumes an answer. Catalog action: reword.

The user already specified “A Legacy deck for an event,” so asking which other format to use assumes they want to abandon Legacy. Reword it to explain the Legacy limitation and ask whether they want to switch formats.

### 35. the user changes the format (probe), turn 3, row `budget`

**Asked:** What spending limit, if any, should I use for cards you do not own?

Faults: duplicate. Catalog action: remove.

The agent already asked, “What spending limit, if any, should I use for cards you do not own?” on turn 1. The user had not answered it, but repeating the same question on turn 3 is still a duplicate.

### 49. copy a known deck (probe), turn 1, row `colors`

**Asked:** Do you have a color preference?

Catalog action: remove.

The user said “whatever is winning,” which already indicates no color preference, so asking about colors was unnecessary.

### 64. terse: a deck as a gift (probe), turn 1, row `budget`

**Asked:** What spending limit, if any, should I use for cards you do not own?

Faults: assumes an answer. Catalog action: reword.

The question presumes the user owns some cards by referring to “cards you do not own,” although user_has_a_card_collection is false. Reword it as: “What total budget, if any, should I use for the deck?”

### 82. terse: the user corrects a card name (probe), turn 1, row `colors`

**Asked:** Do you have a color preference?

Faults: out of scope. Catalog action: remove.

The user already chose Atraxa, Praetor's Voice, whose commander color identity determines the deck's colors; a separate color preference cannot change that identity.

### 88. terse: no theme at all (probe), turn 1, row `theme`

**Asked:** Would you like the deck built around a creature type, a mechanic, or a play style, or should I choose one for you?

Catalog action: remove.

The user already delegated the choice by saying “Surprise me,” which answers whether the agent should choose the theme.

## Slots the eval says nobody asked about

These do not move the ratio. They are the guard against a quiet agent.

- 102. terse: stax with no table named (probe): Clarify the user's power target after they replied "None of those" to the listed brackets.
- 108. a set beside a theme (probe): After the user said “Green and white, and you pick the commander,” the agent should have selected a commander and continued the deck-building process.
- 16. the strongest modern deck: A precise budget amount should have been requested after the user said “The best deck under budget,” before recommending a deck.
- 19. a big library, no theme: Ask which cards or categories from the user's large collection should be prioritized, or request an inventory/list of the collection so the deck can be built from it.
- 23. land destruction: Follow up on the unanswered budget question, such as: “What spending limit, if any, should I use for cards you do not own?”, Clarify the intended scope of land destruction, such as whether mass land destruction is acceptable or whether the deck should focus on targeted land removal.
- 36. the user changes the theme late (probe): After the user changed the theme to tokens, ask whether Meren of Clan Nel Toth and the black-green color identity should remain, and clarify the desired token strategy.
- 44. a card that is not legal (probe): Flag that Black Lotus is not legal in Modern and ask whether the user wants a legal replacement or a different format.
- 48. sideboard help only (probe): format, deck colors/archetype, expected opposing decks or local metagame
- 52. a price cap per card (probe): total budget for the deck
- 63. terse: a competitive request (probe): budget amount after the user changed the request to “The best deck under budget.”
- 69. terse: a Pauper deck named once (probe): Target power level or intended casual competitiveness for the Pauper burn deck.
- 70. terse: a mill deck for a playgroup (probe): commander_color_identity_conflict: Bruvac the Grandiloquent's color identity does not support a black card pool, so ask whether to keep Bruvac and build mono-blue or switch commanders to retain blue-black.
- 74. terse: a card the user does not own (probe): After the user said “Buy it then,” ask which cards may be purchased and what budget, if purchasing is supported.
- 75. terse: cEDH at bracket 5 (probe): Whether the target is bracket 5, or another specific power-level bracket; the user later supplied “Yes, bracket 5” without having been asked.
- 79. terse: everything buried in one long message (probe): Which black and green cards or relevant sacrifice cards do you already own and want considered for the deck?
- 8. mill for a playgroup: The currency for the 100 budget was not clarified.
- 92. terse: proxy the expensive cards only (probe): Follow up on the unresolved commander choice: Bruvac is mono-blue, so the agent should ask whether the deck should remain blue-black with Bruvac in the 99 or become mono-blue with Bruvac as commander., Follow up on the unresolved proxy/budget scope, since the user did not answer whether the $20 rule applies to purchases or total deck value.
- 96. terse: the user asks about a card (probe): Answer the user's direct question about whether Sol Ring is legal in Modern; the agent never addressed card legality.
- 97. terse: Pauper Commander (probe): After the user selects a supported format, ask whether they have a preferred commander or want the agent to recommend one for the black sacrifice strategy.

## Run

- Calls: 103. Time: 580.9 seconds.
- Tokens: 111815 input (0 cached), 62440 output.
- Cost: $0.0973.

## The card check

The check ran against the local snapshot. No refusal rested on a card name the snapshot refutes.
