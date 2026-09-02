# PR-8 deck gate

Run date: 2026-09-02. Card snapshot: 2026-09-01.

Verdict: PASS. 24 of 24 decks passed every block check, 0 invented names reached the user, and 0 summaries stated a false rule of the game. All three bars are zero tolerance.

## Summary

| Measure | Value |
|---|---|
| Prompts | 24 |
| Decks returned | 24 |
| Decks with no block finding | 24 |
| Invented names that reached the user | 0 |
| Decks that needed the repair turn | 3 |
| Summaries judged (F-26) | 24 |
| Summaries that state a rule of the game | 0 |
| Summaries that state a FALSE rule | 0 |
| Judge errors | 0 |
| Errors | 0 |
| Prompt version | 11 |
| Calls | 51 |
| Cost | $1.4596 |
| Time | 1307 seconds |

## Findings by code

A block stops the deck. A warning and a note are reports.

| Code | Count |
|---|---|
| `curve_summary` | 24 |
| `bracket_prose_rules` | 17 |
| `not_owned` | 15 |
| `land_count` | 2 |
| `outside_requested_set` | 2 |
| `cards_trimmed` | 1 |

By severity: BLOCK 0. WARN 19. INFO 42. 

## The set filter (PR-17B)

A deck asked for a set holds cards of that set family alone. Three things earn an exception, and each one is marked: a card the reader named (D-381), a mana card the reader allowed from outside (D-382), and a basic land, which no set limit filters (D-378).

| # | Prompt | Sets | In set | Fill | Deck cards outside | Marked |
|---|---|---|---|---|---|---|
| 19 | the Hobbit family, two colours | `hob,hoc` | 170 | 0 | 0 | 0 |
| 20 | the Hobbit family, a delegated commander | `hob,hoc` | 75 | 0 | 0 | 0 |
| 21 | the Hobbit family, mana from outside | `hob,hoc` | 127 | 26 | 4 | 4 |
| 22 | a set family and a card from outside it | `hob,hoc` | 170 | 0 | 1 | 1 |
| 23 | two set families at once | `blb,blc,hob,hoc,pblb` | 201 | 0 | 0 | 0 |
| 24 | a 60-card deck from one set | `blb,blc,pblb` | 95 | 0 | 0 | 0 |

## Decks

### 1. lifegain Commander, any card

Format: Commander. Theme: lifegain. Pool: any_card. Shortlist: 303 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $347.88 to buy, $347.88 the whole deck.

**Summary:** This deck develops a white-black lifegain board around Karlov of the Ghost Council, using themed creatures, enchantments, and artifacts to turn life gain into growing threats and sustained pressure. Card advantage and mana pieces help it keep deploying its engine, while targeted answers and board wipes prevent opposing boards from taking over. It wins by building a reinforced lifegain battlefield and leveraging its larger payoff creatures, but gives up multicolor utility lands and alternate dedicated win-condition packages for a stable two-color core.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 3.43 over 63 nonland cards

<details><summary>The deck list</summary>

- 18 Plains | land | Provides a consistent white mana base for the deck.
- 18 Swamp | land | Provides a consistent black mana base for the deck.
- 1 Archivist of Oghma | draw | Adds a dedicated draw option.
- 1 Dawn of Hope | draw | Supports the deck with repeatable draw.
- 1 Cosmos Elixir | draw | Adds a resilient draw piece.
- 1 Disciple of Bolas | draw | Converts a creature into card advantage.
- 1 Enduring Innocence | draw | Supplies another draw engine.
- 1 Exemplar of Light | draw | Contributes card advantage alongside the creature plan.
- 1 Mangara, the Diplomat | draw | Provides steady card advantage.
- 1 Markov Purifier | draw | Adds draw within the Vampire portion of the deck.
- 1 Sigarda's Splendor | draw | Provides a draw payoff for the deck's plan.
- 1 Well of Lost Dreams | draw | Turns the lifegain plan into card advantage.
- 1 Altar of the Pantheon | ramp | Provides mana acceleration.
- 1 Angel of Indemnity | ramp | Adds ramp on a lifegain-themed creature.
- 1 Bounty Board | ramp | Provides additional mana acceleration.
- 1 Colossal Plow | ramp | Adds an artifact-based ramp option.
- 1 Crypt Ghast | ramp | Provides a high-impact black mana engine.
- 1 Hierophant's Chalice | ramp | Adds mana development that fits the deck's theme.
- 1 Nuka-Cola Vending Machine | ramp | Supplies artifact-based acceleration.
- 1 Orazca Relic | ramp | Adds a flexible mana artifact.
- 1 Pristine Talisman | ramp | Provides reliable mana acceleration.
- 1 The Celestus | ramp | Adds a versatile ramp artifact.
- 1 Alseid of Life's Bounty | interaction | Protects key pieces of the lifegain board.
- 1 Faith's Shield | interaction | Offers protective interaction for important permanents.
- 1 Restoration Magic | interaction | Provides a flexible interaction slot.
- 1 Sword of Light and Shadow | interaction | Adds equipment-based interaction to support creatures.
- 1 Sword of War and Peace | interaction | Provides another equipment-based interaction tool.
- 1 Werefox Bodyguard | interaction | Adds creature-based interaction.
- 1 Aetherflux Reservoir | removal | Gives the lifegain shell a decisive removal outlet.
- 1 Ayli, Eternal Pilgrim | removal | Provides repeatable removal within the deck's colors.
- 1 Cavalier of Night | removal | Adds a creature-based removal option.
- 1 Murderous Rider // Swift End | removal | Provides flexible targeted removal.
- 1 Nightmare's Thirst | removal | Adds efficient targeted removal.
- 1 Solitude | removal | Provides creature-based removal.
- 1 Vona, Butcher of Magan | removal | Adds a lifegain-adjacent removal threat.
- 1 Witch of the Moors | removal | Provides removal tied to the deck's main plan.
- 1 Ajani, Strength of the Pride | wipe | Supplies a board-clearing option that fits the theme.
- 1 Fumigate | wipe | Provides a reliable full-board reset.
- 1 Kaya's Wrath | wipe | Adds a second dependable board wipe.
- 1 Ajani's Pridemate | synergy | Rewards the deck for gaining life.
- 1 Angel of Vitality | synergy | Supports the deck's lifegain synergies.
- 1 Blood Artist | synergy | Adds a black synergy piece for the creature plan.
- 1 Bloodthirsty Aerialist | synergy | Rewards repeated lifegain with a growing creature.
- 1 Cleric Class | synergy | Strengthens the deck's central lifegain theme.
- 1 Cleric of Life's Bond | synergy | Adds a Cleric-based lifegain payoff.
- 1 Heliod, Sun-Crowned | synergy | Provides a central lifegain synergy engine.
- 1 Indulging Patrician | synergy | Supports the deck's lifegain-focused creature package.
- 1 Leyline of Hope | synergy | Reinforces the lifegain plan from an enchantment slot.
- 1 Resplendent Angel | synergy | Adds a powerful lifegain payoff creature.
- 1 Righteous Valkyrie | synergy | Supports the Angel and lifegain portions of the deck.
- 1 Serra Ascendant | synergy | Rewards maintaining a strong life total.
- 1 Speaker of the Heavens | synergy | Adds another creature that benefits from the lifegain plan.
- 1 Vito, Thorn of the Dusk Rose | synergy | Turns the deck's lifegain theme into pressure.
- 1 Archangel of Thune | threat | Provides a premier lifegain-focused threat.
- 1 Astarion, the Decadent | threat | Adds a powerful black threat to close games.
- 1 Attended Healer | threat | Builds a board while supporting the deck's theme.
- 1 Celestine, the Living Saint | threat | Provides a resilient lifegain-themed threat.
- 1 Cliffhaven Vampire | threat | Adds pressure aligned with repeated lifegain.
- 1 Defiant Bloodlord | threat | Converts the lifegain strategy into a major threat.
- 1 Divinity of Pride | threat | Provides a large threat for a high-life strategy.
- 1 Gideon's Company | threat | Rewards lifegain with a growing battlefield threat.
- 1 Nykthos Paragon | threat | Turns lifegain into a substantial creature payoff.
- 1 Rhox Faithmender | threat | Provides a high-impact lifegain threat.
- 1 Twinblade Paladin | threat | Adds an aggressive payoff for the lifegain plan.
- 1 Valkyrie Harbinger | threat | Provides a top-end threat for the deck's strategy.

</details>

### 2. aristocrats Commander, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 179 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $129.24 the whole deck.

**Summary:** This black-white aristocrats shell develops mana and creatures, uses its synergy pieces and sacrifice-friendly resources to maintain value, and turns a board of recurring pressure into a win with its creature threats. A substantial draw package helps it keep functioning through exchanges, while targeted answers and sweepers prevent opposing boards from taking over. It gives up some threat density and specialized aristocrats redundancy in exchange for relying on the available library and maintaining a broad defensive toolkit.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 2.70 over 63 nonland cards

<details><summary>The deck list</summary>

- 18 Plains | land | Provides reliable white mana for the deck.
- 18 Swamp | land | Provides reliable black mana for the deck.
- 1 Skullclamp | draw | Supplies a draw piece suited to a creature-focused plan.
- 1 Idol of Oblivion | draw | Adds a listed draw resource to keep cards flowing.
- 1 Mask of Memory | draw | Adds another draw option for the deck’s creatures.
- 1 Tome of Legends | draw | Provides a compact listed draw engine.
- 1 Stone of Erech | draw | Fills a draw slot from the available library.
- 1 Lembas | draw | Provides a listed draw resource with artifact utility.
- 1 Buster Sword | draw | Adds a draw-classified Equipment to the creature package.
- 1 Foot Chopper | draw | Provides another draw-classified Equipment.
- 1 Call of the Ring | draw | Adds a black draw option to sustain the deck.
- 1 Grave Venerations | draw | Supports the deck with a draw-classified enchantment.
- 1 Nasty End | draw | Provides an instant-speed draw option.
- 1 Night's Whisper | draw | Adds efficient listed card draw.
- 1 Painful Truths | draw | Provides another draw spell for rebuilding resources.
- 1 Wall of Omens | draw | Adds a creature that occupies a draw slot.
- 1 Puresteel Paladin | draw | Links the Equipment package to an additional draw piece.
- 1 Lightning Greaves | interaction | Protects an important creature or the commander.
- 1 Swiftfoot Boots | interaction | Provides a second Equipment-based protection option.
- 1 Darksteel Plate | interaction | Adds durable protection for a key permanent.
- 1 Champion's Helm | interaction | Helps safeguard the deck’s legendary centerpiece.
- 1 Gift of Immortality | interaction | Provides a protection-oriented enchantment for a key creature.
- 1 Clever Concealment | interaction | Offers a listed defensive interaction spell.
- 1 Duty Beyond Death | interaction | Supports the creature plan with protective interaction.
- 1 Together Forever | interaction | Adds another interaction piece for preserving creatures.
- 1 Sol Ring | ramp | Provides an efficient artifact ramp piece.
- 1 Arcane Signet | ramp | Adds dependable color-fixing ramp.
- 1 Commander's Sphere | ramp | Provides mana acceleration in an artifact slot.
- 1 Fellwar Stone | ramp | Adds another low-cost ramp artifact.
- 1 Thought Vessel | ramp | Supplies artifact-based mana acceleration.
- 1 Relic of Legends | ramp | Provides a ramp option that works with legendary cards.
- 1 Springleaf Drum | ramp | Adds low-cost artifact ramp to the deck.
- 1 Inherited Envelope | ramp | Provides another available ramp piece.
- 1 Deadly Dispute | ramp | Supports the sacrifice plan while filling a ramp slot.
- 1 Sword of the Animist | ramp | Pairs an Equipment with mana development.
- 1 Wayfarer's Bauble | ramp | Adds basic-land ramp to stabilize mana.
- 1 Angel of Serenity | removal | Provides creature-based removal for the deck.
- 1 Banishing Light | removal | Adds flexible permanent removal.
- 1 Bitter Triumph | removal | Provides efficient instant-speed removal.
- 1 Claim the Precious | removal | Adds black sorcery-speed removal.
- 1 Crib Swap | removal | Provides a removal spell that can answer creatures.
- 1 Destroy Evil | removal | Adds flexible white removal.
- 1 Dispatch | removal | Provides a low-cost removal option.
- 1 Fatal Push | removal | Adds efficient black interaction for creatures.
- 1 Fiend Hunter | removal | Provides removal attached to a creature body.
- 1 Generous Gift | removal | Answers a broad range of opposing permanents.
- 1 Get Lost | removal | Adds another flexible white removal spell.
- 1 Infernal Grasp | removal | Provides unconditional creature removal.
- 1 Swords to Plowshares | removal | Adds premium creature-focused removal.
- 1 Austere Command | wipe | Provides a flexible reset when the board gets ahead of the deck.
- 1 Fumigate | wipe | Adds a creature-focused board reset.
- 1 Martial Coup | wipe | Provides a reset that also contributes bodies afterward.
- 1 Dusk // Dawn | wipe | Offers a wipe option with later creature-card value.
- 1 Vanquish the Horde | wipe | Adds another dependable board-clearing spell.
- 1 Al Bhed Salvagers | synergy | Adds a listed synergy creature to the core engine.
- 1 Arcade Cabinet | synergy | Provides a synergy-classified artifact for the deck’s engine.
- 1 Gollum the Abandoned | synergy | Adds a black synergy legend to the aristocrats shell.
- 1 Gollum, Patient Plotter | synergy | Provides another synergy legend for the creature-focused plan.
- 1 Gríma Wormtongue | synergy | Adds a synergy-classified black creature.
- 1 Heirloom Auntie | synergy | Provides another synergy creature from the available library.
- 1 Joo Dee, One of Many | synergy | Adds a white synergy creature to support the core plan.
- 1 Nimble Hobbit | synergy | Provides an additional synergy body for the deck.
- 1 Phantom Train | synergy | Adds a synergy-classified artifact permanent.
- 1 Bill the Pony | threat | Provides a white threat that can pressure opponents.
- 1 Vengeful Villagers | threat | Adds another creature threat to close games.

</details>

### 3. artifacts Commander, bracket 4

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 299 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $4803.20 to buy, $4803.20 the whole deck.

**Summary:** Urza leads a dense blue artifact engine that accelerates into powerful artifact threats while maintaining steady card flow and a broad suite of answers. The deck aims to establish an artifact-heavy board, leverage its synergistic pieces, and close through its large threats after clearing resistance with removal or wipes. Its focused construction gives up broader nonartifact options in favor of keeping nearly every part of the deck aligned with the artifact plan.

- [INFO] `curve_summary`: average mana value 3.94 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Academy Ruins | land | Provides a land slot for the artifact-focused mana base.
- 1 Buried Ruin | land | Provides a land slot for the artifact-focused mana base.
- 1 Gemstone Caverns | land | Provides a land slot for the mana base.
- 1 Inventors' Fair | land | Provides a land slot for the artifact-focused mana base.
- 22 Island | land | Provides reliable blue mana across the deck.
- 1 Mishra's Workshop | land | Provides a land slot for the artifact-focused mana base.
- 1 Mystic Sanctuary | land | Provides a blue land slot for the mana base.
- 1 Otawara, Soaring City | land | Provides a blue land slot for the mana base.
- 1 Phyrexia's Core | land | Provides a land slot for the artifact-focused mana base.
- 1 Power Depot | land | Provides an artifact land slot for the mana base.
- 1 Roadside Reliquary | land | Provides a land slot for the artifact-focused mana base.
- 1 Seat of the Synod | land | Provides an artifact land slot for the mana base.
- 1 The Everflowing Well // The Myriad Pools | land | Provides a land slot for the artifact-focused mana base.
- 1 Urza's Saga | land | Provides a land slot that fits the artifact theme.
- 1 Urza's Workshop | land | Provides a land slot for the artifact-focused mana base.
- 1 Blinkmoth Urn | ramp | Supplies ramp for the artifact-heavy strategy.
- 1 Chief Engineer | ramp | Supplies ramp within the artifact plan.
- 1 Crystal Skull, Isu Spyglass | ramp | Supplies ramp while remaining an artifact.
- 1 Grand Architect | ramp | Supplies ramp for the deck's artifact plan.
- 1 Karn, Legacy Reforged | ramp | Supplies ramp and remains an artifact threat.
- 1 Krark-Clan Ironworks | ramp | Supplies ramp for an artifact-focused deck.
- 1 Metalworker | ramp | Supplies ramp for the artifact-heavy build.
- 1 Moonsnare Prototype | ramp | Supplies ramp while fitting the artifact theme.
- 1 Mox Opal | ramp | Supplies compact artifact-based ramp.
- 1 Tezzeret the Seeker | ramp | Supplies ramp for the artifact strategy.
- 1 Forensic Gadgeteer | draw | Provides card draw within the artifact plan.
- 1 One with the Machine | draw | Provides a dedicated draw spell for the artifact deck.
- 1 Reverse Engineer | draw | Provides card draw for the artifact strategy.
- 1 Riddlesmith | draw | Provides card draw within the artifact theme.
- 1 Sai, Master Thopterist | draw | Provides card draw for the artifact-focused game plan.
- 1 Tezzeret, Artifice Master | draw | Provides a card-draw engine for the deck.
- 1 Thirst for Knowledge | draw | Provides efficient card draw.
- 1 Thought Monitor | draw | Provides card draw on an artifact creature.
- 1 Thoughtcast | draw | Provides card draw that fits the artifact strategy.
- 1 Vedalken Archmage | draw | Provides card draw within the artifact theme.
- 1 Assert Authority | interaction | Provides interaction for protecting the game plan.
- 1 Darksteel Forge | interaction | Provides interaction that supports the artifact board.
- 1 Disruption Protocol | interaction | Provides interaction for key opposing plays.
- 1 Ice Out | interaction | Provides interaction for key opposing plays.
- 1 Metallic Rebuke | interaction | Provides artifact-aligned interaction.
- 1 Welding Jar | interaction | Provides compact interaction for the artifact plan.
- 1 Aether Spellbomb | removal | Provides artifact-based removal.
- 1 Arcum Dagsson | removal | Provides removal in an artifact-focused creature slot.
- 1 Lux Cannon | removal | Provides repeatable artifact-based removal.
- 1 Portal to Phyrexia | removal | Provides a powerful artifact removal option.
- 1 Ravenform | removal | Provides flexible removal.
- 1 Resculpt | removal | Provides flexible removal.
- 1 Skysovereign, Consul Flagship | removal | Provides removal on an artifact vehicle.
- 1 Spine of Ish Sah | removal | Provides artifact-based removal.
- 1 Coercive Portal | wipe | Provides a board-wipe option on an artifact.
- 1 Engineered Explosives | wipe | Provides a compact artifact board wipe.
- 1 Nevinyrral's Disk | wipe | Provides a broad artifact board wipe.
- 1 Clock of Omens | synergy | Supports artifact synergies across the deck.
- 1 Emry, Lurker of the Loch | synergy | Supports the deck's artifact synergies.
- 1 Etherium Sculptor | synergy | Supports the artifact-focused game plan.
- 1 Foundry Inspector | synergy | Supports the artifact-focused game plan.
- 1 Manifold Key | synergy | Supports artifact synergies and utility.
- 1 Mystic Forge | synergy | Supports the artifact-centric engine.
- 1 Power Artifact | synergy | Supports the deck's artifact synergies.
- 1 Reshape | synergy | Supports the artifact-focused game plan.
- 1 Scrap Trawler | synergy | Supports artifact synergies across the deck.
- 1 Simulacrum Synthesizer | synergy | Supports the deck's artifact synergies.
- 1 Transmute Artifact | synergy | Supports the artifact-focused game plan.
- 1 Unwinding Clock | synergy | Supports the artifact-centric engine.
- 1 Voltaic Key | synergy | Supports artifact synergies and utility.
- 1 Whir of Invention | synergy | Supports the artifact-focused game plan.
- 1 Cyberdrive Awakener | threat | Serves as an artifact-themed finishing threat.
- 1 Darksteel Juggernaut | threat | Serves as an artifact creature threat.
- 1 Ironheart, Clever Champion | threat | Serves as a resilient artifact threat.
- 1 Kappa Cannoneer | threat | Serves as a major artifact creature threat.
- 1 Karn, Scion of Urza | threat | Serves as a threat within the artifact strategy.
- 1 Kuldotha Forgemaster | threat | Serves as a high-impact artifact creature threat.
- 1 Master Transmuter | threat | Serves as an artifact creature threat.
- 1 Metalwork Colossus | threat | Serves as a large artifact creature threat.
- 1 Phyrexian Metamorph | threat | Serves as a flexible artifact creature threat.
- 1 The Capitoline Triad | threat | Serves as a major artifact-themed threat.
- 1 Threefold Thunderhulk | threat | Serves as an artifact creature threat.
- 1 Traxos, Scourge of Kroog | threat | Serves as a large artifact creature threat.

</details>

### 4. dinosaur tribal, bracket 2

Format: Commander. Theme: dinosaurs. Pool: any_card. Shortlist: 269 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for copy_limit. Block findings: 0.

Cost: $278.06 to buy, $278.06 the whole deck.

**Summary:** A welcoming Naya Dinosaur deck that ramps into huge creatures, builds a cohesive tribal battlefield, and turns combat into its main path to victory. Its straightforward mana development, creature-focused card advantage, protective tools, and varied removal give a new player clear decisions while preserving dramatic Dinosaur attacks.

- [INFO] `bracket_prose_rules`: bracket 2 (Core, 8+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 4.14 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Command Tower | land | Three-color commander mana
- 1 Path of Ancestry | land | Tribal mana fixing
- 1 Jungle Shrine | land | Three-color mana fixing
- 1 Jetmir's Garden | land | Fetchable three-color source
- 1 Canopy Vista | land | Green-white dual land
- 1 Cinder Glade | land | Red-green dual land
- 1 Sacred Foundry | land | Red-white dual land
- 1 Stomping Ground | land | Red-green dual land
- 1 Temple Garden | land | Green-white dual land
- 1 Exotic Orchard | land | Flexible multicolor source
- 1 Bountiful Promenade | land | Green-white mana fixing
- 1 Spire Garden | land | Red-green mana fixing
- 1 Spectator Seating | land | Red-white mana fixing
- 1 Sunpetal Grove | land | Green-white mana fixing
- 1 Rootbound Crag | land | Red-green mana fixing
- 1 Clifftop Retreat | land | Red-white mana fixing
- 1 Rockfall Vale | land | Red-green mana fixing
- 1 Overgrown Farmland | land | Green-white mana fixing
- 1 Battlefield Forge | land | Red-white mana source
- 1 Brushland | land | Green-white mana source
- 1 Karplusan Forest | land | Red-green mana source
- 1 Game Trail | land | Red-green mana fixing
- 6 Forest | land | Primary green mana for ramp
- 4 Mountain | land | Red mana for Dinosaurs
- 4 Plains | land | White mana for commander and support
- 1 Sol Ring | ramp | Fast mana acceleration
- 1 Arcane Signet | ramp | Reliable color fixing
- 1 Cultivate | ramp | Finds lands and fixes colors
- 1 Kodama's Reach | ramp | Finds lands and fixes colors
- 1 Nature's Lore | ramp | Efficient land ramp
- 1 Three Visits | ramp | Efficient land ramp
- 1 Farseek | ramp | Finds dual lands
- 1 Thunderherd Migration | ramp | Dinosaur-themed land ramp
- 1 Drover of the Mighty | ramp | Mana creature that benefits from Dinosaurs
- 1 Ranging Raptors | ramp | Dinosaur that finds basic lands
- 1 Beast Whisperer | draw | Draws cards from creature spells
- 1 Garruk's Uprising | draw | Draws from large creatures and grants trample
- 1 Guardian Project | draw | Steady creature-based card advantage
- 1 Harmonize | draw | Simple card draw
- 1 Return of the Wildspeaker | draw | Flexible large-creature draw spell
- 1 Ripjaw Raptor | draw | Draws cards when damaged
- 1 Rishkar's Expertise | draw | Large burst of cards and a free spell
- 1 Shamanic Revelation | draw | Rewards a wide Dinosaur board
- 1 Vanquisher's Banner | draw | Tribal anthem and creature-based draw
- 1 Runic Armasaur | draw | Draws from opponents' activated abilities
- 1 Akroma's Will | interaction | Protects the board or enables a decisive attack
- 1 Boros Charm | interaction | Protects permanents from destruction
- 1 Heroic Intervention | interaction | Protects the creature board
- 1 Lightning Greaves | interaction | Protects and hastens key creatures
- 1 Swiftfoot Boots | interaction | Protects and hastens key creatures
- 1 Temple Altisaur | interaction | Reduces damage dealt to Dinosaurs
- 1 Apex Altisaur | removal | Creature removal attached to a Dinosaur
- 1 Bronzebeak Foragers | removal | Removes cards from graveyards
- 1 Burning Sun's Avatar | removal | Removes creatures and players on entry
- 1 Kogla and Yidaro | removal | Flexible Dinosaur removal spell
- 1 Ravenous Sailback | removal | Removes artifacts or enchantments on entry
- 1 Savage Stomp | removal | Efficient Dinosaur creature removal
- 1 Thrashing Brontodon | removal | Repeatable artifact or enchantment removal
- 1 Tranquil Frillback | removal | Flexible graveyard and permanent removal
- 1 Commune with Dinosaurs | synergy | Finds a Dinosaur or land early
- 1 Belligerent Yearling | synergy | Grows alongside Dinosaur development
- 1 Deathgorge Scavenger | synergy | Dinosaur graveyard utility
- 1 Dinosaur Stampede | synergy | Tribal combat finisher
- 1 Dromosaur | synergy | Dinosaur tribal support creature
- 1 Huatli's Raptor | synergy | Supports counters and Dinosaur development
- 1 Hunting Velociraptor | synergy | Dinosaur combat support
- 1 Kinjalli's Caller | synergy | Reduces Dinosaur casting costs
- 1 Kinjalli's Sunwing | synergy | Disrupts opposing creatures
- 1 Marauding Raptor | synergy | Reduces Dinosaur costs and rewards creature entries
- 1 Nest Robber | synergy | Early Dinosaur presence
- 1 Orazca Frillback | synergy | Versatile Dinosaur utility
- 1 Raptor Companion | synergy | Early tribal creature
- 1 Regal Imperiosaur | synergy | Rewards a Dinosaur-heavy board
- 1 Austere Command | wipe | Flexible battlefield reset
- 1 Blasphemous Act | wipe | Efficient creature board wipe
- 1 Wakening Sun's Avatar | wipe | One-sided creature reset for Dinosaurs
- 1 Ancient Brontodon | threat | Large Dinosaur combat threat
- 1 Annoyed Altisaur | threat | Large cascade Dinosaur
- 1 Carnage Tyrant | threat | Resilient attacking threat
- 1 Ghalta, Primal Hunger | threat | Massive creature that can be cast cheaply
- 1 Ghalta, Stampede Tyrant | threat | Deploys a board of large creatures
- 1 Goring Ceratops | threat | Grants double strike while attacking
- 1 Quartzwood Crasher | threat | Creates large trampling Dinosaur tokens
- 1 Regisaur Alpha | threat | Haste enabler and token-producing Dinosaur
- 1 Rampaging Brontodon | threat | Large trampling Dinosaur
- 1 Shifting Ceratops | threat | Hasty and difficult-to-answer threat
- 1 Thundering Spineback | threat | Tribal anthem and token maker
- 1 Zetalpa, Primal Dawn | threat | Evasive, resilient top-end Dinosaur

</details>

### 5. blink Commander, owned first

Format: Commander. Theme: blink. Pool: owned_first. Shortlist: 168 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for copy_limit. Block findings: 0.

Cost: $0.00 to buy, $187.79 the whole deck.

**Summary:** A mono-white blink deck built around recycling enter-the-battlefield abilities, protecting its legendary centerpiece, and turning a stream of value creatures into an evasive late-game board. It balances broad answers and reset buttons with steady card advantage, mana development, and resilient creature pressure.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 2.81 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Command Tower | land | Reliable color access.
- 1 Abandoned Air Temple | land | Utility mana source.
- 1 Ash Barrens | land | Fixes basic land access.
- 1 Avengers Tower | land | Utility mana source.
- 1 Capital City | land | Utility mana source.
- 1 Eclipsed Realms | land | Utility mana source.
- 1 Evolving Wilds | land | Fetches a basic Plains.
- 1 Exotic Orchard | land | Flexible color access.
- 1 Fabled Passage | land | Fetches a basic Plains.
- 1 Grand Coliseum | land | Flexible color access.
- 1 Great Hall of the Citadel | land | Utility mana source.
- 1 Marsh Flats | land | Fetches a Plains.
- 1 Minas Tirith | land | Legendary utility land.
- 1 Path of Ancestry | land | Creature-focused fixing.
- 1 Plaza of Heroes | land | Supports legendary permanents.
- 1 Reliquary Tower | land | Preserves a full hand.
- 1 Secluded Courtyard | land | Creature-type mana fixing.
- 1 Sidequest: Catch a Fish // Cooking Campsite | land | Flexible land slot.
- 1 Study Hall | land | Utility mana source.
- 1 Terramorphic Expanse | land | Fetches a basic Plains.
- 1 The Grey Havens | land | Legendary utility land.
- 1 Unclaimed Territory | land | Creature-type mana fixing.
- 1 War Room | land | Land-based card advantage.
- 1 White Lotus Hideout | land | Utility mana source.
- 1 Windbrisk Heights | land | Provides a spell-bearing land.
- 11 Plains | land | Stable white mana base.
- 1 Arcane Signet | ramp | Efficient mana acceleration.
- 1 Astral Cornucopia | ramp | Scalable mana acceleration.
- 1 Bender's Waterskin | ramp | Early mana development.
- 1 Chromatic Lantern | ramp | Mana acceleration and fixing.
- 1 Commander's Sphere | ramp | Mana now and a card later.
- 1 Fellwar Stone | ramp | Efficient mana acceleration.
- 1 Sol Ring | ramp | Fast colorless acceleration.
- 1 Sword of the Animist | ramp | Combat-based land ramp.
- 1 Thought Vessel | ramp | Mana acceleration with hand support.
- 1 Wayfarer's Bauble | ramp | Finds a basic Plains.
- 1 Adventurer's Airship | draw | Repeatable card advantage.
- 1 Buster Sword | draw | Equipment-based card advantage.
- 1 Crown of Gondor | draw | Supports attacks and cards.
- 1 Diary of Dreams | draw | Dedicated card advantage.
- 1 Energybending | draw | Efficient card selection.
- 1 Idol of Oblivion | draw | Low-cost repeatable draw.
- 1 Instant Ramen | draw | Cheap card advantage.
- 1 Mask of Memory | draw | Combat-triggered card filtering.
- 1 Mirror of Galadriel | draw | Legendary card advantage.
- 1 Skullclamp | draw | Efficient creature-based draw.
- 1 Angel of Condemnation | removal | Creature-based exile interaction.
- 1 Angel of Sanctions | removal | Blinkable exile removal.
- 1 Crib Swap | removal | Flexible creature exile.
- 1 Fiend Hunter | removal | Blinkable creature removal.
- 1 Generous Gift | removal | Answers any permanent.
- 1 Get Lost | removal | Efficient permanent removal.
- 1 Journey to Nowhere | removal | Reliable creature exile.
- 1 Swords to Plowshares | removal | Premium creature exile.
- 1 Austere Command | wipe | Flexible battlefield reset.
- 1 Fumigate | wipe | Creature reset with life gain.
- 1 Vanquish the Horde | wipe | Efficient creature reset.
- 1 Boromir, Warden of the Tower | interaction | Protects against disruptive spells.
- 1 Crystal Fragments // Summon: Alexander | interaction | Versatile protection effect.
- 1 Duty Beyond Death | interaction | Keeps key creatures alive.
- 1 Reprieve | interaction | Temporizes opposing spells.
- 1 Ultimate Magic: Holy | interaction | Protects the board at instant speed.
- 1 Unbreakable Formation | interaction | Defends the creature board.
- 1 Champions of Minas Tirith | synergy | Creature value improves after repeated returns.
- 1 Ennis, Debate Moderator | synergy | Supports the deck's recurring-creature plan.
- 1 Exemplar of Light | synergy | Creature value benefits from repeat appearances.
- 1 Faramir, Field Commander | synergy | Provides reusable creature value.
- 1 Flickerwisp | synergy | Core blink effect for creatures and permanents.
- 1 Inspiring Overseer | synergy | Profits from being returned to the battlefield.
- 1 Jocasta, Automaton Avenger | synergy | Supports the recurring-creature plan.
- 1 Lembas | synergy | Reusable enter-the-battlefield value.
- 1 Palace Jailer | synergy | Blinkable value and disruption.
- 1 Personify | synergy | Directly supports the blink strategy.
- 1 South Pole Voyager | synergy | Reusable creature value.
- 1 Stiltzkin, Moogle Merchant | synergy | Benefits from repeated battlefield entries.
- 1 Wall of Omens | synergy | Reusable enter-the-battlefield value.
- 1 Weapons Vendor | synergy | Recurring creature-based value.
- 1 Angel of Serenity | threat | Large evasive finisher with impactful triggers.
- 1 Bronze Guardian | threat | Resilient artifact-based attacker.
- 1 Frontline Medic | threat | Combat-focused creature that pressures opponents.
- 1 Giada, Font of Hope | threat | Builds a powerful flying creature force.
- 1 Kataki, War's Wage | threat | Punishes artifact-heavy opponents while attacking.
- 1 Summon: Primal Garuda | threat | Evasive saga creature that closes games.
- 1 The Vision | threat | Durable legendary threat.
- 1 Westfold Rider | threat | Creature pressure with useful utility.
- 1 Zack Fair | threat | Legendary combat threat.
- 1 Bastion Protector | threat | Protective body that applies pressure.
- 1 Darksteel Plate | threat | Makes a major attacker hard to answer.
- 1 Lightning Greaves | threat | Lets key creatures attack safely and quickly.

</details>

### 6. Modern tempo, tournament

Format: Modern. Theme: tempo. Pool: any_card. Shortlist: 167 names.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $358.88 to buy, $358.88 the whole deck.

**Summary:** This blue-red tempo deck establishes creature pressure while using draw, interaction, and removal to keep the game focused on its own plan. It wins by sustaining that pressure through a tightly supported board and uses its temporal cards as additional synergy. It gives up broader ramp packages and dedicated sweepers in favor of staying focused on proactive threats and efficient support.

- [INFO] `curve_summary`: average mana value 2.28 over 36 nonland cards

<details><summary>The deck list</summary>

- 4 Steam Vents | land | Blue-red land for the deck's mana base.
- 4 Spirebluff Canal | land | Blue-red land for the deck's mana base.
- 4 Riverglide Pathway // Lavaglide Pathway | land | Flexible blue-red land for the deck's mana base.
- 4 Shivan Reef | land | Blue-red land for the deck's mana base.
- 4 Island | land | Basic blue land for the deck's mana base.
- 4 Mountain | land | Basic red land for the deck's mana base.
- 4 Ledger Shredder | threat | Creature threat for the tempo plan.
- 4 Ragavan, Nimble Pilferer | threat | Creature threat for the tempo plan.
- 4 Faerie Mastermind | threat | Creature threat for the tempo plan.
- 2 Consider | draw | Draw spell that supports the tempo plan.
- 2 Opt | draw | Draw spell that supports the tempo plan.
- 2 Preordain | draw | Draw spell that supports the tempo plan.
- 4 Counterspell | interaction | Core interaction for protecting the tempo plan.
- 2 Spell Pierce | interaction | Additional interaction for the tempo plan.
- 4 Lightning Bolt | removal | Efficient removal for clearing the way.
- 2 Into the Flood Maw | removal | Removal that supports the tempo plan.
- 2 Untimely Malfunction | removal | Removal that supports the tempo plan.
- 2 Temporal Mastery | synergy | Temporal synergy piece for the deck's plan.
- 2 Temporal Trespass | synergy | Temporal synergy piece for the deck's plan.

</details>

### 7. Modern burn, casual

Format: Modern. Theme: burn. Pool: any_card. Shortlist: 290 names.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $34.98 to buy, $34.98 the whole deck.

**Summary:** This mono-red burn deck opens with direct damage and persistent damage-oriented permanents, using its draw spells to keep the pressure going rather than letting the opponent stabilize. Its creatures give it a second route to victory by attacking after the early burn has softened up the opponent, with larger threats helping close games that go longer. The tradeoff is a narrow, proactive plan: it prioritizes damage and momentum over broad answers or defensive tools.

- [INFO] `curve_summary`: average mana value 2.78 over 36 nonland cards

<details><summary>The deck list</summary>

- 24 Mountain | land | Provides the reliable red mana base for a mono-red deck.
- 4 Risk Factor | draw | Keeps pressure on the opponent while refilling resources.
- 2 Browbeat | draw | Adds more card access while presenting a punishing choice.
- 2 Sudden Breakthrough | ramp | Supplies the requested mana acceleration while fitting the spell-heavy plan.
- 4 Lightning Bolt | removal | A direct, efficient burn spell for removing blockers or pressuring life totals.
- 2 Lava Dart | removal | Provides additional cheap damage and flexible creature removal.
- 2 Eidolon of the Great Revel | synergy | Punishes opponents for relying on inexpensive spells while advancing the damage plan.
- 2 Thermo-Alchemist | synergy | Turns the deck's steady stream of spells into repeated damage pressure.
- 2 Firebrand Archer | synergy | Adds incremental damage whenever the deck casts its burn and draw spells.
- 2 Roiling Vortex | synergy | Maintains passive pressure alongside the deck's direct-damage spells.
- 4 Akki Lavarunner // Tok-Tok, Volcano Born | threat | A red threat that helps maintain a creature presence on the battlefield.
- 4 Keral Keep Disciples | threat | Provides a dedicated attacking threat for the red damage strategy.
- 4 Hazoret the Fervent | threat | Gives the deck a resilient top-end threat after early burn pressure.
- 2 Torbran, Thane of Red Fell | threat | Supports the deck's damage-focused battlefield presence as a finishing threat.

</details>

### 8. Modern lifegain, FNM

Format: Modern. Theme: lifegain. Pool: any_card. Shortlist: 301 names.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $222.09 to buy, $222.09 the whole deck.

**Summary:** This white-black lifegain deck establishes its early lifegain synergies, turns those pieces into pressure with creatures such as Ajani's Pridemate and Bloodbond Vampire, and keeps cards flowing through its draw package. It wins by building a board of lifegain-focused threats and protecting its momentum with removal. The deck gives up some speed and flexibility for a focused creature-and-synergy game plan.

- [INFO] `curve_summary`: average mana value 3.17 over 36 nonland cards

<details><summary>The deck list</summary>

- 7 Plains | land | Provides a basic white land base for the deck.
- 6 Swamp | land | Provides a basic black land base for the deck.
- 4 Godless Shrine | land | Provides a white-black land option.
- 3 Shattered Sanctum | land | Provides another white-black land option.
- 2 Caves of Koilos | land | Adds white-black land coverage.
- 2 Isolated Chapel | land | Rounds out the white-black land base.
- 2 Arcane Signet | ramp | Supplies the deck's dedicated ramp package.
- 4 Murderous Rider // Swift End | removal | Provides reliable removal slots while remaining a creature card.
- 2 Solitude | removal | Completes the removal package.
- 4 Dawn of Hope | draw | Forms the main card-draw package for the lifegain plan.
- 2 Inspiring Overseer | draw | Adds card draw on a creature body.
- 4 Soul Warden | synergy | A core lifegain synergy piece.
- 4 Ajani's Pridemate | synergy | A central payoff for the deck's lifegain synergies.
- 4 Attended Healer | threat | Provides a lifegain-focused threat.
- 4 Bloodbond Vampire | threat | Adds another threat aligned with the lifegain plan.
- 4 Angel of Invention | threat | Provides a substantial creature threat.
- 2 Archangel of Thune | threat | Finishes the threat package with a powerful lifegain-themed creature.

</details>

### 9. Standard midrange, FNM

Format: Standard. Theme: midrange. Pool: any_card. Shortlist: 164 names.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $229.20 to buy, $229.20 the whole deck.

**Summary:** This black-green midrange deck develops its mana with creatures and permanents, uses removal to keep opposing pieces from taking over, and refuels with its dedicated draw cards. It aims to win by converting its growing creature board into sustained pressure once the table has been stabilized. It gives up early aggression for a stronger board-focused long game.

- [INFO] `curve_summary`: average mana value 2.33 over 36 nonland cards

<details><summary>The deck list</summary>

- 4 Forest | land | Basic green source for the deck's black-green mana base.
- 4 Swamp | land | Basic black source for the deck's black-green mana base.
- 4 Blooming Marsh | land | Black-green land that supports the two-color mana base.
- 4 Deathcap Glade | land | Black-green land that supports the two-color mana base.
- 4 Overgrown Tomb | land | Black-green land that supports the two-color mana base.
- 4 Underground Mortuary | land | Black-green land that supports the two-color mana base.
- 4 Phyrexian Arena | draw | Dedicated draw source for keeping resources flowing in longer games.
- 2 Darkstar Augur | draw | Creature-based draw that contributes to the deck's board presence.
- 2 Llanowar Elves | ramp | Early ramp for advancing the deck's larger plays.
- 4 Bitter Triumph | removal | Efficient removal for answering opposing problems.
- 2 Maelstrom Pulse | removal | Additional removal that broadens the deck's answers.
- 4 Aftermath Analyst | ramp | Ramp creature that helps build a substantial board.
- 4 Elvish Archdruid | ramp | Ramp creature that supports the deck's mana-focused creature plan.
- 4 Enduring Vitality | ramp | Ramp permanent that reinforces the deck's board development.
- 4 Goldvein Hydra | ramp | Ramp creature that provides a substantial body alongside mana development.
- 4 Insidious Roots | ramp | Ramp enchantment that supports the deck's long-game development.
- 2 Loot, Exuberant Explorer | ramp | Additional ramp creature for the deck's midrange curve.

</details>

### 10. Standard aggro, tournament

Format: Standard. Theme: aggro. Pool: any_card. Shortlist: 294 names.

Cards: 64 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $83.42 to buy, $83.42 the whole deck.

**Summary:** This red-white aggro deck aims to apply pressure early with a dense suite of threats, using its synergy cards to support that attack and its draw cards to maintain momentum. Removal and interaction help push through the final damage or disrupt opposing plans. It gives up heavier late-game options in favor of a focused, proactive game plan.

- [WARN] `land_count`: 28 lands: the guide range for this format is 20 to 27
- [INFO] `curve_summary`: average mana value 2.28 over 36 nonland cards

<details><summary>The deck list</summary>

- 10 Mountain | land | Provides red mana in the red-white mana base.
- 10 Plains | land | Provides white mana in the red-white mana base.
- 2 Inspiring Vantage | land | Supports the red-white mana base.
- 2 Sacred Foundry | land | Supports the red-white mana base.
- 4 Sundown Pass | land | Supports the red-white mana base.
- 4 Slickshot Show-Off | synergy | Serves as an aggressive synergy piece.
- 4 Fugitive Codebreaker | draw | Supplies card draw for an aggressive deck.
- 2 Reckless Lackey | draw | Adds card draw while fitting the red aggressive plan.
- 4 Boros Charm | interaction | Provides flexible interaction.
- 2 Dawn's Truce | interaction | Adds interaction to protect the aggressive plan.
- 4 Harsh Annotation | removal | Provides efficient removal support.
- 4 Case of the Gateway Express | removal | Adds removal to clear opposing obstacles.
- 4 Burnout Bashtronaut | threat | Provides an aggressive threat.
- 4 Redcap Gutter-Dweller | threat | Provides an aggressive threat.
- 4 Teapot Slinger | threat | Rounds out the deck's threat base.

</details>

### 11. Commander with a locked card

Format: Commander. Theme: sacrifice. Pool: any_card. Shortlist: 296 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $1876.48 to buy, $1876.48 the whole deck.

**Summary:** Karlov of the Ghost Council leads a focused sacrifice deck that builds resources through its sacrifice synergies, keeps pressure on opponents with removal and interaction, and closes through its larger threats. It aims to win by turning the sacrifice-focused core into sustained pressure while retaining sweeping answers for crowded boards. The deck gives up broader themes and relies on its sacrifice pieces working together.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 3.22 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Arguel's Blood Fast // Temple of Aclazotz | land | Included as a land option from the shortlist.
- 1 Cavern of Souls | land | Included as a land option from the shortlist.
- 1 Caves of Koilos | land | Included as a land option from the shortlist.
- 1 City of Brass | land | Included as a land option from the shortlist.
- 1 Command Tower | land | Included as a land option from the shortlist.
- 1 Diamond Valley | land | Included as a land option from the shortlist.
- 1 Dust Bowl | land | Included as a land option from the shortlist.
- 1 Exotic Orchard | land | Included as a land option from the shortlist.
- 1 Fetid Heath | land | Included as a land option from the shortlist.
- 1 Fountainport | land | Included as a land option from the shortlist.
- 1 Gemstone Caverns | land | Included as a land option from the shortlist.
- 1 Godless Shrine | land | Included as a land option from the shortlist.
- 1 High Market | land | Included as a land option from the shortlist.
- 1 Hostile Hostel // Creeping Inn | land | Included as a land option from the shortlist.
- 1 Ifnir Deadlands | land | Included as a land option from the shortlist.
- 1 Isolated Chapel | land | Included as a land option from the shortlist.
- 1 Lake of the Dead | land | Included as a land option from the shortlist.
- 1 Lazotep Quarry | land | Included as a land option from the shortlist.
- 1 Mana Confluence | land | Included as a land option from the shortlist.
- 1 Miren, the Moaning Well | land | Included as a land option from the shortlist.
- 1 Nykthos, Shrine to Nyx | land | Included as a land option from the shortlist.
- 1 Phyrexia's Core | land | Included as a land option from the shortlist.
- 1 Phyrexian Tower | land | Included as a land option from the shortlist.
- 1 Reflecting Pool | land | Included as a land option from the shortlist.
- 1 Scavenger Grounds | land | Included as a land option from the shortlist.
- 1 Scrubland | land | Included as a land option from the shortlist.
- 1 Shattered Sanctum | land | Included as a land option from the shortlist.
- 1 Shefet Dunes | land | Included as a land option from the shortlist.
- 1 Tarrian's Journal // The Tomb of Aclazotz | land | Included as a land option from the shortlist.
- 1 Temple of Silence | land | Included as a land option from the shortlist.
- 1 Treasure Map // Treasure Cove | land | Included as a land option from the shortlist.
- 1 Vault of Champions | land | Included as a land option from the shortlist.
- 1 Westvale Abbey // Ormendahl, Profane Prince | land | Included as a land option from the shortlist.
- 2 Plains | land | Basic land support for the deck.
- 1 Swamp | land | Basic land support for the deck.
- 1 Sol Ring | ramp | Kept as requested and used for ramp.
- 1 Ashnod's Altar | ramp | Provides ramp for the sacrifice plan.
- 1 Culling the Weak | ramp | Provides ramp for the sacrifice plan.
- 1 Deadly Dispute | ramp | Provides ramp for the sacrifice plan.
- 1 Phyrexian Altar | ramp | Provides ramp for the sacrifice plan.
- 1 Pitiless Plunderer | ramp | Provides ramp for the sacrifice plan.
- 1 Priest of Forgotten Gods | ramp | Provides ramp for the sacrifice plan.
- 1 Sifter of Skulls | ramp | Provides ramp for the sacrifice plan.
- 1 Skullport Merchant | ramp | Provides ramp for the sacrifice plan.
- 1 Warren Soultrader | ramp | Provides ramp for the sacrifice plan.
- 1 Baron Bertram Graywater | draw | Supplies draw within the deck's sacrifice-focused shell.
- 1 Bushmeat Poacher | draw | Supplies draw within the deck's sacrifice-focused shell.
- 1 Corrupted Conviction | draw | Supplies draw within the deck's sacrifice-focused shell.
- 1 Disciple of Bolas | draw | Supplies draw within the deck's sacrifice-focused shell.
- 1 Lord Skitter's Butcher | draw | Supplies draw within the deck's sacrifice-focused shell.
- 1 Shadowheart, Dark Justiciar | draw | Supplies draw within the deck's sacrifice-focused shell.
- 1 Smothering Abomination | draw | Supplies draw within the deck's sacrifice-focused shell.
- 1 Tevesh Szat, Doom of Fools | draw | Supplies draw within the deck's sacrifice-focused shell.
- 1 Vampiric Rites | draw | Supplies draw within the deck's sacrifice-focused shell.
- 1 Village Rites | draw | Supplies draw within the deck's sacrifice-focused shell.
- 1 Cartel Aristocrat | interaction | Offers interaction that supports the sacrifice plan.
- 1 Fanatical Devotion | interaction | Offers interaction that supports the sacrifice plan.
- 1 Flare of Fortitude | interaction | Offers interaction that supports the sacrifice plan.
- 1 Gift of Doom | interaction | Offers interaction that supports the sacrifice plan.
- 1 Nightmare Shepherd | interaction | Offers interaction that supports the sacrifice plan.
- 1 Promise of Tomorrow | interaction | Offers interaction that supports the sacrifice plan.
- 1 Attrition | removal | Provides removal for opposing problems.
- 1 Ayli, Eternal Pilgrim | removal | Provides removal for opposing problems.
- 1 Bone Shards | removal | Provides removal for opposing problems.
- 1 Dictate of Erebos | removal | Provides removal for opposing problems.
- 1 Eaten Alive | removal | Provides removal for opposing problems.
- 1 Grave Pact | removal | Provides removal for opposing problems.
- 1 Teysa, Orzhov Scion | removal | Provides removal for opposing problems.
- 1 Yawgmoth, Thran Physician | removal | Provides removal for opposing problems.
- 1 Altar of Dementia | synergy | Advances the deck's sacrifice synergy.
- 1 Bartolomé del Presidio | synergy | Advances the deck's sacrifice synergy.
- 1 Bastion of Remembrance | synergy | Advances the deck's sacrifice synergy.
- 1 Carrion Feeder | synergy | Advances the deck's sacrifice synergy.
- 1 Chthonian Nightmare | synergy | Advances the deck's sacrifice synergy.
- 1 Elas il-Kor, Sadistic Pilgrim | synergy | Advances the deck's sacrifice synergy.
- 1 Fleshtaker | synergy | Advances the deck's sacrifice synergy.
- 1 Hidden Stockpile | synergy | Advances the deck's sacrifice synergy.
- 1 Nantuko Husk | synergy | Advances the deck's sacrifice synergy.
- 1 Open the Graves | synergy | Advances the deck's sacrifice synergy.
- 1 Spawning Pit | synergy | Advances the deck's sacrifice synergy.
- 1 Viscera Seer | synergy | Advances the deck's sacrifice synergy.
- 1 Woe Strider | synergy | Advances the deck's sacrifice synergy.
- 1 Zulaport Cutthroat | synergy | Advances the deck's sacrifice synergy.
- 1 Abhorrent Overlord | threat | Serves as a major threat for closing games.
- 1 Basri's Lieutenant | threat | Serves as a major threat for closing games.
- 1 Felisa, Fang of Silverquill | threat | Serves as a major threat for closing games.
- 1 Ghoulcaller Gisa | threat | Serves as a major threat for closing games.
- 1 Kuldotha Forgemaster | threat | Serves as a major threat for closing games.
- 1 Liesa, Forgotten Archangel | threat | Serves as a major threat for closing games.
- 1 Liliana, Heretical Healer // Liliana, Defiant Necromancer | threat | Serves as a major threat for closing games.
- 1 Mondrak, Glory Dominus | threat | Serves as a major threat for closing games.
- 1 Ratadrabik of Urborg | threat | Serves as a major threat for closing games.
- 1 Razaketh, the Foulblooded | threat | Serves as a major threat for closing games.
- 1 Requiem Angel | threat | Serves as a major threat for closing games.
- 1 Vindictive Vampire | threat | Serves as a major threat for closing games.
- 1 Liliana, Dreadhorde General | wipe | Acts as a sweeping answer when the table gets ahead.
- 1 The Meathook Massacre | wipe | Acts as a sweeping answer when the table gets ahead.
- 1 Toxic Deluge | wipe | Acts as a sweeping answer when the table gets ahead.

</details>

### 12. Commander on a budget

Format: Commander. Theme: tokens. Pool: any_card. Shortlist: 296 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $44.31 to buy, $44.31 the whole deck.

**Summary:** Adeline leads a board-focused white token deck that develops a wide force, backs it with token-focused permanents, and turns that board into sustained combat pressure. The deck wins by keeping enough threats on the table to overwhelm opponents while draw, ramp, removal, and a few reset buttons keep the plan moving. Its tradeoff is a straightforward budget mana base and a reliance on maintaining a creature board rather than pursuing a fast standalone finish.

- [INFO] `bracket_prose_rules`: bracket 2 (Core, 8+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 3.25 over 63 nonland cards

<details><summary>The deck list</summary>

- 35 Plains | land | Provides the deck’s primary white mana base.
- 1 Command Tower | land | Provides reliable mana for Adeline’s deck.
- 1 Angelic Sell-Sword | draw | Adds a budget draw option.
- 1 Bygone Bishop | draw | Adds a budget draw option.
- 1 Faramir, Field Commander | draw | Adds a budget draw option.
- 1 Dawn of Hope | draw | Adds a budget draw option.
- 1 Idol of Oblivion | draw | Adds a budget draw option.
- 1 Glimmer Seeker | draw | Adds a budget draw option.
- 1 Platoon Dispenser | draw | Adds a budget draw option.
- 1 Sarah Jane Smith | draw | Adds a budget draw option.
- 1 Staff of the Storyteller | draw | Adds a budget draw option.
- 1 Sunpearl Kirin | draw | Adds a budget draw option.
- 1 Goldvein Pick | ramp | Provides a low-cost ramp piece.
- 1 Prying Blade | ramp | Provides a low-cost ramp piece.
- 1 Noble's Purse | ramp | Provides a low-cost ramp piece.
- 1 Currency Converter | ramp | Provides a low-cost ramp piece.
- 1 Collector's Vault | ramp | Provides a low-cost ramp piece.
- 1 Charisma Bobblehead | ramp | Provides a low-cost ramp piece.
- 1 Coin of Mastery | ramp | Provides a low-cost ramp piece.
- 1 Druidic Satchel | ramp | Provides a low-cost ramp piece.
- 1 Karn, Living Legacy | ramp | Provides a low-cost ramp piece.
- 1 Monologue Tax | ramp | Provides a low-cost ramp piece.
- 1 Aerial Assault | removal | Supplies efficient targeted removal.
- 1 Banishing Slash | removal | Supplies efficient targeted removal.
- 1 Battle Menu | removal | Supplies efficient targeted removal.
- 1 Citizen's Crowbar | removal | Supplies efficient targeted removal.
- 1 Generous Gift | removal | Supplies efficient targeted removal.
- 1 Kellan's Lightblades | removal | Supplies efficient targeted removal.
- 1 Righteous Confluence | removal | Supplies efficient targeted removal.
- 1 Stroke of Midnight | removal | Supplies efficient targeted removal.
- 1 Ainok Strike Leader | interaction | Adds a protective interaction piece for the board.
- 1 Basri Ket | interaction | Adds a protective interaction piece for the board.
- 1 Blessed Sanctuary | interaction | Adds a protective interaction piece for the board.
- 1 Lena, Selfless Champion | interaction | Adds a protective interaction piece for the board.
- 1 Rootborn Defenses | interaction | Adds a protective interaction piece for the board.
- 1 Spirit Bonds | interaction | Adds a protective interaction piece for the board.
- 1 Ceaseless Conflict | wipe | Provides a budget board reset.
- 1 Crisis of Conscience | wipe | Provides a budget board reset.
- 1 Descend upon the Sinful | wipe | Provides a budget board reset.
- 1 Aligned Heart | synergy | Supports the deck’s token-focused plan.
- 1 Anafenza, Unyielding Lineage | synergy | Supports the deck’s token-focused plan.
- 1 Animation Module | synergy | Supports the deck’s token-focused plan.
- 1 Anointer Priest | synergy | Supports the deck’s token-focused plan.
- 1 Automated Assembly Line | synergy | Supports the deck’s token-focused plan.
- 1 Cat Collector | synergy | Supports the deck’s token-focused plan.
- 1 Cathar's Call | synergy | Supports the deck’s token-focused plan.
- 1 Clarion Spirit | synergy | Supports the deck’s token-focused plan.
- 1 Divine Visitation | synergy | Supports the deck’s token-focused plan.
- 1 Eyes in the Skies | synergy | Supports the deck’s token-focused plan.
- 1 Felidar Retreat | synergy | Supports the deck’s token-focused plan.
- 1 Hero of Precinct One | synergy | Supports the deck’s token-focused plan.
- 1 Intangible Virtue | synergy | Supports the deck’s token-focused plan.
- 1 Retrofitter Foundry | synergy | Supports the deck’s token-focused plan.
- 1 Ajani's Chosen | threat | Provides a budget token-oriented threat.
- 1 Archon of Sun's Grace | threat | Provides a budget token-oriented threat.
- 1 Attended Healer | threat | Provides a budget token-oriented threat.
- 1 Basri's Lieutenant | threat | Provides a budget token-oriented threat.
- 1 Cemetery Protector | threat | Provides a budget token-oriented threat.
- 1 Defiler of Faith | threat | Provides a budget token-oriented threat.
- 1 Dragonback Lancer | threat | Provides a budget token-oriented threat.
- 1 Drogskol Cavalry | threat | Provides a budget token-oriented threat.
- 1 Emeria Angel | threat | Provides a budget token-oriented threat.
- 1 Gideon, Ally of Zendikar | threat | Provides a budget token-oriented threat.
- 1 Go-Shintai of Shared Purpose | threat | Provides a budget token-oriented threat.
- 1 God-Eternal Oketra | threat | Provides a budget token-oriented threat.

</details>

### 13. owned first, and the commander is not owned

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 229 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $144.74 the whole deck.

**Summary:** Karlov of the Ghost Council leads a white-black lifegain deck that develops with artifact mana, supportive creatures, and protective equipment before turning its lifegain synergy into pressure. A mix of Angels, legendary creatures, and other threats gives the deck several ways to close games, while efficient removal and board wipes keep opposing boards manageable. The deck gives up multicolor utility and relies on its commander, creature presence, and synergy pieces to maintain momentum.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [WARN] `not_owned`: Karlov of the Ghost Council: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 2.98 over 63 nonland cards

<details><summary>The deck list</summary>

- 18 Plains | land | Provides a reliable white mana base for the deck.
- 18 Swamp | land | Provides a reliable black mana base for the deck.
- 1 Arcane Signet | ramp | Provides efficient mana acceleration.
- 1 Astral Cornucopia | ramp | Adds another mana-producing artifact to the deck.
- 1 Bender's Waterskin | ramp | Fills a mana acceleration slot from the library.
- 1 Commander's Sphere | ramp | Provides flexible mana acceleration.
- 1 Fellwar Stone | ramp | Provides early mana acceleration.
- 1 Sol Ring | ramp | Provides efficient mana acceleration.
- 1 Springleaf Drum | ramp | Adds low-cost mana acceleration.
- 1 Thought Vessel | ramp | Provides an additional mana source.
- 1 Wayfarer's Bauble | ramp | Provides mana acceleration from a colorless card.
- 1 White Lotus Tile | ramp | Rounds out the artifact-based ramp package.
- 1 Buster Sword | draw | Provides card access while fitting the deck's artifact package.
- 1 Call of the Ring | draw | Adds a black draw option from the library.
- 1 Exemplar of Light | draw | Supports the deck's white creature core while providing draw.
- 1 Idol of Oblivion | draw | Provides card access through a colorless permanent.
- 1 Inspiring Overseer | draw | Supports the deck's white creature plan and provides draw.
- 1 Lembas | draw | Provides card access from a colorless card.
- 1 Mask of Memory | draw | Adds repeatable card access to the equipment package.
- 1 Night's Whisper | draw | Provides a compact black draw spell.
- 1 Puresteel Paladin | draw | Supports equipment while contributing to card access.
- 1 Skullclamp | draw | Provides card access through the creature base.
- 1 Bastion Protector | interaction | Helps protect the deck's central legendary creature.
- 1 Boromir, Warden of the Tower | interaction | Provides a creature-based interaction piece.
- 1 Champion's Helm | interaction | Protects an important legendary permanent.
- 1 Clever Concealment | interaction | Provides broad protective interaction.
- 1 Lightning Greaves | interaction | Protects and supports a key creature.
- 1 Swiftfoot Boots | interaction | Adds another efficient protection piece.
- 1 Banishing Light | removal | Provides flexible white removal.
- 1 Bitter Triumph | removal | Provides efficient black removal.
- 1 Crib Swap | removal | Adds creature-focused removal.
- 1 Dispatch | removal | Provides a low-cost removal option alongside the artifacts.
- 1 Dismember | removal | Adds removal that fits the deck's colors.
- 1 Generous Gift | removal | Provides flexible permanent removal.
- 1 Path to Exile | removal | Provides efficient creature removal.
- 1 Swords to Plowshares | removal | Provides another efficient creature removal spell.
- 1 Austere Command | wipe | Provides a flexible battlefield reset.
- 1 Fumigate | wipe | Provides a white board-clearing option for difficult boards.
- 1 Vanquish the Horde | wipe | Adds a further white battlefield reset.
- 1 Aerith Gainsborough | synergy | Its listed synergy role supports the lifegain-focused plan.
- 1 Aettir and Priwen | synergy | Supports the deck's synergy and equipment elements.
- 1 Angel of Vitality | synergy | Its listed synergy role directly supports the lifegain plan.
- 1 Compassionate Healer | synergy | Its listed synergy role fits the deck's lifegain focus.
- 1 Elixir | synergy | Provides a colorless synergy piece for the deck.
- 1 Kor Firewalker | synergy | Its listed synergy role supports the white lifegain shell.
- 1 Light of Promise | synergy | Supports the deck's lifegain-centered synergy package.
- 1 Momo, Playful Pet | synergy | Adds another synergy creature from the library.
- 1 Night Nurse, Healer of Heroes | synergy | Its listed synergy role fits the deck's lifegain plan.
- 1 Prideful Feastling | synergy | Adds a synergy creature in the deck's colors.
- 1 Rosie Cotton of South Lane | synergy | Provides a white synergy piece for the creature plan.
- 1 Second Breakfast | synergy | Supports the deck's lifegain-focused synergy package.
- 1 Stone Docent | synergy | Adds another synergy creature from the library.
- 1 White Mage's Staff | synergy | Supports the deck's artifact and synergy package.
- 1 Angel of Invention | threat | Provides a white creature threat for closing games.
- 1 Bill the Pony | threat | Adds a creature threat from the library.
- 1 Dawnhand Eulogist | threat | Provides a threat that fits the deck's colors.
- 1 Gollum, Silent Slinker // Meager Meal | threat | Adds a black threat with an attached adventure.
- 1 Invisible Woman, Sue Storm | threat | Provides a legendary creature threat.
- 1 Lo and Li, Twin Tutors | threat | Adds a legendary creature threat from the library.
- 1 Lyra Dawnbringer | threat | Provides a powerful white creature threat.
- 1 Minwu, White Mage | threat | Adds another white creature threat.
- 1 Reaping Willow | threat | Provides a black creature threat.
- 1 Shattered Angel | threat | Adds a white creature threat that fits the deck's theme.
- 1 Sneering Shadewriter | threat | Provides another black creature threat.
- 1 Victory's Herald | threat | Provides a top-end white creature threat.

</details>

### 14. the user delegates the commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 244 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $988.01 to buy, $988.01 the whole deck.

**Summary:** This deck ramps into a dense line of Dragons, then uses Atarka, World Render to make combat attacks especially punishing. Dragon-tribal support, card advantage, and protective interaction help keep the board moving toward large combat turns, while removal and sweepers clear opposition that gets in the way. It wins primarily by maintaining a force of flying Dragon threats and turning that board into decisive attacks; in exchange, it is built around expensive creatures and can be slower when its early mana development is disrupted.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 4.24 over 62 nonland cards
- [INFO] `cards_trimmed`: the list was 1 card over, so the builder cut this card: Dragonstorm

<details><summary>The deck list</summary>

- 1 Cavern of Souls | land | Dragon-focused land slot.
- 1 Cinder Glade | land | Red-green mana source.
- 1 Command Tower | land | Reliable commander-color mana.
- 1 Commercial District | land | Red-green mana source.
- 1 Cragcrown Pathway // Timbercrown Pathway | land | Flexible red or green mana source.
- 1 Crucible of the Spirit Dragon | land | Dragon-focused land slot.
- 1 Exotic Orchard | land | Flexible color-fixing land.
- 1 Game Trail | land | Red-green mana source.
- 1 Gruul Turf | land | Red-green mana source.
- 1 Haven of the Spirit Dragon | land | Dragon-focused land slot.
- 1 Karplusan Forest | land | Reliable red-green mana source.
- 1 Mana Confluence | land | Flexible color-fixing land.
- 1 Mossfire Valley | land | Red-green mana source.
- 1 Path of Ancestry | land | Dragon-tribal mana source.
- 1 Reflecting Pool | land | Flexible color-fixing land.
- 1 Rockfall Vale | land | Red-green mana source.
- 1 Rootbound Crag | land | Red-green mana source.
- 1 Secluded Courtyard | land | Dragon-tribal mana source.
- 1 Sheltered Thicket | land | Red-green mana source.
- 1 Spire Garden | land | Red-green mana source.
- 1 Stomping Ground | land | Reliable red-green mana source.
- 1 Taiga | land | Reliable red-green mana source.
- 1 Temple of Abandon | land | Red-green mana source.
- 1 Temple of the Dragon Queen | land | Dragon-focused color fixing.
- 1 Thornspire Verge | land | Red-green mana source.
- 1 Three Tree City | land | Dragon-tribal land slot.
- 1 Unclaimed Territory | land | Dragon-tribal mana source.
- 5 Forest | land | Basic green mana base.
- 5 Mountain | land | Basic red mana base.
- 1 Ancient Copper Dragon | ramp | Dragon-based mana acceleration.
- 1 Carnelian Orb of Dragonkind | ramp | Dragon-focused mana acceleration.
- 1 Dragon's Hoard | ramp | Dragon-focused mana acceleration.
- 1 Ganax, Astral Hunter | ramp | Dragon-based mana acceleration.
- 1 Goldspan Dragon | ramp | Dragon-based mana acceleration.
- 1 Jade Orb of Dragonkind | ramp | Dragon-focused mana acceleration.
- 1 Klauth, Unrivaled Ancient | ramp | High-impact Dragon mana acceleration.
- 1 Old Gnawbone | ramp | High-impact Dragon mana acceleration.
- 1 Orb of Dragonkind | ramp | Dragon-focused mana acceleration.
- 1 Scaled Nurturer | ramp | Early Dragon-focused ramp.
- 1 Dragonborn Champion | draw | Dragon-themed card advantage.
- 1 Elemental Bond | draw | Card advantage for large creatures.
- 1 Garruk's Uprising | draw | Card advantage for large creatures.
- 1 Guardian Project | draw | Creature-based card advantage.
- 1 Return of the Wildspeaker | draw | Flexible card-advantage spell.
- 1 Rishkar's Expertise | draw | High-impact card-advantage spell.
- 1 Skullclamp | draw | Efficient card-advantage equipment.
- 1 Sylvan Library | draw | Reliable card selection and advantage.
- 1 Toski, Bearer of Secrets | draw | Combat-based card advantage.
- 1 Vanquisher's Banner | draw | Dragon-tribal card advantage.
- 1 Heroic Intervention | interaction | Protects the developed board.
- 1 Lightning Greaves | interaction | Protects a key creature.
- 1 Mithril Coat | interaction | Protects a key creature.
- 1 Swiftfoot Boots | interaction | Protects a key creature.
- 1 Tibalt's Trickery | interaction | Flexible stack interaction.
- 1 Veil of Summer | interaction | Efficient protective interaction.
- 1 Dragon Tempest | removal | Dragon-themed removal support.
- 1 Dragonlord Atarka | removal | Dragon threat that also supplies removal.
- 1 Drakuseth, Maw of Flames | removal | Large Dragon with removal value.
- 1 Glorybringer | removal | Dragon-based removal option.
- 1 Scourge of Valkas | removal | Dragon-themed removal payoff.
- 1 Terror of the Peaks | removal | Dragon-based removal payoff.
- 1 Wrathful Red Dragon | removal | Dragon-themed removal support.
- 1 Invasion of Tarkir // Defiant Thundermaw | removal | Dragon-themed removal with a Dragon follow-up.
- 1 Breaching Dragonstorm | synergy | Dragon-focused engine piece.
- 1 Crucible of Fire | synergy | Dragon-tribal payoff.
- 1 Dracogenesis | synergy | Dragon-focused engine piece.
- 1 Dragonkin Berserker | synergy | Dragon-focused support creature.
- 1 Dragonlord's Servant | synergy | Dragon-focused support creature.
- 1 Dragonspeaker Shaman | synergy | Dragon-focused support creature.
- 1 Firespitter Whelp | synergy | Dragon-focused support creature.
- 1 Minion of the Mighty | synergy | Early Dragon-focused support.
- 1 Sarkhan's Triumph | synergy | Finds a needed Dragon.
- 1 Scaleguard Sentinels | synergy | Dragon-focused support creature.
- 1 Shivan Devastator | synergy | Flexible Dragon-themed payoff.
- 1 Slumbering Dragon | synergy | Early Dragon-themed board presence.
- 1 The Dragon-Kami Reborn // Dragon-Kami's Egg | synergy | Dragon-focused engine piece.
- 1 Ancient Bronze Dragon | threat | Powerful Dragon combat threat.
- 1 Blast-Furnace Hellkite | threat | High-impact Dragon threat.
- 1 Chiss-Goria, Forge Tyrant | threat | Large Dragon threat.
- 1 Dragon Broodmother | threat | Dragon threat that builds board presence.
- 1 Hellkite Charger | threat | Aggressive Dragon threat.
- 1 Lathliss, Dragon Queen | threat | Dragon-tribal board-building threat.
- 1 Scourge of the Throne | threat | High-impact combat threat.
- 1 Terror of Mount Velus | threat | High-impact Dragon combat threat.
- 1 Thrakkus the Butcher | threat | Dragon-tribal combat threat.
- 1 Thunderbreak Regent | threat | Reliable Dragon threat.
- 1 Twinflame Tyrant | threat | High-impact Dragon combat threat.
- 1 Utvara Hellkite | threat | Dragon-tribal late-game threat.
- 1 Balefire Dragon | wipe | Dragon-based battlefield reset.
- 1 Draconic Intervention | wipe | Flexible sweeping removal.
- 1 Incinerator of the Guilty | wipe | Dragon-based sweeping threat.

</details>

### 15. delegated commander, owned first

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 229 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $125.39 the whole deck.

**Summary:** Astarion leads a black-white lifegain deck that develops its board through dedicated synergy pieces, then turns to Angels and other threats to pressure the table. It supports that plan with steady cards, mana development, targeted answers, and reset buttons when the board gets away from it. The deck gives up explosive speed for a more measured game built around maintaining its lifegain focus and deploying threats over time.

- [INFO] `bracket_prose_rules`: bracket 2 (Core, 8+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [WARN] `not_owned`: Astarion, the Decadent: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 2.98 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Buster Sword | draw | Included as a draw card.
- 1 Exemplar of Light | draw | Included as a draw card.
- 1 Inspiring Overseer | draw | Included as a draw card.
- 1 Instant Ramen | draw | Included as a draw card.
- 1 Lembas | draw | Included as a draw card.
- 1 Mask of Memory | draw | Included as a draw card.
- 1 Night's Whisper | draw | Included as a draw card.
- 1 Skullclamp | draw | Included as a draw card.
- 1 Tome of Legends | draw | Included as a draw card.
- 1 Wall of Omens | draw | Included as a draw card.
- 1 Bastion Protector | interaction | Included as an interaction card.
- 1 Champion's Helm | interaction | Included as an interaction card.
- 1 Clever Concealment | interaction | Included as an interaction card.
- 1 Darksteel Plate | interaction | Included as an interaction card.
- 1 Lightning Greaves | interaction | Included as an interaction card.
- 1 Swiftfoot Boots | interaction | Included as an interaction card.
- 20 Plains | land | Basic land slots for the deck's white base.
- 16 Swamp | land | Basic land slots for the deck's black base.
- 1 Arcane Signet | ramp | Included as a ramp card.
- 1 Astral Cornucopia | ramp | Included as a ramp card.
- 1 Bender's Waterskin | ramp | Included as a ramp card.
- 1 Blitzball | ramp | Included as a ramp card.
- 1 Brass Infiniscope | ramp | Included as a ramp card.
- 1 Commander's Sphere | ramp | Included as a ramp card.
- 1 Fellwar Stone | ramp | Included as a ramp card.
- 1 Sol Ring | ramp | Included as a ramp card.
- 1 Thought Vessel | ramp | Included as a ramp card.
- 1 Wayfarer's Bauble | ramp | Included as a ramp card.
- 1 Banishing Light | removal | Included as a removal card.
- 1 Bitter Triumph | removal | Included as a removal card.
- 1 Crib Swap | removal | Included as a removal card.
- 1 Dispatch | removal | Included as a removal card.
- 1 Fatal Push | removal | Included as a removal card.
- 1 Generous Gift | removal | Included as a removal card.
- 1 Infernal Grasp | removal | Included as a removal card.
- 1 Swords to Plowshares | removal | Included as a removal card.
- 1 Aerith Gainsborough | synergy | Included as a synergy card for the lifegain plan.
- 1 Aettir and Priwen | synergy | Included as a synergy card for the lifegain plan.
- 1 Angel of Vitality | synergy | Included as a synergy card for the lifegain plan.
- 1 Compassionate Healer | synergy | Included as a synergy card for the lifegain plan.
- 1 Elixir | synergy | Included as a synergy card for the lifegain plan.
- 1 Excalibur II | synergy | Included as a synergy card for the lifegain plan.
- 1 Kor Firewalker | synergy | Included as a synergy card for the lifegain plan.
- 1 Light of Promise | synergy | Included as a synergy card for the lifegain plan.
- 1 Night Nurse, Healer of Heroes | synergy | Included as a synergy card for the lifegain plan.
- 1 Prideful Feastling | synergy | Included as a synergy card for the lifegain plan.
- 1 Rosie Cotton of South Lane | synergy | Included as a synergy card for the lifegain plan.
- 1 Second Breakfast | synergy | Included as a synergy card for the lifegain plan.
- 1 Well-Worn Spatula | synergy | Included as a synergy card for the lifegain plan.
- 1 White Mage's Staff | synergy | Included as a synergy card for the lifegain plan.
- 1 Angel of Invention | threat | Included as a threat.
- 1 Bill the Pony | threat | Included as a threat.
- 1 Dawnhand Eulogist | threat | Included as a threat.
- 1 Gollum, Silent Slinker // Meager Meal | threat | Included as a threat.
- 1 Invisible Woman, Sue Storm | threat | Included as a threat.
- 1 Lo and Li, Twin Tutors | threat | Included as a threat.
- 1 Lyra Dawnbringer | threat | Included as a threat.
- 1 Minwu, White Mage | threat | Included as a threat.
- 1 Reaping Willow | threat | Included as a threat.
- 1 Shattered Angel | threat | Included as a threat.
- 1 Sneering Shadewriter | threat | Included as a threat.
- 1 Victory's Herald | threat | Included as a threat.
- 1 Austere Command | wipe | Included as a wipe.
- 1 Fumigate | wipe | Included as a wipe.
- 1 Vanquish the Horde | wipe | Included as a wipe.

</details>

### 16. a tight budget, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 229 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $16.51 to buy, $38.23 the whole deck.

**Summary:** This deck builds an aristocrats battlefield around Denethor, Ruling Steward, using a broad base of sacrifice-oriented synergy cards, draw, and inexpensive threats to keep pressure on the table. It aims to win through its creature-based plan while removal, interaction, and sweepers keep opposing boards from taking over. The deck gives up premium upgrades and costly individual power cards in favor of cards already in the library and a small set of focused, low-cost additions.

- [INFO] `bracket_prose_rules`: bracket 2 (Core, 8+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [WARN] `not_owned`: Blood Artist: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Elas il-Kor, Sadistic Pilgrim: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Falkenrath Noble: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vindictive Vampire: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Zulaport Cutthroat: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ayli, Eternal Pilgrim: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Baron Bertram Graywater: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bartolomé del Presidio: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Erebos, Bleak-Hearted: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ghoulcaller Gisa: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Lord Skitter's Butcher: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Old Flitterfang: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Woe Strider: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.13 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Avengers Tower | land | Owned land for the mana base.
- 1 Baxter Building | land | Owned land for the mana base.
- 1 Big Apple, 3 a.m. | land | Owned land for the mana base.
- 1 Capital City | land | Owned land for the mana base.
- 1 Castle Doom | land | Owned land for the mana base.
- 1 Command Tower | land | Owned land for the mana base.
- 1 Crossroads Village | land | Owned land for the mana base.
- 1 Eclipsed Realms | land | Owned land for the mana base.
- 1 Exotic Orchard | land | Owned land for the mana base.
- 1 Fisk Tower | land | Owned land for the mana base.
- 1 Forum of Amity | land | Owned land for the mana base.
- 1 Grand Coliseum | land | Owned land for the mana base.
- 1 Great Hall of the Citadel | land | Owned land for the mana base.
- 1 Ifnir Deadlands | land | Owned land for the mana base.
- 1 Insomnia, Crown City | land | Owned land for the mana base.
- 1 Jasmine Dragon Tea Shop | land | Owned land for the mana base.
- 1 Misty Palms Oasis | land | Owned land for the mana base.
- 1 Opal Palace | land | Owned land for the mana base.
- 1 Path of Ancestry | land | Owned land for the mana base.
- 1 Rumble Arena | land | Owned land for the mana base.
- 1 Scavenger Grounds | land | Owned land for the mana base.
- 1 Secluded Courtyard | land | Owned land for the mana base.
- 1 Sidequest: Catch a Fish // Cooking Campsite | land | Owned land for the mana base.
- 1 Spire of Industry | land | Owned land for the mana base.
- 1 Study Hall | land | Owned land for the mana base.
- 1 Surveillance Room | land | Owned land for the mana base.
- 5 Swamp | land | Owned basic land for the mana base.
- 5 Plains | land | Owned basic land for the mana base.
- 1 Arcane Signet | ramp | Owned ramp piece.
- 1 Astral Cornucopia | ramp | Owned ramp piece.
- 1 Chromatic Lantern | ramp | Owned ramp piece.
- 1 Commander's Sphere | ramp | Owned ramp piece.
- 1 Deadly Dispute | ramp | Owned ramp option.
- 1 Fellwar Stone | ramp | Owned ramp piece.
- 1 Inherited Envelope | ramp | Owned ramp piece.
- 1 Sol Ring | ramp | Owned ramp piece.
- 1 Thought Vessel | ramp | Owned ramp piece.
- 1 White Auracite | ramp | Owned ramp piece.
- 1 Ahriman | draw | Owned draw card.
- 1 Beetle-Headed Merchants | draw | Owned draw card.
- 1 Buzzard-Wasp Colony | draw | Owned draw card.
- 1 Circle of Power | draw | Owned draw card.
- 1 Cirith Ungol Patrol | draw | Owned draw card.
- 1 Grave Venerations | draw | Owned draw card.
- 1 Inspiring Overseer | draw | Owned draw card.
- 1 June, Bounty Hunter | draw | Owned draw card.
- 1 Nasty End | draw | Owned draw card.
- 1 Tome of Legends | draw | Owned draw card.
- 1 Bastion Protector | interaction | Owned interaction for the deck's key permanents.
- 1 Frontline Medic | interaction | Owned interaction card.
- 1 Gift of Immortality | interaction | Owned interaction card.
- 1 Swiftfoot Boots | interaction | Owned interaction for a key creature.
- 1 Take Up the Shield | interaction | Owned interaction card.
- 1 Zack Fair | interaction | Owned interaction card.
- 1 Angel of Serenity | removal | Owned removal card.
- 1 Banishing Light | removal | Owned removal card.
- 1 Bitter Triumph | removal | Owned removal card.
- 1 Claim the Precious | removal | Owned removal card.
- 1 Crib Swap | removal | Owned removal card.
- 1 Destroy Evil | removal | Owned removal card.
- 1 Generous Gift | removal | Owned removal card.
- 1 Infernal Grasp | removal | Owned removal card.
- 1 Archfiend of Ifnir | wipe | Owned board-wipe card.
- 1 Austere Command | wipe | Owned board-wipe card.
- 1 Dusk // Dawn | wipe | Owned board-wipe card.
- 1 Al Bhed Salvagers | synergy | Owned synergy piece for the requested aristocrats plan.
- 1 Gollum the Abandoned | synergy | Owned synergy piece for the requested aristocrats plan.
- 1 Gollum, Patient Plotter | synergy | Owned synergy piece for the requested aristocrats plan.
- 1 Gríma Wormtongue | synergy | Owned synergy piece for the requested aristocrats plan.
- 1 Heirloom Auntie | synergy | Owned synergy piece for the requested aristocrats plan.
- 1 Joo Dee, One of Many | synergy | Owned synergy piece for the requested aristocrats plan.
- 1 Nimble Hobbit | synergy | Owned synergy piece for the requested aristocrats plan.
- 1 Phantom Train | synergy | Owned synergy piece for the requested aristocrats plan.
- 1 Rhovanion Rampager | synergy | Owned synergy piece for the requested aristocrats plan.
- 1 Blood Artist | synergy | Low-cost synergy addition for the requested aristocrats plan.
- 1 Elas il-Kor, Sadistic Pilgrim | synergy | Synergy addition for the requested aristocrats plan.
- 1 Falkenrath Noble | synergy | Low-cost synergy addition for the requested aristocrats plan.
- 1 Vindictive Vampire | synergy | Low-cost synergy addition for the requested aristocrats plan.
- 1 Zulaport Cutthroat | synergy | Synergy addition for the requested aristocrats plan.
- 1 Bill the Pony | threat | Owned threat for the battlefield plan.
- 1 Hei Bai, Spirit of Balance | threat | Owned threat for the battlefield plan.
- 1 Namazu Trader | threat | Owned threat for the battlefield plan.
- 1 Vengeful Villagers | threat | Owned threat for the battlefield plan.
- 1 Ayli, Eternal Pilgrim | threat | Low-cost threat that suits the requested aristocrats plan.
- 1 Baron Bertram Graywater | threat | Low-cost threat that suits the requested aristocrats plan.
- 1 Bartolomé del Presidio | threat | Low-cost threat that suits the requested aristocrats plan.
- 1 Erebos, Bleak-Hearted | threat | Threat addition for the requested aristocrats plan.
- 1 Ghoulcaller Gisa | threat | Threat addition for the requested aristocrats plan.
- 1 Lord Skitter's Butcher | threat | Low-cost threat that suits the requested aristocrats plan.
- 1 Old Flitterfang | threat | Low-cost threat that suits the requested aristocrats plan.
- 1 Woe Strider | threat | Low-cost threat that suits the requested aristocrats plan.

</details>

### 17. upgrade a precon, owned first

Format: Commander. Theme: goblins. Pool: owned_first. Shortlist: 196 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $85.27 to buy, $278.32 the whole deck.

**Summary:** This deck builds a Goblin board while using cheap spells, draw, and bursts of ramp to create forceful turns around Zada. It wins by turning its Goblin synergies and threats into sustained pressure, with removal and wipes available to clear resistance. It gives up broader multicolor options and relies on its creature core staying established for its most powerful turns.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 2.65 over 62 nonland cards

<details><summary>The deck list</summary>

- 1 Ancestors' Aid | other | A retained precon spell that stays with the deck's established plan.
- 1 Arena of Glory | land | A retained precon land that remains part of the mana base.
- 1 Battle Hymn | ramp | A retained precon ramp spell that supports explosive Goblin turns.
- 1 Blasphemous Act | wipe | A retained precon wipe for recovering from crowded boards.
- 1 Boggart Shenanigans | removal | A retained precon Goblin payoff that provides removal.
- 1 Castle Embereth | land | A retained precon land that remains part of the mana base.
- 1 Chaos Warp | removal | A retained precon piece of flexible removal.
- 1 Conspicuous Snoop | synergy | A retained precon Goblin synergy card.
- 1 Crimson Wisps | draw | A retained precon draw spell for keeping turns moving.
- 1 Daring Discovery | other | A retained precon spell that preserves the original deck structure.
- 1 Den of the Bugbear | land | A retained precon land that remains part of the mana base.
- 1 Dragon Fodder | other | A retained precon spell that preserves the Goblin-focused plan.
- 1 Empty the Warrens | other | A retained precon spell that remains part of the storm theme.
- 1 Expedite | draw | A retained precon draw spell that helps the deck keep flowing.
- 1 Faithless Looting | draw | A retained precon draw spell that improves card access.
- 1 Fists of Flame | draw | A retained precon draw spell that supports the spell-heavy plan.
- 1 Forgotten Cave | land | A retained precon land that remains part of the mana base.
- 1 Fountainport | land | A retained precon land that remains part of the mana base.
- 1 Frontline Heroism | other | A retained precon card that preserves the original deck structure.
- 1 Gempalm Incinerator | removal | A retained precon Goblin removal card.
- 1 General Kreat, the Boltbringer | synergy | A retained precon Goblin synergy card.
- 1 Goblin Bushwhacker | synergy | A retained precon Goblin synergy card.
- 1 Goblin Chieftain | synergy | A retained precon Goblin synergy card.
- 1 Goblin Dark-Dwellers | threat | A retained precon Goblin threat.
- 1 Goblin Lackey | synergy | A retained precon Goblin synergy card.
- 1 Goblin Matron | synergy | A retained precon Goblin synergy card.
- 1 Goblin Negotiation | removal | A retained precon Goblin removal spell.
- 1 Goblin Trashmaster | removal | A retained precon Goblin removal card.
- 1 Goblin Warchief | synergy | A retained precon Goblin synergy card.
- 1 Grapeshot | removal | A retained precon removal spell that stays with the storm theme.
- 1 Haze of Rage | other | A retained precon spell that preserves the original storm plan.
- 1 Hidden Volcano | land | A retained precon land that remains part of the mana base.
- 1 Idol of Oblivion | draw | A retained precon draw artifact for card flow.
- 1 Impact Tremors | other | A retained precon card that preserves the Goblin-focused plan.
- 1 Kher Keep | land | A retained precon land that remains part of the mana base.
- 1 Krenko's Command | other | A retained precon spell that preserves the Goblin-focused plan.
- 1 Krenko, Mob Boss | threat | A retained precon Goblin threat.
- 1 Mana Geyser | ramp | A retained precon ramp spell for large turns.
- 1 Mogg War Marshal | synergy | A retained precon Goblin synergy card.
- 1 Past in Flames | other | A retained precon spell that remains part of the storm plan.
- 1 Quest for the Goblin Lord | synergy | A retained precon Goblin synergy card.
- 1 Reliquary Tower | land | A retained precon land that remains part of the mana base.
- 1 Roaming Throne | other | A retained precon card that preserves the original deck structure.
- 1 Ruby Medallion | other | A retained precon card that remains part of the spell-focused plan.
- 1 Sazacap's Brew | draw | A retained precon draw spell for card flow.
- 1 Searslicer Goblin | synergy | A retained precon Goblin synergy card.
- 1 Seething Song | ramp | A retained precon ramp spell for explosive turns.
- 1 Shinka, the Bloodsoaked Keep | land | A retained precon land that remains part of the mana base.
- 1 Siege-Gang Commander | removal | A retained precon Goblin removal card.
- 1 Siege-Gang Lieutenant | removal | A retained precon Goblin removal card.
- 1 Skirk Prospector | ramp | A retained precon Goblin ramp card.
- 1 Skullclamp | draw | A retained precon draw artifact for sustained card flow.
- 1 Sol Ring | ramp | A retained precon ramp artifact.
- 1 Spreading Insurrection | other | A retained precon spell that preserves the original deck structure.
- 1 Storm-Kiln Artist | ramp | A retained precon ramp card for the spell-heavy plan.
- 1 Swiftfoot Boots | interaction | A retained precon interaction piece.
- 1 Throne of Eldraine | ramp | A retained precon ramp artifact.
- 1 Vandalblast | wipe | A retained precon wipe for opposing artifacts.
- 1 War Room | land | A retained precon land that remains part of the mana base.
- 1 Wild Ride | other | A retained precon spell that preserves the original deck structure.
- 1 Witch's Mark | draw | A retained precon draw spell for card flow.
- 1 Ancestral Anger | draw | Added draw that supports the deck's spell-focused turns.
- 1 Arcane Signet | ramp | Added reliable ramp for developing the deck sooner.
- 1 Brightstone Ritual | ramp | Added Goblin-focused ramp for explosive turns.
- 1 Fable of the Mirror-Breaker // Reflection of Kiki-Jiki | ramp | Added ramp that strengthens the deck's resource development.
- 1 Goblin Bombardment | removal | Added Goblin-themed removal that keeps the theme intact.
- 1 Lightning Greaves | interaction | Added interaction to better support the commander and key threats.
- 1 Moggcatcher | threat | Added a Goblin-related threat for stronger board presence.
- 1 Pashalik Mons | removal | Added Goblin-themed removal that reinforces the deck's plan.
- 1 Rundvelt Hordemaster | synergy | Added Goblin synergy for a more cohesive creature core.
- 1 Broadside Bombardiers | removal | Added a Goblin removal card that contributes to board control.
- 1 Mask of Memory | draw | Added draw to help sustain resources through longer games.
- 27 Mountain | land | Basic lands provide the deck's primary mana base.

</details>

### 18. upgrade a precon, any card

Format: Commander. Theme: turtles. Pool: any_card. Shortlist: 300 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for game_changer_limit, precon_share. Block findings: 0.

Cost: $86.00 to buy, $86.00 the whole deck.

**Summary:** A five-color Turtle Power deck centered on the precon’s heroes, allies, gadgets, and mutant mayhem. It develops a broad mana base, builds a character-driven board, and uses versatile answers and sweeping effects to keep the action under control before closing through coordinated Turtle-themed pressure.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [WARN] `land_count`: 39 lands: the guide range for this format is 34 to 38
- [INFO] `curve_summary`: average mana value 3.42 over 60 nonland cards

<details><summary>The deck list</summary>

- 1 Acidic Slime | removal | Flexible permanent removal on a creature body.
- 1 April O'Neil, Live on the Scene | draw | Theme-native card advantage.
- 1 Arcade Cabinet | other | Retains the precon's flavorful support package.
- 1 Arcane Signet | ramp | Reliable multicolor acceleration.
- 1 Ash Barrens | land | Fixes colors while maintaining the land base.
- 1 Assassin's Trophy | removal | Versatile answer to problematic permanents.
- 1 Baxter, Fly in the Ointment | other | Retains the Turtle Power cast and theme.
- 1 Bebop, Skull & Crossbones | threat | Retains the Turtle Power cast and theme.
- 1 Big Apple, 3 a.m. | land | Theme-native color source.
- 1 Big Mother Mouser | other | Retains the precon's artifact package.
- 1 Biogenic Ooze | threat | Board-building threat retained from the precon.
- 1 Blasphemous Act | wipe | Efficient creature-board reset.
- 1 Casey Jones, Back Alley Brute | threat | Retains the Turtle Power cast and theme.
- 1 Chromatic Lantern | ramp | Fixes demanding colors and accelerates mana.
- 1 Cinder Glade | land | Dual-color land for the multicolor mana base.
- 1 Coin of Mastery | other | Retains the precon's flavorful support package.
- 1 Command Tower | land | Reliable five-color land.
- 1 Continue? | other | Retains the precon's thematic spell package.
- 1 Corpsejack Menace | synergy | Supports the deck's counter-oriented elements.
- 1 Cultivate | ramp | Stable color fixing and land acceleration.
- 1 Dimension X Pizzasaur | other | Retains the precon's artifact and character theme.
- 1 Donatello, the Brains | synergy | Theme-native Turtle payoff.
- 1 Double Jump // Flying Kick | other | Retains the precon's thematic spell package.
- 1 Dragonskull Summit | land | Dual-color land for the multicolor mana base.
- 1 Electric Seaweed | other | Retains the precon's defensive support package.
- 1 Endless Foot Assault | other | Retains the precon's thematic support package.
- 1 Escape Tunnel | land | Fixes colors and provides utility.
- 1 Everything Pizza | other | Retains the precon's flavorful support package.
- 1 Evolving Wilds | land | Fixes colors through basic-land access.
- 1 Exploding Barrel | other | Retains the precon's artifact package.
- 1 Fast Forward | other | Retains the precon's thematic spell package.
- 1 Foot Chopper | other | Retains the precon's equipment package.
- 1 Game Over | wincon | Retains the precon's thematic finishing package.
- 1 Grand Coliseum | land | Broad color access for a five-color deck.
- 1 Harmonize | draw | Straightforward card advantage.
- 1 Here Comes a New Hero! | other | Retains the precon's thematic support package.
- 1 Hidden Hideout | land | Theme-native mana source.
- 1 High Score | other | Retains the precon's thematic support package.
- 1 Irma, Part-Time Mutant | synergy | Retains the Turtle Power cast and theme.
- 1 Krang, the All-Powerful | threat | Retains the precon's artifact and character theme.
- 1 Leatherhead, Iron Gator | threat | Retains the Turtle Power cast and theme.
- 1 Leonardo, the Balance | threat | Theme-native Turtle payoff.
- 1 Lessons from Life | other | Retains the precon's thematic support package.
- 1 Level Up | synergy | Retains the precon's thematic support package.
- 1 Michelangelo, the Heart | synergy | Theme-native Turtle payoff.
- 1 Mole Module | other | Retains the precon's artifact package.
- 1 Mona Lisa, Science Geek | other | Retains the Turtle Power cast and theme.
- 1 Ninja Pizza | synergy | Retains the precon's thematic support package.
- 1 Path of Ancestry | land | Color fixing that supports the creature theme.
- 1 Rain-Slicked Copse | land | Theme-native dual-color land.
- 1 Raphael, the Muscle | threat | Theme-native Turtle payoff.
- 1 Rat King, Pale Piper | threat | Retains the Turtle Power cast and theme.
- 1 Ray Fillet, Wave Warrior | threat | Retains the Turtle Power cast and theme.
- 1 Roadkill Rodney | other | Retains the precon's artifact package.
- 1 Rocksteady, Mutant Marauder | threat | Retains the Turtle Power cast and theme.
- 1 Rootbound Crag | land | Dual-color land for the multicolor mana base.
- 1 Shellshock | other | Retains the precon's thematic spell package.
- 1 Shredder, Shadow Master | threat | Retains the Turtle Power cast and theme.
- 1 Sodden Verdure | land | Theme-native dual-color land.
- 1 Sol Ring | ramp | Efficient mana acceleration.
- 1 Special Move | other | Retains the precon's thematic spell package.
- 1 Splinter, the Mentor | synergy | Retains the Turtle Power cast and theme.
- 1 Steelbane Hydra | removal | Creature-based artifact and enchantment removal.
- 1 Super Combo | wincon | Retains the precon's thematic finishing package.
- 1 Swift Demise | other | Retains the precon's thematic spell package.
- 1 Tempestra, Dame of Games | other | Retains the Turtle Power cast and theme.
- 1 Thriving Grove | land | Flexible color fixing.
- 1 Thriving Isle | land | Flexible color fixing.
- 1 Thriving Moor | land | Flexible color fixing.
- 1 Together Forever | synergy | Retains the precon's thematic support package.
- 1 Tokka & Rahzar, Unsupervised | ramp | Theme-native acceleration.
- 1 Turtle Lair | land | Theme-native utility land.
- 1 Undergrowth Stadium | land | Dual-color land for the multicolor mana base.
- 1 Vanquish the Horde | wipe | Creature-board reset retained from the precon.
- 1 Vernal Fen | land | Theme-native dual-color land.
- 1 Vibrant Cityscape | land | Flexible color fixing.
- 1 Vigor | interaction | Protects and strengthens the creature board.
- 1 Voracious Hydra | removal | Scalable creature and removal option.
- 1 Wave Goodbye | wipe | Retains the precon's thematic reset option.
- 1 Swords to Plowshares | removal | Low-cost, reliable creature removal.
- 4 Forest | land | Basic green sources for ramp and Turtle spells.
- 4 Island | land | Basic blue sources for the deck's multicolor needs.
- 4 Plains | land | Basic white sources for removal and support.
- 4 Swamp | land | Basic black sources for the multicolor mana base.
- 3 Mountain | land | Basic red sources for the multicolor mana base.

</details>

### 19. the Hobbit family, two colours

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 174 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $355.75 to buy, $355.75 the whole deck.

**Summary:** This deck builds a broad Elf-centered board under Thranduil, using mana acceleration and steady card flow to keep deploying creatures while holding up answers for opposing threats. It aims to win by turning its developed battlefield into sustained combat pressure, with larger woodland creatures providing extra closing power. The tradeoff is that the deck leans heavily on keeping creatures in play, so repeated board resets can slow its momentum and make rebuilding the board its main challenge.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 3.38 over 63 nonland cards

<details><summary>The deck list</summary>

- 16 Forest | land | Provides a dependable green mana base for the Elf-focused core.
- 10 Island | land | Provides blue mana for the deck's card advantage and interaction.
- 10 Swamp | land | Provides black mana for removal and broader support.
- 1 Arcane Signet | ramp | Provides efficient mana development.
- 1 Elven Chorus | ramp | Supports mana development alongside the Elf plan.
- 1 Elvish Archdruid | ramp | An Elf that advances the deck's mana plan.
- 1 Elvish Mystic | ramp | An early Elf for mana development.
- 1 Mox Amber | ramp | Provides additional mana acceleration for the legendary commander.
- 1 Necklace of Girion | ramp | Adds another mana-development piece.
- 1 Thranduil the Strategist | ramp | An Elf legend that advances the mana plan.
- 1 Thranduil's Company | ramp | Supports the Elf-focused mana development plan.
- 1 Through the Forest Gate | ramp | Provides additional mana development.
- 1 Wood Elves | ramp | An Elf that supports mana development.
- 1 Elvish Visionary | draw | An Elf that contributes to the deck's card flow.
- 1 Fateful Discovery | draw | Provides a card-advantage option.
- 1 Hithlain Knots | draw | Provides card flow when resources are needed.
- 1 Ithilien Kingfisher | draw | A creature that contributes to card advantage.
- 1 Key to the Side-Door | draw | Adds a repeatable-looking card-flow resource to the deck.
- 1 Lórien Revealed | draw | Provides card selection and card flow.
- 1 Night's Whisper | draw | A compact card-advantage spell.
- 1 Old Fat Spider | draw | A creature that contributes to card advantage.
- 1 Palantír of Orthanc | draw | Provides a dedicated card-advantage engine.
- 1 Plunder the Trollshaws | draw | Provides another card-flow spell.
- 1 Confusticate and Bebother | interaction | Provides flexible interaction against opposing plays.
- 1 Elrond, Moon-Reader | interaction | An Elf legend that supplies interaction.
- 1 Mithril Coat | interaction | Helps protect an important permanent.
- 1 My Precious // Allure of Power | interaction | Provides a versatile interaction slot.
- 1 Sound the Trumpets | interaction | Adds a reactive answer to opposing plays.
- 1 Stern Scolding | interaction | Provides efficient stack interaction.
- 1 Bilbo's Deadly Slice | removal | Provides a targeted answer to opposing threats.
- 1 Bitter Downfall | removal | A direct removal option.
- 1 Colossal Whale | removal | A creature-based removal option.
- 1 Crude Bent Blade | removal | Provides removal from an equipment slot.
- 1 Enchanted River's Grasp | removal | Provides permanent-based removal.
- 1 Giant's Boulder | removal | Adds another answer to opposing permanents.
- 1 Orcish Bowmasters | removal | A creature that provides a removal option.
- 1 Quarrel | removal | Provides a compact targeted answer.
- 1 Gnashing of Teeth | wipe | Provides a reset when opposing boards become overwhelming.
- 1 Languish | wipe | Provides a broad battlefield reset.
- 1 Raise the Palisade | wipe | Provides another board-clearing option.
- 1 Arwen, Weaver of Hope | synergy | An Elf legend that reinforces the deck's Elf theme.
- 1 Boughside Wanderers | synergy | An Elf creature for the tribal core.
- 1 Cantankerous Keepers | synergy | An Elf creature that builds the tribal board.
- 1 Celeborn the Wise | synergy | An Elf legend that supports the deck's theme.
- 1 Elven Raft-Steerer | synergy | An Elf creature for the tribal core.
- 1 Galadhrim Guide | synergy | An Elf that supports the deck's creature base.
- 1 Galion, Elvenking's Butler | synergy | An Elf legend that fits Thranduil's court.
- 1 Guardian of the Halls | synergy | An Elf creature that strengthens the tribal board.
- 1 Lothlórien Lookout | synergy | An Elf creature for the deck's central theme.
- 1 Mirkwood Meditator | synergy | An Elf creature that supports the tribal plan.
- 1 Mirkwood Nurturer | synergy | An Elf creature that builds the themed board.
- 1 Mirkwood Pathmaker | synergy | An Elf creature for the deck's tribal core.
- 1 Nimrodel Watcher | synergy | An Elf that contributes to the deck's Elf density.
- 1 Supper for Spiders | synergy | Supports the deck's coordinated creature plan.
- 1 Attercop | threat | A substantial creature threat for the battlefield.
- 1 Dreaded Bat-Cloud | threat | Adds another creature threat to pressure opponents.
- 1 Gigantic Big Bear | threat | A large creature threat for closing games through combat.
- 1 Great Fierce Bee | threat | Provides an additional creature threat.
- 1 Large Bear | threat | A straightforward combat threat.
- 1 Little Bear | threat | Adds to the deck's creature pressure.
- 1 Mirkwood Elk | threat | A creature threat that fits the woodland flavor of the deck.
- 1 Nasty Little Rabbit | threat | Provides another body for applying pressure.
- 1 Ordinary Bear | threat | A creature threat for the combat plan.
- 1 Rhovanion Rampager | threat | A creature threat that helps maintain battlefield pressure.
- 1 Willow-Wind | threat | An additional creature threat for the late game.
- 1 Wilderland Scrounger | threat | A creature threat that rounds out the attacking force.

</details>

### 20. the Hobbit family, a delegated commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 77 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $528.48 to buy, $528.48 the whole deck.

**Summary:** Smaug leads a dragon-focused build that uses plentiful ramp to establish its larger threats, while artifact-based draw and a broad removal package help it keep moving through a crowded table. It wins by applying pressure with Smaug, its Dragons, and its supporting creature force, backed by sweepers when the board gets out of hand. The deck gives up deeper dedicated Dragon synergy and a wider range of reactive tools in exchange for a heavily Mountain-based, theme-forward game plan.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 3.17 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Dragon-Cursed Halls | land | Provides a land slot from the shortlist.
- 1 Rogue's Passage | land | Provides a land slot from the shortlist.
- 1 The Lonely Mountain | land | Provides a land slot from the shortlist.
- 1 Treasure Vault | land | Provides a land slot from the shortlist.
- 32 Mountain | land | Supplies the basic-land foundation for the deck.
- 1 Arcane Signet | ramp | Provides the listed ramp support.
- 1 Bag End Banquet | ramp | Provides the listed ramp support.
- 1 Burn, Burn, Tree and Fern | ramp | Provides the listed ramp support.
- 1 Cavern-Hoard Dragon | ramp | Combines a Dragon body with the listed ramp role.
- 1 Dragon's Desire | ramp | Provides the listed ramp support.
- 1 Fíli and Kíli, Joyous | ramp | Provides the listed ramp support.
- 1 Long-Bodied Grey Dog | ramp | Provides the listed ramp support.
- 1 Mox Amber | ramp | Provides the listed ramp support.
- 1 Orcrist, Goblin-cleaver | ramp | Provides the listed ramp support.
- 1 The Misty Mountains Cold | ramp | Provides the listed ramp support.
- 1 The Reaver Cleaver | ramp | Provides the listed ramp support.
- 1 Thorin, Company's Leader | ramp | Provides the listed ramp support.
- 1 Troop of Ponies | ramp | Provides the listed ramp support.
- 1 Wayfarer's Bauble | ramp | Provides the listed ramp support.
- 1 Balin, Loremaster | draw | Provides the listed card-draw support.
- 1 Key to the Side-Door | draw | Provides the listed card-draw support.
- 1 Palantír of Orthanc | draw | Provides the listed card-draw support.
- 1 Ragged Short Spear | draw | Provides the listed card-draw support.
- 1 Thrór's Map | draw | Provides the listed card-draw support.
- 1 Óin the Brave | draw | Provides the listed card-draw support.
- 1 Bilbo's Ring | interaction | Provides the listed interaction support.
- 1 Dwarven Mattock | interaction | Provides the listed interaction support.
- 1 Mithril Coat | interaction | Provides the listed interaction support.
- 1 The One Ring | interaction | Provides the listed interaction support.
- 1 Battle-Scarred Goblin | removal | Provides a removal option from the shortlist.
- 1 Fire of Orthanc | removal | Provides a removal option from the shortlist.
- 1 Giant's Boulder | removal | Provides a removal option from the shortlist.
- 1 Goblin Cratermaker | removal | Provides a removal option from the shortlist.
- 1 Improvised Club | removal | Provides a removal option from the shortlist.
- 1 Inferno Titan | removal | Pairs a creature body with the listed removal role.
- 1 Smaug, the Great Calamity // Spew Flame | removal | Adds a second Smaug card with the listed removal role.
- 1 The Black Arrow | removal | Provides a removal option from the shortlist.
- 1 Call Forth the Tempest | wipe | Provides a board-wipe option.
- 1 Desolation of Smaug | wipe | Provides a board-wipe option tied to Smaug's theme.
- 1 Glóin the Mighty // Easy Pickings | wipe | Provides a board-wipe option.
- 1 Last Light of Durin's Day | synergy | Supplies the shortlist's dedicated synergy piece.
- 1 Desert Were-Worm | threat | Provides a Dragon threat from the shortlist.
- 1 Andúril, Flame of the West | other | Fills the shortlist's other support role.
- 1 Andúril, Narsil Reforged | other | Fills the shortlist's other support role.
- 1 Bombur, Gentle Dreamer | other | Fills the shortlist's other support role.
- 1 Bothersome Noisemaker | other | Fills the shortlist's other support role.
- 1 Dori, Bearer of Friends | other | Fills the shortlist's other support role.
- 1 Dwarven Mauler | other | Fills the shortlist's other support role.
- 1 Dwarven Warriors | other | Fills the shortlist's other support role.
- 1 Dáin Ironfoot | other | Fills the shortlist's other support role.
- 1 Gandalf, Goblins' Bane // Flameshape | other | Fills the shortlist's other support role.
- 1 Getaway Barrel | other | Fills the shortlist's other support role.
- 1 Glamdring | other | Fills the shortlist's other support role.
- 1 Goblin-town Flunkies | other | Fills the shortlist's other support role.
- 1 Gundabad Opportunist | other | Fills the shortlist's other support role.
- 1 Guttersnipe | other | Fills the shortlist's other support role.
- 1 Iron Hills Stalwart | other | Fills the shortlist's other support role.
- 1 Long-Lost Lances | other | Fills the shortlist's other support role.
- 1 Misty Mountains Raider | other | Fills the shortlist's other support role.
- 1 Old Thrush | other | Fills the shortlist's other support role.
- 1 Oliphaunt | other | Fills the shortlist's other support role.
- 1 Olog-hai Crusher | other | Fills the shortlist's other support role.
- 1 Orcish Siegemaster | other | Fills the shortlist's other support role.
- 1 Smaug's Fury | other | Fills the shortlist's other support role.
- 1 Snowslope Hunter | other | Fills the shortlist's other support role.
- 1 Sting, Bilbo's Sword | other | Fills the shortlist's other support role.
- 1 Tidings of War | other | Fills the shortlist's other support role.
- 1 Well-Worn Spatula | other | Fills the shortlist's other support role.

</details>

### 21. the Hobbit family, mana from outside

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 156 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $595.50 to buy, $595.50 the whole deck.

**Summary:** Smaug leads a proactive, board-focused deck that develops its mana, keeps cards flowing, and builds pressure with Dragons, Goblins, Orcs, large creatures, and Equipment. It wins by establishing a threatening board and using its removal and board-reset cards to clear the way for its major creatures. The deck gives up some flexibility for its concentrated creature, Equipment, and Middle-earth thematic packages.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 3.56 over 63 nonland cards
- [WARN] `outside_requested_set`: the sets you named do not hold 4 cards: Blood Crypt, Bojuka Bog, Command Tower, Dragonskull Summit

<details><summary>The deck list</summary>

- 1 Blood Crypt | land | A black-red land for the mana base.
- 1 Bojuka Bog | land | A utility land slot in the mana base.
- 1 Command Tower | land | A multicolor land for the mana base.
- 1 Dragon-Cursed Halls | land | A Dragon-themed land slot.
- 1 Dragonskull Summit | land | A black-red land for the mana base.
- 1 Goblin-town | land | A Goblin-themed land slot.
- 1 Minas Morgul, Dark Fortress | land | A utility land slot in the mana base.
- 1 Mount Doom | land | A thematic land slot.
- 1 The Lonely Mountain | land | A thematic Mountain land slot.
- 15 Mountain | land | Primary basic land for the mana base.
- 12 Swamp | land | Primary basic land for the mana base.
- 1 Arcane Signet | ramp | A compact mana-development piece.
- 1 Bag End Banquet | ramp | A ramp card from the shortlist.
- 1 Bolg's Company | ramp | A creature-based ramp slot.
- 1 Burn, Burn, Tree and Fern | ramp | A Saga that supports mana development.
- 1 Cavern-Hoard Dragon | ramp | A Dragon that contributes to mana development.
- 1 Dragon's Desire | ramp | A Dragon-themed ramp spell.
- 1 Mox Amber | ramp | A compact legendary mana piece.
- 1 Orcrist, Goblin-cleaver | ramp | An Equipment ramp slot.
- 1 Smaug the Magnificent | ramp | A Smaug-themed creature that develops mana.
- 1 Smaug, Wicked Worm | ramp | A Smaug-themed creature that develops mana.
- 1 Balin, Loremaster | draw | A creature-based card-flow slot.
- 1 Gollum, Riddle Master | draw | A card-flow creature for the deck.
- 1 Key to the Side-Door | draw | An artifact card-flow piece.
- 1 Night's Whisper | draw | A direct card-flow spell.
- 1 Palantír of Orthanc | draw | A legendary artifact card-flow piece.
- 1 Rage into the Valley | draw | A card-flow sorcery.
- 1 Ragged Short Spear | draw | An Equipment card-flow slot.
- 1 Reverent Howl | draw | An instant card-flow slot.
- 1 Thrór's Map | draw | A legendary artifact card-flow piece.
- 1 Óin the Brave | draw | A creature-based card-flow slot.
- 1 Bilbo's Ring | interaction | An Equipment interaction piece.
- 1 Dwarven Mattock | interaction | An Equipment interaction piece.
- 1 Mithril Coat | interaction | A protective Equipment interaction piece.
- 1 My Precious // Allure of Power | interaction | A flexible legendary Equipment interaction slot.
- 1 The One Ring | interaction | A legendary artifact interaction piece.
- 1 Getaway Barrel | interaction | An artifact utility slot for interacting with opposing plans.
- 1 Azog, Moria's Ruin | removal | A legendary creature removal slot.
- 1 Bilbo's Deadly Slice | removal | An instant removal spell.
- 1 Bitter Downfall | removal | A dedicated removal spell.
- 1 Fire of Orthanc | removal | A sorcery removal spell.
- 1 Goblin Cratermaker | removal | A Goblin creature removal slot.
- 1 Orcish Bowmasters | removal | An Orc creature removal slot.
- 1 Smaug, the Great Calamity // Spew Flame | removal | A Smaug-themed removal slot with an Adventure.
- 1 Witch-king, Bringer of Ruin | removal | A legendary creature removal slot.
- 1 Bolg, Erebor's Reckoning | wipe | A creature-based board-reset slot.
- 1 Desolation of Smaug | wipe | A Smaug-themed board-reset spell.
- 1 Languish | wipe | A dedicated board-reset spell.
- 1 Supper for Spiders | synergy | A listed synergy piece for the deck's central plan.
- 1 Along the Crooked Way | synergy | A thematic enchantment supporting the deck's central plan.
- 1 Andúril, Flame of the West | synergy | A legendary Equipment that supports the deck's equipment package.
- 1 Andúril, Narsil Reforged | synergy | A legendary Equipment that supports the deck's equipment package.
- 1 Down, Down to Goblin-town | synergy | A Goblin-themed Saga supporting the deck's creature package.
- 1 Goblin Plate Mail | synergy | An Equipment supporting the Goblin and Equipment packages.
- 1 Great Goblin, Foul-Hearted | synergy | A legendary Goblin supporting the Goblin creature package.
- 1 Gundabad Opportunist | synergy | A Goblin creature supporting the Goblin package.
- 1 Long-Lost Lances | synergy | An Equipment supporting the deck's Equipment package.
- 1 Misty Mountains Raider | synergy | A Goblin creature supporting the creature package.
- 1 Orcish Siegemaster | synergy | An Orc creature supporting the aggressive creature plan.
- 1 Smaug's Fury | synergy | A Smaug-themed instant supporting the central plan.
- 1 Sting, Bilbo's Sword | synergy | A legendary Equipment supporting the Equipment package.
- 1 Well-Worn Spatula | synergy | An Equipment supporting the deck's Equipment package.
- 1 Desert Were-Worm | threat | A large Dragon Wurm body for the threat suite.
- 1 Dreaded Bat-Cloud | threat | An evasive creature threat.
- 1 Great Ugly-Looking Goblin // Clap! Snap! | threat | A Goblin threat with an Adventure option.
- 1 Haunt of the Dead Marshes | threat | A Nightmare creature threat.
- 1 Inferno Titan | threat | A large Giant threat.
- 1 Oliphaunt | threat | A large creature threat.
- 1 Olog-hai Crusher | threat | A Troll creature threat.
- 1 Ravening Warg | threat | A creature threat for the board.
- 1 Rhovanion Rampager | threat | A creature threat for the board.
- 1 Sauron, the Lidless Eye | threat | A legendary creature threat.
- 1 Stone-Giant of High Pass | threat | A Giant creature threat.
- 1 Troll of Khazad-dûm | threat | A Troll creature threat.

</details>

### 22. a set family and a card from outside it

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 175 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $471.89 to buy, $471.89 the whole deck.

**Summary:** This deck builds an Elf-heavy creature board under Thranduil, using ramp to develop quickly and draw to keep creatures and support pieces flowing. It wins by turning that board of creature threats sideways, with Equipment and interaction helping the attack get through. It gives up some flexibility for a focused, creature-based Elf plan and relies on its board presence to carry the game.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 3.51 over 63 nonland cards
- [WARN] `outside_requested_set`: the sets you named do not hold 1 card: Sol Ring

<details><summary>The deck list</summary>

- 1 Dragon-Cursed Halls | land | Included as a land.
- 1 Elven Passage | land | Included as a land.
- 1 Elvenking's Halls | land | Included as a land.
- 1 Hobbit Hole | land | Included as a land.
- 1 Minas Morgul, Dark Fortress | land | Included as a land.
- 1 Mirkwood | land | Included as a land.
- 1 Rivendell | land | Included as a land.
- 1 Rogue's Passage | land | Included as a land.
- 1 The Black Gate | land | Included as a land.
- 1 The Shire | land | Included as a land.
- 1 Treasure Vault | land | Included as a land.
- 12 Forest | land | Basic green land base.
- 6 Island | land | Basic blue land base.
- 7 Swamp | land | Basic black land base.
- 1 Sol Ring | ramp | Kept as requested.
- 1 Arcane Signet | ramp | Provides a ramp slot.
- 1 Delighted Halfling | ramp | Provides a ramp slot.
- 1 Elven Chorus | ramp | Provides a ramp slot for the creature plan.
- 1 Elvish Archdruid | ramp | Provides Elf-focused ramp.
- 1 Elvish Mystic | ramp | Provides early ramp.
- 1 Thranduil's Company | ramp | Provides ramp in the Elf plan.
- 1 Wayfarer's Bauble | ramp | Provides a ramp slot.
- 1 Wood Elves | ramp | Provides ramp while adding an Elf creature.
- 1 Woodland Weavemaster | ramp | Provides ramp while supporting the creature plan.
- 1 Elvish Visionary | draw | Provides draw while adding an Elf creature.
- 1 Fateful Discovery | draw | Provides a draw slot.
- 1 Hithlain Knots | draw | Provides a draw slot.
- 1 Last March of the Ents | draw | Provides a draw slot.
- 1 Lórien Revealed | draw | Provides a draw slot.
- 1 Night's Whisper | draw | Provides a draw slot.
- 1 Palantír of Orthanc | draw | Provides a draw slot.
- 1 Plunder the Trollshaws | draw | Provides a draw slot.
- 1 Thrór's Map | draw | Provides a draw slot.
- 1 Uncover the Moon-Letters | draw | Provides a draw slot.
- 1 Bilbo's Ring | interaction | Provides an interaction slot.
- 1 Elrond, Moon-Reader | interaction | Provides interaction while adding an Elf Noble.
- 1 Mithril Coat | interaction | Provides an interaction slot.
- 1 Stern Scolding | interaction | Provides an interaction slot.
- 1 The One Ring | interaction | Provides an interaction slot.
- 1 Thranduil's Decree | interaction | Provides an interaction slot.
- 1 Bilbo's Deadly Slice | removal | Provides a removal slot.
- 1 Bitter Downfall | removal | Provides a removal slot.
- 1 Crude Bent Blade | removal | Provides removal from an Equipment slot.
- 1 Enchanted River's Grasp | removal | Provides a removal slot.
- 1 Orcish Bowmasters | removal | Provides removal on a creature.
- 1 Quarrel | removal | Provides a removal slot.
- 1 Uneasy Partings | removal | Provides a removal slot.
- 1 Witch-king of Angmar | removal | Provides removal on a legendary creature.
- 1 Gnashing of Teeth | wipe | Provides a wipe slot.
- 1 Languish | wipe | Provides a wipe slot.
- 1 Raise the Palisade | wipe | Provides a wipe slot.
- 1 Arwen, Weaver of Hope | synergy | Elf creature supporting the deck's Elf synergy.
- 1 Boughside Wanderers | synergy | Elf creature supporting the deck's Elf synergy.
- 1 Cantankerous Keepers | synergy | Elf creature supporting the deck's Elf synergy.
- 1 Celeborn the Wise | synergy | Elf Noble supporting the deck's Elf synergy.
- 1 Galadhrim Guide | synergy | Elf creature supporting the deck's Elf synergy.
- 1 Galion, Elvenking's Butler | synergy | Elf Advisor supporting the deck's Elf synergy.
- 1 Guardian of the Halls | synergy | Elf creature supporting the deck's Elf synergy.
- 1 Lothlórien Lookout | synergy | Elf creature supporting the deck's Elf synergy.
- 1 Mirkwood Meditator | synergy | Elf creature supporting the deck's Elf synergy.
- 1 Mirkwood Nurturer | synergy | Elf creature supporting the deck's Elf synergy.
- 1 Mirkwood Pathmaker | synergy | Elf creature supporting the deck's Elf synergy.
- 1 Nimrodel Watcher | synergy | Elf creature supporting the deck's Elf synergy.
- 1 Supper for Spiders | synergy | Dedicated synergy piece for the creature plan.
- 1 Thranduil, Sindarin Liege // Silvan Rally | synergy | Elf Noble supporting the deck's Elf synergy.
- 1 Attercop | threat | Creature threat for the combat plan.
- 1 Chief of the Wilds | threat | Legendary Wolf threat for the combat plan.
- 1 Gigantic Big Bear | threat | Creature threat for the combat plan.
- 1 Great Fierce Bee | threat | Creature threat for the combat plan.
- 1 Large Bear | threat | Creature threat for the combat plan.
- 1 Little Bear | threat | Creature threat for the combat plan.
- 1 Mirkwood Elk | threat | Creature threat for the combat plan.
- 1 Ordinary Bear | threat | Creature threat for the combat plan.
- 1 The Chief Warg | threat | Legendary Wolf threat for the combat plan.
- 1 The Lord of the Eagles | threat | Legendary Bird threat for the combat plan.
- 1 Troll of Khazad-dûm | threat | Creature threat for the combat plan.
- 1 Willow-Wind | threat | Creature threat for the combat plan.

</details>

### 23. two set families at once

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 203 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $372.44 to buy, $372.44 the whole deck.

**Summary:** Kíli leads a white creature-and-artifact deck built around Dwarves, legendary Hobbit equipment, and a broad Bloomburrow creature cast. Develop mana, keep cards flowing, and establish a battlefield of creatures backed by equipment before pressing combat with legendary and high-impact threats. The deck can answer key opposing pieces and reset crowded boards, but it gives up a narrow single-tribe focus for a wider collection of thematic creatures and artifacts.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 3.05 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Castle Ardenvale | land | A white mana source for the deck.
- 1 Command Tower | land | A flexible mana source for Kíli's deck.
- 1 Dragon-Cursed Halls | land | A mana source from the available land suite.
- 1 Elven Passage | land | A mana source from the available land suite.
- 1 Evolving Wilds | land | A land slot that helps assemble the mana base.
- 1 Exotic Orchard | land | A flexible mana source.
- 1 Fabled Passage | land | A land slot that helps assemble the mana base.
- 1 Fountainport | land | A mana source that fits the artifact-heavy shell.
- 1 Hidden Grotto | land | A mana source from the available land suite.
- 1 Hobbit Hole | land | A thematic mana source for the deck.
- 1 Lupinflower Village | land | A thematic mana source for the deck.
- 1 Minas Tirith | land | A thematic mana source for the deck.
- 1 Path of Ancestry | land | A mana source that supports the creature base.
- 1 Reliquary Tower | land | A utility land for a deck with substantial card flow.
- 1 Rogue's Passage | land | A utility land that supports combat finishes.
- 1 Swarmyard | land | A utility land that fits several creature types in the deck.
- 1 Three Tree City | land | A thematic mana source for the creature-focused deck.
- 1 Thriving Heath | land | A white mana source.
- 1 Treasure Vault | land | An artifact land that fits the deck's equipment and artifact theme.
- 1 Uncharted Haven | land | A flexible mana source.
- 1 Terramorphic Expanse | land | A land slot that helps assemble the mana base.
- 15 Plains | land | Reliable basic white mana for the deck.
- 1 Arcane Signet | ramp | Efficient artifact mana for developing the board.
- 1 Bag End Banquet | ramp | Thematic mana development from the available pool.
- 1 Burnished Hart | ramp | Creature-based mana development that fits the artifact shell.
- 1 Fellwar Stone | ramp | Low-cost artifact mana.
- 1 Gilded Lotus | ramp | A larger mana boost for the deck's costly plays.
- 1 Hedron Archive | ramp | Artifact mana that also remains useful later.
- 1 Mind Stone | ramp | Early artifact mana with late-game utility.
- 1 Mox Amber | ramp | Legendary-themed artifact mana.
- 1 Ornithopter of Paradise | ramp | Creature-based mana development.
- 1 Sol Ring | ramp | Efficient artifact mana for accelerating the deck.
- 1 Caretaker's Talent | draw | A lasting source of cards for the creature plan.
- 1 Circuit Mender | draw | Artifact-based card flow that fits the deck's shell.
- 1 Idol of Oblivion | draw | An artifact card-flow piece for the deck.
- 1 Jacked Rabbit | draw | A Bloomburrow creature that contributes card flow.
- 1 Jolly Gerbils | draw | A thematic creature that contributes card flow.
- 1 Mentor of the Meek | draw | Card flow tied to the deck's creature development.
- 1 Skullclamp | draw | An equipment-based card-flow tool.
- 1 Spirited Companion | draw | A low-cost creature that contributes a card.
- 1 The Gaffer | draw | A Hobbit-themed card-flow creature.
- 1 The Queen of Dale | draw | A legendary card-flow creature for the deck.
- 1 Bofur, Reliable Guardian // Concerted Care | interaction | A thematic flexible answer that protects the board plan.
- 1 Galadriel's Dismissal | interaction | A flexible instant-speed disruption piece.
- 1 Luminous Broodmoth | interaction | Creature-based protection for the board.
- 1 Mithril Coat | interaction | Equipment protection for an important creature.
- 1 Swiftfoot Boots | interaction | Equipment protection for Kíli or a key threat.
- 1 Stone by Sunlight | interaction | An instant-speed disruption slot.
- 1 Angel of the Ruins | removal | A creature-based answer that fits the artifact-oriented deck.
- 1 Banishing Light | removal | A broad permanent answer.
- 1 Generous Gift | removal | A flexible answer to problematic permanents.
- 1 Loran of the Third Path | removal | A thematic creature that answers opposing pieces.
- 1 Skyclave Apparition | removal | Creature-based permanent removal.
- 1 Swords to Plowshares | removal | An efficient creature answer.
- 1 The Black Arrow | removal | A thematic equipment-based removal piece.
- 1 Westfold Rider | removal | A thematic creature answer.
- 1 Dusk // Dawn | wipe | A reset option that suits the creature-focused build.
- 1 Martial Coup | wipe | A board reset that also supports rebuilding with creatures.
- 1 Promise of Loyalty | wipe | A reset option for crowded boards.
- 1 Academy Manufactor | synergy | An artifact creature that reinforces the deck's artifact theme.
- 1 Blade Splicer | synergy | A creature-and-artifact piece for the deck's core shell.
- 1 Carrot Cake | synergy | A Bloomburrow artifact that supports the deck's board plan.
- 1 Dáin, Lord of the Iron Hills | synergy | A Dwarf legend that reinforces the deck's Dwarf theme.
- 1 Dwarven Provisioner | synergy | A thematic Dwarf that strengthens the Kíli shell.
- 1 Dwarven Shortsword | synergy | A thematic equipment for the creature plan.
- 1 Glamdring | synergy | A legendary equipment that fits the Hobbit artifact package.
- 1 Helm of the Host | synergy | A high-impact equipment for a creature-centered strategy.
- 1 Iron Hills Blacksmith | synergy | A Dwarf Artificer that links the deck's main themes.
- 1 Maskwood Nexus | synergy | An artifact that ties together the deck's varied creature types.
- 1 Ori, Keeper of Songs | synergy | A Dwarf legend that deepens the thematic creature package.
- 1 Tangle Tumbler | synergy | An artifact vehicle that supports the deck's artifact theme.
- 1 Three Tree Mascot | synergy | A creature artifact that supports the varied creature shell.
- 1 Well-Worn Spatula | synergy | An equipment that supports the deck's artifact package.
- 1 Andúril, Flame of the West | threat | A legendary equipment that serves as a major combat payoff.
- 1 Andúril, Narsil Reforged | threat | A legendary equipment that adds another combat payoff.
- 1 Bilbo, Unexpected Adventurer | threat | A legendary Hobbit creature that advances the attacking plan.
- 1 Eagle of the Great Shelf | threat | A creature threat for pressuring opponents in combat.
- 1 Eagles of the North | threat | A creature threat that adds battlefield presence.
- 1 Fíli the Pathfinder | threat | A thematic legendary Dwarf threat alongside Kíli.
- 1 Jazal Goldmane | threat | A legendary creature that serves as a combat payoff.
- 1 Karn, the Great Creator | threat | A resilient noncreature threat from the artifact package.
- 1 Landroval, Horizon Witness | threat | A legendary Bird that adds a substantial threat.
- 1 Serra Redeemer | threat | A creature threat that adds pressure to the board.
- 1 Sunscorch Regent | threat | A large creature threat for closing games.
- 1 Warren Warleader | threat | A Bloomburrow creature that supports combat finishes.

</details>

### 24. a 60-card deck from one set

Format: Modern. Theme: aggro. Pool: any_card. Shortlist: 96 names.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $46.80 to buy, $46.80 the whole deck.

**Summary:** This mono-red aggro deck presses the table with a steady stream of creature threats, using Emberheart Challenger and Hearthborn Battler to reinforce its aggressive creature plan. It wins by keeping pressure on from the opening turns while Agate Assault and Blooming Blast clear away opposing obstacles, with Artist's Talent and Might of the Meek helping it keep moving into the later game. Its tradeoff is a narrow, creature-forward approach that has little room for broad answers or defensive pivots.

- [INFO] `curve_summary`: average mana value 3.11 over 36 nonland cards

<details><summary>The deck list</summary>

- 24 Mountain | land | Provides the mono-red mana base.
- 4 Artist's Talent | draw | Supplies a draw package that supports the aggressive plan.
- 2 Might of the Meek | draw | Adds efficient draw support alongside the creature-heavy plan.
- 2 Patchwork Banner | ramp | Provides the requested ramp while fitting the creature-focused build.
- 4 Agate Assault | removal | Forms the primary removal package.
- 2 Blooming Blast | removal | Rounds out the removal suite with additional red interaction.
- 4 Emberheart Challenger | synergy | A core synergy creature for the red aggro strategy.
- 4 Hearthborn Battler | synergy | Completes the synergy package with another aggressive creature.
- 4 Frilled Sparkshooter | threat | A primary creature threat for applying pressure.
- 4 Reptilian Recruiter | threat | Adds another full set of proactive threats.
- 4 Teapot Slinger | threat | Helps maintain a dense lineup of attacking threats.
- 2 Stormsplitter | threat | Finishes the threat suite with additional impactful creatures.

</details>

