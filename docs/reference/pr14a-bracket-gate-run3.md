# PR-14A bracket gate

Run date: 2026-09-14. Card snapshot: 2026-09-04.

Verdict: FAIL. 9 of 9 decks passed every block check, 9 sat in every band, 9 held no content violation, and the judge agreed with the bracket on 3 of 9 (33 percent, the bar is 80).

## Summary

| Measure | Value |
|---|---|
| Prompts | 9 |
| Decks returned | 9 |
| Decks with no block finding | 9 |
| Decks in every band | 9 |
| Decks with no content violation | 9 |
| Decks the endpoint did not check | 0 |
| Decks that needed a repair turn | 0 |
| Decks judged | 9 |
| Judge agreed with the bracket | 3 |
| Judge errors | 0 |
| Errors | 0 |
| Prompt version | 12 |
| Calls | 18 |
| Cost | $0.7355 |
| Time | 538 seconds |

## Run

- Suite `bracket`, run `pr14a-bracket-gate-run3`, on 2026-09-14, commit `6d61700`.
- Partial run over `1,2,3,4,5,6,7,8,9`. It reads the bars of its own items, and it never stands for the suite.
- Roles: generate on `gpt-5.6-terra` (openai, effort medium), judge on `claude-opus-5` (anthropic, effort medium), repair on `gpt-5.6-terra` (openai, effort medium).
- Versions: card snapshot 2026-09-04, generate prompt version 12.
- Calls: 18. Cost: $0.7355. Time: 538 seconds.

## Off-band features by key

Each row counts the decks whose feature sat outside its band after the repair turns.

| Feature | Decks off band |
|---|---|
| none | 0 |

## Decks

### 1. Bracket 1, Gishath, Sun's Avatar, dinosaurs

Pool 267. 99 cards, 0 block findings, repaired false.

Judge: bracket 3, disagrees. This is a focused Dinosaur tribal deck with no marked Game Changers, no mass land denial, no extra turns and no two-card infinite combos, so it stays legal for lower brackets on the rules checklist. However, the power level is clearly above precon: dual lands (Plateau, Taiga, shocks, Cavern of Souls), efficient two-mana ramp (Nature's Lore, Three Visits, Farseek), cost reducers, and a strong suite of card draw and big finishers that can land Gishath and snowball by turn six or seven. That combination of upgraded mana and consistent threats places it comfortably in Upgraded rather than Core, but it lacks the fast mana, tutors, and combo lines of an Optimized deck.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 38 | 36 to 41 |  |  |
| `tapped_land` | 1 | 0 to 14 |  | Jetmir's Garden |
| `colorless_land` | 0 | 0 to 4 |  |  |
| `color_sources` | 1.03 | 0.75 or more |  | sources of requirement: W 23.25 of 20, R 23.25 of 19, G 22.75 of 22 |
| `avg_mana_value` | 4.11 | 2.4 to 4.4 |  | over 61 nonland cards |
| `ramp` | 9 | 6 to 12 |  |  |
| `draw` | 10 | 6 to 14 |  |  |
| `removal` | 7 | 4 to 10 |  |  |
| `wipe` | 2 | 0 to 4 |  |  |
| `interaction` | 4 | 0 to 8 |  |  |
| `tutor` | 0 | 0 to 1 |  |  |
| `fast_mana` | 0 | 0 to 1 |  |  |
| `game_changer` | 0 | 0 to 0 |  |  |
| `mana_turn_four` | 4.58 | 3.5 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.76 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | 0.7 | -9 to 1.5 |  | the commander costs 8 and comes down on turn 8.7 on average |

Goldfish over 10000 hands: the commander on turn 8.7, 4.58 mana on turn four, and 0.76 of first hands hold two to four lands.

Content: Spellbook tag E. Game Changers none. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 4.11 over 61 nonland cards

Cards:

- 1 Command Tower
- 1 City of Brass
- 1 Mana Confluence
- 1 Exotic Orchard
- 1 Reflecting Pool
- 1 Cavern of Souls
- 1 Secluded Courtyard
- 1 Unclaimed Territory
- 1 Temple Garden
- 1 Stomping Ground
- 1 Sacred Foundry
- 1 Canopy Vista
- 1 Cinder Glade
- 1 Plateau
- 1 Taiga
- 1 Jetmir's Garden
- 1 Bountiful Promenade
- 7 Forest
- 7 Mountain
- 7 Plains
- 1 Arcane Signet
- 1 Birds of Paradise
- 1 Cultivate
- 1 Farseek
- 1 Nature's Lore
- 1 Three Visits
- 1 Thunderherd Migration
- 1 Drover of the Mighty
- 1 Ranging Raptors
- 1 Beast Whisperer
- 1 Garruk's Uprising
- 1 Guardian Project
- 1 Curious Altisaur
- 1 Earthshaker Dreadmaw
- 1 Ripjaw Raptor
- 1 Return of the Wildspeaker
- 1 Rishkar's Expertise
- 1 Shamanic Revelation
- 1 Vaultborn Tyrant
- 1 Akroma's Will
- 1 Boros Charm
- 1 Heroic Intervention
- 1 Lightning Greaves
- 1 Apex Altisaur
- 1 Burning Sun's Avatar
- 1 Bronzebeak Foragers
- 1 Itzquinth, Firstborn of Gishath
- 1 Needletooth Raptor
- 1 Ravenous Sailback
- 1 Thrashing Brontodon
- 1 Blasphemous Act
- 1 Wakening Sun's Avatar
- 1 Armored Kincaller
- 1 Belligerent Yearling
- 1 Commune with Dinosaurs
- 1 Deathgorge Scavenger
- 1 Dinosaur Egg
- 1 Dromosaur
- 1 Huatli's Raptor
- 1 Hunting Velociraptor
- 1 Kinjalli's Caller
- 1 Marauding Raptor
- 1 Otepec Huntmaster
- 1 Pugnacious Hammerskull
- 1 Raptor Companion
- 1 Raptor Hatchling
- 1 Sun-Collared Raptor
- 1 Sunfrill Imitator
- 1 Territorial Hammerskull
- 1 Ancient Imperiosaur
- 1 Carnage Tyrant
- 1 Etali, Primal Conqueror // Etali, Primal Sickness
- 1 Etali, Primal Storm
- 1 Ghalta, Primal Hunger
- 1 Ghalta, Stampede Tyrant
- 1 Goring Ceratops
- 1 Quartzwood Crasher
- 1 Regisaur Alpha
- 1 Tyrranax Rex
- 1 Verdant Sun's Avatar
- 1 Zetalpa, Primal Dawn

### 2. Bracket 1, Karlov of the Ghost Council, lifegain

Pool 303. 99 cards, 0 block findings, repaired false.

Judge: bracket 2, disagrees. This is a grindy lifegain-synergy deck with no marked Game Changers, no mass land denial, no extra turns, and no two-card infinite combo (Heliod and Vito/Defiant Bloodlord need a third piece like Walking Ballista, which isn't here). Its acceleration is clunky (Hierophant's Chalice, Hot Dog Cart, Altar of the Pantheon), the mana base is mostly basics with taplands, and there are no tutors or fast mana, so wins come from slowly stacking lifegain triggers over many turns. A few premium cards like Solitude, Archangel of Thune and Enduring Innocence push it to the top of Core, but the overall speed and consistency sit right around a strong precon rather than an Upgraded build.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 38 | 36 to 41 |  |  |
| `tapped_land` | 2 | 0 to 14 |  | Scoured Barrens, Temple of Silence |
| `colorless_land` | 0 | 0 to 4 |  |  |
| `color_sources` | 1.04 | 0.75 or more |  | sources of requirement: W 29 of 23, B 27 of 26 |
| `avg_mana_value` | 3.3 | 2.4 to 4.4 |  | over 61 nonland cards |
| `ramp` | 9 | 6 to 12 |  |  |
| `draw` | 10 | 6 to 14 |  |  |
| `removal` | 7 | 4 to 10 |  |  |
| `wipe` | 2 | 0 to 4 |  |  |
| `interaction` | 4 | 0 to 8 |  |  |
| `tutor` | 0 | 0 to 1 |  |  |
| `fast_mana` | 0 | 0 to 1 |  |  |
| `game_changer` | 0 | 0 to 0 |  |  |
| `mana_turn_four` | 4.29 | 3.5 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.76 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | 0.02 | -9 to 1.5 |  | the commander costs 2 and comes down on turn 2.02 on average |

Goldfish over 10000 hands: the commander on turn 2.02, 4.29 mana on turn four, and 0.76 of first hands hold two to four lands.

Content: Spellbook tag E. Game Changers none. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 3.30 over 61 nonland cards

Cards:

- 14 Plains
- 12 Swamp
- 1 Command Tower
- 1 Godless Shrine
- 1 Caves of Koilos
- 1 City of Brass
- 1 Mana Confluence
- 1 Exotic Orchard
- 1 Isolated Chapel
- 1 Scoured Barrens
- 1 Shattered Sanctum
- 1 Vault of Champions
- 1 Temple of Silence
- 1 Secluded Courtyard
- 1 Archivist of Oghma
- 1 Dawn of Hope
- 1 Cosmos Elixir
- 1 Enduring Innocence
- 1 Exemplar of Light
- 1 Mangara, the Diplomat
- 1 Sigarda's Splendor
- 1 Survival Cache
- 1 Well of Lost Dreams
- 1 The Gaffer
- 1 Alseid of Life's Bounty
- 1 Faith's Shield
- 1 Metropolis Reformer
- 1 Werefox Bodyguard
- 1 Altar of the Pantheon
- 1 Beza, the Bounding Spring
- 1 Hierophant's Chalice
- 1 Hot Dog Cart
- 1 Nuka-Cola Vending Machine
- 1 Orazca Relic
- 1 Pristine Talisman
- 1 Potioner's Trove
- 1 The Celestus
- 1 Ayli, Eternal Pilgrim
- 1 Nightmare's Thirst
- 1 Murderous Rider // Swift End
- 1 Henrika Domnathi // Henrika, Infernal Seer
- 1 Solitude
- 1 Summon: Ixion
- 1 Vona, Butcher of Magan
- 1 Ajani's Pridemate
- 1 Angel of Vitality
- 1 Cleric Class
- 1 Cleric of Life's Bond
- 1 Bloodthirsty Aerialist
- 1 Case of the Uneaten Feast
- 1 Celestial Unicorn
- 1 Griffin Aerie
- 1 Heliod, Sun-Crowned
- 1 Leyline of Hope
- 1 Light of Promise
- 1 Resplendent Angel
- 1 Righteous Valkyrie
- 1 Speaker of the Heavens
- 1 Vito, Thorn of the Dusk Rose
- 1 Voice of the Blessed
- 1 Wall of Limbs
- 1 Archangel of Thune
- 1 Attended Healer
- 1 Celestine, the Living Saint
- 1 Defiant Bloodlord
- 1 Divinity of Pride
- 1 Exalted Sunborn
- 1 Gideon's Company
- 1 Nykthos Paragon
- 1 Rhox Faithmender
- 1 Valkyrie Harbinger
- 1 Twinblade Paladin
- 1 Bloodbond Vampire
- 1 Fumigate
- 1 The Battle of Bywater

### 3. Bracket 1, Adeline, Resplendent Cathar, tokens

Pool 296. 99 cards, 0 block findings, repaired false.

Judge: bracket 3, disagrees. This is a focused mono-white go-wide token deck that is clearly tuned past precon level, with dense payoff stacking (Divine Visitation, Hero of Bladehold, Myrel, Intangible Virtue, Caretaker's Talent) and a solid removal/protection package, but it carries zero marked Game Changers, no mass land denial, no extra turns, and no two-card infinite combos. Its mana is slow and honest, 27 Plains with no Sol Ring, no fast mana, and no tutors, so it needs several turns of board building before Adeline turns lethal, typically around turn seven or eight. That combination of above-precon consistency with no optimized acceleration puts it at the lower end of Upgraded rather than Core or Optimized.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 38 | 36 to 41 |  |  |
| `tapped_land` | 3 | 0 to 14 |  | Path of Ancestry, Windbrisk Heights, Witch Enchanter // Witch-Blessed Meadow |
| `colorless_land` | 0 | 0 to 4 |  |  |
| `color_sources` | 1.68 | 0.75 or more |  | sources of requirement: W 38.75 of 23 |
| `avg_mana_value` | 3.46 | 2.4 to 4.4 |  | over 61 nonland cards |
| `ramp` | 9 | 6 to 12 |  |  |
| `draw` | 10 | 6 to 14 |  |  |
| `removal` | 7 | 4 to 10 |  |  |
| `wipe` | 2 | 0 to 4 |  |  |
| `interaction` | 4 | 0 to 8 |  |  |
| `tutor` | 0 | 0 to 1 |  |  |
| `fast_mana` | 0 | 0 to 1 |  |  |
| `game_changer` | 0 | 0 to 0 |  |  |
| `mana_turn_four` | 3.85 | 3.5 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.76 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | 0.12 | -9 to 1.5 |  | the commander costs 3 and comes down on turn 3.12 on average |

Goldfish over 10000 hands: the commander on turn 3.12, 3.85 mana on turn four, and 0.76 of first hands hold two to four lands.

Content: Spellbook tag E. Game Changers none. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 3.46 over 61 nonland cards

Cards:

- 27 Plains
- 1 Ancient Den
- 1 Castle Ardenvale
- 1 Command Tower
- 1 Eiganjo, Seat of the Empire
- 1 Kjeldoran Outpost
- 1 Minas Tirith
- 1 Path of Ancestry
- 1 Plaza of Heroes
- 1 Secluded Courtyard
- 1 Windbrisk Heights
- 1 Witch Enchanter // Witch-Blessed Meadow
- 1 Collector's Vault
- 1 Druidic Satchel
- 1 Keeper of the Accord
- 1 Monologue Tax
- 1 Nuka-Cola Vending Machine
- 1 Ojer Taq, Deepest Foundation // Temple of Civilization
- 1 Smuggler's Share
- 1 Tempting Contract
- 1 The Restoration of Eiganjo // Architect of Restoration
- 1 Bennie Bracks, Zoologist
- 1 Bygone Bishop
- 1 Caretaker's Talent
- 1 Chivalric Alliance
- 1 Court of Grace
- 1 Dawn of Hope
- 1 Faramir, Field Commander
- 1 Idol of Oblivion
- 1 Staff of the Storyteller
- 1 Wedding Announcement // Wedding Festivity
- 1 Aerial Assault
- 1 Generous Gift
- 1 Kellan's Lightblades
- 1 Righteous Confluence
- 1 Skyclave Apparition
- 1 Stroke of Midnight
- 1 The Wandering Emperor
- 1 Animation Module
- 1 Attended Healer
- 1 Charismatic Conqueror
- 1 Clarion Spirit
- 1 Divine Visitation
- 1 Felidar Retreat
- 1 Hanweir Militia Captain // Westvale Cult Leader
- 1 Horn of Gondor
- 1 Intangible Virtue
- 1 Luminarch Ascension
- 1 Oketra's Monument
- 1 Retrofitter Foundry
- 1 Rosie Cotton of South Lane
- 1 Skrelv's Hive
- 1 Song of the Worldsoul
- 1 Twilight Drover
- 1 Worthy Knight
- 1 Ajani's Chosen
- 1 Archangel Elspeth
- 1 Emeria Angel
- 1 God-Eternal Oketra
- 1 Hero of Bladehold
- 1 Mite Overseer
- 1 Myrel, Shield of Argive
- 1 Oketra the True
- 1 Phantom General
- 1 Requiem Angel
- 1 Sanctuary Warden
- 1 Threefold Thunderhulk
- 1 Hour of Reckoning
- 1 Martial Coup
- 1 Blessed Sanctuary
- 1 Lena, Selfless Champion
- 1 Rootborn Defenses
- 1 Spirit Bonds

### 4. Bracket 2, Gishath, Sun's Avatar, dinosaurs

Pool 267. 99 cards, 0 block findings, repaired false.

Judge: bracket 3, disagrees. This is a well-tuned but fair Dinosaur tribal deck: strong dual/tri lands, a solid ramp package (Sol Ring, Nature's Lore, Three Visits), and good card draw, but no Game Changers are marked and there are no two-card infinite combos or mass land denial. Its clock depends on resolving a six-mana commander and connecting, which is faster than a precon but not turn-four capable. That upgraded-beyond-precon power with clean fair-play restrictions places it squarely in Bracket 3.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 38 | 35 to 40 |  |  |
| `tapped_land` | 3 | 0 to 12 |  | Fortified Village, Game Trail, Jetmir's Garden |
| `colorless_land` | 0 | 0 to 4 |  |  |
| `color_sources` | 1.08 | 0.8 or more |  | sources of requirement: W 21.5 of 20, R 21.5 of 19, G 24.75 of 23 |
| `avg_mana_value` | 3.69 | 2.3 to 3.9 |  | over 61 nonland cards |
| `ramp` | 10 | 8 to 12 |  |  |
| `draw` | 11 | 8 to 14 |  |  |
| `removal` | 7 | 5 to 10 |  |  |
| `wipe` | 3 | 1 to 5 |  |  |
| `interaction` | 6 | 2 to 10 |  |  |
| `tutor` | 0 | 0 to 2 |  |  |
| `fast_mana` | 1 | 0 to 2 |  | Sol Ring |
| `game_changer` | 0 | 0 to 0 |  |  |
| `mana_turn_four` | 4.72 | 3.8 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.76 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | 0.41 | -9 to 1.2 |  | the commander costs 8 and comes down on turn 8.41 on average |

Goldfish over 10000 hands: the commander on turn 8.41, 4.72 mana on turn four, and 0.76 of first hands hold two to four lands.

Content: Spellbook tag E. Game Changers none. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 3.69 over 61 nonland cards
- INFO `bracket_cut`: bracket 2 does not allow this card, so the builder cut this card: Polyraptor, and added 1 basic land

Cards:

- 9 Forest
- 5 Plains
- 5 Mountain
- 1 Battlefield Forge
- 1 Bountiful Promenade
- 1 Brushland
- 1 Canopy Vista
- 1 Cavern of Souls
- 1 City of Brass
- 1 Clifftop Retreat
- 1 Command Tower
- 1 Exotic Orchard
- 1 Fortified Village
- 1 Game Trail
- 1 Jetmir's Garden
- 1 Karplusan Forest
- 1 Mana Confluence
- 1 Plateau
- 1 Rockfall Vale
- 1 Rootbound Crag
- 1 Sacred Foundry
- 1 Secluded Courtyard
- 1 Sol Ring
- 1 Arcane Signet
- 1 Birds of Paradise
- 1 Farseek
- 1 Nature's Lore
- 1 Three Visits
- 1 Rampant Growth
- 1 Cultivate
- 1 Kodama's Reach
- 1 Thunderherd Migration
- 1 Garruk's Uprising
- 1 Guardian Project
- 1 Curious Altisaur
- 1 Kutzil, Malamet Exemplar
- 1 Ripjaw Raptor
- 1 Runic Armasaur
- 1 Return of the Wildspeaker
- 1 Shamanic Revelation
- 1 Skullclamp
- 1 Sylvan Library
- 1 Toski, Bearer of Secrets
- 1 Akroma's Will
- 1 Boros Charm
- 1 Heroic Intervention
- 1 Lightning Greaves
- 1 Swiftfoot Boots
- 1 Temple Altisaur
- 1 Bronzebeak Foragers
- 1 Burning Sun's Avatar
- 1 Itzquinth, Firstborn of Gishath
- 1 Needletooth Raptor
- 1 Ravenous Sailback
- 1 Savage Stomp
- 1 Territorial Allosaurus
- 1 Blasphemous Act
- 1 Raging Swordtooth
- 1 Wakening Sun's Avatar
- 1 Amped Raptor
- 1 Armored Kincaller
- 1 Belligerent Yearling
- 1 Commune with Dinosaurs
- 1 Deathgorge Scavenger
- 1 Dromosaur
- 1 Huatli's Raptor
- 1 Kinjalli's Caller
- 1 Kinjalli's Sunwing
- 1 Marauding Raptor
- 1 Otepec Huntmaster
- 1 Rampaging Ferocidon
- 1 Raptor Companion
- 1 Carnage Tyrant
- 1 Etali, Primal Conqueror // Etali, Primal Sickness
- 1 Ghalta and Mavren
- 1 Ghalta, Stampede Tyrant
- 1 Goring Ceratops
- 1 Quartzwood Crasher
- 1 Regisaur Alpha
- 1 Shifting Ceratops
- 1 Thundering Spineback
- 1 Tyrranax Rex
- 1 Zetalpa, Primal Dawn

### 5. Bracket 2, Karlov of the Ghost Council, lifegain

Pool 303. 99 cards, 0 block findings, repaired false.

Judge: bracket 3, disagrees. This is a battlecruiser lifegain/Karlov deck with no marked Game Changers, no fast mana, and a basic-heavy mana base, but it does pack a genuine two-card infinite combo in Heliod, Sun-Crowned plus Archangel of Thune (infinite counters and life), which is too expensive to assemble in the first six turns. That combo, plus strong tutorless-but-efficient payoffs like Vito, Solitude, Umezawa's Jitte and three board wipes, puts it clearly past precon power. It lacks the speed, tutors, and staples of Bracket 4, so it sits comfortably in Upgraded.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 37 | 35 to 40 |  |  |
| `tapped_land` | 0 | 0 to 12 |  |  |
| `colorless_land` | 0 | 0 to 4 |  |  |
| `color_sources` | 1.02 | 0.8 or more |  | sources of requirement: W 26.5 of 26, B 23.75 of 23 |
| `avg_mana_value` | 3.27 | 2.3 to 3.9 |  | over 62 nonland cards |
| `ramp` | 10 | 8 to 12 |  |  |
| `draw` | 11 | 8 to 14 |  |  |
| `removal` | 7 | 5 to 10 |  |  |
| `wipe` | 3 | 1 to 5 |  |  |
| `interaction` | 6 | 2 to 10 |  |  |
| `tutor` | 0 | 0 to 2 |  |  |
| `fast_mana` | 1 | 0 to 2 |  | Legion's Landing // Adanto, the First Fort |
| `game_changer` | 0 | 0 to 0 |  |  |
| `mana_turn_four` | 4.42 | 3.8 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.75 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | 0.01 | -9 to 1.2 |  | the commander costs 2 and comes down on turn 2.01 on average |

Goldfish over 10000 hands: the commander on turn 2.01, 4.42 mana on turn four, and 0.75 of first hands hold two to four lands.

Content: Spellbook tag E. Game Changers none. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 3.27 over 62 nonland cards
- INFO `mana_pass`: the builder moved 1 card of the mana base to bring the deck inside its power level

Cards:

- 17 Plains
- 15 Swamp
- 1 Command Tower
- 1 Godless Shrine
- 1 Caves of Koilos
- 1 Vault of Champions
- 1 Shattered Sanctum
- 1 Legion's Landing // Adanto, the First Fort
- 1 Altar of the Pantheon
- 1 Pristine Talisman
- 1 Nuka-Cola Vending Machine
- 1 Orazca Relic
- 1 Phial of Galadriel
- 1 The Celestus
- 1 Potioner's Trove
- 1 Life Insurance
- 1 Cryptolith Fragment // Aurora of Emrakul
- 1 Archivist of Oghma
- 1 Dawn of Hope
- 1 Cosmos Elixir
- 1 Convalescent Care
- 1 Enduring Innocence
- 1 Exemplar of Light
- 1 Mangara, the Diplomat
- 1 Markov Purifier
- 1 Sigarda's Splendor
- 1 Well of Lost Dreams
- 1 Vampiric Rites
- 1 Alseid of Life's Bounty
- 1 Courageous Resolve
- 1 Faith's Shield
- 1 Metropolis Reformer
- 1 Restoration Magic
- 1 Werefox Bodyguard
- 1 Ayli, Eternal Pilgrim
- 1 Nightmare's Thirst
- 1 Murderous Rider // Swift End
- 1 Solitude
- 1 Umezawa's Jitte
- 1 Vona, Butcher of Magan
- 1 Witch of the Moors
- 1 Fumigate
- 1 Kaya's Wrath
- 1 The Battle of Bywater
- 1 Ajani's Pridemate
- 1 Angel of Vitality
- 1 Blood Artist
- 1 Bloodthirsty Aerialist
- 1 Cleric Class
- 1 Cleric of Life's Bond
- 1 Griffin Aerie
- 1 Heliod, Sun-Crowned
- 1 Light of Promise
- 1 Resplendent Angel
- 1 Righteous Valkyrie
- 1 Vito, Thorn of the Dusk Rose
- 1 Voice of the Blessed
- 1 Archangel of Thune
- 1 Attended Healer
- 1 Celestine, the Living Saint
- 1 Cliffhaven Vampire
- 1 Defiant Bloodlord
- 1 Divinity of Pride
- 1 Enduring Tenacity
- 1 Gideon's Company
- 1 Liesa, Forgotten Archangel
- 1 Lyra Dawnbringer
- 1 Nykthos Paragon
- 1 Valkyrie Harbinger

### 6. Bracket 2, Adeline, Resplendent Cathar, tokens

Pool 296. 99 cards, 0 block findings, repaired false.

Judge: bracket 2, agrees. This is a straightforward mono-white go-wide tokens deck with 37 Plains, no fast mana, no tutors, and joke-tier ramp artifacts (Goldvein Pick, Bucknard's Everfull Purse, Coin of Mastery), so it grinds out a board over many turns rather than closing fast. There are no marked Game Changers, no mass land denial, no extra turns, and no two-card infinite combos — just token payoffs like Divine Visitation and Mondrak plus basic removal. Its power and clunky consistency sit right around precon level, making it a Core deck.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 37 | 35 to 40 |  |  |
| `tapped_land` | 0 | 0 to 12 |  |  |
| `colorless_land` | 0 | 0 to 4 |  |  |
| `color_sources` | 1.48 | 0.8 or more |  | sources of requirement: W 38.5 of 26 |
| `avg_mana_value` | 3.37 | 2.3 to 3.9 |  | over 62 nonland cards |
| `ramp` | 10 | 8 to 12 |  |  |
| `draw` | 11 | 8 to 14 |  |  |
| `removal` | 7 | 5 to 10 |  |  |
| `wipe` | 3 | 1 to 5 |  |  |
| `interaction` | 6 | 2 to 10 |  |  |
| `tutor` | 0 | 0 to 2 |  |  |
| `fast_mana` | 1 | 0 to 2 |  | Legion's Landing // Adanto, the First Fort |
| `game_changer` | 0 | 0 to 0 |  |  |
| `mana_turn_four` | 3.99 | 3.8 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.75 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | 0.02 | -9 to 1.2 |  | the commander costs 3 and comes down on turn 3.02 on average |

Goldfish over 10000 hands: the commander on turn 3.02, 3.99 mana on turn four, and 0.75 of first hands hold two to four lands.

Content: Spellbook tag E. Game Changers none. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 3.37 over 62 nonland cards

Cards:

- 37 Plains
- 1 Bennie Bracks, Zoologist
- 1 Bygone Bishop
- 1 Caretaker's Talent
- 1 Chivalric Alliance
- 1 Court of Grace
- 1 Dawn of Hope
- 1 Idol of Oblivion
- 1 Staff of the Storyteller
- 1 Thorough Investigation
- 1 Wedding Announcement // Wedding Festivity
- 1 Wojek Investigator
- 1 Appa, Steadfast Guardian
- 1 Blessed Sanctuary
- 1 Rootborn Defenses
- 1 Spirit Bonds
- 1 Squad Commander
- 1 Summon: Knights of Round
- 1 Bucknard's Everfull Purse
- 1 Coin of Mastery
- 1 Druidic Satchel
- 1 Fishing Gear
- 1 Golden Guardian // Gold-Forge Garrison
- 1 Goldvein Pick
- 1 Karn, Living Legacy
- 1 Keeper of the Accord
- 1 Legion's Landing // Adanto, the First Fort
- 1 Monologue Tax
- 1 Aerial Assault
- 1 Banishing Slash
- 1 Generous Gift
- 1 Kellan's Lightblades
- 1 Skyclave Apparition
- 1 Stroke of Midnight
- 1 The Wandering Emperor
- 1 Anointer Priest
- 1 Clarion Spirit
- 1 Divine Visitation
- 1 Felidar Retreat
- 1 Horn of Gondor
- 1 Intangible Virtue
- 1 Luminarch Ascension
- 1 Ocelot Pride
- 1 Oketra's Monument
- 1 Rosie Cotton of South Lane
- 1 Skrelv's Hive
- 1 Song of the Worldsoul
- 1 Worthy Knight
- 1 Archangel Elspeth
- 1 Attended Healer
- 1 Basri's Lieutenant
- 1 Cemetery Protector
- 1 Emeria Angel
- 1 Gideon, Ally of Zendikar
- 1 God-Eternal Oketra
- 1 Hero of Bladehold
- 1 Mondrak, Glory Dominus
- 1 Myrel, Shield of Argive
- 1 Oketra the True
- 1 Requiem Angel
- 1 Elspeth Tirel
- 1 Hour of Reckoning
- 1 Martial Coup

### 7. Bracket 3, Karlov of the Ghost Council, lifegain

Pool 303. 99 cards, 0 block findings, repaired false.

Judge: bracket 3, agrees. This is a lifegain-payoff Karlov deck with no marked Game Changers, no mass land denial, and no reliable two-card infinite (Heliod, Sanguine Bond and Vito all need a missing third piece like Walking Ballista or Exquisite Blood), so it clears the Bracket 3 restrictions. Its power sits above precon level thanks to premium cards like Solitude, Vein Ripper, Umezawa's Jitte, Archangel of Thune and Aetherflux Reservoir, plus real sweepers and a dual/pain-land mana base. The clunky ramp (Colossal Plow, Hot Dog Cart) and grindy value plan keep it well short of Bracket 4 speed.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 36 | 34 to 38 |  |  |
| `tapped_land` | 2 | 0 to 9 |  | Scoured Barrens, Temple of Silence |
| `colorless_land` | 0 | 0 to 4 |  |  |
| `color_sources` | 1.03 | 0.85 or more |  | sources of requirement: W 26.75 of 23, B 28.75 of 28 |
| `avg_mana_value` | 3.21 | 2 to 3.5 |  | over 63 nonland cards |
| `ramp` | 10 | 8 to 13 |  |  |
| `draw` | 11 | 8 to 14 |  |  |
| `removal` | 9 | 6 to 12 |  |  |
| `wipe` | 3 | 2 to 5 |  |  |
| `interaction` | 8 | 4 to 12 |  |  |
| `tutor` | 0 | 0 to 5 |  |  |
| `fast_mana` | 0 | 0 to 3 |  |  |
| `game_changer` | 0 | 0 to 3 |  |  |
| `mana_turn_four` | 4.28 | 4.2 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.74 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | 0.03 | -9 to 0.8 |  | the commander costs 2 and comes down on turn 2.03 on average |

Goldfish over 10000 hands: the commander on turn 2.03, 4.28 mana on turn four, and 0.74 of first hands hold two to four lands.

Content: Spellbook tag E. Game Changers none. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 3.21 over 63 nonland cards
- INFO `mana_pass`: the builder moved 1 card of the mana base to bring the deck inside its power level

Cards:

- 11 Plains
- 13 Swamp
- 1 Command Tower
- 1 Godless Shrine
- 1 Caves of Koilos
- 1 Isolated Chapel
- 1 Shattered Sanctum
- 1 Vault of Champions
- 1 Mana Confluence
- 1 City of Brass
- 1 Reflecting Pool
- 1 Exotic Orchard
- 1 Scoured Barrens
- 1 Temple of Silence
- 1 Colossal Plow
- 1 Hot Dog Cart
- 1 Oasis Gardener
- 1 Altar of the Pantheon
- 1 Orazca Relic
- 1 Pristine Talisman
- 1 Nuka-Cola Vending Machine
- 1 Phial of Galadriel
- 1 The Celestus
- 1 Crypt Ghast
- 1 Archivist of Oghma
- 1 Dawn of Hope
- 1 Enduring Innocence
- 1 Well of Lost Dreams
- 1 Cosmos Elixir
- 1 Mangara, the Diplomat
- 1 Sigarda's Splendor
- 1 The Gaffer
- 1 Vampiric Rites
- 1 Markov Purifier
- 1 Restless Bloodseeker // Bloodsoaked Reveler
- 1 Alseid of Life's Bounty
- 1 Faith's Shield
- 1 Courageous Resolve
- 1 Caduceus, Staff of Hermes
- 1 Metropolis Reformer
- 1 Restoration Magic
- 1 Sword of Light and Shadow
- 1 Werefox Bodyguard
- 1 Aetherflux Reservoir
- 1 Ayli, Eternal Pilgrim
- 1 Murderous Rider // Swift End
- 1 Solitude
- 1 Umezawa's Jitte
- 1 Vein Ripper
- 1 Vona, Butcher of Magan
- 1 Nightmare's Thirst
- 1 Henrika Domnathi // Henrika, Infernal Seer
- 1 Fumigate
- 1 Kaya's Wrath
- 1 The Meathook Massacre
- 1 Ajani's Pridemate
- 1 Bloodthirsty Aerialist
- 1 Cleric Class
- 1 Heliod, Sun-Crowned
- 1 Light of Promise
- 1 Resplendent Angel
- 1 Righteous Valkyrie
- 1 Sanguine Bond
- 1 Vito, Thorn of the Dusk Rose
- 1 Voice of the Blessed
- 1 Archangel of Thune
- 1 Attended Healer
- 1 Bloodbond Vampire
- 1 Cliffhaven Vampire
- 1 Enduring Tenacity
- 1 Gideon's Company
- 1 Qala, Ajani's Pridemate
- 1 Star Charter
- 1 Twinblade Paladin
- 1 Vampire Scrivener
- 1 Wax-Wane Witness
- 1 Celestine, the Living Saint

### 8. Bracket 3, Denethor, Ruling Steward, aristocrats sacrifice

Pool 295. 99 cards, 0 block findings, repaired false.

Judge: bracket 4, disagrees. Despite carrying no marked Game Changers, this is a tightly tuned aristocrats engine with a dual-heavy, painland/Cavern mana base and multiple redundant two-card infinite loops \u2014 Ashnod's Altar or Phyrexian Altar with Pawn of Ulamog/Sifter of Skulls (or Pitiless Plunderer lines) goes infinite, and the deck runs a dozen Blood Artist-style drains plus Woe Strider/Carrion Feeder free sac outlets as payoffs. Cheap enablers like Culling the Weak, Village Rites and Deadly Dispute let it assemble and win well before turn six, which exceeds Bracket 3's 'no cheap early two-card combo' guidance. It isn't metagame-tuned cEDH, so Optimized is the right home.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 36 | 34 to 38 |  |  |
| `tapped_land` | 1 | 0 to 9 |  | Path of Ancestry |
| `colorless_land` | 0 | 0 to 4 |  |  |
| `color_sources` | 1.13 | 0.85 or more |  | sources of requirement: W 26 of 19, B 26 of 23 |
| `avg_mana_value` | 2.92 | 2 to 3.5 |  | over 63 nonland cards |
| `ramp` | 11 | 8 to 13 |  |  |
| `draw` | 11 | 8 to 14 |  |  |
| `removal` | 9 | 6 to 12 |  |  |
| `wipe` | 3 | 2 to 5 |  |  |
| `interaction` | 8 | 4 to 12 |  |  |
| `tutor` | 0 | 0 to 5 |  |  |
| `fast_mana` | 2 | 0 to 3 |  | Culling the Weak, Sacrifice |
| `game_changer` | 0 | 0 to 3 |  |  |
| `mana_turn_four` | 4.26 | 4.2 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.74 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.11 | -9 to 0.8 |  | the commander costs 3 and comes down on turn 2.89 on average |

Goldfish over 10000 hands: the commander on turn 2.89, 4.26 mana on turn four, and 0.74 of first hands hold two to four lands.

Content: Spellbook tag E. Game Changers none. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 2.92 over 63 nonland cards
- INFO `mana_pass`: the builder moved 2 cards of the mana base to bring the deck inside its power level

Cards:

- 10 Plains
- 10 Swamp
- 1 Caves of Koilos
- 1 Cavern of Souls
- 1 City of Brass
- 1 Command Tower
- 1 Exotic Orchard
- 1 Fetid Heath
- 1 Godless Shrine
- 1 Isolated Chapel
- 1 Mana Confluence
- 1 Path of Ancestry
- 1 Reflecting Pool
- 1 Scrubland
- 1 Secluded Courtyard
- 1 Shattered Sanctum
- 1 Unclaimed Territory
- 1 Vault of Champions
- 1 Ashnod's Altar
- 1 Culling the Weak
- 1 Deadly Dispute
- 1 Pawn of Ulamog
- 1 Phyrexian Altar
- 1 Pitiless Plunderer
- 1 Priest of Forgotten Gods
- 1 Sifter of Skulls
- 1 Skullport Merchant
- 1 Warren Soultrader
- 1 Baron Bertram Graywater
- 1 Bushmeat Poacher
- 1 Corrupted Conviction
- 1 Disciple of Bolas
- 1 Ecstatic Awakener // Awoken Demon
- 1 Relic Vial
- 1 Shadowheart, Dark Justiciar
- 1 Smothering Abomination
- 1 Tevesh Szat, Doom of Fools
- 1 Thraxodemon
- 1 Village Rites
- 1 Attrition
- 1 Ayli, Eternal Pilgrim
- 1 Blasting Station
- 1 Bone Shards
- 1 Bone Splinters
- 1 Dictate of Erebos
- 1 Eaten Alive
- 1 Grave Pact
- 1 Yawgmoth, Thran Physician
- 1 Cartel Aristocrat
- 1 Cultist of the Absolute
- 1 Dark Privilege
- 1 Fanatical Devotion
- 1 Flare of Fortitude
- 1 Gift of Doom
- 1 Nightmare Shepherd
- 1 Promise of Tomorrow
- 1 Bastion of Remembrance
- 1 Bartolomé del Presidio
- 1 Blood Artist
- 1 Carrion Feeder
- 1 Cruel Celebrant
- 1 Elas il-Kor, Sadistic Pilgrim
- 1 Fleshtaker
- 1 Viscera Seer
- 1 Woe Strider
- 1 Zulaport Cutthroat
- 1 Corpse Harvester
- 1 Elenda, the Dusk Rose
- 1 Falkenrath Noble
- 1 Ghoulcaller Gisa
- 1 Liesa, Forgotten Archangel
- 1 Mondrak, Glory Dominus
- 1 Ratadrabik of Urborg
- 1 Sadistic Hypnotist
- 1 Syr Konrad, the Grim
- 1 Vindictive Vampire
- 1 Minister of Pain
- 1 The Meathook Massacre
- 1 Toxic Deluge
- 1 Master of Dark Rites
- 1 Sacrifice

### 9. Bracket 3, Zada, Hedron Grinder, goblins

Pool 294. 99 cards, 0 block findings, repaired false.

Judge: bracket 3, agrees. This is a well-tuned mono-red goblin/Zada deck with one Game Changer (Jeska's Will), Wheel of Fortune, Sol Ring and rituals, plus efficient tribal payoffs like Muxus and Krenko that can end games around turn six or seven. It sits above precon power but lacks the fast mana density, tutor suite, and reliable early two-card kill of Bracket 4 — the Kiki-Jiki/Conspicuous Snoop line is fragile, needs setup, and isn't cheap or consistent in the first six turns. No mass land denial or extra turns, and only a single marked Game Changer, so Upgraded is the right home.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 37 | 34 to 38 |  |  |
| `tapped_land` | 0 | 0 to 9 |  |  |
| `colorless_land` | 0 | 0 to 4 |  |  |
| `color_sources` | 1.67 | 0.85 or more |  | sources of requirement: R 38.5 of 23 |
| `avg_mana_value` | 2.77 | 2 to 3.5 |  | over 62 nonland cards |
| `ramp` | 10 | 8 to 13 |  |  |
| `draw` | 11 | 8 to 14 |  |  |
| `removal` | 9 | 6 to 12 |  |  |
| `wipe` | 3 | 2 to 5 |  |  |
| `interaction` | 8 | 4 to 12 |  |  |
| `tutor` | 0 | 0 to 5 |  |  |
| `fast_mana` | 2 | 0 to 3 |  | Brightstone Ritual, Sol Ring |
| `game_changer` | 1 | 0 to 3 |  | Jeska's Will |
| `mana_turn_four` | 4.68 | 4.2 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.75 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.44 | -9 to 0.8 |  | the commander costs 4 and comes down on turn 3.56 on average |

Goldfish over 10000 hands: the commander on turn 3.56, 4.68 mana on turn four, and 0.75 of first hands hold two to four lands.

Content: Spellbook tag P. Game Changers Jeska's Will. Mass land denial none. Extra turns none. Combos 2.
- Krenko, Mob Boss + Skirk Prospector + Goblin Warchief (speed 5)
- Krenko, Mob Boss + Skirk Prospector + Goblin Chieftain (speed 5)

Findings:

- INFO `curve_summary`: average mana value 2.77 over 62 nonland cards
- INFO `bracket_cut`: bracket 3 does not allow this card, so the builder cut this card: Akki Battle Squad, and added 1 basic land

Cards:

- 37 Mountain
- 1 Sol Ring
- 1 Arcane Signet
- 1 Fellwar Stone
- 1 Mind Stone
- 1 Wayfarer's Bauble
- 1 Skirk Prospector
- 1 Wily Goblin
- 1 Brightstone Ritual
- 1 Jeska's Will
- 1 Krark-Clan Stoker
- 1 Faithless Looting
- 1 Cathartic Reunion
- 1 Demand Answers
- 1 Thrill of Possibility
- 1 Skullclamp
- 1 Mask of Memory
- 1 Idol of Oblivion
- 1 Sensation Gorger
- 1 Wheel of Fortune
- 1 Rummaging Goblin
- 1 Goblin Picker
- 1 Pyroblast
- 1 Red Elemental Blast
- 1 Tibalt's Trickery
- 1 Lightning Greaves
- 1 Swiftfoot Boots
- 1 Champion's Helm
- 1 Mithril Coat
- 1 Whispersilk Cloak
- 1 Arms Dealer
- 1 Gempalm Incinerator
- 1 Goblin Cratermaker
- 1 Goblin Grenade
- 1 Goblin Sharpshooter
- 1 Pashalik Mons
- 1 Siege-Gang Commander
- 1 Sparksmith
- 1 Volley Veteran
- 1 Conspicuous Snoop
- 1 Goblin Chieftain
- 1 Goblin Warchief
- 1 Mogg War Marshal
- 1 Goblin Instigator
- 1 Krenko, Tin Street Kingpin
- 1 Goblin Rabblemaster
- 1 Rundvelt Hordemaster
- 1 Legion Loyalist
- 1 Krenko, Baron of Tin Street
- 1 Battle Squadron
- 1 Beetleback Chief
- 1 Battle-Rattle Shaman
- 1 Goblin Gang Leader
- 1 Krenko, Mob Boss
- 1 Kiki-Jiki, Mirror Breaker
- 1 Muxus, Goblin Grandee
- 1 Reckless One
- 1 Goblin Goon
- 1 Goblin Marshal
- 1 Goblin Ringleader
- 1 Blasphemous Act
- 1 Chain Reaction
- 1 Krark-Clan Shaman

