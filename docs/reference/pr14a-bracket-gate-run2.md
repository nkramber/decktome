# PR-14A bracket gate

Run date: 2026-09-13. Card snapshot: 2026-09-04.

Verdict: FAIL. 14 of 15 decks passed every block check, 10 sat in every band, 14 held no content violation, and the judge agreed with the bracket on 6 of 15 (40 percent, the bar is 80).

## Summary

| Measure | Value |
|---|---|
| Prompts | 15 |
| Decks returned | 15 |
| Decks with no block finding | 14 |
| Decks in every band | 10 |
| Decks with no content violation | 14 |
| Decks the endpoint did not check | 0 |
| Decks that needed a repair turn | 1 |
| Decks judged | 15 |
| Judge agreed with the bracket | 6 |
| Judge errors | 0 |
| Errors | 0 |
| Prompt version | 12 |
| Calls | 31 |
| Cost | $1.3720 |
| Time | 925 seconds |

## Run

- Suite `bracket`, run `pr14a-bracket-gate-run2`, on 2026-09-13, commit `320fcbb`.
- Roles: generate on `gpt-5.6-terra` (openai, effort medium), judge on `claude-opus-5` (anthropic, effort medium), repair on `gpt-5.6-terra` (openai, effort medium).
- Versions: card snapshot 2026-09-04, generate prompt version 12.
- Calls: 31. Cost: $1.3720. Time: 925 seconds.

## Off-band features by key

Each row counts the decks whose feature sat outside its band after the repair turns.

| Feature | Decks off band |
|---|---|
| `mana_turn_four` | 5 |
| `color_sources` | 1 |
| `avg_mana_value` | 1 |

## Decks

### 1. Bracket 1, Gishath, Sun's Avatar, dinosaurs

Pool 267. 99 cards, 0 block findings, repaired false.

Judge: bracket 3, disagrees. This is a focused Dinosaur tribal deck with a strong dual-land mana base (including Plateau and Spire Garden) and efficient ramp like Nature's Lore/Three Visits, putting it clearly above precon power, but its payoffs are big creatures rather than combos. It has no two-card infinite combos, no mass land denial, no extra turns, and essentially no Game Changers, while interaction is limited to a few creature-based removal effects. That combination of upgraded consistency with fair, creature-centric gameplay lands it squarely in Bracket 3.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 38 | 36 to 41 |  |  |
| `tapped_land` | 5 | 0 to 14 |  | Fortified Village, Game Trail, Jetmir's Garden, Jungle Shrine, Path of Ancestry |
| `colorless_land` | 0 | 0 to 4 |  |  |
| `color_sources` | 1.26 | 0.75 or more |  | sources of requirement: W 25.25 of 20, R 26.25 of 19, G 28.5 of 22 |
| `avg_mana_value` | 4.15 | 2.4 to 4.4 |  | over 61 nonland cards |
| `ramp` | 9 | 6 to 12 |  |  |
| `draw` | 10 | 6 to 14 |  |  |
| `removal` | 7 | 4 to 10 |  |  |
| `wipe` | 2 | 0 to 4 |  |  |
| `interaction` | 4 | 0 to 8 |  |  |
| `tutor` | 0 | 0 to 1 |  |  |
| `fast_mana` | 0 | 0 to 1 |  |  |
| `game_changer` | 0 | 0 to 0 |  |  |
| `mana_turn_four` | 4.45 | 3.5 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.76 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | 0.8 | -9 to 1.5 |  | the commander costs 8 and comes down on turn 8.8 on average |

Goldfish over 10000 hands: the commander on turn 8.8, 4.45 mana on turn four, and 0.76 of first hands hold two to four lands.

Content: Spellbook tag E. Game Changers none. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 4.15 over 61 nonland cards

Cards:

- 1 Battlefield Forge
- 1 Bountiful Promenade
- 1 Brushland
- 1 Canopy Vista
- 1 Cinder Glade
- 1 City of Brass
- 1 Clifftop Retreat
- 1 Command Tower
- 1 Exotic Orchard
- 1 Fortified Village
- 1 Game Trail
- 1 Jetmir's Garden
- 1 Jungle Shrine
- 1 Karplusan Forest
- 1 Mana Confluence
- 1 Overgrown Farmland
- 1 Path of Ancestry
- 1 Plateau
- 1 Reflecting Pool
- 1 Rockfall Vale
- 1 Rootbound Crag
- 1 Sacred Foundry
- 1 Secluded Courtyard
- 1 Spectator Seating
- 1 Spire Garden
- 1 Stomping Ground
- 1 Temple Garden
- 5 Forest
- 3 Mountain
- 3 Plains
- 1 Arcane Signet
- 1 Cultivate
- 1 Farseek
- 1 Nature's Lore
- 1 Three Visits
- 1 Thunderherd Migration
- 1 Drover of the Mighty
- 1 Huatli, Poet of Unity // Roar of the Fifth People
- 1 Topiary Stomper
- 1 Beast Whisperer
- 1 Curious Altisaur
- 1 Earthshaker Dreadmaw
- 1 Garruk's Uprising
- 1 Guardian Project
- 1 Harmonize
- 1 Ripjaw Raptor
- 1 Rishkar's Expertise
- 1 Return of the Wildspeaker
- 1 Vanquisher's Banner
- 1 Heroic Intervention
- 1 Lightning Greaves
- 1 Temple Altisaur
- 1 Thrasta, Tempest's Roar
- 1 Apex Altisaur
- 1 Bronzebeak Foragers
- 1 Burning Sun's Avatar
- 1 Itzquinth, Firstborn of Gishath
- 1 Ravenous Sailback
- 1 Thrashing Brontodon
- 1 Trumpeting Carnosaur
- 1 Raging Swordtooth
- 1 Wakening Sun's Avatar
- 1 Amped Raptor
- 1 Armored Kincaller
- 1 Belligerent Yearling
- 1 Deathgorge Scavenger
- 1 Dinosaur Egg
- 1 Dromosaur
- 1 Drowsing Tyrannodon
- 1 Frenzied Raptor
- 1 Huatli's Raptor
- 1 Huatli's Snubhorn
- 1 Hunting Velociraptor
- 1 Kinjalli's Caller
- 1 Marauding Raptor
- 1 Nest Robber
- 1 Otepec Huntmaster
- 1 Orazca Frillback
- 1 Raptor Companion
- 1 Carnage Tyrant
- 1 Etali, Primal Storm
- 1 Ghalta, Primal Hunger
- 1 Ghalta, Stampede Tyrant
- 1 Goring Ceratops
- 1 Palani's Hatcher
- 1 Pantlaza, Sun-Favored
- 1 Quartzwood Crasher
- 1 Regisaur Alpha
- 1 Tyrranax Rex
- 1 Zetalpa, Primal Dawn
- 1 Zilortha, Strength Incarnate

### 2. Bracket 1, Karlov of the Ghost Council, lifegain

Pool 303. 99 cards, 0 block findings, repaired false.

Judge: bracket 3, disagrees. This is a straightforward lifegain-payoff build with a heavy basic-land mana base, no fast mana, and no tutors, but it does play Ocelot Pride, a Game Changer, which alone puts it past Bracket 2's restrictions. Its clock is grindy "gain life, drain out" value rather than any cheap two-card infinite, and interaction is modest spot removal plus two wraths. That lands it comfortably in Upgraded territory, at the low end of Bracket 3.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 38 | 36 to 41 |  |  |
| `tapped_land` | 0 | 0 to 14 |  |  |
| `colorless_land` | 0 | 0 to 4 |  |  |
| `color_sources` | 1.01 | 0.75 or more |  | sources of requirement: W 26.25 of 26, B 26.25 of 26 |
| `avg_mana_value` | 3.2 | 2.4 to 4.4 |  | over 61 nonland cards |
| `ramp` | 9 | 6 to 12 |  |  |
| `draw` | 10 | 6 to 14 |  |  |
| `removal` | 7 | 4 to 10 |  |  |
| `wipe` | 2 | 0 to 4 |  |  |
| `interaction` | 4 | 0 to 8 |  |  |
| `tutor` | 0 | 0 to 1 |  |  |
| `fast_mana` | 0 | 0 to 1 |  |  |
| `game_changer` | 0 | 0 to 0 |  |  |
| `mana_turn_four` | 4.31 | 3.5 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.76 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | 0.01 | -9 to 1.5 |  | the commander costs 2 and comes down on turn 2.01 on average |

Goldfish over 10000 hands: the commander on turn 2.01, 4.31 mana on turn four, and 0.76 of first hands hold two to four lands.

Content: Spellbook tag E. Game Changers none. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 3.20 over 61 nonland cards
- INFO `mana_pass`: the builder moved 1 card of the mana base to bring the deck inside its power level

Cards:

- 14 Plains
- 14 Swamp
- 1 Command Tower
- 1 Godless Shrine
- 1 Caves of Koilos
- 1 Vault of Champions
- 1 Shattered Sanctum
- 1 Isolated Chapel
- 1 City of Brass
- 1 Mana Confluence
- 1 Exotic Orchard
- 1 Reflecting Pool
- 1 Altar of the Pantheon
- 1 Angel of Indemnity
- 1 Beza, the Bounding Spring
- 1 Colossal Plow
- 1 Nuka-Cola Vending Machine
- 1 Orazca Relic
- 1 Phial of Galadriel
- 1 Pristine Talisman
- 1 The Celestus
- 1 Archivist of Oghma
- 1 Convalescent Care
- 1 Cosmos Elixir
- 1 Dawn of Hope
- 1 Enduring Innocence
- 1 Exemplar of Light
- 1 Sigarda's Splendor
- 1 Survival Cache
- 1 The Gaffer
- 1 Well of Lost Dreams
- 1 Alseid of Life's Bounty
- 1 Courageous Resolve
- 1 Faith's Shield
- 1 Restoration Magic
- 1 Ayli, Eternal Pilgrim
- 1 Consuming Corruption
- 1 Gumdrop Poisoner // Tempt with Treats
- 1 Henrika Domnathi // Henrika, Infernal Seer
- 1 Murderous Rider // Swift End
- 1 Nightmare's Thirst
- 1 Vona, Butcher of Magan
- 1 Ajani, Strength of the Pride
- 1 Fumigate
- 1 Ajani's Pridemate
- 1 Angel of Vitality
- 1 Bloodthirsty Aerialist
- 1 Cleric Class
- 1 Cleric of Life's Bond
- 1 Essence Channeler
- 1 Griffin Aerie
- 1 Heliod, Sun-Crowned
- 1 Leyline of Hope
- 1 Marauding Blight-Priest
- 1 Ocelot Pride
- 1 Resplendent Angel
- 1 Righteous Valkyrie
- 1 Serra Ascendant
- 1 Speaker of the Heavens
- 1 Vito, Thorn of the Dusk Rose
- 1 Voice of the Blessed
- 1 Archangel of Thune
- 1 Attended Healer
- 1 Celestine, the Living Saint
- 1 Cliffhaven Vampire
- 1 Defiant Bloodlord
- 1 Divinity of Pride
- 1 Enduring Tenacity
- 1 Lyra Dawnbringer
- 1 Nykthos Paragon
- 1 Rhox Faithmender
- 1 Twinblade Paladin
- 1 Valkyrie Harbinger

### 3. Bracket 1, Adeline, Resplendent Cathar, tokens

Pool 296. 99 cards, 0 block findings, repaired false.

Judge: bracket 3, disagrees. This is a battle-cruiser mono-white token swarm with no fast mana, no tutors, and a basic-heavy mana base, but it runs Ocelot Pride, a Game Changer, which rules out Bracket 2 outright. The deck is clearly upgraded past precon level with strong payoffs like Mondrak, Divine Visitation, Hero of Bladehold and Ojer Taq, plus a decent removal/wrath suite, yet it has no two-card infinite combos and wins through combat over several turns. That places it comfortably in Bracket 3 (one Game Changer, well under the three allowed).

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 38 | 36 to 41 |  |  |
| `tapped_land` | 2 | 0 to 14 |  | Path of Ancestry, Windbrisk Heights |
| `colorless_land` | 1 | 0 to 4 |  | Fountainport |
| `color_sources` | 1.45 | 0.75 or more |  | sources of requirement: W 37.75 of 26 |
| `avg_mana_value` | 3.25 | 2.4 to 4.4 |  | over 61 nonland cards |
| `ramp` | 9 | 6 to 12 |  |  |
| `draw` | 10 | 6 to 14 |  |  |
| `removal` | 7 | 4 to 10 |  |  |
| `wipe` | 2 | 0 to 4 |  |  |
| `interaction` | 4 | 0 to 8 |  |  |
| `tutor` | 1 | 0 to 1 |  | Myr Turbine |
| `fast_mana` | 0 | 0 to 1 |  |  |
| `game_changer` | 0 | 0 to 0 |  |  |
| `mana_turn_four` | 3.79 | 3.5 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.76 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | 0.12 | -9 to 1.5 |  | the commander costs 3 and comes down on turn 3.12 on average |

Goldfish over 10000 hands: the commander on turn 3.12, 3.79 mana on turn four, and 0.76 of first hands hold two to four lands.

Content: Spellbook tag E. Game Changers none. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 3.25 over 61 nonland cards

Cards:

- 28 Plains
- 1 Ancient Den
- 1 Castle Ardenvale
- 1 Eiganjo, Seat of the Empire
- 1 Fountainport
- 1 Kjeldoran Outpost
- 1 Minas Tirith
- 1 Mirrex
- 1 Path of Ancestry
- 1 Secluded Courtyard
- 1 Windbrisk Heights
- 1 Druidic Satchel
- 1 Fishing Gear
- 1 Karn, Living Legacy
- 1 Keeper of the Accord
- 1 Monologue Tax
- 1 Nuka-Cola Vending Machine
- 1 Ojer Taq, Deepest Foundation // Temple of Civilization
- 1 Smuggler's Share
- 1 The Restoration of Eiganjo // Architect of Restoration
- 1 Bennie Bracks, Zoologist
- 1 Bygone Bishop
- 1 Caretaker's Talent
- 1 Chivalric Alliance
- 1 Court of Grace
- 1 Dawn of Hope
- 1 Idol of Oblivion
- 1 Search the Premises
- 1 Staff of the Storyteller
- 1 Wedding Announcement // Wedding Festivity
- 1 Aerial Assault
- 1 Banishing Slash
- 1 Generous Gift
- 1 Kellan's Lightblades
- 1 Skyclave Apparition
- 1 Stroke of Midnight
- 1 The Wandering Emperor
- 1 Hour of Reckoning
- 1 Martial Coup
- 1 Appa, Steadfast Guardian
- 1 Blessed Sanctuary
- 1 Rootborn Defenses
- 1 Spirit Bonds
- 1 Aligned Heart
- 1 Animation Module
- 1 Anointer Priest
- 1 Cathar's Call
- 1 Charismatic Conqueror
- 1 Clarion Spirit
- 1 Divine Visitation
- 1 Felidar Retreat
- 1 Hanweir Militia Captain // Westvale Cult Leader
- 1 Horn of Gondor
- 1 Intangible Virtue
- 1 Luminarch Ascension
- 1 Mavren Fein, Dusk Apostle
- 1 Myr Turbine
- 1 Ocelot Pride
- 1 Oketra's Monument
- 1 Retrofitter Foundry
- 1 Ajani's Chosen
- 1 Archangel Elspeth
- 1 Archon of Sun's Grace
- 1 Attended Healer
- 1 Basri's Lieutenant
- 1 Cemetery Protector
- 1 Defiler of Faith
- 1 Emeria Angel
- 1 Gideon, Ally of Zendikar
- 1 God-Eternal Oketra
- 1 Hero of Bladehold
- 1 Mondrak, Glory Dominus

### 4. Bracket 2, Gishath, Sun's Avatar, dinosaurs

Pool 267. 99 cards, 0 block findings, repaired false.

Judge: bracket 3, disagrees. This is a straightforward Dinosaur tribal beatdown deck with no Game Changers, no mass land denial, no extra turns, and no infinite combos, so nothing pushes it into bracket 4. But the mana base (Taiga, Plateau, shocks, Cavern of Souls), the mana dorks and Sol Ring, and the density of efficient ramp and card draw put it clearly above precon strength, with Gishath threatening to snowball around turn six. That upgraded-but-fair profile lands it in Bracket 3, at the higher end of casual play.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 37 | 35 to 40 |  |  |
| `tapped_land` | 0 | 0 to 12 |  |  |
| `colorless_land` | 0 | 0 to 4 |  |  |
| `color_sources` | 0.96 | 0.8 or more |  | sources of requirement: W 19.25 of 19, R 18.25 of 19, G 25.75 of 26 |
| `avg_mana_value` | 3.71 | 2.3 to 3.9 |  | over 62 nonland cards |
| `ramp` | 10 | 8 to 12 |  |  |
| `draw` | 11 | 8 to 14 |  |  |
| `removal` | 7 | 5 to 10 |  |  |
| `wipe` | 3 | 1 to 5 |  |  |
| `interaction` | 6 | 2 to 10 |  |  |
| `tutor` | 0 | 0 to 2 |  |  |
| `fast_mana` | 1 | 0 to 2 |  | Sol Ring |
| `game_changer` | 0 | 0 to 0 |  |  |
| `mana_turn_four` | 4.79 | 3.8 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.75 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | 0.44 | -9 to 1.2 |  | the commander costs 8 and comes down on turn 8.44 on average |

Goldfish over 10000 hands: the commander on turn 8.44, 4.79 mana on turn four, and 0.75 of first hands hold two to four lands.

Content: Spellbook tag E. Game Changers none. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 3.71 over 62 nonland cards
- INFO `mana_pass`: the builder moved 3 cards of the mana base to bring the deck inside its power level

Cards:

- 13 Forest
- 4 Plains
- 4 Mountain
- 1 Battlefield Forge
- 1 Bountiful Promenade
- 1 Brushland
- 1 Canopy Vista
- 1 Cavern of Souls
- 1 City of Brass
- 1 Clifftop Retreat
- 1 Command Tower
- 1 Exotic Orchard
- 1 Karplusan Forest
- 1 Mana Confluence
- 1 Plateau
- 1 Sacred Foundry
- 1 Stomping Ground
- 1 Temple Garden
- 1 Taiga
- 1 Sol Ring
- 1 Arcane Signet
- 1 Birds of Paradise
- 1 Drover of the Mighty
- 1 Llanowar Elves
- 1 Nature's Lore
- 1 Farseek
- 1 Cultivate
- 1 Rampant Growth
- 1 Thunderherd Migration
- 1 Beast Whisperer
- 1 Cloudpiercer
- 1 Curious Altisaur
- 1 Earthshaker Dreadmaw
- 1 Garruk's Uprising
- 1 Guardian Project
- 1 Harmonize
- 1 Ripjaw Raptor
- 1 Return of the Wildspeaker
- 1 Rishkar's Expertise
- 1 Runic Armasaur
- 1 Akroma's Will
- 1 Boros Charm
- 1 Dawn's Truce
- 1 Heroic Intervention
- 1 Lightning Greaves
- 1 Swiftfoot Boots
- 1 Itzquinth, Firstborn of Gishath
- 1 Needletooth Raptor
- 1 Raging Regisaur
- 1 Ravenous Sailback
- 1 Territorial Allosaurus
- 1 Thrashing Brontodon
- 1 Trumpeting Carnosaur
- 1 Armored Kincaller
- 1 Belligerent Yearling
- 1 Commune with Dinosaurs
- 1 Deathgorge Scavenger
- 1 Dinosaur Egg
- 1 Dinosaur Stampede
- 1 Huatli's Raptor
- 1 Kinjalli's Caller
- 1 Kinjalli's Sunwing
- 1 Marauding Raptor
- 1 Otepec Huntmaster
- 1 Raptor Companion
- 1 Ravenous Daggertooth
- 1 Carnage Tyrant
- 1 Colossal Dreadmaw
- 1 Ghalta, Primal Hunger
- 1 Ghalta, Stampede Tyrant
- 1 Quartzwood Crasher
- 1 Regisaur Alpha
- 1 Sawhorn Nemesis
- 1 Shifting Ceratops
- 1 Snapping Sailback
- 1 Sun-Crested Pterodon
- 1 Tyrranax Rex
- 1 Zetalpa, Primal Dawn
- 1 Blasphemous Act
- 1 Chain Reaction
- 1 Raging Swordtooth

### 5. Bracket 2, Karlov of the Ghost Council, lifegain

Pool 303. 99 cards, 0 block findings, repaired false.

Judge: bracket 3, disagrees. This is a focused lifegain/Karlov aggro-drain build well above precon power, with tuned payoffs (Heliod, Archangel of Thune, Vito, Aetherflux Reservoir) and premium interaction, but no fast mana, no tutors, and a plain mostly-basic mana base. It contains one Game Changer in Solitude and no reliable cheap two-card infinite, so it fits comfortably within Upgraded limits. The clock is a grindy turn six-plus rather than an optimized kill, keeping it out of Bracket 4.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 37 | 35 to 40 |  |  |
| `tapped_land` | 1 | 0 to 12 |  | Scoured Barrens |
| `colorless_land` | 0 | 0 to 4 |  |  |
| `color_sources` | 1.1 | 0.8 or more |  | sources of requirement: W 28.5 of 23, B 28.5 of 26 |
| `avg_mana_value` | 3.11 | 2.3 to 3.9 |  | over 62 nonland cards |
| `ramp` | 10 | 8 to 12 |  |  |
| `draw` | 11 | 8 to 14 |  |  |
| `removal` | 7 | 5 to 10 |  |  |
| `wipe` | 3 | 1 to 5 |  |  |
| `interaction` | 6 | 2 to 10 |  |  |
| `tutor` | 0 | 0 to 2 |  |  |
| `fast_mana` | 1 | 0 to 2 |  | Legion's Landing // Adanto, the First Fort |
| `game_changer` | 0 | 0 to 0 |  |  |
| `mana_turn_four` | 4.43 | 3.8 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.75 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | 0.01 | -9 to 1.2 |  | the commander costs 2 and comes down on turn 2.01 on average |

Goldfish over 10000 hands: the commander on turn 2.01, 4.43 mana on turn four, and 0.75 of first hands hold two to four lands.

Content: Spellbook tag E. Game Changers none. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 3.11 over 62 nonland cards

Cards:

- 13 Plains
- 13 Swamp
- 1 Command Tower
- 1 Godless Shrine
- 1 Caves of Koilos
- 1 City of Brass
- 1 Mana Confluence
- 1 Vault of Champions
- 1 Shattered Sanctum
- 1 Isolated Chapel
- 1 Reflecting Pool
- 1 Exotic Orchard
- 1 Scoured Barrens
- 1 Altar of the Pantheon
- 1 Arguel's Blood Fast // Temple of Aclazotz
- 1 Legion's Landing // Adanto, the First Fort
- 1 Nuka-Cola Vending Machine
- 1 Pristine Talisman
- 1 Orazca Relic
- 1 Phial of Galadriel
- 1 The Celestus
- 1 Potioner's Trove
- 1 Hot Dog Cart
- 1 Archivist of Oghma
- 1 Dawn of Hope
- 1 Cosmos Elixir
- 1 Enduring Innocence
- 1 Exemplar of Light
- 1 Mangara, the Diplomat
- 1 Markov Purifier
- 1 Sigarda's Splendor
- 1 Well of Lost Dreams
- 1 Survival Cache
- 1 Vampiric Rites
- 1 Alseid of Life's Bounty
- 1 Courageous Resolve
- 1 Bofur, Reliable Guardian // Concerted Care
- 1 Caduceus, Staff of Hermes
- 1 Faith's Shield
- 1 Distinguished Conjurer
- 1 Aetherflux Reservoir
- 1 Ayli, Eternal Pilgrim
- 1 Murderous Rider // Swift End
- 1 Nightmare's Thirst
- 1 Solitude
- 1 Umezawa's Jitte
- 1 Vona, Butcher of Magan
- 1 Fumigate
- 1 Kaya's Wrath
- 1 The Battle of Bywater
- 1 Ajani's Pridemate
- 1 Angel of Vitality
- 1 Bloodthirsty Aerialist
- 1 Cleric Class
- 1 Cleric of Life's Bond
- 1 Heliod, Sun-Crowned
- 1 Leyline of Hope
- 1 Resplendent Angel
- 1 Righteous Valkyrie
- 1 Serra Ascendant
- 1 Speaker of the Heavens
- 1 Vito, Thorn of the Dusk Rose
- 1 Voice of the Blessed
- 1 Archangel of Thune
- 1 Attended Healer
- 1 Celestine, the Living Saint
- 1 Cliffhaven Vampire
- 1 Enduring Tenacity
- 1 Gideon's Company
- 1 Lyra Dawnbringer
- 1 Nykthos Paragon
- 1 Regal Bloodlord
- 1 Rhox Faithmender
- 1 Valkyrie Harbinger
- 1 Twinblade Paladin

### 6. Bracket 2, Adeline, Resplendent Cathar, tokens

Pool 296. 99 cards, 0 block findings, repaired false.

Judge: bracket 3, disagrees. This is a well-tuned mono-white go-wide tokens deck with a clean curve, strong payoffs (Hero of Bladehold, Divine Visitation, Defiler of Faith) and plenty of card advantage, but no fast mana, no tutors, and no two-card infinite combos. Ocelot Pride is a Game Changer, which pushes it out of Bracket 2, though a single Game Changer sits comfortably under the Bracket 3 cap of three. The clock is fast for a fair deck but relies on creature-based attrition rather than turn-four kills, so it lands squarely as Upgraded.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 37 | 35 to 40 |  |  |
| `tapped_land` | 1 | 0 to 12 |  | Windbrisk Heights |
| `colorless_land` | 1 | 0 to 4 |  | Fountainport |
| `color_sources` | 1.41 | 0.8 or more |  | sources of requirement: W 36.75 of 26 |
| `avg_mana_value` | 3.16 | 2.3 to 3.9 |  | over 62 nonland cards |
| `ramp` | 10 | 8 to 12 |  |  |
| `draw` | 11 | 8 to 14 |  |  |
| `removal` | 7 | 5 to 10 |  |  |
| `wipe` | 3 | 1 to 5 |  |  |
| `interaction` | 6 | 2 to 10 |  |  |
| `tutor` | 0 | 0 to 2 |  |  |
| `fast_mana` | 1 | 0 to 2 |  | Legion's Landing // Adanto, the First Fort |
| `game_changer` | 0 | 0 to 0 |  |  |
| `mana_turn_four` | 4.03 | 3.8 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.75 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | 0.01 | -9 to 1.2 |  | the commander costs 3 and comes down on turn 3.01 on average |

Goldfish over 10000 hands: the commander on turn 3.01, 4.03 mana on turn four, and 0.75 of first hands hold two to four lands.

Content: Spellbook tag E. Game Changers none. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 3.16 over 62 nonland cards

Cards:

- 25 Plains
- 1 Ancient Den
- 1 Castle Ardenvale
- 1 Cavern of Souls
- 1 Command Tower
- 1 Eiganjo, Seat of the Empire
- 1 Fountainport
- 1 Kjeldoran Outpost
- 1 Minas Tirith
- 1 Nykthos, Shrine to Nyx
- 1 Plaza of Heroes
- 1 Secluded Courtyard
- 1 Windbrisk Heights
- 1 Bucknard's Everfull Purse
- 1 Collector's Vault
- 1 Druidic Satchel
- 1 Goldvein Pick
- 1 Karn, Living Legacy
- 1 Keeper of the Accord
- 1 Legion's Landing // Adanto, the First Fort
- 1 Monologue Tax
- 1 Smuggler's Share
- 1 Sword of Wealth and Power
- 1 Bennie Bracks, Zoologist
- 1 Bygone Bishop
- 1 Caretaker's Talent
- 1 Chivalric Alliance
- 1 Court of Grace
- 1 Dawn of Hope
- 1 Faramir, Field Commander
- 1 Idol of Oblivion
- 1 Staff of the Storyteller
- 1 Thorough Investigation
- 1 Wedding Announcement // Wedding Festivity
- 1 Ainok Strike Leader
- 1 Appa, Steadfast Guardian
- 1 Blessed Sanctuary
- 1 Elspeth, Knight-Errant
- 1 Rootborn Defenses
- 1 Spirit Bonds
- 1 Aerial Assault
- 1 Banishing Slash
- 1 Citizen's Crowbar
- 1 Generous Gift
- 1 Kellan's Lightblades
- 1 Skyclave Apparition
- 1 Stroke of Midnight
- 1 Ceaseless Conflict
- 1 Hour of Reckoning
- 1 Martial Coup
- 1 Charismatic Conqueror
- 1 Clarion Spirit
- 1 Divine Visitation
- 1 Felidar Retreat
- 1 Hanweir Militia Captain // Westvale Cult Leader
- 1 Horn of Gondor
- 1 Intangible Virtue
- 1 Myrsmith
- 1 Ocelot Pride
- 1 Oketra's Monument
- 1 Retrofitter Foundry
- 1 Rosie Cotton of South Lane
- 1 Siege Veteran
- 1 Archangel Elspeth
- 1 Attended Healer
- 1 Basri's Lieutenant
- 1 Cemetery Protector
- 1 Defiler of Faith
- 1 Emeria Angel
- 1 God-Eternal Oketra
- 1 Hero of Bladehold
- 1 Myrel, Shield of Argive
- 1 Oketra the True
- 1 Sanctuary Warden
- 1 Warren Warleader

### 7. Bracket 3, Karlov of the Ghost Council, lifegain

Pool 303. 99 cards, 0 block findings, repaired false.

Judge: bracket 3, agrees. This is a linear lifegain deck with a couple of high-power inclusions anging up to Game Changer territory (Tymna the Weaver, plus premium cards like Solitude, Vein Ripper and Aetherflux Reservoir), which pushes it past precon strength. Its mana base is 28 basics with no fast mana and it has almost no tutoring, so it can't reliably assemble anything early, and its payoff engines (Archangel of Thune, Vito, Heliod) need creatures and turns to convert into a kill. That lands it as an upgraded but fair deck ending games around turn six to eight, well under three Game Changers.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 36 | 34 to 38 |  |  |
| `tapped_land` | 2 | 0 to 9 |  | Scoured Barrens, Temple of Silence |
| `colorless_land` | 0 | 0 to 4 |  |  |
| `color_sources` | 0.96 | 0.85 or more |  | sources of requirement: W 23 of 23, B 27 of 28 |
| `avg_mana_value` | 3.21 | 2 to 3.5 |  | over 63 nonland cards |
| `ramp` | 10 | 8 to 13 |  |  |
| `draw` | 11 | 8 to 14 |  |  |
| `removal` | 9 | 6 to 12 |  |  |
| `wipe` | 3 | 2 to 5 |  |  |
| `interaction` | 8 | 4 to 12 |  |  |
| `tutor` | 0 | 0 to 5 |  |  |
| `fast_mana` | 1 | 0 to 3 |  | Legion's Landing // Adanto, the First Fort |
| `game_changer` | 0 | 0 to 3 |  |  |
| `mana_turn_four` | 4.34 | 4.2 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.74 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | 0.02 | -9 to 0.8 |  | the commander costs 2 and comes down on turn 2.02 on average |

Goldfish over 10000 hands: the commander on turn 2.02, 4.34 mana on turn four, and 0.74 of first hands hold two to four lands.

Content: Spellbook tag E. Game Changers none. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 3.21 over 63 nonland cards
- INFO `mana_pass`: the builder moved 2 cards of the mana base to bring the deck inside its power level

Cards:

- 12 Plains
- 16 Swamp
- 1 Caves of Koilos
- 1 Command Tower
- 1 Godless Shrine
- 1 Isolated Chapel
- 1 Scoured Barrens
- 1 Shattered Sanctum
- 1 Temple of Silence
- 1 Vault of Champions
- 1 Archivist of Oghma
- 1 Dawn of Hope
- 1 Enduring Innocence
- 1 Mangara, the Diplomat
- 1 Markov Purifier
- 1 Restless Bloodseeker // Bloodsoaked Reveler
- 1 Sigarda's Splendor
- 1 The Gaffer
- 1 Tymna the Weaver
- 1 Vampiric Rites
- 1 Well of Lost Dreams
- 1 Alseid of Life's Bounty
- 1 Bofur, Reliable Guardian // Concerted Care
- 1 Caduceus, Staff of Hermes
- 1 Courageous Resolve
- 1 Faith's Shield
- 1 Metropolis Reformer
- 1 Restoration Magic
- 1 Werefox Bodyguard
- 1 Altar of the Pantheon
- 1 Arguel's Blood Fast // Temple of Aclazotz
- 1 Colossal Plow
- 1 Hot Dog Cart
- 1 Legion's Landing // Adanto, the First Fort
- 1 Life Insurance
- 1 Nuka-Cola Vending Machine
- 1 Orazca Relic
- 1 Phial of Galadriel
- 1 Pristine Talisman
- 1 Aetherflux Reservoir
- 1 Ayli, Eternal Pilgrim
- 1 Consuming Corruption
- 1 Gumdrop Poisoner // Tempt with Treats
- 1 Murderous Rider // Swift End
- 1 Nightmare's Thirst
- 1 Solitude
- 1 Vein Ripper
- 1 Vona, Butcher of Magan
- 1 Ajani's Pridemate
- 1 Angel of Vitality
- 1 Bloodthirsty Aerialist
- 1 Cleric Class
- 1 Heliod, Sun-Crowned
- 1 Indulging Patrician
- 1 Resplendent Angel
- 1 Righteous Valkyrie
- 1 Vito, Thorn of the Dusk Rose
- 1 Voice of the Blessed
- 1 Archangel of Thune
- 1 Attended Healer
- 1 Blood Baron of Vizkopa
- 1 Celestine, the Living Saint
- 1 Cliffhaven Vampire
- 1 Defiant Bloodlord
- 1 Divinity of Pride
- 1 Enduring Tenacity
- 1 Nykthos Paragon
- 1 Rhox Faithmender
- 1 Twinblade Paladin
- 1 Valkyrie Harbinger
- 1 Fumigate
- 1 Kaya's Wrath
- 1 The Meathook Massacre

### 8. Bracket 3, Denethor, Ruling Steward, aristocrats sacrifice

Pool 295. 99 cards, 0 block findings, repaired false.

Judge: bracket 3, agrees. This is a well-tuned aristocrats deck with two Game Changers (Yawgmoth, Thran Physician and Grave Pact), a good but not broken dual/shock/pain mana base, and no real fast mana beyond Culling the Weak and the sac-altars. It has plenty of combo potential (Chthonian Nightmare/Pitiless Plunderer loops, Blasting Station and drain payoffs), but these are engine-based and generally assemble past turn six rather than being cheap two-card kills. No mass land denial, no extra turns, and interaction is mostly creature removal and wipes — squarely Upgraded rather than Optimized.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 38 | 34 to 38 |  |  |
| `tapped_land` | 0 | 0 to 9 |  |  |
| `colorless_land` | 2 | 0 to 4 |  | High Market, Westvale Abbey // Ormendahl, Profane Prince |
| `color_sources` | 0.99 | 0.85 or more |  | sources of requirement: W 26 of 26, B 22.75 of 23 |
| `avg_mana_value` | 3.02 | 2 to 3.5 |  | over 61 nonland cards |
| `ramp` | 11 | 8 to 13 |  |  |
| `draw` | 11 | 8 to 14 |  |  |
| `removal` | 9 | 6 to 12 |  |  |
| `wipe` | 2 | 2 to 5 |  |  |
| `interaction` | 8 | 4 to 12 |  |  |
| `tutor` | 0 | 0 to 5 |  |  |
| `fast_mana` | 1 | 0 to 3 |  | Culling the Weak |
| `game_changer` | 0 | 0 to 3 |  |  |
| `mana_turn_four` | 4.28 | 4.2 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.76 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.08 | -9 to 0.8 |  | the commander costs 3 and comes down on turn 2.92 on average |

Goldfish over 10000 hands: the commander on turn 2.92, 4.28 mana on turn four, and 0.76 of first hands hold two to four lands.

Content: Spellbook tag E. Game Changers none. Mass land denial none. Extra turns none. Combos 1.
- Chthonian Nightmare + Pitiless Plunderer + Mondrak, Glory Dominus (speed 5)

Findings:

- INFO `curve_summary`: average mana value 3.02 over 61 nonland cards
- INFO `mana_pass`: the builder moved 7 cards of the mana base to bring the deck inside its power level

Cards:

- 9 Swamp
- 14 Plains
- 1 Godless Shrine
- 1 Scrubland
- 1 Caves of Koilos
- 1 Command Tower
- 1 City of Brass
- 1 Exotic Orchard
- 1 Fetid Heath
- 1 Mana Confluence
- 1 Vault of Champions
- 1 Shattered Sanctum
- 1 Isolated Chapel
- 1 Reflecting Pool
- 1 High Market
- 1 Phyrexian Tower
- 1 Westvale Abbey // Ormendahl, Profane Prince
- 1 Corrupted Conviction
- 1 Village Rites
- 1 Disciple of Bolas
- 1 Smothering Abomination
- 1 Shadowheart, Dark Justiciar
- 1 Relic Vial
- 1 Grave Venerations
- 1 Jerren, Corrupted Bishop // Ormendahl, the Corrupter
- 1 Lord Skitter's Butcher
- 1 Skyclave Shadowcat
- 1 Vampire Gourmand
- 1 Cartel Aristocrat
- 1 Cultist of the Absolute
- 1 Dark Privilege
- 1 Fanatical Devotion
- 1 Flare of Fortitude
- 1 Gift of Doom
- 1 Kinzu of the Bleak Coven
- 1 Nightmare Shepherd
- 1 Ashnod's Altar
- 1 Phyrexian Altar
- 1 Pitiless Plunderer
- 1 Pawn of Ulamog
- 1 Priest of Forgotten Gods
- 1 Skullport Merchant
- 1 Warren Soultrader
- 1 Culling the Weak
- 1 Deadly Dispute
- 1 Tarrian's Journal // The Tomb of Aclazotz
- 1 Attrition
- 1 Ayli, Eternal Pilgrim
- 1 Bone Shards
- 1 Eaten Alive
- 1 Dictate of Erebos
- 1 Grave Pact
- 1 Blasting Station
- 1 Chittering Witch
- 1 Yawgmoth, Thran Physician
- 1 Blood Artist
- 1 Carrion Feeder
- 1 Viscera Seer
- 1 Zulaport Cutthroat
- 1 Bastion of Remembrance
- 1 Elas il-Kor, Sadistic Pilgrim
- 1 Cruel Celebrant
- 1 Woe Strider
- 1 Chthonian Nightmare
- 1 Victimize
- 1 Elenda, the Dusk Rose
- 1 Ghoulcaller Gisa
- 1 Liesa, Forgotten Archangel
- 1 Mondrak, Glory Dominus
- 1 Ratadrabik of Urborg
- 1 Sadistic Hypnotist
- 1 Syr Konrad, the Grim
- 1 Vindictive Vampire
- 1 Falkenrath Noble
- 1 Corpse Harvester
- 1 Toxic Deluge
- 1 The Meathook Massacre
- 1 Master of Dark Rites

### 9. Bracket 3, Zada, Hedron Grinder, goblins

Pool 294. 99 cards, 0 block findings, repaired false.

Judge: bracket 3, agrees. This is a goblin aristocrats/tribal beatdown list with exactly three Game Changers (Jeska's Will, The One Ring, Kiki-Jiki, Mirror Breaker), modest fast mana (Sol Ring, rituals, Skirk Prospector) and no reliable cheap two-card infinite combo. The mana base is a mono-red pile of basics with a few utility lands, and interaction is mostly red blasts and damage-based removal rather than efficient stack interaction. That's clearly a souped-up-beyond-precon deck that can close around turn six or seven with Muxus/Krenko/Bushwhacker swarms, fitting Upgraded.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 36 | 34 to 38 |  |  |
| `tapped_land` | 3 | 0 to 9 |  | Shatterskull Smashing // Shatterskull, the Hammer Pass, Spinerock Knoll, Valakut Awakening // Valakut Stoneforge |
| `colorless_land` | 0 | 0 to 4 |  |  |
| `color_sources` | 1.7 | 0.85 or more |  | sources of requirement: R 37.5 of 22 |
| `avg_mana_value` | 2.68 | 2 to 3.5 |  | over 63 nonland cards |
| `ramp` | 10 | 8 to 13 |  |  |
| `draw` | 11 | 8 to 14 |  |  |
| `removal` | 9 | 6 to 12 |  |  |
| `wipe` | 3 | 2 to 5 |  |  |
| `interaction` | 8 | 4 to 12 |  |  |
| `tutor` | 0 | 0 to 5 |  |  |
| `fast_mana` | 2 | 0 to 3 |  | Brightstone Ritual, Sol Ring |
| `game_changer` | 2 | 0 to 3 |  | Jeska's Will, The One Ring |
| `mana_turn_four` | 4.69 | 4.2 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.74 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.48 | -9 to 0.8 |  | the commander costs 4 and comes down on turn 3.52 on average |

Goldfish over 10000 hands: the commander on turn 3.52, 4.69 mana on turn four, and 0.74 of first hands hold two to four lands.

Content: Spellbook tag S. Game Changers Jeska's Will, The One Ring. Mass land denial none. Extra turns none. Combos 3.
- Krenko, Mob Boss + Skirk Prospector + Goblin Warchief (speed 5)
- Krenko, Mob Boss + Skirk Prospector + Goblin Chieftain (speed 5)
- Kiki-Jiki, Mirror Breaker + Akki Battle Squad (speed 5, two cards)

Findings:

- INFO `curve_summary`: average mana value 2.68 over 63 nonland cards
- WARN `two_card_combo`: bracket 3 allows no two-card infinite combo that needs four mana or less, and the deck holds Kiki-Jiki, Mirror Breaker + Akki Battle Squad (speed 5, near two-card)
- INFO `basics_added`: the list was 1 card short, so the builder added 1 basic land

Cards:

- 27 Mountain
- 1 Arena of Glory
- 1 Castle Embereth
- 1 Cavern of Souls
- 1 Command Tower
- 1 Great Furnace
- 1 Sokenzan, Crucible of Defiance
- 1 Spinerock Knoll
- 1 Shatterskull Smashing // Shatterskull, the Hammer Pass
- 1 Valakut Awakening // Valakut Stoneforge
- 1 Sol Ring
- 1 Arcane Signet
- 1 Fellwar Stone
- 1 Mind Stone
- 1 Wayfarer's Bauble
- 1 Skirk Prospector
- 1 Brightstone Ritual
- 1 Wily Goblin
- 1 Impulsive Pilferer
- 1 Jeska's Will
- 1 Skullclamp
- 1 Faithless Looting
- 1 Demand Answers
- 1 Thrill of Possibility
- 1 Cathartic Reunion
- 1 Mask of Memory
- 1 Idol of Oblivion
- 1 Vanquisher's Banner
- 1 Reckless Lackey
- 1 Fissure Wizard
- 1 Goblin Picker
- 1 Lightning Greaves
- 1 Swiftfoot Boots
- 1 Pyroblast
- 1 Red Elemental Blast
- 1 Tibalt's Trickery
- 1 Goblin Chirurgeon
- 1 Slobad, Goblin Tinkerer
- 1 The One Ring
- 1 Goblin Grenade
- 1 Goblin War Strike
- 1 Gempalm Incinerator
- 1 Goblin Cratermaker
- 1 Goblin Sharpshooter
- 1 Pashalik Mons
- 1 Siege-Gang Commander
- 1 Broadside Bombardiers
- 1 Sparksmith
- 1 Blasphemous Act
- 1 Vandalblast
- 1 Goblin Chainwhirler
- 1 Goblin Chieftain
- 1 Goblin Warchief
- 1 Rundvelt Hordemaster
- 1 Goblin Rabblemaster
- 1 Mogg War Marshal
- 1 Krenko, Tin Street Kingpin
- 1 Goblin Instigator
- 1 Skirk Drill Sergeant
- 1 Goblin Piledriver
- 1 Goblin Bushwhacker
- 1 Krenko, Mob Boss
- 1 Muxus, Goblin Grandee
- 1 Kiki-Jiki, Mirror Breaker
- 1 Beetleback Chief
- 1 Battle Squadron
- 1 Goblin Marshal
- 1 Reckless One
- 1 Goblin Ringleader
- 1 Goblin Dark-Dwellers
- 1 Akki Battle Squad
- 1 Goblin Assault Team
- 1 Chasm Guide

### 10. Bracket 4, Urza, Lord High Artificer, artifacts

Pool 299. 99 cards, 0 block findings, repaired false.

Judge: bracket 4, agrees. Urza himself plus Mishra's Workshop are Game Changers, and the deck backs them with genuine fast mana (Mox Opal, Metalworker, Vedalken Engineer, Chief Engineer) that can deploy huge artifact turns very early. It also packs combo-capable pieces — Power Artifact, Clock of Omens/Unwinding Clock with untappers, Blasting Station with Arcbound recursion, and Aetherflux Reservoir as a kill — plus counterspell protection, so it can close games well before turn six. The 28-Island mana base and lack of tutors keep it out of cEDH consistency, but the power level and speed sit squarely in Optimized.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 36 | 31 to 36 |  |  |
| `tapped_land` | 0 | 0 to 5 |  |  |
| `colorless_land` | 5 | 0 to 5 |  | Academy Ruins, Inventors' Fair, Mishra's Workshop, Urza's Saga, Urza's Workshop |
| `color_sources` | 1.31 | 0.9 or more |  | sources of requirement: U 34 of 26 |
| `avg_mana_value` | 2.81 | 1.6 to 3 |  | over 63 nonland cards |
| `ramp` | 16 | 10 to 16 |  |  |
| `draw` | 9 | 8 to 16 |  |  |
| `removal` | 10 | 6 to 14 |  |  |
| `wipe` | 2 | 0 to 5 |  |  |
| `interaction` | 10 | 6 to 16 |  |  |
| `tutor` | 2 | none |  | Arcum Dagsson, Treasure Chest |
| `fast_mana` | 2 | none |  | Moonsnare Prototype, Mox Opal |
| `game_changer` | 1 | none |  | Mishra's Workshop |
| `mana_turn_four` | 4.72 | 4.8 or more | yes | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.74 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.17 | -9 to 0.5 |  | the commander costs 4 and comes down on turn 3.83 on average |

Goldfish over 10000 hands: the commander on turn 3.83, 4.72 mana on turn four, and 0.74 of first hands hold two to four lands.

Content: Spellbook tag P. Game Changers Mishra's Workshop. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 2.81 over 63 nonland cards
- WARN `profile_off_band`: the mana available on turn four is 4.72, and bracket 4 wants 4.8 or more (mean over 10000 hands)
- INFO `mana_pass`: the builder moved 10 cards of the mana base to bring the deck inside its power level

Cards:

- 1 Academy Ruins
- 1 Inventors' Fair
- 28 Island
- 1 Mishra's Workshop
- 1 Mystic Sanctuary
- 1 Otawara, Soaring City
- 1 Seat of the Synod
- 1 Urza's Saga
- 1 Urza's Workshop
- 1 Chief Engineer
- 1 Grand Architect
- 1 Inspiring Statuary
- 1 Metalworker
- 1 Moonsnare Prototype
- 1 Mox Opal
- 1 Powerstone Shard
- 1 Treasure Chest
- 1 Vedalken Engineer
- 1 Braided Net // Braided Quipu
- 1 Era of Innovation
- 1 Esoteric Duplicator
- 1 Forensic Gadgeteer
- 1 Riddlesmith
- 1 Sai, Master Thopterist
- 1 Thirst for Knowledge
- 1 Trading Post
- 1 Vedalken Archmage
- 1 Curator's Ward
- 1 Disruption Protocol
- 1 Escape Protocol
- 1 Etched Champion
- 1 Ghostly Flicker
- 1 Ice Out
- 1 Metallic Rebuke
- 1 Padeem, Consul of Innovation
- 1 Stoic Rebuttal
- 1 Welding Jar
- 1 Aether Spellbomb
- 1 Aetherflux Reservoir
- 1 Arcum Dagsson
- 1 Blasting Station
- 1 Contagion Clasp
- 1 Cyber Conversion
- 1 Ravenform
- 1 Resculpt
- 1 Shape Anew
- 1 Unable to Scream
- 1 Engineered Explosives
- 1 Hurkyl's Recall
- 1 Arcbound Crusher
- 1 Arcbound Reclaimer
- 1 Frogmite
- 1 Jhoira's Familiar
- 1 Karn, Scion of Urza
- 1 Lodestone Golem
- 1 Lodestone Myr
- 1 Master Transmuter
- 1 Phyrexian Metamorph
- 1 Traxos, Scourge of Kroog
- 1 Clock of Omens
- 1 Manifold Key
- 1 Mystic Forge
- 1 Power Artifact
- 1 Unwinding Clock
- 1 Voltaic Key
- 1 Cargo Ship
- 1 Kitesail Larcenist
- 1 Luck Bobblehead
- 1 Network Terminal
- 1 Solar Array
- 1 Splitting the Powerstone
- 1 Strixhaven Stadium

### 11. Bracket 4, Korvold, Fae-Cursed King, treasure sacrifice

Pool 300. 99 cards, 0 block findings, repaired false.

Judge: bracket 3, disagrees. This is a solid aristocrats/Korvold sacrifice build with two Game Changers (Yawgmoth, Thran Physician and Grist, the Hunger Tide), a good but not fast mana base (no Sol Ring-tier fast mana, mostly duals and basics), and value-driven ramp via Treasure tokens rather than explosive acceleration. It has strong synergy engines and sac outlets (Blasting Station, Altar of Dementia, Goblin Bombardment) that can assemble grindy loops, but no cheap, reliable two-card infinite that goes off by turn six. That places it comfortably above precon power yet short of optimized: Bracket 3.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 36 | 31 to 36 |  |  |
| `tapped_land` | 0 | 0 to 5 |  |  |
| `colorless_land` | 0 | 0 to 5 |  |  |
| `color_sources` | 1.08 | 0.9 or more |  | sources of requirement: B 23.5 of 19, R 20.5 of 19, G 20.5 of 19 |
| `avg_mana_value` | 2.68 | 1.6 to 3 |  | over 63 nonland cards |
| `ramp` | 15 | 10 to 16 |  |  |
| `draw` | 12 | 8 to 16 |  |  |
| `removal` | 10 | 6 to 14 |  |  |
| `wipe` | 2 | 0 to 5 |  |  |
| `interaction` | 10 | 6 to 16 |  |  |
| `tutor` | 4 | none |  | Diabolic Intent, Dimir House Guard, Magda, Brazen Outlaw, Savage Order |
| `fast_mana` | 0 | none |  |  |
| `game_changer` | 0 | none |  |  |
| `mana_turn_four` | 4.6 | 4.8 or more | yes | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.74 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.33 | -9 to 0.5 |  | the commander costs 5 and comes down on turn 4.67 on average |

Goldfish over 10000 hands: the commander on turn 4.67, 4.6 mana on turn four, and 0.74 of first hands hold two to four lands.

Content: Spellbook tag E. Game Changers none. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 2.68 over 63 nonland cards
- WARN `profile_off_band`: the mana available on turn four is 4.6, and bracket 4 wants 4.8 or more (mean over 10000 hands)
- INFO `mana_pass`: the builder moved 12 cards of the mana base to bring the deck inside its power level

Cards:

- 1 Forsworn Paladin
- 1 Captain Lannery Storm
- 1 Professional Face-Breaker
- 1 Skullport Merchant
- 1 Warren Soultrader
- 1 Magda, Brazen Outlaw
- 1 Magda, the Hoardmaster
- 1 Treasure Map // Treasure Cove
- 1 Dina, Essence Brewer
- 1 Gas Guzzler
- 1 Relic Vial
- 1 Sanguine Spy
- 1 Smothering Abomination
- 1 Shadowheart, Dark Justiciar
- 1 Vampire Gourmand
- 1 Skyclave Shadowcat
- 1 Stormclaw Rager
- 1 Swarm Culler
- 1 Lord Skitter's Butcher
- 1 Disciple of Bolas
- 1 Constant Mists
- 1 Cultist of the Absolute
- 1 Dark Privilege
- 1 Goblin Chirurgeon
- 1 Sylvan Safekeeper
- 1 My Precious // Allure of Power
- 1 Campsite Cuisine
- 1 Altar of the Wretched // Wretched Bonemass
- 1 Savage Order
- 1 Slobad, Goblin Tinkerer
- 1 Attrition
- 1 Bone Shards
- 1 Eaten Alive
- 1 Goblin Bombardment
- 1 Blasting Station
- 1 Broadside Bombardiers
- 1 Fling
- 1 Yawgmoth, Thran Physician
- 1 Grist, the Hunger Tide
- 1 Sawblade Skinripper
- 1 Academy Manufactor
- 1 Xorn
- 1 Witch's Oven
- 1 Viscera Seer
- 1 Diabolic Intent
- 1 Altar of Dementia
- 1 Dimir House Guard
- 1 Meren of Clan Nel Toth
- 1 Vito's Inquisitor
- 1 Vindictive Vampire
- 1 Redcap Gutter-Dweller
- 1 Novice Dissector
- 1 Thallid Omnivore
- 1 The Meathook Massacre
- 1 Spontaneous Combustion
- 1 Blood Crypt
- 1 Cinder Glade
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
- 1 Smoldering Marsh
- 1 Stomping Ground
- 1 Sulfurous Springs
- 1 Undergrowth Stadium
- 1 Woodland Cemetery
- 6 Forest
- 5 Mountain
- 7 Swamp
- 1 Luck Bobblehead
- 1 Master of Dark Rites
- 1 Illuminor Szeras
- 1 Mastermind Plum
- 1 Wight of the Reliquary
- 1 Ruthless Knave
- 1 Cavern of Souls
- 1 Swashbuckler Extraordinaire
- 1 The Misty Mountains Cold

### 12. Bracket 4, Prosper, Tome-Bound, treasure

Pool 188. 99 cards, 0 block findings, repaired false.

Judge: bracket 4, agrees. Four Game Changers (Ragavan, Orcish Bowmasters, The One Ring, Black Market Connections) already push it past Bracket 3's cap of three, and the deck backs them with an expensive dual-heavy mana base, efficient one-mana removal, and free interaction like Deadly Rollick and Snuff Out. The treasure engine (Storm-Kiln Artist, Pitiless Plunderer, Academy Manufactor, Xorn) plus Prosper generates explosive value and easy loop potential well before turn six. That's an optimized build, but without a focused cEDH combo-and-stax plan, so Bracket 4.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 36 | 31 to 36 |  |  |
| `tapped_land` | 0 | 0 to 5 |  |  |
| `colorless_land` | 0 | 0 to 5 |  |  |
| `color_sources` | 1.45 | 0.9 or more |  | sources of requirement: B 27.5 of 19, R 28.5 of 19 |
| `avg_mana_value` | 2.4 | 1.6 to 3 |  | over 63 nonland cards |
| `ramp` | 16 | 10 to 16 |  |  |
| `draw` | 10 | 8 to 16 |  |  |
| `removal` | 11 | 6 to 14 |  |  |
| `wipe` | 2 | 0 to 5 |  |  |
| `interaction` | 12 | 6 to 16 |  |  |
| `tutor` | 1 | none |  | Magda, Brazen Outlaw |
| `fast_mana` | 0 | none |  |  |
| `game_changer` | 2 | none |  | Orcish Bowmasters, The One Ring |
| `mana_turn_four` | 4.42 | 4.8 or more | yes | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.74 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.22 | -9 to 0.5 |  | the commander costs 4 and comes down on turn 3.78 on average |

Goldfish over 10000 hands: the commander on turn 3.78, 4.42 mana on turn four, and 0.74 of first hands hold two to four lands.

Content: Spellbook tag P. Game Changers Orcish Bowmasters, The One Ring. Mass land denial none. Extra turns none. Combos 1.
- Prosper, Tome-Bound + Professional Face-Breaker + Sensei's Divining Top + Xorn (speed 5)

Findings:

- INFO `curve_summary`: average mana value 2.40 over 63 nonland cards
- WARN `profile_off_band`: the mana available on turn four is 4.42, and bracket 4 wants 4.8 or more (mean over 10000 hands)
- INFO `mana_pass`: the builder moved 8 cards of the mana base to bring the deck inside its power level
- INFO `basics_added`: the list was 1 card short, so the builder added 1 basic land

Cards:

- 1 Abrade
- 1 Academy Manufactor
- 1 Badlands
- 1 Bedevil
- 1 Black Market Connections
- 1 Blazemire Verge
- 1 Blood Crypt
- 1 Chaos Warp
- 1 City of Brass
- 1 Command Tower
- 1 Contract Hero
- 1 Corrupted Conviction
- 1 Crossover Collaboration
- 1 Deadly Rollick
- 1 Demand Answers
- 1 Dragonskull Summit
- 1 Exotic Orchard
- 1 Evereth, Viceroy of Plunder
- 1 Faithless Looting
- 1 Feed the Swarm
- 1 Flamekin Gildweaver
- 1 Forsworn Paladin
- 1 Gemcutter Buccaneer
- 1 Gilded Assault Cart
- 1 Go for the Throat
- 1 Graven Cairns
- 1 Grim Hireling
- 1 Haunted Ridge
- 1 Hexing Squelcher
- 1 Idol of Oblivion
- 1 Infernal Grasp
- 1 Jaded Sell-Sword
- 1 Kalain, Reclusive Painter
- 1 Kaya's Ghostform
- 1 Knuckles the Echidna
- 1 Lightning Greaves
- 1 Luck Bobblehead
- 1 Luxury Suite
- 1 Magda, Brazen Outlaw
- 1 Mana Confluence
- 1 Meticulous Artisan
- 1 Mithril Coat
- 10 Mountain
- 1 Namazu Trader
- 1 Night's Whisper
- 1 Not Dead After All
- 1 Orcish Bowmasters
- 1 Phyrexian Arena
- 1 Pitiless Plunderer
- 1 Professional Face-Breaker
- 1 Pyroblast
- 1 Ragavan, Nimble Pilferer
- 1 Rakdos Charm
- 1 Red Elemental Blast
- 1 Reflecting Pool
- 1 Saw in Half
- 1 Sensei's Divining Top
- 1 Shadowblood Ridge
- 1 Sign in Blood
- 1 Skullclamp
- 1 Smoldering Marsh
- 1 Snuff Out
- 1 Storm-Kiln Artist
- 1 Sulfurous Springs
- 9 Swamp
- 1 Swiftfoot Boots
- 1 Tainted Peak
- 1 The One Ring
- 1 The Reaver Cleaver
- 1 Tibalt's Trickery
- 1 Toxic Deluge
- 1 Treasure Map // Treasure Cove
- 1 Undying Malice
- 1 Vandalblast
- 1 Village Rites
- 1 Vexing Bauble
- 1 Xorn
- 1 Reckless Lackey
- 1 Blooming Blast
- 1 Ticket Tortoise
- 1 Wanted Scoundrels
- 1 Cavern of Souls

### 13. Bracket 5, Kinnan, Bonder Prodigy, combo

Pool 303. 99 cards, 1 block findings, repaired true (copy_limit, profile_off_band).

Judge: bracket 5, agrees. This is a Kinnan cEDH list: fast mana (Mana Vault, Grim Monolith, Basalt Monolith, Earthcraft), fetch/dual-heavy mana base, free/cheap counterspells (Flusterstorm, Muddle the Mixture, Narset's Reversal) plus protection like Shore Up, and Thassa's Oracle as the win. It is packed with two-card infinite mana combos (Devoted Druid + untappers, Freed from the Real/Pili-Pala, Basalt Monolith + Kinnan, Staff of Domination, Voltaic Construct) and clone/untap redundancy, plus Game Changers like Cyclonic Rift, Thassa's Oracle and Kinnan himself. It is optimized around the best combo strategy rather than a theme and can win on very early turns.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 33 | 27 to 33 |  |  |
| `tapped_land` | 0 | 0 to 2 |  |  |
| `colorless_land` | 0 | 0 to 6 |  |  |
| `color_sources` | 1.32 | 0.9 or more |  | sources of requirement: U 25 of 19, G 27.25 of 19 |
| `avg_mana_value` | 2.33 | 1.2 to 2.4 |  | over 66 nonland cards |
| `ramp` | 22 | 12 to 22 |  |  |
| `draw` | 11 | 8 to 18 |  |  |
| `removal` | 6 | 4 to 14 |  |  |
| `wipe` | 1 | 0 to 3 |  |  |
| `interaction` | 11 | 8 to 22 |  |  |
| `tutor` | 1 | none |  | Muddle the Mixture |
| `fast_mana` | 1 | none |  | Mana Vault |
| `game_changer` | 4 | none |  | Cyclonic Rift, Grim Monolith, Mana Vault, Thassa's Oracle |
| `mana_turn_four` | 4.7 | 5.5 or more | yes | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.71 | 0.65 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.03 | -9 to 0.3 |  | the commander costs 2 and comes down on turn 1.97 on average |

Goldfish over 10000 hands: the commander on turn 1.97, 4.7 mana on turn four, and 0.71 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Cyclonic Rift, Grim Monolith, Mana Vault, Thassa's Oracle. Mass land denial none. Extra turns none. Combos 6.
- Ioreth of the Healing House + Kiora's Follower (speed 5)
- Ioreth of the Healing House + Vizier of Tumbling Sands (speed 5)
- Ioreth of the Healing House + Kelpie Guide (speed 5)
- Kelpie Guide + Freed from the Real (speed 5, two cards)
- Kinnan, Bonder Prodigy + Grim Monolith + Spark Double (speed 5, two cards)
- Kinnan, Bonder Prodigy + Basalt Monolith (speed 5, two cards)

Findings:

- BLOCK `copy_limit`: : 2 copies, the limit is 1
- INFO `curve_summary`: average mana value 2.33 over 66 nonland cards
- WARN `profile_off_band`: the mana available on turn four is 4.7, and bracket 5 wants 5.5 or more (mean over 10000 hands)
- INFO `mana_pass`: the builder moved 7 cards of the mana base to bring the deck inside its power level

Cards:

- 8 Island
- 11 Forest
- 1 Breeding Pool
- 1 Command Tower
- 1 City of Brass
- 1 Mana Confluence
- 1 Yavimaya Coast
- 1 Dreamroot Cascade
- 1 Rejuvenating Springs
- 1 Exotic Orchard
- 1 Misty Rainforest
- 1 Flooded Strand
- 1 Polluted Delta
- 1 Scalding Tarn
- 1 Verdant Catacombs
- 1 Windswept Heath
- 1 Mana Vault
- 1 Grim Monolith
- 1 Basalt Monolith
- 1 Arbor Elf
- 1 Devoted Druid
- 1 Earthcraft
- 1 Gwenna, Eyes of Gaea
- 1 Ioreth of the Healing House
- 1 Kiora's Follower
- 1 Magus of the Candelabra
- 1 Sonic Screwdriver
- 1 Azusa's Many Journeys // Likeness of the Seeker
- 1 Spelunking
- 1 Horizon Explorer
- 1 Patriar's Seal
- 1 Kelpie Guide
- 1 Lost Jitte
- 1 Bender's Waterskin
- 1 Honor-Worn Shaku
- 1 Vizier of Tumbling Sands
- 1 Kiora, Behemoth Beckoner
- 1 Cerulean Wisps
- 1 Benefactor's Draught
- 1 Cloud of Faeries
- 1 Frantic Search
- 1 Refocus
- 1 Silent Hallcreeper
- 1 Twitch
- 1 Pip-Boy 3000
- 1 Sewer-veillance Cam
- 1 Staff of Domination
- 1 Yavimaya Elder
- 1 Flusterstorm
- 1 Biosynthic Burst
- 1 Fleeting Reflection
- 1 Legolas's Quick Reflexes
- 1 Magic Damper
- 1 Muddle the Mixture
- 1 Narset's Reversal
- 1 Shore Up
- 1 Unwind
- 1 Amphibian Downpour
- 1 Apathy
- 1 Bushwhack
- 1 Emerald Charm
- 1 Snap
- 1 Provoke
- 1 Cyclonic Rift
- 1 Voltaic Construct
- 1 Phyrexian Metamorph
- 1 Breaching Hippocamp
- 1 Mirrorhall Mimic // Ghastly Mimicry
- 1 Altered Ego
- 1 Clone
- 1 Clever Impersonator
- 1 Spark Double
- 1 Stunt Double
- 1 Sakashima's Student
- 1 Freed from the Real
- 1 Pili-Pala
- 1 Dramatic Reversal
- 1 Thassa's Oracle
- 1 Dream's Grip
- 1 Kinnan, Bonder Prodigy
- 1 High Stride
- 1 Helix Pinnacle

### 14. Bracket 5, Yuriko, the Tiger's Shadow, ninjas

Pool 213. 99 cards, 0 block findings, repaired false.

Judge: bracket 4, disagrees. This is a tuned Yuriko list with heavy fast mana (Sol Ring, Mana Vault, four Moxen, Lotus Petal, Dark Ritual), a dual-land-quality mana base including Underground Sea, and a dense free-interaction suite (Force of Negation, Pact, Fierce Guardianship, Deadly Rollick, Snuff Out). It blows well past the three-Game-Changer ceiling of Bracket 3 with Rhystic Study, Mystic Remora, Mana Drain, Fierce Guardianship and more, and can realistically kill around turn four off unblockable ninja triggers. It still leans on the ninja tribal theme rather than a dedicated compact win combo, so it sits at Optimized rather than cEDH.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 30 | 27 to 33 |  |  |
| `tapped_land` | 1 | 0 to 2 |  | Choked Estuary |
| `colorless_land` | 0 | 0 to 6 |  |  |
| `color_sources` | 1.51 | 0.9 or more |  | sources of requirement: U 32.75 of 19, B 28.75 of 19 |
| `avg_mana_value` | 2.09 | 1.2 to 2.4 |  | over 69 nonland cards |
| `ramp` | 20 | 12 to 22 |  |  |
| `draw` | 13 | 8 to 18 |  |  |
| `removal` | 9 | 4 to 14 |  |  |
| `wipe` | 1 | 0 to 3 |  |  |
| `interaction` | 14 | 8 to 22 |  |  |
| `tutor` | 0 | none |  |  |
| `fast_mana` | 10 | none |  | An Offer You Can't Refuse, Chrome Mox, Dark Ritual, Everflowing Chalice, Lotus Petal, Mana Vault, Mox Amber, Mox Diamond, Mox Opal, Sol Ring |
| `game_changer` | 5 | none |  | Chrome Mox, Fierce Guardianship, Mana Vault, Mox Diamond, Rhystic Study |
| `mana_turn_four` | 5.52 | 5.5 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.66 | 0.65 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.59 | -9 to 0.3 |  | the commander costs 3 and comes down on turn 2.41 on average |

Goldfish over 10000 hands: the commander on turn 2.41, 5.52 mana on turn four, and 0.66 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Chrome Mox, Fierce Guardianship, Mana Vault, Mox Diamond, Rhystic Study. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 2.09 over 69 nonland cards
- INFO `mana_pass`: the builder moved 4 cards of the mana base to bring the deck inside its power level

Cards:

- 8 Island
- 4 Swamp
- 1 Cavern of Souls
- 1 Choked Estuary
- 1 Clearwater Pathway // Murkwater Pathway
- 1 City of Brass
- 1 Command Tower
- 1 Darkwater Catacombs
- 1 Drowned Catacomb
- 1 Exotic Orchard
- 1 Gloomlake Verge
- 1 Mana Confluence
- 1 Morphic Pool
- 1 Secluded Courtyard
- 1 Sunken Ruins
- 1 Tainted Isle
- 1 Underground River
- 1 Underground Sea
- 1 Unclaimed Territory
- 1 Watery Grave
- 1 Arcane Signet
- 1 Chrome Mox
- 1 Dark Ritual
- 1 Dimir Signet
- 1 Fellwar Stone
- 1 Lotus Petal
- 1 Mana Drain
- 1 Mana Vault
- 1 Mind Stone
- 1 Mox Amber
- 1 Mox Diamond
- 1 Mox Opal
- 1 Ornithopter of Paradise
- 1 Prosperous Thief
- 1 Sol Ring
- 1 Talisman of Dominance
- 1 Wayfarer's Bauble
- 1 Baleful Strix
- 1 Brainstorm
- 1 Consider
- 1 Ingenious Infiltrator
- 1 Moon-Circuit Hacker
- 1 Mystic Remora
- 1 Ninja of the Deep Hours
- 1 Ponder
- 1 Preordain
- 1 Rhystic Study
- 1 Sensei's Divining Top
- 1 Skullclamp
- 1 Opt
- 1 An Offer You Can't Refuse
- 1 Arcane Denial
- 1 Counterspell
- 1 Fierce Guardianship
- 1 Flusterstorm
- 1 Force of Negation
- 1 Kaito, Bane of Nightmares
- 1 Lightning Greaves
- 1 Mental Misstep
- 1 Mithril Coat
- 1 Negate
- 1 Pact of Negation
- 1 Saiba Cryptomancer
- 1 Swan Song
- 1 Deadly Rollick
- 1 Feed the Swarm
- 1 Go for the Throat
- 1 Infernal Grasp
- 1 Into the Flood Maw
- 1 Pongify
- 1 Rapid Hybridization
- 1 Reality Shift
- 1 Snuff Out
- 1 Toxic Deluge
- 1 Biting-Palm Ninja
- 1 Silver-Fur Master
- 1 Foot Mystic
- 1 Futurist Operative
- 1 Kami of Restless Shadows
- 1 Kotose, the Silent Spider
- 1 Moonblade Shinobi
- 1 Mukotai Ambusher
- 1 Sakashima's Student
- 1 Shredder, Shadow Master
- 1 Taeko, the Patient Avalanche
- 1 Everflowing Chalice
- 1 Ashnod's Altar
- 1 Deadly Dispute
- 1 Thought Vessel

### 15. Bracket 5, Najeela, the Blade-Blossom, warriors combat

Pool 306. 99 cards, 0 block findings, repaired false.

Judge: bracket 3, disagrees. It's a solid warrior-tribal aggro build with two Game Changers (Najeela in the command zone plus Winota, Joiner of Forces) and the classic Najeela + Combat Celebrant loop, but that loop needs WUBRG per iteration and there's no fast mana or tutor package to assemble it early. The mana base is a fair but unoptimized mix of shocklands and basics, and the deck relies on grindy creature beats and small-body value rather than efficient interaction or stack-based protection. That puts it beyond a precon but well short of optimized — squarely Upgraded.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 33 | 27 to 33 |  |  |
| `tapped_land` | 2 | 0 to 2 |  | Base Camp, Path of Ancestry |
| `colorless_land` | 0 | 0 to 6 |  |  |
| `color_sources` | 0.64 | 0.9 or more | yes | sources of requirement: W 12.25 of 19, U 12.25 of 19, B 12.25 of 19, R 12.5 of 19, G 17.75 of 26 |
| `avg_mana_value` | 2.74 | 1.2 to 2.4 | yes | over 66 nonland cards |
| `ramp` | 16 | 12 to 22 |  |  |
| `draw` | 12 | 8 to 18 |  |  |
| `removal` | 9 | 4 to 14 |  |  |
| `wipe` | 1 | 0 to 3 |  |  |
| `interaction` | 14 | 8 to 22 |  |  |
| `tutor` | 1 | none |  | Cazur, Ruthless Stalker |
| `fast_mana` | 0 | none |  |  |
| `game_changer` | 0 | none |  |  |
| `mana_turn_four` | 3.97 | 5.5 or more | yes | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.71 | 0.65 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | 0.11 | -9 to 0.3 |  | the commander costs 3 and comes down on turn 3.11 on average |

Goldfish over 10000 hands: the commander on turn 3.11, 3.97 mana on turn four, and 0.71 of first hands hold two to four lands.

Content: Spellbook tag S. Game Changers none. Mass land denial none. Extra turns none. Combos 1.
- Najeela, the Blade-Blossom + Professional Face-Breaker (speed 5, two cards)

Findings:

- INFO `curve_summary`: average mana value 2.74 over 66 nonland cards
- WARN `profile_off_band`: the worst color's share of the sources it needs is 0.64, and bracket 5 wants 0.9 or more (sources of requirement: W 12.25 of 19, U 12.25 of 19, B 12.25 of 19, R 12.5 of 19, G 17.75 of 26)
- WARN `profile_off_band`: the average mana value of the nonland cards is 2.74, and bracket 5 wants 1.2 to 2.4
- WARN `profile_off_band`: the mana available on turn four is 3.97, and bracket 5 wants 5.5 or more (mean over 10000 hands)
- INFO `mana_pass`: the builder moved 13 cards of the mana base to bring the deck inside its power level

Cards:

- 1 Command Tower
- 1 Exotic Orchard
- 1 Base Camp
- 1 Path of Ancestry
- 1 Blood Crypt
- 1 Breeding Pool
- 1 Godless Shrine
- 1 Hallowed Fountain
- 1 Overgrown Tomb
- 1 Sacred Foundry
- 1 Steam Vents
- 1 Stomping Ground
- 1 Temple Garden
- 2 Plains
- 3 Island
- 3 Swamp
- 3 Mountain
- 9 Forest
- 1 Ardent Electromancer
- 1 Baylen, the Haymaker
- 1 Brigid, Clachan's Heart // Brigid, Doun's Mind
- 1 Evendo Brushrazer
- 1 Gimli of the Glittering Caves
- 1 Heronblade Elite
- 1 Neheb, Dreadhorde Champion
- 1 Professional Face-Breaker
- 1 Radha, Heir to Keld
- 1 Scuzzback Scrounger
- 1 Sonic the Hedgehog
- 1 Svella, Ice Shaper
- 1 Swashbuckler Extraordinaire
- 1 Ziatora's Envoy
- 1 Akki Ronin
- 1 Azra Oddsmaker
- 1 Mindblade Render
- 1 Nimble Trapfinder
- 1 Oakhame Adversary
- 1 Samut, Vizier of Naktamun
- 1 Savvy Hunter
- 1 Setessan Champion
- 1 Shakedown Heavy
- 1 Vanguard Suppressor
- 1 Yathan Tombguard
- 1 Zurgo Stormrender
- 1 Concerted Defense
- 1 Akiri, Fearless Voyager
- 1 Ainok Strike Leader
- 1 Eladamri, Lord of Leaves
- 1 Ezuri, Renegade Leader
- 1 Hakoda, Selfless Commander
- 1 Kashi-Tribe Elite
- 1 Linvala, Shield of Sea Gate
- 1 Multiclass Baldric
- 1 Rhys, the Evermore
- 1 Selfless Samurai
- 1 Seasoned Dungeoneer
- 1 Vexilus Praetor
- 1 Winota, Joiner of Forces
- 1 Accursed Marauder
- 1 Cacophony Scamp
- 1 Dreadhorde Butcher
- 1 Fleshbag Marauder
- 1 Goblin Cratermaker
- 1 Practiced Tactics
- 1 Aspiring Champion
- 1 Aysen Crusader
- 1 Bull-Rush Bruiser
- 1 Cazur, Ruthless Stalker
- 1 Doomskar Warrior
- 1 Heiko Yamazaki, the General
- 1 Kabira Outrider
- 1 Kimahri, Valiant Guardian
- 1 Raiyuu, Storm's Edge
- 1 Surrak, the Hunt Caller
- 1 Combat Celebrant
- 1 Rograkh, Son of Rohgahh
- 1 Goblin Chainwhirler
- 1 Elvish Reclaimer
- 1 Shang-Chi, Master of Kung Fu
- 1 Vorpal Sword
- 1 Archpriest of Iona
- 1 Stadium Headliner
- 1 Resolute Strike
- 1 Mudbutton Clanger

