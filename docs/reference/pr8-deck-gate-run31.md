# PR-8 deck gate

Run date: 2026-09-21. Card snapshot: 2026-09-04.

Verdict: FAIL. 25 of 25 decks passed every block check, 0 invented names reached the user, and 1 summaries stated a false rule of the game. All three bars are zero tolerance.

## Summary

| Measure | Value |
|---|---|
| Prompts | 25 |
| Decks returned | 25 |
| Decks with no block finding | 25 |
| Invented names that reached the user | 0 |
| Decks that needed the repair turn | 0 |
| Summaries judged (F-26) | 25 |
| Summaries that state a rule of the game | 2 |
| Summaries that state a FALSE rule | 1 |
| Judge errors | 0 |
| Case assertions missed (PR-28b) | 0 |
| Decks the plan judge read (PR-15, information) | 25 |
| Mean plan score, 0 to 1 | 0.74 |
| Plan reasons the judge left empty | 2 |
| Errors | 0 |
| Prompt version | 15 |
| Calls | 75 |
| Cost | $2.8348 |
| Time | 2306 seconds |

## Run

- Suite `decks`, run `pr8-deck-gate-run31`, on 2026-09-21, commit `1f8f15d`.
- Roles: generate on `gpt-5.6-terra` (openai, effort medium), judge on `claude-opus-5` (anthropic, effort medium), repair on `gpt-5.6-terra` (openai, effort medium).
- Versions: card snapshot 2026-09-04, generate prompt version 15, plan_rubric prompt version 4, precons `5.3.0+20260903`, quality_model `20260914T154223Z`.
- Calls: 75. Cost: $2.8348. Time: 2306 seconds.

## Findings by code

A block stops the deck. A warning and a note are reports.

| Code | Count |
|---|---|
| `curve_summary` | 25 |
| `not_owned` | 22 |
| `profile_off_band` | 10 |
| `mana_pass` | 5 |
| `basics_added` | 4 |
| `outside_requested_set` | 2 |
| `finisher_short` | 2 |
| `cards_trimmed` | 2 |
| `bracket_cut` | 1 |

By severity: BLOCK 0. WARN 36. INFO 37. 

## The set filter (PR-17B)

A deck asked for a set holds cards of that set family alone. Three things earn an exception, and each one is marked: a card the reader named (D-381), a mana card the reader allowed from outside (D-382), and a basic land, which no set limit filters (D-378).

| # | Prompt | Sets | In set | Fill | Deck cards outside | Marked |
|---|---|---|---|---|---|---|
| 19 | the Hobbit family, two colours | `hob,hoc` | 170 | 0 | 0 | 0 |
| 20 | the Hobbit family, a delegated commander | `hob,hoc` | 75 | 0 | 0 | 0 |
| 21 | the Hobbit family, mana from outside | `hob,hoc` | 127 | 26 | 17 | 17 |
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

Grade: bad, score 0.02, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $876.17 to buy, $876.17 the whole deck.

**Summary:** This lifegain deck develops a resilient creature board, uses Karlov of the Ghost Council as a major source of pressure and control, and backs combat with several dedicated alternate ways to close the game. It has broad answers, card flow, and enough mana development to support its stronger threats, but it gives up some raw speed and can be vulnerable when its board is repeatedly cleared.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core is a consistent lifegain-payoff shell (Karlov, Heliod, Vito, Archangel of Thune, Well of Lost Dreams, Aetherflux, Test of Endurance) but the ten-card artifact "ramp" package is a grab-bag of slow, low-impact pieces that barely produce mana, and three board wipes sit awkwardly with a plan built on developing a wide creature board.
- PLAN theme_fit=yes: It is a two-color W/B Commander list headed by Karlov with lifegain triggers, payoffs, and life-total win conditions throughout, at a power level consistent with bracket 3.
- PLAN useful_as_built=yes: 38 lands with solid duals and fixing, a playable curve, plenty of removal and card draw, and several distinct routes to victory make it immediately playable, even if the artifact slots are weak.
- PLAN summary_honest=yes: Every claim — creature board, Karlov pressure, alternate wins (Approach, Test of Endurance, Angel of Destiny, Aetherflux Reservoir), card flow, removal — is backed by cards in the list, and it openly admits the slow mana and vulnerability to wraths.
- [INFO] `curve_summary`: average mana value 3.35 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Archivist of Oghma | draw | Provides a dedicated draw piece for sustained resources.
- 1 Cosmos Elixir | draw | Provides a dedicated draw piece for sustained resources.
- 1 Dawn of Hope | draw | Provides a dedicated draw piece for sustained resources.
- 1 Enduring Innocence | draw | Provides a dedicated draw piece for sustained resources.
- 1 Exemplar of Light | draw | Provides a dedicated draw piece for sustained resources.
- 1 Mangara, the Diplomat | draw | Provides a dedicated draw piece for sustained resources.
- 1 Sigarda's Splendor | draw | Provides a dedicated draw piece for sustained resources.
- 1 The Gaffer | draw | Provides a dedicated draw piece for sustained resources.
- 1 Tymna the Weaver | draw | Provides a dedicated draw piece for sustained resources.
- 1 Vampiric Rites | draw | Provides a dedicated draw piece for sustained resources.
- 1 Well of Lost Dreams | draw | Provides a dedicated draw piece for sustained resources.
- 1 Alseid of Life's Bounty | interaction | Protects the board with a flexible interaction slot.
- 1 Courageous Resolve | interaction | Protects key permanents with a flexible interaction slot.
- 1 Faith's Shield | interaction | Protects key permanents with a flexible interaction slot.
- 1 Metropolis Reformer | interaction | Adds a protective interaction body to the creature plan.
- 1 Paladin Danse, Steel Maverick | interaction | Adds a versatile interaction piece to the board.
- 1 Restoration Magic | interaction | Provides a flexible interaction slot for key moments.
- 1 Sword of Light and Shadow | interaction | Provides a durable interaction tool that fits the creature plan.
- 1 Werefox Bodyguard | interaction | Adds interaction while contributing to the board.
- 1 Brightclimb Pathway // Grimclimb Pathway | land | Provides an untapped choice of either deck color.
- 1 Caves of Koilos | land | Provides flexible untapped black or white mana.
- 1 City of Brass | land | Provides flexible access to both deck colors.
- 1 Command Tower | land | Provides reliable access to both deck colors.
- 1 Concealed Courtyard | land | Provides efficient black or white mana early.
- 1 Eclipsed Steppe | land | Provides access to both deck colors.
- 1 Godless Shrine | land | Provides reliable black and white mana.
- 1 Mana Confluence | land | Provides flexible access to both deck colors.
- 1 Marsh Flats | land | Finds the basic mana needed for the opening turns.
- 11 Plains | land | Forms a reliable white basic-land base.
- 1 Prismatic Vista | land | Finds the basic mana needed for the opening turns.
- 1 Scrubland | land | Provides reliable black and white mana.
- 1 Shattered Sanctum | land | Provides efficient black or white mana.
- 1 Silent Clearing | land | Provides flexible black or white mana.
- 11 Swamp | land | Forms a reliable black basic-land base.
- 1 Vault of Champions | land | Provides efficient black or white mana.
- 1 Altar of the Pantheon | ramp | Adds mana development while helping cover both colors.
- 1 Colossal Plow | ramp | Adds an early mana-development piece.
- 1 Druidic Satchel | ramp | Adds mana development while contributing utility.
- 1 Hot Dog Cart | ramp | Adds an early mana-development piece.
- 1 Nuka-Cola Vending Machine | ramp | Adds mana development to support the midgame.
- 1 Orazca Relic | ramp | Adds mana development to support the midgame.
- 1 Phial of Galadriel | ramp | Adds mana development to support the midgame.
- 1 Potioner's Trove | ramp | Adds mana development for the deck's larger spells.
- 1 Pristine Talisman | ramp | Adds mana development that suits the lifegain plan.
- 1 The Celestus | ramp | Adds mana development with useful late-game utility.
- 1 Aetherflux Reservoir | removal | Supplies a dedicated removal option for problematic creatures.
- 1 Ayli, Eternal Pilgrim | removal | Provides a repeatable removal option on a low-cost body.
- 1 Gumdrop Poisoner // Tempt with Treats | removal | Provides flexible creature removal.
- 1 Murderous Rider // Swift End | removal | Provides flexible removal in a compact slot.
- 1 Nightmare's Thirst | removal | Provides efficient targeted removal.
- 1 Solitude | removal | Provides an efficient creature-removal option.
- 1 Umezawa's Jitte | removal | Provides repeatable removal utility for the creature plan.
- 1 Vona, Butcher of Magan | removal | Adds removal on a threatening creature body.
- 1 Witch of the Moors | removal | Adds removal that rewards the deck's core plan.
- 1 Cleric Class | synergy | Strengthens the deck's central lifegain theme.
- 1 Heliod, Sun-Crowned | synergy | Strengthens the deck's central lifegain theme.
- 1 Righteous Valkyrie | synergy | Rewards the lifegain plan while building the board.
- 1 Serra Ascendant | synergy | Provides an efficient creature for the lifegain plan.
- 1 Vito, Thorn of the Dusk Rose | synergy | Turns the lifegain theme into pressure on opponents.
- 1 Voice of the Blessed | synergy | Provides a growing creature for the lifegain plan.
- 1 Bloodthirsty Aerialist | synergy | Provides a growing evasive creature for the lifegain plan.
- 1 Archangel of Thune | threat | Provides a powerful top-end threat for the creature plan.
- 1 Archon of Sun's Grace | threat | Adds a resilient threat that develops the board.
- 1 Attended Healer | threat | Adds a board-building threat at an efficient point on the curve.
- 1 Celestine, the Living Saint | threat | Provides a substantial threat for the midgame.
- 1 Cliffhaven Vampire | threat | Turns the deck's theme into meaningful pressure.
- 1 Divinity of Pride | threat | Provides a large evasive threat for combat finishes.
- 1 Enduring Tenacity | threat | Provides a durable threat that fits the deck's theme.
- 1 Evra, Halcyon Witness | threat | Adds a substantial combat threat.
- 1 Nykthos Paragon | threat | Rewards the lifegain plan with a high-impact threat.
- 1 Rhox Faithmender | threat | Strengthens the lifegain plan on a sturdy creature.
- 1 Twinblade Paladin | threat | Adds a combat-focused threat for the lifegain plan.
- 1 Valkyrie Harbinger | threat | Provides a high-impact top-end threat.
- 1 Angel of Destiny | wincon | Provides a dedicated alternate route to closing the game.
- 1 Approach of the Second Sun | wincon | Provides a dedicated alternate route to closing the game.
- 1 Test of Endurance | wincon | Provides a dedicated alternate route to closing the game.
- 1 Fumigate | wipe | Resets crowded creature boards while fitting the deck's plan.
- 1 Kaya's Wrath | wipe | Provides a clean reset for overwhelming creature boards.
- 1 The Meathook Massacre | wipe | Provides a scalable board reset with lasting relevance.

</details>

### 2. aristocrats Commander, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 181 names.

Commander: Denethor, Ruling Steward.

Grade: typical, score 0.41, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $199.42 the whole deck.

**Summary:** This is a white-black aristocrats deck built to keep creatures flowing, turn their deaths into value, and use Denethor, Ruling Steward to make Soldier tokens and drain the table through sacrifice. Early mana and a large draw package keep the board supplied, while protection effects help Denethor and key creatures survive long enough to establish a sacrifice engine. The deck closes through steady life-loss pressure or its larger finishers, with broad removal and a few board wipes to recover from opposing boards. It gives up raw speed for a more incremental, board-dependent plan and can be vulnerable when its creatures or commander are repeatedly contained.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=no: The cards pull in two directions: a large equipment/protection suite (Puresteel Paladin, Sword of the Animist, Skullclamp, Champion's Helm, Darksteel Plate, Lightning Greaves, Together Forever, Clever Concealment) plus three sweepers and big white finishers, while the sacrifice engine the summary names barely exists — there are almost no free sac outlets, no token swarm, and no drain payoffs like Blood Artist effects.
- PLAN theme_fit=partly: The commander and a handful of cards (Nasty End, Gollum, Vengeful Villagers, Grave Venerations) gesture at aristocrats, but the bulk of the 99 is an equipment/protection midrange goodstuff pile, so the requested aristocrats deck is not really what was delivered.
- PLAN useful_as_built=partly: The mana base, ramp, draw, and removal are solid enough to sit down and play a midrange game, but the win path is thin and incoherent — the three board wipes fight the creature-based plan, and the promised drain kill simply is not there.
- PLAN summary_honest=no: It promises a sacrifice engine, creatures flowing, and draining the table through life loss, but the list has only a couple of one-shot sacrifice effects, no recurring token production beyond the commander, and no aristocrat drain triggers, and it never mentions that the deck is actually built around equipment and protection spells.
- [INFO] `curve_summary`: average mana value 2.84 over 63 nonland cards

<details><summary>The deck list</summary>

- 12 Plains | land | Provides reliable white mana for the deck.
- 12 Swamp | land | Provides reliable black mana for the deck.
- 1 Command Tower | land | Provides flexible commander-color mana.
- 1 City of Brass | land | Provides flexible colored mana.
- 1 Exotic Orchard | land | Provides flexible colored mana.
- 1 Marsh Flats | land | Finds a needed basic land color.
- 1 Fabled Passage | land | Finds a needed basic land color.
- 1 Evolving Wilds | land | Finds a needed basic land color.
- 1 Terramorphic Expanse | land | Finds a needed basic land color.
- 1 Plaza of Heroes | land | Adds another flexible mana source.
- 1 Spire of Industry | land | Adds flexible colored mana alongside artifacts.
- 1 Grand Coliseum | land | Provides access to both deck colors.
- 1 Secluded Courtyard | land | Supports the deck's creature base with colored mana.
- 1 Unclaimed Territory | land | Supports the deck's creature base with colored mana.
- 1 Sol Ring | ramp | Efficient early mana acceleration.
- 1 Arcane Signet | ramp | Reliable colored mana acceleration.
- 1 Fellwar Stone | ramp | Low-cost mana acceleration.
- 1 Chromatic Lantern | ramp | Fixes colors while accelerating mana.
- 1 Commander's Sphere | ramp | Provides colored mana acceleration.
- 1 Relic of Legends | ramp | Provides flexible mana acceleration.
- 1 Thought Vessel | ramp | Provides mana acceleration.
- 1 Wayfarer's Bauble | ramp | Develops the mana base early.
- 1 Springleaf Drum | ramp | Turns a creature into early mana.
- 1 Sword of the Animist | ramp | Grows the mana base while supporting creatures.
- 1 Skullclamp | draw | Turns expendable creatures into cards.
- 1 Idol of Oblivion | draw | Repeatable card advantage for the creature-token plan.
- 1 Lembas | draw | Provides a compact card-advantage piece.
- 1 Mask of Memory | draw | Rewards creature combat with card selection.
- 1 Tome of Legends | draw | Builds card advantage around the commander.
- 1 Night's Whisper | draw | Efficient early card flow.
- 1 Painful Truths | draw | Refills the hand efficiently.
- 1 Call of the Ring | draw | Provides continuing card advantage.
- 1 Inspiring Overseer | draw | Adds a creature while drawing a card.
- 1 Wall of Omens | draw | Early board presence that replaces itself.
- 1 Puresteel Paladin | draw | Provides card advantage in the Equipment package.
- 1 Buster Sword | draw | An Equipment-based source of card advantage.
- 1 Foot Chopper | draw | Adds another Equipment-based draw piece.
- 1 Nasty End | draw | Provides card advantage while supporting creature losses.
- 1 Lightning Greaves | interaction | Protects an important creature immediately.
- 1 Swiftfoot Boots | interaction | Protects a key creature while keeping it active.
- 1 Darksteel Plate | interaction | Durable protection for a key creature.
- 1 Champion's Helm | interaction | Protects legendary creatures, especially the commander.
- 1 Gift of Immortality | interaction | Helps preserve an important creature through removal.
- 1 Together Forever | interaction | Helps recover creatures after they die.
- 1 Unbreakable Formation | interaction | Protects the creature board at a pivotal moment.
- 1 Clever Concealment | interaction | Protects the board from opposing disruption.
- 1 Bastion Protector | interaction | Provides protection for the commander.
- 1 Boromir, Warden of the Tower | interaction | Offers a disruptive creature presence.
- 1 Frontline Medic | interaction | Protects the attacking creature force.
- 1 Duty Beyond Death | interaction | Protects creatures when an opponent tries to remove them.
- 1 Bitter Triumph | removal | Flexible single-target answer.
- 1 Claim the Precious | removal | Direct answer to a problematic creature.
- 1 Fatal Push | removal | Efficient creature removal.
- 1 Infernal Grasp | removal | Broad single-target removal.
- 1 Generous Gift | removal | Answers a wide range of permanents.
- 1 Get Lost | removal | Efficient answer to troublesome permanents.
- 1 Stroke of Midnight | removal | Flexible permanent removal.
- 1 Swords to Plowshares | removal | Efficient creature removal.
- 1 Fiend Hunter | removal | Creature-based removal that fits the sacrifice plan.
- 1 Oblation | removal | Versatile answer to a troublesome permanent.
- 1 Arcade Cabinet | synergy | A synergy piece for the creature-focused game plan.
- 1 Gollum the Abandoned | synergy | Supports the deck's sacrifice-focused synergies.
- 1 Gollum, Patient Plotter | synergy | Supports repeated creature-loss synergies.
- 1 Gríma Wormtongue | synergy | Adds a black synergy creature to the board.
- 1 Heirloom Auntie | synergy | Adds another creature-based synergy piece.
- 1 Phantom Train | synergy | Artifact synergy support for the deck.
- 1 Nimble Hobbit | synergy | A low-cost creature synergy piece.
- 1 Bill the Pony | threat | A white creature threat for applying pressure.
- 1 Vengeful Villagers | threat | Adds a creature threat that fits the board-focused plan.
- 1 Angel of Serenity | wincon | A high-impact finisher that also serves as removal.
- 1 Archfiend of Ifnir | wincon | A black finisher for closing longer games.
- 1 Grave Venerations | wincon | A dedicated finishing piece for the creature plan.
- 1 Vraska, Betrayal's Sting | wincon | A powerful finisher that also contributes mana development.
- 1 Witch-king of Angmar | wincon | A resilient black finisher with removal utility.
- 1 Austere Command | wipe | Flexible reset button for difficult boards.
- 1 Fumigate | wipe | Creature-board reset that helps stabilize.
- 1 Vanquish the Horde | wipe | Efficient creature-board reset when needed.

</details>

### 3. artifacts Commander, bracket 4

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 319 names.

Commander: Urza, Lord High Artificer.

Grade: typical, score 0.63, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $3872.70 to buy, $3872.70 the whole deck.

**Summary:** Urza turns the early artifact mana into a fast, resilient board of artifact threats while holding up efficient protection and targeted answers. The deck wins by overwhelming the table with its artifact board, by deploying its dedicated artifact finishers, or through Thassa's Oracle. It gives up some flexibility for a dense artifact core, so its strongest games come from establishing mana quickly and protecting the key turn.

The quality model grades this deck typical against the top lists of the format: the deck makes more mana on turn four than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

- PLAN plan_coherent=partly: The artifact-mana-into-artifact-threats core is consistent and Urza-aligned, but the Thassa's Oracle is a stranded piece with no Demonic Consultation/Tainted Pact/self-mill support, and cards like Into Thin Air, Reality Ripple, Unable to Scream, Frogmite and Foundry Assembler are filler that pulls away from the focused plan.
- PLAN theme_fit=partly: It is a mono-blue Urza artifact Commander deck as asked, with real fast mana and free counters, but the density of low-impact filler (Chrome Steed, Foundry Assembler, Into Thin Air, Transmogrifying Wand) and the absence of the tutors and Urza combo lines expected at bracket 4 leave it short of the requested power tier.
- PLAN useful_as_built=yes: 33 lands plus a heavy fast-mana package supports a mostly cheap blue curve, and there are multiple genuine finishers (Kappa Cannoneer, Cyberdrive Awakener, Mechanized Production, Mirrodin Besieged, Urza tokens) so it can be picked up and played.
- PLAN summary_honest=no: It advertises winning "through Thassa's Oracle" when the list contains no way to empty the library, so a named win route simply does not exist in the deck.
- [INFO] `curve_summary`: average mana value 2.92 over 66 nonland cards

<details><summary>The deck list</summary>

- 20 Island | land | Provides a deep base of blue mana for Urza and the deck's blue spells.
- 1 City of Brass | land | Adds flexible colored mana while entering untapped.
- 1 Mana Confluence | land | Adds flexible colored mana while entering untapped.
- 1 Command Tower | land | Provides reliable access to blue mana.
- 1 Glimmervoid | land | Provides flexible mana for an artifact-heavy board.
- 1 Spire of Industry | land | Provides blue mana in the artifact shell.
- 1 Exotic Orchard | land | Offers flexible colored mana from opponents' lands.
- 1 Reflecting Pool | land | Adds another flexible blue source alongside the multicolor lands.
- 1 Seat of the Synod | land | Is an artifact land that supports the artifact plan.
- 1 Mystic Sanctuary | land | Adds a blue-producing utility land.
- 1 Otawara, Soaring City | land | Provides blue mana from a useful legendary land slot.
- 1 Mistrise Village | land | Adds an untapped blue-producing utility land.
- 1 Urza's Saga | land | Contributes an artifact-focused land slot.
- 1 Power Depot | land | Supports the artifact mana base as an artifact land.
- 1 Chrome Mox | ramp | Supplies fast mana for explosive early turns.
- 1 Lion's Eye Diamond | ramp | Supplies fast mana for high-power artifact turns.
- 1 Mana Vault | ramp | Provides a major burst of fast mana.
- 1 Mox Diamond | ramp | Provides fast mana while helping cast Urza ahead of schedule.
- 1 Mox Opal | ramp | Provides fast mana in the dense artifact shell.
- 1 Lotus Petal | ramp | Provides immediate fast mana when speed matters.
- 1 Sol Ring | ramp | Efficiently accelerates artifact-heavy development.
- 1 Moonsnare Prototype | ramp | Provides a cheap mana source that works with artifacts.
- 1 Springleaf Drum | ramp | Turns small creatures or artifact tokens into mana.
- 1 Metalworker | ramp | Offers substantial artifact-based mana production.
- 1 Krark-Clan Ironworks | ramp | Converts artifacts into mana for powerful turns.
- 1 Grand Architect | ramp | Provides artifact-oriented mana acceleration.
- 1 Urza's Command | ramp | Adds another artifact-focused acceleration spell.
- 1 Rhystic Study | draw | Provides a premier ongoing source of cards.
- 1 Thoughtcast | draw | Efficiently refills the hand in an artifact-heavy deck.
- 1 Sai, Master Thopterist | draw | Provides artifact-linked card advantage.
- 1 Riddlesmith | draw | Helps filter through the deck as artifacts are deployed.
- 1 Forensic Gadgeteer | draw | Adds artifact-connected card advantage.
- 1 Thirst for Knowledge | draw | Provides efficient instant-speed card selection.
- 1 Vedalken Archmage | draw | Rewards the deck's high artifact density with cards.
- 1 Cerebral Download | draw | Adds a compact draw spell to keep resources flowing.
- 1 Era of Innovation | draw | Provides artifact-oriented card advantage.
- 1 Esoteric Duplicator | draw | Supports the deck with artifact-based card flow.
- 1 Braided Net // Braided Quipu | draw | Provides a flexible artifact draw slot.
- 1 Katara, Bending Prodigy | draw | Adds another source of cards to sustain longer games.
- 1 Force of Will | interaction | Provides high-impact protection without demanding mana on the key turn.
- 1 Fierce Guardianship | interaction | Protects the board and important plays once Urza is out.
- 1 An Offer You Can't Refuse | interaction | Offers highly efficient stack interaction.
- 1 Disruption Protocol | interaction | Uses artifacts to support efficient counterplay.
- 1 Metallic Rebuke | interaction | Offers artifact-enabled protection at a low cost.
- 1 Stoic Rebuttal | interaction | Provides dependable artifact-based counterplay.
- 1 Ice Out | interaction | Adds flexible counterspell coverage.
- 1 Reality Ripple | interaction | Provides a low-cost disruptive effect.
- 1 Welding Jar | interaction | Protects important artifacts from opposing answers.
- 1 Padeem, Consul of Innovation | interaction | Helps defend the artifact board.
- 1 Escape Protocol | interaction | Adds defensive flexibility around the artifact plan.
- 1 Aether Spellbomb | removal | Provides cheap artifact-based creature interaction.
- 1 Arcum Dagsson | removal | Supplies removal utility while serving as an artifact tutor.
- 1 Cyber Conversion | removal | Provides efficient targeted removal.
- 1 Resculpt | removal | Answers troublesome targets at instant speed.
- 1 Ravenform | removal | Offers versatile permanent removal.
- 1 Unable to Scream | removal | Provides a low-cost answer to opposing creatures.
- 1 Into Thin Air | removal | Adds flexible targeted removal.
- 1 Kitesail Larcenist | removal | Provides creature-based removal utility.
- 1 Skysovereign, Consul Flagship | removal | Pairs an artifact body with repeated removal value.
- 1 Transmogrifying Wand | removal | Provides reusable artifact-based creature removal.
- 1 Clock of Omens | synergy | Lets the artifact board support powerful tap-and-untap lines.
- 1 Mystic Forge | synergy | Keeps artifact-heavy turns moving from the top of the library.
- 1 Whir of Invention | synergy | Finds a needed artifact directly from the library.
- 1 Arcbound Crusher | threat | Provides an artifact creature threat that benefits from the deck's core plan.
- 1 Chrome Steed | threat | Adds an artifact creature body for board pressure.
- 1 Cyberdrive Awakener | threat | Turns a developed artifact board into a major threat.
- 1 Darksteel Juggernaut | threat | Provides a scalable artifact creature threat.
- 1 Filigree Attendant | threat | Adds an artifact creature that pressures opponents.
- 1 Foundry Assembler | threat | Provides a substantial artifact creature threat.
- 1 Frogmite | threat | Adds an artifact body that benefits from a full board.
- 1 Kappa Cannoneer | threat | Provides a powerful artifact-based attacker.
- 1 Karn, Scion of Urza | threat | Creates a persistent artifact-oriented threat.
- 1 Master Transmuter | threat | Provides a dangerous artifact creature for high-impact board development.
- 1 Phyrexian Metamorph | threat | Adapts into the strongest available artifact or creature threat.
- 1 Traxos, Scourge of Kroog | threat | Provides an efficient large artifact attacker.
- 1 Thassa's Oracle | wincon | Provides a compact alternate win condition.
- 1 Mechanized Production | wincon | Provides an artifact-focused finishing route.
- 1 Mirrodin Besieged | wincon | Provides another artifact-centered route to victory.
- 1 Engineered Explosives | wipe | Provides a scalable reset against problematic boards.
- 1 Desynchronization | wipe | Adds a broad reset for difficult board states.

</details>

### 4. dinosaur tribal, bracket 2

Format: Commander. Theme: dinosaurs. Pool: any_card. Shortlist: 299 names.

Commander: Gishath, Sun's Avatar.

Grade: bad, score 0.03, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $456.16 to buy, $456.16 the whole deck.

**Summary:** This deck ramps into a broad Dinosaur board, uses tribal support to make its creatures more effective, and lets Gishath, Sun's Avatar turn successful attacks into even more Dinosaurs. It wins primarily by overwhelming opponents in combat with large threats and its finisher cards, while retaining enough removal, protection, and wipes to keep the board manageable. It gives up fast, highly precise play for a straightforward creature plan with several expensive cards at the top of the curve.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The curve sits high for the format, and that lowers the grade.

Summary rules claims (F-26): the format play

- JUDGE [true]: "lets Gishath, Sun's Avatar turn successful attacks into even more Dinosaurs". This asserts what the card does: Gishath, Sun's Avatar triggers on dealing combat damage to a player, revealing cards equal to its power and putting the revealed Dinosaurs onto the battlefield tapped and attacking. The loose phrasing ('successful attacks' rather than 'combat damage to a player') is imprecise but substantively correct.
- PLAN plan_coherent=yes: Every slot points at one Dinosaur-tribal plan: mana ramp and cost reducers into big Dinosaurs, tribal lords/anthems and Cavern-style lands, with Gishath converting combat damage into more Dinosaurs.
- PLAN theme_fit=yes: It is a Gishath Dinosaur deck in WRG with a simple attack-and-swing gameplan suited to a new player, and its power sits in casual bracket-2 territory without fast mana, tutors, or combo lines.
- PLAN useful_as_built=yes: 37 lands with strong fixing plus roughly ten ramp pieces and cost reducers support the top-heavy curve, and there are ample ways to close a game through combat and finishers, so it plays fine out of the box.
- PLAN summary_honest=yes: Every claim is backed by cards in the list 
combat-oriented finishers, ramp, tribal payoffs, removal, protection equipment/instants, and board wipes 
and it openly admits the high curve and lack of precision.
- [INFO] `curve_summary`: average mana value 3.85 over 62 nonland cards

<details><summary>The deck list</summary>

- 1 Command Tower | land | A land for the three-color mana base.
- 1 Path of Ancestry | land | A land that fits the Dinosaur tribal plan.
- 1 Cavern of Souls | land | A land for the creature-focused mana base.
- 1 Secluded Courtyard | land | A land that supports the tribal mana base.
- 1 Unclaimed Territory | land | A land that supports the tribal mana base.
- 1 City of Brass | land | A land for broad color access.
- 1 Mana Confluence | land | A land for broad color access.
- 1 Temple Garden | land | A green-white land for the mana base.
- 1 Stomping Ground | land | A red-green land for the mana base.
- 1 Sacred Foundry | land | A red-white land for the mana base.
- 1 Canopy Vista | land | A green-white land for the mana base.
- 1 Cinder Glade | land | A red-green land for the mana base.
- 1 Bountiful Promenade | land | A green-white land for the mana base.
- 1 Spire Garden | land | A red-green land for the mana base.
- 1 Spectator Seating | land | A red-white land for the mana base.
- 1 Karplusan Forest | land | A red-green land for the mana base.
- 1 Brushland | land | A green-white land for the mana base.
- 1 Battlefield Forge | land | A red-white land for the mana base.
- 8 Forest | land | Basic green lands support the deck's green-heavy creature plan.
- 6 Mountain | land | Basic red lands support Dinosaur threats and removal.
- 5 Plains | land | Basic white lands support the deck's protection and support cards.
- 1 Arcane Signet | ramp | Reliable mana ramp for all three colors.
- 1 Fellwar Stone | ramp | Low-cost mana ramp.
- 1 Drover of the Mighty | ramp | Creature-based ramp that belongs in a Dinosaur-focused deck.
- 1 Atzocan Seer | ramp | Creature-based ramp for the Dinosaur plan.
- 1 Ixalli's Lorekeeper | ramp | Creature-based ramp for developing the board.
- 1 Thunderherd Migration | ramp | Dinosaur-themed mana ramp.
- 1 Topiary Stomper | ramp | Dinosaur ramp that also adds to the creature plan.
- 1 Wayward Swordtooth | ramp | Dinosaur ramp that helps reach expensive threats.
- 1 Ranging Raptors | ramp | Dinosaur ramp for building toward larger plays.
- 1 Birds of Paradise | ramp | Early creature-based mana ramp.
- 1 Garruk's Uprising | draw | Card draw for a deck built around large creatures.
- 1 Ripjaw Raptor | draw | A Dinosaur draw source for the creature plan.
- 1 Curious Altisaur | draw | A Dinosaur draw source.
- 1 Earthshaker Dreadmaw | draw | A Dinosaur draw source that remains a substantial creature.
- 1 Return of the Wildspeaker | draw | Flexible card draw for a creature-heavy board.
- 1 Rishkar's Expertise | draw | A powerful draw spell for large-creature boards.
- 1 Runic Armasaur | draw | A Dinosaur draw source.
- 1 Vanquisher's Banner | draw | Tribal card draw for the Dinosaur plan.
- 1 Beast Whisperer | draw | Creature-focused card draw.
- 1 Kutzil, Malamet Exemplar | draw | A creature-based draw option.
- 1 Vaultborn Tyrant | draw | A Dinosaur draw source and imposing creature.
- 1 Heroic Intervention | interaction | Protection interaction for the creature board.
- 1 Lightning Greaves | interaction | Equipment-based protection interaction.
- 1 Swiftfoot Boots | interaction | Equipment-based protection interaction.
- 1 Temple Altisaur | interaction | A Dinosaur interaction piece for protecting the board.
- 1 Steely Resolve | interaction | Tribal protection interaction.
- 1 Legolas's Quick Reflexes | interaction | Low-cost protective interaction.
- 1 Itzquinth, Firstborn of Gishath | removal | Low-cost Dinosaur removal.
- 1 Needletooth Raptor | removal | Dinosaur-based removal.
- 1 Raging Regisaur | removal | A Dinosaur that contributes removal.
- 1 Ravenous Sailback | removal | Dinosaur-based removal for troublesome permanents.
- 1 Thrashing Brontodon | removal | A Dinosaur removal option.
- 1 Triumphant Chomp | removal | Dinosaur-themed removal.
- 1 Trumpeting Carnosaur | removal | A Dinosaur removal option that remains a threat.
- 1 Blasphemous Act | wipe | A broad board wipe when creatures get out of hand.
- 1 Raging Swordtooth | wipe | A Dinosaur-based board wipe.
- 1 Harsh Mercy | wipe | A tribal-friendly board wipe option.
- 1 Kinjalli's Caller | synergy | Dinosaur synergy that supports the creature curve.
- 1 Otepec Huntmaster | synergy | Dinosaur synergy for developing the board.
- 1 Marauding Raptor | synergy | A Dinosaur synergy piece.
- 1 Herald's Horn | synergy | Tribal support for Dinosaur cards.
- 1 Urza's Incubator | synergy | Tribal support for the expensive Dinosaur plan.
- 1 Descendants' Path | synergy | Tribal synergy that rewards a Dinosaur-heavy deck.
- 1 Icon of Ancestry | synergy | A tribal support artifact for Dinosaurs.
- 1 Shared Animosity | synergy | Combat-focused tribal synergy.
- 1 Kindred Summons | synergy | Tribal synergy that helps assemble a Dinosaur board.
- 1 Hunting Velociraptor | synergy | A Dinosaur synergy creature.
- 1 Pantlaza, Sun-Favored | threat | A Dinosaur threat that rewards the creature-focused plan.
- 1 Regisaur Alpha | threat | A Dinosaur threat for applying pressure.
- 1 Quartzwood Crasher | threat | A substantial Dinosaur threat.
- 1 Carnage Tyrant | threat | A large Dinosaur threat.
- 1 Ghalta and Mavren | threat | A legendary Dinosaur threat.
- 1 Etali, Primal Storm | threat | A high-impact Dinosaur threat and finisher.
- 1 Pterodon Knight | threat | A lower-cost threat that helps fill out the curve.
- 1 Majestic Heliopterus | threat | A Dinosaur threat for the combat plan.
- 1 Nurturing Bristleback | threat | A Dinosaur threat for the creature board.
- 1 Sun-Blessed Mount | threat | A Dinosaur threat for attacking opponents.
- 1 Shifting Ceratops | threat | A Dinosaur threat marked as a finisher.
- 1 Goring Ceratops | threat | A large Dinosaur threat marked as a finisher.
- 1 Bonehoard Dracosaur | wincon | A Dinosaur finisher for closing games.
- 1 Dinosaurs on a Spaceship | wincon | A Dinosaur finisher for the deck's endgame.
- 1 Collective Inferno | wincon | A finisher that complements the creature-heavy plan.

</details>

### 5. blink Commander, owned first

Format: Commander. Theme: blink. Pool: owned_first. Shortlist: 169 names.

Commander: Gilraen, Dúnedain Protector.

Grade: typical, score 0.41, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $184.17 the whole deck.

**Summary:** Gilraen, Dúnedain Protector keeps a creature-heavy board moving through blink play, repeatedly leaning on creatures that supply draw, removal, and interaction. The deck develops through a substantial artifact ramp package, protects Gilraen and its key permanents, and uses broad reset buttons when the table gets out of hand. Angel of Serenity is chief among its finishers. The tradeoff is that the list is deliberately reactive and board-focused, with its closing power concentrated in a small finishing package rather than a fast combo.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The land count sits above the norm of the format, and that lowers the grade.

This deck plan aims for 2 or more finishers, the cards that win the game, and the deck holds 1.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The removal/draw/protection shell is coherent enough, but the pieces pull apart: only a handful of actual blink effects support the stated blink plan, while a large equipment/artifact-ramp package and four one-sided-less board wipes work against the creature-heavy board the summary describes.
- PLAN theme_fit=partly: It is a mono-white Commander deck led by a blink-flavored commander at a reasonable bracket-3 power, but the list reads as a generic white goodstuff/protection/equipment pile with only a few true blink enablers and ETB payoffs, so the requested blink theme is only lightly served.
- PLAN useful_as_built=yes: Mono-white with 39 lands and a low curve, it casts its spells reliably and can grind out wins through creatures, removal, and Angel of Serenity, even though the win route is slow and the wipes sometimes fight its own board.
- PLAN summary_honest=partly: It correctly flags the reactive, board-focused nature and the thin single-finisher package, but it oversells "blink play" and "creatures that supply draw, removal, and interaction" when most of the draw comes from artifacts and only about four cards actually flicker anything.
- [INFO] `curve_summary`: average mana value 2.92 over 61 nonland cards
- [WARN] `profile_off_band`: the ramp count is 14, and bracket 3 wants 8 to 13
- [WARN] `finisher_short`: the finisher count is 1, and this deck plan wants 2 or more (Angel of Serenity)
- [INFO] `cards_trimmed`: the list was 1 card over, so the builder cut this card: Dusk // Dawn

<details><summary>The deck list</summary>

- 30 Plains | land | Provides a stable white mana base for the deck.
- 1 Command Tower | land | Provides flexible mana from an untapped land slot.
- 1 City of Brass | land | Provides flexible colored mana.
- 1 Exotic Orchard | land | Provides flexible colored mana.
- 1 Marsh Flats | land | Helps access Plains while supporting the land count.
- 1 Fabled Passage | land | Helps access Plains while supporting the land count.
- 1 Ash Barrens | land | Helps convert a land slot into reliable white mana.
- 1 Plaza of Heroes | land | Provides a utility land for the legendary commander deck.
- 1 Minas Tirith | land | Provides a white utility land.
- 1 Sol Ring | ramp | Provides efficient early mana acceleration.
- 1 Arcane Signet | ramp | Provides reliable mana acceleration.
- 1 Springleaf Drum | ramp | Provides low-cost mana acceleration.
- 1 Fellwar Stone | ramp | Provides additional mana acceleration.
- 1 Wayfarer's Bauble | ramp | Builds mana development from an early play.
- 1 Bender's Waterskin | ramp | Adds another mana-development piece.
- 1 Blitzball | ramp | Adds mana acceleration while remaining within the artifact package.
- 1 Coin of Mastery | ramp | Provides additional mana development.
- 1 Commander's Sphere | ramp | Provides mana acceleration and a flexible late-game slot.
- 1 Chromatic Lantern | ramp | Provides mana acceleration and fixing support.
- 1 Relic of Legends | ramp | Provides mana acceleration for the creature-heavy deck.
- 1 Sword of the Animist | ramp | Provides repeatable mana development from the equipment package.
- 1 Giada, Font of Hope | ramp | Provides creature-based mana acceleration.
- 1 Wickersmith's Tools | ramp | Adds another artifact mana-development piece.
- 1 Adventurer's Airship | draw | Supplies a draw-focused permanent for the deck.
- 1 Buster Sword | draw | Supports card flow from the equipment package.
- 1 Champions of Minas Tirith | draw | Adds a creature-based draw piece to blink around.
- 1 Crown of Gondor | draw | Supplies draw support from an artifact slot.
- 1 Diary of Dreams | draw | Adds a dedicated draw permanent.
- 1 Energybending | draw | Provides a compact draw spell.
- 1 Exemplar of Light | draw | Adds creature-based card flow to the deck.
- 1 Faramir, Field Commander | draw | Supplies a legendary creature draw piece.
- 1 Idol of Oblivion | draw | Provides an efficient artifact draw slot.
- 1 Inspiring Overseer | draw | Adds a creature-based draw piece for blink play.
- 1 Lembas | draw | Provides card flow from a low-commitment artifact.
- 1 Mask of Memory | draw | Supports card flow through the equipment package.
- 1 Puresteel Paladin | draw | Adds creature-based draw support for the artifact suite.
- 1 Wall of Omens | draw | Provides a creature-based draw piece that fits the blink plan.
- 1 Angel of Condemnation | removal | Provides creature-based removal that fits blink play.
- 1 Angel of Sanctions | removal | Provides an Angel removal threat for the creature suite.
- 1 Crib Swap | removal | Provides flexible instant-speed removal.
- 1 Dispatch | removal | Provides efficient removal for the artifact-heavy build.
- 1 Fiend Hunter | removal | Provides creature-based removal for blink play.
- 1 Generous Gift | removal | Answers a broad range of opposing permanents.
- 1 Get Lost | removal | Provides efficient flexible removal.
- 1 Journey to Nowhere | removal | Adds a permanent-based removal option.
- 1 Palace Jailer | removal | Provides creature-based removal for the blink plan.
- 1 Swords to Plowshares | removal | Provides efficient creature removal.
- 1 Stroke of Midnight | removal | Provides flexible instant-speed removal.
- 1 Westfold Rider | removal | Adds another creature-based removal piece.
- 1 Bastion Protector | interaction | Helps protect the commander.
- 1 Boromir, Warden of the Tower | interaction | Provides creature-based interaction.
- 1 Bronze Guardian | interaction | Protects the artifact portion of the deck.
- 1 Champion's Helm | interaction | Supports protection for the commander.
- 1 Clever Concealment | interaction | Provides a protective interaction spell.
- 1 Darksteel Plate | interaction | Provides a durable protection piece.
- 1 Gift of Immortality | interaction | Adds a protective Aura for key creatures.
- 1 Lightning Greaves | interaction | Provides efficient creature protection.
- 1 Reprieve | interaction | Provides flexible stack interaction.
- 1 Slip On the Ring | interaction | Provides blink-oriented interaction.
- 1 Swiftfoot Boots | interaction | Provides another commander-protection equipment.
- 1 Together Forever | interaction | Provides resilient creature interaction.
- 1 Ennis, Debate Moderator | synergy | Supports the deck's blink-focused creature plan.
- 1 Flickerwisp | synergy | Directly supports repeated blink play.
- 1 Jocasta, Automaton Avenger | synergy | Adds another permanent aligned with the blink plan.
- 1 Personify | synergy | Provides a spell slot aligned with the blink plan.
- 1 Austere Command | wipe | Provides a flexible reset when the board gets ahead.
- 1 Fumigate | wipe | Provides a clean creature-board reset.
- 1 Martial Coup | wipe | Provides a sweeping reset that can leave a board presence.
- 1 Vanquish the Horde | wipe | Provides another creature-board reset.
- 1 Angel of Serenity | wincon | Serves as the deck's marked finisher and win condition.

</details>

### 6. Modern tempo, tournament

Format: Modern. Theme: tempo. Pool: any_card. Shortlist: 167 names.

Grade: typical, score 0.62, model 20260914T154223Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $407.06 to buy, $407.06 the whole deck.

**Summary:** This blue-red tempo deck establishes pressure with a compact creature suite, then uses cheap removal and interaction to preserve that lead while cantrips keep the hand moving. It wins by maintaining steady pressure while denying the opponent room to stabilize. The tradeoff is a lean, low-end curve that gives up broader late-game power and more specialized answers.

The quality model grades this deck typical against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade. The curve sits low for the format, and that raises the grade.

- PLAN plan_coherent=yes: Cheap evasive threats, eight one-mana removal spells, light counterspells, and cantrips plus Bauble all pull in the same tempo direction with a consistent one-to-two mana curve.
- PLAN theme_fit=partly: It is a legal blue-red Modern tempo deck, but the requested Delver shell has no Delver of Secrets or comparable spell-count payoff finisher, so it reads as a generic Izzet tempo list rather than the named archetype.
- PLAN useful_as_built=yes: 24 lands with four duals, four fastlands, and two painlands support the low curve, and twelve evasive threats backed by eight burn spells give clear ways to close a game.
- PLAN summary_honest=yes: Every claim — compact creature suite, cheap removal and interaction, cantrips, and a lean low curve giving up late-game power — matches the list as written, with no invented payoffs.
- [INFO] `curve_summary`: average mana value 1.28 over 36 nonland cards

<details><summary>The deck list</summary>

- 8 Island | land | Basic land slot supporting the blue-red mana base.
- 6 Mountain | land | Basic land slot supporting the blue-red mana base.
- 4 Steam Vents | land | Dual land slot for the two-color mana base.
- 4 Spirebluff Canal | land | Dual land slot that keeps the mana base focused on early plays.
- 2 Shivan Reef | land | Additional dual-color land slots for consistent casting.
- 2 Consider | draw | Low-cost card selection for keeping spells flowing.
- 2 Preordain | draw | Early card selection that helps find the needed tempo piece.
- 2 Opt | draw | Low-cost card selection that fits the reactive plan.
- 4 Mishra's Bauble | synergy | Low-cost setup piece supporting the spell-dense tempo plan.
- 2 Counterspell | interaction | Reliable interaction for protecting a lead.
- 2 Spell Pierce | interaction | Cheap interaction that supports an early-pressure plan.
- 2 Force of Negation | interaction | Interaction that helps the deck stay reactive while applying pressure.
- 4 Lightning Bolt | removal | Efficient removal for clearing opposing pressure.
- 4 Galvanic Discharge | removal | Additional efficient removal for keeping the battlefield clear.
- 4 Ragavan, Nimble Pilferer | threat | An early threat that begins the deck's pressure plan.
- 4 Ledger Shredder | threat | A core creature threat for maintaining pressure through the middle turns.
- 4 Faerie Mastermind | threat | A compact threat that fits the deck's reactive posture.

</details>

### 7. Modern burn, casual

Format: Modern. Theme: burn. Pool: any_card. Shortlist: 290 names.

Grade: baseline, score 0.41, model 20260914T154223Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $18.34 to buy, $18.34 the whole deck.

**Summary:** This mono-red burn deck applies pressure with efficient removal, a dense package of spell-focused creatures, and a steady stream of attackers. It wins by combining burn pressure with combat damage, using draw to keep action flowing into the middle turns. It gives up broad answers and matchup-specific tools for a focused, consistent red game plan.

The quality model grades this deck at the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. The color sources cover the pips, and that raises the grade.

- PLAN plan_coherent=no: The pinger/spell-density package (Thermo-Alchemist, Firebrand Archer, Ancestral Anger, Expedite) wants a deck full of cheap spells, but the list is actually built around twelve-plus four-mana creatures (Akki Lavarunner, Fuming Effigy, Hazoret, Ashcloud Phoenix) with only 16 nonland spells, so the two halves undercut each other.
- PLAN theme_fit=no: The request was a mono-red burn deck, and while the colors are right, six burn spells amid twenty-two creatures and eight pump-cantrips makes this a creature aggro deck rather than burn.
- PLAN useful_as_built=partly: The mana base is clean and the deck can attack and win, but 24 lands against a top curve of four with only 16 spells makes it flood-prone and the cantrip-pump slots contribute little, so it plays sloppily though it does function.
- PLAN summary_honest=partly: The creature and removal descriptions match, but "burn pressure" overstates a deck with only six actual damage spells, and calling Ancestral Anger and Expedite a "steady stream of draw" oversells two cantrip-pump effects.
- [INFO] `curve_summary`: average mana value 2.50 over 36 nonland cards

<details><summary>The deck list</summary>

- 24 Mountain | land | Reliable red mana for the deck's consistently red costs.
- 4 Ancestral Anger | draw | Provides draw while keeping the spell count high.
- 2 Expedite | draw | Adds inexpensive draw to keep the burn plan moving.
- 2 Chandra, Dressed to Kill | ramp | Supplies the deck's ramp component while staying on plan.
- 4 Lightning Bolt | removal | Efficient removal for clearing blockers or pressuring life totals.
- 2 Burst Lightning | removal | Additional flexible removal for the burn shell.
- 4 Thermo-Alchemist | synergy | A spell-focused synergy piece that rewards the deck's heavy instant and sorcery count.
- 4 Firebrand Archer | synergy | Supports the spell-heavy burn strategy as a synergy threat.
- 4 Akki Lavarunner // Tok-Tok, Volcano Born | threat | An early threat that contributes combat pressure.
- 4 Ashcloud Phoenix | threat | A substantial threat for continuing the deck's offensive plan.
- 4 Fuming Effigy | threat | A red threat that broadens the creature pressure package.
- 2 Hazoret the Fervent | threat | A high-impact threat to close games after early burn pressure.

</details>

### 8. Modern lifegain, FNM

Format: Modern. Theme: lifegain. Pool: any_card. Shortlist: 301 names.

Grade: bad, score 0.28, model 20260914T154223Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $354.30 to buy, $354.30 the whole deck.

**Summary:** This white-black lifegain deck establishes early creatures, turns repeated life gain into increasingly serious combat pressure, and uses Vito, Thorn of the Dusk Rose to make that plan harder to race. Solitude and The Wandering Emperor clear away opposing threats while the draw package helps keep the board stocked. It gives up some speed and relies on its creatures surviving long enough for the lifegain payoffs to take over.

The quality model grades this deck below the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade. Many of the cards appear in no top list, and that lowers the grade.

- PLAN plan_coherent=yes: Nearly every card either gains life (Soul Warden, Attended Healer, Dawn of Hope, Twinblade Paladin) or converts life gain into cards, counters, or damage (Enduring Innocence, Bloodbond Vampire, Archangel of Thune, Vito), with only Altar of the Pantheon as a marginal inclusion.
- PLAN theme_fit=yes: It is exactly a white-black Modern lifegain deck at a casual-competitive FNM level, with no other constraint named by the requester left unmet.
- PLAN useful_as_built=partly: The deck is playable and has clear win conditions, but 24 lands split 15 white / 15 black sources is thin for the many double-pip costs (Solitude WW, Enduring Innocence WW, Bloodbond Vampire BB, The Wandering Emperor WW, Archangel WW) and the curve is heavily clustered at four mana.
- PLAN summary_honest=yes: placeholder
- [INFO] `curve_summary`: average mana value 3.44 over 36 nonland cards
- [INFO] `mana_pass`: the builder moved 2 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 9 Plains | land | Provides a reliable white mana base for the deck's creature core and white removal.
- 9 Swamp | land | Provides black mana for Vito, Thorn of the Dusk Rose while keeping the mana base stable.
- 4 Godless Shrine | land | Supplies either deck color without adding to the tapped-land total.
- 2 Caves of Koilos | land | Adds flexible white-black mana for a low-cost dual land slot.
- 4 Enduring Innocence | draw | Keeps cards flowing as the deck develops its board.
- 2 Dawn of Hope | draw | Provides a lifegain-focused source of continued card advantage.
- 2 Altar of the Pantheon | ramp | Accelerates mana while supporting both colors.
- 4 Solitude | removal | Provides efficient creature removal without diluting the lifegain plan.
- 2 The Wandering Emperor | removal | Adds flexible removal that also contributes to board presence.
- 4 Soul Warden | synergy | Starts the lifegain engine early and rewards either player building a board.
- 4 Vito, Thorn of the Dusk Rose | synergy | Turns the deck's lifegain plan into direct pressure on the opponent.
- 4 Attended Healer | threat | Provides a lifegain-centered threat that helps build a wider battlefield.
- 4 Bloodbond Vampire | threat | Becomes a substantial attacker as the lifegain engine turns on.
- 4 Twinblade Paladin | threat | Pressures opponents through the deck's repeated lifegain effects.
- 2 Archangel of Thune | threat | Serves as a powerful top-end threat for a battlefield built around gaining life.

</details>

### 9. Standard midrange, FNM

Format: Standard. Theme: midrange. Pool: any_card. Shortlist: 164 names.

Grade: bad, score 0.40, model 20260914T154223Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $76.42 to buy, $76.42 the whole deck.

**Summary:** This black-green midrange deck develops its mana, establishes a steady creature board, and uses efficient removal to keep opposing threats from taking over. Insidious Roots and Corrupted Conviction support the creature-heavy core, while Darkstar Augur and Phyrexian Arena help maintain resources through longer games. It wins by applying sustained pressure with its creatures, chief among them Goldvein Hydra and Vein Ripper, but gives up some speed against decks built to end the game before its larger threats matter.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The cards pair in ways the top lists do not, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core is a coherent BG midrange shell (ramp, removal, card draw, escalating creature threats), but the Insidious Roots package is grafted on without the graveyard-exit enablers it needs, and Corrupted Conviction's sacrifice plan cuts against a board-building midrange gameplan aside from a light pairing with Vein Ripper.
- PLAN theme_fit=yes: It is a 60-card Standard black-green midrange list with creatures, removal, and card advantage at a plausible FNM power level, exactly what was asked.
- PLAN useful_as_built=partly: The 24-land dual-heavy mana base and curve support the deck fine and it has plenty of ways to win, but roughly four slots (Insidious Roots) are close to blank as built, which meaningfully thins the functional deck.
- PLAN summary_honest=partly: The summary accurately describes the mana, removal, card-draw, and finishers, and honestly names the deck's weakness to fast starts, but it asserts that Insidious Roots supports the creature core when nothing in the list returns or exiles creature cards from the graveyard to turn it on.
- [INFO] `curve_summary`: average mana value 2.39 over 36 nonland cards

<details><summary>The deck list</summary>

- 4 Overgrown Tomb | land | A black-green dual land that anchors the mana base.
- 4 Blooming Marsh | land | Adds black-green coverage while keeping the mana base concentrated in its two colors.
- 4 Deathcap Glade | land | Provides another set of black-green lands for consistent casting.
- 6 Forest | land | Reliable green mana for the deck's green spells.
- 6 Swamp | land | Reliable black mana for the deck's black spells.
- 2 Llanowar Elves | ramp | Early mana acceleration that helps deploy the deck's larger creatures.
- 4 Darkstar Augur | draw | A creature-based draw option that keeps the deck advancing its board and hand.
- 2 Phyrexian Arena | draw | A dedicated draw engine for longer midrange games.
- 4 Bitter Triumph | removal | Efficient removal for clearing opposing threats.
- 2 Maelstrom Pulse | removal | Flexible removal that answers troublesome permanents.
- 4 Insidious Roots | synergy | A core support piece for the deck's creature-focused plan.
- 4 Corrupted Conviction | synergy | Supports the creature plan while helping convert expendable resources into cards.
- 4 Goldvein Hydra | threat | A scalable creature threat that rewards developing mana.
- 4 Thrashing Brontodon | threat | A sturdy creature that applies pressure in the middle turns.
- 4 Chomping Changeling | threat | A creature threat that adds to the deck's board presence.
- 2 Vein Ripper | threat | A high-impact top-end creature for closing games.

</details>

### 10. Standard aggro, tournament

Format: Standard. Theme: aggro. Pool: any_card. Shortlist: 294 names.

Grade: baseline, score 0.43, model 20260914T154223Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $127.86 to buy, $127.86 the whole deck.

**Summary:** This red-white aggro deck aims to establish pressure early with a focused creature suite, then clear obstacles with removal and protect its tempo with interaction. Its draw cards help sustain that pressure when the first wave is answered, while Warleader's Call supports the attacking plan. It wins by keeping opponents under consistent combat pressure, giving up broader late-game flexibility for a streamlined, proactive approach.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The cards pair in ways the top lists do not, and that lowers the grade. The color sources cover the pips, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: Everything nominally points at attacking with an anthem and removal backup, but the creature suite is three different four-drops with only one two-drop, so the beatdown clock and the "early pressure" plan pull against the actual curve.
- PLAN theme_fit=partly: It is a Standard red-white deck built around creatures and combat with an attack-triggered anthem, but its top-heavy curve and twelve removal/interaction slots make it more of a midrange list than the aggro deck requested.
- PLAN useful_as_built=yes: Twenty-four lands with eight duals plus Verge and Arena support both pips reliably, and sixteen creatures backed by removal and an anthem give a real, playable route to winning through combat even if it is slower than intended.
- PLAN summary_honest=partly: It correctly names the removal, draw pieces and Warleader's Call, but the claim of establishing pressure early is not supported by a list whose creatures are twelve four-mana bodies and four two-drops.
- [INFO] `curve_summary`: average mana value 2.83 over 36 nonland cards

<details><summary>The deck list</summary>

- 4 Plains | land | Reliable white mana for the deck’s early plays.
- 4 Mountain | land | Reliable red mana for the deck’s early plays.
- 4 Sacred Foundry | land | Supplies both deck colors without using a tapped-land slot.
- 4 Inspiring Vantage | land | Provides red-white mana for an aggressive opening.
- 4 Sunbillow Verge | land | Adds another red-white land option while supporting the mana base.
- 4 Rumble Arena | land | Rounds out the land count with a red-white option.
- 4 Fugitive Codebreaker | draw | Provides draw while fitting the deck’s red aggressive plan.
- 2 Caretaker's Talent | draw | Supplies additional draw support for longer games.
- 4 Boros Charm | interaction | Provides efficient red-white interaction.
- 2 Sheltered by Ghosts | interaction | Adds white interaction while retaining an aggressive focus.
- 4 Harsh Annotation | removal | Forms a compact package of red removal.
- 4 Case of the Gateway Express | removal | Adds white removal to answer opposing cards.
- 4 Warleader's Call | synergy | A red-white synergy piece that supports the aggressive game plan.
- 4 Diversion Specialist | threat | A low-end threat for applying early pressure.
- 4 Dragonback Lancer | threat | A red-white threat that helps maintain battlefield pressure.
- 4 Frilled Sparkshooter | threat | Completes the threat package with a red aggressive creature.

</details>

### 11. Commander with a locked card

Format: Commander. Theme: sacrifice. Pool: any_card. Shortlist: 299 names.

Commander: Karlov of the Ghost Council.

Grade: baseline, score 0.08, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $842.44 to buy, $842.44 the whole deck.

**Summary:** Karlov of the Ghost Council leads a low-curve Orzhov sacrifice deck that builds resources from expendable creatures, protects its important pieces, and turns repeated creature losses into pressure. Karlov grows from the deck’s life-focused synergies and provides a dependable creature-control outlet, while sacrifice payoffs such as Bastion of Remembrance and Relic Vial help close games. The deck gives up some raw speed to keep a broad creature package, but it has strong staying power through recursion, card flow, and reset buttons.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. Few lands enter tapped, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every slot is a sacrifice outlet, a token/body source, a death payoff, or recursion, and the removal and protection suites (Grave Pact, Dictate of Erebos, Cartel Aristocrat, Chthonian Nightmare) all feed the same aristocrats plan the commander wants.
- PLAN theme_fit=yes: It is a Commander deck in Orzhov led by Karlov of the Ghost Council built squarely around sacrificing creatures, exactly what was asked.
- PLAN useful_as_built=yes: 35 lands with a strong Orzhov dual/fetch base plus Sol Ring and altars support a very low curve, and there are multiple redundant win paths (Bastion/Meathook drains, Blasting Station, Karlov and token beatdown) so it plays out fine as written.
- PLAN summary_honest=yes: Each claim maps to real cards — lifegain drains for Karlov (Elas il-Kor, Bastion, Vito's Inquisitor, Ayli), protection (Flare of Fortitude, Gift of Doom, Dark Privilege), recursion (Victimize, Liesa, Ratadrabik, Nightmare Shepherd), card flow (Deadly Dispute, Village Rites, Smothering Abomination) and the three wipes — with only the mild overstatement of Relic Vial as a closer.
- [INFO] `curve_summary`: average mana value 3.16 over 63 nonland cards

<details><summary>The deck list</summary>

- 11 Plains | land | Basic white source for casting Karlov and the deck’s white spells.
- 12 Swamp | land | Basic black source for the sacrifice engine and black spells.
- 1 Command Tower | land | Reliable untapped fixing for both deck colors.
- 1 Godless Shrine | land | Dual-color land that supports both halves of the deck.
- 1 Scrubland | land | Untapped dual-color source for early development.
- 1 Caves of Koilos | land | Flexible white and black mana source.
- 1 City of Brass | land | Fixes either color when needed.
- 1 Mana Confluence | land | Flexible source for either deck color.
- 1 Concealed Courtyard | land | Early dual-color source for the low curve.
- 1 Brightclimb Pathway // Grimclimb Pathway | land | Flexible land choice for either color.
- 1 Shattered Sanctum | land | Dual-color source that supports the main mana base.
- 1 Vault of Champions | land | Dual-color source for multiplayer games.
- 1 Silent Clearing | land | Dual-color land with late-game utility.
- 1 Marsh Flats | land | Fetch land that finds either basic color.
- 1 Prismatic Vista | land | Fetch land that fixes the basic mana base.
- 1 Sol Ring | ramp | Kept as requested and accelerates the deck’s early development.
- 1 Ashnod's Altar | ramp | Sacrifice-based mana acceleration for explosive turns.
- 1 Phyrexian Altar | ramp | Turns disposable creatures into colored mana.
- 1 Pitiless Plunderer | ramp | Rewards creature sacrifices with additional resources.
- 1 Priest of Forgotten Gods | ramp | Sacrifice outlet that also advances the mana plan.
- 1 Pawn of Ulamog | ramp | Converts dying creatures into mana resources.
- 1 Warren Soultrader | ramp | Supports the sacrifice plan while producing mana.
- 1 Deadly Dispute | ramp | Efficient sacrifice-based resource boost.
- 1 Skullport Merchant | ramp | Provides a sacrifice-friendly mana outlet.
- 1 Culling the Weak | ramp | Explosive mana burst that uses a creature as fuel.
- 1 Corrupted Conviction | draw | Cheap card flow that rewards sacrificing a creature.
- 1 Village Rites | draw | Efficient card flow from a disposable creature.
- 1 Vampiric Rites | draw | Repeatable sacrifice outlet for card flow.
- 1 Disciple of Bolas | draw | Turns a large creature into a substantial refill.
- 1 Baron Bertram Graywater | draw | Creature-based card advantage for the longer game.
- 1 Tevesh Szat, Doom of Fools | draw | Supplies ongoing card advantage and expendable bodies.
- 1 Shadowheart, Dark Justiciar | draw | Converts a sacrificed creature into cards.
- 1 Smothering Abomination | draw | Strong card-flow payoff for repeated sacrifices.
- 1 Lord Skitter's Butcher | draw | Adds card advantage on a sacrifice-focused body.
- 1 Jerren, Corrupted Bishop // Ormendahl, the Corrupter | draw | Adds card flow while fitting the creature-sacrifice plan.
- 1 Nasty End | draw | Low-cost refill that uses a creature as fuel.
- 1 Attrition | removal | Repeatable creature control using disposable creatures.
- 1 Ayli, Eternal Pilgrim | removal | Flexible creature-based answer that fits Karlov’s colors.
- 1 Bone Shards | removal | Cheap removal that can turn a creature sacrifice into value.
- 1 Eaten Alive | removal | Creature answer that rewards having fodder available.
- 1 Blasting Station | removal | Sacrifice outlet that converts creatures into direct removal.
- 1 Dictate of Erebos | removal | Turns the deck’s sacrifices into opposing creature losses.
- 1 Grave Pact | removal | Punishes opponents whenever the deck sacrifices creatures.
- 1 Stronghold Assassin | removal | Creature-based removal outlet for sacrifice turns.
- 1 Yawgmoth, Thran Physician | removal | Versatile sacrifice-based creature control.
- 1 Cartel Aristocrat | interaction | Reliable sacrifice outlet that helps protect key creatures.
- 1 Dark Privilege | interaction | Protects an important creature through sacrifice play.
- 1 Fanatical Devotion | interaction | Repeatable protection that turns spare creatures into value.
- 1 Flare of Fortitude | interaction | Protects the board during pivotal turns.
- 1 Gift of Doom | interaction | Defensive protection for a key permanent.
- 1 Eldrazi Monument | interaction | Board protection for a creature-heavy sacrifice strategy.
- 1 Meathook Massacre II | interaction | Interactive permanent that supports the sacrifice game plan.
- 1 Nightmare Shepherd | interaction | Protects value from creature losses and supports resilience.
- 1 Viscera Seer | synergy | Cheap repeatable sacrifice outlet that smooths draws.
- 1 Carrion Feeder | synergy | Efficient sacrifice outlet and growing creature threat.
- 1 Bartolomé del Presidio | synergy | Low-cost sacrifice outlet for converting bodies into pressure.
- 1 Elas il-Kor, Sadistic Pilgrim | synergy | Life-focused sacrifice payoff that works closely with Karlov.
- 1 Fleshtaker | synergy | Rewards the deck for recurring creature sacrifices.
- 1 Chthonian Nightmare | synergy | Supports recurring creature value through sacrifice play.
- 1 Victimize | synergy | Turns expendable creatures into a powerful recursion play.
- 1 Mondrak, Glory Dominus | threat | High-impact creature that strengthens the deck’s board-building plan.
- 1 Felisa, Fang of Silverquill | threat | Threat that benefits from Karlov’s counter-focused growth.
- 1 Ratadrabik of Urborg | threat | Resilient legendary-creature threat for extended games.
- 1 Liesa, Forgotten Archangel | threat | Powerful flying threat that supports a creature-focused plan.
- 1 Ghoulcaller Gisa | threat | Creates a substantial board presence from creature resources.
- 1 Sadistic Hypnotist | threat | Major sacrifice payoff that pressures opponents’ resources.
- 1 Requiem Angel | threat | Evasive finisher that reinforces a creature-heavy board.
- 1 Corpse Harvester | threat | Creature-based threat that supports the deck’s resource engine.
- 1 Basri's Lieutenant | threat | Counter-focused threat that pairs well with Karlov.
- 1 Flesh-Eater Imp | threat | Evasive creature threat that can capitalize on sacrifice turns.
- 1 Prowling Geistcatcher | threat | Threat that helps preserve creature value through sacrifices.
- 1 Vito's Inquisitor | threat | Early creature threat that supports the deck’s main colors.
- 1 Bastion of Remembrance | wincon | Primary sacrifice payoff and a direct route to closing the game.
- 1 Relic Vial | wincon | Sacrifice-focused finisher that turns the deck’s engine into a win.
- 1 Sephiroth, Fabled SOLDIER // Sephiroth, One-Winged Angel | wincon | High-impact finisher for ending games after the engine is established.
- 1 Austere Command | wipe | Flexible reset button for unfavorable boards.
- 1 Toxic Deluge | wipe | Efficient board reset when creature pressure gets overwhelming.
- 1 The Meathook Massacre | wipe | Board wipe that also supports the deck’s life-and-sacrifice plan.

</details>

### 12. Commander on a budget

Format: Commander. Theme: tokens. Pool: any_card. Shortlist: 298 names.

Commander: Adeline, Resplendent Cathar.

Grade: bad, score 0.02, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $89.64 to buy, $89.64 the whole deck.

**Summary:** Adeline, Resplendent Cathar leads a white go-wide deck that attacks early, turns those attacks into a growing Human token force, and reinforces that force with token payoffs, creatures, and enchantments. The deck usually wins by overwhelming the table through combat, while Halo Fountain, Luck Bobblehead, and Sword of Body and Mind provide additional closing pressure. It gives up premium mana acceleration and expensive staple cards in favor of a broad, budget-conscious board-building plan.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The deck makes less mana on turn four than the norm, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The large majority of the list is a coherent white go-wide token plan (token makers, anthems, Divine Visitation, Halo Fountain), but a cluster of off-plan filler artifacts sold as "ramp" (Goldvein Pick, Collector's Vault, Currency Converter, Coin of Mastery, Coin/Karn) and a symmetrical wipe like Phyrexian Rebirth that erases your own board pull against the attack-wide plan.
- PLAN theme_fit=yes: It is a mono-white Adeline-led token deck built from cheap, widely-reprinted commons and uncommons with no expensive staples, matching a budget bracket-2 casual build.
- PLAN useful_as_built=yes: 37 mostly-basic lands in a mono-white deck with a low curve and dozens of token producers and anthems means it casts its spells on time and has ample combat and alternate win routes.
- PLAN summary_honest=yes: placeholder
- [INFO] `curve_summary`: average mana value 3.31 over 62 nonland cards
- [INFO] `cards_trimmed`: the list was 1 card over, so the builder cut this card: Hour of Reckoning

<details><summary>The deck list</summary>

- 30 Plains | land | Provides a deep, reliable white mana base for a mono-white token plan.
- 1 Ancient Den | land | A white-producing land option that also fits the artifact portion of the shortlist.
- 1 Castle Ardenvale | land | A utility land that remains part of the deck's white mana base.
- 1 Command Tower | land | An inexpensive untapped color source for the commander’s identity.
- 1 Exotic Orchard | land | A free-to-buy multicolor source that supplies white in this deck.
- 1 Eiganjo, Seat of the Empire | land | A white source with useful utility from the shortlist.
- 1 Mirrex | land | A low-cost utility land that broadens the mana base.
- 1 Windbrisk Heights | land | A budget utility land for a creature-heavy attacking deck.
- 1 Goldvein Pick | ramp | An inexpensive ramp piece for advancing the deck’s mana.
- 1 Collector's Vault | ramp | A cheap ramp artifact that helps develop mana beyond land drops.
- 1 Currency Converter | ramp | A low-cost ramp selection from the shortlist.
- 1 Coin of Mastery | ramp | Provides another budget ramp slot for the deck.
- 1 Druidic Satchel | ramp | An inexpensive ramp artifact for steady resource development.
- 1 Golden Guardian // Gold-Forge Garrison | ramp | A ramp card that also suits a board-focused token strategy.
- 1 Idol of False Gods | ramp | A very inexpensive ramp option from the available cards.
- 1 Karn, Living Legacy | ramp | A ramp planeswalker that supports the deck’s longer games.
- 1 Keeper of the Accord | ramp | A creature-based ramp choice that fits the deck’s creature plan.
- 1 Legion's Landing // Adanto, the First Fort | ramp | A low-cost ramp card that belongs in a token-oriented shell.
- 1 Angelic Sell-Sword | draw | An inexpensive draw creature for maintaining cards in hand.
- 1 Bygone Bishop | draw | A budget draw creature that supports a creature-heavy build.
- 1 Dawn of Hope | draw | A low-cost draw enchantment for token-focused games.
- 1 Faramir, Field Commander | draw | A cheap draw option that fits the deck’s Human-heavy battlefield.
- 1 Glimmer Seeker | draw | A very inexpensive draw creature from the shortlist.
- 1 Idol of Oblivion | draw | A draw artifact that naturally belongs beside token production.
- 1 Nexus of Becoming | draw | A budget draw artifact for replenishing resources.
- 1 Staff of the Storyteller | draw | A low-cost draw artifact that complements the deck’s token focus.
- 1 Wedding Announcement // Wedding Festivity | draw | A draw card well suited to a board-building white strategy.
- 1 Wojek Investigator | draw | An inexpensive creature-based draw selection.
- 1 Search the Premises | draw | A budget draw enchantment for keeping the deck supplied.
- 1 Basri, Tomorrow's Champion | interaction | A low-cost interaction card that remains on-theme with the creature plan.
- 1 Lena, Selfless Champion | interaction | Creature-based interaction that supports a wide battlefield.
- 1 Mage's Attendant | interaction | An inexpensive interactive creature from the shortlist.
- 1 Rootborn Defenses | interaction | A low-cost interaction spell suited to a token board.
- 1 Spirit Bonds | interaction | An interaction enchantment that fits the deck’s token-oriented approach.
- 1 Squad Commander | interaction | A budget interactive creature for the attacking plan.
- 1 Generous Gift | removal | Flexible, inexpensive removal for troublesome permanents.
- 1 Kellan's Lightblades | removal | A very cheap removal spell from the shortlist.
- 1 Skyclave Apparition | removal | Creature-based removal that contributes to the battlefield.
- 1 Stroke of Midnight | removal | A budget removal spell with broad utility.
- 1 Righteous Confluence | removal | A flexible removal option at a low price.
- 1 Banishing Slash | removal | An inexpensive removal spell for the deck’s answer suite.
- 1 Hanged Executioner | removal | A cheap removal creature that fits a creature-focused deck.
- 1 Martial Coup | wipe | A low-cost wipe that also aligns with the deck’s token plan.
- 1 Phyrexian Rebirth | wipe | An inexpensive board wipe for recovering from opposing boards.
- 1 Divine Visitation | synergy | A token synergy enchantment and a particularly strong payoff for Adeline’s attacks.
- 1 Intangible Virtue | synergy | A low-cost token synergy piece for improving a wide board.
- 1 Felidar Retreat | synergy | An on-theme token synergy enchantment for the deck’s board-building plan.
- 1 Worthy Knight | synergy | A cheap Human-and-Knight-adjacent token synergy creature.
- 1 Siege Veteran | synergy | A low-cost creature that supports the deck’s token strategy.
- 1 Rosie Cotton of South Lane | synergy | An inexpensive token synergy creature for growing the board.
- 1 Clarion Spirit | synergy | A very cheap token synergy creature for early development.
- 1 Cat Collector | synergy | A budget token-focused synergy creature.
- 1 Horn of Gondor | synergy | A dedicated token synergy artifact for a wide battlefield.
- 1 Oketra's Monument | synergy | A lasting token synergy payoff that fits the creature-heavy list.
- 1 Ajani's Chosen | threat | A low-cost threat that benefits from the deck’s enchantment selections.
- 1 Archangel Elspeth | threat | A substantial token-oriented threat for board-focused games.
- 1 Archon of Sun's Grace | threat | A cheap threat that pairs well with the deck’s enchantments.
- 1 Attended Healer | threat | An inexpensive creature threat that supports a token-centric battlefield.
- 1 Basri's Lieutenant | threat | A budget Human Knight threat for the deck’s creature plan.
- 1 Cemetery Protector | threat | A versatile creature threat at a modest cost.
- 1 Defiler of Faith | threat | A low-cost white creature threat for maintaining pressure.
- 1 Emeria Angel | threat | A budget token-oriented threat that rewards continued land drops.
- 1 God-Eternal Oketra | threat | A resilient creature threat for sustained board development.
- 1 Hero of Bladehold | threat | An efficient attacking threat that fits Adeline’s combat plan.
- 1 Gideon, Ally of Zendikar | threat | A token-oriented planeswalker threat for wider boards.
- 1 Mite Overseer | threat | An inexpensive white threat that contributes to the creature count.
- 1 Phantom General | threat | A low-cost threat that directly supports a token battlefield.
- 1 Halo Fountain | wincon | A dedicated win condition and finisher for a deck that builds many tokens.
- 1 Luck Bobblehead | wincon | A budget-listed alternate win condition and finisher.
- 1 Sword of Body and Mind | wincon | A finisher-equipped win condition that gives the deck another way to close.

</details>

### 13. owned first, and the commander is not owned

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 230 names.

Commander: Karlov of the Ghost Council.

Grade: baseline, score 0.05, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $140.75 the whole deck.

**Summary:** Karlov of the Ghost Council is the center of a steady Orzhov lifegain plan: gain life repeatedly to grow Karlov, protect it, and use its creature-control pressure to keep attacks clear. The deck backs that plan with efficient removal, board resets, and a broad creature suite, then closes through Karlov's accumulated power or chief among its finishers, Frodo, Sauron's Bane, Grave Venerations, and Lyra Dawnbringer. It gives up explosive multicolor mana and relies on its basics, artifact acceleration, and steady card flow to win longer games.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. Few lands enter tapped, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: Karlov, the protection equipment, and the lifegain payoffs (Angel of Vitality, Light of Promise, Rosie Cotton, Shattered Angel, Exemplar of Light) point one way, but the repeatable-lifegain density is thinner than a Karlov engine wants and three mass wipes plus a wide creature suite pull against each other.
- PLAN theme_fit=yes: It is an Orzhov Commander deck led by the requested Karlov with a lifegain-leaning creature and enchantment suite at a bracket-3-appropriate power level with only ordinary staples.
- PLAN useful_as_built=yes: 36 basics in a two-color deck plus nine ramp pieces support a reasonable curve, and the removal, wipes, card draw, and several combat finishers give a playable game plan out of the box.
- PLAN summary_honest=partly: Most claims check out (basics-only mana, removal, wipes, card flow), but naming Frodo, Grave Venerations, and Lyra as the chief finishers oversells three modest cards, and it does not flag that the wipes hit its own board.
- [WARN] `not_owned`: Karlov of the Ghost Council: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.10 over 63 nonland cards

<details><summary>The deck list</summary>

- 19 Plains | land | Provides reliable untapped white mana for Karlov and the deck's white spells.
- 17 Swamp | land | Provides reliable untapped black mana for Karlov and the deck's black spells.
- 1 Arcane Signet | ramp | Efficiently fixes mana for the two-color plan.
- 1 Bender's Waterskin | ramp | Adds to the deck's mana development.
- 1 Blitzball | ramp | Adds to the deck's mana development.
- 1 Fellwar Stone | ramp | Provides efficient early mana acceleration.
- 1 Hot Dog Cart | ramp | Adds another mana-development piece.
- 1 Inherited Envelope | ramp | Helps advance mana into the deck's stronger plays.
- 1 Lotho, Corrupt Shirriff | ramp | Contributes to mana acceleration while fitting Karlov's colors.
- 1 Sol Ring | ramp | Provides an efficient burst of early mana.
- 1 Thought Vessel | ramp | Adds colorless mana development.
- 1 Wayfarer's Bauble | ramp | Provides early land-based mana development.
- 1 Buster Sword | draw | Supplies a repeatable card-advantage option.
- 1 Call of the Ring | draw | Adds a dedicated source of cards.
- 1 Cirith Ungol Patrol | draw | Provides card advantage from a creature slot.
- 1 Exemplar of Light | draw | Pairs card advantage with a creature body.
- 1 Hero in Training | draw | Adds another creature-based draw piece.
- 1 Idol of Oblivion | draw | Provides a compact source of card advantage.
- 1 Inspiring Overseer | draw | Supplies cards while contributing to the board.
- 1 Lembas | draw | Provides card advantage in an inexpensive artifact slot.
- 1 Mask of Memory | draw | Turns combat into ongoing card advantage.
- 1 Night's Whisper | draw | Efficiently refills the hand.
- 1 Skullclamp | draw | Provides a powerful card-advantage outlet.
- 1 Bastion Protector | interaction | Helps safeguard Karlov from opposing pressure.
- 1 Boromir, Warden of the Tower | interaction | Adds a proactive piece of protection and disruption.
- 1 Champion's Helm | interaction | Helps keep Karlov protected on the battlefield.
- 1 Darksteel Plate | interaction | Provides durable protection for an important creature.
- 1 Lightning Greaves | interaction | Protects Karlov while helping it act immediately.
- 1 Reprieve | interaction | Offers flexible disruption against an opposing spell.
- 1 Swiftfoot Boots | interaction | Provides another efficient way to protect Karlov.
- 1 Unbreakable Formation | interaction | Protects the board through opposing disruption.
- 1 Banishing Light | removal | Answers a troublesome opposing permanent.
- 1 Bitter Triumph | removal | Provides efficient spot removal.
- 1 Crib Swap | removal | Answers an opposing creature at instant speed.
- 1 Dismember | removal | Provides inexpensive creature removal.
- 1 Dispatch | removal | Adds another low-cost removal spell.
- 1 Fatal Push | removal | Efficiently removes early opposing threats.
- 1 Generous Gift | removal | Answers a broad range of opposing permanents.
- 1 Get Lost | removal | Provides flexible permanent removal.
- 1 Infernal Grasp | removal | Adds clean creature removal.
- 1 Angel of Vitality | synergy | Supports the deck's lifegain-focused plan.
- 1 Compassionate Healer | synergy | Contributes to the deck's lifegain synergies.
- 1 Kor Firewalker | synergy | Provides a lifegain-oriented synergy piece.
- 1 Light of Promise | synergy | Builds on the plan of making Karlov large.
- 1 Momo, Playful Pet | synergy | Supports the central lifegain strategy.
- 1 Night Nurse, Healer of Heroes | synergy | Adds another lifegain-focused support card.
- 1 Rosie Cotton of South Lane | synergy | Rewards the deck's creature and lifegain plan.
- 1 Angel of Invention | threat | Provides a meaningful evasive battlefield threat.
- 1 Dawnhand Eulogist | threat | Adds a substantial creature threat.
- 1 Gollum, Silent Slinker // Meager Meal | threat | Provides pressure while fitting the deck's colors.
- 1 Invisible Woman, Sue Storm | threat | Adds another standalone threat to the board.
- 1 Lo and Li, Twin Tutors | threat | Provides a creature threat that advances the board.
- 1 Minwu, White Mage | threat | Adds a white creature threat to the deck's curve.
- 1 Rabaroo Troop | threat | Supplies additional battlefield pressure.
- 1 Reaping Willow | threat | Adds a larger creature presence to the deck.
- 1 Rooftop Percher | threat | Provides another resilient board threat.
- 1 Shattered Angel | threat | Adds an evasive creature threat.
- 1 Sneering Shadewriter | threat | Provides a black creature threat with finishing potential.
- 1 Victory's Herald | threat | Adds a high-impact evasive threat.
- 1 Frodo, Sauron's Bane | wincon | Provides a dedicated finishing route.
- 1 Grave Venerations | wincon | Adds a finisher for closing longer games.
- 1 Lyra Dawnbringer | wincon | Provides an evasive finisher that pressures life totals.
- 1 Austere Command | wipe | Offers a flexible reset when opponents build ahead.
- 1 Fumigate | wipe | Resets creature-heavy boards while supporting the lifegain plan.
- 1 Vanquish the Horde | wipe | Provides an efficient creature-board reset.

</details>

### 14. the user delegates the commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 299 names.

Commander: Atarka, World Render.

Grade: baseline, score 0.05, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $843.73 to buy, $843.73 the whole deck.

**Summary:** This deck ramps through its early turns, uses tribal cost support to deploy Dragons efficiently, and then turns a wide aerial board into devastating attacks with Atarka, World Render. Lathliss, Dragon Queen, Utvara Hellkite, and Wrathful Red Dragon provide especially potent closing pressure, while the removal suite keeps opposing boards from stabilizing. It gives up some resilience to repeated sweepers and relies on its mana development and creature board to convert its powerful top end into a win.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The curve sits high for the format, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Every piece — cost reducers, dragon-tribal payoffs, ramp orbs, and a red/green beatdown top end — pushes the same plan of deploying dragons early and swinging with Atarka-buffed fliers.
- PLAN theme_fit=yes: It is a Commander dragon-tribal deck with a dragon commander chosen by the builder, and its card choices (no infinite combos, a few strong staples like Taiga and Sylvan Library) sit comfortably in bracket 3.
- PLAN useful_as_built=yes: 36 lands with heavy fixing plus ten ramp/cost-reduction pieces support the high curve, and the numerous dragons plus Lathliss/Utvara/Shared Animosity give plenty of ways to close a game.
- PLAN summary_honest=yes: Each named card and mechanic (ramp artifacts, tribal cost reduction, Lathliss/Utvara/Wrathful Red Dragon, removal suite) is actually in the list, and it openly admits the weakness to sweepers and the high curve.
- [INFO] `curve_summary`: average mana value 3.38 over 63 nonland cards
- [INFO] `basics_added`: the list was 1 card short, so the builder added 1 basic land

<details><summary>The deck list</summary>

- 1 Command Tower | land | Reliable red and green mana for the deck.
- 1 City of Brass | land | Flexible access to either deck color.
- 1 Mana Confluence | land | Flexible access to either deck color.
- 1 Cavern of Souls | land | A tribal mana source for the Dragon-heavy creature suite.
- 1 Copperline Gorge | land | An efficient red-green source.
- 1 Cragcrown Pathway // Timbercrown Pathway | land | A flexible choice of red or green mana.
- 1 Fire-Lit Thicket | land | Fixes between red and green mana.
- 1 Karplusan Forest | land | An untapped red-green source.
- 1 Rockfall Vale | land | Provides both deck colors.
- 1 Rootbound Crag | land | Provides both deck colors.
- 1 Spire Garden | land | An untapped multiplayer red-green source.
- 1 Stomping Ground | land | An efficient red-green source.
- 1 Taiga | land | An efficient red-green source.
- 1 Wooded Foothills | land | Finds a needed basic or dual land.
- 1 Game Trail | land | Provides red or green mana.
- 1 Grove of the Burnwillows | land | Provides red or green mana.
- 1 Thornspire Verge | land | Provides both deck colors.
- 1 Exotic Orchard | land | Flexible color fixing in multiplayer games.
- 1 Forbidden Orchard | land | Flexible access to either deck color.
- 1 Secluded Courtyard | land | A tribal source for Dragons.
- 1 Unclaimed Territory | land | A tribal source for Dragons.
- 8 Forest | land | Stable green mana for early development.
- 7 Mountain | land | Stable red mana for Dragons and removal.
- 1 Carnelian Orb of Dragonkind | ramp | Early mana acceleration tailored to the Dragon plan.
- 1 Dragon's Hoard | ramp | A Dragon-focused mana source.
- 1 Dragonstorm Globe | ramp | Fixes mana while advancing development.
- 1 Jade Orb of Dragonkind | ramp | Early green acceleration for the deck.
- 1 Mox Jasper | ramp | A low-cost mana source for the tribal shell.
- 1 Orb of Dragonkind | ramp | Early mana acceleration for Dragons.
- 1 Patchwork Banner | ramp | Provides mana while supporting the tribe.
- 1 Pillar of Origins | ramp | Fixes mana for the chosen creature type.
- 1 Progenitor's Icon | ramp | A tribal mana source that supports creature deployment.
- 1 Scaled Nurturer | ramp | Early creature-based acceleration for Dragons.
- 1 Avaricious Dragon | draw | A Dragon that supplies card flow.
- 1 Beast Whisperer | draw | Turns creature deployment into continued card flow.
- 1 Demand Answers | draw | Low-cost card selection and replenishment.
- 1 Faithless Looting | draw | Efficient early card selection.
- 1 Garruk's Uprising | draw | Card flow that suits the deck's large creatures.
- 1 Guardian Project | draw | Sustained card flow for a creature-heavy deck.
- 1 Harmonize | draw | Straightforward card replenishment.
- 1 Return of the Wildspeaker | draw | Flexible card replenishment for a large-creature board.
- 1 Shamanic Revelation | draw | Refills the hand from a developed creature board.
- 1 Sylvan Library | draw | Early, repeatable card selection.
- 1 Vanquisher's Banner | draw | Tribal card flow for the Dragon plan.
- 1 Heroic Intervention | interaction | Protects the developed board from opposing disruption.
- 1 Legolas's Quick Reflexes | interaction | A low-cost protective interaction spell.
- 1 Lightning Greaves | interaction | Protects a key creature and helps it act promptly.
- 1 Mithril Coat | interaction | Protects an important creature.
- 1 Pyroblast | interaction | Cheap stack interaction.
- 1 Red Elemental Blast | interaction | Cheap stack interaction.
- 1 Swiftfoot Boots | interaction | Protects a key creature while retaining flexibility.
- 1 Veil of Summer | interaction | Low-cost protection against opposing interaction.
- 1 Draconic Roar | removal | Efficient targeted removal that fits the Dragon shell.
- 1 Dragon's Fire | removal | Efficient targeted removal for a Dragon deck.
- 1 Glorybringer | removal | A Dragon threat that also handles opposing creatures.
- 1 Invasion of Tarkir // Defiant Thundermaw | removal | Early removal that leaves a Dragon payoff behind.
- 1 Magmatic Hellkite | removal | A Dragon that contributes targeted removal.
- 1 Opportunistic Dragon | removal | A Dragon-based answer to opposing permanents.
- 1 Piercing Exhale | removal | Low-cost targeted removal.
- 1 Scourge of Valkas | removal | A Dragon payoff that supplies removal pressure.
- 1 Terror of the Peaks | removal | A powerful Dragon that turns creature deployment into removal.
- 1 Breath Weapon | wipe | A compact reset against smaller creature boards.
- 1 Draconic Intervention | wipe | A scalable board reset.
- 1 Thundermaw Hellkite | wipe | A Dragon threat that can clear opposing fliers.
- 1 Acolyte of Bahamut | synergy | Supports deploying the deck's expensive Dragon spells.
- 1 Crucible of Fire | synergy | A dedicated tribal payoff for Dragons.
- 1 Dragonlord's Servant | synergy | Helps the deck deploy Dragons more efficiently.
- 1 Dragonspeaker Shaman | synergy | Helps the deck deploy Dragons more efficiently.
- 1 Herald's Horn | synergy | Tribal support that smooths Dragon development.
- 1 Shared Animosity | synergy | Rewards attacking with a broad Dragon board.
- 1 Urza's Incubator | synergy | A major tribal cost-reduction piece.
- 1 Backdraft Hellkite | threat | An evasive Dragon attacker with useful value potential.
- 1 Caldera Pyremaw | threat | A sizeable Dragon attacker for the combat plan.
- 1 Dragon Broodmother | threat | A Dragon threat that builds the board over time.
- 1 Manaform Hellkite | threat | An evasive Dragon that adds pressure.
- 1 Mirrorwing Dragon | threat | An evasive Dragon threat that rewards protecting the board.
- 1 Roaming Throne | threat | A tribal payoff threat for the chosen creature type.
- 1 Stormbreath Dragon | threat | A hasty evasive attacker.
- 1 Territorial Hellkite | threat | A Dragon that applies immediate combat pressure.
- 1 Thrakkus the Butcher | threat | A Dragon attacker that amplifies the tribal combat plan.
- 1 Thunderbreak Regent | threat | A resilient evasive Dragon threat.
- 1 Thundermane Dragon | threat | An evasive Dragon attacker.
- 1 Verix Bladewing | threat | A flexible Dragon threat for building the board.
- 1 Lathliss, Dragon Queen | wincon | A Dragon payoff that can close games with a developed board.
- 1 Utvara Hellkite | wincon | A top-end Dragon payoff that overwhelms opponents through combat.
- 1 Wrathful Red Dragon | wincon | A Dragon finisher that creates a decisive damage plan.

</details>

### 15. delegated commander, owned first

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 230 names.

Commander: Astarion, the Decadent.

Grade: baseline, score 0.07, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $203.33 the whole deck.

**Summary:** Astarion, the Decadent leads a white-black lifegain deck that develops mana, establishes creatures and Equipment, and uses life gained during the turn to make the Friends choice at end step especially rewarding. The board is supported by efficient answers and reset buttons, then closes through its dedicated victory cards and finisher-class threats, with Astarion's Feed mode providing a way to press an opponent who has already lost life. The deck gives up some speed for a steady creature-based battlefield plan and a commander that costs six mana.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The land count sits above the norm of the format, and that lowers the grade.

Summary rules claims (F-26): the format play

- JUDGE [false]: "Astarion, the Decadent ... a commander that costs six mana.". Astarion, the Decadent has mana cost {2}{W}{B}, a mana value of four, not six.
- JUDGE [true]: "Astarion's Feed mode provides a way to press an opponent who has already lost life.". Astarion's end-step 'Feed' mode targets an opponent who lost life this turn and drains them, so the description of the card's ability is accurate.
- JUDGE [true]: "Life gained during the turn makes the Friends choice at end step rewarding.". Astarion's 'Friends' mode checks whether you gained life this turn to give counters/benefit, so the conditional described matches the card.
- JUDGE [true]: "Astarion, the Decadent leads a white-black deck.". Astarion's color identity is white-black, so it can legally head a white-black Commander deck.
- PLAN plan_coherent=partly: The lifegain core (Angel of Vitality, Rosie Cotton, Light of Promise, Prideful Feastling, Shattered Angel, Victory's Herald, Lyra) is real and feeds Astarion's end-step choice, but a sizable Equipment/card-draw package (Puresteel Paladin, Skullclamp, Mask of Memory, Sword of the Animist, Buster Sword, Aettir and Priwen) and three sweepers sit awkwardly beside a wide creature-based board plan.
- PLAN theme_fit=yes: It is a Commander lifegain deck at a modest, precon-level power with a commander chosen by the builder whose ability directly rewards gaining life, as asked.
- PLAN useful_as_built=yes: 37 lands plus several rocks support a low curve, removal and wraths are plentiful, and the creature suite with lifelinking threats gives clear ways to close, so it plays fine out of the box (if anything slightly land-heavy).
- PLAN summary_honest=partly: Most claims match the list, but "closes through its dedicated victory cards" overstates what Al Bhed Salvagers, Frodo, and Lyra actually are — there is no alternate-win or dedicated finishing engine, just good creatures.
- [WARN] `not_owned`: Astarion, the Decadent: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.10 over 62 nonland cards

<details><summary>The deck list</summary>

- 15 Plains | land | Provides a deep base of white mana for the deck's lifegain-focused spells.
- 14 Swamp | land | Provides reliable black mana for Astarion and the deck's black spells.
- 1 Command Tower | land | Provides either commander color without restriction.
- 1 City of Brass | land | Provides flexible access to both deck colors.
- 1 Exotic Orchard | land | Provides flexible colored mana in multiplayer games.
- 1 Marsh Flats | land | Finds a needed basic land color.
- 1 Fabled Passage | land | Finds a basic land to smooth early mana.
- 1 Grand Coliseum | land | Adds another flexible colored source.
- 1 Spire of Industry | land | Supports colored mana alongside the artifact package.
- 1 Plaza of Heroes | land | Helps cast the deck's legendary cards while producing colored mana.
- 1 Buster Sword | draw | Supplies repeatable card advantage from an Equipment slot.
- 1 Call of the Ring | draw | Provides an ongoing source of cards.
- 1 Idol of Oblivion | draw | Offers efficient card advantage.
- 1 Inspiring Overseer | draw | Adds a creature-based source of card advantage.
- 1 Lembas | draw | Provides a compact card-advantage artifact.
- 1 Mask of Memory | draw | Turns equipped creature attacks into card selection.
- 1 Night's Whisper | draw | Provides inexpensive immediate card advantage.
- 1 Puresteel Paladin | draw | Rewards the deck's Equipment package with cards.
- 1 Skullclamp | draw | Provides efficient card advantage around small creatures.
- 1 Tome of Legends | draw | Uses the commander to keep cards flowing.
- 1 Wall of Omens | draw | Supplies an early defensive body and a card.
- 1 Bastion Protector | interaction | Helps keep Astarion on the battlefield.
- 1 Clever Concealment | interaction | Protects the developed board from opposing disruption.
- 1 Darksteel Plate | interaction | Gives an important creature durable protection.
- 1 Lightning Greaves | interaction | Protects Astarion while letting it contribute immediately.
- 1 Swiftfoot Boots | interaction | Provides another efficient protection piece for key creatures.
- 1 Take Up the Shield | interaction | Protects a creature during a critical exchange.
- 1 Arcane Signet | ramp | Provides efficient fixing for both commander colors.
- 1 Sol Ring | ramp | Accelerates the deck's setup and six-mana commander.
- 1 Fellwar Stone | ramp | Provides early mana acceleration with color flexibility.
- 1 Wayfarer's Bauble | ramp | Develops the mana base with a basic land.
- 1 Sword of the Animist | ramp | Turns creature attacks into continued land development.
- 1 Giada, Font of Hope | ramp | Accelerates the deck's Angel creatures.
- 1 Lotho, Corrupt Shirriff | ramp | Provides mana development on a relevant creature.
- 1 Relic of Legends | ramp | Uses legendary creatures as an additional mana resource.
- 1 White Auracite | ramp | Adds another artifact-based mana source.
- 1 Commander's Sphere | ramp | Fixes mana early and can later become a card.
- 1 Bitter Triumph | removal | Offers flexible single-target removal.
- 1 Generous Gift | removal | Answers a wide range of opposing permanents.
- 1 Get Lost | removal | Provides efficient removal for key opposing permanents.
- 1 Infernal Grasp | removal | Cleanly answers a problematic creature.
- 1 Swords to Plowshares | removal | Provides efficient creature removal.
- 1 Stroke of Midnight | removal | Answers a broad range of permanent types.
- 1 Crib Swap | removal | Provides an additional creature answer.
- 1 Aerith Gainsborough | synergy | Supports the deck's lifegain-centered plan.
- 1 Angel of Vitality | synergy | Builds on the deck's lifegain theme.
- 1 Compassionate Healer | synergy | Adds another creature dedicated to the lifegain plan.
- 1 Light of Promise | synergy | Converts the deck's lifegain focus into creature support.
- 1 Night Nurse, Healer of Heroes | synergy | Provides another on-theme legendary creature.
- 1 Rosie Cotton of South Lane | synergy | Rewards the deck for pursuing its central theme.
- 1 Second Breakfast | synergy | Supports the deck's lifegain-oriented game plan.
- 1 Prideful Feastling | synergy | Adds another dedicated lifegain synergy piece.
- 1 Elixir | synergy | Provides a compact artifact for the deck's synergy package.
- 1 Aettir and Priwen | synergy | Adds an Equipment-based synergy piece for the creature plan.
- 1 Angel of Invention | threat | Provides a substantial evasive creature for the battlefield.
- 1 Bill the Pony | threat | Adds a creature threat that supports board presence.
- 1 Canyon Crawler | threat | Provides a creature threat to pressure opponents.
- 1 Foggy Swamp Hunters | threat | Adds another creature that advances the combat plan.
- 1 Gollum, Silent Slinker // Meager Meal | threat | Provides a relevant threat with additional flexibility.
- 1 Minwu, White Mage | threat | Adds a legendary creature threat to the board.
- 1 Rabaroo Troop | threat | Provides another creature threat for combat pressure.
- 1 Shattered Angel | threat | Adds an evasive creature to the threat suite.
- 1 Victory's Herald | threat | Provides a high-impact evasive threat.
- 1 Rooftop Percher | threat | Adds another creature threat for the board plan.
- 1 Sneering Shadewriter | threat | Provides a finisher-class creature threat.
- 1 Dawnhand Eulogist | threat | Provides another finisher-class creature threat.
- 1 Al Bhed Salvagers | wincon | Serves as one of the deck's dedicated ways to close a game.
- 1 Frodo, Sauron's Bane | wincon | Provides a dedicated alternate route to end the game.
- 1 Lyra Dawnbringer | wincon | Offers a powerful finisher-class route to victory.
- 1 Austere Command | wipe | Provides a flexible reset when opponents establish too much.
- 1 Fumigate | wipe | Resets creature-heavy boards while fitting the deck's theme.
- 1 Vanquish the Horde | wipe | Supplies an efficient broad creature reset.

</details>

### 16. a tight budget, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 231 names.

Commander: Denethor, Ruling Steward.

Grade: baseline, score 0.33, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $13.30 to buy, $45.97 the whole deck.

**Summary:** Denethor, Ruling Steward anchors a creature-resource game: establish mana and card flow, develop a board of expendable creatures and sacrifice-focused pieces, then turn repeated creature losses into pressure. Bastion of Remembrance, Falkenrath Noble, and Zulaport Cutthroat provide the clearest finishing pressure alongside Denethor’s own sacrifice outlet. The deck trades explosive starts and premium individual card power for a patient, board-based plan that can rebuild through creature synergies and use wipes to reset opponents who get ahead.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- JUDGE [true]: "…alongside Denethor's own sacrifice outlet.". This asserts a card ability (that Denethor, Ruling Steward can sacrifice creatures), which is a claim about what a card may do. Denethor, Ruling Steward does include an activated ability that sacrifices another creature, so the claim appears accurate.
- PLAN plan_coherent=partly: The core is a real aristocrats engine (many sacrifice outlets, Bastion/Cutthroat/Falkenrath drains, token makers like Gisa and Old Flitterfang), but a five-card commander-protection package (Gift of Immortality, Together Forever, Take Up the Shield, Bastion Protector, Swiftfoot Boots) and three mass wipes pull against a go-wide sacrifice board, and dedicated fodder generation is thinner than the outlet count wants.
- PLAN theme_fit=yes: It is a WB aristocrats Commander deck built from cheap, commonly-owned staples at a bracket-2 power level with no fast mana or tutored combo lines, matching the theme, format, and stated constraints.
- PLAN useful_as_built=yes: 38 lands plus nine mana rocks/ramp pieces support a low curve, removal and draw are present in reasonable quantity, and there are multiple drain and combat win routes, so the deck is playable as it stands.
- PLAN summary_honest=yes: It names exactly the drain payoffs the list contains, admits the low ramp and lack of premium cards, and even flags the wipes as a deliberate reset rather than hiding the tension.
- [WARN] `not_owned`: Aron, Benalia's Ruin: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ayli, Eternal Pilgrim: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bartolomé del Presidio: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Woe Strider: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Yahenni, Undying Partisan: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Baron Bertram Graywater: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ghoulcaller Gisa: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Lord Skitter's Butcher: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Old Flitterfang: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Sivriss, Nightmare Speaker: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Skullport Merchant: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Umbral Collar Zealot: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vampire Gourmand: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Shadowheart, Dark Justiciar: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Smothering Abomination: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Disciple of Bolas: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Erebos, Bleak-Hearted: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bastion of Remembrance: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Falkenrath Noble: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Zulaport Cutthroat: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 2.84 over 61 nonland cards
- [INFO] `mana_pass`: the builder moved 1 card of the mana base to bring the deck inside its power level
- [INFO] `basics_added`: the list was 1 card short, so the builder added 1 basic land

<details><summary>The deck list</summary>

- 18 Plains | land | A basic white mana source for consistent early plays.
- 15 Swamp | land | A basic black mana source for consistent early plays.
- 1 Command Tower | land | A reliable two-color land slot.
- 1 Exotic Orchard | land | An owned land slot that supports both deck colors.
- 1 City of Brass | land | An owned flexible color source.
- 1 Grand Coliseum | land | An owned multicolor land slot.
- 1 Capital City | land | An owned land slot for the mana base.
- 1 Sol Ring | ramp | An efficient early mana piece.
- 1 Arcane Signet | ramp | Reliable two-color acceleration.
- 1 Chromatic Lantern | ramp | An owned mana-fixing ramp piece.
- 1 Fellwar Stone | ramp | An owned low-cost mana rock.
- 1 Thought Vessel | ramp | An owned mana rock for early development.
- 1 Springleaf Drum | ramp | A low-cost mana producer that works with the creature plan.
- 1 Deadly Dispute | ramp | A sacrifice-compatible ramp spell.
- 1 Inherited Envelope | ramp | An owned ramp artifact.
- 1 Ring of the Lucii | ramp | An owned mana-development piece.
- 1 White Auracite | ramp | An owned low-cost ramp slot.
- 1 Skullclamp | draw | A draw engine suited to expendable creatures.
- 1 Idol of Oblivion | draw | An owned draw artifact for the token-oriented plan.
- 1 Mask of Memory | draw | An owned equipment-based draw source.
- 1 Stone of Erech | draw | An owned artifact draw slot.
- 1 Tome of Legends | draw | An owned repeatable draw piece.
- 1 Wall of Omens | draw | An early defensive creature that supplies a card.
- 1 Inspiring Overseer | draw | An owned creature-based draw option.
- 1 Exemplar of Light | draw | An owned draw creature.
- 1 Cirith Ungol Patrol | draw | An owned black draw creature.
- 1 Nasty End | draw | A draw spell that fits a creature-heavy list.
- 1 Frontline Medic | interaction | An owned creature that helps protect the board.
- 1 Gift of Immortality | interaction | An owned protection piece for an important creature.
- 1 Swiftfoot Boots | interaction | An owned protection equipment slot.
- 1 Together Forever | interaction | An owned resilience piece for the creature plan.
- 1 Take Up the Shield | interaction | An inexpensive instant-speed protection option.
- 1 Bastion Protector | interaction | An owned protective creature for the commander.
- 1 Bitter Triumph | removal | An owned flexible removal spell.
- 1 Claim the Precious | removal | An owned black removal spell.
- 1 Crib Swap | removal | An owned creature-removal option.
- 1 Deadly Precision | removal | An owned low-cost removal slot.
- 1 Dispatch | removal | An owned white removal spell.
- 1 Infernal Grasp | removal | An owned unconditional removal option.
- 1 Palace Jailer | removal | An owned creature-based removal piece.
- 1 Austere Command | wipe | A flexible owned board-reset option.
- 1 Dusk // Dawn | wipe | An owned wipe that remains useful in a creature deck.
- 1 Martial Coup | wipe | An owned wipe for recovering stalled boards.
- 1 Aron, Benalia's Ruin | synergy | A low-cost aristocrats-focused synergy creature.
- 1 Ayli, Eternal Pilgrim | synergy | A low-cost creature that reinforces the sacrifice plan.
- 1 Bartolomé del Presidio | synergy | A cheap aristocrats synergy creature.
- 1 Woe Strider | synergy | A sacrifice-focused creature for the main plan.
- 1 Yahenni, Undying Partisan | synergy | A resilient sacrifice-focused synergy creature.
- 1 Gollum, Patient Plotter | synergy | An owned synergy piece for the creature-resource plan.
- 1 Gollum the Abandoned | synergy | An owned creature that supports the deck's synergy package.
- 1 Nimble Hobbit | synergy | An owned low-cost synergy creature.
- 1 Joo Dee, One of Many | synergy | An owned synergy creature for the broader creature plan.
- 1 Arcade Cabinet | synergy | An owned artifact synergy slot.
- 1 Baron Bertram Graywater | threat | A creature threat that adds to the board-based plan.
- 1 Ghoulcaller Gisa | threat | A substantial creature threat for the sacrifice deck.
- 1 Lord Skitter's Butcher | threat | A creature threat that contributes to the board.
- 1 Old Flitterfang | threat | A low-cost black creature threat.
- 1 Sivriss, Nightmare Speaker | threat | A low-cost creature threat for attrition games.
- 1 Skullport Merchant | threat | A creature threat that fits the resource-oriented plan.
- 1 Umbral Collar Zealot | threat | A low-cost creature threat.
- 1 Vampire Gourmand | threat | A low-cost black creature threat.
- 1 Shadowheart, Dark Justiciar | threat | A sizeable creature threat for the midgame.
- 1 Smothering Abomination | threat | A powerful black creature threat.
- 1 Disciple of Bolas | threat | A creature threat that complements resource exchanges.
- 1 Erebos, Bleak-Hearted | threat | A durable black threat for longer games.
- 1 Bastion of Remembrance | wincon | A finishing aristocrats payoff.
- 1 Falkenrath Noble | wincon | A creature finisher for the death-focused plan.
- 1 Zulaport Cutthroat | wincon | A compact aristocrats finisher.

</details>

### 17. upgrade a precon, owned first

Format: Commander. Theme: goblins. Pool: owned_first. Shortlist: 205 names.

Commander: Zada, Hedron Grinder.

Grade: typical, score 0.54, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $34.39 to buy, $211.01 the whole deck.

**Summary:** This deck builds a wide Goblin board, then uses Zada to turn targeted spells into explosive team-wide turns. Goblin synergies create pressure while draw and burst mana help assemble a decisive combat, with Great Train Heist and Earthquake as major closing tools. It gives up broad flexibility for a focused Goblin-and-spells game plan that is strongest when its creature board remains established.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The cards pair the way the top lists pair them, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every card serves one line \u2014 cheap token makers and Goblin lords build a wide board, one-mana cantrips and pump spells get multiplied by Zada, and rituals plus Past in Flames/Grapeshot support the storm finish \u2014 with only Blasphemous Act sitting awkwardly against a go-wide board.
- PLAN theme_fit=yes: It is a mono-red Commander deck led by a Goblin commander that fuses Goblin kindred with storm-style rituals and copied cantrips, with a Sol Ring/Skullclamp-level upgrade density and no tutor-chain or fast-combo shell, which sits in the bracket-3 range asked for.
- PLAN useful_as_built=yes: 37 lands plus Sol Ring, Arcane Signet, Skirk Prospector and several rituals support a very cheap curve, and the deck has multiple win routes (Goblin Bombardment/Impact Tremors drain, Shared Animosity/Great Train Heist alpha strikes, Grapeshot storm, Earthquake), so it plays as written.
- PLAN summary_honest=yes: Everything it claims is on the list (Zada spell-copying, Goblin token width, cantrip draw, Battle Hymn/Mana Geyser/Seething Song burst mana, Great Train Heist and Earthquake as closers) and it openly admits the deck's dependence on a surviving creature board.
- [INFO] `curve_summary`: average mana value 2.55 over 62 nonland cards
- [WARN] `profile_off_band`: the interaction count is 1, and bracket 3 wants 4 to 12
- [INFO] `bracket_cut`: to hold bracket 3, the builder cut 1 card: Storm-Kiln Artist (the combo Storm-Kiln Artist + Haze of Rage), and added 1 basic land

<details><summary>The deck list</summary>

- 1 Ancestors' Aid | other | A precon instant retained for the deck's spell-focused plan.
- 1 Arena of Glory | land | A retained precon land for the mana base.
- 1 Battle Hymn | ramp | A retained precon mana burst for creature-heavy turns.
- 1 Blasphemous Act | wipe | A retained precon reset button.
- 1 Boggart Shenanigans | removal | A retained Goblin payoff that supplies removal.
- 1 Brightstone Ritual | ramp | A retained precon mana burst for Goblin-heavy turns.
- 1 Castle Embereth | land | A retained precon land for the mana base.
- 1 Chaos Warp | removal | A retained flexible removal spell.
- 1 Conspicuous Snoop | synergy | A retained Goblin synergy piece.
- 1 Crimson Wisps | draw | A retained cantrip for Zada-focused spell turns.
- 1 Daring Discovery | other | A retained precon card that supports the established list.
- 1 Den of the Bugbear | land | A retained precon land for the mana base.
- 1 Dragon Fodder | synergy | A retained Goblin-producing synergy spell.
- 1 Dwarven Mine | land | A retained precon land for the mana base.
- 1 Empty the Warrens | synergy | A retained spell that builds the Goblin board.
- 1 Expedite | draw | A retained cantrip that fits Zada turns.
- 1 Faithless Looting | draw | A retained draw spell that keeps cards flowing.
- 1 Fists of Flame | draw | A retained draw spell for creature-focused turns.
- 1 Forgotten Cave | land | A retained precon land for the mana base.
- 1 Fountainport | land | A retained precon land for the mana base.
- 1 Frontline Heroism | other | A retained precon card that supports the established plan.
- 1 Gempalm Incinerator | removal | A retained Goblin removal option.
- 1 General Kreat, the Boltbringer | synergy | A retained Goblin synergy piece.
- 1 Glimpse the Impossible | ramp | A retained precon mana-development spell.
- 1 Goblin Bombardment | removal | A retained Goblin-based removal outlet.
- 1 Goblin Bushwhacker | synergy | A retained Goblin synergy creature.
- 1 Goblin Chieftain | synergy | A retained Goblin synergy creature.
- 1 Goblin Dark-Dwellers | threat | A retained Goblin threat.
- 1 Goblin Lackey | synergy | A retained Goblin synergy creature.
- 1 Goblin Matron | synergy | A retained Goblin synergy creature.
- 1 Goblin Negotiation | removal | A retained Goblin removal spell.
- 1 Goblin Trashmaster | removal | A retained Goblin removal creature.
- 1 Goblin Warchief | synergy | A retained Goblin synergy creature.
- 1 Grapeshot | removal | A retained spell-based removal option.
- 1 Great Train Heist | wincon | A retained finisher for closing games.
- 1 Grenzo, Havoc Raiser | synergy | A retained Goblin synergy creature.
- 1 Haze of Rage | other | A retained precon spell for the established plan.
- 1 Heirloom Blade | synergy | A retained Goblin synergy equipment.
- 1 Hidden Volcano | land | A retained precon land for the mana base.
- 1 Howlsquad Heavy | ramp | A retained creature that contributes ramp.
- 1 Idol of Oblivion | draw | A retained draw engine.
- 1 Impact Tremors | other | A retained precon card that supports the established plan.
- 1 Kher Keep | land | A retained precon land for the mana base.
- 1 Krenko's Command | synergy | A retained Goblin-producing synergy spell.
- 1 Mana Geyser | ramp | A retained large mana burst.
- 1 Mogg War Marshal | synergy | A retained Goblin synergy creature.
- 26 Mountain | land | Basic red sources keep the mana base reliable for the single-color deck.
- 1 Past in Flames | other | A retained precon card for the deck's spell-focused turns.
- 1 Quest for the Goblin Lord | synergy | A retained Goblin synergy payoff.
- 1 Reliquary Tower | land | A retained precon land for the mana base.
- 1 Ruby Medallion | other | A retained precon card that supports the established plan.
- 1 Sazacap's Brew | draw | A retained draw spell.
- 1 Searslicer Goblin | synergy | A retained Goblin synergy creature.
- 1 Seething Song | ramp | A retained precon mana burst.
- 1 Shinka, the Bloodsoaked Keep | land | A retained precon land for the mana base.
- 1 Siege-Gang Commander | removal | A retained Goblin removal creature.
- 1 Siege-Gang Lieutenant | removal | A retained Goblin removal creature.
- 1 Skirk Prospector | ramp | A retained Goblin mana source.
- 1 Skullclamp | draw | A retained draw engine for the creature plan.
- 1 Sol Ring | ramp | A retained efficient mana source.
- 1 Spreading Insurrection | other | A retained precon card that supports the established plan.
- 1 Swiftfoot Boots | interaction | A retained protection piece.
- 1 Throne of Eldraine | ramp | A retained precon mana source.
- 1 Vandalblast | wipe | A retained artifact-focused wipe.
- 1 War Room | land | A retained precon land for the mana base.
- 1 Warren Torchmaster | synergy | A retained Goblin synergy creature.
- 1 Wild Ride | other | A retained precon card that supports the established plan.
- 1 Witch's Mark | draw | A retained draw spell for creature-focused turns.
- 1 Ancestral Anger | draw | A low-cost draw spell that scales especially well across Zada's creature board.
- 1 Arcane Signet | ramp | Reliable early ramp helps deploy Zada and develop the board.
- 1 Earthquake | wipe | A flexible wipe that also supplies an additional finishing route.
- 1 Fable of the Mirror-Breaker // Reflection of Kiki-Jiki | ramp | A Goblin-adjacent ramp card that adds value to the existing plan.
- 1 Krenko, Mob Boss | threat | A powerful Goblin threat that reinforces the creature-heavy plan.
- 1 Shared Animosity | synergy | A Goblin payoff that turns a wide board into a serious combat threat.

</details>

### 18. upgrade a precon, any card

Format: Commander. Theme: turtles. Pool: any_card. Shortlist: 364 names.

Commander: Heroes in a Half Shell.

Grade: baseline, score 0.06, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $231.26 to buy, $231.26 the whole deck.

**Summary:** This keeps Turtle Power focused on a broad team of Turtles, Mutants, and Ninjas that wants to attack together, with Heroes in a Half Shell turning successful combat into larger attackers and more cards. The upgraded mana base and compact ramp package make the five-color plan steadier, while Rhystic Study and Kindred Discovery support a longer game. Dimension X Pizzasaur and Everything Pizza remain chief among the finishing threats, backed by efficient removal and the precon’s sweepers. The tradeoff is that the deck remains a combat-focused tribal build rather than a tightly streamlined combo deck.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves one color far better than another, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: The list is a unified tribal-combat deck: Mutant/Turtle/Ninja bodies, counter-and-anthem support (Corpsejack Menace, Cover of Darkness, Level Up, Together Forever), a few sweepers and spot removal, and card draw that keys off attacking creatures, with only a couple of loosely connected goodstuff creatures.
- PLAN theme_fit=yes: It is the Turtle Power precon shell kept intact around Heroes in a Half Shell in five colors, with bracket-3-appropriate additions (Sol Ring, Swords to Plowshares, Assassin's Trophy, Rhystic Study, shocks/fetches) rather than a rebuild into something else.
- PLAN useful_as_built=yes: 40 lands with heavy five-color fixing (Command Tower, City of Brass, Mana Confluence, Exotic Orchard, duals, Thriving lands, fetches) plus five ramp pieces support a playable curve, and the deck has clear win routes through a wide attacking board, the Pizzasaur, and Everything Pizza.
- PLAN summary_honest=yes: Every element it names is present — Rhystic Study, Kindred Discovery, the compact ramp package (Sol Ring, Arcane Signet, Fellwar Stone, Chromatic Lantern, Cultivate), the sweepers, the removal suite, and both pizza finishers — and it openly concedes the deck is a combat tribal build with an uneven, green-leaning mana base.
- [INFO] `curve_summary`: average mana value 3.37 over 59 nonland cards
- [WARN] `profile_off_band`: the land count is 40, and bracket 3 wants 34 to 38
- [WARN] `profile_off_band`: the ramp count is 6, and bracket 3 wants 8 to 13
- [WARN] `profile_off_band`: the draw count is 3, and bracket 3 wants 8 to 14
- [WARN] `profile_off_band`: the removal count is 4, and bracket 3 wants 6 to 12
- [WARN] `profile_off_band`: the wipe count is 1, and bracket 3 wants 2 to 5
- [WARN] `profile_off_band`: the interaction count is 0, and bracket 3 wants 4 to 12
- [INFO] `basics_added`: the list was 1 card short, so the builder added 1 basic land

<details><summary>The deck list</summary>

- 1 Acidic Slime | other | Retained precon creature.
- 1 April O'Neil, Live on the Scene | draw | Retained precon draw card.
- 1 Arcane Signet | ramp | Retained precon ramp artifact.
- 1 Assassin's Trophy | removal | Retained precon removal.
- 1 Baxter, Fly in the Ointment | other | Retained precon creature.
- 1 Bebop, Skull & Crossbones | other | Retained precon creature.
- 1 Beast Within | removal | Added flexible removal.
- 1 Big Mother Mouser | other | Retained precon artifact creature.
- 1 Biogenic Ooze | other | Retained precon creature.
- 1 Blasphemous Act | wipe | Retained precon board wipe.
- 1 Chromatic Lantern | ramp | Retained precon mana artifact.
- 1 Continue? | other | Retained precon instant.
- 1 Corpsejack Menace | other | Retained precon creature.
- 1 Cover of Darkness | synergy | Added tribal synergy.
- 1 Cultivate | ramp | Retained precon ramp spell.
- 1 Dimension X Pizzasaur | wincon | Retained precon finisher.
- 1 Donatello, the Brains | synergy | Retained precon Turtle synergy.
- 1 Double Jump // Flying Kick | other | Retained precon instant.
- 1 Electric Seaweed | other | Retained precon creature.
- 1 Endless Foot Assault | other | Retained precon enchantment.
- 1 Everything Pizza | wincon | Retained precon finisher.
- 1 Fast Forward | other | Retained precon sorcery.
- 1 Fellwar Stone | ramp | Added efficient ramp.
- 1 Foot Chopper | other | Retained precon Equipment.
- 1 Game Over | other | Retained precon sorcery.
- 1 Harmonize | other | Retained precon sorcery.
- 1 Here Comes a New Hero! | other | Retained precon sorcery.
- 1 High Score | other | Retained precon enchantment.
- 1 Irma, Part-Time Mutant | other | Retained precon creature.
- 1 Kindred Discovery | draw | Added tribal card draw.
- 1 Krang, the All-Powerful | other | Retained precon artifact creature.
- 1 Leatherhead, Iron Gator | other | Retained precon creature.
- 1 Leonardo, the Balance | threat | Retained precon Turtle threat.
- 1 Lessons from Life | other | Retained precon sorcery.
- 1 Level Up | other | Retained precon Aura.
- 1 Lita, Little Orphan Amphibian | synergy | Retained precon Turtle synergy.
- 1 Mole Module | other | Retained precon Vehicle.
- 1 Mona Lisa, Science Geek | other | Retained precon creature.
- 1 Ninja Pizza | other | Retained precon enchantment.
- 1 Rat King, Pale Piper | other | Retained precon creature.
- 1 Ray Fillet, Wave Warrior | other | Retained precon creature.
- 1 Rhystic Study | draw | Added repeatable card draw.
- 1 Roadkill Rodney | other | Retained precon artifact creature.
- 1 Rocksteady, Mutant Marauder | other | Retained precon creature.
- 1 Shellshock | other | Retained precon instant.
- 1 Shredder, Shadow Master | other | Retained precon Ninja creature.
- 1 Sol Ring | ramp | Retained precon ramp artifact.
- 1 Special Move | other | Retained precon instant.
- 1 Steelbane Hydra | removal | Retained precon removal creature.
- 1 Super Combo | other | Retained precon sorcery.
- 1 Swords to Plowshares | removal | Added efficient removal.
- 1 Swift Demise | other | Retained precon instant.
- 1 Tempestra, Dame of Games | other | Retained precon creature.
- 1 Together Forever | other | Retained precon enchantment.
- 1 Tokka & Rahzar, Unsupervised | ramp | Retained precon ramp creature.
- 1 Vanquish the Horde | other | Retained precon sorcery.
- 1 Vigor | other | Retained precon creature.
- 1 Voracious Hydra | other | Retained precon creature.
- 1 Wave Goodbye | other | Retained precon sorcery.
- 1 Ash Barrens | land | Retained precon land.
- 1 Big Apple, 3 a.m. | land | Retained precon land.
- 1 Breeding Pool | land | Added untapped blue-green source.
- 1 Cinder Glade | land | Retained precon land.
- 1 City of Brass | land | Retained precon land.
- 1 Command Tower | land | Retained precon land.
- 1 Dragonskull Summit | land | Retained precon land.
- 1 Escape Tunnel | land | Retained precon land.
- 1 Evolving Wilds | land | Retained precon land.
- 1 Exotic Orchard | land | Retained precon land.
- 1 Fabled Passage | land | Retained precon land.
- 1 Godless Shrine | land | Added untapped white-black source.
- 1 Hidden Hideout | land | Retained precon land.
- 1 Hinterland Harbor | land | Retained precon land.
- 1 Mana Confluence | land | Added untapped five-color source.
- 1 Rootbound Crag | land | Retained precon land.
- 1 Smoldering Marsh | land | Retained precon land.
- 1 Sodden Verdure | land | Retained precon land.
- 1 Spire Garden | land | Retained precon land.
- 1 Sunken Hollow | land | Retained precon land.
- 1 Thriving Grove | land | Retained precon land.
- 1 Thriving Isle | land | Retained precon land.
- 1 Thriving Moor | land | Retained precon land.
- 1 Turtle Lair | land | Retained precon land.
- 1 Undergrowth Stadium | land | Retained precon land.
- 1 Vernal Fen | land | Retained precon land.
- 1 Vibrant Cityscape | land | Retained precon land.
- 4 Forest | land | Basic green source for the five-color mana base.
- 3 Island | land | Basic blue source for the five-color mana base.
- 2 Mountain | land | Basic red source for the five-color mana base.
- 2 Plains | land | Basic white source for the five-color mana base.
- 2 Swamp | land | Basic black source for the five-color mana base.

</details>

### 19. the Hobbit family, two colours

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 174 names.

Commander: Thranduil, the Elvenking.

Grade: baseline, score 0.06, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $565.14 to buy, $565.14 the whole deck.

**Summary:** Thranduil leads a Sultai Elf creature deck that develops its mana early, builds pressure with a broad Elf board, and uses legendary Elves to fuel the commander’s draw-and-discard trigger. The deck keeps opponents in check with a substantial removal and interaction suite, then closes through Troll of Khazad-dûm or either Witch-king. Its main concession is that much of its pressure and synergy is creature-based, so repeated board clears can slow its ability to establish a decisive battlefield.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. Many of the cards appear in no top list, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core of legendary Elves, mana dorks, and cheap Elf bodies to feed Thranduil's trigger is consistent, but three sweepers (Languish, Gnashing of Teeth, Raise the Palisade) and a handful of unrelated big finishers (Troll of Khazad-dûm, both Witch-kings) pull against the wide Elf board the rest of the deck is trying to build.
- PLAN theme_fit=yes: It is a Commander deck built from the requested Hobbit sets, led by Thranduil, the Elvenking, in his Sultai identity, with a game-changer count (The One Ring, Orcish Bowmasters, Mox Amber) that sits within bracket 3 expectations.
- PLAN useful_as_built=partly: With 39 lands, several mana dorks, and a low curve the deck will cast its spells, but the base is 29 basics plus mostly colorless utility lands with almost no dual or fixing sources, so a {2}{B}{G}{U} commander and split GG/UU/BB costs will be awkward to cast on time.
- PLAN summary_honest=partly: The removal suite, Elf density, and named finishers are all genuinely present, but the summary flags vulnerability to board clears without disclosing that the deck itself runs three of them against its own creature base, and it glosses over the absence of dual lands in a three-color list.
- [INFO] `curve_summary`: average mana value 3.23 over 60 nonland cards
- [WARN] `profile_off_band`: the land count is 39, and bracket 3 wants 34 to 38
- [INFO] `mana_pass`: the builder moved 2 cards of the mana base to bring the deck inside its power level
- [INFO] `basics_added`: the list was 3 cards short, so the builder added 3 basic lands

<details><summary>The deck list</summary>

- 1 Arcane Signet | ramp | Reliable multicolor mana acceleration.
- 1 Arwen, Weaver of Hope | synergy | A legendary Elf that supports Thranduil's legendary-Elf plan.
- 1 Bitter Downfall | removal | Dedicated removal from the shortlist.
- 1 Bilbo's Deadly Slice | removal | Dedicated removal from the shortlist.
- 1 Boughside Wanderers | threat | An Elf creature that adds board pressure.
- 1 Cantankerous Keepers | threat | An Elf creature that adds board pressure.
- 1 Celeborn the Wise | synergy | A legendary Elf that supports Thranduil's legendary-Elf plan.
- 1 Colossal Whale | removal | A large evasive finisher that also fills a removal slot.
- 1 Confusticate and Bebother | interaction | Flexible instant-speed interaction.
- 1 Crude Bent Blade | removal | Equipment-based removal from the shortlist.
- 1 Delighted Halfling | ramp | Low-cost creature-based mana acceleration.
- 1 Elrond, Moon-Reader | interaction | A legendary Elf that provides interaction and can trigger Thranduil.
- 1 Elven Chorus | ramp | Elf-focused mana acceleration.
- 1 Elven Passage | land | Part of the three-color land base.
- 1 Elven Raft-Steerer | threat | An Elf creature that adds board pressure.
- 1 Elvenking's Halls | land | Part of the three-color land base.
- 1 Elvenking's Harper | threat | An Elf creature that adds board pressure.
- 1 Elvish Archdruid | ramp | Elf-focused mana acceleration.
- 1 Elvish Mystic | ramp | Low-cost creature-based mana acceleration.
- 1 Elvish Visionary | draw | An Elf that supplies card draw.
- 1 Enchanted River's Grasp | removal | Dedicated removal from the shortlist.
- 1 Fateful Discovery | draw | A dedicated card-draw piece.
- 10 Forest | land | Basic green mana for the Elf-heavy core.
- 1 Galadhrim Guide | threat | An Elf creature that adds board pressure.
- 1 Galion, Elvenking's Butler | synergy | A legendary Elf that supports Thranduil's legendary-Elf plan.
- 1 Giant's Boulder | removal | Artifact-based removal from the shortlist.
- 1 Gnashing of Teeth | wipe | A board-clearing effect.
- 1 Grey Havens Navigator | threat | An Elf creature that adds board pressure.
- 1 Guardian of the Halls | threat | An Elf creature that adds board pressure.
- 1 Haunt of the Dead Marshes | threat | An Elf creature that adds board pressure.
- 1 Hithlain Knots | draw | Instant-speed card draw.
- 1 Hobbit Hole | land | Part of the three-color land base.
- 8 Island | land | Basic blue mana for the deck's interaction and draw.
- 1 Key to the Side-Door | draw | Artifact-based card draw.
- 1 Languish | wipe | A board-clearing effect.
- 1 Last March of the Ents | draw | A dedicated card-draw spell.
- 1 Lórien Revealed | draw | A dedicated card-draw spell.
- 1 Lothlórien Lookout | synergy | An Elf that reinforces the deck's creature theme.
- 1 Merciless Executioner | removal | Creature-based removal from the shortlist.
- 1 Minas Morgul, Dark Fortress | land | Part of the three-color land base.
- 1 Mirkwood | land | Part of the three-color land base.
- 1 Mirkwood Meditator | synergy | An Elf that reinforces the deck's creature theme.
- 1 Mirkwood Nurturer | synergy | An Elf that reinforces the deck's creature theme.
- 1 Mirkwood Pathmaker | threat | An Elf creature that adds board pressure.
- 1 Mithril Coat | interaction | Protective interaction for important creatures.
- 1 Mox Amber | ramp | Fast legendary-focused mana acceleration.
- 1 My Precious // Allure of Power | interaction | Equipment-based interaction.
- 1 Night's Whisper | draw | Efficient dedicated card draw.
- 1 Nimrodel Watcher | threat | An Elf creature that adds board pressure.
- 1 Orcish Bowmasters | removal | Creature-based removal from the shortlist.
- 1 Palantír of Orthanc | draw | A dedicated card-draw artifact.
- 1 Plunder the Trollshaws | draw | Instant-speed card draw.
- 1 Quarrel | removal | Dedicated instant-speed removal.
- 1 Raise the Palisade | wipe | A board-clearing effect.
- 1 Rivendell | land | Part of the three-color land base.
- 1 Rogue's Passage | land | A utility land for the mana base.
- 1 Stern Scolding | interaction | Low-cost instant-speed interaction.
- 11 Swamp | land | Basic black mana for removal and finishers.
- 1 The Black Arrow | removal | Equipment-based removal from the shortlist.
- 1 The Black Gate | land | Part of the three-color land base.
- 1 The One Ring | interaction | A legendary artifact that supplies interaction.
- 1 The Shire | land | Part of the three-color land base.
- 1 Thranduil the Strategist | ramp | A legendary Elf that accelerates mana and can trigger Thranduil.
- 1 Thranduil's Decree | interaction | Thematic instant-speed interaction.
- 1 Thranduil, Sindarin Liege // Silvan Rally | synergy | A legendary Elf that supports Thranduil's legendary-Elf plan.
- 1 Treasure Vault | land | A utility land for the mana base.
- 1 Troll of Khazad-dûm | wincon | A marked finisher and primary closing threat.
- 1 Uncover the Moon-Letters | draw | A dedicated card-draw enchantment.
- 1 Wayfarer's Bauble | ramp | Early artifact-based mana acceleration.
- 1 Witch-king of Angmar | wincon | A marked finisher and primary closing threat.
- 1 Witch-king, Bringer of Ruin | wincon | A marked finisher and primary closing threat.
- 1 Wood Elves | ramp | Elf-based mana acceleration.
- 1 Woodland Weavemaster | ramp | Elf-based mana acceleration.

</details>

### 20. the Hobbit family, a delegated commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 77 names.

Commander: Smaug the Magnificent.

Grade: baseline, score 0.06, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $504.69 to buy, $504.69 the whole deck.

**Summary:** Smaug the Magnificent is the center of a red dragon-and-Treasure strategy that develops mana, attacks early, and turns its Treasure stockpile into mounting pressure. Cavern-Hoard Dragon, Desert Were-Worm, Smaug, the Great Calamity // Spew Flame, and chief among the supporting threats Dáin Ironfoot provide the closing force, while a broad spread of removal, interaction, and board resets keeps opponents from assembling an easy defense. The deck gives up some card selection and specialized answers in exchange for a direct, creature-driven plan built around Smaug’s attacks and a strongly themed supporting cast.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: There is a recognizable mono-red aggro core of cheap creatures, burn and equipment that attacks early, but the summary's "dragon-and-Treasure" engine is barely present, and the three mass-damage wipes plus a sprawling equipment pile with few equipment payoffs pull against the wide, cheap-creature board the rest of the list builds.
- PLAN theme_fit=partly: The set restriction (hob/hoc), mono-red identity and dragon commander match the request, but a dragons deck with only three non-commander Dragons and a supporting cast of dwarves, goblins and orcs is not really the tribe the person asked for.
- PLAN useful_as_built=yes: 36 lands in a mono-red deck with a low curve, adequate ramp and draw, plenty of removal and several real finishers means it can be shuffled up and played as printed.
- PLAN summary_honest=partly: It accurately flags the removal suite, board resets and weak card selection, but it oversells a "Treasure stockpile" and a dragon centerpiece that only four cards support, and it never mentions that the deck's real body is dwarves, goblins and a dozen Equipment.
- [INFO] `curve_summary`: average mana value 3.11 over 63 nonland cards
- [WARN] `profile_off_band`: the draw count is 6, and bracket 3 wants 8 to 14

<details><summary>The deck list</summary>

- 32 Mountain | land | Supplies the deck’s red mana base.
- 1 The Lonely Mountain | land | A Mountain land that supports the red mana base.
- 1 Dragon-Cursed Halls | land | Adds a themed land slot to the mana base.
- 1 Hobbit Hole | land | Adds another land to keep opening hands stable.
- 1 Treasure Vault | land | A land slot that fits the deck’s Treasure focus.
- 1 Arcane Signet | ramp | Reliable mana development for casting Smaug and the creature suite.
- 1 Bag End Banquet | ramp | Adds to the deck’s mana development package.
- 1 Dragon's Desire | ramp | Dragon-themed acceleration for the main plan.
- 1 Fíli and Kíli, Joyous | ramp | A ramp creature that also contributes to the creature-heavy board.
- 1 Mox Amber | ramp | A low-cost mana piece alongside the deck’s legendary cards.
- 1 Orcrist, Goblin-cleaver | ramp | An Equipment ramp piece for the deck’s combat plan.
- 1 The Reaver Cleaver | ramp | Supports mana development in a combat-focused shell.
- 1 Thorin, Company's Leader | ramp | A ramp threat and marked finisher for the creature plan.
- 1 Troop of Ponies | ramp | Provides another ramp body for the board.
- 1 Wayfarer's Bauble | ramp | A dependable early mana-development card.
- 1 Balin, Loremaster | draw | One of the available draw cards for maintaining resources.
- 1 Key to the Side-Door | draw | Adds card draw from an artifact slot.
- 1 Palantír of Orthanc | draw | A legendary draw artifact for sustained resources.
- 1 Ragged Short Spear | draw | An Equipment draw card that fits the combat shell.
- 1 Thrór's Map | draw | Provides another draw artifact for the deck.
- 1 Óin the Brave | draw | A creature-based source of card draw.
- 1 Battle-Scarred Goblin | removal | Creature-based removal for opposing problems.
- 1 Fire of Orthanc | removal | A direct removal spell.
- 1 Gandalf, Spark Starter | removal | A removal creature that can remain on the board.
- 1 Giant's Boulder | removal | Artifact-based removal for the deck.
- 1 Goblin Cratermaker | removal | A compact creature-based removal option.
- 1 Goblin Fireleaper | removal | Adds another removal body to the creature package.
- 1 Improvised Club | removal | An instant-speed removal option.
- 1 Pinecone Strike | removal | Adds flexible instant-speed removal.
- 1 Smite the Deathless | removal | A further removal spell for difficult opposing cards.
- 1 Call Forth the Tempest | wipe | One of the deck’s board-reset options.
- 1 Desolation of Smaug | wipe | A second board-reset option with Smaug flavor.
- 1 Glóin the Mighty // Easy Pickings | wipe | A creature card that also supplies a wipe option.
- 1 Bilbo's Ring | interaction | A legendary artifact interaction piece.
- 1 Dwarven Mattock | interaction | An Equipment-based interaction card.
- 1 Mithril Coat | interaction | A protective interaction artifact for key creatures.
- 1 The One Ring | interaction | A legendary artifact interaction piece.
- 1 Andúril, Flame of the West | interaction | An Equipment interaction slot for the combat plan.
- 1 Andúril, Narsil Reforged | interaction | A second legendary Equipment interaction slot.
- 1 The Black Arrow | interaction | An Equipment interaction card with a forceful combat identity.
- 1 Getaway Barrel | interaction | An artifact interaction slot that broadens the deck’s options.
- 1 Last Light of Durin's Day | synergy | The dedicated synergy card from the shortlist.
- 1 Glamdring | synergy | A legendary Equipment that supports the creature-and-combat plan.
- 1 Long-Lost Lances | synergy | An Equipment piece for the deck’s combat-focused board.
- 1 Sting, Bilbo's Sword | synergy | A legendary Equipment that supports the attacking plan.
- 1 Well-Worn Spatula | synergy | An Equipment slot that supports the deck’s creature plan.
- 1 Smaug's Fury | synergy | A Smaug-themed card included for the deck’s central plan.
- 1 Tidings of War | synergy | Supports the deck’s aggressive creature strategy.
- 1 Cavern-Hoard Dragon | wincon | A Dragon win condition and marked finisher.
- 1 Desert Were-Worm | wincon | A marked finisher that gives the deck another closing threat.
- 1 Smaug, the Great Calamity // Spew Flame | wincon | A Smaug-themed marked finisher for closing games.
- 1 Dáin Ironfoot | threat | A marked finisher that adds weight to the creature suite.
- 1 Gandalf, Goblins' Bane // Flameshape | threat | A substantial legendary creature threat.
- 1 Dwarven Mauler | threat | A creature threat that helps establish a board.
- 1 Dwarven Warriors | threat | Adds another creature threat to pressure opponents.
- 1 Goblin-town Flunkies | threat | A low-end creature threat for building a board.
- 1 Gundabad Opportunist | threat | A Goblin creature threat for the attacking plan.
- 1 Guttersnipe | threat | A creature threat that complements the spell package.
- 1 Iron Hills Stalwart | threat | A creature threat that strengthens the board presence.
- 1 Misty Mountains Raider | threat | A Goblin threat for the combat plan.
- 1 Olog-hai Crusher | threat | A large creature threat for pressuring life totals.
- 1 Orcish Siegemaster | threat | An Orc threat that broadens the attacking force.
- 1 Snowslope Hunter | threat | A creature threat that contributes to the board.
- 1 Bombur, Gentle Dreamer | other | A legendary creature included as a supporting card.
- 1 Bothersome Noisemaker | other | A supporting creature slot in the deck’s creature package.
- 1 Dori, Bearer of Friends | other | A legendary creature included as a supporting card.
- 1 Old Thrush | other | A small supporting creature for the board.
- 1 Oliphaunt | other | A larger supporting creature for the deck.

</details>

### 21. the Hobbit family, mana from outside

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 156 names.

Commander: Smaug the Impenetrable.

Grade: baseline, score 0.04, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $768.37 to buy, $768.37 the whole deck.

**Summary:** This is a black-red Smaug deck that develops its mana with artifacts, Treasures, and Dragon-oriented acceleration before bringing Smaug the Impenetrable to the table as its central threat. It maintains pressure through Goblins, Orcs, Wargs, and larger monsters, uses a broad removal package to force attacks through, and closes with Smaug, the Great Calamity, Troll of Khazad-dûm, or Witch-king of Angmar. The deck gives up some early speed for a mana-heavy, thematic midrange plan built to support expensive Dragons and legendary equipment.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core is a coherent black-red attack-and-big-monsters midrange plan, but it splits attention between a wide goblin/warg aggro shell, a heavy eight-piece legendary equipment package with few strong carriers, and expensive top-end dragons, so the pieces support each other only loosely.
- PLAN theme_fit=yes: It is a Commander deck led by Smaug the Impenetrable in BR with every nonland, non-rock card drawn from hob/hoc, and its midrange power level is in line with the requested bracket 3.
- PLAN useful_as_built=partly: There are enough lands, threats, removal and win conditions to sit down and play, but nine fetchlands with only a single Blood Crypt target (and three—Flooded Strand, Misty Rainforest, Windswept Heath—that can fetch nothing at all) leaves roughly eight effectively dead lands under a curve topped by a seven-mana commander.
- PLAN summary_honest=partly: Most claims check out (goblins/orcs/wargs, removal, equipment, slow mana-heavy plan, and it even admits the early-speed weakness), but the named Treasure acceleration is barely present and the summary passes over the large number of off-color fetchlands that make the mana worse than described.
- [INFO] `curve_summary`: average mana value 3.18 over 62 nonland cards
- [INFO] `mana_pass`: the builder moved 2 cards of the mana base to bring the deck inside its power level
- [WARN] `outside_requested_set`: the sets you named do not hold 17 cards: Arid Mesa, Blood Crypt, Bloodstained Mire, Bojuka Bog, Command Tower, Dragonskull Summit, Evolving Wilds, Exotic Orchard, Fabled Passage, Flooded Strand, Marsh Flats, Misty Rainforest, Polluted Delta, Scalding Tarn, Verdant Catacombs, Windswept Heath, Wooded Foothills

<details><summary>The deck list</summary>

- 8 Mountain | land | A basic red source that supports the deck’s red spells and a reliable untapped mana base.
- 7 Swamp | land | A basic black source that supports the deck’s black spells and a reliable untapped mana base.
- 1 Blood Crypt | land | A dual-colored land source for both halves of Smaug’s color identity.
- 1 Bloodstained Mire | land | A flexible land source that helps access the deck’s primary colors.
- 1 Bojuka Bog | land | A black land slot with useful utility.
- 1 Command Tower | land | A dependable dual-colored land source.
- 1 Dragon-Cursed Halls | land | A thematic colored land source for the Dragon plan.
- 1 Dragonskull Summit | land | An efficient black-red land source.
- 1 Evolving Wilds | land | A flexible land source that finds the needed basic color.
- 1 Exotic Orchard | land | A flexible colored land for multiplayer games.
- 1 Fabled Passage | land | A flexible land source that finds the needed basic color.
- 1 Goblin-town | land | A thematic red land source.
- 1 Hobbit Hole | land | A colored utility land slot.
- 1 Minas Morgul, Dark Fortress | land | A thematic black land source.
- 1 Mount Doom | land | A thematic colored land source for the deck.
- 1 Arid Mesa | land | A flexible land source that can find a Mountain.
- 1 Flooded Strand | land | A flexible land source that can find a Swamp or Mountain.
- 1 Marsh Flats | land | A flexible land source that can find a Swamp.
- 1 Misty Rainforest | land | A flexible land source that can find a Swamp or Mountain.
- 1 Polluted Delta | land | A flexible land source that can find a Swamp.
- 1 Scalding Tarn | land | A flexible land source that can find a Mountain.
- 1 Verdant Catacombs | land | A flexible land source that can find a Swamp.
- 1 Windswept Heath | land | A flexible land source that can find a Mountain.
- 1 Wooded Foothills | land | A flexible land source that can find a Mountain.
- 1 Arcane Signet | ramp | Early fixing and acceleration toward Smaug.
- 1 Burn, Burn, Tree and Fern | ramp | A thematic ramp piece that advances the mana plan.
- 1 Dragon's Desire | ramp | A themed acceleration spell for the Dragon plan.
- 1 Mox Amber | ramp | Fast early acceleration alongside the legendary commander and permanents.
- 1 Orcrist, Goblin-cleaver | ramp | A ramp equipment that also fits the deck’s legendary armory.
- 1 Smaug the Magnificent | ramp | A thematic Dragon ramp card that supports the top end.
- 1 Smaug, Wicked Worm | ramp | A powerful Dragon ramp piece that remains threatening later.
- 1 The Reaver Cleaver | ramp | A Treasure-oriented ramp piece that supports Smaug’s plan.
- 1 Wayfarer's Bauble | ramp | Low-cost acceleration that improves early land development.
- 1 Balin, Loremaster | draw | A dedicated draw card that keeps resources flowing.
- 1 Gollum, Riddle Master | draw | A draw threat that provides card flow while adding pressure.
- 1 Key to the Side-Door | draw | A colorless draw tool that fits smoothly into the mana base.
- 1 Night's Whisper | draw | An efficient black draw spell.
- 1 Palantír of Orthanc | draw | A legendary draw engine for longer games.
- 1 Rage into the Valley | draw | A red draw spell that refills the hand.
- 1 Ragged Short Spear | draw | An equipment draw source that remains on theme.
- 1 The Master of Lake-town | draw | A dedicated draw creature for sustained resources.
- 1 The Sackville-Bagginses | draw | A draw permanent that helps maintain card flow.
- 1 Thrór's Map | draw | A colorless draw artifact that supports the deck’s longer game.
- 1 Óin the Brave | draw | A draw creature that adds to the board while replenishing cards.
- 1 Battle-Scarred Goblin | removal | A creature-based removal option that fits the aggressive board plan.
- 1 Bolg of the North | removal | A thematic removal creature for opposing problems.
- 1 Goblin Cratermaker | removal | A flexible Goblin removal piece.
- 1 Goblin Fireleaper | removal | A creature-based removal option that maintains board presence.
- 1 Inferno Titan | removal | A major red threat that also supplies removal.
- 1 Merciless Executioner | removal | A black creature-based answer to opposing threats.
- 1 Stir Up Trouble | removal | A direct removal spell for clearing resistance.
- 1 Stone-Giant of High Pass | removal | A removal creature that contributes to the deck’s pressure.
- 1 The Black Arrow | removal | A thematic equipment-based removal tool.
- 1 Desolation of Smaug | wipe | A thematic board wipe that clears a path for the Dragon plan.
- 1 Languish | wipe | An efficient black sweep for creature-heavy boards.
- 1 Andúril, Flame of the West | interaction | A legendary equipment that provides a flexible protective utility slot.
- 1 Bilbo's Ring | interaction | A legendary equipment interaction piece.
- 1 Dwarven Mattock | interaction | An equipment-based utility answer.
- 1 Glamdring | interaction | A legendary equipment that broadens the deck’s utility options.
- 1 Mithril Coat | interaction | A protective equipment for preserving key threats.
- 1 My Precious // Allure of Power | interaction | A flexible legendary interaction card.
- 1 Sting, Bilbo's Sword | interaction | A low-cost legendary equipment utility piece.
- 1 The One Ring | interaction | A legendary artifact interaction tool for difficult positions.
- 1 Along the Crooked Way | synergy | A support enchantment for the deck’s Smaug-focused game plan.
- 1 Gathering of Darkness | synergy | A black support spell for the deck’s central plan.
- 1 Guttersnipe | synergy | A red support creature that rewards the deck’s spell-heavy utility package.
- 1 Last Light of Durin's Day | synergy | A thematic enchantment support piece.
- 1 Shadow of the Enemy | synergy | A black support spell for the deck’s proactive strategy.
- 1 Smaug's Fury | synergy | A thematic instant that supports the Smaug plan.
- 1 Supper for Spiders | synergy | The dedicated synergy card for Smaug’s plan.
- 1 Bothersome Noisemaker | threat | A low-cost creature that helps establish an attacking board.
- 1 Dreaded Bat-Cloud | threat | A black creature threat that adds pressure.
- 1 Fearsome Goblin Pair | threat | A Goblin threat that develops the board.
- 1 Great Goblin, Foul-Hearted | threat | A legendary Goblin threat for the aggressive plan.
- 1 Great Ugly-Looking Goblin // Clap! Snap! | threat | A flexible Goblin threat that contributes to board pressure.
- 1 Gundabad Opportunist | threat | A Goblin creature that adds early and midgame pressure.
- 1 Haunt of the Dead Marshes | threat | A black creature threat for the battlefield plan.
- 1 Misty Mountains Raider | threat | A Goblin attacker that supports the aggressive creature base.
- 1 Nighthowl Pursuer | threat | A black creature that keeps pressure on opponents.
- 1 Olog-hai Crusher | threat | A large creature threat for closing boards.
- 1 Orcish Siegemaster | threat | An Orc threat that strengthens the combat plan.
- 1 Ravening Warg | threat | A black creature threat that adds pressure at multiple stages.
- 1 Smaug, the Great Calamity // Spew Flame | wincon | A marked finisher that combines a major Dragon threat with removal.
- 1 Troll of Khazad-dûm | wincon | A marked finisher that provides a large closing threat.
- 1 Witch-king of Angmar | wincon | A marked finisher that closes games with a powerful evasive threat.
- 1 Bolg's Company | ramp | the mana pass added it to bring the mana base inside the power level

</details>

### 22. a set family and a card from outside it

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 175 names.

Commander: Thranduil, the Elvenking.

Grade: baseline, score 0.07, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $430.86 to buy, $430.86 the whole deck.

**Summary:** Thranduil leads a creature-centered Sultai deck built around Elves, with legendary Elves chief among them for recurring draw-and-discard triggers. Early mana creatures and artifacts help establish the board, while removal, interaction, and sweepers keep opponents from pulling too far ahead. The deck closes with Troll of Khazad-dûm and the Witch-kings after its Elf force has developed; it gives up some speed and consistency for a broad thematic creature package and a mana base that supports all three colors.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core is consistent — cheap Elves, legendary Elves to trigger the commander, ramp and card draw — but the black sweepers (Languish, Gnashing of Teeth) and the edict effect work against a board of small Elves the rest of the deck is trying to assemble.
- PLAN theme_fit=yes: It is a Commander deck drawn entirely from hob/hoc apart from the explicitly requested Sol Ring, led by Thranduil, the Elvenking in his Sultai identity, at a reasonable bracket-3 power level.
- PLAN useful_as_built=yes: 36 lands with ample basics in each color plus early mana creatures support a low curve, and the deck has enough draw, removal and creature-based pressure plus the Witch-kings and Troll to actually close games.
- PLAN summary_honest=yes: It accurately names the Elf/legendary-Elf core, the mana creatures and artifacts, the removal and sweepers, and the Troll/Witch-king finishers, and it openly concedes the speed and consistency costs of the three-color thematic build.
- [INFO] `curve_summary`: average mana value 3.02 over 63 nonland cards
- [INFO] `mana_pass`: the builder moved 5 cards of the mana base to bring the deck inside its power level
- [WARN] `outside_requested_set`: the sets you named do not hold 1 card: Sol Ring

<details><summary>The deck list</summary>

- 1 Sol Ring | ramp | Provides efficient mana acceleration and is retained as requested.
- 1 Arcane Signet | ramp | Provides mana acceleration for the deck’s colors.
- 1 Bejeweled Warg | ramp | Provides creature-based mana acceleration.
- 1 Elven Chorus | ramp | Provides Elf-focused mana acceleration.
- 1 Elvish Archdruid | ramp | Provides Elf-based mana acceleration.
- 1 Elvish Mystic | ramp | Provides early Elf-based mana acceleration.
- 1 Thranduil the Strategist | ramp | Provides ramp while adding a legendary Elf to the deck.
- 1 Wood Elves | ramp | Provides creature-based mana acceleration.
- 1 Woodland Weavemaster | ramp | Provides additional Elf-based mana acceleration.
- 1 Elvish Visionary | draw | Provides card draw on an Elf body.
- 1 Fateful Discovery | draw | Adds a dedicated card-draw spell.
- 1 Hithlain Knots | draw | Adds instant-speed card draw.
- 1 Ithilien Kingfisher | draw | Provides card draw from a creature.
- 1 Lórien Revealed | draw | Provides flexible card draw.
- 1 Night's Whisper | draw | Provides efficient card draw.
- 1 Palantír of Orthanc | draw | Provides a persistent draw-focused artifact.
- 1 Plunder the Trollshaws | draw | Adds instant-speed card draw.
- 1 Reverent Howl | draw | Provides additional card draw.
- 1 Uncover the Moon-Letters | draw | Adds a draw-focused enchantment.
- 1 Bilbo's Ring | interaction | Provides an interaction piece on a legendary artifact.
- 1 Confusticate and Bebother | interaction | Adds stack-focused interaction.
- 1 Elrond, Moon-Reader | interaction | Provides interaction on a legendary Elf body.
- 1 Mithril Coat | interaction | Provides protective interaction for important creatures.
- 1 My Precious // Allure of Power | interaction | Provides flexible interaction from a legendary artifact.
- 1 Old Fat Spider Can't See Me | interaction | Adds a dedicated interaction enchantment.
- 1 Stern Scolding | interaction | Provides efficient instant-speed interaction.
- 1 Bilbo's Deadly Slice | removal | Provides targeted removal.
- 1 Bitter Downfall | removal | Adds efficient instant-speed removal.
- 1 Crude Bent Blade | removal | Provides removal from an equipment slot.
- 1 Enchanted River's Grasp | removal | Adds enchantment-based removal.
- 1 Merciless Executioner | removal | Provides removal on a creature body.
- 1 Orcish Bowmasters | removal | Provides creature-based removal.
- 1 Quarrel | removal | Adds instant-speed removal.
- 1 Stir Up Trouble | removal | Provides another dedicated removal spell.
- 1 Uneasy Partings | removal | Adds flexible instant-speed removal.
- 1 Arwen, Weaver of Hope | synergy | Is a legendary Elf that supports Thranduil’s legendary-Elf draw trigger.
- 1 Boughside Wanderers | synergy | Adds an Elf card for Thranduil’s graveyard ability.
- 1 Cantankerous Keepers | synergy | Adds an Elf card for Thranduil’s graveyard ability.
- 1 Celeborn the Wise | synergy | Is a legendary Elf that supports Thranduil’s legendary-Elf draw trigger.
- 1 Galion, Elvenking's Butler | synergy | Is a legendary Elf that supports Thranduil’s legendary-Elf draw trigger.
- 1 Supper for Spiders | synergy | Provides a dedicated synergy card for the deck’s creature plan.
- 1 Thranduil, Sindarin Liege // Silvan Rally | synergy | Is a legendary Elf that supports Thranduil’s legendary-Elf draw trigger.
- 1 Elven Raft-Steerer | threat | Adds an Elf creature to pressure opposing life totals.
- 1 Elvenking's Harper | threat | Adds an Elf creature to the board-focused plan.
- 1 Galadhrim Guide | threat | Adds an Elf creature to the board-focused plan.
- 1 Grey Havens Navigator | threat | Adds an Elf creature to the board-focused plan.
- 1 Guardian of the Halls | threat | Adds an Elf creature to the board-focused plan.
- 1 Haunt of the Dead Marshes | threat | Adds an Elf creature to pressure opposing life totals.
- 1 Lothlórien Lookout | threat | Adds an Elf creature to the board-focused plan.
- 1 Mirkwood Meditator | threat | Adds an Elf creature to the board-focused plan.
- 1 Mirkwood Nurturer | threat | Adds an Elf creature to the board-focused plan.
- 1 Mirkwood Pathmaker | threat | Adds an Elf creature to the board-focused plan.
- 1 Nimrodel Watcher | threat | Adds an Elf creature to the board-focused plan.
- 1 Thranduil's Company | threat | Adds an Elf creature to the board-focused plan.
- 1 Troll of Khazad-dûm | wincon | Provides a marked finisher and win condition.
- 1 Witch-king of Angmar | wincon | Provides a marked finisher and win condition.
- 1 Witch-king, Bringer of Ruin | wincon | Provides a marked finisher and win condition.
- 1 Gnashing of Teeth | wipe | Provides a board-wipe effect.
- 1 Languish | wipe | Provides a board-wipe effect.
- 1 Raise the Palisade | wipe | Provides a board-wipe effect.
- 9 Forest | land | Provides green mana for the Elf-heavy core.
- 8 Island | land | Provides blue mana for draw and interaction.
- 10 Swamp | land | Provides black mana for removal and finishers.
- 1 Elven Passage | land | Adds an Elf-themed nonbasic land to the mana base.
- 1 Elvenking's Halls | land | Adds a Thranduil-themed nonbasic land to the mana base.
- 1 Hobbit Hole | land | Adds a nonbasic land to the mana base.
- 1 Minas Morgul, Dark Fortress | land | Adds a nonbasic land to the mana base.
- 1 Mirkwood | land | Adds a thematic nonbasic land to the mana base.
- 1 Rivendell | land | Adds a thematic nonbasic land to the mana base.
- 1 The Black Gate | land | Adds a thematic nonbasic land to the mana base.
- 1 The Shire | land | Adds a thematic nonbasic land to the mana base.
- 1 Treasure Vault | land | Adds a utility artifact land to the mana base.
- 1 Mox Amber | ramp | the mana pass added it to bring the mana base inside the power level
- 1 Delighted Halfling | ramp | the mana pass added it to bring the mana base inside the power level
- 1 Giant's Boulder | ramp | the mana pass added it to bring the mana base inside the power level

</details>

### 23. two set families at once

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 203 names.

Commander: Kíli the Resourceful.

Grade: typical, score 0.46, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $232.34 to buy, $232.34 the whole deck.

**Summary:** Kíli leads a white Dwarf-and-Equipment deck that builds a durable artifact board, keeps cards flowing through its many value pieces, and turns even modest creatures into meaningful attackers. The deck wins through equipped combat, supported by large evasive closers and artifact-based pressure. It gives up multicolor flexibility and depends on keeping creatures and Equipment together, so its protective spells and selective removal are important for maintaining momentum.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

This deck plan aims for 2 or more finishers, the cards that win the game, and the deck holds 1.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core of equipment, value creatures and Kíli pulls one way, but three board wipes (Dusk, Martial Coup, Promise of Loyalty) and off-plan pieces like Psychosis Crawler and Mentor of the Meek work against the deck's own wide, equipped creature board.
- PLAN theme_fit=yes: Every card is marked in the requested Hobbit/Bloomburrow set pool, the builder chose a commander as asked, and the mono-white value-and-Equipment build with Sol Ring and no fast combo sits comfortably in bracket 3.
- PLAN useful_as_built=yes: 37 lands plus five mana rocks support a low mono-white curve, and between equipped attackers, Helm of the Host, Andúril and the flying Angel/Dragon closers there are real ways to finish a game without any color or castability problems.
- PLAN summary_honest=partly: The equipment and value-draw claims are backed by the list and it fairly admits the fragility of creature-plus-Equipment, but calling it a "Dwarf" deck oversells roughly six Dwarves with almost no tribal payoff, and it never mentions the three sweepers that punish its own board.
- [INFO] `curve_summary`: average mana value 2.78 over 63 nonland cards
- [WARN] `finisher_short`: the finisher count is 1, and this deck plan wants 2 or more (Angel of the Ruins)

<details><summary>The deck list</summary>

- 29 Plains | land | Basic white land for a consistently white mana base.
- 1 Command Tower | land | Flexible land slot for the commander's color identity.
- 1 Castle Ardenvale | land | White land that supports the deck's creature-focused board.
- 1 Minas Tirith | land | Legendary white land that fits the Middle-earth theme.
- 1 Evolving Wilds | land | Land slot that helps stabilize access to Plains.
- 1 Fabled Passage | land | Land slot that helps stabilize access to Plains.
- 1 Path of Ancestry | land | The Dwarf creature theme makes this a fitting tribal land.
- 1 Thriving Heath | land | White land that keeps the mana base reliable.
- 1 Sol Ring | ramp | Efficient artifact ramp that also contributes to Kíli's artifact plan.
- 1 Arcane Signet | ramp | Reliable artifact ramp for a mono-white deck.
- 1 Fellwar Stone | ramp | Low-cost artifact ramp that helps reach the deck's larger threats.
- 1 Mind Stone | ramp | Early artifact ramp with value later in the game.
- 1 Wayfarer's Bauble | ramp | Early ramp that supports a Plains-heavy mana base.
- 1 Orcrist, Goblin-cleaver | ramp | Equipment-based ramp that supports the commander’s main theme.
- 1 Patchwork Banner | ramp | Artifact ramp that fits alongside the Dwarf package.
- 1 Ornithopter of Paradise | ramp | Artifact creature ramp that helps establish an artifact board.
- 1 Loyal Warhound | ramp | Creature ramp that helps the deck keep pace on mana.
- 1 Burnished Hart | ramp | Artifact creature ramp for building toward the top end.
- 1 Skullclamp | draw | Equipment-based card draw that belongs in the artifact package.
- 1 Mentor of the Meek | draw | Listed draw creature for a board built around smaller creatures.
- 1 Caretaker's Talent | draw | Listed draw enchantment for sustained value.
- 1 Circuit Mender | draw | Artifact creature that supplies card draw while advancing the artifact count.
- 1 Heirloom Epic | draw | Artifact draw piece that helps maintain resources.
- 1 Idol of Oblivion | draw | Low-cost artifact draw that supports the deck's permanent-heavy plan.
- 1 Inspiring Overseer | draw | Creature-based card draw that adds to the board.
- 1 Mangara, the Diplomat | draw | Listed draw threat that helps the deck avoid running out of cards.
- 1 Spirited Companion | draw | Efficient creature-based card draw.
- 1 The Gaffer | draw | Legendary draw creature that fits the Hobbit portion of the theme.
- 1 The Queen of Dale | draw | Legendary draw creature that supports the Middle-earth package.
- 1 Bilbo's Ring | interaction | Equipment interaction that remains on-theme with Kíli.
- 1 Bofur, Reliable Guardian // Concerted Care | interaction | Dwarf interaction that contributes to the creature theme.
- 1 Dwarven Mattock | interaction | Dwarf-themed Equipment interaction for the deck's central plan.
- 1 Galadriel's Dismissal | interaction | Flexible instant-speed interaction.
- 1 Luminous Broodmoth | interaction | Creature-based interaction that supports a resilient board.
- 1 Mithril Coat | interaction | Protective Equipment that helps preserve key creatures.
- 1 Reprieve | interaction | Low-cost instant interaction that buys time and replaces itself.
- 1 Swiftfoot Boots | interaction | Protective Equipment for Kíli and the deck's important threats.
- 1 Banishing Light | removal | Versatile permanent-based removal.
- 1 Bumbleflower's Sharepot | removal | Artifact removal that contributes to the artifact theme.
- 1 Fiend Hunter | removal | Creature-based removal that leaves a body for Equipment.
- 1 Generous Gift | removal | Broad instant-speed removal for problematic permanents.
- 1 Giant's Boulder | removal | Artifact removal that fits the deck's permanent suite.
- 1 Loran of the Third Path | removal | Legendary creature removal that supports the artifact-focused build.
- 1 Skyclave Apparition | removal | Efficient creature-based removal.
- 1 Swords to Plowshares | removal | Highly efficient creature removal.
- 1 The Black Arrow | removal | Removal Equipment that remains useful in Kíli's strategy.
- 1 Blade Splicer | synergy | Artifact-adjacent creature that adds material for the Equipment plan.
- 1 Carrot Cake | synergy | Artifact synergy piece from the Bloomburrow selection.
- 1 Dáin, Lord of the Iron Hills | synergy | Legendary Dwarf that strengthens the deck's Dwarf identity.
- 1 Iron Hills Blacksmith | synergy | Dwarf Artificer that directly fits both main creature and artifact themes.
- 1 Ori, Keeper of Songs | synergy | Legendary Dwarf that supports the deck's Dwarf core.
- 1 Tangle Tumbler | synergy | Artifact Vehicle that adds to the deck's artifact density.
- 1 Maskwood Nexus | synergy | Artifact tribal support for a list centered on Dwarves.
- 1 Fíli the Pathfinder | threat | Dwarf threat that is naturally aligned with Kíli's board plan.
- 1 Karn, the Great Creator | threat | Artifact-focused threat that leverages the deck's many artifacts.
- 1 Andúril, Flame of the West | threat | Legendary Equipment that turns a creature into a major combat threat.
- 1 Andúril, Narsil Reforged | threat | Legendary Equipment that adds another premium weapon to the deck.
- 1 Glamdring | threat | Legendary Equipment for the deck's concentrated Equipment package.
- 1 Helm of the Host | threat | Legendary Equipment that gives the deck a powerful board-building threat.
- 1 Long-Lost Lances | threat | Equipment threat that gives the creature suite more combat impact.
- 1 Starforged Sword | threat | Equipment threat that rewards maintaining creatures on board.
- 1 Sting, Bilbo's Sword | threat | Legendary Equipment that fits both the Hobbit and Equipment themes.
- 1 Sword of Vengeance | threat | Equipment threat for turning smaller creatures into serious attackers.
- 1 Sword of the Squeak | threat | Bloomburrow Equipment threat that stays within the requested sets.
- 1 Steelburr Champion | threat | Creature threat that can carry the deck's Equipment suite.
- 1 Angel of the Ruins | wincon | Finisher that provides a large evasive endgame threat.
- 1 Psychosis Crawler | wincon | Artifact creature that converts the deck's draw plan into a closing threat.
- 1 Sunscorch Regent | wincon | Large evasive creature that provides a decisive combat finish.
- 1 Dusk // Dawn | wipe | Board wipe that can reset creature-heavy opposing boards.
- 1 Martial Coup | wipe | Board wipe that also helps rebuild a creature board.
- 1 Promise of Loyalty | wipe | Asymmetric-style wipe for clearing away opposing boards while retaining a key creature.

</details>

### 24. a 60-card deck from one set

Format: Modern. Theme: aggro. Pool: any_card. Shortlist: 96 names.

Grade: baseline, score 0.35, model 20260914T154223Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $42.16 to buy, $42.16 the whole deck.

**Summary:** This mono-red aggro deck aims to establish creature pressure early, reinforce it with Emberheart Challenger and Hearthborn Battler, and clear resistance with direct removal. Its threats carry the game through combat, with Dragonhawk, Fate's Tempest chief among them, while its draw package helps keep pressure coming. The deck gives up broad answers and defensive options in favor of a direct, consistent red attacking plan.

The quality model grades this deck at the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. The color sources cover the pips, and that raises the grade.

- PLAN plan_coherent=partly: The removal and card-draw pieces are fine red support, but the deck's body is a stack of four- and five-mana creatures plus two Arcane Signets, which is a midrange ramp shell rather than the single early-pressure aggro plan the summary names, with only four two-drop creatures to open on.
- PLAN theme_fit=partly: It is genuinely mono-red, Modern-legal, and drawn entirely from the requested Bloomburrow sets, but the build lands as slow red midrange rather than the aggro deck that was asked for.
- PLAN useful_as_built=partly: With 24 Mountains and a single color the mana works and the deck does have real threats and removal, so it is playable at a casual table, but the top-heavy curve and five-mana Conduct Electricity mean it frequently stalls in the early turns it is supposed to win in.
- PLAN summary_honest=partly: The removal, draw, and Dragonhawk claims are all accurate to the list, but calling it a deck that 'establishes creature pressure early' hides a curve topped by twelve five-mana-ish threats and quietly omits the two Arcane Signets that exist precisely because the deck is too slow to be aggro.
- [INFO] `curve_summary`: average mana value 3.33 over 36 nonland cards

<details><summary>The deck list</summary>

- 2 Artist's Talent | draw | Provides the deck's red card-draw support.
- 2 Might of the Meek | draw | Adds efficient card draw while fitting the aggressive plan.
- 2 Sazacap's Brew | draw | Supplies additional card draw in mono-red.
- 24 Mountain | land | Provides reliable untapped red mana for the deck.
- 2 Arcane Signet | ramp | Accelerates the deck's mana production.
- 4 Agate Assault | removal | Forms the primary package of opposing-card removal.
- 2 Conduct Electricity | removal | Rounds out the deck's removal suite.
- 4 Emberheart Challenger | synergy | Supports the deck's aggressive creature synergies.
- 4 Hearthborn Battler | synergy | Adds another focused synergy piece for the creature plan.
- 4 Dragonhawk, Fate's Tempest | threat | Provides a powerful threat at the top of the creature curve.
- 4 Frilled Sparkshooter | threat | Supplies a full set of aggressive threats.
- 4 Reptilian Recruiter | threat | Adds more creatures to maintain pressure.
- 2 Teapot Slinger | threat | Completes the creature threat package.

</details>

### 25. use no card of an owned precon

Format: Commander. Theme: Avengers superheroes. Pool: owned_first. Shortlist: 195 names.

Commander: Captain America, Super-Soldier.

Grade: baseline, score 0.35, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $153.62 the whole deck.

**Summary:** Captain America, Super-Soldier anchors a white Avengers and S.H.I.E.L.D. Hero shell that develops a protected combat board, uses Equipment to keep the cards and mana flowing, and presses its advantage with superhero threats. The deck closes through Angel of Serenity, Origin of Spider-Man, and The Sentry, Golden Guardian, while its removal and wipes keep opposing boards from taking over. It gives up multicolor flexibility and depends heavily on maintaining creatures and artifacts on the table.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core plan — a wide white Hero board with Equipment for card flow and heavy protection effects (Bastion Protector, Clever Concealment, Darksteel Plate, Boots/Greaves) — is consistent, but three symmetric sweepers (Austere Command, Fumigate, Vanquish the Horde) pull against a deck that wins by keeping its own creatures and artifacts on the table.
- PLAN theme_fit=yes: It is a mono-white bracket-3 Commander deck led by Captain America, Super-Soldier with a large Marvel superhero core (Agents of S.H.I.E.L.D., Vision, Luke Cage, Okoye, Invisible Woman, Mockingbird, Sentry), with only a modest amount of generic/LOTR filler to fill the shell.
- PLAN useful_as_built=yes: 36 mono-white lands with eight mana rocks and a low, mostly two-and-three-drop curve cast reliably, and the deck has ample removal plus a real combat-damage finish backed by Angel of Serenity and the Sentry.
- PLAN summary_honest=yes: Every element the summary names is present in the list — the white Hero/S.H.I.E.L.D. creatures, the Equipment draw and ramp package, the named closers, the spot removal and the wipes — and it openly admits the mono-color and board-dependence drawbacks.
- [INFO] `curve_summary`: average mana value 3.00 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Plains | land | Provides the deck's basic white mana base.
- 1 Plains | land | Provides the deck's basic white mana base.
- 1 Plains | land | Provides the deck's basic white mana base.
- 1 Plains | land | Provides the deck's basic white mana base.
- 1 Plains | land | Provides the deck's basic white mana base.
- 1 Plains | land | Provides the deck's basic white mana base.
- 1 Plains | land | Provides the deck's basic white mana base.
- 1 Plains | land | Provides the deck's basic white mana base.
- 1 Plains | land | Provides the deck's basic white mana base.
- 1 Plains | land | Provides the deck's basic white mana base.
- 1 Plains | land | Provides the deck's basic white mana base.
- 1 Plains | land | Provides the deck's basic white mana base.
- 1 Plains | land | Provides the deck's basic white mana base.
- 1 Plains | land | Provides the deck's basic white mana base.
- 1 Plains | land | Provides the deck's basic white mana base.
- 1 Plains | land | Provides the deck's basic white mana base.
- 1 Plains | land | Provides the deck's basic white mana base.
- 1 Plains | land | Provides the deck's basic white mana base.
- 1 Plains | land | Provides the deck's basic white mana base.
- 1 Plains | land | Provides the deck's basic white mana base.
- 1 Plains | land | Provides the deck's basic white mana base.
- 1 Plains | land | Provides the deck's basic white mana base.
- 1 Plains | land | Provides the deck's basic white mana base.
- 1 Plains | land | Provides the deck's basic white mana base.
- 1 Plains | land | Provides the deck's basic white mana base.
- 1 Plains | land | Provides the deck's basic white mana base.
- 1 Plains | land | Provides the deck's basic white mana base.
- 1 Plains | land | Provides the deck's basic white mana base.
- 1 Command Tower | land | A reliable land for Captain America's colors.
- 1 City of Brass | land | Adds flexible colored mana to the land base.
- 1 Exotic Orchard | land | Adds flexible colored mana to the land base.
- 1 Secluded Courtyard | land | Supports the deck's Hero creature core.
- 1 Path of Ancestry | land | A tribal-minded land for the Hero-heavy roster.
- 1 Minas Tirith | land | Adds a themed legendary land slot to the mana base.
- 1 Windbrisk Heights | land | Adds another white land to the mana base.
- 1 Evolving Wilds | land | Helps stabilize the basic-land mana base.
- 1 Sol Ring | ramp | An efficient early mana piece.
- 1 Arcane Signet | ramp | Reliable colored acceleration.
- 1 Wayfarer's Bauble | ramp | Early ramp that works with the basic-land base.
- 1 Sword of the Animist | ramp | Equipment-based ramp that fits the combat plan.
- 1 Thought Vessel | ramp | A simple mana artifact for developing the board.
- 1 Chromatic Lantern | ramp | Provides mana acceleration and color support.
- 1 Commander's Sphere | ramp | Flexible mana acceleration.
- 1 White Lotus Tile | ramp | A white-themed artifact ramp piece.
- 1 Bender's Waterskin | ramp | An additional artifact mana source.
- 1 Giada, Font of Hope | ramp | A creature-based ramp piece that also supports Angel inclusions.
- 1 Agent 13, Sharon Carter | draw | A Hero that supplies card advantage.
- 1 Agent Maria Hill | draw | A Hero that supplies card advantage.
- 1 Champions of Minas Tirith | draw | A Soldier-based source of card advantage.
- 1 Buster Sword | draw | Equipment-based card advantage for attacking creatures.
- 1 Crown of Gondor | draw | A legendary Equipment source of card advantage.
- 1 Hero in Training | draw | A Hero that helps keep cards flowing.
- 1 Idol of Oblivion | draw | A compact artifact draw piece.
- 1 Mask of Memory | draw | Equipment-based card advantage for the combat plan.
- 1 The Vision | draw | A Robot Hero that supplies card advantage.
- 1 Viv Vision, Teen Synthezoid | draw | A Robot Hero that supplies card advantage.
- 1 Puresteel Paladin | draw | Supports the Equipment package while drawing cards.
- 1 Skullclamp | draw | An efficient Equipment source of card advantage.
- 1 Weapons Vendor | draw | Supports the Equipment-heavy build with card advantage.
- 1 Lembas | draw | A compact artifact draw piece.
- 1 Bastion Protector | interaction | Helps protect Captain America and the Hero board.
- 1 Boromir, Warden of the Tower | interaction | A creature-based protective interaction piece.
- 1 Champion's Helm | interaction | Protective Equipment for Captain America or a key Hero.
- 1 Clever Concealment | interaction | Protective interaction for a developed board.
- 1 Darksteel Plate | interaction | Durable protective Equipment for a key creature.
- 1 Duty Beyond Death | interaction | A flexible protective response.
- 1 Frontline Medic | interaction | A creature-based interaction piece for combat turns.
- 1 Lightning Greaves | interaction | Efficient protective Equipment for Captain America or a Hero.
- 1 Patriot, Shield Wielder | interaction | A Hero that supports the deck's protective plan.
- 1 Sheltered by Ghosts | interaction | Protective interaction that fits the creature plan.
- 1 Swiftfoot Boots | interaction | Protective Equipment for important creatures.
- 1 Take Up the Shield | interaction | A compact protective combat trick.
- 1 Angel of Sanctions | removal | A creature-based answer that can join the battlefield.
- 1 Banishing Light | removal | A versatile permanent answer.
- 1 Crib Swap | removal | A targeted removal spell.
- 1 Dispatch | removal | A low-cost targeted answer for the artifact build.
- 1 Fiend Hunter | removal | A creature-based removal option.
- 1 Generous Gift | removal | A flexible answer to opposing permanents.
- 1 Get Lost | removal | An efficient targeted answer.
- 1 Journey to Nowhere | removal | An enchantment-based removal option.
- 1 Swords to Plowshares | removal | An efficient targeted removal spell.
- 1 Agent Phil Coulson | synergy | A S.H.I.E.L.D. Hero that reinforces the Avengers theme.
- 1 Agent of Atlas | synergy | A Hero inclusion for the superhero roster.
- 1 Agents of S.H.I.E.L.D. | synergy | Strengthens the S.H.I.E.L.D. side of the Hero shell.
- 1 Avengers Quinjet | synergy | A direct Avengers-themed support card.
- 1 Colleen Wing, Street Samurai | synergy | A Hero inclusion that supports the deck's theme.
- 1 Night Nurse, Healer of Heroes | synergy | A Hero-focused support card.
- 1 Silver Sable, Mercenary Leader | synergy | A superhero-themed support card for the Hero roster.
- 1 Captain Mar-Vell, Space-Born | threat | An Avengers-themed Hero threat.
- 1 Invisible Woman, Sue Storm | threat | A superhero Hero threat for the battlefield.
- 1 Luke Cage, Power Man | threat | A Hero threat that advances the combat plan.
- 1 Mockingbird, Ace Agent | threat | A S.H.I.E.L.D.-aligned Hero threat.
- 1 Okoye, Dora Milaje Leader | threat | A Hero threat that adds to the attacking roster.
- 1 Angel of Serenity | wincon | A marked finisher for closing longer games.
- 1 Origin of Spider-Man | wincon | A superhero-themed marked finisher.
- 1 The Sentry, Golden Guardian | wincon | A Hero-themed marked finisher.
- 1 Austere Command | wipe | A flexible reset button when the board gets away.
- 1 Fumigate | wipe | A dependable battlefield reset.
- 1 Vanquish the Horde | wipe | A board wipe for clearing opposing creatures.

</details>

