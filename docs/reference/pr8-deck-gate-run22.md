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
| Mean plan score, 0 to 1 | 1.00 |
| Plan reasons the judge left empty | 1 |
| Errors | 0 |
| Prompt version | 13 |
| Calls | 3 |
| Cost | $0.1159 |
| Time | 106 seconds |

## Run

- Suite `decks`, run `pr8-deck-gate-run22`, on 2026-09-14, commit `8ca9e58`.
- Partial run over `3`. It reads the bars of its own items, and it never stands for the suite.
- Roles: generate on `gpt-5.6-terra` (openai, effort medium), judge on `claude-opus-5` (anthropic, effort medium), repair on `gpt-5.6-terra` (openai, effort medium).
- Versions: card snapshot 2026-09-04, generate prompt version 13, plan_rubric prompt version 2, quality_model `20260914T154223Z`.
- Calls: 3. Cost: $0.1159. Time: 106 seconds.

## Findings by code

A block stops the deck. A warning and a note are reports.

| Code | Count |
|---|---|
| `curve_summary` | 1 |
| `mana_pass` | 1 |

By severity: BLOCK 0. WARN 0. INFO 2. 

## Decks

### 3. artifacts Commander, bracket 4

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 317 names.

Commander: Urza, Lord High Artificer.

Grade: good, score 0.73, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $7791.09 to buy, $7791.09 the whole deck.

**Summary:** This is an artifact-dense, high-power deck that accelerates quickly, uses efficient card advantage and interaction to maintain momentum, and develops a board of artifact threats backed by compact win conditions. Its strongest games assemble mana and artifact synergy early, then protect a decisive threat or winning line with low-cost answers. The tradeoff is that many slots are specialized around artifacts, so the deck is most effective when its artifact engines remain established.

The quality model grades this deck good against the top lists of the format: the deck makes more mana on turn four than the norm, and that raises the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

- PLAN plan_coherent=yes: The list is one plan: cheap artifacts and fast mana rocks feed Urza's construct and his mana ability, affinity/improvise payoffs and tutors (Whir, Reshape, Transmute Artifact) assemble engines, and Power Artifact plus Mana Vault/Urza gives an infinite-mana line into Brain Freeze/Thassa's Oracle, with Mirrodin Besieged and Mechanized Production as backups; a few clunky slots (Skysovereign, Jhoira's Familiar, Shimmer Dragon) are outliers but do not pull against the plan.
- PLAN theme_fit=yes: It is exactly what was asked: a mono-blue Commander deck helmed by Urza, Lord High Artificer, densely artifact-based, with fast mana, free counterspells and tutors appropriate to a high-power bracket-4 build.
- PLAN useful_as_built=yes: 34 lands in a mono-blue shell plus a dozen cheap rocks means the mana works, the curve is low, and there are multiple redundant ways to close (combo mill into Oracle, Mirrodin Besieged, Mechanized Production, Kappa Cannoneer/Cyberdrive Awakener beats), so it can be picked up and played.
- PLAN summary_honest=yes: placeholder
- [INFO] `curve_summary`: average mana value 2.91 over 65 nonland cards
- [INFO] `mana_pass`: the builder moved 2 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 22 Island | land | Provides reliable blue mana for the deck's blue spells.
- 1 Academy Ruins | land | Utility land slot for the artifact-focused mana base.
- 1 Inventors' Fair | land | Utility land slot that supports the artifact plan.
- 1 Mishra's Workshop | land | High-impact artifact-focused mana source.
- 1 Urza's Saga | land | Artifact-oriented utility land.
- 1 Power Depot | land | Artifact land that contributes to the mana base.
- 1 Seat of the Synod | land | Blue artifact land for both mana and artifact density.
- 1 Mystic Sanctuary | land | Blue-producing utility land.
- 1 Otawara, Soaring City | land | Blue-producing utility land.
- 1 Sink into Stupor // Soporific Springs | land | Flexible blue land slot.
- 1 Glimmervoid | land | Colored artifact-supporting land.
- 1 Spire of Industry | land | Colored land that fits an artifact-heavy mana base.
- 1 Mana Confluence | land | Flexible colored source for the mana base.
- 1 Chrome Mox | ramp | Fast mana that accelerates early development.
- 1 Lion's Eye Diamond | ramp | Fast mana for explosive artifact turns.
- 1 Lotus Petal | ramp | Fast mana that helps deploy early plays.
- 1 Mana Vault | ramp | Fast mana that pushes out expensive artifact plays.
- 1 Mox Diamond | ramp | Fast mana that improves early access to mana.
- 1 Mox Opal | ramp | Fast mana suited to the deck's artifact density.
- 1 Moonsnare Prototype | ramp | Fast mana for the artifact-based game plan.
- 1 Sol Ring | ramp | Efficient artifact acceleration.
- 1 Springleaf Drum | ramp | Fast mana that works with the deck's creatures.
- 1 Metalworker | ramp | Artifact-focused mana production.
- 1 Krark-Clan Ironworks | ramp | Artifact-based mana engine.
- 1 Grand Architect | ramp | Creature-based acceleration for the artifact plan.
- 1 Inspiring Statuary | ramp | Artifact-focused ramp for larger turns.
- 1 Rhystic Study | draw | Premium ongoing card advantage.
- 1 Thoughtcast | draw | Efficient draw supported by artifact density.
- 1 Thought Monitor | draw | Artifact creature that contributes card advantage.
- 1 Sai, Master Thopterist | draw | Artifact-focused source of card advantage.
- 1 Thirst for Knowledge | draw | Efficient draw that fits an artifact-heavy hand.
- 1 Reverse Engineer | draw | Artifact-supported card draw.
- 1 Riddlesmith | draw | Artifact-oriented draw and card selection.
- 1 Vedalken Archmage | draw | Ongoing draw for casting artifacts.
- 1 Forensic Gadgeteer | draw | Artifact-focused card advantage.
- 1 One with the Machine | draw | High-output draw option for an artifact deck.
- 1 Force of Will | interaction | Premium free interaction for protecting the deck's tempo.
- 1 Fierce Guardianship | interaction | High-impact interaction while the commander is established.
- 1 Metallic Rebuke | interaction | Efficient artifact-supported interaction.
- 1 Disruption Protocol | interaction | Artifact-supported stack interaction.
- 1 An Offer You Can't Refuse | interaction | Low-cost interaction that can also provide fast mana.
- 1 Ice Out | interaction | Flexible instant-speed interaction.
- 1 Stoic Rebuttal | interaction | Artifact-supported countermagic.
- 1 Welding Jar | interaction | Cheap protection for important artifacts.
- 1 Padeem, Consul of Innovation | interaction | Protective interaction for the artifact board.
- 1 Shimmer Dragon | interaction | Artifact-based interaction attached to a threat.
- 1 Aether Spellbomb | removal | Cheap artifact removal option.
- 1 Cyber Conversion | removal | Flexible removal for problematic permanents.
- 1 Resculpt | removal | Efficient removal with broad utility.
- 1 Ravenform | removal | Removal that answers key targets.
- 1 Unable to Scream | removal | Low-cost permanent-based removal.
- 1 Into Thin Air | removal | Instant-speed removal option.
- 1 Kitesail Larcenist | removal | Creature-based removal for the artifact strategy.
- 1 Skysovereign, Consul Flagship | removal | Artifact threat that also supplies removal.
- 1 Engineered Explosives | wipe | Flexible artifact board reset.
- 1 Hurkyl's Recall | wipe | Artifact-focused reset for disruptive board states.
- 1 Kappa Cannoneer | threat | Powerful artifact-based combat threat.
- 1 Cyberdrive Awakener | threat | Artifact-focused finisher threat.
- 1 Karn, Scion of Urza | threat | Planeswalker threat that fits the artifact plan.
- 1 Master Transmuter | threat | Artifact creature threat with strong deck synergy.
- 1 Phyrexian Metamorph | threat | Flexible artifact creature threat.
- 1 Myr Enforcer | threat | Artifact-density threat that supports the board.
- 1 Sojourner's Companion | threat | Artifact creature threat for the affinity-style shell.
- 1 Frogmite | threat | Efficient artifact creature threat.
- 1 Lodestone Golem | threat | Disruptive artifact creature threat.
- 1 Jhoira's Familiar | threat | Artifact creature threat that supports artifact deployment.
- 1 Whir of Invention | synergy | Artifact tutor that finds key synergy pieces.
- 1 Transmute Artifact | synergy | Artifact tutor that converts resources into a needed piece.
- 1 Reshape | synergy | Artifact tutor for assembling the deck's core pieces.
- 1 Power Artifact | synergy | Core artifact synergy piece.
- 1 Mystic Forge | synergy | Artifact-focused engine for sustained development.
- 1 Emry, Lurker of the Loch | synergy | Artifact synergy creature that supports recurring value.
- 1 Mystical Tutor | other | Tutor that finds a needed instant or sorcery.
- 1 Brain Freeze | other | Dedicated utility card for the deck's high-power plan.
- 1 Thassa's Oracle | wincon | Compact dedicated win condition.
- 1 Mirrodin Besieged | wincon | Artifact-themed dedicated win condition.
- 1 Mechanized Production | wincon | Artifact-themed dedicated win condition.
- 1 Mox Amber | ramp | the mana pass added it to bring the mana base inside the power level

</details>

