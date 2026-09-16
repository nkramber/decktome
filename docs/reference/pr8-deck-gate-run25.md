# PR-8 deck gate

Run date: 2026-09-16. Card snapshot: 2026-09-04.

Verdict: FAIL. 25 of 25 decks passed every block check, 0 invented names reached the user, and 1 summaries stated a false rule of the game. All three bars are zero tolerance.

## Summary

| Measure | Value |
|---|---|
| Prompts | 25 |
| Decks returned | 25 |
| Decks with no block finding | 25 |
| Invented names that reached the user | 0 |
| Decks that needed the repair turn | 1 |
| Summaries judged (F-26) | 25 |
| Summaries that state a rule of the game | 1 |
| Summaries that state a FALSE rule | 1 |
| Judge errors | 0 |
| Case assertions missed (PR-28b) | 0 |
| Decks the plan judge read (PR-15, information) | 25 |
| Mean plan score, 0 to 1 | 0.70 |
| Plan reasons the judge left empty | 2 |
| Errors | 0 |
| Prompt version | 14 |
| Calls | 76 |
| Cost | $2.5900 |
| Time | 2050 seconds |

## Run

- Suite `decks`, run `pr8-deck-gate-run25`, on 2026-09-16, commit `6ff0244`.
- Roles: generate on `gpt-5.6-terra` (openai, effort medium), judge on `claude-opus-5` (anthropic, effort medium), repair on `gpt-5.6-terra` (openai, effort medium).
- Versions: card snapshot 2026-09-04, generate prompt version 14, plan_rubric prompt version 2, precons `5.3.0+20260903`, quality_model `20260914T154223Z`.
- Calls: 76. Cost: $2.5900. Time: 2050 seconds.

## Findings by code

A block stops the deck. A warning and a note are reports.

| Code | Count |
|---|---|
| `curve_summary` | 25 |
| `not_owned` | 23 |
| `profile_off_band` | 10 |
| `mana_pass` | 8 |
| `cards_trimmed` | 4 |
| `basics_added` | 3 |
| `bracket_cut` | 2 |
| `outside_requested_set` | 2 |
| `finisher_short` | 2 |

By severity: BLOCK 0. WARN 37. INFO 42. 

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

Grade: bad, score 0.03, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $816.83 to buy, $816.83 the whole deck.

**Summary:** Karlov leads a creature-forward lifegain deck that builds a high life total while turning repeated gains into larger attackers and steady pressure. The deck protects its key permanents, uses focused removal and board wipes to keep opponents from getting ahead, and closes through combat or dedicated life-total victory cards such as Angel of Destiny, Felidar Sovereign, and Test of Endurance. It gives up some speed for a broad board presence and needs its lifegain engines to stay active for its strongest turns.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. Many of the cards appear in no top list, and that lowers the grade.

Summary rules claims (F-26): the format play

- JUDGE [true]: "...closes through combat or dedicated life-total victory cards such as Angel of Destiny, Felidar Sovereign, and Test of Endurance.". This asserts what these cards do: all three provide alternate win conditions keyed to a high life total (Felidar Sovereign at 40+ life, Test of Endurance at 50+, Angel of Destiny causing attacked players to lose when you have 15 life more than your starting total). The characterization is accurate.
- JUDGE [true]: "Karlov leads a creature-forward lifegain deck.". Reads as a description of this deck's configuration rather than a statement of format rules, but insofar as it implies Karlov of the Ghost Council can be a commander, that is correct since it is a legendary creature.
- PLAN plan_coherent=yes: Nearly every nonland card either gains life, rewards lifegain (Karlov counters, pridemate effects, drain payoffs like Sanguine Bond/Vito/Enduring Tenacity), or protects the board, so the pieces all feed one lifegain-into-pressure plan.
- PLAN theme_fit=yes: It is a 100-card Karlov-led mono-theme lifegain Commander deck in W/B, and the mix of a few strong cards (Solitude, Vein Ripper, Aetherflux) with many casual lifegain payoffs sits reasonably in bracket 3.
- PLAN useful_as_built=yes: 37 lands with a solid dual/fetch base for a two-color deck, ten mana rocks/ramp, eleven card-draw pieces, a low-to-mid curve, and multiple win routes (combat, drain combos like Heliod+Vito-style drains, and three alt-wins) make it playable out of the box.
- PLAN summary_honest=yes: The claimed elements — lifegain engines, protection (Alseid, Faith's Shield, Sword of Light and Shadow), removal and three wipes, and the named alt-win cards — are all actually present in the list, and it fairly notes the deck's dependence on lifegain engines.
- [INFO] `curve_summary`: average mana value 3.24 over 62 nonland cards
- [INFO] `bracket_cut`: to hold bracket 3, the builder cut 1 card: Enduring Angel // Angelic Enforcer (the combo Enduring Angel // Angelic Enforcer + Enduring Tenacity), and added 1 basic land

<details><summary>The deck list</summary>

- 1 Command Tower | land | A reliable black-or-white source for Karlov's mana base.
- 1 Godless Shrine | land | A dual source that supports both halves of the deck.
- 1 Scrubland | land | An efficient dual source for early white and black spells.
- 1 Caves of Koilos | land | An untapped dual source for the two deck colors.
- 1 City of Brass | land | Flexible color fixing for demanding multicolor turns.
- 1 Mana Confluence | land | Flexible color fixing for the lifegain shell.
- 1 Concealed Courtyard | land | A fast dual source for the early game.
- 1 Vault of Champions | land | A dual land that supports multiplayer starts well.
- 1 Shattered Sanctum | land | A dual source that helps cast the deck's core spells.
- 1 Isolated Chapel | land | A dual land that builds both color counts.
- 1 Brightclimb Pathway // Grimclimb Pathway | land | A flexible early land choice for either deck color.
- 1 Silent Clearing | land | A black-and-white source with useful late-game flexibility.
- 1 Marsh Flats | land | Fetches the basic land type needed in the moment.
- 1 Prismatic Vista | land | Finds either basic color while keeping the mana base stable.
- 12 Plains | land | Core white sources for Karlov and the deck's white spells.
- 11 Swamp | land | Core black sources for the deck's black spells.
- 1 Legion's Landing // Adanto, the First Fort | ramp | A low-cost ramp piece that fits the creature-focused plan.
- 1 Oasis Gardener | ramp | Early mana support for developing the board ahead of schedule.
- 1 Hot Dog Cart | ramp | A low-cost mana piece that helps reach Karlov and follow-up plays.
- 1 Altar of the Pantheon | ramp | Color support and ramp for the two-color mana base.
- 1 Nuka-Cola Vending Machine | ramp | Mana acceleration that advances the deck's resources.
- 1 Orazca Relic | ramp | A mana artifact that keeps the deck progressing into its midgame.
- 1 Pristine Talisman | ramp | Reliable mana acceleration for the lifegain strategy.
- 1 The Celestus | ramp | A ramp piece that supports steady development.
- 1 Phial of Galadriel | ramp | Additional mana acceleration for higher-impact plays.
- 1 Potioner's Trove | ramp | A ramp slot that helps pay for the deck's top end.
- 1 Archivist of Oghma | draw | Efficient card draw to keep resources flowing.
- 1 Dawn of Hope | draw | A draw engine suited to a deck built around gaining life.
- 1 Enduring Innocence | draw | Creature-based card draw that supports the board plan.
- 1 Exemplar of Light | draw | A lifegain-friendly draw creature that adds to the board.
- 1 Mangara, the Diplomat | draw | A durable source of card advantage.
- 1 Markov Purifier | draw | A Vampire draw option that fits the deck's creature base.
- 1 Sigarda's Splendor | draw | A lifegain-oriented source of continuing card advantage.
- 1 The Gaffer | draw | A creature draw engine for sustained value.
- 1 Tymna the Weaver | draw | A low-cost creature that helps refill the hand.
- 1 Well of Lost Dreams | draw | Converts the deck's lifegain focus into card advantage.
- 1 Vampiric Rites | draw | A flexible draw outlet for a creature-heavy deck.
- 1 Alseid of Life's Bounty | interaction | A low-cost protection option for Karlov and key permanents.
- 1 Faith's Shield | interaction | Protects an important threat at a crucial moment.
- 1 Metropolis Reformer | interaction | A protective creature that reinforces the lifegain board.
- 1 Restoration Magic | interaction | A flexible response that protects the deck's position.
- 1 Paladin Danse, Steel Maverick | interaction | A creature-based interaction slot for maintaining board presence.
- 1 Werefox Bodyguard | interaction | A flexible creature answer that keeps pressure on opponents.
- 1 Sword of Light and Shadow | interaction | An equipment interaction piece for protecting and supporting threats.
- 1 Aetherflux Reservoir | removal | A lifegain payoff that doubles as a decisive removal option.
- 1 Ayli, Eternal Pilgrim | removal | A lifegain-themed removal creature with utility on the battlefield.
- 1 Gumdrop Poisoner // Tempt with Treats | removal | A flexible removal spell that fits the deck's colors.
- 1 Murderous Rider // Swift End | removal | A versatile answer that remains useful as a creature afterward.
- 1 Nightmare's Thirst | removal | Cheap creature removal for early opposing threats.
- 1 Solitude | removal | A high-quality creature answer that can also join the board.
- 1 Umezawa's Jitte | removal | Reusable equipment-based removal support.
- 1 Vein Ripper | removal | A powerful Vampire removal threat for the late game.
- 1 Witch of the Moors | removal | A lifegain-aligned removal threat that pressures opposing boards.
- 1 Amalia Benavides Aguirre | wipe | A compact board-wipe option tied to the deck's central theme.
- 1 Fumigate | wipe | A reset button that naturally belongs in a lifegain deck.
- 1 The Meathook Massacre | wipe | A board wipe that remains relevant after clearing creatures.
- 1 Ajani's Pridemate | synergy | A classic lifegain payoff that becomes a meaningful combat threat.
- 1 Bloodthirsty Aerialist | synergy | A growing evasive threat fueled by repeated life gain.
- 1 Heliod, Sun-Crowned | synergy | A central lifegain synergy permanent for building stronger creatures.
- 1 Resplendent Angel | synergy | A lifegain payoff that adds board presence.
- 1 Sanguine Bond | synergy | A key lifegain payoff that turns the deck's plan into pressure.
- 1 Speaker of the Heavens | synergy | An early lifegain synergy creature that rewards a high life total.
- 1 Vito, Thorn of the Dusk Rose | synergy | A lifegain payoff that helps convert life into opponent pressure.
- 1 Archangel of Thune | threat | A premier lifegain threat that strengthens the creature board.
- 1 Attended Healer | threat | A creature threat that fits naturally alongside recurring life gain.
- 1 Bloodbond Vampire | threat | A scalable Vampire threat for the lifegain plan.
- 1 Doctor Jane Foster | threat | A standalone legendary threat that broadens the creature suite.
- 1 Enduring Tenacity | threat | A resilient threat that supports the deck's life-focused pressure.
- 1 Gideon's Company | threat | A creature threat that benefits from the deck's lifegain focus.
- 1 Qala, Ajani's Pridemate | threat | A lifegain-linked threat that reinforces the deck's combat plan.
- 1 Rhox Faithmender | threat | A major lifegain threat that amplifies the deck's core resource.
- 1 Star Charter | threat | A creature threat that adds depth to the lifegain creature plan.
- 1 Twinblade Paladin | threat | A combat threat that rewards the deck's ongoing life gain.
- 1 Vampire Scrivener | threat | A Vampire threat that contributes to the deck's pressure.
- 1 Wax-Wane Witness | threat | A creature threat that rounds out the deck's board presence.
- 1 Angel of Destiny | wincon | A lifegain-oriented finisher that gives the deck a clear closing route.
- 1 Felidar Sovereign | wincon | A dedicated alternate victory threat for games where life totals climb high.
- 1 Test of Endurance | wincon | A dedicated alternate victory route that rewards the deck's central plan.

</details>

### 2. aristocrats Commander, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 181 names.

Commander: Denethor, Ruling Steward.

Grade: typical, score 0.40, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $158.92 the whole deck.

**Summary:** This is a black-white aristocrats deck built around establishing a creature board, using its sacrifice-focused synergy package to gain value, and protecting the key pieces that keep the plan moving. It maintains a broad set of answers while drawing through equipment, artifacts, and black spells, then closes through its dedicated finishers once the table has been worn down. The tradeoff is a straightforward two-color mana base with few utility lands, prioritizing dependable early development and timely access to Denethor over specialized land effects.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=no: The list pulls three ways at once — equipment/commander-protection voltron (Puresteel Paladin, Champion's Helm, Bastion Protector, Sword of the Animist), a control shell with nine removal spells and three board wipes, and a handful of loosely related creatures — with almost no sacrifice outlets or death-trigger payoffs to tie anything together, and the wipes actively fight the stated goal of establishing a creature board.
- PLAN theme_fit=no: The user asked for an aristocrats deck and got a black-white goodstuff/equipment-protection pile under an aristocrats-capable commander, with essentially none of the token generation, sacrifice outlets, or drain payoffs the archetype requires.
- PLAN useful_as_built=partly: The two-color basics-plus-Command-Tower mana base and curve are perfectly functional and the removal suite plays fine game to game, but the deck has no real closing plan — a few midsized creatures and card draw with no drain engine or combat finisher — so it can be picked up and played yet struggles to actually win.
- PLAN summary_honest=no: It advertises a "sacrifice-focused synergy package" and "dedicated finishers" that do not exist in the list (no Blood Artist-style drains, no free sac outlets, and the "wincon" slots are Al Bhed Salvagers, Archfiend of Ifnir, and Grave Venerations), and it omits the three sweepers that undercut the creature plan it describes.
- [INFO] `curve_summary`: average mana value 2.79 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Plains | land | Provides a reliable white mana source.
- 1 Plains | land | Provides a reliable white mana source.
- 1 Plains | land | Provides a reliable white mana source.
- 1 Plains | land | Provides a reliable white mana source.
- 1 Plains | land | Provides a reliable white mana source.
- 1 Plains | land | Provides a reliable white mana source.
- 1 Plains | land | Provides a reliable white mana source.
- 1 Plains | land | Provides a reliable white mana source.
- 1 Plains | land | Provides a reliable white mana source.
- 1 Plains | land | Provides a reliable white mana source.
- 1 Plains | land | Provides a reliable white mana source.
- 1 Plains | land | Provides a reliable white mana source.
- 1 Plains | land | Provides a reliable white mana source.
- 1 Plains | land | Provides a reliable white mana source.
- 1 Plains | land | Provides a reliable white mana source.
- 1 Plains | land | Provides a reliable white mana source.
- 1 Plains | land | Provides a reliable white mana source.
- 1 Swamp | land | Provides a reliable black mana source.
- 1 Swamp | land | Provides a reliable black mana source.
- 1 Swamp | land | Provides a reliable black mana source.
- 1 Swamp | land | Provides a reliable black mana source.
- 1 Swamp | land | Provides a reliable black mana source.
- 1 Swamp | land | Provides a reliable black mana source.
- 1 Swamp | land | Provides a reliable black mana source.
- 1 Swamp | land | Provides a reliable black mana source.
- 1 Swamp | land | Provides a reliable black mana source.
- 1 Swamp | land | Provides a reliable black mana source.
- 1 Swamp | land | Provides a reliable black mana source.
- 1 Swamp | land | Provides a reliable black mana source.
- 1 Swamp | land | Provides a reliable black mana source.
- 1 Swamp | land | Provides a reliable black mana source.
- 1 Swamp | land | Provides a reliable black mana source.
- 1 Swamp | land | Provides a reliable black mana source.
- 1 Swamp | land | Provides a reliable black mana source.
- 1 Command Tower | land | Adds flexible colored mana to the mana base.
- 1 City of Brass | land | Adds another flexible colored mana source.
- 1 Sol Ring | ramp | Efficient early mana acceleration.
- 1 Arcane Signet | ramp | Reliable colored mana acceleration.
- 1 Fellwar Stone | ramp | Low-cost mana acceleration.
- 1 Commander's Sphere | ramp | Flexible mana acceleration for the midgame.
- 1 Lotho, Corrupt Shirriff | ramp | Creature-based mana acceleration for the deck's plan.
- 1 Relic of Legends | ramp | Supports mana development alongside legendary cards.
- 1 Springleaf Drum | ramp | Low-cost mana acceleration alongside the creature base.
- 1 Thought Vessel | ramp | Mana acceleration that remains useful later.
- 1 Wayfarer's Bauble | ramp | Early mana development from a low-cost artifact.
- 1 Sword of the Animist | ramp | Equipment-based mana development for attacking creatures.
- 1 Buster Sword | draw | Equipment-based card advantage.
- 1 Call of the Ring | draw | A dedicated card-advantage piece.
- 1 Circle of Power | draw | Adds a dedicated source of card advantage.
- 1 Foot Chopper | draw | Artifact-based card advantage.
- 1 Idol of Oblivion | draw | A compact artifact draw piece.
- 1 Lembas | draw | Low-cost artifact card advantage.
- 1 Mask of Memory | draw | Equipment-based card advantage for the creature plan.
- 1 Nasty End | draw | Black card advantage that fits the sacrifice-focused shell.
- 1 Night's Whisper | draw | Efficient straightforward card advantage.
- 1 Painful Truths | draw | A dedicated spell-based draw option.
- 1 Skullclamp | draw | Equipment-based card advantage that supports the aristocrats plan.
- 1 Bastion Protector | interaction | Protective interaction for the commander-focused game plan.
- 1 Boromir, Warden of the Tower | interaction | Creature-based disruptive interaction.
- 1 Clever Concealment | interaction | Protective interaction for maintaining the board.
- 1 Champion's Helm | interaction | Equipment protection for key legendary permanents.
- 1 Gift of Immortality | interaction | Protection that supports a sacrifice-oriented creature plan.
- 1 Lightning Greaves | interaction | Efficient equipment protection for key creatures.
- 1 Swiftfoot Boots | interaction | Additional equipment protection for important creatures.
- 1 Unbreakable Formation | interaction | Protective interaction for the creature board.
- 1 Bitter Triumph | removal | Versatile targeted removal.
- 1 Claim the Precious | removal | Dedicated black targeted removal.
- 1 Crib Swap | removal | Flexible creature removal.
- 1 Dismember | removal | Low-cost creature removal.
- 1 Fatal Push | removal | Efficient targeted removal.
- 1 Generous Gift | removal | Broad permanent removal.
- 1 Get Lost | removal | Flexible white targeted removal.
- 1 Infernal Grasp | removal | Reliable creature removal.
- 1 Swords to Plowshares | removal | Efficient creature removal.
- 1 Austere Command | wipe | Flexible board-clearing option.
- 1 Fumigate | wipe | Creature-board reset for stabilizing games.
- 1 Vanquish the Horde | wipe | Additional efficient creature-board reset.
- 1 Arcade Cabinet | synergy | A dedicated synergy piece for the aristocrats shell.
- 1 Gollum the Abandoned | synergy | Supports the deck's sacrifice-focused synergy package.
- 1 Gollum, Patient Plotter | synergy | Another dedicated piece for the aristocrats plan.
- 1 Gríma Wormtongue | synergy | Adds to the deck's dedicated synergy package.
- 1 Heirloom Auntie | synergy | Creature-based support for the aristocrats shell.
- 1 Joo Dee, One of Many | synergy | A dedicated creature synergy piece.
- 1 Phantom Train | synergy | Artifact support for the deck's synergy plan.
- 1 Ahriman | threat | A creature used to build battlefield pressure.
- 1 Bill the Pony | threat | A creature threat that advances the board presence plan.
- 1 Cirith Ungol Patrol | threat | A creature threat for applying pressure.
- 1 Exemplar of Light | threat | A creature threat that strengthens the board.
- 1 Inspiring Overseer | threat | A creature threat that contributes to board development.
- 1 June, Bounty Hunter | threat | A creature threat that supports a proactive battlefield.
- 1 Massacre Girl, Known Killer | threat | A black creature threat for the sacrifice-focused shell.
- 1 Puresteel Paladin | threat | A creature threat that works alongside the equipment package.
- 1 Rat King, Pale Piper | threat | A creature threat for the deck's proactive plan.
- 1 The Sackville-Bagginses | threat | A creature threat that adds battlefield presence.
- 1 Vengeful Villagers | threat | A creature threat suited to the aristocrats strategy.
- 1 Wall of Omens | threat | A creature threat that helps establish the board.
- 1 Al Bhed Salvagers | wincon | A designated finishing threat for closing games.
- 1 Archfiend of Ifnir | wincon | A designated finisher for ending stalled games.
- 1 Grave Venerations | wincon | A dedicated finishing piece for the aristocrats plan.

</details>

### 3. artifacts Commander, bracket 4

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 320 names.

Commander: Urza, Lord High Artificer.

Grade: good, score 0.66, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $7517.67 to buy, $7517.67 the whole deck.

**Summary:** This deck accelerates into a dense artifact board, keeps cards flowing, and uses efficient interaction and removal to protect its momentum. It pressures opponents with artifact creatures while closing through Mechanized Production, Mirrodin Besieged, or Thassa's Oracle. The deck gives up some resilience in slower hands that do not find early mana, and its concentrated artifact plan can be vulnerable when the board is repeatedly cleared.

The quality model grades this deck good against the top lists of the format: the deck makes more mana on turn four than the norm, and that raises the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

- PLAN plan_coherent=partly: The bulk of the list is a coherent Urza artifact plan — fast mana, cheap artifacts for affinity/construct size, tutors like Whir and Transmute, and artifact beaters backed by counterspells — but Thassa's Oracle sits outside it with no self-mill, Consultation, or library-emptying effect to make it a win, and a few filler picks (Reality Ripple, Into Thin Air, Cyber Conversion) are off-plan interaction.
- PLAN theme_fit=yes: It is exactly what was asked: mono-blue Urza, Lord High Artificer with a dense artifact shell, powerful accelerants (Workshop, Mana Vault, Moxen, Metalworker) and free counterspells, sitting comfortably in the high-power bracket-4 range.
- PLAN useful_as_built=yes: 33 lands plus a large fast-mana package in a nearly mono-colored base makes casting trivial, the curve is low and supported by affinity/improvise reducers, and the deck can genuinely win by attacking with Kappa Cannoneer, construct tokens, and Aetherflux Reservoir even without the named alt-wins.
- PLAN summary_honest=partly: placeholder
- [INFO] `curve_summary`: average mana value 3.00 over 65 nonland cards
- [INFO] `mana_pass`: the builder moved 1 card of the mana base to bring the deck inside its power level
- [INFO] `cards_trimmed`: the list was 1 card over, so the builder cut this card: Metalwork Colossus

<details><summary>The deck list</summary>

- 1 Academy Ruins | land | Provides a land slot in the artifact-focused mana base.
- 1 City of Brass | land | Provides a land slot in the mana base.
- 1 Command Tower | land | Provides a land slot in the mana base.
- 1 Exotic Orchard | land | Provides a land slot in the mana base.
- 1 Glimmervoid | land | Provides a land slot for an artifact-heavy deck.
- 1 Inventors' Fair | land | Provides a land slot in the artifact-focused mana base.
- 21 Island | land | Supplies the deck's basic land base.
- 1 Mana Confluence | land | Provides a land slot in the mana base.
- 1 Mishra's Workshop | land | Provides a land slot that supports the artifact plan.
- 1 Mystic Sanctuary | land | Provides a land slot in the mana base.
- 1 Otawara, Soaring City | land | Provides a land slot in the mana base.
- 1 Seat of the Synod | land | Provides an artifact land slot.
- 1 Spire of Industry | land | Provides a land slot for an artifact-heavy deck.
- 1 Urza's Saga | land | Provides a land slot that supports the artifact plan.
- 1 Chrome Mox | ramp | Supplies fast artifact mana.
- 1 Chief Engineer | ramp | Supplies artifact-focused mana acceleration.
- 1 Grand Architect | ramp | Supplies artifact-focused mana acceleration.
- 1 Inspiring Statuary | ramp | Supplies artifact-focused mana acceleration.
- 1 Krark-Clan Ironworks | ramp | Supplies artifact-focused mana acceleration.
- 1 Lion's Eye Diamond | ramp | Supplies fast artifact mana.
- 1 Lotus Petal | ramp | Supplies fast artifact mana.
- 1 Mana Vault | ramp | Supplies fast artifact mana.
- 1 Metalworker | ramp | Supplies artifact-focused mana acceleration.
- 1 Moonsnare Prototype | ramp | Supplies fast artifact mana.
- 1 Mox Diamond | ramp | Supplies fast artifact mana.
- 1 Mox Opal | ramp | Supplies fast artifact mana.
- 1 Sol Ring | ramp | Supplies fast artifact mana.
- 1 Springleaf Drum | ramp | Supplies fast artifact mana.
- 1 Braided Net // Braided Quipu | draw | Provides artifact-based card draw.
- 1 Era of Innovation | draw | Provides a draw engine for the artifact plan.
- 1 Esoteric Duplicator | draw | Provides artifact-based card draw.
- 1 Nexus of Becoming | draw | Provides artifact-based card draw.
- 1 One with the Machine | draw | Provides card draw for the artifact plan.
- 1 Reverse Engineer | draw | Provides card draw for the artifact plan.
- 1 Rhystic Study | draw | Provides a powerful draw engine.
- 1 Riddlesmith | draw | Provides card draw in an artifact-focused shell.
- 1 Sai, Master Thopterist | draw | Provides card draw for the artifact plan.
- 1 Thoughtcast | draw | Provides efficient card draw for the artifact plan.
- 1 Thirst for Knowledge | draw | Provides card draw for the artifact plan.
- 1 An Offer You Can't Refuse | interaction | Provides cheap stack interaction.
- 1 Disruption Protocol | interaction | Provides stack interaction for the artifact plan.
- 1 Fierce Guardianship | interaction | Provides premium interaction.
- 1 Force of Will | interaction | Provides premium interaction.
- 1 Ice Out | interaction | Provides stack interaction.
- 1 Metallic Rebuke | interaction | Provides artifact-focused stack interaction.
- 1 Override | interaction | Provides stack interaction.
- 1 Padeem, Consul of Innovation | interaction | Provides interaction for the artifact plan.
- 1 Reality Ripple | interaction | Provides flexible interaction.
- 1 Stoic Rebuttal | interaction | Provides stack interaction.
- 1 Welding Jar | interaction | Provides artifact-focused interaction.
- 1 Aether Spellbomb | removal | Provides artifact-based removal.
- 1 Aetherflux Reservoir | removal | Provides artifact-based removal.
- 1 Arcum Dagsson | removal | Provides removal tied to artifact tutoring.
- 1 Baral's Expertise | removal | Provides broad removal.
- 1 Contagion Clasp | removal | Provides artifact-based removal.
- 1 Cyber Conversion | removal | Provides targeted removal.
- 1 Into Thin Air | removal | Provides targeted removal.
- 1 Kitesail Larcenist | removal | Provides creature-based removal.
- 1 Resculpt | removal | Provides flexible removal.
- 1 Skysovereign, Consul Flagship | removal | Provides artifact-based removal.
- 1 Engineered Explosives | wipe | Provides a scalable board wipe.
- 1 Rebuild | wipe | Provides a board-reset option.
- 1 Mystic Forge | synergy | Supports the artifact synergy plan.
- 1 Transmute Artifact | synergy | Supports the artifact plan while serving as a tutor.
- 1 Whir of Invention | synergy | Supports the artifact plan while serving as a tutor.
- 1 Filigree Attendant | threat | Provides an artifact creature threat.
- 1 Foundry Assembler | threat | Provides an artifact creature threat.
- 1 Frogmite | threat | Provides an artifact creature threat.
- 1 Jhoira's Familiar | threat | Provides an artifact creature threat.
- 1 Kappa Cannoneer | threat | Provides a major artifact creature threat.
- 1 Master Transmuter | threat | Provides a high-impact artifact creature threat.
- 1 Myr Enforcer | threat | Provides an artifact creature threat.
- 1 Phyrexian Metamorph | threat | Provides a flexible artifact creature threat.
- 1 Sojourner's Companion | threat | Provides an artifact creature threat.
- 1 Traxos, Scourge of Kroog | threat | Provides a legendary artifact creature threat.
- 1 Whirler Rogue | threat | Provides a creature-based threat.
- 1 Mechanized Production | wincon | Provides an artifact-based finishing route.
- 1 Mirrodin Besieged | wincon | Provides an artifact-based finishing route.
- 1 Thassa's Oracle | wincon | Provides a compact finishing route.

</details>

### 4. dinosaur tribal, bracket 2

Format: Commander. Theme: dinosaurs. Pool: any_card. Shortlist: 270 names.

Commander: Gishath, Sun's Avatar.

Grade: baseline, score 0.07, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $978.20 to buy, $978.20 the whole deck.

**Summary:** This deck builds its mana early, develops a board of Dinosaurs, and uses large combat threats to pressure the table before closing with its marked finishers. Its removal and wipes give it ways to reset difficult boards, while its draw cards help it keep deploying creatures. It gives up some early speed for a straightforward, creature-heavy plan with expensive top-end plays.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The curve sits high for the format, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Every slot supports one Gishath dinosaur-tribal plan: cost reducers and ramp into fat dinosaurs, haste/protection to connect, and payoffs that reward big bodies and creature ETBs.
- PLAN theme_fit=partly: It is unmistakably the requested Gishath dinosaur Commander deck at roughly precon power, but the 'new player' part of the request sits awkwardly with a mana base built on fetches, original duals, and four pain/drawback lands including Forbidden Orchard, which gives opponents creatures.
- PLAN useful_as_built=yes: 34 lands plus nine ramp pieces and cost reducers support the heavy curve, the colors are well covered, and there are plenty of ways to close a game through Gishath triggers and oversized attackers.
- PLAN summary_honest=yes: The summary's claims all check out — there is meaningful ramp, a real draw suite, targeted removal plus two wipes, a creature-heavy top end, and it openly admits the slow start and high curve.
- [INFO] `curve_summary`: average mana value 3.90 over 59 nonland cards
- [INFO] `mana_pass`: the builder moved 1 card of the mana base to bring the deck inside its power level
- [INFO] `basics_added`: the list was 2 cards short, so the builder added 2 basic lands

<details><summary>The deck list</summary>

- 1 Beast Whisperer | draw | A draw slot that supports the creature-heavy plan.
- 1 Garruk's Uprising | draw | A draw slot for the Dinosaur-focused game plan.
- 1 Guardian Project | draw | A draw slot for sustained creature-based games.
- 1 Harmonize | draw | A straightforward draw spell for a new player.
- 1 Ripjaw Raptor | draw | A Dinosaur draw slot for the tribal plan.
- 1 Return of the Wildspeaker | draw | A draw slot that fits the large-creature plan.
- 1 Rishkar's Expertise | draw | A draw slot suited to a board of large creatures.
- 1 Runic Armasaur | draw | A Dinosaur draw slot that contributes to the creature base.
- 1 Curious Altisaur | draw | A Dinosaur draw slot for the tribal plan.
- 1 Earthshaker Dreadmaw | draw | A Dinosaur draw slot that remains a substantial creature.
- 1 Heroic Intervention | interaction | An interaction slot for protecting the developing board.
- 1 Flawless Maneuver | interaction | An interaction slot for protecting the creature plan.
- 1 Lightning Greaves | interaction | An interaction slot that supports an important creature.
- 1 Swiftfoot Boots | interaction | An interaction slot for supporting key creatures.
- 1 Akroma's Will | interaction | An interaction slot and marked finisher for closing a game.
- 1 Temple Altisaur | interaction | A Dinosaur interaction slot for the tribal plan.
- 1 Bronzebeak Foragers | removal | A Dinosaur removal slot that adds to the creature count.
- 1 Burning Sun's Avatar | removal | A Dinosaur removal slot with a large body.
- 1 Itzquinth, Firstborn of Gishath | removal | A low-cost Dinosaur removal slot.
- 1 Ravenous Sailback | removal | A Dinosaur removal slot for the creature-focused deck.
- 1 Savage Stomp | removal | A direct removal slot that fits the Dinosaur theme.
- 1 Territorial Allosaurus | removal | A Dinosaur removal slot with a meaningful body.
- 1 Trumpeting Carnosaur | removal | A Dinosaur removal slot for the top of the curve.
- 1 Austere Command | wipe | A flexible wipe slot for difficult board states.
- 1 Raging Swordtooth | wipe | A Dinosaur wipe slot that keeps the theme intact.
- 1 Sol Ring | ramp | A low-cost ramp slot for accelerating the deck.
- 1 Arcane Signet | ramp | A reliable ramp slot for the three-color mana plan.
- 1 Birds of Paradise | ramp | A low-cost ramp slot for early color fixing.
- 1 Llanowar Elves | ramp | A low-cost ramp slot for early development.
- 1 Delighted Halfling | ramp | A low-cost ramp slot for the creature plan.
- 1 Drover of the Mighty | ramp | A ramp slot that fits a Dinosaur deck.
- 1 Ranging Raptors | ramp | A Dinosaur ramp slot for building toward larger spells.
- 1 Thunderherd Migration | ramp | A Dinosaur-themed ramp slot for mana development.
- 1 Topiary Stomper | ramp | A Dinosaur ramp slot with a substantial body.
- 1 Otepec Huntmaster | synergy | A synergy slot for the Dinosaur tribal plan.
- 1 Kinjalli's Caller | synergy | A synergy slot that supports the Dinosaur-heavy curve.
- 1 Marauding Raptor | synergy | A Dinosaur synergy slot for the tribal plan.
- 1 Hunting Velociraptor | synergy | A Dinosaur synergy slot for the creature base.
- 1 Commune with Dinosaurs | synergy | A synergy slot that helps find the deck's theme pieces.
- 1 Dinosaur Stampede | synergy | A Dinosaur synergy slot for the combat plan.
- 1 Kinjalli's Sunwing | synergy | A Dinosaur synergy slot that supports the board plan.
- 1 Huatli's Snubhorn | synergy | A Dinosaur synergy slot for the tribal core.
- 1 Raptor Companion | synergy | A low-cost Dinosaur synergy slot.
- 1 Regal Imperiosaur | synergy | A Dinosaur synergy slot that rewards the creature base.
- 1 Carnage Tyrant | threat | A large Dinosaur threat for the primary combat plan.
- 1 Ghalta, Stampede Tyrant | threat | A high-impact Dinosaur threat for late-game turns.
- 1 Regisaur Alpha | threat | A Dinosaur threat that supports the tribal board.
- 1 Pantlaza, Sun-Favored | threat | A Dinosaur threat for the central creature plan.
- 1 Ghalta and Mavren | threat | A Dinosaur threat for building a powerful board.
- 1 Quartzwood Crasher | threat | A Dinosaur threat suited to combat-focused games.
- 1 Tyrranax Rex | threat | A large Dinosaur threat for finishing pressure.
- 1 Zetalpa, Primal Dawn | threat | A large Dinosaur threat at the top of the curve.
- 1 Thundering Spineback | threat | A Dinosaur threat for the late-game board plan.
- 1 Palani's Hatcher | threat | A Dinosaur threat that adds to the tribal creature base.
- 1 Rampaging Brontodon | threat | A Dinosaur threat for combat-focused turns.
- 1 Needletooth Pack | threat | A lower-cost Dinosaur threat that keeps the curve manageable.
- 1 Dinosaurs on a Spaceship | wincon | A marked Dinosaur finisher and win condition.
- 1 Etali, Primal Storm | wincon | A marked Dinosaur finisher and win condition.
- 1 Huatli, Poet of Unity // Roar of the Fifth People | wincon | A marked finisher and win condition for the deck.
- 1 Arid Mesa | land | A land slot for the three-color mana base.
- 1 Windswept Heath | land | A land slot for the three-color mana base.
- 1 Wooded Foothills | land | A land slot for the three-color mana base.
- 1 Prismatic Vista | land | A land slot for the three-color mana base.
- 1 Command Tower | land | A land slot for consistent commander colors.
- 1 City of Brass | land | A land slot for broad color access.
- 1 Mana Confluence | land | A land slot for broad color access.
- 1 Forbidden Orchard | land | A land slot for broad color access.
- 1 Starting Town | land | A land slot for the three-color mana base.
- 1 Tarnished Citadel | land | A land slot for broad color access.
- 1 Battlefield Forge | land | A land slot supporting red and white needs.
- 1 Needleverge Pathway // Pillarverge Pathway | land | A flexible land slot supporting red and white needs.
- 1 Plateau | land | A land slot supporting red and white needs.
- 1 Sacred Foundry | land | A land slot supporting red and white needs.
- 1 Bountiful Promenade | land | A land slot supporting green and white needs.
- 1 Branchloft Pathway // Boulderloft Pathway | land | A flexible land slot supporting green and white needs.
- 1 Canopy Vista | land | A land slot supporting green and white needs.
- 1 Savannah | land | A land slot supporting green and white needs.
- 1 Cinder Glade | land | A land slot supporting red and green needs.
- 1 Karplusan Forest | land | A land slot supporting red and green needs.
- 1 Spire Garden | land | A land slot supporting red and green needs.
- 1 Stomping Ground | land | A land slot supporting red and green needs.
- 7 Forest | land | Basic land slots supporting green requirements.
- 5 Plains | land | Basic land slots supporting white requirements.
- 6 Mountain | land | Basic land slots supporting red requirements.

</details>

### 5. blink Commander, owned first

Format: Commander. Theme: blink. Pool: owned_first. Shortlist: 169 names.

Commander: Gilraen, Dúnedain Protector.

Grade: baseline, score 0.31, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $195.29 the whole deck.

**Summary:** This is a creature-centered Gilraen blink deck that develops its mana, deploys value creatures and artifact support, then uses blink cards to keep the board working through removal and disruption. It wins by establishing a resilient battlefield and converting that advantage through its dedicated closing cards, led by Angel of Serenity. The deck gives up raw speed for a steady permanent-based game plan that is strongest when its creatures can remain on the table and be reused.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

This deck plan aims for 2 or more finishers, the cards that win the game, and the deck holds 1.

Summary rules claims (F-26): the format play

- JUDGE [false]: "\"led by Angel of Serenity\" (in a Gilraen Commander deck)". This asserts Angel of Serenity leads the Commander deck, i.e., serves as its commander. Angel of Serenity is a non-legendary creature and cannot be a commander; the deck's stated commander is Gilraen.
- PLAN plan_coherent=partly: The blink/ETB-value core (Flickerwisp, Angel of Condemnation, Fiend Hunter, Slip On the Ring, Wall of Omens, Inspiring Overseer) is real but thin, and it is pulled against by six mass sweepers (Austere Command, Dusk // Dawn, Fumigate, Martial Coup, Vanquish the Horde) plus a stray equipment sub-theme that a wide creature board does not want.
- PLAN theme_fit=partly: It is a mono-white Gilraen Commander deck at a reasonable bracket-3 power level, but the blink engine the player asked to headline is only lightly represented (a handful of flicker effects among mostly generic white removal, ramp and sweepers).
- PLAN useful_as_built=yes: Thirty-seven lands in mono-white with solid rocks, a low curve, ample removal and card draw means the deck casts its spells reliably and can close through creature beats and Angel of Serenity, even if only one card is a true finisher.
- PLAN summary_honest=partly: It correctly flags the single finisher and the permanent-based plan, but calling it a deck that "keeps the board working" while quietly omitting the five-plus board wipes that blow up its own creatures, and overstating the blink density, hides a real tension in the list.
- [INFO] `curve_summary`: average mana value 3.13 over 63 nonland cards
- [WARN] `finisher_short`: the finisher count is 1, and this deck plan wants 2 or more (Angel of Serenity)

<details><summary>The deck list</summary>

- 30 Plains | land | Primary white mana source for a consistent mana base.
- 1 Command Tower | land | Flexible color access for the commander deck.
- 1 Marsh Flats | land | Land slot that helps maintain white access.
- 1 Fabled Passage | land | Land slot that helps maintain white access.
- 1 Evolving Wilds | land | Land slot that helps maintain white access.
- 1 Minas Tirith | land | White-producing utility land slot.
- 1 Plaza of Heroes | land | Utility land that supports the legendary commander plan.
- 1 Sol Ring | ramp | Efficient early mana acceleration.
- 1 Arcane Signet | ramp | Reliable colored mana acceleration.
- 1 Fellwar Stone | ramp | Low-cost mana acceleration.
- 1 Thought Vessel | ramp | Mana acceleration in an artifact slot.
- 1 Wayfarer's Bauble | ramp | Early ramp that supports land development.
- 1 Sword of the Animist | ramp | Equipment-based ramp for the creature plan.
- 1 Relic of Legends | ramp | Mana support that works with the deck's legendary cards.
- 1 Chromatic Lantern | ramp | Mana acceleration and color fixing.
- 1 Bender's Waterskin | ramp | Additional artifact mana acceleration.
- 1 White Lotus Tile | ramp | Additional ramp for reaching higher-cost plays.
- 1 Adventurer's Airship | draw | Card advantage attached to a permanent.
- 1 Buster Sword | draw | Equipment-based card advantage.
- 1 Champions of Minas Tirith | draw | Creature-based card advantage for the battlefield plan.
- 1 Crown of Gondor | draw | Card advantage from an equipment slot.
- 1 Diary of Dreams | draw | Dedicated card-advantage artifact.
- 1 Exemplar of Light | draw | Creature-based card advantage.
- 1 Faramir, Field Commander | draw | Legendary creature that contributes card advantage.
- 1 Idol of Oblivion | draw | Low-cost artifact card advantage.
- 1 Inspiring Overseer | draw | Creature-based card advantage that fits the blink plan.
- 1 Lembas | draw | Card advantage on a reusable artifact permanent.
- 1 Mask of Memory | draw | Combat-based card advantage.
- 1 Angel of Sanctions | removal | Creature-based removal that fits the creature-centered plan.
- 1 Banishing Light | removal | Permanent-based answer to opposing threats.
- 1 Crib Swap | removal | Efficient instant-speed removal.
- 1 Destroy Evil | removal | Flexible instant-speed removal.
- 1 Dispatch | removal | Cheap targeted removal.
- 1 Generous Gift | removal | Broad answer to troublesome permanents.
- 1 Get Lost | removal | Flexible targeted removal.
- 1 Swords to Plowshares | removal | Efficient targeted removal.
- 1 Stroke of Midnight | removal | Instant-speed permanent removal.
- 1 Bastion Protector | interaction | Protective interaction for the commander plan.
- 1 Boromir, Warden of the Tower | interaction | Creature-based defensive interaction.
- 1 Clever Concealment | interaction | Protection against opposing disruption.
- 1 Lightning Greaves | interaction | Protects an important creature immediately.
- 1 Reprieve | interaction | Flexible stack interaction.
- 1 Swiftfoot Boots | interaction | Protects key creatures while supporting attacks.
- 1 Unbreakable Formation | interaction | Protective interaction for the creature board.
- 1 Ultimate Magic: Holy | interaction | Additional instant-speed interaction.
- 1 Angel of Condemnation | synergy | Blink-focused creature that reinforces the central plan.
- 1 Ennis, Debate Moderator | synergy | Dedicated synergy creature for the deck's plan.
- 1 Fiend Hunter | synergy | Creature permanent that rewards repeated blink support.
- 1 Flickerwisp | synergy | Core blink card for reusing permanent-based value.
- 1 Jocasta, Automaton Avenger | synergy | Legendary synergy piece for the creature shell.
- 1 Personify | synergy | Blink support for the deck's creature permanents.
- 1 Slip On the Ring | synergy | Instant-speed blink support and protection.
- 1 Bronze Guardian | threat | Artifact creature that adds board presence.
- 1 Frontline Medic | threat | Creature that contributes to the attacking board.
- 1 Giada, Font of Hope | threat | Angel creature that builds the board presence.
- 1 Kataki, War's Wage | threat | Creature threat that pressures artifact-heavy boards.
- 1 Palace Jailer | threat | Creature that advances the battlefield plan.
- 1 Puresteel Paladin | threat | Creature that supports the deck's equipment package.
- 1 South Pole Voyager | threat | Creature threat that supports the board-focused strategy.
- 1 The Vision | threat | Artifact creature that adds a durable body to the board.
- 1 The Walls of Ba Sing Se | threat | Artifact creature that strengthens the creature board.
- 1 Wall of Omens | threat | Low-cost creature that contributes to the board plan.
- 1 Weapons Vendor | threat | Creature that supports the equipment-focused elements.
- 1 Zack Fair | threat | Legendary creature that adds pressure and board presence.
- 1 Angel of Serenity | wincon | The shortlist's dedicated finishing Angel.
- 1 Martial Coup | wincon | Late-game closer that also resets the battlefield.
- 1 Vanquish the Horde | wincon | A sweeping reset that clears the way for the deck's board.
- 1 Austere Command | wipe | Flexible battlefield reset.
- 1 Dusk // Dawn | wipe | Creature-focused board reset with additional value.
- 1 Fumigate | wipe | Reliable full-board reset.

</details>

### 6. Modern tempo, tournament

Format: Modern. Theme: tempo. Pool: any_card. Shortlist: 167 names.

Grade: typical, score 0.52, model 20260914T154223Z.

Cards: 60 main, 15 sideboard. Repair turn: no. Block findings: 0.

Cost: $635.93 to buy, $635.93 the whole deck.

**Summary:** This blue-red tempo deck establishes pressure with Ragavan, Nimble Pilferer, Ledger Shredder, and Faerie Mastermind, then uses draw spells to keep finding removal and interaction while protecting its lead. Lightning Bolt and Galvanic Discharge clear opposing pressure, while Force of Negation and Spell Pierce help win crucial exchanges. Temporal Mastery and Temporal Trespass give the deck a powerful synergy finish once it has taken control of the pace, but the list gives up broad permanent coverage in favor of fast, focused pressure.

The quality model grades this deck typical against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade. The deck runs its spells as playsets, and that lowers the grade.

- PLAN plan_coherent=partly: The core of cheap evasive threats, burn, and cantrips is a real tempo plan, but the four extra-turn cards (Temporal Mastery with only 6 cantrips to set up a miracle, and Temporal Trespass in a deck that barely fills a graveyard) pull toward a payoff the rest of the list doesn't support.
- PLAN theme_fit=partly: It is a blue-red Modern tempo deck with the right colors and format, but the archetype's namesake and usual shell (Delver of Secrets, Dragon's Rage Channeler, Murktide, Counterspell) is absent in favor of extra-turn spells, so it is not quite the Delver deck asked for.
- PLAN useful_as_built=yes: The mana base is clean and plentiful for a curve topping at two, and four playsets of threats plus eight burn spells give enough ways to actually close a game even if the Temporal spells are frequently dead.
- PLAN summary_honest=partly: The threat, burn, and counterspell descriptions match the list, but calling the two Temporal cards a "powerful synergy finish" names a synergy the deck lacks (no delve enablers, too few cantrips for reliable miracle), and it glosses over the awkward 24-land, no-Delver build.
- [INFO] `curve_summary`: average mana value 2.33 over 36 nonland cards

<details><summary>The deck list</summary>

- 4 Consider | draw | Provides efficient card selection and supports the spell-heavy tempo plan.
- 2 Preordain | draw | Finds threats, lands, or answers early in the game.
- 4 Lightning Bolt | removal | A direct, efficient answer that clears the way for early pressure.
- 4 Galvanic Discharge | removal | Adds more efficient removal for opposing creatures.
- 4 Force of Negation | interaction | Protects the tempo plan without requiring mana to be held open.
- 2 Spell Pierce | interaction | Provides cheap protection against key opposing spells.
- 2 Temporal Mastery | synergy | An extra-turn synergy card that rewards setting up efficient turns.
- 2 Temporal Trespass | synergy | Provides another extra-turn payoff for the deck's spell-focused plan.
- 4 Ragavan, Nimble Pilferer | threat | A premier early threat that starts applying pressure immediately.
- 4 Ledger Shredder | threat | A flying threat that fits naturally beside a high density of spells.
- 4 Faerie Mastermind | threat | A flash threat that lets the deck maintain a reactive posture.
- 4 Steam Vents | land | Reliable blue-red dual land for the core mana base.
- 4 Spirebluff Canal | land | Fast blue-red source suited to an early-game tempo deck.
- 4 Riverglide Pathway // Lavaglide Pathway | land | Flexible untapped source for either primary color.
- 4 Shivan Reef | land | Untapped blue-red source that keeps early plays available.
- 2 Fiery Islet | land | Provides either color while occupying a land slot.
- 2 Thundering Falls | land | Additional blue-red dual land that maintains both color counts.
- 2 Island | land | Basic blue source for the deck's blue spells.
- 2 Mountain | land | Basic red source for the deck's red spells.

</details>

### 7. Modern burn, casual

Format: Modern. Theme: burn. Pool: any_card. Shortlist: 290 names.

Grade: typical, score 0.43, model 20260914T154223Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $36.78 to buy, $36.78 the whole deck.

**Summary:** This mono-red burn deck applies early pressure with its threat and synergy creatures while using direct removal to keep opposing creatures from slowing the attack. Card-draw spells help sustain the flow of action, and the deck closes through repeated burn pressure backed by its larger red threats. It gives up flexibility and specialized answers in exchange for a focused, all-red aggressive plan.

The quality model grades this deck typical against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade. The color sources cover the pips, and that raises the grade.

- PLAN plan_coherent=no: The list pulls in several directions at once: a pinger/Eidolon prowess-ish shell, a ramp spell (Spawning Breath), four copies of a marginal combat trick cantrip, and a top end of four-drop fatties, with no single coherent line tying them together.
- PLAN theme_fit=no: The person asked for mono-red burn, and while the colors are right, eight burn spells plus 24 lands and a creature-heavy midrange top end is a different deck than a burn list.
- PLAN useful_as_built=partly: The mana is stable and the creature curve can attack, so it functions at a casual table, but 24 lands with this low-impact spell suite means many flooded, slow games and no reliable reach to finish.
- PLAN summary_honest=no: It claims the deck "closes through repeated burn pressure" and leans on card draw, but the only real burn is four Lightning Bolt, two Lava Dart and two Risk Factor, and Ancestral Anger is a pump trick rather than a draw engine.
- [INFO] `curve_summary`: average mana value 2.56 over 36 nonland cards

<details><summary>The deck list</summary>

- 24 Mountain | land | Provides the mono-red mana base.
- 4 Ancestral Anger | draw | Supplies a dedicated draw package for the burn shell.
- 2 Risk Factor | draw | Adds further card access while fitting the burn plan.
- 2 Spawning Breath | ramp | Fills the requested acceleration slots.
- 4 Lightning Bolt | removal | Primary efficient burn removal.
- 2 Lava Dart | removal | Adds more low-cost removal to support the burn plan.
- 4 Eidolon of the Great Revel | synergy | A synergistic permanent for the deck's aggressive red strategy.
- 4 Thermo-Alchemist | synergy | Supports the spell-heavy burn approach.
- 4 Akki Lavarunner // Tok-Tok, Volcano Born | threat | Provides an early threat package.
- 4 Ashcloud Phoenix | threat | Adds a resilient threat presence.
- 4 Hazoret the Fervent | threat | Provides a powerful top-end threat.
- 2 Torbran, Thane of Red Fell | threat | Rounds out the threat suite with a red-focused finisher.

</details>

### 8. Modern lifegain, FNM

Format: Modern. Theme: lifegain. Pool: any_card. Shortlist: 301 names.

Grade: bad, score 0.35, model 20260914T154223Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $398.72 to buy, $398.72 the whole deck.

**Summary:** This white-black lifegain deck establishes early lifegain synergies, then converts that momentum into pressure through growing creatures and a Vampire threat package. Its removal clears the way for the board to attack, while its draw cards help it continue developing through longer FNM games. The deck gives up a broad suite of specialized interaction in favor of committing heavily to its lifegain plan.

The quality model grades this deck below the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade.

- PLAN plan_coherent=yes: Nearly every card feeds or exploits lifegain — Guide of Souls, Attended Healer, and Altar generate life triggers while Pridemate, Bloodbond Vampire, Scion of the Swarm, Vito, and Dawn of Hope convert them into board presence, with removal and Lembas supporting that single axis.
- PLAN theme_fit=yes: It is a straightforward white-black lifegain deck in Modern, and while Solitude and Sheoldred push toward the high end, the shell is still an FNM-appropriate synergy deck rather than a different archetype.
- PLAN useful_as_built=yes: 24 lands with a full set of duals and painlands supports the two-color curve, the threats top out at four mana, and there are multiple credible win routes (wide lifegain creatures, Vito drain, Sheoldred) plus real removal, so it can be picked up and played.
- PLAN summary_honest=yes: The claims of lifegain engines, growing creatures, a small vampire subtheme (Bloodbond Vampire, Vito), removal, and card draw all correspond to cards actually in the list, and it openly admits it skimps on varied interaction.
- [INFO] `curve_summary`: average mana value 3.28 over 36 nonland cards

<details><summary>The deck list</summary>

- 6 Plains | land | Provides reliable white mana for the deck’s white cards.
- 6 Swamp | land | Provides reliable black mana for the deck’s black cards.
- 4 Godless Shrine | land | Supplies both deck colors without taking a slot from the spell plan.
- 4 Caves of Koilos | land | Supplies flexible white or black mana.
- 4 Shattered Sanctum | land | Adds further white-black color coverage.
- 2 Altar of the Pantheon | ramp | Helps accelerate into the deck’s more expensive spells.
- 4 Solitude | removal | Forms the primary removal package while keeping the deck proactive.
- 2 Murderous Rider // Swift End | removal | Adds flexible removal attached to a creature card.
- 4 Lembas | draw | Provides efficient card flow for a lifegain-focused shell.
- 2 Dawn of Hope | draw | Gives the deck a dedicated source of card advantage.
- 4 Guide of Souls | synergy | Starts the lifegain synergy plan early in the game.
- 2 Ajani's Pridemate | synergy | Rewards the deck for repeatedly pursuing its lifegain plan.
- 2 Vito, Thorn of the Dusk Rose | synergy | Turns the deck’s lifegain theme into a more threatening board presence.
- 4 Bloodbond Vampire | threat | Serves as a scalable creature threat for the lifegain strategy.
- 4 Attended Healer | threat | Builds a board presence while supporting the deck’s main theme.
- 4 Scion of the Swarm | threat | Provides a substantial Vampire threat that fits the lifegain plan.
- 2 Sheoldred, the Apocalypse | threat | Supplies a powerful top-end threat for longer games.

</details>

### 9. Standard midrange, FNM

Format: Standard. Theme: midrange. Pool: any_card. Shortlist: 164 names.

Grade: bad, score 0.31, model 20260914T154223Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $130.60 to buy, $130.60 the whole deck.

**Summary:** This black-green midrange deck develops a creature board, backs it with efficient removal, and uses draw to keep resources flowing into later turns. It wins by applying steady pressure with threats ranging from early creatures to large finishers, while its protective synergy cards help keep important pieces in play. It gives up some speed for a more resource-focused game plan and relies on its creatures to convert its advantages into a win.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The cards pair in ways the top lists do not, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The removal, card draw, and beefy top end do point at one midrange plan, but eight copies of one-mana protection tricks (Snakeskin Veil, Royal Treatment) is a hexproof-shell package with only 14 creatures to protect and no aura/voltron payoff, so a chunk of the list pulls away from the grindy resource game the rest of the cards play.
- PLAN theme_fit=yes: It is a two-color black-green Standard deck built around creatures, removal, and card advantage at a casual FNM power level, which is exactly what was asked for.
- PLAN useful_as_built=partly: 24 lands with dual lands and a dork support a curve that tops at six, and the fat threats give clear ways to win, but the eight protection spells are frequently dead cards when no creature is on board, making the deck prone to clunky, threat-light draws.
- PLAN summary_honest=yes: Every claim — creature board, efficient removal, draw engine, large finishers, protection cards, slower resource-focused plan — maps to cards actually in the list, and it openly concedes the speed tradeoff.
- [INFO] `curve_summary`: average mana value 2.50 over 36 nonland cards
- [INFO] `mana_pass`: the builder moved 2 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 8 Forest | land | Provides green mana without entering tapped.
- 12 Swamp | land | Provides black mana without entering tapped.
- 4 Overgrown Tomb | land | Supplies both deck colors while strengthening the mana base.
- 2 Llanowar Elves | ramp | Early green ramp supports the midrange curve.
- 4 Phyrexian Arena | draw | Provides a dedicated draw package for longer games.
- 2 Garruk's Uprising | draw | Adds draw that fits the creature-heavy plan.
- 4 Bitter Triumph | removal | Efficient black removal for opposing threats.
- 2 Maelstrom Pulse | removal | Flexible removal complements the instant-speed answers.
- 4 Snakeskin Veil | synergy | Supports key creatures and the board-focused plan.
- 4 Royal Treatment | synergy | Supports the creature suite during contested turns.
- 4 Darkstar Augur | threat | A black creature that helps pressure opponents in the midgame.
- 4 Goldvein Hydra | threat | A scalable green creature threat for the midrange plan.
- 2 Thrashing Brontodon | threat | A creature threat that also contributes to the deck's answer suite.
- 2 Vein Ripper | threat | A black creature threat at the top of the curve.
- 2 Vaultborn Tyrant | threat | A large creature threat for closing longer games.

</details>

### 10. Standard aggro, tournament

Format: Standard. Theme: aggro. Pool: any_card. Shortlist: 294 names.

Grade: bad, score 0.08, model 20260914T154223Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $63.40 to buy, $63.40 the whole deck.

**Summary:** This red-white aggro deck develops pressure with a focused creature base, uses Slickshot Show-Off to reinforce its offense, and keeps cards coming through Fugitive Codebreaker and Tersa Lightshatter. Boros Charm and Sheltered by Ghosts provide flexible support while the removal suite clears a path for attackers. It gives up broader late-game options in exchange for a direct, consistent plan built around fast red-white mana and repeated combat pressure.

The quality model grades this deck below the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. Many of the cards appear in no top list, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: Most of the list is a genuine RW aggro shell of cheap creatures, combat tricks, and removal, but 4x Krenko's Buzzcrusher at seven mana plus 4x Redcap Gutter-Dweller pull toward a midrange top end that the low-land, fast-clock plan does not support.
- PLAN theme_fit=partly: It is a Standard 60-card red-white creature deck as asked, but the heavy top end makes it read more like a clunky midrange build than the aggro deck requested.
- PLAN useful_as_built=partly: The mana base (24 lands with Sacred Foundry and Inspiring Vantage) is clean and there are plenty of threats and removal to play games, but four copies of a seven-drop and four five-drops mean a lot of dead early draws and hands that stall.
- PLAN summary_honest=partly: The summary fairly describes the creature-and-removal core, but its claim of a "direct, consistent plan built around fast red-white mana" quietly hides that a third of the nonland cards cost five or more, and the "keeps cards coming" card-advantage claim rests on two thin creatures.
- [INFO] `curve_summary`: average mana value 3.06 over 36 nonland cards
- [INFO] `mana_pass`: the builder moved 1 card of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 9 Plains | land | White basic land that supplies a reliable white source.
- 7 Mountain | land | Red basic land that supplies a reliable red source.
- 4 Sacred Foundry | land | Red-white land that supports both colors.
- 4 Inspiring Vantage | land | Red-white land for the aggressive mana base.
- 4 Fugitive Codebreaker | draw | Low-end red card draw that keeps pressure flowing.
- 2 Tersa Lightshatter | draw | Red card draw that helps replenish the attack.
- 4 Boros Charm | interaction | Red-white interaction that supports the proactive game plan.
- 2 Sheltered by Ghosts | interaction | White interaction for protecting the attack plan.
- 4 Emeritus of Truce // Swords to Plowshares | removal | White removal that clears opposing obstacles.
- 4 Krenko's Buzzcrusher | removal | Red removal that keeps blockers from slowing the deck.
- 4 Slickshot Show-Off | synergy | Red synergy piece that reinforces the aggressive plan.
- 4 Frilled Sparkshooter | threat | Red threat that contributes to early pressure.
- 4 Redcap Gutter-Dweller | threat | Red threat for a concentrated aggressive creature base.
- 4 Dragonback Lancer | threat | Red-white threat that continues the deck's attack-focused plan.

</details>

### 11. Commander with a locked card

Format: Commander. Theme: sacrifice. Pool: any_card. Shortlist: 299 names.

Commander: Karlov of the Ghost Council.

Grade: baseline, score 0.06, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $763.02 to buy, $763.02 the whole deck.

**Summary:** This black-white sacrifice deck develops expendable creatures and sacrifice outlets, then converts those resources into mana, cards, removal, and steady pressure. It wins by leveraging death-focused payoffs alongside dedicated finishers and substantial creature threats. The deck gives up some speed to creature-dependent setup, so protecting key engines and choosing when to commit to the board are important.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. Few lands enter tapped, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every card is a sacrifice outlet, an expendable body/token maker, or a death-trigger payoff, and the mana rocks/altars double as sac-fuel, so the pieces all feed one aristocrats engine.
- PLAN theme_fit=yes: It is a black-white Commander sacrifice deck under Karlov, and the numerous lifegain drain payoffs (Zulaport Cutthroat, Elas il-Kor, Vindictive Vampire, Ayli, Vito's Inquisitor) even tie into Karlov's lifegain-counter ability.
- PLAN useful_as_built=yes: 36 lands with solid duals for a very low curve, plenty of cheap outlets and fodder, and several distinct win routes (combat threats, drain, mill/ping stations) make it immediately playable.
- PLAN summary_honest=yes: The claims about sac outlets, converting bodies into mana (Ashnod's/Phyrexian Altar, Pitiless Plunderer, Priest), cards (Village Rites, Vampiric Rites, Yawgmoth), removal (Grave Pact, Attrition) and finishers (Blasting Station, Altar of Dementia, Vindictive Vampire) are all backed by actual cards, and it admits the creature-dependence weakness.
- [INFO] `curve_summary`: average mana value 3.22 over 63 nonland cards
- [INFO] `mana_pass`: the builder moved 1 card of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 13 Plains | land | Provides reliable white mana.
- 15 Swamp | land | Provides reliable black mana.
- 1 Command Tower | land | Provides either commander color.
- 1 Godless Shrine | land | Provides both deck colors.
- 1 Scrubland | land | Provides both deck colors.
- 1 Caves of Koilos | land | Provides flexible colored mana.
- 1 Brightclimb Pathway // Grimclimb Pathway | land | Provides the needed color on either face.
- 1 Vault of Champions | land | Provides both deck colors.
- 1 Shattered Sanctum | land | Provides both deck colors.
- 1 Concealed Courtyard | land | Provides early colored mana.
- 1 Sol Ring | ramp | Supplies efficient mana acceleration.
- 1 Ashnod's Altar | ramp | Converts sacrifice resources into mana.
- 1 Phyrexian Altar | ramp | Turns sacrificed creatures into mana.
- 1 Pitiless Plunderer | ramp | Supports mana production through the sacrifice plan.
- 1 Priest of Forgotten Gods | ramp | Provides sacrifice-based mana acceleration.
- 1 Warren Soultrader | ramp | Contributes to the deck's mana engine.
- 1 Skullport Merchant | ramp | Provides a sacrifice-focused mana option.
- 1 Culling the Weak | ramp | Offers a burst of mana from a sacrifice.
- 1 Master of Dark Rites | ramp | Adds to the deck's mana production.
- 1 Crowded Crypt | ramp | Builds mana resources from creatures dying.
- 1 Corrupted Conviction | draw | Turns a sacrifice into cards.
- 1 Village Rites | draw | Efficiently converts a creature into cards.
- 1 Vampiric Rites | draw | Provides repeatable sacrifice-based card flow.
- 1 Baron Bertram Graywater | draw | Adds a creature-based draw source.
- 1 Disciple of Bolas | draw | Rewards sacrificing a larger creature with cards.
- 1 Bushmeat Poacher | draw | Provides card flow through creature sacrifices.
- 1 Lord Skitter's Butcher | draw | Supplies a sacrifice-oriented draw effect.
- 1 Shadowheart, Dark Justiciar | draw | Converts creatures into a substantial card refill.
- 1 Smothering Abomination | draw | Rewards the sacrifice plan with cards.
- 1 Tevesh Szat, Doom of Fools | draw | Provides sustained card advantage.
- 1 Thraxodemon | draw | Offers another creature-based draw outlet.
- 1 Cartel Aristocrat | interaction | Provides a protective sacrifice outlet.
- 1 Fanatical Devotion | interaction | Uses sacrifices to protect key creatures.
- 1 Flare of Fortitude | interaction | Protects the board during pivotal turns.
- 1 Gift of Doom | interaction | Protects an important permanent through a sacrifice.
- 1 Promise of Tomorrow | interaction | Helps preserve creatures through opposing pressure.
- 1 Spirit Bonds | interaction | Offers flexible protection for creatures.
- 1 Eldrazi Monument | interaction | Protects the creature board while fitting the sacrifice plan.
- 1 Nightmare Shepherd | interaction | Provides resilience against creature losses.
- 1 Attrition | removal | Uses expendable creatures as repeatable removal.
- 1 Ayli, Eternal Pilgrim | removal | Provides a sacrifice-based removal option.
- 1 Blasting Station | removal | Turns sacrifices into targeted removal.
- 1 Bone Shards | removal | Pairs a sacrifice cost with efficient removal.
- 1 Eaten Alive | removal | Uses expendable creatures to answer threats.
- 1 Grave Pact | removal | Makes creature sacrifices pressure opposing boards.
- 1 Dictate of Erebos | removal | Rewards each creature sacrifice with removal pressure.
- 1 Yawgmoth, Thran Physician | removal | Provides repeatable creature control.
- 1 Teysa, Orzhov Scion | removal | Adds a creature-based removal tool.
- 1 Toxic Deluge | wipe | Resets creature-heavy boards efficiently.
- 1 Austere Command | wipe | Provides a flexible board reset.
- 1 The Meathook Massacre | wipe | Combines a board reset with the deck's sacrifice plan.
- 1 Altar of Dementia | synergy | Provides a sacrifice outlet that advances the plan.
- 1 Carrion Feeder | synergy | Supplies a low-cost repeatable sacrifice outlet.
- 1 Bartolomé del Presidio | synergy | Acts as a sacrifice outlet and growing payoff.
- 1 Fleshtaker | synergy | Rewards creatures leaving the battlefield.
- 1 Elas il-Kor, Sadistic Pilgrim | synergy | Provides a sacrifice-focused life-drain payoff.
- 1 Zulaport Cutthroat | synergy | Turns creature deaths into table-wide pressure.
- 1 Viscera Seer | synergy | Provides a cheap repeatable sacrifice outlet.
- 1 Basri's Lieutenant | threat | Adds a resilient creature threat.
- 1 Corpse Harvester | threat | Provides a substantial creature threat.
- 1 Felisa, Fang of Silverquill | threat | Rewards the deck for losing key creatures.
- 1 Ghoulcaller Gisa | threat | Creates a threatening creature-focused board presence.
- 1 Liesa, Forgotten Archangel | threat | Adds a powerful evasive threat.
- 1 Mondrak, Glory Dominus | threat | Amplifies the deck's creature-making potential.
- 1 Ratadrabik of Urborg | threat | Builds lasting value from legendary creature losses.
- 1 Requiem Angel | threat | Provides an evasive top-end threat.
- 1 Sadistic Hypnotist | threat | Creates major pressure from disposable creatures.
- 1 Titan Hunter | threat | Provides a large sacrifice-oriented threat.
- 1 Vindictive Vampire | threat | Makes creature deaths pressure opponents.
- 1 Vito's Inquisitor | threat | Adds another creature threat to the board.
- 1 Razaketh, the Foulblooded | wincon | Serves as a dedicated finishing threat.
- 1 Relic Vial | wincon | Provides a dedicated route to close the game.
- 1 Sephiroth, Fabled SOLDIER // Sephiroth, One-Winged Angel | wincon | Supplies a high-impact finishing route.

</details>

### 12. Commander on a budget

Format: Commander. Theme: tokens. Pool: any_card. Shortlist: 298 names.

Commander: Adeline, Resplendent Cathar.

Grade: bad, score 0.03, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $83.03 to buy, $83.03 the whole deck.

**Summary:** Adeline leads a white go-wide token deck that develops a broad creature board, supports it with token-focused permanents, and turns repeated attacks into pressure. It has several dedicated ways to close a stalled game, backed by finisher-level threats, while draw, ramp, removal, protection, and board resets keep the plan moving. It gives up premium individual-card power for a cohesive, low-cost board-building strategy that is strongest when it can keep creatures in play and attack consistently.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The deck makes less mana on turn four than the norm, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core is a coherent mono-white go-wide shell (Adeline, anthem/token permanents like Divine Visitation, Intangible Virtue, Worthy Knight, plus convoke-friendly wipes such as Hour of Reckoning), but the ramp package of marginal treasure/artifact trinkets (Coin of Mastery, Collector's Vault, Currency Converter, Goldvein Pick, Bucknard's Everfull Purse, Karn) does little for the token plan and reads as filler pulling away from the board-building theme.
- PLAN theme_fit=yes: It is exactly the requested mono-white Adeline token Commander deck, built from cheap commons and mid-tier synergy pieces consistent with a sub-$100 bracket-2 build rather than fast mana or tutors.
- PLAN useful_as_built=yes: 37 mono-colored lands with no color risk, a reasonable curve of two- to five-drops, removal, wipes, card draw and multiple combat-based and alternate paths to victory mean the deck can be shuffled up and played as printed.
- PLAN summary_honest=partly: The go-wide description matches the list and the weaknesses are disclosed, but calling Luck Bobblehead and Sword of Body and Mind dedicated ways to close a stalled game overstates things (Halo Fountain is the only real alternate wincon), and 'ramp' oversells a pile of low-impact treasure artifacts.
- [INFO] `curve_summary`: average mana value 3.53 over 62 nonland cards

<details><summary>The deck list</summary>

- 31 Plains | land | Provides the reliable white mana base the deck needs.
- 1 Command Tower | land | A low-cost colored land for the mana base.
- 1 Exotic Orchard | land | Adds flexible colored mana without adding to the purchase cost.
- 1 Path of Ancestry | land | Supports the creature-focused mana base.
- 1 Secluded Courtyard | land | Supplies colored mana for the deck's many creatures.
- 1 Unclaimed Territory | land | Supplies colored mana for the creature suite.
- 1 Windbrisk Heights | land | A utility land that remains within the tapped-land allowance.
- 1 Bucknard's Everfull Purse | ramp | Helps advance mana development.
- 1 Coin of Mastery | ramp | Provides additional mana development.
- 1 Collector's Vault | ramp | A cheap artifact ramp piece.
- 1 Currency Converter | ramp | A low-cost artifact ramp option.
- 1 Druidic Satchel | ramp | Adds mana development while keeping costs low.
- 1 Goldvein Pick | ramp | A low-cost ramp artifact for the token plan.
- 1 Idol of False Gods | ramp | Provides inexpensive additional ramp.
- 1 Karn, Living Legacy | ramp | A ramp source that also adds a durable permanent.
- 1 Keeper of the Accord | ramp | Supports the deck's mana development.
- 1 Legion's Landing // Adanto, the First Fort | ramp | A low-cost ramp card aligned with the token strategy.
- 1 Angelic Sell-Sword | draw | Provides card access in the creature suite.
- 1 Bygone Bishop | draw | A low-cost draw option for a creature-heavy deck.
- 1 Court of Grace | draw | Provides ongoing card access for the token plan.
- 1 Dawn of Hope | draw | A flexible draw engine for a go-wide strategy.
- 1 Faramir, Field Commander | draw | Adds card access while fitting the creature plan.
- 1 Glimmer Seeker | draw | A very inexpensive source of card access.
- 1 Idol of Oblivion | draw | A draw artifact suited to a token-focused deck.
- 1 Sarah Jane Smith | draw | Provides additional card access at a low cost.
- 1 Search the Premises | draw | An ongoing draw option for longer games.
- 1 Staff of the Storyteller | draw | A cheap artifact draw engine.
- 1 Wedding Announcement // Wedding Festivity | draw | Combines card access with the deck's board-focused plan.
- 1 Rootborn Defenses | interaction | Protects the token board from opposing pressure.
- 1 Lena, Selfless Champion | interaction | Adds protective interaction on a creature body.
- 1 Mage's Attendant | interaction | Provides interactive support while contributing to the board.
- 1 Pegasus Guardian // Rescue the Foal | interaction | A flexible interaction card for the token strategy.
- 1 Spirit Bonds | interaction | Offers interaction that fits the creature-focused plan.
- 1 Squad Commander | interaction | Adds interactive utility to the creature suite.
- 1 Generous Gift | removal | Flexible removal for problematic opposing permanents.
- 1 Hanged Executioner | removal | Creature-based removal that fits the board plan.
- 1 Kellan's Lightblades | removal | A highly efficient budget removal option.
- 1 Righteous Confluence | removal | Flexible removal for a range of situations.
- 1 Skyclave Apparition | removal | Efficient creature-based removal.
- 1 Stroke of Midnight | removal | Versatile instant-speed removal.
- 1 Trostani's Judgment | removal | A low-cost removal spell for difficult threats.
- 1 Hour of Reckoning | wipe | A board reset that suits a token-oriented deck.
- 1 Martial Coup | wipe | Combines a reset option with the deck's token plan.
- 1 Elspeth, Sun's Champion | wipe | A repeatable threat that also provides a board wipe.
- 1 Divine Visitation | synergy | A core token payoff for making the board more threatening.
- 1 Felidar Retreat | synergy | A token-focused synergy piece that rewards development.
- 1 Horn of Gondor | synergy | Adds another dedicated token-synergy artifact.
- 1 Intangible Virtue | synergy | A low-cost payoff for building a wide token board.
- 1 Mavren Fein, Dusk Apostle | synergy | A creature-based token synergy piece.
- 1 Retreat to Emeria | synergy | Supports token production through the deck's land development.
- 1 Rosie Cotton of South Lane | synergy | Rewards the deck for developing a token board.
- 1 Siege Veteran | synergy | A creature payoff for a growing board.
- 1 Skrelv's Hive | synergy | A dedicated token-synergy permanent.
- 1 Worthy Knight | synergy | Adds another inexpensive token-synergy creature.
- 1 Archon of Sun's Grace | threat | A token-oriented threat that advances the board plan.
- 1 Attended Healer | threat | An inexpensive creature threat for the token strategy.
- 1 Basri's Lieutenant | threat | A board-building threat that fits the creature suite.
- 1 Defiler of Faith | threat | A resilient threat for the white creature plan.
- 1 Emeria Angel | threat | A token-oriented threat that rewards land development.
- 1 God-Eternal Oketra | threat | A durable top-end creature threat.
- 1 Gwaihir, Greatest of the Eagles | threat | An inexpensive finisher threat for closing games.
- 1 Hero of Bladehold | threat | An aggressive threat suited to attacking with a wide board.
- 1 Nahiri, the Lithomancer | threat | A finisher-level threat that adds pressure to the board.
- 1 Oketra the True | threat | A durable creature threat for the top of the curve.
- 1 Requiem Angel | threat | A finisher threat that supports the token theme.
- 1 Silverwing Squadron | threat | A wide-board threat that helps pressure opponents.
- 1 Halo Fountain | wincon | A dedicated token-board win condition.
- 1 Luck Bobblehead | wincon | Provides a distinct finisher-level win condition.
- 1 Sword of Body and Mind | wincon | An equipment-based win condition for the creature board.

</details>

### 13. owned first, and the commander is not owned

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 230 names.

Commander: Karlov of the Ghost Council.

Grade: bad, score 0.02, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $179.99 the whole deck.

**Summary:** This lifegain deck develops its mana, establishes Karlov alongside a focused group of lifegain synergy cards, and uses efficient removal to keep opposing boards under control. It wins by turning that incremental advantage into sustained creature pressure, then closing with evasive Angels and dedicated finishing cards. The deck gives up some raw speed for a resilient, board-centered plan with substantial protection and reset options.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves one color far better than another, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: There is a recognizable Orzhov lifegain-plus-Karlov shell with removal, protection equipment and some gain triggers, but the plan is diluted: only a handful of true lifegain payoffs, three board wipes that fight the deck's own creature-pressure/Karlov-counters plan, and a grab-bag of unrelated threats (Canyon Crawler, Foggy Swamp Hunters, Sneering Shadewriter) that do nothing for lifegain.
- PLAN theme_fit=partly: It is a Karlov Commander deck at roughly bracket-3 power, but the lifegain theme is only lightly supported and the 'from my library first' constraint cannot be verified while the list leans on many recent Universes Beyond/Final Fantasy cards that suggest generic pool-building rather than a curated collection.
- PLAN useful_as_built=partly: 38 lands, decent ramp and a low curve make it playable, but the mana is badly skewed (18 Plains to 10 Swamp) for a deck with double-black and BB-ish costs, and the win plan is diffuse creature beats backed by wipes that undercut it.
- PLAN summary_honest=partly: Claims of removal, protection and resets are accurate, but 'closing with evasive Angels' overstates a thin angel count and the summary never acknowledges the heavy Plains/Swamp imbalance or the scarcity of actual gain-triggered payoffs, though the appended model notes do mention the mana skew.
- [WARN] `not_owned`: Karlov of the Ghost Council: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.22 over 63 nonland cards

<details><summary>The deck list</summary>

- 18 Plains | land | Provides a dependable white mana base for the deck.
- 10 Swamp | land | Provides a dependable black mana base for the deck.
- 1 Command Tower | land | Provides flexible commander-color mana.
- 1 City of Brass | land | Provides flexible colored mana.
- 1 Exotic Orchard | land | Adds flexible colored mana.
- 1 Marsh Flats | land | Finds a needed basic land color.
- 1 Fabled Passage | land | Finds a needed basic land color.
- 1 Evolving Wilds | land | Finds a needed basic land color.
- 1 Spire of Industry | land | Supports the artifact-heavy mana package while fixing colors.
- 1 Plaza of Heroes | land | Supports the legendary commander and colored mana needs.
- 1 Sol Ring | ramp | Efficiently accelerates the deck's development.
- 1 Arcane Signet | ramp | Reliable colored acceleration for the commander deck.
- 1 Bender's Waterskin | ramp | Adds to the deck's mana development.
- 1 Blitzball | ramp | Provides additional mana acceleration.
- 1 Brass Infiniscope | ramp | Contributes to advancing the mana base.
- 1 Commander's Sphere | ramp | Fixes mana while supporting later resources.
- 1 Fellwar Stone | ramp | Low-cost mana acceleration.
- 1 Inherited Envelope | ramp | Adds another mana-development piece.
- 1 Relic of Legends | ramp | Provides flexible mana acceleration.
- 1 Thought Vessel | ramp | Supplies efficient artifact mana.
- 1 Buster Sword | draw | Provides a card-advantage equipment option.
- 1 Call of the Ring | draw | Adds a steady card-advantage source.
- 1 Circle of Power | draw | Provides card advantage for the deck.
- 1 Exemplar of Light | draw | Pairs card advantage with the creature plan.
- 1 Hero in Training | draw | Adds a creature-based draw option.
- 1 Idol of Oblivion | draw | Provides a compact card-advantage piece.
- 1 Inspiring Overseer | draw | Contributes card advantage on a lifegain-oriented body.
- 1 Lembas | draw | Provides a low-cost card-advantage artifact.
- 1 Mask of Memory | draw | Turns creature combat into card selection and advantage.
- 1 Night's Whisper | draw | Efficiently refills resources.
- 1 Skullclamp | draw | Provides efficient creature-based card advantage.
- 1 Banishing Light | removal | Answers a problematic permanent.
- 1 Bitter Triumph | removal | Provides flexible single-target removal.
- 1 Crib Swap | removal | Exiles a creature at instant speed.
- 1 Dismember | removal | Gives the deck cheap creature interaction.
- 1 Dispatch | removal | Provides efficient targeted removal.
- 1 Fatal Push | removal | Supplies low-cost creature removal.
- 1 Generous Gift | removal | Answers a broad range of permanents.
- 1 Infernal Grasp | removal | Offers unconditional creature removal.
- 1 Swords to Plowshares | removal | Provides premium creature removal.
- 1 Angel of Vitality | synergy | Supports the deck's lifegain-focused synergies.
- 1 Aerith Gainsborough | synergy | Contributes to the lifegain-focused core.
- 1 Compassionate Healer | synergy | Reinforces the deck's lifegain plan.
- 1 Light of Promise | synergy | Converts the deck's lifegain focus into a stronger board presence.
- 1 Prideful Feastling | synergy | Supports the deck's lifegain synergies.
- 1 Rosie Cotton of South Lane | synergy | Builds on the deck's lifegain-oriented creature plan.
- 1 Second Breakfast | synergy | Provides another focused lifegain synergy piece.
- 1 Angel of Invention | threat | Adds a substantial creature threat to pressure opponents.
- 1 Canyon Crawler | threat | Provides a dedicated board threat.
- 1 Dawnhand Eulogist | threat | Presents a threatening finisher-class creature.
- 1 Foggy Swamp Hunters | threat | Adds another creature threat for the battlefield.
- 1 Gollum, Silent Slinker // Meager Meal | threat | Supplies a threatening creature with added flexibility.
- 1 Minwu, White Mage | threat | Adds to the deck's creature pressure.
- 1 Rabaroo Troop | threat | Provides an additional board threat.
- 1 Reaping Willow | threat | Adds a durable creature threat.
- 1 Rooftop Percher | threat | Supplies another creature to apply pressure.
- 1 Shattered Angel | threat | Adds an evasive threat aligned with the deck's theme.
- 1 Sneering Shadewriter | threat | Presents a threatening finisher-class body.
- 1 Victory's Herald | threat | Provides a powerful top-end creature threat.
- 1 Frodo, Sauron's Bane | wincon | Serves as a dedicated finishing route.
- 1 Grave Venerations | wincon | Provides a finisher-class route to close games.
- 1 Lyra Dawnbringer | wincon | Offers a powerful evasive finishing threat.
- 1 Austere Command | wipe | Provides a flexible reset when opponents develop too far ahead.
- 1 Fumigate | wipe | Resets creature-heavy boards while fitting the lifegain plan.
- 1 Vanquish the Horde | wipe | Provides an efficient broad creature reset.
- 1 Bastion Protector | interaction | Helps keep the commander protected on the battlefield.
- 1 Boromir, Warden of the Tower | interaction | Adds proactive protection against opposing plays.
- 1 Champion's Helm | interaction | Protects the commander and other key legends.
- 1 Clever Concealment | interaction | Protects the board from opposing interaction.
- 1 Darksteel Plate | interaction | Helps preserve an important creature through removal.
- 1 Lightning Greaves | interaction | Protects key creatures while improving their deployment.
- 1 Swiftfoot Boots | interaction | Provides repeatable protection for the commander or a finisher.
- 1 Unbreakable Formation | interaction | Protects the creature board at a crucial moment.

</details>

### 14. the user delegates the commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 247 names.

Commander: Atarka, World Render.

Grade: baseline, score 0.06, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $1103.22 to buy, $1103.22 the whole deck.

**Summary:** This deck develops red-green mana, uses inexpensive support to keep cards flowing, and then fills the sky with Dragons. Atarka and the other large threats turn combat into the primary route to victory, while dedicated closing cards give the deck ways to end games once it has established a board. It carries plentiful targeted answers and protective effects, but its strongest turns rely on keeping expensive creatures in play and maintaining enough mana to deploy them.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. Many of the cards appear in no top list, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every card supports one plan — cheat dragons out early with cost reducers, orbs and ramp, then win through Atarka-boosted dragon combat backed by dragon-synergy payoffs like Dragon Tempest, Lathliss and Crucible of Fire — with only minor outliers such as Skullclamp in a deck of few small creatures.
- PLAN theme_fit=yes: It is a dragons Commander deck with a builder-chosen dragon commander (Atarka, World Render) in red-green, and its power level — dragon tribal beatdown with one game changer in Wheel of Fortune — sits comfortably at bracket 3.
- PLAN useful_as_built=yes: 36 lands with heavy dual/utility fixing, a real ramp and cost-reduction package for the expensive dragons, eleven draw pieces and ample removal means it can be picked up and played, with many dragon threats plus Hellkite Tyrant and Lathliss to close.
- PLAN summary_honest=yes: The summary accurately describes the red-green mana base, the cheap support and card flow, the combat-first win route and the reliance on expensive creatures; the only quibble is labeling Thundermaw Hellkite a sweeper, which does not overstate the deck's capabilities.
- [INFO] `curve_summary`: average mana value 3.02 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Cavern of Souls | land | A flexible Dragon-focused color source.
- 1 City of Brass | land | Provides either commander color immediately.
- 1 Command Tower | land | Reliable access to both deck colors.
- 1 Copperline Gorge | land | An early untapped red-green source.
- 1 Cragcrown Pathway // Timbercrown Pathway | land | Flexible untapped access to either color.
- 1 Exotic Orchard | land | Flexible multicolor fixing.
- 1 Fabled Passage | land | Finds the needed basic color.
- 1 Fire-Lit Thicket | land | Helps convert mana into the needed colors.
- 1 Forbidden Orchard | land | Supplies either deck color.
- 6 Forest | land | Stable green mana for early development.
- 1 Game Trail | land | A red-green source for the mana base.
- 1 Karplusan Forest | land | An untapped red-green source.
- 1 Mana Confluence | land | Provides either color without delay.
- 6 Mountain | land | Stable red mana for Dragons and interaction.
- 1 Plaza of Heroes | land | Supports the deck's legendary cards while fixing mana.
- 1 Prismatic Vista | land | Finds a basic land of the needed color.
- 1 Rockfall Vale | land | A red-green dual land.
- 1 Rootbound Crag | land | A red-green dual land.
- 1 Secluded Courtyard | land | Dragon-focused color fixing.
- 1 Spire Garden | land | An untapped red-green source in multiplayer.
- 1 Stomping Ground | land | An efficient red-green dual land.
- 1 Taiga | land | An efficient red-green dual land.
- 1 Tarnished Citadel | land | Provides either deck color when needed.
- 1 Thornspire Verge | land | A red-green source for the mana base.
- 1 Unclaimed Territory | land | Dragon-focused color fixing.
- 1 Wooded Foothills | land | Finds a red or green dual or basic land.
- 1 Mox Jasper | ramp | Low-cost Dragon-themed mana acceleration.
- 1 Carnelian Orb of Dragonkind | ramp | Fixes mana while accelerating into Dragons.
- 1 Dragon's Hoard | ramp | Dragon-themed mana development.
- 1 Jade Orb of Dragonkind | ramp | Low-cost green acceleration for the deck.
- 1 Orb of Dragonkind | ramp | Two-mana fixing and acceleration.
- 1 Reckless Barbarian | ramp | A cheap burst of mana toward larger plays.
- 1 Scaled Nurturer | ramp | Early mana development on a Dragon body.
- 1 Sarkhan, Dragon Ascendant | ramp | Develops mana while fitting the Dragon plan.
- 1 Sarkhan, Fireblood | ramp | Accelerates the deck's Dragon-heavy curve.
- 1 Embermouth Sentinel | ramp | Provides another low-cost mana piece.
- 1 Demand Answers | draw | Efficiently refills the hand.
- 1 Elemental Bond | draw | Turns the deck's large creatures into steady cards.
- 1 Faithless Looting | draw | Cheap card selection early or late.
- 1 Garruk's Uprising | draw | Draw support that suits large Dragon threats.
- 1 Guardian Project | draw | Sustained card advantage for a creature-heavy deck.
- 1 Idol of Oblivion | draw | A low-cost source of card advantage.
- 1 Sensei's Divining Top | draw | Improves draw quality throughout the game.
- 1 Skullclamp | draw | Efficient card advantage support.
- 1 Sylvan Library | draw | Reliable early card selection.
- 1 Thrill of Possibility | draw | Instant-speed filtering when mana is available.
- 1 Wheel of Fortune | draw | A powerful hand reset when resources run low.
- 1 Heroic Intervention | interaction | Protects the developed board from opposing disruption.
- 1 Lightning Greaves | interaction | Protects an important creature while supporting immediate pressure.
- 1 Mithril Coat | interaction | Keeps a key creature on the table.
- 1 Snakeskin Veil | interaction | A cheap protective response.
- 1 Swiftfoot Boots | interaction | Protection for the commander or a major threat.
- 1 Tamiyo's Safekeeping | interaction | Flexible protection for a permanent or player.
- 1 Veil of Summer | interaction | Efficient protection against targeted disruption.
- 1 Whispersilk Cloak | interaction | Protects a threat and helps it connect.
- 1 Dragon Tempest | removal | Dragon-focused damage that answers smaller targets.
- 1 Dragon's Fire | removal | Efficient creature removal.
- 1 Draconic Roar | removal | Low-cost Dragon-themed removal.
- 1 Invasion of Tarkir // Defiant Thundermaw | removal | A flexible removal spell that remains on-theme.
- 1 Molten Exhale | removal | Additional targeted removal.
- 1 Obsidian Charmaw | removal | A Dragon threat that also handles problem permanents.
- 1 Piercing Exhale | removal | Cheap instant-speed removal.
- 1 Spit Flame | removal | Repeatable Dragon-linked removal support.
- 1 Smoldering Egg // Ashmouth Dragon | removal | Early removal that can become a Dragon threat.
- 1 Crucible of Fire | synergy | Raises the pressure of the Dragon army.
- 1 Dragonkin Berserker | synergy | A low-cost Dragon payoff.
- 1 Dragonlord's Servant | synergy | Supports the deck's Dragon-heavy plan.
- 1 Dragonspeaker Shaman | synergy | Helps the deck deploy its Dragon cards.
- 1 Firespitter Whelp | synergy | An early Dragon-themed support card.
- 1 Sarkhan's Triumph | synergy | Finds the right Dragon for the current board.
- 1 The Dragon-Kami Reborn // Dragon-Kami's Egg | synergy | A thematic value piece for the Dragon strategy.
- 1 Backdraft Hellkite | threat | A substantial evasive Dragon threat.
- 1 Caldera Pyremaw | threat | Adds a durable Dragon body to the attack plan.
- 1 Dragonhawk, Fate's Tempest | threat | A major evasive Dragon threat.
- 1 Dream Pillager | threat | Applies meaningful combat pressure.
- 1 Hoarding Dragon | threat | A Dragon body that advances the board plan.
- 1 Manaform Hellkite | threat | An efficient flying Dragon attacker.
- 1 Mirrorwing Dragon | threat | A potent evasive threat that rewards protection.
- 1 Sagu Wildling // Roost Seek | threat | A flexible Dragon threat for the creature base.
- 1 Stormbreath Dragon | threat | A hasty evasive attacker.
- 1 Thunderbreak Regent | threat | A reliable mid-curve Dragon threat.
- 1 Thundermane Dragon | threat | Adds another efficient Dragon attacker.
- 1 Territorial Hellkite | threat | Provides early aerial pressure.
- 1 Hellkite Tyrant | wincon | A Dragon finisher with a dedicated victory route.
- 1 Lathliss, Dragon Queen | wincon | A Dragon finisher that scales with continued development.
- 1 Wrathful Red Dragon | wincon | A Dragon-based closing piece.
- 1 Breath Weapon | wipe | Clears away opposing smaller creatures.
- 1 Draconic Intervention | wipe | A scalable board-clearing option.
- 1 Thundermaw Hellkite | wipe | A Dragon threat that also pressures opposing boards.

</details>

### 15. delegated commander, owned first

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 230 names.

Commander: Astarion, the Decadent.

Grade: baseline, score 0.04, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $194.96 the whole deck.

**Summary:** Astarion leads a white-black lifegain deck that develops its mana, builds a creature board around Angels and other lifegain-focused pieces, and keeps cards flowing through its draw package. Targeted answers and protective equipment help preserve the board until Frodo, Grave Venerations, Lyra Dawnbringer, and other marked finishers can close the game. It gives up explosive combo speed for a steadier, board-centered plan that can reset crowded tables and rebuild through its card advantage.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. Many of the cards appear in no top list, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every nonland slot either gains life (Compassionate Healer, Rosie Cotton, Angel of Vitality, Kor Firewalker) or converts that life into counters, cards, or bodies (Light of Promise, Exemplar of Light, Prideful Feastling, Astarion himself), with removal, wipes, and equipment supporting a single go-wide-with-lifegain board plan.
- PLAN theme_fit=yes: The request asked the builder to choose a commander for a lifegain Commander deck at bracket 2, and Astarion with a white-black lifegain shell of creatures, equipment, and one-for-one removal with no fast combos or tutors sits right in that band.
- PLAN useful_as_built=yes: About 37 lands split sensibly between Plains and Swamps with Command Tower, Marsh Flats, City of Brass, and Exotic Orchard, roughly nine ramp pieces, a playable curve, and multiple angel/equipment finishers means it can be shuffled up and played as-is.
- PLAN summary_honest=yes: The claimed pieces are all present — angels and lifegain creatures, a real draw package (Skullclamp, Mask of Memory, Night's Whisper, Idol of Oblivion, Inspiring Overseer), targeted answers, protective equipment, and two sweepers for the promised table reset — and it honestly admits it has no combo speed.
- [WARN] `not_owned`: Astarion, the Decadent: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.03 over 62 nonland cards
- [INFO] `cards_trimmed`: the list was 1 card over, so the builder cut this card: Vanquish the Horde

<details><summary>The deck list</summary>

- 1 Angel of Invention | threat | A white threat that supports the creature plan.
- 1 Angel of Vitality | synergy | A lifegain-focused synergy piece.
- 1 Arcane Signet | ramp | Reliable colored mana acceleration.
- 1 Austere Command | wipe | A flexible board reset.
- 1 Aettir and Priwen | synergy | An equipment-based lifegain synergy piece.
- 1 Aerith Gainsborough | synergy | A central lifegain synergy creature.
- 1 Banishing Light | removal | Versatile permanent removal.
- 1 Bastion Protector | interaction | Protection for an important legendary permanent.
- 1 Bill the Pony | threat | A white creature threat for the board plan.
- 1 Bitter Triumph | removal | Efficient targeted removal.
- 1 Buster Sword | draw | Equipment that supplies card advantage.
- 1 Call of the Ring | draw | A repeatable card-advantage piece.
- 1 Chromatic Lantern | ramp | Mana fixing and acceleration.
- 1 Circle of Power | draw | A draw spell that helps keep resources flowing.
- 1 Clever Concealment | interaction | Protection against opposing disruption.
- 1 Command Tower | land | An untapped source for either deck color.
- 1 Compassionate Healer | synergy | A lifegain synergy creature.
- 1 Crib Swap | removal | Creature removal.
- 1 Dawnhand Eulogist | threat | A black creature threat.
- 1 Darksteel Plate | interaction | Durable protection for a key creature.
- 1 Excalibur II | synergy | An equipment-based synergy piece.
- 1 Exemplar of Light | draw | An Angel that provides card advantage.
- 1 Exotic Orchard | land | Flexible color fixing.
- 1 Fellwar Stone | ramp | Low-cost mana acceleration.
- 1 Foggy Swamp Hunters | threat | A creature threat for the board plan.
- 1 Frodo, Sauron's Bane | wincon | A marked finisher for closing games.
- 1 Fumigate | wipe | A board reset for crowded games.
- 1 Generous Gift | removal | Broad permanent removal.
- 1 Get Lost | removal | Efficient targeted removal.
- 1 Giada, Font of Hope | ramp | Mana acceleration that also supports Angels.
- 1 Gollum, Silent Slinker // Meager Meal | threat | A black creature threat.
- 1 Grave Venerations | wincon | A marked finisher for the lifegain plan.
- 1 Hero in Training | draw | A creature that provides card advantage.
- 1 Idol of Oblivion | draw | Low-cost card advantage.
- 1 Inspiring Overseer | draw | An Angel that supplies card advantage.
- 1 Infernal Grasp | removal | Direct creature removal.
- 1 Kor Firewalker | synergy | A lifegain synergy creature.
- 1 Lembas | draw | A compact card-advantage artifact.
- 1 Light of Promise | synergy | A dedicated lifegain synergy enchantment.
- 1 Lightning Greaves | interaction | Low-cost protection for key creatures.
- 1 Lotho, Corrupt Shirriff | ramp | A mana-producing creature.
- 1 Lyra Dawnbringer | wincon | An Angel finisher for ending games.
- 1 Marsh Flats | land | A flexible land source for the two-color mana base.
- 1 Mask of Memory | draw | Equipment-based card advantage.
- 1 Minwu, White Mage | threat | A white creature threat.
- 1 Night Nurse, Healer of Heroes | synergy | A lifegain synergy creature.
- 1 Night's Whisper | draw | Efficient card advantage.
- 1 Plains | land | A basic white mana source.
- 17 Plains | land | Additional basic white mana sources.
- 1 Prideful Feastling | synergy | A creature dedicated to lifegain synergy.
- 1 Rabaroo Troop | threat | A creature threat for building pressure.
- 1 Reaping Willow | threat | A creature threat for the board plan.
- 1 Rooftop Percher | threat | A creature threat that adds board presence.
- 1 Rosie Cotton of South Lane | synergy | A lifegain synergy creature.
- 1 Second Breakfast | synergy | A lifegain synergy instant.
- 1 Shattered Angel | threat | An Angel threat for the creature plan.
- 1 Skullclamp | draw | A low-cost card-advantage equipment.
- 1 Sneering Shadewriter | threat | A black finisher threat.
- 1 Sol Ring | ramp | Fast, efficient mana acceleration.
- 1 Swords to Plowshares | removal | Efficient creature removal.
- 1 Swiftfoot Boots | interaction | Protection for key creatures.
- 1 Swamp | land | A basic black mana source.
- 14 Swamp | land | Additional basic black mana sources.
- 1 Sword of the Animist | ramp | Equipment that advances the mana plan.
- 1 Thought Vessel | ramp | A mana rock for steady acceleration.
- 1 Unbreakable Formation | interaction | Protection against opposing sweepers and removal.
- 1 Victory's Herald | threat | An Angel threat that pressures opponents.
- 1 Wayfarer's Bauble | ramp | Early mana acceleration.
- 1 White Mage's Staff | synergy | An equipment-based lifegain synergy piece.
- 1 City of Brass | land | An untapped source for either deck color.

</details>

### 16. a tight budget, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 231 names.

Commander: Denethor, Ruling Steward.

Grade: baseline, score 0.29, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $15.70 to buy, $52.57 the whole deck.

**Summary:** This is a black-white aristocrats deck that develops a creature board, uses its sacrifice-focused synergy pieces to turn that board into value, and closes with Blood Artist, Falkenrath Noble, or Zulaport Cutthroat. It has a solid supply of cheap mana, card advantage, targeted answers, and a few board resets to keep playing through opposing pressure. The tradeoff is that its strongest finishes depend on maintaining creatures and payoff pieces rather than presenting a single fast, independently powerful threat.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every slot supports one aristocrats line — roughly a dozen sacrifice outlets, a body-generating creature base, death-trigger drains, and sac-fueled draw like Deadly Dispute, Disciple of Bolas and Smothering Abomination — with only the three board wipes sitting slightly at odds with the go-wide plan, and even those are modal or recursive.
- PLAN theme_fit=partly: It is a genuine black-white aristocrats Commander deck at a fair bracket-2 power level, but the $25 ceiling looks broken by the stack of pricier inclusions (City of Brass, Plaza of Heroes, Blood Artist, Sol Ring, Swords to Plowshares, Chromatic Lantern, Austere Command), which is the one hard limit the player named.
- PLAN useful_as_built=yes: 38 lands plus eight cheap rocks support a low curve, there are plenty of bodies and outlets to get the engine going, and three separate drain payoffs plus incidental combat damage give clear routes to win.
- PLAN summary_honest=yes: The summary accurately describes the drain finishers, the sac-based value engine, the wipes, and honestly flags the deck's dependence on keeping a board rather than a single fast threat, without claiming tutors or combos the list lacks.
- [WARN] `not_owned`: Aron, Benalia's Ruin: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ayli, Eternal Pilgrim: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bartolomé del Presidio: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Umbral Collar Zealot: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Woe Strider: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Baron Bertram Graywater: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Disciple of Bolas: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Erebos, Bleak-Hearted: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Lord Skitter's Butcher: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Old Flitterfang: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ruthless Technomancer: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Shadowheart, Dark Justiciar: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Sivriss, Nightmare Speaker: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Skullport Merchant: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Smothering Abomination: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vampire Gourmand: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Yahenni, Undying Partisan: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Blood Artist: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Falkenrath Noble: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Zulaport Cutthroat: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 2.95 over 62 nonland cards
- [INFO] `mana_pass`: the builder moved 1 card of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 16 Plains | land | Basic white source for reliable early white mana.
- 14 Swamp | land | Basic black source for reliable early black mana.
- 1 Command Tower | land | Flexible dual-color land source.
- 1 Exotic Orchard | land | Flexible nonbasic color source.
- 1 City of Brass | land | Untapped access to both deck colors.
- 1 Grand Coliseum | land | Flexible nonbasic color source.
- 1 Plaza of Heroes | land | Color source that supports the legendary commander plan.
- 1 Spire of Industry | land | Nonbasic color source alongside the artifact package.
- 1 Secluded Courtyard | land | Creature-focused color fixing for the board plan.
- 1 Sol Ring | ramp | Cheap acceleration for developing the board quickly.
- 1 Arcane Signet | ramp | Reliable two-color mana acceleration.
- 1 Fellwar Stone | ramp | Low-cost mana fixing and acceleration.
- 1 Chromatic Lantern | ramp | Fixes both colors while adding mana.
- 1 Commander's Sphere | ramp | Color fixing that remains useful later.
- 1 Thought Vessel | ramp | Efficient artifact mana acceleration.
- 1 Relic of Legends | ramp | Mana acceleration that works with the legendary cards.
- 1 Ring of the Lucii | ramp | Additional artifact acceleration.
- 1 White Auracite | ramp | Low-cost ramp for the white-heavy requirements.
- 1 Deadly Dispute | ramp | Sacrifice-friendly acceleration.
- 1 Cirith Ungol Patrol | draw | Creature-based card advantage for the black deck core.
- 1 Exemplar of Light | draw | Creature-based card advantage.
- 1 Inspiring Overseer | draw | Early creature that replaces itself.
- 1 Mask of Memory | draw | Repeatable card advantage from combat.
- 1 Lembas | draw | Low-cost card advantage artifact.
- 1 Idol of Oblivion | draw | Artifact card advantage for the board plan.
- 1 Nasty End | draw | Black card advantage that fits a creature deck.
- 1 Painful Truths | draw | Efficient black card advantage.
- 1 Stone of Erech | draw | Low-cost artifact card advantage.
- 1 Tome of Legends | draw | Commander-friendly repeatable card advantage.
- 1 Wall of Omens | draw | Early defensive creature that draws a card.
- 1 Bastion Protector | interaction | Protects the commander from opposing interaction.
- 1 Frontline Medic | interaction | Creature-based protection for the board.
- 1 Gift of Immortality | interaction | Keeps an important creature involved after removal.
- 1 Swiftfoot Boots | interaction | Efficient protection for the commander or a payoff creature.
- 1 Take Up the Shield | interaction | Low-cost protection during key turns.
- 1 Together Forever | interaction | Protection that supports a creature-focused game plan.
- 1 Claim the Precious | removal | Efficient black spot removal.
- 1 Crib Swap | removal | Flexible creature removal.
- 1 Fatal Push | removal | Cheap instant-speed removal.
- 1 Fiend Hunter | removal | Creature-based removal that contributes to the board.
- 1 Infernal Grasp | removal | Reliable black creature removal.
- 1 Stroke of Midnight | removal | Flexible white permanent removal.
- 1 Swords to Plowshares | removal | Efficient white creature removal.
- 1 Austere Command | wipe | Flexible reset button against developed boards.
- 1 Dusk // Dawn | wipe | Board reset that remains useful in a creature deck.
- 1 Fumigate | wipe | Creature-board reset for difficult positions.
- 1 Aron, Benalia's Ruin | synergy | Aristocrats-oriented legendary creature for the sacrifice core.
- 1 Ayli, Eternal Pilgrim | synergy | Low-cost aristocrats creature that supports the central plan.
- 1 Bartolomé del Presidio | synergy | Sacrifice-focused creature for the deck's core plan.
- 1 Gollum the Abandoned | synergy | Black synergy creature that supports the sacrifice shell.
- 1 Gollum, Patient Plotter | synergy | Creature synergy piece for the recursive aristocrats plan.
- 1 Gríma Wormtongue | synergy | Black synergy creature for the deck's board plan.
- 1 Heirloom Auntie | synergy | Low-cost black synergy creature.
- 1 Phantom Train | synergy | Artifact synergy piece that broadens the sacrifice board.
- 1 Umbral Collar Zealot | synergy | Low-cost black synergy creature for the aristocrats core.
- 1 Woe Strider | synergy | Sacrifice-oriented creature that anchors the deck's plan.
- 1 Baron Bertram Graywater | threat | Black creature threat that builds the board.
- 1 Disciple of Bolas | threat | Creature threat that rewards a developed board.
- 1 Erebos, Bleak-Hearted | threat | Durable black threat for longer games.
- 1 Lord Skitter's Butcher | threat | Black creature threat for the sacrifice-focused board.
- 1 Old Flitterfang | threat | Black creature threat that contributes to board presence.
- 1 Ruthless Technomancer | threat | Substantial black creature threat for midgame pressure.
- 1 Shadowheart, Dark Justiciar | threat | Black creature threat that supports a creature-heavy board.
- 1 Sivriss, Nightmare Speaker | threat | Low-cost black creature threat.
- 1 Skullport Merchant | threat | Black creature threat that fits the value-oriented creature plan.
- 1 Smothering Abomination | threat | Flying black threat for pressuring opponents.
- 1 Vampire Gourmand | threat | Low-cost black creature threat.
- 1 Yahenni, Undying Partisan | threat | Black aristocrats threat for advancing the board plan.
- 1 Blood Artist | wincon | Primary aristocrats finisher.
- 1 Falkenrath Noble | wincon | Creature-based aristocrats finisher.
- 1 Zulaport Cutthroat | wincon | Low-cost aristocrats finisher.

</details>

### 17. upgrade a precon, owned first

Format: Commander. Theme: goblins. Pool: owned_first. Shortlist: 198 names.

Commander: Zada, Hedron Grinder.

Grade: typical, score 0.53, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $84.65 to buy, $263.00 the whole deck.

**Summary:** This keeps the Goblin Storm core intact: develop a board of Goblins, use compact draw and mana spells to build momentum, and convert a crowded board into pressure through tribal payoffs and storm-oriented turns. Great Train Heist, Assault on Osgiliath, and Earthquake provide decisive closing routes, while the added Goblin Bombardment, Broadside Bombardiers, and Lightning Bolt improve the deck’s ability to answer resistance. The tradeoff is a focused, proactive plan rather than the flexibility of a broader control shell.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The bulk of the list pulls one way — Goblin bodies, cheap one-target cantrips/pumps for Zada, and storm/aristocrat payoffs like Grapeshot, Impact Tremors, Boggart Shenanigans, Past in Flames — but Blasphemous Act and Earthquake actively sweep the deck's own army of x/1 tokens, which fights the go-wide core.
- PLAN theme_fit=yes: It is a mono-red Zada Goblin-storm list at a bracket-3 level: precon core kept, upgrades are Sol Ring/Skullclamp/Goblin Bombardment/Bolt tier rather than fast mana and tutors.
- PLAN useful_as_built=yes: 38 lands in mono-red with a low curve, several ramp pieces, real draw (Skullclamp, War Room, Idol, Faithless Looting), and multiple redundant kill routes means it plays out fine as written.
- PLAN summary_honest=partly: Most claims check out (Zada cantrips, tribal payoffs, storm turns, the named removal additions), but calling Earthquake a decisive closing route is misleading in a deck whose win comes from a board of 1/1 Goblins that Earthquake kills first.
- [INFO] `curve_summary`: average mana value 2.56 over 61 nonland cards
- [WARN] `profile_off_band`: the interaction count is 1, and bracket 3 wants 4 to 12
- [INFO] `bracket_cut`: to hold bracket 3, the builder cut 1 card: Storm-Kiln Artist (the combo Storm-Kiln Artist + Haze of Rage), and added 1 basic land
- [INFO] `basics_added`: the list was 1 card short, so the builder added 1 basic land

<details><summary>The deck list</summary>

- 1 Ancestors' Aid | synergy | A retained precon spell that supports the deck's spell-focused plan.
- 1 Arena of Glory | land | A retained precon land that helps maintain the established mana base.
- 1 Assault on Osgiliath | wincon | An added finisher that gives the deck another decisive way to close a game.
- 1 Battle Hymn | ramp | A retained precon burst-mana spell for explosive Goblin turns.
- 1 Blasphemous Act | wipe | A retained precon board reset for crowded tables.
- 1 Boggart Shenanigans | removal | A retained precon Goblin payoff that contributes to the tribe's pressure.
- 1 Broadside Bombardiers | removal | An added Goblin that improves the deck's removal options while staying on theme.
- 1 Castle Embereth | land | A retained precon land that preserves the original mana base.
- 1 Chaos Warp | removal | A retained precon answer to troublesome permanents.
- 1 Conspicuous Snoop | synergy | A retained precon Goblin synergy piece.
- 1 Crimson Wisps | draw | A retained precon cantrip that keeps spells flowing.
- 1 Daring Discovery | other | A retained precon card that remains part of the deck's spell package.
- 1 Den of the Bugbear | land | A retained precon land that preserves the original mana base.
- 1 Dragon Fodder | synergy | A retained precon token-maker that supports the Goblin plan.
- 1 Dwarven Mine | land | A retained precon Mountain land that supports the Goblin theme.
- 1 Earthquake | wincon | An added finisher that can also clear smaller opposing boards.
- 1 Empty the Warrens | synergy | A retained precon storm-themed token payoff.
- 1 Expedite | draw | A retained precon cantrip for efficient spell velocity.
- 1 Fable of the Mirror-Breaker // Reflection of Kiki-Jiki | ramp | An added value engine that supplies ramp within the deck's creature-focused theme.
- 1 Faithless Looting | draw | A retained precon draw spell that helps find the right action.
- 1 Fists of Flame | draw | A retained precon cantrip that supports a spell-heavy turn.
- 1 Forgotten Cave | land | A retained precon land that keeps the mana base familiar.
- 1 Fountainport | land | A retained precon land that preserves the original mana base.
- 1 Frontline Heroism | synergy | A retained precon card that supports the combat-oriented plan.
- 1 Gempalm Incinerator | removal | A retained precon Goblin removal tool.
- 1 General Kreat, the Boltbringer | synergy | A retained precon legendary Goblin synergy piece.
- 1 Goblin Bombardment | removal | An added Goblin-themed removal outlet for the deck's creature resources.
- 1 Goblin Bushwhacker | synergy | A retained precon Goblin payoff for aggressive turns.
- 1 Goblin Chieftain | synergy | A retained precon tribal Goblin payoff.
- 1 Goblin Dark-Dwellers | threat | A retained precon Goblin threat.
- 1 Goblin Lackey | synergy | A retained precon Goblin synergy piece.
- 1 Goblin Matron | synergy | A retained precon Goblin synergy piece that helps the tribe work together.
- 1 Goblin Negotiation | removal | A retained precon Goblin removal spell.
- 1 Goblin Trashmaster | removal | A retained precon Goblin answer to opposing permanents.
- 1 Goblin Warchief | synergy | A retained precon tribal Goblin payoff.
- 1 Grapeshot | removal | A retained precon removal spell that fits the storm theme.
- 1 Great Train Heist | wincon | A retained precon finisher for closing through a strong combat turn.
- 1 Grenzo, Havoc Raiser | synergy | A retained precon legendary Goblin synergy piece.
- 1 Haze of Rage | synergy | A retained precon spell that supports a large finishing combat.
- 1 Hidden Volcano | land | A retained precon land that maintains the deck's established mana base.
- 1 Howlsquad Heavy | ramp | A retained precon Goblin that contributes to mana development.
- 1 Idol of Oblivion | draw | A retained precon draw engine for the token-oriented plan.
- 1 Impact Tremors | wincon | A retained precon payoff for repeatedly producing Goblins.
- 1 Kher Keep | land | A retained precon land that remains part of the original mana base.
- 1 Krenko's Command | synergy | A retained precon token-maker that supports Goblin synergy.
- 1 Krenko, Mob Boss | threat | A retained precon Goblin threat that strengthens the creature plan.
- 1 Lightning Bolt | removal | An added efficient removal spell for early threats and utility creatures.
- 27 Mountain | land | Basic lands keep the mana base stable and provide the deck's primary color source.
- 1 Mogg War Marshal | synergy | A retained precon Goblin synergy card that supports the token plan.
- 1 Pashalik Mons | removal | A retained precon legendary Goblin removal payoff.
- 1 Past in Flames | synergy | A retained precon spell-focused payoff for explosive late turns.
- 1 Quest for the Goblin Lord | synergy | A retained precon tribal Goblin payoff.
- 1 Reliquary Tower | land | A retained precon land that preserves the existing mana base.
- 1 Roaming Throne | synergy | A retained precon artifact creature that supports the tribal theme.
- 1 Ruby Medallion | ramp | A retained precon cost-reduction piece for the spell-heavy plan.
- 1 Sazacap's Brew | draw | A retained precon draw spell for continued card flow.
- 1 Searslicer Goblin | synergy | A retained precon Goblin synergy piece.
- 1 Seething Song | ramp | A retained precon burst-mana spell for storm turns.
- 1 Shinka, the Bloodsoaked Keep | land | A retained precon land that maintains the established mana base.
- 1 Siege-Gang Commander | removal | A retained precon Goblin removal payoff.
- 1 Siege-Gang Lieutenant | removal | A retained precon Goblin removal payoff.
- 1 Skirk Prospector | ramp | A retained precon Goblin mana source for explosive turns.
- 1 Skullclamp | draw | A retained precon draw engine that works with the deck's small creatures.
- 1 Sol Ring | ramp | A retained precon mana accelerator.
- 1 Spreading Insurrection | wincon | A retained precon high-impact spell that supports the aggressive plan.
- 1 Swiftfoot Boots | interaction | A retained precon protection piece for important creatures.
- 1 Throne of Eldraine | ramp | A retained precon mana artifact.
- 1 Vandalblast | wipe | A retained precon artifact-focused board reset.
- 1 War Room | land | A retained precon land that remains in the existing mana base.
- 1 Warren Torchmaster | synergy | A retained precon Goblin synergy piece.
- 1 Wild Ride | synergy | A retained precon spell that supports the deck's aggressive turns.
- 1 Witch's Mark | draw | A retained precon draw spell that helps maintain card flow.
- 1 Arcane Signet | ramp | An added reliable mana rock that improves early development.

</details>

### 18. upgrade a precon, any card

Format: Commander. Theme: turtles. Pool: any_card. Shortlist: 311 names.

Commander: Heroes in a Half Shell.

Grade: baseline, score 0.32, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $168.30 to buy, $168.30 the whole deck.

**Summary:** This stays a Turtle Power deck built around its Turtle and Mutant Ninja Turtle core, with a few focused additions for early mana, card flow, targeted answers, and additional Turtle-themed finishing pressure. It aims to establish its themed board and close through its marked finishers rather than shifting into a separate combo plan. The tradeoff is that much of the original precon cast remains in place, preserving its character over a broad reconstruction.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. Many lands enter tapped, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: The list is one coherent Turtle/TMNT tribal-and-ninjutsu board deck with mana, draw, and removal support, and no competing sub-plan is smuggled in.
- PLAN theme_fit=yes: It is the requested Turtle Power precon kept intact with modest bracket-3-appropriate upgrades (Sol Ring, Arcane Signet, Rhystic Study, Swords to Plowshares, Assassin's Trophy) rather than a rebuild into another archetype.
- PLAN useful_as_built=yes: Five-color mana is supported by a Command Tower, fixing lands, Chromatic Lantern, Birds, and Path of Ancestry with a reasonable land count, and the deck can close through its Turtle threats and finishers, though the many taplands make it slow.
- PLAN summary_honest=yes: The summary accurately describes a light upgrade that preserves the precon's core, names the added ramp, draw, and answers that are actually present, and openly admits most of the original cast remains.
- [INFO] `curve_summary`: average mana value 3.29 over 59 nonland cards
- [WARN] `profile_off_band`: the land count is 40, and bracket 3 wants 34 to 38
- [WARN] `profile_off_band`: the ramp count is 6, and bracket 3 wants 8 to 13
- [WARN] `profile_off_band`: the draw count is 3, and bracket 3 wants 8 to 14
- [WARN] `profile_off_band`: the removal count is 3, and bracket 3 wants 6 to 12
- [WARN] `profile_off_band`: the wipe count is 1, and bracket 3 wants 2 to 5
- [WARN] `profile_off_band`: the interaction count is 0, and bracket 3 wants 4 to 12
- [INFO] `cards_trimmed`: the list was 1 card over, so the builder cut this card: Blasphemous Act

<details><summary>The deck list</summary>

- 1 Acidic Slime | other | Kept from the Turtle Power precon.
- 1 April O'Neil, Live on the Scene | draw | Kept as a listed draw card.
- 1 Arcade Cabinet | other | Kept from the Turtle Power precon.
- 1 Arcane Signet | ramp | Kept as a listed ramp card.
- 1 Ash Barrens | land | Kept as part of the precon mana base.
- 1 Assassin's Trophy | removal | Kept as a listed removal card.
- 1 Baxter, Fly in the Ointment | other | Kept from the Turtle Power precon.
- 1 Bebop, Skull & Crossbones | other | Kept from the Turtle Power precon.
- 1 Big Apple, 3 a.m. | land | Kept as part of the precon mana base.
- 1 Big Mother Mouser | other | Kept from the Turtle Power precon.
- 1 Birds of Paradise | ramp | Added as a listed ramp card to improve early development.
- 1 Cinder Glade | land | Kept as part of the precon mana base.
- 1 City of Brass | land | Kept as part of the precon mana base.
- 1 Chromatic Lantern | ramp | Kept to support the precon's mana.
- 1 Command Tower | land | Kept as part of the precon mana base.
- 1 Continue? | other | Kept from the Turtle Power precon.
- 1 Cultivate | ramp | Kept to support the precon's mana.
- 1 Dimension X Pizzasaur | wincon | Kept as a listed precon finisher.
- 1 Donatello, the Brains | synergy | Kept as a listed Turtle synergy card.
- 1 Double Jump // Flying Kick | other | Kept from the Turtle Power precon.
- 1 Dragonskull Summit | land | Kept as part of the precon mana base.
- 1 Electric Seaweed | other | Kept from the Turtle Power precon.
- 1 Endless Foot Assault | other | Kept from the Turtle Power precon.
- 1 Escape Tunnel | land | Kept as part of the precon mana base.
- 1 Everything Pizza | wincon | Kept as a listed precon finisher.
- 1 Evolving Wilds | land | Kept as part of the precon mana base.
- 1 Exotic Orchard | land | Kept as part of the precon mana base.
- 1 Exploding Barrel | other | Kept from the Turtle Power precon.
- 1 Fast Forward | other | Kept from the Turtle Power precon.
- 1 Foot Chopper | other | Kept from the Turtle Power precon.
- 4 Forest | land | Kept as a basic land for the mana base.
- 1 Game Over | other | Kept from the Turtle Power precon.
- 1 Grand Coliseum | land | Kept as part of the precon mana base.
- 1 Harmonize | draw | Kept from the Turtle Power precon.
- 1 Here Comes a New Hero! | other | Kept from the Turtle Power precon.
- 1 Hidden Hideout | land | Kept as part of the precon mana base.
- 1 High Score | other | Kept from the Turtle Power precon.
- 3 Island | land | Kept as a basic land for the mana base.
- 1 Irma, Part-Time Mutant | other | Kept from the Turtle Power precon.
- 1 Krang, the All-Powerful | other | Kept from the Turtle Power precon.
- 1 Leatherhead, Iron Gator | other | Kept from the Turtle Power precon.
- 1 Leonardo, the Balance | threat | Kept as a listed Turtle threat.
- 1 Lessons from Life | other | Kept from the Turtle Power precon.
- 1 Level Up | other | Kept from the Turtle Power precon.
- 1 Lita, Little Orphan Amphibian | synergy | Kept as a listed Turtle synergy card.
- 1 Michelangelo, the Heart | synergy | Kept as a listed Turtle synergy card.
- 1 Mole Module | other | Kept from the Turtle Power precon.
- 1 Mona Lisa, Science Geek | other | Kept from the Turtle Power precon.
- 3 Mountain | land | Kept as a basic land for the mana base.
- 1 Ninja Pizza | other | Kept from the Turtle Power precon.
- 1 Path of Ancestry | land | Kept as part of the precon mana base.
- 3 Plains | land | Kept as a basic land for the mana base.
- 1 Rain-Slicked Copse | land | Kept as part of the precon mana base.
- 1 Raphael, the Muscle | wincon | Kept as a listed precon finisher.
- 1 Rat King, Pale Piper | other | Kept from the Turtle Power precon.
- 1 Ray Fillet, Wave Warrior | other | Kept from the Turtle Power precon.
- 1 Roadkill Rodney | other | Kept from the Turtle Power precon.
- 1 Rocksteady, Mutant Marauder | other | Kept from the Turtle Power precon.
- 1 Rootbound Crag | land | Kept as part of the precon mana base.
- 1 Raph & Leo, Sibling Rivals | wincon | Added as a listed Turtle finisher.
- 1 Rhystic Study | draw | Added as a listed draw card.
- 1 Shellshock | other | Kept from the Turtle Power precon.
- 1 Shredder, Shadow Master | other | Kept from the Turtle Power precon.
- 1 Smoldering Marsh | land | Kept as part of the precon mana base.
- 1 Sodden Verdure | land | Kept as part of the precon mana base.
- 1 Sol Ring | ramp | Kept as a listed ramp card.
- 1 Special Move | other | Kept from the Turtle Power precon.
- 1 Steelbane Hydra | removal | Kept from the Turtle Power precon.
- 1 Sunken Hollow | land | Kept as part of the precon mana base.
- 1 Super Combo | other | Kept from the Turtle Power precon.
- 3 Swamp | land | Kept as a basic land for the mana base.
- 1 Swift Demise | other | Kept from the Turtle Power precon.
- 1 Swords to Plowshares | removal | Added as a listed removal card.
- 1 Tempestra, Dame of Games | other | Kept from the Turtle Power precon.
- 1 Thriving Grove | land | Kept as part of the precon mana base.
- 1 Thriving Isle | land | Kept as part of the precon mana base.
- 1 Thriving Moor | land | Kept as part of the precon mana base.
- 1 Together Forever | other | Kept from the Turtle Power precon.
- 1 Tokka & Rahzar, Unsupervised | ramp | Kept as a listed ramp card.
- 1 Turtle Lair | land | Kept as part of the precon mana base.
- 1 Turtle Power! | synergy | Added as a listed Turtle synergy card.
- 1 Undergrowth Stadium | land | Kept as part of the precon mana base.
- 1 Vanquish the Horde | wipe | Kept from the Turtle Power precon.
- 1 Vernal Fen | land | Kept as part of the precon mana base.
- 1 Vibrant Cityscape | land | Kept as part of the precon mana base.
- 1 Vigor | other | Kept from the Turtle Power precon.
- 1 Wave Goodbye | other | Kept from the Turtle Power precon.
- 1 Colossal Skyturtle | wincon | Added as a listed Turtle finisher.

</details>

### 19. the Hobbit family, two colours

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 174 names.

Commander: Thranduil, the Elvenking.

Grade: bad, score 0.02, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $393.06 to buy, $393.06 the whole deck.

**Summary:** This Elf-centered Thranduil deck develops its mana early, builds a steady creature board, and keeps cards flowing so it can continue applying pressure after opponents answer the first wave. Spot removal, protective interaction, and board wipes help it reach the stage where Troll and Witch-king finishers can end the game. It gives up some raw speed for a thematic creature plan that is strongest when it can keep several bodies in play.

The quality model grades this deck below the precon baseline against the top lists of the format: the black sources cover 15.75 of the 23 its spells need.

- PLAN plan_coherent=partly: There is a loose "go wide with Middle-earth creatures and grind value" core, but the self-symmetrical wipes (Languish, Gnashing of Teeth) fight the creature-board plan, and the Elf-tribal payoffs (Elvish Archdruid, Celeborn) are stranded among spiders, bears, bees and one-off goodstuff bodies with no real tribal density.
- PLAN theme_fit=partly: It is a Commander deck led by the requested Thranduil at a reasonable bracket-3 power level, but a large share of the list (Delighted Halfling, Orcish Bowmasters, Lórien Revealed, Mox Amber, Languish, Night's Whisper) comes from outside the Hobbit set the person asked to build from.
- PLAN useful_as_built=partly: The deck can be shuffled up and has creatures, removal and finishers, but a three-color list with 36 lands that are all basics and no duals or fixing beyond a couple of ramp pieces will regularly stumble on the heavy black costs (Witch-king, Languish, Bitter Downfall).
- PLAN summary_honest=partly: Calling the list "Elf-centered" overstates a deck where a clear majority of the creatures are not Elves, and it presents the wipes as clean enablers when several of them would sweep its own board; the admission about giving up speed is fair, but the all-basic mana is not disclosed.
- [INFO] `curve_summary`: average mana value 3.38 over 63 nonland cards
- [WARN] `profile_off_band`: the worst color's share of the sources it needs is 0.68, and bracket 3 wants 0.85 or more (sources of requirement: U 14 of 19, B 15.75 of 23, G 16 of 22)
- [INFO] `mana_pass`: the builder moved 2 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 13 Forest | land | Provides a reliable green mana base.
- 11 Island | land | Provides a reliable blue mana base.
- 12 Swamp | land | Provides a reliable black mana base.
- 1 Arcane Signet | ramp | Efficiently supports the deck's mana development.
- 1 Bag End Banquet | ramp | Adds another mana-development piece.
- 1 Delighted Halfling | ramp | Provides early mana acceleration.
- 1 Elvish Archdruid | ramp | Accelerates mana while fitting the Elf core.
- 1 Elvish Mystic | ramp | Provides early mana acceleration.
- 1 Mox Amber | ramp | Supplies fast mana alongside the deck's legendary cards.
- 1 Necklace of Girion | ramp | Adds mana acceleration for the midgame.
- 1 Silvan Reveler | ramp | Contributes to the deck's ramp package.
- 1 Wood Elves | ramp | Provides mana development on a creature that supports the Elf theme.
- 1 Beorn the Fierce | draw | Supplies a draw option for longer games.
- 1 Bilbo Baggins, Burglar // Take a Glance | draw | Adds card advantage with flexible presentation.
- 1 Elvish Visionary | draw | Provides card advantage while fitting the Elf core.
- 1 Fateful Discovery | draw | Adds a dedicated card-advantage piece.
- 1 Gandalf, Shadow's Foe | draw | Provides card advantage from a substantial legendary body.
- 1 Hithlain Knots | draw | Adds an instant-speed draw option.
- 1 Key to the Side-Door | draw | Provides an additional source of card advantage.
- 1 Lórien Revealed | draw | Adds a draw spell that smooths later turns.
- 1 Night's Whisper | draw | Provides efficient card advantage.
- 1 Palantír of Orthanc | draw | Adds a persistent card-advantage option.
- 1 Plunder the Trollshaws | draw | Provides another draw effect for refueling.
- 1 Bilbo's Deadly Slice | removal | Adds a dedicated removal spell.
- 1 Bitter Downfall | removal | Adds direct removal for troublesome permanents.
- 1 Enchanted River's Grasp | removal | Provides another removal option.
- 1 Giant's Boulder | removal | Adds removal on an artifact.
- 1 Merciless Executioner | removal | Provides creature-based removal.
- 1 Orcish Bowmasters | removal | Adds efficient creature-based removal.
- 1 Quarrel | removal | Provides a compact removal spell.
- 1 The Black Arrow | removal | Adds removal that can be deployed as an artifact.
- 1 Uneasy Partings | removal | Rounds out the deck's spot-removal suite.
- 1 Bilbo's Ring | interaction | Provides a protective and disruptive interaction piece.
- 1 Confusticate and Bebother | interaction | Adds stack interaction.
- 1 Elrond, Moon-Reader | interaction | Provides interaction on a legendary Elf.
- 1 Mithril Coat | interaction | Helps protect an important creature.
- 1 My Precious // Allure of Power | interaction | Adds a flexible interaction option.
- 1 Old Fat Spider Can't See Me | interaction | Provides an interactive saga for the board.
- 1 Stern Scolding | interaction | Adds low-cost stack interaction.
- 1 Thranduil's Decree | interaction | Provides another themed interaction spell.
- 1 Arwen, Weaver of Hope | synergy | Supports the Elf-focused creature plan.
- 1 Boughside Wanderers | synergy | Adds an Elf body to support the deck's tribal core.
- 1 Cantankerous Keepers | synergy | Contributes another Elf creature to the board plan.
- 1 Celeborn the Wise | synergy | Adds a legendary Elf to reinforce the deck's central theme.
- 1 Galadhrim Guide | synergy | Supports the Elf-centered creature base.
- 1 Galion, Elvenking's Butler | synergy | Adds a thematic legendary Elf.
- 1 Lothlórien Lookout | synergy | Provides another Elf for the cohesive creature plan.
- 1 Attercop | threat | Provides a substantial creature threat.
- 1 Dreaded Bat-Cloud | threat | Adds board presence as an evasive creature threat.
- 1 Elvenking's Harper | threat | Adds a thematic creature that helps build pressure.
- 1 Gigantic Big Bear | threat | Provides a large creature for combat pressure.
- 1 Great Fierce Bee | threat | Adds another creature threat to develop the board.
- 1 Haunt of the Dead Marshes | threat | Provides a threatening creature from the deck's colors.
- 1 Mirkwood Elk | threat | Adds a creature that can apply combat pressure.
- 1 Mirkwood Nurturer | threat | Builds the creature board while maintaining the woodland theme.
- 1 Mirkwood Pathmaker | threat | Adds another creature for sustained board pressure.
- 1 Nimrodel Watcher | threat | Provides an Elf creature threat.
- 1 Ravenhill Flock | threat | Adds an evasive creature threat.
- 1 Willow-Wind | threat | Provides an additional creature threat for combat-focused games.
- 1 Troll of Khazad-dûm | wincon | Serves as a dedicated finishing threat.
- 1 Witch-king of Angmar | wincon | Provides a powerful legendary route to closing games.
- 1 Witch-king, Bringer of Ruin | wincon | Adds another dedicated finishing threat.
- 1 Gnashing of Teeth | wipe | Provides a reset button against developed boards.
- 1 Languish | wipe | Adds a broad creature-board answer.
- 1 Raise the Palisade | wipe | Provides a tribal-minded board wipe option.
- 1 Pelargir Survivor | ramp | the mana pass added it to bring the mana base inside the power level

</details>

### 20. the Hobbit family, a delegated commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 77 names.

Commander: Smaug the Magnificent.

Grade: baseline, score 0.05, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $462.76 to buy, $462.76 the whole deck.

**Summary:** Smaug leads a mana-forward Dragon deck that develops its board with legendary creatures and equipment before ending games through Cavern-Hoard Dragon, Desert Were-Worm, Smaug, the Great Calamity, or Dáin Ironfoot. Targeted removal and several reset effects keep opposing boards in check while the card-flow artifacts help maintain pressure. The deck gives up broad defensive flexibility for a focused, thematic board plan built around mana, Dragons, and iconic Hobbit-set artifacts.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: There is a loose red legends/equipment/goblin aggro shell that hangs together, but the stated Dragon axis is essentially absent and many cards (Guttersnipe, dwarf/goblin bodies, Oliphaunt, assorted equipment) pull toward unrelated subthemes rather than one plan.
- PLAN theme_fit=no: The person asked for a Dragons deck and got a mono-red Hobbit-set goodstuff pile of dwarves, goblins and equipment with barely any Dragons, so it answers a different request even though the set and format are right.
- PLAN useful_as_built=yes: Mono-red with 36 lands, plenty of ramp, a reasonable curve, removal and several attackers means it can be picked up and played functionally as it stands.
- PLAN summary_honest=no: It calls the list a "Dragon deck" and a "focused board plan built around... Dragons" when only two or three Dragons appear in the 99 plus the commander, and it lists Desert Were-Worm and Dáin Ironfoot as if they were part of that Dragon finish.
- [INFO] `curve_summary`: average mana value 3.11 over 63 nonland cards
- [WARN] `profile_off_band`: the draw count is 6, and bracket 3 wants 8 to 14

<details><summary>The deck list</summary>

- 1 Arcane Signet | ramp | Provides efficient mana acceleration.
- 1 Bag End Banquet | ramp | Supports the deck's mana development.
- 1 Burn, Burn, Tree and Fern | ramp | Adds another mana-development piece.
- 1 Dragon's Desire | ramp | Helps accelerate into the deck's larger threats.
- 1 Fíli and Kíli, Joyous | ramp | Provides creature-based mana acceleration.
- 1 Long-Bodied Grey Dog | ramp | Contributes to early mana development.
- 1 Mox Amber | ramp | Provides low-cost mana acceleration.
- 1 Orcrist, Goblin-cleaver | ramp | Adds to the deck's ramp package.
- 1 The Misty Mountains Cold | ramp | Supports mana development through a saga.
- 1 The Reaver Cleaver | ramp | Provides additional mana acceleration.
- 1 Wayfarer's Bauble | ramp | Helps establish mana early.
- 1 Balin, Loremaster | draw | Provides card flow from a creature slot.
- 1 Key to the Side-Door | draw | Adds a dedicated card-flow artifact.
- 1 Palantír of Orthanc | draw | Provides repeatable card-flow support.
- 1 Ragged Short Spear | draw | Adds card flow while fitting the equipment package.
- 1 Thrór's Map | draw | Provides another card-flow artifact.
- 1 Óin the Brave | draw | Adds creature-based card flow.
- 1 Bilbo's Ring | interaction | Provides a protective interaction piece.
- 1 Dwarven Mattock | interaction | Adds equipment-based interaction.
- 1 Mithril Coat | interaction | Provides a defensive interaction option.
- 1 The One Ring | interaction | Adds a powerful interaction artifact.
- 1 Battle-Scarred Goblin | removal | Provides creature-based removal.
- 1 Fire of Orthanc | removal | Supplies direct removal.
- 1 Gandalf, Spark Starter | removal | Adds removal on a legendary creature.
- 1 Giant's Boulder | removal | Provides artifact-based removal.
- 1 Goblin Cratermaker | removal | Adds flexible creature-based removal.
- 1 Goblin Fireleaper | removal | Provides another removal creature.
- 1 Improvised Club | removal | Supplies instant-speed removal.
- 1 Pinecone Strike | removal | Adds another instant removal spell.
- 1 Smite the Deathless | removal | Provides additional targeted removal.
- 1 Call Forth the Tempest | wipe | Provides a board-reset option.
- 1 Desolation of Smaug | wipe | Adds a second board-reset effect.
- 1 Glóin the Mighty // Easy Pickings | wipe | Provides a creature card with a wipe option.
- 1 Cavern-Hoard Dragon | wincon | A Dragon finisher for closing games.
- 1 Desert Were-Worm | wincon | A large finisher that supports the Dragon plan.
- 1 Smaug, the Great Calamity // Spew Flame | wincon | An additional Smaug-themed Dragon finisher.
- 1 Last Light of Durin's Day | synergy | Provides the available dedicated synergy piece.
- 1 Andúril, Flame of the West | other | Supports the deck's equipment subtheme.
- 1 Andúril, Narsil Reforged | other | Adds another legendary equipment piece.
- 1 Bombur, Gentle Dreamer | other | Adds a legendary creature to the board plan.
- 1 Bothersome Noisemaker | other | Provides an additional creature body.
- 1 Dori, Bearer of Friends | other | Adds another legendary creature.
- 1 Dwarven Mauler | other | Provides a Dwarf creature for the board.
- 1 Dwarven Warriors | other | Adds another Dwarf creature.
- 1 Dáin Ironfoot | other | Provides a finisher-class legendary creature.
- 1 Gandalf, Goblins' Bane // Flameshape | other | Adds a legendary creature with an Adventure option.
- 1 Getaway Barrel | other | Provides a utility artifact.
- 1 Glamdring | other | Adds another legendary equipment card.
- 1 Goblin-town Flunkies | other | Provides an additional Goblin creature.
- 1 Gundabad Opportunist | other | Adds another Goblin creature.
- 1 Guttersnipe | other | Supports the deck's collection of instants and sorceries.
- 1 Iron Hills Stalwart | other | Provides another Dwarf creature.
- 1 Long-Lost Lances | other | Adds to the equipment package.
- 1 Misty Mountains Raider | other | Provides another Goblin creature.
- 1 Old Thrush | other | Adds a small creature to the board plan.
- 1 Oliphaunt | other | Provides another creature threat.
- 1 Olog-hai Crusher | other | Adds a large creature to the board plan.
- 1 Orcish Siegemaster | other | Provides another creature body.
- 1 Smaug's Fury | other | Adds a Smaug-themed instant.
- 1 Snowslope Hunter | other | Provides another Goblin creature.
- 1 Sting, Bilbo's Sword | other | Adds another legendary equipment card.
- 1 Tidings of War | other | Provides a thematic sorcery slot.
- 1 Well-Worn Spatula | other | Adds another equipment card.
- 31 Mountain | land | Forms the core of the red mana base.
- 1 Dragon-Cursed Halls | land | A thematic nonbasic land for the mana base.
- 1 Hobbit Hole | land | Adds a nonbasic land slot to the mana base.
- 1 Rogue's Passage | land | Provides a utility land slot.
- 1 The Lonely Mountain | land | A thematic Mountain land for the mana base.
- 1 Treasure Vault | land | Provides a utility artifact-land slot.

</details>

### 21. the Hobbit family, mana from outside

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 156 names.

Commander: Smaug the Impenetrable.

Grade: baseline, score 0.04, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $918.72 to buy, $918.72 the whole deck.

**Summary:** This Smaug deck builds mana early, develops a board of dragons and legendary threats, and keeps opposing boards manageable with removal, protective interaction, and reset effects. It wins by turning its substantial threats and marked finishers into sustained pressure after clearing a path. The tradeoff is that the deck leans on its mana development and permanent-based board presence rather than on a highly reactive or combo-driven game plan.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. Many of the cards appear in no top list, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The list is a broadly coherent mono-goodstuff Rakdos legends/equipment pile that ramps into big threats and clears the way with removal and wipes, but the pieces pull between a legendary-equipment package (Andúril, Sting, Glamdring, Orcrist, Mithril Coat), goblin/orc tribal bodies, and one-off value cards, and the 'board of dragons' the summary names is barely present.
- PLAN theme_fit=partly: It is a Commander deck led by Smaug the Impenetrable at roughly precon-plus power with free-range lands and rocks as allowed, but a noticeable slice of the nonland spells come from outside the Hobbit set (Sting, Mithril Coat, Andúril, The One Ring, Palantír, Witch-king of Angmar) and Guttersnipe is not even a Middle-earth card, so the 'from the Hobbit set' limit is only loosely respected.
- PLAN useful_as_built=partly: 37 lands, ten-plus ramp pieces, ample removal and plenty of large finishers make it playable out of the box, but Misty Rainforest and Polluted Delta can fetch nothing in a Mountain/Swamp deck and effectively produce no mana, leaving two dead slots in the mana base.
- PLAN summary_honest=partly: The removal, ramp, protection and reset claims are backed by the list, but 'a board of dragons' is not — there are only a couple of dragon cards — and the summary does not disclose that several fetchlands in the mana base cannot fetch any land in the deck.
- [INFO] `curve_summary`: average mana value 3.27 over 62 nonland cards
- [INFO] `cards_trimmed`: the list was 1 card over, so the builder cut this card: Cavern-Hoard Dragon
- [WARN] `outside_requested_set`: the sets you named do not hold 17 cards: Ancient Tomb, Arid Mesa, Blood Crypt, Bloodstained Mire, Bojuka Bog, Command Tower, Dragonskull Summit, Evolving Wilds, Exotic Orchard, Fabled Passage, Marsh Flats, Misty Rainforest, Polluted Delta, Scalding Tarn, Urborg, Tomb of Yawgmoth, Verdant Catacombs, Wooded Foothills

<details><summary>The deck list</summary>

- 1 Ancient Tomb | land | An untapped land that helps accelerate expensive plays.
- 1 Arid Mesa | land | A flexible land slot that supports the mana base.
- 1 Blood Crypt | land | A dual-color land for the core mana base.
- 1 Bloodstained Mire | land | A flexible land slot that supports the mana base.
- 1 Bojuka Bog | land | A utility land that also supplies black mana.
- 1 Command Tower | land | Reliable access to the commander's colors.
- 1 Dragonskull Summit | land | A dual-color land for the core mana base.
- 1 Evolving Wilds | land | A flexible land slot that supports the mana base.
- 1 Exotic Orchard | land | A flexible multicolor land slot.
- 1 Fabled Passage | land | A flexible land slot that supports the mana base.
- 1 Marsh Flats | land | A flexible land slot that supports the mana base.
- 1 Misty Rainforest | land | A flexible land slot that supports the mana base.
- 10 Mountain | land | Primary red mana for the deck's color requirements.
- 1 Polluted Delta | land | A flexible land slot that supports the mana base.
- 1 Scalding Tarn | land | A flexible land slot that supports the mana base.
- 10 Swamp | land | Primary black mana for the deck's color requirements.
- 1 Urborg, Tomb of Yawgmoth | land | A black source and mana-base utility land.
- 1 Verdant Catacombs | land | A flexible land slot that supports the mana base.
- 1 Wooded Foothills | land | A flexible land slot that supports the mana base.
- 1 Arcane Signet | ramp | Efficient color fixing and mana acceleration.
- 1 Bag End Banquet | ramp | Part of the deck's mana acceleration package.
- 1 Bolg's Company | ramp | A ramp creature that contributes to early development.
- 1 Burn, Burn, Tree and Fern | ramp | Mana development that supports the deck's larger cards.
- 1 Dragon's Desire | ramp | Mana acceleration for the deck's top end.
- 1 Fíli and Kíli, Joyous | ramp | A ramp card that also adds to the board.
- 1 Long-Bodied Grey Dog | ramp | A ramp creature for early mana development.
- 1 Mox Amber | ramp | A fast mana piece backed by the deck's legendary cards.
- 1 Orcrist, Goblin-cleaver | ramp | Equipment-based mana acceleration.
- 1 Wayfarer's Bauble | ramp | Early mana acceleration that helps secure land drops.
- 1 Balin, Loremaster | draw | A creature-based source of cards.
- 1 Key to the Side-Door | draw | A dedicated card-advantage piece.
- 1 Night's Whisper | draw | An efficient source of card advantage.
- 1 Óin the Brave | draw | A creature that contributes to card advantage.
- 1 Palantír of Orthanc | draw | A durable card-advantage engine.
- 1 Rage into the Valley | draw | A card-advantage spell for keeping resources flowing.
- 1 Ragged Short Spear | draw | Equipment that contributes to card advantage.
- 1 Reverent Howl | draw | An instant-speed card-advantage option.
- 1 The Master of Lake-town | draw | A creature-based source of cards.
- 1 The Sackville-Bagginses | draw | A card-advantage creature.
- 1 Thrór's Map | draw | A reliable artifact source of card advantage.
- 1 Azog, Moria's Ruin | removal | A removal creature that advances the board.
- 1 Battle-Scarred Goblin | removal | A creature-based removal option.
- 1 Bilbo's Deadly Slice | removal | A direct removal spell.
- 1 Bitter Downfall | removal | A flexible removal spell.
- 1 Bolg of the North | removal | A legendary creature that provides removal.
- 1 Crude Bent Blade | removal | Equipment-based removal for the creature suite.
- 1 Fire of Orthanc | removal | A direct removal spell.
- 1 Gandalf, Spark Starter | removal | A removal creature that adds board presence.
- 1 Giant's Boulder | removal | An artifact-based removal option.
- 1 Bolg, Erebor's Reckoning | wipe | A creature-based board reset.
- 1 Call Forth the Tempest | wipe | A dedicated board-reset spell.
- 1 Desolation of Smaug | wipe | A thematic board-reset option.
- 1 Bilbo's Ring | interaction | A dedicated interaction piece.
- 1 Dwarven Mattock | interaction | Equipment-based interaction for the board.
- 1 Getaway Barrel | interaction | A utility artifact that broadens the interaction suite.
- 1 Glamdring | interaction | A utility equipment piece for board interaction.
- 1 Mithril Coat | interaction | A protective interaction option for key permanents.
- 1 My Precious // Allure of Power | interaction | A flexible interaction card.
- 1 Sting, Bilbo's Sword | interaction | A utility equipment piece for board interaction.
- 1 The One Ring | interaction | A resilient interaction artifact.
- 1 Gollum, Riddle Master | wincon | A marked finisher that provides a closing threat.
- 1 Smaug, Wicked Worm | wincon | A marked finisher that fits the dragon-focused endgame.
- 1 Witch-king, Bringer of Ruin | wincon | A marked finisher that closes games through board pressure.
- 1 Along the Crooked Way | synergy | A thematic support piece for the deck's synergy package.
- 1 Andúril, Flame of the West | synergy | A thematic equipment support piece.
- 1 Andúril, Narsil Reforged | synergy | A thematic equipment support piece.
- 1 Goblin Plate Mail | synergy | Equipment support for the deck's creature plan.
- 1 Last Light of Durin's Day | synergy | A thematic support permanent.
- 1 Long-Lost Lances | synergy | Equipment support for the deck's creature plan.
- 1 Supper for Spiders | synergy | The shortlist's dedicated synergy card.
- 1 Dáin Ironfoot | threat | A marked finisher that supplies a substantial threat.
- 1 Gandalf, Goblins' Bane // Flameshape | threat | A versatile creature threat with an Adventure option.
- 1 Great Goblin, Foul-Hearted | threat | A legendary creature that adds pressure to the board.
- 1 Guttersnipe | threat | A creature threat that rewards the deck's spells.
- 1 Olog-hai Crusher | threat | A large creature threat for the combat plan.
- 1 Sauron, the Lidless Eye | threat | A legendary creature that adds a meaningful threat.
- 1 Smaug, the Great Calamity // Spew Flame | threat | A marked finisher that gives the deck another dragon threat.
- 1 Troll of Khazad-dûm | threat | A marked finisher that supplies a large closing threat.
- 1 Witch-king of Angmar | threat | A marked finisher that applies substantial pressure.
- 1 Desolation Prowler | threat | An efficient creature threat for the board plan.
- 1 Orcish Siegemaster | threat | A creature threat that helps maintain board pressure.

</details>

### 22. a set family and a card from outside it

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 175 names.

Commander: Thranduil, the Elvenking.

Grade: baseline, score 0.04, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $479.87 to buy, $479.87 the whole deck.

**Summary:** This is a Sultai Elvenking creature deck that develops mana early, builds an Elf-heavy board, and maintains cards and answers through a broad mix of draw, interaction, and removal. It wins by establishing enough creature pressure to support its marked Troll and Witch-king finishers, while its wipes give it a way to recover from stalled boards. The tradeoff is a straightforward three-color basic-land mana base with fewer utility lands and less specialized fixing.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core is coherent — elf bodies, mana dorks, and a few big black finishers backed by cheap interaction — but three sweepers (Languish, Gnashing of Teeth, Raise the Palisade) actively fight the wide creature board the rest of the list builds, so the pieces pull against each other at the edges.
- PLAN theme_fit=partly: It is a legal Thranduil Sultai elf deck at a reasonable bracket-3 power with Sol Ring kept, but a large share of the list (Elvish Mystic, Elvish Archdruid, Night's Whisper, Languish, Orcish Bowmasters, The One Ring, Delighted Halfling, Arcane Signet) comes from outside the requested Hobbit set.
- PLAN useful_as_built=yes: 36 basics across three colors plus multiple dorks and rocks support a curve that tops out modestly, and the elf army plus the Troll and both Witch-kings give clear ways to close a game.
- PLAN summary_honest=yes: The summary accurately describes the Sultai creature/ramp shell, the draw-and-removal suite, the Troll/Witch-king top end, and even flags the plain basic-land mana base as a tradeoff, claiming nothing the list lacks.
- [INFO] `curve_summary`: average mana value 3.13 over 63 nonland cards
- [WARN] `profile_off_band`: the worst color's share of the sources it needs is 0.75, and bracket 3 wants 0.85 or more (sources of requirement: U 14.5 of 19, B 17.25 of 23, G 17 of 22)
- [INFO] `mana_pass`: the builder moved 5 cards of the mana base to bring the deck inside its power level
- [WARN] `outside_requested_set`: the sets you named do not hold 1 card: Sol Ring

<details><summary>The deck list</summary>

- 14 Forest | land | Provides reliable green mana for the Elf-focused core.
- 10 Island | land | Provides reliable blue mana for draw and interaction.
- 12 Swamp | land | Provides reliable black mana for removal and finishers.
- 1 Sol Ring | ramp | Kept as requested and supplies efficient mana acceleration.
- 1 Arcane Signet | ramp | Fixes mana while advancing the board early.
- 1 Delighted Halfling | ramp | Adds an early mana-producing creature to the board.
- 1 Elven Chorus | ramp | Provides Elf-aligned mana acceleration.
- 1 Elvish Archdruid | ramp | Rewards the Elf creature base with mana production.
- 1 Elvish Mystic | ramp | Supplies efficient early green mana.
- 1 Mox Amber | ramp | Provides low-cost acceleration alongside legendary creatures.
- 1 Necklace of Girion | ramp | Adds another noncreature mana source.
- 1 Wayfarer's Bauble | ramp | Supports early development through mana acceleration.
- 1 Wood Elves | ramp | Combines creature presence with mana development.
- 1 Beorn the Fierce | draw | Adds a draw card attached to a substantial creature.
- 1 Elvish Visionary | draw | Provides a low-cost Elf body and card flow.
- 1 Fateful Discovery | draw | Adds steady card access outside the creature suite.
- 1 Gandalf, Wandering Wizard | draw | Adds another creature-based draw option.
- 1 Hithlain Knots | draw | Provides instant-speed card access.
- 1 Lórien Revealed | draw | Adds flexible card access to the deck.
- 1 Night's Whisper | draw | Provides efficient black card draw.
- 1 Palantír of Orthanc | draw | Adds a durable artifact source of card flow.
- 1 Plunder the Trollshaws | draw | Provides another instant-speed draw option.
- 1 Confusticate and Bebother | interaction | Gives the deck a reactive instant-speed option.
- 1 Elrond, Moon-Reader | interaction | Adds interaction on a legendary Elf body.
- 1 Mithril Coat | interaction | Helps protect an important creature on the board.
- 1 My Precious // Allure of Power | interaction | Provides flexible interaction through an Equipment card.
- 1 Old Fat Spider Can't See Me | interaction | Adds an interactive Saga to the noncreature suite.
- 1 Stern Scolding | interaction | Provides a low-cost reactive spell.
- 1 The One Ring | interaction | Offers a resilient legendary interaction piece.
- 1 Bilbo's Deadly Slice | removal | Provides direct removal in black.
- 1 Bitter Downfall | removal | Adds efficient instant-speed removal.
- 1 Crude Bent Blade | removal | Gives the creature suite access to removal through Equipment.
- 1 Enchanted River's Grasp | removal | Adds enchantment-based removal.
- 1 Merciless Executioner | removal | Provides removal on a creature body.
- 1 Orcish Bowmasters | removal | Adds creature-based removal to the deck.
- 1 Quarrel | removal | Provides an additional compact removal spell.
- 1 The Black Arrow | removal | Adds removal on a legendary artifact.
- 1 Uneasy Partings | removal | Provides another instant-speed removal option.
- 1 Gnashing of Teeth | wipe | Provides a board-wide reset option.
- 1 Languish | wipe | Adds a second dependable board wipe.
- 1 Raise the Palisade | wipe | Provides a tribal-aware board reset.
- 1 Supper for Spiders | synergy | Supplies the dedicated synergy card from the shortlist.
- 1 Arwen, Weaver of Hope | synergy | Adds a legendary Elf that strengthens the Elvenking theme.
- 1 Boughside Wanderers | synergy | Builds the deck's Elf creature density.
- 1 Cantankerous Keepers | synergy | Adds another Elf body for the tribal core.
- 1 Celeborn the Wise | synergy | Provides a legendary Elf for the deck's cohesive creature plan.
- 1 Elvenking's Harper | synergy | Reinforces the Elf-centered board presence.
- 1 Galion, Elvenking's Butler | synergy | Adds another Elvenking-themed legendary Elf.
- 1 Elven Raft-Steerer | threat | Adds an Elf creature that contributes to the board.
- 1 Galadhrim Guide | threat | Provides an additional Elf threat.
- 1 Grey Havens Navigator | threat | Adds an Elf creature to pressure opponents.
- 1 Haunt of the Dead Marshes | threat | Provides a creature threat outside the Elf core.
- 1 Lothlórien Lookout | threat | Adds another Elf body to build a board.
- 1 Mirkwood Elk | threat | Provides a straightforward creature threat.
- 1 Mirkwood Meditator | threat | Adds an Elf creature to the developing board.
- 1 Mirkwood Nurturer | threat | Provides another Elf threat for the creature plan.
- 1 Mirkwood Pathmaker | threat | Adds another Elf attacker to the board.
- 1 Nimrodel Watcher | threat | Provides an additional Elf creature threat.
- 1 Ravenhill Flock | threat | Adds a creature body that can apply pressure.
- 1 Willow-Wind | threat | Provides a final creature threat for the board plan.
- 1 Troll of Khazad-dûm | wincon | Serves as one of the deck's marked finishing threats.
- 1 Witch-king of Angmar | wincon | Serves as a marked legendary finishing threat.
- 1 Witch-king, Bringer of Ruin | wincon | Provides a second marked Witch-king finishing threat.
- 1 Giant's Boulder | ramp | the mana pass added it to bring the mana base inside the power level
- 1 Pelargir Survivor | ramp | the mana pass added it to bring the mana base inside the power level
- 1 Woodland Weavemaster | ramp | the mana pass added it to bring the mana base inside the power level

</details>

### 23. two set families at once

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 203 names.

Commander: Kíli the Resourceful.

Grade: typical, score 0.43, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $226.28 to buy, $226.28 the whole deck.

**Summary:** This is a white creature-and-artifact Kíli deck that develops its mana, keeps cards flowing, and builds pressure through legendary Equipment, Dwarves, Rabbits, and artifact creatures. It controls key opposing pieces with a broad removal package, protects its board with compact interaction, and closes through its larger threats and finishers. It gives up multicolor flexibility and relies on its permanents remaining in play to turn its steady development into a winning board.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

This deck plan aims for 2 or more finishers, the cards that win the game, and the deck holds 1.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: Most cards support a white creature/artifact/equipment go-wide plan around Kíli, but the list is padded with unrelated goodstuff (Psychosis Crawler, Karn, stray Rabbit and Dwarf tribal payoffs with almost no tribe behind them) and three board wipes that fight the deck's own wide board.
- PLAN theme_fit=no: The person asked for a deck built from the Hobbit and Bloomburrow sets, yet the large majority of the list (Skyclave Apparition, Blade Splicer, Academy Manufactor, Sun Titan, Luminous Broodmoth, Karn, Jazal Goldmane, Swords to Plowshares, the Andúril equipment, etc.) comes from outside those two sets.
- PLAN useful_as_built=yes: Thirty-seven lands plus solid ramp, a deep removal suite, and a reasonable curve mean the deck casts its spells and can present a board that closes games, even if it leans on a single dedicated finisher.
- PLAN summary_honest=partly: placeholder
- [INFO] `curve_summary`: average mana value 3.00 over 63 nonland cards
- [WARN] `finisher_short`: the finisher count is 1, and this deck plan wants 2 or more (Angel of the Ruins)
- [INFO] `basics_added`: the list was 1 card short, so the builder added 1 basic land

<details><summary>The deck list</summary>

- 25 Plains | land | Provides reliable white mana.
- 1 Command Tower | land | Provides colored mana for the deck.
- 1 Castle Ardenvale | land | Adds a white land to the mana base.
- 1 Minas Tirith | land | Adds a white land to the mana base.
- 1 Evolving Wilds | land | Finds a basic land for mana fixing.
- 1 Fabled Passage | land | Finds a basic land for mana fixing.
- 1 Terramorphic Expanse | land | Finds a basic land for mana fixing.
- 1 Thriving Heath | land | Adds a white source to the mana base.
- 1 Uncharted Haven | land | Adds a flexible colored source.
- 1 Path of Ancestry | land | Provides a colored land for the creature base.
- 1 Lupinflower Village | land | Adds another land to support the curve.
- 1 Hobbit Hole | land | Adds another land to support the curve.
- 1 Sol Ring | ramp | Provides efficient artifact ramp.
- 1 Arcane Signet | ramp | Provides reliable colored ramp.
- 1 Fellwar Stone | ramp | Provides low-cost mana acceleration.
- 1 Mind Stone | ramp | Provides early artifact ramp.
- 1 Patchwork Banner | ramp | Provides artifact-based mana acceleration.
- 1 Wayfarer's Bauble | ramp | Provides land-based mana acceleration.
- 1 Ornithopter of Paradise | ramp | Provides creature-based mana acceleration.
- 1 Loyal Warhound | ramp | Provides creature-based ramp.
- 1 Season of the Burrow | ramp | Provides sorcery-speed ramp.
- 1 Solemn Simulacrum | ramp | Provides ramp on an artifact creature.
- 1 Belladonna Took | draw | Supplies a creature-based draw option.
- 1 Caretaker's Talent | draw | Provides repeatable draw support.
- 1 Circuit Mender | draw | Provides draw on an artifact creature.
- 1 Cut a Deal | draw | Provides a direct draw spell.
- 1 Idol of Oblivion | draw | Provides artifact-based card draw.
- 1 Inspiring Overseer | draw | Provides draw on a creature.
- 1 Mentor of the Meek | draw | Provides creature-focused draw support.
- 1 Skullclamp | draw | Provides equipment-based card draw.
- 1 Spirited Companion | draw | Provides early creature-based draw.
- 1 The Gaffer | draw | Provides a legendary creature draw engine.
- 1 The Queen of Dale | draw | Provides another creature-based draw source.
- 1 Bilbo's Gambit | interaction | Provides instant-speed interaction.
- 1 Dawn's Truce | interaction | Provides protective interaction.
- 1 Galadriel's Dismissal | interaction | Provides flexible instant-speed interaction.
- 1 Luminous Broodmoth | interaction | Provides creature-based interaction.
- 1 Mithril Coat | interaction | Provides equipment-based protection.
- 1 Reprieve | interaction | Provides efficient instant-speed interaction.
- 1 Selfless Spirit | interaction | Provides creature-based protection.
- 1 Swiftfoot Boots | interaction | Provides protective equipment support.
- 1 Banishing Light | removal | Provides enchantment-based removal.
- 1 Fiend Hunter | removal | Provides creature-based removal.
- 1 Generous Gift | removal | Provides flexible instant-speed removal.
- 1 Hanged Executioner | removal | Provides creature-based removal.
- 1 Loran of the Third Path | removal | Provides artifact and enchantment removal.
- 1 Skyclave Apparition | removal | Provides creature-based removal.
- 1 Swords to Plowshares | removal | Provides efficient instant-speed removal.
- 1 The Black Arrow | removal | Provides equipment-based removal.
- 1 Westfold Rider | removal | Provides creature-based removal.
- 1 Academy Manufactor | synergy | Supports the deck's artifact themes.
- 1 Blade Splicer | synergy | Supports the deck's artifact creature themes.
- 1 Carrot Cake | synergy | Supports the deck's artifact themes.
- 1 Dáin, Lord of the Iron Hills | synergy | Supports the Dwarf-focused creature package.
- 1 Iron Hills Blacksmith | synergy | Supports the artifact and Equipment package.
- 1 Ori, Keeper of Songs | synergy | Supports the Dwarf-focused creature package.
- 1 Tangle Tumbler | synergy | Supports the deck's artifact themes.
- 1 Dusk // Dawn | wipe | Provides a flexible board wipe.
- 1 Martial Coup | wipe | Provides a creature-focused board wipe.
- 1 Promise of Loyalty | wipe | Provides a selective board wipe.
- 1 Andúril, Flame of the West | threat | Serves as a legendary Equipment threat.
- 1 Andúril, Narsil Reforged | threat | Serves as a legendary Equipment threat.
- 1 Fíli the Pathfinder | threat | Provides a legendary creature threat.
- 1 Helm of the Host | threat | Provides a high-impact legendary Equipment threat.
- 1 Karn, the Great Creator | threat | Provides a planeswalker threat.
- 1 Psychosis Crawler | threat | Provides an artifact creature threat.
- 1 Serra Redeemer | threat | Provides a creature threat.
- 1 Starforged Sword | threat | Provides an Equipment threat.
- 1 Sun Titan | threat | Provides a large creature threat.
- 1 Sword of Vengeance | threat | Provides an Equipment threat.
- 1 Warren Warleader | threat | Provides a Rabbit creature threat.
- 1 Warren Elder | threat | Provides a Rabbit creature threat.
- 1 Angel of the Ruins | wincon | Provides a marked finisher for closing games.
- 1 Jazal Goldmane | wincon | Provides a legendary creature win condition.
- 1 Sunscorch Regent | wincon | Provides a Dragon win condition.

</details>

### 24. a 60-card deck from one set

Format: Modern. Theme: aggro. Pool: any_card. Shortlist: 96 names.

Grade: bad, score 0.32, model 20260914T154223Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $49.48 to buy, $49.48 the whole deck.

**Summary:** This mono-red Bloomburrow aggro deck builds pressure through a dense creature threat package backed by synergy cards, then uses removal to keep the way clear and draw effects to maintain momentum. It aims to win through sustained creature attacks, giving up broader color access and a more flexible game plan for a focused red assault.

The quality model grades this deck below the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. The color sources cover the pips, and that raises the grade.

- PLAN plan_coherent=yes: 22 cheap red creatures plus burn-style removal and two cheap card-advantage engines all push the same one-dimensional aggro clock, with only the two slow Season of the Bold copies sitting slightly off-curve.
- PLAN theme_fit=yes: It is a 60-card Modern-legal mono-red aggro list built entirely from Bloomburrow cards, exactly the request, and the casual power level matches the stated intent.
- PLAN useful_as_built=yes: 24 Mountains reliably cast a curve topping out at the two Dragonhawks, and the 22 creatures plus burn give plenty of redundant ways to close a game, though the land count is a touch high for a pure aggro deck.
- PLAN summary_honest=yes: It claims creature density, synergy pieces, removal and draw effects, and each of those categories is actually present in the list in the proportions described.
- [INFO] `curve_summary`: average mana value 3.28 over 36 nonland cards

<details><summary>The deck list</summary>

- 24 Mountain | land | Provides the mono-red mana base for the deck.
- 2 Season of the Bold | ramp | Fills the requested mana acceleration slot while staying within the red-focused plan.
- 4 Artist's Talent | draw | Provides a repeatable draw-focused element for the aggressive shell.
- 2 Might of the Meek | draw | Adds efficient draw support alongside the creature-heavy plan.
- 4 Agate Assault | removal | Supplies inexpensive removal to clear opposing obstacles.
- 2 Blooming Blast | removal | Completes the removal package with additional red answers.
- 4 Emberheart Challenger | synergy | Forms a core synergy piece for the creature-based aggro plan.
- 4 Hearthborn Battler | synergy | Supports the deck's aggressive creature synergies.
- 4 Frilled Sparkshooter | threat | Provides a full set of proactive threats for early pressure.
- 4 Reptilian Recruiter | threat | Adds another full playset of creature threats.
- 4 Teapot Slinger | threat | Keeps the threat density high for a dedicated aggro approach.
- 2 Dragonhawk, Fate's Tempest | threat | Rounds out the threat package with larger top-end pressure.

</details>

### 25. use no card of an owned precon

Format: Commander. Theme: Avengers superheroes. Pool: owned_first. Shortlist: 154 names.

Commander: Syr Konrad, the Grim.

Grade: baseline, score 0.31, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: yes, for copy_limit. Block findings: 0.

Cost: $0.00 to buy, $201.76 the whole deck.

**Summary:** An Avengers-inspired mono-black Commander deck built around grim battlefield attrition, villainous pressure, and the steady accumulation of death triggers. Counter placement, creature sacrifice, recursion, and broad creature suppression turn crowded boards into a resource engine, while protected legendary threats and imposing finishers close games once opponents have been worn down.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: The list consistently supports one engine: Syr Konrad plus mass -1/-1 counter placement (Blowfly Infestation, Midnight Banshee, Soul Snuffers, Grim Poppet, Contagion Clasp, Liliana, Black Sun's Zenith) to sweep small creatures and convert each death into drain, backed by removal, sweepers, and commander protection.
- PLAN theme_fit=no: The person asked for an Avengers superheroes deck, and the build delivers a mono-black Syr Konrad -1/-1 grindfest whose flavor is Nazgul, Liliana, and Final Fantasy gods, with only a handful of Marvel cards (Avengers Tower, M.O.D.O.K., Baron Strucker, Stolen Stark Tech) glued on.
- PLAN useful_as_built=yes: Thirty-seven lands with mono-black fixing, cheap ramp, a deep removal suite, card draw, and a clear Konrad-drain plus beatdown win path make this immediately playable.
- PLAN summary_honest=partly: The attrition/counter/death-trigger description matches the cards, but it calls the deck "Avengers-inspired" when the list is overwhelmingly Lord of the Rings/other-IP villains, and it names "creature sacrifice, recursion" though there is barely any sacrifice outlet and essentially no recursion package.
- [WARN] `not_owned`: Syr Konrad, the Grim: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.16 over 63 nonland cards

<details><summary>The deck list</summary>

- 24 Swamp | land | Reliable black mana base for a mono-black graveyard strategy.
- 1 Avengers Tower | land | Thematic headquarters land that supports the Avengers presentation.
- 1 Castle Locthwain | land | Land-based card advantage in longer games.
- 1 Command Tower | land | Efficient commander-focused mana fixing.
- 1 Evolving Wilds | land | Finds a Swamp while stocking the graveyard.
- 1 Exotic Orchard | land | Usually provides flexible mana in multiplayer.
- 1 Fabled Passage | land | Fetches a Swamp and supports graveyard resources.
- 1 Ifnir Deadlands | land | Black source with a useful counter-based utility effect.
- 1 Marsh Flats | land | Efficiently finds a Swamp and fuels graveyard synergies.
- 1 Path of Ancestry | land | Provides black mana and modest creature support.
- 1 Secluded Courtyard | land | Flexible creature-focused black source.
- 1 Takenuma, Abandoned Mire | land | Black source with late-game creature recursion.
- 1 Terramorphic Expanse | land | Finds a Swamp while contributing to graveyard resources.
- 1 Sol Ring | ramp | Efficient acceleration for deploying the deck's engines.
- 1 Arcane Signet | ramp | Reliable mana acceleration.
- 1 Wayfarer's Bauble | ramp | Develops mana while finding a basic land.
- 1 Thought Vessel | ramp | Mana acceleration with a useful hand-size benefit.
- 1 Sword of the Animist | ramp | Turns creature attacks into sustained land ramp.
- 1 Bender's Waterskin | ramp | Additional artifact mana development.
- 1 Chromatic Lantern | ramp | Mana acceleration that keeps utility lands functional.
- 1 Commander's Sphere | ramp | Reliable ramp that can later become a card.
- 1 Inherited Envelope | ramp | Artifact-based mana acceleration.
- 1 Inspiring Statuary | ramp | Lets the artifact package help cast larger spells.
- 1 Dusk Urchins | draw | A death-triggered source of cards that works with sacrifice and combat.
- 1 Hoarder's Greed | draw | Flexible black card advantage.
- 1 Idol of Oblivion | draw | Repeatable artifact-based card advantage.
- 1 Lembas | draw | A small value artifact that replaces itself later.
- 1 Mask of Memory | draw | Attack-based filtering and card advantage.
- 1 Massacre Girl, Known Killer | draw | Rewards creatures dying with sustained card flow.
- 1 Mirror of Galadriel | draw | Artifact card selection for assembling key pieces.
- 1 Mordor Muster | draw | Creature-focused card advantage.
- 1 Nasty End | draw | Efficient cards in exchange for a creature resource.
- 1 Night's Whisper | draw | Low-cost, dependable card draw.
- 1 Skullclamp | draw | Converts expendable creatures into substantial cards.
- 1 Bitter Downfall | removal | Versatile permanent removal.
- 1 Bitter Triumph | removal | Cheap answer to creatures or planeswalkers.
- 1 Claim the Precious | removal | Clean answer to a problematic creature.
- 1 Dismember | removal | Efficient creature removal.
- 1 Epic Downfall | removal | Exiles expensive opposing threats.
- 1 Fatal Push | removal | Low-cost creature interaction.
- 1 Heartless Act | removal | Flexible instant-speed creature answer.
- 1 Infernal Grasp | removal | Unconditional creature removal.
- 1 Nameless Inversion | removal | Instant-speed removal that can also enable creature-type lines.
- 1 Black Sun's Zenith | wipe | Scalable sweeper that leaves -1/-1 counters behind.
- 1 Extinction Event | wipe | Exiling board wipe against selected creature deployments.
- 1 M.O.D.O.K. | wipe | Avengers-themed mass-removal presence.
- 1 Champion's Helm | interaction | Protects and empowers the commander.
- 1 Darksteel Plate | interaction | Keeps a key creature alive through destruction effects.
- 1 Lightning Greaves | interaction | Protects key creatures and grants immediate mobility.
- 1 Swiftfoot Boots | interaction | Repeatable protection for the commander or a finisher.
- 1 Archfiend of Ifnir | wincon | A cycling-fueled board-control finisher.
- 1 Vraska, Betrayal's Sting | wincon | Proliferate pressure and a durable alternate path to victory.
- 1 Witch-king of Angmar | wincon | A resilient evasive finisher that punishes attacks.
- 1 Blowfly Infestation | synergy | Moves -1/-1 counters between creatures to fuel death cascades.
- 1 Contagion Clasp | synergy | Places counters early and proliferates them later.
- 1 Grim Poppet | synergy | Converts counters into repeatable creature control.
- 1 Incremental Blight | synergy | Spreads counters across opposing boards for death-trigger setups.
- 1 Midnight Banshee | synergy | Steadily distributes -1/-1 counters across the table.
- 1 One Ring to Rule Them All | synergy | Creates a thematic attrition sequence that feeds the graveyard plan.
- 1 Soul Snuffers | synergy | Distributes counters widely to set up cascading creature deaths.
- 1 Gaius van Baelsar | threat | A combat-capable villain that also contributes to board control.
- 1 Liliana, Death Wielder | threat | A persistent threat that supports counter-based attrition.
- 1 Orcish Bowmasters | threat | Efficient flash threat that pressures opposing card draw.
- 1 Ringwraiths | threat | A creature threat that contributes to removal pressure.
- 1 Skinrender | threat | Creature-based disruption that leaves behind a body.
- 1 Voracious Fell Beast | threat | A sizable creature that provides removal utility.
- 1 Zodiark, Umbral God | threat | Large, imposing threat with sweeping battlefield impact.
- 1 Rat King, Pale Piper | threat | A recursive creature threat with card-advantage utility.
- 1 Undercity Dire Rat | threat | Early creature development that supports the deck's resource plan.
- 1 Gorbag of Minas Morgul | threat | A creature threat that helps advance mana development.
- 1 Lobelia Sackville-Baggins | threat | Creature-based ramp that contributes to battlefield pressure.
- 1 Oft-Nabbed Goat | threat | A value creature that can help maintain resources.
- 1 Baron Strucker, HYDRA Overlord | other | Avengers-themed villain presence with long-game value.
- 1 Stolen Stark Tech | other | Avengers-themed equipment utility for important creatures.
- 1 Stone of Erech | other | Graveyard utility that protects the deck's longer game plan.
- 1 The Walls of Ba Sing Se | other | A durable defensive artifact creature for stabilizing the battlefield.

</details>

