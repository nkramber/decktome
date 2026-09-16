# PR-14A bracket gate

Run date: 2026-09-16. Card snapshot: 2026-09-04.

Verdict: FAIL. 6 of 6 decks passed every block check, 4 sat in every band, 6 held no content violation, and the judge agreed with the bracket on 5 of 6 (83 percent, the bar is 80).

## Summary

| Measure | Value |
|---|---|
| Prompts | 6 |
| Decks returned | 6 |
| Decks with no block finding | 6 |
| Decks in every band | 4 |
| Decks with no content violation | 6 |
| Decks the endpoint did not check | 0 |
| Decks that needed a repair turn | 0 |
| Decks judged | 6 |
| Judge agreed with the bracket | 5 |
| Judge errors | 0 |
| Errors | 0 |
| Prompt version | 15 |
| Calls | 12 |
| Cost | $0.5318 |
| Time | 489 seconds |

## Run

- Suite `bracket`, run `pr14a-bracket-gate-run8`, on 2026-09-16, commit `6ff0244`.
- Partial run over `10,11,12,13,14,15`. It reads the bars of its own items, and it never stands for the suite.
- Roles: generate on `gpt-5.6-terra` (openai, effort medium), judge on `claude-opus-5` (anthropic, effort medium), repair on `gpt-5.6-terra` (openai, effort medium).
- Versions: card snapshot 2026-09-04, generate prompt version 15.
- Calls: 12. Cost: $0.5318. Time: 489 seconds.

## Off-band features by key

Each row counts the decks whose feature sat outside its band after the repair turns.

| Feature | Decks off band |
|---|---|
| `avg_mana_value` | 1 |
| `hands_two_to_four_lands` | 1 |
| `color_sources` | 1 |

## Decks

### 10. Bracket 4, Urza, Lord High Artificer, artifacts

Pool 320. 99 cards, 0 block findings, repaired false.

Judge: bracket 4, agrees. Nine Game Changers plus Sol Ring, Mox Opal/Amber, Lotus Petal and Metalworker give it explosive turn-two-to-three starts, and Urza with Krark-Clan Ironworks/Lion's Eye Diamond into Thassa's Oracle is a real infinite-mana kill backed by Force of Will and Fierce Guardianship. That speed and combo density blow past Bracket 3's limits. It stops short of cEDH, though: 29 basic Islands with no real fixing or fast-mana rituals beyond artifacts, and filler beaters like Darksteel Juggernaut, Skysovereign and Transmogrifying Wand mark it as an optimized artifact theme rather than a tuned competitive list.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 34 | 31 to 36 |  |  |
| `tapped_land` | 0 | 0 to 5 |  |  |
| `colorless_land` | 5 | 0 to 5 |  | Academy Ruins, Buried Ruin, Inventors' Fair, Mishra's Workshop, Urza's Saga |
| `color_sources` | 1.26 | 0.9 or more |  | sources of requirement: U 32.75 of 26 |
| `avg_mana_value` | 2.95 | 1.6 to 3 |  | over 65 nonland cards |
| `ramp` | 14 | 10 to 16 |  |  |
| `draw` | 11 | 8 to 16 |  |  |
| `removal` | 10 | 6 to 14 |  |  |
| `wipe` | 2 | 0 to 5 |  |  |
| `interaction` | 11 | 6 to 16 |  |  |
| `tutor` | 3 | 2 or more |  | Tezzeret, Artifice Master, Transmute Artifact, Whir of Invention |
| `fast_mana` | 11 | 3 or more |  | An Offer You Can't Refuse, Chrome Mox, Lion's Eye Diamond, Lotus Petal, Mana Vault, Moonsnare Prototype, Mox Amber, Mox Diamond, Mox Opal, Sol Ring, Springleaf Drum |
| `game_changer` | 9 | 4 or more |  | Chrome Mox, Fierce Guardianship, Force of Will, Lion's Eye Diamond, Mana Vault, Mishra's Workshop, Mox Diamond, Rhystic Study, Thassa's Oracle |
| `finisher` | 3 | 2 or more |  | Mechanized Production, Mirrodin Besieged, Thassa's Oracle |
| `mana_turn_four` | 5.31 | 4.4 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.72 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.89 | -9 to 0.5 |  | the commander costs 4 and comes down on turn 3.11 on average |

Goldfish over 10000 hands: the commander on turn 3.11, 5.31 mana on turn four, and 0.72 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Chrome Mox, Fierce Guardianship, Force of Will, Lion's Eye Diamond, Mana Vault, Mishra's Workshop, Mox Diamond, Rhystic Study, Thassa's Oracle. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 2.95 over 65 nonland cards
- INFO `mana_pass`: the builder moved 2 cards of the mana base to bring the deck inside its power level

Cards:

- 29 Island
- 1 Academy Ruins
- 1 Buried Ruin
- 1 Inventors' Fair
- 1 Mishra's Workshop
- 1 Urza's Saga
- 1 Rhystic Study
- 1 Thoughtcast
- 1 Sai, Master Thopterist
- 1 Vedalken Archmage
- 1 Riddlesmith
- 1 Reverse Engineer
- 1 Thirst for Knowledge
- 1 Forensic Gadgeteer
- 1 Nexus of Becoming
- 1 Tezzeret, Artifice Master
- 1 Trading Post
- 1 Force of Will
- 1 Fierce Guardianship
- 1 An Offer You Can't Refuse
- 1 Metallic Rebuke
- 1 Disruption Protocol
- 1 Ice Out
- 1 Ghostly Flicker
- 1 Welding Jar
- 1 Etched Champion
- 1 Padeem, Consul of Innovation
- 1 Jhoira's Toolbox
- 1 Chrome Mox
- 1 Lion's Eye Diamond
- 1 Mana Vault
- 1 Mox Diamond
- 1 Lotus Petal
- 1 Mox Opal
- 1 Sol Ring
- 1 Moonsnare Prototype
- 1 Springleaf Drum
- 1 Metalworker
- 1 Krark-Clan Ironworks
- 1 Automated Artificer
- 1 Vedalken Engineer
- 1 Aether Spellbomb
- 1 Blasting Station
- 1 Contagion Clasp
- 1 Cyber Conversion
- 1 Into Thin Air
- 1 Kitesail Larcenist
- 1 Resculpt
- 1 Ravenform
- 1 Skysovereign, Consul Flagship
- 1 Transmogrifying Wand
- 1 Engineered Explosives
- 1 Hurkyl's Recall
- 1 Whir of Invention
- 1 Transmute Artifact
- 1 Mystic Forge
- 1 Kappa Cannoneer
- 1 Cyberdrive Awakener
- 1 Myr Enforcer
- 1 Sojourner's Companion
- 1 Master Transmuter
- 1 Phyrexian Metamorph
- 1 Traxos, Scourge of Kroog
- 1 Darksteel Juggernaut
- 1 Lodestone Golem
- 1 Karn, Scion of Urza
- 1 Arcbound Reclaimer
- 1 Thassa's Oracle
- 1 Mechanized Production
- 1 Mirrodin Besieged
- 1 Mox Amber

### 11. Bracket 4, Korvold, Fae-Cursed King, treasure sacrifice

Pool 328. 99 cards, 0 block findings, repaired false.

Judge: bracket 4, agrees. Six marked Game Changers (Lion's Eye Diamond, Mana Vault, Mox Diamond, Chrome Mox, Jeska's Will, Orcish Bowmasters) alongside Sol Ring, Lotus Petal, Dark Ritual and original dual lands put this well past Bracket 3's three-GC cap and give it genuine turn-three or four explosiveness. The aristocrats shell has real engine combos (Yawgmoth plus Nightmare Shepherd/undying-style bodies, Goblin Bombardment with Prossh tokens, Pitiless Plunderer loops) plus efficient interaction and two sweepers. Still, the creature suite leans on flavorful sacrifice value pieces rather than a tuned cEDH win line, so it sits at Optimized rather than 5.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 33 | 31 to 36 |  |  |
| `tapped_land` | 0 | 0 to 5 |  |  |
| `colorless_land` | 0 | 0 to 5 |  |  |
| `color_sources` | 1.16 | 0.9 or more |  | sources of requirement: B 23 of 19, R 22 of 19, G 22 of 19 |
| `avg_mana_value` | 2.88 | 1.6 to 3 |  | over 66 nonland cards |
| `ramp` | 13 | 10 to 16 |  |  |
| `draw` | 12 | 8 to 16 |  |  |
| `removal` | 10 | 6 to 14 |  |  |
| `wipe` | 2 | 0 to 5 |  |  |
| `interaction` | 11 | 6 to 16 |  |  |
| `tutor` | 2 | 2 or more |  | Diabolic Intent, Magda, Brazen Outlaw |
| `fast_mana` | 10 | 3 or more |  | Chrome Mox, Dark Ritual, Lion's Eye Diamond, Lotus Petal, Mana Vault, Mox Amber, Mox Diamond, Rite of Flame, Sol Ring, Springleaf Drum |
| `game_changer` | 6 | 4 or more |  | Chrome Mox, Jeska's Will, Lion's Eye Diamond, Mana Vault, Mox Diamond, Orcish Bowmasters |
| `finisher` | 5 | 2 or more |  | Prossh, Skyraider of Kher, Revel in Riches, The Meathook Massacre, Vraska, Golgari Queen, Ziatora, the Incinerator |
| `mana_turn_four` | 4.91 | 4.4 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.71 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.58 | -9 to 0.5 |  | the commander costs 5 and comes down on turn 4.42 on average |

Goldfish over 10000 hands: the commander on turn 4.42, 4.91 mana on turn four, and 0.71 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Chrome Mox, Jeska's Will, Lion's Eye Diamond, Mana Vault, Mox Diamond, Orcish Bowmasters. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 2.88 over 66 nonland cards
- INFO `mana_pass`: the builder moved 2 cards of the mana base to bring the deck inside its power level

Cards:

- 1 Badlands
- 1 Bayou
- 1 Blood Crypt
- 1 City of Brass
- 1 Command Tower
- 1 Blightstep Pathway // Searstep Pathway
- 1 Cragcrown Pathway // Timbercrown Pathway
- 1 Darkbore Pathway // Slitherbore Pathway
- 1 Forbidden Orchard
- 1 Grove of the Burnwillows
- 1 Karplusan Forest
- 1 Llanowar Wastes
- 1 Mana Confluence
- 1 Overgrown Tomb
- 1 Phyrexian Tower
- 1 Stomping Ground
- 1 Sulfurous Springs
- 1 Taiga
- 1 Tarnished Citadel
- 4 Swamp
- 4 Forest
- 4 Mountain
- 1 Chrome Mox
- 1 Dark Ritual
- 1 Jeska's Will
- 1 Lion's Eye Diamond
- 1 Lotus Petal
- 1 Magda, Brazen Outlaw
- 1 Mana Vault
- 1 Mox Amber
- 1 Mox Diamond
- 1 Pitiless Plunderer
- 1 Rite of Flame
- 1 Sol Ring
- 1 Springleaf Drum
- 1 Ahriman
- 1 Beetle-Headed Merchants
- 1 Bushmeat Poacher
- 1 Dina, Essence Brewer
- 1 Gas Guzzler
- 1 Ice Cream Kitty
- 1 Lord Skitter's Butcher
- 1 Sanguine Spy
- 1 Skyclave Shadowcat
- 1 Smothering Abomination
- 1 Stormclaw Rager
- 1 Vampire Gourmand
- 1 Constant Mists
- 1 Cultist of the Absolute
- 1 Dark Privilege
- 1 Eldrazi Monument
- 1 Gift of Doom
- 1 Goblin Chirurgeon
- 1 Gyome, Master Chef
- 1 My Precious // Allure of Power
- 1 Nightmare Shepherd
- 1 Slobad, Goblin Tinkerer
- 1 Sylvan Safekeeper
- 1 Attrition
- 1 Bone Shards
- 1 Broadside Bombardiers
- 1 Eaten Alive
- 1 Goblin Bombardment
- 1 Grist, the Hunger Tide
- 1 Minsc & Boo, Timeless Heroes
- 1 Orcish Bowmasters
- 1 Skyfisher Spider
- 1 Yawgmoth, Thran Physician
- 1 Academy Manufactor
- 1 Diabolic Intent
- 1 Xorn
- 1 Blood Host
- 1 Corpse Harvester
- 1 Flesh-Eater Imp
- 1 Gobbling Ooze
- 1 Hell's Caretaker
- 1 Kethek, Crucible Goliath
- 1 Meren of Clan Nel Toth
- 1 Prossh, Skyraider of Kher
- 1 Redcap Gutter-Dweller
- 1 Sadistic Hypnotist
- 1 Thallid Omnivore
- 1 Undercity Scavenger
- 1 Revel in Riches
- 1 Vraska, Golgari Queen
- 1 Ziatora, the Incinerator
- 1 Spontaneous Combustion
- 1 The Meathook Massacre
- 1 Mount Doom
- 1 Nurturing Peatland

### 12. Bracket 4, Prosper, Tome-Bound, treasure

Pool 212. 99 cards, 0 block findings, repaired false.

Judge: bracket 5, disagrees. This is a cEDH-caliber Prosper/treasure-storm build: dual lands, fetches, and ten-plus Game Changers including Chrome Mox, Mox Diamond, Mana Vault, Lion's Eye Diamond, Ad Nauseam, Underworld Breach, and a full tutor suite (Demonic, Vampiric, Imperial Seal, Gamble). Underworld Breach plus LED with Storm-Kiln Artist/Wheel effects is a fast deterministic kill, backed by free interaction like Deadly Rollick, Snuff Out, and Pyroblast/REB. The card choices are metagame-driven rather than thematic, placing it squarely at the top of the ladder.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 33 | 31 to 36 |  |  |
| `tapped_land` | 1 | 0 to 5 |  | Foreboding Ruins |
| `colorless_land` | 0 | 0 to 5 |  |  |
| `color_sources` | 1.61 | 0.9 or more |  | sources of requirement: B 31.5 of 19, R 30.5 of 19 |
| `avg_mana_value` | 2.38 | 1.6 to 3 |  | over 66 nonland cards |
| `ramp` | 13 | 10 to 16 |  |  |
| `draw` | 12 | 8 to 16 |  |  |
| `removal` | 10 | 6 to 14 |  |  |
| `wipe` | 2 | 0 to 5 |  |  |
| `interaction` | 11 | 6 to 16 |  |  |
| `tutor` | 5 | 2 or more |  | Demonic Tutor, Gamble, Imperial Seal, Magda, Brazen Outlaw, Vampiric Tutor |
| `fast_mana` | 6 | 3 or more |  | Chrome Mox, Lion's Eye Diamond, Lotus Petal, Mana Vault, Mox Diamond, Sol Ring |
| `game_changer` | 11 | 4 or more |  | Ad Nauseam, Chrome Mox, Demonic Tutor, Gamble, Imperial Seal, Lion's Eye Diamond, Mana Vault, Mox Diamond, Orcish Bowmasters, Underworld Breach, Vampiric Tutor |
| `finisher` | 4 | 2 or more |  | Revel in Riches, Smaug the Impenetrable, Smaug, Wicked Worm, Swashbuckler Extraordinaire |
| `mana_turn_four` | 4.68 | 4.4 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.71 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.37 | -9 to 0.5 |  | the commander costs 4 and comes down on turn 3.63 on average |

Goldfish over 10000 hands: the commander on turn 3.63, 4.68 mana on turn four, and 0.71 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Ad Nauseam, Chrome Mox, Demonic Tutor, Gamble, Imperial Seal, Lion's Eye Diamond, Mana Vault, Mox Diamond, Orcish Bowmasters, Underworld Breach, Vampiric Tutor. Mass land denial none. Extra turns none. Combos 2.
- Prosper, Tome-Bound + Professional Face-Breaker + Sensei's Divining Top + Xorn (speed 5)
- Underworld Breach + Lion's Eye Diamond + Wheel of Fortune (speed 5)

Findings:

- INFO `curve_summary`: average mana value 2.38 over 66 nonland cards
- INFO `mana_pass`: the builder moved 3 cards of the mana base to bring the deck inside its power level

Cards:

- 1 Badlands
- 1 Blackcleave Cliffs
- 1 Blazemire Verge
- 1 Blightstep Pathway // Searstep Pathway
- 1 Blood Crypt
- 1 Bloodstained Mire
- 1 Cavern of Souls
- 1 City of Brass
- 1 Command Tower
- 1 Dragonskull Summit
- 1 Exotic Orchard
- 1 Fabled Passage
- 1 Forbidden Orchard
- 1 Foreboding Ruins
- 1 Graven Cairns
- 1 Haunted Ridge
- 1 Luxury Suite
- 1 Mana Confluence
- 1 Prismatic Vista
- 1 Spire of Industry
- 1 Sulfurous Springs
- 1 Tarnished Citadel
- 1 Tainted Peak
- 3 Mountain
- 4 Swamp
- 1 Black Market Connections
- 1 Chrome Mox
- 1 Goldspan Dragon
- 1 Lion's Eye Diamond
- 1 Lotus Petal
- 1 Magda, Brazen Outlaw
- 1 Mana Vault
- 1 Mox Diamond
- 1 Pitiless Plunderer
- 1 Professional Face-Breaker
- 1 Ragavan, Nimble Pilferer
- 1 Sol Ring
- 1 Storm-Kiln Artist
- 1 Ad Nauseam
- 1 Corrupted Conviction
- 1 Demand Answers
- 1 Faithless Looting
- 1 Mishra's Bauble
- 1 Night's Whisper
- 1 Read the Bones
- 1 Reckless Lackey
- 1 Sensei's Divining Top
- 1 Skullclamp
- 1 Village Rites
- 1 Wheel of Fortune
- 1 Abrade
- 1 Bedevil
- 1 Chaos Warp
- 1 Deadly Rollick
- 1 Feed the Swarm
- 1 Go for the Throat
- 1 Infernal Grasp
- 1 Orcish Bowmasters
- 1 Rakdos Charm
- 1 Snuff Out
- 1 Hexing Squelcher
- 1 Kaya's Ghostform
- 1 Lightning Greaves
- 1 Mithril Coat
- 1 Not Dead After All
- 1 Pyroblast
- 1 Red Elemental Blast
- 1 Swiftfoot Boots
- 1 Tibalt's Trickery
- 1 Undying Malice
- 1 Vexing Bauble
- 1 Academy Manufactor
- 1 Contract Hero
- 1 Xorn
- 1 Casey & Raph, Hotheads
- 1 Flamekin Gildweaver
- 1 Jaded Sell-Sword
- 1 Marut
- 1 Meticulous Artisan
- 1 Namazu Trader
- 1 Smaug the Impenetrable
- 1 Revel in Riches
- 1 Smaug, Wicked Worm
- 1 Swashbuckler Extraordinaire
- 1 Toxic Deluge
- 1 Vandalblast
- 1 Demonic Tutor
- 1 Gamble
- 1 Imperial Seal
- 1 Underworld Breach
- 1 Vampiric Tutor
- 1 Mount Doom
- 1 Smoldering Marsh
- 1 Starting Town

### 13. Bracket 5, Kinnan, Bonder Prodigy, combo

Pool 323. 99 cards, 0 block findings, repaired false.

Judge: bracket 5, agrees. This is a tuned Kinnan combo deck: heavy fast mana (Chrome Mox, Mox Diamond/Opal/Amber, Lion's Eye Diamond, Grim Monolith, Basalt Monolith, Lotus Petal), an optimized dual/fetch mana base, and a dozen-plus Game Changers alongside free counterspells like Fierce Guardianship and Force of Will. It packs multiple two-card infinite mana/untap loops with Kinnan (Basalt Monolith, Devoted Druid, Voltaic Construct, Filigree Sages, Dramatic Reversal) plus Thassa's Oracle and Protean Hulk as compact win conditions. There's no theme here beyond winning as fast as possible with protection, which is textbook cEDH.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 29 | 27 to 33 |  |  |
| `tapped_land` | 1 | 0 to 2 |  | Vineglimmer Snarl |
| `colorless_land` | 1 | 0 to 6 |  | Deserted Temple |
| `color_sources` | 1.28 | 0.9 or more |  | sources of requirement: U 29.5 of 23, G 26.75 of 19 |
| `avg_mana_value` | 2.53 | 1.2 to 2.4 | yes | over 70 nonland cards |
| `ramp` | 17 | 12 to 22 |  |  |
| `draw` | 16 | 8 to 18 |  |  |
| `removal` | 9 | 4 to 14 |  |  |
| `wipe` | 1 | 0 to 3 |  |  |
| `interaction` | 15 | 8 to 22 |  |  |
| `tutor` | 4 | 4 or more |  | Kuldotha Forgemaster, Muddle the Mixture, Protean Hulk, Tezzeret, Cruel Captain |
| `fast_mana` | 10 | 6 or more |  | An Offer You Can't Refuse, Chrome Mox, Lion's Eye Diamond, Lotus Petal, Mana Vault, Mox Amber, Mox Diamond, Mox Opal, Sol Ring, Springleaf Drum |
| `game_changer` | 11 | 8 or more |  | Chrome Mox, Cyclonic Rift, Fierce Guardianship, Force of Will, Grim Monolith, Lion's Eye Diamond, Mana Vault, Mox Diamond, Rhystic Study, Seedborn Muse, Thassa's Oracle |
| `finisher` | 2 | 1 or more |  | Dread Linnorm // Scale Deflection, Thassa's Oracle |
| `mana_turn_four` | 5.06 | 4.6 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.64 | 0.65 or more | yes | share of first seven-card hands |
| `commander_turn_over_mv` | -0.39 | -9 to 0.3 |  | the commander costs 2 and comes down on turn 1.61 on average |

Goldfish over 10000 hands: the commander on turn 1.61, 5.06 mana on turn four, and 0.64 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Chrome Mox, Cyclonic Rift, Fierce Guardianship, Force of Will, Grim Monolith, Lion's Eye Diamond, Mana Vault, Mox Diamond, Rhystic Study, Seedborn Muse, Thassa's Oracle. Mass land denial none. Extra turns none. Combos 1.
- Kinnan, Bonder Prodigy + Basalt Monolith (speed 5, two cards)

Findings:

- INFO `curve_summary`: average mana value 2.53 over 70 nonland cards
- WARN `profile_off_band`: the average mana value of the nonland cards is 2.53, and bracket 5 wants 1.2 to 2.4
- WARN `profile_off_band`: the share of opening hands with two to four lands is 0.64, and bracket 5 wants 0.65 or more (share of first seven-card hands)
- INFO `mana_pass`: the builder moved 5 cards of the mana base to bring the deck inside its power level

Cards:

- 1 Breeding Pool
- 1 Tropical Island
- 1 Command Tower
- 1 City of Brass
- 1 Mana Confluence
- 1 Forbidden Orchard
- 1 Rejuvenating Springs
- 1 Dreamroot Cascade
- 1 Hinterland Harbor
- 1 Yavimaya Coast
- 1 Misty Rainforest
- 1 Polluted Delta
- 1 Scalding Tarn
- 1 Verdant Catacombs
- 1 Windswept Heath
- 1 Wooded Foothills
- 1 Prismatic Vista
- 1 Fabled Passage
- 4 Island
- 3 Forest
- 1 Minamo, School at Water's Edge
- 1 Deserted Temple
- 1 Chrome Mox
- 1 Lion's Eye Diamond
- 1 Mana Vault
- 1 Mox Diamond
- 1 Lotus Petal
- 1 Mox Amber
- 1 Mox Opal
- 1 Sol Ring
- 1 Springleaf Drum
- 1 Grim Monolith
- 1 Basalt Monolith
- 1 Arbor Elf
- 1 Devoted Druid
- 1 Earthcraft
- 1 Kiora's Follower
- 1 Gwenna, Eyes of Gaea
- 1 Lost Jitte
- 1 Rhystic Study
- 1 Cloud of Faeries
- 1 Frantic Search
- 1 Lórien Revealed
- 1 Benefactor's Draught
- 1 Cerulean Wisps
- 1 Refocus
- 1 Twitch
- 1 Innocuous Researcher
- 1 Pip-Boy 3000
- 1 Staff of Domination
- 1 Time Spiral
- 1 Shadow of the Second Sun
- 1 Sewer-veillance Cam
- 1 Silent Hallcreeper
- 1 Yavimaya Elder
- 1 Fierce Guardianship
- 1 Force of Will
- 1 Flusterstorm
- 1 Muddle the Mixture
- 1 Narset's Reversal
- 1 An Offer You Can't Refuse
- 1 Katara's Reversal
- 1 Legolas's Quick Reflexes
- 1 Magic Damper
- 1 Octopus Form
- 1 Pemmin's Aura
- 1 Rewind
- 1 Shore Up
- 1 Sokka's Haiku
- 1 Unwind
- 1 Snap
- 1 Emerald Charm
- 1 Provoke
- 1 Bushwhack
- 1 Gleeful Sabotage
- 1 Mists of Lórien
- 1 Aggressive Biomancy
- 1 Apathy
- 1 Amphibian Downpour
- 1 Peregrine Drake
- 1 Filigree Sages
- 1 Voltaic Construct
- 1 Kuldotha Forgemaster
- 1 Protean Hulk
- 1 Seedborn Muse
- 1 Vaultborn Tyrant
- 1 Dread Linnorm // Scale Deflection
- 1 Phyrexian Metamorph
- 1 Dramatic Reversal
- 1 Thassa's Oracle
- 1 Cyclonic Rift
- 1 Tezzeret, Cruel Captain
- 1 Vineglimmer Snarl
- 1 Starting Town

### 14. Bracket 5, Yuriko, the Tiger's Shadow, ninjas

Pool 239. 99 cards, 0 block findings, repaired false.

Judge: bracket 5, agrees. This is a textbook cEDH Yuriko list: dual/fetch mana base with Underground Sea, a huge fast-mana suite (Lion's Eye Diamond, Grim Monolith, Mana Vault, the Moxen, Dark Ritual, Culling the Weak), and ~14 pieces of free/cheap countermagic plus Ad Nauseam and Thassa's Oracle as a win condition. It carries roughly 15 marked Game Changers and a low curve built to kill in a handful of turns, far beyond Bracket 4's casual optimization. The strategy is pure competitive tempo-tutor-combo with no theme concessions.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 30 | 27 to 33 |  |  |
| `tapped_land` | 1 | 0 to 2 |  | Choked Estuary |
| `colorless_land` | 0 | 0 to 6 |  |  |
| `color_sources` | 1.84 | 0.9 or more |  | sources of requirement: U 35 of 19, B 35 of 19 |
| `avg_mana_value` | 2.07 | 1.2 to 2.4 |  | over 69 nonland cards |
| `ramp` | 16 | 12 to 22 |  |  |
| `draw` | 13 | 8 to 18 |  |  |
| `removal` | 9 | 4 to 14 |  |  |
| `wipe` | 1 | 0 to 3 |  |  |
| `interaction` | 15 | 8 to 22 |  |  |
| `tutor` | 4 | 4 or more |  | Demonic Tutor, Imperial Seal, Mystical Tutor, Vampiric Tutor |
| `fast_mana` | 13 | 6 or more |  | An Offer You Can't Refuse, Chrome Mox, Culling the Weak, Dark Ritual, Everflowing Chalice, Lion's Eye Diamond, Lotus Petal, Mana Vault, Mox Amber, Mox Diamond, Mox Opal, Sol Ring, Springleaf Drum |
| `game_changer` | 16 | 8 or more |  | Ad Nauseam, Chrome Mox, Cyclonic Rift, Demonic Tutor, Fierce Guardianship, Force of Will, Grim Monolith, Imperial Seal, Lion's Eye Diamond, Mana Vault, Mox Diamond, Mystical Tutor, Rhystic Study, Thassa's Oracle, The One Ring, Vampiric Tutor |
| `finisher` | 1 | 1 or more |  | Thassa's Oracle |
| `mana_turn_four` | 5.28 | 4.6 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.66 | 0.65 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.71 | -9 to 0.3 |  | the commander costs 3 and comes down on turn 2.29 on average |

Goldfish over 10000 hands: the commander on turn 2.29, 5.28 mana on turn four, and 0.66 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Ad Nauseam, Chrome Mox, Cyclonic Rift, Demonic Tutor, Fierce Guardianship, Force of Will, Grim Monolith, Imperial Seal, Lion's Eye Diamond, Mana Vault, Mox Diamond, Mystical Tutor, Rhystic Study, Thassa's Oracle, The One Ring, Vampiric Tutor. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 2.07 over 69 nonland cards
- INFO `mana_pass`: the builder moved 6 cards of the mana base to bring the deck inside its power level

Cards:

- 1 Underground Sea
- 1 Watery Grave
- 1 Polluted Delta
- 1 Prismatic Vista
- 1 Command Tower
- 1 City of Brass
- 1 Mana Confluence
- 1 Underground River
- 1 Darkslick Shores
- 1 Morphic Pool
- 1 Clearwater Pathway // Murkwater Pathway
- 1 Gloomlake Verge
- 1 Choked Estuary
- 1 Drowned Catacomb
- 1 Shipwreck Marsh
- 1 Sunken Ruins
- 1 Darkwater Catacombs
- 1 Tainted Isle
- 1 Exotic Orchard
- 1 Forbidden Orchard
- 1 Cavern of Souls
- 1 Secluded Courtyard
- 1 Unclaimed Territory
- 1 Spire of Industry
- 1 Plaza of Heroes
- 1 Island
- 1 Swamp
- 1 Chrome Mox
- 1 Culling the Weak
- 1 Dark Ritual
- 1 Everflowing Chalice
- 1 Lion's Eye Diamond
- 1 Lotus Petal
- 1 Mana Vault
- 1 Mox Amber
- 1 Mox Diamond
- 1 Mox Opal
- 1 Sol Ring
- 1 Springleaf Drum
- 1 Grim Monolith
- 1 Arcane Signet
- 1 Fellwar Stone
- 1 Talisman of Dominance
- 1 Ad Nauseam
- 1 Brainstorm
- 1 Faerie Mastermind
- 1 Gitaxian Probe
- 1 Ingenious Infiltrator
- 1 Mystic Remora
- 1 Moon-Circuit Hacker
- 1 Ninja of the Deep Hours
- 1 Ponder
- 1 Preordain
- 1 Rhystic Study
- 1 Sensei's Divining Top
- 1 Windfall
- 1 An Offer You Can't Refuse
- 1 Arcane Denial
- 1 Counterspell
- 1 Dispel
- 1 Fierce Guardianship
- 1 Flusterstorm
- 1 Force of Negation
- 1 Force of Will
- 1 Mental Misstep
- 1 Mindbreak Trap
- 1 Negate
- 1 Pact of Negation
- 1 Swan Song
- 1 Vexing Bauble
- 1 The One Ring
- 1 Chain of Vapor
- 1 Deadly Rollick
- 1 Dismember
- 1 Feed the Swarm
- 1 Infernal Grasp
- 1 Into the Flood Maw
- 1 Pongify
- 1 Rapid Hybridization
- 1 Snuff Out
- 1 Fallen Shinobi
- 1 Foot Mystic
- 1 Futurist Operative
- 1 Kami of Restless Shadows
- 1 Kotose, the Silent Spider
- 1 Moonblade Shinobi
- 1 Mukotai Ambusher
- 1 Sakashima's Student
- 1 Shark Shredder, Killer Clone
- 1 Biting-Palm Ninja
- 1 Thassa's Oracle
- 1 Cyclonic Rift
- 1 Demonic Tutor
- 1 Imperial Seal
- 1 Mystical Tutor
- 1 Vampiric Tutor
- 1 Tarnished Citadel
- 1 Fabled Passage
- 1 Starting Town

### 15. Bracket 5, Najeela, the Blade-Blossom, warriors combat

Pool 339. 99 cards, 0 block findings, repaired false.

Judge: bracket 5, agrees. This is a Najeela cEDH shell: Lion's Eye Diamond, Mana Vault, the full Mox suite, Dark Ritual and Jeska's Will for explosive fast mana, a dozen-plus Game Changers including Imperial Seal, Vampiric/Demonic Tutor, Rhystic Study and Fierce Guardianship, plus Thassa's Oracle as an alternate win. Najeela with Combat Celebrant (or any mana producer off warrior tokens) is an infinite-combat kill available on turn three or four, backed by free counterspells for protection. The only drag is the clunky basic-heavy mana base, but the strategy and card quality are competitive, not thematic.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 33 | 27 to 33 |  |  |
| `tapped_land` | 0 | 0 to 2 |  |  |
| `colorless_land` | 0 | 0 to 6 |  |  |
| `color_sources` | 0.7 | 0.9 or more | yes | sources of requirement: W 13.25 of 19, U 14.25 of 19, B 14.25 of 19, R 14.25 of 19, G 13.75 of 19 |
| `avg_mana_value` | 2.35 | 1.2 to 2.4 |  | over 66 nonland cards |
| `ramp` | 19 | 12 to 22 |  |  |
| `draw` | 11 | 8 to 18 |  |  |
| `removal` | 7 | 4 to 14 |  |  |
| `wipe` | 1 | 0 to 3 |  |  |
| `interaction` | 14 | 8 to 22 |  |  |
| `tutor` | 4 | 4 or more |  | Demonic Tutor, Enlightened Tutor, Imperial Seal, Vampiric Tutor |
| `fast_mana` | 12 | 6 or more |  | An Offer You Can't Refuse, Chrome Mox, Dark Ritual, Lion's Eye Diamond, Lotus Petal, Mana Vault, Mox Amber, Mox Diamond, Mox Opal, Rite of Flame, Sol Ring, Springleaf Drum |
| `game_changer` | 14 | 8 or more |  | Chrome Mox, Demonic Tutor, Enlightened Tutor, Fierce Guardianship, Imperial Seal, Jeska's Will, Lion's Eye Diamond, Mana Vault, Mox Diamond, Orcish Bowmasters, Rhystic Study, Smothering Tithe, Thassa's Oracle, Vampiric Tutor |
| `finisher` | 2 | 1 or more |  | Combat Celebrant, Thassa's Oracle |
| `mana_turn_four` | 5.26 | 4.6 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.71 | 0.65 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.67 | -9 to 0.3 |  | the commander costs 3 and comes down on turn 2.33 on average |

Goldfish over 10000 hands: the commander on turn 2.33, 5.26 mana on turn four, and 0.71 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Chrome Mox, Demonic Tutor, Enlightened Tutor, Fierce Guardianship, Imperial Seal, Jeska's Will, Lion's Eye Diamond, Mana Vault, Mox Diamond, Orcish Bowmasters, Rhystic Study, Smothering Tithe, Thassa's Oracle, Vampiric Tutor. Mass land denial none. Extra turns none. Combos 1.
- Najeela, the Blade-Blossom + Professional Face-Breaker (speed 5, two cards)

Findings:

- INFO `curve_summary`: average mana value 2.35 over 66 nonland cards
- WARN `profile_off_band`: the worst color's share of the sources it needs is 0.7, and bracket 5 wants 0.9 or more (sources of requirement: W 13.25 of 19, U 14.25 of 19, B 14.25 of 19, R 14.25 of 19, G 13.75 of 19)
- INFO `mana_pass`: the builder moved 11 cards of the mana base to bring the deck inside its power level

Cards:

- 1 Command Tower
- 1 Blood Crypt
- 1 Breeding Pool
- 1 Godless Shrine
- 1 Hallowed Fountain
- 1 Overgrown Tomb
- 1 Steam Vents
- 1 Stomping Ground
- 1 Watery Grave
- 5 Plains
- 4 Island
- 4 Swamp
- 5 Mountain
- 6 Forest
- 1 Chrome Mox
- 1 Mana Vault
- 1 Mox Diamond
- 1 Lion's Eye Diamond
- 1 Lotus Petal
- 1 Sol Ring
- 1 Mox Opal
- 1 Mox Amber
- 1 Springleaf Drum
- 1 Dark Ritual
- 1 Rite of Flame
- 1 Jeska's Will
- 1 Smothering Tithe
- 1 Ragavan, Nimble Pilferer
- 1 Professional Face-Breaker
- 1 Grand Warlord Radha
- 1 Rhystic Study
- 1 Mindblade Render
- 1 Azra Oddsmaker
- 1 Akki Ronin
- 1 Eshki Dragonclaw
- 1 Estinien Varlineau
- 1 Ivora, Insatiable Heir
- 1 Jaxis, the Troublemaker
- 1 Kutzil, Malamet Exemplar
- 1 Samut, Vizier of Naktamun
- 1 Tidus, Yuna's Guardian
- 1 An Offer You Can't Refuse
- 1 Fierce Guardianship
- 1 Concerted Defense
- 1 Ainok Strike Leader
- 1 Akiri, Fearless Voyager
- 1 Eladamri, Lord of Leaves
- 1 Hakoda, Selfless Commander
- 1 Kashi-Tribe Elite
- 1 Linvala, Shield of Sea Gate
- 1 Seasoned Dungeoneer
- 1 Selfless Samurai
- 1 Vexilus Praetor
- 1 Orcish Bowmasters
- 1 Accursed Marauder
- 1 Dreadhorde Butcher
- 1 Fleshbag Marauder
- 1 Goblin Cratermaker
- 1 Merciless Executioner
- 1 Combat Celebrant
- 1 Aspiring Champion
- 1 Bull-Rush Bruiser
- 1 Doomskar Warrior
- 1 Emeria Captain
- 1 Heiko Yamazaki, the General
- 1 Kabira Outrider
- 1 Kimahri, Valiant Guardian
- 1 Vengeful Firebrand
- 1 Thassa's Oracle
- 1 Goblin Chainwhirler
- 1 Demonic Tutor
- 1 Enlightened Tutor
- 1 Imperial Seal
- 1 Vampiric Tutor
- 1 Multiclass Baldric
- 1 Rhys, the Evermore
- 1 Cacophony Scamp
- 1 Radha, Heir to Keld
- 1 Shang-Chi, Master of Kung Fu
- 1 Heronblade Elite

