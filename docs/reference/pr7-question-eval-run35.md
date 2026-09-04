# PR-7 question eval

Run: `pr7-question-gate-run35`. Eval model: `gpt-5.6-luna`. Prompt version 3.

Scored 411 questions. 23 were not warranted, and 0 are unsure.

**Bad-question ratio: 7.3% on the holdout.**

The tune split holds 288 questions at 4.9%, and the holdout holds 123 at 7.3%. A fixer reads the tune split alone. A gain that shows only there is a reworded test set (D-134).

## Conversations the eval left unjudged

The verdicts did not line up with the questions asked, one per question. Their verdicts are dropped and not counted (T-9).

- 57. terse: the format is named once (probe): 4 verdicts for 6 questions

## The counters the ratio can not see

A run that asks less scores better and serves the user worse. A run can pass both bars with most of its conversations unanswered, so read these counters with the ratio.

| Counter | Value |
|---|---|
| Questions asked | 118 |
| Questions that closed a slot | 95 |
| Invented by the model | 0 |
| Catalog-only conversations | 27 |
| Premature sessions | 0 |
| Linter findings | 0 |

## Where the bad questions came from

| Row | Bad questions |
|---|---|
| `commander_pick` | 6 |
| `colors` | 3 |
| `budget` | 2 |
| `pool_thin` | 2 |
| `theme` | 2 |
| `format` | 1 |
| `format_store` | 1 |
| `format_unsupported_open` | 1 |
| `named_card_role` | 1 |
| `pool` | 1 |
| `pool_precon` | 1 |
| `power_commander` | 1 |
| `power_sixty` | 1 |

## Faults

| Fault | Count |
|---|---|
| assumes an answer | 9 |
| duplicate | 5 |
| out of scope | 4 |
| inaccurate | 3 |
| omits information | 1 |

## What the catalog needs

| Action | Count |
|---|---|
| reword | 27 |
| remove | 9 |

## Every question the eval refused

The holdout conversations are left out of this list on purpose. 9 of their questions failed, and the fixer may not read which (D-134).

### 5. 60-card anything goes, turn 1, row `format`

**Asked:** Which format should the 60-card deck use: Standard or Modern?

Faults: assumes an answer, omits information. Catalog action: reword.

The user had already specified “a 60-card deck” where “anything goes at our table,” but the question restricts the choice to Standard or Modern and omits the unrestricted house format; reword it as: “Should I build this for your unrestricted house format, or for a named sanctioned format?”

### 7. FNM on Friday (probe), turn 2, row `colors`

**Asked:** Do you have a color preference for the deck?

Faults: duplicate. Catalog action: remove.

This repeats the earlier color-preference question from turn 1, which the user had not answered but which had already been asked.

### 7. FNM on Friday (probe), turn 2, row `budget`

**Asked:** What spending limit, if any, should I use for cards you do not own?

Faults: duplicate. Catalog action: remove.

This repeats the earlier budget question from turn 1; the user had not answered it yet, but repeating it in the next turn is still a duplicate.

### 22. extra turns commander, turn 2, row `pool_thin`

**Asked:** Since the card pool is thin for this theme, should I build around the cards you own and provide a buy list, or use the whole pool?

Faults: assumes an answer. Catalog action: reword.

The question asserts that “the card pool is thin” without evidence from the user's collection. Proposed wording: “Should I prioritize cards in your library and provide a buy list for missing cards, or use the full available card pool?”

### 23. land destruction, turn 2, row `pool_thin`

**Asked:** Your library is small for land destruction. Should I use owned cards first and make a buy list, or use the entire pool?

Faults: inaccurate. Catalog action: reword.

The user had a card collection, but never said that their library was small. Reword as: "Should I use cards you own first and create a buy list, or use the entire available card pool?"

### 31. FNM on Friday, Modern, turn 1, row `format_store`

**Asked:** Which format does Friday’s event run: Standard, Modern, or Commander?

Faults: duplicate. Catalog action: reword.

The user already specified the format: “FNM on Friday, Modern.” Proposed wording: “I see you’re building for Modern; do you have any color preference?”

### 43. a card name with a typo (probe), turn 2, row `named_card_role`

**Asked:** Should Karlov of the Ghost Counsel lead the deck, or should it be included among the other cards?

Faults: inaccurate. Catalog action: reword.

The question usefully distinguishes commander from the 99, but it repeats the apparently mistyped name “Karlov of the Ghost Counsel” without first confirming the intended card; reword it as: “Which Karlov card did you mean, and should it lead the deck or be included in the 99?”

### 53. the user swaps the commander (probe), turn 3, row `commander_pick`

**Asked:** Which commander do you want: Astarion, the Decadent, Karlov of the Ghost Council, or Sorin of House Markov // Sorin, Ravenous Neonate? Say "none" and I name three more.

Faults: out of scope. Catalog action: reword.

The user said “Actually use a different commander,” so Karlov of the Ghost Council is no longer an applicable choice. Proposed wording: “Which different commander do you want: Astarion, the Decadent, or Sorin of House Markov // Sorin, Ravenous Neonate? Say ‘none’ and I’ll name three more.”

### 58. terse: a card for the 99 (probe), turn 3, row `commander_pick`

**Asked:** Which commander do you want: Dina, Essence Brewer, Jarad, Golgari Lich Lord, or Dina, Soul Steeper? Say "none" and I name three more.

Faults: duplicate. Catalog action: reword.

The preceding turn already asked, “Which commander do you want? Name one, or I suggest three.” Because the user did not answer that question, this row repeats the same unresolved commander-selection slot. Reword it to: “You haven’t chosen a commander yet; would you like me to choose one for you?”

### 62. terse: a precon upgrade (probe), turn 2, row `pool_precon`

**Asked:** Should I start with the precon, use only cards from it, or ignore it for a fully optimized deck?

Faults: duplicate. Catalog action: remove.

“Upgrade my Atraxa, Praetors' Voice precon” already establishes that the precon is the starting point and that the goal is not to use only its original cards or ignore it.

### 73. terse: a colorless deck (probe), turn 1, row `commander_pick`

**Asked:** Which commander do you want: Emry, Lurker of the Loch, Urza, Lord High Artificer, or Padeem, Consul of Innovation? Say "none" and I name three more.

Faults: inaccurate. Catalog action: reword.

A commander choice was needed, but every offered option has blue color identity while the user requested "a colorless deck"; reword as "Which colorless-legal commander do you want, or should I suggest three options?"

### 82. terse: the user corrects a card name (probe), turn 1, row `colors`

**Asked:** Do you have a color preference?

Faults: assumes an answer. Catalog action: remove.

With Atraxa, Praetors' Voice as the named commander, the deck's color identity is fixed; a general color preference cannot change the commander's colors.

### 86. terse: enchantress (probe), turn 1, row `budget`

**Asked:** What spending limit, if any, should I use for cards you do not own?

Faults: assumes an answer. Catalog action: reword.

The user asked for a deck from their collection, but this question presumes that cards they do not own will be considered. Proposed wording: “Should I restrict the deck strictly to cards in your collection, or may I suggest cards to acquire?”

### 88. terse: no theme at all (probe), turn 1, row `theme`

**Asked:** Should the deck center on a creature type, a mechanic, or a play style? You can also leave this open for a surprise.

Catalog action: reword.

The user already said “Surprise me,” which leaves the theme choice to the agent; reword this conditional prompt as: “I’ll choose the theme—are there any theme constraints you want to impose?”

## Slots the eval says nobody asked about

These do not move the ratio. They are the guard against a quiet agent.

- 100. terse: a deck for a team event (probe): Whether the team event uses Unified Modern deck-construction rules, and what colors or cards are already reserved by the teammates.
- 102. terse: stax with no table named (probe): After the user said "None of those," ask them to describe the desired power level or provide a custom target so the deck's strength can be chosen.
- 103. terse: two themes at once (probe): Whether artifacts or lifegain should be the primary theme, or how evenly the deck should split between them.
- 104. terse: the user answers with a number (probe): After the user's response "2" to the commander question, ask them to name a commander or confirm that they want suggestions.
- 105. a theme the library can not lead (probe): After the user said "Colorless," ask for or offer a commander with colorless identity that fits an Eldrazi deck, or ask whether "colorless" permits a commander with colored identity.
- 107. a set name that names two base sets (probe): After the user answered the set, color, format, theme, budget, and bracket, the agent should have followed up because the commander question remained unanswered: ask whether to suggest a commander or have the user name one.
- 19. a big library, no theme: Ask how the agent can access the user's card library or request an inventory/export so the deck can be built from cards they already own.
- 2. blink with a thin library: The user's buy-list request leaves the previously unanswered spending limit unresolved; the agent should have followed up on the budget before building the acquisition list.
- 36. the user changes the theme late (probe): After the late theme change, ask whether the user wants to keep Meren of Clan Nel Toth and the black-green color identity, and what token strategy they want.
- 37. the user contradicts the budget (probe): The user's preferred strategy or whether the deck should be built to compete in the shop's aggro-heavy metagame.
- 42. shorthand and slang (probe): Ask which cards or relevant staples from the user's binder/collection they want to use or already own.
- 43. a card name with a typo (probe): card identity clarification: confirm the intended card name before building around “Karlov of the Ghost Counsel”
- 44. a card that is not legal (probe): The agent should have addressed that Black Lotus is not legal in Modern and asked whether to switch formats or replace the centerpiece with a Modern-legal card.
- 48. sideboard help only (probe): Which matchups or opposing decks do you most want the sideboard to address?, How many sideboard slots are available, and are there any cards already in the sideboard?
- 52. a price cap per card (probe): Total deck budget: the user gave only a per-card cap, not how much the complete deck may cost.
- 56. the user answers a different question (probe): The available library inventory or how the agent should access it, since the user asked to build from their library first and no specific cards or collection contents were provided.
- 63. terse: a competitive request (probe): budget: What is your budget for the deck?, optimization target: Should the deck prioritize the strongest overall Modern list or be tuned specifically to fight your shop’s aggro and combo decks?
- 66. terse: sideboard help only (probe): The current sideboard and the matchups or threats the user wants help addressing; knowing only that the deck is Modern mono-red burn is not enough to recommend specific sideboard cards.
- 7. FNM on Friday (probe): After the user specified Pioneer, the agent should have clarified whether they wanted help adapting to a supported format or a Pioneer-specific recommendation, if Pioneer is unsupported., The agent should have resolved the still-unanswered color preference before proceeding.
- 74. terse: a card the user does not own (probe): After the user said “Buy it then,” clarify whether they are changing the “only the cards I own” constraint and, if so, what budget or purchasing scope applies.
- 79. terse: everything buried in one long message (probe): Ask which black-green cards or relevant cards in the user's existing collection they want to build around or have available.
- 8. mill for a playgroup: The stated blue-and-black color preference conflicts with Bruvac the Grandiloquent's mono-blue color identity; ask whether to keep Bruvac and drop black, or choose a commander that permits both colors.
- 94. terse: a collection of lands only (probe): The commander choice remained unanswered after the user gave only the power bracket and budget; ask them to choose among the three listed commanders or request three more.
- 95. terse: a library that does not exist (probe): card collection/library access: ask the user to provide or connect the library, since the app knows user_has_a_card_collection is false
- 96. terse: the user asks about a card (probe): Answer whether Sol Ring is legal in Modern.

## Run

- Suite `question-eval`, run `pr7-question-eval-run35`, on 2026-09-04, commit `e8e7dbc`.
- Roles: eval on `gpt-5.6-luna` (openai, effort low).
- Versions: eval prompt version 3, gate_document `pr7-question-gate-run35`.
- Calls: 103. Cost: $0.0959. Time: 805 seconds.
- Tokens: 114402 input (0 cached), 60824 output.
- Eval cost: $0.0959.

## The card check

The check ran against the local snapshot. No refusal rested on a card name the snapshot refutes.
