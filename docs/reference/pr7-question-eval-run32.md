# PR-7 question eval

Run: `pr7-question-gate-run32`. Eval model: `gpt-5.6-luna`. Prompt version 3.

Scored 393 questions. 22 were not warranted, and 0 are unsure.

**Bad-question ratio: 8.8% on the holdout.**

The tune split holds 268 questions at 4.1%, and the holdout holds 125 at 8.8%. A fixer reads the tune split alone. A gain that shows only there is a reworded test set (D-134).

## Conversations the eval left unjudged

The verdicts did not line up with the questions asked, one per question. Their verdicts are dropped and not counted (T-9).

- 20. dinosaur tribal for a child: 4 verdicts for 5 questions
- 43. a card name with a typo (probe): 6 verdicts for 7 questions
- 77. terse: a Background pair (probe): 2 verdicts for 4 questions

## The counters the ratio can not see

A run that asks less scores better and serves the user worse. A run can pass both bars with most of its conversations unanswered, so read these counters with the ratio.

| Counter | Value |
|---|---|
| Questions asked | 114 |
| Questions that closed a slot | 90 |
| Invented by the model | 0 |
| Catalog-only conversations | 27 |
| Premature sessions | 0 |
| Linter findings | 0 |

## Where the bad questions came from

| Row | Bad questions |
|---|---|
| `colors` | 4 |
| `commander` | 3 |
| `commander_pick` | 3 |
| `theme` | 3 |
| `pool_thin` | 2 |
| `budget` | 1 |
| `format` | 1 |
| `format_unsupported_open` | 1 |
| `house_format_limits` | 1 |
| `pool_precon` | 1 |
| `power_commander` | 1 |
| `power_sixty` | 1 |

## Faults

| Fault | Count |
|---|---|
| assumes an answer | 5 |
| inaccurate | 5 |
| duplicate | 4 |
| out of scope | 4 |
| two questions in one | 1 |

## What the catalog needs

| Action | Count |
|---|---|
| reword | 27 |
| remove | 6 |

## Every question the eval refused

The holdout conversations are left out of this list on purpose. 11 of their questions failed, and the fixer may not read which (D-134).

### 7. FNM on Friday (probe), turn 2, row `colors`

**Asked:** Do you have a color preference for the deck?

Faults: duplicate. Catalog action: remove.

The agent already asked "Do you have a color preference?" on turn 1, and the user had not supplied a color preference; this repeats the same unanswered question.

### 22. extra turns commander, turn 4, row `house_format_limits`

**Asked:** Inside your house format, do the normal 60-card deck limits hold?

Faults: assumes an answer, inaccurate. Catalog action: remove.

The question presumes a separate 'house format' that the user never described, and it incorrectly refers to 'normal 60-card deck limits' after the user selected Commander, whose normal construction rules are not 60-card limits.

### 44. a card that is not legal (probe), turn 1, row `theme`

**Asked:** What should the deck do: focus on a creature type, a mechanic, or a play style?

Faults: duplicate. Catalog action: reword.

No: the user already supplied a deck theme by saying “a Modern deck built around Black Lotus”; the question should instead address the card's Modern legality and ask for a legal replacement or revised centerpiece.

### 49. copy a known deck (probe), turn 1, row `colors`

**Asked:** Do you have a color preference?

Catalog action: reword.

The user asked for the deck that is winning "whatever," which effectively leaves colors open; reword as "I’ll choose the winning colors unless you have a preference—do you?"

### 53. the user swaps the commander (probe), turn 3, row `commander_pick`

**Asked:** Which commander do you want: Astarion, the Decadent, Karlov of the Ghost Council, or Sorin of House Markov // Sorin, Ravenous Neonate? Say "none" and I name three more.

Faults: assumes an answer. Catalog action: reword.

The user explicitly said to use a different commander, but the question still offers Karlov of the Ghost Council, the commander they had already chosen; reword it to offer only commanders other than Karlov.

### 62. terse: a precon upgrade (probe), turn 1, row `pool_precon`

**Asked:** How should I use your precon as the card pool: start with it, use only its cards, or skip it for a fully optimized deck?

Catalog action: reword.

The user already said “Upgrade my Atraxa, Praetors' Voice precon,” which establishes that the precon should be the starting pool; ask only about whether additions are allowed or whether the build must use only precon cards.

### 64. terse: a deck as a gift (probe), turn 1, row `budget`

**Asked:** What spending limit, if any, should I use for cards you do not own?

Faults: assumes an answer. Catalog action: reword.

The question is useful, but user_has_a_card_collection is false, while “cards you do not own” assumes the user owns some cards; reword it as: “What total budget should I use for the deck?”

### 73. terse: a colorless deck (probe), turn 1, row `commander_pick`

**Asked:** Which commander do you want: Emry, Lurker of the Loch, Urza, Lord High Artificer, or Padeem, Consul of Innovation? Say "none" and I name three more.

Faults: inaccurate. Catalog action: reword.

The question offers only blue-identity commanders—“Emry, Lurker of the Loch,” “Urza, Lord High Artificer,” and “Padeem, Consul of Innovation”—for a deck the user explicitly described as colorless; reword it to offer commanders whose color identity fits colorless, or ask whether the user has a commander in mind.

### 82. terse: the user corrects a card name (probe), turn 1, row `colors`

**Asked:** Do you have a color preference?

Catalog action: remove.

Atraxa, Praetor's Voice fixes the deck's color identity, so asking for a color preference is unnecessary unless the user is open to changing commanders.

### 88. terse: no theme at all (probe), turn 1, row `theme`

**Asked:** Should the deck center on a creature type, mechanic, or play style, or would you like me to choose?

Faults: duplicate. Catalog action: reword.

The user already said “Surprise me,” which answers the question by delegating the theme choice to the agent. Reword it as an acknowledgement rather than asking again, such as “I’ll choose the theme for you.”

### 107. a set name that names two base sets (probe), turn 2, row `commander`

**Asked:** Which commander do you want? Name one, or I suggest three.

Faults: two questions in one. Catalog action: reword.

This asks both whether the user has a specific commander in mind and whether they want suggestions; reword as: “Do you have a commander in mind, or would you like me to suggest three options?”

## Slots the eval says nobody asked about

These do not move the ratio. They are the guard against a quiet agent.

- 102. terse: stax with no table named (probe): An open-ended power target after the user rejected all five offered power options.
- 105. a theme the library can not lead (probe): After the user specified “Colorless,” ask for a legal colorless commander from their library or offer to choose one from that library.
- 16. the strongest modern deck: budget amount after the user changed the request to “The best deck under budget.”
- 22. extra turns commander: A follow-up on the unanswered budget question was needed before proceeding, since the user did not provide a spending limit for cards outside their library.
- 28. reanimator with two plans: The deck’s second plan or backup win condition, since the request is labeled “reanimator with two plans.”
- 29. two hundred dollars: Ask which cards or types of cards are available in the user's library, so the deck can be built from that collection first.
- 32. burn on a budget: The agent should have asked which colors the Standard burn deck should use, since the user had not yet specified mono-red.
- 36. the user changes the theme late (probe): After the theme change, ask whether Meren of Clan Nel Toth and the black-green color identity should remain, since the new token-deck request does not explicitly confirm them., Ask whether there is a spending limit for cards outside the user's library; the earlier budget question was never answered.
- 37. the user contradicts the budget (probe): After the user changed the budget, confirm whether the $50 limit was withdrawn and the budget is now unlimited., The user’s aggro preference does not answer the earlier color-preference question; ask which colors, if any, they want. 
- 44. a card that is not legal (probe): The agent should have identified that Black Lotus is not legal in Modern and asked what legal card or concept should replace it.
- 48. sideboard help only (probe): Ask for the current maindeck and any existing sideboard cards, since recommendations need to fit the user's actual Modern mono-red Burn list.
- 50. a format we do not support (probe): After the user said “Bracket 2,” follow up on the still-unanswered commander choice, house-format deck limits, and spending limit before building the deck.
- 56. the user answers a different question (probe): After the user said “My friends hate mill,” ask whether to keep the mill theme for that playgroup or adjust the deck to a less frustrating strategy.
- 63. terse: a competitive request (probe): Ask what budget limit the user means by “under budget,” since it conflicts with the earlier “money is no object.”
- 64. terse: a deck as a gift (probe): A total deck budget should have been requested without assuming the user owns any cards.
- 66. terse: sideboard help only (probe): The agent should have asked for the current sideboard or decklist and any sideboard constraints, since the user requested sideboard help but provided no cards or matchup details.
- 69. terse: a Pauper deck named once (probe): power level
- 73. terse: a colorless deck (probe): A commander choice compatible with the user's stated colorless deck, such as asking whether they already have a colorless commander in mind or offering colorless-identity options.
- 8. mill for a playgroup: Resolve the conflict between the requested blue-black colors and Bruvac the Grandiloquent as commander, whose deck color identity does not include black.
- 84. terse: voltron (probe): The agent did not ask which cards, if any, the user already owns, despite the deckbuilding request and the user's collection status being false.
- 89. terse: teaching a new player (probe): A question about the desired complexity or teaching focus of the deck, such as whether to prioritize very simple mechanics and easy-to-track cards for the new player.
- 92. terse: proxy the expensive cards only (probe): Follow up on the unanswered commander choice, including whether Bruvac the Grandiloquent is intended as commander despite the stated blue-and-black color preference.
- 94. terse: a collection of lands only (probe): A follow-up asking the user to choose a commander, or to say "none" for three more options.
- 96. terse: the user asks about a card (probe): Answer the user's legality question: Sol Ring is not legal in Modern.

## Run

- Calls: 103. Time: 776.7 seconds.
- Tokens: 110902 input (0 cached), 63014 output.
- Cost: $0.0978.

## The card check

The check ran against the local snapshot. No refusal rested on a card name the snapshot refutes.
