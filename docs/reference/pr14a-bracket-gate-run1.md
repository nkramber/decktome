# PR-14A bracket gate

Run date: 2026-09-02. Card snapshot: 2026-09-02.

Verdict: FAIL. 15 of 15 decks passed every block check, 8 sat in every band, 15 held no content violation, and the judge agreed with the bracket on 0 of 0 (0 percent, the bar is 80).

15 decks got no judge verdict, because the judge call failed. The judge bar can not pass without one.

## Summary

| Measure | Value |
|---|---|
| Prompts | 15 |
| Decks returned | 15 |
| Decks with no block finding | 15 |
| Decks in every band | 8 |
| Decks with no content violation | 15 |
| Decks the endpoint did not check | 0 |
| Decks that needed a repair turn | 9 |
| Decks judged | 0 |
| Judge agreed with the bracket | 0 |
| Judge errors | 15 |
| Errors | 0 |
| Prompt version | 12 |
| Calls | 46 |
| Cost | $2.0783 |
| Time | 1966 seconds |

## Off-band features by key

Each row counts the decks whose feature sat outside its band after the repair turns.

| Feature | Decks off band |
|---|---|
| `mana_turn_four` | 7 |
| `avg_mana_value` | 4 |
| `color_sources` | 2 |
| `commander_turn_over_mv` | 2 |
| `tapped_land` | 1 |

## Decks

### 1. Bracket 1, Gishath, Sun's Avatar, dinosaurs

Pool 269. 99 cards, 0 block findings, repaired false.

Judge: error, judge bracket: llm anthropic/claude-opus-5: terminal (http 400): POST "https://api.anthropic.com/v1/messages": 400 Bad Request (Request-ID: req_011Ceev7HhDpwaBAWGXypt6n) (Workspace-ID: wrkspc_01DKn3qmnB3ZMerq6QsCdj6C) {"type":"error","error":{"type":"invalid_request_error","message":"output_config.format.schema: For 'integer' type, properties maximum, minimum are not supported"},"request_id":"req_011Ceev7HhDpwaBAWGXypt6n"}

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 38 | 36 to 41 |  |  |
| `tapped_land` | 2 | 0 to 14 |  | Jetmir's Garden, Path of Ancestry |
| `colorless_land` | 0 | 0 to 4 |  |  |
| `color_sources` | 0.88 | 0.75 or more |  | sources of requirement: W 17.5 of 20, R 17.5 of 19, G 23.25 of 22 |
| `avg_mana_value` | 4.38 | 2.4 to 4.4 |  | over 61 nonland cards |
| `ramp` | 9 | 6 to 12 |  |  |
| `draw` | 10 | 6 to 14 |  |  |
| `removal` | 7 | 4 to 10 |  |  |
| `wipe` | 2 | 0 to 4 |  |  |
| `interaction` | 4 | 0 to 8 |  |  |
| `tutor` | 0 | 0 to 1 |  |  |
| `fast_mana` | 0 | 0 to 1 |  |  |
| `game_changer` | 0 | 0 to 0 |  |  |
| `mana_turn_four` | 4.47 | 3.5 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.76 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | 0.73 | -9 to 1.5 |  | the commander costs 8 and comes down on turn 8.73 on average |

Goldfish over 10000 hands: the commander on turn 8.73, 4.47 mana on turn four, and 0.76 of first hands hold two to four lands.

Content: Spellbook tag E. Game Changers none. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 4.38 over 61 nonland cards

Cards:

- 12 Forest
- 8 Mountain
- 8 Plains
- 1 Command Tower
- 1 Jetmir's Garden
- 1 Canopy Vista
- 1 Cinder Glade
- 1 Stomping Ground
- 1 Temple Garden
- 1 Sacred Foundry
- 1 Path of Ancestry
- 1 Unclaimed Territory
- 1 Cavern of Souls
- 1 Arcane Signet
- 1 Cultivate
- 1 Farseek
- 1 Nature's Lore
- 1 Kodama's Reach
- 1 Thunderherd Migration
- 1 Drover of the Mighty
- 1 Ranging Raptors
- 1 Hulking Raptor
- 1 Beast Whisperer
- 1 Guardian Project
- 1 Garruk's Uprising
- 1 Curious Altisaur
- 1 Earthshaker Dreadmaw
- 1 Ripjaw Raptor
- 1 Rishkar's Expertise
- 1 Return of the Wildspeaker
- 1 Shamanic Revelation
- 1 Vaultborn Tyrant
- 1 Akroma's Will
- 1 Boros Charm
- 1 Heroic Intervention
- 1 Inspiring Call
- 1 Apex Altisaur
- 1 Bronzebeak Foragers
- 1 Burning Sun's Avatar
- 1 Savage Stomp
- 1 Ravenous Sailback
- 1 Trumpeting Carnosaur
- 1 Zacama, Primal Calamity
- 1 Blasphemous Act
- 1 Wakening Sun's Avatar
- 1 Amped Raptor
- 1 Kinjalli's Caller
- 1 Otepec Huntmaster
- 1 Marauding Raptor
- 1 Dinosaur Stampede
- 1 Huatli's Raptor
- 1 Kinjalli's Sunwing
- 1 Hunting Velociraptor
- 1 Raptor Companion
- 1 Raptor Hatchling
- 1 Orazca Frillback
- 1 Pugnacious Hammerskull
- 1 Sky Terror
- 1 Territorial Hammerskull
- 1 Sun-Collared Raptor
- 1 Sunfrill Imitator
- 1 Regal Imperiosaur
- 1 Ancient Brontodon
- 1 Carnage Tyrant
- 1 Etali, Primal Conqueror // Etali, Primal Sickness
- 1 Ghalta and Mavren
- 1 Ghalta, Primal Hunger
- 1 Ghalta, Stampede Tyrant
- 1 Goring Ceratops
- 1 Pantlaza, Sun-Favored
- 1 Quartzwood Crasher
- 1 Regisaur Alpha
- 1 Tyrranax Rex
- 1 Zetalpa, Primal Dawn

### 2. Bracket 1, Karlov of the Ghost Council, lifegain

Pool 303. 99 cards, 0 block findings, repaired false.

Judge: error, judge bracket: llm anthropic/claude-opus-5: terminal (http 400): POST "https://api.anthropic.com/v1/messages": 400 Bad Request (Request-ID: req_011CeevASKioCuvHH2Qscd7d) (Workspace-ID: wrkspc_01DKn3qmnB3ZMerq6QsCdj6C) {"type":"error","error":{"type":"invalid_request_error","message":"output_config.format.schema: For 'integer' type, properties maximum, minimum are not supported"},"request_id":"req_011CeevASKioCuvHH2Qscd7d"}

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 38 | 36 to 41 |  |  |
| `tapped_land` | 1 | 0 to 14 |  | Restless Fortress |
| `colorless_land` | 1 | 0 to 4 |  | High Market |
| `color_sources` | 1.04 | 0.75 or more |  | sources of requirement: W 27 of 26, B 27.75 of 23 |
| `avg_mana_value` | 3.51 | 2.4 to 4.4 |  | over 61 nonland cards |
| `ramp` | 9 | 6 to 12 |  |  |
| `draw` | 10 | 6 to 14 |  |  |
| `removal` | 7 | 4 to 10 |  |  |
| `wipe` | 2 | 0 to 4 |  |  |
| `interaction` | 4 | 0 to 8 |  |  |
| `tutor` | 0 | 0 to 1 |  |  |
| `fast_mana` | 0 | 0 to 1 |  |  |
| `game_changer` | 0 | 0 to 0 |  |  |
| `mana_turn_four` | 4.28 | 3.5 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.76 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | 0.01 | -9 to 1.5 |  | the commander costs 2 and comes down on turn 2.01 on average |

Goldfish over 10000 hands: the commander on turn 2.01, 4.28 mana on turn four, and 0.76 of first hands hold two to four lands.

Content: Spellbook tag E. Game Changers none. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 3.51 over 61 nonland cards

Cards:

- 13 Plains
- 13 Swamp
- 1 Caves of Koilos
- 1 City of Brass
- 1 Command Tower
- 1 Godless Shrine
- 1 High Market
- 1 Isolated Chapel
- 1 Mana Confluence
- 1 Nykthos, Shrine to Nyx
- 1 Reflecting Pool
- 1 Restless Fortress
- 1 Shattered Sanctum
- 1 Vault of Champions
- 1 Altar of the Pantheon
- 1 Bounty Board
- 1 Carmen, Cruel Skymarcher
- 1 Crypt Ghast
- 1 Hierophant's Chalice
- 1 Nuka-Cola Vending Machine
- 1 Orazca Relic
- 1 Potioner's Trove
- 1 Pristine Talisman
- 1 Archivist of Oghma
- 1 Convalescent Care
- 1 Cosmos Elixir
- 1 Dawn of Hope
- 1 Enduring Innocence
- 1 Exemplar of Light
- 1 Haliya, Guided by Light
- 1 Mangara, the Diplomat
- 1 Sigarda's Splendor
- 1 Well of Lost Dreams
- 1 Alseid of Life's Bounty
- 1 Faith's Shield
- 1 Restoration Magic
- 1 Werefox Bodyguard
- 1 Ayli, Eternal Pilgrim
- 1 Cavalier of Night
- 1 Murderous Rider // Swift End
- 1 Nightmare's Thirst
- 1 Solitude
- 1 Vona, Butcher of Magan
- 1 Witch of the Moors
- 1 Aerith Gainsborough
- 1 Ajani's Pridemate
- 1 Alhammarret's Archive
- 1 Angel of Vitality
- 1 Angelic Accord
- 1 Blood Artist
- 1 Bloodthirsty Aerialist
- 1 Cleric Class
- 1 Cleric of Life's Bond
- 1 Cradle of Vitality
- 1 Griffin Aerie
- 1 Heliod, Sun-Crowned
- 1 Leyline of Hope
- 1 Resplendent Angel
- 1 Righteous Valkyrie
- 1 Vito, Thorn of the Dusk Rose
- 1 Voice of the Blessed
- 1 Archangel of Thune
- 1 Astarion, the Decadent
- 1 Attended Healer
- 1 Blood Baron of Vizkopa
- 1 Celestine, the Living Saint
- 1 Cliffhaven Vampire
- 1 Defiant Bloodlord
- 1 Divinity of Pride
- 1 Enduring Tenacity
- 1 Gideon's Company
- 1 Nykthos Paragon
- 1 Valkyrie Harbinger
- 1 Ajani, Strength of the Pride
- 1 Fumigate

### 3. Bracket 1, Adeline, Resplendent Cathar, tokens

Pool 296. 99 cards, 0 block findings, repaired true (copy_limit).

Judge: error, judge bracket: llm anthropic/claude-opus-5: terminal (http 400): POST "https://api.anthropic.com/v1/messages": 400 Bad Request (Request-ID: req_011CeevGauaG3QCPxpH2mLjk) (Workspace-ID: wrkspc_01DKn3qmnB3ZMerq6QsCdj6C) {"type":"error","error":{"type":"invalid_request_error","message":"output_config.format.schema: For 'integer' type, properties maximum, minimum are not supported"},"request_id":"req_011CeevGauaG3QCPxpH2mLjk"}

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 39 | 36 to 41 |  |  |
| `tapped_land` | 3 | 0 to 14 |  | Path of Ancestry, Windbrisk Heights, Witch Enchanter // Witch-Blessed Meadow |
| `colorless_land` | 0 | 0 to 4 |  |  |
| `color_sources` | 1.64 | 0.75 or more |  | sources of requirement: W 42.75 of 26 |
| `avg_mana_value` | 3.48 | 2.4 to 4.4 |  | over 60 nonland cards |
| `ramp` | 9 | 6 to 12 |  |  |
| `draw` | 10 | 6 to 14 |  |  |
| `removal` | 7 | 4 to 10 |  |  |
| `wipe` | 2 | 0 to 4 |  |  |
| `interaction` | 4 | 0 to 8 |  |  |
| `tutor` | 0 | 0 to 1 |  |  |
| `fast_mana` | 0 | 0 to 1 |  |  |
| `game_changer` | 0 | 0 to 0 |  |  |
| `mana_turn_four` | 3.94 | 3.5 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.76 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | 0.09 | -9 to 1.5 |  | the commander costs 3 and comes down on turn 3.09 on average |

Goldfish over 10000 hands: the commander on turn 3.09, 3.94 mana on turn four, and 0.76 of first hands hold two to four lands.

Content: Spellbook tag E. Game Changers none. Mass land denial none. Extra turns none. Combos 0.

Findings:

- WARN `land_count`: 39 lands: the guide range for this format is 34 to 38
- INFO `curve_summary`: average mana value 3.48 over 60 nonland cards

Cards:

- 25 Plains
- 1 Ancient Den
- 1 Castle Ardenvale
- 1 Command Tower
- 1 Eiganjo, Seat of the Empire
- 1 Kjeldoran Outpost
- 1 Legion's Landing // Adanto, the First Fort
- 1 Minas Tirith
- 1 Mirrex
- 1 Path of Ancestry
- 1 Plaza of Heroes
- 1 Secluded Courtyard
- 1 Windbrisk Heights
- 1 Witch Enchanter // Witch-Blessed Meadow
- 1 Charisma Bobblehead
- 1 Coin of Mastery
- 1 Collector's Vault
- 1 Currency Converter
- 1 Druidic Satchel
- 1 Goldvein Pick
- 1 Keeper of the Accord
- 1 Monologue Tax
- 1 Nuka-Cola Vending Machine
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
- 1 Blessed Sanctuary
- 1 Rootborn Defenses
- 1 Spirit Bonds
- 1 Summon: Knights of Round
- 1 Banishing Slash
- 1 Generous Gift
- 1 Kellan's Lightblades
- 1 Righteous Confluence
- 1 Skyclave Apparition
- 1 Stroke of Midnight
- 1 The Wandering Emperor
- 1 Hour of Reckoning
- 1 Martial Coup
- 1 Aligned Heart
- 1 Anointer Priest
- 1 Cathar's Call
- 1 Clarion Spirit
- 1 Divine Visitation
- 1 Felidar Retreat
- 1 Halo Fountain
- 1 Horn of Gondor
- 1 Intangible Virtue
- 1 Luminarch Ascension
- 1 Mavren Fein, Dusk Apostle
- 1 Mondrak, Glory Dominus
- 1 Oketra's Monument
- 1 Retreat to Emeria
- 1 Rosie Cotton of South Lane
- 1 Siege Veteran
- 1 Skrelv's Hive
- 1 Archangel Elspeth
- 1 Ancient Gold Dragon
- 1 Basri's Lieutenant
- 1 Cemetery Protector
- 1 Emeria Angel
- 1 God-Eternal Oketra
- 1 Hero of Bladehold
- 1 Myrel, Shield of Argive
- 1 Ojer Taq, Deepest Foundation // Temple of Civilization
- 1 Requiem Angel
- 1 Sanctuary Warden
- 1 Threefold Thunderhulk

### 4. Bracket 2, Gishath, Sun's Avatar, dinosaurs

Pool 269. 99 cards, 0 block findings, repaired false.

Judge: error, judge bracket: llm anthropic/claude-opus-5: terminal (http 400): POST "https://api.anthropic.com/v1/messages": 400 Bad Request (Request-ID: req_011CeevKaRddYWP3cXYjzMFs) (Workspace-ID: wrkspc_01DKn3qmnB3ZMerq6QsCdj6C) {"type":"error","error":{"type":"invalid_request_error","message":"output_config.format.schema: For 'integer' type, properties maximum, minimum are not supported"},"request_id":"req_011CeevKaRddYWP3cXYjzMFs"}

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 37 | 35 to 40 |  |  |
| `tapped_land` | 2 | 0 to 12 |  | Jetmir's Garden, Path of Ancestry |
| `colorless_land` | 0 | 0 to 4 |  |  |
| `color_sources` | 0.82 | 0.8 or more |  | sources of requirement: W 18.5 of 20, R 15.5 of 19, G 23.25 of 23 |
| `avg_mana_value` | 3.77 | 2.3 to 3.9 |  | over 62 nonland cards |
| `ramp` | 10 | 8 to 12 |  |  |
| `draw` | 11 | 8 to 14 |  |  |
| `removal` | 7 | 5 to 10 |  |  |
| `wipe` | 3 | 1 to 5 |  |  |
| `interaction` | 6 | 2 to 10 |  |  |
| `tutor` | 0 | 0 to 2 |  |  |
| `fast_mana` | 0 | 0 to 2 |  |  |
| `game_changer` | 0 | 0 to 0 |  |  |
| `mana_turn_four` | 4.53 | 3.8 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.75 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | 0.69 | -9 to 1.2 |  | the commander costs 8 and comes down on turn 8.69 on average |

Goldfish over 10000 hands: the commander on turn 8.69, 4.53 mana on turn four, and 0.75 of first hands hold two to four lands.

Content: Spellbook tag E. Game Changers none. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 3.77 over 62 nonland cards

Cards:

- 12 Forest
- 7 Mountain
- 7 Plains
- 1 Battlefield Forge
- 1 Bountiful Promenade
- 1 Brushland
- 1 Canopy Vista
- 1 Cinder Glade
- 1 Clifftop Retreat
- 1 Command Tower
- 1 Exotic Orchard
- 1 Jetmir's Garden
- 1 Path of Ancestry
- 1 Temple Garden
- 1 Beast Whisperer
- 1 Cloudpiercer
- 1 Curious Altisaur
- 1 Earthshaker Dreadmaw
- 1 Garruk's Uprising
- 1 Guardian Project
- 1 Kutzil, Malamet Exemplar
- 1 Ripjaw Raptor
- 1 Return of the Wildspeaker
- 1 Shamanic Revelation
- 1 Vanquisher's Banner
- 1 Akroma's Will
- 1 Boros Charm
- 1 Heroic Intervention
- 1 Lightning Greaves
- 1 Swiftfoot Boots
- 1 Temple Altisaur
- 1 Arcane Signet
- 1 Cultivate
- 1 Farseek
- 1 Drover of the Mighty
- 1 Hulking Raptor
- 1 Nature's Lore
- 1 Ranging Raptors
- 1 Rampant Growth
- 1 Thunderherd Migration
- 1 Topiary Stomper
- 1 Bronzebeak Foragers
- 1 Burning Sun's Avatar
- 1 Itzquinth, Firstborn of Gishath
- 1 Needletooth Raptor
- 1 Ravenous Sailback
- 1 Savage Stomp
- 1 Thrashing Brontodon
- 1 Belligerent Yearling
- 1 Commune with Dinosaurs
- 1 Dinosaur Stampede
- 1 Deathgorge Scavenger
- 1 Huatli's Raptor
- 1 Kinjalli's Caller
- 1 Marauding Raptor
- 1 Otepec Huntmaster
- 1 Raptor Companion
- 1 Regal Imperiosaur
- 1 Sunfrill Imitator
- 1 Territorial Hammerskull
- 1 Triceraton Commander
- 1 Carnage Tyrant
- 1 Cavern Stomper
- 1 Etali, Primal Storm
- 1 Ghalta and Mavren
- 1 Goring Ceratops
- 1 Harnessed Snubhorn
- 1 Majestic Heliopterus
- 1 Palani's Hatcher
- 1 Pantlaza, Sun-Favored
- 1 Quartzwood Crasher
- 1 Regisaur Alpha
- 1 Sun-Crested Pterodon
- 1 Austere Command
- 1 Blasphemous Act
- 1 Wakening Sun's Avatar

### 5. Bracket 2, Karlov of the Ghost Council, lifegain

Pool 303. 99 cards, 0 block findings, repaired false.

Judge: error, judge bracket: llm anthropic/claude-opus-5: terminal (http 400): POST "https://api.anthropic.com/v1/messages": 400 Bad Request (Request-ID: req_011CeevNT1yMswCWmv9aHdx7) (Workspace-ID: wrkspc_01DKn3qmnB3ZMerq6QsCdj6C) {"type":"error","error":{"type":"invalid_request_error","message":"output_config.format.schema: For 'integer' type, properties maximum, minimum are not supported"},"request_id":"req_011CeevNT1yMswCWmv9aHdx7"}

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 37 | 35 to 40 |  |  |
| `tapped_land` | 3 | 0 to 12 |  | Kabira Crossroads, Scoured Barrens, Temple of Silence |
| `colorless_land` | 1 | 0 to 4 |  | Radiant Fountain |
| `color_sources` | 0.95 | 0.8 or more |  | sources of requirement: W 26.75 of 23, B 24.75 of 26 |
| `avg_mana_value` | 3.45 | 2.3 to 3.9 |  | over 62 nonland cards |
| `ramp` | 10 | 8 to 12 |  |  |
| `draw` | 11 | 8 to 14 |  |  |
| `removal` | 7 | 5 to 10 |  |  |
| `wipe` | 3 | 1 to 5 |  |  |
| `interaction` | 6 | 2 to 10 |  |  |
| `tutor` | 0 | 0 to 2 |  |  |
| `fast_mana` | 0 | 0 to 2 |  |  |
| `game_changer` | 0 | 0 to 0 |  |  |
| `mana_turn_four` | 4.24 | 3.8 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.74 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | 0.03 | -9 to 1.2 |  | the commander costs 2 and comes down on turn 2.03 on average |

Goldfish over 10000 hands: the commander on turn 2.03, 4.24 mana on turn four, and 0.74 of first hands hold two to four lands.

Content: Spellbook tag E. Game Changers none. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 3.45 over 62 nonland cards

Cards:

- 1 Archivist of Oghma
- 1 Dawn of Hope
- 1 Enduring Innocence
- 1 Exemplar of Light
- 1 Mangara, the Diplomat
- 1 Markov Purifier
- 1 Sigarda's Splendor
- 1 The Gaffer
- 1 Vampiric Rites
- 1 Well of Lost Dreams
- 1 Cosmos Elixir
- 1 Alseid of Life's Bounty
- 1 Faith's Shield
- 1 Metropolis Reformer
- 1 Restoration Magic
- 1 Werefox Bodyguard
- 1 Sword of Light and Shadow
- 14 Plains
- 13 Swamp
- 1 Caves of Koilos
- 1 Command Tower
- 1 Godless Shrine
- 1 Isolated Chapel
- 1 Kabira Crossroads
- 1 Radiant Fountain
- 1 Scoured Barrens
- 1 Shattered Sanctum
- 1 Temple of Silence
- 1 Vault of Champions
- 1 Altar of the Pantheon
- 1 Angel of Indemnity
- 1 Bounty Board
- 1 Colossal Plow
- 1 Crypt Ghast
- 1 Hierophant's Chalice
- 1 Nuka-Cola Vending Machine
- 1 Orazca Relic
- 1 Potioner's Trove
- 1 Pristine Talisman
- 1 Aetherflux Reservoir
- 1 Ayli, Eternal Pilgrim
- 1 Murderous Rider // Swift End
- 1 Noxious Gearhulk
- 1 Solitude
- 1 Vein Ripper
- 1 Vona, Butcher of Magan
- 1 Aerith Gainsborough
- 1 Ajani's Pridemate
- 1 Angel of Vitality
- 1 Bloodthirsty Aerialist
- 1 Cleric Class
- 1 Cleric of Life's Bond
- 1 Essence Channeler
- 1 Griffin Aerie
- 1 Heliod, Sun-Crowned
- 1 Indulging Patrician
- 1 Marauding Blight-Priest
- 1 Resplendent Angel
- 1 Voice of the Blessed
- 1 Archangel of Thune
- 1 Attended Healer
- 1 Blood Baron of Vizkopa
- 1 Bloodbond Vampire
- 1 Celestine, the Living Saint
- 1 Cliffhaven Vampire
- 1 Defiant Bloodlord
- 1 Divinity of Pride
- 1 Enduring Tenacity
- 1 Nykthos Paragon
- 1 Rhox Faithmender
- 1 Valkyrie Harbinger
- 1 Ajani, Strength of the Pride
- 1 Fumigate
- 1 Kaya's Wrath

### 6. Bracket 2, Adeline, Resplendent Cathar, tokens

Pool 296. 99 cards, 0 block findings, repaired false.

Judge: error, judge bracket: llm anthropic/claude-opus-5: terminal (http 400): POST "https://api.anthropic.com/v1/messages": 400 Bad Request (Request-ID: req_011CeevQtcQR69HyuNoegmor) (Workspace-ID: wrkspc_01DKn3qmnB3ZMerq6QsCdj6C) {"type":"error","error":{"type":"invalid_request_error","message":"output_config.format.schema: For 'integer' type, properties maximum, minimum are not supported"},"request_id":"req_011CeevQtcQR69HyuNoegmor"}

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 37 | 35 to 40 |  |  |
| `tapped_land` | 0 | 0 to 12 |  |  |
| `colorless_land` | 0 | 0 to 4 |  |  |
| `color_sources` | 1.54 | 0.8 or more |  | sources of requirement: W 40 of 26 |
| `avg_mana_value` | 3.53 | 2.3 to 3.9 |  | over 62 nonland cards |
| `ramp` | 10 | 8 to 12 |  |  |
| `draw` | 11 | 8 to 14 |  |  |
| `removal` | 7 | 5 to 10 |  |  |
| `wipe` | 3 | 1 to 5 |  |  |
| `interaction` | 6 | 2 to 10 |  |  |
| `tutor` | 0 | 0 to 2 |  |  |
| `fast_mana` | 0 | 0 to 2 |  |  |
| `game_changer` | 0 | 0 to 0 |  |  |
| `mana_turn_four` | 3.89 | 3.8 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.75 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | 0.1 | -9 to 1.2 |  | the commander costs 3 and comes down on turn 3.1 on average |

Goldfish over 10000 hands: the commander on turn 3.1, 3.89 mana on turn four, and 0.75 of first hands hold two to four lands.

Content: Spellbook tag E. Game Changers none. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 3.53 over 62 nonland cards

Cards:

- 37 Plains
- 1 Bennie Bracks, Zoologist
- 1 Bygone Bishop
- 1 Caretaker's Talent
- 1 Chivalric Alliance
- 1 Court of Grace
- 1 Dawn of Hope
- 1 Idol of Oblivion
- 1 Search the Premises
- 1 Staff of the Storyteller
- 1 Thorough Investigation
- 1 Wedding Announcement // Wedding Festivity
- 1 Blessed Sanctuary
- 1 Elspeth, Knight-Errant
- 1 Lena, Selfless Champion
- 1 Rootborn Defenses
- 1 Spirit Bonds
- 1 Teyo, the Shieldmage
- 1 Battle Angels of Tyr
- 1 Charisma Bobblehead
- 1 Coin of Mastery
- 1 Collector's Vault
- 1 Druidic Satchel
- 1 Karn, Living Legacy
- 1 Keeper of the Accord
- 1 Monologue Tax
- 1 Nuka-Cola Vending Machine
- 1 Smuggler's Share
- 1 Aerial Assault
- 1 Banishing Slash
- 1 Generous Gift
- 1 Kellan's Lightblades
- 1 Righteous Confluence
- 1 Stroke of Midnight
- 1 The Wandering Emperor
- 1 Anointer Priest
- 1 Clarion Spirit
- 1 Divine Visitation
- 1 Felidar Retreat
- 1 Horn of Gondor
- 1 Intangible Virtue
- 1 Luminarch Ascension
- 1 Oketra's Monument
- 1 Retreat to Emeria
- 1 Rosie Cotton of South Lane
- 1 Skrelv's Hive
- 1 Song of the Worldsoul
- 1 Twilight Drover
- 1 Ajani's Chosen
- 1 Archangel Elspeth
- 1 Archon of Sun's Grace
- 1 Attended Healer
- 1 Basri's Lieutenant
- 1 Emeria Angel
- 1 God-Eternal Oketra
- 1 Hero of Bladehold
- 1 Myrel, Shield of Argive
- 1 Oketra the True
- 1 Requiem Angel
- 1 Sanctuary Warden
- 1 Elspeth, Sun's Champion
- 1 Hour of Reckoning
- 1 Martial Coup

### 7. Bracket 3, Karlov of the Ghost Council, lifegain

Pool 303. 99 cards, 0 block findings, repaired true (profile_off_band).

Judge: error, judge bracket: llm anthropic/claude-opus-5: terminal (http 400): POST "https://api.anthropic.com/v1/messages": 400 Bad Request (Request-ID: req_011CeevYpWKm8CR27URWokHS) (Workspace-ID: wrkspc_01DKn3qmnB3ZMerq6QsCdj6C) {"type":"error","error":{"type":"invalid_request_error","message":"output_config.format.schema: For 'integer' type, properties maximum, minimum are not supported"},"request_id":"req_011CeevYpWKm8CR27URWokHS"}

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 36 | 34 to 38 |  |  |
| `tapped_land` | 3 | 0 to 9 |  | Restless Fortress, Scoured Barrens, Temple of Silence |
| `colorless_land` | 0 | 0 to 4 |  |  |
| `color_sources` | 1.04 | 0.85 or more |  | sources of requirement: W 29.25 of 23, B 29.25 of 28 |
| `avg_mana_value` | 3.24 | 2 to 3.5 |  | over 63 nonland cards |
| `ramp` | 10 | 8 to 13 |  |  |
| `draw` | 11 | 8 to 14 |  |  |
| `removal` | 9 | 6 to 12 |  |  |
| `wipe` | 3 | 2 to 5 |  |  |
| `interaction` | 8 | 4 to 12 |  |  |
| `tutor` | 0 | 0 to 5 |  |  |
| `fast_mana` | 0 | 0 to 3 |  |  |
| `game_changer` | 0 | 0 to 3 |  |  |
| `mana_turn_four` | 4.24 | 4.2 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.74 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | 0.03 | -9 to 0.8 |  | the commander costs 2 and comes down on turn 2.03 on average |

Goldfish over 10000 hands: the commander on turn 2.03, 4.24 mana on turn four, and 0.74 of first hands hold two to four lands.

Content: Spellbook tag E. Game Changers none. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 3.24 over 63 nonland cards

Cards:

- 1 Godless Shrine
- 1 Caves of Koilos
- 1 Command Tower
- 1 City of Brass
- 1 Mana Confluence
- 1 Arguel's Blood Fast // Temple of Aclazotz
- 1 Aclazotz, Deepest Betrayal // Temple of the Dead
- 1 Isolated Chapel
- 1 Shattered Sanctum
- 1 Vault of Champions
- 1 Scoured Barrens
- 1 Temple of Silence
- 1 Neglected Manor
- 1 Restless Fortress
- 12 Plains
- 10 Swamp
- 1 Altar of the Pantheon
- 1 Bounty Board
- 1 Colossal Plow
- 1 Crypt Ghast
- 1 Hierophant's Chalice
- 1 Hot Dog Cart
- 1 Nuka-Cola Vending Machine
- 1 Potioner's Trove
- 1 Pristine Talisman
- 1 The Celestus
- 1 Archivist of Oghma
- 1 Dawn of Hope
- 1 Cosmos Elixir
- 1 Enduring Innocence
- 1 Mangara, the Diplomat
- 1 Markov Purifier
- 1 Sigarda's Splendor
- 1 The Gaffer
- 1 Tymna the Weaver
- 1 Well of Lost Dreams
- 1 Vampiric Rites
- 1 Alseid of Life's Bounty
- 1 Bofur, Reliable Guardian // Concerted Care
- 1 Caduceus, Staff of Hermes
- 1 Faith's Shield
- 1 Metropolis Reformer
- 1 Restoration Magic
- 1 Sword of Light and Shadow
- 1 Werefox Bodyguard
- 1 Ayli, Eternal Pilgrim
- 1 Aetherflux Reservoir
- 1 Foolish Fate
- 1 Gumdrop Poisoner // Tempt with Treats
- 1 Murderous Rider // Swift End
- 1 Nightmare's Thirst
- 1 Solitude
- 1 Sorin of House Markov // Sorin, Ravenous Neonate
- 1 Vein Ripper
- 1 Fumigate
- 1 Kaya's Wrath
- 1 The Meathook Massacre
- 1 Ajani's Pridemate
- 1 Angel of Vitality
- 1 Bloodthirsty Aerialist
- 1 Cleric Class
- 1 Heliod, Sun-Crowned
- 1 Resplendent Angel
- 1 Righteous Valkyrie
- 1 Serra Ascendant
- 1 Vito, Thorn of the Dusk Rose
- 1 Voice of the Blessed
- 1 Archangel of Thune
- 1 Angel of Destiny
- 1 Attended Healer
- 1 Celestine, the Living Saint
- 1 Cliffhaven Vampire
- 1 Defiant Bloodlord
- 1 Divinity of Pride
- 1 Nykthos Paragon
- 1 Rhox Faithmender
- 1 Twinblade Paladin
- 1 Valkyrie Harbinger
- 1 Wurmcoil Engine

### 8. Bracket 3, Denethor, Ruling Steward, aristocrats sacrifice

Pool 295. 99 cards, 0 block findings, repaired true (profile_off_band).

Judge: error, judge bracket: llm anthropic/claude-opus-5: terminal (http 400): POST "https://api.anthropic.com/v1/messages": 400 Bad Request (Request-ID: req_011CeevmXrL49PSu6vDe7uQg) (Workspace-ID: wrkspc_01DKn3qmnB3ZMerq6QsCdj6C) {"type":"error","error":{"type":"invalid_request_error","message":"output_config.format.schema: For 'integer' type, properties maximum, minimum are not supported"},"request_id":"req_011CeevmXrL49PSu6vDe7uQg"}

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 36 | 34 to 38 |  |  |
| `tapped_land` | 0 | 0 to 9 |  |  |
| `colorless_land` | 1 | 0 to 4 |  | High Market |
| `color_sources` | 1.11 | 0.85 or more |  | sources of requirement: W 21 of 19, B 30.25 of 19 |
| `avg_mana_value` | 2.95 | 2 to 3.5 |  | over 63 nonland cards |
| `ramp` | 10 | 8 to 13 |  |  |
| `draw` | 11 | 8 to 14 |  |  |
| `removal` | 9 | 6 to 12 |  |  |
| `wipe` | 3 | 2 to 5 |  |  |
| `interaction` | 8 | 4 to 12 |  |  |
| `tutor` | 0 | 0 to 5 |  |  |
| `fast_mana` | 2 | 0 to 3 |  | Culling the Weak, Sacrifice |
| `game_changer` | 0 | 0 to 3 |  |  |
| `mana_turn_four` | 4.17 | 4.2 or more | yes | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.74 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.06 | -9 to 0.8 |  | the commander costs 3 and comes down on turn 2.94 on average |

Goldfish over 10000 hands: the commander on turn 2.94, 4.17 mana on turn four, and 0.74 of first hands hold two to four lands.

Content: Spellbook tag E. Game Changers none. Mass land denial none. Extra turns none. Combos 1.
- Elenda, the Dusk Rose + Felisa, Fang of Silverquill + Phyrexian Altar (speed 5)

Findings:

- INFO `curve_summary`: average mana value 2.95 over 63 nonland cards
- WARN `profile_off_band`: the mana available on turn four is 4.17, and bracket 3 wants 4.2 or more (mean over 10000 hands)

Cards:

- 1 Command Tower
- 1 Godless Shrine
- 1 Scrubland
- 1 Caves of Koilos
- 1 City of Brass
- 1 Mana Confluence
- 1 Vault of Champions
- 1 Exotic Orchard
- 1 Phyrexian Tower
- 1 High Market
- 16 Swamp
- 10 Plains
- 1 Ashnod's Altar
- 1 Ashnod, Flesh Mechanist
- 1 Culling the Weak
- 1 Deadly Dispute
- 1 Lotus Ring
- 1 Phyrexian Altar
- 1 Priest of Forgotten Gods
- 1 Sacrifice
- 1 Transmogrant Altar
- 1 Warren Soultrader
- 1 Baron Bertram Graywater
- 1 Bushmeat Poacher
- 1 Corrupted Conviction
- 1 Disciple of Bolas
- 1 Grave Venerations
- 1 Relic Vial
- 1 Shadowheart, Dark Justiciar
- 1 Smothering Abomination
- 1 Tevesh Szat, Doom of Fools
- 1 Vampire Gourmand
- 1 Village Rites
- 1 Cartel Aristocrat
- 1 Dark Privilege
- 1 Fanatical Devotion
- 1 Flare of Fortitude
- 1 Gift of Doom
- 1 Nightmare Shepherd
- 1 Promise of Tomorrow
- 1 Spirit Bonds
- 1 Attrition
- 1 Ayli, Eternal Pilgrim
- 1 Blasting Station
- 1 Bone Shards
- 1 Eaten Alive
- 1 Grave Pact
- 1 Ruthless Lawbringer
- 1 Teysa, Orzhov Scion
- 1 Yawgmoth, Thran Physician
- 1 Bastion of Remembrance
- 1 Blood Artist
- 1 Carrion Feeder
- 1 Cruel Celebrant
- 1 Elas il-Kor, Sadistic Pilgrim
- 1 Fleshtaker
- 1 Hidden Stockpile
- 1 Viscera Seer
- 1 Woe Strider
- 1 Zulaport Cutthroat
- 1 Basri's Lieutenant
- 1 Blood Host
- 1 Elenda, the Dusk Rose
- 1 Falkenrath Noble
- 1 Felisa, Fang of Silverquill
- 1 Flesh-Eater Imp
- 1 Ghoulcaller Gisa
- 1 Liliana, Heretical Healer // Liliana, Defiant Necromancer
- 1 Ratadrabik of Urborg
- 1 Syr Konrad, the Grim
- 1 Vampire Warlord
- 1 Vindictive Vampire
- 1 Minister of Pain
- 1 The Meathook Massacre
- 1 Toxic Deluge

### 9. Bracket 3, Zada, Hedron Grinder, goblins

Pool 294. 99 cards, 0 block findings, repaired false.

Judge: error, judge bracket: llm anthropic/claude-opus-5: terminal (http 400): POST "https://api.anthropic.com/v1/messages": 400 Bad Request (Request-ID: req_011CeevpZAHZ8AWnyJRSs3aP) (Workspace-ID: wrkspc_01DKn3qmnB3ZMerq6QsCdj6C) {"type":"error","error":{"type":"invalid_request_error","message":"output_config.format.schema: For 'integer' type, properties maximum, minimum are not supported"},"request_id":"req_011CeevpZAHZ8AWnyJRSs3aP"}

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 37 | 34 to 38 |  |  |
| `tapped_land` | 4 | 0 to 9 |  | Path of Ancestry, Shatterskull Smashing // Shatterskull, the Hammer Pass, Spinerock Knoll, Valakut Awakening // Valakut Stoneforge |
| `colorless_land` | 0 | 0 to 4 |  |  |
| `color_sources` | 1.6 | 0.85 or more |  | sources of requirement: R 41.5 of 26 |
| `avg_mana_value` | 2.77 | 2 to 3.5 |  | over 62 nonland cards |
| `ramp` | 10 | 8 to 13 |  |  |
| `draw` | 11 | 8 to 14 |  |  |
| `removal` | 9 | 6 to 12 |  |  |
| `wipe` | 2 | 2 to 5 |  |  |
| `interaction` | 8 | 4 to 12 |  |  |
| `tutor` | 1 | 0 to 5 |  | Moggcatcher |
| `fast_mana` | 2 | 0 to 3 |  | Brightstone Ritual, Sol Ring |
| `game_changer` | 2 | 0 to 3 |  | Jeska's Will, The One Ring |
| `mana_turn_four` | 4.58 | 4.2 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.75 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.37 | -9 to 0.8 |  | the commander costs 4 and comes down on turn 3.63 on average |

Goldfish over 10000 hands: the commander on turn 3.63, 4.58 mana on turn four, and 0.75 of first hands hold two to four lands.

Content: Spellbook tag S. Game Changers Jeska's Will, The One Ring. Mass land denial none. Extra turns none. Combos 3.
- Krenko, Mob Boss + Skirk Prospector + Goblin Warchief (speed 5)
- Krenko, Mob Boss + Skirk Prospector + Goblin Chieftain (speed 5)
- Kiki-Jiki, Mirror Breaker + Akki Battle Squad (speed 5)

Findings:

- INFO `curve_summary`: average mana value 2.77 over 62 nonland cards
- INFO `cards_trimmed`: the list was 1 card over, so the builder cut this card: Blasphemous Act

Cards:

- 24 Mountain
- 1 Arena of Glory
- 1 Castle Embereth
- 1 Cavern of Souls
- 1 Command Tower
- 1 Great Furnace
- 1 Path of Ancestry
- 1 Secluded Courtyard
- 1 Shatterskull Smashing // Shatterskull, the Hammer Pass
- 1 Sokenzan, Crucible of Defiance
- 1 Spinerock Knoll
- 1 Three Tree City
- 1 Unclaimed Territory
- 1 Valakut Awakening // Valakut Stoneforge
- 1 Skullclamp
- 1 Wheel of Fortune
- 1 Faithless Looting
- 1 Demand Answers
- 1 Thrill of Possibility
- 1 Cathartic Reunion
- 1 Mask of Memory
- 1 Rummaging Goblin
- 1 Sensation Gorger
- 1 Goblin Picker
- 1 Fissure Wizard
- 1 Sol Ring
- 1 Arcane Signet
- 1 Fellwar Stone
- 1 Mind Stone
- 1 Thought Vessel
- 1 Wayfarer's Bauble
- 1 Skirk Prospector
- 1 Brightstone Ritual
- 1 Jeska's Will
- 1 Crime Novelist
- 1 Goblin Grenade
- 1 Goblin Cratermaker
- 1 Goblin Sharpshooter
- 1 Pashalik Mons
- 1 Siege-Gang Commander
- 1 Sparksmith
- 1 Gempalm Incinerator
- 1 Broadside Bombardiers
- 1 Goblin Trashmaster
- 1 Chain Reaction
- 1 Skirk Fire Marshal
- 1 Lightning Greaves
- 1 Swiftfoot Boots
- 1 Mithril Coat
- 1 Pyroblast
- 1 Red Elemental Blast
- 1 Tibalt's Trickery
- 1 Goblin Chirurgeon
- 1 The One Ring
- 1 Goblin Chieftain
- 1 Goblin Warchief
- 1 Conspicuous Snoop
- 1 Goblin Lackey
- 1 Goblin Piledriver
- 1 Mogg War Marshal
- 1 Rundvelt Hordemaster
- 1 Krenko, Tin Street Kingpin
- 1 Goblin Rabblemaster
- 1 Battle Cry Goblin
- 1 Krenko, Mob Boss
- 1 Muxus, Goblin Grandee
- 1 Kiki-Jiki, Mirror Breaker
- 1 Akki Battle Squad
- 1 Battle Squadron
- 1 Beetleback Chief
- 1 Goblin Ringleader
- 1 Goblin Marshal
- 1 Reckless One
- 1 Grenzo's Ruffians
- 1 Moggcatcher
- 1 Swarming Goblins

### 10. Bracket 4, Urza, Lord High Artificer, artifacts

Pool 299. 99 cards, 0 block findings, repaired true (profile_off_band, profile_off_band).

Judge: error, judge bracket: llm anthropic/claude-opus-5: terminal (http 400): POST "https://api.anthropic.com/v1/messages": 400 Bad Request (Request-ID: req_011Ceew7QNBLsNmfNpNehDjy) (Workspace-ID: wrkspc_01DKn3qmnB3ZMerq6QsCdj6C) {"type":"error","error":{"type":"invalid_request_error","message":"output_config.format.schema: For 'integer' type, properties maximum, minimum are not supported"},"request_id":"req_011Ceew7QNBLsNmfNpNehDjy"}

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 33 | 31 to 36 |  |  |
| `tapped_land` | 2 | 0 to 5 |  | Power Depot, Sink into Stupor // Soporific Springs |
| `colorless_land` | 3 | 0 to 5 |  | Mishra's Workshop, Urza's Saga, Urza's Workshop |
| `color_sources` | 1.27 | 0.9 or more |  | sources of requirement: U 33 of 26 |
| `avg_mana_value` | 3.11 | 1.6 to 3 | yes | over 66 nonland cards |
| `ramp` | 13 | 10 to 16 |  |  |
| `draw` | 12 | 8 to 16 |  |  |
| `removal` | 10 | 6 to 14 |  |  |
| `wipe` | 2 | 0 to 5 |  |  |
| `interaction` | 11 | 6 to 16 |  |  |
| `tutor` | 0 | none |  |  |
| `fast_mana` | 2 | none |  | Moonsnare Prototype, Mox Opal |
| `game_changer` | 1 | none |  | Mishra's Workshop |
| `mana_turn_four` | 4.24 | 4.8 or more | yes | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.71 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | 0.25 | -9 to 0.5 |  | the commander costs 4 and comes down on turn 4.25 on average |

Goldfish over 10000 hands: the commander on turn 4.25, 4.24 mana on turn four, and 0.71 of first hands hold two to four lands.

Content: Spellbook tag P. Game Changers Mishra's Workshop. Mass land denial none. Extra turns none. Combos 0.

Findings:

- WARN `land_count`: 33 lands: the guide range for this format is 34 to 38
- INFO `curve_summary`: average mana value 3.11 over 66 nonland cards
- WARN `profile_off_band`: the average mana value of the nonland cards is 3.11, and bracket 4 wants 1.6 to 3
- WARN `profile_off_band`: the mana available on turn four is 4.24, and bracket 4 wants 4.8 or more (mean over 10000 hands)

Cards:

- 12 Island
- 1 Command Tower
- 1 Exotic Orchard
- 1 Glimmervoid
- 1 Mana Confluence
- 1 Reflecting Pool
- 1 Seat of the Synod
- 1 Spire of Industry
- 1 Mystic Sanctuary
- 1 Otawara, Soaring City
- 1 Mistrise Village
- 1 Sink into Stupor // Soporific Springs
- 1 Cavern of Souls
- 1 Gemstone Caverns
- 1 Plaza of Heroes
- 1 City of Brass
- 1 Secluded Courtyard
- 1 Mishra's Workshop
- 1 Urza's Workshop
- 1 Nykthos, Shrine to Nyx
- 1 Urza's Saga
- 1 Power Depot
- 1 Mox Opal
- 1 Moonsnare Prototype
- 1 Arc Reactor
- 1 Coin of Mastery
- 1 Crystal Skull, Isu Spyglass
- 1 Metalworker
- 1 Vedalken Engineer
- 1 Chief Engineer
- 1 Grand Architect
- 1 Inspiring Statuary
- 1 Krark-Clan Ironworks
- 1 Slagstone Refinery
- 1 Powerstone Shard
- 1 Thoughtcast
- 1 Thirst for Knowledge
- 1 Riddlesmith
- 1 Sai, Master Thopterist
- 1 Forensic Gadgeteer
- 1 Era of Innovation
- 1 Esoteric Duplicator
- 1 Reverse Engineer
- 1 Thought Monitor
- 1 Vedalken Archmage
- 1 Cerebral Download
- 1 Waterbending Lesson
- 1 Disruption Protocol
- 1 Metallic Rebuke
- 1 Ice Out
- 1 Stoic Rebuttal
- 1 Reality Ripple
- 1 Ghostly Flicker
- 1 Welding Jar
- 1 Curator's Ward
- 1 Escape Protocol
- 1 Etched Champion
- 1 Syr Ginger, the Meal Ender
- 1 Aether Spellbomb
- 1 Contagion Clasp
- 1 Cyber Conversion
- 1 Resculpt
- 1 Ravenform
- 1 Into Thin Air
- 1 Unable to Scream
- 1 Water Whip
- 1 Watery Grasp
- 1 Invasion Submersible
- 1 Engineered Explosives
- 1 Hurkyl's Recall
- 1 Voltaic Key
- 1 Manifold Key
- 1 Etherium Sculptor
- 1 Emry, Lurker of the Loch
- 1 Foundry Inspector
- 1 Mystic Forge
- 1 Panther Robot
- 1 Mechan Assembler
- 1 Memory Guardian
- 1 Chrome Steed
- 1 Frogmite
- 1 Master Transmuter
- 1 Phyrexian Metamorph
- 1 Traxos, Scourge of Kroog
- 1 Whirler Rogue
- 1 Karn, Scion of Urza
- 1 Jhoira's Familiar
- 1 Lodestone Golem

### 11. Bracket 4, Korvold, Fae-Cursed King, treasure sacrifice

Pool 300. 99 cards, 0 block findings, repaired true (profile_off_band, profile_off_band, profile_off_band).

Judge: error, judge bracket: llm anthropic/claude-opus-5: terminal (http 400): POST "https://api.anthropic.com/v1/messages": 400 Bad Request (Request-ID: req_011CeewPbTckjpoP9qv17ev2) (Workspace-ID: wrkspc_01DKn3qmnB3ZMerq6QsCdj6C) {"type":"error","error":{"type":"invalid_request_error","message":"output_config.format.schema: For 'integer' type, properties maximum, minimum are not supported"},"request_id":"req_011CeewPbTckjpoP9qv17ev2"}

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 33 | 31 to 36 |  |  |
| `tapped_land` | 0 | 0 to 5 |  |  |
| `colorless_land` | 0 | 0 to 5 |  |  |
| `color_sources` | 0.89 | 0.9 or more | yes | sources of requirement: B 22.75 of 23, R 17 of 19, G 20 of 19 |
| `avg_mana_value` | 3.12 | 1.6 to 3 | yes | over 66 nonland cards |
| `ramp` | 13 | 10 to 16 |  |  |
| `draw` | 12 | 8 to 16 |  |  |
| `removal` | 10 | 6 to 14 |  |  |
| `wipe` | 2 | 0 to 5 |  |  |
| `interaction` | 11 | 6 to 16 |  |  |
| `tutor` | 4 | none |  | Diabolic Intent, Magda, Brazen Outlaw, Razaketh, the Foulblooded, Sidisi, Undead Vizier |
| `fast_mana` | 0 | none |  |  |
| `game_changer` | 0 | none |  |  |
| `mana_turn_four` | 3.96 | 4.8 or more | yes | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.71 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | 0.81 | -9 to 0.5 | yes | the commander costs 5 and comes down on turn 5.81 on average |

Goldfish over 10000 hands: the commander on turn 5.81, 3.96 mana on turn four, and 0.71 of first hands hold two to four lands.

Content: Spellbook tag E. Game Changers none. Mass land denial none. Extra turns none. Combos 4.
- Stimulus Package + Warren Soultrader + Relic Vial (speed 5)
- Stimulus Package + Pitiless Plunderer + Carrion Feeder (speed 5)
- Stimulus Package + Pitiless Plunderer + Viscera Seer (speed 5)
- Stimulus Package + Pitiless Plunderer + Goblin Bombardment (speed 5)

Findings:

- WARN `land_count`: 33 lands: the guide range for this format is 34 to 38
- INFO `curve_summary`: average mana value 3.12 over 66 nonland cards
- WARN `profile_off_band`: the worst color's share of the sources it needs is 0.89, and bracket 4 wants 0.9 or more (sources of requirement: B 22.75 of 23, R 17 of 19, G 20 of 19)
- WARN `profile_off_band`: the average mana value of the nonland cards is 3.12, and bracket 4 wants 1.6 to 3
- WARN `profile_off_band`: the mana available on turn four is 3.96, and bracket 4 wants 4.8 or more (mean over 10000 hands)
- WARN `profile_off_band`: the turns the commander comes down after its mana value is 0.81, and bracket 4 wants 0.5 at most (the commander costs 5 and comes down on turn 5.81 on average)

Cards:

- 1 Blood Crypt
- 1 Overgrown Tomb
- 1 Stomping Ground
- 1 Command Tower
- 1 City of Brass
- 1 Mana Confluence
- 1 Luxury Suite
- 1 Undergrowth Stadium
- 1 Llanowar Wastes
- 1 Sulfurous Springs
- 1 Cavern of Souls
- 1 Reflecting Pool
- 1 Phyrexian Tower
- 8 Forest
- 7 Swamp
- 5 Mountain
- 1 Forsworn Paladin
- 1 Gemcutter Buccaneer
- 1 Glóin, Dwarf Emissary
- 1 Magda, Brazen Outlaw
- 1 Magda, the Hoardmaster
- 1 Master of Dark Rites
- 1 Pitiless Plunderer
- 1 Professional Face-Breaker
- 1 Captain Lannery Storm
- 1 Jolene, Plundering Pugilist
- 1 Kellogg, Dangerous Mind
- 1 Warren Soultrader
- 1 Stimulus Package
- 1 Anje, Maid of Dishonor
- 1 Beetle-Headed Merchants
- 1 Bushmeat Poacher
- 1 Dina, Essence Brewer
- 1 Ecstatic Awakener // Awoken Demon
- 1 Gas Guzzler
- 1 High-Society Hunter
- 1 Lord Skitter's Butcher
- 1 Relic Vial
- 1 Smothering Abomination
- 1 Skyclave Shadowcat
- 1 Vampire Gourmand
- 1 Cultist of the Absolute
- 1 Dark Privilege
- 1 Gift of Doom
- 1 Goblin Chirurgeon
- 1 My Precious // Allure of Power
- 1 Nightmare Shepherd
- 1 Sylvan Safekeeper
- 1 Constant Mists
- 1 Kinzu of the Bleak Coven
- 1 Strefan, Maurer Progenitor
- 1 Campsite Cuisine
- 1 Attrition
- 1 Bone Shards
- 1 Eaten Alive
- 1 Goblin Bombardment
- 1 Blasting Station
- 1 Broadside Bombardiers
- 1 Grist, the Hunger Tide
- 1 Minsc & Boo, Timeless Heroes
- 1 Street Urchin
- 1 Yawgmoth, Thran Physician
- 1 Academy Manufactor
- 1 Xorn
- 1 Viscera Seer
- 1 Carrion Feeder
- 1 Witch's Oven
- 1 Diabolic Intent
- 1 Flesh-Eater Imp
- 1 Daemogoth Titan
- 1 Demonlord of Ashmouth
- 1 Desecration Elemental
- 1 Ghoulcaller Gisa
- 1 Jarad, Golgari Lich Lord
- 1 Kethek, Crucible Goliath
- 1 Meren of Clan Nel Toth
- 1 Prossh, Skyraider of Kher
- 1 Razaketh, the Foulblooded
- 1 Sidisi, Undead Vizier
- 1 Smaug the Impenetrable
- 1 The Meathook Massacre
- 1 Spontaneous Combustion

### 12. Bracket 4, Prosper, Tome-Bound, treasure

Pool 188. 99 cards, 0 block findings, repaired true (profile_off_band).

Judge: error, judge bracket: llm anthropic/claude-opus-5: terminal (http 400): POST "https://api.anthropic.com/v1/messages": 400 Bad Request (Request-ID: req_011CeewfwTLxU1kbpVf2GGJ6) (Workspace-ID: wrkspc_01DKn3qmnB3ZMerq6QsCdj6C) {"type":"error","error":{"type":"invalid_request_error","message":"output_config.format.schema: For 'integer' type, properties maximum, minimum are not supported"},"request_id":"req_011CeewfwTLxU1kbpVf2GGJ6"}

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 33 | 31 to 36 |  |  |
| `tapped_land` | 6 | 0 to 5 | yes | Canyon Slough, Foreboding Ruins, Path of Ancestry, Rakdos Carnarium, Raucous Theater, Temple of Malice |
| `colorless_land` | 1 | 0 to 5 |  | Dragon-Cursed Halls |
| `color_sources` | 2.13 | 0.9 or more |  | sources of requirement: B 40.5 of 19, R 40.5 of 19 |
| `avg_mana_value` | 2.83 | 1.6 to 3 |  | over 66 nonland cards |
| `ramp` | 13 | 10 to 16 |  |  |
| `draw` | 12 | 8 to 16 |  |  |
| `removal` | 10 | 6 to 14 |  |  |
| `wipe` | 2 | 0 to 5 |  |  |
| `interaction` | 11 | 6 to 16 |  |  |
| `tutor` | 1 | none |  | Magda, Brazen Outlaw |
| `fast_mana` | 0 | none |  |  |
| `game_changer` | 1 | none |  | The One Ring |
| `mana_turn_four` | 3.93 | 4.8 or more | yes | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.71 | 0.7 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | 0.48 | -9 to 0.5 |  | the commander costs 4 and comes down on turn 4.48 on average |

Goldfish over 10000 hands: the commander on turn 4.48, 3.93 mana on turn four, and 0.71 of first hands hold two to four lands.

Content: Spellbook tag P. Game Changers The One Ring. Mass land denial none. Extra turns none. Combos 1.
- Prosper, Tome-Bound + Professional Face-Breaker + Sensei's Divining Top + Xorn (speed 5)

Findings:

- WARN `land_count`: 33 lands: the guide range for this format is 34 to 38
- INFO `curve_summary`: average mana value 2.83 over 66 nonland cards
- WARN `profile_off_band`: the count of lands that enter tapped is 6, and bracket 4 wants 5 at most (Canyon Slough, Foreboding Ruins, Path of Ancestry, Rakdos Carnarium, Raucous Theater, Temple of Malice)
- WARN `profile_off_band`: the mana available on turn four is 3.93, and bracket 4 wants 4.8 or more (mean over 10000 hands)

Cards:

- 1 Badlands
- 1 Blazemire Verge
- 1 Blood Crypt
- 1 Cavern of Souls
- 1 City of Brass
- 1 Command Tower
- 1 Dragon-Cursed Halls
- 1 Dragonskull Summit
- 1 Exotic Orchard
- 1 Foreboding Ruins
- 1 Graven Cairns
- 1 Haunted Ridge
- 1 Luxury Suite
- 1 Mana Confluence
- 1 Opal Palace
- 1 Plaza of Heroes
- 1 Raucous Theater
- 1 Reflecting Pool
- 1 Secluded Courtyard
- 1 Shadowblood Ridge
- 1 Smoldering Marsh
- 1 Spire of Industry
- 1 Sulfurous Springs
- 1 Tainted Peak
- 1 Unclaimed Territory
- 1 Canyon Slough
- 1 Path of Ancestry
- 1 Rakdos Carnarium
- 1 Temple of Malice
- 2 Mountain
- 2 Swamp
- 1 Forsworn Paladin
- 1 Ragavan, Nimble Pilferer
- 1 Alchemist's Talent
- 1 Gemcutter Buccaneer
- 1 Magda, Brazen Outlaw
- 1 Magda, the Hoardmaster
- 1 Kalain, Reclusive Painter
- 1 Black Market Connections
- 1 Descent into Avernus
- 1 Skullport Merchant
- 1 Captain Lannery Storm
- 1 Professional Face-Breaker
- 1 Mastermind Plum
- 1 Faithless Looting
- 1 Night's Whisper
- 1 Sign in Blood
- 1 Demand Answers
- 1 Thrill of Possibility
- 1 Corrupted Conviction
- 1 Skullclamp
- 1 Idol of Oblivion
- 1 Sensei's Divining Top
- 1 Wheel of Fortune
- 1 Read the Bones
- 1 The One Ring
- 1 Pyroblast
- 1 Red Elemental Blast
- 1 Tibalt's Trickery
- 1 Vexing Bauble
- 1 Lightning Greaves
- 1 Swiftfoot Boots
- 1 Mithril Coat
- 1 Kaya's Ghostform
- 1 Not Dead After All
- 1 Undying Malice
- 1 Whispersilk Cloak
- 1 Abrade
- 1 Chaos Warp
- 1 Deadly Rollick
- 1 Go for the Throat
- 1 Infernal Grasp
- 1 Lightning Bolt
- 1 Snuff Out
- 1 Terminate
- 1 Withering Torment
- 1 Vandalblast
- 1 Academy Manufactor
- 1 Xorn
- 1 Contract Hero
- 1 Crossover Collaboration
- 1 Mine Raider
- 1 Gilded Assault Cart
- 1 Revel in Riches
- 1 Knuckles the Echidna
- 1 Luck Bobblehead
- 1 Goldspan Dragon
- 1 Ancient Copper Dragon
- 1 Flamekin Gildweaver
- 1 Jaded Sell-Sword
- 1 Casey & Raph, Hotheads
- 1 Meticulous Artisan
- 1 Marut
- 1 Smaug the Impenetrable
- 1 Torment of Hailfire
- 1 Toxic Deluge
- 1 Blasphemous Act

### 13. Bracket 5, Kinnan, Bonder Prodigy, combo

Pool 303. 99 cards, 0 block findings, repaired true (profile_off_band, profile_off_band).

Judge: error, judge bracket: llm anthropic/claude-opus-5: terminal (http 400): POST "https://api.anthropic.com/v1/messages": 400 Bad Request (Request-ID: req_011CeewzRfV1TBhLnZYmxoWN) (Workspace-ID: wrkspc_01DKn3qmnB3ZMerq6QsCdj6C) {"type":"error","error":{"type":"invalid_request_error","message":"output_config.format.schema: For 'integer' type, properties maximum, minimum are not supported"},"request_id":"req_011CeewzRfV1TBhLnZYmxoWN"}

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 30 | 27 to 33 |  |  |
| `tapped_land` | 0 | 0 to 2 |  |  |
| `colorless_land` | 0 | 0 to 6 |  |  |
| `color_sources` | 1.42 | 0.9 or more |  | sources of requirement: U 28.25 of 19, G 27 of 19 |
| `avg_mana_value` | 2.55 | 1.2 to 2.4 | yes | over 69 nonland cards |
| `ramp` | 17 | 12 to 22 |  |  |
| `draw` | 13 | 8 to 18 |  |  |
| `removal` | 9 | 4 to 14 |  |  |
| `wipe` | 1 | 0 to 3 |  |  |
| `interaction` | 15 | 8 to 22 |  |  |
| `tutor` | 4 | none |  | Fabricate, Muddle the Mixture, Transmutation Font, Worldly Tutor |
| `fast_mana` | 1 | none |  | Mana Vault |
| `game_changer` | 5 | none |  | Cyclonic Rift, Grim Monolith, Mana Vault, Thassa's Oracle, Worldly Tutor |
| `mana_turn_four` | 4.38 | 5.5 or more | yes | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.66 | 0.65 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | 0.01 | -9 to 0.3 |  | the commander costs 2 and comes down on turn 2.01 on average |

Goldfish over 10000 hands: the commander on turn 2.01, 4.38 mana on turn four, and 0.66 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Cyclonic Rift, Grim Monolith, Mana Vault, Thassa's Oracle, Worldly Tutor. Mass land denial none. Extra turns none. Combos 7.
- Ioreth of the Healing House + Kiora's Follower (speed 5)
- Ioreth of the Healing House + Vizier of Tumbling Sands (speed 5)
- Ioreth of the Healing House + Kelpie Guide (speed 5)
- Gwenna, Eyes of Gaea + Ioreth of the Healing House + Saryth, the Viper's Fang (speed 5)
- Kelpie Guide + Pemmin's Aura (speed 5)
- Kinnan, Bonder Prodigy + Grim Monolith + Spark Double (speed 5)
- Kinnan, Bonder Prodigy + Basalt Monolith (speed 5, two cards)

Findings:

- WARN `land_count`: 30 lands: the guide range for this format is 34 to 38
- INFO `curve_summary`: average mana value 2.55 over 69 nonland cards
- WARN `profile_off_band`: the average mana value of the nonland cards is 2.55, and bracket 5 wants 1.2 to 2.4
- WARN `profile_off_band`: the mana available on turn four is 4.38, and bracket 5 wants 5.5 or more (mean over 10000 hands)

Cards:

- 1 Bloodstained Mire
- 1 Breeding Pool
- 1 City of Brass
- 1 Command Tower
- 1 Dreamroot Cascade
- 1 Exotic Orchard
- 1 Flooded Strand
- 1 Misty Rainforest
- 1 Polluted Delta
- 1 Scalding Tarn
- 1 Verdant Catacombs
- 1 Windswept Heath
- 1 Wooded Foothills
- 1 Mana Confluence
- 1 Reflecting Pool
- 1 Rejuvenating Springs
- 1 Yavimaya Coast
- 1 Hinterland Harbor
- 1 Gemstone Caverns
- 1 Minamo, School at Water's Edge
- 1 Nykthos, Shrine to Nyx
- 5 Island
- 4 Forest
- 1 Arbor Elf
- 1 Basalt Monolith
- 1 Devoted Druid
- 1 Earthcraft
- 1 Grim Monolith
- 1 Mana Vault
- 1 Magus of the Candelabra
- 1 Sonic Screwdriver
- 1 Kiora's Follower
- 1 Ioreth of the Healing House
- 1 Vizier of Tumbling Sands
- 1 Lost Jitte
- 1 Honor-Worn Shaku
- 1 Patriar's Seal
- 1 Kelpie Guide
- 1 Gwenna, Eyes of Gaea
- 1 Azusa's Many Journeys // Likeness of the Seeker
- 1 Benefactor's Draught
- 1 Cerulean Wisps
- 1 Cloud of Faeries
- 1 Frantic Search
- 1 Innocuous Researcher
- 1 Refocus
- 1 Silent Hallcreeper
- 1 Twitch
- 1 Yavimaya Elder
- 1 Pip-Boy 3000
- 1 Shadow of the Second Sun
- 1 Transmutation Font
- 1 Staff of Domination
- 1 Biosynthic Burst
- 1 Crab Umbra
- 1 Fleeting Reflection
- 1 Flusterstorm
- 1 Katara's Reversal
- 1 Legolas's Quick Reflexes
- 1 Magic Damper
- 1 Muddle the Mixture
- 1 Narset's Reversal
- 1 Octopus Form
- 1 Shore Up
- 1 Sokka's Haiku
- 1 Unwind
- 1 Pemmin's Aura
- 1 Saryth, the Viper's Fang
- 1 Aggressive Biomancy
- 1 Amphibian Downpour
- 1 Apathy
- 1 Bushwhack
- 1 Emerald Charm
- 1 Gleeful Sabotage
- 1 Provoke
- 1 Snap
- 1 Super Combo
- 1 Fabricate
- 1 Worldly Tutor
- 1 Thassa's Oracle
- 1 Voltaic Construct
- 1 Filigree Sages
- 1 Phyrexian Metamorph
- 1 Cacophodon
- 1 Dross Scorpion
- 1 Breaching Hippocamp
- 1 Clone
- 1 Clever Impersonator
- 1 Spark Double
- 1 Stunt Double
- 1 Altered Ego
- 1 Cyclonic Rift

### 14. Bracket 5, Yuriko, the Tiger's Shadow, ninjas

Pool 213. 99 cards, 0 block findings, repaired true (profile_off_band).

Judge: error, judge bracket: llm anthropic/claude-opus-5: terminal (http 400): POST "https://api.anthropic.com/v1/messages": 400 Bad Request (Request-ID: req_011CeexFawG3nst7MvFfTYag) (Workspace-ID: wrkspc_01DKn3qmnB3ZMerq6QsCdj6C) {"type":"error","error":{"type":"invalid_request_error","message":"output_config.format.schema: For 'integer' type, properties maximum, minimum are not supported"},"request_id":"req_011CeexFawG3nst7MvFfTYag"}

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 31 | 27 to 33 |  |  |
| `tapped_land` | 1 | 0 to 2 |  | Choked Estuary |
| `colorless_land` | 0 | 0 to 6 |  |  |
| `color_sources` | 1.74 | 0.9 or more |  | sources of requirement: U 35.25 of 19, B 33 of 19 |
| `avg_mana_value` | 2.29 | 1.2 to 2.4 |  | over 68 nonland cards |
| `ramp` | 16 | 12 to 22 |  |  |
| `draw` | 13 | 8 to 18 |  |  |
| `removal` | 9 | 4 to 14 |  |  |
| `wipe` | 1 | 0 to 3 |  |  |
| `interaction` | 15 | 8 to 22 |  |  |
| `tutor` | 1 | none |  | Higure, the Still Wind |
| `fast_mana` | 10 | none |  | An Offer You Can't Refuse, Chrome Mox, Dark Ritual, Everflowing Chalice, Lotus Petal, Mana Vault, Mox Amber, Mox Diamond, Mox Opal, Sol Ring |
| `game_changer` | 6 | none |  | Chrome Mox, Fierce Guardianship, Force of Will, Mana Vault, Mox Diamond, Rhystic Study |
| `mana_turn_four` | 5.22 | 5.5 or more | yes | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.68 | 0.65 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.62 | -9 to 0.3 |  | the commander costs 3 and comes down on turn 2.38 on average |

Goldfish over 10000 hands: the commander on turn 2.38, 5.22 mana on turn four, and 0.68 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Chrome Mox, Fierce Guardianship, Force of Will, Mana Vault, Mox Diamond, Rhystic Study. Mass land denial none. Extra turns none. Combos 0.

Findings:

- WARN `land_count`: 31 lands: the guide range for this format is 34 to 38
- INFO `curve_summary`: average mana value 2.29 over 68 nonland cards
- WARN `profile_off_band`: the mana available on turn four is 5.22, and bracket 5 wants 5.5 or more (mean over 10000 hands)
- INFO `basics_added`: the list was 1 card short, so the builder added 1 basic land

Cards:

- 1 Underground Sea
- 1 Watery Grave
- 1 Morphic Pool
- 1 Drowned Catacomb
- 1 Shipwreck Marsh
- 1 Gloomlake Verge
- 1 Underground River
- 1 Clearwater Pathway // Murkwater Pathway
- 1 Choked Estuary
- 1 Sunken Ruins
- 1 Tainted Isle
- 1 Command Tower
- 1 City of Brass
- 1 Mana Confluence
- 1 Exotic Orchard
- 1 Forbidden Orchard
- 1 Cavern of Souls
- 1 Secluded Courtyard
- 1 Unclaimed Territory
- 1 Gemstone Caverns
- 7 Island
- 4 Swamp
- 1 Chrome Mox
- 1 Lotus Petal
- 1 Mana Vault
- 1 Mox Diamond
- 1 Mox Opal
- 1 Sol Ring
- 1 Dark Ritual
- 1 Arcane Signet
- 1 Dimir Signet
- 1 Fellwar Stone
- 1 Talisman of Dominance
- 1 Mox Amber
- 1 Ornithopter of Paradise
- 1 Wayfarer's Bauble
- 1 Mind Stone
- 1 Everflowing Chalice
- 1 Brainstorm
- 1 Baleful Strix
- 1 Faerie Mastermind
- 1 Frantic Search
- 1 Ingenious Infiltrator
- 1 Mystic Remora
- 1 Night's Whisper
- 1 Ninja of the Deep Hours
- 1 Opt
- 1 Ponder
- 1 Preordain
- 1 Rhystic Study
- 1 Skullclamp
- 1 An Offer You Can't Refuse
- 1 Arcane Denial
- 1 Counterspell
- 1 Fierce Guardianship
- 1 Flusterstorm
- 1 Force of Negation
- 1 Force of Will
- 1 Mental Misstep
- 1 Negate
- 1 Pact of Negation
- 1 Saiba Cryptomancer
- 1 Swan Song
- 1 Lightning Greaves
- 1 Swiftfoot Boots
- 1 Azra Smokeshaper
- 1 Accursed Marauder
- 1 Deadly Rollick
- 1 Dokuchi Silencer
- 1 Feed the Swarm
- 1 Go for the Throat
- 1 Infernal Grasp
- 1 Into the Flood Maw
- 1 Mistblade Shinobi
- 1 Pongify
- 1 Silver-Fur Master
- 1 Smoke Shroud
- 1 Fallen Shinobi
- 1 Higure, the Still Wind
- 1 Ink-Eyes, Servant of Oni
- 1 Futurist Operative
- 1 Foot Mystic
- 1 Moonblade Shinobi
- 1 Mukotai Ambusher
- 1 Ninja of the New Moon
- 1 Okiba-Gang Shinobi
- 1 Sakashima's Student
- 1 Silent-Blade Oni
- 1 Taeko, the Patient Avalanche
- 1 Toxic Deluge

### 15. Bracket 5, Najeela, the Blade-Blossom, warriors combat

Pool 306. 99 cards, 0 block findings, repaired true (profile_off_band, profile_off_band, profile_off_band, profile_off_band).

Judge: error, judge bracket: llm anthropic/claude-opus-5: terminal (http 400): POST "https://api.anthropic.com/v1/messages": 400 Bad Request (Request-ID: req_011CeexZ8XGhBenJcifLBtnS) (Workspace-ID: wrkspc_01DKn3qmnB3ZMerq6QsCdj6C) {"type":"error","error":{"type":"invalid_request_error","message":"output_config.format.schema: For 'integer' type, properties maximum, minimum are not supported"},"request_id":"req_011CeexZ8XGhBenJcifLBtnS"}

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 30 | 27 to 33 |  |  |
| `tapped_land` | 2 | 0 to 2 |  | Base Camp, Path of Ancestry |
| `colorless_land` | 0 | 0 to 6 |  |  |
| `color_sources` | 0.6 | 0.9 or more | yes | sources of requirement: W 14 of 19, U 14 of 19, B 14 of 19, R 19.5 of 19, G 15.5 of 26 |
| `avg_mana_value` | 3.42 | 1.2 to 2.4 | yes | over 69 nonland cards |
| `ramp` | 17 | 12 to 22 |  |  |
| `draw` | 13 | 8 to 18 |  |  |
| `removal` | 9 | 4 to 14 |  |  |
| `wipe` | 1 | 0 to 3 |  |  |
| `interaction` | 15 | 8 to 22 |  |  |
| `tutor` | 2 | none |  | Blaring Captain, Kassandra, Eagle Bearer |
| `fast_mana` | 0 | none |  |  |
| `game_changer` | 0 | none |  |  |
| `mana_turn_four` | 3.66 | 5.5 or more | yes | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.66 | 0.65 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | 0.38 | -9 to 0.3 | yes | the commander costs 3 and comes down on turn 3.38 on average |

Goldfish over 10000 hands: the commander on turn 3.38, 3.66 mana on turn four, and 0.66 of first hands hold two to four lands.

Content: Spellbook tag S. Game Changers none. Mass land denial none. Extra turns none. Combos 1.
- Najeela, the Blade-Blossom + Professional Face-Breaker (speed 5)

Findings:

- WARN `land_count`: 30 lands: the guide range for this format is 34 to 38
- INFO `curve_summary`: average mana value 3.42 over 69 nonland cards
- WARN `profile_off_band`: the worst color's share of the sources it needs is 0.6, and bracket 5 wants 0.9 or more (sources of requirement: W 14 of 19, U 14 of 19, B 14 of 19, R 19.5 of 19, G 15.5 of 26)
- WARN `profile_off_band`: the average mana value of the nonland cards is 3.42, and bracket 5 wants 1.2 to 2.4
- WARN `profile_off_band`: the mana available on turn four is 3.66, and bracket 5 wants 5.5 or more (mean over 10000 hands)
- WARN `profile_off_band`: the turns the commander comes down after its mana value is 0.38, and bracket 5 wants 0.3 at most (the commander costs 3 and comes down on turn 3.38 on average)

Cards:

- 1 Blood Crypt
- 1 Breeding Pool
- 1 Command Tower
- 1 Exotic Orchard
- 1 Path of Ancestry
- 1 Base Camp
- 1 Godless Shrine
- 1 Hallowed Fountain
- 1 Overgrown Tomb
- 1 Sacred Foundry
- 1 Steam Vents
- 1 Stomping Ground
- 1 Temple Garden
- 1 Watery Grave
- 3 Plains
- 3 Island
- 3 Swamp
- 4 Mountain
- 3 Forest
- 1 Heronblade Elite
- 1 Radha, Heir to Keld
- 1 Evendo Brushrazer
- 1 Scuzzback Scrounger
- 1 Ardent Electromancer
- 1 Baylen, the Haymaker
- 1 Brigid, Clachan's Heart // Brigid, Doun's Mind
- 1 Clive, Ifrit's Dominant // Ifrit, Warden of Inferno
- 1 Eivor, Wolf-Kissed
- 1 Gimli of the Glittering Caves
- 1 Gladiolus Amicitia
- 1 Hugs, Grisly Guardian
- 1 Jolene, the Plunder Queen
- 1 Neheb, Dreadhorde Champion
- 1 Professional Face-Breaker
- 1 Sun Warriors
- 1 Swashbuckler Extraordinaire
- 1 Akki Ronin
- 1 Azra Oddsmaker
- 1 Jacked Rabbit
- 1 Jaxis, the Troublemaker
- 1 Kassandra, Eagle Bearer
- 1 Kutzil, Malamet Exemplar
- 1 Mindblade Render
- 1 Neyith of the Dire Hunt
- 1 Oakhame Adversary
- 1 Peema Trailblazer
- 1 Samut, Vizier of Naktamun
- 1 Savvy Hunter
- 1 Setessan Champion
- 1 Concerted Defense
- 1 Kashi-Tribe Elite
- 1 Selfless Safewright
- 1 Selfless Samurai
- 1 Absolute Virtue
- 1 Ainok Strike Leader
- 1 Akiri, Fearless Voyager
- 1 Eladamri, Lord of Leaves
- 1 Hakoda, Selfless Commander
- 1 Linvala, Shield of Sea Gate
- 1 Multiclass Baldric
- 1 Thancred Waters
- 1 Vexilus Praetor
- 1 Winota, Joiner of Forces
- 1 Squad Commander
- 1 Cacophony Scamp
- 1 Practiced Tactics
- 1 Deadly Alliance
- 1 Dreadhorde Butcher
- 1 Goblin Cratermaker
- 1 Accursed Marauder
- 1 Fleshbag Marauder
- 1 Journey to Oblivion
- 1 Drana's Silencer
- 1 Rograkh, Son of Rohgahh
- 1 Combat Celebrant
- 1 Aspiring Champion
- 1 Blaring Captain
- 1 Chaos Terminator Lord
- 1 Doomskar Warrior
- 1 Emeria Captain
- 1 Herald of Secret Streams
- 1 Lightning Runner
- 1 Raiyuu, Storm's Edge
- 1 Surrak, the Hunt Caller
- 1 Vengeful Firebrand
- 1 Warg Rider
- 1 Zealous Conscripts
- 1 Goblin Chainwhirler

