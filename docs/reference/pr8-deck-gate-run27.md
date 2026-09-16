# PR-8 deck gate

Run date: 2026-09-16. Card snapshot: 2026-09-04.

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
| Case assertions missed (PR-28b) | 0 |
| Decks the plan judge read (PR-15, information) | 1 |
| Mean plan score, 0 to 1 | 0.62 |
| Plan reasons the judge left empty | 0 |
| Errors | 0 |
| Prompt version | 15 |
| Calls | 3 |
| Cost | $0.1018 |
| Time | 85 seconds |

## Run

- Suite `decks`, run `pr8-deck-gate-run27`, on 2026-09-16, commit `6ff0244`.
- Partial run over `5`. It reads the bars of its own items, and it never stands for the suite.
- Roles: generate on `gpt-5.6-terra` (openai, effort medium), judge on `claude-opus-5` (anthropic, effort medium), repair on `gpt-5.6-terra` (openai, effort medium).
- Versions: card snapshot 2026-09-04, generate prompt version 15, plan_rubric prompt version 2, quality_model `20260914T154223Z`.
- Calls: 3. Cost: $0.1018. Time: 85 seconds.

## Findings by code

A block stops the deck. A warning and a note are reports.

| Code | Count |
|---|---|
| `curve_summary` | 1 |
| `finisher_short` | 1 |

By severity: BLOCK 0. WARN 1. INFO 1. 

## Decks

### 5. blink Commander, owned first

Format: Commander. Theme: blink. Pool: owned_first. Shortlist: 169 names.

Commander: Gilraen, Dúnedain Protector.

Grade: baseline, score 0.33, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $169.60 the whole deck.

**Summary:** Gilraen leads a white blink shell that develops mana, repeatedly leans on creature-based value, and protects its key pieces while applying pressure with Angels, legends, and artifact creatures. The deck closes through a developed creature board, chief among them Angel of Serenity, while spot removal and board resets keep opposing boards manageable. It gives up multicolor flexibility and a wider range of finishers for a consistent Plains-heavy mana base and a focused library-first collection of cards.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

This deck plan aims for 2 or more finishers, the cards that win the game, and the deck holds 1.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core of white value creatures, protection equipment, and removal hangs together around Gilraen's protect-and-blink ability, but a sizable block of cards (Kataki, Grim Poppet, PuPu UFO, Blitzball, Stiltzkin, The Walls of Ba Sing Se) are unconnected filler, and the three symmetrical wraths fight the creature-board plan the summary names.
- PLAN theme_fit=partly: It is a mono-white library-constrained Commander deck at roughly bracket 3 as asked, and Gilraen is a reasonable blink commander, but the blink theme is represented by only about six or seven cards with few ETB payoffs, so the deck plays as generic white value more than blink.
- PLAN useful_as_built=yes: 36 lands plus ten ramp pieces, a dozen draw sources, and nine spot-removal spells give a smooth, castable curve, and while only Angel of Serenity is tagged a finisher, the equipment-and-Angels board with Gilraen can realistically close a game.
- PLAN summary_honest=partly: It fairly admits the mono-white mana, narrow finisher count, and library-first constraint, but calling the list a "blink shell" oversells a deck with only a handful of actual flicker effects and payoffs.
- [INFO] `curve_summary`: average mana value 2.97 over 63 nonland cards
- [WARN] `finisher_short`: the finisher count is 1, and this deck plan wants 2 or more (Angel of Serenity)

<details><summary>The deck list</summary>

- 32 Plains | land | Core white mana base.
- 1 Command Tower | land | Reliable colored mana for Gilraen's deck.
- 1 City of Brass | land | Flexible colored mana source.
- 1 Plaza of Heroes | land | Colored land for the legendary-focused shell.
- 1 Minas Tirith | land | White source with useful deck utility.
- 1 Adventurer's Airship | draw | Card-advantage vehicle from the library.
- 1 Buster Sword | draw | Equipment-based card advantage.
- 1 Champions of Minas Tirith | draw | Creature-based card advantage for blink turns.
- 1 Crown of Gondor | draw | Card-advantage equipment.
- 1 Diary of Dreams | draw | Artifact card-advantage piece.
- 1 Energybending | draw | Efficient card-advantage spell.
- 1 Faramir, Field Commander | draw | Creature-based card advantage.
- 1 Idol of Oblivion | draw | Low-cost card-advantage artifact.
- 1 Inspiring Overseer | draw | Blink-friendly creature-based card advantage.
- 1 Lembas | draw | Cheap artifact card advantage.
- 1 Wall of Omens | draw | Early card advantage on a blinkable creature.
- 1 Sol Ring | ramp | Fast early mana acceleration.
- 1 Arcane Signet | ramp | Reliable mana acceleration.
- 1 Bender's Waterskin | ramp | Artifact mana acceleration.
- 1 Blitzball | ramp | Artifact mana acceleration.
- 1 Chromatic Lantern | ramp | Mana acceleration and fixing.
- 1 Fellwar Stone | ramp | Low-cost mana acceleration.
- 1 Relic of Legends | ramp | Legend-friendly mana acceleration.
- 1 Thought Vessel | ramp | Two-mana artifact acceleration.
- 1 Wayfarer's Bauble | ramp | Early land-based mana acceleration.
- 1 Sword of the Animist | ramp | Equipment-based mana acceleration.
- 1 Banishing Light | removal | Flexible permanent removal.
- 1 Crib Swap | removal | Creature removal that suits the deck's creature focus.
- 1 Destroy Evil | removal | Flexible instant-speed removal.
- 1 Dispatch | removal | Efficient spot removal alongside artifacts.
- 1 Fiend Hunter | removal | Blink-friendly creature removal.
- 1 Generous Gift | removal | Broad permanent removal.
- 1 Get Lost | removal | Versatile spot removal.
- 1 Journey to Nowhere | removal | Low-cost creature removal.
- 1 Swords to Plowshares | removal | Efficient creature removal.
- 1 Bastion Protector | interaction | Protection for the commander plan.
- 1 Boromir, Warden of the Tower | interaction | Creature-based disruption.
- 1 Clever Concealment | interaction | Protection against opposing answers.
- 1 Champion's Helm | interaction | Commander protection equipment.
- 1 Duty Beyond Death | interaction | Protective instant-speed interaction.
- 1 Lightning Greaves | interaction | Cheap protection equipment.
- 1 Reprieve | interaction | Flexible reactive interaction.
- 1 Swiftfoot Boots | interaction | Additional commander protection.
- 1 Ultimate Magic: Holy | interaction | Reactive protection for the board.
- 1 Unbreakable Formation | interaction | Board-protection interaction.
- 1 Angel of Condemnation | synergy | Blink-focused creature that reinforces the deck's central plan.
- 1 Ennis, Debate Moderator | synergy | Supports the deck's blink-focused creature plan.
- 1 Flickerwisp | synergy | Direct blink support on a creature body.
- 1 Jocasta, Automaton Avenger | synergy | Creature synergy piece for the blink shell.
- 1 Personify | synergy | Blink-plan support spell.
- 1 Slip On the Ring | synergy | Instant-speed blink support.
- 1 Sheltered by Ghosts | synergy | Protective support for key creatures.
- 1 Angel of Sanctions | threat | Evasive creature threat that pressures the table.
- 1 Exemplar of Light | threat | Creature threat that advances the board.
- 1 Giada, Font of Hope | threat | Creature threat that supports the Angel contingent.
- 1 Grim Poppet | threat | Artifact creature that adds board pressure.
- 1 Kataki, War's Wage | threat | Creature threat with disruptive pressure.
- 1 PuPu UFO | threat | Artifact creature threat for the battlefield.
- 1 South Pole Voyager | threat | Creature threat that develops the board.
- 1 Stiltzkin, Moogle Merchant | threat | Legendary creature threat for the board.
- 1 The Vision | threat | Artifact creature threat that adds pressure.
- 1 The Walls of Ba Sing Se | threat | Legendary artifact creature for board presence.
- 1 Weapons Vendor | threat | Creature threat that complements the equipment package.
- 1 Zack Fair | threat | Legendary creature threat for combat pressure.
- 1 Angel of Serenity | wincon | Primary finishing creature for closing games.
- 1 Austere Command | wipe | Flexible reset for difficult boards.
- 1 Fumigate | wipe | Creature-board reset.
- 1 Vanquish the Horde | wipe | Efficient creature-board reset.

</details>

