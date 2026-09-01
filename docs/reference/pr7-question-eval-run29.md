# PR-7 question eval

Run: `pr7-question-gate-run29`. Eval model: `gpt-5.6-luna`. Prompt version 3.

Scored 398 questions. 22 were not warranted, and 0 are unsure.

**Bad-question ratio: 11.2% on the holdout.**

The tune split holds 273 questions at 2.9%, and the holdout holds 125 at 11.2%. A fixer reads the tune split alone. A gain that shows only there is a reworded test set (D-134).

## Conversations the eval left unjudged

The verdicts did not line up with the questions asked, one per question. Their verdicts are dropped and not counted (T-9).

- 20. dinosaur tribal for a child: 4 verdicts for 5 questions
- 22. extra turns commander: 1 verdicts for 6 questions
- 3. a named card, role unknown: 3 verdicts for 6 questions
- 53. the user swaps the commander (probe): 1 verdicts for 4 questions

## The counters the ratio can not see

A run that asks less scores better and serves the user worse. A run can pass both bars with most of its conversations unanswered, so read these counters with the ratio.

| Counter | Value |
|---|---|
| Questions asked | 131 |
| Questions that closed a slot | 106 |
| Invented by the model | 1 |
| Catalog-only conversations | 29 |
| Premature sessions | 0 |
| Linter findings | 2 |

## Where the bad questions came from

| Row | Bad questions |
|---|---|
| `budget` | 6 |
| `colors` | 4 |
| `commander_pick` | 4 |
| `commander` | 2 |
| `format_unsupported_open` | 2 |
| `pool` | 1 |
| `power_commander` | 1 |
| `power_sixty` | 1 |
| `theme` | 1 |

## Faults

| Fault | Count |
|---|---|
| assumes an answer | 10 |
| out of scope | 5 |
| inaccurate | 4 |
| duplicate | 3 |

## What the catalog needs

| Action | Count |
|---|---|
| reword | 30 |
| remove | 9 |

## Every question the eval refused

The holdout conversations are left out of this list on purpose. 14 of their questions failed, and the fixer may not read which (D-134).

### 1. lifegain with a collection, turn 3, row `commander_pick`

**Asked:** Which commander do you want: Astarion, the Decadent, Karlov of the Ghost Council, or Sorin of House Markov // Sorin, Ravenous Neonate? Say "none" and I name three more.

Faults: assumes an answer. Catalog action: reword.

The user explicitly said, “Suggest a commander from my library,” so the question should have made the recommendation rather than asking the user to choose; proposed wording: “Based on your library and bracket 3 target, I recommend [commander]. Would you like me to build around it?”

### 7. FNM on Friday (probe), turn 2, row `colors`

**Asked:** Do you have a color preference?

Faults: duplicate. Catalog action: remove.

The agent already asked "Do you have a color preference?" on turn 1, and the user did not provide a new answer that would justify repeating it.

### 49. copy a known deck (probe), turn 1, row `colors`

**Asked:** Do you have a color preference?

Catalog action: reword.

The user had already said "whatever is winning," which indicates no color preference; reword this as a confirmation only when the user's request leaves colors genuinely open.

### 56. the user answers a different question (probe), turn 1, row `budget`

**Asked:** Do you have a budget for cards to buy?

Faults: out of scope. Catalog action: remove.

The user asked for a deck “from my collection,” which indicates that the initial build should use owned cards rather than cards to buy.

### 73. terse: a colorless deck (probe), turn 1, row `commander_pick`

**Asked:** Which commander do you want: Emry, Lurker of the Loch, Urza, Lord High Artificer, or Padeem, Consul of Innovation? Say "none" and I name three more.

Faults: inaccurate. Catalog action: reword.

The proposed commanders all have blue color identity, which does not fit the user's stated colorless deck; reword as: "Which legal colorless commander do you want, or should I suggest three options for a big-artifact strategy?"

### 82. terse: the user corrects a card name (probe), turn 1, row `colors`

**Asked:** Do you have a color preference?

Faults: assumes an answer, out of scope. Catalog action: remove.

Once a deck is built around Atraxa, the commander's color identity determines the deck's colors; the user did not need to choose a separate color preference.

### 86. terse: enchantress (probe), turn 1, row `budget`

**Asked:** Do you have a budget for cards to buy?

Faults: assumes an answer. Catalog action: reword.

The user said “from my collection,” so asking for a budget to buy cards presumes that purchases are allowed. Reword as: “May I suggest cards outside your collection?” and ask about a budget only if the user agrees.

### 95. terse: a library that does not exist (probe), turn 2, row `budget`

**Asked:** Do you have a budget for cards to buy?

Faults: assumes an answer. Catalog action: reword.

The request was to build from the user's library, but this question presumes cards will be bought; reword it as: “May I include cards outside your library, and if so, what budget should I use?”

## Slots the eval says nobody asked about

These do not move the ratio. They are the guard against a quiet agent.

- 102. terse: stax with no table named (probe): Whether the user has any commander preferences or is comfortable with the agent selecting one.
- 103. terse: two themes at once (probe): How should the two themes relate: should artifacts and lifegain be equally central, or should one be the primary strategy? The user said the deck should do “artifacts and lifegain,” but did not specify their relative priority.
- 16. the strongest modern deck: The user's budget limit or budget range, since “under budget” does not specify an amount and conflicts with the earlier “money is no object.”, Whether “the strongest Modern deck” should mean the current competitive best deck or the strongest deck affordable within the user's budget.
- 2. blink with a thin library: The user's owned-card inventory or deckbuilding collection list, so the deck can actually be built owned-first and the buy list can identify missing cards.
- 26. legacy for an event (probe): The agent should have asked what competitive level or expected event environment the Legacy deck should target.
- 29. two hundred dollars: collection_contents: Ask the user to provide an inventory or decklist of the cards in their library so the deck can be built from it first.
- 32. burn on a budget: Color identity, such as whether the burn deck should be mono-red or use other colors.
- 36. the user changes the theme late (probe): After the user changed the theme, the agent should have clarified whether Meren of Clan Nel Toth and the black-green color identity should remain, since the requested deck changed from sacrifice to tokens.
- 40. the user asks a question back (probe): Answer the user's question about what a power bracket means before asking another deck-building question.
- 44. a card that is not legal (probe): The requested centerpiece, Black Lotus, is not legal in Modern. The agent should have flagged this and asked whether to replace it with a legal analog or switch formats., The agent did not ask whether the $200 budget is for the entire deck or only cards to buy.
- 47. a deck as a gift (probe): Turn 3: follow up on the unanswered commander choice, since the user supplied power and budget but did not choose Gisa, Zul Ashur, Hancock, or none.
- 48. sideboard help only (probe): Ask for the current maindeck and sideboard lists, since sideboard recommendations depend on the cards already present., Ask which control archetypes or matchups are common at the shop, since the user only later said, “My shop plays a lot of control.”
- 49. copy a known deck (probe): Whether to optimize for the shop's aggro-heavy metagame or for the broader tournament metagame after the user said, "my shop plays aggro."
- 52. a price cap per card (probe): total budget
- 56. the user answers a different question (probe): Ask for the user's library or collection list so the deck can be built from available cards first.
- 61. terse: a proxy table (probe): The agent should have asked about the table’s ban-list or card-restriction policy, since the user specified Vintage but had not yet said whether Vintage’s normal restrictions applied.
- 66. terse: sideboard help only (probe): Ask for the current sideboard list and any sideboard cards the user already owns, so recommendations can identify gaps rather than duplicate existing cards.
- 7. FNM on Friday (probe): A supported format choice after the user said "Pioneer" and the agent said it does not build Pioneer., Color preference, which the user never answered.
- 70. terse: a mill deck for a playgroup (probe): Clarify the conflict between the requested blue-and-black color identity and Bruvac the Grandiloquent's commander color identity; ask whether to keep Bruvac and build mono-blue or use a blue-black commander., Clarify what currency the stated 100 budget uses.
- 71. terse: one-word answers (probe): Clarify whether the user wants a mono-white weenie deck or a control deck, since “Control” conflicts with the initial “Mono-white weenie deck.”
- 73. terse: a colorless deck (probe): Ask for a legal colorless commander or whether the user wants to change the deck's color identity., Ask whether the user already owns cards they want included or wants the deck built entirely from purchases.
- 75. terse: cEDH at bracket 5 (probe): The target Commander bracket or exact power-level expectation, since “cEDH” alone does not establish whether the user wants bracket 5.
- 85. terse: chaos (probe): commander choice
- 87. terse: landfall (probe): Ask which cards are in the user's binder/library, or request their collection list, before building the deck.
- 89. terse: teaching a new player (probe): A follow-up asking the user to choose a commander or approve suggested commander options after they provided the bracket and budget.
- 9. a commander the library does not hold: After the user said to build from their library first, ask which cards or decklist from their collection should be used, if the app cannot automatically inspect the library.
- 95. terse: a library that does not exist (probe): library/card-pool contents or an import of the user's collection, follow-up on the unresolved commander choice
- 96. terse: the user asks about a card (probe): Answer whether Sol Ring is legal in Modern.
- 97. terse: Pauper Commander (probe): Follow up on the unanswered format choice before proceeding; the user never selected Commander, Standard, or Modern., Ask for a commander choice or whether the agent should choose one, once the format is resolved.

## Run

- Calls: 103. Time: 809.0 seconds.
- Tokens: 110583 input (0 cached), 62982 output.
- Cost: $0.0977.

## The card check

The check ran against the local snapshot. No refusal rested on a card name the snapshot refutes.
