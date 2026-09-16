# PR-8 deck gate

Run date: 2026-09-16. Card snapshot: 2026-09-04.

Verdict: PASS. 25 of 25 decks passed every block check, 0 invented names reached the user, and 0 summaries stated a false rule of the game. All three bars are zero tolerance.

## Summary

| Measure | Value |
|---|---|
| Prompts | 25 |
| Decks returned | 25 |
| Decks with no block finding | 25 |
| Invented names that reached the user | 0 |
| Decks that needed the repair turn | 2 |
| Summaries judged (F-26) | 25 |
| Summaries that state a rule of the game | 0 |
| Summaries that state a FALSE rule | 0 |
| Judge errors | 0 |
| Case assertions missed (PR-28b) | 0 |
| Decks the plan judge read (PR-15, information) | 25 |
| Mean plan score, 0 to 1 | 0.69 |
| Plan reasons the judge left empty | 3 |
| Errors | 0 |
| Prompt version | 15 |
| Calls | 77 |
| Cost | $2.7527 |
| Time | 2375 seconds |

## Run

- Suite `decks`, run `pr8-deck-gate-run28`, on 2026-09-16, commit `6ff0244`.
- Roles: generate on `gpt-5.6-terra` (openai, effort medium), judge on `claude-opus-5` (anthropic, effort medium), repair on `gpt-5.6-terra` (openai, effort medium).
- Versions: card snapshot 2026-09-04, generate prompt version 15, plan_rubric prompt version 2, precons `5.3.0+20260903`, quality_model `20260914T154223Z`.
- Calls: 77. Cost: $2.7527. Time: 2375 seconds.

## Findings by code

A block stops the deck. A warning and a note are reports.

| Code | Count |
|---|---|
| `curve_summary` | 25 |
| `not_owned` | 17 |
| `profile_off_band` | 10 |
| `mana_pass` | 9 |
| `outside_requested_set` | 2 |
| `finisher_short` | 2 |
| `bracket_cut` | 1 |
| `basics_added` | 1 |
| `cards_trimmed` | 1 |

By severity: BLOCK 0. WARN 31. INFO 37. 

## The set filter (PR-17B)

A deck asked for a set holds cards of that set family alone. Three things earn an exception, and each one is marked: a card the reader named (D-381), a mana card the reader allowed from outside (D-382), and a basic land, which no set limit filters (D-378).

| # | Prompt | Sets | In set | Fill | Deck cards outside | Marked |
|---|---|---|---|---|---|---|
| 19 | the Hobbit family, two colours | `hob,hoc` | 170 | 0 | 0 | 0 |
| 20 | the Hobbit family, a delegated commander | `hob,hoc` | 75 | 0 | 0 | 0 |
| 21 | the Hobbit family, mana from outside | `hob,hoc` | 127 | 26 | 12 | 12 |
| 22 | a set family and a card from outside it | `hob,hoc` | 170 | 0 | 1 | 1 |
| 23 | two set families at once | `blb,blc,hob,hoc,pblb` | 201 | 0 | 0 | 0 |
| 24 | a 60-card deck from one set | `blb,blc,pblb` | 95 | 0 | 0 | 0 |

## The precon exclusion (PR-24)

A deck asked to use no card of a precon holds none of its cards. The products' copies leave the owned counts, so a card with a spare copy in the binder stays usable, and a basic land never leaves (D-408, D-37). An excluded card in the deck is a block.

| # | Prompt | Products | Cards excluded | Cards spare | Excluded cards in the deck | Blocked |
|---|---|---|---|---|---|---|
| 25 | use no card of an owned precon | Avengers Assemble | 57 | 27 | 0 | 0 |

## Decks

### 1. lifegain Commander, any card

Format: Commander. Theme: lifegain. Pool: any_card. Shortlist: 306 names.

Commander: Karlov of the Ghost Council.

Grade: bad, score 0.03, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $777.19 to buy, $777.19 the whole deck.

**Summary:** Karlov of the Ghost Council leads a white-black lifegain deck that develops a board of scaling creatures and resilient threats while keeping mana and cards flowing. It closes through dedicated lifegain finishers such as Angel of Destiny, Felidar Sovereign, and Test of Endurance, with Archangel of Thune chief among the pressure pieces. The deck gives up explosive speed for a steadier creature-based plan backed by removal, protective interaction, and board resets.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every slot feeds or exploits lifegain — Karlov counter growth, Ajani's Pridemate/Voice of the Blessed/Bloodbond Vampire scaling, Heliod/Archangel/Righteous Valkyrie payoffs, Vito/Cliffhaven drains, and lifelink-doubling with Rhox Faithmender — with only a few filler artifacts (Colossal Plow, Druidic Satchel) sitting outside the plan.
- PLAN theme_fit=yes: It is exactly the requested white-black Karlov of the Ghost Council lifegain Commander deck, and its power level — no tutor-combo lines, no Heliod/Ballista kill, mostly fair creatures and removal — sits comfortably at bracket 3.
- PLAN useful_as_built=yes: 37 lands with a full suite of duals and painlands in a two-color deck, a low creature curve, ample draw, removal, and multiple win routes (Karlov beats, Vito/Cliffhaven drain, Angel of Destiny, Felidar Sovereign, Test of Endurance) make it playable as it stands, if a bit light on real ramp.
- PLAN summary_honest=yes: Placeholder.
- [INFO] `curve_summary`: average mana value 3.08 over 63 nonland cards

<details><summary>The deck list</summary>

- 9 Plains | land | Provides reliable white mana for the deck.
- 8 Swamp | land | Provides reliable black mana for the deck.
- 1 Brightclimb Pathway // Grimclimb Pathway | land | Flexible access to either deck color.
- 1 Caves of Koilos | land | An untapped dual-color land.
- 1 City of Brass | land | Flexible access to the deck's colors.
- 1 Command Tower | land | A dependable multicolor land for Karlov's deck.
- 1 Concealed Courtyard | land | An efficient white-black land.
- 1 Eclipsed Steppe | land | Adds white-black mana access.
- 1 Godless Shrine | land | An efficient white-black land.
- 1 Isolated Chapel | land | Adds white-black mana access.
- 1 Mana Confluence | land | Flexible access to the deck's colors.
- 1 Marsh Flats | land | Improves access to the deck's basic land colors.
- 1 Prismatic Vista | land | Improves access to the deck's basic land colors.
- 1 Restless Fortress | land | A white-black land that also carries finisher potential.
- 1 Scrubland | land | An efficient white-black land.
- 1 Shattered Sanctum | land | An efficient white-black land.
- 1 Shineshadow Snarl | land | Adds white-black mana access.
- 1 Silent Clearing | land | An efficient white-black land.
- 1 Tarnished Citadel | land | Flexible access to the deck's colors.
- 1 Turbulent Moor | land | Adds white-black mana access.
- 1 Vault of Champions | land | An efficient white-black land.
- 1 Altar of the Pantheon | ramp | Supports the deck's mana development.
- 1 Colossal Plow | ramp | Provides early mana acceleration.
- 1 Druidic Satchel | ramp | Supports the deck's mana development.
- 1 Legion's Landing // Adanto, the First Fort | ramp | A low-cost ramp piece for the deck.
- 1 Nuka-Cola Vending Machine | ramp | Supports the deck's mana development.
- 1 Orazca Relic | ramp | Supports the deck's mana development.
- 1 Phial of Galadriel | ramp | Supports the deck's mana development.
- 1 Potioner's Trove | ramp | Supports the deck's mana development.
- 1 Pristine Talisman | ramp | Provides mana while fitting the lifegain plan.
- 1 The Celestus | ramp | Supports the deck's mana development.
- 1 Archivist of Oghma | draw | A low-cost draw option.
- 1 Convalescent Care | draw | Provides draw for the lifegain plan.
- 1 Dawn of Hope | draw | Provides draw for the lifegain plan.
- 1 Enduring Innocence | draw | Provides draw while contributing to the board.
- 1 Frodo, Adventurous Hobbit | draw | A low-cost draw option.
- 1 Markov Purifier | draw | Provides draw for the lifegain plan.
- 1 Pearl-Ear, Imperial Advisor | draw | A low-cost draw option.
- 1 Scheming Silvertongue // Sign in Blood | draw | A flexible draw card.
- 1 Survival Cache | draw | Provides draw for the lifegain plan.
- 1 Tymna the Weaver | draw | Provides repeatable draw support.
- 1 Vampiric Rites | draw | A low-cost draw outlet.
- 1 Alseid of Life's Bounty | interaction | A low-cost protective interaction piece.
- 1 Caduceus, Staff of Hermes | interaction | Adds flexible interaction to the deck.
- 1 Courageous Resolve | interaction | Protects the deck's key pieces.
- 1 Faith's Shield | interaction | A low-cost protective interaction piece.
- 1 Metropolis Reformer | interaction | Provides interactive support on a creature.
- 1 Paladin Danse, Steel Maverick | interaction | Adds flexible interaction to the deck.
- 1 Restoration Magic | interaction | Protects the deck's key pieces.
- 1 Werefox Bodyguard | interaction | Provides interaction while developing the board.
- 1 Ayli, Eternal Pilgrim | removal | A low-cost removal option.
- 1 Consuming Corruption | removal | A focused removal spell.
- 1 Foolish Fate | removal | A focused removal spell.
- 1 Gumdrop Poisoner // Tempt with Treats | removal | Flexible removal that fits the deck's colors.
- 1 Murderous Rider // Swift End | removal | Flexible removal on a creature card.
- 1 Nightmare's Thirst | removal | A low-cost removal spell.
- 1 Poisoner's Apprentice | removal | Provides creature-based removal.
- 1 Starseer Mentor | removal | Provides creature-based removal.
- 1 Umezawa's Jitte | removal | Repeatable removal utility.
- 1 Ajani, Strength of the Pride | wipe | A wipe that also fits the lifegain theme.
- 1 Kaya's Wrath | wipe | A reliable full-board reset.
- 1 White Sun's Twilight | wipe | A flexible wipe for contested boards.
- 1 Ajani's Pridemate | synergy | A core lifegain synergy creature.
- 1 Cleric Class | synergy | A low-cost lifegain synergy piece.
- 1 Heliod, Sun-Crowned | synergy | A central lifegain synergy permanent.
- 1 Resplendent Angel | synergy | A lifegain synergy threat.
- 1 Righteous Valkyrie | synergy | A lifegain synergy creature.
- 1 Vito, Thorn of the Dusk Rose | synergy | A key lifegain synergy piece.
- 1 Voice of the Blessed | synergy | A scalable lifegain synergy creature.
- 1 Archangel of Thune | threat | A powerful lifegain-focused threat.
- 1 Attended Healer | threat | A board-building lifegain threat.
- 1 Blood Baron of Vizkopa | threat | A substantial vampire threat.
- 1 Bloodbond Vampire | threat | A lifegain-focused creature threat.
- 1 Celestine, the Living Saint | threat | A resilient legendary threat.
- 1 Cliffhaven Vampire | threat | A lifegain-focused vampire threat.
- 1 Divinity of Pride | threat | A large lifegain-focused threat.
- 1 Elenda, Saint of Dusk | threat | A legendary vampire threat.
- 1 Gideon's Company | threat | A lifegain-focused creature threat.
- 1 Nykthos Paragon | threat | A high-impact lifegain threat.
- 1 Rhox Faithmender | threat | A durable lifegain-focused threat.
- 1 Valkyrie Harbinger | threat | A high-end angel threat.
- 1 Angel of Destiny | wincon | A dedicated lifegain finisher.
- 1 Felidar Sovereign | wincon | A dedicated lifegain finisher.
- 1 Test of Endurance | wincon | A dedicated lifegain finisher.

</details>

### 2. aristocrats Commander, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 181 names.

Commander: Denethor, Ruling Steward.

Grade: typical, score 0.41, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $145.36 the whole deck.

**Summary:** This Orzhov aristocrats deck develops a creature-heavy board, uses its sacrifice-focused synergy pieces to turn that board into value, and keeps resources moving with a deep draw package. It protects key pieces, answers troublesome permanents, and resets crowded boards when needed before closing with Grave Venerations, Archfiend of Ifnir, or Al Bhed Salvagers. The deck gives up some threat density in exchange for stronger card flow, flexible answers, and a dependable mana base.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=no: The list is mostly removal, board wipes, and equipment/protection for a single creature, which pulls directly against the creature-heavy sacrifice board the summary describes, and there is no coherent sacrifice-outlet-plus-drain-payoff chain.
- PLAN theme_fit=no: The colors, format, and commander match an Orzhov build, but the requested aristocrats engine of tokens, sac outlets, and death-trigger drain is essentially absent in favor of a removal-and-protection midrange pile.
- PLAN useful_as_built=partly: The mana base and curve are sound and the abundant removal lets it play real games, but with only three vague finishers and no engine it lacks a dependable way to close.
- PLAN summary_honest=no: It claims a creature-heavy board and sacrifice-focused synergy pieces when the deck runs roughly a dozen creatures, no real sacrifice outlets or Blood Artist-style payoffs, and three sweepers that undo its own board.
- [INFO] `curve_summary`: average mana value 2.73 over 63 nonland cards

<details><summary>The deck list</summary>

- 14 Plains | land | Basic white mana source for the two-color base.
- 12 Swamp | land | Basic black mana source for the two-color base.
- 1 Command Tower | land | Flexible colored mana source.
- 1 City of Brass | land | Flexible colored mana source.
- 1 Exotic Orchard | land | Flexible colored mana source.
- 1 Marsh Flats | land | Fetch land for the mana base.
- 1 Fabled Passage | land | Fetch land for the mana base.
- 1 Evolving Wilds | land | Fixes either basic color.
- 1 Terramorphic Expanse | land | Fixes either basic color.
- 1 Grand Coliseum | land | Additional colored mana coverage.
- 1 Path of Ancestry | land | Colored mana source for the creature-focused deck.
- 1 Secluded Courtyard | land | Colored mana support for the creature suite.
- 1 Call of the Ring | draw | Black draw piece that keeps cards flowing.
- 1 Cirith Ungol Patrol | draw | Creature-based draw for the deck's board plan.
- 1 Exemplar of Light | draw | Creature-based draw that contributes to the board.
- 1 Folk Hero | draw | Repeatable draw support for the creature package.
- 1 Idol of Oblivion | draw | Low-cost artifact draw option.
- 1 Inspiring Overseer | draw | Creature-based draw for maintaining resources.
- 1 Lembas | draw | Compact artifact draw support.
- 1 Mask of Memory | draw | Equipment-based draw for attacking creatures.
- 1 Night's Whisper | draw | Efficient black card flow.
- 1 Painful Truths | draw | Additional black draw support.
- 1 Puresteel Paladin | draw | Draw support that works with the equipment package.
- 1 Skullclamp | draw | Efficient draw outlet for disposable creatures.
- 1 Tome of Legends | draw | Persistent artifact draw option.
- 1 Wall of Omens | draw | Early defensive creature that replaces itself.
- 1 Boromir, Warden of the Tower | interaction | Creature-based protection and disruption.
- 1 Champion's Helm | interaction | Protects an important legendary creature.
- 1 Clever Concealment | interaction | Protects the developed board from opposing answers.
- 1 Darksteel Plate | interaction | Durable protection for a key creature.
- 1 Gift of Immortality | interaction | Helps preserve a valuable creature.
- 1 Lightning Greaves | interaction | Low-cost protection for key creatures.
- 1 Reprieve | interaction | Flexible instant-speed disruption.
- 1 Sheltered by Ghosts | interaction | Protection piece for an important permanent.
- 1 Swiftfoot Boots | interaction | Protects a key creature while keeping it active.
- 1 Take Up the Shield | interaction | Combat and removal protection at instant speed.
- 1 Unbreakable Formation | interaction | Protects the creature board at a critical moment.
- 1 Ultimate Magic: Holy | interaction | Additional instant-speed interaction.
- 1 Bitter Triumph | removal | Flexible single-target answer.
- 1 Claim the Precious | removal | Black targeted removal.
- 1 Crib Swap | removal | Creature answer that fits the color base.
- 1 Dismember | removal | Low-cost creature removal option.
- 1 Dispatch | removal | Efficient targeted removal.
- 1 Fatal Push | removal | Cheap black removal for early threats.
- 1 Generous Gift | removal | Broad permanent removal.
- 1 Get Lost | removal | Versatile white removal.
- 1 Infernal Grasp | removal | Reliable creature removal.
- 1 Path to Exile | removal | Efficient targeted creature answer.
- 1 Swords to Plowshares | removal | Efficient targeted creature answer.
- 1 Austere Command | wipe | Flexible reset when the board gets away from you.
- 1 Fumigate | wipe | Creature-board reset for difficult positions.
- 1 Vanquish the Horde | wipe | Efficient creature-board reset.
- 1 Arcane Signet | ramp | Reliable early color fixing.
- 1 Bender's Waterskin | ramp | Artifact ramp for reaching the commander and larger spells.
- 1 Chromatic Lantern | ramp | Mana fixing and ramp for the two-color base.
- 1 Commander's Sphere | ramp | Reliable mana rock with later utility.
- 1 Fellwar Stone | ramp | Cheap mana acceleration.
- 1 Lotho, Corrupt Shirriff | ramp | Creature-based mana acceleration.
- 1 Sol Ring | ramp | Fast colorless acceleration.
- 1 Springleaf Drum | ramp | Low-cost ramp that uses the creature board.
- 1 Thought Vessel | ramp | Mana acceleration with lasting utility.
- 1 Wayfarer's Bauble | ramp | Early land-based mana acceleration.
- 1 Arcade Cabinet | synergy | Supports the deck's sacrifice-focused synergy package.
- 1 Gollum the Abandoned | synergy | Creature synergy piece for the aristocrats plan.
- 1 Gollum, Patient Plotter | synergy | Creature synergy piece for the aristocrats plan.
- 1 Gríma Wormtongue | synergy | Supports the deck's creature synergy plan.
- 1 Heirloom Auntie | synergy | Creature synergy piece for the aristocrats shell.
- 1 Phantom Train | synergy | Artifact synergy piece for the deck's main plan.
- 1 Bill the Pony | threat | Low-cost creature threat that advances the board.
- 1 Hei Bai, Spirit of Balance | threat | Creature threat that adds battlefield pressure.
- 1 Namazu Trader | threat | Creature threat that contributes to the board plan.
- 1 Vengeful Villagers | threat | Creature threat for applying pressure.
- 1 Al Bhed Salvagers | wincon | Finishing threat and one of the deck's closing cards.
- 1 Archfiend of Ifnir | wincon | Finishing threat that can close a stalled game.
- 1 Grave Venerations | wincon | Finishing enchantment for closing games.

</details>

### 3. artifacts Commander, bracket 4

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 320 names.

Commander: Urza, Lord High Artificer.

Grade: good, score 0.68, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $7589.16 to buy, $7589.16 the whole deck.

**Summary:** Urza drives a fast artifact engine built around cheap mana, sustained card flow, and artifact tutors that locate the most important piece for the moment. The deck protects its development with a deep interaction suite, clears obstacles with flexible removal, and can close through its dedicated alternate finishers or a concentrated wave of artifact-creature pressure, chief among them Kappa Cannoneer and Cyberdrive Awakener. It gives up some individual card flexibility in favor of a highly artifact-dependent plan, so its strongest turns come from establishing a dense board early and defending it.

The quality model grades this deck good against the top lists of the format: the deck makes more mana on turn four than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

- PLAN plan_coherent=partly: The bulk of the list coheres as a mono-blue artifact engine with cheap mana rocks, affinity-style draw, tutors and artifact threats, but a few pieces pull elsewhere "Thassa's Oracle with no self-mill or Demonic Consultation/Tainted Pact, Aetherflux Reservoir listed as removal with little storm support, Mechanized Production and Mirrodin Besieged as slow alt-wincons "so the 'alternate finishers' strand is half-built.
- PLAN theme_fit=yes: It is exactly a high-power Urza, Lord High Artificer artifact deck in blue with fast mana, Workshop/Saga lands and free counterspells, sitting plausibly at bracket 4.
- PLAN useful_as_built=yes: 24 Islands plus utility lands and a pile of cheap rocks give a fine mono-blue mana base, the curve is low, and there are enough real threats (Kappa Cannoneer, Cyberdrive Awakener, Urza tokens/construct beats) to actually close a game.
- PLAN summary_honest=partly: Most claims (cheap mana, card flow, tutors, counterspell suite, Kappa Cannoneer/Cyberdrive Awakener as finishers) are backed by the list, but 'dedicated alternate finishers' overstates Thassa's Oracle, which has no enabler here to win with, and the Aetherflux line is unsupported.
- [INFO] `curve_summary`: average mana value 3.00 over 66 nonland cards
- [INFO] `mana_pass`: the builder moved 1 card of the mana base to bring the deck inside its power level
- [INFO] `basics_added`: the list was 1 card short, so the builder added 1 basic land

<details><summary>The deck list</summary>

- 24 Island | land | Provides a deep base of blue mana for the deck's blue spells.
- 1 Academy Ruins | land | Provides a utility land slot for the artifact-focused plan.
- 1 Inventors' Fair | land | Provides an artifact-focused utility land slot.
- 1 Mishra's Workshop | land | Provides an exceptionally strong land-based artifact mana source.
- 1 Mystic Sanctuary | land | Provides a blue source in a utility land slot.
- 1 Otawara, Soaring City | land | Provides a blue source in a flexible utility land slot.
- 1 Sink into Stupor // Soporific Springs | land | Provides a blue land slot with flexibility.
- 1 Glimmervoid | land | Provides a blue source for the artifact-heavy mana base.
- 1 Urza's Saga | land | Provides an artifact-focused utility land slot.
- 1 Urza's Workshop | land | Provides a land-based artifact mana source.
- 1 Chrome Mox | ramp | Supplies immediate acceleration for explosive starts.
- 1 Lion's Eye Diamond | ramp | Supplies fast artifact mana for high-power turns.
- 1 Lotus Petal | ramp | Supplies immediate mana at no mana cost.
- 1 Mana Vault | ramp | Supplies a major burst of early mana.
- 1 Moonsnare Prototype | ramp | Supplies cheap artifact-based mana acceleration.
- 1 Mox Amber | ramp | Supplies fast mana alongside the commander and legendary permanents.
- 1 Mox Diamond | ramp | Supplies immediate mana acceleration.
- 1 Mox Opal | ramp | Supplies fast mana in the artifact-heavy shell.
- 1 Sol Ring | ramp | Supplies efficient early acceleration.
- 1 Springleaf Drum | ramp | Turns the creature suite into additional mana access.
- 1 Metalworker | ramp | Provides powerful artifact-based mana production.
- 1 Krark-Clan Ironworks | ramp | Converts artifacts into large bursts of mana.
- 1 Grand Architect | ramp | Provides additional mana support for the artifact plan.
- 1 Rhystic Study | draw | Provides a premier source of sustained card advantage.
- 1 Sai, Master Thopterist | draw | Supplies artifact-linked card advantage.
- 1 Thoughtcast | draw | Provides efficient card draw in an artifact-dense deck.
- 1 Thirst for Knowledge | draw | Provides efficient card selection and card advantage.
- 1 Reverse Engineer | draw | Provides a substantial artifact-supported draw spell.
- 1 Forensic Gadgeteer | draw | Provides card advantage within the artifact shell.
- 1 Riddlesmith | draw | Provides repeatable filtering as artifacts are deployed.
- 1 Nexus of Becoming | draw | Provides an artifact-based card-advantage piece.
- 1 One with the Machine | draw | Provides a large draw effect tied to artifact size.
- 1 Vedalken Archmage | draw | Provides artifact-linked card advantage.
- 1 Thopter Spy Network | draw | Provides an artifact-focused long-game draw engine.
- 1 An Offer You Can't Refuse | interaction | Provides extremely efficient protection against key opposing spells.
- 1 Assert Authority | interaction | Provides artifact-scaled stack interaction.
- 1 Disruption Protocol | interaction | Provides efficient interaction in an artifact-heavy board.
- 1 Fierce Guardianship | interaction | Provides premium protection once Urza is established.
- 1 Force of Will | interaction | Provides vital zero-mana stack interaction.
- 1 Ice Out | interaction | Provides flexible counterspell interaction.
- 1 Metallic Rebuke | interaction | Provides efficient artifact-supported stack interaction.
- 1 Override | interaction | Provides an additional protective counterspell.
- 1 Padeem, Consul of Innovation | interaction | Helps protect the deck's important artifacts.
- 1 Stoic Rebuttal | interaction | Provides reliable artifact-supported stack interaction.
- 1 Welding Jar | interaction | Provides cheap protection for crucial artifacts.
- 1 Aether Spellbomb | removal | Provides a cheap artifact-based answer to creatures.
- 1 Aetherflux Reservoir | removal | Provides artifact-based removal reach.
- 1 Contagion Clasp | removal | Provides a compact artifact removal option.
- 1 Cyber Conversion | removal | Provides flexible single-target removal.
- 1 Into Thin Air | removal | Provides a versatile answer to a problematic permanent.
- 1 Kitesail Larcenist | removal | Provides creature-based permanent disruption.
- 1 Ravenform | removal | Provides a clean answer to troublesome targets.
- 1 Resculpt | removal | Provides efficient flexible removal.
- 1 Skysovereign, Consul Flagship | removal | Provides repeatable artifact-based removal pressure.
- 1 Spine of Ish Sah | removal | Provides a broad artifact-based answer to permanents.
- 1 Transmute Artifact | synergy | Finds a needed artifact and converts expendable material into it.
- 1 Whir of Invention | synergy | Finds key artifacts at instant speed within the artifact shell.
- 1 Reshape | synergy | Turns a lesser artifact into a needed artifact tutor target.
- 1 Kappa Cannoneer | threat | Provides a formidable artifact-based combat threat.
- 1 Cyberdrive Awakener | threat | Turns the artifact board into a major source of pressure.
- 1 Karn, Scion of Urza | threat | Provides a resilient artifact-themed threat.
- 1 Master Transmuter | threat | Provides a dangerous artifact creature that pressures the table.
- 1 Arcbound Crusher | threat | Provides an artifact creature threat that benefits from the deck's density.
- 1 Argent Sphinx | threat | Provides an evasive artifact-supported creature threat.
- 1 Dross Scorpion | threat | Provides an artifact creature that adds board pressure.
- 1 Ironheart, Clever Champion | threat | Provides a compact legendary artifact creature threat.
- 1 Lodestone Golem | threat | Provides a disruptive artifact creature threat.
- 1 Phyrexian Metamorph | threat | Provides a flexible artifact creature threat.
- 1 Traxos, Scourge of Kroog | threat | Provides an efficient heavy-hitting artifact threat.
- 1 Whirler Rogue | threat | Provides an artifact-linked creature threat.
- 1 Thassa's Oracle | wincon | Provides a compact alternate finishing route.
- 1 Mechanized Production | wincon | Provides an artifact-based alternate finishing route.
- 1 Mirrodin Besieged | wincon | Provides an artifact-themed alternate finishing route.
- 1 Engineered Explosives | wipe | Provides a scalable reset for problematic boards.
- 1 Hurkyl's Recall | wipe | Provides a sweeping answer to artifact-heavy opposing boards.
- 1 Waterbender Ascension | draw | the mana pass added it to bring the mana base inside the power level

</details>

### 4. dinosaur tribal, bracket 2

Format: Commander. Theme: dinosaurs. Pool: any_card. Shortlist: 270 names.

Commander: Gishath, Sun's Avatar.

Grade: bad, score 0.03, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $1167.19 to buy, $1167.19 the whole deck.

**Summary:** Gishath, Sun's Avatar leads a Dinosaur-focused creature deck that develops its mana, keeps cards flowing, and fills the board with threats. The deck aims to win through sustained combat pressure, backed by several large finishing Dinosaurs and combat-oriented finishers. It gives up some speed for a straightforward plan, a sturdy mana base, and a broad mix of creature-based answers.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves one color far better than another, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Every slot supports one clear plan: ramp into Gishath and other big Dinosaurs, cheat more Dinosaurs off combat damage, and win by attacking, with Dinosaur cost-reducers, card draw off creatures, and creature-based removal all pointing the same way.
- PLAN theme_fit=partly: It is unmistakably the Gishath Dinosaur deck asked for, but the power level drifts above a bracket 2 build for a new player: original dual lands, fetches, Sylvan Library (a Game Changer that by definition pushes past bracket 2), and fiddly modal cards like Akroma's Will and Boros Charm.
- PLAN useful_as_built=yes: 37 lands plus ten ramp pieces, a curve that tops out sensibly, plenty of draw, removal, board wipes, and many ways to close a game mean it can be shuffled up and played as printed, with only red sources running a bit thin for a three-color commander.
- PLAN summary_honest=yes: The summary describes exactly what the list does — creature-based combat pressure with ramp, draw, and big finishers — and even volunteers the mana-color imbalance and lack of speed rather than hiding them.
- [INFO] `curve_summary`: average mana value 3.66 over 62 nonland cards
- [INFO] `mana_pass`: the builder moved 3 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 4 Forest | land | Provides a reliable green source for the mana base.
- 1 Plains | land | Provides a reliable white source for the mana base.
- 7 Forest | land | Provides reliable green sources for the mana base.
- 5 Plains | land | Provides reliable white sources for the mana base.
- 2 Mountain | land | Provides reliable red sources for the mana base.
- 1 Command Tower | land | Provides flexible mana for the three-color deck.
- 1 City of Brass | land | Provides flexible mana for the three-color deck.
- 1 Mana Confluence | land | Provides flexible mana for the three-color deck.
- 1 Bountiful Promenade | land | Supports the green-white portion of the mana base.
- 1 Spire Garden | land | Supports the green-red portion of the mana base.
- 1 Spectator Seating | land | Supports the red-white portion of the mana base.
- 1 Temple Garden | land | Supports the green-white portion of the mana base.
- 1 Sacred Foundry | land | Supports the red-white portion of the mana base.
- 1 Stomping Ground | land | Supports the green-red portion of the mana base.
- 1 Savannah | land | Supports the green-white portion of the mana base.
- 1 Plateau | land | Supports the red-white portion of the mana base.
- 1 Taiga | land | Supports the green-red portion of the mana base.
- 1 Battlefield Forge | land | Supports the red-white portion of the mana base.
- 1 Brushland | land | Supports the green-white portion of the mana base.
- 1 Karplusan Forest | land | Supports the green-red portion of the mana base.
- 1 Arid Mesa | land | Helps assemble the colored mana base.
- 1 Windswept Heath | land | Helps assemble the colored mana base.
- 1 Wooded Foothills | land | Helps assemble the colored mana base.
- 1 Arcane Signet | ramp | Provides straightforward color fixing and mana development.
- 1 Sol Ring | ramp | Accelerates the deck's larger Dinosaur plays.
- 1 Birds of Paradise | ramp | Provides early mana fixing.
- 1 Llanowar Elves | ramp | Provides early green mana acceleration.
- 1 Drover of the Mighty | ramp | Provides mana acceleration for the creature plan.
- 1 Intrepid Paleontologist | ramp | Supports mana development alongside Dinosaur cards.
- 1 Ixalli's Lorekeeper | ramp | Provides early mana acceleration.
- 1 Thunderherd Migration | ramp | Builds the mana base for later turns.
- 1 Topiary Stomper | ramp | Develops mana while fitting the creature theme.
- 1 Hulking Raptor | ramp | Supplies additional mana from a Dinosaur body.
- 1 Beast Whisperer | draw | Provides repeatable draw for the creature-heavy deck.
- 1 Garruk's Uprising | draw | Adds draw support for the large-creature plan.
- 1 Guardian Project | draw | Provides ongoing draw as creatures join the deck's board.
- 1 Harmonize | draw | Offers simple, direct card draw.
- 1 Return of the Wildspeaker | draw | Provides a flexible draw option for a creature deck.
- 1 Ripjaw Raptor | draw | Adds Dinosaur-themed card draw.
- 1 Rishkar's Expertise | draw | Provides a substantial draw effect for the large-creature plan.
- 1 Runic Armasaur | draw | Adds a Dinosaur body with draw support.
- 1 Sylvan Library | draw | Improves access to useful cards over the game.
- 1 Toski, Bearer of Secrets | draw | Rewards the deck's combat-focused board with draw.
- 1 Earthshaker Dreadmaw | draw | Provides draw from a Dinosaur threat.
- 1 Heroic Intervention | interaction | Protects the board from opposing interaction.
- 1 Lightning Greaves | interaction | Protects an important creature in a simple package.
- 1 Swiftfoot Boots | interaction | Provides another accessible protection piece.
- 1 Temple Altisaur | interaction | Adds Dinosaur-themed defensive interaction.
- 1 Akroma's Will | interaction | Protects the creature board and serves as a finishing combat option.
- 1 Boros Charm | interaction | Provides flexible protection and a finishing option.
- 1 Itzquinth, Firstborn of Gishath | removal | Provides efficient Dinosaur-themed removal.
- 1 Savage Stomp | removal | Offers on-theme creature removal.
- 1 Thrashing Brontodon | removal | Brings removal on a Dinosaur body.
- 1 Ravenous Sailback | removal | Adds a Dinosaur-based removal option.
- 1 Tranquil Frillback | removal | Provides versatile removal from a Dinosaur.
- 1 Trumpeting Carnosaur | removal | Adds removal attached to a Dinosaur threat.
- 1 Raging Regisaur | removal | Provides Dinosaur-themed removal while advancing combat pressure.
- 1 Kinjalli's Caller | synergy | Supports the deck's Dinosaur-focused plan.
- 1 Otepec Huntmaster | synergy | Supports the deck's Dinosaur-focused plan.
- 1 Marauding Raptor | synergy | Provides a Dinosaur-focused synergy piece.
- 1 Hunting Velociraptor | synergy | Supports the deck's Dinosaur creature plan.
- 1 Dinosaur Egg | synergy | Adds a thematic Dinosaur synergy card.
- 1 Commune with Dinosaurs | synergy | Improves access to the deck's Dinosaur plan.
- 1 Huatli's Raptor | synergy | Adds an early Dinosaur synergy creature.
- 1 Nest Robber | synergy | Adds an early Dinosaur creature for the tribal plan.
- 1 Sunfrill Imitator | synergy | Adds a Dinosaur-themed synergy creature.
- 1 Regal Imperiosaur | synergy | Rewards the deck's Dinosaur focus.
- 1 Harnessed Snubhorn | threat | Provides an early Dinosaur threat.
- 1 Majestic Heliopterus | threat | Adds a substantial Dinosaur threat.
- 1 Regisaur Alpha | threat | Provides a powerful Dinosaur threat for the board.
- 1 Pantlaza, Sun-Favored | threat | Adds a strong Dinosaur threat.
- 1 Ghalta and Mavren | threat | Provides a large Dinosaur threat.
- 1 Quartzwood Crasher | threat | Adds a combat-focused Dinosaur threat.
- 1 Shifting Ceratops | threat | Provides a resilient finishing Dinosaur threat.
- 1 Carnage Tyrant | threat | Provides a major Dinosaur threat.
- 1 Goring Ceratops | threat | Adds a large Dinosaur finisher.
- 1 Rampaging Ceratops | threat | Provides a Dinosaur finisher for combat turns.
- 1 Tyrranax Rex | threat | Adds a powerful Dinosaur threat.
- 1 Zetalpa, Primal Dawn | threat | Provides a large Dinosaur threat for closing games.
- 1 Huatli, Poet of Unity // Roar of the Fifth People | wincon | Provides a thematic route to ending the game.
- 1 Etali, Primal Storm | wincon | Provides a high-impact Dinosaur win condition.
- 1 Dinosaurs on a Spaceship | wincon | Adds a Dinosaur-themed finishing card.
- 1 Forerunner of the Empire | wipe | Provides a Dinosaur-adjacent board wipe option.
- 1 Raging Swordtooth | wipe | Adds a Dinosaur-based board wipe option.
- 1 Vanquish the Horde | wipe | Provides a straightforward reset when the board gets crowded.

</details>

### 5. blink Commander, owned first

Format: Commander. Theme: blink. Pool: owned_first. Shortlist: 169 names.

Commander: Gilraen, Dúnedain Protector.

Grade: baseline, score 0.34, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: yes, for copy_limit. Block findings: 0.

Cost: $0.00 to buy, $155.01 the whole deck.

**Summary:** A mono-white blink deck built around repeatedly reusing enter-the-battlefield value, protecting its centerpiece, and maintaining a broad suite of efficient answers. It develops through artifacts, Angels, Humans, and equipment, then converts accumulated board presence into decisive evasive combat.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

This deck plan aims for 2 or more finishers, the cards that win the game, and the deck holds 1.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core is sound — equipment/aura-style targeting to trigger Gilraen's blink plus a pile of ETB creatures and flicker effects — but several inclusions pull against it, most glaringly Kataki, War's Wage in a deck running roughly twenty artifacts and an equipment-matters card like Puresteel Paladin, along with a set of low-impact draw trinkets that do little for the blink plan.
- PLAN theme_fit=yes: It is a mono-white Commander blink deck led by a blink commander at a fair-game bracket-3 power level, and the heavy Lord of the Rings/Final Fantasy card choices are consistent with building from a personal collection first.
- PLAN useful_as_built=yes: Thirty-six lands plus ten ramp pieces support a low curve, the removal and wipe package is deep, and the Angels, Faramir tokens and equipped bodies give a workable if slow route to closing the game.
- PLAN summary_honest=partly: Artifacts, Angels, Humans, equipment, protection and the broad answer suite are all genuinely present, but 'converts accumulated board presence into decisive evasive combat' oversells a list with only a handful of fliers and a single real finisher, and the summary never plainly states the Gilraen targeting-blink engine that the deck is actually built on.
- [INFO] `curve_summary`: average mana value 2.87 over 63 nonland cards
- [WARN] `finisher_short`: the finisher count is 1, and this deck plan wants 2 or more (Angel of Serenity)

<details><summary>The deck list</summary>

- 29 Plains | land | Reliable white mana base.
- 1 Command Tower | land | Untapped white source.
- 1 City of Brass | land | Flexible untapped colored source.
- 1 Plaza of Heroes | land | Supports the commander and legendary creatures.
- 1 Minas Tirith | land | White source with late-game card advantage.
- 1 Secluded Courtyard | land | Names the deck's common creature types for white mana.
- 1 Unclaimed Territory | land | Creature-focused colored source.
- 1 Great Hall of the Citadel | land | Legendary-focused white source.
- 1 Sol Ring | ramp | Efficient acceleration.
- 1 Arcane Signet | ramp | Reliable colored acceleration.
- 1 Fellwar Stone | ramp | Low-cost mana fixing and acceleration.
- 1 Wayfarer's Bauble | ramp | Finds an additional Plains.
- 1 Sword of the Animist | ramp | Creature attacks convert into land ramp.
- 1 Bender's Waterskin | ramp | Artifact mana acceleration.
- 1 White Lotus Tile | ramp | Low-cost mana development.
- 1 Relic of Legends | ramp | Turns legendary creatures into mana sources.
- 1 Thought Vessel | ramp | Mana rock with hand-size utility.
- 1 Inherited Envelope | ramp | Additional artifact acceleration.
- 1 Idol of Oblivion | draw | Repeatable card advantage alongside token production.
- 1 Lembas | draw | Cheap cantrip artifact that is easy to reuse.
- 1 Mask of Memory | draw | Attack-based card filtering and advantage.
- 1 Tome of Legends | draw | Steady cards from commander activity.
- 1 Skullclamp | draw | Efficient card conversion for small creatures.
- 1 Buster Sword | draw | Equipment-based combat card advantage.
- 1 Crown of Gondor | draw | Legendary equipment that supplies cards.
- 1 Diary of Dreams | draw | Artifact card advantage.
- 1 Instant Ramen | draw | Cheap artifact cantrip support.
- 1 Mirror of Galadriel | draw | Legendary source of ongoing card selection.
- 1 Stone of Erech | draw | Low-cost artifact card advantage.
- 1 Lightning Greaves | interaction | Protects key creatures while granting haste.
- 1 Swiftfoot Boots | interaction | Repeatable commander protection.
- 1 Clever Concealment | interaction | Protects the board from sweepers and removal.
- 1 Gift of Immortality | interaction | Keeps an important creature returning.
- 1 Reprieve | interaction | Tempo interaction that replaces itself.
- 1 Unbreakable Formation | interaction | Protects the board and can create a combat swing.
- 1 Together Forever | interaction | Counters help preserve important creatures.
- 1 Ultimate Magic: Holy | interaction | Flexible protective spell.
- 1 Swords to Plowshares | removal | Premium one-mana creature answer.
- 1 Generous Gift | removal | Answers any problematic permanent.
- 1 Get Lost | removal | Efficient answer to creatures, artifacts, and enchantments.
- 1 Destroy Evil | removal | Flexible answer to large creatures or enchantments.
- 1 Stroke of Midnight | removal | Instant-speed universal permanent answer.
- 1 Crib Swap | removal | Exiles creatures and supports the creature-type plan.
- 1 March of Otherworldly Light | removal | Scalable exile removal.
- 1 Banishing Light | removal | Reusable permanent-based answer.
- 1 Journey to Nowhere | removal | Efficient creature exile effect.
- 1 Austere Command | wipe | Flexible reset that can spare chosen permanents.
- 1 Dusk // Dawn | wipe | Creature reset with a later creature-card refill.
- 1 Vanquish the Horde | wipe | Efficient mass creature answer.
- 1 Flickerwisp | synergy | Reusable blink trigger for value creatures and permanents.
- 1 Angel of Condemnation | synergy | Repeatable blink outlet with removal utility.
- 1 Fiend Hunter | synergy | Blinking can repeatedly disrupt opposing creatures.
- 1 Palace Jailer | synergy | Blinking refreshes its exile effect and monarch pressure.
- 1 Personify | synergy | Supports the blink-focused value plan.
- 1 Slip On the Ring | synergy | Instant-speed blink protection and enter-the-battlefield reuse.
- 1 Wall of Omens | synergy | A cheap blinkable source of card advantage.
- 1 Giada, Font of Hope | threat | Develops the Angel board while increasing its pressure.
- 1 Exemplar of Light | threat | Flying creature that contributes to a growing board.
- 1 Inspiring Overseer | threat | Evasive body that advances the board while providing value.
- 1 Faramir, Field Commander | threat | Builds a resilient Human-led battlefield.
- 1 Champions of Minas Tirith | threat | Creature-based pressure with value attached.
- 1 Puresteel Paladin | threat | Turns equipment into efficient combat pressure.
- 1 Bastion Protector | threat | Protects the commander while attacking effectively.
- 1 Boromir, Warden of the Tower | threat | Disruptive legendary attacker.
- 1 Frontline Medic | threat | Combat-oriented creature that protects attacks.
- 1 Kataki, War's Wage | threat | Creature pressure that taxes artifact-heavy opponents.
- 1 Westfold Rider | threat | Creature pressure with useful removal utility.
- 1 Zack Fair | threat | Legendary attacker that supports the creature plan.
- 1 Angel of Serenity | wincon | Large evasive finisher with substantial battlefield impact.
- 1 Angel of Sanctions | wincon | Evasive threat that can close games while disrupting blockers.
- 1 Bronze Guardian | wincon | Protected artifact-based finisher that scales with the board.

</details>

### 6. Modern tempo, tournament

Format: Modern. Theme: tempo. Pool: any_card. Shortlist: 167 names.

Grade: bad, score 0.39, model 20260914T154223Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $316.84 to buy, $316.84 the whole deck.

**Summary:** This blue-red tempo deck establishes pressure with a compact threat suite, then protects that pressure with efficient interaction and removal while its draw package keeps the hand supplied. Its temporal cards reinforce the proactive plan and give the deck another way to press an advantage. It wins by maintaining the initiative and closing before slower opposing plans take over, giving up broader answers and larger standalone threats for speed, consistency, and a lean curve.

The quality model grades this deck below the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The deck runs its spells as playsets, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade.

- PLAN plan_coherent=partly: The twenty-four cheap threats, counters, and burn form a clean tempo core, but four copies of a seven-mana Temporal Mastery with no miracle enablers or cost reduction, alongside a bloated 28 lands, pull against the lean proactive curve.
- PLAN theme_fit=partly: It is a Modern blue-red tempo shell with Ragavan, Shredder, Bolt, and Counterspell, but it contains no Delver of Secrets (or any one-mana threat) despite the request naming Delver, and the Temporal Mastery package belongs to a different archetype.
- PLAN useful_as_built=partly: The mana base is smooth and the interaction plus three two-drop threats can actually close games, but 28 lands with a two-mana top end plus four effectively dead miracle cards means frequent flood and only twelve real threats to win with.
- PLAN summary_honest=no: It asserts the "temporal cards reinforce the proactive plan and give the deck another way to press an advantage," which misrepresents four largely uncastable seven-drops, and it hides the heavy land count and the absence of the Delver the requester asked for.
- [INFO] `curve_summary`: average mana value 2.00 over 36 nonland cards

<details><summary>The deck list</summary>

- 10 Island | land | Provides a deep blue mana base for the deck's blue spells.
- 6 Mountain | land | Provides reliable red mana for the red half of the deck.
- 4 Steam Vents | land | Supplies both deck colors without using a tapped-land slot.
- 4 Spirebluff Canal | land | Adds efficient blue-red mana for an aggressive tempo opening.
- 4 Consider | draw | Forms the core low-cost draw package.
- 2 Preordain | draw | Adds more inexpensive card selection to keep the deck moving.
- 4 Counterspell | interaction | Provides the primary broad interaction suite.
- 2 Spell Pierce | interaction | Adds low-cost interaction for protecting the tempo plan.
- 4 Lightning Bolt | removal | Supplies efficient removal for clearing the way.
- 4 Galvanic Discharge | removal | Completes the removal package with another instant-speed option.
- 4 Temporal Mastery | synergy | Provides the deck's requested temporal synergy package.
- 4 Ragavan, Nimble Pilferer | threat | Gives the deck an early threat that fits its blue-red pressure plan.
- 4 Ledger Shredder | threat | Adds a compact evasive threat suite for sustained pressure.
- 4 Faerie Mastermind | threat | Rounds out the threat base while keeping the deck's curve low.

</details>

### 7. Modern burn, casual

Format: Modern. Theme: burn. Pool: any_card. Shortlist: 290 names.

Grade: baseline, score 0.38, model 20260914T154223Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $46.32 to buy, $46.32 the whole deck.

**Summary:** This mono-red burn deck applies pressure with direct removal, then reinforces that plan with Eidolon of the Great Revel and Guttersnipe while its larger red creatures keep the battlefield demanding attention. It wins by combining repeated burn pressure with combat from Ashcloud Phoenix, Hazoret the Fervent, Sunspine Lynx, and Torbran, Thane of Red Fell. The deck gives up flexible answers and a deep late-game plan in exchange for a focused, proactive red assault.

The quality model grades this deck at the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade.

- PLAN plan_coherent=no: Guttersnipe and Eidolon of the Great Revel are payoffs for a dense cheap-spell burn shell, yet the list has only six burn spells and fourteen four-and-five-drop creatures, and Hazoret's empty-hand requirement fights the card-draw of Risk Factor and Expedite, so the pieces pull in different directions.
- PLAN theme_fit=no: The request was a mono-red burn deck, but with just four Lightning Bolt, two Burst Lightning, and four Risk Factor against sixteen creatures topping out at five mana, this is a clunky red creature-midrange pile rather than burn.
- PLAN useful_as_built=partly: The mana base is fine (24 Mountains, mono-red) and the creatures give a real clock, but the curve is very top-heavy for a deck whose cheap-spell payoffs and Hazoret want speed, so it plays out awkwardly though it does function.
- PLAN summary_honest=partly: placeholder
- [INFO] `curve_summary`: average mana value 2.83 over 36 nonland cards

<details><summary>The deck list</summary>

- 24 Mountain | land | Provides the deck's red mana base.
- 2 Chandra, Dressed to Kill | ramp | Fills the ramp slots while staying in the deck's primary color.
- 4 Risk Factor | draw | Supplies the deck's primary draw package while fitting the burn plan.
- 2 Expedite | draw | Adds inexpensive draw to keep the deck moving.
- 4 Lightning Bolt | removal | Forms the most direct portion of the burn removal suite.
- 2 Burst Lightning | removal | Adds further efficient removal for the burn plan.
- 4 Eidolon of the Great Revel | synergy | Supports the deck's aggressive burn-focused game plan.
- 4 Guttersnipe | synergy | Rewards the deck for leaning on burn spells.
- 4 Ashcloud Phoenix | threat | Provides a resilient creature threat for games that go longer.
- 4 Hazoret the Fervent | threat | Adds a powerful red creature threat to pressure opponents.
- 4 Sunspine Lynx | threat | Provides another substantial creature threat.
- 2 Torbran, Thane of Red Fell | threat | Rounds out the threat suite with a red legendary creature.

</details>

### 8. Modern lifegain, FNM

Format: Modern. Theme: lifegain. Pool: any_card. Shortlist: 301 names.

Grade: bad, score 0.30, model 20260914T154223Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $213.06 to buy, $213.06 the whole deck.

**Summary:** This white-black lifegain deck establishes early creatures, uses lifegain payoffs to turn that foundation into increasingly dangerous combat pressure, and clears opposing blockers with efficient removal. Its strongest games build a wide board around Soul Warden and Ajani's Pridemate before closing with its larger creature threats, chief among them Attended Healer. The deck gives up some speed and flexibility against strategies that can repeatedly answer its early creatures or invalidate creature combat.

The quality model grades this deck below the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade.

- PLAN plan_coherent=yes: Nearly every card feeds one line — gain life with Soul Warden, Lembas, Solitude, lifelinkers, then convert it with Ajani's Pridemate, Bloodbond Vampire, Attended Healer, Dawn of Hope and Enduring Tenacity — with only Altar of the Pantheon sitting loosely on the edge of the plan.
- PLAN theme_fit=yes: It is a white-black Modern lifegain deck at a casual FNM power level, exactly the request, using the classic Soul Warden/Pridemate shell rather than a different archetype.
- PLAN useful_as_built=yes: 24 lands with four untapped-capable duals support a low curve, the removal is castable or free via evoke, and the creature payoffs provide enough clocks to actually close games.
- PLAN summary_honest=partly: It accurately names the payoff creatures and the removal, but its "build a wide board" framing is only thinly supported (Attended Healer and Dawn of Hope are the sole token makers) and it omits the deck's real alternate win route, the drain from Enduring Tenacity.
- [INFO] `curve_summary`: average mana value 3.11 over 36 nonland cards

<details><summary>The deck list</summary>

- 6 Plains | land | Reliable white mana for the deck's early lifegain and creature plays.
- 6 Swamp | land | Reliable black mana for its black threats and removal.
- 4 Godless Shrine | land | Provides both deck colors without using a tapped-land slot.
- 4 Caves of Koilos | land | Flexible white-black source that supports early plays.
- 4 Shattered Sanctum | land | Dual-color land that reinforces consistent white and black access.
- 4 Lembas | draw | Low-investment card flow that fits the deck's lifegain-focused plan.
- 2 Dawn of Hope | draw | Repeatable card-flow option for longer FNM games.
- 2 Altar of the Pantheon | ramp | Mana acceleration that helps deploy the deck's higher-end threats.
- 4 Solitude | removal | Efficient creature-based answer that preserves board pressure.
- 2 Murderous Rider // Swift End | removal | Flexible removal attached to a creature for a proactive deck.
- 4 Soul Warden | synergy | Early lifegain engine that turns creature-heavy games into synergy fuel.
- 4 Ajani's Pridemate | synergy | A primary payoff for the deck's repeated lifegain triggers.
- 4 Attended Healer | threat | Lifegain-oriented threat that adds board presence as the game develops.
- 4 Bloodbond Vampire | threat | Creature threat that rewards the deck for gaining life repeatedly.
- 4 Twinblade Paladin | threat | Combat-focused threat that fits the white lifegain shell.
- 2 Enduring Tenacity | threat | Black threat that gives the deck a resilient top end.

</details>

### 9. Standard midrange, FNM

Format: Standard. Theme: midrange. Pool: any_card. Shortlist: 164 names.

Grade: bad, score 0.33, model 20260914T154223Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $63.08 to buy, $63.08 the whole deck.

**Summary:** This black-green midrange deck develops its mana, builds a creature-heavy board, and uses flexible removal to keep opposing pressure in check. It wins through sustained creature attacks backed by card advantage and protection for important threats, with Goldvein Hydra chief among its early pressure tools. It gives up some speed for a more resilient, board-centered game plan and can be pressured when its creature board is repeatedly answered.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The cards pair in ways the top lists do not, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core of creatures plus removal plus counter-based protection (Snakeskin Veil, Royal Treatment, Inspiring Call) hangs together as midrange, but eight protection spells sit awkwardly with a creature base whose best bodies are six-mana top-end rather than early threats to defend.
- PLAN theme_fit=yes: It is a two-color black-green Standard midrange list with creatures, removal, and card draw, exactly the archetype and format the request named at a casual FNM level.
- PLAN useful_as_built=partly: The 24-land mana base is solid and there are plenty of ways to win, but with eight cards at five to six mana and only a thin two- and three-drop count, the deck often does nothing for the first few turns and plays clunkier than a midrange curve should.
- PLAN summary_honest=partly: The board-centered, removal-and-card-advantage description matches the list and the weakness is fairly admitted, but naming the six-mana Goldvein Hydra as the deck's chief early pressure tool is plainly false.
- [INFO] `curve_summary`: average mana value 2.56 over 36 nonland cards
- [INFO] `mana_pass`: the builder moved 1 card of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 9 Forest | land | Basic green land for the mana base.
- 11 Swamp | land | Basic black land for the mana base.
- 4 Overgrown Tomb | land | Dual-color land that supports both halves of the deck.
- 2 Llanowar Elves | ramp | Early mana development for the midrange plan.
- 2 Assassin's Trophy | removal | Flexible removal for problematic opposing pieces.
- 2 Bitter Triumph | removal | Efficient removal for protecting the board position.
- 2 Maelstrom Pulse | removal | Versatile removal for permanent-based problems.
- 2 Darkstar Augur | draw | Creature-based card advantage for longer games.
- 2 Phyrexian Arena | draw | Ongoing card advantage for grindy matchups.
- 2 Garruk's Uprising | draw | Card advantage that complements the creature suite.
- 4 Snakeskin Veil | synergy | Part of the creature-focused protection package.
- 2 Royal Treatment | synergy | Supports the deck's plan of preserving key creatures.
- 2 Inspiring Call | synergy | Complements the counter-focused elements of the creature suite.
- 4 Goldvein Hydra | threat | Early creature threat that scales with the game.
- 4 Chomping Changeling | threat | Creature threat for applying pressure in the middle turns.
- 2 Thrashing Brontodon | threat | Sturdy creature threat that contributes to board presence.
- 2 Rottenmouth Viper | threat | Creature threat that strengthens the midrange pressure plan.
- 2 Vein Ripper | threat | Top-end creature threat for closing longer games.

</details>

### 10. Standard aggro, tournament

Format: Standard. Theme: aggro. Pool: any_card. Shortlist: 294 names.

Grade: bad, score 0.34, model 20260914T154223Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $38.25 to buy, $38.25 the whole deck.

**Summary:** This red-white deck wins by applying steady creature pressure and forcing opponents to answer its attacks while removal clears resistance and interaction protects the assault. It gives up a deeper long-game plan in exchange for staying focused on tempo, pressure, and converting an early board into a quick finish.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The creature core plus removal and boots does read as one aggro plan, but the five singleton artifacts (Springleaf Drum, Dependable Quinjet, two Banners, Sunbird Standard) are unfocused filler that pulls toward an artifact/ramp shell the rest of the list never uses.
- PLAN theme_fit=yes: It is a 60-card Standard red-white creature deck with removal and equipment, which is the red-white aggro deck the request asked for.
- PLAN useful_as_built=no: The mana is broken for the spells it holds: the deck's most demanding cards are red (Slickshot Show-Off, Fugitive Codebreaker, Reckless Lackey, the five-mana Redcap Gutter-Dweller) yet it runs 15 Plains to only 9 Mountain with no duals, so the red half will be stranded far too often for the deck to function as drawn.
- PLAN summary_honest=partly: placeholder
- [INFO] `curve_summary`: average mana value 2.58 over 36 nonland cards
- [WARN] `profile_off_band`: the worst color's share of the sources it needs is 0.89, and the tournament level wants 0.9 or more (sources of requirement: W 18.75 of 21, R 12.75 of 14)
- [INFO] `mana_pass`: the builder moved 8 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 9 Mountain | land | Provides a reliable red source for the aggressive core.
- 15 Plains | land | Provides a reliable white source for the deck's support spells.
- 4 Fugitive Codebreaker | draw | Keeps pressure supplied with card advantage.
- 2 Reckless Lackey | draw | Adds further card advantage without leaving the aggressive plan.
- 4 Parting Gust | interaction | Flexible interaction for opposing plays that disrupt the attack.
- 2 Lavaspur Boots | interaction | Provides additional interaction while supporting the creature plan.
- 4 Harsh Annotation | removal | Efficient removal to clear opposing resistance.
- 4 Slickshot Show-Off | synergy | A red aggressive synergy piece that rewards committed attacks.
- 3 Frilled Sparkshooter | threat | An aggressive threat that helps establish early pressure.
- 4 Redcap Gutter-Dweller | threat | A threat that strengthens the deck's creature offense.
- 4 Dragonback Lancer | threat | Rounds out the attacking creature suite with another threat.
- 1 Springleaf Drum | ramp | the mana pass added it to bring the mana base inside the power level
- 1 Dependable Quinjet | ramp | the mana pass added it to bring the mana base inside the power level
- 1 Heraldic Banner | ramp | the mana pass added it to bring the mana base inside the power level
- 1 Patchwork Banner | ramp | the mana pass added it to bring the mana base inside the power level
- 1 Sunbird Standard // Sunbird Effigy | ramp | the mana pass added it to bring the mana base inside the power level

</details>

### 11. Commander with a locked card

Format: Commander. Theme: sacrifice. Pool: any_card. Shortlist: 299 names.

Commander: Karlov of the Ghost Council.

Grade: baseline, score 0.05, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $930.04 to buy, $930.04 the whole deck.

**Summary:** This Karlov deck is built around a concentrated sacrifice plan, using its sacrifice synergies alongside creature threats to pressure the table and designated finishers and win-condition cards to close games. It gives up broad, generic value for that focused plan while maintaining dedicated answers, protective pieces, and several ways to reset the board.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The land count sits above the norm of the format, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every nonland card is a sac outlet, death-trigger payoff, token producer, or recursion piece, and the drain effects (Bastion of Remembrance, Elas il-Kor, Vindictive Vampire, Meathook) feed Karlov's lifegain trigger, so the pieces all push one aristocrats plan.
- PLAN theme_fit=yes: It is a Commander deck led by Karlov of the Ghost Council built explicitly around sacrifice, exactly what was asked.
- PLAN useful_as_built=yes: 37 lands with a strong Orzhov dual/painland base, cheap sac outlets and fodder generators, and multiple repeatable drain or mill win routes make it immediately playable, if slightly land-heavy.
- PLAN summary_honest=yes: The claimed sacrifice core, answers (Toxic Deluge, Attrition, Bone Shards), protection (Flare of Fortitude, Gift of Doom, Cartel Aristocrat), board resets, and finishers (Razaketh, Blasting Station, Altar of Dementia, drain engines) are all actually present, and it even flags the high land count.
- [INFO] `curve_summary`: average mana value 3.29 over 62 nonland cards

<details><summary>The deck list</summary>

- 8 Plains | land | Provides a dependable white source for the two-color mana base.
- 10 Swamp | land | Provides a dependable black source for the two-color mana base.
- 1 Brightclimb Pathway // Grimclimb Pathway | land | Flexible white or black land source.
- 1 Caves of Koilos | land | Provides both deck colors.
- 1 Command Tower | land | Reliable two-color fixing.
- 1 Concealed Courtyard | land | Early white or black source.
- 1 City of Brass | land | Flexible color fixing.
- 1 Godless Shrine | land | Dual-color land source.
- 1 Isolated Chapel | land | Supports both colors in the mana base.
- 1 Mana Confluence | land | Flexible color fixing.
- 1 Marsh Flats | land | Finds a basic land source when needed.
- 1 Scrubland | land | Dual-color land source.
- 1 Shattered Sanctum | land | Supports both deck colors.
- 1 Silent Clearing | land | Dual-color land source.
- 1 Vault of Champions | land | Reliable white or black source.
- 1 Eclipsed Steppe | land | Provides access to both deck colors.
- 1 Turbulent Moor | land | Provides access to both deck colors.
- 1 Prismatic Vista | land | Finds a basic land source when needed.
- 1 High Market | land | Utility land for the sacrifice plan.
- 1 Phyrexian Tower | land | Utility land that supports the sacrifice plan.
- 1 Fountainport | land | Adds a utility land slot without weakening the color base.
- 1 Sol Ring | ramp | Included because it is required to remain in the deck.
- 1 Ashnod's Altar | ramp | Sacrifice-focused mana development.
- 1 Phyrexian Altar | ramp | Sacrifice-focused mana development.
- 1 Pitiless Plunderer | ramp | Supports mana development in the sacrifice shell.
- 1 Priest of Forgotten Gods | ramp | Adds sacrifice-focused mana development.
- 1 Pawn of Ulamog | ramp | Supports mana development alongside sacrificed creatures.
- 1 Sifter of Skulls | ramp | Supports mana development in the creature-heavy plan.
- 1 Warren Soultrader | ramp | Adds sacrifice-oriented mana development.
- 1 Skullport Merchant | ramp | Provides mana development for the deck's midgame.
- 1 Deadly Dispute | ramp | Efficient sacrifice-oriented mana development.
- 1 Corrupted Conviction | draw | Efficient card flow for a sacrifice-focused deck.
- 1 Village Rites | draw | Low-cost card flow that fits the sacrifice plan.
- 1 Vampiric Rites | draw | Repeatable card-flow support for the sacrifice shell.
- 1 Disciple of Bolas | draw | Creature-based card-flow option.
- 1 Smothering Abomination | draw | Draw engine that rewards the deck's core plan.
- 1 Baron Bertram Graywater | draw | Creature-based source of card flow.
- 1 Bushmeat Poacher | draw | Sacrifice-oriented card-flow support.
- 1 Jerren, Corrupted Bishop // Ormendahl, the Corrupter | draw | Adds card flow while remaining within the creature plan.
- 1 Tevesh Szat, Doom of Fools | draw | Provides sustained card-flow support.
- 1 Trading Post | draw | Flexible artifact-based card-flow option.
- 1 Cartel Aristocrat | interaction | Protective interaction that fits the sacrifice shell.
- 1 Fanatical Devotion | interaction | Sacrifice-oriented protective interaction.
- 1 Flare of Fortitude | interaction | Protects the board at a critical moment.
- 1 Gift of Doom | interaction | Provides protective interaction for a key permanent.
- 1 Promise of Tomorrow | interaction | Defensive interaction for the creature plan.
- 1 Spirit Bonds | interaction | Interaction that supports the deck's creature focus.
- 1 Sunstone | interaction | Defensive interaction that buys time to develop.
- 1 Valkyrie's Call | interaction | Protective interaction for the sacrifice strategy.
- 1 Attrition | removal | Repeatable sacrifice-oriented creature control.
- 1 Ayli, Eternal Pilgrim | removal | Creature-based removal that fits the deck's plan.
- 1 Blasting Station | removal | Sacrifice-oriented removal option.
- 1 Bone Shards | removal | Efficient removal with a sacrifice-friendly cost.
- 1 Eaten Alive | removal | Removal that works naturally with sacrificed creatures.
- 1 Grave Pact | removal | Broad removal pressure for the sacrifice strategy.
- 1 Dictate of Erebos | removal | Adds broad removal pressure to creature sacrifices.
- 1 Teysa, Orzhov Scion | removal | Creature-based removal that supports the deck theme.
- 1 Yawgmoth, Thran Physician | removal | Flexible creature-based removal.
- 1 Altar of Dementia | synergy | Core sacrifice synergy piece.
- 1 Bastion of Remembrance | synergy | Sacrifice payoff that advances the deck plan.
- 1 Carrion Feeder | synergy | Low-cost sacrifice-focused synergy piece.
- 1 Chthonian Nightmare | synergy | Supports the deck's sacrifice-focused engine.
- 1 Elas il-Kor, Sadistic Pilgrim | synergy | Creature-based sacrifice payoff.
- 1 Fleshtaker | synergy | Supports the sacrifice-focused creature package.
- 1 Viscera Seer | synergy | Low-cost sacrifice synergy for the deck.
- 1 Basri's Lieutenant | threat | Adds a white creature threat to pressure opponents.
- 1 Corpse Harvester | threat | Creature threat that fits the graveyard-oriented shell.
- 1 Felisa, Fang of Silverquill | threat | Orzhov creature threat for the sacrifice plan.
- 1 Liesa, Forgotten Archangel | threat | High-impact Orzhov creature threat.
- 1 Mondrak, Glory Dominus | threat | Creature threat that complements the token-oriented elements.
- 1 Ratadrabik of Urborg | threat | Legendary creature threat for the Orzhov shell.
- 1 Vindictive Vampire | threat | Sacrifice-themed creature threat.
- 1 Sadistic Hypnotist | threat | Sacrifice-focused threat that pressures opposing resources.
- 1 Marrow-Gnawer | threat | Creature threat that supports a go-wide board plan.
- 1 Prowling Geistcatcher | threat | Creature threat for the sacrifice-focused strategy.
- 1 Requiem Angel | threat | Evasive creature threat and designated finisher.
- 1 Thallid Omnivore | threat | Creature threat that fits the sacrifice theme.
- 1 Razaketh, the Foulblooded | wincon | Designated finisher and win-condition card.
- 1 Relic Vial | wincon | Designated finisher and win-condition card.
- 1 Sephiroth, Fabled SOLDIER // Sephiroth, One-Winged Angel | wincon | Designated finisher and win-condition card.
- 1 Austere Command | wipe | Flexible board-reset option.
- 1 Toxic Deluge | wipe | Efficient board-reset option.
- 1 The Meathook Massacre | wipe | Board-reset card that also serves as a designated finisher.

</details>

### 12. Commander on a budget

Format: Commander. Theme: tokens. Pool: any_card. Shortlist: 298 names.

Commander: Adeline, Resplendent Cathar.

Grade: bad, score 0.02, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $67.76 to buy, $67.76 the whole deck.

**Summary:** Adeline, Resplendent Cathar leads a straightforward mono-white token deck that builds a wide battlefield, reinforces it with token-focused support, and turns repeated attacks into mounting pressure. The deck wins through its finishing threats, dedicated win conditions, and a sustained creature presence, with enough removal and board resets to keep opposing plans from taking over. It gives up premium mana and expensive high-end staples for a consistent Plains-based foundation and a deliberately budget-conscious suite of threats and support cards.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The deck makes less mana on turn four than the norm, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every card either makes creature tokens, buffs a wide board (Intangible Virtue, Divine Visitation, Siege Veteran), draws off attacking/token creation, or resets boards asymmetrically (Hour of Reckoning), so the pieces line up behind Adeline's go-wide attack plan; the handful of trinket ramp artifacts (Coin of Mastery, Goldvein Pick, Currency Converter) are weak filler but do not pull against the plan.
- PLAN theme_fit=yes: It is a mono-white Adeline token Commander deck built from cheap, widely-reprinted commons/uncommons and mid-tier rares with no expensive staples, matching the under-$100, bracket-2 request.
- PLAN useful_as_built=yes: 37 Plains in a mono-color deck cannot stumble on colors, the curve is mostly two-to-four drops with several token payoffs, and there are clear finishers (Adeline plus Hero of Bladehold, Divine Visitation, Halo Fountain), though the ramp package is slow enough that big spells land late.
- PLAN summary_honest=yes: placeholder
- [INFO] `curve_summary`: average mana value 3.37 over 62 nonland cards

<details><summary>The deck list</summary>

- 37 Plains | land | Provides a consistent untapped white mana base.
- 1 Bygone Bishop | draw | Provides token-oriented card draw support.
- 1 Dawn of Hope | draw | Adds a repeatable draw outlet for the token plan.
- 1 Idol of Oblivion | draw | Provides efficient card draw for a deck making tokens.
- 1 Platoon Dispenser | draw | Supplies card advantage while fitting the artifact-token shell.
- 1 Search the Premises | draw | Adds a defensive source of card advantage.
- 1 Staff of the Storyteller | draw | Supports the board-building plan with card draw.
- 1 Wedding Announcement // Wedding Festivity | draw | Contributes to the token plan while providing card draw.
- 1 Wojek Investigator | draw | Adds a low-cost draw piece to keep resources flowing.
- 1 Valiant Rescuer | draw | Provides draw support in the creature-heavy token shell.
- 1 Glimmer Seeker | draw | Adds an inexpensive card-draw body.
- 1 Sanctuary Warden | draw | Provides a high-impact draw threat and a finisher.
- 1 Basri Ket | interaction | Protects the developed board and supports combat turns.
- 1 Lena, Selfless Champion | interaction | Helps preserve a wide creature board.
- 1 Mage's Attendant | interaction | Adds interactive coverage on a token-friendly creature.
- 1 Rootborn Defenses | interaction | Protects the token army during pivotal turns.
- 1 Paladin's Arms | interaction | Provides a low-cost interactive combat tool.
- 1 Squad Commander | interaction | Adds interaction while rewarding a broad battlefield.
- 1 Coin of Mastery | ramp | Provides artifact-based mana acceleration.
- 1 Collector's Vault | ramp | Adds affordable mana acceleration.
- 1 Currency Converter | ramp | Provides a cheap artifact ramp option.
- 1 Druidic Satchel | ramp | Supports mana development with a low-cost permanent.
- 1 Goldvein Pick | ramp | Adds mana acceleration that works with attacking creatures.
- 1 Karn, Living Legacy | ramp | Provides ongoing mana development from a resilient permanent.
- 1 Keeper of the Accord | ramp | Helps maintain mana development through a creature.
- 1 Legion's Landing // Adanto, the First Fort | ramp | Supports early board development and mana progression.
- 1 Monologue Tax | ramp | Provides passive mana acceleration over the game.
- 1 The Restoration of Eiganjo // Architect of Restoration | ramp | Advances mana development while remaining useful later.
- 1 Aerial Assault | removal | Provides efficient creature removal.
- 1 Banishing Slash | removal | Adds flexible sorcery-speed removal.
- 1 Battle Menu | removal | Provides a cheap removal option.
- 1 Generous Gift | removal | Answers a broad range of opposing permanents.
- 1 Kellan's Lightblades | removal | Adds efficient instant-speed removal.
- 1 Skyclave Apparition | removal | Provides removal on a creature body.
- 1 Stroke of Midnight | removal | Adds flexible instant-speed removal.
- 1 Cat Collector | synergy | Supports the deck's token-focused synergies.
- 1 Clarion Spirit | synergy | Rewards the deck for developing its board.
- 1 Divine Visitation | synergy | Strengthens the payoff for making creature tokens.
- 1 Felidar Retreat | synergy | Builds the battlefield in support of the token plan.
- 1 Intangible Virtue | synergy | Improves the effectiveness of the token army.
- 1 Hanweir Militia Captain // Westvale Cult Leader | synergy | Rewards going wide with creatures.
- 1 Rosie Cotton of South Lane | synergy | Turns token development into stronger creatures.
- 1 Siege Veteran | synergy | Supports a creature-heavy board with ongoing value.
- 1 Summoner's Sending | synergy | Adds a low-cost token-focused support piece.
- 1 Twilight Drover | synergy | Benefits from and reinforces the token plan.
- 1 Ajani's Chosen | threat | Provides a token-oriented battlefield threat.
- 1 Archon of Sun's Grace | threat | Adds an evasive token-focused threat.
- 1 Attended Healer | threat | Builds a board presence that pressures opponents.
- 1 Basri's Lieutenant | threat | Provides a combat-ready creature threat.
- 1 Cemetery Protector | threat | Offers a durable creature threat with useful board presence.
- 1 Defiler of Faith | threat | Adds a substantial white creature threat.
- 1 Emeria Angel | threat | Turns ongoing development into airborne board pressure.
- 1 Gwaihir, Greatest of the Eagles | threat | Provides an evasive finishing threat.
- 1 Hero of Bladehold | threat | Pressures opponents while supporting wide attacks.
- 1 Nahiri, the Lithomancer | threat | Provides a resilient finishing threat.
- 1 Requiem Angel | threat | Supplies an evasive finishing threat for the token shell.
- 1 Silverwing Squadron | threat | Provides a wide-board payoff and creature threat.
- 1 Halo Fountain | wincon | Gives the deck a dedicated token-based finishing route.
- 1 Luck Bobblehead | wincon | Adds a standalone finishing route.
- 1 Sword of Body and Mind | wincon | Provides a combat-based finishing route.
- 1 Ceaseless Conflict | wipe | Provides a budget reset when the board gets away.
- 1 Hour of Reckoning | wipe | Clears opposing boards while fitting a creature-heavy strategy.
- 1 Martial Coup | wipe | Resets the battlefield while supporting a later token rebuild.

</details>

### 13. owned first, and the commander is not owned

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 230 names.

Commander: Karlov of the Ghost Council.

Grade: baseline, score 0.06, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $187.77 the whole deck.

**Summary:** Karlov of the Ghost Council leads a lifegain-focused board deck that develops mana early, uses its draw and protection pieces to keep key creatures in play, and pressures the table with Angels and other creature threats. Targeted answers and sweepers keep opposing boards manageable, while Frodo, Sauron's Bane, Grave Venerations, and Lyra Dawnbringer provide the closing power. The deck gives up some flexibility when its lifegain synergies are absent, so preserving Karlov and maintaining a steady board are important.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. Few lands enter tapped, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core is a coherent W/B lifegain-and-protect-Karlov shell (Angel of Vitality, Rosie Cotton, Light of Promise, Shattered Angel, Exemplar of Light plus equipment and protection), but a large slice of the list is generic goodstuff removal, value artifacts, and three sweepers that sit awkwardly beside its own creature-based threats, so the pieces only partly pull one direction.
- PLAN theme_fit=yes: It is a Commander deck led by Karlov of the Ghost Council with a lifegain theme in two colors, built from mostly non-staple, collection-flavored cards at a mid-power level consistent with bracket 3.
- PLAN useful_as_built=yes: 36 lands with Sol Ring and several two-mana rocks in a two-color deck, a sane curve, ample card draw and removal, and multiple creature-based damage routes plus Karlov himself mean it can be shuffled up and played as printed.
- PLAN summary_honest=partly: The lifegain, protection, removal, and sweeper claims match the list, but calling the deck one that "pressures the table with Angels" overstates a handful of angels, and naming Grave Venerations and Lyra as "closing power" dresses up ordinary creature/synergy cards as finishers.
- [WARN] `not_owned`: Karlov of the Ghost Council: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.11 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Plains | land | Provides a dependable white mana base.
- 1 Plains | land | Provides a dependable white mana base.
- 1 Plains | land | Provides a dependable white mana base.
- 1 Plains | land | Provides a dependable white mana base.
- 1 Plains | land | Provides a dependable white mana base.
- 1 Plains | land | Provides a dependable white mana base.
- 1 Plains | land | Provides a dependable white mana base.
- 1 Plains | land | Provides a dependable white mana base.
- 1 Plains | land | Provides a dependable white mana base.
- 1 Plains | land | Provides a dependable white mana base.
- 1 Plains | land | Provides a dependable white mana base.
- 1 Plains | land | Provides a dependable white mana base.
- 1 Plains | land | Provides a dependable white mana base.
- 1 Plains | land | Provides a dependable white mana base.
- 1 Plains | land | Provides a dependable white mana base.
- 1 Plains | land | Provides a dependable white mana base.
- 1 Swamp | land | Provides a dependable black mana base.
- 1 Swamp | land | Provides a dependable black mana base.
- 1 Swamp | land | Provides a dependable black mana base.
- 1 Swamp | land | Provides a dependable black mana base.
- 1 Swamp | land | Provides a dependable black mana base.
- 1 Swamp | land | Provides a dependable black mana base.
- 1 Swamp | land | Provides a dependable black mana base.
- 1 Swamp | land | Provides a dependable black mana base.
- 1 Swamp | land | Provides a dependable black mana base.
- 1 Swamp | land | Provides a dependable black mana base.
- 1 Swamp | land | Provides a dependable black mana base.
- 1 Swamp | land | Provides a dependable black mana base.
- 1 Swamp | land | Provides a dependable black mana base.
- 1 Swamp | land | Provides a dependable black mana base.
- 1 Swamp | land | Provides a dependable black mana base.
- 1 Command Tower | land | Supplies either color for Karlov's deck.
- 1 City of Brass | land | Supplies flexible colored mana.
- 1 Exotic Orchard | land | Supplies flexible colored mana.
- 1 Marsh Flats | land | Finds the basic color needed.
- 1 Fabled Passage | land | Finds the basic color needed.
- 1 Arcane Signet | ramp | Efficiently advances the deck's mana.
- 1 Commander's Sphere | ramp | Provides reliable mana development.
- 1 Fellwar Stone | ramp | Provides inexpensive mana acceleration.
- 1 Hot Dog Cart | ramp | Adds to the deck's mana development.
- 1 Inherited Envelope | ramp | Adds to the deck's mana development.
- 1 Sol Ring | ramp | Provides an early mana boost.
- 1 Sword of the Animist | ramp | Develops mana while supporting creature attacks.
- 1 Thought Vessel | ramp | Provides steady mana acceleration.
- 1 Wayfarer's Bauble | ramp | Converts early resources into lasting mana.
- 1 White Auracite | ramp | Adds to the deck's mana development.
- 1 Buster Sword | draw | Provides card flow from an equipment slot.
- 1 Call of the Ring | draw | Supplies ongoing card flow.
- 1 Exemplar of Light | draw | Adds card flow on a lifegain-friendly body.
- 1 Idol of Oblivion | draw | Provides a compact card-flow option.
- 1 Inspiring Overseer | draw | Adds a creature that supports card flow.
- 1 Lembas | draw | Provides a compact card-flow piece.
- 1 Mask of Memory | draw | Turns combat into card flow.
- 1 Night's Whisper | draw | Provides efficient card flow.
- 1 Painful Truths | draw | Provides efficient card flow.
- 1 Skullclamp | draw | Provides a powerful equipment-based draw outlet.
- 1 Tome of Legends | draw | Supports continued card flow alongside Karlov.
- 1 Bastion Protector | interaction | Helps defend Karlov and the board plan.
- 1 Champion's Helm | interaction | Protects an important legendary creature.
- 1 Clever Concealment | interaction | Guards the developed board at a key moment.
- 1 Darksteel Plate | interaction | Provides durable protection for a key creature.
- 1 Lightning Greaves | interaction | Protects Karlov or a key threat.
- 1 Reprieve | interaction | Offers a flexible tempo response.
- 1 Swiftfoot Boots | interaction | Protects Karlov or a key threat.
- 1 Unbreakable Formation | interaction | Helps preserve the creature board.
- 1 Banishing Light | removal | Answers a problematic opposing permanent.
- 1 Bitter Triumph | removal | Provides a flexible answer to a threat.
- 1 Crib Swap | removal | Answers an opposing creature.
- 1 Fatal Push | removal | Provides an efficient creature answer.
- 1 Generous Gift | removal | Answers a wide range of opposing permanents.
- 1 Get Lost | removal | Provides a flexible permanent answer.
- 1 Infernal Grasp | removal | Provides a direct creature answer.
- 1 Path to Exile | removal | Provides a clean creature answer.
- 1 Swords to Plowshares | removal | Provides a clean creature answer.
- 1 Angel of Vitality | synergy | Reinforces the deck's lifegain theme.
- 1 Aettir and Priwen | synergy | Supports the deck's creature-focused lifegain plan.
- 1 Compassionate Healer | synergy | Reinforces the deck's lifegain theme.
- 1 Kor Firewalker | synergy | Provides a lifegain-focused synergy piece.
- 1 Light of Promise | synergy | Rewards the deck's lifegain plan.
- 1 Rosie Cotton of South Lane | synergy | Supports the deck's lifegain-focused board plan.
- 1 Second Breakfast | synergy | Reinforces the deck's lifegain theme.
- 1 Angel of Invention | threat | Provides a substantial creature threat.
- 1 Bill the Pony | threat | Adds a creature threat to pressure opponents.
- 1 Canyon Crawler | threat | Adds a creature threat to pressure opponents.
- 1 Dawnhand Eulogist | threat | Provides a threatening finisher-class creature.
- 1 Gollum, Silent Slinker // Meager Meal | threat | Adds a creature threat to pressure opponents.
- 1 Melancholic Poet | threat | Provides a threatening finisher-class creature.
- 1 Minwu, White Mage | threat | Adds a creature threat to pressure opponents.
- 1 Rabaroo Troop | threat | Adds a creature threat to pressure opponents.
- 1 Rooftop Percher | threat | Adds a creature threat to pressure opponents.
- 1 Shattered Angel | threat | Adds a creature threat to pressure opponents.
- 1 Sneering Shadewriter | threat | Provides a threatening finisher-class creature.
- 1 Victory's Herald | threat | Adds a creature threat to pressure opponents.
- 1 Frodo, Sauron's Bane | wincon | Provides a dedicated finishing route.
- 1 Grave Venerations | wincon | Provides a dedicated finishing route.
- 1 Lyra Dawnbringer | wincon | Provides a dedicated finishing route.
- 1 Austere Command | wipe | Provides a flexible board reset.
- 1 Fumigate | wipe | Provides a creature-board reset.
- 1 Vanquish the Horde | wipe | Provides an efficient board reset.

</details>

### 14. the user delegates the commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 247 names.

Commander: Atarka, World Render.

Grade: bad, score 0.04, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $576.51 to buy, $576.51 the whole deck.

**Summary:** Atarka, World Render leads a focused Gruul Dragon deck that ramps into a steady stream of Dragons, protects its key board pieces, and uses Dragon-themed removal to clear resistance. The deck closes through its large combat threats and its dedicated Dragon win conditions, with its chief strength being concentrated creature pressure. It gives up broad noncreature flexibility for a committed Dragon strategy and a meaningful top end.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Every slot supports one Dragon-tribal plan: cost reducers and Dragon-flavored ramp, Dragon-triggered draw (Elemental Bond, Dragonborn Champion, Terror of the Peaks), Dragon-conditional removal, and a top end of Lathliss/Utvara/Atarka payoffs.
- PLAN theme_fit=yes: It is a Commander Dragons deck with a builder-chosen Dragon commander, and the density of cheap enablers and mid-sized Dragons rather than fast combos or heavy tutors sits sensibly at bracket 3.
- PLAN useful_as_built=yes: 36 lands in two colors with fixing plus about ten ramp/cost-reduction pieces supports the heavy five-to-six drop curve, and there is ample card draw and several redundant Dragon win conditions, so it plays out of the box even if the eight-mana commander lands late.
- PLAN summary_honest=yes: Each claim maps to real cards — ramp (orbs, Dragonspeaker Shaman, Goldspan), protection (Heroic Intervention, Boots/Greaves, Snakeskin Veil), Dragon-costed removal (Draconic Roar, Dragon's Fire, Spit Flame), and token-based closers — and it openly concedes the thin noncreature interaction and slow mana.
- [INFO] `curve_summary`: average mana value 3.32 over 63 nonland cards

<details><summary>The deck list</summary>

- 12 Forest | land | Provides the deck's basic green mana base.
- 12 Mountain | land | Provides the deck's basic red mana base.
- 1 Command Tower | land | Part of the two-color mana base.
- 1 Stomping Ground | land | Part of the two-color mana base.
- 1 Taiga | land | Part of the two-color mana base.
- 1 Karplusan Forest | land | Part of the two-color mana base.
- 1 Copperline Gorge | land | Part of the two-color mana base.
- 1 Spire Garden | land | Part of the two-color mana base.
- 1 Rockfall Vale | land | Part of the two-color mana base.
- 1 Rootbound Crag | land | Part of the two-color mana base.
- 1 Game Trail | land | Part of the two-color mana base.
- 1 Cragcrown Pathway // Timbercrown Pathway | land | Part of the two-color mana base.
- 1 Fire-Lit Thicket | land | Part of the two-color mana base.
- 1 Wooded Foothills | land | Part of the two-color mana base.
- 1 Jade Orb of Dragonkind | ramp | Early ramp for the Dragon plan.
- 1 Orb of Dragonkind | ramp | Early ramp that fits the Dragon theme.
- 1 Scaled Nurturer | ramp | Low-cost Dragon ramp.
- 1 Reckless Barbarian | ramp | Low-cost ramp toward larger Dragons.
- 1 Carnelian Orb of Dragonkind | ramp | Dragon-themed mana acceleration.
- 1 Dragon's Hoard | ramp | Mana acceleration for the deck's larger spells.
- 1 Dragonstorm Globe | ramp | Mana acceleration in the Dragon shell.
- 1 Encroaching Dragonstorm | ramp | Ramp support for the deck's top end.
- 1 Ganax, Astral Hunter | ramp | A Dragon ramp piece for the creature-heavy plan.
- 1 Goldspan Dragon | ramp | A Dragon finisher-class ramp card.
- 1 Faithless Looting | draw | Efficient early card draw.
- 1 Demand Answers | draw | Low-cost card draw support.
- 1 Thrill of Possibility | draw | Low-cost card draw support.
- 1 Sylvan Library | draw | Early ongoing card draw.
- 1 Skullclamp | draw | Low-cost card draw support.
- 1 Elemental Bond | draw | Card draw for the creature-focused plan.
- 1 Garruk's Uprising | draw | Card draw that fits the large-creature strategy.
- 1 Beast Whisperer | draw | Creature-based card draw support.
- 1 Dragonborn Champion | draw | Dragon-themed card draw.
- 1 Guardian Project | draw | Ongoing card draw for a creature-heavy deck.
- 1 Return of the Wildspeaker | draw | Flexible card draw support.
- 1 Heroic Intervention | interaction | Protective interaction for the board.
- 1 Lightning Greaves | interaction | Protection-focused interaction for key creatures.
- 1 Swiftfoot Boots | interaction | Protection-focused interaction for key creatures.
- 1 Snakeskin Veil | interaction | Low-cost protective interaction.
- 1 Tamiyo's Safekeeping | interaction | Low-cost protective interaction.
- 1 Veil of Summer | interaction | Efficient interaction from the shortlist.
- 1 Pyroblast | interaction | Efficient interaction from the shortlist.
- 1 Red Elemental Blast | interaction | Efficient interaction from the shortlist.
- 1 Draconic Roar | removal | Low-cost Dragon-themed removal.
- 1 Dragon's Fire | removal | Low-cost Dragon-themed removal.
- 1 Piercing Exhale | removal | Low-cost Dragon-themed removal.
- 1 Spit Flame | removal | Low-cost Dragon-themed removal.
- 1 Molten Exhale | removal | Low-cost removal support.
- 1 Dragon Tempest | removal | Dragon-themed removal support.
- 1 Invasion of Tarkir // Defiant Thundermaw | removal | Removal that remains within the Dragon plan.
- 1 Glorybringer | removal | A Dragon removal threat.
- 1 Terror of the Peaks | removal | A finisher-class Dragon removal threat.
- 1 Acolyte of Bahamut | synergy | Dragon-focused synergy support.
- 1 Dragonlord's Servant | synergy | Low-cost Dragon synergy support.
- 1 Dragonspeaker Shaman | synergy | Dragon synergy support.
- 1 Crucible of Fire | synergy | A dedicated Dragon synergy piece.
- 1 Sarkhan's Triumph | synergy | Dragon synergy support from the shortlist.
- 1 Dragonkin Berserker | synergy | Low-cost Dragon synergy support.
- 1 The Dragon-Kami Reborn // Dragon-Kami's Egg | synergy | A Dragon-themed synergy card.
- 1 Ambitious Dragonborn | threat | A Dragon threat for the combat plan.
- 1 Backdraft Hellkite | threat | A Dragon threat for the combat plan.
- 1 Caldera Pyremaw | threat | A Dragon threat for the combat plan.
- 1 Dragon Broodmother | threat | A Dragon threat for the combat plan.
- 1 Herdchaser Dragon | threat | A Dragon threat for the combat plan.
- 1 Manaform Hellkite | threat | A Dragon threat for the combat plan.
- 1 Mirrorwing Dragon | threat | A Dragon threat for the combat plan.
- 1 Red Dragon | threat | A Dragon threat for the combat plan.
- 1 Stormbreath Dragon | threat | A Dragon threat for the combat plan.
- 1 Stormwing Dragon | threat | A Dragon threat for the combat plan.
- 1 Thunderbreak Regent | threat | A Dragon threat for the combat plan.
- 1 Territorial Hellkite | threat | A Dragon threat for the combat plan.
- 1 Lathliss, Dragon Queen | wincon | A finisher-class Dragon win condition.
- 1 Utvara Hellkite | wincon | A finisher-class Dragon win condition.
- 1 Wrathful Red Dragon | wincon | A finisher-class Dragon win condition.
- 1 Breath Weapon | wipe | A low-cost board wipe.
- 1 Draconic Intervention | wipe | A Dragon-themed board wipe.
- 1 Thundermaw Hellkite | wipe | A finisher-class Dragon board wipe.

</details>

### 15. delegated commander, owned first

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 230 names.

Commander: Astarion, the Decadent.

Grade: bad, score 0.03, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $177.42 the whole deck.

**Summary:** Astarion leads a steady Orzhov lifegain deck that establishes mana, develops a board of lifegain-focused creatures and Angels, and uses its Equipment and protection pieces to keep important threats in play. It wins by turning that developed board into sustained pressure or by finding one of its dedicated finishing cards. The deck gives up explosive speed for a stable mana base, broad spot removal, card flow, and several ways to rebuild after a crowded board is cleared.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves one color far better than another, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core is a coherent white-heavy lifegain/Angels board plan that feeds Astarion's counters, but an Equipment sub-package (Puresteel Paladin, Buster Sword, Skullclamp, Mask of Memory) and a handful of generic threats pull a little sideways from the lifegain payoffs, which are fairly thin on the ground.
- PLAN theme_fit=yes: It is an Orzhov lifegain Commander deck with the builder choosing Astarion, the Decadent, and its low-synergy, interaction-heavy, no-fast-mana construction sits comfortably at bracket 2 as requested.
- PLAN useful_as_built=yes: 38 lands plus ten mana rocks/ramp pieces support a modest curve, the Swamp-light base is still workable since black needs are few, and there are real ways to close — Astarion growing huge, Lyra/Victory's Herald/Angel of Invention flying beats, plus Frodo — so the deck plays out fine as it stands.
- PLAN summary_honest=yes: The summary accurately describes what the list does — lands, ramp, lifegain creatures and Angels, Equipment and protection, broad spot removal, wraths — and openly flags the lopsided mana base and lack of speed rather than overclaiming a combo or drain-based kill the deck does not have.
- [WARN] `not_owned`: Astarion, the Decadent: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.03 over 62 nonland cards

<details><summary>The deck list</summary>

- 19 Plains | land | Provides a deep base of white mana for the deck's lifegain cards and spells.
- 10 Swamp | land | Provides reliable black mana for Astarion and the black portion of the deck.
- 1 Command Tower | land | Provides either color of the deck's mana.
- 1 City of Brass | land | Provides flexible color fixing.
- 1 Exotic Orchard | land | Provides flexible color fixing.
- 1 Marsh Flats | land | Helps find the appropriate basic land color.
- 1 Fabled Passage | land | Helps find the appropriate basic land color.
- 1 Evolving Wilds | land | Finds a needed basic land.
- 1 Path of Ancestry | land | Provides both deck colors from the land base.
- 1 Grand Coliseum | land | Provides flexible color fixing.
- 1 Arcane Signet | ramp | Efficiently fixes mana for the deck's colors.
- 1 Sol Ring | ramp | Provides an early mana boost.
- 1 Fellwar Stone | ramp | Adds reliable multicolor mana production.
- 1 Commander's Sphere | ramp | Fixes mana while retaining later utility.
- 1 Chromatic Lantern | ramp | Supports a stable two-color mana base.
- 1 Giada, Font of Hope | ramp | Provides ramp within the Angel portion of the deck.
- 1 Lotho, Corrupt Shirriff | ramp | Contributes to the deck's mana development.
- 1 Oath of the Grey Host | ramp | Provides additional mana development for the midgame.
- 1 Relic of Legends | ramp | Helps turn the deck's legendary cards into mana support.
- 1 Wayfarer's Bauble | ramp | Finds an additional basic land for lasting mana development.
- 1 Buster Sword | draw | Supplies card advantage from an Equipment slot.
- 1 Call of the Ring | draw | Provides a continuing source of cards.
- 1 Exemplar of Light | draw | Adds card draw while fitting the lifegain creature plan.
- 1 Idol of Oblivion | draw | Offers a compact artifact source of cards.
- 1 Inspiring Overseer | draw | Provides card draw on a lifegain-friendly creature.
- 1 Lembas | draw | Provides a small, flexible card advantage piece.
- 1 Mask of Memory | draw | Turns creature combat into card selection and cards.
- 1 Night's Whisper | draw | Provides efficient black card draw.
- 1 Puresteel Paladin | draw | Supports the Equipment package while providing cards.
- 1 Skullclamp | draw | Provides efficient card advantage around small creatures.
- 1 Wall of Omens | draw | Provides an early defensive body and a card.
- 1 Bastion Protector | interaction | Helps protect Astarion from opposing pressure.
- 1 Champion's Helm | interaction | Protects the deck's legendary centerpiece.
- 1 Darksteel Plate | interaction | Adds durable protection for a key creature.
- 1 Lightning Greaves | interaction | Provides immediate protection for an important creature.
- 1 Swiftfoot Boots | interaction | Protects a creature while keeping its abilities available.
- 1 Unbreakable Formation | interaction | Protects the developed creature board.
- 1 Banishing Light | removal | Answers a problematic opposing permanent.
- 1 Bitter Triumph | removal | Provides flexible spot removal.
- 1 Crib Swap | removal | Offers creature removal at instant speed.
- 1 Dispatch | removal | Provides efficient creature removal.
- 1 Generous Gift | removal | Answers a wide range of opposing permanents.
- 1 Infernal Grasp | removal | Provides clean creature removal.
- 1 Swords to Plowshares | removal | Provides efficient creature removal.
- 1 Aerith Gainsborough | synergy | Supports the deck's central lifegain plan.
- 1 Angel of Vitality | synergy | Rewards the deck for pursuing lifegain.
- 1 Compassionate Healer | synergy | Provides another dedicated lifegain synergy piece.
- 1 Crowd of True Believers | synergy | Builds on the deck's lifegain-focused board plan.
- 1 Elixir | synergy | Provides a compact lifegain synergy artifact.
- 1 Eastfarthing Farmer | synergy | Adds a creature-based lifegain synergy piece.
- 1 Light of Promise | synergy | Turns lifegain synergies into a growing board presence.
- 1 Rosie Cotton of South Lane | synergy | Rewards the deck's lifegain and creature development.
- 1 Second Breakfast | synergy | Supports the deck's lifegain-focused game plan.
- 1 Wanderbrine Preacher | synergy | Adds another creature that works with the lifegain shell.
- 1 Angel of Invention | threat | Provides a substantial creature threat for the board.
- 1 Bill the Pony | threat | Adds a creature threat that fits the deck's white core.
- 1 Dawnhand Eulogist | threat | Serves as a high-impact finishing threat.
- 1 Foggy Swamp Hunters | threat | Adds a creature threat to pressure opponents.
- 1 Gollum, Silent Slinker // Meager Meal | threat | Provides a black creature threat with additional flexibility.
- 1 Lo and Li, Twin Tutors | threat | Adds a legendary creature threat to the deck's board plan.
- 1 Minwu, White Mage | threat | Provides a creature threat within the deck's lifegain-oriented colors.
- 1 Rabaroo Troop | threat | Adds a board-facing creature threat.
- 1 Reaping Willow | threat | Provides a resilient-looking creature threat for longer games.
- 1 Rooftop Percher | threat | Adds another creature threat to maintain pressure.
- 1 Shattered Angel | threat | Provides an Angel threat that suits the deck's theme.
- 1 Victory's Herald | threat | Provides a top-end creature threat for closing games.
- 1 Frodo, Sauron's Bane | wincon | Provides one of the deck's dedicated finishing routes.
- 1 Grave Venerations | wincon | Provides a dedicated route to end the game.
- 1 Lyra Dawnbringer | wincon | Provides a powerful finishing creature for the lifegain shell.
- 1 Austere Command | wipe | Provides a flexible reset button against developed boards.
- 1 Fumigate | wipe | Provides a full-board reset that fits the lifegain plan.
- 1 Vanquish the Horde | wipe | Provides an efficient creature-board reset.

</details>

### 16. a tight budget, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 231 names.

Commander: Denethor, Ruling Steward.

Grade: baseline, score 0.29, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $6.87 to buy, $34.13 the whole deck.

**Summary:** Denethor, Ruling Steward leads a budget-minded Orzhov aristocrats shell drawn primarily from the library. The deck develops with creatures, artifact ramp, and steady card access, then uses protection, focused removal, and board wipes to keep opposing boards manageable. It closes through Al Bhed Salvagers, Archfiend of Ifnir, and Grave Venerations, while giving up some speed and raw card quality in exchange for a creature-heavy, synergy-driven plan.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: There is a real aristocrats core (Bartolomé del Presidio, Yahenni, Woe Strider, Ayli, Old Flitterfang, Skullport Merchant, Erebos, Baron Bertram Graywater) that pairs with the commander's death trigger, but three symmetrical board wipes plus a commander-protection subtheme of equipment and combat tricks pull against a creature-count-hungry sacrifice plan.
- PLAN theme_fit=yes: It is an Orzhov Commander deck led by a death-trigger commander with a budget, common-card, library-first feel and a low-interaction bracket 2 power level, which is what was asked.
- PLAN useful_as_built=yes: Roughly 37 lands with heavy cheap ramp, a low curve, plenty of draw, and many repeatable drain-on-death payoffs means it casts its spells and has clear routes to close a game.
- PLAN summary_honest=partly: The ramp, draw, removal and wipe description matches the list, but naming Al Bhed Salvagers, Archfiend of Ifnir and Grave Venerations as the closers misrepresents where the deck's damage actually comes from (the commander plus the drain-on-death payoffs), and it never mentions the aristocrats sacrifice engine that is the deck's real plan.
- [WARN] `not_owned`: Woe Strider: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Yahenni, Undying Partisan: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Aron, Benalia's Ruin: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ayli, Eternal Pilgrim: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Baron Bertram Graywater: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bartolomé del Presidio: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Erebos, Bleak-Hearted: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Lord Skitter's Butcher: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Old Flitterfang: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Shadowheart, Dark Justiciar: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Sivriss, Nightmare Speaker: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Skullport Merchant: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Umbral Collar Zealot: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vampire Gourmand: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 2.94 over 62 nonland cards
- [INFO] `mana_pass`: the builder moved 4 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 5 Plains | land | Provides a reliable white mana source.
- 13 Plains | land | Provides reliable untapped white mana.
- 15 Swamp | land | Provides reliable untapped black mana.
- 1 Command Tower | land | Provides flexible command-zone color fixing.
- 1 Exotic Orchard | land | Adds flexible color fixing.
- 1 Grand Coliseum | land | Supplies both deck colors when needed.
- 1 Spire of Industry | land | Provides flexible mana alongside the artifact package.
- 1 Sol Ring | ramp | Efficiently accelerates the deck's mana development.
- 1 Arcane Signet | ramp | Reliable color-fixing mana rock.
- 1 Fellwar Stone | ramp | Low-cost mana acceleration.
- 1 Thought Vessel | ramp | Adds colorless mana development.
- 1 White Auracite | ramp | Low-cost artifact ramp from the library.
- 1 Blitzball | ramp | Supports early mana development.
- 1 Inherited Envelope | ramp | Adds another low-cost ramp piece.
- 1 Ring of the Lucii | ramp | Provides mana acceleration for the midgame.
- 1 Relic of Legends | ramp | Adds flexible mana production.
- 1 Springleaf Drum | ramp | Provides inexpensive early acceleration.
- 1 Ahriman | draw | Creature-based card access for the deck.
- 1 Circle of Power | draw | Adds card access from the library.
- 1 Exemplar of Light | draw | Creature-based card access that supports board development.
- 1 Idol of Oblivion | draw | Repeatable artifact-based card access.
- 1 Inspiring Overseer | draw | Provides card access while adding a creature.
- 1 Mask of Memory | draw | Equipment-based card access for attacking creatures.
- 1 Nasty End | draw | Efficient black card access.
- 1 Painful Truths | draw | Low-cost card access.
- 1 Skullclamp | draw | A strong draw tool for a creature-focused deck.
- 1 Tome of Legends | draw | Adds steady artifact-based card access.
- 1 Wall of Omens | draw | Early defense paired with card access.
- 1 Bastion Protector | interaction | Helps protect the commander and key creatures.
- 1 Gift of Immortality | interaction | Protects an important creature-based piece.
- 1 Swiftfoot Boots | interaction | Protects a key creature from opposing interaction.
- 1 Take Up the Shield | interaction | Flexible protection for a key creature.
- 1 Together Forever | interaction | Supports the deck's creature-focused plan through protection.
- 1 Unbreakable Formation | interaction | Protects the board at a pivotal moment.
- 1 Bitter Triumph | removal | Flexible single-target answer.
- 1 Claim the Precious | removal | Efficient black single-target removal.
- 1 Crib Swap | removal | Versatile instant-speed creature answer.
- 1 Fiend Hunter | removal | Creature-based removal that contributes to the board.
- 1 Generous Gift | removal | Broad answer to troublesome permanents.
- 1 Infernal Grasp | removal | Straightforward creature removal.
- 1 Swords to Plowshares | removal | Efficient creature removal.
- 1 Austere Command | wipe | Flexible reset option for difficult boards.
- 1 Dusk // Dawn | wipe | Board reset with a later creature-focused follow-up.
- 1 Fumigate | wipe | Reliable creature-board reset.
- 1 Arcade Cabinet | synergy | Artifact support for the deck's synergy package.
- 1 Gollum the Abandoned | synergy | Supports the requested aristocrats-oriented synergy plan.
- 1 Gollum, Patient Plotter | synergy | Adds another creature synergy piece for the plan.
- 1 Gríma Wormtongue | synergy | Supports the deck's creature-synergy core.
- 1 Heirloom Auntie | synergy | Low-cost creature synergy piece.
- 1 Joo Dee, One of Many | synergy | Adds to the creature-based synergy package.
- 1 Nimble Hobbit | synergy | Low-cost support for the synergy plan.
- 1 Phantom Train | synergy | Artifact-based support for the deck's synergies.
- 1 Woe Strider | synergy | Inexpensive creature support for the aristocrats plan.
- 1 Yahenni, Undying Partisan | synergy | Creature-based support for the aristocrats strategy.
- 1 Aron, Benalia's Ruin | threat | Low-cost legendary creature that supports the board plan.
- 1 Ayli, Eternal Pilgrim | threat | Efficient creature threat for the deck's core plan.
- 1 Baron Bertram Graywater | threat | Creature threat that fits the deck's colors and budget.
- 1 Bartolomé del Presidio | threat | Low-cost creature threat for the aristocrats shell.
- 1 Erebos, Bleak-Hearted | threat | Durable black threat for the midgame.
- 1 Lord Skitter's Butcher | threat | Creature threat that builds the deck's board presence.
- 1 Old Flitterfang | threat | Budget creature threat for the core strategy.
- 1 Shadowheart, Dark Justiciar | threat | Adds a black creature threat to the board.
- 1 Sivriss, Nightmare Speaker | threat | Low-cost creature threat that fits the deck's plan.
- 1 Skullport Merchant | threat | Creature threat that suits a value-oriented board.
- 1 Umbral Collar Zealot | threat | Inexpensive creature threat for the aristocrats shell.
- 1 Vampire Gourmand | threat | Low-cost black creature threat.
- 1 Al Bhed Salvagers | wincon | One of the deck's designated finishing cards.
- 1 Archfiend of Ifnir | wincon | A designated finisher that provides a powerful closing threat.
- 1 Grave Venerations | wincon | A designated finishing card for the late game.

</details>

### 17. upgrade a precon, owned first

Format: Commander. Theme: goblins. Pool: owned_first. Shortlist: 198 names.

Commander: Zada, Hedron Grinder.

Grade: typical, score 0.50, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $84.39 to buy, $258.98 the whole deck.

**Summary:** Zada leads a red Goblin Storm deck that develops through Goblin synergy, card draw, and bursts of mana before turning a dense spell turn into pressure. Great Train Heist and Earthquake are the chief finishing cards, while the Goblin removal package and board wipes help keep opponents from stabilizing. The deck gives up broader color access and relies on its commander and Goblin core to make its strongest turns.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core is coherent — Zada plus cheap cantrips/pump, wide goblin token makers, lords, and storm accelerants like Seething Song, Battle Hymn, Mana Geyser and Past in Flames — but Earthquake and Blasphemous Act as named finishers/wipes pull against a board of 1/1 goblins the rest of the deck works to build.
- PLAN theme_fit=yes: Mono-red goblins with token swarms, storm-style rituals and Grapeshot/Empty the Warrens at a bracket-3 power level is exactly the goblin-storm upgrade asked for, and no specific commander was mandated.
- PLAN useful_as_built=yes: 38 lands in mono-red plus Sol Ring, Skirk Prospector, Ruby Medallion and Fable give a smooth curve, and there are several independent ways to close (Krenko, Impact Tremors/Boggart Shenanigans, Goblin Bombardment, Great Train Heist, Zada pump).
- PLAN summary_honest=partly: It correctly names the goblin core, mana bursts, and mono-red tradeoff, but calling Earthquake a chief finisher and the symmetrical wipes 'help' overstates cards that mostly destroy the deck's own token board, and it soft-pedals that the real kill is Zada-copied pump spells across wide goblins.
- [INFO] `curve_summary`: average mana value 2.54 over 61 nonland cards
- [WARN] `profile_off_band`: the interaction count is 1, and bracket 3 wants 4 to 12
- [INFO] `bracket_cut`: to hold bracket 3, the builder cut 1 card: Storm-Kiln Artist (the combo Storm-Kiln Artist + Haze of Rage), and added 1 basic land

<details><summary>The deck list</summary>

- 1 Ancestors' Aid | other | Kept from the precon to preserve its Goblin Storm core.
- 1 Arena of Glory | land | Part of the retained red mana base.
- 1 Battle Hymn | ramp | A retained burst of mana for the Goblin-focused plan.
- 1 Blasphemous Act | wipe | Retained as a board-clearing option.
- 1 Boggart Shenanigans | removal | Retained Goblin-themed removal.
- 1 Castle Embereth | land | Part of the retained red mana base.
- 1 Chaos Warp | removal | Retained flexible removal.
- 1 Conspicuous Snoop | synergy | Retained Goblin synergy.
- 1 Crimson Wisps | draw | Retained card draw for the spell-heavy plan.
- 1 Daring Discovery | other | Kept from the precon to preserve its Goblin Storm core.
- 1 Den of the Bugbear | land | Part of the retained red mana base.
- 1 Dragon Fodder | other | Kept from the precon to preserve its Goblin Storm core.
- 1 Dwarven Mine | land | Part of the retained red mana base.
- 1 Empty the Warrens | other | Kept from the precon to preserve its Goblin Storm core.
- 1 Expedite | draw | Retained card draw for the spell-heavy plan.
- 1 Faithless Looting | draw | Retained card draw to keep cards moving.
- 1 Fists of Flame | draw | Retained card draw for the spell-heavy plan.
- 1 Forgotten Cave | land | Part of the retained red mana base.
- 1 Fountainport | land | Part of the retained mana base.
- 1 Frontline Heroism | other | Kept from the precon to preserve its Goblin Storm core.
- 1 Gempalm Incinerator | removal | Retained Goblin-themed removal.
- 1 General Kreat, the Boltbringer | synergy | Retained Goblin synergy.
- 1 Goblin Bombardment | removal | Retained Goblin-themed removal.
- 1 Goblin Bushwhacker | synergy | Retained Goblin synergy.
- 1 Goblin Chieftain | synergy | Retained Goblin synergy.
- 1 Goblin Dark-Dwellers | threat | Retained as a Goblin threat.
- 1 Goblin Lackey | synergy | Retained Goblin synergy.
- 1 Goblin Matron | synergy | Retained Goblin synergy.
- 1 Goblin Negotiation | removal | Retained Goblin-themed removal.
- 1 Goblin Trashmaster | removal | Retained Goblin-themed removal.
- 1 Goblin Warchief | synergy | Retained Goblin synergy.
- 1 Grapeshot | removal | Retained efficient removal for the storm theme.
- 1 Great Train Heist | wincon | A retained finisher for closing games.
- 1 Grenzo, Havoc Raiser | synergy | Retained Goblin synergy.
- 1 Haze of Rage | other | Kept from the precon to preserve its Goblin Storm core.
- 1 Hidden Volcano | land | Part of the retained red mana base.
- 1 Howlsquad Heavy | ramp | Retained mana support.
- 1 Idol of Oblivion | draw | Retained card draw.
- 1 Impact Tremors | other | Kept from the precon to preserve its Goblin Storm core.
- 1 Kher Keep | land | Part of the retained mana base.
- 1 Krenko's Command | other | Kept from the precon to preserve its Goblin Storm core.
- 1 Krenko, Mob Boss | threat | Retained as a Goblin threat.
- 1 Mana Geyser | ramp | Retained burst mana for explosive turns.
- 1 Mogg War Marshal | synergy | Retained Goblin synergy.
- 1 Mountain | land | The basic red foundation of the mana base.
- 1 Past in Flames | other | Kept from the precon to preserve its Goblin Storm core.
- 1 Quest for the Goblin Lord | synergy | Retained Goblin synergy.
- 1 Reliquary Tower | land | Part of the retained mana base.
- 1 Roaming Throne | other | Kept from the precon to preserve its Goblin Storm core.
- 1 Ruby Medallion | other | Kept from the precon to preserve its Goblin Storm core.
- 1 Rundvelt Hordemaster | synergy | Retained Goblin synergy.
- 1 Sazacap's Brew | draw | Retained card draw.
- 1 Searslicer Goblin | synergy | Retained Goblin synergy.
- 1 Seething Song | ramp | Retained burst mana for explosive turns.
- 1 Shinka, the Bloodsoaked Keep | land | Part of the retained red mana base.
- 1 Siege-Gang Commander | removal | Retained Goblin-themed removal.
- 1 Siege-Gang Lieutenant | removal | Retained Goblin-themed removal.
- 1 Skirk Prospector | ramp | Retained Goblin mana support.
- 1 Skullclamp | draw | Retained card draw.
- 1 Sol Ring | ramp | Retained early mana acceleration.
- 1 Spreading Insurrection | other | Kept from the precon to preserve its Goblin Storm core.
- 1 Swiftfoot Boots | interaction | Retained protective interaction.
- 1 Throne of Eldraine | ramp | Retained mana support.
- 1 Vandalblast | wipe | Retained as a board-clearing option.
- 1 War Room | land | Part of the retained mana base.
- 1 Warren Torchmaster | synergy | Retained Goblin synergy.
- 1 Witch's Mark | draw | Retained card draw.
- 1 Abrade | removal | Efficient additional removal.
- 1 Ancestral Anger | draw | Low-cost card draw that supports the spell-focused plan.
- 1 Burst Lightning | removal | Additional low-cost removal.
- 1 Earthquake | wincon | An additional finisher that can also close stalled games.
- 1 Fable of the Mirror-Breaker // Reflection of Kiki-Jiki | ramp | Additional mana support that remains on the Goblin theme.
- 1 Lightning Bolt | removal | Efficient additional removal.
- 26 Mountain | land | The basic red foundation of the mana base.

</details>

### 18. upgrade a precon, any card

Format: Commander. Theme: turtles. Pool: any_card. Shortlist: 311 names.

Commander: Heroes in a Half Shell.

Grade: baseline, score 0.29, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $250.52 to buy, $250.52 the whole deck.

**Summary:** Heroes in a Half Shell leads a Turtle Power deck that keeps the precon’s broad cast of Turtles, allies, pizza, and arcade-themed cards intact while improving its mana and adding sharper draw and answers. Develop with the retained ramp and Turtle synergy, pressure the table with the deck’s threats, and close through its themed finishers—chief among them Raphael, the Muscle, Everything Pizza, and Dimension X Pizzasaur. The deck gives up some of the precon’s less focused creature package in exchange for smoother colors and more reliable interaction.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core is a coherent go-wide Turtles/allies board built around the commander's themed payoffs, but stacking Blasphemous Act, Vanquish the Horde, Turtles in Time and Cyclonic Rift in a deck whose plan is a wide creature board pulls against itself.
- PLAN theme_fit=yes: It is a Heroes in a Half Shell Turtle Power precon upgrade that retains the turtle/pizza/arcade cast and adds a handful of efficient staples, sitting comfortably at bracket 3.
- PLAN useful_as_built=yes: 39 lands with heavy five-color fixing (Command Tower, City of Brass, Mana Confluence, Grand Coliseum, Chromatic Lantern) support the curve, and there are clear themed finishers plus removal, even if ramp is light and the mana is slow.
- PLAN summary_honest=partly: The claims about smoother mana, themed finishers and added interaction match the list, but "sharper draw" oversells two draw pieces (Rhystic Study, April O'Neil) and the summary never mentions the several sweepers that hit the deck's own board.
- [INFO] `curve_summary`: average mana value 3.35 over 60 nonland cards
- [WARN] `profile_off_band`: the land count is 39, and bracket 3 wants 34 to 38
- [WARN] `profile_off_band`: the ramp count is 5, and bracket 3 wants 8 to 13
- [WARN] `profile_off_band`: the draw count is 2, and bracket 3 wants 8 to 14
- [WARN] `profile_off_band`: the removal count is 3, and bracket 3 wants 6 to 12
- [WARN] `profile_off_band`: the interaction count is 0, and bracket 3 wants 4 to 12
- [INFO] `mana_pass`: the builder moved 1 card of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 1 Ash Barrens | land | Retained from the precon as part of its land base.
- 1 Big Apple, 3 a.m. | land | Retained from the precon as part of its land base.
- 1 Cinder Glade | land | Retained from the precon as part of its land base.
- 1 City of Brass | land | Retained from the precon as part of its land base.
- 1 Command Tower | land | Retained from the precon as part of its land base.
- 1 Dragonskull Summit | land | Retained from the precon as part of its land base.
- 1 Grand Coliseum | land | Retained from the precon as part of its land base.
- 1 Hidden Hideout | land | Retained from the precon as part of its land base.
- 1 Rain-Slicked Copse | land | Retained from the precon as part of its land base.
- 1 Rootbound Crag | land | Retained from the precon as part of its land base.
- 1 Smoldering Marsh | land | Retained from the precon as part of its land base.
- 1 Sodden Verdure | land | Retained from the precon as part of its land base.
- 1 Sunken Hollow | land | Retained from the precon as part of its land base.
- 1 Thriving Moor | land | Retained from the precon as part of its land base.
- 1 Turtle Lair | land | Retained from the precon as part of its Turtle-focused mana base.
- 1 Undergrowth Stadium | land | Retained from the precon as part of its land base.
- 1 Vernal Fen | land | Retained from the precon as part of its land base.
- 1 Vibrant Cityscape | land | Retained from the precon as part of its land base.
- 1 Mana Confluence | land | Strengthens the mana base with an additional color source.
- 1 Forbidden Orchard | land | Strengthens the mana base with an additional color source.
- 1 Breeding Pool | land | Improves access to blue and green.
- 1 Hallowed Fountain | land | Improves access to white and blue.
- 1 Blood Crypt | land | Improves access to black and red.
- 3 Forest | land | Basic green sources retained for a stable land count.
- 3 Island | land | Basic blue sources retained for a stable land count.
- 4 Plains | land | Basic white sources retained for a stable land count.
- 3 Swamp | land | Basic black sources retained for a stable land count.
- 3 Mountain | land | Basic red sources retained for a stable land count.
- 1 April O'Neil, Live on the Scene | draw | Retained precon card that supplies draw.
- 1 Arcade Cabinet | other | Retained from the precon to preserve its original theme.
- 1 Arcane Signet | ramp | Retained precon ramp for earlier development.
- 1 Assassin's Trophy | removal | Retained precon removal.
- 1 Baxter, Fly in the Ointment | other | Retained from the precon to preserve its original theme.
- 1 Bebop, Skull & Crossbones | other | Retained from the precon to preserve its original theme.
- 1 Big Mother Mouser | other | Retained from the precon to preserve its original theme.
- 1 Blasphemous Act | wipe | Retained precon board wipe.
- 1 Casey Jones, Back Alley Brute | other | Retained from the precon to preserve its original theme.
- 1 Chromatic Lantern | ramp | Retained from the precon to support the mana base.
- 1 Coin of Mastery | other | Retained from the precon to preserve its original theme.
- 1 Continue? | other | Retained from the precon to preserve its original theme.
- 1 Cultivate | ramp | Retained precon ramp for dependable development.
- 1 Dimension X Pizzasaur | wincon | Retained precon finisher that keeps the deck's theme intact.
- 1 Donatello, the Brains | synergy | Retained Turtle synergy from the precon.
- 1 Double Jump // Flying Kick | other | Retained from the precon to preserve its original theme.
- 1 Electric Seaweed | other | Retained from the precon to preserve its original theme.
- 1 Endless Foot Assault | other | Retained from the precon to preserve its original theme.
- 1 Everything Pizza | wincon | Retained precon finisher that keeps the deck's theme intact.
- 1 Exploding Barrel | other | Retained from the precon to preserve its original theme.
- 1 Fast Forward | other | Retained from the precon to preserve its original theme.
- 1 Foot Chopper | other | Retained from the precon to preserve its original theme.
- 1 Game Over | other | Retained from the precon to preserve its original theme.
- 1 Here Comes a New Hero! | other | Retained from the precon to preserve its original theme.
- 1 High Score | other | Retained from the precon to preserve its original theme.
- 1 Irma, Part-Time Mutant | other | Retained from the precon to preserve its original theme.
- 1 Krang, the All-Powerful | other | Retained from the precon to preserve its original theme.
- 1 Leatherhead, Iron Gator | other | Retained from the precon to preserve its original theme.
- 1 Leonardo, the Balance | threat | Retained precon Turtle threat.
- 1 Lessons from Life | other | Retained from the precon to preserve its original theme.
- 1 Level Up | other | Retained from the precon to preserve its original theme.
- 1 Lita, Little Orphan Amphibian | synergy | Retained precon Turtle synergy.
- 1 Michelangelo, the Heart | synergy | Retained precon Turtle synergy.
- 1 Mole Module | other | Retained from the precon to preserve its original theme.
- 1 Mona Lisa, Science Geek | other | Retained from the precon to preserve its original theme.
- 1 Ninja Pizza | other | Retained from the precon to preserve its original theme.
- 1 Raphael, the Muscle | wincon | Retained precon finisher that remains central to the deck's closing plan.
- 1 Rat King, Pale Piper | other | Retained from the precon to preserve its original theme.
- 1 Ray Fillet, Wave Warrior | other | Retained from the precon to preserve its original theme.
- 1 Roadkill Rodney | other | Retained from the precon to preserve its original theme.
- 1 Rocksteady, Mutant Marauder | other | Retained from the precon to preserve its original theme.
- 1 Shellshock | other | Retained from the precon to preserve its original theme.
- 1 Shredder, Shadow Master | other | Retained from the precon to preserve its original theme.
- 1 Sol Ring | ramp | Retained precon ramp for faster development.
- 1 Special Move | other | Retained from the precon to preserve its original theme.
- 1 Splinter, the Mentor | other | Retained from the precon to preserve its original theme.
- 1 Steelbane Hydra | removal | Retained precon removal.
- 1 Super Combo | other | Retained from the precon to preserve its original theme.
- 1 Swift Demise | other | Retained from the precon to preserve its original theme.
- 1 Tempestra, Dame of Games | other | Retained from the precon to preserve its original theme.
- 1 Together Forever | other | Retained from the precon to preserve its original theme.
- 1 Tokka & Rahzar, Unsupervised | ramp | Retained precon Turtle ramp.
- 1 Turtles in Time | wipe | Retained precon board wipe.
- 1 Vanquish the Horde | wipe | Retained from the precon as another reset option.
- 1 Vigor | other | Retained from the precon to preserve its original theme.
- 1 Voracious Hydra | other | Retained from the precon to preserve its original theme.
- 1 Wave Goodbye | other | Retained from the precon to preserve its original theme.
- 1 Rhystic Study | draw | Adds a stronger ongoing draw option.
- 1 Swords to Plowshares | removal | Adds efficient removal.
- 1 Cyclonic Rift | wipe | Adds a flexible board-reset option.

</details>

### 19. the Hobbit family, two colours

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 174 names.

Commander: Thranduil, the Elvenking.

Grade: baseline, score 0.04, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $415.62 to buy, $415.62 the whole deck.

**Summary:** Thranduil, the Elvenking leads an Elf-centered creature deck that develops its mana early, keeps cards flowing, and uses a broad suite of removal and reactive tools to protect its position. The board is built through Elf creatures and legendary support, then closes with large marked finishers such as Troll of Khazad-dûm and the Witch-kings. Its tradeoff is that the creature-focused plan can need time to establish a meaningful board after a reset, so careful use of the draw and interaction pieces matters.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core is a green-based Elf creature board with ramp and card flow, but three symmetrical sweepers (Languish, Gnashing of Teeth, Raise the Palisade) and a large pile of one-off reactive spells pull against the go-wide small-Elf body of the list, leaving the tribal payoffs thinly supported.
- PLAN theme_fit=partly: It is a legal Commander deck led by the requested Thranduil, the Elvenking at a reasonable bracket-3 power level, but the "from the Hobbit set" constraint is only loosely honored, with a number of LOTR-set and generic staples (Orcish Bowmasters, Delighted Halfling, Arcane Signet, Mox Amber, Night's Whisper, Languish) mixed in.
- PLAN useful_as_built=yes: 36 lands across a Sultai base with signet, Mystic-class dorks and land fetch support a sensible curve, and the deck has plenty of creature-based and finisher win paths to actually close a game.
- PLAN summary_honest=partly: It correctly names the ramp, draw, removal and big finishers, but calls the deck "Elf-centered" when a sizable share of the creatures are Hobbits, spiders, trolls and Witch-kings, and it only obliquely hints at the fact that its own sweepers punish its creature plan.
- [INFO] `curve_summary`: average mana value 3.17 over 63 nonland cards
- [WARN] `profile_off_band`: the worst color's share of the sources it needs is 0.84, and bracket 3 wants 0.85 or more (sources of requirement: U 16.5 of 19, B 19.25 of 23, G 19 of 22)
- [INFO] `mana_pass`: the builder moved 4 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 1 Elven Passage | land | Fills a land slot in the mana base.
- 1 Elvenking's Halls | land | Fills a land slot in the mana base.
- 1 Hobbit Hole | land | Fills a land slot in the mana base.
- 1 Minas Morgul, Dark Fortress | land | Fills a land slot in the mana base.
- 1 Mirkwood | land | Fills a land slot in the mana base.
- 1 Rivendell | land | Fills a land slot in the mana base.
- 1 The Black Gate | land | Fills a land slot in the mana base.
- 1 The Shire | land | Fills a land slot in the mana base.
- 11 Forest | land | Provides a dependable green source.
- 8 Island | land | Provides a dependable blue source.
- 9 Swamp | land | Provides a dependable black source.
- 1 Arcane Signet | ramp | Efficiently supports the deck's mana development.
- 1 Delighted Halfling | ramp | Provides early mana development on a creature.
- 1 Elven Chorus | ramp | Adds a dedicated mana-development piece.
- 1 Elvish Archdruid | ramp | Provides Elf-based mana development.
- 1 Elvish Mystic | ramp | Provides early mana development.
- 1 Mox Amber | ramp | Adds a low-cost mana-development piece.
- 1 Orcrist, Goblin-cleaver | ramp | Contributes to the ramp package.
- 1 Wayfarer's Bauble | ramp | Provides dependable mana development.
- 1 Wood Elves | ramp | Provides Elf-based mana development.
- 1 Woodland Weavemaster | ramp | Provides mana development while supporting the Elf core.
- 1 Beorn the Fierce | draw | Adds a card-advantage creature.
- 1 Bilbo Baggins, Burglar // Take a Glance | draw | Adds card access through a flexible card.
- 1 Elvish Visionary | draw | Provides card access on an Elf body.
- 1 Fateful Discovery | draw | Adds a dedicated card-advantage piece.
- 1 Gandalf, Shadow's Foe | draw | Adds card advantage to the deck's creature suite.
- 1 Hithlain Knots | draw | Provides instant-speed card access.
- 1 Ithilien Kingfisher | draw | Adds card advantage on a creature.
- 1 Key to the Side-Door | draw | Provides a dedicated card-advantage artifact.
- 1 Night's Whisper | draw | Adds efficient card access.
- 1 Palantír of Orthanc | draw | Provides a lasting card-advantage piece.
- 1 Lórien Revealed | draw | Adds another dedicated card-access spell.
- 1 Confusticate and Bebother | interaction | Provides reactive interaction.
- 1 Bilbo's Ring | interaction | Adds an interaction piece that stays with the board.
- 1 Elrond, Moon-Reader | interaction | Provides interaction on a legendary Elf.
- 1 Dwarven Mattock | interaction | Adds artifact-based interaction.
- 1 Mithril Coat | interaction | Provides protective interaction.
- 1 My Precious // Allure of Power | interaction | Adds flexible interaction.
- 1 Old Fat Spider Can't See Me | interaction | Provides a persistent interaction piece.
- 1 Stern Scolding | interaction | Provides low-cost reactive interaction.
- 1 Bilbo's Deadly Slice | removal | Adds a dedicated removal spell.
- 1 Bitter Downfall | removal | Adds a dedicated removal spell.
- 1 Crude Bent Blade | removal | Provides removal from an artifact slot.
- 1 Enchanted River's Grasp | removal | Adds enchantment-based removal.
- 1 Giant's Boulder | removal | Provides artifact-based removal.
- 1 Merciless Executioner | removal | Adds removal on a creature.
- 1 Orcish Bowmasters | removal | Provides removal on a creature.
- 1 Quarrel | removal | Adds instant-speed removal.
- 1 The Black Arrow | removal | Provides a dedicated removal equipment.
- 1 Gnashing of Teeth | wipe | Provides a board-clearing reset.
- 1 Languish | wipe | Provides a board-clearing reset.
- 1 Raise the Palisade | wipe | Provides another broad board reset.
- 1 Celeborn the Wise | synergy | Supports the deck's Elf-centered creature plan.
- 1 Galadhrim Guide | synergy | Supports the Elf-centered creature core.
- 1 Galion, Elvenking's Butler | synergy | Reinforces the legendary Elf contingent.
- 1 Mirkwood Meditator | synergy | Supports the Elf-centered creature core.
- 1 Mirkwood Nurturer | synergy | Supports the Elf-centered creature core.
- 1 Mirkwood Pathmaker | synergy | Supports the Elf-centered creature core.
- 1 Supper for Spiders | synergy | Provides the shortlist's dedicated synergy piece.
- 1 Boughside Wanderers | threat | Adds an Elf creature to apply board pressure.
- 1 Cantankerous Keepers | threat | Adds an Elf creature to apply board pressure.
- 1 Colossal Whale | threat | Adds a large finishing creature.
- 1 Elven Raft-Steerer | threat | Adds an Elf creature to develop the board.
- 1 Elvenking's Harper | threat | Adds an Elf creature to develop the board.
- 1 Grey Havens Navigator | threat | Adds an Elf creature to develop the board.
- 1 Guardian of the Halls | threat | Adds an Elf creature to hold and pressure the board.
- 1 Lothlórien Lookout | threat | Adds an Elf creature to develop the board.
- 1 Nimrodel Watcher | threat | Adds an Elf creature to develop the board.
- 1 Old Fat Spider | threat | Adds a large finishing creature.
- 1 Thranduil, Sindarin Liege // Silvan Rally | threat | Adds a flexible legendary Elf threat.
- 1 Troll of Khazad-dûm | wincon | Provides a marked finishing route.
- 1 Witch-king of Angmar | wincon | Provides a marked finishing route.
- 1 Witch-king, Bringer of Ruin | wincon | Provides a marked finishing route.
- 1 Pelargir Survivor | ramp | the mana pass added it to bring the mana base inside the power level

</details>

### 20. the Hobbit family, a delegated commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 77 names.

Commander: Smaug the Magnificent.

Grade: baseline, score 0.06, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $455.57 to buy, $455.57 the whole deck.

**Summary:** Smaug the Magnificent leads a red, Hobbit-set dragon deck that develops its mana, builds a creature board, and uses a broad equipment package to turn combat pressure into a win. Cavern-Hoard Dragon, Desert Were-Worm, and Smaug, the Great Calamity are chief among the closing threats, backed by Dáin Ironfoot and a supporting force of Dwarves, Goblins, Orcs, and other creatures. The deck gives up some specialized card selection for a cohesive themed roster, leaning on its mana development, removal, and board resets to create the opening for its finishers.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The bulk of the list is a coherent mono-red creatures-plus-equipment beatdown with ramp and removal, but the stated dragon payoff is nearly absent and the board wipes (Desolation of Smaug, Call Forth the Tempest, Easy Pickings) work against its own wide creature board and equipped attackers.
- PLAN theme_fit=no: The person asked for a dragons deck and the list contains only about two dragons besides the commander, with the remaining creatures being dwarves, goblins, orcs and assorted bodies wearing equipment.
- PLAN useful_as_built=yes: Mono-red with 30 Mountains plus utility lands, cheap ramp, a reasonable curve of threats, removal and a clear combat-damage win path means it can be picked up and played as it stands.
- PLAN summary_honest=partly: It accurately describes the equipment package, ramp, removal and wipes, and even concedes the roster is dwarves/goblins/orcs, but calling it a "dragon deck" and listing Desert Were-Worm among the dragon finishers overstates a tribal theme the list does not carry.
- [INFO] `curve_summary`: average mana value 3.11 over 63 nonland cards
- [WARN] `profile_off_band`: the draw count is 6, and bracket 3 wants 8 to 14

<details><summary>The deck list</summary>

- 30 Mountain | land | Primary red mana base for casting Smaug and the deck’s red spells.
- 1 Dragon-Cursed Halls | land | Dragon-themed land slot that adds to the mana base.
- 1 Elven Passage | land | Nonbasic land slot that broadens the mana base.
- 1 Hobbit Hole | land | Nonbasic land slot for the deck’s land count.
- 1 Rogue's Passage | land | Utility land slot for the creature combat plan.
- 1 The Lonely Mountain | land | Mountain land that reinforces the red mana base and theme.
- 1 Treasure Vault | land | Artifact land that supplies a utility mana slot.
- 1 Balin, Loremaster | draw | Creature-based card advantage for keeping resources available.
- 1 Key to the Side-Door | draw | Artifact card-advantage piece for the deck’s longer games.
- 1 Palantír of Orthanc | draw | Legendary artifact card-advantage piece.
- 1 Ragged Short Spear | draw | Equipment that fills a card-advantage slot while supporting creatures.
- 1 Thrór's Map | draw | Legendary artifact card-advantage piece.
- 1 Óin the Brave | draw | Creature-based card advantage that fits the Hobbit roster.
- 1 Bilbo's Ring | interaction | Equipment-based interaction for protecting or contesting key turns.
- 1 Dwarven Mattock | interaction | Equipment interaction that remains useful alongside a creature-heavy plan.
- 1 Mithril Coat | interaction | Protective interaction for important creatures and equipment targets.
- 1 The One Ring | interaction | Legendary artifact interaction for pivotal board states.
- 1 Gandalf, Goblins' Bane // Flameshape | interaction | Flexible legendary and Adventure card used as an interaction slot.
- 1 Smaug's Fury | interaction | Instant-speed interaction that supports reactive turns.
- 1 Tidings of War | interaction | Sorcery interaction for contesting opposing plans.
- 1 Dori, Bearer of Friends | interaction | Legendary creature interaction slot within the Hobbit-themed roster.
- 1 Arcane Signet | ramp | Efficient fixing and acceleration toward Smaug and the top end.
- 1 Bag End Banquet | ramp | Artifact acceleration for advancing the mana plan.
- 1 Burn, Burn, Tree and Fern | ramp | Saga-based acceleration that supports the deck’s development.
- 1 Dragon's Desire | ramp | Dragon-themed acceleration for reaching the deck’s major threats.
- 1 Long-Bodied Grey Dog | ramp | Creature-based acceleration that develops the board early.
- 1 Mox Amber | ramp | Fast legendary-focused acceleration for early development.
- 1 Orcrist, Goblin-cleaver | ramp | Equipment-based acceleration that also suits the creature plan.
- 1 The Misty Mountains Cold | ramp | Saga acceleration for building toward larger spells.
- 1 The Reaver Cleaver | ramp | Equipment ramp piece for turning creature pressure into mana development.
- 1 Wayfarer's Bauble | ramp | Low-cost mana development that helps secure land drops.
- 1 Battle-Scarred Goblin | removal | Creature removal that contributes to the deck’s board presence.
- 1 Fire of Orthanc | removal | Direct removal for troublesome opposing permanents.
- 1 Gandalf, Spark Starter | removal | Legendary creature removal that fits the Hobbit theme.
- 1 Giant's Boulder | removal | Artifact removal piece for answering opposing threats.
- 1 Goblin Cratermaker | removal | Creature-based removal with a low-curve body.
- 1 Improvised Club | removal | Instant removal for answering threats efficiently.
- 1 Pinecone Strike | removal | Instant removal that preserves flexibility on opposing turns.
- 1 Smite the Deathless | removal | Instant removal for handling resilient opposing threats.
- 1 The Black Arrow | removal | Equipment-based removal that supports the combat plan.
- 1 Call Forth the Tempest | wipe | Board-reset option for recovering from crowded battlefields.
- 1 Desolation of Smaug | wipe | Smaug-themed board reset for clearing opposing development.
- 1 Glóin the Mighty // Easy Pickings | wipe | Legendary creature and Adventure board-reset option.
- 1 Last Light of Durin's Day | synergy | Core synergy piece for the deck’s thematic creature plan.
- 1 Andúril, Flame of the West | synergy | Equipment synergy that rewards maintaining a creature presence.
- 1 Andúril, Narsil Reforged | synergy | Legendary equipment synergy for the combat-focused shell.
- 1 Glamdring | synergy | Legendary equipment that strengthens the deck’s creature theme.
- 1 Long-Lost Lances | synergy | Equipment synergy for supporting attackers.
- 1 Sting, Bilbo's Sword | synergy | Legendary equipment synergy in the Hobbit-themed package.
- 1 Well-Worn Spatula | synergy | Additional equipment synergy for the creature-heavy plan.
- 1 Dáin Ironfoot | threat | Finisher-marked legendary creature that adds a substantial combat threat.
- 1 Dwarven Mauler | threat | Creature threat that helps establish a combat board.
- 1 Dwarven Warriors | threat | Creature threat that contributes to sustained battlefield pressure.
- 1 Goblin-town Flunkies | threat | Low-curve creature threat for developing early pressure.
- 1 Gundabad Opportunist | threat | Creature threat that fills out the Hobbit-themed combat force.
- 1 Guttersnipe | threat | Creature threat that benefits from the deck’s instants and sorceries.
- 1 Iron Hills Stalwart | threat | Dwarf creature threat for a resilient combat presence.
- 1 Misty Mountains Raider | threat | Goblin creature threat that adds to the attacking force.
- 1 Oliphaunt | threat | Large creature threat for applying meaningful combat pressure.
- 1 Olog-hai Crusher | threat | Troll creature threat for the deck’s upper-end combat plan.
- 1 Orcish Siegemaster | threat | Orc creature threat that broadens the board presence.
- 1 Snowslope Hunter | threat | Goblin creature threat that supports steady pressure.
- 1 Cavern-Hoard Dragon | wincon | Dragon finisher and one of the deck’s primary closing threats.
- 1 Desert Were-Worm | wincon | Dragon Wurm finisher for closing games through combat.
- 1 Smaug, the Great Calamity // Spew Flame | wincon | Smaug-themed Dragon finisher that reinforces the deck’s central plan.
- 1 Bombur, Gentle Dreamer | other | Legendary Hobbit-themed creature that rounds out the supporting cast.
- 1 Bothersome Noisemaker | other | Goblin creature that adds another body to the board.
- 1 Getaway Barrel | other | Utility artifact that supports the deck outside its primary packages.
- 1 Old Thrush | other | Creature support piece that fills out the thematic roster.
- 1 Troop of Ponies | other | Creature support piece that adds to the Hobbit-set board presence.

</details>

### 21. the Hobbit family, mana from outside

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 156 names.

Commander: Smaug the Impenetrable.

Grade: bad, score 0.02, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $749.31 to buy, $749.31 the whole deck.

**Summary:** Smaug the Impenetrable leads a threat-heavy deck that develops its mana, keeps cards flowing, and uses a broad creature force backed by equipment and Goblin-themed support. The deck aims to clear away resistance with targeted answers or sweeping effects, then end games through its marked finishers, chief among them the other Smaug cards and Cavern-Hoard Dragon. It gives up some reactive flexibility for a proactive board presence and a concentrated high-end finish.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The list leans on a big-creature/dragon top end with some ramp and card flow, but the goblin package, the equipment cluster, and the assorted standalone fatties pull in separate directions rather than building one engine, and the summary itself describes it only as a "threat-heavy" pile.
- PLAN theme_fit=partly: The commander, colors, format, and rough bracket match, and the lands/rocks exception is used as permitted, but a sizable number of non-land, non-rock cards (The One Ring, Palantír of Orthanc, Andúril, Glamdring, Sauron the Lidless Eye, Witch-king of Angmar, Inferno Titan, Guttersnipe, Night's Whisper) fall outside the requested Hobbit set.
- PLAN useful_as_built=yes: Twenty-seven lands in two colors with duals, fetches, and roughly ten ramp pieces will cast this curve, and there are ample finishers and removal to actually close a game, even if the top end is heavy.
- PLAN summary_honest=yes: Every element it names — ramp, draw, targeted removal, two sweepers, equipment, goblin cards, and the Smaug/Cavern-Hoard high end — is present (Cavern-Hoard Dragon being the one named card I cannot find in the list), and it openly concedes the weak reactive suite and the quality-model shortfalls.
- [INFO] `curve_summary`: average mana value 3.45 over 62 nonland cards
- [INFO] `mana_pass`: the builder moved 2 cards of the mana base to bring the deck inside its power level
- [WARN] `outside_requested_set`: the sets you named do not hold 12 cards: Arid Mesa, Blood Crypt, Bloodstained Mire, Command Tower, Dragonskull Summit, Exotic Orchard, Marsh Flats, Polluted Delta, Scalding Tarn, Urborg, Tomb of Yawgmoth, Verdant Catacombs, Wooded Foothills

<details><summary>The deck list</summary>

- 13 Mountain | land | Primary red mana source with reliable untapped access.
- 12 Swamp | land | Primary black mana source with reliable untapped access.
- 1 Blood Crypt | land | Dual-color land that supports both primary colors.
- 1 Bloodstained Mire | land | Fetch land that improves access to both primary colors.
- 1 Command Tower | land | Reliable multicolor source for the commander deck.
- 1 Dragonskull Summit | land | Additional dual-color land for consistent casting.
- 1 Exotic Orchard | land | Flexible multicolor source.
- 1 Arid Mesa | land | Fetch land that can access the red base.
- 1 Marsh Flats | land | Fetch land that can access the black base.
- 1 Polluted Delta | land | Fetch land that can access the black base.
- 1 Verdant Catacombs | land | Fetch land that can access the black base.
- 1 Wooded Foothills | land | Fetch land that can access the red base.
- 1 Scalding Tarn | land | Fetch land that can access the red base.
- 1 Urborg, Tomb of Yawgmoth | land | Additional black source that strengthens the mana base.
- 1 Arcane Signet | ramp | Efficient color fixing and acceleration.
- 1 Bag End Banquet | ramp | Artifact acceleration from the requested card pool.
- 1 Bolg's Company | ramp | Creature-based acceleration that advances the board.
- 1 Burn, Burn, Tree and Fern | ramp | Ramp piece that helps reach the deck's larger cards.
- 1 Dragon's Desire | ramp | Additional acceleration for the dragon-focused top end.
- 1 Mox Amber | ramp | Low-cost acceleration that helps develop quickly.
- 1 Orcrist, Goblin-cleaver | ramp | Equipment-based ramp from the shortlist.
- 1 Smaug the Magnificent | ramp | Dragon ramp piece that contributes to the creature plan.
- 1 The Reaver Cleaver | ramp | Ramp equipment that supports the deck's threat plan.
- 1 Wayfarer's Bauble | ramp | Early mana development that improves later turns.
- 1 Balin, Loremaster | draw | Creature-based card advantage.
- 1 Key to the Side-Door | draw | Artifact card-advantage piece.
- 1 Night's Whisper | draw | Compact card-advantage spell.
- 1 Palantír of Orthanc | draw | Persistent card-advantage artifact.
- 1 Rage into the Valley | draw | Additional card-advantage spell.
- 1 Ragged Short Spear | draw | Equipment that fills a card-advantage slot.
- 1 Reverent Howl | draw | Instant-speed card advantage.
- 1 The Master of Lake-town | draw | Creature-based card advantage.
- 1 The Sackville-Bagginses | draw | Additional creature-based card advantage.
- 1 Thrór's Map | draw | Artifact card advantage that supports longer games.
- 1 Óin the Brave | draw | Creature-based draw option for maintaining resources.
- 1 Azog, Moria's Ruin | removal | Creature removal attached to a legendary threat.
- 1 Battle-Scarred Goblin | removal | Low-cost creature removal.
- 1 Bilbo's Deadly Slice | removal | Direct removal spell.
- 1 Bitter Downfall | removal | Flexible removal spell.
- 1 Bolg of the North | removal | Legendary creature that provides removal.
- 1 Crude Bent Blade | removal | Equipment-based removal option.
- 1 Fire of Orthanc | removal | Sorcery-speed removal.
- 1 Front Porch Sentries | removal | Creature-based removal option.
- 1 Gandalf, Spark Starter | removal | Legendary removal creature.
- 1 Bolg, Erebor's Reckoning | wipe | Creature-based board reset.
- 1 Desolation of Smaug | wipe | Thematic board reset for recovering difficult boards.
- 1 Bilbo's Ring | interaction | Protection and utility from the equipment suite.
- 1 Dwarven Mattock | interaction | Flexible equipment-based interaction.
- 1 Mithril Coat | interaction | Protective interaction for key creatures.
- 1 My Precious // Allure of Power | interaction | Versatile interaction card with an equipment front.
- 1 The One Ring | interaction | High-impact artifact interaction.
- 1 Supper for Spiders | synergy | Dedicated synergy card from the shortlist.
- 1 Fearsome Goblin Pair | synergy | Creature that supports the deck's Goblin contingent.
- 1 Goblin-town Flunkies | synergy | Goblin body that reinforces the creature theme.
- 1 Gollum the Abandoned | synergy | Thematic legendary creature supporting the core creature plan.
- 1 Great Goblin, Foul-Hearted | synergy | Legendary Goblin that deepens the Goblin package.
- 1 Guttersnipe | synergy | Goblin inclusion that supports the deck's spell-heavy support suite.
- 1 Down, Down to Goblin-town | synergy | Thematic saga that supports the Goblin-focused package.
- 1 Dáin Ironfoot | threat | Marked finisher that presents a substantial closing threat.
- 1 Desert Were-Worm | threat | Marked finisher that gives the deck a large threat.
- 1 Gollum, Riddle Master | threat | Marked finisher that also occupies a draw slot in the card pool.
- 1 Inferno Titan | threat | Large creature threat with removal utility.
- 1 Olog-hai Crusher | threat | Creature threat for pressuring opponents.
- 1 Oliphaunt | threat | Large creature threat for the combat plan.
- 1 Sauron, the Lidless Eye | threat | Legendary creature that adds pressure to the board.
- 1 Stone-Giant of High Pass | threat | Creature threat with removal utility.
- 1 The Great Goblin | threat | Legendary Goblin threat that reinforces the creature plan.
- 1 Troll of Khazad-dûm | threat | Marked finisher that provides a durable top-end threat.
- 1 Witch-king of Angmar | threat | Marked finisher with removal utility.
- 1 Witch-king, Bringer of Ruin | threat | Marked finisher with removal utility.
- 1 Smaug, Wicked Worm | wincon | Marked finisher that provides a dragon-themed route to victory.
- 1 Smaug, the Great Calamity // Spew Flame | wincon | Marked finisher that can close games while retaining removal utility.
- 1 Along the Crooked Way | other | Thematic noncreature inclusion from the requested card pool.
- 1 Andúril, Flame of the West | other | Equipment that broadens the deck's permanent suite.
- 1 Glamdring | other | Additional thematic equipment for the creature-heavy plan.
- 1 Giant's Boulder | ramp | the mana pass added it to bring the mana base inside the power level

</details>

### 22. a set family and a card from outside it

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 175 names.

Commander: Thranduil, the Elvenking.

Grade: bad, score 0.04, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $468.22 to buy, $468.22 the whole deck.

**Summary:** Thranduil, the Elvenking leads a creature-forward deck built around Elves, steady mana development, and a broad set of answers for opposing threats. Establish an active battlefield, keep resources flowing with draw pieces, and use targeted removal or wipes to clear the way for combat pressure. Troll of Khazad-dûm and the two Witch-kings are the primary closing threats, chief among them when the game reaches its later turns. The deck gives up some raw speed for a fuller battlefield plan and relies on maintaining creatures and mana development to convert its advantage into a win.

The quality model grades this deck below the precon baseline against the top lists of the format: the black sources cover 19.25 of the 26 its spells need.

- PLAN plan_coherent=partly: The core is a Sultai elf/creature board plan with removal support, but a heap of generic goodstuff removal and board wipes (Languish, Gnashing of Teeth, Raise the Palisade) cuts against the small-creature elf battlefield the summary says it wants to build, and the three named "wincons" are just unrelated big bodies rather than a payoff for the elf shell.
- PLAN theme_fit=partly: The commander and a good share of the list are Hobbit-set Mirkwood/Elvenking cards, and Sol Ring was kept as asked, but a sizable block of cards (The One Ring, Orcish Bowmasters, Delighted Halfling, Elvish Mystic, Lorien Revealed, Languish, Night's Whisper) comes from outside the requested set.
- PLAN useful_as_built=partly: The curve, ramp, and card draw are reasonable and it can play a normal game, but a three-color deck with 36 basics and no duals or fixing lands—black sources covering only about 19 of 26 needed—will regularly stumble on casting its commander and its black spells.
- PLAN summary_honest=partly: It fairly describes the elf/creature-plus-answers shape and admits the deck trades speed for board development, but it oversells three isolated fatties as a coherent closing plan and says nothing about the all-basic three-color mana base that its own quality note shows is short on black.
- [INFO] `curve_summary`: average mana value 3.02 over 63 nonland cards
- [WARN] `profile_off_band`: the worst color's share of the sources it needs is 0.74, and bracket 3 wants 0.85 or more (sources of requirement: U 14.5 of 19, B 19.25 of 26, G 15 of 19)
- [INFO] `mana_pass`: the builder moved 4 cards of the mana base to bring the deck inside its power level
- [INFO] `cards_trimmed`: the list was 1 card over, so the builder cut this card: Last March of the Ents
- [WARN] `outside_requested_set`: the sets you named do not hold 1 card: Sol Ring

<details><summary>The deck list</summary>

- 12 Forest | land | Provides a reliable green source for the deck’s Elf-heavy core.
- 10 Island | land | Provides a reliable blue source for the deck’s card flow and interaction.
- 14 Swamp | land | Provides a reliable black source for removal and finishing threats.
- 1 Sol Ring | ramp | Kept as requested to accelerate the deck’s early development.
- 1 Arcane Signet | ramp | Flexible mana acceleration for the multicolor plan.
- 1 Delighted Halfling | ramp | Early creature-based mana acceleration that fits the creature plan.
- 1 Elven Chorus | ramp | Elf-focused mana support for developing the board.
- 1 Elvish Archdruid | ramp | An Elf mana engine that rewards committing creatures to the board.
- 1 Elvish Mystic | ramp | Low-cost acceleration for getting established early.
- 1 Mox Amber | ramp | A compact mana piece supported by the deck’s legendary cards.
- 1 Necklace of Girion | ramp | Artifact-based mana acceleration for the midgame.
- 1 Wood Elves | ramp | Creature-based ramp that contributes to the board.
- 1 Woodland Weavemaster | ramp | Elf creature ramp that supports the deck’s development.
- 1 Elvish Visionary | draw | Early card flow attached to an Elf body.
- 1 Hithlain Knots | draw | Adds efficient card flow when resources run low.
- 1 Lórien Revealed | draw | Reliable card flow that helps smooth longer games.
- 1 Night's Whisper | draw | Low-cost card flow to keep land drops and spells coming.
- 1 Palantír of Orthanc | draw | A persistent source of card advantage.
- 1 Plunder the Trollshaws | draw | Instant-speed card flow for rebuilding or pressing an advantage.
- 1 Reverent Howl | draw | Adds another efficient way to keep cards moving.
- 1 The Master of Lake-town | draw | A creature-based card advantage piece for the board plan.
- 1 Thrór's Map | draw | Artifact card flow that helps sustain the deck.
- 1 Uncover the Moon-Letters | draw | Ongoing card flow for creature-heavy games.
- 1 Bilbo's Ring | interaction | Protective interaction for an important creature or permanent.
- 1 Confusticate and Bebother | interaction | A flexible interactive spell for disrupting opposing plans.
- 1 Elrond, Moon-Reader | interaction | Creature-based interaction that remains useful on the battlefield.
- 1 Mithril Coat | interaction | Protects a key creature from opposing answers.
- 1 My Precious // Allure of Power | interaction | Versatile interaction that supports the legendary-heavy shell.
- 1 Old Fat Spider Can't See Me | interaction | A thematic interactive piece that helps defend the board plan.
- 1 Stern Scolding | interaction | Cheap disruption for opposing early plays.
- 1 The One Ring | interaction | A high-impact legendary interactive tool for stabilizing key turns.
- 1 Bilbo's Deadly Slice | removal | Efficient spot removal for troublesome opposing permanents.
- 1 Bitter Downfall | removal | Direct removal for a problematic target.
- 1 Crude Bent Blade | removal | Equipment-based removal that can stay useful on board.
- 1 Enchanted River's Grasp | removal | Blue removal that answers a key opposing permanent.
- 1 Giant's Boulder | removal | Artifact removal that gives the deck another answer type.
- 1 Merciless Executioner | removal | Creature-based removal that contributes to the board.
- 1 Orcish Bowmasters | removal | Efficient removal attached to a creature.
- 1 The Black Arrow | removal | A legendary removal tool that fits the deck’s equipment package.
- 1 Uneasy Partings | removal | Instant-speed removal for answering threats at a key moment.
- 1 Gnashing of Teeth | wipe | A board reset for creature-heavy opposing positions.
- 1 Languish | wipe | A dependable reset when the battlefield gets out of hand.
- 1 Raise the Palisade | wipe | A tribal-aware wipe that can preserve the deck’s own board.
- 1 Elven Raft-Steerer | synergy | An Elf synergy piece that supports Thranduil’s creature-centered plan.
- 1 Elvenking's Harper | synergy | An Elf body that reinforces the deck’s Elven identity.
- 1 Galion, Elvenking's Butler | synergy | A legendary Elf that supports the deck’s Thranduil-focused theme.
- 1 Lothlórien Lookout | synergy | An Elf synergy creature for building a cohesive board.
- 1 Nimrodel Watcher | synergy | An Elf creature that strengthens the deck’s shared creature plan.
- 1 Thranduil, Sindarin Liege // Silvan Rally | synergy | A Thranduil-themed Elf card that reinforces the deck’s central theme.
- 1 Mirkwood Nurturer | synergy | An Elf synergy creature that helps make the board more cohesive.
- 1 Attercop | threat | A substantial creature threat that pressures the battlefield.
- 1 Boughside Wanderers | threat | An Elf creature threat that adds to the attacking board.
- 1 Cantankerous Keepers | threat | An Elf threat that helps establish battlefield presence.
- 1 Desolation Prowler | threat | A creature threat that adds pressure during combat.
- 1 Duskwatch Hunter | threat | A creature threat that helps turn board development into pressure.
- 1 Galadhrim Guide | threat | An Elf threat that supports a wide creature battlefield.
- 1 Haunt of the Dead Marshes | threat | A resilient-looking creature threat for the deck’s combat plan.
- 1 Head of the Hunt | threat | A creature threat that supports attacking with the board.
- 1 Mirkwood Elk | threat | A creature threat that contributes to the deck’s battlefield presence.
- 1 Mirkwood Meditator | threat | An Elf creature threat that fits the Mirkwood core.
- 1 Mirkwood Pathmaker | threat | An Elf threat that expands the deck’s board presence.
- 1 Supper for Spiders | synergy | A dedicated synergy card for the deck’s creature-centered strategy.
- 1 Troll of Khazad-dûm | wincon | A marked finisher that gives the deck a decisive closing threat.
- 1 Witch-king of Angmar | wincon | A marked finisher that can close games once the board is controlled.
- 1 Witch-king, Bringer of Ruin | wincon | A marked finisher that gives the deck another powerful endgame.
- 1 Pelargir Survivor | ramp | the mana pass added it to bring the mana base inside the power level

</details>

### 23. two set families at once

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 203 names.

Commander: Kíli the Resourceful.

Grade: typical, score 0.42, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: yes, for 1 missed names, deck_size. Block findings: 0.

Cost: $188.67 to buy, $188.67 the whole deck.

**Summary:** Kíli leads a white Dwarf-and-artifact strategy that develops mana, Equipment, Food, and creature tokens before leveraging tribal support and resilient value creatures. The deck sustains itself with artifact-based card advantage, protects important permanents, controls opposing threats efficiently, and closes through wide combat forces or large evasive finishers.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

This deck plan aims for 2 or more finishers, the cards that win the game, and the deck holds 1.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core of white token-making, artifacts and Equipment with Kíli hangs together loosely, but the claimed Dwarf tribal payoff rests on a handful of Dwarves and cards like Patchwork Banner/Maskwood Nexus, and three sweepers (Dusk, Martial Coup, Promise of Loyalty) pull against the go-wide token plan the deck otherwise builds.
- PLAN theme_fit=partly: The commander is chosen as asked and a meaningful slice of cards come from the Hobbit and Bloomburrow sets, but well over half the list (Sol Ring, Skullclamp, Swords to Plowshares, Academy Manufactor, Blade Splicer, Karn, Sun Titan, Jazal Goldmane, Psychosis Crawler, and more) falls outside the two named sets, so the named card-pool limit is not honored.
- PLAN useful_as_built=yes: Mono-white with ~37 lands plus eight ramp pieces, a smooth curve, ample draw and removal, and several credible closers (Jazal, Helm of the Host, Psychosis Crawler, Angel of the Ruins) means it can be picked up and played as it stands.
- PLAN summary_honest=partly: Artifact card advantage, protection, removal and wide-combat closing are all really present, but "Dwarf tribal support" and the Food angle are overstated relative to the two or three cards that actually support them.
- [INFO] `curve_summary`: average mana value 3.08 over 63 nonland cards
- [WARN] `finisher_short`: the finisher count is 1, and this deck plan wants 2 or more (Angel of the Ruins)

<details><summary>The deck list</summary>

- 25 Plains | land | Reliable white mana base.
- 1 Command Tower | land | Untapped white source.
- 1 Castle Ardenvale | land | White source with late-game token utility.
- 1 Minas Tirith | land | White source that supports the creature plan.
- 1 Path of Ancestry | land | Tribal fixing for the Dwarf commander.
- 1 Thriving Heath | land | Flexible white fixing.
- 1 Uncharted Haven | land | Additional color fixing.
- 1 Lupinflower Village | land | White-producing utility land.
- 1 Hobbit Hole | land | White source with Food utility.
- 1 Evolving Wilds | land | Finds a Plains while fixing mana.
- 1 Fabled Passage | land | Fetches Plains for reliable early mana.
- 1 Terramorphic Expanse | land | Additional Plains fetch land.
- 1 Arcane Signet | ramp | Efficient commander-color acceleration.
- 1 Bag End Banquet | ramp | Theme-fitting mana acceleration.
- 1 Fellwar Stone | ramp | Low-cost mana fixing.
- 1 Mind Stone | ramp | Early mana that can later become a card.
- 1 Ornithopter of Paradise | ramp | Creature-based colored acceleration.
- 1 Patchwork Banner | ramp | Mana rock that also supports the creature type.
- 1 Sol Ring | ramp | Highly efficient colorless acceleration.
- 1 Thought Vessel | ramp | Reliable mana rock with hand-size utility.
- 1 Wayfarer's Bauble | ramp | Turns early mana into a Plains.
- 1 Loyal Warhound | ramp | Creature ramp that helps catch up on lands.
- 1 Caretaker's Talent | draw | Repeatable card advantage from token-making.
- 1 Circuit Mender | draw | Artifact body that replaces itself.
- 1 Cut a Deal | draw | Efficient burst of cards.
- 1 Dawn of a New Age | draw | Ongoing card advantage for the board plan.
- 1 Feather of Flight | draw | Cantrip that improves a creature in combat.
- 1 Fountainport Bell | draw | Artifact-based repeatable card selection.
- 1 Idol of Oblivion | draw | Efficient draw alongside tokens.
- 1 Inspiring Overseer | draw | Creature that immediately replaces itself.
- 1 Mentor of the Meek | draw | Rewards the deck's smaller creatures and tokens.
- 1 Skullclamp | draw | Strong card flow from expendable creatures.
- 1 Spirited Companion | draw | Low-cost creature that draws on entry.
- 1 Baird, Steward of Argive | interaction | Discourages opponents from attacking freely.
- 1 Dawn's Truce | interaction | Protects the board during critical turns.
- 1 Galadriel's Dismissal | interaction | Flexible defensive tempo spell.
- 1 Mithril Coat | interaction | Protects Kíli or a key creature.
- 1 Parting Gust | interaction | Versatile instant-speed disruption.
- 1 Reprieve | interaction | Temporarily answers a spell while replacing itself.
- 1 Selfless Spirit | interaction | Protects the creature board from destruction.
- 1 Swiftfoot Boots | interaction | Provides haste and protection for key creatures.
- 1 Banishing Light | removal | Broad answer to troublesome permanents.
- 1 Fiend Hunter | removal | Creature-based exile removal.
- 1 Generous Gift | removal | Answers any permanent at instant speed.
- 1 Hanged Executioner | removal | Token-producing creature with a removal mode.
- 1 Loran of the Third Path | removal | Removes artifacts or enchantments while adding utility.
- 1 Skyclave Apparition | removal | Efficient exile effect on a creature.
- 1 Sonar Strike | removal | Useful creature or planeswalker answer.
- 1 Swords to Plowshares | removal | Premium efficient creature removal.
- 1 Westfold Rider | removal | Creature-based permanent interaction.
- 1 Academy Manufactor | synergy | Multiplies the value of Food and artifact-token production.
- 1 Blade Splicer | synergy | Provides an artifact creature and supports go-wide combat.
- 1 Carrot Cake | synergy | Creates Food and a body for the artifact-token plan.
- 1 Dáin, Lord of the Iron Hills | synergy | Supports Kíli's Dwarf-focused combat plan.
- 1 Iron Hills Blacksmith | synergy | Improves the Equipment-focused Dwarf strategy.
- 1 Maskwood Nexus | synergy | Unifies creature types for tribal payoffs.
- 1 Tangle Tumbler | synergy | Artifact payoff that complements the deck's token plan.
- 1 Fíli the Pathfinder | threat | Legendary Dwarf threat that advances the theme.
- 1 Karn, the Great Creator | threat | High-impact noncreature threat with artifact utility.
- 1 Jazal Goldmane | threat | Turns a developed creature board into major combat pressure.
- 1 Landroval, Horizon Witness | threat | Evasive legendary creature for sustained pressure.
- 1 Serra Redeemer | threat | Rewards creature development with larger threats.
- 1 Sun Titan | threat | Large attacker that continually recovers key permanents.
- 1 Warren Warleader | threat | Creates a wide combat presence.
- 1 Eagles of the North | threat | Evasive creature that can pressure opponents from the air.
- 1 Hoofprints of the Stag | threat | Converts steady card flow into recurring flying threats.
- 1 Helm of the Host | threat | Creates compounding copies of the deck's best creatures.
- 1 Psychosis Crawler | threat | Transforms card draw into table-wide life pressure.
- 1 The Eagles Are Coming! | threat | Builds an evasive token force.
- 1 Angel of the Ruins | wincon | Large evasive artifact threat that also removes opposing artifacts or enchantments.
- 1 Realm-Cloaked Giant // Cast Off | wincon | Large creature that can clear the way before attacking.
- 1 Sunscorch Regent | wincon | Evasive dragon that grows into a lethal threat.
- 1 Dusk // Dawn | wipe | Flexible creature reset that can later recover small creatures.
- 1 Martial Coup | wipe | Clears creatures and leaves behind an army.
- 1 Promise of Loyalty | wipe | Resets opposing boards while preserving a key creature.

</details>

### 24. a 60-card deck from one set

Format: Modern. Theme: aggro. Pool: any_card. Shortlist: 96 names.

Grade: baseline, score 0.36, model 20260914T154223Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $20.36 to buy, $20.36 the whole deck.

**Summary:** This mono-red Bloomburrow aggro deck presses early with a dense mix of Mouse and Lizard creatures, using Emberheart Challenger and Hearthborn Battler as chief among its synergy pieces. Its removal clears the way for repeated attacks, while its draw cards help it continue deploying threats after the first wave. It wins by maintaining creature pressure, giving up some flexibility and late-game power for a focused, fast board presence.

The quality model grades this deck at the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. The color sources cover the pips, and that raises the grade.

- PLAN plan_coherent=yes: Every slot supports one low-curve red aggro plan: twelve-plus cheap Lizard/Mouse bodies, eight cheap removal spells to keep attacks going, and a handful of cantrip-style tricks to refuel, with nothing pulling toward control or late-game value.
- PLAN theme_fit=yes: It is mono-red, Modern-legal, built entirely from Bloomburrow cards, and aimed at casual-level aggro exactly as asked, with no off-set staples smuggled in.
- PLAN useful_as_built=yes: Twenty-four Mountains for a curve that stops at four is more land than an aggro deck wants, but the mana is flawless, the threat count is high, and the deck can be shuffled up and played as written without any dead draws from color problems.
- PLAN summary_honest=partly: The removal, threat density, and pressure-based win condition are all accurately described, but calling it a 'dense mix of Mouse and Lizard creatures' oversells the Mouse half (Emberheart Challenger is essentially the only Mouse payoff) and it omits that a 24-land mana base is heavier than the stated fast curve needs.
- [INFO] `curve_summary`: average mana value 2.94 over 36 nonland cards

<details><summary>The deck list</summary>

- 24 Mountain | land | Provides a consistent mono-red mana base for an aggressive curve.
- 2 Brazen Collector | ramp | Adds early mana support while remaining part of the Bloomburrow creature suite.
- 4 Might of the Meek | draw | Supplies card flow for the low-curve creature plan.
- 2 Sazacap's Brew | draw | Adds further card flow without leaving the mono-red plan.
- 4 Abrade | removal | Provides efficient answers to opposing problems.
- 2 Agate Assault | removal | Rounds out the deck's direct removal package.
- 4 Emberheart Challenger | synergy | Forms a core Mouse synergy package for early pressure.
- 4 Hearthborn Battler | synergy | Supports the aggressive creature synergies alongside the Mouse package.
- 4 Frilled Sparkshooter | threat | A full set of Lizard threats to keep pressure on the opponent.
- 4 Reptilian Recruiter | threat | Provides another full set of creature threats for the early and middle turns.
- 4 Teapot Slinger | threat | Adds more Bloomburrow creature pressure to the deck's threat base.
- 2 Stormsplitter | threat | Completes the threat suite with additional aggressive creatures.

</details>

### 25. use no card of an owned precon

Format: Commander. Theme: Avengers superheroes. Pool: owned_first. Shortlist: 154 names.

Commander: Syr Konrad, the Grim.

Grade: baseline, score 0.32, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $199.11 the whole deck.

**Summary:** Syr Konrad, the Grim leads a black attrition deck built around creatures, removal, -1/-1 counter pressure, and a compact superhero-flavored artifact package. It develops with efficient mana artifacts, keeps cards flowing through equipment and black draw spells, then turns a crowded or depleted battlefield into an advantage with sweepers and recurring pressure. Archfiend of Ifnir, Vraska, Betrayal's Sting, and Witch-king of Angmar are the strongest closing cards, while the deck gives up multicolor Avengers characters in exchange for a consistent Swamp-heavy mana base and a focused black game plan.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core of -1/-1 counter proliferation, sweepers and cheap removal does feed Syr Konrad's death triggers, but the equipment/aura-of-protection package (Buster Sword, Mask of Memory, Champion's Helm, Sword of the Animist, Springleaf Drum) wants a wide creature board that the deck's own Black Sun's Zenith, Extinction Event, Midnight Banshee and Massacre Girl actively tear down.
- PLAN theme_fit=no: The person asked for an Avengers superheroes deck and received a mono-black Syr Konrad grindfest whose flavor is Mordor, Final Fantasy and Avatar, with two Marvel cards in 99 and a summary that openly concedes the theme was traded away.
- PLAN useful_as_built=yes: Mono-black with 36 Swamps and ten mana rocks has no mana problems, the curve is low, and abundant removal plus Syr Konrad drain, Archfiend, Vraska and Witch-king give real ways to close a game.
- PLAN summary_honest=partly: The black attrition/-1/-1 counter description matches the list and it honestly admits abandoning multicolor Avengers characters, but the claimed "compact superhero-flavored artifact package" is a fiction — M.O.D.O.K. and Stolen Stark Tech aside, the artifacts are LOTR, Final Fantasy and Avatar cards with no Marvel flavor at all.
- [WARN] `not_owned`: Syr Konrad, the Grim: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.22 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Arcane Signet | ramp | Reliable early mana acceleration.
- 1 Bender's Waterskin | ramp | Adds another mana-producing artifact.
- 1 Chromatic Lantern | ramp | Provides mana acceleration for the deck's higher costs.
- 1 Commander's Sphere | ramp | Flexible artifact mana for the midgame.
- 1 Ring of the Lucii | ramp | Supports the mana development plan.
- 1 Sol Ring | ramp | Efficient early artifact acceleration.
- 1 Springleaf Drum | ramp | Low-cost mana acceleration alongside the creature suite.
- 1 Sword of the Animist | ramp | Equipment-based mana development.
- 1 Thought Vessel | ramp | Artifact mana that remains useful later.
- 1 Wayfarer's Bauble | ramp | Early mana development that supports consistent land access.
- 1 Buster Sword | draw | Equipment-based card advantage.
- 1 Call of the Ring | draw | Ongoing card advantage for a long game.
- 1 Eventide's Shadow | draw | Adds a dedicated card-advantage spell.
- 1 Hoarder's Greed | draw | Helps refill the hand.
- 1 Idol of Oblivion | draw | Low-cost artifact card advantage.
- 1 Lembas | draw | Artifact card advantage that fits the deck's utility package.
- 1 Mask of Memory | draw | Equipment-based card selection and advantage.
- 1 Mirror of Galadriel | draw | Legendary artifact card advantage.
- 1 Mordor Muster | draw | Provides another dedicated draw spell.
- 1 Nasty End | draw | Instant-speed card advantage.
- 1 Night's Whisper | draw | Efficient early card advantage.
- 1 Champion's Helm | interaction | Protects an important legendary creature.
- 1 Darksteel Plate | interaction | Durable protection for a key creature.
- 1 Lightning Greaves | interaction | Efficient protection for Syr Konrad or a finisher.
- 1 Orcish Medicine | interaction | Flexible black interaction.
- 1 Stolen Stark Tech | interaction | Utility equipment that supports the deck's key permanents.
- 1 Stone of Erech | interaction | Artifact utility that disrupts opposing graveyard plans.
- 1 Swiftfoot Boots | interaction | Protection and mobility for key creatures.
- 1 The Walls of Ba Sing Se | interaction | Defensive artifact interaction for the board.
- 1 Bitter Downfall | removal | Direct removal for troublesome permanents.
- 1 Bitter Triumph | removal | Efficient answer to opposing threats.
- 1 Fatal Push | removal | Low-cost creature removal.
- 1 Heartless Act | removal | Flexible instant-speed creature removal.
- 1 Infernal Grasp | removal | Straightforward answer to a major threat.
- 1 Nameless Inversion | removal | Creature removal that works well with the deck's graveyard focus.
- 1 Orcish Bowmasters | removal | Creature-based removal and board pressure.
- 1 Skinrender | removal | Creature-based removal that leaves behind a body.
- 1 Zuko's Exile | removal | Flexible instant-speed removal.
- 1 Black Sun's Zenith | wipe | Scalable board reset for creature-heavy tables.
- 1 Extinction Event | wipe | Selective board wipe that can preserve part of the board.
- 1 One Ring to Rule Them All | wipe | Saga-based reset that fits the deck's darker theme.
- 1 Archfiend of Ifnir | wincon | A marked finisher that turns the deck's resource engine into a closing threat.
- 1 Vraska, Betrayal's Sting | wincon | A marked finisher that gives the deck a separate path to close games.
- 1 Witch-king of Angmar | wincon | A marked finisher and imposing late-game threat.
- 1 Blowfly Infestation | synergy | Supports the deck's counter-based attrition package.
- 1 Contagion Clasp | synergy | Provides a compact piece for the counter package.
- 1 Incremental Blight | synergy | Builds on the deck's -1/-1 counter plan.
- 1 Massacre Girl, Known Killer | synergy | Supports the creature-destruction and counter-focused game plan.
- 1 Midnight Banshee | synergy | Adds recurring pressure to the counter package.
- 1 Soul Snuffers | synergy | Advances the -1/-1 counter plan across the table.
- 1 The Black Breath | synergy | Fits the deck's sweeping, attrition-oriented plan.
- 1 Carnifex Demon | threat | Large creature that pressures the table while contributing to the counter theme.
- 1 Dusk Urchins | threat | Creature threat that also rewards the deck's attrition plan.
- 1 Gaius van Baelsar | threat | Legendary creature that adds board presence to the deck.
- 1 Grim Poppet | threat | Artifact creature that applies pressure within the counter package.
- 1 PuPu UFO | threat | Artifact creature that adds a distinct body to the board.
- 1 Rat King, Pale Piper | threat | Legendary creature that expands the deck's creature pressure.
- 1 Ringwraiths | threat | Creature threat suited to the deck's dark, creature-heavy plan.
- 1 Sinister Gnarlbark | threat | Creature that provides an additional board threat.
- 1 Undercity Dire Rat | threat | Creature-based pressure that supports early development.
- 1 Voracious Fell Beast | threat | Creature threat with removal-oriented utility.
- 1 Zodiark, Umbral God | threat | Large legendary threat for the late game.
- 1 M.O.D.O.K. | threat | A superhero-themed artifact creature that adds board presence.
- 36 Swamp | land | Reliable untapped black mana for Syr Konrad and the deck's black-intensive spells.

</details>

