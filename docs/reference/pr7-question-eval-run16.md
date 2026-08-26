# PR-7 question eval

Run: `pr7-question-gate-run16`. Eval model: `gpt-5.6-luna`. Prompt version 2.

Scored 356 questions. 40 were not warranted, and 0 are unsure.

**Bad-question ratio: 11.3% on the holdout.**

The tune split holds 241 questions at 11.2%, and the holdout holds 115 at 11.3%. A fixer reads the tune split alone. A gain that shows only there is a reworded test set (D-134).

## The counters the ratio can not see

A run that asks less scores better and serves the user worse. Gate run 7 of 2026-08-25 passed both bars with 26 of 30 conversations unanswered.

| Counter | Value |
|---|---|
| Questions asked | 124 |
| Questions that closed a slot | 98 |
| Invented by the model | 1 |
| Catalog-only conversations | 29 |
| Premature sessions | 0 |
| Linter findings | 0 |

## Where the bad questions came from

| Row | Bad questions |
|---|---|
| `commander_pick` | 10 |
| `power_sixty_confirm` | 6 |
| `colors` | 4 |
| `locked` | 4 |
| `format` | 3 |
| `pool_precon` | 2 |
| `power_commander` | 2 |
| `theme_card_named` | 2 |
| `format_unsupported_open` | 1 |
| `meta` | 1 |
| `one_deck` | 1 |
| `plan_choice` | 1 |
| `pool` | 1 |
| `power_sixty` | 1 |
| `theme` | 1 |

## Faults

| Fault | Count |
|---|---|
| duplicate | 22 |
| inaccurate | 8 |
| assumes an answer | 3 |
| out of scope | 3 |
| two questions in one | 1 |
| vague | 1 |

## What the catalog needs

| Action | Count |
|---|---|
| reword | 31 |
| remove | 20 |

## Every question the eval refused

The holdout conversations are left out of this list on purpose. 13 of their questions failed, and the fixer may not read which (D-134).

### 1. lifegain with a collection, turn 2, row `commander_pick`

**Asked:** Which one do you want: Cecil, Dark Knight // Cecil, Redeemed Paladin, Denethor, Ruling Steward, or Cloak and Dagger, Entwined? Say "none" and I name three more.

Faults: inaccurate. Catalog action: reword.

The user had not chosen a commander, so asking was warranted, but “Cloak and Dagger, Entwined” is not an applicable commander option; use wording such as “Which commander would you like to use, or should I suggest one from your collection?”

### 1. lifegain with a collection, turn 3, row `commander_pick`

**Asked:** Which one do you want: Cecil, Dark Knight // Cecil, Redeemed Paladin, Denethor, Ruling Steward, or Cloak and Dagger, Entwined? Say "none" and I name three more.

Faults: duplicate. Catalog action: remove.

This repeats the same commander-choice question from turn 2, while the user instead asked the agent to suggest a commander from their library: “Suggest a commander from my library.”

### 6. the strongest deck, no collection, turn 2, row `power_sixty_confirm`

**Asked:** Should I aim for tournament-level power, or a lower power level?

Faults: duplicate. Catalog action: remove.

The user had already asked for “the strongest deck possible” and then “The best deck under budget,” which answers the power-level question in favor of the highest competitive level.

### 13. a card to keep that is not the commander, turn 1, row `locked`

**Asked:** Should the deck keep Sanguine Bond, or may I cut it if it does not fit the plan?

Catalog action: remove.

The user already explicitly said “keep Sanguine Bond in it,” so asking whether to keep or cut that same card was unnecessary.

### 21. vintage with proxies, turn 1, row `format`

**Asked:** Which format should the deck use?

Faults: duplicate. Catalog action: reword.

The user had already identified the format in the opening message: “vintage with proxies.” Reword this row to ask only when no format has been provided.

### 25. poison in a tournament, turn 1, row `power_sixty_confirm`

**Asked:** Should I build the deck for tournament-level power, or would you prefer FNM-level or casual power?

Faults: duplicate. Catalog action: reword.

The user already asked for "an infect deck for a Modern tournament," which establishes the requested tournament-level context; the question should not ask them to choose tournament, FNM, or casual power again.

### 33. the user contradicts the budget (probe), turn 1, row `power_sixty_confirm`

**Asked:** What level should the deck target?

Catalog action: remove.

The user already said they wanted a competitive Modern deck, which supplies the intended power target; asking what level to target was unnecessary at this point.

### 51. a returning player (probe), turn 1, row `theme`

**Asked:** What would you like the deck to focus on: a creature type, a mechanic, or a play style?

Catalog action: reword.

The user already gave a focus—“I like big green creatures.” A better follow-up would be: “Would you like to refine that into a specific creature type, mechanic, or play style?”

### 51. a returning player (probe), turn 1, row `colors`

**Asked:** Do you have any preferred colors?

Catalog action: reword.

The user already stated a color preference with “big green creatures.” A better follow-up would be: “Would you like to add any other colors alongside green?”

### 54. terse: a card for the 99 (probe), turn 1, row `theme_card_named`

**Asked:** What should the deck focus on: a creature type, a mechanic, or a play style?

Faults: duplicate. Catalog action: remove.

The user already specified the theme: “Build around Grist, the Hunger Tide, but not as my commander.”

### 54. terse: a card for the 99 (probe), turn 2, row `locked`

**Asked:** Must the deck keep every card you named, or may I cut cards that do not fit the sacrifice plan?

Faults: duplicate. Catalog action: remove.

The user already required the deck to be built around Grist, so the named card was implicitly locked in; asking whether to cut “every card you named” repeats that constraint.

### 58. terse: a precon upgrade (probe), turn 1, row `pool_precon`

**Asked:** Should I use the precon as the base and add upgrades, or restrict the deck to cards from the precon?

Faults: duplicate. Catalog action: reword.

The user already said “Upgrade my Atraxa, Praetors' Voice precon,” which establishes using the precon as the base and adding upgrades.

### 60. terse: a deck as a gift (probe), turn 2, row `commander_pick`

**Asked:** Which one do you want: The Scarab God, Wilhelt, the Rotcleaver, or Ratadrabik of Urborg? Say "none" and I name three more.

Faults: duplicate. Catalog action: remove.

The exact same commander-choice question had already been asked on turn 1, and the user had not yet selected a commander; repeating it added no new information.

### 60. terse: a deck as a gift (probe), turn 3, row `commander_pick`

**Asked:** Which one do you want: The Scarab God, Wilhelt, the Rotcleaver, or Ratadrabik of Urborg? Say "none" and I name three more.

Faults: duplicate. Catalog action: remove.

The exact same commander-choice question had already been asked on turns 1 and 2, and the user still had not selected a commander; it should not be sent again.

### 61. terse: a named commander and a locked card (probe), turn 1, row `locked`

**Asked:** Should Sanguine Bond stay in the deck no matter what, or may I cut cards that do not fit the plan?

Catalog action: remove.

The user had already explicitly said “keep Sanguine Bond in it,” so asking whether it should stay no matter what was redundant.

### 64. terse: the strongest deck, no collection (probe), turn 1, row `power_sixty_confirm`

**Asked:** Should I build the deck for tournament-level competition?

Faults: duplicate. Catalog action: reword.

The user already asked for “the strongest Modern deck possible,” which establishes the desired competitive power level; ask only if the intended setting needs clarification.

### 67. terse: one-word answers (probe), turn 3, row `power_sixty_confirm`

**Asked:** What level of competitiveness do you want for the deck: tournament level, FNM level, or casual?

Faults: duplicate. Catalog action: remove.

The user had already answered the earlier competitiveness question with “FNM,” and the agent had already asked the same power-level question on turn 2.

### 69. terse: a colorless deck (probe), turn 1, row `colors`

**Asked:** Do you want the deck to remain colorless?

Catalog action: remove.

The user already said "A colorless Commander deck," so asking whether the deck should remain colorless was redundant.

### 70. terse: a card the user does not own (probe), turn 2, row `locked`

**Asked:** Should Sol Ring stay in the deck?

Faults: duplicate. Catalog action: remove.

The user already said “Sol Ring goes in it,” which answers whether Sol Ring should stay; this confirmation question should not be asked after an explicit inclusion request.

### 73. terse: a Background pair (probe), turn 2, row `commander_pick`

**Asked:** Which one do you want: Jaheira, Friend of the Forest, Karlach, Fury of Avernus, or Lae'zel, Vlaakith's Champion? Say "none" and I name three more.

Faults: duplicate, inaccurate. Catalog action: remove.

The commander question repeats the unanswered turn-1 question, and after the user said “Red and white,” Jaheira is not a legal red-white commander choice because its color identity includes green.

### 73. terse: a Background pair (probe), turn 3, row `commander_pick`

**Asked:** Which one do you want: Jaheira, Friend of the Forest, Karlach, Fury of Avernus, or Lae'zel, Vlaakith's Champion? Say "none" and I name three more.

Faults: duplicate, inaccurate. Catalog action: remove.

The commander question repeats the same unanswered question again, and Jaheira still does not apply to the user's stated “Red and white” color identity.

### 75. terse: everything buried in one long message (probe), turn 1, row `power_commander`

**Asked:** Which power bracket should the deck target?

Faults: duplicate. Catalog action: reword.

The user already said they want a casual deck at “bracket 2 or 3,” so asking which bracket to target repeats information already provided.

### 76. terse: all caps, no punctuation (probe), turn 2, row `power_sixty_confirm`

**Asked:** Should I build the deck for tournament-level competition?

Faults: duplicate. Catalog action: remove.

The user already answered this in turn 2 with “TOURNAMENT,” which directly selects tournament-level competition.

### 81. terse: chaos (probe), turn 1, row `commander_pick`

**Asked:** Which one do you want: Vincent, Vengeful Atoner, Shadow the Hedgehog, or Be'lakor, the Dark Master? Say "none" and I name three more.

Faults: inaccurate. Catalog action: reword.

The user had not named a commander, but two offered options—“Vincent, Vengeful Atoner” and “Shadow the Hedgehog”—are not valid Magic commanders; replace the options with real legal commanders.

### 93. terse: Pauper Commander (probe), turn 1, row `one_deck`

**Asked:** I build one deck at a time. Which deck do you want first?

Faults: assumes an answer. Catalog action: remove.

The user already requested a single deck by saying “A Pauper Commander deck,” so asking “Which deck do you want first?” presumes multiple decks without evidence.

### 93. terse: Pauper Commander (probe), turn 2, row `format`

**Asked:** Which format would you like for the black sacrifice deck?

Faults: duplicate, inaccurate. Catalog action: reword.

The user already specified “A Pauper Commander deck,” while the options separately list “Commander” and “Pauper” rather than the requested Pauper Commander variant; reword to “You specified Pauper Commander—what commander should lead the deck?”

### 96. terse: a deck for a team event (probe), turn 3, row `meta`

**Asked:** What decks do you expect to face? I’ll tune the 15-card sideboard to that.

Faults: duplicate. Catalog action: reword.

The user already said, “My teammates play aggro and combo,” which answers what decks the agent should expect to face; the question should instead ask about any additional expected archetypes or specific metagame details.

## Slots the eval says nobody asked about

These do not move the ratio. They are the guard against a quiet agent.

- 15. standard at the store, no library: budget
- 18. pauper burn on a budget: Exact budget cap
- 25. poison in a tournament: Ask for the user's budget, since the initial request did not specify one and budget is needed to build the best feasible Modern deck.
- 29. two hundred dollars: The agent should ask which cards in the user's library are available for deck construction, or access the collection inventory before building from it.
- 30. another version, same plan: On turn 2, clarify which earlier deck or plan this should match, since no source deck or plan is present in the transcript., On turn 2, ask which cards must be replaced or what card-availability constraints apply, since the user only said “different cards.”
- 32. the user changes the theme late (probe): Ask whether Meren of Clan Nel Toth and the black-green color identity should be retained for the token deck after the user changed themes., Ask whether the user has a preferred token strategy or token type for the new deck.
- 33. the user contradicts the budget (probe): A follow-up asking for the user's color preference, since the user never answered the turn 1 colors question.
- 38. shorthand and slang (probe): The agent should ask which commander the user wants, or whether it may choose one, since only “edh gruul dino stompy” was specified.
- 39. a card name with a typo (probe): A card-identity confirmation for the typo: whether the user meant Karlov of the Ghost Council.
- 40. a card that is not legal (probe): Acknowledge that Black Lotus is not legal in Modern and confirm that the replacement deck should remain Modern-legal.
- 43. a deck as a gift (probe): budget
- 45. copy a known deck (probe): budget
- 48. a price cap per card (probe): A question about the desired deck archetype or game plan, since the user had not yet said they wanted burn., A question about the total deck budget, since a per-card price cap does not establish a total budget.
- 5. 60-card anything goes: budget
- 50. two decks at once (probe): After the user gave only “Bracket 3, and 150 dollars,” the agent should have followed up on the unanswered commander choice, asking whether they have a commander in mind or want a suggestion.
- 52. the user answers a different question (probe): After “My friends hate mill,” ask whether the user still wants a mill deck despite the playgroup’s dislike of the strategy.
- 55. terse: two decks at once (probe): budget
- 56. terse: a format we do not build (probe): commander selection: the agent should have followed up after the user provided only the power bracket, since no commander had been chosen.
- 58. terse: a precon upgrade (probe): Upgrade budget or spending limit., Any specific cards already owned or required to include.
- 62. terse: sideboard help only (probe): The agent should ask for the current maindeck list or at least the key cards and flex slots, since sideboard choices depend on the exact Modern Burn build., The agent should clarify which control decks or colors the shop commonly plays; “a lot of control” is too broad to tune a tournament-level sideboard precisely.
- 68. terse: a five-color deck (probe): budget
- 69. terse: a colorless deck (probe): Budget
- 7. FNM on Friday: budget amount
- 71. terse: cEDH at bracket 5 (probe): budget
- 73. terse: a Background pair (probe): A question selecting the Background partner for the commander pair; the user requested “a Background pair,” but no Background was offered or selected.
- 75. terse: everything buried in one long message (probe): A concrete budget or spending limit for any cards the user may need to buy.
- 8. mill for a playgroup: collection contents or decklist, so the agent can build from the user's library first
- 80. terse: voltron (probe): budget
- 85. terse: teaching a new player (probe): color identity, budget
- 90. terse: a collection of lands only (probe): The agent did not ask whether the user has a preferred commander or wants the agent to choose one.
- 91. terse: a library that does not exist (probe): Whether the deck should use an unrestricted card pool, since the app does not have access to a card collection despite the request to build from the user's library., Which commander or commander constraints the deck should use.
- 92. terse: the user asks about a card (probe): Answer whether Sol Ring is legal in Modern; it is banned in Modern.
- 93. terse: Pauper Commander (probe): The agent should have asked which legal uncommon creature or legendary creature will be the commander, since the user has not specified one.
- 94. terse: Oathbreaker (probe): budget
- 95. terse: Historic on Arena (probe): After the user chose Modern, clarify that Modern is not an Arena format and ask whether they want a paper/online Modern deck or an Arena-supported format.
- 97. terse: mill with no table named (probe): Budget

## Run

- Calls: 98. Time: 697.6 seconds.
- Tokens: 93559 input (0 cached), 58460 output.
- Cost: $0.0889.

## The card check

The eval called these cards unreal, and the snapshot holds every one. The `inaccurate` fault was dropped from each verdict (D-149).

| Conversation and card | Row |
|---|---|
| 10. no strong commander in the library: Cloak and Dagger, Entwined | `commander_pick` |
