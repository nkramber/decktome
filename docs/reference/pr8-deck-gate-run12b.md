# PR-8 deck gate

Run date: 2026-09-02. Card snapshot: 2026-09-02.

Verdict: PASS. 2 of 2 decks passed every block check, 0 invented names reached the user, and 0 summaries stated a false rule of the game. All three bars are zero tolerance.

## Summary

| Measure | Value |
|---|---|
| Prompts | 2 |
| Decks returned | 2 |
| Decks with no block finding | 2 |
| Invented names that reached the user | 0 |
| Decks that needed the repair turn | 1 |
| Summaries judged (F-26) | 2 |
| Summaries that state a rule of the game | 0 |
| Summaries that state a FALSE rule | 0 |
| Judge errors | 0 |
| Errors | 0 |
| Prompt version | 12 |
| Calls | 6 |
| Cost | $0.2605 |
| Time | 266 seconds |

## Findings by code

A block stops the deck. A warning and a note are reports.

| Code | Count |
|---|---|
| `curve_summary` | 2 |
| `profile_off_band` | 2 |

By severity: BLOCK 0. WARN 2. INFO 2. 

## The set filter (PR-17B)

A deck asked for a set holds cards of that set family alone. Three things earn an exception, and each one is marked: a card the reader named (D-381), a mana card the reader allowed from outside (D-382), and a basic land, which no set limit filters (D-378).

| # | Prompt | Sets | In set | Fill | Deck cards outside | Marked |
|---|---|---|---|---|---|---|
| 23 | two set families at once | `blb,blc,hob,hoc,pblb` | 201 | 0 | 0 | 0 |

## Decks

### 3. artifacts Commander, bracket 4

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 299 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band, profile_off_band. Block findings: 0.

Cost: $3974.63 to buy, $3974.63 the whole deck.

**Summary:** A high-power mono-blue artifact strategy that converts a dense field of cheap artifacts into exceptional mana, sustained card flow, and dominating board presence. It protects its engines with efficient stack interaction, disrupts opposing development through flexible answers, and closes games through resilient artifact pressure, recursive value, and explosive mana turns.

- [INFO] `curve_summary`: average mana value 3.12 over 66 nonland cards
- [WARN] `profile_off_band`: the average mana value of the nonland cards is 3.12, and bracket 4 wants 1.6 to 3
- [WARN] `profile_off_band`: the mana available on turn four is 4.36, and bracket 4 wants 4.8 or more (mean over 10000 hands)

<details><summary>The deck list</summary>

- 24 Island | land | Reliable blue mana base.
- 1 Mishra's Workshop | land | Explosive artifact-only mana production.
- 1 Urza's Saga | land | Artifact utility land with Construct pressure.
- 1 Inventors' Fair | land | Artifact tutor land and incremental life.
- 1 Buried Ruin | land | Recovers an important destroyed artifact.
- 1 Gemstone Caverns | land | Occasional opening-hand acceleration.
- 1 Seat of the Synod | land | Blue source that also increases artifact density.
- 1 Spire of Industry | land | Flexible blue-producing artifact land support.
- 1 Glimmervoid | land | Consistent blue source in an artifact-heavy deck.
- 1 Otawara, Soaring City | land | Blue source with uncounterable utility interaction.
- 1 Mox Opal | ramp | Fast mana enabled by dense artifacts.
- 1 Moonsnare Prototype | ramp | Cheap artifact acceleration.
- 1 Vedalken Engineer | ramp | Efficient mana creature for artifact spells.
- 1 Chief Engineer | ramp | Convoke support accelerates artifact deployment.
- 1 Slagstone Refinery | ramp | Early artifact mana development.
- 1 Grand Architect | ramp | Turns blue creatures into powerful artifact mana sources.
- 1 Metalworker | ramp | High-output artifact mana engine.
- 1 Arc Reactor | ramp | Artifact-based mana acceleration.
- 1 Coin of Mastery | ramp | Mana production within the artifact shell.
- 1 Crystal Skull, Isu Spyglass | ramp | Mana rock with additional utility.
- 1 Network Terminal | ramp | Fixes mana and supports artifact development.
- 1 Inspiring Statuary | ramp | Artifacts help cast the deck's nonartifact spells.
- 1 Krark-Clan Ironworks | ramp | Converts expendable artifacts into burst mana.
- 1 Thoughtcast | draw | Highly efficient artifact-enabled card draw.
- 1 Riddlesmith | draw | Filters cards whenever artifacts are deployed.
- 1 Era of Innovation | draw | Artifact creature casting converts into cards.
- 1 Esoteric Duplicator | draw | Artifact value piece that supplies cards.
- 1 Forensic Gadgeteer | draw | Artifact-focused card advantage.
- 1 Sai, Master Thopterist | draw | Creates artifact bodies and turns them into cards.
- 1 Thirst for Knowledge | draw | Efficient instant-speed artifact filtering.
- 1 Cerebral Download | draw | Flexible card-advantage spell.
- 1 Katara, Bending Prodigy | draw | Provides steady card advantage.
- 1 Trading Post | draw | Versatile artifact value and card-draw outlet.
- 1 Vedalken Archmage | draw | Rewards repeated artifact casting with cards.
- 1 Reverse Engineer | draw | Artifact-enabled burst card draw.
- 1 Welding Jar | interaction | Protects key artifacts for no mana.
- 1 Metallic Rebuke | interaction | Efficient artifact-enabled counterspell.
- 1 Disruption Protocol | interaction | Artifact-assisted countermagic.
- 1 Ice Out | interaction | Flexible counterspell for crucial turns.
- 1 Override | interaction | Cheap protection and stack interaction.
- 1 Reality Ripple | interaction | Temporarily answers problematic permanents.
- 1 Ghostly Flicker | interaction | Protects permanents while reusing value effects.
- 1 Curator's Ward | interaction | Protects a premium permanent and replaces itself.
- 1 Escape Protocol | interaction | Reusable artifact protection and value.
- 1 Etched Champion | interaction | Resilient artifact threat that blocks effectively.
- 1 Padeem, Consul of Innovation | interaction | Protects the artifact core from opposing removal.
- 1 Aether Spellbomb | removal | Cheap, recurrable creature bounce.
- 1 Contagion Clasp | removal | Early creature removal with proliferate utility.
- 1 Cyber Conversion | removal | Efficient answer to a problematic permanent.
- 1 Resculpt | removal | Exiles an artifact or creature at instant speed.
- 1 Ravenform | removal | Exiles troublesome creatures or artifacts.
- 1 Unable to Scream | removal | Low-cost answer that blanks a creature.
- 1 Watery Grasp | removal | Versatile aura-based creature answer.
- 1 Zuko's Exile | removal | Efficient removal spell.
- 1 Into Thin Air | removal | Tempo-positive permanent interaction.
- 1 Kitesail Larcenist | removal | Turns opposing permanents into manageable artifacts.
- 1 Arcbound Ravager | synergy | Sacrifice outlet that converts artifacts into combat power.
- 1 Coretapper | synergy | Accelerates charge-counter artifact plans.
- 1 Energy Chamber | synergy | Builds counters across artifact permanents.
- 1 Drafna, Founder of Lat-Nam | synergy | Reuses artifacts and supplies artifact velocity.
- 1 Emry, Lurker of the Loch | synergy | Recasts artifacts from the graveyard.
- 1 Foundry Inspector | synergy | Reduces the cost of artifact spells.
- 1 Canoptek Scarab Swarm | threat | Artifact creature that grows into a meaningful attacker.
- 1 Filigree Attendant | threat | Evasive artifact threat that rewards broad development.
- 1 Frogmite | threat | Cheap affinity attacker that builds artifact count.
- 1 Chrome Steed | threat | Efficient metalcraft attacker.
- 1 Jhoira's Familiar | threat | Evasive artifact creature that also reduces costs.
- 1 Phyrexian Metamorph | threat | Copies the strongest artifact or creature in play.
- 1 Traxos, Scourge of Kroog | threat | Large, efficient artifact commander-style attacker.
- 1 Whirler Rogue | threat | Creates artifact bodies and enables evasive attacks.
- 1 Master Transmuter | threat | Cheats expensive artifacts into play at instant speed.
- 1 Arcbound Crusher | threat | Scales rapidly as artifacts enter the battlefield.
- 1 Panther Robot | threat | Efficient artifact creature for board pressure.
- 1 Kappa Cannoneer | threat | Powerful warded finisher in a dense artifact board.
- 1 Engineered Explosives | wipe | Flexible low-cost board reset.
- 1 Hurkyl's Recall | wipe | Sweeps opposing artifacts and creates a major tempo swing.

</details>

### 23. two set families at once

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 203 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $200.81 to buy, $200.81 the whole deck.

**Summary:** This is a white Hobbit-and-Bloomburrow creature deck built around Kíli the Resourceful, using Dwarves, small woodland creatures, artifacts, and Equipment to establish a board and keep cards flowing. It pressures the table through a mixture of creature and Equipment threats, while carrying broad removal, protective responses, and several reset buttons to recover contested games. The deck gives up some speed for its thematic mix of creatures and artifacts, and its larger plays ask the mana base and ramp package to do steady work before the strongest threats arrive.

- [INFO] `curve_summary`: average mana value 3.13 over 63 nonland cards

<details><summary>The deck list</summary>

- 22 Plains | land | Provides the deck's primary white mana base.
- 1 Castle Ardenvale | land | Adds a white land slot to the mana base.
- 1 Command Tower | land | Provides reliable commander-color mana.
- 1 Dragon-Cursed Halls | land | Adds a themed white mana source.
- 1 Elven Passage | land | Adds another white land to support the creature plan.
- 1 Evolving Wilds | land | Helps stabilize access to Plains.
- 1 Exotic Orchard | land | Provides flexible colored mana.
- 1 Fabled Passage | land | Helps stabilize access to Plains.
- 1 Hidden Grotto | land | Adds a flexible themed land slot.
- 1 Hobbit Hole | land | Adds a themed white mana source.
- 1 Lupinflower Village | land | Adds a Bloomburrow-themed land slot.
- 1 Minas Tirith | land | Adds a themed white mana source.
- 1 Path of Ancestry | land | Supports the deck's creature-focused theme.
- 1 Thriving Heath | land | Provides another white mana source.
- 1 Uncharted Haven | land | Provides flexible colored mana.
- 1 Arcane Signet | ramp | Efficiently accelerates the deck's mana development.
- 1 Bag End Banquet | ramp | Provides themed mana acceleration.
- 1 Beza, the Bounding Spring | ramp | Provides creature-based mana development.
- 1 Burnished Hart | ramp | Provides artifact-based mana development.
- 1 Fellwar Stone | ramp | Efficiently accelerates mana development.
- 1 Ghirapur Orrery | ramp | Supplies additional mana development.
- 1 Gilded Lotus | ramp | Provides substantial artifact mana acceleration.
- 1 Hedron Archive | ramp | Provides artifact mana acceleration.
- 1 Long-Bodied Grey Dog | ramp | Provides a Bloomburrow creature ramp piece.
- 1 Mind Stone | ramp | Provides efficient artifact mana acceleration.
- 1 Belladonna Took | draw | Provides Hobbit-themed card advantage.
- 1 Caretaker's Talent | draw | Provides ongoing card advantage for the creature plan.
- 1 Circuit Mender | draw | Provides artifact-based card advantage.
- 1 Cut a Deal | draw | Provides a direct card-advantage refill.
- 1 Dawn of a New Age | draw | Provides persistent card advantage.
- 1 Feather of Flight | draw | Provides low-cost card advantage.
- 1 Fountainport Bell | draw | Provides artifact-based card advantage.
- 1 Harvestrite Host | draw | Provides Bloomburrow-themed card advantage.
- 1 Heirloom Epic | draw | Provides a thematic card-advantage engine.
- 1 Idol of Oblivion | draw | Provides efficient artifact card advantage.
- 1 Inspiring Overseer | draw | Provides creature-based card advantage.
- 1 Baird, Steward of Argive | interaction | Adds a defensive interaction piece.
- 1 Bilbo's Gambit | interaction | Provides a flexible Hobbit-themed response.
- 1 Bofur, Reliable Guardian // Concerted Care | interaction | Provides a creature and interactive Adventure option.
- 1 Crumb and Get It | interaction | Provides a low-cost interactive response.
- 1 Dawn's Truce | interaction | Provides a protective interactive option.
- 1 Galadriel's Dismissal | interaction | Provides a versatile instant-speed response.
- 1 Luminous Broodmoth | interaction | Provides resilient creature-focused interaction.
- 1 Mithril Coat | interaction | Provides protective Equipment interaction.
- 1 Angel of the Ruins | removal | Provides a large creature with removal utility.
- 1 Banishing Light | removal | Provides broad permanent removal.
- 1 Bumbleflower's Sharepot | removal | Provides a themed artifact removal option.
- 1 Celebrate the Mountain-king | removal | Provides thematic enchantment-based removal.
- 1 Driftgloom Coyote | removal | Provides Bloomburrow creature-based removal.
- 1 Fiend Hunter | removal | Provides creature-based removal.
- 1 Fog on the Barrow-Downs | removal | Provides thematic Aura-based removal.
- 1 Generous Gift | removal | Provides flexible instant-speed removal.
- 1 Giant's Boulder | removal | Provides artifact-based removal.
- 1 Dusk // Dawn | wipe | Provides a flexible board-reset option.
- 1 Martial Coup | wipe | Provides a creature-focused board reset.
- 1 Promise of Loyalty | wipe | Provides a selective board-reset option.
- 1 Academy Manufactor | synergy | Supports the deck's artifact and token-themed pieces.
- 1 Blade Splicer | synergy | Supports the deck's artifact creature theme.
- 1 Carrot Cake | synergy | Provides a Bloomburrow-themed artifact synergy piece.
- 1 Dáin, Lord of the Iron Hills | synergy | Supports the deck's Dwarf-focused Hobbit package.
- 1 Iron Hills Blacksmith | synergy | Supports the deck's Dwarf and artifact package.
- 1 Maskwood Nexus | synergy | Helps unify the deck's mixed creature themes.
- 1 Ori, Keeper of Songs | synergy | Supports the deck's Dwarf-focused Hobbit package.
- 1 Patchwork Banner | synergy | Supports the deck's creature-type theme while adding mana.
- 1 Tangle Tumbler | synergy | Provides a Bloomburrow artifact synergy piece.
- 1 Three Tree Mascot | synergy | Supports the deck's mixed creature-type theme.
- 1 Andúril, Flame of the West | threat | Provides a legendary Equipment threat for the creature package.
- 1 Andúril, Narsil Reforged | threat | Provides a second thematic legendary Equipment threat.
- 1 Fíli the Pathfinder | threat | Provides a Hobbit-themed creature threat.
- 1 Helm of the Host | threat | Provides a high-impact legendary Equipment threat.
- 1 Jazal Goldmane | threat | Provides a creature threat for a wide board.
- 1 Karn, the Great Creator | threat | Provides a resilient artifact-centered threat.
- 1 Psychosis Crawler | threat | Provides an artifact creature threat alongside the draw package.
- 1 Serra Redeemer | threat | Provides a substantial creature threat.
- 1 Starforged Sword | threat | Provides an Equipment threat for the creature package.
- 1 Sting, Bilbo's Sword | threat | Provides a thematic legendary Equipment threat.
- 1 Sunscorch Regent | threat | Provides a large creature threat.
- 1 Sword of Vengeance | threat | Provides an Equipment threat for the creature package.

</details>

