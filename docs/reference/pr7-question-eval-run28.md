# PR-7 question eval

Run: `pr7-question-gate-run28`. Eval model: `gpt-5.6-luna`. Prompt version 3.

Scored 391 questions. 14 were not warranted, and 0 are unsure.

**Bad-question ratio: 6.8% on the holdout.**

The tune split holds 273 questions at 2.2%, and the holdout holds 118 at 6.8%. A fixer reads the tune split alone. A gain that shows only there is a reworded test set (D-134).

## The counters the ratio can not see

A run that asks less scores better and serves the user worse. A run can pass both bars with most of its conversations unanswered, so read these counters with the ratio.

| Counter | Value |
|---|---|
| Questions asked | 115 |
| Questions that closed a slot | 96 |
| Invented by the model | 1 |
| Catalog-only conversations | 26 |
| Premature sessions | 0 |
| Linter findings | 0 |

## Where the bad questions came from

| Row | Bad questions |
|---|---|
| `budget` | 4 |
| `format_unsupported_open` | 3 |
| `commander_pick` | 2 |
| `colors` | 1 |
| `pool_precon` | 1 |
| `power_commander` | 1 |
| `power_sixty` | 1 |
| `theme` | 1 |

## Faults

| Fault | Count |
|---|---|
| assumes an answer | 8 |
| out of scope | 6 |
| duplicate | 1 |
| jargon | 1 |

## What the catalog needs

| Action | Count |
|---|---|
| reword | 15 |
| remove | 8 |

## Every question the eval refused

The holdout conversations are left out of this list on purpose. 8 of their questions failed, and the fixer may not read which (D-134).

### 35. the user changes the format (probe), turn 3, row `budget`

**Asked:** Do you have a budget for cards to buy?

Faults: duplicate. Catalog action: remove.

The agent had already asked for a budget on turn 1, and the user had not supplied a different budget-related need; this repeats the earlier question.

### 50. a format we do not support (probe), turn 2, row `budget`

**Asked:** Do you have a budget for cards to buy?

Faults: assumes an answer. Catalog action: reword.

The question presumes that the user may be buying cards and has a budget, although the user only described the format. Reword as: “Will you be buying cards for this deck?”

### 53. the user swaps the commander (probe), turn 1, row `budget`

**Asked:** Do you have a budget for cards to buy?

Faults: assumes an answer. Catalog action: remove.

The user asked for a deck “from my library,” so the question presumes that cards will be bought even though the stated source is the user's existing library; no budget row is needed here.

### 53. the user swaps the commander (probe), turn 3, row `commander_pick`

**Asked:** Which commander do you want: Astarion, the Decadent, Karlov of the Ghost Council, or Sorin of House Markov // Sorin, Ravenous Neonate? Say "none" and I name three more.

Faults: assumes an answer. Catalog action: reword.

The user explicitly asked for “a different commander,” but the choices still include Karlov of the Ghost Council, the previously selected commander. Reword the row to offer only commanders different from Karlov, such as: “Which different commander do you want: Astarion, the Decadent, or Sorin of House Markov // Sorin, Ravenous Neonate? Say ‘none’ and I’ll name three more.”

### 62. terse: a precon upgrade (probe), turn 1, row `pool_precon`

**Asked:** How should I use the precon: start with it, use only its cards, or build without it?

Catalog action: remove.

The user already said “Upgrade my Atraxa, Praetors' Voice precon,” which establishes that the precon should be the starting point; asking how to use it is redundant.

### 73. terse: a colorless deck (probe), turn 1, row `commander_pick`

**Asked:** Which commander do you want: Emry, Lurker of the Loch, Urza, Lord High Artificer, or Padeem, Consul of Innovation? Say "none" and I name three more.

Faults: out of scope. Catalog action: reword.

The user specified “A colorless Commander deck,” but every offered commander has blue color identity, so none of the listed options applies; offer colorless-identity commanders instead.

## Slots the eval says nobody asked about

These do not move the ratio. They are the guard against a quiet agent.

- 103. terse: two themes at once (probe): Which theme should be the deck's primary focus—artifacts or lifegain?, Do you want the agent to choose the commander, or do you have commander preferences?
- 105. a theme the library can not lead (probe): The agent did not ask which colorless Eldrazi commander the user wants, or whether it may choose one from the user's library.
- 16. the strongest modern deck: budget amount
- 2. blink with a thin library: Ask the user to provide or upload the list of cards in their library/collection so the deck can be built owned-first.
- 23. land destruction: After the user delegated the commander choice, the agent should have suggested commander options or asked for any remaining commander constraints., The budget question was not answered by “My table does not mind,” so the agent should have clarified whether there is a card-purchase budget.
- 29. two hundred dollars: Which cards in your library should the deck prioritize, or can I access your full collection list?
- 36. the user changes the theme late (probe): After the user changed the theme to tokens, ask whether to keep Commander, black-green, and Meren of Clan Nel Toth as the commander, or revise any of those constraints., After the theme change, ask what kind of token strategy the user wants, such as creature tokens, a specific token type, go-wide combat, or sacrifice-focused tokens.
- 39. the user stays vague (probe): A deck concept, strategy, or commander choice was needed because the user provided no direction beyond “Make me a good deck.”, A budget or spending constraint was not established before offering a buy-list path.
- 44. a card that is not legal (probe): The agent should have addressed that Black Lotus is not legal in Modern and asked whether to replace it with a legal centerpiece or change the format.
- 45. a commander that can not lead (probe): Turn 1: The agent should have explained that Lightning Bolt cannot be a legal commander and asked whether the user wanted to replace it with a legal red commander, rather than proceeding as though it could lead the deck.
- 48. sideboard help only (probe): The agent should have asked for the current deck and sideboard lists, or at least the cards already available for sideboarding, since the user wants help with an existing Modern burn deck.
- 49. copy a known deck (probe): Whether to optimize for the shop's aggro metagame or for the broader Standard tournament metagame.
- 50. a format we do not support (probe): Arena card-pool or card-availability constraints after the user said the deck is for Arena.
- 52. a price cap per card (probe): Ask what strategy, archetype, or theme the Modern deck should use., Ask whether the user has a total deck budget in addition to the $5-per-card cap.
- 56. the user answers a different question (probe): Ask the user to provide or make available their library/card collection list so the deck can be built from it.
- 57. terse: the format is named once (probe): Whether the user wants mass land destruction, targeted land destruction, or a land-denial theme with a particular play-pattern or social restriction.
- 63. terse: a competitive request (probe): Turn 2: Ask what budget limit the user means by “under budget,” since no amount was provided.
- 66. terse: sideboard help only (probe): The agent should have asked for the current sideboard list, and likely the relevant control decks or colors at the shop, since the user only said the shop plays "a lot of control."
- 7. FNM on Friday (probe): After the user said “Pioneer” and the agent said it does not build Pioneer, the agent needed to resolve whether the user would accept a supported format for FNM; it also needed to ask again for the still-unanswered color preference before recommending a deck.
- 71. terse: one-word answers (probe): Clarify the conflicting deck strategy: does the user want a mono-white weenie deck or a mono-white control deck?
- 77. terse: a Background pair (probe): Which legendary creature and Background should form the commander pair, or should the deck-builder choose them? The user specified only “A Commander deck with a Background commander pair.”
- 79. terse: everything buried in one long message (probe): Ask which specific cards from the user's collection they want included or can provide for the deck.
- 8. mill for a playgroup: Resolve the color-identity conflict between the user's stated blue-and-black preference and mono-blue commander Bruvac the Grandiloquent, such as asking whether to keep Bruvac and omit black cards or choose a blue-black commander.
- 80. terse: all caps, no punctuation (probe): Ask about the expected local metagame or how heavily the control-heavy shop environment should influence the build.
- 85. terse: chaos (probe): commander
- 87. terse: landfall (probe): The agent should have asked which cards are in the user's binder or collection, since the deck was requested specifically "from my binder."
- 88. terse: no theme at all (probe): The agent should have asked about the user's preferred playstyle or deck experience, since “Surprise me” leaves the theme and gameplay direction open.
- 89. terse: teaching a new player (probe): A beginner-oriented commander recommendation or brief explanation of each commander’s play pattern, since the user has never played before.
- 92. terse: proxy the expensive cards only (probe): Resolve the conflict between the requested blue-black color identity and Bruvac the Grandiloquent as a possible commander, since Bruvac cannot support black cards as commander.
- 94. terse: a collection of lands only (probe): Follow up on the unanswered commander selection before building the deck.
- 96. terse: the user asks about a card (probe): Answer whether Sol Ring is legal in Modern: it is not legal in Modern because it is banned., Confirm whether the user wants deck-building help before starting deck-preference intake.
- 97. terse: Pauper Commander (probe): commander choice

## Run

- Calls: 100. Time: 651.0 seconds.
- Tokens: 107146 input (0 cached), 58483 output.
- Cost: $0.0916.

## The card check

The check ran against the local snapshot. No refusal rested on a card name the snapshot refutes.
