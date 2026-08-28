# PR-7 question eval

Run: `pr7-question-gate-run25`. Eval model: `gpt-5.6-luna`. Prompt version 3.

Scored 382 questions. 19 were not warranted, and 0 are unsure.

**Bad-question ratio: 7.2% on the holdout.**

The tune split holds 271 questions at 4.1%, and the holdout holds 111 at 7.2%. A fixer reads the tune split alone. A gain that shows only there is a reworded test set (D-134).

## Conversations the eval left unjudged

The verdicts did not line up with the questions asked, one per question. Their verdicts are dropped and not counted (T-9).

- 21. vintage with proxies (probe): 4 verdicts for 6 questions
- 27. a card for the 99: a verdict on turn 1 row power, which asked no such question
- 4. fun and janky: 5 verdicts for 6 questions
- 86. terse: enchantress (probe): 3 verdicts for 4 questions
- 98. terse: Oathbreaker (probe): 3 verdicts for 6 questions

## The counters the ratio can not see

A run that asks less scores better and serves the user worse. Gate run 7 of 2026-08-25 passed both bars with 26 of 30 conversations unanswered.

| Counter | Value |
|---|---|
| Questions asked | 117 |
| Questions that closed a slot | 98 |
| Invented by the model | 1 |
| Catalog-only conversations | 26 |
| Premature sessions | 0 |
| Linter findings | 0 |

## Where the bad questions came from

| Row | Bad questions |
|---|---|
| `colors` | 5 |
| `budget` | 2 |
| `commander_pick` | 2 |
| `commander` | 1 |
| `format` | 1 |
| `format_unsupported_open` | 1 |
| `house_format_limits` | 1 |
| `named_card_role` | 1 |
| `pool` | 1 |
| `pool_precon` | 1 |
| `pool_thin` | 1 |
| `power_sixty` | 1 |
| `theme` | 1 |

## Faults

| Fault | Count |
|---|---|
| assumes an answer | 5 |
| duplicate | 5 |
| out of scope | 5 |
| inaccurate | 3 |
| omits information | 1 |

## What the catalog needs

| Action | Count |
|---|---|
| reword | 21 |
| remove | 8 |

## Every question the eval refused

The holdout conversations are left out of this list on purpose. 8 of their questions failed, and the fixer may not read which (D-134).

### 5. 60-card anything goes, turn 1, row `format`

**Asked:** Which format should the 60-card deck use: Commander, Standard, or Modern?

Faults: inaccurate, omits information. Catalog action: reword.

The user already specified “a 60-card deck” and that “anything goes at our table,” while Commander is not a 60-card format and the options omit an anything-goes house format. Reword as: “Should I follow your anything-goes house rules, or the rules and card legality of a named sanctioned format?”

### 8. mill for a playgroup, turn 4, row `house_format_limits`

**Asked:** Inside your house format, do the normal 60-card deck limits hold?

Faults: inaccurate. Catalog action: reword.

The user already established the format as Commander, whose normal deck construction does not use 60-card limits; reword this as “Are there any house rules or card bans at your playgroup beyond the official Commander rules?”

### 17. precon upgrade at bracket 2, turn 1, row `pool_precon`

**Asked:** Should I start with your precon, restrict the deck to its cards, or ignore it for a fully optimized build?

Faults: duplicate. Catalog action: reword.

The user already said “Upgrade my Atraxa, Praetors' Voice precon,” which establishes that the precon should be the starting point; ask instead whether precon cards may be cut freely or whether some must be retained.

### 20. dinosaur tribal for a child, turn 3, row `pool_thin`

**Asked:** To reach at least 30 Dinosaurs, should I prioritize cards you already own and add a buy list, or use the entire available pool?

Faults: assumes an answer. Catalog action: reword.

The question assumes the user wants at least 30 Dinosaurs, but the user only said “A dinosaur deck for my kid” and did not provide that numerical target. Proposed wording: “Should I prioritize cards you already own and add a buy list, or use the full available card pool?”

### 43. a card name with a typo (probe), turn 2, row `named_card_role`

**Asked:** Should Karlov of the Ghost Counsel lead the deck, or be one card in the 99?

Faults: inaccurate. Catalog action: reword.

The question uses the apparent incorrect name “Karlov of the Ghost Counsel” instead of clarifying the user's typo; reword it as “Should Karlov of the Ghost Council lead the deck, or be one card in the 99?”

### 49. copy a known deck (probe), turn 1, row `colors`

**Asked:** Do you have a color preference?

Catalog action: remove.

No: the user already said “whatever is winning,” which answers that they have no meaningful color preference.

### 53. the user swaps the commander (probe), turn 2, row `budget`

**Asked:** Do you have a budget for cards to buy?

Faults: out of scope. Catalog action: remove.

The user asked for a deck from their library, so asking for a budget for cards to buy does not apply to the requested build.

### 56. the user answers a different question (probe), turn 1, row `budget`

**Asked:** Do you have a budget for cards to buy?

Faults: assumes an answer. Catalog action: reword.

The user asked for a deck "from my collection," so asking about "cards to buy" assumes they are open to purchases; reword as: "Should I restrict the deck to your collection, or may I include cards to buy?"

### 70. terse: a mill deck for a playgroup (probe), turn 1, row `commander_pick`

**Asked:** Which one do you want: Hope Estheim, The Mouth of Sauron, or Saruman of Many Colors? Say "none" and I name three more.

Faults: assumes an answer. Catalog action: remove.

The user already specified blue and black and named Bruvac the Grandiloquent as commander in turn 2, so asking them to choose among three different commanders was no longer appropriate at that point.

### 70. terse: a mill deck for a playgroup (probe), turn 1, row `colors`

**Asked:** Do you have a color preference?

Faults: duplicate. Catalog action: remove.

The user already gave the deck's colors as blue and black in turn 2; the color-preference question was therefore redundant when evaluated against the full preceding context.

### 82. terse: the user corrects a card name (probe), turn 1, row `colors`

**Asked:** Do you have a color preference?

Faults: out of scope. Catalog action: remove.

Atraxa already determines the deck’s color identity, so a general color preference is not actionable unless it asks about preferences within Atraxa’s colors.

## Slots the eval says nobody asked about

These do not move the ratio. They are the guard against a quiet agent.

- 100. terse: a deck for a team event (probe): The agent should have asked which Modern deck archetype or control strategy the user wants, and whether the deck should complement the teammates' aggro and combo roles., The agent should have asked whether the $600 budget is for the entire deck or only cards to buy, including whether the user already owns any cards.
- 102. terse: stax with no table named (probe): The agent should ask about the intended playgroup or table expectations for a stax strategy, since stax decks can be socially inappropriate or mismatched for some groups.
- 103. terse: two themes at once (probe): Turn 2: ask how to resolve the conflict between the stated white-blue color preference and the offered mono-color commanders, or provide commanders within white-blue identity.
- 104. terse: the user answers with a number (probe): Clarify what the user's reply “2” refers to and, if it answers the budget question, what currency or amount they mean., Obtain the user's commander choice or a clear request to suggest one; “2” does not match either commander option.
- 16. the strongest modern deck: Ask what budget limit the user means by “under budget.”
- 19. a big library, no theme: Whether the deck must use only cards from the user's collection or may recommend purchases after building from the collection first.
- 36. the user changes the theme late (probe): After the user changed the theme, ask whether the deck should now be built as a token deck while retaining Commander, black-green colors, and Meren of Clan Nel Toth, or whether any of those choices should change., After the user changed the theme, ask what kind of token strategy or token preferences they want, since “token deck” is broader than the original sacrifice theme.
- 43. a card name with a typo (probe): Clarify the apparent card-name typo before using the card as the commander candidate (the user wrote “Karlov of the Ghost Counsel”).
- 44. a card that is not legal (probe): Address that Black Lotus is not legal in Modern and ask whether to replace it with a legal centerpiece or switch formats.
- 48. sideboard help only (probe): Ask for the current maindeck and any existing sideboard, since sideboard recommendations depend on the cards already in the deck., Ask what opposing decks or metagame the sideboard should target. The user later identifies a control-heavy shop, but that information was not available when the questions were asked.
- 5. 60-card anything goes: Clarify whether “Modern” is only a label and the deck may use any cards, or whether normal Modern card legality still applies despite the stated no-ban-list house rule.
- 52. a price cap per card (probe): total budget, whether the user already owns cards to use
- 56. the user answers a different question (probe): Whether the user still wants a mill deck despite saying, "My friends hate mill," or would prefer an alternative strategy., Whether "from my library first" means the deck must use only the collection or may later include purchased cards.
- 62. terse: a precon upgrade (probe): Which cards from the user's collection should be prioritized or added to the precon? The user has a collection, and this information is needed to build around their available cards.
- 63. terse: a competitive request (probe): Ask what the budget limit is after the user changed the request to “The best deck under budget.”
- 66. terse: sideboard help only (probe): Ask which control archetypes the shop commonly plays and what the current sideboard or available decklist is, since those details are needed to tailor Modern sideboard recommendations.
- 7. FNM on Friday (probe): Whether the user is willing to switch from Pioneer to one of the supported formats after stating Pioneer is their desired format., Whether the user has any color or archetype preference, since the repeated color question was never answered and the shop's aggro-heavy metagame may affect the recommendation.
- 71. terse: one-word answers (probe): Resolve the conflict between the user's initial "Mono-white weenie deck" description and the later "Control" strategy preference.
- 76. terse: a companion (probe): The agent should have addressed whether the requested Modern deck is intended to be format-legal with Lurrus of the Dream-Den as companion, including the companion’s legality and deckbuilding restriction.
- 77. terse: a Background pair (probe): After the user specified "Red and white, aggressive," ask them to choose a Background commander pair compatible with that color preference, or clarify whether they want to relax the color requirement.
- 8. mill for a playgroup: The user's currency for the 100 budget., The user's actual library/card list or a way to provide it, since they asked to build from their library first.
- 84. terse: voltron (probe): Clarify the contradictory power target: the user said both “The first one” (apparently bracket 1) and “Bracket 3.”
- 9. a commander the library does not hold: Confirm the commander and format before building, since “Brago blink deck” does not explicitly state that Brago is the commander or that the format is Commander.
- 95. terse: a library that does not exist (probe): Ask the user to provide or connect the cards in their library, since no card collection is available to the app.
- 96. terse: the user asks about a card (probe): Answer whether Sol Ring is legal in Modern before asking deck-building questions.

## Run

- Calls: 100. Time: 716.2 seconds.
- Tokens: 108889 input (0 cached), 56656 output.
- Cost: $0.0898.

## The card check

The check ran against the local snapshot. No refusal rested on a card name the snapshot refutes.
