# PR-8 deck gate

Run date: 2026-09-14. Card snapshot: 2026-09-04.

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
| Mean plan score, 0 to 1 | 0.88 |
| Plan reasons the judge left empty | 0 |
| Errors | 0 |
| Prompt version | 13 |
| Calls | 3 |
| Cost | $0.1321 |
| Time | 94 seconds |

## Run

- Suite `decks`, run `pr8-deck-gate-run20`, on 2026-09-14, commit `5cdc4cc`.
- Partial run over `3`. It reads the bars of its own items, and it never stands for the suite.
- Roles: generate on `gpt-5.6-terra` (openai, effort medium), judge on `claude-opus-5` (anthropic, effort medium), repair on `gpt-5.6-terra` (openai, effort medium).
- Versions: card snapshot 2026-09-04, generate prompt version 13, plan_rubric prompt version 2, quality_model `20260914T154223Z`.
- Calls: 3. Cost: $0.1321. Time: 94 seconds.

## Findings by code

A block stops the deck. A warning and a note are reports.

| Code | Count |
|---|---|
| `curve_summary` | 1 |

By severity: BLOCK 0. WARN 0. INFO 1. 

## Decks

### 3. artifacts Commander, bracket 4

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 302 names.

Commander: Urza, Lord High Artificer.

Grade: typical, score 0.63, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $4187.38 to buy, $4187.38 the whole deck.

**Summary:** This Urza deck builds a dense artifact board quickly with fast mana, then uses artifact synergy, card draw, and tutors to keep its momentum through longer games. It wins by deploying a succession of artifact threats and turning its established board into sustained pressure, while a deep suite of interaction and removal protects key turns. The deck gives up land utility for a highly reliable blue mana base and prioritizes artifact-focused cards over broader multicolor options.

The quality model grades this deck typical against the top lists of the format: the deck makes more mana on turn four than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

- PLAN plan_coherent=yes: Nearly everything supports one line — cheap artifacts and fast mana to fuel Urza's construct and his tap-for-mana ability, plus affinity/improvise draw, artifact tutors and artifact beaters — with only a couple of loose ends (Lion's Eye Diamond with no discard payoff, Shape Anew in a list with few big artifact creatures) that don't actively fight the plan.
- PLAN theme_fit=partly: It is genuinely a mono-blue Urza artifact Commander deck, but for a stated bracket 4 build the removal suite is padded with low-impact blue auras and one-shot effects (Water Whip, Unable to Scream, Watery Grasp, Into Thin Air, Zuko's Exile) and there is no fast combo finish, so the power level sits below what was asked.
- PLAN useful_as_built=yes: 33 Islands in mono-blue plus a large fast-mana package makes casting trivially reliable, the curve is low, and Urza's construct alongside Kappa Cannoneer, Darksteel Juggernaut, Cyberdrive Awakener and other beaters give real ways to close a game.
- PLAN summary_honest=yes: The claims — fast mana, artifact synergy draw, tutors, a wide interaction/removal count, artifact threats as the win route, and no utility lands in favor of basics — all match the list, and it does not promise a combo kill the deck lacks.
- [INFO] `curve_summary`: average mana value 2.89 over 66 nonland cards

<details><summary>The deck list</summary>

- 33 Island | land | Provides reliable blue mana for the deck's blue spells.
- 1 Chrome Mox | ramp | Fast artifact mana that accelerates the artifact plan.
- 1 Lion's Eye Diamond | ramp | Fast artifact mana for explosive starts.
- 1 Mana Vault | ramp | Fast artifact mana that pushes out major plays early.
- 1 Mox Diamond | ramp | Fast artifact mana that supports early development.
- 1 Mox Opal | ramp | Fast artifact mana in a deck built around artifacts.
- 1 Lotus Petal | ramp | Fast artifact mana for early acceleration.
- 1 Sol Ring | ramp | Efficient artifact ramp for deploying the deck's core pieces.
- 1 Mox Amber | ramp | Fast artifact mana that supports the legendary commander plan.
- 1 Moonsnare Prototype | ramp | Low-cost artifact ramp that fits the deck's artifact density.
- 1 Springleaf Drum | ramp | Fast artifact-based acceleration.
- 1 Metalworker | ramp | Artifact-focused ramp for powering out threats and engines.
- 1 Krark-Clan Ironworks | ramp | Artifact ramp that supports the deck's resource-heavy turns.
- 1 Inspiring Statuary | ramp | Artifact ramp that turns the board into mana support.
- 1 Rhystic Study | draw | A Game Changer draw engine that keeps resources flowing.
- 1 Thoughtcast | draw | Artifact-oriented card draw.
- 1 Thought Monitor | draw | An artifact threat that also fills the draw role.
- 1 Thirst for Knowledge | draw | Efficient draw for an artifact-heavy hand.
- 1 Sai, Master Thopterist | draw | Artifact-focused draw support.
- 1 Vedalken Archmage | draw | A draw engine aligned with the artifact plan.
- 1 Riddlesmith | draw | Low-cost artifact-related card selection and draw.
- 1 Reverse Engineer | draw | Artifact-oriented draw to reload after deploying the hand.
- 1 Forensic Gadgeteer | draw | Draw support within the artifact strategy.
- 1 Nexus of Becoming | draw | Artifact-based card draw.
- 1 Tezzeret, Artifice Master | draw | A draw-focused planeswalker for the artifact shell.
- 1 One with the Machine | draw | A high-impact draw spell for an artifact deck.
- 1 Fierce Guardianship | interaction | A Game Changer interaction spell that protects the plan.
- 1 Force of Will | interaction | A Game Changer interaction spell for defending key turns.
- 1 Metallic Rebuke | interaction | Artifact-aligned stack interaction.
- 1 Disruption Protocol | interaction | Low-cost interaction supported by artifacts.
- 1 Stoic Rebuttal | interaction | Reliable artifact-oriented interaction.
- 1 Ice Out | interaction | Flexible instant-speed interaction.
- 1 An Offer You Can't Refuse | interaction | Cheap interaction that can also contribute fast mana.
- 1 Escape Protocol | interaction | Interaction that fits the artifact-heavy game plan.
- 1 Fugitive Droid | interaction | Artifact creature interaction that adds to the board.
- 1 Welding Jar | interaction | Cheap artifact-focused protection.
- 1 Etched Champion | interaction | An artifact threat that also occupies an interaction role.
- 1 Aether Spellbomb | removal | Low-cost artifact-based removal.
- 1 Contagion Clasp | removal | Artifact removal that supports the deck's board control.
- 1 Cyber Conversion | removal | Efficient removal for problematic permanents.
- 1 Into Thin Air | removal | Low-cost removal to disrupt opposing boards.
- 1 Resculpt | removal | Flexible removal at instant speed.
- 1 Unable to Scream | removal | Low-cost permanent-based removal.
- 1 Watery Grasp | removal | Removal that handles an opposing problem piece.
- 1 Water Whip | removal | Additional removal coverage.
- 1 Zuko's Exile | removal | Additional low-cost removal.
- 1 Shape Anew | removal | Removal that also fits the artifact shell.
- 1 Engineered Explosives | wipe | A flexible artifact board wipe.
- 1 Rebuild | wipe | A reset option against artifact-heavy boards.
- 1 Whir of Invention | synergy | Artifact synergy and a tutor for important pieces.
- 1 Transmute Artifact | synergy | Artifact synergy and a tutor for key artifacts.
- 1 Reshape | synergy | Artifact synergy and a tutor for the deck's core tools.
- 1 Mystic Forge | synergy | An artifact synergy engine for sustained development.
- 1 Clock of Omens | synergy | Supports artifact-based synergy turns.
- 1 Unwinding Clock | synergy | An artifact synergy piece for maintaining pressure.
- 1 Frogmite | threat | A low-cost artifact threat for applying early pressure.
- 1 Chrome Steed | threat | An artifact creature that contributes to board pressure.
- 1 Filigree Attendant | threat | An artifact threat that supports the deck's artifact density.
- 1 Jhoira's Familiar | threat | An artifact threat that complements the deck's core plan.
- 1 Lodestone Golem | threat | A resilient artifact threat for pressuring opponents.
- 1 Master Transmuter | threat | A major artifact threat for the deck's top end.
- 1 Karn, Scion of Urza | threat | A threat that rewards a board full of artifacts.
- 1 Traxos, Scourge of Kroog | threat | A large artifact threat that closes games through combat.
- 1 Phyrexian Metamorph | threat | A flexible artifact threat for adapting to the table.
- 1 Kappa Cannoneer | threat | A powerful artifact threat for finishing games.
- 1 Cyberdrive Awakener | threat | An artifact threat that converts the board into a closing attack.
- 1 Darksteel Juggernaut | threat | A scaling artifact threat for combat finishes.

</details>

