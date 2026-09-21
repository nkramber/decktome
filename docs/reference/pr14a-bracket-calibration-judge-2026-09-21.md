# PR-14A bracket gate, judge lane

Run date: 2026-09-21. Card snapshot: 2026-09-04. Decks read from `/Users/nate/repos/decktome/docs/reference/pr14a-bracket-calibration-decks-2026-09-13.md`.

Verdict: FAIL. The judge agreed with the bracket on 12 of 21 decks (57 percent, the bar is 80), with 0 judge errors. This document reads the judge bar alone: the block, band, and content bars are in the source document.

Calls 21. Cost $0.3113. Time 120 seconds.

## Run

- Suite `bracket-judge`, run `pr14a-bracket-calibration-judge-2026-09-21`, on 2026-09-21, commit `23b67ce`.
- Roles: judge on `claude-opus-5` (anthropic, effort medium).
- Versions: card snapshot 2026-09-04, bracket_judge prompt version 2, generate prompt version 15, source `pr14a-bracket-calibration-decks-2026-09-13.md`.
- Calls: 21. Cost: $0.3113. Time: 120 seconds.

| # | Built for | Judged | Agrees | Commander |
|---|---|---|---|---|
| 1 | 2 | 2 | yes | Captain America, Team Leader |
| 2 | 2 | 3 | no | Cloud, Ex-SOLDIER |
| 3 | 2 | 2 | yes | Esika, God of the Tree // The Prismatic Bridge |
| 4 | 2 | 3 | no | Zada, Hedron Grinder |
| 5 | 2 | 3 | no | Saheeli, Radiant Creator |
| 6 | 2 | 3 | no | Auntie Ool, Cursewretch |
| 7 | 2 | 3 | no | Éowyn, Shieldmaiden |
| 8 | 2 | 3 | no | Omo, Queen of Vesuva |
| 9 | 2 | 2 | yes | Heroes in a Half Shell |
| 10 | 5 | 5 | yes | Sisay, Weatherlight Captain |
| 11 | 5 | 5 | yes | Kinnan, Bonder Prodigy |
| 12 | 5 | 4 | no | Grenzo, Dungeon Warden |
| 13 | 5 | 5 | yes | Flubs, the Fool |
| 14 | 5 | 5 | yes | Heliod, the Radiant Dawn // Heliod, the Warped Eclipse |
| 15 | 5 | 5 | yes | Crystal, Inhuman Princess |
| 16 | 5 | 3 | no | Ghave, Guru of Spores |
| 17 | 5 | 4 | no | Magda, Brazen Outlaw |
| 18 | 5 | 5 | yes | K'rrik, Son of Yawgmoth |
| 19 | 5 | 5 | yes | Nick Fury, Agent of S.H.I.E.L.D. |
| 20 | 5 | 5 | yes | Terra, Magical Adept // Esper Terra |
| 21 | 5 | 5 | yes | Shorikai, Genesis Engine |

## Why

### 1. Bracket 2, Captain America, Team Leader, precon Avengers Assemble

Judge: bracket 2. This is a straightforward Avengers/hero tribal deck with anthem effects, card-draw payoffs, and a basic mana rocks package — no marked Game Changers, no combos, no mass land denial, and no extra turns. Interaction is limited to a handful of fair removal spells like Swords to Plowshares, Rip Apart, and Austere Command, and the clock is slow, relying on wide tribal boards. That puts it right around precon strength, i.e. Bracket 2 Core, with only mild upgrades like Sol Ring and Kindred Discovery.

Combos the judge read: none, Commander Spellbook finds no combo.

### 2. Bracket 2, Cloud, Ex-SOLDIER, precon Limit Break

Judge: bracket 3. This is an equipment/Voltron deck built well beyond precon level with efficient ramp (Sol Ring, Arcane Signet, Nature's Lore), a tuned dual-land mana base, and strong support pieces like Puresteel Paladin, Colossus Hammer, Lightning Greaves and Skullclamp, plus real interaction (Vandalblast, Chaos Warp, Austere Command, Vanquish the Horde). No Game Changers are marked and Spellbook finds no combos, so it isn't pushed into Optimized territory, but the commander-damage clock and protection suite are faster and more resilient than Core. That lands it squarely in Upgraded.

Combos the judge read: none, Commander Spellbook finds no combo.

### 3. Bracket 2, Esika, God of the Tree // The Prismatic Bridge, precon From Cute to Brute

Judge: bracket 2. This is a casual modal-DFC theme deck: five-color goodstuff creatures and planeswalkers with a slow mana base of taplands and Pathways, minimal ramp beyond Sol Ring and signets, and no tutors or fast mana. Spellbook finds no combos, there are no Game Changers marked, and the only marked power card is Sol Ring, so the clock is slow and grindy. The theme-first build and light interaction put it right around precon strength, Bracket 2.

Combos the judge read: none, Commander Spellbook finds no combo.

### 4. Bracket 2, Zada, Hedron Grinder, precon Goblin Storm

Judge: bracket 3. This is an upgraded Zada goblin-storm deck: no marked Game Changers, no mass land denial, and a basic-heavy mono-red mana base with only Sol Ring/Ruby Medallion and rituals for acceleration, so it isn't operating at optimized speed. It does pack a real engine (Mana Geyser, Past in Flames, Empty the Warrens, Grapeshot, Impact Tremors) plus the Storm-Kiln Artist + Haze of Rage two-card loop, but that costs roughly six mana and generally comes online around turn five or six rather than earlier. The Krenko/Skirk Prospector lines are three-card combos, keeping it comfortably in the Upgraded band rather than Bracket 4.

Combos the judge read: Krenko, Mob Boss + Skirk Prospector + Howlsquad Heavy; Krenko, Mob Boss + Skirk Prospector + Goblin Warchief; Krenko, Mob Boss + Skirk Prospector + Goblin Chieftain; Storm-Kiln Artist + Haze of Rage.

### 5. Bracket 2, Saheeli, Radiant Creator, precon Living Energy

Judge: bracket 3. No Game Changers and a fairly casual energy/artifact theme with basic ramp (Sol Ring, signets) and a durdly, high-curve mana base, but it does pack a genuine two-card infinite combo in Lightning Runner + Aetherwind Basker (plus Saheeli lines), which pushes it past Bracket 2. The pieces cost 5 and 7 mana, so it can't assemble cheaply in the first six turns, fitting Bracket 3's allowance. Interaction is present but modest (Disallow, Chaos Warp, board wipes), so it's an upgraded casual deck rather than an optimized one.

Combos the judge read: Saheeli, Radiant Creator + Lightning Runner + Aetherwind Basker; Saheeli, Radiant Creator + Lightning Runner + Peema Aether-Seer; Aethertide Whale + Nissa, Worldsoul Speaker + Panharmonicon; Lightning Runner + Aetherwind Basker.

### 6. Bracket 2, Auntie Ool, Cursewretch, precon Blight Curse

Judge: bracket 3. No marked Game Changers, no mass land denial, no extra turns, and the two listed synergies are explicitly not infinite two-card combos, so nothing pushes it past the Bracket 3 ceiling. Still, the card quality clearly exceeds a precon \u2014 Sol Ring, Ignoble Hierarch, Fire Covenant, Assassin's Trophy, Grave Titan, a tuned dual/triple-land base and a focused -1/-1 counter engine with plenty of efficient removal. That's a grindy, upgraded value deck that wins around turns six to eight rather than an optimized one.

Combos the judge read: Blowfly Infestation + Flourishing Defenses; Blowfly Infestation + Hapatra, Vizier of Poisons.

### 7. Bracket 2, Éowyn, Shieldmaiden, precon Riders of Rohan

Judge: bracket 3. This is a themed human/Rohan tribal go-wide deck with mostly on-theme creatures, basic ramp (Sol Ring, signets, talismans) and a modest interaction suite, well above precon polish but far from optimized. No marked Game Changers and no mass land denial or chained extra turns, so it fits comfortably under Bracket 4. The only listed two-card combo (Gilraen + Village Bell-Ringer) is just an infinite blink/untap loop that requires assembly and isn't a reliable early win, which is acceptable at Upgraded.

Combos the judge read: Gilraen, Dúnedain Protector + Village Bell-Ringer.

### 8. Bracket 2, Omo, Queen of Vesuva, precon Tricky Terrain

Judge: bracket 3. This is a tuned Simic landfall/land-type deck with no marked Game Changers, no mass land denial, and no combos found, so it stays legal for lower brackets — but its card quality, Sol Ring plus land tutors (Expedition Map, Sylvan Scrying, Hour of Promise), heavy card draw engines (Tatyova, Hydroid Krasis, Scute Swarm) and real interaction (Pongify, Beast Within, Oblivion Stone, Summary Dismissal) push it well beyond a precon. It wins through incremental value and big creatures rather than fast combo kills, so games run to the mid-game, matching Upgraded rather than Optimized.

Combos the judge read: none, Commander Spellbook finds no combo.

### 9. Bracket 2, Heroes in a Half Shell, precon Turtle Power

Judge: bracket 2. This is a five-color TMNT tribal/theme deck with no marked Game Changers, no mass land denial, no extra-turn chaining, and Commander Spellbook finds no combos. The mana base is largely taplands and basics with only Sol Ring/Signet/Lantern for ramp, and the interaction is slow removal and wraths rather than fast, efficient answers. That puts it right at preconstructed power, expecting long games — Bracket 2, Core.

Combos the judge read: none, Commander Spellbook finds no combo.

### 10. Bracket 5, Sisay, Weatherlight Captain, cedh 2026-09-06, place 1 of 65

Judge: bracket 5. This is a cEDH Sisay/5-color list: dual lands and fetches, heavy fast mana (Lion's Eye Diamond, Mana Vault, Moxen, rituals, Spirit Guides), a huge tutor package, and free counterspells like Pact of Negation, Fierce Guardianship, and Force of Will as interaction. It wins with Thassa's Oracle plus Demonic Consultation or Tainted Pact, backed by Underworld Breach/Brain Freeze, and packs ~18 Game Changers. It is a metagame-tuned combo deck, not a theme, so it can end the game on nearly any turn.

Combos the judge read: Faerie Mastermind + Smothering Tithe + Kinnan, Bonder Prodigy; Underworld Breach + Lotus Petal + Brain Freeze; Tainted Pact + Thassa's Oracle; Underworld Breach + Lion's Eye Diamond + Brain Freeze; Demonic Consultation + Thassa's Oracle.

### 11. Bracket 5, Kinnan, Bonder Prodigy, cedh 2026-09-06, place 1 of 40

Judge: bracket 5. This is a textbook cEDH Kinnan list: heavy free interaction (Force of Will, Fierce Guardianship, Mana Drain, Mindbreak Trap), a pile of fast mana (Sol Ring, Mana Vault, Mox Diamond/Opal/Amber, Chrome Mox, Lotus Petal, Ancient Tomb, Gaea's Cradle) and a dozen-plus Game Changers. It packs multiple two-card infinite mana combos (Kinnan + Basalt Monolith, untapper plus dork) closing into Thassa's Oracle/Thrasios for a turn-two-or-three win. Optimized tutor suite and metagame-tuned staples, not a theme, place it firmly at Bracket 5.

Combos the judge read: Hullbreaker Horror + Mox Amber; Kinnan, Bonder Prodigy + Delighted Halfling + Pemmin's Aura; Kinnan, Bonder Prodigy + Delighted Halfling + Freed from the Real; Badgermole Cub + Pemmin's Aura; Badgermole Cub + Freed from the Real; Devoted Druid + Agatha's Soul Cauldron + Incubation Druid; Hullbreaker Horror; Hullbreaker Horror + Sol Ring; Hullbreaker Horror + Mana Vault; Machine God's Effigy + Devoted Druid; Incubation Druid + Freed from the Real; Incubation Druid + Pemmin's Aura; Kinnan, Bonder Prodigy + Birds of Paradise + Pemmin's Aura; Kinnan, Bonder Prodigy + Birds of Paradise + Freed from the Real; Kinnan, Bonder Prodigy + Basalt Monolith; Kinnan, Bonder Prodigy + Bloom Tender + Freed from the Real; Bloom Tender + Pemmin's Aura; Bloom Tender + Freed from the Real.

### 12. Bracket 5, Grenzo, Dungeon Warden, cedh 2026-09-06, place 2 of 69

Judge: bracket 4. Eight marked Game Changers (Ancient Tomb, Chrome Mox, Demonic Tutor, Mana Vault, Mox Diamond, Opposition Agent, Orcish Bowmasters, Vampiric Tutor) far exceeds the three allowed in Bracket 3, and the Kiki-Jiki + Zealous Conscripts two-card infinite is castable early off heavy fast mana. The deck is packed with free/cheap acceleration (Sol Ring, Lotus Petal, Mox Opal/Amber, Simian Spirit Guide, City of Traitors, Gemstone Caverns), tutors, fetch lands, and sacrifice/treasure loops for fast, redundant kills. That optimized speed and combo density puts it at Bracket 4 rather than cEDH, since it's still a Grenzo tribal-ish engine deck rather than a pure metagame build.

Combos the judge read: Pitiless Plunderer + Umbral Collar Zealot + Stridehangar Automaton; Pitiless Plunderer + Ashnod's Altar + Stridehangar Automaton; Pitiless Plunderer + Woe Strider + Stridehangar Automaton; Pitiless Plunderer + Viscera Seer + Stridehangar Automaton; Pitiless Plunderer + Carrion Feeder + Stridehangar Automaton; Warren Soultrader + Zulaport Cutthroat + Stridehangar Automaton; Grenzo, Dungeon Warden + Workhorse + Hoverstone Pilgrim; Grenzo, Dungeon Warden + Workhorse + Barkform Harvester; Grenzo, Dungeon Warden + Orcish Bowmasters + Ashnod's Altar + Barkform Harvester; Grenzo, Dungeon Warden + Orcish Bowmasters + Ashnod's Altar + Hoverstone Pilgrim; Kiki-Jiki, Mirror Breaker + Zealous Conscripts.

### 13. Bracket 5, Flubs, the Fool, cedh 2026-09-06, place 4 of 65

Judge: bracket 5. This is a cEDH-caliber list: 10 Game Changers, heavy fast mana (Lion's Eye Diamond, Mox Diamond, Chrome Mox, Lotus Petal, Spirit Guides, Sol Ring), power-level duals and Gaea's Cradle, and a full tutor suite. It packs the Food Chain + Squee two-card infinite plus Underworld Breach/Brain Freeze storm kills, with free interaction (Pyroblast, Autumn's Veil, Snap) to protect them. The build is a tuned combo engine, not a theme, and can win in the first few turns.

Combos the judge read: Cloudstone Curio + Crimson Kobolds + Rograkh, Son of Rohgahh; Cloudstone Curio + Crimson Kobolds + Kobolds of Kher Keep; Cloudstone Curio + Rograkh, Son of Rohgahh + Crookshank Kobolds; Cloudstone Curio + Crookshank Kobolds + Kobolds of Kher Keep; Cloudstone Curio + Crimson Kobolds + Crookshank Kobolds; Oboro Breezecaller + Talon Gates of Madara + Gaea's Cradle; Flubs, the Fool + Underworld Breach + Lion's Eye Diamond + Life from the Loam; Underworld Breach + Lotus Petal + Brain Freeze; Underworld Breach + Wheel of Fortune + Jeska's Will; Cloudstone Curio + Rograkh, Son of Rohgahh + Kobolds of Kher Keep; Squee, the Immortal + Food Chain; Underworld Breach + Lion's Eye Diamond + Wheel of Fortune; Underworld Breach + Lion's Eye Diamond + Brain Freeze.

### 14. Bracket 5, Heliod, the Radiant Dawn // Heliod, the Warped Eclipse, cedh 2026-09-05, place 1 of 45

Judge: bracket 5. This is a cEDH-caliber Heliod flash/artifact-storm list: Timetwister, Mox Diamond, Chrome Mox, Mox Opal, Lotus Petal, Mana Vault, Grim Monolith, Ancient Tomb and City of Traitors alongside dual/fetch mana, plus a dozen-plus Game Changers. Free interaction (Force of Will, Fierce Guardianship, Pact of Negation, Force of Negation, Mindbreak Trap) and stacked tutors (Intuition, Gifts Ungiven, Enlightened/Mystical Tutor, Merchant Scroll, Spellseeker, Transmute Artifact) support fast Displacer Kitten/Paradoxical Outcome/Hullbreaker Horror engine wins. It is built as the best competitive strategy rather than a theme, so it belongs in bracket 5.

Combos the judge read: Faerie Mastermind + Smothering Tithe + Mirrormade; Faerie Mastermind + Smothering Tithe + Clever Impersonator; Hullbreaker Horror; Hullbreaker Horror + Sol Ring; Hullbreaker Horror + Mana Vault.

### 15. Bracket 5, Crystal, Inhuman Princess, cedh 2026-09-05, place 1 of 64

Judge: bracket 5. This is a cEDH storm/combo list: full power fast mana (Chrome Mox, Mox Diamond, LED, Lotus Petal, spirit guides), dual/original duals and fetches, ~12 Game Changers, heavy free interaction (Force of Will, Pact, Fierce Guardianship, Deflecting Swat), and a Underworld Breach + LED/Petal + Brain Freeze win. It's a metagame-tuned, no-theme competitive deck that can win on very early turns.

Combos the judge read: Underworld Breach + Lotus Petal + Brain Freeze; Underworld Breach + Lion's Eye Diamond + Brain Freeze.

### 16. Bracket 5, Ghave, Guru of Spores, cedh 2026-09-05, place 1 of 37

Judge: bracket 3. This is essentially the Commander 2011 'Devour for Power' precon list \u2014 clunky mana, slow high-cost finishers like Storm Herd and Celestial Force, and no tutors, fast mana beyond Sol Ring, or listed combos. However, it does run one Game Changer (Aura Shards), which pushes it out of Bracket 2's zero-Game-Changer requirement. So it lands at the floor of Bracket 3 on technicality while playing at a Core-level speed.

Combos the judge read: none, Commander Spellbook finds no combo.

### 17. Bracket 5, Magda, Brazen Outlaw, cedh 2026-09-05, place 1 of 64

Judge: bracket 4. This is a highly optimized mono-red Magda combo deck with heavy fast mana (Sol Ring, Mana Vault, Mox Opal, Chrome Mox, Lotus Petal, Rite of Flame, Simian Spirit Guide, Ancient Tomb, City of Traitors, Jeska's Will) plus tutor-adjacent Urza's Saga and Magda's own treasure-to-artifact search. It has multiple cheap two-card infinite combos with the commander (Clock of Omens plus any of several changelings/artifacts) that assemble well before turn six, alongside free interaction like Deflecting Swat, Pyroblast, and Pyrokinesis and stax pieces (Torpor Orb, Grafdigger's Cage, Vexing Bauble). That speed and redundancy is Optimized, though it's still a dwarf/artifact theme rather than a cEDH-tuned list.

Combos the judge read: Magda, Brazen Outlaw + Battered Golem + Stalactite Dagger; Battered Golem + Dwarven Bloodboiler + Magda, Brazen Outlaw + Stalactite Dagger; Magda, Brazen Outlaw + Clock of Omens + Three Tree Mascot; Magda, Brazen Outlaw + Clock of Omens + Barkform Harvester; Magda, Brazen Outlaw + Battered Golem + Maskwood Nexus; Magda, Brazen Outlaw + Clock of Omens + Roaming Throne; Battered Golem + Dwarven Bloodboiler + Magda, Brazen Outlaw + Maskwood Nexus; Magda, Brazen Outlaw + Clock of Omens + Liquimetal Torque; Magda, Brazen Outlaw + Clock of Omens + Universal Automaton.

### 18. Bracket 5, K'rrik, Son of Yawgmoth, cedh 2026-09-05, place 1 of 39

Judge: bracket 5. This is a mono-black K'rrik storm/reanimator shell with a dozen Game Changers (Ancient Tomb, Necropotence, LED, Mana Vault, Grim Monolith, Mox Diamond, Chrome Mox, Imperial Seal, Vampiric/Demonic Tutor, Bowmasters), dual lands/fetches, and a full suite of fast mana and rituals. It tutors aggressively into cheap two-card kills (Sheoldred/Vilis with K'rrik, Peer into the Abyss lines) with Beseech the Mirror, Entomb/Reanimate, and Yawgmoth's Will payoffs. That's a tuned, metagame-level competitive build rather than a themed optimized deck, so it plays at cEDH.

Combos the judge read: Asmodeus the Archfiend + Necrotic Ooze + K'rrik, Son of Yawgmoth + Sheoldred, the Apocalypse; Sheoldred, the Apocalypse + Vilis, Broker of Blood + Necropotence; Sheoldred, the Apocalypse + Vilis, Broker of Blood + Blood Celebrant; Orcish Bowmasters + Peer into the Abyss; Sheoldred, the Apocalypse + Vilis, Broker of Blood + Skirge Familiar; Sheoldred, the Apocalypse + Vilis, Broker of Blood + K'rrik, Son of Yawgmoth; Sheoldred, the Apocalypse + Peer into the Abyss; Asmodeus the Archfiend + Necrotic Ooze + Skirge Familiar.

### 19. Bracket 5, Nick Fury, Agent of S.H.I.E.L.D., cedh 2026-09-05, place 1 of 63

Judge: bracket 5. This is a textbook cEDH storm/consultation shell: dual lands and fetches, Lion's Eye Diamond, Mana Vault, Mox Diamond/Opal/Amber, rituals and Spirit Guides, plus a dozen-plus tutors and a full suite of free interaction (Force of Will, Fierce Guardianship, Pact of Negation, Mindbreak Trap). It packs roughly twenty Game Changers and multiple compact wins in Thassa's Oracle with Demonic Consultation or Tainted Pact, backed by Underworld Breach + LED + Brain Freeze. The deck is optimized for winning as early as possible and playing the best strategy rather than a theme, which is Bracket 5.

Combos the judge read: Faerie Mastermind + Smothering Tithe + Copy Enchantment; Underworld Breach + Lotus Petal + Brain Freeze; Tainted Pact + Thassa's Oracle; Underworld Breach + Lion's Eye Diamond + Brain Freeze; Demonic Consultation + Thassa's Oracle.

### 20. Bracket 5, Terra, Magical Adept // Esper Terra, cedh 2026-09-05, place 2 of 64

Judge: bracket 5. This is a cEDH list: dual lands and fetches, a full fast-mana suite (Lotus Petal, Mox Diamond/Opal/Amber, Chrome Mox, Mana Vault, LED, Ancient Tomb, City of Traitors, spirit guides), free interaction (Force of Will, Fierce Guardianship, Pact of Negation, Force of Negation, Subtlety, Deflecting Swat), and a deep tutor package. It packs over a dozen marked Game Changers plus multiple two-card commander combos and the Underworld Breach/Brain Freeze storm line. Speed, redundancy, and stax pieces like Opposition Agent and Drannith Magistrate mark it as metagame-focused competitive play, not a theme deck.

Combos the judge read: Terra, Magical Adept // Esper Terra + Mirrormade; Terra, Magical Adept // Esper Terra + Estrid's Invocation; Terra, Magical Adept // Esper Terra + Copy Enchantment; Terra, Magical Adept // Esper Terra + Spark Double; Faerie Mastermind + Smothering Tithe + Mirrormade; Faerie Mastermind + Smothering Tithe + Estrid's Invocation; Faerie Mastermind + Smothering Tithe + Copy Enchantment; Underworld Breach + Lotus Petal + Brain Freeze; Underworld Breach + Lion's Eye Diamond + Brain Freeze; Keen Sense + Niv-Mizzet, Parun.

### 21. Bracket 5, Shorikai, Genesis Engine, cedh 2026-09-05, place 2 of 60

Judge: bracket 5. This is a cEDH-caliber Shorikai list: duals/fetches, Ancient Tomb, City of Traitors, Mox Diamond, Chrome Mox, Mox Opal, Lotus Petal, Mana Vault, Grim Monolith, plus Enlightened Tutor, Transmute Artifact and Whir of Invention. It carries ten Game Changers and a dense free-interaction suite (Force of Will, Fierce Guardianship, Pact of Negation, Mindbreak Trap, Silence/Orim's Chant) alongside Displacer Kitten/Teferi and Dramatic Reversal engines that can win out of nowhere. That combination of fast mana, tutors, stack protection and combo finish is competitive metagame play, not a theme deck.

Combos the judge read: Teferi, Time Raveler + Displacer Kitten + Grim Monolith; Faerie Mastermind + Smothering Tithe + Mirrormade; Faerie Mastermind + Smothering Tithe + Copy Enchantment; Faerie Mastermind + Smothering Tithe + Clever Impersonator; Hullbreaker Horror; Hullbreaker Horror + Sol Ring; Hullbreaker Horror + Mana Vault; The One Ring + Displacer Kitten + Teferi, Time Raveler; Teferi, Time Raveler + Displacer Kitten + Sol Ring; Teferi, Time Raveler + Displacer Kitten + Mox Opal.

