# PR-14A bracket gate, judge lane

Run date: 2026-09-29. Card snapshot: 2026-09-04. Decks read from `/Volumes/SSD-1TB/decktome/docs/reference/pr14a-bracket-calibration-decks-2026-09-13.md`.

Verdict: PASS. The judge agreed with the bracket on 11 of 12 decks (91 percent, the bar is 80), with 0 judge errors. This document reads the judge bar alone: the block, band, and content bars are in the source document.

The judge read 9 of 9 precons at or above the lowest bracket their rules allow. A precon anchors no bracket, so the agreement leaves it out (F-162, D-793).

Calls 21. Cost $0.1226. Time 77 seconds.

## Run

- Suite `bracket-judge`, run `pr14a-bracket-calibration-judge-2026-09-29`, on 2026-09-29, commit `9fb2569`.
- Roles: judge on `claude-sonnet-5-5` (anthropic, effort medium).
- Versions: card snapshot 2026-09-04, bracket_judge prompt version 3, generate prompt version 16, source `pr14a-bracket-calibration-decks-2026-09-13.md`.
- Calls: 21. Cost: $0.1226. Time: 77 seconds.

| # | Built for | Judged | Agrees | Commander |
|---|---|---|---|---|
| 1 | floor 1 | 3 | at or above | Captain America, Team Leader |
| 2 | floor 1 | 3 | at or above | Cloud, Ex-SOLDIER |
| 3 | floor 1 | 2 | at or above | Esika, God of the Tree // The Prismatic Bridge |
| 4 | floor 4 | 4, raised from 3 to the rules floor | at or above | Zada, Hedron Grinder |
| 5 | floor 4 | 4 | at or above | Saheeli, Radiant Creator |
| 6 | floor 1 | 2 | at or above | Auntie Ool, Cursewretch |
| 7 | floor 4 | 4, raised from 3 to the rules floor | at or above | Éowyn, Shieldmaiden |
| 8 | floor 1 | 3 | at or above | Omo, Queen of Vesuva |
| 9 | floor 1 | 2 | at or above | Heroes in a Half Shell |
| 10 | 5 | 5 | yes | Sisay, Weatherlight Captain |
| 11 | 5 | 5 | yes | Kinnan, Bonder Prodigy |
| 12 | 5 | 5 | yes | Grenzo, Dungeon Warden |
| 13 | 5 | 5 | yes | Flubs, the Fool |
| 14 | 5 | 5 | yes | Heliod, the Radiant Dawn // Heliod, the Warped Eclipse |
| 15 | 5 | 5 | yes | Crystal, Inhuman Princess |
| 16 | 5 | 3 | no | Ghave, Guru of Spores |
| 17 | 5 | 5 | yes | Magda, Brazen Outlaw |
| 18 | 5 | 5 | yes | K'rrik, Son of Yawgmoth |
| 19 | 5 | 5 | yes | Nick Fury, Agent of S.H.I.E.L.D. |
| 20 | 5 | 5 | yes | Terra, Magical Adept // Esper Terra |
| 21 | 5 | 5 | yes | Shorikai, Genesis Engine |

## Why

### 1. Floor 1, Captain America, Team Leader, precon Avengers Assemble

Judge: bracket 3. The list has no marked Game Changers, no combos, no tutors, and no extra turns. It does run Sol Ring, Arcane Signet, Fellwar Stone and three Talismans for ramp, plus solid removal (Swords, Austere Command, Dismantling Wave) and a tidy three-color mana base, so it is stronger than a basic precon. It is a synergistic tribal deck with good card quality and no explosive or fast win, which fits Upgraded better than Core.

Combos the judge read: none, Commander Spellbook finds no combo.

### 2. Floor 1, Cloud, Ex-SOLDIER, precon Limit Break

Judge: bracket 3. The card list marks no Game Changers, and the combo check finds no combo. There is no mass land denial beyond Decimate, which is a single-target spell. The deck does run Sol Ring, Arcane Signet, Skullclamp, Lightning Greaves, several wraths and a fair amount of ramp, with Boros-Selesnya equipment synergy and solid card quality, which puts it above a basic precon feel. It has no tutors or fast combo, so it isn't optimized, and Bracket 3 fits better than 2 given the powered-up synergy and card quality.

Combos the judge read: none, Commander Spellbook finds no combo.

### 3. Floor 1, Esika, God of the Tree // The Prismatic Bridge, precon From Cute to Brute

Judge: bracket 2. The card list marks no Game Changers, and Commander Spellbook finds no combos. There is no mass land denial or extra turns, and the only fast mana is Sol Ring plus a few signets and rocks. It is a five-color good-stuff pile of double-faced cards with modest removal and a few sweepers, so it wins incrementally through board presence and is not tuned for speed or consistency.

Combos the judge read: none, Commander Spellbook finds no combo.

### 4. Floor 4, Zada, Hedron Grinder, precon Goblin Storm

Judge: bracket 4, raised from 3 to the rules floor. No cards are marked as Game Changers, and there is no mass land denial or chained extra turns. It is a synergistic Goblin and storm deck with Sol Ring, Skullclamp, rituals, Past in Flames, Grapeshot and Krenko, Mob Boss, which is stronger than a Core deck. The Storm-Kiln Artist + Haze of Rage combo and the three-card Krenko lines add speed, but none look fast or refined enough for Optimized.

Combos the judge read: Krenko, Mob Boss + Skirk Prospector + Howlsquad Heavy; Krenko, Mob Boss + Skirk Prospector + Goblin Warchief; Krenko, Mob Boss + Skirk Prospector + Goblin Chieftain; Storm-Kiln Artist + Haze of Rage.

### 5. Floor 4, Saheeli, Radiant Creator, precon Living Energy

Judge: bracket 4. The marked Game Changers aren't identified in the list, so I judged from the combos and power level. Commander Spellbook finds several two-card infinite combos that need no mana past the cards, such as Lightning Runner with Aetherwind Basker. That exceeds the Bracket 3 limit of no two-card infinite combo needing four mana or less. The deck is a fairly focused Aether/energy shell without heavy tutors or a full fast-mana suite, so it sits in Optimized rather than cEDH.

Combos the judge read: Saheeli, Radiant Creator + Lightning Runner + Aetherwind Basker; Saheeli, Radiant Creator + Lightning Runner + Peema Aether-Seer; Aethertide Whale + Nissa, Worldsoul Speaker + Panharmonicon; Lightning Runner + Aetherwind Basker.

### 6. Floor 1, Auntie Ool, Cursewretch, precon Blight Curse

Judge: bracket 2. No cards are marked as Game Changers, and there is no mass land denial, no extra turns, and no real two-card infinite combo. The listed Blowfly Infestation pairings are value interactions, not infinite combos. The deck is a -1/-1 counter and persist/undying value pile with Sol Ring, Arcane Signet and a few mana dorks as its only acceleration, plus ordinary removal and a few sweepers. It plays a straightforward, incremental game with no tutors and no fast win.

Combos the judge read: Blowfly Infestation + Flourishing Defenses; Blowfly Infestation + Hapatra, Vizier of Poisons.

### 7. Floor 4, Éowyn, Shieldmaiden, precon Riders of Rohan

Judge: bracket 4, raised from 3 to the rules floor. The list has no cards marked as Game Changers, no mass land denial, and no chained extra turns. It does play Sol Ring, a Supreme Verdict sweeper, and a tribal-style Human/soldier synergy core with good removal. However, Commander Spellbook lists a two-card infinite combo (Gilraen + Village Bell-Ringer) that costs no mana past the cards, which exceeds what brackets 1-3 allow, so the deck sits at Bracket 4 by the rules, though its overall power is closer to a synergistic midrange deck.

Combos the judge read: Gilraen, Dúnedain Protector + Village Bell-Ringer.

### 8. Floor 1, Omo, Queen of Vesuva, precon Tricky Terrain

Judge: bracket 3. The card list marks no Game Changers, and Commander Spellbook finds no combos. The deck is a land-ramp and value pile with Sol Ring, Arcane Signet and Chromatic Lantern, plus a Tron package, Cloudpost and Vesuva. It has strong synergy and payoffs such as Avenger, Terastodon and Hydroid Krasis, and it can Terastodon-destroy lands. Oblivion Stone and Evacuation are only board wipes, not mass land denial. Its power is above a precon, but it has no tutors or fast combo finish, so it sits at Upgraded rather than Optimized.

Combos the judge read: none, Commander Spellbook finds no combo.

### 9. Floor 1, Heroes in a Half Shell, precon Turtle Power

Judge: bracket 2. The card list has no cards marked as Game Changers, no combos found, no tutors, no extra-turn chains and no mass land denial. The only fast mana is Sol Ring and Arcane Signet, and the rest is a five-color Ninja Turtles theme pile with a mix of creatures, ramp and some sweepers like Blasphemous Act. That makes it a straightforward, incremental deck that plays at Core level.

Combos the judge read: none, Commander Spellbook finds no combo.

### 10. Bracket 5, Sisay, Weatherlight Captain, cedh 2026-09-06, place 1 of 65

Judge: bracket 5. The deck is a five-color, high-power list built to win fast and not to play a theme: original dual lands, a large stack of free and cheap counterspells (Force of Will, Fierce Guardianship, Pact of Negation, Misdirection), and a lot of fast mana (Mana Vault, Chrome Mox, Mox Diamond, Lotus Petal, LED, Dark Ritual). It has 18 marked Game Changers and many tutors (Demonic, Vampiric, Mystical, Enlightened, Gamble, Crop Rotation). It also has the listed two-card combos Demonic Consultation + Thassa's Oracle and Tainted Pact + Thassa's Oracle, plus Underworld Breach + Brain Freeze lines, so it can win on any turn, which is cEDH.

Combos the judge read: Faerie Mastermind + Smothering Tithe + Kinnan, Bonder Prodigy; Underworld Breach + Lotus Petal + Brain Freeze; Tainted Pact + Thassa's Oracle; Underworld Breach + Lion's Eye Diamond + Brain Freeze; Demonic Consultation + Thassa's Oracle.

### 11. Bracket 5, Kinnan, Bonder Prodigy, cedh 2026-09-06, place 1 of 40

Judge: bracket 5. The list is a Kinnan/Thrasios-style blue-green combo-control deck built to win as fast as possible. It runs 13 marked Game Changers (Ancient Tomb, Chrome Mox, Consecrated Sphinx, Crop Rotation, Fierce Guardianship, Force of Will, Gaea's Cradle, Mana Vault, Mox Diamond, Rhystic Study, Seedborn Muse, Thassa's Oracle, The One Ring), a large amount of fast mana (Sol Ring, Moxen, Lotus Petal, Mana Vault, Elvish Spirit Guide), free counterspells (Force of Will, Force of Negation, Fierce Guardianship, Mental Misstep, Mana Drain), and tutors (Chord of Calling, Crop Rotation, Whir of Invention, Transmute Artifact). Commander Spellbook finds many two-card infinite combos that need no extra mana, such as Hullbreaker Horror + Mox Amber, Kinnan + Basalt Monolith, Bloom Tender + Pemmin's Aura, and Incubation Druid + Freed from the Real, so it can win on any turn and plays the best strategy rather than a theme.

Combos the judge read: Hullbreaker Horror + Mox Amber; Kinnan, Bonder Prodigy + Delighted Halfling + Pemmin's Aura; Kinnan, Bonder Prodigy + Delighted Halfling + Freed from the Real; Badgermole Cub + Pemmin's Aura; Badgermole Cub + Freed from the Real; Devoted Druid + Agatha's Soul Cauldron + Incubation Druid; Hullbreaker Horror; Hullbreaker Horror + Sol Ring; Hullbreaker Horror + Mana Vault; Machine God's Effigy + Devoted Druid; Incubation Druid + Freed from the Real; Incubation Druid + Pemmin's Aura; Kinnan, Bonder Prodigy + Birds of Paradise + Pemmin's Aura; Kinnan, Bonder Prodigy + Birds of Paradise + Freed from the Real; Kinnan, Bonder Prodigy + Basalt Monolith; Kinnan, Bonder Prodigy + Bloom Tender + Freed from the Real; Bloom Tender + Pemmin's Aura; Bloom Tender + Freed from the Real.

### 12. Bracket 5, Grenzo, Dungeon Warden, cedh 2026-09-06, place 2 of 69

Judge: bracket 5. The deck has seven marked Game Changers (Ancient Tomb, Chrome Mox, Demonic Tutor, Mana Vault, Mox Diamond, Opposition Agent, Orcish Bowmasters, Vampiric Tutor), which is well past the Bracket 3 limit. It also has a large pile of fast mana (Sol Ring, Mana Vault, Moxen, Lotus Petal, Simian Spirit Guide, City of Traitors), cheap tutors, and a two-card infinite combo of Kiki-Jiki and Zealous Conscripts that needs no extra mana, plus several three-card combos. The list is built around the best strategy with a fast, efficient, metagame-style shell rather than a theme, so it plays as cEDH.

Combos the judge read: Pitiless Plunderer + Umbral Collar Zealot + Stridehangar Automaton; Pitiless Plunderer + Ashnod's Altar + Stridehangar Automaton; Pitiless Plunderer + Woe Strider + Stridehangar Automaton; Pitiless Plunderer + Viscera Seer + Stridehangar Automaton; Pitiless Plunderer + Carrion Feeder + Stridehangar Automaton; Warren Soultrader + Zulaport Cutthroat + Stridehangar Automaton; Grenzo, Dungeon Warden + Workhorse + Hoverstone Pilgrim; Grenzo, Dungeon Warden + Workhorse + Barkform Harvester; Grenzo, Dungeon Warden + Orcish Bowmasters + Ashnod's Altar + Barkform Harvester; Grenzo, Dungeon Warden + Orcish Bowmasters + Ashnod's Altar + Hoverstone Pilgrim; Kiki-Jiki, Mirror Breaker + Zealous Conscripts.

### 13. Bracket 5, Flubs, the Fool, cedh 2026-09-06, place 4 of 65

Judge: bracket 5. The deck is a compressed, metagame-style storm and combo build. It runs a lot of fast mana (Sol Ring, Mox Diamond, Chrome Mox, Lotus Petal, Lion's Eye Diamond, Ancient Tomb, the spirit guides) and Mystical Tutor, Worldly Tutor and Crop Rotation as tutors. It has 9 marked Game Changers, well over the bracket 3 limit. Commander Spellbook lists several combos, including the two-card Squee + Food Chain, plus Underworld Breach lines with Brain Freeze, Lion's Eye Diamond and Wheel of Fortune that need no mana past the cards, so it can win on any turn.

Combos the judge read: Cloudstone Curio + Crimson Kobolds + Rograkh, Son of Rohgahh; Cloudstone Curio + Crimson Kobolds + Kobolds of Kher Keep; Cloudstone Curio + Rograkh, Son of Rohgahh + Crookshank Kobolds; Cloudstone Curio + Crookshank Kobolds + Kobolds of Kher Keep; Cloudstone Curio + Crimson Kobolds + Crookshank Kobolds; Oboro Breezecaller + Talon Gates of Madara + Gaea's Cradle; Flubs, the Fool + Underworld Breach + Lion's Eye Diamond + Life from the Loam; Underworld Breach + Lotus Petal + Brain Freeze; Underworld Breach + Wheel of Fortune + Jeska's Will; Cloudstone Curio + Rograkh, Son of Rohgahh + Kobolds of Kher Keep; Squee, the Immortal + Food Chain; Underworld Breach + Lion's Eye Diamond + Wheel of Fortune; Underworld Breach + Lion's Eye Diamond + Brain Freeze.

### 14. Bracket 5, Heliod, the Radiant Dawn // Heliod, the Warped Eclipse, cedh 2026-09-05, place 1 of 45

Judge: bracket 5. The list carries 14 marked Game Changers (Ancient Tomb, Chrome Mox, Enlightened Tutor, Fierce Guardianship, Force of Will, Gifts Ungiven, Grim Monolith, Intuition, Mana Vault, Mox Diamond, Mystical Tutor, Rhystic Study, Smothering Tithe, The One Ring), far past the bracket 3 limit. It also plays a lot of fast mana (Sol Ring, Lotus Petal, Mox Opal, City of Traitors, Gemstone Caverns), cheap free counterspells, Timetwister, and a dense tutor suite, with a fetch and dual mana base. Its lines are built around Hullbreaker Horror, Paradoxical Outcome and Displacer Kitten, and it plays as a tight, metagame-focused cEDH shell.

Combos the judge read: Faerie Mastermind + Smothering Tithe + Mirrormade; Faerie Mastermind + Smothering Tithe + Clever Impersonator; Hullbreaker Horror; Hullbreaker Horror + Sol Ring; Hullbreaker Horror + Mana Vault.

### 15. Bracket 5, Crystal, Inhuman Princess, cedh 2026-09-05, place 1 of 64

Judge: bracket 5. This is a storm-style cEDH shell with free counterspells (Force of Will, Daze, Mental Misstep, Misdirection, Pact of Negation, Fierce Guardianship), free mana (Chrome Mox, Mox Diamond, Lotus Petal, LED, Spirit Guides, Sol Ring), and tutors (Enlightened Tutor, Gamble, Intuition, Crop Rotation). It runs 10 Game Changers, and the Underworld Breach + LED/Lotus Petal + Brain Freeze lines win with no mana past the cards. The list follows one best-strategy plan built to win on any turn, and the fast mana, tutors, free interaction and a dense fetch/dual mana base all put it in cEDH.

Combos the judge read: Underworld Breach + Lotus Petal + Brain Freeze; Underworld Breach + Lion's Eye Diamond + Brain Freeze.

### 16. Bracket 5, Ghave, Guru of Spores, cedh 2026-09-05, place 1 of 37

Judge: bracket 3. The deck has one marked Game Changer (Aura Shards), no combos, no extra turns and no mass land denial, so it fits within Bracket 3's limits. Sol Ring, signets and Skullclamp add some acceleration and card advantage, but the rest is a loose Abzan token and counters pile with modest removal and little tutoring. It plays a little above the precon level of Bracket 2 because of Aura Shards, Skullclamp and Sol Ring, but it is not tuned or fast enough for Bracket 4.

Combos the judge read: none, Commander Spellbook finds no combo.

### 17. Bracket 5, Magda, Brazen Outlaw, cedh 2026-09-05, place 1 of 64

Judge: bracket 5. The deck is a fast, all-in Dwarf and artifact build with a dense pile of fast mana: Sol Ring, Mana Vault, Chrome Mox, Mox Opal, Lotus Petal, Simian Spirit Guide, Ancient Tomb, City of Traitors, Gemstone Caverns and Jeska's Will. It also plays free counterspell-style cards like Pyroblast and Red Elemental Blast, so it is built around speed and metagame interaction rather than a theme. Commander Spellbook lists several two-card infinite combos with no mana needed, such as Magda + Clock of Omens with Firdoch Core, Three Tree Mascot or Universal Automaton. That is well past what Bracket 3 allows, and with four Game Changers and this much fast mana, it plays as cEDH.

Combos the judge read: Magda, Brazen Outlaw + Clock of Omens + Firdoch Core; Magda, Brazen Outlaw + Battered Golem + Stalactite Dagger; Battered Golem + Dwarven Bloodboiler + Magda, Brazen Outlaw + Stalactite Dagger; Magda, Brazen Outlaw + Clock of Omens + Three Tree Mascot; Magda, Brazen Outlaw + Clock of Omens + Barkform Harvester; Magda, Brazen Outlaw + Battered Golem + Maskwood Nexus; Magda, Brazen Outlaw + Clock of Omens + Roaming Throne; Battered Golem + Dwarven Bloodboiler + Magda, Brazen Outlaw + Maskwood Nexus; Magda, Brazen Outlaw + Clock of Omens + Liquimetal Torque; Magda, Brazen Outlaw + Clock of Omens + Universal Automaton.

### 18. Bracket 5, K'rrik, Son of Yawgmoth, cedh 2026-09-05, place 1 of 39

Judge: bracket 5. The deck has 10 marked Game Changers (Ancient Tomb, Chrome Mox, Demonic Tutor, Grim Monolith, Imperial Seal, Lion's Eye Diamond, Mana Vault, Mox Diamond, Necropotence, Vampiric Tutor, Orcish Bowmasters), far past the Bracket 3 limit. It is built around a lot of fast mana (Sol Ring, Moxen, Dark Ritual, Cabal Ritual, LED, Lotus Petal), cheap tutors, and reanimation (Entomb, Reanimate, Animate Dead, Necromancy, Buried Alive, Shallow Grave) to cheat big threats out early. It also has cheap two-card combos such as Sheoldred with Vilis and K'rrik, so it plays as a stripped-down, speed-focused cEDH shell rather than a themed deck.

Combos the judge read: Asmodeus the Archfiend + Necrotic Ooze + K'rrik, Son of Yawgmoth + Sheoldred, the Apocalypse; Sheoldred, the Apocalypse + Vilis, Broker of Blood + Necropotence; Sheoldred, the Apocalypse + Vilis, Broker of Blood + Blood Celebrant; Orcish Bowmasters + Peer into the Abyss; Sheoldred, the Apocalypse + Vilis, Broker of Blood + Skirge Familiar; Sheoldred, the Apocalypse + Vilis, Broker of Blood + K'rrik, Son of Yawgmoth; Sheoldred, the Apocalypse + Peer into the Abyss; Asmodeus the Archfiend + Necrotic Ooze + Skirge Familiar.

### 19. Bracket 5, Nick Fury, Agent of S.H.I.E.L.D., cedh 2026-09-05, place 1 of 63

Judge: bracket 5. The list is a compressed, metagame-focused storm and combo deck. It runs a full set of free counterspells (Force of Will, Fierce Guardianship, Daze, Force of Negation), the classic fast mana (Mana Vault, Chrome Mox, Mox Diamond, LED, Lotus Petal, Sol Ring), a dual and fetch mana base, and many tutors. It has 17 marked Game Changers and several ways to win out of nowhere, including Demonic Consultation + Thassa's Oracle and Tainted Pact + Thassa's Oracle, which are both two-card combos needing at most four mana, plus Underworld Breach + Brain Freeze lines. It plays the best strategy and not a theme, so it sits at cEDH.

Combos the judge read: Faerie Mastermind + Smothering Tithe + Copy Enchantment; Underworld Breach + Lotus Petal + Brain Freeze; Tainted Pact + Thassa's Oracle; Underworld Breach + Lion's Eye Diamond + Brain Freeze; Demonic Consultation + Thassa's Oracle.

### 20. Bracket 5, Terra, Magical Adept // Esper Terra, cedh 2026-09-05, place 2 of 64

Judge: bracket 5. The deck is a pile of the strongest efficient cards: free counterspells (Force of Will, Force of Negation, Fierce Guardianship, Pact of Negation), a full suite of fast mana (Mana Vault, Chrome Mox, Mox Diamond, LED, Sol Ring, Lotus Petal, spirit guides), and many tutors (Demonic Tutor, Vampiric, Enlightened, Worldly, Gamble, Intuition). It has far more than three marked Game Changers, plus several cheap two-card infinite combos, such as Terra with Spark Double and Keen Sense with Niv-Mizzet, and Underworld Breach with Brain Freeze. With a fetch, dual and Ancient Tomb mana base, and a mix of stax and interaction built to win fast and hold up protection, this is a competitive, metagame-focused cEDH list and not a themed deck.

Combos the judge read: Terra, Magical Adept // Esper Terra + Mirrormade; Terra, Magical Adept // Esper Terra + Estrid's Invocation; Terra, Magical Adept // Esper Terra + Copy Enchantment; Terra, Magical Adept // Esper Terra + Spark Double; Faerie Mastermind + Smothering Tithe + Mirrormade; Faerie Mastermind + Smothering Tithe + Estrid's Invocation; Faerie Mastermind + Smothering Tithe + Copy Enchantment; Underworld Breach + Lotus Petal + Brain Freeze; Underworld Breach + Lion's Eye Diamond + Brain Freeze; Keen Sense + Niv-Mizzet, Parun.

### 21. Bracket 5, Shorikai, Genesis Engine, cedh 2026-09-05, place 2 of 60

Judge: bracket 5. The list is a competitive, metagame-focused shell. It has a dense free-counterspell suite (Force of Will, Force of Negation, Misdirection, Mental Misstep, Pact of Negation, Fierce Guardianship, Flusterstorm, Swan Song) and a large amount of fast mana (Sol Ring, Mana Vault, Grim Monolith, Chrome Mox, Mox Diamond, Mox Opal, Lotus Petal, Ancient Tomb, City of Traitors). It also has tutors (Enlightened Tutor, Whir of Invention, Transmute Artifact) and Game Changers well past the bracket 3 limit, including Rhystic Study, Smothering Tithe and The One Ring. The mana base is optimized with fetches and duals, and the deck wins through three-card combos like Teferi + Displacer Kitten + a mana source and Faerie Mastermind + Smothering Tithe + a clone. It plays the best strategy rather than a theme, so it sits at cEDH.

Combos the judge read: Teferi, Time Raveler + Displacer Kitten + Grim Monolith; Faerie Mastermind + Smothering Tithe + Mirrormade; Faerie Mastermind + Smothering Tithe + Copy Enchantment; Faerie Mastermind + Smothering Tithe + Clever Impersonator; Hullbreaker Horror; Hullbreaker Horror + Sol Ring; Hullbreaker Horror + Mana Vault; The One Ring + Displacer Kitten + Teferi, Time Raveler; Teferi, Time Raveler + Displacer Kitten + Sol Ring; Teferi, Time Raveler + Displacer Kitten + Mox Opal.

