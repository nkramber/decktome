# PR-8 deck gate

Run date: 2026-09-30. Card snapshot: 2026-09-04.

Verdict: PASS. 25 of 25 decks passed every block check, 0 invented names reached the user, and 0 summaries stated a false rule of the game. All three bars are zero tolerance.

## Summary

| Measure | Value |
|---|---|
| Prompts | 25 |
| Decks returned | 25 |
| Decks with no block finding | 25 |
| Invented names that reached the user | 0 |
| Decks that needed the repair turn | 0 |
| Summaries judged (F-26) | 25 |
| Summaries that state a rule of the game | 0 |
| Summaries that state a FALSE rule | 0 |
| Judge errors | 0 |
| Case assertions missed (PR-28b) | 0 |
| Decks the plan judge read (PR-15, information) | 25 |
| Mean plan score, 0 to 1 | 0.69 |
| Plan reasons the judge left empty | 2 |
| Errors | 0 |
| Prompt version | 16 |
| Calls | 81 |
| Cost | $1.6749 |
| Time | 1914 seconds |

## Run

- Suite `decks`, run `pr8-deck-gate-run38`, on 2026-09-29, commit `6cdb2e1`.
- Roles: generate on `gpt-6.1-sol` (openai, effort medium), judge on `claude-sonnet-5-5` (anthropic, effort medium), repair on `gpt-6.1-sol` (openai, effort medium).
- Versions: card snapshot 2026-09-04, generate prompt version 16, plan_rubric prompt version 5, summary_judge prompt version 2, precons `5.3.0+20260923`, quality_model `20260923T202806Z`.
- Calls: 81. Cost: $1.6749. Time: 1914 seconds.

## Findings by code

A block stops the deck. A warning and a note are reports.

| Code | Count |
|---|---|
| `not_owned` | 28 |
| `curve_summary` | 25 |
| `mana_pass` | 12 |
| `profile_off_band` | 9 |
| `outside_requested_set` | 2 |
| `finisher_short` | 2 |
| `bracket_cut` | 1 |

By severity: BLOCK 0. WARN 41. INFO 38. 

## The set filter (PR-17B)

A deck asked for a set holds cards of that set family alone. Three things earn an exception, and each one is marked: a card the reader named (D-381), a mana card the reader allowed from outside (D-382), and a basic land, which no set limit filters (D-378).

| # | Prompt | Sets | In set | Fill | Deck cards outside | Marked |
|---|---|---|---|---|---|---|
| 19 | the Hobbit family, two colours | `hob,hoc` | 170 | 0 | 0 | 0 |
| 20 | the Hobbit family, a delegated commander | `hob,hoc` | 75 | 0 | 0 | 0 |
| 21 | the Hobbit family, mana from outside | `hob,hoc` | 127 | 26 | 14 | 14 |
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

Grade: baseline, score 0.05, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $962.22 to buy, $962.22 the whole deck.

**Summary:** Karlov turns repeated lifegain into a growing attacker and a reserve of creature removal. Develop your mana and lifegain support early, then use draw engines to sustain pressure through creatures and dedicated finishers. The deck can win through combat or its alternate victory plans, while protection and sweepers help it survive longer games. It gives up explosive starts for a fuller mana base and a steady midgame, and its supporting creatures need protection to keep the engine running.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every card gains life or rewards it: Karlov, Aetherflux Reservoir, Archangel of Thune, Elenda, Vito, Righteous Valkyrie, Cosmos Elixir and Test of Endurance all feed one lifegain-into-combat plan.
- PLAN theme_fit=yes: It is a white-black lifegain deck led by Karlov of the Ghost Council with a lifegain payoff suite and power level that fits bracket 3.
- PLAN useful_as_built=partly: The 40-land mana base with many duals plus a dozen mana rocks casts everything reliably, but the ramp packages include several low-impact artifacts and the deck is clunky and over-landed, so it may flood and lack pressure early.
- PLAN summary_honest=partly: The claim that Karlov turns lifegain into a growing attacker is backed by the list, but the claim of dedicated finishers and steady midgame draw overstates it, since several draw and ramp slots are weak artifacts like Orazca Relic and Phial of Galadriel that do little for the plan.
- [INFO] `curve_summary`: average mana value 3.05 over 61 nonland cards
- [INFO] `mana_pass`: the builder moved 2 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 1 Brightclimb Pathway // Grimclimb Pathway | land | Provides an untapped source of whichever color the opening hand needs.
- 1 Caves of Koilos | land | Provides immediate access to both colors.
- 1 City of Brass | land | Fixes either color without slowing early development.
- 1 Command Tower | land | Provides dependable access to either color.
- 1 Concealed Courtyard | land | Supports early deployment of Karlov and inexpensive support cards.
- 1 Forbidden Orchard | land | Provides untapped fixing at the cost of giving opponents small creatures.
- 1 Godless Shrine | land | Provides both colors and a fetchable land with basic land types.
- 1 Isolated Chapel | land | Provides both colors alongside the basic-heavy mana base.
- 1 Mana Confluence | land | Provides flexible untapped mana, with lifegain helping offset its cost.
- 1 Marsh Flats | land | Finds the white or black source needed for the current hand.
- 1 Prismatic Vista | land | Finds an untapped basic of the needed color.
- 1 Restless Fortress | land | Fixes both colors and supplies a late-game attacking option.
- 1 Scoured Barrens | land | Fixes both colors while providing a lifegain trigger.
- 1 Scrubland | land | Provides untapped, fetchable access to both colors.
- 1 Shattered Sanctum | land | Provides both colors and usually enters untapped after the opening turns.
- 1 Shineshadow Snarl | land | Provides both colors with numerous basic lands available to reveal.
- 1 Silent Clearing | land | Provides both colors and can become a card when mana is plentiful.
- 1 Starting Town | land | Offers early untapped fixing and flexible mana later.
- 1 Tarnished Citadel | land | Provides either color when needed, with lifegain cushioning the damage.
- 1 Vault of Champions | land | Provides efficient access to both colors in multiplayer games.
- 11 Plains | land | Supplies reliable untapped white mana for early support and white-heavy spells.
- 7 Swamp | land | Supplies reliable untapped black mana and supports Crypt Ghast.
- 1 Legion's Landing // Adanto, the First Fort | ramp | Provides an early lifelinking body and can develop into an additional mana source.
- 1 Thieving Varmint | ramp | Accelerates creature deployment and helps pay for creature abilities.
- 1 Pristine Talisman | ramp | Combines repeatable lifegain with additional mana.
- 1 Phial of Galadriel | ramp | Fixes mana while providing recovery when life or cards run low.
- 1 The Celestus | ramp | Provides colored mana with occasional lifegain and hand filtering.
- 1 Hierophant's Chalice | ramp | Adds mana while its arrival gains life and drains an opponent.
- 1 Orazca Relic | ramp | Provides mana early and can later become life and a card.
- 1 Altar of the Pantheon | ramp | Fixes mana and can add lifegain alongside a legendary enchantment.
- 1 Cryptolith Fragment // Aurora of Emrakul | ramp | Fixes mana and adds gradual pressure through its life-loss effect.
- 1 Crypt Ghast | ramp | Expands black mana production and adds an extort outlet.
- 1 Archivist of Oghma | draw | Turns opponents' library searches into cards and lifegain.
- 1 Enduring Innocence | draw | Rewards the deck's small creatures with repeatable card draw.
- 1 Dawn of Hope | draw | Converts lifegain triggers into cards when spare mana is available.
- 1 Lunar Convocation | draw | Uses the deck's life-total movement to generate cards and creatures.
- 1 Tymna the Weaver | draw | Turns successful attacks into additional cards.
- 1 The Gaffer | draw | Rewards substantial lifegain with a steady supply of cards.
- 1 Well of Lost Dreams | draw | Converts lifegain into a scalable burst of card draw.
- 1 Cosmos Elixir | draw | Provides cards at a high life total and lifegain while recovering.
- 1 Mangara, the Diplomat | draw | Draws cards when opponents commit to aggressive attacks or busy turns.
- 1 Vampiric Rites | draw | Turns expendable creatures into cards and lifegain.
- 1 Alseid of Life's Bounty | interaction | Provides early lifelink and protection for an important permanent.
- 1 Faith's Shield | interaction | Protects Karlov or another important permanent for little mana.
- 1 Cecil, Dark Knight // Cecil, Redeemed Paladin | interaction | Provides an inexpensive body that can develop into a protective lifelinking creature.
- 1 Rune-Tail, Kitsune Ascendant // Rune-Tail's Essence | interaction | Uses a high life total to protect the creature board from damage.
- 1 Metropolis Reformer | interaction | Protects against targeted effects while supporting lifegain.
- 1 Sword of Light and Shadow | interaction | Protects an attacker while adding lifegain and creature recovery.
- 1 Werefox Bodyguard | interaction | Offers flexible temporary exile and a lifegain sacrifice option.
- 1 Courageous Resolve | interaction | Protects an important creature and replaces itself.
- 1 Ayli, Eternal Pilgrim | removal | Turns expendable creatures into lifegain and offers removal at a high life total.
- 1 Nightmare's Thirst | removal | Provides inexpensive creature removal that scales with life gained.
- 1 Murderous Rider // Swift End | removal | Answers a creature or planeswalker and later supplies a lifelinking body.
- 1 Solitude | removal | Provides flexible creature exile and a lifelinking threat.
- 1 Gumdrop Poisoner // Tempt with Treats | removal | Combines lifegain support with creature removal.
- 1 Sorin of House Markov // Sorin, Ravenous Neonate | removal | Connects lifegain with removal and direct damage.
- 1 Umezawa's Jitte | removal | Turns combat into reusable creature control or lifegain.
- 1 Aetherflux Reservoir | removal | Builds life through spellcasting and converts a large life reserve into damage.
- 1 Witch of the Moors | removal | Rewards recurring lifegain with opposing sacrifices and creature recovery.
- 1 Blind Obedience | synergy | Slows opposing development while extort supplies recurring lifegain and drain.
- 1 Cleric Class | synergy | Improves lifegain and turns it into additional creature growth.
- 1 Case of the Uneaten Feast | synergy | Makes creature arrivals supply frequent lifegain triggers.
- 1 Serra Ascendant | synergy | Provides an inexpensive lifelinking attacker suited to a high life total.
- 1 Righteous Valkyrie | synergy | Adds creature-based lifegain and rewards maintaining a high life total.
- 1 Ocelot Pride | synergy | Pairs early lifelink with creature production when life has been gained.
- 1 Vito, Thorn of the Dusk Rose | synergy | Converts lifegain into pressure on opponents' life totals.
- 1 Archangel of Thune | threat | Turns repeated lifegain into growth across the creature board.
- 1 Elenda, Saint of Dusk | threat | Provides a resilient lifelinking attacker that rewards a high life total.
- 1 Elenda, the Dusk Rose | threat | Grows through creature deaths and leaves behind lifelinking bodies.
- 1 Gisela, the Broken Blade | threat | Provides efficient flying pressure with lifelink.
- 1 Liesa, Forgotten Archangel | threat | Combines aerial lifelink with creature recovery and disruption.
- 1 Lyra Dawnbringer | threat | Provides a substantial flying finisher and strengthens supporting Angels.
- 1 Cliffhaven Vampire | threat | Turns recurring lifegain into pressure against every opponent.
- 1 Enduring Tenacity | threat | Converts lifegain into damage and can persist beyond its creature form.
- 1 Celestine, the Living Saint | threat | Provides flying lifelink and recovers creatures as life is gained.
- 1 Rhox Faithmender | threat | Magnifies lifegain and provides a durable lifelinking blocker.
- 1 Attended Healer | threat | Builds a wider creature board from repeated lifegain.
- 1 Angel of Destiny | wincon | Provides a lifegain-oriented finisher alongside the combat plan.
- 1 Test of Endurance | wincon | Provides a dedicated payoff for building and protecting a large life total.
- 1 Approach of the Second Sun | wincon | Provides a noncombat finish supported by the deck's draw engines.
- 1 Amalia Benavides Aguirre | wipe | Uses lifegain to improve draws and can clear other creatures after growing.
- 1 Fumigate | wipe | Resets crowded creature boards while producing a burst of lifegain.
- 1 The Meathook Massacre | wipe | Clears smaller creatures and turns subsequent deaths into life-total pressure.

</details>

### 2. aristocrats Commander, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 181 names.

Commander: Denethor, Ruling Steward.

Grade: baseline, score 0.36, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $191.78 the whole deck.

**Summary:** Denethor, Ruling Steward leads a patient aristocrats deck that turns expendable creatures into steady life loss and replacement tokens. Develop mana and card flow early, then use protection and removal to keep the sacrifice plan operating through longer games. Accumulated drain provides the main route to victory, backed by larger finishers when the table needs a stronger closing threat. The deck favors sustained resources and defensive interaction over explosive starts, but depends on maintaining creatures and enough spare mana to keep Denethor working.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The mana base serves one color far better than another, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: Denethor, Skullclamp, Grave Venerations and a few sacrifice outlets point at aristocrats, but four board wipes, many equipment pieces, and a pile of generic ramp and removal pull the list toward a durable good-stuff pile rather than one sacrifice engine.
- PLAN theme_fit=partly: The deck is white-black Commander led by Denethor at a reasonable power level, but it carries few actual sacrifice outlets, death-trigger payoffs, or token makers, so it only loosely answers an aristocrats request.
- PLAN useful_as_built=partly: With 25 basic lands plus many fixers and 10+ mana rocks it will cast its spells, but the 17 Plains to 8 Swamp split strains black cards like Massacre Girl and Archfiend, and the thin creature and sacrifice count means it often lacks fodder or a real closing plan.
- PLAN summary_honest=partly: The summary promises accumulated drain as the main route to victory, but the list has almost no drain effects beyond Grave Venerations and Archfiend of Ifnir, and it also skips mentioning the four sweepers that work against the 'keep creatures' plan.
- [INFO] `curve_summary`: average mana value 2.76 over 63 nonland cards
- [INFO] `mana_pass`: the builder moved 4 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 1 Command Tower | land | Provides flexible color fixing for the deck.
- 1 City of Brass | land | Helps keep both colors available early.
- 1 Exotic Orchard | land | Adds multiplayer color fixing.
- 1 Marsh Flats | land | Finds the basic land needed for the current hand.
- 1 Fabled Passage | land | Balances the deck's white and black mana.
- 1 Evolving Wilds | land | Trades early speed for dependable color fixing.
- 1 Terramorphic Expanse | land | Finds whichever basic color is missing.
- 1 Path of Ancestry | land | Provides both commander colors.
- 1 Thriving Moor | land | Supports the deck's two-color mana base.
- 1 Grand Coliseum | land | Adds another flexible colored source.
- 1 Spire of Industry | land | Uses the artifact package to support color fixing.
- 17 Plains | land | Supplies reliable untapped white mana.
- 8 Swamp | land | Supplies reliable untapped black mana.
- 1 Sol Ring | ramp | Accelerates the commander and supporting permanents.
- 1 Arcane Signet | ramp | Provides inexpensive acceleration and color fixing.
- 1 Fellwar Stone | ramp | Adds another inexpensive mana rock.
- 1 Thought Vessel | ramp | Helps develop mana without requiring a creature.
- 1 Springleaf Drum | ramp | Turns spare creatures into early mana support.
- 1 Lotho, Corrupt Shirriff | ramp | Adds creature-based acceleration to the aristocrats shell.
- 1 Deadly Dispute | ramp | Connects expendable permanents with mana development.
- 1 Wayfarer's Bauble | ramp | Develops the basic-land mana base.
- 1 Relic of Legends | ramp | Supports mana development alongside the legendary creatures.
- 1 Chromatic Lantern | ramp | Makes the two-color mana base more dependable.
- 1 Commander's Sphere | ramp | Adds flexible colored acceleration.
- 1 Inherited Envelope | ramp | Provides another source of colored ramp.
- 1 Skullclamp | draw | Makes expendable creatures a strong source of cards.
- 1 Idol of Oblivion | draw | Supports sustained card flow in the token plan.
- 1 Nasty End | draw | Turns a planned sacrifice into fresh cards.
- 1 Night's Whisper | draw | Provides inexpensive card replenishment.
- 1 Call of the Ring | draw | Supplies continuing card advantage.
- 1 Wall of Omens | draw | Combines card flow with an expendable defensive body.
- 1 Inspiring Overseer | draw | Adds cards while maintaining creature density.
- 1 Folk Hero | draw | Supports card flow alongside the commander and Humans.
- 1 Lembas | draw | Adds inexpensive card flow and a useful artifact.
- 1 Tome of Legends | draw | Provides card advantage tied to the commander.
- 1 The Sackville-Bagginses | draw | Adds a creature-based draw option to the sacrifice shell.
- 1 Massacre Girl, Known Killer | draw | Adds a substantial creature-based card-advantage engine.
- 1 Exemplar of Light | draw | Supplies another creature-based source of card advantage.
- 1 Mask of Memory | draw | Gives the creature package an additional draw tool.
- 1 Boromir, Warden of the Tower | interaction | Protects the board while fitting the sacrifice plan.
- 1 Bastion Protector | interaction | Helps keep the commander on the battlefield.
- 1 Clever Concealment | interaction | Protects an established board from disruption.
- 1 Duty Beyond Death | interaction | Connects defensive interaction with creature sacrifice.
- 1 Reprieve | interaction | Provides inexpensive interaction at a critical moment.
- 1 Lightning Greaves | interaction | Protects important creatures with a low equipment cost.
- 1 Swiftfoot Boots | interaction | Adds another protection option for key creatures.
- 1 Gift of Immortality | interaction | Helps preserve a valuable creature through attrition.
- 1 Unbreakable Formation | interaction | Provides protection for a developed creature board.
- 1 Zack Fair | interaction | Adds protective interaction on a sacrifice-friendly body.
- 1 Darksteel Plate | interaction | Offers lasting protection for an important creature.
- 1 Swords to Plowshares | removal | Answers a dangerous creature efficiently.
- 1 Bitter Triumph | removal | Provides flexible removal at a low mana cost.
- 1 Generous Gift | removal | Answers a broad range of troublesome permanents.
- 1 Get Lost | removal | Adds inexpensive, flexible spot removal.
- 1 Infernal Grasp | removal | Removes a threatening creature at instant speed.
- 1 Fatal Push | removal | Provides cheap removal that suits a sacrifice deck.
- 1 Heartless Act | removal | Adds another efficient creature answer.
- 1 Crib Swap | removal | Provides an additional exile-based creature answer.
- 1 Orcish Bowmasters | removal | Adds removal pressure without reducing creature density.
- 1 Fiend Hunter | removal | Places creature removal on a useful body.
- 1 Palace Jailer | removal | Combines a creature answer with an attrition-oriented body.
- 1 Banishing Light | removal | Expands the range of permanents the deck can answer.
- 1 Gollum the Abandoned | synergy | Adds another creature suited to the aristocrats package.
- 1 Gollum, Patient Plotter | synergy | Provides a reusable piece for sacrifice-based play.
- 1 Gríma Wormtongue | synergy | Supports the deck's creature-based sacrifice plan.
- 1 Nimble Hobbit | synergy | Adds an inexpensive supporting creature.
- 1 Phantom Train | synergy | Adds an artifact synergy piece to the creature-heavy shell.
- 1 Bill the Pony | threat | Adds board presence to support the creature plan.
- 1 Witch-king of Angmar | threat | Provides a marked finisher that also contributes removal pressure.
- 1 Angel of Serenity | threat | Provides a substantial marked finisher with removal utility.
- 1 Archfiend of Ifnir | wincon | Adds a marked finisher for closing longer games.
- 1 Grave Venerations | wincon | Supplies a marked finishing option for the aristocrats plan.
- 1 The Battle of Bywater | wipe | Provides an inexpensive board-reset option.
- 1 Dusk // Dawn | wipe | Supports board control and recovery in a small-creature deck.
- 1 Austere Command | wipe | Offers a flexible reset when the board becomes unfavorable.
- 1 Fumigate | wipe | Provides a dependable answer to a crowded creature board.

</details>

### 3. artifacts Commander, bracket 4

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 318 names.

Commander: Urza, Lord High Artificer.

Grade: typical, score 0.60, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $6837.81 to buy, $6837.81 the whole deck.

**Summary:** Develop artifacts quickly, establish the commander's mana engine, and keep cards flowing while protecting your strongest turns. The deck wins through artifact-backed combat pressure and alternative finishers, with enough interaction to disrupt opponents before committing to a decisive push. It trades broad answers and lean combo lines for an artifact-heavy board that can generate sustained value, but mass artifact removal can force a lengthy rebuild.

The quality model grades this deck typical against the top lists of the format: the deck makes more mana on turn four than the norm, and that raises the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

- PLAN plan_coherent=yes: Nearly every card is a blue or colorless artifact, artifact-synergy piece, cheap ramp, cantrip or counterspell, all feeding Urza's artifact engine, with Aetherflux Reservoir, Mechanized Production and Kappa Cannoneer as clear finishers.
- PLAN theme_fit=yes: The deck is a mono-blue artifact list led by Urza, Lord High Artificer with fast mana such as Sol Ring, Mana Vault and multiple Moxen, free counterspells, and a high-power feel appropriate for bracket 4.
- PLAN useful_as_built=partly: The deck has abundant fast mana, card draw and interaction and is playable, but the mana base includes painlands and multicolor-fixing lands like City of Brass, Mana Confluence and Command Tower that do nothing for a mono-blue deck, and some cards like Unable to Scream and Ravenform are weak filler.
- PLAN summary_honest=partly: The summary claims the deck wins through artifact-backed combat pressure and alternative finishers, which the creature and Aetherflux Reservoir lines support, but it says it trades away lean combo lines while the list holds Krark-Clan Ironworks, Mono-U Transmute Artifact and Whir of Invention, and it is vague about how many finishers are actually present.
- [INFO] `curve_summary`: average mana value 2.97 over 66 nonland cards

<details><summary>The deck list</summary>

- 21 Island | land | Provides reliable blue mana without slowing early development.
- 1 Seat of the Synod | land | Supplies blue mana while contributing to artifact synergies.
- 1 Otawara, Soaring City | land | Provides blue mana and a useful interaction option.
- 1 Command Tower | land | Provides an untapped blue source.
- 1 Spire of Industry | land | Supports blue spells in an artifact-heavy board.
- 1 Mystic Sanctuary | land | Provides blue mana and recovers an important spell.
- 1 City of Brass | land | Provides dependable blue mana for early development.
- 1 Mana Confluence | land | Adds another dependable blue source.
- 1 Academy Ruins | land | Recovers important artifacts for longer games.
- 1 Buried Ruin | land | Returns a key artifact after removal.
- 1 Inventors' Fair | land | Finds a useful artifact once the board is established.
- 1 Mishra's Workshop | land | Accelerates artifact deployment.
- 1 Urza's Saga | land | Develops artifact threats and finds a small utility artifact.
- 1 Sol Ring | ramp | Accelerates the commander and larger artifact plays.
- 1 Mana Vault | ramp | Provides an explosive early burst of mana.
- 1 Chrome Mox | ramp | Trades a card for immediate colored acceleration.
- 1 Mox Diamond | ramp | Turns an extra land into immediate colored acceleration.
- 1 Mox Opal | ramp | Rewards the deck's artifact density with inexpensive acceleration.
- 1 Mox Amber | ramp | Provides efficient mana alongside the deck's legendary permanents.
- 1 Lotus Petal | ramp | Enables faster opening sequences and supplies an artifact cheaply.
- 1 Moonsnare Prototype | ramp | Turns spare permanents into early acceleration.
- 1 Metalworker | ramp | Produces substantial mana from artifact-heavy hands.
- 1 Grand Architect | ramp | Helps deploy artifacts ahead of schedule.
- 1 Krark-Clan Ironworks | ramp | Converts expendable artifacts into mana for larger turns.
- 1 Tezzeret the Seeker | ramp | Untaps mana artifacts and finds useful artifact pieces.
- 1 Inspiring Statuary | ramp | Uses the artifact board to help pay for nonartifact spells.
- 1 Rhystic Study | draw | Maintains card flow while opponents develop their boards.
- 1 Thoughtcast | draw | Provides efficient cards once artifacts are established.
- 1 Thought Monitor | draw | Refills the hand while adding an artifact creature.
- 1 Thirst for Knowledge | draw | Digs for useful cards and stocks the graveyard.
- 1 Sai, Master Thopterist | draw | Turns artifact development into tokens and additional cards.
- 1 Forensic Gadgeteer | draw | Provides repeatable value from casting artifacts.
- 1 Riddlesmith | draw | Filters draws during artifact-heavy turns.
- 1 Vedalken Archmage | draw | Keeps artifact chains supplied with fresh cards.
- 1 Reverse Engineer | draw | Uses the artifact board to support a substantial refill.
- 1 Universal Surveillance | draw | Converts a developed board into a scalable refill.
- 1 Era of Innovation | draw | Turns artifact development into additional cards.
- 1 Esoteric Duplicator | draw | Provides cards and preserves value from sacrificed artifacts.
- 1 Force of Will | interaction | Protects important turns even when mana is committed elsewhere.
- 1 Fierce Guardianship | interaction | Protects the established commander and artifact engine.
- 1 An Offer You Can't Refuse | interaction | Provides inexpensive protection against a critical noncreature spell.
- 1 Metallic Rebuke | interaction | Uses spare artifacts to keep countermagic inexpensive.
- 1 Stoic Rebuttal | interaction | Provides efficient countermagic with an established artifact board.
- 1 Disruption Protocol | interaction | Uses an available artifact to support efficient countermagic.
- 1 Welding Jar | interaction | Protects an important artifact without requiring mana.
- 1 Reality Ripple | interaction | Temporarily removes a problem permanent or protects a key piece.
- 1 Padeem, Consul of Innovation | interaction | Protects artifacts and supports sustained card advantage.
- 1 Ghostly Flicker | interaction | Protects valuable permanents and reuses enter-the-battlefield effects.
- 1 Syr Ginger, the Meal Ender | interaction | Adds a resilient artifact creature that rewards artifact losses.
- 1 Aether Spellbomb | removal | Provides inexpensive creature interaction on an artifact body.
- 1 Unable to Scream | removal | Suppresses a troublesome creature at low cost.
- 1 Resculpt | removal | Answers a problematic artifact or creature.
- 1 Cyber Conversion | removal | Neutralizes a threatening creature efficiently.
- 1 Kitesail Larcenist | removal | Disrupts opposing artifacts and creatures while adding an evasive body.
- 1 Ravenform | removal | Provides another answer to troublesome artifacts and creatures.
- 1 Arcum Dagsson | removal | Converts expendable artifact creatures into important artifact pieces.
- 1 Aetherflux Reservoir | removal | Turns sustained spellcasting into a powerful damage outlet.
- 1 Transmogrifying Wand | removal | Provides repeated creature answers from an artifact permanent.
- 1 Baral's Expertise | removal | Clears several obstacles while advancing your own development.
- 1 Transmute Artifact | synergy | Exchanges an expendable artifact for a more important piece.
- 1 Whir of Invention | synergy | Finds the artifact needed for the current board state.
- 1 Emry, Lurker of the Loch | synergy | Reuses artifacts from the graveyard to sustain the engine.
- 1 Kappa Cannoneer | threat | Provides a powerful threat that rewards continued artifact deployment.
- 1 Cyberdrive Awakener | threat | Turns a developed artifact board into a major aerial attack.
- 1 Phyrexian Metamorph | threat | Copies the most useful artifact or creature available.
- 1 Kuldotha Forgemaster | threat | Turns expendable artifacts into a decisive artifact threat.
- 1 Karn, Scion of Urza | threat | Produces artifact-scaled threats and supports longer games.
- 1 Lodestone Golem | threat | Applies pressure while taxing nonartifact development.
- 1 Master Transmuter | threat | Deploys larger artifacts while reusing valuable artifact effects.
- 1 Filigree Attendant | threat | Provides an evasive threat that scales with the artifact board.
- 1 Frogmite | threat | Adds an inexpensive artifact body once the engine develops.
- 1 Traxos, Scourge of Kroog | threat | Adds substantial combat pressure to artifact-heavy turns.
- 1 Arcbound Reclaimer | threat | Adds an artifact body while recovering important artifacts.
- 1 Argent Sphinx | threat | Provides evasive pressure with protection supported by artifacts.
- 1 Mechanized Production | wincon | Builds artifact value toward an alternative victory.
- 1 Mirrodin Besieged | wincon | Provides an alternative finish supported by the artifact strategy.
- 1 Shimmer Dragon | wincon | Combines a substantial evasive finisher with sustained card advantage.
- 1 Hurkyl's Recall | wipe | Resets an opposing artifact board or rescues your own artifacts.
- 1 Engineered Explosives | wipe | Clears clusters of inexpensive permanents.

</details>

### 4. dinosaur tribal, bracket 2

Format: Commander. Theme: dinosaurs. Pool: any_card. Shortlist: 299 names.

Commander: Gishath, Sun's Avatar.

Grade: bad, score 0.04, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $1106.93 to buy, $1106.93 the whole deck.

**Summary:** Build mana early, establish a Dinosaur herd, then send Gishath into combat to turn damage into more Dinosaurs on the battlefield. The deck wins through large creature attacks, with Bonehoard Dracosaur and Dinosaurs on a Spaceship chief among its dedicated finishers. Card draw helps sustain pressure, while protection, removal, and sweepers support the herd. This creature-focused plan gives up extensive instant-speed answers and quick combo wins, and its expensive threats can make rebuilding after repeated board clears slow.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The land count sits above the norm of the format, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every nonland card is a Dinosaur, a ramp piece, or a draw or protection piece, and together they push toward casting big Dinosaurs and attacking with Gishath.
- PLAN theme_fit=yes: It is a Gishath Dinosaur Commander deck in his three colors and is built around a friendly, creature-based plan suitable for a new player at bracket 2.
- PLAN useful_as_built=partly: The ramp and Dinosaur count make it playable, but the mana base has many tapped and painful lands across three colors, and Blasphemous Act and Harsh Mercy sit awkwardly in a creature deck, so a new player could find it clunky.
- PLAN summary_honest=yes: The summary says Bonehoard Dracosaur and Dinosaurs on a Spaceship are finishers, and both are in the list tagged as wincons, alongside the Dinosaur herd, sweepers like Blasphemous Act and Harsh Mercy, and the equipment protection it describes.
- [INFO] `curve_summary`: average mana value 3.45 over 62 nonland cards

<details><summary>The deck list</summary>

- 1 Sol Ring | ramp | Accelerates the deck toward its large Dinosaurs and commander.
- 1 Birds of Paradise | ramp | Provides early acceleration and fixes all three colors.
- 1 Arcane Signet | ramp | Offers dependable acceleration in any needed color.
- 1 Fellwar Stone | ramp | Adds inexpensive acceleration with potential color fixing.
- 1 Drover of the Mighty | ramp | Fixes mana early and becomes a useful body alongside Dinosaurs.
- 1 Intrepid Paleontologist | ramp | Accelerates mana and helps recover Dinosaurs later.
- 1 Hulking Raptor | ramp | Combines a Dinosaur body with substantial mana acceleration.
- 1 Patchwork Banner | ramp | Fixes mana while strengthening the Dinosaur herd.
- 1 Pillar of Origins | ramp | Provides inexpensive colored mana for Dinosaur spells.
- 1 Topiary Stomper | ramp | Develops the mana base while adding another Dinosaur.
- 1 Garruk's Uprising | draw | Rewards large creatures with cards and helps attacks connect.
- 1 Ripjaw Raptor | draw | Turns damage into cards on a sturdy Dinosaur body.
- 1 Runic Armasaur | draw | Provides recurring card advantage from opposing activated abilities.
- 1 Curious Altisaur | draw | Rewards Dinosaur combat with additional cards.
- 1 Earthshaker Dreadmaw | draw | Refills the hand when a Dinosaur herd is already established.
- 1 Beast Whisperer | draw | Keeps cards flowing when casting creature spells.
- 1 Return of the Wildspeaker | draw | Converts a large creature into cards or strengthens a finishing attack.
- 1 Faithless Looting | draw | Filters early draws and exchanges excess lands for fresh cards.
- 1 Thrill of Possibility | draw | Trades an expendable card for fresh options at instant speed.
- 1 Vanquisher's Banner | draw | Rewards casting Dinosaurs while strengthening the herd.
- 1 Folk Hero | draw | Turns Dinosaur spells into a steady source of cards alongside Gishath.
- 1 Heroic Intervention | interaction | Protects an established board from removal and sweepers.
- 1 Swiftfoot Boots | interaction | Protects an important creature and helps it attack promptly.
- 1 Lightning Greaves | interaction | Offers reusable protection and haste without an equip mana payment.
- 1 And They Shall Know No Fear | interaction | Protects the Dinosaur herd during combat or against damage.
- 1 Temple Altisaur | interaction | Reduces damage to other Dinosaurs and supports safer combat.
- 1 Pteron Ghost | interaction | Provides an inexpensive Dinosaur that can protect another creature.
- 1 Itzquinth, Firstborn of Gishath | removal | Uses a larger Dinosaur to remove an opposing creature.
- 1 Thrashing Brontodon | removal | Provides artifact or enchantment removal on a Dinosaur body.
- 1 Tranquil Frillback | removal | Offers flexible answers to artifacts, enchantments, and graveyards.
- 1 Savage Stomp | removal | Lets a Dinosaur fight an opposing creature efficiently.
- 1 Triumphant Chomp | removal | Turns the herd's largest power into inexpensive creature removal.
- 1 Trapjaw Tyrant | removal | Turns damage into removal for opposing creatures.
- 1 Wrathful Raptors | removal | Turns damage dealt to Dinosaurs into retaliatory damage.
- 1 Kinjalli's Caller | synergy | Makes expensive Dinosaurs easier to cast.
- 1 Otepec Huntmaster | synergy | Reduces Dinosaur costs and helps new Dinosaurs attack immediately.
- 1 Siegehorn Ceratops | synergy | Provides an early Dinosaur that grows when damaged.
- 1 Armored Kincaller | synergy | Adds an inexpensive Dinosaur with flexible defensive utility.
- 1 Belligerent Yearling | synergy | Provides early pressure that scales with larger Dinosaurs.
- 1 Raptor Hatchling | synergy | Turns damage into a larger Dinosaur token.
- 1 Sunfrill Imitator | synergy | Copies a useful Dinosaur while attacking.
- 1 Herald's Horn | synergy | Reduces Dinosaur costs and helps find more members of the herd.
- 1 Urza's Incubator | synergy | Makes the deck's expensive Dinosaur spells substantially easier to cast.
- 1 Kinjalli's Sunwing | synergy | Slows opposing creatures and helps Dinosaur attacks get through.
- 1 Carnage Tyrant | threat | Provides a resilient trampling threat.
- 1 Etali, Primal Storm | threat | Turns attacks into additional spells and mounting pressure.
- 1 Goring Ceratops | threat | Makes a wide Dinosaur attack much more dangerous.
- 1 Crested Herdcaller | threat | Adds two Dinosaur bodies from one card.
- 1 Palani's Hatcher | threat | Develops the herd and helps Dinosaurs attack quickly.
- 1 Pantlaza, Sun-Favored | threat | Rewards Dinosaur arrivals with additional resources.
- 1 Quartzwood Crasher | threat | Turns trampling combat damage into additional large threats.
- 1 Regisaur Alpha | threat | Adds another Dinosaur and gives the herd haste.
- 1 Shifting Ceratops | threat | Provides a flexible attacker with useful combat abilities.
- 1 Rampaging Ceratops | threat | Makes profitable blocking difficult for opponents.
- 1 Thrash of Raptors | threat | Provides a straightforward attacker that improves alongside another Dinosaur.
- 1 Majestic Heliopterus | threat | Helps another Dinosaur attack through the air.
- 1 Bonehoard Dracosaur | wincon | Supplies an evasive finisher that also generates resources.
- 1 Dinosaurs on a Spaceship | wincon | Provides an evasive finisher that strengthens the Dinosaur herd.
- 1 Huatli, Poet of Unity // Roar of the Fifth People | wincon | Smooths early mana and develops into a Dinosaur-focused finishing engine.
- 1 Blasphemous Act | wipe | Resets crowded creature boards and works with damage-triggered Dinosaurs.
- 1 Harsh Mercy | wipe | Clears many opposing creatures while preserving the Dinosaur theme.
- 1 Raging Swordtooth | wipe | Clears small creatures and triggers the herd's damage-based abilities.
- 1 Command Tower | land | Provides unrestricted fixing for all three colors.
- 1 City of Brass | land | Provides any needed color without entering tapped.
- 1 Mana Confluence | land | Provides dependable untapped access to all three colors.
- 1 Path of Ancestry | land | Fixes all three colors and supports Dinosaur card selection.
- 1 Cavern of Souls | land | Fixes Dinosaur mana and helps important Dinosaur spells resolve.
- 1 Secluded Courtyard | land | Provides untapped colored mana for the Dinosaur herd.
- 1 Unclaimed Territory | land | Provides untapped fixing for Dinosaur spells.
- 1 Taiga | land | Provides untapped red and green mana.
- 1 Stomping Ground | land | Fixes red and green with an untapped option.
- 1 Spire Garden | land | Provides red and green mana with reliable multiplayer untapped access.
- 1 Karplusan Forest | land | Provides untapped red and green fixing.
- 1 Grove of the Burnwillows | land | Provides untapped red and green mana.
- 1 Rootbound Crag | land | Provides red and green fixing supported by the typed lands.
- 1 Rockfall Vale | land | Provides red and green fixing with untapped access after the early turns.
- 1 Cinder Glade | land | Provides red and green fixing with useful basic land types.
- 1 Cragcrown Pathway // Timbercrown Pathway | land | Offers an untapped choice between red and green.
- 1 Restless Ridgeline | land | Fixes red and green and supplies a late-game creature option.
- 1 Savannah | land | Provides untapped green and white mana.
- 1 Temple Garden | land | Fixes green and white with an untapped option.
- 1 Bountiful Promenade | land | Provides green and white mana with reliable multiplayer untapped access.
- 1 Brushland | land | Provides untapped green and white fixing.
- 1 Sunpetal Grove | land | Provides green and white fixing supported by the typed lands.
- 1 Canopy Vista | land | Provides green and white fixing with useful basic land types.
- 1 Horizon Canopy | land | Provides green and white mana and can become a card later.
- 1 Branchloft Pathway // Boulderloft Pathway | land | Offers an untapped choice between green and white.
- 1 Plateau | land | Provides untapped red and white mana.
- 1 Sacred Foundry | land | Fixes red and white with an untapped option.
- 1 Spectator Seating | land | Provides red and white mana with reliable multiplayer untapped access.
- 1 Battlefield Forge | land | Provides untapped red and white fixing.
- 1 Clifftop Retreat | land | Provides red and white fixing supported by the typed lands.
- 1 Sunbaked Canyon | land | Provides red and white mana and can become a card later.
- 1 Needleverge Pathway // Pillarverge Pathway | land | Offers an untapped choice between red and white.
- 3 Forest | land | Supports early green spells and supplies basic-land search options.
- 1 Mountain | land | Provides an untapped red source and a basic-land search option.
- 1 Plains | land | Provides an untapped white source and a basic-land search option.

</details>

### 5. blink Commander, owned first

Format: Commander. Theme: blink. Pool: owned_first. Shortlist: 169 names.

Commander: Gilraen, Dúnedain Protector.

Grade: typical, score 0.43, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $89.92 the whole deck.

**Summary:** Gilraen leads a creature-focused blink deck that balances draw and removal bodies with inexpensive mana support. Develop the board, keep mana available for Gilraen, and use the interaction package to preserve your most valuable creatures. Angels provide the closing pressure, chief among them Angel of Serenity, while the other creatures help turn a stable board into a winning attack. The deck favors steady value and combat over explosive finishes, so it gives up speed and depends on keeping its creature engine intact.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

This deck plan aims for 2 or more finishers, the cards that win the game, and the deck holds 1.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: Mono-white blink pieces like Flickerwisp, Palace Jailer, Fiend Hunter and Angel of Serenity share a plan, but three board wipes, a stack of mana rocks and equipment and non-blink filler pull it toward a generic pile of goodstuff.
- PLAN theme_fit=partly: It is a mono-white Commander deck led by Gilraen with real blink enablers such as Flickerwisp, Palace Jailer and Fiend Hunter, but the request for a blink deck is only lightly served, since few cards actually reflicker and Gilraen is not a blink commander.
- PLAN useful_as_built=partly: With 37 lands, plenty of cheap rocks and draw, and lots of removal it will function in play, but the 33 Plains plus rocks is flood-prone, and there are few ways to win beyond Angel of Serenity and Angel of Sanctions.
- PLAN summary_honest=partly: The summary names Angel of Serenity as the chief closer and Angels as the finishing pressure, and those are in the list, but it calls the deck a balanced creature blink deck while the list runs three wipes and about a dozen mana rocks and holds only one real finisher.
- [INFO] `curve_summary`: average mana value 2.82 over 62 nonland cards
- [WARN] `finisher_short`: the finisher count is 1, and this deck plan wants 2 or more (Angel of Serenity)

<details><summary>The deck list</summary>

- 33 Plains | land | Provides a dependable white mana base with minimal setup.
- 1 Command Tower | land | Adds another dependable white source.
- 1 Minas Tirith | land | Adds utility without moving away from the white mana base.
- 1 War Room | land | Provides a useful mana-base slot for longer games.
- 1 Windbrisk Heights | land | Supports the deck's creature-heavy battlefield plan.
- 1 Wall of Omens | draw | Provides an inexpensive creature-based draw piece for the blink plan.
- 1 Inspiring Overseer | draw | Combines the draw package with a creature for Gilraen.
- 1 Lembas | draw | Adds inexpensive draw support without requiring another creature.
- 1 Mask of Memory | draw | Supports card flow alongside the deck's attacking creatures.
- 1 Skullclamp | draw | Adds efficient draw support to a creature-focused deck.
- 1 Tome of Legends | draw | Provides a compact draw piece for a commander-centered plan.
- 1 Folk Hero | draw | Adds another source of cards to sustain creature deployment.
- 1 Mirror of Galadriel | draw | Supplies draw support for slower, resource-intensive games.
- 1 Faramir, Field Commander | draw | Keeps the draw package attached to a creature.
- 1 Exemplar of Light | draw | Adds an Angel to the creature-based draw package.
- 1 Joined Researchers // Secret Rendezvous | draw | Adds another draw option while preserving the creature-focused theme.
- 1 Puresteel Paladin | draw | Supports card flow alongside the deck's equipment package.
- 1 Sol Ring | ramp | Accelerates both creature deployment and Gilraen's activation budget.
- 1 Arcane Signet | ramp | Adds inexpensive mana support for white spells.
- 1 Thought Vessel | ramp | Provides early acceleration for the deck's middle turns.
- 1 Fellwar Stone | ramp | Adds another inexpensive mana rock.
- 1 Springleaf Drum | ramp | Turns the creature-heavy board into additional mana support.
- 1 Wayfarer's Bauble | ramp | Develops the mana base early.
- 1 Sword of the Animist | ramp | Pairs mana development with the combat plan.
- 1 Relic of Legends | ramp | Supports the mana demands of the deck's legendary creatures.
- 1 Commander's Sphere | ramp | Provides another mana source for sustained development.
- 1 Giada, Font of Hope | ramp | Supports the Angel portion of the creature package.
- 1 White Auracite | ramp | Adds mana support for the deck's larger plays.
- 1 PuPu UFO | ramp | Keeps an additional ramp slot attached to a creature.
- 1 Boromir, Warden of the Tower | interaction | Adds creature-based interaction to support a developed board.
- 1 Lightning Greaves | interaction | Supports keeping Gilraen available throughout the game.
- 1 Swiftfoot Boots | interaction | Adds another equipment-based interaction piece for key creatures.
- 1 Clever Concealment | interaction | Reserves an interaction slot for protecting the deck's investment.
- 1 Reprieve | interaction | Provides inexpensive interaction while developing the board.
- 1 Unbreakable Formation | interaction | Supports committing creatures without abandoning defensive interaction.
- 1 Zack Fair | interaction | Adds inexpensive creature-based interaction.
- 1 Gift of Immortality | interaction | Supports retaining an important creature through disruption.
- 1 Bastion Protector | interaction | Reinforces the commander-centered game plan.
- 1 Darksteel Plate | interaction | Adds a durable interaction piece for an important creature.
- 1 Swords to Plowshares | removal | Provides efficient removal when a creature must be answered.
- 1 Generous Gift | removal | Adds a flexible answer to the removal package.
- 1 Get Lost | removal | Provides inexpensive removal while preserving mana for development.
- 1 March of Otherworldly Light | removal | Adds a flexible removal option across different stages of the game.
- 1 Crib Swap | removal | Adds another instant-speed answer.
- 1 Destroy Evil | removal | Diversifies the inexpensive removal package.
- 1 Palace Jailer | removal | Places removal on a creature for the blink-focused shell.
- 1 Fiend Hunter | removal | Adds another creature-based removal piece.
- 1 Westfold Rider | removal | Keeps part of the removal package on an inexpensive creature.
- 1 Stroke of Midnight | removal | Adds another flexible answer for troublesome permanents.
- 1 Flickerwisp | synergy | Reinforces the central blink strategy with another creature.
- 1 Personify | synergy | Adds dedicated support for the blink plan.
- 1 Ennis, Debate Moderator | synergy | Adds a creature-based synergy piece alongside Gilraen.
- 1 Jocasta, Automaton Avenger | synergy | Provides another creature-based synergy slot.
- 1 Together Forever | synergy | Supports maintaining the creature engine over a longer game.
- 1 Angel of Condemnation | synergy | Bridges the blink plan and the creature-based removal package.
- 1 Slip On the Ring | synergy | Adds blink-focused interaction without requiring another permanent.
- 1 Champions of Minas Tirith | threat | Adds a substantial creature while retaining draw support.
- 1 Bronze Guardian | threat | Adds battlefield pressure alongside the artifact-heavy support package.
- 1 South Pole Voyager | threat | Adds another creature that combines pressure with draw support.
- 1 The Vision | threat | Provides a creature threat without giving up a draw-oriented slot.
- 1 Summon: Primal Garuda | threat | Adds a larger creature threat with a removal role.
- 1 Frontline Medic | threat | Contributes to the attacking creature group while retaining interaction value.
- 1 Angel of Serenity | wincon | Provides the deck's explicitly identified finisher.
- 1 Angel of Sanctions | wincon | Combines a closing Angel threat with the removal plan.
- 1 Austere Command | wipe | Provides a flexible reset when the battlefield becomes unfavorable.
- 1 Dusk // Dawn | wipe | Adds a reset suited to a creature-focused deck.
- 1 Split Up | wipe | Adds another reset without crowding the deck's upper curve.

</details>

### 6. Modern tempo, tournament

Format: Modern. Theme: tempo. Pool: any_card. Shortlist: 167 names.

Grade: typical, score 0.54, model 20260923T202806Z.

Cards: 60 main, 15 sideboard. Repair turn: no. Block findings: 0.

Cost: $652.08 to buy, $652.08 the whole deck.

**Summary:** This blue-red tempo deck establishes creature pressure, then uses cheap removal and counters to keep the opponent from stabilizing. Card selection keeps useful spells flowing, while larger flying threats provide staying power when early attacks stall. It wins primarily through combat, with burn offering a final push. The trade-offs are a painful mana base, limited answers to resolved noncreature permanents, and less inevitability than a dedicated control deck.

The quality model grades this deck typical against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The deck runs its spells as playsets, and that lowers the grade. The cards are ones the top lists of the format play, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Cheap creatures, burn, counters and cantrips all serve one tempo plan of pressuring early and protecting the lead.
- PLAN theme_fit=partly: It is a blue-red tempo deck in Modern, but it has no Delver of Secrets or other Delver-style flip creature, so the named Delver theme is only loosely met.
- PLAN useful_as_built=yes: It has 24 lands with many dual lands, a very low curve, 12 cantrips or cheap artifacts, and enough creatures and burn to win, so it plays as it stands.
- PLAN summary_honest=yes: x
- [INFO] `curve_summary`: average mana value 1.67 over 36 nonland cards
- [INFO] `mana_pass`: the builder moved 1 card of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 4 Ragavan, Nimble Pilferer | threat | Provides early pressure and rewards connecting with extra resources.
- 4 Ledger Shredder | threat | Turns frequent cheap spells into card selection and a growing flying attacker.
- 4 Quantum Riddler | threat | Supplies a substantial flying threat and helps replenish resources.
- 4 Consider | draw | Finds threats or interaction while supporting instant-speed play.
- 2 Preordain | draw | Smooths opening draws and searches for the right spell.
- 4 Mishra's Bauble | synergy | Enables inexpensive double-spell turns for Ledger Shredder while replacing itself.
- 4 Counterspell | interaction | Protects your advantage by answering opposing spells cleanly.
- 2 Force of Negation | interaction | Provides protection against important noncreature spells when mana is committed elsewhere.
- 4 Lightning Bolt | removal | Clears small blockers and supplies reach to finish close games.
- 4 Galvanic Discharge | removal | Offers efficient creature removal to keep attacks profitable.
- 4 Steam Vents | land | Provides both colors with an untapped option for early tempo plays.
- 4 Spirebluff Canal | land | Supplies both colors efficiently during the opening turns.
- 4 Shivan Reef | land | Provides reliable untapped access to either color.
- 2 Fiery Islet | land | Fixes both colors and converts surplus mana into another card.
- 7 Island | land | Supports double-blue spells without additional life loss.
- 3 Mountain | land | Provides painless red mana for early threats and removal.

</details>

### 7. Modern burn, casual

Format: Modern. Theme: burn. Pool: any_card. Shortlist: 290 names.

Grade: typical, score 0.46, model 20260923T202806Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $55.67 to buy, $55.67 the whole deck.

**Summary:** This mono-red burn deck pairs direct damage with a substantial creature package, aiming to establish early pressure and finish through repeated attacks and burn. Draw support helps sustain the plan when games run longer. It leans toward burn-based midrange rather than all-out speed, giving up the fastest starts in exchange for a heavier threat suite and a straightforward mana base.

The quality model grades this deck typical against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The burn spells and Torbran point toward a burn plan, but Eidolon of the Great Revel hurts the deck's own casting of its four-drops, and the 24 lands with few cheap creatures leave the aggro and midrange halves poorly joined.
- PLAN theme_fit=yes: The deck is entirely mono-red, is built around Lightning Bolt, Burst Lightning, Boltwave and Torbran, and is casual in tone, so it matches the request for a mono-red burn deck.
- PLAN useful_as_built=partly: The mana base is clean and it can be played as it stands, but 24 lands is too many for a burn deck, only 6 cheap burn spells besides Boltwave leave few ways to close, and Eidolon punishes its own four-mana threats.
- PLAN summary_honest=partly: The summary claims a substantial creature package and draw support, which the list does carry through Ashcloud Phoenix, Tectonic Giant, Hazoret, Risk Factor and Browbeat, but it presents them as sustaining draw when Risk Factor and Browbeat are really burn-to-face options and it never mentions the heavy 24-land count.
- [INFO] `curve_summary`: average mana value 2.72 over 36 nonland cards

<details><summary>The deck list</summary>

- 20 Mountain | land | Forms the dependable core of the mono-red mana base.
- 4 Ramunap Ruins | land | Rounds out the mana base without adding another color.
- 4 Lightning Bolt | removal | Provides a core removal slot for the burn plan.
- 2 Burst Lightning | removal | Complements the primary removal package.
- 4 Boltwave | synergy | Keeps the spell package focused on the burn strategy.
- 4 Eidolon of the Great Revel | synergy | Adds a creature-based synergy piece alongside the burn spells.
- 4 Risk Factor | draw | Supplies the main draw package for sustained pressure.
- 2 Browbeat | draw | Adds further draw support for games that last beyond the opening turns.
- 2 Chandra, Dressed to Kill | ramp | Provides ramp support for the substantial creature package.
- 4 Ashcloud Phoenix | threat | Supplies a full playset of threats for the creature-pressure plan.
- 4 Tectonic Giant | threat | Adds another substantial threat package for longer games.
- 3 Hazoret the Fervent | threat | Diversifies the threats used to close the game.
- 3 Torbran, Thane of Red Fell | threat | Rounds out the mono-red threat suite.

</details>

### 8. Modern lifegain, FNM

Format: Modern. Theme: lifegain. Pool: any_card. Shortlist: 301 names.

Grade: typical, score 0.42, model 20260923T202806Z.

Cards: 60 main, 15 sideboard. Repair turn: no. Block findings: 0.

Cost: $702.99 to buy, $702.99 the whole deck.

**Summary:** This white-black lifegain deck plays a creature-focused midrange game: establish a lifegain engine, keep cards flowing, and use removal to make room for increasingly substantial threats. It wins through sustained pressure rather than an all-in combo, with Enduring Tenacity, Sheoldred, the Apocalypse, and the Angels providing its finishing muscle. The tradeoff is speed: its heavier threats favor longer games, and its creature-heavy plan can struggle against repeated board wipes or opponents that largely ignore the battlefield.

The quality model grades this deck typical against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade. Many of the cards appear in no top list, and that lowers the grade.

- PLAN plan_coherent=yes: Soul Warden, Guide of Souls, Enduring Tenacity, Archangel of Thune and Elenda all reward gaining life, while Solitude and Gumdrop Poisoner remove blockers and Sheoldred adds pressure, so the cards serve one lifegain-midrange plan.
- PLAN theme_fit=yes: The deck is white-black, built around lifegain payoffs like Guide of Souls, Soul Warden and Archangel of Thune, and made of Modern-playable cards at a casual FNM power level.
- PLAN useful_as_built=partly: The mana base of 24 lands with many dual lands supports the WW and BB costs, but Enduring Innocence at 4 copies and Soul Warden are weak, the deck has few ways to gain life, and the many one-drops and top-end creatures pull the curve in two directions.
- PLAN summary_honest=yes: The summary names Enduring Tenacity, Sheoldred, and the Angels as finishers, and the list runs 4 Tenacity, 3 Sheoldred, 2 Archangel of Thune and 1 Gisela, so the claim of a creature-heavy midrange deck with removal matches the cards.
- [INFO] `curve_summary`: average mana value 3.06 over 36 nonland cards

<details><summary>The deck list</summary>

- 4 Godless Shrine | land | Core fixing for the deck's white and black spells.
- 4 Caves of Koilos | land | Supports both colors while keeping the mana base responsive.
- 4 Fetid Heath | land | Helps support the demanding colored costs among the threats.
- 4 Isolated Chapel | land | Additional white-black fixing alongside the basic lands.
- 4 Plains | land | Reliable white mana for the lifegain engine.
- 4 Swamp | land | Reliable black mana for threats and removal.
- 4 Soul Warden | synergy | A foundational piece of the lifegain engine.
- 4 Guide of Souls | synergy | Reinforces the creature-focused lifegain plan.
- 2 Legion's Landing // Adanto, the First Fort | ramp | Provides development that fits the creature-heavy strategy.
- 4 Enduring Innocence | draw | The main source of sustained card advantage.
- 2 Preacher of the Schism | draw | Adds another card-advantage option without abandoning the creature plan.
- 4 Solitude | removal | Primary creature removal for protecting the deck's board position.
- 2 Gumdrop Poisoner // Tempt with Treats | removal | Additional removal that fits the lifegain theme.
- 2 Ajani, Caller of the Pride | threat | Adds a noncreature threat to diversify the pressure.
- 4 Enduring Tenacity | threat | A central threat for the lifegain-focused game plan.
- 3 Sheoldred, the Apocalypse | threat | A substantial threat for grinding through longer games.
- 2 Archangel of Thune | threat | A high-impact threat that complements the lifegain engine.
- 2 Elenda, Saint of Dusk | threat | Adds another substantial threat aligned with the deck's theme.
- 1 Gisela, the Broken Blade | threat | Diversifies the finishing threats without crowding the deck with extra copies.

</details>

### 9. Standard midrange, FNM

Format: Standard. Theme: midrange. Pool: any_card. Shortlist: 164 names.

Grade: baseline, score 0.43, model 20260923T202806Z.

Cards: 60 main, 15 sideboard. Repair turn: no. Block findings: 0.

Cost: $157.04 to buy, $157.04 the whole deck.

**Summary:** This black-green midrange deck develops a creature board, clears away blockers, and keeps drawing cards through drawn-out exchanges. Its creatures combine pressure with card advantage, while protection and creature-enhancing support help maintain a threatening battlefield. It wins through sustained combat rather than a single explosive turn, trading outright speed for resilience and staying power. The sideboard broadens its answers against aggressive boards, graveyard strategies, and troublesome artifacts or enchantments.

The quality model grades this deck at the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. The color sources cover the pips, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The deck is black-green midrange with removal, card draw and creatures, but Massacre Girl's board wipe, Not Dead After All and the Room and Siege enchantments pull in slightly different directions from a creature-pressure plan.
- PLAN theme_fit=yes: The deck is a black-green midrange list in Standard with creatures, removal and card advantage, which matches the FNM-level request.
- PLAN useful_as_built=partly: The 24-land mana base covers both colors well with dual lands, but only 12 creatures and just 2 Llanowar Elves give a thin board, and the top end of 4 Massacre Girl-type effects plus several situational spells makes it clunky to play.
- PLAN summary_honest=partly: The summary says the deck clears away blockers and keeps drawing cards, which the removal, Massacre Girl, Darkstar Augur and Midnight Reaper support, but it mentions a sideboard that broadens answers to graveyard strategies and artifacts, and no sideboard appears in the list.
- [INFO] `curve_summary`: average mana value 2.97 over 36 nonland cards

<details><summary>The deck list</summary>

- 4 Blooming Marsh | land | Early black-green fixing keeps the opening turns smooth.
- 4 Overgrown Tomb | land | Flexible black-green fixing supports demanding creature costs.
- 4 Wastewood Verge | land | Provides additional fixing alongside the basic lands.
- 4 Restless Cottage | land | Fixes both colors and provides a threat in longer games.
- 4 Forest | land | Reliable green mana for early development.
- 4 Swamp | land | Reliable black mana for removal and card advantage.
- 2 Llanowar Elves | ramp | Accelerates the deck into its impactful three- and four-mana plays.
- 4 Unholy Annex // Ritual Chamber | draw | Supplies sustained card advantage and a late-game threat.
- 2 Hollowmurk Siege | draw | Adds a flexible source of advantage for grinding games.
- 2 Assassin's Trophy | removal | Answers troublesome permanents beyond opposing creatures.
- 2 Bitter Triumph | removal | Provides efficient interaction against creatures and planeswalkers.
- 2 Nowhere to Run | removal | Handles smaller creatures while improving subsequent removal.
- 4 Innkeeper's Talent | synergy | Builds up the creature board and rewards keeping threats in play.
- 2 Snakeskin Veil | synergy | Protects an invested creature while contributing a counter.
- 2 Not Dead After All | synergy | Helps preserve valuable creatures through removal and combat.
- 4 Darkstar Augur | threat | Combines evasive pressure with ongoing card advantage.
- 4 Surrak, Elusive Hunter | threat | Provides efficient pressure and punishes opposing interaction.
- 3 Midnight Reaper | threat | Attacks while helping replenish resources when creatures die.
- 3 Massacre Girl, Known Killer | threat | Provides a substantial combat threat with additional card-advantage potential.

</details>

### 10. Standard aggro, tournament

Format: Standard. Theme: aggro. Pool: any_card. Shortlist: 294 names.

Grade: bad, score 0.30, model 20260923T202806Z.

Cards: 60 main, 15 sideboard. Repair turn: no. Block findings: 0.

Cost: $135.06 to buy, $135.06 the whole deck.

**Summary:** This red-white aggro deck applies creature pressure, uses removal and interaction to keep attacks productive, and draws on Fugitive Codebreaker and Tersa Lightshatter to sustain its offense. It wins through repeated combat rather than a separate combo finish. The sideboard offers more disruption, removal, and board resets when straightforward aggression is not enough. The deck gives up some long-game staying power and flexibility to keep its main plan focused on attacking.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The cards pair in ways the top lists do not, and that lowers the grade. The curve sits high for the format, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The deck is a red-white creature deck with cheap removal and auras, but the curve tops out with four copies each of 5-drop Starfield Shepherd and 6-drop Twinmaw Stormbrood, which pull against a fast aggro plan.
- PLAN theme_fit=yes: The deck is built in red and white with creatures, burn and auras aimed at attacking, in the Standard format the person asked for.
- PLAN useful_as_built=partly: The mana base of 12 dual lands plus basics supports both colors, but the deck has only 4 two-drops of real board presence, plenty of four- to six-drops and no sideboard visible in the list, so it will stumble against faster decks.
- PLAN summary_honest=partly: The summary says Fugitive Codebreaker and Tersa Lightshatter sustain the offense, and both are in the list, but Tersa is only a 2-of and the summary omits the top-heavy curve that undercuts its claim of focused aggression.
- [INFO] `curve_summary`: average mana value 3.28 over 36 nonland cards
- [INFO] `mana_pass`: the builder moved 1 card of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 4 Delivery Moogle | threat | Provides creature pressure for the aggressive plan.
- 4 Starfield Shepherd | threat | Adds another creature threat to sustain attacks.
- 4 Diversion Specialist | threat | Gives the deck a substantial follow-up threat.
- 4 Skyknight Squire | synergy | Supplies creature-based synergy alongside the main attackers.
- 4 Fugitive Codebreaker | draw | Combines a creature slot with card-flow support.
- 2 Tersa Lightshatter | draw | Adds card-flow support without crowding out the attacking core.
- 4 Case of the Gateway Express | removal | Provides removal support for the creature-heavy plan.
- 4 Twinmaw Stormbrood // Charring Bite | removal | Adds removal while retaining a creature option.
- 4 Sheltered by Ghosts | interaction | Supplies interaction to support continued pressure.
- 2 Boros Charm | interaction | Adds instant-speed interaction to the aggressive core.
- 4 Inspiring Vantage | land | Provides red-white fixing for the creature curve.
- 4 Sacred Foundry | land | Supports both colors without requiring a tapped opening.
- 4 Sunbillow Verge | land | Adds red-white fixing alongside the typed lands.
- 7 Plains | land | Supports white spells and the fixing package.
- 5 Mountain | land | Supports red spells and the fixing package.

</details>

### 11. Commander with a locked card

Format: Commander. Theme: sacrifice. Pool: any_card. Shortlist: 299 names.

Commander: Karlov of the Ghost Council.

Grade: baseline, score 0.06, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $898.43 to buy, $898.43 the whole deck.

**Summary:** Karlov leads a sacrifice engine that turns expendable creatures into cards, mana, removal, and repeated life gain. Develop fodder and sacrifice outlets, then layer death payoffs so each exchange advances your board while wearing opponents down. Life gain builds Karlov into a combat threat and fuels his creature removal, while sacrifice-driven life loss and evasive attackers provide other ways to finish. A substantial mana base and recovery effects support longer games, at the cost of explosive speed and direct interaction with opposing spells.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The land count sits above the norm of the format, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every card is a sacrifice outlet, fodder, death payoff, or recursion piece, so the list serves one aristocrats plan around Karlov.
- PLAN theme_fit=yes: It is a Commander deck led by Karlov of the Ghost Council in white-black and is built around sacrifice outlets and death payoffs, exactly as requested.
- PLAN useful_as_built=partly: The deck has plenty of outlets, fodder, and payoffs, but the land count is high with heavy black requirements and several painful or tapped lands, and the fodder count is modest, and the two board wipes work against its own creature plan, so it will function but with clunky draws.
- PLAN summary_honest=partly: The claim of repeated life gain that grows Karlov is only lightly supported, since few cards gain life (Cartel Aristocrat, Ayli, Vindictive Vampire, Zulaport), and Toxic Deluge and Austere Command are wipes that the summary does not mention.
- [INFO] `curve_summary`: average mana value 3.11 over 61 nonland cards
- [INFO] `mana_pass`: the builder moved 2 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 1 Brightclimb Pathway // Grimclimb Pathway | land | Provides an untapped source of the color most needed early.
- 1 Caves of Koilos | land | Supplies either color without slowing early development.
- 1 City of Brass | land | Keeps both colors available for the creature engine.
- 1 Command Tower | land | Provides dependable access to both colors.
- 1 Concealed Courtyard | land | Supports early creatures and an on-time Karlov.
- 1 Forbidden Orchard | land | Provides flexible colored mana for early development.
- 1 Godless Shrine | land | Provides both colors and supports the fetch lands.
- 1 Isolated Chapel | land | Provides both colors alongside the basic and typed lands.
- 1 Mana Confluence | land | Supplies whichever color the sacrifice engine needs.
- 1 Marsh Flats | land | Finds the appropriate basic or white-black dual land.
- 1 Prismatic Vista | land | Finds the missing basic color early.
- 1 Scrubland | land | Provides both colors untapped and supports Marsh Flats.
- 1 Shattered Sanctum | land | Provides both colors for the deck's middle turns.
- 1 Shineshadow Snarl | land | Uses the substantial basic-land base to provide early fixing.
- 1 Silent Clearing | land | Provides both colors and converts into a card when mana is plentiful.
- 1 Starting Town | land | Supports early colored-mana development.
- 1 Tarnished Citadel | land | Provides additional access to either color when necessary.
- 1 Vault of Champions | land | Provides efficient white-black fixing at a multiplayer table.
- 1 Phyrexian Tower | land | Turns an expendable creature into a burst of black mana.
- 1 High Market | land | Adds a sacrifice outlet that also supplies a life-gain trigger.
- 7 Plains | land | Provides reliable white mana for Karlov and the supporting creatures.
- 11 Swamp | land | Supports the deck's black creatures and demanding black costs.
- 1 Sol Ring | ramp | Accelerates the deck's engines and larger threats.
- 1 Ashnod, Flesh Mechanist | ramp | Turns attacking sacrifice fodder into Powerstone mana.
- 1 Culling the Weak | ramp | Converts a disposable creature into an explosive mana burst.
- 1 Priest of Forgotten Gods | ramp | Turns spare creatures into mana while pressuring opposing boards.
- 1 Soldevi Adnate | ramp | Converts suitable creatures into black mana for larger plays.
- 1 Deadly Dispute | ramp | Trades expendable material for cards and a Treasure.
- 1 Ashnod's Altar | ramp | Provides a free sacrifice outlet that funds additional plays.
- 1 Phyrexian Altar | ramp | Converts creatures into whichever colored mana is needed.
- 1 Pitiless Plunderer | ramp | Makes creature deaths replenish mana through Treasures.
- 1 Warren Soultrader | ramp | Turns creatures into Treasures while triggering sacrifice payoffs.
- 1 Village Rites | draw | Converts an expendable or threatened creature into fresh cards.
- 1 Corrupted Conviction | draw | Provides inexpensive card draw through a sacrifice.
- 1 Nasty End | draw | Refills the hand and rewards sacrificing a legendary creature.
- 1 Vampiric Rites | draw | Provides repeatable sacrifice-based draw with accompanying life gain.
- 1 Disciple of Bolas | draw | Turns a substantial creature into cards and life.
- 1 Smothering Abomination | draw | Rewards repeated sacrifices with a steady stream of cards.
- 1 Jerren, Corrupted Bishop // Ormendahl, the Corrupter | draw | Supports the creature engine and offers a powerful late-game draw outlet.
- 1 High-Society Hunter | draw | Combines a sacrifice engine with card draw from nontoken creature deaths.
- 1 Ecstatic Awakener // Awoken Demon | draw | Turns early sacrifice fodder into a card and a larger body.
- 1 Tevesh Szat, Doom of Fools | draw | Produces sacrifice fodder and converts creatures into cards.
- 1 Baron Bertram Graywater | draw | Supports token production and turns tokens into cards.
- 1 Cartel Aristocrat | interaction | Provides a free sacrifice outlet that can protect itself.
- 1 Fanatical Devotion | interaction | Uses expendable creatures to protect important creatures.
- 1 Flare of Fortitude | interaction | Protects the developed board during a critical opposing turn.
- 1 Spirit Bonds | interaction | Adds Spirit fodder and helps protect key nontoken creatures.
- 1 Nightmare Shepherd | interaction | Preserves useful nontoken creature abilities through token copies.
- 1 Promise of Tomorrow | interaction | Helps recover creatures after the board is emptied.
- 1 Rescue from the Underworld | interaction | Uses a sacrifice to recover an important creature from the graveyard.
- 1 Ayli, Eternal Pilgrim | removal | Combines sacrifice-based life gain with late-game permanent removal.
- 1 Yawgmoth, Thran Physician | removal | Turns spare creatures into cards and opposing-creature suppression.
- 1 Grave Pact | removal | Makes each friendly creature death pressure opposing creature boards.
- 1 Dictate of Erebos | removal | Adds another persistent payoff that converts deaths into opposing sacrifices.
- 1 Attrition | removal | Turns expendable creatures into repeatable targeted removal.
- 1 Bone Shards | removal | Provides cheap creature or planeswalker removal using spare material.
- 1 Eaten Alive | removal | Exiles a problematic creature or planeswalker while enabling a sacrifice.
- 1 Teysa, Orzhov Scion | removal | Produces Spirit fodder and converts white creatures into exile removal.
- 1 Ruthless Lawbringer | removal | Turns a spare creature into removal for a troublesome nonland permanent.
- 1 Viscera Seer | synergy | Provides a cheap free sacrifice outlet and improves future draws.
- 1 Carrion Feeder | synergy | Provides another inexpensive free outlet that grows with sacrifices.
- 1 Bastion of Remembrance | synergy | Supplies fodder and turns friendly creature deaths into life swings.
- 1 Elas il-Kor, Sadistic Pilgrim | synergy | Links incoming creatures to life gain and departing creatures to opponent life loss.
- 1 Zulaport Cutthroat | synergy | Makes friendly creature deaths drain opponents while gaining life.
- 1 Vengeful Bloodwitch | synergy | Adds another creature-death payoff that gains life and drains an opponent.
- 1 Victimize | synergy | Trades one expendable creature for two important creatures from the graveyard.
- 1 Basri's Lieutenant | threat | Rewards counters by leaving replacement bodies after creature deaths.
- 1 Felisa, Fang of Silverquill | threat | Converts counters on dying nontoken creatures into flying token pressure.
- 1 Ghoulcaller Gisa | threat | Converts a large creature into a supply of Zombie attackers and sacrifice fodder.
- 1 Liesa, Forgotten Archangel | threat | Provides a lifelinking attacker and helps recover nontoken creatures.
- 1 Mondrak, Glory Dominus | threat | Amplifies token production and can protect itself through sacrifices.
- 1 Ratadrabik of Urborg | threat | Keeps useful legendary-creature abilities present after those creatures die.
- 1 Prowling Geistcatcher | threat | Grows through sacrifices and helps recover the sacrificed creatures later.
- 1 Requiem Angel | threat | Turns many creature deaths into flying Spirit replacements.
- 1 Demon of Catastrophes | threat | Turns expendable fodder into a substantial evasive attacker.
- 1 Vindictive Vampire | threat | Adds another death-triggered drain effect to the creature board.
- 1 Sadistic Hypnotist | threat | Converts spare creatures into pressure on opponents' hands.
- 1 Razaketh, the Foulblooded | wincon | Turns sacrifice fodder into the pieces needed to close the game.
- 1 Relic Vial | wincon | Pairs with the deck's Clerics to turn repeated creature deaths into table-wide drain.
- 1 Sephiroth, Fabled SOLDIER // Sephiroth, One-Winged Angel | wincon | Rewards the sacrifice engine and supplies a dedicated finishing payoff.
- 1 The Meathook Massacre | wipe | Clears smaller creatures while adding life-gain and life-loss payoffs.
- 1 Toxic Deluge | wipe | Provides an efficient reset against overwhelming creature boards.
- 1 Austere Command | wipe | Offers a selective reset that can preserve useful parts of the engine.

</details>

### 12. Commander on a budget

Format: Commander. Theme: tokens. Pool: any_card. Shortlist: 298 names.

Commander: Adeline, Resplendent Cathar.

Grade: bad, score 0.03, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $93.97 to buy, $93.97 the whole deck.

**Summary:** Adeline leads an attack-focused token deck that grows a Human army and turns a wide board into sustained creature pressure. Token synergies, supporting threats, and card draw keep the deck developing beyond its opening attacks, while removal and board resets help clear difficult positions. Combat is the main route to victory, backed by dedicated finishing pieces such as Halo Fountain and Sword of Body and Mind. The deck favors a dependable white mana base and gradual development over explosive acceleration, so it gives opponents more time to establish their own plans.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The deck makes less mana on turn four than the norm, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every card is a white token maker, anthem, Human or Knight attacker, or equipment, so the whole list supports Adeline's go-wide combat plan.
- PLAN theme_fit=yes: It is a mono-white token deck led by Adeline with many token makers, and it avoids expensive staples so it plausibly fits the budget and bracket 2 casual power.
- PLAN useful_as_built=partly: The deck has plenty of token makers, removal, and planeswalkers, but 39 lands including 35 Plains plus about ten low-impact rocks such as Currency Converter and Collector's Vault crowd out real threats, and the deck makes less mana on turn four than the norm.
- PLAN summary_honest=partly: The summary claims Halo Fountain and Sword of Body and Mind as dedicated finishers, and both are in the list, but Halo Fountain needs many untapped creatures and the Sword is only a modest closer; it also skips that the deck holds 39 lands and a lot of filler mana rocks, and it calls the mana base dependable while the fast-mana claim is downplayed.
- [INFO] `curve_summary`: average mana value 3.42 over 60 nonland cards

<details><summary>The deck list</summary>

- 35 Plains | land | Provides a dependable white mana base for Adeline and the supporting spells.
- 1 Castle Ardenvale | land | Adds a utility land without moving away from the white mana base.
- 1 Windbrisk Heights | land | Provides another utility land for the creature-heavy plan.
- 1 Command Tower | land | Supports consistent access to white mana.
- 1 Path of Ancestry | land | Adds a white source suited to the deck's many Humans.
- 1 Goldvein Pick | ramp | Provides inexpensive ramp for an attack-focused deck.
- 1 Currency Converter | ramp | Adds a low-cost resource-development piece.
- 1 Collector's Vault | ramp | Supports mana development in a compact artifact slot.
- 1 Bucknard's Everfull Purse | ramp | Adds another inexpensive ramp option.
- 1 Staff of Titania | ramp | Supports mana development alongside the combat plan.
- 1 Legion's Landing // Adanto, the First Fort | ramp | Adds an early ramp piece suited to a creature-heavy deck.
- 1 Keeper of the Accord | ramp | Provides creature-based ramp without abandoning board development.
- 1 The Restoration of Eiganjo // Architect of Restoration | ramp | Supports the mana base while fitting the deck's gradual development.
- 1 Monologue Tax | ramp | Adds a supplementary source of ramp.
- 1 Glaring Fleshraker | ramp | Provides another creature-based ramp option.
- 1 Staff of the Storyteller | draw | Adds an affordable draw engine to support the token plan.
- 1 Idol of Oblivion | draw | Provides a compact draw piece for the token-focused shell.
- 1 Wedding Announcement // Wedding Festivity | draw | Supports sustained development with an inexpensive draw slot.
- 1 Bygone Bishop | draw | Adds a draw-supporting creature to the board.
- 1 Faramir, Field Commander | draw | Provides card advantage while reinforcing the Human theme.
- 1 Glimmer Seeker | draw | Adds an inexpensive Human to the draw package.
- 1 Platoon Dispenser | draw | Supports card advantage at the upper end of the curve.
- 1 Sanctuary Warden | draw | Combines a draw slot with a designated finishing threat.
- 1 Court of Grace | draw | Adds an enchantment-based source of card advantage.
- 1 Wojek Investigator | draw | Provides another draw-supporting creature.
- 1 Rootborn Defenses | interaction | Supplies interaction for protecting the creature-based plan.
- 1 Spirit Bonds | interaction | Adds an inexpensive supporting interaction piece.
- 1 Lena, Selfless Champion | interaction | Provides interaction on a Human creature at the top of the curve.
- 1 Mage's Attendant | interaction | Adds creature-based interaction while developing the board.
- 1 Basri, Tomorrow's Champion | interaction | Supplies another affordable Human interaction slot.
- 1 Moogles' Valor | interaction | Adds instant-speed interaction for important combat turns.
- 1 Generous Gift | removal | Provides an affordable answer to an opposing problem permanent.
- 1 Stroke of Midnight | removal | Adds another removal spell for obstacles to the attack plan.
- 1 Skyclave Apparition | removal | Supplies removal while adding a creature to the board.
- 1 Citizen's Crowbar | removal | Adds a low-cost removal option in an artifact slot.
- 1 Hanged Executioner | removal | Provides another creature-based removal piece.
- 1 Ajani Fells the Godsire | removal | Combines a removal slot with a designated finishing piece.
- 1 Righteous Confluence | removal | Provides a higher-curve removal option for later turns.
- 1 Intangible Virtue | synergy | Adds inexpensive support for the token army.
- 1 Inspiring Leader | synergy | Supports the deck's commander-centered token strategy.
- 1 Divine Visitation | synergy | Adds a high-impact enchantment to the token synergy package.
- 1 Horn of Gondor | synergy | Reinforces the Human-focused token theme.
- 1 Hanweir Militia Captain // Westvale Cult Leader | synergy | Adds a low-cost Human synergy piece.
- 1 Clarion Spirit | synergy | Provides inexpensive support for building a wide board.
- 1 Felidar Retreat | synergy | Adds an enchantment-based payoff for sustained development.
- 1 Rosie Cotton of South Lane | synergy | Provides another affordable creature synergy piece.
- 1 Siege Veteran | synergy | Reinforces the Human Soldier side of the deck.
- 1 Oketra's Monument | synergy | Adds a supporting artifact for the white creature strategy.
- 1 Hero of Bladehold | threat | Adds a Human Knight threat to the combat plan.
- 1 God-Eternal Oketra | threat | Provides a substantial threat for the middle and late game.
- 1 Defiler of Faith | threat | Adds a larger Human threat to sustain pressure.
- 1 Basri's Lieutenant | threat | Provides an affordable Human Knight threat.
- 1 Emeria Angel | threat | Adds an Angel threat to diversify the creature board.
- 1 Warren Warleader | threat | Provides another threat suited to the attacking strategy.
- 1 Mite Overseer | threat | Adds an inexpensive Soldier threat.
- 1 Phantom General | threat | Reinforces the token-focused threat package.
- 1 Requiem Angel | threat | Provides a designated finisher at the upper end of the curve.
- 1 Nahiri, the Lithomancer | threat | Adds a planeswalker threat that is also a designated finisher.
- 1 Gideon, Ally of Zendikar | threat | Provides another planeswalker threat for sustained pressure.
- 1 Halo Fountain | wincon | Provides a designated win condition for the token-focused deck.
- 1 Luck Bobblehead | wincon | Adds a secondary designated win condition.
- 1 Sword of Body and Mind | wincon | Provides a designated finishing piece for the combat plan.
- 1 Hour of Reckoning | wipe | Adds a board reset for games that become crowded.
- 1 Martial Coup | wipe | Provides another affordable board-reset option.
- 1 Elspeth, Sun's Champion | wipe | Adds a planeswalker-based wipe for difficult late-game boards.

</details>

### 13. owned first, and the commander is not owned

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 230 names.

Commander: Karlov of the Ghost Council.

Grade: baseline, score 0.08, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $233.88 the whole deck.

**Summary:** Karlov leads a creature-focused lifegain deck that builds a threatening board while turning life gained into counters and creature removal. Develop your mana and supporting creatures early, protect Karlov, and use selective removal to keep attacks productive. Angels provide a second line of pressure, chief among them Lyra Dawnbringer, while the other finishers offer ways to close longer games. The deck favors steady development and protected combat over explosive combo turns, so it gives up speed and remains vulnerable when its supporting board is repeatedly cleared.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The land count sits above the norm of the format, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: Most of the cards support a lifegain creature board around Karlov, but a lot of the list is generic value, equipment and removal, and the three wraths (Fumigate, Austere Command, Dusk // Dawn) pull against the creature-board plan.
- PLAN theme_fit=partly: Karlov is not in the list as a card line and the colors are right, yet the deck is a generic white-black good-stuff pile with only some lifegain payoffs, and nothing shows it was built from the person's library first.
- PLAN useful_as_built=yes: The deck has plenty of lands and mana rocks, cheap draw, a wide spread of removal and several strong finishers, so it can be played as it stands even though the land count is high.
- PLAN summary_honest=partly: The summary says Angels are a second line of pressure led by Lyra Dawnbringer, and the list does carry Lyra, Angel of Invention, Shattered Angel, Victory's Herald and others, but it never mentions the three board wipes or the large amount of ramp and equipment.
- [WARN] `not_owned`: Karlov of the Ghost Council: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 2.94 over 63 nonland cards
- [INFO] `mana_pass`: the builder moved 1 card of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 1 Command Tower | land | Provides fixing for the deck's two colors.
- 1 City of Brass | land | Supports both white and black spells.
- 1 Exotic Orchard | land | Adds another source of multicolor fixing.
- 1 Marsh Flats | land | Helps find the basic color needed early.
- 1 Fabled Passage | land | Balances access to white and black mana.
- 1 Evolving Wilds | land | Provides basic-land fixing.
- 1 Terramorphic Expanse | land | Helps smooth hands missing one color.
- 1 Escape Tunnel | land | Adds another flexible route to basic lands.
- 1 Path of Ancestry | land | Supports the deck's two-color mana base.
- 1 Thriving Moor | land | Adds black mana with complementary fixing.
- 1 Grand Coliseum | land | Provides flexible colored mana.
- 1 Spire of Industry | land | Pairs with the artifact package to support colored mana.
- 1 Plaza of Heroes | land | Supports Karlov and the other legendary creatures.
- 14 Plains | land | Provides reliable white mana for the creature and protection packages.
- 9 Swamp | land | Provides reliable black mana for removal and supporting creatures.
- 1 Sol Ring | ramp | Accelerates deployment of the deck's supporting pieces.
- 1 Springleaf Drum | ramp | Adds early ramp alongside the creature-heavy plan.
- 1 Arcane Signet | ramp | Combines acceleration with two-color fixing.
- 1 Fellwar Stone | ramp | Adds inexpensive artifact ramp.
- 1 Thought Vessel | ramp | Helps build mana for threats and interaction.
- 1 Lotho, Corrupt Shirriff | ramp | Provides a creature-based ramp slot.
- 1 Giada, Font of Hope | ramp | Supports the deck's Angel threats.
- 1 Chromatic Lantern | ramp | Supports acceleration and consistent colored mana.
- 1 Commander's Sphere | ramp | Adds another source of colored artifact ramp.
- 1 Relic of Legends | ramp | Fits the deck's substantial legendary-creature package.
- 1 Call of the Ring | draw | Provides a continuing source of cards.
- 1 Exemplar of Light | draw | Places card draw in an Angel creature slot.
- 1 Inspiring Overseer | draw | Adds card flow while developing the creature board.
- 1 Lembas | draw | Provides card flow through a Food artifact.
- 1 Mask of Memory | draw | Supports card flow through the combat plan.
- 1 Night's Whisper | draw | Provides an efficient way to refill the hand.
- 1 Skullclamp | draw | Turns the small-creature package into additional card flow.
- 1 Stiltzkin, Moogle Merchant | draw | Adds another low-curve creature to the draw package.
- 1 Tome of Legends | draw | Supports card flow in a commander-centered strategy.
- 1 Wall of Omens | draw | Combines early board presence with card flow.
- 1 Buster Sword | draw | Adds draw support to the equipment and combat package.
- 1 Bastion Protector | interaction | Helps safeguard the commander-centered game plan.
- 1 Boromir, Warden of the Tower | interaction | Provides interaction on a legendary creature.
- 1 Cecil, Dark Knight // Cecil, Redeemed Paladin | interaction | Adds a creature-based interaction option.
- 1 Clever Concealment | interaction | Protects the board investment behind Karlov.
- 1 Lightning Greaves | interaction | Helps preserve a key creature for the combat plan.
- 1 Swiftfoot Boots | interaction | Adds another protective equipment option.
- 1 Sheltered by Ghosts | interaction | Adds interactive support to the creature-focused strategy.
- 1 Take Up the Shield | interaction | Provides inexpensive interaction for creature combat.
- 1 Swords to Plowshares | removal | Answers a creature that threatens the deck's position.
- 1 Get Lost | removal | Adds a flexible answer to opposing threats.
- 1 Generous Gift | removal | Broadens the removal package beyond creatures.
- 1 Stroke of Midnight | removal | Provides another flexible removal spell.
- 1 Bitter Triumph | removal | Adds a low-cost answer to an important threat.
- 1 Infernal Grasp | removal | Keeps opposing creatures from dominating combat.
- 1 Dismember | removal | Adds an efficient creature-removal option.
- 1 Orcish Bowmasters | removal | Places removal in a creature slot.
- 1 Banishing Light | removal | Adds enchantment-based removal for a troublesome permanent.
- 1 Angel of Invention | threat | Adds an Angel threat to the creature-based pressure plan.
- 1 Bill the Pony | threat | Provides another creature for building board presence.
- 1 Gollum, Silent Slinker // Meager Meal | threat | Adds a threat with an accompanying Adventure.
- 1 Minwu, White Mage | threat | Provides a legendary threat for the lifegain strategy.
- 1 Shattered Angel | threat | Adds another substantial Angel to the battlefield.
- 1 Victory's Herald | threat | Provides a top-end Angel for the combat plan.
- 1 Reaping Willow | threat | Adds a substantial creature for longer games.
- 1 Dawnhand Eulogist | threat | Adds a marked finisher to the threat package.
- 1 Sneering Shadewriter | threat | Provides another marked finisher outside the Angel package.
- 1 Aerith Gainsborough | synergy | Adds a legendary Cleric to the deck's synergy package.
- 1 Angel of Vitality | synergy | Supports the lifegain theme alongside the other Angels.
- 1 Denethor, Ruling Steward | synergy | Adds a legendary synergy piece and marked finisher.
- 1 Eastfarthing Farmer | synergy | Provides supporting creature synergy for the lifegain plan.
- 1 Graveyard Trespasser // Graveyard Glutton | synergy | Adds a synergy creature that also serves as a marked finisher.
- 1 Kor Firewalker | synergy | Adds another creature to the lifegain support package.
- 1 Light of Promise | synergy | Reinforces the deck's life-and-counters theme.
- 1 Rosie Cotton of South Lane | synergy | Adds a legendary support creature to the synergy package.
- 1 The Darkness Crystal | synergy | Adds a legendary artifact to support the deck's theme.
- 1 White Mage's Staff | synergy | Provides thematic equipment support for the creature plan.
- 1 Lyra Dawnbringer | wincon | Provides a marked finisher that fits the Angel package.
- 1 Frodo, Sauron's Bane | wincon | Adds a distinct legendary finishing threat.
- 1 Al Bhed Salvagers | wincon | Supplies another marked finisher for closing longer games.
- 1 Fumigate | wipe | Resets an opposing creature board when the deck falls behind.
- 1 Austere Command | wipe | Provides a flexible board-reset option.
- 1 Dusk // Dawn | wipe | Adds another reset suited to the creature-centered strategy.

</details>

### 14. the user delegates the commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 299 names.

Commander: Atarka, World Render.

Grade: bad, score 0.02, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $1264.99 to buy, $1264.99 the whole deck.

**Summary:** Develop mana early, build a board of Dragons, and use your commander's double-strike attacks to close the game. Card-flow engines keep threats coming, while inexpensive interaction and removal help preserve decisive combat turns. Lathliss, Dragon Queen, Hellkite Tyrant, and Utvara Hellkite provide additional finishing threats. The deck favors sustained tribal pressure over combo wins, giving up some speed and broad answers in exchange for a strong Dragon theme; its expensive creatures make rebuilding after repeated board clears challenging.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves one color far better than another, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core of Dragon creatures with ramp and card draw pulls together, but filler like Wheel of Fortune, Pyroblast, Red Elemental Blast, Guardian Project and Beast Whisperer serves no Dragon plan.
- PLAN theme_fit=yes: It is a Dragon tribal Commander deck led by a Dragon commander in red-green, with roughly two dozen Dragons and Dragon-support cards, which matches the request for a dragons deck at bracket 3.
- PLAN useful_as_built=partly: The deck has plenty of lands, ramp and Dragons to play, but it runs 10 basic Mountains against only 2 Forests, has several green-heavy cards like Beast Whisperer and Harmonize that are hard to cast, and includes dead-ish cards such as Pyroblast and Red Elemental Blast.
- PLAN summary_honest=no: The summary says the commander makes 'double-strike attacks' to close the game, but Atarka, World Render is not listed with double strike here, and the deck's real payoff is Dragon count and haste-style beats rather than that claim.
- [INFO] `curve_summary`: average mana value 3.17 over 63 nonland cards
- [INFO] `mana_pass`: the builder moved 4 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 1 Command Tower | land | Provides flexible fixing for both deck colors.
- 1 Taiga | land | Supports red and green without slowing early development.
- 1 Stomping Ground | land | Provides both colors for ramp and Dragons.
- 1 Spire Garden | land | Helps cast colored spells on schedule.
- 1 Karplusan Forest | land | Provides early access to either color.
- 1 Grove of the Burnwillows | land | Keeps the two-color mana base consistent.
- 1 Copperline Gorge | land | Supports the deck's early setup turns.
- 1 Rootbound Crag | land | Adds reliable red-green fixing.
- 1 Rockfall Vale | land | Supports colored mana through the middle turns.
- 1 Game Trail | land | Adds another red-green source.
- 1 Cinder Glade | land | Provides both colors alongside the basic lands.
- 1 Fire-Lit Thicket | land | Helps support demanding colored costs.
- 1 Mossfire Valley | land | Helps balance red and green mana.
- 1 Cragcrown Pathway // Timbercrown Pathway | land | Offers flexible color selection during development.
- 1 Thornspire Verge | land | Strengthens the deck's red-green fixing.
- 1 City of Brass | land | Provides flexible colored mana.
- 1 Mana Confluence | land | Prioritizes dependable fixing over life conservation.
- 1 Forbidden Orchard | land | Adds flexible mana for the deck's colored spells.
- 1 Exotic Orchard | land | Adds multiplayer-friendly color fixing.
- 1 Wooded Foothills | land | Helps find the land color needed for development.
- 1 Prismatic Vista | land | Improves access to the appropriate basic land.
- 1 Starting Town | land | Adds flexible fixing to the opening turns.
- 1 Cavern of Souls | land | Supports the Dragon-heavy creature suite.
- 1 Secluded Courtyard | land | Provides tribal fixing for Dragons.
- 2 Forest | land | Provides dependable green mana for setup and support spells.
- 10 Mountain | land | Provides dependable red mana for Dragons and removal.
- 1 Mox Jasper | ramp | Adds inexpensive acceleration to the Dragon plan.
- 1 Orb of Dragonkind | ramp | Helps bridge early development into expensive Dragons.
- 1 Pillar of Origins | ramp | Provides early acceleration for the creature-heavy plan.
- 1 Scaled Nurturer | ramp | Adds a low-cost ramp creature to the setup turns.
- 1 Sarkhan, Dragon Ascendant | ramp | Supports the transition from setup to Dragon deployment.
- 1 Magda, Brazen Outlaw | ramp | Adds an inexpensive ramp option before the larger threats.
- 1 Carnelian Orb of Dragonkind | ramp | Helps finance the deck's red-heavy threats.
- 1 Jade Orb of Dragonkind | ramp | Adds Dragon-focused mana development.
- 1 Dragon's Hoard | ramp | Supports sustained mana development for expensive threats.
- 1 Patchwork Banner | ramp | Adds mana support to the tribal game plan.
- 1 Faithless Looting | draw | Improves early card selection without a large mana investment.
- 1 Demand Answers | draw | Provides inexpensive access to fresh cards.
- 1 Thrill of Possibility | draw | Keeps cards moving during the setup turns.
- 1 Sylvan Library | draw | Supports consistent access to useful cards.
- 1 Garruk's Uprising | draw | Supplies card support for the large-creature plan.
- 1 Dragonborn Champion | draw | Pairs a Dragon body with the deck's card-flow needs.
- 1 Beast Whisperer | draw | Supports continued development in a creature-heavy deck.
- 1 Guardian Project | draw | Helps sustain a sequence of creature threats.
- 1 Harmonize | draw | Provides straightforward replenishment after setup.
- 1 Return of the Wildspeaker | draw | Adds substantial card support to the large-creature strategy.
- 1 Wheel of Fortune | draw | Provides a way to replenish resources after committing threats.
- 1 Heroic Intervention | interaction | Helps preserve the board investment needed for decisive attacks.
- 1 Legolas's Quick Reflexes | interaction | Adds inexpensive interaction around important combat turns.
- 1 Lightning Greaves | interaction | Supports keeping an important creature ready for the attack plan.
- 1 Swiftfoot Boots | interaction | Adds protection-oriented support for key creatures.
- 1 Veil of Summer | interaction | Provides inexpensive interaction when committing important spells.
- 1 Pyroblast | interaction | Adds a low-cost answer to opposing interference.
- 1 Red Elemental Blast | interaction | Provides another inexpensive interactive option.
- 1 Tibalt's Trickery | interaction | Adds a reactive option against a critical opposing play.
- 1 Dragon Tempest | removal | Connects the Dragon theme with the removal package.
- 1 Draconic Roar | removal | Provides inexpensive removal suited to a Dragon deck.
- 1 Dragon's Fire | removal | Adds another low-cost removal option.
- 1 Invasion of Tarkir // Defiant Thundermaw | removal | Adds removal while reinforcing the Dragon theme.
- 1 Molten Exhale | removal | Keeps the removal package inexpensive.
- 1 Spit Flame | removal | Provides removal support for repeated Dragon deployment.
- 1 Piercing Exhale | removal | Adds another answer to obstacles in the combat plan.
- 1 Glorybringer | removal | Combines a Dragon threat with the deck's removal needs.
- 1 Terror of the Peaks | removal | Supplies removal pressure and a substantial finishing threat.
- 1 Acolyte of Bahamut | synergy | Provides inexpensive support for the Dragon strategy.
- 1 Dragonlord's Servant | synergy | Supports deploying the deck's expensive Dragons.
- 1 Dragonspeaker Shaman | synergy | Helps the Dragon-heavy plan develop efficiently.
- 1 Urza's Incubator | synergy | Supports committing multiple tribal threats over a game.
- 1 Herald's Horn | synergy | Adds lasting support for the Dragon creature suite.
- 1 Shared Animosity | synergy | Reinforces the deck's commitment to attacking with a tribal board.
- 1 Realmwalker | synergy | Adds creature-based support to the tribal strategy.
- 1 Thunderbreak Regent | threat | Provides a Dragon threat before the most expensive finishers.
- 1 Stormbreath Dragon | threat | Adds a midgame attacker to maintain pressure.
- 1 Kura, the Boundless Sky | threat | Adds another substantial Dragon to the attack plan.
- 1 Manaform Hellkite | threat | Provides a relatively accessible Dragon threat.
- 1 Mirrorwing Dragon | threat | Broadens the midgame Dragon threat suite.
- 1 Vengeful Ancestor | threat | Adds a Dragon body without extending the top of the curve.
- 1 Verix Bladewing | threat | Provides another midgame Dragon for sustained pressure.
- 1 Thrakkus the Butcher | threat | Reinforces the deck's Dragon combat plan.
- 1 Roaming Throne | threat | Adds a substantial creature to the tribal threat suite.
- 1 Territorial Hellkite | threat | Provides a finishing-capable Dragon at a moderate cost.
- 1 Hellkite Courser | threat | Adds a larger Dragon for the deck's decisive turns.
- 1 Twinflame Tyrant | threat | Provides another finishing-capable Dragon threat.
- 1 Lathliss, Dragon Queen | wincon | Provides a thematic finisher for the Dragon-heavy board.
- 1 Hellkite Tyrant | wincon | Adds a distinct finishing threat alongside the combat plan.
- 1 Utvara Hellkite | wincon | Provides a high-end finisher for a developed Dragon board.
- 1 Breath Weapon | wipe | Adds an inexpensive way to manage a crowded battlefield.
- 1 Thundermaw Hellkite | wipe | Pairs battlefield management with a finishing-capable Dragon.
- 1 Harbinger of the Hunt | wipe | Adds another Dragon-based option for managing opposing boards.

</details>

### 15. delegated commander, owned first

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 230 names.

Commander: Astarion, the Decadent.

Grade: baseline, score 0.05, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $153.70 the whole deck.

**Summary:** Astarion, the Decadent leads a patient lifegain deck that develops mana, builds a creature board, and keeps protection and removal available for important turns. His end-step choices turn life gained into a larger cushion or amplify an opponent's life loss. Angels provide much of the combat pressure, chief among them Lyra Dawnbringer, while Frodo, Sauron's Bane and Grave Venerations offer additional finishing routes. The deck favors sustained board presence and longer games over explosive starts, so it gives up speed and depends on keeping its creatures relevant.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The land count sits above the norm of the format, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The cards mostly point at a white-heavy lifegain creature deck with Angels, protection and removal, but two board wipes plus a third in Dusk // Dawn work against a creature-based plan, and much of the ramp and draw is generic filler that doesn't advance lifegain.
- PLAN theme_fit=yes: It is a white-black lifegain Commander deck led by Astarion, the Decadent, whose end-step life choices fit the theme, and it carries many lifegain payoffs and Angels at a modest power level suited to bracket 2, though the request's 'from my library first' cannot be verified from the list.
- PLAN useful_as_built=partly: The deck has plenty of ramp, draw, removal and finishers, but 30 basics plus 7 nonbasic lands is a very high land count next to 10 mana rocks and a mana-producing creature, so flooding is likely, and few cards actually gain life or use it as a payoff.
- PLAN summary_honest=partly: The claim that Angels provide combat pressure, with Lyra Dawnbringer chief among them, checks out against the list (Angel of Invention, Shattered Angel, Victory's Herald, Exemplar of Light, and others), but the summary never mentions the three board wipes and calls the deck patient without noting how many generic artifacts it runs.
- [WARN] `not_owned`: Astarion, the Decadent: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.11 over 62 nonland cards
- [INFO] `mana_pass`: the builder moved 3 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 18 Plains | land | Provides a dependable foundation for the deck's white spells.
- 12 Swamp | land | Provides a dependable foundation for the deck's black spells.
- 1 Command Tower | land | Helps supply both colors for the commander and supporting spells.
- 1 City of Brass | land | Adds flexible access to both colors.
- 1 Exotic Orchard | land | Adds another flexible colored-mana source.
- 1 Grand Coliseum | land | Supports the deck's two-color mana requirements.
- 1 Path of Ancestry | land | Provides fixing for the creature-heavy plan.
- 1 Thriving Moor | land | Helps balance access to white and black mana.
- 1 Spire of Industry | land | Adds fixing alongside the deck's artifact package.
- 1 Sol Ring | ramp | Accelerates the deck toward its commander and larger threats.
- 1 Arcane Signet | ramp | Combines early acceleration with colored-mana support.
- 1 Fellwar Stone | ramp | Adds inexpensive acceleration and potential fixing.
- 1 Thought Vessel | ramp | Provides another inexpensive mana rock.
- 1 Chromatic Lantern | ramp | Supports both acceleration and reliable colored mana.
- 1 Commander's Sphere | ramp | Adds colored mana for the deck's middle turns.
- 1 Relic of Legends | ramp | Supports a deck with several legendary creatures.
- 1 Inherited Envelope | ramp | Adds another source of colored acceleration.
- 1 Lotho, Corrupt Shirriff | ramp | Provides creature-based acceleration within the deck's colors.
- 1 Bender's Waterskin | ramp | Rounds out the mana-development package.
- 1 Call of the Ring | draw | Supports a steady supply of cards.
- 1 Exemplar of Light | draw | Adds card support on an Angel body.
- 1 Inspiring Overseer | draw | Pairs card access with another Angel creature.
- 1 Lembas | draw | Adds card support through a Food artifact.
- 1 Mask of Memory | draw | Provides card support for the creature-combat plan.
- 1 Night's Whisper | draw | Supplies straightforward card replenishment.
- 1 Skullclamp | draw | Adds an efficient draw tool to the creature-heavy deck.
- 1 Wall of Omens | draw | Provides card support on a defensive creature.
- 1 Tome of Legends | draw | Supports card access alongside the legendary-creature package.
- 1 Stiltzkin, Moogle Merchant | draw | Adds another creature-based source of card support.
- 1 Buster Sword | draw | Adds draw support through combat-oriented Equipment.
- 1 Boromir, Warden of the Tower | interaction | Adds interaction on a creature body.
- 1 Lightning Greaves | interaction | Helps safeguard the commander and important creatures.
- 1 Swiftfoot Boots | interaction | Provides another protective Equipment option.
- 1 Clever Concealment | interaction | Helps preserve the deck's developed board.
- 1 Sheltered by Ghosts | interaction | Adds interactive support to the creature plan.
- 1 Take Up the Shield | interaction | Adds a compact defensive combat option.
- 1 Swords to Plowshares | removal | Provides inexpensive creature removal.
- 1 Generous Gift | removal | Adds flexible removal for opposing problems.
- 1 Get Lost | removal | Provides another inexpensive removal option.
- 1 Infernal Grasp | removal | Adds direct creature removal.
- 1 Bitter Triumph | removal | Broadens the instant-speed removal package.
- 1 Stroke of Midnight | removal | Adds another flexible answer.
- 1 Witch-king of Angmar | removal | Combines the removal package with a marked finisher.
- 1 Aerith Gainsborough | synergy | Adds a legendary creature to the deck's synergy package.
- 1 Angel of Vitality | synergy | Connects the synergy package with the Angel theme.
- 1 Light of Promise | synergy | Adds an Aura payoff to the lifegain-focused plan.
- 1 Rosie Cotton of South Lane | synergy | Adds another legendary support creature.
- 1 Denethor, Ruling Steward | synergy | Provides thematic creature support and another marked finisher.
- 1 Graveyard Trespasser // Graveyard Glutton | synergy | Adds a creature-based synergy piece with finishing potential.
- 1 Eastfarthing Farmer | synergy | Adds a supporting creature to the synergy package.
- 1 The Darkness Crystal | synergy | Provides artifact-based support for the deck's central plan.
- 1 White Mage's Staff | synergy | Adds supportive Equipment to the creature-focused strategy.
- 1 Elixir | synergy | Adds another artifact to the lifegain-focused support package.
- 1 Compassionate Healer | synergy | Provides another small support creature for the central theme.
- 1 Angel of Invention | threat | Adds an Angel threat to the deck's battlefield pressure.
- 1 Bill the Pony | threat | Adds a legendary creature to the threat package.
- 1 Dawnhand Eulogist | threat | Provides another threat with marked finishing potential.
- 1 Gollum, Silent Slinker // Meager Meal | threat | Adds a threat with an Adventure component.
- 1 Minwu, White Mage | threat | Adds a legendary Cleric threat to the creature plan.
- 1 Rabaroo Troop | threat | Adds another creature for sustained battlefield pressure.
- 1 Reaping Willow | threat | Provides a Treefolk Cleric threat for the middle and later turns.
- 1 Rooftop Percher | threat | Diversifies the creature threat package.
- 1 Shattered Angel | threat | Adds another Angel to the deck's threat suite.
- 1 Sneering Shadewriter | threat | Adds a Vampire threat with marked finishing potential.
- 1 Victory's Herald | threat | Provides another substantial Angel threat.
- 1 Lyra Dawnbringer | wincon | Supplies an Angel finisher for the creature-focused plan.
- 1 Frodo, Sauron's Bane | wincon | Provides a dedicated creature-based route to victory.
- 1 Grave Venerations | wincon | Adds an enchantment-based finisher alongside the creatures.
- 1 Fumigate | wipe | Provides a reset when opposing creature boards become overwhelming.
- 1 Austere Command | wipe | Adds a flexible battlefield reset.
- 1 Dusk // Dawn | wipe | Rounds out the deck's board-reset options.

</details>

### 16. a tight budget, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 231 names.

Commander: Denethor, Ruling Steward.

Grade: baseline, score 0.32, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $24.93 to buy, $141.07 the whole deck.

**Summary:** Build a board of expendable creatures, use Denethor to turn sacrifices into steady life loss, and refill with draw engines. Aristocrats finishers, chief among them Bastion of Remembrance and Zulaport Cutthroat, close out the game while larger creatures provide a combat backup. Protection and rebuilding tools support patient attrition, but the deck gives up explosive finishes and can lose momentum when repeated board clears disrupt its creature engines.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every card feeds the aristocrats plan: cheap sacrifice fodder, Skullclamp and Blood Artist style drain payoffs, sacrifice outlets, and a supporting suite of ramp, removal, and draw.
- PLAN theme_fit=partly: It is a white-black aristocrats deck led by Denethor at a bracket-2 power level, but the 25-dollar budget is doubtful with Sol Ring, Skullclamp, Lightning Greaves, Swiftfoot Boots, Marsh Flats and other pricey cards, and nothing shows it was built from the person's own library first.
- PLAN useful_as_built=partly: It has plenty of ramp, draw, removal, and drain payoffs so it plays as built, but the mana base of 27 basics with 10 nonbasics plus 10 mana rocks is heavy, and the model found turn-four mana below the norm and few cards with strong performance.
- PLAN summary_honest=yes: The summary says Bastion of Remembrance and Zulaport Cutthroat close out the game, and both are in the list as wincons alongside Denethor and the draw engines, while the wipes it warns about (Austere Command, Martial Coup, Dusk // Dawn) are also present.
- [WARN] `not_owned`: Erebos, Bleak-Hearted: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Disciple of Bolas: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Smothering Abomination: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Skullport Merchant: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bartolomé del Presidio: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Woe Strider: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Yahenni, Undying Partisan: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Baron Bertram Graywater: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ayli, Eternal Pilgrim: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Aron, Benalia's Ruin: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Lord Skitter's Butcher: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Shilgengar, Sire of Famine: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Sivriss, Nightmare Speaker: the deck needs 1, the collection has 0
- [WARN] `not_owned`: High-Society Hunter: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Falkenrath Noble: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vindictive Vampire: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vengeful Bloodwitch: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Elas il-Kor, Sadistic Pilgrim: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Blood Artist: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Elder Arthur Maxson: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ghoulcaller Gisa: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Umbral Collar Zealot: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Old Flitterfang: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Teysa, Orzhov Scion: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bastion of Remembrance: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Zulaport Cutthroat: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 2.97 over 62 nonland cards
- [INFO] `mana_pass`: the builder moved 3 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 1 Skullclamp | draw | Turns expendable creatures into cards for the attrition plan.
- 1 Idol of Oblivion | draw | Provides repeatable draw alongside the token plan.
- 1 Nasty End | draw | Combines card draw with a useful sacrifice.
- 1 Night's Whisper | draw | Provides inexpensive card draw without needing a developed board.
- 1 Call of the Ring | draw | Supplies ongoing cards for a longer game.
- 1 Wall of Omens | draw | Adds a defensive body while replacing itself.
- 1 Inspiring Overseer | draw | Adds a useful creature without exhausting the hand.
- 1 Erebos, Bleak-Hearted | draw | Supports card advantage in a creature-sacrifice strategy.
- 1 Disciple of Bolas | draw | Converts a creature into resources for rebuilding.
- 1 Smothering Abomination | draw | Rewards the deck's recurring sacrifices with cards.
- 1 Skullport Merchant | draw | Provides a creature-based way to turn spare resources into cards.
- 1 Clever Concealment | interaction | Protects an established board from disruptive plays.
- 1 Boromir, Warden of the Tower | interaction | Adds protection on a creature that fits the deck's sacrifice plan.
- 1 Swiftfoot Boots | interaction | Protects Denethor or an important engine creature.
- 1 Lightning Greaves | interaction | Helps keep a key creature available through opposing interaction.
- 1 Gift of Immortality | interaction | Helps preserve an important creature through removal and sacrifices.
- 1 Reprieve | interaction | Buys time against a spell that would disrupt the board.
- 1 Sol Ring | ramp | Accelerates development and supports Denethor's activated ability.
- 1 Arcane Signet | ramp | Provides early acceleration and both-color fixing.
- 1 Fellwar Stone | ramp | Adds inexpensive acceleration with potential color fixing.
- 1 Thought Vessel | ramp | Provides early mana for deploying engines.
- 1 Springleaf Drum | ramp | Uses spare creatures to support early mana development.
- 1 Chromatic Lantern | ramp | Smooths the mana base for spells in both colors.
- 1 Commander's Sphere | ramp | Provides fixing and remains useful when additional mana is unnecessary.
- 1 Relic of Legends | ramp | Supports acceleration alongside the deck's legendary creatures.
- 1 Inherited Envelope | ramp | Adds another source of dependable colored mana.
- 1 Lotho, Corrupt Shirriff | ramp | Adds creature-based acceleration to support longer turns.
- 1 Swords to Plowshares | removal | Answers a dangerous creature efficiently.
- 1 Bitter Triumph | removal | Provides flexible removal at a low mana cost.
- 1 Generous Gift | removal | Answers troublesome permanents beyond creatures.
- 1 Infernal Grasp | removal | Provides straightforward instant-speed creature removal.
- 1 Destroy Evil | removal | Covers threatening creatures and problematic enchantments.
- 1 Palace Jailer | removal | Combines creature removal with a body for the board.
- 1 Witch-king of Angmar | removal | Adds a substantial creature with a removal-oriented role.
- 1 Bartolomé del Presidio | synergy | Provides a low-cost sacrifice outlet.
- 1 Woe Strider | synergy | Adds a sacrifice outlet that fits the expendable-creature plan.
- 1 Yahenni, Undying Partisan | synergy | Provides another sacrifice outlet on a resilient creature.
- 1 Gollum, Patient Plotter | synergy | Supports recurring creature resources for the sacrifice plan.
- 1 Baron Bertram Graywater | synergy | Supports the token-and-sacrifice engine.
- 1 Ayli, Eternal Pilgrim | synergy | Adds another useful way to sacrifice expendable creatures.
- 1 Aron, Benalia's Ruin | synergy | Connects creature sacrifices with stronger board pressure.
- 1 Lord Skitter's Butcher | synergy | Provides flexible support for a creature-heavy sacrifice strategy.
- 1 Shilgengar, Sire of Famine | synergy | Supports sacrificing creatures and rebuilding creature resources.
- 1 Sivriss, Nightmare Speaker | synergy | Adds a repeatable use for expendable creatures.
- 1 Bill the Pony | threat | Adds an owned creature that supports the deck's resource-based plan.
- 1 High-Society Hunter | threat | Provides a substantial threat suited to a sacrifice-heavy board.
- 1 Falkenrath Noble | threat | Adds a finishing creature to the aristocrats plan.
- 1 Vindictive Vampire | threat | Adds another creature-based finishing payoff.
- 1 Vengeful Bloodwitch | threat | Provides an inexpensive payoff creature for the deck's main plan.
- 1 Elas il-Kor, Sadistic Pilgrim | threat | Adds a low-cost aristocrats payoff to the board.
- 1 Blood Artist | threat | Provides a central finishing payoff for creature attrition.
- 1 Elder Arthur Maxson | threat | Adds creature-based pressure alongside the token strategy.
- 1 Ghoulcaller Gisa | threat | Builds a threatening creature board from sacrifice resources.
- 1 Umbral Collar Zealot | threat | Adds another creature to support the sacrifice-and-pressure plan.
- 1 Old Flitterfang | threat | Provides a larger creature for the deck's resource-oriented game.
- 1 Teysa, Orzhov Scion | threat | Connects the creature-token plan with additional board control.
- 1 Bastion of Remembrance | wincon | Provides a dedicated aristocrats finisher.
- 1 Zulaport Cutthroat | wincon | Provides a compact finishing payoff for the sacrifice plan.
- 1 Grave Venerations | wincon | Adds an owned finishing option to close longer games.
- 1 Austere Command | wipe | Provides a flexible reset when opposing boards become stronger.
- 1 Dusk // Dawn | wipe | Offers a board reset and support for rebuilding smaller creatures.
- 1 Martial Coup | wipe | Combines a late-game reset with a fresh token board.
- 1 Command Tower | land | Provides untapped fixing for both deck colors.
- 1 City of Brass | land | Provides untapped access to either color.
- 1 Exotic Orchard | land | Adds flexible multiplayer color fixing.
- 1 Grand Coliseum | land | Adds another land with access to both colors.
- 1 Path of Ancestry | land | Provides both colors for the creature-focused deck.
- 1 Thriving Moor | land | Supports black mana while adding white fixing.
- 1 Spire of Industry | land | Adds color fixing alongside the deck's artifacts.
- 1 Marsh Flats | land | Finds the basic color needed for the opening hand.
- 1 Fabled Passage | land | Finds either basic color to smooth mana development.
- 1 The Grey Havens | land | Adds fixing supported by the deck's legendary creatures.
- 16 Plains | land | Provides dependable untapped white mana.
- 11 Swamp | land | Provides dependable untapped black mana.

</details>

### 17. upgrade a precon, owned first

Format: Commander. Theme: goblins. Pool: owned_first. Shortlist: 205 names.

Commander: Zada, Hedron Grinder.

Grade: typical, score 0.58, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $33.89 to buy, $233.99 the whole deck.

**Summary:** Zada, Hedron Grinder leads a Goblin-heavy spell deck that builds a broad creature board, develops its mana, and uses draw spells to keep its turns moving. Aim to establish your creatures before committing to a big spell sequence, then turn that momentum into a finishing push, with Great Train Heist and Assault on Osgiliath chief among the finishers. The deck preserves its Goblin Storm character and supporting removal, but favors creature-and-spell synergy over standalone threats, leaving it more dependent on keeping a board together.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every card is a red Goblin creature, token maker, ritual, or cheap draw spell, and they all feed one Goblin storm plan with Impact Tremors, Goblin Bombardment, Grapeshot, and Past in Flames as payoffs.
- PLAN theme_fit=yes: It is a mono-red Goblin storm Commander deck led by Zada, Hedron Grinder, which matches the request to upgrade a Goblin Storm precon at bracket 3, with Sol Ring, Skullclamp, and rituals as upgrades.
- PLAN useful_as_built=yes: With 37 lands, 24 of them Mountains, plus many rituals and cheap Goblins, the deck casts its spells on curve and has several ways to win, so it plays as it stands.
- PLAN summary_honest=partly: The summary names Great Train Heist and Assault on Osgiliath as chief finishers, and both are in the list, but they are marginal ones compared with Grapeshot, Goblin Bombardment, Krenko, and Impact Tremors, which do most of the finishing.
- [INFO] `curve_summary`: average mana value 2.40 over 62 nonland cards
- [WARN] `profile_off_band`: the interaction count is 1, and bracket 3 wants 4 to 12
- [INFO] `bracket_cut`: to hold bracket 3, the builder cut 1 card: Storm-Kiln Artist (the combo Storm-Kiln Artist + Haze of Rage), and added 1 basic land

<details><summary>The deck list</summary>

- 24 Mountain | land | Provides the deck's main supply of red mana.
- 1 Arena of Glory | land | Retains a land from the precon's mana base.
- 1 Castle Embereth | land | Preserves the precon's land package.
- 1 Den of the Bugbear | land | Keeps the Goblin-themed mana base intact.
- 1 Dwarven Mine | land | Retains a Mountain land from the precon.
- 1 Forgotten Cave | land | Preserves an existing land slot.
- 1 Fountainport | land | Keeps part of the precon's utility-land package.
- 1 Goblin Burrows | land | Retains a thematic land from the original deck.
- 1 Hidden Volcano | land | Preserves an existing land slot.
- 1 Kher Keep | land | Keeps part of the precon's utility-land package.
- 1 Shinka, the Bloodsoaked Keep | land | Retains a legendary land from the original mana base.
- 1 Smoldering Crater | land | Preserves an existing land slot.
- 1 Spinerock Knoll | land | Keeps the precon's land package intact.
- 1 War Room | land | Retains an existing utility land.
- 1 Arcane Signet | ramp | Adds ramp to help establish mana before the deck's bigger turns.
- 1 Battle Hymn | ramp | Preserves ramp for the spell-heavy game plan.
- 1 Brightstone Ritual | ramp | Retains ramp in the Goblin-focused core.
- 1 Glimpse the Impossible | ramp | Keeps the precon's supporting ramp package.
- 1 Howlsquad Heavy | ramp | Combines a Goblin creature slot with the ramp package.
- 1 Impulsive Pilferer | ramp | Retains a Goblin in the deck's ramp package.
- 1 Mana Geyser | ramp | Supports the mana demands of a large spell turn.
- 1 Ruby Medallion | ramp | Preserves the precon's mana-support package.
- 1 Seething Song | ramp | Retains ramp for committing multiple spells in a turn.
- 1 Skirk Prospector | ramp | Keeps a Goblin creature in the ramp package.
- 1 Sol Ring | ramp | Provides an established ramp piece from the precon.
- 1 Thought Vessel | ramp | Adds another ramp piece for developing the deck's mana.
- 1 Ancestral Anger | draw | Adds another sorcery to the draw package.
- 1 Crimson Wisps | draw | Retains instant-speed draw in the spell package.
- 1 Expedite | draw | Preserves an instant in the draw package.
- 1 Faithless Looting | draw | Keeps the precon's draw support.
- 1 Fists of Flame | draw | Retains a draw spell in the Zada-focused core.
- 1 Idol of Oblivion | draw | Provides an artifact slot in the draw package.
- 1 Renegade Tactics | draw | Preserves sorcery-based draw support.
- 1 Sazacap's Brew | draw | Keeps another instant in the draw package.
- 1 Skullclamp | draw | Retains equipment-based draw support.
- 1 Wild Ride | draw | Preserves the precon's supporting draw package.
- 1 Witch's Mark | draw | Keeps sorcery-based draw in the original core.
- 1 Boggart Shenanigans | removal | Retains removal in a Goblin-themed enchantment slot.
- 1 Broadside Bombardiers | removal | Combines a Goblin body with the removal package.
- 1 Chaos Warp | removal | Preserves an instant in the removal package.
- 1 Gempalm Incinerator | removal | Keeps removal within the Goblin creature package.
- 1 Goblin Bombardment | removal | Retains an enchantment in the removal package.
- 1 Goblin Negotiation | removal | Preserves a sorcery in the removal package.
- 1 Goblin Trashmaster | removal | Keeps a Goblin creature in the removal package.
- 1 Grapeshot | removal | Retains removal within the storm-themed spell core.
- 1 Pashalik Mons | removal | Preserves a legendary Goblin in the removal package.
- 1 Siege-Gang Commander | removal | Keeps a Goblin creature in the removal package.
- 1 Siege-Gang Lieutenant | removal | Retains another Goblin-based removal slot.
- 1 Blasphemous Act | wipe | Preserves a board-clearing option for difficult positions.
- 1 Vandalblast | wipe | Retains another wipe from the precon.
- 1 Krenko, Mob Boss | threat | Keeps a central Goblin threat in the original creature core.
- 1 Redcap Gutter-Dweller | threat | Preserves a substantial Goblin threat.
- 1 Roaming Throne | threat | Retains an artifact creature in the threat package.
- 1 Swiftfoot Boots | interaction | Keeps the precon's equipment-based interaction support.
- 1 Ancestors' Aid | synergy | Preserves an instant in the supporting spell package.
- 1 Conspicuous Snoop | synergy | Keeps a Goblin synergy piece from the precon.
- 1 Dragon Fodder | synergy | Retains a sorcery in the Goblin synergy package.
- 1 Empty the Warrens | synergy | Preserves the precon's storm-themed synergy core.
- 1 Frontline Heroism | synergy | Keeps an enchantment in the supporting synergy package.
- 1 General Kreat, the Boltbringer | synergy | Retains a legendary Goblin synergy piece.
- 1 Goblin Bushwhacker | synergy | Preserves the original Goblin synergy package.
- 1 Goblin Chieftain | synergy | Keeps a Goblin synergy piece in the creature core.
- 1 Goblin Lackey | synergy | Retains a Goblin synergy slot from the precon.
- 1 Goblin Matron | synergy | Preserves a Goblin synergy piece in the original core.
- 1 Goblin Warchief | synergy | Keeps the Goblin synergy package intact.
- 1 Grenzo, Havoc Raiser | synergy | Retains a legendary Goblin synergy slot.
- 1 Haze of Rage | synergy | Preserves a sorcery in the storm-themed support package.
- 1 Impact Tremors | synergy | Keeps an enchantment in the creature-focused synergy package.
- 1 Krenko's Command | synergy | Retains sorcery support for the Goblin plan.
- 1 Mogg War Marshal | synergy | Preserves a Goblin creature in the synergy package.
- 1 Past in Flames | synergy | Keeps the precon's spell-focused synergy package.
- 1 Quest for the Goblin Lord | synergy | Retains a Goblin-themed enchantment in the synergy package.
- 1 Rundvelt Hordemaster | synergy | Keeps a Goblin synergy piece in the original creature core.
- 1 Searslicer Goblin | synergy | Preserves another Goblin synergy slot.
- 1 Assault on Osgiliath | wincon | Adds a dedicated finisher to turn a developed position into a win.
- 1 Great Train Heist | wincon | Retains the precon's dedicated instant finisher.

</details>

### 18. upgrade a precon, any card

Format: Commander. Theme: turtles. Pool: any_card. Shortlist: 365 names.

Commander: Heroes in a Half Shell.

Grade: baseline, score 0.06, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $2668.45 to buy, $2668.45 the whole deck.

**Summary:** Heroes in a Half Shell leads a combat-focused Turtle, Ninja, and Mutant deck that turns successful attacks into growing threats and a steady flow of cards. Develop your mana, build a broad creature board, and keep attacking, with Shared Animosity supporting the tribal plan and Raphael, the Muscle among the thematic finishers. The deck preserves the precon's character and supporting packages rather than pursuing a combo finish; its main tradeoff is dependence on an established attacking board, making repeated removal and battlefield resets difficult to recover from.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The land count sits above the norm of the format, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The deck leans toward a Turtle, Mutant, and Ninja attack plan with plenty of ramp, but a five-color spread of mostly one-of legends, a few wipes that fight the creature board, and loose pieces like Fast Forward and Lessons from Life pull it in different directions.
- PLAN theme_fit=yes: It is built around Heroes in a Half Shell with the Turtle Power precon's Mutant Ninja Turtle characters and Universes Beyond cards, in the commander's five colors, and the added staples like Sol Ring and dual lands fit an upgrade at bracket 3.
- PLAN useful_as_built=partly: It has plenty of ramp, fixing, and creatures at every point of the curve so it can be played, but the very high land count plus mana rocks leaves fewer real threats and card flow, and expensive cards like Vanquish the Horde and Leatherhead clog the hand while few cards actually win the game.
- PLAN summary_honest=partly: The summary says Shared Animosity supports a tribal plan and Raphael is a finisher, and both are in the list, but Shared Animosity only rewards sharing a creature type and the list has many singleton types, and the summary does not mention that the deck packs three board wipes (Blasphemous Act, Wave Goodbye, Vanquish the Horde) that work against its own attacking board, though it does admit the vulnerability to resets.
- [INFO] `curve_summary`: average mana value 3.37 over 60 nonland cards
- [WARN] `profile_off_band`: the land count is 39, and bracket 3 wants 34 to 38
- [WARN] `profile_off_band`: the ramp count is 7, and bracket 3 wants 8 to 13
- [WARN] `profile_off_band`: the draw count is 1, and bracket 3 wants 8 to 14
- [WARN] `profile_off_band`: the removal count is 3, and bracket 3 wants 6 to 12

<details><summary>The deck list</summary>

- 1 Acidic Slime | removal | Keeps a creature-based utility slot in the precon's board-focused plan.
- 1 April O'Neil, Live on the Scene | draw | Preserves a thematic source of card advantage.
- 1 Arcade Cabinet | synergy | Retains the precon's themed artifact package.
- 1 Arcane Signet | ramp | Provides early acceleration for the five-color plan.
- 1 Assassin's Trophy | removal | Preserves an efficient answer slot.
- 1 Baxter, Fly in the Ointment | synergy | Adds another Mutant to the commander's combat-focused creature package.
- 1 Bebop, Skull & Crossbones | threat | Preserves a thematic Mutant threat.
- 1 Big Mother Mouser | threat | Keeps the precon's Robot creature package intact.
- 1 Biogenic Ooze | threat | Retains another creature threat alongside the themed roster.
- 1 Blasphemous Act | wipe | Keeps a reset available when opposing boards become overwhelming.
- 1 Casey Jones, Back Alley Brute | threat | Preserves a thematic creature in the combat plan.
- 1 Chromatic Lantern | ramp | Supports the deck's five-color mana requirements.
- 1 Coin of Mastery | synergy | Preserves the precon's themed artifact support.
- 1 Continue? | interaction | Keeps a thematic instant in the interaction package.
- 1 Corpsejack Menace | synergy | Supports the commander's counter-focused combat plan.
- 1 Cultivate | ramp | Preserves land-based acceleration and access to basic lands.
- 1 Dimension X Pizzasaur | wincon | Provides a thematic Mutant finisher.
- 1 Donatello, the Brains | synergy | Keeps a core Turtle in the commander's creature package.
- 1 Double Jump // Flying Kick | interaction | Preserves a thematic instant slot for the combat-focused plan.
- 1 Fellwar Stone | ramp | Adds inexpensive acceleration for the multicolor plan.
- 1 Endless Foot Assault | synergy | Keeps the precon's themed enchantment package intact.
- 1 Everything Pizza | wincon | Preserves another thematic finishing piece.
- 1 Exploding Barrel | other | Retains a themed artifact utility slot.
- 1 Fast Forward | other | Preserves a thematic sorcery from the precon.
- 1 Foot Chopper | synergy | Keeps Equipment support for the creature-heavy plan.
- 1 Game Over | other | Retains a themed sorcery in the precon's support package.
- 1 Birds of Paradise | ramp | Adds early acceleration and support for five-color casting.
- 1 Here Comes a New Hero! | other | Keeps a thematic sorcery supporting the precon's overall plan.
- 1 High Score | synergy | Preserves the precon's themed enchantment support.
- 1 Irma, Part-Time Mutant | synergy | Adds another Mutant to the commander's creature package.
- 1 Krang, the All-Powerful | threat | Keeps a thematic artifact creature threat.
- 1 Leatherhead, Iron Gator | threat | Preserves a Mutant threat for the combat plan.
- 1 Leonardo, the Balance | threat | Keeps a core Turtle threat in the deck.
- 1 Lessons from Life | other | Retains a thematic sorcery from the precon's support package.
- 1 Shared Animosity | synergy | Adds focused support for the deck's tribal combat plan.
- 1 Lita, Little Orphan Amphibian | synergy | Preserves another Turtle for the commander's creature package.
- 1 Michelangelo, the Heart | synergy | Keeps a core Turtle in the precon's thematic engine.
- 1 Mole Module | synergy | Preserves the precon's Vehicle support.
- 1 Mona Lisa, Science Geek | synergy | Adds another thematic Mutant to the combat roster.
- 1 Ninja Pizza | synergy | Retains the precon's themed enchantment package.
- 1 Raphael, the Muscle | threat | Provides a core Turtle threat and combat finisher.
- 1 Rat King, Pale Piper | threat | Keeps another thematic creature threat.
- 1 Ray Fillet, Wave Warrior | threat | Preserves a Mutant threat in the creature package.
- 1 Roadkill Rodney | threat | Retains a themed Robot alongside the precon's artifact creatures.
- 1 Rocksteady, Mutant Marauder | threat | Keeps another thematic Mutant threat.
- 1 Shellshock | interaction | Preserves a thematic instant in the support package.
- 1 Shredder, Shadow Master | threat | Adds a Ninja threat to the commander's combat roster.
- 1 Sol Ring | ramp | Preserves efficient acceleration for developing the board.
- 1 Special Move | interaction | Retains a thematic instant for the deck's interactive turns.
- 1 Splinter, the Mentor | synergy | Keeps a Mutant Ninja in the commander's creature package.
- 1 Steelbane Hydra | removal | Preserves a Turtle-based removal slot.
- 1 Super Combo | other | Keeps a thematic sorcery in the precon's support package.
- 1 Swift Demise | interaction | Preserves an instant-speed support slot from the precon.
- 1 Tempestra, Dame of Games | threat | Retains another thematic creature threat.
- 1 Together Forever | synergy | Supports the deck's counter-focused creature plan.
- 1 Tokka & Rahzar, Unsupervised | ramp | Preserves thematic acceleration on a Turtle Mutant creature.
- 1 Vanquish the Horde | wipe | Keeps another way to reset an unfavorable battlefield.
- 1 Vigor | synergy | Supports the deck's creature-heavy combat plan.
- 1 Voracious Hydra | threat | Retains a substantial creature threat.
- 1 Wave Goodbye | wipe | Preserves a battlefield-reset slot from the precon.
- 1 Ash Barrens | land | Preserves access to the retained basic lands.
- 1 Big Apple, 3 a.m. | land | Keeps a thematic land from the precon's mana base.
- 1 Cinder Glade | land | Preserves red and green mana support.
- 1 City of Brass | land | Provides flexible colored mana for the five-color plan.
- 1 Command Tower | land | Provides flexible colored mana for the commander and supporting spells.
- 1 Dragonskull Summit | land | Preserves black and red mana support.
- 1 Escape Tunnel | land | Keeps access to the retained basic lands.
- 1 Exotic Orchard | land | Preserves flexible multicolor mana support.
- 1 Fabled Passage | land | Provides access to the retained basic lands.
- 1 Hidden Hideout | land | Keeps a thematic land in the precon's mana base.
- 1 Hinterland Harbor | land | Preserves green and blue mana support.
- 1 Path of Ancestry | land | Retains a fixing land suited to the tribal plan.
- 1 Rain-Slicked Copse | land | Preserves a Forest and Island land in the mana base.
- 1 Rootbound Crag | land | Keeps red and green mana support.
- 1 Smoldering Marsh | land | Preserves a Swamp and Mountain land in the mana base.
- 1 Spire Garden | land | Keeps red and green mana support.
- 1 Sunken Hollow | land | Preserves an Island and Swamp land in the mana base.
- 1 Turtle Lair | land | Keeps a thematic land supporting the Turtle-focused deck.
- 1 Undergrowth Stadium | land | Preserves black and green mana support.
- 1 Vernal Fen | land | Keeps a Swamp and Forest land in the mana base.
- 1 Vibrant Cityscape | land | Preserves a land from the precon's existing mana base.
- 1 Mana Confluence | land | Adds flexible colored mana for the five-color plan.
- 1 Secluded Courtyard | land | Adds tribal-focused fixing for the creature package.
- 1 Unclaimed Territory | land | Adds tribal-focused fixing for the creature package.
- 1 Cavern of Souls | land | Supports colored mana for the deck's tribal creatures.
- 1 Tundra | land | Adds a Plains and Island land to strengthen white and blue access.
- 1 Savannah | land | Adds a Forest and Plains land to support green and white casting.
- 1 Scrubland | land | Adds a Plains and Swamp land to support white and black casting.
- 1 Bayou | land | Adds a Swamp and Forest land for black and green access.
- 1 Badlands | land | Adds a Swamp and Mountain land for black and red access.
- 1 Underground Sea | land | Adds an Island and Swamp land for blue and black access.
- 1 Volcanic Island | land | Adds an Island and Mountain land for blue and red access.
- 1 Tropical Island | land | Adds a Forest and Island land for green and blue access.
- 1 Hallowed Fountain | land | Strengthens white and blue access with another Plains and Island land.
- 1 Plains | land | Provides a basic white source and a basic-land search option.
- 1 Island | land | Provides a basic blue source and a basic-land search option.
- 1 Swamp | land | Provides a basic black source and a basic-land search option.
- 1 Mountain | land | Provides a basic red source and a basic-land search option.
- 1 Forest | land | Provides a basic green source and a basic-land search option.

</details>

### 19. the Hobbit family, two colours

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 174 names.

Commander: Thranduil, the Elvenking.

Grade: baseline, score 0.06, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $546.18 to buy, $546.18 the whole deck.

**Summary:** Thranduil leads an Elf-heavy deck that develops mana, builds a creature board, and keeps cards flowing through legendary Elves and supporting draw spells. An Elf-filled graveyard gives the commander additional options, while removal and board wipes help contest stronger opposing boards. Win through creature pressure backed by larger finishers, chief among them the Witch-kings and Troll of Khazad-dûm. The deck favors woodland flavor and a sustained battlefield presence over explosive starts, with a basic-heavy mana base and considerable reliance on creature ramp.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The Elf tribal core with draw and ramp is a clear plan, but the black finishers, Witch-kings and Troll, sit awkwardly with a green-blue Elf base, and several cards such as the Rings and Mox Amber do little for it.
- PLAN theme_fit=yes: It is a Thranduil, the Elvenking commander deck in UBG with every card from the Hobbit sets and an Elf-heavy list, which matches the request for a Hobbit set deck.
- PLAN useful_as_built=yes: With 28 basics and nonbasics split across three colors, plenty of cheap ramp, draw, removal and sweepers, the deck can be played as it stands, though a few lands and the three-color split may cause color friction.
- PLAN summary_honest=partly: The summary names Witch-king of Angmar, Witch-king, Bringer of Ruin and Troll of Khazad-dûm as finishers and they are in the list, but it says an Elf-filled graveyard gives options when the deck has almost no self-mill or recursion, so the graveyard claim is thin.
- [INFO] `curve_summary`: average mana value 2.97 over 63 nonland cards
- [WARN] `profile_off_band`: the count of lands that make two or more of the deck's colors is 4, and bracket 3 wants 17 or more (Elven Passage, Elvenking's Halls, Hobbit Hole, Mirkwood)
- [INFO] `mana_pass`: the builder moved 7 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 9 Forest | land | Supports the deck's green Elf core.
- 9 Island | land | Supplies blue mana for the commander and supporting spells.
- 10 Swamp | land | Supplies black mana for the commander and removal.
- 1 Elven Passage | land | Supports the mana base while retaining the Elven setting.
- 1 Elvenking's Halls | land | Provides a thematic home for the commander.
- 1 Hobbit Hole | land | Adds another land from the deck's setting.
- 1 Minas Morgul, Dark Fortress | land | Supports the black portion of the mana base.
- 1 Mirkwood | land | Keeps the mana base connected to the woodland theme.
- 1 Rivendell | land | Supports the blue portion of the mana base.
- 1 The Black Gate | land | Adds a land supporting the deck's black spells.
- 1 The Shire | land | Supports the green portion of the mana base.
- 1 Arcane Signet | ramp | Provides artifact ramp alongside the creature-based mana package.
- 1 Delighted Halfling | ramp | Adds creature-based ramp for the legendary-heavy plan.
- 1 Elvish Mystic | ramp | Combines ramp with an Elf body.
- 1 Elvish Archdruid | ramp | Adds Elf-based ramp to support board development.
- 1 Wood Elves | ramp | Develops mana while adding another Elf.
- 1 Woodland Weavemaster | ramp | Reinforces the deck's Elf-based ramp package.
- 1 Thranduil the Strategist | ramp | Combines ramp with a legendary Elf for the commander.
- 1 Silvan Reveler | ramp | Adds ramp without moving away from the Elf theme.
- 1 Thranduil's Company | ramp | Supports mana development with another Elf creature.
- 1 Wayfarer's Bauble | ramp | Provides another artifact-based ramp option.
- 1 Captain of Umbar | draw | Adds a creature-based source of card flow.
- 1 Elvish Visionary | draw | Pairs card draw with an Elf body.
- 1 Hithlain Knots | draw | Adds instant-speed card flow.
- 1 Ithilien Kingfisher | draw | Contributes card flow through another creature.
- 1 Lórien Revealed | draw | Helps replenish the hand during longer games.
- 1 Night's Whisper | draw | Provides a dedicated card-draw spell.
- 1 Palantír of Orthanc | draw | Adds an artifact source of card flow.
- 1 Key to the Side-Door | draw | Broadens the artifact-based draw package.
- 1 Uncover the Moon-Letters | draw | Adds an enchantment-based draw option.
- 1 Gollum, Riddle Master | draw | Combines card flow with an additional finishing threat.
- 1 Long Lake Nuisance | draw | Adds another creature-based draw option.
- 1 Stern Scolding | interaction | Provides an instant interaction option.
- 1 Mithril Coat | interaction | Adds equipment-based interaction to the creature plan.
- 1 Bilbo's Ring | interaction | Provides another equipment-based interaction option.
- 1 The One Ring | interaction | Adds a powerful artifact to the interaction package.
- 1 My Precious // Allure of Power | interaction | Offers interaction through an equipment and Adventure card.
- 1 Confusticate and Bebother | interaction | Adds another instant for contesting opposing plays.
- 1 Old Fat Spider Can't See Me | interaction | Diversifies interaction with a Saga.
- 1 Bitter Downfall | removal | Adds dedicated instant-speed removal.
- 1 Bilbo's Deadly Slice | removal | Provides another removal spell for opposing threats.
- 1 Enchanted River's Grasp | removal | Adds enchantment-based removal.
- 1 Merciless Executioner | removal | Combines removal with a creature body.
- 1 Orcish Bowmasters | removal | Adds creature-based removal to the supporting cast.
- 1 Quarrel | removal | Broadens the instant-speed removal package.
- 1 Stir Up Trouble | removal | Provides an additional sorcery-speed answer.
- 1 Uneasy Partings | removal | Adds another instant removal option.
- 1 Arwen, Weaver of Hope | synergy | Adds a legendary Elf to fuel the commander's card-flow engine.
- 1 Celeborn the Wise | synergy | Strengthens the legendary Elf package around Thranduil.
- 1 Galion, Elvenking's Butler | synergy | Provides another legendary Elf for the commander.
- 1 Mirkwood Meditator | synergy | Deepens the deck's Elf-focused creature package.
- 1 Mirkwood Nurturer | synergy | Supports the woodland Elf theme.
- 1 Supper for Spiders | synergy | Adds the shortlist's dedicated synergy spell.
- 1 Thranduil, Sindarin Liege // Silvan Rally | synergy | Adds a legendary Elf with an Adventure to the commander-focused plan.
- 1 Elvenking's Harper | threat | Adds another Elf to the developing creature board.
- 1 Elven Raft-Steerer | threat | Expands the Elf-based creature presence.
- 1 Cantankerous Keepers | threat | Reinforces the battlefield with another Elf creature.
- 1 Galadhrim Guide | threat | Keeps the creature package concentrated on Elves.
- 1 Grey Havens Navigator | threat | Adds another Elf to the deck's battlefield presence.
- 1 Guardian of the Halls | threat | Contributes an Elf Soldier to the creature plan.
- 1 Haunt of the Dead Marshes | threat | Adds an Elf creature to the black supporting cast.
- 1 Lothlórien Lookout | threat | Supplies another Elf for sustained creature pressure.
- 1 Nimrodel Watcher | threat | Expands the Elf-heavy battlefield.
- 1 Mirkwood Pathmaker | threat | Adds another woodland Elf to the creature suite.
- 1 Willow-Wind | threat | Diversifies the creature threats beyond Elves.
- 1 Troll of Khazad-dûm | wincon | Provides a dedicated finishing creature.
- 1 Witch-king of Angmar | wincon | Adds a legendary finisher for closing games.
- 1 Witch-king, Bringer of Ruin | wincon | Provides another substantial endgame threat.
- 1 Languish | wipe | Provides a board reset when opposing creatures take over.
- 1 Raise the Palisade | wipe | Adds a reset option to support the creature-focused game plan.
- 1 Gnashing of Teeth | wipe | Rounds out the board-wipe package.
- 1 Mox Amber | synergy | the mana pass added it to bring the mana base inside the power level
- 1 Giant's Boulder | removal | the mana pass added it to bring the mana base inside the power level
- 1 Pelargir Survivor | ramp | the mana pass added it to bring the mana base inside the power level

</details>

### 20. the Hobbit family, a delegated commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 77 names.

Commander: Smaug the Magnificent.

Grade: baseline, score 0.06, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $518.81 to buy, $518.81 the whole deck.

**Summary:** Your commander leads a Treasure-fueled combat deck: develop mana, build up Treasures, and turn attacks into damage while preparing the larger Dragon finishers. Dwarves, Goblins, and Equipment provide a supporting battlefield, with removal and board wipes helping manage opposing threats. The deck wins through commander pressure and its Dragon finishers, trading a deep roster of Dragons and abundant card draw for a broader Hobbit-themed supporting cast and a substantial mana-development package.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The land count sits above the norm of the format, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The red Treasure, Dwarf, Goblin and equipment cards loosely support a combat plan, but the list is a pile of 1-of pieces with many weak equipment and filler cards, and only three Dragons back the summary's claim of a deep Dragon roster.
- PLAN theme_fit=partly: It is a mono-red Hobbit-set Commander deck led by Smaug the Magnificent, which fits the set and the dragon commander, but Dragons are barely present in the 99 so the dragons theme is thin.
- PLAN useful_as_built=partly: With 38 lands (34 Mountains) plus many mana rocks the deck will cast its spells, but it is flooded with mana, has few threats that win the game, and includes many low-impact equipment and filler cards, so it plays clunkily.
- PLAN summary_honest=no: The summary claims a deep roster of Dragons and abundant card draw, but the list has only Smaug plus three Dragon finishers (Cavern-Hoard Dragon, Desert Were-Worm, Smaug the Great Calamity) and just a handful of modest draw cards.
- [INFO] `curve_summary`: average mana value 3.23 over 61 nonland cards
- [WARN] `profile_off_band`: the draw count is 6, and bracket 3 wants 8 to 14

<details><summary>The deck list</summary>

- 34 Mountain | land | Provides a dependable red mana base for the commander and supporting spells.
- 1 Dragon-Cursed Halls | land | Adds a thematic nonbasic land to the mana base.
- 1 Rogue's Passage | land | Adds a utility land alongside the basic-heavy mana base.
- 1 The Lonely Mountain | land | Supports the red mana base while reinforcing the deck's theme.
- 1 Treasure Vault | land | Adds a thematic artifact land to the mana base.
- 1 Arcane Signet | ramp | Accelerates the commander and subsequent threats.
- 1 Bag End Banquet | ramp | Adds another source of mana development.
- 1 Burn, Burn, Tree and Fern | ramp | Supports the mana development needed for expensive Dragons.
- 1 Dragon's Desire | ramp | Adds ramp that fits the Dragon-centered plan.
- 1 Fíli and Kíli, Joyous | ramp | Provides creature-based ramp among the supporting Dwarves.
- 1 Long-Bodied Grey Dog | ramp | Adds another creature-based ramp option.
- 1 Mox Amber | ramp | Provides a low-cost mana piece for the legendary-heavy supporting cast.
- 1 Orcrist, Goblin-cleaver | ramp | Combines the Equipment theme with mana development.
- 1 The Misty Mountains Cold | ramp | Adds another ramp option for the deck's larger threats.
- 1 The Reaver Cleaver | ramp | Supports mana development through the Equipment package.
- 1 Thorin, Company's Leader | ramp | Supplies ramp while also serving as a marked finisher.
- 1 Troop of Ponies | ramp | Adds creature-based mana development to the supporting cast.
- 1 Wayfarer's Bauble | ramp | Provides an inexpensive ramp piece for the opening turns.
- 1 Balin, Loremaster | draw | Adds card flow on a thematic Dwarf creature.
- 1 Key to the Side-Door | draw | Supplies card flow from an artifact slot.
- 1 Palantír of Orthanc | draw | Helps sustain resources beyond the opening hand.
- 1 Ragged Short Spear | draw | Adds card flow within the Equipment package.
- 1 Thrór's Map | draw | Provides another artifact-based source of card flow.
- 1 Óin the Brave | draw | Adds card flow while maintaining the Dwarf supporting cast.
- 1 Bilbo's Ring | interaction | Adds interaction within the Equipment package.
- 1 Dwarven Mattock | interaction | Provides another Equipment-based interaction option.
- 1 Mithril Coat | interaction | Adds interaction suited to the creature-focused plan.
- 1 The One Ring | interaction | Broadens the deck's artifact-based interaction package.
- 1 Battle-Scarred Goblin | removal | Adds removal on a supporting Goblin body.
- 1 Fire of Orthanc | removal | Provides removal to help clear the way for attacks.
- 1 Gandalf, Spark Starter | removal | Adds removal on a legendary creature.
- 1 Giant's Boulder | removal | Provides an artifact-based removal option.
- 1 Goblin Cratermaker | removal | Adds a removal option among the supporting Goblins.
- 1 Goblin Fireleaper | removal | Combines a Goblin body with the removal package.
- 1 Improvised Club | removal | Provides instant-speed removal support.
- 1 Inferno Titan | removal | Adds removal on a substantial creature.
- 1 Pinecone Strike | removal | Adds another instant-speed removal option.
- 1 Smite the Deathless | removal | Provides additional instant-speed removal.
- 1 The Black Arrow | removal | Adds removal within the Equipment package.
- 1 Thorin, Mountain-king | removal | Supplies removal on a thematic legendary Dwarf.
- 1 Call Forth the Tempest | wipe | Provides a board-clearing option when opposing forces grow too large.
- 1 Desolation of Smaug | wipe | Adds a thematic board wipe to the Dragon plan.
- 1 Glóin the Mighty // Easy Pickings | wipe | Adds another board-clearing option to the supporting cast.
- 1 Bothersome Noisemaker | threat | Adds a Goblin threat to the supporting attack force.
- 1 Dwarven Mauler | threat | Adds a Dwarf threat alongside the larger Dragons.
- 1 Dwarven Warriors | threat | Supplies another creature for the combat-focused plan.
- 1 Dáin Ironfoot | threat | Adds a legendary Dwarf threat and a marked finisher.
- 1 Gandalf, Goblins' Bane // Flameshape | threat | Adds another creature threat to the supporting cast.
- 1 Goblin-town Flunkies | threat | Adds a Goblin body to maintain combat pressure.
- 1 Gundabad Opportunist | threat | Broadens the deck's supporting Goblin threats.
- 1 Guttersnipe | threat | Adds another Goblin threat alongside the spell package.
- 1 Iron Hills Stalwart | threat | Provides another Dwarf for the supporting attack force.
- 1 Misty Mountains Raider | threat | Adds a thematic Goblin attacker.
- 1 Olog-hai Crusher | threat | Adds a Troll threat to diversify the creature force.
- 1 Orcish Siegemaster | threat | Adds an Orc threat alongside the Dwarves and Goblins.
- 1 Last Light of Durin's Day | synergy | Provides a dedicated synergy piece for the supporting cast.
- 1 Andúril, Flame of the West | synergy | Adds Equipment support to the creature-based attack plan.
- 1 Andúril, Narsil Reforged | synergy | Broadens the Equipment package supporting combat.
- 1 Glamdring | synergy | Adds another Equipment option for the deck's creatures.
- 1 Long-Lost Lances | synergy | Supports the combat plan through the Equipment package.
- 1 Smaug's Fury | synergy | Adds a thematic spell alongside the Dragon finishers.
- 1 Sting, Bilbo's Sword | synergy | Adds another Equipment piece to support the creature force.
- 1 Cavern-Hoard Dragon | wincon | Provides a Dragon finisher for closing out the game.
- 1 Desert Were-Worm | wincon | Adds another Dragon finisher to the top end.
- 1 Smaug, the Great Calamity // Spew Flame | wincon | Provides a central Dragon finisher for the late game.
- 1 Getaway Barrel | other | Rounds out the thematic supporting artifact package.

</details>

### 21. the Hobbit family, mana from outside

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 156 names.

Commander: Smaug the Impenetrable.

Grade: baseline, score 0.03, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $708.61 to buy, $708.61 the whole deck.

**Summary:** Smaug the Impenetrable leads a black-red resource-building deck with thematic creatures, equipment, and Dragon finishers. Develop your mana and card supply, use removal to keep opposing threats manageable, then push toward a finish with Smaug and the supporting threats. The equipment package keeps creature combat central to the plan, while board wipes offer a reset when the table gets crowded. The deck gives up early speed and a streamlined combo finish in favor of a broader, creature-heavy game.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The cards mostly share a black-red Hobbit goblin-and-equipment shell with ramp into Dragons, but about a dozen equipment pieces and several vanilla-ish goblins have no real payoff, and the board wipes pull against a creature-heavy plan.
- PLAN theme_fit=yes: The deck is led by Smaug the Impenetrable, stays in black-red, draws nearly every nonland card from the Hobbit sets, and takes only its lands and Arcane Signet from outside, as the request allowed.
- PLAN useful_as_built=partly: The mana base of 19 basics plus duals and fetches is playable and hits two to four lands often, but the low count of real ramp, the number of low-impact equipment cards, and few strong finishers leave it slow and clunky at bracket 3.
- PLAN summary_honest=partly: The claim that equipment keeps creature combat central holds since the list has around a dozen equipment pieces, but the claim of a resource-building plan with Dragon finishers is thin because only a few Dragons appear and many ramp slots are just equipment or sagas.
- [INFO] `curve_summary`: average mana value 3.24 over 63 nonland cards
- [INFO] `mana_pass`: the builder moved 6 cards of the mana base to bring the deck inside its power level
- [WARN] `outside_requested_set`: the sets you named do not hold 14 cards: Arid Mesa, Blood Crypt, Bloodstained Mire, Command Tower, Dragonskull Summit, Evolving Wilds, Exotic Orchard, Fabled Passage, Marsh Flats, Path of Ancestry, Polluted Delta, Terramorphic Expanse, Urborg, Tomb of Yawgmoth, Wooded Foothills

<details><summary>The deck list</summary>

- 1 Command Tower | land | Provides fixing for the black-red mana base.
- 1 Blood Crypt | land | Provides both colors and supports the fetch lands.
- 1 Dragonskull Summit | land | Adds another source of both deck colors.
- 1 Exotic Orchard | land | Adds flexible color fixing.
- 1 Mount Doom | land | Provides both colors in a thematic land slot.
- 1 Path of Ancestry | land | Adds color fixing to support the creature-heavy plan.
- 1 Arid Mesa | land | Adds fixing through the dual-land package.
- 1 Bloodstained Mire | land | Connects the basic lands and dual-land package.
- 1 Marsh Flats | land | Supports access to the deck's black mana and dual land.
- 1 Polluted Delta | land | Supports access to the deck's black mana and dual land.
- 1 Wooded Foothills | land | Supports access to the deck's red mana and dual land.
- 1 Urborg, Tomb of Yawgmoth | land | Reinforces the black side of the mana base.
- 1 The Lonely Mountain | land | Adds a thematic Mountain to Smaug's mana base.
- 10 Swamp | land | Supplies reliable black mana for the heavier black requirements.
- 9 Mountain | land | Supplies reliable red mana for the ramp and creature packages.
- 1 Arcane Signet | ramp | Adds a mana rock to help develop toward Smaug.
- 1 Mox Amber | ramp | Adds a compact mana rock alongside the legendary creatures.
- 1 Bag End Banquet | ramp | Adds thematic artifact-based ramp.
- 1 Bolg's Company | ramp | Combines the ramp package with a Goblin creature slot.
- 1 Burn, Burn, Tree and Fern | ramp | Adds Saga-based ramp to the resource plan.
- 1 Dragon's Desire | ramp | Adds a thematic ramp spell for the Dragon plan.
- 1 Long-Bodied Grey Dog | ramp | Adds creature-based ramp to diversify mana development.
- 1 Orcrist, Goblin-cleaver | ramp | Connects the equipment package with mana development.
- 1 The Misty Mountains Cold | ramp | Adds another thematic ramp Saga.
- 1 The Reaver Cleaver | ramp | Adds equipment-based ramp to support the creature plan.
- 1 Cavern-Hoard Dragon | ramp | Combines ramp with a Dragon finisher.
- 1 Balin, Loremaster | draw | Adds card draw on a legendary creature.
- 1 Gollum, Riddle Master | draw | Combines card draw with a marked finisher.
- 1 Key to the Side-Door | draw | Adds artifact-based card draw.
- 1 Night's Whisper | draw | Adds a dedicated draw spell to maintain resources.
- 1 Palantír of Orthanc | draw | Adds another artifact to the card-draw package.
- 1 Rage into the Valley | draw | Adds thematic spell-based card draw.
- 1 Ragged Short Spear | draw | Connects card draw with the equipment package.
- 1 The Master of Lake-town | draw | Adds a thematic legendary creature to the draw package.
- 1 Thrór's Map | draw | Adds another thematic artifact for card draw.
- 1 Óin the Brave | draw | Adds card draw while maintaining the creature presence.
- 1 Bilbo's Ring | interaction | Adds thematic equipment to the interaction package.
- 1 Dwarven Mattock | interaction | Adds another equipment-based interaction slot.
- 1 Mithril Coat | interaction | Adds legendary equipment to the interaction package.
- 1 My Precious // Allure of Power | interaction | Adds an equipment-and-Adventure interaction slot.
- 1 The One Ring | interaction | Adds a thematic artifact to the interaction package.
- 1 Andúril, Flame of the West | interaction | Adds another equipment option for contested combat.
- 1 Andúril, Narsil Reforged | interaction | Broadens the equipment package for creature-based play.
- 1 Getaway Barrel | interaction | Adds a supporting artifact to the interaction package.
- 1 Bilbo's Deadly Slice | removal | Adds instant-speed removal to the answer package.
- 1 Bitter Downfall | removal | Adds another dedicated removal instant.
- 1 Goblin Cratermaker | removal | Combines removal with a Goblin creature slot.
- 1 Goblin Fireleaper | removal | Adds another Goblin to the creature-based removal package.
- 1 Improvised Club | removal | Adds a removal instant alongside the creature package.
- 1 Orcish Bowmasters | removal | Adds creature-based removal without leaving the setting.
- 1 Smite the Deathless | removal | Rounds out the dedicated instant removal.
- 1 Smaug, the Great Calamity // Spew Flame | removal | Combines removal with a Dragon finisher.
- 1 Witch-king of Angmar | removal | Combines creature-based removal with a marked finisher.
- 1 Supper for Spiders | synergy | Provides a dedicated thematic synergy spell.
- 1 Smaug's Fury | synergy | Adds a Smaug-themed spell to the supporting package.
- 1 Guttersnipe | synergy | Adds a Goblin supporting creature alongside the spell package.
- 1 Sting, Bilbo's Sword | synergy | Extends the equipment package supporting the creatures.
- 1 Gundabad Opportunist | synergy | Adds a Goblin Rogue to the supporting creature package.
- 1 Great Ugly-Looking Goblin // Clap! Snap! | synergy | Adds a Goblin-and-Adventure slot to the supporting package.
- 1 Stony-Voiced Goblins | synergy | Adds another thematic Goblin supporting creature.
- 1 Desert Were-Worm | threat | Adds a marked creature finisher to the threat package.
- 1 Dreaded Bat-Cloud | threat | Diversifies the creature threats beyond Goblins.
- 1 Fearsome Goblin Pair | threat | Adds another Goblin threat for the creature plan.
- 1 Goblin-town Flunkies | threat | Maintains the thematic Goblin presence.
- 1 Gollum the Abandoned | threat | Adds a legendary creature to the threat package.
- 1 Great Goblin, Foul-Hearted | threat | Adds a thematic legendary Goblin threat.
- 1 Misty Mountains Raider | threat | Adds another Goblin Soldier to maintain creature pressure.
- 1 Nighthowl Pursuer | threat | Diversifies the threat package with a Wolf.
- 1 Olog-hai Crusher | threat | Adds a Troll Soldier to the creature threats.
- 1 Orcish Siegemaster | threat | Adds an Orc Soldier to the creature-pressure plan.
- 1 The Great Goblin | threat | Adds another legendary Goblin threat.
- 1 Sauron, the Lidless Eye | threat | Adds a thematic legendary threat alongside the Goblins.
- 1 Down, Down to Goblin-town | wincon | Adds a marked Saga finisher to close developed games.
- 1 Inside Information | wincon | Adds a marked sorcery finisher alongside the creatures.
- 1 Smaug, Wicked Worm | wincon | Adds another Dragon finisher with a ramp role.
- 1 Desolation of Smaug | wipe | Adds a thematic board wipe for crowded games.
- 1 Glóin the Mighty // Easy Pickings | wipe | Adds a creature-and-Adventure slot to the wipe package.
- 1 Languish | wipe | Adds a dedicated board wipe to complement spot removal.
- 1 Fabled Passage | land | the mana pass added it to bring the mana base inside the power level
- 1 Evolving Wilds | land | the mana pass traded a basic land for this land, which makes two or more of the deck's colors
- 1 Terramorphic Expanse | land | the mana pass traded a basic land for this land, which makes two or more of the deck's colors
- 1 Hobbit Hole | land | the mana pass traded a basic land for this land, which makes two or more of the deck's colors

</details>

### 22. a set family and a card from outside it

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 175 names.

Commander: Thranduil, the Elvenking.

Grade: baseline, score 0.05, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $550.97 to buy, $550.97 the whole deck.

**Summary:** Thranduil leads an Elf-centered midrange deck that develops mana, builds a creature board, and turns arriving legendary Elves into fresh cards. Interaction and board resets help sustain the game until creature pressure and the designated finishers can close it out. The deck favors a broad thematic supporting cast over a narrow combo plan. Color fixing is its main weakness: the shortlist does not contain enough nonbasic lands to provide the requested density of multicolor sources, so awkward opening hands remain a significant compromise.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The Elf creature core, card draw, and ramp all point toward a Thranduil midrange plan, but the black cards (Languish, Bowmasters, Witch-kings, removal) and the many blue cards spread the deck across three colors with little link to the Elf theme.
- PLAN theme_fit=yes: The deck is led by Thranduil, the Elvenking, is drawn almost entirely from the Hobbit sets, stays in his UBG colors, is built around Elves, and includes the requested Sol Ring.
- PLAN useful_as_built=partly: It has plenty of ramp, draw, removal and creatures, but 27 basics plus only eight nonbasics leave a three-color mana base strained by cards like BB wipes and GG creatures, and the finishers are thin.
- PLAN summary_honest=partly: The summary's claim of an Elf midrange deck with ramp, interaction, wipes and a broad creature cast matches the list, but it says legendary Elves become fresh cards while few cards in the list actually do that, and it lists 'designated finishers' when the wincon slot holds only a Troll and two Witch-kings.
- [INFO] `curve_summary`: average mana value 3.09 over 64 nonland cards
- [WARN] `profile_off_band`: the worst color's share of the sources it needs is 0.78, and bracket 3 wants 0.85 or more (sources of requirement: U 14.75 of 19, B 18.75 of 23, G 17.25 of 22)
- [WARN] `profile_off_band`: the count of lands that make two or more of the deck's colors is 4, and bracket 3 wants 17 or more (Elven Passage, Elvenking's Halls, Hobbit Hole, Mirkwood)
- [INFO] `mana_pass`: the builder moved 5 cards of the mana base to bring the deck inside its power level
- [WARN] `outside_requested_set`: the sets you named do not hold 1 card: Sol Ring

<details><summary>The deck list</summary>

- 10 Forest | land | Prioritizes green mana for the Elf core and early ramp.
- 7 Island | land | Supplies blue mana for card flow and interaction.
- 10 Swamp | land | Supplies black mana for removal and finishers.
- 1 Elven Passage | land | Adds an Elf-themed land to the mana base.
- 1 Elvenking's Halls | land | Keeps the mana base rooted in the Elvenking's setting.
- 1 Hobbit Hole | land | Adds another available nonbasic land.
- 1 Minas Morgul, Dark Fortress | land | Adds a legendary land to the mana base.
- 1 Mirkwood | land | Reinforces the woodland theme in the land slots.
- 1 Rivendell | land | Adds another Elven location to the mana base.
- 1 The Black Gate | land | Adds another legendary land to the mana base.
- 1 The Shire | land | Keeps a thematic land alongside the woodland locations.
- 1 Sol Ring | ramp | Keeps the requested mana accelerator.
- 1 Arcane Signet | ramp | Provides noncreature mana development.
- 1 Delighted Halfling | ramp | Adds an early creature-based ramp option.
- 1 Elvish Mystic | ramp | Combines early ramp with an Elf body.
- 1 Elvish Archdruid | ramp | Keeps a ramp slot within the Elf core.
- 1 Wood Elves | ramp | Develops mana while increasing the Elf presence.
- 1 Silvan Reveler | ramp | Adds another Elf-based ramp piece.
- 1 Thranduil the Strategist | ramp | Combines ramp with a legendary Elf for the commander's trigger.
- 1 Thranduil's Company | ramp | Supports mana development without leaving the Elf theme.
- 1 Woodland Weavemaster | ramp | Adds an Elf Druid to the ramp package.
- 1 Elvish Visionary | draw | Supplies card flow on an Elf body.
- 1 Night's Whisper | draw | Provides a compact spell-based draw option.
- 1 Hithlain Knots | draw | Adds instant-speed card flow.
- 1 Lórien Revealed | draw | Provides a sorcery-based draw option.
- 1 Palantír of Orthanc | draw | Diversifies the draw package with a legendary artifact.
- 1 Key to the Side-Door | draw | Adds another artifact-based card-flow option.
- 1 Uncover the Moon-Letters | draw | Provides card flow from an enchantment slot.
- 1 Fateful Discovery | draw | Broadens the enchantment-based draw package.
- 1 Great Gilded Boat | draw | Adds a Vehicle-based draw option.
- 1 Captain of Umbar | draw | Provides card flow from a creature slot.
- 1 Ithilien Kingfisher | draw | Adds another creature-based draw option.
- 1 Stern Scolding | interaction | Provides an instant interaction slot.
- 1 Mithril Coat | interaction | Diversifies interaction with legendary Equipment.
- 1 Bilbo's Ring | interaction | Adds a thematic Equipment-based interaction option.
- 1 Elrond, Moon-Reader | interaction | Combines interaction with a legendary Elf for Thranduil.
- 1 Confusticate and Bebother | interaction | Provides another instant interaction option.
- 1 Old Fat Spider Can't See Me | interaction | Diversifies interaction with a Saga.
- 1 The One Ring | interaction | Adds a legendary artifact to the interaction package.
- 1 Bitter Downfall | removal | Provides an instant removal option.
- 1 Merciless Executioner | removal | Adds removal from a creature slot.
- 1 Orcish Bowmasters | removal | Provides another creature-based removal option.
- 1 Enchanted River's Grasp | removal | Diversifies removal with an Aura.
- 1 Bilbo's Deadly Slice | removal | Adds a thematic instant removal option.
- 1 Quarrel | removal | Broadens the instant removal package.
- 1 Stir Up Trouble | removal | Provides a sorcery-based removal option.
- 1 Troll Negotiations | removal | Adds another sorcery to the removal package.
- 1 Uneasy Partings | removal | Keeps another removal option available at instant speed.
- 1 Languish | wipe | Provides a board-reset option when opponents build ahead.
- 1 Raise the Palisade | wipe | Adds another board-reset option to the Elf strategy.
- 1 Gnashing of Teeth | wipe | Rounds out the package for handling crowded boards.
- 1 Arwen, Weaver of Hope | synergy | Adds a legendary Elf to support Thranduil's draw trigger.
- 1 Celeborn the Wise | synergy | Provides another legendary Elf for the commander's trigger.
- 1 Galion, Elvenking's Butler | synergy | Strengthens the legendary Elf core around Thranduil.
- 1 Thranduil, Sindarin Liege // Silvan Rally | synergy | Adds a thematic legendary Elf with an Adventure.
- 1 Haunt of the Dead Marshes | synergy | Keeps an additional Elf in the creature package.
- 1 Supper for Spiders | synergy | Adds the shortlist's dedicated synergy spell.
- 1 Elvenking's Harper | synergy | Reinforces the commander's Elf-centered supporting cast.
- 1 Boughside Wanderers | threat | Adds an Elf body for the tribal combat plan.
- 1 Cantankerous Keepers | threat | Provides another Elf Soldier for board presence.
- 1 Elven Raft-Steerer | threat | Expands the Elf creature base.
- 1 Galadhrim Guide | threat | Adds an Elf Scout to the combat plan.
- 1 Grey Havens Navigator | threat | Keeps another Elf body in the threat package.
- 1 Guardian of the Halls | threat | Provides a thematic Elf Soldier for board presence.
- 1 Lothlórien Lookout | threat | Adds another Elf Scout to the battlefield plan.
- 1 Mirkwood Meditator | threat | Keeps the creature package focused on woodland Elves.
- 1 Mirkwood Nurturer | threat | Adds a thematic Elf Ranger to the board.
- 1 Mirkwood Pathmaker | threat | Provides another Elf Ranger for the combat plan.
- 1 Nimrodel Watcher | threat | Rounds out the Elf Scout contingent.
- 1 Mirkwood Elk | threat | Adds a woodland creature alongside the Elf core.
- 1 Troll of Khazad-dûm | wincon | Provides a designated finisher for closing games.
- 1 Witch-king of Angmar | wincon | Adds a legendary finisher to the late-game plan.
- 1 Witch-king, Bringer of Ruin | wincon | Provides another designated finisher when the Elf board needs support.
- 1 Mox Amber | synergy | the mana pass added it to bring the mana base inside the power level
- 1 Giant's Boulder | ramp | the mana pass added it to bring the mana base inside the power level

</details>

### 23. two set families at once

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 203 names.

Commander: Kíli the Resourceful.

Grade: typical, score 0.45, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $320.46 to buy, $320.46 the whole deck.

**Summary:** Kíli leads a white Dwarf-and-Equipment deck that builds an enduring story through artifacts and legendary permanents, then turns incoming Dwarves and Equipment into a steady supply of cards. Develop mana early, build a board of creatures and Equipment, and use Kíli's equip support to keep combat pressure moving. Equipment-backed attacks are the main route to victory, with Angel of the Ruins providing a dedicated finishing threat. Removal, interaction, and board resets help keep opponents in check, but the deck gives up combo speed and depends on maintaining a useful battlefield.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

This deck plan aims for 2 or more finishers, the cards that win the game, and the deck holds 1.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every card is a white creature, Equipment, or artifact that feeds the Dwarf-and-Equipment plan, with ramp and draw pieces supporting it and no off-plan cards pulling elsewhere.
- PLAN theme_fit=yes: Every card is from the Hobbit and Bloomburrow sets, the commander is a Hobbit Dwarf, and the mono-white Equipment plan is a sensible pick for a bracket 3 game.
- PLAN useful_as_built=partly: The mana base of 30 Plains plus utility lands and a dozen mana rocks casts everything smoothly and the deck has plenty of ramp, draw, and removal, but with a single finisher and many low-impact Equipment pieces it may struggle to close out games.
- PLAN summary_honest=partly: The summary says Angel of the Ruins is a dedicated finishing threat and it is in the list, but it also calls the Equipment-backed attacks the main route, and the deck holds only one real finisher, and Kíli's 'equip support' is not clearly what the commander does.
- [INFO] `curve_summary`: average mana value 2.73 over 62 nonland cards
- [WARN] `finisher_short`: the finisher count is 1, and this deck plan wants 2 or more (Angel of the Ruins)

<details><summary>The deck list</summary>

- 30 Plains | land | Provides the white mana needed throughout the deck.
- 1 Command Tower | land | Adds a white source to the mana base.
- 1 Castle Ardenvale | land | Adds a white-producing utility land.
- 1 Minas Tirith | land | Adds a legendary land to support Kíli's story plan.
- 1 Lupinflower Village | land | Adds a Bloomburrow land to the white mana base.
- 1 Rogue's Passage | land | Adds a utility land alongside the deck's combat threats.
- 1 Treasure Vault | land | Its artifact type contributes to Kíli's enduring story.
- 1 Fountainport | land | Adds a utility land without diluting the spell package.
- 1 Sol Ring | ramp | Supplies inexpensive artifact ramp.
- 1 Mox Amber | ramp | Adds low-cost ramp alongside the legendary commander.
- 1 Arcane Signet | ramp | Supplies artifact ramp for the deck's white spells.
- 1 Mind Stone | ramp | Adds an inexpensive artifact to the ramp package.
- 1 Thought Vessel | ramp | Develops mana while contributing an artifact to the board.
- 1 Fellwar Stone | ramp | Adds another inexpensive mana rock.
- 1 Ornithopter of Paradise | ramp | Combines creature presence with artifact-based ramp.
- 1 Wayfarer's Bauble | ramp | Provides an early artifact-based ramp option.
- 1 Loyal Warhound | ramp | Adds creature-based ramp to balance the mana rocks.
- 1 Patchwork Banner | ramp | Adds ramp that fits the creature-focused board.
- 1 Orcrist, Goblin-cleaver | ramp | Combines the ramp package with Kíli's Equipment theme.
- 1 Burnished Hart | ramp | Adds another artifact creature to the ramp package.
- 1 Skullclamp | draw | Combines card draw with Kíli's Equipment support.
- 1 Belladonna Took | draw | Adds a legendary draw option that supports the story plan.
- 1 Caretaker's Talent | draw | Provides an enchantment-based source of cards.
- 1 Circuit Mender | draw | Adds card draw on an artifact creature.
- 1 Dawn of a New Age | draw | Diversifies the draw package with an enchantment.
- 1 Errand-Rider of Gondor | draw | Provides card draw while adding a creature to the board.
- 1 Esgaroth Garrison | draw | Adds another creature-based draw option.
- 1 Key to the Side-Door | draw | Supplies artifact-based draw and supports the story plan.
- 1 The Arkenstone // Seek the Heart | draw | Adds a legendary artifact to the draw package.
- 1 The Queen of Dale | draw | Combines a legendary creature with the deck's draw needs.
- 1 Bofur, Reliable Guardian // Concerted Care | interaction | Adds interaction while its Dwarf type supports Kíli's draw engine.
- 1 Bilbo's Ring | interaction | Combines interaction with a legendary Equipment for Kíli.
- 1 Dwarven Mattock | interaction | Adds an Equipment-based interaction option.
- 1 Mithril Coat | interaction | Adds legendary Equipment to the interaction package.
- 1 Swiftfoot Boots | interaction | Provides interaction through an Equipment that fits the commander.
- 1 Reprieve | interaction | Adds an instant-speed interaction option.
- 1 Dawn's Truce | interaction | Diversifies the deck's instant-speed interaction.
- 1 Flowering of the White Tree | interaction | Adds interaction on a legendary enchantment.
- 1 Swords to Plowshares | removal | Provides inexpensive removal.
- 1 Generous Gift | removal | Adds instant-speed removal to the answer package.
- 1 Repel Calamity | removal | Provides another removal option from Bloomburrow.
- 1 Banishing Light | removal | Diversifies removal with an enchantment.
- 1 Skyclave Apparition | removal | Adds removal on a creature body.
- 1 Loran of the Third Path | removal | Combines removal with a legendary creature for the story plan.
- 1 Nettle Guard | removal | Adds a Bloomburrow creature to the removal package.
- 1 The Black Arrow | removal | Combines removal with legendary Equipment support.
- 1 Blade Splicer | synergy | Adds a creature-based synergy piece to the board.
- 1 Carrot Cake | synergy | Adds a Food artifact that contributes to the story plan.
- 1 Dáin, Lord of the Iron Hills | synergy | Its legendary Dwarf type fits both sides of Kíli's engine.
- 1 Iron Hills Blacksmith | synergy | Adds a Dwarf that supports Kíli's draw engine.
- 1 Ori, Keeper of Songs | synergy | Adds another legendary Dwarf to the deck's central theme.
- 1 Maskwood Nexus | synergy | Adds an artifact synergy piece to support the developed board.
- 1 Dwarven Provisioner | synergy | Its Dwarf type supports Kíli's recurring card draw.
- 1 Fíli the Pathfinder | threat | Adds a legendary Dwarf threat alongside Kíli.
- 1 Andúril, Flame of the West | threat | Adds legendary Equipment to the combat plan.
- 1 Andúril, Narsil Reforged | threat | Provides another legendary Equipment threat.
- 1 Dwarven Shortsword | threat | Adds Equipment for Kíli's draw and equip support.
- 1 Dúnedain Blade | threat | Expands the Equipment package for creature combat.
- 1 Sting, Bilbo's Sword | threat | Adds legendary Equipment to the attacking board.
- 1 Sword of the Squeak | threat | Connects the Bloomburrow theme with the Equipment plan.
- 1 Brightblade Stoat | threat | Adds a creature to carry the deck's Equipment.
- 1 Flowerfoot Swordmaster | threat | Adds a Bloomburrow creature to the combat package.
- 1 Shrike Force | threat | Adds another creature threat for the Equipment plan.
- 1 Warren Warleader | threat | Adds a creature threat to the deck's developed board.
- 1 Steelburr Champion | threat | Adds another Bloomburrow attacker to equip.
- 1 Angel of the Ruins | wincon | Provides the shortlist's explicitly marked finishing threat.
- 1 Helm of the Host | wincon | Provides a top-end Equipment payoff for the combat plan.
- 1 Dusk // Dawn | wipe | Provides a board-reset option.
- 1 Ori, Plate Stacker | wipe | Combines a wipe option with a legendary Dwarf.
- 1 Starfall Invocation | wipe | Adds another board reset for crowded battlefields.

</details>

### 24. a 60-card deck from one set

Format: Modern. Theme: aggro. Pool: any_card. Shortlist: 96 names.

Grade: baseline, score 0.36, model 20260923T202806Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $25.22 to buy, $25.22 the whole deck.

**Summary:** This Bloomburrow mono-red aggro deck centers its offense on Mice, backed by a few other red creatures. Build an attacking board, use removal to clear the way, and keep pressure on with the draw package; Dragonhawk, Fate's Tempest provides a finishing threat when combat alone falls short. The deck favors creature pressure over a broad spell toolbox, leaving it more vulnerable to stalled battlefields and repeated creature removal.

The quality model grades this deck at the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every card is a cheap red creature, mostly Mice, backed by cheap burn and Might of the Meek to push damage, so the list carries one aggressive plan.
- PLAN theme_fit=yes: The deck is mono-red aggro built entirely from Bloomburrow cards and lands, matching the request for casual Modern.
- PLAN useful_as_built=yes: With 24 lands, a low curve topped by two five-drops, and a full set of one-drops and two-drops, the deck can be played as it stands, though 24 lands is a bit heavy for such a low curve.
- PLAN summary_honest=yes: The summary says the offense centers on Mice with red creatures and burn, and the list has 14 Mice-type creatures plus Take Out the Trash and Rabid Gnaw, with Dragonhawk as a two-copy top-end finisher.
- [INFO] `curve_summary`: average mana value 1.94 over 36 nonland cards

<details><summary>The deck list</summary>

- 22 Mountain | land | Provides the dependable red mana needed for the aggressive creature plan.
- 2 Rockface Village | land | Complements the creature-heavy plan while keeping the mana base mostly basic.
- 4 Heartfire Hero | threat | Forms part of the Mouse-focused attacking core.
- 4 Hired Claw | threat | Adds another aggressive creature to maintain pressure.
- 4 Manifold Mouse | threat | Strengthens the Mouse contingent and keeps the deck focused on combat.
- 2 Dragonhawk, Fate's Tempest | threat | Provides a finishing threat when the initial assault is not enough.
- 4 Emberheart Challenger | synergy | Connects the Mouse creature package with the deck's supporting spells.
- 4 Roughshod Duo | synergy | Adds creature-based synergy without moving away from the attacking plan.
- 2 Brazen Collector | ramp | Supplies ramp through a creature rather than a dedicated mana artifact.
- 4 Might of the Meek | draw | Provides instant-speed card flow alongside the creature-heavy strategy.
- 2 Whiskerquill Scribe | draw | Adds card flow while increasing the deck's Mouse presence.
- 4 Take Out the Trash | removal | Supplies instant-speed removal to help keep attacks productive.
- 2 Rabid Gnaw | removal | Rounds out the removal package for a deck built around creature combat.

</details>

### 25. use no card of an owned precon

Format: Commander. Theme: Avengers superheroes. Pool: owned_first. Shortlist: 195 names.

Commander: Captain America, Super-Soldier.

Grade: typical, score 0.43, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $138.50 the whole deck.

**Summary:** Build an Avengers team around your commander's shield-backed protection, using artifact ramp and steady card flow to develop a threatening board. The S.H.I.E.L.D. contingent supplies thematic support while removal and protective interaction help the Heroes stay in the fight. Win through creature pressure and dedicated finishers, chief among them The Sentry, Golden Guardian, with Origin of Spider-Man and Angel of Serenity providing additional closing options. The deck favors a broad superhero roster and sustained battlefield play over combo speed, so it needs time to assemble its team and can lose momentum to repeated board wipes.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The land count sits above the norm of the format, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The white Hero creatures, equipment, and removal broadly pull toward a mono-white superhero team, but 13 mana rocks and three board wipes sit awkwardly beside a creature-pressure plan with few real finishers.
- PLAN theme_fit=partly: The commander is a white Avengers hero and the list is full of Hero creatures, but it includes Avengers Assemble! and Origin of the Avengers, which look like cards from the precon the person excluded, and the library-first constraint cannot be confirmed.
- PLAN useful_as_built=partly: The deck is castable with mono-white mana and plenty of cheap interaction, but 38 lands plus 13 rocks leaves too little action, and off-color lands like Command Tower and Chromatic Lantern add little.
- PLAN summary_honest=partly: placeholder
- [INFO] `curve_summary`: average mana value 2.97 over 61 nonland cards

<details><summary>The deck list</summary>

- 30 Plains | land | Provides the dependable white mana needed to deploy the commander and Heroes.
- 1 Command Tower | land | Supports the deck's white mana base.
- 1 Path of Ancestry | land | Adds a land slot suited to the creature-focused plan.
- 1 Secluded Courtyard | land | Supports the mana base for the Hero roster.
- 1 Minas Tirith | land | Adds a supporting land to the predominantly white mana base.
- 1 War Room | land | Provides another land slot for sustained development.
- 1 Rogue's Passage | land | Rounds out the supporting nonbasic lands.
- 1 Avengers Tower | land | Gives the superhero team a thematic headquarters in the mana base.
- 1 Windbrisk Heights | land | Completes the mana base for the creature-heavy plan.
- 1 Sol Ring | ramp | Accelerates development toward deploying several team members.
- 1 Arcane Signet | ramp | Adds mana support for the white-heavy spell suite.
- 1 Thought Vessel | ramp | Helps develop mana alongside the deck's draw package.
- 1 Springleaf Drum | ramp | Adds inexpensive acceleration to the creature-focused shell.
- 1 Wayfarer's Bauble | ramp | Provides another early mana-development option.
- 1 Sword of the Animist | ramp | Pairs an Equipment slot with the deck's ramp plan.
- 1 Inherited Envelope | ramp | Adds artifact-based mana support.
- 1 Commander's Sphere | ramp | Helps fund the commander and the supporting roster.
- 1 Chromatic Lantern | ramp | Provides another source of mana development.
- 1 Astral Cornucopia | ramp | Broadens the artifact ramp package.
- 1 White Auracite | ramp | Supports deployment of the more expensive finishers.
- 1 White Lotus Tile | ramp | Adds mana support for sustained board development.
- 1 Wickersmith's Tools | ramp | Completes the ramp package with another supporting artifact.
- 1 Agent 13, Sharon Carter | draw | Combines a Hero body with the deck's draw plan.
- 1 Agent Maria Hill | draw | Keeps the superhero theme present in the draw package.
- 1 Avengers Assemble! | draw | Provides thematic draw support for the team.
- 1 Origin of the Avengers | draw | Adds card flow while reinforcing the Avengers story.
- 1 The Vision | draw | Adds a recognizable Avenger to the draw suite.
- 1 Viv Vision, Teen Synthezoid | draw | Provides another Hero in the card-flow package.
- 1 Hero in Training | draw | Supports card flow without leaving the Hero theme.
- 1 Skullclamp | draw | Adds an Equipment-based draw option.
- 1 Mask of Memory | draw | Provides another Equipment slot dedicated to card flow.
- 1 Puresteel Paladin | draw | Adds a creature-based draw option alongside the Equipment package.
- 1 Inspiring Overseer | draw | Supplies supporting card flow on a creature.
- 1 Vanquisher's Banner | draw | Adds draw support suited to a themed creature roster.
- 1 Patriot, Shield Wielder | interaction | Keeps a Hero in the deck's interaction package.
- 1 Bastion Protector | interaction | Adds a supporting creature to the protection package.
- 1 Boromir, Warden of the Tower | interaction | Provides creature-based interaction alongside the superhero roster.
- 1 Clever Concealment | interaction | Adds instant-speed interaction to help preserve the board.
- 1 Swiftfoot Boots | interaction | Contributes Equipment-based protection for important creatures.
- 1 Lightning Greaves | interaction | Adds another protective Equipment option.
- 1 Unbreakable Formation | interaction | Supports the team with instant-speed interaction.
- 1 Take Up the Shield | interaction | Adds a shield-themed interaction spell.
- 1 Swords to Plowshares | removal | Provides an instant removal option for opposing threats.
- 1 Get Lost | removal | Adds another instant answer to troublesome permanents.
- 1 Generous Gift | removal | Broadens the removal package.
- 1 Stroke of Midnight | removal | Provides another removal spell for clearing obstacles.
- 1 Dispatch | removal | Adds removal alongside the substantial artifact package.
- 1 March of Otherworldly Light | removal | Gives the deck another instant removal option.
- 1 Banishing Light | removal | Adds enchantment-based removal.
- 1 Web Up | removal | Supplies superhero-themed removal.
- 1 Crib Swap | removal | Completes the removal suite with another instant answer.
- 1 Agent Phil Coulson | synergy | Adds a supporting Hero to the S.H.I.E.L.D. contingent.
- 1 Agent of Atlas | synergy | Reinforces the Human Spy Hero theme.
- 1 Agents of S.H.I.E.L.D. | synergy | Expands the supporting superhero team.
- 1 Avengers Quinjet | synergy | Adds a thematic Vehicle to the Avengers support package.
- 1 Night Nurse, Healer of Heroes | synergy | Provides another thematic supporting Hero.
- 1 Quake, Agent of S.H.I.E.L.D. | synergy | Strengthens the S.H.I.E.L.D. portion of the roster.
- 1 Silver Sable, Mercenary Leader | synergy | Rounds out the team-supporting Hero package.
- 1 Captain Mar-Vell, Space-Born | threat | Adds a superhero threat to the battlefield.
- 1 Invisible Woman, Sue Storm | threat | Provides another prominent Hero threat.
- 1 Luke Cage, Power Man | threat | Adds a frontline Hero to the combat plan.
- 1 Mockingbird, Ace Agent | threat | Provides an additional Hero threat from the spy contingent.
- 1 Okoye, Dora Milaje Leader | threat | Adds a warrior Hero to the attacking roster.
- 1 Roaming Throne | threat | Supplies a supporting creature threat.
- 1 The Sentry, Golden Guardian | wincon | Provides a superhero-themed finisher.
- 1 Origin of Spider-Man | wincon | Adds a superhero-themed finishing plan.
- 1 Angel of Serenity | wincon | Provides an additional finisher beyond the Hero roster.
- 1 Austere Command | wipe | Provides a board reset when the team falls behind.
- 1 Split Up | wipe | Adds another way to reset a crowded battlefield.
- 1 Dusk // Dawn | wipe | Completes the board-wipe package.

</details>

