# PR-14A bracket gate

Run date: 2026-09-14. Card snapshot: 2026-09-04.

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
| Decks that needed a repair turn | 1 |
| Decks judged | 6 |
| Judge agreed with the bracket | 4 |
| Judge errors | 0 |
| Errors | 0 |
| Prompt version | 13 |
| Calls | 13 |
| Cost | $0.5619 |
| Time | 564 seconds |

## Run

- Suite `bracket`, run `pr14a-bracket-gate-run6`, on 2026-09-14, commit `8ca9e58`.
- Partial run over `10,11,12,13,14,15`. It reads the bars of its own items, and it never stands for the suite.
- Roles: generate on `gpt-5.6-terra` (openai, effort medium), judge on `claude-opus-5` (anthropic, effort medium), repair on `gpt-5.6-terra` (openai, effort medium).
- Versions: card snapshot 2026-09-04, generate prompt version 13.
- Calls: 13. Cost: $0.5619. Time: 564 seconds.

## Off-band features by key

Each row counts the decks whose feature sat outside its band after the repair turns.

| Feature | Decks off band |
|---|---|
| `color_sources` | 1 |

## Decks

### 10. Bracket 4, Urza, Lord High Artificer, artifacts

Pool 317. 99 cards, 0 block findings, repaired false.

Judge: bracket 4, agrees. Eight marked Game Changers plus a pile of unmarked fast mana (Sol Ring, Lotus Petal, Mox Opal, Metalworker, KCI) and artifact tutors (Whir of Invention, Transmute Artifact, Reshape) put it far past Bracket 3's three-Game-Changer cap, and Urza with LED/Ironworks/Metalworker can generate explosive turn-three or four mana with free counterspell protection. It isn't cEDH-tight, though: the threat suite is full of filler beaters and weak removal like Water Whip, Zuko's Exile, and Unable to Scream rather than a streamlined win line. That mix of optimized power with casual-grade card choices lands it squarely in Bracket 4.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 34 | 31 to 36 |  |  |
| `tapped_land` | 0 | 0 to 5 |  |  |
| `colorless_land` | 2 | 0 to 5 |  | Academy Ruins, Mishra's Workshop |
| `color_sources` | 1.25 | 0.9 or more |  | sources of requirement: U 35 of 28 |
| `avg_mana_value` | 2.82 | 1.6 to 3 |  | over 65 nonland cards |
| `ramp` | 12 | 10 to 16 |  |  |
| `draw` | 12 | 8 to 16 |  |  |
| `removal` | 10 | 6 to 14 |  |  |
| `wipe` | 2 | 0 to 5 |  |  |
| `interaction` | 11 | 6 to 16 |  |  |
| `tutor` | 4 | 2 or more |  | Repurposing Bay, Reshape, Transmute Artifact, Whir of Invention |
| `fast_mana` | 10 | 3 or more |  | An Offer You Can't Refuse, Chrome Mox, Lion's Eye Diamond, Lotus Petal, Mana Vault, Moonsnare Prototype, Mox Diamond, Mox Opal, Sol Ring, Springleaf Drum |
| `game_changer` | 8 | 4 or more |  | Chrome Mox, Fierce Guardianship, Force of Will, Lion's Eye Diamond, Mana Vault, Mishra's Workshop, Mox Diamond, Rhystic Study |
| `mana_turn_four` | 5.13 | 4.4 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.72 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.69 | -9 to 0.5 |  | the commander costs 4 and comes down on turn 3.31 on average |

Goldfish over 10000 hands: the commander on turn 3.31, 5.13 mana on turn four, and 0.72 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Chrome Mox, Fierce Guardianship, Force of Will, Lion's Eye Diamond, Mana Vault, Mishra's Workshop, Mox Diamond, Rhystic Study. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 2.82 over 65 nonland cards
- INFO `basics_added`: the list was 1 card short, so the builder added 1 basic land

Cards:

- 29 Island
- 1 Seat of the Synod
- 1 Mystic Sanctuary
- 1 Otawara, Soaring City
- 1 Academy Ruins
- 1 Mishra's Workshop
- 1 Chrome Mox
- 1 Lion's Eye Diamond
- 1 Mana Vault
- 1 Mox Diamond
- 1 Lotus Petal
- 1 Mox Opal
- 1 Moonsnare Prototype
- 1 Sol Ring
- 1 Springleaf Drum
- 1 Metalworker
- 1 Krark-Clan Ironworks
- 1 Grand Architect
- 1 Rhystic Study
- 1 Thoughtcast
- 1 Sai, Master Thopterist
- 1 Vedalken Archmage
- 1 Thirst for Knowledge
- 1 Riddlesmith
- 1 Era of Innovation
- 1 Esoteric Duplicator
- 1 Cerebral Download
- 1 Forensic Gadgeteer
- 1 Katara, Bending Prodigy
- 1 Edgar, King of Figaro
- 1 Force of Will
- 1 Fierce Guardianship
- 1 An Offer You Can't Refuse
- 1 Metallic Rebuke
- 1 Disruption Protocol
- 1 Ice Out
- 1 Stoic Rebuttal
- 1 Welding Jar
- 1 Ghostly Flicker
- 1 Reality Ripple
- 1 Override
- 1 Aether Spellbomb
- 1 Contagion Clasp
- 1 Cyber Conversion
- 1 Resculpt
- 1 Ravenform
- 1 Unable to Scream
- 1 Water Whip
- 1 Watery Grasp
- 1 Zuko's Exile
- 1 Shape Anew
- 1 Transmute Artifact
- 1 Whir of Invention
- 1 Reshape
- 1 Repurposing Bay
- 1 Mystic Forge
- 1 Unwinding Clock
- 1 Frogmite
- 1 Arcbound Crusher
- 1 Arcbound Reclaimer
- 1 Chrome Steed
- 1 Filigree Attendant
- 1 Jhoira's Familiar
- 1 Lodestone Golem
- 1 Master Transmuter
- 1 Phyrexian Metamorph
- 1 Panther Robot
- 1 Traxos, Scourge of Kroog
- 1 Kappa Cannoneer
- 1 Engineered Explosives
- 1 Hurkyl's Recall

### 11. Bracket 4, Korvold, Fae-Cursed King, treasure sacrifice

Pool 326. 99 cards, 0 block findings, repaired false.

Judge: bracket 4, agrees. Five marked Game Changers (Chrome Mox, Lion's Eye Diamond, Mana Vault, Mox Diamond, Orcish Bowmasters) already push it past Bracket 3, and the supporting suite of Sol Ring, Lotus Petal, Mox Opal, Dark Ritual and Rite of Flame makes for genuinely explosive turn-two/three Korvold starts. Tutors like Razaketh, Birthing Pod and Diabolic Intent assemble sacrifice engines with Pitiless Plunderer, Blasting Station and Goblin Bombardment for fast, repeatable drain kills. It is still a themed aristocrats/Korvold build rather than a metagame-tuned cEDH list, so it sits at Optimized rather than Bracket 5.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 33 | 31 to 36 |  |  |
| `tapped_land` | 0 | 0 to 5 |  |  |
| `colorless_land` | 0 | 0 to 5 |  |  |
| `color_sources` | 1.01 | 0.9 or more |  | sources of requirement: B 23.25 of 22, R 19.25 of 19, G 19.25 of 19 |
| `avg_mana_value` | 2.98 | 1.6 to 3 |  | over 66 nonland cards |
| `ramp` | 13 | 10 to 16 |  |  |
| `draw` | 11 | 8 to 16 |  |  |
| `removal` | 10 | 6 to 14 |  |  |
| `wipe` | 2 | 0 to 5 |  |  |
| `interaction` | 11 | 6 to 16 |  |  |
| `tutor` | 4 | 2 or more |  | Birthing Pod, Diabolic Intent, Magda, Brazen Outlaw, Razaketh, the Foulblooded |
| `fast_mana` | 9 | 3 or more |  | Chrome Mox, Dark Ritual, Lion's Eye Diamond, Lotus Petal, Mana Vault, Mox Diamond, Mox Opal, Rite of Flame, Sol Ring |
| `game_changer` | 5 | 4 or more |  | Chrome Mox, Lion's Eye Diamond, Mana Vault, Mox Diamond, Orcish Bowmasters |
| `mana_turn_four` | 4.73 | 4.4 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.71 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.35 | -9 to 0.5 |  | the commander costs 5 and comes down on turn 4.65 on average |

Goldfish over 10000 hands: the commander on turn 4.65, 4.73 mana on turn four, and 0.71 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Chrome Mox, Lion's Eye Diamond, Mana Vault, Mox Diamond, Orcish Bowmasters. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 2.98 over 66 nonland cards
- INFO `mana_pass`: the builder moved 1 card of the mana base to bring the deck inside its power level

Cards:

- 1 Blood Crypt
- 1 Cavern of Souls
- 1 City of Brass
- 1 Command Tower
- 1 Dragonskull Summit
- 1 Exotic Orchard
- 1 Llanowar Wastes
- 1 Luxury Suite
- 1 Mana Confluence
- 1 Overgrown Tomb
- 1 Reflecting Pool
- 1 Rootbound Crag
- 1 Stomping Ground
- 1 Sulfurous Springs
- 1 Undergrowth Stadium
- 1 Woodland Cemetery
- 7 Swamp
- 5 Forest
- 5 Mountain
- 1 Chrome Mox
- 1 Lion's Eye Diamond
- 1 Lotus Petal
- 1 Mana Vault
- 1 Mox Diamond
- 1 Mox Opal
- 1 Sol Ring
- 1 Dark Ritual
- 1 Rite of Flame
- 1 Pitiless Plunderer
- 1 Professional Face-Breaker
- 1 Goldspan Dragon
- 1 Magda, Brazen Outlaw
- 1 Ahriman
- 1 Bushmeat Poacher
- 1 Dina, Essence Brewer
- 1 Gas Guzzler
- 1 Ice Cream Kitty
- 1 Relic Vial
- 1 Shadowheart, Dark Justiciar
- 1 Smothering Abomination
- 1 Stormclaw Rager
- 1 Tevesh Szat, Doom of Fools
- 1 Vampire Gourmand
- 1 Attrition
- 1 Blasting Station
- 1 Bone Shards
- 1 Broadside Bombardiers
- 1 Eaten Alive
- 1 Goblin Bombardment
- 1 Grist, the Hunger Tide
- 1 Minsc & Boo, Timeless Heroes
- 1 Orcish Bowmasters
- 1 Yawgmoth, Thran Physician
- 1 Altar of the Wretched // Wretched Bonemass
- 1 Constant Mists
- 1 Cultist of the Absolute
- 1 Dark Privilege
- 1 Gift of Doom
- 1 Goblin Chirurgeon
- 1 Kinzu of the Bleak Coven
- 1 Nightmare Shepherd
- 1 Rescue from the Underworld
- 1 Sylvan Safekeeper
- 1 Yarus, Roar of the Old Gods
- 1 Liliana, Dreadhorde General
- 1 The Meathook Massacre
- 1 Academy Manufactor
- 1 Birthing Pod
- 1 Diabolic Intent
- 1 Viscera Seer
- 1 Woe Strider
- 1 Xorn
- 1 Flesh-Eater Imp
- 1 Gobbling Ooze
- 1 Meren of Clan Nel Toth
- 1 Prossh, Skyraider of Kher
- 1 Rakdos, the Muscle
- 1 Razaketh, the Foulblooded
- 1 Sadistic Hypnotist
- 1 Sivriss, Nightmare Speaker
- 1 Thallid Omnivore
- 1 Undercity Scavenger
- 1 Venerated Stormsinger
- 1 Vindictive Vampire
- 1 Revel in Riches

### 12. Bracket 4, Prosper, Tome-Bound, treasure

Pool 210. 99 cards, 0 block findings, repaired false.

Judge: bracket 5, disagrees. This is a cEDH-caliber Prosper list: eight-plus Game Changers (Ad Nauseam, The One Ring, Imperial Seal, Vampiric/Demonic Tutor, Mana Vault, Mox Diamond, Chrome Mox, LED) plus Lotus Petal, Dark Ritual, Rite of Flame, Sol Ring and Gemstone Caverns for explosive turn 1-3 starts. The dual/fetchable premium manabase, free interaction (Deadly Rollick, Snuff Out, Pyroblast/REB), and Ad Nauseam/Peer into the Abyss burst-draw wincons mark it as a tuned competitive build rather than a theme deck. Even if only some pods call it Bracket 4, the fast mana density and tutor-to-win consistency place it at the top end.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 33 | 31 to 36 |  |  |
| `tapped_land` | 0 | 0 to 5 |  |  |
| `colorless_land` | 1 | 0 to 5 |  | Treasure Vault |
| `color_sources` | 1.46 | 0.9 or more |  | sources of requirement: B 29.75 of 19, R 27.75 of 19 |
| `avg_mana_value` | 2.42 | 1.6 to 3 |  | over 66 nonland cards |
| `ramp` | 13 | 10 to 16 |  |  |
| `draw` | 13 | 8 to 16 |  |  |
| `removal` | 10 | 6 to 14 |  |  |
| `wipe` | 2 | 0 to 5 |  |  |
| `interaction` | 9 | 6 to 16 |  |  |
| `tutor` | 3 | 2 or more |  | Demonic Tutor, Imperial Seal, Vampiric Tutor |
| `fast_mana` | 10 | 3 or more |  | Chrome Mox, Dark Ritual, Lion's Eye Diamond, Lotus Petal, Mana Vault, Mox Amber, Mox Diamond, Mox Opal, Rite of Flame, Sol Ring |
| `game_changer` | 9 | 4 or more |  | Ad Nauseam, Chrome Mox, Demonic Tutor, Imperial Seal, Lion's Eye Diamond, Mana Vault, Mox Diamond, The One Ring, Vampiric Tutor |
| `mana_turn_four` | 5.15 | 4.4 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.71 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.74 | -9 to 0.5 |  | the commander costs 4 and comes down on turn 3.26 on average |

Goldfish over 10000 hands: the commander on turn 3.26, 5.15 mana on turn four, and 0.71 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Ad Nauseam, Chrome Mox, Demonic Tutor, Imperial Seal, Lion's Eye Diamond, Mana Vault, Mox Diamond, The One Ring, Vampiric Tutor. Mass land denial none. Extra turns none. Combos 1.
- Prosper, Tome-Bound + Professional Face-Breaker + Sensei's Divining Top + Xorn (speed 5)

Findings:

- INFO `curve_summary`: average mana value 2.42 over 66 nonland cards

Cards:

- 1 Badlands
- 1 Blood Crypt
- 1 Cavern of Souls
- 1 City of Brass
- 1 Command Tower
- 1 Dragonskull Summit
- 1 Exotic Orchard
- 1 Forbidden Orchard
- 1 Gemstone Caverns
- 1 Graven Cairns
- 1 Haunted Ridge
- 1 Luxury Suite
- 1 Mana Confluence
- 1 Plaza of Heroes
- 1 Reflecting Pool
- 1 Spire of Industry
- 1 Sulfurous Springs
- 1 Tainted Peak
- 1 Treasure Vault
- 8 Swamp
- 6 Mountain
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
- 1 Pitiless Plunderer
- 1 Professional Face-Breaker
- 1 Goldspan Dragon
- 1 Ad Nauseam
- 1 Corrupted Conviction
- 1 Demand Answers
- 1 Faithless Looting
- 1 Hired Hexblade
- 1 Idol of Oblivion
- 1 Night's Whisper
- 1 Peer into the Abyss
- 1 Sensei's Divining Top
- 1 Skullclamp
- 1 Wheel of Fortune
- 1 Reckless Lackey
- 1 Village Rites
- 1 Abrade
- 1 Bedevil
- 1 Chaos Warp
- 1 Deadly Rollick
- 1 Dismember
- 1 Feed the Swarm
- 1 Go for the Throat
- 1 Infernal Grasp
- 1 Snuff Out
- 1 Withering Torment
- 1 The One Ring
- 1 Pyroblast
- 1 Red Elemental Blast
- 1 Tibalt's Trickery
- 1 Vexing Bauble
- 1 Lightning Greaves
- 1 Swiftfoot Boots
- 1 Not Dead After All
- 1 Undying Malice
- 1 Academy Manufactor
- 1 Contract Hero
- 1 Crossover Collaboration
- 1 Devour Intellect
- 1 Gilded Assault Cart
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
- 1 Imperial Seal
- 1 Vampiric Tutor
- 1 Knuckles the Echidna
- 1 Luck Bobblehead
- 1 Revel in Riches

### 13. Bracket 5, Kinnan, Bonder Prodigy, combo

Pool 322. 99 cards, 0 block findings, repaired false.

Judge: bracket 5, agrees. This is a textbook cEDH Kinnan deck: full fast-mana suite (LED, Moxen, Grim Monolith, Mana Vault, Lotus Petal, Gemstone Caverns), ten fetches with duals, ten-plus counterspells including free interaction, and Thassa's Oracle/Lab Man/Jace as the wincon. It packs numerous two-card infinite mana combos with Kinnan (Basalt Monolith, Devoted Druid, Voltaic Construct, Filigree Sages) that convert straight into a turn-two-or-three win. Roughly a dozen Game Changers and a pure competitive, non-thematic build put it squarely at bracket 5.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 31 | 27 to 33 |  |  |
| `tapped_land` | 0 | 0 to 2 |  |  |
| `colorless_land` | 0 | 0 to 6 |  |  |
| `color_sources` | 1.2 | 0.9 or more |  | sources of requirement: U 27.5 of 23, G 29.75 of 19 |
| `avg_mana_value` | 2.38 | 1.2 to 2.4 |  | over 68 nonland cards |
| `ramp` | 17 | 12 to 22 |  |  |
| `draw` | 11 | 8 to 18 |  |  |
| `removal` | 9 | 4 to 14 |  |  |
| `wipe` | 1 | 0 to 3 |  |  |
| `interaction` | 14 | 8 to 22 |  |  |
| `tutor` | 5 | 4 or more |  | Chakram Retriever, Intuition, Muddle the Mixture, Mystical Tutor, Primal Command |
| `fast_mana` | 10 | 6 or more |  | An Offer You Can't Refuse, Chrome Mox, Lion's Eye Diamond, Lotus Petal, Mana Vault, Mox Amber, Mox Diamond, Mox Opal, Sol Ring, Springleaf Drum |
| `game_changer` | 11 | 8 or more |  | Chrome Mox, Cyclonic Rift, Fierce Guardianship, Grim Monolith, Intuition, Lion's Eye Diamond, Mana Vault, Mox Diamond, Mystical Tutor, Rhystic Study, Thassa's Oracle |
| `mana_turn_four` | 5.19 | 4.6 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.68 | 0.65 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.41 | -9 to 0.3 |  | the commander costs 2 and comes down on turn 1.59 on average |

Goldfish over 10000 hands: the commander on turn 1.59, 5.19 mana on turn four, and 0.68 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Chrome Mox, Cyclonic Rift, Fierce Guardianship, Grim Monolith, Intuition, Lion's Eye Diamond, Mana Vault, Mox Diamond, Mystical Tutor, Rhystic Study, Thassa's Oracle. Mass land denial none. Extra turns none. Combos 5.
- Ioreth of the Healing House + Kiora's Follower (speed 5)
- Sakashima of a Thousand Faces + Ioreth of the Healing House (speed 5, two cards)
- Kinnan, Bonder Prodigy + Grim Monolith + Spark Double (speed 5, two cards)
- Kinnan, Bonder Prodigy + Grim Monolith + Sakashima of a Thousand Faces (speed 5, two cards)
- Kinnan, Bonder Prodigy + Basalt Monolith (speed 5, two cards)

Findings:

- INFO `curve_summary`: average mana value 2.38 over 68 nonland cards
- INFO `mana_pass`: the builder moved 4 cards of the mana base to bring the deck inside its power level

Cards:

- 8 Forest
- 5 Island
- 1 Breeding Pool
- 1 Command Tower
- 1 City of Brass
- 1 Mana Confluence
- 1 Yavimaya Coast
- 1 Dreamroot Cascade
- 1 Rejuvenating Springs
- 1 Hinterland Harbor
- 1 Misty Rainforest
- 1 Flooded Strand
- 1 Polluted Delta
- 1 Scalding Tarn
- 1 Verdant Catacombs
- 1 Windswept Heath
- 1 Wooded Foothills
- 1 Marsh Flats
- 1 Bloodstained Mire
- 1 Gemstone Caverns
- 1 Chrome Mox
- 1 Lion's Eye Diamond
- 1 Lotus Petal
- 1 Mana Vault
- 1 Mox Amber
- 1 Mox Diamond
- 1 Mox Opal
- 1 Sol Ring
- 1 Springleaf Drum
- 1 Basalt Monolith
- 1 Grim Monolith
- 1 Arbor Elf
- 1 Devoted Druid
- 1 Earthcraft
- 1 Gwenna, Eyes of Gaea
- 1 Kiora's Follower
- 1 Ioreth of the Healing House
- 1 Rhystic Study
- 1 Frantic Search
- 1 Cloud of Faeries
- 1 Cerulean Wisps
- 1 Refocus
- 1 Twitch
- 1 Pip-Boy 3000
- 1 Staff of Domination
- 1 Kitsa, Otterball Elite
- 1 Innocuous Researcher
- 1 Yavimaya Elder
- 1 An Offer You Can't Refuse
- 1 Fierce Guardianship
- 1 Flusterstorm
- 1 Muddle the Mixture
- 1 Narset's Reversal
- 1 Rewind
- 1 Shore Up
- 1 Sokka's Haiku
- 1 Unwind
- 1 Katara's Reversal
- 1 Legolas's Quick Reflexes
- 1 Magic Damper
- 1 Biosynthic Burst
- 1 Fleeting Reflection
- 1 Primal Command
- 1 Snap
- 1 Emerald Charm
- 1 Provoke
- 1 Bushwhack
- 1 Aggressive Biomancy
- 1 Apathy
- 1 Amphibian Downpour
- 1 Dream Tides
- 1 Intuition
- 1 Mystical Tutor
- 1 Cyclonic Rift
- 1 Filigree Sages
- 1 Voltaic Construct
- 1 Peregrine Drake
- 1 Breaching Hippocamp
- 1 Chakram Retriever
- 1 Dross Scorpion
- 1 Clever Impersonator
- 1 Phyrexian Metamorph
- 1 Sakashima of a Thousand Faces
- 1 Spark Double
- 1 Mirrorhall Mimic // Ghastly Mimicry
- 1 Thassa's Oracle
- 1 Laboratory Maniac
- 1 Jace, Wielder of Mysteries

### 14. Bracket 5, Yuriko, the Tiger's Shadow, ninjas

Pool 238. 99 cards, 0 block findings, repaired false.

Judge: bracket 5, agrees. This is a textbook cEDH Yuriko list: Underground Sea plus a full fast-mana suite (LED, Mox Diamond, Chrome Mox, Mana Vault, Dark Ritual, Culling the Weak), a dozen-plus free/cheap counterspells including Force of Will, Fierce Guardianship and Pact of Negation, and black tutors (Imperial Seal, Vampiric, Demonic, Mystical) into Ad Nauseam and Thassa's Oracle. It carries roughly fifteen marked Game Changers and can kill or combo out in the first few turns, far past Bracket 4's casual optimized range. The build is metagame-tuned competitive rather than thematic.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 30 | 27 to 33 |  |  |
| `tapped_land` | 1 | 0 to 2 |  | Choked Estuary |
| `colorless_land` | 0 | 0 to 6 |  |  |
| `color_sources` | 1.41 | 0.9 or more |  | sources of requirement: U 31.75 of 19, B 26.75 of 19 |
| `avg_mana_value` | 2.14 | 1.2 to 2.4 |  | over 69 nonland cards |
| `ramp` | 16 | 12 to 22 |  |  |
| `draw` | 13 | 8 to 18 |  |  |
| `removal` | 8 | 4 to 14 |  |  |
| `wipe` | 1 | 0 to 3 |  |  |
| `interaction` | 15 | 8 to 22 |  |  |
| `tutor` | 4 | 4 or more |  | Demonic Tutor, Imperial Seal, Mystical Tutor, Vampiric Tutor |
| `fast_mana` | 12 | 6 or more |  | Chrome Mox, Culling the Weak, Dark Ritual, Everflowing Chalice, Lion's Eye Diamond, Lotus Petal, Mana Vault, Mox Amber, Mox Diamond, Mox Opal, Sol Ring, Springleaf Drum |
| `game_changer` | 15 | 8 or more |  | Ad Nauseam, Chrome Mox, Cyclonic Rift, Demonic Tutor, Fierce Guardianship, Force of Will, Imperial Seal, Lion's Eye Diamond, Mana Vault, Mox Diamond, Mystical Tutor, Rhystic Study, Thassa's Oracle, The One Ring, Vampiric Tutor |
| `mana_turn_four` | 5.12 | 4.6 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.66 | 0.65 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.67 | -9 to 0.3 |  | the commander costs 3 and comes down on turn 2.33 on average |

Goldfish over 10000 hands: the commander on turn 2.33, 5.12 mana on turn four, and 0.66 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Ad Nauseam, Chrome Mox, Cyclonic Rift, Demonic Tutor, Fierce Guardianship, Force of Will, Imperial Seal, Lion's Eye Diamond, Mana Vault, Mox Diamond, Mystical Tutor, Rhystic Study, Thassa's Oracle, The One Ring, Vampiric Tutor. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 2.14 over 69 nonland cards
- INFO `mana_pass`: the builder moved 3 cards of the mana base to bring the deck inside its power level

Cards:

- 10 Island
- 5 Swamp
- 1 Underground Sea
- 1 Watery Grave
- 1 Morphic Pool
- 1 Drowned Catacomb
- 1 Choked Estuary
- 1 Gloomlake Verge
- 1 City of Brass
- 1 Mana Confluence
- 1 Command Tower
- 1 Exotic Orchard
- 1 Forbidden Orchard
- 1 Clearwater Pathway // Murkwater Pathway
- 1 Underground River
- 1 Secluded Courtyard
- 1 Unclaimed Territory
- 1 Chrome Mox
- 1 Mox Diamond
- 1 Mox Opal
- 1 Mox Amber
- 1 Lion's Eye Diamond
- 1 Mana Vault
- 1 Lotus Petal
- 1 Sol Ring
- 1 Springleaf Drum
- 1 Everflowing Chalice
- 1 Dark Ritual
- 1 Culling the Weak
- 1 Arcane Signet
- 1 Talisman of Dominance
- 1 Dimir Signet
- 1 Fellwar Stone
- 1 Ad Nauseam
- 1 Rhystic Study
- 1 Mystic Remora
- 1 Brainstorm
- 1 Gitaxian Probe
- 1 Ponder
- 1 Preordain
- 1 Sensei's Divining Top
- 1 Ingenious Infiltrator
- 1 Moon-Circuit Hacker
- 1 Ninja of the Deep Hours
- 1 Satoru, the Infiltrator
- 1 Windfall
- 1 Fierce Guardianship
- 1 Force of Will
- 1 Flusterstorm
- 1 Mental Misstep
- 1 Pact of Negation
- 1 Swan Song
- 1 Dispel
- 1 Negate
- 1 Counterspell
- 1 Force of Negation
- 1 Arcane Denial
- 1 Vexing Bauble
- 1 The One Ring
- 1 Kaito, Bane of Nightmares
- 1 Azra Smokeshaper
- 1 Chain of Vapor
- 1 Deadly Rollick
- 1 Dismember
- 1 Feed the Swarm
- 1 Snuff Out
- 1 Pongify
- 1 Rapid Hybridization
- 1 Reality Shift
- 1 Fallen Shinobi
- 1 Foot Mystic
- 1 Futurist Operative
- 1 Kotose, the Silent Spider
- 1 Moonblade Shinobi
- 1 Mukotai Ambusher
- 1 Ninja of the New Moon
- 1 Okiba-Gang Shinobi
- 1 Sakashima's Student
- 1 Silver-Fur Master
- 1 Mist-Syndicate Naga
- 1 Cyclonic Rift
- 1 Demonic Tutor
- 1 Imperial Seal
- 1 Mystical Tutor
- 1 Vampiric Tutor
- 1 Thassa's Oracle

### 15. Bracket 5, Najeela, the Blade-Blossom, warriors combat

Pool 338. 99 cards, 0 block findings, repaired true (deck_size, profile_off_band).

Judge: bracket 4, disagrees. This is a 14+ Game Changer pile with LED + Underworld Breach + Thassa's Oracle, Najeela's own infinite combat loop, a full tutor suite (Imperial Seal, Vampiric, Demonic, Mystical, Gamble, Wishclaw), and heavy fast mana (Moxen, Mana Vault, Dark Ritual, Lotus Petal), so it can win by turn three or four. Free interaction like Fierce Guardianship and An Offer You Can't Refuse protects the kill. It falls short of true cEDH only because the mana base is a clunky basic-heavy build with a warrior-tribal beatdown shell rather than a purely optimized combo list, but it's far past Bracket 3.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 31 | 27 to 33 |  |  |
| `tapped_land` | 2 | 0 to 2 |  | Base Camp, Path of Ancestry |
| `colorless_land` | 0 | 0 to 6 |  |  |
| `color_sources` | 0.75 | 0.9 or more | yes | sources of requirement: W 14.25 of 19, U 15.25 of 19, B 14.25 of 19, R 14.25 of 19, G 19.75 of 26 |
| `avg_mana_value` | 2.4 | 1.2 to 2.4 |  | over 68 nonland cards |
| `ramp` | 17 | 12 to 22 |  |  |
| `draw` | 8 | 8 to 18 |  |  |
| `removal` | 6 | 4 to 14 |  |  |
| `wipe` | 1 | 0 to 3 |  |  |
| `interaction` | 13 | 8 to 22 |  |  |
| `tutor` | 7 | 4 or more |  | Demonic Tutor, Enlightened Tutor, Gamble, Imperial Seal, Mystical Tutor, Vampiric Tutor, Wishclaw Talisman |
| `fast_mana` | 11 | 6 or more |  | An Offer You Can't Refuse, Chrome Mox, Dark Ritual, Lion's Eye Diamond, Lotus Petal, Mana Vault, Mox Amber, Mox Diamond, Mox Opal, Sol Ring, Springleaf Drum |
| `game_changer` | 16 | 8 or more |  | Chrome Mox, Demonic Tutor, Enlightened Tutor, Fierce Guardianship, Gamble, Imperial Seal, Lion's Eye Diamond, Mana Vault, Mox Diamond, Mystical Tutor, Orcish Bowmasters, Rhystic Study, Smothering Tithe, Thassa's Oracle, Underworld Breach, Vampiric Tutor |
| `mana_turn_four` | 4.91 | 4.6 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.68 | 0.65 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.6 | -9 to 0.3 |  | the commander costs 3 and comes down on turn 2.4 on average |

Goldfish over 10000 hands: the commander on turn 2.4, 4.91 mana on turn four, and 0.68 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Chrome Mox, Demonic Tutor, Enlightened Tutor, Fierce Guardianship, Gamble, Imperial Seal, Lion's Eye Diamond, Mana Vault, Mox Diamond, Mystical Tutor, Orcish Bowmasters, Rhystic Study, Smothering Tithe, Thassa's Oracle, Underworld Breach, Vampiric Tutor. Mass land denial none. Extra turns none. Combos 1.
- Najeela, the Blade-Blossom + Professional Face-Breaker (speed 5, two cards)

Findings:

- INFO `curve_summary`: average mana value 2.40 over 68 nonland cards
- WARN `profile_off_band`: the worst color's share of the sources it needs is 0.75, and bracket 5 wants 0.9 or more (sources of requirement: W 14.25 of 19, U 15.25 of 19, B 14.25 of 19, R 14.25 of 19, G 19.75 of 26)
- INFO `mana_pass`: the builder moved 3 cards of the mana base to bring the deck inside its power level

Cards:

- 1 Command Tower
- 1 Exotic Orchard
- 1 Base Camp
- 1 Path of Ancestry
- 1 Blood Crypt
- 1 Breeding Pool
- 1 Godless Shrine
- 1 Hallowed Fountain
- 1 Steam Vents
- 1 Watery Grave
- 11 Forest
- 3 Plains
- 2 Island
- 2 Swamp
- 3 Mountain
- 1 Chrome Mox
- 1 Lion's Eye Diamond
- 1 Lotus Petal
- 1 Mana Vault
- 1 Mox Amber
- 1 Mox Diamond
- 1 Mox Opal
- 1 Sol Ring
- 1 Springleaf Drum
- 1 Smothering Tithe
- 1 Ragavan, Nimble Pilferer
- 1 Professional Face-Breaker
- 1 Grand Warlord Radha
- 1 Heronblade Elite
- 1 Burakos, Party Leader
- 1 Rhystic Study
- 1 Akki Ronin
- 1 Azra Oddsmaker
- 1 Mindblade Render
- 1 Nimble Trapfinder
- 1 Oakhame Adversary
- 1 Raiders' Spoils
- 1 Samut, Vizier of Naktamun
- 1 An Offer You Can't Refuse
- 1 Fierce Guardianship
- 1 Concerted Defense
- 1 Eladamri, Lord of Leaves
- 1 Ezuri, Renegade Leader
- 1 Kashi-Tribe Elite
- 1 Linvala, Shield of Sea Gate
- 1 Selfless Safewright
- 1 Selfless Samurai
- 1 Multiclass Baldric
- 1 Akiri, Fearless Voyager
- 1 Winota, Joiner of Forces
- 1 Rhys, the Evermore
- 1 Orcish Bowmasters
- 1 Accursed Marauder
- 1 Cacophony Scamp
- 1 Dreadhorde Butcher
- 1 Fleshbag Marauder
- 1 Goblin Cratermaker
- 1 Combat Celebrant
- 1 Bramblewood Paragon
- 1 Asari Captain
- 1 Aspiring Champion
- 1 Aysen Crusader
- 1 Bull-Rush Bruiser
- 1 Doomskar Warrior
- 1 Heiko Yamazaki, the General
- 1 Kabira Outrider
- 1 Proud Wildbonder
- 1 Raiyuu, Storm's Edge
- 1 Surrak, the Hunt Caller
- 1 Thrasher Brute
- 1 Vengeful Firebrand
- 1 Goblin Chainwhirler
- 1 Demonic Tutor
- 1 Enlightened Tutor
- 1 Imperial Seal
- 1 Mystical Tutor
- 1 Vampiric Tutor
- 1 Gamble
- 1 Wishclaw Talisman
- 1 Underworld Breach
- 1 Thassa's Oracle
- 1 Shang-Chi, Master of Kung Fu
- 1 Dark Ritual

