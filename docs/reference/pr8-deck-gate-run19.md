# PR-8 deck gate

Run date: 2026-09-12. Card snapshot: 2026-09-04.

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
| Mean plan score, 0 to 1 | 0.70 |
| Plan reasons the judge left empty | 2 |
| Errors | 0 |
| Prompt version | 12 |
| Calls | 76 |
| Cost | $2.7504 |
| Time | 2153 seconds |

## Run

- Suite `decks`, run `pr8-deck-gate-run19`, on 2026-09-12, commit `5ef19ec`.
- Roles: generate on `gpt-5.6-terra` (openai, effort medium), judge on `claude-opus-5` (anthropic, effort medium), repair on `gpt-5.6-terra` (openai, effort medium).
- Versions: card snapshot 2026-09-04, generate prompt version 12, plan_rubric prompt version 2, precons `5.3.0+20260903`, quality_model `20260910T012734Z`.
- Calls: 76. Cost: $2.7504. Time: 2153 seconds.

## Findings by code

A block stops the deck. A warning and a note are reports.

| Code | Count |
|---|---|
| `curve_summary` | 25 |
| `not_owned` | 23 |
| `mana_pass` | 10 |
| `profile_off_band` | 9 |
| `basics_added` | 3 |
| `outside_requested_set` | 2 |
| `two_card_combo` | 1 |

By severity: BLOCK 0. WARN 35. INFO 38. 

## The set filter (PR-17B)

A deck asked for a set holds cards of that set family alone. Three things earn an exception, and each one is marked: a card the reader named (D-381), a mana card the reader allowed from outside (D-382), and a basic land, which no set limit filters (D-378).

| # | Prompt | Sets | In set | Fill | Deck cards outside | Marked |
|---|---|---|---|---|---|---|
| 19 | the Hobbit family, two colours | `hob,hoc` | 170 | 0 | 0 | 0 |
| 20 | the Hobbit family, a delegated commander | `hob,hoc` | 75 | 0 | 0 | 0 |
| 21 | the Hobbit family, mana from outside | `hob,hoc` | 127 | 26 | 18 | 18 |
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

Grade: bad, score 0.00, model 20260910T012734Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $454.13 to buy, $454.13 the whole deck.

**Summary:** Karlov leads a lifegain-focused permanent deck that develops quickly, builds a board through synergistic creatures and enchantments, and turns that board into sustained combat pressure. It wins by backing its growing threats with protection, removal, and timely sweepers, while draw engines help it continue deploying threats into the middle and late game. The tradeoff is a board-centered approach that relies on maintaining its synergy pieces rather than pursuing a highly explosive finish.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every card is a lifegain trigger, payoff, or life-based engine (Karlov, Ajani's Pridemate, Heliod, Archangel of Thune, Vito, Well of Lost Dreams), so the pieces point at one plan; the two wraths sit slightly against the creature-heavy board but both gain life and are only a small counterweight.
- PLAN theme_fit=yes: It is exactly the asked-for Karlov-led W/B lifegain Commander deck, and its power level (a few strong cards like Solitude and Serra Ascendant, but no infinite combos or fast mana) sits reasonably in bracket 3.
- PLAN useful_as_built=yes: 37 lands with dual/pain lands support the two-color curve, the creature count is high enough to keep the board stocked, and Vito, Archangel of Thune, Karlov, and the flying angels give several real ways to close a game, though the ramp package is slow and clunky.
- PLAN summary_honest=yes: Everything the summary claims — lifegain permanents, protection (Alseid, Faith's Shield, Metropolis Reformer), spot removal, sweepers, and draw engines — is actually present in the list, and it openly flags the deck's lack of an explosive finish.
- [INFO] `curve_summary`: average mana value 3.06 over 63 nonland cards
- [INFO] `mana_pass`: the builder moved 6 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 11 Plains | land | Primary white source for the deck’s white-heavy spells.
- 16 Swamp | land | Provides reliable black mana for Karlov and black support cards.
- 1 Caves of Koilos | land | Untapped dual-color source.
- 1 Command Tower | land | Reliable source for both deck colors.
- 1 Godless Shrine | land | Dual-color land for the core mana base.
- 1 Mana Confluence | land | Flexible untapped color fixing.
- 1 Vault of Champions | land | Dual-color source that supports early development.
- 1 City of Brass | land | Flexible source for either deck color.
- 1 Exotic Orchard | land | Additional flexible color fixing.
- 1 Isolated Chapel | land | Dual-color source for the mana base.
- 1 Shattered Sanctum | land | Dual-color source for the mana base.
- 1 Bounty Board | ramp | Low-cost acceleration for developing ahead of curve.
- 1 Hot Dog Cart | ramp | Early ramp that helps deploy the deck’s permanents.
- 1 Life of Toshiro Umezawa // Memory of Toshiro | ramp | Supports mana development while fitting the deck’s colors.
- 1 Misfortune Teller | ramp | Provides ramp to advance the board.
- 1 Nuka-Cola Vending Machine | ramp | Mana acceleration for the deck’s midgame.
- 1 Orazca Relic | ramp | Color-fixing ramp for the two-color plan.
- 1 Phial of Galadriel | ramp | Additional mana acceleration.
- 1 Spoils of Evil | ramp | Efficient ramp option for explosive follow-up turns.
- 1 Thieving Varmint | ramp | Low-cost ramp that improves early sequencing.
- 1 The Celestus | ramp | Reliable artifact ramp for sustained development.
- 1 Archivist of Oghma | draw | Efficient card advantage for keeping resources flowing.
- 1 Dawn of Hope | draw | Repeatable draw support for a lifegain-focused strategy.
- 1 Enduring Innocence | draw | Creature-based card advantage that supports the permanent plan.
- 1 Exemplar of Light | draw | A draw source that also contributes to board presence.
- 1 Mangara, the Diplomat | draw | Steady card advantage from a resilient permanent.
- 1 Markov Purifier | draw | Draw support that fits the deck’s Vampire-adjacent themes.
- 1 Restless Bloodseeker // Bloodsoaked Reveler | draw | Early card advantage that supports the deck’s creature plan.
- 1 Scheming Silvertongue // Sign in Blood | draw | Flexible draw option with an efficient adventure half.
- 1 Sigarda's Splendor | draw | A dedicated draw engine for the lifegain shell.
- 1 Survival Cache | draw | Efficient card advantage for maintaining momentum.
- 1 Well of Lost Dreams | draw | A strong draw outlet for the deck’s central theme.
- 1 Alseid of Life's Bounty | interaction | Low-cost protection for Karlov or another important permanent.
- 1 Courageous Resolve | interaction | Protective interaction for key board pieces.
- 1 Faith's Shield | interaction | Efficient protection to preserve threats and synergy pieces.
- 1 Distinguished Conjurer | interaction | Creature-based interaction that fits the deck’s plan.
- 1 Metropolis Reformer | interaction | Protective interaction attached to a board presence.
- 1 Restoration Magic | interaction | Flexible interaction for protecting the board.
- 1 Sword of Light and Shadow | interaction | Equipment-based interaction that supports creature pressure.
- 1 Werefox Bodyguard | interaction | Versatile creature interaction for disrupting opposing plans.
- 1 Ayli, Eternal Pilgrim | removal | Efficient repeatable removal on a low-cost creature.
- 1 Foolish Fate | removal | Direct removal for problematic opposing cards.
- 1 Gumdrop Poisoner // Tempt with Treats | removal | Flexible removal with an additional adventure option.
- 1 Murderous Rider // Swift End | removal | Versatile removal that remains useful as a creature.
- 1 Nightmare's Thirst | removal | Low-cost removal for controlling early threats.
- 1 Papalymo Totolymo | removal | Removal attached to a permanent for board presence.
- 1 Solitude | removal | Premium removal that can answer a major threat quickly.
- 1 Summon: Ixion | removal | Removal that contributes an additional permanent to the plan.
- 1 Umezawa's Jitte | removal | Reusable removal utility that rewards attacking creatures.
- 1 Ajani's Pridemate | synergy | A classic lifegain payoff that grows into meaningful pressure.
- 1 Angel of Vitality | synergy | Supports the deck’s lifegain-focused board plan.
- 1 Bloodthirsty Aerialist | synergy | Turns the central theme into an evasive growing threat.
- 1 Cleric Class | synergy | A low-cost permanent dedicated to lifegain synergy.
- 1 Essence Channeler | synergy | Supports the deck’s core synergy on an efficient creature.
- 1 Heliod, Sun-Crowned | synergy | A central synergy permanent for the lifegain strategy.
- 1 Ocelot Pride | synergy | Low-cost synergy piece that helps establish the board.
- 1 Resplendent Angel | synergy | A powerful creature payoff for the lifegain plan.
- 1 Serra Ascendant | synergy | Early pressure that benefits from the deck’s central theme.
- 1 Vito, Thorn of the Dusk Rose | synergy | A key lifegain payoff that adds reach to the deck.
- 1 Archangel of Thune | threat | High-impact threat that rewards the deck’s main plan.
- 1 Attended Healer | threat | A board-building threat that fits the lifegain shell.
- 1 Celestine, the Living Saint | threat | A resilient midgame threat for maintaining pressure.
- 1 Divinity of Pride | threat | Large flying threat suited to a lifegain deck.
- 1 Exalted Sunborn | threat | A top-end Angel threat for closing games.
- 1 Gisela, the Broken Blade | threat | Efficient legendary flying threat.
- 1 Liesa, Forgotten Archangel | threat | A substantial Angel threat for the deck’s upper curve.
- 1 Nykthos Paragon | threat | A powerful payoff threat for the lifegain strategy.
- 1 Rhox Faithmender | threat | A durable threat that supports the deck’s central plan.
- 1 Twinblade Paladin | threat | Creature pressure that belongs naturally in the lifegain shell.
- 1 Valkyrie Harbinger | threat | A large Angel threat and late-game payoff.
- 1 Vanguard Seraph | threat | Flying threat that strengthens the deck’s combat plan.
- 1 Amalia Benavides Aguirre | wipe | A low-cost wipe that can reset crowded boards.
- 1 Fumigate | wipe | A dependable full-board reset when behind.
- 1 Kaya's Wrath | wipe | A reliable sweeper for clearing opposing creature boards.

</details>

### 2. aristocrats Commander, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 179 names.

Commander: Denethor, Ruling Steward.

Grade: baseline, score 0.02, model 20260910T012734Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $151.43 the whole deck.

**Summary:** This is a resilient aristocrats deck that develops mana, establishes disposable creature and artifact resources, and uses its synergy pieces plus draw engines to keep generating value through a long game. It controls key opposing permanents, protects Denethor and important support creatures, and can reset crowded boards before rebuilding with its creature threats. Its tradeoff is that it leans on assembling several supporting permanents rather than presenting a single fast finish.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck draws more than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: As a white-black removal-and-protection value pile it hangs together, but the pieces pull against each other and against the stated plan: heavy equipment/protection and four board wipes sit next to a commander that wants a stream of expendable bodies, and there is almost nothing that makes fodder to sacrifice.
- PLAN theme_fit=no: The request was aristocrats, but beyond Denethor himself the list has essentially no sacrifice outlets, no token generators, and no death-trigger drain payoffs (no Blood Artist/Zulaport/Bastion of Remembrance effects), so what was built is a removal-and-protection control deck wearing an aristocrats commander.
- PLAN useful_as_built=partly: The mana base and curve are fine and it will cast its spells, but the win routes are very thin — a handful of medium creatures plus a Denethor drain that has almost no fodder to feed it — so games can stall out after the removal runs dry.
- PLAN summary_honest=no: placeholder
- [INFO] `curve_summary`: average mana value 2.57 over 63 nonland cards

<details><summary>The deck list</summary>

- 16 Plains | land | Provides reliable white mana for the deck’s spells.
- 16 Swamp | land | Provides reliable black mana for the deck’s spells.
- 1 Command Tower | land | Provides flexible commander-color mana.
- 1 City of Brass | land | Provides flexible colored mana.
- 1 Exotic Orchard | land | Provides flexible colored mana.
- 1 Plaza of Heroes | land | Supports the deck’s legendary commander mana needs.
- 1 Sol Ring | ramp | Provides an efficient early mana boost.
- 1 Arcane Signet | ramp | Provides reliable colored mana acceleration.
- 1 Fellwar Stone | ramp | Provides low-cost mana acceleration.
- 1 Wayfarer's Bauble | ramp | Helps develop mana early.
- 1 Thought Vessel | ramp | Provides mana acceleration.
- 1 Relic of Legends | ramp | Provides flexible mana acceleration.
- 1 Commander's Sphere | ramp | Provides colored mana acceleration.
- 1 Springleaf Drum | ramp | Turns a creature into early mana production.
- 1 Chromatic Lantern | ramp | Fixes colors while accelerating mana.
- 1 Deadly Dispute | ramp | Converts a disposable resource into mana development.
- 1 Skullclamp | draw | Supplies a powerful draw engine for expendable creatures.
- 1 Night's Whisper | draw | Provides efficient card flow.
- 1 Painful Truths | draw | Refills the hand with cards.
- 1 Nasty End | draw | Turns a creature resource into cards.
- 1 Grave Venerations | draw | Provides ongoing card advantage.
- 1 Idol of Oblivion | draw | Provides repeatable card advantage.
- 1 Lembas | draw | Provides a card while supporting the artifact package.
- 1 Mask of Memory | draw | Rewards creature combat with card selection.
- 1 Tome of Legends | draw | Builds steady card advantage around the commander.
- 1 Wall of Omens | draw | Provides an early blocker and a card.
- 1 Call of the Ring | draw | Provides a continuing draw source.
- 1 Buster Sword | draw | Adds an equipment-based draw source.
- 1 Foot Chopper | draw | Adds another artifact draw option.
- 1 Stone of Erech | draw | Provides card advantage from an inexpensive artifact.
- 1 Lightning Greaves | interaction | Protects an important creature immediately.
- 1 Swiftfoot Boots | interaction | Protects a key creature while keeping it active.
- 1 Boromir, Warden of the Tower | interaction | Provides disruptive protection for the board.
- 1 Bastion Protector | interaction | Helps keep the commander protected.
- 1 Clever Concealment | interaction | Protects permanents from opposing disruption.
- 1 Gift of Immortality | interaction | Helps preserve a valuable creature.
- 1 Reprieve | interaction | Offers flexible stack interaction.
- 1 Take Up the Shield | interaction | Protects a creature during combat or removal.
- 1 Champion's Helm | interaction | Provides commander-focused protection.
- 1 Darksteel Plate | interaction | Makes a key creature much harder to remove.
- 1 Sheltered by Ghosts | interaction | Protects a creature while adding interaction.
- 1 Unbreakable Formation | interaction | Protects the creature board from sweepers.
- 1 Bitter Triumph | removal | Provides flexible single-target removal.
- 1 Claim the Precious | removal | Answers a problematic opposing creature.
- 1 Crib Swap | removal | Provides creature removal at instant speed.
- 1 Fiend Hunter | removal | Adds creature-based removal to the board.
- 1 Generous Gift | removal | Answers a wide range of permanent types.
- 1 Get Lost | removal | Provides efficient permanent removal.
- 1 Infernal Grasp | removal | Provides unconditional creature removal.
- 1 Path to Exile | removal | Efficiently answers an opposing creature.
- 1 Swords to Plowshares | removal | Efficiently answers an opposing creature.
- 1 Angel of Serenity | removal | Provides a creature-based removal option.
- 1 Banishing Light | removal | Answers a problematic permanent.
- 1 Dispatch | removal | Provides low-cost creature removal.
- 1 Austere Command | wipe | Provides a flexible board reset.
- 1 Fumigate | wipe | Resets opposing creature boards.
- 1 One Ring to Rule Them All | wipe | Provides a sweeping reset option.
- 1 Black Sun's Zenith | wipe | Clears creature boards while scaling with available mana.
- 1 Al Bhed Salvagers | synergy | Supports the deck’s sacrifice-focused synergy plan.
- 1 Arcade Cabinet | synergy | Supports the deck’s broader synergy package.
- 1 Gollum the Abandoned | synergy | Adds a sacrifice-oriented synergy creature.
- 1 Gollum, Patient Plotter | synergy | Adds a recurring synergy piece for the deck’s plan.
- 1 Gríma Wormtongue | synergy | Supports the deck’s value-oriented synergy plan.
- 1 Heirloom Auntie | synergy | Adds a creature-based synergy piece.
- 1 Joo Dee, One of Many | synergy | Supports the deck’s creature synergy plan.
- 1 Nimble Hobbit | synergy | Adds an inexpensive creature synergy piece.
- 1 Phantom Train | synergy | Adds an artifact-based synergy piece.
- 1 Bill the Pony | threat | Provides a creature threat for closing games.
- 1 Hei Bai, Spirit of Balance | threat | Provides a substantial creature threat.

</details>

### 3. artifacts Commander, bracket 4

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 299 names.

Commander: Urza, Lord High Artificer.

Grade: baseline, score 0.01, model 20260910T012734Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $4301.36 to buy, $4301.36 the whole deck.

**Summary:** This focused blue artifact deck develops mana quickly, turns its artifact synergies into sustained card flow and pressure, and protects its board with a dense suite of interaction. It closes through its artifact creature threats while using targeted removal and selective resets to clear the way. The deck gives up broad color access and more varied answers in exchange for a streamlined artifact shell.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core artifact-ramp-and-cost-reduction shell hangs together, but several pieces pull elsewhere — Ghostly Flicker/Escape Protocol blink with almost no ETB payoffs, Inspiring Statuary's improvise has few nonartifact spells worth cheating, and the affinity/metalcraft cards (Frogmite, Thoughtcast, Mox Opal, Darksteel Juggernaut) sit alongside a low count of actual cheap permanent artifacts.
- PLAN theme_fit=partly: It is unmistakably a mono-blue Urza artifact deck as requested, but the requested bracket-4 high-power ceiling is not met — there is no Thopter-Sword, Power Artifact, or comparable engine, and the top end is Chrome Steed, Rust Elemental and Jhoira's Familiar.
- PLAN useful_as_built=partly: It casts its spells reliably and has real win routes (Urza's construct, Whirler Rogue evasion, Skysovereign, Darksteel Juggernaut), but 36 lands stacked on top of a dozen ramp pieces in a low-curve deck will flood often, and the beaters are slow enough that games can stall.
- PLAN summary_honest=partly: Card flow, interaction and artifact-creature finishers are genuinely present, but "develops mana quickly" oversells a build with 36 lands, no Sol Ring-class fast rocks, and mostly conditional accelerants, and the summary never signals that the deck has no combo or fast kill despite the bracket-4 ask.
- [INFO] `curve_summary`: average mana value 2.89 over 63 nonland cards
- [WARN] `profile_off_band`: the mana available on turn four is 4.64, and bracket 4 wants 4.8 or more (mean over 10000 hands)
- [INFO] `mana_pass`: the builder moved 7 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 26 Island | land | Reliable blue mana for the focused artifact shell.
- 1 Seat of the Synod | land | An artifact land that supports the artifact-focused mana base.
- 1 Mystic Sanctuary | land | A blue land option for the mana base.
- 1 Otawara, Soaring City | land | A blue land option that also occupies a flexible deck slot.
- 1 Glimmervoid | land | Colored mana support for the artifact-heavy plan.
- 1 Spire of Industry | land | Colored mana support alongside the deck's artifacts.
- 1 Urza's Saga | land | Artifact-focused utility land for the mana base.
- 1 Academy Ruins | land | Artifact-focused utility land.
- 1 Inventors' Fair | land | Artifact-focused utility land.
- 1 Mishra's Workshop | land | A powerful artifact-focused mana land.
- 1 Buried Ruin | land | Artifact-focused utility land.
- 1 Cargo Ship | ramp | Provides additional mana development.
- 1 Chief Engineer | ramp | Provides ramp within the artifact shell.
- 1 Coin of Mastery | ramp | An artifact ramp piece for early development.
- 1 Crystal Skull, Isu Spyglass | ramp | Provides artifact-based ramp.
- 1 Grand Architect | ramp | A blue ramp piece for the artifact strategy.
- 1 Inspiring Statuary | ramp | Provides ramp while rewarding an artifact-heavy build.
- 1 Metalworker | ramp | A dedicated artifact ramp piece.
- 1 Moonsnare Prototype | ramp | Low-cost artifact ramp for faster development.
- 1 Mox Opal | ramp | Fast artifact ramp for explosive openings.
- 1 Network Terminal | ramp | Artifact-based mana development.
- 1 Powerstone Shard | ramp | Additional artifact ramp.
- 1 Solar Array | ramp | Supports the deck's mana development.
- 1 Vedalken Engineer | ramp | A blue artifact-focused ramp card.
- 1 Urza's Command | ramp | Flexible ramp support for the artifact plan.
- 1 Braided Net // Braided Quipu | draw | An artifact draw option.
- 1 Esoteric Duplicator | draw | Artifact-based card advantage.
- 1 Forensic Gadgeteer | draw | Provides card advantage in the artifact shell.
- 1 One with the Machine | draw | Draw support for a deck with substantial artifacts.
- 1 Riddlesmith | draw | Provides draw support for the artifact plan.
- 1 Sai, Master Thopterist | draw | An artifact-focused source of card advantage.
- 1 Thirst for Knowledge | draw | Efficient blue card advantage.
- 1 Thoughtcast | draw | Artifact-focused card advantage.
- 1 Curator's Ward | interaction | Interaction that supports a key permanent.
- 1 Disruption Protocol | interaction | A blue interaction spell for protecting the game plan.
- 1 Escape Protocol | interaction | Flexible interaction within the artifact shell.
- 1 Ghostly Flicker | interaction | A flexible blue interaction option.
- 1 Ice Out | interaction | Blue stack interaction.
- 1 Metallic Rebuke | interaction | Artifact-focused interaction.
- 1 Override | interaction | Blue interaction for defending the plan.
- 1 Padeem, Consul of Innovation | interaction | An artifact-focused interaction piece.
- 1 Stoic Rebuttal | interaction | Reliable blue interaction.
- 1 Welding Jar | interaction | Low-cost artifact protection.
- 1 Aether Spellbomb | removal | Low-cost artifact removal.
- 1 Arcum Dagsson | removal | Artifact-focused removal utility.
- 1 Cyber Conversion | removal | Blue targeted removal.
- 1 Ravenform | removal | Blue removal that handles problematic permanents.
- 1 Resculpt | removal | Efficient blue removal.
- 1 Skysovereign, Consul Flagship | removal | An artifact removal threat.
- 1 Transmogrifying Wand | removal | Repeatable artifact-based removal.
- 1 Emry, Lurker of the Loch | synergy | A core artifact synergy piece.
- 1 Etherium Sculptor | synergy | Supports the deck's artifact synergies.
- 1 Foundry Inspector | synergy | An artifact-focused synergy piece.
- 1 Mystic Forge | synergy | A powerful artifact synergy engine.
- 1 Transmute Artifact | synergy | Artifact-focused synergy and deck access.
- 1 Whir of Invention | synergy | Instant-speed artifact synergy and deck access.
- 1 Arcbound Crusher | threat | An artifact creature threat.
- 1 Chrome Steed | threat | An artifact creature for board pressure.
- 1 Darksteel Juggernaut | threat | A durable artifact threat.
- 1 Frogmite | threat | A low-cost artifact creature threat.
- 1 Jhoira's Familiar | threat | An artifact creature that advances the board.
- 1 Lodestone Golem | threat | An artifact creature threat.
- 1 Lodestone Myr | threat | An artifact creature that pressures opponents.
- 1 Master Transmuter | threat | A high-impact artifact creature threat.
- 1 Rust Elemental | threat | An artifact creature threat for closing games.
- 1 Traxos, Scourge of Kroog | threat | A large artifact creature threat.
- 1 Whirler Rogue | threat | A blue artifact-focused threat.
- 1 Engineered Explosives | wipe | A compact reset option.
- 1 Hurkyl's Recall | wipe | A blue artifact-focused board reset.
- 1 Voltaic Key | draw | the mana pass added it to bring the mana base inside the power level
- 1 Lady Octopus, Inspired Inventor | draw | the mana pass added it to bring the mana base inside the power level
- 1 Artificer's Assistant | removal | the mana pass added it to bring the mana base inside the power level
- 1 Kitesail Larcenist | ramp | the mana pass added it to bring the mana base inside the power level
- 1 Luck Bobblehead | ramp | the mana pass added it to bring the mana base inside the power level

</details>

### 4. dinosaur tribal, bracket 2

Format: Commander. Theme: dinosaurs. Pool: any_card. Shortlist: 269 names.

Commander: Gishath, Sun's Avatar.

Grade: bad, score 0.00, model 20260910T012734Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $674.94 to buy, $674.94 the whole deck.

**Summary:** Gishath leads a straightforward Dinosaur deck that develops its mana, fills the board with themed creatures, and presses opponents through combat with increasingly large threats. Creature-based draw keeps the deck moving, while Dinosaur removal and a small suite of protective cards help preserve momentum. The deck gives up speed and compact finishes for a clear tribal game plan, high mana requirements, and a reliance on developing a creature board.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The curve sits high for the format, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Every slot supports one Gishath dinosaur-tribal plan: heavy ramp and cost reducers to land the commander, dinosaurs to hit off his trigger, creature-based draw, and combat-oriented payoffs like Garruk's Uprising and Regisaur Alpha.
- PLAN theme_fit=yes: It is a Gishath-led Dinosaur Commander deck of roughly precon-level power with a simple attack-and-ramp game plan, matching the request for a beginner bracket 2 build, though a few pieces (Taiga, Cavern of Souls, Spire Garden) are pricier than a typical new-player list would use since no budget was named.
- PLAN useful_as_built=yes: 38 lands with deep fixing plus ten ramp pieces support the admittedly high curve, and there are plenty of win routes through big trampling bodies, Gishath's cascade of dinosaurs, and Goring Ceratops/Garruk's Uprising style combat pushes.
- PLAN summary_honest=partly: The described plan (ramp, tribal creatures, creature draw, dino removal, protection) matches the list accurately, but the summary never mentions the two symmetrical board wipes (Blasphemous Act, Austere Command) that can sweep the deck's own creature board — a real element a new pilot should know about.
- [INFO] `curve_summary`: average mana value 3.90 over 61 nonland cards
- [INFO] `mana_pass`: the builder moved 3 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 1 Command Tower | land | Three-color land for a consistent mana base.
- 1 Cavern of Souls | land | Tribal land that supports the Dinosaur-heavy creature base.
- 1 City of Brass | land | Flexible three-color fixing.
- 1 Mana Confluence | land | Flexible three-color fixing.
- 1 Bountiful Promenade | land | Green-white source for the mana base.
- 1 Brushland | land | Green-white source that can enter ready to use.
- 1 Canopy Vista | land | Green-white land that supports the basic-land package.
- 1 Temple Garden | land | Green-white source for early green and white needs.
- 1 Cinder Glade | land | Green-red land that works with the basic-land package.
- 1 Karplusan Forest | land | Green-red source that can enter ready to use.
- 1 Rockfall Vale | land | Green-red source for the mana base.
- 1 Rootbound Crag | land | Green-red source supported by the other lands.
- 1 Stomping Ground | land | Green-red source for early mana development.
- 1 Taiga | land | Green-red source that helps cast Dinosaur spells.
- 1 Spire Garden | land | Green-red source for multiplayer games.
- 1 Sacred Foundry | land | Red-white source for the mana base.
- 1 Spectator Seating | land | Red-white source for multiplayer games.
- 7 Forest | land | Basic green mana for ramp and Dinosaur spells.
- 9 Plains | land | Basic white mana for the deck's white cards.
- 5 Mountain | land | Basic red mana for the deck's red cards.
- 1 Sol Ring | ramp | Efficient early mana acceleration.
- 1 Arcane Signet | ramp | Reliable fixing and acceleration for all three colors.
- 1 Birds of Paradise | ramp | Early creature-based color fixing.
- 1 Farseek | ramp | Early land-based color fixing.
- 1 Nature's Lore | ramp | Early land ramp that supports the mana base.
- 1 Three Visits | ramp | Additional efficient land ramp.
- 1 Rampant Growth | ramp | Straightforward early mana development.
- 1 Cultivate | ramp | Land ramp that helps make later land drops.
- 1 Kodama's Reach | ramp | Additional land ramp for a mana-hungry deck.
- 1 Thunderherd Migration | ramp | Dinosaur-themed land ramp.
- 1 Beast Whisperer | draw | Creature-focused draw for the Dinosaur plan.
- 1 Garruk's Uprising | draw | Draw support for a deck built around large creatures.
- 1 Guardian Project | draw | Ongoing card advantage for the creature-heavy plan.
- 1 Harmonize | draw | Simple, dependable card draw.
- 1 Return of the Wildspeaker | draw | Flexible draw support for a large-creature deck.
- 1 Rishkar's Expertise | draw | High-impact card draw for the Dinosaur strategy.
- 1 Ripjaw Raptor | draw | Dinosaur draw option that stays on theme.
- 1 Earthshaker Dreadmaw | draw | Dinosaur card draw that also adds a substantial body.
- 1 Curious Altisaur | draw | On-theme Dinosaur draw support.
- 1 Toski, Bearer of Secrets | draw | Creature-based draw support for an attacking deck.
- 1 Vanquisher's Banner | draw | Tribal draw support for the Dinosaur creature base.
- 1 Heroic Intervention | interaction | Flexible protection for the board.
- 1 Boros Charm | interaction | Low-cost protection and combat interaction.
- 1 Akroma's Will | interaction | Versatile combat-focused interaction.
- 1 Lightning Greaves | interaction | Protection equipment for an important creature.
- 1 Swiftfoot Boots | interaction | Additional protective equipment for key creatures.
- 1 Temple Altisaur | interaction | Dinosaur-based defensive interaction.
- 1 Savage Stomp | removal | Efficient Dinosaur-themed creature removal.
- 1 Ravenous Sailback | removal | Dinosaur removal that contributes to the creature plan.
- 1 Territorial Allosaurus | removal | Creature removal attached to a Dinosaur.
- 1 Thrashing Brontodon | removal | Versatile Dinosaur removal option.
- 1 Trumpeting Carnosaur | removal | Dinosaur removal that remains a meaningful threat.
- 1 Itzquinth, Firstborn of Gishath | removal | Low-cost Dinosaur-based removal.
- 1 Zacama, Primal Calamity | removal | Large Dinosaur removal option for late games.
- 1 Commune with Dinosaurs | synergy | Early Dinosaur-themed setup for the deck.
- 1 Kinjalli's Caller | synergy | Early support creature for the Dinosaur plan.
- 1 Otepec Huntmaster | synergy | Dinosaur-focused support creature.
- 1 Marauding Raptor | synergy | Dinosaur-focused support that helps build the board.
- 1 Huatli's Raptor | synergy | Low-cost Dinosaur that supports the deck's theme.
- 1 Dromosaur | synergy | On-theme Dinosaur support creature.
- 1 Deathgorge Scavenger | synergy | Utility Dinosaur that supports the tribal plan.
- 1 Giant Cindermaw | synergy | Dinosaur utility creature for the creature-heavy strategy.
- 1 Hunting Velociraptor | synergy | Dinosaur support creature for the deck's main plan.
- 1 Kinjalli's Sunwing | synergy | Dinosaur utility that supports combat-oriented games.
- 1 Orazca Frillback | synergy | Flexible Dinosaur utility for the tribal shell.
- 1 Sunfrill Imitator | synergy | Low-cost Dinosaur synergy piece.
- 1 Tyrox, Saurid Tyrant | synergy | Legendary Dinosaur that supports the tribal plan.
- 1 Regisaur Alpha | threat | Dinosaur threat that fits the tribal strategy.
- 1 Carnage Tyrant | threat | Large Dinosaur threat for closing games through combat.
- 1 Ghalta, Stampede Tyrant | threat | Large Dinosaur threat for powerful late turns.
- 1 Etali, Primal Storm | threat | Legendary Dinosaur threat with major table presence.
- 1 Ghalta and Mavren | threat | Large legendary Dinosaur threat.
- 1 Quartzwood Crasher | threat | Combat-focused Dinosaur threat.
- 1 Palani's Hatcher | threat | Substantial Dinosaur threat for building the board.
- 1 Pantlaza, Sun-Favored | threat | Legendary Dinosaur threat that remains on theme.
- 1 Zetalpa, Primal Dawn | threat | High-end Dinosaur finisher.
- 1 Goring Ceratops | threat | Large Dinosaur attacker for combat pressure.
- 1 Rampaging Brontodon | threat | Large Dinosaur threat that rewards a developed board.
- 1 Blasphemous Act | wipe | Efficient reset when opposing creatures get ahead.
- 1 Austere Command | wipe | Flexible board reset for difficult battlefield states.
- 1 Raging Swordtooth | wipe | Dinosaur-themed board-control option.

</details>

### 5. blink Commander, owned first

Format: Commander. Theme: blink. Pool: owned_first. Shortlist: 168 names.

Commander: Gilraen, Dúnedain Protector.

Grade: bad, score 0.01, model 20260910T012734Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $214.01 the whole deck.

**Summary:** This deck develops a stable white mana base, uses blink-focused creatures and effects to revisit its value pieces, and protects its commander and best threats while building a board. It wins by turning that accumulated creature value into sustained combat pressure, backed by targeted answers and a few board resets. It gives up some speed for a creature-centric plan that is strongest when it can keep its board intact and repeatedly leverage its library cards.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The commander does not place in cEDH events, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: There is a real blink/ETB core (Flickerwisp, Angel of Condemnation, Fiend Hunter, Palace Jailer, Angel of Sanctions/Serenity, Wall of Omens, Slip On the Ring, Personify) plus targeting spells to trigger the commander, but a sizable equipment/voltron-ish package (Puresteel Paladin, Buster Sword, Sword of the Animist, Champion's Helm, Mask of Memory, Bastion Protector) and three sweepers pull against the go-wide ETB-value board the rest of the deck builds.
- PLAN theme_fit=yes: It is a mono-white Commander deck led by Gilraen, Dúnedain Protector, a blink commander, at a bracket-3-appropriate power level with no fast mana beyond Sol Ring or compact combos, matching what was asked.
- PLAN useful_as_built=yes: 37 lands with mostly Plains for a mono-white curve, ten mana rocks/ramp, a solid removal and card-draw suite, and plenty of creatures to attack with means it plays fine straight out of the box.
- PLAN summary_honest=partly: It correctly flags the white mana base, protection suite, board resets, and the creature-combat win condition, but calling the deck 'blink-focused' oversells a list where a large fraction of slots are equipment, generic ramp, and standalone removal with no flicker interaction.
- [INFO] `curve_summary`: average mana value 2.94 over 63 nonland cards

<details><summary>The deck list</summary>

- 28 Plains | land | Primary land slot for a dependable white mana base.
- 1 Command Tower | land | Land slot that supports the mana base.
- 1 City of Brass | land | Land slot that supports the mana base.
- 1 Marsh Flats | land | Land slot that supports the mana base.
- 1 Fabled Passage | land | Land slot that supports the mana base.
- 1 Plaza of Heroes | land | Land slot that supports the mana base.
- 1 Secluded Courtyard | land | Land slot that supports the mana base.
- 1 Spire of Industry | land | Land slot that supports the mana base.
- 1 Minas Tirith | land | Land slot that supports the mana base.
- 1 Sol Ring | ramp | Efficient acceleration for deploying the commander and value creatures.
- 1 Arcane Signet | ramp | Reliable mana acceleration for the deck's white-heavy plan.
- 1 Fellwar Stone | ramp | Additional early mana acceleration.
- 1 Wayfarer's Bauble | ramp | Ramp slot that helps the deck develop its mana.
- 1 Thought Vessel | ramp | Mana acceleration that supports longer value turns.
- 1 Bender's Waterskin | ramp | Ramp slot for reaching the deck's creature and blink plays.
- 1 Inherited Envelope | ramp | Mana development for the deck's value plan.
- 1 White Lotus Tile | ramp | Ramp slot that helps advance the board early.
- 1 Sword of the Animist | ramp | Ramp option that fits the creature-focused plan.
- 1 Relic of Legends | ramp | Additional mana acceleration for multispell turns.
- 1 Adventurer's Airship | draw | Card-advantage slot for maintaining resources.
- 1 Buster Sword | draw | Equipment-based card-advantage option.
- 1 Diary of Dreams | draw | Card-advantage piece for longer games.
- 1 Energybending | draw | Draw spell that keeps the value engine supplied.
- 1 Idol of Oblivion | draw | Low-cost card-advantage support.
- 1 Inspiring Overseer | draw | Creature-based card advantage that fits the blink shell.
- 1 Lembas | draw | Card-advantage piece that supports incremental value.
- 1 Mask of Memory | draw | Equipment-based card advantage for the creature board.
- 1 Skullclamp | draw | Efficient card-advantage equipment.
- 1 Tome of Legends | draw | Card-advantage support for sustained games.
- 1 Crown of Gondor | draw | Additional card-advantage option for the creature plan.
- 1 Banishing Light | removal | Flexible removal slot for troublesome permanents.
- 1 Crib Swap | removal | Creature removal that fits the deck's reactive suite.
- 1 Destroy Evil | removal | Efficient answer for opposing threats.
- 1 Dispatch | removal | Low-cost targeted removal.
- 1 Generous Gift | removal | Broad answer for opposing permanents.
- 1 Get Lost | removal | Flexible targeted removal.
- 1 Journey to Nowhere | removal | Permanent-based creature answer.
- 1 Swords to Plowshares | removal | Efficient creature removal.
- 1 Aang, the Last Airbender | removal | Additional removal body for the deck's creature-focused game.
- 1 Boromir, Warden of the Tower | interaction | Protective interaction for the board and commander.
- 1 Clever Concealment | interaction | Defensive interaction for preserving the board.
- 1 Gift of Immortality | interaction | Protection piece for an important creature.
- 1 Lightning Greaves | interaction | Protection equipment for key creatures.
- 1 Reprieve | interaction | Flexible stack interaction that buys time.
- 1 Swiftfoot Boots | interaction | Additional protection for the commander or a threat.
- 1 Ultimate Magic: Holy | interaction | Interaction slot for protecting the game plan.
- 1 Champion's Helm | interaction | Equipment-based protection for legendary creatures.
- 1 Austere Command | wipe | Flexible reset when opposing boards get ahead.
- 1 Fumigate | wipe | Board-reset option for stabilizing the game.
- 1 Vanquish the Horde | wipe | Efficient mass answer for creature-heavy boards.
- 1 Angel of Condemnation | synergy | Blink-focused creature that supports repeated value turns.
- 1 Angel of Sanctions | synergy | Creature value piece that benefits from blink support.
- 1 Ennis, Debate Moderator | synergy | Dedicated synergy piece for the deck's blink plan.
- 1 Fiend Hunter | synergy | Creature-based value piece to reuse with blink effects.
- 1 Flickerwisp | synergy | Core blink effect for reusing the deck's creatures.
- 1 Jocasta, Automaton Avenger | synergy | Dedicated synergy piece for the creature-value plan.
- 1 Palace Jailer | synergy | Creature value piece that rewards repeated blinking.
- 1 Personify | synergy | Blink-plan support for reusing creature value.
- 1 Slip On the Ring | synergy | Blink effect that protects or refreshes a creature.
- 1 Wall of Omens | synergy | Creature value piece that works well with blink support.
- 1 Angel of Serenity | threat | Large creature threat for closing games.
- 1 Giada, Font of Hope | threat | Creature threat that advances the board.
- 1 Bronze Guardian | threat | Resilient artifact creature threat.
- 1 Champions of Minas Tirith | threat | Creature threat for building board pressure.
- 1 Faramir, Field Commander | threat | Creature threat that supports the attacking plan.
- 1 Frontline Medic | threat | Creature threat for developing the board.
- 1 Kataki, War's Wage | threat | Creature threat that pressures opposing development.
- 1 Puresteel Paladin | threat | Creature threat that fits the equipment subtheme.
- 1 Bastion Protector | threat | Creature threat that supports key legendary permanents.
- 1 Westfold Rider | threat | Creature threat for applying combat pressure.
- 1 Zack Fair | threat | Legendary creature threat for the board.
- 1 The Walls of Ba Sing Se | threat | Artifact creature threat that adds board presence.

</details>

### 6. Modern tempo, tournament

Format: Modern. Theme: tempo. Pool: any_card. Shortlist: 167 names.

Grade: typical, score 0.61, model 20260910T012734Z.

Cards: 60 main, 15 sideboard. Repair turn: no. Block findings: 0.

Cost: $574.30 to buy, $574.30 the whole deck.

**Summary:** This blue-red tempo deck develops early pressure with Ragavan, Nimble Pilferer, Ledger Shredder, and Faerie Mastermind while cheap draw, interaction, and removal keep the opponent off balance. It aims to stay ahead on the board and convert that lead into a finish, with Temporal Mastery reinforcing its spell-focused plan. It gives up heavier permanent-based threats and main-deck sweepers for speed, mana consistency, and a dense suite of low-cost plays.

The quality model grades this deck typical against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The deck runs its spells as playsets, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade.

- PLAN plan_coherent=partly: The creature suite, cheap cantrips, Bolt/Into the Flood Maw, and light counters form a coherent UR tempo shell, but four copies of Temporal Mastery — a seven-mana sorcery outside the narrow miracle window — pull hard against the low-curve tempo plan.
- PLAN theme_fit=partly: It is a Modern blue-red tempo list with the expected threats and interaction, but it contains no Delver of Secrets (or Murktide-style payoff) despite the request naming a Delver deck, and the Temporal Mastery package is not part of that archetype.
- PLAN useful_as_built=partly: The 24-land mana base, low curve, twelve two-drop-or-less threats, and eight removal spells make it playable out of the box, but four dead-on-arrival Temporal Masteries mean a meaningful fraction of draws are effectively blanks.
- PLAN summary_honest=partly: The summary accurately lists the threats, draw, and removal, but it dresses up Temporal Mastery as something that 'reinforces' the plan rather than admitting it is a four-of miracle gamble, and it never notes the absence of the namesake Delver.
- [INFO] `curve_summary`: average mana value 2.11 over 36 nonland cards

<details><summary>The deck list</summary>

- 4 Steam Vents | land | Primary blue-red dual land for casting the deck's spells.
- 4 Shivan Reef | land | Provides both deck colors without entering tapped.
- 4 Spirebluff Canal | land | Early blue-red source suited to a tempo curve.
- 4 Riverglide Pathway // Lavaglide Pathway | land | Flexible untapped source for either required color.
- 4 Island | land | Reliable blue source for blue-intensive spells.
- 4 Mountain | land | Reliable red source for Lightning Bolt.
- 4 Consider | draw | Low-cost card selection that keeps spells flowing.
- 2 Preordain | draw | Efficient card selection for finding the needed tempo piece.
- 4 Force of Negation | interaction | Free interaction helps protect the deck's tempo plan.
- 2 Spell Pierce | interaction | Cheap interaction for fighting opposing noncreature plays.
- 4 Lightning Bolt | removal | Efficient removal that fits a low-cost tempo plan.
- 4 Into the Flood Maw | removal | Low-cost removal that clears a path for the deck's creatures.
- 4 Temporal Mastery | synergy | Supports the deck's temporal-spell synergy package.
- 4 Ragavan, Nimble Pilferer | ramp | Early creature that contributes to the deck's mana development.
- 4 Ledger Shredder | draw | Creature-based card flow that rewards the deck's dense spell count.
- 4 Faerie Mastermind | draw | Creature-based card draw that supports the tempo plan.

</details>

### 7. Modern burn, casual

Format: Modern. Theme: burn. Pool: any_card. Shortlist: 290 names.

Grade: bad, score 0.39, model 20260910T012734Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $66.02 to buy, $66.02 the whole deck.

**Summary:** This mono-red burn deck establishes pressure with a threat-heavy creature suite, then uses efficient removal and synergistic red permanents to keep damage flowing. Card draw helps it continue presenting action after the opening turns, while its ramp cards support its more demanding threats. It aims to finish games by combining early pressure with repeated burn-oriented plays, giving up broader answers and defensive flexibility for a focused aggressive plan.

The quality model grades this deck below the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: Everything points at red damage, but the pieces fight each other: Hazoret wants an empty hand while Risk Factor/Browbeat refill it, Galvanic Blast is near-dead with zero artifacts, and a true burn plan is undercut by a top-heavy pile of four-mana permanents (Torbran, Ashcloud Phoenix).
- PLAN theme_fit=partly: It is mono-red and does burn face, but a burn deck asked for should be built on cheap direct-damage spells; this is closer to a clunky mono-red midrange list with only ten true burn spells and a glut of four-drops.
- PLAN useful_as_built=partly: The mana base is fine and the deck can function in a casual pod, but the curve is far too high for burn, Galvanic Blast is a blank, Hazoret is hard to turn on, and the reach to close games out of nowhere is thin.
- PLAN summary_honest=partly: The aggressive-red framing is broadly fair, but it claims "ramp cards support its more demanding threats" when Chandra, Dressed to Kill is not real ramp and nothing else ramps, and it calls Galvanic Blast efficient removal in a deck with no artifacts.
- [INFO] `curve_summary`: average mana value 2.72 over 36 nonland cards

<details><summary>The deck list</summary>

- 20 Mountain | land | Provides the deck's primary red mana source.
- 4 Ramunap Ruins | land | Adds red land slots while fitting the burn plan.
- 4 Risk Factor | draw | Provides card draw to keep the deck supplied with action.
- 2 Browbeat | draw | Adds further card draw for longer games.
- 2 Chandra, Dressed to Kill | ramp | Fills the ramp slots while remaining aligned with the mono-red plan.
- 4 Lightning Bolt | removal | Efficient removal that supports the burn-focused game plan.
- 2 Galvanic Blast | removal | Adds more low-cost removal to clear opposing threats.
- 4 Eidolon of the Great Revel | synergy | A synergistic permanent that reinforces the deck's aggressive burn direction.
- 4 Boltwave | synergy | Supplies additional synergy for the deck's damage-oriented plan.
- 4 Akki Lavarunner // Tok-Tok, Volcano Born | threat | An early threat that helps establish pressure.
- 2 Ashcloud Phoenix | threat | A resilient threat for maintaining pressure through the middle of the game.
- 4 Hazoret the Fervent | threat | A powerful threat that gives the deck a durable closing option.
- 4 Torbran, Thane of Red Fell | threat | A threat that strongly supports the mono-red damage plan.

</details>

### 8. Modern lifegain, FNM

Format: Modern. Theme: lifegain. Pool: any_card. Shortlist: 301 names.

Grade: bad, score 0.30, model 20260910T012734Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $141.75 to buy, $141.75 the whole deck.

**Summary:** This deck establishes its white-black lifegain synergies early, then uses growing payoffs and a steady creature curve to pressure the opponent. Solitude and Murderous Rider // Swift End keep opposing threats from derailing the board plan, while Angel of Invention provides a stronger closing threat. It gives up some speed and broad matchup flexibility for a focused creature-based lifegain game plan.

The quality model grades this deck below the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core is a consistent lifegain-payoff shell (Soul Warden/Attended Healer feeding Ajani's Pridemate, Twinblade Paladin, Bloodbond Vampire), but 2 Arcane Signet in a 60-card aggro deck and a lone Hateful Eidolon with zero auras pull against that plan.
- PLAN theme_fit=yes: It is a white-black Modern lifegain deck at a casual FNM power level, exactly what was asked.
- PLAN useful_as_built=yes: 24 lands with Godless Shrine and Caves of Koilos support a low, mostly two-and-three-drop curve, and the creature-plus-payoff core gives a clear route to winning games as printed.
- PLAN summary_honest=yes: Every claim — early lifegain synergies, growing payoffs, Solitude/Murderous Rider as interaction, Angel of Invention as a closer, and the admitted lack of speed and flexibility — matches the list, and it openly reports the quality model's low grade.
- [INFO] `curve_summary`: average mana value 3.11 over 36 nonland cards
- [INFO] `mana_pass`: the builder moved 1 card of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 8 Plains | land | Basic white source for the lifegain core.
- 8 Swamp | land | Basic black source for threats and removal.
- 4 Godless Shrine | land | Provides both deck colors without using a tapped-land slot.
- 4 Caves of Koilos | land | Provides flexible white-black color coverage.
- 2 Arcane Signet | ramp | Supplies early mana acceleration and both colors.
- 3 Solitude | removal | Primary removal option that also occupies a threat-capable creature slot.
- 2 Murderous Rider // Swift End | removal | Additional black removal attached to a creature card.
- 4 Inspiring Overseer | draw | Draw option that supports the white creature plan.
- 2 Lembas | draw | Low-investment draw support for the deck's early turns.
- 4 Soul Warden | synergy | Early lifegain synergy piece for the deck's creature plan.
- 4 Ajani's Pridemate | synergy | Core lifegain payoff that turns the early plan into board pressure.
- 4 Attended Healer | threat | Midgame creature threat suited to the lifegain strategy.
- 4 Twinblade Paladin | threat | Efficient white threat for converting the lifegain plan into attacks.
- 4 Bloodbond Vampire | threat | Black lifegain-focused threat that reinforces the deck's central theme.
- 2 Angel of Invention | threat | Top-end creature threat for closing games after establishing the board.
- 1 Hateful Eidolon | removal | the mana pass added it to bring the mana base inside the power level

</details>

### 9. Standard midrange, FNM

Format: Standard. Theme: midrange. Pool: any_card. Shortlist: 164 names.

Grade: baseline, score 0.39, model 20260910T012734Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $266.36 to buy, $266.36 the whole deck.

**Summary:** This black-green midrange deck develops its mana, answers opposing threats, and uses steady card draw to keep presenting creatures that demand an answer. It aims to win through sustained battlefield pressure, with protection and legendary-focused support helping its best creatures stay relevant. In return, it gives up the speed of a lower-curve aggressive strategy for stronger late-game threats and longer exchanges.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The cards pair in ways the top lists do not, and that lowers the grade. The curve sits high for the format, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core of removal, card draw, and fat black-green threats points one direction, but the protection tricks (Snakeskin Veil, Undying Malice) and the legendary-matters Vraska Joins Up serve a different, thinner sub-plan, and running every slot as a 2-of means no piece is reliably drawn to support another.
- PLAN theme_fit=yes: It is a Standard-legal-shaped black-green midrange list with a normal 24-land base and a mix of removal and threats, which is exactly the FNM-level deck asked for.
- PLAN useful_as_built=partly: The mana base is clean and there are plenty of ways to win, but the creature curve is stacked at five and six mana with only two Llanowar Elves and almost nothing to do on turns two and three, so games will regularly be decided before the deck's plan comes online.
- PLAN summary_honest=yes: The summary accurately describes the removal, draw, protection, and legendary support present in the list, and it openly concedes the high curve and slow clock rather than hiding them.
- [INFO] `curve_summary`: average mana value 3.50 over 36 nonland cards

<details><summary>The deck list</summary>

- 6 Forest | land | Basic green source for the black-green mana base.
- 6 Swamp | land | Basic black source for the black-green mana base.
- 4 Blooming Marsh | land | Dual-color land that supports both primary colors.
- 4 Overgrown Tomb | land | Dual-color land that strengthens access to black and green.
- 4 Underground Mortuary | land | Dual-color land that completes the colored-source base.
- 2 Llanowar Elves | ramp | Early creature ramp that helps deploy the deck's larger threats.
- 2 Bitter Triumph | removal | Efficient black removal for opposing threats.
- 2 Hero's Downfall | removal | Versatile removal for troublesome permanents.
- 2 Withering Torment | removal | Additional black removal to maintain board control.
- 2 Darkstar Augur | draw | A creature-based draw piece that contributes to the board.
- 2 Phyrexian Arena | draw | Reliable ongoing card draw for longer games.
- 2 Unholy Annex // Ritual Chamber | draw | A black draw engine that supports the midrange plan.
- 2 Garruk's Uprising | synergy | Supports the deck's creature-heavy game plan.
- 2 Vraska Joins Up | synergy | Supports the deck's legendary creature package.
- 2 Snakeskin Veil | synergy | Protects key creatures while supporting a board-focused plan.
- 2 Undying Malice | synergy | Helps preserve important creatures through opposing removal.
- 2 Goldvein Hydra | threat | Creature threat that can take over the battlefield.
- 2 Rottenmouth Viper | threat | Black-green creature threat that pressures opposing boards.
- 2 Vein Ripper | threat | A substantial black creature threat for closing games.
- 2 Vaultborn Tyrant | threat | Large creature threat that rewards reaching the late game.
- 2 Earthshaker Dreadmaw | threat | Large green creature that provides a decisive battlefield threat.
- 2 Scavenger Regent // Exude Toxin | threat | Flexible creature threat that fits the black-green top end.
- 2 Sephiroth, Fabled SOLDIER // Sephiroth, One-Winged Angel | threat | Legendary creature threat that works with the deck's legendary support.

</details>

### 10. Standard aggro, tournament

Format: Standard. Theme: aggro. Pool: any_card. Shortlist: 294 names.

Grade: bad, score 0.27, model 20260910T012734Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $116.66 to buy, $116.66 the whole deck.

**Summary:** This red-white aggro deck aims to keep presenting threats and maintain pressure while removal and interaction address opposing resistance. It wins by sustaining that proactive board plan, with draw helping it keep finding action instead of running out of threats. The tradeoff is a direct, board-focused strategy that is less interested in playing a long defensive game.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The cards pair in ways the top lists do not, and that lowers the grade. The curve sits high for the format, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: Most cards point at a proactive board plan (Frilled Sparkshooter, Warleader's Call, cheap removal like Case of the Gateway Express), but four copies of a six-mana Krenko's Buzzcrusher plus other top-end pull hard against an aggro clock, and Redcap Gutter-Dweller's black cost sits outside what the two-color mana supports.
- PLAN theme_fit=partly: It is a Standard-frame red-white deck built around creatures and burn as asked, but the six-drops and four-mana anthem give it a midrange curve rather than the aggro deck the request named.
- PLAN useful_as_built=no: With only 4 dual lands and no black sources, the four Redcap Gutter-Dwellers are effectively dead cards, and the pile of five- and six-mana spells on 24 lands means an aggro deck that repeatedly stalls before it can deploy its top end.
- PLAN summary_honest=partly: The summary fairly describes the removal-plus-threats shell and the modest draw from Fugitive Codebreaker and Caretaker's Talent, and it does concede the high curve, but it never acknowledges that a full playset of the deck's threats cannot be reliably cast off this mana base.
- [INFO] `curve_summary`: average mana value 3.06 over 36 nonland cards
- [INFO] `mana_pass`: the builder moved 1 card of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 11 Mountain | land | Red mana source for the aggressive red-white mana base.
- 9 Plains | land | White mana source for the aggressive red-white mana base.
- 4 Elegant Parlor | land | Dual-color land that supports both halves of the deck.
- 4 Fugitive Codebreaker | draw | Red draw option that keeps the deck supplied with action.
- 2 Caretaker's Talent | draw | White draw option for maintaining pressure after the opening turns.
- 4 Boros Charm | interaction | Red-white interaction that supports the deck's proactive plan.
- 2 Sheltered by Ghosts | interaction | White interaction for answering opposing resistance.
- 4 Case of the Gateway Express | removal | Efficient removal slot for clearing blockers and opposing threats.
- 4 Krenko's Buzzcrusher | removal | Red removal that helps force damage through.
- 4 Warleader's Call | synergy | Red-white synergy piece that reinforces the creature-focused plan.
- 4 Frilled Sparkshooter | threat | Red threat that advances the deck's aggressive battlefield plan.
- 4 Dragonback Lancer | threat | Creature threat for maintaining a steady attack.
- 4 Redcap Gutter-Dweller | threat | Red threat that fills out the deck's pressure package.

</details>

### 11. Commander with a locked card

Format: Commander. Theme: sacrifice. Pool: any_card. Shortlist: 296 names.

Commander: Karlov of the Ghost Council.

Grade: baseline, score 0.01, model 20260910T012734Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $1500.63 to buy, $1500.63 the whole deck.

**Summary:** This deck builds a creature-heavy sacrifice engine around Karlov, using inexpensive outlets and creature-death payoffs to turn expendable bodies into mana, cards, pressure, and removal. It wins by keeping the board under pressure while its threats and sacrifice synergies accumulate advantage, with several sweepers available to reset crowded games. The tradeoff is that the deck is heavily creature-dependent, so it is at its best when it can establish an outlet and enough bodies to fuel it.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every card serves one aristocrats plan — cheap sacrifice outlets, expendable/recursive bodies, and death-trigger drain payoffs that also feed Karlov's lifegain counters — with removal and edicts that run off the same engine.
- PLAN theme_fit=yes: It is exactly a Commander sacrifice deck headed by Karlov of the Ghost Council, in Orzhov, with drain effects that pump Karlov's counters as requested.
- PLAN useful_as_built=yes: 36 well-balanced dual/utility lands plus Sol Ring and altars support a low curve, and there are multiple win routes (Zulaport/Bastion/Vindictive Vampire drain loops, Karlov's exile removal, and beefy threats), so it plays fine out of the box.
- PLAN summary_honest=yes: The claims of cheap outlets, death payoffs turning bodies into mana/cards/removal, and 'several sweepers' (Toxic Deluge, Meathook Massacre, Austere Command) all match cards actually in the list, and it openly admits the creature-dependence.
- [INFO] `curve_summary`: average mana value 3.02 over 63 nonland cards

<details><summary>The deck list</summary>

- 10 Plains | land | Basic white source for the mana base.
- 9 Swamp | land | Basic black source for the mana base.
- 1 Command Tower | land | Reliable two-color land for the mana base.
- 1 Godless Shrine | land | White-black source for the mana base.
- 1 Scrubland | land | White-black source for the mana base.
- 1 Caves of Koilos | land | White-black source for the mana base.
- 1 City of Brass | land | Flexible colored source for the mana base.
- 1 Mana Confluence | land | Flexible colored source for the mana base.
- 1 Vault of Champions | land | White-black source for the mana base.
- 1 Shattered Sanctum | land | White-black source for the mana base.
- 1 Isolated Chapel | land | White-black source for the mana base.
- 1 Fetid Heath | land | White-black source for the mana base.
- 1 Reflecting Pool | land | Flexible colored source alongside the dual lands.
- 1 Exotic Orchard | land | Flexible colored source for the mana base.
- 1 Cavern of Souls | land | Creature-focused colored source for the mana base.
- 1 Secluded Courtyard | land | Creature-focused colored source for the mana base.
- 1 High Market | land | Sacrifice-friendly utility land.
- 1 Phyrexian Tower | land | Sacrifice-friendly utility land.
- 1 Diamond Valley | land | Sacrifice-friendly utility land.
- 1 Sol Ring | ramp | Required efficient mana acceleration.
- 1 Ashnod's Altar | ramp | Sacrifice outlet that converts creatures into mana.
- 1 Phyrexian Altar | ramp | Sacrifice outlet that converts creatures into mana.
- 1 Pitiless Plunderer | ramp | Mana payoff for the sacrifice plan.
- 1 Priest of Forgotten Gods | ramp | Sacrifice-based mana acceleration.
- 1 Warren Soultrader | ramp | Sacrifice-based mana acceleration.
- 1 Pawn of Ulamog | ramp | Mana payoff for creatures dying.
- 1 Skullport Merchant | ramp | Sacrifice-oriented mana support.
- 1 Deadly Dispute | ramp | Sacrifice-enabled burst of resources.
- 1 Culling the Weak | ramp | Fast sacrifice-based mana burst.
- 1 Corrupted Conviction | draw | Low-cost card flow that uses a sacrifice.
- 1 Disciple of Bolas | draw | Sacrifice payoff that refills resources.
- 1 Ecstatic Awakener // Awoken Demon | draw | Sacrifice-themed card flow and a body for the board.
- 1 Lord Skitter's Butcher | draw | Creature-based card flow for the sacrifice shell.
- 1 Shadowheart, Dark Justiciar | draw | Sacrifice payoff that converts creatures into cards.
- 1 Smothering Abomination | draw | Creature-death card flow for the deck's core plan.
- 1 Tevesh Szat, Doom of Fools | draw | Repeatable card flow with sacrifice support.
- 1 Thraxodemon | draw | Sacrifice-oriented creature that provides card flow.
- 1 Vampiric Rites | draw | Repeatable sacrifice outlet that provides card flow.
- 1 Village Rites | draw | Efficient card flow from a disposable creature.
- 1 Baron Bertram Graywater | draw | Creature-based card flow for the sacrifice strategy.
- 1 Attrition | removal | Sacrifice-powered removal for troublesome creatures.
- 1 Ayli, Eternal Pilgrim | removal | Sacrifice-oriented creature removal.
- 1 Bone Shards | removal | Low-cost removal that can use a sacrifice.
- 1 Bone Splinters | removal | Low-cost removal that turns a creature into an answer.
- 1 Eaten Alive | removal | Sacrifice-compatible creature removal.
- 1 Grave Pact | removal | Creature deaths become removal pressure.
- 1 Dictate of Erebos | removal | Creature deaths become removal pressure.
- 1 Spark Harvest | removal | Removal that fits the sacrifice shell.
- 1 Yawgmoth, Thran Physician | removal | Sacrifice-based removal utility.
- 1 Cartel Aristocrat | interaction | Flexible sacrifice outlet that protects the board plan.
- 1 Fanatical Devotion | interaction | Sacrifice-based protection for key permanents.
- 1 Flare of Fortitude | interaction | Protective interaction for a committed board.
- 1 Gift of Doom | interaction | Protection that works well with expendable creatures.
- 1 Kinzu of the Bleak Coven | interaction | Interactive creature that supports the sacrifice plan.
- 1 Nightmare Shepherd | interaction | Creature-focused interaction for the sacrifice shell.
- 1 Promise of Tomorrow | interaction | Protective interaction that rewards creature losses.
- 1 Spirit Bonds | interaction | Board-protection interaction for creature-heavy games.
- 1 Toxic Deluge | wipe | Efficient reset when opposing boards get ahead.
- 1 The Meathook Massacre | wipe | Board reset that fits the creature-death theme.
- 1 Austere Command | wipe | Flexible board reset for difficult board states.
- 1 Elas il-Kor, Sadistic Pilgrim | synergy | Central creature-death payoff for the sacrifice plan.
- 1 Zulaport Cutthroat | synergy | Creature-death payoff that advances the deck's plan.
- 1 Bastion of Remembrance | synergy | Persistent creature-death payoff.
- 1 Viscera Seer | synergy | Low-cost repeatable sacrifice outlet.
- 1 Carrion Feeder | synergy | Low-cost repeatable sacrifice outlet.
- 1 Bartolomé del Presidio | synergy | Efficient sacrifice outlet for the creature package.
- 1 Woe Strider | synergy | Sacrifice outlet that supplies a body for the plan.
- 1 Hidden Stockpile | synergy | Ongoing support for sacrifice-focused turns.
- 1 Chthonian Nightmare | synergy | Recurring sacrifice support for the deck's plan.
- 1 Victimize | synergy | Sacrifice-based return spell that rebuilds the board.
- 1 Basri's Lieutenant | threat | Creature threat that contributes to the board plan.
- 1 Demon of Catastrophes | threat | Large sacrifice-compatible creature threat.
- 1 Flesh-Eater Imp | threat | Creature threat that benefits from the sacrifice shell.
- 1 Liliana, Heretical Healer // Liliana, Defiant Necromancer | threat | Creature-death-themed threat for longer games.
- 1 Mondrak, Glory Dominus | threat | Powerful creature threat for a token-oriented board.
- 1 Ratadrabik of Urborg | threat | Legendary creature threat that supports a creature-heavy plan.
- 1 Requiem Angel | threat | High-impact creature threat for attrition games.
- 1 Vindictive Vampire | threat | Creature-death payoff that pressures opponents.
- 1 Vito's Inquisitor | threat | Low-cost vampire threat for the board.
- 1 Threefold Thunderhulk | threat | Large artifact creature threat for closing games.
- 1 Titan Hunter | threat | Creature threat that supports an aggressive finish.
- 1 Tomb Tyrant | threat | Creature threat that provides a resilient top end.

</details>

### 12. Commander on a budget

Format: Commander. Theme: tokens. Pool: any_card. Shortlist: 296 names.

Commander: Adeline, Resplendent Cathar.

Grade: bad, score 0.00, model 20260910T012734Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $68.59 to buy, $68.59 the whole deck.

**Summary:** Adeline leads a proactive white token deck that builds a wide board, reinforces it with token-focused synergy pieces, and turns repeated attacks into pressure. It wins chiefly through combat backed by creature and enchantment threats, while broad removal and a few board wipes help clear paths or recover from crowded boards. The deck gives up premium acceleration and expensive finishers for a low-cost, board-centric game plan that can be slower to rebuild after repeated disruption.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The deck makes less mana on turn four than the norm, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every slot supports one plan — make creature tokens, pump them with anthems (Intangible Virtue, Wedding Announcement, Divine Visitation) and swing with Adeline — and even the odd treasure-trinket ramp package doesn't pull against that, though it is a weak, off-theme cluster of nine slots.
- PLAN theme_fit=yes: It is a mono-white Adeline token deck with no expensive staples or fast mana, matching the requested commander, theme, budget and bracket-2 power level.
- PLAN useful_as_built=yes: 37 lands in mono-white with a low curve and abundant token makers means the deck casts its spells and closes games through Adeline plus anthems, even if the treasure-artifact ramp is slow.
- PLAN summary_honest=yes: The summary's claims — combat-centric wins, enchantment/creature threats, spot removal plus a few wipes (Hour of Reckoning is token-friendly), and an admitted lack of premium acceleration and big finishers — all match the list as written.
- [INFO] `curve_summary`: average mana value 3.26 over 62 nonland cards

<details><summary>The deck list</summary>

- 28 Plains | land | Provides the deck’s primary white mana.
- 1 Ancient Den | land | Adds a white-producing land with artifact synergies available elsewhere in the deck.
- 1 Castle Ardenvale | land | Provides a white land slot with useful late-game utility.
- 1 Command Tower | land | Provides reliable mana for Adeline’s color identity.
- 1 Eiganjo, Seat of the Empire | land | Provides white mana while occupying a flexible utility-land slot.
- 1 Exotic Orchard | land | Adds an inexpensive land that can supply the deck’s color.
- 1 Legion's Landing // Adanto, the First Fort | land | Fits the token plan while filling a land slot after transformation.
- 1 Path of Ancestry | land | Provides white mana and supports the deck’s creature-focused plan.
- 1 Secluded Courtyard | land | Supplies white mana for the deck’s many creature spells.
- 1 Windbrisk Heights | land | Provides white mana in a low-cost utility-land slot.
- 1 Bucknard's Everfull Purse | ramp | Adds a budget artifact ramp piece.
- 1 Charisma Bobblehead | ramp | Provides low-cost artifact ramp.
- 1 Coin of Mastery | ramp | Fills a budget ramp slot.
- 1 Collector's Vault | ramp | Provides artifact-based mana development.
- 1 Druidic Satchel | ramp | Adds a cheap artifact ramp option.
- 1 Goldvein Pick | ramp | Supports mana development while fitting the combat-focused deck.
- 1 Idol of False Gods | ramp | Provides a low-cost ramp slot.
- 1 Noble's Purse | ramp | Adds inexpensive artifact ramp.
- 1 Prying Blade | ramp | Supports mana development through the deck’s attacking plan.
- 1 The Restoration of Eiganjo // Architect of Restoration | ramp | Provides white ramp while remaining on the token-focused plan.
- 1 Angelic Sell-Sword | draw | Supplies a budget card-advantage creature.
- 1 Bygone Bishop | draw | Provides card advantage from the deck’s creature-heavy construction.
- 1 Court of Grace | draw | Adds continuing card advantage that suits the token strategy.
- 1 Dawn of Hope | draw | Offers repeatable card advantage for a token deck.
- 1 Faramir, Field Commander | draw | Provides card advantage on a creature that fits the board-focused plan.
- 1 Glimmer Seeker | draw | Adds an inexpensive card-advantage creature.
- 1 Idol of Oblivion | draw | Provides efficient card advantage for a token-focused deck.
- 1 Platoon Dispenser | draw | Adds card advantage on a relevant artifact creature.
- 1 Search the Premises | draw | Provides ongoing card advantage during creature combat.
- 1 Staff of the Storyteller | draw | Adds low-cost artifact card advantage.
- 1 Wedding Announcement // Wedding Festivity | draw | Combines card advantage with the deck’s token-focused board plan.
- 1 Basri Ket | interaction | Adds protective, board-oriented interaction.
- 1 Basri, Tomorrow's Champion | interaction | Provides interaction while supporting the creature plan.
- 1 Lena, Selfless Champion | interaction | Protects the token board from opposing answers.
- 1 Moogles' Valor | interaction | Provides a low-cost interactive combat trick.
- 1 Rootborn Defenses | interaction | Helps preserve a wide token board against disruption.
- 1 Spirit Bonds | interaction | Provides flexible interaction that fits the token theme.
- 1 Aerial Assault | removal | Provides inexpensive creature removal.
- 1 Banishing Slash | removal | Adds efficient white removal.
- 1 Battle Menu | removal | Provides a flexible budget removal spell.
- 1 Citizen's Crowbar | removal | Offers artifact and enchantment removal in an equipment slot.
- 1 Generous Gift | removal | Answers a broad range of opposing permanents.
- 1 Kellan's Lightblades | removal | Provides efficient removal at a low cost.
- 1 Skyclave Apparition | removal | Puts removal on a creature that contributes to the board.
- 1 Ceaseless Conflict | wipe | Provides a budget reset for overwhelming boards.
- 1 Crisis of Conscience | wipe | Adds another inexpensive board-clearing option.
- 1 Hour of Reckoning | wipe | Fits a token deck’s need for a favorable board reset.
- 1 Aligned Heart | synergy | Supports the deck’s central token strategy.
- 1 Anafenza, Unyielding Lineage | synergy | Adds a token-focused synergy creature.
- 1 Anointer Priest | synergy | Supports the deck’s wide-board token plan.
- 1 Animation Module | synergy | Provides artifact-based token synergy.
- 1 Automated Assembly Line | synergy | Adds a low-cost artifact token payoff.
- 1 Cat Collector | synergy | Contributes directly to the deck’s token theme.
- 1 Clarion Spirit | synergy | Provides an inexpensive creature-based token synergy piece.
- 1 Divine Visitation | synergy | Provides a powerful payoff for producing tokens.
- 1 Felidar Retreat | synergy | Supports a land-driven token plan.
- 1 Horn of Gondor | synergy | Adds a thematic token-focused artifact payoff.
- 1 Intangible Virtue | synergy | Rewards building a wide token board.
- 1 Oketra's Monument | synergy | Supports the creature-heavy token strategy.
- 1 Rosie Cotton of South Lane | synergy | Provides a low-cost payoff for a growing token board.
- 1 Ajani's Chosen | threat | Provides a token-oriented threat that rewards the deck’s enchantments.
- 1 Archangel Elspeth | threat | Adds a resilient top-end threat for a wide board.
- 1 Archon of Sun's Grace | threat | Provides an enchantment-linked token threat.
- 1 Attended Healer | threat | Adds a low-cost token-producing threat.
- 1 Basri's Lieutenant | threat | Provides a creature threat that supports the board plan.
- 1 Cemetery Protector | threat | Adds a useful creature threat at a modest cost.
- 1 Defiler of Faith | threat | Provides a substantial white creature threat.
- 1 Emeria Angel | threat | Turns land development into a token-based threat.
- 1 Gideon, Ally of Zendikar | threat | Adds a planeswalker threat suited to a creature-focused strategy.
- 1 God-Eternal Oketra | threat | Provides a durable high-impact creature threat.
- 1 Hero of Bladehold | threat | Pressures opponents through the deck’s combat plan.
- 1 Oketra the True | threat | Adds a durable creature threat for the white token shell.

</details>

### 13. owned first, and the commander is not owned

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 229 names.

Commander: Karlov of the Ghost Council.

Grade: bad, score 0.00, model 20260910T012734Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $155.95 the whole deck.

**Summary:** Karlov leads a white-black lifegain strategy that develops a steady mana base, uses lifegain-focused support and equipment to build a commanding board, and follows with a diverse creature threat suite. It wins by converting that sustained board presence into pressure while using targeted answers, protective interaction, and selective battlefield resets to preserve favorable positions. The deck gives up some raw speed for a fuller land base, a broad answer package, and resilient midgame play.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves one color far better than another, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core of lifegain triggers, equipment, and commander protection does support Karlov's counter-and-exile engine, but three symmetrical wipes (including Vanquish the Horde) and a wide, mostly incidental-lifegain creature suite pull against the "commanding board" plan the summary describes.
- PLAN theme_fit=yes: It is a white-black Karlov lifegain Commander deck with a bracket-3-appropriate, low-tutor, low-game-changer build, matching the request as written.
- PLAN useful_as_built=yes: 34 lands plus nine ramp pieces, eleven draw sources, nine removal spells, and a real set of creature and commander-based win paths make this a deck a player can pick up and cast, despite the uneven Plains/Swamp split and lack of duals.
- PLAN summary_honest=yes: Every element claimed — mana base and ramp, lifegain support and equipment, a creature threat suite, targeted removal, protection pieces, and selective wipes — is present in the list, and it openly concedes the heavy basic-land count and speed tradeoff.
- [WARN] `not_owned`: Karlov of the Ghost Council: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.10 over 63 nonland cards

<details><summary>The deck list</summary>

- 18 Plains | land | Forms the white basic-land base.
- 12 Swamp | land | Forms the black basic-land base.
- 1 Command Tower | land | Adds a dual-color land slot.
- 1 Exotic Orchard | land | Adds a dual-color land slot.
- 1 City of Brass | land | Adds a flexible colored land slot.
- 1 Grand Coliseum | land | Adds a flexible colored land slot.
- 1 Path of Ancestry | land | Adds a dual-color land slot.
- 1 Spire of Industry | land | Adds a flexible land slot alongside the artifact package.
- 1 Arcane Signet | ramp | Provides efficient color fixing and ramp.
- 1 Sol Ring | ramp | Provides efficient early ramp.
- 1 Fellwar Stone | ramp | Provides another inexpensive mana source.
- 1 Commander's Sphere | ramp | Provides mana fixing for the deck.
- 1 Chromatic Lantern | ramp | Provides broad mana fixing.
- 1 Thought Vessel | ramp | Provides a durable mana source.
- 1 Sword of the Animist | ramp | Supports the ramp package through the deck's creatures.
- 1 Inherited Envelope | ramp | Fills a ramp slot from the available library.
- 1 White Auracite | ramp | Fills a ramp slot from the available library.
- 1 Wayfarer's Bauble | ramp | Provides an early ramp option.
- 1 Buster Sword | draw | Adds a listed draw option to the equipment package.
- 1 Call of the Ring | draw | Provides a draw slot from the available library.
- 1 Exemplar of Light | draw | Adds draw while remaining part of the creature plan.
- 1 Inspiring Overseer | draw | Adds a creature-based draw option.
- 1 Lembas | draw | Provides a compact draw option.
- 1 Mask of Memory | draw | Adds draw to the equipment package.
- 1 Night's Whisper | draw | Provides a low-cost draw spell.
- 1 Skullclamp | draw | Adds draw to the creature package.
- 1 Tome of Legends | draw | Provides repeatable access to the draw package.
- 1 Wall of Omens | draw | Adds an early defensive draw option.
- 1 Idol of Oblivion | draw | Fills a draw slot from the available library.
- 1 Banishing Light | removal | Supplies a listed removal effect.
- 1 Bitter Triumph | removal | Supplies efficient black removal.
- 1 Dismember | removal | Adds a flexible removal option.
- 1 Dispatch | removal | Adds inexpensive creature interaction.
- 1 Fatal Push | removal | Provides low-cost removal.
- 1 Generous Gift | removal | Answers a wide range of opposing permanents.
- 1 Get Lost | removal | Adds efficient white removal.
- 1 Infernal Grasp | removal | Provides direct creature removal.
- 1 Stroke of Midnight | removal | Provides another flexible removal effect.
- 1 Bastion Protector | interaction | Protects the commander-centered game plan.
- 1 Champion's Helm | interaction | Supports protection for the deck's key creature.
- 1 Clever Concealment | interaction | Provides a defensive interaction slot.
- 1 Darksteel Plate | interaction | Supports protection for an important creature.
- 1 Lightning Greaves | interaction | Provides protection within the equipment package.
- 1 Reprieve | interaction | Adds flexible stack interaction.
- 1 Swiftfoot Boots | interaction | Provides another protection option.
- 1 Unbreakable Formation | interaction | Provides a protective combat-oriented interaction spell.
- 1 Angel of Vitality | synergy | Supports the requested lifegain synergy plan.
- 1 Compassionate Healer | synergy | Supports the requested lifegain synergy plan.
- 1 Kor Firewalker | synergy | Supports the requested lifegain synergy plan.
- 1 Light of Promise | synergy | Supports the commander-centered lifegain plan.
- 1 Rosie Cotton of South Lane | synergy | Supports the deck's synergy package.
- 1 Second Breakfast | synergy | Supports the requested lifegain synergy plan.
- 1 Aettir and Priwen | synergy | Adds equipment-based synergy.
- 1 Dancer's Chakrams | synergy | Adds equipment-based synergy.
- 1 Elixir | synergy | Fills a synergy slot from the available library.
- 1 Well-Worn Spatula | synergy | Adds another equipment synergy piece.
- 1 Angel of Invention | threat | Provides a substantial white threat.
- 1 Bill the Pony | threat | Adds a threat from the available library.
- 1 Dawnhand Eulogist | threat | Adds a threat from the available library.
- 1 Gollum, Silent Slinker // Meager Meal | threat | Provides a black threat with an additional card face.
- 1 Lo and Li, Twin Tutors | threat | Adds a legendary threat to the creature suite.
- 1 Lyra Dawnbringer | threat | Provides a powerful white threat.
- 1 Minwu, White Mage | threat | Adds a legendary white threat.
- 1 Rabaroo Troop | threat | Adds a threat from the available library.
- 1 Rooftop Percher | threat | Adds a threat from the available library.
- 1 Shattered Angel | threat | Provides another white threat.
- 1 Sneering Shadewriter | threat | Provides a black threat.
- 1 Victory's Herald | threat | Adds a high-impact white threat.
- 1 Austere Command | wipe | Provides a flexible battlefield reset.
- 1 Fumigate | wipe | Provides a full-board reset.
- 1 Vanquish the Horde | wipe | Adds an efficient reset option.

</details>

### 14. the user delegates the commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 244 names.

Commander: Atarka, World Render.

Grade: bad, score 0.01, model 20260910T012734Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $751.95 to buy, $751.95 the whole deck.

**Summary:** This deck develops its mana, deploys a steady sequence of Dragons, and turns its creature board into decisive combat attacks. It backs that pressure with Dragon-focused removal, sweepers, and protection for key threats, while several draw cards help it recover after trading resources. The deck gives up some speed and stack control in exchange for a committed creature-based game plan, so it performs best when it can establish a board and keep its largest flyers attacking.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The deck makes less mana on turn four than the norm, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every card serves the same plan of ramping into Dragons and swinging with a doubled-damage Atarka, with cost reducers, haste/protection equipment, dragon-tribal payoffs like Terror of the Peaks and Dragon Tempest, and only a couple of oddities (Tibalt's Trickery, Pyroblast) that stray from it.
- PLAN theme_fit=yes: It is a Gruul Dragons Commander deck with a Dragon commander chosen by the builder, and the mix of a few strong staples with mostly midrange tribal cards and no fast combo sits comfortably at bracket 3.
- PLAN useful_as_built=yes: 36 lands with solid Gruul fixing plus a full ramp/cost-reduction package supports the admittedly heavy curve, and there are ample win conditions in the many large fliers and Atarka's double strike.
- PLAN summary_honest=yes: The claims of ramp, dragon-tribal removal, sweepers (Breath Weapon, Draconic Intervention, Steel Hellkite), protection (Veils, boots/greaves, Heroic Intervention) and card draw all match actual cards, and it openly concedes the lack of counterspells and speed.
- [INFO] `curve_summary`: average mana value 3.30 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Carnelian Orb of Dragonkind | ramp | Early mana development for casting the deck’s larger Dragons.
- 1 Jade Orb of Dragonkind | ramp | Early mana development that supports the green half of the deck.
- 1 Mox Jasper | ramp | A low-cost mana piece for accelerating Dragon turns.
- 1 Scaled Nurturer | ramp | Low-cost creature ramp that fits the Dragon theme.
- 1 Reckless Barbarian | ramp | Low-cost acceleration to help reach the top of the curve.
- 1 Embermouth Sentinel | ramp | Provides additional early mana development.
- 1 Dragon's Hoard | ramp | Dragon-focused mana development for the midgame.
- 1 Atsushi, the Blazing Sky | ramp | A Dragon ramp card that remains useful alongside the combat plan.
- 1 Ganax, Astral Hunter | ramp | Dragon-themed ramp for chaining larger creatures.
- 1 Goldspan Dragon | ramp | A threatening Dragon that also supplies ramp.
- 1 Faithless Looting | draw | Efficient early card selection and draw.
- 1 Thrill of Possibility | draw | Low-cost draw to keep hands moving.
- 1 Demand Answers | draw | Flexible early draw support.
- 1 Skullclamp | draw | Low-cost repeatable draw support.
- 1 Sensei's Divining Top | draw | Early card selection and ongoing draw support.
- 1 Sylvan Library | draw | Reliable early draw support.
- 1 Elemental Bond | draw | Draw support that rewards the deck’s large creatures.
- 1 Garruk's Uprising | draw | Draw support suited to a large-creature strategy.
- 1 Beast Whisperer | draw | Creature-focused draw for rebuilding and extending the board.
- 1 Dragonborn Champion | draw | Dragon-themed draw that supports combat pressure.
- 1 Guardian Project | draw | Long-game draw engine for a creature-heavy deck.
- 1 Heroic Intervention | interaction | Protects the board from opposing interaction.
- 1 Lightning Greaves | interaction | Protects a key creature while supporting immediate attacks.
- 1 Swiftfoot Boots | interaction | Additional protection for the commander and major threats.
- 1 Tamiyo's Safekeeping | interaction | Cheap protection for a crucial permanent.
- 1 Veil of Summer | interaction | Efficient defensive interaction.
- 1 Snakeskin Veil | interaction | Low-cost protection for a Dragon in combat or against removal.
- 1 Pyroblast | interaction | Efficient stack interaction.
- 1 Tibalt's Trickery | interaction | Broad reactive interaction when an answer is needed.
- 1 Dragon's Fire | removal | Efficient Dragon-themed creature removal.
- 1 Draconic Roar | removal | Low-cost removal that fits the Dragon shell.
- 1 Molten Exhale | removal | Direct removal for problematic creatures.
- 1 Piercing Exhale | removal | Additional cheap creature removal.
- 1 Spit Flame | removal | Dragon-focused removal support.
- 1 Dragon Tempest | removal | Dragon synergy that also supplies removal pressure.
- 1 Invasion of Tarkir // Defiant Thundermaw | removal | Dragon-themed removal that can become an additional threat.
- 1 Glorybringer | removal | A flying Dragon that provides repeatable removal pressure.
- 1 Terror of the Peaks | removal | A major Dragon threat that also adds removal reach.
- 1 Dragonlord's Servant | synergy | Cost support for the deck’s Dragon-heavy spell suite.
- 1 Dragonspeaker Shaman | synergy | Further cost support for deploying large Dragons.
- 1 Crucible of Fire | synergy | A tribal payoff for the Dragon combat plan.
- 1 Dragonkin Berserker | synergy | Low-cost Dragon-focused synergy piece.
- 1 Firespitter Whelp | synergy | Early Dragon synergy that contributes to the board.
- 1 Dragon Egg | synergy | Low-cost Dragon-themed board presence.
- 1 Dragon Hatchling | synergy | Early Dragon body that supports the tribal plan.
- 1 Minion of the Mighty | synergy | Cheap Dragon-focused synergy for explosive creature turns.
- 1 Sarkhan's Triumph | synergy | Finds a needed Dragon for the current game state.
- 1 Breaching Dragonstorm | synergy | A Dragon-themed payoff for developing the board.
- 1 Breath Weapon | wipe | Efficient sweep effect that fits a Dragon deck.
- 1 Draconic Intervention | wipe | Flexible sweeping interaction for creature-heavy boards.
- 1 Steel Hellkite | wipe | A Dragon threat with access to a sweeping effect.
- 1 Ancient Bronze Dragon | threat | High-impact Dragon threat for closing games through combat.
- 1 Backdraft Hellkite | threat | A substantial Dragon threat for the attacking plan.
- 1 Blast-Furnace Hellkite | threat | A powerful flying Dragon that increases combat pressure.
- 1 Dragon Broodmother | threat | A Dragon threat that develops the creature plan over time.
- 1 Hellkite Charger | threat | An aggressive Dragon threat for finishing games.
- 1 Lathliss, Dragon Queen | threat | A premier Dragon tribal threat and board-building payoff.
- 1 Scourge of the Throne | threat | A high-pressure Dragon attacker.
- 1 Terror of Mount Velus | threat | A large Dragon threat that strengthens combat finishes.
- 1 Thunderbreak Regent | threat | Efficient Dragon threat that supports the tribal board.
- 1 Thrakkus the Butcher | threat | Dragon-focused combat threat.
- 1 Twinflame Tyrant | threat | A Dragon threat that enhances the aggressive game plan.
- 1 Utvara Hellkite | threat | Top-end Dragon threat for ending longer games.
- 10 Mountain | land | Basic red source for the deck’s heavily red Dragon suite.
- 10 Forest | land | Basic green source for ramp and green Dragon cards.
- 1 Cavern of Souls | land | Tribal land that supplies colored mana for Dragons.
- 1 Cinder Glade | land | Red-green dual land for both halves of the deck.
- 1 City of Brass | land | Flexible colored source for consistent casting.
- 1 Command Tower | land | Reliable dual-color commander land.
- 1 Commercial District | land | Red-green dual land for the mana base.
- 1 Cragcrown Pathway // Timbercrown Pathway | land | Flexible early source for either deck color.
- 1 Exotic Orchard | land | Flexible colored land for multiplayer games.
- 1 Game Trail | land | Red-green dual land support.
- 1 Karplusan Forest | land | Untapped red-green source for early plays.
- 1 Mana Confluence | land | Reliable flexible colored source.
- 1 Rockfall Vale | land | Red-green dual land for stable development.
- 1 Rootbound Crag | land | Red-green dual land that supports an untapped mana base.
- 1 Secluded Courtyard | land | Tribal colored source for casting Dragons.
- 1 Spire Garden | land | Red-green multiplayer dual land.
- 1 Stomping Ground | land | Untapped red-green dual source.
- 1 Taiga | land | Untapped red-green dual source.

</details>

### 15. delegated commander, owned first

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 229 names.

Commander: Astarion, the Decadent.

Grade: bad, score 0.00, model 20260910T012734Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $147.28 the whole deck.

**Summary:** Astarion anchors a patient white-black lifegain strategy that develops mana, establishes a creature-and-equipment board, and turns its dedicated synergy package into steady pressure. Angels and other creature threats provide the primary route to closing games, while efficient spot answers and a small reset suite keep opposing boards manageable. The deck gives up explosive speed and relies on assembling its board over several turns, but it has strong staying power through plentiful mana, card access, and protection.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The commander does not place in cEDH events, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: The list consistently pairs lifegain triggers (Rosie Cotton, Angel of Vitality, Light of Promise, Shattered Angel, Exemplar of Light) with a commander that converts life into counters and a small equipment/protection suite to keep Astarion attacking, so the pieces pull the same way apart from a mild tension between three wraths and its own creature board.
- PLAN theme_fit=yes: It is a white-black lifegain Commander deck with the builder choosing a lifegain-themed commander and a low-power, no-combo, few-staples build appropriate to bracket 2, matching everything the request named that can be checked.
- PLAN useful_as_built=yes: 37 lands plus ten ramp pieces support a curve topping out around five, and there are plenty of creature threats, evasive angels, and removal to actually play and close games.
- PLAN summary_honest=yes: placeholder
- [WARN] `not_owned`: Astarion, the Decadent: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.08 over 62 nonland cards

<details><summary>The deck list</summary>

- 1 Buster Sword | draw | Provides a card-advantage slot while fitting the deck's artifact package.
- 1 Call of the Ring | draw | Provides repeatable card access for the lifegain plan.
- 1 Exemplar of Light | draw | Pairs card access with the deck's lifegain-oriented creature suite.
- 1 Grave Venerations | draw | Adds a card-advantage piece in black.
- 1 Idol of Oblivion | draw | Supplies an efficient artifact-based draw option.
- 1 Inspiring Overseer | draw | Adds card access on a creature that fits the white creature base.
- 1 Lembas | draw | Offers a compact artifact draw slot.
- 1 Mask of Memory | draw | Turns equipped attackers into a source of card access.
- 1 Night's Whisper | draw | Provides straightforward early card access.
- 1 Skullclamp | draw | Adds a low-cost equipment-based draw engine.
- 1 Tome of Legends | draw | Rewards repeated access to the commander with cards.
- 1 Bastion Protector | interaction | Helps preserve the commander and supports the creature plan.
- 1 Clever Concealment | interaction | Protects the developed board from opposing disruption.
- 1 Darksteel Plate | interaction | Gives a key creature a durable protective equipment slot.
- 1 Lightning Greaves | interaction | Protects an important creature at a low mana cost.
- 1 Swiftfoot Boots | interaction | Provides another efficient protection equipment option.
- 1 Unbreakable Formation | interaction | Protects the creature board while advancing combat pressure.
- 15 Plains | land | Provides dependable basic white mana.
- 14 Swamp | land | Provides dependable basic black mana.
- 1 City of Brass | land | Adds a flexible nonbasic mana slot.
- 1 Command Tower | land | Adds a reliable multicolor land slot.
- 1 Exotic Orchard | land | Adds flexible nonbasic mana support.
- 1 Grand Coliseum | land | Expands the mana base's color flexibility.
- 1 Path of Ancestry | land | Adds a multicolor land that supports the creature-heavy build.
- 1 Plaza of Heroes | land | Supports the legendary commander and creature base.
- 1 Secluded Courtyard | land | Adds creature-focused color fixing.
- 1 Unclaimed Territory | land | Adds another creature-focused color source.
- 1 Arcane Signet | ramp | Provides efficient early mana fixing.
- 1 Bender's Waterskin | ramp | Adds an artifact ramp piece to accelerate the deck.
- 1 Blitzball | ramp | Supplies another artifact mana source.
- 1 Commander's Sphere | ramp | Provides fixing early and utility later.
- 1 Fellwar Stone | ramp | Adds efficient two-mana acceleration.
- 1 Lotho, Corrupt Shirriff | ramp | Contributes to the deck's mana development on a creature.
- 1 Oath of the Grey Host | ramp | Adds a thematic mana-development piece.
- 1 Sol Ring | ramp | Provides one of the deck's limited fast mana slots.
- 1 Wayfarer's Bauble | ramp | Develops mana while supporting a land-focused opening.
- 1 White Auracite | ramp | Rounds out the artifact ramp suite.
- 1 Banishing Light | removal | Provides broad permanent-based removal.
- 1 Bitter Triumph | removal | Adds flexible black creature interaction.
- 1 Crib Swap | removal | Provides a creature-removal option.
- 1 Dispatch | removal | Adds efficient spot interaction alongside the artifacts.
- 1 Infernal Grasp | removal | Provides direct black creature removal.
- 1 Swords to Plowshares | removal | Gives the deck a premium low-cost removal option.
- 1 Stroke of Midnight | removal | Covers a wide range of opposing permanents.
- 1 Aerith Gainsborough | synergy | Supports the deck's central lifegain plan.
- 1 Angel of Vitality | synergy | Adds a lifegain-focused creature to the board.
- 1 Compassionate Healer | synergy | Strengthens the deck's dedicated lifegain package.
- 1 Doom Reigns Supreme | synergy | Provides a black synergy piece for the main plan.
- 1 Elixir | synergy | Adds a compact artifact synergy slot.
- 1 Kor Firewalker | synergy | Contributes another lifegain-oriented creature.
- 1 Light of Promise | synergy | Supports the deck's lifegain theme.
- 1 Night Nurse, Healer of Heroes | synergy | Adds a thematic legendary lifegain piece.
- 1 Prideful Feastling | synergy | Contributes to the dedicated lifegain package.
- 1 Rosie Cotton of South Lane | synergy | Adds a white creature that supports the main synergy plan.
- 1 Second Breakfast | synergy | Provides a low-cost lifegain synergy card.
- 1 Stone Docent | synergy | Adds another creature-based synergy piece.
- 1 White Mage's Staff | synergy | Rounds out the equipment-based synergy suite.
- 1 Angel of Invention | threat | Supplies an evasive top-end threat.
- 1 Bill the Pony | threat | Adds a creature threat that fits the white base.
- 1 Dawnhand Eulogist | threat | Provides a black threat for the board plan.
- 1 Foggy Swamp Hunters | threat | Adds another pressure piece in black.
- 1 Gollum, Silent Slinker // Meager Meal | threat | Supplies a thematic black threat with added flexibility.
- 1 Lyra Dawnbringer | threat | Adds a powerful Angel threat to close games.
- 1 Minwu, White Mage | threat | Provides a legendary white threat.
- 1 Rabaroo Troop | threat | Adds creature pressure to the battlefield plan.
- 1 Rooftop Percher | threat | Provides an additional creature threat.
- 1 Shattered Angel | threat | Adds an Angel threat that complements the creature suite.
- 1 Sneering Shadewriter | threat | Supplies another black creature threat.
- 1 Victory's Herald | threat | Adds a high-impact Angel finisher.
- 1 Austere Command | wipe | Provides flexible reset coverage when the board gets away.
- 1 Fumigate | wipe | Gives the deck a dependable full-board reset.
- 1 Vanquish the Horde | wipe | Adds an efficient creature-board reset.

</details>

### 16. a tight budget, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 229 names.

Commander: Denethor, Ruling Steward.

Grade: bad, score 0.01, model 20260910T012734Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $21.93 to buy, $45.67 the whole deck.

**Summary:** Denethor leads a black-white aristocrats deck that develops its mana, establishes a creature board, and turns sacrifice-oriented synergy into steady pressure while draw keeps the hand supplied. It wins by combining its aristocrats core with a broad suite of black creature threats, using targeted removal and board wipes to clear the way. The deck gives up premium-speed acceleration and expensive protection in exchange for a stable basic-land mana base, a low purchase cost, and a longer, board-focused game.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The sac outlets (Woe Strider, Bartolom	tps, Ayli, Yahenni, Aron, Woe Strider) and drain payoffs (Blood Artist, Zulaport Cutthroat, Elas il-Kor, Vengeful Bloodwitch, Meathook Massacre II) line up on one plan, but the deck is thin on token/fodder generation to feed those outlets and carries three symmetrical wipes plus equipment/protection pieces that pull against a wide creature board.
- PLAN theme_fit=yes: It is a black-white Commander aristocrats deck under Denethor built from cheap, commonly-owned staples at a power level consistent with bracket 2, matching the request as stated.
- PLAN useful_as_built=yes: 37 lands plus Sol Ring, Arcane Signet, Fellwar Stone and other cheap rocks support a low curve, and the drain payoffs plus a wide creature suite give clear, repeatable ways to close a game.
- PLAN summary_honest=yes: Every claim in the summary	tp aristocrats core, black creature threats, targeted removal, wipes, all-basic mana, low cost	tp is visible in the list, and it openly names its tradeoffs and weak spots rather than hiding them.
- [WARN] `not_owned`: Aron, Benalia's Ruin: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ayli, Eternal Pilgrim: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bartolomé del Presidio: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Blood Artist: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Elas il-Kor, Sadistic Pilgrim: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Woe Strider: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Yahenni, Undying Partisan: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Zulaport Cutthroat: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Erebos, Bleak-Hearted: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ghoulcaller Gisa: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Lord of the Forsaken: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Lord Skitter's Butcher: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Meathook Massacre II: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Old Flitterfang: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ruthless Lawbringer: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Shadowheart, Dark Justiciar: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Shilgengar, Sire of Famine: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Sivriss, Nightmare Speaker: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Smothering Abomination: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vengeful Bloodwitch: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.03 over 62 nonland cards
- [INFO] `mana_pass`: the builder moved 1 card of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 19 Plains | land | Provides reliable white mana without entering tapped.
- 18 Swamp | land | Provides reliable black mana without entering tapped.
- 1 Arcane Signet | ramp | Provides efficient color fixing and acceleration.
- 1 Chromatic Lantern | ramp | Adds mana acceleration and supports both deck colors.
- 1 Commander's Sphere | ramp | Provides flexible mana acceleration.
- 1 Deadly Dispute | ramp | Converts a disposable permanent into mana development.
- 1 Explorer's Scope | ramp | Offers early mana development from an equipped attacker.
- 1 Fellwar Stone | ramp | Provides inexpensive mana acceleration.
- 1 Inherited Envelope | ramp | Adds another affordable mana source.
- 1 Ring of the Lucii | ramp | Supplies mana acceleration for the midgame.
- 1 Sol Ring | ramp | Provides the deck's strongest early acceleration.
- 1 Thought Vessel | ramp | Adds colorless mana acceleration.
- 1 Cirith Ungol Patrol | draw | Supplies a black draw option for the deck.
- 1 Exemplar of Light | draw | Provides a creature-based draw option.
- 1 Folk Hero | draw | Provides ongoing card access for the creature plan.
- 1 Inspiring Overseer | draw | Adds a simple creature-based draw effect.
- 1 Mask of Memory | draw | Turns combat into continued card access.
- 1 Nasty End | draw | Provides efficient black card draw.
- 1 Painful Truths | draw | Refills the hand efficiently.
- 1 Rat King, Pale Piper | draw | Adds a black creature that supports card access.
- 1 Skullclamp | draw | Provides repeatable card draw for expendable creatures.
- 1 Stone of Erech | draw | Adds artifact-based card access.
- 1 Wall of Omens | draw | Provides early defense while replacing itself.
- 1 Bastion Protector | interaction | Helps protect the commander from opposing interaction.
- 1 Gift of Immortality | interaction | Protects an important creature through removal.
- 1 Reprieve | interaction | Provides a flexible answer to an opposing spell.
- 1 Swiftfoot Boots | interaction | Protects a key creature while keeping it active.
- 1 Take Up the Shield | interaction | Offers inexpensive protection in response to removal.
- 1 Unbreakable Formation | interaction | Protects the creature board from opposing disruption.
- 1 Claim the Precious | removal | Provides efficient black creature removal.
- 1 Crib Swap | removal | Answers a creature at instant speed.
- 1 Fatal Push | removal | Provides low-cost creature removal.
- 1 Fiend Hunter | removal | Adds creature-based removal to the board.
- 1 Infernal Grasp | removal | Answers a wide range of opposing creatures.
- 1 Stroke of Midnight | removal | Provides flexible removal for troublesome permanents.
- 1 Swords to Plowshares | removal | Provides efficient creature removal.
- 1 Austere Command | wipe | Provides a flexible reset against developed boards.
- 1 Dusk // Dawn | wipe | Offers a board reset that fits a creature-focused deck.
- 1 Vanquish the Horde | wipe | Provides an affordable creature-board reset.
- 1 Arcade Cabinet | synergy | Fills a dedicated synergy slot from the existing library.
- 1 Gollum the Abandoned | synergy | Supports the deck's sacrifice-oriented synergy plan.
- 1 Gollum, Patient Plotter | synergy | Supports the deck's sacrifice-oriented synergy plan.
- 1 Gríma Wormtongue | synergy | Adds an on-theme black synergy creature.
- 1 Phantom Train | synergy | Provides a colorless synergy permanent from the existing library.
- 1 Aron, Benalia's Ruin | synergy | Adds a low-cost creature for the aristocrats core.
- 1 Ayli, Eternal Pilgrim | synergy | Adds a sacrifice-focused creature to the core plan.
- 1 Bartolomé del Presidio | synergy | Provides another low-cost sacrifice-oriented creature.
- 1 Blood Artist | synergy | Adds a central aristocrats payoff to the deck.
- 1 Elas il-Kor, Sadistic Pilgrim | synergy | Adds a compact black-white aristocrats payoff.
- 1 Woe Strider | synergy | Adds a sacrifice-oriented creature to support the plan.
- 1 Yahenni, Undying Partisan | synergy | Adds another sacrifice-focused black creature.
- 1 Zulaport Cutthroat | synergy | Provides a focused aristocrats payoff.
- 1 Erebos, Bleak-Hearted | threat | Provides a resilient black threat for longer games.
- 1 Ghoulcaller Gisa | threat | Adds a substantial black creature threat.
- 1 Lord of the Forsaken | threat | Provides another black threat for the top end.
- 1 Lord Skitter's Butcher | threat | Adds a low-cost black threat.
- 1 Meathook Massacre II | threat | Provides an impactful black threat.
- 1 Old Flitterfang | threat | Adds a black creature threat that suits the deck's theme.
- 1 Ruthless Lawbringer | threat | Provides an inexpensive additional threat.
- 1 Shadowheart, Dark Justiciar | threat | Adds a black threat that complements the sacrifice plan.
- 1 Shilgengar, Sire of Famine | threat | Provides a larger black-white threat.
- 1 Sivriss, Nightmare Speaker | threat | Adds a low-cost black threat.
- 1 Smothering Abomination | threat | Provides a creature threat suited to an aristocrats shell.
- 1 Vengeful Bloodwitch | threat | Adds an inexpensive black threat to pressure opponents.

</details>

### 17. upgrade a precon, owned first

Format: Commander. Theme: goblins. Pool: owned_first. Shortlist: 196 names.

Commander: Zada, Hedron Grinder.

Grade: typical, score 0.05, model 20260910T012734Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $84.65 to buy, $256.60 the whole deck.

**Summary:** This Goblin Storm build develops a Goblin board, uses its draw and ramp package to set up explosive spell-heavy turns, and closes through Goblin pressure alongside storm and creature-focused payoffs. It keeps the precon’s aggressive tribal identity rather than broadening into a more general red deck, so its answers and resilience remain tied closely to that focused game plan.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: The list is a single coherent engine: Zada plus one-mana cantrips (Crimson Wisps, Expedite, Ancestral Anger, Fists of Flame, Wild Ride) fed by wide Goblin token makers, supported by rituals and treasure makers, and cashed in on storm payoffs like Grapeshot, Empty the Warrens, Haze of Rage, Past in Flames and Impact Tremors.
- PLAN theme_fit=yes: It is a mono-red Goblin storm Commander deck at a bracket-3 power level (Sol Ring and rituals but no fast-mana or tutor-heavy cEDH package, only incidental Snoop/Kiki lines), which matches the requested precon upgrade.
- PLAN useful_as_built=yes: 38 mono-red lands with a low curve, ample cheap token makers, card draw and multiple distinct win routes (Grapeshot/Haze of Rage, Impact Tremors, Krenko/Siege-Gang beatdown) make it immediately playable, if slightly land-heavy for a storm shell.
- PLAN summary_honest=yes: Every claim — Goblin board development, a draw/ramp package, spell-heavy explosive turns, storm and creature payoffs, and answers kept narrow and tribal — is backed by actual cards in the list, and it does not overstate interaction or tutoring the deck lacks.
- [INFO] `curve_summary`: average mana value 2.62 over 61 nonland cards
- [WARN] `profile_off_band`: the interaction count is 1, and bracket 3 wants 4 to 12
- [WARN] `two_card_combo`: bracket 3 allows no two-card infinite combo that needs four mana or less, and the deck holds Storm-Kiln Artist + Haze of Rage (speed 5, near two-card)
- [INFO] `basics_added`: the list was 2 cards short, so the builder added 2 basic lands

<details><summary>The deck list</summary>

- 1 Arena of Glory | land | Part of the precon mana base.
- 1 Castle Embereth | land | Part of the precon mana base.
- 1 Den of the Bugbear | land | Part of the precon mana base.
- 1 Forgotten Cave | land | Part of the precon mana base.
- 1 Fountainport | land | Part of the precon mana base.
- 1 Hidden Volcano | land | Part of the precon mana base.
- 1 Kher Keep | land | Part of the precon mana base.
- 28 Mountain | land | Basic land allocation for a consistent mana base.
- 1 Reliquary Tower | land | Part of the precon mana base.
- 1 Shinka, the Bloodsoaked Keep | land | Part of the precon mana base.
- 1 War Room | land | Part of the precon mana base.
- 1 Ancestors' Aid | synergy | Retained as a precon spell for the deck's storm-focused plan.
- 1 Ancestral Anger | draw | Adds a low-cost draw spell to support explosive turns.
- 1 Arcane Signet | ramp | Adds dependable early ramp.
- 1 Battle Hymn | ramp | Retained as a precon ramp spell.
- 1 Blasphemous Act | wipe | Retained as the precon's board-reset option.
- 1 Boggart Shenanigans | removal | Retained as a Goblin-themed removal piece.
- 1 Brightstone Ritual | ramp | Adds Goblin-focused ramp for larger turns.
- 1 Chaos Warp | removal | Retained as flexible precon removal.
- 1 Conspicuous Snoop | synergy | Retained as a Goblin synergy card.
- 1 Crimson Wisps | draw | Retained as a precon draw spell.
- 1 Cunning Maneuver | draw | Adds a low-cost draw spell for the deck's spell-heavy turns.
- 1 Daring Discovery | other | Retained as part of the precon's storm package.
- 1 Dragon Fodder | synergy | Retained as a precon card that supports the Goblin theme.
- 1 Empty the Warrens | wincon | Retained as a storm-themed closing card.
- 1 Expedite | draw | Retained as a precon draw spell.
- 1 Fable of the Mirror-Breaker // Reflection of Kiki-Jiki | ramp | Adds a strong ramp card while staying within the deck's creature-focused theme.
- 1 Faithless Looting | draw | Retained as a precon draw spell.
- 1 Fists of Flame | draw | Retained as a precon draw spell.
- 1 Frontline Heroism | synergy | Retained as part of the precon's creature-focused plan.
- 1 Gempalm Incinerator | removal | Retained as Goblin-themed removal.
- 1 General Kreat, the Boltbringer | synergy | Retained as a Goblin synergy card.
- 1 Glimpse the Impossible | ramp | Retained as a precon ramp spell.
- 1 Goblin Bombardment | removal | Retained as a Goblin-themed removal outlet.
- 1 Goblin Bushwhacker | synergy | Retained as a Goblin synergy card.
- 1 Goblin Chieftain | synergy | Retained as a Goblin synergy card.
- 1 Goblin Dark-Dwellers | threat | Retained as a Goblin threat.
- 1 Goblin Lackey | synergy | Retained as a Goblin synergy card.
- 1 Goblin Matron | synergy | Retained as a Goblin synergy card.
- 1 Goblin Negotiation | removal | Retained as precon removal.
- 1 Goblin Trashmaster | removal | Retained as Goblin-themed removal.
- 1 Goblin Warchief | synergy | Retained as a central Goblin synergy card.
- 1 Grapeshot | removal | Retained as a storm-themed removal spell.
- 1 Haze of Rage | wincon | Retained as a precon storm payoff.
- 1 Idol of Oblivion | draw | Retained as a precon draw card.
- 1 Impact Tremors | wincon | Retained as a creature-focused payoff.
- 1 Krenko's Command | synergy | Retained as a precon Goblin synergy card.
- 1 Krenko, Mob Boss | threat | Retained as a Goblin threat.
- 1 Mana Geyser | ramp | Retained as a precon ramp spell.
- 1 Mogg War Marshal | synergy | Retained as a Goblin synergy card.
- 1 Moria Marauder | synergy | Adds another Goblin synergy card without changing the precon's theme.
- 1 Pashalik Mons | removal | Adds Goblin-themed removal while supporting the existing creature plan.
- 1 Past in Flames | synergy | Retained as part of the precon's storm package.
- 1 Quest for the Goblin Lord | synergy | Retained as a Goblin synergy card.
- 1 Roaming Throne | synergy | Retained as part of the precon's Goblin-focused package.
- 1 Ruby Medallion | ramp | Retained as a precon ramp card.
- 1 Rundvelt Hordemaster | synergy | Adds a focused Goblin synergy upgrade.
- 1 Sazacap's Brew | draw | Retained as a precon draw spell.
- 1 Searslicer Goblin | synergy | Retained as a Goblin synergy card.
- 1 Seething Song | ramp | Retained as a precon ramp spell.
- 1 Siege-Gang Commander | removal | Retained as Goblin-themed removal.
- 1 Siege-Gang Lieutenant | removal | Retained as Goblin-themed removal.
- 1 Skirk Prospector | ramp | Retained as a precon ramp card.
- 1 Skullclamp | draw | Retained as a precon draw card.
- 1 Sol Ring | ramp | Retained as a precon ramp card.
- 1 Spreading Insurrection | wincon | Retained as a precon finishing card.
- 1 Storm-Kiln Artist | ramp | Retained as a precon ramp card.
- 1 Swiftfoot Boots | interaction | Retained as precon interaction.
- 1 Throne of Eldraine | ramp | Retained as a precon ramp card.
- 1 Vandalblast | wipe | Retained as a precon board-reset option.
- 1 Wild Ride | synergy | Retained as part of the precon's creature-focused plan.
- 1 Witch's Mark | draw | Retained as a precon draw spell.

</details>

### 18. upgrade a precon, any card

Format: Commander. Theme: turtles. Pool: any_card. Shortlist: 300 names.

Commander: Heroes in a Half Shell.

Grade: baseline, score 0.02, model 20260910T012734Z.

Cards: 99 main, 0 sideboard. Repair turn: yes, for precon_share, profile_off_band, profile_off_band, profile_off_band, profile_off_band, profile_off_band, profile_off_band. Block findings: 0.

Cost: $83.14 to buy, $83.14 the whole deck.

**Summary:** A five-color Turtle Power deck built around the heroes, their allies, their villains, and an energetic mix of gadgets, pizza, arcade flair, and mutant creatures. It develops its mana, builds a character-driven board, draws through its themed value pieces, and uses versatile answers and sweeping effects to keep the action moving before closing with a united team of oversized threats.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: Most cards feed a TMNT character/team-up board with a small +1/+1 counter sub-theme (Corpsejack Menace, Vigor, Biogenic Ooze, Level Up), but four sweepers (Blasphemous Act, Game Over, Vanquish the Horde, Wave Goodbye) work against a deck whose whole plan is amassing a wide character board.
- PLAN theme_fit=partly: It is unmistakably a five-color Turtle Power list on the right commander, but as an "upgrade" it lands at the precon baseline with few added staples, so the bracket-3 lift asked for is barely delivered.
- PLAN useful_as_built=partly: It has plenty of threats, draw, and removal and will function, but a five-color mana base of 23 basics with only Command Tower, Path of Ancestry, Grand Coliseum, Ash Barrens, and Chromatic Lantern for fixing will regularly strand multicolor cards and explains the below-norm turn-four mana.
- PLAN summary_honest=partly: The themed value, wipes, and character board are all genuinely present, but "develops its mana" oversells a five-color deck with only six ramp pieces and a mostly-basic land base.
- [INFO] `curve_summary`: average mana value 3.41 over 61 nonland cards
- [WARN] `profile_off_band`: the ramp count is 6, and bracket 3 wants 8 to 13
- [WARN] `profile_off_band`: the draw count is 5, and bracket 3 wants 8 to 14
- [INFO] `mana_pass`: the builder moved 3 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 1 Acidic Slime | removal | A Turtle-adjacent utility creature that answers troublesome permanents.
- 1 April O'Neil, Live on the Scene | draw | Provides thematic card advantage.
- 1 Arcade Cabinet | other | Supports the precon's playful artifact and game-themed package.
- 1 Arcane Signet | ramp | Reliable fixing for a five-color Turtle strategy.
- 1 Assassin's Trophy | removal | Flexible answer to any problematic permanent.
- 1 Baxter, Fly in the Ointment | other | Retains a signature Turtle Power character.
- 1 Bebop, Skull & Crossbones | other | Retains a signature Turtle Power character.
- 1 Big Mother Mouser | other | Keeps the deck's mutant-and-machine flavor intact.
- 1 Biogenic Ooze | threat | Adds a resilient board-building threat.
- 1 Blasphemous Act | wipe | Efficiently resets creature-heavy boards.
- 1 Casey Jones, Back Alley Brute | threat | Keeps the Turtle Power cast of allies present.
- 1 Chromatic Lantern | ramp | Improves fixing while accelerating the deck.
- 1 Coin of Mastery | other | Supports the precon's game-focused artifact package.
- 1 Continue? | interaction | Maintains the video-game themed suite.
- 1 Corpsejack Menace | synergy | Strengthens counter-based creature development.
- 1 Cultivate | ramp | Provides dependable land-based acceleration and fixing.
- 1 Dimension X Pizzasaur | other | Keeps the deck's distinctive Turtle Power flavor.
- 1 Donatello, the Brains | synergy | Retains a core Turtle legend and synergistic payoff.
- 1 Double Jump // Flying Kick | interaction | A flexible thematic combat trick.
- 1 Electric Seaweed | other | Preserves the precon's defensive creature package.
- 1 Endless Foot Assault | other | Supports the Foot Clan portion of the theme.
- 1 Everything Pizza | other | Maintains the deck's Food and pizza subtheme.
- 1 Exploding Barrel | removal | A thematic artifact answer with utility.
- 1 Farseek | ramp | Adds efficient early fixing and land acceleration.
- 1 Fast Forward | other | Keeps the game-themed spell package intact.
- 1 Foot Chopper | other | Provides thematic equipment utility.
- 1 Game Over | wipe | Contributes a thematic high-impact spell.
- 1 Harmonize | draw | Straightforward card advantage for a creature-focused deck.
- 1 Here Comes a New Hero! | draw | Keeps the heroic Turtle Power narrative and creature support.
- 1 High Score | draw | Supports the arcade-game flavor while providing value.
- 1 Irma, Part-Time Mutant | other | Retains a thematic supporting legend.
- 1 Krang, the All-Powerful | threat | A major thematic threat for the deck's villain package.
- 1 Leatherhead, Iron Gator | threat | Keeps the mutant creature theme broad and flavorful.
- 1 Leonardo, the Balance | threat | Retains a core Turtle legend and combat-oriented threat.
- 1 Lessons from Life | draw | Provides thematic value and card flow.
- 1 Level Up | synergy | Supports creature development within the theme.
- 1 Lita, Little Orphan Amphibian | synergy | Retains a thematic Turtle legend and synergy piece.
- 1 Michelangelo, the Heart | synergy | Retains a core Turtle legend and team payoff.
- 1 Mole Module | other | Maintains the precon's Vehicle and gadget flavor.
- 1 Mona Lisa, Science Geek | other | Retains a supporting thematic legend.
- 1 Ninja Pizza | other | Supports the deck's pizza-centered theme.
- 1 Raphael, the Muscle | threat | Retains a core Turtle legend and aggressive threat.
- 1 Rat King, Pale Piper | threat | Keeps the precon's villain cast intact.
- 1 Ray Fillet, Wave Warrior | other | Retains a thematic supporting character.
- 1 Roadkill Rodney | other | Supports the deck's Robot and villain flavor.
- 1 Rocksteady, Mutant Marauder | threat | Keeps the mutant villain package intact.
- 1 Shellshock | interaction | A thematic defensive interaction spell.
- 1 Shredder, Shadow Master | threat | Retains an iconic villain and thematic threat.
- 1 Sol Ring | ramp | Efficient acceleration for deploying the deck's legends and threats.
- 1 Special Move | interaction | Maintains the combat-trick portion of the theme.
- 1 Splinter, the Mentor | synergy | Retains a key mentor character for the Turtle team.
- 1 Steelbane Hydra | removal | Creature-based artifact and enchantment interaction.
- 1 Super Combo | synergy | A thematic payoff for the deck's action-oriented plan.
- 1 Swift Demise | removal | Thematic spot interaction against opposing threats.
- 1 Tempestra, Dame of Games | other | Retains the game-themed legendary supporting cast.
- 1 Together Forever | interaction | Helps preserve important creatures through removal.
- 1 Tokka & Rahzar, Unsupervised | ramp | A thematic creature that contributes mana development.
- 1 Vanquish the Horde | wipe | Provides an additional broad creature reset.
- 1 Vigor | synergy | Protects and enhances the deck's creature board.
- 1 Voracious Hydra | removal | A scalable creature that can remove smaller threats.
- 1 Wave Goodbye | wipe | Provides a thematic board-clearing effect.
- 1 Ash Barrens | land | Fixes basic-land colors without entering tapped.
- 1 Big Apple, 3 a.m. | land | A flavorful fixing land for the Turtle Power setting.
- 1 Cinder Glade | land | Contributes red and green fixing.
- 1 Command Tower | land | Reliable access to every color in the deck.
- 1 Escape Tunnel | land | Fixes basics and offers late-game utility.
- 1 Grand Coliseum | land | Provides broad color fixing.
- 1 Hidden Hideout | land | A thematic utility land.
- 1 Path of Ancestry | land | Fixes colors for the deck's heavily creature-based plan.
- 1 Rain-Slicked Copse | land | Provides blue and green fixing.
- 1 Rootbound Crag | land | Provides red and green fixing.
- 1 Sodden Verdure | land | Provides green and blue fixing.
- 1 Turtle Lair | land | A central thematic land for the Turtle team.
- 1 Undergrowth Stadium | land | Provides black and green fixing.
- 1 Vernal Fen | land | Provides black and green fixing.
- 1 Vibrant Cityscape | land | Broadens access to the deck's colors.
- 7 Plains | land | Basic white sources for the deck's multicolor requirements.
- 3 Island | land | Basic blue sources for the deck's multicolor requirements.
- 4 Swamp | land | Basic black sources for the deck's multicolor requirements.
- 4 Mountain | land | Basic red sources for the deck's multicolor requirements.
- 5 Forest | land | Basic green sources for the deck's multicolor requirements.

</details>

### 19. the Hobbit family, two colours

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 174 names.

Commander: Thranduil, the Elvenking.

Grade: bad, score 0.00, model 20260910T012734Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $496.35 to buy, $496.35 the whole deck.

**Summary:** Thranduil leads an Elf-heavy Sultai creature deck that develops its mana early, builds a broad board of themed creatures, and keeps cards flowing with a deep draw package. It wins by maintaining battlefield pressure with its creatures while clearing away key opposition through targeted removal and a few reset buttons. The deck gives up some raw speed for a more board-focused, thematic game plan, so it is strongest when its mana creatures and Elf presence remain established.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core is consistent — mana elves, Middle-earth creatures, and a large draw suite supporting a wide board — but three sweepers (Languish, Gnashing of Teeth, Raise the Palisade) plus edict effects work against a deck whose whole plan is keeping small Elves on the battlefield, and there is no strong payoff engine tying the Elf count to a win condition beyond generic beatdown.
- PLAN theme_fit=yes: It is a Commander deck led by the requested commander at a modest bracket-3 power level and is overwhelmingly built from Middle-earth/Hobbit-set cards, with only a handful of generic staples (Arcane Signet, Elvish Mystic, Languish, Night's Whisper) rounding it out.
- PLAN useful_as_built=partly: Curve, land count (36 plus dorks and a signet), and multiple win avenues are fine, but a three-color Sultai deck running essentially all basics with almost no dual or fixing lands and mostly green mana creatures will stumble on casting its blue and black spells on time.
- PLAN summary_honest=yes: The summary accurately describes an Elf-leaning Sultai creature deck with ramp, heavy card draw, targeted removal, and a couple of resets, and it honestly flags the deck's slower, board-dependent nature rather than overselling a combo or engine the list lacks.
- [INFO] `curve_summary`: average mana value 3.19 over 63 nonland cards
- [INFO] `mana_pass`: the builder moved 3 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 1 Elvish Visionary | draw | A green creature in the shortlist’s draw slot.
- 1 Fateful Discovery | draw | A dedicated draw card for keeping resources flowing.
- 1 Gandalf, Shadow's Foe | draw | A legendary draw creature that adds to the board.
- 1 Gollum, Riddle Master | draw | A draw creature that provides another permanent for the battlefield plan.
- 1 Hithlain Knots | draw | A compact instant-speed draw option.
- 1 Ithilien Kingfisher | draw | A draw creature for the deck’s creature-heavy approach.
- 1 Key to the Side-Door | draw | An artifact draw source that does not require another creature.
- 1 Lórien Revealed | draw | A blue draw spell that helps avoid running out of cards.
- 1 Night's Whisper | draw | An efficient black draw spell.
- 1 Palantír of Orthanc | draw | A persistent artifact draw card.
- 1 Uncover the Moon-Letters | draw | An enchantment draw source for longer games.
- 1 Bilbo's Ring | interaction | An artifact interaction piece.
- 1 Confusticate and Bebother | interaction | A blue instant for responding to opposing plays.
- 1 Elrond, Moon-Reader | interaction | A legendary Elf that fills an interaction slot.
- 1 Mithril Coat | interaction | An equipment interaction card for protecting a key creature.
- 1 My Precious // Allure of Power | interaction | A flexible legendary artifact interaction card.
- 1 Old Fat Spider Can't See Me | interaction | An interaction Saga that remains on the board.
- 1 Stern Scolding | interaction | A low-cost blue interaction spell.
- 1 The One Ring | interaction | A legendary artifact interaction piece.
- 1 Arcane Signet | ramp | Reliable two-mana color fixing.
- 1 Delighted Halfling | ramp | A one-mana creature accelerator.
- 1 Elven Chorus | ramp | An Elf-focused enchantment ramp card.
- 1 Elvish Archdruid | ramp | An Elf Druid that supports the deck’s mana development.
- 1 Elvish Mystic | ramp | A one-mana Elf mana creature.
- 1 Mox Amber | ramp | A zero-mana legendary artifact accelerator.
- 1 Necklace of Girion | ramp | A legendary artifact mana source.
- 1 Silvan Reveler | ramp | An Elf ramp creature for the tribal shell.
- 1 Wood Elves | ramp | An Elf ramp creature that develops the board.
- 1 Woodland Weavemaster | ramp | An Elf Druid ramp card.
- 1 Bilbo's Deadly Slice | removal | A dedicated removal instant.
- 1 Bitter Downfall | removal | A dedicated black removal spell.
- 1 Colossal Whale | removal | A blue creature occupying a removal slot.
- 1 Crude Bent Blade | removal | An artifact equipment removal option.
- 1 Enchanted River's Grasp | removal | An enchantment-based removal card.
- 1 Merciless Executioner | removal | A creature-based removal effect.
- 1 Orcish Bowmasters | removal | A low-cost creature removal card.
- 1 The Black Arrow | removal | A legendary equipment removal card.
- 1 Uneasy Partings | removal | An instant-speed removal option.
- 1 Gnashing of Teeth | wipe | A board-clearing spell.
- 1 Languish | wipe | A black mass-removal spell.
- 1 Raise the Palisade | wipe | A tribal board-reset option.
- 1 Arwen, Weaver of Hope | synergy | A legendary Elf that reinforces the deck’s Elf theme.
- 1 Celeborn the Wise | synergy | A legendary Elf for the creature-focused core.
- 1 Elven Raft-Steerer | synergy | An Elf creature supporting the tribal density.
- 1 Elvenking's Harper | synergy | An Elf creature that fits Thranduil’s court.
- 1 Galadhrim Guide | synergy | An Elf Scout that increases the Elf count.
- 1 Galion, Elvenking's Butler | synergy | A legendary Elf tied directly to the Elvenking’s court.
- 1 Lothlórien Lookout | synergy | An Elf Scout for the tribal creature base.
- 1 Mirkwood Meditator | synergy | An Elf Druid that supports the Elf-centered plan.
- 1 Mirkwood Nurturer | synergy | An Elf Ranger that adds to the themed board.
- 1 Supper for Spiders | synergy | The shortlist’s dedicated synergy card.
- 1 Attercop | threat | A Spider threat that adds battlefield pressure.
- 1 Boughside Wanderers | threat | An Elf Scout threat for the creature plan.
- 1 Cantankerous Keepers | threat | An Elf Soldier threat that adds to the board.
- 1 Dreaded Bat-Cloud | threat | A creature threat outside the Elf core.
- 1 Grey Havens Navigator | threat | An Elf Pilot threat that maintains tribal density.
- 1 Guardian of the Halls | threat | An Elf Soldier threat for sustained board pressure.
- 1 Haunt of the Dead Marshes | threat | A Nightmare Elf that provides another threat.
- 1 Mirkwood Elk | threat | A creature threat for attacking and board presence.
- 1 Mirkwood Pathmaker | threat | An Elf Ranger threat that supports the tribal board.
- 1 Nimrodel Watcher | threat | An Elf Scout threat for the creature curve.
- 1 Thranduil, Sindarin Liege // Silvan Rally | threat | A legendary Elf threat fitting the Elvenking theme.
- 1 Willow-Wind | threat | An Elemental threat that broadens the creature suite.
- 12 Forest | land | Basic green sources for the Elf-heavy mana base.
- 9 Island | land | Basic blue sources for draw and interaction.
- 9 Swamp | land | Basic black sources for removal and sweepers.
- 1 Elven Passage | land | A themed nonbasic land for the Elf deck.
- 1 Elvenking's Halls | land | A themed nonbasic land for Thranduil’s court.
- 1 Mirkwood | land | A thematic nonbasic land source.
- 1 Rivendell | land | A blue-aligned nonbasic land source.
- 1 The Black Gate | land | A black-aligned legendary land source.
- 1 Hobbit Hole | land | the mana pass added it to bring the mana base inside the power level

</details>

### 20. the Hobbit family, a delegated commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 77 names.

Commander: Smaug the Magnificent.

Grade: bad, score 0.00, model 20260910T012734Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $490.61 to buy, $490.61 the whole deck.

**Summary:** Smaug the Magnificent leads a red Hobbit-themed Dragon deck that accelerates into its commander, Dragon creatures, and other substantial threats. It wins through creature pressure backed by a deep package of targeted removal and sweeping effects, while its artifacts and legendary characters provide broad support. The deck gives up the depth of a dedicated Dragon-only synergy package in exchange for a wider collection of flavorful creatures, Equipment, and utility cards.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The commander does not place in cEDH events, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The list is a coherent mono-red goodstuff pile of creatures, Equipment and removal that can attack and burn, but it lacks a unifying engine — the Dragon payoff the commander suggests never materializes and the Equipment, legendary-matters, and goblin bits pull mildly in separate directions.
- PLAN theme_fit=no: The person asked for a dragons deck and got a Smaug-led pile with essentially two Dragon bodies amid dwarves, goblins, and Equipment, so the central request is unmet even though the set and bracket are right.
- PLAN useful_as_built=yes: 37 lands in a mono-red deck with ramp rocks, a reasonable curve, ample removal and enough evasive/large threats means it can be shuffled up and played as-is.
- PLAN summary_honest=partly: It is candid that the deck trades Dragon synergy for flavorful red creatures and Equipment, which is honest, but it still bills itself as a 'Dragon-themed' deck accelerating into 'Dragon creatures' when only a couple of Dragons are present.
- [INFO] `curve_summary`: average mana value 3.15 over 62 nonland cards
- [WARN] `profile_off_band`: the draw count is 6, and bracket 3 wants 8 to 14
- [INFO] `basics_added`: the list was 1 card short, so the builder added 1 basic land

<details><summary>The deck list</summary>

- 1 Andúril, Flame of the West | other | Flexible supporting Equipment from the shortlist.
- 1 Andúril, Narsil Reforged | other | Flexible supporting Equipment from the shortlist.
- 1 Arcane Signet | ramp | Reliable early mana development.
- 1 Balin, Loremaster | draw | A designated source of card draw.
- 1 Bilbo's Ring | interaction | A designated interaction piece.
- 1 Bombur, Gentle Dreamer | other | A supporting legendary creature for the Hobbit theme.
- 1 Bothersome Noisemaker | other | A supporting creature from the shortlist.
- 1 Burn, Burn, Tree and Fern | ramp | A designated ramp card.
- 1 Call Forth the Tempest | wipe | A dedicated board wipe.
- 1 Cavern-Hoard Dragon | ramp | A Dragon that also supports mana development.
- 1 Desert Were-Worm | threat | A primary Dragon threat.
- 1 Desolation of Smaug | wipe | A dedicated board wipe.
- 1 Dori, Bearer of Friends | other | A supporting legendary creature from the shortlist.
- 1 Dwarven Mattock | interaction | A designated interaction piece.
- 1 Dwarven Mauler | other | A supporting creature from the shortlist.
- 1 Dwarven Warriors | other | A supporting creature from the shortlist.
- 1 Dáin Ironfoot | other | A supporting legendary creature for the Hobbit theme.
- 1 Fire of Orthanc | removal | A dedicated removal spell.
- 1 Fíli and Kíli, Joyous | ramp | A designated ramp creature.
- 1 Gandalf, Goblins' Bane // Flameshape | other | A flexible supporting legendary card.
- 1 Getaway Barrel | other | A flexible supporting artifact.
- 1 Glóin the Mighty // Easy Pickings | wipe | A dedicated board wipe.
- 1 Goblin Cratermaker | removal | A dedicated removal creature.
- 1 Goblin Fireleaper | removal | A dedicated removal creature.
- 1 Goblin-town Flunkies | other | A supporting creature from the shortlist.
- 1 Gundabad Opportunist | other | A supporting creature from the shortlist.
- 1 Guttersnipe | other | A supporting creature from the shortlist.
- 1 Improvised Club | removal | A dedicated removal spell.
- 1 Inferno Titan | removal | A large creature that provides removal.
- 1 Iron Hills Stalwart | other | A supporting creature from the shortlist.
- 1 Key to the Side-Door | draw | A designated source of card draw.
- 1 Last Light of Durin's Day | synergy | The deck's dedicated synergy piece.
- 1 Long-Bodied Grey Dog | ramp | A designated ramp creature.
- 1 Long-Lost Lances | other | A flexible supporting Equipment.
- 1 Misty Mountains Raider | other | A supporting creature from the shortlist.
- 1 Mithril Coat | interaction | A designated interaction piece.
- 1 Mox Amber | ramp | An efficient ramp artifact.
- 1 Old Thrush | other | A supporting creature from the shortlist.
- 1 Oliphaunt | other | A supporting creature from the shortlist.
- 1 Olog-hai Crusher | other | A supporting creature from the shortlist.
- 1 Orcish Siegemaster | other | A supporting creature from the shortlist.
- 1 Orcrist, Goblin-cleaver | ramp | A designated ramp Equipment.
- 1 Palantír of Orthanc | draw | A designated source of card draw.
- 1 Pinecone Strike | removal | A dedicated removal spell.
- 1 Ragged Short Spear | draw | A designated source of card draw.
- 1 Smaug's Fury | other | A flexible supporting instant.
- 1 Smaug, the Great Calamity // Spew Flame | removal | A Dragon that supplies removal.
- 1 Smite the Deathless | removal | A dedicated removal spell.
- 1 Snowslope Hunter | other | A supporting creature from the shortlist.
- 1 Sting, Bilbo's Sword | other | A flexible supporting Equipment.
- 1 The Black Arrow | removal | A dedicated removal Equipment.
- 1 The Lonely Mountain | land | A flavorful land for the Smaug theme.
- 1 The Misty Mountains Cold | ramp | A designated ramp enchantment.
- 1 The One Ring | interaction | A designated interaction artifact.
- 1 The Reaver Cleaver | ramp | A designated ramp Equipment.
- 1 Thorin, Company's Leader | ramp | A designated ramp creature.
- 1 Thorin, Mountain-king | removal | A legendary creature that provides removal.
- 1 Thrór's Map | draw | A designated source of card draw.
- 1 Tidings of War | other | A flexible supporting sorcery.
- 1 Troop of Ponies | ramp | A designated ramp creature.
- 1 Wayfarer's Bauble | ramp | An efficient early ramp artifact.
- 1 Well-Worn Spatula | other | A flexible supporting Equipment.
- 1 Óin the Brave | draw | A designated source of card draw.
- 36 Mountain | land | Provides a consistent untapped red mana base.

</details>

### 21. the Hobbit family, mana from outside

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 156 names.

Commander: Smaug the Impenetrable.

Grade: bad, score 0.01, model 20260910T012734Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $923.24 to buy, $923.24 the whole deck.

**Summary:** This is a Smaug-led red-black pressure deck that develops its mana, keeps cards flowing, and builds a board around dragons, goblins, and equipment. It wins by forcing opponents to answer a steady stream of threats, with Smaug as a central finisher, while targeted removal and reset effects keep opposing boards from taking over. The deck gives up some reactive flexibility in favor of committing to its own board and themed support pieces.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The deck makes less mana on turn four than the norm, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The pieces are broadly aggressive red-black creatures, but the deck splits itself between dragons, goblins, and an equipment/Legendary-weapon subtheme with little that ties them together, so it plays as a goodstuff pile rather than one focused plan (which the summary itself half-admits).
- PLAN theme_fit=partly: The commander, colors, format, and roughly bracket-3 game-changer count are right, but the explicit "Hobbit set only, except rocks and lands" limit is broken by a sizable block of LTR and outside cards (The One Ring, Mithril Coat, Sauron the Lidless Eye, Andúril, Glamdring, Sting, Gollum the Abandoned, Goblin Cratermaker, and more).
- PLAN useful_as_built=partly: Thirty-seven lands with a solid dual/fetch base will cast spells and there are plenty of finishers, but only three real accelerants (Signet, Mox Amber, Wayfarer's Bauble) plus Temple of the False God under a top-heavy dragon curve means the deck often stalls before its payoffs land.
- PLAN summary_honest=yes: The summary accurately describes a scattered dragon/goblin/equipment board deck with removal and three sweepers, and it openly flags the weak early mana rather than overselling ramp or card flow.
- [INFO] `curve_summary`: average mana value 3.08 over 62 nonland cards
- [INFO] `mana_pass`: the builder moved 2 cards of the mana base to bring the deck inside its power level
- [INFO] `basics_added`: the list was 1 card short, so the builder added 1 basic land
- [WARN] `outside_requested_set`: the sets you named do not hold 18 cards: Ancient Tomb, Arid Mesa, Blood Crypt, Bloodstained Mire, Bojuka Bog, Command Tower, Dragonskull Summit, Evolving Wilds, Exotic Orchard, Fabled Passage, Marsh Flats, Path of Ancestry, Polluted Delta, Scalding Tarn, Temple of the False God, Terramorphic Expanse, Urborg, Tomb of Yawgmoth, Wooded Foothills

<details><summary>The deck list</summary>

- 1 Ancient Tomb | land | Land slot in the mana base.
- 1 Arid Mesa | land | Land slot in the mana base.
- 1 Blood Crypt | land | Land slot in the mana base.
- 1 Bloodstained Mire | land | Land slot in the mana base.
- 1 Bojuka Bog | land | Land slot in the mana base.
- 1 Command Tower | land | Land slot in the mana base.
- 1 Dragon-Cursed Halls | land | Land slot in the mana base.
- 1 Dragonskull Summit | land | Land slot in the mana base.
- 1 Evolving Wilds | land | Land slot in the mana base.
- 1 Exotic Orchard | land | Land slot in the mana base.
- 1 Fabled Passage | land | Land slot in the mana base.
- 1 Goblin-town | land | Land slot in the mana base.
- 1 Marsh Flats | land | Land slot in the mana base.
- 1 Minas Morgul, Dark Fortress | land | Land slot in the mana base.
- 1 Mount Doom | land | Land slot in the mana base.
- 1 Path of Ancestry | land | Land slot in the mana base.
- 1 Polluted Delta | land | Land slot in the mana base.
- 1 Scalding Tarn | land | Land slot in the mana base.
- 1 Temple of the False God | land | Land slot in the mana base.
- 1 Terramorphic Expanse | land | Land slot in the mana base.
- 1 The Black Gate | land | Land slot in the mana base.
- 1 The Lonely Mountain | land | Land slot in the mana base.
- 1 Treasure Vault | land | Land slot in the mana base.
- 1 Urborg, Tomb of Yawgmoth | land | Land slot in the mana base.
- 1 Wooded Foothills | land | Land slot in the mana base.
- 6 Mountain | land | Basic land slots that stabilize the mana base.
- 5 Swamp | land | Basic land slots that stabilize the mana base.
- 1 Balin, Loremaster | draw | Draw slot that keeps resources flowing.
- 1 Gollum, Riddle Master | draw | Draw slot that keeps resources flowing.
- 1 Key to the Side-Door | draw | Draw slot that keeps resources flowing.
- 1 Palantír of Orthanc | draw | Draw slot that keeps resources flowing.
- 1 Rage into the Valley | draw | Draw slot that keeps resources flowing.
- 1 Ragged Short Spear | draw | Draw slot that keeps resources flowing.
- 1 Reverent Howl | draw | Draw slot that keeps resources flowing.
- 1 The Sackville-Bagginses | draw | Draw slot that keeps resources flowing.
- 1 Thrór's Map | draw | Draw slot that keeps resources flowing.
- 1 Óin the Brave | draw | Draw slot that keeps resources flowing.
- 1 Bilbo's Ring | interaction | Interaction slot for opposing plans and key turns.
- 1 Dwarven Mattock | interaction | Interaction slot for opposing plans and key turns.
- 1 Mithril Coat | interaction | Interaction slot for opposing plans and key turns.
- 1 My Precious // Allure of Power | interaction | Interaction slot for opposing plans and key turns.
- 1 The One Ring | interaction | Interaction slot for opposing plans and key turns.
- 1 Arcane Signet | ramp | Efficient mana development for the early turns.
- 1 Bag End Banquet | ramp | Mana development that supports the main plan.
- 1 Burn, Burn, Tree and Fern | ramp | Mana development that supports the main plan.
- 1 Dragon's Desire | ramp | Mana development that supports the main plan.
- 1 Mox Amber | ramp | Fast mana for early development.
- 1 Orcrist, Goblin-cleaver | ramp | Ramp piece that supports the deck's board plan.
- 1 Smaug the Magnificent | ramp | Ramp piece that supports the dragon-focused plan.
- 1 The Reaver Cleaver | ramp | Ramp piece that supports the deck's board plan.
- 1 Wayfarer's Bauble | ramp | Early mana development.
- 1 Azog, Moria's Ruin | removal | Removal for problematic opposing permanents.
- 1 Battle-Scarred Goblin | removal | Removal that remains on theme.
- 1 Bilbo's Deadly Slice | removal | Removal for problematic opposing permanents.
- 1 Bitter Downfall | removal | Removal for problematic opposing permanents.
- 1 Bolg of the North | removal | Removal that remains on theme.
- 1 Crude Bent Blade | removal | Removal that supports the equipment package.
- 1 Fire of Orthanc | removal | Removal for problematic opposing permanents.
- 1 Giant's Boulder | removal | Removal that supports the board plan.
- 1 Goblin Cratermaker | removal | Removal that remains on theme.
- 1 Bolg, Erebor's Reckoning | wipe | Board-reset option when opponents get ahead.
- 1 Call Forth the Tempest | wipe | Board-reset option when opponents get ahead.
- 1 Desolation of Smaug | wipe | Board-reset option that fits the Smaug theme.
- 1 Andúril, Flame of the West | synergy | Equipment support for the deck's creature plan.
- 1 Andúril, Narsil Reforged | synergy | Equipment support for the deck's creature plan.
- 1 Down, Down to Goblin-town | synergy | Themed support for the deck's aggressive board plan.
- 1 Dwarven Mauler | synergy | Creature support for the deck's board plan.
- 1 Dwarven Warriors | synergy | Creature support for the deck's board plan.
- 1 Getaway Barrel | synergy | Support piece for the deck's central plan.
- 1 Glamdring | synergy | Equipment support for the deck's creature plan.
- 1 Goblin Plate Mail | synergy | Equipment support for the deck's creature plan.
- 1 Long-Lost Lances | synergy | Equipment support for the deck's creature plan.
- 1 Sting, Bilbo's Sword | synergy | Equipment support for the deck's creature plan.
- 1 Desert Were-Worm | threat | Threat that adds pressure to the board.
- 1 Desolation Prowler | threat | Threat that adds pressure to the board.
- 1 Dreaded Bat-Cloud | threat | Threat that adds pressure to the board.
- 1 Gandalf, Goblins' Bane // Flameshape | threat | Threat that contributes to the deck's pressure.
- 1 Great Goblin, Foul-Hearted | threat | Threat that adds pressure to the board.
- 1 Haunt of the Dead Marshes | threat | Threat that adds pressure to the board.
- 1 Nighthowl Pursuer | threat | Threat that adds pressure to the board.
- 1 Olog-hai Crusher | threat | Threat that adds pressure to the board.
- 1 Sauron, the Lidless Eye | threat | Threat that adds pressure to the board.
- 1 Smaug, the Great Calamity // Spew Flame | threat | Central dragon threat for the deck's finishing plan.
- 1 The Great Goblin | threat | Threat that adds pressure to the board.
- 1 Troll of Khazad-dûm | threat | Threat that adds pressure to the board.
- 1 Bothersome Noisemaker | other | Flexible themed creature slot.
- 1 Gollum the Abandoned | other | Themed creature slot for the board plan.
- 1 Tidings of War | other | Themed noncreature support slot.
- 1 Elven Passage | land | the mana pass added it to bring the mana base inside the power level
- 1 Bolg's Company | ramp | the mana pass added it to bring the mana base inside the power level

</details>

### 22. a set family and a card from outside it

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 175 names.

Commander: Thranduil, the Elvenking.

Grade: bad, score 0.00, model 20260910T012734Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $485.03 to buy, $485.03 the whole deck.

**Summary:** This deck develops its mana with Elves and artifacts, then establishes a creature-heavy Elvenking board backed by legendary cards. It keeps cards flowing while using removal, reactive spells, and broad resets to keep opposing boards manageable, then wins by maintaining enough creatures to apply steady combat pressure. It gives up some speed and redundancy for a thematic spread of creatures, artifacts, and signature Middle-earth locations.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: There is a loose Middle-earth creature-goodstuff shell with ramp, removal and card draw, but the pieces pull several ways at once — an Elf tribal sub-theme (Archdruid, Chorus, Harper), a legendary sub-theme (Mox Amber, Celeborn, Elrond), a pile of vanilla-ish bears and Mirkwood bodies, and three symmetric wipes that fight the creature-heavy plan the summary names.
- PLAN theme_fit=partly: It is a Commander deck under Thranduil with Sol Ring kept and a broadly Middle-earth feel, but the request named the Hobbit set specifically and a large share of the list is LTR-proper or entirely non-Middle-earth filler (Arcane Signet, Night's Whisper, Languish, Elvish Mystic, Colossal Whale, Mox Amber, Wood Elves).
- PLAN useful_as_built=yes: 36 lands with heavy basics in three colors plus early ramp support the curve, and there are enough creatures, removal and draw that the deck functions and can close games through combat, even if the wipes are awkward alongside its own board.
- PLAN summary_honest=partly: Ramp, draw, removal, wipes and the Middle-earth lands are all genuinely present as described, but the claim of a board 'backed by legendary cards' overstates a handful of scattered legends with no real legendary payoff beyond Mox Amber, and it does not mention that its own wipes conflict with the creature plan.
- [INFO] `curve_summary`: average mana value 3.30 over 63 nonland cards
- [WARN] `profile_off_band`: the worst color's share of the sources it needs is 0.76, and bracket 3 wants 0.85 or more (sources of requirement: U 14.75 of 19, B 20.75 of 26, G 15.25 of 20)
- [INFO] `mana_pass`: the builder moved 4 cards of the mana base to bring the deck inside its power level
- [WARN] `outside_requested_set`: the sets you named do not hold 1 card: Sol Ring

<details><summary>The deck list</summary>

- 10 Forest | land | Provides a reliable green mana base.
- 8 Island | land | Provides a reliable blue mana base.
- 13 Swamp | land | Provides a reliable black mana base.
- 1 Elven Passage | land | Adds an Elf-themed colored land slot.
- 1 Minas Morgul, Dark Fortress | land | Adds a black-aligned legendary land.
- 1 Mirkwood | land | Adds a thematic land for the creature base.
- 1 Rivendell | land | Adds a blue-aligned legendary land.
- 1 Sol Ring | ramp | The required kept card and an efficient mana source.
- 1 Arcane Signet | ramp | Provides flexible color fixing.
- 1 Bag End Banquet | ramp | Adds to the deck's mana development.
- 1 Delighted Halfling | ramp | Provides early creature-based mana.
- 1 Elven Chorus | ramp | Supports mana development in the Elf shell.
- 1 Elvish Archdruid | ramp | Rewards the deck's concentration of Elves with mana.
- 1 Elvish Mystic | ramp | Supplies early mana acceleration.
- 1 Mox Amber | ramp | Works alongside the deck's legendary creatures.
- 1 Wayfarer's Bauble | ramp | Provides dependable early land-based ramp.
- 1 Wood Elves | ramp | Combines creature presence with mana development.
- 1 Beorn the Fierce | draw | Provides card access from a substantial creature.
- 1 Elvish Visionary | draw | Offers a small Elf body with card access.
- 1 Hithlain Knots | draw | Adds an efficient card-access spell.
- 1 Ithilien Kingfisher | draw | Provides card access on a creature.
- 1 Key to the Side-Door | draw | Adds a compact source of card access.
- 1 Last March of the Ents | draw | Provides a larger card-access spell.
- 1 Lórien Revealed | draw | Adds flexible blue card access.
- 1 Night's Whisper | draw | Provides low-cost card access.
- 1 Palantír of Orthanc | draw | Adds repeatable card access from a legendary artifact.
- 1 Reverent Howl | draw | Provides instant-speed card access.
- 1 Uncover the Moon-Letters | draw | Adds an enchantment-based card-access source.
- 1 Bilbo's Deadly Slice | removal | Provides a direct answer to opposing permanents.
- 1 Bitter Downfall | removal | Adds an efficient removal spell.
- 1 Colossal Whale | removal | Puts removal on a large creature.
- 1 Crude Bent Blade | removal | Makes an Equipment slot serve as removal.
- 1 Enchanted River's Grasp | removal | Adds enchantment-based removal.
- 1 Giant's Boulder | removal | Provides a colorless removal option.
- 1 Orcish Bowmasters | removal | Adds removal attached to a creature.
- 1 Quarrel | removal | Provides a compact interaction piece for creatures.
- 1 Uneasy Partings | removal | Adds instant-speed removal.
- 1 Confusticate and Bebother | interaction | Provides reactive stack interaction.
- 1 Dwarven Mattock | interaction | Adds a reusable artifact interaction tool.
- 1 Elrond, Moon-Reader | interaction | Provides interaction on a legendary Elf.
- 1 Mithril Coat | interaction | Protects an important creature from opposing answers.
- 1 Old Fat Spider Can't See Me | interaction | Adds an interaction-focused Saga.
- 1 Stern Scolding | interaction | Provides efficient instant-speed interaction.
- 1 The One Ring | interaction | Adds a powerful legendary interaction piece.
- 1 Thranduil's Decree | interaction | Provides a thematic reactive spell.
- 1 Gnashing of Teeth | wipe | Provides a board-clearing option.
- 1 Languish | wipe | Adds a second broad creature reset.
- 1 Raise the Palisade | wipe | Provides a tribal-aware board reset.
- 1 Arwen, Weaver of Hope | synergy | Strengthens the deck's Elf and legendary creature core.
- 1 Boughside Wanderers | synergy | Adds to the Elf creature density.
- 1 Cantankerous Keepers | synergy | Adds another Elf body for the tribal plan.
- 1 Celeborn the Wise | synergy | Supports the legendary Elf-centered game plan.
- 1 Elven Raft-Steerer | synergy | Increases the deck's Elf presence.
- 1 Elvenking's Harper | synergy | Supports the Elvenking theme with an Elf creature.
- 1 Galadhrim Guide | synergy | Provides another thematic Elf body.
- 1 Galion, Elvenking's Butler | synergy | Directly reinforces the Elvenking flavor and legendary core.
- 1 Grey Havens Navigator | synergy | Adds another Elf for the creature shell.
- 1 Supper for Spiders | synergy | Provides a dedicated synergy piece for the creature plan.
- 1 Attercop | threat | Provides a large creature threat.
- 1 Gigantic Big Bear | threat | Adds a substantial green creature to pressure opponents.
- 1 Great Fierce Bee | threat | Provides another creature threat.
- 1 Large Bear | threat | Adds a straightforward board presence.
- 1 Little Bear | threat | Contributes an efficient creature body.
- 1 Mirkwood Elk | threat | Adds a thematic creature threat.
- 1 Mirkwood Meditator | threat | Provides an Elf threat for the main creature plan.
- 1 Mirkwood Nurturer | threat | Adds another Elf body that benefits the tribal shell.
- 1 Mirkwood Pathmaker | threat | Provides a further Elf threat.
- 1 Nasty Little Rabbit | threat | Adds a low-profile creature threat.
- 1 Ordinary Bear | threat | Provides an additional creature body.
- 1 Willow-Wind | threat | Adds a distinct creature threat at the top of the board.
- 1 Hobbit Hole | land | the mana pass added it to bring the mana base inside the power level

</details>

### 23. two set families at once

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 203 names.

Commander: Kíli the Resourceful.

Grade: baseline, score 0.03, model 20260910T012734Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $302.42 to buy, $302.42 the whole deck.

**Summary:** Kíli leads a white creature-and-artifact deck centered on dwarves, Equipment, and a steady board presence. The deck develops its mana, keeps cards flowing, and uses a broad suite of answers to protect its position while its creature threats and larger finishers apply pressure. It trades raw speed for a more resilient, board-focused game built from Hobbit and Bloomburrow cards.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The commander does not place in cEDH events, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: As a mono-white creature/artifact goodstuff pile with removal, ramp, and card draw it functions as one deck, but the stated dwarf-and-Equipment focus is thin (only a handful of dwarves and about five Equipment) and is diluted by unrelated Bloomburrow rabbit/mouse payoffs, Maskwood Nexus, and Psychosis Crawler pulling in different directions.
- PLAN theme_fit=no: The one hard constraint the player named was that the deck come from the Hobbit and Bloomburrow sets, and roughly half the nonland cards are from unrelated sets, so this is not the deck that was asked for despite the Hobbit commander.
- PLAN useful_as_built=yes: Mono-white with 36 lands plus rocks, a sensible curve, ample removal, wraths, card draw, and plenty of creature-based win paths — it can be shuffled up and played as is.
- PLAN summary_honest=no: It advertises a deck 'built from Hobbit and Bloomburrow cards' when a large share of the list (Sol Ring, Skullclamp, Swords to Plowshares, Sun Titan, Karn, Academy Manufactor, Skyclave Apparition, Dusk // Dawn, Jazal Goldmane, etc.) comes from elsewhere, and it overstates a dwarves-and-Equipment center the cards barely support.
- [INFO] `curve_summary`: average mana value 2.95 over 63 nonland cards

<details><summary>The deck list</summary>

- 21 Plains | land | Reliable white mana base.
- 1 Castle Ardenvale | land | White-producing utility land.
- 1 Command Tower | land | Reliable commander mana.
- 1 Elven Passage | land | Additional colored land option.
- 1 Exotic Orchard | land | Flexible colored land option.
- 1 Fabled Passage | land | Mana-base fixing land.
- 1 Hidden Grotto | land | Flexible land option.
- 1 Hobbit Hole | land | Thematic utility land.
- 1 Minas Tirith | land | White-producing legendary land.
- 1 Path of Ancestry | land | Thematic creature land.
- 1 Three Tree City | land | Creature-focused utility land.
- 1 Thriving Heath | land | Additional white source.
- 1 Uncharted Haven | land | Flexible colored land option.
- 1 Evolving Wilds | land | Mana-base fixing land.
- 1 Terramorphic Expanse | land | Mana-base fixing land.
- 1 Lupinflower Village | land | Thematic utility land.
- 1 Sol Ring | ramp | Efficient early mana acceleration.
- 1 Mox Amber | ramp | Efficient legendary-focused mana acceleration.
- 1 Ornithopter of Paradise | ramp | Creature-based mana fixing.
- 1 Arcane Signet | ramp | Reliable color fixing and acceleration.
- 1 Fellwar Stone | ramp | Low-cost mana acceleration.
- 1 Mind Stone | ramp | Early acceleration with later utility.
- 1 Patchwork Banner | ramp | Creature-theme mana support.
- 1 Wayfarer's Bauble | ramp | Early land-based acceleration.
- 1 Loyal Warhound | ramp | Creature-based land acceleration.
- 1 Season of the Burrow | ramp | Land-based mana development.
- 1 Belladonna Took | draw | Thematic draw engine.
- 1 Caretaker's Talent | draw | Repeatable draw support.
- 1 Circuit Mender | draw | Artifact creature that replaces itself.
- 1 Cut a Deal | draw | Efficient draw spell.
- 1 Idol of Oblivion | draw | Low-cost artifact draw outlet.
- 1 Inspiring Overseer | draw | Creature-based card advantage.
- 1 Mentor of the Meek | draw | Creature-focused draw support.
- 1 Skullclamp | draw | Efficient Equipment-based draw.
- 1 Spirited Companion | draw | Low-cost creature cantrip.
- 1 The Gaffer | draw | Thematic draw engine.
- 1 Thrór's Map | draw | Thematic artifact card advantage.
- 1 Baird, Steward of Argive | interaction | Creature-based interaction piece.
- 1 Bilbo's Gambit | interaction | Flexible instant-speed interaction.
- 1 Bofur, Reliable Guardian // Concerted Care | interaction | Thematic protective interaction.
- 1 Crumb and Get It | interaction | Low-cost instant interaction.
- 1 Dawn's Truce | interaction | Protective interaction for the board.
- 1 Galadriel's Dismissal | interaction | Flexible instant-speed interaction.
- 1 Mithril Coat | interaction | Equipment-based protection.
- 1 Reprieve | interaction | Efficient tempo interaction.
- 1 Angel of the Ruins | removal | Artifact creature with removal utility.
- 1 Banishing Light | removal | Versatile permanent answer.
- 1 Fiend Hunter | removal | Creature-based removal.
- 1 Generous Gift | removal | Flexible permanent removal.
- 1 Hanged Executioner | removal | Creature-based removal option.
- 1 Loran of the Third Path | removal | Thematic artifact creature removal.
- 1 Skyclave Apparition | removal | Efficient creature-based removal.
- 1 Swords to Plowshares | removal | Efficient single-target removal.
- 1 The Black Arrow | removal | Thematic Equipment removal.
- 1 Dusk // Dawn | wipe | Flexible creature-board reset.
- 1 Martial Coup | wipe | Board reset that also leaves material behind.
- 1 Settle the Wreckage | wipe | Instant-speed board control.
- 1 Academy Manufactor | synergy | Artifact-based synergy piece.
- 1 Andúril, Flame of the West | synergy | Thematic Equipment for the artifact core.
- 1 Blade Splicer | synergy | Creature-and-artifact synergy piece.
- 1 Carrot Cake | synergy | Thematic artifact synergy piece.
- 1 Dáin, Lord of the Iron Hills | synergy | Dwarf-focused synergy piece.
- 1 Helm of the Host | synergy | High-impact Equipment synergy piece.
- 1 Iron Hills Blacksmith | synergy | Dwarf artifact-synergy creature.
- 1 Maskwood Nexus | synergy | Creature-type synergy support.
- 1 Ori, Keeper of Songs | synergy | Dwarf-focused synergy piece.
- 1 Tangle Tumbler | synergy | Artifact synergy piece.
- 1 Fíli the Pathfinder | threat | Thematic legendary creature threat.
- 1 Karn, the Great Creator | threat | Planeswalker threat for the artifact shell.
- 1 Brightblade Stoat | threat | Low-cost creature threat.
- 1 Flowerfoot Swordmaster | threat | Thematic creature threat.
- 1 Jazal Goldmane | threat | Legendary creature finisher.
- 1 Psychosis Crawler | threat | Artifact creature threat.
- 1 Realm-Cloaked Giant // Cast Off | threat | Large creature threat with flexibility.
- 1 Serra Redeemer | threat | Creature threat for the board-focused plan.
- 1 Steelburr Champion | threat | Thematic creature threat.
- 1 Sun Titan | threat | Large creature threat with value potential.
- 1 Sunscorch Regent | threat | Large flying creature threat.
- 1 Warren Warleader | threat | Thematic creature finisher.

</details>

### 24. a 60-card deck from one set

Format: Modern. Theme: aggro. Pool: any_card. Shortlist: 96 names.

Grade: bad, score 0.30, model 20260910T012734Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $38.40 to buy, $38.40 the whole deck.

**Summary:** This deck presses forward with a dense lineup of red creature threats, led by Mouse and Lizard synergy, and uses removal to keep attacks moving through resistance. Card draw helps it continue committing pressure after the opening turns, while Dragonhawk, Fate's Tempest provides a more substantial threat near the top of the curve. It gives up broad answers and defensive play in exchange for a focused, creature-first attack plan.

The quality model grades this deck below the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core is a coherent one-plan red aggro curve, but the removal package is a scattershot pile of six different one-ofs including four- and five-mana burn that fights the low curve, and the tribal payoffs are split between Mice and Lizards so Patchwork Banner and Reptilian Recruiter only reward half the board.
- PLAN theme_fit=yes: Every nonland card is from Bloomburrow, it is mono-red, creature-based aggro, and the power level is squarely casual Modern as asked.
- PLAN useful_as_built=yes: Sixty cards with 24 Mountains and a mostly one-to-three-drop curve means it always casts its spells and has plenty of creatures to close games, even if 24 lands is a touch greedy for this curve.
- PLAN summary_honest=partly: The creature-first plan, card draw, and Dragonhawk top end are all genuinely present, but calling it 'Mouse and Lizard synergy' oversells a split tribal base whose anthem effects can only name one type, and the summary glosses the clunky, expensive removal as 'keeping attacks moving.'
- [INFO] `curve_summary`: average mana value 3.22 over 36 nonland cards

<details><summary>The deck list</summary>

- 24 Mountain | land | Provides reliable red mana for the deck's all-red plan.
- 2 Patchwork Banner | ramp | Supplies ramp while fitting the creature-focused build.
- 2 Artist's Talent | draw | Provides card draw without moving away from the red core.
- 2 Might of the Meek | draw | A low-cost draw option that supports an aggressive game plan.
- 2 Sazacap's Brew | draw | Adds more card draw to help sustain pressure.
- 1 Agate Assault | removal | A red removal option for clearing opposing obstacles.
- 1 Blooming Blast | removal | Adds another removal effect from the Bloomburrow shortlist.
- 1 Conduct Electricity | removal | Helps answer opposing permanents that impede attacks.
- 1 Flame Lash | removal | Provides efficient removal alongside the creature assault.
- 1 Playful Shove | removal | Gives the deck another direct removal tool.
- 1 Rabid Gnaw | removal | Rounds out the removal suite with a red instant.
- 4 Emberheart Challenger | synergy | A Mouse Warrior that supplies core aggressive synergy.
- 4 Hearthborn Battler | synergy | A Lizard Warlock that deepens the deck's creature synergy.
- 4 Frilled Sparkshooter | threat | A red creature threat that contributes to the attacking plan.
- 4 Reptilian Recruiter | threat | A Lizard Warrior threat that keeps the creature count high.
- 4 Teapot Slinger | threat | A Raccoon Warrior threat for maintaining board pressure.
- 2 Dragonhawk, Fate's Tempest | threat | A larger threat that gives the deck a stronger top end.

</details>

### 25. use no card of an owned precon

Format: Commander. Theme: Avengers superheroes. Pool: owned_first. Shortlist: 152 names.

Commander: Syr Konrad, the Grim.

Grade: bad, score 0.01, model 20260910T012734Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $185.95 the whole deck.

**Summary:** This Syr Konrad deck plays a black attrition game: develop mana through artifacts, keep cards flowing with equipment and black draw spells, and use extensive removal plus sweepers to keep the table from establishing a board. It aims to close through Syr Konrad after opponents have been worn down and resources have traded away. The deck gives up proactive creature-heavy pressure and broad color flexibility in exchange for a consistent, resilient mono-black control plan.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The deck draws more than the norm, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The removal/sweeper shell does feed Syr Konrad death triggers coherently, but the pile of equipment and aura-like gear (Skullclamp, Mask of Memory, Buster Sword, Champion's Helm, Greaves, Boots, Crown of Gondor) is stranded in a deck with barely half a dozen creatures, so a whole sub-package pulls against the creature-light control plan.
- PLAN theme_fit=no: The player asked for an Avengers superheroes deck and got a generic mono-black Syr Konrad control pile whose flavor cards are Final Fantasy, Lord of the Rings, and Avatar, with nothing Marvel or superhero about it.
- PLAN useful_as_built=partly: 36 Swamps plus heavy ramp and premium removal will function turn to turn, but with only a handful of creatures the win condition rests almost entirely on Syr Konrad pings that the deck rarely triggers on its own, and many equipment slots are effectively blank cards.
- PLAN summary_honest=partly: The control-and-removal description matches the list, but it advertises "keep cards flowing with equipment" when the equipment has almost no creatures to attach to, and it never acknowledges that the deck ignores the requested Avengers theme entirely.
- [WARN] `not_owned`: Syr Konrad, the Grim: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.05 over 63 nonland cards
- [WARN] `profile_off_band`: the ramp count is 15, and bracket 3 wants 8 to 13
- [WARN] `profile_off_band`: the draw count is 16, and bracket 3 wants 8 to 14
- [WARN] `profile_off_band`: the removal count is 21, and bracket 3 wants 6 to 12

<details><summary>The deck list</summary>

- 36 Swamp | land | Provides a deep, untapped black mana base for the deck.
- 1 Arcane Signet | ramp | Efficient mana development for the early turns.
- 1 Astral Cornucopia | ramp | Adds to the deck's artifact-based mana development.
- 1 Bender's Waterskin | ramp | Fills a ramp slot while supporting the artifact package.
- 1 Blitzball | ramp | Provides another mana-development piece.
- 1 Brass Infiniscope | ramp | Contributes to the deck's ramp density.
- 1 Chromatic Lantern | ramp | A mana-development artifact that keeps the deck accelerating.
- 1 Coin of Mastery | ramp | Adds reliable ramp to reach the commander and larger spells.
- 1 Commander's Sphere | ramp | A flexible artifact ramp piece.
- 1 Deadly Dispute | ramp | Provides mana development while fitting the black value plan.
- 1 Explorer's Scope | ramp | An equipment-based ramp option for the artifact shell.
- 1 Inherited Envelope | ramp | Adds another ramp piece without straining the curve.
- 1 Inspiring Statuary | ramp | Supports the deck's numerous artifacts while developing mana.
- 1 Sol Ring | ramp | A highly efficient early ramp artifact.
- 1 Thought Vessel | ramp | Provides artifact mana development.
- 1 Wayfarer's Bauble | ramp | Helps develop mana early in the game.
- 1 Buster Sword | draw | Supplies card advantage from an equipment slot.
- 1 Call of the Ring | draw | A dedicated black card-advantage piece.
- 1 Crown of Gondor | draw | Adds card advantage while remaining part of the artifact package.
- 1 Dusk Urchins | draw | A creature-based source of card advantage.
- 1 Eventide's Shadow | draw | Provides another dedicated draw effect.
- 1 Grave Venerations | draw | Supports the deck's steady card flow.
- 1 Hoarder's Greed | draw | A black draw spell for refilling after trading resources.
- 1 Idol of Oblivion | draw | An artifact source of card advantage.
- 1 Instant Ramen | draw | A compact artifact draw option.
- 1 Lembas | draw | Provides card advantage from a low-commitment artifact.
- 1 Mask of Memory | draw | An equipment-based source of continued card advantage.
- 1 Massacre Girl, Known Killer | draw | Adds card advantage on a black permanent.
- 1 Mirror of Galadriel | draw | A durable artifact draw source.
- 1 Nasty End | draw | An efficient black draw spell.
- 1 Night's Whisper | draw | A straightforward early card-advantage spell.
- 1 Skullclamp | draw | An efficient equipment-based draw engine.
- 1 Champion's Helm | interaction | An equipment slot for protecting important permanents.
- 1 Darksteel Plate | interaction | Provides a resilient protective interaction piece.
- 1 Lightning Greaves | interaction | Protects key creatures in the deck's main game plan.
- 1 Orcish Medicine | interaction | A flexible black interaction slot.
- 1 Swiftfoot Boots | interaction | Additional equipment-based protection for key creatures.
- 1 The Walls of Ba Sing Se | interaction | A permanent-based interaction option.
- 1 Bitter Downfall | removal | A dedicated black removal spell.
- 1 Bitter Triumph | removal | Efficient removal for troublesome opposing cards.
- 1 Blight Rot | removal | Adds another black removal option.
- 1 Blowfly Infestation | removal | A removal piece that supports an attrition plan.
- 1 Contagion Clasp | removal | Artifact-based removal that fits the deck's permanent suite.
- 1 Cornered by Black Mages | removal | A black removal spell for the control package.
- 1 Dismember | removal | Low-cost removal for early opposing threats.
- 1 Epic Downfall | removal | A targeted removal spell.
- 1 Fatal Push | removal | Efficient early interaction against opposing creatures.
- 1 Grim Poppet | removal | A permanent-based removal tool.
- 1 Heartless Act | removal | Flexible instant-speed removal.
- 1 Incremental Blight | removal | A removal spell that helps break up developed boards.
- 1 Infernal Grasp | removal | Reliable targeted black removal.
- 1 Lash of the Balrog | removal | Adds another direct removal effect.
- 1 Liliana, Death Wielder | removal | A repeatable permanent-based removal option.
- 1 Meteor Sword | removal | Equipment-based removal within the artifact shell.
- 1 Nameless Inversion | removal | A compact instant-speed removal spell.
- 1 Orcish Bowmasters | removal | A creature-based removal piece.
- 1 Overkill | removal | Additional instant-speed removal.
- 1 Skinrender | removal | A creature that contributes to the removal suite.
- 1 Voracious Fell Beast | removal | A black creature-based removal option.
- 1 Archfiend of Ifnir | wipe | A creature-based board-clearing option.
- 1 Black Sun's Zenith | wipe | A scalable black board wipe.
- 1 Carnifex Demon | wipe | A permanent-based sweeper for creature-heavy boards.
- 1 Extinction Event | wipe | A black reset button against opposing boards.
- 1 Midnight Banshee | wipe | Provides another creature-based board wipe.

</details>

