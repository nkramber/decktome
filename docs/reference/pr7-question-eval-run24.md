# PR-7 question eval

Run: `pr7-question-gate-run24`. Eval model: `gpt-5.6-luna`. Prompt version 3.

Scored 435 questions. 20 were not warranted, and 0 are unsure.

**Bad-question ratio: 8.6% on the holdout.**

The tune split holds 295 questions at 2.7%, and the holdout holds 140 at 8.6%. A fixer reads the tune split alone. A gain that shows only there is a reworded test set (D-134).

## The counters the ratio can not see

A run that asks less scores better and serves the user worse. Gate run 7 of 2026-08-25 passed both bars with 26 of 30 conversations unanswered.

| Counter | Value |
|---|---|
| Questions asked | 133 |
| Questions that closed a slot | 101 |
| Invented by the model | 2 |
| Catalog-only conversations | 28 |
| Premature sessions | 0 |
| Linter findings | 0 |

## Where the bad questions came from

| Row | Bad questions |
|---|---|
| `budget_scope` | 2 |
| `colors` | 2 |
| `format` | 2 |
| `format_unsupported_open` | 2 |
| `pool_thin` | 2 |
| `theme_card_named` | 2 |
| `budget` | 1 |
| `commander` | 1 |
| `commander_pick` | 1 |
| `house_format_limits` | 1 |
| `pool` | 1 |
| `pool_precon` | 1 |
| `power_commander` | 1 |
| `power_sixty` | 1 |

## Faults

| Fault | Count |
|---|---|
| assumes an answer | 7 |
| inaccurate | 4 |
| out of scope | 3 |
| duplicate | 2 |
| two questions in one | 1 |
| vague | 1 |

## What the catalog needs

| Action | Count |
|---|---|
| reword | 26 |
| remove | 8 |

## Every question the eval refused

The holdout conversations are left out of this list on purpose. 12 of their questions failed, and the fixer may not read which (D-134).

### 2. blink with a thin library, turn 4, row `budget_scope`

**Asked:** Is the $40 limit for the cards you buy or for the deck’s total value?

Faults: duplicate. Catalog action: remove.

The user had already answered the preceding budget question with “No more than 40 dollars,” and then requested “Build owned-first with a buy list,” making the intended scope the cards to buy sufficiently clear.

### 5. 60-card anything goes, turn 1, row `format`

**Asked:** Which format should the 60-card deck use: Commander, Standard, or Modern?

Faults: inaccurate. Catalog action: reword.

The user already specified a 60-card deck where “anything goes at our table,” so offering only Commander, Standard, or Modern does not cover the stated house format. Proposed wording: “Should I treat this as your custom 60-card house format, or apply the rules of a sanctioned format?”

### 17. precon upgrade at bracket 2, turn 1, row `pool_precon`

**Asked:** Should I start with your precon, restrict the deck to cards from it, or build without using it?

Faults: vague. Catalog action: reword.

The user already said “Upgrade my Atraxa, Praetors' Voice precon,” which establishes that the precon should be the starting point; reword as: “Should upgrades be limited to cards in the precon, or may I add cards from your library and purchases?”

### 22. extra turns commander, turn 2, row `pool_thin`

**Asked:** The card pool is thin for this theme. Should I build owned-first with a buy list, or use the whole pool?

Faults: assumes an answer. Catalog action: reword.

The question asserts that “The card pool is thin” without evidence from the user. Reword it as: “Should I prioritize cards from your library first, or use the full card pool?”

### 53. the user swaps the commander (probe), turn 2, row `budget`

**Asked:** Is there a budget for cards to buy?

Faults: assumes an answer. Catalog action: remove.

The user asked for a deck “from my library,” so asking about “cards to buy” presumes purchases that the user had not requested.

### 61. terse: a proxy table (probe), turn 2, row `house_format_limits`

**Asked:** Inside your house format, do the normal 60-card deck limits hold?

Faults: assumes an answer. Catalog action: reword.

The user said “Any card, no ban list,” but never established a “house format”; reword as “For the supported format you choose, should normal 60-card deck limits apply?”

### 92. terse: proxy the expensive cards only (probe), turn 2, row `budget_scope`

**Asked:** Is your budget a cap on the cards you buy or on the deck's total value?

Faults: assumes an answer. Catalog action: reword.

The user gave a proxy rule rather than a conventional budget cap: “I proxy anything over 20 dollars.” Reword as: “Besides your $20 proxy threshold, do you have a limit on what you will spend on cards?”

### 104. terse: the user answers with a number (probe), turn 3, row `commander`

**Asked:** Do you have a black-green commander in mind, or should I suggest one?

Faults: two questions in one. Catalog action: reword.

The user had specified a black-green sacrifice deck but had not named a commander; ask one question first, such as “Do you already have a black-green commander in mind?” rather than asking both whether they have one and whether the agent should suggest one.

## Slots the eval says nobody asked about

These do not move the ratio. They are the guard against a quiet agent.

- 100. terse: a deck for a team event (probe): Whether the team event has team-specific deckbuilding or role constraints beyond Modern legality.
- 102. terse: stax with no table named (probe): Ask about the intended playgroup or table expectations for the stax deck, since no table or meta was named., Resolve the commander choice and reconcile the stated white-blue color preference with the available options; none of the offered commanders has white-blue color identity.
- 105. probe: a theme the library can not lead (probe): A question checking whether the user wants to proceed if their library cannot support a viable Eldrazi theme, or whether the deck should relax the Eldrazi-only theme.
- 23. land destruction: A follow-up asking the user to clarify their card-purchase budget, since “My table does not mind” does not answer the budget question.
- 28. reanimator with two plans: The deck’s reanimation plan: whether to focus on reanimating one large threat or on a recurring value loop.
- 29. two hundred dollars: Ask which cards are in the user's library or request an inventory/export so the deck can be built from it.
- 30. another version, same plan: Ask which existing deck or deck list the user wants to keep as the reference, since “same deck” and “same plan” do not identify it in the transcript.
- 36. the user changes the theme late (probe): After the user changed the theme, ask whether Meren of Clan Nel Toth should remain the commander and whether the deck should now be built as a black-green token deck., After the user said to build from the library first, clarify whether they want to avoid buying cards entirely or merely prioritize cards already owned.
- 43. a card name with a typo (probe): Confirm the intended card name, since the user wrote “Karlov of the Ghost Counsel,” which appears to be a misspelling or ambiguous card name.
- 48. sideboard help only (probe): The agent should ask for the current maindeck and existing sideboard list, or at least the deck's key cards and weaknesses, so the sideboard can be tailored to this burn deck.
- 52. a price cap per card (probe): The user's total deck budget, separate from the $5-per-card cap.
- 56. the user answers a different question (probe): social constraint: whether to keep the mill theme despite the user's friends disliking mill
- 61. terse: a proxy table (probe): A supported format still needed to be selected after the user declined to name one; the agent should not proceed as though a house format had been established.
- 63. terse: a competitive request (probe): The user's budget amount or ceiling after saying “The best deck under budget.”
- 66. terse: sideboard help only (probe): The current main-deck and sideboard list, including any cards the user wants to keep, is needed to give actionable sideboard advice., The specific control decks or archetypes expected at the shop should be clarified after the user says the shop plays a lot of control.
- 7. FNM on Friday (probe): Turn 2: ask for the exact budget amount before selecting a replacement format, since “under budget” is not a usable budget., Turn 3: ask whether the user has a card collection or existing cards to use, since the user has no collection and the deck must be built under a $200 budget., Turn 3: ask whether the user prioritizes maximizing tournament competitiveness or specifically countering the shop’s mostly aggro metagame.
- 71. terse: one-word answers (probe): After the user answered “Control,” ask which control archetypes or decks they expect to face, since that answer is too broad to build a targeted Modern sideboard.
- 74. terse: a card the user does not own (probe): Clarify whether the deck may include purchased cards despite the original constraint “only the cards I own.”, Obtain or re-ask for the commander choice after the user did not answer whether they had one in mind or wanted a suggestion.
- 76. terse: a companion (probe): format legality of Lurrus of the Dream-Den as a Modern companion
- 78. terse: partners named once (probe): Desired power level or playgroup bracket
- 8. mill for a playgroup: Request the user's library or collection list so the deck can be built from it first.
- 80. terse: all caps, no punctuation (probe): After the user said “MY SHOP PLAYS CONTROL,” ask which control decks or archetypes they expect to face so the sideboard can be tailored.
- 87. terse: landfall (probe): Whether the deck must use only cards in the binder or may include cards to buy.
- 96. terse: the user asks about a card (probe): Answer whether Sol Ring is legal in Modern before starting deck-building clarification.
- 99. terse: Historic on Arena (probe): Clarify whether the user wants an Arena-supported format or a paper Modern deck, since Modern is not an Arena format.

## Run

- Calls: 103. Time: 722.9 seconds.
- Tokens: 113634 input (0 cached), 63295 output.
- Cost: $0.0987.

## The card check

The check ran against the local snapshot. No refusal rested on a card name the snapshot refutes.
