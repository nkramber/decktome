# PR-14A bracket gate

Run date: 2026-09-23. Card snapshot: 2026-09-04.

Verdict: FAIL. 1 of 1 decks passed every block check, 1 sat in every band, 1 held no content violation, and the judge agreed with the bracket on 0 of 1 (0 percent, the bar is 80).

## Summary

| Measure | Value |
|---|---|
| Prompts | 1 |
| Decks returned | 1 |
| Decks with no block finding | 1 |
| Decks in every band | 1 |
| Decks with no content violation | 1 |
| Decks the endpoint did not check | 0 |
| Decks that needed a repair turn | 0 |
| Decks judged | 1 |
| Judge agreed with the bracket | 0 |
| Judge errors | 0 |
| Errors | 0 |
| Prompt version | 16 |
| Calls | 2 |
| Cost | $0.1083 |
| Time | 107 seconds |

## Run

- Suite `bracket`, run `pr14a-bracket-gate-run10`, on 2026-09-23, commit `433d055`.
- Partial run over `15`. It reads the bars of its own items, and it never stands for the suite.
- Roles: generate on `gpt-5.6-terra` (openai, effort medium), judge on `claude-opus-5` (anthropic, effort medium), repair on `gpt-5.6-terra` (openai, effort medium).
- Versions: card snapshot 2026-09-04, bracket_judge prompt version 3, generate prompt version 16.
- Calls: 2. Cost: $0.1083. Time: 107 seconds.

## Off-band features by key

Each row counts the decks whose feature sat outside its band after the repair turns.

| Feature | Decks off band |
|---|---|
| none | 0 |

## Decks

### 15. Bracket 5, Najeela, the Blade-Blossom, warriors combat

Pool 338. 99 cards, 0 block findings, repaired false.

Judge: bracket 4, disagrees. Twelve-plus marked Game Changers, a full fast-mana suite (Lion's Eye Diamond, Mana Vault, the Moxen, Lotus Petal, Jeska's Will) and a black tutor package (Demonic, Vampiric, Imperial Seal) put this far beyond Upgraded, and Najeela plus Thassa's Oracle gives it explosive turn-three or four kills. However, the build is still a Warriors tribal theme with filler like Aysen Crusader, Icewind Stalwart and Pact of the Serpent, and no listed combo line, so it isn't a metagame-tuned cEDH list. That lands it squarely in Bracket 4, Optimized.

| Feature | Value | Band | Off band | Note |
|---|---|---|---|---|
| `land` | 32 | 27 to 33 |  |  |
| `tapped_land` | 1 | 0 to 2 |  | Base Camp |
| `colorless_land` | 0 | 0 to 6 |  |  |
| `color_sources` | 1.18 | 0.9 or more |  | sources of requirement: W 22.5 of 19, U 23.5 of 19, B 22.5 of 19, R 22.5 of 19, G 22.75 of 19 |
| `fixing_land` | 24 | 21 or more |  | Adarkar Wastes, Base Camp, Battlefield Forge, Blood Crypt, Breeding Pool, Cavern of Souls, Caves of Koilos, City of Brass, Command Tower, Godless Shrine, Hallowed Fountain, Llanowar Wastes, Mana Confluence, Overgrown Tomb, Sacred Foundry, Secluded Courtyard, Shivan Reef, Steam Vents, Stomping Ground, Temple Garden, Unclaimed Territory, Underground River, Watery Grave, Yavimaya Coast |
| `avg_mana_value` | 2.39 | 1.2 to 2.4 |  | over 67 nonland cards |
| `ramp` | 18 | 12 to 22 |  |  |
| `draw` | 8 | 8 to 18 |  |  |
| `removal` | 9 | 4 to 14 |  |  |
| `wipe` | 1 | 0 to 3 |  |  |
| `interaction` | 13 | 8 to 22 |  |  |
| `tutor` | 5 | 4 or more |  | Blaring Captain, Demonic Tutor, Enlightened Tutor, Imperial Seal, Vampiric Tutor |
| `fast_mana` | 12 | 6 or more |  | An Offer You Can't Refuse, Chrome Mox, Dark Ritual, Lion's Eye Diamond, Lotus Petal, Mana Vault, Mox Amber, Mox Diamond, Mox Opal, Rite of Flame, Sol Ring, Springleaf Drum |
| `game_changer` | 14 | 8 or more |  | Chrome Mox, Demonic Tutor, Enlightened Tutor, Fierce Guardianship, Imperial Seal, Jeska's Will, Lion's Eye Diamond, Mana Vault, Mox Diamond, Orcish Bowmasters, Rhystic Study, Smothering Tithe, Thassa's Oracle, Vampiric Tutor |
| `finisher` | 1 | 1 or more |  | Thassa's Oracle |
| `mana_turn_four` | 5.25 | 4.6 or more |  | mean over 10000 hands |
| `hands_two_to_four_lands` | 0.69 | 0.65 or more |  | share of first seven-card hands |
| `commander_turn_over_mv` | -0.65 | -9 to 0.3 |  | the commander costs 3 and comes down on turn 2.35 on average |

Goldfish over 10000 hands: the commander on turn 2.35, 5.25 mana on turn four, and 0.69 of first hands hold two to four lands.

Content: Spellbook tag R. Game Changers Chrome Mox, Lion's Eye Diamond, Mana Vault, Mox Diamond, Jeska's Will, Smothering Tithe, Rhystic Study, Fierce Guardianship, Orcish Bowmasters, Thassa's Oracle, Demonic Tutor, Enlightened Tutor, Imperial Seal, Vampiric Tutor. Mass land denial none. Extra turns none. Combos 0.

Findings:

- INFO `curve_summary`: average mana value 2.39 over 67 nonland cards
- INFO `mana_pass`: the builder moved 5 cards of the mana base to bring the deck inside its power level
- INFO `cards_trimmed`: the list was 1 card over, so the builder cut this card: Harmonized Crescendo

Cards:

- 1 Adarkar Wastes
- 1 Battlefield Forge
- 1 Blood Crypt
- 1 Breeding Pool
- 1 Cavern of Souls
- 1 Caves of Koilos
- 1 City of Brass
- 1 Command Tower
- 1 Godless Shrine
- 1 Hallowed Fountain
- 1 Llanowar Wastes
- 1 Mana Confluence
- 1 Overgrown Tomb
- 1 Sacred Foundry
- 1 Secluded Courtyard
- 1 Shivan Reef
- 1 Steam Vents
- 1 Stomping Ground
- 1 Temple Garden
- 1 Underground River
- 1 Unclaimed Territory
- 1 Watery Grave
- 1 Yavimaya Coast
- 3 Forest
- 1 Island
- 1 Mountain
- 1 Plains
- 1 Swamp
- 1 Base Camp
- 1 Fire Nation Palace
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
- 1 White Lotus Tile
- 1 Smothering Tithe
- 1 Rhystic Study
- 1 Mindblade Render
- 1 Folk Hero
- 1 For the Ancestors
- 1 Kutzil, Malamet Exemplar
- 1 Oakhame Adversary
- 1 Pact of the Serpent
- 1 Samut, Vizier of Naktamun
- 1 An Offer You Can't Refuse
- 1 Fierce Guardianship
- 1 Ainok Strike Leader
- 1 Akiri, Fearless Voyager
- 1 And They Shall Know No Fear
- 1 Concerted Defense
- 1 Etchings of the Chosen
- 1 Galadhrim Ambush
- 1 Haunted One
- 1 Multiclass Baldric
- 1 Selfless Samurai
- 1 Steely Resolve
- 1 Seasoned Dungeoneer
- 1 Orcish Bowmasters
- 1 Accursed Marauder
- 1 Cacophony Scamp
- 1 Chatterfang, Squirrel General
- 1 Dreadhorde Butcher
- 1 Fleshbag Marauder
- 1 Merciless Executioner
- 1 Practiced Tactics
- 1 Shared Animosity
- 1 Aspiring Champion
- 1 Aysen Crusader
- 1 Blaring Captain
- 1 Cascade Seer
- 1 Icewind Stalwart
- 1 Kabira Outrider
- 1 Lovisa Coldeyes
- 1 Surrak, the Hunt Caller
- 1 Tazri, Beacon of Unity
- 1 Vengeful Firebrand
- 1 Warg Rider
- 1 Thassa's Oracle
- 1 Crippling Fear
- 1 Demonic Tutor
- 1 Enlightened Tutor
- 1 Imperial Seal
- 1 Vampiric Tutor
- 1 Shang-Chi, Master of Kung Fu
- 1 Stadium Headliner
- 1 Archpriest of Iona

