# PR-8 deck gate

Run date: 2026-09-21. Card snapshot: 2026-09-04.

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
| Summaries that state a rule of the game | 2 |
| Summaries that state a FALSE rule | 0 |
| Judge errors | 0 |
| Case assertions missed (PR-28b) | 0 |
| Decks the plan judge read (PR-15, information) | 25 |
| Mean plan score, 0 to 1 | 0.74 |
| Plan reasons the judge left empty | 3 |
| Errors | 0 |
| Prompt version | 15 |
| Calls | 76 |
| Cost | $2.8263 |
| Time | 1906 seconds |

## Run

- Suite `decks`, run `pr8-deck-gate-run30`, on 2026-09-21, commit `2acf610`.
- Roles: generate on `gpt-5.6-terra` (openai, effort medium), judge on `claude-opus-5` (anthropic, effort medium), repair on `gpt-5.6-terra` (openai, effort medium).
- Versions: card snapshot 2026-09-04, generate prompt version 15, plan_rubric prompt version 3, precons `5.3.0+20260903`, quality_model `20260914T154223Z`.
- Calls: 76. Cost: $2.8263. Time: 1906 seconds.

## Findings by code

A block stops the deck. A warning and a note are reports.

| Code | Count |
|---|---|
| `curve_summary` | 25 |
| `not_owned` | 24 |
| `profile_off_band` | 10 |
| `mana_pass` | 6 |
| `outside_requested_set` | 2 |
| `bracket_cut` | 2 |
| `finisher_short` | 2 |
| `cards_trimmed` | 1 |
| `basics_added` | 1 |
| `precon_share` | 1 |

By severity: BLOCK 0. WARN 39. INFO 35. 

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

Grade: bad, score 0.02, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $805.25 to buy, $805.25 the whole deck.

**Summary:** Karlov of the Ghost Council turns the deck’s steady lifegain focus into growing pressure and creature control, while a mix of supporting creatures, enchantments, and interaction keeps the board manageable. The deck closes through Angel of Destiny, Approach of the Second Sun, or Test of Endurance, with a deep creature suite providing a more conventional combat path when needed. It gives up explosive acceleration for a steady, permanent-focused plan that develops its resources over several turns.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every card feeds or rewards lifegain — payoffs (Ajani's Pridemate, Voice of the Blessed, Vito, Sanguine Bond, Nykthos Paragon), draw-on-lifegain engines, lifelink/drain removal, and even the wraths (Fumigate, Kaya's Wrath, Meathook) gain or drain life, so the pieces point one way with only mild tension from three sweepers in a creature-dense list.
- PLAN theme_fit=yes: It is a W/B Karlov-led Commander lifegain deck as requested, with a casual-tier payoff shell and no tutored fast combo, consistent with bracket 3.
- PLAN useful_as_built=yes: 35 lands with solid duals/painlands for a mostly two-color curve topping around five, plus removal, sweepers, card draw, and multiple distinct win routes, make it playable out of the box despite modest ramp.
- PLAN summary_honest=yes: The named win conditions (Angel of Destiny, Approach of the Second Sun, Test of Endurance) are all present, the creature suite is indeed deep, and the admission of slow, artifact-only acceleration matches the ramp package of three-mana rocks.
- [INFO] `curve_summary`: average mana value 3.35 over 62 nonland cards
- [INFO] `bracket_cut`: to hold bracket 3, the builder cut 1 card: Enduring Angel // Angelic Enforcer (the combo Enduring Angel // Angelic Enforcer + Defiant Bloodlord), and added 1 basic land
- [INFO] `mana_pass`: the builder moved 2 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 1 Brightclimb Pathway // Grimclimb Pathway | land | Provides flexible black or white mana.
- 1 Caves of Koilos | land | Provides dependable black and white mana.
- 1 Command Tower | land | Provides either color for the deck.
- 1 Concealed Courtyard | land | Provides efficient black or white mana.
- 1 Godless Shrine | land | Provides both deck colors.
- 1 Scrubland | land | Provides both deck colors.
- 1 Vault of Champions | land | Provides black or white mana.
- 1 Shattered Sanctum | land | Provides black or white mana.
- 1 Mana Confluence | land | Provides flexible colored mana.
- 1 City of Brass | land | Provides flexible colored mana.
- 1 Tarnished Citadel | land | Provides flexible colored mana.
- 1 Silent Clearing | land | Provides black or white mana.
- 1 Marsh Flats | land | Finds an appropriate basic land or dual land.
- 1 Prismatic Vista | land | Finds an appropriate basic land.
- 1 Restless Fortress | land | Adds a black-white land with finisher utility.
- 1 High Market | land | Adds utility alongside the colored mana base.
- 10 Plains | land | Supplies a strong base of white mana.
- 11 Swamp | land | Supplies a strong base of black mana.
- 1 Archivist of Oghma | draw | Provides card-draw support.
- 1 Cosmos Elixir | draw | Provides card-draw support for the lifegain plan.
- 1 Dawn of Hope | draw | Provides card-draw support for the lifegain plan.
- 1 Enduring Innocence | draw | Provides ongoing card-draw support.
- 1 Exemplar of Light | draw | Provides card-draw support.
- 1 Mangara, the Diplomat | draw | Provides card-draw support.
- 1 Markov Purifier | draw | Provides card-draw support.
- 1 Sigarda's Splendor | draw | Provides card-draw support for the lifegain plan.
- 1 Survival Cache | draw | Provides card-draw support.
- 1 Well of Lost Dreams | draw | Provides card-draw support for the lifegain plan.
- 1 The Gaffer | draw | Provides card-draw support.
- 1 Alseid of Life's Bounty | interaction | Provides protective interaction.
- 1 Courageous Resolve | interaction | Provides protective interaction.
- 1 Faith's Shield | interaction | Provides protective interaction.
- 1 Metropolis Reformer | interaction | Provides defensive interaction.
- 1 Restoration Magic | interaction | Provides flexible interaction.
- 1 Sword of Light and Shadow | interaction | Provides interactive equipment support.
- 1 Werefox Bodyguard | interaction | Provides creature-based interaction.
- 1 Legion's Landing // Adanto, the First Fort | ramp | Provides low-cost ramp support.
- 1 Hot Dog Cart | ramp | Provides low-cost ramp support.
- 1 Misfortune Teller | ramp | Provides ramp support.
- 1 Oasis Gardener | ramp | Provides ramp support.
- 1 Nuka-Cola Vending Machine | ramp | Provides ramp support.
- 1 Orazca Relic | ramp | Provides ramp support.
- 1 Pristine Talisman | ramp | Provides ramp support for the lifegain plan.
- 1 Druidic Satchel | ramp | Provides ramp support.
- 1 Phial of Galadriel | ramp | Provides ramp support.
- 1 Carmen, Cruel Skymarcher | ramp | Provides ramp support on a relevant creature.
- 1 Ayli, Eternal Pilgrim | removal | Provides efficient creature removal.
- 1 Nightmare's Thirst | removal | Provides low-cost removal.
- 1 Umezawa's Jitte | removal | Provides repeatable removal utility.
- 1 Murderous Rider // Swift End | removal | Provides flexible removal.
- 1 Gumdrop Poisoner // Tempt with Treats | removal | Provides flexible removal.
- 1 Poisoner's Apprentice | removal | Provides low-cost removal.
- 1 Starseer Mentor | removal | Provides creature-based removal.
- 1 Consuming Corruption | removal | Provides direct removal.
- 1 Foolish Fate | removal | Provides direct removal.
- 1 Cleric Class | synergy | Supports the lifegain theme.
- 1 Ajani's Pridemate | synergy | Supports the lifegain theme.
- 1 Bloodthirsty Aerialist | synergy | Supports the lifegain theme.
- 1 Voice of the Blessed | synergy | Supports the lifegain theme.
- 1 Heliod, Sun-Crowned | synergy | Supports the lifegain theme.
- 1 Vito, Thorn of the Dusk Rose | synergy | Supports the lifegain theme.
- 1 Sanguine Bond | synergy | Supports the lifegain theme.
- 1 Archangel of Thune | threat | Provides a powerful lifegain-focused threat.
- 1 Attended Healer | threat | Provides a lifegain-focused board threat.
- 1 Celestine, the Living Saint | threat | Provides a resilient threat.
- 1 Cliffhaven Vampire | threat | Provides a lifegain-focused threat.
- 1 Defiant Bloodlord | threat | Provides a high-impact threat.
- 1 Divinity of Pride | threat | Provides a substantial lifegain-focused threat.
- 1 Enduring Tenacity | threat | Provides a durable threat.
- 1 Exalted Sunborn | threat | Provides an Angel threat.
- 1 Gideon's Company | threat | Provides a lifegain-focused threat.
- 1 Nykthos Paragon | threat | Provides a lifegain-focused threat.
- 1 Rhox Faithmender | threat | Provides a durable lifegain threat.
- 1 Righteous Valkyrie | threat | Provides an Angel-focused threat.
- 1 Angel of Destiny | wincon | Provides a finisher for the lifegain plan.
- 1 Approach of the Second Sun | wincon | Provides a dedicated finishing route.
- 1 Test of Endurance | wincon | Provides a dedicated lifegain finishing route.
- 1 Fumigate | wipe | Provides a reset against crowded boards.
- 1 Kaya's Wrath | wipe | Provides a reliable board reset.
- 1 The Meathook Massacre | wipe | Provides a flexible board reset.

</details>

### 2. aristocrats Commander, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 181 names.

Commander: Denethor, Ruling Steward.

Grade: typical, score 0.46, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $190.70 the whole deck.

**Summary:** This is a resilient Orzhov aristocrats deck that develops its mana and creature base, then turns repeated creature attrition into steady table pressure behind Denethor. Sacrifice-focused synergy pieces, token-friendly draw, and protective equipment keep the plan moving, while efficient removal and several sweepers buy time against stronger boards. Its finishers provide the closing push once opponents have been worn down. The deck gives up raw speed and a deep bench of dedicated threats in exchange for a stable mana base, broad answers, and strong play from the available library.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=no: The list pulls three ways at once — commander-protection equipment and pump/fog-style tricks (Bastion Protector, Champion's Helm, Take Up the Shield, Duty Beyond Death, Clever Concealment), a small toolbox of removal, and three symmetric board wipes — while carrying almost no sacrifice outlets, token generators, or death-trigger drains to tie any of it into a recurring sacrifice engine.
- PLAN theme_fit=no: The person asked for aristocrats, and while Denethor is an apt commander for it, the deck has essentially no sacrifice outlets, no Blood Artist-style drain payoffs, and no token production, delivering instead an Orzhov goodstuff/protect-the-commander pile.
- PLAN useful_as_built=partly: The mana base is clean and the curve and removal suite are playable, but with roughly fifteen creatures, three wipes that hit its own board, and only Archfiend of Ifnir and Vraska as real closers, it can hold the table without having a clear route to actually win.
- PLAN summary_honest=no: It advertises "sacrifice-focused synergy pieces" and "token-friendly draw" when the deck's sacrifice content amounts to Deadly Dispute, Nasty End and a couple of Gollums with no token engine behind Skullclamp or Idol of Oblivion, and it glosses over the equipment/protection subtheme that is actually the largest block of cards.
- [INFO] `curve_summary`: average mana value 2.70 over 63 nonland cards

<details><summary>The deck list</summary>

- 12 Plains | land | Basic white mana source for the two-color mana base.
- 12 Swamp | land | Basic black mana source for the two-color mana base.
- 1 Command Tower | land | Flexible colored land for a two-color commander deck.
- 1 City of Brass | land | Flexible colored land that supports both deck colors.
- 1 Exotic Orchard | land | Flexible colored land for the mana base.
- 1 Marsh Flats | land | Color-fixing land that can find a basic land.
- 1 Grand Coliseum | land | Additional colored source for the mana base.
- 1 Path of Ancestry | land | Colored land that fits the creature-heavy build.
- 1 Secluded Courtyard | land | Creature-focused color fixing for the mana base.
- 1 Unclaimed Territory | land | Creature-focused color fixing for the mana base.
- 1 Plaza of Heroes | land | Colored source that supports the legendary commander plan.
- 1 Spire of Industry | land | Colored source in an artifact-equipped deck.
- 1 Evolving Wilds | land | Color-fixing land that finds the needed basic.
- 1 Terramorphic Expanse | land | Additional basic-land color fixing.
- 1 Sol Ring | ramp | Efficient artifact ramp for early development.
- 1 Arcane Signet | ramp | Reliable two-color artifact ramp.
- 1 Fellwar Stone | ramp | Low-cost artifact ramp for developing quickly.
- 1 Springleaf Drum | ramp | Low-cost mana production alongside expendable creatures.
- 1 Wayfarer's Bauble | ramp | Early ramp that strengthens the land base.
- 1 Thought Vessel | ramp | Artifact ramp for steady resource development.
- 1 Sword of the Animist | ramp | Creature-based ramp that rewards attacking.
- 1 Relic of Legends | ramp | Artifact ramp that works with the legendary commander.
- 1 Commander's Sphere | ramp | Flexible mana rock for both deck colors.
- 1 Chromatic Lantern | ramp | Fixing and ramp for consistent casting.
- 1 Deadly Dispute | ramp | Listed ramp that also fits a sacrifice-focused game plan.
- 1 Skullclamp | draw | Efficient draw option for a creature-focused deck.
- 1 Idol of Oblivion | draw | Repeatable draw option for the token-oriented plan.
- 1 Tome of Legends | draw | Low-cost draw support for the commander deck.
- 1 Mask of Memory | draw | Combat-based card flow for the creature suite.
- 1 Lembas | draw | Low-cost artifact draw support.
- 1 Wall of Omens | draw | Early defensive creature that provides draw.
- 1 Inspiring Overseer | draw | Creature-based draw for the board-focused plan.
- 1 Night's Whisper | draw | Efficient black card draw.
- 1 Call of the Ring | draw | Ongoing draw support for the deck.
- 1 Puresteel Paladin | draw | Creature-based draw that complements the equipment package.
- 1 Folk Hero | draw | Draw support for the creature-heavy build.
- 1 Nasty End | draw | Black draw that fits the sacrifice-oriented strategy.
- 1 Painful Truths | draw | Efficient draw spell for refilling resources.
- 1 Buster Sword | draw | Artifact-based draw support for the creature plan.
- 1 Bastion Protector | interaction | Interaction that helps preserve the commander plan.
- 1 Boromir, Warden of the Tower | interaction | Creature-based interaction for protecting the board state.
- 1 Bronze Guardian | interaction | Artifact-friendly interaction for the permanent base.
- 1 Champion's Helm | interaction | Equipment interaction that supports the commander.
- 1 Clever Concealment | interaction | Protective interaction against opposing disruption.
- 1 Darksteel Plate | interaction | Protective equipment for an important creature.
- 1 Duty Beyond Death | interaction | Black protective interaction for the creature plan.
- 1 Gift of Immortality | interaction | Creature protection that fits attrition games.
- 1 Lightning Greaves | interaction | Low-cost protective equipment for key creatures.
- 1 Reprieve | interaction | Flexible white interaction against opposing plays.
- 1 Swiftfoot Boots | interaction | Additional protective equipment for the commander or finishers.
- 1 Take Up the Shield | interaction | Protective combat interaction for a key creature.
- 1 Bitter Triumph | removal | Flexible black removal for troublesome permanents.
- 1 Claim the Precious | removal | Black creature removal for opposing threats.
- 1 Fatal Push | removal | Efficient early removal.
- 1 Generous Gift | removal | Broad white permanent removal.
- 1 Get Lost | removal | Flexible white removal.
- 1 Infernal Grasp | removal | Direct black creature removal.
- 1 Path to Exile | removal | Efficient white creature removal.
- 1 Swords to Plowshares | removal | Efficient white creature removal.
- 1 Fiend Hunter | removal | Creature-based removal that supports the board plan.
- 1 Arcade Cabinet | synergy | Artifact synergy piece for the deck's permanent base.
- 1 Gollum the Abandoned | synergy | Black synergy creature for the sacrifice-focused plan.
- 1 Gollum, Patient Plotter | synergy | Synergy creature that supports an attrition game.
- 1 Gríma Wormtongue | synergy | Synergy creature for the deck's creature-centric plan.
- 1 Joo Dee, One of Many | synergy | Synergy creature for the board-focused strategy.
- 1 Nimble Hobbit | synergy | Low-cost synergy creature for the aristocrats shell.
- 1 Phantom Train | synergy | Artifact synergy piece for the deck's permanent package.
- 1 Bill the Pony | threat | Creature threat that adds pressure to the board.
- 1 Hei Bai, Spirit of Balance | threat | Creature threat for closing through board presence.
- 1 Namazu Trader | threat | Creature threat that contributes to the creature count.
- 1 Vengeful Villagers | threat | Creature threat for sustained board pressure.
- 1 Archfiend of Ifnir | wincon | Finisher-class win condition for ending longer games.
- 1 Grave Venerations | wincon | Finisher-class win condition for the sacrifice strategy.
- 1 Vraska, Betrayal's Sting | wincon | Finisher-class win condition that also carries the listed ramp role.
- 1 Austere Command | wipe | Flexible board wipe for resetting unfavorable boards.
- 1 Black Sun's Zenith | wipe | Black board wipe for clearing opposing creatures.
- 1 Dusk // Dawn | wipe | Board wipe that fits the creature-heavy build.

</details>

### 3. artifacts Commander, bracket 4

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 319 names.

Commander: Urza, Lord High Artificer.

Grade: good, score 0.67, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $7623.71 to buy, $7623.71 the whole deck.

**Summary:** Urza leads a fast, artifact-heavy blue deck that uses early mana and steady card advantage to deploy its key pieces while holding up dense interaction. The deck presses opponents with artifact threats, chief among them Kappa Cannoneer, and closes through Thassa's Oracle, Mechanized Production, or Mirrodin Besieged. It gives up a broad creature package for a focused artifact shell that rewards keeping a high artifact density on the table.

The quality model grades this deck good against the top lists of the format: the deck makes more mana on turn four than the norm, and that raises the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

- PLAN plan_coherent=partly: The bulk of the list is a coherent Urza artifact shell — fast mana, artifact-count payoffs, tutors like Whir/Transmute/Kuldotha Forgemaster, and counterspells — but the named Thassa's Oracle finish has no self-mill or library-emptying support, and a few cards (Escape Protocol, Unable to Scream, Waterbender Ascension) pull off the main line.
- PLAN theme_fit=yes: This is exactly the requested mono-blue Urza, Lord High Artificer artifact deck at a high power level, with Workshop, Moxen, LED, and free counterspells matching a bracket 4 build.
- PLAN useful_as_built=yes: Thirty-four lands plus abundant fast mana support a low curve in a single color, and even discounting the Oracle there are real closers in Kappa Cannoneer, Mechanized Production, Mirrodin Besieged, Aetherflux Reservoir, and a stack of artifact beaters.
- PLAN summary_honest=partly: It accurately describes the artifact density, fast mana, and counterspell suite, but it claims the deck 'closes through Thassa's Oracle' when there is no Demonic Consultation-style enabler or mill to make the Oracle a win rather than a small scry.
- [INFO] `curve_summary`: average mana value 2.98 over 65 nonland cards
- [INFO] `mana_pass`: the builder moved 5 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 2 Island | land | Provides a reliable blue source.
- 1 Island | land | Provides a reliable blue source.
- 1 Island | land | Provides a reliable blue source.
- 1 Island | land | Provides a reliable blue source.
- 1 Island | land | Provides a reliable blue source.
- 1 Island | land | Provides a reliable blue source.
- 1 Island | land | Provides a reliable blue source.
- 1 Island | land | Provides a reliable blue source.
- 1 Island | land | Provides a reliable blue source.
- 1 Island | land | Provides a reliable blue source.
- 1 Island | land | Provides a reliable blue source.
- 1 Island | land | Provides a reliable blue source.
- 1 Island | land | Provides a reliable blue source.
- 1 Island | land | Provides a reliable blue source.
- 1 Island | land | Provides a reliable blue source.
- 1 Island | land | Provides a reliable blue source.
- 1 Island | land | Provides a reliable blue source.
- 1 Island | land | Provides a reliable blue source.
- 1 Island | land | Provides a reliable blue source.
- 1 Island | land | Provides a reliable blue source.
- 1 Academy Ruins | land | Adds an artifact-focused utility land.
- 1 Buried Ruin | land | Adds an artifact-focused utility land.
- 1 Inventors' Fair | land | Adds an artifact-focused utility land.
- 1 Mishra's Workshop | land | Provides an artifact-focused mana land.
- 1 Urza's Workshop | land | Provides an artifact-focused mana land.
- 1 Seat of the Synod | land | Supplies blue mana while adding to the artifact count.
- 1 Mystic Sanctuary | land | Supplies a blue-producing utility land.
- 1 Otawara, Soaring City | land | Supplies blue mana with utility.
- 1 Glimmervoid | land | Provides flexible colored mana for the artifact shell.
- 1 Spire of Industry | land | Provides flexible colored mana for the artifact shell.
- 1 Mana Confluence | land | Provides dependable colored mana.
- 1 City of Brass | land | Provides dependable colored mana.
- 1 Command Tower | land | Provides a reliable blue source.
- 1 Chrome Mox | ramp | Provides early artifact acceleration.
- 1 Lion's Eye Diamond | ramp | Provides explosive artifact acceleration.
- 1 Mana Vault | ramp | Provides explosive artifact acceleration.
- 1 Mox Diamond | ramp | Provides early artifact acceleration.
- 1 Lotus Petal | ramp | Provides immediate artifact acceleration.
- 1 Mox Opal | ramp | Provides efficient artifact acceleration.
- 1 Sol Ring | ramp | Provides efficient artifact acceleration.
- 1 Moonsnare Prototype | ramp | Provides low-cost artifact acceleration.
- 1 Springleaf Drum | ramp | Provides low-cost artifact acceleration.
- 1 Metalworker | ramp | Provides artifact-focused mana production.
- 1 Krark-Clan Ironworks | ramp | Provides artifact-focused mana production.
- 1 Grand Architect | ramp | Provides artifact-focused mana production.
- 1 Karn, Legacy Reforged | ramp | Provides artifact-focused mana production.
- 1 Rhystic Study | draw | Provides persistent card advantage.
- 1 Thoughtcast | draw | Provides efficient artifact-based card draw.
- 1 Sai, Master Thopterist | draw | Provides artifact-focused card advantage.
- 1 Vedalken Archmage | draw | Provides artifact-focused card draw.
- 1 Riddlesmith | draw | Provides artifact-focused card selection.
- 1 Reverse Engineer | draw | Provides artifact-based card draw.
- 1 Thirst for Knowledge | draw | Provides efficient card draw.
- 1 One with the Machine | draw | Provides artifact-based card draw.
- 1 Forensic Gadgeteer | draw | Provides artifact-focused card advantage.
- 1 Nexus of Becoming | draw | Provides artifact-based card draw.
- 1 Tezzeret, Artifice Master | draw | Provides artifact-focused card advantage.
- 1 Force of Will | interaction | Provides premium stack interaction.
- 1 Fierce Guardianship | interaction | Provides premium protection and stack interaction.
- 1 An Offer You Can't Refuse | interaction | Provides efficient stack interaction.
- 1 Metallic Rebuke | interaction | Provides artifact-focused stack interaction.
- 1 Disruption Protocol | interaction | Provides artifact-focused stack interaction.
- 1 Ice Out | interaction | Provides efficient stack interaction.
- 1 Stoic Rebuttal | interaction | Provides artifact-focused stack interaction.
- 1 Padeem, Consul of Innovation | interaction | Provides artifact-focused protection.
- 1 Welding Jar | interaction | Provides low-cost artifact protection.
- 1 Aether Spellbomb | removal | Provides flexible artifact-based removal.
- 1 Aetherflux Reservoir | removal | Provides artifact-based removal.
- 1 Cyber Conversion | removal | Provides targeted removal.
- 1 Resculpt | removal | Provides flexible targeted removal.
- 1 Ravenform | removal | Provides targeted removal.
- 1 Kitesail Larcenist | removal | Provides creature-based removal.
- 1 Skysovereign, Consul Flagship | removal | Provides repeatable artifact-based removal.
- 1 Into Thin Air | removal | Provides targeted removal.
- 1 Unable to Scream | removal | Provides targeted removal.
- 1 Engineered Explosives | wipe | Provides a scalable artifact-based reset.
- 1 Hurkyl's Recall | wipe | Provides an artifact-focused reset.
- 1 Whir of Invention | synergy | Finds key artifacts while fitting the artifact plan.
- 1 Transmute Artifact | synergy | Finds key artifacts while fitting the artifact plan.
- 1 Mystic Forge | synergy | Supports the artifact-heavy plan.
- 1 Kappa Cannoneer | threat | Provides a major artifact-focused threat.
- 1 Master Transmuter | threat | Provides a high-impact artifact threat.
- 1 Kuldotha Forgemaster | threat | Provides a high-impact artifact threat.
- 1 Phyrexian Metamorph | threat | Provides a flexible artifact threat.
- 1 Cyberdrive Awakener | threat | Provides an artifact-focused threat.
- 1 Gearseeker Serpent | threat | Provides an artifact-focused threat.
- 1 Darksteel Juggernaut | threat | Provides an artifact-focused threat.
- 1 Lodestone Golem | threat | Provides an artifact creature threat.
- 1 Traxos, Scourge of Kroog | threat | Provides a large artifact threat.
- 1 Karn, Scion of Urza | threat | Provides an artifact-focused threat.
- 1 Myr Enforcer | threat | Provides an artifact creature threat.
- 1 Thassa's Oracle | wincon | Provides a compact finishing option.
- 1 Mechanized Production | wincon | Provides an artifact-focused finishing option.
- 1 Mirrodin Besieged | wincon | Provides an artifact-focused finishing option.
- 1 Mox Amber | ramp | the mana pass added it to bring the mana base inside the power level
- 1 Waterbender Ascension | draw | the mana pass added it to bring the mana base inside the power level
- 1 Fugitive Droid | interaction | the mana pass added it to bring the mana base inside the power level
- 1 Escape Protocol | interaction | the mana pass added it to bring the mana base inside the power level

</details>

### 4. dinosaur tribal, bracket 2

Format: Commander. Theme: dinosaurs. Pool: any_card. Shortlist: 299 names.

Commander: Gishath, Sun's Avatar.

Grade: bad, score 0.03, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $1432.19 to buy, $1432.19 the whole deck.

**Summary:** This deck develops its mana, establishes Dinosaur support pieces, and then deploys a steady stream of large creatures before bringing Gishath, Sun's Avatar into combat. It wins primarily through overwhelming Dinosaur attacks, backed by powerful finishers such as Bonehoard Dracosaur, Dinosaurs on a Spaceship, and Collective Inferno. The deck gives up some early speed and has a high mana curve, so protecting key creatures and using the ramp package to reach its larger spells are central to its plan.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The curve sits high for the format, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Every slot supports one Gishath dinosaur-tribal plan — ramp into big Dinosaurs, cost reducers and tribal payoffs, then connect with the commander for free dino hits — with only the two wipes (Blasphemous Act, Austere Command) sitting slightly against the board-building plan, and those are excused by the deck's oversized bodies and Wakening Sun's Avatar's one-sidedness.
- PLAN theme_fit=partly: It is exactly the Gishath dinosaur deck asked for, but the bracket 2 / new-player framing is strained by cards like Sylvan Library and Esper Sentinel plus a fetch-and-original-dual mana base that is both above the intended power tier and needlessly complex and costly for a beginner.
- PLAN useful_as_built=yes: Thirty-five lands with strong fixing plus a deep ramp package (Sol Ring, Arcane Signet, Birds, Drover, Topiary Stomper, Wayward Swordtooth) supports the admittedly top-heavy curve, and there are plenty of ways to close a game between Gishath triggers, Shared Animosity, Garruk's Uprising trample, and the large threat suite.
- PLAN summary_honest=yes: placeholder
- [INFO] `curve_summary`: average mana value 3.74 over 62 nonland cards
- [INFO] `basics_added`: the list was 1 card short, so the builder added 1 basic land

<details><summary>The deck list</summary>

- 9 Forest | land | Provides dependable green mana for the deck's early development.
- 5 Plains | land | Provides dependable white mana for the deck's support cards.
- 4 Mountain | land | Provides dependable red mana for Dinosaur spells.
- 1 Command Tower | land | Provides flexible mana for all three colors.
- 1 City of Brass | land | Provides flexible mana for all three colors.
- 1 Mana Confluence | land | Provides flexible mana for all three colors.
- 1 Cavern of Souls | land | Supports the Dinosaur creature base.
- 1 Unclaimed Territory | land | Provides tribal-focused color fixing.
- 1 Secluded Courtyard | land | Provides tribal-focused color fixing.
- 1 Temple Garden | land | Provides green and white mana.
- 1 Sacred Foundry | land | Provides red and white mana.
- 1 Stomping Ground | land | Provides red and green mana.
- 1 Savannah | land | Provides green and white mana.
- 1 Taiga | land | Provides red and green mana.
- 1 Plateau | land | Provides red and white mana.
- 1 Windswept Heath | land | Finds a needed green or white land.
- 1 Wooded Foothills | land | Finds a needed red or green land.
- 1 Arid Mesa | land | Finds a needed red or white land.
- 1 Prismatic Vista | land | Finds whichever basic land is needed.
- 1 Spire Garden | land | Provides red or green mana.
- 1 Bountiful Promenade | land | Provides green or white mana.
- 1 Spectator Seating | land | Provides red or white mana.
- 1 Sol Ring | ramp | Accelerates into the deck's costly Dinosaur spells.
- 1 Arcane Signet | ramp | Fixes colors while accelerating mana development.
- 1 Fellwar Stone | ramp | Provides early mana acceleration and color fixing.
- 1 Birds of Paradise | ramp | Provides early acceleration for any needed color.
- 1 Drover of the Mighty | ramp | Supports mana development for the Dinosaur plan.
- 1 Intrepid Paleontologist | ramp | Helps produce mana for Dinosaur spells.
- 1 Thunderherd Migration | ramp | Finds more mana to support the top end.
- 1 Wayward Swordtooth | ramp | Helps advance mana development while fitting the creature theme.
- 1 Topiary Stomper | ramp | Adds a substantial mana boost while remaining a Dinosaur.
- 1 Hulking Raptor | ramp | Provides extra mana from a Dinosaur body.
- 1 Esper Sentinel | draw | Provides early card advantage.
- 1 Sylvan Library | draw | Improves card selection through the game.
- 1 Garruk's Uprising | draw | Rewards the deck's large creatures with cards.
- 1 Beast Whisperer | draw | Turns creature-heavy development into cards.
- 1 Curious Altisaur | draw | Adds a Dinosaur that supports card flow.
- 1 Ripjaw Raptor | draw | Provides card advantage from a Dinosaur body.
- 1 Runic Armasaur | draw | Offers card advantage while contributing to the Dinosaur count.
- 1 Return of the Wildspeaker | draw | Refills the hand around the deck's large creatures.
- 1 Rishkar's Expertise | draw | Converts a large creature into a major refill.
- 1 Vanquisher's Banner | draw | Supports the tribe while providing ongoing cards.
- 1 For the Ancestors | draw | Provides a low-cost tribal card advantage option.
- 1 Heroic Intervention | interaction | Protects the board from opposing disruption.
- 1 Lightning Greaves | interaction | Protects a key creature and helps it get into combat.
- 1 Swiftfoot Boots | interaction | Protects an important creature while supporting attacks.
- 1 Temple Altisaur | interaction | Helps the Dinosaur board endure opposing pressure.
- 1 Steely Resolve | interaction | Protects the Dinosaur creature base.
- 1 Legolas's Quick Reflexes | interaction | Offers a flexible protective response for a creature.
- 1 Itzquinth, Firstborn of Gishath | removal | Provides efficient Dinosaur-based removal.
- 1 Savage Stomp | removal | Uses a Dinosaur to answer an opposing creature efficiently.
- 1 Needletooth Raptor | removal | Provides removal from a Dinosaur body.
- 1 Thrashing Brontodon | removal | Offers a flexible Dinosaur answer to troublesome permanents.
- 1 Raging Regisaur | removal | Provides repeatable removal pressure from a Dinosaur.
- 1 Ravenous Sailback | removal | Answers a problematic permanent while adding a Dinosaur.
- 1 Burning Sun's Avatar | removal | Provides a large Dinosaur that also answers threats.
- 1 Otepec Huntmaster | synergy | Supports efficient Dinosaur deployment.
- 1 Kinjalli's Caller | synergy | Makes the deck's Dinosaur spells easier to deploy.
- 1 Marauding Raptor | synergy | Supports the Dinosaur-focused creature plan.
- 1 Herald's Horn | synergy | Supports Dinosaur casting and tribal card flow.
- 1 Descendants' Path | synergy | Rewards the deck for maintaining its Dinosaur theme.
- 1 Shared Animosity | synergy | Turns a wide Dinosaur attack into greater combat pressure.
- 1 Icon of Ancestry | synergy | Supports the tribe while improving creature access.
- 1 Urza's Incubator | synergy | Reduces the cost of the Dinosaur-heavy creature base.
- 1 Realmwalker | synergy | Works with the tribal creature base to extend development.
- 1 Hunting Velociraptor | synergy | Contributes to the Dinosaur plan while helping deploy creatures.
- 1 Regisaur Alpha | threat | Provides a powerful Dinosaur threat that supports the board.
- 1 Pantlaza, Sun-Favored | threat | Provides a substantial Dinosaur threat.
- 1 Ghalta and Mavren | threat | Adds a large legendary Dinosaur threat.
- 1 Quartzwood Crasher | threat | Creates strong combat pressure as a Dinosaur threat.
- 1 Carnage Tyrant | threat | Provides a durable, high-impact Dinosaur threat.
- 1 Etali, Primal Storm | threat | Provides a major Dinosaur threat and a finisher.
- 1 Rampaging Ceratops | threat | Provides a large Dinosaur finisher for combat-focused games.
- 1 Shifting Ceratops | threat | Provides a versatile Dinosaur finisher.
- 1 Goring Ceratops | threat | Provides a large Dinosaur finisher for attacks.
- 1 Sun-Crested Pterodon | threat | Adds a reliable Dinosaur threat to the creature base.
- 1 Crested Herdcaller | threat | Builds combat pressure with a Dinosaur threat.
- 1 Roaming Throne | threat | Strengthens the deck's chosen tribal theme from a sizeable body.
- 1 Bonehoard Dracosaur | wincon | Provides a high-impact Dinosaur finisher.
- 1 Dinosaurs on a Spaceship | wincon | Provides a distinctive Dinosaur finisher.
- 1 Collective Inferno | wincon | Provides a tribal-themed finishing route.
- 1 Austere Command | wipe | Provides a flexible reset when the board gets out of hand.
- 1 Blasphemous Act | wipe | Provides an efficient creature-board reset.
- 1 Wakening Sun's Avatar | wipe | Provides a Dinosaur-based board reset.

</details>

### 5. blink Commander, owned first

Format: Commander. Theme: blink. Pool: owned_first. Shortlist: 169 names.

Commander: Gilraen, Dúnedain Protector.

Grade: baseline, score 0.31, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $180.65 the whole deck.

**Summary:** Gilraen leads a white blink deck that develops mana, protects its important creatures, and repeatedly turns creature-focused value into a durable board. The deck controls troublesome permanents with a deep removal suite, rebuilds with efficient draw and blink synergies, and closes through its Angels and artifact creatures, with Angel of Serenity chief among them. It gives up explosive multicolor options for a stable Plains-heavy mana base and a measured creature-centric game plan.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

This deck plan aims for 2 or more finishers, the cards that win the game, and the deck holds 1.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: There is a recognizable mono-white value core (protect the commander, remove threats, draw off artifacts), but the pieces pull three ways at once — a small blink package, a sizable equipment/artifact-draw package, and three symmetrical wraths that undercut the creature board the deck is supposed to build.
- PLAN theme_fit=partly: The request was specifically a blink deck, and while Gilraen and a few pieces (Flickerwisp, Angel of Condemnation, Fiend Hunter, Angel of Sanctions/Serenity, Wall of Omens) gesture at it, the bulk of the list is generic mono-white ramp, equipment and removal rather than a blink engine with repeatable flicker outlets.
- PLAN useful_as_built=yes: 36 lands plus ten mana rocks in a mono-white deck with a low curve means the spells reliably cast, the removal and draw are plentiful, and Angels plus equipped creatures give a real if unspectacular clock.
- PLAN summary_honest=partly: The removal suite, draw, mana base and Angel finishers are all really there and the finisher shortfall is disclosed, but 'blink synergies' oversells a handful of cards, and the summary never mentions the large equipment/artifact subtheme or the three board wipes that clash with its own creatures.
- [INFO] `curve_summary`: average mana value 3.00 over 63 nonland cards
- [WARN] `finisher_short`: the finisher count is 1, and this deck plan wants 2 or more (Angel of Serenity)

<details><summary>The deck list</summary>

- 1 Plains | land | Primary untapped white source for a consistent mana base.
- 1 Command Tower | land | Flexible colored land for casting the mono-white deck.
- 1 City of Brass | land | Colored fixing that remains available across the game.
- 1 Marsh Flats | land | Fetch land that can find Plains.
- 1 Fabled Passage | land | Fetch land that can find Plains.
- 1 Evolving Wilds | land | Fetch land that can find Plains.
- 1 Terramorphic Expanse | land | Fetch land that can find Plains.
- 1 Ash Barrens | land | Land slot that can help find Plains.
- 1 Plaza of Heroes | land | Utility land supporting the legendary commander.
- 27 Plains | land | Primary untapped white source for a consistent mana base.
- 1 Arcane Signet | ramp | Efficient mana development for deploying Gilraen and the creature suite.
- 1 Bender's Waterskin | ramp | Mana-development piece for the deck's midgame.
- 1 Fellwar Stone | ramp | Low-cost mana development.
- 1 Relic of Legends | ramp | Mana-development artifact that supports the commander-focused plan.
- 1 Ring of the Lucii | ramp | Mana-development artifact for the deck's creature-heavy plan.
- 1 Sol Ring | ramp | Efficient early mana development.
- 1 Thought Vessel | ramp | Mana development that also provides a useful utility effect.
- 1 Wayfarer's Bauble | ramp | Early mana development that supports reliable land access.
- 1 White Lotus Tile | ramp | Low-cost mana development.
- 1 Wickersmith's Tools | ramp | Mana-development artifact for reaching the deck's upper curve.
- 1 Adventurer's Airship | draw | Card-advantage piece for keeping resources flowing.
- 1 Buster Sword | draw | Equipment-based card-advantage piece.
- 1 Crown of Gondor | draw | Legendary equipment that supplies card advantage.
- 1 Diary of Dreams | draw | Card-advantage artifact for the longer game.
- 1 Energybending | draw | Flexible card-advantage spell.
- 1 Idol of Oblivion | draw | Low-cost repeatable card-advantage slot.
- 1 Lembas | draw | Compact artifact-based card advantage.
- 1 Mask of Memory | draw | Equipment-based card-advantage piece for the creature plan.
- 1 Mirror of Galadriel | draw | Legendary artifact that provides card advantage.
- 1 Skullclamp | draw | Efficient equipment-based card-advantage piece.
- 1 Tome of Legends | draw | Card advantage that works well beside the commander.
- 1 Banishing Light | removal | Broad permanent-answer slot.
- 1 Crib Swap | removal | Creature-answer slot that remains useful with the deck's creature theme.
- 1 Destroy Evil | removal | Flexible removal for problematic permanents.
- 1 Dispatch | removal | Cheap targeted removal.
- 1 Generous Gift | removal | Broad answer to a troublesome permanent.
- 1 Get Lost | removal | Efficient flexible removal.
- 1 Journey to Nowhere | removal | Creature-focused removal that fits the deck's white core.
- 1 Swords to Plowshares | removal | Efficient targeted creature removal.
- 1 Stroke of Midnight | removal | Flexible instant-speed permanent answer.
- 1 Ennis, Debate Moderator | synergy | Synergy creature supporting Gilraen's blink-focused plan.
- 1 Flickerwisp | synergy | Blink synergy that complements Gilraen's repeatable creature plan.
- 1 Personify | synergy | Synergy spell selected for the blink-centered strategy.
- 1 Angel of Condemnation | synergy | Creature synergy that gives the blink plan another useful body.
- 1 Fiend Hunter | synergy | Creature with useful overlap between blink value and board control.
- 1 Palace Jailer | synergy | Creature synergy that rewards repeatedly using creature-focused effects.
- 1 Wall of Omens | synergy | Low-cost creature value that is a strong blink target.
- 1 Clever Concealment | interaction | Protective interaction for preserving a developed board.
- 1 Champion's Helm | interaction | Protection equipment for Gilraen or an important creature.
- 1 Darksteel Plate | interaction | Durable protection for a key creature.
- 1 Duty Beyond Death | interaction | Protective interaction for important creatures.
- 1 Gift of Immortality | interaction | Protection piece for a creature the deck wants to keep using.
- 1 Lightning Greaves | interaction | Efficient protection and haste support for a key creature.
- 1 Slip On the Ring | interaction | Blink-oriented interaction that protects or resets a creature.
- 1 Swiftfoot Boots | interaction | Reliable protection for Gilraen or another key creature.
- 1 Austere Command | wipe | Flexible reset when the board demands a wipe.
- 1 Dusk // Dawn | wipe | Board-reset option with creature-focused utility.
- 1 Vanquish the Horde | wipe | Efficient creature-board reset.
- 1 Angel of Serenity | wincon | Primary high-impact closing card and the shortlist's marked finisher.
- 1 Angel of Sanctions | wincon | High-impact Angel that can help close games after the deck stabilizes.
- 1 Bronze Guardian | wincon | Substantial artifact creature that serves as a closing threat.
- 1 Aang, the Last Airbender | threat | Legendary creature that adds board pressure to the blink shell.
- 1 Bastion Protector | threat | Creature threat that adds meaningful board presence.
- 1 Boromir, Warden of the Tower | threat | Legendary creature that provides board pressure.
- 1 Champions of Minas Tirith | threat | Creature threat that supports a creature-heavy battlefield.
- 1 Exemplar of Light | threat | Angel creature that contributes to the deck's pressure plan.
- 1 Faramir, Field Commander | threat | Legendary creature that builds the deck's battlefield presence.
- 1 Frontline Medic | threat | Creature threat that helps establish a board.
- 1 Giada, Font of Hope | threat | Angel creature that adds pressure while fitting the deck's curve.
- 1 Inspiring Overseer | threat | Creature threat that supports the deck's steady battlefield plan.
- 1 Jocasta, Automaton Avenger | threat | Legendary artifact creature that adds a resilient board presence.
- 1 Kataki, War's Wage | threat | Creature threat that applies pressure while disrupting artifact-heavy boards.
- 1 Puresteel Paladin | threat | Creature threat that fits the deck's equipment package.

</details>

### 6. Modern tempo, tournament

Format: Modern. Theme: tempo. Pool: any_card. Shortlist: 167 names.

Grade: typical, score 0.47, model 20260914T154223Z.

Cards: 60 main, 15 sideboard. Repair turn: no. Block findings: 0.

Cost: $516.65 to buy, $516.65 the whole deck.

**Summary:** This blue-red tempo deck establishes pressure with Ragavan, Nimble Pilferer, Ledger Shredder, and Subtlety, then uses efficient draw, interaction, and removal to preserve that advantage. Temporal Mastery is chief among the synergy cards for pressing a lead, while the sideboard adds more disruption and flexible answers. The deck wins by keeping opponents off balance while its threats stay active, giving up the heavier top end and broader late-game plans of slower blue-red builds.

The quality model grades this deck typical against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The deck runs its spells as playsets, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade.

- PLAN plan_coherent=partly: Most of the list is a consistent cheap-threat-plus-cantrip-plus-interaction tempo shell, but a full playset of the seven-mana Temporal Mastery in a 24-land, low-curve deck pulls hard against that plan and is rarely castable on curve or reliably miracled with only six cantrips.
- PLAN theme_fit=partly: It is a Modern blue-red tempo deck with cantrips, counters, and burn as asked, but the requested Delver archetype's namesake one-drop flipper is absent and the deck instead leans on Ragavan/Ledger Shredder plus a bulky seven-drop.
- PLAN useful_as_built=partly: The mana base and curve support the cheap threats and interaction well enough to sit down and play, but four copies of a spell the deck can almost never cast amount to effectively four dead cards, and no sideboard is provided for the tournament the player named.
- PLAN summary_honest=no: It touts a sideboard of "more disruption and flexible answers" that does not exist in the list, and it elevates Temporal Mastery to the deck's chief synergy engine when the list has no miracle-enabling top-deck manipulation beyond a handful of cantrips.
- [INFO] `curve_summary`: average mana value 2.22 over 36 nonland cards

<details><summary>The deck list</summary>

- 6 Island | land | Basic blue land for the mana base.
- 5 Mountain | land | Basic red land for the mana base.
- 4 Steam Vents | land | Blue-red land slot that supports both deck colors.
- 4 Spirebluff Canal | land | Blue-red land slot for the tempo mana base.
- 4 Shivan Reef | land | Blue-red land slot that supports both deck colors.
- 1 Riverglide Pathway // Lavaglide Pathway | land | Flexible blue-red land slot.
- 4 Consider | draw | Efficient draw selection for a spell-heavy tempo plan.
- 2 Preordain | draw | Additional early draw selection and card flow.
- 4 Counterspell | interaction | Core interaction for protecting the tempo plan.
- 2 Spell Pierce | interaction | Low-cost interaction that supports an early lead.
- 4 Lightning Bolt | removal | Efficient removal for clearing opposing pressure.
- 4 Galvanic Discharge | removal | Additional removal that fits the blue-red shell.
- 4 Temporal Mastery | synergy | Synergy piece for converting a tempo advantage into more pressure.
- 4 Ragavan, Nimble Pilferer | threat | Early threat that applies pressure in the tempo plan.
- 4 Ledger Shredder | threat | Core evasive threat for maintaining pressure.
- 4 Subtlety | threat | Threat that also supports the deck's interactive posture.

</details>

### 7. Modern burn, casual

Format: Modern. Theme: burn. Pool: any_card. Shortlist: 290 names.

Grade: baseline, score 0.39, model 20260914T154223Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $23.34 to buy, $23.34 the whole deck.

**Summary:** This mono-red burn deck applies pressure with efficient removal, spell-focused synergy creatures, and a compact set of hard-hitting threats. It aims to reduce the opponent quickly through burn pressure while creatures maintain the attack, with Hazoret the Fervent chief among the larger threats. The deck gives up broad answers and elaborate multicolor options for a streamlined, consistent red plan.

The quality model grades this deck at the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade.

- PLAN plan_coherent=partly: Everything points at dealing damage, but the spell-count payoffs (Firebrand Archer, Thermo-Alchemist) only have 10 instants to trigger them while the rest of the list is four-drop creatures, and Hazoret wants an empty hand that this creature-heavy curve fights against.
- PLAN theme_fit=partly: It is mono-red and Modern-legal at a casual power level, but a burn deck with only four Bolts, two Lava Darts and two Risk Factors is really a red creature-aggro deck rather than the burn deck requested.
- PLAN useful_as_built=yes: 24 Mountains support an all-red curve topping at four, the creatures and reach spells give clear ways to close a game, and nothing in the list is uncastable or dead on arrival.
- PLAN summary_honest=partly: It correctly flags the synergy creatures and Hazoret, but calling six burn spells "efficient removal" that reduces the opponent "quickly through burn pressure" oversells a list whose damage actually comes from four-mana bodies.
- [INFO] `curve_summary`: average mana value 2.61 over 36 nonland cards

<details><summary>The deck list</summary>

- 24 Mountain | land | Provides the mono-red mana base.
- 4 Voldaren Epicure | draw | Fills a draw slot while staying within the red game plan.
- 2 Risk Factor | draw | Completes the draw package with a burn-focused card.
- 2 Flick a Coin | ramp | Supplies the requested ramp allocation.
- 4 Lightning Bolt | removal | Efficiently fills removal slots in the burn shell.
- 2 Lava Dart | removal | Completes the removal package with another low-cost red option.
- 4 Firebrand Archer | synergy | Supports the deck's spell-heavy burn plan.
- 4 Thermo-Alchemist | synergy | Provides another synergy piece for repeated burn pressure.
- 4 Akki Lavarunner // Tok-Tok, Volcano Born | threat | A red threat that contributes to the creature pressure plan.
- 4 Ashcloud Phoenix | threat | A resilient-looking threat slot for sustained pressure.
- 2 Hazoret the Fervent | threat | A powerful legendary red threat for the top end.
- 4 Sunspine Lynx | threat | Rounds out the threat package with additional red pressure.

</details>

### 8. Modern lifegain, FNM

Format: Modern. Theme: lifegain. Pool: any_card. Shortlist: 301 names.

Grade: bad, score 0.32, model 20260914T154223Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $435.76 to buy, $435.76 the whole deck.

**Summary:** This white-black lifegain deck establishes early momentum with Soul Warden and Ajani's Pridemate, then turns that foundation into pressure through creatures such as Enduring Tenacity, Twinblade Paladin, and Gideon's Company. Solitude and Murderous Rider keep opposing threats in check while Lembas and Dawn of Hope help sustain the deck into longer games. It wins by building a threatening board around its lifegain theme, giving up some raw speed for a more resilient creature-focused plan.

The quality model grades this deck below the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade.

- PLAN plan_coherent=yes: Every nonland slot either gains life (Soul Warden, Lembas, Twinblade Paladin, Gideon's Company, Legion's Landing, Dawn of Hope) or converts that life into pressure (Ajani's Pridemate, Enduring Tenacity, Sheoldred), with Solitude and Murderous Rider as the only interaction, so the pieces pull one way.
- PLAN theme_fit=yes: It is a 60-card Modern-legal white-black lifegain list with a straightforward creature plan, matching the requested colors, format, and roughly FNM-appropriate ambition even with a few premium cards like Solitude and Sheoldred.
- PLAN useful_as_built=yes: 24 lands with a heavy dual/fixing suite support a curve topping out at four, and the deck has plenty of redundant threats and two removal spells, so it plays fine out of the box.
- PLAN summary_honest=yes: placeholder
- [INFO] `curve_summary`: average mana value 3.00 over 36 nonland cards

<details><summary>The deck list</summary>

- 4 Godless Shrine | land | Provides both white and black mana for the deck’s core spells.
- 4 Caves of Koilos | land | Supplies flexible white or black mana early.
- 4 Shattered Sanctum | land | Adds reliable dual-color mana to the base.
- 2 Fetid Heath | land | Helps convert mana into the colors the deck needs.
- 2 Restless Fortress | land | Provides a white-black land slot with late-game threat potential.
- 4 Plains | land | Supports the deck’s white spells without entering tapped.
- 4 Swamp | land | Supports the deck’s black spells without entering tapped.
- 4 Lembas | draw | Adds card flow while fitting the deck’s lifegain plan.
- 2 Dawn of Hope | draw | Provides a repeatable source of cards for a lifegain-focused game plan.
- 2 Legion's Landing // Adanto, the First Fort | ramp | Helps develop mana resources while contributing to the board.
- 4 Solitude | removal | Premium removal that answers opposing threats while preserving the deck’s white count.
- 2 Murderous Rider // Swift End | removal | Flexible creature-based removal that also remains a threat afterward.
- 4 Soul Warden | synergy | A cheap lifegain synergy piece that starts the deck’s engine early.
- 4 Ajani's Pridemate | synergy | A central lifegain payoff that turns repeated life gain into board pressure.
- 4 Enduring Tenacity | threat | A durable midgame threat for the lifegain shell.
- 4 Twinblade Paladin | threat | An efficient threat that rewards the deck’s lifegain focus.
- 4 Gideon's Company | threat | A scalable creature threat for prolonged games.
- 2 Sheoldred, the Apocalypse | threat | A powerful top-end threat that pressures opponents and strengthens the black core.

</details>

### 9. Standard midrange, FNM

Format: Standard. Theme: midrange. Pool: any_card. Shortlist: 164 names.

Grade: bad, score 0.32, model 20260914T154223Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $207.62 to buy, $207.62 the whole deck.

**Summary:** This black-green midrange deck develops its mana early, uses efficient removal to keep the board manageable, and turns the corner with a dense creature package led by Goldvein Hydra and Vein Ripper. Insidious Roots, Vraska Joins Up, and Unholy Annex // Ritual Chamber give its permanent-heavy plan extra cohesion, while Darkstar Augur and Stocking the Pantry help it stay active in longer games. It gives up broad sideboard flexibility and relies on establishing its board to finish games.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The cards pair in ways the top lists do not, and that lowers the grade. The curve sits high for the format, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core of ramp, removal, and a fat creature curve pulls one way, but the 4 Insidious Roots have essentially no graveyard-recursion or reanimation to trigger them and Unholy Annex has no Demon other than its own Ritual Chamber side, so a chunk of the enchantment package sits outside the actual gameplan.
- PLAN theme_fit=yes: It is a 60-card Standard black-green deck with mana creatures, instant-speed removal, and midrange threats, which is exactly the archetype, format, and colors requested at a casual FNM level.
- PLAN useful_as_built=partly: 24 lands with 18 black sources, 16 green sources, and Llanowar Elves support the curve well enough to cast BBB Vein Ripper and the six-drops, but the high curve plus several dead or near-dead enchantment slots means a real number of draws do nothing.
- PLAN summary_honest=partly: The removal, ramp, and creature descriptions are accurate, but calling Insidious Roots and Unholy Annex sources of 'extra cohesion' overstates cards the list cannot reliably turn on, and the sideboard remark implies a sideboard the list does not show.
- [INFO] `curve_summary`: average mana value 3.17 over 36 nonland cards

<details><summary>The deck list</summary>

- 6 Forest | land | Green basic land for the mana base.
- 6 Swamp | land | Black basic land for the mana base.
- 4 Blooming Marsh | land | Black-green land that supports both main colors.
- 4 Overgrown Tomb | land | Black-green land that strengthens dual-color access.
- 4 Underground Mortuary | land | Black-green land for reliable color coverage.
- 2 Llanowar Elves | ramp | Early acceleration that helps deploy the midrange creature package.
- 4 Darkstar Augur | draw | Creature-based card advantage for the grindy game plan.
- 2 Stocking the Pantry | draw | Additional card advantage for longer games.
- 2 Bitter Triumph | removal | Efficient answer to opposing problems.
- 2 Withering Torment | removal | Flexible removal for troublesome permanents.
- 2 Maelstrom Pulse | removal | Versatile removal that handles a range of opposing cards.
- 4 Insidious Roots | synergy | Core support piece for the deck's creature-centered synergy package.
- 2 Vraska Joins Up | synergy | Supports the deck's black-green permanent plan.
- 2 Unholy Annex // Ritual Chamber | synergy | Adds another durable piece to the deck's synergy package.
- 4 Goldvein Hydra | threat | Early-to-midgame creature threat that rewards a proactive start.
- 4 Vein Ripper | threat | Powerful black creature threat for closing games.
- 2 Massacre Girl, Known Killer | threat | Creature threat that fits the deck's removal-heavy posture.
- 2 Rottenmouth Viper | threat | Black-green creature threat for applying pressure.
- 2 Sephiroth, Fabled SOLDIER // Sephiroth, One-Winged Angel | threat | Legendary creature threat that gives the deck another strong finisher.

</details>

### 10. Standard aggro, tournament

Format: Standard. Theme: aggro. Pool: any_card. Shortlist: 294 names.

Grade: baseline, score 0.43, model 20260914T154223Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $56.82 to buy, $56.82 the whole deck.

**Summary:** This red-white aggro deck establishes pressure with Diversion Specialist, Dragonback Lancer, and Redcap Gutter-Dweller, then keeps attacking while using Harsh Annotation and Case of the Gateway Express to remove resistance. Fugitive Codebreaker and Reckless Lackey help sustain the push, while Boros Charm and Parting Gust provide timely answers in close games. Slickshot Show-Off is chief among the cards that reinforce the proactive plan. The deck gives up long-game resilience in exchange for a streamlined, mana-consistent attack.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Every nonland card is either a pressure creature or cheap interaction to clear blockers, so the whole list pulls toward one attack-oriented plan even if the top end is heavier than a typical aggro curve.
- PLAN theme_fit=yes: It is a 60-card red-white creature-and-removal deck built for Standard play, exactly the color pair, format, and aggressive posture requested.
- PLAN useful_as_built=yes: 24 lands with a full playset of duals supports both colors, the curve tops at four, and there are 22 creatures plus eight removal spells, so it can be shuffled up and played as is, though the four-drop glut makes early turns clunkier than ideal.
- PLAN summary_honest=partly: The named cards and roles all exist in the list and the admission of poor long-game resilience is fair, but calling it a 'streamlined' attack glosses over twelve four-drops that make the deck slower than the summary implies.
- [INFO] `curve_summary`: average mana value 2.61 over 36 nonland cards

<details><summary>The deck list</summary>

- 8 Plains | land | Provides reliable white mana without entering tapped.
- 8 Mountain | land | Provides reliable red mana without entering tapped.
- 4 Sacred Foundry | land | Supplies both deck colors from a single land slot.
- 4 Inspiring Vantage | land | Adds flexible red-white mana for the aggressive curve.
- 4 Diversion Specialist | threat | A creature threat that supports early pressure.
- 4 Dragonback Lancer | threat | A creature threat for maintaining the attack.
- 4 Redcap Gutter-Dweller | threat | A creature threat that helps fill out the aggressive core.
- 4 Fugitive Codebreaker | draw | Provides card flow while contributing to the proactive plan.
- 2 Reckless Lackey | draw | Adds more card flow without moving away from aggression.
- 4 Boros Charm | interaction | Flexible instant-speed disruption for protecting pressure or answering opposing plays.
- 2 Parting Gust | interaction | Additional instant-speed interaction for contested turns.
- 4 Harsh Annotation | removal | Efficient removal to clear blockers and disrupt opposing threats.
- 4 Case of the Gateway Express | removal | Removal that helps the deck push its creatures through resistance.
- 4 Slickshot Show-Off | synergy | Supports the deck's aggressive creature plan and rewards proactive sequencing.

</details>

### 11. Commander with a locked card

Format: Commander. Theme: sacrifice. Pool: any_card. Shortlist: 299 names.

Commander: Karlov of the Ghost Council.

Grade: baseline, score 0.07, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $845.79 to buy, $845.79 the whole deck.

**Summary:** Karlov leads a sacrifice-focused Orzhov deck that develops disposable creatures, converts them into mana and cards, and turns those sacrifices into pressure and removal. Karlov grows from the deck’s life-gain elements and supplies a repeatable answer to opposing creatures, while the deck closes through its larger creature threats and dedicated finishing cards. It gives up some speed and consistency to maintain a broad creature engine with several expensive threats and board wipes.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. Few lands enter tapped, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Every slot pulls toward the same aristocrats loop: cheap disposable bodies and token makers (Teysa, Ratadrabik, Gisa, Spirit Bonds, Requiem Angel), a deep stack of free sac outlets (Carrion Feeder, Viscera Seer, Cartel Aristocrat, Ashnod's/Phyrexian Altar, Woe Strider), and payoffs that drain or draw on each death (Bastion of Remembrance, Vindictive Vampire, Elas il-Kor, Grave Pact, Dictate of Erebos, Meathook Massacre).
- PLAN theme_fit=yes: It is exactly the requested Orzhov Karlov sacrifice Commander deck, with the named commander, WB identity, and a sacrifice engine as the core plan.
- PLAN useful_as_built=yes: 37 lands with strong Orzhov fixing and mostly untapped duals, a low curve with plenty of one- and two-drops, and multiple credible win routes (aristocrat drain loops, Razaketh tutoring, Dictate/Grave Pact lockouts, and token beatdown) make it playable off the shelf.
- PLAN summary_honest=yes: The claims map to real cards — altars and Deadly Dispute/Village Rites for mana and cards, Attrition/Ayli/Yawgmoth/Grave Pact as repeatable removal, three sweepers, and the drain payoffs that feed Karlov's lifegain trigger — and it openly concedes the expensive threats and consistency cost.
- [INFO] `curve_summary`: average mana value 3.19 over 63 nonland cards

<details><summary>The deck list</summary>

- 10 Plains | land | Provides white mana for the deck.
- 11 Swamp | land | Provides black mana for the deck.
- 1 Brightclimb Pathway // Grimclimb Pathway | land | Provides flexible white or black mana.
- 1 Caves of Koilos | land | Provides both commander colors.
- 1 City of Brass | land | Provides flexible colored mana.
- 1 Command Tower | land | Provides both commander colors.
- 1 Concealed Courtyard | land | Provides both commander colors.
- 1 Eclipsed Steppe | land | Provides white and black mana.
- 1 Godless Shrine | land | Provides both commander colors.
- 1 High Market | land | Provides a land-based sacrifice outlet.
- 1 Isolated Chapel | land | Provides both commander colors.
- 1 Mana Confluence | land | Provides flexible colored mana.
- 1 Marsh Flats | land | Supports access to both commander colors.
- 1 Phyrexian Tower | land | Provides a land-based sacrifice outlet and mana.
- 1 Scrubland | land | Provides both commander colors.
- 1 Shattered Sanctum | land | Provides both commander colors.
- 1 Vault of Champions | land | Provides both commander colors.
- 1 Sol Ring | ramp | Provides efficient mana acceleration and is kept as requested.
- 1 Ashnod's Altar | ramp | Turns sacrificed creatures into mana.
- 1 Phyrexian Altar | ramp | Turns sacrificed creatures into colored mana.
- 1 Pitiless Plunderer | ramp | Supports sacrifice-based mana production.
- 1 Priest of Forgotten Gods | ramp | Supports the sacrifice plan while accelerating mana.
- 1 Warren Soultrader | ramp | Provides sacrifice-based mana acceleration.
- 1 Skullport Merchant | ramp | Supports sacrifice-based mana production.
- 1 Master of Dark Rites | ramp | Provides mana acceleration for the sacrifice plan.
- 1 Culling the Weak | ramp | Converts a creature into a burst of mana.
- 1 Deadly Dispute | ramp | Converts a disposable permanent into mana and resources.
- 1 Corrupted Conviction | draw | Turns a sacrificed creature into cards.
- 1 Village Rites | draw | Turns a sacrificed creature into cards.
- 1 Vampiric Rites | draw | Provides repeatable sacrifice-based card flow.
- 1 Baron Bertram Graywater | draw | Provides card flow for the sacrifice plan.
- 1 Disciple of Bolas | draw | Turns a creature into card flow.
- 1 Shadowheart, Dark Justiciar | draw | Provides sacrifice-based card flow.
- 1 Smothering Abomination | draw | Rewards the deck's sacrifice plan with cards.
- 1 Tevesh Szat, Doom of Fools | draw | Provides card flow and bodies for the sacrifice plan.
- 1 Lord Skitter's Butcher | draw | Provides card flow for the sacrifice plan.
- 1 Thraxodemon | draw | Provides sacrifice-based card flow.
- 1 Tribute to Horobi // Echo of Death's Wail | draw | Provides card flow for the deck.
- 1 Cartel Aristocrat | interaction | Provides a sacrifice outlet that protects itself.
- 1 Dark Privilege | interaction | Provides protective interaction through sacrifice.
- 1 Fanatical Devotion | interaction | Provides protective sacrifice-based interaction.
- 1 Flare of Fortitude | interaction | Protects the deck's board presence.
- 1 Gift of Doom | interaction | Provides protective interaction for a key permanent.
- 1 Promise of Tomorrow | interaction | Provides resilience against opposing interaction.
- 1 Spirit Bonds | interaction | Provides protective interaction and supports the creature plan.
- 1 Eldrazi Monument | interaction | Protects the board while rewarding disposable creatures.
- 1 Attrition | removal | Uses sacrificed creatures as repeatable creature removal.
- 1 Ayli, Eternal Pilgrim | removal | Provides creature removal within the sacrifice plan.
- 1 Bone Shards | removal | Converts a disposable creature into efficient removal.
- 1 Dictate of Erebos | removal | Turns the deck's sacrifices into opposing creature removal.
- 1 Eaten Alive | removal | Uses a disposable creature to remove a creature.
- 1 Grave Pact | removal | Turns the deck's sacrifices into opposing creature removal.
- 1 Spark Harvest | removal | Uses a disposable creature to remove a permanent.
- 1 Teysa, Orzhov Scion | removal | Provides creature removal within the sacrifice plan.
- 1 Yawgmoth, Thran Physician | removal | Provides repeatable creature removal through sacrifice.
- 1 Carrion Feeder | synergy | Provides a low-cost repeatable sacrifice outlet.
- 1 Bartolomé del Presidio | synergy | Provides a repeatable sacrifice outlet.
- 1 Elas il-Kor, Sadistic Pilgrim | synergy | Rewards the deck's sacrifice plan.
- 1 Fleshtaker | synergy | Rewards sacrificing creatures and artifacts.
- 1 Viscera Seer | synergy | Provides a low-cost repeatable sacrifice outlet.
- 1 Woe Strider | synergy | Provides a sacrifice outlet for the deck.
- 1 Bastion of Remembrance | synergy | Rewards creatures leaving the battlefield.
- 1 Abhorrent Overlord | threat | Provides a substantial finishing threat.
- 1 Corpse Harvester | threat | Provides a threatening creature for the sacrifice plan.
- 1 Felisa, Fang of Silverquill | threat | Builds on Karlov's counter-focused presence.
- 1 Ghoulcaller Gisa | threat | Provides a threatening creature for the sacrifice plan.
- 1 Hell's Caretaker | threat | Provides a threatening sacrifice-based creature.
- 1 Liesa, Forgotten Archangel | threat | Provides a powerful creature threat.
- 1 Mondrak, Glory Dominus | threat | Provides a powerful creature threat for the deck.
- 1 Morkrut Necropod | threat | Provides a substantial finishing threat.
- 1 Prowling Geistcatcher | threat | Provides a resilient creature threat.
- 1 Ratadrabik of Urborg | threat | Provides a powerful legendary creature threat.
- 1 Requiem Angel | threat | Provides a substantial finishing threat.
- 1 Vindictive Vampire | threat | Provides a finishing threat for the sacrifice plan.
- 1 Razaketh, the Foulblooded | wincon | Provides a dedicated finishing card for the deck.
- 1 Relic Vial | wincon | Provides a dedicated finishing card for the sacrifice plan.
- 1 Sephiroth, Fabled SOLDIER // Sephiroth, One-Winged Angel | wincon | Provides a dedicated finishing threat.
- 1 Austere Command | wipe | Provides flexible board-clearing interaction.
- 1 The Meathook Massacre | wipe | Provides a board wipe that supports the sacrifice plan.
- 1 Toxic Deluge | wipe | Provides efficient board-clearing interaction.

</details>

### 12. Commander on a budget

Format: Commander. Theme: tokens. Pool: any_card. Shortlist: 298 names.

Commander: Adeline, Resplendent Cathar.

Grade: bad, score 0.02, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $76.94 to buy, $76.94 the whole deck.

**Summary:** Adeline turns each attack into more Human tokens, and this deck builds on that pressure with token-focused permanents, board-building threats, and steady card flow. It aims to win by going wide in combat or by using its dedicated finishing cards once the board is established, while keeping enough removal and resets to recover from opposing development. It gives up premium mana, expensive token staples, and high-speed explosiveness in exchange for a straightforward, durable white token plan.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The deck makes less mana on turn four than the norm, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- JUDGE [true]: "Adeline turns each attack into more Human tokens". Adeline, Resplendent Cathar's ability creates a 1/1 white Human creature token for each opponent whenever you attack, so the statement about what the card does is accurate.
- PLAN plan_coherent=partly: The core is a consistent mono-white go-wide shell (anthems, token makers, Divine Visitation, Hour of Reckoning, Adeline attack triggers), but a cluster of eight weak Treasure/utility artifacts labeled "ramp" (Bucknard's Everfull Purse, Coin of Mastery, Collector's Vault, Currency Converter, Goldvein Pick, Karn) neither ramps meaningfully nor makes tokens and pulls slots away from the stated plan.
- PLAN theme_fit=yes: It is a mono-white Adeline token Commander deck built from cheap, mostly bulk-priced cards with no infinite combos or fast mana, which matches the budget and bracket 2 request.
- PLAN useful_as_built=yes: 38 mana sources of a single color support a curve topping around six, and the deck has clear paths to a win via wide attacks, anthems, and a few finishers, though its ramp package is largely inert and the clock is slow.
- PLAN summary_honest=yes: Every claim — token permanents, card flow, removal, wipes, dedicated finishers (Halo Fountain, Sword of Body and Mind), and the admitted lack of premium mana and explosiveness — is backed by cards actually in the list.
- [INFO] `curve_summary`: average mana value 3.24 over 62 nonland cards

<details><summary>The deck list</summary>

- 37 Plains | land | Provides reliable white mana for Adeline and the deck’s white spells.
- 1 Bucknard's Everfull Purse | ramp | Adds to the deck’s mana-development package.
- 1 Coin of Mastery | ramp | Adds to the deck’s mana-development package.
- 1 Collector's Vault | ramp | Adds to the deck’s mana-development package.
- 1 Currency Converter | ramp | A low-cost piece of the mana-development package.
- 1 Druidic Satchel | ramp | Adds to the deck’s mana-development package.
- 1 Goldvein Pick | ramp | Adds to the deck’s mana-development package.
- 1 Idol of False Gods | ramp | Adds to the deck’s mana-development package.
- 1 Karn, Living Legacy | ramp | Adds to the deck’s mana-development package.
- 1 Keeper of the Accord | ramp | Adds to the deck’s mana-development package.
- 1 Legion's Landing // Adanto, the First Fort | ramp | Adds to the deck’s mana-development package.
- 1 Angelic Sell-Sword | draw | A draw option that helps keep resources flowing.
- 1 Bygone Bishop | draw | A draw option that helps keep resources flowing.
- 1 Dawn of Hope | draw | A draw option that helps keep resources flowing.
- 1 Faramir, Field Commander | draw | A draw option that helps keep resources flowing.
- 1 Glimmer Seeker | draw | A draw option that helps keep resources flowing.
- 1 Idol of Oblivion | draw | A draw option that helps keep resources flowing.
- 1 Platoon Dispenser | draw | A draw option that helps keep resources flowing.
- 1 Sarah Jane Smith | draw | A draw option that helps keep resources flowing.
- 1 Search the Premises | draw | A draw option that helps keep resources flowing.
- 1 Staff of the Storyteller | draw | A draw option that helps keep resources flowing.
- 1 Wedding Announcement // Wedding Festivity | draw | A draw option that helps keep resources flowing.
- 1 Basri Ket | interaction | Provides interaction while fitting the token-focused plan.
- 1 Basri, Tomorrow's Champion | interaction | Provides interaction while fitting the token-focused plan.
- 1 Lena, Selfless Champion | interaction | Provides interaction while fitting the token-focused plan.
- 1 Mage's Attendant | interaction | Provides interaction while fitting the token-focused plan.
- 1 Rootborn Defenses | interaction | Provides interaction for the developing board.
- 1 Squad Commander | interaction | Provides interaction while fitting the token-focused plan.
- 1 Aerial Assault | removal | Efficient removal for problematic opposing pieces.
- 1 Banishing Slash | removal | Efficient removal for problematic opposing pieces.
- 1 Battle Menu | removal | Efficient removal for problematic opposing pieces.
- 1 Generous Gift | removal | Flexible removal for problematic opposing pieces.
- 1 Kellan's Lightblades | removal | Efficient removal for problematic opposing pieces.
- 1 Skyclave Apparition | removal | Creature-based removal that supports the board-focused plan.
- 1 Stroke of Midnight | removal | Flexible removal for problematic opposing pieces.
- 1 Aligned Heart | synergy | Supports the deck’s token-oriented synergies.
- 1 Anointer Priest | synergy | Supports the deck’s token-oriented synergies.
- 1 Animation Module | synergy | Supports the deck’s token-oriented synergies.
- 1 Automated Assembly Line | synergy | Supports the deck’s token-oriented synergies.
- 1 Cat Collector | synergy | Supports the deck’s token-oriented synergies.
- 1 Clarion Spirit | synergy | Supports the deck’s token-oriented synergies.
- 1 Divine Visitation | synergy | A high-impact token synergy piece.
- 1 Felidar Retreat | synergy | Supports the deck’s token-oriented synergies.
- 1 Intangible Virtue | synergy | Supports the deck’s token-oriented synergies.
- 1 Horn of Gondor | synergy | Supports the deck’s token-oriented synergies.
- 1 Ajani's Chosen | threat | A board-building threat for the token plan.
- 1 Archon of Sun's Grace | threat | A board-building threat for the token plan.
- 1 Attended Healer | threat | A board-building threat for the token plan.
- 1 Basri's Lieutenant | threat | A board-building threat for the token plan.
- 1 Cemetery Protector | threat | A resilient threat for the creature-heavy plan.
- 1 Defiler of Faith | threat | A board-building threat for the token plan.
- 1 Emeria Angel | threat | A board-building threat for the token plan.
- 1 Gideon, Ally of Zendikar | threat | A threat that supports the token-oriented game plan.
- 1 God-Eternal Oketra | threat | A durable threat for the creature-heavy plan.
- 1 Hero of Bladehold | threat | An attacking threat that fits the go-wide plan.
- 1 Nahiri, the Lithomancer | threat | A finisher-class threat that adds a different angle of pressure.
- 1 Requiem Angel | threat | A finisher-class threat for the token-oriented plan.
- 1 Halo Fountain | wincon | A finisher that gives the wide board a dedicated closing route.
- 1 Luck Bobblehead | wincon | A finisher that gives the deck a separate closing route.
- 1 Sword of Body and Mind | wincon | A finisher that rewards connecting with the deck’s attackers.
- 1 Ceaseless Conflict | wipe | A reset button for boards that get out of hand.
- 1 Hour of Reckoning | wipe | A reset button suited to a creature-focused deck.
- 1 Martial Coup | wipe | A reset button that remains aligned with the token plan.

</details>

### 13. owned first, and the commander is not owned

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 230 names.

Commander: Karlov of the Ghost Council.

Grade: baseline, score 0.04, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $150.82 the whole deck.

**Summary:** Karlov is the focal point of a white-black lifegain and counters strategy: establish mana early, deploy lifegain-focused creatures and equipment, then turn Karlov into both a growing combat threat and a source of creature control. The deck backs that pressure with broad targeted answers, protective spells, and several reset buttons, before closing through its larger creatures or dedicated finishers. It gives up some raw speed for a more board-centric plan that needs its creatures and commander to stay in play.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves one color far better than another, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core of lifegain creatures, cheap equipment and Karlov-boosting payoffs holds together, but a sizable block of generic black/white beaters (Canyon Crawler, Rooftop Percher, Foggy Swamp Hunters, Sneering Shadewriter) contributes little to lifegain, and three symmetric board wipes pull against a plan the summary itself calls creature-dependent.
- PLAN theme_fit=yes: It is a WB Commander deck led by Karlov of the Ghost Council built around gaining life, with an interaction and ramp package appropriate to a mid-power bracket 3 table.
- PLAN useful_as_built=yes: 36 lands plus Sol Ring and several two-mana rocks support a curve that tops out reasonably, and there is ample removal and multiple ways to close, though the 17-Plains/10-Swamp split makes the black double-costs occasionally awkward.
- PLAN summary_honest=partly: It fairly discloses the wipes, the board-centric plan, and the lopsided mana, but calling it a 'counters strategy' overstates a deck where Karlov is nearly the only counters payoff, and 'lifegain-focused creatures' glosses over the handful of off-theme bodies.
- [WARN] `not_owned`: Karlov of the Ghost Council: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 2.98 over 63 nonland cards

<details><summary>The deck list</summary>

- 17 Plains | land | Provides a dependable white mana base.
- 10 Swamp | land | Provides a dependable black mana base.
- 1 Command Tower | land | Supports the two-color mana base.
- 1 City of Brass | land | Supports access to both deck colors.
- 1 Exotic Orchard | land | Adds flexible mana support.
- 1 Marsh Flats | land | Helps stabilize the land base.
- 1 Fabled Passage | land | Helps stabilize the land base.
- 1 Spire of Industry | land | Fits the artifact-heavy mana package.
- 1 Plaza of Heroes | land | Supports the legendary commander plan.
- 1 Evolving Wilds | land | Helps stabilize the land base.
- 1 Escape Tunnel | land | Rounds out the land base.
- 1 Call of the Ring | draw | Provides a dedicated card-advantage slot.
- 1 Exemplar of Light | draw | Adds card advantage while fitting the creature plan.
- 1 Hero in Training | draw | Supplies card advantage from a low-cost creature slot.
- 1 Idol of Oblivion | draw | Provides repeatable card-advantage support.
- 1 Inspiring Overseer | draw | Contributes a creature body and card advantage.
- 1 Lembas | draw | Provides a compact card-advantage piece.
- 1 Mask of Memory | draw | Turns combat into card advantage.
- 1 Night's Whisper | draw | Provides efficient card advantage.
- 1 Puresteel Paladin | draw | Rewards the deck's equipment package with card advantage.
- 1 Skullclamp | draw | Provides efficient card-advantage support.
- 1 Wall of Omens | draw | Adds early defense alongside card advantage.
- 1 Bastion Protector | interaction | Helps keep Karlov on the table.
- 1 Champion's Helm | interaction | Protects the commander while supporting the equipment package.
- 1 Clever Concealment | interaction | Guards the board against opposing disruption.
- 1 Darksteel Plate | interaction | Provides durable protection for Karlov or a major threat.
- 1 Lightning Greaves | interaction | Protects Karlov and lets key creatures act quickly.
- 1 Swiftfoot Boots | interaction | Adds a second efficient protection piece.
- 1 Reprieve | interaction | Offers flexible stack interaction while keeping cards flowing.
- 1 Unbreakable Formation | interaction | Protects a developed creature board.
- 1 Sol Ring | ramp | Accelerates early mana development.
- 1 Arcane Signet | ramp | Provides reliable two-color acceleration.
- 1 Fellwar Stone | ramp | Adds efficient mana acceleration.
- 1 Wayfarer's Bauble | ramp | Develops mana while improving land access.
- 1 Springleaf Drum | ramp | Provides low-cost mana acceleration.
- 1 Sword of the Animist | ramp | Combines the equipment plan with ongoing mana development.
- 1 Giada, Font of Hope | ramp | Accelerates the deck's Angel threats.
- 1 Lotho, Corrupt Shirriff | ramp | Provides mana support from a creature slot.
- 1 Deadly Dispute | ramp | Converts a spare permanent into mana development.
- 1 Commander's Sphere | ramp | Provides reliable midgame mana support.
- 1 Bitter Triumph | removal | Answers a troublesome opposing permanent efficiently.
- 1 Dismember | removal | Provides low-cost creature removal.
- 1 Dispatch | removal | Adds efficient targeted removal.
- 1 Fatal Push | removal | Provides early targeted creature removal.
- 1 Generous Gift | removal | Answers a wide range of opposing permanents.
- 1 Get Lost | removal | Provides flexible permanent removal.
- 1 Infernal Grasp | removal | Offers clean creature removal.
- 1 Swords to Plowshares | removal | Provides efficient creature removal.
- 1 Stroke of Midnight | removal | Adds another flexible permanent answer.
- 1 Aerith Gainsborough | synergy | Supports the deck's lifegain-focused plan.
- 1 Angel of Vitality | synergy | Rewards the deck for gaining life.
- 1 Compassionate Healer | synergy | Adds another lifegain-focused creature.
- 1 Kor Firewalker | synergy | Contributes to the lifegain theme.
- 1 Light of Promise | synergy | Supports Karlov's counter-based game plan.
- 1 Night Nurse, Healer of Heroes | synergy | Fits the deck's healing-focused synergies.
- 1 Rosie Cotton of South Lane | synergy | Rewards the deck's creature and counter plan.
- 1 Angel of Invention | threat | Provides an evasive creature threat.
- 1 Canyon Crawler | threat | Adds a substantial creature threat.
- 1 Dawnhand Eulogist | threat | Provides a high-impact creature threat.
- 1 Foggy Swamp Hunters | threat | Adds pressure through the creature suite.
- 1 Gollum, Silent Slinker // Meager Meal | threat | Provides a resilient threat with added flexibility.
- 1 Minwu, White Mage | threat | Adds a legendary creature threat.
- 1 Rabaroo Troop | threat | Provides another creature that can pressure opponents.
- 1 Reaping Willow | threat | Adds a durable creature threat.
- 1 Rooftop Percher | threat | Adds a creature threat to the board plan.
- 1 Shattered Angel | threat | Provides an evasive threat.
- 1 Sneering Shadewriter | threat | Adds a potent finishing creature threat.
- 1 Victory's Herald | threat | Strengthens combat-focused closing turns.
- 1 Frodo, Sauron's Bane | wincon | Provides a dedicated finishing route.
- 1 Lyra Dawnbringer | wincon | Provides a powerful evasive finishing threat.
- 1 Grave Venerations | wincon | Offers a dedicated route to close the game.
- 1 Austere Command | wipe | Provides a flexible reset when the board gets out of hand.
- 1 Fumigate | wipe | Resets opposing creature boards while supporting the life plan.
- 1 Vanquish the Horde | wipe | Provides an efficient creature-board reset.

</details>

### 14. the user delegates the commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 299 names.

Commander: Atarka, World Render.

Grade: bad, score 0.05, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $753.72 to buy, $753.72 the whole deck.

**Summary:** Atarka, World Render leads a red-green Dragon deck that builds its mana, deploys cost-reduced Dragons, and turns each attack into a severe double-strike combat step. The deck wins primarily by developing a critical mass of evasive attackers, with its largest Dragons and dedicated Dragon payoffs providing decisive finishing turns. It gives up some early-game speed for expensive threats, so protecting key Dragons and choosing the right moment to commit to the board are central to its game plan.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every card supports one plan — cast Dragons cheaply (Dragonspeaker Shaman, Urza's Incubator, Herald's Horn, Orbs), then swing with Atarka's double strike and Dragon payoffs like Scourge of Valkas, Terror of the Peaks, Dragon Tempest, and Shared Animosity — with only a couple of loose picks (Tibalt's Trickery, Skullclamp) pulling slightly off-axis.
- PLAN theme_fit=yes: It is a Commander Dragon tribal deck with a builder-chosen Dragon commander at a bracket-3-appropriate power level: no fast mana, no tutor-combo lines, just tribal synergy and beatdown.
- PLAN useful_as_built=yes: 37 lands with a clean two-color base, roughly a dozen ramp/cost-reduction pieces, real card draw, removal and two sweepers, plus many redundant Dragon-damage wincons means it can be shuffled up and played as printed, even if the curve is top-heavy.
- PLAN summary_honest=yes: The summary's claims — cost reduction, double-strike attacks, expensive threats with protection (Greaves, Boots, Heroic Intervention), and Dragon payoffs as finishers — all match the list, and it openly admits the slow early game.
- [INFO] `curve_summary`: average mana value 3.48 over 62 nonland cards
- [INFO] `mana_pass`: the builder moved 1 card of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 1 Command Tower | land | Reliable red and green mana for the Dragon plan.
- 1 Stomping Ground | land | An efficient dual-color land.
- 1 Taiga | land | An efficient dual-color land.
- 1 Karplusan Forest | land | Provides either needed color early.
- 1 Copperline Gorge | land | Provides early red or green mana.
- 1 Spire Garden | land | Provides both deck colors.
- 1 Mana Confluence | land | Flexible access to either deck color.
- 1 City of Brass | land | Flexible access to either deck color.
- 1 Exotic Orchard | land | Flexible color fixing in multiplayer games.
- 1 Rootbound Crag | land | A dual-color land for red and green spells.
- 1 Rockfall Vale | land | A dual-color land for red and green spells.
- 1 Game Trail | land | A dual-color land for red and green spells.
- 1 Cinder Glade | land | A Mountain and Forest dual land.
- 1 Cragcrown Pathway // Timbercrown Pathway | land | Flexible access to either deck color.
- 1 Fire-Lit Thicket | land | Filters mana into the colors the deck needs.
- 1 Wooded Foothills | land | Finds a red or green land source.
- 1 Fabled Passage | land | Finds a basic land of the needed color.
- 1 Prismatic Vista | land | Finds a basic land of the needed color.
- 10 Forest | land | Basic green mana supports early ramp and green spells.
- 9 Mountain | land | Basic red mana supports Dragons and red spells.
- 1 Carnelian Orb of Dragonkind | ramp | Fixes mana while supporting Dragon deployment.
- 1 Jade Orb of Dragonkind | ramp | Provides green acceleration for the deck.
- 1 Dragon's Hoard | ramp | Dragon-focused mana acceleration.
- 1 Orb of Dragonkind | ramp | Provides color fixing and acceleration.
- 1 Pillar of Origins | ramp | Produces mana for the chosen creature type.
- 1 Patchwork Banner | ramp | Tribal mana acceleration for Dragons.
- 1 Progenitor's Icon | ramp | Provides Dragon-focused mana support.
- 1 Scaled Nurturer | ramp | Early creature-based mana acceleration.
- 1 Ganax, Astral Hunter | ramp | A Dragon that contributes to mana production.
- 1 Goldspan Dragon | ramp | A Dragon that supports explosive mana turns.
- 1 Faithless Looting | draw | Low-cost card selection and card flow.
- 1 Demand Answers | draw | Efficient card flow when resources are available.
- 1 Thrill of Possibility | draw | Turns spare cards into fresh options.
- 1 Skullclamp | draw | A cheap source of sustained card flow.
- 1 Sylvan Library | draw | Improves draw quality over multiple turns.
- 1 Garruk's Uprising | draw | Supports card flow alongside large creatures.
- 1 Beast Whisperer | draw | Rewards the deck for deploying creatures.
- 1 Guardian Project | draw | Sustains cards through creature deployment.
- 1 Return of the Wildspeaker | draw | Flexible card flow for a creature-heavy board.
- 1 Harmonize | draw | Straightforward refill for the hand.
- 1 Dragonborn Champion | draw | A Dragon-aligned source of card flow.
- 1 Heroic Intervention | interaction | Protects the board from opposing disruption.
- 1 Legolas's Quick Reflexes | interaction | A low-cost protective combat trick.
- 1 Lightning Greaves | interaction | Protects a key creature immediately.
- 1 Swiftfoot Boots | interaction | Protects important creatures while keeping them active.
- 1 Veil of Summer | interaction | Efficient protection against targeted disruption.
- 1 Pyroblast | interaction | Low-cost stack interaction.
- 1 Red Elemental Blast | interaction | Low-cost stack interaction.
- 1 Tibalt's Trickery | interaction | Broad emergency interaction.
- 1 Draconic Roar | removal | Efficient Dragon-themed targeted removal.
- 1 Dragon's Fire | removal | Targeted removal that fits the Dragon shell.
- 1 Dragon Tempest | removal | Dragon synergy that also pressures opposing creatures.
- 1 Invasion of Tarkir // Defiant Thundermaw | removal | Dragon-based targeted removal with a Dragon follow-up.
- 1 Glorybringer | removal | A Dragon that supplies repeatable creature pressure.
- 1 Scourge of Valkas | removal | Turns Dragon deployment into targeted pressure.
- 1 Foe-Razer Regent | removal | A large Dragon that answers opposing creatures.
- 1 Magmatic Hellkite | removal | A Dragon with removal utility.
- 1 Terror of the Peaks | removal | Converts creature deployment into removal pressure.
- 1 Dragonlord's Servant | synergy | Reduces the cost of the deck's Dragons.
- 1 Dragonspeaker Shaman | synergy | Helps deploy expensive Dragons earlier.
- 1 Herald's Horn | synergy | Provides tribal support for Dragons.
- 1 Shared Animosity | synergy | Makes a wide Dragon attack far more threatening.
- 1 Crucible of Fire | synergy | A dedicated anthem for Dragons.
- 1 Descendants' Path | synergy | Rewards a library built around one creature type.
- 1 Urza's Incubator | synergy | Substantially reduces Dragon costs.
- 1 Thunderbreak Regent | threat | A durable Dragon threat for the attack step.
- 1 Stormbreath Dragon | threat | An evasive Dragon attacker.
- 1 Territorial Hellkite | threat | An aggressive Dragon for consistent combat pressure.
- 1 Manaform Hellkite | threat | A Dragon threat that rewards active spellcasting.
- 1 Mirrorwing Dragon | threat | An evasive Dragon that creates difficult combat decisions.
- 1 Hoarding Dragon | threat | A Dragon body that advances the artifact package.
- 1 Dragon Broodmother | threat | Builds a growing Dragon presence over time.
- 1 Hellkite Charger | threat | A finishing Dragon attacker.
- 1 Thrakkus the Butcher | threat | A Dragon that amplifies the combat plan.
- 1 Twinflame Tyrant | threat | A powerful finishing Dragon for combat turns.
- 1 Stormwing Dragon | threat | An evasive Dragon that adds reliable pressure.
- 1 Scourge of the Throne | threat | A finishing Dragon that rewards attacking opponents.
- 1 Wrathful Red Dragon | wincon | A Dragon-based route to closing the game.
- 1 Lathliss, Dragon Queen | wincon | A dedicated Dragon payoff that can overwhelm opponents.
- 1 Breath Weapon | wipe | A low-cost board wipe option.
- 1 Draconic Intervention | wipe | A flexible Dragon-themed board wipe.
- 1 Thunder Dragon | wipe | A Dragon finisher that also clears smaller boards.

</details>

### 15. delegated commander, owned first

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 230 names.

Commander: Astarion, the Decadent.

Grade: baseline, score 0.05, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $162.79 the whole deck.

**Summary:** Astarion leads a white-black lifegain deck that establishes a board of Angels, Clerics, and other themed creatures, using equipment and steady card advantage to keep the plan moving. The deck turns that growing board and its lifegain focus into pressure, with Frodo, Sauron's Bane, Lyra Dawnbringer, Grave Venerations, and finisher-marked threats providing decisive endgames. It gives up explosive speed for a dependable mana base, substantial protection for key pieces, and a broad package of removal and board resets.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The land count sits above the norm of the format, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core is a coherent WB lifegain midrange built around Astarion with payoffs like Angel of Vitality, Rosie Cotton, Light of Promise, Exemplar of Light and Frodo, but a separate equipment/Puresteel package and three unconditional board wipes pull against the creature-board plan the summary describes.
- PLAN theme_fit=yes: It is a Commander deck in white-black with a builder-chosen lifegain commander (Astarion, the Decadent) at a modest, precon-level power consistent with bracket 2.
- PLAN useful_as_built=yes: 37 lands plus seven ramp pieces support a curve topping out around five, and there are enough creature threats, removal, and card draw to play games out and close them.
- PLAN summary_honest=yes: placeholder
- [WARN] `not_owned`: Astarion, the Decadent: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.10 over 62 nonland cards

<details><summary>The deck list</summary>

- 18 Plains | land | Basic white land slot for a stable mana base.
- 15 Swamp | land | Basic black land slot for a stable mana base.
- 1 Command Tower | land | Nonbasic land slot supporting the two-color deck.
- 1 City of Brass | land | Nonbasic land slot supporting the two-color deck.
- 1 Exotic Orchard | land | Nonbasic land slot supporting the two-color deck.
- 1 Plaza of Heroes | land | Nonbasic land slot supporting the legendary commander.
- 1 Sol Ring | ramp | Efficient early mana acceleration.
- 1 Arcane Signet | ramp | Reliable artifact-based mana acceleration.
- 1 Fellwar Stone | ramp | Additional early artifact mana.
- 1 Commander's Sphere | ramp | Mana acceleration that remains useful later.
- 1 Wayfarer's Bauble | ramp | Land-based mana development.
- 1 Sword of the Animist | ramp | Equipment-based mana development.
- 1 Giada, Font of Hope | ramp | Creature-based ramp that fits the Angel contingent.
- 1 Lotho, Corrupt Shirriff | ramp | Ramp creature for the deck's game plan.
- 1 Chromatic Lantern | ramp | Artifact mana support for consistent casting.
- 1 Path to Exile | ramp | Ramp slot that also fits the deck's white cards.
- 1 Buster Sword | draw | Equipment-based card advantage.
- 1 Call of the Ring | draw | Persistent card-advantage piece.
- 1 Exemplar of Light | draw | Angel card-advantage creature.
- 1 Idol of Oblivion | draw | Low-cost artifact card advantage.
- 1 Inspiring Overseer | draw | Creature-based card advantage for the Angel package.
- 1 Lembas | draw | Artifact card-advantage slot.
- 1 Mask of Memory | draw | Equipment-based card advantage.
- 1 Night's Whisper | draw | Straightforward black card advantage.
- 1 Puresteel Paladin | draw | Card advantage tied to the equipment package.
- 1 Skullclamp | draw | Efficient equipment-based card advantage.
- 1 Wall of Omens | draw | Early defensive card-advantage creature.
- 1 Bastion Protector | interaction | Commander-focused protection.
- 1 Boromir, Warden of the Tower | interaction | Legendary interaction creature.
- 1 Champion's Helm | interaction | Equipment protection for Astarion or a major threat.
- 1 Clever Concealment | interaction | Protective instant for the established board.
- 1 Lightning Greaves | interaction | Low-cost protection equipment.
- 1 Swiftfoot Boots | interaction | Additional protection equipment.
- 1 Banishing Light | removal | Flexible permanent removal.
- 1 Bitter Triumph | removal | Efficient black removal.
- 1 Generous Gift | removal | Broad white removal.
- 1 Get Lost | removal | Flexible white removal.
- 1 Infernal Grasp | removal | Reliable black removal.
- 1 Stroke of Midnight | removal | Additional broad white removal.
- 1 Swords to Plowshares | removal | Efficient creature removal.
- 1 Aerith Gainsborough | synergy | Legendary lifegain synergy piece.
- 1 Angel of Vitality | synergy | Angel lifegain synergy creature.
- 1 Compassionate Healer | synergy | Dedicated lifegain synergy creature.
- 1 Kor Firewalker | synergy | Lifegain-focused synergy creature.
- 1 Light of Promise | synergy | Lifegain synergy enchantment.
- 1 Night Nurse, Healer of Heroes | synergy | Lifegain-themed legendary synergy piece.
- 1 Prideful Feastling | synergy | Creature slot dedicated to the lifegain plan.
- 1 Rosie Cotton of South Lane | synergy | Lifegain synergy from the creature suite.
- 1 Second Breakfast | synergy | Lifegain-oriented instant synergy.
- 1 Wanderbrine Preacher | synergy | Cleric synergy creature for the central plan.
- 1 Angel of Invention | threat | Angel threat that helps build a meaningful board.
- 1 Bill the Pony | threat | Low-investment legendary threat.
- 1 Dawnhand Eulogist | threat | Finisher-marked threat for closing games.
- 1 Foggy Swamp Hunters | threat | Creature threat for maintaining pressure.
- 1 Gollum, Silent Slinker // Meager Meal | threat | Black creature threat with a useful alternate half.
- 1 Invisible Woman, Sue Storm | threat | Legendary creature threat.
- 1 Minwu, White Mage | threat | Legendary white creature threat.
- 1 Rabaroo Troop | threat | Creature threat that broadens the board presence.
- 1 Rooftop Percher | threat | Additional creature threat.
- 1 Shattered Angel | threat | Angel threat that complements the lifegain theme.
- 1 Sneering Shadewriter | threat | Finisher-marked black threat.
- 1 Victory's Herald | threat | High-impact Angel threat.
- 1 Frodo, Sauron's Bane | wincon | Finisher-marked win condition.
- 1 Grave Venerations | wincon | Finisher-marked black win condition.
- 1 Lyra Dawnbringer | wincon | Finisher-marked Angel win condition.
- 1 Dusk // Dawn | wipe | Board-reset option that can favor the creature plan.
- 1 Fumigate | wipe | Lifegain-aligned board reset.
- 1 Vanquish the Horde | wipe | Efficient white board-reset option.

</details>

### 16. a tight budget, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 231 names.

Commander: Denethor, Ruling Steward.

Grade: typical, score 0.39, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $18.19 to buy, $44.31 the whole deck.

**Summary:** Denethor leads a steady Orzhov aristocrats game: establish disposable creatures and sacrifice outlets, turn deaths into value, then drain the table with Bastion of Remembrance and Falkenrath Noble while the commander keeps the board supplied. The deck can play a patient removal-and-wipe game before committing its finishers, with Archfiend of Ifnir chief among its larger closing threats. It gives up explosive speed and relies on keeping creatures and sacrifice material available to maintain pressure.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The land count sits above the norm of the format, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core is coherent — many cheap sacrifice outlets, token/body makers, and drain payoffs — but a commander-protection sub-package (Bastion Protector, Frontline Medic, Take Up the Shield, Reprieve, Swiftfoot Boots) and three symmetrical wipes pull against a plan that wants a wide board of disposable creatures.
- PLAN theme_fit=yes: It is an Orzhov aristocrats Commander deck under Denethor with sac outlets, death triggers, and drain effects, at a modest budget-style power level consistent with bracket 2.
- PLAN useful_as_built=yes: 37 lands plus several cheap rocks, a low curve, ample removal, and multiple win routes (drain plus combat) make it immediately playable as written.
- PLAN summary_honest=partly: The drain engines, removal, and wipes it names are all present, but calling Archfiend of Ifnir a chief closing threat oversells a card the deck has almost no discard/cycling support for, and it understates how few dedicated drain payoffs (essentially Bastion, Falkenrath Noble) the list actually holds.
- [WARN] `not_owned`: Aron, Benalia's Ruin: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ayli, Eternal Pilgrim: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bartolomé del Presidio: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Woe Strider: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Yahenni, Undying Partisan: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Sivriss, Nightmare Speaker: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Skullport Merchant: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Smothering Abomination: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Lord Skitter's Butcher: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Old Flitterfang: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ruthless Lawbringer: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Shadowheart, Dark Justiciar: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ghoulcaller Gisa: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vampire Gourmand: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Umbral Collar Zealot: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Elder Arthur Maxson: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Baron Bertram Graywater: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Master of Dark Rites: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ruthless Technomancer: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Lord of the Forsaken: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bastion of Remembrance: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Falkenrath Noble: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 2.89 over 62 nonland cards

<details><summary>The deck list</summary>

- 15 Plains | land | Primary white mana source for the deck.
- 14 Swamp | land | Primary black mana source for the deck.
- 1 Command Tower | land | Reliable dual-color land for Denethor’s colors.
- 1 Exotic Orchard | land | Flexible colored mana from the table.
- 1 Path of Ancestry | land | Provides both colors while supporting the creature-heavy plan.
- 1 Secluded Courtyard | land | Names a shared creature type to provide colored mana.
- 1 Unclaimed Territory | land | Creature-focused colored mana fixing.
- 1 Evolving Wilds | land | Finds the basic color currently needed.
- 1 Grand Coliseum | land | Additional access to either deck color.
- 1 Capital City | land | A dual-color utility land slot.
- 1 Skullclamp | draw | Turns small disposable creatures into cards.
- 1 Idol of Oblivion | draw | Repeatable card flow alongside creature tokens.
- 1 Wall of Omens | draw | Early body that replaces itself.
- 1 Inspiring Overseer | draw | Creature-based card draw that helps fill the board.
- 1 Nasty End | draw | Uses a creature loss to refill the hand.
- 1 Painful Truths | draw | Efficient spell-based card draw.
- 1 Massacre Girl, Known Killer | draw | Card advantage attached to a creature.
- 1 Tome of Legends | draw | Steady card flow while using the commander.
- 1 Mask of Memory | draw | Equipment-based card selection and draw.
- 1 Stone of Erech | draw | Low-cost artifact card advantage.
- 1 Lembas | draw | Early artifact that replaces itself.
- 1 Sol Ring | ramp | Fast early acceleration for deploying the board.
- 1 Arcane Signet | ramp | Reliable two-color mana fixing.
- 1 Fellwar Stone | ramp | Low-cost colored mana acceleration.
- 1 Thought Vessel | ramp | Two-mana acceleration that stays useful late.
- 1 Commander's Sphere | ramp | Fixes mana early and can become a card later.
- 1 Deadly Dispute | ramp | Converts a sacrificed creature into mana and resources.
- 1 Astral Cornucopia | ramp | Flexible artifact mana for longer games.
- 1 Inherited Envelope | ramp | Low-cost artifact mana support.
- 1 Relic of Legends | ramp | Mana production that works well with the legendary creatures.
- 1 White Auracite | ramp | Additional inexpensive mana acceleration.
- 1 Bitter Triumph | removal | Efficient answer to a troublesome permanent.
- 1 Claim the Precious | removal | Direct answer to an opposing creature.
- 1 Crib Swap | removal | Creature removal that fits the creature-centric deck.
- 1 Fiend Hunter | removal | Creature-based removal that can also be sacrificed later.
- 1 Infernal Grasp | removal | Straightforward answer to key threats.
- 1 Generous Gift | removal | Flexible removal for nearly any permanent.
- 1 Stroke of Midnight | removal | Broad instant-speed permanent removal.
- 1 Aron, Benalia's Ruin | synergy | A sacrifice-focused creature that supports Denethor’s plan.
- 1 Ayli, Eternal Pilgrim | synergy | Provides another sacrifice outlet on a creature.
- 1 Bartolomé del Presidio | synergy | A low-cost sacrifice outlet that benefits from creatures dying.
- 1 Gollum the Abandoned | synergy | Supports recurring creature and sacrifice play patterns.
- 1 Gollum, Patient Plotter | synergy | A recurring creature that gives the deck sacrifice material.
- 1 Woe Strider | synergy | Supplies both a sacrifice outlet and extra bodies.
- 1 Yahenni, Undying Partisan | synergy | Free sacrifice outlet that benefits as creatures die.
- 1 Sivriss, Nightmare Speaker | synergy | Creature-based graveyard and sacrifice support.
- 1 Skullport Merchant | synergy | Turns spare creatures into cards and resources.
- 1 Smothering Abomination | synergy | Rewards the deck for sacrificing creatures.
- 1 Lord Skitter's Butcher | threat | Creature threat that contributes to the sacrifice-focused board.
- 1 Old Flitterfang | threat | Creature threat that helps keep material on the battlefield.
- 1 Ruthless Lawbringer | threat | Creature threat that adds pressure while advancing the board.
- 1 Shadowheart, Dark Justiciar | threat | Sacrifice-oriented creature threat with strong late-game value.
- 1 Ghoulcaller Gisa | threat | Builds a threatening creature board from creatures already used.
- 1 Vampire Gourmand | threat | Low-cost creature that supports the deck’s attrition plan.
- 1 Umbral Collar Zealot | threat | Early creature that contributes to the sacrifice package.
- 1 Elder Arthur Maxson | threat | A durable creature threat for the midgame.
- 1 Baron Bertram Graywater | threat | Creature threat that helps produce material for the board.
- 1 Master of Dark Rites | threat | Powerful creature threat for the sacrifice strategy.
- 1 Ruthless Technomancer | threat | Turns creature material into a meaningful resource swing.
- 1 Lord of the Forsaken | threat | Large black creature that gives the deck a substantial threat.
- 1 Archfiend of Ifnir | wincon | A finishing threat that pressures opposing creature boards.
- 1 Bastion of Remembrance | wincon | Death-drain finisher that rewards sacrificing creatures.
- 1 Falkenrath Noble | wincon | Creature death-drain effect that closes games through attrition.
- 1 Austere Command | wipe | Flexible reset button for unfavorable boards.
- 1 Dusk // Dawn | wipe | Board reset with later creature recovery value.
- 1 Martial Coup | wipe | Wipes the board while leaving creature material behind.
- 1 Gift of Immortality | interaction | Helps preserve an important creature through removal.
- 1 Swiftfoot Boots | interaction | Protects Denethor or a key payoff creature.
- 1 Take Up the Shield | interaction | Combat and removal protection for a key creature.
- 1 Bastion Protector | interaction | Creature-based protection for the commander.
- 1 Frontline Medic | interaction | Protective creature that helps stabilize the board.
- 1 Reprieve | interaction | Flexible instant-speed disruption against a pivotal spell.

</details>

### 17. upgrade a precon, owned first

Format: Commander. Theme: goblins. Pool: owned_first. Shortlist: 205 names.

Commander: Zada, Hedron Grinder.

Grade: typical, score 0.59, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: yes, for precon_share, profile_off_band. Block findings: 0.

Cost: $33.57 to buy, $221.81 the whole deck.

**Summary:** A mono-red Goblin storm deck that develops a wide board, converts small creatures into mana and cards, and turns targeted spells into team-wide bursts of velocity. It combines tribal haste and anthem effects with token production, sacrifice damage, scalable removal, and explosive combat finishes.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: The cheap targeted cantrips and pump spells (Crimson Wisps, Expedite, Fists of Flame, Ancestral Anger, Haze of Rage) pair directly with Zada and the many token/goblin bodies, while Skirk Prospector, Battle Hymn, Mana Geyser and Seething Song fuel Grapeshot, Empty the Warrens and Past in Flames — all pieces feed the same go-wide storm plan.
- PLAN theme_fit=yes: It is a mono-red Goblin storm Commander deck under Zada, upgraded with staples and goblin payoffs but without tutored fast-combo or heavy free-interaction packages, which sits comfortably at bracket 3.
- PLAN useful_as_built=yes: 37 lands plus Sol Ring, Arcane Signet, Fellwar Stone and ritual effects support a low mono-red curve, and the deck has multiple redundant win routes (storm Grapeshot, wide alpha strikes with Shared Animosity/Bushwhacker, Bombardment sacrifice damage).
- PLAN summary_honest=yes: Every claim — wide boards, creatures into mana and cards (Prospector, Battle Hymn, Skullclamp), copied targeted spells, anthems/haste lords, sacrifice damage (Bombardment, Boggart Shenanigans, Pashalik Mons), scalable removal (Earthquake, Goblin Negotiation) and combat finishes (Shared Animosity, Great Train Heist) — is backed by cards in the list.
- [INFO] `curve_summary`: average mana value 2.34 over 62 nonland cards
- [WARN] `precon_share`: the deck keeps 60 of the 77 nonbasic Goblin Storm precon names, and the rule asks for 66
- [WARN] `profile_off_band`: the removal count is 14, and bracket 3 wants 6 to 12
- [INFO] `bracket_cut`: to hold bracket 3, the builder cut 1 card: Storm-Kiln Artist (the combo Storm-Kiln Artist + Haze of Rage), and added 1 basic land
- [INFO] `cards_trimmed`: the list was 2 cards over, so the builder cut these cards: Blasphemous Act, Daring Discovery

<details><summary>The deck list</summary>

- 1 Arena of Glory | land | Red-producing utility land.
- 1 Castle Embereth | land | Creature-pump utility land for wide Goblin attacks.
- 1 Den of the Bugbear | land | Creature-land threat that survives sorcery-speed sweepers.
- 1 Dwarven Mine | land | Mountain land with an extra Goblin body payoff.
- 1 Forgotten Cave | land | Cycling land that can convert excess mana into another card.
- 1 Fountainport | land | Utility land for converting tokens into value.
- 1 Hidden Volcano | land | Red mana source with late-game utility.
- 1 Kher Keep | land | Token-producing utility land that adds bodies for go-wide payoffs.
- 26 Mountain | land | Reliable untapped red sources for the deck’s heavily red requirements.
- 1 Reliquary Tower | land | Utility land for retaining a large hand after explosive draw turns.
- 1 Shinka, the Bloodsoaked Keep | land | Red source that can give a key legendary attacker first strike.
- 1 War Room | land | Repeatable card-draw utility land.
- 1 Abrade | removal | Flexible answer to creatures or artifacts.
- 1 Ancestors' Aid | synergy | A targeted cantrip-style combat trick that scales with Zada.
- 1 Ancestral Anger | draw | Efficient targeted card draw and combat scaling with Zada.
- 1 Arcane Signet | ramp | Efficient early mana acceleration.
- 1 Battle Hymn | ramp | Converts a Goblin board into a large burst of red mana.
- 1 Boggart Shenanigans | removal | Turns sacrificed and defeated Goblins into direct damage.
- 1 Champion's Helm | interaction | Protects the commander and supports attacking with it.
- 1 Chaos Warp | removal | Versatile answer to troublesome permanents.
- 1 Conspicuous Snoop | synergy | Goblin value engine that helps deploy creatures from the library.
- 1 Crimson Wisps | draw | Cheap Zada-friendly cantrip that grants haste.
- 1 Darksteel Plate | interaction | Durable protection for the commander or a key Goblin.
- 1 Dragon Fodder | synergy | Efficiently adds multiple bodies for Goblin and Zada payoffs.
- 1 Earthquake | wipe | Scalable board control that can also finish weakened opponents.
- 1 Empty the Warrens | synergy | Produces a large Goblin board after a spell-heavy turn.
- 1 Expedite | draw | One-mana haste cantrip that becomes powerful with Zada.
- 1 Fable of the Mirror-Breaker // Reflection of Kiki-Jiki | ramp | Provides Treasure acceleration, card selection, and creature-copying value.
- 1 Faithless Looting | draw | Efficient hand filtering and graveyard setup.
- 1 Fellwar Stone | ramp | Low-cost mana acceleration.
- 1 Fists of Flame | draw | Zada-scalable card draw and a potentially lethal combat boost.
- 1 Gempalm Incinerator | removal | Goblin-scaled creature removal that can cycle for value.
- 1 Glimpse the Impossible | ramp | Impulse-card advantage with mana support for explosive turns.
- 1 Goblin Bombardment | removal | Sacrifice outlet that turns Goblin tokens into repeatable damage.
- 1 Goblin Bushwhacker | synergy | Haste and team-wide power boost for token-fueled attacks.
- 1 Goblin Chieftain | synergy | Anthem and haste support for the Goblin army.
- 1 Goblin Fireleaper | removal | Goblin creature that supplies repeatable damage-based removal.
- 1 Goblin Lackey | synergy | Accelerates high-impact Goblins onto the battlefield through combat.
- 1 Goblin Matron | synergy | Finds the most useful Goblin for the current board state.
- 1 Goblin Negotiation | removal | Goblin-themed removal option.
- 1 Goblin Trashmaster | removal | Goblin anthem with artifact-removal utility.
- 1 Goblin Warchief | synergy | Reduces Goblin costs and gives the team haste.
- 1 Grapeshot | removal | Storm-based removal that rewards long spell chains.
- 1 Great Train Heist | wincon | A powerful combat finisher for a developed creature board.
- 1 Haze of Rage | other | Repeatable storm combat boost that scales dramatically with a wide board.
- 1 Howlsquad Heavy | ramp | Goblin mana acceleration attached to a creature.
- 1 Impact Tremors | other | Turns every Goblin and token entering into opponent damage.
- 1 Krenko's Command | synergy | Efficiently creates multiple Goblin bodies.
- 1 Krenko, Mob Boss | threat | Major repeatable Goblin-token producer and battlefield threat.
- 1 Lightning Bolt | removal | Highly efficient creature removal or finishing damage.
- 1 Lightning Greaves | interaction | Protects Zada while enabling immediate activation-free attacks.
- 1 Lightning Strike | removal | Efficient damage-based removal with player-targeting reach.
- 1 Mana Geyser | ramp | Creates explosive mana for storm and token turns.
- 1 Mogg War Marshal | synergy | Multiple Goblin bodies from one card, ideal for sacrifice and go-wide plans.
- 1 Moria Marauder | synergy | Goblin attacker that advances the tribal pressure plan.
- 1 Pashalik Mons | removal | Goblin payoff that converts tokens and sacrifices into damage.
- 1 Past in Flames | other | Lets a stocked graveyard fuel another explosive spell turn.
- 1 Quest for the Goblin Lord | synergy | Tribal payoff that rewards continued Goblin development.
- 1 Ruby Medallion | other | Reduces the cost of the deck’s many red spells.
- 1 Rundvelt Hordemaster | synergy | Goblin anthem and card-advantage engine.
- 1 Sazacap's Brew | draw | Cheap card selection for finding action or lands.
- 1 Seething Song | ramp | Burst mana for high-output turns.
- 1 Shared Animosity | synergy | Makes wide Goblin attacks capable of ending games quickly.
- 1 Siege-Gang Commander | removal | Creates several Goblins and converts them into direct damage.
- 1 Siege-Gang Lieutenant | removal | Goblin token-maker with damage utility.
- 1 Skirk Prospector | ramp | Converts expendable Goblins into mana for explosive turns.
- 1 Skullclamp | draw | Exceptional card draw with disposable Goblin tokens.
- 1 Sol Ring | ramp | Fast colorless acceleration.
- 1 Spreading Insurrection | other | High-impact spell for swinging combat in a crowded game.
- 1 Swiftfoot Boots | interaction | Protects Zada while granting haste.
- 1 Throne of Eldraine | ramp | Mana rock that also supplies late-game card advantage.
- 1 Vandalblast | wipe | Artifact sweeper that can be cast efficiently when needed.
- 1 Wild Ride | other | Combat-focused spell that supports aggressive creature turns.
- 1 Witch's Mark | draw | Card filtering with a targeted combat boost that works well with Zada.

</details>

### 18. upgrade a precon, any card

Format: Commander. Theme: turtles. Pool: any_card. Shortlist: 364 names.

Commander: Heroes in a Half Shell.

Grade: baseline, score 0.08, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $210.95 to buy, $210.95 the whole deck.

**Summary:** Heroes in a Half Shell keeps Turtle Power focused on a broad cast of Turtles, Mutants, Ninjas, and its original character package. The deck develops its board, supports the team with its retained tribal pieces, and closes through combat pressure from its larger threats and finishers, chief among them Dimension X Pizzasaur, Everything Pizza, and Raphael, the Muscle. The upgraded mana base is smoother and Rhystic Study supplies added staying power, while the deck still gives up some speed and consistency to retain the precon’s many thematic cards.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core is a coherent creature-heavy Turtle/Mutant tribal board-building deck, but it also runs four sweepers (Blasphemous Act, Vanquish the Horde, Wave Goodbye, Game Over) that fight its own wide board, and it keeps basic-land fetchers (Cultivate, Evolving Wilds, Ash Barrens) in a list with zero basics.
- PLAN theme_fit=yes: It is the Turtle Power precon kept intact in theme (Turtles, Mutants, Ninjas, original characters) with modest bracket-3-appropriate additions like Rhystic Study, Assassin's Trophy, Sol Ring, and a shock/pain dual land base.
- PLAN useful_as_built=partly: 38 well-fixed lands and a deep creature base make it castable and able to win through combat, but ramp is thin for a five-color deck and several slots (Cultivate, Evolving Wilds, Ash Barrens cycling) are near-dead with no basics in the list.
- PLAN summary_honest=yes: The summary accurately names the tribal/combat plan, the actual finishers present in the list, the Rhystic Study and mana-base additions, and openly concedes the lost speed and consistency from retaining precon cards.
- [INFO] `curve_summary`: average mana value 3.46 over 61 nonland cards
- [WARN] `profile_off_band`: the ramp count is 3, and bracket 3 wants 8 to 13
- [WARN] `profile_off_band`: the draw count is 2, and bracket 3 wants 8 to 14
- [WARN] `profile_off_band`: the removal count is 2, and bracket 3 wants 6 to 12
- [WARN] `profile_off_band`: the wipe count is 1, and bracket 3 wants 2 to 5
- [WARN] `profile_off_band`: the interaction count is 0, and bracket 3 wants 4 to 12
- [WARN] `profile_off_band`: the mana available on turn four is 4.19, and bracket 3 wants 4.2 or more (mean over 10000 hands)

<details><summary>The deck list</summary>

- 1 Acidic Slime | other | Retained from the precon to preserve its established creature package.
- 1 April O'Neil, Live on the Scene | draw | Retained precon draw support.
- 1 Arcade Cabinet | other | Retained from the precon to preserve its original theme.
- 1 Arcane Signet | ramp | Retained precon ramp support.
- 1 Assassin's Trophy | removal | Retained precon removal.
- 1 Baxter, Fly in the Ointment | other | Retained from the precon to preserve its established character package.
- 1 Bebop, Skull & Crossbones | other | Retained from the precon to preserve its established character package.
- 1 Big Mother Mouser | other | Retained from the precon to preserve its established creature package.
- 1 Biogenic Ooze | other | Retained from the precon to preserve its established creature package.
- 1 Blasphemous Act | wipe | Retained precon board control.
- 1 Chromatic Lantern | other | Retained from the precon as part of its mana package.
- 1 Continue? | other | Retained from the precon to preserve its original package.
- 1 Corpsejack Menace | other | Retained from the precon to preserve its established creature package.
- 1 Cultivate | other | Retained from the precon as part of its mana package.
- 1 Dimension X Pizzasaur | wincon | Retained precon finisher.
- 1 Donatello, the Brains | synergy | Retained precon Turtle synergy.
- 1 Double Jump // Flying Kick | other | Retained from the precon to preserve its original package.
- 1 Electric Seaweed | other | Retained from the precon to preserve its established creature package.
- 1 Endless Foot Assault | other | Retained from the precon to preserve its original package.
- 1 Everything Pizza | wincon | Retained precon finisher.
- 1 Exploding Barrel | other | Retained from the precon to preserve its original package.
- 1 Fast Forward | other | Retained from the precon to preserve its original package.
- 1 Foot Chopper | other | Retained from the precon to preserve its original package.
- 1 From the Rubble | synergy | Retained precon tribal support.
- 1 Game Over | other | Retained from the precon to preserve its original package.
- 1 Harmonize | other | Retained from the precon to preserve its original package.
- 1 Here Comes a New Hero! | other | Retained from the precon to preserve its original package.
- 1 High Score | other | Retained from the precon to preserve its original package.
- 1 Irma, Part-Time Mutant | other | Retained from the precon to preserve its established character package.
- 1 Casey Jones, Back Alley Brute | other | Retained from the precon to preserve its established character package.
- 1 Krang, the All-Powerful | other | Retained from the precon to preserve its established character package.
- 1 Leatherhead, Iron Gator | other | Retained from the precon to preserve its established character package.
- 1 Leonardo, the Balance | threat | Retained precon Turtle threat.
- 1 Lessons from Life | other | Retained from the precon to preserve its original package.
- 1 Level Up | other | Retained from the precon to preserve its original package.
- 1 Lita, Little Orphan Amphibian | synergy | Retained precon Turtle synergy.
- 1 Michelangelo, the Heart | synergy | Retained precon Turtle synergy.
- 1 Mole Module | other | Retained from the precon to preserve its original package.
- 1 Mona Lisa, Science Geek | other | Retained from the precon to preserve its established character package.
- 1 Ninja Pizza | other | Retained from the precon to preserve its original package.
- 1 Raphael, the Muscle | threat | Retained precon finisher.
- 1 Rat King, Pale Piper | other | Retained from the precon to preserve its established character package.
- 1 Ray Fillet, Wave Warrior | other | Retained from the precon to preserve its established character package.
- 1 Roadkill Rodney | other | Retained from the precon to preserve its established creature package.
- 1 Rocksteady, Mutant Marauder | other | Retained from the precon to preserve its established character package.
- 1 Shellshock | other | Retained from the precon to preserve its original package.
- 1 Shredder, Shadow Master | other | Retained from the precon to preserve its established character package.
- 1 Sol Ring | ramp | Retained precon ramp.
- 1 Special Move | other | Retained from the precon to preserve its original package.
- 1 Splinter, the Mentor | other | Retained from the precon to preserve its established character package.
- 1 Steelbane Hydra | removal | Retained from the precon to preserve its established creature package.
- 1 Super Combo | other | Retained from the precon to preserve its original package.
- 1 Swift Demise | other | Retained from the precon to preserve its original package.
- 1 Tempestra, Dame of Games | other | Retained from the precon to preserve its established character package.
- 1 Together Forever | other | Retained from the precon to preserve its original package.
- 1 Tokka & Rahzar, Unsupervised | ramp | Retained precon ramp support.
- 1 Vanquish the Horde | other | Retained from the precon to preserve its board-control package.
- 1 Vigor | other | Retained from the precon to preserve its established creature package.
- 1 Voracious Hydra | other | Retained from the precon to preserve its established creature package.
- 1 Wave Goodbye | other | Retained from the precon to preserve its original package.
- 1 Rhystic Study | draw | Added as a durable source of card advantage.
- 1 Ash Barrens | land | Retained precon land slot.
- 1 Big Apple, 3 a.m. | land | Retained precon land slot.
- 1 Cinder Glade | land | Retained precon land slot.
- 1 City of Brass | land | Retained precon land slot.
- 1 Command Tower | land | Retained precon land slot.
- 1 Dragonskull Summit | land | Retained precon land slot.
- 1 Evolving Wilds | land | Retained precon land slot.
- 1 Grand Coliseum | land | Retained precon land slot.
- 1 Hidden Hideout | land | Retained precon land slot.
- 1 Path of Ancestry | land | Retained precon land slot.
- 1 Rain-Slicked Copse | land | Retained precon land slot.
- 1 Rootbound Crag | land | Retained precon land slot.
- 1 Smoldering Marsh | land | Retained precon land slot.
- 1 Sodden Verdure | land | Retained precon land slot.
- 1 Sunken Hollow | land | Retained precon land slot.
- 1 Thriving Grove | land | Retained precon land slot.
- 1 Thriving Isle | land | Retained precon land slot.
- 1 Thriving Moor | land | Retained precon land slot.
- 1 Turtle Lair | land | Retained precon land slot.
- 1 Undergrowth Stadium | land | Retained precon land slot.
- 1 Vernal Fen | land | Retained precon land slot.
- 1 Vibrant Cityscape | land | Retained precon land slot.
- 1 Adarkar Wastes | land | Added to strengthen the mana base.
- 1 Battlefield Forge | land | Added to strengthen the mana base.
- 1 Breeding Pool | land | Added to strengthen the mana base.
- 1 Brushland | land | Added to strengthen the mana base.
- 1 Caves of Koilos | land | Added to strengthen the mana base.
- 1 Exotic Orchard | land | Added to strengthen the mana base.
- 1 Hallowed Fountain | land | Added to strengthen the mana base.
- 1 Karplusan Forest | land | Added to strengthen the mana base.
- 1 Llanowar Wastes | land | Added to strengthen the mana base.
- 1 Mana Confluence | land | Added to strengthen the mana base.
- 1 Overgrown Tomb | land | Added to strengthen the mana base.
- 1 Sacred Foundry | land | Added to strengthen the mana base.
- 1 Shivan Reef | land | Added to strengthen the mana base.
- 1 Spire Garden | land | Added to strengthen the mana base.
- 1 Steam Vents | land | Added to strengthen the mana base.
- 1 Watery Grave | land | Added to strengthen the mana base.

</details>

### 19. the Hobbit family, two colours

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 174 names.

Commander: Thranduil, the Elvenking.

Grade: bad, score 0.03, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $406.93 to buy, $406.93 the whole deck.

**Summary:** Thranduil, the Elvenking anchors an Elf-heavy Sultai value deck that develops mana, fills the board with Elves and legendary Elves, and keeps resources flowing through draw and graveyard-focused synergy. It controls key opposing threats while building toward decisive closers such as Troll of Khazad-dûm and Witch-king of Angmar, chief among them. The deck gives up some explosive speed for a steady creature board, broad answers, and a resilient longer-game plan.

The quality model grades this deck below the precon baseline against the top lists of the format: the black sources cover 19.25 of the 26 its spells need.

- PLAN plan_coherent=partly: The Elf tribal/legendary-Elf value core with ramp, draw, and removal mostly pulls one way, but the three board wipes (Languish, Gnashing of Teeth, Raise the Palisade) actively fight the wide small-Elf board the deck builds, and the 'decisive closers' are two lone big black creatures unconnected to the Elf plan.
- PLAN theme_fit=yes: It is a Commander deck led by Thranduil, the Elvenking, drawn almost entirely from the Hobbit/Middle-earth card pool with a mid-power, non-combo build appropriate to bracket 3.
- PLAN useful_as_built=partly: The curve, ramp, and creature count are playable, but with only 36 lands and a reported 19.25 black sources for 26 black-demanding pips, the heavy double-black cards (Languish, Witch-king, Bilbo's Deadly Slice) will strand regularly.
- PLAN summary_honest=partly: It accurately describes the Elf-heavy Sultai value shell and admits the slower plan and mana weakness, but calling Troll of Khazad-dûm and Witch-king 'decisive closers' overstates a deck whose real win path is creature beatdown, and it omits the sweepers that punish its own board.
- [INFO] `curve_summary`: average mana value 2.98 over 63 nonland cards
- [WARN] `profile_off_band`: the worst color's share of the sources it needs is 0.74, and bracket 3 wants 0.85 or more (sources of requirement: U 14.5 of 19, B 19.25 of 26, G 15 of 19)
- [INFO] `mana_pass`: the builder moved 4 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 12 Forest | land | Provides dependable green mana for the Elf core and green ramp.
- 10 Island | land | Provides dependable blue mana for draw and interaction spells.
- 14 Swamp | land | Provides dependable black mana for removal and finishers.
- 1 Elvish Visionary | draw | An Elf that supplies card draw while fitting the creature theme.
- 1 Fateful Discovery | draw | Dedicated card-advantage spell.
- 1 Hithlain Knots | draw | Adds a flexible draw spell to keep resources flowing.
- 1 Key to the Side-Door | draw | Provides repeatable card-advantage support.
- 1 Lórien Revealed | draw | A draw spell that helps maintain a full hand.
- 1 Night's Whisper | draw | Efficient early card draw.
- 1 Palantír of Orthanc | draw | A persistent source of card advantage.
- 1 Plunder the Trollshaws | draw | Adds instant-speed card draw.
- 1 Reverent Howl | draw | Provides another draw effect within the deck's colors.
- 1 Uncover the Moon-Letters | draw | Adds long-game card advantage.
- 1 Thrór's Map | draw | Provides additional card selection and draw support.
- 1 Bilbo's Ring | interaction | A compact legendary interaction piece.
- 1 Confusticate and Bebother | interaction | Provides a reactive answer to opposing plays.
- 1 Dwarven Mattock | interaction | Gives the deck an equipment-based interaction tool.
- 1 Elrond, Moon-Reader | interaction | A legendary Elf that adds interaction and works with Thranduil's Elf plan.
- 1 Mithril Coat | interaction | Protects a key creature through an interaction slot.
- 1 My Precious // Allure of Power | interaction | Offers a versatile legendary interaction card.
- 1 Stern Scolding | interaction | An efficient reactive spell for early opposing threats.
- 1 Bilbo's Deadly Slice | removal | A direct removal spell for problematic permanents.
- 1 Bitter Downfall | removal | Provides dependable targeted removal.
- 1 Crude Bent Blade | removal | An artifact-based removal option.
- 1 Enchanted River's Grasp | removal | Adds blue removal to answer opposing threats.
- 1 Giant's Boulder | removal | A color-flexible removal tool.
- 1 Merciless Executioner | removal | A creature-based removal effect that can be reused from the graveyard by Thranduil.
- 1 Orcish Bowmasters | removal | Provides efficient creature-based removal pressure.
- 1 Quarrel | removal | Adds another low-cost removal answer.
- 1 The Black Arrow | removal | A legendary equipment that fills a removal slot.
- 1 Gnashing of Teeth | wipe | A board-clearing option against creature-heavy tables.
- 1 Languish | wipe | An efficient reset for opposing creature boards.
- 1 Raise the Palisade | wipe | A tribal-leaning sweeper that can spare the Elf board.
- 1 Arcane Signet | ramp | Reliable multicolor mana acceleration.
- 1 Delighted Halfling | ramp | Early creature ramp that helps deploy the commander and legends.
- 1 Elven Chorus | ramp | Elf-focused mana acceleration for a creature-heavy build.
- 1 Elvish Archdruid | ramp | A strong Elf mana producer for powering out larger plays.
- 1 Elvish Mystic | ramp | A one-mana Elf accelerator.
- 1 Mox Amber | ramp | Fast legendary-oriented mana acceleration.
- 1 Necklace of Girion | ramp | A legendary mana source that supports the deck's curve.
- 1 Silvan Reveler | ramp | An Elf ramp creature that supports the tribal plan.
- 1 Thranduil the Strategist | ramp | A legendary Elf ramp piece that also works with the commander.
- 1 Wood Elves | ramp | Creature-based green ramp that fits the Elf shell.
- 1 Arwen, Weaver of Hope | synergy | A legendary Elf that strengthens the commander-focused creature theme.
- 1 Celeborn the Wise | synergy | A legendary Elf that supports Thranduil's legendary-Elf plan.
- 1 Galion, Elvenking's Butler | synergy | A legendary Elf included for direct thematic synergy with Thranduil.
- 1 Mirkwood Meditator | synergy | An Elf body that deepens the deck's graveyard and tribal focus.
- 1 Mirkwood Pathmaker | synergy | An Elf supporting the deck's creature-based game plan.
- 1 Supper for Spiders | synergy | The dedicated synergy card from the available pool.
- 1 Thranduil, Sindarin Liege // Silvan Rally | synergy | A legendary Elf that supports Thranduil and the wider Elf plan.
- 1 Boughside Wanderers | threat | An Elf creature that provides board presence for the tribal plan.
- 1 Cantankerous Keepers | threat | An Elf threat that helps establish a creature board.
- 1 Elven Raft-Steerer | threat | A thematic Elf threat that adds to the deck's board development.
- 1 Elvenking's Harper | threat | An Elf threat that supports the Elvenking theme.
- 1 Galadhrim Guide | threat | An Elf threat for maintaining creature pressure.
- 1 Grey Havens Navigator | threat | An Elf threat that broadens the creature suite.
- 1 Guardian of the Halls | threat | An Elf creature that contributes to the deck's board presence.
- 1 Haunt of the Dead Marshes | threat | A Nightmare Elf threat that remains relevant to Thranduil's graveyard focus.
- 1 Lothlórien Lookout | threat | An Elf threat that reinforces the tribal creature plan.
- 1 Mirkwood Elk | threat | A green creature threat for applying battlefield pressure.
- 1 Nimrodel Watcher | threat | An Elf threat that helps keep the board populated.
- 1 Woodland Weavemaster | threat | An Elf Druid threat that fits the deck's creature density.
- 1 Inside Information | wincon | A finisher from the available pool that gives the deck a distinct closing option.
- 1 Troll of Khazad-dûm | wincon | A marked finisher that closes games after the deck has developed its resources.
- 1 Witch-king of Angmar | wincon | A marked finisher and powerful top-end threat.
- 1 Pelargir Survivor | ramp | the mana pass added it to bring the mana base inside the power level

</details>

### 20. the Hobbit family, a delegated commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 77 names.

Commander: Smaug the Magnificent.

Grade: baseline, score 0.04, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $500.83 to buy, $500.83 the whole deck.

**Summary:** Smaug the Magnificent builds Treasure over time, then turns each attack into direct pressure while a host of Dwarves, Goblins, Equipment, and Dragons develops the board. Cavern-Hoard Dragon, Desert Were-Worm, and Smaug, the Great Calamity provide the biggest closing threats, with sweepers and focused removal clearing the way when combat stalls. The tradeoff is a narrow, combat-forward plan that relies heavily on red mana and committed battlefield development.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The list is a broadly coherent mono-red combat/equipment aggro pile with some Treasure and burn payoffs, but the pieces are scattered across Dwarf, Goblin, Equipment, saga, and big-creature subplans with no single payoff tying them together, and the 'Treasure over time' engine the summary names has few actual Treasure producers backing it.
- PLAN theme_fit=no: The person asked for a dragons deck and got a mono-red Hobbit-set goodstuff/Dwarf-Goblin aggro list with only three non-commander Dragons in 99 cards.
- PLAN useful_as_built=yes: 36 lands in mono-red with a low curve, several ramp and card-draw pieces, plenty of removal and sweepers, and enough creature and equipment pressure means it can be shuffled up and played as it stands.
- PLAN summary_honest=partly: It is candid that the board is mostly Dwarves, Goblins, and Equipment and that the plan is narrow and combat-forward, but it overstates a Treasure-accumulation engine the card pool barely supports and never acknowledges that the requested dragon theme is essentially absent.
- [INFO] `curve_summary`: average mana value 3.33 over 63 nonland cards
- [WARN] `profile_off_band`: the draw count is 6, and bracket 3 wants 8 to 14

<details><summary>The deck list</summary>

- 34 Mountain | land | Primary red mana source for casting Smaug and the deck's red spells.
- 1 Rogue's Passage | land | Utility land slot that keeps the mana base largely untapped.
- 1 The Lonely Mountain | land | A red-producing themed land for the mana base.
- 1 Balin, Loremaster | draw | A dedicated card-advantage piece.
- 1 Key to the Side-Door | draw | A dedicated card-advantage artifact.
- 1 Palantír of Orthanc | draw | A dedicated card-advantage artifact.
- 1 Ragged Short Spear | draw | An Equipment slot that contributes card advantage.
- 1 Thrór's Map | draw | A themed artifact for card advantage.
- 1 Óin the Brave | draw | A Dwarf that contributes card advantage.
- 1 Bilbo's Ring | interaction | A flexible protective interaction piece.
- 1 Dwarven Mattock | interaction | A themed Equipment interaction slot.
- 1 Mithril Coat | interaction | A protective Equipment interaction piece.
- 1 The One Ring | interaction | A powerful legendary interaction artifact.
- 1 Arcane Signet | ramp | Reliable early mana fixing and acceleration.
- 1 Bag End Banquet | ramp | A themed ramp artifact.
- 1 Burn, Burn, Tree and Fern | ramp | A ramp Saga that supports the deck's mana development.
- 1 Dragon's Desire | ramp | Dragon-themed mana acceleration.
- 1 Mox Amber | ramp | A low-cost legendary mana source.
- 1 Orcrist, Goblin-cleaver | ramp | An Equipment slot that also advances mana.
- 1 The Misty Mountains Cold | ramp | A themed ramp Saga.
- 1 The Reaver Cleaver | ramp | A Treasure-focused Equipment ramp piece.
- 1 Thorin, Company's Leader | ramp | A ramp creature that also adds a substantial body.
- 1 Wayfarer's Bauble | ramp | Early mana acceleration for reaching Smaug promptly.
- 1 Battle-Scarred Goblin | removal | A creature-based removal option.
- 1 Fire of Orthanc | removal | A direct removal spell.
- 1 Gandalf, Spark Starter | removal | A legendary removal creature.
- 1 Giant's Boulder | removal | An artifact removal option.
- 1 Goblin Cratermaker | removal | A compact creature-based removal piece.
- 1 Goblin Fireleaper | removal | A Goblin removal creature.
- 1 Improvised Club | removal | A low-cost removal spell.
- 1 Inferno Titan | removal | A large removal creature that adds battlefield presence.
- 1 Smite the Deathless | removal | A dedicated instant-speed removal slot.
- 1 Stone-Giant of High Pass | removal | A removal creature with a meaningful body.
- 1 The Black Arrow | removal | A legendary Equipment removal piece.
- 1 Thorin, Mountain-king | removal | A legendary Dwarf removal option.
- 1 Last Light of Durin's Day | synergy | A dedicated thematic synergy piece.
- 1 Cavern-Hoard Dragon | wincon | A Dragon finisher that supports the Treasure-centered plan.
- 1 Desert Were-Worm | wincon | A large evasive finisher.
- 1 Smaug, the Great Calamity // Spew Flame | wincon | A second Smaug-themed Dragon finisher.
- 1 Call Forth the Tempest | wipe | A reset button for crowded boards.
- 1 Desolation of Smaug | wipe | A Smaug-themed board reset.
- 1 Glóin the Mighty // Easy Pickings | wipe | A creature slot that can serve as a board wipe.
- 1 Andúril, Flame of the West | other | A legendary themed Equipment for the deck's creature package.
- 1 Andúril, Narsil Reforged | other | A second legendary themed Equipment.
- 1 Bombur, Gentle Dreamer | other | A themed legendary Dwarf for the creature package.
- 1 Bothersome Noisemaker | other | A Goblin body that supports the deck's Middle-earth theme.
- 1 Dáin Ironfoot | other | A substantial legendary Dwarf for the battlefield.
- 1 Dori, Bearer of Friends | other | A themed legendary Dwarf body.
- 1 Dwarven Mauler | other | A Dwarf creature for the deck's supporting creature suite.
- 1 Dwarven Warriors | other | A thematic Dwarf creature.
- 1 Gandalf, Goblins' Bane // Flameshape | other | A legendary themed creature with an Adventure.
- 1 Getaway Barrel | other | A low-cost themed artifact utility slot.
- 1 Glamdring | other | A legendary Equipment for the deck's creature suite.
- 1 Goblin-town Flunkies | other | A Goblin creature that broadens the battlefield presence.
- 1 Gundabad Opportunist | other | A Goblin creature for the supporting creature package.
- 1 Guttersnipe | other | A Goblin Shaman that rewards the deck's spells.
- 1 Iron Hills Stalwart | other | A Dwarf creature for the themed creature base.
- 1 Long-Lost Lances | other | An Equipment that supports the creature plan.
- 1 Misty Mountains Raider | other | A Goblin creature for board presence.
- 1 Olog-hai Crusher | other | A large Troll body for the creature package.
- 1 Oliphaunt | other | A large themed creature for battlefield pressure.
- 1 Orcish Siegemaster | other | An Orc creature that adds to the deck's attacking force.
- 1 Smaug's Fury | other | A Smaug-themed instant utility spell.
- 1 Snowslope Hunter | other | A Goblin creature that supports the themed roster.
- 1 Sting, Bilbo's Sword | other | A legendary Equipment for the creature suite.
- 1 Well-Worn Spatula | other | A themed Equipment utility slot.

</details>

### 21. the Hobbit family, mana from outside

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 156 names.

Commander: Smaug the Impenetrable.

Grade: baseline, score 0.03, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $701.29 to buy, $701.29 the whole deck.

**Summary:** Smaug the Impenetrable is the centerpiece of a red-black Hobbit-themed pressure deck that develops mana, deploys Dragons and Goblin-heavy creatures, and clears away resistance with a deep removal suite. Smaug’s haste and flying make it an immediate threat, while Treasure production helps sustain the deck’s expensive threats and equipment. The deck closes through combat from its finishing creatures, chief among them its Dragons, Troll, and Gollum, but gives up some speed and consistency by leaning on a broad thematic creature-and-equipment package rather than a narrow combo plan.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The deck makes less mana on turn four than the norm, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- JUDGE [unknown]: "Smaug’s haste and flying make it an immediate threat". This asserts that the card Smaug the Impenetrable has the haste and flying abilities, i.e., what the card can do. Flying is plausible for a Dragon, but I cannot verify that this particular Smaug printing grants haste, so the combined claim is unverifiable here.
- PLAN plan_coherent=partly: The list is a broadly coherent Rakdos midrange pile — ramp, removal, goblin bodies and big finishers — but the dozen-plus Equipment package has almost no equipment payoffs or free-attach effects and pulls against the top-heavy dragon/troll plan, so the pieces only loosely serve one line.
- PLAN theme_fit=partly: Smaug the Impenetrable leads a red-black Hobbit-flavored list at a reasonable bracket-3 power, but the request limited non-land, non-rock cards to the Hobbit set and the deck still includes Orcish Bowmasters, Languish, Night's Whisper, Guttersnipe, Goblin Cratermaker, The One Ring and several LTR equipment.
- PLAN useful_as_built=yes: 37 lands with duals, fetches that reliably find Swamp/Mountain, a handful of rocks, plenty of interaction and multiple real win conditions make it a deck someone can shuffle up and play, even if it is a bit top-heavy and light on ramp.
- PLAN summary_honest=partly: It is admirably candid about the lack of a combo and the broad package, but it oversells 'Dragons' (only Smaug the Magnificent joins the commander), oversells Treasure production (a couple of sources), and calls a one-drop Gollum a finishing creature.
- [INFO] `curve_summary`: average mana value 3.13 over 62 nonland cards
- [INFO] `mana_pass`: the builder moved 2 cards of the mana base to bring the deck inside its power level
- [WARN] `outside_requested_set`: the sets you named do not hold 12 cards: Arid Mesa, Blood Crypt, Bloodstained Mire, Command Tower, Dragonskull Summit, Evolving Wilds, Exotic Orchard, Fabled Passage, Marsh Flats, Polluted Delta, Scalding Tarn, Terramorphic Expanse

<details><summary>The deck list</summary>

- 12 Swamp | land | Basic black source for the deck’s core mana.
- 12 Mountain | land | Basic red source for the deck’s core mana.
- 1 Blood Crypt | land | Dual-color land for reliable early access to both colors.
- 1 Bloodstained Mire | land | Fetch land that finds the deck’s basic color sources.
- 1 Command Tower | land | Reliable dual-color commander land.
- 1 Dragonskull Summit | land | Dual-color land supporting the red-black base.
- 1 Exotic Orchard | land | Flexible color access in multiplayer games.
- 1 Arid Mesa | land | Fetch land that finds a Mountain or Blood Crypt.
- 1 Marsh Flats | land | Fetch land that finds a Swamp or Blood Crypt.
- 1 Polluted Delta | land | Fetch land that finds a Swamp or Blood Crypt.
- 1 Scalding Tarn | land | Fetch land that finds a Mountain or Blood Crypt.
- 1 Fabled Passage | land | Fetch land that fixes either core color.
- 1 Evolving Wilds | land | Budget-neutral fixing for either basic color.
- 1 Terramorphic Expanse | land | Additional basic-land fixing.
- 1 Arcane Signet | ramp | Efficient two-color mana development.
- 1 Bag End Banquet | ramp | Artifact mana support for the larger threats.
- 1 Burn, Burn, Tree and Fern | ramp | Ramp piece that advances the mana plan.
- 1 Dragon's Desire | ramp | Dedicated mana development for the Dragon plan.
- 1 Mox Amber | ramp | Cheap mana acceleration alongside the legendary commander.
- 1 Orcrist, Goblin-cleaver | ramp | Equipment-based mana support.
- 1 Smaug the Magnificent | ramp | Additional Dragon ramp for reaching the deck’s top end.
- 1 The Misty Mountains Cold | ramp | Ramp saga supporting the mana-heavy strategy.
- 1 The Reaver Cleaver | ramp | Treasure-focused equipment ramp.
- 1 Balin, Loremaster | draw | Creature-based card advantage.
- 1 Key to the Side-Door | draw | Artifact card advantage piece.
- 1 Night's Whisper | draw | Efficient black card draw.
- 1 Palantír of Orthanc | draw | Repeatable artifact card advantage.
- 1 Rage into the Valley | draw | Red card-draw support.
- 1 Ragged Short Spear | draw | Equipment that contributes to card flow.
- 1 The Sackville-Bagginses | draw | Creature-based card advantage.
- 1 Thrór's Map | draw | Artifact card advantage that fits the equipment package.
- 1 Azog, Moria's Ruin | removal | Creature-based answer to opposing permanents.
- 1 Bilbo's Deadly Slice | removal | Cheap targeted removal.
- 1 Bitter Downfall | removal | Flexible black removal.
- 1 Bolg of the North | removal | Creature removal attached to a threat.
- 1 Fire of Orthanc | removal | Red targeted removal.
- 1 Goblin Cratermaker | removal | Versatile creature-based removal.
- 1 Improvised Club | removal | Low-cost instant-speed removal.
- 1 Orcish Bowmasters | removal | Efficient creature-based removal.
- 1 Smite the Deathless | removal | Additional instant-speed removal.
- 1 Bilbo's Ring | interaction | Protective and disruptive equipment option.
- 1 Dwarven Mattock | interaction | Equipment-based interaction.
- 1 Mithril Coat | interaction | Protection for important creatures.
- 1 My Precious // Allure of Power | interaction | Flexible legendary equipment interaction.
- 1 The One Ring | interaction | Powerful legendary defensive interaction.
- 1 Call Forth the Tempest | wipe | Broad reset against developed boards.
- 1 Desolation of Smaug | wipe | Theme-fitting board reset.
- 1 Languish | wipe | Efficient black creature sweeper.
- 1 Supper for Spiders | synergy | The deck’s dedicated synergy card.
- 1 Dreaded Bat-Cloud | threat | Black creature that adds battlefield pressure.
- 1 Fearsome Goblin Pair | threat | Goblin creature presence for attacking pressure.
- 1 Goblin-town Flunkies | threat | Early Goblin battlefield threat.
- 1 Great Goblin, Foul-Hearted | threat | Legendary Goblin that reinforces the aggressive creature plan.
- 1 Great Ugly-Looking Goblin // Clap! Snap! | threat | Goblin threat with additional flexibility.
- 1 Gundabad Opportunist | threat | Goblin Rogue that helps fill out the attacking force.
- 1 Guttersnipe | threat | Red creature threat that rewards the spell package.
- 1 Haunt of the Dead Marshes | threat | Black creature pressure for the midgame.
- 1 Misty Mountains Raider | threat | Goblin attacker for the creature base.
- 1 Orcish Siegemaster | threat | Orc creature that adds to the combat plan.
- 1 Ravening Warg | threat | Black creature threat for pressuring opponents.
- 1 Stone-Giant of High Pass | threat | Large red creature that supports combat finishes.
- 1 Down, Down to Goblin-town | wincon | Saga finisher for closing longer games.
- 1 Gollum, Riddle Master | wincon | Legendary finisher that gives the deck another closing threat.
- 1 Troll of Khazad-dûm | wincon | Large finishing creature for ending games through combat.
- 1 Along the Crooked Way | other | Theme-supporting enchantment utility.
- 1 Andúril, Flame of the West | other | Legendary equipment for the creature package.
- 1 Andúril, Narsil Reforged | other | Additional legendary equipment support.
- 1 Bothersome Noisemaker | other | Goblin creature utility.
- 1 Gathering of Darkness | other | Black thematic utility spell.
- 1 Getaway Barrel | other | Artifact utility for the board.
- 1 Glamdring | other | Legendary equipment supporting equipped attackers.
- 1 Goblin Plate Mail | other | Equipment that supports the Goblin contingent.
- 1 Gollum the Abandoned | other | Additional black legendary creature utility.
- 1 Long-Lost Lances | other | Equipment support for the creature suite.
- 1 Well-Worn Spatula | other | Additional artifact equipment utility.
- 1 Wayfarer's Bauble | ramp | the mana pass added it to bring the mana base inside the power level
- 1 Mount Doom | land | the mana pass added it to bring the mana base inside the power level

</details>

### 22. a set family and a card from outside it

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 175 names.

Commander: Thranduil, the Elvenking.

Grade: baseline, score 0.05, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $329.35 to buy, $329.35 the whole deck.

**Summary:** This Thranduil deck develops a dense Elven board, uses its commander as the centerpiece for a legendary- and graveyard-focused game, and keeps cards flowing while answering key opposing pieces. It closes through sustained creature pressure, with Troll of Khazad-dûm, Witch-king of Angmar, and Gollum, Riddle Master serving as the biggest end-game threats. The tradeoff is a strongly creature-centered plan that is less flexible when the board and graveyard are repeatedly disrupted.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core of the list is a coherent Mirkwood-elf go-wide board with ramp, card draw and removal, but three sweepers (Languish, Gnashing of Teeth, Raise the Palisade) actively fight the creature-dense plan, and the 'graveyard-focused' angle the summary names is barely supported by the cards.
- PLAN theme_fit=partly: It is a bracket-3-feeling Commander deck led by the requested Thranduil with Sol Ring kept, and the bulk of the list is Mirkwood/Hobbit-set material, but a noticeable slice (Elvish Mystic, Elvish Archdruid, Wood Elves, Elvish Visionary, Night's Whisper, Languish, Merciless Executioner, Mox Amber, Arcane Signet) comes from outside the named set.
- PLAN useful_as_built=partly: 36 lands plus ramp and a reasonable curve mean it functions, but the mana base is entirely basics in three colors while the deck asks for {G}{G}, {U}{U} and {B}{B} plus a four-color-pip commander, and the self-damaging sweepers make some draws work against the main plan.
- PLAN summary_honest=partly: The elf board, card flow, removal and the three named big finishers are all really there, but the claimed graveyard focus is not carried by the list, and the summary omits that the deck runs board wipes that punish its own wide board.
- [INFO] `curve_summary`: average mana value 3.00 over 63 nonland cards
- [WARN] `profile_off_band`: the worst color's share of the sources it needs is 0.76, and bracket 3 wants 0.85 or more (sources of requirement: U 14.75 of 19, B 17.5 of 23, G 15 of 19)
- [INFO] `mana_pass`: the builder moved 3 cards of the mana base to bring the deck inside its power level
- [WARN] `outside_requested_set`: the sets you named do not hold 1 card: Sol Ring

<details><summary>The deck list</summary>

- 12 Forest | land | Provides reliable green mana for the Elf-heavy core.
- 11 Island | land | Provides reliable blue mana for Thranduil and blue support spells.
- 13 Swamp | land | Provides reliable black mana for removal and finishers.
- 1 Sol Ring | ramp | Kept as requested and accelerates the deck's larger plays.
- 1 Arcane Signet | ramp | Flexible color fixing and early acceleration.
- 1 Elvish Mystic | ramp | A cheap Elf mana source that also supports Thranduil's Elf focus.
- 1 Delighted Halfling | ramp | Early acceleration for the deck's legendary-heavy plan.
- 1 Elvish Archdruid | ramp | Elf-based mana production that rewards building a broad Elf board.
- 1 Wayfarer's Bauble | ramp | Reliable early mana development.
- 1 Wood Elves | ramp | Develops mana while adding another Elf to the deck's creature base.
- 1 Woodland Weavemaster | ramp | Elf-themed mana acceleration.
- 1 Thranduil's Company | ramp | Builds mana while staying within the Elven theme.
- 1 Elvish Visionary | draw | An Elf that contributes immediate card flow.
- 1 Fateful Discovery | draw | Provides ongoing card access.
- 1 Hithlain Knots | draw | Instant-speed card flow.
- 1 Ithilien Kingfisher | draw | Creature-based card advantage.
- 1 Key to the Side-Door | draw | Artifact card advantage that does not rely on the creature board.
- 1 Lórien Revealed | draw | Blue card access that helps keep resources flowing.
- 1 Night's Whisper | draw | Efficient black card advantage.
- 1 Palantír of Orthanc | draw | A persistent source of card advantage.
- 1 Plunder the Trollshaws | draw | Adds green card flow to the deck.
- 1 Uncover the Moon-Letters | draw | Enchantment-based card access.
- 1 Bilbo's Deadly Slice | removal | Low-cost spot removal.
- 1 Bitter Downfall | removal | Flexible black removal.
- 1 Crude Bent Blade | removal | Equipment-based removal that remains useful alongside creatures.
- 1 Giant's Boulder | removal | Artifact removal option for the deck's broad answer suite.
- 1 Merciless Executioner | removal | Creature-based sacrifice removal.
- 1 Orcish Bowmasters | removal | Efficient creature-based removal pressure.
- 1 The Black Arrow | removal | Legendary Equipment that supplies another removal tool.
- 1 Troll Negotiations | removal | Additional black removal coverage.
- 1 Uneasy Partings | removal | A final flexible removal slot.
- 1 Confusticate and Bebother | interaction | Blue interaction for opposing plays.
- 1 Elrond, Moon-Reader | interaction | A legendary Elf interaction piece that also works with Thranduil.
- 1 Bilbo's Ring | interaction | Legendary artifact interaction.
- 1 Mithril Coat | interaction | Protective interaction for key creatures.
- 1 My Precious // Allure of Power | interaction | Flexible legendary-artifact interaction.
- 1 Old Fat Spider Can't See Me | interaction | Persistent interaction from an enchantment.
- 1 Stern Scolding | interaction | Cheap blue stack interaction.
- 1 Thranduil's Decree | interaction | Themed instant-speed interaction.
- 1 Gnashing of Teeth | wipe | One of the deck's broad battlefield reset options.
- 1 Languish | wipe | Efficient black creature sweep.
- 1 Raise the Palisade | wipe | A typal-friendly reset that can spare a chosen creature type.
- 1 Galion, Elvenking's Butler | synergy | A legendary Elf that supports Thranduil's central plan.
- 1 Thranduil, Sindarin Liege // Silvan Rally | synergy | A legendary Elf closely aligned with the deck's commander.
- 1 Boughside Wanderers | synergy | Adds another Elf card to support the graveyard-facing theme.
- 1 Mirkwood Meditator | synergy | An Elf that deepens the creature and graveyard core.
- 1 Mirkwood Nurturer | synergy | An Elf body that reinforces the deck's tribal density.
- 1 Mirkwood Pathmaker | synergy | Adds to the concentrated Elf package.
- 1 Supper for Spiders | synergy | A dedicated synergy card from the shortlist.
- 1 Attercop | threat | A substantial creature threat for pressuring opponents.
- 1 Cantankerous Keepers | threat | Creature pressure within the deck's Elven creature base.
- 1 Elven Raft-Steerer | threat | An Elf threat that adds to the board presence plan.
- 1 Elvenking's Harper | threat | An Elf creature that contributes to the deck's pressure.
- 1 Galadhrim Guide | threat | A thematic Elf attacker.
- 1 Grey Havens Navigator | threat | An Elf creature that helps establish a board.
- 1 Guardian of the Halls | threat | A sturdy Elf presence for the creature plan.
- 1 Haunt of the Dead Marshes | threat | A Nightmare Elf threat that also supports the Elf count.
- 1 Lothlórien Lookout | threat | An Elf threat for building board pressure.
- 1 Nimrodel Watcher | threat | Another Elf creature for the deck's combat plan.
- 1 Mirkwood Elk | threat | A creature threat outside the core Elf package.
- 1 Willow-Wind | threat | A final creature threat for board development.
- 1 Troll of Khazad-dûm | wincon | A marked finisher that provides a direct closing threat.
- 1 Witch-king of Angmar | wincon | A marked finisher for closing stalled games.
- 1 Gollum, Riddle Master | wincon | A marked finisher that also fits the deck's card-flow plan.
- 1 Mox Amber | ramp | the mana pass added it to bring the mana base inside the power level
- 1 Pelargir Survivor | ramp | the mana pass added it to bring the mana base inside the power level

</details>

### 23. two set families at once

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 203 names.

Commander: Kíli the Resourceful.

Grade: typical, score 0.41, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $284.94 to buy, $284.94 the whole deck.

**Summary:** Kíli the Resourceful leads a mono-white Dwarf-and-Equipment deck that develops a board of artifacts, legendary permanents, and Dwarves before turning its creatures into equipped attackers. The deck sustains itself with numerous draw permanents, clears blockers through broad removal and board wipes, and closes through upgraded combat threats, Angel of the Ruins, Jazal Goldmane, or Psychosis Crawler. It gives up multicolor flexibility and relies on establishing a board before its strongest turns.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

This deck plan aims for 2 or more finishers, the cards that win the game, and the deck holds 1.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: There is a real equipment/artifact core with white removal and draw supporting it, but the list sprawls into unconnected pieces (Maskwood Nexus, Academy Manufactor, Psychosis Crawler, Karn, Dusk//Dawn which punishes the very equipped 3+ power attackers the plan builds), so the pieces do not all pull one way.
- PLAN theme_fit=no: The one hard limit the person named was that cards come from the Hobbit and Bloomburrow sets, and the large majority of the list is generic Commander staples and cards from unrelated sets (Sol Ring, Skullclamp, Swords to Plowshares, Sun Titan, Blade Splicer, Skyclave Apparition, Karn, Helm of the Host, Psychosis Crawler, and so on).
- PLAN useful_as_built=yes: 36 lands in a mono-white deck with eight cheap rocks/ramp, a low curve, ten-plus removal spells, three wipes, and plenty of card draw makes a list that casts its spells and functions at the table as it stands.
- PLAN summary_honest=partly: The draw, removal, and equipment claims are borne out and the finisher shortfall is disclosed, but calling it a "Dwarf-and-Equipment" deck oversells a list with only about four Dwarves and no Dwarf payoffs, and it never admits the set restriction was abandoned.
- [INFO] `curve_summary`: average mana value 2.79 over 63 nonland cards
- [WARN] `finisher_short`: the finisher count is 1, and this deck plan wants 2 or more (Angel of the Ruins)

<details><summary>The deck list</summary>

- 30 Plains | land | Provides the dense white mana base needed for Kíli and the deck’s white spells.
- 1 Command Tower | land | Provides reliable commander-color mana.
- 1 Minas Tirith | land | Adds a themed white land slot.
- 1 Castle Ardenvale | land | Provides a white land with late-game utility.
- 1 Evolving Wilds | land | Finds a Plains while contributing to the land count.
- 1 Fabled Passage | land | Finds a Plains while improving land access.
- 1 Terramorphic Expanse | land | Finds a Plains while supporting a stable mana base.
- 1 Caretaker's Talent | draw | Supplies a dedicated card-advantage engine.
- 1 Circuit Mender | draw | Provides a card while adding an artifact body.
- 1 Dawn of a New Age | draw | Adds an enduring source of card advantage.
- 1 Errand-Rider of Gondor | draw | Contributes a creature-based draw option.
- 1 Esgaroth Garrison | draw | Provides card advantage on a creature.
- 1 Fountainport Bell | draw | Adds an artifact-based draw piece for Kíli’s plan.
- 1 Heirloom Epic | draw | Provides thematic artifact card advantage.
- 1 Idol of Oblivion | draw | Offers efficient artifact-based card advantage.
- 1 Inspiring Overseer | draw | Adds reliable creature-based card advantage.
- 1 Mentor of the Meek | draw | Rewards the deck’s many smaller creatures with cards.
- 1 Skullclamp | draw | Provides a highly efficient Equipment-based draw outlet.
- 1 Baird, Steward of Argive | interaction | Discourages opposing attacks and protects the developing board.
- 1 Bilbo's Gambit | interaction | Provides a flexible instant-speed interactive option.
- 1 Bilbo's Ring | interaction | Adds a legendary Equipment that also protects a key creature.
- 1 Bofur, Reliable Guardian // Concerted Care | interaction | Provides a Dwarf body and a useful interactive Adventure.
- 1 Crumb and Get It | interaction | Supplies a reactive spell for protecting the board position.
- 1 Dawn's Truce | interaction | Helps preserve the board through opposing pressure.
- 1 Galadriel's Dismissal | interaction | Provides flexible instant-speed disruption.
- 1 Swiftfoot Boots | interaction | Protects Kíli or an important equipped attacker.
- 1 Sol Ring | ramp | Provides the deck’s most efficient artifact acceleration.
- 1 Mox Amber | ramp | Uses the commander and legendary permanents to accelerate mana.
- 1 Arcane Signet | ramp | Provides dependable early mana acceleration.
- 1 Fellwar Stone | ramp | Adds efficient artifact mana acceleration.
- 1 Mind Stone | ramp | Accelerates early while retaining later utility.
- 1 Ornithopter of Paradise | ramp | Provides early mana and counts as an artifact for Kíli’s plan.
- 1 Patchwork Banner | ramp | Supports the Dwarf plan while adding mana.
- 1 Wayfarer's Bauble | ramp | Turns early mana into a lasting Plains-based mana boost.
- 1 Loyal Warhound | ramp | Helps catch up on lands with a low-cost creature.
- 1 Orcrist, Goblin-cleaver | ramp | Adds a themed Equipment that also supports mana development.
- 1 Banishing Light | removal | Provides broad permanent-based removal.
- 1 Bumbleflower's Sharepot | removal | Adds artifact-based removal to the deck’s permanent count.
- 1 Fiend Hunter | removal | Provides creature-based removal.
- 1 Generous Gift | removal | Answers a wide range of opposing permanents.
- 1 Hanged Executioner | removal | Adds a creature that can convert into removal.
- 1 Loran of the Third Path | removal | Provides a versatile removal creature.
- 1 Skyclave Apparition | removal | Offers efficient creature-based removal.
- 1 Sonar Strike | removal | Provides a compact removal spell.
- 1 Swords to Plowshares | removal | Supplies efficient single-target removal.
- 1 Blade Splicer | synergy | Builds the artifact-creature board while supporting Kíli’s artifact count.
- 1 Carrot Cake | synergy | Adds a low-cost artifact permanent to establish Kíli’s story.
- 1 Dáin, Lord of the Iron Hills | synergy | Supports the deck’s Dwarf-centered creature package.
- 1 Iron Hills Blacksmith | synergy | Connects the Dwarf and artifact portions of the deck.
- 1 Tangle Tumbler | synergy | Provides an artifact permanent that advances the equipment theme.
- 1 Academy Manufactor | synergy | Adds another artifact creature for the board-building plan.
- 1 Maskwood Nexus | synergy | Helps unify the creature suite around the Dwarf theme.
- 1 Fíli the Pathfinder | threat | Provides a named Dwarf threat for the commander’s creature plan.
- 1 Karn, the Great Creator | threat | Adds a resilient noncreature threat.
- 1 Andúril, Flame of the West | threat | Gives an attacker a legendary Equipment upgrade.
- 1 Andúril, Narsil Reforged | threat | Adds another legendary Equipment threat.
- 1 Glamdring | threat | Provides a legendary Equipment for the attacking plan.
- 1 Helm of the Host | threat | Adds a high-impact Equipment threat.
- 1 Long-Lost Lances | threat | Supplies another Equipment for building dangerous attackers.
- 1 Starforged Sword | threat | Adds a substantial Equipment threat to the board.
- 1 Sting, Bilbo's Sword | threat | Provides another legendary Equipment attacker upgrade.
- 1 Sword of Vengeance | threat | Turns a creature into a more meaningful combat threat.
- 1 Serra Redeemer | threat | Provides a standalone creature threat.
- 1 Sun Titan | threat | Provides a powerful top-end creature threat.
- 1 Angel of the Ruins | wincon | Serves as a marked finisher for closing games.
- 1 Jazal Goldmane | wincon | Provides a combat-focused route to end the game.
- 1 Psychosis Crawler | wincon | Turns the deck’s card-advantage plan into a finishing threat.
- 1 Dusk // Dawn | wipe | Provides a flexible board reset.
- 1 Martial Coup | wipe | Resets the battlefield while leaving behind a board presence.
- 1 Promise of Loyalty | wipe | Breaks up established opposing creature boards.

</details>

### 24. a 60-card deck from one set

Format: Modern. Theme: aggro. Pool: any_card. Shortlist: 96 names.

Grade: bad, score 0.33, model 20260914T154223Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $22.58 to buy, $22.58 the whole deck.

**Summary:** This mono-red aggro deck aims to establish pressure with Emberheart Challenger, Hearthborn Battler, Frilled Sparkshooter, Reptilian Recruiter, and Teapot Slinger, then keep attacks flowing while Agate Assault and Rabid Gnaw clear the way. Dragonhawk, Fate's Tempest gives the deck a heavier threat for games that go longer, while Might of the Meek and Sazacap's Brew help it keep finding action. It gives up broad answers and late-game flexibility in exchange for a direct, creature-forward red game plan.

The quality model grades this deck below the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. The color sources cover the pips, and that raises the grade.

- PLAN plan_coherent=yes: Every card supports one line — cheap red creatures backed by removal and cantrip pump spells to push damage through — with no off-plan inclusions, though the density of four- and five-drops makes it a slower aggro build than the summary implies.
- PLAN theme_fit=yes: It is a 60-card mono-red Modern creature-aggro deck built entirely from Bloomburrow cards, exactly as requested.
- PLAN useful_as_built=partly: 24 Mountains support the pips and the deck has plenty of creature-based win paths, but with twelve four-drops and six five-drops against only six two-drops the curve is far too top-heavy to function as a true aggro deck, so it will often durdle before its threats land.
- PLAN summary_honest=yes: Each named card is in the list and does what the summary says (removal, card flow, a late threat), and it openly admits the deck lacks broad answers and late-game flexibility; only the two Brazen Collectors go unmentioned.
- [INFO] `curve_summary`: average mana value 3.06 over 36 nonland cards

<details><summary>The deck list</summary>

- 24 Mountain | land | Keeps the mana base entirely red and supports fast, untapped starts.
- 4 Might of the Meek | draw | Low-cost card flow suited to an aggressive red shell.
- 2 Sazacap's Brew | draw | Adds more card flow without moving away from the deck's red plan.
- 2 Brazen Collector | ramp | Provides the requested mana acceleration while remaining part of the creature plan.
- 4 Agate Assault | removal | Efficiently fills the deck's need for opposing-board answers.
- 2 Rabid Gnaw | removal | Rounds out the removal suite with additional red interaction.
- 4 Emberheart Challenger | synergy | A key synergy creature for the deck's aggressive Bloomburrow core.
- 4 Hearthborn Battler | synergy | Reinforces the deck's creature-based synergy package.
- 4 Frilled Sparkshooter | threat | An aggressive red threat that helps apply early pressure.
- 4 Reptilian Recruiter | threat | Adds a full set of red creature threats for consistent attacks.
- 4 Teapot Slinger | threat | Provides another proactive threat in the Bloomburrow creature suite.
- 2 Dragonhawk, Fate's Tempest | threat | Supplies a larger top-end threat once the early assault has developed.

</details>

### 25. use no card of an owned precon

Format: Commander. Theme: Avengers superheroes. Pool: owned_first. Shortlist: 195 names.

Commander: Captain America, Super-Soldier.

Grade: baseline, score 0.32, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $228.15 the whole deck.

**Summary:** Captain America leads a white Avengers and S.H.I.E.L.D. team built around deploying Heroes, protecting the commander’s shield-backed board, and turning equipped attackers into steady pressure and card flow. The deck plays a patient combat game, clears away key obstacles, and closes with The Sentry, Golden Guardian, Origin of Spider-Man, or Angel of Serenity once the table has been worn down. Its tradeoff is a relatively fair, creature-and-Equipment-focused plan that relies on maintaining a board rather than explosive combo finishes.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: The list is one consistent mono-white Heroes/Equipment plan — cheap Hero bodies, ten Equipment plus Puresteel Paladin and Weapons Vendor to carry them, commander-protection pieces (Champion's Helm, Darksteel Plate, Bastion Protector, boots/greaves), and white removal — with the three symmetrical wipes softened by Unbreakable Formation and Clever Concealment.
- PLAN theme_fit=yes: It is a mono-white Commander deck built from Marvel Avengers/S.H.I.E.L.D. characters with a builder-chosen Captain America at the helm and no combo finish, matching the requested theme and bracket, with only a handful of non-Marvel filler (Minas Tirith, Buster Sword, Crown of Gondor) diluting flavor.
- PLAN useful_as_built=yes: 36 lands plus Sol Ring, Arcane Signet, Wayfarer's Bauble and other rocks support a low curve topping out at eight mana, and the deck wins straightforwardly through equipped Hero combat backed by ample removal and card draw.
- PLAN summary_honest=yes: Every element the summary names is on the list (Captain America, S.H.I.E.L.D./Hero creatures, Equipment card flow, targeted removal, and the four named closers), and it openly admits the fair, board-dependent, non-combo nature of the deck.
- [INFO] `curve_summary`: average mana value 3.03 over 63 nonland cards

<details><summary>The deck list</summary>

- 27 Plains | land | Provides dependable white mana for the deck’s largely white Hero and Equipment plan.
- 1 Command Tower | land | Provides flexible commander-focused mana.
- 1 City of Brass | land | Provides access to white mana without slowing the opening turns.
- 1 Marsh Flats | land | Finds a Plains to keep white mana reliable.
- 1 Fabled Passage | land | Finds a Plains while thinning toward action.
- 1 Evolving Wilds | land | Finds a Plains for stable early mana.
- 1 Terramorphic Expanse | land | Finds a Plains for stable early mana.
- 1 Minas Tirith | land | Adds a white-producing land with a Heroic flavor fit.
- 1 Secluded Courtyard | land | Supports the deck’s many Human Heroes.
- 1 Avengers Tower | land | Adds a thematic headquarters to the mana base.
- 1 Sol Ring | ramp | Accelerates the deck into its Heroes, Equipment, and commander.
- 1 Arcane Signet | ramp | Provides efficient mana acceleration.
- 1 Wayfarer's Bauble | ramp | Develops mana while preserving a high Plains count.
- 1 Sword of the Animist | ramp | Turns equipped attackers into ongoing mana development.
- 1 Chromatic Lantern | ramp | Provides reliable mana acceleration.
- 1 Thought Vessel | ramp | Supplies low-cost mana development.
- 1 Springleaf Drum | ramp | Converts an early Hero into mana acceleration.
- 1 Ring of the Lucii | ramp | Adds a compact mana-producing artifact.
- 1 Inherited Envelope | ramp | Provides additional artifact-based mana development.
- 1 White Lotus Tile | ramp | Adds another low-investment ramp piece.
- 1 Agent 13, Sharon Carter | draw | A Hero who helps keep cards flowing.
- 1 Agent Maria Hill | draw | A Hero that supplies card advantage.
- 1 Buster Sword | draw | An Equipment-based source of cards for equipped attackers.
- 1 Champions of Minas Tirith | draw | Contributes a creature body while helping refill the hand.
- 1 Crown of Gondor | draw | An Equipment that supports the deck’s card flow.
- 1 Hero in Training | draw | A Hero that advances the deck while providing cards.
- 1 Idol of Oblivion | draw | A low-cost artifact source of repeatable card advantage.
- 1 Mask of Memory | draw | Rewards combat with equipped Heroes by finding more action.
- 1 Origin of the Avengers | draw | A thematic Avengers card that provides card advantage.
- 1 Puresteel Paladin | draw | Supports the Equipment package while drawing cards.
- 1 Skullclamp | draw | A highly efficient Equipment source of cards.
- 1 The Vision | draw | An Avengers Hero that contributes card advantage.
- 1 Viv Vision, Teen Synthezoid | draw | A Heroic threat that keeps the hand supplied.
- 1 Weapons Vendor | draw | Supports the Equipment plan while adding card advantage.
- 1 Bastion Protector | interaction | Helps preserve the commander and the protected Hero board.
- 1 Champion's Helm | interaction | Protects a key legendary Hero, especially the commander.
- 1 Clever Concealment | interaction | Shields the developed board from opposing answers.
- 1 Darksteel Plate | interaction | Keeps an important Hero or commander on the battlefield.
- 1 Lightning Greaves | interaction | Protects a key creature immediately after deployment.
- 1 Patriot, Shield Wielder | interaction | A thematic Hero that helps defend the team.
- 1 Swiftfoot Boots | interaction | Adds another efficient layer of protection for key Heroes.
- 1 Unbreakable Formation | interaction | Protects the team while supporting a decisive attack.
- 1 Banishing Light | removal | Handles a troublesome opposing permanent.
- 1 Crib Swap | removal | Provides flexible creature removal.
- 1 Dispatch | removal | Offers efficient removal alongside the artifact-heavy Equipment package.
- 1 Fiend Hunter | removal | Places creature removal on a body that can wear Equipment.
- 1 Generous Gift | removal | Answers nearly any problematic permanent.
- 1 Get Lost | removal | Efficiently removes a troublesome nonland permanent.
- 1 Glider Staff | removal | A flavorful Equipment-based answer that fits the superhero battle plan.
- 1 March of Otherworldly Light | removal | Scales to answer important opposing threats.
- 1 Palace Jailer | removal | Combines removal with a creature body for the Equipment plan.
- 1 Swords to Plowshares | removal | Efficiently answers an opposing creature.
- 1 Stroke of Midnight | removal | Provides a broad answer to problematic permanents.
- 1 Web Up | removal | A superhero-flavored answer for opposing threats.
- 1 Agent Phil Coulson | synergy | A key S.H.I.E.L.D. Hero that ties the team together.
- 1 Agent of Atlas | synergy | Expands the Hero-heavy superhero roster.
- 1 Agents of S.H.I.E.L.D. | synergy | Adds another on-theme S.H.I.E.L.D. piece to the Hero team.
- 1 Avengers Quinjet | synergy | A thematic Avengers artifact that supports the team plan.
- 1 Colleen Wing, Street Samurai | synergy | A Hero that supports the deck’s combat-focused team.
- 1 Night Nurse, Healer of Heroes | synergy | Directly supports the deck’s Hero roster.
- 1 Quake, Agent of S.H.I.E.L.D. | synergy | A thematic Hero that strengthens the S.H.I.E.L.D. contingent.
- 1 Captain Mar-Vell, Space-Born | threat | An Avengers Hero that supplies a substantial battlefield threat.
- 1 Invisible Woman, Sue Storm | threat | A superhero threat that broadens the team’s roster.
- 1 Luke Cage, Power Man | threat | A durable superhero threat for the combat plan.
- 1 Mockingbird, Ace Agent | threat | A Heroic attacker that fits the S.H.I.E.L.D. theme.
- 1 Okoye, Dora Milaje Leader | threat | A combat-ready Hero that adds pressure to the board.
- 1 Roaming Throne | threat | Strengthens the deck’s concentrated Hero creature plan.
- 1 Angel of Serenity | wincon | A large finisher that can close games after the team establishes control.
- 1 Origin of Spider-Man | wincon | A superhero-themed finisher for converting a developed board into a win.
- 1 The Sentry, Golden Guardian | wincon | A major Avengers finisher and the deck’s chief closing threat.
- 1 Austere Command | wipe | A flexible reset that can spare the parts of the board most important to the plan.
- 1 Fumigate | wipe | Resets creature-heavy boards when combat has stalled.
- 1 Vanquish the Horde | wipe | Provides a practical board reset after opponents overextend.

</details>

