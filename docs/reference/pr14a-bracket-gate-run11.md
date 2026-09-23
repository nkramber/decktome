# PR-14A bracket gate

Run date: 2026-09-23. Card snapshot: 2026-09-04.

Verdict: PASS. 3 of 3 decks passed every block check, 3 sat in every band, 3 held no content violation, and the judge agreed with the bracket on 3 of 3 (100 percent, the bar is 80).

## Summary

| Measure | Value |
|---|---|
| Prompts | 3 |
| Decks returned | 3 |
| Decks with no block finding | 3 |
| Decks in every band | 3 |
| Decks with no content violation | 3 |
| Decks the endpoint did not check | 0 |
| Decks that needed a repair turn | 0 |
| Decks judged | 3 |
| Judge agreed with the bracket | 3 |
| Judge errors | 0 |
| Errors | 0 |
| Prompt version | 16 |
| Calls | 6 |
| Cost | $0.2967 |
| Time | 225 seconds |

## Run

- Suite `bracket`, run `pr14a-bracket-gate-run11`, on 2026-09-23, commit `a2c8090`.
- Partial run over `13,14,15`. It reads the bars of its own items, and it never stands for the suite.
- Roles: generate on `gpt-5.6-terra` (openai, effort medium), judge on `claude-opus-5` (anthropic, effort medium), repair on `gpt-5.6-terra` (openai, effort medium).
- Versions: card snapshot 2026-09-04, bracket_judge prompt version 3, generate prompt version 16.
- Calls: 6. Cost: $0.2967. Time: 225 seconds.

## Off-band features by key

Each row counts the decks whose feature sat outside its band after the repair turns.

| Feature | Decks off band |
|---|---|
| none | 0 |

## Decks

### 13. Bracket 5, Kinnan, Bonder Prodigy, combo

Pool 321. 99 cards, 0 block findings, repaired false.

Judge: bracket 5, agrees. This is a textbook cEDH Kinnan build: Dual lands, fetches, Lion's Eye Diamond, Mox Diamond/Opal/Chrome Mox, Ancient Tomb and Gaea's Cradle, with a dozen-plus free and cheap counterspells (Force of Will, Fierce Guardianship, Pact of Negation, Flusterstorm). It packs 11 Game Changers and multiple zero-to-two-card infinite mana lines (Basalt Monolith, Freed from the Real, Spark Double) that convert into Walking Ballista/Staff/Thassa's Oracle wins on early turns. It is metagame-tuned rather than themed, so it plays at the highest bracket.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 30 | 27 to 33 |  |  |
| `tapped_land` | 0 | 0 to 2 |  |  |
| `colorless_land` | 1 | 0 to 6 |  | Ancient Tomb |
| `color_sources` | 1.46 | 0.9 or more |  | sources of requirement: U 33.5 of 23, G 32 of 19 |
| `fixing_land` | 17 | 9 or more |  | Barkchannel Pathway // Tidechannel Pathway, Botanical Sanctum, Breeding Pool, City of Brass, Command Tower, Dreamroot Cascade, Exotic Orchard, Forbidden Orchard, Hinterland Harbor, Mana Confluence, Misty Rainforest, Prismatic Vista, Rejuvenating Springs, Sodden Verdure, Tropical Island, Turbulent Wilderness, Waterlogged Grove |
| `avg_mana_value` | 2.38 | 1.2 to 2.4 |  | over 69 nonland cards |
| `ramp` | 17 | 12 to 22 |  |  |
| `draw` | 13 | 8 to 18 |  |  |
| `removal` | 9 | 4 to 14 |  |  |
| `wipe` | 1 | 0 to 3 |  |  |
| `interaction` | 15 | 8 to 22 |  |  |
| `tutor` | 5 | 4 or more |  | Alphinaud Leveilleur, Arcum Dagsson, Kuldotha Forgemaster, Muddle the Mixture, Tezzeret the Seeker |
| `fast_mana` | 10 | 6 or more |  | An Offer You Can't Refuse, Chrome Mox, Lion's Eye Diamond, Lotus Petal, Mana Vault, Mox Diamond, Mox Opal, Paradise Mantle, Sol Ring, Springleaf Drum |
| `game_changer` | 13 | 8 or more |  | Ancient Tomb, Chrome Mox, Consecrated Sphinx, Cyclonic Rift, Fierce Guardianship, Force of Will, Gaea's Cradle, Grim Monolith, Lion's Eye Diamond, Mana Vault, Mox Diamond, Rhystic Study, Thassa's Oracle |
| `finisher` | 1 | 1 or more |  | Thassa's Oracle |
| `mana_turn_four` | 5.42 | 4.6 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.66 | 0.65 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.33 | -9 to 0.3 |  | the commander costs 2 and comes down on turn 1.67 on average |

Goldfish over 10000 hands: the commander on turn 1.67, 5.42 mana on turn four, and 0.66 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Ancient Tomb, Gaea's Cradle, Chrome Mox, Grim Monolith, Lion's Eye Diamond, Mana Vault, Mox Diamond, Consecrated Sphinx, Rhystic Study, Fierce Guardianship, Force of Will, Thassa's Oracle, Cyclonic Rift. Mass land denial none. Extra turns none. Combos 5.
- Kinnan, Bonder Prodigy + Delighted Halfling + Freed from the Real (speed 5)
- Kinnan, Bonder Prodigy + Paradise Mantle + Freed from the Real (speed 5, two cards)
- Kinnan, Bonder Prodigy + Grim Monolith + Spark Double (speed 5, two cards)
- Kinnan, Bonder Prodigy + Birds of Paradise + Freed from the Real (speed 5, two cards)
- Kinnan, Bonder Prodigy + Basalt Monolith (speed 5, two cards)

Findings:

- INFO `curve_summary`: average mana value 2.38 over 69 nonland cards

Cards:

- 1 Ancient Tomb
- 1 Barkchannel Pathway // Tidechannel Pathway
- 1 Boseiju, Who Endures
- 1 Botanical Sanctum
- 1 Breeding Pool
- 1 Cephalid Coliseum
- 1 City of Brass
- 1 Command Tower
- 1 Dreamroot Cascade
- 1 Exotic Orchard
- 1 Flooded Strand
- 1 Forbidden Orchard
- 1 Forest
- 1 Gaea's Cradle
- 1 Gemstone Caverns
- 1 Hinterland Harbor
- 1 Island
- 1 Mana Confluence
- 1 Misty Rainforest
- 1 Otawara, Soaring City
- 1 Polluted Delta
- 1 Prismatic Vista
- 1 Rejuvenating Springs
- 1 Scalding Tarn
- 1 Sodden Verdure
- 1 Tropical Island
- 1 Turbulent Wilderness
- 1 Waterlogged Grove
- 1 Windswept Heath
- 1 Wooded Foothills
- 1 Arcane Signet
- 1 Basalt Monolith
- 1 Birds of Paradise
- 1 Chrome Mox
- 1 Delighted Halfling
- 1 Fellwar Stone
- 1 Grim Monolith
- 1 Lion's Eye Diamond
- 1 Lotus Petal
- 1 Mana Vault
- 1 Mox Diamond
- 1 Mox Opal
- 1 Paradise Mantle
- 1 Sol Ring
- 1 Springleaf Drum
- 1 Talisman of Curiosity
- 1 Tezzeret the Seeker
- 1 Alphinaud Leveilleur
- 1 Borne Upon a Wind
- 1 Cloud of Faeries
- 1 Consecrated Sphinx
- 1 Faerie Mastermind
- 1 Frantic Search
- 1 Kitsa, Otterball Elite
- 1 Lórien Revealed
- 1 Mystic Remora
- 1 Rhystic Study
- 1 Silent Hallcreeper
- 1 Staff of Domination
- 1 Sylvan Library
- 1 Arcum Dagsson
- 1 Blasting Station
- 1 Chain of Vapor
- 1 Emerald Charm
- 1 Force of Vigor
- 1 Into the Flood Maw
- 1 Snap
- 1 Ty Lee, Chi Blocker
- 1 Walking Ballista
- 1 An Offer You Can't Refuse
- 1 Dispel
- 1 Fierce Guardianship
- 1 Flusterstorm
- 1 Force of Negation
- 1 Force of Will
- 1 Legolas's Quick Reflexes
- 1 Mental Misstep
- 1 Muddle the Mixture
- 1 Narset's Reversal
- 1 Pact of Negation
- 1 Spellskite
- 1 Subtlety
- 1 Swan Song
- 1 Veil of Summer
- 1 Altered Ego
- 1 Auton Soldier
- 1 Bramble Sovereign
- 1 Clever Impersonator
- 1 Clone
- 1 Faerie Artisans
- 1 High Fae Trickster
- 1 Kuldotha Forgemaster
- 1 Malleable Impostor
- 1 Peregrine Drake
- 1 Phyrexian Metamorph
- 1 Spark Double
- 1 Freed from the Real
- 1 Thassa's Oracle
- 1 Cyclonic Rift

### 14. Bracket 5, Yuriko, the Tiger's Shadow, ninjas

Pool 295. 99 cards, 0 block findings, repaired false.

Judge: bracket 5, agrees. This is a cEDH Yuriko list: dual/fetch mana base with Underground Sea, heavy fast mana (Lion's Eye Diamond, Mana Vault, Grim Monolith, multiple Moxen, rituals), a full suite of free counterspells (Force of Will, Force of Negation, Fierce Guardianship, Pact of Negation, Daze), and top-tier tutors into the Thassa's Oracle + Demonic Consultation two-card win for zero extra mana. It has roughly a dozen and a half Game Changers, far past Bracket 3/4 casual expectations, and is built to win on the earliest turns rather than express a theme.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 33 | 27 to 33 |  |  |
| `tapped_land` | 2 | 0 to 2 |  | Agadeem's Awakening // Agadeem, the Undercrypt, Choked Estuary |
| `colorless_land` | 1 | 0 to 6 |  | Ancient Tomb |
| `color_sources` | 1.55 | 0.9 or more |  | sources of requirement: U 35.75 of 23, B 33.75 of 19 |
| `fixing_land` | 19 | 9 or more |  | Choked Estuary, City of Brass, Clearwater Pathway // Murkwater Pathway, Command Tower, Darkslick Shores, Drowned Catacomb, Exotic Orchard, Forbidden Orchard, Gloomlake Verge, Mana Confluence, Morphic Pool, Polluted Delta, Prismatic Vista, Shipwreck Marsh, Sunken Hollow, Tarnished Citadel, Underground River, Underground Sea, Watery Grave |
| `avg_mana_value` | 2.3 | 1.2 to 2.4 |  | over 66 nonland cards |
| `ramp` | 17 | 12 to 22 |  |  |
| `draw` | 12 | 8 to 18 |  |  |
| `removal` | 9 | 4 to 14 |  |  |
| `wipe` | 1 | 0 to 3 |  |  |
| `interaction` | 10 | 8 to 22 |  |  |
| `tutor` | 5 | 4 or more |  | Demonic Tutor, Higure, the Still Wind, Imperial Seal, Mystical Tutor, Vampiric Tutor |
| `fast_mana` | 12 | 6 or more |  | An Offer You Can't Refuse, Chrome Mox, Culling the Weak, Dark Ritual, Lion's Eye Diamond, Lotus Petal, Mana Vault, Mox Amber, Mox Diamond, Mox Opal, Sol Ring, Springleaf Drum |
| `game_changer` | 17 | 8 or more |  | Ancient Tomb, Chrome Mox, Cyclonic Rift, Demonic Tutor, Fierce Guardianship, Force of Will, Grim Monolith, Imperial Seal, Lion's Eye Diamond, Mana Vault, Mox Diamond, Mystical Tutor, Orcish Bowmasters, Rhystic Study, Thassa's Oracle, The One Ring, Vampiric Tutor |
| `finisher` | 1 | 1 or more |  | Thassa's Oracle |
| `mana_turn_four` | 5.45 | 4.6 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.71 | 0.65 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.67 | -9 to 0.3 |  | the commander costs 3 and comes down on turn 2.33 on average |

Goldfish over 10000 hands: the commander on turn 2.33, 5.45 mana on turn four, and 0.71 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Ancient Tomb, Chrome Mox, Grim Monolith, Lion's Eye Diamond, Mana Vault, Mox Diamond, Rhystic Study, Fierce Guardianship, Force of Will, The One Ring, Orcish Bowmasters, Cyclonic Rift, Thassa's Oracle, Demonic Tutor, Imperial Seal, Mystical Tutor, Vampiric Tutor. Mass land denial none. Extra turns none. Combos 1.
- Demonic Consultation + Thassa's Oracle (speed 5, two cards)

Findings:

- INFO `curve_summary`: average mana value 2.30 over 66 nonland cards
- INFO `mana_pass`: the builder moved 4 cards of the mana base to bring the deck inside its power level
- INFO `basics_added`: the list was 1 card short, so the builder added 1 basic land

Cards:

- 1 Agadeem's Awakening // Agadeem, the Undercrypt
- 1 Ancient Tomb
- 1 Bloodstained Mire
- 1 Cephalid Coliseum
- 1 City of Brass
- 1 Clearwater Pathway // Murkwater Pathway
- 1 Choked Estuary
- 1 Command Tower
- 1 Darkslick Shores
- 1 Drowned Catacomb
- 1 Exotic Orchard
- 1 Flooded Strand
- 1 Gloomlake Verge
- 2 Island
- 1 Mana Confluence
- 1 Marsh Flats
- 1 Misty Rainforest
- 1 Morphic Pool
- 1 Multiversal Passage
- 1 Otawara, Soaring City
- 1 Polluted Delta
- 1 Prismatic Vista
- 1 Scalding Tarn
- 1 Shipwreck Marsh
- 1 Swamp
- 1 Underground River
- 1 Underground Sea
- 1 Watery Grave
- 1 Tarnished Citadel
- 1 Verdant Catacombs
- 1 Arcane Signet
- 1 Chrome Mox
- 1 Culling the Weak
- 1 Dark Ritual
- 1 Dimir Signet
- 1 Fellwar Stone
- 1 Grim Monolith
- 1 Lion's Eye Diamond
- 1 Lotus Petal
- 1 Mana Drain
- 1 Mana Vault
- 1 Mox Amber
- 1 Mox Diamond
- 1 Mox Opal
- 1 Sol Ring
- 1 Springleaf Drum
- 1 Talisman of Dominance
- 1 Brainstorm
- 1 Gitaxian Probe
- 1 Gush
- 1 Ingenious Infiltrator
- 1 Mystic Remora
- 1 Moon-Circuit Hacker
- 1 Ninja of the Deep Hours
- 1 Rhystic Study
- 1 Scroll Rack
- 1 Sensei's Divining Top
- 1 Satoru, the Infiltrator
- 1 Treasure Cruise
- 1 An Offer You Can't Refuse
- 1 Daze
- 1 Fierce Guardianship
- 1 Flusterstorm
- 1 Force of Negation
- 1 Force of Will
- 1 Mental Misstep
- 1 Pact of Negation
- 1 Swan Song
- 1 The One Ring
- 1 Chain of Vapor
- 1 Deadly Rollick
- 1 Dismember
- 1 Feed the Swarm
- 1 Orcish Bowmasters
- 1 Pongify
- 1 Rapid Hybridization
- 1 Reality Shift
- 1 Snuff Out
- 1 Cyclonic Rift
- 1 Fallen Shinobi
- 1 Futurist Operative
- 1 Higure, the Still Wind
- 1 Ink-Eyes, Servant of Oni
- 1 Kotose, the Silent Spider
- 1 Moonblade Shinobi
- 1 Mukotai Ambusher
- 1 Okiba-Gang Shinobi
- 1 Sakashima's Student
- 1 Silent-Blade Oni
- 1 Thassa's Oracle
- 1 Changeling Outcast
- 1 Demonic Consultation
- 1 Demonic Tutor
- 1 Imperial Seal
- 1 Mystical Tutor
- 1 Vampiric Tutor
- 1 Forbidden Orchard
- 1 Sunken Hollow

### 15. Bracket 5, Najeela, the Blade-Blossom, warriors combat

Pool 335. 99 cards, 0 block findings, repaired false.

Judge: bracket 5, agrees. This is a textbook cEDH Najeela list: dual/fetch mana base, Ancient Tomb, Lion's Eye Diamond, Moxen, Jeska's Will, and a dense free-interaction suite (Force of Will, Fierce Guardianship, Pact of Negation, Flusterstorm). It packs ~15 Game Changers, the full unrestricted tutor package, and a zero-mana Thassa's Oracle + Demonic Consultation win, backed by Najeela's infinite combat line. It's optimized for winning as early as possible in a competitive metagame, not for any theme.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 30 | 27 to 33 |  |  |
| `tapped_land` | 0 | 0 to 2 |  |  |
| `colorless_land` | 1 | 0 to 6 |  | Ancient Tomb |
| `color_sources` | 1.2 | 0.9 or more |  | sources of requirement: W 25 of 19, U 28 of 19, B 25 of 19, R 24 of 19, G 22.75 of 19 |
| `fixing_land` | 25 | 21 or more |  | Arid Mesa, Badlands, Bayou, Bloodstained Mire, Breeding Pool, City of Brass, Command Tower, Flooded Strand, Hallowed Fountain, Mana Confluence, Marsh Flats, Misty Rainforest, Polluted Delta, Sacred Foundry, Scalding Tarn, Scrubland, Steam Vents, Temple Garden, Tropical Island, Tundra, Underground Sea, Verdant Catacombs, Volcanic Island, Windswept Heath, Wooded Foothills |
| `avg_mana_value` | 2.04 | 1.2 to 2.4 |  | over 69 nonland cards |
| `ramp` | 17 | 12 to 22 |  |  |
| `draw` | 11 | 8 to 18 |  |  |
| `removal` | 9 | 4 to 14 |  |  |
| `wipe` | 1 | 0 to 3 |  |  |
| `interaction` | 14 | 8 to 22 |  |  |
| `tutor` | 4 | 4 or more |  | Demonic Tutor, Enlightened Tutor, Mystical Tutor, Vampiric Tutor |
| `fast_mana` | 9 | 6 or more |  | An Offer You Can't Refuse, Chrome Mox, Culling the Weak, Lion's Eye Diamond, Lotus Petal, Mana Vault, Mox Diamond, Mox Opal, Sol Ring |
| `game_changer` | 17 | 8 or more |  | Ad Nauseam, Ancient Tomb, Chrome Mox, Cyclonic Rift, Demonic Tutor, Enlightened Tutor, Fierce Guardianship, Force of Will, Jeska's Will, Lion's Eye Diamond, Mana Vault, Mox Diamond, Mystical Tutor, Orcish Bowmasters, Rhystic Study, Thassa's Oracle, Vampiric Tutor |
| `finisher` | 1 | 1 or more |  | Thassa's Oracle |
| `mana_turn_four` | 5.29 | 4.6 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.66 | 0.65 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.68 | -9 to 0.3 |  | the commander costs 3 and comes down on turn 2.32 on average |

Goldfish over 10000 hands: the commander on turn 2.32, 5.29 mana on turn four, and 0.66 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Ancient Tomb, Chrome Mox, Lion's Eye Diamond, Mana Vault, Mox Diamond, Jeska's Will, Ad Nauseam, Rhystic Study, Fierce Guardianship, Force of Will, Orcish Bowmasters, Demonic Tutor, Enlightened Tutor, Vampiric Tutor, Mystical Tutor, Thassa's Oracle, Cyclonic Rift. Mass land denial none. Extra turns none. Combos 1.
- Demonic Consultation + Thassa's Oracle (speed 5, two cards)

Findings:

- INFO `curve_summary`: average mana value 2.04 over 69 nonland cards
- INFO `mana_pass`: the builder moved 2 cards of the mana base to bring the deck inside its power level

Cards:

- 1 Arid Mesa
- 1 Bloodstained Mire
- 1 Flooded Strand
- 1 Marsh Flats
- 1 Misty Rainforest
- 1 Polluted Delta
- 1 Scalding Tarn
- 1 Verdant Catacombs
- 1 Windswept Heath
- 1 Wooded Foothills
- 1 Badlands
- 1 Bayou
- 1 Breeding Pool
- 1 Hallowed Fountain
- 1 Sacred Foundry
- 1 Steam Vents
- 1 Temple Garden
- 1 Tropical Island
- 1 Tundra
- 1 Underground Sea
- 1 Volcanic Island
- 1 Command Tower
- 1 City of Brass
- 1 Mana Confluence
- 1 Ancient Tomb
- 1 Gemstone Caverns
- 1 Forest
- 1 Island
- 1 Swamp
- 1 Chrome Mox
- 1 Lion's Eye Diamond
- 1 Lotus Petal
- 1 Mana Vault
- 1 Mox Diamond
- 1 Mox Opal
- 1 Sol Ring
- 1 Arcane Signet
- 1 Fellwar Stone
- 1 Birds of Paradise
- 1 Deathrite Shaman
- 1 Delighted Halfling
- 1 Ignoble Hierarch
- 1 Noble Hierarch
- 1 Bloom Tender
- 1 Culling the Weak
- 1 Jeska's Will
- 1 Ad Nauseam
- 1 Mystic Remora
- 1 Rhystic Study
- 1 Esper Sentinel
- 1 Archivist of Oghma
- 1 Faerie Mastermind
- 1 Borne Upon a Wind
- 1 Wheel of Fortune
- 1 Kutzil, Malamet Exemplar
- 1 Samut, Vizier of Naktamun
- 1 Mask of Memory
- 1 An Offer You Can't Refuse
- 1 Fierce Guardianship
- 1 Flusterstorm
- 1 Force of Negation
- 1 Force of Will
- 1 Mental Misstep
- 1 Pact of Negation
- 1 Pyroblast
- 1 Red Elemental Blast
- 1 Swan Song
- 1 Veil of Summer
- 1 Vexing Bauble
- 1 Skrelv, Defector Mite
- 1 Legolas's Quick Reflexes
- 1 Abrupt Decay
- 1 Chain of Vapor
- 1 Deadly Rollick
- 1 Fire Covenant
- 1 Into the Flood Maw
- 1 Orcish Bowmasters
- 1 Snap
- 1 Swords to Plowshares
- 1 Umezawa's Jitte
- 1 Demonic Tutor
- 1 Enlightened Tutor
- 1 Vampiric Tutor
- 1 Mystical Tutor
- 1 Demonic Consultation
- 1 Alexios, Deimos of Kosmos
- 1 Auton Soldier
- 1 Bruenor Battlehammer
- 1 Herald of Secret Streams
- 1 Ivy Lane Denizen
- 1 Krenko, Mob Boss
- 1 Lys Alana Huntmaster
- 1 Maja, Bretagard Protector
- 1 Roaming Throne
- 1 Toph, Hardheaded Teacher
- 1 Zealous Conscripts
- 1 Thassa's Oracle
- 1 Cyclonic Rift
- 1 Scrubland

