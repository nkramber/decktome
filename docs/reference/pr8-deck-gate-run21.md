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
| Mean plan score, 0 to 1 | 0.75 |
| Plan reasons the judge left empty | 0 |
| Errors | 0 |
| Prompt version | 13 |
| Calls | 3 |
| Cost | $0.1174 |
| Time | 97 seconds |

## Run

- Suite `decks`, run `pr8-deck-gate-run21`, on 2026-09-14, commit `f45791a`.
- Partial run over `3`. It reads the bars of its own items, and it never stands for the suite.
- Roles: generate on `gpt-5.6-terra` (openai, effort medium), judge on `claude-opus-5` (anthropic, effort medium), repair on `gpt-5.6-terra` (openai, effort medium).
- Versions: card snapshot 2026-09-04, generate prompt version 13, plan_rubric prompt version 2, quality_model `20260914T154223Z`.
- Calls: 3. Cost: $0.1174. Time: 97 seconds.

## Findings by code

A block stops the deck. A warning and a note are reports.

| Code | Count |
|---|---|
| `mana_pass` | 1 |
| `curve_summary` | 1 |

By severity: BLOCK 0. WARN 0. INFO 2. 

## Decks

### 3. artifacts Commander, bracket 4

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 317 names.

Commander: Urza, Lord High Artificer.

Grade: good, score 0.70, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $7628.00 to buy, $7628.00 the whole deck.

**Summary:** This is a fast artifact-centered Urza deck that develops mana early, uses artifact engines and tutors to assemble its strongest pieces, and protects its turns with dense stack interaction. It can close through Thassa's Oracle or by applying sustained pressure with artifact threats, while its flexible removal and resets keep opposing boards from taking over. The tradeoff is that several of its most powerful cards depend on maintaining a substantial artifact presence, so well-timed disruption can slow its momentum.

The quality model grades this deck good against the top lists of the format: the deck makes more mana on turn four than the norm, and that raises the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

- PLAN plan_coherent=partly: Nearly every card supports one artifact-density plan under Urza — fast mana, artifact payoffs, tutors, and counterspells — but the lone Thassa's Oracle has no library-emptying support and sits outside the plan the rest of the list carries.
- PLAN theme_fit=yes: It is exactly the requested mono-blue high-power artifact deck helmed by Urza, Lord High Artificer, with a card pool (Workshop, LED, Mox suite, free counters) appropriate to bracket 4.
- PLAN useful_as_built=yes: 35 lands plus a deep fast-mana package support the curve, the artifact count comfortably turns on affinity/improvise cards, and it can actually win via Urza tokens, Kappa Cannoneer, Cyberdrive Awakener, and Aetherflux Reservoir even without the Oracle line.
- PLAN summary_honest=partly: The descriptions of fast mana, tutors, dense counterspells, and artifact pressure are all accurate, but naming Thassa's Oracle as a way to close is false: there is no Demonic Consultation, Tainted Pact, or self-mill in the list, so the Oracle is a near-blank.
- [INFO] `curve_summary`: average mana value 2.97 over 65 nonland cards
- [INFO] `mana_pass`: the builder moved 3 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 23 Island | land | Primary blue source for a reliable untapped mana base.
- 1 Academy Ruins | land | Utility land slot in the artifact-focused mana base.
- 1 Buried Ruin | land | Utility land slot in the artifact-focused mana base.
- 1 Inventors' Fair | land | Utility land slot in the artifact-focused mana base.
- 1 Mishra's Workshop | land | High-impact artifact mana-base land.
- 1 Mistrise Village | land | Blue utility land for the mana base.
- 1 Mystic Sanctuary | land | Island-based utility land slot.
- 1 Otawara, Soaring City | land | Blue utility land for the mana base.
- 1 Seat of the Synod | land | Artifact land that supports the artifact plan.
- 1 Sink into Stupor // Soporific Springs | land | Flexible blue land slot.
- 1 Spire of Industry | land | Artifact-focused colored land slot.
- 1 Urza's Saga | land | Artifact-focused utility land.
- 1 Chrome Mox | ramp | Fast mana that accelerates the opening turns.
- 1 Lion's Eye Diamond | ramp | Fast mana for explosive artifact turns.
- 1 Lotus Petal | ramp | Fast mana that helps deploy early plays.
- 1 Mana Vault | ramp | Fast mana for major early acceleration.
- 1 Mox Diamond | ramp | Fast mana that improves early development.
- 1 Mox Opal | ramp | Fast mana for an artifact-dense deck.
- 1 Moonsnare Prototype | ramp | Low-cost artifact acceleration.
- 1 Sol Ring | ramp | Efficient early artifact acceleration.
- 1 Springleaf Drum | ramp | Low-cost mana production alongside artifact creatures.
- 1 Grand Architect | ramp | Artifact-focused mana acceleration.
- 1 Inspiring Statuary | ramp | Artifact-focused mana acceleration.
- 1 Krark-Clan Ironworks | ramp | Artifact mana engine for powerful turns.
- 1 Metalworker | ramp | Artifact-focused mana acceleration.
- 1 Forensic Gadgeteer | draw | Artifact-focused card advantage.
- 1 Nexus of Becoming | draw | Artifact card-advantage piece.
- 1 Reverse Engineer | draw | Artifact-focused card draw.
- 1 Rhystic Study | draw | Premium ongoing card advantage.
- 1 Riddlesmith | draw | Artifact-focused card selection and draw.
- 1 Sai, Master Thopterist | draw | Artifact-focused card advantage.
- 1 Tezzeret, Artifice Master | draw | Card advantage and a tutor option.
- 1 Thirst for Knowledge | draw | Efficient instant-speed card draw.
- 1 Thoughtcast | draw | Efficient artifact-focused card draw.
- 1 Transmutation Font | draw | Card draw with a tutor option.
- 1 Vedalken Archmage | draw | Artifact-focused card advantage.
- 1 An Offer You Can't Refuse | interaction | Efficient stack interaction with fast-mana upside.
- 1 Disruption Protocol | interaction | Artifact-focused stack interaction.
- 1 Fierce Guardianship | interaction | Premium protective interaction.
- 1 Force of Will | interaction | Premium free stack interaction.
- 1 Ice Out | interaction | Flexible stack interaction.
- 1 Metallic Rebuke | interaction | Efficient artifact-focused interaction.
- 1 Override | interaction | Low-cost stack interaction.
- 1 Padeem, Consul of Innovation | interaction | Artifact-focused protective interaction.
- 1 Stoic Rebuttal | interaction | Artifact-focused stack interaction.
- 1 Welding Jar | interaction | Low-cost artifact protection.
- 1 Aether Spellbomb | removal | Cheap artifact-based removal.
- 1 Aetherflux Reservoir | removal | Artifact removal option and proactive payoff.
- 1 Cyber Conversion | removal | Flexible targeted removal.
- 1 Kitesail Larcenist | removal | Artifact-compatible creature removal.
- 1 Ravenform | removal | Flexible removal spell.
- 1 Resculpt | removal | Efficient flexible removal.
- 1 Skysovereign, Consul Flagship | removal | Artifact threat that supplies removal.
- 1 Spine of Ish Sah | removal | Artifact-based universal removal.
- 1 Engineered Explosives | wipe | Flexible artifact board wipe.
- 1 Hurkyl's Recall | wipe | Efficient artifact-focused board reset.
- 1 Clock of Omens | synergy | Artifact synergy engine.
- 1 Mystic Forge | synergy | Central artifact synergy engine.
- 1 Reshape | synergy | Artifact tutor that advances the synergy plan.
- 1 Transmute Artifact | synergy | Artifact tutor for key synergy pieces.
- 1 Unwinding Clock | synergy | Artifact synergy engine for extended turns.
- 1 Whir of Invention | synergy | Artifact tutor that finds key synergy pieces.
- 1 Arcbound Crusher | threat | Artifact creature threat.
- 1 Cyberdrive Awakener | threat | Artifact-focused finisher threat.
- 1 Darksteel Juggernaut | threat | Large artifact creature threat.
- 1 Filigree Attendant | threat | Efficient artifact creature threat.
- 1 Frogmite | threat | Low-cost artifact creature threat.
- 1 Kappa Cannoneer | threat | High-impact artifact creature threat.
- 1 Karn, Scion of Urza | threat | Artifact-focused planeswalker threat.
- 1 Kuldotha Forgemaster | threat | Artifact threat with a tutor option.
- 1 Master Transmuter | threat | Artifact-focused creature threat.
- 1 Phyrexian Metamorph | threat | Flexible artifact creature threat.
- 1 Traxos, Scourge of Kroog | threat | Large artifact creature threat.
- 1 Whirler Rogue | threat | Artifact-focused creature threat.
- 1 Thassa's Oracle | wincon | Compact dedicated win condition.
- 1 Mox Amber | removal | the mana pass added it to bring the mana base inside the power level
- 1 Mystical Tutor | draw | the mana pass added it to bring the mana base inside the power level

</details>

