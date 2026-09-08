# PR-8 deck gate

Run date: 2026-09-08. Card snapshot: 2026-09-03.

Verdict: PASS. 25 of 25 decks passed every block check, 0 invented names reached the user, and 0 summaries stated a false rule of the game. All three bars are zero tolerance.

## Summary

| Measure | Value |
|---|---|
| Prompts | 25 |
| Decks returned | 25 |
| Decks with no block finding | 25 |
| Invented names that reached the user | 0 |
| Decks that needed the repair turn | 3 |
| Summaries judged (F-26) | 25 |
| Summaries that state a rule of the game | 0 |
| Summaries that state a FALSE rule | 0 |
| Judge errors | 0 |
| Decks the plan judge read (PR-15, information) | 25 |
| Mean plan score, 0 to 1 | 0.73 |
| Plan reasons the judge left empty | 1 |
| Errors | 0 |
| Prompt version | 12 |
| Calls | 78 |
| Cost | $2.7104 |
| Time | 2195 seconds |

## Run

- Suite `decks`, run `pr8-deck-gate-run18`, on 2026-09-08, commit `f211a9b`.
- Roles: generate on `gpt-5.6-terra` (openai, effort medium), judge on `claude-opus-5` (anthropic, effort medium), repair on `gpt-5.6-terra` (openai, effort medium).
- Versions: card snapshot 2026-09-03, generate prompt version 12, plan_rubric prompt version 2, precons `5.3.0+20260903`, quality_model `20260907T081149Z`.
- Calls: 78. Cost: $2.7104. Time: 2195 seconds.

## Findings by code

A block stops the deck. A warning and a note are reports.

| Code | Count |
|---|---|
| `not_owned` | 26 |
| `curve_summary` | 25 |
| `profile_off_band` | 10 |
| `mana_pass` | 5 |
| `outside_requested_set` | 2 |
| `two_card_combo` | 1 |
| `cards_trimmed` | 1 |
| `basics_added` | 1 |

By severity: BLOCK 0. WARN 39. INFO 32. 

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

Format: Commander. Theme: lifegain. Pool: any_card. Shortlist: 303 names.

Commander: Karlov of the Ghost Council.

Grade: bad, score 0.04, model 20260907T081149Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $452.15 to buy, $452.15 the whole deck.

**Summary:** Karlov leads a creature-forward lifegain deck that develops its mana, repeatedly turns life gain into larger threats and card advantage, and uses targeted answers or resets to keep opposing boards manageable. The deck wins by building pressure through its lifegain payoffs, Angels, Vampires, and Clerics, with powerful payoff enchantments helping turn a stable life total into a closing position. It gives up some speed and top-end inevitability for a broad, combat-focused board presence and a fair multiplayer game plan.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every nonland card is either a lifegain enabler, a lifegain payoff (Pridemate-style growers, Righteous Valkyrie, Vito/Sanguine Bond drains, Well of Lost Dreams/Sigarda's Splendor draw) or a way to feed Karlov's counters, so the pieces all point at one lifegain plan, with only the two sweepers sitting slightly awkwardly in a creature-heavy list.
- PLAN theme_fit=yes: It is a Karlov of the Ghost Council Commander lifegain deck in the requested colors with a fair, creature-based multiplayer plan and few game changers, matching the bracket-3 request apart from the undiscussed two-card infinite.
- PLAN useful_as_built=yes: 37 lands with solid W/B duals plus about eight rocks support a curve that tops out cheaply, and the deck can close games through combat, Karlov, Aetherflux Reservoir, or Vito/Sanguine Bond drains.
- PLAN summary_honest=partly: The summary accurately describes the lifegain payoffs, tribal creatures, removal and resets, but it does not disclose that the list contains an infinite combo (Archangel of Thune plus Sanguine Bond, and Heliod-style engines) which is relevant to the requested bracket-3 framing.
- [INFO] `curve_summary`: average mana value 3.22 over 63 nonland cards

<details><summary>The deck list</summary>

- 12 Plains | land | Basic white source for the early lifegain shell.
- 12 Swamp | land | Basic black source for the deck's black spells.
- 1 Command Tower | land | Reliable dual-color land for Karlov's colors.
- 1 Godless Shrine | land | Dual white-black source.
- 1 Caves of Koilos | land | Untapped dual-color source.
- 1 City of Brass | land | Flexible color fixing.
- 1 Mana Confluence | land | Flexible color fixing.
- 1 Exotic Orchard | land | Additional multicolor land source.
- 1 Reflecting Pool | land | Fixes either deck color alongside the dual lands.
- 1 Vault of Champions | land | Dual white-black land for multiplayer games.
- 1 Shattered Sanctum | land | Dual white-black source.
- 1 Isolated Chapel | land | Dual white-black land support.
- 1 Scoured Barrens | land | A dual land that also supports the lifegain plan.
- 1 Temple of Silence | land | Dual-color land with card selection.
- 1 Altar of the Pantheon | ramp | Fixes mana while advancing the mana base.
- 1 Bounty Board | ramp | Low-cost mana development.
- 1 Hot Dog Cart | ramp | Early ramp option for reaching the commander sooner.
- 1 Nuka-Cola Vending Machine | ramp | Artifact-based mana development.
- 1 Orazca Relic | ramp | Mana acceleration with later utility.
- 1 Pristine Talisman | ramp | Ramp piece that fits the lifegain plan.
- 1 Spoils of Evil | ramp | Efficient burst of mana when available.
- 1 The Celestus | ramp | Color fixing and mana development in one card.
- 1 Angel of Indemnity | ramp | A lifegain-adjacent creature that supplies ramp.
- 1 Battle Angels of Tyr | ramp | Creature-based ramp that also adds board presence.
- 1 Archivist of Oghma | draw | Efficient creature-based card advantage.
- 1 Dawn of Hope | draw | Repeatable card-advantage outlet for the lifegain plan.
- 1 Enduring Innocence | draw | Persistent draw source on a relevant permanent.
- 1 Exemplar of Light | draw | Creature that provides card advantage.
- 1 Mangara, the Diplomat | draw | Reliable multiplayer card-advantage piece.
- 1 Markov Purifier | draw | Vampire-aligned draw for the deck's creature core.
- 1 Sigarda's Splendor | draw | Lifegain-themed source of cards.
- 1 The Gaffer | draw | Creature-based draw that rewards the deck's plan.
- 1 Tymna the Weaver | draw | Low-cost ongoing card advantage.
- 1 Well of Lost Dreams | draw | Converts lifegain into additional cards.
- 1 Vampiric Rites | draw | Flexible sacrifice-based draw outlet.
- 1 Aetherflux Reservoir | removal | Lifegain payoff that serves as a removal option.
- 1 Ayli, Eternal Pilgrim | removal | Low-cost removal creature that fits the deck's colors.
- 1 Murderous Rider // Swift End | removal | Flexible creature and removal card.
- 1 Nightmare's Thirst | removal | Cheap black removal.
- 1 Solitude | removal | Creature-based removal for problematic permanents.
- 1 Umezawa's Jitte | removal | Reusable equipment-based removal.
- 1 Vona, Butcher of Magan | removal | Lifegain-compatible creature removal.
- 1 Witch of the Moors | removal | A lifegain-themed removal threat.
- 1 Vein Ripper | removal | High-impact Vampire removal option.
- 1 Ajani's Pridemate | synergy | Classic lifegain payoff that grows into pressure.
- 1 Cleric Class | synergy | Low-cost enchantment support for the lifegain strategy.
- 1 Heliod, Sun-Crowned | synergy | Central lifegain synergy permanent.
- 1 Sanguine Bond | synergy | Major lifegain payoff enchantment.
- 1 Vito, Thorn of the Dusk Rose | synergy | Lifegain payoff within the Vampire creature base.
- 1 Vizkopa Guildmage | synergy | Flexible lifegain synergy creature.
- 1 Voice of the Blessed | synergy | Efficient creature that rewards repeated lifegain.
- 1 Resplendent Angel | synergy | Lifegain payoff that adds an evasive board presence.
- 1 Righteous Valkyrie | synergy | Angel and lifegain synergy for the creature plan.
- 1 Serra Ascendant | synergy | Low-cost lifegain-themed threat and payoff.
- 1 Archangel of Thune | threat | Powerful Angel threat for a lifegain-focused board.
- 1 Attended Healer | threat | Board-building threat that fits the Cleric theme.
- 1 Bloodbond Vampire | threat | Creature threat that rewards the deck's lifegain direction.
- 1 Doctor Jane Foster | threat | Legendary creature threat for the board.
- 1 Gideon's Company | threat | Creature threat suited to a lifegain shell.
- 1 Nykthos Paragon | threat | Lifegain payoff that can turn a board into pressure.
- 1 Radiant Solar | threat | Efficient Angel threat.
- 1 Rhox Faithmender | threat | Lifegain-focused creature that demands an answer.
- 1 Star Charter | threat | Low-cost Bat Cleric threat.
- 1 Twinblade Paladin | threat | Creature threat that benefits from the deck's plan.
- 1 Vampire Scrivener | threat | Vampire threat for the black creature core.
- 1 Wax-Wane Witness | threat | Creature threat that supports the lifegain theme.
- 1 Amalia Benavides Aguirre | wipe | Low-cost lifegain-connected board reset.
- 1 The Battle of Bywater | wipe | Efficient white board wipe.
- 1 Kaya's Wrath | wipe | Reliable full board reset in the deck's colors.
- 1 Alseid of Life's Bounty | interaction | Cheap protection for Karlov or a key permanent.
- 1 Bofur, Reliable Guardian // Concerted Care | interaction | Flexible protective interaction.
- 1 Courageous Resolve | interaction | Protective instant-speed interaction.
- 1 Faith's Shield | interaction | Low-cost protection for key threats.
- 1 Distinguished Conjurer | interaction | Creature-based interaction that fits the white core.
- 1 Metropolis Reformer | interaction | Protective Angel Cleric interaction.
- 1 Restoration Magic | interaction | Flexible instant-speed protection.
- 1 Werefox Bodyguard | interaction | Creature-based interactive utility.

</details>

### 2. aristocrats Commander, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 179 names.

Commander: Denethor, Ruling Steward.

Grade: bad, score 0.37, model 20260907T081149Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $145.09 the whole deck.

**Summary:** This Denethor aristocrats deck develops a low-cost creature and artifact board, uses its synergy pieces to turn routine exchanges into lasting advantage, and keeps cards flowing while answering key opposing permanents. It wins by steadily converting its creature-focused board into pressure and value after trades or board resets. The mana base is deliberately basic-heavy for dependable white and black access, giving up utility-land flexibility for consistent early development.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=no: The list mixes a commander-protection/equipment voltron package (Champion's Helm, Darksteel Plate, Bastion Protector, Lightning Greaves, Swiftfoot Boots, Bronze Guardian), a heavy removal/five-board-wipe control shell, and a handful of small value creatures, with no sacrifice outlets or token engines to tie any of it together — the pieces actively fight each other, since the wipes undercut the creature board the summary says it builds.
- PLAN theme_fit=no: The request was aristocrats at bracket 3, and while Denethor is a plausible white-black choice, the deck contains essentially no sacrifice outlets, token generation, or death-trigger payoffs — it answers a completely different request (protection/control) than the one asked.
- PLAN useful_as_built=partly: The 36-basic mana base and ample ramp will function and the deck can cast its spells, but with a handful of small creatures, five sweepers that hit its own board, and no real win condition beyond incidental beatdown, it is very hard to actually close a game with as it stands.
- PLAN summary_honest=no: It calls itself an aristocrats deck that converts a creature board into pressure, but there are almost no creatures, no sacrifice outlets, no death triggers or drain payoffs, and the 'answering key permanents' framing hides that this is really a removal-and-wipe pile with an equipment protection suite.
- [INFO] `curve_summary`: average mana value 2.92 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Buster Sword | draw | A draw slot that helps keep resources flowing.
- 1 Call of the Ring | draw | A black draw option for sustained card access.
- 1 Cirith Ungol Patrol | draw | A draw creature that contributes to the board.
- 1 Exemplar of Light | draw | A white draw creature for the deck's creature-focused plan.
- 1 Folk Hero | draw | A draw engine for a creature-heavy game plan.
- 1 Grave Venerations | draw | A black draw piece that supports longer games.
- 1 Idol of Oblivion | draw | A low-cost draw artifact.
- 1 Inspiring Overseer | draw | A draw creature that adds to the board.
- 1 Lembas | draw | A draw artifact that fits the low curve.
- 1 Mask of Memory | draw | An equipment-based draw option for creature combat.
- 1 Night's Whisper | draw | An efficient black draw spell.
- 1 Puresteel Paladin | draw | A draw creature that works alongside the deck's equipment.
- 1 Skullclamp | draw | A compact draw artifact for the aristocrats shell.
- 1 Tome of Legends | draw | A repeatable draw artifact for the commander deck.
- 1 Bastion Protector | interaction | A creature-based interaction slot.
- 1 Boromir, Warden of the Tower | interaction | A legendary creature providing interaction.
- 1 Bronze Guardian | interaction | An artifact creature interaction piece.
- 1 Champion's Helm | interaction | An equipment interaction piece for the commander plan.
- 1 Clever Concealment | interaction | A white interaction spell.
- 1 Darksteel Plate | interaction | An equipment-based interaction option.
- 1 Duty Beyond Death | interaction | A black interaction spell for protecting the board plan.
- 1 Frontline Medic | interaction | A creature interaction option.
- 1 Gift of Immortality | interaction | An interaction aura suited to a creature-focused deck.
- 1 Lightning Greaves | interaction | A low-cost equipment interaction piece.
- 1 Swiftfoot Boots | interaction | An equipment interaction option for key creatures.
- 1 Arcane Signet | ramp | Reliable color fixing and acceleration.
- 1 Bender's Waterskin | ramp | A low-cost ramp artifact.
- 1 Blitzball | ramp | An artifact ramp slot that maintains the curve.
- 1 Chromatic Lantern | ramp | Artifact acceleration and color support.
- 1 Coin of Mastery | ramp | An artifact ramp piece.
- 1 Commander's Sphere | ramp | Color-fixing artifact acceleration.
- 1 Deadly Dispute | ramp | A black ramp spell that suits the aristocrats direction.
- 1 Fellwar Stone | ramp | Efficient artifact mana fixing.
- 1 Inherited Envelope | ramp | An artifact ramp option.
- 1 Lotho, Corrupt Shirriff | ramp | A creature ramp piece in Denethor's colors.
- 1 Relic of Legends | ramp | A mana artifact that supports the legendary commander deck.
- 1 Sol Ring | ramp | Efficient colorless acceleration.
- 1 Thought Vessel | ramp | A mana artifact for stable development.
- 1 Angel of Serenity | removal | A creature-based removal option.
- 1 Banishing Light | removal | A white removal enchantment.
- 1 Bitter Triumph | removal | A flexible black removal spell.
- 1 Claim the Precious | removal | A black removal spell.
- 1 Contagion Clasp | removal | Artifact-based removal.
- 1 Crib Swap | removal | A white removal instant.
- 1 Dismember | removal | Efficient creature removal.
- 1 Dispatch | removal | A low-cost white removal spell.
- 1 Fiend Hunter | removal | A creature-based removal piece.
- 1 Generous Gift | removal | Broad white removal.
- 1 Get Lost | removal | A flexible white removal spell.
- 1 Heartless Act | removal | An efficient black removal instant.
- 1 Archfiend of Ifnir | wipe | A black creature that occupies a wipe slot.
- 1 Black Sun's Zenith | wipe | A scalable black board wipe.
- 1 Dusk // Dawn | wipe | A white board wipe with added flexibility.
- 1 Split Up | wipe | A low-curve white wipe option.
- 1 Vanquish the Horde | wipe | A white board wipe for resetting crowded boards.
- 1 Arcade Cabinet | synergy | A colorless synergy piece for the aristocrats plan.
- 1 Gollum the Abandoned | synergy | A black synergy creature.
- 1 Gollum, Patient Plotter | synergy | A black synergy creature for the deck's core plan.
- 1 Gríma Wormtongue | synergy | A black legendary synergy piece.
- 1 Heirloom Auntie | synergy | A black synergy creature.
- 1 Joo Dee, One of Many | synergy | A white synergy creature.
- 1 Nimble Hobbit | synergy | A low-curve white synergy creature.
- 1 Phantom Train | synergy | A colorless synergy artifact.
- 18 Plains | land | Primary white mana source.
- 18 Swamp | land | Primary black mana source.

</details>

### 3. artifacts Commander, bracket 4

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 299 names.

Commander: Urza, Lord High Artificer.

Grade: bad, score 0.37, model 20260907T081149Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $4012.80 to buy, $4012.80 the whole deck.

**Summary:** This deck develops a dense artifact board quickly, keeps cards flowing, and uses efficient interaction and removal to protect its momentum. Artifact creatures provide the primary pressure, while the artifact-synergy pieces reinforce the central plan. It gives up some broad reset capacity in exchange for speed, focused development, and a proactive board presence.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core is a consistent mono-blue artifact plan (cost reducers, Metalworker/Powerstone ramp, Urza construct fodder, artifact creatures), but a meaningful slice of the list wanders — Twiddle, Reality Ripple, Ghostly Flicker/Escape Protocol blink with almost nothing worth blinking, Water Whip, Splitting the Powerstone, Luck Bobblehead — filler that serves no single line and dilutes the artifact engine.
- PLAN theme_fit=partly: It is genuinely a mono-blue Urza artifact deck, but it does not answer the "high-power, bracket 4" part of the request: no fast mana beyond Mox Opal and Workshop, essentially one tutor, no Urza combo package or Thopter/Sword-style finish, and dozens of draft-chaff slots that sit far below bracket 4.
- PLAN useful_as_built=yes: The mana base has ample blue sources and a reasonable land count, the curve is low and supported by cost reducers and Powerstone ramp, and there are clear ways to close — Urza's construct token, Traxos/Lodestone/Metamorph beats and Skysovereign — so the deck functions at the table even if several tribal lands (Cavern, Unclaimed Territory, Secluded Courtyard) are near-dead in a monocolor list.
- PLAN summary_honest=partly: It fairly discloses the low quality-model grade and the thin sweeper count, but calling the removal suite "efficient interaction and removal" oversells a package built on Unable to Scream, Water Whip, Watery Grasp, Cyber Conversion and Twiddle, and it implies a speed the list's clunky ramp does not really deliver.
- [INFO] `curve_summary`: average mana value 3.32 over 66 nonland cards
- [WARN] `profile_off_band`: the mana available on turn four is 4.68, and bracket 4 wants 4.8 or more (mean over 10000 hands)
- [INFO] `mana_pass`: the builder moved 9 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 15 Island | land | Reliable blue mana base for the artifact plan.
- 1 Academy Ruins | land | Utility land slot in the artifact-focused mana base.
- 1 Buried Ruin | land | Utility land slot for the artifact-focused shell.
- 1 Inventors' Fair | land | Artifact-oriented utility land.
- 1 Mishra's Workshop | land | Colorless artifact-focused mana land.
- 1 Urza's Saga | land | Artifact-focused utility land.
- 1 Seat of the Synod | land | Blue artifact land for the mana base.
- 1 Mystic Sanctuary | land | Blue utility land.
- 1 Otawara, Soaring City | land | Blue legendary utility land.
- 1 Command Tower | land | Reliable blue source.
- 1 Exotic Orchard | land | Flexible blue-source land.
- 1 City of Brass | land | Flexible untapped blue source.
- 1 Mana Confluence | land | Flexible untapped blue source.
- 1 Glimmervoid | land | Artifact-friendly blue source.
- 1 Spire of Industry | land | Artifact-friendly flexible source.
- 1 Cavern of Souls | land | Creature-supporting blue-source land.
- 1 Secluded Courtyard | land | Creature-supporting blue-source land.
- 1 Unclaimed Territory | land | Creature-supporting blue-source land.
- 1 Plaza of Heroes | land | Legendary-permanent utility land.
- 1 Archway of Innovation | land | Artifact-plan utility land.
- 1 Mistrise Village | land | Blue utility land.
- 1 Reflecting Pool | land | Flexible mana-base support.
- 1 Thirst for Knowledge | draw | Dedicated card-flow slot.
- 1 Riddlesmith | draw | Artifact-plan draw support.
- 1 Sai, Master Thopterist | draw | Artifact-focused draw engine.
- 1 Vedalken Archmage | draw | Artifact-focused draw support.
- 1 Forensic Gadgeteer | draw | Artifact-plan card-flow piece.
- 1 Braided Net // Braided Quipu | draw | Artifact draw slot.
- 1 One with the Machine | draw | Artifact-plan draw payoff.
- 1 Metallic Rebuke | interaction | Efficient artifact-shell interaction.
- 1 Disruption Protocol | interaction | Low-cost interaction for protecting the plan.
- 1 Ice Out | interaction | Flexible interaction slot.
- 1 Override | interaction | Dedicated stack-interaction slot.
- 1 Stoic Rebuttal | interaction | Artifact-friendly interaction spell.
- 1 Reality Ripple | interaction | Flexible instant-speed interaction.
- 1 Ghostly Flicker | interaction | Flexible protection and interaction piece.
- 1 Welding Jar | interaction | Artifact protection interaction.
- 1 Padeem, Consul of Innovation | interaction | Artifact-focused protective interaction.
- 1 Curator's Ward | interaction | Permanent-protection interaction slot.
- 1 Escape Protocol | interaction | Artifact-plan interaction support.
- 1 Mox Opal | ramp | Fast artifact-compatible mana development.
- 1 Moonsnare Prototype | ramp | Low-cost artifact ramp.
- 1 Vedalken Engineer | ramp | Artifact-focused mana development.
- 1 Chief Engineer | ramp | Artifact-focused mana development.
- 1 Grand Architect | ramp | Artifact-focused mana development.
- 1 Coin of Mastery | ramp | Artifact ramp slot.
- 1 Crystal Skull, Isu Spyglass | ramp | Artifact ramp support.
- 1 Network Terminal | ramp | Artifact ramp slot.
- 1 Inspiring Statuary | ramp | Artifact-plan mana support.
- 1 Metalworker | ramp | High-impact artifact ramp piece.
- 1 Slagstone Refinery | ramp | Artifact ramp support.
- 1 Aether Spellbomb | removal | Cheap artifact-based removal.
- 1 Resculpt | removal | Flexible removal spell.
- 1 Unable to Scream | removal | Low-cost permanent removal.
- 1 Cyber Conversion | removal | Dedicated removal slot.
- 1 Ravenform | removal | Flexible removal spell.
- 1 Water Whip | removal | Dedicated removal slot.
- 1 Watery Grasp | removal | Permanent-removal support.
- 1 Kitesail Larcenist | removal | Creature-based removal option.
- 1 Lumengrid Drake | removal | Creature-based removal option.
- 1 Skysovereign, Consul Flagship | removal | Artifact threat that provides removal.
- 1 Emry, Lurker of the Loch | synergy | Core artifact-synergy permanent.
- 1 Etherium Sculptor | synergy | Artifact-synergy support piece.
- 1 Foundry Inspector | synergy | Artifact-synergy support piece.
- 1 Mystic Forge | synergy | Artifact-plan synergy engine.
- 1 Clock of Omens | synergy | Artifact-focused synergy piece.
- 1 Whir of Invention | synergy | Artifact-plan synergy spell.
- 1 Frogmite | threat | Low-curve artifact threat.
- 1 Chrome Steed | threat | Artifact creature pressure.
- 1 Traxos, Scourge of Kroog | threat | Efficient artifact threat.
- 1 Lodestone Golem | threat | Artifact creature threat.
- 1 Jhoira's Familiar | threat | Artifact creature pressure.
- 1 Master Transmuter | threat | Artifact-focused threat.
- 1 Phyrexian Metamorph | threat | Flexible artifact creature threat.
- 1 Arcbound Crusher | threat | Artifact creature threat.
- 1 Synth Infiltrator | threat | Artifact creature threat.
- 1 Whirler Rogue | threat | Creature threat for the artifact plan.
- 1 Engineered Explosives | wipe | Flexible artifact board reset.
- 1 Hurkyl's Recall | wipe | Broad reset option against developed boards.
- 1 Cargo Ship | ramp | the mana pass added it to bring the mana base inside the power level
- 1 Twiddle | draw | the mana pass added it to bring the mana base inside the power level
- 1 Luck Bobblehead | ramp | the mana pass added it to bring the mana base inside the power level
- 1 Powerstone Shard | ramp | the mana pass added it to bring the mana base inside the power level
- 1 Solar Array | ramp | the mana pass added it to bring the mana base inside the power level
- 1 Splitting the Powerstone | ramp | the mana pass added it to bring the mana base inside the power level

</details>

### 4. dinosaur tribal, bracket 2

Format: Commander. Theme: dinosaurs. Pool: any_card. Shortlist: 269 names.

Commander: Gishath, Sun's Avatar.

Grade: bad, score 0.05, model 20260907T081149Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $678.23 to buy, $678.23 the whole deck.

**Summary:** This deck develops its mana early, builds a board of dinosaurs, and uses Gishath as the centerpiece of a straightforward combat plan. Large dinosaur threats provide the main way to finish games, while dinosaur-based removal, a few sweepers, and protective interaction help the board survive long enough to attack. It gives up the speed and precision of a highly tuned list in exchange for a clear tribal theme, resilient creature play, and splashy late-game turns.

The quality model grades this deck below the precon baseline against the top lists of the format: more opening hands hold two to four lands than the norm, and that raises the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. The deck holds no fast mana, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Every slot points at one plan: ramp into dinosaurs, cast Gishath, swing and cheat more dinosaurs in, supported by tribal payoffs (Vanquisher's Banner, Garruk's Uprising, Pantlaza) and dino-based removal.
- PLAN theme_fit=yes: It is a mono-dinosaur Gishath Commander deck with no game changers or combo lines, sitting comfortably in bracket 2 and using mostly simple attack-and-block cards suited to a newer player.
- PLAN useful_as_built=yes: The five-color-ish Naya mana base is deep and well-fixed with ten ramp pieces, and there are numerous large creatures plus trample/double-strike enablers to close games, so it functions out of the box despite a very land-heavy top-end curve.
- PLAN summary_honest=yes: The claims — early mana development, dinosaur removal (Trumpeting Carnosaur, Savage Stomp, Thrashing Brontodon), a few sweepers (Blasphemous Act, Austere Command, Raging Swordtooth), protective interaction (Heroic Intervention, Akroma's Will, boots/greaves) and big finishers — all match the list, and it openly notes the lack of fast mana.
- [INFO] `curve_summary`: average mana value 3.85 over 62 nonland cards

<details><summary>The deck list</summary>

- 8 Forest | land | Basic green source for the dinosaur-heavy mana base.
- 5 Mountain | land | Basic red source for the mana base.
- 5 Plains | land | Basic white source for the mana base.
- 1 Command Tower | land | Reliable multicolor land for the three-color deck.
- 1 City of Brass | land | Flexible source for any needed color.
- 1 Mana Confluence | land | Flexible source for any needed color.
- 1 Exotic Orchard | land | Multicolor fixing for a multiplayer table.
- 1 Reflecting Pool | land | Supports the deck's broad range of colored sources.
- 1 Jetmir's Garden | land | Three-color land that supports all of Gishath's colors.
- 1 Canopy Vista | land | Green-white dual land for the mana base.
- 1 Cinder Glade | land | Green-red dual land for the mana base.
- 1 Sacred Foundry | land | Red-white dual land for the mana base.
- 1 Temple Garden | land | Green-white dual land for the mana base.
- 1 Stomping Ground | land | Green-red dual land for the mana base.
- 1 Taiga | land | Green-red dual land for the mana base.
- 1 Bountiful Promenade | land | Green-white fixing for multiplayer games.
- 1 Brushland | land | Green-white source that helps early development.
- 1 Battlefield Forge | land | Red-white source that supports the deck's spells.
- 1 Karplusan Forest | land | Green-red source that supports the deck's spells.
- 1 Spire Garden | land | Green-red fixing for multiplayer games.
- 1 Sunpetal Grove | land | Green-white dual land for the mana base.
- 1 Three Tree City | land | Tribal land that supports the dinosaur-focused deck.
- 1 Sol Ring | ramp | Efficient early mana to help deploy the deck's expensive dinosaurs.
- 1 Arcane Signet | ramp | Reliable three-color mana acceleration.
- 1 Fellwar Stone | ramp | Early mana rock that helps fix colors in multiplayer.
- 1 Birds of Paradise | ramp | Low-cost creature ramp that fixes every deck color.
- 1 Llanowar Elves | ramp | Low-cost green acceleration for early turns.
- 1 Drover of the Mighty | ramp | Creature-based ramp that fits the dinosaur plan.
- 1 Nature's Lore | ramp | Efficient land ramp that improves color access.
- 1 Three Visits | ramp | Efficient land ramp that improves color access.
- 1 Farseek | ramp | Early land ramp for the three-color mana base.
- 1 Rampant Growth | ramp | Straightforward land ramp for a new-player-friendly plan.
- 1 Beast Whisperer | draw | Creature-focused draw that rewards casting the deck's many dinosaurs.
- 1 Garruk's Uprising | draw | Draw support suited to a deck of large creatures.
- 1 Guardian Project | draw | Ongoing draw for the creature-heavy strategy.
- 1 Kutzil, Malamet Exemplar | draw | Creature-based draw support for the combat plan.
- 1 Ripjaw Raptor | draw | Dinosaur draw option that remains a meaningful board presence.
- 1 Earthshaker Dreadmaw | draw | Large dinosaur that provides card advantage.
- 1 Return of the Wildspeaker | draw | Flexible draw spell for a deck built around large creatures.
- 1 Rishkar's Expertise | draw | High-impact card advantage for a battlefield with large dinosaurs.
- 1 Harmonize | draw | Simple, dependable card draw.
- 1 Toski, Bearer of Secrets | draw | Combat-oriented draw support for an attacking creature deck.
- 1 Vanquisher's Banner | draw | Tribal draw support for the dinosaur creature type.
- 1 Akroma's Will | interaction | Flexible interaction for protecting or forcing through the creature plan.
- 1 Boros Charm | interaction | Versatile protection and interaction option.
- 1 Heroic Intervention | interaction | Protection for the developed dinosaur board.
- 1 Lightning Greaves | interaction | Protects an important creature while supporting the attack plan.
- 1 Swiftfoot Boots | interaction | Repeatable creature protection for key dinosaurs.
- 1 Temple Altisaur | interaction | Dinosaur interaction that helps support the creature board.
- 1 Bronzebeak Foragers | removal | Dinosaur-based removal that advances the tribal plan.
- 1 Itzquinth, Firstborn of Gishath | removal | Low-cost dinosaur removal for early board control.
- 1 Ravenous Sailback | removal | Dinosaur removal attached to a sturdy creature.
- 1 Savage Stomp | removal | Efficient dinosaur-themed removal.
- 1 Thrashing Brontodon | removal | Flexible dinosaur removal on a creature body.
- 1 Trumpeting Carnosaur | removal | Large dinosaur that also supplies removal.
- 1 Tranquil Frillback | removal | Versatile dinosaur removal for troublesome permanents.
- 1 Blasphemous Act | wipe | Reset button when opposing creatures overwhelm the board.
- 1 Austere Command | wipe | Flexible board wipe that can be chosen around the current battlefield.
- 1 Raging Swordtooth | wipe | Dinosaur-themed sweep effect that keeps the tribal focus.
- 1 Commune with Dinosaurs | synergy | Early dinosaur synergy that helps set up the creature plan.
- 1 Kinjalli's Caller | synergy | Supports deploying the deck's expensive dinosaurs.
- 1 Otepec Huntmaster | synergy | Dinosaur-focused support for the aggressive game plan.
- 1 Marauding Raptor | synergy | Early dinosaur synergy that rewards creature development.
- 1 Kinjalli's Sunwing | synergy | Dinosaur support that improves the deck's board presence.
- 1 Huatli's Raptor | synergy | Low-cost dinosaur that contributes to tribal development.
- 1 Huatli's Snubhorn | synergy | Dinosaur support creature for the tribal shell.
- 1 Hunting Velociraptor | synergy | Dinosaur synergy that supports the creature-heavy plan.
- 1 Orazca Frillback | synergy | Flexible dinosaur synergy for the midgame.
- 1 Priest of the Wakening Sun | synergy | Early tribal support for the dinosaur strategy.
- 1 Raptor Companion | synergy | Low-cost dinosaur that establishes the tribal board.
- 1 Relentless Raptor | synergy | Efficient dinosaur body for the attacking plan.
- 1 Territorial Hammerskull | synergy | Dinosaur synergy that supports combat-focused turns.
- 1 Ancient Brontodon | threat | Large dinosaur threat that helps close games through combat.
- 1 Carnage Tyrant | threat | Powerful dinosaur threat for the top end.
- 1 Etali, Primal Storm | threat | High-impact dinosaur threat for the late game.
- 1 Ghalta and Mavren | threat | Legendary dinosaur threat that strengthens the combat plan.
- 1 Ghalta, Primal Hunger | threat | Massive dinosaur finisher for a creature-heavy board.
- 1 Ghalta, Stampede Tyrant | threat | Large dinosaur threat that rewards building a big board.
- 1 Goring Ceratops | threat | Combat-focused dinosaur threat for attacking turns.
- 1 Pantlaza, Sun-Favored | threat | Dinosaur threat that reinforces the tribal centerpiece.
- 1 Quartzwood Crasher | threat | Combat-focused dinosaur threat that pressures opponents.
- 1 Regisaur Alpha | threat | Dinosaur threat that helps the deck turn creatures sideways.
- 1 Tyrranax Rex | threat | Large dinosaur finisher for the late game.
- 1 Zetalpa, Primal Dawn | threat | Evasive dinosaur threat that can finish games in combat.

</details>

### 5. blink Commander, owned first

Format: Commander. Theme: blink. Pool: owned_first. Shortlist: 168 names.

Commander: Gilraen, Dúnedain Protector.

Grade: bad, score 0.38, model 20260907T081149Z.

Cards: 99 main, 0 sideboard. Repair turn: yes, for deck_size, copy_limit, profile_off_band, profile_off_band. Block findings: 0.

Cost: $0.00 to buy, $178.49 the whole deck.

**Summary:** This mono-white blink deck develops a durable creature board, repeatedly reuses enter-the-battlefield effects for cards and removal, and protects its most important permanents. Equipment, resilient legends, and evasive Angels provide a steady combat finish after the deck has controlled the battlefield.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The list gathers ETB value creatures, removal, protection gear, and a wipe package around a protection-blink commander, but the many hexproof/shroud equipment pieces (Champion's Helm, Swiftfoot Boots, Lightning Greaves) work against a commander whose payoff comes from being targeted, and the equipment/legends/Angels subthemes pull attention away from a tight blink engine.
- PLAN theme_fit=partly: It is mono-white Commander at a reasonable bracket-3 power with a blink-flavored commander, but the actual blink count is thin (Flickerwisp, Angel of Condemnation, Fiend Hunter, Angel of Sanctions, Personify) next to a large protection/equipment and removal shell, so it reads more as a white protection-value deck than the blink deck asked for.
- PLAN useful_as_built=yes: 37 lands with mostly Plains, ten ramp pieces, plenty of removal, card draw, and a real combat finish in Angels and equipped legends make it a fully playable pick-up-and-go list, even though multicolor fixing like City of Brass, Grand Coliseum, and Chromatic Lantern is wasted in a mono-white deck.
- PLAN summary_honest=partly: placeholder
- [INFO] `curve_summary`: average mana value 2.95 over 63 nonland cards

<details><summary>The deck list</summary>

- 28 Plains | land | Reliable white mana base.
- 1 Command Tower | land | Untapped white source.
- 1 City of Brass | land | Flexible untapped color source.
- 1 Exotic Orchard | land | Flexible white-producing land in multiplayer.
- 1 Grand Coliseum | land | Additional color fixing.
- 1 Minas Tirith | land | White source with useful card advantage potential.
- 1 Plaza of Heroes | land | White source that can protect legendary permanents.
- 1 Secluded Courtyard | land | Human-focused white fixing.
- 1 Fabled Passage | land | Finds a Plains while smoothing draws.
- 1 Sol Ring | ramp | Efficient colorless acceleration.
- 1 Arcane Signet | ramp | Reliable commander-color acceleration.
- 1 Bender's Waterskin | ramp | Early mana development.
- 1 Fellwar Stone | ramp | Efficient multiplayer mana rock.
- 1 Wayfarer's Bauble | ramp | Turns early mana into a permanent Plains.
- 1 Sword of the Animist | ramp | Equipment-based land ramp for attacking creatures.
- 1 Relic of Legends | ramp | Fixes mana and rewards legendary creatures.
- 1 Thought Vessel | ramp | Mana rock with hand-size utility.
- 1 Commander's Sphere | ramp | Mana fixing that can become a card later.
- 1 Chromatic Lantern | ramp | Broad mana fixing and acceleration.
- 1 Adventurer's Airship | draw | Repeatable source of cards from combat.
- 1 Buster Sword | draw | Equipment that supplies ongoing card advantage.
- 1 Champions of Minas Tirith | draw | Creature-based card advantage.
- 1 Crown of Gondor | draw | Equipment that supports steady card flow.
- 1 Diary of Dreams | draw | Artifact card-advantage engine.
- 1 Idol of Oblivion | draw | Low-cost recurring draw.
- 1 Inspiring Overseer | draw | An evasive enter-the-battlefield draw creature.
- 1 Lembas | draw | Immediate card value that can be reused.
- 1 Mask of Memory | draw | Combat-based filtering and card advantage.
- 1 Skullclamp | draw | Efficient creature-based card draw.
- 1 Wall of Omens | draw | Cheap defensive enter-the-battlefield cantrip.
- 1 Bastion Protector | interaction | Keeps the commander protected while developing the board.
- 1 Champion's Helm | interaction | Protects the commander and other legends.
- 1 Clever Concealment | interaction | Protects a developed board from opposing answers.
- 1 Darksteel Plate | interaction | Durable protection for a key creature.
- 1 Lightning Greaves | interaction | Gives key creatures haste and protection.
- 1 Reprieve | interaction | Temporarily stops a problematic spell while replacing itself.
- 1 Slip On the Ring | interaction | Protects a creature while retriggering its enter-the-battlefield ability.
- 1 Swiftfoot Boots | interaction | Reliable haste and hexproof protection.
- 1 Angel of Sanctions | removal | Creature-based exile removal that remains valuable when reused.
- 1 Banishing Light | removal | Flexible answer to troublesome nonlands.
- 1 Crib Swap | removal | Instant-speed creature exile.
- 1 Dispatch | removal | Efficient creature answer alongside artifacts.
- 1 Fiend Hunter | removal | Enter-the-battlefield removal that works well with blink effects.
- 1 Generous Gift | removal | Answers any permanent at instant speed.
- 1 Get Lost | removal | Efficient answer to creatures, enchantments, and planeswalkers.
- 1 Palace Jailer | removal | Removal creature that also creates a card-advantage incentive.
- 1 Swords to Plowshares | removal | Premium low-cost creature exile.
- 1 Austere Command | wipe | Flexible reset that can spare important parts of the board.
- 1 Dusk // Dawn | wipe | Creature reset with a creature-recursion aftermath.
- 1 Vanquish the Horde | wipe | Efficient creature-board reset.
- 1 Angel of Condemnation | synergy | Repeatably blinks value creatures and can answer opposing attackers.
- 1 Angel of Serenity | synergy | Provides a powerful enter-the-battlefield effect to reuse.
- 1 Duty Beyond Death | synergy | Preserves an important creature for another enter-the-battlefield trigger.
- 1 Ennis, Debate Moderator | synergy | Supports the creature-centered value plan.
- 1 Flickerwisp | synergy | Efficient blink creature for reusing value permanents.
- 1 Gift of Immortality | synergy | Returns a key value creature after it dies.
- 1 Jocasta, Automaton Avenger | synergy | Supports the deck's recurring creature-value plan.
- 1 Personify | synergy | Blink effect for protecting creatures and reusing triggers.
- 1 Sheltered by Ghosts | synergy | Protects a value creature while supporting repeated creature deployment.
- 1 Together Forever | synergy | Keeps key creatures available for repeated value.
- 1 Exemplar of Light | threat | Evasive creature that pressures opponents while providing value.
- 1 Faramir, Field Commander | threat | Legendary creature that builds a meaningful board presence.
- 1 Frontline Medic | threat | Combat-focused creature that protects an attacking board.
- 1 Giada, Font of Hope | threat | Develops Angels into a substantial evasive force.
- 1 Kataki, War's Wage | threat | Creature pressure that disrupts artifact-heavy opponents.
- 1 Puresteel Paladin | threat | Efficient creature that turns equipment into a battlefield advantage.
- 1 South Pole Voyager | threat | Creature that contributes pressure and card value.
- 1 The Vision | threat | Resilient artifact creature with strong battlefield presence.
- 1 The Walls of Ba Sing Se | threat | Large defensive artifact creature that can become a win condition.
- 1 Troop of Ponies | threat | Creature-based development that contributes to the board.
- 1 Westfold Rider | threat | Attacking creature with useful removal utility.
- 1 Zack Fair | threat | Legendary creature that adds pressure and supports the board.

</details>

### 6. Modern tempo, tournament

Format: Modern. Theme: tempo. Pool: any_card. Shortlist: 167 names.

Grade: typical, score 0.46, model 20260907T081149Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $324.38 to buy, $324.38 the whole deck.

**Summary:** This is a blue-red tempo deck that establishes pressure with Ragavan, Nimble Pilferer, Ledger Shredder, and Faerie Mastermind, then protects that pressure with efficient interaction while using removal to keep attacks clear. It wins by forcing the opponent to answer its threats under disruption, with temporal cards providing a higher-impact secondary angle. The deck gives up broad permanent coverage and long-game resilience in exchange for a streamlined, spell-heavy pressure plan.

The quality model grades this deck typical against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade. The deck runs its spells as playsets, and that lowers the grade.

- PLAN plan_coherent=partly: The core of cheap threats, counters, cantrips, and burn is a single clean tempo plan, but the four extra-turn spells (especially a Temporal Trespass that needs heavy delve fuel this list barely generates) pull toward a different, clunkier axis.
- PLAN theme_fit=partly: It is a Modern blue-red spell-heavy tempo deck as asked, but it contains no Delver of Secrets (or Murktide-style payoff) despite the request naming a Delver deck, and the extra-turn package is not part of that archetype.
- PLAN useful_as_built=yes: 24 lands with duals and fetch-free untapped sources support a curve topping at three real mana, and twelve evasive threats plus Bolt give plenty of ways to close, so it can be picked up and played despite the awkward Temporal Trespass copies.
- PLAN summary_honest=partly: It fairly describes the threat suite, interaction, and lack of long-game resilience, but calling the extra-turn spells a "higher-impact secondary angle" oversells cards that are frequently uncastable here, and it never notes that the requested Delver is absent.
- [INFO] `curve_summary`: average mana value 2.22 over 36 nonland cards

<details><summary>The deck list</summary>

- 10 Island | land | Provides a reliable blue mana base for the deck’s blue spells.
- 6 Mountain | land | Provides red mana for the deck’s red spells.
- 4 Steam Vents | land | Supplies either deck color while counting as both an Island and a Mountain.
- 4 Spirebluff Canal | land | Adds flexible blue-red mana without using a forced tapped land.
- 4 Consider | draw | Efficient draw that helps the deck keep cards flowing.
- 2 Preordain | draw | Additional early draw for finding the needed part of the plan.
- 4 Counterspell | interaction | Core interaction for protecting pressure and disrupting opposing plays.
- 2 Spell Pierce | interaction | Cheap interaction that fits a tempo-oriented game plan.
- 4 Lightning Bolt | removal | Efficient removal that supports an aggressive tempo plan.
- 2 Into the Flood Maw | removal | Low-cost removal for clearing a path for threats.
- 2 Pongify | removal | Additional efficient removal for creatures that get in the way.
- 2 Temporal Mastery | synergy | Supports the deck’s temporal synergy package.
- 2 Temporal Trespass | synergy | Adds another temporal synergy card to the deck’s plan.
- 4 Ragavan, Nimble Pilferer | threat | Creature threat that gives the deck early pressure.
- 4 Ledger Shredder | threat | Creature threat that supports the deck’s pressure-focused plan.
- 4 Faerie Mastermind | threat | Flying creature threat that helps maintain pressure in the air.

</details>

### 7. Modern burn, casual

Format: Modern. Theme: burn. Pool: any_card. Shortlist: 290 names.

Grade: typical, score 0.44, model 20260907T081149Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $23.82 to buy, $23.82 the whole deck.

**Summary:** This mono-red burn deck applies pressure with direct-damage removal and a creature package, while spell-focused synergy turns repeated casting into more pressure. It aims to keep damage flowing and use its draw spells to avoid running out of action, finishing with resilient threats when the game goes longer. It gives up broad answers and flexibility for a focused, aggressive red plan.

The quality model grades this deck typical against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The color sources cover the pips, and that raises the grade. Many of the cards appear in no top list, and that lowers the grade.

- PLAN plan_coherent=no: The pieces fight each other: Thermo-Alchemist and Fuming Effigy want a dense instant/sorcery count that this deck (12 real spells) does not have, Lambholt Raconteur is a werewolf that wants you *not* to cast spells, and Hazoret wants an empty hand while Risk Factor and Ancestral Anger refill it.
- PLAN theme_fit=partly: It is mono-red, casual, and Modern-legal, but the request was burn and the list is a creature-aggro deck with only six actual burn spells, leaning on pingers rather than reach to the face.
- PLAN useful_as_built=partly: Twenty-four Mountains cast everything and the curve is low enough to function, but 24 lands is several too many for a two-drop-heavy list and the win path relies on mediocre creatures whose payoffs (Alchemist, Effigy, Hazoret) rarely turn on.
- PLAN summary_honest=partly: It is candid about giving up flexibility, but it oversells a "spell-focused synergy" and "draw spells" engine that a 12-spell, four-cantrip shell barely supports, and it hides the werewolf/Hazoret tension and the mislabeling of Sudden Breakthrough as ramp.
- [INFO] `curve_summary`: average mana value 2.56 over 36 nonland cards

<details><summary>The deck list</summary>

- 24 Mountain | land | Reliable mono-red mana base for the burn plan.
- 4 Ancestral Anger | draw | Low-cost card-flow slot that keeps spells coming.
- 2 Risk Factor | draw | Additional card-flow option for maintaining pressure.
- 2 Sudden Breakthrough | ramp | Provides the requested mana acceleration.
- 4 Lightning Bolt | removal | Efficient direct-damage removal for the core burn package.
- 2 Lava Dart | removal | Completes the low-cost removal suite.
- 4 Eidolon of the Great Revel | synergy | Spell-focused synergy piece that supports the aggressive plan.
- 4 Thermo-Alchemist | synergy | Spell-focused synergy threat that rewards repeated burn spells.
- 4 Akki Lavarunner // Tok-Tok, Volcano Born | threat | Early creature threat for applying pressure.
- 4 Fuming Effigy | threat | Creature threat that broadens the board presence.
- 4 Lambholt Raconteur // Lambholt Ravager | threat | Creature threat that fills out the aggressive package.
- 2 Hazoret the Fervent | threat | Top-end creature threat for closing games.

</details>

### 8. Modern lifegain, FNM

Format: Modern. Theme: lifegain. Pool: any_card. Shortlist: 301 names.

Grade: bad, score 0.28, model 20260907T081149Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $441.92 to buy, $441.92 the whole deck.

**Summary:** This is a white-black creature-based lifegain deck that establishes early synergy pieces, builds card advantage through its draw package, and turns those incremental gains into pressure from Ajani's Pridemate, Attended Healer, Twinblade Paladin, and larger black threats. Solitude and Murderous Rider keep key opposing permanents in check while the deck develops. It gives up some speed to play a fuller mana base and a higher-impact midgame, so its strongest games come from establishing its lifegain engine and keeping it on the table.

The quality model grades this deck below the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade.

- PLAN plan_coherent=yes: Every nonland card either gains life (Soul Warden, Attended Healer, Twinblade Paladin, Solitude, Murderous Rider) or rewards it (Pridemate, Enduring Innocence, Enduring Tenacity, Dawn of Hope, Sheoldred), so the pieces all feed one engine.
- PLAN theme_fit=yes: It is a Modern-legal white-black creature lifegain deck as asked, with a straightforward synergy shell appropriate to FNM play.
- PLAN useful_as_built=yes: 24 dual-heavy lands plus Altar support a curve that is admittedly four-and-five-drop heavy, but Soul Warden and Pridemate give early plays, Solitude can be evoked, and multiple creature payoffs plus Sheoldred provide clear ways to close a game.
- PLAN summary_honest=yes: Each named piece and role — draw package, lifegain payoffs, removal, top-heavy curve traded for a fuller mana base — is actually present in the list, and the summary openly reports the quality model's low grade.
- [INFO] `curve_summary`: average mana value 3.22 over 36 nonland cards

<details><summary>The deck list</summary>

- 4 Godless Shrine | land | Primary white-black dual land for reliable early color access.
- 4 Caves of Koilos | land | Untapped white-black source that supports the two-color curve.
- 4 Shattered Sanctum | land | Additional white-black dual source for the mana base.
- 6 Plains | land | Basic white source for the deck’s white spells.
- 6 Swamp | land | Basic black source for the deck’s black spells.
- 2 Altar of the Pantheon | ramp | Mana acceleration that also supports both deck colors.
- 4 Enduring Innocence | draw | Draw engine that rewards the deck’s creature-focused plan.
- 2 Dawn of Hope | draw | Repeatable card advantage for a lifegain-focused game plan.
- 4 Solitude | removal | Efficient removal that keeps opposing threats from taking over.
- 2 Murderous Rider // Swift End | removal | Flexible removal attached to a creature threat.
- 4 Soul Warden | synergy | Low-cost lifegain synergy piece that starts the plan early.
- 4 Ajani's Pridemate | synergy | Early payoff creature for repeated lifegain.
- 4 Attended Healer | threat | Midgame creature threat that fits the lifegain strategy.
- 4 Enduring Tenacity | threat | Resilient black threat for pressuring opponents.
- 4 Twinblade Paladin | threat | White threat that benefits from the deck’s lifegain plan.
- 2 Sheoldred, the Apocalypse | threat | High-impact threat that provides a powerful top end.

</details>

### 9. Standard midrange, FNM

Format: Standard. Theme: midrange. Pool: any_card. Shortlist: 164 names.

Grade: bad, score 0.32, model 20260907T081149Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $108.56 to buy, $108.56 the whole deck.

**Summary:** This black-green midrange deck builds a board around an Elf creature core, develops its mana, and keeps resources coming while using flexible removal to clear opposing problems. It wins by maintaining pressure with a steady stream of creature threats and turning a stable board into a combat advantage. The tradeoff is a focused, board-centric plan that has less room for narrow answers and specialized interaction.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The removal-plus-ramp-plus-fat midrange shell hangs together, but the advertised "Elf creature core" is only ten elves whose payoff (Elvish Archdruid) sits awkwardly beside a top end of non-Elf fatties, so the synergy slot pulls against the generic midrange plan.
- PLAN theme_fit=yes: It is a two-color black-green Standard midrange list at an unremarkable FNM power level, exactly what was asked for, with no extra restriction violated.
- PLAN useful_as_built=yes: 24 lands with dual lands in both colors, cheap ramp, eight removal spells, six card-draw pieces, and plenty of threats make it a functional pick-up-and-play deck despite a somewhat top-heavy five- and six-drop count.
- PLAN summary_honest=partly: The midrange, removal, and card-draw claims match the list, but calling the deck built "around an Elf creature core" oversells three Elf cards that do not meaningfully feed the Hydra/Vein Ripper/Brontodon top end.
- [INFO] `curve_summary`: average mana value 2.89 over 36 nonland cards

<details><summary>The deck list</summary>

- 6 Forest | land | Basic land slot for the mana base.
- 6 Swamp | land | Basic land slot for the mana base.
- 4 Overgrown Tomb | land | Dual land slot supporting the black-green mana base.
- 4 Blooming Marsh | land | Dual land slot supporting the black-green mana base.
- 4 Deathcap Glade | land | Dual land slot supporting the black-green mana base.
- 2 Llanowar Elves | ramp | Ramp option that helps the deck develop ahead of its threats.
- 4 Assassin's Trophy | removal | Flexible removal for opposing problems.
- 2 Bitter Triumph | removal | Additional efficient removal for opposing threats.
- 4 Darkstar Augur | draw | Draw source that keeps resources flowing during a longer game.
- 2 Phyrexian Arena | draw | Persistent draw source for the midrange plan.
- 4 Elvish Archdruid | synergy | Elf creature included for the deck's cohesive Elf theme.
- 4 Wood Elves | synergy | Elf creature included for the deck's cohesive Elf theme.
- 4 Goldvein Hydra | threat | Creature threat for a proactive board presence.
- 4 Vein Ripper | threat | Creature threat that pressures the opponent.
- 4 Thrashing Brontodon | threat | Creature threat that supports the deck's board-focused plan.
- 2 Massacre Girl, Known Killer | threat | Creature threat that adds another angle of pressure.

</details>

### 10. Standard aggro, tournament

Format: Standard. Theme: aggro. Pool: any_card. Shortlist: 294 names.

Grade: bad, score 0.01, model 20260907T081149Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $130.64 to buy, $130.64 the whole deck.

**Summary:** This red-white aggro deck aims to establish a fast board of threats, use its synergy cards to keep attacks productive, and clear resistance with removal. Its draw cards help it keep applying pressure after the first wave, while interaction protects its momentum against opposing plans. It gives up larger top-end threats and broad sweepers in favor of staying focused on an assertive, creature-driven game.

The quality model grades this deck below the precon baseline against the top lists of the format: the deck runs many spells as one copy, and that lowers the grade. Many of the cards appear in no top list, and that raises the grade. The cards pair in ways the top lists do not, and that lowers the grade.

- PLAN plan_coherent=partly: The broad shape is a low-curve creature deck with cheap removal, but it is a pile of one-ofs pulling in several directions — Caretaker's Talent, Dawnsire, Redcap Gutter-Dweller and Warleader's Call each want a different game than the two-drop beatdown plan, and no single synergy (prowess, tokens, equipment) is supported enough to matter.
- PLAN theme_fit=yes: It is a Standard-sized 60-card red-white creature deck with a mostly cheap curve, burn-ish interaction and a Boros dual-heavy mana base, which is what the request named.
- PLAN useful_as_built=partly: The 24-land Boros base is clean and the curve is castable, so it can be shuffled up and played, but with almost every noncreature a singleton the deck draws wildly different games and lacks the redundancy an aggro deck needs to close consistently.
- PLAN summary_honest=partly: It accurately admits the low-count, off-meta nature, but it overstates a "draw" package (Reckless Lackey and Irreverent Gremlin are not card draw) and claims the deck gives up larger top-end threats while it runs Dawnsire, Aurelia's Vindicator and Redcap Gutter-Dweller at the top.
- [INFO] `curve_summary`: average mana value 3.08 over 36 nonland cards

<details><summary>The deck list</summary>

- 4 Plains | land | Provides a reliable white mana source.
- 4 Mountain | land | Provides a reliable red mana source.
- 4 Sacred Foundry | land | Supplies both deck colors.
- 4 Inspiring Vantage | land | Supplies both deck colors for the aggressive plan.
- 4 Sunbillow Verge | land | Adds flexible red-white mana coverage.
- 4 Sundown Pass | land | Rounds out the two-color mana base.
- 2 Dinotomaton | threat | Adds to the deck's creature pressure.
- 2 Diversion Specialist | threat | Provides another proactive threat.
- 2 Dragonback Lancer | threat | Helps maintain a steady attack force.
- 2 Edgewall Pack | threat | Contributes to the aggressive threat density.
- 2 Frilled Sparkshooter | threat | Adds another threat for early pressure.
- 2 Redcap Gutter-Dweller | threat | Fills out the attacking creature base.
- 1 Slickshot Show-Off | synergy | Supports the deck's aggressive synergies.
- 1 Burnout Bashtronaut | synergy | Adds another synergy piece for the creature plan.
- 1 Shock Brigade | synergy | Reinforces the deck's proactive game plan.
- 1 Warleader's Call | synergy | Provides a dedicated synergy card for the deck.
- 1 Fugitive Codebreaker | draw | Provides card draw while remaining part of the aggressive shell.
- 1 Reckless Lackey | draw | Adds card draw to help sustain pressure.
- 1 Irreverent Gremlin | draw | Supplies another draw option.
- 1 Caretaker's Talent | draw | Provides ongoing card draw support.
- 1 Inspiring Overseer | draw | Adds draw alongside the creature plan.
- 1 Tersa Lightshatter | draw | Completes the deck's draw package.
- 2 Boros Charm | interaction | Provides flexible interaction for protecting the attack plan.
- 1 Sheltered by Ghosts | interaction | Adds interaction that supports the board presence.
- 1 Parting Gust | interaction | Provides a reactive interaction option.
- 1 Dawn's Truce | interaction | Helps defend the deck's proactive position.
- 1 Aven Interrupter | interaction | Adds an interactive creature to the deck.
- 2 Harsh Annotation | removal | Provides efficient removal for opposing permanents.
- 1 Case of the Gateway Express | removal | Adds another removal tool.
- 1 Krenko's Buzzcrusher | removal | Provides removal while fitting the red side of the deck.
- 1 Aurelia's Vindicator | removal | Adds white removal to the main deck.
- 1 Emeritus of Truce // Swords to Plowshares | removal | Provides a dedicated removal option.
- 1 Dawnsire, Sunstar Dreadnought | removal | Rounds out the removal suite.
- 1 Mongoose Lizard | removal | Completes the deck's removal package.

</details>

### 11. Commander with a locked card

Format: Commander. Theme: sacrifice. Pool: any_card. Shortlist: 296 names.

Commander: Karlov of the Ghost Council.

Grade: bad, score 0.06, model 20260907T081149Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $764.33 to buy, $764.33 the whole deck.

**Summary:** Karlov leads a creature-heavy sacrifice deck that builds a board of sacrifice synergies, converts those pieces into cards and mana, and uses removal to keep opposing boards under pressure. The deck closes through its larger threats and the cumulative pressure created by sacrificing creatures while preserving key pieces. It gives up some speed to support a full land base and several expensive finishers, so its strongest turns come after its board has developed.

The quality model grades this deck below the precon baseline against the top lists of the format: more opening hands hold two to four lands than the norm, and that raises the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. The deck holds no fast mana, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: The list is a tight aristocrats package: free/cheap sac outlets (Carrion Feeder, Viscera Seer, Ashnod's Altar, Altar of Dementia, Blasting Station), token and body generators (Sifter of Skulls, Pawn of Ulamog, Ghoulcaller Gisa, Abhorrent Overlord, Spirit Bonds), drain payoffs (Zulaport Cutthroat, Bastion of Remembrance, Vindictive Vampire, Fleshtaker) whose lifegain also feeds Karlov, plus edict engines (Grave Pact, Dictate of Erebos, Attrition, Yawgmoth) and recursion — everything pulls the same direction.
- PLAN theme_fit=yes: It is exactly a Commander sacrifice deck under Karlov of the Ghost Council, in Orzhov, with the sacrifice-for-lifedrain angle that best supports him.
- PLAN useful_as_built=yes: 37 lands with solid duals and pain lands, Sol Ring plus creature-based mana, a low curve with only a couple of top-end cards, and several distinct win routes (drain loops, big Karlov beats, Altar of Dementia mill, board-lock edicts) make it immediately playable.
- PLAN summary_honest=yes: The summary's claims — sacrifice synergies converted into cards and mana, removal pressure, a full land base with a few expensive finishers and slower early turns — all match the list, and nothing in the deck is concealed, though it undersells the Karlov lifegain-drain interaction.
- [INFO] `curve_summary`: average mana value 3.29 over 63 nonland cards

<details><summary>The deck list</summary>

- 14 Plains | land | Basic mana-base land.
- 12 Swamp | land | Basic mana-base land.
- 1 Command Tower | land | Mana-base land.
- 1 Godless Shrine | land | Mana-base land.
- 1 Scrubland | land | Mana-base land.
- 1 Caves of Koilos | land | Mana-base land.
- 1 City of Brass | land | Mana-base land.
- 1 Mana Confluence | land | Mana-base land.
- 1 Vault of Champions | land | Mana-base land.
- 1 Fetid Heath | land | Mana-base land.
- 1 Shattered Sanctum | land | Mana-base land.
- 1 Exotic Orchard | land | Mana-base land.
- 1 Sol Ring | ramp | Required ramp piece.
- 1 Ashnod's Altar | ramp | Sacrifice-focused ramp piece.
- 1 Phyrexian Altar | ramp | Sacrifice-focused ramp piece.
- 1 Pitiless Plunderer | ramp | Sacrifice-focused ramp piece.
- 1 Priest of Forgotten Gods | ramp | Sacrifice-focused ramp piece.
- 1 Warren Soultrader | ramp | Sacrifice-focused ramp piece.
- 1 Skullport Merchant | ramp | Sacrifice-focused ramp piece.
- 1 Pawn of Ulamog | ramp | Sacrifice-focused ramp piece.
- 1 Sifter of Skulls | ramp | Sacrifice-focused ramp piece.
- 1 Crowded Crypt | ramp | Sacrifice-focused ramp piece.
- 1 Corrupted Conviction | draw | Efficient sacrifice-focused draw.
- 1 Baron Bertram Graywater | draw | Draw support for the sacrifice plan.
- 1 Disciple of Bolas | draw | Sacrifice-focused draw support.
- 1 Relic Vial | draw | Draw support for the deck.
- 1 Shadowheart, Dark Justiciar | draw | Sacrifice-focused draw support.
- 1 Smothering Abomination | draw | Draw support for the sacrifice plan.
- 1 Vampiric Rites | draw | Repeatable sacrifice-focused draw.
- 1 Village Rites | draw | Efficient sacrifice-focused draw.
- 1 Tevesh Szat, Doom of Fools | draw | Draw support for the sacrifice plan.
- 1 Thraxodemon | draw | Draw support for the sacrifice plan.
- 1 Sanguine Spy | draw | Draw support for the deck.
- 1 Cartel Aristocrat | interaction | Sacrifice-focused interaction piece.
- 1 Fanatical Devotion | interaction | Protective interaction for the sacrifice plan.
- 1 Flare of Fortitude | interaction | Protective interaction piece.
- 1 Gift of Doom | interaction | Protective interaction piece.
- 1 Kinzu of the Bleak Coven | interaction | Interaction supporting the sacrifice plan.
- 1 Promise of Tomorrow | interaction | Interaction supporting the sacrifice plan.
- 1 Spirit Bonds | interaction | Interaction supporting the sacrifice plan.
- 1 Eldrazi Monument | interaction | Protective interaction for the board.
- 1 Attrition | removal | Sacrifice-focused removal.
- 1 Ayli, Eternal Pilgrim | removal | Removal supporting the sacrifice plan.
- 1 Blasting Station | removal | Sacrifice-focused removal.
- 1 Bone Shards | removal | Efficient sacrifice-focused removal.
- 1 Dictate of Erebos | removal | Removal supporting the sacrifice plan.
- 1 Grave Pact | removal | Removal supporting the sacrifice plan.
- 1 Eaten Alive | removal | Efficient sacrifice-focused removal.
- 1 Spark Harvest | removal | Efficient sacrifice-focused removal.
- 1 Yawgmoth, Thran Physician | removal | Removal supporting the sacrifice plan.
- 1 Altar of Dementia | synergy | Core sacrifice synergy piece.
- 1 Bastion of Remembrance | synergy | Core sacrifice synergy piece.
- 1 Carrion Feeder | synergy | Low-cost sacrifice synergy piece.
- 1 Viscera Seer | synergy | Low-cost sacrifice synergy piece.
- 1 Bartolomé del Presidio | synergy | Sacrifice synergy piece.
- 1 Fleshtaker | synergy | Sacrifice synergy piece.
- 1 Chthonian Nightmare | synergy | Sacrifice synergy piece.
- 1 Hidden Stockpile | synergy | Sacrifice synergy piece.
- 1 Victimize | synergy | Sacrifice synergy piece.
- 1 Zulaport Cutthroat | synergy | Core sacrifice synergy piece.
- 1 Felisa, Fang of Silverquill | threat | Threat that complements Karlov's plan.
- 1 Liesa, Forgotten Archangel | threat | High-impact threat.
- 1 Mondrak, Glory Dominus | threat | High-impact threat.
- 1 Ratadrabik of Urborg | threat | Threat for the creature-focused plan.
- 1 Liliana, Heretical Healer // Liliana, Defiant Necromancer | threat | Threat for the sacrifice plan.
- 1 Requiem Angel | threat | Threat for the creature-focused plan.
- 1 Razaketh, the Foulblooded | threat | Top-end sacrifice threat.
- 1 Vindictive Vampire | threat | Threat supporting the sacrifice plan.
- 1 Ghoulcaller Gisa | threat | Threat for the sacrifice plan.
- 1 Sadistic Hypnotist | threat | Threat supporting the sacrifice plan.
- 1 Threefold Thunderhulk | threat | High-impact threat.
- 1 Abhorrent Overlord | threat | Top-end threat.
- 1 Toxic Deluge | wipe | Efficient board wipe.
- 1 The Meathook Massacre | wipe | Board wipe that suits the sacrifice plan.
- 1 Austere Command | wipe | Flexible board wipe.

</details>

### 12. Commander on a budget

Format: Commander. Theme: tokens. Pool: any_card. Shortlist: 296 names.

Commander: Adeline, Resplendent Cathar.

Grade: bad, score 0.03, model 20260907T081149Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $75.32 to buy, $75.32 the whole deck.

**Summary:** Adeline, Resplendent Cathar leads a wide white token board that builds pressure through combat, token-focused enchantments, and creature payoffs. The deck develops mana and cards early, uses efficient answers to keep opposing threats manageable, then wins by turning an accumulated board of tokens into repeated attacks. It gives up premium acceleration and expensive standalone finishers in exchange for a consistent, budget-minded token plan with several ways to rebuild after a reset.

The quality model grades this deck below the precon baseline against the top lists of the format: more opening hands hold two to four lands than the norm, and that raises the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Nearly every card either makes tokens, pumps or protects a wide board, or draws off token/creature triggers, and the wipes chosen (Hour of Reckoning, Martial Coup) are ones that spare or rebuild the token board, so the pieces point one way; the only soft spot is the pile of marginal treasure trinkets standing in for ramp.
- PLAN theme_fit=yes: It is a mono-white Adeline token Commander deck built out of cheap, widely available cards at a modest, precon-adjacent power level, matching the theme, commander, budget, and bracket the requester named.
- PLAN useful_as_built=yes: 38 lands in a mono-colored deck with a low curve means it casts its spells reliably, and Adeline plus the token generators, anthems, and go-wide payoffs give clear, repeatable ways to close a game even if the trinket ramp is slow.
- PLAN summary_honest=yes: The summary claims a wide white token plan with token enchantments, creature payoffs, budget answers, and post-wipe rebuilding, all of which the list actually contains, and it openly admits it forgoes premium acceleration and big finishers rather than hiding the weak ramp.
- [INFO] `curve_summary`: average mana value 3.19 over 62 nonland cards

<details><summary>The deck list</summary>

- 28 Plains | land | Provides the reliable white mana base the deck needs.
- 1 Ancient Den | land | A white-producing land that also fits the artifact portion of the deck.
- 1 Castle Ardenvale | land | A white land slot with useful late-game token support.
- 1 Command Tower | land | Provides dependable commander-color mana.
- 1 Eiganjo, Seat of the Empire | land | A white source that occupies a flexible land slot.
- 1 Exotic Orchard | land | Adds another flexible color-producing land.
- 1 Legion's Landing // Adanto, the First Fort | land | A token-focused land slot that supports the deck's main plan.
- 1 Minas Tirith | land | Provides a white-producing utility land slot.
- 1 Path of Ancestry | land | A flexible land that supports the creature-heavy build.
- 1 Secluded Courtyard | land | Provides creature-focused color fixing.
- 1 Angelic Sell-Sword | draw | A low-cost draw card that keeps resources flowing.
- 1 Bygone Bishop | draw | Supports the deck's creature-based draw package.
- 1 Court of Grace | draw | Provides ongoing card advantage for the token strategy.
- 1 Dawn of Hope | draw | A flexible draw engine for the deck's longer games.
- 1 Faramir, Field Commander | draw | Adds card advantage while fitting the token-focused creature plan.
- 1 Glimmer Seeker | draw | A cheap draw option that helps maintain momentum.
- 1 Idol of Oblivion | draw | A dedicated draw artifact for a deck making tokens.
- 1 Sarah Jane Smith | draw | A low-cost draw creature that helps refill the hand.
- 1 Search the Premises | draw | Provides continuing card advantage during multiplayer games.
- 1 Staff of the Storyteller | draw | A budget artifact draw piece for the token shell.
- 1 Sunpearl Kirin | draw | Adds another inexpensive draw card to the deck.
- 1 Charisma Bobblehead | ramp | An artifact mana source that helps advance the board.
- 1 Coin of Mastery | ramp | Provides budget artifact-based mana acceleration.
- 1 Collector's Vault | ramp | A mana artifact that supports the deck's resource development.
- 1 Currency Converter | ramp | A cheap artifact ramp slot for early development.
- 1 Discerning Financier | ramp | Adds creature-based mana development to the deck.
- 1 Druidic Satchel | ramp | A low-cost artifact that supports mana development.
- 1 Goldvein Pick | ramp | An Equipment-based ramp option that works with attackers.
- 1 Keeper of the Accord | ramp | Provides creature-based ramp for a white deck.
- 1 Noble's Purse | ramp | A budget artifact mana source.
- 1 Prying Blade | ramp | An Equipment ramp piece that fits a combat-oriented deck.
- 1 Aerial Assault | removal | A cheap white answer for opposing permanents.
- 1 Banishing Slash | removal | An efficient removal spell for handling threats.
- 1 Battle Menu | removal | A flexible removal option at a low cost.
- 1 Citizen's Crowbar | removal | An Equipment-based answer that fits the creature plan.
- 1 Generous Gift | removal | A broad answer to troublesome opposing permanents.
- 1 Kellan's Lightblades | removal | An inexpensive instant-speed removal option.
- 1 Righteous Confluence | removal | Provides flexible removal coverage.
- 1 Ainok Strike Leader | interaction | Adds a creature-based interaction slot.
- 1 Basri Ket | interaction | Provides interactive support while contributing to the board.
- 1 Lena, Selfless Champion | interaction | A token-compatible creature that supports board protection.
- 1 Moonsilver Spear | interaction | An Equipment that supports the deck's combat plan and interaction suite.
- 1 Rootborn Defenses | interaction | Protects the developed token board from disruption.
- 1 Spirit Bonds | interaction | A token-oriented interaction piece for protecting key creatures.
- 1 Anafenza, Unyielding Lineage | synergy | A cheap token-synergy creature for the main plan.
- 1 Cat Collector | synergy | Adds a low-cost token-focused synergy piece.
- 1 Clarion Spirit | synergy | A creature that reinforces the token-making plan.
- 1 Divine Visitation | synergy | A powerful token-focused enchantment for converting a wide board into pressure.
- 1 Felidar Retreat | synergy | Supports the deck's land and token game plan.
- 1 Hanweir Militia Captain // Westvale Cult Leader | synergy | A creature that rewards building a wide board.
- 1 Horn of Gondor | synergy | A token-focused artifact that supports the central strategy.
- 1 Intangible Virtue | synergy | An efficient anthem-style token payoff.
- 1 Oketra's Monument | synergy | Supports the creature-heavy token plan.
- 1 Retrofitter Foundry | synergy | A compact artifact token engine.
- 1 Rosie Cotton of South Lane | synergy | A low-cost creature that rewards token development.
- 1 Siege Veteran | synergy | A Soldier-oriented payoff for a token board.
- 1 Skrelv's Hive | synergy | An enchantment-based token producer for steady board growth.
- 1 Ajani's Chosen | threat | A creature threat that fits the enchantment and token shell.
- 1 Archangel Elspeth | threat | A planeswalker threat that contributes to the token game plan.
- 1 Attended Healer | threat | A low-cost threat that works alongside token development.
- 1 Basri's Lieutenant | threat | A creature threat that rewards a developed board.
- 1 Cemetery Protector | threat | Provides a resilient creature threat for the midgame.
- 1 Defiler of Faith | threat | A creature threat that supports the deck's white permanent base.
- 1 Emeria Angel | threat | A land-focused token threat.
- 1 Gideon, Ally of Zendikar | threat | A planeswalker threat suited to building a wide board.
- 1 Hero of Bladehold | threat | A combat-focused creature threat for the token strategy.
- 1 Oketra the True | threat | A durable creature threat for the creature-heavy build.
- 1 Phantom General | threat | A token-focused creature threat and payoff.
- 1 Requiem Angel | threat | A flying creature threat that supports the token theme.
- 1 Ceaseless Conflict | wipe | A low-cost board-reset option.
- 1 Hour of Reckoning | wipe | A token-friendly sweep effect for recovering control of the board.
- 1 Martial Coup | wipe | Combines a board reset with the deck's token plan.

</details>

### 13. owned first, and the commander is not owned

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 229 names.

Commander: Karlov of the Ghost Council.

Grade: bad, score 0.33, model 20260907T081149Z.

Cards: 99 main, 0 sideboard. Repair turn: yes, for deck_size, copy_limit, profile_off_band, profile_off_band. Block findings: 0.

Cost: $0.00 to buy, $174.32 the whole deck.

**Summary:** This Orzhov lifegain deck builds a steady stream of life-gain triggers to grow its commander into a fast, dangerous attacker. Cheap artifacts, equipment, and value creatures maintain momentum, while broad removal and board resets keep opponents from establishing superior positions. Evasive Angels and enlarged creatures provide multiple ways to finish once the battlefield is under control.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The land count sits above the norm of the format, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: The bulk of the list serves one plan — incidental and food-based lifegain triggers (Rosie Cotton, Angel of Vitality, Second Breakfast, Lembas, Meager Meal, Compassionate Healer) feeding Karlov, plus protection equipment and evasive finishers — with only the two symmetrical wipes pulling slightly against its own creature base.
- PLAN theme_fit=yes: It is an Orzhov Karlov lifegain Commander deck at a bracket-3 power level with no combos or fast mana beyond Sol Ring, matching the request as written.
- PLAN useful_as_built=yes: The mana base is clean and generous (arguably too many lands at ~38 plus eight ramp pieces), the curve is low, and there are several evasive angels and a growing commander to actually close games.
- PLAN summary_honest=yes: Each claim maps to real cards: lifegain triggers and Karlov growth, cheap artifacts/equipment, a deep removal suite with Austere Command and Fumigate, and Angel finishers like Lyra and Victory's Herald, and it openly repeats the quality model's criticisms about land count and card overlap.
- [WARN] `not_owned`: Karlov of the Ghost Council: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 2.95 over 60 nonland cards
- [INFO] `mana_pass`: the builder moved 2 cards of the mana base to bring the deck inside its power level
- [INFO] `basics_added`: the list was 3 cards short, so the builder added 3 basic lands

<details><summary>The deck list</summary>

- 1 Aerith Gainsborough | synergy | Lifegain-focused creature that supports Karlov’s plan.
- 1 Angel of Invention | threat | Evasive lifelinking-style battlefield pressure for the deck’s combat plan.
- 1 Angel of Vitality | synergy | Rewards repeated life gain with a larger flying body.
- 1 Arcane Signet | ramp | Efficient color fixing and acceleration.
- 1 Austere Command | wipe | Flexible reset button for problematic boards.
- 1 Banishing Light | removal | Versatile answer to a troublesome permanent.
- 1 Bastion Protector | interaction | Keeps Karlov safer while advancing the board.
- 1 Bender's Waterskin | ramp | Early mana acceleration.
- 1 Bill the Pony | threat | Creature pressure that fits the deck’s life-positive theme.
- 1 Bitter Triumph | removal | Cheap, flexible creature or planeswalker answer.
- 1 Boromir, Warden of the Tower | interaction | Protects the board from disruptive spells.
- 1 Buster Sword | draw | Equipment-based card advantage.
- 1 Call of the Ring | draw | Steady card flow with life as a useful resource.
- 1 Champion's Helm | interaction | Protects Karlov and other legendary threats.
- 1 Clever Concealment | interaction | Protects a developed board through opposing removal.
- 1 Command Tower | land | Reliable access to both colors.
- 1 Compassionate Healer | synergy | Repeated life gain for Karlov and other payoffs.
- 1 Crib Swap | removal | Exiles a creature at instant speed.
- 1 City of Brass | land | Untapped access to either deck color.
- 1 Dawnhand Eulogist | threat | A single copy provides a resilient thematic threat.
- 1 Darksteel Plate | interaction | Durable protection for Karlov.
- 1 Dispatch | removal | Efficient creature interaction alongside the artifact package.
- 1 Eclipsed Realms | land | Additional colored land source.
- 1 Exemplar of Light | draw | Card advantage attached to a lifegain-friendly creature.
- 1 Exotic Orchard | land | Flexible dual-color fixing.
- 1 Fatal Push | removal | Low-cost creature removal.
- 1 Fellwar Stone | ramp | Two-mana fixing and acceleration.
- 1 Frodo, Sauron's Bane | threat | A scalable creature that can become a meaningful win condition.
- 1 Fumigate | wipe | Board reset that also replenishes life.
- 1 Generous Gift | removal | Answers nearly any permanent type.
- 1 Get Lost | removal | Efficient answer to creatures, enchantments, and planeswalkers.
- 1 Gollum, Silent Slinker // Meager Meal | threat | The adventure and creature provide pressure with useful flexibility.
- 1 Grand Coliseum | land | Fixes both colors while remaining an untapped land.
- 1 Idol of Oblivion | draw | Repeatable card advantage with token-producing support.
- 1 Infernal Grasp | removal | Clean answer to any creature.
- 1 Inspiring Overseer | draw | Flying body that replaces itself and supports life gain.
- 1 Kor Firewalker | synergy | Life gain payoff creature that helps establish the board.
- 1 Lembas | draw | Low-cost cantrip artifact that can later recur.
- 1 Light of Promise | synergy | Turns repeated life gain into a major creature threat.
- 1 Lightning Greaves | interaction | Haste and efficient protection for Karlov.
- 1 Lyra Dawnbringer | threat | Powerful lifelink attacker that stabilizes races.
- 1 Mask of Memory | draw | Combat-triggered filtering and card advantage.
- 1 Minas Tirith | land | Colored land that can provide card advantage later.
- 1 Minwu, White Mage | threat | Life-positive threat that supports the main game plan.
- 1 Night's Whisper | draw | Efficient card draw.
- 1 Path of Ancestry | land | Dual-color land with incidental creature support.
- 14 Plains | land | Reliable white sources for early lifegain creatures and Karlov.
- 1 Plaza of Heroes | land | Protects Karlov while providing colored mana.
- 1 Prideful Feastling | synergy | Life gain synergy that strengthens the deck’s central engine.
- 1 Rabaroo Troop | threat | Creature-based pressure for closing games.
- 1 Rosie Cotton of South Lane | synergy | Converts life gain into lasting battlefield growth.
- 1 Rooftop Percher | threat | Additional evasive creature pressure.
- 1 Second Breakfast | synergy | Supports the deck’s recurring life gain triggers.
- 1 Secluded Courtyard | land | Creature-focused color fixing.
- 1 Shattered Angel | threat | Flying lifegain threat that punishes opposing land development.
- 1 Skullclamp | draw | Highly efficient creature-based card draw.
- 1 Sneering Shadewriter | threat | A black creature threat that pressures life totals.
- 1 Sol Ring | ramp | Premium early acceleration.
- 1 Spire of Industry | land | Untapped fixing supported by the artifact package.
- 1 Stone Docent | synergy | Supports the deck’s life-gain-oriented creature plan.
- 12 Swamp | land | Reliable black sources for removal and Karlov.
- 1 Swiftfoot Boots | interaction | Protects Karlov while enabling immediate attacks.
- 1 Sword of the Animist | ramp | Equipment-based ramp that rewards attacking.
- 1 Swords to Plowshares | removal | Efficient exile-based creature removal.
- 1 Takenuma, Abandoned Mire | land | Black source with late-game creature recursion.
- 1 Thought Vessel | ramp | Colorless acceleration with an expanded hand size.
- 1 Tome of Legends | draw | Karlov helps keep this efficient draw engine stocked.
- 1 Unbreakable Formation | interaction | Protects the creature board and supports an aggressive finish.
- 1 Unclaimed Territory | land | Additional creature-based color fixing.
- 1 Victory's Herald | threat | Flying combat finisher with lifelink support.
- 1 Wall of Omens | draw | Early defense that immediately draws a card.
- 1 Wayfarer's Bauble | ramp | Reliable early land ramp and color fixing.
- 1 Well-Worn Spatula | synergy | Equipment support for the life-gain strategy.
- 1 Astral Cornucopia | ramp | the mana pass added it to bring the mana base inside the power level
- 1 Bofur, Reliable Guardian // Concerted Care | synergy | the mana pass added it to bring the mana base inside the power level

</details>

### 14. the user delegates the commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 244 names.

Commander: Atarka, World Render.

Grade: bad, score 0.05, model 20260907T081149Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $529.24 to buy, $529.24 the whole deck.

**Summary:** This is a red-green Dragon deck that develops its mana early, uses low-cost draw and protection to keep moving, and then presses the table with a steady sequence of Dragon threats led by Atarka, World Render. It wins by turning that Dragon pressure into combat finishes while removal and sweepers clear resistance. The tradeoff is that its strongest plays remain creature-centered, so it is less explosive and less stack-focused than faster high-power strategies.

The quality model grades this deck below the precon baseline against the top lists of the format: more opening hands hold two to four lands than the norm, and that raises the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. The deck holds no fast mana, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: Every slot serves the Dragon plan: cost reducers and Orbs to cheat on cost, Dragon-conditional removal (Draconic Roar, Dragon's Fire, Molten Exhale), creature-based draw, protection for the tribe, and payoffs like Dragon Tempest, Scourge of Valkas, Crucible of Fire and Thrakkus that convert Dragons into damage.
- PLAN theme_fit=yes: It is a Dragon tribal Commander deck with a builder-chosen Dragon commander (Atarka, World Render) and a mid-power, no-fast-mana, no-combo build appropriate to bracket 3.
- PLAN useful_as_built=yes: 36 lands plus Orbs, Dragon's Hoard, Sarkhan Fireblood and two Dragon cost reducers support a top-heavy but functional curve, and there are ample win routes (Atarka double strike, Thrakkus, Dragon Tempest/Scourge burn) with real interaction, though a few filler bodies like Dragon Hatchling and Kargan Dragonrider are weak.
- PLAN summary_honest=yes: The claims of ramp, cheap draw and protection, removal and sweepers, and creature-centered combat wins all match the actual list, and it openly concedes the lack of fast mana and low overlap with top lists.
- [INFO] `curve_summary`: average mana value 3.11 over 63 nonland cards

<details><summary>The deck list</summary>

- 12 Forest | land | Basic green land for dependable early mana.
- 13 Mountain | land | Basic red land for dependable early mana.
- 1 Command Tower | land | Flexible red-green land source.
- 1 Stomping Ground | land | Red-green dual land source.
- 1 Taiga | land | Red-green dual land source.
- 1 Spire Garden | land | Red-green dual land source.
- 1 Karplusan Forest | land | Red-green dual land source.
- 1 Rockfall Vale | land | Red-green dual land source.
- 1 Rootbound Crag | land | Red-green dual land source.
- 1 Commercial District | land | Red-green dual land source.
- 1 Cragcrown Pathway // Timbercrown Pathway | land | Flexible choice of red or green mana.
- 1 Cinder Glade | land | Red-green land source with basic land types.
- 1 Game Trail | land | Additional red-green land source.
- 1 Scaled Nurturer | ramp | Early Dragon-themed mana development.
- 1 Carnelian Orb of Dragonkind | ramp | Dragon-focused mana acceleration.
- 1 Jade Orb of Dragonkind | ramp | Dragon-focused mana acceleration.
- 1 Orb of Dragonkind | ramp | Flexible Dragon-themed mana acceleration.
- 1 Dragon's Hoard | ramp | Dragon-themed mana source for the midgame.
- 1 Ganax, Astral Hunter | ramp | Dragon ramp piece that fits the creature plan.
- 1 Sarkhan, Fireblood | ramp | Planeswalker-based mana development.
- 1 Goldspan Dragon | ramp | Dragon threat that also advances mana.
- 1 Gadrak, the Crown-Scourge | ramp | Dragon ramp option for building toward larger plays.
- 1 Atsushi, the Blazing Sky | ramp | Dragon ramp option with value in the main plan.
- 1 Faithless Looting | draw | Low-cost card selection and draw.
- 1 Demand Answers | draw | Efficient draw spell for finding action.
- 1 Thrill of Possibility | draw | Low-cost draw and hand filtering.
- 1 Skullclamp | draw | Cheap repeatable draw support.
- 1 Sylvan Library | draw | Early ongoing card selection.
- 1 Elemental Bond | draw | Ongoing draw support for the creature-heavy plan.
- 1 Garruk's Uprising | draw | Creature-focused draw engine.
- 1 Beast Whisperer | draw | Creature-based draw support.
- 1 Dragonborn Champion | draw | Dragon-themed draw source.
- 1 Guardian Project | draw | Long-game creature draw engine.
- 1 Return of the Wildspeaker | draw | Flexible draw spell for a board of large creatures.
- 1 Heroic Intervention | interaction | Broad protective interaction for the board.
- 1 Lightning Greaves | interaction | Protection equipment for key creatures.
- 1 Swiftfoot Boots | interaction | Additional protection equipment for key creatures.
- 1 Snakeskin Veil | interaction | Low-cost protective interaction.
- 1 Veil of Summer | interaction | Efficient reactive protection.
- 1 Tamiyo's Safekeeping | interaction | Protective interaction for an important permanent.
- 1 Pyroblast | interaction | Cheap red interaction.
- 1 Red Elemental Blast | interaction | Cheap red interaction.
- 1 Dragon's Fire | removal | Efficient Dragon-themed removal.
- 1 Draconic Roar | removal | Low-cost Dragon-themed removal.
- 1 Molten Exhale | removal | Direct removal option.
- 1 Piercing Exhale | removal | Instant-speed removal option.
- 1 Spit Flame | removal | Dragon-themed removal for creature boards.
- 1 Invasion of Tarkir // Defiant Thundermaw | removal | Dragon-themed removal that adds a Dragon threat.
- 1 Dragon Tempest | removal | Dragon synergy that supplies removal pressure.
- 1 Glorybringer | removal | Dragon threat that provides removal.
- 1 Scourge of Valkas | removal | Dragon-themed removal payoff.
- 1 Crucible of Fire | synergy | Core Dragon tribal payoff.
- 1 Dragonlord's Servant | synergy | Cost support for the Dragon-heavy deck.
- 1 Dragonspeaker Shaman | synergy | Additional Dragon cost support.
- 1 Dragon Egg | synergy | Early Dragon-themed board presence.
- 1 Dragon Hatchling | synergy | Low-cost Dragon creature for tribal density.
- 1 Dragonkin Berserker | synergy | Dragon-focused creature payoff.
- 1 Firespitter Whelp | synergy | Dragon creature that supports the tribal plan.
- 1 Minion of the Mighty | synergy | Low-cost creature supporting Dragon deployment.
- 1 Sarkhan's Triumph | synergy | Dragon-focused library search.
- 1 Kargan Dragonrider | synergy | Low-cost creature supporting the Dragon plan.
- 1 Dragonloft Idol | threat | Low-cost Dragon threat.
- 1 Fang, Roku's Companion | threat | Efficient Dragon threat for the curve.
- 1 Caldera Pyremaw | threat | Midrange Dragon threat.
- 1 Sagu Wildling // Roost Seek | threat | Flexible Dragon threat.
- 1 Thunderbreak Regent | threat | Reliable Dragon threat for the tribal board.
- 1 Verix Bladewing | threat | Midrange legendary Dragon threat.
- 1 Territorial Hellkite | threat | Dragon attacker that maintains pressure.
- 1 Stormbreath Dragon | threat | Hasty Dragon threat.
- 1 Red Dragon | threat | Straightforward Dragon threat.
- 1 Thrakkus the Butcher | threat | Legendary Dragon threat for combat-focused games.
- 1 Lathliss, Dragon Queen | threat | High-impact Dragon tribal threat.
- 1 Twinflame Tyrant | threat | Top-end Dragon threat.
- 1 Breath Weapon | wipe | Efficient creature-board reset.
- 1 Draconic Intervention | wipe | Flexible Dragon-themed board wipe.
- 1 Ryusei, the Falling Star | wipe | Dragon threat with a wipe role.

</details>

### 15. delegated commander, owned first

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 229 names.

Commander: Astarion, the Decadent.

Grade: bad, score 0.08, model 20260907T081149Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $171.88 the whole deck.

**Summary:** This lifegain deck develops its mana, establishes a board of creatures, and uses lifegain-focused support pieces to make its threats matter more over a long game. Astarion, the Decadent anchors a steady white-black plan that protects key creatures, removes opposing obstacles, and can reset crowded boards before rebuilding. It aims to win through sustained creature pressure rather than explosive early turns, trading raw speed for resilience and a broad set of answers.

The quality model grades this deck below the precon baseline against the top lists of the format: more opening hands hold two to four lands than the norm, and that raises the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. The deck holds no fast mana, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core is a coherent white-black lifegain value pile (Angel of Vitality, Rosie Cotton, Light of Promise, Exemplar of Light, Shattered Angel, Lyra) backed by ramp, draw and removal, but three unconditional wraths plus commander-protection equipment sit awkwardly against a creature-pressure win plan that the deck itself is building.
- PLAN theme_fit=yes: It is a Commander lifegain deck with a builder-chosen lifegain commander in Astarion, the Decadent, and the power level and lack of fast mana beyond Sol Ring sit comfortably at bracket 2.
- PLAN useful_as_built=yes: 37 lands plus ten mana rocks/ramp, a dozen card-draw pieces, seven spot removal spells and multiple evasive finishers (Lyra, Victory's Herald, Angel of Invention) make it a playable, self-sufficient list out of the box.
- PLAN summary_honest=yes: It accurately describes the deck as slow creature-based value with protection, removal and wraths, and it does not claim combos, tutors or fast mana the list lacks — it even flags the absence of fast mana.
- [WARN] `not_owned`: Astarion, the Decadent: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 2.98 over 62 nonland cards

<details><summary>The deck list</summary>

- 17 Plains | land | Provides reliable white mana for the deck’s white-heavy creatures and spells.
- 12 Swamp | land | Provides reliable black mana for the deck’s black spells.
- 1 Command Tower | land | Provides flexible mana for the commander’s colors.
- 1 City of Brass | land | Provides flexible colored mana.
- 1 Exotic Orchard | land | Provides flexible colored mana in multiplayer games.
- 1 Path of Ancestry | land | Provides fixing while supporting the commander’s creature type.
- 1 Castle Locthwain | land | Adds a black source with additional utility.
- 1 Takenuma, Abandoned Mire | land | Adds a black source with additional utility.
- 1 Thriving Moor | land | Adds another black source and color fixing.
- 1 White Lotus Hideout | land | Adds another white source to support the deck’s white cards.
- 1 Sol Ring | ramp | Provides efficient early mana acceleration.
- 1 Arcane Signet | ramp | Provides reliable color fixing and acceleration.
- 1 Fellwar Stone | ramp | Provides efficient mana acceleration in multiplayer games.
- 1 Wayfarer's Bauble | ramp | Helps develop mana early while finding a basic land.
- 1 Sword of the Animist | ramp | Turns attacking creatures into ongoing mana development.
- 1 Thought Vessel | ramp | Provides mana acceleration without taking up a colored source.
- 1 Relic of Legends | ramp | Provides mana acceleration and supports the deck’s legendary creatures.
- 1 Commander's Sphere | ramp | Provides fixing early and can become a card later.
- 1 Chromatic Lantern | ramp | Provides broad color fixing alongside mana acceleration.
- 1 Lotho, Corrupt Shirriff | ramp | Contributes to the mana plan from a lifegain-friendly color pair.
- 1 Buster Sword | draw | Provides card advantage from an Equipment slot.
- 1 Call of the Ring | draw | Supplies repeatable card advantage.
- 1 Exemplar of Light | draw | Adds card advantage on a creature that fits the deck’s plan.
- 1 Idol of Oblivion | draw | Provides a low-cost source of card advantage.
- 1 Inspiring Overseer | draw | Provides card advantage on a creature body.
- 1 Lembas | draw | Provides immediate card advantage from a low-cost artifact.
- 1 Mask of Memory | draw | Turns combat connections into card selection and advantage.
- 1 Night's Whisper | draw | Provides efficient black card advantage.
- 1 Skullclamp | draw | Provides an efficient Equipment-based card advantage engine.
- 1 Tome of Legends | draw | Works well with repeatedly deploying and using the commander.
- 1 Wall of Omens | draw | Provides a defensive early body while replacing itself.
- 1 Bastion Protector | interaction | Helps keep the commander protected on the battlefield.
- 1 Boromir, Warden of the Tower | interaction | Provides a creature-based disruptive element.
- 1 Champion's Helm | interaction | Protects the commander from opposing disruption.
- 1 Clever Concealment | interaction | Protects the established board from opposing interaction.
- 1 Darksteel Plate | interaction | Gives an important creature durable protection.
- 1 Lightning Greaves | interaction | Protects key creatures while helping them contribute immediately.
- 1 Banishing Light | removal | Answers a problematic opposing permanent.
- 1 Bitter Triumph | removal | Provides flexible spot removal.
- 1 Dispatch | removal | Provides efficient creature removal.
- 1 Generous Gift | removal | Answers a wide range of opposing permanents.
- 1 Infernal Grasp | removal | Provides dependable creature removal.
- 1 Swords to Plowshares | removal | Provides efficient creature removal.
- 1 Stroke of Midnight | removal | Answers a broad range of opposing permanents.
- 1 Austere Command | wipe | Offers a flexible reset when the board gets out of hand.
- 1 Fumigate | wipe | Resets creature-heavy boards while fitting the lifegain plan.
- 1 Vanquish the Horde | wipe | Provides an efficient creature-board reset.
- 1 Aerith Gainsborough | synergy | Supports the deck’s central lifegain theme.
- 1 Angel of Vitality | synergy | Rewards the deck for pursuing its lifegain plan.
- 1 Compassionate Healer | synergy | Adds another creature dedicated to the lifegain theme.
- 1 Elixir | synergy | Provides a compact artifact piece for the deck’s synergy package.
- 1 Kor Firewalker | synergy | Supports the deck’s lifegain-focused game plan.
- 1 Light of Promise | synergy | Converts the deck’s lifegain focus into a stronger creature presence.
- 1 Night Nurse, Healer of Heroes | synergy | Adds a legendary lifegain-themed synergy piece.
- 1 Prideful Feastling | synergy | Adds another creature that fits the lifegain plan.
- 1 Pull from the Grave | synergy | Returns an important creature from the deck’s graveyard.
- 1 Rosie Cotton of South Lane | synergy | Rewards the deck for building around lifegain and creatures.
- 1 Second Breakfast | synergy | Provides a lifegain-themed support spell.
- 1 Well-Worn Spatula | synergy | Adds an Equipment-based piece to the deck’s synergy package.
- 1 White Mage's Staff | synergy | Adds another Equipment that supports the deck’s game plan.
- 1 Angel of Invention | threat | Provides a substantial creature threat for the battlefield.
- 1 Bill the Pony | threat | Adds a creature threat that supports the deck’s board presence.
- 1 Dawnhand Eulogist | threat | Provides another meaningful creature threat.
- 1 Foggy Swamp Hunters | threat | Adds a creature threat that helps pressure opponents.
- 1 Gollum, Silent Slinker // Meager Meal | threat | Provides a black creature threat with an additional spell option.
- 1 Lo and Li, Twin Tutors | threat | Adds a legendary creature threat to diversify the board.
- 1 Lyra Dawnbringer | threat | Provides a powerful Angel threat for the deck’s creature plan.
- 1 Minwu, White Mage | threat | Adds a legendary creature threat that fits the white creature base.
- 1 Reaping Willow | threat | Provides a durable creature threat for longer games.
- 1 Shattered Angel | threat | Adds an Angel threat that complements the lifegain theme.
- 1 Sneering Shadewriter | threat | Provides an additional black creature threat.
- 1 Victory's Herald | threat | Provides a high-impact creature threat for closing games.

</details>

### 16. a tight budget, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 229 names.

Commander: Denethor, Ruling Steward.

Grade: bad, score 0.37, model 20260907T081149Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $17.69 to buy, $39.85 the whole deck.

**Summary:** This deck builds a white-black creature board around a sacrifice-focused aristocrats plan, using its inexpensive mana pieces and draw cards to keep deploying creatures. Its threats give the deck a way to convert that board presence into a win while removal and several sweepers prevent opposing boards from taking over. The deck gives up premium lands, expensive protection, and the fastest available acceleration in favor of an owned-card-first build that stays comfortably within the budget.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The sac outlets plus drain payoffs (Blood Artist, Zulaport, Bastion of Remembrance) form one clear aristocrats engine, but the deck has essentially no token or fodder generation to feed it and packs three symmetrical sweepers plus a commander-protection subpackage that work against holding a wide board.
- PLAN theme_fit=yes: It is a W/B aristocrats Commander deck under Denethor with cheap, low-power cards consistent with a $25 bracket-2 build from an existing collection.
- PLAN useful_as_built=yes: 37 basics in a two-color deck with several rocks supports the curve, and the many drain-on-death effects plus creature bodies give real, repeatable ways to close a game.
- PLAN summary_honest=yes: Every claim - sacrifice payoffs, cheap rocks, card draw, removal and multiple sweepers, and the budget/owned-cards constraint - matches the list, and it openly admits the missing lands and slower ramp.
- [WARN] `not_owned`: Bastion of Remembrance: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ayli, Eternal Pilgrim: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Aron, Benalia's Ruin: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bartolomé del Presidio: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Baron Bertram Graywater: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Relic Vial: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Skullport Merchant: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Sivriss, Nightmare Speaker: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Umbral Collar Zealot: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Woe Strider: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ruthless Technomancer: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Blood Artist: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Elas il-Kor, Sadistic Pilgrim: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Falkenrath Noble: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vindictive Vampire: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Zulaport Cutthroat: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Yahenni, Undying Partisan: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Shadowheart, Dark Justiciar: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Disciple of Bolas: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Smothering Abomination: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vampire Gourmand: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vengeful Bloodwitch: the deck needs 1, the collection has 0
- [WARN] `not_owned`: High-Society Hunter: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 2.82 over 62 nonland cards

<details><summary>The deck list</summary>

- 19 Plains | land | Untapped white mana base for the deck.
- 18 Swamp | land | Untapped black mana base for the deck.
- 1 Arcane Signet | ramp | Reliable early mana fixing and acceleration.
- 1 Sol Ring | ramp | Efficient acceleration for developing the board.
- 1 Fellwar Stone | ramp | Low-cost mana acceleration.
- 1 Thought Vessel | ramp | Mana acceleration that remains useful on board.
- 1 Blitzball | ramp | An owned ramp piece for the early turns.
- 1 Astral Cornucopia | ramp | Flexible artifact mana acceleration.
- 1 Chromatic Lantern | ramp | Fixes colors while accelerating mana.
- 1 Commander's Sphere | ramp | Color-fixing mana acceleration.
- 1 Deadly Dispute | ramp | A sacrifice-compatible ramp option.
- 1 Inherited Envelope | ramp | An owned artifact ramp piece.
- 1 Skullclamp | draw | A draw engine suited to a creature-focused plan.
- 1 Wall of Omens | draw | An efficient white draw creature.
- 1 Inspiring Overseer | draw | An owned creature-based draw option.
- 1 Massacre Girl, Known Killer | draw | A black draw option that also adds to the creature count.
- 1 Painful Truths | draw | Efficient black draw spell.
- 1 Tome of Legends | draw | A low-cost repeatable draw artifact.
- 1 Stone of Erech | draw | An owned artifact draw piece.
- 1 Idol of Oblivion | draw | An artifact draw option for the board plan.
- 1 Lembas | draw | An owned artifact draw piece.
- 1 Mask of Memory | draw | Equipment-based card flow for the creature suite.
- 1 Nasty End | draw | An owned black draw spell.
- 1 Bastion Protector | interaction | Protects an important creature in the deck's plan.
- 1 Bronze Guardian | interaction | Artifact-centered protection for the board.
- 1 Frontline Medic | interaction | Creature-based defensive interaction.
- 1 Gift of Immortality | interaction | Keeps a key creature involved in the game plan.
- 1 Swiftfoot Boots | interaction | Low-cost protection for important creatures.
- 1 Together Forever | interaction | Protective interaction for the creature board.
- 1 Bitter Triumph | removal | Flexible black spot removal.
- 1 Claim the Precious | removal | An owned black removal spell.
- 1 Crib Swap | removal | White creature removal.
- 1 Deadly Precision | removal | An owned removal option.
- 1 Destroy Evil | removal | Flexible white removal.
- 1 Fatal Push | removal | Efficient black spot removal.
- 1 Infernal Grasp | removal | Reliable creature removal.
- 1 Austere Command | wipe | Flexible board reset when the table gets ahead.
- 1 Dusk // Dawn | wipe | A board reset that fits the creature-focused deck.
- 1 Vanquish the Horde | wipe | An owned creature-board reset.
- 1 Gollum the Abandoned | synergy | A black synergy creature for the sacrifice-focused shell.
- 1 Gollum, Patient Plotter | synergy | An owned black synergy creature for the aristocrats plan.
- 1 Bastion of Remembrance | synergy | A dedicated black synergy permanent for the aristocrats shell.
- 1 Ayli, Eternal Pilgrim | synergy | A low-cost legendary creature supporting the sacrifice-focused plan.
- 1 Aron, Benalia's Ruin | synergy | A low-cost white-black synergy creature.
- 1 Bartolomé del Presidio | synergy | A low-cost white-black synergy creature.
- 1 Baron Bertram Graywater | synergy | A black legendary creature selected for aristocrats synergy.
- 1 Relic Vial | synergy | A low-cost artifact support piece for the deck's creature plan.
- 1 Skullport Merchant | synergy | A black creature supporting the sacrifice-focused shell.
- 1 Sivriss, Nightmare Speaker | synergy | A low-cost black synergy creature.
- 1 Umbral Collar Zealot | synergy | A black creature selected to support the aristocrats core.
- 1 Woe Strider | synergy | A black creature that strengthens the sacrifice-focused plan.
- 1 Ruthless Technomancer | synergy | A black creature providing a higher-impact synergy piece.
- 1 Blood Artist | threat | A premium aristocrats threat for closing games through the creature plan.
- 1 Elas il-Kor, Sadistic Pilgrim | threat | A white-black aristocrats threat.
- 1 Falkenrath Noble | threat | A low-cost black threat for the creature-focused plan.
- 1 Vindictive Vampire | threat | A black Vampire threat for the aristocrats shell.
- 1 Zulaport Cutthroat | threat | A compact black aristocrats threat.
- 1 Yahenni, Undying Partisan | threat | A black legendary creature that adds pressure to the board.
- 1 Shadowheart, Dark Justiciar | threat | A black legendary threat that works with the creature plan.
- 1 Disciple of Bolas | threat | A black creature threat with value-oriented deck positioning.
- 1 Smothering Abomination | threat | A black creature threat for the sacrifice-focused board.
- 1 Vampire Gourmand | threat | A low-cost Vampire threat.
- 1 Vengeful Bloodwitch | threat | A black Vampire threat for the deck's creature suite.
- 1 High-Society Hunter | threat | A black creature threat that helps maintain board pressure.

</details>

### 17. upgrade a precon, owned first

Format: Commander. Theme: goblins. Pool: owned_first. Shortlist: 196 names.

Commander: Zada, Hedron Grinder.

Grade: typical, score 0.52, model 20260907T081149Z.

Cards: 99 main, 0 sideboard. Repair turn: yes, for deck_size, precon_share, profile_off_band, profile_off_band, profile_off_band, profile_off_band, profile_off_band, profile_off_band, profile_off_band. Block findings: 0.

Cost: $84.58 to buy, $271.30 the whole deck.

**Summary:** A mono-red Goblin swarm deck built around multiplying the value of targeted spells, token makers, and tribal haste effects. It develops a broad battlefield, converts creature density into explosive mana and card flow, and closes through massed combat, storm-fueled pressure, and direct damage from Goblin synergies.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: The cheap single-target cantrips and pump spells, wide goblin token makers, sacrifice/ritual mana, and Grapeshot/Empty the Warrens/Past in Flames all feed the same Zada mass-copy storm plan, with protective equipment sensibly included to keep the commander alive.
- PLAN theme_fit=yes: It is a mono-red goblin storm list with a modest, non-combo-oriented power level consistent with a bracket 3 precon upgrade.
- PLAN useful_as_built=yes: 37 lands plus rituals and rocks, a low curve, and several distinct win routes (Krenko/anthem swarms, Impact Tremors and Boggart Shenanigans drain, Grapeshot/Haze of Rage) make it immediately playable, though its turn-four mana is a bit below average.
- PLAN summary_honest=yes: Every claim — targeted-spell multiplication, token makers, haste anthems, creature-to-mana conversion, combat and storm/direct-damage finishes — maps to actual cards in the list.
- [INFO] `curve_summary`: average mana value 2.60 over 63 nonland cards
- [WARN] `profile_off_band`: the count of nonbasic lands that make only colorless mana is 5, and bracket 3 wants 4 at most (Fountainport, Goblin Burrows, Kher Keep, Reliquary Tower, War Room)
- [WARN] `two_card_combo`: bracket 3 allows no two-card infinite combo that needs four mana or less, and the deck holds Storm-Kiln Artist + Haze of Rage (speed 5, near two-card)

<details><summary>The deck list</summary>

- 1 Arena of Glory | land | Utility land that supports aggressive combat turns.
- 1 Castle Embereth | land | Goblin-focused mana source with a team-pump option.
- 1 Den of the Bugbear | land | Manland that remains a threat after a wipe.
- 1 Forgotten Cave | land | Cycling land that can become a fresh spell late.
- 1 Fountainport | land | Flexible utility land for the token strategy.
- 1 Goblin Burrows | land | Tribal land that makes a Goblin attack more dangerous.
- 1 Hidden Volcano | land | Red source with useful late-game utility.
- 1 Kher Keep | land | Produces bodies for token payoffs and targeted spells.
- 1 Reliquary Tower | land | Keeps a large hand after burst draw turns.
- 1 Shinka, the Bloodsoaked Keep | land | Legendary red source that improves combat.
- 1 War Room | land | Repeatable card advantage from a land slot.
- 25 Mountain | land | Reliable untapped red mana for a spell-heavy deck.
- 1 Ancestors' Aid | synergy | Efficient targeted combat trick for a creature-heavy board.
- 1 Battle Hymn | ramp | Converts a wide Goblin board into explosive red mana.
- 1 Blasphemous Act | wipe | Efficient emergency reset against overwhelming creature boards.
- 1 Boggart Shenanigans | removal | Turns Goblin deaths into direct damage.
- 1 Brightstone Ritual | ramp | Produces a burst of mana from Goblin creatures.
- 1 Chaos Warp | removal | Answers problematic permanents of any type.
- 1 Conspicuous Snoop | synergy | Provides tribal velocity and works with Goblin density.
- 1 Crimson Wisps | draw | Cantrip that grants haste to a key creature.
- 1 Daring Discovery | synergy | Flexible thematic spell that supports the deck's game plan.
- 1 Dragon Fodder | synergy | Creates multiple Goblin bodies for go-wide payoffs.
- 1 Empty the Warrens | synergy | Produces a large Goblin board after chaining spells.
- 1 Expedite | draw | Cheap haste enabler that replaces itself.
- 1 Faithless Looting | draw | Efficient hand filtering with flashback value.
- 1 Fists of Flame | draw | Scales with spellcasting while drawing a card.
- 1 Frontline Heroism | synergy | Supports the token swarm's combat plan.
- 1 Gempalm Incinerator | removal | Goblin-scaled creature removal with cycling value.
- 1 General Kreat, the Boltbringer | synergy | Goblin legend that reinforces the tribal plan.
- 1 Goblin Bushwhacker | synergy | Turns a developing token board into immediate pressure.
- 1 Goblin Chieftain | synergy | Tribal anthem and haste effect for Goblins.
- 1 Goblin Dark-Dwellers | threat | Recasts useful instants or sorceries while adding a body.
- 1 Goblin Lackey | synergy | Accelerates Goblin deployment through combat.
- 1 Goblin Matron | synergy | Finds the most useful Goblin for the current board.
- 1 Goblin Negotiation | removal | Thematic answer to opposing creatures or permanents.
- 1 Goblin Trashmaster | removal | Tribal body that provides artifact removal.
- 1 Goblin Warchief | synergy | Reduces Goblin costs and grants haste.
- 1 Grapeshot | removal | Storm payoff that can clear small creatures or finish players.
- 1 Haze of Rage | wincon | Storm-fueled combat finisher for a wide board.
- 1 Idol of Oblivion | draw | Reliable draw engine alongside frequent token creation.
- 1 Impact Tremors | wincon | Makes every Goblin token contribute direct damage.
- 1 Krenko's Command | synergy | Efficiently adds multiple Goblins to the battlefield.
- 1 Krenko, Mob Boss | threat | Major repeatable token producer and primary battlefield threat.
- 1 Mana Geyser | ramp | Explosive ritual that enables high-impact turns.
- 1 Mogg War Marshal | synergy | Multiple Goblin bodies from one inexpensive card.
- 1 Past in Flames | synergy | Lets a spell-heavy graveyard fuel a decisive turn.
- 1 Pashalik Mons | removal | Converts Goblins and their deaths into damage.
- 1 Quest for the Goblin Lord | synergy | Tribal payoff that rewards continuous Goblin deployment.
- 1 Roaming Throne | synergy | Doubles relevant Goblin triggered abilities.
- 1 Ruby Medallion | ramp | Reduces the cost of the deck's many red spells.
- 1 Sazacap's Brew | draw | Low-cost filtering that keeps spells flowing.
- 1 Searslicer Goblin | synergy | Goblin attacker that supports the tribal combat plan.
- 1 Seething Song | ramp | Ritual mana for explosive multi-spell turns.
- 1 Siege-Gang Commander | removal | Creates a Goblin squad and turns it into removal or reach.
- 1 Siege-Gang Lieutenant | removal | Adds Goblin material and provides useful board control.
- 1 Skullclamp | draw | Exceptional card flow with disposable Goblin tokens.
- 1 Sol Ring | ramp | Efficient early mana acceleration.
- 1 Spreading Insurrection | wincon | High-impact aggressive payoff for a developed game.
- 1 Swiftfoot Boots | interaction | Protects a key creature while enabling haste.
- 1 Throne of Eldraine | ramp | Mana development that also offers late-game utility.
- 1 Vandalblast | wipe | Flexible artifact answer with a powerful overload mode.
- 1 Wild Ride | synergy | Combat-focused support spell for an aggressive board.
- 1 Witch's Mark | draw | Card filtering that can also grant haste.
- 1 Arcane Signet | ramp | Consistent low-cost mana acceleration.
- 1 Skirk Prospector | ramp | Converts Goblin bodies into burst mana.
- 1 Lightning Greaves | interaction | Protects vital creatures and gives immediate haste.
- 1 Darksteel Plate | interaction | Durable protection for an important creature.
- 1 Champion's Helm | interaction | Protects the commander while improving combat.
- 1 Goblin Bombardment | removal | Turns spare Goblins into repeatable creature control and reach.
- 1 Fable of the Mirror-Breaker // Reflection of Kiki-Jiki | ramp | Provides mana, filtering, and token-copying value.
- 1 Rundvelt Hordemaster | synergy | Tribal anthem that provides card advantage when Goblins die.
- 1 Moria Marauder | synergy | Aggressive Goblin that rewards a broad creature base.
- 1 Ancestral Anger | draw | Cheap cantrip that grows a creature for combat.
- 1 Renegade Tactics | draw | Inexpensive cantrip that disrupts a blocker.
- 1 Storm-Kiln Artist | ramp | Generates Treasure from the deck's frequent instants and sorceries.

</details>

### 18. upgrade a precon, any card

Format: Commander. Theme: turtles. Pool: any_card. Shortlist: 300 names.

Commander: Heroes in a Half Shell.

Grade: bad, score 0.29, model 20260907T081149Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $220.08 to buy, $220.08 the whole deck.

**Summary:** This Turtle Power deck keeps the precon’s broad Turtle and TMNT character theme while improving its mana base and adding a few focused Turtle payoffs. It develops its board through ramp and card draw, then presses the table with Turtle threats supported by synergy pieces and removal. The deck gives up some raw efficiency for a theme-forward, varied creature strategy with the precon’s distinctive cards intact.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The list is largely the precon's grab-bag of TMNT characters, equipment-ish gadgets and combat tricks pointing in several directions at once, with only a handful of actual Turtle payoffs (Fecund Greenshell, The Pride of Hull Clade, Turtle Power!) tying it together — the summary is honest that this is a broad, theme-forward pile rather than a single tight plan.
- PLAN theme_fit=yes: It is a Heroes in a Half Shell Commander deck that keeps the Turtle Power precon's identity, upgrades fixing and adds Turtle synergy, and stays at a modest, no-game-changer power level appropriate to bracket 3.
- PLAN useful_as_built=partly: The land count and five-color fixing are adequate to cast spells, but ramp is thin (Sol Ring, Signet, Lantern, Tokka & Rahzar, Greenshell) for a WUBRG deck full of multicolor cards, so it plays out slowly and its win path is just generic creature beatdown.
- PLAN summary_honest=yes: The summary claims exactly what the list shows — precon shell preserved, mana base improved with Cavern of Souls/City of Brass/Mana Confluence and the tri/dual cycle, a few added Turtle payoffs — and openly concedes the loss of raw efficiency rather than overselling a plan the cards don't carry.
- [INFO] `curve_summary`: average mana value 3.58 over 60 nonland cards
- [WARN] `profile_off_band`: the land count is 39, and bracket 3 wants 34 to 38
- [WARN] `profile_off_band`: the average mana value of the nonland cards is 3.58, and bracket 3 wants 2 to 3.5
- [WARN] `profile_off_band`: the ramp count is 5, and bracket 3 wants 8 to 13
- [WARN] `profile_off_band`: the draw count is 2, and bracket 3 wants 8 to 14
- [WARN] `profile_off_band`: the removal count is 4, and bracket 3 wants 6 to 12
- [WARN] `profile_off_band`: the wipe count is 1, and bracket 3 wants 2 to 5
- [WARN] `profile_off_band`: the interaction count is 0, and bracket 3 wants 4 to 12

<details><summary>The deck list</summary>

- 1 Acidic Slime | removal | Kept from Turtle Power precon for removal.
- 1 April O'Neil, Live on the Scene | draw | Kept from Turtle Power precon for draw.
- 1 Arcade Cabinet | other | Kept from Turtle Power precon to preserve its theme.
- 1 Arcane Signet | ramp | Kept from Turtle Power precon for ramp.
- 1 Ash Barrens | land | Kept from Turtle Power precon as a land.
- 1 Assassin's Trophy | removal | Kept from Turtle Power precon for removal.
- 1 Baxter, Fly in the Ointment | other | Kept from Turtle Power precon to preserve its theme.
- 1 Bebop, Skull & Crossbones | other | Kept from Turtle Power precon to preserve its theme.
- 1 Big Apple, 3 a.m. | land | Kept from Turtle Power precon as a land.
- 1 Big Mother Mouser | other | Kept from Turtle Power precon to preserve its theme.
- 1 Biogenic Ooze | other | Kept from Turtle Power precon to preserve its theme.
- 1 Blasphemous Act | wipe | Kept from Turtle Power precon for board control.
- 1 Casey Jones, Back Alley Brute | other | Kept from Turtle Power precon to preserve its theme.
- 1 Chromatic Lantern | ramp | Kept from Turtle Power precon for ramp.
- 1 Cinder Glade | land | Kept from Turtle Power precon as a land.
- 1 Coin of Mastery | other | Kept from Turtle Power precon to preserve its theme.
- 1 Command Tower | land | Kept from Turtle Power precon as a land.
- 1 Continue? | other | Kept from Turtle Power precon to preserve its theme.
- 1 Corpsejack Menace | other | Kept from Turtle Power precon to preserve its theme.
- 1 Dimension X Pizzasaur | other | Kept from Turtle Power precon to preserve its theme.
- 1 Donatello, the Brains | synergy | Kept from Turtle Power precon for Turtle synergy.
- 1 Double Jump // Flying Kick | other | Kept from Turtle Power precon to preserve its theme.
- 1 Dragonskull Summit | land | Kept from Turtle Power precon as a land.
- 1 Electric Seaweed | other | Kept from Turtle Power precon to preserve its theme.
- 1 Endless Foot Assault | other | Kept from Turtle Power precon to preserve its theme.
- 1 Escape Tunnel | land | Kept from Turtle Power precon as a land.
- 1 Everything Pizza | other | Kept from Turtle Power precon to preserve its theme.
- 1 Evolving Wilds | land | Kept from Turtle Power precon as a land.
- 1 Exploding Barrel | other | Kept from Turtle Power precon to preserve its theme.
- 1 Fast Forward | other | Kept from Turtle Power precon to preserve its theme.
- 1 Foot Chopper | other | Kept from Turtle Power precon to preserve its theme.
- 1 Game Over | other | Kept from Turtle Power precon to preserve its theme.
- 1 Grand Coliseum | land | Kept from Turtle Power precon as a land.
- 1 Harmonize | other | Kept from Turtle Power precon to preserve its theme.
- 1 Here Comes a New Hero! | other | Kept from Turtle Power precon to preserve its theme.
- 1 Hidden Hideout | land | Kept from Turtle Power precon as a land.
- 1 High Score | other | Kept from Turtle Power precon to preserve its theme.
- 1 Irma, Part-Time Mutant | other | Kept from Turtle Power precon to preserve its theme.
- 1 Krang, the All-Powerful | other | Kept from Turtle Power precon to preserve its theme.
- 1 Leatherhead, Iron Gator | other | Kept from Turtle Power precon to preserve its theme.
- 1 Leonardo, the Balance | threat | Kept from Turtle Power precon as a threat.
- 1 Lessons from Life | other | Kept from Turtle Power precon to preserve its theme.
- 1 Level Up | other | Kept from Turtle Power precon to preserve its theme.
- 1 Michelangelo, the Heart | synergy | Kept from Turtle Power precon for Turtle synergy.
- 1 Mole Module | other | Kept from Turtle Power precon to preserve its theme.
- 1 Mona Lisa, Science Geek | other | Kept from Turtle Power precon to preserve its theme.
- 1 Ninja Pizza | other | Kept from Turtle Power precon to preserve its theme.
- 1 Path of Ancestry | land | Kept from Turtle Power precon as a land.
- 1 Rain-Slicked Copse | land | Kept from Turtle Power precon as a land.
- 1 Rat King, Pale Piper | other | Kept from Turtle Power precon to preserve its theme.
- 1 Ray Fillet, Wave Warrior | other | Kept from Turtle Power precon to preserve its theme.
- 1 Roadkill Rodney | other | Kept from Turtle Power precon to preserve its theme.
- 1 Rocksteady, Mutant Marauder | other | Kept from Turtle Power precon to preserve its theme.
- 1 Rootbound Crag | land | Kept from Turtle Power precon as a land.
- 1 Shellshock | other | Kept from Turtle Power precon to preserve its theme.
- 1 Shredder, Shadow Master | other | Kept from Turtle Power precon to preserve its theme.
- 1 Smoldering Marsh | land | Kept from Turtle Power precon as a land.
- 1 Sodden Verdure | land | Kept from Turtle Power precon as a land.
- 1 Sol Ring | ramp | Kept from Turtle Power precon for ramp.
- 1 Special Move | other | Kept from Turtle Power precon to preserve its theme.
- 1 Steelbane Hydra | removal | Kept from Turtle Power precon for removal.
- 1 Super Combo | other | Kept from Turtle Power precon to preserve its theme.
- 1 Swift Demise | other | Kept from Turtle Power precon to preserve its theme.
- 1 Tempestra, Dame of Games | other | Kept from Turtle Power precon to preserve its theme.
- 1 Thriving Grove | land | Kept from Turtle Power precon as a land.
- 1 Thriving Isle | land | Kept from Turtle Power precon as a land.
- 1 Thriving Moor | land | Kept from Turtle Power precon as a land.
- 1 Together Forever | other | Kept from Turtle Power precon to preserve its theme.
- 1 Tokka & Rahzar, Unsupervised | ramp | Kept from Turtle Power precon for ramp.
- 1 Turtle Lair | land | Kept from Turtle Power precon as a land.
- 1 Undergrowth Stadium | land | Kept from Turtle Power precon as a land.
- 1 Vanquish the Horde | other | Kept from Turtle Power precon to preserve its theme.
- 1 Vernal Fen | land | Kept from Turtle Power precon as a land.
- 1 Vibrant Cityscape | land | Kept from Turtle Power precon as a land.
- 1 Vigor | other | Kept from Turtle Power precon to preserve its theme.
- 1 Voracious Hydra | other | Kept from Turtle Power precon to preserve its theme.
- 1 Wave Goodbye | other | Kept from Turtle Power precon to preserve its theme.
- 3 Forest | land | Basic mana base for the deck.
- 2 Island | land | Basic mana base for the deck.
- 2 Mountain | land | Basic mana base for the deck.
- 2 Plains | land | Basic mana base for the deck.
- 2 Swamp | land | Basic mana base for the deck.
- 1 Cavern of Souls | land | Adds a land that supports the Turtle creature theme.
- 1 City of Brass | land | Adds flexible mana to the land base.
- 1 Exotic Orchard | land | Adds flexible mana to the land base.
- 1 Mana Confluence | land | Adds flexible mana to the land base.
- 1 Morphic Pool | land | Adds a blue-black land to the mana base.
- 1 Rejuvenating Springs | land | Adds a green-blue land to the mana base.
- 1 Spire Garden | land | Adds a red-green land to the mana base.
- 1 Fecund Greenshell | ramp | Adds Turtle-themed ramp.
- 1 Kappa Tech-Wrecker | removal | Adds Turtle-themed removal.
- 1 The Pride of Hull Clade | draw | Adds Turtle-themed draw.
- 1 Turtle Power! | synergy | Strengthens the Turtle synergy plan.

</details>

### 19. the Hobbit family, two colours

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 174 names.

Commander: Thranduil, the Elvenking.

Grade: bad, score 0.03, model 20260907T081149Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $527.66 to buy, $527.66 the whole deck.

**Summary:** This deck builds an Elf-heavy board around Thranduil and develops its mana early with Elves and artifacts. It aims to keep cards flowing, disrupt key opposing plays, clear away troublesome permanents when needed, and win by turning a broad creature board into sustained combat pressure. It gives up some individual-card flexibility for a strongly themed collection of Elves, legendary items, and locations from the Hobbit setting.

The quality model grades this deck below the precon baseline against the top lists of the format: more opening hands hold two to four lands than the norm, and that raises the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core is a coherent Sultai Elves/tribal-go-wide plan with ramp, draw and removal, but the mana base is oddly skewed (14 Swamps, only 6 Forests) for a green Elf-heavy list, and three board wipes plus a mass sacrifice pull against the summary's stated go-wide combat pressure plan.
- PLAN theme_fit=yes: It is a Commander deck led by Thranduil, the Elvenking, built almost entirely from Hobbit-set/Middle-earth cards at a middling, non-optimized power level appropriate to bracket 3.
- PLAN useful_as_built=partly: It has enough lands, ramp, draw and win-by-creatures to be picked up and played, but the color balance of the mana base badly under-supports the green-heavy Elf creatures and early green ramp.
- PLAN summary_honest=partly: It accurately describes the Elf/artifact/location theme, ramp, draw and removal, but 'win by turning a broad creature board into sustained combat pressure' is undercut by the wipes it never mentions and by no anthem or overrun-style payoff.
- [INFO] `curve_summary`: average mana value 3.19 over 63 nonland cards
- [INFO] `mana_pass`: the builder moved 5 cards of the mana base to bring the deck inside its power level

<details><summary>The deck list</summary>

- 6 Forest | land | Basic green source for the Elf-heavy creature base.
- 8 Island | land | Basic blue source for the deck's blue cards.
- 14 Swamp | land | Basic black source for the deck's black cards.
- 1 Dragon-Cursed Halls | land | Themed nonbasic land that broadens the mana base.
- 1 Elven Passage | land | Themed nonbasic land for the Elf-led deck.
- 1 Elvenking's Halls | land | Themed nonbasic land for Thranduil's court.
- 1 Hobbit Hole | land | A themed nonbasic land that adds mana-base variety.
- 1 Minas Morgul, Dark Fortress | land | A themed nonbasic land for the black portion of the deck.
- 1 Mirkwood | land | A thematic land alongside the deck's Mirkwood Elves.
- 1 Rivendell | land | A thematic land that supports the Elven setting.
- 1 The Shire | land | A themed nonbasic land that completes the land package.
- 1 Arcane Signet | ramp | Reliable artifact mana acceleration.
- 1 Delighted Halfling | ramp | Early creature-based mana acceleration.
- 1 Elven Chorus | ramp | An Elf-themed mana engine.
- 1 Elvish Archdruid | ramp | Elf creature that contributes to mana development.
- 1 Elvish Mystic | ramp | Low-cost Elf mana acceleration.
- 1 Mox Amber | ramp | Low-cost legendary-themed mana acceleration.
- 1 Necklace of Girion | ramp | Themed artifact mana acceleration.
- 1 Wayfarer's Bauble | ramp | Early artifact ramp that helps develop mana.
- 1 Wood Elves | ramp | Elf-based mana acceleration that fits the creature plan.
- 1 Woodland Weavemaster | ramp | Elf Druid that adds to mana development.
- 1 Beorn the Fierce | draw | A themed card that supplies card advantage.
- 1 Elvish Visionary | draw | Elf creature that provides card advantage.
- 1 Fateful Discovery | draw | A dedicated card-advantage enchantment.
- 1 Hithlain Knots | draw | Instant-speed card advantage.
- 1 Key to the Side-Door | draw | Themed artifact card advantage.
- 1 Last March of the Ents | draw | A themed spell that replenishes resources.
- 1 Lórien Revealed | draw | A Lórien-themed card-advantage spell.
- 1 Night's Whisper | draw | Efficient dedicated card advantage.
- 1 Palantír of Orthanc | draw | Legendary artifact card advantage.
- 1 Thrór's Map | draw | Themed legendary artifact card advantage.
- 1 Uncover the Moon-Letters | draw | Enchantment-based card advantage.
- 1 Bilbo's Ring | interaction | Legendary equipment that protects the board plan.
- 1 Confusticate and Bebother | interaction | A flexible instant for disrupting opposing plays.
- 1 Dwarven Mattock | interaction | Equipment-based utility for the creature board.
- 1 Elrond, Moon-Reader | interaction | Legendary Elf utility that supports the game plan.
- 1 Mithril Coat | interaction | Protective legendary equipment for key creatures.
- 1 My Precious // Allure of Power | interaction | Legendary equipment with a utility Adventure.
- 1 Stern Scolding | interaction | Efficient instant-speed disruption.
- 1 The One Ring | interaction | Legendary artifact utility for the deck's plan.
- 1 Bilbo's Deadly Slice | removal | Dedicated spot removal.
- 1 Bitter Downfall | removal | Efficient instant-speed removal.
- 1 Colossal Whale | removal | A creature that also serves as removal.
- 1 Crude Bent Blade | removal | Equipment-based removal that remains on theme.
- 1 Enchanted River's Grasp | removal | Enchantment-based removal.
- 1 Giant's Boulder | removal | Themed artifact removal.
- 1 Merciless Executioner | removal | Creature-based removal for the board plan.
- 1 Orcish Bowmasters | removal | Efficient creature-based removal.
- 1 The Black Arrow | removal | Legendary equipment that provides removal.
- 1 Gnashing of Teeth | wipe | A themed board-clearing spell.
- 1 Languish | wipe | A dependable creature-board reset.
- 1 Raise the Palisade | wipe | A board reset that can favor the Elf creature base.
- 1 Arwen, Weaver of Hope | synergy | Legendary Elf that reinforces the Elf-centered plan.
- 1 Boughside Wanderers | synergy | Elf creature that expands the thematic board.
- 1 Cantankerous Keepers | synergy | Elf Soldier that supports the creature-focused strategy.
- 1 Celeborn the Wise | synergy | Legendary Elf that strengthens the Elven theme.
- 1 Galadhrim Guide | synergy | Elf Scout that builds the Elf presence.
- 1 Galion, Elvenking's Butler | synergy | Thranduil-themed legendary Elf support.
- 1 Guardian of the Halls | synergy | Elf Soldier that develops the themed board.
- 1 Lothlórien Lookout | synergy | Elf Scout that supports the Elf creature density.
- 1 Nimrodel Watcher | synergy | Elf Scout that reinforces the tribal core.
- 1 Thranduil's Company | synergy | Elf Soldier creature that directly fits Thranduil's court.
- 1 Attercop | threat | A substantial themed creature for pressuring opponents.
- 1 Elven Raft-Steerer | threat | Elf creature that adds to the attacking board.
- 1 Elvenking's Harper | threat | Elf Bard that adds another body to the board plan.
- 1 Gigantic Big Bear | threat | Large creature threat for closing games through combat.
- 1 Grey Havens Navigator | threat | Elf creature that broadens the creature offense.
- 1 Haunt of the Dead Marshes | threat | Nightmare Elf that adds pressure and thematic depth.
- 1 Mirkwood Elk | threat | A Mirkwood creature that contributes to combat pressure.
- 1 Mirkwood Meditator | threat | Elf Druid that adds another themed attacker.
- 1 Mirkwood Nurturer | threat | Elf Ranger that builds the creature board.
- 1 Mirkwood Pathmaker | threat | Elf Ranger that adds to the combat plan.
- 1 Thranduil, Sindarin Liege // Silvan Rally | threat | A Thranduil-themed legendary creature that supports the primary board.
- 1 Willow-Wind | threat | Elemental creature that offers a distinct combat threat.

</details>

### 20. the Hobbit family, a delegated commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 77 names.

Commander: Smaug the Magnificent.

Grade: bad, score 0.06, model 20260907T081149Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $484.11 to buy, $484.11 the whole deck.

**Summary:** This deck ramps toward Smaug the Magnificent and its Dragon creatures, using supporting creatures and equipment to maintain a combat-focused board. Removal and sweeping effects clear away opposition so the larger threats can close the game. It gives up depth in card flow, defensive interaction, and dedicated Dragon synergy in exchange for a straightforward, mountain-heavy creature plan.

The quality model grades this deck below the precon baseline against the top lists of the format: more opening hands hold two to four lands than the norm, and that raises the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The list is a coherent mono-red goodstuff creatures-and-equipment beatdown built around a big commander, but it is a grab-bag of dwarves, goblins, equipment and one-off effects rather than a focused plan, and it lacks the Dragon subtheme the summary and request point at.
- PLAN theme_fit=no: The person asked for a dragons deck, and apart from Smaug the Magnificent and one or two Smaug-named spells the list is dwarves, goblins and equipment with essentially no Dragon creatures or Dragon payoffs.
- PLAN useful_as_built=partly: Mono-red with 36 Mountains plus ramp rocks casts everything reliably and there are plenty of bodies and equipment to attack with, but a single color with almost no card draw and few real finishers makes it a flat, grindy deck that can stall out.
- PLAN summary_honest=yes: The summary openly admits the deck gives up card flow, interaction and dedicated Dragon synergy, which matches a list that has almost no other Dragons and thin draw.
- [INFO] `curve_summary`: average mana value 3.16 over 63 nonland cards
- [WARN] `profile_off_band`: the draw count is 6, and bracket 3 wants 8 to 14
- [INFO] `cards_trimmed`: the list was 1 card over, so the builder cut this card: Cavern-Hoard Dragon

<details><summary>The deck list</summary>

- 1 Andúril, Flame of the West | other | Equipment support for the creature-focused plan.
- 1 Andúril, Narsil Reforged | other | Equipment support for the creature-focused plan.
- 1 Arcane Signet | ramp | Reliable listed mana acceleration.
- 1 Balin, Loremaster | draw | Included as a listed card-draw source.
- 1 Bilbo's Ring | interaction | Included as a listed interaction piece.
- 1 Bombur, Gentle Dreamer | other | A supporting creature for the deck's board presence.
- 1 Bothersome Noisemaker | other | A supporting creature for the deck's board presence.
- 1 Burn, Burn, Tree and Fern | ramp | Included as a listed ramp card.
- 1 Call Forth the Tempest | wipe | A listed board-wipe option.
- 1 Desert Were-Worm | threat | A listed creature threat.
- 1 Desolation of Smaug | wipe | A listed board-wipe option.
- 1 Dori, Bearer of Friends | other | A supporting legendary creature.
- 1 Dragon's Desire | ramp | Included as a listed ramp card.
- 1 Dwarven Mattock | interaction | Included as a listed interaction piece.
- 1 Dwarven Mauler | other | A supporting creature for the deck's board presence.
- 1 Dáin Ironfoot | other | A supporting legendary creature.
- 1 Fire of Orthanc | removal | Included as a listed removal spell.
- 1 Fíli and Kíli, Joyous | ramp | Included as a listed ramp card.
- 1 Gandalf, Goblins' Bane // Flameshape | other | A flexible supporting legendary creature.
- 1 Gandalf, Spark Starter | removal | Included as a listed removal option.
- 1 Getaway Barrel | other | A supporting artifact for the deck's board plan.
- 1 Giant's Boulder | removal | Included as a listed removal option.
- 1 Glamdring | other | Equipment support for the creature-focused plan.
- 1 Glóin the Mighty // Easy Pickings | wipe | A listed board-wipe option.
- 1 Goblin Cratermaker | removal | Included as a listed removal creature.
- 1 Gundabad Opportunist | other | A supporting creature for the deck's board presence.
- 1 Guttersnipe | other | A supporting creature for the deck's board presence.
- 1 Improvised Club | removal | Included as a listed removal spell.
- 1 Inferno Titan | removal | A substantial creature that supplies listed removal.
- 1 Iron Hills Stalwart | other | A supporting creature for the deck's board presence.
- 1 Key to the Side-Door | draw | Included as a listed card-draw source.
- 1 Last Light of Durin's Day | synergy | The shortlist's dedicated synergy card.
- 1 Long-Bodied Grey Dog | ramp | Included as a listed ramp creature.
- 1 Long-Lost Lances | other | Equipment support for the creature-focused plan.
- 1 Misty Mountains Raider | other | A supporting creature for the deck's board presence.
- 1 Mithril Coat | interaction | Included as a listed interaction piece.
- 36 Mountain | land | Provides a consistent basic-land mana base.
- 1 Mox Amber | ramp | A low-cost listed ramp piece.
- 1 Old Thrush | other | A supporting creature for the deck's board presence.
- 1 Oliphaunt | other | A supporting creature for the deck's board presence.
- 1 Olog-hai Crusher | other | A supporting creature for the deck's board presence.
- 1 Orcish Siegemaster | other | A supporting creature for the deck's board presence.
- 1 Orcrist, Goblin-cleaver | ramp | Equipment that also supplies listed ramp.
- 1 Palantír of Orthanc | draw | Included as a listed card-draw source.
- 1 Pinecone Strike | removal | Included as a listed removal spell.
- 1 Ragged Short Spear | draw | Included as a listed card-draw source.
- 1 Smaug's Fury | other | A supporting instant for the deck's plan.
- 1 Smaug, the Great Calamity // Spew Flame | removal | A Dragon that also supplies listed removal.
- 1 Smite the Deathless | removal | Included as a listed removal spell.
- 1 Snowslope Hunter | other | A supporting creature for the deck's board presence.
- 1 Sting, Bilbo's Sword | other | Equipment support for the creature-focused plan.
- 1 Stone-Giant of High Pass | removal | Included as a listed removal creature.
- 1 The Black Arrow | removal | Included as a listed removal option.
- 1 The Misty Mountains Cold | ramp | Included as a listed ramp card.
- 1 The One Ring | interaction | Included as a listed interaction piece.
- 1 The Reaver Cleaver | ramp | Equipment that also supplies listed ramp.
- 1 Thorin, Company's Leader | ramp | Included as a listed ramp creature.
- 1 Thorin, Mountain-king | removal | Included as a listed removal creature.
- 1 Thrór's Map | draw | Included as a listed card-draw source.
- 1 Tidings of War | other | A supporting sorcery for the deck's plan.
- 1 Troop of Ponies | ramp | Included as a listed ramp creature.
- 1 Wayfarer's Bauble | ramp | A low-cost listed ramp piece.
- 1 Well-Worn Spatula | other | Equipment support for the creature-focused plan.
- 1 Óin the Brave | draw | Included as a listed card-draw source.

</details>

### 21. the Hobbit family, mana from outside

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 156 names.

Commander: Smaug the Impenetrable.

Grade: bad, score 0.05, model 20260907T081149Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $859.13 to buy, $859.13 the whole deck.

**Summary:** This deck ramps into Smaug the Impenetrable and follows with a steady stream of Dragons, Goblins, Trolls, and other battlefield threats. It aims to win by maintaining combat pressure, supporting its key creatures with equipment, and clearing away resistance with targeted removal or sweeping effects. Its tradeoff is a creature- and board-focused approach: it is strongest when it can keep developing permanents, rather than relying on a highly reactive or combo-oriented finish.

The quality model grades this deck below the precon baseline against the top lists of the format: more opening hands hold two to four lands than the norm, and that raises the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The list is broadly a red-black creature-and-equipment beatdown pile that matches the summary's stated board-pressure plan, but the goblin payoffs, dragon payoffs, troll bodies, and equipment package pull in several directions and few cards actually build toward Smaug's own hoard/treasure-style payoff.
- PLAN theme_fit=partly: Smaug the Impenetrable leads a heavily Middle-earth list at a reasonable bracket-3 power, but the request restricted non-land, non-rock cards to the Hobbit set and roughly a dozen-plus slots come from Tales of Middle-earth or from outside the LOTR product entirely (The One Ring, Orcish Bowmasters, both Andúrils, Palantír of Orthanc, Night's Whisper, Goblin Cratermaker, Bitter Downfall).
- PLAN useful_as_built=yes: Thirty-six lands with fetches, duals, and Command Tower support the two-color curve, and the deck has ample creatures, equipment, removal, and card draw to play a full game out of the box.
- PLAN summary_honest=partly: The board-focused, equipment-and-removal description matches the list and the quality caveats are stated plainly, but 'a steady stream of Dragons' oversells a deck holding only the two alternate Smaug cards as dragons.
- [INFO] `curve_summary`: average mana value 3.37 over 63 nonland cards
- [INFO] `mana_pass`: the builder moved 2 cards of the mana base to bring the deck inside its power level
- [WARN] `outside_requested_set`: the sets you named do not hold 14 cards: Ancient Tomb, Arid Mesa, Blood Crypt, Bloodstained Mire, Bojuka Bog, Command Tower, Dragonskull Summit, Exotic Orchard, Fabled Passage, Polluted Delta, Scalding Tarn, Urborg, Tomb of Yawgmoth, Verdant Catacombs, Wooded Foothills

<details><summary>The deck list</summary>

- 1 Ancient Tomb | land | Fast colorless land for getting key spells onto the table sooner.
- 1 Arid Mesa | land | Fixes red mana through the deck’s Mountain lands and Blood Crypt.
- 1 Blood Crypt | land | Reliable dual source for both deck colors.
- 1 Bloodstained Mire | land | Finds the deck’s red or black dual and basic sources.
- 1 Bojuka Bog | land | Black source with a useful disruptive land slot.
- 1 Command Tower | land | Consistent access to either commander color.
- 1 Dragonskull Summit | land | Dual land supporting the red-black mana base.
- 1 Exotic Orchard | land | Flexible dual-color fixing in multiplayer games.
- 1 Fabled Passage | land | Fetches the basic color needed most.
- 1 Goblin-town | land | Themed red source for the deck.
- 1 Mount Doom | land | Themed land slot that supports the mana base.
- 10 Mountain | land | Primary untapped red source for the deck’s red-heavy threats and ramp.
- 1 Polluted Delta | land | Finds black basics or Blood Crypt.
- 1 Scalding Tarn | land | Finds red basics or Blood Crypt.
- 9 Swamp | land | Primary untapped black source for draw and removal.
- 1 The Lonely Mountain | land | Red source that keeps the land base on theme.
- 1 Urborg, Tomb of Yawgmoth | land | Strengthens black access across the mana base.
- 1 Verdant Catacombs | land | Finds black basics or Blood Crypt.
- 1 Wooded Foothills | land | Finds red basics or Blood Crypt.
- 1 Balin, Loremaster | draw | Dedicated card-advantage piece for maintaining resources.
- 1 Gollum, Riddle Master | draw | Dedicated card-advantage piece for maintaining resources.
- 1 Key to the Side-Door | draw | Adds a repeatable card-advantage slot.
- 1 Night's Whisper | draw | Efficient early card-advantage spell.
- 1 Palantír of Orthanc | draw | Provides a durable card-advantage engine.
- 1 Rage into the Valley | draw | Adds another card-advantage spell to avoid running out of pressure.
- 1 Ragged Short Spear | draw | Equipment slot that also contributes card advantage.
- 1 Reverent Howl | draw | Instant-speed card-advantage option.
- 1 The Master of Lake-town | draw | Creature-based card advantage that supports the board plan.
- 1 Thrór's Map | draw | Artifact card-advantage piece that fits the deck’s permanent plan.
- 1 Óin the Brave | draw | Adds card advantage while remaining a creature for board presence.
- 1 Bilbo's Ring | interaction | Protective and disruptive utility for important permanents.
- 1 Dwarven Mattock | interaction | Flexible equipment utility slot.
- 1 Getaway Barrel | interaction | Flexible support artifact for responding to opposing plans.
- 1 Long-Lost Lances | interaction | Equipment utility that helps the deck interact through combat.
- 1 Mithril Coat | interaction | Protection for Smaug or another key threat.
- 1 My Precious // Allure of Power | interaction | Versatile interaction slot with an equipment body.
- 1 The One Ring | interaction | High-impact defensive utility and resilience piece.
- 1 Well-Worn Spatula | interaction | Additional equipment-based utility for the board plan.
- 1 Arcane Signet | ramp | Efficient fixing and acceleration for casting Smaug on schedule.
- 1 Bag End Banquet | ramp | Additional acceleration toward the deck’s larger spells.
- 1 Bolg's Company | ramp | Creature-based acceleration that also develops the board.
- 1 Burn, Burn, Tree and Fern | ramp | Themed ramp that advances the main game plan.
- 1 Dragon's Desire | ramp | Ramp spell that helps bridge into Smaug and other threats.
- 1 Mox Amber | ramp | Low-cost acceleration alongside the deck’s legendary commander.
- 1 Orcrist, Goblin-cleaver | ramp | Equipment-based acceleration for the creature plan.
- 1 Smaug the Magnificent | ramp | Themed Dragon acceleration that complements the commander.
- 1 Smaug, Wicked Worm | ramp | Additional Smaug-themed ramp and battlefield presence.
- 1 Azog, Moria's Ruin | removal | Creature-based answer that also leaves behind pressure.
- 1 Battle-Scarred Goblin | removal | Low-cost creature removal for keeping blockers and threats in check.
- 1 Bilbo's Deadly Slice | removal | Efficient targeted answer.
- 1 Bitter Downfall | removal | Flexible targeted answer for troublesome permanents.
- 1 Bolg of the North | removal | Themed creature answer that contributes to board presence.
- 1 Fire of Orthanc | removal | Direct removal for clearing key opposing pieces.
- 1 Goblin Cratermaker | removal | Flexible creature-based removal.
- 1 Improvised Club | removal | Cheap removal option that fits the aggressive plan.
- 1 Orcish Bowmasters | removal | Efficient removal attached to a creature.
- 1 Bolg, Erebor's Reckoning | wipe | Themed reset button when opponents establish larger boards.
- 1 Desolation of Smaug | wipe | Smaug-themed sweeper for resetting crowded boards.
- 1 Along the Crooked Way | synergy | Support enchantment for the Smaug-focused game plan.
- 1 Andúril, Flame of the West | synergy | Equipment support for turning a major creature into a stronger threat.
- 1 Andúril, Narsil Reforged | synergy | Additional equipment support for the creature-focused plan.
- 1 Down, Down to Goblin-town | synergy | Themed support card for the deck’s red-black creature plan.
- 1 Gathering of Darkness | synergy | Black support piece that reinforces the deck’s central plan.
- 1 Goblin Plate Mail | synergy | Equipment support that rewards keeping creatures on the battlefield.
- 1 Gollum the Abandoned | synergy | Themed creature support for the deck’s broader battlefield plan.
- 1 Gollum, Silent Slinker // Meager Meal | synergy | Flexible themed support card with a creature body.
- 1 Great Goblin, Foul-Hearted | synergy | Themed legendary support piece for the aggressive plan.
- 1 Smaug's Fury | synergy | Smaug-themed support spell for the deck’s main plan.
- 1 Desert Were-Worm | threat | Large battlefield threat that helps pressure opponents.
- 1 Dreaded Bat-Cloud | threat | Creature threat that broadens the deck’s board presence.
- 1 Fearsome Goblin Pair | threat | Efficient creature threat for developing pressure early.
- 1 Great Ugly-Looking Goblin // Clap! Snap! | threat | Flexible creature threat that remains useful across game stages.
- 1 Gundabad Opportunist | threat | Creature threat supporting the deck’s aggressive board plan.
- 1 Misty Mountains Raider | threat | Themed attacker that adds to the pressure package.
- 1 Nighthowl Pursuer | threat | Black creature threat for maintaining battlefield pressure.
- 1 Olog-hai Crusher | threat | Substantial creature threat for closing games.
- 1 Orcish Siegemaster | threat | Creature threat that reinforces the red-black board plan.
- 1 Ravening Warg | threat | Creature threat that adds another attacking body.
- 1 The Great Goblin | threat | Legendary Goblin threat for sustained pressure.
- 1 Troll of Khazad-dûm | threat | Large thematic threat for finishing games through combat.
- 1 Giant's Boulder | ramp | the mana pass added it to bring the mana base inside the power level
- 1 Wayfarer's Bauble | wipe | the mana pass added it to bring the mana base inside the power level

</details>

### 22. a set family and a card from outside it

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 175 names.

Commander: Thranduil, the Elvenking.

Grade: bad, score 0.05, model 20260907T081149Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $427.49 to buy, $427.49 the whole deck.

**Summary:** This is a Thranduil-led, Elf-centered creature deck that develops mana early, refuels with a broad draw package, and builds a board of thematic creatures backed by equipment and interaction. It aims to win through sustained creature pressure after clearing key blockers or resetting crowded boards with sweepers. The deck gives up some raw speed for a balanced, flavorful board plan with varied answers and several larger creatures to close games.

The quality model grades this deck below the precon baseline against the top lists of the format: more opening hands hold two to four lands than the norm, and that raises the grade. Many of the cards appear in no top list, and that lowers the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The mana/draw shell and Elf ramp point one way, but the top-end "threat" package is a grab-bag of non-Elf beasts and the deck runs three symmetrical sweepers (Languish, Gnashing of Teeth) that punish its own wide creature board, so the pieces do not all pull together.
- PLAN theme_fit=partly: It is a Commander deck under the requested Thranduil at a bracket-3-appropriate power level with Sol Ring kept, but a noticeable slice of the list (Sol Ring, Arcane Signet, Mox Amber, Languish, Night's Whisper, Elvish Mystic/Archdruid/Visionary) comes from outside the Hobbit set the person asked to build within.
- PLAN useful_as_built=partly: 36 lands, cheap ramp, a deep draw package, and plenty of removal make it functional out of the box, but a three-color deck leaning almost entirely on basics with only two rocks will stumble on colors, and its win condition is diffuse creature pressure repeatedly undercut by its own board wipes.
- PLAN summary_honest=partly: It fairly reports the sweepers, equipment, and slower speed, but calling the list "Elf-centered" oversells a creature suite where a large share of bodies (bears, bees, whales, trolls, wargs, bats) are not Elves and get no support from the tribal pieces.
- [INFO] `curve_summary`: average mana value 3.35 over 63 nonland cards
- [INFO] `mana_pass`: the builder moved 2 cards of the mana base to bring the deck inside its power level
- [WARN] `outside_requested_set`: the sets you named do not hold 1 card: Sol Ring

<details><summary>The deck list</summary>

- 9 Forest | land | Basic land for reliable green mana.
- 9 Island | land | Basic land for reliable blue mana.
- 8 Swamp | land | Basic land for reliable black mana.
- 1 Elven Passage | land | Thematic nonbasic land for the mana base.
- 1 Elvenking's Halls | land | Thematic nonbasic land for the mana base.
- 1 Hobbit Hole | land | Nonbasic land that broadens the mana base.
- 1 Minas Morgul, Dark Fortress | land | Nonbasic land for the mana base.
- 1 Mirkwood | land | Thematic nonbasic land for the mana base.
- 1 Rivendell | land | Thematic nonbasic land for the mana base.
- 1 The Black Gate | land | Nonbasic land for the mana base.
- 1 The Shire | land | Thematic nonbasic land for the mana base.
- 1 Treasure Vault | land | Artifact land that adds a mana-base slot without using a spell slot.
- 1 Sol Ring | ramp | Efficient artifact mana kept as requested.
- 1 Arcane Signet | ramp | Reliable artifact mana fixing.
- 1 Delighted Halfling | ramp | Low-cost creature ramp for an early board.
- 1 Elvish Mystic | ramp | Low-cost Elf ramp that fits the creature plan.
- 1 Elvish Archdruid | ramp | Elf-based mana development for the tribal board.
- 1 Elven Chorus | ramp | Enchantment-based mana development.
- 1 Thranduil the Strategist | ramp | Thematic legendary ramp piece.
- 1 Through the Forest Gate | ramp | Sorcery ramp to develop mana.
- 1 Wood Elves | ramp | Elf ramp that adds to the creature suite.
- 1 Woodland Weavemaster | ramp | Elf-based ramp for the midgame.
- 1 Elvish Visionary | draw | Elf card draw that keeps the creature theme dense.
- 1 Fateful Discovery | draw | Dedicated card-draw enchantment.
- 1 Hithlain Knots | draw | Instant-speed card draw.
- 1 Ithilien Kingfisher | draw | Creature-based card draw.
- 1 Key to the Side-Door | draw | Artifact card-draw option.
- 1 Lórien Revealed | draw | Dedicated sorcery card draw.
- 1 Night's Whisper | draw | Efficient dedicated card draw.
- 1 Palantír of Orthanc | draw | Legendary artifact card draw.
- 1 Plunder the Trollshaws | draw | Instant card draw for refueling.
- 1 Reverent Howl | draw | Card draw that stays on theme.
- 1 Uncover the Moon-Letters | draw | Enchantment-based card draw.
- 1 Bilbo's Ring | interaction | Equipment interaction for protecting the game plan.
- 1 Confusticate and Bebother | interaction | Instant interaction for opposing plays.
- 1 Elrond, Moon-Reader | interaction | Legendary Elf interaction that fits the theme.
- 1 Mithril Coat | interaction | Equipment interaction for a creature-focused deck.
- 1 My Precious // Allure of Power | interaction | Flexible artifact-based interaction.
- 1 Sound the Trumpets | interaction | Instant interaction option.
- 1 Stern Scolding | interaction | Low-cost instant interaction.
- 1 Thranduil's Decree | interaction | Thematic instant interaction.
- 1 Bilbo's Deadly Slice | removal | Efficient instant removal.
- 1 Bitter Downfall | removal | Dedicated instant removal.
- 1 Crude Bent Blade | removal | Artifact removal option.
- 1 Enchanted River's Grasp | removal | Enchantment-based removal.
- 1 Orcish Bowmasters | removal | Creature-based removal for the board.
- 1 Quarrel | removal | Instant removal option.
- 1 Stir Up Trouble | removal | Sorcery-speed removal.
- 1 Uneasy Partings | removal | Flexible instant removal.
- 1 Witch-king of Angmar | removal | Legendary creature removal that can also pressure opponents.
- 1 Gnashing of Teeth | wipe | Sweeper for crowded boards.
- 1 Languish | wipe | Sweeper for resetting creature-heavy positions.
- 1 Raise the Palisade | wipe | Sweeper that can clear opposing boards.
- 1 Arwen, Weaver of Hope | synergy | Elf creature chosen to reinforce the Elf-focused board.
- 1 Boughside Wanderers | synergy | Elf creature that adds to the shared tribal suite.
- 1 Cantankerous Keepers | synergy | Elf creature supporting the creature-centered theme.
- 1 Celeborn the Wise | synergy | Legendary Elf that strengthens the deck's thematic core.
- 1 Elven Raft-Steerer | synergy | Elf creature for a denser Elf board.
- 1 Galadhrim Guide | synergy | Elf creature aligned with the deck's theme.
- 1 Galion, Elvenking's Butler | synergy | Thematic legendary Elf for the commander’s court.
- 1 Grey Havens Navigator | synergy | Elf creature that supports the tribal creature plan.
- 1 Guardian of the Halls | synergy | Elf creature for the shared board plan.
- 1 Lothlórien Lookout | synergy | Elf creature that rounds out the thematic suite.
- 1 Attercop | threat | Creature threat that adds board presence.
- 1 Colossal Whale | threat | Large creature threat for closing games.
- 1 Dreaded Bat-Cloud | threat | Creature threat for sustained board pressure.
- 1 Gigantic Big Bear | threat | Creature threat that attacks from the board.
- 1 Great Fierce Bee | threat | Creature threat for combat pressure.
- 1 Haunt of the Dead Marshes | threat | Nightmare Elf threat that also fits the creature plan.
- 1 Mirkwood Elk | threat | Creature threat for the battlefield.
- 1 Nighthowl Pursuer | threat | Wolf threat for applying pressure.
- 1 Ravening Warg | threat | Wolf creature threat for combat.
- 1 Rhovanion Rampager | threat | Wolf threat that broadens the creature base.
- 1 Troll of Khazad-dûm | threat | Creature threat for the top end.
- 1 Willow-Wind | threat | Elemental creature threat that helps finish games.
- 1 Mox Amber | synergy | the mana pass added it to bring the mana base inside the power level

</details>

### 23. two set families at once

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 203 names.

Commander: Kíli the Resourceful.

Grade: typical, score 0.52, model 20260907T081149Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $250.85 to buy, $250.85 the whole deck.

**Summary:** This is a white Dwarf-and-Equipment Commander deck built to establish a creature board, develop it with a dense artifact package, and keep resources flowing through creature and artifact draw. Kíli is backed by Fíli, Dáin, Ori, and a suite of thematic blades, while versatile removal and several reset buttons keep opposing boards manageable. The deck wins by maintaining a developed creature board through interaction and turning its equipped threats sideways; it gives up multicolor flexibility and relies on its artifacts and creatures remaining in play to sustain its momentum.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=partly: The core is a legible mono-white equipment/artifact go-wide build with draw and removal, but the three symmetrical wraths (Martial Coup, Promise of Loyalty, Dusk // Dawn) and the changeling-typal pieces (Maskwood Nexus, Three Tree Mascot, Patchwork Banner) pull against the stated plan of keeping a developed creature board equipped and attacking.
- PLAN theme_fit=partly: The commander choice is from the requested Hobbit set and a genuine Bloomburrow slice is present (Fountainport, Carrot Cake, Caretaker's Talent, Sword of the Squeak, Bumbleflower's Sharepot), but a large fraction of the deck is generic or Tales of Middle-earth/other-set material (Swords to Plowshares, Skullclamp, Blade Splicer, Skyclave Apparition, Karn, Andúril, Glamdring), so it only loosely honors the two-set restriction.
- PLAN useful_as_built=yes: 36 lands in a single color plus ten ramp pieces, a smooth curve, plenty of removal and card draw, and a clear combat-damage win with equipped creatures make it straightforward to pick up and play.
- PLAN summary_honest=partly: It fairly describes the equipment/artifact/draw shell and admits the mono-color tradeoff, but calling the sweepers cards that "keep opposing boards manageable" glosses over that they wipe the deck's own board, the "Dwarf" billing rests on only a handful of actual Dwarves, and it never notes that much of the list comes from outside the two sets the player named.
- [INFO] `curve_summary`: average mana value 2.71 over 63 nonland cards

<details><summary>The deck list</summary>

- 21 Plains | land | Provides a reliable untapped white mana base.
- 1 Castle Ardenvale | land | Adds a white-producing land slot.
- 1 Command Tower | land | Provides dependable commander-color mana.
- 1 Evolving Wilds | land | Provides flexible access to Plains.
- 1 Exotic Orchard | land | Adds flexible colored mana.
- 1 Fabled Passage | land | Provides another flexible Plains-finding land.
- 1 Fountainport | land | Adds a utility land without reducing the Plains count too far.
- 1 Hobbit Hole | land | Adds a thematic white land.
- 1 Minas Tirith | land | Provides a thematic white mana source.
- 1 Path of Ancestry | land | Adds colored mana for the creature-focused deck.
- 1 Reliquary Tower | land | Provides a utility land slot.
- 1 Rogue's Passage | land | Gives the creature board a utility land option.
- 1 Terramorphic Expanse | land | Provides additional access to Plains.
- 1 Thriving Heath | land | Adds another white-producing land.
- 1 Treasure Vault | land | Adds a thematic artifact land.
- 1 Uncharted Haven | land | Provides flexible colored mana.
- 1 Caretaker's Talent | draw | Supplies repeatable card advantage.
- 1 Circuit Mender | draw | Provides card advantage on an artifact creature.
- 1 Cut a Deal | draw | Adds a straightforward card-advantage spell.
- 1 Feather of Flight | draw | Supports the creature plan while providing a card.
- 1 Fountainport Bell | draw | Provides an artifact-based draw outlet.
- 1 Heirloom Epic | draw | Adds a thematic artifact draw piece.
- 1 Idol of Oblivion | draw | Provides efficient artifact-based card advantage.
- 1 Inspiring Overseer | draw | Adds card advantage on a creature.
- 1 Mangara, the Diplomat | draw | Provides a durable source of card advantage.
- 1 Mentor of the Meek | draw | Rewards the deck's smaller creatures with cards.
- 1 Skullclamp | draw | Turns expendable creatures into cards.
- 1 Bilbo's Ring | interaction | Adds a protective artifact option.
- 1 Dawn's Truce | interaction | Provides a defensive instant-speed option.
- 1 Flowering of the White Tree | interaction | Supports the legendary creature board.
- 1 Luminous Broodmoth | interaction | Helps preserve the creature board.
- 1 Perch Protection | interaction | Provides a flexible defensive response.
- 1 Reprieve | interaction | Offers a flexible instant-speed answer.
- 1 Selfless Spirit | interaction | Protects the creature board from opposing answers.
- 1 Swiftfoot Boots | interaction | Protects an important creature or commander.
- 1 Arcane Signet | ramp | Provides efficient colored mana acceleration.
- 1 Burnished Hart | ramp | Provides creature-based land acceleration.
- 1 Fellwar Stone | ramp | Adds efficient mana acceleration.
- 1 Loyal Warhound | ramp | Provides early creature-based ramp.
- 1 Mind Stone | ramp | Accelerates early development.
- 1 Mox Amber | ramp | Provides inexpensive legendary-focused acceleration.
- 1 Ornithopter of Paradise | ramp | Provides fixing on an artifact creature.
- 1 Patchwork Banner | ramp | Supports the creature theme while accelerating mana.
- 1 Sol Ring | ramp | Provides efficient artifact acceleration.
- 1 Wayfarer's Bauble | ramp | Provides early land-based acceleration.
- 1 Angel of the Ruins | removal | Adds a large artifact creature that answers opposing permanents.
- 1 Banishing Light | removal | Provides a broad permanent answer.
- 1 Bumbleflower's Sharepot | removal | Adds thematic artifact-based removal.
- 1 Fiend Hunter | removal | Places removal on a creature body.
- 1 Generous Gift | removal | Answers a wide range of opposing permanents.
- 1 Giant's Boulder | removal | Provides a thematic artifact removal option.
- 1 Loran of the Third Path | removal | Adds removal on a legendary creature.
- 1 Skyclave Apparition | removal | Provides efficient creature-based removal.
- 1 Swords to Plowshares | removal | Provides a highly efficient creature answer.
- 1 Blade Splicer | synergy | Supports the artifact-creature and token-oriented board.
- 1 Carrot Cake | synergy | Adds a thematic artifact synergy piece.
- 1 Dáin, Lord of the Iron Hills | synergy | Supports the Dwarf-focused creature plan.
- 1 Iron Hills Blacksmith | synergy | Supports the deck's Equipment and Dwarf themes.
- 1 Ori, Keeper of Songs | synergy | Adds another Dwarf-focused synergy card.
- 1 Tangle Tumbler | synergy | Supports the artifact-focused board plan.
- 1 Fíli the Pathfinder | threat | Serves as a named legendary threat alongside Kíli.
- 1 Karn, the Great Creator | threat | Provides a resilient colorless threat.
- 1 Dusk // Dawn | wipe | Provides a flexible board-reset option.
- 1 Martial Coup | wipe | Resets the board while fitting the creature plan.
- 1 Promise of Loyalty | wipe | Provides a board-control reset.
- 1 Academy Manufactor | other | Adds a colorless artifact creature to develop the board.
- 1 Andúril, Flame of the West | other | Adds a thematic legendary Equipment.
- 1 Andúril, Narsil Reforged | other | Adds a second thematic legendary Equipment.
- 1 Dúnedain Blade | other | Adds another Equipment for the creature board.
- 1 Dwarven Shortsword | other | Adds a low-cost thematic Equipment.
- 1 Glamdring | other | Adds a legendary Equipment to the artifact package.
- 1 Helm of the Host | other | Provides a powerful Equipment centerpiece.
- 1 Long-Lost Lances | other | Adds another thematic Equipment.
- 1 Maskwood Nexus | other | Adds an artifact that supports the creature themes.
- 1 Starforged Sword | other | Adds an Equipment for the creature plan.
- 1 Sting, Bilbo's Sword | other | Adds a thematic legendary Equipment.
- 1 Sword of the Squeak | other | Adds a Bloomburrow-themed Equipment.
- 1 Sword of Vengeance | other | Adds another Equipment payoff for the board.
- 1 Three Tree Mascot | other | Adds a colorless artifact creature to the board plan.

</details>

### 24. a 60-card deck from one set

Format: Modern. Theme: aggro. Pool: any_card. Shortlist: 96 names.

Grade: baseline, score 0.35, model 20260907T081149Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $22.00 to buy, $22.00 the whole deck.

**Summary:** This mono-red Bloomburrow aggro deck presses early with a Mouse and Lizard creature core, then keeps attacking behind a dense package of threats. Removal clears the way for combat, while the draw cards help sustain pressure once the first wave has been answered. It aims to win through repeated creature attacks and gives up broad answers and late-game power for a focused, fast plan.

The quality model grades this deck at the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. The color sources cover the pips, and that raises the grade.

- PLAN plan_coherent=partly: The creature core, burn removal and cheap draw all point at one aggressive plan, but 24 lands plus two Arcane Signets is ramp-flavored filler that pulls against a curve topping at three.
- PLAN theme_fit=partly: It is mono-red aggro legal for Modern with a Bloomburrow creature and spell core, but Abrade and Arcane Signet come from outside the requested set, and Signet in particular is not an aggro card.
- PLAN useful_as_built=yes: Twenty-four Mountains cast everything, twenty-two cheap creatures plus burn give plenty of ways to close a game, and a casual player could shuffle it up and play immediately despite the excess lands.
- PLAN summary_honest=yes: The summary accurately describes the Mouse/Lizard creature core, removal, and draw package and honestly admits the deck trades late-game power and broad answers for speed.
- [INFO] `curve_summary`: average mana value 2.89 over 36 nonland cards

<details><summary>The deck list</summary>

- 24 Mountain | land | Provides the consistent mono-red mana base.
- 2 Arcane Signet | ramp | Fills the requested acceleration slots without straining the mana base.
- 4 Might of the Meek | draw | Provides efficient card flow alongside the deck's Meek creature theme.
- 2 Sazacap's Brew | draw | Adds more card flow to keep the pressure supplied.
- 4 Abrade | removal | Provides flexible low-cost removal.
- 2 Blooming Blast | removal | Completes the removal suite with another red answer.
- 4 Emberheart Challenger | synergy | Forms a core Bloomburrow Mouse synergy package.
- 4 Hearthborn Battler | synergy | Forms a core Bloomburrow Lizard synergy package.
- 4 Frilled Sparkshooter | threat | Supplies a full set of aggressive creature threats.
- 4 Reptilian Recruiter | threat | Adds another full set of Lizard threats.
- 4 Teapot Slinger | threat | Provides an additional aggressive threat package.
- 2 Stormsplitter | threat | Rounds out the threat suite with Otter Wizard threats.

</details>

### 25. use no card of an owned precon

Format: Commander. Theme: Avengers superheroes. Pool: owned_first. Shortlist: 152 names.

Commander: Syr Konrad, the Grim.

Grade: bad, score 0.34, model 20260907T081149Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $280.53 the whole deck.

**Summary:** This is a mono-black Avengers-flavored attrition deck built around Syr Konrad, the Grim. It develops mana early, uses equipment and card advantage to keep resources flowing, then trades creatures and clears crowded boards while Syr Konrad turns that attrition into pressure. Its victories come from sustained graveyard-focused damage backed by a roster of large creature threats; in exchange, it favors incremental setup and board control over explosive early finishes.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- PLAN plan_coherent=yes: The list is a consistent mono-black attrition build: ramp, heavy spot removal, sweepers and a -1/-1 counter subtheme that repeatedly kills creatures, all feeding Syr Konrad's death/mill triggers, with equipment mostly serving protection and card draw rather than pulling toward a competing voltron plan.
- PLAN theme_fit=no: The player asked for an Avengers superheroes deck, but the commander is Syr Konrad and only a handful of cards (Baron Strucker, M.O.D.O.K., Stolen Stark Tech) are even Marvel — and those are villains, not Avengers — so this answers a different request despite plausibly meeting the bracket-3 and no-precon constraints.
- PLAN useful_as_built=yes: 37 lands with strong fixing, ample cheap ramp, a low-curve interaction suite and a dozen real threats plus Syr Konrad as an inevitable damage engine make this straightforwardly playable out of the box.
- PLAN summary_honest=no: The mechanical description matches the cards, but calling it an "Avengers-flavored" deck asserts a flavor identity the list does not have and quietly hides that nearly every card is LOTR, Final Fantasy, Avatar or generic fantasy rather than Marvel heroes.
- [WARN] `not_owned`: Syr Konrad, the Grim: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.19 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Swamp | land | Primary black mana source.
- 1 Command Tower | land | Flexible mana source for the commander deck.
- 1 Castle Locthwain | land | Black-producing utility land.
- 1 Takenuma, Abandoned Mire | land | Black-producing utility land.
- 1 Ifnir Deadlands | land | Black-producing utility land.
- 1 Marsh Flats | land | Fetch land for the mana base.
- 1 Fabled Passage | land | Fetch land for the mana base.
- 1 Evolving Wilds | land | Fetch land for the mana base.
- 1 Terramorphic Expanse | land | Fetch land for the mana base.
- 1 Exotic Orchard | land | Flexible mana source.
- 1 City of Brass | land | Flexible mana source.
- 1 Secluded Courtyard | land | Creature-focused mana source.
- 1 Spire of Industry | land | Artifact-friendly mana source.
- 1 War Room | land | Utility land for the long game.
- 1 Reliquary Tower | land | Utility land for drawn cards.
- 1 Rogue's Passage | land | Utility land that helps a threat connect.
- 1 Sol Ring | ramp | Efficient early acceleration.
- 1 Arcane Signet | ramp | Reliable early acceleration.
- 1 Wayfarer's Bauble | ramp | Early mana development.
- 1 Thought Vessel | ramp | Mana development with later utility.
- 1 Bender's Waterskin | ramp | Artifact-based acceleration.
- 1 Blitzball | ramp | Artifact-based acceleration.
- 1 Commander's Sphere | ramp | Mana development with later utility.
- 1 Deadly Dispute | ramp | Turns expendable material into resources.
- 1 Sword of the Animist | ramp | Equipment that advances mana development.
- 1 Springleaf Drum | ramp | Low-cost mana acceleration.
- 1 Baron Strucker, HYDRA Overlord | draw | Avengers-themed source of card advantage.
- 1 Buster Sword | draw | Equipment-based card advantage.
- 1 Call of the Ring | draw | Ongoing card advantage.
- 1 Dusk Urchins | draw | Creature-based card advantage.
- 1 Grave Venerations | draw | Card advantage for the attrition plan.
- 1 Idol of Oblivion | draw | Low-cost card advantage.
- 1 Lembas | draw | Artifact-based card advantage.
- 1 Mask of Memory | draw | Equipment-based card selection.
- 1 Nasty End | draw | Efficient card advantage.
- 1 Night's Whisper | draw | Efficient card advantage.
- 1 Skullclamp | draw | Efficient card advantage around expendable creatures.
- 1 Champion's Helm | interaction | Protects an important creature.
- 1 Darksteel Plate | interaction | Protects an important creature.
- 1 Lightning Greaves | interaction | Protects and supports key creatures.
- 1 Swiftfoot Boots | interaction | Protects and supports key creatures.
- 1 Stolen Stark Tech | interaction | Avengers-themed protective equipment.
- 1 Orcish Medicine | interaction | Flexible response for opposing plays.
- 1 The Walls of Ba Sing Se | interaction | Defensive permanent for the board.
- 1 Stone of Erech | interaction | Utility interaction for opposing graveyard plans.
- 1 Bitter Downfall | removal | Clean answer to a problematic permanent.
- 1 Bitter Triumph | removal | Flexible answer to a problematic creature.
- 1 Claim the Precious | removal | Reliable creature answer.
- 1 Dismember | removal | Efficient creature answer.
- 1 Epic Downfall | removal | Answer for a major opposing threat.
- 1 Fatal Push | removal | Low-cost creature answer.
- 1 Heartless Act | removal | Flexible creature answer.
- 1 Infernal Grasp | removal | Unconditional creature answer.
- 1 Overkill | removal | Additional creature answer.
- 1 Black Sun's Zenith | wipe | Scalable mass-removal option.
- 1 Extinction Event | wipe | Mass-removal option against creature boards.
- 1 Game Over | wipe | Additional board-reset option.
- 1 Blowfly Infestation | synergy | Supports Syr Konrad's attrition-focused plan.
- 1 Contagion Clasp | synergy | Supports the deck's counter-based attrition tools.
- 1 Grim Poppet | synergy | Creature-based support for the attrition plan.
- 1 Incremental Blight | synergy | Supports the deck's counter-based attrition tools.
- 1 Liliana, Death Wielder | synergy | Supports the deck's long-game attrition plan.
- 1 Nameless Inversion | synergy | Supports creature-focused graveyard play.
- 1 Orcish Bowmasters | synergy | Creature-based support for the attrition plan.
- 1 Summon: Anima | synergy | Supports the deck's creature-focused plan.
- 1 The Black Breath | synergy | Supports the deck's attrition plan.
- 1 Obsessive Pursuit | synergy | Supports the deck's long-game plan.
- 1 Archfiend of Ifnir | threat | Large creature that pressures opponents.
- 1 Carnifex Demon | threat | Large creature that pressures opponents.
- 1 M.O.D.O.K. | threat | Avengers-themed creature threat.
- 1 Midnight Banshee | threat | Creature threat for the attrition plan.
- 1 Soul Snuffers | threat | Creature threat that fits the attrition shell.
- 1 Zodiark, Umbral God | threat | High-impact creature threat.
- 1 Rat King, Pale Piper | threat | Creature threat for the long game.
- 1 Massacre Girl, Known Killer | threat | Creature threat that rewards attrition.
- 1 Witch-king of Angmar | threat | Creature threat with a commanding board presence.
- 1 Ringwraiths | threat | Creature threat for the attacking plan.
- 1 Skinrender | threat | Creature threat that contributes to attrition.
- 1 Voracious Fell Beast | threat | Creature threat for closing games.
- 20 Swamp | land | Primary black mana source.

</details>

