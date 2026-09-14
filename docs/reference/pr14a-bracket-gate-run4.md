# PR-14A bracket gate

Run date: 2026-09-14. Card snapshot: 2026-09-04.

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
| Prompt version | 13 |
| Calls | 12 |
| Cost | $0.4979 |
| Time | 424 seconds |

## Run

- Suite `bracket`, run `pr14a-bracket-gate-run4`, on 2026-09-14, commit `5cdc4cc`.
- Partial run over `10,11,12,13,14,15`. It reads the bars of its own items, and it never stands for the suite.
- Roles: generate on `gpt-5.6-terra` (openai, effort medium), judge on `claude-opus-5` (anthropic, effort medium), repair on `gpt-5.6-terra` (openai, effort medium).
- Versions: card snapshot 2026-09-04, generate prompt version 13.
- Calls: 12. Cost: $0.4979. Time: 424 seconds.

## Off-band features by key

Each row counts the decks whose feature sat outside its band after the repair turns.

| Feature | Decks off band |
|---|---|
| `color_sources` | 2 |

## Decks

### 10. Bracket 4, Urza, Lord High Artificer, artifacts

Pool 302. 99 cards, 0 block findings, repaired false.

Judge: bracket 4, agrees. Eight Game Changers plus Sol Ring, Lotus Petal, Mox Opal and Urza's Saga give it a fast-mana core that can deploy the commander turn two, and Urza with KCI/LED/Aetherflux Reservoir or Blasting Station gives real infinite-combo potential backed by free counterspells. That blows past Bracket 3's three-Game-Changer cap and early-combo restriction. It isn't cEDH, though: 31 basic Islands, thin tutoring, and filler like Twiddle, Unable to Scream and Lux Cannon mark it as an optimized artifact deck rather than a tuned competitive list.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 34 | 31 to 36 |  |  |
| `tapped_land` | 0 | 0 to 5 |  |  |
| `colorless_land` | 2 | 0 to 5 |  | Mishra's Workshop, Urza's Saga |
| `color_sources` | 1.38 | 0.9 or more |  | sources of requirement: U 35.75 of 26 |
| `avg_mana_value` | 3 | 1.6 to 3 |  | over 65 nonland cards |
| `ramp` | 13 | 10 to 16 |  |  |
| `draw` | 12 | 8 to 16 |  |  |
| `removal` | 10 | 6 to 14 |  |  |
| `wipe` | 2 | 0 to 5 |  |  |
| `interaction` | 11 | 6 to 16 |  |  |
| `tutor` | 4 | 2 or more |  | Arcum Dagsson, Tezzeret, Artifice Master, Transmute Artifact, Whir of Invention |
| `fast_mana` | 10 | 3 or more |  | Chrome Mox, Lion's Eye Diamond, Lotus Petal, Mana Vault, Moonsnare Prototype, Mox Amber, Mox Diamond, Mox Opal, Sol Ring, Springleaf Drum |
| `game_changer` | 8 | 4 or more |  | Chrome Mox, Fierce Guardianship, Force of Will, Lion's Eye Diamond, Mana Vault, Mishra's Workshop, Mox Diamond, Rhystic Study |
| `mana_turn_four` | 5.1 | 4.4 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.72 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.7 | -9 to 0.5 |  | the commander costs 4 and comes down on turn 3.3 on average |

Goldfish over 10000 hands: the commander on turn 3.3, 5.1 mana on turn four, and 0.72 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Chrome Mox, Fierce Guardianship, Force of Will, Lion's Eye Diamond, Mana Vault, Mishra's Workshop, Mox Diamond, Rhystic Study. Mass land denial none. Extra turns none. Combos 2.
- Aetherflux Reservoir + Tezzeret, Betrayer of Flesh (speed 5, two cards)
- Aetherflux Reservoir + Cyberdrive Awakener (speed 4, two cards)

Findings:

- INFO `curve_summary`: average mana value 3.00 over 65 nonland cards
- INFO `mana_pass`: the builder moved 5 cards of the mana base to bring the deck inside its power level

Cards:

- 31 Island
- 1 Mishra's Workshop
- 1 Otawara, Soaring City
- 1 Urza's Saga
- 1 Chrome Mox
- 1 Lion's Eye Diamond
- 1 Mana Vault
- 1 Lotus Petal
- 1 Mox Diamond
- 1 Mox Opal
- 1 Moonsnare Prototype
- 1 Springleaf Drum
- 1 Sol Ring
- 1 Metalworker
- 1 Krark-Clan Ironworks
- 1 Inspiring Statuary
- 1 The Mightstone and Weakstone
- 1 Rhystic Study
- 1 Thoughtcast
- 1 Reverse Engineer
- 1 Thirst for Knowledge
- 1 Sai, Master Thopterist
- 1 Vedalken Archmage
- 1 Forensic Gadgeteer
- 1 Riddlesmith
- 1 One with the Machine
- 1 Tezzeret, Artifice Master
- 1 Tezzeret, Betrayer of Flesh
- 1 Fierce Guardianship
- 1 Force of Will
- 1 Metallic Rebuke
- 1 Disruption Protocol
- 1 Ice Out
- 1 Stoic Rebuttal
- 1 Override
- 1 Padeem, Consul of Innovation
- 1 Shimmer Dragon
- 1 Welding Jar
- 1 Aether Spellbomb
- 1 Aetherflux Reservoir
- 1 Arcum Dagsson
- 1 Blasting Station
- 1 Contagion Clasp
- 1 Cyber Conversion
- 1 Dawnsire, Sunstar Dreadnought
- 1 Into Thin Air
- 1 Kitesail Larcenist
- 1 Lux Cannon
- 1 Engineered Explosives
- 1 Nevinyrral's Disk
- 1 Transmute Artifact
- 1 Whir of Invention
- 1 Clock of Omens
- 1 Etherium Sculptor
- 1 Foundry Inspector
- 1 Mystic Forge
- 1 Kappa Cannoneer
- 1 Cyberdrive Awakener
- 1 Master Transmuter
- 1 Phyrexian Metamorph
- 1 Karn, Scion of Urza
- 1 Threefold Thunderhulk
- 1 Traxos, Scourge of Kroog
- 1 Lodestone Golem
- 1 Darksteel Juggernaut
- 1 Mox Amber
- 1 Unable to Scream
- 1 Twiddle
- 1 Artificer's Assistant

### 11. Bracket 4, Korvold, Fae-Cursed King, treasure sacrifice

Pool 304. 99 cards, 0 block findings, repaired false.

Judge: bracket 4, agrees. Six marked Game Changers (Chrome Mox, Lion's Eye Diamond, Mana Vault, Mox Diamond, Jeska's Will, Orcish Bowmasters) blow well past Bracket 3's cap of three, and the fast-mana suite (Sol Ring, Lotus Petal, Mox Opal, Mox Amber) with tutors like Sidisi and Diabolic Intent supports very early starts. The sacrifice engine also assembles cheap loops — Chthonian Nightmare/Pitiless Plunderer with free sac outlets plus Goblin Bombardment or Vindictive Vampire as the kill — making this optimized rather than upgraded. It's still a themed Korvold aristocrats build with clunky cards, so not cEDH.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 33 | 31 to 36 |  |  |
| `tapped_land` | 0 | 0 to 5 |  |  |
| `colorless_land` | 2 | 0 to 5 |  | Dragon-Cursed Halls, The Gold Saucer |
| `color_sources` | 0.86 | 0.9 or more | yes | sources of requirement: B 16.25 of 19, R 16.25 of 19, G 16.25 of 19 |
| `avg_mana_value` | 2.77 | 1.6 to 3 |  | over 66 nonland cards |
| `ramp` | 14 | 10 to 16 |  |  |
| `draw` | 12 | 8 to 16 |  |  |
| `removal` | 10 | 6 to 14 |  |  |
| `wipe` | 2 | 0 to 5 |  |  |
| `interaction` | 11 | 6 to 16 |  |  |
| `tutor` | 2 | 2 or more |  | Diabolic Intent, Sidisi, Undead Vizier |
| `fast_mana` | 9 | 3 or more |  | Chrome Mox, Lion's Eye Diamond, Lotus Petal, Mana Vault, Mox Amber, Mox Diamond, Mox Opal, Sol Ring, Springleaf Drum |
| `game_changer` | 6 | 4 or more |  | Chrome Mox, Jeska's Will, Lion's Eye Diamond, Mana Vault, Mox Diamond, Orcish Bowmasters |
| `mana_turn_four` | 5.03 | 4.4 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.71 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.76 | -9 to 0.5 |  | the commander costs 5 and comes down on turn 4.24 on average |

Goldfish over 10000 hands: the commander on turn 4.24, 5.03 mana on turn four, and 0.71 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Chrome Mox, Jeska's Will, Lion's Eye Diamond, Mana Vault, Mox Diamond, Orcish Bowmasters. Mass land denial none. Extra turns none. Combos 1.
- Chthonian Nightmare + Pitiless Plunderer + Xorn (speed 5)

Findings:

- INFO `curve_summary`: average mana value 2.77 over 66 nonland cards
- WARN `profile_off_band`: the worst color's share of the sources it needs is 0.86, and bracket 4 wants 0.9 or more (sources of requirement: B 16.25 of 19, R 16.25 of 19, G 16.25 of 19)
- INFO `mana_pass`: the builder moved 3 cards of the mana base to bring the deck inside its power level

Cards:

- 10 Forest
- 10 Swamp
- 10 Mountain
- 1 Dragon-Cursed Halls
- 1 Heap Gate
- 1 The Gold Saucer
- 1 Chrome Mox
- 1 Lion's Eye Diamond
- 1 Lotus Petal
- 1 Mana Vault
- 1 Mox Diamond
- 1 Mox Opal
- 1 Sol Ring
- 1 Springleaf Drum
- 1 Pitiless Plunderer
- 1 Professional Face-Breaker
- 1 Captain Lannery Storm
- 1 Jeska's Will
- 1 Anje, Maid of Dishonor
- 1 Bushmeat Poacher
- 1 Dina, Essence Brewer
- 1 Disciple of Bolas
- 1 Gas Guzzler
- 1 Lord Skitter's Butcher
- 1 Relic Vial
- 1 Shadowheart, Dark Justiciar
- 1 Smothering Abomination
- 1 Vampire Gourmand
- 1 Skyclave Shadowcat
- 1 Stormclaw Rager
- 1 Constant Mists
- 1 Eldrazi Monument
- 1 Gift of Doom
- 1 Goblin Chirurgeon
- 1 Gyome, Master Chef
- 1 Kinzu of the Bleak Coven
- 1 Nightmare Shepherd
- 1 Rescue from the Underworld
- 1 Slobad, Goblin Tinkerer
- 1 Sylvan Safekeeper
- 1 Strefan, Maurer Progenitor
- 1 Attrition
- 1 Goblin Bombardment
- 1 Orcish Bowmasters
- 1 Broadside Bombardiers
- 1 Bone Shards
- 1 Eaten Alive
- 1 Fling
- 1 Grist, the Hunger Tide
- 1 Minsc & Boo, Timeless Heroes
- 1 Yawgmoth, Thran Physician
- 1 Academy Manufactor
- 1 Xorn
- 1 Viscera Seer
- 1 Woe Strider
- 1 Chthonian Nightmare
- 1 Diabolic Intent
- 1 Sidisi, Undead Vizier
- 1 Meren of Clan Nel Toth
- 1 Sadistic Hypnotist
- 1 Jarad, Golgari Lich Lord
- 1 Immersturm Predator
- 1 Flesh-Eater Imp
- 1 Vindictive Vampire
- 1 Vito's Inquisitor
- 1 Redcap Gutter-Dweller
- 1 Thallid Omnivore
- 1 The Meathook Massacre
- 1 Spiteful Banditry
- 1 Mox Amber
- 1 Treasure Map // Treasure Cove
- 1 Luck Bobblehead

### 12. Bracket 4, Prosper, Tome-Bound, treasure

Pool 210. 99 cards, 0 block findings, repaired false.

Judge: bracket 4, agrees. This is a heavily optimized build with ~11 Game Changers, a full fast-mana suite (LED, Mana Vault, Mox Diamond/Opal/Chrome, Lotus Petal, Dark Ritual, Sol Ring), and the best tutors (Vampiric, Imperial Seal, Demonic, Gamble) plus Ad Nauseam and Underworld Breach/LED/Wheel storm lines that can win well before turn six. That far exceeds Bracket 3's three-Game-Changer cap and combo restrictions. It stops short of 5 because it is still a Prosper treasure-theme deck with clunky payoffs like Smaug, Revel in Riches and assorted treasure creatures rather than a pure metagame-tuned cEDH list.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 33 | 31 to 36 |  |  |
| `tapped_land` | 5 | 0 to 5 |  | Canyon Slough, Foreboding Ruins, Rakdos Carnarium, Raucous Theater, Temple of Malice |
| `colorless_land` | 1 | 0 to 5 |  | Treasure Vault |
| `color_sources` | 1.42 | 0.9 or more |  | sources of requirement: B 33 of 19, R 27 of 19 |
| `avg_mana_value` | 2.39 | 1.6 to 3 |  | over 66 nonland cards |
| `ramp` | 13 | 10 to 16 |  |  |
| `draw` | 12 | 8 to 16 |  |  |
| `removal` | 10 | 6 to 14 |  |  |
| `wipe` | 2 | 0 to 5 |  |  |
| `interaction` | 10 | 6 to 16 |  |  |
| `tutor` | 4 | 2 or more |  | Demonic Tutor, Gamble, Imperial Seal, Vampiric Tutor |
| `fast_mana` | 10 | 3 or more |  | Chrome Mox, Dark Ritual, Lion's Eye Diamond, Lotus Petal, Mana Vault, Mox Diamond, Mox Opal, Rite of Flame, Sol Ring, Springleaf Drum |
| `game_changer` | 13 | 4 or more |  | Ad Nauseam, Chrome Mox, Demonic Tutor, Gamble, Imperial Seal, Jeska's Will, Lion's Eye Diamond, Mana Vault, Mox Diamond, Orcish Bowmasters, The One Ring, Underworld Breach, Vampiric Tutor |
| `mana_turn_four` | 4.96 | 4.4 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.71 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.64 | -9 to 0.5 |  | the commander costs 4 and comes down on turn 3.36 on average |

Goldfish over 10000 hands: the commander on turn 3.36, 4.96 mana on turn four, and 0.71 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Ad Nauseam, Chrome Mox, Demonic Tutor, Gamble, Imperial Seal, Jeska's Will, Lion's Eye Diamond, Mana Vault, Mox Diamond, Orcish Bowmasters, The One Ring, Underworld Breach, Vampiric Tutor. Mass land denial none. Extra turns none. Combos 4.
- Orcish Bowmasters + Peer into the Abyss (speed 3, two cards)
- Underworld Breach + Wheel of Fortune + Jeska's Will (speed 5)
- Prosper, Tome-Bound + Professional Face-Breaker + Sensei's Divining Top + Xorn (speed 5)
- Underworld Breach + Lion's Eye Diamond + Wheel of Fortune (speed 5)

Findings:

- INFO `curve_summary`: average mana value 2.39 over 66 nonland cards

Cards:

- 1 Badlands
- 1 Blazemire Verge
- 1 Blood Crypt
- 1 Canyon Slough
- 1 City of Brass
- 1 Command Tower
- 1 Dragonskull Summit
- 1 Exotic Orchard
- 1 Foreboding Ruins
- 1 Gemstone Caverns
- 1 Haunted Ridge
- 1 Luxury Suite
- 1 Mana Confluence
- 1 Raucous Theater
- 1 Rakdos Carnarium
- 1 Reflecting Pool
- 1 Shadowblood Ridge
- 1 Smoldering Marsh
- 1 Spire of Industry
- 1 Sulfurous Springs
- 1 Tainted Peak
- 1 Temple of Malice
- 1 Treasure Vault
- 8 Swamp
- 2 Mountain
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
- 1 Rite of Flame
- 1 Pitiless Plunderer
- 1 Professional Face-Breaker
- 1 Ad Nauseam
- 1 Corrupted Conviction
- 1 Demand Answers
- 1 Faithless Looting
- 1 Mishra's Bauble
- 1 Night's Whisper
- 1 Peer into the Abyss
- 1 Phyrexian Arena
- 1 Reckless Lackey
- 1 Sensei's Divining Top
- 1 Skullclamp
- 1 Wheel of Fortune
- 1 Abrade
- 1 Chaos Warp
- 1 Deadly Rollick
- 1 Feed the Swarm
- 1 Go for the Throat
- 1 Infernal Grasp
- 1 Orcish Bowmasters
- 1 Snuff Out
- 1 Terminate
- 1 Withering Torment
- 1 The One Ring
- 1 Hexing Squelcher
- 1 Kaya's Ghostform
- 1 Lightning Greaves
- 1 Mithril Coat
- 1 Not Dead After All
- 1 Pyroblast
- 1 Red Elemental Blast
- 1 Tibalt's Trickery
- 1 Vexing Bauble
- 1 Academy Manufactor
- 1 Contract Hero
- 1 Crossover Collaboration
- 1 Evereth, Viceroy of Plunder
- 1 Heroes for Hire
- 1 Xorn
- 1 Casey & Raph, Hotheads
- 1 Flamekin Gildweaver
- 1 Jaded Sell-Sword
- 1 Marut
- 1 Meticulous Artisan
- 1 Namazu Trader
- 1 Smaug the Impenetrable
- 1 Toxic Deluge
- 1 Vandalblast
- 1 Demonic Tutor
- 1 Gamble
- 1 Imperial Seal
- 1 Underworld Breach
- 1 Vampiric Tutor
- 1 Revel in Riches

### 13. Bracket 5, Kinnan, Bonder Prodigy, combo

Pool 303. 99 cards, 0 block findings, repaired false.

Judge: bracket 5, agrees. This is a Kinnan cEDH combo deck: full fetchland/Moxen fast-mana suite (LED, Mox Diamond, Chrome Mox, Grim Monolith, Mana Vault, Sol Ring), free counterspells (Force of Will, Fierce Guardianship), Mystical Tutor, and Thassa's Oracle as the wincon. It packs numerous two-card infinites (Kinnan + Basalt Monolith, Devoted Druid + untappers, Voltaic Construct/Filigree Sages loops) that can win by turn two or three. Nine Game Changers and a metagame-tuned interaction package place it firmly in cEDH.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 31 | 27 to 33 |  |  |
| `tapped_land` | 0 | 0 to 2 |  |  |
| `colorless_land` | 0 | 0 to 6 |  |  |
| `color_sources` | 1.02 | 0.9 or more |  | sources of requirement: U 23.5 of 23, G 24.75 of 19 |
| `avg_mana_value` | 2.35 | 1.2 to 2.4 |  | over 68 nonland cards |
| `ramp` | 16 | 12 to 22 |  |  |
| `draw` | 13 | 8 to 18 |  |  |
| `removal` | 9 | 4 to 14 |  |  |
| `wipe` | 1 | 0 to 3 |  |  |
| `interaction` | 15 | 8 to 22 |  |  |
| `tutor` | 4 | 4 or more |  | Arcum Dagsson, Chakram Retriever, Muddle the Mixture, Mystical Tutor |
| `fast_mana` | 10 | 6 or more |  | An Offer You Can't Refuse, Chrome Mox, Lion's Eye Diamond, Lotus Petal, Mana Vault, Mox Amber, Mox Diamond, Mox Opal, Sol Ring, Springleaf Drum |
| `game_changer` | 10 | 8 or more |  | Chrome Mox, Fierce Guardianship, Force of Will, Grim Monolith, Lion's Eye Diamond, Mana Vault, Mox Diamond, Mystical Tutor, Rhystic Study, Thassa's Oracle |
| `mana_turn_four` | 5.2 | 4.6 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.68 | 0.65 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.41 | -9 to 0.3 |  | the commander costs 2 and comes down on turn 1.59 on average |

Goldfish over 10000 hands: the commander on turn 1.59, 5.2 mana on turn four, and 0.68 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Chrome Mox, Fierce Guardianship, Force of Will, Grim Monolith, Lion's Eye Diamond, Mana Vault, Mox Diamond, Mystical Tutor, Rhystic Study, Thassa's Oracle. Mass land denial Natural Balance. Extra turns none. Combos 3.
- Kinnan, Bonder Prodigy + Grim Monolith + Spark Double (speed 5, two cards)
- Kinnan, Bonder Prodigy + Grim Monolith + Sakashima of a Thousand Faces (speed 5, two cards)
- Kinnan, Bonder Prodigy + Basalt Monolith (speed 5, two cards)

Findings:

- INFO `curve_summary`: average mana value 2.35 over 68 nonland cards
- INFO `mana_pass`: the builder moved 2 cards of the mana base to bring the deck inside its power level

Cards:

- 12 Forest
- 10 Island
- 1 Bloodstained Mire
- 1 Flooded Strand
- 1 Marsh Flats
- 1 Misty Rainforest
- 1 Polluted Delta
- 1 Scalding Tarn
- 1 Verdant Catacombs
- 1 Windswept Heath
- 1 Wooded Foothills
- 1 Chrome Mox
- 1 Lion's Eye Diamond
- 1 Lotus Petal
- 1 Mana Vault
- 1 Mox Diamond
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
- 1 Rhystic Study
- 1 Benefactor's Draught
- 1 Cerulean Wisps
- 1 Cloud of Faeries
- 1 Frantic Search
- 1 Kitsa, Otterball Elite
- 1 Lórien Revealed
- 1 Pip-Boy 3000
- 1 Refocus
- 1 Sewer-veillance Cam
- 1 Staff of Domination
- 1 Twitch
- 1 Yavimaya Elder
- 1 An Offer You Can't Refuse
- 1 Fierce Guardianship
- 1 Force of Will
- 1 Flusterstorm
- 1 Muddle the Mixture
- 1 Narset's Reversal
- 1 Shore Up
- 1 Sokka's Haiku
- 1 Unwind
- 1 Rewind
- 1 Katara's Reversal
- 1 Fleeting Reflection
- 1 Legolas's Quick Reflexes
- 1 Magic Damper
- 1 Pemmin's Aura
- 1 Aggressive Biomancy
- 1 Amphibian Downpour
- 1 Apathy
- 1 Bushwhack
- 1 Emerald Charm
- 1 Gleeful Sabotage
- 1 Provoke
- 1 Snap
- 1 Arcum Dagsson
- 1 Mystical Tutor
- 1 Brain Freeze
- 1 Filigree Sages
- 1 Chakram Retriever
- 1 Breaching Hippocamp
- 1 Cacophodon
- 1 Dross Scorpion
- 1 Peregrine Drake
- 1 Phyrexian Metamorph
- 1 Voltaic Construct
- 1 Spark Double
- 1 Sakashima of a Thousand Faces
- 1 Stunt Double
- 1 Natural Balance
- 1 Thassa's Oracle

### 14. Bracket 5, Yuriko, the Tiger's Shadow, ninjas

Pool 238. 99 cards, 0 block findings, repaired false.

Judge: bracket 5, agrees. This is a tuned cEDH Yuriko list: Underground Sea, a dozen pieces of fast mana including Lion's Eye Diamond, Mox Diamond, Mox Opal and Lotus Petal, a full suite of free counterspells (Force of Will, Fierce Guardianship, Pact of Negation, Mindbreak Trap) and a stack of unconditional tutors. It carries roughly a dozen marked Game Changers, far past Bracket 3's limit, and its ninjutsu clock plus Ad Nauseam/The One Ring draw engines are built to win in the first handful of turns. It plays the format's best strategy rather than a theme, which puts it in cEDH rather than merely Optimized.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 30 | 27 to 33 |  |  |
| `tapped_land` | 2 | 0 to 2 |  | Choked Estuary, Undercity Sewers |
| `colorless_land` | 0 | 0 to 6 |  |  |
| `color_sources` | 1.57 | 0.9 or more |  | sources of requirement: U 34.75 of 19, B 29.75 of 19 |
| `avg_mana_value` | 2.03 | 1.2 to 2.4 |  | over 69 nonland cards |
| `ramp` | 16 | 12 to 22 |  |  |
| `draw` | 13 | 8 to 18 |  |  |
| `removal` | 9 | 4 to 14 |  |  |
| `wipe` | 1 | 0 to 3 |  |  |
| `interaction` | 15 | 8 to 22 |  |  |
| `tutor` | 4 | 4 or more |  | Demonic Tutor, Imperial Seal, Mystical Tutor, Vampiric Tutor |
| `fast_mana` | 13 | 6 or more |  | An Offer You Can't Refuse, Chrome Mox, Culling the Weak, Dark Ritual, Everflowing Chalice, Lion's Eye Diamond, Lotus Petal, Mana Vault, Mox Amber, Mox Diamond, Mox Opal, Sol Ring, Springleaf Drum |
| `game_changer` | 13 | 8 or more |  | Ad Nauseam, Chrome Mox, Demonic Tutor, Fierce Guardianship, Force of Will, Imperial Seal, Lion's Eye Diamond, Mana Vault, Mox Diamond, Mystical Tutor, Rhystic Study, The One Ring, Vampiric Tutor |
| `mana_turn_four` | 5.17 | 4.6 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.66 | 0.65 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.7 | -9 to 0.3 |  | the commander costs 3 and comes down on turn 2.3 on average |

Goldfish over 10000 hands: the commander on turn 2.3, 5.17 mana on turn four, and 0.66 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Ad Nauseam, Chrome Mox, Demonic Tutor, Fierce Guardianship, Force of Will, Imperial Seal, Lion's Eye Diamond, Mana Vault, Mox Diamond, Mystical Tutor, Rhystic Study, The One Ring, Vampiric Tutor. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 2.03 over 69 nonland cards
- INFO `mana_pass`: the builder moved 3 cards of the mana base to bring the deck inside its power level

Cards:

- 1 Underground Sea
- 1 Watery Grave
- 1 Morphic Pool
- 1 Drowned Catacomb
- 1 Choked Estuary
- 1 Clearwater Pathway // Murkwater Pathway
- 1 Command Tower
- 1 City of Brass
- 1 Mana Confluence
- 1 Exotic Orchard
- 1 Forbidden Orchard
- 1 Gloomlake Verge
- 1 Shipwreck Marsh
- 1 Underground River
- 1 Sunken Ruins
- 1 Tainted Isle
- 1 Undercity Sewers
- 1 Cavern of Souls
- 1 Secluded Courtyard
- 1 Unclaimed Territory
- 1 Spire of Industry
- 7 Island
- 2 Swamp
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
- 1 Arcane Signet
- 1 Dimir Signet
- 1 Talisman of Dominance
- 1 Fellwar Stone
- 1 Ad Nauseam
- 1 Rhystic Study
- 1 Mystic Remora
- 1 Brainstorm
- 1 Gitaxian Probe
- 1 Ponder
- 1 Preordain
- 1 Opt
- 1 Sensei's Divining Top
- 1 Skullclamp
- 1 Moon-Circuit Hacker
- 1 Ingenious Infiltrator
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
- 1 Silver-Fur Master
- 1 Satoru Umezawa
- 1 Fallen Shinobi
- 1 Foot Mystic
- 1 Futurist Operative
- 1 Kami of Restless Shadows
- 1 Kotose, the Silent Spider
- 1 Moonblade Shinobi
- 1 Mukotai Ambusher
- 1 Okiba-Gang Shinobi
- 1 Sakashima's Student
- 1 Toxic Deluge
- 1 Demonic Tutor
- 1 Imperial Seal
- 1 Mystical Tutor
- 1 Vampiric Tutor

### 15. Bracket 5, Najeela, the Blade-Blossom, warriors combat

Pool 306. 99 cards, 0 block findings, repaired false.

Judge: bracket 4, disagrees. The deck is packed with Game Changers (Chrome Mox, Mana Vault, Mox Diamond, LED, Jeska's Will, Rhystic Study, Fierce Guardianship, Orcish Bowmasters, and three unconditional tutors) plus a dense fast-mana suite, and Najeela with Combat Celebrant is an early two-card infinite-combat kill, which pushes it well past Bracket 3's limits. However, the all-basic mana base with no duals or fetches and the warrior-tribal theme with filler creatures keep it short of true cEDH consistency and metagame focus.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 33 | 27 to 33 |  |  |
| `tapped_land` | 1 | 0 to 2 |  | Base Camp |
| `colorless_land` | 0 | 0 to 6 |  |  |
| `color_sources` | 0.59 | 0.9 or more | yes | sources of requirement: W 11.25 of 19, U 12.25 of 19, B 11.25 of 19, R 11.25 of 19, G 15.75 of 26 |
| `avg_mana_value` | 2.39 | 1.2 to 2.4 |  | over 66 nonland cards |
| `ramp` | 18 | 12 to 22 |  |  |
| `draw` | 11 | 8 to 18 |  |  |
| `removal` | 9 | 4 to 14 |  |  |
| `wipe` | 1 | 0 to 3 |  |  |
| `interaction` | 12 | 8 to 22 |  |  |
| `tutor` | 6 | 4 or more |  | Blaring Captain, Cazur, Ruthless Stalker, Demonic Tutor, Enlightened Tutor, Kassandra, Eagle Bearer, Mystical Tutor |
| `fast_mana` | 12 | 6 or more |  | An Offer You Can't Refuse, Chrome Mox, Dark Ritual, Lion's Eye Diamond, Lotus Petal, Mana Vault, Mox Amber, Mox Diamond, Mox Opal, Rite of Flame, Sol Ring, Springleaf Drum |
| `game_changer` | 11 | 8 or more |  | Chrome Mox, Demonic Tutor, Enlightened Tutor, Fierce Guardianship, Jeska's Will, Lion's Eye Diamond, Mana Vault, Mox Diamond, Mystical Tutor, Orcish Bowmasters, Rhystic Study |
| `mana_turn_four` | 5.12 | 4.6 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.71 | 0.65 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.66 | -9 to 0.3 |  | the commander costs 3 and comes down on turn 2.34 on average |

Goldfish over 10000 hands: the commander on turn 2.34, 5.12 mana on turn four, and 0.71 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Chrome Mox, Demonic Tutor, Enlightened Tutor, Fierce Guardianship, Jeska's Will, Lion's Eye Diamond, Mana Vault, Mox Diamond, Mystical Tutor, Orcish Bowmasters, Rhystic Study. Mass land denial none. Extra turns none. Combos 1.
- Najeela, the Blade-Blossom + Professional Face-Breaker (speed 5, two cards)

Findings:

- INFO `curve_summary`: average mana value 2.39 over 66 nonland cards
- WARN `profile_off_band`: the worst color's share of the sources it needs is 0.59, and bracket 5 wants 0.9 or more (sources of requirement: W 11.25 of 19, U 12.25 of 19, B 11.25 of 19, R 11.25 of 19, G 15.75 of 26)
- INFO `mana_pass`: the builder moved 9 cards of the mana base to bring the deck inside its power level

Cards:

- 11 Forest
- 6 Island
- 5 Mountain
- 5 Plains
- 5 Swamp
- 1 Base Camp
- 1 Chrome Mox
- 1 Mana Vault
- 1 Mox Diamond
- 1 Lion's Eye Diamond
- 1 Lotus Petal
- 1 Mox Amber
- 1 Mox Opal
- 1 Sol Ring
- 1 Springleaf Drum
- 1 Dark Ritual
- 1 Rite of Flame
- 1 Jeska's Will
- 1 Ragavan, Nimble Pilferer
- 1 Professional Face-Breaker
- 1 Heronblade Elite
- 1 Ardent Electromancer
- 1 Rhystic Study
- 1 Kassandra, Eagle Bearer
- 1 Mindblade Render
- 1 Azra Oddsmaker
- 1 Akki Ronin
- 1 Estinien Varlineau
- 1 Felothar the Steadfast
- 1 Ivora, Insatiable Heir
- 1 Jaxis, the Troublemaker
- 1 Kutzil, Malamet Exemplar
- 1 Samut, Vizier of Naktamun
- 1 An Offer You Can't Refuse
- 1 Fierce Guardianship
- 1 Concerted Defense
- 1 Akiri, Fearless Voyager
- 1 Eladamri, Lord of Leaves
- 1 Ezuri, Renegade Leader
- 1 Hakoda, Selfless Commander
- 1 Kashi-Tribe Elite
- 1 Linvala, Shield of Sea Gate
- 1 Rhys, the Evermore
- 1 Selfless Samurai
- 1 Winota, Joiner of Forces
- 1 Orcish Bowmasters
- 1 Cacophony Scamp
- 1 Dreadhorde Butcher
- 1 Goblin Cratermaker
- 1 Practiced Tactics
- 1 Chatterfang, Squirrel General
- 1 Fleshbag Marauder
- 1 Merciless Executioner
- 1 Sosuke, Son of Seshiro
- 1 Combat Celebrant
- 1 Allied Assault
- 1 Blaring Captain
- 1 Cazur, Ruthless Stalker
- 1 Aspiring Champion
- 1 Bull-Rush Bruiser
- 1 Kabira Outrider
- 1 Kimahri, Valiant Guardian
- 1 Marisi, Breaker of the Coil
- 1 Raiyuu, Storm's Edge
- 1 Surrak, the Hunt Caller
- 1 The Destined Warrior
- 1 Goblin Chainwhirler
- 1 Demonic Tutor
- 1 Enlightened Tutor
- 1 Mystical Tutor
- 1 Shang-Chi, Master of Kung Fu
- 1 Rograkh, Son of Rohgahh

