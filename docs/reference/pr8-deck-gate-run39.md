# PR-8 deck gate

Run date: 2026-10-04. Card snapshot: 2026-09-04.

Verdict: FAIL. 25 of 25 decks passed every block check, 0 invented names reached the user, and 1 summaries stated a false rule of the game. All three bars are zero tolerance.

## Summary

| Measure | Value |
|---|---|
| Prompts | 25 |
| Decks returned | 25 |
| Decks with no block finding | 25 |
| Invented names that reached the user | 0 |
| Decks that needed the repair turn | 3 |
| Summaries judged (F-26) | 25 |
| Summaries that state a rule of the game | 2 |
| Summaries that state a FALSE rule | 1 |
| Judge errors | 0 |
| Case assertions missed (PR-28b) | 0 |
| Decks the plan judge read (PR-15, information) | 25 |
| Mean plan score, 0 to 1 | 0.70 |
| Plan reasons the judge left empty | 0 |
| Errors | 0 |
| Prompt version | 18 |
| Calls | 81 |
| Cost | $1.7725 |
| Time | 2040 seconds |

## Run

- Suite `decks`, run `pr8-deck-gate-run39`, on 2026-10-04, commit `2883f53`.
- Roles: generate on `gpt-6.1-sol` (openai, effort medium), judge on `claude-sonnet-5-5` (anthropic, effort medium), repair on `gpt-6.1-sol` (openai, effort medium).
- Versions: card snapshot 2026-09-04, generate prompt version 18, plan_rubric prompt version 5, summary_judge prompt version 2, precons `5.3.0+20260923`, quality_model `20260923T202806Z`.
- Calls: 81. Cost: $1.7725. Time: 2040 seconds.

## Findings by code

A block stops the deck. A warning and a note are reports.

| Code | Count |
|---|---|
| `not_owned` | 27 |
| `curve_summary` | 25 |
| `mana_pass` | 10 |
| `profile_off_band` | 9 |
| `finisher_short` | 2 |
| `outside_requested_set` | 2 |
| `summary_rules_claim` | 1 |
| `bracket_cut` | 1 |

By severity: BLOCK 0. WARN 41. INFO 36. 

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

Grade: typical, score 0.33, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $1212.20 to buy, $1212.20 the whole deck.

**Summary:** Karlov leads a lifegain deck that develops early, keeps cards flowing, and builds toward sustained battlefield pressure. Use interaction to keep dangerous opposing plays in check while Karlov and the threat package establish control. Win through combat or dedicated lifegain finishers, chief among them Angel of Destiny, Felidar Sovereign, and Test of Endurance. The deck favors a steady, interactive game over an all-in combo plan, giving up some explosive speed and remaining vulnerable when its battlefield is repeatedly cleared.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The cards pair in ways the top lists do not, and that lowers the grade. The commander is a popular one, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: Most of the list is lifegain creatures, Angels, and Clerics that fit Karlov's plan, but three board wipes, several blink and protection tricks, and stray pieces like Mox Amber and Vexing Bauble pull toward a different, more reactive game.
- PLAN theme_fit=yes: Karlov of the Ghost Council leads a white-black deck full of lifegain payoffs like Archangel of Thune, Lyra Dawnbringer, and Voice of the Blessed, which is the lifegain Commander deck the person asked for.
- PLAN useful_as_built=partly: The mana is plentiful and the creature count is high, but the list has many weak ramp and utility cards, and the wipes work against its own creatures, so it plays unevenly.
- PLAN summary_honest=partly: The summary names Angel of Destiny, Felidar Sovereign, and Test of Endurance as finishers and all three are in the list, but it never mentions the three board wipes, which work against a creature-based lifegain plan.
- [INFO] `curve_summary`: average mana value 2.92 over 63 nonland cards
- [INFO] `mana_pass`: the builder moved 2 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 1 Command Tower | land | Supports the white-black mana base.
- 1 Godless Shrine | land | Helps supply the deck's two colors.
- 1 Scrubland | land | Helps supply the deck's two colors.
- 1 Caves of Koilos | land | Supports early white-black plays.
- 1 City of Brass | land | Adds flexibility to the mana base.
- 1 Mana Confluence | land | Adds flexibility to the mana base.
- 1 Forbidden Orchard | land | Supports access to both deck colors.
- 1 Concealed Courtyard | land | Supports the deck's early development.
- 1 Vault of Champions | land | Supports the white-black mana base.
- 1 Silent Clearing | land | Provides another land slot for the two-color mana base.
- 1 Isolated Chapel | land | Helps balance access to white and black.
- 1 Shattered Sanctum | land | Supports colored mana through the middle turns.
- 1 Shineshadow Snarl | land | Adds another source for the deck's colors.
- 1 Brightclimb Pathway // Grimclimb Pathway | land | Adds flexibility to early land sequencing.
- 1 Starting Town | land | Supports the deck's colored mana needs.
- 1 Marsh Flats | land | Helps assemble the white-black mana base.
- 1 Prismatic Vista | land | Helps balance the deck's basic lands.
- 1 Radiant Fountain | land | Adds a utility land slot without crowding out colored sources.
- 11 Plains | land | Provides a dependable foundation for white spells.
- 7 Swamp | land | Provides a dependable foundation for black spells.
- 1 Sol Ring | ramp | Provides efficient early acceleration.
- 1 Arcane Signet | ramp | Supports early development and colored mana needs.
- 1 Mox Amber | ramp | Adds low-cost acceleration to the commander-focused plan.
- 1 Chrome Mox | ramp | Prioritizes faster development over conserving every card.
- 1 Thieving Varmint | ramp | Adds a creature-based ramp option.
- 1 Pristine Talisman | ramp | Provides additional acceleration for the middle turns.
- 1 Phial of Galadriel | ramp | Adds another artifact-based ramp option.
- 1 Nuka-Cola Vending Machine | ramp | Diversifies the deck's ramp package.
- 1 Crypt Ghast | ramp | Provides a higher-impact ramp piece for longer games.
- 1 Azor's Gateway // Sanctum of the Sun | ramp | Adds a developmental ramp piece alongside the faster accelerants.
- 1 Esper Sentinel | draw | Adds an early card-advantage piece.
- 1 Archivist of Oghma | draw | Provides a low-cost source of card advantage.
- 1 Enduring Innocence | draw | Supports continued card flow through creature development.
- 1 Tymna the Weaver | draw | Adds card advantage to the creature-heavy plan.
- 1 The Gaffer | draw | Provides a draw option suited to the lifegain theme.
- 1 Dawn of Hope | draw | Adds a thematic card-advantage engine.
- 1 Well of Lost Dreams | draw | Supports card advantage within the lifegain plan.
- 1 Mangara, the Diplomat | draw | Helps sustain resources in multiplayer games.
- 1 Cosmos Elixir | draw | Adds an artifact-based source of card advantage.
- 1 Exemplar of Light | draw | Combines a creature slot with the draw package.
- 1 Lunar Convocation | draw | Adds another card-advantage piece for sustained play.
- 1 Reprieve | interaction | Provides inexpensive interaction while developing the board.
- 1 Soul Partition | interaction | Adds a flexible interactive spell.
- 1 Aven Interrupter | interaction | Provides interaction in a creature slot.
- 1 Boromir, Warden of the Tower | interaction | Adds a creature-based interactive piece.
- 1 Ephemerate | interaction | Adds a low-cost tactical option for the creature package.
- 1 Restoration Magic | interaction | Provides another instant-speed tactical option.
- 1 Vexing Bauble | interaction | Adds inexpensive interaction without raising the curve.
- 1 Bilbo's Gambit | interaction | Rounds out the deck's instant-speed interaction.
- 1 Swords to Plowshares | removal | Provides an efficient removal spell.
- 1 Dismember | removal | Adds another compact removal option.
- 1 Deadly Rollick | removal | Provides a strong answer alongside the commander.
- 1 Sorin of House Markov // Sorin, Ravenous Neonate | removal | Adds a thematic removal piece.
- 1 Amalia Benavides Aguirre | removal | Provides a low-cost removal-role creature for the lifegain shell.
- 1 Ayli, Eternal Pilgrim | removal | Adds another removal-role creature suited to the theme.
- 1 Gumdrop Poisoner // Tempt with Treats | removal | Adds a thematic removal option in a creature slot.
- 1 Murderous Rider // Swift End | removal | Broadens the removal package with another creature card.
- 1 Witch of the Moors | removal | Provides a higher-impact removal piece for longer games.
- 1 Blind Obedience | synergy | Adds a thematic synergy piece and a marked finisher.
- 1 Cleric Class | synergy | Provides a low-cost foundation for the lifegain theme.
- 1 Ocelot Pride | synergy | Adds an inexpensive creature to the synergy package.
- 1 Righteous Valkyrie | synergy | Supports the lifegain-focused creature plan.
- 1 Voice of the Blessed | synergy | Adds a compact synergy creature to the battlefield plan.
- 1 Vito, Thorn of the Dusk Rose | synergy | Provides another dedicated lifegain synergy piece.
- 1 Tithe Drinker | synergy | Adds an early synergy creature and another marked finisher.
- 1 Rhox Faithmender | threat | Adds a thematic threat to the lifegain shell.
- 1 Attended Healer | threat | Provides another battlefield threat suited to the theme.
- 1 Archangel of Thune | threat | Adds a substantial threat for the middle and late game.
- 1 Elenda, Saint of Dusk | threat | Strengthens the deck's thematic threat package.
- 1 Liesa, Forgotten Archangel | threat | Provides another substantial creature threat.
- 1 Lyra Dawnbringer | threat | Adds a marked finisher to the creature-based attack plan.
- 1 Gisela, the Broken Blade | threat | Adds pressure without reaching the top of the curve.
- 1 Enduring Tenacity | threat | Provides a thematic threat alongside the synergy package.
- 1 Blood Baron of Vizkopa | threat | Adds another substantial creature for sustained pressure.
- 1 Cliffhaven Vampire | threat | Adds a threat suited to the lifegain theme.
- 1 Celestine, the Living Saint | threat | Provides another late-game creature threat.
- 1 Ajani Goldmane | threat | Diversifies the threat package with a planeswalker.
- 1 Angel of Destiny | wincon | Provides a dedicated finisher for the lifegain strategy.
- 1 Felidar Sovereign | wincon | Adds a creature-based route to closing the game.
- 1 Test of Endurance | wincon | Adds a noncreature win condition suited to the theme.
- 1 Damn | wipe | Provides a board-reset option when opposing threats accumulate.
- 1 Toxic Deluge | wipe | Adds a lower-cost reset for difficult battlefields.
- 1 Fumigate | wipe | Rounds out the board-reset package.

</details>

### 2. aristocrats Commander, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 196 names.

Commander: Denethor, Ruling Steward.

Grade: baseline, score 0.35, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $2.25 to buy, $173.67 the whole deck.

**Summary:** Denethor leads a patient aristocrats deck that turns expendable creatures into life loss and replenishes the board with end-step tokens. Develop mana and card draw first, then use sacrifice synergies, protective interaction, and removal to keep the engine working through a long game. Steady drains and larger finishers provide complementary ways to close, with combat offering another route when opponents leave an opening. The deck gives up explosive speed and a dense dedicated-drain package in favor of resilient utility creatures and broad answers.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The mana base serves one color far better than another, and that lowers the grade. The commander does not place in cEDH events, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: Sacrifice outlets, drain effects and tokens exist, but the list also carries three board wipes, many protective instants and a pile of generic ramp that pull away from a go-wide aristocrats board.
- PLAN theme_fit=partly: The deck is the right colors for Denethor at a plausible bracket 3 and has some sacrifice pieces, but the request for aristocrats from the owner's library first is only loosely served by a list heavy on removal, wipes and ramp.
- PLAN useful_as_built=partly: With 17 Plains against 9 Swamp plus plenty of fixing and ramp it will cast spells, but the black-heavy cards like Witch-king and Archfiend strain the white-skewed mana, and the win conditions are thin.
- PLAN summary_honest=partly: The summary claims 'end-step tokens' to replenish the board, but few cards in the list make tokens at end step beyond Denethor himself, and it admits the thin drain package while still calling the deck an aristocrats engine.
- [INFO] `curve_summary`: average mana value 2.73 over 62 nonland cards
- [INFO] `mana_pass`: the builder moved 4 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 17 Plains | land | Provides dependable white mana for Denethor and the supporting creatures.
- 9 Swamp | land | Provides dependable black mana for sacrifice engines and finishers.
- 1 Command Tower | land | Fixes both of the deck's colors without slowing the opening turns.
- 1 City of Brass | land | Provides either color when the opening hand needs flexible mana.
- 1 Exotic Orchard | land | Adds flexible fixing in multiplayer games.
- 1 Marsh Flats | land | Finds the basic land needed for the next play.
- 1 Fabled Passage | land | Finds either basic land and improves color consistency.
- 1 Evolving Wilds | land | Trades early speed for reliable access to either color.
- 1 Grand Coliseum | land | Supplies flexible colored mana after its slower opening turn.
- 1 Path of Ancestry | land | Fixes both colors and complements the deck's Human creatures.
- 1 Thriving Moor | land | Pairs black mana with white fixing.
- 1 Spire of Industry | land | Uses the artifact package to support colored mana.
- 1 Opal Palace | land | Provides additional fixing for casting Denethor.
- 1 Sol Ring | ramp | Accelerates setup and leaves more mana available for sacrifice activations.
- 1 Springleaf Drum | ramp | Turns a spare creature into inexpensive mana fixing.
- 1 Arcane Signet | ramp | Provides efficient acceleration in either color.
- 1 Fellwar Stone | ramp | Adds inexpensive acceleration with multiplayer fixing potential.
- 1 Thought Vessel | ramp | Accelerates development and supports larger hands.
- 1 Wayfarer's Bauble | ramp | Builds the mana base with a basic land.
- 1 Sword of the Animist | ramp | Turns attacks by expendable creatures into lasting mana development.
- 1 Lotho, Corrupt Shirriff | ramp | Adds creature-based acceleration to the deck's resource engine.
- 1 Chromatic Lantern | ramp | Smooths colored mana for demanding spells and multiple plays.
- 1 Relic of Legends | ramp | Pairs mana acceleration with the deck's numerous legendary creatures.
- 1 Gorbag of Minas Morgul | ramp | Adds another creature-based ramp option to the sacrifice plan.
- 1 Skullclamp | draw | Converts expendable creatures into a steady supply of cards.
- 1 Idol of Oblivion | draw | Provides repeatable draw alongside Denethor's token production.
- 1 Night's Whisper | draw | Refills the hand efficiently during setup.
- 1 Nasty End | draw | Turns a creature that is about to die into additional cards.
- 1 Call of the Ring | draw | Supplies ongoing card advantage for longer games.
- 1 Folk Hero | draw | Rewards the Human-heavy supporting creature package.
- 1 Inspiring Overseer | draw | Replaces itself while leaving a creature for combat or sacrifice.
- 1 Wall of Omens | draw | Provides an early blocker and a replacement card.
- 1 Lembas | draw | Adds inexpensive card flow to the artifact package.
- 1 Mask of Memory | draw | Turns successful attacks into hand selection and card advantage.
- 1 The Sackville-Bagginses | draw | Adds a legendary creature-based source of cards.
- 1 Ahriman | draw | Adds another draw creature to support the attrition plan.
- 1 Boromir, Warden of the Tower | interaction | Provides interaction on a creature that supports the sacrifice strategy.
- 1 Reprieve | interaction | Buys tempo against a pivotal opposing spell.
- 1 Clever Concealment | interaction | Preserves the developed board through disruptive turns.
- 1 Unbreakable Formation | interaction | Protects the creature board and supports a decisive attack.
- 1 Take Up the Shield | interaction | Keeps an important creature alive through combat or removal.
- 1 Orcish Medicine | interaction | Provides inexpensive interaction for protecting a useful creature.
- 1 Slip On the Ring | interaction | Protects a creature while reusing a useful enters-the-battlefield effect.
- 1 Duty Beyond Death | interaction | Adds another way to preserve the creatures that keep the engine running.
- 1 Swords to Plowshares | removal | Answers a dangerous creature for a minimal mana investment.
- 1 Generous Gift | removal | Answers a problematic permanent that narrower removal cannot handle.
- 1 Bitter Triumph | removal | Provides flexible removal at instant speed.
- 1 Infernal Grasp | removal | Removes a threatening creature efficiently.
- 1 Fatal Push | removal | Provides cheap creature removal that suits a deck with frequent sacrifices.
- 1 Fiend Hunter | removal | Combines creature removal with a body for the supporting engines.
- 1 Palace Jailer | removal | Adds creature-based removal and pressure for the longer game.
- 1 Witch-king of Angmar | removal | Combines removal pressure with a substantial finishing threat.
- 1 Get Lost | removal | Answers several important permanent types efficiently.
- 1 Orcish Bowmasters | removal | Adds efficient creature-based interaction to the attrition plan.
- 1 Gollum the Abandoned | synergy | Adds a sacrifice-oriented creature to the core engine.
- 1 Gollum, Patient Plotter | synergy | Provides reusable creature material for sacrifice turns.
- 1 Gríma Wormtongue | synergy | Adds another creature that supports the sacrifice plan.
- 1 Nimble Hobbit | synergy | Supports attacks by the deck's smaller creatures.
- 1 Arcade Cabinet | synergy | Adds a supporting artifact to the deck's synergy package.
- 1 Phantom Train | synergy | Adds an artifact-based payoff to the supporting engine.
- 1 Heirloom Auntie | synergy | Adds another creature-based synergy piece for the longer game.
- 1 Bill the Pony | threat | Develops the creature board with a useful legendary threat.
- 1 Bastion Protector | threat | Adds a combat body while helping keep Denethor on the table.
- 1 Bronze Guardian | threat | Provides a substantial combat threat alongside the artifact package.
- 1 Gilraen, Dúnedain Protector | threat | Adds a legendary body with utility for important creatures.
- 1 Zack Fair | threat | Provides an early body with protective utility.
- 1 Exemplar of Light | threat | Combines aerial pressure with a supporting draw role.
- 1 Massacre Girl, Known Killer | threat | Combines combat pressure with card advantage for the attrition plan.
- 1 Cirith Ungol Patrol | threat | Provides a substantial body with additional card-flow utility.
- 1 Grave Venerations | wincon | Provides a dedicated finisher for the sacrifice-centered game plan.
- 1 Al Bhed Salvagers | wincon | Adds a finishing payoff alongside the artifact package.
- 1 Archfiend of Ifnir | wincon | Provides another substantial finisher for games that reach the late turns.
- 1 Austere Command | wipe | Provides a flexible reset when opposing boards outgrow the engine.
- 1 Dusk // Dawn | wipe | Offers a board reset suited to the deck's many small creatures.
- 1 The Battle of Bywater | wipe | Controls larger creature boards while supporting the small-creature plan.

</details>

### 3. artifacts Commander, bracket 4

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 318 names.

Commander: Urza, Lord High Artificer.

Grade: good, score 0.75, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $6895.19 to buy, $6895.19 the whole deck.

**Summary:** Your commander turns a wide artifact board into mana, letting you develop threats while keeping interaction ready. Cheap acceleration, draw engines, and tutors support explosive turns, with artifact-heavy combat and dedicated finishers providing multiple routes to victory. The deck gives up some resilience to artifact hate and some opening-hand consistency to fit both a strong combo plan and a substantial battlefield presence.

The quality model grades this deck good against the top lists of the format: the cards are ones the top lists of the format play, and that raises the grade. The mana base serves the colors evenly, and that raises the grade. The commander places in cEDH events, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The mono-blue artifact shell with Moxen, counterspells, and Urza's mana engine pulls together, but the 23 Islands beside Mishra's Workshop and colorless utility lands, the bounce-and-Aura removal, and the mix of free spells and clunky 7-drops pull in different directions.
- PLAN theme_fit=yes: It is a mono-blue Urza, Lord High Artificer deck packed with artifacts, Moxen, free counterspells, and Mana Vault-style acceleration, which matches a high-power bracket 4 artifact request.
- PLAN useful_as_built=partly: The mana and interaction suite is functional, but 23 Islands plus 11 nonbasics is a heavy land count with 12 mana rocks, and several weak cards like Myr Enforcer, Sojourner's Companion, Amphibian Downpour, and Submerge dilute a deck with few clear ways to win.
- PLAN summary_honest=partly: The summary claims 'dedicated finishers' and a 'strong combo plan', but the list has only Thassa's Oracle, Mechanized Production, and Power Artifact with Urza as real win routes, and the combo plan lacks the tutors the summary promises beyond Transmute Artifact and Whir of Invention.
- [INFO] `curve_summary`: average mana value 2.89 over 66 nonland cards

<details><summary>The deck list</summary>

- 23 Island | land | Provides reliable untapped blue mana for early plays and demanding blue costs.
- 1 Academy Ruins | land | Recovers important artifacts after removal or sacrifice.
- 1 Inventors' Fair | land | Finds a key artifact once your board is established.
- 1 Urza's Saga | land | Produces artifact threats and finds a cheap utility artifact.
- 1 Mishra's Workshop | land | Accelerates artifact deployment without using a spell slot.
- 1 Seat of the Synod | land | Supplies blue mana while increasing your artifact count.
- 1 Otawara, Soaring City | land | Provides untapped blue mana with a useful late-game interaction option.
- 1 Mystic Sanctuary | land | Returns an important instant or sorcery for another use.
- 1 Command Tower | land | Provides dependable blue mana without an entry delay.
- 1 Spire of Industry | land | Supplies blue mana in a deck built to maintain artifacts.
- 1 Archway of Innovation | land | Provides blue mana and helps cast expensive spells through your artifact board.
- 1 Sol Ring | ramp | Accelerates your commander and the deck's more expensive artifacts.
- 1 Mana Vault | ramp | Provides an explosive burst of mana for early development.
- 1 Chrome Mox | ramp | Turns a spare colored card into immediate acceleration.
- 1 Mox Diamond | ramp | Converts an extra land into fast blue mana.
- 1 Mox Opal | ramp | Rewards a dense artifact board with free mana production.
- 1 Mox Amber | ramp | Provides free blue mana once your commander is established.
- 1 Lotus Petal | ramp | Supplies a cheap burst of mana to accelerate an important turn.
- 1 Arcane Signet | ramp | Adds reliable blue acceleration and another artifact to the board.
- 1 Grim Monolith | ramp | Accelerates large plays and combines with Power Artifact for unlimited mana.
- 1 Basalt Monolith | ramp | Provides acceleration and a second unlimited-mana partner for Power Artifact.
- 1 Metalworker | ramp | Turns an artifact-heavy hand into substantial mana production.
- 1 Tezzeret the Seeker | ramp | Finds important artifacts and untaps mana-producing artifacts.
- 1 Fellwar Stone | ramp | Adds inexpensive acceleration while contributing to artifact synergies.
- 1 Mystic Remora | draw | Refills your hand while opponents develop with noncreature spells.
- 1 Rhystic Study | draw | Provides sustained card advantage throughout the game.
- 1 Sensei's Divining Top | draw | Improves draw quality and helps locate interaction or combo pieces.
- 1 Thoughtcast | draw | Converts a developed artifact board into inexpensive card draw.
- 1 Thought Monitor | draw | Refills your hand while adding an evasive artifact body.
- 1 Sai, Master Thopterist | draw | Turns artifact spells into tokens and converts spare artifacts into cards.
- 1 Forensic Gadgeteer | draw | Produces Clues from artifact spells and discounts artifact activations.
- 1 Consecrated Sphinx | draw | Provides a powerful sustained card-advantage engine.
- 1 Thirst for Knowledge | draw | Finds fresh cards while filtering excess artifacts or lands.
- 1 Windfall | draw | Replenishes your hand after a fast deployment of mana and artifacts.
- 1 Faerie Mastermind | draw | Adds an inexpensive source of cards against opposing draw engines.
- 1 Gitaxian Probe | draw | Checks an opponent's hand before committing to a critical turn.
- 1 Force of Will | interaction | Protects a decisive turn even when your mana is committed elsewhere.
- 1 Force of Negation | interaction | Stops dangerous noncreature spells without requiring open mana on opposing turns.
- 1 Fierce Guardianship | interaction | Protects your board and combo once your commander is present.
- 1 Pact of Negation | interaction | Provides immediate protection when attempting to end the game.
- 1 Mental Misstep | interaction | Answers critical one-mana spells efficiently.
- 1 Flusterstorm | interaction | Helps win fights over instants and sorceries.
- 1 Swan Song | interaction | Answers important enchantments, instants, and sorceries for little mana.
- 1 An Offer You Can't Refuse | interaction | Stops a critical noncreature spell at a low mana cost.
- 1 Daze | interaction | Provides early protection while allowing aggressive mana deployment.
- 1 Metallic Rebuke | interaction | Uses your artifacts to make countermagic easier to hold up.
- 1 Hullbreaker Horror | interaction | Turns subsequent spells into repeatable disruption and supports mana loops with cheap artifacts.
- 1 Chain of Vapor | removal | Clears an obstructing nonland permanent efficiently.
- 1 Into the Flood Maw | removal | Provides inexpensive bounce for a troublesome opposing permanent.
- 1 Snap | removal | Removes a creature temporarily while restoring access to land mana.
- 1 Snapback | removal | Provides creature bounce when your mana is otherwise committed.
- 1 Submerge | removal | Offers efficient creature disruption against opponents with Forests.
- 1 Amphibian Downpour | removal | Neutralizes creature abilities and scales during spell-heavy turns.
- 1 Walking Ballista | removal | Removes small creatures and converts unlimited mana into lethal damage.
- 1 Volatile Stormdrake | removal | Takes an opposing creature while adding an evasive body.
- 1 Into Thin Air | removal | Uses artifact density to reduce the cost of bouncing a troublesome permanent.
- 1 Arcum Dagsson | removal | Turns expendable artifact creatures into important noncreature artifacts.
- 1 Power Artifact | synergy | Combines with either monolith to generate unlimited colorless mana.
- 1 Transmute Artifact | synergy | Converts an expendable artifact into a needed engine or combo piece.
- 1 Whir of Invention | synergy | Uses spare artifacts to find a key artifact at instant speed.
- 1 Kappa Cannoneer | threat | Provides a resilient threat that grows as artifacts enter.
- 1 Cyberdrive Awakener | threat | Turns a developed artifact board into a flying combat finish.
- 1 Phyrexian Metamorph | threat | Copies the most valuable artifact or creature available.
- 1 Kuldotha Forgemaster | threat | Converts expendable artifacts into a powerful artifact of your choice.
- 1 Lodestone Golem | threat | Pressures opponents while slowing their nonartifact development.
- 1 Traxos, Scourge of Kroog | threat | Provides a large trampling body that your artifact spells repeatedly untap.
- 1 Filigree Attendant | threat | Turns artifact density into substantial evasive pressure.
- 1 Frogmite | threat | Adds an inexpensive artifact body as your board develops.
- 1 Myr Enforcer | threat | Uses affinity to add a substantial body without consuming much mana.
- 1 Sojourner's Companion | threat | Provides an affinity threat with a land-finding option when needed.
- 1 Master Transmuter | threat | Exchanges cheap artifacts for larger artifacts from your hand.
- 1 Karn, Scion of Urza | threat | Produces growing artifact threats and supplies additional resources.
- 1 Thassa's Oracle | wincon | Finishes after unlimited mana lets your commander exhaust your library.
- 1 Mirrodin Besieged | wincon | Provides an alternate victory route from an artifact-filled graveyard.
- 1 Mechanized Production | wincon | Adds an alternate victory route built around accumulating matching artifacts.
- 1 Cyclonic Rift | wipe | Clears opposing boards while preserving your developed battlefield.
- 1 Desynchronization | wipe | Resets nonartifact boards while leaving most of your engine intact.

</details>

### 4. dinosaur tribal, bracket 2

Format: Commander. Theme: dinosaurs. Pool: any_card. Shortlist: 307 names.

Commander: Gishath, Sun's Avatar.

Grade: bad, score 0.03, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $1590.30 to buy, $1590.30 the whole deck.

**Summary:** Develop your mana early, build a herd of Dinosaurs, and use your commander's attacks to expand the battlefield. Large Dinosaur threats and dedicated finishers, chief among them Bonehoard Dracosaur and Dinosaurs on a Spaceship, support a combat-focused win. Card draw keeps the creatures coming, while removal, protective interaction, and board resets help you recover from opposition. The deck favors a straightforward creature plan over combos, giving up explosive speed and some flexibility for a slower, Dinosaur-heavy game.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The mana base serves one color far better than another, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Ramp, Dinosaur creatures, card draw, and a few protective spells all feed one plan of building a Dinosaur board and attacking with Gishath.
- PLAN theme_fit=yes: It is a Gishath Dinosaur Commander deck with a large Dinosaur count, a creature-focused plan, and mostly low-power cards that fit bracket 2.
- PLAN useful_as_built=partly: The deck has plenty of ramp, draw, and Dinosaurs, but the three-color mana base leans heavily on green with only one Mountain and one Plains among the basics, and the many nonbasic lands and RR or WW costs make it hard for a new player to cast things reliably.
- PLAN summary_honest=partly: The summary's claim of a herd of Dinosaurs with Bonehoard Dracosaur and Dinosaurs on a Spaceship as finishers matches the list, but it does not mention the heavy green lean, the many painful or tricky lands, or that Vandalblast and Blasphemous Act are awkward in a creature deck.
- [INFO] `curve_summary`: average mana value 3.46 over 61 nonland cards
- [INFO] `mana_pass`: the builder moved 2 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 1 Command Tower | land | Provides flexible color fixing for the three-color mana base.
- 1 Path of Ancestry | land | Adds color fixing suited to the Dinosaur theme.
- 1 City of Brass | land | Helps supply whichever color the opening hand needs.
- 1 Mana Confluence | land | Provides flexible access to the deck's colors.
- 1 Forbidden Orchard | land | Adds another flexible source for three-color spells.
- 1 Cavern of Souls | land | Supports the mana needs of the Dinosaur-heavy creature suite.
- 1 Secluded Courtyard | land | Helps cast Dinosaurs across all three colors.
- 1 Unclaimed Territory | land | Adds creature-focused fixing for the Dinosaur plan.
- 1 Battlefield Forge | land | Supports the red and white portions of the deck.
- 1 Brushland | land | Supports green development and white support spells.
- 1 Karplusan Forest | land | Supports green ramp and red Dinosaurs.
- 1 Horizon Canopy | land | Adds another green-white source.
- 1 Sunbaked Canyon | land | Adds another red-white source.
- 1 Grove of the Burnwillows | land | Strengthens access to red and green mana.
- 1 Savannah | land | Provides green-white fixing.
- 1 Plateau | land | Provides red-white fixing.
- 1 Taiga | land | Provides red-green fixing.
- 1 Temple Garden | land | Helps balance green and white requirements.
- 1 Sacred Foundry | land | Helps balance red and white requirements.
- 1 Stomping Ground | land | Helps balance red and green requirements.
- 1 Sunpetal Grove | land | Adds green-white fixing to support development.
- 1 Clifftop Retreat | land | Adds red-white fixing for creatures and interaction.
- 1 Rootbound Crag | land | Adds red-green fixing for the creature-heavy plan.
- 1 Canopy Vista | land | Adds another green-white land to the mana base.
- 1 Cinder Glade | land | Adds another red-green land to the mana base.
- 1 Branchloft Pathway // Boulderloft Pathway | land | Offers a choice between green and white development.
- 1 Cragcrown Pathway // Timbercrown Pathway | land | Offers a choice between red and green development.
- 1 Needleverge Pathway // Pillarverge Pathway | land | Offers a choice between red and white development.
- 8 Forest | land | Gives the deck a dependable foundation for its green spells.
- 1 Mountain | land | Provides basic red sources for Dinosaurs and support spells.
- 1 Plains | land | Provides basic white sources for creatures and interaction.
- 1 Sol Ring | ramp | Helps reach the deck's expensive creatures sooner.
- 1 Birds of Paradise | ramp | Provides early creature-based mana support.
- 1 Arcane Signet | ramp | Adds inexpensive mana development and color support.
- 1 Fellwar Stone | ramp | Adds another inexpensive ramp option.
- 1 Drover of the Mighty | ramp | Supports mana development within the creature-focused plan.
- 1 Pillar of Origins | ramp | Provides ramp suited to a deck centered on one creature type.
- 1 Patchwork Banner | ramp | Adds mana support that fits the Dinosaur theme.
- 1 Hulking Raptor | ramp | Combines the Dinosaur creature type with mana development.
- 1 Atzocan Seer | ramp | Adds creature-based ramp to help reach the commander.
- 1 Fountain of Ichor | ramp | Provides another mana source for the deck's larger plays.
- 1 Garruk's Uprising | draw | Supports card flow for the creature-heavy strategy.
- 1 Beast Whisperer | draw | Helps sustain a steady stream of creatures.
- 1 Curious Altisaur | draw | Adds card support while preserving Dinosaur density.
- 1 Earthshaker Dreadmaw | draw | Fills a draw slot with a Dinosaur creature.
- 1 Ripjaw Raptor | draw | Provides Dinosaur-based card support.
- 1 Runic Armasaur | draw | Adds another Dinosaur to the card-advantage package.
- 1 Return of the Wildspeaker | draw | Helps replenish resources during the creature-based game plan.
- 1 For the Ancestors | draw | Supports continued access to cards for the Dinosaur strategy.
- 1 Vanquisher's Banner | draw | Provides card support that matches the shared creature type.
- 1 Faithless Looting | draw | Adds inexpensive card selection to smooth development.
- 1 Thrill of Possibility | draw | Provides another low-cost way to improve card flow.
- 1 Heroic Intervention | interaction | Helps preserve the board when opponents disrupt it.
- 1 Boros Charm | interaction | Provides flexible interaction and an additional finishing option.
- 1 And They Shall Know No Fear | interaction | Adds interaction suited to a shared-creature-type strategy.
- 1 Flawless Maneuver | interaction | Helps keep the creature plan intact through disruption.
- 1 Akroma's Will | interaction | Offers interaction that also supports a decisive finish.
- 1 Legolas's Quick Reflexes | interaction | Adds inexpensive interaction alongside the larger creatures.
- 1 Swords to Plowshares | removal | Provides inexpensive removal for a troublesome creature.
- 1 Savage Stomp | removal | Adds removal suited to the Dinosaur plan.
- 1 Triumphant Chomp | removal | Provides another inexpensive removal spell.
- 1 Thrashing Brontodon | removal | Fills a removal slot while adding a Dinosaur.
- 1 Tranquil Frillback | removal | Adds Dinosaur-based removal to the support package.
- 1 Bronzebeak Foragers | removal | Combines a Dinosaur body with a removal role.
- 1 Trumpeting Carnosaur | removal | Adds a substantial Dinosaur to the removal package.
- 1 Commune with Dinosaurs | synergy | Provides inexpensive support for the Dinosaur-focused plan.
- 1 Kinjalli's Caller | synergy | Adds an early support creature for the Dinosaur strategy.
- 1 Otepec Huntmaster | synergy | Supports the deck's Dinosaur development.
- 1 Herald's Horn | synergy | Provides support built around the deck's shared creature type.
- 1 Armored Kincaller | synergy | Adds a Dinosaur to the lower end of the curve.
- 1 Belligerent Yearling | synergy | Provides an early Dinosaur for the creature-focused plan.
- 1 Raptor Hatchling | synergy | Adds another inexpensive Dinosaur to establish the herd.
- 1 Kinjalli's Sunwing | synergy | Supports the Dinosaur plan without another expensive threat.
- 1 Sunfrill Imitator | synergy | Adds Dinosaur-focused support to the creature suite.
- 1 Dinosaur Stampede | synergy | Reinforces the deck's combat-oriented Dinosaur strategy.
- 1 Carnage Tyrant | threat | Provides a substantial Dinosaur threat for combat.
- 1 Regisaur Alpha | threat | Adds a Dinosaur threat to the attacking force.
- 1 Crested Herdcaller | threat | Strengthens the deck's creature-based pressure.
- 1 Quartzwood Crasher | threat | Adds another substantial Dinosaur threat.
- 1 Rampaging Ceratops | threat | Provides a Dinosaur threat marked as a finisher.
- 1 Shifting Ceratops | threat | Adds a comparatively modest-cost Dinosaur finisher.
- 1 Pantlaza, Sun-Favored | threat | Adds another Dinosaur threat to the main creature plan.
- 1 Palani's Hatcher | threat | Strengthens the Dinosaur threat package.
- 1 Colossal Dreadmaw | threat | Provides a straightforward Dinosaur combat threat.
- 1 Thundering Spineback | threat | Adds a late-game Dinosaur threat.
- 1 Goring Ceratops | threat | Provides another Dinosaur finisher for the combat plan.
- 1 Bellowing Aegisaur | threat | Adds another Dinosaur to maintain creature pressure.
- 1 Bonehoard Dracosaur | wincon | Provides a Dinosaur finisher for closing games.
- 1 Dinosaurs on a Spaceship | wincon | Adds a second dedicated Dinosaur finisher.
- 1 Blasphemous Act | wipe | Provides a reset when the battlefield becomes unfavorable.
- 1 Harsh Mercy | wipe | Adds a board reset suited to a shared-creature-type deck.
- 1 Vandalblast | wipe | Provides a wipe option against artifact-heavy boards.

</details>

### 5. blink Commander, owned first

Format: Commander. Theme: blink. Pool: owned_first. Shortlist: 183 names.

Commander: Gilraen, Dúnedain Protector.

Grade: typical, score 0.38, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $2.61 to buy, $85.92 the whole deck.

**Summary:** Gilraen, Dúnedain Protector leads a creature-heavy blink deck that builds value through draw creatures, removal creatures, and protective support. Develop mana early, establish a useful creature board, and keep resources available for Gilraen rather than committing everything at once. Smaller creatures provide steady pressure, while the larger closing threats, chief among them Angel of Serenity, help finish the game. The deck gives up explosive speed for a slower value game, and many of its creatures prioritize utility over raw combat power.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The mana base serves the colors evenly, and that raises the grade. The commander does not place in cEDH events, and that lowers the grade.

This deck plan aims for 2 or more finishers, the cards that win the game, and the deck holds 1.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every card is a white creature or support piece that either blinks, draws, removes, or protects, so the list reads as one value-blink plan around Gilraen.
- PLAN theme_fit=partly: It is a mono-white Commander deck full of enter-the-battlefield creatures, but Gilraen is not a blink commander and the list has few actual blink spells, and the 'from my library first' request cannot be seen in the list.
- PLAN useful_as_built=partly: The mana base of 29 Plains plus utility lands and ten rocks casts everything, and the removal and draw are solid, but the deck holds only one real finisher and several wipes that work against its own creature board.
- PLAN summary_honest=partly: The summary names Angel of Serenity as a chief closer, and it is in the list as a seven-mana top end, but it calls the deck a blink deck with larger closing threats when the list holds only one real finisher and few true blink effects beyond Flickerwisp and a few ETB creatures.
- [INFO] `curve_summary`: average mana value 2.79 over 63 nonland cards
- [WARN] `finisher_short`: the finisher count is 1, and this deck plan wants 2 or more (Angel of Serenity)

<details><summary>The deck list</summary>

- 29 Plains | land | Provides a dependable white mana base without slowing early turns.
- 1 Command Tower | land | Adds another source of white mana.
- 1 City of Brass | land | Adds a flexible white mana source.
- 1 Minas Tirith | land | Adds utility to the predominantly basic mana base.
- 1 War Room | land | Gives the mana base a useful late-game outlet.
- 1 Rogue's Passage | land | Provides a utility land for the creature-heavy finishing plan.
- 1 Scavenger Grounds | land | Adds utility without taking a nonland slot.
- 1 Windbrisk Heights | land | Complements the deck's creature-heavy construction.
- 1 Sol Ring | ramp | Accelerates deployment while leaving mana for Gilraen.
- 1 Arcane Signet | ramp | Provides inexpensive mana acceleration.
- 1 Thought Vessel | ramp | Adds another inexpensive ramp piece.
- 1 Fellwar Stone | ramp | Helps develop mana early.
- 1 Springleaf Drum | ramp | Provides low-cost acceleration for a creature-heavy deck.
- 1 Wayfarer's Bauble | ramp | Supports early mana development.
- 1 Sword of the Animist | ramp | Adds ramp suited to a deck that maintains a creature presence.
- 1 Giada, Font of Hope | ramp | Adds creature-based ramp alongside the Angel package.
- 1 Relic of Legends | ramp | Supports the commander and the deck's other legendary creatures.
- 1 Commander's Sphere | ramp | Rounds out the mana acceleration package.
- 1 Wall of Omens | draw | Supplies inexpensive creature-based draw for the blink plan.
- 1 Inspiring Overseer | draw | Combines card draw with a creature for Gilraen.
- 1 South Pole Voyager | draw | Adds another creature-based source of cards.
- 1 Faramir, Field Commander | draw | Supports card flow while contributing to the creature board.
- 1 Exemplar of Light | draw | Adds a substantial creature to the draw package.
- 1 Folk Hero | draw | Provides ongoing draw support for the creature-heavy plan.
- 1 Mask of Memory | draw | Adds card flow suited to a deck with many creatures.
- 1 Tome of Legends | draw | Supports card flow alongside the commander.
- 1 Skullclamp | draw | Provides an inexpensive draw engine.
- 1 Lembas | draw | Adds a compact source of card draw.
- 1 Champions of Minas Tirith | draw | Provides a larger creature in the draw package.
- 1 Boromir, Warden of the Tower | interaction | Adds interaction without reducing the creature count.
- 1 Reprieve | interaction | Provides inexpensive interaction while developing the board.
- 1 Clever Concealment | interaction | Helps preserve the deck's accumulated board investment.
- 1 Unbreakable Formation | interaction | Supports a creature board through opposing interaction.
- 1 Take Up the Shield | interaction | Adds a low-cost way to defend an important creature.
- 1 Airbender's Reversal | interaction | Provides another interactive option for contested turns.
- 1 Airbending Lesson | interaction | Adds inexpensive interaction to complement the blink engine.
- 1 Slip On the Ring | interaction | Fits the blink plan while supplying instant-speed interaction.
- 1 Swords to Plowshares | removal | Provides efficient removal for a problematic creature.
- 1 Generous Gift | removal | Adds flexible removal for opposing threats.
- 1 Get Lost | removal | Provides another inexpensive removal option.
- 1 Fiend Hunter | removal | Combines removal with a creature for the blink plan.
- 1 Palace Jailer | removal | Adds creature-based removal to the value engine.
- 1 Angel of Condemnation | removal | Supplies a creature-based removal tool.
- 1 Journey to Nowhere | removal | Adds inexpensive removal to keep development on pace.
- 1 March of Otherworldly Light | removal | Provides a flexible removal slot.
- 1 Crib Swap | removal | Adds another instant-speed answer to opposing creatures.
- 1 Flickerwisp | synergy | Directly supports the blink strategy.
- 1 Personify | synergy | Adds another dedicated blink support card.
- 1 Ennis, Debate Moderator | synergy | Provides a creature-based synergy piece.
- 1 Jocasta, Automaton Avenger | synergy | Adds another synergy creature to the value plan.
- 1 Gift of Immortality | synergy | Helps preserve the creatures that sustain the engine.
- 1 Swiftfoot Boots | synergy | Supports keeping Gilraen available throughout the game.
- 1 Aang's Iceberg | synergy | Adds protective support for the deck's important creatures.
- 1 Angel of Sanctions | threat | Adds a substantial creature while retaining removal utility.
- 1 Bastion Protector | threat | Contributes to the creature board while supporting Gilraen.
- 1 Frontline Medic | threat | Adds a combat-oriented creature with interactive utility.
- 1 Zack Fair | threat | Adds an early creature with protective utility.
- 1 Hakoda, Selfless Commander | threat | Contributes a creature while retaining protective utility.
- 1 Weapons Vendor | threat | Adds an inexpensive creature with draw utility.
- 1 Elite Interceptor // Rejoinder | threat | Adds a creature-oriented slot without abandoning card flow.
- 1 Joined Researchers // Secret Rendezvous | threat | Supports the creature plan while retaining a draw option.
- 1 Stiltzkin, Moogle Merchant | threat | Adds an inexpensive creature with card-flow utility.
- 1 Puresteel Paladin | threat | Contributes a creature alongside the deck's equipment and draw package.
- 1 Curious Farm Animals | threat | Adds another creature while retaining removal utility.
- 1 The Vision | threat | Adds a larger creature without giving up draw utility.
- 1 Angel of Serenity | wincon | Provides the deck's explicitly identified finisher.
- 1 Bronze Guardian | wincon | Serves as a closing creature while retaining protective utility.
- 1 Summon: Primal Garuda | wincon | Adds a substantial closing creature with removal utility.
- 1 Austere Command | wipe | Provides a broad reset when targeted removal is insufficient.
- 1 Dusk // Dawn | wipe | Adds another board-reset option to the creature-value plan.
- 1 The Battle of Bywater | wipe | Rounds out the board-wipe package.

</details>

### 6. Modern tempo, tournament

Format: Modern. Theme: tempo. Pool: any_card. Shortlist: 182 names.

Grade: typical, score 0.54, model 20260923T202806Z.

Cards: 60 main, 15 sideboard. Repair turn: no. Block findings: 0.

Cost: $565.24 to buy, $565.24 the whole deck.

**Summary:** This blue-red tempo deck establishes early creature pressure, then uses card selection, removal, and countermagic to keep the opponent from stabilizing. It wins through sustained attacks backed by burn rather than a large finisher or combo. The flexible interaction supports a reactive game plan, but the deck gives up explosive closing power and can lose momentum when it draws too many lands or its early threats are repeatedly answered.

The quality model grades this deck typical against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The deck runs its spells as playsets, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade.

- PLAN plan_coherent=yes: Cheap threats like Ragavan, Ledger Shredder and Faerie Mastermind are backed by Bolt, Discharge, Counterspell and cantrips, all serving one tempo plan.
- PLAN theme_fit=partly: It is a blue-red Modern tempo deck, but the request named Delver and the list has no Delver of Secrets or flip payoff, and it uses only 12 creatures with a Ragavan-style shell instead.
- PLAN useful_as_built=yes: With 26 lands including 12 dual-type untapped sources, a curve topping at three, and plentiful cheap threats and interaction, the deck can be played as it stands.
- PLAN summary_honest=partly: The claim of early creature pressure backed by burn and countermagic matches the list, but the deck has no actual Delver of Secrets and the summary does not mention that, and the 12 creatures are fewer than 'sustained attacks' implies.
- [INFO] `curve_summary`: average mana value 1.33 over 36 nonland cards

<details><summary>The deck list</summary>

- 4 Ragavan, Nimble Pilferer | threat | Provides early pressure and rewards connecting while interaction clears the way.
- 4 Ledger Shredder | threat | Supplies an evasive threat that grows while improving card selection.
- 4 Faerie Mastermind | threat | Adds evasive pressure while letting the deck hold up interaction.
- 4 Preordain | draw | Finds early threats, needed lands, and the right interaction.
- 2 Consider | draw | Provides inexpensive card selection without committing mana during your own turn.
- 4 Mishra's Bauble | synergy | Helps trigger Ledger Shredder while replacing itself without a mana investment.
- 4 Lightning Bolt | removal | Removes early blockers and supplies reach to finish close games.
- 4 Galvanic Discharge | removal | Answers creatures efficiently so your attackers can maintain pressure.
- 4 Counterspell | interaction | Stops key opposing plays once an early threat is established.
- 2 Force of Negation | interaction | Provides backup against important noncreature spells when mana is committed elsewhere.
- 4 Steam Vents | land | Supplies both colors for early threats and blue-heavy interaction.
- 4 Spirebluff Canal | land | Provides painless early fixing for the deck's low-cost spells.
- 2 Shivan Reef | land | Adds flexible blue-red fixing without an early tapped-land delay.
- 2 Fiery Islet | land | Fixes both colors and offers a way to cash in excess mana later.
- 8 Island | land | Provides reliable blue mana, especially for Counterspell.
- 4 Mountain | land | Supports early red threats and removal without additional life loss.

</details>

### 7. Modern burn, casual

Format: Modern. Theme: burn. Pool: any_card. Shortlist: 299 names.

Grade: typical, score 0.45, model 20260923T202806Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $55.26 to buy, $55.26 the whole deck.

**Summary:** This mono-red burn deck combines direct damage with a substantial creature plan. Apply pressure with early spells, remove blockers when necessary, and finish with burn or attacks from your larger threats. Draw support helps sustain pressure into the midgame. The deck trades the explosive speed of leaner burn builds for a stronger battlefield presence, so its opening turns are less aggressive and creature removal can disrupt its damage engines.

The quality model grades this deck typical against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: Burn spells, Torbran, Thermo-Alchemist and Sunspine Lynx all push damage at the face, but the four-mana creatures (Lambholt, Lynx, Tectonic Giant, Hazoret) form a midrange plan that sits awkwardly beside a reactive burn core, and Hazoret wants an empty hand that the Risk Factor and Artist's Talent draw package works against.
- PLAN theme_fit=partly: The deck is mono-red with Lightning Bolt, Boltwave, Risk Factor and Torbran, and it plays at a casual level, but the many four-drop creatures make it closer to red midrange than the pure burn deck requested.
- PLAN useful_as_built=partly: The 24 lands and the mono-red mana work fine, but the curve is top-heavy for burn with 14 four-mana creatures and few cheap plays, and only a handful of spells besides Bolt and Boltwave are reliable burn finishers, so the deck's reach is limited.
- PLAN summary_honest=yes: The summary claims a creature-heavy burn deck that trades speed for board presence, and the list matches: it runs 18 creatures and 10 burn or removal spells, with Risk Factor and Artist's Talent as the draw support it names.
- [INFO] `curve_summary`: average mana value 2.78 over 36 nonland cards

<details><summary>The deck list</summary>

- 20 Mountain | land | Provides reliable red mana without slowing the deck's early turns.
- 4 Ramunap Ruins | land | Supplies red mana and turns spare lands into late-game damage.
- 4 Risk Factor | draw | Pressures the opponent's life total or replenishes your hand, with jump-start providing another use.
- 2 Artist's Talent | draw | Filters draws and helps turn excess lands into useful spells.
- 2 Chandra, Dressed to Kill | ramp | Provides extra red mana while supporting the deck's damage and card-access plans.
- 4 Lightning Bolt | removal | Efficiently removes small creatures or delivers damage directly to the opponent.
- 2 Bonecrusher Giant // Stomp | removal | Pairs an early damage spell with a substantial creature for the midgame.
- 4 Boltwave | synergy | Provides inexpensive direct damage while triggering the deck's spell-based payoffs.
- 4 Thermo-Alchemist | synergy | Adds repeatable damage and rewards casting the deck's instants and sorceries.
- 4 Lambholt Raconteur // Lambholt Ravager | threat | Turns repeated noncreature spells into additional damage while contributing a creature threat.
- 4 Sunspine Lynx | threat | Provides a large attacker and punishes opponents relying heavily on nonbasic lands.
- 2 Tectonic Giant | threat | Offers a midgame attacker that can generate damage or additional cards.
- 2 Torbran, Thane of Red Fell | threat | Amplifies the damage from red spells and creatures to help close the game.
- 2 Hazoret the Fervent | threat | Provides a resilient finisher and converts spare cards into damage.

</details>

### 8. Modern lifegain, FNM

Format: Modern. Theme: lifegain. Pool: any_card. Shortlist: 302 names.

Grade: typical, score 0.51, model 20260923T202806Z.

Cards: 60 main, 15 sideboard. Repair turn: no. Block findings: 0.

Cost: $714.03 to buy, $714.03 the whole deck.

**Summary:** This white-black lifegain deck develops a creature board, keeps cards flowing, and wins through steady combat pressure backed by larger midgame threats. Removal helps keep opposing creatures from taking over, while the sideboard offers disruption, protection, and board resets for different matchups. The deck favors a sustained creature game over fast combo finishes, giving up some speed and extensive disruption for a stronger battlefield presence.

The quality model grades this deck typical against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade. Many of the cards appear in no top list, and that lowers the grade.

- PLAN plan_coherent=yes: Lifegain creatures like Guide of Souls, Ocelot Pride, Enduring Innocence, Elenda and Archangel of Thune all feed one sustained creature-combat plan with removal in support.
- PLAN theme_fit=yes: The deck is white-black, leans on lifegain payoffs and enablers, and sits at a casual FNM power level in Modern as requested.
- PLAN useful_as_built=partly: The two-color mana base is heavy on dual lands and the curve of one-drops through five-drops is playable, but the list has only about 8 removal spells and a few uneven inclusions such as Attended Healer and Legion's Landing with a few cards that are clunky, though it can still be played as it stands.
- PLAN summary_honest=partly: The summary's claim of a creature board with steady combat pressure and larger midgame threats matches the list, but it mentions a sideboard with disruption, protection and board resets that is not shown in the cards, and the deck has only a few card-flow pieces.
- [INFO] `curve_summary`: average mana value 3.06 over 36 nonland cards
- [INFO] `mana_pass`: the builder moved 1 card of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 4 Godless Shrine | land | Provides both colors for the deck's early creatures and later threats.
- 4 Caves of Koilos | land | Supports early plays in either color without slowing the curve.
- 4 Isolated Chapel | land | Adds white-black fixing alongside the basic lands and Godless Shrine.
- 2 Shattered Sanctum | land | Provides additional fixing for the midgame.
- 2 Fetid Heath | land | Helps support the heavier colored requirements of the threats.
- 3 Plains | land | Provides reliable white mana for the deck's opening plays.
- 5 Swamp | land | Provides reliable black mana for removal and threats.
- 2 Legion's Landing // Adanto, the First Fort | ramp | Fits the creature-heavy opening while providing a route to additional mana.
- 4 Guide of Souls | synergy | An early creature that ties the developing board to the lifegain plan.
- 4 Ocelot Pride | synergy | Provides an early payoff for combining lifegain with a growing creature board.
- 4 Enduring Innocence | draw | Keeps cards flowing alongside the deck's smaller creatures.
- 2 Preacher of the Schism | draw | Adds a creature-based source of card advantage for longer games.
- 4 Solitude | removal | Provides creature removal without abandoning the creature-heavy plan.
- 2 Gumdrop Poisoner // Tempt with Treats | removal | Adds removal that complements the lifegain strategy.
- 2 Ajani, Caller of the Pride | threat | Adds a planeswalker threat to a deck otherwise centered on creatures.
- 2 Attended Healer | threat | Provides another creature threat that rewards the lifegain plan.
- 2 Gisela, the Broken Blade | threat | Adds an evasive threat for pushing through stalled boards.
- 3 Elenda, Saint of Dusk | threat | Provides a substantial midgame threat for the white-black lifegain shell.
- 3 Sheoldred, the Apocalypse | threat | Adds a powerful threat that supports the deck's life-total pressure.
- 2 Archangel of Thune | threat | Serves as a top-end threat for the lifegain-focused creature board.

</details>

### 9. Standard midrange, FNM

Format: Standard. Theme: midrange. Pool: any_card. Shortlist: 179 names.

Grade: baseline, score 0.41, model 20260923T202806Z.

Cards: 60 main, 15 sideboard. Repair turn: no. Block findings: 0.

Cost: $172.02 to buy, $172.02 the whole deck.

**Summary:** This black-green midrange deck develops creature pressure, uses removal to keep opponents from taking over, and leans on draw to keep presenting threats in longer games. It wins through creature combat, with protection supporting its board and several creatures contributing to the draw or removal plan. It favors sustained pressure over explosive starts, so very fast opponents can force it onto the defensive before its midrange engine gets established.

The quality model grades this deck at the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. The curve sits high for the format, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The black-green midrange core of creatures, removal and card draw fits together, but the 24 lands and a few clunky cards like Massacre Girl, which wipes out the deck's own creatures, and the six-mana Rottenmouth Viper pull against a steady pressure plan.
- PLAN theme_fit=yes: The deck is a black-green midrange list built from creatures, removal and card advantage using Assassin's Trophy, Overgrown Tomb and Blooming Marsh, which matches the request for BG midrange at FNM level.
- PLAN useful_as_built=yes: With 28 lands including 12 dual lands, a low-cost removal suite, and a spread of three-drop creatures, the deck can be played as it stands, though the 28-land count is heavy and leaves it flood-prone.
- PLAN summary_honest=partly: The summary claims the deck wins through creature combat with removal and draw support, which the list does carry, but it never mentions the high curve that the grader notes, and it calls the protection suite present when the list has almost no actual protection spells.
- [INFO] `curve_summary`: average mana value 3.39 over 36 nonland cards

<details><summary>The deck list</summary>

- 4 Blooming Marsh | land | Provides fixing for the black-green mana base.
- 4 Overgrown Tomb | land | Supports both colors throughout the game.
- 4 Wastewood Verge | land | Adds fixing for the deck's two colors.
- 4 Restless Cottage | land | Rounds out the black-green fixing package.
- 4 Forest | land | Provides basic green mana for early development.
- 4 Swamp | land | Provides basic black mana for removal and draw.
- 2 Llanowar Elves | ramp | Accelerates the creature-focused midrange plan.
- 4 Unholy Annex // Ritual Chamber | draw | Supplies draw support for longer games.
- 2 Diresight | draw | Adds draw without requiring another creature on the battlefield.
- 4 Assassin's Trophy | removal | Provides removal to support creature pressure.
- 2 Bitter Triumph | removal | Complements the main removal package.
- 4 Midnight Reaper | synergy | Combines creature presence with draw support for the midrange plan.
- 4 Innkeeper's Talent | synergy | Supports the creature-heavy strategy with protection.
- 4 Surrak, Elusive Hunter | threat | Provides creature pressure while contributing to the draw plan.
- 4 Darkstar Augur | threat | Combines a creature threat with additional draw support.
- 2 Massacre Girl, Known Killer | threat | Adds another creature threat with a draw role.
- 2 Rottenmouth Viper | threat | Pairs creature pressure with the deck's removal plan.
- 2 Thrashing Brontodon | threat | Adds creature presence while supplementing removal.

</details>

### 10. Standard aggro, tournament

Format: Standard. Theme: aggro. Pool: any_card. Shortlist: 302 names.

Grade: baseline, score 0.38, model 20260923T202806Z.

Cards: 60 main, 15 sideboard. Repair turn: no. Block findings: 0.

Cost: $151.11 to buy, $151.11 the whole deck.

**Summary:** Build an attacking board, use removal and interaction to keep the pressure on, and win through combat. Creature-based draw support helps sustain the assault when the first wave is not enough, while Stingerback Terror supplies a substantial finishing threat. The sideboard offers more disruption, removal, and staying power for slower matchups. This deck favors a direct creature-based game plan over late-game inevitability, so keep applying pressure rather than settling into a prolonged exchange of resources.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The cheap red creatures, Boros Charm and the Case point toward aggro, but Twinmaw Stormbrood costs six, Delivery Moogle and Frilled Sparkshooter are four-drops, and the low creature count of about 22 spells is thin for a unified plan.
- PLAN theme_fit=yes: The deck is built in red and white with Sacred Foundry, Inspiring Vantage and Sunbillow Verge, it plays low-curve creatures and Boros Charm, and it follows the request for Standard aggro.
- PLAN useful_as_built=partly: The mana base of 24 lands with 12 dual lands is smooth, but with only 16 or so nonland cards, a small count of cheap creatures, and a six-drop Dragon, the deck has too few threats and a clunky curve to play reliably as it stands.
- PLAN summary_honest=partly: The summary says Stingerback Terror is a substantial finishing threat, which is true as a four-mana Dragon, but it calls the draw creatures sustaining support and the deck holds only six cheap creatures and few real removal spells, so the interaction claim is overstated.
- [INFO] `curve_summary`: average mana value 2.89 over 36 nonland cards

<details><summary>The deck list</summary>

- 4 Sacred Foundry | land | Provides both colors for the deck's aggressive opening turns.
- 4 Inspiring Vantage | land | Supports early red and white plays.
- 4 Sunbillow Verge | land | Adds red-white fixing alongside the basic lands.
- 6 Mountain | land | Provides a dependable red source for the creature-heavy core.
- 6 Plains | land | Supports the deck's white creatures, removal, and interaction.
- 4 Reckless Lackey | draw | Combines a creature slot with support for the deck's draw plan.
- 2 Fugitive Codebreaker | draw | Adds creature-based draw support without abandoning the attacking plan.
- 4 Boros Charm | interaction | Supplies interaction that complements an aggressive red-white strategy.
- 2 Take Up the Shield | interaction | Adds interaction focused on supporting the creature plan.
- 4 Case of the Gateway Express | removal | Provides removal to help the attacking board maintain pressure.
- 4 Twinmaw Stormbrood // Charring Bite | removal | Adds another removal option to keep opposing threats from taking over.
- 4 Stromkirk Noble | synergy | Provides creature-based synergy for the aggressive core.
- 4 Delivery Moogle | threat | Adds white creatures to the deck's attacking core.
- 4 Frilled Sparkshooter | threat | Supplies a full set of red threats for consistent pressure.
- 4 Stingerback Terror | threat | Provides a substantial threat to finish the creature-based assault.

</details>

### 11. Commander with a locked card

Format: Commander. Theme: sacrifice. Pool: any_card. Shortlist: 306 names.

Commander: Karlov of the Ghost Council.

Grade: baseline, score 0.33, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $1329.00 to buy, $1329.00 the whole deck.

**Summary:** Karlov of the Ghost Council leads a sacrifice deck that develops expendable creatures, establishes value engines, and turns its board into sustained pressure. Use early ramp and draw to keep developing, hold interaction for threats to the engine, and commit larger creatures once you can profit from losing them. Win through accumulated sacrifice payoffs or creature combat, with several finishers providing a stronger late game. The deck gives up some immediate aggression for staying power, and hands full of payoffs without enough creatures can develop slowly.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The mana base serves one color far better than another, and that lowers the grade. The cards pair in ways the top lists do not, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every card is a sacrifice outlet, a death-trigger payoff, or a cheap creature or token source, so the list pulls toward the same aristocrats plan under Karlov.
- PLAN theme_fit=yes: It is a Commander deck led by Karlov of the Ghost Council in white-black, packed with sacrifice outlets such as Viscera Seer, Ashnod's Altar, Phyrexian Altar and Yawgmoth, plus payoffs like Zulaport Cutthroat and Grave Pact.
- PLAN useful_as_built=partly: The deck has plenty of ramp, draw, outlets and payoffs and plays as a real sacrifice deck, but the mana base has many untapped dual lands yet 13 Swamps against only 6 Plains for white cards like Mondrak and Liesa, and it has few token makers, so it can stumble.
- PLAN summary_honest=partly: The summary's claim of early ramp and draw is borne out by the rocks and Priest of Forgotten Gods, but the claim of holding interaction is weak because the creature count looks thin and many 'interaction' cards are protection spells, while the deck also runs wipes that conflict with developing a board.
- [INFO] `curve_summary`: average mana value 3.02 over 61 nonland cards
- [INFO] `mana_pass`: the builder moved 3 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 1 Brightclimb Pathway // Grimclimb Pathway | land | Provides an untapped source of whichever color the opening hand needs.
- 1 Caves of Koilos | land | Provides immediate access to both colors.
- 1 City of Brass | land | Keeps both colors available for early creatures and interaction.
- 1 Command Tower | land | Provides reliable fixing throughout the game.
- 1 Concealed Courtyard | land | Supports fast opening turns with both colors.
- 1 Forbidden Orchard | land | Provides untapped fixing despite giving opponents extra creatures.
- 1 Godless Shrine | land | Provides both colors and a useful target for Marsh Flats.
- 1 Isolated Chapel | land | Provides both colors alongside the deck's basic land types.
- 1 Mana Confluence | land | Provides immediate fixing for color-intensive turns.
- 1 Marsh Flats | land | Finds the basic or dual land needed to complete the opening colors.
- 1 Prismatic Vista | land | Finds the missing basic color early.
- 1 Scrubland | land | Provides both colors without slowing the early curve.
- 1 Shattered Sanctum | land | Provides dependable dual-color mana after the opening turns.
- 1 Shineshadow Snarl | land | Uses the deck's basic land types to support early fixing.
- 1 Silent Clearing | land | Provides both colors and becomes a card when extra mana is unnecessary.
- 1 Starting Town | land | Supports early colored mana while remaining useful later.
- 1 Vault of Champions | land | Provides efficient dual-color mana in multiplayer games.
- 1 High Market | land | Adds a sacrifice outlet and a small life-gain trigger without using a spell slot.
- 1 Phyrexian Tower | land | Turns an expendable creature into a burst of black mana.
- 6 Plains | land | Provides stable white mana for Karlov and the deck's white spells.
- 13 Swamp | land | Supports the deck's heavier black requirements.
- 1 Sol Ring | ramp | Accelerates the sacrifice engines and larger threats.
- 1 Arcane Signet | ramp | Accelerates development while fixing either color.
- 1 Chrome Mox | ramp | Provides early colored acceleration when speed matters more than card quantity.
- 1 Mana Vault | ramp | Provides a burst of mana for an early engine or substantial threat.
- 1 Priest of Forgotten Gods | ramp | Turns spare creatures into mana, cards, and pressure on opponents.
- 1 Ashnod's Altar | ramp | Converts creatures into mana through a repeatable sacrifice outlet.
- 1 Phyrexian Altar | ramp | Converts creatures into colored mana for extended sacrifice turns.
- 1 Warren Soultrader | ramp | Turns creatures into Treasure while enabling sacrifice payoffs.
- 1 Pawn of Ulamog | ramp | Replaces dying nontoken creatures with expendable mana-producing bodies.
- 1 Pitiless Plunderer | ramp | Builds a Treasure reserve as other creatures die.
- 1 Esper Sentinel | draw | Supplies early card advantage while the sacrifice board develops.
- 1 Archivist of Oghma | draw | Pairs recurring card advantage with life gain that supports Karlov.
- 1 Village Rites | draw | Turns an expendable or threatened creature into fresh cards.
- 1 Corrupted Conviction | draw | Provides another inexpensive way to cash in a creature for cards.
- 1 Disciple of Bolas | draw | Converts a large creature into cards and life.
- 1 Smothering Abomination | draw | Rewards repeated sacrifices with sustained card draw.
- 1 Vampiric Rites | draw | Provides a repeatable outlet that supplies cards and life.
- 1 Baron Bertram Graywater | draw | Rewards token production and converts spare permanents into cards.
- 1 High-Society Hunter | draw | Combines creature-death card advantage with an evasive lifelinking body.
- 1 Shadowheart, Dark Justiciar | draw | Turns a high-power creature into a substantial refill.
- 1 Aven Interrupter | interaction | Disrupts a key spell while adding an evasive creature.
- 1 Boromir, Warden of the Tower | interaction | Interferes with free spells and offers protection for the creature board.
- 1 Duty Beyond Death | interaction | Uses a sacrifice to protect the remaining creatures.
- 1 Ephemerate | interaction | Protects an important creature or reuses an enters-the-battlefield effect.
- 1 Flare of Fortitude | interaction | Protects an established board during a critical opposing turn.
- 1 Reprieve | interaction | Buys time against a crucial spell without costing a card.
- 1 Soul Partition | interaction | Temporarily answers a troublesome permanent or saves an important engine.
- 1 Swords to Plowshares | removal | Answers a dangerous creature efficiently.
- 1 Yawgmoth, Thran Physician | removal | Turns expendable creatures into cards and opposing-creature suppression.
- 1 Grave Pact | removal | Makes each creature death pressure opposing boards.
- 1 Dictate of Erebos | removal | Turns sacrifice sequences into repeated opposing sacrifices.
- 1 Ayli, Eternal Pilgrim | removal | Provides life-gaining sacrifices and a late-game answer to nonland permanents.
- 1 Deadly Rollick | removal | Provides efficient exile removal while Karlov is on the battlefield.
- 1 Chittering Witch | removal | Supplies sacrifice fodder and converts creatures into removal.
- 1 Ruthless Lawbringer | removal | Trades an expendable creature for a troublesome nonland permanent.
- 1 Orcish Bowmasters | removal | Punishes opposing card draw while creating an additional body.
- 1 Elas il-Kor, Sadistic Pilgrim | synergy | Connects creature arrivals and deaths to life-gain and drain payoffs.
- 1 Zulaport Cutthroat | synergy | Turns creature deaths into life gain and pressure on every opponent.
- 1 Bastion of Remembrance | synergy | Provides a body and a resilient payoff for creature deaths.
- 1 Viscera Seer | synergy | Provides a cheap sacrifice outlet that improves upcoming draws.
- 1 Bartolomé del Presidio | synergy | Provides a free outlet for spare creatures and artifacts.
- 1 Chthonian Nightmare | synergy | Recycles creatures while repeatedly feeding sacrifice payoffs.
- 1 Pious Evangel // Wayward Disciple | synergy | Supports Karlov with life gain before becoming a creature-death payoff.
- 1 Liesa, Forgotten Archangel | threat | Adds an evasive lifelinking threat and helps recover sacrificed nontoken creatures.
- 1 Felisa, Fang of Silverquill | threat | Turns deaths of creatures with counters into more sacrifice fodder.
- 1 Mondrak, Glory Dominus | threat | Multiplies token production and offers a durable attacking body.
- 1 Ratadrabik of Urborg | threat | Preserves value when the deck's other legendary creatures die.
- 1 Ghoulcaller Gisa | threat | Converts one creature into a larger supply of bodies.
- 1 Requiem Angel | threat | Replenishes the board with flying tokens as other non-Spirit creatures die.
- 1 Demonlord of Ashmouth | threat | Provides substantial aerial pressure while feeding the sacrifice plan.
- 1 Demon of Catastrophes | threat | Turns a spare creature into a large flying trampler.
- 1 Vindictive Vampire | threat | Makes deaths of other creatures advance the drain plan.
- 1 Hell's Caretaker | threat | Trades expendable creatures for important creatures in the graveyard.
- 1 Sadistic Hypnotist | threat | Converts spare creatures into pressure on opposing hands.
- 1 Basri's Lieutenant | threat | Adds counters and leaves replacement bodies when counter-bearing creatures die.
- 1 Razaketh, the Foulblooded | wincon | Turns a developed creature board into the pieces needed to close the game.
- 1 Relic Vial | wincon | Combines a sacrifice-and-draw outlet with a finishing death payoff alongside Clerics.
- 1 Sephiroth, Fabled SOLDIER // Sephiroth, One-Winged Angel | wincon | Rewards repeated creature deaths and establishes a lasting finishing payoff.
- 1 Damn | wipe | Serves as focused removal early or a broad creature reset later.
- 1 Toxic Deluge | wipe | Resets threatening creature boards at a flexible size.
- 1 The Meathook Massacre | wipe | Clears smaller creatures while supporting the deck's death-payoff plan.

</details>

### 12. Commander on a budget

Format: Commander. Theme: tokens. Pool: any_card. Shortlist: 305 names.

Commander: Adeline, Resplendent Cathar.

Grade: bad, score 0.05, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band, over_budget. Block findings: 0.

Cost: $96.36 to buy, $96.36 the whole deck.

**Summary:** A budget-conscious, bracket-2 mono-white token deck built around Adeline's combat pressure. Repeatable token engines and anthem effects build a formidable army, while steady card advantage supports longer games. Flexible removal and protective interaction help preserve momentum, with evasive threats and an alternate victory route providing ways to close.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every card is a white token maker, anthem, token-doubler, or a support piece (ramp, draw, removal) that feeds the same go-wide plan led by Adeline.
- PLAN theme_fit=yes: It is a mono-white token deck led by Adeline, Resplendent Cathar, built from budget-level cards that suit bracket 2, with a few exceptions such as Sol Ring and Gideon, Ally of Zendikar.
- PLAN useful_as_built=partly: The deck has plenty of token makers, ramp, draw, and removal, but 37 lands plus 10 ramp pieces is excessive, and Command Tower and Fellwar Stone add little in mono-white, and there are several clunky top-end cards such as Vanquish the Horde and Triskelavus.
- PLAN summary_honest=yes: The summary claims an alternate victory route, and Halo Fountain and Windbrisk Heights are in the list, along with token engines like Felidar Retreat and Hero of Bladehold and anthems like Intangible Virtue.
- [INFO] `curve_summary`: average mana value 3.63 over 62 nonland cards

<details><summary>The deck list</summary>

- 34 Plains | land | Reliable white mana for the deck's spells.
- 1 Castle Ardenvale | land | Produces tokens when spare mana is available.
- 1 Command Tower | land | Untapped white mana.
- 1 Windbrisk Heights | land | Rewards attacking with a wide board.
- 1 Sol Ring | ramp | Accelerates token engines and larger threats.
- 1 Arcane Signet | ramp | Provides dependable white acceleration.
- 1 Fellwar Stone | ramp | Low-cost mana acceleration.
- 1 Springleaf Drum | ramp | Turns a spare creature into a mana source.
- 1 Goldvein Pick | ramp | Rewards combat damage with Treasure.
- 1 Keeper of the Accord | ramp | Helps catch up on lands and creatures.
- 1 Monologue Tax | ramp | Generates Treasure from opponents' busy turns.
- 1 Druidic Satchel | ramp | Offers repeatable land acceleration, tokens, or life.
- 1 The Restoration of Eiganjo // Architect of Restoration | ramp | Develops mana and later supplies an attacking token engine.
- 1 Collector's Vault | ramp | Filters cards while generating Treasure.
- 1 Idol of Oblivion | draw | Provides repeatable draw alongside token production.
- 1 Staff of the Storyteller | draw | Turns creature-token production into cards.
- 1 Glimmer Seeker | draw | Rewards attacking and surviving with card advantage.
- 1 Wedding Announcement // Wedding Festivity | draw | Supplies cards or creatures before boosting the army.
- 1 Bygone Bishop | draw | Provides Clues from smaller creature spells.
- 1 Search the Premises | draw | Generates Clues when opponents attack you.
- 1 Platoon Dispenser | draw | Supports a populated battlefield with cards and tokens.
- 1 Sanctuary Warden | draw | Combines card advantage, token production, and an evasive threat.
- 1 Faramir, Field Commander | draw | Provides cards after creatures die and can produce tokens.
- 1 Dawn of Hope | draw | Converts life gain into cards and offers repeatable token production.
- 1 Wojek Investigator | draw | Provides aerial presence and conditional Clue production.
- 1 Intangible Virtue | synergy | Strengthens tokens and keeps them available to defend.
- 1 Inspiring Leader | synergy | Makes the token army substantially more threatening.
- 1 Divine Visitation | synergy | Turns token production into a powerful aerial army.
- 1 Felidar Retreat | synergy | Converts land drops into creatures or team-wide growth.
- 1 Clarion Spirit | synergy | Rewards sequencing inexpensive spells with extra bodies.
- 1 Hanweir Militia Captain // Westvale Cult Leader | synergy | Rewards a wide battlefield with ongoing token production.
- 1 Prava of the Steel Legion | synergy | Strengthens attacking tokens and provides a mana sink.
- 1 Rosie Cotton of South Lane | synergy | Converts token creation into creature growth.
- 1 Cathar's Call | synergy | Provides steady Human production and vigilance.
- 1 Retreat to Emeria | synergy | Turns land drops into tokens or an offensive boost.
- 1 Hero of Bladehold | threat | Creates attackers and increases combat pressure.
- 1 God-Eternal Oketra | threat | Builds a substantial token army from creature spells.
- 1 Defiler of Faith | threat | Supports white spells with additional Soldiers.
- 1 Basri's Lieutenant | threat | Adds board presence and resilience through creature deaths.
- 1 Gideon, Ally of Zendikar | threat | Produces Allies or strengthens the whole army.
- 1 Emeria Angel | threat | Builds an evasive army from land drops.
- 1 Mite Overseer | threat | Improves token combat and provides additional bodies.
- 1 Romana II | threat | Copies useful tokens for sustained board development.
- 1 Phantom General | threat | Boosts the entire token army.
- 1 Silverwing Squadron | threat | Scales with the battlefield and produces attacking Knights.
- 1 Requiem Angel | threat | Turns creature losses into flying reinforcements.
- 1 Nahiri, the Lithomancer | threat | Creates creatures and offers a powerful Equipment-based endgame.
- 1 Swords to Plowshares | removal | Efficiently answers a dangerous creature.
- 1 Skyclave Apparition | removal | Removes a troublesome nonland permanent.
- 1 Shire Shirriff | removal | Uses expendable tokens to remove a creature.
- 1 The Wandering Emperor | removal | Provides flexible creature removal and combat support.
- 1 Ugin, the Ineffable | removal | Answers colored permanents and generates resources.
- 1 Trostani's Judgment | removal | Exiles a creature while adding another token.
- 1 Triskelavus | removal | Provides flying bodies that can become targeted damage.
- 1 Rootborn Defenses | interaction | Protects the army while copying a creature token.
- 1 Soul Partition | interaction | Temporarily answers a threat or rescues a key permanent.
- 1 Reprieve | interaction | Delays a critical opposing spell and replaces itself.
- 1 Moogles' Valor | interaction | Protects the battlefield and supports the token plan.
- 1 Parting Gust | interaction | Offers flexible creature exile or a way to reuse an ally.
- 1 Bilbo's Gambit | interaction | Provides instant-speed interaction with token support.
- 1 Hour of Reckoning | wipe | Clears nontoken creatures while preserving the token army.
- 1 Austere Command | wipe | Offers a flexible battlefield reset.
- 1 Vanquish the Horde | wipe | Provides an affordable reset on crowded battlefields.
- 1 Ancient Gold Dragon | wincon | Threatens explosive flying-token production.
- 1 Halo Fountain | wincon | Supports attacking creatures, card advantage, and an alternate victory.
- 1 Sword of Body and Mind | wincon | Adds protection and converts combat hits into additional pressure.

</details>

### 13. owned first, and the commander is not owned

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 245 names.

Commander: Karlov of the Ghost Council.

Grade: baseline, score 0.31, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: yes, for deck_size. Block findings: 0.

Cost: $1.98 to buy, $168.36 the whole deck.

**Summary:** An Orzhov lifegain deck built around Karlov's growing combat presence and creature control. Small value creatures, Food, and lifelinking attackers sustain the engine, while efficient removal and protective interaction help maintain momentum. Angel-led attacks and dedicated late-game payoffs provide closing power without relying on an infinite combo.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The cards pair in ways the top lists do not, and that lowers the grade. The mana base serves one color far better than another, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: Lifegain pieces like Angel of Vitality, Exemplar of Light, Aerith and Lembas share one plan, but the three board wipes work against a creature-based Karlov deck and the pile of ramp artifacts and generic removal pulls attention away from the engine.
- PLAN theme_fit=yes: This is a Commander deck led by Karlov of the Ghost Council in Orzhov colors with many lifegain payoffs and lifelink Angels, matching the lifegain request.
- PLAN useful_as_built=partly: The deck can be played as it stands, with 36 lands plus about ten mana rocks and plenty of removal, but the mana base is heavy on Plains with fixing lands that cost tempo, black is lightly supported, and the wipes and ramp suit a different plan.
- PLAN summary_honest=partly: The summary says efficient removal and Angel-led attacks, and the list has Swords, Infernal Grasp, Lyra, Angel of Invention and Victory's Herald, but it never mentions the three board wipes or the heavy ramp package, and Food is thin with only Lembas.
- [WARN] `not_owned`: Karlov of the Ghost Council: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 2.87 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Command Tower | land | Reliable color fixing.
- 1 City of Brass | land | Untapped color fixing.
- 1 Exotic Orchard | land | Multiplayer color fixing.
- 1 Spire of Industry | land | Artifact-supported color fixing.
- 1 Marsh Flats | land | Finds either basic land type.
- 1 Fabled Passage | land | Finds the needed basic land.
- 1 Evolving Wilds | land | Basic-land fixing.
- 1 Terramorphic Expanse | land | Basic-land fixing.
- 1 Path of Ancestry | land | Color fixing with occasional card selection.
- 1 Thriving Moor | land | Supports both deck colors.
- 1 Grand Coliseum | land | Flexible colored mana.
- 1 Opal Palace | land | Fixing that supports commander counters.
- 1 Great Hall of the Citadel | land | Helps cast legendary creatures.
- 14 Plains | land | Untapped white sources.
- 9 Swamp | land | Untapped black sources.
- 1 Sol Ring | ramp | Efficient acceleration.
- 1 Arcane Signet | ramp | Early acceleration and fixing.
- 1 Fellwar Stone | ramp | Low-cost multiplayer fixing.
- 1 Springleaf Drum | ramp | Turns small creatures into mana support.
- 1 Thought Vessel | ramp | Low-cost acceleration.
- 1 Wayfarer's Bauble | ramp | Develops basic lands.
- 1 Lotho, Corrupt Shirriff | ramp | Generates Treasure through active turns.
- 1 Chromatic Lantern | ramp | Acceleration and broad color fixing.
- 1 Inherited Envelope | ramp | Colored mana with Ring support.
- 1 Relic of Legends | ramp | Leverages legendary creatures for mana.
- 1 Exemplar of Light | draw | Connects lifegain and counters with card advantage.
- 1 Inspiring Overseer | draw | A card, life, and an evasive body.
- 1 Wall of Omens | draw | Early defense that replaces itself.
- 1 Call of the Ring | draw | Repeatable card advantage.
- 1 Night's Whisper | draw | Efficient hand replenishment.
- 1 Mask of Memory | draw | Rewards connecting in combat.
- 1 Tome of Legends | draw | Card advantage alongside the commander.
- 1 Lembas | draw | Card selection and a lifegain resource.
- 1 Stiltzkin, Moogle Merchant | draw | Repeatable card advantage with lifegain support.
- 1 Skullclamp | draw | Converts expendable creatures into cards.
- 1 Sunset Revelry | draw | Flexible recovery through cards, life, and creatures.
- 1 Boromir, Warden of the Tower | interaction | Disrupts free spells and protects the board.
- 1 Clever Concealment | interaction | Protects a developed battlefield.
- 1 Reprieve | interaction | Tempo interaction that replaces itself.
- 1 Take Up the Shield | interaction | Combat protection with lifegain potential.
- 1 Unbreakable Formation | interaction | Board protection and offensive support.
- 1 Orcish Medicine | interaction | Flexible creature protection.
- 1 Slip On the Ring | interaction | Protects creatures and reuses arrival abilities.
- 1 Frontline Medic | interaction | Supports combat and disrupts opposing X spells.
- 1 Swords to Plowshares | removal | Efficient creature exile.
- 1 Get Lost | removal | Answers several permanent types.
- 1 Generous Gift | removal | Flexible permanent removal.
- 1 Bitter Triumph | removal | Efficient creature or planeswalker removal.
- 1 Infernal Grasp | removal | Low-cost creature removal.
- 1 Orcish Bowmasters | removal | Punishes extra draws and picks off small creatures.
- 1 Stroke of Midnight | removal | Broad nonland removal.
- 1 Dismember | removal | Flexible-cost creature removal.
- 1 Crib Swap | removal | Exiles troublesome creatures.
- 1 Aerith Gainsborough | synergy | Lifegain and counter synergy.
- 1 Angel of Vitality | synergy | Amplifies lifegain and rewards a high life total.
- 1 Compassionate Healer | synergy | Provides repeatable lifegain.
- 1 Kor Firewalker | synergy | Adds lifegain against red spells.
- 1 Light of Promise | synergy | Turns lifegain into combat pressure.
- 1 Graveyard Trespasser // Graveyard Glutton | synergy | Graveyard pressure paired with life swings.
- 1 Night Nurse, Healer of Heroes | synergy | Supports the lifegain engine.
- 1 Rosie Cotton of South Lane | synergy | Food and token production support counter growth.
- 1 Denethor, Ruling Steward | synergy | Creates sacrifice value and repeatable life swings.
- 1 Angel of Invention | threat | Lifelink pressure and support for the wider board.
- 1 Bill the Pony | threat | Food generation and combat support.
- 1 Dawnhand Eulogist | threat | A lifegain-oriented battlefield threat.
- 1 Gollum, Silent Slinker // Meager Meal | threat | Evasive pressure with Food support.
- 1 Minwu, White Mage | threat | A central lifegain payoff.
- 1 Shattered Angel | threat | Recurring lifegain from opposing land development.
- 1 Sneering Shadewriter | threat | Adds pressure to the lifegain plan.
- 1 Victory's Herald | threat | Enables evasive, lifelinking attacks.
- 1 Lyra Dawnbringer | wincon | Powerful lifelinking finisher that strengthens Angels.
- 1 Frodo, Sauron's Bane | wincon | An alternate route to closing games.
- 1 Grave Venerations | wincon | A dedicated late-game payoff.
- 1 Fumigate | wipe | Resets creatures while replenishing life.
- 1 Dusk // Dawn | wipe | Clears large threats and recovers smaller creatures.
- 1 Austere Command | wipe | Selective sweeping across multiple permanent types.
- 1 Lightning Greaves | protection | Shields key creatures and enables immediate attacks.
- 1 Swiftfoot Boots | protection | Protects important creatures while preserving targeting flexibility.

</details>

### 14. the user delegates the commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 306 names.

Commander: Atarka, World Render.

Grade: baseline, score 0.06, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $917.20 to buy, $917.20 the whole deck.

**Summary:** Atarka, World Render leads a Dragon deck that develops mana early, establishes a threatening board, and turns Dragon attacks into double-striking pressure. Lower-cost threats keep the deck active before its larger finishers arrive, while draw support and inexpensive interaction help sustain the attack. Lathliss, Dragon Queen, Utvara Hellkite, and Wrathful Red Dragon provide additional ways to close the game. The deck favors tribal combat over elaborate combo lines, giving up some speed and resilience in exchange for a strong Dragon theme and explosive attacking turns.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The commander does not place in cEDH events, and that lowers the grade. The mana base serves one color far better than another, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: Most of the deck is red Dragons with Dragon payoffs, but the Green half is a scattering of tricks and protection spells like Legolas's Quick Reflexes, Royal Treatment and Keen Sense that do little for Dragon tribal, and the wipes work against the creature plan.
- PLAN theme_fit=yes: The deck is a Dragon tribal Commander deck led by Atarka, a Dragon commander in Red and Green, with a dense Dragon creature count that fits the request for dragons.
- PLAN useful_as_built=partly: The deck has plenty of ramp and Dragons, but Pyroblast, Red Elemental Blast and Legolas's Quick Reflexes are weak or dead in many games, the 10 Mountains and 4 Forests plus many painful and tapped lands strain a heavily red mana base, and Blasphemous Act and the other wipes can hurt the deck's own board.
- PLAN summary_honest=partly: The summary says Atarka turns Dragon attacks into double-striking pressure and names Lathliss, Utvara Hellkite and Wrathful Red Dragon as finishers, and all three are in the list, but Atarka's double strike works only on Dragons that attack alongside it and a 7-mana commander is a slow centerpiece, and the summary does not mention the board wipes or the Green tricks.
- [INFO] `curve_summary`: average mana value 3.14 over 63 nonland cards
- [INFO] `mana_pass`: the builder moved 3 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 1 Command Tower | land | Supports the deck's two-color mana base.
- 1 Taiga | land | Provides reliable mana for the deck's two colors.
- 1 Stomping Ground | land | Supports early plays and demanding Dragon costs.
- 1 Spire Garden | land | Strengthens the two-color mana base.
- 1 Karplusan Forest | land | Supports early acceleration and interaction.
- 1 Grove of the Burnwillows | land | Adds another source for either deck color.
- 1 City of Brass | land | Helps cast spells with demanding colored costs.
- 1 Mana Confluence | land | Keeps both deck colors readily available.
- 1 Forbidden Orchard | land | Prioritizes consistent access to both colors.
- 1 Exotic Orchard | land | Adds multiplayer-oriented mana fixing.
- 1 Copperline Gorge | land | Supports the deck's early development.
- 1 Rockfall Vale | land | Supports the mana base through the middle turns.
- 1 Rootbound Crag | land | Adds two-color fixing alongside the basic lands.
- 1 Game Trail | land | Helps balance red and green access.
- 1 Cinder Glade | land | Adds another two-color land to the mana base.
- 1 Cragcrown Pathway // Timbercrown Pathway | land | Provides flexible access to the color needed most.
- 1 Wooded Foothills | land | Helps assemble the appropriate colored mana.
- 1 Prismatic Vista | land | Supports reliable access to basic lands.
- 1 Fabled Passage | land | Helps balance the deck's basic-land colors.
- 1 Secluded Courtyard | land | Supports the Dragon-heavy creature package.
- 1 Unclaimed Territory | land | Helps pay for the deck's Dragons.
- 1 Path of Ancestry | land | Adds mana fixing suited to the tribal plan.
- 10 Mountain | land | Provides dependable red mana for Dragons and interaction.
- 4 Forest | land | Provides dependable green mana for development and support.
- 1 Sol Ring | ramp | Accelerates the deck toward its expensive Dragons.
- 1 Birds of Paradise | ramp | Provides early acceleration for the two-color plan.
- 1 Arcane Signet | ramp | Adds inexpensive mana development.
- 1 Orb of Dragonkind | ramp | Supports early development toward Dragon threats.
- 1 Pillar of Origins | ramp | Adds inexpensive acceleration for the creature-heavy plan.
- 1 Scaled Nurturer | ramp | Develops mana while fitting the Dragon theme.
- 1 Sarkhan, Dragon Ascendant | ramp | Provides early acceleration for the Dragon strategy.
- 1 Carnelian Orb of Dragonkind | ramp | Helps bridge the gap to larger Dragons.
- 1 Jade Orb of Dragonkind | ramp | Adds another mana-development piece for the Dragon plan.
- 1 Dragon's Hoard | ramp | Supports sustained mana development.
- 1 Goldspan Dragon | ramp | Combines Dragon pressure with further acceleration.
- 1 Ganax, Astral Hunter | ramp | Adds a Dragon that supports continued mana development.
- 1 Faithless Looting | draw | Provides inexpensive access to fresh cards.
- 1 Thrill of Possibility | draw | Helps refresh the hand without a large mana investment.
- 1 Sylvan Library | draw | Supports consistent access to useful cards.
- 1 Garruk's Uprising | draw | Provides card support for the large-creature strategy.
- 1 Dragonborn Champion | draw | Adds card flow on a Dragon body.
- 1 Beast Whisperer | draw | Supports card flow in a creature-heavy deck.
- 1 Return of the Wildspeaker | draw | Provides substantial card support for the creature plan.
- 1 For the Ancestors | draw | Helps maintain access to the tribal creature package.
- 1 Keen Sense | draw | Adds a low-cost draw option.
- 1 Mishra's Bauble | draw | Provides a minimal-cost contribution to card flow.
- 1 Heroic Intervention | interaction | Helps preserve a developed board against opposing interaction.
- 1 Legolas's Quick Reflexes | interaction | Provides inexpensive interaction for important turns.
- 1 Pyroblast | interaction | Adds a low-cost answer for relevant opposing plays.
- 1 Red Elemental Blast | interaction | Provides another inexpensive interactive option.
- 1 Royal Treatment | interaction | Helps defend a key creature at low cost.
- 1 Tibalt's Trickery | interaction | Provides interaction against a pivotal opposing spell.
- 1 Veil of Summer | interaction | Helps contest opposing interaction during important turns.
- 1 Dragon Tempest | removal | Adds removal support that fits the Dragon strategy.
- 1 Dragon's Fire | removal | Provides inexpensive removal within the Dragon theme.
- 1 Draconic Roar | removal | Adds another low-cost removal spell.
- 1 Lightning Bolt | removal | Keeps an efficient removal option available early.
- 1 Balefire Dragon | removal | Combines removal with a substantial finishing threat.
- 1 Untimely Malfunction | removal | Adds inexpensive removal utility.
- 1 Spit Flame | removal | Provides removal suited to the Dragon-heavy plan.
- 1 Terror of the Peaks | removal | Adds a finishing Dragon to the removal package.
- 1 Invasion of Tarkir // Defiant Thundermaw | removal | Provides removal while reinforcing the Dragon theme.
- 1 Acolyte of Bahamut | synergy | Supports the deck's Dragon-focused development.
- 1 Dragonlord's Servant | synergy | Helps the deck develop its Dragon army.
- 1 Dragonspeaker Shaman | synergy | Supports repeated deployment of Dragons.
- 1 Urza's Incubator | synergy | Strengthens the tribal development plan.
- 1 Herald's Horn | synergy | Provides support for the Dragon-heavy creature package.
- 1 Shared Animosity | synergy | Rewards committing to a tribal combat strategy.
- 1 Molten Echoes | synergy | Reinforces the deck's Dragon-based board development.
- 1 Thunderbreak Regent | threat | Provides a Dragon threat before the most expensive finishers.
- 1 Stormbreath Dragon | threat | Adds another Dragon to the combat plan.
- 1 Territorial Hellkite | threat | Supplies a comparatively inexpensive finishing threat.
- 1 Manaform Hellkite | threat | Adds a Dragon threat at the middle of the curve.
- 1 Mirrorwing Dragon | threat | Strengthens the deck's Dragon combat presence.
- 1 Kura, the Boundless Sky | threat | Adds another substantial Dragon to the board.
- 1 Roaming Throne | threat | Provides a supporting threat for the tribal strategy.
- 1 Vengeful Ancestor | threat | Adds a Dragon threat without pushing the curve too high.
- 1 Verix Bladewing | threat | Provides another Dragon threat in the middle of the curve.
- 1 Thrakkus the Butcher | threat | Reinforces the Dragon combat package.
- 1 Scourge of the Throne | threat | Adds a high-impact finishing threat.
- 1 Twinflame Tyrant | threat | Provides another Dragon capable of closing a game.
- 1 Lathliss, Dragon Queen | wincon | Provides a finishing centerpiece for the Dragon strategy.
- 1 Utvara Hellkite | wincon | Supplies a top-end finisher for a developed Dragon board.
- 1 Wrathful Red Dragon | wincon | Adds a complementary Dragon-based route to victory.
- 1 Blasphemous Act | wipe | Provides a board reset when opponents develop faster.
- 1 Breath Weapon | wipe | Adds a lower-cost board-control option.
- 1 Vandalblast | wipe | Provides a complementary sweeping answer.

</details>

### 15. delegated commander, owned first

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 245 names.

Commander: Astarion, the Decadent.

Grade: bad, score 0.04, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $2.41 to buy, $173.02 the whole deck.

**Summary:** Astarion pairs a lifegain plan with a creature-heavy board, using its end-step choice to grow your life cushion or press an opponent toward defeat. Angels and other threats supply the main attacking force, backed by dedicated finishers for longer games. Card draw, removal, and protective interaction help sustain that pressure. The deck favors steady development and board-based wins over fast combos, so rebuilding after repeated board wipes can take time.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. The commander does not place in cEDH events, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: Most cards support lifegain with Angels and lifelink threats, but three board wipes (Fumigate, Austere Command, Dusk // Dawn) fight the creature-heavy board, and several cards such as Lembas, Buster Sword and Bill the Pony are loose fits.
- PLAN theme_fit=yes: It is a white-black lifegain Commander deck led by Astarion, the Decadent, whose end-step choice is lifegain or life loss, and it has a modest bracket 2 power level, though the request's 'from my library first' limit is not visible in the list.
- PLAN useful_as_built=partly: The deck has plenty of ramp, draw, removal and finishers, but 28 basics with many pure single-color lands, heavy white costs like Victory's Herald and Enduring Angel, and nine mana rocks make it clunky, though it still plays as it stands.
- PLAN summary_honest=yes: The summary says Angels and other threats are the main attackers, and the list has Lyra, Angel of Invention, Victory's Herald, Enduring Angel and Shattered Angel, and it openly admits that board wipes slow rebuilding.
- [WARN] `not_owned`: Astarion, the Decadent: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.18 over 62 nonland cards
- [INFO] `mana_pass`: the builder moved 3 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 17 Plains | land | Provides dependable white mana for the deck's creatures and interaction.
- 11 Swamp | land | Provides dependable black mana for Astarion and the deck's black spells.
- 1 Command Tower | land | Provides either color without slowing early development.
- 1 City of Brass | land | Provides flexible mana throughout the game.
- 1 Exotic Orchard | land | Adds flexible colored mana in multiplayer games.
- 1 Thriving Moor | land | Supports both colors with white chosen as its additional color.
- 1 Grand Coliseum | land | Helps cover demanding colored costs later in the game.
- 1 Spire of Industry | land | Uses the deck's artifacts to support colored mana production.
- 1 Path of Ancestry | land | Provides both colors and occasional filtering with matching creature types.
- 1 Marsh Flats | land | Finds whichever basic land the opening hand needs.
- 1 Fabled Passage | land | Finds either basic land to balance the mana base.
- 1 Sol Ring | ramp | Accelerates the deck toward its larger creatures and commander.
- 1 Arcane Signet | ramp | Provides inexpensive acceleration and both commander colors.
- 1 Fellwar Stone | ramp | Adds inexpensive acceleration with useful multiplayer fixing.
- 1 Thought Vessel | ramp | Accelerates early development while supporting a larger hand.
- 1 Bender's Waterskin | ramp | Adds another ramp artifact to support the deck's midgame.
- 1 Chromatic Lantern | ramp | Accelerates mana while smoothing the deck's colored costs.
- 1 Commander's Sphere | ramp | Provides colored acceleration and can become a card when unnecessary.
- 1 Relic of Legends | ramp | Provides fixing and additional mana from the deck's legendary creatures.
- 1 Inherited Envelope | ramp | Adds colored acceleration for Astarion and larger threats.
- 1 Lotho, Corrupt Shirriff | ramp | Produces Treasure as players cast multiple spells.
- 1 Exemplar of Light | draw | Connects the lifegain plan to repeatable card advantage.
- 1 Inspiring Overseer | draw | Replaces itself while contributing life and an attacking body.
- 1 Wall of Omens | draw | Replaces itself and helps defend the early turns.
- 1 Night's Whisper | draw | Provides inexpensive cards to keep development moving.
- 1 Call of the Ring | draw | Supplies ongoing card advantage with a life cost the deck can absorb.
- 1 Lembas | draw | Provides card selection and a later source of life.
- 1 Mask of Memory | draw | Turns successful attacks into cards and hand filtering.
- 1 Tome of Legends | draw | Rewards deploying and attacking with Astarion.
- 1 Buster Sword | draw | Adds combat-based card advantage to the creature plan.
- 1 Sunset Revelry | draw | Offers catch-up value when opponents pull ahead.
- 1 Nasty End | draw | Converts a creature facing removal into fresh cards.
- 1 Boromir, Warden of the Tower | interaction | Protects the creature board and disrupts spells cast without mana.
- 1 Reprieve | interaction | Buys time against an important opposing spell while replacing itself.
- 1 Unbreakable Formation | interaction | Protects the board or strengthens a committed attack.
- 1 Clever Concealment | interaction | Preserves important permanents through a dangerous turn.
- 1 Take Up the Shield | interaction | Protects a creature while supporting lifegain in combat.
- 1 Orcish Medicine | interaction | Offers flexible creature protection and defensive support.
- 1 Swords to Plowshares | removal | Answers a dangerous creature for very little mana.
- 1 Generous Gift | removal | Answers troublesome permanents across several card types.
- 1 Get Lost | removal | Provides an efficient answer to creatures, enchantments, and planeswalkers.
- 1 Infernal Grasp | removal | Removes a threatening creature at instant speed.
- 1 Bitter Triumph | removal | Answers a creature or planeswalker with a manageable additional cost.
- 1 Crib Swap | removal | Exiles a creature that would be troublesome to destroy.
- 1 Witch-king of Angmar | removal | Combines sacrifice-based interaction with a substantial attacking threat.
- 1 Angel of Vitality | synergy | Strengthens life gains and rewards maintaining a high life total.
- 1 Light of Promise | synergy | Turns repeated life gains into growth for an important creature.
- 1 Denethor, Ruling Steward | synergy | Connects creature sacrifice with life gain and opposing life loss.
- 1 Graveyard Trespasser // Graveyard Glutton | synergy | Combines graveyard pressure with life gain and opposing life loss.
- 1 Rosie Cotton of South Lane | synergy | Rewards the deck's token production with creature growth.
- 1 Eastfarthing Farmer | synergy | Supplies Food and turns a healthy life total into combat pressure.
- 1 Kor Firewalker | synergy | Adds a low-cost creature with a recurring lifegain opportunity.
- 1 Aerith Gainsborough | synergy | Adds another creature-based synergy piece to the lifegain shell.
- 1 Elixir | synergy | Adds artifact-based support for the deck's lifegain plan.
- 1 White Mage's Staff | synergy | Adds equipment-based support to the creature-focused lifegain plan.
- 1 Angel of Invention | threat | Develops the board and strengthens the deck's attacking creatures.
- 1 Bill the Pony | threat | Provides a creature threat alongside Food-based support.
- 1 Dawnhand Eulogist | threat | Adds a marked finisher to the creature threat package.
- 1 Gollum, Silent Slinker // Meager Meal | threat | Adds another threat with an Adventure for additional flexibility.
- 1 Minwu, White Mage | threat | Adds a legendary creature threat to the midgame board.
- 1 Rabaroo Troop | threat | Broadens the creature threat package for sustained attacking pressure.
- 1 Reaping Willow | threat | Adds another substantial creature threat to the board-focused plan.
- 1 Shattered Angel | threat | Builds life as opponents develop their lands while supplying an attacker.
- 1 Sneering Shadewriter | threat | Provides another marked finisher among the creature threats.
- 1 Victory's Herald | threat | Supports a decisive attack with team-wide flying and lifelink.
- 1 Archfiend of Ifnir | threat | Provides a large flying threat with additional creature-control potential.
- 1 Enduring Angel // Angelic Enforcer | threat | Combines an attacking threat with insurance against a collapsing life total.
- 1 Lyra Dawnbringer | wincon | Provides a combat finisher that strengthens the deck's other Angels.
- 1 Frodo, Sauron's Bane | wincon | Offers a creature-based alternative route to finishing an opponent.
- 1 Grave Venerations | wincon | Adds a dedicated enchantment finisher alongside the creature plan.
- 1 Fumigate | wipe | Resets a crowded creature board while rebuilding the deck's life cushion.
- 1 Austere Command | wipe | Provides a flexible reset tailored to the most troublesome permanents.
- 1 Dusk // Dawn | wipe | Clears larger creatures and later recovers smaller creatures.

</details>

### 16. a tight budget, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 246 names.

Commander: Denethor, Ruling Steward.

Grade: baseline, score 0.30, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $19.97 to buy, $127.58 the whole deck.

**Summary:** Denethor leads a steady aristocrats deck that develops mana, builds a supply of creatures, and turns sacrifices into accumulating advantage. Keep the engine stocked with draw, use interaction to preserve important pieces, and finish through repeated life-loss pressure or a larger creature board. Board resets give you a way back when opponents pull ahead. The deck favors resilient, incremental progress over explosive turns, so it gives up some speed and depends on keeping creatures and payoff pieces available.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The commander does not place in cEDH events, and that lowers the grade. Few decks lead with the commander, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every card feeds one aristocrats plan: sacrifice outlets and creature-death engines like Blood Artist, Zulaport Cutthroat, Falkenrath Noble and Teysa, plus Skullclamp, ramp and card draw to keep the engine stocked.
- PLAN theme_fit=partly: The deck is a clear white-black aristocrats list led by Denethor, but the request asked to use cards from the player's own library first and nothing in the list shows that; it also leans on Sol Ring, Swords and Austere Command at a bracket 2 power level, and the $25 budget is unconfirmed.
- PLAN useful_as_built=yes: It has 38 lands plus about 11 mana rocks, a low curve with many two- and three-drops, ample removal and draw, and plenty of drain payoffs, so it can be played as it stands.
- PLAN summary_honest=yes: The summary says the deck finishes through repeated life-loss pressure, and Blood Artist, Zulaport Cutthroat, Falkenrath Noble, Vindictive Vampire and Bastion of Remembrance all carry that plan, with three wraths as the board resets it names.
- [WARN] `not_owned`: Erebos, Bleak-Hearted: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Smothering Abomination: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Shadowheart, Dark Justiciar: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Disciple of Bolas: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ruthless Lawbringer: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Woe Strider: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Yahenni, Undying Partisan: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bartolomé del Presidio: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ayli, Eternal Pilgrim: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Aron, Benalia's Ruin: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Baron Bertram Graywater: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Teysa, Orzhov Scion: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ghoulcaller Gisa: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Shilgengar, Sire of Famine: the deck needs 1, the collection has 0
- [WARN] `not_owned`: High-Society Hunter: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Falkenrath Noble: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vindictive Vampire: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vengeful Bloodwitch: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Elas il-Kor, Sadistic Pilgrim: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Lord Skitter's Butcher: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Old Flitterfang: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Lord of the Forsaken: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Zulaport Cutthroat: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Blood Artist: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bastion of Remembrance: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.00 over 62 nonland cards
- [INFO] `mana_pass`: the builder moved 3 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 1 Skullclamp | draw | Turns expendable creatures into a steady supply of cards.
- 1 Idol of Oblivion | draw | Provides repeatable draw alongside the deck's token production.
- 1 Night's Whisper | draw | An inexpensive way to refill your hand early.
- 1 Nasty End | draw | Converts a creature into cards while supporting the sacrifice plan.
- 1 Call of the Ring | draw | An owned source of ongoing card advantage.
- 1 Inspiring Overseer | draw | Replaces itself and leaves a creature available for combat or sacrifice.
- 1 Wall of Omens | draw | Replaces itself while providing an early defensive body.
- 1 Erebos, Bleak-Hearted | draw | Helps turn creature deaths into additional cards.
- 1 Smothering Abomination | draw | Rewards the deck's repeated sacrifices with card draw.
- 1 Shadowheart, Dark Justiciar | draw | Provides another way to turn creatures into a fresh hand.
- 1 Disciple of Bolas | draw | Converts a larger creature into cards and life.
- 1 Boromir, Warden of the Tower | interaction | An owned interactive creature that helps defend your board.
- 1 Clever Concealment | interaction | Preserves your established board through opposing disruption.
- 1 Reprieve | interaction | Buys time against an important opposing spell without exhausting your hand.
- 1 Unbreakable Formation | interaction | Helps your creature board survive a dangerous turn.
- 1 Orcish Medicine | interaction | Provides inexpensive flexibility when a key creature is threatened.
- 1 Slip On the Ring | interaction | Protects a useful creature and can reuse an arrival effect.
- 1 Sol Ring | ramp | Accelerates development and helps fund sacrifice activations later.
- 1 Springleaf Drum | ramp | Turns an available creature into inexpensive mana fixing.
- 1 Arcane Signet | ramp | Provides early acceleration in either deck color.
- 1 Fellwar Stone | ramp | Adds inexpensive acceleration with potential color fixing.
- 1 Thought Vessel | ramp | An owned two-mana accelerator for a smoother opening.
- 1 Deadly Dispute | ramp | Combines a sacrifice with cards and temporary acceleration.
- 1 Lotho, Corrupt Shirriff | ramp | An owned source of additional mana during active turns.
- 1 Relic of Legends | ramp | Makes the deck's many legendary creatures useful for mana development.
- 1 Commander's Sphere | ramp | Fixes both colors and remains useful when additional mana is unnecessary.
- 1 Chromatic Lantern | ramp | Provides acceleration while smoothing the two-color mana base.
- 1 Swords to Plowshares | removal | Answers a dangerous creature efficiently.
- 1 Generous Gift | removal | Provides a flexible answer to a troublesome permanent.
- 1 Infernal Grasp | removal | Offers straightforward instant-speed creature removal.
- 1 Bitter Triumph | removal | Answers a creature or planeswalker at a low mana cost.
- 1 Fiend Hunter | removal | Combines creature interaction with a body for the sacrifice engine.
- 1 Ruthless Lawbringer | removal | Turns expendable material into an answer to an opposing permanent.
- 1 Witch-king of Angmar | removal | An owned removal threat that also supplies a substantial late-game presence.
- 1 Woe Strider | synergy | Provides a repeatable sacrifice outlet and additional creature material.
- 1 Yahenni, Undying Partisan | synergy | Adds a sacrifice outlet that can remain resilient through removal.
- 1 Bartolomé del Presidio | synergy | Provides an inexpensive sacrifice outlet that grows with use.
- 1 Ayli, Eternal Pilgrim | synergy | Adds a low-cost sacrifice outlet with useful defensive value.
- 1 Aron, Benalia's Ruin | synergy | Converts sacrifices into a stronger creature board.
- 1 Baron Bertram Graywater | synergy | Supports the token-and-sacrifice engine while helping replenish resources.
- 1 Teysa, Orzhov Scion | synergy | Links creature deaths, replacement bodies, and creature interaction.
- 1 Gollum, Patient Plotter | synergy | Offers reusable creature material for the sacrifice plan.
- 1 Gollum the Abandoned | synergy | An owned creature that gives expendable bodies another useful purpose.
- 1 Gríma Wormtongue | synergy | Adds another owned sacrifice outlet to the creature engine.
- 1 Bill the Pony | threat | An owned creature that adds supporting material to the board.
- 1 Vengeful Villagers | threat | Uses an owned threat to maintain creature pressure.
- 1 Ghoulcaller Gisa | threat | Turns a creature into a wider board for pressure and future sacrifices.
- 1 Shilgengar, Sire of Famine | threat | Adds a substantial threat that supports sacrifice and recovery.
- 1 High-Society Hunter | threat | Supplies a finishing threat that rewards creature deaths.
- 1 Falkenrath Noble | threat | Adds a finishing creature that makes repeated deaths costly for opponents.
- 1 Vindictive Vampire | threat | Makes losing your creatures contribute to the deck's closing pressure.
- 1 Vengeful Bloodwitch | threat | Adds an inexpensive death payoff to the creature board.
- 1 Elas il-Kor, Sadistic Pilgrim | threat | Connects creature development and creature deaths to life-total pressure.
- 1 Lord Skitter's Butcher | threat | Provides a flexible creature that supports board development.
- 1 Old Flitterfang | threat | Adds a larger creature that benefits from the deck's attrition plan.
- 1 Lord of the Forsaken | threat | Provides a large finishing threat when smaller creatures cannot break through.
- 1 Zulaport Cutthroat | wincon | Turns repeated creature deaths into pressure on every opponent.
- 1 Blood Artist | wincon | Makes the sacrifice engine a direct route toward finishing opponents.
- 1 Bastion of Remembrance | wincon | Provides a persistent death payoff and an additional creature.
- 1 Dusk // Dawn | wipe | Offers a board reset with a later route to recovering small creatures.
- 1 Austere Command | wipe | Provides a flexible reset when the table develops beyond your position.
- 1 Martial Coup | wipe | Combines a late-game board reset with replacement creature material.
- 1 Command Tower | land | Provides untapped fixing for both deck colors.
- 1 City of Brass | land | An owned untapped source of either color.
- 1 Exotic Orchard | land | Adds flexible colored mana without a tapped opening.
- 1 Marsh Flats | land | Finds the basic land needed to balance your colors.
- 1 Fabled Passage | land | Provides another flexible route to the needed basic land.
- 1 Thriving Moor | land | Adds black mana and a chosen second color.
- 1 Path of Ancestry | land | Provides both colors with additional value for matching creatures.
- 1 Spire of Industry | land | The artifact package helps this land provide colored mana.
- 1 Grand Coliseum | land | Adds another owned source capable of producing either color.
- 1 The Grey Havens | land | The legendary creature package supports its colored-mana utility.
- 16 Plains | land | Provides reliable untapped white mana using owned basics.
- 11 Swamp | land | Provides reliable untapped black mana for the deck's heavier black requirements.

</details>

### 17. upgrade a precon, owned first

Format: Commander. Theme: goblins. Pool: owned_first. Shortlist: 236 names.

Commander: Zada, Hedron Grinder.

Grade: typical, score 0.52, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $36.19 to buy, $200.81 the whole deck.

**Summary:** Build a wide Goblin board, then use your commander to spread targeted spells across your creatures and sustain an explosive storm turn. The Goblin synergy package supplies steady pressure, while Great Train Heist and Assault on Osgiliath provide finishing plays. Early mana development helps establish the commander before committing to a big turn. The deck favors coordinated creature-and-spell turns over standalone threats, leaving it vulnerable when opponents repeatedly clear the board or disrupt its setup.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The mana base serves the colors evenly, and that raises the grade. The deck holds less interaction than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: Most cards feed a mono-red Goblin token and storm plan, but Blasphemous Act and Vandalblast wipe the board the deck builds, and a pile of cheap cantrips and rituals competes with the Goblin creature count.
- PLAN theme_fit=yes: This is a mono-red Zada, Hedron Grinder deck built around Goblins, tokens and cheap spells, which is the Goblin Storm upgrade the person asked for at a bracket 3 power level.
- PLAN useful_as_built=partly: The deck has plenty of lands, ramp and Goblin bodies, but it runs many low-impact cantrips, two board wipes, and little interaction, and some win cards like Haze of Rage and Great Train Heist need a wide board to pay off.
- PLAN summary_honest=partly: The summary says Zada spreads targeted spells across the creatures, which fits Ancestors' Aid, Fists of Flame and Frontline Heroism, but it never mentions the wipes that cut against the wide board it describes, and it overstates early mana development when ramp is mostly rituals.
- [INFO] `curve_summary`: average mana value 2.53 over 62 nonland cards
- [WARN] `profile_off_band`: the interaction count is 0, and bracket 3 wants 4 to 12
- [INFO] `bracket_cut`: to hold bracket 3, the builder cut 1 card: Storm-Kiln Artist (the combo Storm-Kiln Artist + Haze of Rage), and added 1 basic land

<details><summary>The deck list</summary>

- 1 Ancestors' Aid | synergy | Adds another targeted spell to the commander's spell-copying plan.
- 1 Ancestral Anger | draw | Adds inexpensive card flow to the spell-heavy game plan.
- 1 Arcane Signet | ramp | Provides early mana development for the commander and follow-up spells.
- 1 Arena of Glory | land | Preserves the precon's land base.
- 1 Assault on Osgiliath | wincon | Provides a finishing play for the deck's explosive turns.
- 1 Battle Hymn | ramp | Supports the mana needed for a large spell turn.
- 1 Blasphemous Act | wipe | Provides a reset when opposing boards become overwhelming.
- 1 Boggart Shenanigans | removal | Keeps a Goblin-themed removal element in the deck.
- 1 Brightstone Ritual | ramp | Supplies burst mana for the storm plan.
- 1 Broadside Bombardiers | removal | Combines a Goblin body with the deck's removal package.
- 1 Castle Embereth | land | Preserves a land from the precon's mana base.
- 1 Chaos Warp | removal | Provides interaction with troublesome opposing permanents.
- 1 Conspicuous Snoop | synergy | Maintains the precon's Goblin synergy package.
- 1 Crimson Wisps | draw | Supports card flow during the commander's spell turns.
- 1 Den of the Bugbear | land | Preserves the precon's land base.
- 1 Dragon Fodder | synergy | Supports the board-building half of the Goblin storm plan.
- 1 Dwarven Mine | land | Keeps a Mountain land in the mana base.
- 1 Empty the Warrens | synergy | Connects the spell-heavy plan with building a wide board.
- 1 Expedite | draw | Adds card flow to the commander's targeted-spell package.
- 1 Faithless Looting | draw | Provides card selection to assemble productive turns.
- 1 Fists of Flame | draw | Supports card flow within the deck's spell-driven plan.
- 1 Forgotten Cave | land | Preserves the precon's land base.
- 1 Fountainport | land | Keeps an existing utility land without expanding the colorless land package.
- 1 Frontline Heroism | synergy | Retains a synergy piece for the creature-and-spell plan.
- 1 Gempalm Incinerator | removal | Keeps removal attached to the Goblin theme.
- 1 General Kreat, the Boltbringer | synergy | Preserves a central Goblin synergy piece.
- 1 Glimpse the Impossible | draw | Helps replenish resources for further development.
- 1 Goblin Bombardment | removal | Maintains interaction within the creature-heavy strategy.
- 1 Goblin Burrows | land | Preserves an existing utility land.
- 1 Goblin Bushwhacker | synergy | Supports the deck's aggressive Goblin plan.
- 1 Goblin Chieftain | synergy | Keeps the Goblin synergy core intact.
- 1 Goblin Dark-Dwellers | threat | Provides a substantial Goblin threat.
- 1 Goblin Lackey | synergy | Maintains an inexpensive creature in the Goblin synergy package.
- 1 Goblin Matron | synergy | Preserves the precon's Goblin support package.
- 1 Goblin Negotiation | removal | Keeps another removal spell available during development.
- 1 Goblin Trashmaster | removal | Provides removal while preserving the Goblin creature core.
- 1 Goblin Warchief | synergy | Supports the deck's Goblin-focused development.
- 1 Grapeshot | removal | Adds removal to the spell-heavy storm package.
- 1 Great Train Heist | wincon | Preserves a finishing play for a decisive turn.
- 1 Grenzo, Havoc Raiser | synergy | Keeps another synergistic Goblin in the creature base.
- 1 Haze of Rage | wincon | Supports the deck's explosive closing turns.
- 1 Hidden Volcano | land | Preserves the precon's land base.
- 1 Howlsquad Heavy | ramp | Contributes mana development without abandoning the Goblin theme.
- 1 Idol of Oblivion | draw | Provides continuing card flow for the board-building plan.
- 1 Impact Tremors | wincon | Preserves an additional route to closing the game.
- 1 Kher Keep | land | Retains an existing utility land.
- 1 Krenko's Command | synergy | Supports the wide-board Goblin strategy.
- 1 Krenko, Mob Boss | threat | Provides a major Goblin threat around which to build pressure.
- 1 Mana Geyser | ramp | Supplies burst mana for an ambitious spell turn.
- 1 Mogg War Marshal | synergy | Maintains the precon's creature-building synergy package.
- 24 Mountain | land | Provides the dependable red mana foundation for the deck.
- 1 Pashalik Mons | removal | Keeps removal integrated with the Goblin theme.
- 1 Past in Flames | synergy | Supports the precon's spell-heavy storm engine.
- 1 Quest for the Goblin Lord | synergy | Maintains support for the Goblin board plan.
- 1 Redcap Gutter-Dweller | threat | Adds a substantial threat to the Goblin creature base.
- 1 Renegade Tactics | draw | Adds another draw spell to keep turns moving.
- 1 Roaming Throne | threat | Preserves a substantial threat from the precon.
- 1 Ruby Medallion | ramp | Supports mana efficiency across the spell-heavy plan.
- 1 Rundvelt Hordemaster | synergy | Keeps the Goblin synergy core intact.
- 1 Sazacap's Brew | draw | Provides card flow for assembling and sustaining spell turns.
- 1 Searslicer Goblin | synergy | Maintains another Goblin synergy creature.
- 1 Seething Song | ramp | Provides burst mana for the commander and follow-up plays.
- 1 Shinka, the Bloodsoaked Keep | land | Preserves the precon's land base.
- 1 Siege-Gang Commander | removal | Retains removal within the Goblin creature package.
- 1 Siege-Gang Lieutenant | removal | Keeps another Goblin-based removal option.
- 1 Skirk Prospector | ramp | Supplies early mana support while preserving the Goblin theme.
- 1 Skullclamp | draw | Provides card flow for the creature-heavy strategy.
- 1 Smoldering Crater | land | Preserves the precon's land base.
- 1 Sol Ring | ramp | Provides early acceleration toward the commander.
- 1 Spinerock Knoll | land | Keeps an existing land from the precon.
- 1 Spreading Insurrection | wincon | Preserves another payoff for a large spell turn.
- 1 Swiftfoot Boots | protection | Provides protection for an important creature.
- 1 Thought Vessel | ramp | Adds early mana development for more reliable setup turns.
- 1 Vandalblast | wipe | Provides a sweeping answer within the interaction package.
- 1 War Room | land | Preserves an existing utility land.
- 1 Witch's Mark | draw | Adds card flow to sustain the deck's spell plan.

</details>

### 18. upgrade a precon, any card

Format: Commander. Theme: turtles. Pool: any_card. Shortlist: 381 names.

Commander: Heroes in a Half Shell.

Grade: baseline, score 0.04, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $870.40 to buy, $870.40 the whole deck.

**Summary:** This five-color Turtle deck develops its mana, establishes a cast of Mutants, Ninjas, and Turtles, then turns successful attacks into growing creatures and fresh cards through its commander. Keep attacking to build an advantage, using removal and sweepers when opponents threaten to take over. Raphael, the Muscle, Dimension X Pizzasaur, and Everything Pizza supply finishing threats within the precon's character-driven plan. The deck favors sustained creature combat over compact combo wins, so it gives up some speed and remains dependent on maintaining a board and finding profitable attacks.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. The land count sits above the norm of the format, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The deck leans on five-color Mutant, Ninja and Turtle creatures with ramp and a few sweepers, but it also carries scattered cards like Fast Forward, Electric Seaweed, Wave Goodbye and Lessons from Life that don't clearly serve the attack-and-grow plan.
- PLAN theme_fit=yes: This is a five-color Heroes in a Half Shell Commander deck full of Turtle Power cards like the Mutant Ninja Turtle legends, Pizza cards and Mutant creatures, with Sol Ring, Chromatic Lantern and fixing lands as upgrades at bracket 3.
- PLAN useful_as_built=partly: The mana base is heavy on fixing and has plenty of rocks and dorks, so it should cast its spells, but the count of lands plus ramp is very high and the amount of real card advantage and ways to win is thin, with several low-impact cards like Exploding Barrel and Coin of Mastery.
- PLAN summary_honest=partly: The summary names Raphael, the Muscle, Dimension X Pizzasaur and Everything Pizza as finishers and they are all in the list, but Everything Pizza is a 2-mana artifact that does little to close a game, so calling it a finishing threat overstates what it does.
- [INFO] `curve_summary`: average mana value 3.33 over 60 nonland cards
- [WARN] `profile_off_band`: the land count is 39, and bracket 3 wants 34 to 38
- [WARN] `profile_off_band`: the ramp count is 7, and bracket 3 wants 8 to 13
- [WARN] `profile_off_band`: the draw count is 2, and bracket 3 wants 8 to 14
- [WARN] `profile_off_band`: the removal count is 2, and bracket 3 wants 6 to 12
- [WARN] `profile_off_band`: the interaction count is 2, and bracket 3 wants 4 to 12

<details><summary>The deck list</summary>

- 1 April O'Neil, Live on the Scene | draw | Provides card draw while preserving the deck's supporting cast.
- 1 Arcade Cabinet | synergy | Preserves the precon's arcade-themed support package.
- 1 Arcane Signet | ramp | Accelerates development and fixes the commander's five colors.
- 1 Assassin's Trophy | removal | Provides flexible removal for troublesome permanents.
- 1 Baxter, Fly in the Ointment | synergy | Adds a thematic Mutant to the commander's combat plan.
- 1 Bebop, Skull & Crossbones | threat | Keeps a thematic Mutant in the attacking creature package.
- 1 Big Mother Mouser | threat | Preserves the precon's Robot creature package.
- 1 Biogenic Ooze | threat | Maintains the precon's Ooze creature package.
- 1 Birds of Paradise | ramp | Adds inexpensive acceleration and five-color fixing.
- 1 Blasphemous Act | wipe | Provides a reset when opposing creature boards become overwhelming.
- 1 Chromatic Lantern | ramp | Supports reliable access to all five colors.
- 1 Coin of Mastery | synergy | Keeps the precon's thematic artifact support intact.
- 1 Continue? | other | Preserves the precon's supporting instant package.
- 1 Corpsejack Menace | synergy | Supports the commander's +1/+1-counter strategy.
- 1 Cultivate | ramp | Develops mana using the retained basic lands.
- 1 Dimension X Pizzasaur | wincon | Provides a thematic finisher that also belongs to the Mutant package.
- 1 Donatello, the Brains | synergy | Keeps a core Turtle in the commander's combat engine.
- 1 Double Jump // Flying Kick | interaction | Preserves the precon's combat-themed spell package.
- 1 Electric Seaweed | other | Maintains the precon's supporting creature package.
- 1 Endless Foot Assault | synergy | Keeps the precon's thematic enchantment package intact.
- 1 Everything Pizza | wincon | Retains a thematic finisher alongside the creature-based threats.
- 1 Exploding Barrel | other | Preserves the precon's arcade-themed artifact package.
- 1 Fast Forward | other | Keeps the precon's supporting sorcery package intact.
- 1 Fellwar Stone | ramp | Adds early acceleration and additional color fixing.
- 1 Foot Chopper | synergy | Maintains the precon's equipment package.
- 1 Game Over | other | Preserves the precon's thematic sorcery package.
- 1 Harmonize | draw | Replenishes cards independently of successful combat.
- 1 Here Comes a New Hero! | synergy | Keeps the precon's hero-themed support package intact.
- 1 High Score | synergy | Preserves the precon's arcade-themed enchantment package.
- 1 Irma, Part-Time Mutant | synergy | Adds a thematic Mutant to the commander's creature package.
- 1 Krang, the All-Powerful | threat | Keeps a signature character in the deck's threat package.
- 1 Leatherhead, Iron Gator | threat | Adds another thematic Mutant for the combat plan.
- 1 Leonardo, the Balance | threat | Supplies a core Turtle threat for repeated attacks.
- 1 Lessons from Life | other | Preserves the precon's supporting sorcery package.
- 1 Level Up | synergy | Maintains the precon's creature-focused Aura package.
- 1 Lita, Little Orphan Amphibian | synergy | Keeps another Turtle connected to the commander's combat engine.
- 1 Michelangelo, the Heart | synergy | Preserves a core Turtle in the deck's synergy package.
- 1 Mole Module | synergy | Maintains the precon's Vehicle package.
- 1 Mona Lisa, Science Geek | synergy | Adds a thematic Mutant to the creature package.
- 1 Ninja Pizza | synergy | Keeps the precon's thematic enchantment package intact.
- 1 Raphael, the Muscle | threat | Provides a core Turtle threat and a creature-based finisher.
- 1 Rat King, Pale Piper | threat | Preserves a signature character in the creature package.
- 1 Ray Fillet, Wave Warrior | threat | Adds another thematic Mutant to the attacking roster.
- 1 Roadkill Rodney | threat | Maintains the precon's Robot creature package.
- 1 Rocksteady, Mutant Marauder | threat | Keeps a thematic Mutant in the deck's threat package.
- 1 Shellshock | other | Preserves the precon's supporting instant package.
- 1 Shredder, Shadow Master | threat | Adds a signature Ninja to the commander's combat plan.
- 1 Sol Ring | ramp | Accelerates deployment of the commander and supporting permanents.
- 1 Special Move | interaction | Keeps the precon's combat-themed instant package intact.
- 1 Splinter, the Mentor | synergy | Adds a thematic Mutant Ninja to the commander's creature package.
- 1 Steelbane Hydra | removal | Provides removal while retaining a Turtle creature.
- 1 Super Combo | other | Preserves the precon's arcade-themed sorcery package.
- 1 Swift Demise | other | Keeps the precon's supporting instant package intact.
- 1 Tempestra, Dame of Games | threat | Maintains a signature character in the creature package.
- 1 Together Forever | synergy | Preserves support for the deck's counter-focused creatures.
- 1 Tokka & Rahzar, Unsupervised | ramp | Provides thematic acceleration on a Mutant creature.
- 1 Vanquish the Horde | wipe | Provides another answer to overwhelming creature boards.
- 1 Vigor | protection | Supports the deck's creature-combat and counter strategy.
- 1 Voracious Hydra | threat | Maintains a scalable creature in the precon's threat package.
- 1 Wave Goodbye | other | Preserves the precon's supporting sorcery package.
- 1 Ash Barrens | land | Provides a retained route to the basic colors.
- 1 Big Apple, 3 a.m. | land | Keeps the precon's thematic mana-base foundation.
- 1 Cinder Glade | land | Retains red-green fixing.
- 1 City of Brass | land | Provides flexible access to all five colors.
- 1 Command Tower | land | Supplies whichever commander color is needed.
- 1 Dragonskull Summit | land | Retains black-red fixing.
- 1 Escape Tunnel | land | Keeps a basic-land fixing option in the mana base.
- 1 Exotic Orchard | land | Provides flexible multicolor fixing.
- 1 Fabled Passage | land | Finds a retained basic land to complete the needed colors.
- 1 Hidden Hideout | land | Preserves a thematic land from the precon.
- 1 Hinterland Harbor | land | Retains green-blue fixing.
- 1 Path of Ancestry | land | Keeps five-color fixing suited to the creature theme.
- 1 Rootbound Crag | land | Retains red-green fixing.
- 1 Smoldering Marsh | land | Retains black-red fixing.
- 1 Spire Garden | land | Supports red-green development.
- 1 Sunken Hollow | land | Retains blue-black fixing.
- 1 Turtle Lair | land | Keeps the precon's Turtle-themed mana-base support.
- 1 Undergrowth Stadium | land | Supports black-green development.
- 1 Vibrant Cityscape | land | Preserves part of the precon's existing mana base.
- 1 Blood Crypt | land | Adds black-red fixing with an untapped option.
- 1 Breeding Pool | land | Adds green-blue fixing with an untapped option.
- 1 Godless Shrine | land | Adds white-black fixing with an untapped option.
- 1 Hallowed Fountain | land | Adds white-blue fixing with an untapped option.
- 1 Overgrown Tomb | land | Adds black-green fixing with an untapped option.
- 1 Sacred Foundry | land | Adds red-white fixing with an untapped option.
- 1 Steam Vents | land | Adds blue-red fixing with an untapped option.
- 1 Stomping Ground | land | Adds red-green fixing with an untapped option.
- 1 Temple Garden | land | Adds green-white fixing with an untapped option.
- 1 Watery Grave | land | Adds blue-black fixing with an untapped option.
- 1 Mana Confluence | land | Provides untapped access to every commander color.
- 1 Cavern of Souls | land | Supports colored mana for the central creature theme.
- 1 Secluded Courtyard | land | Adds untapped fixing for the central creature theme.
- 1 Unclaimed Territory | land | Adds another tribal fixing source.
- 1 Tropical Island | land | Provides untapped green-blue fixing for early development.
- 1 Forest | land | Provides a basic green source for the land-search package.
- 1 Island | land | Provides a basic blue source for the land-search package.
- 1 Swamp | land | Provides a basic black source for the land-search package.
- 1 Mountain | land | Provides a basic red source for the land-search package.
- 1 Plains | land | Provides a basic white source for the land-search package.

</details>

### 19. the Hobbit family, two colours

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 174 names.

Commander: Thranduil, the Elvenking.

Grade: bad, score 0.04, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $335.42 to buy, $335.42 the whole deck.

**Summary:** Thranduil leads an Elf-focused midrange deck that develops mana, builds a creature presence, and uses legendary Elves to keep cards flowing. Removal and battlefield resets help it reach a longer game, where its Elf board and finishers—chief among them Witch-king of Angmar and Troll of Khazad-dûm—provide the closing pressure. It favors thematic creature development over explosive wins, while the basic-heavy mana base makes color-intensive draws less dependable.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. The cards pair in ways the top lists do not, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The Elf creature core, ramp and card draw agree with Thranduil, but the black removal, two wipes and Witch-king and Troll finishers pull away from an Elf board plan, and Languish and Gnashing of Teeth would wreck the Elves.
- PLAN theme_fit=yes: The deck is led by Thranduil, the Elvenking in UBG, uses only cards from the Hobbit sets, and is built around Elves and Middle-earth legends as the request asks.
- PLAN useful_as_built=partly: It has about 36 lands, enough ramp and draw, and a reasonable curve, but 30 basics across three colors with BB and UU costs strain the mana, and the finishers are modest, so it works but is clunky.
- PLAN summary_honest=yes: The summary says Witch-king of Angmar and Troll of Khazad-dûm are finishers, and both are in the list as wincons, and it also correctly admits the basic-heavy mana base is a weakness.
- [INFO] `curve_summary`: average mana value 3.15 over 61 nonland cards
- [WARN] `profile_off_band`: the count of lands that make two or more of the deck's colors is 4, and bracket 3 wants 17 or more (Elven Passage, Elvenking's Halls, Hobbit Hole, Mirkwood)
- [INFO] `mana_pass`: the builder moved 4 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 10 Forest | land | Provides green mana for the Elf-heavy creature base.
- 9 Island | land | Provides blue mana for card draw and interaction.
- 11 Swamp | land | Provides black mana for removal and finishing threats.
- 1 Elven Passage | land | Supports mana development while preserving the woodland theme.
- 1 Elvenking's Halls | land | Adds a thematic land to Thranduil's mana base.
- 1 Hobbit Hole | land | Adds another land for consistent mana development.
- 1 Minas Morgul, Dark Fortress | land | Contributes a land to support the deck's larger threats.
- 1 Mirkwood | land | Supports the mana base and the woodland setting.
- 1 Rivendell | land | Adds an Elven location to the mana base.
- 1 The Black Gate | land | Provides another land for sustained mana development.
- 1 The Shire | land | Rounds out the mana base with another thematic land.
- 1 Arcane Signet | ramp | Adds artifact-based mana acceleration.
- 1 Delighted Halfling | ramp | Adds creature-based acceleration alongside the legendary creatures.
- 1 Elvish Mystic | ramp | Combines mana acceleration with the deck's Elf theme.
- 1 Mox Amber | ramp | Adds mana acceleration to a deck containing numerous legends.
- 1 Elvish Archdruid | ramp | Provides acceleration while increasing the Elf presence.
- 1 Wood Elves | ramp | Supports mana development without sacrificing Elf density.
- 1 Woodland Weavemaster | ramp | Adds another Elf to the mana-development package.
- 1 Silvan Reveler | ramp | Combines the tribal creature plan with acceleration.
- 1 Thranduil the Strategist | ramp | Provides ramp and a legendary Elf for the commander's draw trigger.
- 1 Thranduil's Company | ramp | Adds Elf-based acceleration to support the larger threats.
- 1 Elven Chorus | ramp | Diversifies the ramp package with an enchantment.
- 1 Wayfarer's Bauble | ramp | Provides another artifact-based route to mana development.
- 1 Elvish Visionary | draw | Adds card draw while keeping the creature base focused on Elves.
- 1 Night's Whisper | draw | Provides a dedicated source of card draw.
- 1 Hithlain Knots | draw | Adds instant-speed card draw to the resource package.
- 1 Lórien Revealed | draw | Supplies card draw for longer games.
- 1 Palantír of Orthanc | draw | Adds artifact-based card draw.
- 1 Captain of Umbar | draw | Adds a creature-based source of card draw.
- 1 Fateful Discovery | draw | Diversifies card draw with an enchantment.
- 1 Uncover the Moon-Letters | draw | Adds another enchantment to the card-draw package.
- 1 Key to the Side-Door | draw | Provides card draw from an artifact slot.
- 1 Gollum, Riddle Master | draw | Supports card draw while also contributing finishing pressure.
- 1 Ithilien Kingfisher | draw | Adds another creature to the card-draw package.
- 1 Confusticate and Bebother | interaction | Provides instant-speed interaction against opposing plans.
- 1 Stern Scolding | interaction | Adds another instant to the interaction package.
- 1 Thranduil's Decree | interaction | Supports the deck with thematic instant-speed interaction.
- 1 Warg Tactics | interaction | Rounds out the instant-speed interaction package.
- 1 Bitter Downfall | removal | Provides instant-speed removal for opposing threats.
- 1 Bilbo's Deadly Slice | removal | Adds another instant-speed removal option.
- 1 Enchanted River's Grasp | removal | Diversifies removal with an Aura.
- 1 Merciless Executioner | removal | Adds removal attached to a creature.
- 1 Orcish Bowmasters | removal | Provides another creature-based removal option.
- 1 Stir Up Trouble | removal | Adds a sorcery to the removal package.
- 1 Uneasy Partings | removal | Provides an additional instant-speed answer.
- 1 Colossal Whale | removal | Combines removal with a substantial finishing threat.
- 1 Galion, Elvenking's Butler | synergy | Adds a legendary Elf to support Thranduil's draw trigger.
- 1 Arwen, Weaver of Hope | synergy | Strengthens the legendary Elf core around Thranduil.
- 1 Celeborn the Wise | synergy | Provides another legendary Elf for the commander's draw engine.
- 1 Elrond, Moon-Reader | synergy | Combines the shortlist's protection role with legendary Elf synergy.
- 1 Grey Havens Navigator | synergy | Keeps the creature package concentrated on Elves.
- 1 Mirkwood Meditator | synergy | Adds another Elf to the commander's tribal shell.
- 1 Supper for Spiders | synergy | Supplies a dedicated synergy spell alongside the creature core.
- 1 Boughside Wanderers | threat | Adds an Elf body to develop the battlefield.
- 1 Cantankerous Keepers | threat | Expands the Elf creature presence for the combat plan.
- 1 Elven Raft-Steerer | threat | Adds another Elf to the deck's battlefield development.
- 1 Elvenking's Harper | threat | Maintains Elf density in the creature package.
- 1 Galadhrim Guide | threat | Contributes an Elf creature to the combat plan.
- 1 Guardian of the Halls | threat | Adds another Elf Soldier to the board.
- 1 Haunt of the Dead Marshes | threat | Adds a Nightmare Elf to the tribal creature base.
- 1 Lothlórien Lookout | threat | Supports battlefield development with another Elf Scout.
- 1 Mirkwood Nurturer | threat | Adds an Elf Ranger to the creature-based game plan.
- 1 Mirkwood Pathmaker | threat | Keeps the combat creature package centered on Elves.
- 1 Nimrodel Watcher | threat | Provides another Elf body for battlefield presence.
- 1 Thranduil, Sindarin Liege // Silvan Rally | threat | Adds a legendary Elf threat that supports the commander's draw trigger.
- 1 Troll of Khazad-dûm | wincon | Provides a dedicated finishing threat.
- 1 Witch-king of Angmar | wincon | Adds a legendary finisher to close out developed games.
- 1 Witch-king, Bringer of Ruin | wincon | Supplies another dedicated finishing threat.
- 1 Gnashing of Teeth | wipe | Provides a battlefield reset when opposing boards become overwhelming.
- 1 Languish | wipe | Adds another sweeper for recovering from unfavorable boards.
- 1 Raise the Palisade | wipe | Rounds out the deck's battlefield-reset options.
- 1 Pelargir Survivor | ramp | the mana pass added it to bring the mana base inside the power level

</details>

### 20. the Hobbit family, a delegated commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 77 names.

Commander: Smaug the Magnificent.

Grade: baseline, score 0.07, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $428.55 to buy, $428.55 the whole deck.

**Summary:** Your commander leads a Dragon-focused deck that develops mana, builds up Treasure for powerful attacks, and uses removal and protective artifacts to keep its key creatures relevant. Dwarves, Goblins, and Equipment supply the supporting battlefield presence before Cavern-Hoard Dragon, Desert Were-Worm, and Smaug, the Great Calamity // Spew Flame close the game. The plan favors steady mana and combat over explosive combinations, but gives up a dense Dragon roster and plentiful card draw for its thematic supporting cast.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The mana base serves the colors evenly, and that raises the grade. The cards pair in ways the top lists do not, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: Mono-red Treasure, Dwarves, Goblins and Equipment all lean toward a loose red-aggro pile, but the many singleton Equipment and Goblin/Dwarf bodies do not clearly feed the Dragon or Treasure payoffs.
- PLAN theme_fit=partly: It is mono-red, all from the Hobbit sets, and led by Smaug the Magnificent, but with only four Dragons including the commander it barely delivers the dragons deck that was requested.
- PLAN useful_as_built=yes: With 38 lands, mostly Mountains, plus a spread of cheap ramp, removal and creatures, the deck casts its spells smoothly and has a few big finishers, so it can be played as it stands.
- PLAN summary_honest=yes: The summary's claim that the closers are Cavern-Hoard Dragon, Desert Were-Worm and Smaug, the Great Calamity matches the wincon cards in the list, and it openly admits the thin Dragon count and low card draw.
- [INFO] `curve_summary`: average mana value 3.20 over 61 nonland cards
- [WARN] `profile_off_band`: the draw count is 6, and bracket 3 wants 8 to 14

<details><summary>The deck list</summary>

- 34 Mountain | land | Provides a dependable red mana base for the commander and supporting spells.
- 1 Dragon-Cursed Halls | land | Adds a thematic nonbasic land without crowding out red sources.
- 1 The Lonely Mountain | land | Supports the mana base while reinforcing the Dragon theme.
- 1 Rogue's Passage | land | Provides a utility land slot for the creature-focused plan.
- 1 Treasure Vault | land | Adds an artifact land to the deck's supporting mana base.
- 1 Arcane Signet | ramp | Provides inexpensive acceleration toward the commander and Dragons.
- 1 Bag End Banquet | ramp | Adds another source of acceleration for the larger threats.
- 1 Burn, Burn, Tree and Fern | ramp | Supplies enchantment-based ramp for the deck's mana development.
- 1 Dragon's Desire | ramp | Supports the mana needed to deploy the Dragon finishers.
- 1 Fíli and Kíli, Joyous | ramp | Combines a supporting Dwarf creature with mana development.
- 1 Glóin the Mighty // Easy Pickings | ramp | Adds a Dwarf ramp option to support the larger spells.
- 1 Mox Amber | ramp | Provides a low-cost mana piece alongside the legendary creatures.
- 1 Orcrist, Goblin-cleaver | ramp | Contributes Equipment-based acceleration to the supporting package.
- 1 The Misty Mountains Cold | ramp | Adds another ramp option for sustained mana development.
- 1 The Reaver Cleaver | ramp | Connects the Equipment package with the deck's acceleration plan.
- 1 Thorin, Company's Leader | ramp | Provides ramp while also adding a marked finisher to the creature suite.
- 1 Wayfarer's Bauble | ramp | Adds an inexpensive ramp piece to help establish mana early.
- 1 Balin, Loremaster | draw | Provides card draw on a thematic supporting creature.
- 1 Key to the Side-Door | draw | Adds artifact-based card draw to sustain the deck's resources.
- 1 Palantír of Orthanc | draw | Provides another draw option for finding threats and answers.
- 1 Ragged Short Spear | draw | Combines the Equipment package with additional card draw.
- 1 Thrór's Map | draw | Adds a draw piece that fits the deck's supporting artifact theme.
- 1 Óin the Brave | draw | Supplies card draw while maintaining the Dwarf supporting cast.
- 1 Battle-Scarred Goblin | removal | Adds creature-based removal to help clear opposing threats.
- 1 Fire of Orthanc | removal | Provides a removal spell for troublesome opposing permanents.
- 1 Gandalf, Spark Starter | removal | Adds removal on a legendary supporting creature.
- 1 Giant's Boulder | removal | Provides an artifact-based removal option.
- 1 Goblin Fireleaper | removal | Adds another removal creature to the supporting army.
- 1 Inferno Titan | removal | Combines a substantial creature with a removal role.
- 1 Pinecone Strike | removal | Provides an instant-speed removal option.
- 1 Thorin, Mountain-king | removal | Adds another thematic creature that serves as removal.
- 1 Bilbo's Ring | interaction | Provides protection for an important creature through the Equipment package.
- 1 Dwarven Mattock | interaction | Adds another protective Equipment option for key creatures.
- 1 Mithril Coat | interaction | Protects a key creature while fitting the artifact support package.
- 1 The One Ring | interaction | Adds a protective artifact to help preserve the deck's position.
- 1 Goblin Cratermaker | interaction | Provides a creature-based answer to opposing threats.
- 1 Improvised Club | interaction | Adds an instant answer for a timely exchange.
- 1 Smite the Deathless | interaction | Provides another instant answer to an opposing threat.
- 1 Smaug's Fury | interaction | Adds a thematic instant to the deck's tactical spell package.
- 1 Bombur, Gentle Dreamer | threat | Adds a supporting Dwarf body to the creature plan.
- 1 Bothersome Noisemaker | threat | Provides another Goblin creature for the supporting battlefield presence.
- 1 Dori, Bearer of Friends | threat | Adds a legendary Dwarf to the supporting creature suite.
- 1 Dwarven Mauler | threat | Provides a combat-focused Dwarf threat beneath the Dragon finishers.
- 1 Dwarven Warriors | threat | Adds another supporting creature to carry the combat plan.
- 1 Dáin Ironfoot | threat | Adds a marked finisher that complements the Dragon threats.
- 1 Goblin-town Flunkies | threat | Contributes another Goblin body to the deck's supporting army.
- 1 Gundabad Opportunist | threat | Adds another creature threat below the top-end Dragons.
- 1 Iron Hills Stalwart | threat | Reinforces the Dwarf portion of the creature suite.
- 1 Misty Mountains Raider | threat | Adds a Goblin creature to maintain battlefield pressure.
- 1 Orcish Siegemaster | threat | Diversifies the supporting army with an Orc threat.
- 1 Snowslope Hunter | threat | Provides another supporting Goblin creature for the combat plan.
- 1 Andúril, Flame of the West | synergy | Adds Equipment support for the creature-heavy strategy.
- 1 Getaway Barrel | synergy | Rounds out the deck's supporting artifact package.
- 1 Glamdring | synergy | Provides another Equipment option for the supporting creatures.
- 1 Guttersnipe | synergy | Adds a Goblin Shaman alongside the deck's supporting spell package.
- 1 Last Light of Durin's Day | synergy | Adds a dedicated synergy enchantment to the supporting engine.
- 1 Long-Lost Lances | synergy | Expands the Equipment package for the creature-focused plan.
- 1 Old Thrush | synergy | Adds a small supporting creature alongside the Dragons and their allies.
- 1 Sting, Bilbo's Sword | synergy | Provides another thematic Equipment piece for the supporting army.
- 1 The Black Arrow | synergy | Adds a dedicated synergy piece within the Equipment package.
- 1 Well-Worn Spatula | synergy | Rounds out the Equipment support for the deck's creatures.
- 1 Cavern-Hoard Dragon | wincon | Provides a Dragon finisher for closing out the game.
- 1 Desert Were-Worm | wincon | Adds a Dragon Wurm finisher to diversify the top-end threats.
- 1 Smaug, the Great Calamity // Spew Flame | wincon | Provides a thematic Dragon finisher for the deck's closing turns.
- 1 Call Forth the Tempest | wipe | Provides a board wipe when opposing battlefields become overwhelming.
- 1 Desolation of Smaug | wipe | Adds a second board wipe that reinforces the Dragon theme.

</details>

### 21. the Hobbit family, mana from outside

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 156 names.

Commander: Smaug the Impenetrable.

Grade: baseline, score 0.05, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $501.88 to buy, $501.88 the whole deck.

**Summary:** Smaug leads a creature-heavy midrange deck that develops its mana, keeps cards flowing, and uses removal to make room for combat pressure. Goblins and thematic Equipment support the early and middle turns, while the Dragon finishers and Witch-king provide the closing threats. Smaug's noncombat-damage Treasure engine can help fuel the late game. The deck favors a broad Hobbit-themed battlefield over a compact combo finish, giving up some speed and protection in exchange for sustained creature pressure.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. The cards pair in ways the top lists do not, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: Most cards are Goblin and Hobbit creatures, equipment and removal that point toward midrange combat, but the board wipes (Languish, Desolation of Smaug, Gnashing of Teeth) work against a deck full of small creatures, and several cards (Verdant Catacombs, Exotic Orchard) do not fit the plan.
- PLAN theme_fit=yes: Smaug the Impenetrable leads a black-red deck built almost entirely from Hobbit-set cards, with lands and rocks from other sets as the request allows.
- PLAN useful_as_built=partly: The deck has plenty of two-drops, ramp and draw, but the mana base has off-color fetches (Verdant Catacombs, Marsh Flats) and Exotic Orchard, plus a tail of weak cards, so it runs rough though it can be played.
- PLAN summary_honest=partly: The Goblins, Equipment and Dragon finishers in the summary are in the list, but the claim that removal 'makes room for combat pressure' leaves out that three sweepers cut against the creature-heavy plan.
- [INFO] `curve_summary`: average mana value 3.11 over 61 nonland cards
- [INFO] `mana_pass`: the builder moved 4 cards of the mana base to bring the deck inside its power level
- [WARN] `outside_requested_set`: the sets you named do not hold 14 cards: Blood Crypt, Bloodstained Mire, Command Tower, Dragonskull Summit, Evolving Wilds, Exotic Orchard, Fabled Passage, Marsh Flats, Path of Ancestry, Polluted Delta, Scalding Tarn, Terramorphic Expanse, Urborg, Tomb of Yawgmoth, Verdant Catacombs

<details><summary>The deck list</summary>

- 1 Blood Crypt | land | Provides both colors for the deck's spells.
- 1 Bloodstained Mire | land | Finds the appropriate colored land.
- 1 Polluted Delta | land | Finds a black source or Blood Crypt.
- 1 Marsh Flats | land | Finds a black source or Blood Crypt.
- 1 Verdant Catacombs | land | Finds a black source or Blood Crypt.
- 1 Scalding Tarn | land | Finds a red source or Blood Crypt.
- 1 Command Tower | land | Provides flexible colored mana.
- 1 Dragonskull Summit | land | Supports both halves of the mana base.
- 1 Exotic Orchard | land | Adds another flexible colored source.
- 1 Mount Doom | land | Supplies both of the deck's colors.
- 1 Path of Ancestry | land | Adds another source of either color.
- 1 Urborg, Tomb of Yawgmoth | land | Strengthens access to black mana.
- 11 Mountain | land | Provides dependable untapped red mana.
- 11 Swamp | land | Provides dependable untapped black mana.
- 1 Arcane Signet | ramp | Provides inexpensive acceleration and color fixing.
- 1 Mox Amber | ramp | Adds a low-cost mana source alongside the legendary creatures.
- 1 Bag End Banquet | ramp | Adds artifact-based acceleration to the thematic core.
- 1 Bolg's Company | ramp | Combines a Goblin presence with the deck's ramp package.
- 1 Burn, Burn, Tree and Fern | ramp | Adds thematic acceleration through a Saga.
- 1 Dragon's Desire | ramp | Supports the mana needed for the Dragon endgame.
- 1 Glóin the Mighty // Easy Pickings | ramp | Adds a thematic creature to the acceleration package.
- 1 Long-Bodied Grey Dog | ramp | Provides another creature-based ramp option.
- 1 Orcrist, Goblin-cleaver | ramp | Adds Equipment-based acceleration.
- 1 The Reaver Cleaver | ramp | Connects the Equipment package with mana development.
- 1 The Misty Mountains Cold | ramp | Adds another thematic source of acceleration.
- 1 Troop of Ponies | ramp | Rounds out the creature-based mana support.
- 1 Night's Whisper | draw | Provides inexpensive card draw.
- 1 Palantír of Orthanc | draw | Adds an artifact-based source of card advantage.
- 1 Key to the Side-Door | draw | Supports the deck's card-draw package.
- 1 Ragged Short Spear | draw | Adds card draw through Equipment.
- 1 Thrór's Map | draw | Provides another thematic artifact for card advantage.
- 1 Rage into the Valley | draw | Helps replenish resources for the midgame.
- 1 The Master of Lake-town | draw | Adds a legendary creature to the draw package.
- 1 The Sackville-Bagginses | draw | Provides thematic creature-based card advantage.
- 1 Gollum, Riddle Master | draw | Supports card advantage while adding a marked finisher.
- 1 Óin the Brave | draw | Adds another legendary source of card draw.
- 1 Bilbo's Deadly Slice | interaction | Provides instant-speed interaction.
- 1 Bitter Downfall | interaction | Adds an instant answer to opposing threats.
- 1 Smite the Deathless | interaction | Adds another instant removal option.
- 1 Improvised Club | interaction | Provides a red instant for answering threats.
- 1 Pinecone Strike | interaction | Expands the instant-speed interaction package.
- 1 Fire of Orthanc | interaction | Adds a thematic removal spell.
- 1 Goblin Cratermaker | interaction | Places an interactive option on a creature.
- 1 Orcish Bowmasters | removal | Adds creature-based removal to the board.
- 1 Goblin Fireleaper | removal | Combines a Goblin body with a removal role.
- 1 Merciless Executioner | removal | Provides removal through a creature.
- 1 Azog, Moria's Ruin | removal | Adds a legendary Goblin to the removal package.
- 1 Bolg of the North | removal | Supports the Goblin theme while answering threats.
- 1 Bolg, Erebor's Reckoning | removal | Adds another thematic removal creature.
- 1 Gandalf, Spark Starter | removal | Provides a legendary creature with a removal role.
- 1 Giant's Boulder | removal | Adds artifact-based removal.
- 1 Supper for Spiders | synergy | Provides a dedicated thematic synergy spell.
- 1 Smaug's Fury | synergy | Reinforces the Smaug-centered theme.
- 1 The Black Arrow | synergy | Adds thematic Equipment to the creature package.
- 1 Crude Bent Blade | synergy | Supports the deck's Equipment subtheme.
- 1 Goblin Plate Mail | synergy | Connects the Goblin and Equipment themes.
- 1 Bothersome Noisemaker | synergy | Adds another Goblin to the supporting creature package.
- 1 Stony-Voiced Goblins | synergy | Reinforces the Goblin-heavy thematic core.
- 1 Goblin-town Flunkies | threat | Adds a Goblin body for battlefield pressure.
- 1 Fearsome Goblin Pair | threat | Expands the deck's creature pressure.
- 1 Great Goblin, Foul-Hearted | threat | Adds a thematic legendary threat.
- 1 The Great Goblin | threat | Provides another legendary Goblin threat.
- 1 Gundabad Opportunist | threat | Adds a Goblin threat to the creature curve.
- 1 Misty Mountains Raider | threat | Reinforces the attacking Goblin package.
- 1 Olog-hai Crusher | threat | Adds a substantial creature to the threat package.
- 1 Ravening Warg | threat | Adds a thematic Wolf threat.
- 1 Desolation Prowler | threat | Diversifies the deck's creature pressure.
- 1 Sauron, the Lidless Eye | threat | Adds a legendary threat beyond the Goblin package.
- 1 Troll of Khazad-dûm | threat | Provides a marked finisher among the creature threats.
- 1 Smaug, Wicked Worm | wincon | Provides a marked Dragon finisher.
- 1 Smaug, the Great Calamity // Spew Flame | wincon | Combines a marked finisher with a removal role.
- 1 Witch-king of Angmar | wincon | Adds a marked finisher outside the Dragon package.
- 1 Desolation of Smaug | wipe | Provides a thematic battlefield reset.
- 1 Gnashing of Teeth | wipe | Adds another option for clearing a crowded battlefield.
- 1 Languish | wipe | Provides a midgame battlefield reset.
- 1 Fabled Passage | land | the mana pass traded a basic land for this land, which makes two or more of the deck's colors
- 1 Evolving Wilds | land | the mana pass traded a basic land for this land, which makes two or more of the deck's colors
- 1 Terramorphic Expanse | land | the mana pass traded a basic land for this land, which makes two or more of the deck's colors
- 1 Hobbit Hole | land | the mana pass traded a basic land for this land, which makes two or more of the deck's colors

</details>

### 22. a set family and a card from outside it

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 175 names.

Commander: Thranduil, the Elvenking.

Grade: baseline, score 0.06, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: yes, for deck_size, profile_off_band, profile_off_band. Block findings: 0.

Cost: $372.08 to buy, $372.08 the whole deck.

**Summary:** An Elf-centered Sultai deck built around legendary-creature value, graveyard utility, and steady combat pressure. Mana acceleration and card advantage sustain development, while interaction and protective equipment support powerful evasive finishers. The supplied shortlist cannot satisfy the multicolor-land requirement under the singleton restriction, so full compliance is not possible.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. The cards pair in ways the top lists do not, and that lowers the grade.

Summary rules claims (F-26): the format play, singleton

- JUDGE [unknown]: "The supplied shortlist cannot satisfy the multicolor-land requirement under the singleton restriction, so full compliance is not possible.". It asserts a deckbuilding requirement (a multicolor-land requirement and a singleton restriction) as a rule. Singleton is a real Commander rule, but no multicolor-land requirement exists in the game rules. The card list also has 29 basic lands, which breaks the singleton rule, so the statement mixes accurate and inaccurate rule content.
- PLAN plan_coherent=partly: The Elf-tribal creatures, ramp, draw, and removal broadly fit a Sultai Elf value plan, but there is a lot of generic filler, three Wraith/Orc/Troll black cards, and few clear payoffs tying the pieces together.
- PLAN theme_fit=yes: The deck is led by Thranduil, the Elvenking in his Sultai colors, is drawn almost entirely from the Hobbit sets, and includes the requested Sol Ring.
- PLAN useful_as_built=partly: The 37 lands plus plenty of ramp and draw make it playable, but the 29 basics with few dual lands make three-color mana shaky, and the finishers are modest, with few evasive win conditions despite the summary.
- PLAN summary_honest=partly: The summary claims a graveyard utility theme that the list barely carries, though the Elf and legendary-creature value claims hold up, and the note about the multicolor-land shortfall is an honest admission.
- [INFO] `curve_summary`: average mana value 3.00 over 63 nonland cards
- [WARN] `profile_off_band`: the count of lands that make two or more of the deck's colors is 4, and bracket 3 wants 17 or more (Elven Passage, Elvenking's Halls, Hobbit Hole, Mirkwood)
- [WARN] `outside_requested_set`: the sets you named do not hold 1 card: Sol Ring
- [WARN] `summary_rules_claim`: the summary states a rule of the game, "singleton". The engine reports the rules, and the summary must not.

<details><summary>The deck list</summary>

- 9 Forest | land | Green mana for Elves and early development.
- 9 Island | land | Blue mana for card selection and interaction.
- 11 Swamp | land | Black mana for removal and late-game threats.
- 1 Elven Passage | land | Multicolor fixing.
- 1 Elvenking's Halls | land | Multicolor fixing.
- 1 Mirkwood | land | Multicolor fixing.
- 1 Rivendell | land | Blue source with utility.
- 1 The Black Gate | land | Black source with combat utility.
- 1 The Shire | land | Green source with utility.
- 1 Hobbit Hole | land | Multicolor fixing.
- 1 Sol Ring | ramp | Efficient acceleration into the commander and larger threats.
- 1 Arcane Signet | ramp | Acceleration and color fixing.
- 1 Elvish Mystic | ramp | Early acceleration on an Elf.
- 1 Delighted Halfling | ramp | Early acceleration supporting legendary creatures.
- 1 Wayfarer's Bauble | ramp | Basic-land acceleration.
- 1 Wood Elves | ramp | Land acceleration on an Elf.
- 1 Woodland Weavemaster | ramp | Elf-based mana development.
- 1 Silvan Reveler | ramp | Mana development supporting the Elf plan.
- 1 Thranduil the Strategist | ramp | Legendary Elf supporting mana development.
- 1 Thranduil's Company | ramp | Elf-based acceleration.
- 1 Pelargir Survivor | ramp | Additional mana support.
- 1 Necklace of Girion | ramp | Artifact-based mana development.
- 1 Elvish Visionary | draw | Card flow attached to an Elf.
- 1 Night's Whisper | draw | Efficient card replenishment.
- 1 Hithlain Knots | draw | Card flow with tempo utility.
- 1 Lórien Revealed | draw | Card replenishment with land-finding flexibility.
- 1 Palantír of Orthanc | draw | Repeatable card advantage.
- 1 Captain of Umbar | draw | Card selection to support the graveyard plan.
- 1 Old Fat Spider | draw | Card advantage and a substantial threat.
- 1 Gollum, Riddle Master | draw | Card advantage with finishing potential.
- 1 Ithilien Kingfisher | draw | Card replenishment on a creature.
- 1 Key to the Side-Door | draw | Artifact-based card advantage.
- 1 Thrór's Map | draw | Additional artifact-based card flow.
- 1 Stern Scolding | interaction | Efficient stack interaction.
- 1 Confusticate and Bebother | interaction | Disrupts opposing plays.
- 1 Thranduil's Decree | interaction | Interaction supporting the Elf strategy.
- 1 Warg Tactics | interaction | Tactical interaction.
- 1 Bitter Downfall | removal | Answers opposing creatures.
- 1 Bilbo's Deadly Slice | removal | Targeted removal.
- 1 Enchanted River's Grasp | removal | Neutralizes an opposing threat.
- 1 Merciless Executioner | removal | Creature-based sacrifice removal.
- 1 Orcish Bowmasters | removal | Small-creature removal and pressure.
- 1 Uneasy Partings | removal | Flexible removal.
- 1 Arwen, Weaver of Hope | synergy | Legendary Elf supporting creature growth.
- 1 Celeborn the Wise | synergy | Legendary Elf supporting the tribal plan.
- 1 Galion, Elvenking's Butler | synergy | Legendary Elf supporting commander value.
- 1 Thranduil, Sindarin Liege // Silvan Rally | synergy | Legendary Elf with tribal utility.
- 1 Elrond, Moon-Reader | synergy | Legendary Elf supporting the board.
- 1 Elvenking's Harper | synergy | Elf utility for the commander strategy.
- 1 Supper for Spiders | synergy | Supports the deck's creature and graveyard plan.
- 1 Mox Amber | synergy | Mana support alongside legendary creatures.
- 1 Giant's Boulder | synergy | Board utility and supplemental removal.
- 1 Boughside Wanderers | threat | Elf board presence.
- 1 Cantankerous Keepers | threat | Elf combat pressure.
- 1 Elven Raft-Steerer | threat | Elf board development.
- 1 Galadhrim Guide | threat | Elf support and combat presence.
- 1 Grey Havens Navigator | threat | Elf utility and board presence.
- 1 Guardian of the Halls | threat | Elf combat presence.
- 1 Haunt of the Dead Marshes | threat | Graveyard-oriented Elf threat.
- 1 Lothlórien Lookout | threat | Early Elf presence.
- 1 Mirkwood Nurturer | threat | Supports the developing Elf board.
- 1 Mirkwood Pathmaker | threat | Helps turn creature pressure into damage.
- 1 Nimrodel Watcher | threat | Elf combat pressure.
- 1 Mirkwood Meditator | threat | Elf utility and board presence.
- 1 Troll of Khazad-dûm | wincon | Large evasive finisher with early land-finding utility.
- 1 Witch-king of Angmar | wincon | Resilient evasive finisher.
- 1 Witch-king, Bringer of Ruin | wincon | Large evasive closing threat.
- 1 Languish | wipe | Sweeps smaller opposing creatures.
- 1 Raise the Palisade | wipe | Tribal board reset.
- 1 Mithril Coat | protection | Protects an important legendary creature.
- 1 Bilbo's Ring | protection | Equipment-based protection and combat support.
- 1 My Precious // Allure of Power | protection | Flexible protection and equipment utility.
- 1 Wizard's Staff | other | Equipment utility for the creature plan.

</details>

### 23. two set families at once

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 203 names.

Commander: Kíli the Resourceful.

Grade: typical, score 0.44, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $306.72 to buy, $306.72 the whole deck.

**Summary:** Kíli the Resourceful turns a steady procession of Dwarves and Equipment into cards, while artifacts and legendary permanents build toward an enduring story. Develop mana early, establish an attacker, and use the free equip opportunity to keep combat pressure moving. The deck wins through its creature and Equipment package, with Angel of the Ruins providing a top-end finisher. Draw engines and protective interaction help sustain that pressure, but the deck relies on its commander and battlefield rather than a compact combo, so repeated board clears can slow it considerably.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The mana base serves the colors evenly, and that raises the grade. The commander does not place in cEDH events, and that lowers the grade.

This deck plan aims for 2 or more finishers, the cards that win the game, and the deck holds 1.

Summary rules claims (F-26): the format play

- JUDGE [unknown]: "use the free equip opportunity to keep combat pressure moving". This describes what the deck does on the table, not a rule. It does not say what a card may do or what a format allows, so it is not a rules claim.
- PLAN plan_coherent=yes: Mono-white Kíli with Dwarves, many Equipment, artifact ramp and draw all feed one equipment-and-creature plan, though the three wraths sit slightly against the creature plan.
- PLAN theme_fit=yes: Every card is from the Hobbit or Bloomburrow sets, the commander is a Hobbit Dwarf chosen by the builder, and the power level fits bracket 3 with Sol Ring, Mox Amber and a few tutors-free staples.
- PLAN useful_as_built=partly: The mono-white mana base of 31 Plains plus utility lands and plenty of ramp and draw works smoothly, but the deck has few real finishers, relying on Angel of the Ruins and Helm of the Host, and many cheap Equipment pieces with a thin creature count may leave it short on ways to win.
- PLAN summary_honest=partly: The summary claims Angel of the Ruins is a finisher and that the deck wins through creatures and Equipment, which the list does carry, but it says the commander turns Dwarves and Equipment into cards while the deck is mostly Equipment and artifacts with only a handful of Dwarves.
- [INFO] `curve_summary`: average mana value 2.74 over 62 nonland cards
- [WARN] `finisher_short`: the finisher count is 1, and this deck plan wants 2 or more (Angel of the Ruins)

<details><summary>The deck list</summary>

- 31 Plains | land | Provides a dependable white mana base for early creatures and interaction.
- 1 Command Tower | land | Adds another reliable white source.
- 1 Castle Ardenvale | land | Provides white mana and a late-game source of creatures.
- 1 Minas Tirith | land | Adds white mana and another way to maintain cards in hand.
- 1 Lupinflower Village | land | Supports the white mana base without using another colorless land slot.
- 1 Rogue's Passage | land | Helps an equipped attacker push through a crowded battlefield.
- 1 Treasure Vault | land | Adds an artifact permanent to help establish Kíli's enduring story.
- 1 Sol Ring | ramp | Accelerates the deployment of creatures and Equipment.
- 1 Mox Amber | ramp | Pairs with the inexpensive legendary commander to accelerate development.
- 1 Arcane Signet | ramp | Provides inexpensive white-producing acceleration.
- 1 Mind Stone | ramp | Accelerates early plays and remains useful when additional mana is unnecessary.
- 1 Thought Vessel | ramp | Adds inexpensive acceleration while supporting a growing hand.
- 1 Ornithopter of Paradise | ramp | Provides colored acceleration on an artifact creature.
- 1 Fellwar Stone | ramp | Adds another inexpensive mana rock.
- 1 Wayfarer's Bauble | ramp | Turns an early artifact into lasting land development.
- 1 Loyal Warhound | ramp | Helps recover mana development while adding a creature to the battlefield.
- 1 Patchwork Banner | ramp | Combines mana development with support for the creature plan.
- 1 Orcrist, Goblin-cleaver | ramp | Fills a ramp slot while adding Equipment for Kíli.
- 1 Burnished Hart | ramp | Builds the basic-land mana base from an artifact creature.
- 1 Skullclamp | draw | Provides card flow in an Equipment package that already rewards deployment.
- 1 Caretaker's Talent | draw | Adds a continuing draw engine alongside the deck's token support.
- 1 Palantír of Orthanc | draw | Provides card advantage from a legendary artifact.
- 1 Dawn of a New Age | draw | Rewards developing a substantial creature presence.
- 1 Inspiring Overseer | draw | Replaces itself while adding an evasive Equipment carrier.
- 1 Cut a Deal | draw | Refills the hand without requiring an established battlefield.
- 1 Circuit Mender | draw | Adds card flow and an artifact creature for the battlefield plan.
- 1 The Arkenstone // Seek the Heart | draw | Adds thematic card advantage and a legendary artifact.
- 1 Errand-Rider of Gondor | draw | Adds a creature-based source of card flow.
- 1 Belladonna Took | draw | Provides another draw option on a legendary creature.
- 1 Galadriel's Dismissal | interaction | Provides flexible interaction for protecting or clearing a combat step.
- 1 Reprieve | interaction | Buys time against an important opposing spell.
- 1 Dawn's Truce | interaction | Helps preserve the battlefield against opposing interaction.
- 1 Crumb and Get It | interaction | Offers inexpensive protection for an important creature.
- 1 Parting Gust | interaction | Adds another flexible instant-speed interaction slot.
- 1 Mithril Coat | interaction | Protects a key legendary creature while supporting the Equipment plan.
- 1 Swiftfoot Boots | interaction | Helps keep the commander or an important attacker available.
- 1 Swords to Plowshares | removal | Answers a dangerous creature for a small mana investment.
- 1 Generous Gift | removal | Provides a broad answer to a troublesome permanent.
- 1 Skyclave Apparition | removal | Combines targeted removal with another Equipment carrier.
- 1 Loran of the Third Path | removal | Adds permanent removal on a legendary creature.
- 1 Repel Calamity | removal | Provides an inexpensive answer to a substantial threat.
- 1 Nettle Guard | removal | Adds removal utility to an inexpensive creature.
- 1 Ori, Plate Stacker | removal | Fills a removal slot while retaining a legendary Dwarf for Kíli.
- 1 Fog on the Barrow-Downs | removal | Adds another creature-control option to the removal package.
- 1 Dáin, Lord of the Iron Hills | synergy | Supports the Dwarf theme while adding another legendary permanent.
- 1 Iron Hills Blacksmith | synergy | Adds a Dwarf synergy piece that also works with Kíli's draw trigger.
- 1 Ori, Keeper of Songs | synergy | Supports the Dwarf plan and contributes a legendary permanent.
- 1 Blade Splicer | synergy | Develops both creature and artifact presence for the battlefield plan.
- 1 Maskwood Nexus | synergy | Connects the broader creature package to Kíli's Dwarf trigger.
- 1 Dwarven Provisioner | synergy | Adds another Dwarf to help sustain the commander's card flow.
- 1 Carrot Cake | synergy | Adds artifact and token support for the deck's overlapping engines.
- 1 Fíli the Pathfinder | threat | Provides a thematic legendary Dwarf threat.
- 1 Andúril, Flame of the West | threat | Adds a legendary Equipment threat to the combat plan.
- 1 Andúril, Narsil Reforged | threat | Provides another legendary Equipment payoff for developed attackers.
- 1 Dwarven Shortsword | threat | Adds Equipment that supports both attacking and Kíli's card flow.
- 1 Dúnedain Blade | threat | Adds another Equipment option for building a threatening attacker.
- 1 Long-Lost Lances | threat | Broadens the Equipment package used to maintain combat pressure.
- 1 Sting, Bilbo's Sword | threat | Adds a legendary Equipment option to the attacking package.
- 1 Sword of the Squeak | threat | Provides a combat payoff for the deck's smaller creatures.
- 1 Shrike Force | threat | Adds an evasive creature suited to carrying Equipment.
- 1 Warren Warleader | threat | Provides a substantial creature threat for the combat plan.
- 1 Jacked Rabbit | threat | Adds a creature threat that also contributes to card flow.
- 1 Steelburr Champion | threat | Adds another creature to convert the Equipment package into pressure.
- 1 Angel of the Ruins | wincon | Provides the marked top-end finisher.
- 1 Helm of the Host | wincon | Turns a strong creature into a growing combat-based closing plan.
- 1 Jazal Goldmane | wincon | Provides a closing payoff for a developed creature battlefield.
- 1 Dusk // Dawn | wipe | Provides a battlefield reset with a recovery option for smaller creatures.
- 1 Martial Coup | wipe | Resets a crowded battlefield while supporting the creature plan.
- 1 Starfall Invocation | wipe | Provides another reset when opposing creatures become overwhelming.

</details>

### 24. a 60-card deck from one set

Format: Modern. Theme: aggro. Pool: any_card. Shortlist: 96 names.

Grade: baseline, score 0.36, model 20260923T202806Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $23.38 to buy, $23.38 the whole deck.

**Summary:** This mono-red Bloomburrow aggro deck applies early creature pressure, with a Mouse-focused core and supporting attackers. Combat support and removal keep attacks productive, while draw spells help sustain momentum and larger threats provide a finish when games run longer. It favors consistent red mana and proactive play over defensive flexibility, leaving it vulnerable to repeated board clears and opponents that establish stronger late-game engines.

The quality model grades this deck at the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every card is a cheap red creature or a combat trick or burn spell that supports attacking, so the whole list pulls toward one Mouse-flavored aggro plan.
- PLAN theme_fit=yes: The deck is mono-red, built entirely from Bloomburrow-set cards, and plays as a casual aggro deck in Modern as the request asked.
- PLAN useful_as_built=yes: With 24 lands, a low curve topping out at two Dragonhawks, and enough creatures and burn to close games, it can be played as it stands, though 22 Mountains plus two utility lands is slightly heavy on land for such a low curve.
- PLAN summary_honest=yes: The summary claims a Mouse-focused core with early pressure, and the list does carry eight-plus Mouse creatures such as Heartfire Hero, Manifold Mouse, Emberheart Challenger and Roughshod Duo, plus Take Out the Trash and Rabid Gnaw as removal and Dragonhawk as a top-end finisher.
- [INFO] `curve_summary`: average mana value 1.94 over 36 nonland cards

<details><summary>The deck list</summary>

- 22 Mountain | land | Provides dependable red mana for an aggressive opening and later threats.
- 2 Rockface Village | land | Supports the creature-heavy mana base and gives relevant attackers a late-game boost.
- 4 Heartfire Hero | threat | An early Mouse threat that rewards targeting it with your combat spells.
- 4 Hired Claw | threat | Adds another one-mana attacker to keep opening hands aggressive.
- 4 Manifold Mouse | threat | Makes your Mouse attackers substantially more dangerous in combat.
- 2 Dragonhawk, Fate's Tempest | threat | Provides a larger flying threat when early aggression does not finish the game.
- 4 Emberheart Challenger | synergy | Connects the Mouse attack plan with spells that target your own creatures.
- 4 Roughshod Duo | synergy | Turns subsequent creature plays into targeted boosts for your attackers.
- 4 Might of the Meek | draw | Replaces itself while targeting an attacker and helping damage get through.
- 2 Sazacap's Brew | draw | Refreshes your hand, with its gift option adding another targeted combat boost.
- 2 Brazen Collector | ramp | Supplies attack-generated mana while contributing to the creature pressure.
- 4 Take Out the Trash | removal | Removes opposing creatures that would otherwise interrupt your attacks.
- 2 Rabid Gnaw | removal | Uses an attacker to remove a blocker while also targeting your own creature.

</details>

### 25. use no card of an owned precon

Format: Commander. Theme: Avengers superheroes. Pool: owned_first. Shortlist: 210 names.

Commander: Captain America, Super-Soldier.

Grade: baseline, score 0.36, model 20260923T202806Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $2.79 to buy, $130.69 the whole deck.

**Summary:** Captain America leads a white Avengers team that develops its mana, keeps cards flowing, and pressures opponents with Heroes. Protect Captain America, maintain the team's presence, and close with finishers, chief among them The Sentry, Golden Guardian. The deck favors sustained creature combat over explosive combos and trades some superhero flavor for practical support. The Avengers Assemble precon inventory was not supplied, so exclusion of its cards cannot be confirmed.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The mana base serves the colors evenly, and that raises the grade. The commander does not place in cEDH events, and that lowers the grade.

Summary rules claims (F-26): the format play

- JUDGE [unknown]: "The Avengers Assemble precon inventory was not supplied, so exclusion of its cards cannot be confirmed.". This describes a limit of the data given to the builder, not a rule of the game. It is not a rules claim, so it is listed only for completeness.
- JUDGE [unknown]: "The commander does not place in cEDH events, and that lowers the grade.". This is a statement about the commander's competitive results and the grading model, not a rule about what a card may do or what a format allows. It is not a rules claim.
- PLAN plan_coherent=partly: The Hero creatures, protection pieces, and removal all support a white creature plan, but the pile of ten mana rocks, three wraths, and a handful of reactive instants pull against a Hero-pressure plan.
- PLAN theme_fit=partly: It is a Captain America mono-white Hero deck with many Marvel Avengers cards, but it includes Avengers Assemble!, which may be from the precon the player excluded, and it leans on generic staples like Sol Ring and Skullclamp instead of more Avengers cards.
- PLAN useful_as_built=partly: The deck has 35 lands with ten mana rocks, plenty of cheap Heroes, and ample removal, so it plays smoothly, but the very high mana count and three wraths leave few threats and a thin route to winning.
- PLAN summary_honest=yes: The summary says the deck is white, led by Captain America, closes with The Sentry, Golden Guardian, and may include cards from the precon; each of these matches the list, which has Captain America, a mono-white mana base, The Sentry as a wincon, and Avengers Assemble!, so the caveat is accurate.
- [INFO] `curve_summary`: average mana value 2.87 over 63 nonland cards

<details><summary>The deck list</summary>

- 31 Plains | land | The main white mana source for Captain America and the team.
- 1 Avengers Tower | land | A thematic land slot for the Avengers mana base.
- 1 Minas Tirith | land | A utility land alongside the basic white mana base.
- 1 War Room | land | Adds a utility land without crowding out white sources.
- 1 Secluded Courtyard | land | Supports the creature-focused mana base.
- 1 Nesting Grounds | land | Provides another utility land slot.
- 1 Agent 13, Sharon Carter | draw | Hero-themed card advantage for sustained pressure.
- 1 Agent Maria Hill | draw | Keeps the deck supplied with cards while reinforcing its theme.
- 1 Avengers Assemble! | draw | A central Avengers-themed card-advantage piece.
- 1 Origin of the Avengers | draw | Adds card advantage through an Avengers story card.
- 1 The Vision | draw | Combines a Hero presence with the draw package.
- 1 Viv Vision, Teen Synthezoid | draw | Another Hero-themed source of card advantage.
- 1 Hero in Training | draw | Supports the Hero theme while filling a draw slot.
- 1 Mask of Memory | draw | Card-advantage support for the combat plan.
- 1 Skullclamp | draw | An efficient card-advantage tool for a creature-heavy deck.
- 1 Inspiring Overseer | draw | Adds a creature to the card-advantage package.
- 1 Vanquisher's Banner | draw | Card-advantage support for the shared creature theme.
- 1 Sol Ring | ramp | Early acceleration toward the commander and supporting threats.
- 1 Arcane Signet | ramp | Mana acceleration for deploying the team.
- 1 Thought Vessel | ramp | Additional artifact acceleration.
- 1 Springleaf Drum | ramp | A low-cost ramp piece for the creature-focused plan.
- 1 Wayfarer's Bauble | ramp | Early mana development to support later turns.
- 1 Sword of the Animist | ramp | Ramp support that fits the combat-oriented build.
- 1 Bender's Waterskin | ramp | Another mana-development piece from the library.
- 1 White Lotus Tile | ramp | Supports consistent mana development.
- 1 Wickersmith's Tools | ramp | Adds acceleration without increasing the threat curve.
- 1 Page, Loose Leaf | ramp | A creature-based addition to the ramp package.
- 1 Boromir, Warden of the Tower | interaction | A supporting creature in the interaction package.
- 1 Frontline Medic | interaction | Interaction support suited to a creature-heavy combat plan.
- 1 Clever Concealment | interaction | Helps the team navigate opposing interaction.
- 1 Reprieve | interaction | A flexible interaction slot for pivotal turns.
- 1 Take Up the Shield | interaction | Combat-oriented interaction that suits Captain America's theme.
- 1 Unbreakable Formation | interaction | Supports the team during contested combat turns.
- 1 Slip On the Ring | interaction | An inexpensive addition to the interaction package.
- 1 Duty Beyond Death | interaction | Provides another response option for protecting the game plan.
- 1 Swords to Plowshares | removal | Efficient removal for an opposing threat.
- 1 Generous Gift | removal | Broad removal support against troublesome permanents.
- 1 Get Lost | removal | A low-cost answer in the removal package.
- 1 Stroke of Midnight | removal | Another flexible removal option.
- 1 Dispatch | removal | An inexpensive removal slot alongside the artifact package.
- 1 Crib Swap | removal | Additional removal to clear the way for the team.
- 1 March of Otherworldly Light | removal | Adds flexibility to the removal suite.
- 1 Journey to Nowhere | removal | A low-curve removal piece.
- 1 Palace Jailer | removal | A creature-based removal option.
- 1 Agent Phil Coulson | synergy | A thematic support Hero for the Avengers team.
- 1 Agent of Atlas | synergy | Adds another Hero to the shared creature theme.
- 1 Agents of S.H.I.E.L.D. | synergy | Reinforces the Avengers support-team theme.
- 1 Colleen Wing, Street Samurai | synergy | A Hero support piece for the creature plan.
- 1 Night Nurse, Healer of Heroes | synergy | Keeps the support package closely tied to the Hero theme.
- 1 Quake, Agent of S.H.I.E.L.D. | synergy | Another thematic support Hero for the team.
- 1 Silver Sable, Mercenary Leader | synergy | Rounds out the Hero synergy package.
- 1 Captain Mar-Vell, Space-Born | threat | An Avengers-themed Hero threat for the combat plan.
- 1 Invisible Woman, Sue Storm | threat | Adds another superhero threat to the battlefield.
- 1 Luke Cage, Power Man | threat | A Hero threat that keeps the deck focused on its theme.
- 1 Mockingbird, Ace Agent | threat | Provides another Avengers-associated combat threat.
- 1 Okoye, Dora Milaje Leader | threat | Adds a Hero threat to the team.
- 1 Roaming Throne | threat | A supporting threat for the creature-focused strategy.
- 1 The Sentry, Golden Guardian | wincon | The primary superhero-themed finisher.
- 1 Origin of Spider-Man | wincon | A superhero-themed finisher that diversifies the closing threats.
- 1 Angel of Serenity | wincon | An additional finisher for longer games.
- 1 Austere Command | wipe | A board-reset option when opponents get ahead.
- 1 Dusk // Dawn | wipe | A second board-reset option for difficult creature boards.
- 1 Split Up | wipe | Rounds out the board-wipe package.
- 1 Patriot, Shield Wielder | protection | Hero-themed protection for the team.
- 1 Bastion Protector | protection | Protection support focused on keeping the commander present.
- 1 Lightning Greaves | protection | A reusable protection piece for important creatures.
- 1 Swiftfoot Boots | protection | Additional reusable protection for Captain America and key Heroes.
- 1 Sheltered by Ghosts | protection | A low-curve protection slot for an important creature.
- 1 Together Forever | protection | Protection support for the creature-heavy game plan.

</details>

