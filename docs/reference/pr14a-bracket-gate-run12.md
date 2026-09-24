# PR-14A bracket gate

Run date: 2026-09-23. Card snapshot: 2026-09-04.

Verdict: FAIL. 1 of 1 decks passed every block check, 0 sat in every band, 1 held no content violation, and the judge agreed with the bracket on 1 of 1 (100 percent, the bar is 80).

## Summary

| Measure | Value |
|---|---|
| Prompts | 1 |
| Decks returned | 1 |
| Decks with no block finding | 1 |
| Decks in every band | 0 |
| Decks with no content violation | 1 |
| Decks the endpoint did not check | 0 |
| Decks that needed a repair turn | 0 |
| Decks judged | 1 |
| Judge agreed with the bracket | 1 |
| Judge errors | 0 |
| Errors | 0 |
| Prompt version | 16 |
| Calls | 2 |
| Cost | $0.1082 |
| Time | 69 seconds |

## Run

- Suite `bracket`, run `pr14a-bracket-gate-run12`, on 2026-09-23, commit `2b2b466`.
- Partial run over `15`. It reads the bars of its own items, and it never stands for the suite.
- Roles: generate on `gpt-5.6-terra` (openai, effort medium), judge on `claude-opus-5` (anthropic, effort medium), repair on `gpt-5.6-terra` (openai, effort medium).
- Versions: card snapshot 2026-09-04, bracket_judge prompt version 3, generate prompt version 16.
- Calls: 2. Cost: $0.1082. Time: 69 seconds.

## Off-band features by key

Each row counts the decks whose feature sat outside its band after the repair turns.

| Feature | Decks off band |
|---|---|
| `hands_two_to_four_lands` | 1 |

## Decks

### 15. Bracket 5, Najeela, the Blade-Blossom, warriors combat

Pool 335. 99 cards, 0 block findings, repaired false.

Judge: bracket 5, agrees. This is a full-power Najeela build with Ancient Tomb, Gaea's Cradle, dual lands, fetches, and a pile of fast mana (LED, Mox Diamond/Opal/Amber, Chrome Mox, Mana Vault, Dark Ritual, Culling the Weak) backed by free counterspells like Force of Will, Fierce Guardianship, Pact of Negation and Mindbreak Trap. It carries a dozen-plus marked Game Changers plus Ad Nauseam, The One Ring and efficient tutors, with Najeela's own infinite-combat kill available off Zealous Conscripts/Chord toolbox. The card choices are metagame-driven rather than thematic, so it plays as cEDH.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 27 | 27 to 33 |  |  |
| `tapped_land` | 0 | 0 to 2 |  |  |
| `colorless_land` | 1 | 0 to 6 |  | Ancient Tomb |
| `color_sources` | 1.14 | 0.9 or more |  | sources of requirement: W 23.25 of 19, U 23.25 of 19, B 23 of 19, R 22 of 19, G 21.75 of 19 |
| `fixing_land` | 25 | 21 or more |  | Arid Mesa, Badlands, Bayou, Bloodstained Mire, Breeding Pool, Cavern of Souls, City of Brass, Command Tower, Exotic Orchard, Flooded Strand, Godless Shrine, Hallowed Fountain, Mana Confluence, Marsh Flats, Misty Rainforest, Plateau, Polluted Delta, Savannah, Scrubland, Taiga, Tarnished Citadel, Tropical Island, Tundra, Underground Sea, Volcanic Island |
| `avg_mana_value` | 2.32 | 1.2 to 2.4 |  | over 72 nonland cards |
| `ramp` | 17 | 12 to 22 |  |  |
| `draw` | 13 | 8 to 18 |  |  |
| `removal` | 9 | 4 to 14 |  |  |
| `wipe` | 1 | 0 to 3 |  |  |
| `interaction` | 15 | 8 to 22 |  |  |
| `tutor` | 4 | 4 or more |  | Chord of Calling, Demonic Tutor, Fighter Class, Tezzeret, Cruel Captain |
| `fast_mana` | 12 | 6 or more |  | An Offer You Can't Refuse, Chrome Mox, Culling the Weak, Dark Ritual, Lion's Eye Diamond, Lotus Petal, Mana Vault, Mox Amber, Mox Diamond, Mox Opal, Rite of Flame, Sol Ring |
| `game_changer` | 14 | 8 or more |  | Ad Nauseam, Ancient Tomb, Chrome Mox, Cyclonic Rift, Demonic Tutor, Fierce Guardianship, Force of Will, Gaea's Cradle, Lion's Eye Diamond, Mana Vault, Mox Diamond, Orcish Bowmasters, Rhystic Study, The One Ring |
| `finisher` | 1 | 1 or more |  | Angel of Destiny |
| `mana_turn_four` | 5.18 | 4.6 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.6 | 0.65 or more | yes | share of first seven-card hands |
| `commander_turn_over_mv` | -0.72 | -9 to 0.3 |  | the commander costs 3 and comes down on turn 2.28 on average |

Goldfish over 10000 hands: the commander on turn 2.28, 5.18 mana on turn four, and 0.6 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Ancient Tomb, Gaea's Cradle, Chrome Mox, Mox Diamond, Mana Vault, Lion's Eye Diamond, Ad Nauseam, Rhystic Study, Force of Will, Fierce Guardianship, The One Ring, Orcish Bowmasters, Demonic Tutor, Cyclonic Rift. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 2.32 over 72 nonland cards
- WARN `profile_off_band`: the share of opening hands with two to four lands is 0.6, and bracket 5 wants 0.65 or more (share of first seven-card hands)

Cards:

- 1 Ancient Tomb
- 1 Gaea's Cradle
- 1 Command Tower
- 1 City of Brass
- 1 Mana Confluence
- 1 Tarnished Citadel
- 1 Exotic Orchard
- 1 Cavern of Souls
- 1 Badlands
- 1 Bayou
- 1 Breeding Pool
- 1 Godless Shrine
- 1 Hallowed Fountain
- 1 Plateau
- 1 Savannah
- 1 Scrubland
- 1 Taiga
- 1 Tropical Island
- 1 Tundra
- 1 Underground Sea
- 1 Volcanic Island
- 1 Arid Mesa
- 1 Bloodstained Mire
- 1 Flooded Strand
- 1 Marsh Flats
- 1 Misty Rainforest
- 1 Polluted Delta
- 1 Chrome Mox
- 1 Mox Diamond
- 1 Mana Vault
- 1 Lion's Eye Diamond
- 1 Lotus Petal
- 1 Sol Ring
- 1 Mox Opal
- 1 Mox Amber
- 1 Culling the Weak
- 1 Dark Ritual
- 1 Rite of Flame
- 1 Birds of Paradise
- 1 Deathrite Shaman
- 1 Delighted Halfling
- 1 Ignoble Hierarch
- 1 Bloom Tender
- 1 Arcane Signet
- 1 Ad Nauseam
- 1 Mystic Remora
- 1 Rhystic Study
- 1 Esper Sentinel
- 1 Archivist of Oghma
- 1 Faerie Mastermind
- 1 Borne Upon a Wind
- 1 Wheel of Fortune
- 1 Enduring Curiosity
- 1 Distant Melody
- 1 Mask of Memory
- 1 Samut, Vizier of Naktamun
- 1 Species Specialist
- 1 Force of Will
- 1 Fierce Guardianship
- 1 Flusterstorm
- 1 Pact of Negation
- 1 Mental Misstep
- 1 Swan Song
- 1 Veil of Summer
- 1 Pyroblast
- 1 Red Elemental Blast
- 1 Mindbreak Trap
- 1 An Offer You Can't Refuse
- 1 Vexing Bauble
- 1 Skrelv, Defector Mite
- 1 Legolas's Quick Reflexes
- 1 The One Ring
- 1 Abrupt Decay
- 1 Chain of Vapor
- 1 Deadly Rollick
- 1 Fire Covenant
- 1 Into the Flood Maw
- 1 Swords to Plowshares
- 1 Orcish Bowmasters
- 1 Fighter Class
- 1 Umezawa's Jitte
- 1 Demonic Tutor
- 1 Abdel Adrian, Gorion's Ward
- 1 Alexios, Deimos of Kosmos
- 1 Brago, King Eternal
- 1 Bruenor Battlehammer
- 1 Celestine, the Living Saint
- 1 Dwynen, Gilt-Leaf Daen
- 1 Krenko, Mob Boss
- 1 Maja, Bretagard Protector
- 1 Railway Brawler
- 1 Toph, Hardheaded Teacher
- 1 Zealous Conscripts
- 1 Herald of Secret Streams
- 1 Tezzeret, Cruel Captain
- 1 Auton Soldier
- 1 Angel of Destiny
- 1 Cyclonic Rift
- 1 Chord of Calling

