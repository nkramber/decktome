# PR-8 deck gate

Run date: 2026-09-16. Card snapshot: 2026-09-04.

Verdict: FAIL. 23 of 25 decks passed every block check, 0 invented names reached the user, and 1 summaries stated a false rule of the game. All three bars are zero tolerance.

## Summary

| Measure | Value |
|---|---|
| Prompts | 25 |
| Decks returned | 25 |
| Decks with no block finding | 23 |
| Invented names that reached the user | 0 |
| Decks that needed the repair turn | 4 |
| Summaries judged (F-26) | 25 |
| Summaries that state a rule of the game | 1 |
| Summaries that state a FALSE rule | 1 |
| Judge errors | 0 |
| Case assertions missed (PR-28b) | 0 |
| Decks the plan judge read (PR-15, information) | 25 |
| Mean plan score, 0 to 1 | 0.72 |
| Plan reasons the judge left empty | 2 |
| Errors | 0 |
| Prompt version | 14 |
| Calls | 79 |
| Cost | $2.7876 |
| Time | 2320 seconds |

## Run

- Suite `decks`, run `pr8-deck-gate-run24`, on 2026-09-16, commit `6ff0244`.
- Roles: generate on `gpt-5.6-terra` (openai, effort medium), judge on `claude-opus-5` (anthropic, effort medium), repair on `gpt-5.6-terra` (openai, effort medium).
- Versions: card snapshot 2026-09-04, generate prompt version 14, plan_rubric prompt version 2, precons `5.3.0+20260903`, quality_model `20260914T154223Z`.
- Calls: 79. Cost: $2.7876. Time: 2320 seconds.

## Findings by code

A block stops the deck. A warning and a note are reports.

| Code | Count |
|---|---|
| `curve_summary` | 24 |
| `not_owned` | 18 |
| `profile_off_band` | 15 |
| `mana_pass` | 7 |
| `bracket_cut` | 2 |
| `cards_trimmed` | 2 |
| `outside_requested_set` | 2 |
| `deck_size` | 2 |
| `land_count` | 1 |
| `precon_cards_restored` | 1 |
| `game_changer_limit` | 1 |

By severity: BLOCK 3. WARN 36. INFO 36. 

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

Cost: $837.67 to buy, $837.67 the whole deck.

**Summary:** This Karlov of the Ghost Council lifegain deck establishes its dedicated synergy pieces, develops pressure through a broad creature threat package, and keeps the table in check with removal, protection, and board wipes. It closes through Angel of Destiny, Felidar Sovereign, and Test of Endurance, with several additional finisher-marked cards supporting the endgame. The tradeoff is a focused lifegain plan that prioritizes its internal synergies over a more general-purpose strategy.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every card feeds one W/B lifegain axis — life-trigger payoffs (Ajani's Pridemate, Bloodthirsty Aerialist, Archangel of Thune, Heliod), lifelink/doubling, life-based draw, and life-total wincons like Felidar Sovereign and Test of Endurance — with only the two symmetrical wraths mildly at odds with its own creature count.
- PLAN theme_fit=yes: It is exactly the requested Karlov-led Commander lifegain deck in Orzhov colors, with a power level (a couple of strong staples but no fast mana or tutor chains) that sits comfortably at bracket 3.
- PLAN useful_as_built=yes: 38 lands with a solid dual/painland base for a two-color deck, several mana rocks, a low-to-mid curve, and multiple redundant win routes mean it can be picked up and played, even if turn-four mana is a touch slow.
- PLAN summary_honest=yes: The summary's claims — synergy pieces, creature threats, removal/protection/wipes, and the named finishers plus extras like Aetherflux Reservoir and Vito — all appear in the list, and it admits the inward-facing nature of the plan.
- [INFO] `curve_summary`: average mana value 3.46 over 61 nonland cards
- [INFO] `bracket_cut`: to hold bracket 3, the builder cut 1 card: Enduring Angel // Angelic Enforcer (the combo Enduring Angel // Angelic Enforcer + Enduring Tenacity), and added 1 basic land
- [INFO] `mana_pass`: the builder moved 1 card of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 13 Plains | land | A basic white land slot for the mana base.
- 11 Swamp | land | A basic black land slot for the mana base.
- 1 Brightclimb Pathway // Grimclimb Pathway | land | A white-black land slot for the mana base.
- 1 Caves of Koilos | land | A white-black land slot for the mana base.
- 1 Command Tower | land | A land slot that supports the commander's colors.
- 1 Concealed Courtyard | land | A white-black land slot for the mana base.
- 1 Godless Shrine | land | A white-black land slot for the mana base.
- 1 Isolated Chapel | land | A white-black land slot for the mana base.
- 1 Marsh Flats | land | A land slot for the mana base.
- 1 Prismatic Vista | land | A land slot for the mana base.
- 1 Restless Fortress | land | A white-black land slot that also carries a finisher marking.
- 1 Scrubland | land | A white-black land slot for the mana base.
- 1 Shattered Sanctum | land | A white-black land slot for the mana base.
- 1 Shineshadow Snarl | land | A white-black land slot for the mana base.
- 1 Silent Clearing | land | A white-black land slot for the mana base.
- 1 Vault of Champions | land | A white-black land slot for the mana base.
- 1 Archivist of Oghma | draw | Provides a designated draw card.
- 1 Dawn of Hope | draw | Provides a designated draw card.
- 1 Cosmos Elixir | draw | Provides a designated draw card.
- 1 Enduring Innocence | draw | Provides a designated draw card.
- 1 Exemplar of Light | draw | Provides a designated draw card.
- 1 Mangara, the Diplomat | draw | Provides a designated draw card.
- 1 Markov Purifier | draw | Provides a designated draw card.
- 1 Sigarda's Splendor | draw | Provides a designated draw card.
- 1 The Gaffer | draw | Provides a designated draw card.
- 1 Tymna the Weaver | draw | Provides a designated draw card.
- 1 Well of Lost Dreams | draw | Provides a designated draw card.
- 1 Alseid of Life's Bounty | interaction | Provides a designated interaction card.
- 1 Courageous Resolve | interaction | Provides a designated interaction card.
- 1 Faith's Shield | interaction | Provides a designated interaction card.
- 1 Metropolis Reformer | interaction | Provides a designated interaction card.
- 1 Restoration Magic | interaction | Provides a designated interaction card.
- 1 Sword of Light and Shadow | interaction | Provides a designated interaction card.
- 1 Werefox Bodyguard | interaction | Provides a designated interaction card.
- 1 Altar of the Pantheon | ramp | Provides a designated ramp card.
- 1 Battle Angels of Tyr | ramp | Provides a designated ramp card.
- 1 Beza, the Bounding Spring | ramp | Provides a designated ramp card.
- 1 Crypt Ghast | ramp | Provides ramp and carries a finisher marking.
- 1 Nuka-Cola Vending Machine | ramp | Provides a designated ramp card.
- 1 Oasis Gardener | ramp | Provides a designated ramp card.
- 1 Orazca Relic | ramp | Provides a designated ramp card.
- 1 Potioner's Trove | ramp | Provides a designated ramp card.
- 1 Pristine Talisman | ramp | Provides a designated ramp card.
- 1 The Celestus | ramp | Provides a designated ramp card.
- 1 Aetherflux Reservoir | removal | Provides a designated removal card for the lifegain shell.
- 1 Ayli, Eternal Pilgrim | removal | Provides a designated removal card.
- 1 Gumdrop Poisoner // Tempt with Treats | removal | Provides a designated removal card.
- 1 Murderous Rider // Swift End | removal | Provides a designated removal card.
- 1 Nightmare's Thirst | removal | Provides a designated removal card.
- 1 Solitude | removal | Provides a designated removal card.
- 1 Umezawa's Jitte | removal | Provides a designated removal card.
- 1 Vona, Butcher of Magan | removal | Provides a designated removal card.
- 1 Witch of the Moors | removal | Provides a designated removal card.
- 1 Ajani's Pridemate | synergy | Provides a designated lifegain-synergy card.
- 1 Angel of Vitality | synergy | Provides a designated lifegain-synergy card.
- 1 Angelic Accord | synergy | Provides a designated lifegain-synergy card.
- 1 Bloodthirsty Aerialist | synergy | Provides a designated lifegain-synergy card.
- 1 Heliod, Sun-Crowned | synergy | Provides a designated lifegain-synergy card.
- 1 Resplendent Angel | synergy | Provides a designated lifegain-synergy card.
- 1 Vito, Thorn of the Dusk Rose | synergy | Provides a designated lifegain-synergy card.
- 1 Archangel of Thune | threat | A designated threat for the lifegain plan.
- 1 Attended Healer | threat | A designated threat for the lifegain plan.
- 1 Blood Baron of Vizkopa | threat | A designated threat for the lifegain plan.
- 1 Celestine, the Living Saint | threat | A designated threat for the lifegain plan.
- 1 Enduring Tenacity | threat | A designated threat for the lifegain plan.
- 1 Gideon's Company | threat | A designated threat for the lifegain plan.
- 1 Rhox Faithmender | threat | A designated threat for the lifegain plan.
- 1 Twinblade Paladin | threat | A designated threat for the lifegain plan.
- 1 Valkyrie Harbinger | threat | A designated threat for the lifegain plan.
- 1 Vampire Scrivener | threat | A designated threat for the lifegain plan.
- 1 Vanguard Seraph | threat | A designated threat for the lifegain plan.
- 1 Angel of Destiny | wincon | A finisher-marked designated win condition.
- 1 Felidar Sovereign | wincon | A finisher-marked designated win condition.
- 1 Test of Endurance | wincon | A finisher-marked designated win condition.
- 1 Fumigate | wipe | Provides a designated board wipe.
- 1 Kaya's Wrath | wipe | Provides a designated board wipe.
- 1 The Meathook Massacre | wipe | Provides a board wipe and carries a finisher marking.

</details>

### 2. aristocrats Commander, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 181 names.

Commander: Denethor, Ruling Steward.

Grade: baseline, score 0.29, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $139.82 the whole deck.

**Summary:** This is a black-white aristocrats deck that develops a protected creature board, uses its dedicated synergy pieces to make creature exchanges favorable, and keeps the hand full with equipment and black draw spells. It wins by pressing with creature threats or landing one of its larger finishing cards after opponents have spent resources fighting the board. The deck gives up the explosive starts and dense stack interaction of the strongest lists in favor of a steadier battlefield-focused game.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The cards do hang together as a black-white removal/equipment midrange pile that protects a few creatures and draws off gear, but the summary's headline plan (aristocrats) is nowhere in the list — there are essentially no free sacrifice outlets, token makers, or death-trigger drain payoffs — so the stated plan and the built plan diverge.
- PLAN theme_fit=no: The person asked for an aristocrats deck and got a Denethor good-stuff pile of nine spot removal spells, two wraths, and six equipment with no sacrifice engine, no token production, and no Blood Artist-type payoff; only Nasty End and Bitter Triumph even sacrifice anything.
- PLAN useful_as_built=yes: 36 lands in a two-color deck with eleven mana rocks/ramp pieces, a reasonable curve, heavy interaction, and enough creature threats plus Vraska and Archfiend of Ifnir means it can be shuffled up and played functionally as written.
- PLAN summary_honest=no: It opens by calling the list 'a black-white aristocrats deck' and credits 'dedicated synergy pieces' for favorable creature exchanges, when the synergy slots are Gollum/Gríma/Phantom Train-style value creatures rather than sac outlets or drain effects, so the central claim is not backed by the cards.
- [INFO] `curve_summary`: average mana value 2.92 over 63 nonland cards
- [INFO] `mana_pass`: the builder moved 1 card of the mana base to bring the deck inside its power level
- [INFO] `cards_trimmed`: the list was 1 card over, so the builder cut this card: Vanquish the Horde

<details><summary>The deck list</summary>

- 19 Plains | land | Provides a reliable white source while keeping the mana base untapped.
- 17 Swamp | land | Provides a reliable black source while keeping the mana base untapped.
- 1 Skullclamp | draw | Efficient card-draw engine for a creature-focused aristocrats plan.
- 1 Idol of Oblivion | draw | Repeatable card-draw support for the deck's token and sacrifice-oriented game.
- 1 Lembas | draw | Low-cost card selection and draw support.
- 1 Mask of Memory | draw | Turns creature combat into steady card draw.
- 1 Tome of Legends | draw | Keeps cards flowing alongside the commander.
- 1 Stone of Erech | draw | Provides a compact draw piece for the deck's longer games.
- 1 Call of the Ring | draw | Supplies continuing card draw for the creature plan.
- 1 Night's Whisper | draw | Efficiently refills the hand.
- 1 Painful Truths | draw | Provides a larger burst of card draw.
- 1 Nasty End | draw | Provides draw that fits the black creature shell.
- 1 Buster Sword | draw | Adds another equipment-based source of card draw.
- 1 Sol Ring | ramp | Accelerates early development with a compact mana investment.
- 1 Arcane Signet | ramp | Fixes both deck colors while accelerating the commander and threats.
- 1 Wayfarer's Bauble | ramp | Early ramp that also supports reliable colored mana.
- 1 Fellwar Stone | ramp | Low-cost mana acceleration and color fixing.
- 1 Thought Vessel | ramp | Provides inexpensive mana acceleration.
- 1 Commander's Sphere | ramp | Fixes mana early and remains useful later.
- 1 Chromatic Lantern | ramp | Strengthens color fixing while adding mana.
- 1 Relic of Legends | ramp | Provides flexible mana in a legendary creature deck.
- 1 Ring of the Lucii | ramp | Adds mana acceleration for the deck's creature-heavy plan.
- 1 Sword of the Animist | ramp | Turns creature attacks into ongoing mana development.
- 1 Bender's Waterskin | ramp | Maintains the ramp count with a colorless mana piece.
- 1 Swords to Plowshares | removal | Efficient answer to an opposing creature.
- 1 Generous Gift | removal | Flexible removal for troublesome permanents.
- 1 Get Lost | removal | Versatile answer to key opposing permanents.
- 1 Destroy Evil | removal | Efficient removal with useful coverage.
- 1 Crib Swap | removal | Creature removal that fits the deck's tribal-adjacent creature plan.
- 1 Dispatch | removal | Low-cost creature interaction.
- 1 Stroke of Midnight | removal | Instant-speed answer to a wide range of permanents.
- 1 Banishing Light | removal | Broad permanent removal for problems that survive creature combat.
- 1 Bitter Triumph | removal | Flexible black removal for creatures and planeswalkers.
- 1 Austere Command | wipe | Flexible reset button that can preserve the most useful parts of the board.
- 1 Fumigate | wipe | Creature-board reset for games that get out of hand.
- 1 Lightning Greaves | interaction | Protects a key creature or the commander while supporting immediate pressure.
- 1 Swiftfoot Boots | interaction | Provides repeatable protection for important creatures.
- 1 Darksteel Plate | interaction | Keeps a key creature on the battlefield through removal and combat.
- 1 Champion's Helm | interaction | Protection for the commander and other legendary threats.
- 1 Gift of Immortality | interaction | Helps preserve a central creature through opposing interaction.
- 1 Clever Concealment | interaction | Protects the developed board from opposing disruption.
- 1 Reprieve | interaction | Provides flexible stack interaction while replacing itself.
- 1 Unbreakable Formation | interaction | Protects the creature board and helps maintain pressure.
- 1 Gollum the Abandoned | synergy | Core aristocrats synergy piece for the sacrifice-focused plan.
- 1 Gollum, Patient Plotter | synergy | Adds another dedicated synergy creature for the aristocrats shell.
- 1 Gríma Wormtongue | synergy | Supports the deck's sacrifice-oriented creature synergies.
- 1 Heirloom Auntie | synergy | Adds a creature-based synergy piece to the aristocrats core.
- 1 Joo Dee, One of Many | synergy | Contributes to the deck's creature synergy package.
- 1 Arcade Cabinet | synergy | Colorless synergy support that fits the deck's artifact and creature elements.
- 1 Phantom Train | synergy | Adds a resilient artifact-based synergy piece.
- 1 Bill the Pony | threat | Creature threat that contributes to battlefield pressure.
- 1 Vengeful Villagers | threat | Creature threat that reinforces the deck's board presence.
- 1 Ahriman | threat | Creature threat that also offers access to cards.
- 1 Angel of Serenity | threat | Large evasive creature threat with a removal role.
- 1 Exemplar of Light | threat | Creature threat that also supports the deck's card flow.
- 1 Inspiring Overseer | threat | Flying creature threat that contributes to card advantage.
- 1 Cirith Ungol Patrol | threat | Creature threat that helps keep the hand supplied.
- 1 Frontline Medic | threat | Creature threat that also provides interaction for the board.
- 1 Massacre Girl, Known Killer | threat | Creature threat that contributes additional card draw.
- 1 Palace Jailer | threat | Creature threat that doubles as removal.
- 1 Fiend Hunter | threat | Creature threat that supplies creature removal.
- 1 Witch-king of Angmar | threat | Finishing creature threat that also provides removal.
- 1 Archfiend of Ifnir | wincon | Dedicated finishing card for closing creature-heavy games.
- 1 Grave Venerations | wincon | Dedicated finishing card for the aristocrats strategy.
- 1 Vraska, Betrayal's Sting | wincon | High-impact finisher that also contributes to the deck's mana plan.

</details>

### 3. artifacts Commander, bracket 4

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 320 names.

Commander: Urza, Lord High Artificer.

Grade: typical, score 0.62, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $4225.02 to buy, $4225.02 the whole deck.

**Summary:** This is a fast artifact pressure deck that uses early mana to establish Urza, develop a dense artifact board, and keep pivotal turns safe with efficient interaction. It wins through its artifact-based alternate win conditions or by turning a developed board into combat pressure, while tutors help locate the most important artifact pieces. The deck gives up some creature-specific flexibility and relies heavily on maintaining its artifact infrastructure through opposing disruption.

The quality model grades this deck typical against the top lists of the format: the deck makes more mana on turn four than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

- PLAN plan_coherent=yes: Nearly every card supports one plan — cheap artifacts and fast mana to power out Urza, a construct/board that turns into pressure, artifact tutors, and artifact-count payoffs like Mechanized Production and Mirrodin Besieged — with only Thassa's Oracle sitting oddly since there is no library-emptying piece to pair with it.
- PLAN theme_fit=yes: It is exactly the requested mono-blue Urza, Lord High Artificer artifact deck at high power, using bracket-4-appropriate fast mana, free counterspells, and tutors without stepping into a two-card infinite-combo cEDH shell.
- PLAN useful_as_built=yes: Thirty-three lands plus a dense fast-mana suite, a low curve, ample card draw and interaction, and multiple routes to victory (construct beats, Kappa Cannoneer, Mechanized Production, Mirrodin Besieged) mean it can be shuffled up and played as printed.
- PLAN summary_honest=yes: placeholder
- [INFO] `curve_summary`: average mana value 2.97 over 66 nonland cards

<details><summary>The deck list</summary>

- 25 Island | land | Primary blue mana base for the artifact-heavy plan.
- 1 Mystic Sanctuary | land | Blue-producing utility land.
- 1 Otawara, Soaring City | land | Blue-producing utility land.
- 1 Seat of the Synod | land | Artifact land that supports the deck's artifact density.
- 1 Command Tower | land | Reliable blue source for Urza.
- 1 Glimmervoid | land | Flexible colored source in an artifact deck.
- 1 Spire of Industry | land | Flexible colored source alongside artifacts.
- 1 Inventors' Fair | land | Artifact-focused utility land.
- 1 Urza's Saga | land | Artifact-focused utility land.
- 1 Chrome Mox | ramp | Fast mana for explosive early development.
- 1 Lion's Eye Diamond | ramp | Fast mana that raises the deck's explosive potential.
- 1 Lotus Petal | ramp | Fast mana for early acceleration.
- 1 Mana Vault | ramp | Fast mana that accelerates Urza and larger artifacts.
- 1 Mox Diamond | ramp | Fast mana that helps deploy the board quickly.
- 1 Mox Opal | ramp | Fast mana supported by the dense artifact suite.
- 1 Moonsnare Prototype | ramp | Low-cost artifact acceleration.
- 1 Sol Ring | ramp | Efficient artifact acceleration.
- 1 Springleaf Drum | ramp | Low-cost mana production for early turns.
- 1 Metalworker | ramp | Artifact-based mana acceleration.
- 1 Grand Architect | ramp | Accelerates the artifact-focused board.
- 1 Inspiring Statuary | ramp | Turns the artifact board into additional mana support.
- 1 Vedalken Engineer | ramp | Creature-based artifact acceleration.
- 1 Rhystic Study | draw | Persistent card advantage for a proactive deck.
- 1 Thirst for Knowledge | draw | Efficient card selection in an artifact shell.
- 1 Sai, Master Thopterist | draw | Artifact-based value engine.
- 1 Forensic Gadgeteer | draw | Provides card advantage while supporting artifacts.
- 1 Riddlesmith | draw | Cheap artifact-driven card selection.
- 1 Braided Net // Braided Quipu | draw | Artifact-based card advantage.
- 1 Era of Innovation | draw | Supports artifact development with card advantage.
- 1 Esoteric Duplicator | draw | Artifact card advantage option.
- 1 Cerebral Download | draw | Low-cost blue card advantage.
- 1 Thopter Spy Network | draw | Ongoing artifact-based card advantage.
- 1 Nexus of Becoming | draw | Artifact card advantage for longer games.
- 1 Waterbender Ascension | draw | Additional blue card advantage.
- 1 Force of Will | interaction | Premium protection for pivotal turns.
- 1 Fierce Guardianship | interaction | Protects Urza and key artifacts efficiently.
- 1 An Offer You Can't Refuse | interaction | Low-cost stack interaction with incidental acceleration.
- 1 Metallic Rebuke | interaction | Artifact-supported stack interaction.
- 1 Disruption Protocol | interaction | Artifact-supported stack interaction.
- 1 Ice Out | interaction | Flexible blue stack interaction.
- 1 Stoic Rebuttal | interaction | Artifact-supported countermagic.
- 1 Padeem, Consul of Innovation | interaction | Protective interaction for important artifacts.
- 1 Welding Jar | interaction | Cheap protection for core artifacts.
- 1 Vronos, Masked Inquisitor | interaction | Protective interaction for the board.
- 1 Ghostly Flicker | interaction | Flexible protection and utility interaction.
- 1 Aether Spellbomb | removal | Cheap artifact-based answer to creatures.
- 1 Cyber Conversion | removal | Flexible blue removal.
- 1 Into Thin Air | removal | Blue removal for problematic permanents.
- 1 Ravenform | removal | Versatile blue removal.
- 1 Resculpt | removal | Efficient removal option.
- 1 Skysovereign, Consul Flagship | removal | Repeatable artifact-based removal pressure.
- 1 Spine of Ish Sah | removal | Broad artifact answer for difficult permanents.
- 1 Transmogrifying Wand | removal | Artifact-based creature removal.
- 1 Unable to Scream | removal | Low-cost blue creature answer.
- 1 Portal to Phyrexia | removal | High-impact artifact removal option.
- 1 Transmute Artifact | synergy | Finds a key artifact and supports artifact synergies.
- 1 Whir of Invention | synergy | Instant-speed artifact tutor that rewards the artifact board.
- 1 Clock of Omens | synergy | Supports artifact-based sequencing and engines.
- 1 Kappa Cannoneer | threat | Major artifact threat that pressures opponents quickly.
- 1 Master Transmuter | threat | Threat that supports artifact-centric development.
- 1 Phyrexian Metamorph | threat | Flexible artifact creature threat.
- 1 Traxos, Scourge of Kroog | threat | Efficient large artifact attacker.
- 1 Arcbound Crusher | threat | Artifact creature that adds combat pressure.
- 1 Arcbound Reclaimer | threat | Artifact creature threat with artifact relevance.
- 1 Dross Scorpion | threat | Artifact creature that contributes board pressure.
- 1 Frogmite | threat | Low-cost artifact creature for board development.
- 1 Lodestone Golem | threat | Artifact creature that applies meaningful pressure.
- 1 Rust Elemental | threat | Large artifact creature threat.
- 1 Whirler Rogue | threat | Creature threat that fits the artifact plan.
- 1 Chrome Steed | threat | Artifact creature that adds combat pressure.
- 1 Thassa's Oracle | wincon | Compact alternate win condition.
- 1 Mechanized Production | wincon | Artifact-centered alternate win condition.
- 1 Mirrodin Besieged | wincon | Artifact-centered alternate win condition.
- 1 Engineered Explosives | wipe | Flexible low-cost board reset.
- 1 Rebuild | wipe | Artifact-focused board reset when the table gets ahead.

</details>

### 4. dinosaur tribal, bracket 2

Format: Commander. Theme: dinosaurs. Pool: any_card. Shortlist: 270 names.

Commander: Gishath, Sun's Avatar.

Grade: bad, score 0.04, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $336.73 to buy, $336.73 the whole deck.

**Summary:** This deck develops its mana, builds a board around Dinosaur synergy, and keeps cards flowing while it deploys increasingly threatening creatures. It wins by attacking with its Dinosaur threats and dedicated finishers, with protective interaction and removal to support key turns. It gives up some speed and all-purpose answers in exchange for a clear, creature-forward plan that is easy to follow at the table.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The curve sits high for the format, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Every slot supports a Gishath dinosaur-tribal beatdown plan: ramp into big dinos, cost reducers, card draw off creatures, protection for the commander, and dinosaur-flavored removal, with no competing subtheme.
- PLAN theme_fit=yes: It is a mono-dinosaur Gishath Commander deck of roughly precon-level power, matching the requested tribe, format, and bracket 2 target.
- PLAN useful_as_built=partly: 36 lands plus mana rocks and dorks in a three-color deck with dual lands is playable and the win path is clear, but the very high curve, several 7-8 drops, and self-inflicted wraths like Blasphemous Act make it clunkier than ideal for a new player.
- PLAN summary_honest=partly: The description of ramp, draw, creature-forward wins, protection and removal is accurate, but it omits the three board wipes (Blasphemous Act, Austere Command, Raging Swordtooth) that sit oddly with a go-wide creature deck, and it glosses over the fact that the deck's protection is largely commander-focused.
- [INFO] `curve_summary`: average mana value 3.68 over 62 nonland cards
- [INFO] `mana_pass`: the builder moved 2 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 3 Forest | land | Basic green source for early development and reliable ramp casting.
- 1 Forest | land | Basic green source for early development and reliable ramp casting.
- 1 Forest | land | Basic green source for early development and reliable ramp casting.
- 1 Forest | land | Basic green source for early development and reliable ramp casting.
- 1 Forest | land | Basic green source for early development and reliable ramp casting.
- 1 Forest | land | Basic green source for early development and reliable ramp casting.
- 1 Forest | land | Basic green source for early development and reliable ramp casting.
- 1 Forest | land | Basic green source for early development and reliable ramp casting.
- 1 Forest | land | Basic green source for early development and reliable ramp casting.
- 1 Plains | land | Basic white source for the deck's white Dinosaur cards and support spells.
- 1 Plains | land | Basic white source for the deck's white Dinosaur cards and support spells.
- 1 Plains | land | Basic white source for the deck's white Dinosaur cards and support spells.
- 1 Plains | land | Basic white source for the deck's white Dinosaur cards and support spells.
- 1 Plains | land | Basic white source for the deck's white Dinosaur cards and support spells.
- 1 Plains | land | Basic white source for the deck's white Dinosaur cards and support spells.
- 1 Mountain | land | Basic red source for the deck's red Dinosaur cards and support spells.
- 1 Mountain | land | Basic red source for the deck's red Dinosaur cards and support spells.
- 1 Mountain | land | Basic red source for the deck's red Dinosaur cards and support spells.
- 1 Mountain | land | Basic red source for the deck's red Dinosaur cards and support spells.
- 1 Mountain | land | Basic red source for the deck's red Dinosaur cards and support spells.
- 1 Mountain | land | Basic red source for the deck's red Dinosaur cards and support spells.
- 1 Mountain | land | Basic red source for the deck's red Dinosaur cards and support spells.
- 1 Mountain | land | Basic red source for the deck's red Dinosaur cards and support spells.
- 1 Mountain | land | Basic red source for the deck's red Dinosaur cards and support spells.
- 1 Mountain | land | Basic red source for the deck's red Dinosaur cards and support spells.
- 1 Command Tower | land | Flexible three-color land for a smooth mana base.
- 1 City of Brass | land | Flexible three-color land that covers any needed color.
- 1 Mana Confluence | land | Flexible three-color land that covers any needed color.
- 1 Forbidden Orchard | land | Additional flexible source for all three deck colors.
- 1 Sacred Foundry | land | Red-white dual land for dependable color access.
- 1 Temple Garden | land | Green-white dual land for dependable color access.
- 1 Stomping Ground | land | Red-green dual land for dependable color access.
- 1 Bountiful Promenade | land | Green-white dual land that supports the multiplayer mana base.
- 1 Brushland | land | Green-white dual land that adds another early colored source.
- 1 Karplusan Forest | land | Red-green dual land that adds another early colored source.
- 1 Sol Ring | ramp | Efficient early mana acceleration for reaching the deck's larger cards.
- 1 Arcane Signet | ramp | Reliable fixing and acceleration across all three colors.
- 1 Fellwar Stone | ramp | Early mana acceleration that helps cast multicolored spells.
- 1 Birds of Paradise | ramp | Early creature-based fixing and acceleration.
- 1 Delighted Halfling | ramp | Early creature-based acceleration for the deck's game plan.
- 1 Drover of the Mighty | ramp | Creature-based mana acceleration that fits the Dinosaur deck.
- 1 Thunderherd Migration | ramp | Dinosaur-focused ramp that improves land development.
- 1 Topiary Stomper | ramp | Dinosaur ramp piece that contributes to the creature plan.
- 1 Hulking Raptor | ramp | Dinosaur ramp piece that remains a meaningful board presence.
- 1 Intrepid Paleontologist | ramp | Early ramp support that helps deploy Dinosaur spells.
- 1 Beast Whisperer | draw | Creature-based draw support for a deck built around creatures.
- 1 Garruk's Uprising | draw | Draw support that rewards the deck's threat-heavy creature plan.
- 1 Guardian Project | draw | Ongoing draw support for deploying a steady stream of creatures.
- 1 Harmonize | draw | Straightforward draw spell that helps refill the hand.
- 1 Return of the Wildspeaker | draw | Flexible draw support for a board with substantial threats.
- 1 Ripjaw Raptor | draw | Dinosaur draw source that also contributes to the creature plan.
- 1 Rishkar's Expertise | draw | High-impact draw spell for rebuilding or extending the board.
- 1 Earthshaker Dreadmaw | draw | Dinosaur draw option that keeps the deck's creature density high.
- 1 Kutzil, Malamet Exemplar | draw | Creature-based draw support that fits the combat-focused plan.
- 1 Archivist of Oghma | draw | Low-cost draw support for keeping cards available through the game.
- 1 Toski, Bearer of Secrets | draw | Combat-oriented draw support for a creature-heavy board.
- 1 Heroic Intervention | interaction | Protective interaction for preserving the board of Dinosaur threats.
- 1 Flawless Maneuver | interaction | Protective interaction for keeping an established board intact.
- 1 Lightning Greaves | interaction | Equipment-based protection for important creatures.
- 1 Swiftfoot Boots | interaction | Equipment-based protection for important creatures.
- 1 Boros Charm | interaction | Versatile interaction that also carries finisher potential.
- 1 Akroma's Will | interaction | Protective combat interaction with finisher potential.
- 1 Itzquinth, Firstborn of Gishath | removal | Low-cost Dinosaur removal that supports the tribe.
- 1 Savage Stomp | removal | Dinosaur-focused removal for answering opposing problems.
- 1 Ravenous Sailback | removal | Dinosaur removal attached to a creature body.
- 1 Tranquil Frillback | removal | Flexible Dinosaur removal that remains part of the creature plan.
- 1 Burning Sun's Avatar | removal | Dinosaur removal option that adds a substantial creature.
- 1 Kogla and Yidaro | removal | Versatile Dinosaur-themed removal threat.
- 1 Trumpeting Carnosaur | removal | Dinosaur removal that supports the deck's main creature plan.
- 1 Commune with Dinosaurs | synergy | Dinosaur synergy piece that improves access to the deck's plan.
- 1 Kinjalli's Caller | synergy | Dinosaur-focused support for deploying the creature suite.
- 1 Otepec Huntmaster | synergy | Dinosaur-focused support that helps the deck develop its board.
- 1 Marauding Raptor | synergy | Dinosaur synergy creature that strengthens the tribal plan.
- 1 Hunting Velociraptor | synergy | Dinosaur synergy creature for advancing the creature strategy.
- 1 Dinosaur Stampede | synergy | Tribal combat support for a board full of Dinosaurs.
- 1 Raptor Companion | synergy | Low-cost Dinosaur that supports tribal density and early board presence.
- 1 Territorial Hammerskull | synergy | Dinosaur synergy creature that contributes to the combat plan.
- 1 Orazca Frillback | synergy | Dinosaur synergy creature that reinforces the tribal core.
- 1 Sunfrill Imitator | synergy | Dinosaur synergy creature that benefits from the deck's Dinosaur focus.
- 1 Regisaur Alpha | threat | Core Dinosaur threat that helps establish a dangerous board.
- 1 Quartzwood Crasher | threat | Dinosaur threat for pressuring opponents through combat.
- 1 Pantlaza, Sun-Favored | threat | Dinosaur threat that strengthens the deck's tribal presence.
- 1 Shifting Ceratops | threat | Efficient Dinosaur threat with finisher potential.
- 1 Rampaging Ceratops | threat | Dinosaur threat with finisher potential in combat.
- 1 Majestic Heliopterus | threat | Dinosaur threat that adds meaningful board pressure.
- 1 Harnessed Snubhorn | threat | Lower-cost Dinosaur threat that supports a smoother curve.
- 1 Pathfinding Axejaw | threat | Dinosaur threat that advances the creature-focused plan.
- 1 Needletooth Pack | threat | Dinosaur threat that adds to the deck's board pressure.
- 1 Palani's Hatcher | threat | Dinosaur threat that supplies a strong tribal board presence.
- 1 Ghalta and Mavren | threat | Dinosaur threat for turning a developed board into pressure.
- 1 Cavern Stomper | threat | Dinosaur threat with finisher potential.
- 1 Dinosaurs on a Spaceship | wincon | Dedicated Dinosaur win condition with finisher potential.
- 1 Etali, Primal Storm | wincon | Dedicated Dinosaur win condition with finisher potential.
- 1 Huatli, Poet of Unity // Roar of the Fifth People | wincon | Dedicated win condition that fits the deck's creature-focused plan.
- 1 Blasphemous Act | wipe | Board wipe for recovering when opposing creatures get ahead.
- 1 Austere Command | wipe | Flexible board wipe for resetting difficult board states.
- 1 Raging Swordtooth | wipe | Dinosaur-based board wipe that keeps the tribal theme present.

</details>

### 5. blink Commander, owned first

Format: Commander. Theme: blink. Pool: owned_first. Shortlist: 169 names.

Commander: Gilraen, Dúnedain Protector.

Grade: bad, score 0.00, model 20260914T154223Z.

Cards: 0 main, 0 sideboard. Repair turn: yes, for copy_limit, profile_off_band. Block findings: 1.

Cost: $0.00 to buy, $0.00 the whole deck.

**Summary:** No legal repaired deck can be returned from the supplied information. The prior deck list is absent, and the shortlist contains only one uniquely legal card marked as a finisher: Angel of Serenity. The required minimum is two, while duplicate nonbasic names are prohibited.

The quality model grades this deck below the precon baseline against the top lists of the format: the deck holds 0 lands, and its curve needs about 31.

Bracket 3 wants 2 or more finishers, the cards that win the game, and the deck holds 0. To close the gap, add Angel of Serenity (owned).

- JUDGE [true]: "duplicate nonbasic names are prohibited". This states the Commander singleton rule: a Commander deck may contain only one of each card other than basic lands. That is an accurate statement of the format's deckbuilding rule, but it is a rule of the game and should not appear in the summary.
- JUDGE [unknown]: "Bracket 3 wants 2 or more finishers ... The required minimum is two". This describes the builder's bracket heuristic for how many game-ending cards the deck should hold, not a rule of Magic. Bracket guidance is advisory and not a legality or deckbuilding rule, so its truth as a rule cannot be judged.
- PLAN plan_coherent=no: There are no cards at all beyond the commander, so no plan of any kind is carried by the list.
- PLAN theme_fit=no: The person asked for a bracket 3 blink deck from their library and received an empty list with no blink pieces, no lands, and no payoffs.
- PLAN useful_as_built=no: An empty 0-card list with no mana and no win conditions cannot be picked up and played.
- PLAN summary_honest=yes: The summary openly states that no deck could be built, that the list holds 0 lands and 0 finishers, and explains why, hiding nothing.
- [BLOCK] `deck_size`: deck has 1 cards, the format needs exactly 100 (commander included)
- [WARN] `land_count`: 0 lands: the guide range for this format is 27 to 41
- [WARN] `profile_off_band`: the land count is 0, and bracket 3 wants 34 to 38
- [WARN] `profile_off_band`: the average mana value of the nonland cards is 0, and bracket 3 wants 2 to 3.5
- [WARN] `profile_off_band`: the ramp count is 0, and bracket 3 wants 8 to 13
- [WARN] `profile_off_band`: the draw count is 0, and bracket 3 wants 8 to 14
- [WARN] `profile_off_band`: the removal count is 0, and bracket 3 wants 6 to 12
- [WARN] `profile_off_band`: the wipe count is 0, and bracket 3 wants 2 to 5
- [WARN] `profile_off_band`: the interaction count is 0, and bracket 3 wants 4 to 12
- [WARN] `profile_off_band`:  is 0, and bracket 3 wants 2 or more

<details><summary>The deck list</summary>


</details>

### 6. Modern tempo, tournament

Format: Modern. Theme: tempo. Pool: any_card. Shortlist: 167 names.

Grade: typical, score 0.44, model 20260914T154223Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $356.90 to buy, $356.90 the whole deck.

**Summary:** This blue-red tempo deck establishes early pressure with Ledger Shredder, Ragavan, Nimble Pilferer, and Faerie Mastermind, then uses cheap draw, removal, and counterspells to keep an opponent from stabilizing. It wins by maintaining that pressure while protecting its threats and finding additional spells, with Temporal Trespass supporting the deck’s spell-focused plan. In exchange, it favors efficient early exchanges over broad permanent answers and a deeply resilient late game.

The quality model grades this deck typical against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The deck runs its spells as playsets, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade.

- PLAN plan_coherent=partly: The creatures, cheap cantrips, burn, and counterspells all point at one clean tempo plan, but four copies of Temporal Trespass — a delve extra-turn spell with no fetchland or self-mill support and no real payoff for an extra turn beyond a couple of 2-power attacks — pulls against the low-curve interaction plan.
- PLAN theme_fit=partly: It is a Modern blue-red tempo deck in the modern UR shell (Ragavan, Shredder, Bolt, Counterspell), but the person asked for Delver and there is no Delver of Secrets or equivalent one-drop flipper, and the Trespass package is not part of any tempo build.
- PLAN useful_as_built=partly: The mana base is clean and the curve is low enough to function, and eight to twelve evasive threats plus burn can close games, but four dead-ish Temporal Trespass in a 36-spell deck means a real share of draws stall on a card the deck cannot cast.
- PLAN summary_honest=partly: The threat, draw, removal, and counter descriptions match the list exactly, but calling Temporal Trespass a supporter of the 'spell-focused plan' dresses up four copies of a card the deck can rarely cast cheaply, and the summary never acknowledges the absent Delver the request named.
- [INFO] `curve_summary`: average mana value 2.44 over 36 nonland cards

<details><summary>The deck list</summary>

- 4 Island | land | Provides a reliable blue source for the deck's early tempo spells.
- 4 Mountain | land | Provides a reliable red source for the deck's removal suite.
- 4 Steam Vents | land | Supplies both deck colors while keeping the mana base focused.
- 4 Spirebluff Canal | land | Adds another blue-red land slot for the early game.
- 4 Shivan Reef | land | Adds flexible blue-red color access.
- 2 Riverglide Pathway // Lavaglide Pathway | land | Serves as a flexible choice of either deck color.
- 2 Thundering Falls | land | Rounds out the blue-red land base.
- 4 Consider | draw | Efficiently fills the deck's draw package while supporting a spell-heavy plan.
- 2 Preordain | draw | Adds more early card selection to smooth tempo hands.
- 4 Counterspell | interaction | Provides broad stack interaction to protect the tempo plan.
- 2 Spell Pierce | interaction | Gives the deck inexpensive early interaction.
- 4 Lightning Bolt | removal | Forms a fast, efficient red removal package.
- 4 Galvanic Discharge | removal | Adds more instant-speed removal for clearing pressure blockers.
- 4 Temporal Trespass | synergy | Fills the requested synergy slots in the blue spell shell.
- 4 Ledger Shredder | threat | Provides a low-cost blue threat for the tempo plan.
- 4 Ragavan, Nimble Pilferer | threat | Provides an early red threat that supports proactive openings.
- 4 Faerie Mastermind | threat | Completes the threat base with an additional blue tempo creature.

</details>

### 7. Modern burn, casual

Format: Modern. Theme: burn. Pool: any_card. Shortlist: 290 names.

Grade: baseline, score 0.38, model 20260914T154223Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $44.44 to buy, $44.44 the whole deck.

**Summary:** This mono-red burn deck applies steady pressure with cheap removal, burn-focused synergy cards, and a compact threat package that can finish games once the opponent is under pressure. It wins by maintaining that damage-focused pace from the opening turns through its larger threats, while giving up broad answers and matchup-specific sideboard tools for a direct, streamlined plan.

The quality model grades this deck at the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=no: The pieces pull apart: Thermo-Alchemist and Eidolon want a dense instant/sorcery count but the deck plays only about ten spells (half of them the near-blank Ancestral Anger), while Hazoret wants an empty hand and Torbran/Imodane want a board of creatures, so no single damage engine is actually supported.
- PLAN theme_fit=partly: It is mono-red, Modern-legal, and aggressive, but a burn deck asks for a stack of direct-damage spells and this is a creature deck with four Bolts, two Lava Darts, and two Risk Factors.
- PLAN useful_as_built=partly: The mana base and curve are fine and it can attack and win, but four copies each of three legends fight the legend rule, Ancestral Anger and Ancestors' Aid do almost nothing, and the stated burn clock is not there, so games will play out clunkily.
- PLAN summary_honest=partly: It correctly flags the lack of sideboard and broad answers, but calling this a package of "cheap removal" and "burn-focused synergy" overstates a list with six removal spells and burn payoffs that have almost no burn to feed them.
- [INFO] `curve_summary`: average mana value 2.56 over 36 nonland cards

<details><summary>The deck list</summary>

- 24 Mountain | land | Provides the mono-red mana base.
- 4 Ancestral Anger | draw | Fills the deck's draw slot while supporting the burn-focused plan.
- 2 Risk Factor | draw | Adds further draw for a burn-focused game plan.
- 2 Ancestors' Aid | ramp | Supplies the requested ramp support.
- 4 Lightning Bolt | removal | Core efficient removal for a burn deck.
- 2 Lava Dart | removal | Completes the removal package with additional red burn pressure.
- 4 Eidolon of the Great Revel | synergy | A synergy piece aligned with the aggressive burn plan.
- 4 Thermo-Alchemist | synergy | Supports the deck's spell-heavy burn theme.
- 4 Barret Wallace | threat | Provides a threat that can maintain pressure alongside burn spells.
- 4 Hazoret the Fervent | threat | Adds a resilient threat for closing games.
- 4 Torbran, Thane of Red Fell | threat | Provides a high-impact threat for the red damage plan.
- 2 Imodane, the Pyrohammer | threat | Rounds out the threat package with additional burn-oriented pressure.

</details>

### 8. Modern lifegain, FNM

Format: Modern. Theme: lifegain. Pool: any_card. Shortlist: 301 names.

Grade: bad, score 0.32, model 20260914T154223Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $442.06 to buy, $442.06 the whole deck.

**Summary:** This white-black lifegain deck starts with Soul Warden and Ajani's Pridemate, then builds a board of creatures whose pressure is reinforced by repeated life gain. Enduring Tenacity, Attended Healer, Gideon's Company, and Sheoldred, the Apocalypse provide the closing power, while Solitude and Deep-Cavern Bat clear the way and Enduring Innocence with Lembas helps keep cards flowing. It gives up some speed and mana efficiency for a board-centric plan with several higher-cost threats, so early interaction and careful sequencing matter against faster decks.

The quality model grades this deck below the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade.

- PLAN plan_coherent=yes: Nearly every card feeds or exploits repeated lifegain triggers — Soul Warden, Attended Healer, Ajani's Pridemate, Gideon's Company, Enduring Tenacity, Sheoldred — with a small removal package, so the pieces pull one direction.
- PLAN theme_fit=yes: It is a Modern white-black lifegain list at roughly FNM power, exactly the request, using both classic (Soul Warden, Pridemate) and modern lifegain payoffs.
- PLAN useful_as_built=partly: The mana base is fine (12 duals plus basics) and there are plenty of win conditions, but with Solitude needing white cards to pitch and a heavy cluster of four-drops plus only two Altars of ramp, the deck often stalls before its payoffs land.
- PLAN summary_honest=yes: It names cards actually in the list, credits the right roles to them, and openly admits the deck is slow, top-heavy and mana-inefficient, which matches the curve of four-drops.
- [INFO] `curve_summary`: average mana value 3.17 over 36 nonland cards

<details><summary>The deck list</summary>

- 4 Godless Shrine | land | Reliable white-black land for the two-color mana base.
- 4 Caves of Koilos | land | Provides either deck color without entering tapped.
- 4 Shattered Sanctum | land | Dual land that supports both primary colors.
- 6 Plains | land | Basic white source for the lifegain core.
- 6 Swamp | land | Basic black source for black threats and removal.
- 2 Altar of the Pantheon | ramp | Fills the requested mana acceleration slot while supporting both colors.
- 4 Solitude | removal | Primary removal option that also contributes a creature body.
- 2 Deep-Cavern Bat | removal | Additional creature-based removal for an FNM-oriented game plan.
- 4 Enduring Innocence | draw | Main repeatable draw card in a creature-heavy build.
- 2 Lembas | draw | Low-commitment draw support that fits the lifegain shell.
- 4 Soul Warden | synergy | Early lifegain synergy piece that turns creatures entering into progress.
- 4 Ajani's Pridemate | synergy | Efficient lifegain payoff that grows alongside Soul Warden.
- 4 Enduring Tenacity | threat | Black lifegain threat that pressures opponents as the deck gains life.
- 4 Attended Healer | threat | Creature threat that reinforces the deck's lifegain-focused board plan.
- 4 Gideon's Company | threat | Lifegain-oriented threat that benefits from the deck's core plan.
- 2 Sheoldred, the Apocalypse | threat | High-impact black threat that gives the deck a strong top end.

</details>

### 9. Standard midrange, FNM

Format: Standard. Theme: midrange. Pool: any_card. Shortlist: 164 names.

Grade: bad, score 0.34, model 20260914T154223Z.

Cards: 62 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $96.16 to buy, $96.16 the whole deck.

**Summary:** This black-green midrange deck establishes its mana early, builds cards through Darkstar Augur and Unholy Annex // Ritual Chamber, and controls the table with a dense package of removal. Goldvein Hydra, Rottenmouth Viper, Vein Ripper, and Massacre Wurm give it creature pressure while Season of Loss can recover a board that has fallen behind. It gives up specialized matchup tools in favor of a straightforward main-deck plan built around resilient mana, protection spells, and broad answers.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The cards pair in ways the top lists do not, and that lowers the grade. The curve sits high for the format, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core of threats plus a dense removal suite is a recognizable midrange plan, but the eight slots of protection spells (Blossoming Defense, Snakeskin Veil, Not Dead After All, Royal Treatment) belong to a hexproof/aura-style deck and are poorly supported by only a handful of creatures worth protecting, and the pile of two-ofs makes the plan inconsistent.
- PLAN theme_fit=yes: It is a black-green Standard midrange list with creatures, removal, and a fair curve, matching exactly what was asked for at FNM level.
- PLAN useful_as_built=partly: 24 lands plus four duals support the colors and there are enough finishers to close a game, but the curve is top-heavy (multiple five- and six-drops) with almost no ramp, and eight protection spells crowd out proactive plays, so the deck will often hold reactive cards it cannot profitably use.
- PLAN summary_honest=yes: Every element named — Darkstar Augur/Annex card advantage, the removal density, the listed finishers, Season of Loss as a recovery tool, and the lack of specialized tools — is present in the list, and it openly reports the quality model's criticisms rather than hiding them, though "establishes its mana early" overstates a mere two Llanowar Elves.
- [INFO] `curve_summary`: average mana value 3.00 over 38 nonland cards

<details><summary>The deck list</summary>

- 10 Forest | land | Provides a reliable green mana base for the deck.
- 10 Swamp | land | Provides a reliable black mana base for the deck.
- 2 Blooming Marsh | land | Adds black-green mana while keeping the land base compact.
- 2 Underground Mortuary | land | Adds another black-green land source.
- 2 Llanowar Elves | ramp | Supplies the deck's early acceleration.
- 4 Darkstar Augur | draw | Provides a creature-based source of cards.
- 2 Unholy Annex // Ritual Chamber | draw | Adds a durable draw-focused permanent.
- 4 Bitter Triumph | removal | Forms the efficient core of the deck's removal suite.
- 4 Nowhere to Run | removal | Gives the deck additional early removal.
- 2 Maelstrom Pulse | removal | Broadens removal coverage for problematic permanents.
- 2 Withering Torment | removal | Adds flexible instant-speed removal.
- 2 Season of Loss | wipe | Provides a reset option when opposing boards get ahead.
- 2 Massacre Wurm | wipe | Supplies a creature-based board-reset option.
- 2 Blossoming Defense | interaction | Helps protect important creatures from opposing answers.
- 2 Snakeskin Veil | interaction | Provides another efficient protection effect.
- 2 Not Dead After All | interaction | Helps a key creature survive removal.
- 2 Royal Treatment | interaction | Adds further protection for the deck's creatures.
- 2 Goldvein Hydra | ramp | Provides a scalable creature that also supports the mana plan.
- 2 Rottenmouth Viper | removal | Adds a creature that contributes to the removal plan.
- 2 Vein Ripper | removal | Provides a high-impact creature with removal utility.

</details>

### 10. Standard aggro, tournament

Format: Standard. Theme: aggro. Pool: any_card. Shortlist: 294 names.

Grade: bad, score 0.30, model 20260914T154223Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $80.24 to buy, $80.24 the whole deck.

**Summary:** This red-white aggro deck leads with a dense creature package and keeps pressure on through red-white synergy cards. Removal clears the way for its threats, while interaction helps preserve the attacking position and draw cards help sustain the hand. It gives up heavier late-game options in exchange for a direct, proactive plan built around deploying threats and forcing the opponent to respond.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Eighteen cheap creatures, an anthem-plus-damage effect, and cheap removal all push the same proactive attacking plan, with only Warleader's Call at four mana sitting slightly above the curve.
- PLAN theme_fit=yes: It is a two-color red-white creature aggro deck built for Standard, exactly the request, with no off-theme late-game or control elements.
- PLAN useful_as_built=partly: The curve, threat count, and win paths are fine, but the mana is strained: only four dual lands supporting a triple-red Boros Charm alongside double-white cards will produce regular color failures.
- PLAN summary_honest=yes: Every element the summary names — dense creatures, removal, interaction, card draw, no late game — is actually present in the list, and it does not claim any mechanic the cards lack.
- [INFO] `curve_summary`: average mana value 2.94 over 36 nonland cards
- [INFO] `mana_pass`: the builder moved 1 card of the mana base to bring the deck inside its power level
- [INFO] `cards_trimmed`: the list was 1 card over, so the builder cut this card: Inspiring Vantage

<details><summary>The deck list</summary>

- 10 Plains | land | White mana source for the deck.
- 10 Mountain | land | Red mana source for the deck.
- 2 Sacred Foundry | land | Red-white mana source.
- 1 Sunbillow Verge | land | Red-white mana source.
- 1 Elegant Parlor | land | Red-white mana source.
- 4 Fugitive Codebreaker | draw | A red card from the shortlist assigned to draw.
- 2 Inspiring Overseer | draw | A white card from the shortlist assigned to draw.
- 4 Boros Charm | interaction | A red-white interaction card for protecting the attack plan.
- 2 Sheltered by Ghosts | interaction | A white interaction card for the creature-focused plan.
- 4 Emeritus of Truce // Swords to Plowshares | removal | A white removal option for clearing opposing cards.
- 4 Case of the Gateway Express | removal | A removal option from the shortlist.
- 4 Warleader's Call | synergy | A red-white synergy piece for the aggressive deck.
- 4 Dragonback Lancer | threat | An aggressive threat for applying pressure.
- 4 Frilled Sparkshooter | threat | A red threat for maintaining pressure.
- 4 Vanguard Seraph | threat | A white threat that broadens the creature base.

</details>

### 11. Commander with a locked card

Format: Commander. Theme: sacrifice. Pool: any_card. Shortlist: 299 names.

Commander: Karlov of the Ghost Council.

Grade: baseline, score 0.04, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $817.25 to buy, $817.25 the whole deck.

**Summary:** Karlov of the Ghost Council leads a black-white sacrifice deck that develops a creature board, uses sacrifice-focused support to maintain momentum, and backs that plan with removal and protective interaction. It wins by turning its established sacrifice engine into pressure through its marked finishers and win conditions. The deck gives up some speed to maintain a substantial mana base and a broad suite of board control.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves one color far better than another, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every slot is a sacrifice outlet, a free-fodder token maker, a death-trigger drainer, or a payoff that converts sacrifices into cards or damage, and the drain effects also feed Karlov's lifegain counters, so the pieces all push one engine.
- PLAN theme_fit=yes: It is a Commander deck led by Karlov of the Ghost Council built squarely around sacrificing creatures, exactly what was asked.
- PLAN useful_as_built=yes: 36 lands with a healthy dual/fetch suite, a low curve, ample cheap fodder and token generation to fuel the outlets, and multiple redundant win routes (Zulaport/Bastion/Vindictive Vampire drains, Altar of Dementia mill, Karlov himself) make it playable as printed.
- PLAN summary_honest=yes: The summary is generic but makes no claim the list fails to support: the sacrifice engine, drain-based finishers, three wraths plus spot removal, 36 lands, and the black-leaning mana base are all genuinely present.
- [INFO] `curve_summary`: average mana value 3.02 over 63 nonland cards

<details><summary>The deck list</summary>

- 14 Swamp | land | Basic black source for the deck's spells.
- 7 Plains | land | Basic white source for the deck's spells.
- 1 Brightclimb Pathway // Grimclimb Pathway | land | Flexible black or white source.
- 1 Caves of Koilos | land | Black and white mana source.
- 1 City of Brass | land | Flexible colored mana source.
- 1 Command Tower | land | Reliable commander-deck color source.
- 1 Concealed Courtyard | land | Black and white mana source.
- 1 Eclipsed Steppe | land | Black and white mana source.
- 1 Godless Shrine | land | Black and white mana source.
- 1 Mana Confluence | land | Flexible colored mana source.
- 1 Marsh Flats | land | Land slot that supports the mana base.
- 1 Prismatic Vista | land | Land slot that supports the mana base.
- 1 Scrubland | land | Black and white mana source.
- 1 Shattered Sanctum | land | Black and white mana source.
- 1 Silent Clearing | land | Black and white mana source.
- 1 Turbulent Moor | land | Black and white mana source.
- 1 Vault of Champions | land | Black and white mana source.
- 1 Sol Ring | ramp | Efficient mana acceleration the deck is required to retain.
- 1 Ashnod's Altar | ramp | Mana acceleration that belongs in the sacrifice plan.
- 1 Culling the Weak | ramp | Fast mana acceleration for the sacrifice plan.
- 1 Deadly Dispute | ramp | Mana acceleration that also supports sacrificing.
- 1 Master of Dark Rites | ramp | Creature-based mana acceleration.
- 1 Phyrexian Altar | ramp | Mana acceleration for the sacrifice plan.
- 1 Pitiless Plunderer | ramp | Sacrifice-focused mana acceleration.
- 1 Priest of Forgotten Gods | ramp | Creature-based mana acceleration for the sacrifice plan.
- 1 Skullport Merchant | ramp | Mana acceleration on a sacrifice-oriented creature.
- 1 Warren Soultrader | ramp | Mana acceleration that fits the sacrifice plan.
- 1 Baron Bertram Graywater | draw | Card advantage for the sacrifice deck.
- 1 Corrupted Conviction | draw | Efficient card advantage that supports sacrificing.
- 1 Disciple of Bolas | draw | Card advantage for a creature-focused sacrifice plan.
- 1 Ecstatic Awakener // Awoken Demon | draw | Card advantage that fits the sacrifice theme.
- 1 Jerren, Corrupted Bishop // Ormendahl, the Corrupter | draw | Card advantage on a thematic creature.
- 1 Shadowheart, Dark Justiciar | draw | Card advantage for the creature-heavy plan.
- 1 Smothering Abomination | draw | Card advantage in the sacrifice shell.
- 1 Tevesh Szat, Doom of Fools | draw | Repeatable card advantage for the deck.
- 1 Thraxodemon | draw | Card advantage on a thematic creature.
- 1 Vampiric Rites | draw | Card advantage that supports sacrificing.
- 1 Village Rites | draw | Low-cost card advantage for the sacrifice plan.
- 1 Cartel Aristocrat | interaction | Protective interaction that fits the sacrifice plan.
- 1 Dark Privilege | interaction | Protective interaction for key creatures.
- 1 Fanatical Devotion | interaction | Protective interaction that uses the sacrifice theme.
- 1 Flare of Fortitude | interaction | Protective interaction for the board.
- 1 Gift of Doom | interaction | Protective interaction for an important permanent.
- 1 Promise of Tomorrow | interaction | Interaction that supports the creature plan.
- 1 Rescue from the Underworld | interaction | Reactive interaction for the sacrifice deck.
- 1 Spirit Bonds | interaction | Creature-focused protective interaction.
- 1 Attrition | removal | Sacrifice-oriented creature removal.
- 1 Ayli, Eternal Pilgrim | removal | Creature-based removal for the deck.
- 1 Blasting Station | removal | Sacrifice-oriented removal option.
- 1 Bone Shards | removal | Low-cost removal that supports sacrificing.
- 1 Bone Splinters | removal | Low-cost removal that supports sacrificing.
- 1 Eaten Alive | removal | Removal that fits the sacrifice plan.
- 1 Grave Pact | removal | Sacrifice-focused removal for opposing creatures.
- 1 Teysa, Orzhov Scion | removal | Thematic creature-based removal.
- 1 Yawgmoth, Thran Physician | removal | Creature-based removal for the sacrifice shell.
- 1 Altar of Dementia | synergy | Central sacrifice synergy piece.
- 1 Bastion of Remembrance | synergy | Sacrifice synergy with finisher potential.
- 1 Carrion Feeder | synergy | Low-cost sacrifice synergy creature.
- 1 Elas il-Kor, Sadistic Pilgrim | synergy | Sacrifice synergy with finisher potential.
- 1 Viscera Seer | synergy | Low-cost sacrifice synergy creature.
- 1 Witch's Oven | synergy | Artifact-based sacrifice synergy.
- 1 Zulaport Cutthroat | synergy | Sacrifice synergy with finisher potential.
- 1 Basri's Lieutenant | threat | Creature threat for the board plan.
- 1 Blood Host | threat | Creature threat that fits the sacrifice theme.
- 1 Corpse Harvester | threat | Creature threat for the deck.
- 1 Felisa, Fang of Silverquill | threat | Creature threat in the commander’s colors.
- 1 Flesh-Eater Imp | threat | Creature threat for the sacrifice shell.
- 1 Ghoulcaller Gisa | threat | Creature threat for the board plan.
- 1 Prowling Geistcatcher | threat | Creature threat that fits the creature plan.
- 1 Ratadrabik of Urborg | threat | Creature threat in the commander’s colors.
- 1 Reaper of Flight Moonsilver | threat | Creature threat for the deck.
- 1 Thallid Omnivore | threat | Creature threat that fits the sacrifice theme.
- 1 Vindictive Vampire | threat | Creature threat with finisher potential.
- 1 Vito's Inquisitor | threat | Creature threat for the deck.
- 1 Razaketh, the Foulblooded | wincon | A marked finisher and win condition.
- 1 Relic Vial | wincon | A marked finisher and win condition.
- 1 Sephiroth, Fabled SOLDIER // Sephiroth, One-Winged Angel | wincon | A marked finisher and win condition.
- 1 Austere Command | wipe | Broad board-reset option.
- 1 The Meathook Massacre | wipe | Board-reset option with finisher potential.
- 1 Toxic Deluge | wipe | Efficient board-reset option.

</details>

### 12. Commander on a budget

Format: Commander. Theme: tokens. Pool: any_card. Shortlist: 298 names.

Commander: Adeline, Resplendent Cathar.

Grade: bad, score 0.01, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $81.48 to buy, $81.48 the whole deck.

**Summary:** Adeline leads a focused white token deck that develops its mana, builds a broad board, and keeps cards flowing while using affordable removal and protective interaction to preserve momentum. It closes through its marked finishers and win-condition cards once the board is established. The deck gives up expensive premium mana and broader multicolor flexibility in favor of a consistent, budget-minded token plan.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The deck makes less mana on turn four than the norm, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core is a consistent mono-white go-wide plan (Adeline, Intangible Virtue, Divine Visitation, Felidar Retreat, anthem/token threats, Hour of Reckoning as a one-sided sweeper), but a cluster of off-plan filler — Coin of Mastery, Collector's Vault, Currency Converter, Goldvein Pick, Druidic Satchel — and Hero of Precinct One, which is nearly dead in a deck with almost no multicolored spells, dilute the plan.
- PLAN theme_fit=yes: It is exactly what was asked: a mono-white Adeline token Commander deck built from cheap, low-power cards consistent with a budget bracket-2 build.
- PLAN useful_as_built=yes: 37 lands in a nearly mono-colored deck with a low curve, plenty of cheap token makers and draw, plus Halo Fountain and a wide board with Adeline as clear routes to win, all play fine out of the box.
- PLAN summary_honest=partly: Token payoffs, card flow, removal, protection and Halo Fountain as a finisher are all genuinely present, but the claim that the deck "develops its mana" oversells a ramp package made largely of marginal trinket-treasure artifacts.
- [INFO] `curve_summary`: average mana value 3.44 over 62 nonland cards

<details><summary>The deck list</summary>

- 29 Plains | land | Core white mana base.
- 1 Ancient Den | land | Part of the mana base.
- 1 Command Tower | land | Part of the mana base.
- 1 Eiganjo, Seat of the Empire | land | Part of the mana base.
- 1 Exotic Orchard | land | Part of the mana base.
- 1 Mirrex | land | Part of the mana base.
- 1 Plaza of Heroes | land | Part of the mana base.
- 1 Spire of Industry | land | Part of the mana base.
- 1 Windbrisk Heights | land | Part of the mana base.
- 1 Coin of Mastery | ramp | Low-cost mana development.
- 1 Collector's Vault | ramp | Low-cost mana development.
- 1 Currency Converter | ramp | Low-cost mana development.
- 1 Druidic Satchel | ramp | Low-cost mana development.
- 1 Goldvein Pick | ramp | Low-cost mana development.
- 1 Karn, Living Legacy | ramp | Supports mana development.
- 1 Keeper of the Accord | ramp | Supports mana development.
- 1 Monologue Tax | ramp | Supports mana development.
- 1 The Restoration of Eiganjo // Architect of Restoration | ramp | Supports mana development.
- 1 Thousand Moons Smithy // Barracks of the Thousand | ramp | Supports mana development.
- 1 Angelic Sell-Sword | draw | Card-advantage option for the token plan.
- 1 Bygone Bishop | draw | Card-advantage option for the token plan.
- 1 Court of Grace | draw | Card-advantage option for the token plan.
- 1 Dawn of Hope | draw | Card-advantage option for the token plan.
- 1 Faramir, Field Commander | draw | Card-advantage option for the token plan.
- 1 Glimmer Seeker | draw | Card-advantage option for the token plan.
- 1 Idol of Oblivion | draw | Card-advantage option for the token plan.
- 1 Search the Premises | draw | Card-advantage option for the token plan.
- 1 Staff of the Storyteller | draw | Card-advantage option for the token plan.
- 1 Valiant Rescuer | draw | Card-advantage option for the token plan.
- 1 Wedding Announcement // Wedding Festivity | draw | Card-advantage option for the token plan.
- 1 Aerial Assault | removal | Efficient removal coverage.
- 1 Banishing Slash | removal | Efficient removal coverage.
- 1 Battle Menu | removal | Efficient removal coverage.
- 1 Citizen's Crowbar | removal | Efficient removal coverage.
- 1 Generous Gift | removal | Flexible removal coverage.
- 1 Kellan's Lightblades | removal | Efficient removal coverage.
- 1 Skyclave Apparition | removal | Creature-based removal coverage.
- 1 Ceaseless Conflict | wipe | Reset option when the board gets away.
- 1 Crisis of Conscience | wipe | Reset option when the board gets away.
- 1 Hour of Reckoning | wipe | Reset option for a creature-heavy table.
- 1 Aligned Heart | synergy | Token-plan synergy piece.
- 1 Anointer Priest | synergy | Token-plan synergy piece.
- 1 Automated Assembly Line | synergy | Token-plan synergy piece.
- 1 Cat Collector | synergy | Token-plan synergy piece.
- 1 Clarion Spirit | synergy | Token-plan synergy piece.
- 1 Divine Visitation | synergy | Token-plan synergy piece.
- 1 Felidar Retreat | synergy | Token-plan synergy piece.
- 1 Hanweir Militia Captain // Westvale Cult Leader | synergy | Token-plan synergy piece.
- 1 Hero of Precinct One | synergy | Token-plan synergy piece.
- 1 Intangible Virtue | synergy | Token-plan synergy piece.
- 1 Archon of Sun's Grace | threat | Token-focused board presence.
- 1 Attended Healer | threat | Token-focused board presence.
- 1 Basri's Lieutenant | threat | Token-focused board presence.
- 1 Cemetery Protector | threat | Token-focused board presence.
- 1 Defiler of Faith | threat | Token-focused board presence.
- 1 Emeria Angel | threat | Token-focused board presence.
- 1 Gwaihir, Greatest of the Eagles | threat | A finisher that also advances the board plan.
- 1 Hero of Bladehold | threat | Token-focused board presence.
- 1 Nahiri, the Lithomancer | threat | A finisher that also advances the board plan.
- 1 Requiem Angel | threat | A finisher that also advances the board plan.
- 1 Silverwing Squadron | threat | Token-focused board presence.
- 1 Threefold Thunderhulk | threat | Token-focused board presence.
- 1 Halo Fountain | wincon | Dedicated finisher for the established board.
- 1 Luck Bobblehead | wincon | Dedicated finisher option.
- 1 Sword of Body and Mind | wincon | Dedicated finisher option.
- 1 Basri, Tomorrow's Champion | interaction | Protective interaction for the board plan.
- 1 Lena, Selfless Champion | interaction | Protective interaction for the board plan.
- 1 Mage's Attendant | interaction | Creature-based interaction.
- 1 Rootborn Defenses | interaction | Protective interaction for the board plan.
- 1 Squad Commander | interaction | Interactive support for the board plan.
- 1 Teyo, the Shieldmage | interaction | Protective interaction for the board plan.

</details>

### 13. owned first, and the commander is not owned

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 230 names.

Commander: Karlov of the Ghost Council.

Grade: baseline, score 0.06, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $142.02 the whole deck.

**Summary:** Karlov leads a white-black lifegain strategy that develops a resilient creature board, protects its key pieces with Equipment and protective spells, and keeps opposing boards under control with efficient removal. The deck wins by turning its lifegain support into sustained pressure from its creatures, then closing with its marked finishers. It gives up some speed for a basic-land mana base and a combat-centered plan, so preserving Karlov and deploying threats in measured waves is important.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck draws more than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core is a coherent Karlov build — lifegain triggers, equipment/protection to keep him alive, and removal — but the lifegain-trigger density is thin for a commander that depends on it, and three board wipes pull against the summary's "resilient creature board" plan.
- PLAN theme_fit=yes: It is a white-black Karlov lifegain Commander deck at a mid-power, no-fast-combo level consistent with bracket 3, and the all-basic mana base is at least consistent with a build-from-owned-cards constraint.
- PLAN useful_as_built=yes: 36 lands plus ten mana rocks, a real draw package, ten pieces of spot removal, and several creature finishers make this a functional, immediately playable list with a clear route to winning through combat.
- PLAN summary_honest=yes: The summary accurately names the equipment/protection package, the removal suite, and the marked finishers, and it openly discloses the all-basic mana base and slower combat plan rather than overselling the deck.
- [WARN] `not_owned`: Karlov of the Ghost Council: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 2.86 over 63 nonland cards

<details><summary>The deck list</summary>

- 21 Plains | land | Provides a reliable white mana base for Karlov’s white cards.
- 15 Swamp | land | Provides a reliable black mana base for Karlov’s black cards.
- 1 Sol Ring | ramp | Provides efficient early acceleration.
- 1 Arcane Signet | ramp | Provides flexible color fixing and acceleration.
- 1 Chromatic Lantern | ramp | Provides color fixing alongside mana acceleration.
- 1 Commander's Sphere | ramp | Provides flexible mana acceleration.
- 1 Fellwar Stone | ramp | Provides inexpensive mana acceleration.
- 1 Relic of Legends | ramp | Provides mana acceleration for the creature-heavy plan.
- 1 Springleaf Drum | ramp | Provides low-cost mana acceleration.
- 1 Sword of the Animist | ramp | Provides ramp while supporting the combat plan.
- 1 Thought Vessel | ramp | Provides mana acceleration.
- 1 Wayfarer's Bauble | ramp | Provides early land-based mana acceleration.
- 1 Buster Sword | draw | Provides a marked source of card draw.
- 1 Call of the Ring | draw | Provides a marked source of card draw.
- 1 Exemplar of Light | draw | Provides card draw on a lifegain-themed creature.
- 1 Idol of Oblivion | draw | Provides a dedicated card-draw artifact.
- 1 Inspiring Overseer | draw | Provides card draw on a white creature.
- 1 Instant Ramen | draw | Provides a marked source of card draw.
- 1 Lembas | draw | Provides a marked source of card draw.
- 1 Mask of Memory | draw | Provides card draw through combat.
- 1 Night's Whisper | draw | Provides efficient card draw.
- 1 Painful Truths | draw | Provides a marked source of card draw.
- 1 Puresteel Paladin | draw | Provides card draw while fitting the Equipment package.
- 1 Skullclamp | draw | Provides efficient card draw.
- 1 Tome of Legends | draw | Provides recurring card draw for the commander plan.
- 1 Wall of Omens | draw | Provides card draw on an early defensive creature.
- 1 Bastion Protector | interaction | Provides protection-focused interaction for Karlov.
- 1 Boromir, Warden of the Tower | interaction | Provides a creature-based interaction piece.
- 1 Champion's Helm | interaction | Provides protective Equipment interaction.
- 1 Clever Concealment | interaction | Provides broad protective interaction.
- 1 Darksteel Plate | interaction | Provides durable protective interaction.
- 1 Lightning Greaves | interaction | Provides efficient protective Equipment interaction.
- 1 Swiftfoot Boots | interaction | Provides protective Equipment interaction.
- 1 Unbreakable Formation | interaction | Provides protection for the creature board.
- 1 Banishing Light | removal | Provides versatile permanent removal.
- 1 Bitter Triumph | removal | Provides efficient black removal.
- 1 Crib Swap | removal | Provides creature removal.
- 1 Dismember | removal | Provides low-cost creature removal.
- 1 Dispatch | removal | Provides efficient removal.
- 1 Fatal Push | removal | Provides efficient creature removal.
- 1 Generous Gift | removal | Provides flexible permanent removal.
- 1 Get Lost | removal | Provides flexible removal.
- 1 Infernal Grasp | removal | Provides dependable creature removal.
- 1 Swords to Plowshares | removal | Provides efficient creature removal.
- 1 Angel of Vitality | synergy | Supports the deck’s lifegain synergy.
- 1 Compassionate Healer | synergy | Supports the deck’s lifegain synergy.
- 1 Kor Firewalker | synergy | Supports the deck’s lifegain synergy.
- 1 Light of Promise | synergy | Supports Karlov’s lifegain-focused plan.
- 1 Rosie Cotton of South Lane | synergy | Supports the deck’s lifegain synergy.
- 1 Second Breakfast | synergy | Supports the deck’s lifegain synergy.
- 1 Well-Worn Spatula | synergy | Supports the deck’s Equipment and lifegain themes.
- 1 Angel of Invention | threat | Provides a white creature threat.
- 1 Bill the Pony | threat | Provides an on-theme creature threat.
- 1 Dawnhand Eulogist | threat | Provides a marked finisher threat.
- 1 Gollum, Silent Slinker // Meager Meal | threat | Provides a black creature threat.
- 1 Minwu, White Mage | threat | Provides a white creature threat.
- 1 Shattered Angel | threat | Provides a lifegain-aligned creature threat.
- 1 Sneering Shadewriter | threat | Provides a marked finisher threat.
- 1 Victory's Herald | threat | Provides a high-impact white creature threat.
- 1 Frodo, Sauron's Bane | wincon | Provides a marked finisher for closing games.
- 1 Grave Venerations | wincon | Provides a marked finisher for closing games.
- 1 Lyra Dawnbringer | wincon | Provides a marked Angel finisher.
- 1 Austere Command | wipe | Provides a flexible board wipe.
- 1 Fumigate | wipe | Provides a lifegain-aligned board wipe.
- 1 Vanquish the Horde | wipe | Provides an efficient creature board wipe.

</details>

### 14. the user delegates the commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 247 names.

Commander: Atarka, World Render.

Grade: bad, score 0.04, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $707.22 to buy, $707.22 the whole deck.

**Summary:** Atarka, World Render leads a Dragon-heavy ramp deck that builds its mana and card flow before applying pressure with a steady stream of Dragon threats. Dedicated Dragon win conditions and finishers give the deck strong ways to close once it establishes a board, while removal, wipes, and protective interaction help preserve that position. The tradeoff is a focused, creature-centered plan that is most effective when its mana development and Dragon board come together.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. Many of the cards appear in no top list, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Dragon cost reducers, ramp, card draw off big creatures, protection, and a top end of dragons all feed one plan of deploying Atarka and swinging with double-striking dragons.
- PLAN theme_fit=yes: It is a Commander dragon-tribal deck with a builder-chosen dragon commander at a reasonable mid-power, no-combo bracket 3 level.
- PLAN useful_as_built=yes: 36 lands with solid Gruul fixing plus cost reducers and mana rocks support the high curve, and there are multiple dragon-based ways to close the game.
- PLAN summary_honest=yes: It accurately describes a ramp-into-dragons creature deck with dragon payoffs, some removal and protection, and honestly flags the creature-dependent tradeoff, though calling Thundermaw Hellkite a wipe is a mild stretch.
- [INFO] `curve_summary`: average mana value 3.40 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Cavern of Souls | land | Dragon-focused land-base support.
- 1 Cinder Glade | land | Red-green land-base support.
- 1 City of Brass | land | Flexible land-base support.
- 1 Command Tower | land | Flexible land-base support.
- 1 Copperline Gorge | land | Red-green land-base support.
- 1 Cragcrown Pathway // Timbercrown Pathway | land | Flexible red-green land-base support.
- 1 Exotic Orchard | land | Flexible land-base support.
- 1 Fabled Passage | land | Land-base fixing support.
- 1 Fire-Lit Thicket | land | Red-green land-base support.
- 1 Game Trail | land | Red-green land-base support.
- 1 Karplusan Forest | land | Red-green land-base support.
- 1 Mana Confluence | land | Flexible land-base support.
- 1 Prismatic Vista | land | Land-base fixing support.
- 1 Rockfall Vale | land | Red-green land-base support.
- 1 Rootbound Crag | land | Red-green land-base support.
- 1 Spire Garden | land | Red-green land-base support.
- 1 Stomping Ground | land | Red-green land-base support.
- 1 Taiga | land | Red-green land-base support.
- 9 Forest | land | Reliable green land-base support.
- 9 Mountain | land | Reliable red land-base support.
- 1 Carnelian Orb of Dragonkind | ramp | Dragon-focused mana development.
- 1 Dragon's Hoard | ramp | Dragon-focused mana development.
- 1 Dragonstorm Globe | ramp | Mana development for the Dragon plan.
- 1 Ganax, Astral Hunter | ramp | A Dragon ramp piece.
- 1 Jade Orb of Dragonkind | ramp | Dragon-focused mana development.
- 1 Mox Jasper | ramp | Early mana development.
- 1 Orb of Dragonkind | ramp | Dragon-focused mana development.
- 1 Reckless Barbarian | ramp | Early mana development.
- 1 Sarkhan, Fireblood | ramp | Mana development for the Dragon plan.
- 1 Scaled Nurturer | ramp | Early Dragon-focused mana development.
- 1 Avaricious Dragon | draw | A Dragon source of card flow.
- 1 Beast Whisperer | draw | Creature-focused card flow.
- 1 Demand Answers | draw | Efficient card flow.
- 1 Dragonborn Champion | draw | A Dragon-themed card-flow piece.
- 1 Elemental Bond | draw | Creature-focused card flow.
- 1 Faithless Looting | draw | Early card flow.
- 1 Garruk's Uprising | draw | Large-creature card flow.
- 1 Guardian Project | draw | Creature-focused card flow.
- 1 Harmonize | draw | Straightforward card flow.
- 1 Return of the Wildspeaker | draw | Creature-focused card flow.
- 1 Sylvan Library | draw | Early card flow.
- 1 Heroic Intervention | interaction | Protective interaction.
- 1 Inspiring Call | interaction | Protective interaction for the creature plan.
- 1 Legolas's Quick Reflexes | interaction | Low-cost protective interaction.
- 1 Lightning Greaves | interaction | Protection support for key creatures.
- 1 Mithril Coat | interaction | Protection support for key creatures.
- 1 Snakeskin Veil | interaction | Low-cost protective interaction.
- 1 Swiftfoot Boots | interaction | Protection support for key creatures.
- 1 Tamiyo's Safekeeping | interaction | Protective interaction.
- 1 Draconic Roar | removal | Efficient Dragon-themed removal.
- 1 Dragon's Fire | removal | Efficient Dragon-themed removal.
- 1 Glorybringer | removal | A Dragon removal piece.
- 1 Invasion of Tarkir // Defiant Thundermaw | removal | Dragon-themed removal.
- 1 Magmatic Hellkite | removal | A Dragon removal piece.
- 1 Molten Exhale | removal | Efficient removal.
- 1 Piercing Exhale | removal | Efficient removal.
- 1 Scourge of Valkas | removal | A Dragon removal piece.
- 1 Spit Flame | removal | Dragon-themed removal.
- 1 Acolyte of Bahamut | synergy | Dragon-plan synergy support.
- 1 Crucible of Fire | synergy | Dragon-plan synergy support.
- 1 Dragonlord's Servant | synergy | Dragon-plan synergy support.
- 1 Dragonspeaker Shaman | synergy | Dragon-plan synergy support.
- 1 Firespitter Whelp | synergy | Dragon-plan synergy support.
- 1 Sarkhan's Triumph | synergy | Dragon-plan synergy support.
- 1 The Dragon-Kami Reborn // Dragon-Kami's Egg | synergy | Dragon-plan synergy support.
- 1 Ambitious Dragonborn | threat | A Dragon threat for board pressure.
- 1 Backdraft Hellkite | threat | A Dragon threat for board pressure.
- 1 Caldera Pyremaw | threat | A Dragon threat for board pressure.
- 1 Dragon Broodmother | threat | A Dragon threat for board pressure.
- 1 Dream Pillager | threat | A Dragon threat for board pressure.
- 1 Hoarding Dragon | threat | A Dragon threat for board pressure.
- 1 Manaform Hellkite | threat | A Dragon threat for board pressure.
- 1 Mirrorwing Dragon | threat | A Dragon threat for board pressure.
- 1 Sagu Wildling // Roost Seek | threat | A Dragon threat for board pressure.
- 1 Stormbreath Dragon | threat | A Dragon threat for board pressure.
- 1 Territorial Hellkite | threat | A Dragon threat for board pressure.
- 1 Vengeful Ancestor | threat | A Dragon threat for board pressure.
- 1 Lathliss, Dragon Queen | wincon | A Dragon win condition.
- 1 Utvara Hellkite | wincon | A Dragon win condition.
- 1 Wrathful Red Dragon | wincon | A Dragon win condition.
- 1 Breath Weapon | wipe | A Dragon-friendly board wipe.
- 1 Draconic Intervention | wipe | A Dragon-themed board wipe.
- 1 Thundermaw Hellkite | wipe | A Dragon board-wipe option.

</details>

### 15. delegated commander, owned first

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 230 names.

Commander: Astarion, the Decadent.

Grade: baseline, score 0.06, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $135.32 the whole deck.

**Summary:** This lifegain Commander deck develops a resilient board of creatures and synergistic permanents, using its mana acceleration and card-advantage pieces to keep deploying meaningful threats. It wins by building toward its marked finishing cards and win conditions while using targeted answers and board wipes to keep the table manageable. It gives up explosive, highly streamlined play for a fairer board-facing game that relies on establishing and preserving its battlefield presence.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The land count sits above the norm of the format, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core is a coherent white-black lifegain-payoff board deck (Angel of Vitality, Rosie Cotton, Shattered Angel, Compassionate Healer, Lyra, Victory's Herald) with ramp, removal, and wipes, but an equipment/commander-protection subpackage (Puresteel Paladin, Buster Sword, Aettir and Priwen, Bastion Protector, Champion's Helm, Greaves/Boots) and a few off-plan bodies pull mildly toward a different voltron axis.
- PLAN theme_fit=yes: The requester asked for a lifegain Commander deck at bracket 2 with the builder choosing the commander, and Astarion, the Decadent with a life-triggered payoff shell at a fair, precon-level power fits that ask.
- PLAN useful_as_built=yes: 37 lands plus ten ramp pieces, a low-to-mid curve, eleven draw sources, seven removal spells and three wipes, and several combat-based finishers make it immediately castable and able to close games, though the land count is a bit high.
- PLAN summary_honest=partly: Nothing it claims is absent (ramp, draw, targeted removal, three wipes, finishers), but it is so generic that it never mentions the deck's actual lifegain engine, which is the main thing the list does.
- [WARN] `not_owned`: Astarion, the Decadent: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.16 over 62 nonland cards

<details><summary>The deck list</summary>

- 19 Plains | land | Provides reliable white mana for the deck.
- 18 Swamp | land | Provides reliable black mana for the deck.
- 1 Arcane Signet | ramp | Provides efficient color fixing and mana acceleration.
- 1 Astral Cornucopia | ramp | Adds flexible mana production.
- 1 Chromatic Lantern | ramp | Supports broad mana fixing and acceleration.
- 1 Commander's Sphere | ramp | Provides dependable mana acceleration.
- 1 Fellwar Stone | ramp | Adds inexpensive mana acceleration.
- 1 Giada, Font of Hope | ramp | Supports the deck's mana development.
- 1 Inherited Envelope | ramp | Contributes to the ramp package.
- 1 Lotho, Corrupt Shirriff | ramp | Provides mana acceleration from a black permanent.
- 1 Relic of Legends | ramp | Adds flexible mana production.
- 1 Sol Ring | ramp | Provides efficient early mana acceleration.
- 1 Buster Sword | draw | Supplies a source of card advantage.
- 1 Call of the Ring | draw | Provides ongoing card advantage.
- 1 Cirith Ungol Patrol | draw | Adds a black card-advantage piece.
- 1 Exemplar of Light | draw | Adds card advantage within the creature suite.
- 1 Idol of Oblivion | draw | Provides artifact-based card advantage.
- 1 Inspiring Overseer | draw | Adds card advantage to the creature plan.
- 1 Lembas | draw | Provides compact artifact-based card advantage.
- 1 Mask of Memory | draw | Adds repeatable equipment-based card advantage.
- 1 Night's Whisper | draw | Provides efficient black card advantage.
- 1 Puresteel Paladin | draw | Supports the equipment package while drawing cards.
- 1 Wall of Omens | draw | Provides early defense and card advantage.
- 1 Bastion Protector | interaction | Helps protect an important legendary creature.
- 1 Champion's Helm | interaction | Protects a key legendary creature.
- 1 Clever Concealment | interaction | Provides a defensive interaction option.
- 1 Lightning Greaves | interaction | Protects an important creature efficiently.
- 1 Swiftfoot Boots | interaction | Adds another reliable protection piece.
- 1 Unbreakable Formation | interaction | Provides a creature-focused protection option.
- 1 Banishing Light | removal | Provides flexible permanent removal.
- 1 Bitter Triumph | removal | Provides efficient black removal.
- 1 Crib Swap | removal | Provides creature removal.
- 1 Dismember | removal | Provides low-cost creature removal.
- 1 Dispatch | removal | Provides efficient creature removal.
- 1 Fatal Push | removal | Provides efficient black removal.
- 1 Generous Gift | removal | Answers a broad range of opposing permanents.
- 1 Aettir and Priwen | synergy | Supports the deck's synergy package.
- 1 Aerith Gainsborough | synergy | Supports the lifegain-focused synergy plan.
- 1 Angel of Vitality | synergy | Adds a lifegain-focused synergy creature.
- 1 Compassionate Healer | synergy | Supports the deck's core synergy plan.
- 1 Kor Firewalker | synergy | Adds another lifegain-focused synergy piece.
- 1 Light of Promise | synergy | Supports the lifegain synergy package.
- 1 Night Nurse, Healer of Heroes | synergy | Adds a themed synergy creature.
- 1 Rosie Cotton of South Lane | synergy | Supports the creature-based synergy plan.
- 1 Second Breakfast | synergy | Provides a synergy-focused instant.
- 1 Well-Worn Spatula | synergy | Supports the deck's equipment and synergy package.
- 1 Angel of Invention | threat | Provides a white creature threat.
- 1 Bill the Pony | threat | Adds a creature threat to develop the board.
- 1 Dawnhand Eulogist | threat | Provides a substantial finisher-marked threat.
- 1 Foggy Swamp Hunters | threat | Adds a creature threat to the board plan.
- 1 Gollum, Silent Slinker // Meager Meal | threat | Provides a black creature threat.
- 1 Invisible Woman, Sue Storm | threat | Adds a legendary creature threat.
- 1 Minwu, White Mage | threat | Provides another white creature threat.
- 1 Rabaroo Troop | threat | Adds to the deck's board presence.
- 1 Reaping Willow | threat | Provides a larger creature threat.
- 1 Rooftop Percher | threat | Adds a creature threat to pressure opponents.
- 1 Shattered Angel | threat | Provides an Angel threat for the deck.
- 1 Victory's Herald | threat | Provides a top-end creature threat.
- 1 Frodo, Sauron's Bane | wincon | Provides a marked finishing route.
- 1 Grave Venerations | wincon | Provides a marked finishing route.
- 1 Lyra Dawnbringer | wincon | Provides a marked finishing creature.
- 1 Austere Command | wipe | Provides a flexible board reset.
- 1 Fumigate | wipe | Provides a creature-focused board reset.
- 1 Vanquish the Horde | wipe | Provides an additional board reset.

</details>

### 16. a tight budget, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 231 names.

Commander: Denethor, Ruling Steward.

Grade: baseline, score 0.31, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $7.44 to buy, $46.54 the whole deck.

**Summary:** This white-black aristocrats deck builds a creature board, uses sacrifice-focused synergy to turn that board into pressure, and keeps resources flowing with a substantial draw and ramp package. It closes games through its dedicated finisher package while using targeted answers, protection, and board wipes to keep opposing plans contained. The tradeoff is a deliberately fair, board-centered pace: the deck is strongest when it has time to establish creatures and supporting permanents rather than racing ahead immediately.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The deck makes less mana on turn four than the norm, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: Most of the creature slots are genuine sacrifice payoffs and outlets (Bartolome, Woe Strider, Old Flitterfang, Erebos, Skullclamp, Deadly Dispute), but the deck also runs three board wipes and Archfiend of Ifnir alongside a plan that wants to keep a wide creature board, and it is notably short on token/expendable-body generators to feed the sacrifice engines.
- PLAN theme_fit=yes: It is a mono-white/black budget-priced aristocrats Commander deck under a low-power commander, matching the requested theme, bracket, and price ceiling.
- PLAN useful_as_built=yes: 38 lands plus several cheap rocks, a low curve, plentiful removal and card draw make this a playable, functional list out of the box even if its win routes are grindy.
- PLAN summary_honest=partly: The draw, removal, protection and wipe claims check out, but the 'dedicated finisher package' is three cards, only one of which (Shilgengar) is a real closer, and the summary never mentions the wipes' tension with its own board-centered plan.
- [WARN] `not_owned`: Woe Strider: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Yahenni, Undying Partisan: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Aron, Benalia's Ruin: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ayli, Eternal Pilgrim: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Baron Bertram Graywater: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bartolomé del Presidio: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Elder Arthur Maxson: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Lord Skitter's Butcher: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Old Flitterfang: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Skullport Merchant: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Umbral Collar Zealot: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vampire Gourmand: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Shadowheart, Dark Justiciar: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Erebos, Bleak-Hearted: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Shilgengar, Sire of Famine: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 2.87 over 62 nonland cards

<details><summary>The deck list</summary>

- 14 Plains | land | Basic white source for the two-color mana base.
- 14 Swamp | land | Basic black source for the two-color mana base.
- 1 Command Tower | land | Flexible color source for the commander deck.
- 1 Exotic Orchard | land | Flexible color source that supports both deck colors.
- 1 City of Brass | land | Flexible color source for consistently casting colored spells.
- 1 Grand Coliseum | land | Additional multicolor land for the mana base.
- 1 Path of Ancestry | land | A color-fixing land that supports the creature-focused list.
- 1 Secluded Courtyard | land | Creature-oriented color fixing.
- 1 Unclaimed Territory | land | Creature-oriented color fixing.
- 1 Spire of Industry | land | Artifact-supported color fixing.
- 1 Evolving Wilds | land | Flexible access to either basic land type.
- 1 Sol Ring | ramp | Low-cost acceleration that helps develop the board early.
- 1 Arcane Signet | ramp | Reliable color-fixing acceleration.
- 1 Fellwar Stone | ramp | Additional efficient mana acceleration.
- 1 Chromatic Lantern | ramp | Mana acceleration and broad color support.
- 1 Commander's Sphere | ramp | Flexible mana rock for the commander deck.
- 1 Deadly Dispute | ramp | Listed sacrifice-friendly acceleration for the deck's plan.
- 1 Explorer's Scope | ramp | Low-cost ramp support for the creature board.
- 1 Inherited Envelope | ramp | Additional artifact-based mana acceleration.
- 1 Relic of Legends | ramp | Mana support that works well with the legendary commander and creatures.
- 1 White Auracite | ramp | Cheap additional ramp for the deck.
- 1 Ahriman | draw | Creature-based card draw for maintaining resources.
- 1 Wall of Omens | draw | Early creature that supplies card draw.
- 1 Inspiring Overseer | draw | Creature-based card draw that adds to the board.
- 1 Nasty End | draw | Efficient black card draw.
- 1 Painful Truths | draw | Dedicated draw spell for refilling the hand.
- 1 Skullclamp | draw | Draw engine that fits a creature-focused strategy.
- 1 Tome of Legends | draw | Repeatable draw support around the commander.
- 1 Stone of Erech | draw | Artifact-based card draw support.
- 1 Mask of Memory | draw | Equipment-based draw for attacking creatures.
- 1 Idol of Oblivion | draw | Artifact draw support for the deck's board plan.
- 1 Lembas | draw | Low-cost artifact card draw.
- 1 Bastion Protector | interaction | Creature-based protection for important permanents.
- 1 Frontline Medic | interaction | Protective creature interaction for the board.
- 1 Gift of Immortality | interaction | Protection for a key creature in the sacrifice-focused core.
- 1 Swiftfoot Boots | interaction | Efficient protection for the commander or a key creature.
- 1 Take Up the Shield | interaction | Low-cost protective interaction.
- 1 Unbreakable Formation | interaction | Board-protection interaction for a creature-heavy deck.
- 1 Banishing Light | removal | Flexible permanent removal.
- 1 Bitter Triumph | removal | Efficient single-target removal.
- 1 Claim the Precious | removal | Black spot removal for problematic creatures.
- 1 Crib Swap | removal | Creature removal that fits the deck's colors.
- 1 Fatal Push | removal | Low-cost targeted removal.
- 1 Infernal Grasp | removal | Reliable unconditional creature removal.
- 1 Stroke of Midnight | removal | Flexible white permanent removal.
- 1 Arcade Cabinet | synergy | Listed synergy piece for the deck's central plan.
- 1 Gollum the Abandoned | synergy | Listed synergy creature for the sacrifice-focused core.
- 1 Gollum, Patient Plotter | synergy | Listed synergy creature that reinforces the core plan.
- 1 Gríma Wormtongue | synergy | Listed synergy piece for the creature package.
- 1 Heirloom Auntie | synergy | Listed synergy creature for the deck's board plan.
- 1 Joo Dee, One of Many | synergy | Listed synergy creature for the core strategy.
- 1 Nimble Hobbit | synergy | Listed synergy creature that helps fill out the board.
- 1 Phantom Train | synergy | Listed artifact synergy for the deck.
- 1 Woe Strider | synergy | Sacrifice-oriented creature that supports the aristocrats plan.
- 1 Yahenni, Undying Partisan | synergy | Sacrifice-oriented creature that supports the aristocrats plan.
- 1 Aron, Benalia's Ruin | threat | Creature threat that contributes to the aristocrats board.
- 1 Ayli, Eternal Pilgrim | threat | Creature threat suited to the sacrifice-focused creature package.
- 1 Baron Bertram Graywater | threat | Creature threat that adds to the deck's board presence.
- 1 Bartolomé del Presidio | threat | Creature threat for the sacrifice-focused board.
- 1 Elder Arthur Maxson | threat | Creature threat that broadens the board plan.
- 1 Lord Skitter's Butcher | threat | Creature threat that adds pressure while supporting the creature count.
- 1 Old Flitterfang | threat | Creature threat for the deck's black creature package.
- 1 Skullport Merchant | threat | Creature threat that fits the value-oriented aristocrats shell.
- 1 Umbral Collar Zealot | threat | Creature threat for the black creature package.
- 1 Vampire Gourmand | threat | Low-cost creature threat that helps establish the board.
- 1 Shadowheart, Dark Justiciar | threat | Creature threat that supports the sacrifice-focused plan.
- 1 Erebos, Bleak-Hearted | threat | Durable black threat for the deck's creature strategy.
- 1 Archfiend of Ifnir | wincon | Listed finisher that gives the deck a high-impact closing threat.
- 1 Grave Venerations | wincon | Listed finisher for closing games through the core plan.
- 1 Shilgengar, Sire of Famine | wincon | Listed finisher that provides another top-end closer.
- 1 Austere Command | wipe | Flexible board wipe for resetting difficult boards.
- 1 Dusk // Dawn | wipe | Board wipe that complements a creature-focused build.
- 1 Fumigate | wipe | Reliable full-board reset when pressure gets out of hand.

</details>

### 17. upgrade a precon, owned first

Format: Commander. Theme: goblins. Pool: owned_first. Shortlist: 198 names.

Commander: Zada, Hedron Grinder.

Grade: typical, score 0.60, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: yes, for deck_size, precon_share, profile_off_band, profile_off_band, profile_off_band, profile_off_band, profile_off_band, profile_off_band, profile_off_band, profile_off_band. Block findings: 0.

Cost: $84.15 to buy, $276.98 the whole deck.

**Summary:** A mono-red Goblin Storm deck that develops a dense board of Goblins, converts that board into bursts of mana, and chains inexpensive spells for amplified Zada turns. Token production, tribal lords, sacrifice damage, and combat pumps provide several overlapping routes to overwhelm the table, while efficient damage and board-clearing effects preserve momentum.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The cards pair the way the top lists pair them, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every card feeds the same engine — goblin token production, tribal lords, cheap single-target cantrips/pumps that Zada copies, rituals and Grapeshot/Empty the Warrens storm payoffs, plus sac-damage outlets — with only Blasphemous Act and Earthquake pulling slightly against the wide board.
- PLAN theme_fit=yes: It is exactly the asked-for upgrade: mono-red Zada Goblin Storm in Commander with a tuned but not degenerate power level consistent with bracket 3 (no fast mana beyond Sol Ring/medallion, no infinite combo lines).
- PLAN useful_as_built=yes: 37 lands in a mono-color deck with a low curve, plenty of ramp and card draw, and several independent win routes (Grapeshot/Impact Tremors drain, Siege-Gang sac damage, Zada-pumped alpha strikes) make it immediately playable.
- PLAN summary_honest=partly: The described token swarm, ritual mana, Zada spell-chaining, lords, sacrifice damage and pumps are all genuinely present, but calling the sweepers effects that 'preserve momentum' misrepresents Blasphemous Act and Earthquake, which mostly erase the deck's own goblins.
- [INFO] `curve_summary`: average mana value 2.54 over 61 nonland cards
- [INFO] `precon_cards_restored`: the deck was 2 cards short of the Goblin Storm precon share, so the builder put 2 cards back
- [INFO] `bracket_cut`: to hold bracket 3, the builder cut 1 card: Storm-Kiln Artist (the combo Storm-Kiln Artist + Haze of Rage), and added 1 basic land

<details><summary>The deck list</summary>

- 27 Mountain | land | Primary red mana base.
- 1 Arena of Glory | land | Precon utility land with red access.
- 1 Castle Embereth | land | Precon land that supports attacking Goblins.
- 1 Den of the Bugbear | land | Precon creature-land threat.
- 1 Forgotten Cave | land | Precon cycling land.
- 1 Fountainport | land | Precon utility land.
- 1 Hidden Volcano | land | Precon red utility land.
- 1 Kher Keep | land | Precon token-making utility land.
- 1 Reliquary Tower | land | Precon hand-size utility land.
- 1 Shinka, the Bloodsoaked Keep | land | Precon legendary combat utility land.
- 1 War Room | land | Precon card-advantage utility land.
- 1 Ancestors' Aid | synergy | Copies across the board with Zada and improves combat.
- 1 Battle Hymn | ramp | Converts a Goblin board into explosive mana.
- 1 Blasphemous Act | wipe | Efficient emergency creature reset.
- 1 Boggart Shenanigans | removal | Turns Goblin deaths into reach.
- 1 Brightstone Ritual | ramp | Burst mana from the Goblin army.
- 1 Chaos Warp | removal | Flexible answer to troublesome permanents.
- 1 Conspicuous Snoop | synergy | Goblin-focused card access and tribal payoff.
- 1 Crimson Wisps | draw | Cheap Zada cantrip that grants haste.
- 1 Dragon Fodder | synergy | Produces multiple bodies for storm and tribal payoffs.
- 1 Empty the Warrens | synergy | Builds a large Goblin board after chaining spells.
- 1 Expedite | draw | Cheap cantrip that gives a key creature haste.
- 1 Fable of the Mirror-Breaker // Reflection of Kiki-Jiki | ramp | Provides mana, selection, and a Goblin-copying threat.
- 1 Faithless Looting | draw | Efficient card selection for spell-heavy turns.
- 1 Fists of Flame | draw | Zada cantrip that can create a lethal combat swing.
- 1 Gempalm Incinerator | removal | Goblin-scaled creature removal.
- 1 General Kreat, the Boltbringer | synergy | Tribal payoff that strengthens the Goblin plan.
- 1 Goblin Bushwhacker | synergy | Turns token-heavy boards into immediate attackers.
- 1 Goblin Chieftain | synergy | Haste and power boost for the Goblin army.
- 1 Goblin Dark-Dwellers | threat | Recasts a useful instant or sorcery while adding pressure.
- 1 Goblin Lackey | synergy | Accelerates Goblins onto the battlefield through combat.
- 1 Goblin Matron | synergy | Finds the most useful Goblin for the situation.
- 1 Goblin Negotiation | removal | Tribal-flavored answer to opposing threats.
- 1 Goblin Trashmaster | removal | Goblin lord that also answers artifacts.
- 1 Goblin Warchief | synergy | Reduces Goblin costs and grants haste.
- 1 Grapeshot | removal | Storm payoff that can clear small creatures or finish opponents.
- 1 Great Train Heist | wincon | High-impact combat finisher for a wide board.
- 1 Grenzo, Havoc Raiser | synergy | Rewards Goblin combat with disruption and impulse access.
- 1 Haze of Rage | synergy | Repeatable storm combat payoff for a token board.
- 1 Howlsquad Heavy | ramp | Goblin body that helps generate mana.
- 1 Idol of Oblivion | draw | Reliable card draw alongside frequent token production.
- 1 Impact Tremors | synergy | Converts every token wave into direct damage.
- 1 Impulsive Pilferer | ramp | Early Goblin that leaves behind mana when sacrificed.
- 1 Krenko's Command | synergy | Efficiently adds multiple Goblins to the board.
- 1 Krenko, Mob Boss | threat | Rapidly snowballs a Goblin board into a major threat.
- 1 Mana Geyser | ramp | Explosive midgame mana for a storm turn.
- 1 Mogg War Marshal | synergy | Multiple Goblin bodies from one card.
- 1 Pashalik Mons | removal | Turns Goblin sacrifices and deaths into damage.
- 1 Past in Flames | synergy | Lets a stocked graveyard fuel a second spell chain.
- 1 Quest for the Goblin Lord | synergy | Tribal anthem that rewards steady Goblin deployment.
- 1 Roaming Throne | synergy | Doubles key Goblin triggered abilities.
- 1 Ruby Medallion | ramp | Reduces the cost of the deck's dense red spell suite.
- 1 Rundvelt Hordemaster | synergy | Goblin anthem with card advantage when the army dies.
- 1 Searslicer Goblin | synergy | Tribal creature that contributes to wide-board pressure.
- 1 Seething Song | ramp | Burst mana for pivotal multi-spell turns.
- 1 Siege-Gang Commander | removal | Creates a board and converts Goblins into removal.
- 1 Siege-Gang Lieutenant | removal | Adds Goblins while providing sacrifice-based damage.
- 1 Skirk Prospector | ramp | Converts Goblins into mana for explosive turns.
- 1 Skullclamp | draw | Refills the hand from disposable Goblin tokens.
- 1 Sol Ring | ramp | Fast colorless acceleration.
- 1 Swiftfoot Boots | interaction | Protects the commander while enabling immediate value.
- 1 Throne of Eldraine | ramp | Mana acceleration with useful late-game flexibility.
- 1 Vandalblast | wipe | Efficient artifact answer with a sweeping overload mode.
- 1 Vanquisher's Banner | draw | Tribal anthem and ongoing Goblin card draw.
- 1 Warren Torchmaster | synergy | Goblin payoff that supports aggressive token combat.
- 1 Wild Ride | synergy | Combat-focused spell that supports the aggressive Goblin theme.
- 1 Witch's Mark | draw | Card filtering that can enhance a combat turn.
- 1 Arcane Signet | ramp | Reliable early acceleration into the commander and board development.
- 1 Champion's Helm | interaction | Protects the legendary commander from targeted disruption.
- 1 Darksteel Plate | interaction | Keeps a key creature alive through damage and destroy effects.
- 1 Earthquake | wincon | Scalable direct-damage finisher that also clears smaller boards.
- 1 Spinerock Knoll | removal | the deck upgrades a precon, and this card is one the precon holds
- 1 Frontline Heroism | interaction | the deck upgrades a precon, and this card is one the precon holds

</details>

### 18. upgrade a precon, any card

Format: Commander. Theme: turtles. Pool: any_card. Shortlist: 311 names.

Commander: Heroes in a Half Shell.

Grade: baseline, score 0.35, model 20260914T154223Z.

Cards: 105 main, 0 sideboard. Repair turn: yes, for game_changer_limit, profile_off_band, profile_off_band, profile_off_band, profile_off_band, profile_off_band. Block findings: 2.

Cost: $504.32 to buy, $504.32 the whole deck.

**Summary:** A five-color Turtle Power deck built around the heroes, allies, inventions, and mutant mayhem of the Turtle Power roster. It develops a broad mana base, builds a resilient creature board, and mixes character-driven combat with artifact utility, card advantage, targeted answers, and dramatic finishing attacks.

The quality model grades this deck at the precon baseline against the top lists of the format: many lands enter tapped, and that raises the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: There is a real backbone of toughness/counters payoffs (Bedrock Tortoise, Fecund Greenshell, Ancient Adamantoise, Corpsejack Menace, Vigor) plus the TMNT character shell, but a large share of the list is unconnected five-color goodstuff and one-off gadget synergy cards that pull toward generic value rather than one plan.
- PLAN theme_fit=partly: It is the requested five-color Turtle Power upgrade with the precon commander and theme, but it packs Rhystic Study, Mystic Remora, Smothering Tithe, The One Ring and Fierce Guardianship, well past the bracket 3 ceiling the person named.
- PLAN useful_as_built=yes: Roughly 37 lands with heavy five-color fixing, rainbow lands, Chromatic Lantern and mana creatures support the curve, and there are plenty of creature threats, wraths and finishers to actually close games, even if many lands enter tapped.
- PLAN summary_honest=partly: The described mana base, creature board, artifacts, card advantage and removal all exist in the list, but the summary stays vague and omits that the deck runs five game-changer-tier staples, which is relevant to the bracket the person asked for.
- [BLOCK] `deck_size`: deck has 106 cards, the format needs exactly 100 (commander included)
- [BLOCK] `game_changer_limit`: bracket 3 allows 3 Game Changers, the deck has 4: Fierce Guardianship, Rhystic Study, Smothering Tithe, The One Ring
- [INFO] `curve_summary`: average mana value 3.38 over 68 nonland cards
- [WARN] `profile_off_band`: the draw count is 7, and bracket 3 wants 8 to 14
- [WARN] `profile_off_band`: the removal count is 5, and bracket 3 wants 6 to 12

<details><summary>The deck list</summary>

- 1 Acidic Slime | removal | Versatile Turtle-adjacent permanent answer retained from the precon.
- 1 Ancient Adamantoise | ramp | Turtle-themed mana development.
- 1 An Offer You Can't Refuse | interaction | Efficient protection against pivotal noncreature spells.
- 1 April O'Neil, Live on the Scene | draw | Card advantage that remains part of the precon’s cast.
- 1 Arcade Cabinet | synergy | Retained thematic artifact.
- 1 Arcane Signet | ramp | Reliable multicolor acceleration.
- 1 Ash Barrens | land | Fixes colors while preserving the precon mana base.
- 1 Assassin's Trophy | removal | Flexible answer to troublesome permanents.
- 1 Baxter, Fly in the Ointment | synergy | Retained supporting character.
- 1 Bedrock Tortoise | interaction | Turtle-themed stack and board interaction.
- 1 Big Apple, 3 a.m. | land | Retained precon fixing land.
- 1 Big Mother Mouser | synergy | Retained thematic artifact creature.
- 1 Biogenic Ooze | threat | Retained creature-producing threat.
- 1 Birds of Paradise | ramp | Early color fixing and acceleration.
- 1 Blasphemous Act | wipe | Efficient emergency creature reset.
- 1 Casey Jones, Back Alley Brute | threat | Retained thematic creature.
- 1 Chromatic Lantern | ramp | Fixes demanding five-color mana while accelerating.
- 1 Cinder Glade | land | Retained dual-color land.
- 1 City of Brass | land | Untapped access to every color.
- 1 Coin of Mastery | synergy | Retained precon artifact.
- 1 Command Tower | land | Best unrestricted multicolor land.
- 1 Continue? | synergy | Retained thematic spell.
- 1 Corpsejack Menace | synergy | Supports the deck’s counter-focused creature plan.
- 1 Cultivate | ramp | Stable land-based acceleration and fixing.
- 1 Dimension X Pizzasaur | wincon | Thematic high-impact closing threat.
- 1 Donatello, the Brains | synergy | Retained Turtle leader and engine piece.
- 1 Double Jump // Flying Kick | interaction | Retained thematic combat trick.
- 1 Dragonskull Summit | land | Retained dual-color land.
- 1 Electric Seaweed | synergy | Retained defensive thematic creature.
- 1 Endless Foot Assault | synergy | Retained thematic enchantment.
- 1 Escape Tunnel | land | Budget fixing with late-game evasion utility.
- 1 Esper Sentinel | draw | Efficient early card-advantage engine.
- 1 Everything Pizza | wincon | Thematic finisher for longer games.
- 1 Evolving Wilds | land | Fixes basic-color access.
- 1 Exotic Orchard | land | Usually provides broad multiplayer fixing.
- 1 Exploding Barrel | synergy | Retained thematic artifact.
- 1 Fabled Passage | land | Reliable basic-land fixing.
- 1 Fecund Greenshell | ramp | Turtle-themed acceleration that also supports sustained resources.
- 1 Fellwar Stone | ramp | Low-cost multicolor acceleration.
- 1 Fierce Guardianship | interaction | Free commander-based protection for key turns.
- 1 Foot Chopper | synergy | Retained thematic Equipment.
- 1 Grand Coliseum | land | Broad color access for the five-color shell.
- 1 Harmonize | draw | Straightforward refill spell retained from the precon.
- 1 Here Comes a New Hero! | synergy | Retained thematic spell.
- 1 Hidden Hideout | land | Retained precon fixing land.
- 1 High Score | synergy | Retained thematic enchantment.
- 1 Hinterland Harbor | land | Reliable blue-green source.
- 1 Irma, Part-Time Mutant | synergy | Retained supporting character.
- 1 Kappa Tech-Wrecker | removal | Turtle Ninja that answers artifacts and enchantments.
- 1 Krang, the All-Powerful | threat | Retained thematic artifact threat.
- 1 Leatherhead, Iron Gator | threat | Retained thematic creature.
- 1 Leonardo, the Balance | threat | Retained Turtle legend and combat presence.
- 1 Lessons from Life | draw | Retained thematic value spell.
- 1 Level Up | synergy | Retained thematic Aura.
- 1 Lita, Little Orphan Amphibian | synergy | Retained thematic Turtle legend.
- 1 Michelangelo, the Heart | synergy | Retained Turtle legend and cohesion piece.
- 1 Mole Module | synergy | Retained thematic Vehicle.
- 1 Mona Lisa, Science Geek | synergy | Retained supporting character.
- 1 Mystic Remora | draw | Early-game card advantage against spell-heavy tables.
- 1 Ninja Pizza | synergy | Retained thematic enchantment.
- 1 Path of Ancestry | land | Color fixing with incidental creature selection.
- 1 Rain-Slicked Copse | land | Retained blue-green dual land.
- 1 Raphael, the Muscle | wincon | Thematic finisher retained from the precon.
- 1 Rat King, Pale Piper | threat | Retained supporting legend.
- 1 Ray Fillet, Wave Warrior | threat | Retained thematic creature.
- 1 Rhystic Study | draw | Persistent card-advantage engine.
- 1 Roadkill Rodney | synergy | Retained thematic artifact creature.
- 1 Rocksteady, Mutant Marauder | threat | Retained thematic creature.
- 1 Rootbound Crag | land | Reliable red-green source.
- 1 Shellshock | interaction | Retained thematic instant.
- 1 Shredder, Shadow Master | threat | Retained thematic legend.
- 1 Smoldering Marsh | land | Retained black-red dual land.
- 1 Smothering Tithe | ramp | Powerful mana engine for a multiplayer game.
- 1 Sol Ring | ramp | Efficient staple acceleration retained from the precon.
- 1 Sodden Verdure | land | Retained blue-green dual land.
- 1 Special Move | interaction | Retained thematic instant.
- 1 Splinter, the Mentor | synergy | Retained thematic legend.
- 1 Steelbane Hydra | removal | Creature-based artifact and enchantment removal.
- 1 Sunken Hollow | land | Retained blue-black dual land.
- 1 Swords to Plowshares | removal | Premium low-cost creature removal.
- 1 The One Ring | interaction | Protection and substantial card flow for grindy games.
- 1 The Pride of Hull Clade | draw | Turtle-themed card-advantage payoff.
- 1 Thriving Grove | land | Flexible green-based fixing.
- 1 Thriving Isle | land | Flexible blue-based fixing.
- 1 Thriving Moor | land | Flexible black-based fixing.
- 1 Together Forever | synergy | Retained thematic resilience piece.
- 1 Tokka & Rahzar, Unsupervised | ramp | Turtle-themed mana development.
- 1 Turtle Lair | land | Retained thematic land.
- 1 Undergrowth Stadium | land | Reliable black-green source.
- 1 Vernal Fen | land | Retained black-green dual land.
- 1 Vibrant Cityscape | land | Broad five-color fixing.
- 1 Vigor | synergy | Protects and strengthens the creature board.
- 1 Voracious Hydra | threat | Flexible scalable creature retained from the precon.
- 1 Wave Goodbye | wipe | Retained thematic reset spell.
- 3 Forest | land | Green-heavy basic support.
- 3 Island | land | Blue basic support.
- 2 Plains | land | White basic support.
- 2 Swamp | land | Black basic support.
- 1 Mountain | land | Red basic support.

</details>

### 19. the Hobbit family, two colours

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 174 names.

Commander: Thranduil, the Elvenking.

Grade: baseline, score 0.05, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $504.15 to buy, $504.15 the whole deck.

**Summary:** This deck develops an Elf-centered board around Thranduil, using early mana pieces to establish its legends and creatures while draw cards keep the hand supplied. It protects that board with interaction, clears away key opposing pieces with removal and battlefield resets, then converts its board pressure into wins through its marked Troll and Witch-king finishers. The deck gives up some raw speed for a thematic creature core and a broad spread of answers.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. Many of the cards appear in no top list, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core is a broad creature-plus-removal goodstuff pile that plays to the board, but the self-damaging sweepers (Languish, Gnashing of Teeth, Raise the Palisade) work against its own small elves and the labeled "wincons" are not real finishing engines, so the pieces only loosely pull one way.
- PLAN theme_fit=yes: It is a Commander deck led by the requested Thranduil, the Elvenking, drawn almost entirely from the Hobbit set's card pool in his colors, and its mix of a few fast mana pieces with mostly midrange value cards sits comfortably at bracket 3.
- PLAN useful_as_built=yes: Thirty-seven lands with fixing and ramp support a three-color curve that is mostly cheap, and there are plenty of bodies plus removal to actually play games out.
- PLAN summary_honest=partly: It correctly describes the ramp/draw/removal spread, but it oversells an "Elf-centered" core when most creatures are spiders, bears, trolls and wraiths, and calls Troll of Khazad-dûm and the Witch-kings "finishers" when they are just big bodies.
- [INFO] `curve_summary`: average mana value 3.10 over 62 nonland cards
- [WARN] `profile_off_band`: the worst color's share of the sources it needs is 0.83, and bracket 3 wants 0.85 or more (sources of requirement: U 15.75 of 19, B 19.5 of 23, G 19 of 23)
- [INFO] `mana_pass`: the builder moved 7 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 1 Elven Passage | land | Fills a land slot in the mana base.
- 1 Elvenking's Halls | land | Fills a land slot in the mana base.
- 11 Forest | land | Provides reliable green mana in the mana base.
- 1 Mirkwood | land | Fills a land slot in the mana base.
- 8 Island | land | Provides reliable blue mana in the mana base.
- 1 Rivendell | land | Fills a land slot in the mana base.
- 11 Swamp | land | Provides reliable black mana in the mana base.
- 1 The Black Gate | land | Fills a land slot in the mana base.
- 1 The Shire | land | Fills a land slot in the mana base.
- 1 Arcane Signet | ramp | Provides efficient mana acceleration.
- 1 Delighted Halfling | ramp | Provides early mana acceleration.
- 1 Elvish Archdruid | ramp | Adds mana acceleration while fitting the Elf core.
- 1 Elvish Mystic | ramp | Provides early mana acceleration.
- 1 Mox Amber | ramp | Provides low-cost mana acceleration.
- 1 Necklace of Girion | ramp | Adds to the deck's mana development.
- 1 Orcrist, Goblin-cleaver | ramp | Adds to the deck's mana development.
- 1 Wayfarer's Bauble | ramp | Provides dependable mana development.
- 1 Wood Elves | ramp | Provides mana acceleration while supporting the Elf theme.
- 1 Beorn the Fierce | draw | Provides card draw.
- 1 Elvish Visionary | draw | Provides card draw while fitting the Elf core.
- 1 Fateful Discovery | draw | Provides card draw.
- 1 Hithlain Knots | draw | Provides card draw.
- 1 Key to the Side-Door | draw | Provides card draw.
- 1 Lórien Revealed | draw | Provides card draw.
- 1 Night's Whisper | draw | Provides efficient card draw.
- 1 Palantír of Orthanc | draw | Provides card draw.
- 1 Plunder the Trollshaws | draw | Provides card draw.
- 1 Thrór's Map | draw | Provides card draw.
- 1 Uncover the Moon-Letters | draw | Provides card draw.
- 1 Confusticate and Bebother | interaction | Provides stack interaction.
- 1 Elrond, Moon-Reader | interaction | Provides interaction while fitting the Elf core.
- 1 Mithril Coat | interaction | Provides protective interaction.
- 1 My Precious // Allure of Power | interaction | Provides flexible interaction.
- 1 Old Fat Spider Can't See Me | interaction | Provides interaction.
- 1 Stern Scolding | interaction | Provides efficient interaction.
- 1 The One Ring | interaction | Provides interaction.
- 1 Bilbo's Deadly Slice | removal | Provides targeted removal.
- 1 Bitter Downfall | removal | Provides targeted removal.
- 1 Crude Bent Blade | removal | Provides removal from a durable artifact slot.
- 1 Enchanted River's Grasp | removal | Provides targeted removal.
- 1 Giant's Boulder | removal | Provides removal.
- 1 Merciless Executioner | removal | Provides creature-based removal.
- 1 Orcish Bowmasters | removal | Provides efficient creature-based removal.
- 1 The Black Arrow | removal | Provides targeted removal.
- 1 Troll Negotiations | removal | Provides targeted removal.
- 1 Gnashing of Teeth | wipe | Provides a battlefield reset.
- 1 Languish | wipe | Provides a battlefield reset.
- 1 Raise the Palisade | wipe | Provides a battlefield reset.
- 1 Arwen, Weaver of Hope | synergy | An Elf legend that reinforces the deck's Elf-centered plan.
- 1 Celeborn the Wise | synergy | An Elf legend that reinforces the deck's Elf-centered plan.
- 1 Galadhrim Guide | synergy | An Elf that reinforces the deck's Elf-centered plan.
- 1 Galion, Elvenking's Butler | synergy | An Elf court card that reinforces the Thranduil theme.
- 1 Mirkwood Meditator | synergy | An Elf that reinforces the deck's Elf-centered plan.
- 1 Mirkwood Pathmaker | synergy | An Elf that reinforces the deck's Elf-centered plan.
- 1 Supper for Spiders | synergy | Provides the shortlist's dedicated synergy piece.
- 1 Attercop | threat | A creature threat that adds pressure to the board.
- 1 Boughside Wanderers | threat | An Elf creature that adds pressure to the board.
- 1 Cantankerous Keepers | threat | An Elf creature that adds pressure to the board.
- 1 Guardian of the Halls | threat | An Elf creature that adds pressure to the board.
- 1 Haunt of the Dead Marshes | threat | A creature threat that adds pressure to the board.
- 1 Large Bear | threat | A creature threat that adds board pressure.
- 1 Lothlórien Lookout | threat | An Elf creature that adds pressure to the board.
- 1 Mirkwood Elk | threat | A creature threat that adds board presence.
- 1 Nimrodel Watcher | threat | An Elf creature that adds pressure to the board.
- 1 Rhovanion Rampager | threat | A creature threat that adds pressure to the board.
- 1 Willow-Wind | threat | A creature threat that adds board pressure.
- 1 Troll of Khazad-dûm | wincon | A marked finisher that closes games.
- 1 Witch-king of Angmar | wincon | A marked finisher that closes games.
- 1 Witch-king, Bringer of Ruin | wincon | A marked finisher that closes games.
- 1 Hobbit Hole | land | the mana pass added it to bring the mana base inside the power level
- 1 Pelargir Survivor | ramp | the mana pass added it to bring the mana base inside the power level
- 1 Woodland Weavemaster | ramp | the mana pass added it to bring the mana base inside the power level

</details>

### 20. the Hobbit family, a delegated commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 77 names.

Commander: Smaug the Magnificent.

Grade: baseline, score 0.04, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $469.51 to buy, $469.51 the whole deck.

**Summary:** This deck builds toward Smaug the Magnificent with a broad red mana base, plentiful acceleration, and a mix of creature pressure and equipment utility. Cavern-Hoard Dragon, Desert Were-Worm, and Smaug, the Great Calamity // Spew Flame provide the closing power, while the removal suite and three sweepers clear the way. It gives up some dragon density for a deeper package of artifacts, supporting creatures, and board-control tools drawn from the Hobbit shortlist.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: It functions as a generic mono-red Hobbit goodstuff pile — ramp, equipment, small goblin/dwarf bodies — but the three sweepers work against its own wide board of small creatures and no single engine ties the pieces together.
- PLAN theme_fit=no: The person asked for a dragons deck and the list contains only a handful of dragons, the bulk being dwarves, goblins, and equipment, so it answers a different request even if the set and format are right.
- PLAN useful_as_built=yes: Mono-red with roughly 36 lands plus abundant cheap ramp and a low curve, it casts its spells reliably and has several real threats and evasion routes to close a game.
- PLAN summary_honest=yes: The summary names real cards in the list, correctly counts the sweepers, and openly admits the deck sacrifices dragon density, hiding nothing material.
- [INFO] `curve_summary`: average mana value 3.24 over 63 nonland cards
- [WARN] `profile_off_band`: the draw count is 6, and bracket 3 wants 8 to 14

<details><summary>The deck list</summary>

- 1 Andúril, Flame of the West | other | Provides a utility equipment slot for the deck.
- 1 Andúril, Narsil Reforged | other | Provides a utility equipment slot for the deck.
- 1 Arcane Signet | ramp | Provides reliable mana acceleration.
- 1 Bag End Banquet | ramp | Adds to the deck's mana development.
- 1 Balin, Loremaster | draw | Adds a dedicated card-flow piece.
- 1 Bilbo's Ring | interaction | Supplies a defensive utility option.
- 1 Bothersome Noisemaker | threat | Fills out the deck's creature pressure.
- 1 Burn, Burn, Tree and Fern | ramp | Contributes to mana acceleration.
- 1 Call Forth the Tempest | wipe | Provides a full-board reset option.
- 1 Cavern-Hoard Dragon | wincon | Serves as a dragon finisher.
- 1 Desert Were-Worm | wincon | Serves as a large finisher.
- 1 Desolation of Smaug | wipe | Provides another full-board reset.
- 1 Dragon's Desire | ramp | Helps accelerate into the deck's larger cards.
- 1 Dwarven Mattock | interaction | Adds a utility interaction piece.
- 1 Dwarven Mauler | threat | Adds a creature threat to the board.
- 1 Dwarven Warriors | threat | Adds another creature threat.
- 1 Dáin Ironfoot | threat | Provides a high-impact creature threat.
- 1 Fire of Orthanc | removal | Provides targeted removal.
- 1 Gandalf, Goblins' Bane // Flameshape | threat | Adds a creature threat with extra utility.
- 1 Gandalf, Spark Starter | removal | Provides a removal option.
- 1 Getaway Barrel | other | Fills a flexible utility slot.
- 1 Giant's Boulder | removal | Provides targeted removal.
- 1 Glamdring | other | Provides a utility equipment slot.
- 1 Glóin the Mighty // Easy Pickings | wipe | Adds a third board-reset effect.
- 1 Goblin Cratermaker | removal | Provides creature-based removal.
- 1 Goblin Fireleaper | removal | Adds another removal tool.
- 1 Goblin-town Flunkies | threat | Adds creature pressure.
- 1 Gundabad Opportunist | threat | Adds another creature threat.
- 1 Guttersnipe | threat | Adds a creature threat alongside the deck's spells.
- 1 Hobbit Hole | land | Provides a nonbasic land slot.
- 1 Improvised Club | removal | Provides efficient targeted removal.
- 1 Inferno Titan | removal | Provides a large removal threat.
- 1 Iron Hills Stalwart | threat | Adds creature pressure.
- 1 Key to the Side-Door | draw | Adds to the deck's card flow.
- 1 Last Light of Durin's Day | synergy | Provides the deck's dedicated synergy piece.
- 1 Long-Lost Lances | other | Provides a utility equipment slot.
- 1 Misty Mountains Raider | threat | Adds another creature threat.
- 1 Mithril Coat | interaction | Provides a protective interaction option.
- 31 Mountain | land | Provides the primary mana base for the deck.
- 1 Mox Amber | ramp | Provides early mana acceleration.
- 1 Old Thrush | other | Fills a utility creature slot.
- 1 Oliphaunt | threat | Adds a substantial creature threat.
- 1 Olog-hai Crusher | threat | Adds a substantial creature threat.
- 1 Orcish Siegemaster | other | Fills a creature utility slot.
- 1 Orcrist, Goblin-cleaver | ramp | Contributes to mana acceleration.
- 1 Palantír of Orthanc | draw | Provides sustained card flow.
- 1 Pinecone Strike | removal | Provides targeted removal.
- 1 Ragged Short Spear | draw | Adds another card-flow source.
- 1 Rogue's Passage | land | Provides a nonbasic land utility slot.
- 1 Smaug's Fury | other | Provides a flexible utility spell.
- 1 Smaug, the Great Calamity // Spew Flame | wincon | Serves as a dragon finisher.
- 1 Smite the Deathless | removal | Provides targeted removal.
- 1 Snowslope Hunter | other | Fills a creature utility slot.
- 1 Sting, Bilbo's Sword | other | Provides a utility equipment slot.
- 1 Stone-Giant of High Pass | removal | Provides another removal body.
- 1 The Black Arrow | removal | Provides targeted removal.
- 1 The Lonely Mountain | land | Provides a red-producing nonbasic land slot.
- 1 The Misty Mountains Cold | ramp | Adds to the deck's mana acceleration.
- 1 The One Ring | interaction | Provides a premium interaction option.
- 1 The Reaver Cleaver | ramp | Helps accelerate the deck's mana development.
- 1 Thorin, Mountain-king | removal | Provides creature-based removal.
- 1 Thrór's Map | draw | Adds to the deck's card flow.
- 1 Tidings of War | other | Provides a flexible utility spell.
- 1 Treasure Vault | land | Provides a nonbasic land utility slot.
- 1 Troop of Ponies | ramp | Contributes to mana acceleration.
- 1 Wayfarer's Bauble | ramp | Provides early mana development.
- 1 Well-Worn Spatula | other | Provides a utility equipment slot.
- 1 Óin the Brave | draw | Adds another dedicated card-flow piece.
- 1 Dragon-Cursed Halls | land | Provides a nonbasic land slot.

</details>

### 21. the Hobbit family, mana from outside

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 156 names.

Commander: Smaug the Impenetrable.

Grade: baseline, score 0.07, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $914.87 to buy, $914.87 the whole deck.

**Summary:** This Smaug deck builds its mana with artifacts, themed ramp, and a stable red-black land base before taking over combat with Dragons, Goblins, Orcs, Trolls, and legendary threats. Its strongest closing pressure comes from the marked finishers alongside equipped creatures, while removal, protective artifacts, and board wipes keep opposing boards from getting too far ahead. It gives up some sustained card flow and relies on its creature board and individual high-impact threats to finish games.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: There is a loose Rakdos creatures-plus-equipment midrange plan (roughly a dozen Equipment alongside cheap goblins/orcs and removal), but the summary's framing as Dragons, Goblins, Orcs, Trolls and legends describes a tribal soup rather than one plan, and the deck has little that ties those groups together or that specifically feeds Smaug.
- PLAN theme_fit=partly: It is a Commander deck led by Smaug the Impenetrable at a reasonable bracket-3 power, but the request limited non-land, non-rock cards to the Hobbit set and the list includes many LTR-era and generic cards (Orcish Bowmasters, The One Ring, Andúril, Glamdring, Languish, Night's Whisper, Guttersnipe, Goblin Cratermaker).
- PLAN useful_as_built=yes: 33 lands plus eight ramp pieces, a low curve, plenty of removal and multiple combat-based and equipment-based win routes make this immediately playable, with only a stray off-color fetch (Misty Rainforest) as dead weight.
- PLAN summary_honest=partly: The equipment, removal, three wipes and thin card flow are described accurately, but it oversells a "Dragon" dimension (essentially just the Smaug cards) and a "stable red-black land base" while quietly including off-color fetchlands, and it never flags that many nonland, non-rock cards come from outside the requested set.
- [INFO] `curve_summary`: average mana value 2.77 over 62 nonland cards
- [WARN] `profile_off_band`: the draw count is 6, and bracket 3 wants 8 to 14
- [INFO] `mana_pass`: the builder moved 4 cards of the mana base to bring the deck inside its power level
- [WARN] `outside_requested_set`: the sets you named do not hold 17 cards: Ancient Tomb, Arid Mesa, Blood Crypt, Bloodstained Mire, Bojuka Bog, Command Tower, Dragonskull Summit, Evolving Wilds, Exotic Orchard, Fabled Passage, Marsh Flats, Misty Rainforest, Path of Ancestry, Polluted Delta, Scalding Tarn, Terramorphic Expanse, Urborg, Tomb of Yawgmoth

<details><summary>The deck list</summary>

- 1 Ancient Tomb | land | An untapped land slot that helps accelerate the deck's mana.
- 1 Arid Mesa | land | A flexible land slot for the mana base.
- 1 Blood Crypt | land | A red-black land slot for the mana base.
- 1 Bloodstained Mire | land | A flexible land slot for the mana base.
- 1 Bojuka Bog | land | A utility land slot with a black mana option.
- 1 Command Tower | land | A reliable colored land slot.
- 1 Dragonskull Summit | land | A red-black land slot for the mana base.
- 1 Evolving Wilds | land | A flexible land slot for the mana base.
- 1 Exotic Orchard | land | A flexible colored land slot.
- 1 Fabled Passage | land | A flexible land slot for the mana base.
- 1 Marsh Flats | land | A flexible land slot for the mana base.
- 1 Misty Rainforest | land | A flexible land slot for the mana base.
- 1 Minas Morgul, Dark Fortress | land | A themed utility land slot.
- 1 Mount Doom | land | A themed utility land slot.
- 10 Mountain | land | A basic red mana source.
- 1 Path of Ancestry | land | A themed colored land slot.
- 1 Polluted Delta | land | A flexible land slot for the mana base.
- 1 Scalding Tarn | land | A flexible land slot for the mana base.
- 8 Swamp | land | A basic black mana source.
- 1 Terramorphic Expanse | land | A flexible land slot for the mana base.
- 1 Urborg, Tomb of Yawgmoth | land | A black-focused utility land slot.
- 1 Arcane Signet | ramp | A compact mana source that supports the commander and colored spells.
- 1 Dragon's Desire | ramp | A themed ramp spell.
- 1 Mox Amber | ramp | A low-cost legendary-focused mana source.
- 1 Orcrist, Goblin-cleaver | ramp | A themed Equipment ramp card.
- 1 Smaug the Magnificent | ramp | A second Smaug that advances the deck's mana plan.
- 1 Smaug, Wicked Worm | ramp | A Dragon ramp card with a finisher marking.
- 1 The Reaver Cleaver | ramp | A themed Equipment ramp card.
- 1 Wayfarer's Bauble | ramp | An early artifact ramp piece.
- 1 Key to the Side-Door | draw | A colorless card-advantage piece.
- 1 Night's Whisper | draw | An efficient black draw spell.
- 1 Palantír of Orthanc | draw | A themed legendary draw artifact.
- 1 Ragged Short Spear | draw | A themed Equipment draw card.
- 1 Rage into the Valley | draw | A red draw spell that fits the deck's theme.
- 1 Thrór's Map | draw | A themed legendary draw artifact.
- 1 Bilbo's Ring | interaction | A themed Equipment that supplies interaction.
- 1 Dwarven Mattock | interaction | A themed Equipment interaction piece.
- 1 Mithril Coat | interaction | A protective Equipment interaction piece.
- 1 My Precious // Allure of Power | interaction | A themed legendary interaction card.
- 1 The One Ring | interaction | A powerful legendary interaction artifact.
- 1 Azog, Moria's Ruin | removal | A themed legendary removal creature.
- 1 Battle-Scarred Goblin | removal | A Goblin removal creature for the board plan.
- 1 Bilbo's Deadly Slice | removal | A direct black removal spell.
- 1 Bitter Downfall | removal | A direct black removal spell.
- 1 Crude Bent Blade | removal | A themed Equipment removal card.
- 1 Fire of Orthanc | removal | A red removal spell.
- 1 Goblin Cratermaker | removal | A Goblin creature that fills a removal slot.
- 1 Improvised Club | removal | A compact red removal spell.
- 1 Orcish Bowmasters | removal | An Orc removal creature with a low mana cost.
- 1 Desolation of Smaug | wipe | A Smaug-themed board wipe.
- 1 Gnashing of Teeth | wipe | A thematic wipe for resetting the board.
- 1 Languish | wipe | A reliable black board wipe.
- 1 Supper for Spiders | synergy | The available dedicated synergy card from the shortlist.
- 1 Along the Crooked Way | other | A themed noncreature card that supports the deck's overall plan.
- 1 Andúril, Flame of the West | other | A legendary Equipment for the creature-heavy plan.
- 1 Andúril, Narsil Reforged | other | A second legendary Equipment for the creature-heavy plan.
- 1 Bothersome Noisemaker | other | A Goblin creature that reinforces the Middle-earth creature suite.
- 1 Desolation Prowler | other | A creature slot that adds board presence.
- 1 Down, Down to Goblin-town | other | A themed Saga with a finisher marking.
- 1 Dreaded Bat-Cloud | other | A creature slot that adds board presence.
- 1 Dwarven Mauler | other | A creature that broadens the deck's combat board.
- 1 Dwarven Warriors | other | A themed creature for the combat plan.
- 1 Fearsome Goblin Pair | other | A Goblin creature that strengthens the creature suite.
- 1 Gandalf, Goblins' Bane // Flameshape | other | A thematic legendary creature-adventure card.
- 1 Gathering of Darkness | other | A thematic black noncreature card.
- 1 Getaway Barrel | other | A utility artifact for the deck's board plan.
- 1 Glamdring | other | A legendary Equipment that supports the creature plan.
- 1 Gollum the Abandoned | other | A thematic legendary creature for board presence.
- 1 Gollum, Silent Slinker // Meager Meal | other | A thematic legendary creature-adventure card.
- 1 Goblin Plate Mail | other | An Equipment that supports the Goblin-heavy creature suite.
- 1 Goblin-town Flunkies | other | A Goblin creature that supports the combat board.
- 1 Great Goblin, Foul-Hearted | other | A legendary Goblin that adds a meaningful creature presence.
- 1 Gundabad Opportunist | other | A Goblin creature that adds board presence.
- 1 Guttersnipe | other | A Goblin creature that complements the spell suite.
- 1 Haunt of the Dead Marshes | other | A creature slot that adds board presence.
- 1 Nighthowl Pursuer | other | A creature slot that adds board presence.
- 1 Olog-hai Crusher | other | A large themed creature for the combat plan.
- 1 Orcish Siegemaster | other | An Orc creature that strengthens the creature suite.
- 1 Ravening Warg | other | A creature slot that adds board presence.
- 1 Sauron, the Lidless Eye | other | A thematic legendary creature for the deck's top end.
- 1 Giant's Boulder | ramp | the mana pass added it to bring the mana base inside the power level
- 1 Bolg's Company | ramp | the mana pass added it to bring the mana base inside the power level
- 1 Dori, Bearer of Friends | ramp | the mana pass added it to bring the mana base inside the power level

</details>

### 22. a set family and a card from outside it

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 175 names.

Commander: Thranduil, the Elvenking.

Grade: baseline, score 0.04, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $406.09 to buy, $406.09 the whole deck.

**Summary:** Thranduil leads a creature-centered Middle-earth deck built around Elf support, steady mana development, and a broad mix of draw, removal, and protection. Establish a board, use the interaction package to preserve tempo, then turn the corner with the deck's large marked finishers and Witch-kings. The deck favors its themed creature plan and resilient midgame over the raw speed of a more artifact-heavy build.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: There is a loose creature-midrange spine with a few Elf mana dorks and a fat pile of removal, but the Elf support is thin and the finishers (Witch-kings, Trolls, big bears) don't connect to any of the Elf or "marked" synergies the summary implies, so the list reads as goodstuff rather than one plan.
- PLAN theme_fit=partly: The requested commander, bracket, and the kept Sol Ring are all honored, but the request asked for a Hobbit-set deck and a large share of the list (Orcish Bowmasters, Delighted Halfling, Lórien Revealed, Palantír of Orthanc, Troll of Khazad-dûm) plus a non-Middle-earth Mox Amber comes from elsewhere.
- PLAN useful_as_built=partly: The curve, draw, and threat count are fine and it will function, but 36 basics with no dual or fetch land in a three-color deck full of blue instant-speed interaction will strand cards more often than it should.
- PLAN summary_honest=partly: Draw, removal, ramp, and the Witch-king finishers are all genuinely present, but "Elf support" overstates a handful of dorks with no payoff web and "large marked finishers" names a synergy the list barely carries.
- [INFO] `curve_summary`: average mana value 3.10 over 63 nonland cards
- [WARN] `profile_off_band`: the worst color's share of the sources it needs is 0.75, and bracket 3 wants 0.85 or more (sources of requirement: U 14.5 of 19, B 17.25 of 23, G 17 of 22)
- [INFO] `mana_pass`: the builder moved 5 cards of the mana base to bring the deck inside its power level
- [WARN] `outside_requested_set`: the sets you named do not hold 1 card: Sol Ring

<details><summary>The deck list</summary>

- 14 Forest | land | Untapped green mana for the deck's Elf-heavy core.
- 10 Island | land | Untapped blue mana for draw and interaction.
- 12 Swamp | land | Untapped black mana for removal and finishers.
- 1 Sol Ring | ramp | Kept as requested and provides efficient mana development.
- 1 Arcane Signet | ramp | Reliable color fixing and mana acceleration.
- 1 Delighted Halfling | ramp | Early mana development on a themed creature.
- 1 Elven Chorus | ramp | Elf-themed mana acceleration.
- 1 Elvish Archdruid | ramp | Elf-based mana production that supports the creature plan.
- 1 Elvish Mystic | ramp | Low-cost early mana acceleration.
- 1 Thranduil the Strategist | ramp | Themed mana development for Thranduil's deck.
- 1 Wayfarer's Bauble | ramp | Early artifact-based mana development.
- 1 Wood Elves | ramp | Creature-based mana development that fits the Elf core.
- 1 Beorn the Fierce | draw | Provides a draw option on a substantial themed creature.
- 1 Elvish Visionary | draw | Efficient Elf-themed card advantage.
- 1 Fateful Discovery | draw | Dedicated card advantage for the deck.
- 1 Hithlain Knots | draw | Instant-speed card advantage.
- 1 Ithilien Kingfisher | draw | Creature-based card advantage.
- 1 Lórien Revealed | draw | Flexible card advantage from a blue spell.
- 1 Night's Whisper | draw | Low-cost black card advantage.
- 1 Palantír of Orthanc | draw | Repeatable artifact-based card advantage.
- 1 Plunder the Trollshaws | draw | Instant-speed card advantage.
- 1 Uncover the Moon-Letters | draw | Additional dedicated card advantage.
- 1 Bilbo's Ring | interaction | A protective interaction piece.
- 1 Confusticate and Bebother | interaction | Flexible instant-speed interaction.
- 1 Dwarven Mattock | interaction | Artifact-based interaction.
- 1 Elrond, Moon-Reader | interaction | Legendary interaction that supports the Middle-earth theme.
- 1 Mithril Coat | interaction | Protects an important creature or commander.
- 1 My Precious // Allure of Power | interaction | Versatile interaction from an equipment-adventure card.
- 1 Old Fat Spider Can't See Me | interaction | Themed interaction on a Saga.
- 1 Stern Scolding | interaction | Efficient blue instant-speed interaction.
- 1 Bilbo's Deadly Slice | removal | Dedicated removal spell.
- 1 Bitter Downfall | removal | Dedicated black removal.
- 1 Enchanted River's Grasp | removal | Blue removal that answers a problematic permanent.
- 1 Giant's Boulder | removal | Artifact-based removal.
- 1 Merciless Executioner | removal | Creature-based removal for the board.
- 1 Orcish Bowmasters | removal | Efficient creature-based removal.
- 1 Quarrel | removal | Instant-speed removal.
- 1 The Black Arrow | removal | Themed equipment-based removal.
- 1 Uneasy Partings | removal | Additional instant-speed removal.
- 1 Gnashing of Teeth | wipe | A board-clearing option.
- 1 Languish | wipe | Efficient black board wipe.
- 1 Raise the Palisade | wipe | A themed board-reset option.
- 1 Boughside Wanderers | synergy | Elf-themed support for the creature core.
- 1 Elven Raft-Steerer | synergy | An Elf that reinforces the deck's tribal theme.
- 1 Galadhrim Guide | synergy | Elf-themed support for the board plan.
- 1 Galion, Elvenking's Butler | synergy | A Thranduil-themed Elf support card.
- 1 Lothlórien Lookout | synergy | Elf-themed support for the creature base.
- 1 Mirkwood Meditator | synergy | An Elf support piece for the themed plan.
- 1 Mirkwood Pathmaker | synergy | An Elf that deepens the Mirkwood creature theme.
- 1 Attercop | threat | A sizable themed creature that adds board pressure.
- 1 Desolation Prowler | threat | Creature pressure that supports the board-focused plan.
- 1 Dreaded Bat-Cloud | threat | A threatening creature for combat pressure.
- 1 Duskwatch Hunter | threat | Creature-based pressure from the darker side of the theme.
- 1 Gigantic Big Bear | threat | A large creature that advances the combat plan.
- 1 Gollum, Riddle Master | threat | A marked finisher that also supplies card advantage.
- 1 Haunt of the Dead Marshes | threat | A threatening Nightmare Elf for board presence.
- 1 Mirkwood Elk | threat | A solid green creature for combat pressure.
- 1 Nasty Little Rabbit | threat | A themed creature that contributes to the board.
- 1 Old Fat Spider | threat | A marked finisher that creates a serious battlefield threat.
- 1 Ravening Warg | threat | Creature pressure that complements the deck's darker threats.
- 1 Troll of Khazad-dûm | wincon | A marked finisher and dedicated closing threat.
- 1 Witch-king of Angmar | wincon | A marked finisher that serves as a primary closer.
- 1 Witch-king, Bringer of Ruin | wincon | A marked finisher for closing games.
- 1 Mox Amber | ramp | the mana pass added it to bring the mana base inside the power level
- 1 Pelargir Survivor | ramp | the mana pass added it to bring the mana base inside the power level
- 1 Woodland Weavemaster | ramp | the mana pass added it to bring the mana base inside the power level

</details>

### 23. two set families at once

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 203 names.

Commander: Kíli the Resourceful.

Grade: typical, score 0.44, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: yes, for deck_size, profile_off_band, profile_off_band, profile_off_band, profile_off_band, profile_off_band, profile_off_band, profile_off_band, profile_off_band. Block findings: 0.

Cost: $281.35 to buy, $281.35 the whole deck.

**Summary:** This mono-white Hobbit and Bloomburrow deck develops a durable board of Dwarves, small creatures, artifact tokens, and Equipment. It uses efficient mana rocks and land ramp to establish steady resources, then keeps cards flowing through creature, token, and artifact value engines. Flexible exile effects, protective spells, and sweepers let it control hostile boards while its upgraded creature force closes through combat and recurring permanent value.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

Bracket 3 wants 2 or more finishers, the cards that win the game, and the deck holds 1.

Summary rules claims (F-26): the format play

- JUDGE [false]: "Bracket 3 wants 2 or more finishers, the cards that win the game, and the deck holds 1.". This asserts a deck-construction requirement of a format tier. The official Commander bracket system defines brackets by game changers, tutors, mass land denial, extra turns, and combo speed; it sets no minimum number of 'finishers'. 'Finisher count' is an internal heuristic, not a format rule.
- PLAN plan_coherent=partly: The cards broadly point at a mono-white go-wide board of small creatures, artifact tokens and Equipment with card-draw payoffs, but the pieces are spread thin across a half-hearted Dwarf subtheme, Equipment voltron pieces, token payoffs and three symmetrical sweepers that fight its own wide board.
- PLAN theme_fit=no: The one explicit constraint was that cards come from the Hobbit and Bloomburrow sets, yet a large share of the deck (Sol Ring, Swords to Plowshares, Skullclamp, Sun Titan, Jazal Goldmane, Karn, Academy Manufactor, Blade Splicer, Luminous Broodmoth, Andúril, Landroval and more) is from outside them.
- PLAN useful_as_built=yes: Thirty-eight lands plus cheap rocks and land ramp support a low mono-white curve, and the deck has ample removal, draw and several combat finishers (Jazal, Warren Warleader, Andúril) to actually close games.
- PLAN summary_honest=partly: The described ramp, draw engines, exile removal, protection and sweepers are all genuinely in the list, but calling it a "Hobbit and Bloomburrow deck" hides that roughly half the cards come from neither set, and the Dwarf emphasis is overstated for five or so Dwarves.
- [INFO] `curve_summary`: average mana value 3.03 over 63 nonland cards
- [WARN] `profile_off_band`:  is 1, and bracket 3 wants 2 or more (Angel of the Ruins)

<details><summary>The deck list</summary>

- 25 Plains | land | Reliable untapped white sources.
- 1 Castle Ardenvale | land | White source with late-game token utility.
- 1 Command Tower | land | Untapped commander color source.
- 1 Elven Passage | land | Additional colored land utility.
- 1 Evolving Wilds | land | Finds a basic Plains while fixing mana.
- 1 Exotic Orchard | land | Flexible colored source in multiplayer.
- 1 Fabled Passage | land | Fetches a Plains and improves land access.
- 1 Hobbit Hole | land | Thematic utility land and colored source.
- 1 Minas Tirith | land | White source with card-flow utility.
- 1 Path of Ancestry | land | Supports the creature-heavy tribal plan.
- 1 Thriving Heath | land | Additional white fixing.
- 1 Uncharted Haven | land | Flexible colored mana source.
- 1 Arcane Signet | ramp | Efficient color fixing and acceleration.
- 1 Bag End Banquet | ramp | Thematic mana acceleration.
- 1 Burnished Hart | ramp | Turns spare mana into additional basic lands.
- 1 Fellwar Stone | ramp | Low-cost multiplayer mana rock.
- 1 Mind Stone | ramp | Early mana that can become a card later.
- 1 Mox Amber | ramp | Fast legendary-focused acceleration.
- 1 Ornithopter of Paradise | ramp | Creature-based mana fixing.
- 1 Patchwork Banner | ramp | Mana rock that also supports the creature type plan.
- 1 Sol Ring | ramp | Efficient artifact acceleration.
- 1 Wayfarer's Bauble | ramp | Early ramp that permanently improves land count.
- 1 Belladonna Took | draw | Token production converts into steady card flow.
- 1 Caretaker's Talent | draw | Rewards token creation with cards and board growth.
- 1 Circuit Mender | draw | Artifact body that replaces itself.
- 1 Cut a Deal | draw | Refills a depleted hand.
- 1 Dawn of a New Age | draw | Long-game card advantage engine.
- 1 Errand-Rider of Gondor | draw | Creature-based card advantage.
- 1 Fountainport Bell | draw | Repeatable artifact card selection.
- 1 Idol of Oblivion | draw | Efficient draw alongside token makers.
- 1 Inspiring Overseer | draw | Flying body that replaces itself.
- 1 Mentor of the Meek | draw | Turns the many smaller creatures into cards.
- 1 Skullclamp | draw | Converts expendable creature tokens into a strong refill.
- 1 Banishing Light | removal | Versatile answer to problematic permanents.
- 1 Bumbleflower's Sharepot | removal | Thematic permanent-based answer.
- 1 Fiend Hunter | removal | Creature-based exile effect.
- 1 Generous Gift | removal | Answers any permanent at instant speed.
- 1 Hanged Executioner | removal | Flexible creature threat and exile effect.
- 1 Loran of the Third Path | removal | Removes artifacts or enchantments while offering card utility.
- 1 Nettle Guard | removal | Creature-based removal for the board plan.
- 1 Skyclave Apparition | removal | Efficient exile attached to a creature.
- 1 Swords to Plowshares | removal | Premium low-cost creature answer.
- 1 Dusk // Dawn | wipe | Resets larger boards and can recover smaller creatures.
- 1 Martial Coup | wipe | Board reset that leaves behind an army.
- 1 Promise of Loyalty | wipe | Answers creature-heavy boards while preserving a key creature.
- 1 Baird, Steward of Argive | interaction | Taxes opposing attacks and protects the developing board.
- 1 Bilbo's Gambit | interaction | Flexible instant-speed disruption.
- 1 Dawn's Truce | interaction | Protects the board from opposing interaction.
- 1 Galadriel's Dismissal | interaction | Defensive tempo interaction that protects a winning board.
- 1 Mithril Coat | interaction | Protects a crucial creature or commander.
- 1 Parting Gust | interaction | Versatile instant-speed disruption.
- 1 Reprieve | interaction | Delays a key spell while replacing itself.
- 1 Swiftfoot Boots | interaction | Protects important creatures and grants haste.
- 1 Academy Manufactor | synergy | Multiplies the value of Food and other artifact tokens.
- 1 Blade Splicer | synergy | Produces an artifact creature body for the artifact plan.
- 1 Carrot Cake | synergy | Creates Food and bodies for token-based value.
- 1 Dáin, Lord of the Iron Hills | synergy | Rewards the Dwarf-heavy creature suite.
- 1 Iron Hills Blacksmith | synergy | Supports Equipment and artifact creatures.
- 1 Ori, Keeper of Songs | synergy | Builds value around Dwarves and artifacts.
- 1 Tangle Tumbler | synergy | Artifact creature that supports the deck's token and artifact themes.
- 1 Fíli the Pathfinder | threat | Thematic legendary attacker that pressures opponents.
- 1 Karn, the Great Creator | threat | A resilient value threat that disrupts opposing artifacts.
- 1 Jazal Goldmane | threat | Turns a wide creature board into a lethal combat threat.
- 1 Landroval, Horizon Witness | threat | Evasive legendary creature that adds pressure and value.
- 1 Luminous Broodmoth | threat | Makes creature-based board development resilient.
- 1 Psychosis Crawler | threat | Converts the deck's card draw into table-wide damage.
- 1 Serra Redeemer | threat | Grows the board with recurring creature value.
- 1 Steelburr Champion | threat | Efficient creature pressure for the combat plan.
- 1 Sun Titan | threat | Recurring permanents generates overwhelming long-game value.
- 1 Sunscorch Regent | threat | Evasive scaling threat that gains life while attacking.
- 1 Warren Warleader | threat | Creates a wide attacking force and rewards combat.
- 1 Angel of the Ruins | wincon | Large evasive artifact threat that answers key permanents.
- 1 Realm-Cloaked Giant // Cast Off | wincon | Can clear the battlefield before ending games with a huge creature.
- 1 Andúril, Flame of the West | other | Thematic Equipment that upgrades the creature offense.
- 1 Glamdring | other | Thematic legendary Equipment for the combat plan.

</details>

### 24. a 60-card deck from one set

Format: Modern. Theme: aggro. Pool: any_card. Shortlist: 96 names.

Grade: bad, score 0.31, model 20260914T154223Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $27.64 to buy, $27.64 the whole deck.

**Summary:** This mono-red Bloomburrow aggro deck aims to establish pressure with its creature threats, reinforce combat with its synergy cards, and use removal to keep opposing resistance out of the way. It wins by sustaining an aggressive attack while draw cards help refill its options. The deck gives up broader utility and a deep long-game package in favor of a focused, creature-forward red plan.

The quality model grades this deck below the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every card supports one curve-out aggro plan — cheap red creatures, cheap removal to clear blockers, and pump/cantrip effects — with only the two five-mana Season of the Bold copies sitting slightly outside the aggressive line.
- PLAN theme_fit=yes: It is a mono-red Modern-legal aggro list built entirely from Bloomburrow cards, matching every constraint the requester named, at a casual power level.
- PLAN useful_as_built=yes: 24 Mountains is a couple more land than a lean aggro deck wants, but the mana is flawless, the curve is heavily one-to-three drops, and there are plenty of creature threats to close a game.
- PLAN summary_honest=yes: The summary claims creature pressure, combat tricks, removal, and a little card draw, all of which the list actually contains, and it openly admits the deck lacks long-game depth.
- [INFO] `curve_summary`: average mana value 3.17 over 36 nonland cards

<details><summary>The deck list</summary>

- 24 Mountain | land | Provides the red mana base for a mono-red deck.
- 4 Might of the Meek | draw | Supplies card flow while fitting the aggressive creature plan.
- 2 Sazacap's Brew | draw | Adds additional card flow to keep threats coming.
- 2 Season of the Bold | ramp | Provides the requested mana acceleration without leaving mono-red.
- 4 Abrade | removal | Gives the deck efficient removal support.
- 2 Agate Assault | removal | Completes the removal suite for clearing resistance.
- 4 Emberheart Challenger | synergy | A central red synergy creature for the aggressive plan.
- 4 Hearthborn Battler | synergy | Adds another full set of synergy creatures to support combat pressure.
- 4 Frilled Sparkshooter | threat | A full set of red threats for maintaining pressure.
- 4 Reptilian Recruiter | threat | Provides a reliable threat package for the creature assault.
- 4 Teapot Slinger | threat | Adds more threats to make the deck consistently aggressive.
- 2 Dragonhawk, Fate's Tempest | threat | Rounds out the threat suite with a higher-impact finisher.

</details>

### 25. use no card of an owned precon

Format: Commander. Theme: Avengers superheroes. Pool: owned_first. Shortlist: 154 names.

Commander: Syr Konrad, the Grim.

Grade: typical, score 0.36, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $204.09 the whole deck.

**Summary:** This mono-black Syr Konrad deck develops its mana, builds resources through a deep draw package, and uses targeted removal plus board wipes to keep the table under pressure. Its creature-and-attrition shell gives the commander a steady supporting board, while Archfiend of Ifnir, Vraska, Betrayal's Sting, and Witch-king of Angmar provide decisive late-game paths. It gives up color flexibility for a highly consistent black mana base and a direct, grind-oriented game plan.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: The list is a unified mono-black attrition engine: a -1/-1 counter/death subtheme (Blowfly Infestation, Contagion Clasp, Midnight Banshee, Grim Poppet, Soul Snuffers, Black Sun's Zenith, Archfiend of Ifnir) plus nine removal spells and wipes all feed Syr Konrad's death-trigger drain, backed by draw and commander protection.
- PLAN theme_fit=no: The request was an Avengers superheroes deck, and the builder delivered a Lord of the Rings/Avatar/Final Fantasy-flavored Syr Konrad grind pile with no Marvel or superhero cards or flavor connection whatsoever.
- PLAN useful_as_built=yes: 36 Swamps in mono-black with seven ramp pieces, a low curve, and a clear drain/attrition win path means the deck casts its spells and can close games as it stands.
- PLAN summary_honest=partly: placeholder
- [WARN] `not_owned`: Syr Konrad, the Grim: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.06 over 63 nonland cards

<details><summary>The deck list</summary>

- 36 Swamp | land | Provides reliable black mana for Syr Konrad and the deck's black-heavy spells.
- 1 Arcane Signet | ramp | Provides dependable mana acceleration.
- 1 Sol Ring | ramp | Provides early mana acceleration.
- 1 Springleaf Drum | ramp | Provides low-cost mana acceleration.
- 1 Wayfarer's Bauble | ramp | Provides early mana development.
- 1 Thought Vessel | ramp | Provides mana acceleration while supporting a full hand.
- 1 Commander's Sphere | ramp | Provides flexible mana acceleration.
- 1 Chromatic Lantern | ramp | Provides mana acceleration and supports the mana base.
- 1 Sword of the Animist | ramp | Provides repeatable mana development.
- 1 Deadly Dispute | ramp | Contributes mana development while advancing resources.
- 1 Coin of Mastery | ramp | Fills out the deck's mana acceleration package.
- 1 Call of the Ring | draw | Provides a steady source of cards.
- 1 Buster Sword | draw | Supplies card advantage from an equipment slot.
- 1 Hoarder's Greed | draw | Provides a dedicated card-advantage spell.
- 1 Idol of Oblivion | draw | Provides a compact source of cards.
- 1 Instant Ramen | draw | Adds inexpensive card advantage.
- 1 Lembas | draw | Provides card advantage from an artifact slot.
- 1 Mask of Memory | draw | Provides repeatable card advantage.
- 1 Mirror of Galadriel | draw | Provides card advantage in a durable permanent.
- 1 Nasty End | draw | Provides an instant-speed card-advantage option.
- 1 Night's Whisper | draw | Provides efficient card advantage.
- 1 Skullclamp | draw | Provides an efficient source of card advantage.
- 1 Bitter Downfall | removal | Provides targeted removal.
- 1 Bitter Triumph | removal | Provides flexible targeted removal.
- 1 Claim the Precious | removal | Provides targeted removal.
- 1 Dismember | removal | Provides low-cost creature removal.
- 1 Fatal Push | removal | Provides efficient targeted removal.
- 1 Heartless Act | removal | Provides flexible creature removal.
- 1 Infernal Grasp | removal | Provides unconditional targeted removal.
- 1 Nameless Inversion | removal | Provides a compact removal spell.
- 1 Overkill | removal | Provides another dedicated removal effect.
- 1 Black Sun's Zenith | wipe | Provides a scalable board-clearing effect.
- 1 Darkness Descends | wipe | Provides a full-board reset option.
- 1 Extinction Event | wipe | Provides a selective board-clearing effect.
- 1 Champion's Helm | interaction | Helps protect an important legendary permanent.
- 1 Darksteel Plate | interaction | Provides a durable protection piece.
- 1 Lightning Greaves | interaction | Protects a key creature while helping it act immediately.
- 1 Swiftfoot Boots | interaction | Provides another protection option for key creatures.
- 1 Orcish Medicine | interaction | Provides a reactive support spell.
- 1 The Walls of Ba Sing Se | interaction | Provides a defensive interaction permanent.
- 1 Bender's Waterskin | interaction | Occupies a flexible utility slot that supports the deck's key permanents.
- 1 Stone of Erech | interaction | Provides a utility interaction piece.
- 1 Blowfly Infestation | synergy | Supports the deck's incremental creature-focused plan.
- 1 Contagion Clasp | synergy | Provides a dedicated synergy piece for the deck's counter-oriented elements.
- 1 Incremental Blight | synergy | Supports the deck's incremental attrition plan.
- 1 Grave Venerations | synergy | Provides a dedicated permanent that supports the central plan.
- 1 Eventide's Shadow | synergy | Supports the deck's broader Syr Konrad plan.
- 1 Mordor Muster | synergy | Contributes to the deck's creature-focused synergy package.
- 1 Obsessive Pursuit | synergy | Provides another dedicated synergy permanent.
- 1 Carnifex Demon | threat | Serves as a sizable creature threat with broad table impact.
- 1 Dusk Urchins | threat | Provides an early creature threat that fits the deck's creature plan.
- 1 Gorbag of Minas Morgul | threat | Adds a thematic creature body to pressure opponents.
- 1 Grim Poppet | threat | Provides a persistent artifact-creature threat.
- 1 Massacre Girl, Known Killer | threat | Adds a legendary creature threat to the board.
- 1 Midnight Banshee | threat | Provides a high-impact creature threat.
- 1 Orcish Bowmasters | threat | Provides a compact creature threat with immediate board relevance.
- 1 Rat King, Pale Piper | threat | Adds a legendary creature threat to the creature package.
- 1 Ringwraiths | threat | Provides an additional black creature threat.
- 1 Skinrender | threat | Adds a creature threat that contributes to board pressure.
- 1 Soul Snuffers | threat | Provides a creature threat that fits the attrition plan.
- 1 Zodiark, Umbral God | threat | Provides a large, high-impact creature threat.
- 1 Archfiend of Ifnir | wincon | Acts as a designated finishing threat.
- 1 Vraska, Betrayal's Sting | wincon | Provides a dedicated endgame route.
- 1 Witch-king of Angmar | wincon | Acts as a resilient finishing threat.

</details>

