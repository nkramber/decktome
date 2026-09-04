# PR-8 deck gate

Run date: 2026-09-04. Card snapshot: 2026-09-03.

Verdict: PASS. 1 of 1 decks passed every block check, 0 invented names reached the user, and 0 summaries stated a false rule of the game. All three bars are zero tolerance.

## Summary

| Measure | Value |
|---|---|
| Prompts | 1 |
| Decks returned | 1 |
| Decks with no block finding | 1 |
| Invented names that reached the user | 0 |
| Decks that needed the repair turn | 0 |
| Summaries judged (F-26) | 1 |
| Summaries that state a rule of the game | 0 |
| Summaries that state a FALSE rule | 0 |
| Judge errors | 0 |
| Decks the plan judge read (PR-15, information) | 1 |
| Mean plan score, 0 to 1 | 0.25 |
| Errors | 0 |
| Prompt version | 12 |
| Calls | 3 |
| Cost | $0.1089 |
| Time | 82 seconds |

## Run

- Suite `decks`, run `pr8-deck-gate-run15`, on 2026-09-04, commit `ec058b4`.
- Roles: generate on `gpt-5.6-terra` (openai, effort medium), judge on `claude-opus-5` (anthropic, effort medium), repair on `gpt-5.6-terra` (openai, effort medium).
- Versions: card snapshot 2026-09-03, generate prompt version 12, plan_rubric prompt version 1, precons `5.3.0+20260903`, quality_model `20260903T210658Z`.
- Calls: 3. Cost: $0.1089. Time: 82 seconds.

## Findings by code

A block stops the deck. A warning and a note are reports.

| Code | Count |
|---|---|
| `not_owned` | 1 |
| `curve_summary` | 1 |

By severity: BLOCK 0. WARN 1. INFO 1. 

## The precon exclusion (PR-24)

A deck asked to use no card of a precon holds none of its cards. The products' copies leave the owned counts, so a card with a spare copy in the binder stays usable, and a basic land never leaves (D-408, D-37). An excluded card in the deck is a block.

| # | Prompt | Products | Cards excluded | Cards spare | Excluded cards in the deck | Blocked |
|---|---|---|---|---|---|---|
| 25 | use no card of an owned precon | Avengers Assemble | 57 | 27 | 0 | 0 |

## Decks

### 25. use no card of an owned precon

Format: Commander. Theme: Avengers superheroes. Pool: owned_first. Shortlist: 152 names.

Commander: Syr Konrad, the Grim.

Grade: typical, score 0.47, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $249.94 the whole deck.

**Summary:** This Syr Konrad deck develops its mana with artifacts and lands, keeps cards flowing, and establishes a creature-heavy board backed by a broad package of removal and board clears. It aims to win by pressing with its larger creature threats while maintaining enough protection and interaction to keep its key pieces in play. It gives up the speed of the format's most compressed lists for a steadier, more board-focused game.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: There is a discernible mono-black -1/-1 counter / attrition subtheme (Blowfly Infestation, Midnight Banshee, Soul Snuffers, Carnifex Demon, Contagion Clasp, Black Sun's Zenith) that also feeds Syr Konrad's death triggers, but the deck carries no self-mill, no meaningful token/sacrifice engine, and a large pile of generic equipment and single-target removal that does little for Konrad, so the pieces only loosely converge.
- PLAN theme_fit=no: The request was an Avengers superhero deck, and this is a mono-black Syr Konrad list full of Lord of the Rings, Final Fantasy, and Avatar cards with no Marvel or superhero identity whatsoever.
- PLAN useful_as_built=partly: The mono-black mana base and curve are functional and playable, but the win path is vague — few evasive or large threats, no assembled Konrad loop — and several wipes (Black Sun's Zenith, Midnight Banshee, Massacre Girl) actively fight the deck's own small creatures.
- PLAN summary_honest=no: The summary describes a generic 'creature-heavy board pressing with larger threats' plan and never mentions Syr Konrad's drain-on-death engine or the -1/-1 counter subtheme, and the claim of a creature-heavy board is false for a list with roughly twenty creatures and dozens of artifacts and removal spells.
- [WARN] `not_owned`: Syr Konrad, the Grim: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 2.97 over 63 nonland cards

<details><summary>The deck list</summary>

- 25 Swamp | land | Provides the deck's primary black mana.
- 1 Castle Locthwain | land | A black-producing land from the shortlist.
- 1 Takenuma, Abandoned Mire | land | A black-producing land from the shortlist.
- 1 Command Tower | land | Provides flexible commander-deck mana.
- 1 Exotic Orchard | land | Provides flexible mana.
- 1 Evolving Wilds | land | Helps establish the mana base.
- 1 Fabled Passage | land | Helps establish the mana base.
- 1 Marsh Flats | land | Helps establish the mana base.
- 1 Path of Ancestry | land | A land from the shortlist for the mana base.
- 1 Thriving Moor | land | Provides black mana.
- 1 Spire of Industry | land | A flexible land for the artifact-heavy support package.
- 1 Grand Coliseum | land | Provides flexible mana.
- 1 Sol Ring | ramp | An efficient mana artifact from the shortlist.
- 1 Arcane Signet | ramp | Provides reliable mana acceleration.
- 1 Astral Cornucopia | ramp | Adds artifact-based mana support.
- 1 Bender's Waterskin | ramp | Adds mana support from the shortlist.
- 1 Brass Infiniscope | ramp | Adds artifact-based mana support.
- 1 Commander's Sphere | ramp | Provides mana support.
- 1 Deadly Dispute | ramp | Contributes to the deck's resource development.
- 1 Explorer's Scope | ramp | Adds mana development support.
- 1 Thought Vessel | ramp | Provides artifact-based mana support.
- 1 Wayfarer's Bauble | ramp | Provides early mana development.
- 1 Buster Sword | draw | A card-advantage artifact from the shortlist.
- 1 Call of the Ring | draw | Provides card advantage.
- 1 Circle of Power | draw | Provides card advantage.
- 1 Hoarder's Greed | draw | Provides card advantage.
- 1 Idol of Oblivion | draw | Provides artifact-based card advantage.
- 1 Lembas | draw | Provides card advantage through an artifact.
- 1 Mask of Memory | draw | Provides equipment-based card advantage.
- 1 Nasty End | draw | Provides card advantage.
- 1 Night's Whisper | draw | Provides efficient card advantage.
- 1 Skullclamp | draw | Provides equipment-based card advantage.
- 1 Wizard's Rockets | draw | Provides artifact-based card advantage.
- 1 Bitter Downfall | removal | A removal option from the shortlist.
- 1 Bitter Triumph | removal | A removal option from the shortlist.
- 1 Blight Rot | removal | A removal option from the shortlist.
- 1 Claim the Precious | removal | A removal option from the shortlist.
- 1 Dismember | removal | A removal option from the shortlist.
- 1 Epic Downfall | removal | A removal option from the shortlist.
- 1 Fatal Push | removal | An efficient removal option.
- 1 Heartless Act | removal | A removal option from the shortlist.
- 1 Infernal Grasp | removal | A removal option from the shortlist.
- 1 Black Sun's Zenith | wipe | A board-clearing option from the shortlist.
- 1 Extinction Event | wipe | A board-clearing option from the shortlist.
- 1 One Ring to Rule Them All | wipe | A board-clearing Saga from the shortlist.
- 1 Champion's Helm | interaction | Protects an important creature.
- 1 Darksteel Plate | interaction | Protects an important creature.
- 1 Lightning Greaves | interaction | Protects an important creature.
- 1 Nameless Inversion | interaction | A flexible instant-speed interaction piece.
- 1 Orcish Medicine | interaction | An interaction option from the shortlist.
- 1 Stone of Erech | interaction | An artifact interaction piece.
- 1 Swiftfoot Boots | interaction | Protects an important creature.
- 1 The Walls of Ba Sing Se | interaction | A creature-based interaction piece.
- 1 Archfiend of Ifnir | synergy | A creature that supports the deck's shared plan.
- 1 Blowfly Infestation | synergy | An enchantment that supports the deck's shared plan.
- 1 Contagion Clasp | synergy | An artifact that supports the deck's shared plan.
- 1 Grave Venerations | synergy | An enchantment that supports the deck's shared plan.
- 1 Massacre Girl, Known Killer | synergy | A legendary creature that supports the deck's shared plan.
- 1 Midnight Banshee | synergy | A creature that supports the deck's shared plan.
- 1 Obsessive Pursuit | synergy | An enchantment that supports the deck's shared plan.
- 1 Rat King, Pale Piper | synergy | A legendary creature that supports the deck's shared plan.
- 1 Soul Snuffers | synergy | A creature that supports the deck's shared plan.
- 1 Undercity Dire Rat | synergy | A creature that supports the deck's shared plan.
- 1 Carnifex Demon | threat | A substantial creature threat from the shortlist.
- 1 Dusk Urchins | threat | A creature threat from the shortlist.
- 1 Gorbag of Minas Morgul | threat | A legendary creature threat from the shortlist.
- 1 Grim Poppet | threat | An artifact creature threat from the shortlist.
- 1 Lobelia Sackville-Baggins | threat | A legendary creature threat from the shortlist.
- 1 Orcish Bowmasters | threat | A creature threat from the shortlist.
- 1 Ringwraiths | threat | A creature threat from the shortlist.
- 1 Sinister Gnarlbark | threat | A creature threat from the shortlist.
- 1 Skinrender | threat | A creature threat from the shortlist.
- 1 Voracious Fell Beast | threat | A creature threat from the shortlist.
- 1 Witch-king of Angmar | threat | A legendary creature threat from the shortlist.
- 1 Zodiark, Umbral God | threat | A legendary creature threat from the shortlist.

</details>

