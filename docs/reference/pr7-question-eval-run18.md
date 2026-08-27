# PR-7 question eval

Run: `pr7-question-gate-run18`. Eval model: `gpt-5.6-luna`. Prompt version 3.

Scored 369 questions. 32 were not warranted, and 0 are unsure.

**Bad-question ratio: 10.4% on the holdout.**

The tune split holds 244 questions at 7.8%, and the holdout holds 125 at 10.4%. A fixer reads the tune split alone. A gain that shows only there is a reworded test set (D-134).

## The counters the ratio can not see

A run that asks less scores better and serves the user worse. Gate run 7 of 2026-08-25 passed both bars with 26 of 30 conversations unanswered.

| Counter | Value |
|---|---|
| Questions asked | 124 |
| Questions that closed a slot | 97 |
| Invented by the model | 2 |
| Catalog-only conversations | 28 |
| Premature sessions | 0 |
| Linter findings | 0 |

## Where the bad questions came from

| Row | Bad questions |
|---|---|
| `commander_pick` | 5 |
| `power_sixty_confirm` | 5 |
| `format_unsupported_open` | 4 |
| `power_commander` | 4 |
| `colors` | 3 |
| `locked` | 3 |
| `meta` | 3 |
| `commander` | 1 |
| `house_format_limits` | 1 |
| `pool_precon` | 1 |
| `power_sixty` | 1 |
| `theme_card_named` | 1 |

## Faults

| Fault | Count |
|---|---|
| duplicate | 16 |
| assumes an answer | 6 |
| out of scope | 5 |
| two questions in one | 3 |
| inaccurate | 2 |

## What the catalog needs

| Action | Count |
|---|---|
| reword | 38 |
| remove | 11 |

## Every question the eval refused

The holdout conversations are left out of this list on purpose. 13 of their questions failed, and the fixer may not read which (D-134).

### 1. lifegain with a collection, turn 3, row `commander_pick`

**Asked:** Which one do you want: Cecil, Dark Knight // Cecil, Redeemed Paladin, Denethor, Ruling Steward, or Cloak and Dagger, Entwined? Say "none" and I name three more.

Faults: duplicate. Catalog action: remove.

The same commander-selection question was already asked on turn 2, and the user instead requested a commander suggestion from their library: "Suggest a commander from my library, and build from my library first."

### 10. no strong commander in the library, turn 3, row `commander_pick`

**Asked:** Which one do you want: Cecil, Dark Knight // Cecil, Redeemed Paladin, Denethor, Ruling Steward, or Cloak and Dagger, Entwined? Say "none" and I name three more.

Faults: assumes an answer. Catalog action: remove.

The user already asked the agent to buy the best lifegain commander, so asking the user to choose among candidates improperly assumes they want to make the selection; the agent should evaluate the listed legal commanders using the user's owned-first and $60 constraints.

### 13. a card to keep that is not the commander, turn 1, row `locked`

**Asked:** Should Sanguine Bond stay in the deck, or may I cut cards that do not fit the plan?

Faults: duplicate. Catalog action: reword.

The user already said “keep Sanguine Bond in it,” so this should instead ask about any other named cards: “Are there any other cards you want locked in?”

### 18. pauper burn on a budget (probe), turn 1, row `format_unsupported_open`

**Asked:** I do not build Pauper. Which format should I build instead: Commander, Standard, or Modern?

Faults: assumes an answer. Catalog action: reword.

No: the user had already specified Pauper in “A Pauper burn deck,” so this presumes they want to abandon the requested format. Reword it as: “Pauper is not supported here. Would you like to switch to Commander, Standard, or Modern?”

### 18. pauper burn on a budget (probe), turn 2, row `format_unsupported_open`

**Asked:** I do not build Pauper. Which format should I build instead: Commander, Standard, or Modern?

Faults: duplicate. Catalog action: reword.

No: this repeats the same format-switch question after the user again explicitly answered “Pauper.” The unsupported-format prompt should be conditional and not sent again after the user declines to switch.

### 25. poison in a tournament, turn 1, row `power_sixty_confirm`

**Asked:** Should I build this for tournament-level competition?

Faults: duplicate. Catalog action: reword.

The user already specified “in a tournament,” so asking “Should I build this for tournament-level competition?” repeats information already provided. Proposed wording: “What level of competition should I target if it differs from the tournament setting you mentioned?”

### 34. no ban list at our table, turn 2, row `house_format_limits`

**Asked:** For this custom format, should the deck use a 60-card minimum, four copies per name, and a 15-card sideboard?

Faults: two questions in one. Catalog action: reword.

The question asks for three separate construction limits at once: “60-card minimum,” “four copies per name,” and “a 15-card sideboard.” Reword it into separate questions or a single question asking the user to specify the format's deck-size, copy-limit, and sideboard rules.

### 37. the user contradicts the budget (probe), turn 2, row `meta`

**Asked:** What decks do you expect to play against? I’ll use that to tune the 15-card sideboard.

Faults: assumes an answer. Catalog action: reword.

At this point the user had not mentioned a shop, playgroup, tournament, or expected opponents; reword it as "If you know your local metagame, which decks do you expect to face?"

### 39. the user stays vague (probe), turn 2, row `commander`

**Asked:** Do you have a commander in mind, or should I suggest one?

Faults: duplicate. Catalog action: reword.

The user had already said “I dunno, you pick,” which answers the choice by authorizing the agent to suggest and select a commander. Reword to: “I’ll choose the commander unless you have a specific constraint.”

### 63. terse: a competitive request (probe), turn 1, row `power_sixty_confirm`

**Asked:** Should I build for tournament-level competition?

Catalog action: reword.

The user already asked for “the strongest Modern deck,” which answers the tournament-level power question; reword this row to ask only when the user has not already specified a competitive power level.

### 73. terse: a colorless deck (probe), turn 1, row `colors`

**Asked:** What color preference should the deck use?

Faults: assumes an answer. Catalog action: reword.

The user already specified “a colorless Commander deck,” so asking for a color preference assumes that information was not provided. Reword as a conditional question asked only when the user has not already given colors.

### 75. terse: cEDH at bracket 5 (probe), turn 1, row `power_commander`

**Asked:** Which power bracket should the deck target?

Faults: duplicate. Catalog action: reword.

The user had already said “A cEDH deck,” which supplies the requested power level; reword this row to ask for a power bracket only when the user has not already stated one.

### 79. terse: everything buried in one long message (probe), turn 1, row `power_commander`

**Asked:** Which power bracket should the deck target?

Faults: duplicate. Catalog action: remove.

The user already specified the desired power level as “maybe bracket 2 or 3,” so asking which power bracket to target adds no needed information.

### 90. terse: a deck for a spouse (probe), turn 3, row `commander_pick`

**Asked:** Which one do you want: Arahbo, Roar of the World, Kaheera, the Orphanguard, or Kutzil, Malamet Exemplar? Say "none" and I name three more.

Faults: duplicate. Catalog action: remove.

This repeats the commander-selection question from turn 2 verbatim, while the user’s intervening reply, “Under 60 dollars,” did not answer or change that question.

### 93. terse: two budget shapes at once (probe), turn 3, row `meta`

**Asked:** What decks do you expect to face? I’ll tune the 15 sideboard cards to that.

Faults: assumes an answer. Catalog action: reword.

The question presumes the user has a particular set of decks or an event metagame in mind, which the transcript does not establish; reword it as: "Should I build a general-purpose sideboard, or tune it for particular Modern matchups?"

### 96. terse: the user asks about a card (probe), turn 1, row `theme_card_named`

**Asked:** What should the deck focus on: a creature type, a mechanic, or a play style?

Faults: out of scope. Catalog action: remove.

The user asked only, “Is Sol Ring legal in Modern?” and did not ask for deck-building help, so questions about deck theme, power, and colors were not relevant at this point.

### 96. terse: the user asks about a card (probe), turn 1, row `power_sixty`

**Asked:** How strong should the deck be?

Faults: out of scope. Catalog action: remove.

The user asked only, “Is Sol Ring legal in Modern?” and did not ask for deck-building help, so deck-strength preferences were not needed.

### 96. terse: the user asks about a card (probe), turn 1, row `colors`

**Asked:** Do you have a color preference?

Faults: out of scope. Catalog action: remove.

The user asked only, “Is Sol Ring legal in Modern?” and did not ask for deck-building help, so a color preference was not needed.

### 99. terse: Historic on Arena (probe), turn 1, row `format_unsupported_open`

**Asked:** I do not build Historic. Which format should I build instead: Commander, Standard, or Modern?

Faults: inaccurate. Catalog action: reword.

The user specified Arena, but Modern and Commander are not applicable Arena deck formats. Reword as: “I do not build Historic for Arena. Would you like an Arena-supported format instead, such as Standard?”

## Slots the eval says nobody asked about

These do not move the ratio. They are the guard against a quiet agent.

- 100. terse: a deck for a team event (probe): budget
- 101. terse: mill with no table named (probe): budget
- 102. terse: stax with no table named (probe): Budget or spending limit, Playgroup/table expectations and restrictions, Whether the deck should use cards the user already owns
- 103. terse: two themes at once (probe): budget
- 25. poison in a tournament: budget
- 26. legacy for an event (probe): Ask about the user's budget before recommending a deck., Ask about the desired competitiveness or power level for the event.
- 27. a card for the 99: commander: after the user replied "Commander. A sacrifice deck, black and green," the agent should have clarified which commander to use, since none had been chosen.
- 30. another version, same plan: Ask the user to provide or identify the original deck/list, since “the same deck” and “the same plan” do not specify which deck or plan is meant.
- 31. FNM on Friday, Modern: The user's budget amount, since “under budget” did not specify a dollar limit.
- 32. burn on a budget: The agent should have asked for the user's budget cap, since “as cheap as possible” does not specify an actionable dollar limit and the user later supplied “25 dollars is the cap.”
- 33. a tempo deck for an event: A direct question about the user's budget or maximum spend should have been asked on turn 1, rather than only offering a budget-based option.
- 36. the user changes the theme late (probe): After the user changed the theme to tokens, ask whether Commander and Meren of Clan Nel Toth remain the format and commander, and clarify the desired token strategy before building.
- 39. the user stays vague (probe): At turn 3, the agent should have confirmed any practical constraints still needed to build the deck, especially budget or whether to prioritize the user's collection, rather than ending without a follow-up.
- 47. a deck as a gift (probe): budget or spending limit
- 48. sideboard help only (probe): The agent did not ask for the existing Modern burn deck’s mainboard/list or which cards the user already owns, both of which are needed to build a compatible sideboard.
- 49. copy a known deck (probe): budget, card collection or acquisition constraints
- 52. a price cap per card (probe): Ask for the user's total deck budget, since a per-card cap does not establish an overall spending limit.
- 55. a returning player (probe): budget
- 59. terse: two decks at once (probe): Ask for the Commander deck's budget before proceeding.
- 61. terse: a proxy table (probe): The intended supported format remains unresolved after the user said “Any card, no ban list”; the agent should have followed up to select or define a supported format., The deck’s specific artifact-prison plan or win condition was not elicited after “Casual power, mono brown.”
- 63. terse: a competitive request (probe): budget amount or ceiling after the user changed the request to “The best deck under budget.”
- 64. terse: a deck as a gift (probe): A budget or price limit for the gift deck.
- 69. terse: a Pauper deck named once (probe): After the user supplied power level and budget without choosing an alternative format, ask again which supported format to use.
- 7. FNM on Friday (probe): Turn 3: confirm which supported format to build after the user chose Pioneer, since the user did not select Commander, Standard, or Modern., Turn 3: follow up on the unanswered color-preference question, or explicitly ask whether the user has no color preference.
- 71. terse: one-word answers (probe): Budget or card-acquisition constraints, since user_has_a_card_collection is false.
- 73. terse: a colorless deck (probe): budget
- 76. terse: a companion (probe): The agent should have addressed that Lurrus of the Dream-Den is not currently legal as a Modern companion and asked whether the user wants a legal replacement or is targeting a historical rules environment.
- 8. mill for a playgroup: The currency for the 100 budget, since the user gave only “budget 100.”, Whether Bruvac’s mono-blue color identity is acceptable despite the earlier preference for “Blue and black.”
- 80. terse: all caps, no punctuation (probe): Budget or spending limit for the Modern burn deck.
- 82. terse: the user corrects a card name (probe): budget
- 83. terse: group hug (probe): budget
- 84. terse: voltron (probe): budget: What is your budget for the deck?, card availability: Should I build from cards you already own, or can I suggest cards to buy?
- 85. terse: chaos (probe): Budget for the deck, since the user has no known card collection., Whether “as random as possible” means random effects, unpredictable game outcomes, or a specific chaos theme.
- 89. terse: teaching a new player (probe): budget
- 90. terse: a deck for a spouse (probe): The agent did not ask whether the $60 budget includes the commander and all 100 cards, or whether the user has any preferred spending limits for individual cards.
- 92. terse: proxy the expensive cards only (probe): Clarify how to handle the conflict between the requested blue-black identity and Bruvac the Grandiloquent as commander, since Bruvac's color identity does not include black.
- 95. terse: a library that does not exist (probe): Ask the user to provide or connect their card library/card pool, since the deck is supposed to be built from it and no collection is available.
- 96. terse: the user asks about a card (probe): Answer the user's legality question: Sol Ring is banned in Modern and therefore not legal in a Modern deck.
- 98. terse: Oathbreaker (probe): The agent should have asked for the budget context explicitly: whether the $150 budget is for the entire deck and whether it is a hard maximum, since the user supplied only “150 dollars.”

## Run

- Calls: 102. Time: 815.4 seconds.
- Tokens: 110813 input (0 cached), 60875 output.
- Cost: $0.0952.

## The card check

The check ran against the local snapshot. No refusal rested on a card name the snapshot refutes.
