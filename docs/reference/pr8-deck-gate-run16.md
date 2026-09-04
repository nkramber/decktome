# PR-8 deck gate

Run date: 2026-09-04. Card snapshot: 2026-09-03.

Verdict: PASS. 25 of 25 decks passed every block check, 0 invented names reached the user, and 0 summaries stated a false rule of the game. All three bars are zero tolerance.

## Summary

| Measure | Value |
|---|---|
| Prompts | 25 |
| Decks returned | 25 |
| Decks with no block finding | 25 |
| Invented names that reached the user | 0 |
| Decks that needed the repair turn | 12 |
| Summaries judged (F-26) | 25 |
| Summaries that state a rule of the game | 0 |
| Summaries that state a FALSE rule | 0 |
| Judge errors | 0 |
| Decks the plan judge read (PR-15, information) | 25 |
| Mean plan score, 0 to 1 | 0.69 |
| Errors | 0 |
| Prompt version | 12 |
| Calls | 93 |
| Cost | $3.7874 |
| Time | 3018 seconds |

## Run

- Suite `decks`, run `pr8-deck-gate-run16`, on 2026-09-04, commit `ec058b4`.
- Roles: generate on `gpt-5.6-terra` (openai, effort medium), judge on `claude-opus-5` (anthropic, effort medium), repair on `gpt-5.6-terra` (openai, effort medium).
- Versions: card snapshot 2026-09-03, generate prompt version 12, plan_rubric prompt version 1, precons `5.3.0+20260903`, quality_model `20260903T210658Z`.
- Calls: 93. Cost: $3.7874. Time: 3018 seconds.

## Findings by code

A block stops the deck. A warning and a note are reports.

| Code | Count |
|---|---|
| `curve_summary` | 25 |
| `not_owned` | 25 |
| `profile_off_band` | 10 |
| `outside_requested_set` | 2 |
| `precon_cards_restored` | 1 |
| `repair_kept_earlier` | 1 |

By severity: BLOCK 0. WARN 37. INFO 27. 

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

Grade: bad, score 0.02, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $500.76 to buy, $500.76 the whole deck.

**Summary:** Karlov leads a white-black lifegain deck that develops a resilient creature board, turns repeated life gain into larger threats and pressure, and uses efficient protection and removal to keep its key pieces in play. It wins by snowballing its lifegain synergies into a commanding battlefield, while retaining several board resets for games that get away from it. The deck gives up some speed for a fuller land base and a more creature-focused midgame plan.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The bulk of the list is a unified white-black lifegain creature build (payoffs like Ajani's Pridemate, Heliod, Archangel of Thune, Vito, plus lifelink drain threats), but three symmetrical board wipes in a deck whose whole plan is building a wide creature board pull against the rest of the cards.
- PLAN theme_fit=yes: It is exactly what was asked: a 99-card white-black Karlov of the Ghost Council lifegain Commander deck with power consistent with bracket 3 (no fast mana, no infinite combos, just synergy cards and a couple of premium staples).
- PLAN useful_as_built=yes: 36 lands with duals and fixing, ten ramp pieces, eleven draw sources, and many redundant lifegain payoffs and combat threats mean the deck casts its spells and closes games out of the box.
- PLAN summary_honest=yes: It accurately describes the lifegain-into-threats plan, openly names the board resets and the trade of speed for a heavier land count, and claims no combo or tutor package the list lacks.
- [INFO] `curve_summary`: average mana value 3.37 over 63 nonland cards

<details><summary>The deck list</summary>

- 13 Plains | land | Basic white source for the mana base.
- 11 Swamp | land | Basic black source for the mana base.
- 1 Godless Shrine | land | Dual-color land for reliable white and black access.
- 1 Caves of Koilos | land | Untapped dual-color land.
- 1 Command Tower | land | Reliable land for the commander’s colors.
- 1 Vault of Champions | land | Dual-color land for multiplayer games.
- 1 Shattered Sanctum | land | Dual-color land that supports the early curve.
- 1 Isolated Chapel | land | Dual-color land for the white-black mana base.
- 1 Mana Confluence | land | Flexible source for either deck color.
- 1 Reflecting Pool | land | Flexible dual-color land alongside the rest of the mana base.
- 1 Radiant Fountain | land | Land slot that supports the lifegain plan.
- 1 Seraph Sanctuary | land | Land slot that supports the Angel portion of the deck.
- 1 Restless Fortress | land | Dual-color utility land.
- 1 High Market | land | Utility land that supplies a sacrifice outlet.
- 1 Altar of the Pantheon | ramp | Early mana acceleration for the deck’s colored spells.
- 1 Battle Angels of Tyr | ramp | Threat-based acceleration that advances the board.
- 1 Carmen, Cruel Skymarcher | ramp | Creature-based mana acceleration for the midgame.
- 1 Crypt Ghast | ramp | High-impact black mana acceleration.
- 1 Nuka-Cola Vending Machine | ramp | Artifact ramp that helps develop resources.
- 1 Orazca Relic | ramp | Mana rock for consistent development.
- 1 Pristine Talisman | ramp | Mana acceleration that fits the lifegain theme.
- 1 Phial of Galadriel | ramp | Artifact ramp for the deck’s midgame.
- 1 The Celestus | ramp | Versatile artifact mana source.
- 1 Spoils of Evil | ramp | Efficient burst mana acceleration.
- 1 Archivist of Oghma | draw | Efficient creature-based card advantage.
- 1 Ayara, First of Locthwain | draw | Repeatable card advantage on a relevant creature.
- 1 Dawn of Hope | draw | Lifegain-focused source of card advantage.
- 1 Enduring Innocence | draw | Creature-based draw engine.
- 1 Mangara, the Diplomat | draw | Reliable multiplayer card advantage.
- 1 Markov Purifier | draw | Vampire-themed card advantage.
- 1 Sigarda's Splendor | draw | Lifegain-aligned card advantage.
- 1 The Gaffer | draw | Creature-based draw that fits the life-total plan.
- 1 Vampiric Rites | draw | Low-cost draw outlet with sacrifice utility.
- 1 Well of Lost Dreams | draw | Dedicated lifegain card-advantage engine.
- 1 Survival Cache | draw | Lifegain-oriented card advantage.
- 1 Alseid of Life's Bounty | interaction | Low-cost protection for Karlov or a key permanent.
- 1 Faith's Shield | interaction | Efficient protective interaction.
- 1 Courageous Resolve | interaction | Protection spell for the board’s key pieces.
- 1 Metropolis Reformer | interaction | Creature-based defensive interaction.
- 1 Restoration Magic | interaction | Flexible protective interaction.
- 1 Rune-Tail, Kitsune Ascendant // Rune-Tail's Essence | interaction | Defensive piece for protecting the creature board.
- 1 Sword of Light and Shadow | interaction | Equipment-based protection and utility.
- 1 Werefox Bodyguard | interaction | Creature interaction that supports the board plan.
- 1 Ayli, Eternal Pilgrim | removal | Low-cost repeatable removal on a lifegain creature.
- 1 Gumdrop Poisoner // Tempt with Treats | removal | Flexible removal with an additional spell option.
- 1 Henrika Domnathi // Henrika, Infernal Seer | removal | Versatile creature-based removal.
- 1 Murderous Rider // Swift End | removal | Flexible targeted removal.
- 1 Nightmare's Thirst | removal | Cheap targeted creature removal.
- 1 Solitude | removal | Flexible creature removal.
- 1 Umezawa's Jitte | removal | Reusable equipment-based removal.
- 1 Vona, Butcher of Magan | removal | Lifegain-compatible repeatable removal.
- 1 Cavalier of Night | removal | Creature-based removal for the top end.
- 1 Fumigate | wipe | Board wipe that aligns with the life-total plan.
- 1 Kaya's Wrath | wipe | Reliable white-black board wipe.
- 1 The Battle of Bywater | wipe | Additional reset button for opposing boards.
- 1 Ajani's Pridemate | synergy | Low-cost lifegain payoff.
- 1 Angel of Vitality | synergy | Lifegain-focused creature synergy.
- 1 Cleric Class | synergy | Early enchantment that supports repeated lifegain.
- 1 Heliod, Sun-Crowned | synergy | Central lifegain payoff and support piece.
- 1 Indulging Patrician | synergy | Creature payoff for the lifegain strategy.
- 1 Ocelot Pride | synergy | Low-cost creature that rewards the lifegain plan.
- 1 Resplendent Angel | synergy | Powerful lifegain payoff threat.
- 1 Righteous Valkyrie | synergy | Angel-focused lifegain synergy.
- 1 Vito, Thorn of the Dusk Rose | synergy | Lifegain payoff that pressures opponents.
- 1 Voice of the Blessed | synergy | Efficient scalable lifegain payoff.
- 1 Archangel of Thune | threat | High-impact creature threat for a lifegain board.
- 1 Astarion, the Decadent | threat | Large lifegain-themed finisher.
- 1 Attended Healer | threat | Board-building creature threat.
- 1 Celestine, the Living Saint | threat | Recurring value threat for longer games.
- 1 Cliffhaven Vampire | threat | Lifegain-themed source of pressure.
- 1 Defiant Bloodlord | threat | High-impact lifegain payoff threat.
- 1 Divinity of Pride | threat | Efficient large threat for high life totals.
- 1 Enduring Tenacity | threat | Lifegain-themed creature threat.
- 1 Nykthos Paragon | threat | Powerful lifegain payoff at the top of the curve.
- 1 Rhox Faithmender | threat | Lifegain multiplier that becomes a substantial board presence.
- 1 Valkyrie Harbinger | threat | Large lifegain-themed Angel threat.
- 1 Wurmcoil Engine | threat | Resilient artifact creature threat.

</details>

### 2. aristocrats Commander, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 179 names.

Commander: Denethor, Ruling Steward.

Grade: typical, score 0.45, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band. Block findings: 0.

Cost: $0.00 to buy, $186.10 the whole deck.

**Summary:** A white-black aristocrats deck centered on expendable creatures, graveyard value, and steady attrition. It develops a resilient board, converts losses into resources, protects its key engine pieces, and closes games through evasive legends and creature-driven pressure.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=no: The list mixes a sacrifice/attrition shell with only a couple of true sac outlets and almost no token or expendable-body generation, while simultaneously running three board wipes (Austere Command, Fumigate, Dusk // Dawn) plus a pile of spot removal and equipment, so the creature-death engine and the sweeper-control plan pull against each other.
- PLAN theme_fit=no: The requester asked for an aristocrats deck, and the build is mostly generic removal, wipes, card-draw rocks, and LOTR/other legends with little sacrifice payoff density; worse, Denethor, Ruling Steward is mono-black, so the Plains and the entire white half of the list fall outside the commander's color identity.
- PLAN useful_as_built=no: As written the deck is not legal to sleeve up — roughly half the nonland cards and 14 Plains sit outside a mono-black commander's identity — and even setting that aside, the sacrifice payoffs lack the fodder to fire consistently.
- PLAN summary_honest=no: It advertises "expendable creatures," "converts losses into resources," and a protected engine, but the deck has few sacrifice outlets, no fodder producers, minimal recursion, and it calls the deck white-black without flagging that the commander's identity does not support the white half.
- [INFO] `curve_summary`: average mana value 2.87 over 63 nonland cards

<details><summary>The deck list</summary>

- 14 Plains | land | Reliable white mana.
- 14 Swamp | land | Reliable black mana.
- 1 Command Tower | land | Unrestricted commander-color fixing.
- 1 City of Brass | land | Flexible colored mana.
- 1 Exotic Orchard | land | Multiplayer color fixing.
- 1 Grand Coliseum | land | Additional color fixing.
- 1 Minas Tirith | land | White source with late-game value.
- 1 Plaza of Heroes | land | Colored mana and legendary support.
- 1 Spire of Industry | land | Artifact-enabled color fixing.
- 1 Takenuma, Abandoned Mire | land | Black source with graveyard utility.
- 1 Sol Ring | ramp | Efficient colorless acceleration.
- 1 Arcane Signet | ramp | Reliable colored acceleration.
- 1 Fellwar Stone | ramp | Efficient multiplayer fixing.
- 1 Thought Vessel | ramp | Colorless acceleration with hand support.
- 1 Wayfarer's Bauble | ramp | Permanent-based land acceleration.
- 1 Commander's Sphere | ramp | Fixing that can later become a card.
- 1 Chromatic Lantern | ramp | Broad mana fixing.
- 1 Relic of Legends | ramp | Acceleration that leverages legendary creatures.
- 1 Springleaf Drum | ramp | Turns expendable creatures into mana.
- 1 Deadly Dispute | ramp | Converts a disposable permanent into mana and cards.
- 1 Skullclamp | draw | Turns small creatures into efficient card flow.
- 1 Night's Whisper | draw | Low-cost card selection.
- 1 Painful Truths | draw | Efficient burst of cards.
- 1 Idol of Oblivion | draw | Repeatable draw alongside creature tokens.
- 1 Lembas | draw | Early selection with a later card return.
- 1 Mask of Memory | draw | Combat-based filtering and card advantage.
- 1 Tome of Legends | draw | Steady cards from commander activity.
- 1 Wall of Omens | draw | Early defense that replaces itself.
- 1 Inspiring Overseer | draw | Creature body with immediate card value.
- 1 Call of the Ring | draw | Recurring card flow for the creature plan.
- 1 Grave Venerations | draw | Graveyard-oriented card advantage.
- 1 Lightning Greaves | interaction | Protects key creatures immediately.
- 1 Swiftfoot Boots | interaction | Repeatable protection for important bodies.
- 1 Darksteel Plate | interaction | Durable creature protection.
- 1 Gift of Immortality | interaction | Keeps a valuable creature available through removal.
- 1 Clever Concealment | interaction | Protects the board from hostile effects.
- 1 Reprieve | interaction | Temporarily answers a pivotal spell while cycling.
- 1 Unbreakable Formation | interaction | Safeguards the creature board.
- 1 Together Forever | interaction | Preserves creatures through the graveyard.
- 1 Swords to Plowshares | removal | Efficient creature answer.
- 1 Generous Gift | removal | Answers any problematic permanent.
- 1 Get Lost | removal | Flexible answer to major nonland permanents.
- 1 Bitter Triumph | removal | Broad instant-speed creature or planeswalker answer.
- 1 Fatal Push | removal | Low-cost creature removal.
- 1 Heartless Act | removal | Flexible instant-speed creature removal.
- 1 Infernal Grasp | removal | Unconditional creature removal.
- 1 Fiend Hunter | removal | Creature-based exile effect.
- 1 Claim the Precious | removal | Clean answer to dangerous creatures.
- 1 Austere Command | wipe | Modular reset for troublesome boards.
- 1 Dusk // Dawn | wipe | Creature reset with graveyard value.
- 1 Fumigate | wipe | Creature sweep that restores life.
- 1 Gollum the Abandoned | synergy | A disposable creature that supports death-focused play.
- 1 Gollum, Patient Plotter | synergy | Repeatedly returns to support sacrifice lines.
- 1 Gríma Wormtongue | synergy | Supports the deck's attrition strategy.
- 1 Nimble Hobbit | synergy | Low-cost creature for sacrifice and board development.
- 1 Phantom Train | synergy | Artifact body that contributes to the sacrifice plan.
- 1 Arcade Cabinet | synergy | Artifact value piece for the permanent-heavy build.
- 1 Stone of Erech | synergy | Supports graveyard pressure and attrition.
- 1 Contagion Clasp | synergy | Works with counters and incremental creature control.
- 1 Blowfly Infestation | synergy | Turns creature deaths into continuing board pressure.
- 1 The Sackville-Bagginses | synergy | Creature-based value that fits the death-and-token plan.
- 1 Bill the Pony | threat | Efficient creature pressure with thematic support.
- 1 Bronze Guardian | threat | Large resilient artifact threat.
- 1 Gaius van Baelsar | threat | Legendary creature that applies meaningful pressure.
- 1 Massacre Girl, Known Killer | threat | Punishes opposing creature boards while advancing attrition.
- 1 Orcish Bowmasters | threat | Early board presence with recurring pressure.
- 1 Rat King, Pale Piper | threat | Creature threat with graveyard-oriented value.
- 1 The Walls of Ba Sing Se | threat | Substantial artifact creature for board control.
- 1 Witch-king of Angmar | threat | Evasive finisher that punishes attacks.
- 1 Frontline Medic | threat | Creature pressure that also improves combat resilience.
- 1 Palace Jailer | threat | Adds pressure and sustained value while disrupting a creature.
- 1 Exemplar of Light | threat | Evasive creature that provides value while attacking.
- 1 Archfiend of Ifnir | threat | Large evasive threat that constrains opposing creatures.

</details>

### 3. artifacts Commander, bracket 4

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 299 names.

Commander: Urza, Lord High Artificer.

Grade: bad, score 0.04, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band, profile_off_band; then profile_off_band, profile_off_band. Block findings: 0.

Cost: $1037.53 to buy, $1037.53 the whole deck.

**Summary:** This high-power mono-blue artifact strategy deploys inexpensive artifacts quickly, turns them into mana and card velocity, and uses flexible permission to protect its strongest sequences. Its board develops through cost reduction, recursion, tutoring, and artifact synergies before ending games with heavy artifact pressure and scalable threats.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The commander leads high-bracket decks, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every card supports one artifact plan — cheap artifacts, cost reducers and mana-from-artifacts effects, affinity/improvise payoffs, artifact recursion and tutors, and artifact-based bodies — with no competing subtheme pulling the list apart.
- PLAN theme_fit=partly: Commander, colors, format and artifact theme match the request exactly, but the requested bracket-4 power level is only half-met: a few strong pieces (Mox Opal, KCI, Metalworker, Mystic Forge, Urza's Saga) sit alongside a lot of low-impact filler and weak removal like Water Whip and Unable to Scream.
- PLAN useful_as_built=yes: 33 lands plus abundant artifact ramp support a very low curve, mono-blue mana is trouble-free, and Urza's construct token, Metalworker/KCI mana and a dozen artifact creatures give real ways to close a game.
- PLAN summary_honest=partly: The claims of cost reduction, recursion (Academy Ruins, Emry, Mystic Sanctuary), tutoring (Whir, Transmute Artifact, Urza's Saga, Inventors' Fair) and permission are all genuinely present, but calling it 'high-power' with 'heavy artifact pressure' oversells a finish that leans on modest bodies like Panther Robot, Chrome Steed and Memory Guardian.
- [INFO] `curve_summary`: average mana value 3.23 over 66 nonland cards
- [WARN] `profile_off_band`: the average mana value of the nonland cards is 3.23, and bracket 4 wants 1.6 to 3
- [WARN] `profile_off_band`: the mana available on turn four is 4.34, and bracket 4 wants 4.8 or more (mean over 10000 hands)

<details><summary>The deck list</summary>

- 22 Island | land | Reliable blue mana base for an artifact-heavy strategy.
- 1 Academy Ruins | land | Recurs key artifacts after they are destroyed.
- 1 Command Tower | land | Provides dependable colored mana.
- 1 Exotic Orchard | land | Usually supplies blue mana in multiplayer games.
- 1 Glimmervoid | land | Colored source supported by the deck's artifact density.
- 1 Inventors' Fair | land | Finds an important artifact while offering incidental life gain.
- 1 Mystic Sanctuary | land | Reuses an important instant or sorcery.
- 1 Otawara, Soaring City | land | Blue source with flexible channel interaction.
- 1 Seat of the Synod | land | Artifact land that improves affinity, improvise, and artifact count.
- 1 Sink into Stupor // Soporific Springs | land | Flexible blue land that can also answer a problematic permanent.
- 1 Spire of Industry | land | Reliable blue source in an artifact deck.
- 1 Urza's Saga | land | Makes artifact constructs and finds a low-cost utility artifact.
- 1 Mox Opal | ramp | Fast artifact-based acceleration.
- 1 Moonsnare Prototype | ramp | Early mana acceleration that uses artifacts and creatures.
- 1 Arc Reactor | ramp | Artifact-based mana acceleration.
- 1 Chief Engineer | ramp | Convoke-style acceleration for artifact spells.
- 1 Crystal Skull, Isu Spyglass | ramp | Early mana rock with useful late-game selection.
- 1 Grand Architect | ramp | Turns blue creatures into substantial artifact mana.
- 1 Inspiring Statuary | ramp | Uses spare artifacts to help cast larger spells.
- 1 Jetfire, Ingenious Scientist // Jetfire, Air Guardian | ramp | Artifact threat that also supports mana development.
- 1 Krark-Clan Ironworks | ramp | Converts artifacts into explosive mana for decisive turns.
- 1 Metalworker | ramp | Produces large amounts of mana from artifact-heavy hands.
- 1 Slagstone Refinery | ramp | Provides repeatable artifact mana production.
- 1 Solar Array | ramp | Supports the deck's mana development.
- 1 Vedalken Engineer | ramp | Efficient colored acceleration for artifact spells.
- 1 Braided Net // Braided Quipu | draw | Artifact card advantage that supports the deck's core mechanics.
- 1 Cerebral Download | draw | Efficient spell-based card advantage.
- 1 Era of Innovation | draw | Rewards artifact development with card selection and cards.
- 1 Esoteric Duplicator | draw | Artifact-based source of recurring card advantage.
- 1 Forensic Gadgeteer | draw | Generates value while advancing the artifact plan.
- 1 Reverse Engineer | draw | Refills the hand efficiently with artifact support.
- 1 Riddlesmith | draw | Filters cards whenever artifacts are deployed.
- 1 Sai, Master Thopterist | draw | Builds a Thopter board while converting artifacts into cards.
- 1 Thirst for Knowledge | draw | Instant-speed card draw that benefits from spare artifacts.
- 1 Thought Monitor | draw | Affinity makes this a powerful artifact-based refill.
- 1 Thoughtcast | draw | Highly efficient card draw with a developed artifact board.
- 1 Transmutation Font | draw | Provides repeatable artifact-based card advantage.
- 1 Welding Jar | interaction | Protects a crucial artifact for no mana.
- 1 Disruption Protocol | interaction | Artifact-assisted countermagic.
- 1 Fugitive Droid | interaction | Artifact creature that supplies disruptive utility.
- 1 Ghostly Flicker | interaction | Protects permanents and reuses valuable enter-the-battlefield effects.
- 1 Ice Out | interaction | Flexible counterspell that becomes efficient with artifacts.
- 1 Metallic Rebuke | interaction | Efficient improvise countermagic.
- 1 Override | interaction | Low-cost permission for protecting key turns.
- 1 Padeem, Consul of Innovation | interaction | Protects important artifacts from opposing interaction.
- 1 Reality Ripple | interaction | Flexible instant-speed disruption.
- 1 Stoic Rebuttal | interaction | Reliable artifact-enabled counterspell.
- 1 Curator's Ward | interaction | Protects a key permanent while replacing itself.
- 1 Aether Spellbomb | removal | Cheap, reusable creature interaction that remains an artifact.
- 1 Contagion Clasp | removal | Early removal with proliferate utility.
- 1 Cyber Conversion | removal | Efficient answer to a troublesome creature.
- 1 Into Thin Air | removal | Flexible tempo-oriented permanent answer.
- 1 Kitesail Larcenist | removal | Turns opposing permanents into manageable artifacts.
- 1 Lumengrid Drake | removal | Artifact-supported creature removal.
- 1 Ravenform | removal | Exiles a creature or artifact cleanly.
- 1 Resculpt | removal | Exiles problematic artifacts or creatures at instant speed.
- 1 Unable to Scream | removal | Cheap permanent-based creature shutdown.
- 1 Water Whip | removal | Versatile removal that fits the deck's blue identity.
- 1 Engineered Explosives | wipe | Scalable board control with a minimal mana investment.
- 1 Hurkyl's Recall | wipe | Sweeps opposing artifacts and can create a major tempo swing.
- 1 Emry, Lurker of the Loch | synergy | Recasts inexpensive artifacts from the graveyard.
- 1 Etherium Sculptor | synergy | Reduces the cost of the deck's artifact spells.
- 1 Foundry Inspector | synergy | Additional artifact cost reduction for explosive development.
- 1 Mystic Forge | synergy | Lets artifacts be played from the top of the library.
- 1 Transmute Artifact | synergy | Converts a spare artifact into the most important artifact for the situation.
- 1 Whir of Invention | synergy | Instant-speed artifact tutoring that scales with the board.
- 1 Arcbound Crusher | threat | Grows into a substantial artifact-fueled attacker.
- 1 Chrome Steed | threat | Efficient artifact creature that attacks well with metalcraft.
- 1 Filigree Attendant | threat | Artifact creature that becomes a meaningful evasive body.
- 1 Frogmite | threat | Can be deployed cheaply to pressure opponents and fuel artifacts-matter effects.
- 1 Jhoira's Familiar | threat | Evasive artifact body that makes legendary spells easier to cast.
- 1 Lodestone Golem | threat | Pressures opponents while taxing nonartifact responses.
- 1 Lodestone Myr | threat | Converts a wide artifact board into a large attacker.
- 1 Mechan Assembler | threat | Artifact creature that provides resilient combat pressure.
- 1 Memory Guardian | threat | Artifact body that applies pressure while fitting the deck's synergies.
- 1 Panther Robot | threat | Efficient artifact creature for board presence.
- 1 Phyrexian Metamorph | threat | Copies the strongest artifact or creature on the battlefield.
- 1 Traxos, Scourge of Kroog | threat | Large, efficient artifact attacker that untaps through normal development.

</details>

### 4. dinosaur tribal, bracket 2

Format: Commander. Theme: dinosaurs. Pool: any_card. Shortlist: 269 names.

Commander: Gishath, Sun's Avatar.

Grade: bad, score 0.07, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band. Block findings: 0.

Cost: $264.91 to buy, $264.91 the whole deck.

**Summary:** A welcoming Naya dinosaur deck built around accelerating into large creatures, gaining value from tribal payoffs, and attacking with an imposing board. It uses a balanced suite of creature-based answers, protective tricks, and board resets to keep games interactive while rewarding straightforward combat decisions.

The quality model grades this deck below the precon baseline against the top lists of the format: more opening hands hold two to four lands than the norm, and that raises the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. The deck holds no fast mana, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core is a consistent ramp-into-fatties Gishath plan with dino tribal cost reducers and payoffs, but three symmetrical board wipes in a deck whose whole win condition is building an imposing creature board (and filler like Faithless Looting/Demand Answers) pull against that plan, especially for a new pilot.
- PLAN theme_fit=yes: It is a Naya Gishath dinosaur Commander deck with an accessible, combat-focused game plan and no combos or heavy stax, which matches the requested tribe, colors, format, and roughly bracket-2 power level.
- PLAN useful_as_built=yes: 37 lands with a solid Naya fixing base, ten ramp pieces, and multiple redundant threats and card-draw outlets mean the deck casts its spells and can close games through combat as written.
- PLAN summary_honest=yes: Every element named — ramp into large creatures, tribal payoffs, creature-based removal, protective tricks, and board resets — is actually present in the list, and it openly admits the wipes rather than hiding them.
- [INFO] `curve_summary`: average mana value 3.08 over 62 nonland cards

<details><summary>The deck list</summary>

- 1 Beast Whisperer | draw | Creature-based card advantage for the dinosaur-heavy game plan.
- 1 Garruk's Uprising | draw | Rewards large creatures with ongoing card draw and grants trample.
- 1 Demand Answers | draw | Low-cost card filtering when extra lands or artifacts are available.
- 1 Faithless Looting | draw | Efficient hand filtering with a later flashback option.
- 1 Kutzil, Malamet Exemplar | draw | Encourages combat and turns successful attacks into cards.
- 1 Ripjaw Raptor | draw | A dinosaur that converts damage into card advantage.
- 1 Runic Armasaur | draw | Provides steady incidental card draw against common activated abilities.
- 1 Skullclamp | draw | Turns smaller creatures and tokens into efficient card advantage.
- 1 Sylvan Library | draw | Improves draw quality throughout the game.
- 1 Return of the Wildspeaker | draw | Refills the hand from a large board or serves as a combat boost.
- 1 Cloudpiercer | draw | A dinosaur with flexible hand filtering.
- 1 Boros Charm | interaction | Protects the board from destruction or helps finish with combat damage.
- 1 Ephemerate | interaction | Protects an important creature while reusing its enter-the-battlefield ability.
- 1 Heroic Intervention | interaction | Protects the creature board from removal and sweepers.
- 1 Lightning Greaves | interaction | Provides immediate haste and durable protection for key creatures.
- 1 Snakeskin Veil | interaction | Low-cost protection that can also improve a creature in combat.
- 1 Swiftfoot Boots | interaction | Provides haste and repeatable protection for major threats.
- 1 Arcane Signet | ramp | Reliable multicolor mana acceleration.
- 1 Birds of Paradise | ramp | Early creature-based access to every deck color.
- 1 Cultivate | ramp | Finds lands while ensuring the next land drop.
- 1 Farseek | ramp | Efficiently finds a dual land with a needed basic land type.
- 1 Drover of the Mighty | ramp | Early mana creature that also benefits from controlling dinosaurs.
- 1 Llanowar Elves | ramp | One-mana acceleration toward larger creatures.
- 1 Nature's Lore | ramp | Efficiently finds a Forest dual land untapped.
- 1 Rampant Growth | ramp | Simple color fixing and land acceleration.
- 1 Sol Ring | ramp | Efficient acceleration for deploying the deck's large spells.
- 1 Thunderherd Migration | ramp | Dinosaur-themed land ramp and color fixing.
- 1 Itzquinth, Firstborn of Gishath | removal | Low-cost dinosaur that can fight opposing creatures.
- 1 Savage Stomp | removal | Efficient creature removal that is especially strong with dinosaurs.
- 1 Thrashing Brontodon | removal | Dinosaur body with flexible artifact or enchantment removal.
- 1 Triumphant Chomp | removal | Low-cost dinosaur-themed creature interaction.
- 1 Ravenous Sailback | removal | Removes an artifact or enchantment while adding a dinosaur body.
- 1 Raging Regisaur | removal | Combat-triggered creature removal attached to a dinosaur.
- 1 Needletooth Raptor | removal | Uses enrage to pick off opposing creatures.
- 1 Amped Raptor | synergy | Low-cost dinosaur that supports the tribe's creature curve.
- 1 Armored Kincaller | synergy | Early dinosaur presence that supports tribal development.
- 1 Belligerent Yearling | synergy | Efficient dinosaur that grows alongside larger creatures.
- 1 Commune with Dinosaurs | synergy | Finds either a needed land or an important dinosaur.
- 1 Deathgorge Scavenger | synergy | Provides useful graveyard control on a dinosaur body.
- 1 Dromosaur | synergy | A low-cost tribal creature that helps establish the board.
- 1 Drowsing Tyrannodon | synergy | Early dinosaur that becomes a meaningful attacker with a larger creature.
- 1 Huatli's Raptor | synergy | A cheap dinosaur that supports counter-based development.
- 1 Hunting Velociraptor | synergy | Rewards attacking with dinosaurs by helping deploy more creatures.
- 1 Kinjalli's Caller | synergy | Reduces the cost of the deck's large dinosaurs.
- 1 Marauding Raptor | synergy | Reduces dinosaur costs and rewards creatures entering the battlefield.
- 1 Nest Robber | synergy | Early tribal creature that helps begin attacking quickly.
- 1 Otepec Huntmaster | synergy | Reduces dinosaur costs and gives larger creatures haste.
- 1 Regisaur Alpha | threat | Creates additional dinosaur power and gives the tribe haste.
- 1 Quartzwood Crasher | threat | A trampling attacker that builds a growing army after combat damage.
- 1 Shifting Ceratops | threat | Resilient attacker with flexible combat abilities.
- 1 Goring Ceratops | threat | A large attacker that makes the combat step much stronger.
- 1 Majestic Heliopterus | threat | A substantial flying dinosaur threat for breaking stalled boards.
- 1 Pterodon Knight | threat | An efficient evasive attacker that benefits from the tribal plan.
- 1 Gigantosaurus | threat | A large, efficient dinosaur that pressures life totals.
- 1 Harnessed Snubhorn | threat | A sturdy dinosaur that contributes meaningful combat power.
- 1 Snapping Sailback | threat | A sizable flash threat that can disrupt combat.
- 1 Soaring Sandwing | threat | An evasive dinosaur that provides late-game pressure.
- 1 Stampeding Horncrest | threat | A hasty dinosaur attacker for immediate pressure.
- 1 Zilortha, Strength Incarnate | threat | Makes the deck's high-power creatures especially dangerous in combat.
- 1 Blasphemous Act | wipe | A powerful reset button when creature boards get crowded.
- 1 Chain Reaction | wipe | Creature sweep that scales with the number of creatures in play.
- 1 Vanquish the Horde | wipe | Often becomes an inexpensive answer to a crowded battlefield.
- 1 Battlefield Forge | land | Flexible red-white mana source.
- 1 Bountiful Promenade | land | Reliable green-white dual land for multiplayer games.
- 1 Brushland | land | Flexible green-white mana source.
- 1 Karplusan Forest | land | Flexible red-green mana source.
- 1 City of Brass | land | Provides access to every deck color.
- 1 Mana Confluence | land | Provides access to every deck color.
- 1 Command Tower | land | Reliable all-color commander mana.
- 1 Exotic Orchard | land | Usually supplies several needed colors in multiplayer.
- 1 Sacred Foundry | land | Fetchable red-white dual land.
- 1 Stomping Ground | land | Fetchable red-green dual land.
- 1 Temple Garden | land | Fetchable green-white dual land.
- 1 Cinder Glade | land | Fetchable red-green source that works well with basic lands.
- 1 Canopy Vista | land | Fetchable green-white source that works well with basic lands.
- 1 Jetmir's Garden | land | Three-color fixing with useful basic land types.
- 1 Path of Ancestry | land | Tribal color fixing with occasional card selection.
- 1 Secluded Courtyard | land | Tribal color fixing for dinosaur spells.
- 1 Unclaimed Territory | land | Tribal color fixing for dinosaur spells.
- 10 Forest | land | Primary basic source for early green ramp and dinosaur spells.
- 6 Mountain | land | Basic red sources for tribal creatures and interaction.
- 4 Plains | land | Basic white sources for protection and removal.

</details>

### 5. blink Commander, owned first

Format: Commander. Theme: blink. Pool: owned_first. Shortlist: 168 names.

Commander: Gilraen, Dúnedain Protector.

Grade: typical, score 0.41, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $215.35 the whole deck.

**Summary:** This deck develops mana, deploys value creatures, and uses blink synergies to keep generating value from its creature-based draw and removal pieces. It wins by building a protected creature board and turning that board into sustained combat pressure, with Angels, legendary creatures, and equipment providing resilient threats. It gives up the explosive speed of the format’s fastest lists in exchange for a steadier battlefield plan, permanent-based card advantage, and several broad reset buttons.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core is a coherent mono-white value pile (ETB creatures, removal, wipes, protection), but the blink engine is thin — only a handful of actual flicker effects (Flickerwisp, Angel of Condemnation, Slip On the Ring, Personify) — while a competing equipment/legendary-matters subtheme (Puresteel Paladin, Buster Sword, Sword of the Animist, Bastion Protector, Plaza of Heroes) pulls in a different direction.
- PLAN theme_fit=partly: It is a Commander deck at a reasonable bracket-3 power level with a legitimate white blink commander in Gilraen, but the requested blink theme is only lightly supported relative to the generic white goodstuff and equipment cards, and the off-color fixing lands suggest a build not tuned to the ask.
- PLAN useful_as_built=yes: 36 lands with plenty of ramp, a sensible curve, nine spot removal spells, three wipes, and multiple creature/equipment win routes make this immediately playable at a table.
- PLAN summary_honest=partly: It fairly reports the grindy value plan, the equipment/legend/Angel threats and the wipes, but leads with "blink synergies" as a driving engine when the list carries only a few flicker effects, and the boast that "the mana base serves the colors evenly" is empty in a mono-white deck that also runs pointless multicolor fixing (City of Brass, Spire of Industry, Exotic Orchard, Chromatic Lantern).
- [INFO] `curve_summary`: average mana value 2.94 over 63 nonland cards

<details><summary>The deck list</summary>

- 20 Plains | land | Provides dependable white mana for the deck’s creature-heavy plan.
- 1 Command Tower | land | Provides flexible commander-deck mana.
- 1 City of Brass | land | Provides broad color fixing.
- 1 Exotic Orchard | land | Provides flexible color fixing.
- 1 Marsh Flats | land | Provides land-based mana fixing.
- 1 Fabled Passage | land | Provides land-based mana fixing.
- 1 Evolving Wilds | land | Provides land-based mana fixing.
- 1 Terramorphic Expanse | land | Provides land-based mana fixing.
- 1 Ash Barrens | land | Helps smooth early land access.
- 1 Minas Tirith | land | Adds a utility land to the mana base.
- 1 Plaza of Heroes | land | Adds utility for the legendary commander plan.
- 1 Spire of Industry | land | Provides flexible mana alongside the artifact package.
- 1 Secluded Courtyard | land | Provides creature-focused mana fixing.
- 1 Path of Ancestry | land | Provides creature-focused mana fixing.
- 1 Windbrisk Heights | land | Adds a utility land to the mana base.
- 1 Great Hall of the Citadel | land | Adds utility land support for the commander plan.
- 1 The Grey Havens | land | Adds a utility land to the mana base.
- 1 Sol Ring | ramp | Efficiently accelerates the deck’s development.
- 1 Arcane Signet | ramp | Provides reliable mana acceleration.
- 1 Fellwar Stone | ramp | Provides inexpensive mana acceleration.
- 1 Wayfarer's Bauble | ramp | Provides early land-based acceleration.
- 1 Sword of the Animist | ramp | Combines creature pressure with mana development.
- 1 Relic of Legends | ramp | Provides mana acceleration for the creature shell.
- 1 Thought Vessel | ramp | Provides artifact-based mana acceleration.
- 1 Commander's Sphere | ramp | Provides flexible mana acceleration.
- 1 Chromatic Lantern | ramp | Provides mana acceleration and fixing.
- 1 White Lotus Tile | ramp | Adds another artifact-based mana source.
- 1 Inspiring Overseer | draw | Supplies creature-based card draw for the value plan.
- 1 Wall of Omens | draw | Provides an early defensive body and card draw.
- 1 Exemplar of Light | draw | Adds a creature-based draw option.
- 1 Energybending | draw | Provides straightforward card advantage.
- 1 Lembas | draw | Adds a compact artifact draw piece.
- 1 Idol of Oblivion | draw | Provides repeatable artifact-based card draw.
- 1 Tome of Legends | draw | Provides a durable source of card advantage.
- 1 Skullclamp | draw | Adds efficient equipment-based card draw.
- 1 Mask of Memory | draw | Turns creature combat into card selection and draw.
- 1 Buster Sword | draw | Adds an equipment-based draw source.
- 1 Mirror of Galadriel | draw | Provides another permanent-based draw engine.
- 1 Swords to Plowshares | removal | Provides efficient creature removal.
- 1 Generous Gift | removal | Answers a wide range of problematic permanents.
- 1 Get Lost | removal | Provides flexible permanent removal.
- 1 Destroy Evil | removal | Provides efficient targeted removal.
- 1 Crib Swap | removal | Provides creature removal that fits the creature-focused shell.
- 1 Banishing Light | removal | Provides permanent-based removal.
- 1 Dispatch | removal | Provides efficient targeted removal alongside artifacts.
- 1 Oblation | removal | Provides flexible instant-speed removal.
- 1 Stroke of Midnight | removal | Provides broad permanent removal.
- 1 Lightning Greaves | interaction | Protects important creatures and the commander.
- 1 Swiftfoot Boots | interaction | Provides repeatable creature protection.
- 1 Clever Concealment | interaction | Protects the developed board from opposing answers.
- 1 Reprieve | interaction | Supplies flexible stack interaction.
- 1 Gift of Immortality | interaction | Helps preserve an important creature on the battlefield.
- 1 Boromir, Warden of the Tower | interaction | Adds a protective creature to the board.
- 1 Bastion Protector | interaction | Helps safeguard the commander.
- 1 Frontline Medic | interaction | Provides creature-based protection for combat turns.
- 1 Austere Command | wipe | Provides a flexible reset when the board gets away.
- 1 Fumigate | wipe | Provides a full battlefield reset.
- 1 Vanquish the Horde | wipe | Provides an efficient creature-board reset.
- 1 Flickerwisp | synergy | A core creature for the blink-synergy package.
- 1 Angel of Sanctions | synergy | A creature-based removal piece for recurring blink value.
- 1 Angel of Serenity | synergy | A high-impact creature for the recurring blink plan.
- 1 Fiend Hunter | synergy | A creature-based removal piece that supports the blink shell.
- 1 Palace Jailer | synergy | A creature-based removal piece for the value plan.
- 1 Personify | synergy | A listed synergy piece for the blink shell.
- 1 Ennis, Debate Moderator | synergy | A listed synergy creature for the deck’s core plan.
- 1 Jocasta, Automaton Avenger | synergy | A listed synergy creature for the deck’s core plan.
- 1 Slip On the Ring | synergy | Supports the blink-focused strategy.
- 1 Sheltered by Ghosts | synergy | Supports the deck’s creature-centered synergy plan.
- 1 Giada, Font of Hope | threat | A creature threat that supports an Angel-focused combat board.
- 1 Bronze Guardian | threat | A durable artifact creature threat.
- 1 The Vision | threat | A creature threat that contributes to the board.
- 1 Zack Fair | threat | A legendary creature threat for the combat plan.
- 1 Westfold Rider | threat | A creature threat that adds battlefield pressure.
- 1 Kataki, War's Wage | threat | A creature threat with relevant board presence.
- 1 Puresteel Paladin | threat | A creature threat that works with the equipment package.
- 1 Troop of Ponies | threat | Adds another creature threat to develop the board.
- 1 The Walls of Ba Sing Se | threat | A substantial artifact creature threat.
- 1 Champions of Minas Tirith | threat | A creature threat that supports the combat plan.
- 1 Faramir, Field Commander | threat | A creature threat that helps build a board presence.
- 1 Angel of Condemnation | threat | A creature threat that fits the blink-focused creature base.

</details>

### 6. Modern tempo, tournament

Format: Modern. Theme: tempo. Pool: any_card. Shortlist: 167 names.

Grade: typical, score 0.44, model 20260903T210658Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $539.76 to buy, $539.76 the whole deck.

**Summary:** This blue-red tempo deck establishes pressure with inexpensive creatures, then uses efficient removal and permission to keep opponents from stabilizing. Card selection keeps the stream of spells and threats moving, while Temporal Mastery supplies a momentum-focused payoff that can turn a narrow lead into a win. The deck gives up broad answers and late-game inevitability in exchange for a lean, proactive game plan.

The quality model grades this deck typical against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade. The deck runs its spells as playsets, and that lowers the grade.

- PLAN plan_coherent=partly: The cheap threats, burn, and counters do form one tempo plan, but four copies of Temporal Mastery — a seven-mana sorcery with only six cantrips to enable the miracle — pull hard against the lean curve the rest of the list is built on.
- PLAN theme_fit=partly: It is a Modern blue-red tempo deck with the right colors and shape, but the person asked for Delver specifically and the list contains no Delver of Secrets, substituting a four-of Temporal Mastery package instead.
- PLAN useful_as_built=partly: The mana base is clean (18 blue and 14 red sources over 24 lands) and the threat/removal suite is functional, but four dead-in-hand seven-drops is a real drag on consistency in games where the miracle never lines up.
- PLAN summary_honest=partly: It accurately describes the pressure-plus-interaction shell, but it sells Temporal Mastery as a reliable 'momentum payoff' when four copies with minimal top-of-library setup will often be uncastable bricks, and it never notes that the requested Delver is absent.
- [INFO] `curve_summary`: average mana value 2.22 over 36 nonland cards

<details><summary>The deck list</summary>

- 10 Island | land | Basic blue source for the deck's early spells.
- 6 Mountain | land | Basic red source for the deck's early spells.
- 4 Steam Vents | land | Blue-red dual land for consistent two-color mana.
- 4 Spirebluff Canal | land | Blue-red land that supports the low curve.
- 4 Consider | draw | Low-cost card selection that keeps spells flowing.
- 2 Preordain | draw | Early card selection for finding the needed tempo piece.
- 4 Force of Negation | interaction | Free interaction for protecting the tempo plan.
- 2 Counterspell | interaction | Reliable permission against opposing spells.
- 4 Lightning Bolt | removal | Efficient red removal for clearing resistance.
- 2 Into the Flood Maw | removal | Cheap removal that helps preserve board pressure.
- 2 Untimely Malfunction | removal | Flexible removal slot for disrupting opposing permanents.
- 4 Temporal Mastery | synergy | Extra-turn synergy payoff for maintaining momentum.
- 4 Ragavan, Nimble Pilferer | threat | Early legendary creature threat for the tempo plan.
- 4 Ledger Shredder | threat | Cheap evasive creature threat that fits the spell-heavy shell.
- 4 Faerie Mastermind | threat | Flash creature threat that supports holding up interaction.

</details>

### 7. Modern burn, casual

Format: Modern. Theme: burn. Pool: any_card. Shortlist: 290 names.

Grade: bad, score 0.36, model 20260903T210658Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $47.58 to buy, $47.58 the whole deck.

**Summary:** This mono-red burn deck applies steady pressure with aggressive threats, then uses removal and spell-focused synergy to finish the opponent. Card draw helps it keep casting spells after the first rush, while a small ramp package supports its red threats. It gives up broad answers and sideboard flexibility in exchange for a focused, direct game plan.

The quality model grades this deck below the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade.

- PLAN plan_coherent=partly: The core is red aggro plus burn, but the pieces pull apart: Ancestral Anger x4 is a filler cantrip pump spell that doesn't support the creature suite, Guttersnipe/Risk Factor want a dense spell count the deck lacks (only 12 non-creature spells), and Hazoret's hellbent condition fights the four-and-five-drop threats like Ashcloud Phoenix and Sunspine Lynx.
- PLAN theme_fit=partly: It is mono-red and Modern-legal and aggressive, but it is a midrange creature deck with a burn splash rather than a burn deck — only 6 burn/removal spells and no Lava Spike/Boros Charm/Skewer-style reach at all.
- PLAN useful_as_built=partly: 24 Mountains and a legal 60 make it playable and it can win with creature beats, but the curve is far too top-heavy for a 'burn' plan and it has almost no direct reach to close out games, so it plays much worse than the summary implies.
- PLAN summary_honest=partly: It correctly flags removal, draw, and a ramp bit, but calling Chandra a 'ramp package' and Ancestral Anger 'card draw' oversells filler, and 'spell-focused synergy' (Guttersnipe) is not supported by the actual spell density.
- [INFO] `curve_summary`: average mana value 2.72 over 36 nonland cards

<details><summary>The deck list</summary>

- 24 Mountain | land | Provides a consistent mono-red mana base.
- 2 Chandra, Dressed to Kill | ramp | Fills the deck’s ramp slot while staying in the mono-red plan.
- 4 Lightning Bolt | removal | Efficient burn-oriented removal.
- 2 Lava Dart | removal | Adds additional low-cost removal for the burn plan.
- 4 Ancestral Anger | draw | Provides card draw to keep spells flowing.
- 2 Risk Factor | draw | Provides additional card draw for longer games.
- 4 Eidolon of the Great Revel | synergy | A synergy piece that supports the deck’s aggressive burn focus.
- 4 Guttersnipe | synergy | Rewards the deck for casting its many burn spells.
- 4 Ashcloud Phoenix | threat | A resilient red threat for applying pressure.
- 4 Hazoret the Fervent | threat | A powerful threat that fits an aggressive red strategy.
- 2 Torbran, Thane of Red Fell | threat | A red threat that reinforces the deck’s damage-focused plan.
- 4 Sunspine Lynx | threat | A finishing threat for games that go beyond the early burn turns.

</details>

### 8. Modern lifegain, FNM

Format: Modern. Theme: lifegain. Pool: any_card. Shortlist: 301 names.

Grade: bad, score 0.07, model 20260903T210658Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $235.17 to buy, $235.17 the whole deck.

**Summary:** This white-black lifegain deck takes a creature-forward approach, pairing Soul Warden and Ajani's Pridemate with a threat base of Clerics, Vampires, and Knights. It aims to win through sustained battlefield pressure, using Solitude and Deep-Cavern Bat to clear the way while Lembas and Inspiring Overseer keep cards flowing. The deck gives up some speed for a heavier midrange creature package and a straightforward, synergy-focused game plan.

The quality model grades this deck below the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The deck runs its spells as playsets, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade.

- PLAN plan_coherent=partly: The core is a consistent lifegain-payoff build (Soul Warden, Attended Healer, Pridemate, Bloodbond Vampire, Twinblade Paladin, Enduring Tenacity all feed each other), but the singleton Arcane Signet and Fellwar Stone are ramp with no home in a 60-card two-color aggro-midrange list and pull against the curve.
- PLAN theme_fit=partly: It is clearly a white-black lifegain 60-card deck as asked, but Fellwar Stone has no Modern-legal printing and the free-spell Solitude playset sits well above the FNM budget the request implies.
- PLAN useful_as_built=partly: Twenty-four lands with a strong dual-heavy base support the low curve and the deck can win through Pridemate/Bloodbond growth and Tenacity drain, but the two random mana rocks are dead draws and one is not legal for sanctioned Modern play.
- PLAN summary_honest=partly: placeholder
- [INFO] `curve_summary`: average mana value 3.06 over 36 nonland cards

<details><summary>The deck list</summary>

- 4 Godless Shrine | land | Primary white-black land for the deck.
- 4 Caves of Koilos | land | Provides consistent white-black mana.
- 4 Shattered Sanctum | land | Supports the two-color mana base.
- 4 Plains | land | Basic white source for the lifegain core.
- 4 Swamp | land | Basic black source for the deck's threats and removal.
- 2 Fetid Heath | land | Additional white-black color fixing.
- 2 Isolated Chapel | land | Completes the two-color land package.
- 1 Arcane Signet | ramp | Early mana acceleration for the deck's larger creatures.
- 1 Fellwar Stone | ramp | Additional mana acceleration in a two-color deck.
- 4 Lembas | draw | Low-investment card advantage that fits the lifegain plan.
- 2 Inspiring Overseer | draw | A creature-based source of card advantage.
- 4 Solitude | removal | Primary removal suite for opposing threats.
- 2 Deep-Cavern Bat | removal | Creature-based removal that contributes to board presence.
- 4 Soul Warden | synergy | Core lifegain synergy piece for the early game.
- 4 Ajani's Pridemate | synergy | Central payoff for the deck's lifegain synergies.
- 4 Attended Healer | threat | Creature threat that supports the deck's lifegain direction.
- 4 Enduring Tenacity | threat | Midgame threat for the white-black lifegain shell.
- 4 Bloodbond Vampire | threat | A creature threat that fits the lifegain theme.
- 2 Twinblade Paladin | threat | Additional creature pressure for closing games.

</details>

### 9. Standard midrange, FNM

Format: Standard. Theme: midrange. Pool: any_card. Shortlist: 164 names.

Grade: bad, score 0.02, model 20260903T210658Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $170.82 to buy, $170.82 the whole deck.

**Summary:** This black-green midrange deck establishes its mana, answers opposing permanents, and keeps cards flowing while deploying a steady stream of creatures. It wins by forcing opponents to handle resilient threats from Goldvein Hydra through Vaultborn Tyrant, with protection spells helping preserve key creatures. The deck gives up some matchup-specific flexibility for a focused main-deck plan built around creatures, removal, and sustained card advantage.

The quality model grades this deck below the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. Many of the cards appear in no top list, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: The list holds together as one plan: a bit of ramp and removal into resilient four-to-six drop threats, with Snakeskin Veil and Undying Malice protecting them, which is exactly the midrange curve the summary describes.
- PLAN theme_fit=no: The colors and midrange shape match, but the request named Standard and a large share of the deck is not Standard-legal (Maelstrom Pulse, Llanowar Elves, Phyrexian Arena, Overgrown Tomb, Blooming Marsh, Snakeskin Veil, Undying Malice), so it is not a deck that can be sleeved up at FNM Standard.
- PLAN useful_as_built=partly: As a 60-card pile it has a workable 24-land base, real removal, and enough finishers to close, but the top end is heavy on four-to-six drops with only two Llanowar Elves to bridge, and it would be rejected at the event the player asked about.
- PLAN summary_honest=partly: placeholder
- [INFO] `curve_summary`: average mana value 3.00 over 36 nonland cards

<details><summary>The deck list</summary>

- 8 Forest | land | Basic green source for the black-green mana base.
- 8 Swamp | land | Basic black source for the black-green mana base.
- 4 Overgrown Tomb | land | Dual black-green land for consistent access to both colors.
- 4 Blooming Marsh | land | Additional black-green land to support the early game.
- 2 Llanowar Elves | ramp | Early mana creature that helps advance into the deck's larger cards.
- 4 Bitter Triumph | removal | Efficient black removal that keeps opposing threats in check.
- 2 Maelstrom Pulse | removal | Flexible black-green removal for problematic permanents.
- 4 Darkstar Augur | draw | Creature-based card advantage that also contributes to board presence.
- 2 Phyrexian Arena | draw | Dedicated long-game card advantage for midrange mirrors.
- 4 Snakeskin Veil | synergy | Protection that helps the creature suite stay on the battlefield.
- 4 Undying Malice | synergy | Creature protection that supports trading and rebuilding through removal.
- 4 Goldvein Hydra | threat | A scalable green creature that serves as a primary battlefield threat.
- 4 Vein Ripper | threat | A powerful black creature that pressures opponents in the late game.
- 4 Rottenmouth Viper | threat | A black-green creature that adds to the deck's threat density.
- 2 Vaultborn Tyrant | threat | Large creature finisher for games that go long.

</details>

### 10. Standard aggro, tournament

Format: Standard. Theme: aggro. Pool: any_card. Shortlist: 294 names.

Grade: baseline, score 0.41, model 20260903T210658Z.

Cards: 60 main, 0 sideboard. Repair turn: yes, for profile_off_band; then profile_off_band. Block findings: 0.

Cost: $100.10 to buy, $100.10 the whole deck.

**Summary:** A fast red-white aggro deck built to establish early combat pressure, remove opposing blockers efficiently, and preserve momentum with flexible disruption and card access. Its mana base emphasizes dependable red access while retaining strong white support for its tactical elements.

The quality model grades this deck at the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. Many of the cards appear in no top list, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The bulk of the list is a coherent RW aggro plan of cheap threats plus burn/removal, but two copies of the six-mana Krenko's Buzzcrusher and the Rakdos-costed Redcap Gutter-Dweller (four copies, a top-end threat needing black) pull against the fast-clock plan.
- PLAN theme_fit=partly: It is red-white and aggressive as asked, but it is not a Standard-legal deck: Boros Charm, Inspiring Vantage, and Sacred Foundry are outside the Standard card pool, so it could not be registered at the tournament named.
- PLAN useful_as_built=no: Beyond the illegal cards, four copies of a threat requiring a black pip cannot be cast off a Plains/Mountain/Sacred Foundry/Sunbillow Verge base, so a meaningful chunk of the deck is dead in hand.
- PLAN summary_honest=partly: The aggro/removal description broadly matches the list, but the claim of a mana base with "dependable red access" and "strong white support" quietly papers over a four-of that needs black mana and several lands/spells that are not legal in the requested format.
- [INFO] `curve_summary`: average mana value 2.72 over 36 nonland cards

<details><summary>The deck list</summary>

- 4 Fugitive Codebreaker | draw | Early attacker that keeps pressure flowing while providing card access.
- 2 Reckless Lackey | draw | Low-cost aggressive creature that converts into additional cards.
- 4 Boros Charm | interaction | Flexible protection and reach for an attacking board.
- 2 Sheltered by Ghosts | interaction | Protects a key attacker while disrupting opposing development.
- 4 Harsh Annotation | removal | Efficient red removal to clear blockers.
- 2 Case of the Gateway Express | removal | Additional white removal that supports the aggressive plan.
- 2 Krenko's Buzzcrusher | removal | Creature-based removal that maintains battlefield pressure.
- 4 Slickshot Show-Off | synergy | Cheap evasive threat that rewards a spell-heavy aggressive turn.
- 4 Redcap Gutter-Dweller | threat | Reliable red attacker for the early combat plan.
- 4 Frilled Sparkshooter | threat | Efficient creature threat that pressures opponents early.
- 4 Dragonback Lancer | threat | Red-white attacker that gives the deck another assertive threat tier.
- 4 Sacred Foundry | land | Untapped dual source for both primary colors.
- 4 Inspiring Vantage | land | Fast dual land for early red-white development.
- 4 Sunbillow Verge | land | Dual-color land supporting both halves of the deck.
- 4 Plains | land | Reliable white source for white interaction and removal.
- 8 Mountain | land | Heavy red basic allocation supports the deck's red-intensive aggression and removal.

</details>

### 11. Commander with a locked card

Format: Commander. Theme: sacrifice. Pool: any_card. Shortlist: 296 names.

Commander: Karlov of the Ghost Council.

Grade: bad, score 0.30, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $322.90 to buy, $322.90 the whole deck.

**Summary:** This Karlov sacrifice deck builds a creature-based engine of sacrifice outlets, payoffs, and resource conversion, then turns that board presence into steady pressure. It wins by assembling durable creature threats while using sacrifice-enabled removal and wipes to keep opposing boards contained. The deck gives up some mana-base utility and broad card selection in favor of a stable two-color foundation and a dense, focused sacrifice package.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. Many of the cards appear in no top list, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every nonland card is a sacrifice outlet, a death/sacrifice payoff, or a token/body source, and the drain payoffs (Bastion of Remembrance, Vindictive Vampire, Elas il-Kor, Vito's Inquisitor) feed Karlov's lifegain trigger, so the pieces pull one direction.
- PLAN theme_fit=yes: It is a 99-card Commander deck in Orzhov led by Karlov of the Ghost Council built squarely around sacrifice, exactly what was asked.
- PLAN useful_as_built=yes: 36 basics in a two-color deck with Sol Ring, altars and cheap rituals supports a low curve, and it has multiple real win routes (Blasting Station/Altar of Dementia loops, aristocrat drain, Karlov beats), so it plays out of the box despite the plain mana base.
- PLAN summary_honest=yes: placeholder
- [INFO] `curve_summary`: average mana value 2.87 over 63 nonland cards

<details><summary>The deck list</summary>

- 18 Plains | land | Basic white mana for the deck.
- 18 Swamp | land | Basic black mana for the deck.
- 1 Corrupted Conviction | draw | Efficient draw that supports the sacrifice plan.
- 1 Disciple of Bolas | draw | Draw that rewards the deck's creature-focused plan.
- 1 Relic Vial | draw | Repeatable draw support for the creature base.
- 1 Shadowheart, Dark Justiciar | draw | A draw source that works with expendable creatures.
- 1 Skyclave Shadowcat | draw | Draw that fits the sacrifice-focused creature suite.
- 1 Smothering Abomination | draw | A substantial draw engine for the sacrifice strategy.
- 1 Vampiric Rites | draw | Low-cost draw support for sacrificed creatures.
- 1 Village Rites | draw | Fast draw that converts a creature into resources.
- 1 Baron Bertram Graywater | draw | A creature-based draw option for the deck's board plan.
- 1 Thraxodemon | draw | Creature-based draw that belongs in the sacrifice shell.
- 1 Bushmeat Poacher | draw | A draw source that uses the deck's creatures as resources.
- 1 Sol Ring | ramp | Fast early mana that the deck is required to keep.
- 1 Ashnod's Altar | ramp | A sacrifice outlet that turns creatures into mana.
- 1 Phyrexian Altar | ramp | Sacrifice-based mana production for the core plan.
- 1 Pitiless Plunderer | ramp | Mana support that rewards the sacrifice strategy.
- 1 Pawn of Ulamog | ramp | Creature-based mana support for sacrifice turns.
- 1 Priest of Forgotten Gods | ramp | A sacrifice-focused source of mana acceleration.
- 1 Skullport Merchant | ramp | A creature-based ramp piece for the resource plan.
- 1 Warren Soultrader | ramp | Mana acceleration that belongs in a sacrifice deck.
- 1 Culling the Weak | ramp | Burst mana that converts a creature into development.
- 1 Deadly Dispute | ramp | Resource conversion that helps advance the mana plan.
- 1 Bone Shards | removal | Low-cost removal that fits a sacrifice-oriented deck.
- 1 Bone Splinters | removal | Cheap removal using the deck's expendable creatures.
- 1 Attrition | removal | Repeatable removal tied to sacrificing creatures.
- 1 Ayli, Eternal Pilgrim | removal | A sacrifice-friendly creature that provides removal.
- 1 Blasting Station | removal | A sacrifice outlet that also supplies removal.
- 1 Chittering Witch | removal | Creature-based removal that supports the board plan.
- 1 Dictate of Erebos | removal | A payoff that turns sacrifices into opposing losses.
- 1 Eaten Alive | removal | Efficient removal that can use a creature as its cost.
- 1 Yawgmoth, Thran Physician | removal | A sacrifice-centric removal engine.
- 1 Altar of Dementia | synergy | A sacrifice outlet that advances the deck's synergy plan.
- 1 Bastion of Remembrance | synergy | A sacrifice payoff for the deck's creature plan.
- 1 Bloodflow Connoisseur | synergy | A dedicated creature sacrifice outlet.
- 1 Carrion Feeder | synergy | A low-cost sacrifice outlet for the engine.
- 1 Chthonian Nightmare | synergy | Recursion support for sacrificed creatures.
- 1 Elas il-Kor, Sadistic Pilgrim | synergy | A creature-based payoff for sacrificing creatures.
- 1 Fleshtaker | synergy | A sacrifice payoff that strengthens the core plan.
- 1 Hidden Stockpile | synergy | Ongoing support for the sacrifice strategy.
- 1 Viscera Seer | synergy | An inexpensive and reliable sacrifice outlet.
- 1 Woe Strider | synergy | A sacrifice outlet that fits the creature engine.
- 1 Basri's Lieutenant | threat | A creature threat that supports the deck's board presence.
- 1 Felisa, Fang of Silverquill | threat | A threat that fits alongside Karlov and sacrifice synergies.
- 1 Flesh-Eater Imp | threat | A compact creature threat for closing games.
- 1 Liliana, Heretical Healer // Liliana, Defiant Necromancer | threat | A resilient threat for the creature-heavy strategy.
- 1 Mondrak, Glory Dominus | threat | A powerful board-building threat for the deck.
- 1 Ratadrabik of Urborg | threat | A creature threat that reinforces the legendary core.
- 1 Sidisi, Undead Vizier | threat | A high-impact creature threat for sacrifice turns.
- 1 Vito's Inquisitor | threat | A creature threat that maintains offensive pressure.
- 1 Vindictive Vampire | threat | A threat that supports the sacrifice-focused creature plan.
- 1 Blood Host | threat | A creature threat that benefits from the deck's sacrifices.
- 1 Demon of Catastrophes | threat | A sizeable threat that can use an expendable creature.
- 1 Thallid Omnivore | threat | A creature threat that fits the sacrifice shell.
- 1 Cartel Aristocrat | interaction | A sacrifice outlet that provides protective interaction.
- 1 Dark Privilege | interaction | Protection interaction for an important creature.
- 1 Duty Beyond Death | interaction | Low-cost interaction that protects the creature plan.
- 1 Fanatical Devotion | interaction | A sacrifice-based way to protect key creatures.
- 1 Gift of Doom | interaction | Protective interaction for a central permanent.
- 1 Promise of Tomorrow | interaction | Interaction that safeguards the creature strategy.
- 1 Spirit Bonds | interaction | Board protection that supports the creature plan.
- 1 Sunstone | interaction | Defensive interaction that buys time to develop.
- 1 Toxic Deluge | wipe | A flexible board wipe for difficult boards.
- 1 The Meathook Massacre | wipe | A board wipe that complements the sacrifice strategy.
- 1 Minister of Pain | wipe | A creature-based board wipe option.

</details>

### 12. Commander on a budget

Format: Commander. Theme: tokens. Pool: any_card. Shortlist: 296 names.

Commander: Adeline, Resplendent Cathar.

Grade: bad, score 0.02, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: yes, for deck_size, profile_off_band, profile_off_band, profile_off_band, profile_off_band. Block findings: 0.

Cost: $68.48 to buy, $68.48 the whole deck.

**Summary:** A mono-white go-wide strategy that develops a dense creature board, amplifies tokens with anthem effects and permanent growth, and turns repeated attacks into overwhelming combat pressure. The deck pairs resilient token engines with broad removal, protective effects, and board resets that favor rebuilding a wide battlefield.

The quality model grades this deck below the precon baseline against the top lists of the format: more opening hands hold two to four lands than the norm, and that raises the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every slot points at one go-wide token plan — token makers, anthems, convoke/token-friendly wipes, and payoffs like Divine Visitation and Hero of Bladehold — with only the grab-bag of treasure trinkets (Coin of Mastery, Prying Blade, Collector's Vault) reading as budget filler rather than a competing plan.
- PLAN theme_fit=yes: It is a mono-white Adeline token Commander deck with clearly budget-tier card choices and no combos or fast mana, matching the under-$100, bracket 2 request.
- PLAN useful_as_built=yes: 99 cards with 37 lands in a single color, a sensible curve, card draw, seven removal spells and multiple token-friendly wipes, so it can be shuffled up and played as printed.
- PLAN summary_honest=yes: placeholder
- [INFO] `curve_summary`: average mana value 3.24 over 62 nonland cards

<details><summary>The deck list</summary>

- 29 Plains | land | Reliable untapped white sources.
- 1 Ancient Den | land | White source with artifact synergy.
- 1 Castle Ardenvale | land | White source that can produce late-game tokens.
- 1 Command Tower | land | Reliable commander-color fixing.
- 1 Eiganjo, Seat of the Empire | land | White source with flexible creature interaction.
- 1 Mirrex | land | Token-producing utility land.
- 1 Plaza of Heroes | land | Protects the commander while producing white mana.
- 1 Secluded Courtyard | land | Names Human for the deck's substantial Human base.
- 1 Unclaimed Territory | land | Tribal white fixing for the creature-heavy plan.
- 1 Charisma Bobblehead | ramp | Artifact mana acceleration.
- 1 Coin of Mastery | ramp | Flexible artifact mana acceleration.
- 1 Collector's Vault | ramp | Treasure production with hand filtering.
- 1 Currency Converter | ramp | Low-cost Treasure support.
- 1 Goldvein Pick | ramp | Turns attacks into Treasure.
- 1 Karn, Living Legacy | ramp | Produces mana artifacts while providing value.
- 1 Keeper of the Accord | ramp | Catches up on lands and builds a board.
- 1 Monologue Tax | ramp | Converts opponents' spell-heavy turns into Treasure.
- 1 Noble's Purse | ramp | Early Treasure access on a cheap artifact.
- 1 Prying Blade | ramp | Attack-triggered Treasure production.
- 1 Angelic Sell-Sword | draw | Creature-based card advantage.
- 1 Bygone Bishop | draw | Creates Clues from inexpensive creatures.
- 1 Dawn of Hope | draw | Repeatable token and card-draw outlet.
- 1 Faramir, Field Commander | draw | Rewards attacking with cards and additional board presence.
- 1 Glimmer Seeker | draw | Efficient creature-based card access.
- 1 Idol of Oblivion | draw | Draws efficiently after token creation.
- 1 Sarah Jane Smith | draw | Provides steady card advantage.
- 1 Search the Premises | draw | Turns attacks against you into Clues.
- 1 Staff of the Storyteller | draw | Supplies a token and repeatable card draw.
- 1 Sunpearl Kirin | draw | Creature-based card selection and advantage.
- 1 Wedding Announcement // Wedding Festivity | draw | Develops the board before improving the whole token army.
- 1 Basri Ket | interaction | Builds tokens and protects the attacking plan.
- 1 Blessed Sanctuary | interaction | Prevents damage while creating tokens.
- 1 Lena, Selfless Champion | interaction | Protects a wide creature board from removal.
- 1 Rootborn Defenses | interaction | Protects the army and populates.
- 1 Spirit Bonds | interaction | Makes tokens and can shield key creatures.
- 1 Squad Commander | interaction | Provides a wide-board combat payoff.
- 1 Aerial Assault | removal | Efficient flying-creature removal with life gain.
- 1 Banishing Slash | removal | Flexible sorcery-speed permanent removal.
- 1 Generous Gift | removal | Answers any problematic permanent.
- 1 Kellan's Lightblades | removal | Low-cost instant-speed creature interaction.
- 1 Righteous Confluence | removal | Flexible answer for creatures, artifacts, and enchantments.
- 1 Skyclave Apparition | removal | Creature-based exile removal.
- 1 Stroke of Midnight | removal | Instant-speed answer to any nonland permanent.
- 1 Hour of Reckoning | wipe | Token-friendly board reset.
- 1 Martial Coup | wipe | Sweeps the battlefield and rebuilds with Soldiers.
- 1 Phyrexian Rebirth | wipe | Board reset that leaves behind a sizable threat.
- 1 Animation Module | synergy | Turns counters and mana into additional creatures.
- 1 Anointer Priest | synergy | Rewards the steady stream of creature tokens.
- 1 Clarion Spirit | synergy | Creates Spirit tokens while developing the turn.
- 1 Divine Visitation | synergy | Upgrades small tokens into powerful evasive threats.
- 1 Felidar Retreat | synergy | Converts land drops into creatures or team-wide counters.
- 1 Horn of Gondor | synergy | Scales Soldier production with the existing board.
- 1 Intangible Virtue | synergy | Improves tokens with vigilance and extra power.
- 1 Oketra's Monument | synergy | Reduces creature costs while creating Warriors.
- 1 Retrofitter Foundry | synergy | Converts spare mana into increasingly relevant tokens.
- 1 Rosie Cotton of South Lane | synergy | Turns incoming tokens into permanent team growth.
- 1 Siege Veteran | synergy | Strengthens Soldiers and replaces them when they fall.
- 1 Skrelv's Hive | synergy | Provides a persistent stream of artifact creature tokens.
- 1 Worthy Knight | synergy | Rewards Knights with extra creature tokens.
- 1 Ajani's Chosen | threat | Creates Cats from the deck's enchantment package.
- 1 Archon of Sun's Grace | threat | Produces evasive Pegasi from enchantments.
- 1 Attended Healer | threat | Builds a Cat army alongside incidental life gain.
- 1 Basri's Lieutenant | threat | Makes a resilient board of Knight threats.
- 1 Cemetery Protector | threat | Pressures opponents while improving creature production.
- 1 Emeria Angel | threat | Makes evasive tokens from land drops.
- 1 Hero of Bladehold | threat | Creates attackers and dramatically improves combat damage.
- 1 God-Eternal Oketra | threat | Turns creature spells into a stream of vigilant Zombies.
- 1 Mite Overseer | threat | Builds a growing army of artifact creature tokens.
- 1 Phantom General | threat | Provides a substantial anthem for token combat.
- 1 Requiem Angel | threat | Replaces fallen non-Spirit creatures with fliers.
- 1 Warren Warleader | threat | Brings an immediate Rabbit force and scales with attacks.

</details>

### 13. owned first, and the commander is not owned

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 229 names.

Commander: Karlov of the Ghost Council.

Grade: bad, score 0.08, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $138.45 the whole deck.

**Summary:** This Karlov deck builds around lifegain synergies while using efficient mana, steady card flow, and protective interaction to keep its key creatures on the table. It wins by turning that supported board into sustained pressure through its creature threats, while removal and sweepers prevent opposing boards from taking over. The deck gives up specialized utility lands and explosive mana for a consistent basic-land mana base and a straightforward, resilient lifegain game.

The quality model grades this deck below the precon baseline against the top lists of the format: more opening hands hold two to four lands than the norm, and that raises the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. The deck holds no fast mana, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: There is a broad lifegain-creature core with plenty of removal and protection, but the deck barely uses Karlov's actual engine (few incidental lifegain triggers per turn, no lifelink-matters payoffs beyond a handful) and it also runs three board wipes plus Austere Command/Dusk//Dawn that fight its own go-wide creature threats.
- PLAN theme_fit=yes: It is a mono-W/B Karlov of the Ghost Council Commander deck built on lifegain with no fast mana or high-power staples, matching the requested bracket-3, library-first ask.
- PLAN useful_as_built=partly: 36 basics plus six rocks gives clean mana and the curve is castable, but the win path is vague creature beatdown that the deck's own three wraths undercut, and Karlov gets few counters from the small number of repeatable lifegain sources.
- PLAN summary_honest=no: placeholder
- [WARN] `not_owned`: Karlov of the Ghost Council: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.06 over 63 nonland cards

<details><summary>The deck list</summary>

- 18 Plains | land | Basic white mana source for the deck.
- 18 Swamp | land | Basic black mana source for the deck.
- 1 Sol Ring | ramp | Efficiently fills a ramp slot.
- 1 Arcane Signet | ramp | Provides reliable color fixing and ramp.
- 1 Fellwar Stone | ramp | Adds another early ramp piece.
- 1 Wayfarer's Bauble | ramp | Provides a ramp slot without adding color demands.
- 1 Thought Vessel | ramp | Provides artifact-based ramp.
- 1 Commander's Sphere | ramp | Provides flexible mana as a ramp piece.
- 1 Sword of the Animist | ramp | Supports the deck while filling a ramp slot.
- 1 Oath of the Grey Host | ramp | Adds ramp that fits the deck's colors.
- 1 Lotho, Corrupt Shirriff | ramp | Provides a creature-based ramp slot.
- 1 Giada, Font of Hope | ramp | Provides ramp alongside the creature plan.
- 1 Buster Sword | draw | Fills a draw slot while remaining useful to the creature plan.
- 1 Call of the Ring | draw | Provides a dedicated draw piece.
- 1 Cirith Ungol Patrol | draw | Adds draw from a creature slot.
- 1 Exemplar of Light | draw | Provides draw within the deck's creature plan.
- 1 Gollum, Riddle Master | draw | Adds a creature-based draw option.
- 1 Grave Venerations | draw | Provides a dedicated draw piece.
- 1 Hero in Training | draw | Adds draw while contributing to the board.
- 1 Inspiring Overseer | draw | Provides draw from a creature.
- 1 Lembas | draw | Provides a compact draw piece.
- 1 Mask of Memory | draw | Supports the creature plan while filling a draw slot.
- 1 Night's Whisper | draw | Provides efficient dedicated draw.
- 1 Lightning Greaves | interaction | Protects an important creature.
- 1 Swiftfoot Boots | interaction | Provides repeatable creature protection.
- 1 Darksteel Plate | interaction | Adds a durable protection piece.
- 1 Champion's Helm | interaction | Protects a key legendary creature.
- 1 Clever Concealment | interaction | Provides protection against opposing interaction.
- 1 Unbreakable Formation | interaction | Helps preserve the creature board.
- 1 Take Up the Shield | interaction | Provides a compact protection effect.
- 1 Reprieve | interaction | Adds flexible stack interaction.
- 1 Banishing Light | removal | Provides a removal slot for problematic permanents.
- 1 Bitter Triumph | removal | Provides efficient targeted removal.
- 1 Crib Swap | removal | Adds versatile creature removal.
- 1 Dispatch | removal | Provides a low-cost removal option.
- 1 Fatal Push | removal | Adds efficient targeted removal.
- 1 Generous Gift | removal | Provides broad permanent removal.
- 1 Get Lost | removal | Adds another flexible removal option.
- 1 Infernal Grasp | removal | Provides reliable targeted removal.
- 1 Swords to Plowshares | removal | Provides efficient creature removal.
- 1 Aerith Gainsborough | synergy | Supports the lifegain synergy plan.
- 1 Angel of Vitality | synergy | Supports the deck's lifegain focus.
- 1 Compassionate Healer | synergy | Adds a lifegain-focused synergy piece.
- 1 Elixir | synergy | Provides a compact synergy piece for the plan.
- 1 Kor Firewalker | synergy | Supports the lifegain strategy from the creature suite.
- 1 Light of Promise | synergy | Builds on the deck's lifegain theme.
- 1 Night Nurse, Healer of Heroes | synergy | Adds another lifegain-oriented synergy card.
- 1 Prideful Feastling | synergy | Supports the central lifegain plan.
- 1 Rosie Cotton of South Lane | synergy | Adds a creature-based synergy piece.
- 1 Second Breakfast | synergy | Provides an additional lifegain synergy card.
- 1 Angel of Invention | threat | Provides a substantial creature threat.
- 1 Bill the Pony | threat | Adds a creature threat to pressure opponents.
- 1 Dawnhand Eulogist | threat | Contributes another board-based threat.
- 1 Foggy Swamp Hunters | threat | Adds a creature threat to the deck.
- 1 Gollum, Silent Slinker // Meager Meal | threat | Provides a threat with additional card flexibility.
- 1 Lyra Dawnbringer | threat | Adds a powerful legendary creature threat.
- 1 Minwu, White Mage | threat | Provides another creature threat.
- 1 Reaping Willow | threat | Adds a resilient-looking board threat.
- 1 Rooftop Percher | threat | Contributes to the deck's threat density.
- 1 Shattered Angel | threat | Adds an Angel threat aligned with the creature plan.
- 1 Sneering Shadewriter | threat | Provides another black creature threat.
- 1 Victory's Herald | threat | Serves as a top-end creature threat.
- 1 Austere Command | wipe | Provides a flexible board wipe.
- 1 Fumigate | wipe | Adds a dedicated reset for crowded boards.
- 1 Dusk // Dawn | wipe | Provides a board wipe with added flexibility.

</details>

### 14. the user delegates the commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 244 names.

Commander: Atarka, World Render.

Grade: bad, score 0.02, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $603.35 to buy, $603.35 the whole deck.

**Summary:** This deck ramps through its early turns, develops a Dragon-focused board, and keeps cards flowing with creature-based draw engines and efficient refills. It wins by applying sustained combat pressure with large flying threats under Atarka, World Render, while Dragon-based damage and removal keep blockers and opposing threats manageable. It gives up some speed and precision for a creature-heavy plan that needs mana development and a stable board to deliver its strongest turns.

The quality model grades this deck below the precon baseline against the top lists of the format: the mana base serves one color far better than another, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every card supports one plan — ramp and cost reducers into big Dragons, Dragon-tribal payoffs (Dragon Tempest, Scourge of Valkas, Crucible of Fire, Lathliss, Utvara) and creature-based draw to fuel attacks with Atarka's double strike — with only a couple of odd stragglers like Tibalt's Trickery.
- PLAN theme_fit=yes: It is a Commander-legal, singleton Dragon tribal deck with a builder-chosen Dragon commander and a mid-power, non-combo build appropriate to bracket 3.
- PLAN useful_as_built=yes: 36 lands plus ten ramp pieces and Dragon cost reducers support the heavy top end, and there are plenty of ways to close a game, though the green sources are thin for the handful of green spells and the curve is demanding.
- PLAN summary_honest=yes: placeholder
- [INFO] `curve_summary`: average mana value 3.21 over 63 nonland cards

<details><summary>The deck list</summary>

- 18 Mountain | land | Provides the red mana base for the deck’s dragon-heavy spells.
- 8 Forest | land | Provides green mana for ramp, draw, and protection spells.
- 1 Command Tower | land | Provides flexible commander-color mana.
- 1 Taiga | land | Provides both deck colors without slowing development.
- 1 Stomping Ground | land | Provides both deck colors for early casting.
- 1 Karplusan Forest | land | Provides reliable red or green mana.
- 1 Spire Garden | land | Provides both deck colors in multiplayer games.
- 1 Cragcrown Pathway // Timbercrown Pathway | land | Supplies the needed color on either face.
- 1 Rockfall Vale | land | Adds a red-green dual land to the mana base.
- 1 Rootbound Crag | land | Adds another efficient red-green source.
- 1 Game Trail | land | Helps supply either deck color.
- 1 Cavern of Souls | land | Provides colored mana for the deck’s Dragon creature type.
- 1 Mox Jasper | ramp | Provides an inexpensive mana source for accelerating the deck.
- 1 Jade Orb of Dragonkind | ramp | Develops mana while supporting Dragon spells.
- 1 Orb of Dragonkind | ramp | Provides color fixing and mana development.
- 1 Carnelian Orb of Dragonkind | ramp | Adds Dragon-focused mana acceleration.
- 1 Dragon's Hoard | ramp | Builds mana resources for the deck’s expensive creatures.
- 1 Dragonstorm Globe | ramp | Provides fixing and ramp for multicolor Dragon turns.
- 1 Scaled Nurturer | ramp | Supplies early creature-based mana acceleration.
- 1 Reckless Barbarian | ramp | Provides a burst of mana to reach larger plays.
- 1 Ganax, Astral Hunter | ramp | Turns the Dragon plan into additional mana resources.
- 1 Goldspan Dragon | ramp | Provides a Dragon threat that also advances mana.
- 1 Faithless Looting | draw | Provides efficient card selection and card flow.
- 1 Thrill of Possibility | draw | Turns a spare card into fresh cards.
- 1 Demand Answers | draw | Provides inexpensive card flow.
- 1 Sylvan Library | draw | Provides ongoing card selection.
- 1 Skullclamp | draw | Provides a compact source of cards.
- 1 Elemental Bond | draw | Rewards the deck for deploying substantial creatures.
- 1 Garruk's Uprising | draw | Turns the creature plan into ongoing card flow.
- 1 Beast Whisperer | draw | Provides cards as the deck develops its creatures.
- 1 Dragonborn Champion | draw | Supports combat-driven card flow.
- 1 Return of the Wildspeaker | draw | Provides a flexible refill for the creature-heavy plan.
- 1 Rishkar's Expertise | draw | Provides a large card refill around a major creature.
- 1 Lightning Greaves | interaction | Protects an important creature with a low-cost Equipment.
- 1 Swiftfoot Boots | interaction | Provides repeatable protection for key creatures.
- 1 Heroic Intervention | interaction | Protects the board from opposing disruption.
- 1 Snakeskin Veil | interaction | Offers efficient protection for a key permanent.
- 1 Tamiyo's Safekeeping | interaction | Protects a creature or land at a low cost.
- 1 Veil of Summer | interaction | Provides efficient defensive interaction.
- 1 Pyroblast | interaction | Provides low-cost stack interaction.
- 1 Tibalt's Trickery | interaction | Provides flexible stack interaction.
- 1 Draconic Roar | removal | Provides efficient Dragon-themed removal.
- 1 Dragon's Fire | removal | Provides efficient creature removal for the Dragon deck.
- 1 Piercing Exhale | removal | Adds a low-cost removal spell.
- 1 Spit Flame | removal | Provides Dragon-connected creature removal.
- 1 Molten Exhale | removal | Adds direct removal to clear blockers or threats.
- 1 Invasion of Tarkir // Defiant Thundermaw | removal | Combines Dragon-focused removal with a later Dragon body.
- 1 Dragon Tempest | removal | Adds Dragon-based removal pressure as creatures enter.
- 1 Glorybringer | removal | Provides a Dragon body with creature removal.
- 1 Scourge of Valkas | removal | Turns Dragon deployment into targeted removal pressure.
- 1 Dragonlord's Servant | synergy | Supports the deck’s Dragon-focused casting plan.
- 1 Dragonspeaker Shaman | synergy | Helps the deck deploy its Dragon spells.
- 1 Dragonkin Berserker | synergy | Rewards the deck for continuing to cast Dragons.
- 1 Firespitter Whelp | synergy | Adds a low-cost Dragon that supports the tribal plan.
- 1 Minion of the Mighty | synergy | Supports fast development into the deck’s Dragon threats.
- 1 Dragon Egg | synergy | Adds a Dragon-themed early permanent.
- 1 Dragon Hatchling | synergy | Adds another inexpensive Dragon for tribal density.
- 1 Kargan Dragonrider | synergy | Supports the deck’s Dragon-centered board presence.
- 1 Slumbering Dragon | synergy | Adds an early Dragon permanent that fits the creature plan.
- 1 Crucible of Fire | synergy | Strengthens the deck’s Dragon board.
- 1 Dragonloft Idol | threat | Provides an efficient Dragon threat.
- 1 Manaform Hellkite | threat | Provides a Dragon threat that rewards spellcasting.
- 1 Furnace Hellkite | threat | Adds a substantial Dragon attacker.
- 1 Thunderbreak Regent | threat | Provides a resilient Dragon threat for the tribal plan.
- 1 Stormbreath Dragon | threat | Adds a hasty evasive Dragon threat.
- 1 Thrakkus the Butcher | threat | Provides a powerful Dragon-focused combat threat.
- 1 Lathliss, Dragon Queen | threat | Builds a larger Dragon board presence.
- 1 Scourge of the Throne | threat | Adds a high-impact combat threat.
- 1 Hellkite Charger | threat | Provides an aggressive Dragon threat.
- 1 Twinflame Tyrant | threat | Adds a major combat-focused Dragon threat.
- 1 Terror of Mount Velus | threat | Provides a high-end Dragon finisher.
- 1 Utvara Hellkite | threat | Provides a top-end Dragon threat for prolonged games.
- 1 Breath Weapon | wipe | Provides a Dragon-friendly board-clearing effect.
- 1 Draconic Intervention | wipe | Provides a scalable board wipe for problematic creatures.
- 1 Balefire Dragon | wipe | Provides a Dragon threat that can clear opposing boards.

</details>

### 15. delegated commander, owned first

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 229 names.

Commander: Astarion, the Decadent.

Grade: bad, score 0.07, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $168.30 the whole deck.

**Summary:** Astarion leads a white-black lifegain deck that develops a creature board, layers lifegain-focused synergies with Angels and Equipment, and keeps cards flowing while it pressures opponents through combat. It wins by turning that sustained board presence into repeated attacks, backed by efficient removal and a small set of reset buttons. The deck gives up explosive combo speed for a steadier, board-centered plan that needs its creatures and synergies to remain established.

The quality model grades this deck below the precon baseline against the top lists of the format: more opening hands hold two to four lands than the norm, and that raises the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. The deck holds no fast mana, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: Most cards point at a W/B lifegain creature board with equipment and angel payoffs, but three unconditional board wipes (Austere Command, Fumigate, Vanquish the Horde) and the go-wide/equip-heavy creature plan work against each other, and the equipment subtheme is only loosely tied to the lifegain payoffs.
- PLAN theme_fit=yes: It is a Commander deck in white-black built around gaining life with a lifegain commander the builder picked, and its no-fast-mana, no-combo, mostly-uncompetitive card pool sits comfortably at bracket 2.
- PLAN useful_as_built=yes: 37 lands plus ten ramp pieces in a two-color mostly-basic mana base, a low curve, ~25 creatures and plenty of removal make this immediately playable with a clear combat-damage win path.
- PLAN summary_honest=yes: The summary accurately describes a creature-board lifegain deck with angel/equipment synergies, card draw, removal, a small number of wipes, and no fast-combo speed, and it even flags the deck's dependence on keeping creatures on board.
- [WARN] `not_owned`: Astarion, the Decadent: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.10 over 62 nonland cards

<details><summary>The deck list</summary>

- 1 Buster Sword | draw | Draw slot that supports the deck's proactive board plan.
- 1 Call of the Ring | draw | Draw slot for maintaining resources through the game.
- 1 Exemplar of Light | draw | Draw slot that also fits the deck's creature-focused plan.
- 1 Idol of Oblivion | draw | Draw slot for steady card access.
- 1 Inspiring Overseer | draw | Draw slot that contributes to the board.
- 1 Lembas | draw | Draw slot with a low commitment to deploy.
- 1 Mask of Memory | draw | Draw slot for a combat-oriented game plan.
- 1 Night's Whisper | draw | Draw slot for efficient resource refilling.
- 1 Puresteel Paladin | draw | Draw slot that complements the Equipment package.
- 1 Skullclamp | draw | Draw slot for converting board material into cards.
- 1 Tome of Legends | draw | Draw slot that works alongside the commander.
- 1 Bastion Protector | interaction | Interaction slot for protecting the central game plan.
- 1 Champion's Helm | interaction | Interaction slot for keeping a key legend safe.
- 1 Clever Concealment | interaction | Interaction slot for safeguarding the board.
- 1 Darksteel Plate | interaction | Interaction slot for durable protection.
- 1 Lightning Greaves | interaction | Interaction slot for protecting an important creature.
- 1 Swiftfoot Boots | interaction | Interaction slot for flexible creature protection.
- 18 Plains | land | Basic white mana for the deck's primary color needs.
- 14 Swamp | land | Basic black mana for the deck's secondary color needs.
- 1 Command Tower | land | Reliable multicolor land for the commander deck.
- 1 City of Brass | land | Flexible color-fixing land.
- 1 Exotic Orchard | land | Flexible multicolor mana source.
- 1 Path of Ancestry | land | Color-fixing land that fits the creature-heavy shell.
- 1 Takenuma, Abandoned Mire | land | Black mana source with additional utility.
- 1 Arcane Signet | ramp | Efficient mana fixing for the two-color deck.
- 1 Bender's Waterskin | ramp | Ramp slot for advancing the mana plan.
- 1 Blitzball | ramp | Ramp slot that helps deploy the deck's board.
- 1 Chromatic Lantern | ramp | Ramp and color-fixing support.
- 1 Commander's Sphere | ramp | Mana fixing that remains useful later.
- 1 Fellwar Stone | ramp | Low-cost mana acceleration.
- 1 Giada, Font of Hope | ramp | Ramp slot that fits the Angel presence.
- 1 Lotho, Corrupt Shirriff | ramp | Ramp slot in the deck's colors.
- 1 Sol Ring | ramp | Fast early mana acceleration.
- 1 Springleaf Drum | ramp | Fast low-cost mana acceleration.
- 1 Banishing Light | removal | Removal slot for handling a problematic permanent.
- 1 Bitter Triumph | removal | Removal slot for efficient answers.
- 1 Crib Swap | removal | Removal slot with broad creature coverage.
- 1 Dismember | removal | Removal slot for answering creatures efficiently.
- 1 Dispatch | removal | Removal slot that benefits from the artifact package.
- 1 Fatal Push | removal | Low-cost removal for early threats.
- 1 Generous Gift | removal | Flexible removal for varied permanent types.
- 1 Aerith Gainsborough | synergy | Lifegain synergy slot that contributes to the creature board.
- 1 Aettir and Priwen | synergy | Synergy slot supporting the Equipment subtheme.
- 1 Angel of Vitality | synergy | Lifegain synergy slot that fits the Angel suite.
- 1 Compassionate Healer | synergy | Lifegain synergy slot for the deck's core plan.
- 1 Crowd of True Believers | synergy | Synergy slot that builds the creature board.
- 1 Eastfarthing Farmer | synergy | Synergy slot for the lifegain-focused shell.
- 1 Kor Firewalker | synergy | Synergy slot that supports the deck's core theme.
- 1 Light of Promise | synergy | Synergy slot for rewarding the lifegain plan.
- 1 Night Nurse, Healer of Heroes | synergy | Lifegain synergy slot for the deck's central plan.
- 1 Prideful Feastling | synergy | Synergy slot that fits the lifegain shell.
- 1 Rosie Cotton of South Lane | synergy | Synergy slot for building a board around the theme.
- 1 Second Breakfast | synergy | Synergy slot that reinforces the deck's theme.
- 1 White Mage's Staff | synergy | Synergy slot that also supports the Equipment package.
- 1 Angel of Invention | threat | Threat slot for applying pressure on board.
- 1 Bill the Pony | threat | Threat slot that adds to the creature count.
- 1 Dawnhand Eulogist | threat | Threat slot for advancing the board presence.
- 1 Gollum, Silent Slinker // Meager Meal | threat | Threat slot that provides pressure in the midgame.
- 1 Lyra Dawnbringer | threat | Threat slot that fits the Angel package.
- 1 Minwu, White Mage | threat | Threat slot that supports the white creature base.
- 1 Rabaroo Troop | threat | Threat slot for maintaining combat pressure.
- 1 Reaping Willow | threat | Threat slot for a resilient board plan.
- 1 Rooftop Percher | threat | Threat slot that adds a meaningful body.
- 1 Shattered Angel | threat | Threat slot that reinforces the Angel suite.
- 1 Sneering Shadewriter | threat | Threat slot for black-side board pressure.
- 1 Victory's Herald | threat | Threat slot that helps close games through combat.
- 1 Austere Command | wipe | Wipe slot for resetting unfavorable boards.
- 1 Fumigate | wipe | Wipe slot for recovering from wide boards.
- 1 Vanquish the Horde | wipe | Wipe slot for efficient board control.

</details>

### 16. a tight budget, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 229 names.

Commander: Denethor, Ruling Steward.

Grade: bad, score 0.06, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $13.71 to buy, $35.32 the whole deck.

**Summary:** This aristocrats deck develops a creature-heavy board, layers sacrifice-focused synergy pieces, and uses attrition to pressure the table while keeping cards and mana flowing. It closes by making creature exchanges favorable over a long game, backed by broad removal and reset buttons when opponents establish stronger boards. The deck gives up explosive early wins and premium individual finishers in favor of a steady, budget-conscious creature engine.

The quality model grades this deck below the precon baseline against the top lists of the format: more opening hands hold two to four lands than the norm, and that raises the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: The list is a coherent sacrifice-and-drain build: a dozen sac outlets (Woe Strider, Yahenni, Bartolom\u00e9, Ayli, Skullport Merchant, Aron) feed a full suite of drain payoffs (Blood Artist, Zulaport, Falkenrath Noble, Bastion of Remembrance, Vindictive Vampire), and even the three wraths convert into drain damage rather than fighting the plan, though token fodder is a bit thin.
- PLAN theme_fit=partly: It is genuinely an aristocrats Commander deck at a bracket-2 power level with a fitting W/B commander, but the collection of staples (Skullclamp, Blood Artist, Swords to Plowshares, Sol Ring, Austere Command, Shilgengar, June) plausibly runs well past the stated $25 ceiling.
- PLAN useful_as_built=yes: 37 lands in a two-color deck plus ten mana rocks, a reasonable curve, eleven draw pieces and seven spot removal spells make this immediately playable, with only minor slop like Secluded Courtyard/Unclaimed Territory in a tribally scattered creature base.
- PLAN summary_honest=yes: The summary's claims\u2014creature-heavy board, sacrifice synergy, card and mana flow, broad removal plus resets, no premium finishers\u2014all match what is actually in the 99, and it openly flags the deck's weaknesses rather than overselling.
- [WARN] `not_owned`: Blood Artist: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bastion of Remembrance: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Elas il-Kor, Sadistic Pilgrim: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Falkenrath Noble: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Sivriss, Nightmare Speaker: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Skullport Merchant: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Smothering Abomination: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Disciple of Bolas: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Woe Strider: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Yahenni, Undying Partisan: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Zulaport Cutthroat: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Aron, Benalia's Ruin: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ayli, Eternal Pilgrim: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Baron Bertram Graywater: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bartolomé del Presidio: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Elder Arthur Maxson: the deck needs 1, the collection has 0
- [WARN] `not_owned`: High-Society Hunter: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Lord Skitter's Butcher: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Old Flitterfang: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Shilgengar, Sire of Famine: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vampire Gourmand: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vindictive Vampire: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.15 over 62 nonland cards

<details><summary>The deck list</summary>

- 18 Plains | land | Provides the primary basic white land base.
- 14 Swamp | land | Provides the primary basic black land base.
- 1 Command Tower | land | Provides a multicolor land slot in the mana base.
- 1 Exotic Orchard | land | Provides a multicolor land slot in the mana base.
- 1 Secluded Courtyard | land | Provides a creature-focused land slot in the mana base.
- 1 Unclaimed Territory | land | Provides a creature-focused land slot in the mana base.
- 1 Spire of Industry | land | Provides a nonbasic land slot alongside the artifact package.
- 1 Sol Ring | ramp | Provides efficient mana development for the creature-heavy plan.
- 1 Arcane Signet | ramp | Provides reliable color fixing and mana development.
- 1 Chromatic Lantern | ramp | Supports mana development and fixing.
- 1 Fellwar Stone | ramp | Provides an inexpensive mana-development slot.
- 1 Commander's Sphere | ramp | Provides mana development for the midgame.
- 1 Thought Vessel | ramp | Provides an artifact-based mana-development slot.
- 1 Deadly Dispute | ramp | Fits the sacrifice-oriented plan while filling a ramp slot.
- 1 Ring of the Lucii | ramp | Provides an artifact-based mana-development slot.
- 1 Relic of Legends | ramp | Provides mana development for the legendary creature package.
- 1 White Auracite | ramp | Provides an inexpensive artifact ramp slot.
- 1 Skullclamp | draw | Provides a draw engine for the creature-focused plan.
- 1 Ahriman | draw | Provides a creature-based draw slot.
- 1 Beetle-Headed Merchants | draw | Provides an inexpensive creature-based draw slot.
- 1 Buzzard-Wasp Colony | draw | Provides a creature-based draw slot.
- 1 Circle of Power | draw | Provides a dedicated draw slot.
- 1 Cirith Ungol Patrol | draw | Provides a creature-based draw slot.
- 1 Grave Venerations | draw | Provides a dedicated draw slot for the creature deck.
- 1 Inspiring Overseer | draw | Provides a creature-based draw slot.
- 1 June, Bounty Hunter | draw | Provides an inexpensive creature-based draw slot.
- 1 Nasty End | draw | Provides a draw slot that fits the creature strategy.
- 1 Tome of Legends | draw | Provides a persistent artifact draw slot.
- 1 Airbending Lesson | interaction | Provides a flexible interaction slot.
- 1 Bastion Protector | interaction | Provides creature-based interaction for the board.
- 1 Bronze Guardian | interaction | Provides artifact-based interaction for the board.
- 1 Frontline Medic | interaction | Provides a creature-based interaction slot.
- 1 Gift of Immortality | interaction | Provides protective interaction for a key creature.
- 1 Swiftfoot Boots | interaction | Provides protective interaction for a key creature.
- 1 Bitter Triumph | removal | Provides an efficient removal slot.
- 1 Claim the Precious | removal | Provides a dedicated removal slot.
- 1 Crib Swap | removal | Provides flexible creature removal.
- 1 Dispatch | removal | Provides an efficient removal slot alongside the artifacts.
- 1 Generous Gift | removal | Provides broad removal coverage.
- 1 Infernal Grasp | removal | Provides direct creature removal.
- 1 Swords to Plowshares | removal | Provides efficient creature removal.
- 1 Austere Command | wipe | Provides a flexible board-wipe slot.
- 1 Dusk // Dawn | wipe | Provides a creature-focused board-wipe slot.
- 1 Fumigate | wipe | Provides a dedicated board-wipe slot.
- 1 Blood Artist | synergy | Provides a core creature-synergy piece for the aristocrats plan.
- 1 Bastion of Remembrance | synergy | Provides a persistent synergy piece for the aristocrats plan.
- 1 Elas il-Kor, Sadistic Pilgrim | synergy | Provides a creature-based synergy piece for the aristocrats plan.
- 1 Falkenrath Noble | synergy | Provides a creature-based synergy piece for the aristocrats plan.
- 1 Gollum the Abandoned | synergy | Provides an owned creature-synergy piece.
- 1 Gollum, Patient Plotter | synergy | Provides an owned creature-synergy piece.
- 1 Sivriss, Nightmare Speaker | synergy | Provides a creature-synergy piece for the aristocrats plan.
- 1 Skullport Merchant | synergy | Provides a creature-synergy piece for the aristocrats plan.
- 1 Smothering Abomination | synergy | Provides a creature-synergy piece for the aristocrats plan.
- 1 Disciple of Bolas | synergy | Provides a creature-synergy piece for the aristocrats plan.
- 1 Woe Strider | synergy | Provides a creature-synergy piece for the aristocrats plan.
- 1 Yahenni, Undying Partisan | synergy | Provides a creature-synergy piece for the aristocrats plan.
- 1 Zulaport Cutthroat | synergy | Provides a core creature-synergy piece for the aristocrats plan.
- 1 Aron, Benalia's Ruin | threat | Provides a low-cost legendary creature threat.
- 1 Ayli, Eternal Pilgrim | threat | Provides a low-cost legendary creature threat.
- 1 Baron Bertram Graywater | threat | Provides a creature threat for the aristocrats board.
- 1 Bartolomé del Presidio | threat | Provides a creature threat for the aristocrats board.
- 1 Bill the Pony | threat | Provides an owned creature threat.
- 1 Elder Arthur Maxson | threat | Provides a legendary creature threat.
- 1 High-Society Hunter | threat | Provides a creature threat for the board.
- 1 Lord Skitter's Butcher | threat | Provides a creature threat for the board.
- 1 Old Flitterfang | threat | Provides a creature threat for the board.
- 1 Shilgengar, Sire of Famine | threat | Provides a larger legendary creature threat.
- 1 Vampire Gourmand | threat | Provides a low-cost creature threat.
- 1 Vindictive Vampire | threat | Provides a creature threat that fits the deck's creature core.

</details>

### 17. upgrade a precon, owned first

Format: Commander. Theme: goblins. Pool: owned_first. Shortlist: 196 names.

Commander: Zada, Hedron Grinder.

Grade: typical, score 0.56, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band, profile_off_band, two_card_combo; then profile_off_band. Block findings: 0.

Cost: $78.99 to buy, $259.13 the whole deck.

**Summary:** A mono-red Goblin storm deck that builds a wide battlefield, converts creature density into burst mana, and uses targeted cantrips to turn a single spell into a team-wide surge of cards and combat power. Goblin lords, token makers, sacrifice outlets, and enter-the-battlefield damage provide several ways to pressure the table, while protective equipment and flexible removal keep the central engine online.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Every piece supports the Zada plan: cheap targeted cantrips (Crimson Wisps, Expedite, Fists of Flame, Wild Ride), a wide goblin token base to copy them onto, rituals and Storm-Kiln Artist for burst mana, and storm/aristocrat payoffs (Grapeshot, Empty the Warrens, Impact Tremors, Boggart Shenanigans) to convert the draws into damage.
- PLAN theme_fit=yes: It is a mono-red 100-card Commander deck led by a legal goblin-storm commander, upgraded with genuine goblin lords, tutors and storm engines while staying clear of tutor-chains and infinite combos, which sits sensibly at bracket 3.
- PLAN useful_as_built=yes: 38 lands plus Sol Ring, Arcane Signet, Ruby Medallion and several rituals support a low curve, and the deck has multiple realistic finishes (token swarm with lords/Bushwhacker, Impact Tremors/Shenanigans drain, Grapeshot, Empty the Warrens) rather than a single fragile line.
- PLAN summary_honest=yes: placeholder
- [INFO] `curve_summary`: average mana value 2.64 over 61 nonland cards
- [INFO] `precon_cards_restored`: the deck was 2 cards short of the Goblin Storm precon share, so the builder put 2 cards back
- [WARN] `profile_off_band`: the removal count is 13, and bracket 3 wants 6 to 12
- [INFO] `repair_kept_earlier`: the repair turn answered a deck that failed a block check, so the deck before it stands with its warnings

<details><summary>The deck list</summary>

- 1 Ancestors' Aid | synergy | A cheap targeting spell that scales across the board with the commander.
- 1 Ancestral Anger | draw | Efficient cantrip that rewards targeting a creature.
- 1 Arcane Signet | ramp | Reliable early mana acceleration.
- 1 Arena of Glory | land | Creature land utility that supports aggressive turns.
- 1 Battle Hymn | ramp | Converts a wide Goblin board into explosive mana.
- 1 Battle-Scarred Goblin | removal | Goblin body with repeatable creature control utility.
- 1 Blasphemous Act | wipe | Efficient emergency reset against creature-heavy boards.
- 1 Boggart Shenanigans | removal | Turns Goblin deaths into direct damage and supports sacrifice play.
- 1 Brightstone Ritual | ramp | Produces a large burst of red mana from the creature swarm.
- 1 Broadside Bombardiers | removal | A Goblin threat that converts expendable permanents into removal.
- 1 Castle Embereth | land | Mana source with a late-game team-pump option.
- 1 Champion's Helm | interaction | Protects the commander while improving combat pressure.
- 1 Chaos Warp | removal | Flexible answer to problematic permanents.
- 1 Conspicuous Snoop | synergy | Goblin-focused card advantage and topdeck utility.
- 1 Crimson Wisps | draw | A cheap cantrip that grants haste and scales with creature targets.
- 1 Daring Discovery | synergy | Supports the deck's spell-driven Goblin game plan.
- 1 Darksteel Plate | interaction | Durable protection for the commander or a key Goblin.
- 1 Den of the Bugbear | land | Mana source that becomes a Goblin-producing threat.
- 1 Dragon Fodder | synergy | Creates multiple bodies for swarm, sacrifice, and targeting effects.
- 1 Empty the Warrens | wincon | Produces a large Goblin force after a spell-heavy turn.
- 1 Expedite | draw | Cheap haste cantrip that works especially well on key creatures.
- 1 Faithless Looting | draw | Efficient filtering that helps find action and stock the graveyard.
- 1 Fists of Flame | draw | Card draw and a powerful combat boost for a targeted creature.
- 1 Forgotten Cave | land | Red mana source with cycling when a land is unneeded.
- 1 Frontline Heroism | synergy | Supports the deck's creature-focused combat plan.
- 1 Gempalm Incinerator | removal | Goblin-scaled creature removal with cycling utility.
- 1 General Kreat, the Boltbringer | synergy | A Goblin-focused payoff for building a broad board.
- 1 Goblin Bombardment | removal | Converts tokens and dying Goblins into reliable damage.
- 1 Goblin Burrows | land | Land-based combat boost for an attacking Goblin.
- 1 Goblin Bushwhacker | synergy | Haste and power boost help the token army close games.
- 1 Goblin Chieftain | synergy | Lord effect and haste make the Goblin army immediately dangerous.
- 1 Goblin Dark-Dwellers | threat | A substantial Goblin threat that recasts a useful spell.
- 1 Goblin Fireleaper | removal | Goblin creature that can pick off opposing creatures.
- 1 Goblin Glasswright // Craft with Pride | ramp | Goblin body with mana-production utility.
- 1 Goblin Lackey | synergy | Accelerates Goblins onto the battlefield after connecting.
- 1 Goblin Matron | synergy | Finds the most useful Goblin for the current board state.
- 1 Goblin Negotiation | removal | Goblin-themed removal that keeps opponents' boards manageable.
- 1 Goblin Trashmaster | removal | Goblin lord that also answers artifacts.
- 1 Goblin Warchief | synergy | Reduces Goblin costs and grants the team haste.
- 1 Grapeshot | removal | Storm-based removal that benefits from spell-heavy turns.
- 1 Hidden Volcano | land | Red mana source with useful land utility.
- 1 Idol of Oblivion | draw | Efficient repeatable card draw alongside token production.
- 1 Impact Tremors | wincon | Turns every Goblin and token entry into opponent-wide damage.
- 1 Kher Keep | land | Land that creates expendable creature tokens.
- 1 Krenko's Command | synergy | Low-cost source of multiple Goblin bodies.
- 1 Krenko, Mob Boss | threat | Major Goblin token engine and primary battlefield threat.
- 1 Dwarven Mine | interaction | the deck upgrades a precon, and this card is one the precon holds
- 26 Mountain | land | Consistent red mana for a mono-red spell and Goblin strategy.
- 1 Mana Geyser | ramp | Explosive ritual mana for high-impact turns.
- 1 Mogg War Marshal | synergy | Multiple Goblin bodies from one card fuel the swarm plan.
- 1 Spinerock Knoll | synergy | the deck upgrades a precon, and this card is one the precon holds
- 1 Pashalik Mons | removal | Goblin payoff that turns token production and sacrifice into damage.
- 1 Past in Flames | synergy | Lets the deck reuse a graveyard full of cheap spells.
- 1 Quest for the Goblin Lord | synergy | Permanent Goblin payoff that rewards a growing board.
- 1 Reliquary Tower | land | Utility land that preserves a large hand after draw-heavy turns.
- 1 Roaming Throne | synergy | Amplifies chosen Goblin triggered abilities.
- 1 Ruby Medallion | ramp | Reduces the cost of the deck's many red spells.
- 1 Rundvelt Hordemaster | synergy | Goblin lord and card-advantage engine for creature-heavy games.
- 1 Sazacap's Brew | draw | Cheap spell that supplies card selection.
- 1 Seething Song | ramp | Burst mana enables explosive Goblin and spell turns.
- 1 Shinka, the Bloodsoaked Keep | land | Red source with combat utility for legendary creatures.
- 1 Siege-Gang Commander | removal | Creates a Goblin squad and turns it into direct damage.
- 1 Siege-Gang Lieutenant | removal | Provides more Goblin bodies and sacrifice-based reach.
- 1 Skirk Prospector | ramp | Sacrifices Goblins for burst mana and enables explosive sequencing.
- 1 Skullclamp | draw | Exceptional card draw with expendable Goblin tokens.
- 1 Sol Ring | ramp | Efficient universal mana acceleration.
- 1 Spreading Insurrection | wincon | High-impact spell that can convert opposing creatures into a finishing attack.
- 1 Storm-Kiln Artist | ramp | Turns the deck's many instants and sorceries into mana resources.
- 1 Swiftfoot Boots | interaction | Protects a key creature while granting haste.
- 1 Throne of Eldraine | ramp | Mana rock that supports the deck's long-game resources.
- 1 Vandalblast | wipe | Strong artifact cleanup with an overloaded reset option.
- 1 War Room | land | Land-based card draw when mana is available.
- 1 Wild Ride | synergy | Combat-oriented spell that supports aggressive creature turns.
- 1 Witch's Mark | draw | Filtering spell that can also grant haste to a key creature.

</details>

### 18. upgrade a precon, any card

Format: Commander. Theme: turtles. Pool: any_card. Shortlist: 300 names.

Commander: Heroes in a Half Shell.

Grade: typical, score 0.51, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: yes, for precon_share, profile_off_band, profile_off_band, profile_off_band, profile_off_band, profile_off_band, profile_off_band; then profile_off_band. Block findings: 0.

Cost: $84.92 to buy, $84.92 the whole deck.

**Summary:** A five-color Turtle-themed creature deck that develops its board through flavorful heroes, allies, and mutant threats, using broad mana fixing to support its colors. It balances combat-forward pressure with artifact utility, counter synergies, creature protection, targeted answers, and several ways to reset crowded battlefields.

The quality model grades this deck typical against the top lists of the format: the cards pair the way the top lists pair them, and that raises the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core is coherent enough — go-wide TMNT creatures with +1/+1 counter payoffs (Corpsejack Menace, Donatello, Level Up, Biogenic Ooze) and combat pressure — but four sweepers plus several tap-out utility artifacts pull against a board-centric creature plan, leaving it more precon goodstuff than a single tight line (Vigor only partially redeems the Blasphemous Act tension).
- PLAN theme_fit=yes: It keeps Heroes in a Half Shell as commander, retains the TMNT flavor shell in five colors, and layers in modest bracket-3-appropriate upgrades (Sol Ring, Arcane Signet, Cultivate, Farseek, Assassin's Trophy, Vigor, Corpsejack Menace) without pushing into combo or fast-mana territory.
- PLAN useful_as_built=yes: 99 cards with 39 lands, tri-lands, Command Tower, Path of Ancestry, Grand Coliseum and Chromatic Lantern make the five-color mana workable, the curve is supported by real ramp, and there are plenty of creature threats plus removal to close and interact.
- PLAN summary_honest=yes: Placeholder
- [INFO] `curve_summary`: average mana value 3.40 over 60 nonland cards
- [WARN] `profile_off_band`: the land count is 39, and bracket 3 wants 34 to 38
- [WARN] `profile_off_band`: the ramp count is 7, and bracket 3 wants 8 to 13
- [WARN] `profile_off_band`: the draw count is 2, and bracket 3 wants 8 to 14
- [WARN] `profile_off_band`: the removal count is 3, and bracket 3 wants 6 to 12
- [WARN] `profile_off_band`: the interaction count is 0, and bracket 3 wants 4 to 12

<details><summary>The deck list</summary>

- 1 Acidic Slime | removal | Flexible permanent interaction attached to a creature.
- 1 April O'Neil, Live on the Scene | draw | Provides thematic card advantage.
- 1 Arcade Cabinet | other | Retains the precon’s flavorful support package.
- 1 Arcane Signet | ramp | Reliable early color fixing.
- 1 Assassin's Trophy | removal | Answers any problematic permanent.
- 1 Baxter, Fly in the Ointment | other | Retains the precon’s villain package.
- 1 Bebop, Skull & Crossbones | other | Retains the precon’s mutant-themed cast.
- 1 Big Mother Mouser | other | Retains a thematic artifact creature.
- 1 Biogenic Ooze | threat | Supplies a resilient board presence.
- 1 Blasphemous Act | wipe | Efficient emergency battlefield reset.
- 1 Casey Jones, Back Alley Brute | threat | Retains a thematic combat threat.
- 1 Chromatic Lantern | ramp | Fixes the deck’s broad color requirements.
- 1 Coin of Mastery | other | Retains the precon’s themed utility package.
- 1 Continue? | other | Retains thematic utility from the original list.
- 1 Corpsejack Menace | synergy | Enhances counter-based creature development.
- 1 Cultivate | ramp | Provides dependable land ramp and color access.
- 1 Dimension X Pizzasaur | other | Retains a thematic artifact creature.
- 1 Donatello, the Brains | synergy | Supports the deck’s Turtle-focused synergies.
- 1 Double Jump // Flying Kick | other | Retains flexible thematic utility.
- 1 Electric Seaweed | other | Retains the precon’s defensive utility.
- 1 Endless Foot Assault | other | Retains the themed combat package.
- 1 Everything Pizza | other | Retains flavorful artifact support.
- 1 Exploding Barrel | other | Retains thematic utility and interaction.
- 1 Fast Forward | other | Retains the precon’s themed spell package.
- 1 Farseek | ramp | Finds dual land types while accelerating mana.
- 1 Foot Chopper | other | Retains the precon’s Equipment package.
- 1 Harmonize | draw | Provides straightforward card advantage.
- 1 Here Comes a New Hero! | other | Retains thematic support for the creature core.
- 1 High Score | other | Retains the precon’s value engine.
- 1 Irma, Part-Time Mutant | other | Retains a thematic supporting character.
- 1 Krang, the All-Powerful | threat | Retains the precon’s major thematic threat.
- 1 Leatherhead, Iron Gator | threat | Retains a thematic creature threat.
- 1 Leonardo, the Balance | threat | Retains a central Turtle combat threat.
- 1 Lessons from Life | other | Retains the precon’s themed utility.
- 1 Level Up | synergy | Retains a thematic enhancement piece.
- 1 Lita, Little Orphan Amphibian | synergy | Supports Turtle-focused synergies.
- 1 Mole Module | other | Retains the precon’s artifact package.
- 1 Mona Lisa, Science Geek | other | Retains a thematic supporting character.
- 1 Ninja Pizza | other | Retains the precon’s flavorful support package.
- 1 Rat King, Pale Piper | other | Retains the precon’s villain package.
- 1 Ray Fillet, Wave Warrior | threat | Retains a thematic creature threat.
- 1 Roadkill Rodney | other | Retains the precon’s artifact creature package.
- 1 Rocksteady, Mutant Marauder | threat | Retains a thematic creature threat.
- 1 Shellshock | other | Retains thematic instant-speed utility.
- 1 Shredder, Shadow Master | other | Retains the precon’s villain package.
- 1 Sol Ring | ramp | Provides efficient early acceleration.
- 1 Special Move | other | Retains the themed utility package.
- 1 Splinter, the Mentor | other | Retains a central thematic supporting character.
- 1 Steelbane Hydra | removal | Provides repeatable artifact and enchantment interaction.
- 1 Super Combo | other | Retains the precon’s themed payoff package.
- 1 Swift Demise | other | Retains thematic instant-speed utility.
- 1 Tempestra, Dame of Games | other | Retains a thematic supporting character.
- 1 Together Forever | synergy | Helps preserve important creatures through removal.
- 1 Tokka & Rahzar, Unsupervised | ramp | Provides Turtle-themed mana acceleration.
- 1 Turtles in Time | wipe | Retains a thematic battlefield reset.
- 1 Vanquish the Horde | wipe | Provides an efficient creature reset.
- 1 Vigor | synergy | Makes creature combat more favorable and supports counters.
- 1 Voracious Hydra | threat | Offers a scalable creature and removal option.
- 1 Wave Goodbye | wipe | Retains a thematic battlefield reset.
- 1 Elvish Mystic | ramp | Adds early mana acceleration for stronger development.
- 1 Ash Barrens | land | Fixes basic-land access when needed.
- 1 Big Apple, 3 a.m. | land | Retains a thematic land from the original mana base.
- 1 Cinder Glade | land | Provides red and green mana.
- 1 Command Tower | land | Provides flexible commander color fixing.
- 1 Escape Tunnel | land | Finds a needed basic land.
- 1 Evolving Wilds | land | Provides basic-land fixing.
- 1 Grand Coliseum | land | Provides access to every deck color.
- 1 Hidden Hideout | land | Retains a thematic land from the original mana base.
- 1 Path of Ancestry | land | Fixes colors for the Turtle creature core.
- 1 Rain-Slicked Copse | land | Provides green and blue mana.
- 1 Rootbound Crag | land | Provides red and green mana.
- 1 Sodden Verdure | land | Provides green and blue mana.
- 1 Thriving Grove | land | Provides flexible color fixing.
- 1 Thriving Isle | land | Provides flexible color fixing.
- 1 Thriving Moor | land | Provides flexible color fixing.
- 1 Turtle Lair | land | Retains the deck’s signature thematic land.
- 1 Undergrowth Stadium | land | Provides black and green mana.
- 1 Vernal Fen | land | Provides black and green mana.
- 1 Vibrant Cityscape | land | Provides broad color fixing.
- 5 Forest | land | Supplies green mana for early acceleration and Turtle spells.
- 5 Island | land | Supplies blue mana for the deck’s utility and creatures.
- 4 Plains | land | Supplies white mana for support spells.
- 3 Swamp | land | Supplies black mana for the villain and utility package.
- 3 Mountain | land | Supplies red mana for combat and removal effects.

</details>

### 19. the Hobbit family, two colours

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 174 names.

Commander: Thranduil, the Elvenking.

Grade: bad, score 0.01, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band; then profile_off_band. Block findings: 0.

Cost: $341.38 to buy, $341.38 the whole deck.

**Summary:** This Sultai Elf deck develops a woodland court of mana creatures, legendary allies, and resilient tribal bodies before turning that board into sustained card advantage and combat pressure. Its game plan combines efficient disruption with creature-based acceleration, allowing it to protect its key permanents, contain opposing threats, and rebuild well after a sweeping exchange.

The quality model grades this deck below the precon baseline against the top lists of the format: more opening hands hold two to four lands than the norm, and that raises the grade. Many of the cards appear in no top list, and that lowers the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: There is a loose green-black-blue creature goodstuff shell with some ramp and removal, but the elf-lord payoffs (Elvish Archdruid, Woodland Weavemaster) sit next to a dozen non-Elf spiders, bears and beasts and a pile of unrelated card-draw and wipes, so the pieces only half pull the same direction.
- PLAN theme_fit=partly: The commander, colors and Sultai identity match the request, but a large share of the list is drawn from the earlier LOTR set and from ordinary Magic staples (Night's Whisper, Languish, Elvish Mystic, Colossal Whale) rather than the Hobbit set the person asked for.
- PLAN useful_as_built=partly: It has creatures, removal and win conditions and can be shuffled up, but the deck runs one card long and its three-color mana is 36 untapped basics with almost no fixing, which will strand double-colored cards like Elvish Archdruid and Languish.
- PLAN summary_honest=partly: The claims of disruption, ramp and card advantage are borne out, but calling it a "tribal" Elf deck overstates a list where most creatures are not Elves and there is no real tribal engine beyond one lord.
- [INFO] `curve_summary`: average mana value 3.24 over 63 nonland cards
- [WARN] `profile_off_band`: the worst color's share of the sources it needs is 0.71, and bracket 3 wants 0.85 or more (sources of requirement: U 15.5 of 19, B 13.5 of 19, G 22 of 20)

<details><summary>The deck list</summary>

- 1 Elvish Visionary | draw | Efficient creature-based card advantage that also supports the Elf theme.
- 1 Hithlain Knots | draw | Flexible instant-speed card draw.
- 1 Ithilien Kingfisher | draw | Creature-based card advantage.
- 1 Key to the Side-Door | draw | Artifact card-advantage piece.
- 1 Last March of the Ents | draw | High-impact refill for a creature-heavy board.
- 1 Lórien Revealed | draw | Card selection early and card advantage later.
- 1 Night's Whisper | draw | Low-cost, reliable card draw.
- 1 Palantír of Orthanc | draw | Repeatable value engine.
- 1 Plunder the Trollshaws | draw | Instant-speed card advantage.
- 1 Thrór's Map | draw | Thematic artifact that provides ongoing value.
- 1 Uncover the Moon-Letters | draw | Enchantment-based card advantage.
- 1 Bilbo's Ring | interaction | Protects an important creature or disrupts opposing combat.
- 1 Confusticate and Bebother | interaction | Versatile stack interaction.
- 1 Elrond, Moon-Reader | interaction | Thematic interactive legend.
- 1 Mithril Coat | interaction | Protects the commander or a key Elf.
- 1 My Precious // Allure of Power | interaction | Equipment protection with a useful Adventure mode.
- 1 Old Fat Spider Can't See Me | interaction | Provides defensive and disruptive utility.
- 1 Stern Scolding | interaction | Efficient answer to cheap opposing creatures and spells.
- 1 Thranduil's Decree | interaction | Thematic flexible disruption.
- 1 Arcane Signet | ramp | Reliable fixing for all commander colors.
- 1 Delighted Halfling | ramp | Early mana creature that helps deploy legendary spells.
- 1 Elven Chorus | ramp | Turns the Elf board into a substantial mana engine.
- 1 Elvish Archdruid | ramp | Elf payoff and powerful creature-based mana production.
- 1 Elvish Mystic | ramp | Efficient early acceleration.
- 1 Mox Amber | ramp | Fast legendary-focused fixing and acceleration.
- 1 Wayfarer's Bauble | ramp | Stable land-based mana acceleration.
- 1 Wood Elves | ramp | Finds a Forest while advancing the Elf board.
- 1 Woodland Weavemaster | ramp | Thematic mana creature.
- 1 Through the Forest Gate | ramp | Land-ramp spell that strengthens green access.
- 1 Bitter Downfall | removal | Flexible black spot removal.
- 1 Bilbo's Deadly Slice | removal | Efficient targeted removal.
- 1 Crude Bent Blade | removal | Reusable thematic removal tool.
- 1 Enchanted River's Grasp | removal | Blue permanent-based answer.
- 1 Giant's Boulder | removal | Colorless removal accessible through any mana mix.
- 1 Merciless Executioner | removal | Creature-based sacrifice removal.
- 1 Orcish Bowmasters | removal | Efficient creature removal with pressure.
- 1 Quarrel | removal | Low-cost targeted answer.
- 1 The Black Arrow | removal | Thematic removal attached to an artifact.
- 1 Gnashing of Teeth | wipe | Board reset for creature-heavy opposing positions.
- 1 Languish | wipe | Efficient creature sweeper that can preserve larger threats.
- 1 Raise the Palisade | wipe | Tribal-aware sweeper that can favor the Elf board.
- 1 Arwen, Weaver of Hope | synergy | Legendary Elf payoff that builds resilient boards.
- 1 Boughside Wanderers | synergy | Elf body that supports the tribal plan.
- 1 Cantankerous Keepers | synergy | Thematic Elf creature for board development.
- 1 Galadhrim Guide | synergy | Elf support creature.
- 1 Galion, Elvenking's Butler | synergy | Legendary Elf that reinforces the court theme.
- 1 Guardian of the Halls | synergy | Elf board presence that benefits tribal payoffs.
- 1 Lothlórien Lookout | synergy | Thematic Elf support body.
- 1 Mirkwood Meditator | synergy | Elf creature for tribal density.
- 1 Mirkwood Nurturer | synergy | Elf creature supporting the woodland strategy.
- 1 Mirkwood Pathmaker | synergy | Elf creature that broadens tribal board presence.
- 1 Attercop | threat | Large thematic creature that applies pressure.
- 1 Colossal Whale | threat | Large blue threat with disruptive upside.
- 1 Gigantic Big Bear | threat | Straightforward green combat threat.
- 1 Haunt of the Dead Marshes | threat | Black creature threat with thematic flavor.
- 1 Mirkwood Elk | threat | Efficient green creature for combat pressure.
- 1 Nasty Little Rabbit | threat | Low-curve creature that contributes to battlefield pressure.
- 1 Nighthowl Pursuer | threat | Aggressive woodland threat.
- 1 Old Fat Spider | threat | Creature threat that also contributes value.
- 1 Ravenhill Flock | threat | Evasive creature pressure.
- 1 Rhovanion Rampager | threat | Green combat-focused threat.
- 1 Willow-Wind | threat | Elemental threat that broadens the creature suite.
- 1 Wilderland Scrounger | threat | Woodland creature that supports sustained combat.
- 16 Forest | land | Primary green mana base for the Elf-heavy core and green acceleration.
- 11 Island | land | Reliable blue mana for interaction and value spells.
- 9 Swamp | land | Reliable black mana for removal and sweepers.

</details>

### 20. the Hobbit family, a delegated commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 77 names.

Commander: Smaug the Magnificent.

Grade: bad, score 0.05, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band. Block findings: 0.

Cost: $514.88 to buy, $514.88 the whole deck.

**Summary:** Smaug leads a red Dragon-and-Treasure assault built around accelerating into imposing creatures, equipping the best attacker, and forcing damage through crowded battlefields. Artifact tools supply resilient resources and protection, while a deep suite of burn, creature answers, and sweeping effects keeps opposing boards from stabilizing. The deck wins by turning mana advantages into devastating combat turns backed by dragonfire.

The quality model grades this deck below the precon baseline against the top lists of the format: more opening hands hold two to four lands than the norm, and that raises the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: There is a discernible mono-red core of ramp, big creatures and a large equipment package that fits the "equip the best attacker" line, but the list also scatters into unrelated goblin/dwarf bodies, burn, and wipes that dilute rather than reinforce any single axis, and the Treasure payoff count is thin.
- PLAN theme_fit=no: The person asked for a Dragons deck; apart from Smaug and Cavern-Hoard Dragon the creature base is dwarves, goblins, orcs and giants, so this is a mono-red equipment/beatdown deck wearing a Smaug hat rather than the requested Dragon deck.
- PLAN useful_as_built=yes: 36 lands in a single color, a reasonable curve with cheap ramp, plentiful removal, and many creature and equipment threats mean it can be shuffled up and played as it stands.
- PLAN summary_honest=partly: The equipment, ramp, removal and sweeper claims are backed by the list, but calling it a "Dragon-and-Treasure assault" oversells a deck with essentially two Dragons and only a handful of Treasure makers/payoffs.
- [INFO] `curve_summary`: average mana value 3.30 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Arcane Signet | ramp | Reliable color fixing and mana acceleration.
- 1 Bag End Banquet | ramp | Treasure-focused acceleration for expensive threats.
- 1 Burn, Burn, Tree and Fern | ramp | Advances mana while supporting the deck's red theme.
- 1 Cavern-Hoard Dragon | ramp | A Dragon that converts combat success into mana resources.
- 1 Dragon's Desire | ramp | Provides additional mana for deploying large creatures.
- 1 Fíli and Kíli, Joyous | ramp | Creature-based acceleration with tribal flavor.
- 1 Long-Bodied Grey Dog | ramp | Early mana development attached to a creature.
- 1 Mox Amber | ramp | Cheap legendary-dependent acceleration.
- 1 Orcrist, Goblin-cleaver | ramp | Equipment-based mana support for attacking creatures.
- 1 The Reaver Cleaver | ramp | Turns combat damage into a substantial Treasure supply.
- 1 Balin, Loremaster | draw | Provides repeatable card advantage.
- 1 Key to the Side-Door | draw | Helps keep cards flowing while supporting artifact play.
- 1 Palantír of Orthanc | draw | Long-game card advantage engine.
- 1 Ragged Short Spear | draw | An inexpensive equipment source of card access.
- 1 Thrór's Map | draw | Generates card advantage while fitting the artifact package.
- 1 Óin the Brave | draw | Creature-based card advantage for the midgame.
- 1 The One Ring | draw | A resilient source of sustained card advantage.
- 1 Tidings of War | draw | Refills resources while advancing the aggressive game plan.
- 1 Battle-Scarred Goblin | removal | Creature removal that can also pressure opponents.
- 1 Fire of Orthanc | removal | Flexible spell-based answer to troublesome permanents.
- 1 Gandalf, Spark Starter | removal | A repeatable removal body with useful board presence.
- 1 Goblin Cratermaker | removal | Efficient answer to creatures and artifacts.
- 1 Goblin Fireleaper | removal | Adds direct creature control to the board.
- 1 Improvised Club | removal | Efficient instant-speed removal.
- 1 Inferno Titan | removal | A major threat that also clears smaller creatures.
- 1 Pinecone Strike | removal | Low-cost interaction for key targets.
- 1 Smite the Deathless | removal | Clean answer to durable opposing threats.
- 1 Call Forth the Tempest | wipe | Resets crowded boards when combat stalls.
- 1 Desolation of Smaug | wipe | A thematic board reset for unfavorable positions.
- 1 Glóin the Mighty // Easy Pickings | wipe | Offers a creature early and a sweeping effect when needed.
- 1 Last Light of Durin's Day | synergy | Supports the deck's Treasure, artifact, and combat plan.
- 1 Andúril, Flame of the West | synergy | Legendary equipment that rewards committing to combat.
- 1 Andúril, Narsil Reforged | synergy | A durable combat enhancement for valuable attackers.
- 1 Glamdring | synergy | Equipment support for the creature-heavy strategy.
- 1 Long-Lost Lances | synergy | Boosts attackers and complements the equipment package.
- 1 Sting, Bilbo's Sword | synergy | Low-cost equipment that improves combat pressure.
- 1 Well-Worn Spatula | synergy | Additional equipment support for equipped attackers.
- 1 Smaug's Fury | synergy | A thematic payoff for the deck's aggressive resource plan.
- 1 Dwarven Mauler | synergy | A sturdy body that benefits from the combat-focused shell.
- 1 Dwarven Warriors | synergy | Adds another creature that supports combat development.
- 1 Bilbo's Ring | interaction | Protects important creatures and disrupts opposing plans.
- 1 Dwarven Mattock | interaction | Versatile equipment-based interaction.
- 1 Mithril Coat | interaction | Protects the commander or a key threat from removal.
- 1 Getaway Barrel | interaction | Provides utility against hostile board states.
- 1 Giant's Boulder | interaction | A flexible artifact answer with tactical utility.
- 1 The Black Arrow | interaction | Equipment-based answer to opposing creatures.
- 1 Bothersome Noisemaker | interaction | Disruptive creature utility for contested boards.
- 1 Smaug, the Great Calamity // Spew Flame | interaction | A Dragon threat with an attached interactive spell.
- 1 Desert Were-Worm | threat | A large Dragon-themed finisher.
- 1 Oliphaunt | threat | A sizeable battlefield presence for closing games.
- 1 Olog-hai Crusher | threat | A high-impact combat threat.
- 1 Orcish Siegemaster | threat | Pressures opponents through combat.
- 1 Stone-Giant of High Pass | threat | A powerful creature that contributes to board control.
- 1 Troop of Ponies | threat | Adds a wide, aggressive creature presence.
- 1 Dáin Ironfoot | threat | A legendary creature that strengthens the attacking board.
- 1 Thorin, Company's Leader | threat | A resilient legendary threat with useful utility.
- 1 Thorin, Mountain-king | threat | A strong legendary body for the midgame.
- 1 Gundabad Opportunist | threat | Provides pressure while advancing creature presence.
- 1 Guttersnipe | threat | Converts the deck's spells into additional damage.
- 1 Iron Hills Stalwart | threat | A reliable creature for equipment and combat.
- 1 Bombur, Gentle Dreamer | other | Flavorful utility creature for the supporting cast.
- 1 Goblin-town Flunkies | other | Early board presence and combat support.
- 1 Misty Mountains Raider | other | Additional creature density for the aggressive plan.
- 31 Mountain | land | Primary red mana source.
- 1 Dragon-Cursed Halls | land | Thematic nonbasic mana source.
- 1 Hobbit Hole | land | Utility land that supports the deck's resource plan.
- 1 Rogue's Passage | land | Helps a large attacker break through stalled boards.
- 1 The Lonely Mountain | land | Thematic red-producing land.
- 1 Treasure Vault | land | Land slot with late-game Treasure utility.

</details>

### 21. the Hobbit family, mana from outside

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 156 names.

Commander: Smaug the Impenetrable.

Grade: bad, score 0.03, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band, profile_off_band, profile_off_band. Block findings: 0.

Cost: $846.20 to buy, $846.20 the whole deck.

**Summary:** A Rakdos Middle-earth battle deck centered on immense Dragons, Goblins, and legendary villains. It develops mana quickly, uses equipment to turn its creatures into dangerous attackers, keeps opposing boards in check with plentiful answers, and closes games through repeated combat pressure backed by dramatic sweepers and card advantage.

The quality model grades this deck below the precon baseline against the top lists of the format: more opening hands hold two to four lands than the norm, and that raises the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. The deck holds no fast mana, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The list hangs together loosely as a Rakdos midrange beatdown pile with removal and equipment, but the pieces pull several ways at once — a light equipment/Andúril subtheme, a thin Goblin subtheme, only two or three actual Dragons to support the "immense Dragons" angle, and three sweepers that punish its own creature-heavy board.
- PLAN theme_fit=partly: It is a Rakdos Middle-earth deck led by the requested commander at roughly bracket-3 trappings, but the request restricted nonland, non-rock cards to the Hobbit set and the deck is stuffed with LTR and unrelated-set cards (The One Ring, Orcish Bowmasters, Sauron the Lidless Eye, Witch-king of Angmar, Andúril, Palantír, plus Languish, Night's Whisper, Guttersnipe, Inferno Titan).
- PLAN useful_as_built=yes: Thirty-seven lands in two colors with fetches, a dual, Command Tower and Signet-style rocks, plus a reasonable curve of ramp, draw, removal and a dozen finishers, means the deck casts its spells and has multiple ways to close a game.
- PLAN summary_honest=partly: placeholder
- [INFO] `curve_summary`: average mana value 3.24 over 63 nonland cards
- [WARN] `outside_requested_set`: the sets you named do not hold 12 cards: Ancient Tomb, Arid Mesa, Blood Crypt, Bloodstained Mire, Command Tower, Dragonskull Summit, Exotic Orchard, Marsh Flats, Polluted Delta, Scalding Tarn, Verdant Catacombs, Wooded Foothills

<details><summary>The deck list</summary>

- 14 Mountain | land | Red basic mana for early development and heavy red spells.
- 10 Swamp | land | Black basic mana for reliable removal and draw spells.
- 1 Ancient Tomb | land | Fast colorless acceleration for expensive threats.
- 1 Arid Mesa | land | Fetches a Mountain or Blood Crypt while thinning the mana base.
- 1 Blood Crypt | land | Untapped dual source when needed.
- 1 Bloodstained Mire | land | Fetches either primary basic color or Blood Crypt.
- 1 Command Tower | land | Reliable two-color fixing.
- 1 Dragonskull Summit | land | Dual-color source with minimal entry delay.
- 1 Exotic Orchard | land | Usually supplies either deck color in multiplayer.
- 1 Marsh Flats | land | Fetches Swamps or Blood Crypt.
- 1 Polluted Delta | land | Fetches Swamps or Blood Crypt.
- 1 Scalding Tarn | land | Fetches Mountains or Blood Crypt.
- 1 Verdant Catacombs | land | Fetches Swamps or Blood Crypt.
- 1 Wooded Foothills | land | Fetches Mountains or Blood Crypt.
- 1 Arcane Signet | ramp | Efficient color fixing and acceleration.
- 1 Mox Amber | ramp | Low-cost acceleration alongside the deck's legendary creatures.
- 1 Wayfarer's Bauble | ramp | Early land-based acceleration that fixes mana.
- 1 Bag End Banquet | ramp | Early mana development for the deck's larger spells.
- 1 Burn, Burn, Tree and Fern | ramp | Incremental mana acceleration that supports the top end.
- 1 Dragon's Desire | ramp | Treasure-focused acceleration for Dragons and large turns.
- 1 Bolg's Company | ramp | Creature-based acceleration that advances the board.
- 1 Fíli and Kíli, Joyous | ramp | Legendary ramp piece that also enables Mox Amber.
- 1 Long-Bodied Grey Dog | ramp | Low-curve ramp creature for smoother early turns.
- 1 Troop of Ponies | ramp | Additional creature-based mana development.
- 1 Balin, Loremaster | draw | Legendary card advantage engine.
- 1 Gollum, Riddle Master | draw | Creature-based card selection and advantage.
- 1 Night's Whisper | draw | Efficient early card draw.
- 1 Palantír of Orthanc | draw | Persistent card advantage with late-game pressure.
- 1 Rage into the Valley | draw | Refills resources while maintaining the deck's aggressive plan.
- 1 Ragged Short Spear | draw | Equipment-based card advantage for attacking creatures.
- 1 Reverent Howl | draw | Instant-speed card draw.
- 1 The Master of Lake-town | draw | Legendary source of repeatable card advantage.
- 1 The Sackville-Bagginses | draw | Creature-based value engine.
- 1 Thrór's Map | draw | Artifact card advantage that works with the equipment package.
- 1 Óin the Brave | draw | Legendary draw source that supports a creature-heavy board.
- 1 Bilbo's Ring | interaction | Protects key creatures and disrupts opposing combat plans.
- 1 Dwarven Mattock | interaction | Flexible equipment interaction for contested boards.
- 1 Mithril Coat | interaction | Protects the deck's important legendary threats.
- 1 My Precious // Allure of Power | interaction | Versatile protection and disruption from one card.
- 1 The One Ring | interaction | Provides a protected window to stabilize and rebuild.
- 1 Gathering of Darkness | interaction | Disruptive spell that pressures opposing plans.
- 1 Getaway Barrel | interaction | Flexible artifact utility against opposing permanents.
- 1 Smaug's Fury | interaction | Instant-speed combat and board-control utility.
- 1 Azog, Moria's Ruin | removal | Legendary removal attached to a relevant body.
- 1 Battle-Scarred Goblin | removal | Creature-based removal that contributes to board presence.
- 1 Bilbo's Deadly Slice | removal | Efficient targeted removal.
- 1 Bitter Downfall | removal | Flexible answer to problematic permanents.
- 1 Fire of Orthanc | removal | Direct removal for creatures and other threats.
- 1 Goblin Cratermaker | removal | Low-cost answer to creatures and artifacts.
- 1 Improvised Club | removal | Cheap instant-speed removal.
- 1 Orcish Bowmasters | removal | Efficient creature control with pressure.
- 1 Stir Up Trouble | removal | Additional targeted answer that keeps the board clear.
- 1 Call Forth the Tempest | wipe | Broad reset for boards that get out of hand.
- 1 Desolation of Smaug | wipe | Thematic mass removal that clears the way for large threats.
- 1 Languish | wipe | Efficient sweeper against creature-heavy boards.
- 1 Desert Were-Worm | threat | Large Dragon threat that pressures life totals.
- 1 Gandalf, Goblins' Bane // Flameshape | threat | Flexible legendary threat with a useful adventure.
- 1 Great Goblin, Foul-Hearted | threat | A resilient legendary attacker for the Goblin element.
- 1 Guttersnipe | threat | Converts the deck's spells into repeated damage.
- 1 Inferno Titan | threat | High-impact finisher that controls smaller creatures.
- 1 Sauron, the Lidless Eye | threat | Powerful legendary threat for prolonged games.
- 1 Smaug, the Great Calamity // Spew Flame | threat | Dragon finisher with removal flexibility.
- 1 Stone-Giant of High Pass | threat | Large red threat that also affects the board.
- 1 The Great Goblin | threat | Legendary Goblin threat that supports the deck's creature plan.
- 1 Troll of Khazad-dûm | threat | Large black attacker with early-game utility.
- 1 Witch-king of Angmar | threat | Evasive legendary finisher with removal pressure.
- 1 Witch-king, Bringer of Ruin | threat | A second evasive finisher that punishes vulnerable boards.
- 1 Along the Crooked Way | synergy | Thematic value piece supporting the deck's dark Middle-earth plan.
- 1 Andúril, Flame of the West | synergy | Equipment that turns legendary attackers into stronger finishers.
- 1 Andúril, Narsil Reforged | synergy | A second equipment payoff for the attacking creature package.
- 1 Down, Down to Goblin-town | synergy | Supports the Goblin subtheme and ongoing board development.
- 1 Goblin Plate Mail | synergy | Equipment payoff for Goblins and other attackers.
- 1 Last Light of Durin's Day | synergy | Thematic enchantment value for the creature plan.
- 1 Long-Lost Lances | synergy | Equipment that improves combat pressure from large creatures.
- 1 Sting, Bilbo's Sword | synergy | Efficient equipment for enabling profitable attacks.
- 1 Supper for Spiders | synergy | Thematic utility that supports the deck's attrition plan.
- 1 Well-Worn Spatula | synergy | Additional equipment support for combat-focused threats.

</details>

### 22. a set family and a card from outside it

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 175 names.

Commander: Thranduil, the Elvenking.

Grade: bad, score 0.06, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band; then profile_off_band. Block findings: 0.

Cost: $523.19 to buy, $523.19 the whole deck.

**Summary:** A Sultai Elf-focused Hobbit deck that develops a broad mana base, builds a resilient board of Elves and legends, and supports its pressure with steady card advantage. It balances targeted answers, protective disruption, and sweeping resets while maintaining a strong late-game threat package.

The quality model grades this deck below the precon baseline against the top lists of the format: more opening hands hold two to four lands than the norm, and that raises the grade. Many of the cards appear in no top list, and that lowers the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: There is a loose Elf/legend go-wide core with card advantage, but three symmetrical sweepers plus a pile of unrelated one-off legends and utility creatures pull against the wide Elf board the deck is trying to assemble.
- PLAN theme_fit=partly: It is a Commander deck led by the requested Thranduil at a bracket-3-legal power level with Sol Ring intact, but a noticeable share of the list (The One Ring, Orcish Bowmasters, Delighted Halfling, Lórien Revealed, Mithril Coat, Languish, Colossal Whale) comes from outside the Hobbit set the person asked to build from.
- PLAN useful_as_built=partly: 36 lands and a reasonable curve make it playable, but with only 7 Forests and a couple of green-producing nonbasics the mana badly underserves the green half of the deck (Elvish Mystic, Elvish Archdruid, Radagast, Last March of the Ents, and the many green Elves), so green spells will regularly be stranded.
- PLAN summary_honest=yes: The claims — Sultai colors, an Elf/legend body count, roughly a dozen draw pieces, targeted removal, counter/protection effects, and three wipes — are all actually present in the list, and it does not assert combos or engines the cards lack.
- [INFO] `curve_summary`: average mana value 3.14 over 63 nonland cards
- [WARN] `profile_off_band`: the worst color's share of the sources it needs is 0.79, and bracket 3 wants 0.85 or more (sources of requirement: U 18.5 of 19, B 20.5 of 26, G 17.75 of 20)
- [WARN] `outside_requested_set`: the sets you named do not hold 1 card: Sol Ring

<details><summary>The deck list</summary>

- 10 Island | land | Blue mana base
- 11 Swamp | land | Black mana base
- 7 Forest | land | Green mana base
- 1 Dragon-Cursed Halls | land | Colored fixing land
- 1 Elven Passage | land | Elf-themed colored land
- 1 Elvenking's Halls | land | Colored fixing land
- 1 Hobbit Hole | land | Colored fixing land
- 1 Minas Morgul, Dark Fortress | land | Black source
- 1 Mirkwood | land | Elf-themed colored land
- 1 Rivendell | land | Blue source
- 1 The Black Gate | land | Black source
- 1 Elvish Visionary | draw | Efficient Elf cantrip
- 1 Fateful Discovery | draw | Persistent card advantage
- 1 Gandalf, Shadow's Foe | draw | Card-advantage threat
- 1 Hithlain Knots | draw | Instant-speed card selection
- 1 Lórien Revealed | draw | Early fixing or late card draw
- 1 Night's Whisper | draw | Efficient card draw
- 1 Plunder the Trollshaws | draw | Flexible instant draw
- 1 Reverent Howl | draw | Creature-based card advantage
- 1 Thrór's Map | draw | Repeatable card advantage
- 1 Uncover the Moon-Letters | draw | Enchanting card advantage
- 1 Last March of the Ents | draw | Creature-driven refill
- 1 Sol Ring | ramp | Requested efficient mana acceleration
- 1 Arcane Signet | ramp | Reliable color fixing
- 1 Mox Amber | ramp | Legendary mana acceleration
- 1 Necklace of Girion | ramp | Color fixing and acceleration
- 1 Delighted Halfling | ramp | Fixing for the commander and legends
- 1 Elvish Mystic | ramp | Early Elf mana creature
- 1 Elvish Archdruid | ramp | Elf-scaled mana production
- 1 Woodland Weavemaster | ramp | Elf mana development
- 1 Pelargir Survivor | ramp | Mana development
- 1 Radagast of Rhosgobel | ramp | Mana acceleration on a durable body
- 1 Bitter Downfall | removal | Versatile permanent removal
- 1 Bilbo's Deadly Slice | removal | Efficient creature removal
- 1 Enchanted River's Grasp | removal | Permanent-based answer
- 1 Giant's Boulder | removal | Repeatable removal tool
- 1 Merciless Executioner | removal | Edict creature removal
- 1 Orcish Bowmasters | removal | Creature control with board presence
- 1 Quarrel | removal | Flexible instant answer
- 1 Stir Up Trouble | removal | Creature removal
- 1 Uneasy Partings | removal | Instant-speed removal
- 1 Gnashing of Teeth | wipe | Mass creature control
- 1 Languish | wipe | Efficient creature sweeper
- 1 Raise the Palisade | wipe | Tribal-friendly board reset
- 1 Confusticate and Bebother | interaction | Stack interaction
- 1 Elrond, Moon-Reader | interaction | Interactive legendary Elf
- 1 Bilbo's Ring | interaction | Protective utility
- 1 Mithril Coat | interaction | Commander protection
- 1 My Precious // Allure of Power | interaction | Flexible protection and disruption
- 1 Old Fat Spider Can't See Me | interaction | Defensive disruption
- 1 Stern Scolding | interaction | Low-cost countermagic
- 1 The One Ring | interaction | Protective utility
- 1 Boughside Wanderers | synergy | Elf tribal body
- 1 Cantankerous Keepers | synergy | Elf tribal body
- 1 Elven Raft-Steerer | synergy | Elf tribal support
- 1 Elvenking's Harper | synergy | Elf tribal support
- 1 Galadhrim Guide | synergy | Elf tribal support
- 1 Grey Havens Navigator | synergy | Elf tribal support
- 1 Lothlórien Lookout | synergy | Elf tribal support
- 1 Mirkwood Meditator | synergy | Elf tribal support
- 1 Mirkwood Nurturer | synergy | Elf tribal support
- 1 Nimrodel Watcher | synergy | Elf tribal support
- 1 Celeborn the Wise | threat | Legendary Elf payoff
- 1 Colossal Whale | threat | Large blue finisher
- 1 Galion, Elvenking's Butler | threat | Legendary Elf threat
- 1 Gandalf, Wandering Wizard | threat | High-impact legendary threat
- 1 Gollum, Riddle Master | threat | Evasive value threat
- 1 Guardian of the Halls | threat | Elf board presence
- 1 Haunt of the Dead Marshes | threat | Black creature threat
- 1 Mirkwood Pathmaker | threat | Elf combat threat
- 1 The Master of Lake-town | threat | Value-oriented threat
- 1 Thranduil, Sindarin Liege // Silvan Rally | threat | Elf-themed legendary payoff
- 1 Willow-Wind | threat | Large creature threat
- 1 Witch-king of Angmar | threat | Evasive black finisher

</details>

### 23. two set families at once

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 203 names.

Commander: Kíli the Resourceful.

Grade: typical, score 0.45, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $292.87 to buy, $292.87 the whole deck.

**Summary:** Kíli leads a white creature-and-artifact shell built around Dwarves, Bloomburrow creatures, and a substantial Equipment package. The deck develops its board with artifact mana and steady card flow, uses removal and protective interaction to maintain momentum, and wins by converting its creature board and threat cards into sustained pressure. It gives up a more explosive, spell-heavy approach for a thematic permanent-based plan that is strongest when several creatures and artifacts are established.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: The list is a consistent mono-white creatures-plus-equipment/artifact board deck — cheap bodies, token/artifact makers, Skullclamp and Mentor draw, Maskwood Nexus to glue the typal bits, and equipment to push damage — with only mild tension from the three sweepers, which are mostly asymmetric or scalable.
- PLAN theme_fit=partly: The commander and a good chunk of the deck do come from the Hobbit and Bloomburrow sets and the power level is plausibly bracket 3, but roughly half the list is outside the two named sets (Skullclamp, Sol Ring, Swords to Plowshares, Skyclave Apparition, Karn, Academy Manufactor, plus a pile of LTR cards), so the one limit the person named is only half honored.
- PLAN useful_as_built=yes: 36 lands in mono-white with ten ramp pieces, a low curve, plenty of card draw, real removal and multiple equipment-based damage routes means it casts its spells and can close a game as it stands.
- PLAN summary_honest=partly: The white creature-artifact-equipment description matches the cards, but it oversells the 'Dwarves and Bloomburrow' theming when much of the list is generic staples and LTR (not Hobbit) cards, and it never mentions the three board wipes or the near-dead Karn (no sideboard/wishboard support).
- [INFO] `curve_summary`: average mana value 2.70 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Command Tower | land | Reliable colored land for the mana base.
- 1 Exotic Orchard | land | Flexible colored land for the mana base.
- 1 Fountainport | land | Utility land slot in the mana base.
- 1 Minas Tirith | land | Thematic white land for the mana base.
- 1 Reliquary Tower | land | Utility land slot in the mana base.
- 1 Rogue's Passage | land | Utility land slot in the mana base.
- 1 Treasure Vault | land | Artifact land that supports the deck's artifact theme.
- 29 Plains | land | Consistent basic white mana for the deck.
- 1 Caretaker's Talent | draw | Dedicated card-draw piece from the shortlist.
- 1 Circuit Mender | draw | Artifact creature that supplies card draw.
- 1 Cut a Deal | draw | Dedicated card-draw spell from the shortlist.
- 1 Errand-Rider of Gondor | draw | Creature-based card draw for the board-focused plan.
- 1 Fountainport Bell | draw | Artifact card-draw option.
- 1 Heirloom Epic | draw | Thematic artifact card-draw piece.
- 1 Idol of Oblivion | draw | Low-cost artifact card draw.
- 1 Inspiring Overseer | draw | Creature-based card draw.
- 1 Jolly Gerbils | draw | Bloomburrow creature that provides card draw.
- 1 Mentor of the Meek | draw | Repeatable creature-focused card draw.
- 1 Skullclamp | draw | Efficient artifact card-draw tool.
- 1 Bofur, Reliable Guardian // Concerted Care | interaction | Thematic Dwarf interaction with an Adventure option.
- 1 Bilbo's Gambit | interaction | Instant-speed interaction.
- 1 Dawn's Truce | interaction | Instant interaction for protecting the board plan.
- 1 Galadriel's Dismissal | interaction | Flexible instant-speed interaction.
- 1 Luminous Broodmoth | interaction | Creature-based interaction for the board-focused shell.
- 1 Mithril Coat | interaction | Equipment-based interaction for key permanents.
- 1 Parting Gust | interaction | Instant-speed interaction.
- 1 Selfless Spirit | interaction | Creature-based interaction that supports the board.
- 1 Arcane Signet | ramp | Efficient artifact mana acceleration.
- 1 Bag End Banquet | ramp | Thematic artifact ramp.
- 1 Burnished Hart | ramp | Creature-based mana acceleration.
- 1 Fellwar Stone | ramp | Low-cost artifact mana acceleration.
- 1 Mind Stone | ramp | Low-cost artifact mana acceleration.
- 1 Mox Amber | ramp | Fast legendary-themed mana acceleration.
- 1 Ornithopter of Paradise | ramp | Artifact creature that accelerates mana.
- 1 Patchwork Banner | ramp | Artifact ramp that fits the creature-focused shell.
- 1 Sol Ring | ramp | Efficient artifact mana acceleration.
- 1 Wayfarer's Bauble | ramp | Reliable early mana acceleration.
- 1 Angel of the Ruins | removal | High-impact artifact creature removal.
- 1 Banishing Light | removal | Versatile permanent removal.
- 1 Fiend Hunter | removal | Creature-based removal.
- 1 Generous Gift | removal | Flexible instant-speed removal.
- 1 Loran of the Third Path | removal | Thematic creature-based removal.
- 1 Nettle Guard | removal | Bloomburrow creature removal.
- 1 Skyclave Apparition | removal | Efficient creature-based removal.
- 1 Swords to Plowshares | removal | Efficient instant-speed removal.
- 1 Westfold Rider | removal | Thematic creature-based removal.
- 1 Dusk // Dawn | wipe | Flexible board-clearing spell.
- 1 Martial Coup | wipe | Board wipe for resetting crowded tables.
- 1 Promise of Loyalty | wipe | Board wipe that disrupts opposing creature boards.
- 1 Blade Splicer | synergy | Artifact-creature synergy piece.
- 1 Carrot Cake | synergy | Thematic Food synergy piece.
- 1 Dáin, Lord of the Iron Hills | synergy | Dwarf synergy for Kíli's themed shell.
- 1 Iron Hills Blacksmith | synergy | Dwarf Artificer synergy.
- 1 Ori, Keeper of Songs | synergy | Thematic Dwarf synergy piece.
- 1 Tangle Tumbler | synergy | Artifact synergy for the deck's permanent core.
- 1 Fíli the Pathfinder | threat | Legendary Dwarf threat for the themed creature plan.
- 1 Karn, the Great Creator | threat | Standalone threat from the shortlist.
- 1 Academy Manufactor | other | Artifact creature that reinforces the deck's artifact core.
- 1 Andúril, Flame of the West | other | Thematic legendary Equipment for the creature shell.
- 1 Bilbo, Unexpected Adventurer | other | Thematic Hobbit legendary creature.
- 1 Dwarven Provisioner | other | Additional thematic Dwarf creature.
- 1 Dwarven Shortsword | other | Low-cost thematic Equipment.
- 1 Glamdring | other | Thematic legendary Equipment.
- 1 Helm of the Host | other | Artifact Equipment for the permanent-focused plan.
- 1 Maskwood Nexus | other | Artifact support for the creature-theme shell.
- 1 Mouse Trapper | other | Bloomburrow creature for the board plan.
- 1 Sting, Bilbo's Sword | other | Thematic legendary Equipment.
- 1 Sword of the Squeak | other | Bloomburrow Equipment for the creature shell.
- 1 Three Tree Mascot | other | Artifact creature that fits the themed permanent core.
- 1 Well-Worn Spatula | other | Thematic artifact Equipment.
- 1 Whiskervale Forerunner | other | Bloomburrow creature for the board-focused plan.

</details>

### 24. a 60-card deck from one set

Format: Modern. Theme: aggro. Pool: any_card. Shortlist: 96 names.

Grade: baseline, score 0.33, model 20260903T210658Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $23.40 to buy, $23.40 the whole deck.

**Summary:** This mono-red Bloomburrow aggro deck deploys a steady stream of creature threats, uses its synergy creatures to strengthen its combat pressure, and clears resistance with direct removal. It wins by maintaining attacks through the early game and following up with larger threats such as Dragonhawk, Fate's Tempest when needed. The deck gives up broad answers and defensive flexibility in favor of a focused red mana base, proactive creatures, and a straightforward attacking plan.

The quality model grades this deck at the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. The color sources cover the pips, and that raises the grade.

- PLAN plan_coherent=yes: Every slot serves one low-curve aggro plan: cheap red creatures, a couple of pump/card-flow spells to keep pressure, burn to clear blockers, and Dragonhawk as a top-end finisher.
- PLAN theme_fit=yes: It is mono-red, 60 cards, Modern-legal, and drawn essentially entirely from Bloomburrow's red pool with a casual aggro shell, exactly as asked.
- PLAN useful_as_built=yes: 24 Mountains cast every spell and the curve tops at Dragonhawk, so it plays out fine as-is, though 24 lands is a couple more than a deck this cheap really wants.
- PLAN summary_honest=yes: It accurately describes the creature-first attack plan, the burn as removal, the all-Mountain base, and openly concedes the lack of interaction and defensive tools.
- [INFO] `curve_summary`: average mana value 2.94 over 36 nonland cards

<details><summary>The deck list</summary>

- 24 Mountain | land | Provides the consistent red mana base for a mono-red deck.
- 4 Might of the Meek | draw | Supplies card flow while fitting the low-cost aggressive plan.
- 2 Sazacap's Brew | draw | Adds further card flow without diluting the red mana base.
- 2 Brazen Collector | ramp | Provides the requested mana acceleration on a creature body.
- 4 Blooming Blast | removal | Efficiently clears opposing pieces that interfere with attacks.
- 2 Rabid Gnaw | removal | Rounds out the removal package while staying in the deck's color.
- 4 Emberheart Challenger | synergy | A key aggressive synergy creature for the deck's combat plan.
- 4 Hearthborn Battler | synergy | Reinforces the deck's synergy package and adds pressure.
- 4 Frilled Sparkshooter | threat | A creature threat that helps keep pressure on the opponent.
- 4 Reptilian Recruiter | threat | A core attacking threat for the early and middle turns.
- 4 Teapot Slinger | threat | Adds another reliable creature threat to the aggressive curve.
- 2 Dragonhawk, Fate's Tempest | threat | Provides a larger top-end threat when early pressure is answered.

</details>

### 25. use no card of an owned precon

Format: Commander. Theme: Avengers superheroes. Pool: owned_first. Shortlist: 152 names.

Commander: Syr Konrad, the Grim.

Grade: typical, score 0.46, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band. Block findings: 0.

Cost: $0.00 to buy, $226.15 the whole deck.

**Summary:** A mono-black Avengers-inspired roster built around grim battlefield attrition, disposable allies, sinister villains, and superhero technology. It develops resources through equipment and artifacts, keeps opposing boards under pressure, and closes with a procession of imposing legendary and demonic threats.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: There is a real spine — Syr Konrad plus a -1/-1 counter attrition package (Blowfly Infestation, Soul Snuffers, Midnight Banshee, Carnifex Demon, Grim Poppet, Black Sun's Zenith) that turns board sweeps into drain — but the equipment/ramp package and the pile of unrelated big legends pull loosely in another direction, and the summary never names the actual engine.
- PLAN theme_fit=no: The request was an Avengers superhero deck, and this is a mono-black villains-and-Middle-earth pile led by Syr Konrad with only a handful of Marvel-flavored cards (Avengers Tower, Baxter Building, M.O.D.O.K., Baron Strucker, Stolen Stark Tech), so it answers a different request.
- PLAN useful_as_built=yes: 36 lands with solid fixing for a mono-black deck, ten ramp pieces, eleven draw sources, twelve removal spells and three sweepers, and clear win paths through Konrad drain plus a dozen sizable threats — it plays fine out of the box.
- PLAN summary_honest=partly: placeholder
- [WARN] `not_owned`: Syr Konrad, the Grim: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.02 over 63 nonland cards

<details><summary>The deck list</summary>

- 24 Swamp | land | Reliable black mana base.
- 1 Avengers Tower | land | Thematic headquarters land.
- 1 Baxter Building | land | Thematic city landmark.
- 1 Big Apple, 3 a.m. | land | Thematic urban setting.
- 1 Castle Locthwain | land | Black source with card advantage utility.
- 1 Command Tower | land | Reliable commander mana.
- 1 Evolving Wilds | land | Finds a Swamp and supports graveyard texture.
- 1 Fabled Passage | land | Fetches a Swamp while filling the graveyard.
- 1 Ifnir Deadlands | land | Black source with -1/-1 counter utility.
- 1 Marsh Flats | land | Fetches the primary basic land.
- 1 Minas Morgul, Dark Fortress | land | Flavorful utility land.
- 1 Secluded Courtyard | land | Supports the creature-heavy roster.
- 1 Takenuma, Abandoned Mire | land | Black source with recursion utility.
- 1 Arcane Signet | ramp | Reliable mana acceleration.
- 1 Bender's Waterskin | ramp | Artifact-based mana support.
- 1 Coin of Mastery | ramp | Durable midgame mana acceleration.
- 1 Commander's Sphere | ramp | Mana fixing that can become a card.
- 1 Sol Ring | ramp | Efficient artifact acceleration.
- 1 Springleaf Drum | ramp | Turns spare creatures into mana.
- 1 Sword of the Animist | ramp | Equipment that develops the mana base.
- 1 Thought Vessel | ramp | Artifact acceleration with hand support.
- 1 Wayfarer's Bauble | ramp | Finds a basic land.
- 1 Wizard's Rockets | ramp | Early mana support with later utility.
- 1 Baron Strucker, HYDRA Overlord | draw | Thematic card advantage engine.
- 1 Call of the Ring | draw | Steady source of cards.
- 1 Dusk Urchins | draw | Turns counters and death into cards.
- 1 Grave Venerations | draw | Rewards the graveyard-focused plan.
- 1 Hoarder's Greed | draw | Flexible burst of card advantage.
- 1 Idol of Oblivion | draw | Efficient artifact card advantage.
- 1 Lembas | draw | Low-cost card selection and recursion value.
- 1 Mask of Memory | draw | Combat-based card filtering.
- 1 Night's Whisper | draw | Efficient card advantage.
- 1 Painful Truths | draw | Efficient refill spell.
- 1 Skullclamp | draw | Converts disposable creatures into cards.
- 1 Champion's Helm | interaction | Protects the commander or a major threat.
- 1 Darksteel Plate | interaction | Durable protection for key creatures.
- 1 Lightning Greaves | interaction | Immediate protection and mobility.
- 1 Orcish Medicine | interaction | Flexible defensive utility.
- 1 Stone of Erech | interaction | Disrupts opposing graveyard plans.
- 1 Stolen Stark Tech | interaction | Thematic protection equipment.
- 1 Swiftfoot Boots | interaction | Protects key creatures.
- 1 The Walls of Ba Sing Se | interaction | Defensive artifact utility.
- 1 Bitter Downfall | removal | Answers problematic creatures or planeswalkers.
- 1 Bitter Triumph | removal | Flexible instant-speed answer.
- 1 Claim the Precious | removal | Clean answer to a major threat.
- 1 Dismember | removal | Efficient creature removal.
- 1 Fatal Push | removal | Low-cost creature interaction.
- 1 Heartless Act | removal | Flexible creature removal.
- 1 Infernal Grasp | removal | Unconditional creature answer.
- 1 Nameless Inversion | removal | Tribal-compatible creature removal.
- 1 Orcish Bowmasters | removal | Creature-based damage and removal pressure.
- 1 Black Sun's Zenith | wipe | Scalable battlefield reset that supports counter themes.
- 1 Extinction Event | wipe | Efficient selective battlefield reset.
- 1 Game Over | wipe | Broad reset for difficult boards.
- 1 Blowfly Infestation | synergy | Moves counters through dying creatures for repeated value.
- 1 Contagion Clasp | synergy | Places and multiplies useful counters.
- 1 Deadly Dispute | synergy | Turns expendable permanents into resources.
- 1 Grim Poppet | synergy | Counter manipulation pairs with creature attrition.
- 1 Mordor Muster | synergy | Builds a disposable creature force.
- 1 Nasty End | synergy | Rewards sacrificing a creature at the right moment.
- 1 Obsessive Pursuit | synergy | Supports the deck's attrition plan.
- 1 Oft-Nabbed Goat | synergy | Provides value when it leaves the battlefield.
- 1 Soul Snuffers | synergy | Spreads counters across creature-heavy boards.
- 1 Undercity Dire Rat | synergy | Creature-based support for the graveyard plan.
- 1 Archfiend of Ifnir | threat | Dominating evasive finisher with counter pressure.
- 1 Carnifex Demon | threat | Large threat that controls smaller creatures.
- 1 M.O.D.O.K. | threat | Thematic villain and resilient payoff.
- 1 Massacre Girl, Known Killer | threat | Turns creature attrition into a threatening advantage engine.
- 1 Midnight Banshee | threat | Pressures opposing creature boards over time.
- 1 Rat King, Pale Piper | threat | Recurring creature payoff.
- 1 Ringwraiths | threat | Evasive pressure with removal utility.
- 1 Sinister Gnarlbark | threat | Creature threat that contributes to the value plan.
- 1 Skinrender | threat | Creature-based board pressure.
- 1 Voracious Fell Beast | threat | Large finishing body.
- 1 Witch-king of Angmar | threat | Evasive legendary finisher.
- 1 Zodiark, Umbral God | threat | Powerful top-end battlefield presence.

</details>

