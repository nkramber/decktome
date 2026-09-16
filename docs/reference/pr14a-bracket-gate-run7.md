# PR-14A bracket gate

Run date: 2026-09-16. Card snapshot: 2026-09-04.

Verdict: FAIL. 6 of 6 decks passed every block check, 5 sat in every band, 6 held no content violation, and the judge agreed with the bracket on 5 of 6 (83 percent, the bar is 80).

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
| Judge agreed with the bracket | 5 |
| Judge errors | 0 |
| Errors | 0 |
| Prompt version | 13 |
| Calls | 12 |
| Cost | $0.5235 |
| Time | 406 seconds |

## Run

- Suite `bracket`, run `pr14a-bracket-gate-run7`, on 2026-09-16, commit `76e53c7`.
- Partial run over `10,11,12,13,14,15`. It reads the bars of its own items, and it never stands for the suite.
- Roles: generate on `gpt-5.6-terra` (openai, effort medium), judge on `claude-opus-5` (anthropic, effort medium), repair on `gpt-5.6-terra` (openai, effort medium).
- Versions: card snapshot 2026-09-04, generate prompt version 13.
- Calls: 12. Cost: $0.5235. Time: 406 seconds.

## Off-band features by key

Each row counts the decks whose feature sat outside its band after the repair turns.

| Feature | Decks off band |
|---|---|
| `color_sources` | 1 |

## Decks

### 10. Bracket 4, Urza, Lord High Artificer, artifacts

Pool 317. 99 cards, 0 block findings, repaired false.

Judge: bracket 4, agrees. Eight marked Game Changers plus a fast-mana suite (Workshop, LED, Mana Vault, the Moxen, Lotus Petal, Sol Ring) blow past Bracket 3's three-Game-Changer cap and set up turn-two or -three Urza with huge construct mana. Artifact tutors (Whir of Invention, Transmute Artifact, Reshape, Arcum Dagsson, Mystic Forge) assemble KCI/Aetherflux-style engines behind Force of Will and Fierce Guardianship. Still, the clunky 29-Island mana base and filler bodies like Chrome Steed, Jhoira's Familiar and Panther Robot keep it short of a tuned cEDH list, so it's Optimized rather than 5.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 33 | 31 to 36 |  |  |
| `tapped_land` | 0 | 0 to 5 |  |  |
| `colorless_land` | 4 | 0 to 5 |  | Academy Ruins, Inventors' Fair, Mishra's Workshop, Urza's Saga |
| `color_sources` | 1.14 | 0.9 or more |  | sources of requirement: U 32 of 28 |
| `avg_mana_value` | 2.98 | 1.6 to 3 |  | over 66 nonland cards |
| `ramp` | 13 | 10 to 16 |  |  |
| `draw` | 12 | 8 to 16 |  |  |
| `removal` | 10 | 6 to 14 |  |  |
| `wipe` | 2 | 0 to 5 |  |  |
| `interaction` | 11 | 6 to 16 |  |  |
| `tutor` | 4 | 2 or more |  | Arcum Dagsson, Reshape, Transmute Artifact, Whir of Invention |
| `fast_mana` | 10 | 3 or more |  | An Offer You Can't Refuse, Chrome Mox, Lion's Eye Diamond, Lotus Petal, Mana Vault, Moonsnare Prototype, Mox Diamond, Mox Opal, Sol Ring, Springleaf Drum |
| `game_changer` | 8 | 4 or more |  | Chrome Mox, Fierce Guardianship, Force of Will, Lion's Eye Diamond, Mana Vault, Mishra's Workshop, Mox Diamond, Rhystic Study |
| `mana_turn_four` | 5.18 | 4.4 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.71 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.73 | -9 to 0.5 |  | the commander costs 4 and comes down on turn 3.27 on average |

Goldfish over 10000 hands: the commander on turn 3.27, 5.18 mana on turn four, and 0.71 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Chrome Mox, Fierce Guardianship, Force of Will, Lion's Eye Diamond, Mana Vault, Mishra's Workshop, Mox Diamond, Rhystic Study. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 2.98 over 66 nonland cards

Cards:

- 29 Island
- 1 Academy Ruins
- 1 Inventors' Fair
- 1 Mishra's Workshop
- 1 Urza's Saga
- 1 Chrome Mox
- 1 Lion's Eye Diamond
- 1 Lotus Petal
- 1 Mana Vault
- 1 Mox Diamond
- 1 Mox Opal
- 1 Moonsnare Prototype
- 1 Sol Ring
- 1 Springleaf Drum
- 1 Metalworker
- 1 Krark-Clan Ironworks
- 1 Vedalken Engineer
- 1 Grand Architect
- 1 Rhystic Study
- 1 Thoughtcast
- 1 Thirst for Knowledge
- 1 Sai, Master Thopterist
- 1 Riddlesmith
- 1 Forensic Gadgeteer
- 1 Era of Innovation
- 1 Esoteric Duplicator
- 1 Katara, Bending Prodigy
- 1 Nexus of Becoming
- 1 Uthros Research Craft
- 1 Reverse Engineer
- 1 Force of Will
- 1 Fierce Guardianship
- 1 Metallic Rebuke
- 1 Stoic Rebuttal
- 1 Disruption Protocol
- 1 An Offer You Can't Refuse
- 1 Ice Out
- 1 Override
- 1 Padeem, Consul of Innovation
- 1 Welding Jar
- 1 Shimmer Dragon
- 1 Aether Spellbomb
- 1 Aetherflux Reservoir
- 1 Arcum Dagsson
- 1 Cyber Conversion
- 1 Resculpt
- 1 Ravenform
- 1 Caltrops
- 1 Contagion Clasp
- 1 Into Thin Air
- 1 Transmogrifying Wand
- 1 Reshape
- 1 Transmute Artifact
- 1 Whir of Invention
- 1 Mystic Forge
- 1 Power Conduit
- 1 Voltaic Key
- 1 Frogmite
- 1 Arcbound Crusher
- 1 Chrome Steed
- 1 Filigree Attendant
- 1 Foundry Assembler
- 1 Ironheart, Clever Champion
- 1 Jhoira's Familiar
- 1 Lodestone Golem
- 1 Panther Robot
- 1 Phyrexian Metamorph
- 1 Traxos, Scourge of Kroog
- 1 Valkyrie Aerial Unit
- 1 Engineered Explosives
- 1 Nevinyrral's Disk

### 11. Bracket 4, Korvold, Fae-Cursed King, treasure sacrifice

Pool 326. 99 cards, 0 block findings, repaired false.

Judge: bracket 4, agrees. Eight marked Game Changers plus Sol Ring, Lotus Petal, Mox Opal and original dual lands put this well past Bracket 3's three-Game-Changer ceiling. Heavy tutor density (Razaketh, Sidisi, Demonic/Vampiric Tutor, Eldritch Evolution) and easy two-card loops (Prossh with Goblin Bombardment/Blasting Station, Yawgmoth lines) mean it can close games very early. It's still a Korvold sacrifice/aristocrats theme with some cute filler rather than a lean cEDH list, so Bracket 4 rather than 5.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 33 | 31 to 36 |  |  |
| `tapped_land` | 0 | 0 to 5 |  |  |
| `colorless_land` | 0 | 0 to 5 |  |  |
| `color_sources` | 1.07 | 0.9 or more |  | sources of requirement: B 25.25 of 23, R 19.25 of 18, G 20.25 of 19 |
| `avg_mana_value` | 2.95 | 1.6 to 3 |  | over 66 nonland cards |
| `ramp` | 13 | 10 to 16 |  |  |
| `draw` | 11 | 8 to 16 |  |  |
| `removal` | 10 | 6 to 14 |  |  |
| `wipe` | 1 | 0 to 5 |  |  |
| `interaction` | 10 | 6 to 16 |  |  |
| `tutor` | 6 | 2 or more |  | Demonic Tutor, Diabolic Intent, Eldritch Evolution, Razaketh, the Foulblooded, Sidisi, Undead Vizier, Vampiric Tutor |
| `fast_mana` | 9 | 3 or more |  | Chrome Mox, Dark Ritual, Lion's Eye Diamond, Lotus Petal, Mana Vault, Mox Diamond, Mox Opal, Rite of Flame, Sol Ring |
| `game_changer` | 8 | 4 or more |  | Chrome Mox, Demonic Tutor, Jeska's Will, Lion's Eye Diamond, Mana Vault, Mox Diamond, Orcish Bowmasters, Vampiric Tutor |
| `mana_turn_four` | 4.83 | 4.4 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.71 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.49 | -9 to 0.5 |  | the commander costs 5 and comes down on turn 4.51 on average |

Goldfish over 10000 hands: the commander on turn 4.51, 4.83 mana on turn four, and 0.71 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Chrome Mox, Demonic Tutor, Jeska's Will, Lion's Eye Diamond, Mana Vault, Mox Diamond, Orcish Bowmasters, Vampiric Tutor. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 2.95 over 66 nonland cards
- INFO `mana_pass`: the builder moved 3 cards of the mana base to bring the deck inside its power level

Cards:

- 1 Badlands
- 1 Bayou
- 1 Blood Crypt
- 1 Blightstep Pathway // Searstep Pathway
- 1 City of Brass
- 1 Command Tower
- 1 Cragcrown Pathway // Timbercrown Pathway
- 1 Darkbore Pathway // Slitherbore Pathway
- 1 Forbidden Orchard
- 1 Karplusan Forest
- 1 Llanowar Wastes
- 1 Mana Confluence
- 1 Overgrown Tomb
- 1 Phyrexian Tower
- 1 Stomping Ground
- 1 Sulfurous Springs
- 1 Taiga
- 1 Tarnished Citadel
- 3 Forest
- 7 Swamp
- 2 Mountain
- 1 Chrome Mox
- 1 Lion's Eye Diamond
- 1 Mana Vault
- 1 Mox Diamond
- 1 Mox Opal
- 1 Lotus Petal
- 1 Sol Ring
- 1 Dark Ritual
- 1 Rite of Flame
- 1 Jeska's Will
- 1 Pitiless Plunderer
- 1 Professional Face-Breaker
- 1 Grim Hireling
- 1 Anje, Maid of Dishonor
- 1 Disciple of Bolas
- 1 Smothering Abomination
- 1 Tevesh Szat, Doom of Fools
- 1 Shadowheart, Dark Justiciar
- 1 Relic Vial
- 1 Lord Skitter's Butcher
- 1 Sanguine Spy
- 1 Skyclave Shadowcat
- 1 Vampire Gourmand
- 1 Tribute to Horobi // Echo of Death's Wail
- 1 Altar of the Wretched // Wretched Bonemass
- 1 Constant Mists
- 1 Cultist of the Absolute
- 1 Dark Privilege
- 1 Eldrazi Monument
- 1 Gift of Doom
- 1 Goblin Chirurgeon
- 1 Gyome, Master Chef
- 1 Kinzu of the Bleak Coven
- 1 Nightmare Shepherd
- 1 Attrition
- 1 Bone Shards
- 1 Broadside Bombardiers
- 1 Eaten Alive
- 1 Goblin Bombardment
- 1 Orcish Bowmasters
- 1 Yawgmoth, Thran Physician
- 1 Grist, the Hunger Tide
- 1 Minsc & Boo, Timeless Heroes
- 1 Blasting Station
- 1 Academy Manufactor
- 1 Xorn
- 1 Witch's Oven
- 1 Viscera Seer
- 1 Diabolic Intent
- 1 Eldritch Evolution
- 1 Blood Host
- 1 Brawl-Bash Ogre
- 1 Flesh-Eater Imp
- 1 Ghoulcaller Gisa
- 1 Kethek, Crucible Goliath
- 1 Meren of Clan Nel Toth
- 1 Prossh, Skyraider of Kher
- 1 Razaketh, the Foulblooded
- 1 Sadistic Hypnotist
- 1 Sidisi, Undead Vizier
- 1 Undercity Scavenger
- 1 Vindictive Vampire
- 1 The Meathook Massacre
- 1 Demonic Tutor
- 1 Vampiric Tutor
- 1 Revel in Riches
- 1 Mount Doom
- 1 Nurturing Peatland
- 1 Grove of the Burnwillows

### 12. Bracket 4, Prosper, Tome-Bound, treasure

Pool 210. 99 cards, 0 block findings, repaired false.

Judge: bracket 5, disagrees. The manabase (Badlands, Blood Crypt, fetches, City of Brass/Mana Confluence, Cavern) plus a dozen pieces of fast mana 
top-tier acceleration (Sol Ring, Mana Vault, Lion's Eye Diamond, three Moxen, Dark Ritual, Lotus Petal, Jeska's Will) and the full suite of black tutors (Demonic, Vampiric, Imperial Seal, Gamble) put this far past a merely optimized build. It packs roughly a dozen marked Game Changers and a genuine cEDH engine in Lion's Eye Diamond plus Underworld Breach with Ad Nauseam and Wheel to fuel storm turns, backed by free interaction (Deadly Rollick, Snuff Out, Pyroblast, Red Elemental Blast) to protect or stop a combo. That is a metagame-facing, turn-three-or-four capable combo deck rather than a treasure theme deck, so it sits at cEDH.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 33 | 31 to 36 |  |  |
| `tapped_land` | 1 | 0 to 5 |  | Foreboding Ruins |
| `colorless_land` | 0 | 0 to 5 |  |  |
| `color_sources` | 1.68 | 0.9 or more |  | sources of requirement: B 33 of 19, R 32 of 19 |
| `avg_mana_value` | 2.21 | 1.6 to 3 |  | over 66 nonland cards |
| `ramp` | 13 | 10 to 16 |  |  |
| `draw` | 12 | 8 to 16 |  |  |
| `removal` | 10 | 6 to 14 |  |  |
| `wipe` | 2 | 0 to 5 |  |  |
| `interaction` | 11 | 6 to 16 |  |  |
| `tutor` | 4 | 2 or more |  | Demonic Tutor, Gamble, Imperial Seal, Vampiric Tutor |
| `fast_mana` | 10 | 3 or more |  | Chrome Mox, Dark Ritual, Lion's Eye Diamond, Lotus Petal, Mana Vault, Mox Diamond, Mox Opal, Rite of Flame, Sol Ring, Springleaf Drum |
| `game_changer` | 13 | 4 or more |  | Ad Nauseam, Chrome Mox, Demonic Tutor, Gamble, Imperial Seal, Jeska's Will, Lion's Eye Diamond, Mana Vault, Mox Diamond, Orcish Bowmasters, The One Ring, Underworld Breach, Vampiric Tutor |
| `mana_turn_four` | 4.97 | 4.4 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.71 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.64 | -9 to 0.5 |  | the commander costs 4 and comes down on turn 3.36 on average |

Goldfish over 10000 hands: the commander on turn 3.36, 4.97 mana on turn four, and 0.71 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Ad Nauseam, Chrome Mox, Demonic Tutor, Gamble, Imperial Seal, Jeska's Will, Lion's Eye Diamond, Mana Vault, Mox Diamond, Orcish Bowmasters, The One Ring, Underworld Breach, Vampiric Tutor. Mass land denial none. Extra turns none. Combos 3.
- Underworld Breach + Wheel of Fortune + Jeska's Will (speed 5)
- Prosper, Tome-Bound + Professional Face-Breaker + Sensei's Divining Top + Xorn (speed 5)
- Underworld Breach + Lion's Eye Diamond + Wheel of Fortune (speed 5)

Findings:

- INFO `curve_summary`: average mana value 2.21 over 66 nonland cards
- INFO `mana_pass`: the builder moved 2 cards of the mana base to bring the deck inside its power level

Cards:

- 1 Badlands
- 1 Blackcleave Cliffs
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
- 1 Graven Cairns
- 1 Haunted Ridge
- 1 Luxury Suite
- 1 Mana Confluence
- 3 Mountain
- 1 Prismatic Vista
- 1 Secluded Courtyard
- 1 Smoldering Marsh
- 1 Spire of Industry
- 1 Starting Town
- 4 Swamp
- 1 Sulfurous Springs
- 1 Tainted Peak
- 1 Tarnished Citadel
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
- 1 Ad Nauseam
- 1 Corrupted Conviction
- 1 Demand Answers
- 1 Faithless Looting
- 1 Idol of Oblivion
- 1 Mishra's Bauble
- 1 Night's Whisper
- 1 Sensei's Divining Top
- 1 Sign in Blood
- 1 Skullclamp
- 1 Village Rites
- 1 Wheel of Fortune
- 1 Kaya's Ghostform
- 1 Lightning Greaves
- 1 Mithril Coat
- 1 Not Dead After All
- 1 Pyroblast
- 1 Red Elemental Blast
- 1 Swiftfoot Boots
- 1 The One Ring
- 1 Tibalt's Trickery
- 1 Undying Malice
- 1 Vexing Bauble
- 1 Abrade
- 1 Chaos Warp
- 1 Deadly Rollick
- 1 Feed the Swarm
- 1 Go for the Throat
- 1 Infernal Grasp
- 1 Mayhem Devil
- 1 Orcish Bowmasters
- 1 Snuff Out
- 1 Terminate
- 1 Academy Manufactor
- 1 Crossover Collaboration
- 1 Evereth, Viceroy of Plunder
- 1 Gilded Assault Cart
- 1 Heroes for Hire
- 1 Xorn
- 1 Casey & Raph, Hotheads
- 1 Flamekin Gildweaver
- 1 Jaded Sell-Sword
- 1 Marut
- 1 Meticulous Artisan
- 1 Namazu Trader
- 1 Toxic Deluge
- 1 Vandalblast
- 1 Demonic Tutor
- 1 Gamble
- 1 Imperial Seal
- 1 Revel in Riches
- 1 Underworld Breach
- 1 Vampiric Tutor
- 1 Mount Doom
- 1 Foreboding Ruins

### 13. Bracket 5, Kinnan, Bonder Prodigy, combo

Pool 322. 99 cards, 0 block findings, repaired false.

Judge: bracket 5, agrees. This is a textbook cEDH Kinnan build: eight-plus Game Changers including the full fast-mana suite (Chrome Mox, Grim Monolith, LED, Mana Vault, Mox Diamond, Sol Ring, Lotus Petal), an untapper package (Freed from the Real, Dramatic Reversal, Voltaic Construct, Filigree Sages, Devoted Druid/Basalt Monolith) for turn-two or -three infinite mana, and Thassa's Oracle plus Lab Man/Jace as the deterministic wincon. Interaction is all free or cheap counterspells (Fierce Guardianship, Flusterstorm, An Offer You Can't Refuse) protecting the combo, with an optimized fetch/dual mana base. It is a metagame-focused combo deck, not a theme, and can win on nearly any turn.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 30 | 27 to 33 |  |  |
| `tapped_land` | 1 | 0 to 2 |  | Vineglimmer Snarl |
| `colorless_land` | 0 | 0 to 6 |  |  |
| `color_sources` | 1.21 | 0.9 or more |  | sources of requirement: U 31.5 of 26, G 29.75 of 19 |
| `avg_mana_value` | 2.32 | 1.2 to 2.4 |  | over 69 nonland cards |
| `ramp` | 17 | 12 to 22 |  |  |
| `draw` | 12 | 8 to 18 |  |  |
| `removal` | 9 | 4 to 14 |  |  |
| `wipe` | 1 | 0 to 3 |  |  |
| `interaction` | 14 | 8 to 22 |  |  |
| `tutor` | 4 | 4 or more |  | Alphinaud Leveilleur, Kuldotha Forgemaster, Muddle the Mixture, Tezzeret, Cruel Captain |
| `fast_mana` | 10 | 6 or more |  | An Offer You Can't Refuse, Chrome Mox, Lion's Eye Diamond, Lotus Petal, Mana Vault, Mox Amber, Mox Diamond, Mox Opal, Sol Ring, Springleaf Drum |
| `game_changer` | 9 | 8 or more |  | Chrome Mox, Cyclonic Rift, Fierce Guardianship, Grim Monolith, Lion's Eye Diamond, Mana Vault, Mox Diamond, Rhystic Study, Thassa's Oracle |
| `mana_turn_four` | 5.16 | 4.6 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.66 | 0.65 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.4 | -9 to 0.3 |  | the commander costs 2 and comes down on turn 1.61 on average |

Goldfish over 10000 hands: the commander on turn 1.61, 5.16 mana on turn four, and 0.66 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Chrome Mox, Cyclonic Rift, Fierce Guardianship, Grim Monolith, Lion's Eye Diamond, Mana Vault, Mox Diamond, Rhystic Study, Thassa's Oracle. Mass land denial none. Extra turns none. Combos 2.
- Kitsa, Otterball Elite + Dramatic Reversal (speed 5, two cards)
- Kinnan, Bonder Prodigy + Basalt Monolith (speed 5, two cards)

Findings:

- INFO `curve_summary`: average mana value 2.32 over 69 nonland cards
- INFO `mana_pass`: the builder moved 9 cards of the mana base to bring the deck inside its power level

Cards:

- 1 Tropical Island
- 1 Breeding Pool
- 1 Command Tower
- 1 City of Brass
- 1 Mana Confluence
- 1 Forbidden Orchard
- 1 Yavimaya Coast
- 1 Misty Rainforest
- 1 Flooded Strand
- 1 Polluted Delta
- 1 Scalding Tarn
- 1 Verdant Catacombs
- 1 Windswept Heath
- 1 Wooded Foothills
- 1 Prismatic Vista
- 1 Fabled Passage
- 1 Evolving Wilds
- 1 Terramorphic Expanse
- 4 Island
- 3 Forest
- 1 Chrome Mox
- 1 Grim Monolith
- 1 Lion's Eye Diamond
- 1 Mana Vault
- 1 Mox Diamond
- 1 Mox Amber
- 1 Mox Opal
- 1 Lotus Petal
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
- 1 Frantic Search
- 1 Cloud of Faeries
- 1 Cerulean Wisps
- 1 Refocus
- 1 Twitch
- 1 Benefactor's Draught
- 1 Innocuous Researcher
- 1 Kitsa, Otterball Elite
- 1 Pip-Boy 3000
- 1 Staff of Domination
- 1 Alphinaud Leveilleur
- 1 Fierce Guardianship
- 1 Flusterstorm
- 1 An Offer You Can't Refuse
- 1 Muddle the Mixture
- 1 Narset's Reversal
- 1 Shore Up
- 1 Sokka's Haiku
- 1 Unwind
- 1 Rewind
- 1 Katara's Reversal
- 1 Legolas's Quick Reflexes
- 1 Illusionist's Gambit
- 1 Octopus Form
- 1 Magic Damper
- 1 Snap
- 1 Emerald Charm
- 1 Provoke
- 1 Apathy
- 1 Amphibian Downpour
- 1 Singing Bell Strike
- 1 Ringing Strike Mastery
- 1 Aggressive Biomancy
- 1 Bushwhack
- 1 Dramatic Reversal
- 1 Freed from the Real
- 1 Filigree Sages
- 1 Voltaic Construct
- 1 Peregrine Drake
- 1 Kuldotha Forgemaster
- 1 Phyrexian Metamorph
- 1 Clever Impersonator
- 1 Sakashima's Student
- 1 Dross Scorpion
- 1 Breaching Hippocamp
- 1 Cacophodon
- 1 Cyclonic Rift
- 1 Thassa's Oracle
- 1 Laboratory Maniac
- 1 Jace, Wielder of Mysteries
- 1 Tezzeret, Cruel Captain
- 1 Hinterland Harbor
- 1 Rejuvenating Springs
- 1 Dreamroot Cascade
- 1 Vineglimmer Snarl
- 1 Starting Town

### 14. Bracket 5, Yuriko, the Tiger's Shadow, ninjas

Pool 238. 99 cards, 0 block findings, repaired false.

Judge: bracket 5, agrees. This is a tuned cEDH-caliber Yuriko list: a dozen-plus Game Changers including Lion's Eye Diamond, Mox Diamond, Grim Monolith, Chrome Mox and a full suite of one-card tutors (Imperial Seal, Vampiric/Mystical Tutor, Demonic Tutor) alongside Ad Nauseam. The interaction is free-spell competitive staples — Force of Will, Fierce Guardianship, Pact of Negation, Deadly Rollick, Mindbreak Trap, Flusterstorm — and the mana base is dual-land/fetch optimized with original duals. It plays the strongest known version of the archetype rather than a theme, and can close games in the first few turns.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 30 | 27 to 33 |  |  |
| `tapped_land` | 1 | 0 to 2 |  | Choked Estuary |
| `colorless_land` | 0 | 0 to 6 |  |  |
| `color_sources` | 1.83 | 0.9 or more |  | sources of requirement: U 34.75 of 19, B 35.75 of 19 |
| `avg_mana_value` | 2.25 | 1.2 to 2.4 |  | over 69 nonland cards |
| `ramp` | 17 | 12 to 22 |  |  |
| `draw` | 12 | 8 to 18 |  |  |
| `removal` | 9 | 4 to 14 |  |  |
| `wipe` | 1 | 0 to 3 |  |  |
| `interaction` | 14 | 8 to 22 |  |  |
| `tutor` | 5 | 4 or more |  | Demonic Tutor, Higure, the Still Wind, Imperial Seal, Mystical Tutor, Vampiric Tutor |
| `fast_mana` | 13 | 6 or more |  | An Offer You Can't Refuse, Chrome Mox, Culling the Weak, Dark Ritual, Everflowing Chalice, Lion's Eye Diamond, Lotus Petal, Mana Vault, Mox Amber, Mox Diamond, Mox Opal, Sol Ring, Springleaf Drum |
| `game_changer` | 16 | 8 or more |  | Ad Nauseam, Chrome Mox, Consecrated Sphinx, Cyclonic Rift, Demonic Tutor, Fierce Guardianship, Force of Will, Grim Monolith, Imperial Seal, Lion's Eye Diamond, Mana Vault, Mox Diamond, Mystical Tutor, Orcish Bowmasters, Rhystic Study, Vampiric Tutor |
| `mana_turn_four` | 5.36 | 4.6 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.66 | 0.65 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.7 | -9 to 0.3 |  | the commander costs 3 and comes down on turn 2.3 on average |

Goldfish over 10000 hands: the commander on turn 2.3, 5.36 mana on turn four, and 0.66 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Ad Nauseam, Chrome Mox, Consecrated Sphinx, Cyclonic Rift, Demonic Tutor, Fierce Guardianship, Force of Will, Grim Monolith, Imperial Seal, Lion's Eye Diamond, Mana Vault, Mox Diamond, Mystical Tutor, Orcish Bowmasters, Rhystic Study, Vampiric Tutor. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 2.25 over 69 nonland cards
- INFO `mana_pass`: the builder moved 7 cards of the mana base to bring the deck inside its power level

Cards:

- 1 Cavern of Souls
- 1 Clearwater Pathway // Murkwater Pathway
- 1 Command Tower
- 1 City of Brass
- 1 Darkslick Shores
- 1 Darkwater Catacombs
- 1 Exotic Orchard
- 1 Fabled Passage
- 1 Gloomlake Verge
- 1 Morphic Pool
- 1 Polluted Delta
- 1 Prismatic Vista
- 1 Secluded Courtyard
- 1 Spire of Industry
- 1 Starting Town
- 1 Sunken Ruins
- 1 Tainted Isle
- 1 Underground River
- 1 Underground Sea
- 1 Watery Grave
- 1 Island
- 2 Swamp
- 1 Arcane Signet
- 1 Chrome Mox
- 1 Culling the Weak
- 1 Dark Ritual
- 1 Dimir Signet
- 1 Everflowing Chalice
- 1 Fellwar Stone
- 1 Grim Monolith
- 1 Lion's Eye Diamond
- 1 Lotus Petal
- 1 Mana Vault
- 1 Mox Amber
- 1 Mox Diamond
- 1 Mox Opal
- 1 Sol Ring
- 1 Springleaf Drum
- 1 Talisman of Dominance
- 1 Ad Nauseam
- 1 Brainstorm
- 1 Consecrated Sphinx
- 1 Gitaxian Probe
- 1 Mystic Remora
- 1 Rhystic Study
- 1 Sensei's Divining Top
- 1 Ponder
- 1 Preordain
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
- 1 Chain of Vapor
- 1 Deadly Rollick
- 1 Dismember
- 1 Dokuchi Silencer
- 1 Feed the Swarm
- 1 Into the Flood Maw
- 1 Orcish Bowmasters
- 1 Pongify
- 1 Snuff Out
- 1 Cyclonic Rift
- 1 Dokuchi Shadow-Walker
- 1 Fallen Shinobi
- 1 Foot Mystic
- 1 Futurist Operative
- 1 Higure, the Still Wind
- 1 Ink-Eyes, Servant of Oni
- 1 Kami of Restless Shadows
- 1 Kotose, the Silent Spider
- 1 Moonblade Shinobi
- 1 Silent-Blade Oni
- 1 Satoru Umezawa
- 1 Silver-Fur Master
- 1 Demonic Tutor
- 1 Imperial Seal
- 1 Mystical Tutor
- 1 Vampiric Tutor
- 1 Mana Confluence
- 1 Forbidden Orchard
- 1 Tarnished Citadel
- 1 Sunken Hollow
- 1 Drowned Catacomb
- 1 Shipwreck Marsh
- 1 Choked Estuary

### 15. Bracket 5, Najeela, the Blade-Blossom, warriors combat

Pool 338. 99 cards, 0 block findings, repaired false.

Judge: bracket 5, agrees. Najeela with Tainted Pact + Thassa's Oracle, Underworld Breach + Lion's Eye Diamond, and the Najeela infinite-combat line is a cEDH shell, not a warrior theme deck. It's packed with ten-plus Game Changers, a full fast-mana suite (Chrome Mox, Mox Diamond/Opal/Amber, Lotus Petal, Mana Vault, Dark Ritual, Jeska's Will), free interaction like Fierce Guardianship, and black tutors. Only the mediocre basic-heavy mana base holds it back slightly, but the strategy and win conditions are firmly competitive.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 33 | 27 to 33 |  |  |
| `tapped_land` | 1 | 0 to 2 |  | Base Camp |
| `colorless_land` | 0 | 0 to 6 |  |  |
| `color_sources` | 0.7 | 0.9 or more | yes | sources of requirement: W 13.25 of 19, U 14.25 of 19, B 14.25 of 19, R 13.25 of 19, G 18.75 of 26 |
| `avg_mana_value` | 2.39 | 1.2 to 2.4 |  | over 66 nonland cards |
| `ramp` | 20 | 12 to 22 |  |  |
| `draw` | 12 | 8 to 18 |  |  |
| `removal` | 8 | 4 to 14 |  |  |
| `wipe` | 1 | 0 to 3 |  |  |
| `interaction` | 12 | 8 to 22 |  |  |
| `tutor` | 4 | 4 or more |  | Blaring Captain, Cazur, Ruthless Stalker, Demonic Tutor, Kassandra, Eagle Bearer |
| `fast_mana` | 12 | 6 or more |  | An Offer You Can't Refuse, Chrome Mox, Dark Ritual, Lion's Eye Diamond, Lotus Petal, Mana Vault, Mox Amber, Mox Diamond, Mox Opal, Rite of Flame, Sol Ring, Springleaf Drum |
| `game_changer` | 12 | 8 or more |  | Chrome Mox, Demonic Tutor, Fierce Guardianship, Jeska's Will, Lion's Eye Diamond, Mana Vault, Mox Diamond, Orcish Bowmasters, Rhystic Study, Smothering Tithe, Thassa's Oracle, Underworld Breach |
| `mana_turn_four` | 5.24 | 4.6 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.71 | 0.65 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.67 | -9 to 0.3 |  | the commander costs 3 and comes down on turn 2.33 on average |

Goldfish over 10000 hands: the commander on turn 2.33, 5.24 mana on turn four, and 0.71 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Chrome Mox, Demonic Tutor, Fierce Guardianship, Jeska's Will, Lion's Eye Diamond, Mana Vault, Mox Diamond, Orcish Bowmasters, Rhystic Study, Smothering Tithe, Thassa's Oracle, Underworld Breach. Mass land denial none. Extra turns none. Combos 2.
- Najeela, the Blade-Blossom + Professional Face-Breaker (speed 5, two cards)
- Tainted Pact + Thassa's Oracle (speed 4, two cards)

Findings:

- INFO `curve_summary`: average mana value 2.39 over 66 nonland cards
- WARN `profile_off_band`: the worst color's share of the sources it needs is 0.7, and bracket 5 wants 0.9 or more (sources of requirement: W 13.25 of 19, U 14.25 of 19, B 14.25 of 19, R 13.25 of 19, G 18.75 of 26)
- INFO `mana_pass`: the builder moved 12 cards of the mana base to bring the deck inside its power level
- INFO `basics_added`: the list was 1 card short, so the builder added 1 basic land

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
- 10 Forest
- 3 Island
- 3 Mountain
- 4 Plains
- 3 Swamp
- 1 Chrome Mox
- 1 Lion's Eye Diamond
- 1 Lotus Petal
- 1 Mana Vault
- 1 Mox Amber
- 1 Mox Diamond
- 1 Mox Opal
- 1 Sol Ring
- 1 Springleaf Drum
- 1 Dark Ritual
- 1 Rite of Flame
- 1 Jeska's Will
- 1 Smothering Tithe
- 1 Ragavan, Nimble Pilferer
- 1 Professional Face-Breaker
- 1 Grand Warlord Radha
- 1 Rhystic Study
- 1 Kassandra, Eagle Bearer
- 1 Mindblade Render
- 1 Azra Oddsmaker
- 1 Akki Ronin
- 1 Kutzil, Malamet Exemplar
- 1 Samut, Vizier of Naktamun
- 1 Nimble Trapfinder
- 1 Oakhame Adversary
- 1 Neyith of the Dire Hunt
- 1 Raiders' Spoils
- 1 Jaxis, the Troublemaker
- 1 Fierce Guardianship
- 1 Concerted Defense
- 1 Ainok Strike Leader
- 1 Akiri, Fearless Voyager
- 1 Eladamri, Lord of Leaves
- 1 Ezuri, Renegade Leader
- 1 Hakoda, Selfless Commander
- 1 Kashi-Tribe Elite
- 1 Linvala, Shield of Sea Gate
- 1 Rhys, the Evermore
- 1 Seasoned Dungeoneer
- 1 Selfless Samurai
- 1 Orcish Bowmasters
- 1 Accursed Marauder
- 1 Dreadhorde Butcher
- 1 Goblin Cratermaker
- 1 Chatterfang, Squirrel General
- 1 Fleshbag Marauder
- 1 Practiced Tactics
- 1 Aspiring Champion
- 1 Doomskar Warrior
- 1 Imaryll, Elfhame Elite
- 1 Jazal Goldmane
- 1 Raiyuu, Storm's Edge
- 1 Blaring Captain
- 1 Cazur, Ruthless Stalker
- 1 Combat Celebrant
- 1 Arashin Foremost
- 1 Goblin Chainwhirler
- 1 Demonic Tutor
- 1 Underworld Breach
- 1 Tainted Pact
- 1 Thassa's Oracle
- 1 Shang-Chi, Master of Kung Fu
- 1 Cacophony Scamp
- 1 An Offer You Can't Refuse
- 1 Radha, Heir to Keld
- 1 Heronblade Elite

