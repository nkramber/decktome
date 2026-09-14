# PR-14A bracket gate, judge lane

Run date: 2026-09-13. Card snapshot: 2026-09-04. Decks read from `/Users/nate/Repos/decktome/docs/reference/pr14a-bracket-calibration-decks-2026-09-13.md`.

Verdict: FAIL. The judge agreed with the bracket on 16 of 21 decks (76 percent, the bar is 80), with 0 judge errors. This document reads the judge bar alone: the block, band, and content bars are in the source document.

Calls 21. Cost $0.3283. Time 149 seconds.

## Run

- Suite `bracket-judge`, run `pr14a-bracket-calibration-judge-2026-09-13`, on 2026-09-13, commit `9c8aa3a`.
- Roles: judge on `claude-opus-5` (anthropic, effort medium).
- Versions: card snapshot 2026-09-04, generate prompt version 12, source `pr14a-bracket-calibration-decks-2026-09-13.md`.
- Calls: 21. Cost: $0.3283. Time: 149 seconds.

| # | Built for | Judged | Agrees | Commander |
|---|---|---|---|---|
| 1 | 2 | 2 | yes | Captain America, Team Leader |
| 2 | 2 | 3 | no | Cloud, Ex-SOLDIER |
| 3 | 2 | 2 | yes | Esika, God of the Tree // The Prismatic Bridge |
| 4 | 2 | 3 | no | Zada, Hedron Grinder |
| 5 | 2 | 2 | yes | Saheeli, Radiant Creator |
| 6 | 2 | 3 | no | Auntie Ool, Cursewretch |
| 7 | 2 | 2 | yes | Éowyn, Shieldmaiden |
| 8 | 2 | 3 | no | Omo, Queen of Vesuva |
| 9 | 2 | 2 | yes | Heroes in a Half Shell |
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

### 1. Bracket 2, Captain America, Team Leader, precon Avengers Assemble

Judge: bracket 2. This is a straightforward Avengers/hero tribal deck built around typal payoffs (Door of Destinies, Kindred Discovery, Metallic Mimic, Herald's Horn) with no marked Game Changers, no mass land denial, no extra turns, and no two-card infinite combos. Ramp is modest rock-based (Sol Ring, Talismans, Signet) and interaction is a light suite of Swords to Plowshares, Rip Apart, and a couple of wraths, so games will play out over many turns. That puts it squarely at precon-plus power, Bracket 2, with only slight upgrade pressure toward 3.

### 2. Bracket 2, Cloud, Ex-SOLDIER, precon Limit Break

Judge: bracket 3. This is a well-tuned Cloud equipment/voltron build that goes clearly beyond precon power: a real dual/fastland-heavy mana base, Sol Ring and Arcane Signet, tutor-like value engines (Sword of the Animist, Armory Automaton, Puresteel Paladin with Colossus Hammer for fast one-shot kills), plus solid interaction in Vandalblast, Chaos Warp, Austere Command and Vanquish the Horde. It carries no marked Game Changers, no mass land denial, no extra-turn chaining and no cheap two-card infinite, so it stays under the Bracket 4 ceiling. Its consistent turn-5-6 commander-damage clock is too strong for Core, making Upgraded the right home.

### 3. Bracket 2, Esika, God of the Tree // The Prismatic Bridge, precon From Cute to Brute

Judge: bracket 2. This is a casual double-faced-card/legendary theme pile: the top end is clunky (Triplicate Titan, Zetalpa, Meteor Golem, Sphinx of the Second Sun) and the payoffs are grindy value rather than fast wins. No cards are flagged as Game Changers, there are no two-card infinite combos, no mass land denial, and no extra turns; interaction is a handful of removal spells and Sol Ring is the only real fast mana. The Pathway-heavy mana base is nicer than a precon's, but the gameplan is slow enough that games will routinely run past turn eight, putting it squarely in Core with only mild upgrade leanings.

### 4. Bracket 2, Zada, Hedron Grinder, precon Goblin Storm

Judge: bracket 3. No marked Game Changers and no mass land denial, but the deck is clearly souped up past precon level: rituals (Battle Hymn, Brightstone Ritual, Seething Song, Mana Geyser), Sol Ring/Ruby Medallion, Skullclamp, and a Zada-fueled storm/token engine with Impact Tremors, Goblin Bombardment, Grapeshot and Past in Flames that can win out of nowhere around turn six or so. Its combo pieces (Conspicuous Snoop lines, Pashalik Mons + Bombardment) generally need three cards or more, and the interaction suite is thin, so it sits comfortably in Upgraded rather than Optimized.

### 5. Bracket 2, Saheeli, Radiant Creator, precon Living Energy

Judge: bracket 2. This is a themed energy/artifact-token deck with no marked Game Changers, no fast mana beyond Sol Ring, no tutors, and no two-card infinite combos. Its payoffs (Aetherworks Marvel, Panharmonicon, Thopter Spy Network, big energy beaters) are slow, mana-hungry value plays, and interaction is a handful of counterspells and board wipes. It plays at roughly upgraded-precon power with a grindy eight-plus turn clock, placing it squarely in Core.

### 6. Bracket 2, Auntie Ool, Cursewretch, precon Blight Curse

Judge: bracket 3. This is a focused Jund -1/-1 counters attrition deck with no marked Game Changers, no mass land denial, no extra turns, and no true two-card infinite (Devoted Druid's untaps are capped by its own toughness, and payoffs like Flourishing Defenses only net a couple of tokens). Card quality and mana are clearly above precon level u2014 Sol Ring, Ignoble Hierarch, Fire Covenant, Glissa Sunslayer, Massacre Girl, Grave Titan, plus a deep suite of duals and efficient removal u2014 but the plan is grindy board-control that wins over many turns. That combination of upgraded consistency with fair, slow win conditions sits squarely in Bracket 3.

### 7. Bracket 2, Éowyn, Shieldmaiden, precon Riders of Rohan

Judge: bracket 2. This is a LOTR-flavored Human/token go-wide tribal deck with no marked Game Changers, no mass land denial, and no two-card infinite combos — its only extra-turn-ish effect is extra combat steps, not chained turns. The mana base is basics-plus-taplands with only Sol Ring and signets for acceleration, and the win route is slow anthem-and-swarm beatdown backed by a few clean removal spells. That casual precon-plus-a-few-upgrades power level lands it squarely in Core, with only a light lean toward Upgraded.

### 8. Bracket 2, Omo, Queen of Vesuva, precon Tricky Terrain

Judge: bracket 3. This is a tuned Simic lands/landfall deck with clear upgrades past precon level fffff
interaction (Beast Within, Pongify, Summary Dismissal), heavy land tutoring and Urzatron/Cloudpost ramp fffff
but it lists no marked Game Changers and no fast mana beyond Sol Ring. Its payoff is grindy value (Scute Swarm, Avenger, Tatyova, Dark Depths + Thespian's Stage/Vesuva for a 20/20), which wins around turn six to eight rather than presenting a cheap two-card infinite. That places it comfortably in Upgraded, above a precon but short of Optimized.

### 9. Bracket 2, Heroes in a Half Shell, precon Turtle Power

Judge: bracket 2. This is a five-color TMNT theme deck built almost entirely out of the new precon-style cards, with no marked Game Changers, no fast mana beyond Sol Ring, and no tutors or two-card infinite combos. The mana base leans on taplands, Thriving lands and basics, and the interaction is slow removal and wraths like Blasphemous Act and Vanquish the Horde. It plays a grindy counters/mutant tribal game that wins around turn eight or later, squarely Core.

### 10. Bracket 5, Sisay, Weatherlight Captain, cedh 2026-09-06, place 1 of 65

Judge: bracket 5. This is a textbook cEDH list: dual lands and fetches, Sol Ring/Moxen/rituals for turn-one explosiveness, a dozen-plus tutors, and a free-spell interaction suite (Force of Will, Fierce Guardianship, Pact of Negation, Mindbreak Trap). It wins with Thassa's Oracle plus Demonic Consultation/Tainted Pact and has Underworld Breach/Brain Freeze and Ad Nauseam lines as backup. With roughly twenty Game Changers and zero concession to theme, it's purely competitive.

### 11. Bracket 5, Kinnan, Bonder Prodigy, cedh 2026-09-06, place 1 of 40

Judge: bracket 5. This is a textbook cEDH Kinnan list: dual lands, fetches, a pile of fast mana (Mana Vault, Basalt Monolith, Moxen, Lotus Petal, Ancient Tomb, Gaea's Cradle) and free interaction like Force of Will, Fierce Guardianship, Mindbreak Trap and Mana Drain. It packs many two-card infinite combos (Kinnan + Basalt Monolith, Devoted Druid + Freed from the Real/Pemmin's Aura, Dramatic Reversal + Isochron-style mana loops) with Thassa's Oracle and Thrasios as win conditions. A dozen-plus Game Changers and a metagame-tuned, non-thematic build put it squarely at competitive level.

### 12. Bracket 5, Grenzo, Dungeon Warden, cedh 2026-09-06, place 2 of 69

Judge: bracket 5. This is a cEDH-style Grenzo combo engine: heavy fast mana (Sol Ring, Mana Vault, Chrome Mox, Mox Opal/Amber/Diamond, Lotus Petal, City of Traitors, Ancient Tomb, Simian Spirit Guide), fetch/dual mana base, and unrestricted tutors (Demonic, Vampiric, Goblin Recruiter, Kuldotha Forgemaster). It packs eight-plus marked Game Changers plus compact infinite loops (Kiki-Jiki + Zealous Conscripts, Ashnod's Altar/Pitiless Plunderer sac-treasure engines with Zulaport Cutthroat/Sling-Gang) and stax-style hate like Magus of the Moon, Opposition Agent and Deflecting Swat. That power level and turn-three-or-four kill potential is far past Bracket 4's casual optimized ceiling and firmly competitive.

### 13. Bracket 5, Flubs, the Fool, cedh 2026-09-06, place 4 of 65

Judge: bracket 5. This is a cEDH Flubs storm/combo list: dual lands, fetches, and a pile of fast mana (Lotus Petal, LED, Mox Diamond, Chrome Mox, Mox Amber, spirit guides, Jeska's Will, Sol Ring) enabling turn-one-to-three kills. It packs nine+ marked Game Changers plus Underworld Breach/LED and Food Chain/Cloudstone Curio infinite lines with Brain Freeze/Walking Ballista as wincons, backed by free interaction (Pyroblast, Autumn's Veil, Into the Flood Maw). It's a metagame-tuned best-strategy build, not a theme deck.

### 14. Bracket 5, Heliod, the Radiant Dawn // Heliod, the Warped Eclipse, cedh 2026-09-05, place 1 of 45

Judge: bracket 5. This is a cEDH-caliber Azorius build: Power-9 level fast mana (Timetwister, Mana Vault, Grim Monolith, Mox Diamond/Opal/Chrome Mox, Lotus Petal, Ancient Tomb, City of Traitors), a dual/fetch-heavy optimal mana base, and a dense free-interaction suite (Force of Will, Force of Negation, Fierce Guardianship, Pact of Negation, Mindbreak Trap, Subtlety, Silence/Orim's Chant). It packs roughly a dozen Game Changers plus a tutor package (Intuition, Gifts Ungiven, Mystical/Enlightened Tutor, Spellseeker, Merchant Scroll, Transmute Artifact) feeding Displacer Kitten/Paradoxical Outcome artifact-storm lines that can win very early. Nothing here is themed or self-limiting u2014 it's a metagame-facing competitive list.

### 15. Bracket 5, Crystal, Inhuman Princess, cedh 2026-09-05, place 1 of 64

Judge: bracket 5. This is a cEDH storm/combo list: duals and fetches, a pile of fast mana (Lotus Petal, LED, Moxen, Spirit Guides, Jeska's Will), free interaction (Force of Will, Fierce Guardianship, Pact of Negation, Misdirection), and a dozen-plus Game Changers. It wins early with Jeskai Ascendancy/Underworld Breach loops and Curiosity-style draw engines on Crystal, backed by stax-lite protection (Silence, Orim's Chant, Grand Abolisher). Nothing here is thematic — it's a tuned, metagame-facing combo deck that can win on turn two or three.

### 16. Bracket 5, Ghave, Guru of Spores, cedh 2026-09-05, place 1 of 37

Judge: bracket 3. This is essentially a lightly tuned Commander 2011 Ghave precon 
inc
large clunky token payoffs, no fast mana beyond Sol Ring, few tutors, and a slow tapland-heavy mana base, so its actual gameplay speed is Core level. However, it runs Aura Shards, a marked Game Changer, which disqualifies it from Bracket 2 outright. With one Game Changer, no mass land denial, and no cheap two-card infinite (the Spike Feeder/Ghave loop costs mana per iteration and is not a quick kill), it sits at the very bottom of Bracket 3.

### 17. Bracket 5, Magda, Brazen Outlaw, cedh 2026-09-05, place 1 of 64

Judge: bracket 5. This is the known mono-red Magda combo shell: heavy fast mana (Ancient Tomb, Mana Vault, Mox Opal, Chrome Mox, Lotus Petal, City of Traitors, Simian Spirit Guide, Rite of Flame, Jeska's Will), Gemstone Caverns/Urza's Saga, and free or one-mana interaction like Deflecting Swat, Pyrokinesis, Pyroblast, Flare of Duplication and Return the Favor. Magda turns dwarves/treasures into artifact tutors that assemble compact loops (Clock of Omens, Grinding Station, Sculpting Steel, Maskwood Nexus/Universal Automaton) for very early wins, backed by stax pieces (Torpor Orb, Grafdigger's Cage, Vexing Bauble, Disruptor Flute) that read as metagame answers. It is playing the strongest available strategy rather than a dwarf theme, so it belongs in cEDH.

### 18. Bracket 5, K'rrik, Son of Yawgmoth, cedh 2026-09-05, place 1 of 39

Judge: bracket 5. This is a mono-black K'rrik storm/reanimator list with a dozen-plus Game Changers (Necropotence, Mana Vault, Grim Monolith, LED, Mox Diamond, Imperial Seal, Demonic/Vampiric Tutor, Ancient Tomb, Bowmasters, etc.), stacked fast mana rituals, and a full tutor suite that assembles lines like Entomb/Reanimate into Razaketh or Necrotic Ooze combos. The mana base (duals, fetches, City of Traitors, Gemstone Caverns, Urza's Saga) and free interaction like Deadly Rollick are competitive-grade, not thematic. It can win on very early turns and is metagame-tuned, placing it squarely in cEDH.

### 19. Bracket 5, Nick Fury, Agent of S.H.I.E.L.D., cedh 2026-09-05, place 1 of 63

Judge: bracket 5. This is a full cEDH storm/combo list: dual lands and fetches, Lion's Eye Diamond, Mox Diamond/Opal/Amber/Chrome Mox, rituals, and roughly twenty Game Changers including Necropotence, Ad Nauseam, and Rhystic Study. It packs Thassa's Oracle with Demonic Consultation/Tainted Pact plus Underworld Breach-Brain Freeze lines, backed by free interaction like Force of Will, Fierce Guardianship, and Pact of Negation. It is metagame-optimized rather than themed and can win on turn one to three.

### 20. Bracket 5, Terra, Magical Adept // Esper Terra, cedh 2026-09-05, place 2 of 64

Judge: bracket 5. This is a full cEDH build: original dual lands and fetches, a stack of fast mana (Lion's Eye Diamond, Mana Vault, Moxen, Lotus Petal, spirit guides, Ancient Tomb, City of Traitors) and a deep tutor suite backed by free interaction like Force of Will, Fierce Guardianship, Pact of Negation, Subtlety and Mindbreak Trap. It packs roughly twenty Game Changers plus the Underworld Breach + LED + Brain Freeze storm kill, so it can win on very early turns. Nothing about it is themed or casual — it's a metagame-tuned competitive list.

### 21. Bracket 5, Shorikai, Genesis Engine, cedh 2026-09-05, place 2 of 60

Judge: bracket 5. This is a fully optimized cEDH Shorikai artifact-combo list: Ancient Tomb, Mana Vault, Grim Monolith, Mox Diamond, Chrome Mox, Mox Opal, Lotus Petal and City of Traitors alongside a Tundra/dual-and-fetch mana base, plus tutors like Enlightened Tutor, Transmute Artifact and Whir of Invention. It carries a full free-spell counterwall (Force of Will, Fierce Guardianship, Pact of Negation, Force of Negation, Subtlety, Misdirection) and soft-lock pieces (Silence, Orim's Chant, Teferi, Time Raveler) to protect wins via Dramatic Reversal/Displacer Kitten/Unwinding Clock untap loops with Shorikai. Roughly ten Game Changers and a competitive, metagame-tuned build put it squarely in Bracket 5.

