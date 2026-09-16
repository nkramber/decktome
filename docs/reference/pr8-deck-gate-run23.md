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
| Decks that needed the repair turn | 1 |
| Summaries judged (F-26) | 25 |
| Summaries that state a rule of the game | 0 |
| Summaries that state a FALSE rule | 0 |
| Judge errors | 0 |
| Case assertions missed (PR-28b) | 0 |
| Decks the plan judge read (PR-15, information) | 25 |
| Mean plan score, 0 to 1 | 0.72 |
| Plan reasons the judge left empty | 4 |
| Errors | 0 |
| Prompt version | 13 |
| Calls | 76 |
| Cost | $2.6509 |
| Time | 2156 seconds |

## Run

- Suite `decks`, run `pr8-deck-gate-run23`, on 2026-09-16, commit `76e53c7`.
- Roles: generate on `gpt-5.6-terra` (openai, effort medium), judge on `claude-opus-5` (anthropic, effort medium), repair on `gpt-5.6-terra` (openai, effort medium).
- Versions: card snapshot 2026-09-04, generate prompt version 13, plan_rubric prompt version 2, precons `5.3.0+20260903`, quality_model `20260914T154223Z`.
- Calls: 76. Cost: $2.6509. Time: 2156 seconds.

## Findings by code

A block stops the deck. A warning and a note are reports.

| Code | Count |
|---|---|
| `curve_summary` | 25 |
| `not_owned` | 14 |
| `profile_off_band` | 12 |
| `mana_pass` | 9 |
| `basics_added` | 4 |
| `outside_requested_set` | 2 |
| `cards_trimmed` | 1 |
| `land_count` | 1 |
| `bracket_cut` | 1 |

By severity: BLOCK 0. WARN 29. INFO 40. 

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

Format: Commander. Theme: lifegain. Pool: any_card. Shortlist: 303 names.

Commander: Karlov of the Ghost Council.

Grade: bad, score 0.03, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $750.13 to buy, $750.13 the whole deck.

**Summary:** Karlov leads a lifegain-focused deck that develops a creature board through its synergy pieces, keeps resources flowing with draw and ramp, and uses focused interaction alongside board wipes to stay in the game. It wins by turning that developing board and its threat package into sustained pressure, giving up some raw speed for a broader mix of protection, removal, and longer-game tools.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core is a consistent lifegain-payoff shell (Karlov, Ajani's Pridemate, Heliod, Righteous Valkyrie, Vito, drain finishers) supported by lifelink and life-triggered draw, but the three symmetrical wraths and a scattering of low-impact 'ramp' artifacts (Colossal Plow, Hot Dog Cart, Nuka-Cola Vending Machine, Azor's Gateway) pull against a board-centric creature plan.
- PLAN theme_fit=yes: It is exactly a W/B Karlov lifegain Commander deck at a mid power level, with counter-growth, life-triggered payoffs, and drain finishers, and nothing pushing it above a bracket-3 ceiling.
- PLAN useful_as_built=partly: The 37-land mana base with duals and both basics is solid and there are several real win routes (Vito, Defiant Bloodlord, Cliffhaven Vampire, Karlov beats), but the acceleration slots are mostly junk artifacts that don't actually ramp, so a top-heavy curve of five- and six-drops will often land late.
- PLAN summary_honest=yes: placeholder
- [INFO] `curve_summary`: average mana value 3.21 over 62 nonland cards
- [INFO] `basics_added`: the list was 1 card short, so the builder added 1 basic land

<details><summary>The deck list</summary>

- 1 Command Tower | land | Land slot for the two-color mana base.
- 1 Godless Shrine | land | Land slot for the two-color mana base.
- 1 Scrubland | land | Land slot for the two-color mana base.
- 1 Caves of Koilos | land | Land slot for the two-color mana base.
- 1 Concealed Courtyard | land | Land slot for the two-color mana base.
- 1 Isolated Chapel | land | Land slot for the two-color mana base.
- 1 Shattered Sanctum | land | Land slot for the two-color mana base.
- 1 Silent Clearing | land | Land slot for the two-color mana base.
- 1 Vault of Champions | land | Land slot for the two-color mana base.
- 1 Brightclimb Pathway // Grimclimb Pathway | land | Flexible land slot for the two-color mana base.
- 1 Marsh Flats | land | Land slot for the two-color mana base.
- 14 Plains | land | Basic land allocation emphasizing white mana.
- 12 Swamp | land | Basic land allocation emphasizing black mana.
- 1 Arguel's Blood Fast // Temple of Aclazotz | ramp | Ramp slot that also fits the deck's lifegain theme.
- 1 Azor's Gateway // Sanctum of the Sun | ramp | Ramp slot for developing into the midgame.
- 1 Colossal Plow | ramp | Early ramp slot supporting the mana plan.
- 1 Hot Dog Cart | ramp | Ramp slot supporting the mana plan.
- 1 Legion's Landing // Adanto, the First Fort | ramp | Low-cost ramp slot that supports board development.
- 1 Life Insurance | ramp | Ramp slot for sustaining the deck's development.
- 1 Nuka-Cola Vending Machine | ramp | Ramp slot for the deck's mana development.
- 1 Phial of Galadriel | ramp | Ramp slot for the deck's mana development.
- 1 Pristine Talisman | ramp | Ramp slot that fits the lifegain-focused plan.
- 1 The Celestus | ramp | Ramp slot for consistent mana development.
- 1 Archivist of Oghma | draw | Draw slot that helps keep resources flowing.
- 1 Dawn of Hope | draw | Draw slot aligned with the lifegain plan.
- 1 Cosmos Elixir | draw | Draw slot for maintaining resources through the game.
- 1 Enduring Innocence | draw | Draw slot supporting continued board development.
- 1 Mangara, the Diplomat | draw | Draw slot for maintaining resources.
- 1 Markov Purifier | draw | Draw slot that fits the deck's creature theme.
- 1 Sigarda's Splendor | draw | Draw slot aligned with the lifegain plan.
- 1 The Gaffer | draw | Draw slot for sustained resources.
- 1 Vampiric Rites | draw | Efficient draw slot for the creature-heavy build.
- 1 Well of Lost Dreams | draw | Draw slot aligned with the lifegain plan.
- 1 Ayli, Eternal Pilgrim | removal | Dedicated removal option on a thematic creature.
- 1 Gumdrop Poisoner // Tempt with Treats | removal | Flexible removal slot for opposing permanents.
- 1 Henrika Domnathi // Henrika, Infernal Seer | removal | Removal slot that also adds a creature body.
- 1 Murderous Rider // Swift End | removal | Flexible removal option for problematic threats.
- 1 Nightmare's Thirst | removal | Low-cost removal slot.
- 1 Solitude | removal | Removal option for answering opposing threats.
- 1 Umezawa's Jitte | removal | Repeatable removal slot for creature combat.
- 1 Vona, Butcher of Magan | removal | Removal slot on a lifegain-themed creature.
- 1 Witch of the Moors | removal | Removal slot that fits the lifegain plan.
- 1 Alseid of Life's Bounty | interaction | Low-cost interaction to protect key pieces.
- 1 Bofur, Reliable Guardian // Concerted Care | interaction | Interaction slot for protecting the board.
- 1 Courageous Resolve | interaction | Interaction slot for protecting important threats.
- 1 Faith's Shield | interaction | Efficient interaction for protecting key permanents.
- 1 Metropolis Reformer | interaction | Interaction creature that supports the defensive plan.
- 1 Restoration Magic | interaction | Interaction slot for recovering or protecting resources.
- 1 Sword of Light and Shadow | interaction | Interaction equipment for supporting creature combat.
- 1 Werefox Bodyguard | interaction | Interaction creature for disrupting opponents.
- 1 Ajani's Pridemate | synergy | Core lifegain synergy piece for Karlov's plan.
- 1 Angel of Vitality | synergy | Lifegain synergy creature supporting the main plan.
- 1 Bloodthirsty Aerialist | synergy | Lifegain synergy creature that grows alongside Karlov.
- 1 Cleric Class | synergy | Lifegain synergy piece for the deck's central theme.
- 1 Heliod, Sun-Crowned | synergy | Central lifegain synergy piece.
- 1 Resplendent Angel | synergy | Lifegain synergy threat that supports board development.
- 1 Righteous Valkyrie | synergy | Lifegain synergy creature for the deck's board plan.
- 1 Serra Ascendant | synergy | Low-cost lifegain synergy creature.
- 1 Vito, Thorn of the Dusk Rose | synergy | Lifegain synergy piece that supports closing games.
- 1 Voice of the Blessed | synergy | Lifegain synergy creature that develops with the plan.
- 1 Archangel of Thune | threat | High-impact threat for the lifegain strategy.
- 1 Attended Healer | threat | Threat that supports the deck's creature plan.
- 1 Astarion, the Decadent | threat | Lifegain-themed threat for finishing pressure.
- 1 Celestine, the Living Saint | threat | Threat that supports the deck's creature plan.
- 1 Cliffhaven Vampire | threat | Lifegain-themed threat for applying pressure.
- 1 Defiant Bloodlord | threat | Lifegain-themed threat for closing games.
- 1 Divinity of Pride | threat | Large lifegain-themed threat.
- 1 Enduring Tenacity | threat | Threat that fits the deck's lifegain focus.
- 1 Exalted Sunborn | threat | Threat supporting the lifegain creature strategy.
- 1 Nykthos Paragon | threat | Lifegain-themed threat for the board plan.
- 1 Rhox Faithmender | threat | Threat aligned with the lifegain strategy.
- 1 Valkyrie Harbinger | threat | Lifegain-themed threat for sustained pressure.
- 1 Fumigate | wipe | Board wipe for resetting crowded boards.
- 1 Kaya's Wrath | wipe | Board wipe for answering developed creature boards.
- 1 The Meathook Massacre | wipe | Board wipe that fits the deck's black-based plan.

</details>

### 2. aristocrats Commander, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 179 names.

Commander: Denethor, Ruling Steward.

Grade: baseline, score 0.36, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $137.30 the whole deck.

**Summary:** Denethor leads a white-black aristocrats deck that develops a creature board, turns sacrifice-focused synergies into lasting advantage, and keeps cards flowing through a deep draw package. It wins by maintaining pressure with its creatures and threats while using efficient removal and a few resets to clear resistance. The tradeoff is that the deck leans hard on establishing creatures and support pieces, so it is less explosive than a faster combo-focused build.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck draws more than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The removal, ramp, and draw packages are individually fine, but three board wipes plus protective equipment and graveyard hate (Stone of Erech) pull against the stated plan of developing a creature board and recurring sacrifice value, leaving a goodstuff pile rather than one plan.
- PLAN theme_fit=partly: It is a W/B Denethor Commander deck at a reasonable bracket-3 power level, but the aristocrats core — repeatable sacrifice outlets, token fodder, and death-trigger drains like Blood Artist effects — is nearly absent, with only Deadly Dispute, Nasty End, Skullclamp, and a couple of others gesturing at the theme.
- PLAN useful_as_built=yes: 37 lands with fixing, ten-plus ramp pieces, nine removal spells, a low curve, and about twenty creatures make this a functional, castable deck that can present a board and close games, even if not on the intended axis.
- PLAN summary_honest=partly: The draw package, removal, and resets are genuinely there, but the claim that the deck 'turns sacrifice-focused synergies into lasting advantage' overstates a list with only a handful of sac outlets and essentially no drain payoffs.
- [INFO] `curve_summary`: average mana value 2.79 over 61 nonland cards
- [WARN] `profile_off_band`: the draw count is 17, and bracket 3 wants 8 to 14

<details><summary>The deck list</summary>

- 15 Plains | land | Reliable white source for the deck’s core spells.
- 14 Swamp | land | Reliable black source for the deck’s core spells.
- 1 Command Tower | land | Flexible color fixing for the two-color mana base.
- 1 City of Brass | land | Provides access to either deck color.
- 1 Exotic Orchard | land | Flexible color fixing in multiplayer games.
- 1 Marsh Flats | land | Fetches a needed basic color source.
- 1 Fabled Passage | land | Finds a basic land to smooth early mana.
- 1 Evolving Wilds | land | Finds a basic land to fix colors.
- 1 Terramorphic Expanse | land | Finds a basic land to fix colors.
- 1 Path of Ancestry | land | A dual-color source that supports the creature-focused plan.
- 1 Secluded Courtyard | land | Creature-focused color fixing.
- 1 Call of the Ring | draw | Draw support that keeps resources flowing.
- 1 Cirith Ungol Patrol | draw | Creature-based draw support for the board-focused plan.
- 1 Exemplar of Light | draw | Adds a body while contributing card advantage.
- 1 Grave Venerations | draw | Draw support for a graveyard-oriented game plan.
- 1 Idol of Oblivion | draw | Low-cost card advantage support.
- 1 Inspiring Overseer | draw | Creature-based card advantage.
- 1 June, Bounty Hunter | draw | Adds a creature and draw support.
- 1 Kingpin's Enforcers | draw | Creature-based draw support.
- 1 Lembas | draw | Affordable artifact-based card advantage.
- 1 Mask of Memory | draw | Equipment-based card advantage for the creature suite.
- 1 Massacre Girl, Known Killer | draw | A creature that contributes card advantage.
- 1 Nasty End | draw | Efficient black draw support.
- 1 Night's Whisper | draw | Straightforward early card advantage.
- 1 Painful Truths | draw | Efficient card advantage.
- 1 Skullclamp | draw | Key card advantage for a creature-heavy sacrifice plan.
- 1 Stone of Erech | draw | Artifact-based draw support.
- 1 Tome of Legends | draw | Sustained card advantage alongside the commander.
- 1 Boromir, Warden of the Tower | interaction | Creature-based interaction that helps protect the plan.
- 1 Champion's Helm | interaction | Protects an important legendary creature.
- 1 Clever Concealment | interaction | Broad protection against opposing disruption.
- 1 Darksteel Plate | interaction | Durable protection for a key creature.
- 1 Gift of Immortality | interaction | Protection that supports a sacrifice-oriented board.
- 1 Lightning Greaves | interaction | Efficient protection for key creatures.
- 1 Reprieve | interaction | Flexible instant-speed disruption.
- 1 Swiftfoot Boots | interaction | Repeatable protection for important creatures.
- 1 Arcane Signet | ramp | Efficient two-color mana acceleration.
- 1 Bender's Waterskin | ramp | Artifact ramp for reaching the commander and higher-cost spells.
- 1 Chromatic Lantern | ramp | Mana acceleration and broad color fixing.
- 1 Commander's Sphere | ramp | Reliable fixing and mana acceleration.
- 1 Deadly Dispute | ramp | Turns a disposable resource into mana support.
- 1 Fellwar Stone | ramp | Low-cost mana acceleration and fixing.
- 1 Lotho, Corrupt Shirriff | ramp | Creature-based mana acceleration.
- 1 Sol Ring | ramp | Efficient early mana acceleration.
- 1 Springleaf Drum | ramp | Low-cost mana support for the creature-heavy build.
- 1 Wayfarer's Bauble | ramp | Early land-based mana development.
- 1 Bitter Triumph | removal | Flexible black spot removal.
- 1 Claim the Precious | removal | Focused creature removal.
- 1 Crib Swap | removal | Flexible creature removal.
- 1 Dismember | removal | Low-cost creature removal.
- 1 Fatal Push | removal | Efficient early spot removal.
- 1 Fiend Hunter | removal | Creature-based removal that fits the board plan.
- 1 Generous Gift | removal | Versatile permanent removal.
- 1 Get Lost | removal | Efficient flexible removal.
- 1 Swords to Plowshares | removal | Efficient creature removal.
- 1 Al Bhed Salvagers | synergy | Supports the deck’s central sacrifice-focused synergies.
- 1 Arcade Cabinet | synergy | Artifact synergy piece for the deck’s broader engine.
- 1 Gollum the Abandoned | synergy | A thematic synergy creature for the aristocrats shell.
- 1 Gollum, Patient Plotter | synergy | Supports the sacrifice-focused creature package.
- 1 Gríma Wormtongue | synergy | Adds another synergy-focused creature to the plan.
- 1 Heirloom Auntie | synergy | Supports the deck’s creature-based synergies.
- 1 Joo Dee, One of Many | synergy | A synergy creature that helps develop the board plan.
- 1 Nimble Hobbit | synergy | Low-cost creature synergy for the aristocrats plan.
- 1 Phantom Train | synergy | Artifact synergy that complements the creature strategy.
- 1 Rhovanion Rampager | synergy | Creature synergy that helps pressure the table.
- 1 Bill the Pony | threat | A creature threat that contributes to the board presence.
- 1 Hei Bai, Spirit of Balance | threat | A substantial threat for closing games.
- 1 Namazu Trader | threat | A creature threat that advances the proactive plan.
- 1 Vengeful Villagers | threat | A creature threat for converting board development into pressure.
- 1 Austere Command | wipe | Flexible board reset when the table gets ahead.
- 1 Dusk // Dawn | wipe | A board wipe that suits the creature-oriented shell.
- 1 Fumigate | wipe | Reliable full-board reset.

</details>

### 3. artifacts Commander, bracket 4

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 317 names.

Commander: Urza, Lord High Artificer.

Grade: good, score 0.70, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $7771.83 to buy, $7771.83 the whole deck.

**Summary:** This artifact-focused Urza deck develops through a dense package of artifact mana, draw, tutors, and protective interaction, then turns that infrastructure into pressure with artifact creatures and planeswalkers. It can close through Thassa's Oracle, Mechanized Production, Mirrodin Besieged, or sustained artifact threats while keeping removal and reset tools available. It gives up broader nonartifact flexibility in favor of a concentrated artifact plan and a demanding utility-land package.

The quality model grades this deck good against the top lists of the format: the deck makes more mana on turn four than the norm, and that raises the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

- PLAN plan_coherent=partly: The bulk of the list pulls one way — cheap artifact mana, affinity-style draw, artifact tutors, and artifact threats that Urza can convert into mana and Constructs — but Thassa's Oracle sits in a deck with no self-mill, Demonic Consultation, or Tainted Pact, so it is a stranded piece of a combo the deck never assembles.
- PLAN theme_fit=yes: It is exactly what was asked: a mono-blue Urza, Lord High Artificer artifact deck with fast mana, free counterspells, and strong tutors at a plausible bracket-4 power level.
- PLAN useful_as_built=yes: 36 lands-plus-fast-mana in a mono-colored deck with a very low artifact curve, ample draw, tutors, interaction, and multiple real finishers (Mechanized Production, Mirrodin Besieged, Aetherflux Reservoir, Kappa Cannoneer/Construct beats) is entirely playable as it stands.
- PLAN summary_honest=partly: The artifact engine, removal, planeswalkers, and the Mechanized Production/Mirrodin Besieged win routes are all genuinely present, but naming Thassa's Oracle as a way to close asserts a line the list cannot execute since nothing empties the library.
- [INFO] `curve_summary`: average mana value 2.94 over 65 nonland cards
- [INFO] `mana_pass`: the builder moved 2 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 26 Island | land | Basic blue mana source for the artifact-heavy shell.
- 1 Seat of the Synod | land | Artifact land that supplies a blue source.
- 1 Mystic Sanctuary | land | Blue-producing land slot.
- 1 Otawara, Soaring City | land | Blue-producing land slot.
- 1 Sink into Stupor // Soporific Springs | land | Blue-producing land slot.
- 1 Academy Ruins | land | Nonbasic utility land slot.
- 1 Inventors' Fair | land | Nonbasic utility land slot.
- 1 Mishra's Workshop | land | Game Changer land that supports the artifact plan.
- 1 Urza's Workshop | land | Nonbasic artifact-focused land slot.
- 1 Chrome Mox | ramp | Game Changer fast-mana artifact.
- 1 Lion's Eye Diamond | ramp | Game Changer fast-mana artifact.
- 1 Lotus Petal | ramp | Fast-mana artifact for early acceleration.
- 1 Mana Vault | ramp | Game Changer fast-mana artifact.
- 1 Mox Diamond | ramp | Game Changer fast-mana artifact.
- 1 Mox Opal | ramp | Fast-mana artifact for the artifact shell.
- 1 Moonsnare Prototype | ramp | Fast-mana artifact that supports early development.
- 1 Sol Ring | ramp | Fast-mana artifact for efficient acceleration.
- 1 Springleaf Drum | ramp | Fast-mana artifact for early acceleration.
- 1 Metalworker | ramp | Artifact-focused mana producer.
- 1 Chief Engineer | ramp | Artifact-focused ramp creature.
- 1 Grand Architect | ramp | Artifact-focused ramp creature.
- 1 Krark-Clan Ironworks | ramp | Artifact-based ramp piece.
- 1 Rhystic Study | draw | Game Changer draw engine.
- 1 Thoughtcast | draw | Artifact-friendly card draw.
- 1 Thought Monitor | draw | Artifact creature that provides card draw.
- 1 Sai, Master Thopterist | draw | Artifact-focused draw engine.
- 1 Thirst for Knowledge | draw | Efficient card-draw spell.
- 1 Reverse Engineer | draw | Artifact-friendly card-draw spell.
- 1 Riddlesmith | draw | Artifact-focused draw creature.
- 1 Vedalken Archmage | draw | Artifact-focused draw engine.
- 1 Forensic Gadgeteer | draw | Artifact-focused draw creature.
- 1 Nexus of Becoming | draw | Artifact card-draw piece.
- 1 One with the Machine | draw | Artifact-focused card-draw spell.
- 1 Tezzeret, Artifice Master | draw | Card-draw planeswalker and tutor.
- 1 Fierce Guardianship | interaction | Game Changer interaction spell.
- 1 Force of Will | interaction | Game Changer interaction spell.
- 1 An Offer You Can't Refuse | interaction | Interaction spell that is also marked fast mana.
- 1 Metallic Rebuke | interaction | Artifact-friendly interaction spell.
- 1 Disruption Protocol | interaction | Artifact-friendly interaction spell.
- 1 Ice Out | interaction | Interaction spell for protecting the plan.
- 1 Welding Jar | interaction | Artifact interaction piece.
- 1 Padeem, Consul of Innovation | interaction | Artifact-focused interaction creature.
- 1 Aether Spellbomb | removal | Low-cost artifact removal.
- 1 Aetherflux Reservoir | removal | Artifact removal option.
- 1 Arcum Dagsson | removal | Artifact removal creature that is also a tutor.
- 1 Cyber Conversion | removal | Removal spell for answering permanents.
- 1 Into Thin Air | removal | Removal spell for answering permanents.
- 1 Kitesail Larcenist | removal | Removal creature for the artifact deck.
- 1 Resculpt | removal | Flexible removal spell.
- 1 Ravenform | removal | Removal spell for answering permanents.
- 1 Unable to Scream | removal | Removal enchantment for answering threats.
- 1 Skysovereign, Consul Flagship | removal | Artifact vehicle that supplies removal.
- 1 Mystic Forge | synergy | Central artifact-synergy piece.
- 1 Voltaic Key | synergy | Low-cost artifact-synergy piece.
- 1 Power Artifact | synergy | Artifact-synergy enchantment.
- 1 Reshape | synergy | Artifact-synergy tutor.
- 1 Transmute Artifact | synergy | Artifact-synergy tutor.
- 1 Whir of Invention | synergy | Artifact-synergy tutor.
- 1 Kuldotha Forgemaster | threat | Artifact threat that is also a tutor.
- 1 Kappa Cannoneer | threat | Artifact creature threat.
- 1 Cyberdrive Awakener | threat | Artifact creature threat.
- 1 Phyrexian Metamorph | threat | Artifact creature threat.
- 1 Master Transmuter | threat | Artifact creature threat.
- 1 Karn, Scion of Urza | threat | Artifact-focused planeswalker threat.
- 1 Traxos, Scourge of Kroog | threat | Legendary artifact creature threat.
- 1 Lodestone Golem | threat | Artifact creature threat.
- 1 Myr Enforcer | threat | Artifact creature threat.
- 1 Jhoira's Familiar | threat | Artifact creature threat.
- 1 Engineered Explosives | wipe | Artifact board-wipe option.
- 1 Hurkyl's Recall | wipe | Board-wipe spell for resetting artifact-heavy boards.
- 1 Thassa's Oracle | wincon | Game Changer compact win condition.
- 1 Mechanized Production | wincon | Artifact-focused win condition.
- 1 Mirrodin Besieged | wincon | Artifact-focused win condition.
- 1 Mox Amber | ramp | the mana pass added it to bring the mana base inside the power level

</details>

### 4. dinosaur tribal, bracket 2

Format: Commander. Theme: dinosaurs. Pool: any_card. Shortlist: 267 names.

Commander: Gishath, Sun's Avatar.

Grade: baseline, score 0.06, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $854.39 to buy, $854.39 the whole deck.

**Summary:** This deck builds mana early and develops a Dinosaur board, using Gishath alongside a range of large Dinosaur threats to pressure opponents through combat. Creature-based draw keeps the deck moving, while Dinosaur-themed removal and a small suite of protective interaction help preserve momentum. It gives up fast, highly optimized lines for a straightforward game plan with plenty of large creatures and clear decisions for a newer player.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The curve sits high for the format, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every card supports one plan of ramping into and cheating out big Dinosaurs behind Gishath, with cost reducers, creature-based draw, and combat payoffs, the only mild tension being symmetric wipes like Wrath of God and Chain Reaction in a board-centric creature deck.
- PLAN theme_fit=yes: It is a Gishath-led Dinosaur tribal Commander deck with no tutors, combos, or fast mana beyond Sol Ring, which sits comfortably at bracket 2, and the decision points are simple enough for a newer player.
- PLAN useful_as_built=yes: 38 lands with strong Naya fixing plus ten ramp pieces support the heavy top end, and there are ample win conditions in the large threats and Gishath triggers, so the deck plays fine as written.
- PLAN summary_honest=yes: The summary accurately describes the ramp package, creature-based draw (Beast Whisperer, Guardian Project, Ripjaw Raptor, Return of the Wildspeaker), Dinosaur removal, and small protection suite, and honestly flags the high curve and unoptimized lines, though it glosses over the two symmetric board wipes.
- [INFO] `curve_summary`: average mana value 3.61 over 61 nonland cards
- [INFO] `basics_added`: the list was 1 card short, so the builder added 1 basic land

<details><summary>The deck list</summary>

- 7 Forest | land | Forms a stable part of the mana base.
- 5 Plains | land | Forms a stable part of the mana base.
- 4 Mountain | land | Forms a stable part of the mana base.
- 1 Arid Mesa | land | Adds flexible mana-base support.
- 1 Battlefield Forge | land | Adds red and white mana-base support.
- 1 Bountiful Promenade | land | Adds green and white mana-base support.
- 1 Branchloft Pathway // Boulderloft Pathway | land | Adds flexible green or white mana-base support.
- 1 Brushland | land | Adds green and white mana-base support.
- 1 Cinder Glade | land | Adds red and green mana-base support.
- 1 City of Brass | land | Adds broad mana-base support.
- 1 Command Tower | land | Adds broad mana-base support.
- 1 Cragcrown Pathway // Timbercrown Pathway | land | Adds flexible red or green mana-base support.
- 1 Forbidden Orchard | land | Adds broad mana-base support.
- 1 Karplusan Forest | land | Adds red and green mana-base support.
- 1 Mana Confluence | land | Adds broad mana-base support.
- 1 Needleverge Pathway // Pillarverge Pathway | land | Adds flexible red or white mana-base support.
- 1 Plateau | land | Adds red and white mana-base support.
- 1 Prismatic Vista | land | Adds flexible mana-base support.
- 1 Sacred Foundry | land | Adds red and white mana-base support.
- 1 Savannah | land | Adds green and white mana-base support.
- 1 Spire Garden | land | Adds red and green mana-base support.
- 1 Stomping Ground | land | Adds red and green mana-base support.
- 1 Temple Garden | land | Adds green and white mana-base support.
- 1 Windswept Heath | land | Adds flexible mana-base support.
- 1 Wooded Foothills | land | Adds flexible mana-base support.
- 1 Birds of Paradise | ramp | Provides early ramp and broad color support.
- 1 Sol Ring | ramp | Provides efficient early ramp.
- 1 Arcane Signet | ramp | Provides reliable mana acceleration.
- 1 Fellwar Stone | ramp | Provides mana acceleration.
- 1 Drover of the Mighty | ramp | Supports the Dinosaur plan while ramping.
- 1 Intrepid Paleontologist | ramp | Provides ramp for the creature-focused deck.
- 1 Ixalli's Lorekeeper | ramp | Provides early mana acceleration.
- 1 Ranging Raptors | ramp | Adds Dinosaur-themed ramp.
- 1 Thunderherd Migration | ramp | Provides Dinosaur-themed mana acceleration.
- 1 Topiary Stomper | ramp | Provides ramp on a Dinosaur body.
- 1 Beast Whisperer | draw | Supplies draw for the creature-heavy plan.
- 1 Demand Answers | draw | Provides efficient card draw.
- 1 Garruk's Uprising | draw | Provides draw support for large creatures.
- 1 Guardian Project | draw | Provides ongoing draw support.
- 1 Kutzil, Malamet Exemplar | draw | Provides draw support on a creature.
- 1 Ripjaw Raptor | draw | Adds Dinosaur-themed card draw.
- 1 Return of the Wildspeaker | draw | Provides draw support for the large-creature plan.
- 1 Runic Armasaur | draw | Adds Dinosaur-themed card draw.
- 1 Sylvan Library | draw | Provides consistent draw support.
- 1 Toski, Bearer of Secrets | draw | Provides draw support for combat-focused play.
- 1 Akroma's Will | interaction | Provides flexible interaction.
- 1 Boros Charm | interaction | Provides compact interaction support.
- 1 Heroic Intervention | interaction | Provides protective interaction.
- 1 Lightning Greaves | interaction | Provides protective interaction for key creatures.
- 1 Swiftfoot Boots | interaction | Provides protective interaction for key creatures.
- 1 Temple Altisaur | interaction | Provides interaction on a Dinosaur body.
- 1 Bronzebeak Foragers | removal | Provides removal on a Dinosaur body.
- 1 Itzquinth, Firstborn of Gishath | removal | Provides efficient Dinosaur-themed removal.
- 1 Needletooth Raptor | removal | Provides removal on a Dinosaur body.
- 1 Ravenous Sailback | removal | Provides removal on a Dinosaur body.
- 1 Savage Stomp | removal | Provides Dinosaur-themed removal.
- 1 Thrashing Brontodon | removal | Provides removal on a Dinosaur body.
- 1 Trumpeting Carnosaur | removal | Provides removal on a Dinosaur body.
- 1 Chain Reaction | wipe | Provides a board wipe.
- 1 Wrath of God | wipe | Provides a dependable board wipe.
- 1 Wakening Sun's Avatar | wipe | Provides a Dinosaur-themed board wipe.
- 1 Amped Raptor | synergy | Supports the Dinosaur theme.
- 1 Belligerent Yearling | synergy | Supports the Dinosaur theme.
- 1 Commune with Dinosaurs | synergy | Supports the Dinosaur-focused plan.
- 1 Dinosaur Stampede | synergy | Supports Dinosaur combat turns.
- 1 Huatli's Raptor | synergy | Supports the Dinosaur theme.
- 1 Kinjalli's Caller | synergy | Supports the Dinosaur-focused plan.
- 1 Kinjalli's Sunwing | synergy | Provides Dinosaur-themed support.
- 1 Marauding Raptor | synergy | Supports the Dinosaur-focused plan.
- 1 Otepec Huntmaster | synergy | Supports the Dinosaur-focused plan.
- 1 Raptor Companion | synergy | Adds to the Dinosaur creature base.
- 1 Relentless Raptor | synergy | Adds to the Dinosaur creature base.
- 1 Sky Terror | synergy | Adds to the Dinosaur creature base.
- 1 Sunfrill Imitator | synergy | Provides Dinosaur-themed support.
- 1 Carnage Tyrant | threat | Provides a substantial Dinosaur threat.
- 1 Etali, Primal Storm | threat | Provides a legendary Dinosaur threat.
- 1 Ghalta, Primal Hunger | threat | Provides a major Dinosaur threat.
- 1 Ghalta, Stampede Tyrant | threat | Provides a major Dinosaur threat.
- 1 Goring Ceratops | threat | Provides a Dinosaur combat threat.
- 1 Quartzwood Crasher | threat | Provides a Dinosaur combat threat.
- 1 Regisaur Alpha | threat | Provides a Dinosaur threat.
- 1 Shifting Ceratops | threat | Provides a Dinosaur threat.
- 1 Thundering Spineback | threat | Provides a Dinosaur threat.
- 1 Tyrranax Rex | threat | Provides a major Dinosaur threat.
- 1 Zetalpa, Primal Dawn | threat | Provides a major Dinosaur threat.
- 1 Zilortha, Strength Incarnate | threat | Provides a legendary Dinosaur threat.

</details>

### 5. blink Commander, owned first

Format: Commander. Theme: blink. Pool: owned_first. Shortlist: 168 names.

Commander: Gilraen, Dúnedain Protector.

Grade: typical, score 0.38, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $123.76 the whole deck.

**Summary:** This is a mono-white blink deck built around Gilraen and a creature-heavy battlefield. Flickerwisp, Personify, and creature-based draw and removal keep the deck moving while Angels, legendary creatures, and equipment-bearing bodies provide the pressure to close games. It wins through a sustained board presence backed by flexible answers and protection, giving up explosive speed and tutoring for a steadier, board-oriented bracket 3 plan.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core of ETB creatures, protection and removal hangs together, but the list pulls in several directions at once — an equipment package (Puresteel Paladin, Sword of the Animist, Mask of Memory, Skullclamp), a legendary-matters package (Plaza of Heroes, Bastion Protector, Boromir), an Angel package, and only a handful of actual repeatable blink enablers (Flickerwisp, Personify, Angel of Condemnation) to tie the ETB bodies together.
- PLAN theme_fit=partly: It is a mono-white Gilraen Commander deck at a plausible bracket-3 level (three wraths, no fast mana beyond Sol Ring, no tutors), but the blink theme the person asked for is thin — a few flicker effects and ETB creatures amid a lot of generic white removal, equipment and anthem-style protection.
- PLAN useful_as_built=yes: 36 lands plus nine mana rocks and Wayfarer's Bauble support a low, creature-dense curve in a single color, and roughly 25 creatures with equipment and Gilraen's protection give a clear, playable route to closing games.
- PLAN summary_honest=yes: placeholder
- [INFO] `curve_summary`: average mana value 2.84 over 63 nonland cards

<details><summary>The deck list</summary>

- 25 Plains | land | Primary untapped white mana base.
- 1 Command Tower | land | Reliable white source for the commander deck.
- 1 City of Brass | land | Flexible colored mana source.
- 1 Exotic Orchard | land | Additional flexible colored mana source.
- 1 Marsh Flats | land | Fetch land that supports the white-heavy mana base.
- 1 Fabled Passage | land | Fetch land for mana fixing.
- 1 Evolving Wilds | land | Fetch land that finds a Plains.
- 1 Terramorphic Expanse | land | Additional fetch land for Plains access.
- 1 Ash Barrens | land | Land slot with basic-land fixing.
- 1 Grand Coliseum | land | Flexible colored mana source.
- 1 Plaza of Heroes | land | White source suited to the legendary commander.
- 1 Secluded Courtyard | land | Colored source for the creature-heavy deck.
- 1 Sol Ring | ramp | Efficient early mana acceleration.
- 1 Arcane Signet | ramp | Reliable mana fixing and acceleration.
- 1 Fellwar Stone | ramp | Low-cost mana acceleration.
- 1 Wayfarer's Bauble | ramp | Early land-based ramp.
- 1 Thought Vessel | ramp | Mana acceleration in an artifact slot.
- 1 Bender's Waterskin | ramp | Additional mana development for the deck.
- 1 Sword of the Animist | ramp | Equipment-based mana acceleration for the creature plan.
- 1 Relic of Legends | ramp | Mana acceleration that works with the legendary commander.
- 1 Chromatic Lantern | ramp | Mana fixing and acceleration for the deck.
- 1 Giada, Font of Hope | ramp | Creature-based ramp that also contributes to the Angel presence.
- 1 Wall of Omens | draw | Low-cost creature-based card advantage for blink turns.
- 1 Inspiring Overseer | draw | Creature-based card advantage that fits the blink shell.
- 1 Lembas | draw | Compact artifact card advantage.
- 1 Idol of Oblivion | draw | Repeatable artifact draw slot.
- 1 Skullclamp | draw | Efficient equipment-based card advantage.
- 1 Tome of Legends | draw | Card advantage alongside the commander.
- 1 Mask of Memory | draw | Equipment-based card advantage for attacking creatures.
- 1 Mirror of Galadriel | draw | Legendary artifact draw source.
- 1 Diary of Dreams | draw | Artifact card advantage for longer games.
- 1 Energybending | draw | Instant-speed card advantage.
- 1 Champions of Minas Tirith | draw | Creature-based draw that supports the board plan.
- 1 Angel of Condemnation | removal | Creature-based removal that fits the blink-oriented creature suite.
- 1 Angel of Sanctions | removal | Angel removal body for the deck's board plan.
- 1 Angel of Serenity | removal | High-impact Angel removal option.
- 1 Fiend Hunter | removal | Creature removal that belongs in a blink-focused shell.
- 1 Palace Jailer | removal | Creature-based removal with added table presence.
- 1 Swords to Plowshares | removal | Efficient single-target removal.
- 1 Generous Gift | removal | Flexible answer to opposing permanents.
- 1 Get Lost | removal | Versatile low-cost removal.
- 1 Stroke of Midnight | removal | Broad instant-speed permanent answer.
- 1 Crib Swap | removal | Tribal instant removal for a creature-heavy deck.
- 1 Dispatch | removal | Low-cost targeted removal.
- 1 Journey to Nowhere | removal | Additional permanent-based creature removal.
- 1 Slip On the Ring | interaction | Blink-oriented protection and interaction.
- 1 Clever Concealment | interaction | Protects the developed board from opposing answers.
- 1 Lightning Greaves | interaction | Protects the commander or a key creature.
- 1 Swiftfoot Boots | interaction | Additional protection for important creatures.
- 1 Bastion Protector | interaction | Creature-based commander protection.
- 1 Boromir, Warden of the Tower | interaction | Legendary creature interaction for protecting the board plan.
- 1 Reprieve | interaction | Flexible instant-speed disruption.
- 1 Unbreakable Formation | interaction | Protective interaction for a creature board.
- 1 Duty Beyond Death | interaction | Protection for a key creature or commander.
- 1 Gift of Immortality | interaction | Persistent protection for an important creature.
- 1 Together Forever | interaction | Protective interaction that supports the creature plan.
- 1 Austere Command | wipe | Flexible reset button against developed boards.
- 1 Fumigate | wipe | Board reset for creature-heavy opposing boards.
- 1 Vanquish the Horde | wipe | Efficient mass-creature answer.
- 1 Flickerwisp | synergy | Core blink synergy creature.
- 1 Personify | synergy | Dedicated blink-synergy instant.
- 1 Ennis, Debate Moderator | synergy | Synergy creature for the deck's central plan.
- 1 Jocasta, Automaton Avenger | synergy | Additional synergy permanent for the creature-focused shell.
- 1 Aang, the Last Airbender | threat | Legendary creature that supplies meaningful board pressure.
- 1 Exemplar of Light | threat | Angel creature that contributes to the attacking board.
- 1 Faramir, Field Commander | threat | Legendary creature that adds board presence.
- 1 Frontline Medic | threat | Creature threat that supports a committed board.
- 1 Kataki, War's Wage | threat | Creature threat with useful table presence.
- 1 Puresteel Paladin | threat | Creature threat that pairs well with the equipment package.
- 1 South Pole Voyager | threat | Creature threat for the deck's board-focused game.
- 1 Stiltzkin, Moogle Merchant | threat | Legendary creature that adds to the battlefield plan.
- 1 The Vision | threat | Artifact creature threat with resilient board presence.
- 1 Weapons Vendor | threat | Creature threat that complements the deck's equipment cards.
- 1 Westfold Rider | threat | Creature threat that adds combat pressure.
- 1 Bronze Guardian | threat | Artifact creature threat that complements the artifact package.

</details>

### 6. Modern tempo, tournament

Format: Modern. Theme: tempo. Pool: any_card. Shortlist: 167 names.

Grade: typical, score 0.59, model 20260914T154223Z.

Cards: 64 main, 15 sideboard. Repair turn: no. Block findings: 0.

Cost: $473.33 to buy, $473.33 the whole deck.

**Summary:** This blue-red tempo deck establishes pressure with Ragavan, Nimble Pilferer, Ledger Shredder, and Faerie Mastermind, then uses efficient removal and interaction to preserve its lead. It wins by staying ahead on the board while its creatures finish the game, giving up larger late-game threats and sweeping effects for a focused, proactive plan.

The quality model grades this deck typical against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The deck runs its spells as playsets, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade.

- PLAN plan_coherent=yes: Every card serves one low-curve tempo plan: cheap evasive/value threats, four cantrips plus Bauble to fuel Ledger Shredder's connive and Ragavan's dash-turns, and a suite of one- and two-mana removal and counters to protect the clock.
- PLAN theme_fit=partly: It is a Modern blue-red tempo deck as asked, but the person specifically asked for a Delver deck and the list contains no Delver of Secrets nor any other one-mana threat in its place, so it is the adjacent 'Izzet Prowess/Shredder' build rather than the requested one.
- PLAN useful_as_built=partly: The mana base is clean and the curve tops at two so the spells always cast, but 28 lands behind a deck whose most expensive card costs two is far more land than this plan can use and the deck will flood out regularly.
- PLAN summary_honest=yes: The summary accurately describes what the list does and openly concedes what it gives up (big late-game threats, sweepers); it names only threats and effects that are actually present.
- [WARN] `land_count`: 28 lands: the guide range for this format is 20 to 27
- [INFO] `curve_summary`: average mana value 1.22 over 36 nonland cards

<details><summary>The deck list</summary>

- 4 Ragavan, Nimble Pilferer | threat | Early creature pressure for the tempo plan.
- 4 Ledger Shredder | threat | A creature that keeps pressure on the opponent.
- 4 Faerie Mastermind | threat | Adds evasive creature pressure to the threat base.
- 4 Consider | draw | Low-cost card selection for finding threats, lands, and answers.
- 2 Preordain | draw | Additional card selection that keeps the deck consistent.
- 4 Counterspell | interaction | Core broad interaction for protecting the tempo lead.
- 2 Spell Pierce | interaction | Efficient interaction that supports a proactive game plan.
- 4 Lightning Bolt | removal | Efficient removal that also supports closing games.
- 4 Galvanic Discharge | removal | Additional instant-speed removal for opposing threats.
- 4 Mishra's Bauble | synergy | Supports the deck's low-cost spell synergy.
- 8 Island | land | Reliable blue source for the deck's blue spells.
- 4 Mountain | land | Reliable red source for the deck's red spells.
- 4 Steam Vents | land | Dual-color land that supplies both primary colors.
- 4 Riverglide Pathway // Lavaglide Pathway | land | Flexible dual-color source for the streamlined mana base.
- 4 Shivan Reef | land | Untapped dual-color source for early plays.
- 4 Spirebluff Canal | land | Dual-color land suited to the deck's early game.

</details>

### 7. Modern burn, casual

Format: Modern. Theme: burn. Pool: any_card. Shortlist: 290 names.

Grade: typical, score 0.46, model 20260914T154223Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $39.42 to buy, $39.42 the whole deck.

**Summary:** This mono-red burn deck applies pressure with a steady stream of threats while using direct removal to clear resistance and keep the damage plan moving. Its card-access spells help it continue presenting action, and its synergy pieces reward the deck for leaning heavily into red pressure. The deck gives up broad answers and a flexible mana base in exchange for a focused, consistent attack.

The quality model grades this deck typical against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=no: The pieces pull apart: Eidolon of the Great Revel and Thermo-Alchemist want a low curve stuffed with cheap burn spells, but the deck has only six actual damage spells and instead fills up on four-drop creatures like Ashcloud Phoenix, Tectonic Giant, and Hazoret alongside 24 lands, so no single plan is served.
- PLAN theme_fit=no: The colors and format match, but the request was a burn deck and this is a clunky mono-red creature midrange pile with only Lightning Bolt and Lava Dart as real reach, missing staples like Lava Spike, Skewer the Critics, or Boros Charm entirely.
- PLAN useful_as_built=partly: It will cast its spells off 24 Mountains and can win with creature beats, but the land count is several too high for an aggro shell, Ancestral Anger is close to a dead card, and the deck lacks the reach to close games it claims to pressure.
- PLAN summary_honest=partly: It sidesteps the word 'burn' and honestly notes the lack of answers, but calling Ancestral Anger and two Risk Factors 'card-access spells' that keep it 'presenting action,' and framing the synergy pieces as rewarded for 'leaning heavily into red pressure,' overstates what a deck with six burn spells and a 24-land creature curve actually does.
- [INFO] `curve_summary`: average mana value 2.61 over 36 nonland cards

<details><summary>The deck list</summary>

- 24 Mountain | land | Provides a consistent mono-red mana base.
- 2 Chandra, Dressed to Kill | ramp | Supplies the deck's allotted acceleration while staying in red.
- 4 Ancestral Anger | draw | Helps maintain access to more action.
- 2 Risk Factor | draw | Adds additional card access for the burn plan.
- 4 Lightning Bolt | removal | Efficiently supports the deck's direct-damage game plan.
- 2 Lava Dart | removal | Provides additional low-cost removal.
- 4 Eidolon of the Great Revel | synergy | Supports the deck's pressure-oriented burn shell.
- 4 Thermo-Alchemist | synergy | Works alongside the deck's large instant and sorcery package.
- 4 Akki Lavarunner // Tok-Tok, Volcano Born | threat | Provides an early threat to apply pressure.
- 4 Ashcloud Phoenix | threat | Adds a resilient threat slot for longer games.
- 2 Hazoret the Fervent | threat | Provides a powerful top-end threat.
- 2 Torbran, Thane of Red Fell | threat | Adds a high-impact threat to support the red damage plan.
- 2 Tectonic Giant | threat | Rounds out the threat suite with additional staying power.

</details>

### 8. Modern lifegain, FNM

Format: Modern. Theme: lifegain. Pool: any_card. Shortlist: 301 names.

Grade: bad, score 0.34, model 20260914T154223Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $344.60 to buy, $344.60 the whole deck.

**Summary:** This white-black lifegain deck establishes early synergy pieces, turns repeated life gain into creature pressure, and uses efficient removal to clear a path for its threats. It wins primarily by building a board that rewards the lifegain plan and attacking through the opponent’s defenses. The focused creature-and-synergy approach gives up broader utility and relies on keeping its key permanents in play.

The quality model grades this deck below the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade.

- PLAN plan_coherent=yes: Nearly every nonland card either produces repeated life gain (Soul Warden, Attended Healer, Lembas, Solitude) or pays it off (Pridemate, Twinblade Paladin, Bloodbond Vampire, Cliffhaven Vampire, Enduring Tenacity, Dawn of Hope), so the pieces pull in one direction, with only minor tension from the painlands.
- PLAN theme_fit=yes: It is a Modern-legal white-black lifegain creature deck at a reasonable FNM power level, exactly what was asked.
- PLAN useful_as_built=yes: 24 lands with solid dual/painland fixing support a two-to-four-mana curve, and the payoff creatures plus Cliffhaven Vampire and lifelinkers give clear ways to close a game.
- PLAN summary_honest=yes: The summary accurately describes a creature-based lifegain board that leans on removal and admits its dependence on keeping permanents around, claiming nothing the list lacks.
- [INFO] `curve_summary`: average mana value 3.11 over 36 nonland cards

<details><summary>The deck list</summary>

- 6 Plains | land | Provides a reliable white source for the deck's early plays.
- 6 Swamp | land | Provides a reliable black source for the deck's black spells.
- 4 Godless Shrine | land | Supplies both deck colors without using a tapped-land slot.
- 4 Caves of Koilos | land | Adds flexible white-black mana for the two-color base.
- 4 Mana Confluence | land | Completes the color requirements while keeping the mana base untapped.
- 2 Altar of the Pantheon | ramp | Fills the ramp slot while fitting the deck's colors.
- 4 Solitude | removal | Provides the deck's most efficient removal package.
- 2 Murderous Rider // Swift End | removal | Adds versatile black removal alongside a creature body.
- 4 Lembas | draw | Low-cost card access that supports the deck's lifegain theme.
- 2 Dawn of Hope | draw | Repeatable draw support for a deck built around gaining life.
- 4 Soul Warden | synergy | A core early lifegain synergy piece.
- 4 Ajani's Pridemate | synergy | Turns the deck's lifegain plan into growing battlefield pressure.
- 4 Attended Healer | threat | A lifegain-focused creature that serves as a primary threat.
- 4 Twinblade Paladin | threat | A white threat that fits naturally alongside the lifegain package.
- 2 Bloodbond Vampire | threat | A black lifegain payoff that helps pressure the opponent.
- 2 Enduring Tenacity | threat | A resilient black threat for the lifegain strategy.
- 2 Cliffhaven Vampire | threat | Adds a black lifegain-oriented finishing threat.

</details>

### 9. Standard midrange, FNM

Format: Standard. Theme: midrange. Pool: any_card. Shortlist: 164 names.

Grade: baseline, score 0.42, model 20260914T154223Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $137.70 to buy, $137.70 the whole deck.

**Summary:** This black-green midrange deck establishes a creature presence early, uses efficient removal to keep opposing pressure in check, and maintains resources for longer exchanges. It wins by steadily attacking with its creature package while protecting key threats and drawing into more action. In return, it gives up explosive speed for a measured, board-focused game plan.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The cards pair in ways the top lists do not, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: Most pieces do point one way — resilient creatures, cheap removal, and repeating card draw — but the four Undying Malice are a sacrifice-payoff/combat-trick card with no sacrifice outlets or death triggers in the list, so eight protection spells supporting only fourteen creatures pulls against the grindy card-advantage plan the rest of the deck is built on.
- PLAN theme_fit=yes: It is a two-color black-green Standard creature-and-removal midrange list at a casual FNM power level, exactly the request, with no stray colors, formats, or ambitions.
- PLAN useful_as_built=yes: Twenty-four lands with sixteen untapped-capable duals support a clean one-through-five curve, and fourteen creatures plus six removal spells and recurring draw give it enough board presence and staying power to be picked up and played as-is.
- PLAN summary_honest=yes: Every claim — early creatures, efficient removal, resource upkeep, protecting threats, no explosive speed — maps to cards actually in the list (Bitter Triumph/Trophy, Phyrexian Arena/Unholy Annex, Snakeskin Veil), and it openly concedes the deck's slowness rather than overselling it.
- [INFO] `curve_summary`: average mana value 2.25 over 36 nonland cards

<details><summary>The deck list</summary>

- 4 Overgrown Tomb | land | A black-green land slot that supports both primary colors.
- 4 Underground Mortuary | land | A black-green land slot that deepens access to both primary colors.
- 4 Blooming Marsh | land | A black-green land slot for an efficient two-color mana base.
- 4 Deathcap Glade | land | A black-green land slot that rounds out the dual-color mana base.
- 4 Forest | land | Reliable green mana for the deck's green spells.
- 4 Swamp | land | Reliable black mana for the deck's black spells.
- 2 Llanowar Elves | ramp | Early mana acceleration helps the deck move ahead into its midgame.
- 4 Phyrexian Arena | draw | A dedicated card-advantage engine for longer games.
- 2 Unholy Annex // Ritual Chamber | draw | Additional card advantage that keeps the deck supplied in midrange battles.
- 4 Bitter Triumph | removal | Efficient removal for clearing opposing problems.
- 2 Assassin's Trophy | removal | Flexible removal that answers a wide range of opposing cards.
- 4 Snakeskin Veil | synergy | Protects the creature-heavy threat package and reinforces its board presence.
- 4 Undying Malice | synergy | Supports the creature plan by helping key threats stay relevant through removal-heavy exchanges.
- 4 Goldvein Hydra | threat | A scalable green creature that serves as a central offensive threat.
- 4 Chomping Changeling | threat | A creature body that contributes to the deck's steady board pressure.
- 3 Thrashing Brontodon | threat | A sturdy green creature for applying pressure in the midgame.
- 3 Syr Ginger, the Meal Ender | threat | A low-cost creature threat that helps start pressure early.

</details>

### 10. Standard aggro, tournament

Format: Standard. Theme: aggro. Pool: any_card. Shortlist: 294 names.

Grade: bad, score 0.09, model 20260914T154223Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $59.65 to buy, $59.65 the whole deck.

**Summary:** This red-white aggro deck aims to establish pressure with its threat package and turn that pressure into a fast win through attacking. Its removal and interaction clear a path or protect the plan, while its draw cards help it keep presenting action after the opening push. It gives up heavier late-game cards in favor of a streamlined, proactive game.

The quality model grades this deck below the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. Many of the cards appear in no top list, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core of one- to three-drop aggressive creatures plus burn and cheap removal points one way, but the 5-mana Redcap Gutter-Dweller, the huge Krenko's Buzzcrusher, and a lone Springleaf Drum are stray pieces that pull toward a slower, clunkier game the rest of the list does not support.
- PLAN theme_fit=partly: It is largely a red-white beatdown list for Standard as asked, but several slots (a black-red five-drop, a seven-mana finisher, a singleton mana rock) answer a different, slower brief than red-white aggro.
- PLAN useful_as_built=partly: 24 lands with only six untapped duals will cast the cheap double-costed spells often enough to function, and the creature count gives real clocks, but the expensive off-curve cards will regularly be stranded in hand.
- PLAN summary_honest=partly: The threat/removal/draw description roughly matches, but the claim that it "gives up heavier late-game cards in favor of a streamlined, proactive game" is contradicted by the top end it actually plays, and 'draw' is a generous label for Codebreaker and Tersa.
- [INFO] `curve_summary`: average mana value 2.97 over 36 nonland cards
- [INFO] `mana_pass`: the builder moved 1 card of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 10 Plains | land | Provides a reliable white source for the deck's white spells.
- 8 Mountain | land | Provides a reliable red source for the deck's red spells.
- 4 Sacred Foundry | land | Supplies both deck colors without adding a dedicated tapped land.
- 2 Inspiring Vantage | land | Adds flexible red-white mana for the aggressive opening turns.
- 4 Fugitive Codebreaker | draw | A red draw card that supports the deck's proactive game plan.
- 2 Tersa Lightshatter | draw | Adds more red card access without leaving the aggressive plan.
- 4 Boros Charm | interaction | A flexible red-white interaction spell for protecting the attack plan.
- 2 Sheltered by Ghosts | interaction | Provides additional white interaction alongside the creature pressure.
- 4 Emeritus of Truce // Swords to Plowshares | removal | Efficient white removal to clear opposing resistance.
- 3 Krenko's Buzzcrusher | removal | Red removal that keeps the deck's pressure supported.
- 4 Slickshot Show-Off | synergy | A focused red synergy piece for the aggressive shell.
- 4 Redcap Gutter-Dweller | threat | A red threat that contributes to the deck's attacking plan.
- 4 Dragonback Lancer | threat | A white threat that broadens the creature pressure.
- 4 Frilled Sparkshooter | threat | A red threat that helps maintain a fast clock.
- 1 Springleaf Drum | ramp | the mana pass added it to bring the mana base inside the power level

</details>

### 11. Commander with a locked card

Format: Commander. Theme: sacrifice. Pool: any_card. Shortlist: 296 names.

Commander: Karlov of the Ghost Council.

Grade: baseline, score 0.08, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $713.59 to buy, $713.59 the whole deck.

**Summary:** Karlov leads a black-white sacrifice shell that develops mana and card flow while assembling sacrifice synergies around a broad creature package. The deck pressures the table through its threats, uses removal and sweeping effects to control opposing boards, and aims to turn a stable sacrifice engine into a winning position. It gives up some speed to maintain a dependable two-color mana base and a substantial set of answers.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. Few lands enter tapped, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every card is a sacrifice outlet, fodder generator, death-trigger drain, or sac payoff, and the removal/wipe suite (Grave Pact, Dictate of Erebos, Meathook Massacre) reinforces the same aristocrats plan rather than pulling elsewhere.
- PLAN theme_fit=yes: It is a black-white Commander deck led by Karlov of the Ghost Council built squarely around sacrifice, with lifegain drains (Zulaport Cutthroat, Bastion of Remembrance, Elas il-Kor, Vito's Inquisitor) that feed the commander's counters.
- PLAN useful_as_built=yes: Thirty-six lands with a clean two-color base, a low curve, plenty of cheap ramp and outlets, and multiple win routes (drain triggers, Altar of Dementia mill, Karlov beats) make it immediately playable.
- PLAN summary_honest=yes: The claims of ramp, card flow, removal, sweepers, and a sacrifice engine all map to real cards in the list; it is vague about Karlov's lifegain angle but asserts nothing the deck lacks.
- [INFO] `curve_summary`: average mana value 3.13 over 63 nonland cards
- [INFO] `mana_pass`: the builder moved 3 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 13 Plains | land | Provides a reliable white mana base.
- 14 Swamp | land | Provides a reliable black mana base.
- 1 Brightclimb Pathway // Grimclimb Pathway | land | Flexible dual-color land slot.
- 1 Caves of Koilos | land | Dual-color land slot for the mana base.
- 1 Command Tower | land | Dual-color land slot for the mana base.
- 1 Concealed Courtyard | land | Dual-color land slot for the mana base.
- 1 Godless Shrine | land | Dual-color land slot for the mana base.
- 1 Isolated Chapel | land | Dual-color land slot for the mana base.
- 1 Scrubland | land | Dual-color land slot for the mana base.
- 1 Shattered Sanctum | land | Dual-color land slot for the mana base.
- 1 Vault of Champions | land | Dual-color land slot for the mana base.
- 1 Bushmeat Poacher | draw | Included as a draw card for the sacrifice shell.
- 1 Corrupted Conviction | draw | Included as a draw card for the sacrifice shell.
- 1 Disciple of Bolas | draw | Included as a draw card for the sacrifice shell.
- 1 Ecstatic Awakener // Awoken Demon | draw | Included as a draw card for the sacrifice shell.
- 1 Jerren, Corrupted Bishop // Ormendahl, the Corrupter | draw | Included as a draw card for the sacrifice shell.
- 1 Relic Vial | draw | Included as a draw card for the sacrifice shell.
- 1 Shadowheart, Dark Justiciar | draw | Included as a draw card for the sacrifice shell.
- 1 Smothering Abomination | draw | Included as a draw card for the sacrifice shell.
- 1 Vampiric Rites | draw | Included as a draw card for the sacrifice shell.
- 1 Village Rites | draw | Included as a draw card for the sacrifice shell.
- 1 Vampire Gourmand | draw | Included as a draw card for the sacrifice shell.
- 1 Cartel Aristocrat | interaction | Included to provide interaction in the sacrifice plan.
- 1 Dark Privilege | interaction | Included to provide interaction in the sacrifice plan.
- 1 Fanatical Devotion | interaction | Included to provide interaction in the sacrifice plan.
- 1 Flare of Fortitude | interaction | Included to provide interaction in the sacrifice plan.
- 1 Gift of Doom | interaction | Included to provide interaction in the sacrifice plan.
- 1 Kinzu of the Bleak Coven | interaction | Included to provide interaction in the sacrifice plan.
- 1 Promise of Tomorrow | interaction | Included to provide interaction in the sacrifice plan.
- 1 Rescue from the Underworld | interaction | Included to provide interaction in the sacrifice plan.
- 1 Sol Ring | ramp | Required card that supplies a ramp slot.
- 1 Ashnod's Altar | ramp | Included to supply ramp for the sacrifice plan.
- 1 Culling the Weak | ramp | Included to supply ramp for the sacrifice plan.
- 1 Deadly Dispute | ramp | Included to supply ramp for the sacrifice plan.
- 1 Pawn of Ulamog | ramp | Included to supply ramp for the sacrifice plan.
- 1 Phyrexian Altar | ramp | Included to supply ramp for the sacrifice plan.
- 1 Pitiless Plunderer | ramp | Included to supply ramp for the sacrifice plan.
- 1 Priest of Forgotten Gods | ramp | Included to supply ramp for the sacrifice plan.
- 1 Skullport Merchant | ramp | Included to supply ramp for the sacrifice plan.
- 1 Warren Soultrader | ramp | Included to supply ramp for the sacrifice plan.
- 1 Attrition | removal | Included as removal for the sacrifice shell.
- 1 Ayli, Eternal Pilgrim | removal | Included as removal for the sacrifice shell.
- 1 Blasting Station | removal | Included as removal for the sacrifice shell.
- 1 Bone Shards | removal | Included as removal for the sacrifice shell.
- 1 Bone Splinters | removal | Included as removal for the sacrifice shell.
- 1 Dictate of Erebos | removal | Included as removal for the sacrifice shell.
- 1 Eaten Alive | removal | Included as removal for the sacrifice shell.
- 1 Grave Pact | removal | Included as removal for the sacrifice shell.
- 1 Yawgmoth, Thran Physician | removal | Included as removal for the sacrifice shell.
- 1 Altar of Dementia | synergy | Included for sacrifice synergy.
- 1 Bartolomé del Presidio | synergy | Included for sacrifice synergy.
- 1 Bastion of Remembrance | synergy | Included for sacrifice synergy.
- 1 Carrion Feeder | synergy | Included for sacrifice synergy.
- 1 Elas il-Kor, Sadistic Pilgrim | synergy | Included for sacrifice synergy.
- 1 Fleshtaker | synergy | Included for sacrifice synergy.
- 1 Hidden Stockpile | synergy | Included for sacrifice synergy.
- 1 Viscera Seer | synergy | Included for sacrifice synergy.
- 1 Woe Strider | synergy | Included for sacrifice synergy.
- 1 Zulaport Cutthroat | synergy | Included for sacrifice synergy.
- 1 Abhorrent Overlord | threat | Included as a major threat.
- 1 Basri's Lieutenant | threat | Included as a major threat.
- 1 Felisa, Fang of Silverquill | threat | Included as a major threat.
- 1 Ghoulcaller Gisa | threat | Included as a major threat.
- 1 Liesa, Forgotten Archangel | threat | Included as a major threat.
- 1 Mondrak, Glory Dominus | threat | Included as a major threat.
- 1 Ratadrabik of Urborg | threat | Included as a major threat.
- 1 Razaketh, the Foulblooded | threat | Included as a major threat.
- 1 Requiem Angel | threat | Included as a major threat.
- 1 Sidisi, Undead Vizier | threat | Included as a major threat.
- 1 Vindictive Vampire | threat | Included as a major threat.
- 1 Vito's Inquisitor | threat | Included as a major threat.
- 1 Austere Command | wipe | Included as a reset option.
- 1 The Meathook Massacre | wipe | Included as a reset option.
- 1 Toxic Deluge | wipe | Included as a reset option.

</details>

### 12. Commander on a budget

Format: Commander. Theme: tokens. Pool: any_card. Shortlist: 296 names.

Commander: Adeline, Resplendent Cathar.

Grade: bad, score 0.03, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $73.14 to buy, $73.14 the whole deck.

**Summary:** Adeline leads a white token deck that builds a broad board, turns that board into combat pressure, and closes through its supporting threats. It carries draw and mana development to keep deploying pieces, while removal, protective interaction, and board wipes help it recover the initiative when opponents establish stronger positions. The deck gives up premium acceleration and expensive token staples in favor of a straightforward, lower-powered token plan.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The deck makes less mana on turn four than the norm, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every card feeds one plan — make wide white tokens, pump them with anthems/Divine Visitation/Intangible Virtue, and swing with Adeline's trigger — with only the two symmetric wraths (Martial Coup, Phyrexian Rebirth) pulling slightly against the board the deck builds.
- PLAN theme_fit=yes: It is a mono-white Adeline token Commander deck built from cheap, mostly non-staple cards with no fast mana or high-end bombs, which matches the requested budget and bracket-2 power level.
- PLAN useful_as_built=yes: 38 lands in a nearly mono-colored deck with a low, deployable curve and many redundant token producers plus anthems means it can be shuffled up and played, and it wins straightforwardly through wide combat damage.
- PLAN summary_honest=yes: Every element the summary names — token making, combat pressure, card draw engines, modest mana rocks, spot removal, protection like Rootborn Defenses/Sunpearl Kirin, and wipes — is actually present, and it openly admits it skips premium acceleration and expensive staples.
- [INFO] `curve_summary`: average mana value 3.23 over 62 nonland cards

<details><summary>The deck list</summary>

- 23 Plains | land | Provides the deck's primary white mana source.
- 1 Ancient Den | land | Provides a listed land slot.
- 1 Castle Ardenvale | land | Provides a listed land slot.
- 1 Command Tower | land | Provides a listed land slot.
- 1 Eiganjo, Seat of the Empire | land | Provides a listed land slot.
- 1 Exotic Orchard | land | Provides a listed land slot.
- 1 Path of Ancestry | land | Provides a listed land slot.
- 1 Plaza of Heroes | land | Provides a listed land slot.
- 1 Secluded Courtyard | land | Provides a listed land slot.
- 1 Unclaimed Territory | land | Provides a listed land slot.
- 1 Windbrisk Heights | land | Provides a listed land slot.
- 1 Fountainport | land | Provides a listed land slot.
- 1 Hall of Tagsin | land | Provides a listed land slot.
- 1 Mirrex | land | Provides a listed land slot.
- 1 Urza's Factory | land | Provides a listed land slot.
- 1 Collector's Vault | ramp | Chosen for its listed ramp role.
- 1 Currency Converter | ramp | Chosen for its listed ramp role.
- 1 Druidic Satchel | ramp | Chosen for its listed ramp role.
- 1 Goldvein Pick | ramp | Chosen for its listed ramp role.
- 1 Idol of False Gods | ramp | Chosen for its listed ramp role.
- 1 Karn, Living Legacy | ramp | Chosen for its listed ramp role.
- 1 Keeper of the Accord | ramp | Chosen for its listed ramp role.
- 1 Legion's Landing // Adanto, the First Fort | ramp | Chosen for its listed ramp role.
- 1 Monologue Tax | ramp | Chosen for its listed ramp role.
- 1 The Restoration of Eiganjo // Architect of Restoration | ramp | Chosen for its listed ramp role.
- 1 Angelic Sell-Sword | draw | Chosen for its listed draw role.
- 1 Bygone Bishop | draw | Chosen for its listed draw role.
- 1 Court of Grace | draw | Chosen for its listed draw role.
- 1 Dawn of Hope | draw | Chosen for its listed draw role.
- 1 Faramir, Field Commander | draw | Chosen for its listed draw role.
- 1 Glimmer Seeker | draw | Chosen for its listed draw role.
- 1 Idol of Oblivion | draw | Chosen for its listed draw role.
- 1 Search the Premises | draw | Chosen for its listed draw role.
- 1 Staff of the Storyteller | draw | Chosen for its listed draw role.
- 1 Sunpearl Kirin | draw | Chosen for its listed draw role.
- 1 Wedding Announcement // Wedding Festivity | draw | Chosen for its listed draw role.
- 1 Aerial Assault | removal | Chosen for its listed removal role.
- 1 Banishing Slash | removal | Chosen for its listed removal role.
- 1 Battle Menu | removal | Chosen for its listed removal role.
- 1 Citizen's Crowbar | removal | Chosen for its listed removal role.
- 1 Generous Gift | removal | Chosen for its listed removal role.
- 1 Kellan's Lightblades | removal | Chosen for its listed removal role.
- 1 Skyclave Apparition | removal | Chosen for its listed removal role.
- 1 Hour of Reckoning | wipe | Chosen for its listed wipe role.
- 1 Martial Coup | wipe | Chosen for its listed wipe role.
- 1 Phyrexian Rebirth | wipe | Chosen for its listed wipe role.
- 1 Aligned Heart | synergy | Chosen for its listed synergy role.
- 1 Anafenza, Unyielding Lineage | synergy | Chosen for its listed synergy role.
- 1 Animation Module | synergy | Chosen for its listed synergy role.
- 1 Automated Assembly Line | synergy | Chosen for its listed synergy role.
- 1 Cat Collector | synergy | Chosen for its listed synergy role.
- 1 Clarion Spirit | synergy | Chosen for its listed synergy role.
- 1 Divine Visitation | synergy | Chosen for its listed synergy role.
- 1 Felidar Retreat | synergy | Chosen for its listed synergy role.
- 1 Hanweir Militia Captain // Westvale Cult Leader | synergy | Chosen for its listed synergy role.
- 1 Horn of Gondor | synergy | Chosen for its listed synergy role.
- 1 Intangible Virtue | synergy | Chosen for its listed synergy role.
- 1 Mavren Fein, Dusk Apostle | synergy | Chosen for its listed synergy role.
- 1 Oketra's Monument | synergy | Chosen for its listed synergy role.
- 1 Archangel Elspeth | threat | Chosen for its listed threat role.
- 1 Archon of Sun's Grace | threat | Chosen for its listed threat role.
- 1 Attended Healer | threat | Chosen for its listed threat role.
- 1 Basri's Lieutenant | threat | Chosen for its listed threat role.
- 1 Cemetery Protector | threat | Chosen for its listed threat role.
- 1 Defiler of Faith | threat | Chosen for its listed threat role.
- 1 Dragonback Lancer | threat | Chosen for its listed threat role.
- 1 Drogskol Cavalry | threat | Chosen for its listed threat role.
- 1 Emeria Angel | threat | Chosen for its listed threat role.
- 1 Gideon, Ally of Zendikar | threat | Chosen for its listed threat role.
- 1 God-Eternal Oketra | threat | Chosen for its listed threat role.
- 1 Hero of Bladehold | threat | Chosen for its listed threat role.
- 1 Ainok Strike Leader | interaction | Chosen for its listed interaction role.
- 1 Basri Ket | interaction | Chosen for its listed interaction role.
- 1 Lena, Selfless Champion | interaction | Chosen for its listed interaction role.
- 1 Rootborn Defenses | interaction | Chosen for its listed interaction role.
- 1 Spirit Bonds | interaction | Chosen for its listed interaction role.
- 1 Squad Commander | interaction | Chosen for its listed interaction role.

</details>

### 13. owned first, and the commander is not owned

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 229 names.

Commander: Karlov of the Ghost Council.

Grade: baseline, score 0.07, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $199.04 the whole deck.

**Summary:** Karlov leads a white-black lifegain deck that develops steady mana, uses repeatable lifegain synergies to build pressure, and protects its key creatures with equipment and instant-speed support. It wins by turning that incremental advantage into a threatening board of Angels and other creatures while keeping opponents’ key permanents in check with efficient removal and timely resets. The deck gives up explosive multicolor options for a consistent two-color mana base and a focused, board-centered game plan.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. Few lands enter tapped, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The lifegain payoffs and Karlov-growth cards hang together, but the list also carries a sizable equipment/protection voltron package and three board wipes that pull against the summary's "threatening board of Angels" plan, so the pieces do not all point one way.
- PLAN theme_fit=yes: It is a white-black Karlov lifegain deck at a mid-power, mostly owned-looking card level with no expensive staples forced in, matching the bracket-3 and "from my library" framing.
- PLAN useful_as_built=yes: 36 lands with heavy basics, a healthy ramp suite, ten draw pieces, and a low curve mean the deck casts its spells reliably and can close games through Karlov plus the angel and equipment threats.
- PLAN summary_honest=yes: Everything it claims is present—lifegain engines, equipment and instant-speed protection, nine pieces of efficient removal, three wipes as "resets," and a two-color mana base of mostly basics—and it does not hide the sweepers or the modest power level.
- [WARN] `not_owned`: Karlov of the Ghost Council: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.00 over 63 nonland cards
- [INFO] `mana_pass`: the builder moved 1 card of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 15 Plains | land | Basic white mana source for the deck’s early white requirements.
- 16 Swamp | land | Basic black mana source for the deck’s early black requirements.
- 1 Command Tower | land | Flexible untapped color source for the commander’s colors.
- 1 City of Brass | land | Flexible color source that supports either half of the deck.
- 1 Exotic Orchard | land | Additional flexible color fixing.
- 1 Marsh Flats | land | Fetch land that finds the deck’s basic color sources.
- 1 Fabled Passage | land | Fetch land that improves access to basic color sources.
- 1 Buster Sword | draw | Equipment-based card advantage for the draw package.
- 1 Call of the Ring | draw | Dedicated card-advantage piece.
- 1 Exemplar of Light | draw | Creature-based card advantage that fits the lifegain plan.
- 1 Idol of Oblivion | draw | Low-cost card-advantage option.
- 1 Inspiring Overseer | draw | Creature-based card advantage for the deck’s value plan.
- 1 Lembas | draw | Compact artifact card-advantage piece.
- 1 Mask of Memory | draw | Equipment-based card advantage.
- 1 Night's Whisper | draw | Efficient dedicated card-advantage spell.
- 1 Puresteel Paladin | draw | Equipment-focused card advantage.
- 1 Skullclamp | draw | Efficient equipment-based card advantage.
- 1 Wall of Omens | draw | Early defensive creature that supplies card advantage.
- 1 Bastion Protector | interaction | Protective interaction for Karlov and other key creatures.
- 1 Boromir, Warden of the Tower | interaction | Creature-based disruption for opposing plays.
- 1 Champion's Helm | interaction | Equipment protection for Karlov.
- 1 Clever Concealment | interaction | Broad protective interaction for the board.
- 1 Darksteel Plate | interaction | Durable equipment protection for a key creature.
- 1 Lightning Greaves | interaction | Low-cost protection for Karlov or a threat.
- 1 Moogles' Valor | interaction | Flexible instant-speed protective interaction.
- 1 Swiftfoot Boots | interaction | Repeatable equipment protection for important creatures.
- 1 Banishing Light | removal | Versatile permanent removal.
- 1 Bitter Triumph | removal | Efficient single-target removal.
- 1 Dispatch | removal | Low-cost targeted removal.
- 1 Fatal Push | removal | Efficient targeted removal.
- 1 Generous Gift | removal | Flexible answer to a problematic permanent.
- 1 Get Lost | removal | Versatile targeted removal.
- 1 Infernal Grasp | removal | Reliable creature removal.
- 1 Path to Exile | removal | Efficient targeted removal.
- 1 Swords to Plowshares | removal | Efficient targeted removal.
- 1 Arcane Signet | ramp | Reliable early color fixing and mana acceleration.
- 1 Commander's Sphere | ramp | Color fixing that also remains useful later.
- 1 Fellwar Stone | ramp | Low-cost mana acceleration and fixing.
- 1 Giada, Font of Hope | ramp | Creature-based mana acceleration that supports the Angel threats.
- 1 Lotho, Corrupt Shirriff | ramp | Lifegain-adjacent mana acceleration.
- 1 Relic of Legends | ramp | Color fixing and mana acceleration.
- 1 Sol Ring | ramp | Efficient colorless mana acceleration.
- 1 Sword of the Animist | ramp | Equipment-based mana development.
- 1 Thought Vessel | ramp | Low-cost colorless mana acceleration.
- 1 Wayfarer's Bauble | ramp | Early basic-land mana development.
- 1 Aerith Gainsborough | synergy | Lifegain-focused synergy creature.
- 1 Angel of Vitality | synergy | Lifegain synergy creature that supports the core plan.
- 1 Compassionate Healer | synergy | Dedicated lifegain synergy piece.
- 1 Kor Firewalker | synergy | Lifegain-oriented creature for Karlov’s plan.
- 1 Light of Promise | synergy | Lifegain payoff that helps turn incremental gains into pressure.
- 1 Night Nurse, Healer of Heroes | synergy | Lifegain-focused legendary synergy piece.
- 1 Prideful Feastling | synergy | Creature-based lifegain synergy.
- 1 Rosie Cotton of South Lane | synergy | Lifegain synergy that develops the board.
- 1 Second Breakfast | synergy | Instant-speed lifegain synergy support.
- 1 Well-Worn Spatula | synergy | Artifact-based lifegain synergy piece.
- 1 Angel of Invention | threat | White creature threat that adds pressure to the board.
- 1 Bill the Pony | threat | Creature threat that advances the board presence.
- 1 Canyon Crawler | threat | Creature threat for applying combat pressure.
- 1 Dawnhand Eulogist | threat | Creature threat that contributes to the deck’s closing pressure.
- 1 Gollum, Silent Slinker // Meager Meal | threat | Black creature threat with useful flexibility.
- 1 Lyra Dawnbringer | threat | Angel threat that suits the deck’s lifegain theme.
- 1 Minwu, White Mage | threat | White creature threat for the midgame board.
- 1 Rabaroo Troop | threat | Creature threat that adds battlefield pressure.
- 1 Rooftop Percher | threat | Creature threat for maintaining pressure after interaction.
- 1 Shattered Angel | threat | Angel threat aligned with the lifegain strategy.
- 1 Sneering Shadewriter | threat | Black creature threat that broadens the attack plan.
- 1 Victory's Herald | threat | High-impact Angel threat for closing games.
- 1 Austere Command | wipe | Flexible board-reset option.
- 1 Fumigate | wipe | Board reset that supports the lifegain plan.
- 1 Vanquish the Horde | wipe | Efficient creature-board reset.

</details>

### 14. the user delegates the commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 244 names.

Commander: Atarka, World Render.

Grade: bad, score 0.04, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $632.86 to buy, $632.86 the whole deck.

**Summary:** Atarka, World Render leads a red-green dragon deck that develops mana early, uses dragon-focused support to establish a threatening board, and refills through creature-based card advantage. The deck aims to close through sustained combat pressure from large dragons while carrying enough removal and protective interaction to force key threats through. Its tradeoff is that the strongest plays are concentrated in the midgame, so it is less focused on disrupting every early opposing setup.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. Many of the cards appear in no top list, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Every slot supports one line: dragon cost reducers and dragon-flavored ramp (Orbs, Dragonspeaker Shaman, Dragonlord's Servant, Sarkhan Fireblood) into a dense curve of big dragons, with dragon-tribal payoffs (Scourge of Valkas, Lathliss, Crucible of Fire, Terror of the Peaks) and creature-based draw feeding the same board, plus haste/protection to make Atarka's double strike land.
- PLAN theme_fit=yes: It is a red-green dragon Commander deck with a builder-chosen dragon commander (Atarka, World Render), and the power level — dragon tribal beatdown with no fast combo lines, just Taiga/Mox Jasper style staples — sits comfortably in bracket 3.
- PLAN useful_as_built=yes: 36 lands plus roughly ten ramp/cost-reduction pieces adequately support the heavy 5-7 drop curve, and there are many redundant win routes (Atarka swings, Utvara/Lathliss token engines, Scourge of Valkas and Terror of the Peaks direct damage) so the deck can be picked up and played as-is.
- PLAN summary_honest=yes: The claims of early mana development, dragon support, creature-based refill, removal and protection all map to actual cards in the list, and it openly concedes the midgame-heavy, low-early-disruption profile that the top-end curve confirms.
- [INFO] `curve_summary`: average mana value 3.27 over 63 nonland cards
- [INFO] `mana_pass`: the builder moved 1 card of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 12 Forest | land | Basic green land for the mana base.
- 14 Mountain | land | Basic red land for the mana base.
- 1 Command Tower | land | Flexible red-green land for the mana base.
- 1 Stomping Ground | land | Red-green land for the mana base.
- 1 Taiga | land | Red-green land for the mana base.
- 1 Karplusan Forest | land | Red-green land for the mana base.
- 1 Spire Garden | land | Red-green land for the mana base.
- 1 Copperline Gorge | land | Red-green land for the mana base.
- 1 Rootbound Crag | land | Red-green land for the mana base.
- 1 Rockfall Vale | land | Red-green land for the mana base.
- 1 Game Trail | land | Red-green land for the mana base.
- 1 Cragcrown Pathway // Timbercrown Pathway | land | Flexible red or green land for the mana base.
- 1 Carnelian Orb of Dragonkind | ramp | Early mana development for the dragon plan.
- 1 Dragon's Hoard | ramp | Mana development that fits the dragon theme.
- 1 Dragonstorm Globe | ramp | Mana development for expensive dragon turns.
- 1 Jade Orb of Dragonkind | ramp | Early mana development for the dragon plan.
- 1 Mox Jasper | ramp | Low-cost mana development.
- 1 Orb of Dragonkind | ramp | Early mana development for the dragon plan.
- 1 Reckless Barbarian | ramp | Low-cost mana development.
- 1 Scaled Nurturer | ramp | Early mana development within the dragon theme.
- 1 Sarkhan, Dragon Ascendant | ramp | Mana development that supports the dragon plan.
- 1 Sarkhan, Fireblood | ramp | Mana development that supports the dragon plan.
- 1 Avaricious Dragon | draw | Dragon-based card advantage.
- 1 Beast Whisperer | draw | Creature-focused card advantage.
- 1 Demand Answers | draw | Efficient card selection and card advantage.
- 1 Elemental Bond | draw | Card advantage alongside large creatures.
- 1 Faithless Looting | draw | Low-cost card selection.
- 1 Garruk's Uprising | draw | Card advantage for the large-creature plan.
- 1 Guardian Project | draw | Ongoing creature-based card advantage.
- 1 Return of the Wildspeaker | draw | Flexible card advantage for a creature-heavy deck.
- 1 Skullclamp | draw | Low-cost card advantage.
- 1 Sylvan Library | draw | Early card selection and card advantage.
- 1 Thrill of Possibility | draw | Efficient card selection.
- 1 Heroic Intervention | interaction | Protects the board from opposing interaction.
- 1 Lightning Greaves | interaction | Protects an important creature.
- 1 Legolas's Quick Reflexes | interaction | Low-cost protection for a key creature.
- 1 Mithril Coat | interaction | Protects an important creature.
- 1 Snakeskin Veil | interaction | Low-cost protection for a key creature.
- 1 Swiftfoot Boots | interaction | Protects an important creature.
- 1 Tamiyo's Safekeeping | interaction | Low-cost protection for a key permanent.
- 1 Veil of Summer | interaction | Efficient protection against opposing interaction.
- 1 Dragon's Fire | removal | Efficient spot removal.
- 1 Draconic Roar | removal | Efficient dragon-themed spot removal.
- 1 Glorybringer | removal | Dragon threat that also supplies removal.
- 1 Invasion of Tarkir // Defiant Thundermaw | removal | Dragon-themed removal that remains a relevant body.
- 1 Magmatic Hellkite | removal | Dragon-based spot removal.
- 1 Piercing Exhale | removal | Low-cost spot removal.
- 1 Scourge of Valkas | removal | Dragon-themed removal for the creature plan.
- 1 Spit Flame | removal | Dragon-themed spot removal.
- 1 Terror of the Peaks | removal | Dragon threat that also supplies removal.
- 1 Acolyte of Bahamut | synergy | Supports the high-end dragon game plan.
- 1 Crucible of Fire | synergy | Direct support for the dragon theme.
- 1 Dragon Egg | synergy | Low-cost creature that fits the dragon theme.
- 1 Dragon Hatchling | synergy | Low-cost dragon for the tribal plan.
- 1 Dragonkin Berserker | synergy | Direct support for the dragon theme.
- 1 Dragonlord's Servant | synergy | Supports deploying the deck's dragons.
- 1 Dragonspeaker Shaman | synergy | Supports deploying the deck's dragons.
- 1 Firespitter Whelp | synergy | Low-cost dragon for the tribal plan.
- 1 Sarkhan's Triumph | synergy | Finds a key dragon for the game plan.
- 1 Shivan Devastator | synergy | Flexible dragon that supports the tribal plan.
- 1 Backdraft Hellkite | threat | Midgame dragon threat.
- 1 Blast-Furnace Hellkite | threat | Dragon threat that pressures opponents.
- 1 Dragon Broodmother | threat | High-impact dragon threat.
- 1 Hellkite Charger | threat | Dragon threat for closing games.
- 1 Lathliss, Dragon Queen | threat | High-impact dragon threat.
- 1 Manaform Hellkite | threat | Efficient dragon threat.
- 1 Scourge of the Throne | threat | Dragon threat for closing games.
- 1 Stormbreath Dragon | threat | Midgame dragon threat.
- 1 Thunderbreak Regent | threat | Efficient dragon threat.
- 1 Twinflame Tyrant | threat | Dragon threat for closing games.
- 1 Utvara Hellkite | threat | High-end dragon threat.
- 1 Terror of Mount Velus | threat | High-end dragon threat.
- 1 Breath Weapon | wipe | Low-cost board cleanup.
- 1 Draconic Intervention | wipe | Flexible board cleanup.
- 1 Ryusei, the Falling Star | wipe | Dragon-based board cleanup.

</details>

### 15. delegated commander, owned first

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 229 names.

Commander: Astarion, the Decadent.

Grade: bad, score 0.01, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $65.49 the whole deck.

**Summary:** Astarion leads a black lifegain shell that develops mana, builds advantage with durable black and colorless support pieces, and turns its themed creature package into steady pressure. It closes through its larger threats while using focused removal and a few reset buttons to keep opposing boards manageable. The deck gives up multicolor lifegain staples and explosive mana for a consistent mono-black base built around Astarion's plan.

The quality model grades this deck below the precon baseline against the top lists of the format: the mana base serves one color far better than another, and that lowers the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=no: The list is a pile of generic black removal, mana rocks, and unrelated creatures with almost no lifegain triggers or payoffs to feed Astarion, so nothing ties the pieces to a single plan and the "themed creature package" is a grab-bag of unconnected bodies.
- PLAN theme_fit=no: The person asked for a lifegain deck and got a black midrange goodstuff pile with only a couple of incidental life-gaining cards and no lifegain payoffs beyond the commander.
- PLAN useful_as_built=partly: It can be cast and played on curve off 26 Swamps plus rocks, but a third of the mana base is Plains supporting essentially no white spells, and the win path is a scatter of small unrelated creatures.
- PLAN summary_honest=no: It calls the deck a lifegain shell it does not carry and describes a "consistent mono-black base" while the list runs 12 Plains alongside 26 Swamps, which is the opposite of what is claimed.
- [WARN] `not_owned`: Astarion, the Decadent: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 2.74 over 61 nonland cards
- [INFO] `mana_pass`: the builder moved 11 cards of the mana base to bring the deck inside its power level
- [INFO] `basics_added`: the list was 1 card short, so the builder added 1 basic land

<details><summary>The deck list</summary>

- 26 Swamp | land | Reliable black mana base.
- 1 Arcane Signet | ramp | Color-fixing mana acceleration.
- 1 Astral Cornucopia | ramp | Flexible mana acceleration.
- 1 Bender's Waterskin | ramp | Supports the mana development plan.
- 1 Blitzball | ramp | Supports the mana development plan.
- 1 Brass Infiniscope | ramp | Supports the mana development plan.
- 1 Chromatic Lantern | ramp | Mana acceleration and fixing.
- 1 Commander's Sphere | ramp | Reliable mana acceleration.
- 1 Deadly Dispute | ramp | Turns a resource into mana acceleration.
- 1 Fellwar Stone | ramp | Efficient mana acceleration.
- 1 Sol Ring | ramp | Efficient early mana acceleration.
- 1 Call of the Ring | draw | Ongoing card advantage.
- 1 Cirith Ungol Patrol | draw | Adds to the deck's card-advantage package.
- 1 Gollum, Riddle Master | draw | Adds card advantage on a thematic creature.
- 1 Grave Venerations | draw | Ongoing card advantage.
- 1 Idol of Oblivion | draw | Efficient card-advantage piece.
- 1 Kingpin's Enforcers | draw | Adds card advantage on a creature.
- 1 Nasty End | draw | Efficient card-advantage spell.
- 1 Night's Whisper | draw | Efficient card-advantage spell.
- 1 Obsessive Pursuit | draw | Ongoing card advantage.
- 1 Skullclamp | draw | Efficient equipment-based card advantage.
- 1 Stone of Erech | draw | Colorless card-advantage option.
- 1 Champion's Helm | interaction | Protects a key legendary permanent.
- 1 Darksteel Plate | interaction | Protects an important creature.
- 1 Lightning Greaves | interaction | Efficient creature protection.
- 1 Orcish Medicine | interaction | Flexible black interaction.
- 1 Recuperate | interaction | Supports the deck's protection package.
- 1 Swiftfoot Boots | interaction | Protects an important creature.
- 1 Bitter Triumph | removal | Flexible black removal.
- 1 Blowfly Infestation | removal | Provides repeatable removal pressure.
- 1 Contagion Clasp | removal | Colorless removal option.
- 1 Dismember | removal | Efficient creature removal.
- 1 Fatal Push | removal | Efficient early creature removal.
- 1 Infernal Grasp | removal | Reliable creature removal.
- 1 Liminal Hold | removal | Versatile black removal.
- 1 Adventurous Eater // Have a Bite | synergy | Supports the lifegain-focused synergy plan.
- 1 Bob, Reluctant HYDRA Agent | synergy | Contributes to the deck's themed synergy package.
- 1 Doom Reigns Supreme | synergy | Supports the deck's central synergy plan.
- 1 Elixir | synergy | Colorless support for the lifegain synergy plan.
- 1 Graveyard Trespasser // Graveyard Glutton | synergy | Adds a resilient synergy creature.
- 1 HYDRA Infiltration | synergy | Supports the deck's themed synergy package.
- 1 Melancholic Poet | synergy | Adds to the lifegain-supporting creature base.
- 1 Morlun, Devourer of Spiders | synergy | Supports the deck's central synergy plan.
- 1 Pull from the Grave | synergy | Recovers an important synergy piece.
- 1 Ravening Warg | synergy | Adds to the deck's creature synergy package.
- 1 Scarblade's Malice | synergy | Supports the deck's black synergy package.
- 1 Send in the Pest | synergy | Supports the deck's central synergy plan.
- 1 The Darkness Crystal | synergy | Colorless support for the synergy plan.
- 1 Well-Worn Spatula | synergy | Equipment support for the synergy package.
- 1 Whiplash, Vengeful Engineer | synergy | Adds to the deck's themed synergy package.
- 1 Yellowjacket, Heartless Marauder | synergy | Adds to the deck's black synergy package.
- 1 Canyon Crawler | threat | Creature threat that helps pressure opponents.
- 1 Dawnhand Eulogist | threat | Creature threat for the deck's closing plan.
- 1 Foggy Swamp Hunters | threat | Creature threat that fits the black creature base.
- 1 Gollum, Silent Slinker // Meager Meal | threat | Thematic creature threat.
- 1 Rabaroo Troop | threat | Creature threat for applying pressure.
- 1 Reaping Willow | threat | Creature threat for the late game.
- 1 Rooftop Percher | threat | Creature threat that broadens the board presence.
- 1 Sneering Shadewriter | threat | Black creature threat for the closing plan.
- 1 Archfiend of Ifnir | wipe | Creature-based board-clearing effect.
- 1 M.O.D.O.K. | wipe | Additional board-clearing option.
- 1 Withering Curse | wipe | Black board-clearing option.
- 12 Plains | land | the builder added this basic land to reach the deck size

</details>

### 16. a tight budget, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 229 names.

Commander: Denethor, Ruling Steward.

Grade: baseline, score 0.30, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $9.56 to buy, $37.77 the whole deck.

**Summary:** This is a budget-minded Orzhov aristocrats deck built to establish a board of creatures and synergy pieces, use sacrifice-compatible resources to keep cards and mana flowing, and turn its themed threats and payoffs into a steady path to victory. It has a broad package of creature removal, protective interaction, and board wipes to recover tempo when opponents get ahead. The deck gives up premium mana, expensive staples, and explosive combo finishes in exchange for a consistent library-first build with a deliberate, creature-focused game plan.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: A real aristocrats core exists (Blood Artist, Bastion of Remembrance, Falkenrath Noble, Vindictive Vampire, Elas il-Kor plus sac outlets like Bartolomé, Vampire Gourmand, Old Flitterfang, Shilgengar), but it is diluted by a sizable block of off-plan filler creatures (Arcade Cabinet, Joo Dee, Namazu Trader, Phantom Train, Heirloom Auntie, Beetle-Headed Merchants) and thin token/fodder generation, while three board wipes and equipment-style protection tug against the go-wide sacrifice plan.
- PLAN theme_fit=yes: It is an Orzhov aristocrats Commander deck under a tight budget with a low-power, bracket-2 feel and a collection-first card pool, matching everything the request named.
- PLAN useful_as_built=yes: 38 lands with mostly basics plus Sol Ring/Signet-level ramp, a sane curve, seven removal spells and multiple drain payoffs mean it can be shuffled up and played without fixing anything.
- PLAN summary_honest=partly: Nothing is outright false — the removal, protection, and wipes are all present and it honestly flags the lack of premium mana and combo finishes — but it hedges into vagueness ("sacrifice-compatible resources," "themed threats and payoffs") rather than admitting how many slots are unfocused filler, and calling seven removal spells a "broad package" is generous.
- [WARN] `not_owned`: Bastion of Remembrance: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Blood Artist: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Elas il-Kor, Sadistic Pilgrim: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Baron Bertram Graywater: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bartolomé del Presidio: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Falkenrath Noble: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Lord Skitter's Butcher: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Old Flitterfang: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Shilgengar, Sire of Famine: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vampire Gourmand: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vindictive Vampire: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 2.97 over 62 nonland cards

<details><summary>The deck list</summary>

- 15 Plains | land | Basic white source that keeps the mana base reliable.
- 15 Swamp | land | Basic black source that keeps the mana base reliable.
- 1 Command Tower | land | Flexible color source for the two-color mana base.
- 1 Exotic Orchard | land | Flexible color source for the two-color mana base.
- 1 City of Brass | land | Reliable multicolor land already in the library.
- 1 Grand Coliseum | land | Additional multicolor source for the deck's colored spells.
- 1 Evolving Wilds | land | Basic-land access that supports color consistency.
- 1 Fabled Passage | land | Basic-land access that supports color consistency.
- 1 Terramorphic Expanse | land | Basic-land access that supports color consistency.
- 1 Ahriman | draw | Library-owned draw option for maintaining resources.
- 1 Beetle-Headed Merchants | draw | Low-cost library-owned draw support.
- 1 Buzzard-Wasp Colony | draw | Library-owned draw piece that adds to the creature count.
- 1 Circle of Power | draw | Inexpensive draw support from the library.
- 1 Cirith Ungol Patrol | draw | Creature-based draw support from the library.
- 1 Exemplar of Light | draw | Library-owned draw creature for sustained resources.
- 1 Foot Chopper | draw | Artifact draw support from the library.
- 1 Grave Venerations | draw | Draw support that fits the deck's graveyard-oriented plan.
- 1 Inspiring Overseer | draw | Efficient creature-based draw support.
- 1 June, Bounty Hunter | draw | Low-cost draw support already in the library.
- 1 Kingpin's Enforcers | draw | Additional creature-based draw support.
- 1 Arcane Signet | ramp | Efficient fixing and acceleration already in the library.
- 1 Astral Cornucopia | ramp | Artifact acceleration from the library.
- 1 Blitzball | ramp | Low-cost artifact ramp already available.
- 1 Chromatic Lantern | ramp | Library-owned mana fixing and acceleration.
- 1 Commander's Sphere | ramp | Reliable artifact ramp for the mana base.
- 1 Deadly Dispute | ramp | Sacrifice-compatible ramp option from the library.
- 1 Fellwar Stone | ramp | Efficient library-owned mana acceleration.
- 1 Inherited Envelope | ramp | Additional artifact ramp from the library.
- 1 Sol Ring | ramp | Fast early acceleration already in the library.
- 1 Thought Vessel | ramp | Library-owned artifact ramp for developing the board.
- 1 Bitter Triumph | removal | Flexible single-target removal from the library.
- 1 Claim the Precious | removal | Efficient black removal already in the library.
- 1 Crib Swap | removal | Creature removal that is easy on the mana base.
- 1 Deadly Precision | removal | Low-cost removal from the library.
- 1 Destroy Evil | removal | Flexible removal for problematic permanents.
- 1 Generous Gift | removal | Broad permanent removal already in the library.
- 1 Infernal Grasp | removal | Reliable creature removal from the library.
- 1 Airbending Lesson | interaction | Low-cost interaction already in the library.
- 1 Bastion Protector | interaction | Creature-based protection for the deck's key permanents.
- 1 Frontline Medic | interaction | Creature-based interaction that supports a board-focused plan.
- 1 Gift of Immortality | interaction | Protection for an important sacrifice-plan creature.
- 1 Swiftfoot Boots | interaction | Efficient protection already in the library.
- 1 Take Up the Shield | interaction | Low-cost protective interaction from the library.
- 1 Black Sun's Zenith | wipe | Flexible board wipe already in the library.
- 1 Dusk // Dawn | wipe | Board-control option that supports creature-centric games.
- 1 Martial Coup | wipe | Board wipe that also supplies a threatening board presence.
- 1 Al Bhed Salvagers | synergy | Library-owned synergy piece for the aristocrats shell.
- 1 Arcade Cabinet | synergy | Library-owned synergy permanent for the deck's core plan.
- 1 Gollum the Abandoned | synergy | Creature synergy piece already in the library.
- 1 Gollum, Patient Plotter | synergy | Library-owned creature synergy for the sacrifice-focused plan.
- 1 Gríma Wormtongue | synergy | Low-cost synergy creature from the library.
- 1 Heirloom Auntie | synergy | Creature synergy support already in the library.
- 1 Joo Dee, One of Many | synergy | Library-owned synergy piece for building the board.
- 1 Nimble Hobbit | synergy | Low-cost creature synergy from the library.
- 1 Phantom Train | synergy | Artifact synergy piece already in the library.
- 1 Rhovanion Rampager | synergy | Creature synergy support from the library.
- 1 Bastion of Remembrance | synergy | Inexpensive aristocrats payoff for the sacrifice plan.
- 1 Blood Artist | synergy | Central aristocrats payoff that rewards the deck's core plan.
- 1 Elas il-Kor, Sadistic Pilgrim | synergy | Aristocrats-focused creature that reinforces the deck's plan.
- 1 Bill the Pony | threat | Library-owned creature threat for pressuring opponents.
- 1 Hei Bai, Spirit of Balance | threat | Low-cost creature threat already in the library.
- 1 Namazu Trader | threat | Additional library-owned creature threat.
- 1 Vengeful Villagers | threat | Creature threat that helps maintain battlefield pressure.
- 1 Baron Bertram Graywater | threat | Inexpensive creature threat that fits the deck's theme.
- 1 Bartolomé del Presidio | threat | Low-cost themed creature threat for the aristocrats shell.
- 1 Falkenrath Noble | threat | Creature threat that supports the sacrifice-focused strategy.
- 1 Lord Skitter's Butcher | threat | Themed creature threat at a very low cost.
- 1 Old Flitterfang | threat | Inexpensive creature threat for the deck's board plan.
- 1 Shilgengar, Sire of Famine | threat | Themed creature threat that complements the aristocrats shell.
- 1 Vampire Gourmand | threat | Low-cost creature threat that fits the deck's theme.
- 1 Vindictive Vampire | threat | Creature threat that supports the sacrifice-focused plan.

</details>

### 17. upgrade a precon, owned first

Format: Commander. Theme: goblins. Pool: owned_first. Shortlist: 196 names.

Commander: Zada, Hedron Grinder.

Grade: typical, score 0.57, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: yes, for deck_size, precon_share, profile_off_band, profile_off_band, profile_off_band, profile_off_band, profile_off_band, profile_off_band, profile_off_band. Block findings: 0.

Cost: $80.64 to buy, $256.45 the whole deck.

**Summary:** A mono-red Goblin swarm deck built to turn cheap creatures, token makers, and targeted spells into overwhelming combat turns. It develops a broad battlefield, protects its centerpiece, converts creature deaths and arrivals into damage, and can pivot from pressure to explosive spell-fueled finishing turns.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: The cards all serve one plan: flood the board with cheap goblins and tokens, then use Zada-friendly cantrip pump spells, rituals, and death/ETB damage payoffs (Impact Tremors, Boggart Shenanigans, Goblin Bombardment) to convert the wide board into a lethal turn.
- PLAN theme_fit=yes: It is a mono-red goblin deck led by Zada with a genuine storm subtheme (rituals, one-target-copies-to-all cantrips, Grapeshot/Empty the Warrens), upgraded with staples but staying clear of infinite combos or fast mana beyond Sol Ring, which suits bracket 3.
- PLAN useful_as_built=yes: 37 lands with utility lands that still produce red, a very low curve full of two-drops and cantrips, plenty of card draw (Skullclamp, War Room, Idol of Oblivion) and multiple distinct win routes mean it can be picked up and played as-is.
- PLAN summary_honest=yes: Every claim — token swarm, commander protection via Helm/Plate/Greaves/Boots, death-and-arrival damage engines, and ritual-fueled storm finishers like Grapeshot, Empty the Warrens, and Haze of Rage — is backed by cards actually in the list.
- [INFO] `curve_summary`: average mana value 2.55 over 62 nonland cards
- [WARN] `profile_off_band`: the count of nonbasic lands that make only colorless mana is 5, and bracket 3 wants 4 at most (Fountainport, Goblin Burrows, Kher Keep, Reliquary Tower, War Room)
- [WARN] `profile_off_band`: the removal count is 13, and bracket 3 wants 6 to 12
- [INFO] `bracket_cut`: to hold bracket 3, the builder cut 1 card: Storm-Kiln Artist (the combo Storm-Kiln Artist + Haze of Rage), and added 1 basic land

<details><summary>The deck list</summary>

- 1 Abrade | removal | Flexible answer to creatures or artifacts.
- 1 Ancestors' Aid | synergy | Targeted spell that supports Zada-style combat turns.
- 1 Ancestral Anger | draw | Cheap targeted cantrip for the commander and creature swarm.
- 1 Arcane Signet | ramp | Reliable early mana acceleration.
- 1 Arena of Glory | land | Precon utility land.
- 1 Battle Hymn | ramp | Converts a wide Goblin board into explosive mana.
- 1 Battle-Scarred Goblin | removal | Goblin-bodied creature interaction.
- 1 Blasphemous Act | wipe | Efficient emergency creature reset.
- 1 Boggart Shenanigans | removal | Turns Goblin deaths into direct damage.
- 1 Brightstone Ritual | ramp | Burst mana fueled by Goblins.
- 1 Castle Embereth | land | Precon land that boosts a wide attack.
- 1 Champion's Helm | interaction | Protects and enhances the commander.
- 1 Chaos Warp | removal | Versatile answer to problematic permanents.
- 1 Conspicuous Snoop | synergy | Goblin synergy and card-access engine.
- 1 Crimson Wisps | draw | Targeted cantrip that can give the team haste through Zada.
- 1 Daring Discovery | synergy | Precon spell that supports the deck's proactive plan.
- 1 Darksteel Plate | interaction | Durable commander protection.
- 1 Den of the Bugbear | land | Precon creature-land utility.
- 1 Dragon Fodder | synergy | Efficiently builds a Goblin board.
- 1 Dwarven Mine | land | Precon land that contributes a Goblin token.
- 1 Empty the Warrens | synergy | Creates a large Goblin force after a spell-heavy turn.
- 1 Expedite | draw | Cheap targeted cantrip and haste effect.
- 1 Faithless Looting | draw | Efficient card selection and graveyard setup.
- 1 Fists of Flame | draw | Targeted card draw and a powerful Zada combat payoff.
- 1 Forgotten Cave | land | Precon land with late-game cycling utility.
- 1 Fountainport | land | Precon utility land.
- 1 Frontline Heroism | synergy | Supports the deck's aggressive creature plan.
- 1 Gempalm Incinerator | removal | Goblin-scaled creature removal with cycling value.
- 1 General Kreat, the Boltbringer | synergy | Goblin payoff that reinforces the creature theme.
- 1 Glimpse the Impossible | ramp | Provides mana and impulse-style resources.
- 1 Goblin Bombardment | removal | Sacrifice outlet and repeatable reach.
- 1 Goblin Burrows | land | Precon land that improves Goblin combat.
- 1 Goblin Bushwhacker | synergy | Token-board haste and power payoff.
- 1 Goblin Chieftain | synergy | Lord effect that makes Goblin swarms immediately threatening.
- 1 Goblin Dark-Dwellers | threat | Threat that reuses an impactful instant or sorcery.
- 1 Goblin Lackey | synergy | Accelerates Goblins onto the battlefield through combat.
- 1 Goblin Matron | synergy | Finds the Goblin needed for the current board state.
- 1 Goblin Negotiation | removal | Thematic Goblin removal spell.
- 1 Goblin Trashmaster | removal | Goblin lord with artifact removal utility.
- 1 Goblin Warchief | synergy | Cost reduction and haste for Goblin turns.
- 1 Grapeshot | removal | Storm-based removal and finishing reach.
- 1 Haze of Rage | wincon | Storm combat finisher for a wide board.
- 1 Hidden Volcano | land | Precon land with creature utility.
- 1 Idol of Oblivion | draw | Repeatable card draw alongside token production.
- 1 Impact Tremors | wincon | Converts every Goblin token into damage.
- 1 Kher Keep | land | Precon token-producing utility land.
- 1 Krenko's Command | synergy | Efficient Goblin token production.
- 1 Krenko, Mob Boss | threat | Major Goblin token engine and primary threat.
- 1 Lightning Bolt | removal | Low-cost, instant-speed creature removal or reach.
- 1 Lightning Greaves | interaction | Protects the commander while enabling immediate attacks.
- 1 Mana Geyser | ramp | Explosive mana for high-spell storm turns.
- 1 Mogg War Marshal | synergy | Efficient multi-body Goblin producer.
- 24 Mountain | land | Reliable untapped red mana base.
- 1 Pashalik Mons | removal | Goblin payoff that turns sacrifices and losses into damage.
- 1 Path of Ancestry | land | Precon tribal fixing and card-selection land.
- 1 Quest for the Goblin Lord | synergy | Tribal payoff for repeated Goblin deployment.
- 1 Reliquary Tower | land | Precon utility land for keeping a full hand.
- 1 Roaming Throne | synergy | Doubles key Goblin triggered abilities.
- 1 Ruby Medallion | ramp | Red spell cost reduction supports spell chaining.
- 1 Sazacap's Brew | draw | Targeted draw spell that works well with Zada.
- 1 Searslicer Goblin | synergy | Goblin creature that advances the tribal attack plan.
- 1 Seething Song | ramp | Ritual mana for explosive turns.
- 1 Shinka, the Bloodsoaked Keep | land | Precon legendary land with combat utility.
- 1 Siege-Gang Commander | removal | Produces Goblins and converts them into damage.
- 1 Siege-Gang Lieutenant | removal | Goblin token payoff with removal utility.
- 1 Skirk Prospector | ramp | Goblin sacrifice outlet that produces mana.
- 1 Skullclamp | draw | Turns expendable Goblin tokens into cards.
- 1 Sol Ring | ramp | Efficient colorless acceleration.
- 1 Spreading Insurrection | wincon | High-impact aggressive spell for a developed board.
- 1 Swiftfoot Boots | interaction | Commander protection with haste.
- 1 Throne of Eldraine | ramp | Precon mana rock with late-game utility.
- 1 Vandalblast | wipe | Artifact sweep that can be cast efficiently early.
- 1 War Room | land | Precon land that provides repeatable card draw.
- 1 Warren Torchmaster | synergy | Goblin creature that supports the tribal game plan.
- 1 Wild Ride | synergy | Aggressive spell that supports combat-focused turns.
- 1 Witch's Mark | draw | Targeted filtering spell with haste utility.

</details>

### 18. upgrade a precon, any card

Format: Commander. Theme: turtles. Pool: any_card. Shortlist: 310 names.

Commander: Heroes in a Half Shell.

Grade: baseline, score 0.08, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $92.64 to buy, $92.64 the whole deck.

**Summary:** This remains a Turtle Power deck built around its Turtle-themed characters, artifacts, and enchantments, with its creatures supplying the main pressure while a few broad reset and removal cards help clear the way. Cowabunga!, Turtle Power!, and Turtles in Time deepen that theme without pulling the list toward a different strategy. The deck gives up the tight consistency and dense interaction of a heavily rebuilt list in exchange for preserving the precon’s character-driven, creature-forward play pattern.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. Many lands enter tapped, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: The list is a unified five-color creature/turtle-character deck: the named TMNT characters, the Cowabunga!/Turtle Power!/High Score synergy pieces, and the ramp/fixing all support the same go-wide, creature-forward plan, with a small removal and wipe package as support rather than a competing angle.
- PLAN theme_fit=partly: The colors, commander, and Turtle theme are exactly what was asked, but the request was for a bracket 3 upgrade and the deck is only lightly touched (Sol Ring, Arcane Signet, Cultivate, Assassin's Trophy, a few hydras and wipes) while the summary itself concedes it declines to tighten consistency or interaction.
- PLAN useful_as_built=yes: About 38 lands with Command Tower, City of Brass, Grand Coliseum, Path of Ancestry, thriving lands, fetch-style fixing plus Chromatic Lantern and Arcane Signet make five colors castable, and there are plenty of creature threats to actually close a game, even if the heavy tapland count slows it down.
- PLAN summary_honest=partly: The summary is candid that this stays close to the precon's play pattern rather than overclaiming a rebuilt engine, but the appended quality notes are confusing to the point of misleading (asserting that many lands entering tapped 'raises the grade'), which misrepresents a real weakness of the mana base as a positive.
- [INFO] `curve_summary`: average mana value 3.43 over 60 nonland cards
- [WARN] `profile_off_band`: the land count is 39, and bracket 3 wants 34 to 38
- [WARN] `profile_off_band`: the ramp count is 5, and bracket 3 wants 8 to 13
- [WARN] `profile_off_band`: the draw count is 2, and bracket 3 wants 8 to 14
- [WARN] `profile_off_band`: the removal count is 2, and bracket 3 wants 6 to 12
- [WARN] `profile_off_band`: the interaction count is 0, and bracket 3 wants 4 to 12

<details><summary>The deck list</summary>

- 1 Acidic Slime | other | Retains a precon creature piece.
- 1 April O'Neil, Live on the Scene | draw | Retains the precon's draw piece.
- 1 Arcade Cabinet | synergy | Retains a precon theme piece.
- 1 Arcane Signet | ramp | Retains the precon's ramp piece.
- 1 Ash Barrens | land | Retains a precon land.
- 1 Assassin's Trophy | removal | Retains the precon's removal piece.
- 1 Baxter, Fly in the Ointment | synergy | Retains a precon character card.
- 1 Bebop, Skull & Crossbones | threat | Retains a precon character card.
- 1 Big Apple, 3 a.m. | land | Retains a precon land.
- 1 Big Mother Mouser | synergy | Retains a precon artifact creature.
- 1 Biogenic Ooze | threat | Retains a precon creature threat.
- 1 Blasphemous Act | wipe | Retains the precon's wipe.
- 1 Casey Jones, Back Alley Brute | threat | Retains a precon character card.
- 1 Chromatic Lantern | ramp | Retains a precon mana piece.
- 1 Cinder Glade | land | Retains a precon land.
- 1 City of Brass | land | Retains a precon land.
- 1 Coin of Mastery | other | Retains a precon artifact piece.
- 1 Command Tower | land | Retains a precon land.
- 1 Continue? | other | Retains a precon spell.
- 1 Corpsejack Menace | synergy | Retains a precon creature piece.
- 1 Cowabunga! | synergy | Adds a Turtle-themed synergy card.
- 1 Cultivate | ramp | Retains a precon mana piece.
- 1 Dimension X Pizzasaur | synergy | Retains a precon artifact creature.
- 1 Donatello, the Brains | synergy | Retains a Turtle-themed precon card.
- 1 Double Jump // Flying Kick | other | Retains a precon spell.
- 1 Dragonskull Summit | land | Retains a precon land.
- 1 Electric Seaweed | other | Retains a precon creature.
- 1 Endless Foot Assault | synergy | Retains a precon enchantment.
- 1 Escape Tunnel | land | Retains a precon land.
- 1 Everything Pizza | synergy | Retains a precon artifact piece.
- 1 Evolving Wilds | land | Retains a precon land.
- 1 Exploding Barrel | other | Retains a precon artifact piece.
- 1 Fast Forward | other | Retains a precon spell.
- 4 Forest | land | Keeps green basic-land coverage in the mana base.
- 1 Foot Chopper | other | Retains a precon Equipment.
- 1 Game Over | other | Retains a precon spell.
- 1 Grand Coliseum | land | Retains a precon land.
- 1 Harmonize | draw | Retains a precon spell.
- 1 Here Comes a New Hero! | other | Retains a precon spell.
- 1 Hidden Hideout | land | Retains a precon land.
- 1 High Score | synergy | Retains a precon enchantment.
- 1 Irma, Part-Time Mutant | synergy | Retains a precon character card.
- 4 Island | land | Keeps blue basic-land coverage in the mana base.
- 1 Krang, the All-Powerful | threat | Retains a precon artifact creature.
- 1 Leatherhead, Iron Gator | threat | Retains a precon character card.
- 1 Lessons from Life | other | Retains a precon spell.
- 1 Level Up | synergy | Retains a precon Aura.
- 1 Lita, Little Orphan Amphibian | synergy | Retains a Turtle-themed precon card.
- 1 Michelangelo, the Heart | synergy | Retains a Turtle-themed precon card.
- 1 Mole Module | other | Retains a precon Vehicle.
- 1 Mona Lisa, Science Geek | synergy | Retains a precon character card.
- 2 Mountain | land | Keeps red basic-land coverage in the mana base.
- 1 Ninja Pizza | synergy | Retains a precon enchantment.
- 1 Path of Ancestry | land | Retains a precon land.
- 3 Plains | land | Keeps white basic-land coverage in the mana base.
- 1 Rain-Slicked Copse | land | Retains a precon land.
- 1 Raphael, the Muscle | threat | Retains a Turtle-themed precon threat.
- 1 Rat King, Pale Piper | threat | Retains a precon character card.
- 1 Ray Fillet, Wave Warrior | synergy | Retains a precon character card.
- 1 Roadkill Rodney | other | Retains a precon artifact creature.
- 1 Rocksteady, Mutant Marauder | threat | Retains a precon character card.
- 1 Rootbound Crag | land | Retains a precon land.
- 1 Shellshock | other | Retains a precon spell.
- 1 Smoldering Marsh | land | Retains a precon land.
- 1 Sodden Verdure | land | Retains a precon land.
- 1 Sol Ring | ramp | Retains the precon's ramp piece.
- 1 Special Move | other | Retains a precon spell.
- 1 Steelbane Hydra | removal | Retains a precon creature.
- 1 Sunken Hollow | land | Retains a precon land.
- 1 Super Combo | other | Retains a precon spell.
- 3 Swamp | land | Keeps black basic-land coverage in the mana base.
- 1 Swift Demise | other | Retains a precon spell.
- 1 Tempestra, Dame of Games | synergy | Retains a precon character card.
- 1 Thriving Grove | land | Retains a precon land.
- 1 Thriving Isle | land | Retains a precon land.
- 1 Thriving Moor | land | Retains a precon land.
- 1 Together Forever | synergy | Retains a precon enchantment.
- 1 Tokka & Rahzar, Unsupervised | ramp | Retains the precon's Turtle-themed ramp piece.
- 1 Turtle Lair | land | Retains a precon land.
- 1 Turtle Power! | synergy | Adds a Turtle-focused synergy centerpiece.
- 1 Turtles in Time | wipe | Adds a Turtle-themed board reset.
- 1 Undergrowth Stadium | land | Retains a precon land.
- 1 Vanquish the Horde | wipe | Retains a precon spell.
- 1 Vernal Fen | land | Retains a precon land.
- 1 Vibrant Cityscape | land | Retains a precon land.
- 1 Vigor | synergy | Retains a precon creature.
- 1 Voracious Hydra | threat | Retains a precon creature threat.
- 1 Wave Goodbye | other | Retains a precon spell.

</details>

### 19. the Hobbit family, two colours

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 174 names.

Commander: Thranduil, the Elvenking.

Grade: bad, score 0.02, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $499.26 to buy, $499.26 the whole deck.

**Summary:** This is a Sultai Elf-centered Thranduil deck that develops its mana with a dense creature-and-artifact ramp package, keeps cards flowing, and uses its Elf core alongside a broad creature suite to establish pressure. It wins through sustained board presence backed by removal, interaction, and several reset buttons. The tradeoff is that the deck leans on creatures and a straightforward mana base rather than premium fast mana or highly specialized combo lines.

The quality model grades this deck below the precon baseline against the top lists of the format: the black sources cover 16.5 of the 23 its spells need.

- PLAN plan_coherent=partly: There is a recognizable core of Elf ramp, card draw, and removal, but a large slice of the list is generic Middle-earth creatures (wargs, bears, bees, spiders) with no shared payoff, so the Elf-tribal plan the summary names is diluted by goodstuff bodies.
- PLAN theme_fit=partly: The commander, colors, format, and most of the card pool match the requested Hobbit-set build, but several inclusions (The One Ring, Orcish Bowmasters, Delighted Halfling, Elvish Mystic/Archdruid, Arcane Signet, Mox Amber, Languish, Wood Elves) come from outside that set and push past the stated set restriction.
- PLAN useful_as_built=partly: The curve, removal, and creature count are fine and there are plenty of ways to close a game, but a 36-land mana base of nothing but basics in three colors will stumble on the black-heavy and double-costed spells, as the quality note itself flags.
- PLAN summary_honest=yes: placeholder
- [INFO] `curve_summary`: average mana value 3.17 over 63 nonland cards
- [WARN] `profile_off_band`: the worst color's share of the sources it needs is 0.72, and bracket 3 wants 0.85 or more (sources of requirement: U 14.5 of 19, B 16.5 of 23, G 16.25 of 22)
- [INFO] `mana_pass`: the builder moved 4 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 14 Forest | land | Basic green land source.
- 10 Island | land | Basic blue land source.
- 12 Swamp | land | Basic black land source.
- 1 Arcane Signet | ramp | Listed ramp artifact for early mana.
- 1 Delighted Halfling | ramp | Listed ramp creature.
- 1 Elven Chorus | ramp | Listed ramp enchantment.
- 1 Elvish Archdruid | ramp | Listed Elf ramp creature.
- 1 Elvish Mystic | ramp | Listed Elf ramp creature.
- 1 Thranduil the Strategist | ramp | Listed ramp Elf legend.
- 1 Thranduil's Company | ramp | Listed Elf ramp creature.
- 1 Wayfarer's Bauble | ramp | Listed ramp artifact.
- 1 Wood Elves | ramp | Listed Elf ramp creature.
- 1 Woodland Weavemaster | ramp | Listed Elf ramp creature.
- 1 Beorn the Fierce | draw | Listed draw creature.
- 1 Elvish Visionary | draw | Listed Elf draw creature.
- 1 Fateful Discovery | draw | Listed draw enchantment.
- 1 Gandalf, Wandering Wizard | draw | Listed draw legend.
- 1 Hithlain Knots | draw | Listed draw instant.
- 1 Ithilien Kingfisher | draw | Listed draw creature.
- 1 Lórien Revealed | draw | Listed draw sorcery.
- 1 Plunder the Trollshaws | draw | Listed draw instant.
- 1 Reverent Howl | draw | Listed draw instant.
- 1 Thrór's Map | draw | Listed draw artifact.
- 1 Bilbo's Ring | interaction | Listed interaction Equipment.
- 1 Confusticate and Bebother | interaction | Listed interaction instant.
- 1 Dwarven Mattock | interaction | Listed interaction Equipment.
- 1 Elrond, Moon-Reader | interaction | Listed interaction Elf legend.
- 1 Mithril Coat | interaction | Listed interaction Equipment.
- 1 Old Fat Spider Can't See Me | interaction | Listed interaction Saga.
- 1 Stern Scolding | interaction | Listed interaction instant.
- 1 The One Ring | interaction | Listed interaction artifact.
- 1 Bilbo's Deadly Slice | removal | Listed removal instant.
- 1 Bitter Downfall | removal | Listed removal instant.
- 1 Enchanted River's Grasp | removal | Listed removal Aura.
- 1 Merciless Executioner | removal | Listed removal creature.
- 1 Orcish Bowmasters | removal | Listed removal creature.
- 1 Uneasy Partings | removal | Listed removal instant.
- 1 Witch-king of Angmar | removal | Listed removal legend.
- 1 Witch-king, Bringer of Ruin | removal | Listed removal legend.
- 1 Gnashing of Teeth | wipe | Listed wipe sorcery.
- 1 Languish | wipe | Listed wipe sorcery.
- 1 Raise the Palisade | wipe | Listed wipe sorcery.
- 1 Arwen, Weaver of Hope | synergy | Elf creature supporting the deck's Elf theme.
- 1 Boughside Wanderers | synergy | Elf creature supporting the deck's Elf theme.
- 1 Cantankerous Keepers | synergy | Elf creature supporting the deck's Elf theme.
- 1 Celeborn the Wise | synergy | Elf legend supporting the deck's Elf theme.
- 1 Elven Raft-Steerer | synergy | Elf creature supporting the deck's Elf theme.
- 1 Elvenking's Harper | synergy | Elf creature supporting the deck's Elf theme.
- 1 Galadhrim Guide | synergy | Elf creature supporting the deck's Elf theme.
- 1 Galion, Elvenking's Butler | synergy | Elf legend supporting the deck's Elf theme.
- 1 Grey Havens Navigator | synergy | Elf creature supporting the deck's Elf theme.
- 1 Supper for Spiders | synergy | Listed synergy instant.
- 1 Attercop | threat | Creature threat for board presence.
- 1 Bejeweled Warg | threat | Creature threat for board presence.
- 1 Chief Warg's Company | threat | Creature threat for board presence.
- 1 Desolation Prowler | threat | Creature threat for board presence.
- 1 Dreaded Bat-Cloud | threat | Creature threat for board presence.
- 1 Duskwatch Hunter | threat | Creature threat for board presence.
- 1 Great Fierce Bee | threat | Creature threat for board presence.
- 1 Head of the Hunt | threat | Creature threat for board presence.
- 1 Large Bear | threat | Creature threat for board presence.
- 1 Mirkwood Elk | threat | Creature threat for board presence.
- 1 Nighthowl Pursuer | threat | Creature threat for board presence.
- 1 Ravening Warg | threat | Creature threat for board presence.
- 1 Mox Amber | ramp | the mana pass added it to bring the mana base inside the power level
- 1 Giant's Boulder | removal | the mana pass added it to bring the mana base inside the power level

</details>

### 20. the Hobbit family, a delegated commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 77 names.

Commander: Smaug the Magnificent.

Grade: baseline, score 0.06, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $484.40 to buy, $484.40 the whole deck.

**Summary:** This deck uses a heavy red mana base and a broad ramp package to bring Smaug and other large creatures onto the table quickly, then presses the game through Dragon, Giant, Troll, Orc, Goblin, and Dwarf bodies backed by a suite of legendary equipment. It can clear troublesome permanents and reset crowded boards before rebuilding with its creature density and card-flow artifacts. The deck gives up some consistency for a strongly Hobbit-themed collection of creatures, equipment, and dramatic top-end plays.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The list hangs together loosely as a mono-red ramp-into-fatties/equipment pile, but the equipment package, the goblin/dwarf small bodies, and the big top-end don't reinforce one another, and there is no dragon payoff to tie Smaug's presence to the rest of the cards.
- PLAN theme_fit=partly: The commander choice, mono-red color, and bracket-3 power level match the request, but the deck is a Hobbit-flavored dwarves/goblins/giants pile with barely any actual Dragons, and several inclusions (Guttersnipe, Inferno Titan, Arcane Signet, Mox Amber, The One Ring) fall outside the named set.
- PLAN useful_as_built=yes: 36 Mountains plus roughly a dozen ramp pieces supports a mono-red curve, and the deck has ample removal, sweepers, card draw, and beaters to actually close games at a casual table.
- PLAN summary_honest=partly: It is candid about consistency loss and correctly names the equipment, ramp, and wipe elements, but claiming the game is pressed through "Dragon" bodies overstates a deck whose only dragons are the Smaug cards themselves.
- [INFO] `curve_summary`: average mana value 3.13 over 63 nonland cards
- [WARN] `profile_off_band`: the draw count is 6, and bracket 3 wants 8 to 14
- [INFO] `cards_trimmed`: the list was 1 card over, so the builder cut this card: Cavern-Hoard Dragon

<details><summary>The deck list</summary>

- 36 Mountain | land | Provides a reliable untapped red mana base for the deck.
- 1 Balin, Loremaster | draw | Supplies one of the deck's listed card-advantage pieces.
- 1 Key to the Side-Door | draw | Adds to the deck's available card flow.
- 1 Palantír of Orthanc | draw | Provides a dedicated draw slot.
- 1 Ragged Short Spear | draw | Supports the deck's card-advantage package.
- 1 Thrór's Map | draw | Adds another source of card flow.
- 1 Óin the Brave | draw | Helps keep cards coming while contributing a creature body.
- 1 Bilbo's Ring | interaction | Provides a flexible interaction slot in an artifact package.
- 1 Dwarven Mattock | interaction | Adds interaction while remaining part of the equipment suite.
- 1 Mithril Coat | interaction | Offers a dedicated interaction piece for important permanents.
- 1 The One Ring | interaction | Provides a high-impact interaction option.
- 1 Arcane Signet | ramp | Efficiently advances mana toward the commander and larger threats.
- 1 Bag End Banquet | ramp | Contributes to the deck's mana development.
- 1 Burn, Burn, Tree and Fern | ramp | Provides another ramp effect for reaching larger plays.
- 1 Dragon's Desire | ramp | Adds a dedicated ramp spell to accelerate the deck.
- 1 Fíli and Kíli, Joyous | ramp | Supports the mana plan with a legendary creature.
- 1 Long-Bodied Grey Dog | ramp | Provides a creature-based ramp slot.
- 1 Mox Amber | ramp | Supplies inexpensive acceleration alongside the deck's legendary cards.
- 1 Orcrist, Goblin-cleaver | ramp | Contributes ramp while fitting the equipment theme.
- 1 The Reaver Cleaver | ramp | Adds mana development in an equipment slot.
- 1 The Misty Mountains Cold | ramp | Provides another ramp effect for the top end.
- 1 Thorin, Company's Leader | ramp | Adds creature-based acceleration to the deck.
- 1 Troop of Ponies | ramp | Rounds out the ramp package with another creature.
- 1 Wayfarer's Bauble | ramp | Provides early mana development from an artifact.
- 1 Battle-Scarred Goblin | removal | Supplies a creature-based removal option.
- 1 Fire of Orthanc | removal | Provides a direct removal spell.
- 1 Giant's Boulder | removal | Adds flexible artifact-based removal.
- 1 Goblin Cratermaker | removal | Provides compact creature-based removal.
- 1 Goblin Fireleaper | removal | Adds another removal creature to the board.
- 1 Improvised Club | removal | Provides a low-commitment removal option.
- 1 Inferno Titan | removal | Combines a substantial body with removal utility.
- 1 Smaug, the Great Calamity // Spew Flame | removal | Adds a Dragon threat with a removal role.
- 1 Stone-Giant of High Pass | removal | Provides another large creature that fills a removal slot.
- 1 Call Forth the Tempest | wipe | Serves as a board-reset option when opponents get ahead.
- 1 Desolation of Smaug | wipe | Adds a second sweeping answer for crowded boards.
- 1 Glóin the Mighty // Easy Pickings | wipe | Provides a creature card that also occupies a wipe slot.
- 1 Last Light of Durin's Day | synergy | Provides the deck's dedicated synergy piece.
- 1 Desert Were-Worm | threat | Adds a Dragon Wurm as a standalone threat.
- 1 Andúril, Flame of the West | other | Expands the legendary equipment package.
- 1 Andúril, Narsil Reforged | other | Provides another legendary equipment option.
- 1 Bombur, Gentle Dreamer | other | Adds a legendary creature to support the deck's board presence.
- 1 Bothersome Noisemaker | other | Provides another creature for the deck's battlefield plan.
- 1 Dori, Bearer of Friends | other | Adds a legendary creature to the creature suite.
- 1 Dwarven Mauler | other | Provides an additional Dwarf body.
- 1 Dwarven Warriors | other | Adds another creature to establish the board.
- 1 Dáin Ironfoot | other | Provides another legendary Dwarf for the creature package.
- 1 Gandalf, Goblins' Bane // Flameshape | other | Adds a versatile legendary creature card.
- 1 Getaway Barrel | other | Provides a utility artifact slot.
- 1 Glamdring | other | Adds to the legendary equipment suite.
- 1 Goblin-town Flunkies | other | Provides another inexpensive creature body.
- 1 Gundabad Opportunist | other | Adds another Goblin creature to the board plan.
- 1 Guttersnipe | other | Provides a Goblin body alongside the deck's spells.
- 1 Iron Hills Stalwart | other | Adds another Dwarf to the creature base.
- 1 Long-Lost Lances | other | Expands the equipment package.
- 1 Misty Mountains Raider | other | Provides another Goblin creature for pressure.
- 1 Old Thrush | other | Adds a low-cost creature to establish the board.
- 1 Oliphaunt | other | Provides an additional large creature body.
- 1 Olog-hai Crusher | other | Adds a substantial creature to the attacking force.
- 1 Orcish Siegemaster | other | Provides another aggressive creature body.
- 1 Smaug's Fury | other | Adds a thematic spell to the Smaug deck.
- 1 Snowslope Hunter | other | Provides another Goblin creature for board presence.
- 1 Sting, Bilbo's Sword | other | Adds another legendary equipment card.
- 1 Tidings of War | other | Provides a thematic noncreature spell slot.
- 1 Well-Worn Spatula | other | Rounds out the deck's equipment package.

</details>

### 21. the Hobbit family, mana from outside

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 156 names.

Commander: Smaug the Impenetrable.

Grade: baseline, score 0.06, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $801.39 to buy, $801.39 the whole deck.

**Summary:** This Smaug deck develops its mana with artifacts, lands, and Dragon-themed acceleration, then puts pressure on the table with Goblins, Orcs, Trolls, Dragons, and heavily equipped attackers. It keeps cards flowing through legendary creatures and artifacts while using a broad suite of removal, instant-speed disruption, and a few board resets to clear a path. The deck aims to win through sustained creature combat backed by its Equipment package; in exchange, it is more focused on building a board than on holding up constant reactive answers.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The deck makes less mana on turn four than the norm, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The equipment-plus-attackers core is real (Andúril x2, Glamdring, Sting, Orcrist, Improvised Club, Goblin Plate Mail, Reaver Cleaver, Mithril Coat) and supports the summary's combat plan, but the creature base is a scatter of unconnected Goblins, Orcs, Trolls and one-off dragon pieces with no shared payoff, so the deck's bodies and its ramp/draw pieces pull only loosely in the same direction.
- PLAN theme_fit=partly: It is the right commander, the right format, and roughly the right power level, but the request limited non-mana cards to the Hobbit set and the list leans on a fair number of outside-set spells and creatures (Orcish Bowmasters, The One Ring, Languish, Night's Whisper, Guttersnipe, Goblin Cratermaker, Andúril, Glamdring, Sting) that are neither rocks nor lands.
- PLAN useful_as_built=yes: 36 lands with duals, fetches and fixing plus nine ramp pieces support a low-to-mid curve, and there are plenty of creature and equipment threats to actually close a game, so the deck can be picked up and played as it stands.
- PLAN summary_honest=yes: Every claim maps to cards present — artifact and land ramp, dragon-flavored acceleration, the Equipment package, legendary/artifact card flow, a broad removal suite and exactly the 'few' board resets it admits to (Languish, Desolation of Smaug) — and it openly concedes the deck is proactive rather than reactive.
- [INFO] `curve_summary`: average mana value 2.81 over 63 nonland cards
- [INFO] `mana_pass`: the builder moved 2 cards of the mana base to bring the deck inside its power level
- [WARN] `outside_requested_set`: the sets you named do not hold 12 cards: Ancient Tomb, Blood Crypt, Bloodstained Mire, Bojuka Bog, Command Tower, Dragonskull Summit, Evolving Wilds, Exotic Orchard, Fabled Passage, Marsh Flats, Polluted Delta, Scalding Tarn

<details><summary>The deck list</summary>

- 1 Ancient Tomb | land | An untapped land that accelerates the deck's mana development.
- 1 Blood Crypt | land | A dual land source for the deck's core colors.
- 1 Bloodstained Mire | land | A flexible land source that finds the needed basics or dual land.
- 1 Bojuka Bog | land | A black source with useful graveyard-focused utility.
- 1 Command Tower | land | A reliable source of the commander's colors.
- 1 Dragon-Cursed Halls | land | A themed colored land for the mana base.
- 1 Dragonskull Summit | land | A dual land source for the deck's core colors.
- 1 Evolving Wilds | land | Fixes colors by finding a basic land.
- 1 Exotic Orchard | land | Flexible color fixing in multiplayer games.
- 1 Fabled Passage | land | Finds the appropriate basic land while improving color access.
- 1 Goblin-town | land | A themed land that contributes to the mana base.
- 1 Marsh Flats | land | A fetch land that improves black access and finds dual lands.
- 1 Minas Morgul, Dark Fortress | land | A themed black land with utility.
- 1 Mount Doom | land | A themed colored land for the deck's mana base.
- 1 Polluted Delta | land | A fetch land that finds the needed black sources.
- 1 Scalding Tarn | land | A fetch land that improves red access and finds dual lands.
- 1 The Lonely Mountain | land | A Mountain land that supports red mana production.
- 10 Swamp | land | Basic black sources stabilize early plays and fetch-land targets.
- 9 Mountain | land | Basic red sources stabilize early plays and fetch-land targets.
- 1 Arcane Signet | ramp | Efficient color fixing and mana acceleration.
- 1 Bag End Banquet | ramp | A mana artifact that advances the deck's development.
- 1 Burn, Burn, Tree and Fern | ramp | The Saga contributes to the deck's mana acceleration plan.
- 1 Dragon's Desire | ramp | A thematic spell that advances the mana plan.
- 1 Mox Amber | ramp | A cheap legendary mana source alongside the commander and legendary permanents.
- 1 Orcrist, Goblin-cleaver | ramp | An Equipment that also supports mana development.
- 1 The Reaver Cleaver | ramp | A legendary Equipment that contributes to the mana plan.
- 1 Smaug the Magnificent | ramp | A Dragon that reinforces the deck's mana development.
- 1 Wayfarer's Bauble | ramp | Early ramp that also helps fix basic land access.
- 1 Balin, Loremaster | draw | A legendary creature included to maintain card flow.
- 1 Gollum, Riddle Master | draw | A thematic creature source of card advantage.
- 1 Key to the Side-Door | draw | An artifact source of ongoing card access.
- 1 Night's Whisper | draw | An efficient spell for replenishing cards.
- 1 Palantír of Orthanc | draw | A legendary artifact that supports sustained card access.
- 1 Rage into the Valley | draw | A red card-advantage spell for keeping pressure up.
- 1 Ragged Short Spear | draw | An Equipment that contributes to the deck's card flow.
- 1 The Master of Lake-town | draw | A legendary creature that supplies card advantage.
- 1 The Sackville-Bagginses | draw | A legendary creature that helps keep the hand supplied.
- 1 Thrór's Map | draw | A legendary artifact source of card access.
- 1 Óin the Brave | draw | A legendary creature included for card advantage.
- 1 Battle-Scarred Goblin | removal | A low-cost Goblin that handles opposing permanents.
- 1 Bilbo's Deadly Slice | removal | A focused removal spell.
- 1 Bitter Downfall | removal | A direct answer to a troublesome opposing permanent.
- 1 Crude Bent Blade | removal | An Equipment slot that also serves as removal.
- 1 Fire of Orthanc | removal | A red removal spell for answering threats.
- 1 Goblin Cratermaker | removal | A Goblin body that doubles as a removal option.
- 1 Improvised Club | removal | A flexible instant-speed removal option.
- 1 Orcish Bowmasters | removal | An efficient creature-based removal piece.
- 1 The Black Arrow | removal | A legendary Equipment that provides another answer.
- 1 Desolation of Smaug | wipe | A thematic board-reset effect.
- 1 Languish | wipe | A compact sweeper for clearing opposing creatures.
- 1 Desert Were-Worm | threat | A large creature threat that adds battlefield pressure.
- 1 Dreaded Bat-Cloud | threat | A creature threat that contributes to the attack plan.
- 1 Gollum the Abandoned | threat | A legendary creature threat for the deck's combat plan.
- 1 Great Goblin, Foul-Hearted | threat | A legendary Goblin that adds pressure to the board.
- 1 Guttersnipe | threat | A creature threat that rewards the deck's spells.
- 1 Haunt of the Dead Marshes | threat | A creature threat that broadens the board presence.
- 1 Nighthowl Pursuer | threat | A creature threat that helps pressure opponents.
- 1 Olog-hai Crusher | threat | A heavy creature threat for closing games through combat.
- 1 Orcish Siegemaster | threat | An Orc threat that adds to the attacking board.
- 1 Ravening Warg | threat | A creature threat that supports the combat plan.
- 1 The Great Goblin | threat | A legendary Goblin threat that adds board pressure.
- 1 Troll of Khazad-dûm | threat | A large creature threat for the late game.
- 1 Andúril, Flame of the West | synergy | A legendary Equipment that strengthens the deck's artifact-and-combat package.
- 1 Andúril, Narsil Reforged | synergy | A second legendary Equipment for the deck's combat package.
- 1 Down, Down to Goblin-town | synergy | A Goblin-themed Saga that supports the deck's creature package.
- 1 Glamdring | synergy | A legendary Equipment that complements the artifact package.
- 1 Goblin Plate Mail | synergy | An Equipment that supports the Goblin and combat themes.
- 1 Goblin-town Flunkies | synergy | A Goblin creature that reinforces the deck's tribal board presence.
- 1 Long-Lost Lances | synergy | An Equipment that supports creature-based combat.
- 1 Sting, Bilbo's Sword | synergy | A legendary Equipment that complements the combat package.
- 1 Supper for Spiders | synergy | A dedicated thematic synergy piece.
- 1 Well-Worn Spatula | synergy | An Equipment that adds to the deck's artifact synergies.
- 1 Bilbo's Ring | interaction | A legendary artifact interaction piece.
- 1 Dwarven Mattock | interaction | An Equipment that provides a flexible interaction slot.
- 1 Mithril Coat | interaction | A protective legendary artifact for key permanents.
- 1 My Precious // Allure of Power | interaction | A legendary artifact with an interaction role.
- 1 The One Ring | interaction | A powerful legendary artifact interaction piece.
- 1 Gathering of Darkness | interaction | A black spell held for disruptive interaction.
- 1 Getaway Barrel | interaction | An artifact utility piece for responding to opponents.
- 1 Smaug's Fury | interaction | A thematic instant-speed interaction option.
- 1 Giant's Boulder | ramp | the mana pass added it to bring the mana base inside the power level
- 1 Bolg's Company | ramp | the mana pass added it to bring the mana base inside the power level

</details>

### 22. a set family and a card from outside it

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 175 names.

Commander: Thranduil, the Elvenking.

Grade: baseline, score 0.32, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $394.37 to buy, $394.37 the whole deck.

**Summary:** This is a creature-forward Thranduil deck that develops its mana early, builds an Elf-centered board, and keeps cards flowing while applying pressure with a broad creature suite. It wins by maintaining battlefield presence and pushing that pressure through while using targeted answers, reactive spells, and a few reset buttons to keep opposing boards from taking over. The deck gives up some flexibility for a strongly themed creature core and a mana base centered on dependable basic lands.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: There is a workable creature-plus-interaction shell with ramp, draw, and removal, but the creature core pulls two ways at once — an Elf subtheme (Elvish Archdruid, Galadhrim Guide, Elvenking's Halls, Thranduil's Company) alongside a sizable Warg/spider package (Attercop, Ravening Warg, Wargling, Chief Warg's Company, Bejeweled Warg) that shares no payoff with the Elf half.
- PLAN theme_fit=partly: The commander, colors, format, and the requested Sol Ring are all correct and much of the list is Middle-earth flavored, but a large slice of the deck (Orcish Bowmasters, Lórien Revealed, Delighted Halfling, Celeborn the Wise, Arwen, Mithril Coat, plus generic staples like Languish, Night's Whisper, Elvish Mystic/Archdruid) comes from outside the Hobbit set the person named.
- PLAN useful_as_built=yes: Thirty-six lands with a heavily basic base in three colors, several cheap mana accelerants, a good spread of removal, three sweepers, and plenty of card draw means the deck casts its spells and can close a game on creature damage.
- PLAN summary_honest=partly: placeholder
- [INFO] `curve_summary`: average mana value 2.87 over 63 nonland cards
- [INFO] `mana_pass`: the builder moved 1 card of the mana base to bring the deck inside its power level
- [WARN] `outside_requested_set`: the sets you named do not hold 1 card: Sol Ring

<details><summary>The deck list</summary>

- 10 Forest | land | Basic green source for the mana base.
- 10 Island | land | Basic blue source for the mana base.
- 9 Swamp | land | Basic black source for the mana base.
- 1 Elven Passage | land | Elf-themed land for the mana base.
- 1 Elvenking's Halls | land | Themed land supporting the mana base.
- 1 Minas Morgul, Dark Fortress | land | Nonbasic land slot for the mana base.
- 1 Mirkwood | land | Themed nonbasic land for the mana base.
- 1 Rivendell | land | Nonbasic land slot for the mana base.
- 1 The Black Gate | land | Nonbasic land slot for the mana base.
- 1 The Shire | land | Themed nonbasic land for the mana base.
- 1 Sol Ring | ramp | Efficient artifact mana and a required inclusion.
- 1 Arcane Signet | ramp | Reliable artifact mana fixing.
- 1 Mox Amber | ramp | Low-cost legendary-focused mana acceleration.
- 1 Elvish Mystic | ramp | Early Elf mana acceleration.
- 1 Wayfarer's Bauble | ramp | Early mana development from an artifact.
- 1 Elvish Archdruid | ramp | Elf-based mana acceleration.
- 1 Wood Elves | ramp | Elf creature that advances mana development.
- 1 Thranduil's Company | ramp | Elf-themed mana acceleration.
- 1 Woodland Weavemaster | ramp | Elf mana acceleration for the creature plan.
- 1 Elvish Visionary | draw | Elf creature card advantage.
- 1 Hithlain Knots | draw | Instant-speed card advantage.
- 1 Key to the Side-Door | draw | Artifact card advantage.
- 1 Lórien Revealed | draw | Flexible sorcery card advantage.
- 1 Night's Whisper | draw | Efficient sorcery card advantage.
- 1 Palantír of Orthanc | draw | Legendary artifact card advantage.
- 1 Plunder the Trollshaws | draw | Instant card advantage.
- 1 Reverent Howl | draw | Instant card advantage.
- 1 Uncover the Moon-Letters | draw | Enchantment-based card advantage.
- 1 Last March of the Ents | draw | Sorcery card advantage.
- 1 Ithilien Kingfisher | draw | Creature-based card advantage.
- 1 Confusticate and Bebother | interaction | Instant interaction.
- 1 Bilbo's Ring | interaction | Equipment-based interaction.
- 1 Mithril Coat | interaction | Equipment-based interaction.
- 1 My Precious // Allure of Power | interaction | Flexible equipment and instant interaction.
- 1 Old Fat Spider Can't See Me | interaction | Saga-based interaction.
- 1 Sound the Trumpets | interaction | Instant interaction.
- 1 Stern Scolding | interaction | Low-cost instant interaction.
- 1 Thranduil's Decree | interaction | Themed instant interaction.
- 1 Bilbo's Deadly Slice | removal | Instant removal.
- 1 Bitter Downfall | removal | Instant removal.
- 1 Crude Bent Blade | removal | Equipment-based removal.
- 1 Enchanted River's Grasp | removal | Aura removal.
- 1 Giant's Boulder | removal | Artifact-based removal.
- 1 Orcish Bowmasters | removal | Creature-based removal.
- 1 Quarrel | removal | Instant removal.
- 1 Stir Up Trouble | removal | Sorcery removal.
- 1 The Black Arrow | removal | Themed equipment removal.
- 1 Gnashing of Teeth | wipe | Sorcery board wipe.
- 1 Languish | wipe | Sorcery board wipe.
- 1 Raise the Palisade | wipe | Sorcery board wipe.
- 1 Arwen, Weaver of Hope | synergy | Legendary Elf supporting the Elf-centered creature suite.
- 1 Boughside Wanderers | synergy | Elf creature supporting the deck's Elf theme.
- 1 Cantankerous Keepers | synergy | Elf creature supporting the deck's Elf theme.
- 1 Celeborn the Wise | synergy | Legendary Elf supporting the Elf-centered creature suite.
- 1 Galadhrim Guide | synergy | Elf creature supporting the deck's Elf theme.
- 1 Galion, Elvenking's Butler | synergy | Legendary Elf tied to the Elvenking theme.
- 1 Guardian of the Halls | synergy | Elf creature supporting the deck's Elf theme.
- 1 Lothlórien Lookout | synergy | Elf creature supporting the deck's Elf theme.
- 1 Mirkwood Meditator | synergy | Elf creature supporting the deck's Elf theme.
- 1 Supper for Spiders | synergy | Dedicated synergy card from the shortlist.
- 1 Attercop | threat | Creature threat that adds battlefield pressure.
- 1 Bejeweled Warg | threat | Creature threat that also fits the mana-focused creature base.
- 1 Chief Warg's Company | threat | Wolf creature threat for battlefield pressure.
- 1 Chief of the Wilds | threat | Legendary Wolf creature threat.
- 1 Desolation Prowler | threat | Wolf creature threat for battlefield pressure.
- 1 Duskwatch Hunter | threat | Wolf creature threat for battlefield pressure.
- 1 Head of the Hunt | threat | Wolf creature threat for battlefield pressure.
- 1 Nighthowl Pursuer | threat | Wolf creature threat for battlefield pressure.
- 1 Ravening Warg | threat | Wolf creature threat for battlefield pressure.
- 1 Rhovanion Rampager | threat | Wolf creature threat for battlefield pressure.
- 1 Wargling | threat | Wolf creature threat for battlefield pressure.
- 1 Willow-Wind | threat | Creature threat that broadens the battlefield presence.
- 1 Delighted Halfling | ramp | the mana pass added it to bring the mana base inside the power level

</details>

### 23. two set families at once

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 203 names.

Commander: Kíli the Resourceful.

Grade: typical, score 0.47, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $315.76 to buy, $315.76 the whole deck.

**Summary:** Kíli leads a white creature-and-Equipment deck built around Dwarves, artifacts, and recognizable Hobbit and Bloomburrow characters. Develop mana early, keep cards flowing through creatures and artifacts, then establish equipped attackers and durable legendary threats while using removal and sweepers to reset unfavorable boards. The deck gives up some raw finishing density for a broad, thematic board presence and a large package of protective interaction.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The bulk of the list coheres as a mono-white creature-plus-Equipment build with token/artifact support and plenty of draw and removal, but a number of pieces (Maskwood Nexus, Karn the Great Creator, Dawn of a New Age, Mentor of the Meek's small-creature payoff) point at plans the rest of the deck doesn't really support, making it read as white goodstuff with an equipment lean.
- PLAN theme_fit=partly: The person named a hard limit—cards from The Hobbit and Bloomburrow—and while Kíli, Fíli, Dáin, Ori, Bilbo, Fountainport, Three Tree City, Carrot Cake and friends honor it, a large slice of the deck (Swords to Plowshares, Sol Ring, Skullclamp, Elspeth, Skyclave Apparition, Blade Splicer, Karn, Academy Manufactor, Loran, the LOTR-proper Equipment) comes from well outside the two requested sets.
- PLAN useful_as_built=yes: 37 lands in essentially mono-white with ten ramp pieces, a sane curve, eleven draw sources, ample spot removal and three sweepers, and real win paths in equipped legends, Elspeth and Helm of the Host means it can be shuffled up and played as written.
- PLAN summary_honest=partly: It is candid about the thin finishing density and the broad board-presence plan, and the Equipment/Dwarf/legendary elements do exist, but it quietly presents the deck as a Hobbit/Bloomburrow build when roughly half the nonland cards come from neither set, and 'Dwarves' is a handful of bodies rather than a tribal engine.
- [INFO] `curve_summary`: average mana value 2.74 over 62 nonland cards
- [INFO] `basics_added`: the list was 1 card short, so the builder added 1 basic land

<details><summary>The deck list</summary>

- 1 Castle Ardenvale | land | White-producing land for the mana base.
- 1 Command Tower | land | Reliable colored land for the commander deck.
- 1 Exotic Orchard | land | Flexible colored land.
- 1 Fountainport | land | Utility land slot that stays within the colorless-land limit.
- 1 Minas Tirith | land | White land with a Hobbit setting identity.
- 29 Plains | land | Primary white sources to support the deck's colored cards.
- 1 Reliquary Tower | land | Colorless utility land.
- 1 Rogue's Passage | land | Colorless utility land.
- 1 Three Tree City | land | Colorless utility land that fits the creature-focused build.
- 1 Arcane Signet | ramp | Efficient fixing and acceleration.
- 1 Fellwar Stone | ramp | Two-mana mana acceleration.
- 1 Loyal Warhound | ramp | Creature-based mana development.
- 1 Mind Stone | ramp | Low-cost mana acceleration.
- 1 Mox Amber | ramp | Legendary-focused acceleration.
- 1 Ornithopter of Paradise | ramp | Early creature-based mana fixing.
- 1 Patchwork Banner | ramp | Mana acceleration that supports the creature base.
- 1 Sol Ring | ramp | Fast artifact acceleration.
- 1 Thought Vessel | ramp | Two-mana artifact acceleration.
- 1 Wayfarer's Bauble | ramp | Early land-based mana development.
- 1 Caretaker's Talent | draw | Dedicated card-flow piece.
- 1 Circuit Mender | draw | Artifact creature that supplies card flow.
- 1 Cut a Deal | draw | Straightforward card-flow spell.
- 1 Dawn of a New Age | draw | Persistent card-flow option.
- 1 Fountainport Bell | draw | Artifact-based card flow.
- 1 Heirloom Epic | draw | Artifact card-flow engine.
- 1 Idol of Oblivion | draw | Low-cost artifact card flow.
- 1 Inspiring Overseer | draw | Creature-based card flow.
- 1 Mentor of the Meek | draw | Creature-focused card-flow engine.
- 1 Skullclamp | draw | Efficient equipment-based card flow.
- 1 Spirited Companion | draw | Low-cost creature that supplies card flow.
- 1 Baird, Steward of Argive | interaction | Creature-based table interaction.
- 1 Bilbo's Ring | interaction | Legendary Equipment interaction piece.
- 1 Dawn's Truce | interaction | Instant-speed interaction.
- 1 Luminous Broodmoth | interaction | Creature-based protection and interaction.
- 1 Mithril Coat | interaction | Equipment-based protection.
- 1 Reprieve | interaction | Flexible instant-speed interaction.
- 1 Selfless Spirit | interaction | Creature-based protective interaction.
- 1 Swiftfoot Boots | interaction | Equipment-based protection for key creatures.
- 1 Angel of the Ruins | removal | Artifact creature that answers opposing pieces.
- 1 Banishing Light | removal | Broad permanent removal.
- 1 Fiend Hunter | removal | Creature-based removal.
- 1 Generous Gift | removal | Flexible instant-speed removal.
- 1 Hanged Executioner | removal | Creature-based removal option.
- 1 Loran of the Third Path | removal | Utility creature with removal.
- 1 Skyclave Apparition | removal | Efficient creature-based removal.
- 1 Swords to Plowshares | removal | Efficient single-target removal.
- 1 The Black Arrow | removal | Legendary Equipment removal piece.
- 1 Dusk // Dawn | wipe | Flexible sweep effect.
- 1 Elspeth, Sun's Champion | wipe | Planeswalker with a board-clearing role.
- 1 Martial Coup | wipe | Creature-focused sweep effect.
- 1 Blade Splicer | synergy | Artifact-creature synergy piece.
- 1 Carrot Cake | synergy | Artifact synergy piece from Bloomburrow.
- 1 Dáin, Lord of the Iron Hills | synergy | Dwarf-focused synergy card.
- 1 Iron Hills Blacksmith | synergy | Dwarf and artifact synergy piece.
- 1 Ori, Keeper of Songs | synergy | Dwarf-focused synergy card.
- 1 Tangle Tumbler | synergy | Artifact synergy piece.
- 1 Fíli the Pathfinder | threat | Legendary creature threat with Hobbit flavor.
- 1 Karn, the Great Creator | threat | Standalone planeswalker threat.
- 1 Academy Manufactor | other | Colorless artifact creature that supports the artifact shell.
- 1 Andúril, Flame of the West | other | Signature legendary Equipment for the combat package.
- 1 Andúril, Narsil Reforged | other | Additional legendary Equipment for the combat package.
- 1 Bilbo, Unexpected Adventurer | other | Hobbit-themed legendary creature.
- 1 Boss's Chauffeur | other | Creature that broadens the board presence.
- 1 Dwarven Shortsword | other | Low-cost Equipment for the creature package.
- 1 Dúnedain Blade | other | Equipment that fits the Middle-earth theme.
- 1 Glamdring | other | Legendary Equipment for the artifact package.
- 1 Helm of the Host | other | High-impact legendary Equipment.
- 1 Maskwood Nexus | other | Artifact support for the mixed creature roster.
- 1 Sting, Bilbo's Sword | other | Legendary Equipment with Hobbit flavor.
- 1 Three Tree Mascot | other | Artifact creature that fits the Bloomburrow setting.
- 1 Well-Worn Spatula | other | Low-cost Equipment for the artifact shell.

</details>

### 24. a 60-card deck from one set

Format: Modern. Theme: aggro. Pool: any_card. Shortlist: 96 names.

Grade: baseline, score 0.37, model 20260914T154223Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $44.76 to buy, $44.76 the whole deck.

**Summary:** This mono-red Bloomburrow aggro deck presses the table with a dense lineup of creatures, using Emberheart Challenger and Hearthborn Battler to support its attacking core. Efficient removal clears the way, while Artist's Talent and Might of the Meek help maintain pressure after the first wave. It wins through sustained creature attacks, with Dragonhawk, Fate's Tempest providing a stronger finishing threat, but gives up broader defensive tools and long-game flexibility for speed and consistency.

The quality model grades this deck at the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The color sources cover the pips, and that raises the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The creature core plus cheap removal and Artist's Talent does point at one aggressive plan, but 24 lands and two Mind Stone are ramp-and-grind inclusions that pull against the low-curve beatdown the rest of the list is built for.
- PLAN theme_fit=partly: It is a mono-red Modern aggro deck built mostly from Bloomburrow, but Abrade and Mind Stone come from outside the requested set, so the "from Bloomburrow" limit is not fully honored.
- PLAN useful_as_built=yes: Mana is clean with 24 Mountains for a curve topping at four or five, and there are twenty threats plus removal, so the deck functions and closes games as presented.
- PLAN summary_honest=partly: It describes the creatures, removal, and Dragonhawk finisher accurately, but claiming the list trades long-game flexibility "for speed" glosses over the 24-land, Mind Stone configuration that is notably slower than an aggro build implies.
- [INFO] `curve_summary`: average mana value 3.06 over 36 nonland cards

<details><summary>The deck list</summary>

- 24 Mountain | land | Provides a consistent mono-red mana base for early aggression and higher-cost threats.
- 2 Mind Stone | ramp | Adds a small amount of mana acceleration while retaining a draw role later in the game.
- 4 Artist's Talent | draw | Supplies repeatable card access to keep the aggressive plan moving.
- 2 Might of the Meek | draw | Provides efficient card draw support for the low-curve creature plan.
- 4 Abrade | removal | Gives the deck flexible, efficient removal to clear opposing obstacles.
- 2 Agate Assault | removal | Adds more direct removal so attackers can keep connecting.
- 4 Emberheart Challenger | synergy | Supports the deck's aggressive red creature synergies.
- 4 Hearthborn Battler | synergy | Builds out the synergy package alongside the deck's creature pressure.
- 4 Frilled Sparkshooter | threat | Provides a full set of proactive threats for the early and middle turns.
- 4 Reptilian Recruiter | threat | Adds another reliable set of red creature threats.
- 4 Teapot Slinger | threat | Keeps the creature count high and the attack-focused plan dense.
- 2 Dragonhawk, Fate's Tempest | threat | Serves as a larger top-end threat when the game goes beyond the opening turns.

</details>

### 25. use no card of an owned precon

Format: Commander. Theme: Avengers superheroes. Pool: owned_first. Shortlist: 152 names.

Commander: Syr Konrad, the Grim.

Grade: baseline, score 0.36, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $235.28 the whole deck.

**Summary:** This is a black Avengers-flavored Syr Konrad deck that develops its mana, keeps cards flowing, and uses a deep suite of answers to clear space for its creatures and equipment. It wins by sustaining pressure with its threats and having Syr Konrad help close games as the table trades resources. The deck gives up dedicated creature-threat density and specialized synergy pieces in exchange for unusually broad removal, several reset buttons, and a strong Marvel setting through its lands and featured cards.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The deck draws more than the norm, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=no: The list is a pile of generic mono-black removal, equipment, and card draw with only incidental support for Syr Konrad's death/mill triggers — no mill, no sacrifice engine, no recursion loop — and the -1/-1 counter cards (Blowfly Infestation, Contagion Clasp, Grim Poppet, Liliana Death Wielder, Black Sun's Zenith) point at a counters subtheme the rest of the deck doesn't build on, so pieces pull in different directions.
- PLAN theme_fit=partly: The request was an Avengers superheroes deck, yet the overwhelming majority of the list is Lord of the Rings orcs, Nazgûl and Mordor cards plus generic black staples, with only a handful of Marvel cards (and several of those, like Baron Strucker, M.O.D.O.K., HYDRA/Doom, are villains rather than Avengers); the mono-black bracket-3 library-first constraints are otherwise respected.
- PLAN useful_as_built=partly: The mana base is fine and the curve is playable, but with roughly a dozen creatures, no evasive or recursive win engine, and a commander whose triggers are barely enabled, the deck can answer things all game while having very few real routes to actually finish a table.
- PLAN summary_honest=partly: It is candid about lacking creature density and synergy pieces and about the removal-heavy build, but claims the deck 'sustains pressure with its threats' when it has barely a dozen creatures, and it says Syr Konrad helps close games without noting nothing in the list actually feeds him.
- [WARN] `not_owned`: Syr Konrad, the Grim: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.10 over 63 nonland cards
- [WARN] `profile_off_band`: the draw count is 18, and bracket 3 wants 8 to 14
- [WARN] `profile_off_band`: the removal count is 23, and bracket 3 wants 6 to 12

<details><summary>The deck list</summary>

- 28 Swamp | land | Primary black mana base.
- 1 Avengers Tower | land | Avengers-themed land slot.
- 1 Baxter Building | land | Superhero setting land slot.
- 1 Big Apple, 3 a.m. | land | Superhero setting land slot.
- 1 Castle Doom | land | Marvel villain setting land slot.
- 1 Castle Locthwain | land | Black-aligned utility land.
- 1 Ifnir Deadlands | land | Black utility land.
- 1 Minas Morgul, Dark Fortress | land | Black-themed utility land.
- 1 Takenuma, Abandoned Mire | land | Black utility land.
- 1 Sol Ring | ramp | Efficient early mana development.
- 1 Arcane Signet | ramp | Reliable mana acceleration.
- 1 Wayfarer's Bauble | ramp | Early mana development.
- 1 Chromatic Lantern | ramp | Mana acceleration for the deck's setup turns.
- 1 Commander's Sphere | ramp | Flexible mana acceleration.
- 1 Thought Vessel | ramp | Mana acceleration.
- 1 Springleaf Drum | ramp | Low-cost mana development.
- 1 Ring of the Lucii | ramp | Mana acceleration.
- 1 Wickersmith's Tools | ramp | Mana acceleration.
- 1 Inherited Envelope | ramp | Mana acceleration.
- 1 Baron Strucker, HYDRA Overlord | draw | Marvel-themed card-advantage option.
- 1 Buster Sword | draw | Equipment-based card-advantage option.
- 1 Call of the Ring | draw | Ongoing card-advantage option.
- 1 Dusk Urchins | draw | Creature-based card-advantage option.
- 1 Eventide's Shadow | draw | Card-advantage option.
- 1 Grave Venerations | draw | Card-advantage option.
- 1 Hoarder's Greed | draw | Card-advantage option.
- 1 Idol of Oblivion | draw | Artifact-based card-advantage option.
- 1 Lembas | draw | Low-cost card-advantage option.
- 1 Mask of Memory | draw | Equipment-based card-advantage option.
- 1 Massacre Girl, Known Killer | draw | Creature-based card-advantage option.
- 1 Mordor Muster | draw | Card-advantage option.
- 1 Nasty End | draw | Instant-speed card-advantage option.
- 1 Night's Whisper | draw | Efficient card-advantage option.
- 1 Obsessive Pursuit | draw | Ongoing card-advantage option.
- 1 Skullclamp | draw | Efficient equipment-based card advantage.
- 1 Stone of Erech | draw | Artifact-based card-advantage option.
- 1 The Vision | draw | Avengers-themed card-advantage option.
- 1 Champion's Helm | interaction | Protective equipment for key creatures.
- 1 Darksteel Plate | interaction | Protective equipment for key creatures.
- 1 Lightning Greaves | interaction | Protective equipment for key creatures.
- 1 Orcish Medicine | interaction | Flexible interaction slot.
- 1 Stolen Stark Tech | interaction | Marvel-themed equipment interaction.
- 1 Swiftfoot Boots | interaction | Protective equipment for key creatures.
- 1 The Walls of Ba Sing Se | interaction | Creature-based interaction option.
- 1 Bitter Downfall | removal | Single-target answer.
- 1 Bitter Triumph | removal | Flexible single-target answer.
- 1 Blight Rot | removal | Instant-speed answer.
- 1 Blowfly Infestation | removal | Persistent removal tool.
- 1 Claim the Precious | removal | Single-target answer.
- 1 Contagion Clasp | removal | Artifact-based removal option.
- 1 Dismember | removal | Low-cost removal option.
- 1 Epic Downfall | removal | Single-target answer.
- 1 Fatal Push | removal | Efficient instant-speed answer.
- 1 Grim Poppet | removal | Creature-based removal option.
- 1 Heartless Act | removal | Flexible instant-speed answer.
- 1 Incremental Blight | removal | Multi-target removal option.
- 1 Infernal Grasp | removal | Reliable single-target answer.
- 1 Lash of the Balrog | removal | Sorcery-speed removal option.
- 1 Liliana, Death Wielder | removal | Planeswalker-based removal option.
- 1 Meteor Sword | removal | Equipment-based removal option.
- 1 Nameless Inversion | removal | Flexible instant-speed answer.
- 1 Orcish Bowmasters | removal | Creature-based removal option.
- 1 Overkill | removal | Instant-speed removal option.
- 1 Ringwraiths | removal | Creature-based removal option.
- 1 Skinrender | removal | Creature-based removal option.
- 1 Voracious Fell Beast | removal | Creature-based removal option.
- 1 Witch-king of Angmar | removal | Legendary creature removal option.
- 1 Black Sun's Zenith | wipe | Scalable board reset.
- 1 Extinction Event | wipe | Board-reset option.
- 1 M.O.D.O.K. | wipe | Marvel-themed board-reset option.
- 1 One Ring to Rule Them All | wipe | Saga-based board reset.
- 1 Zodiark, Umbral God | wipe | Creature-based board-reset option.

</details>

