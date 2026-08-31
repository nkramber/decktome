# PR-8 deck gate

Run date: 2026-08-31. Card snapshot: 2026-08-31.

Verdict: PASS. 18 of 18 decks passed every block check, 0 invented names reached the user, and 0 summaries stated a false rule of the game. All three bars are zero tolerance.

## Summary

| Measure | Value |
|---|---|
| Prompts | 18 |
| Decks returned | 18 |
| Decks with no block finding | 18 |
| Invented names that reached the user | 0 |
| Decks that needed the repair turn | 1 |
| Summaries judged (F-26) | 18 |
| Summaries that state a rule of the game | 3 |
| Summaries that state a FALSE rule | 0 |
| Judge errors | 0 |
| Errors | 0 |
| Prompt version | 10 |
| Calls | 37 |
| Cost | $1.0787 |
| Time | 852 seconds |

## Findings by code

A block stops the deck. A warning and a note are reports.

| Code | Count |
|---|---|
| `not_owned` | 26 |
| `curve_summary` | 18 |
| `bracket_prose_rules` | 12 |
| `basics_added` | 1 |
| `land_count` | 1 |

By severity: BLOCK 0. WARN 27. INFO 31. 

## Decks

### 1. lifegain Commander, any card

Format: Commander. Theme: lifegain. Pool: any_card. Shortlist: 303 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $1109.40 to buy, $1109.40 the whole deck.

**Summary:** This Karlov-led deck repeatedly leans on lifegain synergies to develop a threatening board, then wins by turning that advantage into sustained combat pressure from Angels, Vampires, and other large threats. It has a broad mix of removal, protection, and board clears to keep opposing plans contained while its card-advantage pieces keep the engine moving. The tradeoff is that the deck commits heavily to creatures and themed permanents, so it is most effective when it has time to establish its board and maintain momentum.

- JUDGE [true]: "This Karlov-led deck". Implies Karlov of the Ghost Council can serve as a Commander, which is true since he is a legendary creature eligible to be a commander.
- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 3.56 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Arguel's Blood Fast // Temple of Aclazotz | land | A black source that also occupies a land slot.
- 1 Blighted Steppe | land | A white land for the mana base.
- 1 Diamond Valley | land | A utility land in the mana base.
- 1 Glimmerpost | land | A utility land that supports the deck's life-focused theme.
- 1 High Market | land | A utility land for the mana base.
- 1 Kabira Crossroads | land | A white land that supports the life-focused theme.
- 1 Miren, the Moaning Well | land | A utility land for the mana base.
- 1 Nomad Stadium | land | A white land that supports the life-focused theme.
- 12 Plains | land | Core white mana for the deck.
- 1 Radiant Fountain | land | A utility land that supports the life-focused theme.
- 1 Scoured Barrens | land | A dual-color land that supports the life-focused theme.
- 1 Shambling Vent | land | A black-white land for the mana base.
- 12 Swamp | land | Core black mana for the deck.
- 1 Vault of the Archangel | land | A black-white utility land for the mana base.
- 1 Angel of Indemnity | ramp | A white ramp piece for advancing the board.
- 1 Battle Angels of Tyr | ramp | A ramp threat that helps develop mana.
- 1 Beza, the Bounding Spring | ramp | A white ramp piece for catching up on resources.
- 1 Carmen, Cruel Skymarcher | ramp | A black-white ramp option that fits the creature plan.
- 1 Crypt Ghast | ramp | A powerful black mana-development piece.
- 1 Hierophant's Chalice | ramp | A colorless ramp artifact for the mana base.
- 1 Life Insurance | ramp | A black ramp piece that fits the deck's theme.
- 1 Nuka-Cola Vending Machine | ramp | A colorless ramp artifact for developing resources.
- 1 Orazca Relic | ramp | A colorless ramp artifact that supports the deck's resources.
- 1 Pristine Talisman | ramp | A mana artifact that also fits the life-focused strategy.
- 1 Archivist of Oghma | draw | A white creature that supplies card advantage.
- 1 Ayara, First of Locthwain | draw | A black card-advantage engine on a creature.
- 1 Convalescent Care | draw | A life-focused source of card advantage.
- 1 Cosmos Elixir | draw | A colorless card-advantage artifact.
- 1 Dawn of Hope | draw | A life-focused card-advantage engine.
- 1 Enduring Innocence | draw | A white permanent that provides card advantage.
- 1 Mangara, the Diplomat | draw | A white creature that supplies card advantage.
- 1 Sigarda's Splendor | draw | A life-focused source of card advantage.
- 1 The Gaffer | draw | A white card-advantage creature for the life plan.
- 1 Well of Lost Dreams | draw | A life-focused artifact for converting the plan into cards.
- 1 Alseid of Life's Bounty | interaction | A low-cost protective interaction piece.
- 1 Courageous Resolve | interaction | A white instant for protecting the board plan.
- 1 Faith's Shield | interaction | A flexible white protection spell.
- 1 Metropolis Reformer | interaction | A white creature that contributes protective interaction.
- 1 Rune-Tail, Kitsune Ascendant // Rune-Tail's Essence | interaction | A life-focused protective permanent.
- 1 Werefox Bodyguard | interaction | A white creature that provides flexible interaction.
- 1 Ayli, Eternal Pilgrim | removal | A black-white creature that provides repeatable removal.
- 1 Cavalier of Night | removal | A black creature that answers opposing permanents.
- 1 Murderous Rider // Swift End | removal | A flexible black removal spell with a creature body.
- 1 Nightmare's Thirst | removal | A black removal spell that fits the life-focused theme.
- 1 Solitude | removal | A white creature that provides removal.
- 1 Sorin Markov | removal | A black planeswalker that supplies removal.
- 1 Umezawa's Jitte | removal | A colorless equipment piece that supplies removal.
- 1 Vona, Butcher of Magan | removal | A black-white creature that provides removal.
- 1 Ajani, Strength of the Pride | wipe | A white planeswalker that can reset the board.
- 1 Fumigate | wipe | A white board-clearing spell that fits the life plan.
- 1 Kaya's Wrath | wipe | A black-white board-clearing spell.
- 1 Ajani's Pridemate | synergy | A signature lifegain synergy creature.
- 1 Angel of Vitality | synergy | An Angel that supports repeated life gain.
- 1 Angelic Accord | synergy | A life-focused enchantment for building the board.
- 1 Blood Artist | synergy | A black creature that supports the deck's life-focused synergies.
- 1 Bloodthirsty Aerialist | synergy | A lifegain synergy creature that can become a major threat.
- 1 Cleric Class | synergy | A life-focused enchantment that reinforces the core plan.
- 1 Cleric of Life's Bond | synergy | A black-white Cleric for the lifegain theme.
- 1 Heliod, Sun-Crowned | synergy | A central white permanent for lifegain synergies.
- 1 Indulging Patrician | synergy | A black-white creature that rewards the life-focused plan.
- 1 Resplendent Angel | synergy | An Angel that rewards sustained life gain.
- 1 Righteous Valkyrie | synergy | An Angel Cleric that supports the life-focused creature plan.
- 1 Serra Ascendant | synergy | An efficient white creature for the life-focused strategy.
- 1 Vito, Thorn of the Dusk Rose | synergy | A black legendary creature that supports lifegain synergies.
- 1 Voice of the Blessed | synergy | A white lifegain synergy creature that develops into pressure.
- 1 Archangel of Thune | threat | A premier Angel threat for the life-focused creature plan.
- 1 Astarion, the Decadent | threat | A black-white legendary threat that fits the deck's theme.
- 1 Celestine, the Living Saint | threat | A white legendary threat for the life-focused strategy.
- 1 Cliffhaven Vampire | threat | A black creature that pressures opponents through the life plan.
- 1 Defiant Bloodlord | threat | A black Vampire threat for the life-focused strategy.
- 1 Divinity of Pride | threat | A black-white creature that serves as a substantial threat.
- 1 Gideon's Company | threat | A white creature threat that fits the lifegain theme.
- 1 Liesa, Forgotten Archangel | threat | A black-white Angel threat with a strong battlefield presence.
- 1 Lyra Dawnbringer | threat | A white Angel threat for closing games in combat.
- 1 Nykthos Paragon | threat | A white life-focused threat for turning the plan into pressure.
- 1 Rhox Faithmender | threat | A life-focused creature that is also a meaningful threat.
- 1 Tivash, Gloom Summoner | threat | A black legendary threat for the life-focused strategy.

</details>

### 2. aristocrats Commander, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 179 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $158.67 the whole deck.

**Summary:** This is a black-white aristocrats deck built around establishing a creature-based value engine, protecting the pieces that matter, and using sacrifice-focused synergies to keep resources flowing. It wins by turning a developed board into steady pressure while its larger creatures close games after opponents have been forced to spend resources. The deck gives up some speed for a durable, board-centered plan, so its strongest games come from assembling its engine and keeping it intact through the table's interaction.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 2.79 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Command Tower | land | Included in the mana base.
- 1 Path of Ancestry | land | Included in the mana base.
- 1 Marsh Flats | land | Included in the mana base.
- 1 Fabled Passage | land | Included in the mana base.
- 1 Evolving Wilds | land | Included in the mana base.
- 1 Terramorphic Expanse | land | Included in the mana base.
- 1 Castle Locthwain | land | Included in the mana base.
- 1 Takenuma, Abandoned Mire | land | Included in the mana base.
- 1 Minas Tirith | land | Included in the mana base.
- 1 Plaza of Heroes | land | Included in the mana base.
- 1 Secluded Courtyard | land | Included in the mana base.
- 1 War Room | land | Included in the mana base.
- 16 Plains | land | Provides a reliable basic land base.
- 8 Swamp | land | Provides a reliable basic land base.
- 1 Sol Ring | ramp | Provides an early mana boost.
- 1 Arcane Signet | ramp | Provides dependable mana acceleration.
- 1 Commander's Sphere | ramp | Fills a mana acceleration slot.
- 1 Fellwar Stone | ramp | Fills a mana acceleration slot.
- 1 Thought Vessel | ramp | Fills a mana acceleration slot.
- 1 Wayfarer's Bauble | ramp | Fills a mana acceleration slot.
- 1 Relic of Legends | ramp | Fills a mana acceleration slot.
- 1 Springleaf Drum | ramp | Fills a mana acceleration slot.
- 1 Giada, Font of Hope | ramp | Adds to the deck's mana development.
- 1 Lotho, Corrupt Shirriff | ramp | Adds to the deck's mana development.
- 1 Skullclamp | draw | A focused card-advantage piece for the creature plan.
- 1 Idol of Oblivion | draw | Adds ongoing card advantage.
- 1 Nasty End | draw | Provides a card-advantage spell.
- 1 Night's Whisper | draw | Provides efficient card advantage.
- 1 Call of the Ring | draw | Adds a card-advantage engine.
- 1 Lembas | draw | Provides card advantage within the artifact package.
- 1 Mask of Memory | draw | Adds a card-advantage option.
- 1 Tome of Legends | draw | Adds a persistent card-advantage piece.
- 1 Wall of Omens | draw | Provides card advantage while occupying an early board slot.
- 1 Puresteel Paladin | draw | Adds card advantage alongside the equipment package.
- 1 Bitter Triumph | removal | A flexible answer for opposing problems.
- 1 Claim the Precious | removal | A dedicated answer for opposing threats.
- 1 Crib Swap | removal | A dedicated answer for opposing creatures.
- 1 Fatal Push | removal | An efficient spot-answer option.
- 1 Generous Gift | removal | A broad answer for opposing permanents.
- 1 Infernal Grasp | removal | A dedicated creature answer.
- 1 Swords to Plowshares | removal | An efficient creature answer.
- 1 Stroke of Midnight | removal | A broad answer for opposing permanents.
- 1 Clever Concealment | interaction | Protects the established board.
- 1 Darksteel Plate | interaction | Supports protection for a key creature.
- 1 Lightning Greaves | interaction | Protects a key creature.
- 1 Swiftfoot Boots | interaction | Protects a key creature.
- 1 Unbreakable Formation | interaction | Provides a defensive response for the board.
- 1 Champion's Helm | interaction | Supports protection for the commander.
- 1 Austere Command | wipe | Provides a reset option when the board gets away from you.
- 1 Fumigate | wipe | Provides a full-board reset option.
- 1 Vanquish the Horde | wipe | Provides another reset for creature-heavy boards.
- 1 Arcade Cabinet | synergy | Supports the deck's sacrifice-focused core.
- 1 Phantom Train | synergy | Supports the deck's sacrifice-focused core.
- 1 Gollum the Abandoned | synergy | A thematic piece for the aristocrats plan.
- 1 Gollum, Patient Plotter | synergy | A thematic piece for the aristocrats plan.
- 1 Gríma Wormtongue | synergy | Supports the deck's creature-focused value plan.
- 1 Heirloom Auntie | synergy | Supports the deck's creature-focused value plan.
- 1 Joo Dee, One of Many | synergy | Supports the deck's creature-focused value plan.
- 1 Nimble Hobbit | synergy | Supports the deck's creature-focused value plan.
- 1 Deadly Dispute | synergy | Supports the sacrifice-focused game plan.
- 1 Gift of Immortality | synergy | Helps preserve an important creature in the core plan.
- 1 Together Forever | synergy | Supports the deck's creature-focused value plan.
- 1 Fiend Hunter | synergy | Contributes to the creature-based aristocrats shell.
- 1 Blowfly Infestation | synergy | Adds another engine-oriented piece to the core plan.
- 1 Sheltered by Ghosts | synergy | Supports the creature-focused core plan.
- 1 Angel of Serenity | threat | A substantial creature to pressure opponents.
- 1 Archfiend of Ifnir | threat | A substantial creature for the top end.
- 1 Bill the Pony | threat | A creature that contributes to board pressure.
- 1 Exemplar of Light | threat | A creature that contributes to board pressure.
- 1 Inspiring Overseer | threat | A creature that contributes to board pressure.
- 1 June, Bounty Hunter | threat | A creature that contributes to board pressure.
- 1 Massacre Girl, Known Killer | threat | A meaningful creature threat with board presence.
- 1 Palace Jailer | threat | A creature that contributes to board pressure.
- 1 Rat King, Pale Piper | threat | A thematic creature threat.
- 1 The Sackville-Bagginses | threat | A creature that contributes to board pressure.
- 1 Vengeful Villagers | threat | A creature that contributes to board pressure.
- 1 Witch-king of Angmar | threat | A substantial creature for the top end.

</details>

### 3. artifacts Commander, bracket 4

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 299 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $4753.68 to buy, $4753.68 the whole deck.

**Summary:** Urza leads a dense artifact engine that develops mana quickly, keeps cards flowing, and uses tutors and utility artifacts to assemble the right board for the moment. The deck wins by turning its artifact advantage into sustained pressure from large artifact threats while maintaining answers and protection for key pieces. It gives up some flexibility in favor of committing heavily to artifacts, so opposing disruption aimed at that core can slow its momentum.

- [INFO] `curve_summary`: average mana value 4.06 over 63 nonland cards

<details><summary>The deck list</summary>

- 18 Island | land | Provides the deck's primary blue mana base.
- 1 Academy Ruins | land | An artifact-focused utility land for the mana base.
- 1 Buried Ruin | land | An artifact-focused utility land for the mana base.
- 1 Inventors' Fair | land | An artifact-focused utility land for the mana base.
- 1 Urza's Saga | land | A thematic artifact land that contributes to the mana base.
- 1 Mishra's Workshop | land | A high-impact artifact land for powering out the deck's plan.
- 1 Power Depot | land | An artifact land that supports the deck's mana base.
- 1 Glimmervoid | land | A flexible land for an artifact-heavy mana base.
- 1 Spire of Industry | land | A flexible land for an artifact-heavy mana base.
- 1 Blinkmoth Nexus | land | A utility land that fits the artifact theme.
- 1 Inkmoth Nexus | land | A utility land that fits the artifact theme.
- 1 Mirrex | land | A utility land that supports the artifact strategy.
- 1 Phyrexia's Core | land | A utility land for the artifact-focused mana base.
- 1 Treasure Map // Treasure Cove | land | A land-slot card that supports the deck's artifact theme.
- 1 Primal Amulet // Primal Wellspring | land | A land-slot card that supports the deck's spell package.
- 1 Conqueror's Galleon // Conqueror's Foothold | land | A land-slot artifact card that fits the deck's plan.
- 1 Matzalantli, the Great Door // The Core | land | A utility land-slot artifact for the mana base.
- 1 Roadside Reliquary | land | A utility land for the artifact-heavy mana base.
- 1 Otawara, Soaring City | land | A blue utility land that supports the mana base.
- 1 Mox Opal | ramp | Provides efficient artifact-based acceleration.
- 1 Metalworker | ramp | A dedicated artifact mana accelerator.
- 1 Krark-Clan Ironworks | ramp | An artifact-based source of explosive acceleration.
- 1 Moonsnare Prototype | ramp | Efficient early artifact acceleration.
- 1 Grand Architect | ramp | Supports mana development in an artifact-heavy deck.
- 1 Chief Engineer | ramp | Helps accelerate the deck's artifact deployment.
- 1 Tezzeret the Seeker | ramp | A planeswalker included for artifact-focused acceleration.
- 1 Karn, Legacy Reforged | ramp | A major artifact-focused mana engine.
- 1 Inspiring Statuary | ramp | Turns the artifact board into additional acceleration.
- 1 The Mightstone and Weakstone | ramp | An artifact accelerator that advances the main plan.
- 1 Thoughtcast | draw | Efficient card flow for an artifact-heavy deck.
- 1 Thought Monitor | draw | An artifact threat that also supports card flow.
- 1 Sai, Master Thopterist | draw | An artifact-focused source of ongoing card flow.
- 1 Thopter Spy Network | draw | A persistent artifact-themed draw engine.
- 1 Vedalken Archmage | draw | Rewards the deck's dense artifact deployment with cards.
- 1 Riddlesmith | draw | Provides card selection alongside artifact casting.
- 1 Reverse Engineer | draw | A dedicated draw spell for the artifact strategy.
- 1 One with the Machine | draw | A dedicated draw spell that fits the artifact plan.
- 1 Tezzeret, Artifice Master | draw | A planeswalker that supplies recurring card flow.
- 1 Tezzeret, Betrayer of Flesh | draw | A planeswalker that supplies artifact-focused card flow.
- 1 Metallic Rebuke | interaction | Efficient stack interaction for an artifact-heavy board.
- 1 Disruption Protocol | interaction | Artifact-friendly interaction for protecting the plan.
- 1 Ice Out | interaction | Flexible instant-speed interaction.
- 1 Stoic Rebuttal | interaction | Reliable interaction in an artifact-focused shell.
- 1 Padeem, Consul of Innovation | interaction | Helps shield the deck's key artifact plan.
- 1 Welding Jar | interaction | Low-cost protection for important artifacts.
- 1 Aetherflux Reservoir | removal | A powerful artifact-based answer and payoff.
- 1 Arcum Dagsson | removal | An artifact-focused removal piece.
- 1 Portal to Phyrexia | removal | A large artifact answer that also pressures opponents.
- 1 Spine of Ish Sah | removal | A versatile artifact answer for difficult permanents.
- 1 Skysovereign, Consul Flagship | removal | A threatening artifact that provides removal.
- 1 Resculpt | removal | Flexible spot removal at instant speed.
- 1 Ravenform | removal | A dedicated answer for problematic targets.
- 1 Transmogrifying Wand | removal | Repeatable artifact-based spot removal.
- 1 Nevinyrral's Disk | wipe | A broad reset button when the board gets away.
- 1 Contagion Engine | wipe | An artifact board-control option.
- 1 Coercive Portal | wipe | A flexible artifact-based board reset.
- 1 Emry, Lurker of the Loch | synergy | A centerpiece for recurring and leveraging artifacts.
- 1 Etherium Sculptor | synergy | Makes the artifact-heavy deck operate more efficiently.
- 1 Foundry Inspector | synergy | Improves the efficiency of artifact deployment.
- 1 Mystic Forge | synergy | Provides an artifact-centric engine for continued development.
- 1 Unwinding Clock | synergy | Supports an artifact board through repeated untapping.
- 1 Voltaic Key | synergy | A cheap artifact utility piece for the deck's engines.
- 1 Manifold Key | synergy | A low-cost artifact utility piece for the deck's engines.
- 1 Clock of Omens | synergy | Converts a broad artifact board into engine value.
- 1 Power Artifact | synergy | An enchantment engine that supports the artifact plan.
- 1 Transmute Artifact | synergy | Finds the artifact piece best suited to the current board.
- 1 Whir of Invention | synergy | An instant-speed artifact tutor for key engine pieces.
- 1 Scrap Trawler | synergy | Supports artifact recursion and long-game resilience.
- 1 Shimmer Myr | synergy | Adds flexibility to artifact deployment.
- 1 The Antiquities War | synergy | A thematic artifact payoff that advances the board.
- 1 Kappa Cannoneer | threat | A premier artifact threat that rewards developing the board.
- 1 Kuldotha Forgemaster | threat | A high-impact artifact threat for the deck's core plan.
- 1 Cyberdrive Awakener | threat | A major artifact-based finisher threat.
- 1 Metalwork Colossus | threat | A large artifact threat for closing games.
- 1 Mycosynth Golem | threat | A powerful artifact creature that rewards the deck's density.
- 1 Traxos, Scourge of Kroog | threat | An efficient legendary artifact threat.
- 1 Karn, Scion of Urza | threat | A planeswalker threat that fits the artifact strategy.
- 1 Master Transmuter | threat | An artifact-focused creature threat with high impact.
- 1 Broodstar | threat | A large threat that rewards a full artifact board.
- 1 Darksteel Juggernaut | threat | A scaling artifact threat for the late game.
- 1 Lodestone Golem | threat | A disruptive artifact creature that applies pressure.
- 1 Threefold Thunderhulk | threat | A substantial artifact creature threat for closing games.

</details>

### 4. dinosaur tribal, bracket 2

Format: Commander. Theme: dinosaurs. Pool: any_card. Shortlist: 269 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $215.68 to buy, $215.68 the whole deck.

**Summary:** This deck develops its mana, builds a board of Dinosaur creatures, and uses its commander alongside large threats to pressure opponents through combat. Its draw cards help keep creatures coming, while its removal and wipes give it ways to recover when opposing boards get ahead. The deck gives up some speed and fine-tuned answers in exchange for a clear, creature-focused plan with many large plays.

- [INFO] `bracket_prose_rules`: bracket 2 (Core, 8+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 4.25 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Command Tower | land | Land for the deck’s three-color mana base.
- 1 Path of Ancestry | land | Land that supports the Dinosaur-focused deck.
- 1 Evolving Wilds | land | Land for finding a needed basic land.
- 1 Terramorphic Expanse | land | Land for finding a needed basic land.
- 1 Myriad Landscape | land | Land slot with additional mana development value.
- 1 Canopy Vista | land | Green-white land for the mana base.
- 1 Cinder Glade | land | Red-green land for the mana base.
- 1 Clifftop Retreat | land | Red-white land for the mana base.
- 1 Rootbound Crag | land | Red-green land for the mana base.
- 1 Sacred Foundry | land | Red-white land for the mana base.
- 1 Stomping Ground | land | Red-green land for the mana base.
- 1 Sunpetal Grove | land | Green-white land for the mana base.
- 1 Temple Garden | land | Green-white land for the mana base.
- 1 Battlefield Forge | land | Red-white land for the mana base.
- 1 Exotic Orchard | land | Flexible land for the mana base.
- 10 Forest | land | Basic green land for a green-heavy Dinosaur deck.
- 6 Mountain | land | Basic red land for the mana base.
- 5 Plains | land | Basic white land for the mana base.
- 1 Sol Ring | ramp | Efficient artifact ramp to advance the deck’s mana.
- 1 Arcane Signet | ramp | Artifact ramp that supports all deck colors.
- 1 Cultivate | ramp | Sorcery ramp for steady land development.
- 1 Farseek | ramp | Sorcery ramp that helps fix mana.
- 1 Nature's Lore | ramp | Sorcery ramp for early mana development.
- 1 Kodama's Reach | ramp | Sorcery ramp that builds the mana base.
- 1 Thunderherd Migration | ramp | Dinosaur-themed sorcery ramp.
- 1 Drover of the Mighty | ramp | Creature ramp that fits the deck’s creature plan.
- 1 Ranging Raptors | ramp | Dinosaur ramp that supports the tribal plan.
- 1 Hulking Raptor | ramp | Dinosaur ramp that also adds to the creature board.
- 1 Beast Whisperer | draw | Creature-based draw for a deck built around creatures.
- 1 Garruk's Uprising | draw | Enchantment draw that supports large creatures.
- 1 Guardian Project | draw | Ongoing draw for the creature-heavy plan.
- 1 Harmonize | draw | Straightforward sorcery draw for a new player.
- 1 Return of the Wildspeaker | draw | Flexible instant-speed draw for a large-creature deck.
- 1 Ripjaw Raptor | draw | Dinosaur draw that contributes to the tribal board.
- 1 Rishkar's Expertise | draw | High-impact sorcery draw for a deck with large threats.
- 1 Shamanic Revelation | draw | Sorcery draw that rewards building a creature board.
- 1 Vanquisher's Banner | draw | Tribal artifact draw for the Dinosaur plan.
- 1 Vaultborn Tyrant | draw | Dinosaur draw attached to a major creature.
- 1 Akroma's Will | interaction | Instant interaction for protecting or pushing the creature board.
- 1 Boros Charm | interaction | Versatile instant interaction in the deck’s colors.
- 1 Heroic Intervention | interaction | Protective instant interaction for the permanent board.
- 1 Lightning Greaves | interaction | Equipment interaction that protects an important creature.
- 1 Swiftfoot Boots | interaction | Equipment interaction that protects an important creature.
- 1 Temple Altisaur | interaction | Dinosaur interaction that supports the creature board.
- 1 Apex Altisaur | removal | Dinosaur removal that doubles as a sizable creature.
- 1 Bronzebeak Foragers | removal | Dinosaur removal that fits the tribal game plan.
- 1 Burning Sun's Avatar | removal | Dinosaur removal with a meaningful board presence.
- 1 Itzquinth, Firstborn of Gishath | removal | Low-cost Dinosaur removal for the early game.
- 1 Ravenous Sailback | removal | Dinosaur removal that remains useful as a creature.
- 1 Thrashing Brontodon | removal | Dinosaur removal that is easy to deploy.
- 1 Trumpeting Carnosaur | removal | Dinosaur removal that supports the top end.
- 1 Zacama, Primal Calamity | removal | Large Dinosaur removal for established boards.
- 1 Austere Command | wipe | Flexible board wipe for difficult board states.
- 1 Blasphemous Act | wipe | Efficient wipe when the board is crowded.
- 1 Wakening Sun's Avatar | wipe | Dinosaur-themed wipe that fits the deck’s plan.
- 1 Commune with Dinosaurs | synergy | Dinosaur synergy that helps support the tribal plan.
- 1 Dinosaur Stampede | synergy | Dinosaur synergy for turning the creature board into pressure.
- 1 Huatli's Raptor | synergy | Low-cost Dinosaur synergy creature.
- 1 Kinjalli's Caller | synergy | Tribal support for deploying Dinosaur creatures.
- 1 Kinjalli's Sunwing | synergy | Dinosaur synergy that supports combat-focused play.
- 1 Marauding Raptor | synergy | Dinosaur synergy creature for the tribal shell.
- 1 Nest Robber | synergy | Early Dinosaur synergy creature.
- 1 Otepec Huntmaster | synergy | Tribal support for Dinosaur creatures.
- 1 Raptor Companion | synergy | Simple Dinosaur synergy creature for the early board.
- 1 Regal Imperiosaur | synergy | Dinosaur synergy creature that adds to the tribal board.
- 1 Relentless Raptor | synergy | Efficient Dinosaur synergy creature.
- 1 Territorial Hammerskull | synergy | Dinosaur synergy creature for combat pressure.
- 1 Triceraton Commander | synergy | Dinosaur synergy creature for the tribal plan.
- 1 Tyrox, Saurid Tyrant | synergy | Legendary Dinosaur synergy creature for the deck theme.
- 1 Carnage Tyrant | threat | Large Dinosaur threat that pressures opponents.
- 1 Etali, Primal Storm | threat | Legendary Dinosaur threat at the top of the curve.
- 1 Ghalta and Mavren | threat | Legendary Dinosaur threat that advances the combat plan.
- 1 Ghalta, Primal Hunger | threat | Large legendary Dinosaur threat.
- 1 Ghalta, Stampede Tyrant | threat | Large legendary Dinosaur threat for finishing games.
- 1 Goring Ceratops | threat | Dinosaur threat that strengthens combat turns.
- 1 Pantlaza, Sun-Favored | threat | Legendary Dinosaur threat that supports the creature plan.
- 1 Quartzwood Crasher | threat | Dinosaur threat for a combat-focused strategy.
- 1 Regisaur Alpha | threat | Dinosaur threat that contributes to the tribal board.
- 1 Tyrranax Rex | threat | Large Dinosaur threat for closing games.
- 1 Verdant Sun's Avatar | threat | Dinosaur threat that supports staying power.
- 1 Zetalpa, Primal Dawn | threat | Large legendary Dinosaur threat at the top end.

</details>

### 5. blink Commander, owned first

Format: Commander. Theme: blink. Pool: owned_first. Shortlist: 168 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for copy_limit. Block findings: 0.

Cost: $0.00 to buy, $214.45 the whole deck.

**Summary:** A mono-white blink deck built around repeatedly reusing creature entry effects while protecting its legendary centerpiece. Evasive Angels, disruptive creatures, and artifact-backed value provide pressure, with broad answers and reset effects keeping the board manageable.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 2.83 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Abandoned Air Temple | land | Flexible white source.
- 1 Adventurer's Inn | land | Theme-friendly mana source.
- 1 Ash Barrens | land | Basic-land fixing.
- 1 City of Brass | land | Unrestricted color fixing.
- 1 Command Tower | land | Reliable commander mana.
- 1 Crossroads Village | land | Theme-friendly mana source.
- 1 Eclipsed Realms | land | Utility mana source.
- 1 Escape Tunnel | land | Fixing and utility.
- 1 Evolving Wilds | land | Basic-land fixing.
- 1 Exotic Orchard | land | Flexible mana source.
- 1 Fabled Passage | land | Efficient basic-land fixing.
- 1 Field of Ruin | land | Answers troublesome lands.
- 1 Fountainport | land | Utility mana source.
- 1 Ghost Quarter | land | Land disruption.
- 1 Great Hall of the Citadel | land | Legendary-theme mana source.
- 1 Marsh Flats | land | Finds Plains.
- 1 Minas Tirith | land | Legendary-theme utility land.
- 1 Opal Palace | land | Supports the commander.
- 1 Path of Ancestry | land | Creature-theme fixing.
- 1 Plains | land | Reliable white mana.
- 1 Plaza of Heroes | land | Protects legendary creatures.
- 1 Reliquary Tower | land | Hand-size utility.
- 1 Rogue's Passage | land | Provides evasion.
- 1 Scavenger Grounds | land | Graveyard control.
- 1 Secluded Courtyard | land | Creature-type fixing.
- 1 Shire Terrace | land | Basic-land fixing.
- 1 Spire of Industry | land | Artifact-supported fixing.
- 1 Study Hall | land | Commander utility.
- 1 Temple of the False God | land | High-output mana source.
- 1 Terramorphic Expanse | land | Basic-land fixing.
- 1 The Gold Saucer | land | Theme-friendly utility land.
- 1 The Grey Havens | land | Legendary-theme mana source.
- 1 Unclaimed Territory | land | Creature-type fixing.
- 1 Vibrant Cityscape | land | Flexible mana source.
- 1 War Room | land | Repeatable card advantage.
- 1 Windbrisk Heights | land | Value-oriented utility land.
- 1 Arcane Signet | ramp | Efficient mana acceleration.
- 1 Astral Cornucopia | ramp | Scalable mana production.
- 1 Commander's Sphere | ramp | Mana source with later utility.
- 1 Fellwar Stone | ramp | Efficient mana acceleration.
- 1 Sol Ring | ramp | Fast colorless acceleration.
- 1 Springleaf Drum | ramp | Turns creatures into mana.
- 1 Sword of the Animist | ramp | Combat-based land ramp.
- 1 Thought Vessel | ramp | Mana acceleration and hand utility.
- 1 Wayfarer's Bauble | ramp | Finds a basic Plains.
- 1 White Lotus Tile | ramp | Mana acceleration.
- 1 Buster Sword | draw | Equipment-based card advantage.
- 1 Crown of Gondor | draw | Legendary equipment card advantage.
- 1 Diary of Dreams | draw | Repeatable card advantage.
- 1 Instant Ramen | draw | Low-cost card advantage.
- 1 Mask of Memory | draw | Combat-based card selection.
- 1 Mirror of Galadriel | draw | Repeatable card filtering.
- 1 Skullclamp | draw | Efficient creature-based card advantage.
- 1 Stone of Erech | draw | Card advantage with utility.
- 1 Tome of Legends | draw | Commander-supported card advantage.
- 1 Wizard's Rockets | draw | Artifact-based card advantage.
- 1 Banishing Light | removal | Versatile permanent exile.
- 1 Crib Swap | removal | Exiles a creature at instant speed.
- 1 Dispatch | removal | Efficient creature exile.
- 1 Generous Gift | removal | Answers any permanent.
- 1 Get Lost | removal | Efficient broad removal.
- 1 Journey to Nowhere | removal | Efficient creature exile.
- 1 Oblation | removal | Flexible permanent answer.
- 1 Swords to Plowshares | removal | Premium creature exile.
- 1 Champion's Helm | interaction | Protects a legendary commander.
- 1 Clever Concealment | interaction | Protects the board from removal.
- 1 Darksteel Plate | interaction | Durable creature protection.
- 1 Lightning Greaves | interaction | Immediate protection and haste.
- 1 Reprieve | interaction | Temporizes opposing spells.
- 1 Swiftfoot Boots | interaction | Reliable commander protection.
- 1 Angel of Condemnation | synergy | Repeatable blink outlet and creature control.
- 1 Angel of Sanctions | synergy | Reusable exile trigger.
- 1 Champions of Minas Tirith | synergy | Creature-based value for repeated triggers.
- 1 Exemplar of Light | synergy | Angel value piece that benefits from creature recursion.
- 1 Faramir, Field Commander | synergy | Legendary creature value engine.
- 1 Fiend Hunter | synergy | Blinkable creature exile effect.
- 1 Flickerwisp | synergy | Additional blink effect with an evasive body.
- 1 Gift of Immortality | synergy | Keeps key creatures returning.
- 1 Inspiring Overseer | synergy | Blinkable card-advantage trigger.
- 1 Palace Jailer | synergy | Blinkable removal and card advantage.
- 1 Sheltered by Ghosts | synergy | Protects a key creature while adding value.
- 1 Slip On the Ring | synergy | Instant-speed blink protection.
- 1 Together Forever | synergy | Recovers important creatures.
- 1 Wall of Omens | synergy | Low-cost blinkable card-advantage trigger.
- 1 Adventurer's Airship | threat | Evasive artifact threat.
- 1 Angel of Serenity | threat | Large evasive finisher with graveyard utility.
- 1 Bastion Protector | threat | Protective legendary-creature threat.
- 1 Boromir, Warden of the Tower | threat | Legendary disruptive attacker.
- 1 Bronze Guardian | threat | Large artifact threat with protection.
- 1 Frontline Medic | threat | Combat-focused creature threat.
- 1 Giada, Font of Hope | threat | Evasive legendary Angel threat.
- 1 Grim Poppet | threat | Artifact creature that pressures opposing boards.
- 1 Kataki, War's Wage | threat | Artifact-taxing disruptive creature.
- 1 PuPu UFO | threat | Artifact creature threat.
- 1 Puresteel Paladin | threat | Equipment-supported creature threat.
- 1 Westfold Rider | threat | Creature threat with useful removal pressure.
- 1 Austere Command | wipe | Flexible board reset.
- 1 Fumigate | wipe | Creature reset with life recovery.
- 1 Vanquish the Horde | wipe | Efficient creature board reset.

</details>

### 6. Modern tempo, tournament

Format: Modern. Theme: tempo. Pool: any_card. Shortlist: 167 names.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $463.92 to buy, $463.92 the whole deck.

**Summary:** This blue-red tempo deck establishes pressure with cheap creature threats, then uses draw spells, removal, and interaction to keep the opponent from stabilizing. Temporal Mastery provides a focused synergy element alongside the spell-heavy core. It gives up broad permanent types and sweeping effects in favor of a streamlined plan built around early threats and efficient support.

- [INFO] `curve_summary`: average mana value 2.00 over 36 nonland cards

<details><summary>The deck list</summary>

- 8 Island | land | Primary blue source for the deck's tempo spells.
- 4 Mountain | land | Red source for the removal suite.
- 4 Steam Vents | land | Blue-red land that supports both halves of the deck.
- 4 Scalding Tarn | land | Land slot supporting the blue-red mana base.
- 4 Shivan Reef | land | Blue-red land for consistent early mana.
- 4 Consider | draw | Cheap draw spell that keeps the deck moving.
- 2 Preordain | draw | Additional early draw and selection.
- 4 Counterspell | interaction | Core interaction for protecting the tempo plan.
- 2 Spell Pierce | interaction | Low-cost interaction that fits an early-pressure strategy.
- 4 Lightning Bolt | removal | Efficient removal for clearing the way.
- 4 Into the Flood Maw | removal | Additional removal to maintain tempo.
- 4 Temporal Mastery | synergy | Temporal card that supplies the deck's synergy package.
- 4 Ledger Shredder | threat | Early creature threat that supports the spell-heavy plan.
- 4 Ragavan, Nimble Pilferer | threat | Early creature threat for applying pressure.
- 4 Faerie Mastermind | threat | Evasive creature threat that complements the tempo shell.

</details>

### 7. Modern burn, casual

Format: Modern. Theme: burn. Pool: any_card. Shortlist: 290 names.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $28.62 to buy, $28.62 the whole deck.

**Summary:** This mono-red burn deck presses opponents with early creatures and damage-focused synergy pieces, then uses efficient removal and direct burn to finish the game. Its threats make the spell-heavy plan less dependent on drawing only burn, while its draw cards help maintain pressure. The deck gives up broad answers and heavier late-game options in exchange for a focused, aggressive game plan.

- [INFO] `curve_summary`: average mana value 2.83 over 36 nonland cards

<details><summary>The deck list</summary>

- 20 Mountain | land | Provides the stable mono-red land base.
- 4 Ramunap Ruins | land | Fills land slots while supporting the red-focused mana base.
- 2 Chandra, Dressed to Kill | ramp | Supplies the requested ramp element without diluting the burn plan.
- 2 Ancestral Anger | draw | Low-cost draw that keeps spells moving through the deck.
- 4 Risk Factor | draw | Draw option that fits a direct-damage-focused strategy.
- 4 Lightning Bolt | removal | Efficient removal for clearing opposing pressure or advancing the burn plan.
- 2 Lava Dart | removal | Additional inexpensive removal that works well with spell-focused synergies.
- 4 Eidolon of the Great Revel | synergy | A synergy piece that reinforces the deck's aggressive burn focus.
- 4 Thermo-Alchemist | synergy | A synergy creature for a deck built around casting burn spells.
- 4 Akki Lavarunner // Tok-Tok, Volcano Born | threat | A low-end threat that adds creature pressure early.
- 4 Keral Keep Disciples | threat | A threat that supports the deck's proactive board presence.
- 4 Onakke Javelineer | threat | An additional threat for maintaining early pressure.
- 2 Hazoret the Fervent | threat | A durable top-end threat for closing games after early burn pressure.

</details>

### 8. Modern lifegain, FNM

Format: Modern. Theme: lifegain. Pool: any_card. Shortlist: 301 names.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $160.13 to buy, $160.13 the whole deck.

**Summary:** This white-black lifegain deck establishes its synergy pieces early, uses them to support a steady stream of creature threats, and clears away opposing pressure with focused removal. It wins by turning its lifegain plan into sustained board pressure, with larger threats providing a stronger finish. The tradeoff is a concentrated game plan with limited ramp and fewer answers outside its removal suite.

- [INFO] `curve_summary`: average mana value 2.78 over 36 nonland cards

<details><summary>The deck list</summary>

- 8 Plains | land | Provides a reliable white mana base for the deck.
- 7 Swamp | land | Provides a reliable black mana base for the deck.
- 4 Scoured Barrens | land | Supports both deck colors while fitting the lifegain plan.
- 3 Shambling Vent | land | Provides white-black mana and functions as a land slot.
- 2 Restless Fortress | land | Provides white-black mana while remaining within the land base.
- 4 Dawn of Hope | draw | Supplies card draw for a deck built around lifegain synergies.
- 2 Well of Lost Dreams | draw | Adds further card draw tied to the deck's central plan.
- 2 Arcane Signet | ramp | Accelerates access to the deck's white and black spells.
- 4 Path to Exile | removal | Provides efficient removal for opposing threats.
- 2 Murderous Rider // Swift End | removal | Adds flexible removal in a card that can also contribute to the board.
- 4 Soul Warden | synergy | A core lifegain synergy piece that supports the rest of the deck.
- 4 Ajani's Pridemate | synergy | Rewards the deck for repeatedly pursuing its lifegain plan.
- 4 Attended Healer | threat | Provides a lifegain-focused threat for the board.
- 4 Bloodbond Vampire | threat | Serves as a threat that fits the deck's lifegain theme.
- 2 Archangel of Thune | threat | Provides a powerful top-end threat for the lifegain strategy.
- 4 Twinblade Paladin | threat | Adds another lifegain-oriented threat to pressure opponents.

</details>

### 9. Standard midrange, FNM

Format: Standard. Theme: midrange. Pool: any_card. Shortlist: 164 names.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $218.94 to buy, $218.94 the whole deck.

**Summary:** This black-green midrange deck develops its mana, trades resources with flexible removal, and keeps cards coming while deploying a steady stream of substantial creature threats. It wins by maintaining pressure through the midgame and forcing opponents to answer successive threats, with support cards helping that creature-centered plan hold together. The tradeoff is a focused main deck with little room for specialized answers outside its core creature, removal, and card-advantage approach.

- [INFO] `curve_summary`: average mana value 3.06 over 34 nonland cards

<details><summary>The deck list</summary>

- 8 Forest | land | Basic green land for the black-green mana base.
- 8 Swamp | land | Basic black land for the black-green mana base.
- 4 Blooming Marsh | land | Dual land supporting both deck colors.
- 2 Deathcap Glade | land | Dual land supporting the midrange mana base.
- 2 Underground Mortuary | land | Swamp Forest land that supports both colors.
- 2 Llanowar Elves | ramp | Early ramp that helps the deck reach its larger cards.
- 2 Phyrexian Arena | draw | Dedicated draw source for keeping resources flowing.
- 2 Unholy Annex // Ritual Chamber | draw | Draw card that supports the deck's longer game.
- 2 Corrupted Conviction | draw | Low-cost draw option for refilling resources.
- 2 Bitter Triumph | removal | Efficient removal for problematic opposing cards.
- 2 Assassin's Trophy | removal | Flexible removal for a broad range of threats.
- 2 Maelstrom Pulse | removal | Versatile removal that suits a midrange plan.
- 2 Goldvein Hydra | threat | Creature threat that adds pressure in the midgame.
- 2 Darkstar Augur | threat | Creature threat that contributes to the deck's board presence.
- 2 Massacre Girl, Known Killer | threat | Legendary creature threat for attrition-focused games.
- 2 Vein Ripper | threat | Large creature threat for closing games.
- 2 Vaultborn Tyrant | threat | Top-end creature threat for the late game.
- 2 Aclazotz, Deepest Betrayal // Temple of the Dead | threat | Creature-land threat that remains useful in longer games.
- 2 Rottenmouth Viper | threat | Creature threat that adds to the deck's pressure plan.
- 2 Snakeskin Veil | synergy | Supporting interaction that helps the creature-focused plan.
- 2 Royal Treatment | synergy | Supporting interaction for the deck's key threats.
- 2 Undying Malice | synergy | Supporting interaction that reinforces the creature plan.
- 2 Innkeeper's Talent | synergy | Support card that fits the deck's threat-focused strategy.

</details>

### 10. Standard aggro, tournament

Format: Standard. Theme: aggro. Pool: any_card. Shortlist: 294 names.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $71.42 to buy, $71.42 the whole deck.

**Summary:** This red-white aggro deck aims to establish threats quickly, support them with Warleader's Call, and keep attacking while removal clears resistance. Draw cards help maintain pressure after the first wave, while interaction protects the tempo of the assault. It gives up larger standalone finishers and broad reset effects in favor of a focused, proactive plan.

- [INFO] `curve_summary`: average mana value 2.72 over 36 nonland cards

<details><summary>The deck list</summary>

- 12 Mountain | land | Provides red mana for the deck’s aggressive plan.
- 8 Plains | land | Provides white mana for the deck’s aggressive plan.
- 4 Sacred Foundry | land | Provides red-white mana in the land base.
- 4 Redcap Gutter-Dweller | threat | Supplies a threat for early board pressure.
- 4 Dragonback Lancer | threat | Supplies a threat that advances the attacking plan.
- 4 Teapot Slinger | threat | Supplies another threat for sustained pressure.
- 4 Fugitive Codebreaker | draw | Provides draw support while fitting the aggressive strategy.
- 2 Reckless Lackey | draw | Adds further draw support to keep the deck moving.
- 4 Boros Charm | interaction | Provides flexible interaction for protecting the attack plan.
- 2 Dawn's Truce | interaction | Provides additional interaction when the opponent disrupts the board.
- 4 Harsh Annotation | removal | Provides efficient removal to clear opposing resistance.
- 4 Case of the Gateway Express | removal | Adds removal that supports forcing attacks through.
- 4 Warleader's Call | synergy | Provides synergy for the red-white creature assault.

</details>

### 11. Commander with a locked card

Format: Commander. Theme: sacrifice. Pool: any_card. Shortlist: 296 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $1126.47 to buy, $1126.47 the whole deck.

**Summary:** This Karlov of the Ghost Council deck centers on a creature-heavy sacrifice engine, using ramp and draw to keep deploying creatures while sacrifice synergies turn them into steady pressure. Its threats provide the closing power, with targeted removal and board wipes clearing resistance when the table gets ahead. The deck gives up some flexibility in exchange for committing heavily to creatures, sacrifice outlets, and death-focused payoffs.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 3.17 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Baron Bertram Graywater | draw | Draw support that fits the deck's black-white creature base.
- 1 Bushmeat Poacher | draw | Draw support for the sacrifice-focused plan.
- 1 Corrupted Conviction | draw | Efficient draw that supports sacrificing resources.
- 1 Disciple of Bolas | draw | Creature-based draw for a deck built around expendable bodies.
- 1 Ecstatic Awakener // Awoken Demon | draw | Draw option that also contributes a creature to the sacrifice shell.
- 1 Relic Vial | draw | Repeatable draw support in the artifact slot.
- 1 Shadowheart, Dark Justiciar | draw | Draw support on a creature that fits the deck's plan.
- 1 Smothering Abomination | draw | Dedicated draw threat for repeated sacrifice turns.
- 1 Vampiric Rites | draw | Low-cost draw support for sacrificing creatures.
- 1 Village Rites | draw | Efficient draw that turns a sacrificed creature into cards.
- 1 Cartel Aristocrat | interaction | Interactive creature that supports the sacrifice plan.
- 1 Dark Privilege | interaction | Interaction that protects a key permanent while using the sacrifice theme.
- 1 Fanatical Devotion | interaction | Protective interaction that rewards having creatures available.
- 1 Gift of Doom | interaction | Protective interaction for an important threat or engine piece.
- 1 Rescue from the Underworld | interaction | Interaction that helps preserve a creature-focused board.
- 1 Sunstone | interaction | Artifact-based interaction for stabilizing the table.
- 1 Command Tower | land | Reliable color fixing for the commander deck.
- 1 Bojuka Bog | land | Black mana source with utility.
- 1 High Market | land | Utility land that supports the sacrifice theme.
- 1 Phyrexian Tower | land | Utility land that supports creature sacrifices.
- 1 Westvale Abbey // Ormendahl, Profane Prince | land | Land slot that also contributes a late-game threat option.
- 1 Diamond Valley | land | Utility land for the deck's creature-centric plan.
- 1 Miren, the Moaning Well | land | Utility land suited to sacrificing creatures.
- 1 Reliquary Tower | land | Colorless utility land for retaining drawn cards.
- 1 Rogue's Passage | land | Utility land that helps a threat press through stalled boards.
- 1 Myriad Landscape | land | Land slot that supports mana development.
- 1 Evolving Wilds | land | Fixing land for the two-color mana base.
- 1 Terramorphic Expanse | land | Additional fixing land for the mana base.
- 1 Scavenger Grounds | land | Utility land with a useful disruptive role.
- 1 Ifnir Deadlands | land | Black utility land.
- 1 Lazotep Quarry | land | Black utility land.
- 1 Bloodstained Mire | land | Fetch land that improves mana access.
- 10 Plains | land | Basic white mana sources for the deck's white cards.
- 10 Swamp | land | Basic black mana sources for the deck's black-heavy sacrifice plan.
- 1 Sol Ring | ramp | Fast artifact mana required for the deck.
- 1 Ashnod's Altar | ramp | Sacrifice outlet that converts creatures into mana.
- 1 Phyrexian Altar | ramp | Sacrifice outlet and mana engine for creature loops and deployment.
- 1 Pitiless Plunderer | ramp | Creature-based ramp that rewards the sacrifice plan.
- 1 Pawn of Ulamog | ramp | Turns creature losses into additional mana resources.
- 1 Priest of Forgotten Gods | ramp | Sacrifice-focused creature that supplies mana support.
- 1 Sifter of Skulls | ramp | Ramp payoff for a deck expecting creatures to die.
- 1 Skullport Merchant | ramp | Creature-based mana support for the sacrifice shell.
- 1 Culling the Weak | ramp | Burst mana that turns a creature into explosive development.
- 1 Deadly Dispute | ramp | Sacrifice-based resource spell that contributes mana development.
- 1 Attrition | removal | Sacrifice-based removal engine.
- 1 Ayli, Eternal Pilgrim | removal | Creature-based removal that fits the deck's colors and theme.
- 1 Bone Shards | removal | Low-cost removal that can use a disposable creature.
- 1 Dictate of Erebos | removal | Removal payoff for repeated creature sacrifices.
- 1 Grave Pact | removal | Powerful removal payoff for the sacrifice engine.
- 1 Eaten Alive | removal | Sacrifice-compatible creature removal.
- 1 Stronghold Assassin | removal | Repeatable creature-based removal for the board.
- 1 Yawgmoth, Thran Physician | removal | Versatile removal creature that belongs in the sacrifice shell.
- 1 Elas il-Kor, Sadistic Pilgrim | synergy | Core black-white sacrifice synergy piece.
- 1 Bastion of Remembrance | synergy | Enchantment payoff for the deck's sacrifice plan.
- 1 Zulaport Cutthroat | synergy | Creature payoff for repeated sacrifices.
- 1 Viscera Seer | synergy | Low-cost sacrifice outlet that smooths the deck's game plan.
- 1 Carrion Feeder | synergy | Reliable sacrifice outlet on an early creature.
- 1 Woe Strider | synergy | Sacrifice outlet that also contributes a creature body.
- 1 Bloodflow Connoisseur | synergy | Creature-based sacrifice outlet and payoff.
- 1 Nantuko Husk | synergy | Sacrifice outlet that can become a meaningful attacker.
- 1 Fleshtaker | synergy | Black-white sacrifice synergy creature.
- 1 Hidden Stockpile | synergy | Enchantment engine for the sacrifice-oriented board.
- 1 Open the Graves | synergy | Creature-focused payoff that sustains the sacrifice plan.
- 1 Chthonian Nightmare | synergy | Sacrifice synergy piece for recurring creature-focused value.
- 1 Victimize | synergy | Sacrifice-based recursion that rebuilds the board.
- 1 Witch's Oven | synergy | Cheap artifact sacrifice outlet.
- 1 Razaketh, the Foulblooded | threat | High-impact Demon threat for the top end.
- 1 Liesa, Forgotten Archangel | threat | Powerful black-white creature threat.
- 1 Mondrak, Glory Dominus | threat | High-impact white creature threat for a token-friendly shell.
- 1 Felisa, Fang of Silverquill | threat | Black-white threat that fits a creature-focused strategy.
- 1 Ghoulcaller Gisa | threat | Creature threat that supports a sacrifice-oriented board.
- 1 Requiem Angel | threat | Flying white threat for the deck's creature plan.
- 1 Ratadrabik of Urborg | threat | Legendary black-white threat that supports the creature theme.
- 1 Sidisi, Undead Vizier | threat | Threat that contributes to the sacrifice-focused creature package.
- 1 Abhorrent Overlord | threat | Large Demon threat for closing games.
- 1 Liliana, Heretical Healer // Liliana, Defiant Necromancer | threat | Creature-based threat that fits a deck where creatures frequently die.
- 1 Vindictive Vampire | threat | Vampire threat that reinforces the sacrifice theme.
- 1 Sadistic Hypnotist | threat | High-impact creature threat for a sacrifice-heavy board.
- 1 Toxic Deluge | wipe | Efficient board wipe for resetting difficult boards.
- 1 The Meathook Massacre | wipe | Board wipe that fits the deck's creature-death theme.
- 1 Liliana, Dreadhorde General | wipe | Planeswalker-based board wipe for the top end.

</details>

### 12. Commander on a budget

Format: Commander. Theme: tokens. Pool: any_card. Shortlist: 296 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $43.22 to buy, $43.22 the whole deck.

**Summary:** Adeline leads a white token strategy that aims to establish a broad board, keep resources flowing, and pressure opponents through sustained combat backed by token-focused synergies and creature threats. The deck can clear away problematic permanents or reset crowded boards before rebuilding its presence. It gives up premium individual card power and expensive mana acceleration in favor of a cohesive, budget-conscious board-building plan.

- JUDGE [true]: "Adeline leads a white token strategy". Adeline, Resplendent Cathar is a legendary creature and is eligible to serve as a Commander, so the claim that she leads the deck is a true statement about commander eligibility.
- [INFO] `bracket_prose_rules`: bracket 2 (Core, 8+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 3.41 over 63 nonland cards

<details><summary>The deck list</summary>

- 36 Plains | land | Provides the basic white land base for the deck.
- 1 Bucknard's Everfull Purse | ramp | Included as an inexpensive ramp piece.
- 1 Charisma Bobblehead | ramp | Included as an inexpensive ramp piece.
- 1 Coin of Mastery | ramp | Included as an inexpensive ramp piece.
- 1 Collector's Vault | ramp | Included as an inexpensive ramp piece.
- 1 Currency Converter | ramp | Included as an inexpensive ramp piece.
- 1 Druidic Satchel | ramp | Included as an inexpensive ramp piece.
- 1 Goldvein Pick | ramp | Included as an inexpensive ramp piece.
- 1 Karn, Living Legacy | ramp | Adds a ramp option that also fits the deck's permanent-heavy plan.
- 1 Keeper of the Accord | ramp | Provides a white ramp option for the deck.
- 1 Monologue Tax | ramp | Adds efficient ramp while staying within the budget plan.
- 1 Angelic Sell-Sword | draw | Provides a low-cost draw option.
- 1 Bygone Bishop | draw | Adds draw support in a creature slot.
- 1 Dawn of Hope | draw | Provides repeatable draw support for the token plan.
- 1 Idol of Oblivion | draw | Adds artifact-based draw support for a token deck.
- 1 Platoon Dispenser | draw | Provides draw support while fitting the deck's board-focused plan.
- 1 Staff of the Storyteller | draw | Adds a low-cost draw option.
- 1 Tamiyo's Journal | draw | Provides a durable draw piece.
- 1 Wedding Announcement // Wedding Festivity | draw | Supports the deck's token plan while filling a draw slot.
- 1 Wojek Investigator | draw | Adds draw support in a creature slot.
- 1 Sunpearl Kirin | draw | Provides another inexpensive draw option.
- 1 Paladin's Arms | interaction | Provides a low-cost interaction piece.
- 1 Pegasus Guardian // Rescue the Foal | interaction | Adds flexible interaction while fitting the deck's creature plan.
- 1 Rootborn Defenses | interaction | Provides interaction that suits a wide-board strategy.
- 1 Spirit Bonds | interaction | Adds interaction that supports the deck's token theme.
- 1 Squad Commander | interaction | Provides interaction in a creature slot.
- 1 Teyo, the Shieldmage | interaction | Adds a budget interaction option.
- 1 Aerial Assault | removal | Provides efficient targeted removal.
- 1 Banishing Slash | removal | Adds a low-cost removal spell.
- 1 Citizen's Crowbar | removal | Provides artifact-based removal support.
- 1 Generous Gift | removal | Adds flexible targeted removal.
- 1 Kellan's Lightblades | removal | Provides inexpensive targeted removal.
- 1 Righteous Confluence | removal | Adds versatile removal coverage.
- 1 Skyclave Apparition | removal | Provides removal attached to a creature.
- 1 Stroke of Midnight | removal | Adds flexible instant-speed removal.
- 1 Hour of Reckoning | wipe | Provides a board-reset option suited to the deck's token focus.
- 1 Martial Coup | wipe | Provides a wipe that remains aligned with the token strategy.
- 1 Phyrexian Rebirth | wipe | Adds a budget board-reset option.
- 1 Clarion Spirit | synergy | Supports the deck's token-focused synergies.
- 1 Divine Visitation | synergy | Provides a high-impact token synergy piece.
- 1 Felidar Retreat | synergy | Supports the deck's token-focused game plan.
- 1 Intangible Virtue | synergy | Strengthens the deck's token-focused plan.
- 1 Mavren Fein, Dusk Apostle | synergy | Adds a creature-based token synergy.
- 1 Retrofitter Foundry | synergy | Provides a repeatable artifact-based token synergy.
- 1 Rosie Cotton of South Lane | synergy | Supports a board built around many creatures and tokens.
- 1 Siege Veteran | synergy | Adds a Soldier-oriented token synergy.
- 1 Sigil of the Empty Throne | synergy | Provides a token payoff within the enchantment package.
- 1 Skrelv's Hive | synergy | Adds a focused token synergy piece.
- 1 Three Blind Mice | synergy | Provides an inexpensive token-focused synergy card.
- 1 Twilight Drover | synergy | Supports the deck's token strategy in a creature slot.
- 1 Visions of Glory | synergy | Adds a low-cost token synergy spell.
- 1 Worthy Knight | synergy | Provides a creature-based token synergy.
- 1 Ajani's Chosen | threat | Serves as a token-oriented threat.
- 1 Archon of Sun's Grace | threat | Provides a token-focused threat.
- 1 Attended Healer | threat | Adds a low-cost creature threat.
- 1 Basri's Lieutenant | threat | Provides a creature threat that fits the deck's board plan.
- 1 Cemetery Protector | threat | Adds a resilient creature threat.
- 1 Defiler of Faith | threat | Provides a white creature threat.
- 1 Dragonback Lancer | threat | Adds an inexpensive creature threat.
- 1 Drogskol Cavalry | threat | Provides a larger token-oriented threat.
- 1 Emeria Angel | threat | Adds a creature threat aligned with the token plan.
- 1 Gideon, Ally of Zendikar | threat | Provides a board-focused planeswalker threat.
- 1 God-Eternal Oketra | threat | Adds a powerful creature threat for the deck's permanent plan.
- 1 Hero of Bladehold | threat | Provides a combat-focused creature threat.

</details>

### 13. owned first, and the commander is not owned

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 229 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $132.72 the whole deck.

**Summary:** This deck centers Karlov of the Ghost Council in a lifegain-focused creature strategy, developing value through healers, Angels, equipment, and supporting permanents before pressuring the table with its threat suite. It wins by maintaining a resilient board and converting that development into combat pressure, while removal, protection, and sweepers keep opposing boards from taking over. The tradeoff is a board-focused plan that is less explosive when its creatures and synergy pieces cannot stay in play.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [WARN] `not_owned`: Karlov of the Ghost Council: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 2.98 over 63 nonland cards

<details><summary>The deck list</summary>

- 18 Plains | land | Provides dependable white mana for the deck’s core spells.
- 10 Swamp | land | Provides dependable black mana for the deck’s core spells.
- 1 Command Tower | land | Provides flexible mana for Karlov’s colors.
- 1 Marsh Flats | land | Finds a needed basic land early.
- 1 Fabled Passage | land | Finds a needed basic land while improving mana access.
- 1 Evolving Wilds | land | Finds either basic land color.
- 1 Ash Barrens | land | Helps convert into the basic land color needed.
- 1 Takenuma, Abandoned Mire | land | Supplies black mana from a utility land slot.
- 1 War Room | land | Adds a utility land to the mana base.
- 1 Scavenger Grounds | land | Adds utility without taking a spell slot.
- 1 Sol Ring | ramp | Provides efficient early mana acceleration.
- 1 Arcane Signet | ramp | Fixes and accelerates the deck’s mana.
- 1 Commander's Sphere | ramp | Provides flexible mana acceleration.
- 1 Fellwar Stone | ramp | Adds another inexpensive mana source.
- 1 Thought Vessel | ramp | Accelerates mana development.
- 1 Wayfarer's Bauble | ramp | Turns an artifact slot into lasting mana development.
- 1 Sword of the Animist | ramp | Pairs mana development with the deck’s creature plan.
- 1 Relic of Legends | ramp | Provides additional mana acceleration.
- 1 Springleaf Drum | ramp | Lets the creature-heavy deck produce extra mana.
- 1 Deadly Dispute | ramp | Converts a spare permanent into cards and mana.
- 1 Night's Whisper | draw | Efficiently replenishes cards.
- 1 Idol of Oblivion | draw | Provides a reusable source of cards.
- 1 Mask of Memory | draw | Rewards attacking creatures with card flow.
- 1 Skullclamp | draw | Turns expendable creatures into cards.
- 1 Tome of Legends | draw | Supports continued card flow alongside Karlov.
- 1 Wall of Omens | draw | Provides an early body while replacing itself.
- 1 Lembas | draw | Adds card flow in an artifact slot.
- 1 Painful Truths | draw | Provides a strong burst of cards.
- 1 Puresteel Paladin | draw | Supports the equipment package with card flow.
- 1 Exemplar of Light | draw | Adds a creature-based draw option.
- 1 Lightning Greaves | interaction | Protects Karlov or a key threat.
- 1 Swiftfoot Boots | interaction | Provides repeatable protection for important creatures.
- 1 Darksteel Plate | interaction | Helps keep a central creature on the battlefield.
- 1 Clever Concealment | interaction | Protects the developed board from opposing disruption.
- 1 Unbreakable Formation | interaction | Defends the creature board at a critical moment.
- 1 Take Up the Shield | interaction | Protects a key creature during combat or removal.
- 1 Swords to Plowshares | removal | Provides efficient creature removal.
- 1 Fatal Push | removal | Offers low-cost targeted removal.
- 1 Generous Gift | removal | Answers a wide range of problematic permanents.
- 1 Get Lost | removal | Provides flexible targeted removal.
- 1 Infernal Grasp | removal | Answers opposing creatures cleanly.
- 1 Bitter Triumph | removal | Provides flexible instant-speed removal.
- 1 Dispatch | removal | Adds another efficient removal spell.
- 1 Stroke of Midnight | removal | Answers a troublesome permanent at instant speed.
- 1 Aerith Gainsborough | synergy | Supports the deck’s lifegain-focused synergy plan.
- 1 Angel of Vitality | synergy | Builds on repeated lifegain.
- 1 Compassionate Healer | synergy | Adds another lifegain-oriented creature.
- 1 Elixir | synergy | Provides a compact piece for the deck’s synergy package.
- 1 Kor Firewalker | synergy | Contributes to the lifegain theme on a creature body.
- 1 Light of Promise | synergy | Rewards the deck for pursuing lifegain.
- 1 Night Nurse, Healer of Heroes | synergy | Fits the healer and lifegain game plan.
- 1 Prideful Feastling | synergy | Supports the deck’s food and lifegain-adjacent plan.
- 1 Rosie Cotton of South Lane | synergy | Builds value from the deck’s token and lifegain synergies.
- 1 Second Breakfast | synergy | Adds another focused lifegain synergy piece.
- 1 White Mage's Staff | synergy | Supports the deck’s healer-themed synergy package.
- 1 Well-Worn Spatula | synergy | Adds an equipment-based synergy piece.
- 1 Eastfarthing Farmer | synergy | Contributes to the deck’s value-oriented synergy plan.
- 1 Aettir and Priwen | synergy | Adds another equipment synergy piece for the creature plan.
- 1 Angel of Invention | threat | Provides a substantial evasive threat.
- 1 Canyon Crawler | threat | Adds a durable creature threat.
- 1 Dawnhand Eulogist | threat | Provides another creature that can pressure opponents.
- 1 Gollum, Silent Slinker // Meager Meal | threat | Adds a black threat that fits the deck’s creature plan.
- 1 Lyra Dawnbringer | threat | Supplies a powerful Angel threat.
- 1 Minwu, White Mage | threat | Adds a focused white creature threat.
- 1 Shattered Angel | threat | Provides an additional Angel threat.
- 1 Sneering Shadewriter | threat | Adds a black creature threat to diversify the board.
- 1 Victory's Herald | threat | Provides a high-impact creature threat.
- 1 Lo and Li, Twin Tutors | threat | Adds a legendary creature threat.
- 1 Rooftop Percher | threat | Provides another creature for applying pressure.
- 1 Rabaroo Troop | threat | Rounds out the deck’s creature threat suite.
- 1 Austere Command | wipe | Provides a flexible reset when the board gets out of hand.
- 1 Fumigate | wipe | Resets opposing creature boards while fitting the deck’s plan.
- 1 Vanquish the Horde | wipe | Provides an efficient creature-board reset.

</details>

### 14. the user delegates the commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 244 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $838.68 to buy, $838.68 the whole deck.

**Summary:** Atarka, World Render leads a Gruul Dragon deck that builds its mana, develops a board of Dragon threats, and supports that plan with Dragon-focused cards. It aims to win by turning its large threats into sustained combat pressure while using removal, wipes, and protective interaction to keep opponents from stabilizing. The deck gives up some speed and flexibility to devote most of its resources to the Dragon plan.

- JUDGE [true]: "Atarka, World Render leads a Gruul Dragon deck". Atarka, World Render is a legendary creature and thus can legally serve as a commander, and its colors are Red/Green (Gruul).
- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 4.14 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Command Tower | land | Included as a land in the mana base.
- 1 Stomping Ground | land | Included as a land in the mana base.
- 1 Cinder Glade | land | Included as a land in the mana base.
- 1 Rootbound Crag | land | Included as a land in the mana base.
- 1 Path of Ancestry | land | Included as a land in the mana base.
- 1 Temple of the Dragon Queen | land | Included as a land in the mana base.
- 1 Crucible of the Spirit Dragon | land | Included as a land in the mana base.
- 1 Haven of the Spirit Dragon | land | Included as a land in the mana base.
- 1 Maelstrom of the Spirit Dragon | land | Included as a land in the mana base.
- 1 Three Tree City | land | Included as a land in the mana base.
- 1 Cavern of Souls | land | Included as a land in the mana base.
- 1 Boseiju, Who Endures | land | Included as a land in the mana base.
- 1 Yavimaya, Cradle of Growth | land | Included as a land in the mana base.
- 1 Ancient Tomb | land | Included as a land in the mana base.
- 1 Nykthos, Shrine to Nyx | land | Included as a land in the mana base.
- 1 Myriad Landscape | land | Included as a land in the mana base.
- 1 Evolving Wilds | land | Included as a land in the mana base.
- 1 Fabled Passage | land | Included as a land in the mana base.
- 1 Terramorphic Expanse | land | Included as a land in the mana base.
- 1 Rogue's Passage | land | Included as a land in the mana base.
- 1 War Room | land | Included as a land in the mana base.
- 1 Reliquary Tower | land | Included as a land in the mana base.
- 8 Forest | land | Basic lands that anchor the mana base.
- 6 Mountain | land | Basic lands that anchor the mana base.
- 1 Carnelian Orb of Dragonkind | ramp | Included to ramp into the deck's Dragon plan.
- 1 Dragon's Hoard | ramp | Included to ramp into the deck's Dragon plan.
- 1 Jade Orb of Dragonkind | ramp | Included to ramp into the deck's Dragon plan.
- 1 Orb of Dragonkind | ramp | Included to ramp into the deck's Dragon plan.
- 1 Scaled Nurturer | ramp | Included to ramp into the deck's Dragon plan.
- 1 Ganax, Astral Hunter | ramp | Included to ramp into the deck's Dragon plan.
- 1 Goldspan Dragon | ramp | Included to ramp into the deck's Dragon plan.
- 1 Klauth, Unrivaled Ancient | ramp | Included to ramp into the deck's Dragon plan.
- 1 Old Gnawbone | ramp | Included to ramp into the deck's Dragon plan.
- 1 Ancient Copper Dragon | ramp | Included to ramp into the deck's Dragon plan.
- 1 Beast Whisperer | draw | Included as a draw piece for the creature-heavy plan.
- 1 Dragonborn Champion | draw | Included as a draw piece for the Dragon plan.
- 1 Elemental Bond | draw | Included to provide card draw.
- 1 Garruk's Uprising | draw | Included to provide card draw.
- 1 Guardian Project | draw | Included to provide card draw.
- 1 Return of the Wildspeaker | draw | Included to provide card draw.
- 1 Rishkar's Expertise | draw | Included to provide card draw.
- 1 Shamanic Revelation | draw | Included to provide card draw.
- 1 Sylvan Library | draw | Included to provide card draw.
- 1 Vanquisher's Banner | draw | Included to provide card draw.
- 1 Heroic Intervention | interaction | Included as interaction for key permanents.
- 1 Lightning Greaves | interaction | Included as interaction for key creatures.
- 1 Swiftfoot Boots | interaction | Included as interaction for key creatures.
- 1 Tamiyo's Safekeeping | interaction | Included as interaction for important permanents.
- 1 Tibalt's Trickery | interaction | Included as flexible interaction.
- 1 Brotherhood Regalia | interaction | Included as interaction for the Dragon plan.
- 1 Dragon Tempest | removal | Included as removal that supports the Dragon plan.
- 1 Draconic Roar | removal | Included as removal.
- 1 Dragon's Fire | removal | Included as removal.
- 1 Dragonlord Atarka | removal | Included as a Dragon removal piece.
- 1 Drakuseth, Maw of Flames | removal | Included as a Dragon removal piece.
- 1 Scourge of Valkas | removal | Included as a Dragon removal piece.
- 1 Terror of the Peaks | removal | Included as a Dragon removal piece.
- 1 Wrathful Red Dragon | removal | Included as a Dragon removal piece.
- 1 Acolyte of Bahamut | synergy | Included to support the Dragon-focused strategy.
- 1 Breaching Dragonstorm | synergy | Included to support the Dragon-focused strategy.
- 1 Crucible of Fire | synergy | Included to support the Dragon-focused strategy.
- 1 Dracogenesis | synergy | Included to support the Dragon-focused strategy.
- 1 Dragon Egg | synergy | Included to support the Dragon-focused strategy.
- 1 Dragon Hatchling | synergy | Included to support the Dragon-focused strategy.
- 1 Dragonkin Berserker | synergy | Included to support the Dragon-focused strategy.
- 1 Dragonlord's Servant | synergy | Included to support the Dragon-focused strategy.
- 1 Dragonspeaker Shaman | synergy | Included to support the Dragon-focused strategy.
- 1 Firespitter Whelp | synergy | Included to support the Dragon-focused strategy.
- 1 Kargan Dragonrider | synergy | Included to support the Dragon-focused strategy.
- 1 Minion of the Mighty | synergy | Included to support the Dragon-focused strategy.
- 1 Sarkhan's Triumph | synergy | Included to support the Dragon-focused strategy.
- 1 The Dragon-Kami Reborn // Dragon-Kami's Egg | synergy | Included to support the Dragon-focused strategy.
- 1 Ancient Bronze Dragon | threat | Included as a high-impact Dragon threat.
- 1 Blast-Furnace Hellkite | threat | Included as a Dragon threat.
- 1 Caldera Pyremaw | threat | Included as a Dragon threat.
- 1 Hellkite Charger | threat | Included as a Dragon threat.
- 1 Lathliss, Dragon Queen | threat | Included as a Dragon threat.
- 1 Scourge of the Throne | threat | Included as a Dragon threat.
- 1 Terror of Mount Velus | threat | Included as a Dragon threat.
- 1 Thrakkus the Butcher | threat | Included as a Dragon threat.
- 1 Thunderbreak Regent | threat | Included as a Dragon threat.
- 1 Twinflame Tyrant | threat | Included as a Dragon threat.
- 1 Utvara Hellkite | threat | Included as a Dragon threat.
- 1 Worldgorger Dragon | threat | Included as a Dragon threat.
- 1 Balefire Dragon | wipe | Included as a Dragon-based wipe.
- 1 Breath Weapon | wipe | Included as a wipe.
- 1 Draconic Intervention | wipe | Included as a wipe.

</details>

### 15. delegated commander, owned first

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 229 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $178.81 the whole deck.

**Summary:** This deck builds a lifegain-focused board around Astarion, then turns that steady resource advantage into pressure through Angels, equipped creatures, and other sizeable threats. It has a broad set of answers and ways to protect important creatures while keeping cards and mana flowing. The tradeoff is that the deck is built for a measured, board-centric game rather than the fastest possible finish.

- [INFO] `bracket_prose_rules`: bracket 2 (Core, 8+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [WARN] `not_owned`: Astarion, the Decadent: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.14 over 63 nonland cards

<details><summary>The deck list</summary>

- 18 Plains | land | Provides reliable white mana for the lifegain-focused core.
- 12 Swamp | land | Provides reliable black mana for the deck’s support spells.
- 1 Ash Barrens | land | Flexible land slot for smoothing the mana base.
- 1 Command Tower | land | Reliable commander-focused color fixing.
- 1 Evolving Wilds | land | Fetch land that helps stabilize the mana base.
- 1 Fabled Passage | land | Fetch land that helps stabilize the mana base.
- 1 Marsh Flats | land | Fetch land for the deck’s two colors.
- 1 Terramorphic Expanse | land | Fetch land that helps stabilize the mana base.
- 1 Sol Ring | ramp | Efficient artifact acceleration.
- 1 Arcane Signet | ramp | Reliable color fixing and acceleration.
- 1 Bender's Waterskin | ramp | Artifact-based mana acceleration.
- 1 Commander's Sphere | ramp | Mana fixing that remains useful later.
- 1 Fellwar Stone | ramp | Low-cost artifact acceleration.
- 1 Giada, Font of Hope | ramp | Creature-based acceleration that fits the Angel contingent.
- 1 Lotho, Corrupt Shirriff | ramp | Creature-based mana acceleration.
- 1 Relic of Legends | ramp | Mana acceleration for a legendary-heavy deck.
- 1 Thought Vessel | ramp | Artifact acceleration with lasting utility.
- 1 Wayfarer's Bauble | ramp | Early mana development from an artifact.
- 1 Buster Sword | draw | Equipment-based card advantage.
- 1 Call of the Ring | draw | Ongoing card-advantage source.
- 1 Exemplar of Light | draw | Creature-based card advantage that fits the deck’s theme.
- 1 Grave Venerations | draw | Card-advantage enchantment.
- 1 Idol of Oblivion | draw | Low-cost artifact card advantage.
- 1 Inspiring Overseer | draw | Theme-fitting creature card advantage.
- 1 Lembas | draw | Flexible artifact card advantage.
- 1 Mask of Memory | draw | Equipment-based card selection and advantage.
- 1 Night's Whisper | draw | Efficient black card advantage.
- 1 Skullclamp | draw | Efficient equipment-based card advantage.
- 1 Bastion Protector | interaction | Protects the commander from opposing disruption.
- 1 Champion's Helm | interaction | Equipment protection for the commander or a key threat.
- 1 Clever Concealment | interaction | Protective response for the board.
- 1 Darksteel Plate | interaction | Durable equipment protection for an important creature.
- 1 Lightning Greaves | interaction | Efficient protection for the commander or a threat.
- 1 Swiftfoot Boots | interaction | Flexible equipment protection for key creatures.
- 1 Banishing Light | removal | Versatile permanent-based answer.
- 1 Bitter Triumph | removal | Flexible instant-speed answer.
- 1 Crib Swap | removal | Creature answer that fits the deck’s colors.
- 1 Generous Gift | removal | Broad answer to troublesome permanents.
- 1 Get Lost | removal | Efficient white removal.
- 1 Infernal Grasp | removal | Efficient black creature removal.
- 1 Swords to Plowshares | removal | Premium creature removal.
- 1 Stroke of Midnight | removal | Flexible white permanent removal.
- 1 Aerith Gainsborough | synergy | Lifegain-focused synergy creature.
- 1 Angel of Vitality | synergy | Angel that supports the lifegain plan.
- 1 Aettir and Priwen | synergy | Equipment slot supporting the deck’s synergies.
- 1 Adventurous Eater // Have a Bite | synergy | Theme-supporting lifegain synergy card.
- 1 Compassionate Healer | synergy | Core lifegain synergy creature.
- 1 Dancer's Chakrams | synergy | Equipment-based synergy piece.
- 1 Elixir | synergy | Artifact synergy piece for the lifegain shell.
- 1 Kor Firewalker | synergy | Creature that supports the lifegain strategy.
- 1 Light of Promise | synergy | Lifegain-focused Aura payoff.
- 1 Night Nurse, Healer of Heroes | synergy | Theme-fitting lifegain synergy creature.
- 1 Prideful Feastling | synergy | Creature synergy for the lifegain plan.
- 1 Rosie Cotton of South Lane | synergy | Lifegain-focused creature synergy.
- 1 Second Breakfast | synergy | Theme-supporting lifegain synergy spell.
- 1 Wanderbrine Preacher | synergy | Creature synergy for the lifegain plan.
- 1 Angel of Invention | threat | Flying creature threat for closing games.
- 1 Bill the Pony | threat | Creature threat that adds pressure to the board.
- 1 Dawnhand Eulogist | threat | Creature threat for the deck’s midgame.
- 1 Foggy Swamp Hunters | threat | Creature threat that helps build board pressure.
- 1 Gollum, Silent Slinker // Meager Meal | threat | Black creature threat with an attached spell option.
- 1 Lyra Dawnbringer | threat | Powerful Angel threat for finishing games.
- 1 Minwu, White Mage | threat | Legendary creature threat that fits the deck’s plan.
- 1 Rabaroo Troop | threat | Creature threat for maintaining combat pressure.
- 1 Rooftop Percher | threat | Creature threat that adds another board presence.
- 1 Shattered Angel | threat | Angel threat aligned with the lifegain shell.
- 1 Sneering Shadewriter | threat | Black creature threat for the deck’s curve.
- 1 Victory's Herald | threat | Large Angel threat for closing games.
- 1 Austere Command | wipe | Flexible board reset.
- 1 Fumigate | wipe | Board reset that fits the lifegain plan.
- 1 Vanquish the Horde | wipe | Efficient creature-board reset.

</details>

### 16. a tight budget, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 229 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $23.03 to buy, $47.29 the whole deck.

**Summary:** This is a board-focused aristocrats deck that develops mana and cards, then combines its synergy pieces with a steady stream of threats. It wins by maintaining pressure through that creature-centered plan while using targeted removal, interaction, and wipes to keep opposing boards in check. Its tradeoff is a focused, incremental game plan rather than expensive standalone cards or a highly elaborate mana base.

- [INFO] `bracket_prose_rules`: bracket 2 (Core, 8+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [WARN] `not_owned`: Aron, Benalia's Ruin: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ayli, Eternal Pilgrim: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bartolomé del Presidio: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bastion of Remembrance: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Blood Artist: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Elas il-Kor, Sadistic Pilgrim: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Falkenrath Noble: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Relic Vial: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Skullport Merchant: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Woe Strider: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Yahenni, Undying Partisan: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Zulaport Cutthroat: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Disciple of Bolas: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Erebos, Bleak-Hearted: the deck needs 1, the collection has 0
- [WARN] `not_owned`: High-Society Hunter: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Lord Skitter's Butcher: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Lord of the Forsaken: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Meathook Massacre II: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Old Flitterfang: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ruthless Lawbringer: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ruthless Technomancer: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Shadowheart, Dark Justiciar: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Shilgengar, Sire of Famine: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Sivriss, Nightmare Speaker: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 2.90 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Command Tower | land | Provides a land slot in the mana base.
- 1 Exotic Orchard | land | Provides a land slot in the mana base.
- 1 Path of Ancestry | land | Provides a land slot in the mana base.
- 1 Ash Barrens | land | Provides a land slot in the mana base.
- 1 Castle Doom | land | Provides a land slot in the mana base.
- 1 Escape Tunnel | land | Provides a land slot in the mana base.
- 1 Evolving Wilds | land | Provides a land slot in the mana base.
- 1 Fabled Passage | land | Provides a land slot in the mana base.
- 1 Field of Ruin | land | Provides a land slot in the mana base.
- 1 Ghost Quarter | land | Provides a land slot in the mana base.
- 1 Ifnir Deadlands | land | Provides a land slot in the mana base.
- 1 Midgar, City of Mako // Reactor Raid | land | Provides a land slot in the mana base.
- 1 Opal Palace | land | Provides a land slot in the mana base.
- 1 Scavenger Grounds | land | Provides a land slot in the mana base.
- 1 Secluded Courtyard | land | Provides a land slot in the mana base.
- 1 Sidequest: Catch a Fish // Cooking Campsite | land | Provides a land slot in the mana base.
- 1 Thriving Moor | land | Provides a land slot in the mana base.
- 1 Terramorphic Expanse | land | Provides a land slot in the mana base.
- 1 The Gold Saucer | land | Provides a land slot in the mana base.
- 1 The Grey Havens | land | Provides a land slot in the mana base.
- 1 Unclaimed Territory | land | Provides a land slot in the mana base.
- 1 Vibrant Cityscape | land | Provides a land slot in the mana base.
- 1 Windbrisk Heights | land | Provides a land slot in the mana base.
- 1 Spire of Industry | land | Provides a land slot in the mana base.
- 6 Plains | land | Provides basic land slots for the mana base.
- 6 Swamp | land | Provides basic land slots for the mana base.
- 1 Arcane Signet | ramp | Provides the listed ramp role.
- 1 Astral Cornucopia | ramp | Provides the listed ramp role.
- 1 Chromatic Lantern | ramp | Provides the listed ramp role.
- 1 Commander's Sphere | ramp | Provides the listed ramp role.
- 1 Deadly Dispute | ramp | Provides the listed ramp role.
- 1 Fellwar Stone | ramp | Provides the listed ramp role.
- 1 Inherited Envelope | ramp | Provides the listed ramp role.
- 1 Sol Ring | ramp | Provides the listed ramp role.
- 1 Thought Vessel | ramp | Provides the listed ramp role.
- 1 White Auracite | ramp | Provides the listed ramp role.
- 1 Grave Venerations | draw | Provides the listed draw role.
- 1 Inspiring Overseer | draw | Provides the listed draw role.
- 1 Nasty End | draw | Provides the listed draw role.
- 1 Painful Truths | draw | Provides the listed draw role.
- 1 Wall of Omens | draw | Provides the listed draw role.
- 1 Mask of Memory | draw | Provides the listed draw role.
- 1 Foot Chopper | draw | Provides the listed draw role.
- 1 Stone of Erech | draw | Provides the listed draw role.
- 1 Tome of Legends | draw | Provides the listed draw role.
- 1 Lembas | draw | Provides the listed draw role.
- 1 Bastion Protector | interaction | Provides the listed interaction role.
- 1 Frontline Medic | interaction | Provides the listed interaction role.
- 1 Gift of Immortality | interaction | Provides the listed interaction role.
- 1 Swiftfoot Boots | interaction | Provides the listed interaction role.
- 1 Take Up the Shield | interaction | Provides the listed interaction role.
- 1 Ultimate Magic: Holy | interaction | Provides the listed interaction role.
- 1 Bitter Triumph | removal | Provides the listed removal role.
- 1 Claim the Precious | removal | Provides the listed removal role.
- 1 Crib Swap | removal | Provides the listed removal role.
- 1 Destroy Evil | removal | Provides the listed removal role.
- 1 Generous Gift | removal | Provides the listed removal role.
- 1 Infernal Grasp | removal | Provides the listed removal role.
- 1 Swords to Plowshares | removal | Provides the listed removal role.
- 1 Vayne's Treachery | removal | Provides the listed removal role.
- 1 Austere Command | wipe | Provides the listed wipe role.
- 1 Black Sun's Zenith | wipe | Provides the listed wipe role.
- 1 Dusk // Dawn | wipe | Provides the listed wipe role.
- 1 Aron, Benalia's Ruin | synergy | Fills a synergy role in the requested aristocrats plan.
- 1 Ayli, Eternal Pilgrim | synergy | Fills a synergy role in the requested aristocrats plan.
- 1 Bartolomé del Presidio | synergy | Fills a synergy role in the requested aristocrats plan.
- 1 Bastion of Remembrance | synergy | Fills a synergy role in the requested aristocrats plan.
- 1 Blood Artist | synergy | Fills a synergy role in the requested aristocrats plan.
- 1 Elas il-Kor, Sadistic Pilgrim | synergy | Fills a synergy role in the requested aristocrats plan.
- 1 Falkenrath Noble | synergy | Fills a synergy role in the requested aristocrats plan.
- 1 Gollum, Patient Plotter | synergy | Fills a synergy role in the requested aristocrats plan.
- 1 Nimble Hobbit | synergy | Fills a synergy role in the requested aristocrats plan.
- 1 Relic Vial | synergy | Fills a synergy role in the requested aristocrats plan.
- 1 Skullport Merchant | synergy | Fills a synergy role in the requested aristocrats plan.
- 1 Woe Strider | synergy | Fills a synergy role in the requested aristocrats plan.
- 1 Yahenni, Undying Partisan | synergy | Fills a synergy role in the requested aristocrats plan.
- 1 Zulaport Cutthroat | synergy | Fills a synergy role in the requested aristocrats plan.
- 1 Disciple of Bolas | threat | Fills a threat role for the deck's board-focused plan.
- 1 Erebos, Bleak-Hearted | threat | Fills a threat role for the deck's board-focused plan.
- 1 High-Society Hunter | threat | Fills a threat role for the deck's board-focused plan.
- 1 Lord Skitter's Butcher | threat | Fills a threat role for the deck's board-focused plan.
- 1 Lord of the Forsaken | threat | Fills a threat role for the deck's board-focused plan.
- 1 Meathook Massacre II | threat | Fills a threat role for the deck's board-focused plan.
- 1 Old Flitterfang | threat | Fills a threat role for the deck's board-focused plan.
- 1 Ruthless Lawbringer | threat | Fills a threat role for the deck's board-focused plan.
- 1 Ruthless Technomancer | threat | Fills a threat role for the deck's board-focused plan.
- 1 Shadowheart, Dark Justiciar | threat | Fills a threat role for the deck's board-focused plan.
- 1 Shilgengar, Sire of Famine | threat | Fills a threat role for the deck's board-focused plan.
- 1 Sivriss, Nightmare Speaker | threat | Fills a threat role for the deck's board-focused plan.

</details>

### 17. upgrade a precon, owned first

Format: Commander. Theme: goblins. Pool: owned_first. Shortlist: 196 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $78.47 to buy, $266.95 the whole deck.

**Summary:** This Goblin Storm deck develops a Goblin board, builds toward explosive spell-heavy turns, and keeps pressure on opponents with Goblin threats backed by direct removal and sweeping answers. It wins by leaning into its Goblin and storm synergies, with card flow and mana bursts helping assemble stronger sequences. The tradeoff is that its focused, proactive plan leaves fewer slots for broad defensive tools and long-game flexibility.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 2.66 over 62 nonland cards
- [INFO] `basics_added`: the list was 1 card short, so the builder added 1 basic land

<details><summary>The deck list</summary>

- 1 Ancestors' Aid | synergy | A retained synergy piece for the Goblin Storm plan.
- 1 Ancestral Anger | draw | Adds efficient card flow to support the spell-focused plan.
- 1 Arena of Glory | land | A retained land in the precon mana base.
- 1 Battle Hymn | ramp | A retained burst-mana piece for explosive turns.
- 1 Blasphemous Act | wipe | A retained reset button when the board gets crowded.
- 1 Boggart Shenanigans | removal | A retained Goblin-focused removal piece.
- 1 Broadside Bombardiers | removal | A retained Goblin removal piece.
- 1 Castle Embereth | land | A retained land in the precon mana base.
- 1 Chaos Warp | removal | A retained flexible removal spell.
- 1 Conspicuous Snoop | synergy | A retained Goblin synergy piece.
- 1 Crimson Wisps | draw | A retained card-flow spell for the storm plan.
- 1 Cunning Maneuver | draw | Adds card flow for chaining through active turns.
- 1 Daring Discovery | synergy | A retained support piece for the Goblin Storm plan.
- 1 Den of the Bugbear | land | A retained land in the precon mana base.
- 1 Dragon Fodder | synergy | A retained Goblin Storm synergy piece.
- 1 Dwarven Mine | land | A retained land in the precon mana base.
- 1 Empty the Warrens | synergy | A retained storm-focused synergy piece.
- 1 Expedite | draw | A retained card-flow spell for active turns.
- 1 Fable of the Mirror-Breaker // Reflection of Kiki-Jiki | ramp | Adds ramp while staying aligned with the Goblin theme.
- 1 Faithless Looting | draw | A retained card-flow spell for finding key pieces.
- 1 Fists of Flame | draw | A retained card-flow spell for the storm plan.
- 1 Forgotten Cave | land | A retained land in the precon mana base.
- 1 Fountainport | land | A retained land in the precon mana base.
- 1 Frontline Heroism | synergy | A retained support piece for the Goblin Storm plan.
- 1 Gempalm Incinerator | removal | A retained Goblin removal piece.
- 1 General Kreat, the Boltbringer | synergy | A retained Goblin synergy piece.
- 1 Glimpse the Impossible | ramp | A retained ramp piece for explosive turns.
- 1 Goblin Bombardment | removal | A retained Goblin-focused removal piece.
- 1 Goblin Burrows | land | A retained land in the precon mana base.
- 1 Goblin Bushwhacker | synergy | A retained Goblin synergy piece.
- 1 Goblin Chieftain | synergy | A retained Goblin synergy piece.
- 1 Goblin Dark-Dwellers | threat | A retained Goblin threat.
- 1 Goblin Lackey | synergy | A retained Goblin synergy piece.
- 1 Goblin Matron | synergy | A retained Goblin synergy piece.
- 1 Goblin Negotiation | removal | A retained Goblin removal piece.
- 1 Goblin Trashmaster | removal | A retained Goblin removal piece.
- 1 Goblin Warchief | synergy | A retained Goblin synergy piece.
- 1 Grapeshot | removal | A retained storm-oriented removal spell.
- 1 Haze of Rage | synergy | A retained storm-focused synergy piece.
- 1 Hidden Volcano | land | A retained land in the precon mana base.
- 1 Idol of Oblivion | draw | A retained card-flow artifact.
- 1 Impact Tremors | wincon | A retained Goblin Storm payoff.
- 1 Kher Keep | land | A retained land in the precon mana base.
- 1 Krenko's Command | synergy | A retained Goblin Storm synergy piece.
- 1 Krenko, Mob Boss | threat | A retained Goblin threat.
- 1 Lightning Greaves | interaction | Adds protection for the deck's important creatures.
- 22 Mountain | land | Provides the deck's basic red mana base.
- 1 Mana Geyser | ramp | A retained burst-mana piece for explosive turns.
- 1 Mogg War Marshal | synergy | A retained Goblin synergy piece.
- 1 Pashalik Mons | removal | A retained Goblin removal piece.
- 1 Past in Flames | synergy | A retained spell-focused synergy piece.
- 1 Path of Ancestry | land | Adds a Goblin-aligned land to the mana base.
- 1 Quest for the Goblin Lord | synergy | A retained Goblin synergy piece.
- 1 Renegade Tactics | draw | Adds card flow for chaining through active turns.
- 1 Reliquary Tower | land | A retained land in the precon mana base.
- 1 Roaming Throne | synergy | A retained Goblin Storm support piece.
- 1 Ruby Medallion | ramp | A retained mana-development piece for the spell plan.
- 1 Rundvelt Hordemaster | synergy | A retained Goblin synergy piece.
- 1 Rush the Room | synergy | Adds another spell-focused synergy piece.
- 1 Sazacap's Brew | draw | A retained card-flow spell.
- 1 Secluded Courtyard | land | Adds a Goblin-aligned land to the mana base.
- 1 Seething Song | ramp | A retained burst-mana spell for explosive turns.
- 1 Shinka, the Bloodsoaked Keep | land | Adds a red land while preserving the mana-base shape.
- 1 Siege-Gang Commander | removal | A retained Goblin removal piece.
- 1 Siege-Gang Lieutenant | removal | A retained Goblin removal piece.
- 1 Skirk Prospector | ramp | A retained Goblin ramp piece.
- 1 Skullclamp | draw | A retained card-flow artifact.
- 1 Smoldering Crater | land | A retained land in the precon mana base.
- 1 Sol Ring | ramp | A retained early mana-development piece.
- 1 Spreading Insurrection | synergy | A retained support piece for the Goblin Storm plan.
- 1 Storm-Kiln Artist | ramp | A retained ramp piece for the spell-focused plan.
- 1 Swiftfoot Boots | interaction | A retained protection piece for important creatures.
- 1 Throne of Eldraine | ramp | A retained mana-development artifact.
- 1 Vanquisher's Banner | draw | Adds card flow while staying aligned with the Goblin theme.
- 1 Vandalblast | wipe | A retained artifact-focused board reset.
- 1 War Room | land | A retained land in the precon mana base.
- 1 Wild Ride | synergy | A retained support piece for the Goblin Storm plan.
- 1 Witch's Mark | draw | A retained card-flow spell.

</details>

### 18. upgrade a precon, any card

Format: Commander. Theme: turtles. Pool: any_card. Shortlist: 303 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $227.51 to buy, $227.51 the whole deck.

**Summary:** This Turtle Power deck keeps the precon’s character-driven Turtle theme intact, building its board through Turtle-themed synergy, mana acceleration, and card draw before closing with its many Turtle threats. The upgraded land base and added interaction improve consistency and give the deck cleaner answers when opponents interfere. It gives up a more narrowly focused finish in order to preserve the precon’s broad cast of heroes, allies, artifacts, and theme pieces.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [WARN] `land_count`: 44 lands: the guide range for this format is 34 to 38
- [INFO] `curve_summary`: average mana value 3.33 over 55 nonland cards

<details><summary>The deck list</summary>

- 1 Acidic Slime | other | Retained as a precon creature piece.
- 1 April O'Neil, Live on the Scene | draw | Retained card draw from the precon.
- 1 Arcane Signet | ramp | Retained mana acceleration from the precon.
- 1 Ash Barrens | land | Retained precon land.
- 1 Assassin's Trophy | removal | Retained removal from the precon.
- 1 Baxter, Fly in the Ointment | other | Retained as a precon character piece.
- 1 Bebop, Skull & Crossbones | other | Retained as a precon character piece.
- 1 Big Apple, 3 a.m. | land | Retained precon land.
- 1 Big Mother Mouser | other | Retained as a precon artifact creature.
- 1 Biogenic Ooze | other | Retained as a precon creature piece.
- 1 Blasphemous Act | wipe | Retained board wipe from the precon.
- 1 Breeding Pool | land | Added to strengthen the land base.
- 1 Casey Jones, Back Alley Brute | other | Retained as a precon character piece.
- 1 Chromatic Lantern | ramp | Retained mana acceleration from the precon.
- 1 Cinder Glade | land | Retained precon land.
- 1 City of Brass | land | Retained precon land.
- 1 Command Tower | land | Retained precon land.
- 1 Continue? | other | Retained as a precon spell.
- 1 Corpsejack Menace | other | Retained as a precon creature piece.
- 1 Cultivate | ramp | Retained mana acceleration from the precon.
- 1 Dimension X Pizzasaur | other | Retained as a precon artifact creature.
- 1 Donatello, the Brains | synergy | Retained Turtle-themed synergy from the precon.
- 1 Double Jump // Flying Kick | other | Retained as a precon spell.
- 1 Endless Foot Assault | other | Retained as a precon enchantment.
- 1 Escape Tunnel | land | Retained precon land.
- 1 Everything Pizza | other | Retained as a precon artifact.
- 1 Fierce Guardianship | interaction | Added efficient interaction.
- 4 Forest | land | Retained basic land base.
- 1 Grand Coliseum | land | Retained precon land.
- 1 Hallowed Fountain | land | Added to strengthen the land base.
- 1 Harmonize | other | Retained as a precon spell.
- 1 Here Comes a New Hero! | other | Retained as a precon spell.
- 1 Hidden Hideout | land | Retained precon land.
- 1 High Score | other | Retained as a precon enchantment.
- 4 Island | land | Retained basic land base.
- 1 Irma, Part-Time Mutant | other | Retained as a precon character piece.
- 1 Krang, the All-Powerful | other | Retained as a precon artifact creature.
- 1 Leatherhead, Iron Gator | other | Retained as a precon character piece.
- 1 Lessons from Life | other | Retained as a precon spell.
- 1 Level Up | other | Retained as a precon enchantment.
- 1 Lita, Little Orphan Amphibian | synergy | Retained Turtle-themed synergy from the precon.
- 1 Michelangelo, the Heart | synergy | Retained Turtle-themed synergy from the precon.
- 1 Mole Module | other | Retained as a precon vehicle.
- 1 Mona Lisa, Science Geek | other | Retained as a precon character piece.
- 3 Mountain | land | Retained basic land base.
- 1 Ninja Pizza | other | Retained as a precon enchantment.
- 1 Overgrown Tomb | land | Added to strengthen the land base.
- 1 Path of Ancestry | land | Retained precon land.
- 4 Plains | land | Retained basic land base.
- 1 Rain-Slicked Copse | land | Retained precon land.
- 1 Raphael, the Muscle | threat | Retained Turtle-themed threat from the precon.
- 1 Rat King, Pale Piper | other | Retained as a precon character piece.
- 1 Ray Fillet, Wave Warrior | other | Retained as a precon character piece.
- 1 Rhystic Study | draw | Added reliable card draw.
- 1 Roadkill Rodney | other | Retained as a precon artifact creature.
- 1 Rocksteady, Mutant Marauder | other | Retained as a precon character piece.
- 1 Rootbound Crag | land | Retained precon land.
- 1 Shellshock | other | Retained as a precon spell.
- 1 Shredder, Shadow Master | other | Retained as a precon character piece.
- 1 Smoldering Marsh | land | Retained precon land.
- 1 Sodden Verdure | land | Retained precon land.
- 1 Sol Ring | ramp | Retained mana acceleration from the precon.
- 1 Special Move | other | Retained as a precon spell.
- 1 Spire Garden | land | Retained precon land.
- 1 Splinter, the Mentor | other | Retained as a precon character piece.
- 1 Steelbane Hydra | removal | Retained removal from the precon.
- 1 Sunken Hollow | land | Retained precon land.
- 1 Super Combo | other | Retained as a precon spell.
- 4 Swamp | land | Retained basic land base.
- 1 Swords to Plowshares | removal | Added efficient removal.
- 1 Swift Demise | other | Retained as a precon spell.
- 1 Tempestra, Dame of Games | other | Retained as a precon character piece.
- 1 Thriving Grove | land | Retained precon land.
- 1 Thriving Isle | land | Retained precon land.
- 1 Thriving Moor | land | Retained precon land.
- 1 Together Forever | other | Retained as a precon enchantment.
- 1 Tokka & Rahzar, Unsupervised | ramp | Retained Turtle-themed mana acceleration from the precon.
- 1 Turtle Lair | land | Retained precon land.
- 1 Undergrowth Stadium | land | Retained precon land.
- 1 Vanquish the Horde | other | Retained as a precon spell.
- 1 Vernal Fen | land | Retained precon land.
- 1 Vibrant Cityscape | land | Retained precon land.
- 1 Vigor | other | Retained as a precon creature piece.
- 1 Voracious Hydra | other | Retained as a precon creature piece.
- 1 Wave Goodbye | other | Retained as a precon spell.

</details>

