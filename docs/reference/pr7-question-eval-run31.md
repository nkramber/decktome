# PR-7 question eval

Run: `pr7-question-gate-run31`. Eval model: `gpt-5.6-luna`. Prompt version 3.

Scored 413 questions. 27 were not warranted, and 0 are unsure.

**Bad-question ratio: 7.9% on the holdout.**

The tune split holds 286 questions at 5.9%, and the holdout holds 127 at 7.9%. A fixer reads the tune split alone. A gain that shows only there is a reworded test set (D-134).

## The counters the ratio can not see

A run that asks less scores better and serves the user worse. A run can pass both bars with most of its conversations unanswered, so read these counters with the ratio.

| Counter | Value |
|---|---|
| Questions asked | 115 |
| Questions that closed a slot | 93 |
| Invented by the model | 0 |
| Catalog-only conversations | 27 |
| Premature sessions | 0 |
| Linter findings | 0 |

## Where the bad questions came from

| Row | Bad questions |
|---|---|
| `budget` | 6 |
| `colors` | 4 |
| `format_unsupported_open` | 3 |
| `power_commander` | 3 |
| `commander_pick` | 2 |
| `format` | 2 |
| `pool_thin` | 2 |
| `theme` | 2 |
| `commander` | 1 |
| `named_card_role` | 1 |
| `power_sixty` | 1 |

## Faults

| Fault | Count |
|---|---|
| assumes an answer | 9 |
| inaccurate | 6 |
| out of scope | 5 |
| duplicate | 4 |
| omits information | 1 |
| two questions in one | 1 |
| vague | 1 |

## What the catalog needs

| Action | Count |
|---|---|
| reword | 37 |
| remove | 10 |
| no | 1 |

## Every question the eval refused

The holdout conversations are left out of this list on purpose. 10 of their questions failed, and the fixer may not read which (D-134).

### 1. lifegain with a collection, turn 1, row `budget`

**Asked:** Do you have a budget for cards to buy?

Faults: assumes an answer. Catalog action: reword.

The question presumes the user has a budget for cards to buy, although the user only said, “I have a collection.” Proposed wording: “Would you like the deck to use only cards from your collection, or may I suggest cards to buy as well?”

### 5. 60-card anything goes, turn 1, row `format`

**Asked:** Which format should the 60-card deck use: Commander, Standard, or Modern?

Faults: inaccurate, omits information. Catalog action: reword.

A format question was needed, but the options omit the user's stated anything-goes house format and include Commander for a 60-card deck, even though Commander uses a 100-card deck.

### 13. a card to keep that is not the commander, turn 3, row `budget`

**Asked:** Do you have a budget for cards to buy?

Faults: assumes an answer. Catalog action: reword.

The question presumes that a budget is relevant or exists, although the user has not mentioned one. Reword it as: “What spending limit, if any, should I use for cards not in your library?”

### 20. dinosaur tribal for a child, turn 3, row `pool_thin`

**Asked:** There are only 4 Dinosaur cards in your pool, so how should I handle the card pool: start with what you own and add a buy list, or use every card in the pool?

Faults: assumes an answer, inaccurate. Catalog action: reword.

The transcript establishes only that the user has a card collection, not that there are exactly four Dinosaur cards in the pool. Reword as: “How should I use your card collection: prioritize cards you own and suggest additions, or build without restricting myself to your collection?”

### 26. legacy for an event (probe), turn 1, row `format_unsupported_open`

**Asked:** I do not build Legacy. Which format should I build instead: Commander, Standard, or Modern?

Faults: assumes an answer. Catalog action: reword.

No—the user already specified “A Legacy deck,” but this asks them to choose a different format and assumes they want to switch. Reword it to: “I can’t build a Legacy deck. Would you like to stop, or should I adapt the request to Commander, Standard, or Modern?”

### 35. the user changes the format (probe), turn 3, row `budget`

**Asked:** Is there a budget for cards to buy?

Faults: duplicate. Catalog action: remove.

The agent had already asked “Do you have a budget for cards to buy?” on turn 1, and the user had not changed or withdrawn that request; asking “Is there a budget for cards to buy?” repeated the same question.

### 43. a card name with a typo (probe), turn 2, row `named_card_role`

**Asked:** Do you want Karlov of the Ghost Counsel as your commander, or as one card in the 99?

Faults: inaccurate. Catalog action: reword.

The question asks about “Karlov of the Ghost Counsel,” but the user's card name contains a typo; reword it as “Do you want Karlov of the Ghost Council as your commander, or as one card in the 99?”

### 49. copy a known deck (probe), turn 1, row `colors`

**Asked:** Do you have a color preference?

Catalog action: reword.

The user asked for the best deck "whatever is winning," which already indicates that color preference should not constrain the recommendation. Reword this row to ask about color preference only if the user is unwilling to play the current best colors.

### 79. terse: everything buried in one long message (probe), turn 1, row `power_commander`

**Asked:** Which power bracket should the deck target?

Faults: duplicate. Catalog action: remove.

No; the user already said they wanted bracket 2 or 3: "maybe bracket 2 or 3".

### 82. terse: the user corrects a card name (probe), turn 1, row `format`

**Asked:** Which format should the deck use: Commander, Standard, or Modern?

Faults: inaccurate. Catalog action: reword.

The user named Atraxa, Praetor's Voice, which strongly indicates Commander, while Standard and Modern are not applicable choices for this commander. Proposed wording: “Should I build this as a Commander deck?”

### 82. terse: the user corrects a card name (probe), turn 1, row `colors`

**Asked:** Do you have a color preference?

Faults: assumes an answer. Catalog action: remove.

Atraxa, Praetor's Voice fixes the deck's color identity, so the question assumes the user can or needs to choose colors independently.

### 88. terse: no theme at all (probe), turn 1, row `theme`

**Asked:** What would you like the deck to center on: a creature type, a mechanic, or a play style?

Catalog action: reword.

The user said “Surprise me,” which delegates the theme choice to the agent; reword as “I’ll choose a theme for the deck unless you want to specify one.”

### 88. terse: no theme at all (probe), turn 1, row `power_commander`

**Asked:** Which power bracket should the deck target?

Catalog action: reword.

The user said “Surprise me,” so the agent should choose the power target rather than ask the user to select a bracket; reword as “I’ll choose an appropriate power level for the deck unless you want to specify one.”

### 88. terse: no theme at all (probe), turn 1, row `colors`

**Asked:** Do you have a color preference?

Catalog action: reword.

The user said “Surprise me,” which also delegates the color choice; reword as “I’ll choose the deck’s colors unless you have a preference.”

### 91. terse: a budget of zero (probe), turn 2, row `budget`

**Asked:** Do you have a budget for cards to buy?

Faults: duplicate. Catalog action: reword.

The user already said, “I can not spend anything,” so asking for a budget repeats information; reword as “Since you said you cannot spend anything, should I restrict the deck to cards you already own?”

### 95. terse: a library that does not exist (probe), turn 2, row `budget`

**Asked:** Do you have a budget for cards to buy?

Faults: assumes an answer. Catalog action: reword.

The user asked to build from their library and never mentioned buying cards; reword as: “If you want to buy cards beyond your library, what budget should I use?”

### 107. a set name that names two base sets (probe), turn 2, row `budget`

**Asked:** Do you have a budget for cards to buy?

Faults: assumes an answer. Catalog action: no.

The question presumes the user plans to buy cards, but the user had not said that; a deck can instead be built from owned cards or without a purchase plan.

## Slots the eval says nobody asked about

These do not move the ratio. They are the guard against a quiet agent.

- 1. lifegain with a collection: Ask which specific cards from the user's library/collection are available, or request an inventory/list, because “build from my library first” requires knowing the collection contents., Ask whether the $50 buy-list limit is a hard cap and whether it includes shipping/taxes, if that distinction affects purchasing recommendations.
- 100. terse: a deck for a team event (probe): deck strategy or team role, such as whether the user wants a control deck to complement teammates' aggro and combo decks
- 105. a theme the library can not lead (probe): A commander-selection question offering only legal, colorless Eldrazi commanders from the user's library, or asking whether the agent may choose one if no such commander is available.
- 107. a set name that names two base sets (probe): The agent did not ask which Tarkir: Dragonstorm cards, if any, the user already owns or wants to use, despite the collection-status context being relevant to deck construction., The agent did not clarify whether the user wants a Dragonstorm-only card pool or merely a Tarkir: Dragonstorm theme, after the set was resolved.
- 16. the strongest modern deck: Ask what budget the user means, since “under budget” does not specify an amount.
- 17. precon upgrade at bracket 2: Ask which cards are in the user's library/collection so the upgrade can be built from it first.
- 18. pauper burn on a budget (probe): color identity
- 19. a big library, no theme: Ask how the user's library is available to the app (for example, whether there is an inventory to build from), since the user requested building from their library first.
- 23. land destruction: Clarify the budget for cards to buy, since “My table does not mind” does not answer the budget question.
- 26. legacy for an event (probe): Whether the user already owns any cards or needs the entire Legacy deck purchased, since the app knows the user has no card collection.
- 29. two hundred dollars: Ask which cards or relevant portions of the user's library are available, so the deck can be built from that collection first.
- 34. no ban list at our table: A follow-up asking for the deck's color preference, since the user described the strategy but did not yet specify colors.
- 36. the user changes the theme late (probe): After the user changed the theme, ask whether the deck should remain black-green Commander with Meren of Clan Nel Toth or whether the colors and commander should also change., Ask what kind of token strategy the user wants, such as creature tokens, sacrifice-oriented tokens, or another token theme.
- 40. the user asks a question back (probe): Explain what a Commander power bracket means in response to the user's question.
- 43. a card name with a typo (probe): Clarification of the intended card name, since “Karlov of the Ghost Counsel” contains a typo and should be confirmed as Karlov of the Ghost Council before proceeding.
- 44. a card that is not legal (probe): Flag that Black Lotus is not legal in Modern and ask whether to replace it with a legal card or change the format.
- 48. sideboard help only (probe): Ask for the current sideboard/decklist or the cards the user wants help replacing, since sideboard recommendations depend on the existing configuration.
- 56. the user answers a different question (probe): Ask how the deck should address the stated social constraint that the user's friends hate mill., Ask the user to provide their library or a card/deck inventory so the build can be made from their collection first.
- 63. terse: a competitive request (probe): The user's budget amount after saying "The best deck under budget."
- 66. terse: sideboard help only (probe): Request the current maindeck and existing sideboard so recommendations can fit the user's actual Modern Burn list., Ask which control decks or colors are common at the user's shop, once the user says the shop plays a lot of control.
- 75. terse: cEDH at bracket 5 (probe): The agent should have asked which bracket-5 power-level expectations or constraints the user wants the deck to meet, since “cEDH” alone does not specify the requested bracket context.
- 76. terse: a companion (probe): Modern legality: ask whether the user wants a currently Modern-legal deck or a house-rule build allowing Lurrus of the Dream-Den as companion, since Lurrus is not currently legal in Modern.
- 85. terse: chaos (probe): commander choice
- 89. terse: teaching a new player (probe): Whether the deck should prioritize especially simple, beginner-friendly gameplay and low card complexity.
- 92. terse: proxy the expensive cards only (probe): Clarify whether Bruvac the Grandiloquent is the commander or a card in the 99, and resolve the blue-black color identity if Bruvac is intended as commander., Clarify whether the $20 threshold applies to purchased cards only or to the total deck value; the user did not answer that question.
- 95. terse: a library that does not exist (probe): Clarify how to proceed because the user requested a deck from their library, but user_has_a_card_collection is false; ask whether they want a deck built from cards to buy instead.
- 96. terse: the user asks about a card (probe): Answer whether Sol Ring is legal in Modern.
- 97. terse: Pauper Commander (probe): The intended power level or competitiveness of the deck., The user's preferred commander, or whether the agent should choose one, once a supported Commander format is selected.
- 99. terse: Historic on Arena (probe): The agent should have clarified whether the user wanted an Arena-supported format or a paper Modern deck, since Modern is not an Arena format.

## Run

- Calls: 103. Time: 675.4 seconds.
- Tokens: 110316 input (0 cached), 59674 output.
- Cost: $0.0937.

## The card check

The check ran against the local snapshot. No refusal rested on a card name the snapshot refutes.
