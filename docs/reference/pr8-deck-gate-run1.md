# PR-8 deck gate

Run date: 2026-08-27. Card snapshot: 2026-08-24.

Verdict: FAIL. 10 of 12 decks passed every block check, and 1 invented names reached the user. Both bars are zero tolerance.

## Summary

| Measure | Value |
|---|---|
| Prompts | 12 |
| Decks returned | 12 |
| Decks with no block finding | 10 |
| Invented names that reached the user | 1 |
| Decks that needed the repair turn | 3 |
| Errors | 0 |
| Prompt version | 2 |
| Calls | 15 |
| Cost | $0.9117 |
| Time | 732 seconds |

## Findings by code

A block stops the deck. A warning and a note are reports.

| Code | Count |
|---|---|
| `curve_summary` | 12 |
| `bracket_prose_rules` | 6 |
| `deck_size` | 2 |
| `not_owned` | 2 |

By severity: BLOCK 2. WARN 2. INFO 18. 

## Decks

### 1. lifegain Commander, any card

Format: Commander. Theme: lifegain. Pool: any_card. Shortlist: 301 names.

Cards: 99. Repair turn: false. Block findings: 0.

**Summary:** This deck builds around repeated lifegain and Karlov, using those gains to support a broad board of lifegain-synergy permanents and threatening creatures. It wins by establishing those threats, pressing combat advantages, and turning the deck's lifegain plan into sustained pressure while removal and sweepers keep opposing boards in check. The tradeoff is that the deck commits many slots to its lifegain engine, so its strongest games depend on assembling and protecting that connected board.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 3.56 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Aclazotz, Deepest Betrayal // Temple of the Dead | land | Fills a land slot in the mana base.
- 1 Adventurer's Inn | land | Fills a land slot in the mana base.
- 1 Arguel's Blood Fast // Temple of Aclazotz | land | Fills a land slot in the mana base.
- 1 Azor's Gateway // Sanctum of the Sun | land | Fills a land slot in the mana base.
- 1 Blighted Steppe | land | Fills a land slot in the mana base.
- 1 Brokers Hideout | land | Fills a land slot in the mana base.
- 1 Cabaretti Courtyard | land | Fills a land slot in the mana base.
- 1 Diamond Valley | land | Fills a land slot in the mana base.
- 1 Glasswing Grace // Age-Graced Chapel | land | Fills a land slot in the mana base.
- 1 Glimmerpost | land | Fills a land slot in the mana base.
- 1 Hidden Hideout | land | Fills a land slot in the mana base.
- 1 High Market | land | Fills a land slot in the mana base.
- 1 Hostile Hostel // Creeping Inn | land | Fills a land slot in the mana base.
- 1 Interplanar Beacon | land | Fills a land slot in the mana base.
- 1 Inventors' Fair | land | Fills a land slot in the mana base.
- 1 Kabira Crossroads | land | Fills a land slot in the mana base.
- 1 Legion's Landing // Adanto, the First Fort | land | Fills a land slot in the mana base.
- 1 Miren, the Moaning Well | land | Fills a land slot in the mana base.
- 1 Neglected Manor | land | Fills a land slot in the mana base.
- 1 Nomad Stadium | land | Fills a land slot in the mana base.
- 1 Obscura Storefront | land | Fills a land slot in the mana base.
- 1 Phyrexia's Core | land | Fills a land slot in the mana base.
- 1 Plaza of Harmony | land | Fills a land slot in the mana base.
- 1 Radiant Fountain | land | Fills a land slot in the mana base.
- 1 Restless Fortress | land | Fills a land slot in the mana base.
- 1 Scoured Barrens | land | Fills a land slot in the mana base.
- 1 Seraph Sanctuary | land | Fills a land slot in the mana base.
- 1 Shambling Vent | land | Fills a land slot in the mana base.
- 1 Skyclave Cleric // Skyclave Basilica | land | Fills a land slot in the mana base.
- 1 Springjack Pasture | land | Fills a land slot in the mana base.
- 1 Starlit Sanctum | land | Fills a land slot in the mana base.
- 1 Tomb of the Spirit Dragon | land | Fills a land slot in the mana base.
- 1 Vault of the Archangel | land | Fills a land slot in the mana base.
- 1 Westvale Abbey // Ormendahl, Profane Prince | land | Fills a land slot in the mana base.
- 1 Witch's Clinic | land | Fills a land slot in the mana base.
- 1 Zof Consumption // Zof Bloodbog | land | Fills a land slot in the mana base.
- 1 Alhammarret's Archive | draw | Provides draw support for the deck.
- 1 Archivist of Oghma | draw | Provides draw support for the deck.
- 1 Ayara, First of Locthwain | draw | Provides draw support for the deck.
- 1 Convalescent Care | draw | Provides draw support for the deck.
- 1 Dawn of Hope | draw | Provides draw support for the deck.
- 1 Enduring Innocence | draw | Provides draw support for the deck.
- 1 Exemplar of Light | draw | Provides draw support for the deck.
- 1 Mangara, the Diplomat | draw | Provides draw support for the deck.
- 1 Markov Purifier | draw | Provides draw support for the deck.
- 1 Well of Lost Dreams | draw | Provides draw support for the deck.
- 1 Altar of the Pantheon | ramp | Provides ramp for the deck.
- 1 Angel of Indemnity | ramp | Provides ramp for the deck.
- 1 Bounty Board | ramp | Provides ramp for the deck.
- 1 Carmen, Cruel Skymarcher | ramp | Provides ramp for the deck.
- 1 Crypt Ghast | ramp | Provides ramp for the deck.
- 1 Hierophant's Chalice | ramp | Provides ramp for the deck.
- 1 Nuka-Cola Vending Machine | ramp | Provides ramp for the deck.
- 1 Orazca Relic | ramp | Provides ramp for the deck.
- 1 Pristine Talisman | ramp | Provides ramp for the deck.
- 1 Redemption Choir | ramp | Provides ramp for the deck.
- 1 Alseid of Life's Bounty | interaction | Provides interaction for the deck.
- 1 Faith's Shield | interaction | Provides interaction for the deck.
- 1 Metropolis Reformer | interaction | Provides interaction for the deck.
- 1 Restoration Magic | interaction | Provides interaction for the deck.
- 1 Sword of Light and Shadow | interaction | Provides interaction for the deck.
- 1 Werefox Bodyguard | interaction | Provides interaction for the deck.
- 1 Aetherflux Reservoir | removal | Provides removal for the deck.
- 1 Ayli, Eternal Pilgrim | removal | Provides removal for the deck.
- 1 Murderous Rider // Swift End | removal | Provides removal for the deck.
- 1 Solitude | removal | Provides removal for the deck.
- 1 Tithing Blade // Consuming Sepulcher | removal | Provides removal for the deck.
- 1 Umezawa's Jitte | removal | Provides removal for the deck.
- 1 Vona, Butcher of Magan | removal | Provides removal for the deck.
- 1 Witch of the Moors | removal | Provides removal for the deck.
- 1 Aerith Gainsborough | synergy | Supports the deck's lifegain synergy plan.
- 1 Ajani's Pridemate | synergy | Supports the deck's lifegain synergy plan.
- 1 Angel of Vitality | synergy | Supports the deck's lifegain synergy plan.
- 1 Angelic Accord | synergy | Supports the deck's lifegain synergy plan.
- 1 Blood Artist | synergy | Supports the deck's lifegain synergy plan.
- 1 Cleric Class | synergy | Supports the deck's lifegain synergy plan.
- 1 Cleric of Life's Bond | synergy | Supports the deck's lifegain synergy plan.
- 1 Heliod, Sun-Crowned | synergy | Supports the deck's lifegain synergy plan.
- 1 Resplendent Angel | synergy | Supports the deck's lifegain synergy plan.
- 1 Righteous Valkyrie | synergy | Supports the deck's lifegain synergy plan.
- 1 Sanguine Bond | synergy | Supports the deck's lifegain synergy plan.
- 1 Vito, Thorn of the Dusk Rose | synergy | Supports the deck's lifegain synergy plan.
- 1 Vizkopa Guildmage | synergy | Supports the deck's lifegain synergy plan.
- 1 Voice of the Blessed | synergy | Supports the deck's lifegain synergy plan.
- 1 Archangel of Thune | threat | Serves as a major threat for the deck.
- 1 Astarion, the Decadent | threat | Serves as a major threat for the deck.
- 1 Attended Healer | threat | Serves as a major threat for the deck.
- 1 Celestine, the Living Saint | threat | Serves as a major threat for the deck.
- 1 Cliffhaven Vampire | threat | Serves as a major threat for the deck.
- 1 Defiant Bloodlord | threat | Serves as a major threat for the deck.
- 1 Divinity of Pride | threat | Serves as a major threat for the deck.
- 1 Exalted Sunborn | threat | Serves as a major threat for the deck.
- 1 Liesa, Forgotten Archangel | threat | Serves as a major threat for the deck.
- 1 Lyra Dawnbringer | threat | Serves as a major threat for the deck.
- 1 Nykthos Paragon | threat | Serves as a major threat for the deck.
- 1 Valkyrie Harbinger | threat | Serves as a major threat for the deck.
- 1 Ajani, Strength of the Pride | wipe | Provides a board wipe option.
- 1 Fumigate | wipe | Provides a board wipe option.
- 1 Kaya's Wrath | wipe | Provides a board wipe option.

</details>

### 2. aristocrats Commander, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 177 names.

Cards: 99. Repair turn: false. Block findings: 0.

**Summary:** This Karlov deck plays a creature-focused aristocrats game, building resources while using removal and board resets to keep the table from getting ahead. It aims to win through repeated sacrifice-driven value and steady board pressure, with protective pieces helping preserve the core of the plan. It gives up a deep dedicated finisher package in exchange for a flexible, attrition-oriented game.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [WARN] `not_owned`: Karlov of the Ghost Council: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 2.86 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Abandoned Air Temple | land | Listed land for the mana base.
- 1 Adventurer's Inn | land | Listed land for the mana base.
- 1 Ash Barrens | land | Listed land for the mana base.
- 1 Bonders' Enclave | land | Listed land for the mana base.
- 1 Castle Doom | land | Listed land for the mana base.
- 1 Castle Locthwain | land | Listed land for the mana base.
- 1 City of Brass | land | Listed land for the mana base.
- 1 Command Tower | land | Listed land for the mana base.
- 1 Escape Tunnel | land | Listed land for the mana base.
- 1 Evolving Wilds | land | Listed land for the mana base.
- 1 Exotic Orchard | land | Listed land for the mana base.
- 1 Fabled Passage | land | Listed land for the mana base.
- 1 Field of Ruin | land | Listed land for the mana base.
- 1 Ghost Quarter | land | Listed land for the mana base.
- 1 Ifnir Deadlands | land | Listed land for the mana base.
- 1 Marsh Flats | land | Listed land for the mana base.
- 1 Minas Tirith | land | Listed land for the mana base.
- 1 Opal Palace | land | Listed land for the mana base.
- 1 Path of Ancestry | land | Listed land for the mana base.
- 1 Plaza of Heroes | land | Listed land for the mana base.
- 1 Reliquary Tower | land | Listed land for the mana base.
- 1 Riveteers Overlook | land | Listed land for the mana base.
- 1 Rogue's Passage | land | Listed land for the mana base.
- 1 Scavenger Grounds | land | Listed land for the mana base.
- 1 Secluded Courtyard | land | Listed land for the mana base.
- 1 Spire of Industry | land | Listed land for the mana base.
- 1 Takenuma, Abandoned Mire | land | Listed land for the mana base.
- 1 Temple of the False God | land | Listed land for the mana base.
- 1 Terramorphic Expanse | land | Listed land for the mana base.
- 1 The Gold Saucer | land | Listed land for the mana base.
- 1 The Grey Havens | land | Listed land for the mana base.
- 1 Thriving Moor | land | Listed land for the mana base.
- 1 Unclaimed Territory | land | Listed land for the mana base.
- 1 Vibrant Cityscape | land | Listed land for the mana base.
- 1 War Room | land | Listed land for the mana base.
- 1 Windbrisk Heights | land | Listed land for the mana base.
- 1 Buster Sword | draw | Listed draw that helps keep cards flowing.
- 1 Call of the Ring | draw | Listed draw for sustained resources.
- 1 Cirith Ungol Patrol | draw | Listed draw for sustained resources.
- 1 Exemplar of Light | draw | Listed draw attached to a creature.
- 1 Grave Venerations | draw | Listed draw for the long game.
- 1 Idol of Oblivion | draw | Listed draw for the aristocrats plan.
- 1 Inspiring Overseer | draw | Listed draw attached to a creature.
- 1 Lembas | draw | Listed draw for steady resources.
- 1 Mask of Memory | draw | Listed draw for steady resources.
- 1 Night's Whisper | draw | Listed draw for efficient card flow.
- 1 Painful Truths | draw | Listed draw for efficient card flow.
- 1 Puresteel Paladin | draw | Listed draw attached to a creature.
- 1 Skullclamp | draw | Listed draw that suits an aristocrats plan.
- 1 Stone of Erech | draw | Listed draw for sustained resources.
- 1 Tome of Legends | draw | Listed draw for the long game.
- 1 Wall of Omens | draw | Listed draw attached to a creature.
- 1 Arcane Signet | ramp | Listed ramp for early development.
- 1 Astral Cornucopia | ramp | Listed ramp for mana development.
- 1 Bender's Waterskin | ramp | Listed ramp for mana development.
- 1 Blitzball | ramp | Listed ramp for mana development.
- 1 Commander's Sphere | ramp | Listed ramp for mana development.
- 1 Deadly Dispute | ramp | Listed ramp that supports the sacrifice-focused plan.
- 1 Explorer's Scope | ramp | Listed ramp for mana development.
- 1 Fellwar Stone | ramp | Listed ramp for early development.
- 1 Inherited Envelope | ramp | Listed ramp for mana development.
- 1 Relic of Legends | ramp | Listed ramp for mana development.
- 1 Sol Ring | ramp | Listed ramp for early development.
- 1 Springleaf Drum | ramp | Listed ramp for early development.
- 1 Sword of the Animist | ramp | Listed ramp that supports creature pressure.
- 1 Bastion Protector | interaction | Listed interaction to support the commander plan.
- 1 Bronze Guardian | interaction | Listed interaction for board support.
- 1 Champion's Helm | interaction | Listed interaction to protect an important creature.
- 1 Clever Concealment | interaction | Listed interaction for protecting the board.
- 1 Darksteel Plate | interaction | Listed interaction to protect an important creature.
- 1 Gift of Immortality | interaction | Listed interaction that suits a creature-focused plan.
- 1 Lightning Greaves | interaction | Listed interaction to protect an important creature.
- 1 Swiftfoot Boots | interaction | Listed interaction to protect an important creature.
- 1 Angel of Serenity | removal | Listed removal attached to a creature.
- 1 Banishing Light | removal | Listed removal for opposing permanents.
- 1 Bitter Triumph | removal | Listed removal for opposing threats.
- 1 Blowfly Infestation | removal | Listed removal that supports the attrition plan.
- 1 Claim the Precious | removal | Listed removal for opposing threats.
- 1 Contagion Clasp | removal | Listed removal for opposing threats.
- 1 Crib Swap | removal | Listed removal for opposing creatures.
- 1 Destroy Evil | removal | Listed removal for opposing permanents.
- 1 Dismember | removal | Listed removal for opposing creatures.
- 1 Dispatch | removal | Listed removal for opposing threats.
- 1 Fiend Hunter | removal | Listed removal attached to a creature.
- 1 Generous Gift | removal | Listed removal for opposing permanents.
- 1 Get Lost | removal | Listed removal for opposing permanents.
- 1 Austere Command | wipe | Listed wipe for resetting difficult boards.
- 1 Black Sun's Zenith | wipe | Listed wipe for resetting creature boards.
- 1 Dusk // Dawn | wipe | Listed wipe for resetting creature boards.
- 1 Fumigate | wipe | Listed wipe for resetting creature boards.
- 1 Martial Coup | wipe | Listed wipe for resetting difficult boards.
- 1 Vanquish the Horde | wipe | Listed wipe for resetting creature boards.
- 1 Arcade Cabinet | synergy | Listed synergy for the aristocrats plan.
- 1 Denethor, Ruling Steward | synergy | Listed synergy for the aristocrats plan.
- 1 Gollum, Patient Plotter | synergy | Listed synergy for the aristocrats plan.
- 1 Gríma Wormtongue | synergy | Listed synergy for the aristocrats plan.
- 1 Heirloom Auntie | synergy | Listed synergy for the aristocrats plan.
- 1 Joo Dee, One of Many | synergy | Listed synergy for the aristocrats plan.
- 1 Phantom Train | synergy | Listed synergy for the aristocrats plan.

</details>

### 3. artifacts Commander, bracket 4

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 298 names.

Cards: 99. Repair turn: true. Block findings: 0.

**Summary:** Urza, Lord High Artificer leads this 99-card list. It contains 36 lands, 10 ramp cards, 10 draw cards, 6 interaction cards, 8 removal cards, 14 synergy cards, 12 threats, and 3 wipes.

- [INFO] `curve_summary`: average mana value 3.94 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Academy Ruins | land | Artifact recursion land.
- 1 Archway of Innovation | land | Artifact-focused mana base land.
- 1 Blinkmoth Nexus | land | Manland that benefits from artifact support.
- 1 Buried Ruin | land | Recovers key artifacts.
- 1 Conqueror's Galleon // Conqueror's Foothold | land | Artifact land utility.
- 1 Fomori Vault | land | Artifact utility land.
- 1 Glimmervoid | land | Reliable colored source in an artifact-heavy list.
- 1 Golden Guardian // Gold-Forge Garrison | land | Threat-producing artifact land.
- 1 Hall of Tagsin | land | Artifact utility land.
- 1 Heap Gate | land | Mana-base utility.
- 1 Inkmoth Nexus | land | Artifact manland and alternate pressure.
- 1 Inventors' Fair | land | Artifact tutoring utility.
- 1 Matzalantli, the Great Door // The Core | land | Filtering artifact land.
- 1 Mech Hangar | land | Vehicle and artifact mana support.
- 1 Mirrex | land | Produces artifact tokens.
- 1 Mishra's Factory | land | Artifact manland.
- 1 Mishra's Foundry | land | Artifact manland.
- 1 Mishra's Workshop | land | Explosive artifact mana source.
- 1 Otawara, Soaring City | land | Utility interaction land.
- 1 Phyrexia's Core | land | Artifact sacrifice outlet.
- 1 Power Depot | land | Artifact land with mana utility.
- 1 Primal Amulet // Primal Wellspring | land | Spell-supporting utility land.
- 1 Public Thoroughfare | land | Mana-base depth.
- 1 Roadside Reliquary | land | Card-advantage utility land.
- 1 Secluded Starforge | land | Artifact-supporting land.
- 1 Spire of Industry | land | Colored mana for an artifact deck.
- 1 Thaumatic Compass // Spires of Orazca | land | Land selection and defensive utility.
- 1 The Everflowing Well // The Myriad Pools | land | Artifact land utility.
- 1 The Gold Saucer | land | Mana-base utility.
- 1 The Monumental Facade | land | Counter utility for artifacts.
- 1 The Mycosynth Gardens | land | Copies important artifacts.
- 1 Treasure Map // Treasure Cove | land | Selection that becomes a card-advantage land.
- 1 Underdark Rift | land | Mana-base utility.
- 1 Urza's Factory | land | Late-game Construct production.
- 1 Urza's Saga | land | Artifact tutoring and Construct pressure.
- 1 Urza's Workshop | land | Artifact mana source.
- 1 Mox Opal | ramp | Premium early artifact acceleration.
- 1 Metalworker | ramp | High-output artifact mana engine.
- 1 Krark-Clan Ironworks | ramp | Converts artifacts into explosive mana.
- 1 Moonsnare Prototype | ramp | Cheap artifact acceleration.
- 1 Inspiring Statuary | ramp | Turns artifacts into spell mana.
- 1 Grand Architect | ramp | Artifact creature mana engine.
- 1 Chief Engineer | ramp | Convoke support for artifact threats.
- 1 Tezzeret the Seeker | ramp | Finds key mana artifacts and untaps them.
- 1 Karn, Legacy Reforged | ramp | Scales mana production with the artifact board.
- 1 Arc Reactor | ramp | Artifact mana acceleration.
- 1 Thoughtcast | draw | Efficient artifact-enabled card advantage.
- 1 Thirst for Knowledge | draw | Instant-speed artifact filtering.
- 1 Reverse Engineer | draw | Large artifact-enabled draw spell.
- 1 Sai, Master Thopterist | draw | Artifact casts create material and cards.
- 1 Vedalken Archmage | draw | Sustained draw from artifact spells.
- 1 Thopter Spy Network | draw | Repeatable Thopter production and draw.
- 1 Thought Monitor | draw | Artifact-enabled draw attached to a body.
- 1 Tezzeret, Artifice Master | draw | Repeatable card advantage and tokens.
- 1 Riddlesmith | draw | Cheap artifact-triggered filtering.
- 1 Nexus of Becoming | draw | Artifact card-advantage engine.
- 1 Assert Authority | interaction | Artifact-scaled stack interaction.
- 1 Disruption Protocol | interaction | Efficient protection for the artifact board.
- 1 Metallic Rebuke | interaction | Low-cost artifact-enabled counterspell.
- 1 Stoic Rebuttal | interaction | Reliable artifact-enabled counterspell.
- 1 Darksteel Forge | interaction | Protects the central artifact board.
- 1 Padeem, Consul of Innovation | interaction | Protects valuable artifacts while supporting cards.
- 1 Aether Spellbomb | removal | Reusable, tutor-friendly creature disruption.
- 1 Resculpt | removal | Flexible answer to artifacts or creatures.
- 1 Ravenform | removal | Exiles troublesome artifacts or creatures.
- 1 Skysovereign, Consul Flagship | removal | Repeatable damage-based removal.
- 1 Portal to Phyrexia | removal | Major permanent removal with a threatening body engine.
- 1 Spine of Ish Sah | removal | Tutorable answer to any problematic permanent.
- 1 Lux Cannon | removal | Repeatable permanent removal with untap support.
- 1 Tormod's Crypt | removal | Zero-cost graveyard interaction.
- 1 Arcbound Ravager | synergy | Sacrifice outlet and modular payoff.
- 1 Emry, Lurker of the Loch | synergy | Recasts cheap artifacts from the graveyard.
- 1 Etherium Sculptor | synergy | Reduces artifact spell costs.
- 1 Foundry Inspector | synergy | Additional artifact cost reduction.
- 1 Mystic Forge | synergy | Extends artifact casting from the library.
- 1 Clock of Omens | synergy | Converts spare artifacts into powerful untaps.
- 1 Unwinding Clock | synergy | Untaps the artifact engine every turn cycle.
- 1 Voltaic Key | synergy | Efficiently untaps mana and utility artifacts.
- 1 Manifold Key | synergy | Untaps major artifacts and enables attacks.
- 1 Scrap Trawler | synergy | Recovers sacrificed and destroyed artifacts.
- 1 Transmute Artifact | synergy | Converts small artifacts into decisive engines.
- 1 Whir of Invention | synergy | Instant-speed artifact tutoring.
- 1 Reshape | synergy | Direct artifact conversion tutor.
- 1 Panharmonicon | synergy | Doubles the deck's artifact creature triggers.
- 1 Kappa Cannoneer | threat | Fast, protected artifact payoff threat.
- 1 Cyberdrive Awakener | threat | Turns the artifact board into a lethal attack.
- 1 Kuldotha Forgemaster | threat | Threat that tutors premier artifacts into play.
- 1 Master Transmuter | threat | Deploys expensive artifacts while protecting them.
- 1 Phyrexian Metamorph | threat | Copies the strongest artifact or creature available.
- 1 Karn, Scion of Urza | threat | Produces large artifact-scaled Constructs.
- 1 Traxos, Scourge of Kroog | threat | Efficient large artifact attacker.
- 1 Lodestone Golem | threat | Large disruptive artifact body.
- 1 Darksteel Juggernaut | threat | Indestructible artifact-scaled attacker.
- 1 Broodstar | threat | Large affinity-based finisher.
- 1 Threefold Thunderhulk | threat | Creates a substantial artifact board presence.
- 1 The Capitoline Triad | threat | Top-end artifact payoff threat.
- 1 Nevinyrral's Disk | wipe | Broad reset accessible through artifact tutoring.
- 1 Oblivion Stone | wipe | Flexible board reset.
- 1 Engineered Explosives | wipe | Efficient scalable artifact sweeper.

</details>

### 4. dinosaur tribal, bracket 2

Format: Commander. Theme: dinosaurs. Pool: any_card. Shortlist: 264 names.

Cards: 99. Repair turn: false. Block findings: 0.

**Summary:** This is a straightforward creature-heavy Dinosaur deck: develop mana, deploy Dinosaur synergy pieces and large threats, then attack to finish games with Gishath leading the charge. It backs that plan with draw, removal, wipes, and interaction, while giving up some flexibility in exchange for committing heavily to creatures and combat.

- [INFO] `bracket_prose_rules`: bracket 2 (Core, 8+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 4.21 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Command Tower | land | A land for the deck's mana base.
- 1 Path of Ancestry | land | A land for the deck's mana base.
- 1 Canopy Vista | land | A land for the deck's mana base.
- 1 Cinder Glade | land | A land for the deck's mana base.
- 1 Sacred Foundry | land | A land for the deck's mana base.
- 1 Stomping Ground | land | A land for the deck's mana base.
- 1 Temple Garden | land | A land for the deck's mana base.
- 1 Sunpetal Grove | land | A land for the deck's mana base.
- 1 Rootbound Crag | land | A land for the deck's mana base.
- 1 Clifftop Retreat | land | A land for the deck's mana base.
- 1 Battlefield Forge | land | A land for the deck's mana base.
- 1 Spectator Seating | land | A land for the deck's mana base.
- 1 Exotic Orchard | land | A land for the deck's mana base.
- 1 City of Brass | land | A land for the deck's mana base.
- 1 Mana Confluence | land | A land for the deck's mana base.
- 1 Evolving Wilds | land | A land for the deck's mana base.
- 1 Terramorphic Expanse | land | A land for the deck's mana base.
- 1 Fabled Passage | land | A land for the deck's mana base.
- 1 Myriad Landscape | land | A land for the deck's mana base.
- 1 Boseiju, Who Endures | land | A land for the deck's mana base.
- 1 Cavern of Souls | land | A land for the deck's mana base.
- 1 Rogue's Passage | land | A land for the deck's mana base.
- 1 War Room | land | A land for the deck's mana base.
- 1 Yavimaya, Cradle of Growth | land | A land for the deck's mana base.
- 1 Arid Mesa | land | A land for the deck's mana base.
- 1 Windswept Heath | land | A land for the deck's mana base.
- 1 Wooded Foothills | land | A land for the deck's mana base.
- 1 Verdant Catacombs | land | A land for the deck's mana base.
- 1 Bloodstained Mire | land | A land for the deck's mana base.
- 1 Marsh Flats | land | A land for the deck's mana base.
- 1 Misty Rainforest | land | A land for the deck's mana base.
- 1 Flooded Strand | land | A land for the deck's mana base.
- 1 Scalding Tarn | land | A land for the deck's mana base.
- 1 Reliquary Tower | land | A land for the deck's mana base.
- 1 Nykthos, Shrine to Nyx | land | A land for the deck's mana base.
- 1 Welcome to . . . // Jurassic Park | land | A land for the deck's mana base.
- 1 Arcane Signet | ramp | A straightforward ramp piece.
- 1 Sol Ring | ramp | A straightforward ramp piece.
- 1 Cultivate | ramp | A ramp spell for developing the deck.
- 1 Kodama's Reach | ramp | A ramp spell for developing the deck.
- 1 Nature's Lore | ramp | A ramp spell for developing the deck.
- 1 Farseek | ramp | A ramp spell for developing the deck.
- 1 Thunderherd Migration | ramp | A Dinosaur-themed ramp spell.
- 1 Drover of the Mighty | ramp | A ramp creature for the Dinosaur deck.
- 1 Topiary Stomper | ramp | A Dinosaur ramp creature.
- 1 Hulking Raptor | ramp | A Dinosaur ramp creature.
- 1 Beast Whisperer | draw | A creature-based draw card.
- 1 Cloudpiercer | draw | A Dinosaur draw card.
- 1 Curious Altisaur | draw | A Dinosaur draw card.
- 1 Earthshaker Dreadmaw | draw | A Dinosaur draw card.
- 1 Garruk's Uprising | draw | A draw card for the creature-heavy plan.
- 1 Guardian Project | draw | A draw card for the creature-heavy plan.
- 1 Harmonize | draw | A straightforward draw spell.
- 1 Imposing Vantasaur | draw | A Dinosaur draw card.
- 1 Return of the Wildspeaker | draw | A draw card for the creature-heavy plan.
- 1 Ripjaw Raptor | draw | A Dinosaur draw card.
- 1 Akroma's Will | interaction | An interaction spell for protecting the attack plan.
- 1 Boros Charm | interaction | An interaction spell for the deck.
- 1 Heroic Intervention | interaction | An interaction spell for the deck.
- 1 Lightning Greaves | interaction | An interaction artifact for supporting key creatures.
- 1 Swiftfoot Boots | interaction | An interaction artifact for supporting key creatures.
- 1 Whispersilk Cloak | interaction | An interaction artifact for supporting key creatures.
- 1 Apex Altisaur | removal | A Dinosaur removal creature.
- 1 Burning Sun's Avatar | removal | A Dinosaur removal creature.
- 1 Itzquinth, Firstborn of Gishath | removal | A Dinosaur removal card tied to Gishath.
- 1 Ravenous Sailback | removal | A Dinosaur removal creature.
- 1 Savage Stomp | removal | A Dinosaur-themed removal spell.
- 1 Thrashing Brontodon | removal | A Dinosaur removal creature.
- 1 Tranquil Frillback | removal | A Dinosaur removal creature.
- 1 Zacama, Primal Calamity | removal | A large Dinosaur removal threat.
- 1 Commune with Dinosaurs | synergy | A Dinosaur synergy card.
- 1 Kinjalli's Caller | synergy | A Dinosaur synergy creature.
- 1 Otepec Huntmaster | synergy | A Dinosaur synergy creature.
- 1 Marauding Raptor | synergy | A Dinosaur synergy creature.
- 1 Dinosaur Stampede | synergy | A Dinosaur synergy spell.
- 1 Kinjalli's Sunwing | synergy | A Dinosaur synergy creature.
- 1 Hunting Velociraptor | synergy | A Dinosaur synergy creature.
- 1 Invasion of Ixalan // Belligerent Regisaur | synergy | A Dinosaur synergy card.
- 1 Invasion of Ikoria // Zilortha, Apex of Ikoria | synergy | A Dinosaur synergy card.
- 1 Huatli's Raptor | synergy | A Dinosaur synergy creature.
- 1 Raptor Companion | synergy | A Dinosaur synergy creature.
- 1 Raptor Hatchling | synergy | A Dinosaur synergy creature.
- 1 Territorial Hammerskull | synergy | A Dinosaur synergy creature.
- 1 Sunfrill Imitator | synergy | A Dinosaur synergy creature.
- 1 Ancient Brontodon | threat | A large Dinosaur threat.
- 1 Carnage Tyrant | threat | A Dinosaur threat for closing games.
- 1 Etali, Primal Conqueror // Etali, Primal Sickness | threat | A legendary Dinosaur threat.
- 1 Etali, Primal Storm | threat | A legendary Dinosaur threat.
- 1 Ghalta, Primal Hunger | threat | A legendary Dinosaur threat.
- 1 Ghalta, Stampede Tyrant | threat | A legendary Dinosaur threat.
- 1 Goring Ceratops | threat | A Dinosaur threat for combat.
- 1 Polyraptor | threat | A Dinosaur threat for combat.
- 1 Quartzwood Crasher | threat | A Dinosaur threat for combat.
- 1 Regisaur Alpha | threat | A Dinosaur threat for the tribal plan.
- 1 Zetalpa, Primal Dawn | threat | A legendary Dinosaur threat.
- 1 Tyrranax Rex | threat | A Dinosaur threat for closing games.
- 1 Austere Command | wipe | A wipe for resetting difficult boards.
- 1 Blasphemous Act | wipe | A wipe for resetting difficult boards.
- 1 Wakening Sun's Avatar | wipe | A Dinosaur wipe creature.

</details>

### 5. blink Commander, owned first

Format: Commander. Theme: blink. Pool: owned_first. Shortlist: 171 names.

Cards: 98. Repair turn: true. Block findings: 1.

**Summary:** Repaired Brago blink deck at bracket 3: 99 listed cards plus Brago, King Eternal, with 36 lands and the requested 10 ramp, 10 draw, 6 interaction, 8 removal, 3 wipes, 14 synergy cards, and 12 threats.

- NOTE: I could not place "Plains", so it is not in the deck.
- [BLOCK] `deck_size`: deck has 99 cards, the format needs exactly 100 (commander included)
- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [WARN] `not_owned`: Brago, King Eternal: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 2.83 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Abandoned Air Temple | land | Azorius source.
- 1 Adventurer's Inn | land | Mana base.
- 1 Ash Barrens | land | Fixing and basic-land cycling.
- 1 Bonders' Enclave | land | Late-game card advantage land.
- 1 Command Tower | land | Reliable commander-color fixing.
- 1 Escape Tunnel | land | Fixes mana and can make Brago evasive.
- 1 Evolving Wilds | land | Basic-land fixing.
- 1 Exotic Orchard | land | Multiplayer color fixing.
- 1 Fabled Passage | land | Basic-land fixing.
- 1 Field of Ruin | land | Answers problematic lands.
- 1 Fountainport | land | Flexible utility land.
- 1 Glacial Fortress | land | Azorius dual land.
- 1 Great Hall of the Citadel | land | White source and utility.
- 1 Irrigated Farmland | land | Azorius dual with cycling.
- 1 Minas Tirith | land | White source with card-draw utility.
- 1 Opal Palace | land | Commander-supporting mana source.
- 1 Path of Ancestry | land | Fixing for the creature-heavy build.
- 1 Plaza of Heroes | land | Protects Brago while fixing mana.
- 1 Port Town | land | Azorius dual land.
- 1 Prairie Stream | land | Azorius dual land.
- 1 Reliquary Tower | land | No maximum hand-size utility.
- 1 Rivendell | land | Blue source and creature utility.
- 1 Rogue's Passage | land | Helps Brago connect.
- 1 Scavenger Grounds | land | Graveyard disruption.
- 1 Secluded Courtyard | land | Creature-based color fixing.
- 1 Spire of Industry | land | Artifact-enabled color fixing.
- 1 Temple of the False God | land | Midgame mana acceleration.
- 1 Terramorphic Expanse | land | Basic-land fixing.
- 1 The Grey Havens | land | Blue source and utility.
- 1 Thriving Isle | land | Flexible blue fixing.
- 1 Tranquil Cove | land | Azorius fixing.
- 1 Unclaimed Territory | land | Creature-based color fixing.
- 1 Vibrant Cityscape | land | Flexible color fixing.
- 1 War Room | land | Repeatable card-draw land.
- 1 Windbrisk Heights | land | Value-oriented utility land.
- 1 Arcane Signet | ramp | Efficient color fixing.
- 1 Astral Cornucopia | ramp | Blinkable mana artifact.
- 1 Bender's Waterskin | ramp | Mana artifact that benefits from Brago resets.
- 1 Chromatic Lantern | ramp | Fixes the utility-heavy mana base.
- 1 Commander's Sphere | ramp | Mana early and cards later.
- 1 Fellwar Stone | ramp | Efficient multiplayer mana rock.
- 1 Sol Ring | ramp | Fast colorless acceleration.
- 1 Talisman of Progress | ramp | Efficient Azorius acceleration.
- 1 Thought Vessel | ramp | Mana rock with hand-size utility.
- 1 Wayfarer's Bauble | ramp | Permanent land-based acceleration.
- 1 Champions of Minas Tirith | draw | Creature-based card advantage.
- 1 Faramir, Steward of Gondor | draw | Creates a board while providing cards.
- 1 Idol of Oblivion | draw | Efficient repeatable card draw.
- 1 Mask of Memory | draw | Brago can carry it to turn combat into cards.
- 1 Pull from Tomorrow | draw | Scalable instant-speed refill.
- 1 Reconnaissance Mission | draw | Rewards evasive combat damage.
- 1 Skullclamp | draw | Strong creature-token card advantage.
- 1 Stone of Erech | draw | Card advantage with graveyard utility.
- 1 Thirst for Knowledge | draw | Efficient instant-speed selection.
- 1 Tome of Legends | draw | Brago keeps it supplied with counters.
- 1 Arcane Denial | interaction | Flexible stack interaction.
- 1 Clever Concealment | interaction | Protects the developed board from sweepers.
- 1 Dovin's Veto | interaction | Reliable protection against noncreature spells.
- 1 Lightning Greaves | interaction | Protects Brago and grants haste.
- 1 Reprieve | interaction | Tempo interaction that replaces itself.
- 1 Swiftfoot Boots | interaction | Protects Brago while preserving equip flexibility.
- 1 Banishing Light | removal | Versatile permanent-based answer.
- 1 Crib Swap | removal | Exiles troublesome creatures.
- 1 Fiend Hunter | removal | Blinkable creature removal.
- 1 Generous Gift | removal | Answers any permanent.
- 1 Get Lost | removal | Efficient broad permanent removal.
- 1 Journey to Nowhere | removal | Low-cost creature exile.
- 1 Swords to Plowshares | removal | Premium creature exile.
- 1 Ty Lee, Chi Blocker | removal | Creature-based disruptive removal.
- 1 Austere Command | wipe | Flexible asymmetrical reset.
- 1 Perplexing Test | wipe | Preserves much of the artifact board while resetting creatures.
- 1 Supreme Verdict | wipe | Dependable creature-board reset.
- 1 Contagion Clasp | synergy | Brago can repeatedly reset its enters-the-battlefield value.
- 1 Crystal Fragments // Summon: Alexander | synergy | A value permanent Brago can reset.
- 1 Ennis, Debate Moderator | synergy | Supports the deck's blink-value plan.
- 1 Flickerwisp | synergy | Additional blink effect and reusable disruption.
- 1 Gift of Immortality | synergy | Keeps an important value creature or Brago available.
- 1 Inspiring Overseer | synergy | Reusable enters-the-battlefield life and card value.
- 1 Ioreth of the Healing House | synergy | Untaps Brago or other key legends for additional value.
- 1 Lembas | synergy | Blinking it repeatedly converts into card selection.
- 1 Personify | synergy | Supports the deck's flicker-focused value plan.
- 1 Roll-Roll-Roll-Roll | synergy | A permanent whose chapter value can be reset.
- 1 Slip On the Ring | synergy | Instant-speed blink protection and value.
- 1 The Bath Song | synergy | Saga value that benefits from being reset.
- 1 Wall of Omens | synergy | Cheap repeatable blink card advantage.
- 1 Waterbender Ascension | synergy | Rewards the deck's incremental value engine.
- 1 Aang, the Last Airbender | threat | High-impact legendary threat with removal utility.
- 1 Angel of Condemnation | threat | Evasive threat and repeatable flicker effect.
- 1 Angel of Serenity | threat | Large evasive finisher with major enters-the-battlefield impact.
- 1 Bronze Guardian | threat | Growing artifact payoff that protects the mana engine.
- 1 Exemplar of Light | threat | Evasive Angel threat.
- 1 Giada, Font of Hope | threat | Accelerates and strengthens the Angel threat package.
- 1 Lake-town Mariners // Gone Fishing | threat | Evasive creature threat with additional value.
- 1 Meneldor, Swift Savior | threat | Evasive blink-oriented threat.
- 1 PuPu UFO | threat | Artifact creature that adds pressure while supporting the artifact shell.
- 1 S.H.I.E.L.D. Flying Car | threat | Vehicle threat that complements the artifact package.
- 1 Summon: Primal Garuda | threat | Saga creature that supplies recurring value and pressure.
- 1 Summon: Shiva | threat | Saga creature threat with value on resolution.

</details>

### 6. Modern tempo, tournament

Format: Modern. Theme: tempo. Pool: any_card. Shortlist: 165 names.

Cards: 60. Repair turn: false. Block findings: 0.

**Summary:** This blue-red tempo deck aims to establish creature pressure, use its instant-heavy suite to keep the opponent off balance, and close while that pressure remains in place. It gives up some card quality and late-game inevitability for a proactive plan built around efficient spells, removal, and disruption.

- [INFO] `curve_summary`: average mana value 2.22 over 36 nonland cards

<details><summary>The deck list</summary>

- 4 Steam Vents | land | Land slot for the blue-red mana base.
- 4 Shivan Reef | land | Land slot for the blue-red mana base.
- 4 Stormcarved Coast | land | Land slot for the blue-red mana base.
- 4 Sulfur Falls | land | Land slot for the blue-red mana base.
- 4 Scalding Tarn | land | Land slot for the mana base.
- 4 Sink into Stupor // Soporific Springs | land | Land slot that also fits the deck's spell-focused plan.
- 4 Consider | draw | Cheap draw selection for a proactive spell-heavy plan.
- 2 Preordain | draw | Additional draw selection to smooth the early game.
- 4 Counterspell | interaction | Core instant-speed interaction for protecting the tempo plan.
- 2 Spell Pierce | interaction | Low-cost interaction that helps maintain tempo.
- 4 Lightning Bolt | removal | Efficient removal for clearing resistance to the deck's pressure.
- 4 Pongify | removal | Additional cheap removal for problematic creatures.
- 4 Temporal Mastery | synergy | Synergy piece for the deck's temporal package.
- 4 Ragavan, Nimble Pilferer | ramp | Ramp creature that supports the deck's proactive board presence.
- 4 Storm-Kiln Artist | ramp | Ramp creature for the spell-focused portion of the plan.
- 4 Izzet Signet | ramp | Ramp support that helps develop the deck's resources.

</details>

### 7. Modern burn, casual

Format: Modern. Theme: burn. Pool: any_card. Shortlist: 289 names.

Cards: 60. Repair turn: false. Block findings: 0.

**Summary:** This deck takes a board-based mono-red burn approach: deploy threats early, use its synergy pieces to keep damage pressure building, and use burn removal to clear a path or finish the last stretch. It is strongest when it can stay proactive and make every turn advance the opponent toward defeat. In exchange, the focused main deck gives up some flexibility against strategies that demand specialized answers, making the sideboard important for adjusting after the first game.

- [INFO] `curve_summary`: average mana value 2.83 over 36 nonland cards

<details><summary>The deck list</summary>

- 4 Barbarian Ring | land | Forms part of the mono-red land base.
- 4 Blighted Gorge | land | Forms part of the mono-red land base.
- 4 Keldon Megaliths | land | Forms part of the mono-red land base.
- 4 Ramunap Ruins | land | Forms part of the mono-red land base.
- 4 Sunscorched Desert | land | Forms part of the mono-red land base.
- 4 Valakut, the Molten Pinnacle | land | Forms part of the mono-red land base.
- 4 Risk Factor | draw | Provides a substantial portion of the deck's card draw.
- 2 Browbeat | draw | Adds more card draw to keep the pressure flowing.
- 2 Ancestors' Aid | ramp | Supplies the deck's ramp slots.
- 4 Lightning Bolt | removal | Efficient removal supports the burn plan.
- 2 Lightning Strike | removal | Additional removal keeps opposing threats in check.
- 4 Eidolon of the Great Revel | synergy | A synergy piece for the deck's aggressive damage plan.
- 4 Thermo-Alchemist | synergy | Supports repeated burn-oriented pressure.
- 4 Ashcloud Phoenix | threat | A threat that helps maintain battlefield pressure.
- 4 Champion of the Path | threat | A threat that contributes to the aggressive plan.
- 2 Fuming Effigy | threat | A threat that adds to the deck's pressure package.
- 2 Teapot Slinger | threat | A threat for continuing the deck's offensive plan.
- 2 Tectonic Giant | threat | A larger threat for closing out longer games.

</details>

### 8. Modern lifegain, FNM

Format: Modern. Theme: lifegain. Pool: any_card. Shortlist: 299 names.

Cards: 60. Repair turn: false. Block findings: 0.

**Summary:** This deck builds around early lifegain synergies and follows them with a steady creature assault, using its larger Angels and Vampires to finish games once the battlefield is under control. It can draw into more action and clear away opposing pressure, but it gives up speed through a mana base built entirely from nonbasic options and a plan that needs its pieces to stay together.

- [INFO] `curve_summary`: average mana value 3.28 over 36 nonland cards

<details><summary>The deck list</summary>

- 4 Scoured Barrens | land | A core land slot in the white-black mana base.
- 4 Shambling Vent | land | A land slot that supports the deck's white-black base.
- 4 Restless Fortress | land | A land slot for the white-black mana base.
- 4 Kabira Crossroads | land | A land slot that supports the lifegain plan.
- 4 Zof Consumption // Zof Bloodbog | land | A flexible land slot for the black side of the mana base.
- 2 Orzhova, the Church of Deals | land | A white-black land slot for the deck's mana base.
- 2 Vault of the Archangel | land | A utility land slot in the mana base.
- 2 Orzhov Keyrune | ramp | Mana development support for the deck's higher-end cards.
- 4 Inspiring Overseer | draw | A creature-based source of card draw.
- 2 Dawn of Hope | draw | Ongoing card-draw support for the lifegain strategy.
- 2 Solitude | removal | Efficient removal support for problematic opposing cards.
- 4 Murderous Rider // Swift End | removal | Flexible removal that also contributes to the creature plan.
- 4 Soul Warden | synergy | An early lifegain synergy piece.
- 4 Ajani's Pridemate | synergy | A central payoff for the lifegain plan.
- 4 Attended Healer | threat | A midgame threat that fits the deck's lifegain theme.
- 4 Bloodbond Vampire | threat | A threat that rewards the deck's core strategy.
- 2 Archangel of Thune | threat | A high-impact threat for closing games.
- 2 Angel of Invention | threat | A resilient top-end threat.
- 2 Blood Baron of Vizkopa | threat | An additional hard-hitting threat for the midgame.

</details>

### 9. Standard midrange, FNM

Format: Standard. Theme: midrange. Pool: any_card. Shortlist: 162 names.

Cards: 60. Repair turn: false. Block findings: 0.

**Summary:** This black-green midrange deck establishes its mana, develops a creature-focused board, and uses removal plus protective support to keep pressure on the opponent. It wins by turning its substantial threat package into a finishing board once opposing pieces have been handled. In exchange, it takes a deliberate midrange approach, prioritizing trading resources and developing over maximum early speed or highly narrow main-deck answers.

- [INFO] `curve_summary`: average mana value 3.33 over 36 nonland cards

<details><summary>The deck list</summary>

- 1 Blooming Marsh | land | Part of the black-green land base.
- 1 Deathcap Glade | land | Part of the black-green land base.
- 1 Golgari Guildgate | land | Part of the black-green land base.
- 1 Jungle Hollow | land | Part of the black-green land base.
- 1 Overgrown Tomb | land | Part of the black-green land base.
- 1 Temple of Malady | land | Part of the black-green land base.
- 1 Underground Mortuary | land | Part of the black-green land base.
- 1 Wastewood Verge | land | Part of the black-green land base.
- 1 Fabled Passage | land | Flexible land-base slot.
- 1 Evolving Wilds | land | Flexible land-base slot.
- 1 Terramorphic Expanse | land | Flexible land-base slot.
- 1 Escape Tunnel | land | Flexible land-base slot.
- 1 Demolition Field | land | Utility land slot.
- 1 Conduit Pylons | land | Utility land slot.
- 1 Echoing Deeps | land | Utility land slot.
- 1 Hidden Grotto | land | Utility land slot.
- 1 Rogue's Passage | land | Utility land slot.
- 1 Secluded Courtyard | land | Creature-focused land-base slot.
- 1 Starting Town | land | Land-base slot for the midrange plan.
- 1 Adventurer's Inn | land | Land-base slot for the midrange plan.
- 1 Fountainport | land | Utility land slot.
- 1 Multiversal Passage | land | Flexible land-base slot.
- 1 Valgavoth's Lair | land | Land-base slot for the midrange plan.
- 1 Ojer Kaslem, Deepest Growth // Temple of Cultivation | land | Creature-land slot that supports the midrange plan.
- 1 Aftermath Analyst | ramp | Early ramp for the deck's heavier cards.
- 1 Archdruid's Charm | ramp | Additional ramp support.
- 1 Phyrexian Arena | draw | Repeatable card-draw slot.
- 1 Darkstar Augur | draw | Creature-based card-draw slot.
- 1 Midnight Reaper | draw | Creature-based card-draw slot.
- 1 Insatiable Avarice | draw | Dedicated card-draw spell.
- 1 Unholy Annex // Ritual Chamber | draw | Card-draw enchantment slot.
- 1 Vraska Joins Up | draw | Card-draw support for the long game.
- 1 Assassin's Trophy | removal | Flexible removal for problematic permanents.
- 1 Bitter Triumph | removal | Efficient removal slot.
- 1 Hero's Downfall | removal | Direct removal for opposing threats.
- 1 Murder | removal | Straightforward removal slot.
- 1 Withering Torment | removal | Instant-speed removal support.
- 1 Pick Your Poison | removal | Flexible removal option.
- 1 Blossoming Defense | synergy | Protective support for the creature plan.
- 1 Snakeskin Veil | synergy | Protective support for key creatures.
- 1 Royal Treatment | synergy | Protective support for the board.
- 1 Undying Malice | synergy | Supports the deck's creature-focused plan.
- 1 Not Dead After All | synergy | Supports the deck's creature-focused plan.
- 1 Overprotect | synergy | Protective support for important permanents.
- 1 Swiftfoot Boots | synergy | Equipment support for the threat package.
- 1 Meathook Massacre II | synergy | Interaction that supports the deck's board plan.
- 1 Vein Ripper | threat | Creature threat for closing the game.
- 1 Rottenmouth Viper | threat | Creature threat for the midgame.
- 1 Lord Skitter, Sewer King | threat | Creature threat that adds board presence.
- 1 Chomping Changeling | threat | Creature threat for the board.
- 1 Reclamation Sage | threat | Creature threat that fits the midrange body plan.
- 1 Thrashing Brontodon | threat | Creature threat for sustained pressure.
- 1 Scavenging Ooze | threat | Creature threat for the midgame.
- 1 Summon: Bahamut | threat | Creature threat for the top end.
- 1 Abyssal Harvester | threat | Large creature threat for stabilizing the board.
- 1 Bringer of the Last Gift | threat | Large creature threat for the late game.
- 1 Massacre Wurm | threat | Large creature threat for the late game.
- 1 Steel Hellkite | threat | Creature threat that broadens the top end.
- 1 Zenos yae Galvus // Shinryu, Transcendent Rival | threat | Legendary creature threat for finishing games.
- 1 Zodiark, Umbral God | threat | Legendary creature threat for the top end.

</details>

### 10. Standard aggro, tournament

Format: Standard. Theme: aggro. Pool: any_card. Shortlist: 292 names.

Cards: 60. Repair turn: false. Block findings: 0.

**Summary:** This red-white aggro deck aims to establish a board of threats, reinforce that pressure with Goblin Oriflamme, and use removal and interaction to keep opposing resistance from taking over the game. It wins by maintaining an attacking force while its draw cards help sustain that pressure. In return, it is built to stay proactive rather than to devote many slots to a broad range of late-game answers.

- [INFO] `curve_summary`: average mana value 2.61 over 36 nonland cards

<details><summary>The deck list</summary>

- 4 Sacred Foundry | land | Forms part of the deck's land base.
- 4 Inspiring Vantage | land | Forms part of the deck's land base.
- 4 Elegant Parlor | land | Forms part of the deck's land base.
- 4 Sundown Pass | land | Forms part of the deck's land base.
- 4 Sunbillow Verge | land | Forms part of the deck's land base.
- 4 Temple of Triumph | land | Forms part of the deck's land base.
- 4 Fugitive Codebreaker | draw | Provides draw while staying within the deck's aggressive card pool.
- 2 Reckless Lackey | draw | Adds further draw support for games that run longer.
- 4 Boros Charm | interaction | Supplies interaction to back up the deck's pressure.
- 2 Sheltered by Ghosts | interaction | Adds interaction that supports the main proactive plan.
- 4 Harsh Annotation | removal | Provides focused removal to clear opposing resistance.
- 4 Case of the Gateway Express | removal | Adds more removal so threats can keep applying pressure.
- 4 Goblin Oriflamme | synergy | Provides the deck's central synergy piece for its aggressive plan.
- 4 Redcap Gutter-Dweller | threat | Serves as a core threat for applying pressure.
- 4 Frilled Sparkshooter | threat | Provides another set of threats for the aggressive plan.
- 4 Dragonback Lancer | threat | Rounds out the threat base with additional proactive creatures.

</details>

### 11. Commander with a locked card

Format: Commander. Theme: sacrifice. Pool: any_card. Shortlist: 294 names.

Cards: 99. Repair turn: false. Block findings: 0.

**Summary:** Karlov leads a sacrifice-focused Orzhov deck that develops its board with ramp, uses draw and synergy pieces to keep the plan moving, and presses its advantage with a broad set of threats. It wins by turning that sacrifice plan into sustained pressure while removal, interaction, and wipes keep opposing boards from taking over. The tradeoff is that the deck is committed to assembling and maintaining its sacrifice-focused board rather than pursuing a single narrow finish.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 3.05 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Arguel's Blood Fast // Temple of Aclazotz | land | Included as a land slot from the shortlist.
- 1 Bloodstained Mire | land | Included as a land slot from the shortlist.
- 1 Bojuka Bog | land | Included as a land slot from the shortlist.
- 1 Castle Doom | land | Included as a land slot from the shortlist.
- 1 Command Tower | land | Included as a land slot from the shortlist.
- 1 Diamond Valley | land | Included as a land slot from the shortlist.
- 1 Dust Bowl | land | Included as a land slot from the shortlist.
- 1 Evolving Wilds | land | Included as a land slot from the shortlist.
- 1 Exotic Orchard | land | Included as a land slot from the shortlist.
- 1 Fountainport | land | Included as a land slot from the shortlist.
- 1 Great Hall of Starnheim | land | Included as a land slot from the shortlist.
- 1 High Market | land | Included as a land slot from the shortlist.
- 1 Hostile Hostel // Creeping Inn | land | Included as a land slot from the shortlist.
- 1 Ifnir Deadlands | land | Included as a land slot from the shortlist.
- 1 Lake of the Dead | land | Included as a land slot from the shortlist.
- 1 Lazotep Quarry | land | Included as a land slot from the shortlist.
- 1 Miren, the Moaning Well | land | Included as a land slot from the shortlist.
- 1 Myriad Landscape | land | Included as a land slot from the shortlist.
- 1 Path of Ancestry | land | Included as a land slot from the shortlist.
- 1 Phyrexia's Core | land | Included as a land slot from the shortlist.
- 1 Phyrexian Tower | land | Included as a land slot from the shortlist.
- 1 Polluted Delta | land | Included as a land slot from the shortlist.
- 1 Rath's Edge | land | Included as a land slot from the shortlist.
- 1 Reliquary Tower | land | Included as a land slot from the shortlist.
- 1 Rogue's Passage | land | Included as a land slot from the shortlist.
- 1 Scavenger Grounds | land | Included as a land slot from the shortlist.
- 1 Shefet Dunes | land | Included as a land slot from the shortlist.
- 1 Springjack Pasture | land | Included as a land slot from the shortlist.
- 1 Starlit Sanctum | land | Included as a land slot from the shortlist.
- 1 Susur Secundi, Void Altar | land | Included as a land slot from the shortlist.
- 1 Tarrian's Journal // The Tomb of Aclazotz | land | Included as a land slot from the shortlist.
- 1 Terramorphic Expanse | land | Included as a land slot from the shortlist.
- 1 The Gold Saucer | land | Included as a land slot from the shortlist.
- 1 Treasure Map // Treasure Cove | land | Included as a land slot from the shortlist.
- 1 Westvale Abbey // Ormendahl, Profane Prince | land | Included as a land slot from the shortlist.
- 1 Windswept Heath | land | Included as a land slot from the shortlist.
- 1 Ashnod's Altar | ramp | A shortlist ramp piece that fits the sacrifice plan.
- 1 Crowded Crypt | ramp | A shortlist ramp piece that fits the sacrifice plan.
- 1 Culling the Weak | ramp | A shortlist ramp piece that fits the sacrifice plan.
- 1 Fain, the Broker | ramp | A shortlist ramp piece that fits the sacrifice plan.
- 1 Pitiless Plunderer | ramp | A shortlist ramp piece that fits the sacrifice plan.
- 1 Phyrexian Altar | ramp | A shortlist ramp piece that fits the sacrifice plan.
- 1 Priest of Forgotten Gods | ramp | A shortlist ramp piece that fits the sacrifice plan.
- 1 Sifter of Skulls | ramp | A shortlist ramp piece that fits the sacrifice plan.
- 1 Skullport Merchant | ramp | A shortlist ramp piece that fits the sacrifice plan.
- 1 Sol Ring | ramp | Required by the requested deck and assigned to the ramp role.
- 1 Baron Bertram Graywater | draw | A shortlist draw card for keeping the sacrifice plan supplied.
- 1 Corrupted Conviction | draw | A shortlist draw card for keeping the sacrifice plan supplied.
- 1 Disciple of Bolas | draw | A shortlist draw card for keeping the sacrifice plan supplied.
- 1 Ecstatic Awakener // Awoken Demon | draw | A shortlist draw card for keeping the sacrifice plan supplied.
- 1 Jerren, Corrupted Bishop // Ormendahl, the Corrupter | draw | A shortlist draw card for keeping the sacrifice plan supplied.
- 1 Lord Skitter's Butcher | draw | A shortlist draw card for keeping the sacrifice plan supplied.
- 1 Relic Vial | draw | A shortlist draw card for keeping the sacrifice plan supplied.
- 1 Smothering Abomination | draw | A shortlist draw card for keeping the sacrifice plan supplied.
- 1 Vampiric Rites | draw | A shortlist draw card for keeping the sacrifice plan supplied.
- 1 Village Rites | draw | A shortlist draw card for keeping the sacrifice plan supplied.
- 1 Cartel Aristocrat | interaction | A shortlist interaction card that supports the sacrifice plan.
- 1 Fanatical Devotion | interaction | A shortlist interaction card that supports the sacrifice plan.
- 1 Gift of Doom | interaction | A shortlist interaction card that supports the sacrifice plan.
- 1 Nightmare Shepherd | interaction | A shortlist interaction card that supports the sacrifice plan.
- 1 Promise of Tomorrow | interaction | A shortlist interaction card that supports the sacrifice plan.
- 1 Spirit Bonds | interaction | A shortlist interaction card that supports the sacrifice plan.
- 1 Attrition | removal | A shortlist removal card for answering opposing pieces.
- 1 Ayli, Eternal Pilgrim | removal | A shortlist removal card for answering opposing pieces.
- 1 Bone Shards | removal | A shortlist removal card for answering opposing pieces.
- 1 Eaten Alive | removal | A shortlist removal card for answering opposing pieces.
- 1 Grave Pact | removal | A shortlist removal card for answering opposing pieces.
- 1 Stronghold Assassin | removal | A shortlist removal card for answering opposing pieces.
- 1 Teysa, Orzhov Scion | removal | A shortlist removal card for answering opposing pieces.
- 1 Yawgmoth, Thran Physician | removal | A shortlist removal card for answering opposing pieces.
- 1 Altar of Dementia | synergy | A shortlist synergy piece for the sacrifice plan.
- 1 Bartolomé del Presidio | synergy | A shortlist synergy piece for the sacrifice plan.
- 1 Bastion of Remembrance | synergy | A shortlist synergy piece for the sacrifice plan.
- 1 Carrion Feeder | synergy | A shortlist synergy piece for the sacrifice plan.
- 1 Elas il-Kor, Sadistic Pilgrim | synergy | A shortlist synergy piece for the sacrifice plan.
- 1 Fleshtaker | synergy | A shortlist synergy piece for the sacrifice plan.
- 1 Hidden Stockpile | synergy | A shortlist synergy piece for the sacrifice plan.
- 1 Open the Graves | synergy | A shortlist synergy piece for the sacrifice plan.
- 1 Pious Evangel // Wayward Disciple | synergy | A shortlist synergy piece for the sacrifice plan.
- 1 Spawning Pit | synergy | A shortlist synergy piece for the sacrifice plan.
- 1 Viscera Seer | synergy | A shortlist synergy piece for the sacrifice plan.
- 1 Witch's Oven | synergy | A shortlist synergy piece for the sacrifice plan.
- 1 Woe Strider | synergy | A shortlist synergy piece for the sacrifice plan.
- 1 Zulaport Cutthroat | synergy | A shortlist synergy piece for the sacrifice plan.
- 1 Abhorrent Overlord | threat | A shortlist threat that helps close games.
- 1 Felisa, Fang of Silverquill | threat | A shortlist threat that helps close games.
- 1 Ghoulcaller Gisa | threat | A shortlist threat that helps close games.
- 1 Kuldotha Forgemaster | threat | A shortlist threat that helps close games.
- 1 Liesa, Forgotten Archangel | threat | A shortlist threat that helps close games.
- 1 Liliana, Heretical Healer // Liliana, Defiant Necromancer | threat | A shortlist threat that helps close games.
- 1 Mondrak, Glory Dominus | threat | A shortlist threat that helps close games.
- 1 Ratadrabik of Urborg | threat | A shortlist threat that helps close games.
- 1 Razaketh, the Foulblooded | threat | A shortlist threat that helps close games.
- 1 Requiem Angel | threat | A shortlist threat that helps close games.
- 1 Sidisi, Undead Vizier | threat | A shortlist threat that helps close games.
- 1 Vindictive Vampire | threat | A shortlist threat that helps close games.
- 1 Elesh Norn // The Argent Etchings | wipe | A shortlist wipe for resetting crowded boards.
- 1 The Meathook Massacre | wipe | A shortlist wipe for resetting crowded boards.
- 1 Toxic Deluge | wipe | A shortlist wipe for resetting crowded boards.

</details>

### 12. Commander on a budget

Format: Commander. Theme: tokens. Pool: any_card. Shortlist: 295 names.

Cards: 98. Repair turn: true. Block findings: 1.

**Summary:** Repaired Adeline token deck: 99 listed cards plus Adeline, Resplendent Cathar, with the requested 36 lands and exact role targets. The build emphasizes inexpensive white token production, anthem effects, resilient card flow, and token-favored sweepers for a bracket 2 play pattern.

- [BLOCK] `deck_size`: deck has 99 cards, the format needs exactly 100 (commander included)
- [INFO] `bracket_prose_rules`: bracket 2 (Core, 8+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 3.34 over 62 nonland cards

<details><summary>The deck list</summary>

- 1 Abundant Countryside | land | Token-friendly mana source.
- 1 Adagia, Windswept Bastion | land | White-compatible land slot.
- 1 Big Apple, 3 a.m. | land | Utility land slot.
- 1 Blighted Steppe | land | Land that can later support the board.
- 1 Castle Ardenvale | land | Mana source with a token-making outlet.
- 1 Country Roads | land | White-compatible land slot.
- 1 Cradle of the Accursed | land | Utility land that can become a body.
- 1 Dalkovan Encampment | land | White-compatible land slot.
- 1 Dowsing Dagger // Lost Vale | land | Early utility that can become an explosive land.
- 1 Dragon-Cursed Halls | land | White-compatible land slot.
- 1 Dunes of the Dead | land | Land that leaves behind a token.
- 1 Emeria's Call // Emeria, Shattered Skyclave | land | Flexible land slot with a late-game token spell.
- 1 Forbidden Orchard | land | Fixing and utility land.
- 1 Fountainport | land | Utility land with token support.
- 1 Gargoyle Castle | land | Land that can turn into an evasive threat.
- 1 Gods' Eye, Gate to the Reikai | land | Produces a token when replaced.
- 1 Golden Guardian // Gold-Forge Garrison | land | Land-producing transform card with a token engine.
- 1 Hall of Tagsin | land | White-compatible land slot.
- 1 Heap Gate | land | Fixing land for the artifact package.
- 1 Jasmine Dragon Tea Shop | land | White-compatible land slot.
- 1 Kjeldoran Outpost | land | Repeatable token production from a land.
- 1 Lazotep Quarry | land | Utility land slot.
- 1 Legion's Landing // Adanto, the First Fort | land | Early token maker that becomes a repeatable token land.
- 1 Mirrex | land | Land that produces a useful token.
- 1 Mirrorpool | land | Utility land for copying key spells.
- 1 Ojer Taq, Deepest Foundation // Temple of Civilization | land | Token multiplier with a land fallback.
- 1 Secluded Starforge | land | White-compatible land slot.
- 1 Spawning Bed | land | Land with token-making utility.
- 1 Springjack Pasture | land | Token-centric utility land.
- 1 Stardew Valley | land | White-compatible land slot.
- 1 The Gold Saucer | land | Utility land slot.
- 1 The Great Mound | land | White-compatible land slot.
- 1 Thousand Moons Smithy // Barracks of the Thousand | land | Token payoff that supplies a land side.
- 1 Urza's Factory | land | Late-game token-producing land.
- 1 Volatile Fault | land | Utility land slot.
- 1 Voldaren Estate | land | White-compatible land slot.
- 1 Battle Angels of Tyr | ramp | Can generate Treasure against ahead opponents.
- 1 Bucknard's Everfull Purse | ramp | Low-cost artifact mana option.
- 1 Coin of Mastery | ramp | Artifact acceleration.
- 1 Collector's Vault | ramp | Mana source with hand filtering.
- 1 Currency Converter | ramp | Cheap artifact-based resource engine.
- 1 Druidic Satchel | ramp | Incremental mana development and value.
- 1 Fishing Gear | ramp | Equipment-based Treasure production.
- 1 Goldvein Pick | ramp | Turns combat damage into Treasure.
- 1 Karn, Living Legacy | ramp | Creates Powerstone mana while building value.
- 1 Keeper of the Accord | ramp | Helps catch up on lands while producing bodies.
- 1 Bennie Bracks, Zoologist | draw | Reliable cards from Adeline and the token makers.
- 1 Bygone Bishop | draw | Creates Clues as the deck develops its board.
- 1 Caretaker's Talent | draw | Strong recurring card advantage for token production.
- 1 Chivalric Alliance | draw | Turns spare mana into cards and bodies.
- 1 Court of Grace | draw | Monarch-based card flow with token production.
- 1 Dawn of Hope | draw | Mana sink that converts life gain into cards and tokens.
- 1 Faramir, Field Commander | draw | Rewards attacking with cards and extra bodies.
- 1 Idol of Oblivion | draw | Efficient repeatable draw in a token deck.
- 1 Staff of the Storyteller | draw | Produces a token and converts token creation into cards.
- 1 Wedding Announcement // Wedding Festivity | draw | Provides steady tokens, cards, and a team buff.
- 1 Aerial Assault | removal | Efficient creature removal that rewards a wide board.
- 1 Banishing Slash | removal | Flexible permanent removal.
- 1 Battle Menu | removal | Versatile removal spell.
- 1 Citizen's Crowbar | removal | Token-friendly answer to artifacts and enchantments.
- 1 Hanged Executioner | removal | Token-producing creature with a built-in answer.
- 1 Kellan's Lightblades | removal | Efficient instant-speed creature answer.
- 1 Path to Redemption | removal | Permanent-based removal that fits the board plan.
- 1 Trostani's Judgment | removal | Exile removal with populate upside.
- 1 Basri Ket | interaction | Protects and grows the board while supplying tokens.
- 1 Lena, Selfless Champion | interaction | Protects a developed creature board.
- 1 Parting Gust | interaction | Flexible instant-speed defensive interaction.
- 1 Rootborn Defenses | interaction | Protects tokens and adds another copy of the best token.
- 1 Spirit Bonds | interaction | Creates Spirits and protects important creatures.
- 1 Summon: Knights of Round | interaction | Board-impacting defensive and offensive interaction.
- 1 Elspeth, Sun's Champion | wipe | Reset button that also rebuilds with tokens.
- 1 Hour of Reckoning | wipe | Token-friendly board reset.
- 1 The Battle of Bywater | wipe | Low-cost sweep that favors the smaller token board.
- 1 Animation Module | synergy | Converts counter placement into additional bodies.
- 1 Automated Assembly Line | synergy | Artifact token engine for a wide-board plan.
- 1 Cat Collector | synergy | Steady token production tied to the deck's creature plan.
- 1 Charismatic Conqueror | synergy | Builds an army as opponents develop.
- 1 Clarion Spirit | synergy | Rewards multiple spells with flying tokens.
- 1 Divine Visitation | synergy | Upgrades small tokens into powerful evasive threats.
- 1 Felidar Retreat | synergy | Land drops make bodies or strengthen the army.
- 1 Horn of Gondor | synergy | Scalable Human-token production alongside Adeline.
- 1 Intangible Virtue | synergy | Efficient anthem and vigilance for tokens.
- 1 Oketra's Monument | synergy | Reduces creature costs and makes Warrior tokens.
- 1 Rosie Cotton of South Lane | synergy | Turns repeated token creation into a growing threat.
- 1 Skrelv's Hive | synergy | Early recurring token source.
- 1 Twilight Drover | synergy | Transforms departing tokens into a larger board.
- 1 Ajani's Chosen | threat | Turns enchantment development into Cat tokens.
- 1 Archangel Elspeth | threat | Makes evasive bodies and pressures opponents quickly.
- 1 Archon of Sun's Grace | threat | Enchantment payoffs create a growing Pegasus force.
- 1 Attended Healer | threat | Life gain produces a continuing token presence.
- 1 Defiler of Faith | threat | Discounts white development and adds board presence.
- 1 Emeria Angel | threat | Land drops create evasive attackers.
- 1 Hero of Bladehold | threat | A potent attacking token-maker and anthem effect.
- 1 Illustrious Wanderglyph | threat | Anthem-style artifact creature that supports a wide board.
- 1 Mite Overseer | threat | Produces and enhances Phyrexian tokens.
- 1 Oketra the True | threat | Resilient attacker that supplies Warrior tokens.
- 1 Requiem Angel | threat | Turns creature losses into a flying token army.
- 1 Warren Warleader | threat | Creates Rabbits and converts a wide board into pressure.

</details>

