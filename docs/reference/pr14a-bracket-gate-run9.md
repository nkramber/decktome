# PR-14A bracket gate

Run date: 2026-09-23. Card snapshot: 2026-09-04.

Verdict: FAIL. 6 of 6 decks passed every block check, 5 sat in every band, 6 held no content violation, and the judge agreed with the bracket on 4 of 6 (66 percent, the bar is 80).

## Summary

| Measure | Value |
|---|---|
| Prompts | 6 |
| Decks returned | 6 |
| Decks with no block finding | 6 |
| Decks in every band | 5 |
| Decks with no content violation | 6 |
| Decks the endpoint did not check | 0 |
| Decks that needed a repair turn | 0 |
| Decks judged | 6 |
| Judge agreed with the bracket | 4 |
| Judge errors | 0 |
| Errors | 0 |
| Prompt version | 16 |
| Calls | 12 |
| Cost | $0.5312 |
| Time | 471 seconds |

## Run

- Suite `bracket`, run `pr14a-bracket-gate-run9`, on 2026-09-23, commit `44a03e6`.
- Partial run over `10,11,12,13,14,15`. It reads the bars of its own items, and it never stands for the suite.
- Roles: generate on `gpt-5.6-terra` (openai, effort medium), judge on `claude-opus-5` (anthropic, effort medium), repair on `gpt-5.6-terra` (openai, effort medium).
- Versions: card snapshot 2026-09-04, bracket_judge prompt version 3, generate prompt version 16.
- Calls: 12. Cost: $0.5312. Time: 471 seconds.

## Off-band features by key

Each row counts the decks whose feature sat outside its band after the repair turns.

| Feature | Decks off band |
|---|---|
| `tapped_land` | 1 |
| `fixing_land` | 1 |

## Decks

### 10. Bracket 4, Urza, Lord High Artificer, artifacts

Pool 319. 99 cards, 0 block findings, repaired false.

Judge: bracket 4, agrees. Seven marked Game Changers (Chrome Mox, Lion's Eye Diamond, Mana Vault, Mox Diamond, Rhystic Study, Fierce Guardianship, Force of Will, Thassa's Oracle) blow past the three-GC cap, and the density of fast mana plus Urza's token/mana engine lets it deploy huge turns very early. Free counterspells and artifact tutors like Whir of Invention/Transmute Artifact protect a fast Thassa's Oracle finish. It falls short of cEDH, though, since the mana base is 29 basic Islands with a themed artifact shell rather than an optimized metagame list.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 34 | 31 to 36 |  |  |
| `tapped_land` | 1 | 0 to 5 |  | Sink into Stupor // Soporific Springs |
| `colorless_land` | 1 | 0 to 5 |  | Academy Ruins |
| `color_sources` | 1.38 | 0.9 or more |  | sources of requirement: U 36 of 26 |
| `avg_mana_value` | 2.91 | 1.6 to 3 |  | over 65 nonland cards |
| `ramp` | 13 | 10 to 16 |  |  |
| `draw` | 12 | 8 to 16 |  |  |
| `removal` | 10 | 6 to 14 |  |  |
| `wipe` | 2 | 0 to 5 |  |  |
| `interaction` | 11 | 6 to 16 |  |  |
| `tutor` | 3 | 2 or more |  | Tezzeret the Seeker, Transmute Artifact, Whir of Invention |
| `fast_mana` | 10 | 3 or more |  | An Offer You Can't Refuse, Chrome Mox, Lion's Eye Diamond, Lotus Petal, Mana Vault, Moonsnare Prototype, Mox Diamond, Mox Opal, Sol Ring, Springleaf Drum |
| `game_changer` | 8 | 4 or more |  | Chrome Mox, Fierce Guardianship, Force of Will, Lion's Eye Diamond, Mana Vault, Mox Diamond, Rhystic Study, Thassa's Oracle |
| `finisher` | 3 | 2 or more |  | Mechanized Production, Mirrodin Besieged, Thassa's Oracle |
| `mana_turn_four` | 5.03 | 4.4 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.72 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.68 | -9 to 0.5 |  | the commander costs 4 and comes down on turn 3.32 on average |

Goldfish over 10000 hands: the commander on turn 3.32, 5.03 mana on turn four, and 0.72 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Chrome Mox, Lion's Eye Diamond, Mana Vault, Mox Diamond, Rhystic Study, Fierce Guardianship, Force of Will, Thassa's Oracle. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 2.91 over 65 nonland cards
- INFO `mana_pass`: the builder moved 1 card of the mana base to bring the deck inside its power level

Cards:

- 29 Island
- 1 Academy Ruins
- 1 Mystic Sanctuary
- 1 Otawara, Soaring City
- 1 Seat of the Synod
- 1 Sink into Stupor // Soporific Springs
- 1 Chrome Mox
- 1 Lion's Eye Diamond
- 1 Mana Vault
- 1 Mox Diamond
- 1 Mox Opal
- 1 Lotus Petal
- 1 Sol Ring
- 1 Springleaf Drum
- 1 Moonsnare Prototype
- 1 Metalworker
- 1 Krark-Clan Ironworks
- 1 Chief Engineer
- 1 Tezzeret the Seeker
- 1 Rhystic Study
- 1 Braided Net // Braided Quipu
- 1 Cerebral Download
- 1 Era of Innovation
- 1 Esoteric Duplicator
- 1 Forensic Gadgeteer
- 1 Riddlesmith
- 1 Sai, Master Thopterist
- 1 Thirst for Knowledge
- 1 Thoughtcast
- 1 Vedalken Archmage
- 1 Thought Monitor
- 1 An Offer You Can't Refuse
- 1 Assert Authority
- 1 Disruption Protocol
- 1 Fierce Guardianship
- 1 Force of Will
- 1 Ice Out
- 1 Metallic Rebuke
- 1 Padeem, Consul of Innovation
- 1 Stoic Rebuttal
- 1 Welding Jar
- 1 Ghostly Flicker
- 1 Aether Spellbomb
- 1 Blasting Station
- 1 Contagion Clasp
- 1 Cyber Conversion
- 1 Into Thin Air
- 1 Kitesail Larcenist
- 1 Ravenform
- 1 Resculpt
- 1 Unable to Scream
- 1 Watery Grasp
- 1 Engineered Explosives
- 1 Hurkyl's Recall
- 1 Transmute Artifact
- 1 Whir of Invention
- 1 Mystic Forge
- 1 Arcbound Crusher
- 1 Chrome Steed
- 1 Filigree Attendant
- 1 Frogmite
- 1 Ironheart, Clever Champion
- 1 Kappa Cannoneer
- 1 Karn, Scion of Urza
- 1 Mechan Assembler
- 1 Memory Guardian
- 1 Traxos, Scourge of Kroog
- 1 Utrom Monitor
- 1 Mechanized Production
- 1 Mirrodin Besieged
- 1 Thassa's Oracle

### 11. Bracket 4, Korvold, Fae-Cursed King, treasure sacrifice

Pool 328. 99 cards, 0 block findings, repaired false.

Judge: bracket 4, agrees. Eight marked Game Changers (Chrome Mox, Jeska's Will, Lion's Eye Diamond, Mana Vault, Mox Diamond, Orcish Bowmasters, Demonic Tutor, Gamble, Vampiric Tutor) blow far past the three-Game-Changer ceiling of Bracket 3, and the deck backs them with dual/original-duals, Sol Ring, Lotus Petal, Dark Ritual and multiple tutors for explosive starts. Even without a Spellbook-listed combo, the card quality, fast mana density and sacrifice-engine threats (Yawgmoth, Goblin Bombardment, Prossh, Ziatora) let it close games very early. It's a tuned value/aristocrats build rather than a metagame-honed cEDH list, so it sits at Optimized.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 33 | 31 to 36 |  |  |
| `tapped_land` | 2 | 0 to 5 |  | Disciple of Freyalise // Garden of Freyalise, Kazuul's Fury // Kazuul's Cliffs |
| `colorless_land` | 2 | 0 to 5 |  | Dragon-Cursed Halls, Westvale Abbey // Ormendahl, Profane Prince |
| `color_sources` | 1.2 | 0.9 or more |  | sources of requirement: B 22.75 of 19, R 23.75 of 19, G 22.75 of 19 |
| `fixing_land` | 21 | 12 or more |  | Badlands, Bayou, Blightstep Pathway // Searstep Pathway, Blood Crypt, City of Brass, Command Tower, Cragcrown Pathway // Timbercrown Pathway, Darkbore Pathway // Slitherbore Pathway, Forbidden Orchard, Grove of the Burnwillows, Heap Gate, Karplusan Forest, Llanowar Wastes, Mana Confluence, Mount Doom, Nurturing Peatland, Overgrown Tomb, Stomping Ground, Sulfurous Springs, Taiga, Tarnished Citadel |
| `avg_mana_value` | 2.98 | 1.6 to 3 |  | over 66 nonland cards |
| `ramp` | 14 | 10 to 16 |  |  |
| `draw` | 12 | 8 to 16 |  |  |
| `removal` | 10 | 6 to 14 |  |  |
| `wipe` | 2 | 0 to 5 |  |  |
| `interaction` | 8 | 6 to 16 |  |  |
| `tutor` | 4 | 2 or more |  | Demonic Tutor, Gamble, Sidisi, Undead Vizier, Vampiric Tutor |
| `fast_mana` | 11 | 3 or more |  | Chrome Mox, Dark Ritual, Lion's Eye Diamond, Lotus Petal, Mana Vault, Mox Amber, Mox Diamond, Mox Opal, Rite of Flame, Sol Ring, Springleaf Drum |
| `game_changer` | 9 | 4 or more |  | Chrome Mox, Demonic Tutor, Gamble, Jeska's Will, Lion's Eye Diamond, Mana Vault, Mox Diamond, Orcish Bowmasters, Vampiric Tutor |
| `finisher` | 10 | 2 or more |  | Prossh, Skyraider of Kher, Rakdos, the Muscle, Relic Vial, Revel in Riches, Smaug the Impenetrable, Terror Ballista, The Meathook Massacre, Venerated Stormsinger, Vraska, Golgari Queen, Ziatora, the Incinerator |
| `mana_turn_four` | 4.93 | 4.4 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.71 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.65 | -9 to 0.5 |  | the commander costs 5 and comes down on turn 4.35 on average |

Goldfish over 10000 hands: the commander on turn 4.35, 4.93 mana on turn four, and 0.71 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Chrome Mox, Jeska's Will, Lion's Eye Diamond, Mana Vault, Mox Diamond, Orcish Bowmasters, Demonic Tutor, Gamble, Vampiric Tutor. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 2.98 over 66 nonland cards
- INFO `mana_pass`: the builder moved 2 cards of the mana base to bring the deck inside its power level

Cards:

- 1 Badlands
- 1 Bayou
- 1 Blood Crypt
- 1 Blightstep Pathway // Searstep Pathway
- 1 City of Brass
- 1 Command Tower
- 1 Cragcrown Pathway // Timbercrown Pathway
- 1 Darkbore Pathway // Slitherbore Pathway
- 1 Disciple of Freyalise // Garden of Freyalise
- 1 Dragon-Cursed Halls
- 1 Forbidden Orchard
- 2 Forest
- 1 Grove of the Burnwillows
- 1 Heap Gate
- 1 Karplusan Forest
- 1 Kazuul's Fury // Kazuul's Cliffs
- 1 Llanowar Wastes
- 1 Mana Confluence
- 3 Mountain
- 1 Nurturing Peatland
- 1 Overgrown Tomb
- 1 Phyrexian Tower
- 1 Stomping Ground
- 2 Swamp
- 1 Sulfurous Springs
- 1 Taiga
- 1 Tarnished Citadel
- 1 Westvale Abbey // Ormendahl, Profane Prince
- 1 Chrome Mox
- 1 Dark Ritual
- 1 Jeska's Will
- 1 Lion's Eye Diamond
- 1 Lotus Petal
- 1 Mana Vault
- 1 Mox Diamond
- 1 Mox Opal
- 1 Pitiless Plunderer
- 1 Professional Face-Breaker
- 1 Rite of Flame
- 1 Sol Ring
- 1 Springleaf Drum
- 1 Smothering Abomination
- 1 Disciple of Bolas
- 1 Dina, Essence Brewer
- 1 Bushmeat Poacher
- 1 Jerren, Corrupted Bishop // Ormendahl, the Corrupter
- 1 Lord Skitter's Butcher
- 1 Relic Vial
- 1 Sanguine Spy
- 1 Shadowheart, Dark Justiciar
- 1 Skyclave Shadowcat
- 1 Tevesh Szat, Doom of Fools
- 1 Vampire Gourmand
- 1 Attrition
- 1 Bone Shards
- 1 Broadside Bombardiers
- 1 Eaten Alive
- 1 Goblin Bombardment
- 1 Grist, the Hunger Tide
- 1 Minsc & Boo, Timeless Heroes
- 1 Orcish Bowmasters
- 1 Terror Ballista
- 1 Yawgmoth, Thran Physician
- 1 Altar of the Wretched // Wretched Bonemass
- 1 Constant Mists
- 1 Eldrazi Monument
- 1 Gift of Doom
- 1 Gyome, Master Chef
- 1 Nightmare Shepherd
- 1 Sylvan Safekeeper
- 1 Yarus, Roar of the Old Gods
- 1 Academy Manufactor
- 1 Viscera Seer
- 1 Xorn
- 1 Flesh-Eater Imp
- 1 Immersturm Predator
- 1 Kethek, Crucible Goliath
- 1 Meren of Clan Nel Toth
- 1 Prossh, Skyraider of Kher
- 1 Rakdos, the Muscle
- 1 Sadistic Hypnotist
- 1 Sidisi, Undead Vizier
- 1 Smaug the Impenetrable
- 1 Tomb Tyrant
- 1 Venerated Stormsinger
- 1 Revel in Riches
- 1 Vraska, Golgari Queen
- 1 Ziatora, the Incinerator
- 1 The Meathook Massacre
- 1 Spiteful Banditry
- 1 Demonic Tutor
- 1 Gamble
- 1 Vampiric Tutor
- 1 Mox Amber
- 1 Mount Doom

### 12. Bracket 4, Prosper, Tome-Bound, treasure

Pool 212. 99 cards, 0 block findings, repaired false.

Judge: bracket 5, disagrees. This is a cEDH-caliber Prosper build: twelve-plus Game Changers including Ad Nauseam, Lion's Eye Diamond, Mox Diamond, Mana Vault, Chrome Mox, Underworld Breach, and a full tutor suite (Demonic, Vampiric, Imperial Seal, Gamble). The mana base is dual/fetch-heavy with heavy fast mana, supporting Breach + LED + Wheel storm lines that can win on turn two or three, plus free interaction (Deadly Rollick, Snuff Out, Pyroblast, REB) to protect them. That speed, redundancy, and metagame-tuned interaction package is competitive rather than merely optimized.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 33 | 31 to 36 |  |  |
| `tapped_land` | 1 | 0 to 5 |  | Foreboding Ruins |
| `colorless_land` | 0 | 0 to 5 |  |  |
| `color_sources` | 1.68 | 0.9 or more |  | sources of requirement: B 33 of 19, R 32 of 19 |
| `fixing_land` | 26 | 9 or more |  | Badlands, Blackcleave Cliffs, Blazemire Verge, Blightstep Pathway // Searstep Pathway, Blood Crypt, Bloodstained Mire, City of Brass, Command Tower, Dragonskull Summit, Exotic Orchard, Fabled Passage, Forbidden Orchard, Foreboding Ruins, Graven Cairns, Haunted Ridge, Luxury Suite, Mana Confluence, Mount Doom, Prismatic Vista, Shadowblood Ridge, Smoldering Marsh, Spire of Industry, Starting Town, Sulfurous Springs, Tainted Peak, Tarnished Citadel |
| `avg_mana_value` | 2.53 | 1.6 to 3 |  | over 66 nonland cards |
| `ramp` | 13 | 10 to 16 |  |  |
| `draw` | 12 | 8 to 16 |  |  |
| `removal` | 10 | 6 to 14 |  |  |
| `wipe` | 2 | 0 to 5 |  |  |
| `interaction` | 11 | 6 to 16 |  |  |
| `tutor` | 4 | 2 or more |  | Demonic Tutor, Gamble, Imperial Seal, Vampiric Tutor |
| `fast_mana` | 9 | 3 or more |  | Chrome Mox, Dark Ritual, Lion's Eye Diamond, Lotus Petal, Mana Vault, Mox Diamond, Mox Opal, Sol Ring, Springleaf Drum |
| `game_changer` | 13 | 4 or more |  | Ad Nauseam, Chrome Mox, Demonic Tutor, Gamble, Imperial Seal, Jeska's Will, Lion's Eye Diamond, Mana Vault, Mox Diamond, Orcish Bowmasters, The One Ring, Underworld Breach, Vampiric Tutor |
| `finisher` | 4 | 2 or more |  | Knuckles the Echidna, Revel in Riches, Smaug the Impenetrable, Smaug, Wicked Worm |
| `mana_turn_four` | 4.95 | 4.4 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.71 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.65 | -9 to 0.5 |  | the commander costs 4 and comes down on turn 3.35 on average |

Goldfish over 10000 hands: the commander on turn 3.35, 4.95 mana on turn four, and 0.71 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Chrome Mox, Jeska's Will, Lion's Eye Diamond, Mana Vault, Mox Diamond, Ad Nauseam, The One Ring, Orcish Bowmasters, Demonic Tutor, Gamble, Imperial Seal, Underworld Breach, Vampiric Tutor. Mass land denial none. Extra turns none. Combos 4.
- Orcish Bowmasters + Peer into the Abyss (speed 3, two cards)
- Underworld Breach + Wheel of Fortune + Jeska's Will (speed 5)
- Prosper, Tome-Bound + Professional Face-Breaker + Sensei's Divining Top + Xorn (speed 5)
- Underworld Breach + Lion's Eye Diamond + Wheel of Fortune (speed 5)

Findings:

- INFO `curve_summary`: average mana value 2.53 over 66 nonland cards
- INFO `mana_pass`: the builder moved 2 cards of the mana base to bring the deck inside its power level

Cards:

- 1 Badlands
- 1 Blackcleave Cliffs
- 1 Blazemire Verge
- 1 Blightstep Pathway // Searstep Pathway
- 1 Blood Crypt
- 1 Bloodstained Mire
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
- 1 Shadowblood Ridge
- 1 Smoldering Marsh
- 1 Spire of Industry
- 1 Sulfurous Springs
- 1 Tainted Peak
- 1 Tarnished Citadel
- 4 Swamp
- 3 Mountain
- 1 Chrome Mox
- 1 Dark Ritual
- 1 Jeska's Will
- 1 Lion's Eye Diamond
- 1 Lotus Petal
- 1 Mana Vault
- 1 Mox Diamond
- 1 Mox Opal
- 1 Sol Ring
- 1 Springleaf Drum
- 1 Black Market Connections
- 1 Pitiless Plunderer
- 1 Professional Face-Breaker
- 1 Ad Nauseam
- 1 Corrupted Conviction
- 1 Demand Answers
- 1 Faithless Looting
- 1 Idol of Oblivion
- 1 Mishra's Bauble
- 1 Night's Whisper
- 1 Peer into the Abyss
- 1 Reckless Lackey
- 1 Sensei's Divining Top
- 1 Skullclamp
- 1 Wheel of Fortune
- 1 Brotherhood Regalia
- 1 Kaya's Ghostform
- 1 Lightning Greaves
- 1 Mithril Coat
- 1 Not Dead After All
- 1 Pyroblast
- 1 Red Elemental Blast
- 1 Swiftfoot Boots
- 1 Tibalt's Trickery
- 1 The One Ring
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
- 1 Knuckles the Echidna
- 1 Revel in Riches
- 1 Smaug, Wicked Worm
- 1 Blasphemous Act
- 1 Toxic Deluge
- 1 Abrade
- 1 Bedevil
- 1 Chaos Warp
- 1 Deadly Rollick
- 1 Feed the Swarm
- 1 Go for the Throat
- 1 Infernal Grasp
- 1 Orcish Bowmasters
- 1 Snuff Out
- 1 Terminate
- 1 Demonic Tutor
- 1 Gamble
- 1 Imperial Seal
- 1 Underworld Breach
- 1 Vampiric Tutor
- 1 Mount Doom
- 1 Starting Town

### 13. Bracket 5, Kinnan, Bonder Prodigy, combo

Pool 322. 99 cards, 0 block findings, repaired false.

Judge: bracket 5, agrees. This is a textbook cEDH Kinnan build: free counterspells (Force of Will, Fierce Guardianship), a dozen+ Game Changers including Grim Monolith, Mana Vault, LED, Mox Diamond and Rhystic Study, a fetch/dual mana base, and Thassa's Oracle as the win condition. It packs zero-mana two-card infinites with the commander (Basalt Monolith, Grim Monolith + Spark Double) plus untap-enablers and tutors to assemble them by turn two or three. It is metagame-optimized rather than themed, so it sits squarely at the competitive tier.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 32 | 27 to 33 |  |  |
| `tapped_land` | 1 | 0 to 2 |  | Vineglimmer Snarl |
| `colorless_land` | 0 | 0 to 6 |  |  |
| `color_sources` | 1.25 | 0.9 or more |  | sources of requirement: U 32.5 of 26, G 28.75 of 19 |
| `fixing_land` | 15 | 9 or more |  | Breeding Pool, City of Brass, Command Tower, Dreamroot Cascade, Fabled Passage, Forbidden Orchard, Hinterland Harbor, Mana Confluence, Misty Rainforest, Prismatic Vista, Rejuvenating Springs, Starting Town, Tropical Island, Vineglimmer Snarl, Yavimaya Coast |
| `avg_mana_value` | 2.4 | 1.2 to 2.4 |  | over 67 nonland cards |
| `ramp` | 18 | 12 to 22 |  |  |
| `draw` | 11 | 8 to 18 |  |  |
| `removal` | 9 | 4 to 14 |  |  |
| `wipe` | 1 | 0 to 3 |  |  |
| `interaction` | 15 | 8 to 22 |  |  |
| `tutor` | 7 | 4 or more |  | Alphinaud Leveilleur, Arcum Dagsson, Intuition, Kuldotha Forgemaster, Muddle the Mixture, Primal Command, Vedalken Aethermage |
| `fast_mana` | 10 | 6 or more |  | An Offer You Can't Refuse, Chrome Mox, Lion's Eye Diamond, Lotus Petal, Mana Vault, Mox Amber, Mox Diamond, Mox Opal, Sol Ring, Springleaf Drum |
| `game_changer` | 12 | 8 or more |  | Chrome Mox, Cyclonic Rift, Fierce Guardianship, Force of Will, Grim Monolith, Intuition, Lion's Eye Diamond, Mana Vault, Mox Diamond, Rhystic Study, Seedborn Muse, Thassa's Oracle |
| `finisher` | 1 | 1 or more |  | Thassa's Oracle |
| `mana_turn_four` | 5.28 | 4.6 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.69 | 0.65 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.41 | -9 to 0.3 |  | the commander costs 2 and comes down on turn 1.59 on average |

Goldfish over 10000 hands: the commander on turn 1.59, 5.28 mana on turn four, and 0.69 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Chrome Mox, Grim Monolith, Lion's Eye Diamond, Mana Vault, Mox Diamond, Rhystic Study, Force of Will, Fierce Guardianship, Intuition, Seedborn Muse, Thassa's Oracle, Cyclonic Rift. Mass land denial none. Extra turns none. Combos 2.
- Kinnan, Bonder Prodigy + Grim Monolith + Spark Double (speed 5, two cards)
- Kinnan, Bonder Prodigy + Basalt Monolith (speed 5, two cards)

Findings:

- INFO `curve_summary`: average mana value 2.40 over 67 nonland cards
- INFO `mana_pass`: the builder moved 9 cards of the mana base to bring the deck inside its power level
- INFO `basics_added`: the list was 1 card short, so the builder added 1 basic land

Cards:

- 1 Tropical Island
- 1 Breeding Pool
- 1 Command Tower
- 1 City of Brass
- 1 Mana Confluence
- 1 Rejuvenating Springs
- 1 Yavimaya Coast
- 1 Misty Rainforest
- 1 Flooded Strand
- 1 Polluted Delta
- 1 Scalding Tarn
- 1 Windswept Heath
- 1 Verdant Catacombs
- 1 Wooded Foothills
- 1 Prismatic Vista
- 7 Island
- 4 Forest
- 1 Chrome Mox
- 1 Grim Monolith
- 1 Lion's Eye Diamond
- 1 Mana Vault
- 1 Mox Diamond
- 1 Lotus Petal
- 1 Mox Amber
- 1 Mox Opal
- 1 Sol Ring
- 1 Springleaf Drum
- 1 Basalt Monolith
- 1 Devoted Druid
- 1 Arbor Elf
- 1 Kiora's Follower
- 1 Gwenna, Eyes of Gaea
- 1 Earthcraft
- 1 Honor-Worn Shaku
- 1 Rhystic Study
- 1 Alphinaud Leveilleur
- 1 Frantic Search
- 1 Cloud of Faeries
- 1 Cerulean Wisps
- 1 Refocus
- 1 Twitch
- 1 Staff of Domination
- 1 Pip-Boy 3000
- 1 Finale of Revelation
- 1 Force of Will
- 1 Fierce Guardianship
- 1 Flusterstorm
- 1 Muddle the Mixture
- 1 Narset's Reversal
- 1 Unwind
- 1 Rewind
- 1 Shore Up
- 1 Legolas's Quick Reflexes
- 1 Katara's Reversal
- 1 Magic Damper
- 1 Biosynthic Burst
- 1 Octopus Form
- 1 Sokka's Haiku
- 1 Pemmin's Aura
- 1 Arcum Dagsson
- 1 Snap
- 1 Emerald Charm
- 1 Gleeful Sabotage
- 1 Bushwhack
- 1 Provoke
- 1 Primal Command
- 1 Vedalken Aethermage
- 1 Blasting Station
- 1 Intuition
- 1 Filigree Sages
- 1 Voltaic Construct
- 1 Kuldotha Forgemaster
- 1 Seedborn Muse
- 1 Peregrine Drake
- 1 Dross Scorpion
- 1 Breaching Hippocamp
- 1 Phyrexian Metamorph
- 1 Clever Impersonator
- 1 Stunt Double
- 1 Spark Double
- 1 Thassa's Oracle
- 1 Cyclonic Rift
- 1 An Offer You Can't Refuse
- 1 Sewer-veillance Cam
- 1 Forbidden Orchard
- 1 Fabled Passage
- 1 Hinterland Harbor
- 1 Dreamroot Cascade
- 1 Vineglimmer Snarl
- 1 Starting Town

### 14. Bracket 5, Yuriko, the Tiger's Shadow, ninjas

Pool 276. 99 cards, 0 block findings, repaired false.

Judge: bracket 5, agrees. This is a textbook cEDH Yuriko list: Underground Sea, Chrome Mox, Mox Diamond, Lion's Eye Diamond, Mana Vault, Grim Monolith, Dark Ritual and Lotus Petal for explosive turn-one starts, a dozen-plus free/cheap counterspells (Force of Will, Fierce Guardianship, Pact of Negation, Mental Misstep), and a full tutor suite (Imperial Seal, Vampiric/Demonic/Mystical Tutor, Wishclaw Talisman) into Ad Nauseam and Thassa's Oracle. It carries roughly fifteen marked Game Changers, far beyond Bracket 3 or 4 norms, and is metagame-tuned rather than themed. Games can be decided on turn two or three, which is squarely Bracket 5.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 31 | 27 to 33 |  |  |
| `tapped_land` | 1 | 0 to 2 |  | Choked Estuary |
| `colorless_land` | 0 | 0 to 6 |  |  |
| `color_sources` | 1.67 | 0.9 or more |  | sources of requirement: U 34.75 of 19, B 36.75 of 22 |
| `fixing_land` | 27 | 9 or more |  | Cavern of Souls, Choked Estuary, City of Brass, Clearwater Pathway // Murkwater Pathway, Command Tower, Darkslick Shores, Drowned Catacomb, Exotic Orchard, Fabled Passage, Forbidden Orchard, Mana Confluence, Morphic Pool, Plaza of Heroes, Polluted Delta, Prismatic Vista, Secluded Courtyard, Shipwreck Marsh, Spire of Industry, Starting Town, Sunken Hollow, Sunken Ruins, Tainted Isle, Tarnished Citadel, Unclaimed Territory, Underground River, Underground Sea, Watery Grave |
| `avg_mana_value` | 2.32 | 1.2 to 2.4 |  | over 68 nonland cards |
| `ramp` | 16 | 12 to 22 |  |  |
| `draw` | 8 | 8 to 18 |  |  |
| `removal` | 9 | 4 to 14 |  |  |
| `wipe` | 1 | 0 to 3 |  |  |
| `interaction` | 15 | 8 to 22 |  |  |
| `tutor` | 5 | 4 or more |  | Demonic Tutor, Imperial Seal, Mystical Tutor, Vampiric Tutor, Wishclaw Talisman |
| `fast_mana` | 12 | 6 or more |  | An Offer You Can't Refuse, Chrome Mox, Culling the Weak, Dark Ritual, Lion's Eye Diamond, Lotus Petal, Mana Vault, Mox Amber, Mox Diamond, Mox Opal, Sol Ring, Springleaf Drum |
| `game_changer` | 16 | 8 or more |  | Ad Nauseam, Chrome Mox, Cyclonic Rift, Demonic Tutor, Fierce Guardianship, Force of Will, Grim Monolith, Imperial Seal, Lion's Eye Diamond, Mana Vault, Mox Diamond, Mystical Tutor, Rhystic Study, Thassa's Oracle, The One Ring, Vampiric Tutor |
| `finisher` | 1 | 1 or more |  | Thassa's Oracle |
| `mana_turn_four` | 5.33 | 4.6 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.68 | 0.65 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.66 | -9 to 0.3 |  | the commander costs 3 and comes down on turn 2.34 on average |

Goldfish over 10000 hands: the commander on turn 2.34, 5.33 mana on turn four, and 0.68 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Chrome Mox, Lion's Eye Diamond, Mana Vault, Mox Diamond, Grim Monolith, Ad Nauseam, Rhystic Study, Fierce Guardianship, Force of Will, The One Ring, Thassa's Oracle, Cyclonic Rift, Demonic Tutor, Imperial Seal, Mystical Tutor, Vampiric Tutor. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 2.32 over 68 nonland cards
- INFO `mana_pass`: the builder moved 3 cards of the mana base to bring the deck inside its power level
- INFO `basics_added`: the list was 1 card short, so the builder added 1 basic land

Cards:

- 1 Underground Sea
- 1 Watery Grave
- 1 Morphic Pool
- 1 Darkslick Shores
- 1 Choked Estuary
- 1 Drowned Catacomb
- 1 Shipwreck Marsh
- 1 Clearwater Pathway // Murkwater Pathway
- 1 Command Tower
- 1 City of Brass
- 1 Mana Confluence
- 1 Exotic Orchard
- 1 Forbidden Orchard
- 1 Polluted Delta
- 1 Prismatic Vista
- 1 Fabled Passage
- 1 Sunken Ruins
- 1 Underground River
- 1 Tainted Isle
- 1 Secluded Courtyard
- 1 Unclaimed Territory
- 1 Cavern of Souls
- 1 Spire of Industry
- 1 Plaza of Heroes
- 1 Island
- 3 Swamp
- 1 Chrome Mox
- 1 Lion's Eye Diamond
- 1 Mana Vault
- 1 Mox Diamond
- 1 Mox Opal
- 1 Lotus Petal
- 1 Sol Ring
- 1 Springleaf Drum
- 1 Grim Monolith
- 1 Arcane Signet
- 1 Dimir Signet
- 1 Talisman of Dominance
- 1 Fellwar Stone
- 1 Mox Amber
- 1 Dark Ritual
- 1 Culling the Weak
- 1 Ad Nauseam
- 1 Rhystic Study
- 1 Mystic Remora
- 1 Gitaxian Probe
- 1 Sensei's Divining Top
- 1 Ingenious Infiltrator
- 1 Moon-Circuit Hacker
- 1 Ninja of the Deep Hours
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
- 1 Go for the Throat
- 1 Infernal Grasp
- 1 Into the Flood Maw
- 1 Pongify
- 1 Rapid Hybridization
- 1 Dokuchi Shadow-Walker
- 1 Foot Mystic
- 1 Fallen Shinobi
- 1 Futurist Operative
- 1 Ink-Eyes, Servant of Oni
- 1 Kami of Restless Shadows
- 1 Kotose, the Silent Spider
- 1 Moonblade Shinobi
- 1 Mukotai Ambusher
- 1 Ninja of the New Moon
- 1 Okiba-Gang Shinobi
- 1 Silent-Blade Oni
- 1 Satoru Umezawa
- 1 Thassa's Oracle
- 1 Cyclonic Rift
- 1 Demonic Tutor
- 1 Imperial Seal
- 1 Mystical Tutor
- 1 Vampiric Tutor
- 1 Wishclaw Talisman
- 1 Tarnished Citadel
- 1 Sunken Hollow
- 1 Starting Town

### 15. Bracket 5, Najeela, the Blade-Blossom, warriors combat

Pool 338. 99 cards, 0 block findings, repaired false.

Judge: bracket 4, disagrees. Over a dozen marked Game Changers, a full suite of Moxen/Lion's Eye Diamond/Mana Vault fast mana, four unconditional tutors, free counterspells, and Najeela plus Combat Celebrant for an infinite-combat kill put this far past Bracket 3's three-Game-Changer cap. It's clearly optimized and can win by turn three or four off ritual-powered Najeela activations. It stops short of cEDH because the build still leans on a warrior-tribal theme with filler payoffs (Jazal, Pact of the Serpent, tribal lands) rather than the purely best line.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 32 | 27 to 33 |  |  |
| `tapped_land` | 5 | 0 to 2 | yes | Base Camp, Gnottvold Slumbermound, Great Hall of Starnheim, Path of Ancestry, Skemfar Elderhall |
| `colorless_land` | 1 | 0 to 6 |  | Accursed Duneyard |
| `color_sources` | 1.08 | 0.9 or more |  | sources of requirement: W 20.5 of 19, U 21.5 of 19, B 22.5 of 19, R 22.5 of 19, G 24.75 of 19 |
| `fixing_land` | 16 | 21 or more | yes | Abundant Countryside, Base Camp, Blood Crypt, Breeding Pool, Cavern of Souls, Command Tower, Godless Shrine, Hallowed Fountain, Mirrex, Overgrown Tomb, Path of Ancestry, Secluded Courtyard, Steam Vents, Stomping Ground, Unclaimed Territory, Watery Grave |
| `avg_mana_value` | 2.39 | 1.2 to 2.4 |  | over 67 nonland cards |
| `ramp` | 18 | 12 to 22 |  |  |
| `draw` | 10 | 8 to 18 |  |  |
| `removal` | 8 | 4 to 14 |  |  |
| `wipe` | 1 | 0 to 3 |  |  |
| `interaction` | 15 | 8 to 22 |  |  |
| `tutor` | 5 | 4 or more |  | Cazur, Ruthless Stalker, Demonic Tutor, Enlightened Tutor, Mystical Tutor, Vampiric Tutor |
| `fast_mana` | 12 | 6 or more |  | An Offer You Can't Refuse, Chrome Mox, Dark Ritual, Lion's Eye Diamond, Lotus Petal, Mana Vault, Mox Amber, Mox Diamond, Mox Opal, Rite of Flame, Sol Ring, Springleaf Drum |
| `game_changer` | 13 | 8 or more |  | Chrome Mox, Demonic Tutor, Enlightened Tutor, Fierce Guardianship, Jeska's Will, Lion's Eye Diamond, Mana Vault, Mox Diamond, Mystical Tutor, Orcish Bowmasters, Rhystic Study, Smothering Tithe, Vampiric Tutor |
| `finisher` | 3 | 1 or more |  | Combat Celebrant, Frodo, Sauron's Bane, Raiyuu, Storm's Edge |
| `mana_turn_four` | 5.12 | 4.6 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.69 | 0.65 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.61 | -9 to 0.3 |  | the commander costs 3 and comes down on turn 2.39 on average |

Goldfish over 10000 hands: the commander on turn 2.39, 5.12 mana on turn four, and 0.69 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Chrome Mox, Lion's Eye Diamond, Mana Vault, Mox Diamond, Jeska's Will, Smothering Tithe, Rhystic Study, Fierce Guardianship, Orcish Bowmasters, Demonic Tutor, Enlightened Tutor, Mystical Tutor, Vampiric Tutor. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 2.39 over 67 nonland cards
- WARN `profile_off_band`: the count of lands that enter tapped is 5, and bracket 5 wants 2 at most (Base Camp, Gnottvold Slumbermound, Great Hall of Starnheim, Path of Ancestry, Skemfar Elderhall)
- WARN `profile_off_band`: the count of lands that make two or more of the deck's colors is 16, and bracket 5 wants 21 or more (Abundant Countryside, Base Camp, Blood Crypt, Breeding Pool, Cavern of Souls, Command Tower, Godless Shrine, Hallowed Fountain, Mirrex, Overgrown Tomb, Path of Ancestry, Secluded Courtyard, Steam Vents, Stomping Ground, Unclaimed Territory, Watery Grave)
- INFO `mana_pass`: the builder moved 12 cards of the mana base to bring the deck inside its power level
- INFO `cards_trimmed`: the list was 1 card over, so the builder cut this card: Harmonized Crescendo

Cards:

- 1 Blood Crypt
- 1 Breeding Pool
- 1 Godless Shrine
- 1 Hallowed Fountain
- 1 Overgrown Tomb
- 1 Steam Vents
- 1 Stomping Ground
- 1 Watery Grave
- 1 Command Tower
- 1 Base Camp
- 1 Cavern of Souls
- 1 Fire Nation Palace
- 1 Path of Ancestry
- 1 Secluded Courtyard
- 1 Unclaimed Territory
- 1 Three Tree City
- 1 Abundant Countryside
- 1 Accursed Duneyard
- 1 Dalkovan Encampment
- 1 Great Hall of Starnheim
- 1 Skemfar Elderhall
- 1 Gnottvold Slumbermound
- 5 Forest
- 1 Island
- 1 Plains
- 1 Swamp
- 1 Mountain
- 1 Chrome Mox
- 1 Dark Ritual
- 1 Lion's Eye Diamond
- 1 Lotus Petal
- 1 Mana Vault
- 1 Mox Amber
- 1 Mox Diamond
- 1 Mox Opal
- 1 Rite of Flame
- 1 Sol Ring
- 1 Springleaf Drum
- 1 Jeska's Will
- 1 Patchwork Banner
- 1 Pillar of Origins
- 1 Progenitor's Icon
- 1 Smothering Tithe
- 1 White Lotus Tile
- 1 Rhystic Study
- 1 Mindblade Render
- 1 Raiders' Spoils
- 1 Pact of the Serpent
- 1 Distant Melody
- 1 Species Specialist
- 1 Kutzil, Malamet Exemplar
- 1 Samut, Vizier of Naktamun
- 1 Felothar the Steadfast
- 1 For the Ancestors
- 1 An Offer You Can't Refuse
- 1 And They Shall Know No Fear
- 1 Ainok Strike Leader
- 1 Akiri, Fearless Voyager
- 1 Concerted Defense
- 1 Etchings of the Chosen
- 1 Fierce Guardianship
- 1 Galadhrim Ambush
- 1 Hakoda, Selfless Commander
- 1 Kindred Boon
- 1 Linvala, Shield of Sea Gate
- 1 Oko, Lorwyn Liege // Oko, Shadowmoor Scion
- 1 Steely Resolve
- 1 Vexilus Praetor
- 1 Orcish Bowmasters
- 1 Cacophony Scamp
- 1 Dreadhorde Butcher
- 1 Fleshbag Marauder
- 1 Merciless Executioner
- 1 Practiced Tactics
- 1 Combat Celebrant
- 1 Aspiring Champion
- 1 Aven Wind Guide
- 1 Cazur, Ruthless Stalker
- 1 Emeria Captain
- 1 Jazal Goldmane
- 1 Marisi, Breaker of the Coil
- 1 Raiyuu, Storm's Edge
- 1 Surrak, the Hunt Caller
- 1 Crippling Fear
- 1 Demonic Tutor
- 1 Enlightened Tutor
- 1 Mystical Tutor
- 1 Vampiric Tutor
- 1 Multiclass Baldric
- 1 Stadium Headliner
- 1 Invasion of New Capenna // Holy Frazzle-Cannon
- 1 Shang-Chi, Master of Kung Fu
- 1 Frodo, Sauron's Bane
- 1 Archpriest of Iona
- 1 Mirrex

