# PR-8 deck gate

Run date: 2026-09-04. Card snapshot: 2026-09-03.

Verdict: FAIL. 23 of 24 decks passed every block check, 0 invented names reached the user, and 0 summaries stated a false rule of the game. All three bars are zero tolerance.

## Summary

| Measure | Value |
|---|---|
| Prompts | 24 |
| Decks returned | 24 |
| Decks with no block finding | 23 |
| Invented names that reached the user | 0 |
| Decks that needed the repair turn | 9 |
| Summaries judged (F-26) | 24 |
| Summaries that state a rule of the game | 1 |
| Summaries that state a FALSE rule | 0 |
| Judge errors | 0 |
| Errors | 0 |
| Prompt version | 12 |
| Calls | 62 |
| Cost | $2.5432 |
| Time | 2081 seconds |

## Findings by code

A block stops the deck. A warning and a note are reports.

| Code | Count |
|---|---|
| `not_owned` | 27 |
| `curve_summary` | 23 |
| `profile_off_band` | 13 |
| `outside_requested_set` | 2 |
| `basics_added` | 1 |
| `land_count` | 1 |
| `precon_share` | 1 |
| `two_card_combo` | 1 |
| `deck_size` | 1 |

By severity: BLOCK 1. WARN 45. INFO 24. 

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

## Decks

### 1. lifegain Commander, any card

Format: Commander. Theme: lifegain. Pool: any_card. Shortlist: 303 names.

Commander: Karlov of the Ghost Council.

Grade: bad, score 0.01, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band. Block findings: 0.

Cost: $555.72 to buy, $555.72 the whole deck.

**Summary:** This Orzhov lifegain strategy builds a steady stream of small life-gain events to grow its commander into a lethal attacker while turning those same triggers into cards, tokens, counters, and opponent life loss. A deep creature suite supplies durable pressure, while efficient answers and sweeping effects keep opposing boards under control. The mana base strongly supports black-intensive spells without sacrificing dependable white access for the deck's lifegain engines.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 3.27 over 63 nonland cards
- [WARN] `profile_off_band`: the mana available on turn four is 4.19, and bracket 3 wants 4.2 or more (mean over 10000 hands)
- [WARN] `two_card_combo`: bracket 3 allows no two-card infinite combo that needs four mana or less, and the deck holds Enduring Angel // Angelic Enforcer + Enduring Tenacity (speed 5, near two-card); Enduring Angel // Angelic Enforcer + Vito, Thorn of the Dusk Rose (speed 5, near two-card)

<details><summary>The deck list</summary>

- 1 Archivist of Oghma | draw | Draws cards while rewarding opponents for searching.
- 1 Dawn of Hope | draw | Turns repeated life gain into cards and creature tokens.
- 1 Enduring Innocence | draw | Provides steady card advantage from small creatures entering.
- 1 Mangara, the Diplomat | draw | Punishes opponents who overextend with cards.
- 1 Markov Purifier | draw | Draws from the deck's Vampire and lifegain plan.
- 1 Sigarda's Splendor | draw | Rewards maintaining a high life total with cards.
- 1 Survival Cache | draw | Efficient life-gain card draw spell.
- 1 The Gaffer | draw | Supplies cards while the life total stays high.
- 1 Tymna the Weaver | draw | Converts evasive combat pressure into cards.
- 1 Vampiric Rites | draw | Sacrifice outlet that turns expendable creatures into cards and life.
- 1 Well of Lost Dreams | draw | Converts large lifegain events into substantial card draw.
- 1 Alseid of Life's Bounty | interaction | Protects key creatures or enchantments from targeted removal.
- 1 Courageous Resolve | interaction | Protects a key permanent while supporting the life-total plan.
- 1 Enduring Angel // Angelic Enforcer | interaction | Provides a resilient defensive threat and protection against losing life.
- 1 Faith's Shield | interaction | Flexible protection for the commander or a decisive attacker.
- 1 Paladin Danse, Steel Maverick | interaction | Offers flexible protection and board utility.
- 1 Restoration Magic | interaction | Protects important permanents and recovers from removal.
- 1 Sword of Light and Shadow | interaction | Protects an attacker and recurs creatures.
- 1 Werefox Bodyguard | interaction | Temporarily answers a problematic creature while advancing the board.
- 1 Altar of the Pantheon | ramp | Fixes mana and accelerates toward the deck's larger spells.
- 1 Bounty Board | ramp | Provides mana acceleration for the midgame.
- 1 Crypt Ghast | ramp | Doubles black mana production for powerful turns.
- 1 Hierophant's Chalice | ramp | Mana rock that also contributes a lifegain trigger.
- 1 Hot Dog Cart | ramp | Early mana acceleration with Food-themed utility.
- 1 Nuka-Cola Vending Machine | ramp | Creates mana resources while supporting incidental life gain.
- 1 Orazca Relic | ramp | Reliable mana rock that can later become card advantage.
- 1 Potioner's Trove | ramp | Builds Treasure resources from repeated life gain.
- 1 Pristine Talisman | ramp | Mana acceleration that triggers lifegain synergies.
- 1 The Celestus | ramp | Fixes mana and supplies recurring filtering and life gain.
- 1 Ayli, Eternal Pilgrim | removal | Sacrifice outlet and repeatable removal tied to a high life total.
- 1 Gumdrop Poisoner // Tempt with Treats | removal | Flexible creature removal with an additional resource option.
- 1 Henrika Domnathi // Henrika, Infernal Seer | removal | Versatile edict-style removal attached to a useful creature.
- 1 Murderous Rider // Swift End | removal | Efficient answer to creatures or planeswalkers that remains a creature afterward.
- 1 Nightmare's Thirst | removal | Cheap creature removal that improves with life gained.
- 1 Solitude | removal | Answers dangerous creatures immediately and can gain life.
- 1 Umezawa's Jitte | removal | Reusable combat-based removal and utility equipment.
- 1 Vona, Butcher of Magan | removal | Repeatable removal fueled by the deck's substantial life total.
- 1 Witch of the Moors | removal | Uses end-step lifegain to pressure opposing hands and boards.
- 1 Aerith Gainsborough | synergy | Rewards lifegain and supports the deck's creature-based engine.
- 1 Ajani's Pridemate | synergy | Efficient scaling threat from every lifegain event.
- 1 Cleric Class | synergy | Enhances lifegain and rewards it with permanent growth.
- 1 Heliod, Sun-Crowned | synergy | Distributes counters from lifegain and enables creature synergies.
- 1 Lurrus of the Dream-Den | synergy | Recasts inexpensive permanents that drive the lifegain engine.
- 1 Ocelot Pride | synergy | Creates a growing board from recurring life gain.
- 1 Resplendent Angel | synergy | Turns large life-gain turns into evasive Angel tokens.
- 1 Serra Ascendant | synergy | A powerful early lifelink attacker at a high life total.
- 1 Speaker of the Heavens | synergy | Uses a high life total to create a stream of Angel tokens.
- 1 Vito, Thorn of the Dusk Rose | synergy | Turns lifegain into direct pressure on opponents.
- 1 Archangel of Thune | threat | Makes every lifegain event permanently grow the entire board.
- 1 Astarion, the Decadent | threat | Converts combat damage and lifegain into major life-total swings.
- 1 Blood Baron of Vizkopa | threat | Resilient lifelink attacker that becomes dominant at high life.
- 1 Celestine, the Living Saint | threat | Recurring creatures builds overwhelming long-game value.
- 1 Enduring Tenacity | threat | Makes lifegain drain opponents while remaining hard to remove permanently.
- 1 Gideon's Company | threat | Rapidly grows into a large lifegain-fueled attacker.
- 1 Nykthos Paragon | threat | Turns a large lifegain event into a massive board-wide power boost.
- 1 Qala, Ajani's Pridemate | threat | A lifegain-scaled attacker that pressures life totals quickly.
- 1 Rhox Faithmender | threat | Doubles lifegain and supplies a substantial lifelink body.
- 1 Valkyrie Harbinger | threat | Creates recurring Angel threats from a high life total.
- 1 Vampire Scrivener | threat | Provides an efficient evasive Vampire threat.
- 1 Wurmcoil Engine | threat | Durable lifelink threat that leaves behind value after removal.
- 1 Fumigate | wipe | Resets creature-heavy boards while restoring life.
- 1 Kaya's Wrath | wipe | Efficient unconditional creature sweep that gains life.
- 1 The Battle of Bywater | wipe | Low-cost creature sweep that can preserve smaller utility creatures.
- 1 Caves of Koilos | land | Untapped dual land for both deck colors.
- 1 Cavern of Souls | land | Provides colored mana for the deck's creature-heavy Vampire plan.
- 1 City of Brass | land | Flexible untapped fixing for either color.
- 1 Command Tower | land | Reliable untapped fixing for the commander's colors.
- 1 Exotic Orchard | land | Usually provides flexible colored mana in multiplayer games.
- 1 Godless Shrine | land | Fetchable dual land that can enter untapped when needed.
- 1 Isolated Chapel | land | Dual land fixing for white and black.
- 1 Mana Confluence | land | Untapped access to either needed color.
- 1 Maestros Theater | land | Finds a Swamp while fixing early black access.
- 1 Obscura Storefront | land | Finds either basic color and supports landfall-style thinning.
- 5 Plains | land | Basic white sources that keep the mana base stable.
- 1 Path of Ancestry | land | Produces commander colors and provides creature selection.
- 1 Reflecting Pool | land | Becomes flexible fixing alongside the many dual lands.
- 1 Restless Fortress | land | Dual land that becomes a lifelink creature threat.
- 1 Scoured Barrens | land | Dual land that adds a useful lifegain trigger.
- 1 Shattered Sanctum | land | Reliable dual land in the early game.
- 12 Swamp | land | Heavy basic black base for black-intensive spells.
- 1 Temple of Silence | land | Dual land with useful card selection.
- 1 Unclaimed Territory | land | Produces colored mana for the deck's many Vampire creatures.
- 1 Vault of Champions | land | Multiplayer dual land that commonly enters untapped.
- 1 Brokers Hideout | land | Finds a basic Plains or Swamp for flexible fixing.

</details>

### 2. aristocrats Commander, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 179 names.

Commander: Denethor, Ruling Steward.

Grade: bad, score 0.36, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band. Block findings: 0.

Cost: $0.00 to buy, $186.02 the whole deck.

**Summary:** Denethor leads a black-white aristocrats strategy built around expendable creatures, recurring graveyard bodies, and death-trigger value. Protective equipment and spells keep the central engine active, while efficient removal and flexible sweepers control the table. The deck closes through resilient creature pressure, incremental life-drain sacrifice sequences, and recurring board presence.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. Few lands enter tapped, and that lowers the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 2.89 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Skullclamp | draw | Turns disposable creatures into efficient card flow.
- 1 Idol of Oblivion | draw | Draws from token-making turns and offers a late-game body.
- 1 Night's Whisper | draw | Low-cost card selection.
- 1 Painful Truths | draw | Efficient refill in the deck's colors.
- 1 Nasty End | draw | Converts a sacrificed creature into cards.
- 1 Lembas | draw | Provides immediate selection and later replaces itself.
- 1 Mask of Memory | draw | Rewards attacking creatures with repeatable filtering.
- 1 Puresteel Paladin | draw | Draws from the equipment package.
- 1 Wall of Omens | draw | Early defense that replaces itself.
- 1 Grave Venerations | draw | Supports ongoing card advantage from creature deaths.
- 1 Call of the Ring | draw | Provides steady card flow while advancing the creature plan.
- 1 Champion's Helm | interaction | Protects the commander and key legendary creatures.
- 1 Darksteel Plate | interaction | Keeps an important creature alive through damage and destruction.
- 1 Gift of Immortality | interaction | Returns a sacrificed or destroyed creature for repeated value.
- 1 Lightning Greaves | interaction | Gives immediate protection and haste.
- 1 Clever Concealment | interaction | Shields the board from hostile sweepers.
- 1 Reprieve | interaction | Temporarily answers a pivotal spell while replacing itself.
- 1 Sheltered by Ghosts | interaction | Protects a creature while disrupting an opposing permanent.
- 1 Unbreakable Formation | interaction | Preserves the creature board and supports combat pressure.
- 16 Plains | land | Reliable white mana.
- 14 Swamp | land | Reliable black mana.
- 1 Command Tower | land | Reliable access to both colors.
- 1 Exotic Orchard | land | Flexible color fixing.
- 1 City of Brass | land | Untapped fixing for both colors.
- 1 Plaza of Heroes | land | Supports the legendary commander and legends.
- 1 Minas Tirith | land | Produces white mana and supports the creature strategy.
- 1 Takenuma, Abandoned Mire | land | Black source with late-game creature recursion.
- 1 Arcane Signet | ramp | Reliable color fixing.
- 1 Sol Ring | ramp | Efficient early mana acceleration.
- 1 Fellwar Stone | ramp | Flexible mana acceleration.
- 1 Bender's Waterskin | ramp | Artifact-based mana development.
- 1 Commander's Sphere | ramp | Fixes mana and can become a card later.
- 1 Lotho, Corrupt Shirriff | ramp | Generates Treasure during opponents' spell-heavy turns.
- 1 Deadly Dispute | ramp | Sacrifices fodder for Treasure and cards.
- 1 Relic of Legends | ramp | Turns legendary creatures into mana sources.
- 1 Wayfarer's Bauble | ramp | Finds a basic land for durable acceleration.
- 1 Sword of the Animist | ramp | Converts attacks into land ramp.
- 1 Bitter Triumph | removal | Flexible creature or planeswalker answer.
- 1 Generous Gift | removal | Answers any troublesome permanent.
- 1 Get Lost | removal | Efficiently removes creatures, enchantments, or planeswalkers.
- 1 Infernal Grasp | removal | Clean creature removal.
- 1 Swords to Plowshares | removal | Highly efficient creature exile.
- 1 Claim the Precious | removal | Exiles a problematic creature.
- 1 Fiend Hunter | removal | Creature-based removal that can be sacrificed for value.
- 1 Orcish Bowmasters | removal | Punishes opposing card draw and removes small creatures.
- 1 Palace Jailer | removal | Exiles a creature while adding monarch pressure.
- 1 Gollum the Abandoned | synergy | A recursive body that rewards sacrificing creatures.
- 1 Gollum, Patient Plotter | synergy | Returns from the graveyard for repeated sacrifice value.
- 1 Gríma Wormtongue | synergy | Disrupts opposing resources while contributing a sacrifice body.
- 1 Heirloom Auntie | synergy | Supports graveyard-oriented creature value.
- 1 Nimble Hobbit | synergy | Provides a low-cost body for sacrifice-based lines.
- 1 Joo Dee, One of Many | synergy | Supports the creature-focused value plan.
- 1 Phantom Train | synergy | Provides an artifact body and recurring board presence.
- 1 Arcade Cabinet | synergy | Adds artifact-based support for the deck's creature plan.
- 1 Blowfly Infestation | synergy | Turns creature deaths into cascading board control.
- 1 Gorbag of Minas Morgul | synergy | Provides sacrifice-friendly creature value and mana support.
- 1 Bill the Pony | threat | Efficient creature pressure that contributes to the board.
- 1 Vengeful Villagers | threat | Builds a resilient creature presence.
- 1 Rat King, Pale Piper | threat | Creates a threatening board presence with graveyard value.
- 1 Angel of Serenity | threat | A high-impact flying finisher with immediate board influence.
- 1 Archfiend of Ifnir | threat | A large evasive threat that suppresses opposing creatures.
- 1 Massacre Girl, Known Killer | threat | Turns opposing creature losses into sustained pressure.
- 1 Witch-king of Angmar | threat | A resilient evasive finisher that punishes attacks.
- 1 The Sackville-Bagginses | threat | Provides a substantial legendary creature threat.
- 1 Frontline Medic | threat | Adds combat pressure while protecting an attack.
- 1 Bronze Guardian | threat | A scalable artifact creature that is difficult to answer.
- 1 Bastion Protector | threat | Strengthens the commander while adding a solid body.
- 1 Exemplar of Light | threat | Provides an evasive threat with value potential.
- 1 Austere Command | wipe | Flexible reset that can spare the preferred board pieces.
- 1 Black Sun's Zenith | wipe | Scalable creature sweep that weakens indestructible boards.
- 1 Dusk // Dawn | wipe | Clears larger creatures and can recover a creature-heavy graveyard.

</details>

### 3. artifacts Commander, bracket 4

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 299 names.

Commander: Urza, Lord High Artificer.

Grade: bad, score 0.03, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band, profile_off_band. Block findings: 0.

Cost: $644.90 to buy, $644.90 the whole deck.

**Summary:** Urza leads a fast mono-blue artifact strategy built to establish mana engines, deploy a dense board of artifacts, and turn that material advantage into overwhelming combat pressure. The deck combines efficient card flow, broad reactive coverage, artifact cost reduction, and powerful permanent-based engines, allowing it to pivot between protecting a dominant board, disrupting opposing plans, and assembling explosive finishing turns.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The commander leads high-bracket decks, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 3.18 over 66 nonland cards
- [WARN] `profile_off_band`: the average mana value of the nonland cards is 3.18, and bracket 4 wants 1.6 to 3
- [WARN] `profile_off_band`: the mana available on turn four is 4.31, and bracket 4 wants 4.8 or more (mean over 10000 hands)

<details><summary>The deck list</summary>

- 25 Island | land | Reliable blue sources for casting spells and supporting Urza.
- 1 Seat of the Synod | land | Blue source that also raises the artifact count.
- 1 Mystic Sanctuary | land | Blue source with late-game instant or sorcery recursion.
- 1 Otawara, Soaring City | land | Blue source and flexible channel-based disruption.
- 1 Mistrise Village | land | Blue source that helps key spells resolve through opposing interaction.
- 1 Glimmervoid | land | Reliable colored source in an artifact-heavy deck.
- 1 Command Tower | land | Untapped blue source for the commander’s identity.
- 1 Urza's Saga | land | Artifact land that develops threats and finds inexpensive artifacts.
- 1 Inventors' Fair | land | Artifact-focused utility land with a powerful late-game search effect.
- 1 Mox Opal | ramp | Fast artifact-enabled mana acceleration.
- 1 Moonsnare Prototype | ramp | Low-cost mana production that turns artifacts into acceleration.
- 1 Arc Reactor | ramp | Artifact-based mana acceleration for early development.
- 1 Coin of Mastery | ramp | Artifact mana source that advances the board early.
- 1 Crystal Skull, Isu Spyglass | ramp | Mana artifact with useful selection attached.
- 1 Chief Engineer | ramp | Convoke-style artifact acceleration for large artifact turns.
- 1 Grand Architect | ramp | Turns blue creatures into substantial artifact mana.
- 1 Inspiring Statuary | ramp | Lets artifact permanents help cast nonartifact spells.
- 1 Metalworker | ramp | High-output artifact mana engine.
- 1 Solar Array | ramp | Artifact mana production that supports large turns.
- 1 Slagstone Refinery | ramp | Artifact-based acceleration for deploying the deck’s engines.
- 1 Vedalken Engineer | ramp | Early artifact mana creature.
- 1 Urza's Command | ramp | Flexible spell that supplies mana development when needed.
- 1 Braided Net // Braided Quipu | draw | Artifact card advantage that works with the deck’s permanent density.
- 1 Cerebral Download | draw | Efficient instant-speed card selection and advantage.
- 1 Era of Innovation | draw | Artifact and creature activity converts into cards.
- 1 Esoteric Duplicator | draw | Artifact-based value engine that provides cards over time.
- 1 Forensic Gadgeteer | draw | Artifact synergy creature that supplies card advantage.
- 1 Katara, Bending Prodigy | draw | Low-cost source of recurring card flow.
- 1 Riddlesmith | draw | Filters draws whenever artifacts are deployed.
- 1 Sai, Master Thopterist | draw | Artifact deployment creates fliers and can be converted into cards.
- 1 Senator Peacock | draw | Persistent artifact-oriented card advantage.
- 1 Spirit Water Revival | draw | Efficient spell-based card advantage.
- 1 Thirst for Knowledge | draw | Instant-speed digging that rewards artifact cards.
- 1 Thoughtcast | draw | Very efficient card draw in an artifact-heavy shell.
- 1 Metallic Rebuke | interaction | Cheap artifact-enabled counterspell.
- 1 Disruption Protocol | interaction | Countermagic that can leverage an artifact for its additional cost.
- 1 Ice Out | interaction | Flexible counterspell that improves with a developed board.
- 1 Reality Ripple | interaction | Tempo interaction that answers problematic permanents temporarily.
- 1 Welding Jar | interaction | Zero-cost protection for important artifacts.
- 1 Curator's Ward | interaction | Protects a major artifact permanent while replacing itself.
- 1 Escape Protocol | interaction | Protects artifacts and reuses enter-the-battlefield value.
- 1 Ghostly Flicker | interaction | Instant-speed protection and value reset for key permanents.
- 1 Override | interaction | Artifact-scaled counterspell for defending a winning board.
- 1 Stoic Rebuttal | interaction | Reliable hard counterspell with artifact upside.
- 1 Waterbender's Restoration | interaction | Versatile instant-speed defensive interaction.
- 1 Aether Spellbomb | removal | Cheap artifact removal that can also cycle.
- 1 Contagion Clasp | removal | Removes small creatures and supports counter-based artifacts.
- 1 Cyber Conversion | removal | Efficient answer to a problematic permanent.
- 1 Into Thin Air | removal | Flexible bounce removal at instant speed.
- 1 Resculpt | removal | Exiles an artifact or creature at a low cost.
- 1 Ravenform | removal | Exile-based answer for artifacts and creatures.
- 1 Unable to Scream | removal | Cheap enchantment-based neutralization for commanders and creatures.
- 1 Watery Grasp | removal | Permanent-based creature suppression.
- 1 Zuko's Exile | removal | Low-cost removal option for key opposing permanents.
- 1 Water Whip | removal | Flexible removal spell that handles creatures cleanly.
- 1 Engineered Explosives | wipe | Scalable, low-cost board control.
- 1 Hurkyl's Recall | wipe | Instant-speed mass reset that can create a decisive tempo swing.
- 1 Etherium Sculptor | synergy | Reduces the cost of the deck’s extensive artifact suite.
- 1 Foundry Inspector | synergy | Additional artifact cost reduction for explosive turns.
- 1 Emry, Lurker of the Loch | synergy | Recurs inexpensive artifacts and rewards a dense artifact board.
- 1 Mystic Forge | synergy | Lets artifact-heavy turns continue from the top of the library.
- 1 Clock of Omens | synergy | Converts spare artifacts into repeated untaps and combo-like value.
- 1 Whir of Invention | synergy | Finds the artifact engine or answer best suited to the situation.
- 1 Canoptek Scarab Swarm | threat | Artifact threat that scales with the game and pressures opponents.
- 1 Cyberdrive Awakener | threat | Turns a developed artifact board into an immediate lethal attack.
- 1 Kappa Cannoneer | threat | Large, evasive artifact payoff with strong protection.
- 1 Karn, Scion of Urza | threat | Produces artifact bodies and converts artifacts into major combat pressure.
- 1 Master Transmuter | threat | Deploys expensive artifacts efficiently while protecting them from removal.
- 1 Phyrexian Metamorph | threat | Copies the strongest artifact or creature on the battlefield.
- 1 Kuldotha Forgemaster | threat | Transforms expendable artifacts into the deck’s most powerful permanents.
- 1 Traxos, Scourge of Kroog | threat | Efficient large artifact attacker that untaps through normal development.
- 1 Threefold Thunderhulk | threat | Creates a substantial artifact board and scales into a finisher.
- 1 Whirler Rogue | threat | Produces artifact tokens and gives a large attacker reliable evasion.
- 1 Ironheart, Clever Champion | threat | Artifact creature payoff that provides meaningful board pressure.
- 1 Myr Enforcer | threat | Large affinity threat that becomes inexpensive in the artifact shell.

</details>

### 4. dinosaur tribal, bracket 2

Format: Commander. Theme: dinosaurs. Pool: any_card. Shortlist: 269 names.

Commander: Gishath, Sun's Avatar.

Grade: bad, score 0.25, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $381.91 to buy, $381.91 the whole deck.

**Summary:** This Gishath deck builds its mana early, develops a board of Dinosaur creatures, and uses steady card flow to keep deploying threats. It aims to win through combat with a growing Dinosaur board and its powerful top-end creatures, while keeping removal, protection, and board wipes available when the table becomes difficult. It gives up some speed and stack-based disruption in exchange for a direct, creature-forward plan that is easy to sequence and learn.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The curve sits high for the format, and that lowers the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 3.73 over 62 nonland cards

<details><summary>The deck list</summary>

- 9 Forest | land | Basic green source for the deck’s Dinosaur-heavy mana needs.
- 5 Plains | land | Basic white source for the deck’s three-color mana base.
- 5 Mountain | land | Basic red source for the deck’s three-color mana base.
- 1 Command Tower | land | Flexible three-color land for consistently casting the deck’s spells.
- 1 Exotic Orchard | land | Flexible multicolor land for the mana base.
- 1 Path of Ancestry | land | Tribal-focused land that supports the Dinosaur plan.
- 1 Unclaimed Territory | land | Creature-type-focused land for casting Dinosaurs.
- 1 Secluded Courtyard | land | Creature-type-focused land for casting Dinosaurs.
- 1 Cavern of Souls | land | Creature-focused land for the Dinosaur creature base.
- 1 City of Brass | land | Flexible multicolor source for casting across all three colors.
- 1 Mana Confluence | land | Flexible multicolor source for the deck’s demanding mana costs.
- 1 Bountiful Promenade | land | Green-white dual land for the mana base.
- 1 Spire Garden | land | Red-green dual land for the mana base.
- 1 Spectator Seating | land | Red-white dual land for the mana base.
- 1 Brushland | land | Green-white dual land for early colored mana.
- 1 Karplusan Forest | land | Red-green dual land for early colored mana.
- 1 Battlefield Forge | land | Red-white dual land for early colored mana.
- 1 Sacred Foundry | land | Red-white dual land that improves color access.
- 1 Stomping Ground | land | Red-green dual land that improves color access.
- 1 Temple Garden | land | Green-white dual land that improves color access.
- 1 Canopy Vista | land | Green-white land that rounds out the color base.
- 1 Beast Whisperer | draw | Creature-based draw that fits a deck full of Dinosaurs.
- 1 Garruk's Uprising | draw | Draw support that fits the deck’s creature-focused plan.
- 1 Guardian Project | draw | Reliable draw support for a creature-heavy game plan.
- 1 Harmonize | draw | Straightforward draw to help a new player keep resources flowing.
- 1 Return of the Wildspeaker | draw | Flexible draw spell that suits a Dinosaur board.
- 1 Ripjaw Raptor | draw | Dinosaur draw option that contributes to the creature plan.
- 1 Kutzil, Malamet Exemplar | draw | Creature-based draw support for the combat-focused deck.
- 1 Demand Answers | draw | Low-cost draw option for smoothing out a game.
- 1 Faithless Looting | draw | Efficient draw option for finding more action.
- 1 Thrill of Possibility | draw | Quick draw support that helps refresh the hand.
- 1 Sylvan Library | draw | Early draw support for finding lands and Dinosaur spells.
- 1 Sol Ring | ramp | Fast early ramp for getting the deck’s larger spells onto the table.
- 1 Arcane Signet | ramp | Reliable multicolor ramp for the three-color deck.
- 1 Birds of Paradise | ramp | Early creature ramp that fixes all of the deck’s colors.
- 1 Cultivate | ramp | Straightforward land ramp for a new-player-friendly mana plan.
- 1 Farseek | ramp | Early land ramp that helps assemble the color base.
- 1 Nature's Lore | ramp | Efficient land ramp for reaching larger Dinosaur spells.
- 1 Three Visits | ramp | Efficient land ramp for reaching the top end.
- 1 Thunderherd Migration | ramp | Dinosaur-themed ramp that supports the tribal plan.
- 1 Drover of the Mighty | ramp | Creature ramp that fits alongside the Dinosaur creature base.
- 1 Huatli, Poet of Unity // Roar of the Fifth People | ramp | Ramp option that supports the deck’s long-game development.
- 1 Bronzebeak Foragers | removal | Dinosaur removal that keeps interaction attached to the creature plan.
- 1 Burning Sun's Avatar | removal | Dinosaur removal that remains a meaningful creature threat.
- 1 Itzquinth, Firstborn of Gishath | removal | Low-cost Dinosaur removal for early board control.
- 1 Needletooth Raptor | removal | Dinosaur removal that supports the tribal creature suite.
- 1 Ravenous Sailback | removal | Dinosaur removal for answering troublesome permanents.
- 1 Savage Stomp | removal | Efficient Dinosaur-themed removal for clearing blockers.
- 1 Thrashing Brontodon | removal | Flexible Dinosaur removal attached to a creature.
- 1 Akroma's Will | interaction | Versatile interaction for protecting the board during decisive turns.
- 1 Boros Charm | interaction | Low-cost interaction that helps safeguard the game plan.
- 1 Heroic Intervention | interaction | Protection interaction for preserving a developed Dinosaur board.
- 1 Lightning Greaves | interaction | Protection equipment for an important Dinosaur or commander.
- 1 Swiftfoot Boots | interaction | Protection equipment that supports key creatures.
- 1 Temple Altisaur | interaction | Dinosaur interaction that supports the creature-focused board.
- 1 Austere Command | wipe | Flexible board wipe for resetting difficult game states.
- 1 Blasphemous Act | wipe | Board wipe for recovering when opponents build larger boards.
- 1 Raging Swordtooth | wipe | Dinosaur board wipe that remains part of the tribal plan.
- 1 Commune with Dinosaurs | synergy | Low-cost Dinosaur synergy piece that improves the tribal plan.
- 1 Armored Kincaller | synergy | Early Dinosaur synergy creature for developing the board.
- 1 Belligerent Yearling | synergy | Low-cost Dinosaur synergy creature for the early game.
- 1 Dinosaur Stampede | synergy | Dinosaur synergy spell that supports combat-focused turns.
- 1 Deathgorge Scavenger | synergy | Dinosaur synergy creature that adds utility to the board.
- 1 Dromosaur | synergy | Dinosaur synergy creature that advances the tribal theme.
- 1 Hunting Velociraptor | synergy | Dinosaur synergy creature for building toward larger plays.
- 1 Kinjalli's Caller | synergy | Tribal synergy creature that supports casting the Dinosaur suite.
- 1 Marauding Raptor | synergy | Dinosaur synergy creature for an aggressive creature plan.
- 1 Otepec Huntmaster | synergy | Tribal support creature for the Dinosaur-focused deck.
- 1 Priest of the Wakening Sun | synergy | Early tribal support creature for the Dinosaur plan.
- 1 Sunfrill Imitator | synergy | Dinosaur synergy creature that adds flexibility to the board.
- 1 Territorial Hammerskull | synergy | Dinosaur synergy creature that supports attacking turns.
- 1 Carnage Tyrant | threat | Dinosaur threat that helps close games through combat.
- 1 Etali, Primal Storm | threat | High-impact Dinosaur threat for the deck’s top end.
- 1 Ghalta, Primal Hunger | threat | Major Dinosaur threat for finishing creature-focused games.
- 1 Ghalta and Mavren | threat | Legendary Dinosaur threat that strengthens the top end.
- 1 Regisaur Alpha | threat | Dinosaur threat that supports the deck’s aggressive board plan.
- 1 Quartzwood Crasher | threat | Dinosaur threat for pressuring opponents in combat.
- 1 Zetalpa, Primal Dawn | threat | Large Dinosaur threat for decisive late-game turns.
- 1 Tyrranax Rex | threat | Dinosaur threat that provides a powerful finishing body.
- 1 Pantlaza, Sun-Favored | threat | Dinosaur threat that rewards keeping the deck creature-focused.
- 1 Goring Ceratops | threat | Dinosaur threat that adds pressure to combat turns.
- 1 Ghalta, Stampede Tyrant | threat | High-end Dinosaur threat for overwhelming late-game boards.
- 1 Verdant Sun's Avatar | threat | Dinosaur threat that adds a substantial body to the finish.

</details>

### 5. blink Commander, owned first

Format: Commander. Theme: blink. Pool: owned_first. Shortlist: 168 names.

Commander: Gilraen, Dúnedain Protector.

Grade: typical, score 0.41, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $178.35 the whole deck.

**Summary:** This blink deck develops its mana, deploys value creatures and artifacts, then repeatedly leans on its blink-focused pieces while maintaining a steady flow of cards. Angels, Humans, and legendary creatures provide the board pressure needed to finish games, while efficient removal, protection, and a small reset package keep opposing boards manageable. It gives up the explosive speed of the format's fastest shells in exchange for a resilient, creature-centered midrange game.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 2.81 over 63 nonland cards

<details><summary>The deck list</summary>

- 24 Plains | land | Reliable white mana base for the deck.
- 1 Command Tower | land | Flexible colored land for the commander deck.
- 1 City of Brass | land | Flexible colored land for the mana base.
- 1 Exotic Orchard | land | Flexible colored land for the mana base.
- 1 Marsh Flats | land | Fetch land that supports the Plains-heavy mana base.
- 1 Fabled Passage | land | Fetch land that supports the Plains-heavy mana base.
- 1 Ash Barrens | land | Mana-base utility land with basic-land support.
- 1 Minas Tirith | land | White-producing utility land.
- 1 Plaza of Heroes | land | Utility land that fits the legendary commander shell.
- 1 Spire of Industry | land | Flexible mana land alongside the deck's artifacts.
- 1 Unclaimed Territory | land | Colored mana support for the creature base.
- 1 Secluded Courtyard | land | Colored mana support for the creature base.
- 1 Opal Palace | land | Commander-focused utility land.
- 1 Sol Ring | ramp | Efficient artifact mana acceleration.
- 1 Arcane Signet | ramp | Reliable mana acceleration for the deck.
- 1 Fellwar Stone | ramp | Low-cost mana acceleration.
- 1 Wayfarer's Bauble | ramp | Early mana development tool.
- 1 Sword of the Animist | ramp | Equipment-based mana development.
- 1 Relic of Legends | ramp | Artifact mana support for the creature deck.
- 1 Thought Vessel | ramp | Artifact mana acceleration.
- 1 Chromatic Lantern | ramp | Mana acceleration and fixing support.
- 1 White Auracite | ramp | Artifact mana acceleration.
- 1 Bender's Waterskin | ramp | Additional artifact-based mana development.
- 1 Adventurer's Airship | draw | Card-advantage piece for keeping resources flowing.
- 1 Buster Sword | draw | Equipment-based card-advantage option.
- 1 Crown of Gondor | draw | Legendary equipment that supports card advantage.
- 1 Diary of Dreams | draw | Artifact card-advantage source.
- 1 Energybending | draw | Flexible card-advantage spell.
- 1 Idol of Oblivion | draw | Low-cost artifact card-advantage piece.
- 1 Instant Ramen | draw | Artifact card-advantage support.
- 1 Mask of Memory | draw | Equipment that supplies card advantage.
- 1 Mirror of Galadriel | draw | Legendary artifact card-advantage source.
- 1 Tome of Legends | draw | Commander-oriented card-advantage artifact.
- 1 Wizard's Rockets | draw | Artifact card-advantage option.
- 1 Banishing Light | removal | Broad permanent removal.
- 1 Crib Swap | removal | Creature removal option.
- 1 Destroy Evil | removal | Flexible targeted removal.
- 1 Dispatch | removal | Efficient targeted removal.
- 1 Generous Gift | removal | Flexible answer to opposing permanents.
- 1 Get Lost | removal | Versatile targeted removal.
- 1 Make Your Move | removal | Instant-speed removal support.
- 1 Swords to Plowshares | removal | Efficient creature removal.
- 1 Stroke of Midnight | removal | Flexible instant-speed removal.
- 1 Flickerwisp | synergy | Core creature for the blink-focused plan.
- 1 Fiend Hunter | synergy | Creature-based blink payoff.
- 1 Inspiring Overseer | synergy | Creature value piece for repeated blink use.
- 1 Wall of Omens | synergy | Low-cost creature value piece for blink lines.
- 1 Palace Jailer | synergy | Creature-based value piece for the blink plan.
- 1 Personify | synergy | Direct support for the deck's blink theme.
- 1 Slip On the Ring | synergy | Instant-speed blink support.
- 1 Lembas | synergy | Artifact value piece that works well with blinking.
- 1 Gift of Immortality | synergy | Recursion support for important creature pieces.
- 1 Together Forever | synergy | Long-game support for the creature-based plan.
- 1 Angel of Serenity | threat | Large Angel threat for closing games.
- 1 Angel of Sanctions | threat | Angel threat that adds pressure to the board.
- 1 Angel of Condemnation | threat | Angel threat that fits the blink-oriented creature plan.
- 1 Exemplar of Light | threat | Angel creature that contributes battlefield pressure.
- 1 Giada, Font of Hope | threat | Legendary Angel threat for the creature plan.
- 1 Champions of Minas Tirith | threat | Human creature threat for developing the board.
- 1 Faramir, Field Commander | threat | Legendary Human threat that supports board presence.
- 1 Frontline Medic | threat | Human creature that contributes combat pressure.
- 1 Puresteel Paladin | threat | Creature threat that complements the equipment package.
- 1 Kataki, War's Wage | threat | Legendary creature threat with disruptive board presence.
- 1 Westfold Rider | threat | Human Knight that adds to the creature offense.
- 1 Zack Fair | threat | Legendary Human Soldier threat for the board.
- 1 Bastion Protector | interaction | Protection piece for the commander and key creatures.
- 1 Boromir, Warden of the Tower | interaction | Creature-based protective interaction.
- 1 Clever Concealment | interaction | Protective interaction for preserving the board.
- 1 Champion's Helm | interaction | Equipment protection for the commander or a threat.
- 1 Lightning Greaves | interaction | Efficient equipment protection.
- 1 Reprieve | interaction | Flexible instant-speed interaction.
- 1 Swiftfoot Boots | interaction | Additional equipment-based protection.
- 1 Unbreakable Formation | interaction | Protective interaction for the creature board.
- 1 Austere Command | wipe | Flexible reset for problematic boards.
- 1 Fumigate | wipe | Creature-board reset option.
- 1 Vanquish the Horde | wipe | Efficient creature-board reset.

</details>

### 6. Modern tempo, tournament

Format: Modern. Theme: tempo. Pool: any_card. Shortlist: 167 names.

Grade: typical, score 0.50, model 20260903T210658Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $438.08 to buy, $438.08 the whole deck.

**Summary:** This blue-red tempo deck establishes pressure with its creature threats, then backs them with draw selection, efficient removal, and a compact interaction suite. Temporal Trespass serves as the deck’s synergy package while the mana base keeps both colors available without relying on lands that enter tapped. It wins by sustaining pressure while disrupting the opponent’s key plays, giving up broader answers and larger standalone finishers for a focused tempo plan.

The quality model grades this deck typical against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade. The color sources cover the pips, and that raises the grade.

- [INFO] `curve_summary`: average mana value 2.56 over 36 nonland cards

<details><summary>The deck list</summary>

- 6 Island | land | Basic blue source for the tempo spell base.
- 4 Mountain | land | Basic red source for removal spells.
- 4 Steam Vents | land | Blue-red dual land for consistent access to both colors.
- 4 Spirebluff Canal | land | Blue-red land supporting the two-color mana base.
- 4 Riverglide Pathway // Lavaglide Pathway | land | Flexible blue or red source for the tempo plan.
- 2 Shivan Reef | land | Additional blue-red source to complete the mana base.
- 4 Ragavan, Nimble Pilferer | threat | Creature threat that applies pressure.
- 4 Ledger Shredder | threat | Creature threat that supports a spell-heavy tempo shell.
- 4 Faerie Mastermind | threat | Creature threat that helps maintain pressure.
- 4 Consider | draw | Efficient draw selection for finding threats and answers.
- 2 Preordain | draw | Draw selection that smooths the deck's game plan.
- 2 Counterspell | interaction | Reliable interaction for protecting pressure or stopping opposing plays.
- 2 Force of Negation | interaction | Interaction that lets the deck contest key opposing plays.
- 2 Spell Pierce | interaction | Cheap interaction for maintaining tempo.
- 4 Lightning Bolt | removal | Efficient removal for clearing blockers and opposing threats.
- 2 Into the Flood Maw | removal | Removal that handles problematic opposing permanents.
- 2 Abrade | removal | Flexible removal for opposing targets.
- 4 Temporal Trespass | synergy | Synergy payoff for the deck's spell-focused plan.

</details>

### 7. Modern burn, casual

Format: Modern. Theme: burn. Pool: any_card. Shortlist: 290 names.

Grade: baseline, score 0.39, model 20260903T210658Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $28.26 to buy, $28.26 the whole deck.

**Summary:** This mono-red burn deck opens with cheap removal and spell-focused synergy pieces, then keeps pressure on the table with a concentrated red threat package. It aims to win by combining repeated burn pressure with creature attacks, using its draw cards to keep the action flowing. The deck gives up broad interaction and defensive flexibility in favor of a direct, proactive plan.

The quality model grades this deck at the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade.

- [INFO] `curve_summary`: average mana value 2.61 over 36 nonland cards

<details><summary>The deck list</summary>

- 24 Mountain | land | Provides a consistent mono-red mana base.
- 2 Chandra, Dressed to Kill | ramp | Fills the requested ramp role while supporting the red game plan.
- 4 Ancestral Anger | draw | Supplies efficient card access for a spell-heavy burn shell.
- 2 Risk Factor | draw | Adds further card access and keeps pressure on the opponent.
- 4 Lightning Bolt | removal | Primary low-cost burn removal.
- 2 Lava Dart | removal | Adds more cheap removal to support the burn plan.
- 4 Eidolon of the Great Revel | synergy | A synergistic red permanent for a low-curve burn strategy.
- 4 Thermo-Alchemist | synergy | Supports the deck's spell-focused burn synergies.
- 4 Barret Wallace | threat | Provides a red threat to maintain board pressure.
- 4 Hazoret the Fervent | threat | A resilient red threat for closing games.
- 2 Ashcloud Phoenix | threat | Adds more threatening creatures for the midgame.
- 4 Keral Keep Disciples | threat | Rounds out the creature pressure package.

</details>

### 8. Modern lifegain, FNM

Format: Modern. Theme: lifegain. Pool: any_card. Shortlist: 301 names.

Grade: bad, score 0.08, model 20260903T210658Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $613.52 to buy, $613.52 the whole deck.

**Summary:** This white-black lifegain deck establishes its board with Soul Warden and Ajani's Pridemate, then develops pressure through lifegain-oriented creatures and larger threats such as Sheoldred, the Apocalypse and Archangel of Thune. Dawn of Hope and Inspiring Overseer help keep cards flowing, while Path to Exile and Solitude clear opposing threats. It wins through sustained creature pressure, giving up some early speed and relying on its creatures remaining in play to make the lifegain plan matter.

The quality model grades this deck below the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The deck runs its spells as playsets, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade.

- [INFO] `curve_summary`: average mana value 2.89 over 36 nonland cards

<details><summary>The deck list</summary>

- 8 Plains | land | Forms the white portion of the mana base.
- 6 Swamp | land | Forms the black portion of the mana base.
- 4 Godless Shrine | land | Supports the two-color mana base.
- 4 Caves of Koilos | land | Supports the two-color mana base.
- 2 Shattered Sanctum | land | Rounds out the two-color mana base.
- 2 Altar of the Pantheon | ramp | Provides the deck's ramp package.
- 4 Dawn of Hope | draw | Provides repeatable card draw for the lifegain plan.
- 2 Inspiring Overseer | draw | Adds card draw while contributing to the creature plan.
- 4 Path to Exile | removal | Forms the primary removal package.
- 2 Solitude | removal | Adds further removal to answer opposing threats.
- 4 Soul Warden | synergy | Supplies a core lifegain synergy piece.
- 4 Ajani's Pridemate | synergy | Rewards the deck's lifegain-focused plan.
- 4 Attended Healer | threat | Provides a lifegain-oriented creature threat.
- 4 Twinblade Paladin | threat | Adds a creature threat suited to the lifegain plan.
- 4 Sheoldred, the Apocalypse | threat | Provides a powerful standalone threat.
- 2 Archangel of Thune | threat | Supplies a high-impact threat for the lifegain strategy.

</details>

### 9. Standard midrange, FNM

Format: Standard. Theme: midrange. Pool: any_card. Shortlist: 164 names.

Grade: bad, score 0.08, model 20260903T210658Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $118.31 to buy, $118.31 the whole deck.

**Summary:** This black-green midrange deck develops its mana, builds around creature-focused permanents, and keeps cards flowing into longer games. It aims to win through sustained pressure from its creature threats while using targeted removal to clear the way. The tradeoff is a focused main deck with little room for broad reset effects or specialized interaction.

The quality model grades this deck below the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. Many of the cards appear in no top list, and that raises the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 2.89 over 36 nonland cards

<details><summary>The deck list</summary>

- 10 Forest | land | Basic land in the black-green mana base.
- 8 Swamp | land | Basic land in the black-green mana base.
- 4 Overgrown Tomb | land | Black-green land for the mana base.
- 2 Blooming Marsh | land | Black-green land for the mana base.
- 2 Llanowar Elves | ramp | Early ramp creature for the midrange plan.
- 2 Darkstar Augur | draw | Creature draw option for maintaining resources.
- 2 Phyrexian Arena | draw | Dedicated draw permanent for longer games.
- 2 Corrupted Conviction | draw | Low-cost draw spell for the deck's resource plan.
- 2 Bitter Triumph | removal | Efficient removal for opposing threats.
- 2 Hero's Downfall | removal | Reliable removal for the midrange shell.
- 2 Maelstrom Pulse | removal | Flexible removal for problematic permanents.
- 4 Insidious Roots | synergy | Enchantment synergy piece for the creature-focused plan.
- 2 Enduring Vitality | synergy | Creature-enchantment synergy piece that supports the deck's permanent base.
- 2 Garruk's Uprising | synergy | Enchantment synergy piece alongside the deck's larger creatures.
- 4 Goldvein Hydra | threat | Creature threat that gives the deck a proactive board presence.
- 4 Chomping Changeling | threat | Creature threat for applying pressure in midrange games.
- 3 Rottenmouth Viper | threat | Creature threat that strengthens the deck's black creature suite.
- 3 Vein Ripper | threat | Creature threat for closing games after the deck establishes control.

</details>

### 10. Standard aggro, tournament

Format: Standard. Theme: aggro. Pool: any_card. Shortlist: 294 names.

Grade: bad, score 0.08, model 20260903T210658Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $51.98 to buy, $51.98 the whole deck.

**Summary:** This red-white aggro deck aims to establish pressure with efficient threats, then use its synergy cards to make that board more punishing. Removal and interaction keep opposing defenses from stabilizing, while draw cards help sustain the attack when the first wave is answered. It gives up slower, resource-heavy options in favor of a direct plan that wants to stay ahead on the battlefield.

The quality model grades this deck below the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that raises the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 2.94 over 36 nonland cards

<details><summary>The deck list</summary>

- 12 Mountain | land | Part of the red portion of the mana base.
- 12 Plains | land | Part of the white portion of the mana base.
- 4 Fugitive Codebreaker | draw | Draw support that helps the deck maintain resources while applying pressure.
- 2 Reckless Lackey | draw | Additional draw support for an aggressive game plan.
- 4 Boros Charm | interaction | Flexible interaction to support the board and disrupt opposing plans.
- 2 Sheltered by Ghosts | interaction | Interaction that helps the deck answer opposing plays.
- 4 Emeritus of Truce // Swords to Plowshares | removal | Reliable removal for clearing opposing threats.
- 4 Case of the Gateway Express | removal | Additional removal to keep blockers and threats off the board.
- 4 Slickshot Show-Off | synergy | Synergy piece that supports the deck's aggressive pressure.
- 4 Bedhead Beastie | threat | Early threat that advances the aggressive battlefield plan.
- 4 Frilled Sparkshooter | threat | Threat that helps keep damage pressure on the opponent.
- 4 Dragonback Lancer | threat | Threat that contributes to the deck's combat-focused finish.

</details>

### 11. Commander with a locked card

Format: Commander. Theme: sacrifice. Pool: any_card. Shortlist: 296 names.

Commander: Karlov of the Ghost Council.

Grade: bad, score 0.07, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $1495.36 to buy, $1495.36 the whole deck.

**Summary:** Karlov leads a black-white sacrifice deck that develops a creature board, converts expendable creatures into mana, cards, and removal, and uses death-focused synergies to maintain pressure. It wins by combining its threats with the sacrifice engine and creature-death payoffs, while its interaction and wipes keep opposing boards manageable. The deck gives up some individual-card independence in exchange for a plan that is strongest when its creatures and sacrifice pieces work together.

The quality model grades this deck below the precon baseline against the top lists of the format: more opening hands hold two to four lands than the norm, and that raises the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. Few lands enter tapped, and that lowers the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 3.08 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Command Tower | land | Provides reliable black or white mana.
- 1 Caves of Koilos | land | Provides flexible black or white mana.
- 1 City of Brass | land | Provides flexible colored mana.
- 1 Godless Shrine | land | Provides both deck colors.
- 1 Fetid Heath | land | Helps filter mana into the deck's colors.
- 1 High Market | land | Provides a repeatable land-based sacrifice outlet.
- 1 Isolated Chapel | land | Provides black or white mana.
- 1 Mana Confluence | land | Provides flexible colored mana.
- 1 Scrubland | land | Provides both deck colors.
- 1 Shattered Sanctum | land | Provides black or white mana.
- 1 Vault of Champions | land | Provides black or white mana.
- 1 Diamond Valley | land | Adds another land-based sacrifice outlet.
- 1 Phyrexian Tower | land | Converts sacrificed creatures into mana.
- 1 Westvale Abbey // Ormendahl, Profane Prince | land | Provides a sacrifice-focused utility land.
- 12 Plains | land | Supplies dependable white mana.
- 10 Swamp | land | Supplies dependable black mana.
- 1 Sol Ring | ramp | Provides efficient early mana and is a required inclusion.
- 1 Ashnod's Altar | ramp | Turns sacrificed creatures into mana.
- 1 Ashnod, Flesh Mechanist | ramp | Supports the sacrifice plan while providing mana.
- 1 Crowded Crypt | ramp | Builds toward additional mana from creature deaths.
- 1 Culling the Weak | ramp | Converts a creature into a burst of mana.
- 1 Fain, the Broker | ramp | Provides a sacrifice-compatible mana outlet.
- 1 Phyrexian Altar | ramp | Converts sacrificed creatures into colored mana.
- 1 Pitiless Plunderer | ramp | Rewards creature deaths with mana resources.
- 1 Priest of Forgotten Gods | ramp | Uses creatures to generate mana while advancing the plan.
- 1 Warren Soultrader | ramp | Provides mana through the deck's creature resources.
- 1 Baron Bertram Graywater | draw | Provides card advantage within the creature plan.
- 1 Bushmeat Poacher | draw | Turns creatures into cards.
- 1 Corrupted Conviction | draw | Trades a creature for efficient card advantage.
- 1 Disciple of Bolas | draw | Converts a sacrificed creature into cards.
- 1 Ecstatic Awakener // Awoken Demon | draw | Provides card advantage while fitting sacrifice play.
- 1 Jerren, Corrupted Bishop // Ormendahl, the Corrupter | draw | Adds card advantage to the creature-heavy plan.
- 1 Lord Skitter's Butcher | draw | Provides creature-based card advantage.
- 1 Relic Vial | draw | Supplies repeatable card advantage for the creature suite.
- 1 Shadowheart, Dark Justiciar | draw | Turns a creature into a substantial card draw effect.
- 1 Smothering Abomination | draw | Rewards creature sacrifices with cards.
- 1 Vampiric Rites | draw | Provides a repeatable sacrifice outlet that draws cards.
- 1 Cartel Aristocrat | interaction | Provides a sacrifice outlet and protects itself.
- 1 Dark Privilege | interaction | Uses sacrifices to protect a key creature.
- 1 Fanatical Devotion | interaction | Turns expendable creatures into protection.
- 1 Flare of Fortitude | interaction | Protects the board at a critical moment.
- 1 Gift of Doom | interaction | Protects an important permanent while using sacrifice resources.
- 1 Promise of Tomorrow | interaction | Helps preserve creatures through opposing pressure.
- 1 Spirit Bonds | interaction | Provides creature protection and supports the board plan.
- 1 Sunstone | interaction | Offers a defensive tool when under attack.
- 1 Attrition | removal | Turns sacrificed creatures into creature removal.
- 1 Ayli, Eternal Pilgrim | removal | Provides a sacrifice outlet with removal utility.
- 1 Blasting Station | removal | Converts sacrificed creatures into direct removal pressure.
- 1 Bone Shards | removal | Provides cheap removal that can use a creature as a cost.
- 1 Bone Splinters | removal | Uses an expendable creature to answer a target.
- 1 Dictate of Erebos | removal | Makes opposing boards pay for the deck's creature deaths.
- 1 Eaten Alive | removal | Provides efficient removal with a sacrifice option.
- 1 Grave Pact | removal | Turns each sacrificed creature into broad opposing creature pressure.
- 1 Yawgmoth, Thran Physician | removal | Offers repeatable creature-based removal utility.
- 1 Austere Command | wipe | Provides a flexible reset when the board gets out of hand.
- 1 The Meathook Massacre | wipe | Provides a board reset that fits creature-death play.
- 1 Toxic Deluge | wipe | Provides an efficient creature-board reset.
- 1 Altar of Dementia | synergy | Provides a sacrifice outlet and an alternate pressure angle.
- 1 Bastion of Remembrance | synergy | Rewards the deck for creatures dying.
- 1 Carrion Feeder | synergy | Provides a cheap repeatable sacrifice outlet.
- 1 Elas il-Kor, Sadistic Pilgrim | synergy | Rewards the deck's creature-death plan.
- 1 Hidden Stockpile | synergy | Builds value around sacrificing permanents.
- 1 Martyr's Cause | synergy | Turns expendable creatures into damage prevention.
- 1 Viscera Seer | synergy | Provides a low-cost repeatable sacrifice outlet.
- 1 Witch's Oven | synergy | Provides a repeatable artifact sacrifice outlet.
- 1 Woe Strider | synergy | Brings a sacrifice outlet attached to a creature.
- 1 Zulaport Cutthroat | synergy | Rewards each creature death with additional pressure.
- 1 Abhorrent Overlord | threat | Provides a large closing threat for the creature plan.
- 1 Basri's Lieutenant | threat | Adds a board-building creature threat.
- 1 Felisa, Fang of Silverquill | threat | Turns the deck's creature losses into continued board pressure.
- 1 Ghoulcaller Gisa | threat | Builds a threatening board from expendable creatures.
- 1 Liesa, Forgotten Archangel | threat | Provides a resilient top-end creature threat.
- 1 Liliana, Heretical Healer // Liliana, Defiant Necromancer | threat | Rewards creature deaths with an enduring threat.
- 1 Mondrak, Glory Dominus | threat | Amplifies the deck's token-oriented board development.
- 1 Ratadrabik of Urborg | threat | Keeps the legendary creature suite producing board pressure.
- 1 Razaketh, the Foulblooded | threat | Provides a powerful sacrifice-centered finisher.
- 1 Requiem Angel | threat | Turns creature deaths into a continuing aerial board presence.
- 1 Sidisi, Undead Vizier | threat | Provides a substantial creature threat that fits sacrifice play.
- 1 Vindictive Vampire | threat | Adds creature-death pressure that helps close games.

</details>

### 12. Commander on a budget

Format: Commander. Theme: tokens. Pool: any_card. Shortlist: 296 names.

Commander: Adeline, Resplendent Cathar.

Grade: bad, score 0.02, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $83.03 to buy, $83.03 the whole deck.

**Summary:** This Adeline-led white token deck develops a broad creature board, supports that board with token-focused permanents, and keeps cards flowing while applying steady combat pressure. It wins by turning a growing army sideways, with larger token payoffs and combat threats providing additional closing power. The deck gives up speed and some flexibility for a budget-minded, board-centric plan that is most effective when its creature presence remains established.

The quality model grades this deck below the precon baseline against the top lists of the format: the deck makes less mana on turn four than the norm, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 3.53 over 62 nonland cards

<details><summary>The deck list</summary>

- 28 Plains | land | Provides the primary white mana base.
- 1 Ancient Den | land | Adds a white-producing land slot.
- 1 Castle Ardenvale | land | Adds a utility land slot to the mana base.
- 1 Eiganjo, Seat of the Empire | land | Adds a white land with utility.
- 1 Exotic Orchard | land | Adds flexible mana support.
- 1 Fountainport | land | Adds a utility land slot.
- 1 Mirrex | land | Adds a utility land slot.
- 1 Plaza of Heroes | land | Adds flexible mana support.
- 1 Secluded Courtyard | land | Supports the deck's creature-heavy plan.
- 1 Windbrisk Heights | land | Adds a low-cost utility land.
- 1 Bygone Bishop | draw | Provides the listed draw role.
- 1 Caretaker's Talent | draw | Provides a strong token-oriented draw slot.
- 1 Dawn of Hope | draw | Provides the listed draw role.
- 1 Faramir, Field Commander | draw | Provides the listed draw role.
- 1 Idol of Oblivion | draw | Provides a token-focused draw slot.
- 1 Platoon Dispenser | draw | Provides the listed draw role.
- 1 Sanctuary Warden | draw | Provides the listed draw role.
- 1 Search the Premises | draw | Provides the listed draw role.
- 1 Staff of the Storyteller | draw | Provides the listed draw role.
- 1 Wedding Announcement // Wedding Festivity | draw | Provides a token-oriented draw slot.
- 1 Wojek Investigator | draw | Provides the listed draw role.
- 1 Charisma Bobblehead | ramp | Provides the listed ramp role.
- 1 Coin of Mastery | ramp | Provides the listed ramp role.
- 1 Collector's Vault | ramp | Provides the listed ramp role.
- 1 Currency Converter | ramp | Provides the listed ramp role.
- 1 Druidic Satchel | ramp | Provides the listed ramp role.
- 1 Goldvein Pick | ramp | Provides the listed ramp role.
- 1 Karn, Living Legacy | ramp | Provides the listed ramp role.
- 1 Keeper of the Accord | ramp | Provides the listed ramp role.
- 1 Monologue Tax | ramp | Provides the listed ramp role.
- 1 Noble's Purse | ramp | Provides the listed ramp role.
- 1 Aerial Assault | removal | Provides a low-cost removal slot.
- 1 Banishing Slash | removal | Provides a low-cost removal slot.
- 1 Generous Gift | removal | Provides flexible removal.
- 1 Kellan's Lightblades | removal | Provides a low-cost removal slot.
- 1 Righteous Confluence | removal | Provides flexible removal.
- 1 Skyclave Apparition | removal | Provides creature-based removal.
- 1 Stroke of Midnight | removal | Provides a low-cost removal slot.
- 1 Blessed Sanctuary | interaction | Provides the listed interaction role.
- 1 Lena, Selfless Champion | interaction | Provides creature-based interaction.
- 1 Rootborn Defenses | interaction | Provides token-friendly interaction.
- 1 Spirit Bonds | interaction | Provides the listed interaction role.
- 1 Squad Commander | interaction | Provides creature-based interaction.
- 1 Teyo, the Shieldmage | interaction | Provides the listed interaction role.
- 1 Elspeth, Sun's Champion | wipe | Provides a wipe that also suits the token plan.
- 1 Hour of Reckoning | wipe | Provides a token-oriented wipe slot.
- 1 Martial Coup | wipe | Provides a wipe for the deck.
- 1 Archon of Sun's Grace | threat | Provides a token-oriented threat.
- 1 Attended Healer | threat | Provides a token-oriented threat.
- 1 Basri's Lieutenant | threat | Provides a creature threat.
- 1 Cemetery Protector | threat | Provides a creature threat.
- 1 Defiler of Faith | threat | Provides a creature threat.
- 1 Emeria Angel | threat | Provides a token-oriented threat.
- 1 Gideon, Ally of Zendikar | threat | Provides a token-oriented threat.
- 1 God-Eternal Oketra | threat | Provides a creature threat.
- 1 Hero of Bladehold | threat | Provides a combat-focused threat.
- 1 Oketra the True | threat | Provides a creature threat.
- 1 Requiem Angel | threat | Provides a token-oriented threat.
- 1 Threefold Thunderhulk | threat | Provides a token-oriented artifact threat.
- 1 Anointer Priest | synergy | Supports the token-focused plan.
- 1 Cathar's Call | synergy | Supports the token-focused plan.
- 1 Clarion Spirit | synergy | Supports the token-focused plan.
- 1 Divine Visitation | synergy | Provides a high-impact token synergy piece.
- 1 Felidar Retreat | synergy | Supports the token-focused plan.
- 1 Horn of Gondor | synergy | Provides a token-focused synergy piece.
- 1 Inspiring Leader | synergy | Supports the token-focused plan.
- 1 Intangible Virtue | synergy | Supports the token-focused plan.
- 1 Oketra's Monument | synergy | Provides a creature-focused synergy piece.
- 1 Rosie Cotton of South Lane | synergy | Supports the token-focused plan.
- 1 Siege Veteran | synergy | Supports the token-focused plan.
- 1 Sigil of the Empty Throne | synergy | Provides a token-focused synergy piece.
- 1 Skrelv's Hive | synergy | Provides a token-focused synergy piece.

</details>

### 13. owned first, and the commander is not owned

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 229 names.

Commander: Karlov of the Ghost Council.

Grade: bad, score 0.32, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $152.96 the whole deck.

**Summary:** This deck develops mana and a steady stream of cards, then builds around Karlov with lifegain-themed creatures, Angels, and equipment. It aims to establish a threatening board while holding efficient answers and protective spells for key pieces, using sweepers to reset opponents when the battlefield gets crowded. It gives up some raw speed for a fuller land base, creature-based pressure, and a more resilient midrange game.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- [WARN] `not_owned`: Karlov of the Ghost Council: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.03 over 63 nonland cards

<details><summary>The deck list</summary>

- 13 Plains | land | Basic white mana for the lifegain-focused core.
- 12 Swamp | land | Basic black mana for the deck’s spells and threats.
- 1 Command Tower | land | Reliable multicolor mana for Karlov’s colors.
- 1 Exotic Orchard | land | Flexible color fixing.
- 1 City of Brass | land | Flexible access to both deck colors.
- 1 Grand Coliseum | land | Additional multicolor fixing.
- 1 Castle Locthwain | land | Black-producing utility land.
- 1 Minas Tirith | land | White-producing utility land.
- 1 Takenuma, Abandoned Mire | land | Black-producing utility land.
- 1 Plaza of Heroes | land | Utility mana for a commander-centered deck.
- 1 Spire of Industry | land | Color fixing alongside the artifact package.
- 1 Study Hall | land | Utility land slot with mana production.
- 1 Windbrisk Heights | land | White-producing utility land.
- 1 Sol Ring | ramp | Efficient early mana acceleration.
- 1 Arcane Signet | ramp | Reliable color fixing and acceleration.
- 1 Fellwar Stone | ramp | Low-cost mana acceleration.
- 1 Wayfarer's Bauble | ramp | Early ramp that supports the land base.
- 1 Thought Vessel | ramp | Mana acceleration in an artifact slot.
- 1 Commander's Sphere | ramp | Color fixing with later utility.
- 1 Sword of the Animist | ramp | Equipment-based ramp for an attacking creature deck.
- 1 Giada, Font of Hope | ramp | Ramp that also fits the Angel threat package.
- 1 Lotho, Corrupt Shirriff | ramp | A themed creature ramp slot.
- 1 White Lotus Tile | ramp | Additional artifact-based mana support.
- 1 Night's Whisper | draw | Efficient card flow.
- 1 Wall of Omens | draw | Early defensive card draw.
- 1 Idol of Oblivion | draw | Repeatable draw support.
- 1 Lembas | draw | Low-cost card flow in an artifact slot.
- 1 Mask of Memory | draw | Equipment-based card advantage.
- 1 Tome of Legends | draw | Card flow for a commander-focused plan.
- 1 Skullclamp | draw | Efficient equipment-based draw support.
- 1 Puresteel Paladin | draw | Card advantage that supports the equipment package.
- 1 Inspiring Overseer | draw | Creature-based draw that fits the Angel suite.
- 1 Exemplar of Light | draw | An Angel draw slot for the creature plan.
- 1 Call of the Ring | draw | Ongoing card flow.
- 1 Swords to Plowshares | removal | Efficient targeted removal.
- 1 Generous Gift | removal | Flexible permanent removal.
- 1 Get Lost | removal | Versatile targeted removal.
- 1 Stroke of Midnight | removal | Flexible answer to opposing permanents.
- 1 Dispatch | removal | Low-cost targeted removal.
- 1 Infernal Grasp | removal | Direct creature removal.
- 1 Bitter Triumph | removal | Flexible targeted removal.
- 1 Fatal Push | removal | Efficient early interaction with creatures.
- 1 Crib Swap | removal | Creature removal in an instant-speed slot.
- 1 Lightning Greaves | interaction | Protects a key creature or the commander.
- 1 Swiftfoot Boots | interaction | Protection for important creatures.
- 1 Darksteel Plate | interaction | A durable protection piece for a threat.
- 1 Champion's Helm | interaction | Commander-focused protection.
- 1 Clever Concealment | interaction | Protective interaction for the board.
- 1 Unbreakable Formation | interaction | Board protection during combat-focused turns.
- 1 Reprieve | interaction | Flexible stack interaction.
- 1 Take Up the Shield | interaction | Protection for a key creature.
- 1 Angel of Vitality | synergy | A lifegain-themed creature that supports Karlov’s plan.
- 1 Compassionate Healer | synergy | A lifegain-focused synergy creature.
- 1 Rosie Cotton of South Lane | synergy | A creature that fits the lifegain synergy package.
- 1 Light of Promise | synergy | A lifegain synergy piece for building a large threat.
- 1 Aerith Gainsborough | synergy | A themed synergy creature for the life-focused strategy.
- 1 Night Nurse, Healer of Heroes | synergy | A healer-themed lifegain synergy slot.
- 1 Kor Firewalker | synergy | A creature-based lifegain synergy piece.
- 1 Prideful Feastling | synergy | Supports the deck’s lifegain theme.
- 1 Second Breakfast | synergy | A lifegain-themed support spell.
- 1 Aettir and Priwen | synergy | Equipment synergy for the creature-heavy plan.
- 1 Angel of Invention | threat | An Angel threat that advances the creature plan.
- 1 Bill the Pony | threat | A low-curve creature threat.
- 1 Dawnhand Eulogist | threat | A creature threat that pressures opponents.
- 1 Gollum, Silent Slinker // Meager Meal | threat | A black creature threat with an attached adventure.
- 1 Lyra Dawnbringer | threat | A top-end Angel threat.
- 1 Minwu, White Mage | threat | A creature threat that fits the white core.
- 1 Rabaroo Troop | threat | A creature threat for developing the board.
- 1 Reaping Willow | threat | A resilient creature threat slot.
- 1 Rooftop Percher | threat | An additional creature threat.
- 1 Shattered Angel | threat | An Angel threat aligned with the lifegain plan.
- 1 Sneering Shadewriter | threat | A black creature threat.
- 1 Victory's Herald | threat | A high-impact Angel threat.
- 1 Austere Command | wipe | Flexible board reset.
- 1 Fumigate | wipe | Board wipe that suits the lifegain theme.
- 1 Dusk // Dawn | wipe | A wipe with later creature value.

</details>

### 14. the user delegates the commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 244 names.

Commander: Atarka, World Render.

Grade: bad, score 0.03, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $579.20 to buy, $579.20 the whole deck.

**Summary:** This is a red-green Dragon deck that develops its mana early, uses Dragon-focused support to deploy a steady stream of flying threats, and pressures the table through combat. Atarka, World Render leads the attack while the deck backs its creatures with removal, protective interaction, and several reset buttons. The tradeoff is that the plan remains creature-forward, so it is most effective when it can keep Dragons on the table and turn them sideways.

The quality model grades this deck below the precon baseline against the top lists of the format: more opening hands hold two to four lands than the norm, and that raises the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. The mana base serves one color far better than another, and that lowers the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 3.03 over 63 nonland cards

<details><summary>The deck list</summary>

- 18 Mountain | land | Basic red mana for the mana base.
- 10 Forest | land | Basic green mana for the mana base.
- 1 Taiga | land | Red-green land for the mana base.
- 1 Stomping Ground | land | Red-green land for the mana base.
- 1 Spire Garden | land | Red-green land for the mana base.
- 1 Karplusan Forest | land | Red-green land for the mana base.
- 1 Rockfall Vale | land | Red-green land for the mana base.
- 1 Rootbound Crag | land | Red-green land for the mana base.
- 1 Commercial District | land | Red-green land for the mana base.
- 1 Command Tower | land | Reliable color-fixing land.
- 1 Faithless Looting | draw | Low-cost draw option.
- 1 Demand Answers | draw | Efficient draw option.
- 1 Thrill of Possibility | draw | Low-cost draw option.
- 1 Sensei's Divining Top | draw | Low-cost draw support.
- 1 Skullclamp | draw | Low-cost draw support.
- 1 Sylvan Library | draw | Early draw support.
- 1 Beast Whisperer | draw | Creature-based draw support.
- 1 Elemental Bond | draw | Draw support for the creature plan.
- 1 Garruk's Uprising | draw | Draw support for the creature plan.
- 1 Dragonborn Champion | draw | Dragon-based draw support.
- 1 Harmonize | draw | Straightforward draw refill.
- 1 Heroic Intervention | interaction | Protective interaction for the board.
- 1 Lightning Greaves | interaction | Equipment-based protection.
- 1 Swiftfoot Boots | interaction | Equipment-based protection.
- 1 Snakeskin Veil | interaction | Low-cost protective interaction.
- 1 Tamiyo's Safekeeping | interaction | Low-cost protective interaction.
- 1 Veil of Summer | interaction | Low-cost interaction option.
- 1 Tibalt's Trickery | interaction | Stack interaction option.
- 1 Whispersilk Cloak | interaction | Equipment-based interaction support.
- 1 Carnelian Orb of Dragonkind | ramp | Low-cost ramp for the Dragon plan.
- 1 Jade Orb of Dragonkind | ramp | Low-cost ramp for the Dragon plan.
- 1 Orb of Dragonkind | ramp | Low-cost ramp for the Dragon plan.
- 1 Mox Jasper | ramp | Low-cost ramp support.
- 1 Dragon's Hoard | ramp | Dragon-themed ramp support.
- 1 Dragonstorm Globe | ramp | Mana-fixing ramp support.
- 1 Encroaching Dragonstorm | ramp | Dragon-themed ramp support.
- 1 Scaled Nurturer | ramp | Creature-based ramp for Dragons.
- 1 Reckless Barbarian | ramp | Low-cost creature ramp.
- 1 Sarkhan, Fireblood | ramp | Planeswalker ramp support.
- 1 Draconic Roar | removal | Low-cost removal option.
- 1 Dragon's Fire | removal | Low-cost removal option.
- 1 Molten Exhale | removal | Low-cost removal option.
- 1 Piercing Exhale | removal | Low-cost removal option.
- 1 Spit Flame | removal | Dragon-themed removal option.
- 1 Invasion of Tarkir // Defiant Thundermaw | removal | Dragon-themed removal with a Dragon back face.
- 1 Dragon Tempest | removal | Dragon-themed removal support.
- 1 Glorybringer | removal | Dragon-based removal threat.
- 1 Terror of the Peaks | removal | Dragon-based removal support.
- 1 Acolyte of Bahamut | synergy | Dragon-focused synergy piece.
- 1 Dragon Egg | synergy | Low-cost Dragon synergy piece.
- 1 Dragon Hatchling | synergy | Low-cost Dragon synergy piece.
- 1 Dragonkin Berserker | synergy | Dragon-focused synergy piece.
- 1 Dragonlord's Servant | synergy | Dragon-focused synergy support.
- 1 Dragonspeaker Shaman | synergy | Dragon-focused synergy support.
- 1 Firespitter Whelp | synergy | Dragon-focused synergy piece.
- 1 Kargan Dragonrider | synergy | Dragon-focused synergy piece.
- 1 Minion of the Mighty | synergy | Low-cost Dragon synergy piece.
- 1 Sarkhan's Triumph | synergy | Dragon-focused synergy support.
- 1 Manaform Hellkite | threat | Efficient Dragon threat.
- 1 Mirrorwing Dragon | threat | Mid-curve Dragon threat.
- 1 Thunderbreak Regent | threat | Mid-curve Dragon threat.
- 1 Stormbreath Dragon | threat | Dragon threat for combat pressure.
- 1 Dragonhawk, Fate's Tempest | threat | Dragon threat for combat pressure.
- 1 Thrakkus the Butcher | threat | Dragon threat for combat pressure.
- 1 Caldera Pyremaw | threat | Dragon threat for combat pressure.
- 1 Backdraft Hellkite | threat | Dragon threat for combat pressure.
- 1 Blast-Furnace Hellkite | threat | Dragon threat for combat pressure.
- 1 Lathliss, Dragon Queen | threat | Dragon threat for the tribal plan.
- 1 Dragon Broodmother | threat | Dragon threat for the tribal plan.
- 1 Twinflame Tyrant | threat | Dragon threat for combat pressure.
- 1 Breath Weapon | wipe | Low-cost board wipe option.
- 1 Draconic Intervention | wipe | Dragon-themed board wipe option.
- 1 Steel Hellkite | wipe | Dragon-based board wipe option.

</details>

### 15. delegated commander, owned first

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 229 names.

Commander: Astarion, the Decadent.

Grade: bad, score 0.06, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $174.30 the whole deck.

**Summary:** This lifegain deck builds a resilient white-black creature board, using healer-themed cards, Angels, and equipment to turn a growing life total into pressure. It wins by developing durable threats and attacking through a stabilized board, backed by broad removal and sweepers when opponents get ahead. The tradeoff is a measured, board-focused game plan rather than an especially explosive start.

The quality model grades this deck below the precon baseline against the top lists of the format: more opening hands hold two to four lands than the norm, and that raises the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. The deck holds no fast mana, and that lowers the grade.

Summary rules claims (F-26): the format play

- [WARN] `not_owned`: Astarion, the Decadent: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.11 over 62 nonland cards

<details><summary>The deck list</summary>

- 1 Buster Sword | draw | Equipment-based card advantage supports the creature plan.
- 1 Call of the Ring | draw | A black card-advantage piece for the deck's longer game.
- 1 Exemplar of Light | draw | An Angel that provides card advantage alongside the lifegain theme.
- 1 Cirith Ungol Patrol | draw | A black creature that contributes card advantage.
- 1 Idol of Oblivion | draw | A compact artifact source of card advantage.
- 1 Nasty End | draw | Efficient black card advantage.
- 1 Night's Whisper | draw | Low-cost black card advantage.
- 1 Puresteel Paladin | draw | Supports the equipment package while providing card advantage.
- 1 Skullclamp | draw | A cheap equipment card-advantage engine.
- 1 Tome of Legends | draw | Reliable artifact card advantage over several turns.
- 1 Wall of Omens | draw | Early defense that replaces itself.
- 1 Bastion Protector | interaction | Helps keep the commander protected on the battlefield.
- 1 Champion's Helm | interaction | Equipment protection for the commander or a key threat.
- 1 Clever Concealment | interaction | Protects the board from opposing disruption.
- 1 Darksteel Plate | interaction | Durable equipment protection for an important creature.
- 1 Lightning Greaves | interaction | Cheap protection for the commander and threats.
- 1 Swiftfoot Boots | interaction | A second efficient equipment protection piece.
- 16 Plains | land | Core white mana for Angels, lifegain pieces, and answers.
- 14 Swamp | land | Core black mana for the commander, draw, and removal.
- 1 Command Tower | land | Reliable fixing for the two-color deck.
- 1 Exotic Orchard | land | Flexible color fixing.
- 1 City of Brass | land | Provides dependable access to either deck color.
- 1 Castle Locthwain | land | A black source with additional utility.
- 1 Minas Tirith | land | A white source with useful utility.
- 1 Plaza of Heroes | land | Utility land that supports the legendary core.
- 1 Takenuma, Abandoned Mire | land | Black mana with useful late-game utility.
- 1 Sol Ring | ramp | Efficient colorless acceleration.
- 1 Arcane Signet | ramp | Reliable two-color mana fixing and acceleration.
- 1 Bender's Waterskin | ramp | Artifact acceleration for the deck's mana base.
- 1 Chromatic Lantern | ramp | Mana acceleration that also smooths colors.
- 1 Commander's Sphere | ramp | Color fixing that remains useful later.
- 1 Fellwar Stone | ramp | Low-cost artifact acceleration.
- 1 Inherited Envelope | ramp | Artifact mana acceleration.
- 1 Relic of Legends | ramp | Mana acceleration that works with the legendary creatures.
- 1 Thought Vessel | ramp | Colorless acceleration with additional utility.
- 1 Wayfarer's Bauble | ramp | Early land-based mana development.
- 1 Banishing Light | removal | Flexible white answer to a problematic permanent.
- 1 Bitter Triumph | removal | Efficient black answer to opposing threats.
- 1 Crib Swap | removal | Flexible creature removal.
- 1 Fatal Push | removal | Low-cost black removal.
- 1 Generous Gift | removal | Broad instant-speed permanent removal.
- 1 Get Lost | removal | Efficient white removal for key permanents.
- 1 Infernal Grasp | removal | Straightforward black creature removal.
- 1 Aerith Gainsborough | synergy | A lifegain-focused legendary creature for the deck's core theme.
- 1 Angel of Vitality | synergy | An Angel that directly supports the lifegain plan.
- 1 Aettir and Priwen | synergy | Equipment support for the creature-focused strategy.
- 1 Compassionate Healer | synergy | A Cleric that reinforces the lifegain theme.
- 1 Crowd of True Believers | synergy | A white creature that supports the deck's cohesive plan.
- 1 Elixir | synergy | A compact artifact piece for the deck's synergies.
- 1 Kor Firewalker | synergy | A lifegain-oriented creature for the board.
- 1 Light of Promise | synergy | Rewards the deck for building its life total.
- 1 Night Nurse, Healer of Heroes | synergy | A healer-themed legend that supports lifegain.
- 1 Prideful Feastling | synergy | Supports the deck's food and lifegain-adjacent plan.
- 1 Rosie Cotton of South Lane | synergy | A creature that benefits from the deck's token and lifegain support.
- 1 Second Breakfast | synergy | A themed support spell for the lifegain strategy.
- 1 White Mage's Staff | synergy | Equipment support that fits the white lifegain shell.
- 1 Angel of Invention | threat | An Angel threat that advances the creature plan.
- 1 Bill the Pony | threat | A legendary creature threat that fits the white creature base.
- 1 Dawnhand Eulogist | threat | A creature threat for pressuring opponents.
- 1 Gollum, Silent Slinker // Meager Meal | threat | A black legendary threat with useful flexibility.
- 1 Lyra Dawnbringer | threat | A premier Angel threat for the lifegain-focused board.
- 1 Minwu, White Mage | threat | A white legendary threat that fits the deck's theme.
- 1 Rabaroo Troop | threat | A creature threat that broadens the board presence.
- 1 Reaping Willow | threat | A durable creature threat for the midgame.
- 1 Rooftop Percher | threat | A black creature threat for the attacking plan.
- 1 Shattered Angel | threat | An Angel threat that complements the lifegain shell.
- 1 Sneering Shadewriter | threat | A black creature threat that diversifies the curve.
- 1 Victory's Herald | threat | An Angel finisher for a developed creature board.
- 1 Austere Command | wipe | Flexible sweeper for resetting difficult boards.
- 1 Fumigate | wipe | A board reset that aligns with the lifegain plan.
- 1 Vanquish the Horde | wipe | An efficient mass-creature answer.

</details>

### 16. a tight budget, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 229 names.

Commander: Denethor, Ruling Steward.

Grade: bad, score 0.34, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $22.30 to buy, $45.28 the whole deck.

**Summary:** This is a black-white aristocrats deck built around establishing a creature board, using sacrifice-focused synergy pieces to turn routine exchanges into lasting value, and keeping cards flowing with creature-friendly draw. It wins by steadily compounding its death-and-sacrifice plan while its creature threats pressure opponents after a reset. The deck gives up speed and premium mana for a deliberately lower-powered, library-first build that leans on basics, inexpensive artifacts, and board development.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

Summary rules claims (F-26): the format play

- [WARN] `not_owned`: Disciple of Bolas: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Smothering Abomination: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Aron, Benalia's Ruin: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ayli, Eternal Pilgrim: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bartolomé del Presidio: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bastion of Remembrance: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Blood Artist: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Elas il-Kor, Sadistic Pilgrim: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Falkenrath Noble: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Teysa, Orzhov Scion: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Woe Strider: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Yahenni, Undying Partisan: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Zulaport Cutthroat: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Baron Bertram Graywater: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Embalmed Ascendant: the deck needs 1, the collection has 0
- [WARN] `not_owned`: High-Society Hunter: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Lord Skitter's Butcher: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Lord of the Forsaken: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Old Flitterfang: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ruthless Technomancer: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Shilgengar, Sire of Famine: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Sivriss, Nightmare Speaker: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Skullport Merchant: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vampire Gourmand: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vengeful Bloodwitch: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.03 over 62 nonland cards

<details><summary>The deck list</summary>

- 15 Plains | land | Basic white land for a stable mana base.
- 14 Swamp | land | Basic black land for a stable mana base.
- 1 Command Tower | land | A flexible land slot for the two-color mana base.
- 1 Exotic Orchard | land | A flexible land slot for the two-color mana base.
- 1 Grand Coliseum | land | A flexible land slot for the two-color mana base.
- 1 Path of Ancestry | land | A color-fixing land that supports the creature-heavy plan.
- 1 Secluded Courtyard | land | A creature-focused color-fixing land.
- 1 Spire of Industry | land | A flexible land alongside the deck's artifacts.
- 1 Thriving Moor | land | A black source that can help cover the secondary color.
- 1 Unclaimed Territory | land | A creature-focused color-fixing land.
- 1 Sol Ring | ramp | Efficient early mana for deploying the board.
- 1 Arcane Signet | ramp | Reliable colored mana acceleration.
- 1 Astral Cornucopia | ramp | A flexible artifact mana source.
- 1 Blitzball | ramp | An inexpensive artifact ramp piece.
- 1 Chromatic Lantern | ramp | Artifact acceleration that supports color consistency.
- 1 Commander's Sphere | ramp | Reliable mana acceleration for the midgame.
- 1 Deadly Dispute | ramp | Provides a sacrifice-friendly mana boost.
- 1 Fellwar Stone | ramp | Efficient early artifact acceleration.
- 1 Relic of Legends | ramp | A mana piece for the creature-heavy deck.
- 1 White Auracite | ramp | Low-cost artifact acceleration.
- 1 Beetle-Headed Merchants | draw | A low-cost draw option from the library.
- 1 Cirith Ungol Patrol | draw | A black creature draw option for the deck.
- 1 Disciple of Bolas | draw | A sacrifice-compatible source of cards.
- 1 Grave Venerations | draw | A black draw engine for the longer game.
- 1 Inspiring Overseer | draw | A creature-based draw option.
- 1 Mask of Memory | draw | An equipment-based source of continued cards.
- 1 Massacre Girl, Known Killer | draw | A creature draw engine that fits the black core.
- 1 Nasty End | draw | A cheap black draw spell.
- 1 Painful Truths | draw | Efficient access to additional cards.
- 1 Skullclamp | draw | A key equipment draw outlet for disposable creatures.
- 1 Smothering Abomination | draw | A creature-based draw engine for the sacrifice plan.
- 1 Gift of Immortality | interaction | Protects an important creature in the engine.
- 1 Reprieve | interaction | A flexible reactive spell for protecting tempo.
- 1 Swiftfoot Boots | interaction | Keeps a key creature safer on the table.
- 1 Together Forever | interaction | Helps preserve important creatures through exchanges.
- 1 Ultimate Magic: Holy | interaction | A reactive spell for defending the board plan.
- 1 Unbreakable Formation | interaction | Protects the creature board at an important moment.
- 1 Bitter Triumph | removal | Flexible creature or planeswalker answer.
- 1 Claim the Precious | removal | A black targeted answer.
- 1 Crib Swap | removal | A versatile creature removal spell.
- 1 Deadly Precision | removal | Low-cost targeted removal from the library.
- 1 Fatal Push | removal | Efficient early targeted removal.
- 1 Heartless Act | removal | A flexible black removal spell.
- 1 Infernal Grasp | removal | Clean targeted creature removal.
- 1 Austere Command | wipe | A flexible reset when the board gets away.
- 1 Dusk // Dawn | wipe | A board reset that suits a creature-based deck.
- 1 Fumigate | wipe | A dependable full-board reset.
- 1 Aron, Benalia's Ruin | synergy | A core creature for the aristocrats synergy package.
- 1 Ayli, Eternal Pilgrim | synergy | A central creature for the sacrifice-focused plan.
- 1 Bartolomé del Presidio | synergy | A low-cost creature that supports the aristocrats package.
- 1 Bastion of Remembrance | synergy | A dedicated aristocrats payoff piece.
- 1 Blood Artist | synergy | A premier payoff for the aristocrats plan.
- 1 Elas il-Kor, Sadistic Pilgrim | synergy | An efficient creature payoff for the deck's core plan.
- 1 Falkenrath Noble | synergy | A creature payoff that reinforces the aristocrats strategy.
- 1 Gollum the Abandoned | synergy | A library-owned creature supporting the sacrifice package.
- 1 Gollum, Patient Plotter | synergy | A library-owned creature for the aristocrats core.
- 1 Teysa, Orzhov Scion | synergy | A powerful Orzhov creature for the sacrifice shell.
- 1 Woe Strider | synergy | A creature that directly supports the aristocrats plan.
- 1 Yahenni, Undying Partisan | synergy | A resilient creature for the sacrifice-focused board.
- 1 Zulaport Cutthroat | synergy | An efficient aristocrats payoff creature.
- 1 Baron Bertram Graywater | threat | A black creature threat that complements the deck's theme.
- 1 Embalmed Ascendant | threat | A creature threat for advancing the board.
- 1 High-Society Hunter | threat | A low-cost black creature threat.
- 1 Lord Skitter's Butcher | threat | A creature threat that suits the black creature core.
- 1 Lord of the Forsaken | threat | A larger black threat for closing games.
- 1 Old Flitterfang | threat | A creature threat that fits the deck's overall plan.
- 1 Ruthless Technomancer | threat | A high-impact creature threat for the midgame.
- 1 Shilgengar, Sire of Famine | threat | A black legendary creature threat.
- 1 Sivriss, Nightmare Speaker | threat | A low-cost black creature threat.
- 1 Skullport Merchant | threat | A creature threat that complements the sacrifice shell.
- 1 Vampire Gourmand | threat | A low-cost Vampire creature threat.
- 1 Vengeful Bloodwitch | threat | A Vampire threat for applying pressure.

</details>

### 17. upgrade a precon, owned first

Format: Commander. Theme: goblins. Pool: owned_first. Shortlist: 196 names.

Commander: Zada, Hedron Grinder.

Grade: bad, score 0.00, model 20260903T210658Z.

Cards: 0 main, 0 sideboard. Repair turn: yes, for profile_off_band, profile_off_band, two_card_combo. Block findings: 1.

Cost: $0.00 to buy, $0.00 the whole deck.

**Summary:** No legal repaired deck can be constructed from the supplied data: the shortlist marks only 62 nonbasic card names as precon, while the required minimum is 67. It also does not include the previously returned deck that is to be repaired. Please provide the missing precon-marked entries or correct the retention requirement, along with the prior decklist.

The quality model grades this deck below the precon baseline against the top lists of the format: the deck makes less mana on turn four than the norm, and that lowers the grade. Fewer opening hands hold two to four lands than the norm, and that lowers the grade. The curve sits low for the format, and that raises the grade.

- JUDGE [unknown]: "No legal repaired deck can be constructed from the supplied data: the shortlist marks only 62 nonbasic card names as precon, while the required minimum is 67.". This asserts a legality/minimum-count condition for the deck. No rule of Magic requires 67 nonbasic precon-marked cards; the figure appears to come from a user-supplied retention requirement rather than the game's rules, so its correctness cannot be judged against the real rules.
- [BLOCK] `deck_size`: deck has 1 cards, the format needs exactly 100 (commander included)
- [WARN] `land_count`: 0 lands: the guide range for this format is 27 to 41
- [WARN] `precon_share`: the deck keeps 1 of the 78 nonbasic Goblin Storm precon names, and the rule asks for 67
- [WARN] `profile_off_band`: the land count is 0, and bracket 3 wants 34 to 38
- [WARN] `profile_off_band`: the average mana value of the nonland cards is 0, and bracket 3 wants 2 to 3.5
- [WARN] `profile_off_band`: the ramp count is 0, and bracket 3 wants 8 to 13
- [WARN] `profile_off_band`: the draw count is 0, and bracket 3 wants 8 to 14
- [WARN] `profile_off_band`: the removal count is 0, and bracket 3 wants 6 to 12
- [WARN] `profile_off_band`: the wipe count is 0, and bracket 3 wants 2 to 5
- [WARN] `profile_off_band`: the interaction count is 0, and bracket 3 wants 4 to 12

<details><summary>The deck list</summary>


</details>

### 18. upgrade a precon, any card

Format: Commander. Theme: turtles. Pool: any_card. Shortlist: 300 names.

Commander: Heroes in a Half Shell.

Grade: typical, score 0.51, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band, profile_off_band. Block findings: 0.

Cost: $78.07 to buy, $78.07 the whole deck.

**Summary:** A five-color Turtle-themed creature deck that develops its mana, builds a board of heroes and allies, and converts that board into decisive combat pressure. It supports the creature plan with resilient card advantage, versatile answers, and sweepers that let the team recover when the table becomes hostile.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The cards pair the way the top lists pair them, and that raises the grade. Many lands enter tapped, and that raises the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 3.32 over 63 nonland cards
- [WARN] `profile_off_band`: the removal count is 5, and bracket 3 wants 6 to 12

<details><summary>The deck list</summary>

- 1 Acidic Slime | removal | Flexible permanent interaction attached to a creature.
- 1 April O'Neil, Live on the Scene | draw | Repeatable card advantage for the team-focused game plan.
- 1 Arcade Cabinet | synergy | Retains the precon's themed artifact package.
- 1 Arcane Signet | ramp | Reliable multicolor acceleration.
- 1 Ash Barrens | land | Fixes colors while remaining a land slot.
- 1 Assassin's Trophy | removal | Answers problematic permanents cleanly.
- 1 Baxter, Fly in the Ointment | synergy | Retains the precon's mutant support.
- 1 Bebop, Skull & Crossbones | threat | Retains a themed aggressive threat.
- 1 Big Apple, 3 a.m. | land | Themed color fixing.
- 1 Big Mother Mouser | synergy | Retains a themed artifact creature.
- 1 Biogenic Ooze | threat | Produces a growing board presence.
- 1 Blasphemous Act | wipe | Efficient reset against creature-heavy boards.
- 1 Casey Jones, Back Alley Brute | threat | Retains the precon's combat package.
- 1 Chromatic Lantern | ramp | Fixes demanding colors while accelerating.
- 1 Cinder Glade | land | On-theme multicolor land.
- 1 Coin of Mastery | ramp | Supports the precon's artifact subtheme and mana development.
- 1 Command Tower | land | Unconditional multicolor fixing.
- 1 Continue? | draw | Retains a thematic source of card flow.
- 1 Corpsejack Menace | synergy | Strengthens counter-based board development.
- 1 Dimension X Pizzasaur | synergy | Retains the themed artifact package.
- 1 Donatello, the Brains | synergy | Core Turtle-themed payoff.
- 1 Double Jump // Flying Kick | interaction | Themed combat trick with flexibility.
- 1 Electric Seaweed | interaction | Retains the precon's defensive utility.
- 1 Endless Foot Assault | synergy | Themed pressure and board development.
- 1 Escape Tunnel | land | Basic-land access and fixing.
- 1 Everything Pizza | synergy | Retains the Food-themed support package.
- 1 Evolving Wilds | land | Color fixing through basic-land access.
- 1 Exploding Barrel | ramp | Themed mana support.
- 1 Fast Forward | draw | Maintains the precon's card-selection package.
- 1 Foot Chopper | synergy | Retains themed equipment support.
- 1 Game Over | threat | Themed high-impact payoff.
- 1 Grand Coliseum | land | Broad color fixing.
- 1 Harmonize | draw | Straightforward card advantage.
- 1 Here Comes a New Hero! | draw | Themed card advantage and development.
- 1 High Score | draw | Supports sustained card flow.
- 1 Hidden Hideout | land | Themed multicolor land.
- 1 Irma, Part-Time Mutant | synergy | Retains the precon's mutant theme.
- 1 Krang, the All-Powerful | threat | Retains a major themed artifact payoff.
- 1 Leatherhead, Iron Gator | threat | Retains the themed creature suite.
- 1 Leonardo, the Balance | threat | Core Turtle-themed threat.
- 1 Lessons from Life | draw | Themed card advantage.
- 1 Level Up | synergy | Supports creature development in combat.
- 1 Lita, Little Orphan Amphibian | synergy | Retains the Turtle and mutant theme.
- 1 Michelangelo, the Heart | synergy | Core Turtle-themed payoff.
- 1 Mole Module | synergy | Retains the artifact subtheme.
- 1 Mona Lisa, Science Geek | synergy | Retains the themed creature suite.
- 1 Ninja Pizza | synergy | Themed support for the creature plan.
- 1 Path of Ancestry | land | Tribal color fixing.
- 1 Rain-Slicked Copse | land | Themed multicolor land.
- 1 Raphael, the Muscle | threat | Core Turtle-themed threat.
- 1 Rat King, Pale Piper | threat | Retains the precon's mutant roster.
- 1 Ray Fillet, Wave Warrior | threat | Retains the themed creature suite.
- 1 Roadkill Rodney | synergy | Retains artifact-creature support.
- 1 Rocksteady, Mutant Marauder | threat | Retains the precon's aggressive mutant theme.
- 1 Rootbound Crag | land | Reliable color fixing.
- 1 Shellshock | interaction | Themed tactical interaction.
- 1 Shredder, Shadow Master | threat | Retains the precon's antagonist package.
- 1 Sodden Verdure | land | Themed multicolor land.
- 1 Sol Ring | ramp | Efficient early acceleration.
- 1 Special Move | interaction | Themed tactical interaction.
- 1 Steelbane Hydra | removal | Creature-based artifact and enchantment interaction.
- 1 Super Combo | wincon | Themed finishing pressure.
- 1 Swift Demise | removal | Themed spot interaction.
- 1 Tempestra, Dame of Games | synergy | Retains the themed creature package.
- 1 Thriving Grove | land | Flexible color fixing.
- 1 Thriving Isle | land | Flexible color fixing.
- 1 Thriving Moor | land | Flexible color fixing.
- 1 Together Forever | interaction | Protects and recovers important creatures.
- 1 Tokka & Rahzar, Unsupervised | ramp | Themed creature-based acceleration.
- 1 Turtle Lair | land | Tribal land support.
- 1 Vanquish the Horde | wipe | Board reset when opponents get ahead.
- 1 Vernal Fen | land | Themed multicolor land.
- 1 Vibrant Cityscape | land | Flexible color fixing.
- 1 Voracious Hydra | threat | Scalable creature and removal option.
- 1 Wave Goodbye | wipe | Themed board-control effect.
- 1 Breeding Pool | land | High-quality blue and green fixing.
- 1 Godless Shrine | land | High-quality white and black fixing.
- 1 Overgrown Tomb | land | High-quality black and green fixing.
- 1 Sacred Foundry | land | High-quality red and white fixing.
- 1 Steam Vents | land | High-quality blue and red fixing.
- 1 Temple Garden | land | High-quality green and white fixing.
- 1 Watery Grave | land | High-quality blue and black fixing.
- 1 Farseek | ramp | Efficient multicolor land acceleration.
- 1 Nature's Lore | ramp | Efficient land-based acceleration.
- 1 Beast Whisperer | draw | Turns the creature-heavy plan into sustained card advantage.
- 1 Garruk's Uprising | draw | Provides ongoing card advantage and rewards large creatures.
- 1 Counterspell | interaction | Efficient protection from pivotal opposing spells.
- 1 Swords to Plowshares | removal | Premium, low-cost creature removal.
- 3 Forest | land | Green sources for creature spells and land ramp.
- 3 Island | land | Blue sources for utility and interaction.
- 2 Plains | land | White sources for support and answers.
- 2 Swamp | land | Black sources for removal and multicolor costs.
- 1 Mountain | land | Red source for themed spells and multicolor costs.

</details>

### 19. the Hobbit family, two colours

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 174 names.

Commander: Thranduil, the Elvenking.

Grade: bad, score 0.04, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band. Block findings: 0.

Cost: $444.32 to buy, $444.32 the whole deck.

**Summary:** Thranduil leads a Sultai Middle-earth court built around an Elf-heavy woodland board, legendary characters, and a steady flow of storybook value. Creature-based mana establishes the realm early, while card advantage keeps the court supplied with allies and answers. The deck protects its key figures, disrupts opposing plans, and converts a developed board of Elves, legends, and imposing denizens of the wild into sustained battlefield pressure.

The quality model grades this deck below the precon baseline against the top lists of the format: more opening hands hold two to four lands than the norm, and that raises the grade. Many of the cards appear in no top list, and that lowers the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 3.48 over 63 nonland cards
- [WARN] `profile_off_band`: the worst color's share of the sources it needs is 0.8, and bracket 3 wants 0.85 or more (sources of requirement: U 17.5 of 20, B 18.5 of 23, G 21 of 22)
- [INFO] `basics_added`: the list was 1 card short, so the builder added 1 basic land

<details><summary>The deck list</summary>

- 9 Island | land | Blue basics support the deck's information, evasive creatures, and reactive spells.
- 9 Swamp | land | Black basics support removal and the darker legends in the deck.
- 10 Forest | land | Green basics support Elf development and creature-based mana production.
- 1 Elven Passage | land | An on-theme colored land for the Elf-focused mana base.
- 1 Elvenking's Halls | land | A thematic fixing land for the commander’s court.
- 1 Hobbit Hole | land | A colored Middle-earth land that broadens the mana base.
- 1 Minas Morgul, Dark Fortress | land | A black-oriented legendary land for the deck’s darker spells.
- 1 Mirkwood | land | A thematic land supporting the deck’s woodland core.
- 1 Rivendell | land | A blue-oriented legendary land for Elf and Wizard magic.
- 1 The Black Gate | land | A thematic black land that strengthens color coverage.
- 1 The Shire | land | A colored legendary land that rounds out the Middle-earth mana base.
- 1 Arcane Signet | ramp | Efficient universal color fixing.
- 1 Bag End Banquet | ramp | Provides thematic artifact acceleration.
- 1 Delighted Halfling | ramp | Early mana development with legendary-spell support.
- 1 Elven Chorus | ramp | Turns the Elf-heavy creature suite into mana access.
- 1 Elvish Archdruid | ramp | A major Elf tribal mana producer.
- 1 Elvish Mystic | ramp | Low-cost creature acceleration.
- 1 Mox Amber | ramp | Legendary-focused acceleration for the deck’s many named characters.
- 1 Wayfarer's Bauble | ramp | Reliable basic-land ramp and fixing.
- 1 Wood Elves | ramp | Develops mana while adding an Elf to the board.
- 1 Woodland Weavemaster | ramp | An on-theme creature-based mana source.
- 1 Beorn the Fierce | draw | A resilient green source of card advantage.
- 1 Elvish Visionary | draw | A cheap Elf that replaces itself.
- 1 Fateful Discovery | draw | Provides sustained access to additional cards.
- 1 Gandalf, Shadow's Foe | draw | A legendary value creature that contributes cards.
- 1 Hithlain Knots | draw | Flexible instant-speed card advantage.
- 1 Last March of the Ents | draw | A powerful green refill for a creature-heavy board.
- 1 Lórien Revealed | draw | Early mana smoothing that becomes late-game card advantage.
- 1 Night's Whisper | draw | Efficient black card selection.
- 1 Palantír of Orthanc | draw | A flavorful artifact engine for steady advantage.
- 1 Plunder the Trollshaws | draw | Instant-speed access to additional resources.
- 1 Uncover the Moon-Letters | draw | A thematic enchantment-based draw engine.
- 1 Bitter Downfall | removal | Clean black answer to a problematic permanent.
- 1 Bilbo's Deadly Slice | removal | Cheap thematic spot removal.
- 1 Crude Bent Blade | removal | An equipment-based answer that remains useful on board.
- 1 Enchanted River's Grasp | removal | Blue removal that disrupts opposing development.
- 1 Giant's Boulder | removal | Versatile artifact removal for troublesome targets.
- 1 Merciless Executioner | removal | Creature-based sacrifice pressure against protected threats.
- 1 Orcish Bowmasters | removal | Efficient creature control with punishing incidental value.
- 1 Quarrel | removal | Low-cost interaction for opposing creatures or plans.
- 1 Uneasy Partings | removal | Instant-speed disruption that fits the deck’s colors.
- 1 Bilbo's Ring | interaction | Protects important creatures and interferes with combat.
- 1 Confusticate and Bebother | interaction | A flexible blue response to opposing spells or actions.
- 1 Elrond, Moon-Reader | interaction | A legendary Elf that supplies meaningful stack interaction.
- 1 Mithril Coat | interaction | Keeps the commander or a key legend safe.
- 1 My Precious // Allure of Power | interaction | A versatile protective equipment with an interactive adventure.
- 1 Old Fat Spider Can't See Me | interaction | A thematic saga that disrupts opponents’ attacks and plans.
- 1 Stern Scolding | interaction | Efficient blue protection against early opposing threats.
- 1 Thranduil's Decree | interaction | Flavorful instant-speed interference for the Elven court.
- 1 Gnashing of Teeth | wipe | A broad reset for crowded creature boards.
- 1 Languish | wipe | Efficient black creature sweeper.
- 1 Raise the Palisade | wipe | A tribal-conscious blue board reset that favors the deck’s creature type.
- 1 Boughside Wanderers | synergy | Elf tribal body that benefits from the deck’s woodland development.
- 1 Cantankerous Keepers | synergy | Adds another Elf presence for tribal payoffs.
- 1 Celeborn the Wise | synergy | A legendary Elf that reinforces the courtly theme.
- 1 Galadhrim Guide | synergy | Supports the deck’s Elf-forward battlefield plan.
- 1 Galion, Elvenking's Butler | synergy | A courtly legendary Elf that complements Thranduil’s retinue.
- 1 Lothlórien Lookout | synergy | An Elf that advances the deck’s coordinated creature plan.
- 1 Mirkwood Meditator | synergy | A thematic Elf body for tribal development.
- 1 Mirkwood Nurturer | synergy | Adds to the Elf core while supporting board presence.
- 1 Mirkwood Pathmaker | synergy | An Elf that helps establish a connected woodland board.
- 1 Thranduil, Sindarin Liege // Silvan Rally | synergy | A second Thranduil card that strongly reinforces the Silvan theme.
- 1 Attercop | threat | A sizable spider threat that pressures the board.
- 1 Colossal Whale | threat | A large blue finisher that also pressures opposing creatures.
- 1 Gollum the Abandoned | threat | A dangerous black legend that adds pressure and flavor.
- 1 Gollum, Riddle Master | threat | A slippery legendary threat with value potential.
- 1 Haunt of the Dead Marshes | threat | A menacing Elf-adjacent threat for the darker side of the deck.
- 1 Mirkwood Elk | threat | A solid green attacker that suits the woodland setting.
- 1 Nimrodel Watcher | threat | An Elf threat that complements the deck’s tribal battlefield.
- 1 Old Fat Spider | threat | A thematic creature that can become a meaningful board presence.
- 1 Troll of Khazad-dûm | threat | A large black threat with a strong Middle-earth identity.
- 1 Witch-king of Angmar | threat | A powerful legendary black finisher.
- 1 Witch-king, Bringer of Ruin | threat | A second imposing Wraith legend that closes games through pressure.
- 1 Willow-Wind | threat | A substantial green creature that helps finish through combat.

</details>

### 20. the Hobbit family, a delegated commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 77 names.

Commander: Smaug the Magnificent.

Grade: bad, score 0.05, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band. Block findings: 0.

Cost: $465.27 to buy, $465.27 the whole deck.

**Summary:** A treasure-fueled dragon deck that uses artifacts, legendary equipment, and forceful combat to build toward imposing attacks. It balances its threatening creature suite with resilient resource engines, direct disruption, and dramatic battlefield resets.

The quality model grades this deck below the precon baseline against the top lists of the format: more opening hands hold two to four lands than the norm, and that raises the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 3.24 over 63 nonland cards

<details><summary>The deck list</summary>

- 31 Mountain | land | Reliable red mana base for a mono-red dragon strategy.
- 1 Dragon-Cursed Halls | land | Thematic utility land that supports the dragon plan.
- 1 Hobbit Hole | land | Utility land slot with minimal strain on the mana base.
- 1 Rogue's Passage | land | Helps a large attacker connect through stalled boards.
- 1 The Lonely Mountain | land | Thematic Mountain land for the deck's primary color.
- 1 Treasure Vault | land | Flexible artifact land that can convert spare mana into resources.
- 1 Arcane Signet | ramp | Efficient color fixing and acceleration.
- 1 Bag End Banquet | ramp | Early artifact-based mana acceleration.
- 1 Burn, Burn, Tree and Fern | ramp | Thematic mana development.
- 1 Cavern-Hoard Dragon | ramp | A dragon that turns combat into mana resources.
- 1 Dragon's Desire | ramp | Supports deployment of costly threats.
- 1 Fíli and Kíli, Joyous | ramp | Creature-based acceleration with thematic support.
- 1 Long-Bodied Grey Dog | ramp | Low-cost creature that advances mana development.
- 1 Mox Amber | ramp | Fast legendary-focused acceleration.
- 1 The Reaver Cleaver | ramp | Turns combat damage into a substantial mana advantage.
- 1 Wayfarer's Bauble | ramp | Reliable land-based mana development.
- 1 Balin, Loremaster | draw | Repeatable card advantage on a thematic creature.
- 1 Key to the Side-Door | draw | Artifact-based card selection and advantage.
- 1 Palantír of Orthanc | draw | Persistent source of cards and pressure.
- 1 Ragged Short Spear | draw | Equipment-based card advantage.
- 1 Thrór's Map | draw | Provides steady access to additional resources.
- 1 Óin the Brave | draw | Creature-based card advantage.
- 1 Glamdring | draw | Equipment that rewards combat with cards.
- 1 Sting, Bilbo's Sword | draw | Combat equipment that provides ongoing value.
- 1 Bilbo's Ring | interaction | Protects an important creature from opposing interaction.
- 1 Dwarven Mattock | interaction | Flexible equipment-based disruption.
- 1 Mithril Coat | interaction | Protects the commander or a key threat.
- 1 The One Ring | interaction | Provides a resilient protective resource engine.
- 1 Smaug's Fury | interaction | Instant-speed combat and board influence.
- 1 Giant's Boulder | interaction | Versatile answer that can disrupt opposing permanents.
- 1 Improvised Club | interaction | Efficient instant-speed answer.
- 1 Pinecone Strike | interaction | Low-cost reactive disruption.
- 1 Battle-Scarred Goblin | removal | Creature-based answer to problematic targets.
- 1 Fire of Orthanc | removal | Direct answer to an opposing permanent.
- 1 Gandalf, Spark Starter | removal | Repeatable removal attached to a creature.
- 1 Goblin Cratermaker | removal | Flexible creature or artifact removal.
- 1 Goblin Fireleaper | removal | Thematic creature-based removal.
- 1 Inferno Titan | removal | Large threat that removes smaller opposing creatures.
- 1 Smaug, the Great Calamity // Spew Flame | removal | Dragon threat with a removal spell attached.
- 1 Smite the Deathless | removal | Efficient answer to a dangerous opposing creature.
- 1 Stone-Giant of High Pass | removal | Creature that supplies removal while advancing the board.
- 1 Call Forth the Tempest | wipe | Resets crowded opposing boards.
- 1 Desolation of Smaug | wipe | Thematic mass removal for overwhelming boards.
- 1 Glóin the Mighty // Easy Pickings | wipe | Flexible creature that can access a board-clearing spell.
- 1 Last Light of Durin's Day | synergy | Supports the deck's legendary and combat-focused plan.
- 1 Andúril, Flame of the West | synergy | Equipment that improves the deck's combat plan.
- 1 Andúril, Narsil Reforged | synergy | Legendary equipment supporting large attacking creatures.
- 1 Getaway Barrel | synergy | Artifact utility that supports the deck's resource-focused plan.
- 1 Long-Lost Lances | synergy | Equipment that rewards committing large creatures to combat.
- 1 Well-Worn Spatula | synergy | Low-cost equipment support for combat threats.
- 1 The Black Arrow | synergy | Thematic legendary equipment supporting the attack plan.
- 1 Tidings of War | synergy | Builds toward forceful combat turns.
- 1 Thorin, Mountain-king | synergy | Legendary creature that supports the deck's thematic core.
- 1 Dwarven Mauler | synergy | Benefits from the deck's equipment and combat emphasis.
- 1 Desert Were-Worm | threat | Large evasive creature that pressures life totals.
- 1 Guttersnipe | threat | Converts the deck's noncreature spells into additional pressure.
- 1 Gundabad Opportunist | threat | Efficient creature that contributes to battlefield pressure.
- 1 Misty Mountains Raider | threat | Aggressive creature for maintaining pressure.
- 1 Olog-hai Crusher | threat | Large body suited to closing games through combat.
- 1 Oliphaunt | threat | High-impact creature that provides a substantial attacking body.
- 1 Orcish Siegemaster | threat | Thematic attacker that strengthens the creature suite.
- 1 Snowslope Hunter | threat | Creature threat that helps keep pressure on opponents.
- 1 Troop of Ponies | threat | Creature presence that can contribute to decisive combat.
- 1 Old Thrush | threat | Low-cost creature that supports a broad attacking board.
- 1 Dwarven Warriors | threat | Thematic body for equipment and combat support.
- 1 Iron Hills Stalwart | threat | Durable attacker that complements the creature plan.
- 1 Bombur, Gentle Dreamer | other | Thematic legendary utility creature.
- 1 Bothersome Noisemaker | other | Flexible thematic utility creature.
- 1 Gandalf, Goblins' Bane // Flameshape | other | Versatile legendary spell-creature for the deck's broader game plan.

</details>

### 21. the Hobbit family, mana from outside

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 156 names.

Commander: Smaug the Impenetrable.

Grade: bad, score 0.03, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band, profile_off_band. Block findings: 0.

Cost: $1042.10 to buy, $1042.10 the whole deck.

**Summary:** A Rakdos Smaug deck built around a treasure-rich equipment package, forceful combat, and a cast of Dragons, Goblins, Trolls, and legendary villains. Efficient mana development supports towering threats, while plentiful removal and sweeping effects keep opposing boards from overwhelming the mountain-hoard strategy.

The quality model grades this deck below the precon baseline against the top lists of the format: more opening hands hold two to four lands than the norm, and that raises the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. The deck holds no fast mana, and that lowers the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 3.24 over 63 nonland cards
- [WARN] `outside_requested_set`: the sets you named do not hold 17 cards: Ancient Tomb, Arid Mesa, Blood Crypt, Bloodstained Mire, Command Tower, Dragonskull Summit, Exotic Orchard, Flooded Strand, Marsh Flats, Misty Rainforest, Polluted Delta, Scalding Tarn, Urborg, Tomb of Yawgmoth, Verdant Catacombs, Windswept Heath, Wooded Foothills, Yavimaya, Cradle of Growth

<details><summary>The deck list</summary>

- 1 Ancient Tomb | land | Fast colorless acceleration for deploying Smaug and expensive threats.
- 1 Arid Mesa | land | Efficiently finds the deck's basic Mountains or Swamps.
- 1 Blood Crypt | land | Reliable dual-color source that can enter untapped.
- 1 Bloodstained Mire | land | Untapped fetch land for either basic color or the dual land.
- 1 Command Tower | land | Consistent access to both commander colors.
- 1 Dragonskull Summit | land | Dual land with strong untapped consistency in this manabase.
- 1 Exotic Orchard | land | Usually supplies either required color in multiplayer games.
- 1 Flooded Strand | land | Fetches either basic color through land types.
- 1 Marsh Flats | land | Fetches either basic color through land types.
- 1 Misty Rainforest | land | Fetches either basic color through land types.
- 1 Polluted Delta | land | Fetches either basic color through land types.
- 1 Scalding Tarn | land | Fetches either basic color through land types.
- 1 Urborg, Tomb of Yawgmoth | land | Improves black access while letting utility lands contribute colored mana.
- 1 Verdant Catacombs | land | Fetches either basic color through land types.
- 1 Windswept Heath | land | Fetches either basic color through land types.
- 1 Wooded Foothills | land | Fetches either basic color through land types.
- 1 Yavimaya, Cradle of Growth | land | Lets utility lands contribute colored mana while maintaining speed.
- 1 The Lonely Mountain | land | On-theme red source for the deck's dragon and Goblin elements.
- 10 Mountain | land | Primary red source for Smaug and the deck's aggressive spells.
- 8 Swamp | land | Stable black source for removal and resilient threats.
- 1 Arcane Signet | ramp | Efficient two-color acceleration.
- 1 Mox Amber | ramp | Zero-cost acceleration supported by the deck's legendary permanents.
- 1 Wayfarer's Bauble | ramp | Early permanent land acceleration and color fixing.
- 1 Bag End Banquet | ramp | Thematic mana development for the midgame.
- 1 Bolg's Company | ramp | Creature-based acceleration that supports the Goblin contingent.
- 1 Burn, Burn, Tree and Fern | ramp | Saga-based mana development that advances the deck's larger plays.
- 1 Dragon's Desire | ramp | Accelerates into Smaug and other costly threats.
- 1 Orcrist, Goblin-cleaver | ramp | Equipment-based mana generation that rewards combat.
- 1 The Misty Mountains Cold | ramp | Thematic saga acceleration for the deck's top end.
- 1 The Reaver Cleaver | ramp | Turns combat damage into a substantial mana burst.
- 1 Balin, Loremaster | draw | Legendary creature that provides recurring card advantage.
- 1 Gollum, Riddle Master | draw | Thematic source of ongoing card selection.
- 1 Key to the Side-Door | draw | Low-cost artifact card advantage.
- 1 Night's Whisper | draw | Efficient early card flow.
- 1 Palantír of Orthanc | draw | Persistent artifact-based card advantage.
- 1 Rage into the Valley | draw | Refills resources while advancing the aggressive plan.
- 1 Ragged Short Spear | draw | Equipment that contributes card advantage.
- 1 Reverent Howl | draw | Flexible instant-speed card advantage.
- 1 The Master of Lake-town | draw | Creature-based draw engine for longer games.
- 1 Thrór's Map | draw | Artifact card advantage that fits the deck's legendary package.
- 1 Óin the Brave | draw | Thematic legendary source of card advantage.
- 1 Azog, Moria's Ruin | removal | Threatening legendary creature that answers opposing permanents.
- 1 Battle-Scarred Goblin | removal | Low-curve creature removal attached to a thematic body.
- 1 Bilbo's Deadly Slice | removal | Efficient targeted removal.
- 1 Bitter Downfall | removal | Flexible answer to problematic creatures or planeswalkers.
- 1 Fire of Orthanc | removal | Direct red removal for creatures and key targets.
- 1 Goblin Cratermaker | removal | Versatile removal on a Goblin body.
- 1 Improvised Club | removal | Low-cost instant-speed removal.
- 1 Orcish Bowmasters | removal | Efficient creature interaction with incidental pressure.
- 1 Pinecone Strike | removal | Flexible instant-speed answer.
- 1 Bolg, Erebor's Reckoning | wipe | Thematic creature-based reset option.
- 1 Call Forth the Tempest | wipe | Broad battlefield reset when opponents pull ahead.
- 1 Languish | wipe | Efficient sweeper against creature-heavy boards.
- 1 Along the Crooked Way | interaction | Flexible protective and disruptive support for the central game plan.
- 1 Bilbo's Ring | interaction | Protective legendary equipment with useful utility.
- 1 Dwarven Mattock | interaction | Equipment-based interaction that remains useful in combat.
- 1 Getaway Barrel | interaction | Artifact utility for preserving important permanents.
- 1 Mithril Coat | interaction | Protects Smaug or another key legendary threat.
- 1 My Precious // Allure of Power | interaction | Versatile legendary equipment and instant-speed disruption.
- 1 Smaug's Fury | interaction | Instant-speed combat interaction and pressure.
- 1 The One Ring | interaction | Legendary protection and sustained resource advantage.
- 1 Andúril, Flame of the West | synergy | Legendary equipment that strengthens the combat-focused plan.
- 1 Andúril, Narsil Reforged | synergy | A thematic weapon that rewards keeping a threat on the battlefield.
- 1 Crude Bent Blade | synergy | Cheap equipment supporting creature combat and battlefield control.
- 1 Glamdring | synergy | Legendary equipment that reinforces the artifact package.
- 1 Goblin Plate Mail | synergy | Supports the Goblin subtheme while improving combat.
- 1 Last Light of Durin's Day | synergy | Thematic enchantment support for the deck's creature plan.
- 1 Long-Lost Lances | synergy | Equipment that helps large creatures convert attacks into pressure.
- 1 Sting, Bilbo's Sword | synergy | Efficient legendary equipment for evasive combat value.
- 1 Supper for Spiders | synergy | Thematic utility that complements the deck's battlefield plan.
- 1 Well-Worn Spatula | synergy | Low-cost equipment that supports the artifact-and-combat theme.
- 1 Cavern-Hoard Dragon | threat | Large Dragon threat that supplies additional resources after connecting.
- 1 Desert Were-Worm | threat | Large, flavorful combat threat for closing games.
- 1 Gandalf, Goblins' Bane // Flameshape | threat | Versatile legendary threat with an attached spell option.
- 1 Great Goblin, Foul-Hearted | threat | Goblin leader that adds resilient battlefield pressure.
- 1 Great Ugly-Looking Goblin // Clap! Snap! | threat | Creature threat with flexible adventure value.
- 1 Guttersnipe | threat | Converts the deck's instants and sorceries into repeatable damage.
- 1 Haunt of the Dead Marshes | threat | Evasive creature pressure with a dark Middle-earth theme.
- 1 Inferno Titan | threat | High-impact finisher that adds immediate board pressure.
- 1 Olog-hai Crusher | threat | Large Troll body for ending stalled games.
- 1 Sauron, the Lidless Eye | threat | Legendary late-game threat fitting the deck's dark-power theme.
- 1 Smaug, the Great Calamity // Spew Flame | threat | Dragon threat with flexible removal attached.
- 1 Troll of Khazad-dûm | threat | Durable, thematic top-end creature.

</details>

### 22. a set family and a card from outside it

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 175 names.

Commander: Thranduil, the Elvenking.

Grade: bad, score 0.03, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band, profile_off_band. Block findings: 0.

Cost: $294.11 to buy, $294.11 the whole deck.

**Summary:** A Sultai Elf-tribal Commander deck that develops its mana through creatures and artifacts, then builds a resilient board of legendary figures, scouts, and other woodland forces. It combines steady card advantage with broad answers and tribal-friendly sweepers, aiming to establish battlefield control before converting its assembled creature force into sustained combat pressure.

The quality model grades this deck below the precon baseline against the top lists of the format: many of the cards appear in no top list, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 2.97 over 63 nonland cards
- [WARN] `profile_off_band`: the worst color's share of the sources it needs is 0.64, and bracket 3 wants 0.85 or more (sources of requirement: U 17.25 of 19, B 18.25 of 19, G 12.75 of 20)
- [WARN] `outside_requested_set`: the sets you named do not hold 1 card: Sol Ring

<details><summary>The deck list</summary>

- 12 Island | land | Blue sources for the deck’s reactive spells and card advantage.
- 12 Swamp | land | Reliable black sources for removal and graveyard-oriented creatures.
- 6 Forest | land | Green sources supporting Elf deployment and mana development.
- 1 Elven Passage | land | On-theme dual-color land.
- 1 Elvenking's Halls | land | On-theme blue-green fixing.
- 1 Mirkwood | land | Black-green fixing for the creature base.
- 1 Rivendell | land | Blue fixing in an Elf-themed land slot.
- 1 Minas Morgul, Dark Fortress | land | Black fixing for the removal suite.
- 1 The Black Gate | land | Additional black source that remains within the tapped-land limit.
- 1 Sol Ring | ramp | Efficient colorless acceleration retained for early development.
- 1 Arcane Signet | ramp | Reliable three-color mana fixing.
- 1 Elvish Archdruid | ramp | Elf-scaled mana production that rewards the tribal board.
- 1 Elvish Mystic | ramp | Low-cost green acceleration.
- 1 Mox Amber | ramp | Efficient mana acceleration alongside the deck’s legends.
- 1 Wayfarer's Bauble | ramp | Fixes basic-land access while advancing mana.
- 1 Wood Elves | ramp | Land-based ramp attached to an Elf body.
- 1 Woodland Weavemaster | ramp | Creature-based mana development for the Elf shell.
- 1 Thranduil's Company | ramp | Thematic mana development that adds to the board.
- 1 Through the Forest Gate | ramp | Color fixing and land ramp.
- 1 Elvish Visionary | draw | Low-cost card replacement on a relevant Elf body.
- 1 Night's Whisper | draw | Efficient unconditional card draw.
- 1 Fateful Discovery | draw | Thematic ongoing card advantage.
- 1 Gollum, Riddle Master | draw | Creature-based card advantage.
- 1 Hithlain Knots | draw | Instant-speed card draw.
- 1 Ithilien Kingfisher | draw | Evasive creature that supplies cards.
- 1 Key to the Side-Door | draw | Utility artifact that converts into card advantage.
- 1 Palantír of Orthanc | draw | Repeatable source of card selection and advantage.
- 1 Plunder the Trollshaws | draw | Green card advantage for a creature-focused deck.
- 1 Reverent Howl | draw | Flexible instant-speed draw.
- 1 Uncover the Moon-Letters | draw | Persistent thematic card advantage.
- 1 Confusticate and Bebother | interaction | Flexible blue disruption.
- 1 Elrond, Moon-Reader | interaction | Legendary Elf utility and stack interaction.
- 1 Mithril Coat | interaction | Protects the commander or a key threat.
- 1 My Precious // Allure of Power | interaction | Protective equipment with flexible interaction.
- 1 Old Fat Spider Can't See Me | interaction | Defensive disruption that supports the deck’s plan.
- 1 Sound the Trumpets | interaction | Instant-speed protection and tactical interaction.
- 1 Stern Scolding | interaction | Cheap answer to an opposing early play.
- 1 Thranduil's Decree | interaction | Thematic interactive spell for protecting tempo.
- 1 Bilbo's Deadly Slice | removal | Efficient creature removal.
- 1 Bitter Downfall | removal | Versatile black permanent removal.
- 1 Crude Bent Blade | removal | Reusable equipment-based removal utility.
- 1 Enchanted River's Grasp | removal | Blue-based answer to troublesome permanents.
- 1 Merciless Executioner | removal | Creature-based sacrifice removal.
- 1 Orcish Bowmasters | removal | Efficient creature control with incidental pressure.
- 1 Quarrel | removal | Low-cost interactive removal.
- 1 The Black Arrow | removal | Thematic removal attached to an equipment slot.
- 1 Uneasy Partings | removal | Flexible instant-speed answer.
- 1 Gnashing of Teeth | wipe | Board reset suited to an attrition plan.
- 1 Languish | wipe | Efficient creature-board sweeper.
- 1 Raise the Palisade | wipe | Tribal-friendly mass bounce that preserves Elf development.
- 1 Boughside Wanderers | synergy | Elf body that reinforces the tribal battlefield.
- 1 Cantankerous Keepers | synergy | Elf-based board presence for tribal payoffs.
- 1 Galadhrim Guide | synergy | Low-cost Elf that advances the creature theme.
- 1 Galion, Elvenking's Butler | synergy | Legendary Elf support for the deck’s central tribe.
- 1 Guardian of the Halls | synergy | Thematic Elf body with battlefield utility.
- 1 Lothlórien Lookout | synergy | Elf tribal support at a low mana cost.
- 1 Mirkwood Meditator | synergy | On-theme Elf that builds the board.
- 1 Mirkwood Nurturer | synergy | Elf synergy that supports the developing battlefield.
- 1 Mirkwood Pathmaker | synergy | Tribal creature that improves board cohesion.
- 1 Nimrodel Watcher | synergy | Elf-focused utility creature.
- 1 Arwen, Weaver of Hope | threat | Resilient legendary threat that strengthens the creature board.
- 1 Attercop | threat | Large black threat that pressures opposing boards.
- 1 Dreaded Bat-Cloud | threat | Evasive black creature for sustained pressure.
- 1 Duskwatch Hunter | threat | Efficient creature threat for combat-focused turns.
- 1 Great Fierce Bee | threat | Green creature pressure that works with a broad board.
- 1 Haunt of the Dead Marshes | threat | Graveyard-adjacent threat with an on-theme creature type.
- 1 Nighthowl Pursuer | threat | Black creature threat with combat utility.
- 1 Ravening Warg | threat | Aggressive creature that pressures life totals.
- 1 Rhovanion Rampager | threat | Green combat threat for closing games.
- 1 The Chief Warg | threat | Legendary threat that adds value while attacking.
- 1 Wargling | threat | Low-cost creature pressure.
- 1 Wilderland Scrounger | threat | Creature threat that supports the deck’s board-forward plan.

</details>

### 23. two set families at once

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 203 names.

Commander: Kíli the Resourceful.

Grade: typical, score 0.47, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $274.75 to buy, $274.75 the whole deck.

**Summary:** Kíli the Resourceful leads a board-focused white Dwarf, artifact, and equipment deck that develops mana, builds a durable creature presence, and uses its equipped threats to pressure opponents through combat. Card advantage and protective interaction help it stay engaged over longer games, while broad removal and several board resets keep opposing battlefields manageable. It gives up some speed for a permanent-heavy strategy that is strongest once several creatures and artifacts are established.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 2.76 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Belladonna Took | draw | Draw support for the deck’s longer board-building games.
- 1 Caretaker's Talent | draw | Repeatable card advantage that rewards maintaining a board.
- 1 Circuit Mender | draw | An artifact body that also supplies card advantage.
- 1 Cut a Deal | draw | A straightforward refill when the hand is running low.
- 1 Dawn of a New Age | draw | An ongoing source of card advantage.
- 1 Fountainport Bell | draw | Artifact-based card advantage for the deck’s permanent plan.
- 1 Heirloom Epic | draw | A dedicated card-advantage artifact.
- 1 Idol of Oblivion | draw | Efficient card advantage in a token-friendly shell.
- 1 Mentor of the Meek | draw | Turns the deck’s smaller creatures into steady card advantage.
- 1 Skullclamp | draw | Low-cost equipment that provides card advantage.
- 1 The Gaffer | draw | A creature-based source of sustained card advantage.
- 1 Baird, Steward of Argive | interaction | Discourages opposing attacks while the board develops.
- 1 Bofur, Reliable Guardian // Concerted Care | interaction | A flexible protective piece that fits the Dwarf theme.
- 1 Crumb and Get It | interaction | A compact instant-speed protective option.
- 1 Dawn's Truce | interaction | Protects the developed board at a key moment.
- 1 Dwarven Mattock | interaction | A Dwarf-themed utility equipment piece.
- 1 Mithril Coat | interaction | Equipment-based protection for an important permanent.
- 1 Reprieve | interaction | A flexible answer that also keeps cards flowing.
- 1 Swiftfoot Boots | interaction | Protects a key creature while supporting the equipment package.
- 1 Arcane Signet | ramp | Reliable early mana fixing and acceleration.
- 1 Bag End Banquet | ramp | Theme-appropriate artifact acceleration.
- 1 Burnished Hart | ramp | Creature-based mana development that works well in longer games.
- 1 Fellwar Stone | ramp | Efficient two-mana acceleration.
- 1 Mind Stone | ramp | Early acceleration with later utility.
- 1 Mox Amber | ramp | Fast acceleration alongside the commander and legendary permanents.
- 1 Ornithopter of Paradise | ramp | Early creature-based mana fixing.
- 1 Patchwork Banner | ramp | Mana development that complements the creature theme.
- 1 Sol Ring | ramp | Efficient acceleration for deploying the deck’s permanents.
- 1 Wayfarer's Bauble | ramp | Early land-based mana development.
- 1 Banishing Light | removal | A versatile permanent-based answer.
- 1 Fiend Hunter | removal | Creature removal attached to a body.
- 1 Generous Gift | removal | Broad instant-speed removal for troublesome permanents.
- 1 Hanged Executioner | removal | A creature-based removal option with useful flexibility.
- 1 Loran of the Third Path | removal | Artifact and enchantment removal on a relevant creature.
- 1 Skyclave Apparition | removal | Efficient creature-based removal.
- 1 Swords to Plowshares | removal | A low-cost answer to an opposing creature.
- 1 The Black Arrow | removal | Equipment-themed targeted removal.
- 1 Westfold Rider | removal | A creature that contributes a removal option.
- 1 Academy Manufactor | synergy | An artifact creature that strengthens the deck’s artifact-focused board.
- 1 Andúril, Flame of the West | synergy | Legendary equipment that supports the deck’s combat package.
- 1 Blade Splicer | synergy | An artifact-oriented creature that adds to the board plan.
- 1 Carrot Cake | synergy | A low-cost artifact piece that supports the permanent-focused strategy.
- 1 Dáin, Lord of the Iron Hills | synergy | A Dwarf-themed legend for the deck’s central creature package.
- 1 Dwarven Provisioner | synergy | Adds another Dwarf to reinforce the tribal core.
- 1 Iron Hills Blacksmith | synergy | A Dwarf Artificer that ties together the tribal and artifact themes.
- 1 Maskwood Nexus | synergy | Helps unify the creature base around the tribal plan.
- 1 Ori, Keeper of Songs | synergy | A Dwarf legend that deepens the themed creature package.
- 1 Tangle Tumbler | synergy | An artifact-based piece that supports the deck’s board development.
- 1 Andúril, Narsil Reforged | threat | A legendary equipment threat for closing through combat.
- 1 Fíli the Pathfinder | threat | A themed legendary creature that adds pressure to the board.
- 1 Helm of the Host | threat | A high-impact equipment threat for a developed board.
- 1 Jazal Goldmane | threat | A creature threat that rewards committing to the battlefield.
- 1 Karn, the Great Creator | threat | A resilient noncreature threat that demands attention.
- 1 Psychosis Crawler | threat | A payoff threat that pairs naturally with the draw package.
- 1 Serra Redeemer | threat | A substantial creature threat for the midgame.
- 1 Starforged Sword | threat | A combat-focused equipment threat.
- 1 Sting, Bilbo's Sword | threat | A low-cost legendary equipment threat.
- 1 Sunscorch Regent | threat | A large evasive creature for finishing games.
- 1 Sword of the Squeak | threat | A compact equipment threat for the creature board.
- 1 Sword of Vengeance | threat | A versatile equipment threat that improves combat pressure.
- 1 Dusk // Dawn | wipe | A reset that can preserve the deck’s ability to rebuild.
- 1 Martial Coup | wipe | A board reset that also supports rebuilding a battlefield presence.
- 1 Promise of Loyalty | wipe | A selective board reset for breaking stalled boards.
- 26 Plains | land | Provides a dependable white mana base.
- 1 Castle Ardenvale | land | A white land with late-game utility.
- 1 Command Tower | land | Reliable access to the commander’s color identity.
- 1 Evolving Wilds | land | Finds a basic Plains while supporting land consistency.
- 1 Exotic Orchard | land | A flexible source of colored mana.
- 1 Fabled Passage | land | Fetches a basic Plains for mana consistency.
- 1 Minas Tirith | land | A themed white land for the mana base.
- 1 Path of Ancestry | land | A tribal land that supports the Dwarf creature base.
- 1 Terramorphic Expanse | land | Finds a basic Plains while smoothing mana.
- 1 Thriving Heath | land | A white source that helps maintain colored access.
- 1 Uncharted Haven | land | Flexible colored mana for the deck’s spells.

</details>

### 24. a 60-card deck from one set

Format: Modern. Theme: aggro. Pool: any_card. Shortlist: 96 names.

Grade: baseline, score 0.34, model 20260903T210658Z.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $45.36 to buy, $45.36 the whole deck.

**Summary:** This mono-red aggro deck builds pressure through its creature threats and synergy pieces, using removal to clear away resistance while its draw cards help it keep moving into the midgame. It aims to win through sustained creature attacks and a few stronger closing threats, giving up broad utility and defensive flexibility in exchange for a focused, proactive game plan.

The quality model grades this deck at the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. The color sources cover the pips, and that raises the grade.

- [INFO] `curve_summary`: average mana value 3.22 over 36 nonland cards

<details><summary>The deck list</summary>

- 24 Mountain | land | Provides the mono-red mana base.
- 4 Artist's Talent | draw | Fills a primary draw slot while staying in the deck's red plan.
- 2 Sazacap's Brew | draw | Adds draw support to help maintain momentum.
- 2 Brazen Collector | ramp | Provides the requested ramp support on a creature.
- 4 Blooming Blast | removal | Forms the core of the deck's removal suite.
- 2 Conduct Electricity | removal | Completes the removal package with additional direct answers.
- 4 Emberheart Challenger | synergy | A central synergy piece for the aggressive creature plan.
- 4 Hearthborn Battler | synergy | Adds more synergy support for the deck's creature pressure.
- 4 Frilled Sparkshooter | threat | An aggressive threat that helps establish pressure.
- 4 Reptilian Recruiter | threat | Provides a reliable creature threat for the attack plan.
- 4 Teapot Slinger | threat | Adds another set of threats to sustain aggression.
- 2 Dragonhawk, Fate's Tempest | threat | Serves as a higher-impact threat to finish games.

</details>

