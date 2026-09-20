# PR-8 deck gate

Run date: 2026-09-20. Card snapshot: 2026-09-04.

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
| Mean plan score, 0 to 1 | 0.70 |
| Plan reasons the judge left empty | 3 |
| Errors | 0 |
| Prompt version | 15 |
| Calls | 75 |
| Cost | $2.7384 |
| Time | 1874 seconds |

## Run

- Suite `decks`, run `pr8-deck-gate-run29`, on 2026-09-20, commit `5fd8085`.
- Roles: generate on `gpt-5.6-terra` (openai, effort medium), judge on `claude-opus-5` (anthropic, effort medium), repair on `gpt-5.6-terra` (openai, effort medium).
- Versions: card snapshot 2026-09-04, generate prompt version 15, plan_rubric prompt version 2, precons `5.3.0+20260903`, quality_model `20260914T154223Z`.
- Calls: 75. Cost: $2.7384. Time: 1874 seconds.

## Findings by code

A block stops the deck. A warning and a note are reports.

| Code | Count |
|---|---|
| `curve_summary` | 25 |
| `not_owned` | 20 |
| `profile_off_band` | 11 |
| `mana_pass` | 6 |
| `bracket_cut` | 2 |
| `finisher_short` | 2 |
| `outside_requested_set` | 2 |
| `precon_cards_restored` | 1 |
| `cards_trimmed` | 1 |
| `basics_added` | 1 |

By severity: BLOCK 0. WARN 35. INFO 36. 

## The set filter (PR-17B)

A deck asked for a set holds cards of that set family alone. Three things earn an exception, and each one is marked: a card the reader named (D-381), a mana card the reader allowed from outside (D-382), and a basic land, which no set limit filters (D-378).

| # | Prompt | Sets | In set | Fill | Deck cards outside | Marked |
|---|---|---|---|---|---|---|
| 19 | the Hobbit family, two colours | `hob,hoc` | 170 | 0 | 0 | 0 |
| 20 | the Hobbit family, a delegated commander | `hob,hoc` | 75 | 0 | 0 | 0 |
| 21 | the Hobbit family, mana from outside | `hob,hoc` | 127 | 26 | 13 | 13 |
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

Cost: $834.22 to buy, $834.22 the whole deck.

**Summary:** Karlov of the Ghost Council leads a white-black lifegain deck that builds a resilient creature board, turns repeated lifegain into growing threats and pressure, and backs that plan with substantial removal and board resets. Angel of Destiny, Felidar Sovereign, and Test of Endurance provide decisive endgames, while Sanguine Bond and Vito, Thorn of the Dusk Rose reinforce the central theme. The deck gives up some speed for a broad creature package and several higher-cost Angels and Vampires.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The curve sits high for the format, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every slot feeds one loop — incidental lifegain triggers (Pristine Talisman, The Celestus, lifelinkers, Angels) that grow Karlov and turn on payoffs like Ajani's Pridemate, Bloodthirsty Aerialist, Sanguine Bond/Vito, Felidar Sovereign and Test of Endurance — with only mild tension from three sweepers in a creature-heavy list, two of which gain life themselves.
- PLAN theme_fit=yes: It is exactly the requested white-black Karlov lifegain Commander deck, with a card pool (no obvious game changers, mostly synergy commons and midrange staples) that sits naturally at bracket 3.
- PLAN useful_as_built=yes: 37 lands heavily weighted to duals and painlands plus roughly ten mana rocks support the admittedly high curve, and the deck has multiple clear routes to victory (Karlov beats, Sanguine Bond/Vito drain, Aetherflux Reservoir, and the two alt-win permanents).
- PLAN summary_honest=yes: Every element named — the lifegain engine, Angel of Destiny/Felidar Sovereign/Test of Endurance as endgames, Sanguine Bond and Vito, the removal and wipes, and the admission of a high curve and slow speed — is actually present in the list, and nothing significant is concealed.
- [INFO] `curve_summary`: average mana value 3.47 over 62 nonland cards
- [INFO] `bracket_cut`: to hold bracket 3, the builder cut 1 card: Enduring Angel // Angelic Enforcer (the combo Enduring Angel // Angelic Enforcer + Enduring Tenacity), and added 1 basic land
- [INFO] `mana_pass`: the builder moved 3 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 1 Godless Shrine | land | Dual-color land for the white-black mana base.
- 1 Scrubland | land | Dual-color land for the white-black mana base.
- 1 Caves of Koilos | land | Provides either deck color early.
- 1 Command Tower | land | Reliable dual-color commander land.
- 1 Brightclimb Pathway // Grimclimb Pathway | land | Flexible choice of either deck color.
- 1 Concealed Courtyard | land | Early dual-color land.
- 1 Vault of Champions | land | Dual-color land for multiplayer games.
- 1 Shattered Sanctum | land | Dual-color land for the mana base.
- 1 Silent Clearing | land | Dual-color land that supports the mana base.
- 1 Marsh Flats | land | Fetch land that finds either basic color source.
- 1 Prismatic Vista | land | Fetch land that fixes either basic color.
- 1 Mana Confluence | land | Flexible source for both deck colors.
- 1 City of Brass | land | Flexible source for both deck colors.
- 11 Plains | land | Primary white basic sources.
- 13 Swamp | land | Primary black basic sources.
- 1 Legion's Landing // Adanto, the First Fort | ramp | Low-cost ramp slot that also fits the deck's creature plan.
- 1 Hot Dog Cart | ramp | Dedicated mana development for the deck.
- 1 Phial of Galadriel | ramp | Mana development in a compact artifact slot.
- 1 Nuka-Cola Vending Machine | ramp | Artifact-based mana development.
- 1 Pristine Talisman | ramp | Ramp piece aligned with the lifegain plan.
- 1 Orazca Relic | ramp | Artifact mana development.
- 1 The Celestus | ramp | Artifact ramp for reaching the deck's higher-cost cards.
- 1 Potioner's Trove | ramp | Mana development from an artifact slot.
- 1 Altar of the Pantheon | ramp | Color-fixing artifact ramp.
- 1 Carmen, Cruel Skymarcher | ramp | Creature-based ramp that retains the deck's theme.
- 1 Archivist of Oghma | draw | Efficient creature-based card advantage.
- 1 Dawn of Hope | draw | Lifegain-aligned card advantage engine.
- 1 Cosmos Elixir | draw | Ongoing card advantage in an artifact slot.
- 1 Convalescent Care | draw | Lifegain-focused source of cards.
- 1 Enduring Innocence | draw | Creature-based card advantage for the deck.
- 1 Exemplar of Light | draw | Lifegain-themed creature that supplies cards.
- 1 Mangara, the Diplomat | draw | Reliable creature-based card advantage.
- 1 Markov Purifier | draw | Thematic Vampire card advantage.
- 1 Sigarda's Splendor | draw | Lifegain-aligned enchantment card advantage.
- 1 Well of Lost Dreams | draw | A dedicated lifegain card-draw outlet.
- 1 The Gaffer | draw | Creature-based card advantage that suits the plan.
- 1 Alseid of Life's Bounty | interaction | Low-cost protective interaction.
- 1 Caduceus, Staff of Hermes | interaction | Flexible artifact interaction.
- 1 Courageous Resolve | interaction | Protective instant-speed interaction.
- 1 Faith's Shield | interaction | Low-cost protective interaction.
- 1 Metropolis Reformer | interaction | Creature-based protection for the deck.
- 1 Paladin Danse, Steel Maverick | interaction | Thematic creature interaction.
- 1 Sword of Light and Shadow | interaction | Equipment-based interaction that fits the creature plan.
- 1 Aetherflux Reservoir | removal | Lifegain-based removal option.
- 1 Ayli, Eternal Pilgrim | removal | Thematic creature-based removal.
- 1 Consuming Corruption | removal | Black removal for opposing problems.
- 1 Cavalier of Night | removal | Creature-based removal with a meaningful body.
- 1 Murderous Rider // Swift End | removal | Flexible removal attached to a creature.
- 1 Nightmare's Thirst | removal | Efficient black removal.
- 1 Solitude | removal | Flexible white creature removal.
- 1 Umezawa's Jitte | removal | Reusable equipment-based removal.
- 1 Vona, Butcher of Magan | removal | Lifegain-themed creature removal.
- 1 Fumigate | wipe | Board reset that suits the lifegain plan.
- 1 Kaya's Wrath | wipe | Reliable white-black board reset.
- 1 The Battle of Bywater | wipe | Additional board reset for creature-heavy tables.
- 1 Ajani's Pridemate | synergy | Classic lifegain payoff for Karlov's plan.
- 1 Bloodthirsty Aerialist | synergy | Lifegain-focused evasive payoff.
- 1 Heliod, Sun-Crowned | synergy | Central lifegain synergy piece.
- 1 Resplendent Angel | synergy | Thematic lifegain payoff.
- 1 Sanguine Bond | synergy | Converts the deck's lifegain theme into pressure.
- 1 Vito, Thorn of the Dusk Rose | synergy | Lifegain synergy that pressures opponents.
- 1 Voice of the Blessed | synergy | Efficient lifegain payoff creature.
- 1 Archangel of Thune | threat | Premium lifegain threat for the creature board.
- 1 Attended Healer | threat | Creature threat that fits the lifegain shell.
- 1 Blood Baron of Vizkopa | threat | Thematic Vampire threat.
- 1 Celestine, the Living Saint | threat | Legendary lifegain-themed threat.
- 1 Cliffhaven Vampire | threat | Lifegain-based Vampire pressure.
- 1 Divinity of Pride | threat | A powerful white-black lifegain threat.
- 1 Enduring Tenacity | threat | Resilient black threat for the deck.
- 1 Exalted Sunborn | threat | Large Angel threat for the top end.
- 1 Liesa, Forgotten Archangel | threat | Legendary Angel threat with strong board presence.
- 1 Lyra Dawnbringer | threat | High-impact Angel threat.
- 1 Valkyrie Harbinger | threat | Large lifegain-themed Angel threat.
- 1 Angel of Destiny | wincon | Thematic Angel finisher.
- 1 Felidar Sovereign | wincon | Lifegain-based alternate victory threat.
- 1 Test of Endurance | wincon | Dedicated lifegain victory condition.
- 1 Arguel's Blood Fast // Temple of Aclazotz | ramp | the mana pass added it to bring the mana base inside the power level

</details>

### 2. aristocrats Commander, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 181 names.

Commander: Denethor, Ruling Steward.

Grade: typical, score 0.39, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $226.44 the whole deck.

**Summary:** Denethor anchors a creature-heavy aristocrats shell that develops mana and card flow early, then turns its synergy creatures into steady pressure while holding up broad interaction and protection. The deck wins through its finisher-tagged cards and an attrition-based board presence, with sweepers available to reset opponents who get ahead. It gives up some raw threat density for resilience, creature synergy, and a substantial package of answers.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck draws more than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=no: The list pulls three ways at once — a commander-protection/equipment voltron package (Bastion Protector, Champion's Helm, Darksteel Plate, Greaves, Boots, Gift of Immortality), a creature board of small synergy bodies, and three unconditional wraths that would blow up that board — with no sacrifice outlets or death-trigger payoffs to tie any of it together.
- PLAN theme_fit=no: The person asked for aristocrats, and the deck contains almost none of the required pieces — no Blood Artist/Zulaport-style drains, no free sacrifice outlets beyond Deadly Dispute, and no recursion loop — so it plays as a generic W/B goodstuff protection deck instead.
- PLAN useful_as_built=partly: The mana base and curve are serviceable and the removal/ramp suite functions, but the deck has no real win condition beyond incidental beatdown from small utility creatures, and its own sweepers work against the board it builds.
- PLAN summary_honest=no: It calls the deck a 'creature-heavy aristocrats shell' when there are roughly a dozen creatures and essentially no sac outlets or drain/death payoffs, and the 'finisher-tagged' cards it points to (Al Bhed Salvagers, Grave Venerations) are not actually win conditions.
- [INFO] `curve_summary`: average mana value 2.73 over 63 nonland cards

<details><summary>The deck list</summary>

- 16 Plains | land | Reliable white source for the mana base.
- 13 Swamp | land | Reliable black source for the mana base.
- 1 Command Tower | land | Flexible nonbasic mana-base slot.
- 1 City of Brass | land | Flexible nonbasic mana-base slot.
- 1 Exotic Orchard | land | Flexible nonbasic mana-base slot.
- 1 Marsh Flats | land | Mana-base fixing slot.
- 1 Fabled Passage | land | Mana-base fixing slot.
- 1 Grand Coliseum | land | Flexible nonbasic mana-base slot.
- 1 Path of Ancestry | land | Mana-base slot that fits the creature-focused shell.
- 1 Ahriman | draw | Provides a draw slot in the creature-heavy plan.
- 1 Beetle-Headed Merchants | draw | Provides a draw slot for sustained resources.
- 1 Buster Sword | draw | Provides a draw slot while supporting the artifact package.
- 1 Call of the Ring | draw | Provides a draw slot for sustained resources.
- 1 Circle of Power | draw | Provides a draw slot for sustained resources.
- 1 Idol of Oblivion | draw | Provides a draw slot in the artifact package.
- 1 Lembas | draw | Provides a draw slot in the artifact package.
- 1 Mask of Memory | draw | Provides a draw slot while rewarding creature combat.
- 1 Night's Whisper | draw | Provides an efficient draw slot.
- 1 Painful Truths | draw | Provides a draw slot for sustained resources.
- 1 Skullclamp | draw | Provides a draw slot suited to disposable creatures.
- 1 Stone of Erech | draw | Provides a draw slot in the artifact package.
- 1 Tome of Legends | draw | Provides a draw slot alongside the commander.
- 1 Wall of Omens | draw | Provides a draw slot on a cheap creature.
- 1 Bastion Protector | interaction | Provides an interaction slot that supports Denethor.
- 1 Boromir, Warden of the Tower | interaction | Provides a creature-based interaction slot.
- 1 Champion's Helm | interaction | Provides a protective interaction slot.
- 1 Clever Concealment | interaction | Provides a flexible interaction slot.
- 1 Darksteel Plate | interaction | Provides a protective interaction slot.
- 1 Duty Beyond Death | interaction | Provides a protection-focused interaction slot.
- 1 Gift of Immortality | interaction | Provides a protection-focused interaction slot.
- 1 Lightning Greaves | interaction | Provides a protective interaction slot.
- 1 Reprieve | interaction | Provides a flexible interaction slot.
- 1 Swiftfoot Boots | interaction | Provides a protective interaction slot.
- 1 Take Up the Shield | interaction | Provides a protection-focused interaction slot.
- 1 Together Forever | interaction | Provides a resilient interaction slot for creatures.
- 1 Arcane Signet | ramp | Provides efficient color fixing and acceleration.
- 1 Bender's Waterskin | ramp | Provides a ramp slot for the mana plan.
- 1 Chromatic Lantern | ramp | Provides fixing alongside acceleration.
- 1 Commander's Sphere | ramp | Provides a ramp slot with flexible mana.
- 1 Deadly Dispute | ramp | Provides a ramp slot aligned with the sacrifice plan.
- 1 Fellwar Stone | ramp | Provides efficient acceleration.
- 1 Lotho, Corrupt Shirriff | ramp | Provides a creature-based ramp slot.
- 1 Relic of Legends | ramp | Provides a ramp slot for the legendary package.
- 1 Sol Ring | ramp | Provides efficient acceleration.
- 1 Springleaf Drum | ramp | Provides low-cost acceleration with creatures.
- 1 Wayfarer's Bauble | ramp | Provides a basic-land ramp slot.
- 1 Bitter Triumph | removal | Provides a focused removal slot.
- 1 Claim the Precious | removal | Provides a focused removal slot.
- 1 Crib Swap | removal | Provides a flexible removal slot.
- 1 Fatal Push | removal | Provides a low-cost removal slot.
- 1 Generous Gift | removal | Provides broad removal coverage.
- 1 Infernal Grasp | removal | Provides a focused removal slot.
- 1 Orcish Bowmasters | removal | Provides creature-based removal.
- 1 Swords to Plowshares | removal | Provides efficient removal coverage.
- 1 Stroke of Midnight | removal | Provides flexible removal coverage.
- 1 Arcade Cabinet | synergy | Provides a dedicated synergy piece.
- 1 Gollum the Abandoned | synergy | Provides a creature synergy piece for the aristocrats shell.
- 1 Gollum, Patient Plotter | synergy | Provides a creature synergy piece for the aristocrats shell.
- 1 Gríma Wormtongue | synergy | Provides a dedicated synergy piece.
- 1 Heirloom Auntie | synergy | Provides a creature synergy piece.
- 1 Joo Dee, One of Many | synergy | Provides a dedicated synergy piece.
- 1 Nimble Hobbit | synergy | Provides a low-cost creature synergy piece.
- 1 Bill the Pony | threat | Provides a creature threat for board presence.
- 1 Hei Bai, Spirit of Balance | threat | Provides a creature threat for board presence.
- 1 Namazu Trader | threat | Provides a creature threat for board presence.
- 1 Vengeful Villagers | threat | Provides a creature threat for board presence.
- 1 Al Bhed Salvagers | wincon | Provides a finisher-tagged win condition.
- 1 Archfiend of Ifnir | wincon | Provides a finisher-tagged win condition.
- 1 Grave Venerations | wincon | Provides a finisher-tagged win condition for the graveyard-focused plan.
- 1 Austere Command | wipe | Provides a flexible board-wipe slot.
- 1 Fumigate | wipe | Provides a board-wipe slot.
- 1 Vanquish the Horde | wipe | Provides a board-wipe slot.

</details>

### 3. artifacts Commander, bracket 4

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 319 names.

Commander: Urza, Lord High Artificer.

Grade: good, score 0.66, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $3871.00 to buy, $3871.00 the whole deck.

**Summary:** This high-power blue artifact deck uses Urza, Lord High Artificer to support a dense artifact core with explosive acceleration, sustained card draw, and strong protective interaction. It can pressure the table with artifact threats while presenting dedicated wins through Mechanized Production, Mirrodin Besieged, and Thassa's Oracle. It gives up broad multicolor access for a focused blue artifact shell and a mana base centered on Islands.

The quality model grades this deck good against the top lists of the format: the deck makes more mana on turn four than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

- PLAN plan_coherent=yes: Nearly every card supports one artifact-density plan around Urza — cheap rocks and affinity creatures to flood the board, cost reducers, artifact-fueled draw, and blue interaction to protect it — with only Thassa's Oracle sitting outside the line as a card with no enabler.
- PLAN theme_fit=yes: It is exactly the requested high-power Urza, Lord High Artificer artifact deck, with a bracket-4-appropriate package (Mana Vault, LED, Mox suite, Force of Will, Urza's Saga) rather than a cEDH tutor-and-fast-combo pile.
- PLAN useful_as_built=yes: 34 lands plus a dozen cheap rocks supports a low curve, the affinity/cost-reduction package lands spells reliably, and Mechanized Production, Mirrodin Besieged, and a wide artifact board plus Cyberdrive Awakener give real ways to close a game.
- PLAN summary_honest=partly: The acceleration, draw, and interaction claims are accurate, but calling Thassa's Oracle a 'dedicated win' is misleading since the list contains no Demonic Consultation-style effect or self-mill to empty the library, leaving it a near-blank card.
- [INFO] `curve_summary`: average mana value 2.98 over 66 nonland cards
- [INFO] `mana_pass`: the builder moved 1 card of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 26 Island | land | Forms the reliable blue mana base.
- 1 Command Tower | land | Adds a blue source to the mana base.
- 1 Glimmervoid | land | Adds a blue source for the artifact shell.
- 1 Spire of Industry | land | Adds a blue source alongside artifacts.
- 1 Seat of the Synod | land | Provides blue mana while fitting the artifact theme.
- 1 Otawara, Soaring City | land | Provides a blue source in a land slot.
- 1 Mystic Sanctuary | land | Provides an additional blue source.
- 1 Urza's Saga | land | Supplies a utility land slot for the artifact plan.
- 1 Chrome Mox | ramp | Provides premium fast mana.
- 1 Lion's Eye Diamond | ramp | Provides premium fast mana.
- 1 Lotus Petal | ramp | Provides immediate fast mana.
- 1 Mana Vault | ramp | Provides explosive fast mana.
- 1 Mox Diamond | ramp | Provides premium fast mana.
- 1 Mox Opal | ramp | Provides fast artifact-based mana.
- 1 Moonsnare Prototype | ramp | Provides low-cost artifact acceleration.
- 1 Sol Ring | ramp | Provides efficient early acceleration.
- 1 Springleaf Drum | ramp | Provides low-cost acceleration.
- 1 Metalworker | ramp | Provides artifact-focused mana production.
- 1 Krark-Clan Ironworks | ramp | Provides artifact-based acceleration.
- 1 Grand Architect | ramp | Provides additional ramp for the artifact plan.
- 1 Tezzeret the Seeker | ramp | Provides ramp while serving as a tutor.
- 1 Rhystic Study | draw | Provides a powerful draw engine.
- 1 Thoughtcast | draw | Provides efficient card draw for artifacts.
- 1 Sai, Master Thopterist | draw | Supports the deck's draw plan.
- 1 Thirst for Knowledge | draw | Provides efficient card draw.
- 1 Reverse Engineer | draw | Provides artifact-focused card draw.
- 1 Riddlesmith | draw | Adds repeatable draw support.
- 1 Forensic Gadgeteer | draw | Supports the deck's draw plan.
- 1 Cerebral Download | draw | Adds to the card-draw suite.
- 1 Era of Innovation | draw | Supports artifact-based card draw.
- 1 Esoteric Duplicator | draw | Adds artifact-based card draw.
- 1 Braided Net // Braided Quipu | draw | Adds another draw option.
- 1 An Offer You Can't Refuse | interaction | Provides efficient stack interaction.
- 1 Fierce Guardianship | interaction | Provides premium protective interaction.
- 1 Force of Will | interaction | Provides premium stack interaction.
- 1 Disruption Protocol | interaction | Adds artifact-friendly interaction.
- 1 Metallic Rebuke | interaction | Adds efficient artifact-focused interaction.
- 1 Ice Out | interaction | Provides stack interaction.
- 1 Reality Ripple | interaction | Adds flexible interaction.
- 1 Ghostly Flicker | interaction | Provides protective interaction.
- 1 Welding Jar | interaction | Provides artifact protection.
- 1 Escape Protocol | interaction | Adds interaction for the artifact shell.
- 1 Stoic Rebuttal | interaction | Provides reliable stack interaction.
- 1 Aether Spellbomb | removal | Provides cheap artifact-based removal.
- 1 Contagion Clasp | removal | Provides artifact-based removal.
- 1 Cyber Conversion | removal | Adds flexible removal.
- 1 Into Thin Air | removal | Adds removal coverage.
- 1 Resculpt | removal | Provides efficient removal.
- 1 Ravenform | removal | Adds removal coverage.
- 1 Lumengrid Drake | removal | Adds creature-based removal.
- 1 Kitesail Larcenist | removal | Adds creature-based removal.
- 1 Invasion Submersible | removal | Adds artifact-based removal.
- 1 Shape Anew | removal | Provides another removal option.
- 1 Engineered Explosives | wipe | Provides a scalable board wipe.
- 1 Rebuild | wipe | Provides a reset against artifact boards.
- 1 Foundry Inspector | synergy | Supports the artifact synergy package.
- 1 Etherium Sculptor | synergy | Supports the artifact synergy package.
- 1 Mystic Forge | synergy | Provides a key artifact synergy piece.
- 1 Kappa Cannoneer | threat | Provides a major artifact-based threat.
- 1 Cyberdrive Awakener | threat | Provides a high-impact artifact threat.
- 1 Darksteel Juggernaut | threat | Provides a durable artifact threat.
- 1 Gearseeker Serpent | threat | Provides a substantial artifact threat.
- 1 Karn, Scion of Urza | threat | Provides a planeswalker threat.
- 1 Kuldotha Forgemaster | threat | Provides a threat that also tutors.
- 1 Master Transmuter | threat | Provides an artifact-focused threat.
- 1 Phyrexian Metamorph | threat | Provides a flexible artifact threat.
- 1 Traxos, Scourge of Kroog | threat | Provides an efficient artifact threat.
- 1 Frogmite | threat | Adds to the artifact threat density.
- 1 Myr Enforcer | threat | Provides another artifact threat.
- 1 Sojourner's Companion | threat | Provides an artifact threat for the board.
- 1 Mechanized Production | wincon | Provides a dedicated artifact win condition.
- 1 Mirrodin Besieged | wincon | Provides a dedicated artifact win condition.
- 1 Thassa's Oracle | wincon | Provides a compact win condition.
- 1 Mox Amber | ramp | the mana pass added it to bring the mana base inside the power level

</details>

### 4. dinosaur tribal, bracket 2

Format: Commander. Theme: dinosaurs. Pool: any_card. Shortlist: 299 names.

Commander: Gishath, Sun's Avatar.

Grade: baseline, score 0.06, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $1210.69 to buy, $1210.69 the whole deck.

**Summary:** Gishath, Sun's Avatar leads a Dinosaur-focused deck that builds mana early, develops a board of tribal creatures, and keeps resources flowing with creature-oriented draw. It wins by turning its Dinosaur threats into sustained combat pressure, with Dinosaurs on a Spaceship, Huatli, Poet of Unity // Roar of the Fifth People, and Tribal Forcemage providing dedicated closing power. The deck gives up some speed for a straightforward creature plan and several larger Dinosaur cards.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The curve sits high for the format, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Every slot supports one Gishath dinosaur-tribal plan: tribal cost reducers and lords, ramp to reach the seven-drop commander, creature-based draw, and big Dinosaur payoffs that Gishath's trigger cheats in.
- PLAN theme_fit=yes: It is exactly the requested Dinosaur Commander deck under Gishath in Naya, with no game changers or combos and a straightforward creature plan appropriate to bracket 2, though the reserved-list duals are a steep ask for a new player if cost mattered.
- PLAN useful_as_built=yes: 38 lands with heavy fixing plus ten ramp pieces support the admittedly high curve, and the deck has plenty of ways to close a game through combat, so it plays fine out of the box.
- PLAN summary_honest=partly: The core claims (ramp, tribal board, creature draw, combat-pressure win) are all borne out by the list, but singling out Dinosaurs on a Spaceship, Huatli, and Tribal Forcemage as the 'dedicated closing power' misdescribes cards that are really removal/ramp/a small anthem, when the actual finish is Gishath plus Shared Animosity and the fatties.
- [INFO] `curve_summary`: average mana value 3.81 over 62 nonland cards

<details><summary>The deck list</summary>

- 7 Forest | land | Basic land for the mana base.
- 5 Plains | land | Basic land for the mana base.
- 5 Mountain | land | Basic land for the mana base.
- 1 Command Tower | land | Part of the three-color mana base.
- 1 Path of Ancestry | land | Part of the Dinosaur-focused mana base.
- 1 Cavern of Souls | land | Part of the creature-focused mana base.
- 1 Secluded Courtyard | land | Part of the creature-focused mana base.
- 1 Unclaimed Territory | land | Part of the creature-focused mana base.
- 1 City of Brass | land | Part of the three-color mana base.
- 1 Mana Confluence | land | Part of the three-color mana base.
- 1 Forbidden Orchard | land | Part of the three-color mana base.
- 1 Bountiful Promenade | land | Part of the green-white mana base.
- 1 Canopy Vista | land | Part of the green-white mana base.
- 1 Savannah | land | Part of the green-white mana base.
- 1 Temple Garden | land | Part of the green-white mana base.
- 1 Branchloft Pathway // Boulderloft Pathway | land | Flexible green-white land slot.
- 1 Battlefield Forge | land | Part of the red-white mana base.
- 1 Sacred Foundry | land | Part of the red-white mana base.
- 1 Needleverge Pathway // Pillarverge Pathway | land | Flexible red-white land slot.
- 1 Plateau | land | Part of the red-white mana base.
- 1 Stomping Ground | land | Part of the red-green mana base.
- 1 Taiga | land | Part of the red-green mana base.
- 1 Rockfall Vale | land | Part of the red-green mana base.
- 1 Sol Ring | ramp | Early ramp for the deck's larger Dinosaur cards.
- 1 Birds of Paradise | ramp | Early creature ramp for the three-color plan.
- 1 Arcane Signet | ramp | Reliable artifact ramp for the commander plan.
- 1 Fellwar Stone | ramp | Artifact ramp that supports the mana plan.
- 1 Drover of the Mighty | ramp | Creature ramp that fits the Dinosaur theme.
- 1 Ixalli's Lorekeeper | ramp | Creature ramp for deploying Dinosaurs.
- 1 Intrepid Paleontologist | ramp | Creature ramp supporting the Dinosaur plan.
- 1 Thunderherd Migration | ramp | Tribal ramp for the mana base.
- 1 Ranging Raptors | ramp | Dinosaur ramp for the creature plan.
- 1 Topiary Stomper | ramp | Dinosaur ramp for the larger threats.
- 1 Beast Whisperer | draw | Creature-based draw for the Dinosaur plan.
- 1 Garruk's Uprising | draw | Draw support for the deck's large creatures.
- 1 Ripjaw Raptor | draw | Dinosaur draw that adds to the creature plan.
- 1 Runic Armasaur | draw | Dinosaur draw for sustained resources.
- 1 Return of the Wildspeaker | draw | Draw support for the deck's creature-heavy plan.
- 1 Rishkar's Expertise | draw | A draw spell suited to the deck's large creatures.
- 1 Curious Altisaur | draw | Dinosaur draw for the tribal plan.
- 1 Earthshaker Dreadmaw | draw | Dinosaur draw attached to a sizable creature.
- 1 For the Ancestors | draw | Tribal draw support.
- 1 Vanquisher's Banner | draw | Tribal draw support for the Dinosaur plan.
- 1 Vaultborn Tyrant | draw | A Dinosaur draw threat for longer games.
- 1 Heroic Intervention | interaction | Interaction supporting the creature-focused plan.
- 1 Lightning Greaves | interaction | Interaction for supporting key creatures.
- 1 Swiftfoot Boots | interaction | Interaction for supporting key creatures.
- 1 Temple Altisaur | interaction | Dinosaur interaction for the tribal board.
- 1 Veil of Summer | interaction | A compact interaction slot.
- 1 And They Shall Know No Fear | interaction | Interaction supporting the creature plan.
- 1 Itzquinth, Firstborn of Gishath | removal | Low-cost Dinosaur removal.
- 1 Savage Stomp | removal | Dinosaur-themed removal.
- 1 Thrashing Brontodon | removal | Dinosaur removal on a creature body.
- 1 Tranquil Frillback | removal | Flexible Dinosaur removal.
- 1 Ravenous Sailback | removal | Dinosaur removal for the creature plan.
- 1 Burning Sun's Avatar | removal | A Dinosaur removal threat.
- 1 Trumpeting Carnosaur | removal | Dinosaur removal that remains on theme.
- 1 Kinjalli's Caller | synergy | Tribal synergy for the Dinosaur plan.
- 1 Otepec Huntmaster | synergy | Tribal synergy supporting Dinosaur deployment.
- 1 Marauding Raptor | synergy | Dinosaur synergy for the creature-focused plan.
- 1 Herald's Horn | synergy | Tribal synergy for the Dinosaur theme.
- 1 Icon of Ancestry | synergy | Tribal synergy supporting the creature plan.
- 1 Descendants' Path | synergy | Tribal synergy for a Dinosaur-heavy deck.
- 1 Shared Animosity | synergy | Tribal synergy that rewards the Dinosaur board.
- 1 Adaptive Automaton | synergy | A tribal synergy creature.
- 1 Metallic Mimic | synergy | A tribal synergy creature.
- 1 Hunting Velociraptor | synergy | Dinosaur synergy for the tribal plan.
- 1 Carnage Tyrant | threat | A substantial Dinosaur threat.
- 1 Regisaur Alpha | threat | A Dinosaur threat for building the board.
- 1 Ghalta and Mavren | threat | A legendary Dinosaur threat.
- 1 Pantlaza, Sun-Favored | threat | A Dinosaur threat for the tribal plan.
- 1 Quartzwood Crasher | threat | A Dinosaur threat that supports combat pressure.
- 1 Palani's Hatcher | threat | A Dinosaur threat for the creature plan.
- 1 Roaming Throne | threat | A tribal threat for the Dinosaur-focused deck.
- 1 Majestic Heliopterus | threat | A Dinosaur threat for the board.
- 1 Rampaging Brontodon | threat | A large Dinosaur threat.
- 1 Tyrranax Rex | threat | A major Dinosaur threat.
- 1 Zilortha, Strength Incarnate | threat | A legendary Dinosaur threat.
- 1 Thundering Spineback | threat | A Dinosaur threat for the tribal plan.
- 1 Dinosaurs on a Spaceship | wincon | A Dinosaur finisher for closing games.
- 1 Huatli, Poet of Unity // Roar of the Fifth People | wincon | A finisher that fits the deck's colors and creature plan.
- 1 Tribal Forcemage | wincon | A tribal finisher for the Dinosaur board.
- 1 Blasphemous Act | wipe | A board-wipe option when the board gets crowded.
- 1 Raging Swordtooth | wipe | A Dinosaur board-wipe option.
- 1 Wakening Sun's Avatar | wipe | A Dinosaur board-wipe option for stabilizing.

</details>

### 5. blink Commander, owned first

Format: Commander. Theme: blink. Pool: owned_first. Shortlist: 169 names.

Commander: Gilraen, Dúnedain Protector.

Grade: baseline, score 0.34, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $139.45 the whole deck.

**Summary:** This is a white blink-value deck built to establish mana, deploy Gilraen, Dúnedain Protector, and repeatedly revisit creature-based value pieces while maintaining protection and broad answers. It pressures the table through a wide creature board, Angels, and artifact creatures, then closes with its dedicated finishers, chief among them Angel of Serenity. The tradeoff is a deliberately fair, board-centric game plan: it relies on creatures staying relevant and gives up the explosive speed and dense stack interaction of stronger combo-focused builds.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

This deck plan aims for 2 or more finishers, the cards that win the game, and the deck holds 1.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core is consistent mono-white ETB-value creatures plus a handful of blink/flicker enablers around Gilraen's protection trigger, but several pieces pull against it — three board wipes and Kataki, War's Wage cut against a deck that wants a persistent creature board and runs roughly fifteen artifacts, and Puresteel Paladin sits on a thin equipment count.
- PLAN theme_fit=yes: It is a mono-white Commander deck led by a blink/protection commander with ETB creatures like Flickerwisp, Wall of Omens, Inspiring Overseer, Fiend Hunter, Angel of Sanctions and Angel of Serenity, at a fair bracket-3 power level with no fast combo or heavy tutoring.
- PLAN useful_as_built=yes: Thirty-seven lands in a single color plus about ten ramp pieces, a smooth curve, ten-plus removal spells and enough card draw means the deck casts its spells and can actually close through creature damage, even if its win routes are slow.
- PLAN summary_honest=partly: It correctly flags the fair, board-centric plan and openly reports the finisher shortfall from the quality note, but it still asserts a 'wide creature board' the list has few token or go-wide effects to produce and speaks of 'dedicated finishers' plural when the deck effectively has one.
- [INFO] `curve_summary`: average mana value 2.95 over 62 nonland cards
- [WARN] `finisher_short`: the finisher count is 1, and this deck plan wants 2 or more (Angel of Serenity)
- [INFO] `cards_trimmed`: the list was 1 card over, so the builder cut this card: The Walls of Ba Sing Se

<details><summary>The deck list</summary>

- 26 Plains | land | Reliable white mana base.
- 1 Command Tower | land | Flexible mana support.
- 1 City of Brass | land | Flexible mana support.
- 1 Exotic Orchard | land | Flexible mana support.
- 1 Marsh Flats | land | Helps stabilize land access.
- 1 Fabled Passage | land | Helps stabilize land access.
- 1 Evolving Wilds | land | Helps stabilize land access.
- 1 Terramorphic Expanse | land | Helps stabilize land access.
- 1 Minas Tirith | land | White-producing utility land.
- 1 Secluded Courtyard | land | Creature-focused mana support.
- 1 Plaza of Heroes | land | Utility land for the legendary package.
- 1 War Room | land | Utility land slot.
- 1 Sol Ring | ramp | Efficient early mana acceleration.
- 1 Arcane Signet | ramp | Reliable mana acceleration.
- 1 Fellwar Stone | ramp | Low-cost mana acceleration.
- 1 Wayfarer's Bauble | ramp | Early land-based acceleration.
- 1 Sword of the Animist | ramp | Creature-based mana development.
- 1 Giada, Font of Hope | ramp | Creature ramp that fits the Angel contingent.
- 1 Chromatic Lantern | ramp | Fixing and mana acceleration.
- 1 Relic of Legends | ramp | Supports the deck's legendary creatures while accelerating mana.
- 1 Thought Vessel | ramp | Low-cost mana acceleration.
- 1 White Lotus Tile | ramp | Additional mana acceleration.
- 1 Wall of Omens | draw | Early creature-based card flow and a useful blink target.
- 1 Inspiring Overseer | draw | Creature-based card flow that fits the blink shell.
- 1 Champions of Minas Tirith | draw | Adds card flow on a creature body.
- 1 Faramir, Field Commander | draw | Adds card flow while contributing to the creature plan.
- 1 Puresteel Paladin | draw | Card flow for the Equipment subpackage.
- 1 Skullclamp | draw | Efficient card-flow tool.
- 1 Mask of Memory | draw | Equipment-based card flow.
- 1 Idol of Oblivion | draw | Low-cost persistent card-flow option.
- 1 Lembas | draw | Compact card-flow piece.
- 1 Tome of Legends | draw | Steady card flow alongside the commander.
- 1 Energybending | draw | Flexible card-flow spell.
- 1 Bastion Protector | interaction | Helps keep the commander protected.
- 1 Boromir, Warden of the Tower | interaction | Protective creature interaction.
- 1 Clever Concealment | interaction | Protects a developed board.
- 1 Lightning Greaves | interaction | Efficient commander protection.
- 1 Swiftfoot Boots | interaction | Additional commander protection.
- 1 Gift of Immortality | interaction | Protects an important creature through removal.
- 1 Reprieve | interaction | Flexible stack interaction.
- 1 Ultimate Magic: Holy | interaction | Additional protective interaction.
- 1 Fiend Hunter | removal | Creature-based removal that works well with blink support.
- 1 Palace Jailer | removal | Creature-based removal for the blink plan.
- 1 Generous Gift | removal | Broad permanent removal.
- 1 Get Lost | removal | Efficient flexible removal.
- 1 Dispatch | removal | Low-cost targeted removal.
- 1 Destroy Evil | removal | Flexible targeted removal.
- 1 Crib Swap | removal | Creature removal that fits the deck's creature focus.
- 1 Swords to Plowshares | removal | Efficient creature removal.
- 1 March of Otherworldly Light | removal | Flexible removal at multiple points on the curve.
- 1 Flickerwisp | synergy | Core blink piece for reusing creature value.
- 1 Personify | synergy | Directly supports the blink-focused plan.
- 1 Ennis, Debate Moderator | synergy | Supports the deck's dedicated synergy package.
- 1 Jocasta, Automaton Avenger | synergy | Supports the deck's dedicated synergy package.
- 1 Slip On the Ring | synergy | A compact blink effect that protects or reuses a creature.
- 1 Angel of Condemnation | synergy | Creature-based blink support for the value plan.
- 1 Joined Researchers // Secret Rendezvous | synergy | A creature-based value piece to revisit with blink effects.
- 1 Exemplar of Light | threat | A meaningful creature threat that also fits the Angel core.
- 1 Frontline Medic | threat | Develops the board while presenting a creature threat.
- 1 Westfold Rider | threat | Adds to the creature pressure plan.
- 1 Zack Fair | threat | Legendary creature pressure that supports the board plan.
- 1 Bronze Guardian | threat | Durable artifact-creature threat.
- 1 Grim Poppet | threat | Creature threat with useful board presence.
- 1 Kataki, War's Wage | threat | Creature pressure that can strain opposing artifact plans.
- 1 PuPu UFO | threat | Adds an artifact-creature body to the pressure plan.
- 1 Summon: Primal Garuda | threat | Creature-based pressure with useful utility.
- 1 Weapons Vendor | threat | Creature pressure that fits the Equipment elements.
- 1 Angel of Sanctions | threat | Evasive creature pressure for the blink shell.
- 1 Angel of Serenity | wincon | Primary dedicated closing threat.
- 1 Aang, the Last Airbender | wincon | High-impact legendary creature for closing games.
- 1 The Vision | wincon | A dedicated artifact-creature closer.
- 1 Austere Command | wipe | Flexible reset button for difficult boards.
- 1 Fumigate | wipe | Creature-board reset when pressure gets overwhelming.
- 1 Vanquish the Horde | wipe | Efficient mass creature answer.

</details>

### 6. Modern tempo, tournament

Format: Modern. Theme: tempo. Pool: any_card. Shortlist: 167 names.

Grade: typical, score 0.59, model 20260914T154223Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $380.09 to buy, $380.09 the whole deck.

**Summary:** This blue-red tempo deck keeps a low curve, uses draw to keep relevant spells coming, and protects its position with efficient interaction and removal. It aims to establish creature pressure, chiefly through Ragavan, Nimble Pilferer and Ledger Shredder, then maintain the initiative with a dense stream of spells and Temporal Mastery. It gives up broad permanent coverage and a heavy late-game threat package in favor of speed, consistency, and early pressure.

The quality model grades this deck typical against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade. The deck runs its spells as playsets, and that lowers the grade.

- PLAN plan_coherent=partly: The cantrip-plus-cheap-interaction shell is a coherent tempo skeleton, but four copies of Temporal Mastery — an extra-turn miracle that only pays off with a fast clock — pull against a threat base of just six creatures, so the 'pressure' half of the plan is undersupported.
- PLAN theme_fit=partly: The colors, format, curve, and tempo shell match the request, but a Delver deck with zero Delver of Secrets (and no comparable one-mana threat) is not quite the deck asked for.
- PLAN useful_as_built=no: The mana and interaction are fine, but six total threats with no other route to damage means the deck often has nothing to pressure with, and Temporal Mastery's extra turns are near-blank without a board, leaving too few ways to actually win.
- PLAN summary_honest=partly: It correctly names Ragavan and Ledger Shredder as the threats and openly concedes the thin late game, but it never mentions that the requested Delver is absent and it oversells 'creature pressure' from only six bodies.
- [INFO] `curve_summary`: average mana value 1.78 over 36 nonland cards

<details><summary>The deck list</summary>

- 6 Island | land | Provides a reliable blue source for the spell-heavy core.
- 4 Mountain | land | Provides a reliable red source for the removal package.
- 4 Steam Vents | land | Supplies both deck colors without using a tapped land slot.
- 4 Spirebluff Canal | land | Supports the early blue-red tempo plan.
- 3 Riverglide Pathway // Lavaglide Pathway | land | Flexibly fills either color requirement.
- 2 Shivan Reef | land | Adds another untapped source for both colors.
- 1 Fiery Islet | land | Completes the mana base with an additional red-blue source.
- 4 Consider | draw | Keeps cards flowing while maintaining a low curve.
- 4 Mishra's Bauble | draw | Provides inexpensive card flow for the tempo shell.
- 2 Ledger Shredder | draw | Adds a creature presence while contributing to card flow.
- 4 Preordain | draw | Helps the deck keep finding its next relevant spell.
- 2 Counterspell | interaction | Provides direct protection against opposing plays.
- 2 Force of Negation | interaction | Adds interaction that supports a proactive tempo posture.
- 2 Spell Pierce | interaction | Gives the deck an efficient early interaction option.
- 4 Lightning Bolt | removal | Forms the most efficient core of the red removal suite.
- 4 Galvanic Discharge | removal | Expands the deck's cheap removal density.
- 4 Temporal Mastery | synergy | Supplies the shortlist's available blue-red tempo synergy piece.
- 4 Ragavan, Nimble Pilferer | ramp | Provides the available low-cost ramp element while adding to the board.

</details>

### 7. Modern burn, casual

Format: Modern. Theme: burn. Pool: any_card. Shortlist: 290 names.

Grade: bad, score 0.02, model 20260914T154223Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $76.16 to buy, $76.16 the whole deck.

**Summary:** This mono-red burn deck keeps the pressure on with efficient removal, a dense suite of spell-driven damage synergies, and creatures that continue attacking when direct damage alone is not enough. It wins by steadily converting red spells and aggressive threats into damage, with Torbran, Thane of Red Fell chief among the creature threats. The deck gives up broad answers and defensive tools in favor of speed, redundancy, and a straightforward damage race.

The quality model grades this deck below the precon baseline against the top lists of the format: the deck runs many spells as one copy, and that lowers the grade. The deck runs few playsets, and that raises the grade. The cards pair in ways the top lists do not, and that lowers the grade.

- PLAN plan_coherent=no: The prowess/pinger shell (Electrostatic Field, Thermo-Alchemist, Guttersnipe, Kessig Flamebreather) wants a dense pile of cheap spells but the deck has only ~14 nonland spells, half of them weak cantrips, while a scattered top-end of Lukka, Ashcloud Phoenix, Hazoret, and Purphoros pulls toward midrange — three plans, none supported.
- PLAN theme_fit=partly: It is mono-red and Modern-legal, but a deck with six actual burn spells (4 Lightning Bolt, 2 Burst Lightning) and a singleton toolbox of creatures and planeswalkers is not the burn deck that was requested.
- PLAN useful_as_built=partly: The mana base and curve are functional enough to shuffle up and cast things casually, but with almost every effect a one-of the deck draws wildly inconsistent hands and rarely assembles the spell density its pingers need to close a game.
- PLAN summary_honest=no: placeholder
- [INFO] `curve_summary`: average mana value 2.53 over 36 nonland cards

<details><summary>The deck list</summary>

- 20 Mountain | land | Provides the mono-red mana base without entering tapped.
- 1 Barbarian Ring | land | Adds a red land slot to the mana base.
- 1 Ramunap Ruins | land | Adds a red land slot to the mana base.
- 1 Sokenzan, Crucible of Defiance | land | Adds a red land slot to the mana base.
- 1 Sunscorched Desert | land | Completes the land package while staying in mono-red.
- 1 Ancestral Anger | draw | Low-cost draw support for the spell-heavy plan.
- 1 Crimson Wisps | draw | Efficient draw support that fits the deck's red spells.
- 1 Expedite | draw | Provides inexpensive draw support.
- 1 Fists of Flame | draw | Adds draw support while contributing to an aggressive turn.
- 1 Renegade Tactics | draw | A cheap draw spell for keeping the deck moving.
- 1 Risk Factor | draw | Adds a damage-focused draw option.
- 1 Flick a Coin | ramp | Provides ramp support in a compact spell slot.
- 1 Sudden Breakthrough | ramp | Provides the second ramp effect for the deck.
- 4 Lightning Bolt | removal | Core efficient burn removal for clearing resistance or pressing damage.
- 2 Burst Lightning | removal | Additional flexible burn removal.
- 1 Eidolon of the Great Revel | synergy | A damage-focused synergy piece for the low-cost spell plan.
- 1 Electrostatic Field | synergy | Rewards the deck's high density of instants and sorceries.
- 1 Firebrand Archer | synergy | Adds spell-driven damage synergy.
- 1 Guttersnipe | synergy | A key payoff for repeatedly casting burn spells.
- 1 Kessig Flamebreather | synergy | Adds another spell-focused damage payoff.
- 1 Thermo-Alchemist | synergy | Supports the repeated-spell damage plan.
- 1 Roiling Vortex | synergy | Provides a persistent damage-oriented synergy piece.
- 1 Unstable Amulet | synergy | Adds another synergy card to support the burn plan.
- 1 Ashcloud Phoenix | threat | A resilient threat for games that go beyond the opening burn exchange.
- 1 Barret Wallace | threat | Provides an additional creature threat.
- 1 Cemetery Gatekeeper | threat | A proactive threat that fits the red creature package.
- 1 Fuming Effigy | threat | Adds a threat at the top of the creature curve.
- 1 Hazoret the Fervent | threat | A powerful damage-focused threat for closing games.
- 1 Imodane, the Pyrohammer | threat | A threat aligned with the deck's damage spells.
- 1 Keral Keep Disciples | threat | Adds another red threat to maintain board pressure.
- 1 Lukka, Coppercoat Outcast | threat | Provides a noncreature threat that broadens the deck's pressure.
- 1 Purphoros, God of the Forge | threat | A durable threat for damage-oriented games.
- 1 Sunspine Lynx | threat | A red threat for finishing longer games.
- 1 Tectonic Giant | threat | Adds a substantial midgame threat.
- 1 Thundering Raiju | threat | An aggressive creature threat.
- 1 Torbran, Thane of Red Fell | threat | Chief among the creature threats for amplifying the deck's damage focus.
- 1 Zo-Zu the Punisher | threat | A pressure-oriented threat that complements the aggressive plan.

</details>

### 8. Modern lifegain, FNM

Format: Modern. Theme: lifegain. Pool: any_card. Shortlist: 301 names.

Grade: bad, score 0.33, model 20260914T154223Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $249.48 to buy, $249.48 the whole deck.

**Summary:** This white-black lifegain deck establishes an early engine with Soul Warden and Ajani's Pridemate, then converts that foundation into creature pressure through cards such as Attended Healer, Enduring Tenacity, and Gideon's Company. Solitude and Deep-Cavern Bat help clear a path while Enduring Innocence and Dawn of Hope keep resources flowing. It wins by building a board of lifegain-focused threats and pressing that advantage, giving up some speed to dedicated aggressive decks in exchange for a more durable, value-oriented game plan.

The quality model grades this deck below the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade.

- PLAN plan_coherent=yes: Every nonland card either gains life (Soul Warden, Attended Healer, Gideon's Company, Solitude) or pays off lifegain (Ajani's Pridemate, Bloodbond Vampire, Enduring Tenacity, Enduring Innocence, Dawn of Hope), so the pieces all feed the same engine.
- PLAN theme_fit=yes: It is exactly a two-color white-black lifegain 60-card Modern list at a reasonable FNM power level, with no named constraint ignored.
- PLAN useful_as_built=partly: The mana base of 24 lands with four WB duals and ample basics supports the colors well and there are clear win routes in Pridemate/Tenacity/Gideon's Company, but the curve is crowded with roughly fourteen four-drops plus a six-mana Solitude, and dedicated removal is thin, so games can stall on clunky draws.
- PLAN summary_honest=partly: The description matches the list closely, though it lists Enduring Tenacity as a source of "creature pressure" when it is a drain enchantment, and it does not mention the Altar of the Pantheon/Bloodbond Vampire filler.
- [INFO] `curve_summary`: average mana value 3.17 over 36 nonland cards

<details><summary>The deck list</summary>

- 4 Godless Shrine | land | Provides both deck colors while keeping the mana base focused.
- 4 Caves of Koilos | land | Provides flexible white or black mana.
- 4 Shattered Sanctum | land | Adds dual-color mana for the early turns.
- 4 Isolated Chapel | land | Supports the white-black mana base.
- 4 Plains | land | Supplies reliable white mana.
- 4 Swamp | land | Supplies reliable black mana.
- 2 Altar of the Pantheon | ramp | Adds the requested mana acceleration.
- 4 Enduring Innocence | draw | Provides a dedicated source of cards.
- 2 Dawn of Hope | draw | Adds further card access for the lifegain plan.
- 4 Solitude | removal | Handles opposing problems while fitting the white half of the deck.
- 2 Deep-Cavern Bat | removal | Provides additional disruption in a creature slot.
- 4 Soul Warden | synergy | Forms an early lifegain engine for the deck's payoffs.
- 4 Ajani's Pridemate | synergy | Turns the lifegain theme into a growing board presence.
- 4 Enduring Tenacity | threat | Serves as a resilient threat within the black lifegain shell.
- 4 Attended Healer | threat | Builds the deck's creature pressure around the lifegain theme.
- 2 Bloodbond Vampire | threat | Adds more lifegain-oriented pressure at the middle of the curve.
- 4 Gideon's Company | threat | Provides a large lifegain-focused finisher for the creature plan.

</details>

### 9. Standard midrange, FNM

Format: Standard. Theme: midrange. Pool: any_card. Shortlist: 164 names.

Grade: bad, score 0.00, model 20260914T154223Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $130.03 to buy, $130.03 the whole deck.

**Summary:** This black-green midrange deck develops its mana, trades efficiently with opposing threats, and steadily builds a creature board backed by protection and card advantage. Its wins come from surviving the early exchanges and turning its larger creatures into sustained pressure, with Goldvein Hydra, Vein Ripper, Vaultborn Tyrant, and Zodiark, Umbral God chief among them. It gives up some speed and consistency for a broad spread of individually useful threats, answers, and support cards.

The quality model grades this deck below the precon baseline against the top lists of the format: the deck runs many spells as one copy, and that lowers the grade. The cards pair in ways the top lists do not, and that lowers the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The deck is broadly a BG goodstuff midrange pile, but its one-of support cards pull several ways at once — combat tricks/protection auras, death-trigger recursion, an Insidious Roots graveyard engine, and a top-heavy pile of six-plus drops — none of which are supported by enough copies or enablers to form a single plan.
- PLAN theme_fit=yes: It is a Standard-legal black-green midrange list with removal, card draw, and creature threats, matching the requested colors, format, and casual FNM ambition.
- PLAN useful_as_built=partly: The 24-land base with Blooming Marsh and Overgrown Tomb casts the spells fine and there are plenty of win conditions, but with only two Llanowar Elves as ramp and a heavy cluster of five-to-seven drops, plus nearly every effect as a single copy, the deck will often stall on clunky draws.
- PLAN summary_honest=yes: The summary accurately describes a slow, threat-dense midrange pile and openly admits the singleton spread costs speed and consistency, claiming no engine the list does not contain.
- [INFO] `curve_summary`: average mana value 3.06 over 36 nonland cards
- [INFO] `mana_pass`: the builder moved 1 card of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 4 Blooming Marsh | land | Provides black-green mana without taking a tapped-land slot.
- 4 Overgrown Tomb | land | Provides both deck colors for the midrange shell.
- 7 Forest | land | Supplies a dependable green base for creature threats and support cards.
- 9 Swamp | land | Supplies a dependable black base for removal and threats.
- 2 Llanowar Elves | ramp | Accelerates the deck into its larger midrange plays.
- 1 Corrupted Conviction | draw | Provides a compact card-advantage option.
- 1 Darkstar Augur | draw | Adds a creature-based source of card advantage.
- 1 Phyrexian Arena | draw | Supplies ongoing card advantage for longer games.
- 1 Midnight Reaper | draw | Keeps cards flowing alongside the creature plan.
- 1 Stocking the Pantry | draw | Adds a dedicated card-advantage piece to the midrange plan.
- 1 Unholy Annex // Ritual Chamber | draw | Provides another durable card-advantage option.
- 1 Assassin's Trophy | removal | Serves as flexible removal.
- 1 Bitter Triumph | removal | Provides efficient removal for opposing problems.
- 1 Hero's Downfall | removal | Adds reliable targeted removal.
- 1 Maelstrom Pulse | removal | Offers a versatile removal slot.
- 1 Pick Your Poison | removal | Gives the deck another low-cost answer.
- 1 Withering Torment | removal | Rounds out the targeted removal package.
- 1 Garruk's Uprising | synergy | Supports a board built around substantial creature threats.
- 1 Insidious Roots | synergy | Rewards the deck for maintaining its creature-focused game plan.
- 1 Innkeeper's Talent | synergy | Supports the deck's permanent-based midrange plan.
- 1 Not Dead After All | synergy | Helps preserve an important creature threat through opposition.
- 1 Royal Treatment | synergy | Protects a key permanent while advancing the creature plan.
- 1 Snakeskin Veil | synergy | Provides inexpensive support for a valuable threat.
- 1 Undying Malice | synergy | Helps the deck retain pressure after an opposing answer.
- 1 Vraska Joins Up | synergy | Supports a midrange board featuring notable creatures.
- 1 Goldvein Hydra | threat | Presents a scalable green creature threat.
- 1 Crossway Troublemakers | threat | Adds a black creature that contributes pressure and card access.
- 1 Earthshaker Dreadmaw | threat | Provides a large green creature for the top of the curve.
- 1 Massacre Girl, Known Killer | threat | Adds a legendary black creature that pressures creature decks.
- 1 Rottenmouth Viper | threat | Supplies a black creature that also fits the interaction-heavy plan.
- 1 Sephiroth, Fabled SOLDIER // Sephiroth, One-Winged Angel | threat | Provides a prominent legendary creature threat.
- 1 Surrak, Elusive Hunter | threat | Adds a green legendary creature to the attacking plan.
- 1 Thrashing Brontodon | threat | Provides a creature threat with utility against opposing permanents.
- 1 Vaultborn Tyrant | threat | Serves as a large creature threat for drawn-out games.
- 1 Vein Ripper | threat | Adds a powerful black creature threat.
- 1 Doomsday Excruciator | threat | Provides a high-impact black creature at the top end.
- 1 Chomping Changeling | threat | Adds another creature that can apply pressure while fitting the removal plan.
- 1 Reclamation Sage | threat | Contributes a creature body while offering useful utility.
- 1 Zodiark, Umbral God | threat | Provides a legendary top-end creature threat.

</details>

### 10. Standard aggro, tournament

Format: Standard. Theme: aggro. Pool: any_card. Shortlist: 294 names.

Grade: bad, score 0.06, model 20260914T154223Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $134.80 to buy, $134.80 the whole deck.

**Summary:** This red-white aggro deck aims to keep combat pressure on the opponent with a dense creature package, backed by removal and interaction to push through resistance. Warleader's Call is chief among its supporting cards, while Fugitive Codebreaker and Caretaker's Talent help keep resources flowing. It wins by maintaining a steady attack, giving up larger late-game threats and board-reset effects for a focused proactive plan.

The quality model grades this deck below the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. Many of the cards appear in no top list, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core is a beatdown plan with anthems and cheap interaction, but the curve is stacked at four mana (Warleader's Call, Caretaker's Talent, Redcap Gutter-Dweller) with no one-drops and a six-mana dragon, and Caretaker's Talent pulls toward a token/value axis the rest of the list barely supports.
- PLAN theme_fit=yes: It is a two-color red-white 60-card Standard list built around creatures, combat anthems, and removal, which is the aggro deck the request named, even if it leans midrange-heavy at the top.
- PLAN useful_as_built=partly: The mana base is clean and consistent and there are ample threats to close a game, but an aggro deck with zero one-drops and a glut of four-mana cards will often be doing nothing on the turns it most needs pressure.
- PLAN summary_honest=partly: placeholder
- [INFO] `curve_summary`: average mana value 3.22 over 36 nonland cards

<details><summary>The deck list</summary>

- 4 Sacred Foundry | land | Red-white land base anchor with both basic land types.
- 4 Inspiring Vantage | land | Provides both deck colors while supporting the aggressive mana base.
- 4 Sunbillow Verge | land | Adds flexible red-white color access.
- 6 Mountain | land | Reliable red source for the red portion of the deck.
- 6 Plains | land | Reliable white source for the white portion of the deck.
- 4 Fugitive Codebreaker | draw | Red creature that supplies the deck's draw package.
- 2 Caretaker's Talent | draw | White draw option that helps sustain pressure.
- 4 Boros Charm | interaction | Red-white interaction suited to the deck's colors.
- 2 Aven Interrupter | interaction | Creature-based interaction that can join the attack.
- 4 Emeritus of Truce // Swords to Plowshares | removal | White removal that fits the main-deck removal suite.
- 4 Krenko's Buzzcrusher | removal | Red removal option for clearing opposing resistance.
- 4 Warleader's Call | synergy | Red-white synergy piece for the creature-focused plan.
- 4 Dragonback Lancer | threat | Creature threat for applying combat pressure.
- 4 Diversion Specialist | threat | Creature threat that broadens the attacking roster.
- 4 Redcap Gutter-Dweller | threat | Red creature threat for the aggressive game plan.

</details>

### 11. Commander with a locked card

Format: Commander. Theme: sacrifice. Pool: any_card. Shortlist: 299 names.

Commander: Karlov of the Ghost Council.

Grade: baseline, score 0.06, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $1625.61 to buy, $1625.61 the whole deck.

**Summary:** Karlov of the Ghost Council leads an Orzhov sacrifice deck that develops its mana, turns creatures into cards and removal, and uses resilient sacrifice outlets to keep the board working in its favor. The deck closes through dedicated finishers, persistent sacrifice payoffs, and a selection of larger creature threats. It gives up some speed to maintain a broad creature package and relies on keeping enough creatures or tokens available to fuel its engines.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. Few lands enter tapped, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every card is a sacrifice outlet, a token/body producer, or an aristocrats drain/lifegain payoff that also feeds Karlov's counters, with only a couple of strays like Kuldotha Forgemaster (few artifacts to fetch) pulling off-plan.
- PLAN theme_fit=yes: It is a Commander deck led by Karlov of the Ghost Council in Orzhov colors built squarely around sacrificing creatures, exactly as requested.
- PLAN useful_as_built=yes: 36 mostly untapped dual/fixing lands support a low curve, there are ample sacrifice outlets and fodder, and multiple real win routes (Zulaport/Bastion/Elas drains, Altar of Dementia mill, Karlov's exile ability, Sephiroth) make it playable out of the box.
- PLAN summary_honest=yes: The claims — mana development, card-draw and removal from sacrificing, resilient outlets, drain finishers, and the caveat about needing bodies to fuel engines — all match what the list actually contains, though calling Relic Vial a finisher is generous.
- [INFO] `curve_summary`: average mana value 3.27 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Brightclimb Pathway // Grimclimb Pathway | land | Flexible land slot for the mana base.
- 1 Caves of Koilos | land | Dual-color land slot for the mana base.
- 1 City of Brass | land | Flexible color source for the mana base.
- 1 Command Tower | land | Reliable color source for the mana base.
- 1 Concealed Courtyard | land | Dual-color land slot for the mana base.
- 1 Diamond Valley | land | Sacrifice-oriented land slot in the mana base.
- 1 Eclipsed Steppe | land | Dual-color land slot for the mana base.
- 1 Forbidden Orchard | land | Flexible color source for the mana base.
- 1 Godless Shrine | land | Dual-color land slot for the mana base.
- 1 High Market | land | Sacrifice-oriented land slot in the mana base.
- 1 Isolated Chapel | land | Dual-color land slot for the mana base.
- 1 Mana Confluence | land | Flexible color source for the mana base.
- 1 Marsh Flats | land | Mana-base fixing land.
- 1 Phyrexian Tower | land | Sacrifice-oriented land slot in the mana base.
- 1 Prismatic Vista | land | Mana-base fixing land.
- 1 Scrubland | land | Dual-color land slot for the mana base.
- 1 Shattered Sanctum | land | Dual-color land slot for the mana base.
- 1 Silent Clearing | land | Dual-color land slot for the mana base.
- 1 Starting Town | land | Dual-color land slot for the mana base.
- 1 Starlit Sanctum | land | Dual-color land slot for the mana base.
- 1 Tarnished Citadel | land | Flexible color source for the mana base.
- 1 Vault of Champions | land | Dual-color land slot for the mana base.
- 7 Plains | land | Basic white mana source.
- 7 Swamp | land | Basic black mana source.
- 1 Sol Ring | ramp | Efficient mana acceleration kept at the deck's core.
- 1 Ashnod's Altar | ramp | Turns sacrificed creatures into mana.
- 1 Culling the Weak | ramp | Sacrifice-based burst mana.
- 1 Deadly Dispute | ramp | Converts a sacrifice into mana support.
- 1 Phyrexian Altar | ramp | Turns sacrificed creatures into mana.
- 1 Pitiless Plunderer | ramp | Rewards the sacrifice plan with mana resources.
- 1 Priest of Forgotten Gods | ramp | Sacrifice-based mana development.
- 1 Skullport Merchant | ramp | Provides sacrifice-oriented mana support.
- 1 Sifter of Skulls | ramp | Rewards creature deaths with mana resources.
- 1 Warren Soultrader | ramp | Adds mana support to the sacrifice shell.
- 1 Ahriman | draw | Adds card advantage to sustain the deck.
- 1 Baron Bertram Graywater | draw | Provides card advantage in the creature shell.
- 1 Bushmeat Poacher | draw | Turns creature resources into cards.
- 1 Corrupted Conviction | draw | Efficient card draw for a sacrificed creature.
- 1 Disciple of Bolas | draw | Converts a creature into card advantage.
- 1 Lord Skitter's Butcher | draw | Supplies card advantage within the sacrifice plan.
- 1 Shadowheart, Dark Justiciar | draw | Turns creature resources into cards.
- 1 Smothering Abomination | draw | Rewards the sacrifice plan with cards.
- 1 Tevesh Szat, Doom of Fools | draw | Provides a substantial card-advantage engine.
- 1 Vampiric Rites | draw | Repeatable sacrifice-based card draw.
- 1 Village Rites | draw | Efficient card draw for a sacrificed creature.
- 1 Cartel Aristocrat | interaction | Flexible protective sacrifice outlet.
- 1 Dark Privilege | interaction | Protects key creatures through sacrifice play.
- 1 Eldrazi Monument | interaction | Protects the board while using creature resources.
- 1 Fanatical Devotion | interaction | Protects important creatures with sacrifice synergies.
- 1 Flare of Fortitude | interaction | Defensive interaction for the board.
- 1 Gift of Doom | interaction | Protects a key permanent.
- 1 Nightmare Shepherd | interaction | Adds resilient interaction to creature exchanges.
- 1 Promise of Tomorrow | interaction | Protects the creature-focused plan.
- 1 Attrition | removal | Sacrifice outlet that converts creatures into removal.
- 1 Ayli, Eternal Pilgrim | removal | Sacrifice-oriented creature removal.
- 1 Blasting Station | removal | Sacrifice outlet that provides removal reach.
- 1 Bone Shards | removal | Low-cost removal that accepts a sacrifice.
- 1 Bone Splinters | removal | Low-cost removal that accepts a sacrifice.
- 1 Dictate of Erebos | removal | Makes creature sacrifices pressure opposing boards.
- 1 Spark Harvest | removal | Removal that can use a creature as its cost.
- 1 Teysa, Orzhov Scion | removal | Creature-based removal suited to the sacrifice plan.
- 1 Yawgmoth, Thran Physician | removal | Sacrifice-based repeatable removal.
- 1 Austere Command | wipe | Flexible board reset.
- 1 The Meathook Massacre | wipe | Board wipe that also serves as a finisher.
- 1 Toxic Deluge | wipe | Efficient reset when the board gets away.
- 1 Altar of Dementia | synergy | Sacrifice outlet that advances the deck's plan.
- 1 Carrion Feeder | synergy | Low-cost repeatable sacrifice outlet.
- 1 Elas il-Kor, Sadistic Pilgrim | synergy | Sacrifice payoff and finisher.
- 1 Viscera Seer | synergy | Low-cost repeatable sacrifice outlet.
- 1 Woe Strider | synergy | Provides a sacrifice outlet for the creature shell.
- 1 Zulaport Cutthroat | synergy | Sacrifice payoff and finisher.
- 1 Bastion of Remembrance | synergy | Persistent sacrifice payoff and finisher.
- 1 Blood Host | threat | Creature threat that benefits from sacrifices.
- 1 Felisa, Fang of Silverquill | threat | Creature threat that supports the board-building plan.
- 1 Ghoulcaller Gisa | threat | Creature threat for the sacrifice-focused board.
- 1 Kuldotha Forgemaster | threat | Large artifact creature threat.
- 1 Liesa, Forgotten Archangel | threat | Powerful creature threat.
- 1 Mondrak, Glory Dominus | threat | Creature threat for a token-oriented board.
- 1 Ratadrabik of Urborg | threat | Legendary creature threat for the board plan.
- 1 Reaper of Flight Moonsilver | threat | Evasive creature threat.
- 1 Sadistic Hypnotist | threat | Creature threat with sacrifice synergy.
- 1 Sidisi, Undead Vizier | threat | Creature threat for the sacrifice shell.
- 1 Titan Hunter | threat | Creature threat that pressures the table.
- 1 Undercity Necrolisk | threat | Creature threat for the midgame board.
- 1 Razaketh, the Foulblooded | wincon | Dedicated finisher and endgame win condition.
- 1 Relic Vial | wincon | Dedicated sacrifice-based finisher.
- 1 Sephiroth, Fabled SOLDIER // Sephiroth, One-Winged Angel | wincon | Dedicated finisher for closing games.

</details>

### 12. Commander on a budget

Format: Commander. Theme: tokens. Pool: any_card. Shortlist: 298 names.

Commander: Adeline, Resplendent Cathar.

Grade: bad, score 0.02, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $96.06 to buy, $96.06 the whole deck.

**Summary:** Adeline leads a steady mono-white token plan: establish a broad board, reinforce that board with token-focused permanents, and keep momentum through plentiful draw, removal, and a compact suite of sweepers. Halo Fountain, Luck Bobblehead, and Sword of Body and Mind provide the closing routes, chief among them alongside the marked finishers Requiem Angel. The deck gives up premium mana and costly standalone threats for a dependable basic-land base, modest curve, and broad access to inexpensive support pieces.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The deck makes less mana on turn four than the norm, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every slot makes creature tokens, anthems them, or protects/refuels the wide board, which is exactly how Adeline converts a crowded battlefield into damage, with only a handful of weak trinket-ramp pieces (Coin of Mastery, Goldvein Pick, Druidic Satchel) sitting loosely in that plan.
- PLAN theme_fit=yes: It is a mono-white Adeline token Commander deck built from cheap, widely available commons and mid-tier rares with no expensive staples or fast mana, matching the budget and bracket 2 target the person asked for.
- PLAN useful_as_built=yes: 37 lands with a heavily basic base support a mono-white curve topping out around five or six mana, and the deck has plenty of token makers, anthems, removal, and both combat and alt-win closers to actually finish games.
- PLAN summary_honest=partly: The draw, removal and sweeper claims check out, but the garbled "closing routes" sentence oversells Sword of Body and Mind and Requiem Angel as finishers while underselling the deck's actual win route of Adeline plus anthems and go-wide combat.
- [INFO] `curve_summary`: average mana value 3.48 over 62 nonland cards

<details><summary>The deck list</summary>

- 1 Ancient Den | land | White-producing artifact land that supports the mana base.
- 1 Castle Ardenvale | land | White land slot for the mana base.
- 1 Minas Tirith | land | White land slot for the mana base.
- 1 Mirrex | land | Utility land slot in the mana base.
- 1 Path of Ancestry | land | Mana-base land that fits the commander’s creature type.
- 30 Plains | land | Reliable untapped white sources for a mono-white deck.
- 1 Secluded Courtyard | land | Creature-focused colored source for the mana base.
- 1 Windbrisk Heights | land | Utility land slot in the mana base.
- 1 Coin of Mastery | ramp | Provides the deck's designated ramp.
- 1 Druidic Satchel | ramp | Provides the deck's designated ramp.
- 1 Goldvein Pick | ramp | Provides the deck's designated ramp.
- 1 Idol of False Gods | ramp | Provides the deck's designated ramp.
- 1 Karn, Living Legacy | ramp | Provides the deck's designated ramp.
- 1 Keeper of the Accord | ramp | Provides the deck's designated ramp.
- 1 Legion's Landing // Adanto, the First Fort | ramp | Provides the deck's designated ramp.
- 1 Monologue Tax | ramp | Provides the deck's designated ramp.
- 1 The Restoration of Eiganjo // Architect of Restoration | ramp | Provides the deck's designated ramp.
- 1 Thousand Moons Smithy // Barracks of the Thousand | ramp | Provides the deck's designated ramp.
- 1 Angelic Sell-Sword | draw | Provides the deck's designated card draw.
- 1 Bygone Bishop | draw | Provides the deck's designated card draw.
- 1 Court of Grace | draw | Provides the deck's designated card draw.
- 1 Dawn of Hope | draw | Provides the deck's designated card draw.
- 1 Glimmer Seeker | draw | Provides the deck's designated card draw.
- 1 Idol of Oblivion | draw | Provides the deck's designated card draw.
- 1 Nexus of Becoming | draw | Provides the deck's designated card draw.
- 1 Platoon Dispenser | draw | Provides the deck's designated card draw.
- 1 Staff of the Storyteller | draw | Provides the deck's designated card draw.
- 1 Sunpearl Kirin | draw | Provides the deck's designated card draw.
- 1 Wedding Announcement // Wedding Festivity | draw | Provides the deck's designated card draw.
- 1 Aerial Assault | removal | Provides the deck's designated removal.
- 1 Banishing Slash | removal | Provides the deck's designated removal.
- 1 Generous Gift | removal | Provides the deck's designated removal.
- 1 Kellan's Lightblades | removal | Provides the deck's designated removal.
- 1 Righteous Confluence | removal | Provides the deck's designated removal.
- 1 Skyclave Apparition | removal | Provides the deck's designated removal.
- 1 Stroke of Midnight | removal | Provides the deck's designated removal.
- 1 Clarion Spirit | synergy | Token-focused synergy piece for the deck.
- 1 Divine Visitation | synergy | Token-focused synergy piece for the deck.
- 1 Felidar Retreat | synergy | Token-focused synergy piece for the deck.
- 1 Intangible Virtue | synergy | Token-focused synergy piece for the deck.
- 1 Oketra's Monument | synergy | Token-focused synergy piece for the deck.
- 1 Retreat to Emeria | synergy | Token-focused synergy piece for the deck.
- 1 Rosie Cotton of South Lane | synergy | Token-focused synergy piece for the deck.
- 1 Siege Veteran | synergy | Token-focused synergy piece for the deck.
- 1 Skrelv's Hive | synergy | Token-focused synergy piece for the deck.
- 1 Twilight Drover | synergy | Token-focused synergy piece for the deck.
- 1 Archon of Sun's Grace | threat | A designated threat that advances board presence.
- 1 Attended Healer | threat | A designated threat that advances board presence.
- 1 Basri's Lieutenant | threat | A designated threat that advances board presence.
- 1 Cemetery Protector | threat | A designated threat that advances board presence.
- 1 Defiler of Faith | threat | A designated threat that advances board presence.
- 1 Emeria Angel | threat | A designated threat that advances board presence.
- 1 Gideon, Ally of Zendikar | threat | A designated threat that advances board presence.
- 1 Hero of Bladehold | threat | A designated threat that advances board presence.
- 1 Oketra the True | threat | A designated threat that advances board presence.
- 1 Phantom General | threat | A designated threat that advances board presence.
- 1 Requiem Angel | threat | A designated threat and marked finisher.
- 1 Threefold Thunderhulk | threat | A designated threat that advances board presence.
- 1 Halo Fountain | wincon | A designated win condition and marked finisher.
- 1 Luck Bobblehead | wincon | A designated win condition and marked finisher.
- 1 Sword of Body and Mind | wincon | A designated win condition and marked finisher.
- 1 Ceaseless Conflict | wipe | Provides the deck's designated board wipe.
- 1 Hour of Reckoning | wipe | Provides the deck's designated board wipe.
- 1 Martial Coup | wipe | Provides the deck's designated board wipe.
- 1 Basri Ket | interaction | Provides the deck's designated interaction.
- 1 Lena, Selfless Champion | interaction | Provides the deck's designated interaction.
- 1 Mage's Attendant | interaction | Provides the deck's designated interaction.
- 1 Rootborn Defenses | interaction | Provides the deck's designated interaction.
- 1 Spirit Bonds | interaction | Provides the deck's designated interaction.
- 1 Teyo, the Shieldmage | interaction | Provides the deck's designated interaction.

</details>

### 13. owned first, and the commander is not owned

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 230 names.

Commander: Karlov of the Ghost Council.

Grade: baseline, score 0.31, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $191.77 the whole deck.

**Summary:** Karlov of the Ghost Council leads a white-black lifegain deck that develops through low-cost ramp, lifegain synergy, and a creature board while keeping opposing permanents in check with plentiful removal. Karlov is protected by a deep defensive package, and the deck closes through its finisher-marked threats, with Frodo, Sauron's Bane, Grave Venerations, and Lyra Dawnbringer chief among them. It gives up some speed and consistency for a broad, library-first collection of creatures, equipment, and varied answers.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core Karlov engine needs frequent, repeatable lifegain triggers, and while a handful of payoffs (Angel of Vitality, Rosie Cotton, Light of Promise, Shattered Angel, Exemplar of Light) point that way, a large block of the creature slots (Canyon Crawler, Rabaroo Troop, Reaping Willow, Sneering Shadewriter, Lo and Li, Minwu) are generic bodies unrelated to gaining life, and three board wipes sit awkwardly beside a creature-and-equipment board plan.
- PLAN theme_fit=partly: It is a W/B Karlov deck at a reasonable bracket-3 power level, but the lifegain theme the player asked for is a minority of the list, which reads more like generic removal-plus-equipment goodstuff than a dedicated lifegain build.
- PLAN useful_as_built=yes: 36 lands with even W/B support plus ten cheap ramp pieces, a workable curve, ample card draw, and a full removal suite mean the deck functions and can close games through Karlov, Lyra, Victory's Herald, and Frodo as it stands.
- PLAN summary_honest=partly: It is candid about trading speed and consistency for a library-first pile, and the protection package it advertises is genuinely present, but calling Grave Venerations a chief finisher and leaning on "lifegain synergy" overstates what a thin set of trigger sources actually delivers.
- [WARN] `not_owned`: Karlov of the Ghost Council: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.10 over 63 nonland cards

<details><summary>The deck list</summary>

- 14 Plains | land | Basic white mana for Karlov and the deck’s white spells.
- 14 Swamp | land | Basic black mana for Karlov and the deck’s black spells.
- 1 Command Tower | land | Flexible two-color mana from the shortlist.
- 1 City of Brass | land | Flexible color access for the two-color mana base.
- 1 Exotic Orchard | land | Additional flexible color access.
- 1 Marsh Flats | land | A land slot that supports access to the basic mana base.
- 1 Fabled Passage | land | A land slot that supports access to the basic mana base.
- 1 Evolving Wilds | land | A land slot that supports access to the basic mana base.
- 1 Path of Ancestry | land | A dual-purpose color source for the creature-heavy deck.
- 1 Plaza of Heroes | land | Color access that also suits a commander deck.
- 1 Arcane Signet | ramp | Reliable early ramp for Karlov’s colors.
- 1 Sol Ring | ramp | Efficient acceleration from the shortlist.
- 1 Fellwar Stone | ramp | Early mana acceleration.
- 1 Commander's Sphere | ramp | Color-fixing ramp with late-game utility.
- 1 Wayfarer's Bauble | ramp | Early ramp that supports the basic-land mana base.
- 1 Thought Vessel | ramp | A straightforward mana accelerator.
- 1 Sword of the Animist | ramp | Ramp that also works with the deck’s creature plan.
- 1 Chromatic Lantern | ramp | Ramp and broad color support.
- 1 Relic of Legends | ramp | Mana acceleration that fits a legendary commander deck.
- 1 Springleaf Drum | ramp | Low-cost acceleration for an early creature board.
- 1 Buster Sword | draw | A listed draw tool that also fits the deck’s creature plan.
- 1 Call of the Ring | draw | A dedicated draw option from the shortlist.
- 1 Exemplar of Light | draw | Draw attached to a creature for the board-focused plan.
- 1 Folk Hero | draw | A listed draw engine for the creature suite.
- 1 Idol of Oblivion | draw | A compact draw option.
- 1 Inspiring Overseer | draw | Creature-based card advantage.
- 1 Lembas | draw | A low-cost listed draw option.
- 1 Mask of Memory | draw | Repeatable draw support for attacking creatures.
- 1 Night's Whisper | draw | Efficient card advantage.
- 1 Painful Truths | draw | A dedicated draw spell.
- 1 Skullclamp | draw | A low-cost draw tool for the creature-heavy build.
- 1 Bastion Protector | interaction | Protection-oriented interaction for Karlov.
- 1 Boromir, Warden of the Tower | interaction | A creature-based interaction slot.
- 1 Champion's Helm | interaction | Protection for Karlov or another key legend.
- 1 Clever Concealment | interaction | A flexible defensive interaction spell.
- 1 Darksteel Plate | interaction | A durable protection tool for a key creature.
- 1 Lightning Greaves | interaction | Efficient protection for Karlov and the creature suite.
- 1 Swiftfoot Boots | interaction | Additional protection for important creatures.
- 1 Unbreakable Formation | interaction | A defensive interaction spell for the board.
- 1 Banishing Light | removal | Versatile dedicated removal.
- 1 Bitter Triumph | removal | Efficient single-target removal.
- 1 Crib Swap | removal | Creature-focused removal from the shortlist.
- 1 Dismember | removal | Low-cost creature removal.
- 1 Dispatch | removal | Efficient removal that works with the artifact package.
- 1 Fatal Push | removal | Low-cost single-target removal.
- 1 Generous Gift | removal | Broad permanent removal.
- 1 Get Lost | removal | Flexible removal coverage.
- 1 Infernal Grasp | removal | Straightforward creature removal.
- 1 Austere Command | wipe | A flexible board-reset option.
- 1 Fumigate | wipe | A reset that fits the lifegain plan.
- 1 Vanquish the Horde | wipe | An efficient creature-board reset.
- 1 Aerith Gainsborough | synergy | A listed synergy creature for the lifegain plan.
- 1 Angel of Vitality | synergy | A listed lifegain-synergy creature.
- 1 Compassionate Healer | synergy | A listed synergy piece for the deck’s central plan.
- 1 Kor Firewalker | synergy | A listed synergy creature.
- 1 Light of Promise | synergy | A synergy card that supports Karlov’s lifegain focus.
- 1 Night Nurse, Healer of Heroes | synergy | A listed lifegain-synergy legend.
- 1 Rosie Cotton of South Lane | synergy | A listed synergy creature for the creature-focused plan.
- 1 Angel of Invention | threat | A white creature threat that develops the board.
- 1 Canyon Crawler | threat | A dedicated creature threat from the shortlist.
- 1 Dawnhand Eulogist | threat | A finisher-marked creature threat.
- 1 Gollum, Silent Slinker // Meager Meal | threat | A creature threat with added flexibility.
- 1 Lo and Li, Twin Tutors | threat | A legendary creature threat for the board plan.
- 1 Minwu, White Mage | threat | A creature threat that fits the deck’s white core.
- 1 Rabaroo Troop | threat | A dedicated creature threat from the shortlist.
- 1 Reaping Willow | threat | A dedicated creature threat from the shortlist.
- 1 Rooftop Percher | threat | A dedicated creature threat from the shortlist.
- 1 Shattered Angel | threat | A white creature threat that supports the deck’s plan.
- 1 Sneering Shadewriter | threat | A finisher-marked creature threat.
- 1 Victory's Herald | threat | A substantial white creature threat.
- 1 Frodo, Sauron's Bane | wincon | A finisher-marked win condition.
- 1 Grave Venerations | wincon | A finisher-marked win condition.
- 1 Lyra Dawnbringer | wincon | A finisher-marked creature win condition.

</details>

### 14. the user delegates the commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 286 names.

Commander: Atarka, World Render.

Grade: baseline, score 0.06, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $751.04 to buy, $751.04 the whole deck.

**Summary:** Atarka, World Render leads a Gruul Dragon deck that develops its mana, discounts and supports its tribal creatures, then turns a growing aerial force into decisive combat pressure. It wins through powerful Dragon attacks backed by dedicated finishers, with removal and sweepers clearing the way when combat stalls. The deck gives up some early-game speed for expensive creature payoffs and is most effective when it can keep its Dragon board intact.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every card feeds one plan — Dragon cost reducers, tribal ramp orbs, anthem/trigger payoffs and a critical mass of Dragons that Atarka's double strike converts into damage — with only stray picks like Tibalt's Trickery and the rummaging spells sitting slightly off-axis.
- PLAN theme_fit=yes: It is exactly a Dragons Commander deck with a self-chosen Dragon commander, and its power — no fast mana, no combos, just tribal value and combat — sits comfortably in bracket 3.
- PLAN useful_as_built=yes: 36 lands with an even Gruul base plus roughly ten ramp/cost-reduction pieces supports the heavy top end, and there is plenty of draw, protection and multiple finishing angles to actually close games.
- PLAN summary_honest=yes: The claims of mana development, tribal discounts, aerial beatdown, removal and a couple of sweepers (Breath Weapon, Draconic Intervention) all map to real cards, and it openly admits the slow, expensive curve and board-dependence.
- [INFO] `curve_summary`: average mana value 3.37 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Command Tower | land | Reliable red and green fixing for the Dragon plan.
- 1 Stomping Ground | land | An efficient dual source for the deck’s two colors.
- 1 Taiga | land | An efficient dual source for the deck’s two colors.
- 1 Cinder Glade | land | A red-green land slot that supports both halves of the mana base.
- 1 Copperline Gorge | land | Provides early access to either deck color.
- 1 Cragcrown Pathway // Timbercrown Pathway | land | Flexible access to the color needed in the moment.
- 1 Karplusan Forest | land | An untapped source for either deck color.
- 1 Rootbound Crag | land | Supports the red-green dual-land core.
- 1 Rockfall Vale | land | Provides both colors in a single land slot.
- 1 Spire Garden | land | An efficient red-green source for multiplayer games.
- 1 Game Trail | land | Additional red-green fixing for the Dragon curve.
- 1 Fire-Lit Thicket | land | Helps convert mana into the required colors.
- 1 Thornspire Verge | land | A dual-color land for the Gruul mana base.
- 1 Wooded Foothills | land | Flexible access to the deck’s basic land types.
- 11 Forest | land | Provides dependable green mana.
- 11 Mountain | land | Provides dependable red mana.
- 1 Faithless Looting | draw | Low-cost card selection to keep plays flowing.
- 1 Demand Answers | draw | Provides an efficient draw option.
- 1 Thrill of Possibility | draw | Helps dig toward lands, Dragons, and action.
- 1 Beast Whisperer | draw | Creature-focused draw support for a creature-heavy deck.
- 1 Garruk's Uprising | draw | Draw support that suits the deck’s large creatures.
- 1 Guardian Project | draw | Sustained draw for deploying a stream of creatures.
- 1 Harmonize | draw | Straightforward card replenishment.
- 1 Return of the Wildspeaker | draw | Flexible draw that works with the deck’s large threats.
- 1 Dragonborn Champion | draw | A Dragon-themed source of ongoing card advantage.
- 1 Vanquisher's Banner | draw | Tribal card advantage for the Dragon roster.
- 1 Sylvan Library | draw | Early card selection and sustained draw support.
- 1 Heroic Intervention | interaction | Protects the board from opposing interaction.
- 1 Lightning Greaves | interaction | Protects an important Dragon or the commander.
- 1 Swiftfoot Boots | interaction | Provides repeatable protection for key creatures.
- 1 Mithril Coat | interaction | A protection piece for the deck’s major threats.
- 1 Veil of Summer | interaction | A low-cost answer to opposing disruption.
- 1 Tibalt's Trickery | interaction | Broad stack interaction in red.
- 1 Legolas's Quick Reflexes | interaction | A flexible protective response for a creature.
- 1 Steely Resolve | interaction | Tribal protection for the Dragon force.
- 1 Mox Jasper | ramp | A low-cost Dragon-themed mana source.
- 1 Jade Orb of Dragonkind | ramp | Early mana development that supports Dragons.
- 1 Orb of Dragonkind | ramp | Color fixing and mana development for the deck.
- 1 Pillar of Origins | ramp | Tribal mana fixing for casting Dragons.
- 1 Scaled Nurturer | ramp | Early creature-based mana acceleration.
- 1 Carnelian Orb of Dragonkind | ramp | Red mana acceleration tailored to Dragons.
- 1 Dragon's Hoard | ramp | Dragon-themed mana development.
- 1 Progenitor's Icon | ramp | Supports the tribal creature plan with mana.
- 1 Patchwork Banner | ramp | A tribal mana piece that develops the board.
- 1 Ganax, Astral Hunter | ramp | A Dragon that contributes to mana development.
- 1 Draconic Roar | removal | Efficient Dragon-themed targeted removal.
- 1 Dragon's Fire | removal | Low-cost removal that fits the deck’s theme.
- 1 Invasion of Tarkir // Defiant Thundermaw | removal | Dragon-focused removal with a Dragon payoff.
- 1 Piercing Exhale | removal | An efficient removal spell for opposing creatures.
- 1 Spit Flame | removal | Repeatable Dragon-oriented creature removal.
- 1 Dragon Tempest | removal | A tribal piece that also supplies removal pressure.
- 1 Akoum Hellkite | removal | A Dragon body that contributes targeted removal.
- 1 Glorybringer | removal | A threatening Dragon that answers creatures.
- 1 Terror of the Peaks | removal | A finisher-class Dragon that supplies removal pressure.
- 1 Dragonlord's Servant | synergy | Reduces the burden of deploying expensive Dragons.
- 1 Dragonspeaker Shaman | synergy | Supports the deck’s Dragon-heavy casting plan.
- 1 Herald's Horn | synergy | A tribal value piece for the chosen creature type.
- 1 Urza's Incubator | synergy | Makes the Dragon curve easier to deploy.
- 1 Crucible of Fire | synergy | A dedicated boost for the Dragon combat plan.
- 1 Shared Animosity | synergy | Rewards attacking with a wide Dragon force.
- 1 Molten Echoes | synergy | A tribal engine that multiplies Dragon value.
- 1 Dragonloft Idol | threat | A low-cost threat that benefits from the Dragon shell.
- 1 Manaform Hellkite | threat | An efficient Dragon threat for the combat plan.
- 1 Roaming Throne | threat | A resilient tribal threat that supports the chosen type.
- 1 Territorial Hellkite | threat | An aggressive Dragon threat for pressuring opponents.
- 1 Backdraft Hellkite | threat | A Dragon threat that adds value to the attack plan.
- 1 Mirrorwing Dragon | threat | A Dragon threat that makes opposing targeting awkward.
- 1 Twinflame Tyrant | threat | A finisher-class Dragon threat for combat turns.
- 1 Stormbreath Dragon | threat | A hasty Dragon threat that applies immediate pressure.
- 1 Verix Bladewing | threat | A Dragon threat that contributes to board presence.
- 1 Hoarding Dragon | threat | A substantial Dragon threat for the midgame.
- 1 Red Dragon | threat | A straightforward large Dragon threat.
- 1 Thundermane Dragon | threat | A Dragon threat that reinforces the creature plan.
- 1 Hellkite Tyrant | wincon | A finisher-class Dragon with a dedicated winning role.
- 1 Lathliss, Dragon Queen | wincon | A finisher-class Dragon that serves as a primary payoff.
- 1 Wrathful Red Dragon | wincon | A finisher-class Dragon with a dedicated winning role.
- 1 Breath Weapon | wipe | A tribal-friendly board-clearing option.
- 1 Draconic Intervention | wipe | A flexible sweeper for resetting opposing boards.
- 1 Thundermaw Hellkite | wipe | A finisher-class Dragon that can clear smaller creatures.

</details>

### 15. delegated commander, owned first

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 230 names.

Commander: Astarion, the Decadent.

Grade: baseline, score 0.28, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $212.07 the whole deck.

**Summary:** Astarion, the Decadent leads a white-black lifegain deck that develops its mana, establishes lifegain-focused creatures and Equipment, then turns that board into steady combat pressure. The deck closes through its Angel and legendary finishers, with several dedicated endgame cards chief among them. It gives up explosive speed for a broad mix of creature pressure, protective interaction, targeted answers, and board resets.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: The pieces point one way: lifegain enablers and payoffs (Angel of Vitality, Shattered Angel, Rosie Cotton, Light of Promise, Compassionate Healer, Lyra) backed by equipment/protection for the commander, ramp, and removal, with the wipes fitting Astarion's reward for creatures dying rather than fighting the plan.
- PLAN theme_fit=yes: It is a Commander lifegain deck with a builder-chosen W/B lifegain commander at roughly precon/bracket-2 power, matching the request; the only unverifiable part is whether the cards came from the person's library.
- PLAN useful_as_built=yes: 37 lands with an even W/B base plus fixing, nine ramp pieces, a reasonable curve, ample removal and multiple combat and Angel finishers means it can be shuffled up and played as printed.
- PLAN summary_honest=yes: It accurately describes a W/B lifegain build with equipment, protection, spot removal and three board resets, and the Angel/legendary finishers it names are present; "several dedicated endgame cards" is a mild stretch for three tagged wincons but nothing is claimed that the list lacks or hidden.
- [WARN] `not_owned`: Astarion, the Decadent: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.18 over 62 nonland cards

<details><summary>The deck list</summary>

- 1 Buster Sword | draw | Draw support from the shortlist.
- 1 Call of the Ring | draw | Draw support from the shortlist.
- 1 Cirith Ungol Patrol | draw | Draw support from the shortlist.
- 1 Exemplar of Light | draw | Draw support that also fits the deck's creature theme.
- 1 Hero in Training | draw | Draw support from the shortlist.
- 1 Idol of Oblivion | draw | A colorless draw option.
- 1 Inspiring Overseer | draw | Draw support that fits the lifegain-focused creature base.
- 1 Lembas | draw | A colorless draw option.
- 1 Mask of Memory | draw | Draw support for the deck's creatures.
- 1 Night's Whisper | draw | Efficient black draw support.
- 1 Puresteel Paladin | draw | Draw support that pairs with the deck's Equipment.
- 1 Bastion Protector | interaction | Interaction that helps protect the commander.
- 1 Boromir, Warden of the Tower | interaction | Creature-based interaction.
- 1 Clever Concealment | interaction | Broad protective interaction.
- 1 Darksteel Plate | interaction | Protective Equipment for a key creature.
- 1 Lightning Greaves | interaction | Efficient protection for the commander or a finisher.
- 1 Swiftfoot Boots | interaction | Additional protective Equipment.
- 1 Banishing Light | removal | Flexible white removal.
- 1 Bitter Triumph | removal | Efficient black removal.
- 1 Crib Swap | removal | Creature removal from the shortlist.
- 1 Dispatch | removal | Low-cost removal.
- 1 Fatal Push | removal | Efficient black removal.
- 1 Generous Gift | removal | Flexible instant-speed removal.
- 1 Get Lost | removal | Flexible white removal.
- 1 Austere Command | wipe | Flexible reset button.
- 1 Dusk // Dawn | wipe | Board reset that fits the creature-heavy plan.
- 1 Fumigate | wipe | A wipe that belongs in the lifegain plan.
- 1 Arcane Signet | ramp | Reliable two-color acceleration.
- 1 Chromatic Lantern | ramp | Mana fixing and acceleration.
- 1 Commander's Sphere | ramp | Color fixing and ramp.
- 1 Fellwar Stone | ramp | Low-cost mana acceleration.
- 1 Lotho, Corrupt Shirriff | ramp | Creature-based ramp.
- 1 Relic of Legends | ramp | Mana acceleration for the legendary creature base.
- 1 Sol Ring | ramp | Fast colorless acceleration.
- 1 Sword of the Animist | ramp | Equipment-based mana development.
- 1 Thought Vessel | ramp | Colorless mana acceleration.
- 1 Wayfarer's Bauble | ramp | Early mana development.
- 1 Aerith Gainsborough | synergy | A lifegain-focused synergy creature.
- 1 Aettir and Priwen | synergy | Equipment synergy for the creature plan.
- 1 Angel of Vitality | synergy | A lifegain-focused synergy creature.
- 1 Compassionate Healer | synergy | A lifegain-focused synergy creature.
- 1 Crowd of True Believers | synergy | Creature-based synergy for the deck's plan.
- 1 Elixir | synergy | A colorless lifegain synergy piece.
- 1 Kor Firewalker | synergy | A lifegain-focused synergy creature.
- 1 Light of Promise | synergy | A payoff for the lifegain plan.
- 1 Night Nurse, Healer of Heroes | synergy | A lifegain-focused synergy creature.
- 1 Rosie Cotton of South Lane | synergy | A creature-based lifegain payoff.
- 1 Angel of Invention | threat | An evasive creature threat.
- 1 Dawnhand Eulogist | threat | A resilient finisher-class creature threat.
- 1 Foggy Swamp Hunters | threat | Creature pressure for the midgame.
- 1 Gollum, Silent Slinker // Meager Meal | threat | A black creature threat.
- 1 Lo and Li, Twin Tutors | threat | A legendary creature threat.
- 1 Minwu, White Mage | threat | A white creature threat.
- 1 Rabaroo Troop | threat | Creature pressure for the board.
- 1 Rooftop Percher | threat | A creature threat from the shortlist.
- 1 Shattered Angel | threat | An evasive creature threat.
- 1 Sneering Shadewriter | threat | A black finisher-class creature threat.
- 1 Victory's Herald | threat | An evasive creature threat.
- 1 Witch-king of Angmar | threat | A finisher-class creature threat.
- 1 Frodo, Sauron's Bane | wincon | A dedicated finisher for closing games.
- 1 Grave Venerations | wincon | A dedicated finisher for the deck's plan.
- 1 Lyra Dawnbringer | wincon | An evasive finisher that fits the lifegain theme.
- 1 Plains | land | Basic white mana source.
- 1 Plains | land | Basic white mana source.
- 1 Plains | land | Basic white mana source.
- 1 Plains | land | Basic white mana source.
- 1 Plains | land | Basic white mana source.
- 1 Plains | land | Basic white mana source.
- 1 Plains | land | Basic white mana source.
- 1 Plains | land | Basic white mana source.
- 1 Plains | land | Basic white mana source.
- 1 Plains | land | Basic white mana source.
- 1 Plains | land | Basic white mana source.
- 1 Plains | land | Basic white mana source.
- 1 Plains | land | Basic white mana source.
- 1 Swamp | land | Basic black mana source.
- 1 Swamp | land | Basic black mana source.
- 1 Swamp | land | Basic black mana source.
- 1 Swamp | land | Basic black mana source.
- 1 Swamp | land | Basic black mana source.
- 1 Swamp | land | Basic black mana source.
- 1 Swamp | land | Basic black mana source.
- 1 Swamp | land | Basic black mana source.
- 1 Swamp | land | Basic black mana source.
- 1 Swamp | land | Basic black mana source.
- 1 Swamp | land | Basic black mana source.
- 1 Swamp | land | Basic black mana source.
- 1 Swamp | land | Basic black mana source.
- 1 City of Brass | land | Flexible color fixing.
- 1 Command Tower | land | Reliable commander-color fixing.
- 1 Evolving Wilds | land | Fixes either basic color.
- 1 Exotic Orchard | land | Flexible color fixing.
- 1 Fabled Passage | land | Fixes either basic color.
- 1 Grand Coliseum | land | Flexible color fixing.
- 1 Marsh Flats | land | Finds a needed basic land.
- 1 Path of Ancestry | land | A commander-color land for the creature deck.
- 1 Plaza of Heroes | land | Supports the legendary creature package.
- 1 Spire of Industry | land | Color fixing for the artifact package.
- 1 Terramorphic Expanse | land | Fixes either basic color.

</details>

### 16. a tight budget, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 231 names.

Commander: Denethor, Ruling Steward.

Grade: baseline, score 0.31, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $14.01 to buy, $57.45 the whole deck.

**Summary:** Denethor leads a creature-focused Orzhov aristocrats deck that develops a board, protects its key legends, and uses removal and board resets to keep opponents from stabilizing. It closes through its dedicated finisher package, chief among them Bastion of Remembrance, Falkenrath Noble, and Grave Venerations. The deck gives up premium mana and expensive staples for a mostly basic, budget-conscious mana base and a straightforward creature-driven game plan.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The deck makes less mana on turn four than the norm, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core is a real aristocrats shell (sac outlets like Bartolom	tdel Presidio, Woe Strider, Ayli, Ghoulcaller Gisa paired with drains like Bastion of Remembrance and Falkenrath Noble), but a legends-matter/protection subtheme (Bastion Protector, Plaza of Heroes, Relic of Legends, tribal-only lands like Unclaimed Territory and Secluded Courtyard that do nothing in a non-tribal build) and three symmetrical wraths pull against a go-wide sacrifice board.
- PLAN theme_fit=yes: It is an Orzhov aristocrats Commander deck under a legal aristocrats-flavored commander at a modest, precon-level power consistent with bracket 2, matching what was asked.
- PLAN useful_as_built=yes: 37 lands with fixing, ten mana rocks, a sensible curve, real card draw, ample removal and multiple drain win conditions mean it can be picked up and played, even if the tribal lands and self-damaging wraths cost it some smoothness.
- PLAN summary_honest=partly: The summary accurately names the finishers, protection, removal and wipes that are in the list, but it frames the deck as a creature-driven aristocrats plan without acknowledging that the wipes fight that plan, and its 'dedicated finisher package' is only three drain cards.
- [WARN] `not_owned`: Ayli, Eternal Pilgrim: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bartolomé del Presidio: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Aron, Benalia's Ruin: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Baron Bertram Graywater: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Erebos, Bleak-Hearted: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ghoulcaller Gisa: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Lord Skitter's Butcher: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Old Flitterfang: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ruthless Lawbringer: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ruthless Technomancer: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Shadowheart, Dark Justiciar: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Sivriss, Nightmare Speaker: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Skullport Merchant: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Smothering Abomination: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Woe Strider: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bastion of Remembrance: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Falkenrath Noble: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.00 over 62 nonland cards
- [INFO] `basics_added`: the list was 1 card short, so the builder added 1 basic land

<details><summary>The deck list</summary>

- 12 Plains | land | Untapped white mana for the deck's white spells.
- 11 Swamp | land | Untapped black mana for the deck's black spells.
- 1 Command Tower | land | A reliable multicolor land for Denethor's deck.
- 1 Exotic Orchard | land | A multicolor land that supports both halves of the deck.
- 1 City of Brass | land | An untapped multicolor source.
- 1 Grand Coliseum | land | A multicolor land slot for the mana base.
- 1 Evolving Wilds | land | A fixing land that finds a needed basic.
- 1 Fabled Passage | land | A fixing land that finds a needed basic.
- 1 Terramorphic Expanse | land | A fixing land that finds a needed basic.
- 1 Path of Ancestry | land | A multicolor land for the creature-heavy build.
- 1 Secluded Courtyard | land | A flexible colored source for the creature suite.
- 1 Unclaimed Territory | land | A flexible colored source for the creature suite.
- 1 Spire of Industry | land | An untapped colored source alongside the artifact package.
- 1 Thriving Moor | land | A black source with additional color fixing.
- 1 Plaza of Heroes | land | A land slot that supports the legendary-heavy creature base.
- 1 Great Hall of the Citadel | land | A mana source for the deck's legendary cards.
- 1 Sol Ring | ramp | Efficient early acceleration.
- 1 Arcane Signet | ramp | Reliable two-color acceleration.
- 1 Fellwar Stone | ramp | Low-cost color fixing and acceleration.
- 1 Thought Vessel | ramp | A cheap mana rock for developing the board.
- 1 Blitzball | ramp | A low-cost ramp piece from the library.
- 1 White Auracite | ramp | A cheap ramp artifact.
- 1 Inherited Envelope | ramp | A ramp card that fits the budget.
- 1 Ring of the Lucii | ramp | A mana artifact for the midgame.
- 1 Relic of Legends | ramp | Color fixing and acceleration for the legendary suite.
- 1 Commander's Sphere | ramp | Color fixing that remains useful later.
- 1 Skullclamp | draw | A premier draw card for the aristocrats plan.
- 1 Wall of Omens | draw | An early creature that supplies a draw card.
- 1 Exemplar of Light | draw | A creature draw option from the library.
- 1 Inspiring Overseer | draw | A creature that contributes to the draw package.
- 1 Nasty End | draw | An efficient black draw spell.
- 1 Painful Truths | draw | A compact draw spell for a two-color deck.
- 1 Mask of Memory | draw | An artifact draw option for the creature base.
- 1 Tome of Legends | draw | A low-cost source of continued cards.
- 1 Idol of Oblivion | draw | An artifact draw piece for the deck's plan.
- 1 Stone of Erech | draw | A low-cost draw artifact.
- 1 Rat King, Pale Piper | draw | A creature draw option that fits the black creature base.
- 1 Bastion Protector | interaction | Protection for Denethor and the important legendary creatures.
- 1 Gift of Immortality | interaction | A protective piece for a key creature.
- 1 Swiftfoot Boots | interaction | Protection for Denethor or a major threat.
- 1 Take Up the Shield | interaction | A low-cost protective instant.
- 1 Together Forever | interaction | A resilient interaction piece for the creature board.
- 1 Ultimate Magic: Holy | interaction | An instant-speed interaction card from the library.
- 1 Bitter Triumph | removal | Flexible creature or planeswalker removal.
- 1 Claim the Precious | removal | A direct black removal spell.
- 1 Cornered by Black Mages | removal | An efficient removal card from the library.
- 1 Crib Swap | removal | Creature removal that fits the color identity.
- 1 Destroy Evil | removal | Flexible white removal.
- 1 Fatal Push | removal | Cheap targeted removal.
- 1 Infernal Grasp | removal | Unconditional creature removal.
- 1 Austere Command | wipe | A flexible board-reset option.
- 1 Dusk // Dawn | wipe | A creature-focused reset that suits the creature plan.
- 1 Fumigate | wipe | A dependable full-board reset.
- 1 Arcade Cabinet | synergy | Included in the aristocrats synergy package.
- 1 Gollum the Abandoned | synergy | Included in the aristocrats synergy package.
- 1 Gollum, Patient Plotter | synergy | Included in the aristocrats synergy package.
- 1 Gríma Wormtongue | synergy | Included in the aristocrats synergy package.
- 1 Joo Dee, One of Many | synergy | Included in the aristocrats synergy package.
- 1 Nimble Hobbit | synergy | Included in the aristocrats synergy package.
- 1 Phantom Train | synergy | An artifact included in the aristocrats synergy package.
- 1 Ayli, Eternal Pilgrim | synergy | A low-cost legendary creature for the aristocrats synergy package.
- 1 Bartolomé del Presidio | synergy | A low-cost creature for the aristocrats synergy package.
- 1 Aron, Benalia's Ruin | synergy | A legendary creature included for the aristocrats plan.
- 1 Baron Bertram Graywater | threat | A creature threat that fits the black creature core.
- 1 Erebos, Bleak-Hearted | threat | A durable black threat for longer games.
- 1 Ghoulcaller Gisa | threat | A high-impact black creature threat.
- 1 Lord Skitter's Butcher | threat | A black creature threat for the board.
- 1 Old Flitterfang | threat | A creature threat that supports the deck's overall plan.
- 1 Ruthless Lawbringer | threat | A low-cost black creature threat.
- 1 Ruthless Technomancer | threat | A substantial creature threat for the midgame.
- 1 Shadowheart, Dark Justiciar | threat | A black legendary threat for the creature suite.
- 1 Sivriss, Nightmare Speaker | threat | A low-cost black threat from the library.
- 1 Skullport Merchant | threat | A creature threat that fits the aristocrats shell.
- 1 Smothering Abomination | threat | A black creature threat for the sacrifice-focused board.
- 1 Woe Strider | threat | A creature threat that fits the aristocrats shell.
- 1 Bastion of Remembrance | wincon | A dedicated aristocrats finisher.
- 1 Falkenrath Noble | wincon | A creature finisher for the aristocrats plan.
- 1 Grave Venerations | wincon | A dedicated finisher already in the library.

</details>

### 17. upgrade a precon, owned first

Format: Commander. Theme: goblins. Pool: owned_first. Shortlist: 204 names.

Commander: Zada, Hedron Grinder.

Grade: typical, score 0.55, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $34.73 to buy, $212.28 the whole deck.

**Summary:** Zada leads a red Goblin Storm deck that develops a Goblin board, converts mana bursts and inexpensive card flow into productive turns, and stacks Goblin synergies to make its creatures increasingly dangerous. Great Train Heist and Assault on Osgiliath are chief among the finishing plays, with Earthquake providing another route to end a crowded game. The deck gives up broad color access and favors committed battlefield development, so its strongest turns come from protecting a built-up board and choosing the right window for a decisive spell-heavy turn.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The cards pair the way the top lists pair them, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every card supports one plan — flood the board with Goblin tokens, use Zada to multiply cheap cantrips/pumps, and cash in with rituals, Impact Tremors/Bombardment and extra-combat or mass-pump finishers — with only minor tension between the wipes (Blasphemous Act/Earthquake) and the wide board.
- PLAN theme_fit=yes: It is a mono-red Zada Goblin storm/token deck with rituals, Grapeshot, Empty the Warrens and Past in Flames at a reasonable mid-power bracket-3 level, matching the requested precon upgrade.
- PLAN useful_as_built=yes: Roughly 37 lands mostly Mountains plus Sol Ring, Arcane Signet and Goblin rituals support a low curve, and the deck has multiple win routes (Bombardment/Impact Tremors drain, Shared Animosity/Great Train Heist alpha strikes, Grapeshot), so it plays fine as written.
- PLAN summary_honest=yes: The claims match the list: Great Train Heist, Assault on Osgiliath and Earthquake are present, rituals and cheap cantrips are numerous, and it fairly admits the mono-red limitation without asserting tutors or combos the deck lacks.
- [INFO] `curve_summary`: average mana value 2.54 over 61 nonland cards
- [INFO] `precon_cards_restored`: the deck was 1 card short of the Goblin Storm precon share, so the builder put 1 card back
- [WARN] `profile_off_band`: the count of nonbasic lands that make only colorless mana is 5, and bracket 3 wants 4 at most (Fountainport, Goblin Burrows, Kher Keep, Reliquary Tower, War Room)
- [WARN] `profile_off_band`: the interaction count is 1, and bracket 3 wants 4 to 12
- [INFO] `bracket_cut`: to hold bracket 3, the builder cut 1 card: Storm-Kiln Artist (the combo Storm-Kiln Artist + Haze of Rage), and added 1 basic land

<details><summary>The deck list</summary>

- 1 Ancestors' Aid | draw | Retained from the precon as a draw spell for the storm-oriented game plan.
- 1 Arena of Glory | land | Retained from the precon mana base.
- 1 Battle Hymn | ramp | Retained from the precon as Goblin-based mana acceleration.
- 1 Blasphemous Act | wipe | Retained from the precon as a board reset.
- 1 Boggart Shenanigans | removal | Retained from the precon as Goblin-linked removal.
- 1 Brightstone Ritual | ramp | Retained from the precon as Goblin-based mana acceleration.
- 1 Castle Embereth | land | Retained from the precon mana base.
- 1 Chaos Warp | removal | Retained from the precon as flexible removal.
- 1 Conspicuous Snoop | synergy | Retained from the precon for Goblin synergy.
- 1 Crimson Wisps | draw | Retained from the precon as inexpensive card flow.
- 1 Daring Discovery | other | Retained from the precon to preserve its original package.
- 1 Den of the Bugbear | land | Retained from the precon mana base.
- 1 Dragon Fodder | synergy | Retained from the precon to support the Goblin-focused plan.
- 1 Empty the Warrens | synergy | Retained from the precon to support the storm-focused plan.
- 1 Expedite | draw | Retained from the precon as inexpensive card flow.
- 1 Faithless Looting | draw | Retained from the precon as card selection and draw.
- 1 Fists of Flame | draw | Retained from the precon as a draw spell.
- 1 Forgotten Cave | land | Retained from the precon mana base.
- 1 Fountainport | land | Retained from the precon mana base.
- 1 Frontline Heroism | other | Retained from the precon to preserve its original package.
- 1 Gempalm Incinerator | removal | Retained from the precon as Goblin-linked removal.
- 1 Great Train Heist | wincon | Retained from the precon as a finishing win condition.
- 1 Glimpse the Impossible | ramp | Retained from the precon as mana acceleration.
- 1 Goblin Bombardment | removal | Retained from the precon as a Goblin-linked removal outlet.
- 1 Goblin Burrows | land | Retained from the precon mana base.
- 1 Goblin Bushwhacker | synergy | Retained from the precon for Goblin synergy.
- 1 Goblin Chieftain | synergy | Retained from the precon for Goblin synergy.
- 1 Goblin Dark-Dwellers | threat | Retained from the precon as a Goblin threat.
- 1 Goblin Lackey | synergy | Retained from the precon for Goblin synergy.
- 1 Goblin Matron | synergy | Retained from the precon for Goblin synergy.
- 1 Goblin Negotiation | removal | Retained from the precon as removal.
- 1 Goblin Trashmaster | removal | Retained from the precon as Goblin-linked removal.
- 1 Goblin Warchief | synergy | Retained from the precon for Goblin synergy.
- 1 Grapeshot | removal | Retained from the precon as storm-oriented removal.
- 1 Haze of Rage | synergy | Retained from the precon to support the storm-focused plan.
- 1 Hidden Volcano | land | Retained from the precon mana base.
- 1 Howlsquad Heavy | ramp | Retained from the precon as mana acceleration.
- 1 Idol of Oblivion | draw | Retained from the precon as repeatable card flow.
- 1 Impact Tremors | synergy | Retained from the precon to support the token-heavy Goblin plan.
- 1 Kher Keep | land | Retained from the precon mana base.
- 1 Krenko's Command | synergy | Retained from the precon to support the Goblin-focused plan.
- 1 Mana Geyser | ramp | Retained from the precon as high-output mana acceleration.
- 1 Mogg War Marshal | synergy | Retained from the precon for Goblin synergy.
- 27 Mountain | land | Keeps the mana base consistent and supports the precon's land count.
- 1 Past in Flames | other | Retained from the precon to preserve its storm-oriented package.
- 1 Quest for the Goblin Lord | synergy | Retained from the precon for Goblin synergy.
- 1 Reliquary Tower | land | Retained from the precon mana base.
- 1 Ruby Medallion | ramp | Retained from the precon as mana acceleration.
- 1 Sazacap's Brew | draw | Retained from the precon as a draw spell.
- 1 Searslicer Goblin | synergy | Retained from the precon for Goblin synergy.
- 1 Seething Song | ramp | Retained from the precon as mana acceleration.
- 1 Siege-Gang Commander | removal | Retained from the precon as Goblin-linked removal.
- 1 Siege-Gang Lieutenant | removal | Retained from the precon as Goblin-linked removal.
- 1 Skirk Prospector | ramp | Retained from the precon as Goblin-based mana acceleration.
- 1 Skullclamp | draw | Retained from the precon as card draw.
- 1 Sol Ring | ramp | Retained from the precon as efficient mana acceleration.
- 1 Spreading Insurrection | other | Retained from the precon to preserve its original package.
- 1 Swiftfoot Boots | interaction | Retained from the precon as protection.
- 1 Throne of Eldraine | ramp | Retained from the precon as mana acceleration.
- 1 Vandalblast | wipe | Retained from the precon as a board wipe.
- 1 Warren Torchmaster | synergy | Retained from the precon for Goblin synergy.
- 1 War Room | land | Retained from the precon mana base.
- 1 Wild Ride | other | Retained from the precon to preserve its original package.
- 1 Witch's Mark | draw | Retained from the precon as card draw.
- 1 Shared Animosity | synergy | Adds a potent Goblin synergy piece without changing the deck's core plan.
- 1 Rundvelt Hordemaster | synergy | Adds another efficient Goblin synergy card.
- 1 Pashalik Mons | removal | Adds Goblin-linked removal to strengthen the existing creature package.
- 1 Fable of the Mirror-Breaker // Reflection of Kiki-Jiki | ramp | Adds flexible mana acceleration while keeping the creature-and-spell plan intact.
- 1 Arcane Signet | ramp | Adds reliable early mana acceleration.
- 1 Abrade | removal | Adds efficient, flexible removal.
- 1 Assault on Osgiliath | wincon | Adds a second dedicated finishing win condition.
- 1 Earthquake | wipe | Adds a scalable board wipe that can also close games.
- 1 Shinka, the Bloodsoaked Keep | synergy | the deck upgrades a precon, and this card is one the precon holds

</details>

### 18. upgrade a precon, any card

Format: Commander. Theme: turtles. Pool: any_card. Shortlist: 364 names.

Commander: Heroes in a Half Shell.

Grade: baseline, score 0.07, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $257.12 to buy, $257.12 the whole deck.

**Summary:** This Turtle Power build keeps the precon’s broad Turtle and character core, using its tribal pieces and creature threats to establish a battlefield before closing with Dimension X Pizzasaur, Everything Pizza, or Raphael, the Muscle. Rhystic Study supplies steadier cards while Mana Confluence and Cavern of Souls tighten the mana without changing the deck’s identity. It remains a thematic five-color creature deck, so it gives up some streamlined speed for its wide cast of support cards and character-driven play.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: The list is one consistent five-color Turtle/character tribal build with a mana base full of fixing, tribal payoffs, team-up character cards, and a handful of big finishers, so the pieces all pull in the same precon direction the summary describes.
- PLAN theme_fit=partly: It is the right commander, format, and Turtle theme, but the request was to upgrade the precon to bracket 3 and only three cards appear to have been changed, which is a very thin answer to 'upgrade.'
- PLAN useful_as_built=yes: With roughly forty lands heavy in five-color fixing, a normal precon curve, ramp, removal, wipes, and several creature finishers, the deck can be shuffled up and played as it stands.
- PLAN summary_honest=yes: The summary accurately says it keeps the precon core, names only the three actual additions (Rhystic Study, Mana Confluence, Cavern of Souls), and openly concedes the deck trades speed for thematic breadth rather than claiming a plan the cards lack.
- [INFO] `curve_summary`: average mana value 3.48 over 60 nonland cards
- [WARN] `profile_off_band`: the land count is 39, and bracket 3 wants 34 to 38
- [WARN] `profile_off_band`: the ramp count is 5, and bracket 3 wants 8 to 13
- [WARN] `profile_off_band`: the draw count is 3, and bracket 3 wants 8 to 14
- [WARN] `profile_off_band`: the removal count is 2, and bracket 3 wants 6 to 12
- [WARN] `profile_off_band`: the interaction count is 0, and bracket 3 wants 4 to 12

<details><summary>The deck list</summary>

- 1 Acidic Slime | other | Retained precon creature support.
- 1 April O'Neil, Live on the Scene | draw | Retained Turtle Power card-draw support.
- 1 Arcade Cabinet | other | Retained precon support piece.
- 1 Arcane Signet | ramp | Retained precon mana acceleration.
- 1 Ash Barrens | land | Retained precon mana fixing land.
- 1 Assassin's Trophy | removal | Retained flexible precon removal.
- 1 Baxter, Fly in the Ointment | other | Retained precon character card.
- 1 Bebop, Skull & Crossbones | other | Retained precon character card.
- 1 Big Apple, 3 a.m. | land | Retained precon mana land.
- 1 Big Mother Mouser | other | Retained precon artifact support.
- 1 Biogenic Ooze | threat | Retained precon creature threat.
- 1 Blasphemous Act | wipe | Retained precon board reset.
- 1 Chromatic Lantern | ramp | Retained precon mana fixing support.
- 1 Cinder Glade | land | Retained precon dual land.
- 1 City of Brass | land | Retained precon five-color source.
- 1 Coin of Mastery | other | Retained precon support artifact.
- 1 Command Tower | land | Retained precon five-color source.
- 1 Continue? | other | Retained precon support spell.
- 1 Corpsejack Menace | other | Retained precon creature support.
- 1 Cultivate | ramp | Retained precon mana development.
- 1 Dimension X Pizzasaur | wincon | Retained precon finishing threat.
- 1 Donatello, the Brains | synergy | Retained Turtle Power synergy piece.
- 1 Double Jump // Flying Kick | other | Retained precon support spell.
- 1 Dragonskull Summit | land | Retained precon dual land.
- 1 Electric Seaweed | other | Retained precon creature support.
- 1 Elvish Soultiller | threat | Retained precon creature threat.
- 1 Endless Foot Assault | other | Retained precon support enchantment.
- 1 Escape Tunnel | land | Retained precon mana land.
- 1 Everything Pizza | wincon | Retained precon finishing card.
- 1 Evolving Wilds | land | Retained precon fixing land.
- 1 Exploding Barrel | other | Retained precon artifact support.
- 1 Fabled Passage | land | Retained precon fixing land.
- 1 Fast Forward | other | Retained precon support spell.
- 1 Foot Chopper | other | Retained precon Equipment support.
- 1 From the Rubble | synergy | Retained Turtle Power synergy card.
- 1 Game Over | other | Retained precon support spell.
- 1 Grand Coliseum | land | Retained precon five-color source.
- 1 Harmonize | draw | Retained precon card draw.
- 1 Here Comes a New Hero! | other | Retained precon support spell.
- 1 Hidden Hideout | land | Retained precon mana land.
- 1 High Score | other | Retained precon support enchantment.
- 1 Hinterland Harbor | land | Retained precon dual land.
- 1 Irma, Part-Time Mutant | other | Retained precon character card.
- 1 Krang, the All-Powerful | other | Retained precon artifact creature support.
- 1 Leatherhead, Iron Gator | other | Retained precon character card.
- 1 Leonardo, the Balance | threat | Retained precon Turtle threat.
- 1 Lessons from Life | other | Retained precon support spell.
- 1 Level Up | other | Retained precon Aura support.
- 1 Lita, Little Orphan Amphibian | synergy | Retained Turtle Power synergy piece.
- 1 Michelangelo, the Heart | synergy | Retained Turtle Power synergy piece.
- 1 Mole Module | other | Retained precon Vehicle support.
- 1 Mona Lisa, Science Geek | other | Retained precon character card.
- 1 Ninja Pizza | other | Retained precon enchantment support.
- 1 Path of Ancestry | land | Retained precon tribal mana land.
- 1 Rain-Slicked Copse | land | Retained precon dual land.
- 1 Rat King, Pale Piper | other | Retained precon character card.
- 1 Ray Fillet, Wave Warrior | other | Retained precon character card.
- 1 Roadkill Rodney | other | Retained precon artifact creature support.
- 1 Rocksteady, Mutant Marauder | other | Retained precon character card.
- 1 Rootbound Crag | land | Retained precon dual land.
- 1 Shellshock | other | Retained precon support spell.
- 1 Shredder, Shadow Master | other | Retained precon character card.
- 1 Smoldering Marsh | land | Retained precon dual land.
- 1 Sodden Verdure | land | Retained precon dual land.
- 1 Sol Ring | ramp | Retained precon mana acceleration.
- 1 Special Move | other | Retained precon support spell.
- 1 Spire Garden | land | Retained precon dual land.
- 1 Steelbane Hydra | removal | Retained precon removal creature.
- 1 Sunken Hollow | land | Retained precon dual land.
- 1 Super Combo | other | Retained precon support spell.
- 1 Swift Demise | other | Retained precon support spell.
- 1 Tempestra, Dame of Games | other | Retained precon character card.
- 1 Thriving Grove | land | Retained precon fixing land.
- 1 Thriving Isle | land | Retained precon fixing land.
- 1 Thriving Moor | land | Retained precon fixing land.
- 1 Together Forever | other | Retained precon support enchantment.
- 1 Tokka & Rahzar, Unsupervised | ramp | Retained precon mana creature.
- 1 Turtle Lair | land | Retained precon tribal mana land.
- 1 Undergrowth Stadium | land | Retained precon dual land.
- 1 Vanquish the Horde | wipe | Retained precon board reset.
- 1 Vernal Fen | land | Retained precon dual land.
- 1 Vibrant Cityscape | land | Retained precon fixing land.
- 1 Vigor | other | Retained precon creature support.
- 1 Voracious Hydra | threat | Retained precon creature threat.
- 1 Wave Goodbye | other | Retained precon support spell.
- 1 Rhystic Study | draw | Adds reliable ongoing card draw.
- 1 Mana Confluence | land | Adds an untapped five-color source.
- 1 Cavern of Souls | land | Adds tribal mana support.
- 3 Forest | land | Supplies green mana for the five-color base.
- 2 Island | land | Supplies blue mana for the five-color base.
- 2 Plains | land | Supplies white mana for the five-color base.
- 2 Swamp | land | Supplies black mana for the five-color base.
- 2 Mountain | land | Supplies red mana for the five-color base.

</details>

### 19. the Hobbit family, two colours

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 174 names.

Commander: Thranduil, the Elvenking.

Grade: bad, score 0.03, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $528.58 to buy, $528.58 the whole deck.

**Summary:** Thranduil leads an Elf-heavy creature deck that develops its mana early, keeps cards flowing, and uses a broad package of removal and interaction to protect its board position. The deck applies pressure through its legendary Elves and supporting creatures, with Troll of Khazad-dûm and the two Witch-kings chief among its closing threats. It gives up utility-land variety for a very stable, untapped three-color basic-land base, and its creature-centric plan is more focused on steady board development than explosive combo finishes.

The quality model grades this deck below the precon baseline against the top lists of the format: the blue sources cover 13.75 of the 19 its spells need.

- PLAN plan_coherent=yes: The list holds together as a Sultai Elf-legends midrange deck: cheap Elf mana dorks and Elven Chorus/Archdruid ramp into legendary Elves and big Witch-king/Troll finishers, backed by a consistent removal and draw suite, which matches Thranduil's Elf/legend-centric leadership.
- PLAN theme_fit=yes: It is a Commander deck with the requested commander, correct colors, Middle-earth flavor throughout, and a power level consistent with bracket 3 (fast mana and The One Ring are present but there is no combo kill).
- PLAN useful_as_built=partly: Curve, draw, removal and multiple creature win conditions are all playable out of the box, but a 36-basic three-color base with only ten Islands will regularly strand Stern Scolding, Lórien Revealed and other blue cards, making games inconsistent.
- PLAN summary_honest=partly: It correctly flags the lack of utility lands and the non-explosive plan, but calling the base "very stable" is spin given the deck is short roughly five blue sources for its blue spells, which the summary never admits.
- [INFO] `curve_summary`: average mana value 3.08 over 63 nonland cards
- [WARN] `profile_off_band`: the worst color's share of the sources it needs is 0.72, and bracket 3 wants 0.85 or more (sources of requirement: U 13.75 of 19, B 17.5 of 23, G 16 of 22)
- [INFO] `mana_pass`: the builder moved 3 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 1 Beorn the Fierce | draw | Draw option from the shortlist.
- 1 Bilbo, Luckwearer // Burglar's Plot | draw | Draw option from the shortlist.
- 1 Elvish Visionary | draw | Low-cost Elf that supplies draw.
- 1 Gandalf, Wandering Wizard | draw | Draw option from the shortlist.
- 1 Hithlain Knots | draw | Draw option from the shortlist.
- 1 Ithilien Kingfisher | draw | Draw option from the shortlist.
- 1 Lórien Revealed | draw | Draw option from the shortlist.
- 1 Night's Whisper | draw | Efficient draw option.
- 1 Palantír of Orthanc | draw | Artifact-based draw option.
- 1 Plunder the Trollshaws | draw | Draw option from the shortlist.
- 1 Bilbo's Ring | interaction | Interaction option from the shortlist.
- 1 Confusticate and Bebother | interaction | Interaction option from the shortlist.
- 1 Elrond, Moon-Reader | interaction | Elf interaction piece.
- 1 Mithril Coat | interaction | Equipment-based interaction option.
- 1 My Precious // Allure of Power | interaction | Interaction option from the shortlist.
- 1 Old Fat Spider Can't See Me | interaction | Saga-based interaction option.
- 1 Stern Scolding | interaction | Low-cost interaction option.
- 1 The One Ring | interaction | Legendary artifact interaction option.
- 1 Arcane Signet | ramp | Reliable multicolor ramp.
- 1 Delighted Halfling | ramp | Early creature ramp.
- 1 Elven Chorus | ramp | Elf-themed ramp piece.
- 1 Elvish Archdruid | ramp | Elf-based ramp.
- 1 Elvish Mystic | ramp | Early creature ramp.
- 1 Mox Amber | ramp | Fast legendary-focused ramp.
- 1 Silvan Reveler | ramp | Elf ramp option.
- 1 Thranduil the Strategist | ramp | Thranduil-themed ramp piece.
- 1 Wayfarer's Bauble | ramp | Two-mana ramp option.
- 1 Wood Elves | ramp | Elf-based ramp.
- 1 Bilbo's Deadly Slice | removal | Removal option from the shortlist.
- 1 Bitter Downfall | removal | Removal option from the shortlist.
- 1 Enchanted River's Grasp | removal | Aura-based removal.
- 1 Giant's Boulder | removal | Artifact-based removal.
- 1 Merciless Executioner | removal | Creature-based removal.
- 1 Orcish Bowmasters | removal | Creature-based removal.
- 1 Quarrel | removal | Low-cost removal option.
- 1 The Black Arrow | removal | Equipment-based removal.
- 1 Uneasy Partings | removal | Instant-speed removal option.
- 1 Gnashing of Teeth | wipe | Board-wipe option.
- 1 Languish | wipe | Board-wipe option.
- 1 Raise the Palisade | wipe | Board-wipe option.
- 1 Troll of Khazad-dûm | wincon | Finishing threat marked as a win condition.
- 1 Witch-king of Angmar | wincon | Finishing threat marked as a win condition.
- 1 Witch-king, Bringer of Ruin | wincon | Finishing threat marked as a win condition.
- 1 Supper for Spiders | synergy | Dedicated synergy card from the shortlist.
- 1 Arwen, Weaver of Hope | other | Legendary Elf Noble for the creature suite.
- 1 Boughside Wanderers | other | Elf Scout creature slot.
- 1 Cantankerous Keepers | other | Elf Soldier creature slot.
- 1 Celeborn the Wise | other | Legendary Elf Noble creature slot.
- 1 Elven Raft-Steerer | other | Elf Pilot creature slot.
- 1 Elvenking's Harper | other | Elf Bard creature slot.
- 1 Galadhrim Guide | other | Elf Scout creature slot.
- 1 Galion, Elvenking's Butler | other | Legendary Elf Advisor creature slot.
- 1 Guardian of the Halls | other | Elf Soldier creature slot.
- 1 Lothlórien Lookout | other | Elf Scout creature slot.
- 1 Mirkwood Meditator | other | Elf Druid creature slot.
- 1 Mirkwood Nurturer | other | Elf Ranger creature slot.
- 1 Mirkwood Pathmaker | other | Elf Ranger creature slot.
- 1 Nimrodel Watcher | other | Elf Scout creature slot.
- 1 Thranduil, Sindarin Liege // Silvan Rally | other | Thranduil-themed legendary Elf slot.
- 1 Andúril, Flame of the West | other | Legendary Equipment slot.
- 1 Glamdring | other | Legendary Equipment slot.
- 1 Sting, Bilbo's Sword | other | Legendary Equipment slot.
- 13 Forest | land | Basic green source with no tapped-land drawback.
- 10 Island | land | Basic blue source with no tapped-land drawback.
- 13 Swamp | land | Basic black source with no tapped-land drawback.
- 1 Pelargir Survivor | ramp | the mana pass added it to bring the mana base inside the power level

</details>

### 20. the Hobbit family, a delegated commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 77 names.

Commander: Smaug the Magnificent.

Grade: baseline, score 0.04, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $463.03 to buy, $463.03 the whole deck.

**Summary:** Smaug the Magnificent leads a red Hobbit dragon deck that develops mana, uses Equipment and legendary creatures to build a board, and clears away resistance with focused answers or broad resets. Its strongest closing plays are its dragon finishers, chief among them Cavern-Hoard Dragon, Desert Were-Worm, and Smaug, the Great Calamity // Spew Flame. The deck gives up some precision and card-selection depth for a heavily themed collection of creatures, artifacts, and dramatic top-end threats.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=no: The list is a scattering of dwarves, goblins, random legends, a dozen unrelated Equipment, and two or three big dragons with no shared engine tying them together, and the summary's own 'Equipment and legendary creatures build a board' framing is not supported by any consistent payoff (no Equipment tutors, no legendary-matters cards, no dragon tribal support).
- PLAN theme_fit=no: The person asked for a dragons deck and, aside from the commander, the list contains roughly two actual dragons (Cavern-Hoard Dragon and Smaug, the Great Calamity) amid a mass of dwarves, goblins, and Equipment, so it answers a different request.
- PLAN useful_as_built=partly: The mono-red mana base with 36 lands and a low curve will function and the deck can attack with equipped creatures, but the ramp is thin for a deck topped by expensive dragons and there are very few real closing threats, so games will often stall into an unfocused board.
- PLAN summary_honest=partly: placeholder
- [INFO] `curve_summary`: average mana value 3.13 over 63 nonland cards
- [WARN] `profile_off_band`: the draw count is 6, and bracket 3 wants 8 to 14

<details><summary>The deck list</summary>

- 1 Arcane Signet | ramp | Provides reliable early mana acceleration.
- 1 Bag End Banquet | ramp | Adds to the deck's mana development.
- 1 Burn, Burn, Tree and Fern | ramp | Supports the mana plan while keeping the Hobbit theme.
- 1 Dragon's Desire | ramp | Helps reach the mana needed for the dragon package.
- 1 Fíli and Kíli, Joyous | ramp | Contributes mana acceleration in a thematic creature slot.
- 1 Mox Amber | ramp | Provides efficient early acceleration alongside the commander and legends.
- 1 Orcrist, Goblin-cleaver | ramp | Adds mana support from an on-theme Equipment.
- 1 The Misty Mountains Cold | ramp | Builds toward the deck's larger plays.
- 1 The Reaver Cleaver | ramp | Supports the mana plan while rewarding creature deployment.
- 1 Troop of Ponies | ramp | Provides another thematic mana source.
- 1 Wayfarer's Bauble | ramp | Gives the deck early mana development.
- 1 Battle-Scarred Goblin | removal | Provides a creature-based answer to opposing problems.
- 1 Fire of Orthanc | removal | Offers a direct removal spell.
- 1 Gandalf, Spark Starter | removal | Adds a legendary removal option.
- 1 Goblin Cratermaker | removal | Provides flexible creature-based removal.
- 1 Goblin Fireleaper | removal | Adds another thematic removal body.
- 1 Improvised Club | removal | Provides a compact removal option.
- 1 Pinecone Strike | removal | Adds instant-speed removal coverage.
- 1 Smite the Deathless | removal | Provides another direct answer.
- 1 The Black Arrow | removal | Brings removal in an Equipment slot.
- 1 Call Forth the Tempest | wipe | Resets crowded boards when single-target answers are insufficient.
- 1 Desolation of Smaug | wipe | Provides a second broad battlefield reset.
- 1 Glóin the Mighty // Easy Pickings | wipe | Adds a creature slot with board-clearing utility.
- 1 Cavern-Hoard Dragon | wincon | Acts as a major dragon payoff and closing threat.
- 1 Desert Were-Worm | wincon | Provides a large finisher for the deck's endgame.
- 1 Smaug, the Great Calamity // Spew Flame | wincon | Adds another powerful dragon-based path to closing games.
- 1 Balin, Loremaster | draw | Supplies a thematic source of cards.
- 1 Key to the Side-Door | draw | Adds card flow from a low-investment artifact.
- 1 Palantír of Orthanc | draw | Provides a persistent card-advantage option.
- 1 Ragged Short Spear | draw | Adds card flow in an Equipment slot.
- 1 Thrór's Map | draw | Provides another artifact-based draw option.
- 1 Óin the Brave | draw | Adds card advantage on a thematic creature.
- 1 Bilbo's Ring | interaction | Protects the deck's plan with a useful interaction piece.
- 1 Dwarven Mattock | interaction | Provides an interactive Equipment option.
- 1 Mithril Coat | interaction | Helps preserve important creatures and the commander.
- 1 The One Ring | interaction | Provides a flexible defensive interaction piece.
- 1 Last Light of Durin's Day | synergy | Supports the deck's central thematic plan.
- 1 Bombur, Gentle Dreamer | threat | Adds a thematic creature to apply pressure.
- 1 Bothersome Noisemaker | threat | Provides an additional creature threat.
- 1 Dori, Bearer of Friends | threat | Adds another legendary creature that can pressure opponents.
- 1 Dwarven Mauler | threat | Provides a thematic battlefield threat.
- 1 Dwarven Warriors | threat | Adds another creature for the deck's board presence.
- 1 Dáin Ironfoot | threat | Provides a legendary finishing-oriented threat.
- 1 Gandalf, Goblins' Bane // Flameshape | threat | Adds a versatile legendary threat.
- 1 Goblin-town Flunkies | threat | Provides another low-cost creature threat.
- 1 Gundabad Opportunist | threat | Adds a creature that helps maintain battlefield pressure.
- 1 Guttersnipe | threat | Provides a threat that complements the deck's spells.
- 1 Iron Hills Stalwart | threat | Adds a sturdy thematic creature threat.
- 1 Misty Mountains Raider | threat | Provides another creature for early pressure.
- 1 Andúril, Flame of the West | other | Adds a legendary Equipment to support the creature plan.
- 1 Andúril, Narsil Reforged | other | Provides another thematic legendary Equipment.
- 1 Getaway Barrel | other | Adds a utility artifact to the deck's support package.
- 1 Glamdring | other | Provides an additional legendary Equipment option.
- 1 Long-Lost Lances | other | Supports the creature package as an Equipment.
- 1 Old Thrush | other | Adds a thematic utility creature.
- 1 Oliphaunt | other | Provides a large thematic creature.
- 1 Olog-hai Crusher | other | Adds another sizeable creature to the board.
- 1 Orcish Siegemaster | other | Provides a thematic creature body.
- 1 Smaug's Fury | other | Adds a thematic spell to support Smaug's game plan.
- 1 Snowslope Hunter | other | Provides another creature for the deck's board presence.
- 1 Sting, Bilbo's Sword | other | Adds a legendary Equipment utility piece.
- 1 Tidings of War | other | Provides a thematic support spell.
- 1 Well-Worn Spatula | other | Adds another low-cost Equipment support piece.
- 1 Dragon-Cursed Halls | land | Provides a themed land slot.
- 1 Hobbit Hole | land | Provides a themed land slot.
- 1 Rogue's Passage | land | Provides a utility land for the creature plan.
- 1 The Lonely Mountain | land | Provides a thematic Mountain land slot.
- 1 Treasure Vault | land | Provides a utility artifact land.
- 31 Mountain | land | Supplies dependable red mana throughout the game.

</details>

### 21. the Hobbit family, mana from outside

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 156 names.

Commander: Smaug the Impenetrable.

Grade: baseline, score 0.06, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $844.92 to buy, $844.92 the whole deck.

**Summary:** Smaug the Impenetrable leads a red-black, creature-forward deck that develops mana quickly, builds an Equipment package around its Goblins, Dragons, and legendary creatures, and keeps opposing boards manageable with plentiful direct answers and several reset cards. It wins by applying sustained combat pressure before closing with its marked Dragon and Goblin-town finishers. The deck gives up some flexibility for its dense thematic core, leaning on its permanent-based board presence and artifact package to carry long games.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The deck makes less mana on turn four than the norm, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The cards do broadly cohere around Smaug plus Equipment, Goblins and Dragons with removal, but the list pulls between a Goblin tribal subtheme, a legendary/Equipment subtheme, and a generic midrange goodstuff shell, so no single engine is clearly supported.
- PLAN theme_fit=partly: The commander, colors and bracket match and most of the deck is Hobbit-set material, but a meaningful number of nonland, non-rock cards (Night's Whisper, Guttersnipe, Inferno Titan, Goblin Cratermaker, The One Ring, Andúril, Glamdring, Sting, Sauron, Palantír, etc.) come from outside the Hobbit set, which the request only permitted for lands and mana rocks.
- PLAN useful_as_built=yes: Thirty-six lands with fixing plus a handful of rocks and ramp supports a reasonable curve, and the deck has plenty of creatures, removal and several distinct finishers to actually close a game.
- PLAN summary_honest=partly: The Equipment package, removal suite and Goblin-town/Smaug finishers are all genuinely present, but 'several reset cards' overstates the two sweepers and 'develops mana quickly' sits awkwardly with the fetch/utility-heavy base the summary itself admits makes less turn-four mana.
- [INFO] `curve_summary`: average mana value 3.22 over 63 nonland cards
- [INFO] `mana_pass`: the builder moved 2 cards of the mana base to bring the deck inside its power level
- [WARN] `outside_requested_set`: the sets you named do not hold 13 cards: Ancient Tomb, Arid Mesa, Blood Crypt, Bloodstained Mire, Bojuka Bog, Command Tower, Dragonskull Summit, Evolving Wilds, Fabled Passage, Marsh Flats, Polluted Delta, Terramorphic Expanse, Urborg, Tomb of Yawgmoth

<details><summary>The deck list</summary>

- 1 Ancient Tomb | land | An untapped land that helps accelerate the deck's mana.
- 1 Arid Mesa | land | A land slot that helps access the deck's basic mana base.
- 1 Blood Crypt | land | A black-red land source for the deck's core colors.
- 1 Bloodstained Mire | land | A land slot that helps access the deck's basic mana base.
- 1 Bojuka Bog | land | A black land slot with additional utility.
- 1 Command Tower | land | A reliable source for the commander's colors.
- 1 Dragon-Cursed Halls | land | A themed land slot in the mana base.
- 1 Dragonskull Summit | land | A black-red land source for the deck's core colors.
- 1 Evolving Wilds | land | A land slot that finds the appropriate basic land.
- 1 Fabled Passage | land | A land slot that finds the appropriate basic land.
- 1 Goblin-town | land | A themed land slot in the mana base.
- 1 Marsh Flats | land | A land slot that helps access the deck's basic mana base.
- 1 Minas Morgul, Dark Fortress | land | A themed land slot in the mana base.
- 1 Mount Doom | land | A themed land slot in the mana base.
- 10 Mountain | land | Provides a deep supply of red mana.
- 1 Polluted Delta | land | A land slot that helps access the deck's basic mana base.
- 8 Swamp | land | Provides a deep supply of black mana.
- 1 Terramorphic Expanse | land | A land slot that finds the appropriate basic land.
- 1 The Lonely Mountain | land | A themed land slot in the mana base.
- 1 Urborg, Tomb of Yawgmoth | land | A black mana source with utility.
- 1 Arcane Signet | ramp | Efficiently fixes and accelerates the deck's mana.
- 1 Bag End Banquet | ramp | Fills a dedicated mana-development slot.
- 1 Burn, Burn, Tree and Fern | ramp | Fills a dedicated mana-development slot.
- 1 Dragon's Desire | ramp | Fills a dedicated mana-development slot.
- 1 Mox Amber | ramp | A low-cost mana accelerator for the legendary-heavy list.
- 1 Orcrist, Goblin-cleaver | ramp | Adds mana development from a themed Equipment slot.
- 1 Smaug the Magnificent | ramp | A Dragon-themed source of mana development.
- 1 The Reaver Cleaver | ramp | An Equipment-based mana-development card.
- 1 Wayfarer's Bauble | ramp | An early mana-development piece that supports the land count.
- 1 Balin, Loremaster | draw | Adds a themed source of cards.
- 1 Gollum, Riddle Master | draw | A card-advantage creature with finisher potential.
- 1 Key to the Side-Door | draw | A compact artifact source of cards.
- 1 Night's Whisper | draw | A direct card-advantage spell.
- 1 Palantír of Orthanc | draw | A legendary artifact source of cards.
- 1 Rage into the Valley | draw | Adds a red card-advantage spell.
- 1 Ragged Short Spear | draw | An Equipment slot that also contributes cards.
- 1 Thrór's Map | draw | A legendary artifact source of cards.
- 1 Azog, Moria's Ruin | removal | A themed creature-based answer.
- 1 Battle-Scarred Goblin | removal | A low-cost Goblin answer.
- 1 Bilbo's Deadly Slice | removal | A direct removal spell.
- 1 Bitter Downfall | removal | A direct removal spell.
- 1 Bolg of the North | removal | A themed creature-based answer.
- 1 Crude Bent Blade | removal | An Equipment-based answer.
- 1 Fire of Orthanc | removal | A direct red removal spell.
- 1 Goblin Cratermaker | removal | A flexible Goblin-based answer.
- 1 Goblin Fireleaper | removal | A Goblin creature that fills a removal slot.
- 1 Bolg, Erebor's Reckoning | wipe | A themed board-reset card.
- 1 Desolation of Smaug | wipe | A Smaug-themed board-reset spell.
- 1 Andúril, Flame of the West | synergy | A legendary Equipment that supports the deck's artifact-and-combat package.
- 1 Andúril, Narsil Reforged | synergy | A legendary Equipment that supports the deck's artifact-and-combat package.
- 1 Glamdring | synergy | A legendary Equipment that supports the deck's artifact-and-combat package.
- 1 Goblin Plate Mail | synergy | An Equipment that supports the creature combat plan.
- 1 Sting, Bilbo's Sword | synergy | A legendary Equipment that supports the creature combat plan.
- 1 Supper for Spiders | synergy | The shortlist's dedicated synergy card.
- 1 Well-Worn Spatula | synergy | An Equipment that supports the deck's artifact-and-combat package.
- 1 Dreaded Bat-Cloud | threat | A creature threat that adds pressure to the board.
- 1 Fearsome Goblin Pair | threat | A Goblin threat that contributes to the attacking force.
- 1 Gandalf, Goblins' Bane // Flameshape | threat | A legendary threat with an additional Adventure face.
- 1 Gollum the Abandoned | threat | A legendary creature threat.
- 1 Gollum, Silent Slinker // Meager Meal | threat | A legendary creature threat with an additional Adventure face.
- 1 Great Goblin, Foul-Hearted | threat | A legendary Goblin threat.
- 1 Great Ugly-Looking Goblin // Clap! Snap! | threat | A creature threat with an additional Adventure face.
- 1 Gundabad Opportunist | threat | A Goblin threat that contributes to the attacking force.
- 1 Guttersnipe | threat | A creature threat that rewards the deck's spell suite.
- 1 Inferno Titan | threat | A large red creature threat.
- 1 Olog-hai Crusher | threat | A large creature threat for combat-focused games.
- 1 Sauron, the Lidless Eye | threat | A legendary Avatar threat at the top of the curve.
- 1 Down, Down to Goblin-town | wincon | A marked finisher that gives the deck a dedicated closing card.
- 1 Smaug, Wicked Worm | wincon | A Dragon finisher that reinforces Smaug's theme.
- 1 Smaug, the Great Calamity // Spew Flame | wincon | A marked Dragon finisher with an Adventure face.
- 1 Bilbo's Ring | interaction | A legendary Equipment that provides a dedicated utility slot.
- 1 Dwarven Mattock | interaction | An Equipment-based utility card.
- 1 Gathering of Darkness | interaction | A black utility spell for handling opposing plans.
- 1 Getaway Barrel | interaction | A flexible artifact utility card.
- 1 Long-Lost Lances | interaction | An Equipment-based utility card.
- 1 Mithril Coat | interaction | A legendary Equipment utility card.
- 1 My Precious // Allure of Power | interaction | A legendary Equipment utility card with an Adventure face.
- 1 Shadow of the Enemy | interaction | A black utility spell for interacting with opposing plans.
- 1 Smaug's Fury | interaction | A Smaug-themed instant-speed utility spell.
- 1 The One Ring | interaction | A legendary artifact utility card.
- 1 The Black Arrow | interaction | A legendary Equipment utility card.
- 1 Giant's Boulder | ramp | the mana pass added it to bring the mana base inside the power level
- 1 Bolg's Company | ramp | the mana pass added it to bring the mana base inside the power level

</details>

### 22. a set family and a card from outside it

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 175 names.

Commander: Thranduil, the Elvenking.

Grade: bad, score 0.02, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $496.14 to buy, $496.14 the whole deck.

**Summary:** This is a Simic Elf-heavy Thranduil deck that develops its mana early, builds a broad creature board, and keeps cards flowing through a deep suite of draw options. It presses combat with its Elf core before closing through Colossal Whale, Old Fat Spider, or Inside Information. The deck gives up some broad answer coverage in exchange for consistent mana, board development, and a strongly thematic creature plan.

The quality model grades this deck below the precon baseline against the top lists of the format: the black sources cover 17.25 of the 28 its spells need.

- PLAN plan_coherent=partly: There is a recognizable Elf/creature go-wide core with ramp and card draw attached to Thranduil, but the list also sprawls into a large pile of unrelated draw, equipment, and one-off effects, and the summary's 'Simic' frame does not match a list that runs 12 Swamps and black spells, so the pieces pull in more than one direction.
- PLAN theme_fit=yes: The deck is a Commander list led by Thranduil, the Elvenking, drawn almost entirely from Middle-earth/Hobbit-flavored cards at a modest bracket-3 power level, and it keeps the requested Sol Ring.
- PLAN useful_as_built=partly: 36 lands that are all basics split across three colors, with black demand well above its sources, means frequent color screw and stranded spells, though the low curve and heavy ramp/draw let it still function as a creature-beatdown deck much of the time.
- PLAN summary_honest=no: It describes the deck as Simic when the list is three-color with a dozen Swamps and black cards, and it presents Colossal Whale/Old Fat Spider/Inside Information as closers when they are incidental value pieces rather than a stated win plan.
- [INFO] `curve_summary`: average mana value 3.29 over 63 nonland cards
- [WARN] `profile_off_band`: the worst color's share of the sources it needs is 0.62, and bracket 3 wants 0.85 or more (sources of requirement: U 14.5 of 19, B 17.25 of 28, G 17 of 22)
- [WARN] `profile_off_band`: the removal count is 5, and bracket 3 wants 6 to 12
- [INFO] `mana_pass`: the builder moved 13 cards of the mana base to bring the deck inside its power level
- [WARN] `outside_requested_set`: the sets you named do not hold 1 card: Sol Ring

<details><summary>The deck list</summary>

- 14 Forest | land | Untapped basic green source for the deck's early Elf development.
- 10 Island | land | Untapped basic blue source for the deck's spells and commander.
- 1 Sol Ring | ramp | Kept as requested and supplies efficient acceleration.
- 1 Arcane Signet | ramp | Reliable color-fixing acceleration.
- 1 Delighted Halfling | ramp | Early creature-based acceleration for the Elf plan.
- 1 Elven Chorus | ramp | Elf-focused mana development.
- 1 Elvish Archdruid | ramp | A powerful Elf-based mana engine.
- 1 Elvish Mystic | ramp | Low-cost green acceleration.
- 1 Mox Amber | ramp | Fast legendary-focused acceleration.
- 1 Orcrist, Goblin-cleaver | ramp | An equipment slot that also contributes mana development.
- 1 Thranduil the Strategist | ramp | A Thranduil-themed source of acceleration.
- 1 Wood Elves | ramp | An Elf body that advances mana development.
- 1 Woodland Weavemaster | ramp | Additional Elf-based mana development.
- 1 Beorn the Fierce | draw | A draw option attached to a substantial creature.
- 1 Elvish Visionary | draw | An inexpensive Elf that supports card flow.
- 1 Fateful Discovery | draw | Dedicated card-flow support.
- 1 Gandalf, Shadow's Foe | draw | A legendary draw piece for the midgame.
- 1 Hithlain Knots | draw | Instant-speed card-flow support.
- 1 Ithilien Kingfisher | draw | Creature-based card-flow support.
- 1 Key to the Side-Door | draw | Artifact-based card-flow support.
- 1 Last March of the Ents | draw | A high-impact draw spell for rebuilding resources.
- 1 Lórien Revealed | draw | Blue card-flow support.
- 1 Palantír of Orthanc | draw | A persistent artifact draw option.
- 1 Uncover the Moon-Letters | draw | Enchantment-based card-flow support.
- 1 Long Lake Nuisance | draw | An additional creature-based draw option.
- 1 Reverent Howl | draw | Extra card-flow support for longer games.
- 1 Plunder the Trollshaws | draw | A further draw spell that keeps resources coming.
- 1 Crude Bent Blade | removal | Equipment-based removal coverage.
- 1 Enchanted River's Grasp | removal | Blue removal coverage for problematic permanents.
- 1 Giant's Boulder | removal | Colorless removal coverage.
- 1 Quarrel | removal | Low-profile removal coverage.
- 1 The Black Arrow | removal | Legendary equipment that supplies another answer.
- 1 Confusticate and Bebother | interaction | A flexible instant-speed interaction slot.
- 1 Dwarven Mattock | interaction | Equipment-based interaction coverage.
- 1 Elrond, Moon-Reader | interaction | A legendary interaction piece that fits the Elf shell.
- 1 Mithril Coat | interaction | Protective interaction for important creatures.
- 1 Old Fat Spider Can't See Me | interaction | A thematic interaction piece.
- 1 Stern Scolding | interaction | Efficient blue stack interaction.
- 1 The One Ring | interaction | A legendary artifact interaction option.
- 1 Thranduil's Decree | interaction | Thranduil-themed interaction coverage.
- 1 Gnashing of Teeth | wipe | A board-reset option when creatures pile up.
- 1 Raise the Palisade | wipe | A tribal-aware board-reset option.
- 1 Arwen, Weaver of Hope | synergy | An Elf legend that reinforces the deck's creature core.
- 1 Celeborn the Wise | synergy | An Elf legend that supports the deck's thematic core.
- 1 Galadhrim Guide | synergy | An Elf that deepens the deck's tribal concentration.
- 1 Galion, Elvenking's Butler | synergy | A Thranduil-associated Elf for the thematic shell.
- 1 Lothlórien Lookout | synergy | An Elf body that reinforces the deck's creature base.
- 1 Mirkwood Meditator | synergy | A Mirkwood Elf that supports the central theme.
- 1 Mirkwood Nurturer | synergy | An Elf inclusion that builds the board-focused plan.
- 1 Attercop | threat | A sizeable green creature threat.
- 1 Boughside Wanderers | threat | An Elf creature that adds pressure to the board.
- 1 Cantankerous Keepers | threat | An Elf creature that adds another board presence.
- 1 Elven Raft-Steerer | threat | An Elf threat that contributes to the creature density.
- 1 Elvenking's Harper | threat | A Thranduil-themed Elf threat.
- 1 Great Fierce Bee | threat | A green creature that provides an additional attacker.
- 1 Grey Havens Navigator | threat | An Elf creature that broadens the board presence.
- 1 Guardian of the Halls | threat | An Elf creature that adds battlefield pressure.
- 1 Mirkwood Elk | threat | A green creature threat from Thranduil's domain.
- 1 Mirkwood Pathmaker | threat | An Elf creature that supports the board plan.
- 1 Nimrodel Watcher | threat | An Elf creature that adds another attacker.
- 1 Willow-Wind | threat | A creature threat that helps maintain pressure.
- 1 Colossal Whale | wincon | A marked finisher that can close games.
- 1 Inside Information | wincon | A marked finisher for ending stalled games.
- 1 Old Fat Spider | wincon | A marked finisher that provides a large closing threat.
- 1 Pelargir Survivor | ramp | the mana pass added it to bring the mana base inside the power level
- 12 Swamp | land | the mana pass added it to bring the mana base inside the power level

</details>

### 23. two set families at once

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 203 names.

Commander: Kíli the Resourceful.

Grade: typical, score 0.44, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $256.25 to buy, $256.25 the whole deck.

**Summary:** Kíli the Resourceful leads a mono-white Hobbit and Bloomburrow creature-and-artifact deck that develops steadily through cheap mana, recurring card flow, Dwarf support, and equipment. It plays a patient battlefield game, answers key opposing permanents, and closes by turning a developed creature board into decisive combat pressure through its large threats and finishing equipment. The deck gives up explosive multicolor lines for a consistent white mana base, resilient board development, and a strong ability to rebuild after a reset.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

This deck plan aims for 2 or more finishers, the cards that win the game, and the deck holds 1.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The list is a broadly coherent mono-white creature/artifact value pile with equipment and recursion, but the named sub-themes pull apart — "Dwarf support" rests on only a handful of Dwarves, while token payoffs (Jazal, Hoofprints, Martial Coup), blink-style ETB removal, Psychosis Crawler, and Karn's artifact lock each point in a slightly different direction.
- PLAN theme_fit=no: The person named a hard limit — cards from The Hobbit and Bloomburrow — and the deck is mostly generic white staples from all over the pool (Sol Ring, Skullclamp, Swords to Plowshares, Sun Titan, Academy Manufactor, Karn, Elspeth, Blade Splicer), with only a thin scattering of cards from the two requested sets.
- PLAN useful_as_built=yes: 36 lands plus ten mana rocks/ramp in a single color, a smooth curve, ample draw and removal, and enough top-end (Elspeth, Angel of the Ruins, Jazal, Helm of the Host, Andúril) that the deck can be shuffled up and closed out games.
- PLAN summary_honest=partly: It is candid about the single-finisher count and the mono-white tradeoff, but it calls the list a "Hobbit and Bloomburrow" deck when the large majority of cards come from neither set, and it oversells "Dwarf support" that amounts to a few creatures and Patchwork Banner.
- [INFO] `curve_summary`: average mana value 3.10 over 63 nonland cards
- [WARN] `finisher_short`: the finisher count is 1, and this deck plan wants 2 or more (Angel of the Ruins)

<details><summary>The deck list</summary>

- 28 Plains | land | Reliable white mana for the deck’s heavily white core.
- 1 Castle Ardenvale | land | A white-producing utility land.
- 1 Command Tower | land | Reliable commander-color mana.
- 1 Evolving Wilds | land | Fixes for a basic Plains while supporting land count.
- 1 Fabled Passage | land | Finds Plains while remaining a flexible land slot.
- 1 Minas Tirith | land | A white land with useful utility.
- 1 Path of Ancestry | land | Provides white mana for the creature-focused list.
- 1 Thriving Heath | land | Adds another dependable white source.
- 1 Uncharted Haven | land | Flexible colored mana in a single-color shell.
- 1 Arcane Signet | ramp | Efficient mana acceleration for Kíli and the deck’s spells.
- 1 Burnished Hart | ramp | Creature-based mana development for longer games.
- 1 Fellwar Stone | ramp | Low-cost mana acceleration.
- 1 Loyal Warhound | ramp | Develops mana while contributing to the battlefield.
- 1 Mind Stone | ramp | Early acceleration with later utility.
- 1 Mox Amber | ramp | Fast legendary-focused mana support.
- 1 Ornithopter of Paradise | ramp | Early colored mana from a creature.
- 1 Patchwork Banner | ramp | Mana support that complements the creature plan.
- 1 Sol Ring | ramp | Fast, efficient colorless acceleration.
- 1 Solemn Simulacrum | ramp | Reliable mana development attached to a creature.
- 1 Caretaker's Talent | draw | A durable card-advantage piece for extended games.
- 1 Circuit Mender | draw | A creature that replaces itself in hand resources.
- 1 Cut a Deal | draw | Refills the hand when resources run low.
- 1 Dawn of a New Age | draw | Ongoing card advantage for the board-building plan.
- 1 Heirloom Epic | draw | Artifact-based card advantage.
- 1 Idol of Oblivion | draw | Efficient repeatable card advantage.
- 1 Inspiring Overseer | draw | A flying creature that supplies a card.
- 1 Mentor of the Meek | draw | Rewards the deck’s stream of smaller creatures.
- 1 Skullclamp | draw | Turns expendable creatures into cards.
- 1 Spirited Companion | draw | Low-cost creature that keeps cards flowing.
- 1 The Gaffer | draw | A creature-based source of continued card advantage.
- 1 Baird, Steward of Argive | interaction | Helps discourage attacks while the board develops.
- 1 Bilbo's Gambit | interaction | A flexible instant-speed disruptive tool.
- 1 Bilbo's Ring | interaction | Protective utility for an important creature.
- 1 Dawn's Truce | interaction | Defensive instant-speed support for the board.
- 1 Mithril Coat | interaction | Protects a key legendary permanent.
- 1 Reprieve | interaction | Temporarily answers a problematic spell while cycling onward.
- 1 Selfless Spirit | interaction | Creature-based protection against opposing pressure.
- 1 Swiftfoot Boots | interaction | Protects Kíli or a major threat.
- 1 Banishing Light | removal | Versatile answer to a troublesome permanent.
- 1 Fiend Hunter | removal | Creature-based answer that supports the board plan.
- 1 Generous Gift | removal | Broad instant-speed permanent answer.
- 1 Hanged Executioner | removal | A body that can convert into creature removal.
- 1 Loran of the Third Path | removal | Useful permanent removal on a creature.
- 1 Skyclave Apparition | removal | Efficient creature-based permanent removal.
- 1 Swords to Plowshares | removal | Premium low-cost creature answer.
- 1 The Black Arrow | removal | Equipment-based removal utility.
- 1 Westfold Rider | removal | Creature-based removal that remains relevant on board.
- 1 Academy Manufactor | synergy | Supports the deck’s artifact-resource package.
- 1 Blade Splicer | synergy | Builds a creature board while fitting artifact themes.
- 1 Carrot Cake | synergy | Supports artifact and token-oriented resource play.
- 1 Dáin, Lord of the Iron Hills | synergy | A Dwarf-focused payoff alongside Kíli.
- 1 Iron Hills Blacksmith | synergy | Supports the deck’s equipment and artifact elements.
- 1 Ori, Keeper of Songs | synergy | A Dwarf-focused support piece.
- 1 Tangle Tumbler | synergy | Artifact synergy that contributes to the board.
- 1 Fíli the Pathfinder | threat | A named threat that pressures opponents through combat.
- 1 Karn, the Great Creator | threat | A high-impact permanent that demands attention.
- 1 Landroval, Horizon Witness | threat | A legendary flying threat for the battlefield.
- 1 Luminous Broodmoth | threat | A resilient flying threat that strengthens creature pressure.
- 1 Psychosis Crawler | threat | A scaling artifact creature that creates pressure.
- 1 Realm-Cloaked Giant // Cast Off | threat | A substantial creature threat after the board is cleared.
- 1 Serra Redeemer | threat | A flying creature that adds sustained pressure.
- 1 Sun Titan | threat | A large creature that generates lasting battlefield value.
- 1 Sunscorch Regent | threat | A sizable flying threat for combat-focused finishes.
- 1 Jazal Goldmane | threat | A legendary combat threat for a wide board.
- 1 Warren Warleader | threat | A Rabbit threat that helps maintain combat pressure.
- 1 Hoofprints of the Stag | threat | A persistent source of board pressure.
- 1 Angel of the Ruins | wincon | A finisher that provides a powerful top-end threat.
- 1 Andúril, Flame of the West | wincon | Equipment that turns an established creature board into a closing force.
- 1 Helm of the Host | wincon | A high-impact equipment for overwhelming value and pressure.
- 1 Dusk // Dawn | wipe | A flexible reset that can preserve the deck’s creature plan.
- 1 Elspeth, Sun's Champion | wipe | A sweeping planeswalker that can reset creature-heavy boards.
- 1 Martial Coup | wipe | A board reset that leaves behind a force to rebuild with.

</details>

### 24. a 60-card deck from one set

Format: Modern. Theme: aggro. Pool: any_card. Shortlist: 96 names.

Grade: baseline, score 0.39, model 20260914T154223Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $33.66 to buy, $33.66 the whole deck.

**Summary:** This mono-red Bloomburrow aggro deck opens with a dense group of small threats, uses Emberheart Challenger and Hearthborn Battler to keep that pressure cohesive, and clears blockers with Agate Assault and Blooming Blast. It wins by turning early creature deployment into sustained combat damage, with Dragonhawk, Fate's Tempest as a stronger follow-up when the game runs longer. The deck gives up broad answers and late-game inevitability in exchange for a streamlined red mana base, fast starts, and enough card flow to keep attacking after the first wave.

The quality model grades this deck at the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. The color sources cover the pips, and that raises the grade.

- PLAN plan_coherent=yes: Nearly every slot is a cheap red creature, a burn spell that doubles as removal, or a cheap trick/draw spell that triggers Emberheart Challenger, so the cards all push the same low-curve beatdown plan.
- PLAN theme_fit=yes: It is a 60-card Modern mono-red aggro list drawn entirely from Bloomburrow, matching every constraint the request named, at a casual power level.
- PLAN useful_as_built=yes: Twenty-four Mountains cast everything with no color risk and the curve tops at the two Dragonhawks, so the deck functions as written, though 24 lands is a touch high for a curve this low.
- PLAN summary_honest=yes: The summary names exactly the cards and roles present (Challenger/Battler as engines, Agate Assault and Blooming Blast as removal, Dragonhawk as the top end) and openly concedes the missing late game and broad answers.
- [INFO] `curve_summary`: average mana value 3.11 over 36 nonland cards

<details><summary>The deck list</summary>

- 24 Mountain | land | Provides reliable red mana without entering tapped.
- 2 Brazen Collector | ramp | Supplies the deck's limited mana acceleration while remaining part of the aggressive creature base.
- 2 Artist's Talent | draw | Adds red card flow for games where the first attack does not finish the job.
- 2 Might of the Meek | draw | Provides efficient card flow alongside the deck's small aggressive creatures.
- 2 Sazacap's Brew | draw | Helps refill resources after committing cards to the board.
- 4 Agate Assault | removal | Gives the deck cheap red removal to clear opposing blockers.
- 2 Blooming Blast | removal | Adds further removal so attacks can keep connecting.
- 4 Emberheart Challenger | synergy | Forms a core aggressive synergy package for early pressure.
- 4 Hearthborn Battler | synergy | Supports the deck's creature-focused aggressive plan.
- 4 Frilled Sparkshooter | threat | Provides a reliable red threat that advances the attack plan.
- 4 Reptilian Recruiter | threat | Adds another efficient threat to maintain battlefield pressure.
- 4 Teapot Slinger | threat | Fills out the low-end threat base for fast starts.
- 2 Dragonhawk, Fate's Tempest | threat | Provides a larger top-end threat when early pressure is answered.

</details>

### 25. use no card of an owned precon

Format: Commander. Theme: Avengers superheroes. Pool: owned_first. Shortlist: 154 names.

Commander: Syr Konrad, the Grim.

Grade: baseline, score 0.08, model 20260914T154223Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $177.54 the whole deck.

**Summary:** Syr Konrad, the Grim leads a mono-black attrition deck that develops through artifact mana, maintains pressure with a broad creature suite, and uses its marked finishers to close games. The list carries plentiful targeted answers, several reset buttons, and protective equipment for its key pieces. It gives up multicolor flexibility for a dependable Swamp-heavy mana base and a direct, board-focused game plan.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: There is a real -1/-1 counter cluster (Blowfly Infestation, Midnight Banshee, Carnifex Demon, Grim Poppet, Soul Snuffers, Contagion Clasp, Black Sun's Zenith, Skinrender) that meshes with Syr Konrad's death triggers, but the summary never names that engine and a scattering of unrelated equipment and draw trinkets (Buster Sword, Meteor Sword, Instant Ramen, Lembas, Mirror of Galadriel) pull toward a generic goodstuff pile rather than one plan.
- PLAN theme_fit=no: The person asked for an Avengers superheroes deck and received a Syr Konrad mono-black attrition list built out of Lord of the Rings, Final Fantasy, Avatar and Phyrexian cards with no Marvel or superhero identity anywhere in the ninety-nine.
- PLAN useful_as_built=yes: Thirty-six Swamps plus ten mana rocks support a mono-black curve comfortably, and the deck has abundant interaction, board wipes and clear finishers in Konrad drain, Vraska and the large fliers, so it can be shuffled up and played as printed.
- PLAN summary_honest=partly: The description of a mono-black attrition deck with artifact mana, removal, wipes and protective equipment matches the list, but it omits the deck's actual -1/-1 counter engine, calls the creature suite 'broad' when it is roughly sixteen bodies, and never acknowledges that the requested Avengers theme is entirely absent.
- [WARN] `not_owned`: Syr Konrad, the Grim: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.19 over 63 nonland cards

<details><summary>The deck list</summary>

- 36 Swamp | land | Provides a reliable mono-black mana base.
- 1 Call of the Ring | draw | Supplies a dedicated card-advantage slot.
- 1 Grave Venerations | draw | Adds card advantage while remaining a marked finisher.
- 1 Hoarder's Greed | draw | Supplies a dedicated card-advantage slot.
- 1 Idol of Oblivion | draw | Adds an artifact-based card-advantage piece.
- 1 Instant Ramen | draw | Fills an inexpensive card-advantage slot.
- 1 Lembas | draw | Provides artifact-based card advantage.
- 1 Mask of Memory | draw | Adds an equipment-based card-advantage option.
- 1 Mirror of Galadriel | draw | Provides a persistent card-advantage piece.
- 1 Nasty End | draw | Adds a black card-advantage spell.
- 1 Night's Whisper | draw | Supplies efficient card advantage.
- 1 Skullclamp | draw | Adds an equipment-based card-advantage piece.
- 1 Champion's Helm | interaction | Protects an important creature or commander.
- 1 Darksteel Plate | interaction | Provides a durable protection option.
- 1 Lightning Greaves | interaction | Provides an efficient protection option.
- 1 Swiftfoot Boots | interaction | Provides another equipment-based protection option.
- 1 Arcane Signet | ramp | Provides efficient color fixing and mana development.
- 1 Bender's Waterskin | ramp | Fills a mana-development slot.
- 1 Blitzball | ramp | Fills a mana-development slot.
- 1 Chromatic Lantern | ramp | Provides mana development and fixing.
- 1 Coin of Mastery | ramp | Fills a mana-development slot.
- 1 Commander's Sphere | ramp | Provides dependable mana development.
- 1 Inherited Envelope | ramp | Fills a mana-development slot.
- 1 Sol Ring | ramp | Provides efficient early mana development.
- 1 Thought Vessel | ramp | Adds artifact-based mana development.
- 1 Wayfarer's Bauble | ramp | Provides early mana development.
- 1 Bitter Downfall | removal | Adds a focused removal spell.
- 1 Bitter Triumph | removal | Adds flexible targeted removal.
- 1 Claim the Precious | removal | Provides a black removal option.
- 1 Dismember | removal | Adds efficient targeted removal.
- 1 Epic Downfall | removal | Provides a dedicated removal spell.
- 1 Fatal Push | removal | Adds efficient targeted removal.
- 1 Heartless Act | removal | Provides flexible targeted removal.
- 1 Infernal Grasp | removal | Adds dependable targeted removal.
- 1 Nameless Inversion | removal | Provides another efficient removal option.
- 1 Black Sun's Zenith | wipe | Provides a scalable battlefield reset.
- 1 Darkness Descends | wipe | Adds a dedicated battlefield reset.
- 1 Extinction Event | wipe | Provides another broad reset option.
- 1 Archfiend of Ifnir | wincon | Serves as a marked finisher for closing games.
- 1 Vraska, Betrayal's Sting | wincon | Serves as a marked finisher for closing games.
- 1 Witch-king of Angmar | wincon | Serves as a marked finisher for closing games.
- 1 Blowfly Infestation | synergy | Supports the deck's synergy package.
- 1 Contagion Clasp | synergy | Supports the deck's synergy package.
- 1 Deadly Dispute | synergy | Supports the deck's synergy package.
- 1 Incremental Blight | synergy | Supports the deck's synergy package.
- 1 Midnight Banshee | synergy | Supports the deck's synergy package.
- 1 Orcish Bowmasters | synergy | Supports the deck's synergy package.
- 1 Stone of Erech | synergy | Supports the deck's synergy package.
- 1 Carnifex Demon | threat | Provides a substantial creature threat and is a marked finisher.
- 1 Dusk Urchins | threat | Adds creature pressure to the board.
- 1 Grim Poppet | threat | Provides a creature-based threat.
- 1 Massacre Girl, Known Killer | threat | Adds a legendary creature threat.
- 1 Rat King, Pale Piper | threat | Adds a legendary creature threat.
- 1 Ringwraiths | threat | Provides a creature-based threat.
- 1 Sinister Gnarlbark | threat | Adds creature pressure to the board.
- 1 Skinrender | threat | Provides a creature-based threat.
- 1 Soul Snuffers | threat | Adds creature pressure to the board.
- 1 Undercity Dire Rat | threat | Provides a creature-based threat.
- 1 Voracious Fell Beast | threat | Adds a substantial creature threat.
- 1 Zodiark, Umbral God | threat | Provides a large creature threat.
- 1 Buster Sword | other | Adds a flexible equipment slot.
- 1 Meteor Sword | other | Adds a flexible equipment slot.
- 1 The Fire Nation Drill | other | Adds a standalone artifact slot.
- 1 Vanquisher's Banner | other | Adds a standalone artifact slot.

</details>

