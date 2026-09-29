# PR-8 deck gate

Run date: 2026-09-29. Card snapshot: 2026-09-04.

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
| Mean plan score, 0 to 1 | 0.76 |
| Plan reasons the judge left empty | 0 |
| Errors | 0 |
| Prompt version | 16 |
| Calls | 77 |
| Cost | $3.1278 |
| Time | 1993 seconds |

## Run

- Suite `decks`, run `pr8-deck-gate-run35`, on 2026-09-29, commit `3c85ca4`.
- Roles: generate on `gpt-5.6-terra` (openai, effort medium), judge on `claude-opus-5` (anthropic, effort medium), repair on `gpt-5.6-terra` (openai, effort medium).
- Versions: card snapshot 2026-09-04, generate prompt version 16, plan_rubric prompt version 4, summary_judge prompt version 2, precons `5.3.0+20260923`, quality_model `20260923T202806Z`.
- Calls: 77. Cost: $3.1278. Time: 1993 seconds.

## Findings by code

A block stops the deck. A warning and a note are reports.

| Code | Count |
|---|---|
| `curve_summary` | 25 |
| `not_owned` | 21 |
| `mana_pass` | 13 |
| `profile_off_band` | 6 |
| `outside_requested_set` | 2 |
| `finisher_short` | 2 |
| `basics_added` | 2 |
| `bracket_cut` | 1 |
| `precon_cards_restored` | 1 |

By severity: BLOCK 0. WARN 31. INFO 42. 

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

Format: Commander. Theme: lifegain. Pool: any_card. Shortlist: 306 names.

Commander: Karlov of the Ghost Council.

Grade: bad, score 0.03, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $761.88 to buy, $761.88 the whole deck.

**Summary:** This lifegain deck uses steady development and card advantage to make Karlov of the Ghost Council the central battlefield threat, with protection to preserve its key pieces and broad removal to clear resistance. Its creature suite applies pressure through lifegain payoffs while Angel of Destiny, Approach of the Second Sun, and Felidar Sovereign provide distinct ways to close a game. It gives up some speed for a resilient midrange plan that is strongest once its engines and mana are established.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every nonland card either gains life, pays off lifegain (Ajani's Pridemate, Archangel of Thune, Vito, Heliod, Nykthos Paragon), or protects/draws off that engine, all of which feeds Karlov's counters, so the pieces pull one direction.
- PLAN theme_fit=yes: It is a W/B Karlov lifegain list with no tutor-chains or two-card infinite combos assembled, matching the requested commander, theme, and bracket-3 power level.
- PLAN useful_as_built=yes: 36 lands with a heavy dual/basic mix support the WB costs, the curve is mostly two-to-five drops, and it has both combat threats and dedicated win conditions, though the ramp package (Colossal Plow, Hot Dog Cart, Orazca Relic) is unusually clunky.
- PLAN summary_honest=yes: The claims—card advantage, protection spells, removal and wraths, and the three named alternate win conditions—are all actually present in the list, and it openly concedes the deck is slow.
- [INFO] `curve_summary`: average mana value 3.48 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Command Tower | land | A reliable dual-color land for the opening turns.
- 1 Godless Shrine | land | A dual land that supports both halves of the deck.
- 1 Scrubland | land | An efficient dual land for white and black access.
- 1 Caves of Koilos | land | An untapped dual-color land.
- 1 Concealed Courtyard | land | An early dual-color land for a low curve.
- 1 Shattered Sanctum | land | A dual land that supports the core colors.
- 1 Vault of Champions | land | A dual land for reliable color fixing.
- 1 Silent Clearing | land | A dual-color land with late-game utility.
- 1 Brightclimb Pathway // Grimclimb Pathway | land | A flexible land choice for either core color.
- 1 Eclipsed Steppe | land | A Plains Swamp that fixes both colors.
- 1 Isolated Chapel | land | A dual land backed by a substantial basic base.
- 1 Shineshadow Snarl | land | A dual land that complements the basic lands.
- 1 Marsh Flats | land | A fetch land that finds either core color.
- 1 Prismatic Vista | land | A fetch land that turns into the needed basic color.
- 11 Plains | land | Basic white sources for the deck's white-heavy spells.
- 11 Swamp | land | Basic black sources for the deck's black spells.
- 1 Azor's Gateway // Sanctum of the Sun | ramp | A ramp piece with a powerful long-game land side.
- 1 Colossal Plow | ramp | An inexpensive ramp option for accelerating the board.
- 1 Hot Dog Cart | ramp | A low-cost ramp piece to improve early development.
- 1 Legion's Landing // Adanto, the First Fort | ramp | An early play that develops into a mana source.
- 1 Oasis Gardener | ramp | A cheap ramp body that helps cast Karlov on schedule.
- 1 Potioner's Trove | ramp | A ramp artifact that supports midgame development.
- 1 Nuka-Cola Vending Machine | ramp | A mana-focused artifact for sustained development.
- 1 Orazca Relic | ramp | A color-fixing ramp artifact.
- 1 Pristine Talisman | ramp | A mana artifact that belongs naturally in the lifegain plan.
- 1 The Celestus | ramp | A versatile mana artifact for the deck's setup turns.
- 1 Archivist of Oghma | draw | An efficient creature-based draw option.
- 1 Dawn of Hope | draw | A draw engine aligned with the deck's central plan.
- 1 Cosmos Elixir | draw | A steady source of cards over a longer game.
- 1 Enduring Innocence | draw | A durable draw engine for the creature suite.
- 1 Exemplar of Light | draw | A lifegain-themed creature that contributes cards.
- 1 Mangara, the Diplomat | draw | A dependable draw threat in multiplayer games.
- 1 Markov Purifier | draw | A Vampire draw piece that fits the creature base.
- 1 Sigarda's Splendor | draw | A dedicated lifegain draw engine.
- 1 The Gaffer | draw | A low-cost draw creature for the lifegain shell.
- 1 Tymna the Weaver | draw | A repeatable source of cards attached to a creature.
- 1 Well of Lost Dreams | draw | A major card-advantage outlet for lifegain turns.
- 1 Alseid of Life's Bounty | interaction | An inexpensive protective interaction piece.
- 1 Bofur, Reliable Guardian // Concerted Care | interaction | A flexible protection card for key permanents.
- 1 Courageous Resolve | interaction | A protection spell for pivotal board states.
- 1 Faith's Shield | interaction | A low-cost defensive interaction spell.
- 1 Metropolis Reformer | interaction | A creature-based protection option.
- 1 Restoration Magic | interaction | A flexible defensive spell for the board.
- 1 Rune-Tail, Kitsune Ascendant // Rune-Tail's Essence | interaction | A lifegain-friendly defensive permanent.
- 1 Werefox Bodyguard | interaction | A creature that provides immediate interaction.
- 1 Ayli, Eternal Pilgrim | removal | A low-cost removal creature that suits the deck's life total plan.
- 1 Gumdrop Poisoner // Tempt with Treats | removal | A flexible removal card with an on-theme alternate face.
- 1 Murderous Rider // Swift End | removal | Reliable flexible removal.
- 1 Nightmare's Thirst | removal | An efficient targeted removal spell.
- 1 Solitude | removal | A premium creature-based removal option.
- 1 Umezawa's Jitte | removal | A reusable equipment-based removal tool.
- 1 Vona, Butcher of Magan | removal | A lifegain-compatible removal threat.
- 1 Witch of the Moors | removal | A recurring removal payoff for lifegain turns.
- 1 Virtue of Persistence // Locthwain Scorn | removal | Flexible early removal with a strong late-game permanent.
- 1 Fumigate | wipe | A full reset that remains aligned with the deck's plan.
- 1 Kaya's Wrath | wipe | A reliable sweep when the board gets away.
- 1 White Sun's Twilight | wipe | A scalable sweeper for difficult board states.
- 1 Ajani's Pridemate | synergy | A classic lifegain payoff that grows beside Karlov.
- 1 Cleric Class | synergy | A dedicated lifegain engine and payoff.
- 1 Heliod, Sun-Crowned | synergy | A premier permanent for the deck's lifegain synergies.
- 1 Resplendent Angel | synergy | A lifegain payoff that adds meaningful board presence.
- 1 Sanguine Bond | synergy | A high-impact payoff for repeated lifegain.
- 1 Vito, Thorn of the Dusk Rose | synergy | A creature-based lifegain payoff.
- 1 Voice of the Blessed | synergy | A growing lifegain payoff that pressures opponents.
- 1 Archangel of Thune | threat | A powerful top-end threat for the lifegain strategy.
- 1 Attended Healer | threat | A creature threat that fits the deck's central theme.
- 1 Blood Baron of Vizkopa | threat | A durable Vampire threat for combat-focused games.
- 1 Celestine, the Living Saint | threat | A legendary threat with strong lifegain alignment.
- 1 Cliffhaven Vampire | threat | A lifegain payoff that adds pressure to the table.
- 1 Defiant Bloodlord | threat | A large black threat for the upper end of the curve.
- 1 Divinity of Pride | threat | A substantial lifegain-themed combat threat.
- 1 Nykthos Paragon | threat | A high-impact threat that rewards large lifegain turns.
- 1 Rhox Faithmender | threat | A major lifegain-focused threat.
- 1 Valkyrie Harbinger | threat | A powerful Angel threat for closing games.
- 1 Wurmcoil Engine | threat | A resilient colorless threat that stabilizes the board.
- 1 Angel of Vitality | threat | An efficient Angel threat that reinforces the theme.
- 1 Angel of Destiny | wincon | A dedicated finisher that gives the deck a distinct closing route.
- 1 Approach of the Second Sun | wincon | A noncombat finishing plan for stalled games.
- 1 Felidar Sovereign | wincon | A lifegain-centered alternate finishing threat.

</details>

### 2. aristocrats Commander, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 181 names.

Commander: Denethor, Ruling Steward.

Grade: baseline, score 0.32, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $166.72 the whole deck.

**Summary:** Denethor turns a steady supply of creatures into sacrifice pressure, draining opponents while building a fresh Soldier force after creatures die. The deck develops through inexpensive mana and card engines, protects its important pieces, and uses broad spot interaction plus selective board resets to keep opponents from getting ahead. It closes through dedicated finishers or by grinding repeated sacrifice turns into an advantage, giving up raw speed for a board-dependent plan that wants creatures and time to establish itself.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The mana base serves one color far better than another, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=no: The list names a sacrifice plan but carries almost no repeatable sacrifice outlets or token generators (Deadly Dispute and Nasty End are one-shots), while devoting slots to three board wipes, commander-protection auras/equipment and a pile of generic spot removal — a controlling good-stuff deck pulling against the handful of death-trigger payoffs.
- PLAN theme_fit=partly: It is a legal WB Commander deck at roughly bracket-3 power with an aristocrats-flavored commander, but the actual build plays as a removal/protection control shell rather than the aristocrats engine the player asked for, and its wipes even work against its own creature base.
- PLAN useful_as_built=partly: 36 lands plus eight mana rocks will cast the spells and the curve is low, but the 17 Plains to 8 Swamps split badly underserves the BB and black-heavy half of the list, and win conditions are thin — mostly incidental commander drain and a couple of midsize creatures.
- PLAN summary_honest=partly: It is candid about the slow, interaction-heavy, board-dependent build, but it asserts a 'steady supply of creatures' feeding 'sacrifice pressure' and 'dedicated finishers' that the 15-odd small creatures, minimal token production and near-absent sac outlets do not actually supply.
- [INFO] `curve_summary`: average mana value 2.68 over 63 nonland cards
- [INFO] `mana_pass`: the builder moved 4 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 1 Command Tower | land | Provides dependable access to both of the deck’s colors.
- 1 City of Brass | land | Provides flexible color fixing.
- 1 Exotic Orchard | land | Provides flexible color fixing in multiplayer games.
- 1 Grand Coliseum | land | Provides access to both deck colors.
- 1 Secluded Courtyard | land | Helps cast the deck’s creature-heavy core.
- 1 Unclaimed Territory | land | Helps cast the deck’s creature-heavy core.
- 1 Path of Ancestry | land | Provides dual-color fixing for the creature base.
- 1 Evolving Wilds | land | Finds whichever basic color is needed.
- 1 Fabled Passage | land | Finds whichever basic color is needed.
- 1 Marsh Flats | land | Finds either of the deck’s basic land types.
- 1 Terramorphic Expanse | land | Finds whichever basic color is needed.
- 17 Plains | land | Supplies reliable white mana.
- 8 Swamp | land | Supplies reliable black mana.
- 1 Sol Ring | ramp | Accelerates the deck’s early development.
- 1 Arcane Signet | ramp | Provides efficient color fixing and mana acceleration.
- 1 Fellwar Stone | ramp | Provides low-cost mana acceleration.
- 1 Wayfarer's Bauble | ramp | Builds the mana base early.
- 1 Springleaf Drum | ramp | Turns a spare creature into mana.
- 1 Inherited Envelope | ramp | Provides additional mana acceleration.
- 1 Relic of Legends | ramp | Provides mana acceleration for the deck’s legendary creatures.
- 1 Thought Vessel | ramp | Provides durable mana acceleration.
- 1 Lotho, Corrupt Shirriff | ramp | Adds mana production on a creature body.
- 1 Deadly Dispute | ramp | Converts a disposable permanent into mana and resources.
- 1 Commander's Sphere | ramp | Fixes colors while advancing mana.
- 1 White Auracite | ramp | Provides another mana-producing artifact.
- 1 Skullclamp | draw | Turns small creatures into a steady source of cards.
- 1 Idol of Oblivion | draw | Supplies repeatable card advantage alongside token production.
- 1 Night's Whisper | draw | Provides efficient early card flow.
- 1 Nasty End | draw | Converts a creature into more cards.
- 1 Lembas | draw | Provides a compact source of card advantage.
- 1 Stone of Erech | draw | Adds card access from a low-cost artifact.
- 1 Tome of Legends | draw | Builds repeatable card advantage around the commander.
- 1 Wall of Omens | draw | Provides an early creature and replaces itself.
- 1 Puresteel Paladin | draw | Rewards the deck’s equipment package with cards.
- 1 Call of the Ring | draw | Provides ongoing card access.
- 1 Inspiring Overseer | draw | Adds a creature while providing card advantage.
- 1 Exemplar of Light | draw | Provides creature-based card advantage.
- 1 June, Bounty Hunter | draw | Adds another creature-based draw engine.
- 1 Lightning Greaves | interaction | Protects Denethor and key creatures efficiently.
- 1 Swiftfoot Boots | interaction | Protects Denethor and key creatures.
- 1 Gift of Immortality | interaction | Helps preserve an important creature through removal.
- 1 Together Forever | interaction | Supports the deck’s creatures through attrition.
- 1 Duty Beyond Death | interaction | Protects the board from opposing disruption.
- 1 Clever Concealment | interaction | Protects the board during a critical turn.
- 1 Reprieve | interaction | Temporarily answers a troublesome opposing spell.
- 1 Bastion Protector | interaction | Keeps Denethor better protected on the battlefield.
- 1 Unbreakable Formation | interaction | Helps preserve the creature board.
- 1 Take Up the Shield | interaction | Provides a compact protective response for a key creature.
- 1 Bitter Triumph | removal | Provides flexible creature or planeswalker removal.
- 1 Claim the Precious | removal | Answers a problematic opposing creature.
- 1 Crib Swap | removal | Provides creature removal at instant speed.
- 1 Fiend Hunter | removal | Places removal on a creature body.
- 1 Generous Gift | removal | Answers a wide range of problematic permanents.
- 1 Get Lost | removal | Provides efficient permanent removal.
- 1 Infernal Grasp | removal | Provides direct creature removal.
- 1 Orcish Bowmasters | removal | Provides creature-based removal utility.
- 1 Swords to Plowshares | removal | Provides efficient creature removal.
- 1 Fatal Push | removal | Provides low-cost creature removal.
- 1 Heartless Act | removal | Provides flexible creature removal.
- 1 Austere Command | wipe | Resets selected opposing board elements.
- 1 Dusk // Dawn | wipe | Provides a creature-board reset with later value.
- 1 Fumigate | wipe | Clears a crowded creature board.
- 1 Gollum the Abandoned | synergy | Supports the deck’s sacrifice-focused game plan.
- 1 Gollum, Patient Plotter | synergy | Supports the deck’s sacrifice-focused game plan.
- 1 Gríma Wormtongue | synergy | Adds another creature that supports the core plan.
- 1 Heirloom Auntie | synergy | Supports the deck’s creature-centric engine.
- 1 Joo Dee, One of Many | synergy | Adds support for the deck’s main creature plan.
- 1 Arcade Cabinet | synergy | Provides a noncreature support piece for the core plan.
- 1 Phantom Train | synergy | Adds a synergistic permanent to the deck’s engine.
- 1 Bill the Pony | threat | Provides an on-board threat that contributes to combat pressure.
- 1 Hei Bai, Spirit of Balance | threat | Provides a substantial creature threat.
- 1 Namazu Trader | threat | Adds another creature that pressures opponents.
- 1 Vengeful Villagers | threat | Adds a board-based threat for the creature plan.
- 1 Al Bhed Salvagers | wincon | Serves as a dedicated finishing threat.
- 1 Archfiend of Ifnir | wincon | Provides a powerful finishing threat.
- 1 Grave Venerations | wincon | Provides a dedicated route to close the game.

</details>

### 3. artifacts Commander, bracket 4

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 318 names.

Commander: Urza, Lord High Artificer.

Grade: typical, score 0.64, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $4291.43 to buy, $4291.43 the whole deck.

**Summary:** Urza leads a fast blue artifact deck that uses early mana, artifact tutors, and sustained card advantage to establish a dense board quickly. The deck protects its position with extensive stack interaction, answers troublesome permanents through artifact-based removal, and closes through artifact creature pressure or alternate endgames from Mechanized Production, Mirrodin Besieged, and Thassa's Oracle. It gives up broad multicolor answers and relies heavily on artifacts remaining central to its plan.

The quality model grades this deck typical against the top lists of the format: the deck makes more mana on turn four than the norm, and that raises the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

- PLAN plan_coherent=partly: The bulk of the list coheres around cheap artifacts, artifact tutors, cost reduction and Urza's token/mana engine, but Thassa's Oracle sits in the deck with no self-mill or library-exiling enabler, and a few slots (Cargo Ship, Arc Reactor, Into Thin Air, Esoteric Duplicator) are filler that does not advance the fast artifact plan.
- PLAN theme_fit=partly: It is genuinely a mono-blue Urza artifact Commander deck with real fast mana (Mox Opal, Mana Vault, Lotus Petal, Sol Ring) and tutors, but several low-impact draft-level cards and some off-color utility lands in a mono-blue shell pull it under the bracket-4 high-power bar the request named.
- PLAN useful_as_built=yes: Thirty-four lands plus a pile of zero- and one-mana rocks support a low curve, and the deck has plenty of real ways to win through artifact creatures, Karn, Aetherflux Reservoir and the two enchantment alt-wins, so it can be picked up and played as printed.
- PLAN summary_honest=partly: The claims about artifact tutors, stack interaction, card draw, Mechanized Production and Mirrodin Besieged are all backed by the list, but calling Thassa's Oracle an alternate endgame is an overclaim since nothing in the deck empties the library.
- [INFO] `curve_summary`: average mana value 3.00 over 66 nonland cards
- [INFO] `mana_pass`: the builder moved 3 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 25 Island | land | Primary blue source for the deck's artifact-heavy plan.
- 1 Command Tower | land | Untapped blue-producing land.
- 1 City of Brass | land | Flexible blue source.
- 1 Exotic Orchard | land | Flexible blue-producing land.
- 1 Glimmervoid | land | Artifact-focused blue source.
- 1 Mana Confluence | land | Flexible blue source.
- 1 Reflecting Pool | land | Additional flexible blue source.
- 1 Seat of the Synod | land | Artifact land that supplies blue.
- 1 Spire of Industry | land | Artifact-focused blue source.
- 1 Braided Net // Braided Quipu | draw | Artifact-based card advantage.
- 1 Cerebral Download | draw | Card advantage for the deck.
- 1 Esoteric Duplicator | draw | Artifact-based card advantage.
- 1 Forensic Gadgeteer | draw | Creature-based card advantage.
- 1 Nexus of Becoming | draw | Artifact-based card advantage.
- 1 Reverse Engineer | draw | Efficient card advantage.
- 1 Rhystic Study | draw | Game-changing source of card advantage.
- 1 Riddlesmith | draw | Artifact-focused card selection.
- 1 Sai, Master Thopterist | draw | Artifact-focused card advantage.
- 1 Thoughtcast | draw | Efficient artifact-focused card advantage.
- 1 Thirst for Knowledge | draw | Instant-speed card advantage.
- 1 An Offer You Can't Refuse | interaction | Low-cost protection and disruption.
- 1 Disruption Protocol | interaction | Artifact-focused stack interaction.
- 1 Etched Champion | interaction | Artifact creature that supplies interaction.
- 1 Fierce Guardianship | interaction | Game-changing free interaction while Urza is in play.
- 1 Force of Will | interaction | Game-changing free stack interaction.
- 1 Ghostly Flicker | interaction | Flexible instant-speed interaction.
- 1 Ice Out | interaction | Efficient stack interaction.
- 1 Metallic Rebuke | interaction | Artifact-focused stack interaction.
- 1 Padeem, Consul of Innovation | interaction | Artifact-focused protective interaction.
- 1 Stoic Rebuttal | interaction | Reliable stack interaction.
- 1 Welding Jar | interaction | Cheap artifact protection.
- 1 Chrome Mox | ramp | Game-changing fast mana.
- 1 Lion's Eye Diamond | ramp | Game-changing fast mana.
- 1 Lotus Petal | ramp | Fast mana for explosive early turns.
- 1 Mana Vault | ramp | Game-changing fast mana.
- 1 Mox Diamond | ramp | Game-changing fast mana.
- 1 Mox Opal | ramp | Fast mana in an artifact-heavy deck.
- 1 Moonsnare Prototype | ramp | Fast artifact-based mana.
- 1 Sol Ring | ramp | Efficient fast mana.
- 1 Arc Reactor | ramp | Artifact-based mana acceleration.
- 1 Automated Artificer | ramp | Artifact creature that accelerates mana.
- 1 Chief Engineer | ramp | Artifact-focused mana acceleration.
- 1 Grand Architect | ramp | Artifact-focused mana acceleration.
- 1 Metalworker | ramp | High-output artifact mana acceleration.
- 1 Aether Spellbomb | removal | Cheap artifact-based removal.
- 1 Aetherflux Reservoir | removal | Artifact removal option and pressure piece.
- 1 Arcum Dagsson | removal | Artifact removal with tutor utility.
- 1 Blasting Station | removal | Artifact-based removal outlet.
- 1 Contagion Clasp | removal | Artifact-based removal.
- 1 Cyber Conversion | removal | Flexible removal spell.
- 1 Into Thin Air | removal | Instant-speed removal.
- 1 Kitesail Larcenist | removal | Creature-based removal.
- 1 Resculpt | removal | Efficient flexible removal.
- 1 Skysovereign, Consul Flagship | removal | Artifact removal that supports the board.
- 1 Reshape | synergy | Artifact tutor that turns available artifacts into a needed piece.
- 1 Transmute Artifact | synergy | Premium artifact tutor for assembling key pieces.
- 1 Whir of Invention | synergy | Instant-speed artifact tutor for the deck's toolbox.
- 1 Cyberdrive Awakener | threat | Artifact-focused closing threat.
- 1 Darksteel Juggernaut | threat | Large artifact creature threat.
- 1 Gearseeker Serpent | threat | Artifact-focused finisher threat.
- 1 Kappa Cannoneer | threat | High-impact artifact creature threat.
- 1 Karn, Scion of Urza | threat | Artifact-focused planeswalker threat.
- 1 Kuldotha Forgemaster | threat | Threat that also provides artifact tutor utility.
- 1 Master Transmuter | threat | Artifact creature threat.
- 1 Myr Enforcer | threat | Efficient artifact creature threat.
- 1 Sojourner's Companion | threat | Efficient artifact creature threat.
- 1 Traxos, Scourge of Kroog | threat | Large legendary artifact threat.
- 1 Mechanized Production | wincon | Artifact-based alternate endgame.
- 1 Mirrodin Besieged | wincon | Artifact-focused alternate endgame.
- 1 Thassa's Oracle | wincon | Game-changing creature-based endgame.
- 1 Engineered Explosives | wipe | Flexible artifact board wipe.
- 1 Hurkyl's Recall | wipe | Artifact-focused board reset.
- 1 Mox Amber | ramp | the mana pass added it to bring the mana base inside the power level
- 1 Cargo Ship | ramp | the mana pass added it to bring the mana base inside the power level
- 1 Era of Innovation | draw | the mana pass added it to bring the mana base inside the power level

</details>

### 4. dinosaur tribal, bracket 2

Format: Commander. Theme: dinosaurs. Pool: any_card. Shortlist: 299 names.

Commander: Gishath, Sun's Avatar.

Grade: bad, score 0.04, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $1314.09 to buy, $1314.09 the whole deck.

**Summary:** This is a straightforward Dinosaur combat deck: build mana, develop Dinosaur support, then bring Gishath, Sun's Avatar into combat to turn its trigger into a large board of Dinosaurs. The main finish is an overwhelming attack backed by large Dinosaur threats and finishers, with removal and a few reset buttons available when the board gets crowded. It favors committing creatures and pressing combat over a highly reactive game plan, so careful sequencing and protecting Gishath matter.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The land count sits above the norm of the format, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Every slot supports one plan — ramp and tribal cost reducers into big Dinosaurs, tribal anthems and evasion, and Gishath's combat trigger to dump more Dinosaurs — with nothing pulling toward an unrelated strategy.
- PLAN theme_fit=partly: It is unmistakably the requested Gishath Dinosaur Commander deck, but the manabase of fetchlands, dual lands, pain/horizon lands with only two basics plus staples like Esper Sentinel, Sylvan Library and Skullclamp is both more expensive and more decision-dense than a bracket-2 deck for a new player.
- PLAN useful_as_built=yes: Thirty-seven lands with heavy fixing, ten-plus ramp pieces, a reasonable curve topping at Gishath, plenty of card draw and multiple ways to close out a game make this immediately playable out of the box.
- PLAN summary_honest=partly: The core description of ramp-into-Gishath combat matches the list, but calling Vandalblast, Forerunner of the Empire and Raging Swordtooth "reset buttons" oversells small pingers/an artifact sweeper, and the summary never flags that the deck leans on fetches, original duals and painlands rather than the beginner-friendly build the request asked for.
- [INFO] `curve_summary`: average mana value 3.26 over 62 nonland cards

<details><summary>The deck list</summary>

- 1 Command Tower | land | Three-color land fixing for the deck.
- 1 City of Brass | land | Flexible land fixing for the deck.
- 1 Mana Confluence | land | Flexible land fixing for the deck.
- 1 Forbidden Orchard | land | Flexible land fixing for the deck.
- 1 Tarnished Citadel | land | Flexible land fixing for the deck.
- 1 Arid Mesa | land | Fetch land fixing for the deck.
- 1 Prismatic Vista | land | Fetch land fixing for the deck.
- 1 Windswept Heath | land | Fetch land fixing for the deck.
- 1 Wooded Foothills | land | Fetch land fixing for the deck.
- 1 Bountiful Promenade | land | Green-white fixing land.
- 1 Spire Garden | land | Red-green fixing land.
- 1 Spectator Seating | land | Red-white fixing land.
- 1 Battlefield Forge | land | Red-white fixing land.
- 1 Brushland | land | Green-white fixing land.
- 1 Karplusan Forest | land | Red-green fixing land.
- 1 Branchloft Pathway // Boulderloft Pathway | land | Green-white fixing land.
- 1 Cragcrown Pathway // Timbercrown Pathway | land | Red-green fixing land.
- 1 Needleverge Pathway // Pillarverge Pathway | land | Red-white fixing land.
- 1 Sacred Foundry | land | Red-white fixing land.
- 1 Stomping Ground | land | Red-green fixing land.
- 1 Temple Garden | land | Green-white fixing land.
- 1 Canopy Vista | land | Green-white fixing land.
- 1 Cinder Glade | land | Red-green fixing land.
- 1 Plateau | land | Red-white fixing land.
- 1 Savannah | land | Green-white fixing land.
- 1 Taiga | land | Red-green fixing land.
- 1 Cavern of Souls | land | Creature-focused color fixing.
- 1 Secluded Courtyard | land | Creature-focused color fixing.
- 1 Unclaimed Territory | land | Creature-focused color fixing.
- 1 Sunbaked Canyon | land | Red-white fixing land.
- 1 Horizon Canopy | land | Green-white fixing land.
- 1 Grove of the Burnwillows | land | Red-green fixing land.
- 1 Rockfall Vale | land | Red-green fixing land.
- 1 Rootbound Crag | land | Red-green fixing land.
- 1 Sunpetal Grove | land | Green-white fixing land.
- 1 Forest | land | Basic green mana source.
- 1 Plains | land | Basic white mana source.
- 1 Arcane Signet | ramp | Reliable color fixing and early ramp.
- 1 Fellwar Stone | ramp | Low-cost mana acceleration.
- 1 Drover of the Mighty | ramp | Creature-based ramp for the Dinosaur plan.
- 1 Intrepid Paleontologist | ramp | Creature-based ramp for the Dinosaur plan.
- 1 Ixalli's Lorekeeper | ramp | Creature-based ramp for the Dinosaur plan.
- 1 Pillar of Origins | ramp | Tribal color fixing and ramp.
- 1 Thunderherd Migration | ramp | Dinosaur-focused mana development.
- 1 Topiary Stomper | ramp | Dinosaur ramp that also supports Gishath hits.
- 1 Wayward Swordtooth | ramp | Dinosaur ramp that supports the board plan.
- 1 Ranging Raptors | ramp | Dinosaur ramp that supports the creature plan.
- 1 Esper Sentinel | draw | Early draw support.
- 1 Faithless Looting | draw | Low-cost card selection and draw support.
- 1 Folk Hero | draw | Creature-focused draw support.
- 1 Idol of Oblivion | draw | Low-cost draw support.
- 1 Skullclamp | draw | Efficient draw support around small creatures.
- 1 Sylvan Library | draw | Early draw support and card selection.
- 1 Beast Whisperer | draw | Creature-focused draw support.
- 1 Garruk's Uprising | draw | Draw support for the large-creature plan.
- 1 Ripjaw Raptor | draw | Dinosaur draw support for Gishath hits.
- 1 Runic Armasaur | draw | Dinosaur draw support for Gishath hits.
- 1 Return of the Wildspeaker | draw | Flexible draw support for a large board.
- 1 Heroic Intervention | interaction | Interaction to preserve the creature board.
- 1 Lightning Greaves | interaction | Interaction that helps safeguard a key creature.
- 1 Swiftfoot Boots | interaction | Interaction that helps safeguard a key creature.
- 1 Steely Resolve | interaction | Tribal interaction for the Dinosaur board.
- 1 Veil of Summer | interaction | Low-cost interaction slot.
- 1 Kindred Boon | interaction | Tribal interaction for the Dinosaur board.
- 1 Itzquinth, Firstborn of Gishath | removal | Low-cost Dinosaur removal for Gishath hits.
- 1 Savage Stomp | removal | Dinosaur-focused removal.
- 1 Thrashing Brontodon | removal | Dinosaur removal that remains a Gishath hit.
- 1 Bronzebeak Foragers | removal | Dinosaur removal that remains a Gishath hit.
- 1 Needletooth Raptor | removal | Dinosaur removal that remains a Gishath hit.
- 1 Ravenous Sailback | removal | Dinosaur removal that remains a Gishath hit.
- 1 Kogla and Yidaro | removal | Dinosaur removal for the creature plan.
- 1 Commune with Dinosaurs | synergy | Dinosaur synergy and early setup.
- 1 Kinjalli's Caller | synergy | Dinosaur tribal support.
- 1 Otepec Huntmaster | synergy | Dinosaur tribal support.
- 1 Marauding Raptor | synergy | Dinosaur tribal support.
- 1 Herald's Horn | synergy | Tribal Dinosaur synergy.
- 1 Urza's Incubator | synergy | Tribal Dinosaur synergy.
- 1 Descendants' Path | synergy | Dinosaur tribal synergy.
- 1 Shared Animosity | synergy | Combat-focused Dinosaur tribal synergy.
- 1 Icon of Ancestry | synergy | Dinosaur tribal synergy.
- 1 Radiant Destiny | synergy | Dinosaur tribal synergy.
- 1 Pantlaza, Sun-Favored | threat | Efficient Dinosaur threat for Gishath hits.
- 1 Regisaur Alpha | threat | Dinosaur threat for the combat plan.
- 1 Quartzwood Crasher | threat | Dinosaur threat for the combat plan.
- 1 Carnage Tyrant | threat | Large Dinosaur threat for Gishath hits.
- 1 Ghalta and Mavren | threat | Legendary Dinosaur threat for Gishath hits.
- 1 Goring Ceratops | threat | Large Dinosaur threat and finisher.
- 1 Shifting Ceratops | threat | Dinosaur threat and finisher.
- 1 Cacophodon | threat | Dinosaur threat for Gishath hits.
- 1 Sun-Blessed Mount | threat | Dinosaur threat for Gishath hits.
- 1 Sun-Crested Pterodon | threat | Dinosaur threat for Gishath hits.
- 1 Stampeding Horncrest | threat | Dinosaur threat for Gishath hits.
- 1 Balamb T-Rexaur | threat | Dinosaur threat for Gishath hits.
- 1 Bonehoard Dracosaur | wincon | Dinosaur finisher for the combat plan.
- 1 Dinosaurs on a Spaceship | wincon | Dinosaur finisher for the combat plan.
- 1 Huatli, Poet of Unity // Roar of the Fifth People | wincon | Finisher that supports the Dinosaur plan.
- 1 Vandalblast | wipe | Low-cost board wipe option.
- 1 Forerunner of the Empire | wipe | Dinosaur-focused board wipe option.
- 1 Raging Swordtooth | wipe | Dinosaur board wipe option for Gishath hits.

</details>

### 5. blink Commander, owned first

Format: Commander. Theme: blink. Pool: owned_first. Shortlist: 169 names.

Commander: Gilraen, Dúnedain Protector.

Grade: typical, score 0.41, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $132.41 the whole deck.

**Summary:** Gilraen anchors a creature-focused blink plan, repeatedly setting up favorable returns while the deck develops mana, cards, equipment, and a protected board. It wins by turning its growing creature force sideways, with Angel of Serenity chief among its closing threats, while removal and wipes clear away resistance. The deck gives up multicolor flexibility for a very consistent Plains-based mana base and relies on creatures staying relevant on the battlefield.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

This deck plan aims for 2 or more finishers, the cards that win the game, and the deck holds 1.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=no: The pieces pull apart: a go-wide creature board is undercut by three unconditional wipes, a heavy artifact/equipment package is taxed by the deck's own Kataki, War's Wage, and the blink plan the summary names is carried by only two or three actual flicker effects.
- PLAN theme_fit=partly: The commander and a handful of ETB creatures (Flickerwisp, Wall of Omens, Inspiring Overseer, Angel of Serenity) gesture at blink, but the bulk of the deck is generic mono-white equipment, ramp, removal, and wipes rather than the blink deck asked for, though format, bracket-3 feel, and library-only constraint are respected.
- PLAN useful_as_built=partly: 36 Plains and ample ramp give a clean mana base and a fine curve, but with only one true finisher and three of its own board wipes pointed at its creature-based clock the deck can struggle to actually close a game.
- PLAN summary_honest=no: It advertises a repeatable blink plan and equipment-based development while the list holds almost no blink enablers and quietly includes symmetrical wipes and an artifact tax that fight its own artifacts.
- [INFO] `curve_summary`: average mana value 2.71 over 63 nonland cards
- [WARN] `finisher_short`: the finisher count is 1, and this deck plan wants 2 or more (Angel of Serenity)

<details><summary>The deck list</summary>

- 36 Plains | land | A reliable white mana base for casting Gilraen and the deck's white spells.
- 1 Sol Ring | ramp | Efficient artifact ramp for accelerating the deck's development.
- 1 Arcane Signet | ramp | Reliable colored ramp in a mono-white shell.
- 1 Fellwar Stone | ramp | Low-cost artifact ramp.
- 1 Wayfarer's Bauble | ramp | Early ramp that supports reaching Gilraen on schedule.
- 1 Sword of the Animist | ramp | A repeatable ramp piece that also supports the creature plan.
- 1 Thought Vessel | ramp | Low-cost artifact ramp.
- 1 Commander's Sphere | ramp | Mana acceleration with later card utility.
- 1 White Lotus Tile | ramp | Artifact ramp for the deck's early turns.
- 1 Springleaf Drum | ramp | Cheap ramp that works alongside the creature-heavy plan.
- 1 Relic of Legends | ramp | Artifact ramp that supports the legendary commander.
- 1 Adventurer's Airship | draw | A draw artifact that helps keep cards flowing.
- 1 Buster Sword | draw | Equipment-based card draw for the creature plan.
- 1 Crown of Gondor | draw | A draw equipment that rewards keeping creatures in play.
- 1 Diary of Dreams | draw | Artifact card draw for longer games.
- 1 Feather of Flight | draw | A draw aura that supports a creature-focused board.
- 1 Idol of Oblivion | draw | Low-cost artifact card draw.
- 1 Lembas | draw | An artifact draw piece with useful blink potential.
- 1 Mask of Memory | draw | Equipment-based card draw that complements combat.
- 1 Mirror of Galadriel | draw | A draw artifact for maintaining resources.
- 1 Skullclamp | draw | An efficient equipment draw engine.
- 1 Tome of Legends | draw | A draw artifact that works with the commander.
- 1 Crib Swap | removal | Flexible creature removal.
- 1 Destroy Evil | removal | Efficient removal for key opposing permanents.
- 1 Dispatch | removal | Low-cost targeted removal.
- 1 Generous Gift | removal | Broad permanent removal.
- 1 Get Lost | removal | Versatile targeted removal.
- 1 March of Otherworldly Light | removal | Flexible removal that scales to the target.
- 1 Make Your Move | removal | Instant-speed targeted removal.
- 1 Oblation | removal | Removal that can answer troublesome permanents.
- 1 Swords to Plowshares | removal | Efficient creature removal.
- 1 Lightning Greaves | interaction | Protection for Gilraen and important creatures.
- 1 Swiftfoot Boots | interaction | Protects key creatures while supporting immediate use.
- 1 Darksteel Plate | interaction | Durable protection for a central creature.
- 1 Champion's Helm | interaction | Protective equipment for the commander and legends.
- 1 Clever Concealment | interaction | A defensive interaction spell for preserving the board.
- 1 Reprieve | interaction | Flexible stack interaction that also replaces itself.
- 1 Slip On the Ring | interaction | Protective interaction that fits the blink plan.
- 1 Unbreakable Formation | interaction | Board-protection interaction for creature-heavy games.
- 1 Flickerwisp | synergy | A blink-focused creature that works naturally with Gilraen.
- 1 Personify | synergy | A synergy spell supporting the deck's blink-focused plan.
- 1 Ennis, Debate Moderator | synergy | A synergy creature for the deck's central plan.
- 1 Jocasta, Automaton Avenger | synergy | A synergy creature that adds to the board presence.
- 1 Sheltered by Ghosts | synergy | A supporting aura for protecting a key creature.
- 1 Gift of Immortality | synergy | A resilience piece for creatures Gilraen wants to reuse.
- 1 Together Forever | synergy | A creature-supporting enchantment for the blink shell.
- 1 Angel of Serenity | wincon | A finisher and primary closing threat.
- 1 Austere Command | wipe | Flexible board wipe for resetting unfavorable boards.
- 1 Fumigate | wipe | A board wipe that helps stabilize against creature boards.
- 1 Vanquish the Horde | wipe | An efficient creature-board reset.
- 1 Bastion Protector | threat | A creature threat that also supports the commander.
- 1 Bronze Guardian | threat | A sturdy artifact creature that adds pressure.
- 1 Champions of Minas Tirith | threat | A creature threat that contributes to the attacking board.
- 1 Exemplar of Light | threat | An Angel threat for applying creature pressure.
- 1 Faramir, Field Commander | threat | A Human creature threat for building the board.
- 1 Frontline Medic | threat | A creature threat that supports combat development.
- 1 Giada, Font of Hope | threat | An Angel threat that supports the deck's creature base.
- 1 Inspiring Overseer | threat | A flying creature threat for pressuring opponents.
- 1 Kataki, War's Wage | threat | A disruptive creature threat.
- 1 Palace Jailer | threat | A creature threat that contributes to board control.
- 1 Puresteel Paladin | threat | A creature threat that complements the equipment package.
- 1 Wall of Omens | threat | An early creature that can be blinked and helps establish the board.
- 1 Westfold Rider | threat | A Human creature threat for the attacking plan.
- 1 Zack Fair | threat | A legendary creature threat that adds combat pressure.

</details>

### 6. Modern tempo, tournament

Format: Modern. Theme: tempo. Pool: any_card. Shortlist: 167 names.

Grade: typical, score 0.50, model 20260923T202806Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $354.72 to buy, $354.72 the whole deck.

**Summary:** This blue-red tempo deck establishes pressure with Ragavan, Nimble Pilferer, Ledger Shredder, and Faerie Mastermind, then protects that pressure with efficient interaction and removal. Preordain and Mishra's Bauble keep the deck moving, while Temporal Mastery and Temporal Trespass reinforce its spell-focused plan. It gives up broader late-game threats and sweeping answers in exchange for a fast, focused game built around trading efficiently and staying ahead.

The quality model grades this deck typical against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade. The deck runs its spells as playsets, and that lowers the grade.

- PLAN plan_coherent=partly: The core of cheap threats, cantrips, burn and counters is a single coherent tempo plan, but the four copies of Temporal Mastery and Temporal Trespass are seven- and eleven-mana sorceries with almost no delve or miracle support in a deck that wants to win by turn five, so four slots pull against the plan.
- PLAN theme_fit=partly: It is a blue-red Modern tempo deck with the right shell of cheap creatures and interaction, but the request named a Delver deck and the list contains no Delver of Secrets or comparable one-drop spell-payoff, instead adding Time Walk effects nobody asked for.
- PLAN useful_as_built=partly: The 24-land dual-heavy mana base, curve, and clock are all functional and the deck can be picked up and played, but four uncastable top-end sorceries mean roughly a tenth of the deck is dead weight in games it is trying to win quickly.
- PLAN summary_honest=partly: It accurately lists the threats, cantrips and interaction and admits the deck gives up late-game power, but calling Temporal Mastery and Temporal Trespass a reinforcement of the "spell-focused plan" dresses up two cards the deck cannot realistically cast.
- [INFO] `curve_summary`: average mana value 2.17 over 36 nonland cards

<details><summary>The deck list</summary>

- 4 Steam Vents | land | Provides an untapped blue-red land base for the deck's early plays.
- 4 Spirebluff Canal | land | Supplies efficient blue-red mana in the opening turns.
- 4 Riverglide Pathway // Lavaglide Pathway | land | Flexible color access without entering tapped.
- 4 Shivan Reef | land | Keeps both colors available for reactive turns.
- 4 Island | land | Provides dependable blue mana for the deck's blue-heavy spells.
- 2 Mountain | land | Provides dependable red mana for removal and threats.
- 2 Fiery Islet | land | Adds blue-red access while supporting the low land curve.
- 4 Ragavan, Nimble Pilferer | threat | An early threat that helps establish pressure.
- 4 Ledger Shredder | threat | A cheap threat that fits the deck's dense spell plan.
- 4 Faerie Mastermind | threat | A flash threat that supports a reactive tempo posture.
- 4 Preordain | draw | Smooths early draws and helps find the needed threat or answer.
- 2 Mishra's Bauble | draw | Provides low-cost card flow for a spell-dense shell.
- 4 Counterspell | interaction | Core permission for protecting pressure and stopping opposing plans.
- 2 Spell Pierce | interaction | Cheap interaction that helps the deck stay ahead in early exchanges.
- 4 Lightning Bolt | removal | Efficient removal that clears blockers and preserves tempo.
- 4 Galvanic Discharge | removal | Additional efficient removal for creatures that challenge the attack plan.
- 2 Temporal Mastery | synergy | Supports the deck's temporal spell package.
- 2 Temporal Trespass | synergy | Completes the temporal synergy package for the deck's spell-focused plan.

</details>

### 7. Modern burn, casual

Format: Modern. Theme: burn. Pool: any_card. Shortlist: 290 names.

Grade: baseline, score 0.40, model 20260923T202806Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $28.34 to buy, $28.34 the whole deck.

**Summary:** This mono-red burn deck applies early pressure with efficient removal and spell-focused synergy creatures, then keeps the cards coming to sustain its assault. Its threats give it a board-based route to victory when direct burn alone is not enough, with Hazoret the Fervent chief among them as a closing threat. The deck gives up flexibility against specialized strategies in exchange for a focused, consistent red mana base and a direct aggressive game plan.

The quality model grades this deck at the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade. Many of the cards appear in no top list, and that lowers the grade.

- PLAN plan_coherent=no: The pieces fight each other: Eidolon of the Great Revel and Thermo-Alchemist want a deck stuffed with cheap spells but there are only six burn spells, twelve clunky four-drop creatures drag the curve away from a burn clock, and four Hazoret the Fervent (which needs an empty hand) sit next to Browbeat and Risk Factor, which refill the hand and even hand the opponent the choice.
- PLAN theme_fit=partly: It is mono-red and Modern-legal as asked, but a burn deck with only four Lightning Bolt and two Lava Dart as actual burn, and sixteen four-mana creatures on top, is really a clunky red midrange deck wearing a burn shell.
- PLAN useful_as_built=partly: With 24 Mountains and a castable curve it functions and can win by attacking with Hazoret, Champion of the Path, and friends, but the mana is flooded for the spell count, Hazoret will often be stranded uncastable-to-attack, and the reach the plan advertises is not there.
- PLAN summary_honest=partly: It is candid about the creature-based route and the plain mana base, but calling six burn spells plus symmetrical Browbeat/Risk Factor an "assault" of "efficient removal" with "spell-focused synergy" oversells a list whose spell count cannot support Thermo-Alchemist, and naming Hazoret the chief closer hides that the deck's own draw spells and four-mana creatures work against its attack condition.
- [INFO] `curve_summary`: average mana value 2.83 over 36 nonland cards

<details><summary>The deck list</summary>

- 24 Mountain | land | Mono-red land base that reliably supplies red mana.
- 4 Lightning Bolt | removal | Efficient burn-oriented removal for clearing blockers or pressuring opponents.
- 2 Lava Dart | removal | Additional low-cost removal that supports the burn plan.
- 4 Browbeat | draw | Card-flow option that keeps the deck supplied with pressure.
- 2 Risk Factor | draw | Additional card-flow spell suited to an aggressive burn shell.
- 2 Chandra, Dressed to Kill | ramp | Provides the deck's ramp component while fitting the mono-red plan.
- 4 Eidolon of the Great Revel | synergy | A synergy piece that reinforces the deck's aggressive spell-focused strategy.
- 4 Thermo-Alchemist | synergy | A synergy creature that supports repeated burn-oriented pressure.
- 4 Hazoret the Fervent | threat | A durable top-end threat for closing games after early burn pressure.
- 4 Barret Wallace | threat | A threat that adds board pressure alongside the burn package.
- 4 Champion of the Path | threat | A threat slot that helps maintain a steady stream of attackers.
- 2 Keral Keep Disciples | threat | Additional threats that fill out the creature pressure plan.

</details>

### 8. Modern lifegain, FNM

Format: Modern. Theme: lifegain. Pool: any_card. Shortlist: 301 names.

Grade: bad, score 0.35, model 20260923T202806Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $898.46 to buy, $898.46 the whole deck.

**Summary:** This white-black lifegain deck establishes its early board with Guide of Souls and Ocelot Pride, then turns that foundation into pressure through Enduring Tenacity, Twinblade Paladin, and its larger threats. Solitude and Murderous Rider keep opposing threats from taking over while Lembas and Enduring Innocence help sustain the hand. It wins by building a resilient lifegain-based board and closing with its powerful creatures; the tradeoff is that several of its strongest threats ask the deck to reach the middle and later stages of the game.

The quality model grades this deck below the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade. The deck runs its spells as playsets, and that lowers the grade.

- PLAN plan_coherent=yes: Nearly every card feeds one lifegain-payoff plan — Guide of Souls/Ocelot Pride into Enduring Tenacity, Twinblade Paladin, Archangel of Thune, Sheoldred — with removal and Lembas as support, the only loose piece being the two Altar of the Pantheon.
- PLAN theme_fit=yes: It is a 60-card Modern white-black deck built entirely around gaining life and its payoffs, exactly the request.
- PLAN useful_as_built=yes: 24 lands with three dual lands supporting both colors, a curve topping at five, cheap creatures plus Twinblade Paladin/Sheoldred/Archangel as clocks, and removal in Solitude and Murderous Rider make it playable as it stands.
- PLAN summary_honest=yes: Every card and role the summary names is present in the list, and it openly notes the deck's reliance on reaching the mid and late game.
- [INFO] `curve_summary`: average mana value 3.11 over 36 nonland cards
- [INFO] `mana_pass`: the builder moved 1 card of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 4 Godless Shrine | land | A white-black dual land that supports both halves of the deck.
- 4 Caves of Koilos | land | An additional white-black source for early plays.
- 4 Shattered Sanctum | land | A white-black dual land that strengthens the mana base.
- 5 Plains | land | Reliable white sources for the deck's white cards.
- 7 Swamp | land | Reliable black sources for the deck's black cards.
- 2 Altar of the Pantheon | ramp | Provides the deck's dedicated mana acceleration.
- 4 Lembas | draw | A low-commitment draw card that fits the lifegain plan.
- 2 Enduring Innocence | draw | A durable draw engine for the creature-heavy build.
- 4 Solitude | removal | Premium removal that keeps opposing threats in check.
- 2 Murderous Rider // Swift End | removal | Flexible removal attached to a creature card.
- 4 Guide of Souls | synergy | A central early lifegain synergy piece.
- 4 Ocelot Pride | synergy | An efficient lifegain-focused synergy creature.
- 4 Enduring Tenacity | threat | A substantial black threat that rewards the deck's core plan.
- 4 Sheoldred, the Apocalypse | threat | A powerful midgame threat for the black half of the deck.
- 4 Twinblade Paladin | threat | A white threat that benefits from the lifegain theme.
- 2 Archangel of Thune | threat | A high-impact finishing threat for longer games.

</details>

### 9. Standard midrange, FNM

Format: Standard. Theme: midrange. Pool: any_card. Shortlist: 164 names.

Grade: bad, score 0.29, model 20260923T202806Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $176.50 to buy, $176.50 the whole deck.

**Summary:** This black-green midrange deck builds its mana around Overgrown Tomb and Wastewood Verge, uses Llanowar Elves to reach its heavier creatures, and keeps cards flowing through Darkstar Augur and Phyrexian Arena. Goldvein Hydra and Chomping Changeling apply steady pressure, while Vein Ripper, Massacre Wurm, and Vaultborn Tyrant provide heavier finishes. Bitter Triumph and Maelstrom Pulse clear the way, and Snakeskin Veil plus Not Dead After All help the creature plan survive resistance. It gives up a broad range of alternate angles in favor of committing to creatures, protection, and direct answers.

The quality model grades this deck below the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: The list is one plan: resilient creature threats backed by cheap protection (Snakeskin Veil, Not Dead After All), spot removal, and card advantage, with the curve topping at Vein Ripper/Massacre Wurm/Vaultborn Tyrant, and the only slightly off-plan pieces are the thin 2 Llanowar Elves.
- PLAN theme_fit=yes: It is a black-green creature midrange 60-card Standard deck with removal, card draw, and a fair top end, which is exactly the FNM-level request.
- PLAN useful_as_built=yes: 24 lands with 8 dual sources support a curve that peaks at six, there are enough threats and removal to close games, and the only stress point is the BBB six-drops, which the 17-odd black sources can reach.
- PLAN summary_honest=yes: Every card and role named in the summary is in the list, the trade-offs it admits (narrow, creature-committed) match the build, and it does not claim any package the deck lacks — only the ramp framing around 2 Llanowar Elves is mildly overstated.
- [INFO] `curve_summary`: average mana value 2.67 over 36 nonland cards
- [INFO] `mana_pass`: the builder moved 1 card of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 4 Overgrown Tomb | land | A black-green fixing land that supports the two-color mana base.
- 4 Wastewood Verge | land | A black-green fixing land that helps cast either half of the deck.
- 9 Swamp | land | A dependable black source for the deck's black-heavy spells.
- 7 Forest | land | A dependable green source for the deck's green spells.
- 4 Darkstar Augur | draw | A creature-based draw card that helps maintain resources while contributing to the board.
- 2 Phyrexian Arena | draw | A dedicated draw engine for longer midrange games.
- 2 Llanowar Elves | ramp | Early mana acceleration that helps deploy the deck's larger creatures sooner.
- 4 Bitter Triumph | removal | Low-cost removal that clears opposing problems efficiently.
- 2 Maelstrom Pulse | removal | Flexible removal for troublesome opposing permanents.
- 4 Snakeskin Veil | synergy | Protects key creatures and helps preserve the deck's board presence.
- 4 Not Dead After All | synergy | Keeps an important creature in the game through opposing removal.
- 4 Goldvein Hydra | threat | A scalable green creature that serves as an early or late-game threat.
- 4 Chomping Changeling | threat | A creature threat that adds pressure while fitting the deck's black-green plan.
- 2 Vein Ripper | threat | A powerful black creature for closing games after the board is stabilized.
- 2 Massacre Wurm | threat | A large creature threat that is especially strong when opponents have built a board.
- 2 Vaultborn Tyrant | threat | A top-end Dinosaur threat for games that go long.

</details>

### 10. Standard aggro, tournament

Format: Standard. Theme: aggro. Pool: any_card. Shortlist: 294 names.

Grade: bad, score 0.30, model 20260923T202806Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $95.42 to buy, $95.42 the whole deck.

**Summary:** This red-white aggro deck commits threats early, backs its attacks with removal and flexible interaction, and uses its draw cards to keep pressure on after the first exchange. Warleader's Call is chief among the cards tying the creature plan together. It aims to win by maintaining a fast board presence and clearing the way rather than by relying on a slow, high-end finish.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The cards pair in ways the top lists do not, and that lowers the grade. The curve sits high for the format, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core is a coherent creature-aggro plan with an anthem (Warleader's Call) and cheap removal, but four copies of a six-mana beast plus four four-drops pull against the fast-clock plan the summary describes.
- PLAN theme_fit=yes: It is a 60-card Standard deck in exactly red-white with an aggressive creature core, removal, and burn, which is what was asked for.
- PLAN useful_as_built=yes: 24 lands including three untapped-capable duals support the two-color costs, and 20 creatures plus an anthem and removal give it a clear, playable route to damage.
- PLAN summary_honest=partly: Removal, the two draw pieces and Warleader's Call are all really present, but the claim that it wins without 'relying on a slow, high-end finish' sits oddly with four copies of a {4}{R}{R} beast, and it does openly report the quality model's criticisms.
- [INFO] `curve_summary`: average mana value 3.17 over 36 nonland cards
- [INFO] `basics_added`: the list was 4 cards short, so the builder added 4 basic lands

<details><summary>The deck list</summary>

- 4 Sacred Foundry | land | A red-white source that keeps both colors available for the aggressive plan.
- 4 Inspiring Vantage | land | A red-white source supporting early plays in either color.
- 4 Sundown Pass | land | A red-white source that rounds out the dual-land base.
- 6 Plains | land | Reliable white mana for the deck's white spells.
- 6 Mountain | land | Reliable red mana for the deck's red spells.
- 4 Fugitive Codebreaker | draw | An aggressive red card that also occupies a draw slot.
- 2 Caretaker's Talent | draw | A white card included to provide additional cards during longer games.
- 4 Boros Charm | interaction | A red-white instant that provides flexible interaction.
- 2 Sheltered by Ghosts | interaction | A white interaction card that helps protect the deck's board position.
- 4 Case of the Gateway Express | removal | A white removal option that fits the deck's colors.
- 4 Emeritus of Truce // Swords to Plowshares | removal | A white removal card for efficiently answering opposing threats.
- 4 Warleader's Call | synergy | A red-white synergy piece for an aggressive creature strategy.
- 4 Bedhead Beastie | threat | A low-end red threat intended to start applying pressure early.
- 4 Dragonback Lancer | threat | A red-white threat that continues the deck's creature pressure.
- 4 Redcap Gutter-Dweller | threat | A red threat that fills out the attacking creature base.

</details>

### 11. Commander with a locked card

Format: Commander. Theme: sacrifice. Pool: any_card. Shortlist: 299 names.

Commander: Karlov of the Ghost Council.

Grade: baseline, score 0.05, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $783.34 to buy, $783.34 the whole deck.

**Summary:** This is an Orzhov sacrifice deck built to turn a broad creature board into steady pressure, with Karlov of the Ghost Council as a growing centerpiece. Sacrifice outlets, death-focused payoffs, and creature-based value keep the board working through exchanges, while high-impact threats and dedicated closing cards end drawn-out games. It gives up some speed to retain a deep creature package and several expensive top-end plays.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. Few lands enter tapped, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every nonland card is a sacrifice outlet, a death-trigger payoff, an expendable body, or a drain effect, and the drain/lifegain payoffs (Zulaport, Bastion, Elas il-Kor, Vindictive Vampire, Ayli) directly feed Karlov's lifegain counters, so the pieces pull in one direction.
- PLAN theme_fit=yes: It is exactly a Commander sacrifice deck under Karlov of the Ghost Council in Orzhov colors, as requested.
- PLAN useful_as_built=yes: 36 lands with strong dual/untapped fixing, cheap ramp, a low curve with a handful of top-end bombs, and multiple redundant win routes (drain triggers, Grave Pact effects, Razaketh, combat) make it immediately playable.
- PLAN summary_honest=yes: The summary's claims — sacrifice outlets, death payoffs, creature value, expensive top-end, and a slower clock — all match what the list actually contains, with no strategy named that the cards do not support.
- [INFO] `curve_summary`: average mana value 3.22 over 63 nonland cards
- [INFO] `mana_pass`: the builder moved 2 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 1 Brightclimb Pathway // Grimclimb Pathway | land | Flexible color support for the mana base.
- 1 Caves of Koilos | land | Flexible color support for the mana base.
- 1 City of Brass | land | Flexible color support for the mana base.
- 1 Command Tower | land | Reliable color support for the mana base.
- 1 Concealed Courtyard | land | Flexible color support for the mana base.
- 1 Eclipsed Steppe | land | Flexible color support for the mana base.
- 1 Godless Shrine | land | Flexible color support for the mana base.
- 1 Isolated Chapel | land | Flexible color support for the mana base.
- 1 Mana Confluence | land | Reliable color support for the mana base.
- 1 Scrubland | land | Flexible color support for the mana base.
- 1 Shattered Sanctum | land | Flexible color support for the mana base.
- 1 Vault of Champions | land | Flexible color support for the mana base.
- 10 Plains | land | Provides a dependable white mana foundation.
- 14 Swamp | land | Provides a dependable black mana foundation.
- 1 Sol Ring | ramp | Required early mana acceleration.
- 1 Ashnod's Altar | ramp | Turns sacrifice fodder into mana for the core plan.
- 1 Culling the Weak | ramp | Efficient burst mana for sacrifice turns.
- 1 Deadly Dispute | ramp | Supports sacrifice turns while advancing mana.
- 1 Pawn of Ulamog | ramp | Provides mana-oriented support for creature sacrifices.
- 1 Phyrexian Altar | ramp | Converts creature sacrifices into flexible mana.
- 1 Pitiless Plunderer | ramp | Rewards the deck's sacrifice-focused game plan with mana support.
- 1 Priest of Forgotten Gods | ramp | A creature-based mana piece that fits the sacrifice plan.
- 1 Skullport Merchant | ramp | A sacrifice-friendly source of mana support.
- 1 Warren Soultrader | ramp | Provides mana support within the creature-sacrifice shell.
- 1 Baron Bertram Graywater | draw | Creature-based card advantage for the sacrifice shell.
- 1 Bushmeat Poacher | draw | Provides card advantage in a creature-focused plan.
- 1 Corrupted Conviction | draw | Efficient card advantage that uses expendable creatures.
- 1 Disciple of Bolas | draw | Converts a creature into a substantial value play.
- 1 Ecstatic Awakener // Awoken Demon | draw | A sacrifice-friendly source of card advantage.
- 1 Jerren, Corrupted Bishop // Ormendahl, the Corrupter | draw | Creature-based card advantage for the deck's plan.
- 1 Lord Skitter's Butcher | draw | Adds card advantage from a creature slot.
- 1 Smothering Abomination | draw | A powerful card-advantage piece for sacrifice turns.
- 1 Tevesh Szat, Doom of Fools | draw | Provides sustained card advantage and board presence.
- 1 Vampiric Rites | draw | Turns expendable creatures into ongoing card advantage.
- 1 Village Rites | draw | Low-cost card advantage when a creature is sacrificed.
- 1 Cartel Aristocrat | interaction | Provides an interactive creature sacrifice outlet.
- 1 Dark Privilege | interaction | Protective interaction that works with sacrifice resources.
- 1 Fanatical Devotion | interaction | Protective interaction supported by expendable creatures.
- 1 Flare of Fortitude | interaction | A protection-focused interaction slot.
- 1 Gift of Doom | interaction | Protective interaction for a key permanent.
- 1 Promise of Tomorrow | interaction | Interaction that supports the creature-focused plan.
- 1 Rescue from the Underworld | interaction | Reactive support for an important creature.
- 1 Spirit Bonds | interaction | Adds flexible interaction to the creature shell.
- 1 Attrition | removal | Sacrifice-based creature removal for repeatable board control.
- 1 Ayli, Eternal Pilgrim | removal | A creature-based removal option that fits Karlov's colors.
- 1 Blasting Station | removal | Provides sacrifice-based removal and reach.
- 1 Bone Shards | removal | Efficient removal that can use a sacrifice as its cost.
- 1 Dictate of Erebos | removal | Turns creature sacrifices into broad removal pressure.
- 1 Eaten Alive | removal | Efficient removal compatible with expendable creatures.
- 1 Grave Pact | removal | Makes creature sacrifices a major removal engine.
- 1 Teysa, Orzhov Scion | removal | Creature-based removal that belongs in the sacrifice shell.
- 1 Yawgmoth, Thran Physician | removal | A repeatable creature-based removal piece.
- 1 Austere Command | wipe | Flexible reset button for difficult board states.
- 1 The Meathook Massacre | wipe | A sweeping effect that also serves as a closing threat.
- 1 Toxic Deluge | wipe | Efficient board reset against opposing creatures.
- 1 Bartolomé del Presidio | synergy | A low-cost sacrifice synergy piece.
- 1 Bastion of Remembrance | synergy | A sacrifice payoff that helps convert creature losses into pressure.
- 1 Carrion Feeder | synergy | A cheap, repeatable sacrifice outlet.
- 1 Elas il-Kor, Sadistic Pilgrim | synergy | A central sacrifice payoff and finisher.
- 1 Fleshtaker | synergy | A sacrifice-focused payoff creature.
- 1 Viscera Seer | synergy | A low-cost, repeatable sacrifice outlet.
- 1 Zulaport Cutthroat | synergy | A sacrifice payoff that can finish weakened opponents.
- 1 Abhorrent Overlord | threat | A high-impact finishing threat for longer games.
- 1 Basri's Lieutenant | threat | A resilient creature threat for the board-focused plan.
- 1 Demon of Catastrophes | threat | A sacrifice-compatible finishing threat.
- 1 Demon of Fate's Design | threat | A substantial finishing threat in the deck's colors.
- 1 Felisa, Fang of Silverquill | threat | A creature threat that supports the counter-focused commander.
- 1 Liesa, Forgotten Archangel | threat | A high-impact creature threat for contested games.
- 1 Mondrak, Glory Dominus | threat | A premium board-building threat.
- 1 Ratadrabik of Urborg | threat | A value-oriented legendary creature threat.
- 1 Requiem Angel | threat | A durable top-end creature threat.
- 1 Sidisi, Undead Vizier | threat | A versatile creature threat for the sacrifice shell.
- 1 Venerated Stormsinger | threat | A finishing creature threat.
- 1 Vindictive Vampire | threat | A sacrifice payoff that also pressures life totals.
- 1 Razaketh, the Foulblooded | wincon | A dedicated late-game finishing card.
- 1 Relic Vial | wincon | A dedicated closing piece for the creature plan.
- 1 Sephiroth, Fabled SOLDIER // Sephiroth, One-Winged Angel | wincon | A dedicated high-impact finisher.

</details>

### 12. Commander on a budget

Format: Commander. Theme: tokens. Pool: any_card. Shortlist: 298 names.

Commander: Adeline, Resplendent Cathar.

Grade: bad, score 0.01, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $92.82 to buy, $92.82 the whole deck.

**Summary:** Adeline, Resplendent Cathar drives a wide white token strategy, turning attacks into an expanding Human force and using the deck’s token-focused support to make that board matter. The deck develops with plentiful mana and card flow, keeps opposing boards in check with targeted answers and resets, then closes through combat pressure or dedicated finishing cards such as Halo Fountain, Luck Bobblehead, and Sword of Body and Mind. It favors steady board development over explosive early-game acceleration, so it is most effective when it can keep creatures on the table and attack repeatedly.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The deck makes less mana on turn four than the norm, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every slot supports the go-wide white token plan — token makers, anthems and Divine Visitation/Intangible Virtue payoffs, token-fueled draw like Idol of Oblivion and Dawn of Hope, and Halo Fountain/Hour of Reckoning to convert a wide board — with only a few weak pseudo-ramp artifacts (Coin of Mastery, Currency Converter) sitting off-plan without contradicting it.
- PLAN theme_fit=yes: It is exactly what was asked: a mono-white Commander deck headed by Adeline, Resplendent Cathar built around creating and exploiting tokens, with a budget-conscious card pool and no combos or fast mana that would push it past a casual bracket-2 ceiling.
- PLAN useful_as_built=yes: Thirty-seven lands in mono-white with a curve topping out reasonably, plenty of cheap token producers, seven removal spells, three sweepers and multiple ways to close (combat, Elspeth, Halo Fountain) make this immediately playable at a table.
- PLAN summary_honest=partly: The token engine, removal suite, wipes and Halo Fountain finish are all genuinely present, but calling the deck's mana "plentiful" clashes with the treasure-trinket ramp package (and the model's own turn-four mana note), and billing Luck Bobblehead and Sword of Body and Mind as "dedicated finishing cards" overstates two incidental value pieces.
- [INFO] `curve_summary`: average mana value 3.55 over 62 nonland cards

<details><summary>The deck list</summary>

- 30 Plains | land | Provides dependable white mana.
- 1 Ancient Den | land | Adds to the mana base.
- 1 Castle Ardenvale | land | Adds to the mana base.
- 1 Command Tower | land | Adds to the mana base.
- 1 Eiganjo, Seat of the Empire | land | Adds to the mana base.
- 1 Exotic Orchard | land | Adds to the mana base.
- 1 Secluded Courtyard | land | Adds to the mana base.
- 1 Unclaimed Territory | land | Adds to the mana base.
- 1 Coin of Mastery | ramp | Provides mana acceleration.
- 1 Collector's Vault | ramp | Provides mana acceleration.
- 1 Currency Converter | ramp | Provides mana acceleration.
- 1 Druidic Satchel | ramp | Provides mana acceleration.
- 1 Goldvein Pick | ramp | Provides mana acceleration.
- 1 Karn, Living Legacy | ramp | Provides mana acceleration.
- 1 Keeper of the Accord | ramp | Provides mana acceleration.
- 1 Monologue Tax | ramp | Provides mana acceleration.
- 1 The Restoration of Eiganjo // Architect of Restoration | ramp | Provides mana acceleration.
- 1 Thousand Moons Smithy // Barracks of the Thousand | ramp | Provides mana acceleration.
- 1 Angelic Sell-Sword | draw | Provides card draw.
- 1 Bygone Bishop | draw | Provides card draw.
- 1 Court of Grace | draw | Provides card draw.
- 1 Dawn of Hope | draw | Provides card draw.
- 1 Faramir, Field Commander | draw | Provides card draw.
- 1 Glimmer Seeker | draw | Provides card draw.
- 1 Idol of Oblivion | draw | Provides card draw.
- 1 Platoon Dispenser | draw | Provides card draw.
- 1 Staff of the Storyteller | draw | Provides card draw.
- 1 Wedding Announcement // Wedding Festivity | draw | Provides card draw.
- 1 Wojek Investigator | draw | Provides card draw.
- 1 Banishing Slash | removal | Answers opposing permanents.
- 1 Generous Gift | removal | Answers opposing permanents.
- 1 Hanged Executioner | removal | Answers opposing permanents.
- 1 Kellan's Lightblades | removal | Answers opposing permanents.
- 1 Righteous Confluence | removal | Answers opposing permanents.
- 1 Skyclave Apparition | removal | Answers opposing permanents.
- 1 Stroke of Midnight | removal | Answers opposing permanents.
- 1 Lena, Selfless Champion | interaction | Provides interaction.
- 1 Mage's Attendant | interaction | Provides interaction.
- 1 Rootborn Defenses | interaction | Provides interaction.
- 1 Spirit Bonds | interaction | Provides interaction.
- 1 Squad Commander | interaction | Provides interaction.
- 1 Teyo, the Shieldmage | interaction | Provides interaction.
- 1 Elspeth, Sun's Champion | wipe | Resets crowded boards.
- 1 Hour of Reckoning | wipe | Resets crowded boards.
- 1 Martial Coup | wipe | Resets crowded boards.
- 1 Divine Visitation | synergy | Supports the token plan.
- 1 Felidar Retreat | synergy | Supports the token plan.
- 1 Horn of Gondor | synergy | Supports the token plan.
- 1 Intangible Virtue | synergy | Supports the token plan.
- 1 Oketra's Monument | synergy | Supports the token plan.
- 1 Rosie Cotton of South Lane | synergy | Supports the token plan.
- 1 Siege Veteran | synergy | Supports the token plan.
- 1 Skrelv's Hive | synergy | Supports the token plan.
- 1 Twilight Drover | synergy | Supports the token plan.
- 1 Worthy Knight | synergy | Supports the token plan.
- 1 Archon of Sun's Grace | threat | Adds a token-focused threat.
- 1 Attended Healer | threat | Adds a token-focused threat.
- 1 Emeria Angel | threat | Adds a token-focused threat.
- 1 God-Eternal Oketra | threat | Adds a token-focused threat.
- 1 Gwaihir, Greatest of the Eagles | threat | Provides a finishing threat.
- 1 Hero of Bladehold | threat | Adds a token-focused threat.
- 1 Nahiri, the Lithomancer | threat | Provides a finishing threat.
- 1 Requiem Angel | threat | Provides a finishing threat.
- 1 Silverwing Squadron | threat | Adds a token-focused threat.
- 1 Threefold Thunderhulk | threat | Adds a token-focused threat.
- 1 Warren Warleader | threat | Adds a token-focused threat.
- 1 Wingmantle Chaplain | threat | Adds a token-focused threat.
- 1 Halo Fountain | wincon | Provides a dedicated finishing route.
- 1 Luck Bobblehead | wincon | Provides a dedicated finishing route.
- 1 Sword of Body and Mind | wincon | Provides a dedicated finishing route.

</details>

### 13. owned first, and the commander is not owned

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 230 names.

Commander: Karlov of the Ghost Council.

Grade: baseline, score 0.06, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $180.37 the whole deck.

**Summary:** Karlov of the Ghost Council is the focal point of a white-black lifegain deck that develops mana, gains cards through a mix of creatures and equipment, and turns steady life gains into a dangerous commander. Protective equipment and instant-speed safeguards help Karlov stay active while targeted answers clear the way. The deck closes through Karlov's accumulated pressure, evasive Angel threats, and finishers such as Frodo, Sauron's Bane, Lyra Dawnbringer, and Grave Venerations; in exchange, it leans heavily on creatures and artifacts remaining in play to sustain its momentum.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The land count sits above the norm of the format, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: There is a real spine — Karlov plus protective equipment, a few lifegain payoffs (Angel of Vitality, Rosie Cotton, Light of Promise, Shattered Angel, Lyra) and cheap interaction — but the actual density of repeatable lifegain triggers Karlov needs is thin, and a block of generic beaters (Canyon Crawler, Rabaroo Troop, Foggy Swamp Hunters, Rooftop Percher, Lo and Li) plus three board wipes in a creature-and-equipment deck pull against the stated plan.
- PLAN theme_fit=partly: It is a legal WB Commander deck led by Karlov at a reasonable bracket-3 power level, but the lifegain theme the request asked for is diluted by a large share of cards that neither gain life nor care about it.
- PLAN useful_as_built=yes: 35 lands plus eight to nine mana rocks, a sane curve, ten-plus pieces of interaction and multiple real threats mean the deck casts its spells and can close games through Karlov and the Angels.
- PLAN summary_honest=partly: The equipment, protection, removal and Angel claims all match the list, but calling a one-mana Frodo and an enchantment 'finishers' overstates them, the summary never mentions the three wraths that undercut its own creature-dependent board, and it soft-pedals how few actual lifegain engines back the 'steady life gains' claim.
- [WARN] `not_owned`: Karlov of the Ghost Council: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 2.98 over 63 nonland cards
- [INFO] `mana_pass`: the builder moved 5 cards of the mana base to bring the deck inside its power level
- [INFO] `basics_added`: the list was 1 card short, so the builder added 1 basic land

<details><summary>The deck list</summary>

- 15 Plains | land | Provides a reliable white source for the mana base.
- 10 Swamp | land | Provides a reliable black source for the mana base.
- 1 Command Tower | land | Provides flexible fixing for Karlov's colors.
- 1 City of Brass | land | Provides flexible fixing for the two-color mana base.
- 1 Exotic Orchard | land | Provides flexible fixing in multiplayer games.
- 1 Marsh Flats | land | Finds a needed basic land color.
- 1 Fabled Passage | land | Finds a needed basic land color.
- 1 Evolving Wilds | land | Finds a needed basic land color.
- 1 Path of Ancestry | land | Provides fixing for the commander-focused mana base.
- 1 Grand Coliseum | land | Adds another flexible color source.
- 1 Sol Ring | ramp | Efficiently accelerates the deck's early development.
- 1 Arcane Signet | ramp | Provides dependable color fixing and acceleration.
- 1 Fellwar Stone | ramp | Adds efficient mana acceleration.
- 1 Commander's Sphere | ramp | Provides fixing while supporting the deck's mana development.
- 1 Wayfarer's Bauble | ramp | Helps develop the mana base early.
- 1 Sword of the Animist | ramp | Turns creature attacks into mana development.
- 1 Thought Vessel | ramp | Provides colorless acceleration.
- 1 Chromatic Lantern | ramp | Supports broad mana fixing and acceleration.
- 1 Relic of Legends | ramp | Adds flexible mana production.
- 1 Path to Exile | ramp | Serves as ramp from the shortlist while retaining useful flexibility.
- 1 Buster Sword | draw | Provides card draw for the equipment portion of the deck.
- 1 Call of the Ring | draw | Provides repeatable card draw.
- 1 Idol of Oblivion | draw | Adds efficient card draw support.
- 1 Lembas | draw | Provides a compact source of card draw.
- 1 Mask of Memory | draw | Rewards the deck for connecting in combat.
- 1 Night's Whisper | draw | Provides efficient card draw.
- 1 Puresteel Paladin | draw | Supports equipment while supplying card draw.
- 1 Skullclamp | draw | Converts expendable creatures into cards.
- 1 Tome of Legends | draw | Works well with a commander-centered game plan.
- 1 Wall of Omens | draw | Provides early defense and a card.
- 1 Inspiring Overseer | draw | Adds card draw on a lifegain-friendly creature.
- 1 Bitter Triumph | removal | Provides flexible spot removal.
- 1 Crib Swap | removal | Answers an opposing creature cleanly.
- 1 Dispatch | removal | Offers efficient creature removal.
- 1 Fatal Push | removal | Provides low-cost spot removal.
- 1 Generous Gift | removal | Answers a wide range of opposing permanents.
- 1 Get Lost | removal | Provides versatile permanent removal.
- 1 Infernal Grasp | removal | Provides unconditional creature removal.
- 1 Swords to Plowshares | removal | Provides efficient creature removal.
- 1 Stroke of Midnight | removal | Answers troublesome nonland permanents.
- 1 Bastion Protector | interaction | Helps keep Karlov on the battlefield.
- 1 Champion's Helm | interaction | Protects the commander during combat-focused games.
- 1 Clever Concealment | interaction | Protects the board from opposing disruption.
- 1 Darksteel Plate | interaction | Provides durable protection for Karlov.
- 1 Lightning Greaves | interaction | Protects Karlov and supports immediate pressure.
- 1 Reprieve | interaction | Temporarily disrupts an opposing spell while replacing itself.
- 1 Swiftfoot Boots | interaction | Protects Karlov and supports attacking safely.
- 1 Unbreakable Formation | interaction | Protects the creature board through key turns.
- 1 Angel of Vitality | synergy | Supports the deck's lifegain-focused plan.
- 1 Compassionate Healer | synergy | Adds another lifegain synergy piece.
- 1 Kor Firewalker | synergy | Contributes to the deck's lifegain theme.
- 1 Light of Promise | synergy | Supports a growing Karlov game plan.
- 1 Rosie Cotton of South Lane | synergy | Rewards the deck for its lifegain-focused development.
- 1 Second Breakfast | synergy | Adds a compact lifegain synergy card.
- 1 White Mage's Staff | synergy | Supports the deck's lifegain and equipment themes.
- 1 Angel of Invention | threat | Provides a substantial creature threat.
- 1 Bill the Pony | threat | Adds a thematic creature threat.
- 1 Foggy Swamp Hunters | threat | Adds a creature body that pressures opponents.
- 1 Gollum, Silent Slinker // Meager Meal | threat | Provides a threatening black creature for the board.
- 1 Lo and Li, Twin Tutors | threat | Adds a legendary creature threat.
- 1 Minwu, White Mage | threat | Adds a white creature threat to the deck.
- 1 Rabaroo Troop | threat | Provides another creature threat for combat.
- 1 Reaping Willow | threat | Adds a durable creature threat.
- 1 Rooftop Percher | threat | Provides a further creature threat.
- 1 Shattered Angel | threat | Adds an evasive threat that fits the deck's theme.
- 1 Victory's Herald | threat | Provides a powerful evasive creature threat.
- 1 Canyon Crawler | threat | Adds another sizable creature threat.
- 1 Frodo, Sauron's Bane | wincon | Provides a dedicated finishing route.
- 1 Lyra Dawnbringer | wincon | Provides a powerful lifegain-oriented finishing threat.
- 1 Grave Venerations | wincon | Provides a dedicated late-game finishing route.
- 1 Austere Command | wipe | Provides a flexible board reset.
- 1 Fumigate | wipe | Resets creature-heavy boards while fitting the lifegain plan.
- 1 Vanquish the Horde | wipe | Provides an efficient creature-board reset.
- 1 Secluded Courtyard | land | the mana pass traded a basic land for this land, which makes two or more of the deck's colors
- 1 Unclaimed Territory | land | the mana pass traded a basic land for this land, which makes two or more of the deck's colors
- 1 Spire of Industry | land | the mana pass traded a basic land for this land, which makes two or more of the deck's colors

</details>

### 14. the user delegates the commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 299 names.

Commander: Atarka, World Render.

Grade: bad, score 0.04, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $872.62 to buy, $872.62 the whole deck.

**Summary:** Atarka, World Render leads a Gruul Dragon deck that ramps into a dense flight of threatening Dragons, then turns attacks into decisive double-strike combat steps. Tribal cost support helps deploy the expensive core, while draw, protection, removal, and sweepers keep the pressure sustained. The primary wins come from overwhelming combat with Atarka and the deck’s chief Dragon finishers, with the tradeoff that the deck remains creature-centered and can need its mana development to line up before its heaviest threats take over.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The land count sits above the norm of the format, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every card supports one line — ramp and tribal cost reducers deploy a dense Dragon curve that Atarka's double strike and payoffs like Shared Animosity, Crucible of Fire and Terror of the Peaks convert into damage, with only the odd Tibalt's Trickery sitting off-plan.
- PLAN theme_fit=yes: It is a Gruul Dragon tribal Commander deck with a builder-chosen Dragon commander and a bracket-3-appropriate power level (one Game Changer, no fast combo).
- PLAN useful_as_built=yes: The mana base is deep in both colors with strong fixing and the deck has many redundant Dragon threats and finishers, though 38 lands plus ten ramp pieces is flood-prone for a curve that tops out around six.
- PLAN summary_honest=yes: Ramp, tribal cost support, card draw, protection equipment, removal and the (mostly one-sided) sweepers it names are all present, and it openly flags the creature-dependence and mana-development risk, with 'sweepers' being only a slight stretch for Thundermaw Hellkite.
- [INFO] `curve_summary`: average mana value 3.48 over 61 nonland cards
- [INFO] `mana_pass`: the builder moved 5 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 1 Cinder Glade | land | Land slot supporting red and green mana.
- 1 City of Brass | land | Flexible fixing for the two-color mana base.
- 1 Command Tower | land | Reliable commander-color fixing.
- 1 Copperline Gorge | land | Red-green source for early development.
- 1 Cragcrown Pathway // Timbercrown Pathway | land | Flexible red or green source.
- 1 Exotic Orchard | land | Flexible fixing land.
- 1 Fabled Passage | land | Fixing land that supports the basic-land base.
- 1 Fire-Lit Thicket | land | Red-green fixing land.
- 1 Game Trail | land | Red-green source for the mana base.
- 1 Karplusan Forest | land | Untapped red-green source.
- 1 Mana Confluence | land | Flexible fixing for either commander color.
- 1 Mossfire Valley | land | Red-green fixing land.
- 12 Mountain | land | Basic red source for Dragon spells.
- 1 Path of Ancestry | land | Tribal fixing land.
- 1 Rockfall Vale | land | Red-green source for the mana base.
- 1 Rootbound Crag | land | Red-green source for the mana base.
- 1 Spire Garden | land | Untapped red-green source.
- 1 Stomping Ground | land | Untapped red-green source.
- 1 Taiga | land | Untapped red-green source.
- 1 Temple of the Dragon Queen | land | Dragon-focused fixing land.
- 1 Thornspire Verge | land | Red-green source for the mana base.
- 1 Wooded Foothills | land | Fetch land that fixes both commander colors.
- 5 Forest | land | Basic green source for ramp and support spells.
- 1 Mox Jasper | ramp | Low-cost Dragon-focused mana acceleration.
- 1 Carnelian Orb of Dragonkind | ramp | Early mana acceleration for the Dragon curve.
- 1 Dragon's Hoard | ramp | Dragon-focused mana acceleration.
- 1 Dragonstorm Globe | ramp | Fixing and acceleration for the deck's mana needs.
- 1 Jade Orb of Dragonkind | ramp | Early mana acceleration for Dragon spells.
- 1 Orb of Dragonkind | ramp | Early mana acceleration and fixing.
- 1 Patchwork Banner | ramp | Tribal mana acceleration.
- 1 Pillar of Origins | ramp | Dragon-focused mana acceleration.
- 1 Progenitor's Icon | ramp | Tribal mana acceleration.
- 1 Scaled Nurturer | ramp | Early creature-based mana acceleration.
- 1 Faithless Looting | draw | Efficient early card selection and card flow.
- 1 Demand Answers | draw | Low-cost card flow.
- 1 Thrill of Possibility | draw | Efficient card flow to keep the hand moving.
- 1 Sylvan Library | draw | Early ongoing card selection.
- 1 Skullclamp | draw | Low-cost card-draw support.
- 1 Garruk's Uprising | draw | Creature-focused card-draw support.
- 1 Guardian Project | draw | Ongoing creature-based card advantage.
- 1 Harmonize | draw | Straightforward refill for a creature-heavy deck.
- 1 Dragonborn Champion | draw | Dragon-themed source of card advantage.
- 1 Return of the Wildspeaker | draw | Flexible card-draw spell for the board state.
- 1 Heroic Intervention | interaction | Protects the Dragon board from opposing disruption.
- 1 Lightning Greaves | interaction | Protects a key Dragon or the commander.
- 1 Mithril Coat | interaction | Protection for a key creature.
- 1 Swiftfoot Boots | interaction | Protects a key Dragon or the commander.
- 1 Steely Resolve | interaction | Tribal protection for the creature core.
- 1 Veil of Summer | interaction | Efficient reactive protection.
- 1 Tibalt's Trickery | interaction | Broad reactive interaction.
- 1 The One Ring | interaction | Resilient interactive support piece.
- 1 Dragon's Fire | removal | Efficient Dragon-themed removal.
- 1 Draconic Roar | removal | Efficient Dragon-themed removal.
- 1 Glorybringer | removal | Dragon threat that also supplies removal.
- 1 Invasion of Tarkir // Defiant Thundermaw | removal | Dragon-themed removal with a Dragon follow-up.
- 1 Magmatic Hellkite | removal | Dragon-based removal option.
- 1 Obsidian Charmaw | removal | Dragon-based removal option.
- 1 Scourge of Valkas | removal | Dragon-themed removal that rewards the tribe.
- 1 Terror of the Peaks | removal | Dragon threat that supplies removal pressure.
- 1 Sarkhan's Rage | removal | Direct removal for problematic targets.
- 1 Breath Weapon | wipe | Efficient board-clearing option.
- 1 Draconic Intervention | wipe | Dragon-themed sweeper for crowded boards.
- 1 Thundermaw Hellkite | wipe | Dragon threat that also functions as a wipe effect.
- 1 Dragonlord's Servant | synergy | Tribal cost support for the Dragon curve.
- 1 Dragonspeaker Shaman | synergy | Tribal cost support for the Dragon curve.
- 1 Herald's Horn | synergy | Dragon tribal support piece.
- 1 Urza's Incubator | synergy | Tribal cost support for expensive Dragons.
- 1 Crucible of Fire | synergy | Dragon-focused combat support.
- 1 Shared Animosity | synergy | Rewards attacking with a Dragon force.
- 1 Sarkhan's Triumph | synergy | Finds an appropriate Dragon for the situation.
- 1 Backdraft Hellkite | threat | Dragon body that adds pressure to the air.
- 1 Dragon Broodmother | threat | High-impact Dragon threat for prolonged games.
- 1 Hellkite Charger | threat | Evasive Dragon threat that attacks well with Atarka.
- 1 Hellkite Courser | threat | Dragon threat that supports the commander plan.
- 1 Manaform Hellkite | threat | Evasive Dragon pressure at an efficient point on the curve.
- 1 Roaming Throne | threat | Tribal threat that reinforces Dragon-focused play.
- 1 Stormbreath Dragon | threat | Evasive Dragon threat for combat pressure.
- 1 Thunderbreak Regent | threat | Reliable Dragon threat for the air game.
- 1 Thrakkus the Butcher | threat | Dragon threat built for decisive combat turns.
- 1 Twinflame Tyrant | threat | Powerful evasive Dragon threat and finisher.
- 1 Terror of Mount Velus | threat | High-impact Dragon finisher for combat turns.
- 1 Scourge of the Throne | threat | Evasive Dragon threat that rewards attacking.
- 1 Lathliss, Dragon Queen | wincon | Dragon-centered finishing threat.
- 1 Wrathful Red Dragon | wincon | Dragon-based finishing threat.

</details>

### 15. delegated commander, owned first

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 230 names.

Commander: Astarion, the Decadent.

Grade: baseline, score 0.05, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: yes, for deck_size. Block findings: 0.

Cost: $0.00 to buy, $187.02 the whole deck.

**Summary:** A black-white lifegain deck built to turn incremental healing into an overwhelming board of creatures, resilient equipment carriers, and powerful flying finishers. It combines efficient removal, protective tools, and several board resets to maintain control while its life-total advantages become decisive pressure.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The land count sits above the norm of the format, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The lifegain core (Astarion, Rosie Cotton, Angel of Vitality, Light of Promise, Second Breakfast, Shattered Angel) is real, but it is diluted by a separate equipment package (Puresteel Paladin, Skullclamp, Champion's Helm, Buster Sword, Sword of the Animist) and by three sweepers that fight the creature-wide board the deck otherwise assembles.
- PLAN theme_fit=yes: The request was a bracket-2 lifegain Commander deck with the builder choosing the commander, and Astarion, the Decadent with a W/B lifegain shell in a low-power, precon-level list matches that ask.
- PLAN useful_as_built=yes: Mana is abundant and well-colored for a mostly-white curve with a light black splash, the curve is low, and there are clear finishers (Lyra, Victory's Herald, Astarion's counters) so the deck can be picked up and played, though 39 lands is more than needed.
- PLAN summary_honest=yes: It names removal, protection, equipment carriers, flying angels, and board resets, all of which appear in the list, and it even flags the high land count and precon-level power rather than hiding them.
- [WARN] `not_owned`: Astarion, the Decadent: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.10 over 60 nonland cards

<details><summary>The deck list</summary>

- 1 Angel of Invention | threat | Evasive lifelink threat that supports the life-gain plan.
- 1 Angel of Vitality | synergy | Rewards repeated life gain with a growing flying body.
- 1 Arcane Signet | ramp | Reliable color fixing and acceleration.
- 1 Austere Command | wipe | Flexible reset for problematic boards.
- 1 Bastion Protector | interaction | Helps keep the commander protected.
- 1 Bill the Pony | threat | Creature pressure that fits the combat plan.
- 1 Bitter Triumph | removal | Flexible answer to creatures and planeswalkers.
- 1 Boromir, Warden of the Tower | interaction | Disrupts opposing free spells and protects the board.
- 1 Buster Sword | draw | Equipment-based card advantage.
- 1 Canyon Crawler | threat | Creature threat for applying board pressure.
- 1 Champion's Helm | interaction | Protects and enhances the commander.
- 1 Clever Concealment | interaction | Protects permanents from sweepers and targeted disruption.
- 1 Command Tower | land | Untapped source for either deck color.
- 1 Compassionate Healer | synergy | Repeated life-gain support for the deck's central plan.
- 1 Darksteel Plate | interaction | Durable commander protection.
- 1 Dispatch | removal | Efficient creature interaction.
- 1 Exemplar of Light | draw | Life-gain payoff that also supplies cards.
- 1 Exotic Orchard | land | Flexible dual-color fixing.
- 1 Fabled Passage | land | Finds needed basics while fixing colors.
- 1 Fellwar Stone | ramp | Cheap fixing and acceleration.
- 1 Foggy Swamp Hunters | threat | Creature threat that contributes to combat pressure.
- 1 Frodo, Sauron's Bane | wincon | Mana sink that can close games.
- 1 Fumigate | wipe | Creature reset with a substantial life-gain payoff.
- 1 Generous Gift | removal | Answers any troublesome permanent.
- 1 Get Lost | removal | Versatile answer to key nonland permanents.
- 1 Giada, Font of Hope | ramp | Accelerates the Angel package and improves its bodies.
- 1 Gollum, Silent Slinker // Meager Meal | threat | Creature threat with useful additional value.
- 1 Grave Venerations | wincon | Life-total pressure that can finish opponents.
- 1 Hero in Training | draw | Creature-based card advantage.
- 1 Idol of Oblivion | draw | Low-cost recurring card draw.
- 1 Infernal Grasp | removal | Clean answer to opposing creatures.
- 1 Inspiring Overseer | draw | Flying body that replaces itself and gains life.
- 1 Instant Ramen | draw | Card advantage that fits the food and life-gain texture.
- 1 Kor Firewalker | synergy | Life-gain synergy on a resilient creature.
- 1 Lembas | draw | Card selection with life-gain utility.
- 1 Light of Promise | synergy | Turns life gain into a large creature threat.
- 1 Lightning Greaves | interaction | Efficient haste and protection for important creatures.
- 1 Lotho, Corrupt Shirriff | ramp | Produces mana while advancing the life-gain plan.
- 1 Lyra Dawnbringer | wincon | Powerful lifelinking flying finisher.
- 1 Marsh Flats | land | Fetch land that fixes both colors.
- 1 Mask of Memory | draw | Combat-based card filtering and advantage.
- 1 Minwu, White Mage | threat | Creature threat that supports sustained board presence.
- 1 Night Nurse, Healer of Heroes | synergy | Life-gain support that reinforces creature combat.
- 1 Night's Whisper | draw | Efficient two-card refill.
- 1 Oath of the Grey Host | ramp | Mana development with later board value.
- 18 Plains | land | Primary white mana base.
- 1 Prideful Feastling | synergy | Food-adjacent life-gain synergy.
- 1 Puresteel Paladin | draw | Equipment-focused card advantage.
- 1 Rabaroo Troop | threat | Creature threat for combat pressure.
- 1 Reaping Willow | threat | Large creature threat that benefits from the deck's plan.
- 1 Rosie Cotton of South Lane | synergy | Converts life gain into permanent creature growth.
- 1 Rooftop Percher | threat | Creature threat that helps pressure opponents.
- 1 Second Breakfast | synergy | Life-gain payoff and support spell.
- 1 Shattered Angel | threat | Flying threat that gains life over the course of the game.
- 1 Skullclamp | draw | Efficient creature-based card advantage.
- 1 Sneering Shadewriter | threat | Evasive Vampire finisher.
- 1 Sol Ring | ramp | Efficient colorless acceleration.
- 1 Sword of the Animist | ramp | Equipment that steadily develops mana.
- 1 Swords to Plowshares | removal | Efficient exile-based creature removal.
- 14 Swamp | land | Primary black mana base.
- 1 Terramorphic Expanse | land | Fetches required basic colors.
- 1 Thought Vessel | ramp | Colorless acceleration with hand-size utility.
- 1 Vanquish the Horde | wipe | Efficient creature-board reset.
- 1 Victory's Herald | threat | Flying finisher that grants combat lifelink.
- 1 Wayfarer's Bauble | ramp | Basic-land ramp and color fixing.
- 1 Well-Worn Spatula | synergy | Equipment synergy for the creature plan.
- 1 White Auracite | ramp | Artifact-based acceleration.
- 1 City of Brass | land | Untapped fixing for both deck colors.
- 1 Evolving Wilds | land | Basic-land fixing for either color.

</details>

### 16. a tight budget, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 231 names.

Commander: Denethor, Ruling Steward.

Grade: baseline, score 0.26, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $11.94 to buy, $39.82 the whole deck.

**Summary:** This is a patient Orzhov aristocrats deck that develops creatures and support pieces, then uses Denethor’s sacrifice ability to turn creature deaths into life swings while building an end-step Soldier force. The finishing cards give the deck ways to close once the board has been established, while the removal, protection, and board resets help it survive longer games. It gives up explosive speed for a library-first mana base and an incremental, board-centered plan.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core is coherently sacrifice-oriented (Ayli, Bartolomé, Woe Strider, Yahenni, Skullport Merchant, Erebos, Smothering Abomination feeding Denethor's lifegain/Soldier engine), but three unconditional board wipes and a heavy card-draw suite sit awkwardly beside a wide creature board, and the deck has almost no drain payoffs or repeatable token fodder to convert the deaths into damage.
- PLAN theme_fit=yes: It is an Orzhov Commander aristocrats deck under Denethor with sac outlets, death payoffs, and an interaction package at a low-power, precon-level build consistent with bracket 2 and a cheap library-first pool.
- PLAN useful_as_built=partly: It is playable and has plenty of lands and ramp, but the mana is skewed the wrong way — 16 Plains to 11 Swamps for a deck that is predominantly black with several double-black costs — and the win path is slow chip damage with no drain engine, so games can stall.
- PLAN summary_honest=partly: The description of Denethor turning deaths into life and end-step Soldiers is accurate and it honestly flags the slow mana, but 'the finishing cards give the deck ways to close' oversells the list, since the cards tagged wincon (Al Bhed Salvagers, Grave Venerations, Archfiend of Ifnir) are value pieces rather than a stated finish.
- [WARN] `not_owned`: Aron, Benalia's Ruin: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ayli, Eternal Pilgrim: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bartolomé del Presidio: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Baron Bertram Graywater: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Shadowheart, Dark Justiciar: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Skullport Merchant: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Woe Strider: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Yahenni, Undying Partisan: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Disciple of Bolas: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Elder Arthur Maxson: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Erebos, Bleak-Hearted: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ghoulcaller Gisa: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Lord Skitter's Butcher: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Old Flitterfang: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ruthless Lawbringer: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Sivriss, Nightmare Speaker: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Smothering Abomination: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Umbral Collar Zealot: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vampire Gourmand: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.02 over 62 nonland cards
- [INFO] `mana_pass`: the builder moved 2 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 1 Ahriman | draw | Provides a library-owned draw option.
- 1 Beetle-Headed Merchants | draw | Provides a library-owned draw option.
- 1 Circle of Power | draw | Provides a library-owned draw option.
- 1 Cirith Ungol Patrol | draw | Provides a library-owned draw option.
- 1 Exemplar of Light | draw | Provides a library-owned draw option.
- 1 Idol of Oblivion | draw | Provides a library-owned draw option.
- 1 Inspiring Overseer | draw | Provides a library-owned draw option.
- 1 June, Bounty Hunter | draw | Provides a library-owned draw option.
- 1 Mask of Memory | draw | Provides a library-owned draw option.
- 1 Skullclamp | draw | Provides a library-owned draw option that suits a creature-focused plan.
- 1 Wall of Omens | draw | Provides a library-owned draw option.
- 1 Bastion Protector | interaction | Helps protect the board plan through interaction.
- 1 Gift of Immortality | interaction | Supports the creature plan with interaction.
- 1 Reprieve | interaction | Provides a flexible library-owned interaction slot.
- 1 Swiftfoot Boots | interaction | Helps keep an important creature in play.
- 1 Take Up the Shield | interaction | Provides a low-cost interaction option.
- 1 Zack Fair | interaction | Provides a creature-based interaction option.
- 1 Bitter Triumph | removal | Provides a library-owned removal option.
- 1 Claim the Precious | removal | Provides a library-owned removal option.
- 1 Crib Swap | removal | Provides a library-owned removal option.
- 1 Generous Gift | removal | Provides broad library-owned removal.
- 1 Infernal Grasp | removal | Provides a library-owned removal option.
- 1 Stroke of Midnight | removal | Provides a library-owned removal option.
- 1 Swords to Plowshares | removal | Provides a library-owned removal option.
- 1 Austere Command | wipe | Provides a flexible library-owned board reset.
- 1 Dusk // Dawn | wipe | Provides a library-owned board reset.
- 1 Fumigate | wipe | Provides a library-owned board reset.
- 1 Arcane Signet | ramp | Provides efficient library-owned mana acceleration.
- 1 Astral Cornucopia | ramp | Provides a library-owned ramp option.
- 1 Chromatic Lantern | ramp | Provides library-owned mana acceleration and fixing.
- 1 Commander's Sphere | ramp | Provides a library-owned ramp option.
- 1 Deadly Dispute | ramp | Provides a library-owned ramp option for the sacrifice-focused plan.
- 1 Fellwar Stone | ramp | Provides efficient library-owned mana acceleration.
- 1 Inherited Envelope | ramp | Provides a library-owned ramp option.
- 1 Relic of Legends | ramp | Provides a library-owned ramp option.
- 1 Sol Ring | ramp | Provides efficient library-owned mana acceleration.
- 1 Thought Vessel | ramp | Provides a library-owned ramp option.
- 1 Aron, Benalia's Ruin | synergy | Fills a synergy slot for the creature-sacrifice plan.
- 1 Ayli, Eternal Pilgrim | synergy | Fills a synergy slot for the creature-sacrifice plan.
- 1 Bartolomé del Presidio | synergy | Fills a synergy slot for the creature-sacrifice plan.
- 1 Baron Bertram Graywater | synergy | Fills a synergy slot for the creature-sacrifice plan.
- 1 Gollum the Abandoned | synergy | Provides library-owned synergy for the deck's main plan.
- 1 Gollum, Patient Plotter | synergy | Provides library-owned synergy for the deck's main plan.
- 1 Shadowheart, Dark Justiciar | synergy | Fills a synergy slot for the creature-sacrifice plan.
- 1 Skullport Merchant | synergy | Fills a synergy slot for the creature-sacrifice plan.
- 1 Woe Strider | synergy | Fills a synergy slot for the creature-sacrifice plan.
- 1 Yahenni, Undying Partisan | synergy | Fills a synergy slot for the creature-sacrifice plan.
- 1 Al Bhed Salvagers | wincon | Serves as a library-owned finishing card.
- 1 Archfiend of Ifnir | wincon | Serves as a library-owned finishing card.
- 1 Grave Venerations | wincon | Serves as a library-owned finishing card.
- 1 Disciple of Bolas | threat | Adds a creature-based threat to the board plan.
- 1 Elder Arthur Maxson | threat | Adds a creature-based threat to the board plan.
- 1 Erebos, Bleak-Hearted | threat | Adds a resilient threat to the board plan.
- 1 Gaius van Baelsar | threat | Adds a library-owned creature threat.
- 1 Ghoulcaller Gisa | threat | Adds a creature-based threat to the board plan.
- 1 Lord Skitter's Butcher | threat | Adds a creature-based threat to the board plan.
- 1 Old Flitterfang | threat | Adds a creature-based threat to the board plan.
- 1 Ruthless Lawbringer | threat | Adds a creature-based threat to the board plan.
- 1 Sivriss, Nightmare Speaker | threat | Adds a creature-based threat to the board plan.
- 1 Smothering Abomination | threat | Adds a creature-based threat to the board plan.
- 1 Umbral Collar Zealot | threat | Adds a creature-based threat to the board plan.
- 1 Vampire Gourmand | threat | Adds a creature-based threat to the board plan.
- 16 Plains | land | Supplies reliable white mana.
- 11 Swamp | land | Supplies reliable black mana.
- 1 Command Tower | land | Provides efficient two-color fixing.
- 1 Exotic Orchard | land | Provides two-color fixing.
- 1 Evolving Wilds | land | Finds a needed basic color.
- 1 Fabled Passage | land | Finds a needed basic color.
- 1 Grand Coliseum | land | Provides two-color fixing.
- 1 Path of Ancestry | land | Provides two-color fixing for the creature-heavy plan.
- 1 Secluded Courtyard | land | Provides creature-focused color fixing.
- 1 Spire of Industry | land | Provides two-color fixing alongside the artifact package.
- 1 Terramorphic Expanse | land | Finds a needed basic color.
- 1 Unclaimed Territory | land | Provides creature-focused color fixing.

</details>

### 17. upgrade a precon, owned first

Format: Commander. Theme: goblins. Pool: owned_first. Shortlist: 205 names.

Commander: Zada, Hedron Grinder.

Grade: typical, score 0.47, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $34.39 to buy, $218.33 the whole deck.

**Summary:** Zada Goblin Storm builds a wide Goblin board and turns creature-targeted spells into broad bursts of cards, pressure, and momentum. The deck wins by converting that developed board into a decisive attack, with Great Train Heist and Collective Inferno chief among its closing cards, while Goblin-linked removal clears resistance. It gives up some consistency to retain the precon's varied Goblin package and remains most effective when it has both creatures and a stocked hand.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every card supports one plan — flood the board with Goblin tokens and multiply cheap cantrip/pump spells through Zada for cards and a lethal alpha strike — with the token makers, lords, sacrifice outlets and one-target spells all pulling the same direction.
- PLAN theme_fit=yes: It is a mono-red Commander deck built on the Goblin storm precon's commander with clear upgrades (Sol Ring, Skullclamp, Shared Animosity, Fable) while staying at an unoptimized, precon-plus bracket-3 level.
- PLAN useful_as_built=yes: 36 lands plus Skirk Prospector, Battle Hymn, Seething Song and Sol Ring support a low curve, and the deck has multiple win routes (combat pump, Impact Tremors/Boggart Shenanigans drain, Bombardment), so it plays fine as presented.
- PLAN summary_honest=yes: The summary accurately describes the Zada copy engine, names cards actually present (Great Train Heist, Collective Inferno), and openly concedes the consistency cost of keeping the precon's scattered Goblin package rather than overclaiming.
- [INFO] `curve_summary`: average mana value 2.62 over 61 nonland cards
- [WARN] `profile_off_band`: the interaction count is 1, and bracket 3 wants 4 to 12
- [INFO] `bracket_cut`: to hold bracket 3, the builder cut 1 card: Storm-Kiln Artist (the combo Storm-Kiln Artist + Haze of Rage), and added 1 basic land

<details><summary>The deck list</summary>

- 1 Ancestors' Aid | other | Retained as a precon spell for the deck's Goblin Storm theme.
- 1 Ancestral Anger | draw | Adds a low-cost draw spell to support spell-focused turns.
- 1 Arena of Glory | land | Retained as part of the precon mana base.
- 1 Battle Hymn | ramp | Retained as a Goblin-based burst-ramp spell.
- 1 Battle-Scarred Goblin | removal | Adds another Goblin removal piece.
- 1 Blasphemous Act | wipe | Retained as the precon's broad creature reset.
- 1 Boggart Shenanigans | removal | Retained as Goblin-linked removal.
- 1 Castle Embereth | land | Retained as part of the precon mana base.
- 1 Chaos Warp | removal | Retained as flexible removal.
- 1 Collective Inferno | wincon | Adds a dedicated finishing card.
- 1 Conspicuous Snoop | synergy | Retained as a Goblin synergy piece.
- 1 Crimson Wisps | draw | Retained as a card-draw spell for spell-heavy turns.
- 1 Cunning Maneuver | draw | Adds a draw spell suited to the deck's creature-focused plan.
- 1 Daring Discovery | other | Retained as a precon spell that preserves its original texture.
- 1 Den of the Bugbear | land | Retained as part of the precon mana base.
- 1 Dragon Fodder | synergy | Adds more Goblin-producing synergy.
- 1 Empty the Warrens | synergy | Retained as a Goblin Storm synergy card.
- 1 Expedite | draw | Retained as an inexpensive draw spell.
- 1 Fable of the Mirror-Breaker // Reflection of Kiki-Jiki | ramp | Adds a strong ramp card while remaining within the creature-and-spell theme.
- 1 Faithless Looting | draw | Retained as efficient card selection and draw.
- 1 Fists of Flame | draw | Retained as a draw spell for the deck's combat-oriented spell turns.
- 1 Forgotten Cave | land | Retained as part of the precon mana base.
- 1 Fountainport | land | Retained as part of the precon mana base.
- 1 Frontline Heroism | other | Retained as a precon card that preserves the original deck's character.
- 1 Gempalm Incinerator | removal | Retained as Goblin-linked removal.
- 1 General Kreat, the Boltbringer | synergy | Retained as a Goblin synergy card.
- 1 Glimpse the Impossible | ramp | Retained as a ramp spell.
- 1 Goblin Bombardment | removal | Adds efficient removal that works with the Goblin theme.
- 1 Goblin Bushwhacker | synergy | Retained as a Goblin synergy piece.
- 1 Goblin Chieftain | synergy | Retained as a Goblin synergy piece.
- 1 Goblin Dark-Dwellers | threat | Retained as a Goblin threat.
- 1 Goblin Lackey | synergy | Retained as a Goblin synergy piece.
- 1 Goblin Matron | synergy | Retained as a Goblin synergy piece.
- 1 Goblin Negotiation | removal | Retained as Goblin-themed removal.
- 1 Goblin Trashmaster | removal | Retained as a Goblin removal card.
- 1 Goblin Warchief | synergy | Retained as a Goblin synergy piece.
- 1 Grapeshot | removal | Retained as a spell-based removal card.
- 1 Great Train Heist | wincon | Retained as one of the deck's finishing cards.
- 1 Haze of Rage | other | Retained as a precon spell that supports the deck's storm theme.
- 1 Hidden Volcano | land | Retained as part of the precon mana base.
- 1 Howlsquad Heavy | ramp | Retained as a ramp creature.
- 1 Idol of Oblivion | draw | Retained as a draw artifact.
- 1 Impact Tremors | other | Retained as a precon card supporting the deck's creature plan.
- 1 Kher Keep | land | Retained as part of the precon mana base.
- 1 Krenko's Command | synergy | Retained as a Goblin synergy spell.
- 1 Krenko, Mob Boss | threat | Retained as a Goblin threat.
- 1 Lightning Bolt | removal | Adds efficient removal.
- 1 Mana Geyser | ramp | Retained as a high-output ramp spell.
- 27 Mountain | land | Provides a reliable untapped red mana base.
- 1 Mogg War Marshal | synergy | Retained as a Goblin synergy creature.
- 1 Pashalik Mons | removal | Adds another Goblin removal card.
- 1 Past in Flames | other | Retained as a precon spell for the deck's spell-focused theme.
- 1 Quest for the Goblin Lord | synergy | Retained as a Goblin synergy card.
- 1 Reliquary Tower | land | Retained as part of the precon mana base.
- 1 Ruby Medallion | other | Retained as a precon card for the deck's spell-focused plan.
- 1 Rundvelt Hordemaster | synergy | Adds a focused Goblin synergy piece.
- 1 Sazacap's Brew | draw | Retained as a draw spell.
- 1 Seething Song | ramp | Retained as burst ramp.
- 1 Shared Animosity | synergy | Adds a powerful Goblin synergy payoff.
- 1 Shinka, the Bloodsoaked Keep | land | Retained as part of the precon mana base.
- 1 Siege-Gang Commander | removal | Retained as Goblin-linked removal.
- 1 Siege-Gang Lieutenant | removal | Retained as Goblin-linked removal.
- 1 Skirk Prospector | ramp | Retained as a Goblin ramp piece.
- 1 Skullclamp | draw | Retained as an efficient draw card.
- 1 Sol Ring | ramp | Retained as efficient ramp.
- 1 Spinerock Knoll | land | Retained as part of the precon mana base.
- 1 Spreading Insurrection | other | Retained as a precon spell that preserves the original deck's character.
- 1 Swiftfoot Boots | interaction | Retained as a protection-oriented interaction card.
- 1 Throne of Eldraine | ramp | Retained as a ramp artifact.
- 1 Vandalblast | wipe | Retained as artifact-focused removal.
- 1 War Room | land | Retained as part of the precon mana base.
- 1 Wild Ride | other | Retained as a precon spell that preserves the original deck's character.
- 1 Witch's Mark | draw | Retained as a draw spell.

</details>

### 18. upgrade a precon, any card

Format: Commander. Theme: turtles. Pool: any_card. Shortlist: 365 names.

Commander: Heroes in a Half Shell.

Grade: bad, score 0.04, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $580.57 to buy, $580.57 the whole deck.

**Summary:** This deck keeps Turtle Power’s broad Mutant, Ninja, and Turtle creature core while giving its mana base a cleaner route to all five colors. Build a board of themed creatures, send them into combat, and let Heroes in a Half Shell turn successful hits into larger attackers and more cards. Raphael, the Muscle, Dimension X Pizzasaur, and Everything Pizza are chief among the cards that can close a game. The deck gives up some of the precon’s slower lands to improve early development, while retaining its varied character-driven creature package and big finishing turns.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The land count sits above the norm of the format, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The bulk of the list is a coherent themed creature-combat deck built around Heroes in a Half Shell's attack triggers with equipment, auras and pump, but three mass board wipes (Blasphemous Act, Vanquish the Horde, Wave Goodbye) work against its own wide creature board.
- PLAN theme_fit=yes: It is a five-color Commander deck led by Heroes in a Half Shell that keeps the precon's Turtle/Mutant/Ninja character cards while adding fixing lands and a few staples, which is exactly the requested modest precon upgrade.
- PLAN useful_as_built=yes: Thirty-one lands with heavy five-color fixing, Sol Ring, Arcane Signet, Chromatic Lantern and Cultivate support a mostly two-to-five-drop curve, with plenty of creature-based and named finishers to actually close a game.
- PLAN summary_honest=partly: The summary accurately describes the creature core, the mana-base rework and the named finishers, but it never mentions the three symmetrical board wipes that cut against the go-wide plan it advertises.
- [INFO] `curve_summary`: average mana value 3.46 over 59 nonland cards
- [INFO] `precon_cards_restored`: the deck was 3 cards short of the Turtle Power precon share, so the builder put 3 cards back
- [WARN] `profile_off_band`: the land count is 40, and bracket 3 wants 34 to 38
- [WARN] `profile_off_band`: the ramp count is 6, and bracket 3 wants 8 to 13
- [WARN] `profile_off_band`: the draw count is 4, and bracket 3 wants 8 to 14
- [INFO] `mana_pass`: the builder moved 3 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 1 Adarkar Wastes | land | Improves the five-color mana base.
- 1 Badlands | land | Improves the five-color mana base.
- 1 Bayou | land | Improves the five-color mana base.
- 1 Battlefield Forge | land | Improves the five-color mana base.
- 1 Big Apple, 3 a.m. | land | Retained from the Turtle Power precon mana base.
- 1 Blood Crypt | land | Improves the five-color mana base.
- 1 Breeding Pool | land | Improves the five-color mana base.
- 1 Brushland | land | Improves the five-color mana base.
- 1 Caves of Koilos | land | Improves the five-color mana base.
- 1 City of Brass | land | Retained from the Turtle Power precon mana base.
- 1 Command Tower | land | Retained from the Turtle Power precon mana base.
- 1 Escape Tunnel | land | Retained from the Turtle Power precon mana base.
- 1 Exotic Orchard | land | Retained from the Turtle Power precon mana base.
- 1 Fabled Passage | land | Retained from the Turtle Power precon mana base.
- 1 Godless Shrine | land | Improves the five-color mana base.
- 1 Grand Coliseum | land | Retained from the Turtle Power precon mana base.
- 1 Hallowed Fountain | land | Improves the five-color mana base.
- 1 Hinterland Harbor | land | Retained from the Turtle Power precon mana base.
- 1 Karplusan Forest | land | Improves the five-color mana base.
- 1 Spire Garden | land | the deck upgrades a precon, and this card is one the precon holds
- 1 Rootbound Crag | land | Retained from the Turtle Power precon mana base.
- 1 Smoldering Marsh | land | the deck upgrades a precon, and this card is one the precon holds
- 1 Undergrowth Stadium | land | Retained from the Turtle Power precon mana base.
- 1 Vernal Fen | land | Retained from the Turtle Power precon mana base.
- 1 Vibrant Cityscape | land | Retained from the Turtle Power precon mana base.
- 4 Forest | land | Provides a basic green source.
- 3 Island | land | Provides a basic blue source.
- 2 Mountain | land | Provides a basic red source.
- 4 Plains | land | Provides a basic white source.
- 1 Swamp | land | Provides a basic black source.
- 1 Acidic Slime | removal | Retained as a precon creature with a removal role.
- 1 April O'Neil, Live on the Scene | draw | Retained as a precon draw card.
- 1 Arcade Cabinet | synergy | Retained as a precon theme piece.
- 1 Arcane Signet | ramp | Retained as efficient precon ramp.
- 1 Assassin's Trophy | removal | Retained as flexible precon removal.
- 1 Baxter, Fly in the Ointment | synergy | Retained as a precon theme creature.
- 1 Bebop, Skull & Crossbones | threat | Retained as a precon threat.
- 1 Big Mother Mouser | threat | Retained as a precon creature threat.
- 1 Biogenic Ooze | threat | Retained as a precon creature threat.
- 1 Blasphemous Act | wipe | Retained as a precon board wipe.
- 1 Casey Jones, Back Alley Brute | threat | Retained as a precon creature threat.
- 1 Chromatic Lantern | ramp | Retained as a precon mana piece.
- 1 Coin of Mastery | synergy | Retained as a precon theme piece.
- 1 Continue? | interaction | Retained as a precon interaction card.
- 1 Corpsejack Menace | synergy | Retained as a precon counter-focused theme piece.
- 1 Cultivate | ramp | Retained as precon ramp.
- 1 Dimension X Pizzasaur | wincon | Retained as a precon finisher.
- 1 Donatello, the Brains | synergy | Retained as a Turtle Ninja theme creature.
- 1 Double Jump // Flying Kick | interaction | Retained as a precon interaction card.
- 1 Electric Seaweed | threat | Retained as a precon creature threat.
- 1 Endless Foot Assault | synergy | Retained as a precon theme piece.
- 1 Everything Pizza | wincon | Retained as a precon finisher.
- 1 Exploding Barrel | removal | Retained as a precon removal piece.
- 1 Fast Forward | ramp | Retained as a precon mana piece.
- 1 Foot Chopper | synergy | Retained as a precon Equipment theme piece.
- 1 Sodden Verdure | synergy | the deck upgrades a precon, and this card is one the precon holds
- 1 Game Over | wincon | Retained as a precon closing card.
- 1 Harmonize | draw | Retained as a precon draw spell.
- 1 Here Comes a New Hero! | synergy | Retained as a precon theme piece.
- 1 High Score | draw | Retained as a precon draw piece.
- 1 Irma, Part-Time Mutant | synergy | Retained as a Mutant theme creature.
- 1 Krang, the All-Powerful | threat | Retained as a precon creature threat.
- 1 Leatherhead, Iron Gator | threat | Retained as a Mutant theme creature.
- 1 Leonardo, the Balance | threat | Retained as a Turtle Ninja threat.
- 1 Lessons from Life | draw | Retained as a precon draw card.
- 1 Level Up | synergy | Retained as a precon theme piece.
- 1 Lita, Little Orphan Amphibian | synergy | Retained as a Turtle Ninja theme creature.
- 1 Michelangelo, the Heart | synergy | Retained as a Turtle Ninja theme creature.
- 1 Mole Module | synergy | Retained as a precon theme piece.
- 1 Mona Lisa, Science Geek | synergy | Retained as a Mutant theme creature.
- 1 Ninja Pizza | synergy | Retained as a precon theme piece.
- 1 Raphael, the Muscle | wincon | Retained as a Turtle Ninja finisher.
- 1 Rat King, Pale Piper | threat | Retained as a precon creature threat.
- 1 Ray Fillet, Wave Warrior | threat | Retained as a Mutant theme creature.
- 1 Roadkill Rodney | threat | Retained as a precon creature threat.
- 1 Rocksteady, Mutant Marauder | threat | Retained as a Mutant theme creature.
- 1 Shellshock | interaction | Retained as a precon interaction card.
- 1 Shredder, Shadow Master | threat | Retained as a Ninja theme creature.
- 1 Sol Ring | ramp | Retained as efficient precon ramp.
- 1 Special Move | interaction | Retained as a precon interaction card.
- 1 Steelbane Hydra | removal | Retained as a precon removal creature.
- 1 Super Combo | wincon | Retained as a precon closing card.
- 1 Swift Demise | removal | Retained as a precon removal spell.
- 1 Tempestra, Dame of Games | threat | Retained as a precon creature threat.
- 1 Together Forever | synergy | Retained as a precon theme piece.
- 1 Tokka & Rahzar, Unsupervised | ramp | Retained as a precon ramp creature.
- 1 Vanquish the Horde | wipe | Retained as a precon board wipe.
- 1 Vigor | interaction | Retained as a precon protection piece.
- 1 Voracious Hydra | removal | Retained as a precon removal creature.
- 1 Wave Goodbye | wipe | Retained as a precon board wipe.

</details>

### 19. the Hobbit family, two colours

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 174 names.

Commander: Thranduil, the Elvenking.

Grade: baseline, score 0.05, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $567.46 to buy, $567.46 the whole deck.

**Summary:** Thranduil leads an Elf-centered Sultai deck that develops its mana, fills the battlefield with Elf creatures, and turns legendary Elf arrivals into fresh cards while using Elf cards in the graveyard as a resource. The deck controls key opposing pieces with a broad mix of removal, disruption, and sweepers, then closes through its large marked finishers, chief among them the two Witch-kings. It gives up some raw speed for a creature-heavy, theme-driven plan that needs its mana and Elf presence to stay established.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. Many of the cards appear in no top list, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: The list is built around a dense Elf creature core, several legendary Elves to trigger the commander, mana elves and ramp, and a normal suite of removal and card draw, all pointing at the same board-based plan; the two black sweepers create mild tension with the creature count but are few enough not to split the deck.
- PLAN theme_fit=yes: It is a Commander deck led by Thranduil, the Elvenking in Sultai colors, drawn entirely from hob/hoc, with a mid-power Elf theme consistent with a bracket 3 ask.
- PLAN useful_as_built=yes: Thirty-five lands plus mana dorks, Signet and other ramp support the curve, and the deck can win through a wide Elf board backed by the Witch-kings and the Troll, though the eleven Swamps are slightly heavy relative to the green- and blue-weighted spell list.
- PLAN summary_honest=partly: The claims about Elf tribal, legendary Elf arrivals, removal, sweepers and the two Witch-kings as finishers all match the list, but the stated use of "Elf cards in the graveyard as a resource" is barely supported by any card here, and it does not mention that Languish and Gnashing of Teeth also hit its own Elves.
- [INFO] `curve_summary`: average mana value 3.11 over 63 nonland cards
- [WARN] `profile_off_band`: the count of lands that make two or more of the deck's colors is 4, and bracket 3 wants 17 or more (Elven Passage, Elvenking's Halls, Hobbit Hole, Mirkwood)
- [INFO] `mana_pass`: the builder moved 5 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 1 Elven Passage | land | Fills a nonbasic land slot in the mana base.
- 1 Elvenking's Halls | land | Fills a nonbasic land slot in the mana base.
- 9 Forest | land | Provides a reliable green basic land base.
- 1 Hobbit Hole | land | Fills a nonbasic land slot in the mana base.
- 8 Island | land | Provides a reliable blue basic land base.
- 1 Minas Morgul, Dark Fortress | land | Fills a nonbasic land slot in the mana base.
- 1 Mirkwood | land | Fills a nonbasic land slot in the mana base.
- 1 Rivendell | land | Fills a nonbasic land slot in the mana base.
- 11 Swamp | land | Provides a reliable black basic land base.
- 1 The Black Gate | land | Adds a nonbasic utility land slot.
- 1 The Shire | land | Fills a nonbasic land slot in the mana base.
- 1 Arcane Signet | ramp | Provides early color fixing and mana acceleration.
- 1 Delighted Halfling | ramp | Provides early creature-based mana acceleration.
- 1 Elven Chorus | ramp | Supports mana development within the Elf-heavy plan.
- 1 Elvish Archdruid | ramp | Adds Elf-based mana acceleration.
- 1 Elvish Mystic | ramp | Provides inexpensive early mana acceleration.
- 1 Mox Amber | ramp | Provides a low-cost mana source alongside the deck's legends.
- 1 Necklace of Girion | ramp | Adds another mana-development artifact.
- 1 Thranduil the Strategist | ramp | Contributes mana development while being a legendary Elf.
- 1 Wood Elves | ramp | Adds creature-based mana development and Elf density.
- 1 Elvish Visionary | draw | Provides card flow while adding an Elf to the deck.
- 1 Fateful Discovery | draw | Adds a dedicated source of card flow.
- 1 Hithlain Knots | draw | Adds a dedicated source of card flow.
- 1 Ithilien Kingfisher | draw | Provides creature-based card flow.
- 1 Key to the Side-Door | draw | Adds a repeatable-looking card-flow artifact.
- 1 Lórien Revealed | draw | Adds flexible card flow.
- 1 Night's Whisper | draw | Provides efficient early card flow.
- 1 Palantír of Orthanc | draw | Adds a persistent card-flow artifact.
- 1 Plunder the Trollshaws | draw | Adds another source of card flow.
- 1 Thrór's Map | draw | Provides artifact-based card flow.
- 1 Uncover the Moon-Letters | draw | Adds a dedicated source of card flow.
- 1 Bilbo's Deadly Slice | removal | Provides a focused removal spell.
- 1 Bitter Downfall | removal | Provides efficient spot removal.
- 1 Enchanted River's Grasp | removal | Adds a permanent-focused removal option.
- 1 Giant's Boulder | removal | Adds artifact-based removal.
- 1 Merciless Executioner | removal | Provides creature-based removal.
- 1 Orcish Bowmasters | removal | Provides efficient creature-based removal.
- 1 Quarrel | removal | Adds another flexible removal spell.
- 1 The Black Arrow | removal | Provides equipment-based removal.
- 1 Uneasy Partings | removal | Adds a further removal option.
- 1 Bilbo's Ring | interaction | Provides a protective interaction piece.
- 1 Confusticate and Bebother | interaction | Adds stack-oriented interaction.
- 1 Dwarven Mattock | interaction | Provides an artifact-based interaction tool.
- 1 Mithril Coat | interaction | Helps protect an important creature.
- 1 My Precious // Allure of Power | interaction | Adds a flexible legendary interaction piece.
- 1 Stern Scolding | interaction | Provides inexpensive stack interaction.
- 1 The One Ring | interaction | Adds a resilient interaction artifact.
- 1 Thranduil's Decree | interaction | Provides another themed interaction spell.
- 1 Gnashing of Teeth | wipe | Provides a battlefield reset when pressure builds.
- 1 Languish | wipe | Provides a compact creature sweep effect.
- 1 Raise the Palisade | wipe | Provides a tribal-aware battlefield reset.
- 1 Arwen, Weaver of Hope | synergy | Adds a legendary Elf for Thranduil's legendary-Elf trigger.
- 1 Celeborn the Wise | synergy | Adds a legendary Elf for Thranduil's legendary-Elf trigger.
- 1 Elrond, Moon-Reader | synergy | Adds a legendary Elf for Thranduil's legendary-Elf trigger.
- 1 Galion, Elvenking's Butler | synergy | Adds a legendary Elf for Thranduil's legendary-Elf trigger.
- 1 Mirkwood Meditator | synergy | Increases Elf density for Thranduil's graveyard plan.
- 1 Mirkwood Nurturer | synergy | Increases Elf density for Thranduil's graveyard plan.
- 1 Thranduil, Sindarin Liege // Silvan Rally | synergy | Adds another legendary Elf for Thranduil's legendary-Elf trigger.
- 1 Boughside Wanderers | threat | Supplies an Elf creature to pressure opponents.
- 1 Cantankerous Keepers | threat | Supplies an Elf creature to pressure opponents.
- 1 Elven Raft-Steerer | threat | Supplies an Elf creature to build the battlefield.
- 1 Elvenking's Harper | threat | Supplies an Elf creature to build the battlefield.
- 1 Galadhrim Guide | threat | Supplies an Elf creature to build the battlefield.
- 1 Grey Havens Navigator | threat | Supplies an Elf creature to build the battlefield.
- 1 Guardian of the Halls | threat | Supplies an Elf creature to build the battlefield.
- 1 Haunt of the Dead Marshes | threat | Adds an Elf creature that contributes battlefield pressure.
- 1 Lothlórien Lookout | threat | Supplies an Elf creature to build the battlefield.
- 1 Mirkwood Pathmaker | threat | Supplies an Elf creature to build the battlefield.
- 1 Nimrodel Watcher | threat | Supplies an Elf creature to build the battlefield.
- 1 Thranduil's Company | threat | Adds a themed Elf body to develop battlefield pressure.
- 1 Troll of Khazad-dûm | wincon | Provides a marked finisher and closing threat.
- 1 Witch-king of Angmar | wincon | Provides a marked finisher and closing threat.
- 1 Witch-king, Bringer of Ruin | wincon | Provides a marked finisher and closing threat.
- 1 Woodland Weavemaster | ramp | the mana pass added it to bring the mana base inside the power level

</details>

### 20. the Hobbit family, a delegated commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 77 names.

Commander: Smaug the Magnificent.

Grade: baseline, score 0.05, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: yes, for copy_limit, profile_off_band. Block findings: 0.

Cost: $503.32 to buy, $503.32 the whole deck.

**Summary:** A mono-red Hobbit Dragon deck centered on building Treasures, equipping resilient creatures, and leveraging combat to create overwhelming pressure. Smaug turns accumulated wealth into direct damage, while Dragons, giants, and a wide supporting cast provide multiple routes to a decisive finish. Removal and sweepers keep opposing boards manageable while legendary artifacts reinforce the deck’s adventurous Middle-earth character.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The list is a coherent mono-red aggro/equipment pile with Treasure-ish ramp and burn, but the summary's Dragon focus is carried by only three Dragons plus the commander, and the equipment package and goblin/dwarf beaters pull somewhat apart from any dragon plan.
- PLAN theme_fit=partly: It is a legal Hobbit-set Commander deck at a reasonable mid-power level with a Dragon commander, but the requested dragons theme is thin, the bulk being dwarves, goblins and equipment.
- PLAN useful_as_built=yes: 36 lands in mono-red with cheap ramp, a low curve, ample removal and plenty of creature-plus-equipment damage means it plays and closes games fine as written.
- PLAN summary_honest=partly: Treasures, equipment, removal, sweepers and legendary artifacts are all genuinely present, but calling Dragons a pillar overstates a three-Dragon list and 'giants' amounts to a single Inferno Titan.
- [INFO] `curve_summary`: average mana value 3.16 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Arcane Signet | ramp | Early colorless acceleration for deploying Smaug and the deck's larger threats.
- 1 Bag End Banquet | ramp | Treasure-focused mana acceleration.
- 1 Burn, Burn, Tree and Fern | ramp | Saga-based ramp that supports the deck's mana development.
- 1 Dragon's Desire | ramp | Mana acceleration suited to the Dragon plan.
- 1 Fíli and Kíli, Joyous | ramp | Creature-based ramp with Hobbit-theme flavor.
- 1 Long-Bodied Grey Dog | ramp | Low-cost mana development.
- 1 Mox Amber | ramp | Fast legendary-matters acceleration.
- 1 Orcrist, Goblin-cleaver | ramp | Equipment that contributes to Treasure-oriented mana growth.
- 1 The Misty Mountains Cold | ramp | Saga ramp for reaching expensive threats.
- 1 The Reaver Cleaver | ramp | Treasure production that turns combat into mana.
- 1 Wayfarer's Bauble | ramp | Reliable early land ramp.
- 1 Balin, Loremaster | draw | Creature-based card advantage.
- 1 Key to the Side-Door | draw | Artifact card selection and advantage.
- 1 Palantír of Orthanc | draw | Repeatable card advantage source.
- 1 Ragged Short Spear | draw | Equipment-based card advantage.
- 1 Thrór's Map | draw | Card advantage attached to a thematic artifact.
- 1 Óin the Brave | draw | Legendary creature that provides card advantage.
- 1 Getaway Barrel | draw | Flexible artifact card advantage.
- 1 Old Thrush | draw | Creature-based card advantage for sustained pressure.
- 1 Bilbo's Ring | interaction | Protects an important creature or advances through hostile boards.
- 1 Dwarven Mattock | interaction | Flexible equipment interaction.
- 1 Mithril Coat | interaction | Protects Smaug and key creatures from removal.
- 1 The One Ring | interaction | Defensive utility against opposing pressure.
- 1 Battle-Scarred Goblin | removal | Creature-based removal.
- 1 Fire of Orthanc | removal | Direct removal for problematic permanents.
- 1 Gandalf, Spark Starter | removal | Legendary removal tool with board presence.
- 1 Giant's Boulder | removal | Artifact-based spot removal.
- 1 Goblin Cratermaker | removal | Efficient creature removal with additional utility.
- 1 Goblin Fireleaper | removal | Additional creature-based spot removal.
- 1 Improvised Club | removal | Instant-speed removal.
- 1 Inferno Titan | removal | Repeatable damage-based removal on a substantial body.
- 1 Smite the Deathless | removal | Efficient answer to opposing threats.
- 1 The Black Arrow | removal | Thematic targeted removal.
- 1 Call Forth the Tempest | wipe | Mass removal for resetting overwhelming boards.
- 1 Desolation of Smaug | wipe | Sweeper that complements the commander’s destructive theme.
- 1 Glóin the Mighty // Easy Pickings | wipe | Creature that also offers a board-clearing Adventure.
- 1 Cavern-Hoard Dragon | wincon | Treasure-driven Dragon finisher.
- 1 Desert Were-Worm | wincon | Large evasive finisher.
- 1 Smaug, the Great Calamity // Spew Flame | wincon | Dragon threat that closes games and offers removal.
- 1 Andúril, Flame of the West | synergy | Legendary Equipment supporting combat-focused threats.
- 1 Andúril, Narsil Reforged | synergy | Equipment that reinforces the deck’s legendary and combat themes.
- 1 Glamdring | synergy | Thematic legendary Equipment for creature combat.
- 1 Last Light of Durin's Day | synergy | Supports the deck’s thematic permanent package.
- 1 Long-Lost Lances | synergy | Equipment that improves combat damage.
- 1 Sting, Bilbo's Sword | synergy | Legendary Equipment supporting evasive attacks.
- 1 Well-Worn Spatula | synergy | Equipment synergy for the creature suite.
- 1 Bothersome Noisemaker | threat | Creature that applies pressure while fitting the setting.
- 1 Dwarven Mauler | threat | Combat-focused creature threat.
- 1 Dwarven Warriors | threat | Creature pressure that benefits from the equipment package.
- 1 Goblin-town Flunkies | threat | Low-cost battlefield presence.
- 1 Gundabad Opportunist | threat | Aggressive creature threat.
- 1 Guttersnipe | threat | Turns the deck’s spells into repeated opponent damage.
- 1 Iron Hills Stalwart | threat | Durable combat threat.
- 1 Misty Mountains Raider | threat | Creature pressure for the midgame.
- 1 Olog-hai Crusher | threat | Large combat threat.
- 1 Orcish Siegemaster | threat | Creature that pressures opponents and supports the aggressive plan.
- 1 Snowslope Hunter | threat | Evasive combat-oriented creature.
- 1 Thorin, Company's Leader | threat | Legendary threat that also rewards the deck’s development plan.
- 1 Bombur, Gentle Dreamer | other | Thematic legendary creature utility.
- 1 Dori, Bearer of Friends | other | Thematic legendary creature utility.
- 1 Gandalf, Goblins' Bane // Flameshape | other | Flexible thematic spell and creature.
- 1 Smaug's Fury | other | Flexible instant-speed thematic utility.
- 1 Tidings of War | other | Thematic sorcery that supports the broader game plan.
- 35 Mountain | land | Untapped red source for a mono-red deck.
- 1 The Lonely Mountain | land | Thematic red source.

</details>

### 21. the Hobbit family, mana from outside

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 156 names.

Commander: Smaug the Impenetrable.

Grade: baseline, score 0.07, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $983.84 to buy, $983.84 the whole deck.

**Summary:** This Rakdos Smaug deck develops its mana early, deploys Smaug the Impenetrable and other large threats, then presses combat while using focused removal and broad resets to clear resistance. Smaug is chief among the finishers, with Dragons and other legendary threats providing several ways to end a stalled game. It gives up some flexibility for a committed battlefield plan and depends on its mana development to bring its expensive threats online quickly.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: The cards line up behind one plan — cheap rocks and ramp spells into Smaug and a stack of large legendary/Dragon threats, with equipment and burn to push damage through and spot removal plus two sweepers to clear blockers — so the pieces largely pull the same direction.
- PLAN theme_fit=yes: It is a Rakdos Commander deck led by Smaug the Impenetrable with essentially every nonland card in hob/hoc and only lands (plus in-set rocks) drawn from outside, matching the request's stated allowance and a mid-bracket power level.
- PLAN useful_as_built=yes: Around thirty-two lands with fixing, fetches and duals in a two-color deck, a spread of cheap rocks and ramp, a normal curve, and plenty of finishers means the deck can be picked up and played as it stands.
- PLAN summary_honest=partly: The described ramp-into-big-threats-plus-removal-and-resets arc is genuinely in the list, but the summary never mentions the sizable Equipment package (Andúril x2, Glamdring, Sting, Mithril Coat, Reaver Cleaver, Orcrist, etc.) that is a real part of how the deck presses combat, and it glosses over that Languish/Desolation can clip its own small creatures.
- [INFO] `curve_summary`: average mana value 3.02 over 62 nonland cards
- [INFO] `mana_pass`: the builder moved 6 cards of the mana base to bring the deck inside its power level
- [WARN] `outside_requested_set`: the sets you named do not hold 18 cards: Ancient Tomb, Arid Mesa, Blood Crypt, Bloodstained Mire, Bojuka Bog, Command Tower, Dragonskull Summit, Evolving Wilds, Exotic Orchard, Fabled Passage, Marsh Flats, Path of Ancestry, Polluted Delta, Scalding Tarn, Terramorphic Expanse, Urborg, Tomb of Yawgmoth, Verdant Catacombs, Wooded Foothills

<details><summary>The deck list</summary>

- 1 Ancient Tomb | land | Accelerates the mana base.
- 1 Arid Mesa | land | Helps assemble the mana base.
- 1 Blood Crypt | land | Provides a dual-color land slot.
- 1 Bloodstained Mire | land | Helps assemble the mana base.
- 1 Bojuka Bog | land | Provides a utility land slot.
- 1 Command Tower | land | Provides reliable color fixing.
- 1 Dragon-Cursed Halls | land | Provides a dual-color land slot.
- 1 Dragonskull Summit | land | Provides a dual-color land slot.
- 1 Evolving Wilds | land | Helps assemble the mana base.
- 1 Exotic Orchard | land | Provides color fixing.
- 1 Fabled Passage | land | Helps assemble the mana base.
- 1 Goblin-town | land | Provides a colored land slot.
- 1 Hobbit Hole | land | Provides a colored land slot.
- 1 Marsh Flats | land | Helps assemble the mana base.
- 1 Minas Morgul, Dark Fortress | land | Provides a colored land slot.
- 1 Mount Doom | land | Provides a colored land slot.
- 8 Mountain | land | Supplies a dependable red mana base.
- 1 Path of Ancestry | land | Provides color fixing.
- 1 Polluted Delta | land | Helps assemble the mana base.
- 1 Rogue's Passage | land | Provides a utility land slot.
- 1 Scalding Tarn | land | Helps assemble the mana base.
- 2 Swamp | land | Supplies a dependable black mana base.
- 1 Terramorphic Expanse | land | Helps assemble the mana base.
- 1 The Black Gate | land | Provides a colored land slot.
- 1 The Lonely Mountain | land | Provides a colored land slot.
- 1 Treasure Vault | land | Provides a utility land slot.
- 1 Urborg, Tomb of Yawgmoth | land | Strengthens black mana access.
- 1 Verdant Catacombs | land | Helps assemble the mana base.
- 1 Wooded Foothills | land | Helps assemble the mana base.
- 1 Arcane Signet | ramp | Provides efficient color fixing and mana development.
- 1 Bag End Banquet | ramp | Provides a ramp slot.
- 1 Burn, Burn, Tree and Fern | ramp | Provides a ramp slot.
- 1 Dragon's Desire | ramp | Provides a ramp slot.
- 1 Mox Amber | ramp | Provides early mana development.
- 1 Orcrist, Goblin-cleaver | ramp | Provides a ramp slot.
- 1 Smaug the Magnificent | ramp | Provides mana development while adding a substantial body.
- 1 Smaug, Wicked Worm | ramp | Provides mana development and a powerful late-game threat.
- 1 The Reaver Cleaver | ramp | Provides a ramp slot.
- 1 Wayfarer's Bauble | ramp | Provides early mana development.
- 1 Balin, Loremaster | draw | Provides a draw slot.
- 1 Key to the Side-Door | draw | Provides a draw slot.
- 1 Night's Whisper | draw | Provides efficient card flow.
- 1 Óin the Brave | draw | Provides a draw slot.
- 1 Palantír of Orthanc | draw | Provides a draw slot.
- 1 Rage into the Valley | draw | Provides a draw slot.
- 1 Ragged Short Spear | draw | Provides a draw slot.
- 1 Reverent Howl | draw | Provides a draw slot.
- 1 The Master of Lake-town | draw | Provides a draw slot.
- 1 The Sackville-Bagginses | draw | Provides a draw slot.
- 1 Thrór's Map | draw | Provides a draw slot.
- 1 Along the Crooked Way | interaction | Fills an interaction slot.
- 1 Bilbo's Ring | interaction | Fills an interaction slot.
- 1 Dwarven Mattock | interaction | Fills an interaction slot.
- 1 Getaway Barrel | interaction | Fills an interaction slot.
- 1 Mithril Coat | interaction | Fills an interaction slot.
- 1 My Precious // Allure of Power | interaction | Fills an interaction slot.
- 1 Smaug's Fury | interaction | Fills an interaction slot.
- 1 The One Ring | interaction | Fills an interaction slot.
- 1 Bilbo's Deadly Slice | removal | Provides focused removal.
- 1 Bitter Downfall | removal | Provides focused removal.
- 1 Fire of Orthanc | removal | Provides focused removal.
- 1 Goblin Cratermaker | removal | Provides creature-based removal.
- 1 Improvised Club | removal | Provides focused removal.
- 1 Inferno Titan | removal | Provides removal on a substantial body.
- 1 Orcish Bowmasters | removal | Provides creature-based removal.
- 1 Smite the Deathless | removal | Provides focused removal.
- 1 Stir Up Trouble | removal | Provides focused removal.
- 1 Andúril, Flame of the West | synergy | Supports the deck's combat-focused core plan.
- 1 Andúril, Narsil Reforged | synergy | Supports the deck's combat-focused core plan.
- 1 Glamdring | synergy | Supports the deck's combat-focused core plan.
- 1 Goblin Plate Mail | synergy | Supports the deck's combat-focused core plan.
- 1 Long-Lost Lances | synergy | Supports the deck's combat-focused core plan.
- 1 Sting, Bilbo's Sword | synergy | Supports the deck's combat-focused core plan.
- 1 Supper for Spiders | synergy | Provides a dedicated synergy slot.
- 1 Bothersome Noisemaker | threat | Adds a proactive creature threat.
- 1 Desert Were-Worm | threat | Adds a high-impact finisher threat.
- 1 Down, Down to Goblin-town | threat | Adds a high-impact finisher threat.
- 1 Dreaded Bat-Cloud | threat | Adds a proactive creature threat.
- 1 Gollum the Abandoned | threat | Adds a proactive creature threat.
- 1 Great Goblin, Foul-Hearted | threat | Adds a proactive creature threat.
- 1 Guttersnipe | threat | Adds a proactive creature threat.
- 1 Haunt of the Dead Marshes | threat | Adds a proactive creature threat.
- 1 Sauron, the Lidless Eye | threat | Adds a proactive creature threat.
- 1 Troll of Khazad-dûm | threat | Adds a high-impact finisher threat.
- 1 Witch-king of Angmar | threat | Adds a high-impact finisher threat.
- 1 Witch-king, Bringer of Ruin | threat | Adds a high-impact finisher threat.
- 1 Gollum, Riddle Master | wincon | Provides a finisher for closing games.
- 1 Desolation of Smaug | wipe | Provides a board-reset option.
- 1 Languish | wipe | Provides a board-reset option.
- 1 Giant's Boulder | ramp | the mana pass added it to bring the mana base inside the power level
- 1 Bolg's Company | ramp | the mana pass added it to bring the mana base inside the power level

</details>

### 22. a set family and a card from outside it

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 175 names.

Commander: Thranduil, the Elvenking.

Grade: baseline, score 0.07, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $587.60 to buy, $587.60 the whole deck.

**Summary:** This is a Sultai Elf deck centered on Thranduil, the Elvenking, using legendary Elves and a full Elf creature base to develop pressure while keeping mana and cards flowing. Targeted answers, protective tools, and sweepers give it room to stabilize, then Troll of Khazad-dûm and the two Witch-kings provide its strongest closing threats. It gives up some raw speed for a creature-centric plan that benefits from a developed board and a stocked graveyard.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. Many of the cards appear in no top list, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core is a consistent Sultai legendary-Elf board-building deck with ramp and card flow, but the three sweepers (Languish, Gnashing of Teeth, Raise the Palisade) work against its own small-Elf board, pulling mildly against the creature plan the rest of the list builds.
- PLAN theme_fit=yes: It is a Commander deck led by Thranduil, the Elvenking with Sol Ring kept and every other card drawn from the requested hob/hoc sets, at a reasonable bracket-3 power level.
- PLAN useful_as_built=yes: 36 lands plus several ramp pieces support a curve topping near six, and it has creatures, evasive top-end threats, removal and draw, so it plays fine as presented.
- PLAN summary_honest=yes: Everything named — Elf base, legendary Elves, targeted answers, equipment protection, sweepers, and the Troll/Witch-king top end — is present in the list, and it openly concedes the deck's slower speed.
- [INFO] `curve_summary`: average mana value 3.11 over 63 nonland cards
- [WARN] `profile_off_band`: the count of lands that make two or more of the deck's colors is 4, and bracket 3 wants 17 or more (Elven Passage, Elvenking's Halls, Hobbit Hole, Mirkwood)
- [INFO] `mana_pass`: the builder moved 4 cards of the mana base to bring the deck inside its power level
- [WARN] `outside_requested_set`: the sets you named do not hold 1 card: Sol Ring

<details><summary>The deck list</summary>

- 1 Sol Ring | ramp | Efficient mana acceleration the deck was asked to retain.
- 1 Arcane Signet | ramp | Reliable color fixing and acceleration.
- 1 Elven Chorus | ramp | Elf-focused mana acceleration.
- 1 Elvish Archdruid | ramp | Elf-focused mana acceleration.
- 1 Mox Amber | ramp | Efficient legendary-focused acceleration.
- 1 Necklace of Girion | ramp | Additional mana acceleration.
- 1 Orcrist, Goblin-cleaver | ramp | Equipment-based mana acceleration.
- 1 Thranduil the Strategist | ramp | A legendary Elf that advances the mana plan.
- 1 Wood Elves | ramp | Elf-based mana acceleration.
- 1 Woodland Weavemaster | ramp | Elf-based mana acceleration.
- 1 Beorn the Fierce | draw | Creature-based card advantage.
- 1 Elvish Visionary | draw | Elf-based card advantage.
- 1 Fateful Discovery | draw | Dedicated card advantage.
- 1 Gandalf, Wandering Wizard | draw | Legendary card advantage.
- 1 Hithlain Knots | draw | Instant-speed card advantage.
- 1 Lórien Revealed | draw | Flexible card advantage.
- 1 Night's Whisper | draw | Efficient card advantage.
- 1 Palantír of Orthanc | draw | Persistent card advantage.
- 1 Plunder the Trollshaws | draw | Instant-speed card advantage.
- 1 Uncover the Moon-Letters | draw | Additional card advantage.
- 1 Bilbo's Ring | interaction | Protection and utility for key creatures.
- 1 Confusticate and Bebother | interaction | Flexible disruptive interaction.
- 1 Elrond, Moon-Reader | interaction | A legendary Elf with interactive utility.
- 1 Mithril Coat | interaction | Protects an important creature.
- 1 My Precious // Allure of Power | interaction | Versatile interactive utility.
- 1 Old Fat Spider Can't See Me | interaction | Ongoing disruptive interaction.
- 1 Stern Scolding | interaction | Efficient stack interaction.
- 1 The One Ring | interaction | Powerful defensive utility.
- 1 Bilbo's Deadly Slice | removal | Efficient targeted removal.
- 1 Bitter Downfall | removal | Dedicated targeted removal.
- 1 Crude Bent Blade | removal | Equipment-based removal utility.
- 1 Enchanted River's Grasp | removal | Aura-based removal.
- 1 Giant's Boulder | removal | Artifact-based removal.
- 1 Merciless Executioner | removal | Creature-based removal.
- 1 Orcish Bowmasters | removal | Efficient creature-based removal.
- 1 Quarrel | removal | Flexible instant-speed removal.
- 1 Uneasy Partings | removal | Additional instant-speed removal.
- 1 Gnashing of Teeth | wipe | Broad battlefield reset.
- 1 Languish | wipe | Efficient creature sweep.
- 1 Raise the Palisade | wipe | Tribal-aware battlefield reset.
- 1 Arwen, Weaver of Hope | synergy | Legendary Elf support for Thranduil's plan.
- 1 Celeborn the Wise | synergy | Legendary Elf support for Thranduil's plan.
- 1 Galion, Elvenking's Butler | synergy | Legendary Elf support for Thranduil's plan.
- 1 Mirkwood Meditator | synergy | Elf support for the creature-focused plan.
- 1 Supper for Spiders | synergy | Dedicated synergy piece.
- 1 Thranduil, Sindarin Liege // Silvan Rally | synergy | Legendary Elf support for Thranduil's plan.
- 1 Cantankerous Keepers | threat | Elf creature pressure.
- 1 Elven Raft-Steerer | threat | Elf creature pressure.
- 1 Elvenking's Harper | threat | Elf creature pressure.
- 1 Galadhrim Guide | threat | Elf creature pressure.
- 1 Grey Havens Navigator | threat | Elf creature pressure.
- 1 Guardian of the Halls | threat | Elf creature pressure.
- 1 Haunt of the Dead Marshes | threat | Elf creature pressure.
- 1 Lothlórien Lookout | threat | Elf creature pressure.
- 1 Mirkwood Nurturer | threat | Elf creature pressure.
- 1 Mirkwood Pathmaker | threat | Elf creature pressure.
- 1 Nimrodel Watcher | threat | Elf creature pressure.
- 1 Thranduil's Company | threat | Elf creature pressure.
- 1 Troll of Khazad-dûm | wincon | A marked finisher and closing threat.
- 1 Witch-king of Angmar | wincon | A marked finisher and closing threat.
- 1 Witch-king, Bringer of Ruin | wincon | A marked finisher and closing threat.
- 1 Elven Passage | land | Nonbasic mana-base slot.
- 1 Elvenking's Halls | land | Nonbasic mana-base slot.
- 1 Hobbit Hole | land | Nonbasic mana-base slot.
- 1 Minas Morgul, Dark Fortress | land | Nonbasic mana-base slot.
- 1 Mirkwood | land | Nonbasic mana-base slot.
- 1 Rivendell | land | Nonbasic mana-base slot.
- 1 The Black Gate | land | Nonbasic mana-base slot.
- 1 The Shire | land | Nonbasic mana-base slot.
- 1 Treasure Vault | land | Nonbasic mana-base slot.
- 9 Forest | land | Primary green mana base.
- 8 Island | land | Primary blue mana base.
- 10 Swamp | land | Primary black mana base.
- 1 Pelargir Survivor | ramp | the mana pass added it to bring the mana base inside the power level
- 1 Delighted Halfling | ramp | the mana pass added it to bring the mana base inside the power level

</details>

### 23. two set families at once

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 203 names.

Commander: Kíli the Resourceful.

Grade: typical, score 0.47, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $264.09 to buy, $264.09 the whole deck.

**Summary:** Kíli the Resourceful leads a mono-white Dwarf and Equipment deck that develops through artifact mana, cheap gear, and a steady flow of creatures. Establish an enduring story early, use Kíli to turn each turn's first equip into efficient pressure, and keep cards flowing as Dwarves and Equipment enter. The deck wins by building one or more well-equipped attackers, with Angel of the Ruins, Helm of the Host, and Sunscorch Regent serving as major closing threats. It gives up broad color access and relies on its white mana base, artifacts, and creature board to carry the game.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

This deck plan aims for 2 or more finishers, the cards that win the game, and the deck holds 1.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core is consistent — cheap Dwarves, a dozen-plus Equipment, artifact ramp, and card flow off creatures/artifacts entering — but three board wipes (Dusk // Dawn, Martial Coup, Promise of Loyalty) pull against a plan built on sticking and suiting up a creature board, and the Dwarf count (~7) is thin for the tribal half of the pitch.
- PLAN theme_fit=yes: Every card is listed as in the requested BLB/HOB set pool, the builder picked a commander as asked, and the curve, interaction suite, and lack of infinite combos read as a reasonable bracket 3 Commander deck.
- PLAN useful_as_built=yes: 36 lands in mono-white with Sol Ring, Arcane Signet, and other cheap ramp support a low curve, and the deck has removal, sweepers, card draw, and several equipped attackers plus fliers to actually close a game.
- PLAN summary_honest=partly: Most claims check out (white mana base, artifact ramp, equipment, draw-on-enter effects, the three named threats are present), but 'establish an enduring story early' points to nothing in the list, and the summary does not mention the sweepers that cut against its own board-building plan.
- [INFO] `curve_summary`: average mana value 2.71 over 63 nonland cards
- [WARN] `finisher_short`: the finisher count is 1, and this deck plan wants 2 or more (Angel of the Ruins)

<details><summary>The deck list</summary>

- 28 Plains | land | Reliable white mana keeps the opening mana base stable.
- 1 Command Tower | land | An untapped white source for the deck.
- 1 Exotic Orchard | land | Provides flexible mana while remaining an untapped land.
- 1 Evolving Wilds | land | Finds a Plains and helps stabilize white mana.
- 1 Fabled Passage | land | Fetches a Plains for dependable color access.
- 1 Path of Ancestry | land | A tribal land that supports the Dwarf-focused creature base.
- 1 Thriving Heath | land | A white source with flexible color selection.
- 1 Uncharted Haven | land | Provides another flexible source of mana.
- 1 Terramorphic Expanse | land | Finds a Plains to support consistent white mana.
- 1 Caretaker's Talent | draw | Provides a dedicated card-advantage engine.
- 1 Circuit Mender | draw | An artifact creature that supplies card advantage and supports the artifact count.
- 1 Cut a Deal | draw | Refills the hand after committing creatures and Equipment.
- 1 Dawn of a New Age | draw | A lasting source of card advantage.
- 1 Heirloom Epic | draw | An artifact-based source of cards that also helps establish the story.
- 1 Idol of Oblivion | draw | A low-cost artifact that provides repeatable card access.
- 1 Inspiring Overseer | draw | Supplies a creature body and immediate card advantage.
- 1 Mangara, the Diplomat | draw | Provides card advantage from a resilient creature.
- 1 Mentor of the Meek | draw | Turns the deck's smaller creatures into steady card advantage.
- 1 Skullclamp | draw | An Equipment-based draw engine that works with the deck's creature plan.
- 1 Spirited Companion | draw | A cheap permanent that replaces itself and supports the creature count.
- 1 Arcane Signet | ramp | Efficient early mana acceleration.
- 1 Bag End Banquet | ramp | Artifact mana support contributes to Kíli's enduring-story setup.
- 1 Burnished Hart | ramp | Provides repeatable creature-based access to more mana.
- 1 Fellwar Stone | ramp | A cheap mana rock for early acceleration.
- 1 Mind Stone | ramp | Early acceleration that remains useful later.
- 1 Mox Amber | ramp | A fast legendary mana source for the commander and legendary permanents.
- 1 Ornithopter of Paradise | ramp | An artifact creature that fixes mana and builds the artifact count.
- 1 Orcrist, Goblin-cleaver | ramp | An Equipment that also advances the mana plan.
- 1 Sol Ring | ramp | The deck's strongest early burst of artifact mana.
- 1 Wayfarer's Bauble | ramp | Finds a Plains and supports early mana development.
- 1 Banishing Light | removal | Flexible permanent-based removal.
- 1 Bumbleflower's Sharepot | removal | Artifact removal that contributes to the artifact count.
- 1 Fiend Hunter | removal | Creature-based removal that can carry Equipment.
- 1 Generous Gift | removal | Answers a broad range of opposing permanents.
- 1 Giant's Boulder | removal | Artifact removal that helps maintain an artifact-heavy board.
- 1 Hanged Executioner | removal | Provides a creature body with removal utility.
- 1 Loran of the Third Path | removal | Versatile removal attached to a creature.
- 1 Skyclave Apparition | removal | Efficient creature-based permanent removal.
- 1 Swords to Plowshares | removal | A cheap, direct answer to opposing creatures.
- 1 Bilbo's Ring | interaction | A legendary Equipment that protects the equipped creature plan.
- 1 Bofur, Reliable Guardian // Concerted Care | interaction | A Dwarf that offers a useful protective option.
- 1 Dwarven Mattock | interaction | An Equipment-based interactive tool for the Dwarf plan.
- 1 Mithril Coat | interaction | Protects an important creature while adding an artifact.
- 1 Parting Gust | interaction | Flexible instant-speed interaction.
- 1 Reprieve | interaction | A compact answer that preserves card flow.
- 1 Selfless Spirit | interaction | Protects the creature and Equipment board from opposing answers.
- 1 Swiftfoot Boots | interaction | Protects key creatures and helps them contribute immediately.
- 1 Dusk // Dawn | wipe | A board reset that can later return smaller creatures.
- 1 Martial Coup | wipe | A scalable reset that can leave behind an attacking force.
- 1 Promise of Loyalty | wipe | A creature-focused wipe that breaks up opposing boards.
- 1 Blade Splicer | synergy | Builds multiple artifact-related bodies for the Equipment plan.
- 1 Carrot Cake | synergy | An artifact permanent that supports the deck's board development.
- 1 Dáin, Lord of the Iron Hills | synergy | A legendary Dwarf that strengthens the tribal core.
- 1 Dwarven Provisioner | synergy | Another Dwarf to trigger Kíli and support tribal pressure.
- 1 Iron Hills Blacksmith | synergy | A Dwarf Artificer that directly fits the Equipment theme.
- 1 Ori, Keeper of Songs | synergy | A legendary Dwarf that reinforces the Dwarf package.
- 1 Tangle Tumbler | synergy | An artifact Vehicle that supports the artifact-heavy strategy.
- 1 Fíli the Pathfinder | threat | A Dwarf threat that benefits from the tribal and Equipment shell.
- 1 Karn, the Great Creator | threat | A durable colorless threat in the artifact-focused deck.
- 1 Andúril, Flame of the West | threat | A legendary Equipment that turns a creature into a major attacker.
- 1 Andúril, Narsil Reforged | threat | A second legendary Equipment for sustained combat pressure.
- 1 Dúnedain Blade | threat | An Equipment that helps convert small creatures into threats.
- 1 Dwarven Shortsword | threat | A cheap Equipment for Kíli's card-draw and combat plan.
- 1 Glamdring | threat | A legendary Equipment that increases the deck's attacking pressure.
- 1 Long-Lost Lances | threat | An Equipment threat for the creature-heavy board.
- 1 Starforged Sword | threat | An Equipment that makes even modest creatures threatening.
- 1 Sting, Bilbo's Sword | threat | A legendary Equipment that supports Kíli's Equipment-focused plan.
- 1 Sword of Vengeance | threat | An Equipment that upgrades creatures into dangerous attackers.
- 1 Sword of the Squeak | threat | A low-cost Equipment threat that helps trigger Kíli.
- 1 Angel of the Ruins | wincon | A marked finisher that gives the deck a powerful closing threat.
- 1 Helm of the Host | wincon | A legendary Equipment that can create overwhelming board pressure.
- 1 Sunscorch Regent | wincon | A large flying threat for closing stalled games.

</details>

### 24. a 60-card deck from one set

Format: Modern. Theme: aggro. Pool: any_card. Shortlist: 96 names.

Grade: baseline, score 0.40, model 20260923T202806Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $33.86 to buy, $33.86 the whole deck.

**Summary:** This mono-red Bloomburrow aggro deck wins by building an attacking force of Mice, Lizards, and Raccoons, then using removal to keep its pressure pointed at the opponent. Emberheart Challenger and Hearthborn Battler reinforce the creature core, while Dragonhawk, Fate's Tempest is chief among the deck's threats. It gives up broad answers and a deep long game in exchange for a focused, consistent attack plan.

The quality model grades this deck at the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Every nonland card either adds a body, removes a blocker, or refills the hand to keep attacking, so the pieces all point at the same creature-pressure plan.
- PLAN theme_fit=partly: It is mono-red, Modern, and drawn entirely from the requested Bloomburrow sets, but with 24 lands and fourteen creatures costing four or five mana against only six one-drops it plays more like red midrange than the aggro deck asked for.
- PLAN useful_as_built=yes: A single basic land type with 24 lands casts everything reliably, and the many four- and five-drop threats plus six burn spells give clear, repeatable ways to close a game even if the curve is slower than ideal.
- PLAN summary_honest=yes: The named cards are all present in the listed counts, and the summary openly concedes the lack of broad answers and late-game staying power, which matches a list of small creatures, burn, and no sweepers or recursion.
- [INFO] `curve_summary`: average mana value 3.11 over 36 nonland cards

<details><summary>The deck list</summary>

- 24 Mountain | land | Provides the consistent red mana base for a mono-red deck.
- 2 Artist's Talent | draw | Supplies card draw within the red aggressive shell.
- 2 Might of the Meek | draw | Provides efficient card draw while supporting the creature-focused plan.
- 2 Sazacap's Brew | draw | Adds red card draw to keep pressure flowing.
- 2 Agate Assault | removal | Provides mono-red removal for opposing threats.
- 2 Flame Lash | removal | Adds direct removal to clear a path for attackers.
- 2 Playful Shove | removal | Provides additional removal for the aggressive plan.
- 4 Emberheart Challenger | synergy | A Mouse Warrior synergy piece for the creature-focused attack plan.
- 4 Hearthborn Battler | synergy | A Lizard Warlock synergy piece that supports the aggressive creature core.
- 4 Frilled Sparkshooter | threat | A Lizard Archer threat that contributes to the attacking creature base.
- 4 Reptilian Recruiter | threat | A Lizard Warrior threat for the deck's creature pressure.
- 4 Teapot Slinger | threat | A Raccoon Warrior threat that broadens the attacking force.
- 2 Dragonhawk, Fate's Tempest | threat | A powerful threat that gives the deck a high-impact creature.
- 2 Heartfire Hero | other | A Mouse Soldier that adds more creatures to the aggressive plan.

</details>

### 25. use no card of an owned precon

Format: Commander. Theme: Avengers superheroes. Pool: owned_first. Shortlist: 195 names.

Commander: Captain America, Super-Soldier.

Grade: baseline, score 0.29, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $179.77 the whole deck.

**Summary:** Captain America leads a mono-white superhero squad built around deploying Heroes, protecting the team, and using Equipment to turn steady attacks into pressure. The deck develops its board with S.H.I.E.L.D. agents and Avengers-adjacent heroes, keeps cards flowing through its creature and Equipment pieces, and closes with major superhero threats such as The Sentry, Golden Guardian and Origin of Spider-Man. It gives up some flexibility in exchange for a focused white mana base and a creature-forward plan that can be vulnerable when the board is repeatedly cleared.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The deck makes less mana on turn four than the norm, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core is a coherent mono-white Hero/Equipment go-wide attack plan with Puresteel Paladin, Skullclamp, Sword of the Animist and protection effects, but three unconditional board wipes (Austere Command, Fumigate, Vanquish the Horde) plus Angel of Serenity pull against a creature-forward board that the deck itself is trying to build and protect.
- PLAN theme_fit=partly: The commander and a solid Marvel core (Vision, Viv Vision, Coulson, Maria Hill, Sharon Carter, Mockingbird, Luke Cage, Okoye, Invisible Woman, The Sentry, Captain Mar-Vell) answer the Avengers superhero ask in mono-white at a reasonable bracket-3 power, but a sizable share of the list is off-theme filler from Final Fantasy, LOTR and Avatar (Buster Sword, Blitzball, Lembas, Crown of Gondor, Champions of Minas Tirith, Bender's Waterskin), and picking the face-name Captain America risks the stated no-precon-cards limit.
- PLAN useful_as_built=yes: 35 Plains plus nine cheap rocks and land-fetch pieces support a low mono-white curve, and the deck has ample removal, protection, card draw and enough evasive/equipped threats plus Angel of Serenity to actually close a game out of the box.
- PLAN summary_honest=partly: Most claims check out (Heroes, S.H.I.E.L.D. agents, Equipment, card flow, white-only mana), but it warns the deck is vulnerable to repeated board clears while omitting that the list runs three sweepers of its own, and it bills a two-mana saga (Origin of Spider-Man) as a major closing threat.
- [INFO] `curve_summary`: average mana value 2.95 over 64 nonland cards
- [INFO] `mana_pass`: the builder moved 1 card of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 35 Plains | land | Provides reliable white mana while keeping the mana base consistent.
- 1 Arcane Signet | ramp | Provides efficient color fixing and mana acceleration.
- 1 Bender's Waterskin | ramp | Adds another early mana source.
- 1 Blitzball | ramp | Helps accelerate the deck's development.
- 1 Chromatic Lantern | ramp | Provides mana acceleration and supports smooth casting.
- 1 Commander's Sphere | ramp | Accelerates mana and remains useful later.
- 1 Explorer's Scope | ramp | Uses the creature-heavy plan to develop mana.
- 1 Inherited Envelope | ramp | Provides another ramp piece for reaching the midgame.
- 1 Sword of the Animist | ramp | Pairs with attacking Heroes to build mana over time.
- 1 Thought Vessel | ramp | Provides efficient mana acceleration.
- 1 Wayfarer's Bauble | ramp | Finds an additional basic land early.
- 1 Agent 13, Sharon Carter | draw | A Hero that contributes card draw to the team plan.
- 1 Agent Maria Hill | draw | Adds another Hero while providing card draw.
- 1 Buster Sword | draw | Provides card draw from an Equipment slot.
- 1 Champions of Minas Tirith | draw | Adds a creature that helps keep cards flowing.
- 1 Crown of Gondor | draw | An Equipment-based source of card draw.
- 1 Hero in Training | draw | A Hero that supports the deck's card advantage.
- 1 Idol of Oblivion | draw | Provides a compact card-draw artifact.
- 1 Lembas | draw | Provides a low-cost source of card draw.
- 1 Mask of Memory | draw | Rewards combat with continued card selection.
- 1 Puresteel Paladin | draw | Supports the Equipment package while drawing cards.
- 1 Skullclamp | draw | Turns small creatures into additional cards.
- 1 The Vision | draw | A Robot Hero that contributes card draw.
- 1 Viv Vision, Teen Synthezoid | draw | Adds another Heroic body and card draw.
- 1 Weapons Vendor | draw | Supports the Equipment theme while drawing cards.
- 1 Bastion Protector | interaction | Protects the commander and reinforces the battlefield.
- 1 Boromir, Warden of the Tower | interaction | Provides a creature-based interaction piece.
- 1 Champion's Helm | interaction | Helps keep the commander protected.
- 1 Clever Concealment | interaction | Protects the assembled team from opposing answers.
- 1 Darksteel Plate | interaction | Provides durable protection for a key creature.
- 1 Frontline Medic | interaction | A battlefield protection piece for combat-focused turns.
- 1 Lightning Greaves | interaction | Protects an important Hero immediately.
- 1 Patriot, Shield Wielder | interaction | A Hero that adds to the deck's protective plan.
- 1 Reprieve | interaction | Offers flexible instant-speed disruption.
- 1 Swiftfoot Boots | interaction | Gives a key creature fast protection.
- 1 Unbreakable Formation | interaction | Protects the team while supporting an aggressive board.
- 1 Banishing Light | removal | Provides broad permanent-based removal.
- 1 Crib Swap | removal | Offers flexible creature removal.
- 1 Dispatch | removal | Provides efficient targeted removal.
- 1 Generous Gift | removal | Answers a wide range of problematic permanents.
- 1 Get Lost | removal | Provides efficient removal for troublesome targets.
- 1 March of Otherworldly Light | removal | Scales into an answer for key opposing permanents.
- 1 Path to Redemption | removal | Adds another answer to opposing threats.
- 1 Stroke of Midnight | removal | Offers instant-speed removal for many permanent types.
- 1 Swords to Plowshares | removal | Provides an efficient answer to opposing creatures.
- 1 Agent Phil Coulson | synergy | A Hero that strengthens the S.H.I.E.L.D. and Avengers-style team.
- 1 Agent of Atlas | synergy | Adds another Heroic team member for the creature plan.
- 1 Agents of S.H.I.E.L.D. | synergy | Reinforces the deck's S.H.I.E.L.D. Hero theme.
- 1 Colleen Wing, Street Samurai | synergy | A Hero that supports the deck's coordinated creature plan.
- 1 Dancer's Chakrams | synergy | An Equipment that supports the combat-focused Hero shell.
- 1 Door of Destinies | synergy | Builds up the shared creature plan over a longer game.
- 1 Heirloom Blade | synergy | An Equipment that rewards the deck's creature theme.
- 1 Captain Mar-Vell, Space-Born | threat | A superhero threat that advances the Avengers roster.
- 1 Invisible Woman, Sue Storm | threat | A superhero threat that adds another powerful Hero to the board.
- 1 Luke Cage, Power Man | threat | A superhero threat for the deck's creature pressure.
- 1 Mockingbird, Ace Agent | threat | A Hero threat that fits the S.H.I.E.L.D. side of the roster.
- 1 Okoye, Dora Milaje Leader | threat | A Hero threat that contributes to the combat plan.
- 1 Roaming Throne | threat | Adds a resilient threat to the creature-heavy battlefield.
- 1 Angel of Serenity | wincon | A finisher that gives the deck a high-impact closing threat.
- 1 Origin of Spider-Man | wincon | A finisher that provides a superhero-themed route to close games.
- 1 The Sentry, Golden Guardian | wincon | A superhero finisher and one of the deck's largest closing threats.
- 1 Austere Command | wipe | Provides a flexible reset when opponents get ahead.
- 1 Fumigate | wipe | Clears crowded creature boards.
- 1 Vanquish the Horde | wipe | Provides an efficient board reset against creature-heavy tables.
- 1 Astral Cornucopia | synergy | the mana pass added it to bring the mana base inside the power level

</details>

