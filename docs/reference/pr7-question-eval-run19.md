# PR-7 question eval

Run: `pr7-question-gate-run19`. Eval model: `gpt-5.6-luna`. Prompt version 3.

Scored 423 questions. 26 were not warranted, and 0 are unsure.

**Bad-question ratio: 8.9% on the holdout.**

The tune split holds 288 questions at 4.9%, and the holdout holds 135 at 8.9%. A fixer reads the tune split alone. A gain that shows only there is a reworded test set (D-134).

## The counters the ratio can not see

A run that asks less scores better and serves the user worse. Gate run 7 of 2026-08-25 passed both bars with 26 of 30 conversations unanswered.

| Counter | Value |
|---|---|
| Questions asked | 132 |
| Questions that closed a slot | 100 |
| Invented by the model | 1 |
| Catalog-only conversations | 29 |
| Premature sessions | 0 |
| Linter findings | 0 |

## Where the bad questions came from

| Row | Bad questions |
|---|---|
| `colors` | 3 |
| `format` | 3 |
| `format_unsupported_open` | 3 |
| `meta` | 3 |
| `budget` | 2 |
| `commander` | 2 |
| `pool_thin` | 2 |
| `power_commander` | 2 |
| `theme_card_named` | 2 |
| `house_format_limits` | 1 |
| `named_card_role` | 1 |
| `pool` | 1 |
| `power_sixty` | 1 |

## Faults

| Fault | Count |
|---|---|
| assumes an answer | 10 |
| inaccurate | 6 |
| out of scope | 6 |
| two questions in one | 3 |
| jargon | 2 |
| vague | 2 |
| duplicate | 1 |

## What the catalog needs

| Action | Count |
|---|---|
| reword | 38 |
| remove | 5 |

## Every question the eval refused

The holdout conversations are left out of this list on purpose. 12 of their questions failed, and the fixer may not read which (D-134).

### 5. 60-card anything goes, turn 1, row `format`

**Asked:** Which format would you like for the 60-card deck?

Faults: inaccurate. Catalog action: reword.

The user already specified a 60-card deck and an anything-goes table, so offering Commander alongside Standard and Modern is not an appropriate set of options for this request; ask whether they want a sanctioned 60-card format or a custom house format.

### 16. the strongest modern deck, turn 1, row `meta`

**Asked:** Should I keep the sideboard general, or will you name the decks you expect to face?

Faults: assumes an answer. Catalog action: reword.

The question is relevant, but it presumes the user has a defined set of expected opposing decks despite no event or metagame being given. Reword as: “Should I build a generally tuned sideboard, or tailor it to a specific metagame you describe?”

### 20. dinosaur tribal for a child, turn 3, row `pool_thin`

**Asked:** You have four dinosaur cards, but the deck needs at least 30. Should I prioritize cards you own and add a buy list, or use the entire pool?

Faults: assumes an answer, inaccurate, vague. Catalog action: reword.

The prior transcript never says the user has four dinosaur cards, and a Commander deck does not need at least 30 dinosaur cards; “the entire pool” is also undefined. Reword as: “Should I prioritize cards from your library and identify only the additional cards you may need, or may I recommend cards outside your library as well?”

### 26. legacy for an event (probe), turn 1, row `format_unsupported_open`

**Asked:** I do not build Legacy. Which format should I build instead: Commander, Standard, or Modern?

Faults: assumes an answer. Catalog action: reword.

The user explicitly requested Legacy, so the question presumes they will accept switching formats. Reword it as: “Legacy isn’t supported. Would you like to switch to Commander, Standard, or Modern?”

### 43. a card name with a typo (probe), turn 1, row `format`

**Asked:** Which format would you like for the deck: Commander, Standard, or Modern?

Faults: inaccurate. Catalog action: reword.

The format is needed, but Standard and Modern do not apply to Karlov of the Ghost Council as the deck's build-around card; ask which Commander variant or other applicable format the user wants.

### 49. copy a known deck (probe), turn 1, row `colors`

**Asked:** Do you have a color preference?

Catalog action: remove.

The user already said “whatever is winning,” which answers the color-preference question by declining to constrain the colors.

### 53. the user swaps the commander (probe), turn 2, row `budget`

**Asked:** Is there a budget for cards to buy?

Faults: out of scope. Catalog action: reword.

The user asked for a deck "from my library," so asking about a budget for cards to buy introduces purchases the user did not request. Reword as: "If you want to buy cards beyond your library, what budget should I use?"

### 61. terse: a proxy table (probe), turn 1, row `format_unsupported_open`

**Asked:** I do not build Vintage. Which format should I build instead: Commander, Standard, or Modern?

Faults: inaccurate. Catalog action: remove.

No—the user explicitly requested “A Vintage artifact prison deck,” but the question incorrectly claims “I do not build Vintage” and offers unrelated replacement formats.

### 61. terse: a proxy table (probe), turn 2, row `house_format_limits`

**Asked:** Inside your house format, do the normal 60-card deck limits hold?

Faults: assumes an answer, vague. Catalog action: reword.

No—the wording presumes an established “house format,” although the user only said “We proxy everything at our table.” Reword as: “Should I use normal Vintage deck-construction rules, or does your table use different deck-size or copy-limit rules?”

### 74. terse: a card the user does not own (probe), turn 1, row `named_card_role`

**Asked:** Should Sol Ring be your commander or one of the 99 cards?

Faults: inaccurate. Catalog action: reword.

The question is not valid as written because Sol Ring cannot be a commander, so it should ask only how to include it: "Should Sol Ring be included in the deck as one of the 99 cards?"

### 88. terse: no theme at all (probe), turn 1, row `power_commander`

**Asked:** Which power bracket should the deck target? Bracket 2 is precon level, 3 is upgraded, and 4 is high power.

Faults: assumes an answer. Catalog action: reword.

The user said “Surprise me” and then “You decide,” so the agent should choose the power level rather than ask the user to select one. Proposed wording: “I’ll choose the deck’s power bracket for you.”

### 88. terse: no theme at all (probe), turn 1, row `colors`

**Asked:** Do you have a color preference?

Faults: assumes an answer. Catalog action: reword.

The user explicitly delegated the decision with “Surprise me” and “You decide,” so asking for a color preference is unnecessary. Proposed wording: “I’ll choose the deck’s colors for you.”

### 88. terse: no theme at all (probe), turn 1, row `budget`

**Asked:** Do you have a budget for cards to buy?

Faults: assumes an answer. Catalog action: reword.

The user said “You decide,” which delegates the budget decision as well; the agent should choose an appropriate budget assumption instead of asking. Proposed wording: “I’ll choose a reasonable card-buying budget for you.”

### 100. terse: a deck for a team event (probe), turn 3, row `meta`

**Asked:** Should I keep the sideboard general, or will you name the decks you expect to face?

Faults: two questions in one. Catalog action: reword.

It asks both whether to keep the sideboard general and whether the user will name expected decks; reword as: "Which decks do you expect to face, or should I build a general sideboard?"

## Slots the eval says nobody asked about

These do not move the ratio. They are the guard against a quiet agent.

- 10. no strong commander in the library: Whether the deck should use only cards from the collection or prioritize owned cards while allowing purchases.
- 2. blink with a thin library: The agent should have asked for the user's available card collection or decklist, since the requested build was to be owned-first.
- 20. dinosaur tribal for a child: The child's age or experience level, since the deck is for a child and that affects card complexity and gameplay suitability.
- 22. extra turns commander: Whether the user's playgroup is comfortable with an extra-turns strategy, since this can affect card and win-condition choices.
- 23. land destruction: A follow-up asking the user to clarify their budget, since “My table does not mind” does not answer whether they have a spending limit.
- 30. another version, same plan: Which cards from the current deck should be replaced or avoided, or can you provide the current decklist so the alternative version can use different cards?, Should the alternative deck be limited to cards from your collection?
- 31. FNM on Friday, Modern: The agent should follow up on the unanswered color-preference question before choosing a deck, unless it is willing to choose colors independently., The agent should clarify any remaining purchase constraints beyond the stated $200 budget, such as whether the entire 75-card deck must fit within that amount.
- 32. burn on a budget: The user's preferred color identity or whether the burn deck should be mono-red or allow other colors.
- 36. the user changes the theme late (probe): After the user changed the theme, ask whether they still want Meren of Clan Nel Toth as commander and whether the token deck should replace the sacrifice theme entirely., Ask for the relevant cards or deckbuilding constraints from the user's library, since the user requested building from their library first.
- 37. the user contradicts the budget (probe): budget clarification after the user changed "$50" to "money is no object, spend what you need", colors follow-up, because the user never answered "Do you have a color preference?", specific expected Modern aggro matchups for sideboard planning, because "My shop plays aggro" does not identify the decks
- 44. a card that is not legal (probe): The agent should have addressed that Black Lotus is not legal in Modern and clarified whether to replace it with a legal card or reinterpret the request before proceeding.
- 46. two partner commanders (probe): Whether Thrasios, Triton Hero and Tymna the Weaver are intended to be partnered commanders.
- 48. sideboard help only (probe): Specific control archetypes or representative decklists the shop plays, since “a lot of control” is too broad for precise sideboard recommendations
- 52. a price cap per card (probe): A question about the total deck budget, since the user gave a per-card cap but not a total budget.
- 55. a returning player (probe): A follow-up asking whether the user has a commander in mind or wants a suggested commander, because the user did not answer that choice on turn 3.
- 56. the user answers a different question (probe): Whether the deck should still focus on mill despite the user's statement that their friends hate mill.
- 6. the strongest deck, no collection: The budget amount: the user said “under budget” but never specified how much they can spend., The desired tournament/deckbuilding priorities beyond raw strength, such as whether tournament legality and current metagame competitiveness should take precedence over other considerations.
- 61. terse: a proxy table (probe): The agent should have asked whether to build a Vintage deck despite its stated inability to build Vintage, rather than forcing an unsupported format choice., The agent should have clarified the requested deck size and construction rules after the user said “Any card, no ban list,” especially whether the deck should use normal Vintage rules or custom table rules., The agent should have asked whether “mono brown” means colorless, rather than relying on the earlier unanswered color-preference question.
- 63. terse: a competitive request (probe): Ask for the user's budget ceiling after they changed the request to “The best deck under budget.”, Clarify whether “under budget” supersedes the earlier “money is no object” constraint.
- 66. terse: sideboard help only (probe): The agent should ask for the current 60-card decklist, or at least the key maindeck cards and existing sideboard, because sideboard recommendations depend on the deck's actual composition., The agent should ask which control decks or color combinations are common at the shop, since “a lot of control” is not specific enough to target a tournament-level Modern sideboard.
- 74. terse: a card the user does not own (probe): A valid commander choice from the user's collection, such as which colorless artifact-appropriate commander they own or want to use.
- 77. terse: a Background pair (probe): Which Background and legendary creature should be the commander pair, or may the deck builder choose them?
- 85. terse: chaos (probe): The agent did not clarify what “as random as possible” means in practice, such as whether randomness should come from random card selection, random game effects, coin flips, or unpredictable politics.
- 92. terse: proxy the expensive cards only (probe): Clarify the unresolved commander choice after the user said “Blue and black” but named Bruvac the Grandiloquent, including whether Bruvac is intended as commander or whether a blue-black commander is needed., Clarify what “library first” means, since that phrase is ambiguous and was not previously solicited.
- 96. terse: the user asks about a card (probe): Answer whether Sol Ring is legal in Modern before asking deck-building questions.
- 97. terse: Pauper Commander (probe): Confirm which supported format the user wants after they did not answer the format-choice question.

## Run

- Calls: 102. Time: 677.5 seconds.
- Tokens: 110984 input (0 cached), 63175 output.
- Cost: $0.0980.

## The card check

The check ran against the local snapshot. No refusal rested on a card name the snapshot refutes.
