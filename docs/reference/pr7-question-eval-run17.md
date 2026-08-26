# PR-7 question eval

Run: `pr7-question-gate-run17`. Eval model: `gpt-5.6-luna`. Prompt version 3.

Scored 378 questions. 49 were not warranted, and 0 are unsure.

**Bad-question ratio: 13.3% on the holdout.**

The tune split holds 250 questions at 12.8%, and the holdout holds 128 at 13.3%. A fixer reads the tune split alone. A gain that shows only there is a reworded test set (D-134).

## The counters the ratio can not see

A run that asks less scores better and serves the user worse. Gate run 7 of 2026-08-25 passed both bars with 26 of 30 conversations unanswered.

| Counter | Value |
|---|---|
| Questions asked | 125 |
| Questions that closed a slot | 97 |
| Invented by the model | 3 |
| Catalog-only conversations | 27 |
| Premature sessions | 0 |
| Linter findings | 0 |

## Where the bad questions came from

| Row | Bad questions |
|---|---|
| `format_unsupported_open` | 14 |
| `commander_pick` | 8 |
| `power_sixty_confirm` | 7 |
| `locked` | 4 |
| `colors` | 3 |
| `power_commander` | 3 |
| `meta` | 2 |
| `format` | 1 |
| `format_unsupported` | 1 |
| `named_card_role` | 1 |
| `plan_choice` | 1 |
| `power_sixty` | 1 |
| `theme` | 1 |
| `theme_card_named` | 1 |
| `theme_competitive` | 1 |

## Faults

| Fault | Count |
|---|---|
| duplicate | 31 |
| assumes an answer | 9 |
| out of scope | 6 |
| inaccurate | 3 |
| omits information | 2 |
| vague | 1 |

## What the catalog needs

| Action | Count |
|---|---|
| reword | 43 |
| remove | 23 |
| duplicate | 1 |

## Every question the eval refused

The holdout conversations are left out of this list on purpose. 17 of their questions failed, and the fixer may not read which (D-134).

### 1. lifegain with a collection, turn 3, row `commander_pick`

**Asked:** Which one do you want: Cecil, Dark Knight // Cecil, Redeemed Paladin, Denethor, Ruling Steward, or Cloak and Dagger, Entwined? Say "none" and I name three more.

Faults: duplicate. Catalog action: remove.

The exact same commander-selection question was already asked on turn 2, and the user instead requested a commander suggestion from their library in turn 3.

### 6. the strongest deck, no collection, turn 2, row `power_sixty_confirm`

**Asked:** Should I aim for tournament-level power, or would you prefer FNM-level or casual power?

Faults: duplicate. Catalog action: remove.

The user had already requested “the strongest deck possible” and then “The best deck,” which supplies the intended tournament-level power target; this confirmation was unnecessary.

### 7. FNM on Friday (probe), turn 3, row `format_unsupported_open`

**Asked:** I do not build Pioneer. Which format should I build instead: Commander, Standard, or Modern?

Faults: duplicate. Catalog action: remove.

This repeats the immediately preceding format question verbatim after the user already said “200 dollars, tournament level. My shop is mostly aggro decks.”

### 13. a card to keep that is not the commander, turn 1, row `locked`

**Asked:** Should Sanguine Bond stay in the deck, or may I cut cards that do not fit the plan?

Faults: duplicate. Catalog action: reword.

The user already said “keep Sanguine Bond in it,” so asking whether Sanguine Bond may be cut repeats an answer already provided. Reword this row to ask only about other named cards that were not explicitly locked in, or skip it when all named cards are already addressed.

### 18. pauper burn on a budget (probe), turn 2, row `format_unsupported_open`

**Asked:** I do not build Pauper. Which format should I build instead: Commander, Standard, or Modern?

Faults: duplicate. Catalog action: reword.

The agent already asked this exact format-substitution question on turn 1, and the user did not select one of the offered formats; re-ask only with an acknowledgment that Pauper remains unsupported and ask whether they want Commander, Standard, or Modern instead.

### 18. pauper burn on a budget (probe), turn 3, row `format_unsupported_open`

**Asked:** I do not build Pauper. Which format should I build instead: Commander, Standard, or Modern?

Faults: duplicate. Catalog action: reword.

This repeats the same unanswered question from turns 1 and 2, despite the user adding that the deck should be casual power with a $25 cap; re-ask only if needed and acknowledge those constraints.

### 21. vintage with proxies (probe), turn 2, row `format_unsupported_open`

**Asked:** I do not build Vintage. Which format should I build instead: Commander, Standard, or Modern?

Faults: assumes an answer. Catalog action: reword.

The user explicitly requested “Any card, no ban list. Call it Vintage,” so asking which other format to use presumes the user will abandon the requested format. If Vintage is unsupported, reword as: “Vintage is not currently supported. Would you like me to build this as Commander, Standard, or Modern instead?”

### 21. vintage with proxies (probe), turn 3, row `format_unsupported_open`

**Asked:** I do not build Vintage. Which format should I build instead: Commander, Standard, or Modern?

Faults: duplicate. Catalog action: remove.

The agent had already asked on turn 2 which supported format to use, and the user had already stated “Call it Vintage”; repeating the same unsupported-format question added no new information.

### 25. poison in a tournament, turn 1, row `power_sixty_confirm`

**Asked:** Should I build this for tournament-level competition?

Faults: duplicate. Catalog action: remove.

The user already specified “in a Modern tournament,” so asking “Should I build this for tournament-level competition?” requests information already provided.

### 27. a card for the 99, turn 2, row `locked`

**Asked:** Should Grist, the Hunger Tide be included in the 99, or may the deck omit Grist entirely?

Catalog action: remove.

“Build around Grist, the Hunger Tide” already indicates that Grist should be in the deck, so asking whether to include Grist or omit it was unnecessary at this point.

### 33. a tempo deck for an event, turn 1, row `theme_competitive`

**Asked:** For the event, do you want a named top-tier deck or the best deck under your budget?

Faults: assumes an answer. Catalog action: reword.

The phrase “the best deck under your budget” presumes the user has already established a budget, but the user only said “A Modern deck for an event.” Reword as: “Do you want a named top-tier deck, or should I optimize for a budget you specify?”

### 43. a card name with a typo (probe), turn 2, row `named_card_role`

**Asked:** Should Karlov of the Ghost Council lead the deck, or be included as one of the other cards?

Faults: assumes an answer. Catalog action: reword.

The question uses “Karlov of the Ghost Council” as though it is certainly the card intended, even though the user wrote “Karlov of the Ghost Counsel”; reword it as “Should the Karlov card you mean lead the deck, or be included in the 99?” and separately confirm the card name.

### 49. copy a known deck (probe), turn 1, row `power_sixty_confirm`

**Asked:** Should I build for tournament-level competition, FNM-level play, or casual play?

Catalog action: reword.

The request for the “current best Standard deck, whatever is winning” already indicates a tournament-oriented power goal; reword as “You asked for the current strongest deck—should I treat that as tournament-level, or would you prefer FNM or casual power?”

### 54. two decks at once (probe), turn 3, row `commander_pick`

**Asked:** Which one do you want: Atarka, World Render, Thrakkus the Butcher, or Dragonlord Atarka? Say "none" and I name three more.

Faults: duplicate. Catalog action: duplicate.

The same commander-choice question was already asked on turn 2, and the user did not provide a commander choice in “Bracket 3, and 150 dollars.” Repeating it is still unnecessary rather than a new information need.

### 61. terse: a proxy table (probe), turn 2, row `format_unsupported_open`

**Asked:** I do not build Vintage. Which format should I build instead: Commander, Standard, or Modern?

Faults: duplicate. Catalog action: remove.

This repeats the unanswered format-selection question from turn 1 verbatim.

### 61. terse: a proxy table (probe), turn 3, row `format_unsupported_open`

**Asked:** I do not build Vintage. Which format should I build instead: Commander, Standard, or Modern?

Faults: duplicate. Catalog action: remove.

This repeats the same unanswered format-selection question asked on turns 1 and 2.

### 63. terse: a competitive request (probe), turn 1, row `power_sixty_confirm`

**Asked:** Should I build the deck for tournament-level competition?

Catalog action: reword.

The user already asked for “the strongest Modern deck,” which establishes tournament-level optimization; ask this only when the requested power level is unclear.

### 69. terse: a Pauper deck named once (probe), turn 1, row `format_unsupported_open`

**Asked:** I do not build Pauper. Which format should I build instead: Commander, Standard, or Modern?

Faults: out of scope, inaccurate. Catalog action: remove.

The user already specified “A Pauper burn deck” and “Mono red,” so asking them to replace Pauper with another format is unnecessary; Pauper is an applicable Magic format for the requested deck.

### 69. terse: a Pauper deck named once (probe), turn 2, row `format_unsupported_open`

**Asked:** I do not build Pauper. Which format should I build instead: Commander, Standard, or Modern?

Faults: duplicate, out of scope, inaccurate. Catalog action: remove.

This repeats the same request after the user has already specified “A Pauper burn deck,” and the offered replacement formats still do not address the requested Pauper deck.

### 73. terse: a colorless deck (probe), turn 1, row `colors`

**Asked:** What color preference should the deck have?

Faults: duplicate. Catalog action: remove.

The user had already specified the deck was colorless: “A colorless Commander deck built around big artifacts.”

### 75. terse: cEDH at bracket 5 (probe), turn 1, row `power_commander`

**Asked:** Which power bracket should the deck target?

Faults: duplicate. Catalog action: reword.

The user already said “A cEDH deck,” which answers the power-bracket question. Reword the row to ask about the target power bracket only when the user has not already specified a power level.

### 88. terse: no theme at all (probe), turn 1, row `theme`

**Asked:** What should the deck focus on: a creature type, a mechanic, or a play style?

Catalog action: reword.

The user said “Surprise me,” which delegates the choice of theme to the agent, so this preference question was not needed at this point. Proposed wording: “I’ll choose a theme unless you want to constrain it.”

### 88. terse: no theme at all (probe), turn 1, row `power_commander`

**Asked:** Which power bracket should the deck target? 2 is precon level, 3 is upgraded, and 4 is high power.

Catalog action: reword.

The user said “Surprise me,” which delegates the power target to the agent, so asking the user to select a bracket was not needed at this point. Proposed wording: “I’ll choose an appropriate power level unless you want to specify one.”

### 88. terse: no theme at all (probe), turn 1, row `colors`

**Asked:** Do you have a color preference?

Catalog action: reword.

The user said “Surprise me,” which delegates the color choice to the agent, so asking for a color preference was not needed at this point. Proposed wording: “I’ll choose the deck’s colors unless you want to restrict them.”

### 90. terse: a deck for a spouse (probe), turn 3, row `commander_pick`

**Asked:** Which one do you want: Arahbo, Roar of the World, Kaheera, the Orphanguard, or Kutzil, Malamet Exemplar? Say "none" and I name three more.

Faults: duplicate. Catalog action: remove.

This repeats the unanswered commander-selection question from turn 2 verbatim, without adding information or changing what the user needs to decide.

### 93. terse: two budget shapes at once (probe), turn 3, row `meta`

**Asked:** What decks do you expect to face at FNM? I’ll tune the 15 sideboard cards to that.

Faults: assumes an answer. Catalog action: reword.

The user said only “FNM level,” which describes desired strength and does not establish that they will attend an FNM or what opponents they expect. Reword as: “Which opposing decks, if any, should the 15-card sideboard prioritize?”

### 96. terse: the user asks about a card (probe), turn 1, row `theme_card_named`

**Asked:** What should the deck focus on: a creature type, a mechanic, or a play style?

Faults: assumes an answer, out of scope. Catalog action: reword.

The user only asked, "Is Sol Ring legal in Modern?"; this question assumes they want to build a deck and asks about its theme before answering the legality question. Proposed wording: "If you want help building a deck, what should it focus on: a creature type, a mechanic, or a play style?"

### 96. terse: the user asks about a card (probe), turn 1, row `power_sixty`

**Asked:** How strong should the deck be?

Faults: assumes an answer, out of scope. Catalog action: reword.

The user asked only about Sol Ring's Modern legality, but this asks how strong an unrequested deck should be. Proposed wording: "If you want help building a deck, how strong should it be?"

### 96. terse: the user asks about a card (probe), turn 1, row `colors`

**Asked:** Do you have a color preference?

Faults: assumes an answer, out of scope. Catalog action: reword.

The user did not say they wanted a deck, so asking for a color preference presumes a deck-building request rather than addressing "Is Sol Ring legal in Modern?" Proposed wording: "If you want help building a deck, do you have a color preference?"

### 97. terse: Pauper Commander (probe), turn 2, row `format_unsupported_open`

**Asked:** I do not build Pauper Commander. Which format should I build instead: Commander, Standard, or Modern?

Faults: duplicate. Catalog action: remove.

This repeats the unanswered format question from turn 1 verbatim; the user had not answered it, but repeating it on the next turn does not add a new question.

### 97. terse: Pauper Commander (probe), turn 3, row `format_unsupported_open`

**Asked:** I do not build Pauper Commander. Which format should I build instead: Commander, Standard, or Modern?

Faults: duplicate. Catalog action: remove.

This repeats the same format question from turns 1 and 2, while the user still had not answered it; the row should not be emitted again.

### 99. terse: Historic on Arena (probe), turn 1, row `format_unsupported_open`

**Asked:** I do not build Historic. Which format should I build instead: Commander, Standard, or Modern?

Faults: inaccurate. Catalog action: reword.

The user specified Arena, but Commander and Modern are not Arena formats; reword as: "I can’t build Historic on Arena. Which Arena format should I use instead: Standard, Explorer, or Alchemy?"

## Slots the eval says nobody asked about

These do not move the ratio. They are the guard against a quiet agent.

- 101. terse: mill with no table named (probe): budget
- 102. terse: stax with no table named (probe): budget
- 103. terse: two themes at once (probe): budget
- 15. standard at the store, no library: Budget or spending limit for the Standard deck.
- 25. poison in a tournament: budget
- 27. a card for the 99: commander: The user never clearly answered the turn 1 question about whether they had a commander in mind or wanted a suggestion; after “Commander. A sacrifice deck, black and green,” the agent should have clarified or offered to choose a legal commander.
- 30. another version, same plan: Whether the replacement cards must come from the user's collection, since the user asked for different cards and the app knows they have a collection.
- 31. FNM on Friday, Modern: budget: The user requested “the best deck under budget” but no budget amount was provided; ask what the budget is.
- 33. a tempo deck for an event: An explicit budget amount was not asked for before proposing a budget-based deck choice.
- 36. the user changes the theme late (probe): After the theme change, ask whether to keep Commander, black-green, and Meren of Clan Nel Toth, or choose a new commander/color identity., Clarify what kind of token deck the user wants (token creatures, sacrifice tokens, a specific token type, or another token-focused plan).
- 43. a card name with a typo (probe): A confirmation of the intended card name before referring to “Karlov of the Ghost Council,” since the user wrote “Karlov of the Ghost Counsel.”
- 44. a card that is not legal (probe): Acknowledge that Black Lotus is not legal in Modern and confirm or propose a legal replacement before building the deck.
- 48. sideboard help only (probe): Ask for the current Modern burn decklist and existing sideboard so the 15-card sideboard can be tuned without duplicating or conflicting with the user's cards., Ask about budget or card availability, since the user has not said which sideboard cards they can obtain.
- 49. copy a known deck (probe): budget
- 54. two decks at once (probe): Ask for the user's budget currency and whether the $150 budget applies to the Commander deck only; also clarify whether the user wants to proceed with the Commander deck before building the Modern deck.
- 55. a returning player (probe): budget
- 56. the user answers a different question (probe): Whether the user still wants a mill deck despite saying their friends hate mill.
- 58. terse: a card for the 99 (probe): Ask which commander to suggest or present a suitable black-green commander, since the user said, “You pick the commander.”
- 59. terse: two decks at once (probe): commander choice or permission to suggest one after the user did not answer that question, the remaining Modern deck's theme, colors, and budget after prioritizing the Commander deck
- 64. terse: a deck as a gift (probe): budget
- 66. terse: sideboard help only (probe): Whether the sideboard must be built from cards the user already owns, or whether they want purchase recommendations, including any budget constraint.
- 7. FNM on Friday (probe): A color preference or color constraints remained unanswered after the user provided the budget and competitive goals.
- 71. terse: one-word answers (probe): A question about the user's budget or card-acquisition constraints, since no budget was provided and the user has no stated collection.
- 72. terse: a five-color deck (probe): budget
- 73. terse: a colorless deck (probe): budget
- 74. terse: a card the user does not own (probe): Clarify whether the deck must remain limited to the user's current collection or whether buying additional cards is now allowed, since the user first said "only the cards I own" and later said "Buy it then."
- 75. terse: cEDH at bracket 5 (probe): budget
- 77. terse: a Background pair (probe): Offer Background commander pairs whose color identity matches the user's stated red-and-white preference, and ask the user to choose among them.
- 8. mill for a playgroup: Whether the playgroup is comfortable with a mill strategy or has any relevant table restrictions., Whether mill should be the primary win condition or merely a theme.
- 82. terse: the user corrects a card name (probe): budget
- 89. terse: teaching a new player (probe): budget
- 91. terse: a budget of zero (probe): Ask which cards from the user's collection are available for a no-cost deck, or whether proxies are acceptable, since the user said they cannot spend anything.
- 95. terse: a library that does not exist (probe): Ask the user to provide or import their card collection, since the app does not have one and cannot build a deck from an unavailable library.
- 96. terse: the user asks about a card (probe): Answer whether Sol Ring is legal in Modern before beginning deck-building intake., Do not infer that the user wants to build a deck from a card-legality question.
- 98. terse: Oathbreaker (probe): budget

## Run

- Calls: 102. Time: 789.1 seconds.
- Tokens: 111664 input (0 cached), 61285 output.
- Cost: $0.0959.

## The card check

The check ran against the local snapshot. No refusal rested on a card name the snapshot refutes.
