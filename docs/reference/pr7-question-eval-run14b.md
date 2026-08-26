# PR-7 question eval

Run: `pr7-question-gate-run14`. Eval model: `gpt-5.6-luna`. Prompt version 2.

Scored 259 questions. 40 were not warranted, and 0 are unsure.

**Bad-question ratio: 17.4% on the holdout.**

The tune split holds 173 questions at 14.5%, and the holdout holds 86 at 17.4%. A fixer reads the tune split alone. A gain that shows only there is a reworded test set (D-134).

## The counters the ratio can not see

A run that asks less scores better and serves the user worse. Gate run 7 of 2026-08-25 passed both bars with 26 of 30 conversations unanswered.

| Counter | Value |
|---|---|
| Questions asked | 134 |
| Questions that closed a slot | 89 |
| Invented by the model | 3 |
| Catalog-only conversations | 27 |
| Premature sessions | 0 |
| Linter findings | 0 |

## Where the bad questions came from

| Row | Bad questions |
|---|---|
| `commander_pick` | 18 |
| `pool` | 7 |
| `power_sixty_confirm` | 4 |
| `locked` | 3 |
| `format` | 2 |
| `house_format_limits` | 2 |
| `commander_illegal` | 1 |
| `format_unsupported` | 1 |
| `pool_precon` | 1 |
| `power_commander` | 1 |

## Faults

| Fault | Count |
|---|---|
| duplicate | 21 |
| assumes an answer | 19 |
| inaccurate | 10 |
| out of scope | 2 |
| two questions in one | 2 |
| vague | 2 |
| omits information | 1 |

## What the catalog needs

| Action | Count |
|---|---|
| reword | 30 |
| remove | 13 |

## Every question the eval refused

The holdout conversations are left out of this list on purpose. 15 of their questions failed, and the fixer may not read which (D-134).

### 1. lifegain with a collection, turn 3, row `commander_pick`

**Asked:** Which one do you want: Aerith Gainsborough, Minwu, White Mage, or Cecil, Dark Knight // Cecil, Redeemed Paladin? Say "none" and I name three more.

Faults: duplicate, assumes an answer. Catalog action: reword.

This repeats the prior commander-choice question even though the user asked to suggest a commander from their library, and it assumes the listed commanders are in that library without any library contents being provided.

### 3. a named card, role unknown, turn 3, row `commander_illegal`

**Asked:** Grist, the Hunger Tide can not lead a deck. Shall I suggest a commander instead?

Faults: duplicate, inaccurate. Catalog action: remove.

The user had already said “Grist is my commander,” and Grist can legally be a Commander because its characteristic-defining ability makes it a creature outside the battlefield; the question incorrectly claims “Grist, the Hunger Tide can not lead a deck.”

### 3. a named card, role unknown, turn 3, row `locked`

**Asked:** Should the deck keep Grist, the Hunger Tide, or may I cut cards that do not fit the plan?

Faults: duplicate, assumes an answer, vague. Catalog action: reword.

The user had already established Grist as the commander, so asking whether to keep Grist is redundant and presumes there are other named cards to cut; if additional named cards exist, ask: “Are there any other cards or deckbuilding constraints you want me to keep?”

### 4. fun and janky, turn 3, row `pool`

**Asked:** Should I use your library first, restrict the deck to your library, or ignore your library?

Faults: assumes an answer. Catalog action: reword.

The question presumes the user has a usable library, although the user had not mentioned owning or providing one; reword it as: “Do you have a card library or decklist you want me to use, or should I build without a supplied card pool?”

### 10. no strong commander in the library, turn 3, row `commander_pick`

**Asked:** Which one do you want: Aerith Gainsborough, Minwu, White Mage, or Cecil, Dark Knight // Cecil, Redeemed Paladin? Say "none" and I name three more.

Faults: assumes an answer. Catalog action: reword.

The user asked to “Buy the best lifegain commander,” so the agent should recommend or select the best candidate rather than presume the user will choose among these three; reword as: “I recommend [best candidate] for your $60, owned-first bracket-3 plan; should I use it?”

### 13. a card to keep that is not the commander, turn 1, row `locked`

**Asked:** Should Sanguine Bond stay in the deck, or may I cut cards that do not fit the plan?

Faults: duplicate. Catalog action: reword.

The user already said “keep Sanguine Bond in it,” so asking whether Sanguine Bond should stay repeats an answered requirement. Proposed wording: “Are there any other cards you want locked in?”

### 19. a big library, no theme, turn 2, row `commander_pick`

**Asked:** Which one do you want: Peregrin Took, Rosie Cotton of South Lane, or Quina, Qu Gourmet? Say "none" and I name three more.

Faults: inaccurate. Catalog action: reword.

The user had not chosen a commander, so a commander-selection question was warranted, but “Quina, Qu Gourmet” is not a valid Magic card option; offer only actual legal commander candidates.

### 19. a big library, no theme, turn 3, row `commander_pick`

**Asked:** Which one do you want: Peregrin Took, Rosie Cotton of South Lane, or Quina, Qu Gourmet? Say "none" and I name three more.

Faults: duplicate. Catalog action: remove.

The same commander-selection question was already asked on turn 2, and the user then explicitly said, “You pick the commander,” so repeating it was unwarranted.

### 21. vintage with proxies, turn 1, row `format`

**Asked:** Which format would you like: Commander, Standard, Modern, Pioneer, Pauper, or something else?

Faults: duplicate. Catalog action: reword.

The user had already specified “vintage” in the initial request, so asking which format they wanted was redundant; reword this as a conditional question asked only when no format is provided.

### 21. vintage with proxies, turn 2, row `house_format_limits`

**Asked:** For this custom Vintage-style deck, should I use a 60-card minimum, a four-copy limit per card name, and a 15-card sideboard?

Faults: two questions in one. Catalog action: reword.

The user said “Any card, no ban list” but had not specified deck size, copy limits, or sideboard size; separate these into individual questions because the prompt asks for three rule decisions at once.

### 22. extra turns commander, turn 3, row `commander_pick`

**Asked:** Which one do you want: Birgi, God of Storytelling // Harnfel, Horn of Bounty, Emrakul, the Promised End, or Vnwxt, Verbose Host? Say "none" and I name three more.

Faults: assumes an answer, inaccurate. Catalog action: reword.

The user said, "You pick the commander," so this should not ask them to choose; it also offers Birgi and Emrakul, whose color identities do not match the requested blue-red deck. Proposed wording: "I’ll choose a legal blue-red commander for the deck."

### 22. extra turns commander, turn 4, row `commander_pick`

**Asked:** Which one do you want: Birgi, God of Storytelling // Harnfel, Horn of Bounty, Emrakul, the Promised End, or Vnwxt, Verbose Host? Say "none" and I name three more.

Faults: duplicate, assumes an answer, inaccurate. Catalog action: reword.

The user had already said, "You pick the commander," and the exact commander-choice question was already asked on turn 3; Birgi and Emrakul also do not have the requested blue-red color identity. Proposed wording: "I’ll choose a legal blue-red commander for the deck."

### 24. stax, owned only, turn 3, row `commander_pick`

**Asked:** Which one do you want: Thalia, Guardian of Thraben, Pippin, Guard of the Citadel, or Kataki, War's Wage? Say "none" and I name three more.

Faults: assumes an answer. Catalog action: remove.

The user explicitly said “You pick the commander,” but this question assumes they want to choose among the listed commanders instead.

### 24. stax, owned only, turn 4, row `commander_pick`

**Asked:** Which one do you want: Thalia, Guardian of Thraben, Pippin, Guard of the Citadel, or Kataki, War's Wage? Say "none" and I name three more.

Faults: duplicate, assumes an answer. Catalog action: remove.

This repeats the same commander-selection question even though the user already delegated the choice with “You pick the commander,” and it again assumes the user should choose.

### 25. poison in a tournament, turn 1, row `power_sixty_confirm`

**Asked:** Should I build this for tournament-level play?

Faults: duplicate. Catalog action: reword.

The user already said “for a Modern tournament,” which answers the question’s tournament-level target; ask about power level only when the user has not already specified it.

### 27. a card for the 99, turn 4, row `commander_pick`

**Asked:** Which one do you want: Gríma Wormtongue, Old Man Willow, or The Sackville-Bagginses? Say "none" and I name three more.

Faults: duplicate, assumes an answer. Catalog action: reword.

The user already said, "You pick the commander," so asking "Which one do you want" repeats the commander decision and assumes the user will choose instead of the agent.

### 28. reanimator with two plans, turn 2, row `pool`

**Asked:** Should I use your library first, restrict the deck to only your library, or ignore it for a fully optimized build?

Faults: assumes an answer. Catalog action: reword.

The question presumes that the user has a library to use, although the user had only said "Commander. Black and green." Reword it as: "Do you have a card collection or library to build from? If so, should I use it first, restrict the deck to it, or ignore it for an optimized build?"

### 45. copy a known deck (probe), turn 1, row `power_sixty_confirm`

**Asked:** Should I build for tournament-level competition?

Faults: duplicate. Catalog action: reword.

The user already asked for “the current best Standard deck, whatever is winning,” which supplies a tournament-level power preference; reword as “Are you optimizing for sanctioned tournament results, or would you prefer a lower-power version?”

### 46. a format we do not support (probe), turn 3, row `commander_pick`

**Asked:** Which one do you want: Lathliss, Dragon Queen, Ganax, Astral Hunter, or Ran and Shaw? Say "none" and I name three more.

Faults: inaccurate. Catalog action: reword.

Commander selection is needed because the user said only “A dragon deck, red,” but “Ran and Shaw” is not a real Magic card or commander option; replace it with a valid option.

### 46. a format we do not support (probe), turn 4, row `commander_pick`

**Asked:** Which one do you want: Lathliss, Dragon Queen, Ganax, Astral Hunter, or Ran and Shaw? Say "none" and I name three more.

Faults: duplicate, inaccurate. Catalog action: remove.

This repeats the unanswered commander question verbatim and again offers the nonexistent “Ran and Shaw” option; the selection prompt should not be emitted twice in this flow.

### 51. a returning player (probe), turn 3, row `commander_pick`

**Asked:** Which one do you want: Nylea, God of the Hunt, Verdeloth the Ancient, or Nemata, Grove Guardian? Say "none" and I name three more.

Faults: duplicate. Catalog action: remove.

This repeats the earlier commander-selection question verbatim; the user had not answered it, but repeating it in the next turn is still a duplicate.

### 54. terse: a card for the 99 (probe), turn 4, row `commander_pick`

**Asked:** Which one do you want: Gríma Wormtongue, Old Man Willow, or The Sackville-Bagginses? Say "none" and I name three more.

Faults: duplicate. Catalog action: remove.

The user already answered the commander-selection question with “You pick the commander,” so asking which commander they want is redundant and contradicts that delegation.

### 58. terse: a precon upgrade (probe), turn 1, row `pool_precon`

**Asked:** Should I start by upgrading the precon, restrict the deck to cards from it, or build without using it?

Faults: duplicate. Catalog action: reword.

The user already said “Upgrade my Atraxa, Praetors' Voice precon,” which establishes that the deck should start from the precon; the question should only appear when the user has not specified the deck's starting pool.

### 61. terse: a named commander and a locked card (probe), turn 1, row `locked`

**Asked:** Must the deck keep Sanguine Bond, or may I cut cards that do not fit the plan?

Faults: duplicate. Catalog action: remove.

The user already explicitly said “keep Sanguine Bond in it,” so asking whether the deck must keep Sanguine Bond repeats an answered requirement.

### 64. terse: the strongest deck, no collection (probe), turn 1, row `power_sixty_confirm`

**Asked:** Should I build for tournament-level competition?

Catalog action: reword.

The user already asked for "the strongest Modern deck possible," which answers the tournament-level intent; reword as "If ‘strongest’ does not mean tournament-level, what competitiveness level should I target?"

## Slots the eval says nobody asked about

These do not move the ratio. They are the guard against a quiet agent.

- 1. lifegain with a collection: library contents or an import/list of the user's collection
- 12. a frozen run asks nothing: owned card pool: Which cards or collection should the deck be restricted to?, commander choice: If using Commander, which legendary creature or commander should lead the deck?
- 15. standard at the store, no library: budget or spending limit for acquiring the deck
- 17. precon upgrade at bracket 2: deck strategy or theme, which relevant cards are available in the user's library
- 18. pauper burn on a budget: Ask for the user's maximum budget or spending cap.
- 19. a big library, no theme: Ask how the agent can access the user's library or receive a card list for the requested library-first build.
- 2. blink with a thin library: Ask for the user's owned-card inventory or library list/export so an owned-first deck can actually be built.
- 20. dinosaur tribal for a child: budget, child's age or desired rules complexity, whether the user has a library/card pool available
- 25. poison in a tournament: budget
- 29. two hundred dollars: How can the app access or receive the user's library/card list so it can build from that library first?
- 3. a named card, role unknown: budget
- 30. another version, same plan: The original deck or card list being referenced, since the transcript does not identify which deck should be recreated with different cards.
- 32. the user changes the theme late (probe): After the user changed the theme, ask whether Commander, black-green colors, and Meren of Clan Nel Toth should remain., After the theme change, ask what token strategy or token types the user wants the deck to emphasize.
- 36. the user asks a question back (probe): budget or spending limit, whether the user has an existing card collection or wants a build from any available cards
- 38. shorthand and slang (probe): Whether the deck must be built exclusively from the user's existing binder, beyond the stated no-proxies constraint.
- 39. a card name with a typo (probe): Card-name confirmation: ask whether “Karlov of the Ghost Counsel” was intended to mean “Karlov of the Ghost Council.”
- 4. fun and janky: Whether the user has provided a card library or decklist that the agent can use, and how to obtain it if not.
- 43. a deck as a gift (probe): After the user gave the bracket and budget without answering the commander question, the agent should have followed up by asking whether it should suggest a black zombie commander or use one the user has in mind.
- 45. copy a known deck (probe): budget
- 47. the user declines everything (probe): commander
- 48. a price cap per card (probe): Total deck budget
- 49. the user swaps the commander (probe): After changing commanders, ask whether the deck should remain restricted to white-black or whether the user is open to the suggested commander's color identity.
- 52. the user answers a different question (probe): Whether the user still wants a mill deck despite saying their friends hate mill., What cards or card list from the user's library are available for the build.
- 53. terse: the format is named once (probe): The actual red-green card library or a way to access its contents is still needed to build from it.
- 55. terse: two decks at once (probe): commander identity or permission to proceed with a suggested commander after the user answered only the power bracket and budget
- 56. terse: a format we do not build (probe): The agent should have clarified whether the user wants a paper Commander deck or an Arena-legal alternative, since the request began with “for Arena” and Commander is not an Arena format.
- 58. terse: a precon upgrade (probe): Budget or card-availability constraints for the upgrades.
- 60. terse: a deck as a gift (probe): budget
- 61. terse: a named commander and a locked card (probe): The agent should have asked the user to provide or upload their library/decklist so it could build from that pool.
- 62. terse: sideboard help only (probe): Which control decks or other major archetypes are expected at the shop?, Are there budget or card-collection constraints for the 15-card sideboard?
- 65. terse: a Pauper deck named once (probe): budget
- 7. FNM on Friday: The actual budget amount: the user asked for the best deck under budget, but no dollar limit had been requested.
- 8. mill for a playgroup: currency for the 100 budget, whether the playgroup is comfortable with mill as the primary strategy, the user's library or decklist after they chose to build from it

## Run

- Calls: 64. Time: 494.4 seconds.
- Tokens: 59886 input (0 cached), 43746 output.
- Cost: $0.0645.
