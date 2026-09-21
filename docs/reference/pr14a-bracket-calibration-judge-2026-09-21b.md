# PR-14A bracket gate, judge lane

Run date: 2026-09-21. Card snapshot: 2026-09-04. Decks read from `/Users/nate/Repos/decktome/docs/reference/pr14a-bracket-calibration-decks-2026-09-13.md`.

Verdict: FAIL. The judge agreed with the bracket on 10 of 12 decks (83 percent, the bar is 80), with 0 judge errors. This document reads the judge bar alone: the block, band, and content bars are in the source document.

The judge read 6 of 9 precons at or above the lowest bracket their rules allow. A precon anchors no bracket, so the agreement leaves it out (F-162, D-793).

Calls 21. Cost $0.3107. Time 120 seconds.

## Run

- Suite `bracket-judge`, run `pr14a-bracket-calibration-judge-2026-09-21b`, on 2026-09-21, commit `88e3870`.
- Roles: judge on `claude-opus-5` (anthropic, effort medium).
- Versions: card snapshot 2026-09-04, bracket_judge prompt version 3, generate prompt version 15, source `pr14a-bracket-calibration-decks-2026-09-13.md`.
- Calls: 21. Cost: $0.3107. Time: 120 seconds.

| # | Built for | Judged | Agrees | Commander |
|---|---|---|---|---|
| 1 | floor 1 | 2 | at or above | Captain America, Team Leader |
| 2 | floor 1 | 2 | at or above | Cloud, Ex-SOLDIER |
| 3 | floor 1 | 2 | at or above | Esika, God of the Tree // The Prismatic Bridge |
| 4 | floor 4 | 3 | no | Zada, Hedron Grinder |
| 5 | floor 4 | 3 | no | Saheeli, Radiant Creator |
| 6 | floor 1 | 2 | at or above | Auntie Ool, Cursewretch |
| 7 | floor 4 | 3 | no | Éowyn, Shieldmaiden |
| 8 | floor 1 | 3 | at or above | Omo, Queen of Vesuva |
| 9 | floor 1 | 2 | at or above | Heroes in a Half Shell |
| 10 | 5 | 5 | yes | Sisay, Weatherlight Captain |
| 11 | 5 | 5 | yes | Kinnan, Bonder Prodigy |
| 12 | 5 | 4 | no | Grenzo, Dungeon Warden |
| 13 | 5 | 5 | yes | Flubs, the Fool |
| 14 | 5 | 5 | yes | Heliod, the Radiant Dawn // Heliod, the Warped Eclipse |
| 15 | 5 | 5 | yes | Crystal, Inhuman Princess |
| 16 | 5 | 2 | no | Ghave, Guru of Spores |
| 17 | 5 | 5 | yes | Magda, Brazen Outlaw |
| 18 | 5 | 5 | yes | K'rrik, Son of Yawgmoth |
| 19 | 5 | 5 | yes | Nick Fury, Agent of S.H.I.E.L.D. |
| 20 | 5 | 5 | yes | Terra, Magical Adept // Esper Terra |
| 21 | 5 | 5 | yes | Shorikai, Genesis Engine |

## Why

### 1. Floor 1, Captain America, Team Leader, precon Avengers Assemble

Judge: bracket 2. This is a straightforward Avengers-typal deck built around creature synergies, anthem/tribal payoffs like Door of Destinies and Kindred Discovery, and a value-oriented board presence rather than any combo finish. No Game Changers are marked, no combos are found, and the interaction is basic removal (Swords to Plowshares, Rip Apart, Austere Command) with only Sol Ring as fast mana. Wins are incremental and telegraphed on the board, which is classic Bracket 2 Core.

Combos the judge read: none, Commander Spellbook finds no combo.

### 2. Floor 1, Cloud, Ex-SOLDIER, precon Limit Break

Judge: bracket 2. This is a straightforward Final Fantasy-themed equipment/voltron deck with no Game Changers, no mass land denial, no extra turns, and no combos found. Its wins are incremental and board-telegraphed via suited-up creatures, and the mana base is mostly taplands and basics with only Sol Ring/Arcane Signet as acceleration. Removal and card draw are present but unoptimized, placing it comfortably in Core.

Combos the judge read: none, Commander Spellbook finds no combo.

### 3. Floor 1, Esika, God of the Tree // The Prismatic Bridge, precon From Cute to Brute

Judge: bracket 2. This is a themed double-faced-card pile with no marked Game Changers and no combos found, winning through incremental value from Esika's Prismatic Bridge hits rather than any focused engine. The mana base is slow and clunky (many taplands, vivid/tri-lands, only Sol Ring and basic signets for ramp), and interaction is a handful of one-for-one removal plus a wrath. That unoptimized, telegraphed, board-based plan lands squarely in Core.

Combos the judge read: none, Commander Spellbook finds no combo.

### 4. Floor 4, Zada, Hedron Grinder, precon Goblin Storm

Judge: bracket 3. No marked Game Changers and no mass land denial, but the deck is a tuned storm/aristocrats goblin build with rituals (Battle Hymn, Brightstone Ritual, Seething Song, Mana Geyser), Sol Ring/Ruby Medallion, and Grapeshot/Empty the Warrens payoffs that can close games around turn six. It also holds the Storm-Kiln Artist + Haze of Rage two-card infinite at roughly six mana, which is above Bracket 3's four-mana floor but disqualifies it from Bracket 2. Card quality and synergy are high, yet it lacks the tutors, fast mana density, and free interaction of Bracket 4.

Combos the judge read: Krenko, Mob Boss + Skirk Prospector + Howlsquad Heavy; Krenko, Mob Boss + Skirk Prospector + Goblin Warchief; Krenko, Mob Boss + Skirk Prospector + Goblin Chieftain; Storm-Kiln Artist + Haze of Rage.

### 5. Floor 4, Saheeli, Radiant Creator, precon Living Energy

Judge: bracket 3. This is an energy/artifact theme deck with solid synergy, decent ramp (Sol Ring, signets, Aetherworks Marvel) and a real dual-land mana base, but wins are board-based and telegraphed rather than fast. It contains no marked Game Changers and no mass land denial, though the Lightning Runner + Aetherwind Basker extra-turn engine and related three-card loops push it past Core. Strong card quality plus a repeatable infinite-turn combo lands it comfortably in Upgraded rather than Optimized.

Combos the judge read: Saheeli, Radiant Creator + Lightning Runner + Aetherwind Basker; Saheeli, Radiant Creator + Lightning Runner + Peema Aether-Seer; Aethertide Whale + Nissa, Worldsoul Speaker + Panharmonicon; Lightning Runner + Aetherwind Basker.

### 6. Floor 1, Auntie Ool, Cursewretch, precon Blight Curse

Judge: bracket 2. This is a straightforward -1/-1 counter tribal deck with incremental, board-based wins and no marked Game Changers, no mass land denial, and no extra turns. Its only listed interactions (Blowfly Infestation with Flourishing Defenses or Hapatra) are flagged as not true two-card infinite combos, and the mana base is casual with just Sol Ring and a few signets as acceleration. Solid removal and card draw keep it above Exhibition, but it's unoptimized enough to sit comfortably in Core.

Combos the judge read: Blowfly Infestation + Flourishing Defenses; Blowfly Infestation + Hapatra, Vizier of Poisons.

### 7. Floor 4, Éowyn, Shieldmaiden, precon Riders of Rohan

Judge: bracket 3. This is a Rohan/Gondor humans-and-tokens go-wide deck with no Game Changers, modest ramp (Sol Ring, signets, talismans) and no real tutor package, but it has solid card quality, anthem synergy and interaction (Swords, Path, Supreme Verdict, Palace Jailer). The one listed two-card combo (Gilraen + Village Bell-Ringer) costs no extra mana, which pushes it past Bracket 2's restrictions. Its incremental, board-based, telegraphed wins land it comfortably in Upgraded rather than Optimized.

Combos the judge read: Gilraen, Dúnedain Protector + Village Bell-Ringer.

### 8. Floor 1, Omo, Queen of Vesuva, precon Tricky Terrain

Judge: bracket 3. This is a well-tuned lands-matter/Omo typal deck with high card quality and strong synergy \u2014 Dark Depths plus Thespian's Stage/Vesuva, Urza tron and locus lands turned on by Omo, Scute Swarm, Mana Reflection, and a real suite of interaction (Beast Within, Pongify, Oblivion Stone, Summary Dismissal) plus land tutors like Expedition Map and Sylvan Scrying. It carries no marked Game Changers, no mass land denial, no extra turns, and Spellbook finds no combo, so it doesn't push into Bracket 4. But the ramp density, Sol Ring, and land-toolbox payoffs make it clearly more powerful and less telegraphed than a Bracket 2 Core deck.

Combos the judge read: none, Commander Spellbook finds no combo.

### 9. Floor 1, Heroes in a Half Shell, precon Turtle Power

Judge: bracket 2. This is a straightforward five-color TMNT theme deck with no marked Game Changers, no combos, no extra-turn chaining and no mass land denial. Its mana base leans on taplands, Thriving lands and basics with only Sol Ring/Signet/Lantern for ramp, and wins come from incremental creature/counter synergies that are fully telegraphed on board. That unoptimized, theme-first build lands squarely in Core.

Combos the judge read: none, Commander Spellbook finds no combo.

### 10. Bracket 5, Sisay, Weatherlight Captain, cedh 2026-09-06, place 1 of 65

Judge: bracket 5. This is a textbook cEDH list: dual lands and fetches, a pile of fast mana (Lotus Petal, Moxen, LED, rituals, Ancient Tomb), a dozen-plus tutors, and free interaction like Force of Will, Pact of Negation, and Fierce Guardianship. It packs ~18 Game Changers and instant-win two-card combos in Thassa's Oracle with Demonic Consultation or Tainted Pact, plus Underworld Breach lines. It's built purely to win as fast as possible, not around a theme.

Combos the judge read: Faerie Mastermind + Smothering Tithe + Kinnan, Bonder Prodigy; Underworld Breach + Lotus Petal + Brain Freeze; Tainted Pact + Thassa's Oracle; Underworld Breach + Lion's Eye Diamond + Brain Freeze; Demonic Consultation + Thassa's Oracle.

### 11. Bracket 5, Kinnan, Bonder Prodigy, cedh 2026-09-06, place 1 of 40

Judge: bracket 5. This is a textbook cEDH Kinnan build: 12 Game Changers including Force of Will, Rhystic Study, Mana Vault, Gaea's Cradle and Thassa's Oracle, plus a dense fast-mana suite (Sol Ring, Moxen, Lotus Petal, Elvish Spirit Guide, Ancient Tomb) and free counterspells for protection. Numerous two-card infinite mana combos (Kinnan + Basalt Monolith, Bloom Tender/Incubation Druid + Freed from the Real) feed Thrasios/Thassa's Oracle wins, and Whir/Transmute/Chord tutors assemble them by turn 2-4. It is metagame-tuned and combo-focused rather than thematic, so it sits squarely in cEDH.

Combos the judge read: Hullbreaker Horror + Mox Amber; Kinnan, Bonder Prodigy + Delighted Halfling + Pemmin's Aura; Kinnan, Bonder Prodigy + Delighted Halfling + Freed from the Real; Badgermole Cub + Pemmin's Aura; Badgermole Cub + Freed from the Real; Devoted Druid + Agatha's Soul Cauldron + Incubation Druid; Hullbreaker Horror; Hullbreaker Horror + Sol Ring; Hullbreaker Horror + Mana Vault; Machine God's Effigy + Devoted Druid; Incubation Druid + Freed from the Real; Incubation Druid + Pemmin's Aura; Kinnan, Bonder Prodigy + Birds of Paradise + Pemmin's Aura; Kinnan, Bonder Prodigy + Birds of Paradise + Freed from the Real; Kinnan, Bonder Prodigy + Basalt Monolith; Kinnan, Bonder Prodigy + Bloom Tender + Freed from the Real; Bloom Tender + Pemmin's Aura; Bloom Tender + Freed from the Real.

### 12. Bracket 5, Grenzo, Dungeon Warden, cedh 2026-09-06, place 2 of 69

Judge: bracket 4. Seven marked Game Changers (Ancient Tomb, Chrome Mox, Demonic Tutor, Mana Vault, Mox Diamond, Opposition Agent, Orcish Bowmasters, Vampiric Tutor) far exceeds Bracket 3's cap of three, and it packs a free two-card infinite in Kiki-Jiki + Zealous Conscripts. Stacked fast mana (Sol Ring, Lotus Petal, Mox Opal/Amber, Simian Spirit Guide, City of Traitors, Gemstone Caverns), tutors, and multiple sacrifice/treasure loops make this a fast, optimized Grenzo combo deck that can win well before turn six.

Combos the judge read: Pitiless Plunderer + Umbral Collar Zealot + Stridehangar Automaton; Pitiless Plunderer + Ashnod's Altar + Stridehangar Automaton; Pitiless Plunderer + Woe Strider + Stridehangar Automaton; Pitiless Plunderer + Viscera Seer + Stridehangar Automaton; Pitiless Plunderer + Carrion Feeder + Stridehangar Automaton; Warren Soultrader + Zulaport Cutthroat + Stridehangar Automaton; Grenzo, Dungeon Warden + Workhorse + Hoverstone Pilgrim; Grenzo, Dungeon Warden + Workhorse + Barkform Harvester; Grenzo, Dungeon Warden + Orcish Bowmasters + Ashnod's Altar + Barkform Harvester; Grenzo, Dungeon Warden + Orcish Bowmasters + Ashnod's Altar + Hoverstone Pilgrim; Kiki-Jiki, Mirror Breaker + Zealous Conscripts.

### 13. Bracket 5, Flubs, the Fool, cedh 2026-09-06, place 4 of 65

Judge: bracket 5. This is a cEDH-caliber list: nine Game Changers, a full suite of fast mana (Lotus Petal, LED, Mox Diamond/Amber, Chrome Mox, spirit guides, Sol Ring), power tutors, and free-spell enablers built around Flubs' card flow. It packs multiple redundant infinite lines (Food Chain + Squee, Underworld Breach + LED + Brain Freeze/Wheel, Cloudstone Curio kobold loops) plus stack interaction like Pyroblast and Autumn's Veil. The deck is optimized to win in the first few turns with no concession to theme or casual play.

Combos the judge read: Cloudstone Curio + Crimson Kobolds + Rograkh, Son of Rohgahh; Cloudstone Curio + Crimson Kobolds + Kobolds of Kher Keep; Cloudstone Curio + Rograkh, Son of Rohgahh + Crookshank Kobolds; Cloudstone Curio + Crookshank Kobolds + Kobolds of Kher Keep; Cloudstone Curio + Crimson Kobolds + Crookshank Kobolds; Oboro Breezecaller + Talon Gates of Madara + Gaea's Cradle; Flubs, the Fool + Underworld Breach + Lion's Eye Diamond + Life from the Loam; Underworld Breach + Lotus Petal + Brain Freeze; Underworld Breach + Wheel of Fortune + Jeska's Will; Cloudstone Curio + Rograkh, Son of Rohgahh + Kobolds of Kher Keep; Squee, the Immortal + Food Chain; Underworld Breach + Lion's Eye Diamond + Wheel of Fortune; Underworld Breach + Lion's Eye Diamond + Brain Freeze.

### 14. Bracket 5, Heliod, the Radiant Dawn // Heliod, the Warped Eclipse, cedh 2026-09-05, place 1 of 45

Judge: bracket 5. This is a cEDH-grade list: Timetwister, Mox Diamond/Chrome Mox/Mox Opal/Lotus Petal/Mana Vault/Grim Monolith/Ancient Tomb fast mana, a dual-land/fetch-heavy optimized mana base, a dozen-plus free and cheap counterspells (Force of Will, Fierce Guardianship, Pact of Negation, Silence, Orim's Chant), and a deep tutor package. It carries roughly a dozen marked Game Changers and a Displacer Kitten/Paradoxical Outcome/Hullbreaker Horror engine that can lock the game out early. The strategy is metagame-driven stax-combo control with no theme concessions, well past Bracket 4 casual optimization.

Combos the judge read: Faerie Mastermind + Smothering Tithe + Mirrormade; Faerie Mastermind + Smothering Tithe + Clever Impersonator; Hullbreaker Horror; Hullbreaker Horror + Sol Ring; Hullbreaker Horror + Mana Vault.

### 15. Bracket 5, Crystal, Inhuman Princess, cedh 2026-09-05, place 1 of 64

Judge: bracket 5. This is a cEDH storm/combo list: dual lands and fetches, Chrome Mox/Mox Diamond/LED/Lotus Petal/spirit guides for fast mana, a dozen free or cheap counterspells (Force of Will, Pact, Fierce Guardianship, Daze, Misstep), heavy tutoring, and eleven marked Game Changers. It wins with Underworld Breach + LED/Lotus Petal + Brain Freeze or Jeskai Ascendancy/Vivi draw loops, capable of turn-one to turn-three kills. Nothing here is thematic — it's pure optimized competitive play.

Combos the judge read: Underworld Breach + Lotus Petal + Brain Freeze; Underworld Breach + Lion's Eye Diamond + Brain Freeze.

### 16. Bracket 5, Ghave, Guru of Spores, cedh 2026-09-05, place 1 of 37

Judge: bracket 2. This is essentially the Ghave C11 precon with light upgrades: slow, clunky cards (Storm Herd, Celestial Force, Alliance of Arms), a basic-heavy mana base with signets and vivid lands, and wins that build incrementally on board through tokens and counters. Only one Game Changer (Aura Shards) plus Sol Ring/Skullclamp, no mass land denial, no extra turns, and Spellbook finds no combos. That's squarely unoptimized Core-level play.

Combos the judge read: none, Commander Spellbook finds no combo.

### 17. Bracket 5, Magda, Brazen Outlaw, cedh 2026-09-05, place 1 of 64

Judge: bracket 5. This is a cEDH-style mono-red Magda combo deck: heavy fast mana (Sol Ring, Mana Vault, Mox Opal, Chrome Mox, Lotus Petal, Ancient Tomb, City of Traitors, Simian Spirit Guide, Jeska's Will) plus free interaction (Deflecting Swat, Pyrokinesis, Pyroblast, Red Elemental Blast) and hate pieces like Torpor Orb and Grafdigger's Cage. Multiple zero-mana two-card infinite combos with Clock of Omens off the commander mean it can win on turn two or three, and the build is strategy-first rather than a dwarf theme.

Combos the judge read: Magda, Brazen Outlaw + Battered Golem + Stalactite Dagger; Battered Golem + Dwarven Bloodboiler + Magda, Brazen Outlaw + Stalactite Dagger; Magda, Brazen Outlaw + Clock of Omens + Three Tree Mascot; Magda, Brazen Outlaw + Clock of Omens + Barkform Harvester; Magda, Brazen Outlaw + Battered Golem + Maskwood Nexus; Magda, Brazen Outlaw + Clock of Omens + Roaming Throne; Battered Golem + Dwarven Bloodboiler + Magda, Brazen Outlaw + Maskwood Nexus; Magda, Brazen Outlaw + Clock of Omens + Liquimetal Torque; Magda, Brazen Outlaw + Clock of Omens + Universal Automaton.

### 18. Bracket 5, K'rrik, Son of Yawgmoth, cedh 2026-09-05, place 1 of 39

Judge: bracket 5. This is a cEDH-caliber K'rrik list: Lion's Eye Diamond, Mox Diamond, Chrome Mox, Mana Vault, Grim Monolith, Ancient Tomb, City of Traitors and rituals fuel turn-two or three kills, backed by a full tutor suite (Imperial Seal, Vampiric, Demonic, Grim Tutor, Beseech the Mirror, Razaketh) and Necropotence. Ten-plus Game Changers, a premium dual/fetch mana base, and multiple compact reanimator/combo lines (Sheoldred + Vilis with K'rrik, Peer into the Abyss kills) place it far past Bracket 4's casual optimized range. The build is a tuned competitive engine rather than a theme deck.

Combos the judge read: Asmodeus the Archfiend + Necrotic Ooze + K'rrik, Son of Yawgmoth + Sheoldred, the Apocalypse; Sheoldred, the Apocalypse + Vilis, Broker of Blood + Necropotence; Sheoldred, the Apocalypse + Vilis, Broker of Blood + Blood Celebrant; Orcish Bowmasters + Peer into the Abyss; Sheoldred, the Apocalypse + Vilis, Broker of Blood + Skirge Familiar; Sheoldred, the Apocalypse + Vilis, Broker of Blood + K'rrik, Son of Yawgmoth; Sheoldred, the Apocalypse + Peer into the Abyss; Asmodeus the Archfiend + Necrotic Ooze + Skirge Familiar.

### 19. Bracket 5, Nick Fury, Agent of S.H.I.E.L.D., cedh 2026-09-05, place 1 of 63

Judge: bracket 5. This is a full cEDH storm/consultation list: Power-9-adjacent duals and fetches, heavy fast mana (Lion's Eye Diamond, Mana Vault, Moxen, rituals, Spirit Guides), a dozen-plus tutors, and free interaction like Force of Will, Fierce Guardianship, Pact of Negation and Mindbreak Trap. It wins with Thassa's Oracle plus Demonic Consultation/Tainted Pact or Underworld Breach storm lines, all for minimal mana and often by turns two or three. With ~18 Game Changers and no theme beyond winning fastest, it is metagame-tuned competitive EDH.

Combos the judge read: Faerie Mastermind + Smothering Tithe + Copy Enchantment; Underworld Breach + Lotus Petal + Brain Freeze; Tainted Pact + Thassa's Oracle; Underworld Breach + Lion's Eye Diamond + Brain Freeze; Demonic Consultation + Thassa's Oracle.

### 20. Bracket 5, Terra, Magical Adept // Esper Terra, cedh 2026-09-05, place 2 of 64

Judge: bracket 5. This is a cEDH-grade list: dual lands and original duals with full fast mana (Lotus Petal, Mox Diamond/Opal/Amber, Chrome Mox, Mana Vault, LED, City of Traitors, Gemstone Caverns), a huge free-interaction suite (Force of Will, Pact of Negation, Fierce Guardianship, Deflecting Swat, Subtlety, Silence/Orim's Chant), and roughly twenty Game Changers. It packs multiple zero-to-low-mana two-card infinites plus the Underworld Breach/LED/Brain Freeze storm kill and a deep tutor package to assemble them, so games can end on almost any turn. The build is a metagame-focused combo deck rather than a theme, placing it firmly in cEDH.

Combos the judge read: Terra, Magical Adept // Esper Terra + Mirrormade; Terra, Magical Adept // Esper Terra + Estrid's Invocation; Terra, Magical Adept // Esper Terra + Copy Enchantment; Terra, Magical Adept // Esper Terra + Spark Double; Faerie Mastermind + Smothering Tithe + Mirrormade; Faerie Mastermind + Smothering Tithe + Estrid's Invocation; Faerie Mastermind + Smothering Tithe + Copy Enchantment; Underworld Breach + Lotus Petal + Brain Freeze; Underworld Breach + Lion's Eye Diamond + Brain Freeze; Keen Sense + Niv-Mizzet, Parun.

### 21. Bracket 5, Shorikai, Genesis Engine, cedh 2026-09-05, place 2 of 60

Judge: bracket 5. This is a cEDH-caliber Shorikai list: 10 Game Changers, heavy fast mana (Mana Vault, Grim Monolith, Ancient Tomb, all the Moxen, Lotus Petal, Sol Ring), an efficient dual/fetch-heavy mana base, and a dense free-interaction suite (Force of Will, Pact of Negation, Fierce Guardianship, Flusterstorm, Silence/Orim's Chant). It packs tutors (Enlightened Tutor, Whir of Invention, Transmute Artifact, Tezzeret) into Displacer Kitten/Hullbreaker Horror storm loops that can win very early, with no theme or casual concessions.

Combos the judge read: Teferi, Time Raveler + Displacer Kitten + Grim Monolith; Faerie Mastermind + Smothering Tithe + Mirrormade; Faerie Mastermind + Smothering Tithe + Copy Enchantment; Faerie Mastermind + Smothering Tithe + Clever Impersonator; Hullbreaker Horror; Hullbreaker Horror + Sol Ring; Hullbreaker Horror + Mana Vault; The One Ring + Displacer Kitten + Teferi, Time Raveler; Teferi, Time Raveler + Displacer Kitten + Sol Ring; Teferi, Time Raveler + Displacer Kitten + Mox Opal.

