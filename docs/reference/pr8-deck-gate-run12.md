# PR-8 deck gate

Run date: 2026-09-02. Card snapshot: 2026-09-02.

Verdict: FAIL. 22 of 24 decks passed every block check, 0 invented names reached the user, and 0 summaries stated a false rule of the game. All three bars are zero tolerance.

2 prompts failed before a deck existed. A gate can not pass with an error.

## Summary

| Measure | Value |
|---|---|
| Prompts | 24 |
| Decks returned | 22 |
| Decks with no block finding | 22 |
| Invented names that reached the user | 0 |
| Decks that needed the repair turn | 11 |
| Summaries judged (F-26) | 22 |
| Summaries that state a rule of the game | 0 |
| Summaries that state a FALSE rule | 0 |
| Judge errors | 0 |
| Errors | 2 |
| Prompt version | 12 |
| Calls | 63 |
| Cost | $2.2397 |
| Time | 2646 seconds |

## Findings by code

A block stops the deck. A warning and a note are reports.

| Code | Count |
|---|---|
| `not_owned` | 26 |
| `curve_summary` | 22 |
| `profile_off_band` | 8 |
| `basics_added` | 3 |
| `outside_requested_set` | 2 |
| `precon_share` | 2 |
| `two_card_combo` | 1 |
| `cards_trimmed` | 1 |

By severity: BLOCK 0. WARN 39. INFO 26. 

## The set filter (PR-17B)

A deck asked for a set holds cards of that set family alone. Three things earn an exception, and each one is marked: a card the reader named (D-381), a mana card the reader allowed from outside (D-382), and a basic land, which no set limit filters (D-378).

| # | Prompt | Sets | In set | Fill | Deck cards outside | Marked |
|---|---|---|---|---|---|---|
| 19 | the Hobbit family, two colours | `hob,hoc` | 170 | 0 | 0 | 0 |
| 20 | the Hobbit family, a delegated commander | `hob,hoc` | 75 | 0 | 0 | 0 |
| 21 | the Hobbit family, mana from outside | `hob,hoc` | 127 | 26 | 10 | 10 |
| 22 | a set family and a card from outside it | `hob,hoc` | 170 | 0 | 1 | 1 |
| 23 | two set families at once | `blb,blc,hob,hoc,pblb` | 201 | 0 | 0 | 0 |
| 24 | a 60-card deck from one set | `blb,blc,pblb` | 95 | 0 | 0 | 0 |

## Decks

### 1. lifegain Commander, any card

Format: Commander. Theme: lifegain. Pool: any_card. Shortlist: 303 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band. Block findings: 0.

Cost: $367.75 to buy, $367.75 the whole deck.

**Summary:** A focused Orzhov lifegain strategy that builds a durable board from incremental life-gain triggers, converts that advantage into growing creatures and opposing life loss, and maintains momentum through resilient card advantage. Efficient fixing and acceleration support a proactive midrange plan, while protection, pinpoint answers, and reset tools keep opposing boards from outpacing the deck's battlefield growth.

- [INFO] `curve_summary`: average mana value 3.25 over 63 nonland cards

<details><summary>The deck list</summary>

- 18 Plains | land | Primary white source for early lifegain creatures and protective spells.
- 9 Swamp | land | Reliable black source for removal and drain effects.
- 1 Command Tower | land | Untapped dual-color fixing.
- 1 Godless Shrine | land | Fetchable untapped dual-color source when needed.
- 1 Caves of Koilos | land | Untapped dual-color fixing.
- 1 Mana Confluence | land | Untapped access to either deck color.
- 1 Vault of Champions | land | Usually untapped multiplayer dual land.
- 1 Shattered Sanctum | land | Early untapped dual-color land.
- 1 Isolated Chapel | land | Dual-color land supported by the basic-heavy mana base.
- 1 Exotic Orchard | land | Multiplayer color fixing without entering tapped.
- 1 Reflecting Pool | land | Efficient fixing alongside abundant colored sources.
- 1 Oasis Gardener | ramp | Early mana acceleration and color fixing.
- 1 Bounty Board | ramp | Low-cost mana development.
- 1 Hot Dog Cart | ramp | Early resource development that supports mana production.
- 1 Altar of the Pantheon | ramp | Colored mana fixing and acceleration.
- 1 Cryptolith Fragment // Aurora of Emrakul | ramp | Reliable artifact mana source.
- 1 Orazca Relic | ramp | Mana rock with late-game card value.
- 1 Pristine Talisman | ramp | Mana acceleration that also triggers lifegain payoffs.
- 1 The Celestus | ramp | Mana fixing with recurring utility.
- 1 Potioner's Trove | ramp | Stores lifegain progress as later mana resources.
- 1 Spoils of Evil | ramp | Efficient burst mana when graveyards are developed.
- 1 Archivist of Oghma | draw | Efficient reactive card advantage with incidental life.
- 1 Dawn of Hope | draw | Turns repeated lifegain into cards and board presence.
- 1 Cosmos Elixir | draw | Reliable recurring card advantage.
- 1 Enduring Innocence | draw | Rewards the deck's inexpensive creature suite with cards.
- 1 Mangara, the Diplomat | draw | Deters opponents while generating cards.
- 1 Markov Purifier | draw | Lifegain-driven card advantage.
- 1 Sigarda's Splendor | draw | Provides ongoing cards while rewarding a high life total.
- 1 The Gaffer | draw | Consistent card draw from life-total advantage.
- 1 Tymna the Weaver | draw | Converts creature combat into flexible card draw.
- 1 Vampiric Rites | draw | Sacrifice outlet that converts creatures into cards and life.
- 1 Well of Lost Dreams | draw | Scales card draw with the deck's frequent life gains.
- 1 Alseid of Life's Bounty | interaction | Protects key creatures or enchantments.
- 1 Bofur, Reliable Guardian // Concerted Care | interaction | Flexible protection for important permanents.
- 1 Courageous Resolve | interaction | Protects the board from targeted disruption.
- 1 Faith's Shield | interaction | Flexible protection that can force through combat.
- 1 Metropolis Reformer | interaction | Shields the player and creatures from harmful effects.
- 1 Restoration Magic | interaction | Versatile defensive response for valuable permanents.
- 1 Rune-Tail, Kitsune Ascendant // Rune-Tail's Essence | interaction | A high-life defensive payoff that protects the creature board.
- 1 Werefox Bodyguard | interaction | Creature-based disruption with useful battlefield presence.
- 1 Ayli, Eternal Pilgrim | removal | Repeatable sacrifice-based removal backed by lifegain.
- 1 Gumdrop Poisoner // Tempt with Treats | removal | Flexible creature removal with a relevant alternate mode.
- 1 Murderous Rider // Swift End | removal | Efficient answer to creatures or planeswalkers.
- 1 Nightmare's Thirst | removal | Cheap creature removal that improves with lifegain.
- 1 Solitude | removal | Premium exile removal with emergency flexibility.
- 1 Umezawa's Jitte | removal | Reusable combat control and creature removal.
- 1 Vein Ripper | removal | A resilient finisher that punishes opposing sacrifices.
- 1 Vona, Butcher of Magan | removal | Repeatable permanent removal paid for with life.
- 1 Witch of the Moors | removal | Lifegain-triggered edict effect that pressures opposing boards.
- 1 Amalia Benavides Aguirre | wipe | Lifegain-enabled creature sweep attached to a cheap threat.
- 1 Fumigate | wipe | Resets creature-heavy boards while restoring life.
- 1 Kaya's Wrath | wipe | Reliable unconditional creature sweeper with lifegain.
- 1 Ajani's Pridemate | synergy | Efficient scalable creature from every life-gain event.
- 1 Angel of Vitality | synergy | Amplifies life gain and becomes a meaningful attacker.
- 1 Cleric Class | synergy | Enhances life gain and provides a long-game payoff.
- 1 Cleric of Life's Bond | synergy | Rewards creature development with life and counters.
- 1 Essence Channeler | synergy | Low-cost life-gain engine that supports counter growth.
- 1 Heliod, Sun-Crowned | synergy | Converts life gain into permanent creature growth.
- 1 Indulging Patrician | synergy | Turns substantial life gain into opponent life loss.
- 1 Lurrus of the Dream-Den | synergy | Replays inexpensive permanents for sustained value.
- 1 Resplendent Angel | synergy | Transforms large life gains into evasive creature pressure.
- 1 Vito, Thorn of the Dusk Rose | synergy | Makes every life-gain trigger pressure opposing life totals.
- 1 Archangel of Thune | threat | Turns repeated lifegain into a team-wide scaling threat.
- 1 Attended Healer | threat | Builds a board while rewarding repeated lifegain.
- 1 Bloodbond Vampire | threat | A scalable attacker that grows from the core game plan.
- 1 Celestine, the Living Saint | threat | Recurring creature advantage and evasive combat pressure.
- 1 Divinity of Pride | threat | Large evasive lifelink threat that rewards a high life total.
- 1 Nykthos Paragon | threat | Converts large life gains into overwhelming creature growth.
- 1 Qala, Ajani's Pridemate | threat | A lifegain-scaled attacker that advances the combat plan.
- 1 Regal Bloodlord | threat | Creates a growing evasive board from life gain.
- 1 Rhox Faithmender | threat | Doubles life gain while presenting a sturdy body.
- 1 Twinblade Paladin | threat | Becomes a potent double-striking attacker after large life gains.
- 1 Valkyrie Harbinger | threat | Produces a stream of large flying threats from high life totals.
- 1 Wurmcoil Engine | threat | Resilient lifelink threat that leaves value behind after removal.

</details>

### 2. aristocrats Commander, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 179 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band. Block findings: 0.

Cost: $0.00 to buy, $191.93 the whole deck.

**Summary:** An Orzhov aristocrats deck built around sacrificing expendable creatures for value, using recursive bodies, death triggers, and artifact support to grind through long multiplayer games. Efficient removal and protective effects preserve its central engine, while resilient creature threats and selective sweepers let it convert attrition into a winning board position.

- [INFO] `curve_summary`: average mana value 2.81 over 63 nonland cards

<details><summary>The deck list</summary>

- 11 Plains | land | Basic white mana for the deck’s core spells.
- 12 Swamp | land | Basic black mana for sacrifice, recursion, and removal.
- 1 Command Tower | land | Reliable two-color commander land.
- 1 City of Brass | land | Flexible untapped color fixing.
- 1 Exotic Orchard | land | Efficient multiplayer color fixing.
- 1 Grand Coliseum | land | Provides either deck color when needed.
- 1 Minas Tirith | land | White source with useful late-game utility.
- 1 Plaza of Heroes | land | Supports the commander and legendary creatures.
- 1 Spire of Industry | land | Colored source supported by the artifact package.
- 1 Fountainport | land | Utility land that converts spare resources into value.
- 1 Takenuma, Abandoned Mire | land | Black source with recursion utility.
- 1 Unclaimed Territory | land | Produces colored mana for the creature-heavy deck.
- 1 Secluded Courtyard | land | Creature-focused color fixing.
- 1 Great Hall of the Citadel | land | Additional white-producing utility land.
- 1 Castle Doom | land | Additional black-producing utility land.
- 1 Arcane Signet | ramp | Efficient color fixing.
- 1 Astral Cornucopia | ramp | Scalable mana production.
- 1 Chromatic Lantern | ramp | Fixes colors while accelerating mana.
- 1 Commander's Sphere | ramp | Color fixing that can become a card.
- 1 Deadly Dispute | ramp | Turns a disposable creature or artifact into mana and cards.
- 1 Fellwar Stone | ramp | Low-cost multiplayer mana rock.
- 1 Lotho, Corrupt Shirriff | ramp | Creates Treasure while advancing a creature plan.
- 1 Sol Ring | ramp | Efficient early mana acceleration.
- 1 Thought Vessel | ramp | Colorless acceleration with hand-size utility.
- 1 Wayfarer's Bauble | ramp | Finds a basic land and permanently fixes mana.
- 1 Ahriman | draw | Creature-based card advantage that can be fed into sacrifice lines.
- 1 Call of the Ring | draw | Steady card selection and card advantage.
- 1 Grave Venerations | draw | Rewards graveyard-focused creature play with cards.
- 1 Idol of Oblivion | draw | Efficient repeatable card draw alongside token production.
- 1 Inspiring Overseer | draw | Provides a body and immediate card advantage.
- 1 Lembas | draw | Early card selection that later replaces itself.
- 1 Mask of Memory | draw | Turns creature combat into filtering and card advantage.
- 1 Night's Whisper | draw | Cheap, dependable card draw.
- 1 Skullclamp | draw | Converts small sacrificed creatures into substantial cards.
- 1 Tome of Legends | draw | Reliable incremental cards around the commander.
- 1 Wall of Omens | draw | Early defense that replaces itself.
- 1 Puresteel Paladin | draw | Draws from the equipment package while adding a creature body.
- 1 Bitter Triumph | removal | Flexible instant-speed creature or planeswalker answer.
- 1 Claim the Precious | removal | Answers a problematic creature cleanly.
- 1 Crib Swap | removal | Exiles a creature and handles recursive threats.
- 1 Dispatch | removal | Low-cost creature interaction.
- 1 Fiend Hunter | removal | Creature-based removal that can participate in sacrifice play.
- 1 Generous Gift | removal | Answers any problematic permanent.
- 1 Get Lost | removal | Efficient answer to creatures and nonland permanents.
- 1 Infernal Grasp | removal | Clean instant-speed creature removal.
- 1 Swords to Plowshares | removal | Premium efficient creature exile.
- 1 Bastion Protector | interaction | Keeps the commander protected while adding a creature.
- 1 Boromir, Warden of the Tower | interaction | Protects the board from opposing free interaction.
- 1 Champion's Helm | interaction | Grants durable protection to the commander.
- 1 Clever Concealment | interaction | Protects the board from sweepers and targeted removal.
- 1 Darksteel Plate | interaction | Makes an important creature difficult to remove.
- 1 Lightning Greaves | interaction | Provides immediate protection and haste.
- 1 Swiftfoot Boots | interaction | Repeatable commander protection.
- 1 Unbreakable Formation | interaction | Safeguards the creature board and can enable an attack.
- 1 Austere Command | wipe | Flexible reset that can preserve selected permanent types.
- 1 Fumigate | wipe | Creature sweeper that stabilizes life total.
- 1 Vanquish the Horde | wipe | Efficient creature-board reset.
- 1 Arcade Cabinet | synergy | Artifact value piece that supports the deck’s disposable-permanent plan.
- 1 Blowfly Infestation | synergy | Turns creature deaths and counters into recurring attrition.
- 1 Buster Sword | synergy | Equipment that helps creatures produce value before sacrifice.
- 1 Foot Chopper | synergy | Equipment-based value piece for the creature plan.
- 1 Gollum the Abandoned | synergy | Benefits from creatures dying and supports graveyard recursion.
- 1 Gollum, Patient Plotter | synergy | Repeatable creature recursion supplies sacrifice material.
- 1 Gríma Wormtongue | synergy | Provides an expendable body with disruptive value.
- 1 Nimble Hobbit | synergy | Low-cost creature that advances the death-focused game plan.
- 1 Phantom Train | synergy | Artifact creature support that works well with sacrifice and recursion.
- 1 Stone of Erech | synergy | Graveyard utility that supports the deck’s attrition plan.
- 1 Angel of Serenity | threat | High-impact creature that disrupts opposing boards and rewards recursion.
- 1 Exemplar of Light | threat | Evasive creature that supplies ongoing board pressure.
- 1 Cirith Ungol Patrol | threat | Creature threat that contributes to a persistent board.
- 1 Frontline Medic | threat | Combat-ready creature that protects an attacking force.
- 1 Gaius van Baelsar | threat | Legendary threat with meaningful board impact.
- 1 Kingpin's Enforcers | threat | Creature pressure that fits an attrition-oriented strategy.
- 1 Massacre Girl, Known Killer | threat | Punishes opposing creature boards while advancing a death-focused plan.
- 1 Palace Jailer | threat | Pressures the table while generating monarch value.
- 1 Rat King, Pale Piper | threat | Recurring black creature threat with graveyard value.
- 1 The Sackville-Bagginses | threat | Legendary creature threat that adds value to the board.
- 1 Witch-king of Angmar | threat | Evasive finisher that punishes attacks and stabilizes the board.

</details>

### 3. artifacts Commander, bracket 4

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 299 names.

ERROR: generate: repair: llm openai/gpt-5.6-terra: budget: context deadline exceeded (attempt 2: llm openai/gpt-5.6-terra: budget: context deadline exceeded)

### 4. dinosaur tribal, bracket 2

Format: Commander. Theme: dinosaurs. Pool: any_card. Shortlist: 269 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band. Block findings: 0.

Cost: $210.05 to buy, $210.05 the whole deck.

**Summary:** A welcoming Naya Dinosaur tribal deck built around accelerating into a stampede of creatures, maintaining a full hand through creature-based value, and using protective tricks to keep its biggest attackers charging. It offers straightforward combat decisions alongside flexible answers and several ways to rebuild after a crowded battlefield is cleared.

- [INFO] `curve_summary`: average mana value 3.02 over 61 nonland cards
- [INFO] `basics_added`: the list was 1 card short, so the builder added 1 basic land

<details><summary>The deck list</summary>

- 10 Forest | land | Basic green source for the deck's creature-heavy mana requirements.
- 6 Mountain | land | Basic red source for Dinosaur threats and damage effects.
- 5 Plains | land | Basic white source for protection and creature support.
- 1 Command Tower | land | Reliable three-color commander land.
- 1 Battlefield Forge | land | Untapped red-white color source.
- 1 Brushland | land | Untapped green-white color source.
- 1 Karplusan Forest | land | Untapped red-green color source.
- 1 Sacred Foundry | land | Red-white dual land.
- 1 Stomping Ground | land | Red-green dual land.
- 1 Temple Garden | land | Green-white dual land.
- 1 Cinder Glade | land | Fetchable red-green dual land.
- 1 Canopy Vista | land | Fetchable green-white dual land.
- 1 Bountiful Promenade | land | Multiplayer green-white dual land.
- 1 Spectator Seating | land | Multiplayer red-white dual land.
- 1 Spire Garden | land | Multiplayer red-green dual land.
- 1 Rootbound Crag | land | Red-green dual land.
- 1 Sunpetal Grove | land | Green-white dual land.
- 1 Clifftop Retreat | land | Red-white dual land.
- 1 Exotic Orchard | land | Flexible multiplayer color fixing.
- 1 Unclaimed Territory | land | Produces any needed color for Dinosaurs.
- 1 Arcane Signet | ramp | Efficient three-color mana fixing.
- 1 Birds of Paradise | ramp | Early color fixing and acceleration.
- 1 Cultivate | ramp | Finds lands and advances mana development.
- 1 Drover of the Mighty | ramp | Dinosaur-friendly mana creature.
- 1 Farseek | ramp | Efficient fixing through dual lands.
- 1 Llanowar Elves | ramp | Early green acceleration.
- 1 Nature's Lore | ramp | Efficiently finds a Forest dual land.
- 1 Rampant Growth | ramp | Reliable basic-land acceleration.
- 1 Thunderherd Migration | ramp | Dinosaur-themed land ramp.
- 1 Huatli, Poet of Unity // Roar of the Fifth People | ramp | Early land development with a Dinosaur-themed late option.
- 1 Beast Whisperer | draw | Rewards casting the deck's many creatures with cards.
- 1 Garruk's Uprising | draw | Draw engine that also supports large creatures.
- 1 Demand Answers | draw | Low-cost card selection.
- 1 Faithless Looting | draw | Efficient filtering for action and lands.
- 1 Curious Altisaur | draw | Dinosaur-based card advantage.
- 1 Idol of Oblivion | draw | Low-cost repeatable card draw support.
- 1 Kutzil, Malamet Exemplar | draw | Creature-based card advantage and protection.
- 1 Ripjaw Raptor | draw | Dinosaur that turns damage into cards.
- 1 Runic Armasaur | draw | Draws cards against common opposing utility abilities.
- 1 Sylvan Library | draw | Powerful early draw selection.
- 1 Esper Sentinel | draw | Efficient early card advantage.
- 1 Boros Charm | interaction | Protects the board or provides a finishing burst.
- 1 Ephemerate | interaction | Protects a creature and reuses an enter-the-battlefield ability.
- 1 Heroic Intervention | interaction | Protects the Dinosaur board from removal.
- 1 Lightning Greaves | interaction | Protects a key Dinosaur and gives it haste.
- 1 Snakeskin Veil | interaction | Low-cost protection for an important creature.
- 1 Swiftfoot Boots | interaction | Repeatable haste and protection.
- 1 Itzquinth, Firstborn of Gishath | removal | Efficient Dinosaur removal spell on a creature.
- 1 Savage Stomp | removal | Cheap Dinosaur-themed creature removal.
- 1 Thrashing Brontodon | removal | Dinosaur that answers artifacts or enchantments.
- 1 Needletooth Raptor | removal | Dinosaur removal attached to an enrage trigger.
- 1 Ravenous Sailback | removal | Removes an artifact or enchantment while adding a body.
- 1 Territorial Allosaurus | removal | Flexible creature removal on a Dinosaur.
- 1 Triumphant Chomp | removal | Efficient Dinosaur-supported removal.
- 1 Amped Raptor | synergy | Low-cost Dinosaur that supports the tribal plan.
- 1 Armored Kincaller | synergy | Early Dinosaur body for tribal development.
- 1 Belligerent Yearling | synergy | Low-cost Dinosaur that grows with the herd.
- 1 Commune with Dinosaurs | synergy | Finds a needed Dinosaur or land early.
- 1 Deathgorge Scavenger | synergy | Dinosaur utility creature with graveyard pressure.
- 1 Dromosaur | synergy | Dinosaur tribal support creature.
- 1 Huatli's Raptor | synergy | Early Dinosaur that supports counter-focused creatures.
- 1 Kinjalli's Caller | synergy | Reduces the cost of larger Dinosaurs.
- 1 Marauding Raptor | synergy | Cost reduction and enrage support for Dinosaurs.
- 1 Nest Robber | synergy | Early Dinosaur that helps establish the board.
- 1 Otepec Huntmaster | synergy | Reduces Dinosaur costs and grants haste.
- 1 Raptor Companion | synergy | Simple early Dinosaur for tribal density.
- 1 Balamb T-Rexaur | threat | A substantial Dinosaur threat for combat pressure.
- 1 Cacophodon | threat | A sizable Dinosaur that rewards the creature-focused plan.
- 1 Cavern Stomper | threat | Large Dinosaur pressure at a manageable cost.
- 1 Charging Tuskodon | threat | Turns combat damage into a threatening attack.
- 1 Gigantosaurus | threat | Efficiently massive Dinosaur attacker.
- 1 Harnessed Snubhorn | threat | Dinosaur threat that benefits from the tribal board.
- 1 Nurturing Bristleback | threat | Resilient Dinosaur threat.
- 1 Palani's Hatcher | threat | Builds a wide Dinosaur board from one card.
- 1 Pantlaza, Sun-Favored | threat | Dinosaur leader that provides recurring value.
- 1 Quartzwood Crasher | threat | Creates enormous trampling Dinosaur tokens after combat.
- 1 Regisaur Alpha | threat | Adds multiple Dinosaur bodies and grants haste.
- 1 Shifting Ceratops | threat | Versatile hasty Dinosaur attacker.
- 1 Forerunner of the Empire | wipe | Tribal payoff that can clear small creatures.
- 1 Raging Swordtooth | wipe | Dinosaur-based board damage that triggers enrage.
- 1 Vanquish the Horde | wipe | Efficient reset button when the board gets out of hand.

</details>

### 5. blink Commander, owned first

Format: Commander. Theme: blink. Pool: owned_first. Shortlist: 168 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $206.32 the whole deck.

**Summary:** This Gilraen deck is a blink-focused white creature deck that develops mana, establishes value creatures, and uses protective pieces to keep its core board intact. It aims to pull ahead by repeatedly leaning on its creature synergies, then closes with Angels, artifact creatures, and equipped attackers. The deck gives up broad color flexibility for a focused permanent-heavy plan with strong creature-based pressure and several ways to reset crowded boards.

- [INFO] `curve_summary`: average mana value 3.06 over 63 nonland cards

<details><summary>The deck list</summary>

- 20 Plains | land | Provides the deck's core white mana base.
- 1 Command Tower | land | Provides reliable colored mana.
- 1 City of Brass | land | Provides flexible colored mana.
- 1 Exotic Orchard | land | Provides flexible colored mana.
- 1 Secluded Courtyard | land | Supports the creature-heavy mana base.
- 1 Unclaimed Territory | land | Supports the creature-heavy mana base.
- 1 Spire of Industry | land | Fits the deck's artifact presence while providing mana.
- 1 Grand Coliseum | land | Provides broad color access when needed.
- 1 Evolving Wilds | land | Finds a basic Plains while improving mana access.
- 1 Fabled Passage | land | Finds a basic Plains while improving mana access.
- 1 Marsh Flats | land | Finds a Plains for consistent mana.
- 1 Terramorphic Expanse | land | Finds a basic Plains for consistent mana.
- 1 Ash Barrens | land | Helps convert into the needed basic land.
- 1 Minas Tirith | land | Adds a white-producing legendary land.
- 1 Plaza of Heroes | land | Supports the deck's legendary cards while providing mana.
- 1 Path of Ancestry | land | Provides colored mana for the creature plan.
- 1 Great Hall of the Citadel | land | Adds another white-oriented mana source.
- 1 Arcane Signet | ramp | Provides efficient mana fixing and acceleration.
- 1 Sol Ring | ramp | Provides early mana acceleration.
- 1 Fellwar Stone | ramp | Adds inexpensive mana acceleration.
- 1 Wayfarer's Bauble | ramp | Helps develop mana through a basic land.
- 1 Sword of the Animist | ramp | Pairs with the creature plan to support mana development.
- 1 Giada, Font of Hope | ramp | Provides creature-based mana support.
- 1 Relic of Legends | ramp | Adds mana acceleration for the legendary creature shell.
- 1 Commander's Sphere | ramp | Provides mana fixing and later utility.
- 1 Thought Vessel | ramp | Adds artifact-based mana acceleration.
- 1 Chromatic Lantern | ramp | Smooths mana while accelerating the deck.
- 1 Champions of Minas Tirith | draw | Supplies card advantage from a creature body.
- 1 Faramir, Field Commander | draw | Adds card advantage to the creature plan.
- 1 Idol of Oblivion | draw | Provides a compact source of card advantage.
- 1 Lembas | draw | Provides card advantage from an artifact permanent.
- 1 Mask of Memory | draw | Turns creature attacks into card selection.
- 1 Mirror of Galadriel | draw | Adds a persistent card-advantage piece.
- 1 Tome of Legends | draw | Provides repeatable card advantage alongside the commander.
- 1 The Vision | draw | Adds card advantage on an artifact creature.
- 1 Instant Ramen | draw | Provides a compact card-advantage option.
- 1 Diary of Dreams | draw | Adds another artifact source of card advantage.
- 1 Skullclamp | draw | Provides efficient card advantage around disposable creatures.
- 1 Bastion Protector | interaction | Helps safeguard the commander and key creatures.
- 1 Boromir, Warden of the Tower | interaction | Provides creature-based disruption.
- 1 Clever Concealment | interaction | Protects the board from opposing interaction.
- 1 Gift of Immortality | interaction | Helps preserve an important creature through removal.
- 1 Lightning Greaves | interaction | Protects a key creature efficiently.
- 1 Champion's Helm | interaction | Provides another layer of protection for legendary creatures.
- 1 Swiftfoot Boots | interaction | Protects a key creature while keeping it active.
- 1 Together Forever | interaction | Helps retain creatures through opposing removal.
- 1 Banishing Light | removal | Provides flexible permanent-based removal.
- 1 Journey to Nowhere | removal | Offers efficient creature removal.
- 1 Generous Gift | removal | Answers a wide range of opposing permanents.
- 1 Get Lost | removal | Provides efficient removal for troublesome permanents.
- 1 Swords to Plowshares | removal | Provides efficient creature removal.
- 1 Stroke of Midnight | removal | Answers a broad range of permanent types.
- 1 Crib Swap | removal | Provides instant-speed creature removal.
- 1 Dispatch | removal | Adds low-cost creature removal.
- 1 Oblation | removal | Answers a problematic permanent at instant speed.
- 1 Flickerwisp | synergy | A central blink-oriented creature for the deck's plan.
- 1 Personify | synergy | Supports the deck's blink-focused game plan.
- 1 Ennis, Debate Moderator | synergy | Adds a creature that supports the deck's focused plan.
- 1 Jocasta, Automaton Avenger | synergy | Adds another synergy-oriented permanent to the board.
- 1 Angel of Condemnation | synergy | Provides a creature that fits the blink plan.
- 1 Angel of Sanctions | synergy | Benefits from the deck's focus on reusable creatures.
- 1 Fiend Hunter | synergy | Fits the deck's plan of revisiting creature value.
- 1 Palace Jailer | synergy | Adds a creature-oriented value piece for the blink shell.
- 1 Wall of Omens | synergy | Provides a low-cost creature for recurring value lines.
- 1 Inspiring Overseer | synergy | Adds another creature that supports repeated value.
- 1 Angel of Serenity | threat | A substantial creature threat for closing games.
- 1 Kataki, War's Wage | threat | Adds a disruptive creature presence to pressure opponents.
- 1 Grim Poppet | threat | Provides a resilient artifact creature threat.
- 1 Westfold Rider | threat | Adds a creature body that contributes to combat pressure.
- 1 Troop of Ponies | threat | Provides another creature threat for the board.
- 1 The Walls of Ba Sing Se | threat | Adds a large artifact creature presence.
- 1 Bronze Guardian | threat | Provides an artifact creature that pressures opponents.
- 1 Frontline Medic | threat | Adds a creature that contributes to board pressure.
- 1 Buster Sword | threat | Turns the deck's creatures into more meaningful attackers.
- 1 Crown of Gondor | threat | Adds an Equipment-based source of combat pressure.
- 1 Meteor Sword | threat | Provides an Equipment threat for creature combat.
- 1 Phoenix Down | threat | Adds a permanent that supports the deck's threat density.
- 1 Austere Command | wipe | Provides a flexible reset against developed boards.
- 1 Fumigate | wipe | Provides a full creature-board reset.
- 1 Vanquish the Horde | wipe | Adds an efficient creature-board reset.

</details>

### 6. Modern tempo, tournament

Format: Modern. Theme: tempo. Pool: any_card. Shortlist: 167 names.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $347.20 to buy, $347.20 the whole deck.

**Summary:** This deck establishes pressure with a compact creature suite, then uses cheap draw, removal, and interaction to keep the opponent from stabilizing. It wins by preserving that pressure while the temporal-spell package helps extend a favorable turn cycle. It gives up broader answers and late-game staying power for a focused, proactive tempo approach.

- [INFO] `curve_summary`: average mana value 2.22 over 36 nonland cards

<details><summary>The deck list</summary>

- 4 Steam Vents | land | Blue-red dual land for the tempo mana base.
- 4 Spirebluff Canal | land | Blue-red land that supports the early tempo plan.
- 4 Riverglide Pathway // Lavaglide Pathway | land | Flexible blue or red source.
- 4 Shivan Reef | land | Reliable blue-red source.
- 5 Island | land | Basic blue source.
- 3 Mountain | land | Basic red source.
- 4 Consider | draw | Efficient card selection for a spell-heavy tempo shell.
- 2 Preordain | draw | Early card selection to find pressure or answers.
- 4 Counterspell | interaction | Broad stack interaction for protecting the tempo lead.
- 2 Spell Pierce | interaction | Low-cost interaction that supports an assertive game plan.
- 4 Lightning Bolt | removal | Efficient removal that clears opposing creatures.
- 4 Into the Flood Maw | removal | Cheap removal for keeping blockers and threats off the table.
- 2 Temporal Mastery | synergy | Supports the deck's temporal-spell package.
- 2 Temporal Trespass | synergy | Supports the deck's temporal-spell package.
- 4 Ledger Shredder | threat | Creature pressure for the tempo plan.
- 4 Ragavan, Nimble Pilferer | threat | Early creature threat that helps establish pressure.
- 4 Faerie Mastermind | threat | Evasive creature threat for maintaining pressure.

</details>

### 7. Modern burn, casual

Format: Modern. Theme: burn. Pool: any_card. Shortlist: 290 names.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $46.36 to buy, $46.36 the whole deck.

**Summary:** This mono-red burn deck uses efficient direct-damage spells alongside Eidolon of the Great Revel and Thermo-Alchemist to keep pressure on the opponent while its draw spells maintain momentum. Hazoret the Fervent, Ashcloud Phoenix, Sunspine Lynx, and Torbran, Thane of Red Fell give it meaningful threats when the game goes longer. It wins by sustaining a fast damage race, trading some flexibility for a focused, highly red-intensive game plan.

- [INFO] `curve_summary`: average mana value 2.89 over 36 nonland cards

<details><summary>The deck list</summary>

- 24 Mountain | land | Provides the consistent red mana base for the deck.
- 2 Chandra, Dressed to Kill | ramp | Fills the ramp slot while staying within the deck's red plan.
- 4 Risk Factor | draw | Provides draw to keep burn resources flowing.
- 2 Browbeat | draw | Adds more draw for a spell-heavy burn plan.
- 4 Lightning Bolt | removal | Efficient direct-damage removal for the burn plan.
- 2 Lightning Strike | removal | Adds reliable direct-damage removal.
- 4 Eidolon of the Great Revel | synergy | Supports the deck's low-cost, damage-focused strategy.
- 4 Thermo-Alchemist | synergy | Rewards the deck for casting its many burn and draw spells.
- 4 Hazoret the Fervent | threat | A resilient threat that gives the deck a strong creature presence.
- 4 Ashcloud Phoenix | threat | Provides a persistent airborne threat.
- 4 Sunspine Lynx | threat | Adds another substantial red threat to pressure opponents.
- 2 Torbran, Thane of Red Fell | threat | A damage-focused threat that fits the deck's burn identity.

</details>

### 8. Modern lifegain, FNM

Format: Modern. Theme: lifegain. Pool: any_card. Shortlist: 301 names.

Cards: 60 main, 0 sideboard. Repair turn: yes, for profile_off_band. Block findings: 0.

Cost: $179.36 to buy, $179.36 the whole deck.

**Summary:** An Orzhov lifegain deck for Modern FNM that establishes steady life gain early, converts it into growing creatures and token pressure, and uses efficient black interaction to clear a path for its attackers. Its threats reward every incremental gain of life, allowing the deck to pivot from stabilizing the board to ending the game quickly.

- [INFO] `curve_summary`: average mana value 2.89 over 36 nonland cards

<details><summary>The deck list</summary>

- 4 Plains | land | Reliable white mana for early lifegain creatures and white-intensive spells.
- 4 Swamp | land | Reliable black mana for the deck’s removal and lifegain payoffs.
- 4 Godless Shrine | land | Untapped access to either color when needed.
- 4 Caves of Koilos | land | Early dual-color fixing without entering tapped.
- 4 Mana Confluence | land | Flexible untapped fixing for both colors.
- 4 Shattered Sanctum | land | Dual-color support that keeps the mana base focused on white and black.
- 2 Altar of the Pantheon | ramp | Provides mana acceleration while contributing incidental life.
- 4 Inspiring Overseer | draw | A lifelinking body that replaces itself.
- 2 Dawn of Hope | draw | Turns repeated life gain into a durable card-advantage engine.
- 4 Murderous Rider // Swift End | removal | Flexible creature that answers problematic creatures or planeswalkers.
- 2 Nightmare's Thirst | removal | Efficient creature removal that improves as life is gained.
- 4 Soul Warden | synergy | Starts the life-gain engine from the first creatures entering the battlefield.
- 4 Ajani's Pridemate | synergy | Rapidly grows into a major attacker from repeated life gain.
- 4 Attended Healer | threat | Creates a stream of lifelinking creature tokens alongside life gain.
- 4 Bloodbond Vampire | threat | A scalable black lifegain payoff that becomes a substantial attacker.
- 4 Twinblade Paladin | threat | Life-gain triggers can turn it into a powerful double-striking finisher.
- 2 Cliffhaven Vampire | threat | Converts each instance of life gain into pressure against the opponent.

</details>

### 9. Standard midrange, FNM

Format: Standard. Theme: midrange. Pool: any_card. Shortlist: 164 names.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $226.42 to buy, $226.42 the whole deck.

**Summary:** This black-green midrange deck develops its mana, trades efficiently with opposing threats, and uses draw engines to stay ahead in longer games. It wins by deploying a steady stream of dangerous creatures after the board has been contained, with its value-focused enchantments helping turn creature-based exchanges into lasting advantage. The deck gives up speed for resilience, so it is strongest when it can settle into a measured, interactive game.

- [INFO] `curve_summary`: average mana value 3.44 over 36 nonland cards

<details><summary>The deck list</summary>

- 8 Forest | land | Basic green mana for the deck’s mana base.
- 8 Swamp | land | Basic black mana for the deck’s mana base.
- 4 Blooming Marsh | land | A black-green land slot that helps cast the deck’s spells.
- 4 Overgrown Tomb | land | A black-green land slot that supports reliable two-color mana.
- 2 Llanowar Elves | ramp | Early acceleration for deploying the midrange game plan ahead of schedule.
- 4 Bitter Triumph | removal | Efficient removal to clear opposing problems for the deck’s threats.
- 2 Nowhere to Run | removal | Additional removal that keeps the board manageable.
- 4 Darkstar Augur | draw | A draw-focused creature that helps keep cards flowing while adding to the board.
- 2 Phyrexian Arena | draw | A dedicated draw engine for sustained midrange resources.
- 4 Insidious Roots | synergy | A central engine that ties the deck’s resource plan to its board development.
- 4 Vampiric Rites | synergy | A supporting engine that turns the deck’s creatures into additional value.
- 4 Vein Ripper | threat | A major creature threat that pressures the opponent once the board is stabilized.
- 4 Vaultborn Tyrant | threat | A high-impact threat that provides a powerful top end.
- 4 Massacre Girl, Known Killer | threat | A creature threat that also supports the deck’s attrition-focused plan.
- 2 Rottenmouth Viper | threat | A flexible creature threat that complements the removal suite.

</details>

### 10. Standard aggro, tournament

Format: Standard. Theme: aggro. Pool: any_card. Shortlist: 294 names.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $57.96 to buy, $57.96 the whole deck.

**Summary:** This red-white aggro deck aims to establish threats quickly, back them with Warleader's Call, and keep attacking while its removal clears the way. Boros Charm and Aven Interrupter provide useful disruption without abandoning the board-focused plan, while Fugitive Codebreaker and Inspiring Overseer help sustain pressure. It gives up broader late-game options in favor of a streamlined, proactive creature strategy.

- [INFO] `curve_summary`: average mana value 3.11 over 36 nonland cards

<details><summary>The deck list</summary>

- 10 Mountain | land | Basic red source for the aggressive red half of the deck.
- 10 Plains | land | Basic white source for the aggressive white half of the deck.
- 4 Sacred Foundry | land | Red-white land that supports both colors without adding a tapped-land slot.
- 4 Fugitive Codebreaker | draw | Low-cost draw option that keeps the deck's pressure supported.
- 2 Inspiring Overseer | draw | White draw creature that helps maintain resources while adding to the board.
- 4 Boros Charm | interaction | Flexible red-white interaction for protecting the attack or disrupting an opponent.
- 2 Aven Interrupter | interaction | Creature-based interaction that can contribute to the battlefield plan.
- 4 Case of the Gateway Express | removal | White removal that clears opposing obstacles for attackers.
- 4 Krenko's Buzzcrusher | removal | Red removal attached to an artifact creature for a proactive deck.
- 4 Warleader's Call | synergy | Red-white synergy piece that reinforces the deck's creature-focused plan.
- 4 Dragonback Lancer | threat | Red-white-friendly creature threat for applying early pressure.
- 4 Frilled Sparkshooter | threat | Red creature threat that advances the aggressive game plan.
- 4 Vanguard Seraph | threat | White creature threat that provides another attacking body.

</details>

### 11. Commander with a locked card

Format: Commander. Theme: sacrifice. Pool: any_card. Shortlist: 296 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $344.06 to buy, $344.06 the whole deck.

**Summary:** Karlov leads a creature-based sacrifice deck that develops mana and card flow through its sacrifice pieces, then uses its threats and board-control cards to take over the table. The deck aims to turn expendable creatures into lasting value while keeping pressure on opposing boards, with multiple wipes available when the battlefield gets out of hand. It gives up a more land-diverse mana base in exchange for dependable access to both colors through basic lands.

- [INFO] `curve_summary`: average mana value 3.10 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Baron Bertram Graywater | draw | Draw support that helps keep the sacrifice plan supplied.
- 1 Corrupted Conviction | draw | Efficient draw support for converting sacrificed resources into cards.
- 1 Disciple of Bolas | draw | Draw support for a creature-focused sacrifice strategy.
- 1 Ecstatic Awakener // Awoken Demon | draw | Draw support that fits the deck's sacrifice theme.
- 1 Relic Vial | draw | Draw support for maintaining cards in hand.
- 1 Shadowheart, Dark Justiciar | draw | Draw support for the creature-heavy plan.
- 1 Smothering Abomination | draw | A draw piece aligned with sacrificing creatures.
- 1 Thraxodemon | draw | Draw support that contributes to the sacrifice shell.
- 1 Vampiric Rites | draw | Repeatable draw support for sacrificed creatures.
- 1 Village Rites | draw | Low-cost draw support when a creature is sacrificed.
- 1 Vampire Gourmand | draw | Draw support that fits the deck's creature theme.
- 1 Cartel Aristocrat | interaction | Interaction that also belongs in a sacrifice-focused creature base.
- 1 Dark Privilege | interaction | Interaction to help protect the board while using sacrifice resources.
- 1 Fanatical Devotion | interaction | Interaction that supports a sacrifice-based defense plan.
- 1 Flare of Fortitude | interaction | Interaction for protecting the deck's board presence.
- 1 Gift of Doom | interaction | Interaction that helps preserve an important permanent.
- 1 Kinzu of the Bleak Coven | interaction | Interaction that fits the deck's creature-focused strategy.
- 1 Knight-Captain of Eos | interaction | Interaction on a creature that can contribute to the board.
- 1 Promise of Tomorrow | interaction | Interaction that supports the deck's long-game resilience.
- 18 Plains | land | Basic white mana source for the deck.
- 18 Swamp | land | Basic black mana source for the deck.
- 1 Sol Ring | ramp | Required mana acceleration that helps deploy the deck's permanents.
- 1 Ashnod's Altar | ramp | Sacrifice-based mana acceleration for the deck's core plan.
- 1 Crowded Crypt | ramp | Mana acceleration that supports a creature-sacrifice strategy.
- 1 Deadly Dispute | ramp | Mana acceleration that works with expendable creatures.
- 1 Fain, the Broker | ramp | A ramp piece that contributes to the sacrifice shell.
- 1 Pawn of Ulamog | ramp | Ramp support for a deck built around creatures leaving the battlefield.
- 1 Phyrexian Altar | ramp | Sacrifice-based ramp that turns creatures into mana.
- 1 Pitiless Plunderer | ramp | Ramp support for repeated creature sacrifices.
- 1 Priest of Forgotten Gods | ramp | Ramp that also advances the sacrifice plan.
- 1 Warren Soultrader | ramp | Ramp support for the deck's creature-based engine.
- 1 Attrition | removal | Removal that uses the deck's sacrifice resources.
- 1 Ayli, Eternal Pilgrim | removal | Removal on a creature that fits the sacrifice strategy.
- 1 Blasting Station | removal | Removal that converts sacrificed creatures into board control.
- 1 Bone Shards | removal | Low-cost removal that fits the sacrifice theme.
- 1 Bone Splinters | removal | Removal that can use an expendable creature as a resource.
- 1 Chittering Witch | removal | Creature-based removal for the deck's board plan.
- 1 Dictate of Erebos | removal | Removal that rewards the deck for sacrificing creatures.
- 1 Grave Pact | removal | Removal support for repeated sacrifice turns.
- 1 Yawgmoth, Thran Physician | removal | Removal on a creature suited to the sacrifice shell.
- 1 Bastion of Remembrance | synergy | Core sacrifice synergy for the deck's game plan.
- 1 Carrion Feeder | synergy | A low-cost sacrifice synergy creature.
- 1 Elas il-Kor, Sadistic Pilgrim | synergy | Sacrifice synergy that supports the deck's central theme.
- 1 Fleshtaker | synergy | Sacrifice synergy in a compact creature package.
- 1 Hidden Stockpile | synergy | Synergy support for recurring sacrifice turns.
- 1 Nantuko Husk | synergy | A sacrifice outlet that advances the deck's plan.
- 1 Pitiless Pontiff | synergy | Creature-based sacrifice synergy.
- 1 Viscera Seer | synergy | A low-cost sacrifice outlet for the deck's engines.
- 1 Woe Strider | synergy | Sacrifice synergy that adds to the creature board.
- 1 Zulaport Cutthroat | synergy | A sacrifice-themed synergy piece for closing games.
- 1 Basri's Lieutenant | threat | A substantial white creature threat for the board.
- 1 Felisa, Fang of Silverquill | threat | A threat that fits the deck's Orzhov creature plan.
- 1 Ghoulcaller Gisa | threat | A creature threat for building a commanding board.
- 1 Liesa, Forgotten Archangel | threat | A resilient top-end threat for the deck.
- 1 Liliana, Heretical Healer // Liliana, Defiant Necromancer | threat | A threat that supports the creature-focused game plan.
- 1 Mondrak, Glory Dominus | threat | A white creature threat for the deck's board-building plan.
- 1 Ratadrabik of Urborg | threat | An Orzhov threat that adds pressure to the board.
- 1 Requiem Angel | threat | A high-impact creature threat for longer games.
- 1 Sidisi, Undead Vizier | threat | A black creature threat that supports the deck's plan.
- 1 Vampire Warlord | threat | A Vampire threat that fits the deck's creature base.
- 1 Vindictive Vampire | threat | A Vampire threat that supports the sacrifice theme.
- 1 Vito's Inquisitor | threat | A creature threat for applying pressure.
- 1 Austere Command | wipe | A flexible wipe for resetting difficult boards.
- 1 The Meathook Massacre | wipe | A wipe that fits the deck's black sacrifice shell.
- 1 Toxic Deluge | wipe | An efficient wipe for clearing the battlefield.

</details>

### 12. Commander on a budget

Format: Commander. Theme: tokens. Pool: any_card. Shortlist: 296 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $85.61 to buy, $85.61 the whole deck.

**Summary:** This Adeline deck develops a broad white creature board, using token-focused permanents and efficient support pieces to keep attacking and rebuilding through resistance. It wins primarily by turning a growing board sideways, backed by creature and planeswalker threats, while its removal and sweepers prevent opposing boards from taking over. The tradeoff is a deliberately budget-conscious mana base and a lower concentration of premium token doublers and top-end finishers.

- [INFO] `curve_summary`: average mana value 3.39 over 62 nonland cards
- [INFO] `basics_added`: the list was 1 card short, so the builder added 1 basic land

<details><summary>The deck list</summary>

- 28 Plains | land | Reliable basic mana for the mono-white token plan.
- 1 Ancient Den | land | An artifact land that supports the mana base.
- 1 Command Tower | land | Reliable commander-deck mana fixing.
- 1 Exotic Orchard | land | Flexible nonbasic mana for the deck.
- 1 Castle Ardenvale | land | A white-producing utility land.
- 1 Eiganjo, Seat of the Empire | land | A white-producing utility land.
- 1 Minas Tirith | land | A white-producing legendary land.
- 1 Secluded Courtyard | land | A tribal-friendly white mana source.
- 1 Path of Ancestry | land | A budget tribal-focused mana source.
- 1 Windbrisk Heights | land | A low-cost utility land for an attacking deck.
- 1 Charisma Bobblehead | ramp | Affordable artifact-based mana support.
- 1 Coin of Mastery | ramp | Budget artifact ramp for developing the board.
- 1 Collector's Vault | ramp | Low-cost artifact mana support.
- 1 Discerning Financier | ramp | Creature-based mana support that fits the deck's color identity.
- 1 Druidic Satchel | ramp | Low-cost artifact ramp support.
- 1 Goldvein Pick | ramp | A cheap equipment-based mana option.
- 1 Karn, Living Legacy | ramp | A durable mana-support permanent.
- 1 Keeper of the Accord | ramp | Creature-based ramp for a board-focused deck.
- 1 Monologue Tax | ramp | Efficient white ramp support.
- 1 Smuggler's Share | ramp | A strong budget-appropriate ramp piece for the deck.
- 1 Bygone Bishop | draw | Creature-based card advantage for the token strategy.
- 1 Dawn of Hope | draw | A low-cost draw engine for a long game.
- 1 Faramir, Field Commander | draw | A creature-based draw option that fits the go-wide plan.
- 1 Idol of Oblivion | draw | An efficient artifact draw piece for token decks.
- 1 Platoon Dispenser | draw | A resilient artifact source of card advantage.
- 1 Search the Premises | draw | An inexpensive enchantment draw engine.
- 1 Staff of the Storyteller | draw | A budget artifact draw option for token play.
- 1 Thorough Investigation | draw | A token-compatible enchantment draw piece.
- 1 Wedding Announcement // Wedding Festivity | draw | A low-cost draw card that suits a creature-heavy board.
- 1 Wojek Investigator | draw | A cheap creature-based draw option.
- 1 Caretaker's Talent | draw | A powerful token-focused source of card advantage.
- 1 Basri Ket | interaction | A low-cost planeswalker interaction piece.
- 1 Lena, Selfless Champion | interaction | Creature-based protection for the developing board.
- 1 Rootborn Defenses | interaction | An inexpensive defensive interaction spell.
- 1 Spirit Bonds | interaction | A flexible enchantment-based interaction piece.
- 1 Teyo, the Shieldmage | interaction | Budget planeswalker interaction for protecting the plan.
- 1 Squad Commander | interaction | A creature interaction option for the token shell.
- 1 Aerial Assault | removal | Cheap white removal.
- 1 Banishing Slash | removal | Low-cost sorcery-speed removal.
- 1 Generous Gift | removal | Flexible instant-speed removal.
- 1 Kellan's Lightblades | removal | An efficient removal spell.
- 1 Righteous Confluence | removal | Versatile removal for varied opposing boards.
- 1 Skyclave Apparition | removal | Creature-based removal that advances the board.
- 1 Stroke of Midnight | removal | Flexible instant-speed removal.
- 1 Hour of Reckoning | wipe | A token-oriented board reset.
- 1 Martial Coup | wipe | A board wipe that remains aligned with the token plan.
- 1 Phyrexian Rebirth | wipe | An inexpensive sweep effect.
- 1 Aligned Heart | synergy | A low-cost synergy piece for the token strategy.
- 1 Anafenza, Unyielding Lineage | synergy | A creature-based token synergy card.
- 1 Automated Assembly Line | synergy | Budget artifact synergy for building a board.
- 1 Cat Collector | synergy | A cheap creature that supports the deck's token theme.
- 1 Clarion Spirit | synergy | An efficient creature synergy piece for going wide.
- 1 Divine Visitation | synergy | A high-impact token synergy enchantment.
- 1 Felidar Retreat | synergy | A flexible enchantment for the token plan.
- 1 Horn of Gondor | synergy | A thematic artifact synergy piece for a wide board.
- 1 Intangible Virtue | synergy | An efficient anthem-style token synergy card.
- 1 Oketra's Monument | synergy | A token-focused artifact synergy piece.
- 1 Rosie Cotton of South Lane | synergy | A low-cost creature that rewards token development.
- 1 Siege Veteran | synergy | A creature synergy piece for a growing board.
- 1 Skrelv's Hive | synergy | An efficient enchantment for the token-focused shell.
- 1 Ajani's Chosen | threat | A budget creature threat for the token strategy.
- 1 Archon of Sun's Grace | threat | An inexpensive evasive threat for the deck.
- 1 Attended Healer | threat | A low-cost creature threat that fits the board-building plan.
- 1 Basri's Lieutenant | threat | A creature threat that supports a counters-and-tokens board.
- 1 Cemetery Protector | threat | A versatile creature threat for combat-focused games.
- 1 Defiler of Faith | threat | A budget creature threat in the deck's color identity.
- 1 Emeria Angel | threat | A low-cost token-oriented creature threat.
- 1 Gideon, Ally of Zendikar | threat | A planeswalker threat suited to a wide-board strategy.
- 1 God-Eternal Oketra | threat | A durable legendary creature threat.
- 1 Hero of Bladehold | threat | An efficient combat threat for an attacking token deck.
- 1 Mite Overseer | threat | A budget creature threat for the token shell.
- 1 Oketra the True | threat | A resilient legendary creature threat.

</details>

### 13. owned first, and the commander is not owned

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 229 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $165.41 the whole deck.

**Summary:** This Karlov deck develops a lifegain-focused board, using its synergy pieces to make the commander and creature suite increasingly threatening. It wins by applying steady combat pressure with Angels and other creatures while protecting key pieces, backed by plentiful targeted answers and a few reset buttons when opposing boards get ahead. The tradeoff is that the deck prioritizes board presence and thematic lifegain support over a more tutor-heavy or fast-mana-driven approach.

- [WARN] `not_owned`: Karlov of the Ghost Council: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.08 over 63 nonland cards

<details><summary>The deck list</summary>

- 17 Plains | land | Provides a reliable white mana base.
- 16 Swamp | land | Provides a reliable black mana base.
- 1 Command Tower | land | Provides flexible commander-color mana.
- 1 City of Brass | land | Adds flexible color access.
- 1 Exotic Orchard | land | Adds flexible color access.
- 1 Sol Ring | ramp | Efficiently accelerates the deck's development.
- 1 Arcane Signet | ramp | Provides dependable mana acceleration.
- 1 Fellwar Stone | ramp | Provides additional mana acceleration.
- 1 Wayfarer's Bauble | ramp | Helps build mana resources early.
- 1 Commander's Sphere | ramp | Provides mana acceleration with later utility.
- 1 Thought Vessel | ramp | Provides artifact-based mana acceleration.
- 1 Sword of the Animist | ramp | Supports mana development while fitting the creature plan.
- 1 Relic of Legends | ramp | Adds flexible mana acceleration.
- 1 Lotho, Corrupt Shirriff | ramp | Provides a ramp option in the deck's colors.
- 1 Giada, Font of Hope | ramp | Accelerates the deck's Angel threats.
- 1 Buster Sword | draw | Provides a draw option that fits the equipment package.
- 1 Call of the Ring | draw | Provides a dedicated draw source.
- 1 Exemplar of Light | draw | Pairs card access with the deck's creature plan.
- 1 Grave Venerations | draw | Provides a dedicated draw source.
- 1 Idol of Oblivion | draw | Offers artifact-based card access.
- 1 Inspiring Overseer | draw | Provides card access on a lifegain-friendly creature.
- 1 Lembas | draw | Provides a compact draw option.
- 1 Mask of Memory | draw | Adds repeatable card access through combat.
- 1 Night's Whisper | draw | Provides efficient card access.
- 1 Skullclamp | draw | Provides equipment-based card access.
- 1 Tome of Legends | draw | Provides ongoing card access around the commander.
- 1 Banishing Light | removal | Provides flexible permanent removal.
- 1 Bitter Triumph | removal | Provides efficient creature removal.
- 1 Crib Swap | removal | Provides instant-speed removal.
- 1 Dismember | removal | Provides efficient removal for problematic creatures.
- 1 Dispatch | removal | Provides low-cost targeted removal.
- 1 Fatal Push | removal | Provides efficient targeted removal.
- 1 Generous Gift | removal | Answers a broad range of permanents.
- 1 Infernal Grasp | removal | Provides unconditional targeted removal.
- 1 Path to Exile | removal | Provides efficient creature removal.
- 1 Bastion Protector | interaction | Helps protect the commander and key legends.
- 1 Boromir, Warden of the Tower | interaction | Provides disruptive interaction on a creature.
- 1 Champion's Helm | interaction | Helps safeguard the commander.
- 1 Clever Concealment | interaction | Provides protection against opposing disruption.
- 1 Darksteel Plate | interaction | Provides durable protection for an important creature.
- 1 Enduring Angel // Angelic Enforcer | interaction | Adds a protective interaction option.
- 1 Lightning Greaves | interaction | Protects a key creature while supporting attacks.
- 1 Swiftfoot Boots | interaction | Provides repeatable creature protection.
- 1 Aerith Gainsborough | synergy | Supports the deck's lifegain-focused synergy plan.
- 1 Angel of Vitality | synergy | Supports the deck's lifegain-focused synergy plan.
- 1 Compassionate Healer | synergy | Reinforces the lifegain theme.
- 1 Kor Firewalker | synergy | Adds a creature that supports the lifegain theme.
- 1 Light of Promise | synergy | Rewards the deck's lifegain focus.
- 1 Night Nurse, Healer of Heroes | synergy | Reinforces the lifegain-focused game plan.
- 1 Prideful Feastling | synergy | Supports the deck's lifegain synergy.
- 1 Rosie Cotton of South Lane | synergy | Rewards the deck's lifegain plan.
- 1 Second Breakfast | synergy | Provides a thematic lifegain synergy piece.
- 1 Wanderbrine Preacher | synergy | Supports the deck's lifegain-focused plan.
- 1 Angel of Invention | threat | Provides an evasive creature threat.
- 1 Dawnhand Eulogist | threat | Adds a meaningful creature threat.
- 1 Gollum, Silent Slinker // Meager Meal | threat | Provides a threatening creature with additional flexibility.
- 1 Invisible Woman, Sue Storm | threat | Adds a legendary threat to pressure opponents.
- 1 Lyra Dawnbringer | threat | Provides a powerful Angel threat.
- 1 Minwu, White Mage | threat | Adds a legendary creature threat.
- 1 Rabaroo Troop | threat | Provides another creature to apply pressure.
- 1 Reaping Willow | threat | Adds a resilient creature threat.
- 1 Rooftop Percher | threat | Provides an additional creature threat.
- 1 Shattered Angel | threat | Adds an evasive threat that fits the theme.
- 1 Sneering Shadewriter | threat | Provides a black creature threat.
- 1 Victory's Herald | threat | Provides an evasive finishing threat.
- 1 Austere Command | wipe | Provides a flexible battlefield reset.
- 1 Fumigate | wipe | Provides a creature-focused battlefield reset.
- 1 Vanquish the Horde | wipe | Provides an efficient creature battlefield reset.

</details>

### 14. the user delegates the commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 244 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $746.22 to buy, $746.22 the whole deck.

**Summary:** Atarka leads a combat-focused Dragon deck that builds mana early, develops a board of Dragons, and turns that board into repeated attacking pressure. It wins primarily through large Dragon attacks, backed by tribal support, removal, and a small protection package to preserve key threats. The deck gives up some early-board speed and broad answers in exchange for a focused, high-impact creature plan.

- [INFO] `curve_summary`: average mana value 3.44 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Cavern of Souls | land | Land slot for the Dragon-focused mana base.
- 1 City of Brass | land | Land slot for the mana base.
- 1 Command Tower | land | Land slot for the mana base.
- 1 Cragcrown Pathway // Timbercrown Pathway | land | Land slot for the mana base.
- 1 Exotic Orchard | land | Land slot for the mana base.
- 9 Forest | land | Basic land slots for the mana base.
- 1 Game Trail | land | Land slot for the mana base.
- 1 Haven of the Spirit Dragon | land | Land slot that supports the Dragon theme.
- 1 Karplusan Forest | land | Land slot for the mana base.
- 1 Mana Confluence | land | Land slot for the mana base.
- 13 Mountain | land | Basic land slots for the mana base.
- 1 Rootbound Crag | land | Land slot for the mana base.
- 1 Rockfall Vale | land | Land slot for the mana base.
- 1 Spire Garden | land | Land slot for the mana base.
- 1 Stomping Ground | land | Land slot for the mana base.
- 1 Taiga | land | Land slot for the mana base.
- 1 Skullclamp | draw | Efficient listed draw piece.
- 1 Sylvan Library | draw | Early listed draw piece.
- 1 Faithless Looting | draw | Low-cost listed draw piece.
- 1 Demand Answers | draw | Listed draw support.
- 1 Thrill of Possibility | draw | Listed draw support.
- 1 Elemental Bond | draw | Draw support for the creature-heavy plan.
- 1 Garruk's Uprising | draw | Draw support for the creature-heavy plan.
- 1 Guardian Project | draw | Ongoing listed draw support.
- 1 Beast Whisperer | draw | Creature-based listed draw support.
- 1 Dragonborn Champion | draw | Dragon-themed listed draw support.
- 1 Harmonize | draw | Reliable listed draw support.
- 1 Lightning Greaves | interaction | Low-cost interaction for key creatures.
- 1 Swiftfoot Boots | interaction | Interaction that supports important creatures.
- 1 Mithril Coat | interaction | Interaction for protecting a key permanent.
- 1 Heroic Intervention | interaction | Broad listed interaction.
- 1 Veil of Summer | interaction | Low-cost listed interaction.
- 1 Snakeskin Veil | interaction | Low-cost listed interaction.
- 1 Tamiyo's Safekeeping | interaction | Low-cost listed interaction.
- 1 Tibalt's Trickery | interaction | Flexible listed interaction.
- 1 Mox Jasper | ramp | Early ramp for deploying the deck's expensive Dragons.
- 1 Carnelian Orb of Dragonkind | ramp | Dragon-themed ramp support.
- 1 Jade Orb of Dragonkind | ramp | Dragon-themed ramp support.
- 1 Orb of Dragonkind | ramp | Dragon-themed ramp support.
- 1 Dragon's Hoard | ramp | Dragon-themed ramp support.
- 1 Dragonstorm Globe | ramp | Ramp support for the mana base.
- 1 Scaled Nurturer | ramp | Creature-based ramp that fits the Dragon theme.
- 1 Ganax, Astral Hunter | ramp | Dragon-themed ramp payoff.
- 1 Savage Ventmaw | ramp | Dragon ramp for larger follow-up plays.
- 1 Klauth, Unrivaled Ancient | ramp | Large Dragon that supports the ramp plan.
- 1 Draconic Roar | removal | Efficient Dragon-themed removal.
- 1 Dragon's Fire | removal | Efficient Dragon-themed removal.
- 1 Molten Exhale | removal | Low-cost listed removal.
- 1 Piercing Exhale | removal | Low-cost listed removal.
- 1 Spit Flame | removal | Dragon-themed removal support.
- 1 Invasion of Tarkir // Defiant Thundermaw | removal | Removal that also fits the Dragon plan.
- 1 Glorybringer | removal | Dragon-based removal option.
- 1 Scourge of Valkas | removal | Dragon-themed removal payoff.
- 1 Terror of the Peaks | removal | Dragon-based removal pressure.
- 1 Dragonlord's Servant | synergy | Cost-support synergy for Dragons.
- 1 Dragonspeaker Shaman | synergy | Cost-support synergy for Dragons.
- 1 Crucible of Fire | synergy | Direct Dragon tribal synergy.
- 1 Dragonkin Berserker | synergy | Low-cost Dragon tribal synergy.
- 1 Sarkhan's Triumph | synergy | Dragon-focused library-search synergy.
- 1 Firespitter Whelp | synergy | Early Dragon synergy piece.
- 1 Kargan Dragonrider | synergy | Low-cost Dragon-related synergy.
- 1 Minion of the Mighty | synergy | Low-cost synergy for the Dragon plan.
- 1 Slumbering Dragon | synergy | Early Dragon synergy piece.
- 1 Shivan Devastator | synergy | Flexible Dragon synergy threat.
- 1 Breath Weapon | wipe | Low-cost board wipe.
- 1 Draconic Intervention | wipe | Dragon-themed board wipe.
- 1 Balefire Dragon | wipe | Dragon-based board wipe option.
- 1 Ancient Bronze Dragon | threat | Large Dragon threat for the combat plan.
- 1 Backdraft Hellkite | threat | Dragon threat that adds pressure in combat.
- 1 Blast-Furnace Hellkite | threat | Large Dragon threat for the combat plan.
- 1 Dragon Broodmother | threat | Dragon threat that develops the board.
- 1 Dragonhawk, Fate's Tempest | threat | Dragon threat for sustained pressure.
- 1 Hellkite Charger | threat | Dragon threat for the attacking plan.
- 1 Lathliss, Dragon Queen | threat | Dragon tribal threat for the main plan.
- 1 Scourge of the Throne | threat | High-impact Dragon combat threat.
- 1 Terror of Mount Velus | threat | Large Dragon threat for closing games.
- 1 Thrakkus the Butcher | threat | Dragon threat that supports aggressive combat.
- 1 Twinflame Tyrant | threat | Dragon threat for closing pressure.
- 1 Utvara Hellkite | threat | Top-end Dragon threat for the combat plan.

</details>

### 15. delegated commander, owned first

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 229 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band. Block findings: 0.

Cost: $0.00 to buy, $143.53 the whole deck.

**Summary:** This Orzhov lifegain deck builds a durable creature board, repeatedly gains life for incremental advantages, and uses equipment to protect and empower its most important attackers. Efficient mana, broad answers, and resilient card advantage support a measured multiplayer game plan that can stabilize early before winning through evasive combat and powerful life-total swings.

- [WARN] `not_owned`: Astarion, the Decadent: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.02 over 62 nonland cards

<details><summary>The deck list</summary>

- 18 Plains | land | Reliable white mana for the lifegain core and white-intensive spells.
- 19 Swamp | land | Reliable black mana for Astarion's ability and black support spells.
- 1 Buster Sword | draw | Equipment-based card advantage that supports creature combat.
- 1 Call of the Ring | draw | Steady card selection and a source of life-loss to leverage lifegain payoffs.
- 1 Exemplar of Light | draw | Lifegain-themed creature that converts the plan into cards.
- 1 Idol of Oblivion | draw | Efficient repeatable card advantage alongside token-producing effects.
- 1 Inspiring Overseer | draw | A flying body that replaces itself while gaining life.
- 1 Lembas | draw | Low-cost card advantage with incidental life gain.
- 1 Mask of Memory | draw | Combat-based filtering and card advantage.
- 1 Night's Whisper | draw | Efficient black card draw.
- 1 Skullclamp | draw | Turns small creatures and tokens into substantial card advantage.
- 1 Tome of Legends | draw | Reliable long-game draw with the commander in play.
- 1 Wall of Omens | draw | Early defense that draws a card.
- 1 Arcane Signet | ramp | Efficient fixing for both colors.
- 1 Bender's Waterskin | ramp | Mana acceleration that fits the artifact package.
- 1 Commander's Sphere | ramp | Color fixing that can become a card later.
- 1 Fellwar Stone | ramp | Low-cost multiplayer mana fixing.
- 1 Inherited Envelope | ramp | Artifact-based mana acceleration.
- 1 Lotho, Corrupt Shirriff | ramp | Generates Treasure while advancing the deck's colors and creature plan.
- 1 Oath of the Grey Host | ramp | Develops mana while contributing to the broader board plan.
- 1 Sol Ring | ramp | Efficient early acceleration.
- 1 Thought Vessel | ramp | Colorless acceleration with a useful hand-size benefit.
- 1 Wayfarer's Bauble | ramp | Early land-based ramp and color fixing.
- 1 Banishing Light | removal | Flexible answer to problematic nonland permanents.
- 1 Bitter Triumph | removal | Efficient answer to creatures or planeswalkers.
- 1 Crib Swap | removal | Exiles a creature and can answer difficult graveyard threats.
- 1 Generous Gift | removal | Versatile instant-speed permanent removal.
- 1 Infernal Grasp | removal | Reliable unconditional creature removal.
- 1 Stroke of Midnight | removal | Flexible answer to troublesome permanents.
- 1 Swords to Plowshares | removal | Premium creature exile that also reinforces the lifegain theme.
- 1 Austere Command | wipe | Flexible reset that can preserve the most useful parts of the board.
- 1 Fumigate | wipe | Creature reset with a meaningful life-gain cushion.
- 1 Vanquish the Horde | wipe | Affordable creature sweeper in multiplayer games.
- 1 Bastion Protector | interaction | Protects the commander while adding a durable body.
- 1 Champion's Helm | interaction | Protects the commander and strengthens combat.
- 1 Clever Concealment | interaction | Safeguards the board from opposing sweepers and removal.
- 1 Darksteel Plate | interaction | Makes a key creature difficult to remove.
- 1 Lightning Greaves | interaction | Provides immediate protection and haste.
- 1 Swiftfoot Boots | interaction | Protects an important creature while retaining flexibility.
- 1 Aettir and Priwen | synergy | Equipment support that rewards attacking with a lifegain-focused creature board.
- 1 Aerith Gainsborough | synergy | A lifegain payoff that helps sustain the creature plan.
- 1 Angel of Vitality | synergy | Turns repeated life gain into a larger evasive threat.
- 1 Compassionate Healer | synergy | Provides repeatable life gain for the deck's payoffs.
- 1 Elixir | synergy | Low-cost artifact support with life-gain utility.
- 1 Kor Firewalker | synergy | Resilient lifegain creature that improves favorable combat positions.
- 1 Light of Promise | synergy | Converts consistent life gain into a major creature threat.
- 1 Night Nurse, Healer of Heroes | synergy | Adds another lifegain engine on a relevant body.
- 1 Prideful Feastling | synergy | Rewards the deck for creating repeated life-gain events.
- 1 Rosie Cotton of South Lane | synergy | Builds a board presence from the deck's steady life gain.
- 1 Second Breakfast | synergy | Efficient lifegain support with useful creature synergies.
- 1 Well-Worn Spatula | synergy | Equipment utility that supports Food and lifegain play patterns.
- 1 White Mage's Staff | synergy | Equipment support that reinforces the deck's healing theme.
- 1 Angel of Invention | threat | Produces a meaningful aerial board presence and pressures life totals.
- 1 Bill the Pony | threat | A resilient creature threat that supports the deck's incremental plan.
- 1 Canyon Crawler | threat | A substantial creature that pressures opponents through combat.
- 1 Dawnhand Eulogist | threat | A scalable creature threat that benefits from the deck's core game plan.
- 1 Frodo, Sauron's Bane | threat | A mana sink that can become a serious late-game threat.
- 1 Gollum, Silent Slinker // Meager Meal | threat | Provides an evasive threat with additional value.
- 1 Lyra Dawnbringer | threat | A powerful flying lifelink finisher that stabilizes races.
- 1 Minwu, White Mage | threat | A lifegain-oriented legend that contributes meaningful battlefield pressure.
- 1 Morlun, Devourer of Spiders | threat | A high-impact black threat for closing games.
- 1 Shattered Angel | threat | Flying pressure that punishes opponents for developing their mana.
- 1 Sneering Shadewriter | threat | An evasive black threat that rewards the deck's attrition plan.
- 1 Victory's Herald | threat | A combat-focused finisher that gives the creature board strong closing power.

</details>

### 16. a tight budget, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 229 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $20.84 to buy, $62.48 the whole deck.

**Summary:** This black-white aristocrats deck develops through a dense creature suite, mana rocks, and steady card advantage, then presses an attrition game with sacrifice-focused synergy and creature threats. It has targeted answers and several board resets to keep opposing boards manageable, while protective interaction helps preserve its key pieces. The deck gives up explosive speed and expensive staples for a focused, library-first build that stays within a modest purchase budget.

- [WARN] `not_owned`: Disciple of Bolas: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bastion of Remembrance: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Blood Artist: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Elas il-Kor, Sadistic Pilgrim: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Falkenrath Noble: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Teysa, Orzhov Scion: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vindictive Vampire: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Zulaport Cutthroat: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Woe Strider: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Yahenni, Undying Partisan: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Shadowheart, Dark Justiciar: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Skullport Merchant: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ayli, Eternal Pilgrim: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Aron, Benalia's Ruin: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Baron Bertram Graywater: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bartolomé del Presidio: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Elder Arthur Maxson: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Embalmed Ascendant: the deck needs 1, the collection has 0
- [WARN] `not_owned`: High-Society Hunter: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Lord Skitter's Butcher: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Lord of the Forsaken: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Old Flitterfang: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ruthless Lawbringer: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Shilgengar, Sire of Famine: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 2.92 over 62 nonland cards

<details><summary>The deck list</summary>

- 13 Plains | land | Basic white source for the two-color mana base.
- 16 Swamp | land | Basic black source for the two-color mana base.
- 1 Command Tower | land | Reliable multicolor land already in the library.
- 1 Exotic Orchard | land | Multicolor land already in the library.
- 1 City of Brass | land | Flexible color source already in the library.
- 1 Grand Coliseum | land | Multicolor land already in the library.
- 1 Spire of Industry | land | Nonbasic color source for the artifact-containing list.
- 1 Plaza of Heroes | land | Nonbasic land already in the library.
- 1 Secluded Courtyard | land | Creature-oriented color source already in the library.
- 1 Unclaimed Territory | land | Creature-oriented color source already in the library.
- 1 Arcane Signet | ramp | Efficient listed mana acceleration.
- 1 Chromatic Lantern | ramp | Color smoothing and ramp already in the library.
- 1 Fellwar Stone | ramp | Low-cost mana acceleration already in the library.
- 1 Sol Ring | ramp | Efficient artifact ramp already in the library.
- 1 Thought Vessel | ramp | Artifact mana acceleration already in the library.
- 1 Commander's Sphere | ramp | Flexible artifact ramp already in the library.
- 1 Deadly Dispute | ramp | Listed ramp that fits the creature-focused plan.
- 1 Bender's Waterskin | ramp | Artifact ramp already in the library.
- 1 Blitzball | ramp | Low-cost artifact ramp already in the library.
- 1 Inherited Envelope | ramp | Artifact ramp already in the library.
- 1 Skullclamp | draw | Creature-focused card advantage already in the library.
- 1 Idol of Oblivion | draw | Artifact card advantage already in the library.
- 1 Mask of Memory | draw | Equipment-based card advantage already in the library.
- 1 Tome of Legends | draw | Repeatable artifact card advantage already in the library.
- 1 Painful Truths | draw | Efficient listed card advantage.
- 1 Wall of Omens | draw | Low-cost creature-based card advantage.
- 1 Inspiring Overseer | draw | Creature-based card advantage already in the library.
- 1 Puresteel Paladin | draw | Equipment-friendly card advantage already in the library.
- 1 Folk Hero | draw | Creature-focused card advantage already in the library.
- 1 Nasty End | draw | Black card advantage already in the library.
- 1 Disciple of Bolas | draw | Creature-based card advantage for the sacrifice-focused plan.
- 1 Swiftfoot Boots | interaction | Protective equipment already in the library.
- 1 Gift of Immortality | interaction | Protective enchantment already in the library.
- 1 Together Forever | interaction | Creature-focused protection already in the library.
- 1 Take Up the Shield | interaction | Low-cost protective interaction already in the library.
- 1 Unbreakable Formation | interaction | Protective interaction for the creature board.
- 1 Bastion Protector | interaction | Creature-based protective interaction.
- 1 Bitter Triumph | removal | Flexible instant-speed listed removal.
- 1 Claim the Precious | removal | Black targeted removal already in the library.
- 1 Crib Swap | removal | Instant-speed creature removal already in the library.
- 1 Deadly Precision | removal | Low-cost listed removal already in the library.
- 1 Destroy Evil | removal | Flexible white removal already in the library.
- 1 Infernal Grasp | removal | Efficient black targeted removal.
- 1 Stroke of Midnight | removal | Flexible white instant removal.
- 1 Austere Command | wipe | Flexible board-reset option already in the library.
- 1 Dusk // Dawn | wipe | Creature-focused board-reset option already in the library.
- 1 Fumigate | wipe | Board-reset option already in the library.
- 1 Bastion of Remembrance | synergy | Core synergy piece for the aristocrats plan.
- 1 Blood Artist | synergy | Core aristocrats synergy piece.
- 1 Elas il-Kor, Sadistic Pilgrim | synergy | Creature-focused aristocrats synergy.
- 1 Falkenrath Noble | synergy | Aristocrats-focused creature synergy.
- 1 Teysa, Orzhov Scion | synergy | Black-white creature synergy for the central plan.
- 1 Vindictive Vampire | synergy | Aristocrats-focused creature synergy.
- 1 Zulaport Cutthroat | synergy | Core creature-based aristocrats synergy.
- 1 Woe Strider | synergy | Creature synergy for the sacrifice-focused plan.
- 1 Yahenni, Undying Partisan | synergy | Creature synergy for the aristocrats shell.
- 1 Gollum, Patient Plotter | synergy | Listed synergy creature already in the library.
- 1 Gollum the Abandoned | synergy | Listed synergy creature already in the library.
- 1 Shadowheart, Dark Justiciar | synergy | Creature synergy for the sacrifice-focused plan.
- 1 Skullport Merchant | synergy | Creature synergy that suits the deck's central plan.
- 1 Ayli, Eternal Pilgrim | threat | Low-cost black-white creature threat.
- 1 Aron, Benalia's Ruin | threat | Low-cost creature threat for the board.
- 1 Baron Bertram Graywater | threat | Creature threat that fits the black-white shell.
- 1 Bartolomé del Presidio | threat | Low-cost black-white creature threat.
- 1 Elder Arthur Maxson | threat | Black-white creature threat for the board.
- 1 Embalmed Ascendant | threat | Creature threat for the white side of the deck.
- 1 High-Society Hunter | threat | Black creature threat for the central plan.
- 1 Lord Skitter's Butcher | threat | Creature threat for the black creature suite.
- 1 Lord of the Forsaken | threat | Black creature threat for the top of the curve.
- 1 Old Flitterfang | threat | Creature threat that fits the black-white shell.
- 1 Ruthless Lawbringer | threat | Low-cost creature threat for the board.
- 1 Shilgengar, Sire of Famine | threat | Black-white creature threat for the deck's upper curve.

</details>

### 17. upgrade a precon, owned first

Format: Commander. Theme: goblins. Pool: owned_first. Shortlist: 196 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for deck_size, precon_share, profile_off_band, profile_off_band, profile_off_band, profile_off_band, profile_off_band, profile_off_band, profile_off_band. Block findings: 0.

Cost: $84.44 to buy, $285.55 the whole deck.

**Summary:** A mono-red Goblin swarm deck built to turn inexpensive creatures into overwhelming combat steps and spell-fueled bursts. It develops a wide board, protects its centerpiece, converts token production into damage, and uses ritual mana plus graveyard recursion to assemble explosive finishing turns.

- [INFO] `curve_summary`: average mana value 2.54 over 63 nonland cards
- [WARN] `precon_share`: the deck keeps 63 of the 78 nonbasic Goblin Storm precon names, and the rule asks for 67
- [WARN] `profile_off_band`: the count of nonbasic lands that make only colorless mana is 5, and bracket 3 wants 4 at most (Fountainport, Goblin Burrows, Kher Keep, Reliquary Tower, War Room)
- [WARN] `profile_off_band`: the removal count is 13, and bracket 3 wants 6 to 12
- [WARN] `profile_off_band`: the wipe count is 1, and bracket 3 wants 2 to 5
- [WARN] `two_card_combo`: bracket 3 allows no two-card infinite combo that needs four mana or less, and the deck holds Storm-Kiln Artist + Haze of Rage (speed 5, near two-card)
- [INFO] `cards_trimmed`: the list was 2 cards over, so the builder cut these cards: Blasphemous Act, Daring Discovery

<details><summary>The deck list</summary>

- 1 Arena of Glory | land | Haste-focused utility land from the precon.
- 1 Castle Embereth | land | Goblin-wide combat boost on a land.
- 1 Den of the Bugbear | land | Creature-producing land that supports the tribal plan.
- 1 Forgotten Cave | land | Cycling land that smooths draws.
- 1 Fountainport | land | Utility land with token synergy.
- 1 Goblin Burrows | land | Tribal combat utility land.
- 1 Hidden Volcano | land | Red utility land from the precon.
- 1 Kher Keep | land | Produces bodies for go-wide turns.
- 1 Reliquary Tower | land | Utility land for retaining a full grip.
- 1 Shinka, the Bloodsoaked Keep | land | Legendary utility land for combat.
- 1 War Room | land | Repeatable card-advantage land.
- 25 Mountain | land | Reliable red mana base for a mono-red spell-and-Goblin strategy.
- 1 Ancestors' Aid | draw | Cantripping combat trick that scales with Zada.
- 1 Battle Hymn | ramp | Explosive ritual mana for large Goblin boards.
- 1 Boggart Shenanigans | removal | Turns Goblin deaths into direct damage.
- 1 Chaos Warp | removal | Flexible answer to troublesome permanents.
- 1 Conspicuous Snoop | synergy | Goblin-based card access and creature synergy.
- 1 Crimson Wisps | draw | Cheap cantrip that grants haste and multiplies with Zada.
- 1 Dragon Fodder | synergy | Creates multiple Goblins for go-wide and targeting turns.
- 1 Empty the Warrens | wincon | Major token payoff for spell-heavy turns.
- 1 Expedite | draw | One-mana cantrip and haste enabler for Zada turns.
- 1 Faithless Looting | draw | Efficient filtering that stocks the graveyard.
- 1 Fists of Flame | draw | Cantripping Zada payoff that turns a board into lethal damage.
- 1 Frontline Heroism | synergy | Combat-oriented precon support card.
- 1 Gempalm Incinerator | removal | Goblin-scaled creature removal.
- 1 General Kreat, the Boltbringer | synergy | Goblin legend that supports aggressive token pressure.
- 1 Goblin Bushwhacker | synergy | Haste and power boost for token-finishing turns.
- 1 Goblin Chieftain | synergy | Lord effect and haste for Goblin attacks.
- 1 Goblin Dark-Dwellers | threat | Creature threat that recasts a useful spell.
- 1 Grapeshot | removal | Storm-based removal and finishing reach.
- 1 Haze of Rage | wincon | Repeatable storm combat finisher with Zada potential.
- 1 Idol of Oblivion | draw | Reliable token-enabled card draw.
- 1 Impact Tremors | wincon | Converts every Goblin token wave into damage.
- 1 Krenko's Command | synergy | Efficient pair of Goblin tokens.
- 1 Mana Geyser | ramp | Large ritual that enables explosive main phases.
- 1 Mogg War Marshal | synergy | Multiple Goblin bodies from one card.
- 1 Past in Flames | wincon | Graveyard recursion for a decisive spell chain.
- 1 Quest for the Goblin Lord | synergy | Tribal payoff for the Goblin swarm.
- 1 Roaming Throne | synergy | Doubles valuable Goblin triggered abilities.
- 1 Ruby Medallion | ramp | Red spell cost reduction supports storm turns.
- 1 Sazacap's Brew | draw | Cheap filtering spell for Zada and storm turns.
- 1 Seething Song | ramp | Fast ritual mana for high-tempo turns.
- 1 Siege-Gang Commander | removal | Creates a board and converts Goblins into damage.
- 1 Siege-Gang Lieutenant | removal | Goblin token maker with direct-damage utility.
- 1 Skirk Prospector | ramp | Converts Goblins into mana during combo-like turns.
- 1 Skullclamp | draw | Exceptional card flow with disposable Goblin tokens.
- 1 Sol Ring | ramp | Efficient early mana acceleration.
- 1 Spreading Insurrection | wincon | High-impact precon finisher.
- 1 Storm-Kiln Artist | ramp | Builds Treasure mana from the deck's instant and sorcery density.
- 1 Swiftfoot Boots | interaction | Protects Zada while enabling immediate activation-free attacks.
- 1 Throne of Eldraine | ramp | Mana-producing artifact support from the precon.
- 1 Vandalblast | wipe | Artifact sweep that can preserve the deck's own board.
- 1 Warren Torchmaster | synergy | Goblin combat payoff.
- 1 Wild Ride | synergy | Precon combat spell that supports aggressive pressure.
- 1 Witch's Mark | draw | Low-cost filtering and targeting spell for Zada.
- 1 Arcane Signet | ramp | Consistent additional mana acceleration.
- 1 Lightning Greaves | interaction | Protects Zada and gives it haste.
- 1 Darksteel Plate | interaction | Durable protection for the commander or a key threat.
- 1 Champion's Helm | interaction | Protects the commander while improving combat.
- 1 Abrade | removal | Flexible, efficient answer to creatures or artifacts.
- 1 Goblin Bombardment | removal | Sacrifice outlet that turns tokens into repeatable damage.
- 1 Broadside Bombardiers | removal | Goblin-based removal that rewards expendable bodies.
- 1 Goblin Matron | synergy | Finds the most useful Goblin for the current board state.
- 1 Goblin Warchief | synergy | Cost reduction and haste improve explosive Goblin turns.
- 1 Rundvelt Hordemaster | synergy | Tribal lord and card advantage for Goblin attacks.
- 1 Goblin Lackey | synergy | Early Goblin deployment accelerator.
- 1 Goblin Trashmaster | removal | Goblin lord that also answers artifacts.
- 1 Pashalik Mons | removal | Token payoff that turns Goblin sacrifices into damage.
- 1 Krenko, Mob Boss | threat | Powerful repeatable Goblin-token engine.
- 1 Moggcatcher | threat | Tutors large Goblin threats directly into play.
- 1 Sword of the Animist | ramp | Combat-based mana development on a creature-heavy board.
- 1 Lightning Bolt | removal | Efficient removal with direct player-targeting flexibility.
- 1 Goblin Fireleaper | removal | Goblin creature that provides additional removal reach.
- 1 Goblin Glasswright // Craft with Pride | ramp | Goblin body and mana-oriented support for explosive turns.
- 1 Boneclub Berserker | threat | Additional Goblin threat for pressure and tribal density.

</details>

### 18. upgrade a precon, any card

Format: Commander. Theme: turtles. Pool: any_card. Shortlist: 300 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band, profile_off_band, profile_off_band, profile_off_band, profile_off_band, profile_off_band. Block findings: 0.

Cost: $152.49 to buy, $152.49 the whole deck.

**Summary:** A five-color Turtle-themed creature deck that develops its mana, assembles a resilient team of heroes and allies, and uses a broad suite of answers to keep opponents from disrupting its board. The strategy leans on character-driven synergies, steady card advantage, and decisive combat turns backed by protective interaction.

- [INFO] `curve_summary`: average mana value 3.77 over 62 nonland cards
- [WARN] `precon_share`: the deck keeps 66 of the 88 nonbasic Turtle Power precon names, and the rule asks for 75
- [WARN] `profile_off_band`: the average mana value of the nonland cards is 3.77, and bracket 3 wants 2 to 3.5
- [WARN] `profile_off_band`: the mana available on turn four is 4.11, and bracket 3 wants 4.2 or more (mean over 10000 hands)
- [INFO] `basics_added`: the list was 1 card short, so the builder added 1 basic land

<details><summary>The deck list</summary>

- 1 Ash Barrens | land | Fixes colors while preserving a preconstructed land.
- 1 Big Apple, 3 a.m. | land | Thematic multicolor land.
- 1 Cinder Glade | land | Color source from the preconstructed mana base.
- 1 Command Tower | land | Reliable five-color fixing.
- 1 Escape Tunnel | land | Finds a needed basic color.
- 1 Evolving Wilds | land | Flexible basic-land fixing.
- 1 Fabled Passage | land | Efficient fixing for the multicolor deck.
- 1 Grand Coliseum | land | Produces every needed color.
- 1 Hidden Hideout | land | Thematic mana source.
- 1 Path of Ancestry | land | Tribal fixing and selection.
- 1 Rain-Slicked Copse | land | Preconstructed dual land.
- 1 Rootbound Crag | land | Red-green fixing.
- 1 Sodden Verdure | land | Preconstructed blue-green source.
- 1 Thriving Grove | land | Flexible color fixing.
- 1 Thriving Isle | land | Flexible color fixing.
- 1 Thriving Moor | land | Flexible color fixing.
- 1 Turtle Lair | land | Thematic tribal land.
- 1 Undergrowth Stadium | land | Black-green fixing.
- 1 Vernal Fen | land | Preconstructed dual land.
- 1 Vibrant Cityscape | land | Five-color fixing.
- 1 Blood Crypt | land | Fast black-red fixing.
- 1 Breeding Pool | land | Fast green-blue fixing.
- 1 Canopy Vista | land | Green-white fixing.
- 1 City of Brass | land | Unconditional five-color mana.
- 1 Exotic Orchard | land | Usually provides excellent multiplayer fixing.
- 1 Godless Shrine | land | Fast white-black fixing.
- 1 Hallowed Fountain | land | Fast white-blue fixing.
- 1 Mana Confluence | land | Unconditional five-color mana.
- 1 Morphic Pool | land | Blue-black fixing.
- 1 Overgrown Tomb | land | Fast black-green fixing.
- 1 Rejuvenating Springs | land | Green-blue fixing in multiplayer.
- 1 Sacred Foundry | land | Fast red-white fixing.
- 1 Steam Vents | land | Fast blue-red fixing.
- 1 Temple Garden | land | Fast green-white fixing.
- 1 Watery Grave | land | Fast blue-black fixing.
- 1 Yavimaya Coast | land | Green-blue fixing.
- 1 Ancient Adamantoise | ramp | A Turtle that advances mana development.
- 1 Blossoming Tortoise | ramp | Thematic creature-based acceleration.
- 1 Fecund Greenshell | ramp | Turtle ramp that supports the creature plan.
- 1 Michelangelo, Improviser | ramp | Thematic mana acceleration.
- 1 Arcane Signet | ramp | Efficient multicolor acceleration.
- 1 Chromatic Lantern | ramp | Fixes colors while accelerating.
- 1 Sol Ring | ramp | Efficient early acceleration.
- 1 Tokka & Rahzar, Unsupervised | ramp | Thematic ramp creature.
- 1 Ambling Stormshell | draw | Turtle-based card advantage.
- 1 Donatello, Turtle Techie | draw | Thematic repeatable card advantage.
- 1 Donnie & April, Adorkable Duo | draw | Supports the team while drawing cards.
- 1 Mikey & Leo, Chaos & Order | draw | Thematic source of card advantage.
- 1 The Pride of Hull Clade | draw | Rewards the deck's large creatures with cards.
- 1 Thunderous Snapper | draw | Turtle card advantage.
- 1 Venus, Torn Between Worlds | draw | Thematic value engine.
- 1 April O'Neil, Live on the Scene | draw | Preconstructed card-advantage piece.
- 1 Acidic Slime | removal | Flexible permanent destruction on a creature.
- 1 Assassin's Trophy | removal | Answers any problematic permanent.
- 1 Steelbane Hydra | removal | Thematic artifact and enchantment removal.
- 1 Colossal Skyturtle | removal | Turtle removal with useful flexibility.
- 1 Kappa Tech-Wrecker | removal | Thematic artifact and enchantment answer.
- 1 Metalhead | removal | Thematic removal attached to a creature.
- 1 Aetherize | wipe | Punishes an opposing combat step and resets attackers.
- 1 Blasphemous Act | wipe | Efficient creature-board reset.
- 1 Bedrock Tortoise | interaction | A Turtle that protects the game plan.
- 1 Saved by the Shell | interaction | Thematic protection for key creatures.
- 1 Swiftfoot Boots | interaction | Protects a key creature while granting haste.
- 1 Vintara Snapper | interaction | Thematic interactive creature.
- 1 Arcade Cabinet | synergy | Thematic preconstructed artifact.
- 1 Baxter, Fly in the Ointment | other | Preconstructed thematic character.
- 1 Bebop, Skull & Crossbones | other | Preconstructed thematic character.
- 1 Big Mother Mouser | other | Preconstructed thematic artifact creature.
- 1 Biogenic Ooze | threat | Preconstructed creature threat.
- 1 Casey Jones, Back Alley Brute | other | Preconstructed thematic character.
- 1 Coin of Mastery | synergy | Preconstructed thematic artifact.
- 1 Continue? | other | Preconstructed thematic instant.
- 1 Corpsejack Menace | synergy | Supports counter-based creature development.
- 1 Dimension X Pizzasaur | other | Preconstructed thematic artifact creature.
- 1 Double Jump // Flying Kick | other | Preconstructed thematic combat trick.
- 1 Electric Seaweed | other | Preconstructed defensive creature.
- 1 Endless Foot Assault | synergy | Preconstructed thematic enchantment.
- 1 Everything Pizza | synergy | Preconstructed thematic artifact.
- 1 Exploding Barrel | other | Preconstructed thematic artifact.
- 1 Fast Forward | other | Preconstructed thematic spell.
- 1 Foot Chopper | synergy | Preconstructed thematic Equipment.
- 1 Game Over | wincon | Preconstructed thematic finisher.
- 1 Harmonize | other | Preconstructed value spell.
- 1 Here Comes a New Hero! | synergy | Preconstructed thematic spell.
- 1 High Score | synergy | Preconstructed thematic enchantment.
- 1 Irma, Part-Time Mutant | other | Preconstructed thematic character.
- 1 Krang, the All-Powerful | threat | Preconstructed thematic artifact creature.
- 1 Leatherhead, Iron Gator | threat | Preconstructed thematic character.
- 1 Lessons from Life | other | Preconstructed thematic spell.
- 1 Level Up | synergy | Preconstructed thematic Aura.
- 1 Lita, Little Orphan Amphibian | synergy | Supports the Turtle-focused strategy.
- 1 Mole Module | other | Preconstructed thematic Vehicle.
- 1 Mona Lisa, Science Geek | other | Preconstructed thematic character.
- 1 Ninja Pizza | synergy | Preconstructed thematic enchantment.
- 1 Rat King, Pale Piper | threat | Preconstructed thematic character.
- 1 Ray Fillet, Wave Warrior | threat | Preconstructed thematic character.
- 1 Roadkill Rodney | other | Preconstructed thematic artifact creature.
- 1 Rocksteady, Mutant Marauder | threat | Preconstructed thematic character.
- 1 Forest | land | the builder added this basic land to reach the deck size

</details>

### 19. the Hobbit family, two colours

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 174 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band. Block findings: 0.

Cost: $308.70 to buy, $308.70 the whole deck.

**Summary:** Thranduil leads a Sultai Elf host that develops quickly through mana creatures, builds a resilient board of woodland allies, and turns that board into sustained combat pressure. Blue card advantage and disruptive magic keep the hand full and opponents off balance, while black removal clears the most dangerous obstacles. Tribal-friendly sweepers allow the host to recover from crowded battlefields and maintain momentum into the late game.

- [INFO] `curve_summary`: average mana value 3.43 over 63 nonland cards
- [WARN] `profile_off_band`: the worst color's share of the sources it needs is 0.63, and bracket 3 wants 0.85 or more (sources of requirement: U 20.75 of 19, B 12 of 19, G 15.25 of 20)

<details><summary>The deck list</summary>

- 1 Beorn the Fierce | draw | Creature-based card advantage that remains a meaningful board presence.
- 1 Bilbo Baggins, Burglar // Take a Glance | draw | Flexible early play that converts into card selection and advantage.
- 1 Bilbo, Luckwearer // Burglar's Plot | draw | Adventure utility provides a compact source of cards.
- 1 Elvish Visionary | draw | An Elf that replaces itself and supports creature synergies.
- 1 Fateful Discovery | draw | Reliable enchantment-based card advantage.
- 1 Gandalf, Shadow's Foe | draw | A high-impact source of cards for the midgame.
- 1 Hithlain Knots | draw | Instant-speed card advantage.
- 1 Ithilien Kingfisher | draw | Evasive creature that contributes cards over time.
- 1 Key to the Side-Door | draw | Artifact card advantage that is difficult for creature removal to answer.
- 1 Lórien Revealed | draw | Flexible early mana smoothing or later card advantage.
- 1 Night's Whisper | draw | Efficient, low-cost card draw.
- 1 Bilbo's Ring | interaction | Protective utility that can disrupt opposing plans.
- 1 Confusticate and Bebother | interaction | Instant-speed disruption for problematic spells or abilities.
- 1 Elrond, Moon-Reader | interaction | Elf-themed interactive legendary creature.
- 1 Mithril Coat | interaction | Protects an important creature from removal.
- 1 My Precious // Allure of Power | interaction | Versatile interaction with a useful Adventure option.
- 1 Old Fat Spider Can't See Me | interaction | A Saga that interferes with opposing development.
- 1 Stern Scolding | interaction | Efficient answer to early opposing creatures.
- 1 Thranduil's Decree | interaction | On-theme instant-speed disruption.
- 1 Bilbo's Deadly Slice | removal | Efficient creature removal.
- 1 Bitter Downfall | removal | Broad answer to a troublesome permanent.
- 1 Colossal Whale | removal | Large threat that repeatedly removes opposing creatures.
- 1 Enchanted River's Grasp | removal | Blue removal that answers a key creature or permanent.
- 1 Giant's Boulder | removal | Repeatable artifact-based removal.
- 1 Merciless Executioner | removal | Creature-based sacrifice removal.
- 1 Orcish Bowmasters | removal | Efficient removal attached to a disruptive creature.
- 1 Quarrel | removal | Low-cost removal for opposing threats.
- 1 Stir Up Trouble | removal | Flexible removal spell.
- 1 Arcane Signet | ramp | Fixes all of the deck's colors while accelerating development.
- 1 Delighted Halfling | ramp | Early green acceleration for legendary and creature-heavy lines.
- 1 Elven Chorus | ramp | Elf-focused mana production and creature support.
- 1 Elvish Archdruid | ramp | Powerful Elf mana engine that also rewards a wide board.
- 1 Elvish Mystic | ramp | Reliable one-mana green acceleration.
- 1 Silvan Reveler | ramp | Elf body that advances mana development.
- 1 Thranduil the Strategist | ramp | Thematic legendary Elf acceleration.
- 1 Thranduil's Company | ramp | Elf-based mana development that reinforces the tribal plan.
- 1 Wayfarer's Bauble | ramp | Color-stable basic-land ramp.
- 1 Wood Elves | ramp | Finds a Forest while adding another Elf to the board.
- 1 Arwen, Weaver of Hope | synergy | Strengthens the Elf board and makes the creature plan more resilient.
- 1 Boughside Wanderers | synergy | Elf creature that benefits from the deck's tribal support.
- 1 Cantankerous Keepers | synergy | Contributes another Elf body for tribal payoffs.
- 1 Celeborn the Wise | synergy | Legendary Elf payoff for the deck's creature-focused strategy.
- 1 Galadhrim Guide | synergy | Cheap Elf that helps establish the tribal board.
- 1 Galion, Elvenking's Butler | synergy | Thematic Elf legend supporting Thranduil's court.
- 1 Mirkwood Meditator | synergy | Elf creature that advances the deck's cohesive board plan.
- 1 Mirkwood Nurturer | synergy | On-theme Elf support creature.
- 1 Nimrodel Watcher | synergy | Elf body that benefits from the deck's tribal density.
- 1 Supper for Spiders | synergy | Tribal-support spell that rewards the creature-focused strategy.
- 1 Attercop | threat | Large, flavorful creature that pressures opponents.
- 1 Dreaded Bat-Cloud | threat | Evasive threat that helps close games through stalled boards.
- 1 Gigantic Big Bear | threat | Large creature that provides combat pressure.
- 1 Great Fierce Bee | threat | Creature threat that adds board presence.
- 1 Large Bear | threat | Efficient combat body for the creature plan.
- 1 Little Bear | threat | Early creature that contributes to a growing board.
- 1 Mirkwood Elk | threat | Solid green creature threat.
- 1 Nasty Little Rabbit | threat | Low-cost creature that helps build pressure early.
- 1 Ordinary Bear | threat | Straightforward creature pressure.
- 1 Ravenhill Flock | threat | Flying threat that can attack around ground defenses.
- 1 Troll of Khazad-dûm | threat | Substantial black creature for the top of the curve.
- 1 Willow-Wind | threat | Creature threat that complements the deck's green core.
- 1 Gnashing of Teeth | wipe | Resets opposing creature boards when combat stalls.
- 1 Languish | wipe | Efficient sweeper against creature-heavy opponents.
- 1 Raise the Palisade | wipe | Tribal-friendly mass bounce that preserves the Elf board.
- 17 Island | land | Provides abundant blue mana for the deck's blue spells.
- 10 Forest | land | Supports the Elf core and green ramp package.
- 9 Swamp | land | Supplies black mana for removal and utility spells.

</details>

### 20. the Hobbit family, a delegated commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 77 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band. Block findings: 0.

Cost: $469.88 to buy, $469.88 the whole deck.

**Summary:** A mono-red Hobbit-set Dragon deck centered on building treasure-fueled momentum, equipping its legendary threats, and turning combat into overwhelming pressure. It combines artifact value, direct damage, and dramatic battlefield-clearing spells to keep opponents contained while its largest creatures close the game.

- [INFO] `curve_summary`: average mana value 3.18 over 61 nonland cards
- [WARN] `profile_off_band`: the draw count is 6, and bracket 3 wants 8 to 14

<details><summary>The deck list</summary>

- 34 Mountain | land | Reliable red mana base.
- 1 Dragon-Cursed Halls | land | Thematic colored land for the Dragon deck.
- 1 Hobbit Hole | land | Utility land slot.
- 1 The Lonely Mountain | land | Thematic red-producing land.
- 1 Treasure Vault | land | Artifact-land utility and late-game mana sink.
- 1 Arcane Signet | ramp | Efficient color fixing and acceleration.
- 1 Bag End Banquet | ramp | Treasure-focused mana acceleration.
- 1 Burn, Burn, Tree and Fern | ramp | Saga-based mana development.
- 1 Cavern-Hoard Dragon | ramp | A Dragon that helps fund larger plays.
- 1 Dragon's Desire | ramp | Dragon-themed mana acceleration.
- 1 Fíli and Kíli, Joyous | ramp | Creature-based ramp with Hobbit flavor.
- 1 Long-Bodied Grey Dog | ramp | Early mana development on a creature.
- 1 Mox Amber | ramp | Low-cost legendary mana acceleration.
- 1 Orcrist, Goblin-cleaver | ramp | Equipment that contributes to mana development.
- 1 The Reaver Cleaver | ramp | Turns combat damage into a major mana resource.
- 1 Balin, Loremaster | draw | Creature-based card advantage.
- 1 Key to the Side-Door | draw | Artifact card selection and advantage.
- 1 Palantír of Orthanc | draw | Persistent card advantage engine.
- 1 Ragged Short Spear | draw | Equipment-based card advantage.
- 1 Thrór's Map | draw | Artifact source of cards and utility.
- 1 Óin the Brave | draw | Legendary creature that supplies cards.
- 1 Bilbo's Ring | interaction | Protective utility for important creatures.
- 1 Dwarven Mattock | interaction | Flexible equipment-based disruption.
- 1 Mithril Coat | interaction | Protection for the commander or a key threat.
- 1 The One Ring | interaction | Protection and resilient value utility.
- 1 Battle-Scarred Goblin | removal | Creature-based targeted removal.
- 1 Fire of Orthanc | removal | Direct damage removal.
- 1 Gandalf, Spark Starter | removal | Legendary source of repeatable removal pressure.
- 1 Giant's Boulder | removal | Artifact-based removal option.
- 1 Goblin Cratermaker | removal | Flexible answer to creatures and artifacts.
- 1 Improvised Club | removal | Efficient instant-speed removal.
- 1 Inferno Titan | removal | Large threat with built-in damage removal.
- 1 Pinecone Strike | removal | Instant-speed targeted answer.
- 1 Smaug, the Great Calamity // Spew Flame | removal | Dragon threat with removal attached.
- 1 Call Forth the Tempest | wipe | Broad battlefield reset.
- 1 Desolation of Smaug | wipe | Thematic mass-removal spell.
- 1 Glóin the Mighty // Easy Pickings | wipe | Creature that provides a sweeping adventure.
- 1 Andúril, Flame of the West | synergy | Combat equipment supporting the legendary Dragon plan.
- 1 Andúril, Narsil Reforged | synergy | A thematic equipment payoff for combat.
- 1 Glamdring | synergy | Equipment that strengthens combat-focused threats.
- 1 Last Light of Durin's Day | synergy | Enchanting value piece for the deck's legendary theme.
- 1 Long-Lost Lances | synergy | Combat equipment supporting attacking threats.
- 1 Smaug's Fury | synergy | Thematic burst of aggressive pressure.
- 1 Sting, Bilbo's Sword | synergy | Low-cost legendary equipment support.
- 1 The Black Arrow | synergy | Thematic equipment with combat utility.
- 1 Tidings of War | synergy | Supports the deck's attack-oriented game plan.
- 1 Well-Worn Spatula | synergy | Additional equipment for combat value.
- 1 Desert Were-Worm | threat | Large Dragon-adjacent finisher.
- 1 Dori, Bearer of Friends | threat | Legendary creature that adds board presence.
- 1 Dwarven Mauler | threat | Combat-focused creature threat.
- 1 Dwarven Warriors | threat | Creature body that develops the board.
- 1 Dáin Ironfoot | threat | Legendary threat with tribal combat relevance.
- 1 Gandalf, Goblins' Bane // Flameshape | threat | Versatile legendary attacker with an adventure option.
- 1 Goblin-town Flunkies | threat | Early creature pressure.
- 1 Gundabad Opportunist | threat | Aggressive creature that advances the board.
- 1 Guttersnipe | threat | Threat that converts spells into damage.
- 1 Oliphaunt | threat | Large creature for finishing combat.
- 1 Olog-hai Crusher | threat | Heavy combat threat.
- 1 Orcish Siegemaster | threat | Creature threat with an aggressive role.
- 1 Bombur, Gentle Dreamer | other | Legendary utility creature.
- 1 Bothersome Noisemaker | other | Flavorful utility creature.
- 1 Getaway Barrel | other | Artifact utility and resilience.
- 1 Iron Hills Stalwart | other | Additional creature presence.
- 1 Old Thrush | other | Low-cost utility creature.
- 1 Snowslope Hunter | other | Creature utility and board development.
- 1 Troop of Ponies | other | Creature-based support piece.

</details>

### 21. the Hobbit family, mana from outside

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 156 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band, profile_off_band. Block findings: 0.

Cost: $833.07 to buy, $833.07 the whole deck.

**Summary:** A Rakdos Middle-earth treasure-and-equipment deck built to accelerate into imposing Dragons, Giants, Trolls, and legendary villains. It combines a broad removal suite with durable card advantage, then turns accumulated mana and combat enhancements into overwhelming late-game pressure.

- [INFO] `curve_summary`: average mana value 3.22 over 63 nonland cards
- [WARN] `outside_requested_set`: the sets you named do not hold 10 cards: Ancient Tomb, Blood Crypt, Bloodstained Mire, Command Tower, Dragonskull Summit, Marsh Flats, Polluted Delta, Scalding Tarn, Urborg, Tomb of Yawgmoth, Wooded Foothills

<details><summary>The deck list</summary>

- 18 Mountain | land | Reliable red mana for the deck’s primary color and high-impact Dragon spells.
- 8 Swamp | land | Reliable black mana for removal and supporting spells.
- 1 Ancient Tomb | land | Provides a substantial burst of mana toward the deck’s expensive threats.
- 1 Blood Crypt | land | Untapped dual-color source when needed.
- 1 Bloodstained Mire | land | Finds either basic color or the dual land while improving color access.
- 1 Command Tower | land | Consistent access to both commander colors.
- 1 Dragonskull Summit | land | Dual-color land that is commonly untapped.
- 1 Marsh Flats | land | Fetches black sources and can find the dual land.
- 1 Polluted Delta | land | Fetches black sources and can find the dual land.
- 1 Scalding Tarn | land | Fetches red sources and can find the dual land.
- 1 Urborg, Tomb of Yawgmoth | land | Improves black access while allowing utility lands to produce mana.
- 1 Wooded Foothills | land | Fetches red sources and can find the dual land.
- 1 Arcane Signet | ramp | Efficient color fixing and early mana acceleration.
- 1 Bag End Banquet | ramp | Treasure-oriented acceleration that supports the deck’s large spells.
- 1 Bolg's Company | ramp | Creature-based mana development that fits the Middle-earth creature suite.
- 1 Burn, Burn, Tree and Fern | ramp | Develops mana while advancing the deck’s saga-adjacent flavor.
- 1 Cavern-Hoard Dragon | ramp | A Dragon that converts combat and treasure resources into further mana.
- 1 Dragon's Desire | ramp | Treasure acceleration tailored to casting costly Dragons.
- 1 Mox Amber | ramp | A low-cost mana source supported by the deck’s legendary creatures.
- 1 Smaug the Magnificent | ramp | Dragon-themed mana acceleration for the top end.
- 1 Smaug, Wicked Worm | ramp | Supports the Dragon plan while producing additional mana resources.
- 1 Wayfarer's Bauble | ramp | Early land-based ramp that improves both mana count and color access.
- 1 Balin, Loremaster | draw | Legendary creature that supplies sustained card advantage.
- 1 Gollum, Riddle Master | draw | A thematic creature that turns the deck’s game actions into cards.
- 1 Key to the Side-Door | draw | Artifact-based card advantage that works alongside the equipment package.
- 1 Night's Whisper | draw | Efficient, inexpensive card selection.
- 1 Palantír of Orthanc | draw | Long-term card advantage from a resilient artifact.
- 1 Rage into the Valley | draw | Red card flow that keeps pressure and resources coming.
- 1 Ragged Short Spear | draw | Equipment that contributes to card advantage while supporting creatures.
- 1 Reverent Howl | draw | Instant-speed card draw for maintaining resources.
- 1 The Master of Lake-town | draw | A legendary source of repeatable card advantage.
- 1 The Sackville-Bagginses | draw | A thematic legendary draw engine.
- 1 Thrór's Map | draw | Artifact card advantage with useful utility.
- 1 Bilbo's Ring | interaction | Protective equipment that helps a key creature survive opposing answers.
- 1 Dwarven Mattock | interaction | Flexible equipment utility for creature combat.
- 1 Getaway Barrel | interaction | Provides flexible protection and tactical creature utility.
- 1 Mithril Coat | interaction | Protects an important legendary creature or threat.
- 1 My Precious // Allure of Power | interaction | Versatile equipment and spell utility.
- 1 Smaug's Fury | interaction | Instant-speed red utility that can influence combat and opposing boards.
- 1 Sting, Bilbo's Sword | interaction | A cheap legendary equipment piece with useful creature utility.
- 1 The One Ring | interaction | A powerful legendary artifact that provides protection and ongoing value.
- 1 Azog, Moria's Ruin | removal | Legendary creature-based answer that contributes to board presence.
- 1 Bilbo's Deadly Slice | removal | Efficient targeted creature removal.
- 1 Bitter Downfall | removal | Broad instant-speed answer to problematic permanents.
- 1 Fire of Orthanc | removal | Red removal for creatures or other troublesome targets.
- 1 Goblin Cratermaker | removal | Flexible creature-based removal for small creatures and artifacts.
- 1 Improvised Club | removal | Low-cost interactive removal spell.
- 1 Orcish Bowmasters | removal | Efficient creature removal attached to a threatening body.
- 1 Pinecone Strike | removal | Instant-speed targeted answer.
- 1 Smite the Deathless | removal | Clean removal for problematic creatures.
- 1 Bolg, Erebor's Reckoning | wipe | Creature-based reset that leaves room to rebuild with threats.
- 1 Call Forth the Tempest | wipe | Sweeper for recovering from developed opposing boards.
- 1 Desolation of Smaug | wipe | On-theme mass removal that clears the way for large attackers.
- 1 Desert Were-Worm | threat | Large creature that applies meaningful combat pressure.
- 1 Dreaded Bat-Cloud | threat | Evasive creature that helps pressure life totals.
- 1 Great Fierce Bee | threat | Creature threat that adds to board pressure.
- 1 Haunt of the Dead Marshes | threat | A resilient thematic threat for the midgame.
- 1 Head of the Hunt | threat | Creature pressure that supports attacking boards.
- 1 Inferno Titan | threat | High-impact finisher that pressures creatures and players.
- 1 Nighthowl Pursuer | threat | Aggressive creature threat with useful combat presence.
- 1 Olog-hai Crusher | threat | Large Troll threat suited to closing games.
- 1 Oliphaunt | threat | A sizable creature that contributes to the deck’s top-end pressure.
- 1 Sauron, the Lidless Eye | threat | Legendary late-game threat with strong battlefield presence.
- 1 The Great Goblin | threat | A thematic legendary threat that supports the Goblin contingent.
- 1 Troll of Khazad-dûm | threat | A substantial threat that remains useful across stages of the game.
- 1 Along the Crooked Way | synergy | Thematic permanent that supports the deck’s value-oriented game plan.
- 1 Andúril, Flame of the West | synergy | Legendary equipment supporting combat-focused creature turns.
- 1 Andúril, Narsil Reforged | synergy | A second legendary equipment payoff for the creature suite.
- 1 Down, Down to Goblin-town | synergy | Saga that reinforces the Goblin and Middle-earth theme.
- 1 Goblin Plate Mail | synergy | Equipment that rewards committing creature threats to combat.
- 1 Guttersnipe | synergy | Converts the deck’s instants and sorceries into additional pressure.
- 1 Last Light of Durin's Day | synergy | Thematic enchantment value that supports the deck’s creature plan.
- 1 Long-Lost Lances | synergy | Equipment that enhances attackers and combat pressure.
- 1 Supper for Spiders | synergy | Thematic support spell for the deck’s creature-centered strategy.
- 1 Well-Worn Spatula | synergy | Low-cost equipment that supports equipped attackers.

</details>

### 22. a set family and a card from outside it

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 175 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band. Block findings: 0.

Cost: $461.42 to buy, $461.42 the whole deck.

**Summary:** A Sultai Hobbit-set Elf deck that develops a broad woodland court, accelerates through creature-based mana, and uses steady card advantage to keep pressure on the table. Its interactive suite protects key legends, disrupts opposing development, and clears the way for a decisive creature-led finish.

- [INFO] `curve_summary`: average mana value 3.33 over 63 nonland cards
- [WARN] `profile_off_band`: the worst color's share of the sources it needs is 0.68, and bracket 3 wants 0.85 or more (sources of requirement: U 13 of 19, B 16.25 of 19, G 18.5 of 19)
- [WARN] `outside_requested_set`: the sets you named do not hold 1 card: Sol Ring

<details><summary>The deck list</summary>

- 11 Forest | land | Green basics support the Elf-heavy core and early mana development.
- 10 Swamp | land | Black basics support removal and the deck's darker Hobbit-set spells.
- 7 Island | land | Blue basics provide reliable mana for card advantage and reactive spells.
- 1 Elven Passage | land | A thematic colored land that supports the deck's Elf strategy.
- 1 Elvenking's Halls | land | A thematic colored land for the Sultai mana base.
- 1 Minas Morgul, Dark Fortress | land | A black-producing utility land for the removal suite.
- 1 Mirkwood | land | A thematic land that reinforces the deck's green and black needs.
- 1 Rivendell | land | A blue-producing legendary land for the deck's spell support.
- 1 The Black Gate | land | A black-producing thematic land.
- 1 Rogue's Passage | land | Utility land that helps an important attacker connect.
- 1 Treasure Vault | land | Flexible artifact utility land without straining colored requirements.
- 1 Elvish Visionary | draw | An inexpensive Elf that replaces itself and supports creature synergies.
- 1 Fateful Discovery | draw | A thematic source of sustained card advantage.
- 1 Hithlain Knots | draw | Efficient card advantage for maintaining resources.
- 1 Ithilien Kingfisher | draw | A creature-based draw piece that contributes to board presence.
- 1 Key to the Side-Door | draw | Repeatable thematic card access.
- 1 Lórien Revealed | draw | Flexible card selection that can also smooth early land drops.
- 1 Night's Whisper | draw | Low-cost, reliable card advantage.
- 1 Palantír of Orthanc | draw | A durable source of ongoing card advantage.
- 1 Plunder the Trollshaws | draw | Thematic instant-speed card advantage.
- 1 Uncover the Moon-Letters | draw | An enduring draw engine for longer games.
- 1 Last March of the Ents | draw | A powerful refill that rewards developing creatures.
- 1 Sol Ring | ramp | Fast, reliable acceleration retained as requested.
- 1 Arcane Signet | ramp | Efficient fixing across the commander's colors.
- 1 Delighted Halfling | ramp | Early green acceleration on a thematic creature.
- 1 Elven Chorus | ramp | Converts the creature-heavy board into additional mana.
- 1 Elvish Archdruid | ramp | Major Elf-based mana production.
- 1 Elvish Mystic | ramp | Efficient early green acceleration.
- 1 Silvan Reveler | ramp | Elf-based mana development for the midgame.
- 1 Thranduil the Strategist | ramp | Thematic acceleration that supports the Elven plan.
- 1 Thranduil's Company | ramp | Creature-based ramp that fits the deck's tribe.
- 1 Wood Elves | ramp | Permanent land-based acceleration and color support.
- 1 Bitter Downfall | removal | Flexible black creature removal.
- 1 Colossal Whale | removal | A large threat that also answers opposing creatures.
- 1 Enchanted River's Grasp | removal | Blue interaction that removes a problematic permanent.
- 1 Merciless Executioner | removal | Creature-based sacrifice removal.
- 1 Orcish Bowmasters | removal | Efficient creature control with incidental pressure.
- 1 Quarrel | removal | Low-cost thematic removal.
- 1 The Black Arrow | removal | A flavorful repeatable removal tool.
- 1 Uneasy Partings | removal | Versatile instant-speed answer.
- 1 Witch-king of Angmar | removal | A substantial threat that controls opposing creatures.
- 1 Gnashing of Teeth | wipe | A board reset against creature swarms.
- 1 Languish | wipe | Efficient mass creature control.
- 1 Raise the Palisade | wipe | A tribal-friendly sweeper that preserves the chosen creature type.
- 1 Confusticate and Bebother | interaction | Reactive protection and disruption.
- 1 Elrond, Moon-Reader | interaction | A thematic interactive legend.
- 1 Mithril Coat | interaction | Protects a key legendary creature from removal.
- 1 My Precious // Allure of Power | interaction | Flexible protection and tactical interaction.
- 1 Old Fat Spider Can't See Me | interaction | A thematic disruption piece.
- 1 Stern Scolding | interaction | Efficient stack interaction against cheap threats.
- 1 The One Ring | interaction | Protection and resilient utility for pivotal turns.
- 1 Thranduil's Decree | interaction | Theme-forward instant-speed interaction.
- 1 Celeborn the Wise | threat | A legendary Elf that provides a meaningful battlefield presence.
- 1 Haunt of the Dead Marshes | threat | A threatening creature aligned with the deck's colors.
- 1 Mirkwood Elk | threat | A thematic green body for combat pressure.
- 1 Mirkwood Nurturer | threat | An Elf creature that advances the board.
- 1 Mirkwood Pathmaker | threat | A thematic creature that contributes to pressure.
- 1 Elven Raft-Steerer | threat | An Elf body that helps establish the board.
- 1 Elvenking's Harper | threat | A thematic Elf that adds battlefield presence.
- 1 Thranduil, Sindarin Liege // Silvan Rally | threat | A major Elven payoff with flexible spell utility.
- 1 Gollum the Abandoned | threat | A resilient thematic threat.
- 1 Dreaded Bat-Cloud | threat | An evasive black creature for pressure.
- 1 Attercop | threat | A large thematic creature that demands an answer.
- 1 Willow-Wind | threat | A substantial green creature for closing games.
- 1 Supper for Spiders | synergy | Supports the deck's creature-focused thematic package.
- 1 Boughside Wanderers | synergy | An Elf that strengthens the deck's tribal density.
- 1 Cantankerous Keepers | synergy | A thematic Elf creature for the core board plan.
- 1 Galadhrim Guide | synergy | An Elf that reinforces the tribal strategy.
- 1 Galion, Elvenking's Butler | synergy | A thematic legendary Elf supporting the Elven court.
- 1 Guardian of the Halls | synergy | An Elf Soldier that contributes to the tribal battlefield.
- 1 Lothlórien Lookout | synergy | A thematic Elf that builds creature synergies.
- 1 Mirkwood Meditator | synergy | An Elf body supporting the deck's unified creature plan.
- 1 Nimrodel Watcher | synergy | A thematic Elf that improves tribal consistency.
- 1 Grey Havens Navigator | synergy | An Elf creature that rounds out the thematic synergy package.

</details>

### 23. two set families at once

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 203 names.

ERROR: generate: generate: llm openai/gpt-5.6-terra: budget: context deadline exceeded (attempt 2: llm openai/gpt-5.6-terra: budget: context deadline exceeded)

### 24. a 60-card deck from one set

Format: Modern. Theme: aggro. Pool: any_card. Shortlist: 96 names.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $59.11 to buy, $59.11 the whole deck.

**Summary:** This mono-red Bloomburrow aggro deck leads with a dense creature threat base, using Emberheart Challenger and Hearthborn Battler to support its aggressive core. It keeps pressure on through a steady mix of threats, backs them with a broad removal package, and uses draw plus a small amount of ramp to maintain momentum. The deck gives up flexibility against specialized opponents in exchange for a focused, straightforward proactive plan.

- [INFO] `curve_summary`: average mana value 3.44 over 36 nonland cards

<details><summary>The deck list</summary>

- 24 Mountain | land | Basic mono-red land base.
- 4 Artist's Talent | draw | Primary source of the deck's draw.
- 2 Sazacap's Brew | draw | Completes the deck's draw package.
- 2 Season of the Bold | ramp | Provides the requested ramp support.
- 1 Agate Assault | removal | Flexible removal slot.
- 1 Blooming Blast | removal | Removal coverage for opposing permanents.
- 1 Conduct Electricity | removal | Removal coverage in the main deck.
- 1 Flame Lash | removal | Additional direct removal.
- 1 Rabid Gnaw | removal | Removal option that fits the Bloomburrow creature focus.
- 1 Take Out the Trash | removal | Sixth removal piece for the deck.
- 4 Emberheart Challenger | synergy | Core aggressive synergy creature.
- 4 Hearthborn Battler | synergy | Complements the deck's aggressive synergy plan.
- 4 Dragonhawk, Fate's Tempest | threat | Top-end threat for closing games.
- 4 Frilled Sparkshooter | threat | Creature threat that supports board pressure.
- 4 Reptilian Recruiter | threat | Reliable creature threat for the aggro plan.
- 2 Teapot Slinger | threat | Additional threat to round out the attacking creature base.

</details>

