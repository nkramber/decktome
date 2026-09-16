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
| Summaries that state a rule of the game | 1 |
| Summaries that state a FALSE rule | 0 |
| Judge errors | 0 |
| Case assertions missed (PR-28b) | 0 |
| Decks the plan judge read (PR-15, information) | 1 |
| Mean plan score, 0 to 1 | 0.50 |
| Plan reasons the judge left empty | 0 |
| Errors | 0 |
| Prompt version | 15 |
| Calls | 3 |
| Cost | $0.1292 |
| Time | 96 seconds |

## Run

- Suite `decks`, run `pr8-deck-gate-run26`, on 2026-09-16, commit `6ff0244`.
- Partial run over `5`. It reads the bars of its own items, and it never stands for the suite.
- Roles: generate on `gpt-5.6-terra` (openai, effort medium), judge on `claude-opus-5` (anthropic, effort medium), repair on `gpt-5.6-terra` (openai, effort medium).
- Versions: card snapshot 2026-09-04, generate prompt version 15, plan_rubric prompt version 2, quality_model `20260914T154223Z`.
- Calls: 3. Cost: $0.1292. Time: 96 seconds.

## Findings by code

A block stops the deck. A warning and a note are reports.

| Code | Count |
|---|---|
| `profile_off_band` | 2 |
| `curve_summary` | 1 |
| `finisher_short` | 1 |
| `basics_added` | 1 |

By severity: BLOCK 0. WARN 3. INFO 2. 

## Decks

### 5. blink Commander, owned first

Format: Commander. Theme: blink. Pool: owned_first. Shortlist: 169 names.

Commander: Gilraen, Dúnedain Protector.

Grade: typical, score 0.42, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $126.28 the whole deck.

**Summary:** Gilraen, Dúnedain Protector leads a white blink deck built around repeatedly using creature-based draw, removal, and synergy pieces while equipment and instant interaction help maintain its board. Angel of Serenity is chief among the Angels and provides the deck’s primary finishing route, while the wide removal suite and three sweepers let it play a measured, board-controlling game. No other card leads it; the tradeoff is that the deck leans heavily on creatures and has less room for additional dedicated finishers or blink synergy from the available library.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

This deck plan aims for 2 or more finishers, the cards that win the game, and the deck holds 1.

Summary rules claims (F-26): the format play

- JUDGE [true]: "Gilraen, Dúnedain Protector leads a white blink deck". Asserts this card can serve as the deck's commander. Gilraen, Dúnedain Protector is a white legendary creature, so it is a legal commander, and a white deck matches its color identity.
- JUDGE [unknown]: "No other card leads it.". Asserts that no other card in the decklist is eligible to be a commander (i.e., no other legendary creature or 'can be your commander' card). Without the decklist this cannot be verified.
- PLAN plan_coherent=partly: The list hangs together as a white attrition deck (heavy targeted removal, three sweepers, ETB value creatures, protection pieces that can target your own creatures to trigger Gilraen), but the stated blink plan is only lightly supported — few actual flicker enablers and only a handful of ETB payoffs — and the equipment/Puresteel-Skullclamp package and sweepers pull somewhat against a creature-value board.
- PLAN theme_fit=partly: It is a mono-white Commander deck at a plausible bracket-3 power level drawn from the owner's library, but the requested blink theme is largely absent — Angel of Condemnation, Fiend Hunter, Flickerwisp and Gilraen's own trigger are nearly the whole package, with the rest being generic removal and protection.
- PLAN useful_as_built=partly: The mana base and curve are sound and the deck can interact all game, but with one real finisher and a board of small utility creatures it has few credible paths to actually closing a game.
- PLAN summary_honest=partly: It is commendably frank about the single finisher and the thin blink synergy, but it still opens by calling this a deck "built around" blink and gestures at an Angel subtheme that amounts to three or four Angels, which overstates what the list actually does.
- [INFO] `curve_summary`: average mana value 2.81 over 62 nonland cards
- [WARN] `profile_off_band`: the removal count is 17, and bracket 3 wants 6 to 12
- [WARN] `profile_off_band`: the interaction count is 14, and bracket 3 wants 4 to 12
- [WARN] `finisher_short`: the finisher count is 1, and this deck plan wants 2 or more (Angel of Serenity)
- [INFO] `basics_added`: the list was 1 card short, so the builder added 1 basic land

<details><summary>The deck list</summary>

- 31 Plains | land | Reliable white mana for the blink shell.
- 1 Command Tower | land | A color-fixing land for the commander’s deck.
- 1 City of Brass | land | Additional color-fixing land support.
- 1 Plaza of Heroes | land | A land slot that supports the legendary commander plan.
- 1 Secluded Courtyard | land | A creature-focused color source for the deck.
- 1 Marsh Flats | land | A land slot that helps maintain access to Plains.
- 1 Fabled Passage | land | A land slot that helps maintain access to Plains.
- 1 Sol Ring | ramp | Efficient artifact ramp for earlier development.
- 1 Arcane Signet | ramp | Reliable artifact ramp in the commander’s colors.
- 1 Fellwar Stone | ramp | Low-cost artifact ramp.
- 1 Bender's Waterskin | ramp | Artifact ramp that supports the deck’s mana development.
- 1 Wayfarer's Bauble | ramp | Early ramp support that works well with the Plains-heavy mana base.
- 1 Sword of the Animist | ramp | Equipment-based ramp for the creature plan.
- 1 Giada, Font of Hope | ramp | Creature-based ramp that fits the Angel contingent.
- 1 Chromatic Lantern | ramp | Artifact ramp and color support.
- 1 Thought Vessel | ramp | Artifact ramp for steady mana development.
- 1 Relic of Legends | ramp | Artifact ramp that complements a legendary commander.
- 1 Champions of Minas Tirith | draw | A draw-tagged creature that supplies a blink-friendly body.
- 1 Diary of Dreams | draw | Dedicated card-draw support.
- 1 Energybending | draw | Low-cost draw support.
- 1 Exemplar of Light | draw | An Angel draw creature for the blink plan.
- 1 Faramir, Field Commander | draw | A draw-tagged Human creature that suits the creature shell.
- 1 Inspiring Overseer | draw | An Angel draw creature that is useful to blink.
- 1 Joined Researchers // Secret Rendezvous | draw | Draw support with a creature face for the blink shell.
- 1 Lembas | draw | Artifact-based card-draw support.
- 1 Mask of Memory | draw | Equipment-based card-draw support for the creature plan.
- 1 Puresteel Paladin | draw | A draw-tagged creature that complements the equipment package.
- 1 Skullclamp | draw | Efficient equipment-based card-draw support.
- 1 Stiltzkin, Moogle Merchant | draw | A draw-tagged creature for the blink-focused shell.
- 1 Wall of Omens | draw | A draw-tagged creature that is a natural blink target.
- 1 Aang, the Last Airbender | removal | Creature-based removal that supports the blink plan.
- 1 Angel of Condemnation | removal | An Angel removal creature for repeated creature-focused play.
- 1 Angel of Sanctions | removal | An Angel removal creature that fits the deck’s theme.
- 1 Banishing Light | removal | Reliable enchantment-based removal.
- 1 Crib Swap | removal | Instant-speed removal support.
- 1 Destroy Evil | removal | Flexible instant removal.
- 1 Dispatch | removal | Low-cost instant removal.
- 1 Fiend Hunter | removal | Creature-based removal and a strong blink-shell fit.
- 1 Generous Gift | removal | Flexible instant removal coverage.
- 1 Get Lost | removal | Efficient removal coverage.
- 1 Journey to Nowhere | removal | Additional enchantment-based removal.
- 1 Make Your Move | removal | Instant removal support.
- 1 March of Otherworldly Light | removal | Flexible instant removal.
- 1 Palace Jailer | removal | Creature-based removal for the blink shell.
- 1 Swords to Plowshares | removal | Efficient instant removal.
- 1 Stroke of Midnight | removal | Broad instant removal coverage.
- 1 Summon: Primal Garuda | removal | Creature-based removal that contributes to the board.
- 1 Bastion Protector | interaction | Creature-based interaction supporting the commander plan.
- 1 Boromir, Warden of the Tower | interaction | Legendary creature interaction for the board-focused shell.
- 1 Clever Concealment | interaction | Instant interaction for protecting the deck’s board presence.
- 1 Darksteel Plate | interaction | Equipment-based interaction for a creature deck.
- 1 Duty Beyond Death | interaction | Low-cost instant interaction.
- 1 Gift of Immortality | interaction | Aura-based interaction that suits the creature plan.
- 1 Lightning Greaves | interaction | Efficient equipment interaction for key creatures.
- 1 Reprieve | interaction | Flexible instant interaction.
- 1 Sheltered by Ghosts | interaction | Aura-based interaction for the board-focused strategy.
- 1 Slip On the Ring | interaction | Instant interaction that directly supports the blink plan.
- 1 Swiftfoot Boots | interaction | Equipment interaction for the commander and creature suite.
- 1 Together Forever | interaction | Enchantment-based interaction for the creature shell.
- 1 Ultimate Magic: Holy | interaction | Instant interaction support.
- 1 Unbreakable Formation | interaction | Instant interaction for maintaining the board.
- 1 Ennis, Debate Moderator | synergy | A dedicated synergy creature for the requested plan.
- 1 Jocasta, Automaton Avenger | synergy | A dedicated synergy creature in the legendary creature shell.
- 1 Flickerwisp | synergy | A dedicated blink synergy creature.
- 1 Personify | synergy | Dedicated synergy support for the blink strategy.
- 1 Angel of Serenity | wincon | The shortlist’s finisher-marked win condition and chief among the Angels.
- 1 Austere Command | wipe | Flexible board-wipe coverage.
- 1 Fumigate | wipe | A dependable board wipe.
- 1 Vanquish the Horde | wipe | Additional board-wipe coverage.

</details>

