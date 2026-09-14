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
| Cost | $0.5247 |
| Time | 442 seconds |

## Run

- Suite `bracket`, run `pr14a-bracket-gate-run5`, on 2026-09-14, commit `f45791a`.
- Partial run over `10,11,12,13,14,15`. It reads the bars of its own items, and it never stands for the suite.
- Roles: generate on `gpt-5.6-terra` (openai, effort medium), judge on `claude-opus-5` (anthropic, effort medium), repair on `gpt-5.6-terra` (openai, effort medium).
- Versions: card snapshot 2026-09-04, generate prompt version 13.
- Calls: 12. Cost: $0.5247. Time: 442 seconds.

## Off-band features by key

Each row counts the decks whose feature sat outside its band after the repair turns.

| Feature | Decks off band |
|---|---|
| `hands_two_to_four_lands` | 1 |
| `color_sources` | 1 |

## Decks

### 10. Bracket 4, Urza, Lord High Artificer, artifacts

Pool 317. 99 cards, 0 block findings, repaired false.

Judge: bracket 4, agrees. Eight marked Game Changers (Mishra's Workshop, Chrome Mox, LED, Mana Vault, Mox Diamond, Rhystic Study, Force of Will, Fierce Guardianship) plus Sol Ring, Lotus Petal, Mox Opal and Metalworker make this a fast-mana artifact deck that can deploy Urza and go off well before turn six, backed by free counterspells. Combo lines like Urza + Krark-Clan Ironworks with zero-cost artifacts into Aetherflux Reservoir push it past Upgraded. It is not quite cEDH, though, since the mana base is 24 basic Islands with no real tutor package and plenty of durdly artifact-theme filler, so Bracket 4 Optimized fits.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 34 | 31 to 36 |  |  |
| `tapped_land` | 0 | 0 to 5 |  |  |
| `colorless_land` | 5 | 0 to 5 |  | Academy Ruins, Buried Ruin, Inventors' Fair, Mishra's Workshop, Urza's Saga |
| `color_sources` | 1.42 | 0.9 or more |  | sources of requirement: U 32.75 of 23 |
| `avg_mana_value` | 2.98 | 1.6 to 3 |  | over 65 nonland cards |
| `ramp` | 13 | 10 to 16 |  |  |
| `draw` | 12 | 8 to 16 |  |  |
| `removal` | 10 | 6 to 14 |  |  |
| `wipe` | 2 | 0 to 5 |  |  |
| `interaction` | 10 | 6 to 16 |  |  |
| `tutor` | 2 | 2 or more |  | Arcum Dagsson, Tezzeret, Artifice Master |
| `fast_mana` | 11 | 3 or more |  | An Offer You Can't Refuse, Chrome Mox, Lion's Eye Diamond, Lotus Petal, Mana Vault, Moonsnare Prototype, Mox Amber, Mox Diamond, Mox Opal, Sol Ring, Springleaf Drum |
| `game_changer` | 8 | 4 or more |  | Chrome Mox, Fierce Guardianship, Force of Will, Lion's Eye Diamond, Mana Vault, Mishra's Workshop, Mox Diamond, Rhystic Study |
| `mana_turn_four` | 5.31 | 4.4 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.72 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.9 | -9 to 0.5 |  | the commander costs 4 and comes down on turn 3.1 on average |

Goldfish over 10000 hands: the commander on turn 3.1, 5.31 mana on turn four, and 0.72 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Chrome Mox, Fierce Guardianship, Force of Will, Lion's Eye Diamond, Mana Vault, Mishra's Workshop, Mox Diamond, Rhystic Study. Mass land denial none. Extra turns none. Combos 1.
- Aetherflux Reservoir + Cyberdrive Awakener (speed 4, two cards)

Findings:

- INFO `curve_summary`: average mana value 2.98 over 65 nonland cards
- INFO `mana_pass`: the builder moved 4 cards of the mana base to bring the deck inside its power level

Cards:

- 24 Island
- 1 Seat of the Synod
- 1 Mystic Sanctuary
- 1 Otawara, Soaring City
- 1 Urza's Saga
- 1 Mishra's Workshop
- 1 Inventors' Fair
- 1 Academy Ruins
- 1 Buried Ruin
- 1 Spire of Industry
- 1 Glimmervoid
- 1 Chrome Mox
- 1 Lion's Eye Diamond
- 1 Mana Vault
- 1 Lotus Petal
- 1 Mox Diamond
- 1 Mox Opal
- 1 Moonsnare Prototype
- 1 Sol Ring
- 1 Springleaf Drum
- 1 Metalworker
- 1 Krark-Clan Ironworks
- 1 Automated Artificer
- 1 Vedalken Engineer
- 1 Rhystic Study
- 1 Thoughtcast
- 1 Sai, Master Thopterist
- 1 Reverse Engineer
- 1 Thirst for Knowledge
- 1 Riddlesmith
- 1 Vedalken Archmage
- 1 Nexus of Becoming
- 1 Tezzeret, Artifice Master
- 1 Forensic Gadgeteer
- 1 Trading Post
- 1 Force of Will
- 1 Fierce Guardianship
- 1 An Offer You Can't Refuse
- 1 Metallic Rebuke
- 1 Disruption Protocol
- 1 Stoic Rebuttal
- 1 Ice Out
- 1 Welding Jar
- 1 Padeem, Consul of Innovation
- 1 Etched Champion
- 1 Aether Spellbomb
- 1 Aetherflux Reservoir
- 1 Arcum Dagsson
- 1 Contagion Clasp
- 1 Cyber Conversion
- 1 Dawnsire, Sunstar Dreadnought
- 1 Into Thin Air
- 1 Kitesail Larcenist
- 1 Ravenform
- 1 Resculpt
- 1 Emry, Lurker of the Loch
- 1 Foundry Inspector
- 1 Etherium Sculptor
- 1 Mystic Forge
- 1 Clock of Omens
- 1 Unwinding Clock
- 1 Kappa Cannoneer
- 1 Master Transmuter
- 1 Phyrexian Metamorph
- 1 Myr Enforcer
- 1 Sojourner's Companion
- 1 Karn, Scion of Urza
- 1 Cyberdrive Awakener
- 1 Traxos, Scourge of Kroog
- 1 Threefold Thunderhulk
- 1 Jhoira's Familiar
- 1 Engineered Explosives
- 1 Hurkyl's Recall
- 1 Mox Amber
- 1 Unable to Scream
- 1 Twiddle

### 11. Bracket 4, Korvold, Fae-Cursed King, treasure sacrifice

Pool 326. 99 cards, 0 block findings, repaired false.

Judge: bracket 4, agrees. Six marked Game Changers (Chrome Mox, Jeska's Will, Lion's Eye Diamond, Mana Vault, Mox Diamond, Orcish Bowmasters) already push it past Bracket 3's limit of three. The mana base is an optimized dual/fast-mana shell with Sol Ring, Lotus Petal, Dark Ritual and multiple Moxen, and the sacrifice package contains cheap infinite loops (Chthonian Nightmare/Pitiless Plunderer, Carrion Feeder plus Blasting Station or Zulaport Cutthroat drains) available well before turn six. It is a highly tuned Korvold aristocrats deck rather than a metagame-defined cEDH list, so Bracket 4 fits.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 34 | 31 to 36 |  |  |
| `tapped_land` | 0 | 0 to 5 |  |  |
| `colorless_land` | 0 | 0 to 5 |  |  |
| `color_sources` | 1.08 | 0.9 or more |  | sources of requirement: B 24.75 of 23, R 21.75 of 19, G 20.75 of 19 |
| `avg_mana_value` | 2.95 | 1.6 to 3 |  | over 65 nonland cards |
| `ramp` | 13 | 10 to 16 |  |  |
| `draw` | 12 | 8 to 16 |  |  |
| `removal` | 10 | 6 to 14 |  |  |
| `wipe` | 2 | 0 to 5 |  |  |
| `interaction` | 10 | 6 to 16 |  |  |
| `tutor` | 2 | 2 or more |  | Diabolic Intent, Magda, Brazen Outlaw |
| `fast_mana` | 10 | 3 or more |  | Chrome Mox, Dark Ritual, Lion's Eye Diamond, Lotus Petal, Mana Vault, Mox Amber, Mox Diamond, Mox Opal, Sol Ring, Springleaf Drum |
| `game_changer` | 6 | 4 or more |  | Chrome Mox, Jeska's Will, Lion's Eye Diamond, Mana Vault, Mox Diamond, Orcish Bowmasters |
| `mana_turn_four` | 4.93 | 4.4 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.72 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.66 | -9 to 0.5 |  | the commander costs 5 and comes down on turn 4.34 on average |

Goldfish over 10000 hands: the commander on turn 4.34, 4.93 mana on turn four, and 0.72 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Chrome Mox, Jeska's Will, Lion's Eye Diamond, Mana Vault, Mox Diamond, Orcish Bowmasters. Mass land denial none. Extra turns none. Combos 1.
- Chthonian Nightmare + Pitiless Plunderer + Xorn (speed 5)

Findings:

- INFO `curve_summary`: average mana value 2.95 over 65 nonland cards
- INFO `mana_pass`: the builder moved 2 cards of the mana base to bring the deck inside its power level

Cards:

- 1 Blood Crypt
- 1 Cavern of Souls
- 1 City of Brass
- 1 Command Tower
- 1 Dragonskull Summit
- 5 Forest
- 1 Exotic Orchard
- 1 Llanowar Wastes
- 1 Luxury Suite
- 1 Mana Confluence
- 6 Mountain
- 1 Overgrown Tomb
- 1 Phyrexian Tower
- 1 Reflecting Pool
- 1 Rootbound Crag
- 1 Stomping Ground
- 6 Swamp
- 1 Sulfurous Springs
- 1 Undergrowth Stadium
- 1 Woodland Cemetery
- 1 Chrome Mox
- 1 Dark Ritual
- 1 Jeska's Will
- 1 Lion's Eye Diamond
- 1 Lotus Petal
- 1 Magda, Brazen Outlaw
- 1 Mana Vault
- 1 Mox Diamond
- 1 Mox Opal
- 1 Pitiless Plunderer
- 1 Professional Face-Breaker
- 1 Sol Ring
- 1 Springleaf Drum
- 1 Anje, Maid of Dishonor
- 1 Beetle-Headed Merchants
- 1 Bushmeat Poacher
- 1 Dina, Essence Brewer
- 1 Disciple of Bolas
- 1 Gas Guzzler
- 1 Greven, Predator Captain
- 1 Izoni, Thousand-Eyed
- 1 Shadowheart, Dark Justiciar
- 1 Smothering Abomination
- 1 Tevesh Szat, Doom of Fools
- 1 Vampire Gourmand
- 1 Attrition
- 1 Blasting Station
- 1 Bone Shards
- 1 Broadside Bombardiers
- 1 Goblin Bombardment
- 1 Grist, the Hunger Tide
- 1 Minsc & Boo, Timeless Heroes
- 1 Orcish Bowmasters
- 1 Yawgmoth, Thran Physician
- 1 Zulaport Cutthroat
- 1 Academy Manufactor
- 1 Carrion Feeder
- 1 Diabolic Intent
- 1 Witch's Oven
- 1 Xorn
- 1 Chthonian Nightmare
- 1 Ghoulcaller Gisa
- 1 Hell's Caretaker
- 1 Kethek, Crucible Goliath
- 1 Meren of Clan Nel Toth
- 1 Prossh, Skyraider of Kher
- 1 Rakdos, the Muscle
- 1 Sadistic Hypnotist
- 1 Thallid Omnivore
- 1 Undercity Scavenger
- 1 Venerated Stormsinger
- 1 Constant Mists
- 1 Eldrazi Monument
- 1 Goblin Chirurgeon
- 1 Gyome, Master Chef
- 1 Kinzu of the Bleak Coven
- 1 My Precious // Allure of Power
- 1 Nightmare Shepherd
- 1 Rescue from the Underworld
- 1 Slobad, Goblin Tinkerer
- 1 Sylvan Safekeeper
- 1 Liliana, Dreadhorde General
- 1 The Meathook Massacre
- 1 Revel in Riches
- 1 Mox Amber

### 12. Bracket 4, Prosper, Tome-Bound, treasure

Pool 210. 99 cards, 0 block findings, repaired false.

Judge: bracket 4, agrees. Ten Game Changers plus a dense fast-mana suite (Sol Ring, Mana Vault, three Moxen, Lotus Petal, LED, Dark Ritual, Jeska's Will), black tutors, and Ad Nauseam/The One Ring card advantage put this far past Upgraded's three-Game-Changer cap. It has genuine turn-three-or-four kill potential via Underworld Breach + Lion's Eye Diamond with Wheel/Jeska's Will, backed by free interaction like Deadly Rollick, Snuff Out and Deflecting Swat. It stops short of 5 because the shell is still built around a Prosper treasure theme with fillers like Smaug, Namazu Trader and Revel in Riches rather than the tightest competitive line.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 33 | 31 to 36 |  |  |
| `tapped_land` | 3 | 0 to 5 |  | Canyon Slough, Foreboding Ruins, Raucous Theater |
| `colorless_land` | 1 | 0 to 5 |  | Volatile Fault |
| `color_sources` | 1.66 | 0.9 or more |  | sources of requirement: B 31.5 of 19, R 32.5 of 19 |
| `avg_mana_value` | 2.36 | 1.6 to 3 |  | over 66 nonland cards |
| `ramp` | 13 | 10 to 16 |  |  |
| `draw` | 12 | 8 to 16 |  |  |
| `removal` | 12 | 6 to 14 |  |  |
| `wipe` | 2 | 0 to 5 |  |  |
| `interaction` | 8 | 6 to 16 |  |  |
| `tutor` | 2 | 2 or more |  | Demonic Tutor, Vampiric Tutor |
| `fast_mana` | 10 | 3 or more |  | Chrome Mox, Dark Ritual, Lion's Eye Diamond, Lotus Petal, Mana Vault, Mox Amber, Mox Diamond, Mox Opal, Sol Ring, Springleaf Drum |
| `game_changer` | 11 | 4 or more |  | Ad Nauseam, Chrome Mox, Demonic Tutor, Jeska's Will, Lion's Eye Diamond, Mana Vault, Mox Diamond, Orcish Bowmasters, The One Ring, Underworld Breach, Vampiric Tutor |
| `mana_turn_four` | 5.13 | 4.4 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.71 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.8 | -9 to 0.5 |  | the commander costs 4 and comes down on turn 3.2 on average |

Goldfish over 10000 hands: the commander on turn 3.2, 5.13 mana on turn four, and 0.71 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Ad Nauseam, Chrome Mox, Demonic Tutor, Jeska's Will, Lion's Eye Diamond, Mana Vault, Mox Diamond, Orcish Bowmasters, The One Ring, Underworld Breach, Vampiric Tutor. Mass land denial none. Extra turns none. Combos 3.
- Underworld Breach + Wheel of Fortune + Jeska's Will (speed 5)
- Prosper, Tome-Bound + Professional Face-Breaker + Sensei's Divining Top + Xorn (speed 5)
- Underworld Breach + Lion's Eye Diamond + Wheel of Fortune (speed 5)

Findings:

- INFO `curve_summary`: average mana value 2.36 over 66 nonland cards

Cards:

- 1 Badlands
- 1 Blood Crypt
- 1 Blazemire Verge
- 1 Canyon Slough
- 1 Cavern of Souls
- 1 City of Brass
- 1 Command Tower
- 1 Dragonskull Summit
- 1 Exotic Orchard
- 1 Forbidden Orchard
- 1 Foreboding Ruins
- 1 Graven Cairns
- 1 Haunted Ridge
- 1 Luxury Suite
- 1 Mana Confluence
- 1 Plaza of Heroes
- 1 Raucous Theater
- 1 Reflecting Pool
- 1 Shadowblood Ridge
- 1 Smoldering Marsh
- 1 Spire of Industry
- 1 Sulfurous Springs
- 1 Tainted Peak
- 1 Volatile Fault
- 5 Mountain
- 4 Swamp
- 1 Chrome Mox
- 1 Dark Ritual
- 1 Jeska's Will
- 1 Lion's Eye Diamond
- 1 Lotus Petal
- 1 Mana Vault
- 1 Mox Amber
- 1 Mox Diamond
- 1 Mox Opal
- 1 Sol Ring
- 1 Springleaf Drum
- 1 Pitiless Plunderer
- 1 Professional Face-Breaker
- 1 Ad Nauseam
- 1 Corrupted Conviction
- 1 Demand Answers
- 1 Faithless Looting
- 1 Idol of Oblivion
- 1 Mishra's Bauble
- 1 Night's Whisper
- 1 Phyrexian Arena
- 1 Reckless Lackey
- 1 Sensei's Divining Top
- 1 Skullclamp
- 1 Wheel of Fortune
- 1 Pyroblast
- 1 Red Elemental Blast
- 1 The One Ring
- 1 Tibalt's Trickery
- 1 Vexing Bauble
- 1 Lightning Greaves
- 1 Swiftfoot Boots
- 1 Not Dead After All
- 1 Abrade
- 1 Bedevil
- 1 Chaos Warp
- 1 Deadly Rollick
- 1 Dismember
- 1 Feed the Swarm
- 1 Go for the Throat
- 1 Orcish Bowmasters
- 1 Rakdos Charm
- 1 Snuff Out
- 1 Terminate
- 1 Withering Torment
- 1 Academy Manufactor
- 1 Gilded Assault Cart
- 1 Heroes for Hire
- 1 Ticket Tortoise
- 1 Wanted Scoundrels
- 1 Xorn
- 1 Casey & Raph, Hotheads
- 1 Flamekin Gildweaver
- 1 Jaded Sell-Sword
- 1 Marut
- 1 Namazu Trader
- 1 Smaug the Impenetrable
- 1 Meticulous Artisan
- 1 Toxic Deluge
- 1 Vandalblast
- 1 Demonic Tutor
- 1 Vampiric Tutor
- 1 Deflecting Swat
- 1 Underworld Breach
- 1 Revel in Riches
- 1 Luck Bobblehead

### 13. Bracket 5, Kinnan, Bonder Prodigy, combo

Pool 322. 99 cards, 0 block findings, repaired false.

Judge: bracket 5, agrees. This is a textbook cEDH Kinnan build: heavy fast mana (LED, Moxen, Grim Monolith, Mana Vault, Lotus Petal, Sol Ring), a fetch/dual-heavy blue-green base, free counterspells like Fierce Guardianship, and ten-plus Game Changers. It wins with Thassa's Oracle/Lab Man/Jace off infinite mana lines (Kinnan + Basalt Monolith, Devoted Druid combos, Voltaic Construct/Filigree Sages untappers), all tutorable and assemblable in the first few turns. It is metagame-optimized rather than themed, so it plays at the competitive tier.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 30 | 27 to 33 |  |  |
| `tapped_land` | 0 | 0 to 2 |  |  |
| `colorless_land` | 0 | 0 to 6 |  |  |
| `color_sources` | 1.36 | 0.9 or more |  | sources of requirement: U 25.75 of 19, G 28.75 of 19 |
| `avg_mana_value` | 2.19 | 1.2 to 2.4 |  | over 69 nonland cards |
| `ramp` | 17 | 12 to 22 |  |  |
| `draw` | 13 | 8 to 18 |  |  |
| `removal` | 9 | 4 to 14 |  |  |
| `wipe` | 1 | 0 to 3 |  |  |
| `interaction` | 13 | 8 to 22 |  |  |
| `tutor` | 5 | 4 or more |  | Arcum Dagsson, Chakram Retriever, Muddle the Mixture, Mystical Tutor, Worldly Tutor |
| `fast_mana` | 10 | 6 or more |  | An Offer You Can't Refuse, Chrome Mox, Lion's Eye Diamond, Lotus Petal, Mana Vault, Mox Amber, Mox Diamond, Mox Opal, Sol Ring, Springleaf Drum |
| `game_changer` | 11 | 8 or more |  | Chrome Mox, Cyclonic Rift, Fierce Guardianship, Grim Monolith, Lion's Eye Diamond, Mana Vault, Mox Diamond, Mystical Tutor, Rhystic Study, Thassa's Oracle, Worldly Tutor |
| `mana_turn_four` | 5.08 | 4.6 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.66 | 0.65 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.42 | -9 to 0.3 |  | the commander costs 2 and comes down on turn 1.58 on average |

Goldfish over 10000 hands: the commander on turn 1.58, 5.08 mana on turn four, and 0.66 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Chrome Mox, Cyclonic Rift, Fierce Guardianship, Grim Monolith, Lion's Eye Diamond, Mana Vault, Mox Diamond, Mystical Tutor, Rhystic Study, Thassa's Oracle, Worldly Tutor. Mass land denial none. Extra turns none. Combos 6.
- Ioreth of the Healing House + Kiora's Follower (speed 5)
- Ioreth of the Healing House + Kelpie Guide (speed 5)
- Sakashima of a Thousand Faces + Ioreth of the Healing House (speed 5, two cards)
- Kinnan, Bonder Prodigy + Grim Monolith + Spark Double (speed 5, two cards)
- Kinnan, Bonder Prodigy + Grim Monolith + Sakashima of a Thousand Faces (speed 5, two cards)
- Kinnan, Bonder Prodigy + Basalt Monolith (speed 5, two cards)

Findings:

- INFO `curve_summary`: average mana value 2.19 over 69 nonland cards
- INFO `mana_pass`: the builder moved 3 cards of the mana base to bring the deck inside its power level

Cards:

- 1 Fabled Passage
- 1 Evolving Wilds
- 1 Flooded Strand
- 1 Misty Rainforest
- 1 Polluted Delta
- 1 Scalding Tarn
- 1 Verdant Catacombs
- 1 Windswept Heath
- 1 Wooded Foothills
- 5 Island
- 8 Forest
- 1 Breeding Pool
- 1 Command Tower
- 1 City of Brass
- 1 Exotic Orchard
- 1 Mana Confluence
- 1 Reflecting Pool
- 1 Rejuvenating Springs
- 1 Yavimaya Coast
- 1 Chrome Mox
- 1 Lion's Eye Diamond
- 1 Lotus Petal
- 1 Mana Vault
- 1 Mox Amber
- 1 Mox Diamond
- 1 Mox Opal
- 1 Sol Ring
- 1 Springleaf Drum
- 1 Grim Monolith
- 1 Basalt Monolith
- 1 Arbor Elf
- 1 Devoted Druid
- 1 Earthcraft
- 1 Kiora's Follower
- 1 Ioreth of the Healing House
- 1 Kelpie Guide
- 1 Rhystic Study
- 1 Benefactor's Draught
- 1 Cerulean Wisps
- 1 Cloud of Faeries
- 1 Frantic Search
- 1 Innocuous Researcher
- 1 Kitsa, Otterball Elite
- 1 Pip-Boy 3000
- 1 Refocus
- 1 Sewer-veillance Cam
- 1 Silent Hallcreeper
- 1 Twitch
- 1 Yavimaya Elder
- 1 An Offer You Can't Refuse
- 1 Fierce Guardianship
- 1 Flusterstorm
- 1 Muddle the Mixture
- 1 Narset's Reversal
- 1 Shore Up
- 1 Unwind
- 1 Magic Damper
- 1 Katara's Reversal
- 1 Legolas's Quick Reflexes
- 1 Biosynthic Burst
- 1 Octopus Form
- 1 Fleeting Reflection
- 1 Arcum Dagsson
- 1 Apathy
- 1 Bushwhack
- 1 Emerald Charm
- 1 Gleeful Sabotage
- 1 Provoke
- 1 Snap
- 1 Vedalken Anatomist
- 1 Singing Bell Strike
- 1 Mystical Tutor
- 1 Worldly Tutor
- 1 Filigree Sages
- 1 Voltaic Construct
- 1 Peregrine Drake
- 1 Phyrexian Metamorph
- 1 Sakashima of a Thousand Faces
- 1 Sakashima the Impostor
- 1 Sakashima's Student
- 1 Spark Double
- 1 Stunt Double
- 1 Chakram Retriever
- 1 Dross Scorpion
- 1 Cyclonic Rift
- 1 Thassa's Oracle
- 1 Laboratory Maniac
- 1 Jace, Wielder of Mysteries

### 14. Bracket 5, Yuriko, the Tiger's Shadow, ninjas

Pool 238. 99 cards, 0 block findings, repaired false.

Judge: bracket 5, agrees. This is a textbook cEDH Yuriko list: Underground Sea and a full fast-mana suite (LED, Chrome Mox, Mox Diamond/Opal/Amber, Lotus Petal, Dark Ritual, Culling the Weak), a dozen-plus free and cheap counterspells, and a Tainted Pact/Demonic Consultation-style Thassa's Oracle win with a deep tutor package. It carries roughly fourteen Game Changers and is built purely to win as fast as possible rather than around a theme.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 27 | 27 to 33 |  |  |
| `tapped_land` | 2 | 0 to 2 |  | Choked Estuary, Undercity Sewers |
| `colorless_land` | 0 | 0 to 6 |  |  |
| `color_sources` | 1.74 | 0.9 or more |  | sources of requirement: U 33 of 19, B 33 of 19 |
| `avg_mana_value` | 2.28 | 1.2 to 2.4 |  | over 72 nonland cards |
| `ramp` | 15 | 12 to 22 |  |  |
| `draw` | 13 | 8 to 18 |  |  |
| `removal` | 9 | 4 to 14 |  |  |
| `wipe` | 1 | 0 to 3 |  |  |
| `interaction` | 14 | 8 to 22 |  |  |
| `tutor` | 5 | 4 or more |  | Demonic Tutor, Higure, the Still Wind, Imperial Seal, Mystical Tutor, Vampiric Tutor |
| `fast_mana` | 13 | 6 or more |  | An Offer You Can't Refuse, Chrome Mox, Culling the Weak, Dark Ritual, Everflowing Chalice, Lion's Eye Diamond, Lotus Petal, Mana Vault, Mox Amber, Mox Diamond, Mox Opal, Sol Ring, Springleaf Drum |
| `game_changer` | 16 | 8 or more |  | Ad Nauseam, Chrome Mox, Cyclonic Rift, Demonic Tutor, Fierce Guardianship, Force of Will, Imperial Seal, Lion's Eye Diamond, Mana Vault, Mox Diamond, Mystical Tutor, Orcish Bowmasters, Rhystic Study, Thassa's Oracle, The One Ring, Vampiric Tutor |
| `mana_turn_four` | 4.92 | 4.6 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.6 | 0.65 or more | yes | share of first seven-card hands |
| `commander_turn_over_mv` | -0.63 | -9 to 0.3 |  | the commander costs 3 and comes down on turn 2.37 on average |

Goldfish over 10000 hands: the commander on turn 2.37, 4.92 mana on turn four, and 0.6 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Ad Nauseam, Chrome Mox, Cyclonic Rift, Demonic Tutor, Fierce Guardianship, Force of Will, Imperial Seal, Lion's Eye Diamond, Mana Vault, Mox Diamond, Mystical Tutor, Orcish Bowmasters, Rhystic Study, Thassa's Oracle, The One Ring, Vampiric Tutor. Mass land denial none. Extra turns none. Combos 1.
- Tainted Pact + Thassa's Oracle (speed 4, two cards)

Findings:

- INFO `curve_summary`: average mana value 2.28 over 72 nonland cards
- WARN `profile_off_band`: the share of opening hands with two to four lands is 0.6, and bracket 5 wants 0.65 or more (share of first seven-card hands)

Cards:

- 1 Cavern of Souls
- 1 Choked Estuary
- 1 City of Brass
- 1 Clearwater Pathway // Murkwater Pathway
- 1 Command Tower
- 1 Darkwater Catacombs
- 1 Drowned Catacomb
- 1 Exotic Orchard
- 1 Forbidden Orchard
- 1 Gemstone Caverns
- 1 Gloomlake Verge
- 1 Mana Confluence
- 1 Morphic Pool
- 1 Opal Palace
- 1 Plaza of Heroes
- 1 Reflecting Pool
- 1 Secluded Courtyard
- 1 Shipwreck Marsh
- 1 Spire of Industry
- 1 Starting Town
- 1 Sunken Ruins
- 1 Tainted Isle
- 1 Undercity Sewers
- 1 Underground River
- 1 Underground Sea
- 1 Unclaimed Territory
- 1 Watery Grave
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
- 1 Ad Nauseam
- 1 Brainstorm
- 1 Gitaxian Probe
- 1 Mystic Remora
- 1 Rhystic Study
- 1 Ponder
- 1 Preordain
- 1 Sensei's Divining Top
- 1 Skullclamp
- 1 Windfall
- 1 Moon-Circuit Hacker
- 1 Ingenious Infiltrator
- 1 Ninja of the Deep Hours
- 1 An Offer You Can't Refuse
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
- 1 The One Ring
- 1 Vexing Bauble
- 1 Chain of Vapor
- 1 Deadly Rollick
- 1 Dismember
- 1 Dokuchi Silencer
- 1 Go for the Throat
- 1 Into the Flood Maw
- 1 Orcish Bowmasters
- 1 Pongify
- 1 Snuff Out
- 1 Biting-Palm Ninja
- 1 Mist-Syndicate Naga
- 1 Dokuchi Shadow-Walker
- 1 Fallen Shinobi
- 1 Foot Mystic
- 1 Futurist Operative
- 1 Higure, the Still Wind
- 1 Ink-Eyes, Servant of Oni
- 1 Kami of Restless Shadows
- 1 Kotose, the Silent Spider
- 1 Krang & Shredder
- 1 Moonblade Shinobi
- 1 Mukotai Ambusher
- 1 Okiba-Gang Shinobi
- 1 Cyclonic Rift
- 1 Demonic Tutor
- 1 Imperial Seal
- 1 Mystical Tutor
- 1 Vampiric Tutor
- 1 Tainted Pact
- 1 Thassa's Oracle

### 15. Bracket 5, Najeela, the Blade-Blossom, warriors combat

Pool 338. 99 cards, 0 block findings, repaired false.

Judge: bracket 4, disagrees. The engine is unambiguously high-power \u2014 a dozen-plus Game Changers including Thassa's Oracle + Tainted Pact, Underworld Breach + LED + Brain Freeze, the full tutor suite, and a stack of fast mana that can deploy Najeela's infinite-combat combo in the first few turns. What holds it back from cEDH is the build itself: the mana base is mostly basics with only a handful of duals for a five-color deck, and a large slice of slots are middling warrior-tribal payoffs (Azra Oddsmaker, Doomskar Warrior, Emeria Captain) rather than the best available cards or free interaction. That's an optimized deck expressing a theme, so it lands at Bracket 4 rather than 5.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 33 | 27 to 33 |  |  |
| `tapped_land` | 0 | 0 to 2 |  |  |
| `colorless_land` | 0 | 0 to 6 |  |  |
| `color_sources` | 0.66 | 0.9 or more | yes | sources of requirement: W 12.5 of 19, U 12.5 of 19, B 12.5 of 19, R 12.5 of 19, G 17.75 of 26 |
| `avg_mana_value` | 2.38 | 1.2 to 2.4 |  | over 66 nonland cards |
| `ramp` | 16 | 12 to 22 |  |  |
| `draw` | 12 | 8 to 18 |  |  |
| `removal` | 6 | 4 to 14 |  |  |
| `wipe` | 1 | 0 to 3 |  |  |
| `interaction` | 11 | 8 to 22 |  |  |
| `tutor` | 5 | 4 or more |  | Demonic Tutor, Enlightened Tutor, Imperial Seal, Kassandra, Eagle Bearer, Vampiric Tutor |
| `fast_mana` | 12 | 6 or more |  | An Offer You Can't Refuse, Chrome Mox, Dark Ritual, Lion's Eye Diamond, Lotus Petal, Mana Vault, Mox Amber, Mox Diamond, Mox Opal, Rite of Flame, Sol Ring, Springleaf Drum |
| `game_changer` | 14 | 8 or more |  | Chrome Mox, Demonic Tutor, Enlightened Tutor, Fierce Guardianship, Imperial Seal, Jeska's Will, Lion's Eye Diamond, Mana Vault, Mox Diamond, Orcish Bowmasters, Rhystic Study, Thassa's Oracle, Underworld Breach, Vampiric Tutor |
| `mana_turn_four` | 5.11 | 4.6 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.71 | 0.65 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.67 | -9 to 0.3 |  | the commander costs 3 and comes down on turn 2.33 on average |

Goldfish over 10000 hands: the commander on turn 2.33, 5.11 mana on turn four, and 0.71 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Chrome Mox, Demonic Tutor, Enlightened Tutor, Fierce Guardianship, Imperial Seal, Jeska's Will, Lion's Eye Diamond, Mana Vault, Mox Diamond, Orcish Bowmasters, Rhystic Study, Thassa's Oracle, Underworld Breach, Vampiric Tutor. Mass land denial none. Extra turns none. Combos 4.
- Najeela, the Blade-Blossom + Professional Face-Breaker (speed 5, two cards)
- Underworld Breach + Lotus Petal + Brain Freeze (speed 5)
- Tainted Pact + Thassa's Oracle (speed 4, two cards)
- Underworld Breach + Lion's Eye Diamond + Brain Freeze (speed 5)

Findings:

- INFO `curve_summary`: average mana value 2.38 over 66 nonland cards
- WARN `profile_off_band`: the worst color's share of the sources it needs is 0.66, and bracket 5 wants 0.9 or more (sources of requirement: W 12.5 of 19, U 12.5 of 19, B 12.5 of 19, R 12.5 of 19, G 17.75 of 26)
- INFO `mana_pass`: the builder moved 8 cards of the mana base to bring the deck inside its power level

Cards:

- 1 Blood Crypt
- 1 Breeding Pool
- 1 Godless Shrine
- 1 Hallowed Fountain
- 1 Steam Vents
- 1 Command Tower
- 1 Exotic Orchard
- 11 Forest
- 3 Island
- 4 Mountain
- 4 Plains
- 4 Swamp
- 1 Chrome Mox
- 1 Lion's Eye Diamond
- 1 Lotus Petal
- 1 Mana Vault
- 1 Mox Amber
- 1 Mox Diamond
- 1 Mox Opal
- 1 Springleaf Drum
- 1 Sol Ring
- 1 Dark Ritual
- 1 Rite of Flame
- 1 Jeska's Will
- 1 Ragavan, Nimble Pilferer
- 1 Professional Face-Breaker
- 1 Neheb, Dreadhorde Champion
- 1 Rhystic Study
- 1 Mindblade Render
- 1 Azra Oddsmaker
- 1 Akki Ronin
- 1 Raiders' Spoils
- 1 Samut, Vizier of Naktamun
- 1 Kutzil, Malamet Exemplar
- 1 Oakhame Adversary
- 1 Nimble Trapfinder
- 1 The Destined Thief
- 1 Kassandra, Eagle Bearer
- 1 Eshki Dragonclaw
- 1 An Offer You Can't Refuse
- 1 Fierce Guardianship
- 1 Concerted Defense
- 1 Eladamri, Lord of Leaves
- 1 Hakoda, Selfless Commander
- 1 Kashi-Tribe Elite
- 1 Linvala, Shield of Sea Gate
- 1 Rhys, the Evermore
- 1 Seasoned Dungeoneer
- 1 Selfless Samurai
- 1 Winota, Joiner of Forces
- 1 Orcish Bowmasters
- 1 Accursed Marauder
- 1 Cacophony Scamp
- 1 Dreadhorde Butcher
- 1 Fleshbag Marauder
- 1 Goblin Cratermaker
- 1 Combat Celebrant
- 1 Arashin Foremost
- 1 Aspiring Champion
- 1 Bull-Rush Bruiser
- 1 Doomskar Warrior
- 1 Emeria Captain
- 1 Kimahri, Valiant Guardian
- 1 Proud Wildbonder
- 1 Raiyuu, Storm's Edge
- 1 Surrak, the Hunt Caller
- 1 Vengeful Firebrand
- 1 The Destined Warrior
- 1 Goblin Chainwhirler
- 1 Demonic Tutor
- 1 Enlightened Tutor
- 1 Imperial Seal
- 1 Vampiric Tutor
- 1 Underworld Breach
- 1 Brain Freeze
- 1 Tainted Pact
- 1 Thassa's Oracle
- 1 Shang-Chi, Master of Kung Fu

