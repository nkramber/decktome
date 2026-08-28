# PR-8 deck gate

Run date: 2026-08-28. Card snapshot: 2026-08-24.

Verdict: PASS. 18 of 18 decks passed every block check, 0 invented names reached the user, and 0 summaries stated a false rule of the game. All three bars are zero tolerance.

## Summary

| Measure | Value |
|---|---|
| Prompts | 18 |
| Decks returned | 18 |
| Decks with no block finding | 18 |
| Invented names that reached the user | 0 |
| Decks that needed the repair turn | 0 |
| Summaries judged (F-26) | 18 |
| Summaries that state a rule of the game | 0 |
| Summaries that state a FALSE rule | 0 |
| Errors | 0 |
| Prompt version | 6 |
| Calls | 36 |
| Cost | $1.0066 |
| Time | 710 seconds |

## Findings by code

A block stops the deck. A warning and a note are reports.

| Code | Count |
|---|---|
| `not_owned` | 24 |
| `curve_summary` | 18 |
| `bracket_prose_rules` | 12 |
| `land_count` | 2 |
| `basics_added` | 1 |

By severity: BLOCK 0. WARN 26. INFO 31. 

## Decks

### 1. lifegain Commander, any card

Format: Commander. Theme: lifegain. Pool: any_card. Shortlist: 303 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $373.08 to buy, $373.08 the whole deck.

**Summary:** This deck builds around Karlov of the Ghost Council and repeated lifegain, using that theme to grow creatures, create board pressure, and turn favorable life totals into a finishing advantage. It wins through a wide selection of lifegain payoffs and large creature threats, while removal and sweepers keep opposing boards manageable. The tradeoff is a focused strategy: the deck leans heavily on its lifegain engines and permanent-based synergies rather than pursuing a broad range of unrelated lines.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 3.70 over 63 nonland cards

<details><summary>The deck list</summary>

- 18 Plains | land | Provides reliable white mana for the deck's lifegain-focused spells.
- 18 Swamp | land | Provides reliable black mana for the deck's lifegain-focused spells.
- 1 Angel of Indemnity | ramp | Supports the mana plan while remaining part of the deck's creature base.
- 1 Battle Angels of Tyr | ramp | Adds mana acceleration from an evasive threat.
- 1 Beza, the Bounding Spring | ramp | Helps advance mana while contributing to the board.
- 1 Crypt Ghast | ramp | Provides powerful black-mana acceleration.
- 1 Hierophant's Chalice | ramp | Adds steady artifact-based mana support.
- 1 Nuka-Cola Vending Machine | ramp | Supplies artifact-based mana acceleration.
- 1 Orazca Relic | ramp | Helps develop mana resources early.
- 1 Pristine Talisman | ramp | Provides mana acceleration that fits the deck's theme.
- 1 Redemption Choir | ramp | Contributes mana support from a creature slot.
- 1 The Celestus | ramp | Provides flexible artifact-based mana acceleration.
- 1 Alhammarret's Archive | draw | Strengthens the deck's card-advantage package.
- 1 Archivist of Oghma | draw | Provides card advantage on a creature.
- 1 Ayara, First of Locthwain | draw | Adds repeatable card advantage in black.
- 1 Convalescent Care | draw | Turns the deck's theme into card advantage.
- 1 Dawn of Hope | draw | Provides a durable source of card advantage.
- 1 Enduring Innocence | draw | Adds card advantage while occupying a resilient permanent slot.
- 1 Exemplar of Light | draw | Brings card advantage on a lifegain-themed creature.
- 1 Mangara, the Diplomat | draw | Offers creature-based card advantage.
- 1 Markov Purifier | draw | Provides card advantage within the Vampire package.
- 1 Well of Lost Dreams | draw | Converts the deck's central theme into cards.
- 1 Alseid of Life's Bounty | interaction | Protects key creatures and enchantments.
- 1 Faith's Shield | interaction | Protects an important permanent at a critical moment.
- 1 Metropolis Reformer | interaction | Provides protective utility on a lifegain-themed creature.
- 1 Rune-Tail, Kitsune Ascendant // Rune-Tail's Essence | interaction | Rewards the lifegain plan with creature protection.
- 1 Sephara, Sky's Blade | interaction | Offers protection for the creature board.
- 1 Werefox Bodyguard | interaction | Provides flexible creature-based disruption.
- 1 Aetherflux Reservoir | removal | Gives the deck a lifegain-fueled way to remove a problem.
- 1 Ayli, Eternal Pilgrim | removal | Provides repeatable removal from a lifegain-themed creature.
- 1 Cavalier of Night | removal | Adds removal attached to a substantial creature.
- 1 Consuming Corruption | removal | Supplies direct black removal.
- 1 Murderous Rider // Swift End | removal | Offers efficient removal with a creature attached.
- 1 Noxious Gearhulk | removal | Pairs removal with a durable artifact creature.
- 1 Solitude | removal | Provides immediate creature interaction.
- 1 Vona, Butcher of Magan | removal | Adds removal on a lifegain-themed Vampire.
- 1 Ajani, Strength of the Pride | wipe | Provides a sweeping reset that fits the deck's theme.
- 1 Fumigate | wipe | Resets a crowded board while supporting the plan.
- 1 Kaya's Wrath | wipe | Provides a reliable full-board reset.
- 1 Aerith Gainsborough | synergy | Supports the deck's lifegain-focused game plan.
- 1 Ajani's Pridemate | synergy | Rewards repeated life gain with a growing creature.
- 1 Angel of Vitality | synergy | Strengthens the deck's lifegain-focused creature plan.
- 1 Angelic Accord | synergy | Turns the central theme into a lasting board presence.
- 1 Blood Artist | synergy | Adds a complementary life-total pressure angle.
- 1 Bloodthirsty Aerialist | synergy | Rewards recurring life gain with a growing evasive threat.
- 1 Cleric Class | synergy | Provides an enchantment-based lifegain payoff.
- 1 Cleric of Life's Bond | synergy | Fits the Cleric and lifegain portions of the deck.
- 1 Heliod, Sun-Crowned | synergy | Provides a powerful permanent-based lifegain payoff.
- 1 Indulging Patrician | synergy | Converts the lifegain theme into pressure on opponents.
- 1 Resplendent Angel | synergy | Rewards sustained life gain with board development.
- 1 Sanguine Bond | synergy | Turns life gain into a direct pressure engine.
- 1 Vito, Thorn of the Dusk Rose | synergy | Adds another lifegain-to-pressure payoff.
- 1 Voice of the Blessed | synergy | Grows into a major threat through repeated life gain.
- 1 Archangel of Thune | threat | Turns lifegain into a powerful team-wide threat.
- 1 Astarion, the Decadent | threat | Provides a high-impact black-white finisher.
- 1 Attended Healer | threat | Builds a board while supporting the lifegain plan.
- 1 Celestine, the Living Saint | threat | Provides a substantial lifegain-themed threat.
- 1 Cliffhaven Vampire | threat | Pressures opponents through the deck's central theme.
- 1 Defiant Bloodlord | threat | Creates a large life-total pressure payoff.
- 1 Divinity of Pride | threat | Serves as a sizeable lifegain-themed attacker.
- 1 Elenda, Saint of Dusk | threat | Adds a resilient Vampire threat to the deck.
- 1 Exalted Sunborn | threat | Provides an Angel threat that supports the theme.
- 1 Nykthos Paragon | threat | Converts lifegain into a powerful creature-board payoff.
- 1 Rhox Faithmender | threat | Amplifies the lifegain plan on a durable body.
- 1 Valkyrie Harbinger | threat | Provides a top-end Angel threat for sustained games.

</details>

### 2. aristocrats Commander, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 178 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $0.00 to buy, $145.33 the whole deck.

**Summary:** This mana-heavy aristocrats shell develops resources early, keeps cards flowing, and uses its synergy creatures alongside removal, protection, and board clears to hold a favorable table position. It aims to win by building sustained creature pressure after opposing boards have been checked. It gives up some dedicated payoff and threat depth for extra acceleration, card flow, and broad answers from the library.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 2.70 over 63 nonland cards

<details><summary>The deck list</summary>

- 18 Plains | land | Provides reliable white mana in the basic mana base.
- 18 Swamp | land | Provides reliable black mana in the basic mana base.
- 1 Sol Ring | ramp | Supplies efficient early acceleration.
- 1 Arcane Signet | ramp | Provides dependable color fixing and acceleration.
- 1 Commander's Sphere | ramp | Adds flexible mana acceleration.
- 1 Fellwar Stone | ramp | Provides another inexpensive mana source.
- 1 Wayfarer's Bauble | ramp | Helps develop the mana base.
- 1 Thought Vessel | ramp | Contributes to the deck's mana development.
- 1 Relic of Legends | ramp | Adds a reusable source of acceleration.
- 1 Springleaf Drum | ramp | Offers low-cost early acceleration.
- 1 Astral Cornucopia | ramp | Provides scalable mana production.
- 1 Deadly Dispute | ramp | Supports the resource plan while accelerating.
- 1 Chromatic Lantern | ramp | Provides broad mana fixing and acceleration.
- 1 Sword of the Animist | ramp | Turns creature deployment into mana development.
- 1 Inherited Envelope | ramp | Adds another artifact-based mana source.
- 1 Lotho, Corrupt Shirriff | ramp | Contributes to the deck's acceleration package.
- 1 White Lotus Tile | ramp | Provides additional mana acceleration.
- 1 Ring of the Lucii | ramp | Adds a further source of acceleration.
- 1 Skullclamp | draw | Provides efficient card flow for the creature-focused plan.
- 1 Idol of Oblivion | draw | Adds a repeatable source of card advantage.
- 1 Night's Whisper | draw | Provides inexpensive card flow.
- 1 Painful Truths | draw | Refills the hand efficiently.
- 1 Nasty End | draw | Supplies card advantage for the creature plan.
- 1 Call of the Ring | draw | Provides continuing card flow.
- 1 Grave Venerations | draw | Supports sustained card advantage.
- 1 Mask of Memory | draw | Turns creature combat into card selection.
- 1 Tome of Legends | draw | Provides a durable card-advantage engine.
- 1 Wall of Omens | draw | Adds an early creature that replaces itself.
- 1 Lembas | draw | Provides compact card value.
- 1 Buster Sword | draw | Adds equipment-based card advantage.
- 1 Puresteel Paladin | draw | Supports the equipment suite while providing card flow.
- 1 Inspiring Overseer | draw | Adds a creature body and card advantage.
- 1 Massacre Girl, Known Killer | draw | Provides card advantage on a relevant creature body.
- 1 Bitter Triumph | removal | Provides flexible spot removal.
- 1 Claim the Precious | removal | Adds a direct answer to opposing permanents.
- 1 Crib Swap | removal | Provides creature removal at instant speed.
- 1 Generous Gift | removal | Answers a wide range of opposing permanents.
- 1 Get Lost | removal | Provides efficient broad removal.
- 1 Infernal Grasp | removal | Adds reliable creature removal.
- 1 Swords to Plowshares | removal | Provides an efficient answer to creatures.
- 1 Stroke of Midnight | removal | Adds flexible instant-speed removal.
- 1 Fatal Push | removal | Provides low-cost creature interaction.
- 1 Heartless Act | removal | Adds another efficient removal spell.
- 1 Dismember | removal | Provides a flexible answer to creatures.
- 1 Gollum the Abandoned | removal | Adds removal on a creature body.
- 1 Lightning Greaves | interaction | Protects key creatures while supporting the equipment package.
- 1 Swiftfoot Boots | interaction | Provides another protection option for important creatures.
- 1 Darksteel Plate | interaction | Helps preserve a key creature through opposing answers.
- 1 Champion's Helm | interaction | Protects an important legendary creature.
- 1 Clever Concealment | interaction | Protects the board from opposing interaction.
- 1 Reprieve | interaction | Provides a flexible tempo-focused answer.
- 1 Take Up the Shield | interaction | Offers creature protection during key exchanges.
- 1 Gift of Immortality | interaction | Supports resilience for a central creature.
- 1 Austere Command | wipe | Provides a flexible reset when the board gets out of hand.
- 1 Fumigate | wipe | Supplies a clean creature-board reset.
- 1 Vanquish the Horde | wipe | Adds an efficient mass answer to creatures.
- 1 Dusk // Dawn | wipe | Provides another creature-focused board reset.
- 1 Arcade Cabinet | synergy | Adds a permanent that supports the deck's synergy plan.
- 1 Gollum, Patient Plotter | synergy | Supports the aristocrats-focused creature plan.
- 1 Gríma Wormtongue | synergy | Adds a synergistic creature to the core plan.
- 1 Heirloom Auntie | synergy | Supports the deck's creature-based synergies.
- 1 Joo Dee, One of Many | synergy | Adds another body aligned with the synergy package.
- 1 Nimble Hobbit | synergy | Supports the deck's creature-focused game plan.
- 1 Phantom Train | synergy | Provides an artifact piece for the synergy package.
- 1 Vengeful Villagers | threat | Provides creature pressure as a closing threat.

</details>

### 3. artifacts Commander, bracket 4

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 299 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $1301.64 to buy, $1301.64 the whole deck.

**Summary:** This deck builds a dense blue artifact board, using artifact ramp and draw to keep resources flowing while synergy pieces make the deck’s many artifacts work together. It wins by turning that accumulated artifact presence into pressure through its artifact creature threats, backed by removal, interaction, and broad reset cards when the table gets ahead. It gives up broader multicolor options and leans heavily on maintaining an artifact-focused board state.

- [INFO] `curve_summary`: average mana value 4.05 over 63 nonland cards
- [INFO] `basics_added`: the list was 1 card short, so the builder added 1 basic land

<details><summary>The deck list</summary>

- 19 Island | land | Basic blue land base for the deck.
- 1 Academy Ruins | land | Blue-compatible land slot with an artifact-focused identity.
- 1 Archway of Innovation | land | Land slot for the artifact-focused mana base.
- 1 Blinkmoth Nexus | land | Land slot that remains within the artifact theme.
- 1 Buried Ruin | land | Land slot suited to an artifact-centered deck.
- 1 Conqueror's Galleon // Conqueror's Foothold | land | Land slot that also carries an artifact card type.
- 1 Fomori Vault | land | Land slot for the mana base.
- 1 Glimmervoid | land | Land slot for an artifact-heavy mana base.
- 1 Hall of Tagsin | land | Land slot for the mana base.
- 1 Inventors' Fair | land | Land slot aligned with the artifact plan.
- 1 Mishra's Factory | land | Artifact-themed land slot.
- 1 Mishra's Foundry | land | Artifact-themed land slot.
- 1 Mishra's Workshop | land | Artifact-themed land slot.
- 1 Otawara, Soaring City | land | Blue land slot for the mana base.
- 1 Power Depot | land | Artifact land slot for the mana base.
- 1 Spire of Industry | land | Land slot supporting an artifact-heavy build.
- 1 Urza's Saga | land | Artifact-themed land slot.
- 1 Urza's Workshop | land | Artifact-themed land slot.
- 1 Mox Opal | ramp | Compact artifact ramp for the early game.
- 1 Metalworker | ramp | Artifact creature dedicated to ramping the deck.
- 1 Krark-Clan Ironworks | ramp | Artifact ramp piece for explosive artifact turns.
- 1 Moonsnare Prototype | ramp | Low-cost artifact ramp piece.
- 1 Chief Engineer | ramp | Artifact-focused creature ramp.
- 1 Grand Architect | ramp | Creature ramp that fits the blue artifact shell.
- 1 Inspiring Statuary | ramp | Artifact ramp support for nonartifact spells.
- 1 The Mightstone and Weakstone | ramp | Legendary artifact ramp piece.
- 1 Tezzeret the Seeker | ramp | Planeswalker ramp that fits the artifact strategy.
- 1 Karn, Legacy Reforged | ramp | Artifact creature ramp for bigger turns.
- 1 Thoughtcast | draw | Efficient artifact-themed card draw.
- 1 Thought Monitor | draw | Artifact creature that fills a draw slot.
- 1 Sai, Master Thopterist | draw | Artifact-focused legendary creature for card draw.
- 1 Thopter Spy Network | draw | Artifact-focused enchantment providing card draw.
- 1 Vedalken Archmage | draw | Artifact-focused creature for card draw.
- 1 Reverse Engineer | draw | Artifact-themed sorcery draw spell.
- 1 Thirst for Knowledge | draw | Instant-speed card draw for the blue shell.
- 1 Tezzeret, Artifice Master | draw | Artifact-themed planeswalker draw source.
- 1 Forensic Gadgeteer | draw | Artifact-focused creature in a draw slot.
- 1 Riddlesmith | draw | Artifact-focused creature for card draw.
- 1 Metallic Rebuke | interaction | Artifact-friendly instant interaction.
- 1 Stoic Rebuttal | interaction | Instant interaction for protecting the game plan.
- 1 Disruption Protocol | interaction | Artifact-themed instant interaction.
- 1 Assert Authority | interaction | High-impact instant interaction.
- 1 Welding Jar | interaction | Artifact interaction piece.
- 1 Jin-Gitaxias, Progress Tyrant | interaction | High-impact blue legendary interaction threat.
- 1 Aether Spellbomb | removal | Low-cost artifact removal that fits the deck's core theme.
- 1 Resculpt | removal | Flexible blue removal spell.
- 1 Ravenform | removal | Blue sorcery removal for problematic permanents.
- 1 Spine of Ish Sah | removal | Artifact removal option for the larger mana turns.
- 1 Portal to Phyrexia | removal | High-impact artifact removal piece.
- 1 Lux Cannon | removal | Artifact-based removal for the long game.
- 1 Arcum Dagsson | removal | Legendary artificer providing a removal slot.
- 1 Skysovereign, Consul Flagship | removal | Artifact Vehicle removal option.
- 1 Nevinyrral's Disk | wipe | Artifact board wipe for reset turns.
- 1 Oblivion Stone | wipe | Artifact board wipe for difficult boards.
- 1 Perilous Vault | wipe | Artifact board wipe for broad cleanup.
- 1 Emry, Lurker of the Loch | synergy | Artifact-focused legendary creature synergy.
- 1 Etherium Sculptor | synergy | Artifact creature that strengthens the deck's central theme.
- 1 Foundry Inspector | synergy | Artifact creature synergy for an artifact-dense list.
- 1 Mystic Forge | synergy | Artifact synergy centerpiece.
- 1 Clock of Omens | synergy | Artifact synergy engine.
- 1 Unwinding Clock | synergy | Artifact synergy support for sustained turns.
- 1 Voltaic Key | synergy | Low-cost artifact synergy piece.
- 1 Manifold Key | synergy | Low-cost artifact synergy piece.
- 1 Shimmer Myr | synergy | Artifact creature synergy for the main plan.
- 1 Scrap Trawler | synergy | Artifact creature synergy for resource-heavy games.
- 1 Transmute Artifact | synergy | Artifact-focused sorcery synergy.
- 1 Whir of Invention | synergy | Artifact-focused instant synergy.
- 1 Reshape | synergy | Artifact-focused sorcery synergy.
- 1 Mirran Spy | synergy | Creature synergy in the artifact shell.
- 1 Arcbound Crusher | threat | Artifact creature threat for pressuring opponents.
- 1 Kappa Cannoneer | threat | Artifact creature threat for closing games.
- 1 Cyberdrive Awakener | threat | Artifact creature threat in the deck's primary theme.
- 1 Metalwork Colossus | threat | Large artifact creature threat.
- 1 Kuldotha Forgemaster | threat | Artifact creature threat for high-power turns.
- 1 Master Transmuter | threat | Artifact creature threat that supports the deck's identity.
- 1 Karn, Scion of Urza | threat | Artifact-themed planeswalker threat.
- 1 Traxos, Scourge of Kroog | threat | Legendary artifact creature threat.
- 1 Mycosynth Golem | threat | Artifact creature threat at the top of the curve.
- 1 Broodstar | threat | Blue creature threat for an artifact-heavy deck.
- 1 Threefold Thunderhulk | threat | Artifact creature threat for ending games.
- 1 Darksteel Juggernaut | threat | Artifact creature threat that reinforces the central theme.

</details>

### 4. dinosaur tribal, bracket 2

Format: Commander. Theme: dinosaurs. Pool: any_card. Shortlist: 267 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $268.19 to buy, $268.19 the whole deck.

**Summary:** This deck builds around Gishath, Sun's Avatar by developing its mana, filling the board with Dinosaurs, and keeping cards flowing as the game progresses. It wins by applying steady pressure with large Dinosaur threats, while its removal, protection, and board-clearing cards help it push through difficult tables. It gives up some speed and precision for a straightforward, creature-focused plan that is easy to pilot.

- [INFO] `bracket_prose_rules`: bracket 2 (Core, 8+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 4.11 over 63 nonland cards

<details><summary>The deck list</summary>

- 12 Forest | land | Basic land slot for the mana base.
- 7 Plains | land | Basic land slot for the mana base.
- 7 Mountain | land | Basic land slot for the mana base.
- 1 Command Tower | land | Land slot for the three-color mana base.
- 1 Path of Ancestry | land | Land slot for the Dinosaur-focused mana base.
- 1 Canopy Vista | land | Land slot for the mana base.
- 1 Cinder Glade | land | Land slot for the mana base.
- 1 Rootbound Crag | land | Land slot for the mana base.
- 1 Clifftop Retreat | land | Land slot for the mana base.
- 1 Sunpetal Grove | land | Land slot for the mana base.
- 1 Temple Garden | land | Land slot for the mana base.
- 1 Sacred Foundry | land | Land slot for the mana base.
- 1 Stomping Ground | land | Land slot for the mana base.
- 1 Sol Ring | ramp | Early ramp for a creature-heavy plan.
- 1 Arcane Signet | ramp | Reliable ramp for the three-color deck.
- 1 Cultivate | ramp | Ramp that supports developing the mana base.
- 1 Farseek | ramp | Ramp to help cast the deck's larger cards.
- 1 Nature's Lore | ramp | Ramp for steady early development.
- 1 Kodama's Reach | ramp | Ramp that supports consistent land development.
- 1 Thunderherd Migration | ramp | Dinosaur-themed ramp for the deck's larger plays.
- 1 Drover of the Mighty | ramp | Creature-based ramp that fits the deck's plan.
- 1 Atzocan Seer | ramp | Creature-based ramp for the Dinosaur deck.
- 1 Hulking Raptor | ramp | Dinosaur ramp that contributes to the board.
- 1 Beast Whisperer | draw | Draw support alongside the deck's many creatures.
- 1 Garruk's Uprising | draw | Draw support for the large-creature plan.
- 1 Guardian Project | draw | Ongoing draw support for a creature-focused deck.
- 1 Harmonize | draw | Straightforward draw to refill the hand.
- 1 Return of the Wildspeaker | draw | Flexible draw support for a deck with large threats.
- 1 Ripjaw Raptor | draw | Dinosaur draw support that adds to the board.
- 1 Rishkar's Expertise | draw | Draw support for the deck's high-power creatures.
- 1 Shamanic Revelation | draw | Draw support for a developed creature board.
- 1 Vanquisher's Banner | draw | Dinosaur-focused draw support.
- 1 Vaultborn Tyrant | draw | A Dinosaur draw option that also advances the board.
- 1 Heroic Intervention | interaction | Interaction that helps preserve the creature plan.
- 1 Swiftfoot Boots | interaction | Interaction that supports a key creature.
- 1 Lightning Greaves | interaction | Interaction that supports a key creature.
- 1 Boros Charm | interaction | Flexible interaction for protecting the plan.
- 1 Flawless Maneuver | interaction | Interaction for protecting the creature board.
- 1 Inspiring Call | interaction | Interaction that supports the creature-focused strategy.
- 1 Apex Altisaur | removal | Dinosaur removal that remains a meaningful body.
- 1 Savage Stomp | removal | Dinosaur-themed removal for opposing problems.
- 1 Thrashing Brontodon | removal | Dinosaur removal that fits the main plan.
- 1 Ravenous Sailback | removal | Dinosaur removal that contributes to the board.
- 1 Territorial Allosaurus | removal | Dinosaur removal for creature-based opposition.
- 1 Itzquinth, Firstborn of Gishath | removal | Dinosaur removal that supports the tribe.
- 1 Sun-Crowned Hunters | removal | Dinosaur removal for advancing the board plan.
- 1 Burning Sun's Avatar | removal | Dinosaur removal that also serves the main theme.
- 1 Blasphemous Act | wipe | Board-clearing option when the table gets ahead.
- 1 Austere Command | wipe | Flexible board-clearing option for difficult boards.
- 1 Wakening Sun's Avatar | wipe | Dinosaur-themed board-clearing option.
- 1 Kinjalli's Caller | synergy | Dinosaur synergy that supports the deck's creature plan.
- 1 Otepec Huntmaster | synergy | Dinosaur synergy for a more focused tribal board.
- 1 Marauding Raptor | synergy | Dinosaur synergy that advances the tribe's plan.
- 1 Commune with Dinosaurs | synergy | Dinosaur synergy that improves access to the theme.
- 1 Dinosaur Stampede | synergy | Dinosaur synergy for closing with the creature board.
- 1 Huatli's Raptor | synergy | Low-cost Dinosaur synergy for building the board.
- 1 Raptor Companion | synergy | Simple Dinosaur synergy for the tribal plan.
- 1 Hunting Velociraptor | synergy | Dinosaur synergy that supports the creature strategy.
- 1 Regal Imperiosaur | synergy | Dinosaur synergy for a committed tribal board.
- 1 Kinjalli's Sunwing | synergy | Dinosaur synergy that supports the deck's game plan.
- 1 Belligerent Yearling | synergy | Dinosaur synergy for an aggressive creature plan.
- 1 Nest Robber | synergy | Early Dinosaur synergy for the tribal board.
- 1 Orazca Frillback | synergy | Dinosaur synergy that adds flexibility to the tribe.
- 1 Pugnacious Hammerskull | synergy | Dinosaur synergy that contributes to board pressure.
- 1 Ghalta, Primal Hunger | threat | Large Dinosaur threat for ending games through pressure.
- 1 Regisaur Alpha | threat | Dinosaur threat that strengthens the board plan.
- 1 Carnage Tyrant | threat | Large Dinosaur threat for the top end.
- 1 Etali, Primal Storm | threat | High-impact Dinosaur threat for the deck's top end.
- 1 Ghalta, Stampede Tyrant | threat | Large Dinosaur threat for powerful late turns.
- 1 Zetalpa, Primal Dawn | threat | Large Dinosaur threat for closing games.
- 1 Quartzwood Crasher | threat | Dinosaur threat that supports sustained board pressure.
- 1 Palani's Hatcher | threat | Dinosaur threat that contributes to a wide board.
- 1 Pantlaza, Sun-Favored | threat | Dinosaur threat that reinforces the tribal plan.
- 1 Goring Ceratops | threat | Dinosaur threat for creature-based pressure.
- 1 Ancient Brontodon | threat | Large Dinosaur threat for the deck's top end.
- 1 Colossal Dreadmaw | threat | Straightforward large Dinosaur threat for a new-player plan.

</details>

### 5. blink Commander, owned first

Format: Commander. Theme: blink. Pool: owned_first. Shortlist: 168 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $0.00 to buy, $155.71 the whole deck.

**Summary:** This deck develops a creature board around Gilraen, then leans on blink-focused pieces to keep that board working through the game. Its creatures supply much of the deck’s momentum, while a broad set of answers and reset buttons keeps opposing boards from taking over. It wins by maintaining a steady board presence and turning that presence into pressure, giving up some specialization for a balanced mix of creatures, support pieces, and answers.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 2.87 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Abandoned Air Temple | land | Land slot in the mana base.
- 1 Adventurer's Inn | land | Land slot in the mana base.
- 1 Ash Barrens | land | Land slot in the mana base.
- 1 Bonders' Enclave | land | Land slot in the mana base.
- 1 Command Tower | land | Land slot in the mana base.
- 1 Eclipsed Realms | land | Land slot in the mana base.
- 1 Escape Tunnel | land | Land slot in the mana base.
- 1 Evolving Wilds | land | Land slot in the mana base.
- 1 Exotic Orchard | land | Land slot in the mana base.
- 1 Fabled Passage | land | Land slot in the mana base.
- 1 Field of Ruin | land | Land slot in the mana base.
- 1 Fountainport | land | Land slot in the mana base.
- 1 Ghost Quarter | land | Land slot in the mana base.
- 1 Great Hall of the Citadel | land | Land slot in the mana base.
- 1 Minas Tirith | land | Land slot in the mana base.
- 1 Opal Palace | land | Land slot in the mana base.
- 1 Path of Ancestry | land | Land slot in the mana base.
- 9 Plains | land | Basic land slots that keep the mana base consistent.
- 1 Plaza of Heroes | land | Land slot in the mana base.
- 1 Reliquary Tower | land | Land slot in the mana base.
- 1 Rogue's Passage | land | Land slot in the mana base.
- 1 Scavenger Grounds | land | Land slot in the mana base.
- 1 Secluded Courtyard | land | Land slot in the mana base.
- 1 Shire Terrace | land | Land slot in the mana base.
- 1 Study Hall | land | Land slot in the mana base.
- 1 Temple of the False God | land | Land slot in the mana base.
- 1 War Room | land | Land slot in the mana base.
- 1 Windbrisk Heights | land | Land slot in the mana base.
- 1 Adventurer's Airship | draw | Draw slot for maintaining access to cards.
- 1 Buster Sword | draw | Draw slot for maintaining access to cards.
- 1 Crown of Gondor | draw | Draw slot for maintaining access to cards.
- 1 Diary of Dreams | draw | Draw slot for maintaining access to cards.
- 1 Energybending | draw | Draw slot for maintaining access to cards.
- 1 Idol of Oblivion | draw | Draw slot for maintaining access to cards.
- 1 Instant Ramen | draw | Draw slot for maintaining access to cards.
- 1 Lembas | draw | Draw slot for maintaining access to cards.
- 1 Mask of Memory | draw | Draw slot for maintaining access to cards.
- 1 Mirror of Galadriel | draw | Draw slot for maintaining access to cards.
- 1 Arcane Signet | ramp | Ramp slot for accelerating the deck's development.
- 1 Bender's Waterskin | ramp | Ramp slot for accelerating the deck's development.
- 1 Fellwar Stone | ramp | Ramp slot for accelerating the deck's development.
- 1 Relic of Legends | ramp | Ramp slot for accelerating the deck's development.
- 1 Ring of the Lucii | ramp | Ramp slot for accelerating the deck's development.
- 1 Sol Ring | ramp | Ramp slot for accelerating the deck's development.
- 1 Sword of the Animist | ramp | Ramp slot for accelerating the deck's development.
- 1 Thought Vessel | ramp | Ramp slot for accelerating the deck's development.
- 1 Wayfarer's Bauble | ramp | Ramp slot for accelerating the deck's development.
- 1 White Lotus Tile | ramp | Ramp slot for accelerating the deck's development.
- 1 Banishing Light | removal | Removal slot for answering opposing cards.
- 1 Battle Menu | removal | Removal slot for answering opposing cards.
- 1 Crib Swap | removal | Removal slot for answering opposing cards.
- 1 Destroy Evil | removal | Removal slot for answering opposing cards.
- 1 Dispatch | removal | Removal slot for answering opposing cards.
- 1 Generous Gift | removal | Removal slot for answering opposing cards.
- 1 Get Lost | removal | Removal slot for answering opposing cards.
- 1 Swords to Plowshares | removal | Removal slot for answering opposing cards.
- 1 Bastion Protector | interaction | Interaction slot that supports the deck's board plan.
- 1 Clever Concealment | interaction | Interaction slot that supports the deck's board plan.
- 1 Lightning Greaves | interaction | Interaction slot that supports the deck's board plan.
- 1 Slip On the Ring | interaction | Interaction slot that supports the deck's board plan.
- 1 Swiftfoot Boots | interaction | Interaction slot that supports the deck's board plan.
- 1 Together Forever | interaction | Interaction slot that supports the deck's board plan.
- 1 Austere Command | wipe | Wipe slot for resetting difficult boards.
- 1 Fumigate | wipe | Wipe slot for resetting difficult boards.
- 1 Vanquish the Horde | wipe | Wipe slot for resetting difficult boards.
- 1 Angel of Condemnation | synergy | Creature-focused synergy slot for the blink plan.
- 1 Angel of Sanctions | synergy | Creature-focused synergy slot for the blink plan.
- 1 Angel of Serenity | synergy | Creature-focused synergy slot for the blink plan.
- 1 Champions of Minas Tirith | synergy | Creature-focused synergy slot for the blink plan.
- 1 Ennis, Debate Moderator | synergy | Dedicated synergy slot for the blink plan.
- 1 Faramir, Field Commander | synergy | Creature-focused synergy slot for the blink plan.
- 1 Fiend Hunter | synergy | Creature-focused synergy slot for the blink plan.
- 1 Flickerwisp | synergy | Dedicated synergy slot for the blink plan.
- 1 Inspiring Overseer | synergy | Creature-focused synergy slot for the blink plan.
- 1 Jocasta, Automaton Avenger | synergy | Dedicated synergy slot for the blink plan.
- 1 Joined Researchers // Secret Rendezvous | synergy | Creature-focused synergy slot for the blink plan.
- 1 Palace Jailer | synergy | Creature-focused synergy slot for the blink plan.
- 1 Personify | synergy | Dedicated synergy slot for the blink plan.
- 1 Wall of Omens | synergy | Creature-focused synergy slot for the blink plan.
- 1 Aang, the Last Airbender | threat | Threat slot that advances the deck's board plan.
- 1 Boromir, Warden of the Tower | threat | Threat slot that advances the deck's board plan.
- 1 Bronze Guardian | threat | Threat slot that advances the deck's board plan.
- 1 Exemplar of Light | threat | Threat slot that advances the deck's board plan.
- 1 Frontline Medic | threat | Threat slot that advances the deck's board plan.
- 1 Giada, Font of Hope | threat | Threat slot that advances the deck's board plan.
- 1 Kataki, War's Wage | threat | Threat slot that advances the deck's board plan.
- 1 Puresteel Paladin | threat | Threat slot that advances the deck's board plan.
- 1 South Pole Voyager | threat | Threat slot that advances the deck's board plan.
- 1 Stiltzkin, Moogle Merchant | threat | Threat slot that advances the deck's board plan.
- 1 The Vision | threat | Threat slot that advances the deck's board plan.
- 1 Zack Fair | threat | Threat slot that advances the deck's board plan.

</details>

### 6. Modern tempo, tournament

Format: Modern. Theme: tempo. Pool: any_card. Shortlist: 167 names.

Cards: 66 main, 15 sideboard. Repair turn: false. Block findings: 0.

Cost: $1331.27 to buy, $1331.27 the whole deck.

**Summary:** This blue-red tempo deck uses draw, removal, and interaction to keep opposing plays in check while seeking to convert its creature cards into a fast win. Temporal Mastery supplies the deck's listed synergy, while the sideboard shifts toward additional removal, interaction, wipes, and draw when those tools matter more. It gives up a dedicated threat package from the available pool, so its pressure is less redundant than a typical Delver tempo build.

- [INFO] `curve_summary`: average mana value 1.62 over 42 nonland cards

<details><summary>The deck list</summary>

- 8 Island | land | Included as a basic land for the land base.
- 4 Mountain | land | Included as a basic land for the land base.
- 4 Steam Vents | land | Included as a land for the blue-red land base.
- 4 Shivan Reef | land | Included as a land for the blue-red land base.
- 4 Stormcarved Coast | land | Included as a land for the blue-red land base.
- 4 Consider | draw | Included as a draw spell.
- 4 Opt | draw | Included as a draw spell.
- 4 Preordain | draw | Included as a draw spell.
- 2 Counterspell | interaction | Included as interaction.
- 4 Spell Pierce | interaction | Included as interaction.
- 4 Lightning Bolt | removal | Included as removal.
- 2 Into the Flood Maw | removal | Included as removal.
- 2 Pongify | removal | Included as removal.
- 4 Temporal Mastery | synergy | Included for synergy.
- 4 Ragavan, Nimble Pilferer | ramp | Included as a ramp card.
- 4 Mox Opal | ramp | Included as a ramp card.
- 4 Talisman of Creativity | ramp | Included as a ramp card.

</details>

### 7. Modern burn, casual

Format: Modern. Theme: burn. Pool: any_card. Shortlist: 290 names.

Cards: 60 main, 15 sideboard. Repair turn: false. Block findings: 0.

Cost: $75.82 to buy, $75.82 the whole deck.

**Summary:** This deck plays as a straightforward mono-red pressure deck: deploy threats, use direct removal to keep attacks moving, and lean on burn to finish the opponent once damage has added up. Draw spells help it continue applying pressure instead of running out of action. It gives up broad answers and defensive depth in exchange for a focused, aggressive damage plan.

- [INFO] `curve_summary`: average mana value 2.89 over 36 nonland cards

<details><summary>The deck list</summary>

- 24 Mountain | land | Provides the mono-red mana base.
- 4 Browbeat | draw | Supplies card draw to keep burn spells and threats coming.
- 2 Risk Factor | draw | Adds more card draw for the attrition plan.
- 2 Chandra, Dressed to Kill | ramp | Fills the ramp slot while supporting the red game plan.
- 4 Lightning Bolt | removal | Efficient removal that also supports the burn plan.
- 2 Lightning Strike | removal | Additional direct removal for clearing blockers or finishing pressure.
- 4 Eidolon of the Great Revel | synergy | A synergy piece for a low-cost red pressure strategy.
- 4 Thermo-Alchemist | synergy | Supports the deck's burn-focused synergy package.
- 4 Akki Lavarunner // Tok-Tok, Volcano Born | threat | A creature threat that helps establish board pressure.
- 4 Ashcloud Phoenix | threat | A resilient-looking threat slot for sustained pressure.
- 2 Hazoret the Fervent | threat | A powerful red threat for closing games.
- 4 Torbran, Thane of Red Fell | threat | A red threat that complements the damage-focused plan.

</details>

### 8. Modern lifegain, FNM

Format: Modern. Theme: lifegain. Pool: any_card. Shortlist: 301 names.

Cards: 60 main, 15 sideboard. Repair turn: false. Block findings: 0.

Cost: $292.92 to buy, $292.92 the whole deck.

**Summary:** This white-black deck builds around lifegain synergies, using them to establish a board before applying pressure with its creature threats while removal handles opposing threats. It wins by converting that lifegain-focused board into sustained attacks. The tradeoff is a creature-heavy plan that takes time to assemble and needs its mana base to support both colors.

- [INFO] `curve_summary`: average mana value 3.28 over 36 nonland cards

<details><summary>The deck list</summary>

- 8 Plains | land | White basic land for the mana base.
- 8 Swamp | land | Black basic land for the mana base.
- 4 Scoured Barrens | land | White-black land for the mana base.
- 2 Shambling Vent | land | White-black land for the mana base.
- 2 Restless Fortress | land | White-black land for the mana base.
- 4 Inspiring Overseer | draw | Draw card that also fits the deck's white creature base.
- 2 Dawn of Hope | draw | Draw card for the lifegain-focused plan.
- 2 Orzhov Keyrune | ramp | White-black ramp option.
- 2 Solitude | removal | Removal for opposing threats.
- 2 Murderous Rider // Swift End | removal | Removal that fits the black half of the deck.
- 2 The Wandering Emperor | removal | Removal option in white.
- 4 Soul Warden | synergy | Core lifegain synergy creature.
- 4 Ajani's Pridemate | synergy | Core lifegain synergy creature.
- 4 Attended Healer | threat | Creature threat for the lifegain plan.
- 4 Bloodbond Vampire | threat | Creature threat for the lifegain plan.
- 2 Archangel of Thune | threat | Top-end creature threat.
- 2 Cliffhaven Vampire | threat | Black creature threat for the lifegain plan.
- 2 Regal Bloodlord | threat | Additional creature threat in black.

</details>

### 9. Standard midrange, FNM

Format: Standard. Theme: midrange. Pool: any_card. Shortlist: 164 names.

Cards: 60 main, 15 sideboard. Repair turn: false. Block findings: 0.

Cost: $407.87 to buy, $407.87 the whole deck.

**Summary:** This black-green midrange deck builds around creature threats, permanent-based synergy pieces, draw sources, and flexible removal. It aims to take control of the table with answers, then apply pressure with increasingly substantial creatures to finish the game. It gives up the speed of a focused aggressive strategy in exchange for a broader main-deck plan and a sideboard that can shift toward more targeted answers.

- [INFO] `curve_summary`: average mana value 3.06 over 34 nonland cards

<details><summary>The deck list</summary>

- 6 Forest | land | Basic green land for the black-green mana base.
- 6 Swamp | land | Basic black land for the black-green mana base.
- 4 Overgrown Tomb | land | Black-green land for consistent access to both deck colors.
- 4 Blooming Marsh | land | Black-green land supporting the two-color mana base.
- 4 Deathcap Glade | land | Black-green land supporting the two-color mana base.
- 2 Llanowar Elves | ramp | Early ramp to support the deck's higher-cost creature threats.
- 4 Phyrexian Arena | draw | Primary repeatable draw selection for the midrange plan.
- 2 Darkstar Augur | draw | Creature-based draw that contributes to the board plan.
- 4 Bitter Triumph | removal | Efficient removal for opposing threats.
- 2 Assassin's Trophy | removal | Flexible removal for problematic opposing permanents.
- 4 Agatha's Soul Cauldron | synergy | A central permanent in the deck's synergy package.
- 4 Insidious Roots | synergy | A second core permanent for the creature-focused synergy package.
- 4 Goldvein Hydra | threat | Creature threat that helps establish a strong midrange board.
- 4 Vein Ripper | threat | Large creature threat for closing games.
- 4 Vaultborn Tyrant | threat | Top-end creature threat for the midrange plan.
- 2 Ojer Kaslem, Deepest Growth // Temple of Cultivation | threat | Legendary creature threat that also offers a land-facing side.

</details>

### 10. Standard aggro, tournament

Format: Standard. Theme: aggro. Pool: any_card. Shortlist: 294 names.

Cards: 60 main, 15 sideboard. Repair turn: false. Block findings: 0.

Cost: $68.62 to buy, $68.62 the whole deck.

**Summary:** This red-white aggro deck looks to build pressure with a dense creature threat base, reinforce that plan with red-white synergy cards, and clear resistance with dedicated removal and interaction. It wins by keeping threats on the table and pressing that advantage, while draw cards help it continue deploying resources. The tradeoff is a focused proactive plan with comparatively few broad reset effects in the main deck, so the sideboard supplies extra answers for matchups that demand them.

- [INFO] `curve_summary`: average mana value 3.06 over 36 nonland cards

<details><summary>The deck list</summary>

- 10 Mountain | land | A basic red land for the deck's red cards.
- 10 Plains | land | A basic white land for the deck's white cards.
- 4 Sacred Foundry | land | A red-white land that supports both main colors.
- 2 Fugitive Codebreaker | draw | A red threat that also fills the deck's draw slot.
- 2 Inspiring Overseer | draw | A white creature that supplies draw support.
- 2 Reckless Lackey | draw | A red creature included for draw support.
- 2 Boros Charm | interaction | A red-white instant included as interaction.
- 2 Aven Interrupter | interaction | A white creature that provides interaction.
- 2 Sheltered by Ghosts | interaction | A white enchantment used for interaction.
- 2 Harsh Annotation | removal | A red instant assigned to the removal package.
- 2 Emeritus of Truce // Swords to Plowshares | removal | A white split card used for removal.
- 2 Krenko's Buzzcrusher | removal | A red artifact creature included for removal.
- 2 Aurelia's Vindicator | removal | A white creature that strengthens the removal suite.
- 2 Shock Brigade | synergy | A red creature included for aggressive synergy.
- 2 Warleader's Call | synergy | A red-white enchantment that supports the deck's synergy plan.
- 2 Redcap Gutter-Dweller | threat | A red creature chosen as an aggressive threat.
- 2 Teapot Slinger | threat | A red creature that adds to the threat base.
- 2 Dragonback Lancer | threat | A red-white creature included as a threat.
- 2 Diversion Specialist | threat | A red-white creature that contributes to board pressure.
- 2 Frilled Sparkshooter | threat | A red creature included as a threat.
- 2 Delivery Moogle | threat | A white creature that rounds out the threat package.

</details>

### 11. Commander with a locked card

Format: Commander. Theme: sacrifice. Pool: any_card. Shortlist: 296 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $406.53 to buy, $406.53 the whole deck.

**Summary:** This deck is built to establish Karlov of the Ghost Council alongside a dense sacrifice package, using expendable resources and supporting permanents to keep the plan moving while building toward a threatening board. It wins through the pressure created by its sacrifice synergies, finishers, and removal-backed board control. The tradeoff is a narrow, committed game plan: the deck favors sacrifice cohesion over broad utility and relies on its key pieces working together.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 3.24 over 63 nonland cards

<details><summary>The deck list</summary>

- 18 Plains | land | Included as a basic land for the mana base.
- 18 Swamp | land | Included as a basic land for the mana base.
- 1 Corrupted Conviction | draw | Included as a draw card for the sacrifice-focused plan.
- 1 Disciple of Bolas | draw | Included as a draw card for the sacrifice-focused plan.
- 1 Relic Vial | draw | Included as a draw card for the sacrifice-focused plan.
- 1 Shadowheart, Dark Justiciar | draw | Included as a draw card for the sacrifice-focused plan.
- 1 Smothering Abomination | draw | Included as a draw card for the sacrifice-focused plan.
- 1 Tevesh Szat, Doom of Fools | draw | Included as a draw card for the sacrifice-focused plan.
- 1 Vampiric Rites | draw | Included as a draw card for the sacrifice-focused plan.
- 1 Village Rites | draw | Included as a draw card for the sacrifice-focused plan.
- 1 Baron Bertram Graywater | draw | Included as a draw card for the sacrifice-focused plan.
- 1 Bushmeat Poacher | draw | Included as a draw card for the sacrifice-focused plan.
- 1 Cartel Aristocrat | interaction | Included to provide interaction within the sacrifice shell.
- 1 Fanatical Devotion | interaction | Included to provide interaction within the sacrifice shell.
- 1 Gift of Doom | interaction | Included to provide interaction within the sacrifice shell.
- 1 Promise of Tomorrow | interaction | Included to provide interaction within the sacrifice shell.
- 1 Rescue from the Underworld | interaction | Included to provide interaction within the sacrifice shell.
- 1 Spirit Bonds | interaction | Included to provide interaction within the sacrifice shell.
- 1 Sol Ring | ramp | Kept in the deck as requested.
- 1 Ashnod's Altar | ramp | Included as ramp that fits the sacrifice-focused plan.
- 1 Crowded Crypt | ramp | Included as ramp that fits the sacrifice-focused plan.
- 1 Pawn of Ulamog | ramp | Included as ramp that fits the sacrifice-focused plan.
- 1 Phyrexian Altar | ramp | Included as ramp that fits the sacrifice-focused plan.
- 1 Pitiless Plunderer | ramp | Included as ramp that fits the sacrifice-focused plan.
- 1 Priest of Forgotten Gods | ramp | Included as ramp that fits the sacrifice-focused plan.
- 1 Sifter of Skulls | ramp | Included as ramp that fits the sacrifice-focused plan.
- 1 Skullport Merchant | ramp | Included as ramp that fits the sacrifice-focused plan.
- 1 Warren Soultrader | ramp | Included as ramp that fits the sacrifice-focused plan.
- 1 Attrition | removal | Included as removal for opposing permanents.
- 1 Ayli, Eternal Pilgrim | removal | Included as removal for opposing permanents.
- 1 Bone Shards | removal | Included as removal for opposing permanents.
- 1 Dictate of Erebos | removal | Included as removal for opposing permanents.
- 1 Eaten Alive | removal | Included as removal for opposing permanents.
- 1 Grave Pact | removal | Included as removal for opposing permanents.
- 1 Teysa, Orzhov Scion | removal | Included as removal for opposing permanents.
- 1 Yawgmoth, Thran Physician | removal | Included as removal for opposing permanents.
- 1 Altar of Dementia | synergy | Included for sacrifice-focused synergy.
- 1 Bartolomé del Presidio | synergy | Included for sacrifice-focused synergy.
- 1 Bastion of Remembrance | synergy | Included for sacrifice-focused synergy.
- 1 Bloodflow Connoisseur | synergy | Included for sacrifice-focused synergy.
- 1 Carrion Feeder | synergy | Included for sacrifice-focused synergy.
- 1 Fleshtaker | synergy | Included for sacrifice-focused synergy.
- 1 Hidden Stockpile | synergy | Included for sacrifice-focused synergy.
- 1 Open the Graves | synergy | Included for sacrifice-focused synergy.
- 1 Spawning Pit | synergy | Included for sacrifice-focused synergy.
- 1 Viscera Seer | synergy | Included for sacrifice-focused synergy.
- 1 Witch's Oven | synergy | Included for sacrifice-focused synergy.
- 1 Woe Strider | synergy | Included for sacrifice-focused synergy.
- 1 Yahenni, Undying Partisan | synergy | Included for sacrifice-focused synergy.
- 1 Zulaport Cutthroat | synergy | Included for sacrifice-focused synergy.
- 1 Abhorrent Overlord | threat | Included as a threat that advances the deck's board presence.
- 1 Felisa, Fang of Silverquill | threat | Included as a threat that advances the deck's board presence.
- 1 Ghoulcaller Gisa | threat | Included as a threat that advances the deck's board presence.
- 1 Liesa, Forgotten Archangel | threat | Included as a threat that advances the deck's board presence.
- 1 Liliana, Heretical Healer // Liliana, Defiant Necromancer | threat | Included as a threat that advances the deck's board presence.
- 1 Marrow-Gnawer | threat | Included as a threat that advances the deck's board presence.
- 1 Mondrak, Glory Dominus | threat | Included as a threat that advances the deck's board presence.
- 1 Ratadrabik of Urborg | threat | Included as a threat that advances the deck's board presence.
- 1 Razaketh, the Foulblooded | threat | Included as a threat that advances the deck's board presence.
- 1 Requiem Angel | threat | Included as a threat that advances the deck's board presence.
- 1 Sidisi, Undead Vizier | threat | Included as a threat that advances the deck's board presence.
- 1 Vindictive Vampire | threat | Included as a threat that advances the deck's board presence.
- 1 Elesh Norn // The Argent Etchings | wipe | Included as a board wipe for reset potential.
- 1 The Meathook Massacre | wipe | Included as a board wipe for reset potential.
- 1 Toxic Deluge | wipe | Included as a board wipe for reset potential.

</details>

### 12. Commander on a budget

Format: Commander. Theme: tokens. Pool: any_card. Shortlist: 296 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $47.65 to buy, $47.65 the whole deck.

**Summary:** This deck builds a token-focused board around Adeline, Resplendent Cathar and applies pressure through combat, using its threats to turn a wide board into a win. It gives up premium individual-card power for a low-cost, straightforward plan that develops steadily while retaining draw, ramp, answers, and board-reset options.

- [INFO] `bracket_prose_rules`: bracket 2 (Core, 8+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 3.25 over 63 nonland cards

<details><summary>The deck list</summary>

- 36 Plains | land | Basic land choice for the deck’s land base.
- 1 Bucknard's Everfull Purse | ramp | A ramp piece for developing the token plan.
- 1 Charisma Bobblehead | ramp | A ramp piece for developing the token plan.
- 1 Coin of Mastery | ramp | A ramp piece for developing the token plan.
- 1 Collector's Vault | ramp | A ramp piece for developing the token plan.
- 1 Currency Converter | ramp | A ramp piece for developing the token plan.
- 1 Druidic Satchel | ramp | A ramp piece for developing the token plan.
- 1 Goldvein Pick | ramp | A ramp piece for developing the token plan.
- 1 Idol of False Gods | ramp | A ramp piece for developing the token plan.
- 1 Karn, Living Legacy | ramp | A ramp piece for developing the token plan.
- 1 Keeper of the Accord | ramp | A ramp piece for developing the token plan.
- 1 Angelic Sell-Sword | draw | A draw piece that helps keep the deck moving.
- 1 Bygone Bishop | draw | A draw piece that helps keep the deck moving.
- 1 Dawn of Hope | draw | A draw piece that helps keep the deck moving.
- 1 Faramir, Field Commander | draw | A draw piece that helps keep the deck moving.
- 1 Glimmer Seeker | draw | A draw piece that helps keep the deck moving.
- 1 Hoofprints of the Stag | draw | A draw piece that helps keep the deck moving.
- 1 Idol of Oblivion | draw | A draw piece that helps keep the deck moving.
- 1 Platoon Dispenser | draw | A draw piece that helps keep the deck moving.
- 1 Sarah Jane Smith | draw | A draw piece that helps keep the deck moving.
- 1 Search the Premises | draw | A draw piece that helps keep the deck moving.
- 1 Ainok Strike Leader | interaction | An interaction piece for protecting the deck’s position.
- 1 Basri Ket | interaction | An interaction piece for protecting the deck’s position.
- 1 Basri, Tomorrow's Champion | interaction | An interaction piece for protecting the deck’s position.
- 1 Darksteel Splicer | interaction | An interaction piece for protecting the deck’s position.
- 1 Lena, Selfless Champion | interaction | An interaction piece for protecting the deck’s position.
- 1 Rootborn Defenses | interaction | An interaction piece for protecting the deck’s position.
- 1 Aerial Assault | removal | A removal option for answering opposing pieces.
- 1 Ajani, Outland Chaperone | removal | A removal option for answering opposing pieces.
- 1 Banishing Slash | removal | A removal option for answering opposing pieces.
- 1 Battle Menu | removal | A removal option for answering opposing pieces.
- 1 Citizen's Crowbar | removal | A removal option for answering opposing pieces.
- 1 Conversion Chamber | removal | A removal option for answering opposing pieces.
- 1 Kellan's Lightblades | removal | A removal option for answering opposing pieces.
- 1 Path to Redemption | removal | A removal option for answering opposing pieces.
- 1 Aligned Heart | synergy | A synergy piece for the deck’s token-focused plan.
- 1 Animation Module | synergy | A synergy piece for the deck’s token-focused plan.
- 1 Anointer Priest | synergy | A synergy piece for the deck’s token-focused plan.
- 1 Automated Assembly Line | synergy | A synergy piece for the deck’s token-focused plan.
- 1 Cat Collector | synergy | A synergy piece for the deck’s token-focused plan.
- 1 Clarion Spirit | synergy | A synergy piece for the deck’s token-focused plan.
- 1 Divine Visitation | synergy | A synergy piece for the deck’s token-focused plan.
- 1 Felidar Retreat | synergy | A synergy piece for the deck’s token-focused plan.
- 1 Hero of Precinct One | synergy | A synergy piece for the deck’s token-focused plan.
- 1 Intangible Virtue | synergy | A synergy piece for the deck’s token-focused plan.
- 1 Oketra's Monument | synergy | A synergy piece for the deck’s token-focused plan.
- 1 Retrofitter Foundry | synergy | A synergy piece for the deck’s token-focused plan.
- 1 Rosie Cotton of South Lane | synergy | A synergy piece for the deck’s token-focused plan.
- 1 Siege Veteran | synergy | A synergy piece for the deck’s token-focused plan.
- 1 Ajani's Chosen | threat | A low-cost threat that supports the deck’s pressure plan.
- 1 Archon of Sun's Grace | threat | A low-cost threat that supports the deck’s pressure plan.
- 1 Attended Healer | threat | A low-cost threat that supports the deck’s pressure plan.
- 1 Basri's Lieutenant | threat | A low-cost threat that supports the deck’s pressure plan.
- 1 Cemetery Protector | threat | A low-cost threat that supports the deck’s pressure plan.
- 1 Defiler of Faith | threat | A low-cost threat that supports the deck’s pressure plan.
- 1 Dragonback Lancer | threat | A low-cost threat that supports the deck’s pressure plan.
- 1 Drogskol Cavalry | threat | A low-cost threat that supports the deck’s pressure plan.
- 1 Emeria Angel | threat | A low-cost threat that supports the deck’s pressure plan.
- 1 Gideon, Ally of Zendikar | threat | A low-cost threat that supports the deck’s pressure plan.
- 1 Hero of Bladehold | threat | A low-cost threat that supports the deck’s pressure plan.
- 1 Mite Overseer | threat | A low-cost threat that supports the deck’s pressure plan.
- 1 Ceaseless Conflict | wipe | A wipe option for resetting a crowded board.
- 1 Crisis of Conscience | wipe | A wipe option for resetting a crowded board.
- 1 Descend upon the Sinful | wipe | A wipe option for resetting a crowded board.

</details>

### 13. owned first, and the commander is not owned

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 225 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $0.00 to buy, $181.41 the whole deck.

**Summary:** This deck centers on Karlov and a lifegain-focused creature board, using support pieces and equipment to keep its main plan moving while drawing into more action. It wins by building sustained pressure with its threats and clearing away opposing obstacles with efficient removal or sweeping resets. In exchange, the list leans heavily on maintaining creatures and key support permanents, so it is less focused on fast standalone finishes.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [WARN] `not_owned`: Karlov of the Ghost Council: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 2.89 over 63 nonland cards

<details><summary>The deck list</summary>

- 14 Plains | land | Provides the deck’s white mana base.
- 14 Swamp | land | Provides the deck’s black mana base.
- 1 Command Tower | land | A reliable multicolor land for the deck’s core colors.
- 1 Ash Barrens | land | Adds flexible land access.
- 1 Evolving Wilds | land | Helps stabilize the mana base.
- 1 Fabled Passage | land | Supports consistent access to the deck’s lands.
- 1 Marsh Flats | land | Finds a needed basic land type.
- 1 Path of Ancestry | land | A fixing land that supports the commander deck.
- 1 Scavenger Grounds | land | Adds utility without taking a nonland slot.
- 1 Takenuma, Abandoned Mire | land | A black source with useful utility.
- 1 Arcane Signet | ramp | Efficient mana development for the early turns.
- 1 Sol Ring | ramp | Accelerates the deck’s mana development.
- 1 Fellwar Stone | ramp | A compact mana accelerator.
- 1 Commander's Sphere | ramp | Provides mana support in a flexible slot.
- 1 Wayfarer's Bauble | ramp | Helps the deck develop its mana base.
- 1 Thought Vessel | ramp | A mana rock that supports longer games.
- 1 Sword of the Animist | ramp | Combines mana development with the deck’s creature plan.
- 1 Relic of Legends | ramp | Adds another efficient source of mana support.
- 1 Springleaf Drum | ramp | Helps turn the board presence into mana development.
- 1 Astral Cornucopia | ramp | A flexible artifact ramp piece.
- 1 Buster Sword | draw | Provides card flow while fitting the equipment package.
- 1 Call of the Ring | draw | A dedicated source of ongoing card flow.
- 1 Exemplar of Light | draw | Adds card flow on a creature that suits the deck’s theme.
- 1 Gollum, Riddle Master | draw | Provides card flow in a creature slot.
- 1 Idol of Oblivion | draw | A compact artifact source of card flow.
- 1 Lembas | draw | Supports the deck with a low-cost draw piece.
- 1 Mask of Memory | draw | Pairs card flow with the creature-focused plan.
- 1 Night's Whisper | draw | A direct source of card flow.
- 1 Skullclamp | draw | An efficient equipment-based draw option.
- 1 Tome of Legends | draw | Supports steady card flow around the commander.
- 1 Bastion Protector | interaction | Helps preserve the commander and key board pieces.
- 1 Clever Concealment | interaction | Protects the board from opposing disruption.
- 1 Darksteel Plate | interaction | A durable protection piece for a key creature.
- 1 Lightning Greaves | interaction | Protects an important creature efficiently.
- 1 Swiftfoot Boots | interaction | Provides another protection tool for key creatures.
- 1 Take Up the Shield | interaction | A flexible answer to targeted pressure.
- 1 Bitter Triumph | removal | A versatile answer to a problematic permanent.
- 1 Dispatch | removal | A low-cost removal option.
- 1 Generous Gift | removal | Answers a broad range of opposing problems.
- 1 Get Lost | removal | An efficient answer for troublesome permanents.
- 1 Infernal Grasp | removal | Straightforward creature removal.
- 1 Path to Exile | removal | A premier answer to opposing creatures.
- 1 Stroke of Midnight | removal | Flexible permanent removal.
- 1 Swords to Plowshares | removal | An efficient answer to opposing creatures.
- 1 Austere Command | wipe | A flexible reset when the board gets out of hand.
- 1 Dusk // Dawn | wipe | A board reset that fits the creature-focused build.
- 1 Fumigate | wipe | A sweeping answer that supports the lifegain plan.
- 1 Aerith Gainsborough | synergy | A central themed piece for the lifegain strategy.
- 1 Angel of Vitality | synergy | Supports the deck’s lifegain-focused plan.
- 1 Aettir and Priwen | synergy | An equipment piece that supports the deck’s main plan.
- 1 Compassionate Healer | synergy | A dedicated lifegain-theme creature.
- 1 Elixir | synergy | A compact support piece for the lifegain plan.
- 1 Kor Firewalker | synergy | Supports the deck’s life-focused strategy.
- 1 Light of Promise | synergy | Rewards the deck for staying focused on lifegain.
- 1 Northern Air Temple | synergy | A persistent support piece for the deck’s theme.
- 1 Prideful Feastling | synergy | Adds another creature devoted to the lifegain plan.
- 1 Pull from the Grave | synergy | Supports the deck’s creature-based strategy.
- 1 Rosie Cotton of South Lane | synergy | A strong thematic creature for the life-focused board.
- 1 Second Breakfast | synergy | A focused support spell for the deck’s central plan.
- 1 Well-Worn Spatula | synergy | An artifact support piece for the creature plan.
- 1 White Mage's Staff | synergy | An equipment support card for the lifegain strategy.
- 1 Angel of Invention | threat | A proactive creature that adds meaningful board pressure.
- 1 Bill the Pony | threat | A creature threat that advances the board plan.
- 1 Dawnhand Eulogist | threat | A threat that helps convert board development into pressure.
- 1 Foggy Swamp Hunters | threat | Adds a substantial creature presence.
- 1 Gollum, Silent Slinker // Meager Meal | threat | A proactive threat that fits the creature plan.
- 1 Lo and Li, Twin Tutors | threat | Adds another impactful threat to the board.
- 1 Lyra Dawnbringer | threat | A high-impact creature threat for closing games.
- 1 Minwu, White Mage | threat | A themed threat that supports the deck’s creature plan.
- 1 Reaping Willow | threat | A resilient board-focused threat.
- 1 Shattered Angel | threat | A creature threat suited to the deck’s life-focused direction.
- 1 Sneering Shadewriter | threat | Adds pressure in a creature slot.
- 1 Victory's Herald | threat | A powerful top-end threat for ending games.

</details>

### 14. the user delegates the commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 244 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $1042.58 to buy, $1042.58 the whole deck.

**Summary:** Led by Atarka, World Render, this deck builds mana, develops a Dragon board, and turns a stream of large flying threats into forceful combat finishes. Its draw package helps it continue deploying Dragons, while removal and sweepers clear away the blockers and opposing threats that slow the attack. The deck gives up a lower curve and broad reactive coverage in exchange for a creature-heavy plan that is strongest once its mana engine and Dragon synergies are established.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 4.46 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Ancient Tomb | land | A land slot that helps support the deck's expensive Dragon plan.
- 1 Arid Mesa | land | A land slot for the red-green mana base.
- 1 Bloodstained Mire | land | A land slot for the red-green mana base.
- 1 Boseiju, Who Endures | land | A land slot that supports the deck's mana base.
- 1 Cavern of Souls | land | A land slot tailored to the Dragon-focused creature base.
- 1 Cinder Glade | land | A red-green land slot.
- 1 City of Brass | land | A flexible land slot for the mana base.
- 1 Command Tower | land | A central land slot for the commander's colors.
- 1 Crucible of the Spirit Dragon | land | A land slot aligned with the Dragon plan.
- 1 Evolving Wilds | land | A land slot for stabilizing the mana base.
- 1 Exotic Orchard | land | A flexible land slot for the mana base.
- 1 Fabled Passage | land | A land slot for stabilizing the mana base.
- 1 Haven of the Spirit Dragon | land | A land slot aligned with the Dragon plan.
- 1 Mana Confluence | land | A flexible land slot for the mana base.
- 1 Misty Rainforest | land | A land slot for the red-green mana base.
- 1 Myriad Landscape | land | A land slot that supports the deck's mana development.
- 1 Nykthos, Shrine to Nyx | land | A land slot that supports the creature-heavy plan.
- 1 Path of Ancestry | land | A land slot tailored to the Dragon creature base.
- 1 Reflecting Pool | land | A flexible land slot for the mana base.
- 1 Rogue's Passage | land | A land slot that supports the attacking plan.
- 1 Rootbound Crag | land | A red-green land slot.
- 1 Scalding Tarn | land | A land slot for the red-green mana base.
- 1 Stomping Ground | land | A red-green land slot.
- 1 Temple of the Dragon Queen | land | A land slot aligned with the Dragon plan.
- 1 Temple of the False God | land | A land slot for the deck's large threats.
- 1 Terramorphic Expanse | land | A land slot for stabilizing the mana base.
- 1 Three Tree City | land | A land slot tailored to the Dragon creature base.
- 1 War Room | land | A land slot that supports the deck's resources.
- 1 Windswept Heath | land | A land slot for the red-green mana base.
- 1 Wooded Foothills | land | A land slot for the red-green mana base.
- 1 Yavimaya, Cradle of Growth | land | A land slot for the green side of the mana base.
- 3 Forest | land | Basic green land slots for the mana base.
- 2 Mountain | land | Basic red land slots for the mana base.
- 1 Ancient Copper Dragon | ramp | A Dragon ramp piece that advances the large-creature plan.
- 1 Atsushi, the Blazing Sky | ramp | A Dragon ramp piece for building toward larger turns.
- 1 Carnelian Orb of Dragonkind | ramp | A ramp piece aligned with the Dragon plan.
- 1 Dragon's Hoard | ramp | A ramp piece aligned with the Dragon plan.
- 1 Ganax, Astral Hunter | ramp | A Dragon ramp piece that supports deploying more threats.
- 1 Goldspan Dragon | ramp | A Dragon ramp piece for explosive development.
- 1 Jade Orb of Dragonkind | ramp | A ramp piece aligned with the Dragon plan.
- 1 Klauth, Unrivaled Ancient | ramp | A Dragon ramp piece that supports powerful combat turns.
- 1 Old Gnawbone | ramp | A Dragon ramp piece for the creature-heavy plan.
- 1 Savage Ventmaw | ramp | A Dragon ramp piece that supports attacking and deploying threats.
- 1 Avaricious Dragon | draw | A Dragon draw piece that keeps the deck supplied.
- 1 Beast Whisperer | draw | A draw piece for the creature-heavy plan.
- 1 Elemental Bond | draw | A draw piece that supports large creatures.
- 1 Garruk's Uprising | draw | A draw piece for the large-Dragon strategy.
- 1 Guardian Project | draw | A draw piece for the creature-heavy plan.
- 1 Knollspine Dragon | draw | A Dragon draw piece that refuels the deck.
- 1 Return of the Wildspeaker | draw | A draw piece that fits the large-creature plan.
- 1 Rishkar's Expertise | draw | A draw piece that rewards large threats.
- 1 Sylvan Library | draw | A draw piece for maintaining resources.
- 1 Vanquisher's Banner | draw | A draw piece aligned with the Dragon creature base.
- 1 Heroic Intervention | interaction | An interaction piece for protecting the developed board.
- 1 Lightning Greaves | interaction | An interaction piece for safeguarding a key creature.
- 1 Mithril Coat | interaction | An interaction piece for protecting a major threat.
- 1 Swiftfoot Boots | interaction | An interaction piece for safeguarding a key creature.
- 1 Tamiyo's Safekeeping | interaction | An interaction piece for protecting an important permanent.
- 1 Veil of Summer | interaction | An interaction piece for defending the deck's plan.
- 1 Dragon Tempest | removal | A Dragon-focused removal piece that supports pressure.
- 1 Dragonlord Atarka | removal | A Dragon removal piece that also adds a major body.
- 1 Drakuseth, Maw of Flames | removal | A Dragon removal piece that supports the attacking plan.
- 1 Foe-Razer Regent | removal | A Dragon removal piece for handling opposing threats.
- 1 Glorybringer | removal | A Dragon removal piece that maintains pressure.
- 1 Scourge of Valkas | removal | A Dragon removal piece aligned with the tribal plan.
- 1 Terror of the Peaks | removal | A Dragon removal piece that supports successive threats.
- 1 Wrathful Red Dragon | removal | A Dragon removal piece for punishing opposing pressure.
- 1 Breaching Dragonstorm | synergy | A synergy piece for the Dragon-focused game plan.
- 1 Crucible of Fire | synergy | A synergy piece that strengthens the Dragon plan.
- 1 Dracogenesis | synergy | A synergy piece for committing Dragons to the board.
- 1 Dragonkin Berserker | synergy | A synergy piece for the Dragon-focused creature base.
- 1 Dragonlord's Servant | synergy | A synergy piece for the deck's Dragon plan.
- 1 Dragonspeaker Shaman | synergy | A synergy piece for the deck's Dragon plan.
- 1 Dragonstorm | synergy | A synergy piece that emphasizes the Dragon theme.
- 1 Firespitter Whelp | synergy | A Dragon synergy piece that supports the tribe.
- 1 Last Light of Durin's Day | synergy | A synergy piece for the Dragon-focused strategy.
- 1 Minion of the Mighty | synergy | A synergy piece for getting the Dragon plan moving early.
- 1 Sarkhan's Triumph | synergy | A synergy piece for finding the right Dragon threat.
- 1 Shivan Devastator | synergy | A Dragon synergy piece that scales with the deck's mana plan.
- 1 Slumbering Dragon | synergy | A Dragon synergy piece for the creature-focused plan.
- 1 The Dragon-Kami Reborn // Dragon-Kami's Egg | synergy | A synergy piece that supports the Dragon strategy.
- 1 Ancient Bronze Dragon | threat | A major Dragon threat for closing games through combat.
- 1 Blast-Furnace Hellkite | threat | A Dragon threat that supports forceful attacks.
- 1 Hellkite Charger | threat | A Dragon threat for the aggressive combat plan.
- 1 Lathliss, Dragon Queen | threat | A Dragon threat that reinforces the tribal plan.
- 1 Scourge of the Throne | threat | A Dragon threat for decisive combat turns.
- 1 Stormbreath Dragon | threat | A Dragon threat that maintains aerial pressure.
- 1 Terror of Mount Velus | threat | A Dragon threat for explosive finishing attacks.
- 1 Thrakkus the Butcher | threat | A Dragon threat that supports the combat-focused plan.
- 1 Thunderbreak Regent | threat | A Dragon threat that rewards the tribal commitment.
- 1 Twinflame Tyrant | threat | A Dragon threat for powerful attacks.
- 1 Utvara Hellkite | threat | A Dragon threat for overwhelming the table.
- 1 Draconic Muralists | threat | A Dragon threat that adds to the deck's top end.
- 1 Balefire Dragon | wipe | A Dragon board wipe that can clear opposing pressure.
- 1 Draconic Intervention | wipe | A board wipe for resetting difficult boards.
- 1 Thundermaw Hellkite | wipe | A Dragon board wipe that also contributes to combat pressure.

</details>

### 15. delegated commander, owned first

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 225 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $0.00 to buy, $194.05 the whole deck.

**Summary:** This is a lifegain-focused Commander deck that develops its board, supports its central synergies, and keeps pressure off with targeted answers and reset buttons. It aims to close games through its creature threats once the board is stabilized. The deck gives up some flexibility for a direct, board-centered lifegain plan.

- [INFO] `bracket_prose_rules`: bracket 2 (Core, 8+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 3.08 over 63 nonland cards

<details><summary>The deck list</summary>

- 18 Plains | land | Provides white mana for the deck.
- 10 Swamp | land | Provides black mana for the deck.
- 1 Command Tower | land | Provides reliable mana for the deck.
- 1 Marsh Flats | land | Provides land access for the deck.
- 1 Path of Ancestry | land | Provides mana for the deck.
- 1 Plaza of Heroes | land | Provides mana for the deck.
- 1 Takenuma, Abandoned Mire | land | Provides black mana for the deck.
- 1 Castle Locthwain | land | Provides black mana for the deck.
- 1 War Room | land | Provides mana for the deck.
- 1 Windbrisk Heights | land | Provides white mana for the deck.
- 1 Sol Ring | ramp | Provides ramp for the deck.
- 1 Arcane Signet | ramp | Provides ramp for the deck.
- 1 Bender's Waterskin | ramp | Provides ramp for the deck.
- 1 Commander's Sphere | ramp | Provides ramp for the deck.
- 1 Fellwar Stone | ramp | Provides ramp for the deck.
- 1 Giada, Font of Hope | ramp | Provides ramp for the deck.
- 1 Lotho, Corrupt Shirriff | ramp | Provides ramp for the deck.
- 1 Oath of the Grey Host | ramp | Provides ramp for the deck.
- 1 Thought Vessel | ramp | Provides ramp for the deck.
- 1 Wayfarer's Bauble | ramp | Provides ramp for the deck.
- 1 Buster Sword | draw | Provides draw for the deck.
- 1 Call of the Ring | draw | Provides draw for the deck.
- 1 Exemplar of Light | draw | Provides draw for the deck.
- 1 Idol of Oblivion | draw | Provides draw for the deck.
- 1 Inspiring Overseer | draw | Provides draw for the deck.
- 1 Instant Ramen | draw | Provides draw for the deck.
- 1 Lembas | draw | Provides draw for the deck.
- 1 Night's Whisper | draw | Provides draw for the deck.
- 1 Puresteel Paladin | draw | Provides draw for the deck.
- 1 Skullclamp | draw | Provides draw for the deck.
- 1 Bastion Protector | interaction | Provides interaction for the deck.
- 1 Champion's Helm | interaction | Provides interaction for the deck.
- 1 Lightning Greaves | interaction | Provides interaction for the deck.
- 1 Sheltered by Ghosts | interaction | Provides interaction for the deck.
- 1 Swiftfoot Boots | interaction | Provides interaction for the deck.
- 1 Unbreakable Formation | interaction | Provides interaction for the deck.
- 1 Banishing Light | removal | Provides removal for the deck.
- 1 Bitter Triumph | removal | Provides removal for the deck.
- 1 Crib Swap | removal | Provides removal for the deck.
- 1 Dispatch | removal | Provides removal for the deck.
- 1 Generous Gift | removal | Provides removal for the deck.
- 1 Get Lost | removal | Provides removal for the deck.
- 1 Infernal Grasp | removal | Provides removal for the deck.
- 1 Swords to Plowshares | removal | Provides removal for the deck.
- 1 Austere Command | wipe | Provides a board wipe for the deck.
- 1 Fumigate | wipe | Provides a board wipe for the deck.
- 1 Vanquish the Horde | wipe | Provides a board wipe for the deck.
- 1 Adventurous Eater // Have a Bite | synergy | Fills a synergy role in the deck.
- 1 Aerith Gainsborough | synergy | Fills a synergy role in the deck.
- 1 Aettir and Priwen | synergy | Fills a synergy role in the deck.
- 1 Angel of Vitality | synergy | Fills a synergy role in the deck.
- 1 Compassionate Healer | synergy | Fills a synergy role in the deck.
- 1 Dancer's Chakrams | synergy | Fills a synergy role in the deck.
- 1 Elixir | synergy | Fills a synergy role in the deck.
- 1 Light of Promise | synergy | Fills a synergy role in the deck.
- 1 Night Nurse, Healer of Heroes | synergy | Fills a synergy role in the deck.
- 1 Prideful Feastling | synergy | Fills a synergy role in the deck.
- 1 Rosie Cotton of South Lane | synergy | Fills a synergy role in the deck.
- 1 Second Breakfast | synergy | Fills a synergy role in the deck.
- 1 Well-Worn Spatula | synergy | Fills a synergy role in the deck.
- 1 White Mage's Staff | synergy | Fills a synergy role in the deck.
- 1 Angel of Invention | threat | Serves as a threat for the deck.
- 1 Canyon Crawler | threat | Serves as a threat for the deck.
- 1 Dawnhand Eulogist | threat | Serves as a threat for the deck.
- 1 Foggy Swamp Hunters | threat | Serves as a threat for the deck.
- 1 Gollum, Silent Slinker // Meager Meal | threat | Serves as a threat for the deck.
- 1 Invisible Woman, Sue Storm | threat | Serves as a threat for the deck.
- 1 Lyra Dawnbringer | threat | Serves as a threat for the deck.
- 1 Minwu, White Mage | threat | Serves as a threat for the deck.
- 1 Rabaroo Troop | threat | Serves as a threat for the deck.
- 1 Shattered Angel | threat | Serves as a threat for the deck.
- 1 Sneering Shadewriter | threat | Serves as a threat for the deck.
- 1 Victory's Herald | threat | Serves as a threat for the deck.

</details>

### 16. a tight budget, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 228 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $22.34 to buy, $46.83 the whole deck.

**Summary:** This is a sacrifice-focused aristocrats deck that develops a creature board, supports that board with mana and cards, and turns its themed synergy pieces and threats into a path to victory. It has removal, protection, and reset buttons to keep opponents from taking over, but it gives up some individual card power for a focused, creature-based plan and a low-cost build.

- [INFO] `bracket_prose_rules`: bracket 2 (Core, 8+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [WARN] `not_owned`: Aron, Benalia's Ruin: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ayli, Eternal Pilgrim: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bartolomé del Presidio: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bastion of Remembrance: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Blood Artist: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Elas il-Kor, Sadistic Pilgrim: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Falkenrath Noble: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Meathook Massacre II: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Skullport Merchant: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Sivriss, Nightmare Speaker: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Woe Strider: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Yahenni, Undying Partisan: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Zulaport Cutthroat: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Baron Bertram Graywater: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Erebos, Bleak-Hearted: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Elder Arthur Maxson: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Lord of the Forsaken: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Master of Dark Rites: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Shadowheart, Dark Justiciar: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Shilgengar, Sire of Famine: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vampire Gourmand: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vengeful Bloodwitch: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vindictive Vampire: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 2.98 over 63 nonland cards

<details><summary>The deck list</summary>

- 10 Plains | land | Provides dependable white mana for the deck.
- 10 Swamp | land | Provides dependable black mana for the deck.
- 1 Command Tower | land | Provides flexible mana for the deck.
- 1 Exotic Orchard | land | Provides flexible mana for the deck.
- 1 Path of Ancestry | land | Adds a themed mana source.
- 1 Secluded Courtyard | land | Supports the creature-heavy mana base.
- 1 Evolving Wilds | land | Finds a basic land when needed.
- 1 Terramorphic Expanse | land | Finds a basic land when needed.
- 1 Ash Barrens | land | Helps smooth basic-land access.
- 1 Fabled Passage | land | Helps smooth basic-land access.
- 1 Escape Tunnel | land | Adds utility to the mana base.
- 1 Opal Palace | land | Adds utility to the mana base.
- 1 Thriving Moor | land | Provides black mana and color flexibility.
- 1 Vibrant Cityscape | land | Provides flexible mana for the deck.
- 1 Windbrisk Heights | land | Provides white mana with added utility.
- 1 Ifnir Deadlands | land | Provides black mana with added utility.
- 1 Sidequest: Catch a Fish // Cooking Campsite | land | Adds a flexible land slot to the mana base.
- 1 The Gold Saucer | land | Adds utility to the mana base.
- 1 Arcane Signet | ramp | Provides the listed ramp role.
- 1 Chromatic Lantern | ramp | Provides the listed ramp role.
- 1 Fellwar Stone | ramp | Provides the listed ramp role.
- 1 Sol Ring | ramp | Provides the listed ramp role.
- 1 Thought Vessel | ramp | Provides the listed ramp role.
- 1 Commander's Sphere | ramp | Provides the listed ramp role.
- 1 White Auracite | ramp | Provides the listed ramp role.
- 1 World Map | ramp | Provides the listed ramp role.
- 1 Inherited Envelope | ramp | Provides the listed ramp role.
- 1 Deadly Dispute | ramp | Provides the listed ramp role.
- 1 Cirith Ungol Patrol | draw | Provides the listed draw role.
- 1 Grave Venerations | draw | Provides the listed draw role.
- 1 Inspiring Overseer | draw | Provides the listed draw role.
- 1 Nasty End | draw | Provides the listed draw role.
- 1 Painful Truths | draw | Provides the listed draw role.
- 1 Skullclamp | draw | Provides the listed draw role.
- 1 Idol of Oblivion | draw | Provides the listed draw role.
- 1 Wall of Omens | draw | Provides the listed draw role.
- 1 Tome of Legends | draw | Provides the listed draw role.
- 1 Foot Chopper | draw | Provides the listed draw role.
- 1 Bastion Protector | interaction | Provides the listed interaction role.
- 1 Bronze Guardian | interaction | Provides the listed interaction role.
- 1 Frontline Medic | interaction | Provides the listed interaction role.
- 1 Gift of Immortality | interaction | Provides the listed interaction role.
- 1 Swiftfoot Boots | interaction | Provides the listed interaction role.
- 1 Together Forever | interaction | Provides the listed interaction role.
- 1 Bitter Triumph | removal | Provides the listed removal role.
- 1 Claim the Precious | removal | Provides the listed removal role.
- 1 Crib Swap | removal | Provides the listed removal role.
- 1 Deadly Precision | removal | Provides the listed removal role.
- 1 Destroy Evil | removal | Provides the listed removal role.
- 1 Gollum the Abandoned | removal | Provides the listed removal role.
- 1 Infernal Grasp | removal | Provides the listed removal role.
- 1 Palace Jailer | removal | Provides the listed removal role.
- 1 Archfiend of Ifnir | wipe | Provides the listed wipe role.
- 1 Austere Command | wipe | Provides the listed wipe role.
- 1 Dusk // Dawn | wipe | Provides the listed wipe role.
- 1 Gollum, Patient Plotter | synergy | Supports the deck's sacrifice-focused plan.
- 1 Aron, Benalia's Ruin | synergy | Supports the deck's sacrifice-focused plan.
- 1 Ayli, Eternal Pilgrim | synergy | Supports the deck's sacrifice-focused plan.
- 1 Bartolomé del Presidio | synergy | Supports the deck's sacrifice-focused plan.
- 1 Bastion of Remembrance | synergy | Supports the deck's sacrifice-focused plan.
- 1 Blood Artist | synergy | Supports the deck's sacrifice-focused plan.
- 1 Elas il-Kor, Sadistic Pilgrim | synergy | Supports the deck's sacrifice-focused plan.
- 1 Falkenrath Noble | synergy | Supports the deck's sacrifice-focused plan.
- 1 Meathook Massacre II | synergy | Supports the deck's sacrifice-focused plan.
- 1 Skullport Merchant | synergy | Supports the deck's sacrifice-focused plan.
- 1 Sivriss, Nightmare Speaker | synergy | Supports the deck's sacrifice-focused plan.
- 1 Woe Strider | synergy | Supports the deck's sacrifice-focused plan.
- 1 Yahenni, Undying Partisan | synergy | Supports the deck's sacrifice-focused plan.
- 1 Zulaport Cutthroat | synergy | Supports the deck's sacrifice-focused plan.
- 1 Bill the Pony | threat | Adds a threat to advance the deck's board presence.
- 1 Baron Bertram Graywater | threat | Adds a threat to advance the deck's board presence.
- 1 Erebos, Bleak-Hearted | threat | Adds a threat to advance the deck's board presence.
- 1 Elder Arthur Maxson | threat | Adds a threat to advance the deck's board presence.
- 1 Lord of the Forsaken | threat | Adds a threat to advance the deck's board presence.
- 1 Master of Dark Rites | threat | Adds a threat to advance the deck's board presence.
- 1 Shadowheart, Dark Justiciar | threat | Adds a threat to advance the deck's board presence.
- 1 Shilgengar, Sire of Famine | threat | Adds a threat to advance the deck's board presence.
- 1 Vampire Gourmand | threat | Adds a threat to advance the deck's board presence.
- 1 Vengeful Bloodwitch | threat | Adds a threat to advance the deck's board presence.
- 1 Vengeful Villagers | threat | Adds a threat to advance the deck's board presence.
- 1 Vindictive Vampire | threat | Adds a threat to advance the deck's board presence.

</details>

### 17. upgrade a precon, owned first

Format: Commander. Theme: goblins. Pool: owned_first. Shortlist: 196 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $73.85 to buy, $248.11 the whole deck.

**Summary:** This Goblin Storm upgrade keeps the precon’s creature-and-spell core, using Goblin synergy alongside ramp and draw to keep pressure on the table and turn its threat cards into the route to victory. Removal, wipes, and protective interaction back up that plan, while the tradeoff is remaining tied to the precon’s broader package instead of concentrating entirely on a rebuilt single-line strategy.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [WARN] `land_count`: 25 lands: the guide range for this format is 34 to 38
- [INFO] `curve_summary`: average mana value 2.49 over 74 nonland cards

<details><summary>The deck list</summary>

- 1 Ancestors' Aid | synergy | Retained as part of the precon's established spell-focused package.
- 1 Arena of Glory | land | Retained from the precon as a land.
- 1 Battle Hymn | ramp | Provides ramp.
- 1 Blasphemous Act | wipe | Provides a wipe.
- 1 Boggart Shenanigans | removal | Provides removal.
- 1 Castle Embereth | land | Retained from the precon as a land.
- 1 Chaos Warp | removal | Provides removal.
- 1 Conspicuous Snoop | synergy | Provides Goblin synergy.
- 1 Crimson Wisps | draw | Provides draw.
- 1 Daring Discovery | synergy | Retained as part of the precon's established package.
- 1 Den of the Bugbear | land | Retained from the precon as a land.
- 1 Dragon Fodder | synergy | Retained as part of the precon's Goblin package.
- 1 Dwarven Mine | land | Retained from the precon as a land.
- 1 Empty the Warrens | synergy | Retained as part of the precon's storm-focused package.
- 1 Expedite | draw | Provides draw.
- 1 Faithless Looting | draw | Provides draw.
- 1 Fists of Flame | draw | Provides draw.
- 1 Forgotten Cave | land | Retained from the precon as a land.
- 1 Frontline Heroism | interaction | Retained as part of the precon's established package.
- 1 Gempalm Incinerator | removal | Provides removal.
- 1 Glimpse the Impossible | ramp | Provides ramp.
- 1 Goblin Bombardment | removal | Provides removal.
- 1 Goblin Burrows | land | Retained from the precon as a land.
- 1 Goblin Bushwhacker | synergy | Provides Goblin synergy.
- 1 Goblin Chieftain | synergy | Provides Goblin synergy.
- 1 Goblin Dark-Dwellers | threat | Provides a threat.
- 1 Goblin Lackey | synergy | Provides Goblin synergy.
- 1 Goblin Matron | synergy | Provides Goblin synergy.
- 1 Goblin Negotiation | removal | Provides removal.
- 1 Goblin Trashmaster | removal | Provides removal.
- 1 Goblin Warchief | synergy | Provides Goblin synergy.
- 1 Grapeshot | removal | Provides removal.
- 1 Great Train Heist | ramp | Provides ramp.
- 1 Grenzo, Havoc Raiser | synergy | Provides Goblin synergy.
- 1 Haze of Rage | synergy | Retained as part of the precon's storm-focused package.
- 1 Hidden Volcano | land | Retained from the precon as a land.
- 1 Idol of Oblivion | draw | Provides draw.
- 1 Impact Tremors | synergy | Retained as part of the precon's Goblin package.
- 1 Kher Keep | land | Retained from the precon as a land.
- 1 Krenko's Command | synergy | Retained as part of the precon's Goblin package.
- 1 Krenko, Mob Boss | threat | Provides a threat.
- 1 Mana Geyser | ramp | Provides ramp.
- 1 Mogg War Marshal | synergy | Provides Goblin synergy.
- 14 Mountain | land | Retained from the precon as basic lands.
- 1 Past in Flames | synergy | Retained as part of the precon's storm-focused package.
- 1 Quest for the Goblin Lord | synergy | Provides Goblin synergy.
- 1 Reliquary Tower | land | Retained from the precon as a land.
- 1 Roaming Throne | synergy | Retained as part of the precon's Goblin package.
- 1 Ruby Medallion | ramp | Retained as part of the precon's spell-focused package.
- 1 Rundvelt Hordemaster | synergy | Provides Goblin synergy.
- 1 Seething Song | ramp | Provides ramp.
- 1 Siege-Gang Commander | removal | Provides removal.
- 1 Siege-Gang Lieutenant | removal | Provides removal.
- 1 Skirk Prospector | ramp | Provides ramp.
- 1 Skullclamp | draw | Provides draw.
- 1 Smoldering Crater | land | Retained from the precon as a land.
- 1 Sol Ring | ramp | Provides ramp.
- 1 Spreading Insurrection | synergy | Retained as part of the precon's established package.
- 1 Storm-Kiln Artist | ramp | Provides ramp.
- 1 Swiftfoot Boots | interaction | Provides interaction.
- 1 Throne of Eldraine | ramp | Provides ramp.
- 1 Vandalblast | wipe | Provides a wipe.
- 1 War Room | land | Retained from the precon as a land.
- 1 Wild Ride | synergy | Retained as part of the precon's established package.
- 1 Witch's Mark | draw | Provides draw.
- 1 Abrade | removal | Adds flexible removal.
- 1 Ancestral Anger | draw | Adds draw to the spell package.
- 1 Arcane Signet | ramp | Adds reliable ramp.
- 1 Battle-Scarred Goblin | removal | Adds a Goblin removal card.
- 1 Broadside Bombardiers | removal | Adds a Goblin removal card.
- 1 Burst Lightning | removal | Adds efficient removal.
- 1 Cunning Maneuver | draw | Adds draw to the spell package.
- 1 Darksteel Plate | interaction | Adds interaction for the deck's key creatures.
- 1 Fable of the Mirror-Breaker // Reflection of Kiki-Jiki | ramp | Adds ramp while staying within the deck's creature package.
- 1 Firebending Lesson | removal | Adds removal.
- 1 Goblin Fireleaper | removal | Adds a Goblin removal card.
- 1 Grishnákh, Brash Instigator | synergy | Adds Goblin synergy.
- 1 Lightning Bolt | removal | Adds efficient removal.
- 1 Lightning Greaves | interaction | Adds interaction for the deck's key creatures.
- 1 Mask of Memory | draw | Adds draw.
- 1 Moria Marauder | synergy | Adds Goblin synergy.
- 1 Pashalik Mons | removal | Adds a Goblin removal card.
- 1 Renegade Tactics | draw | Adds draw to the spell package.
- 1 Rush the Room | synergy | Adds synergy for the deck's focused plan.
- 1 Sazacap's Brew | draw | Adds draw to the spell package.
- 1 Seasoned Pyromancer | draw | Adds draw while supporting the creature package.

</details>

### 18. upgrade a precon, any card

Format: Commander. Theme: turtles. Pool: any_card. Shortlist: 303 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $214.96 to buy, $214.96 the whole deck.

**Summary:** Heroes in a Half Shell leads a Turtle Power deck that keeps the precon’s character-driven core while adding steadier mana, stronger card flow, and broadly useful answers. Build a Turtle-centered board, support it with the retained themed package, and convert that presence into sustained pressure. The deck gives up some raw speed and tightly focused competitive lines in exchange for keeping the precon’s broad cast, flavor, and varied game play.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [WARN] `land_count`: 23 lands: the guide range for this format is 34 to 38
- [INFO] `curve_summary`: average mana value 3.13 over 76 nonland cards

<details><summary>The deck list</summary>

- 1 Acidic Slime | other | Retained from the Turtle Power precon to preserve its established creature base.
- 1 April O'Neil, Live on the Scene | draw | Retained from the Turtle Power precon for card draw.
- 1 Arcade Cabinet | other | Retained from the Turtle Power precon to preserve its original support package.
- 1 Arcane Signet | ramp | Retained from the Turtle Power precon for ramp.
- 1 Ash Barrens | land | Retained from the Turtle Power precon as a land.
- 1 Assassin's Trophy | removal | Retained from the Turtle Power precon for removal.
- 1 Baxter, Fly in the Ointment | other | Retained from the Turtle Power precon to preserve its original character package.
- 1 Bebop, Skull & Crossbones | other | Retained from the Turtle Power precon to preserve its original character package.
- 1 Big Apple, 3 a.m. | land | Retained from the Turtle Power precon as a land.
- 1 Big Mother Mouser | other | Retained from the Turtle Power precon to preserve its original creature package.
- 1 Biogenic Ooze | other | Retained from the Turtle Power precon to preserve its original creature package.
- 1 Blasphemous Act | wipe | Retained from the Turtle Power precon as a board wipe.
- 1 Casey Jones, Back Alley Brute | other | Retained from the Turtle Power precon to preserve its original character package.
- 1 Chromatic Lantern | ramp | Retained from the Turtle Power precon for ramp.
- 1 Cinder Glade | land | Retained from the Turtle Power precon as a land.
- 1 Coin of Mastery | other | Retained from the Turtle Power precon to preserve its original support package.
- 1 Command Tower | land | Retained from the Turtle Power precon as a land.
- 1 Continue? | other | Retained from the Turtle Power precon to preserve its original support package.
- 1 Corpsejack Menace | other | Retained from the Turtle Power precon to preserve its original creature package.
- 1 Dimension X Pizzasaur | other | Retained from the Turtle Power precon to preserve its original creature package.
- 1 Donatello, the Brains | synergy | Retained from the Turtle Power precon for Turtle synergy.
- 1 Double Jump // Flying Kick | other | Retained from the Turtle Power precon to preserve its original support package.
- 1 Electric Seaweed | other | Retained from the Turtle Power precon to preserve its original creature package.
- 1 Endless Foot Assault | other | Retained from the Turtle Power precon to preserve its original support package.
- 1 Escape Tunnel | land | Retained from the Turtle Power precon as a land.
- 1 Everything Pizza | other | Retained from the Turtle Power precon to preserve its original support package.
- 1 Exploding Barrel | other | Retained from the Turtle Power precon to preserve its original support package.
- 1 Fast Forward | other | Retained from the Turtle Power precon to preserve its original support package.
- 1 Foot Chopper | other | Retained from the Turtle Power precon to preserve its original Equipment package.
- 1 Forest | land | Retained from the Turtle Power precon as a basic land.
- 1 Game Over | other | Retained from the Turtle Power precon to preserve its original support package.
- 1 Grand Coliseum | land | Retained from the Turtle Power precon as a land.
- 1 Harmonize | other | Retained from the Turtle Power precon to preserve its original support package.
- 1 Here Comes a New Hero! | other | Retained from the Turtle Power precon to preserve its original support package.
- 1 Hidden Hideout | land | Retained from the Turtle Power precon as a land.
- 1 High Score | other | Retained from the Turtle Power precon to preserve its original support package.
- 1 Hinterland Harbor | land | Retained from the Turtle Power precon as a land.
- 1 Irma, Part-Time Mutant | other | Retained from the Turtle Power precon to preserve its original character package.
- 1 Island | land | Retained from the Turtle Power precon as a basic land.
- 1 Krang, the All-Powerful | other | Retained from the Turtle Power precon to preserve its original creature package.
- 1 Leatherhead, Iron Gator | other | Retained from the Turtle Power precon to preserve its original character package.
- 1 Leonardo, the Balance | threat | Retained from the Turtle Power precon as a threat.
- 1 Lessons from Life | other | Retained from the Turtle Power precon to preserve its original support package.
- 1 Level Up | other | Retained from the Turtle Power precon to preserve its original Aura package.
- 1 Lita, Little Orphan Amphibian | synergy | Retained from the Turtle Power precon for Turtle synergy.
- 1 Michelangelo, the Heart | synergy | Retained from the Turtle Power precon for Turtle synergy.
- 1 Mole Module | other | Retained from the Turtle Power precon to preserve its original Vehicle package.
- 1 Mona Lisa, Science Geek | other | Retained from the Turtle Power precon to preserve its original character package.
- 1 Mountain | land | Retained from the Turtle Power precon as a basic land.
- 1 Ninja Pizza | other | Retained from the Turtle Power precon to preserve its original support package.
- 1 Path of Ancestry | land | Retained from the Turtle Power precon as a land.
- 1 Plains | land | Retained from the Turtle Power precon as a basic land.
- 1 Rain-Slicked Copse | land | Retained from the Turtle Power precon as a land.
- 1 Rat King, Pale Piper | other | Retained from the Turtle Power precon to preserve its original character package.
- 1 Ray Fillet, Wave Warrior | other | Retained from the Turtle Power precon to preserve its original character package.
- 1 Roadkill Rodney | other | Retained from the Turtle Power precon to preserve its original creature package.
- 1 Rocksteady, Mutant Marauder | other | Retained from the Turtle Power precon to preserve its original character package.
- 1 Rootbound Crag | land | Retained from the Turtle Power precon as a land.
- 1 Shellshock | other | Retained from the Turtle Power precon to preserve its original support package.
- 1 Shredder, Shadow Master | other | Retained from the Turtle Power precon to preserve its original character package.
- 1 Sol Ring | ramp | Retained from the Turtle Power precon for ramp.
- 1 Special Move | other | Retained from the Turtle Power precon to preserve its original support package.
- 1 Spire Garden | land | Retained from the Turtle Power precon as a land.
- 1 Splinter, the Mentor | other | Retained from the Turtle Power precon to preserve its original character package.
- 1 Steelbane Hydra | removal | Retained from the Turtle Power precon as a removal option.
- 1 Super Combo | other | Retained from the Turtle Power precon to preserve its original support package.
- 1 Swamp | land | Retained from the Turtle Power precon as a basic land.
- 1 Swift Demise | other | Retained from the Turtle Power precon to preserve its original support package.
- 1 Tempestra, Dame of Games | other | Retained from the Turtle Power precon to preserve its original character package.
- 1 Thriving Grove | land | Retained from the Turtle Power precon as a land.
- 1 Thriving Isle | land | Retained from the Turtle Power precon as a land.
- 1 Thriving Moor | land | Retained from the Turtle Power precon as a land.
- 1 Together Forever | other | Retained from the Turtle Power precon to preserve its original support package.
- 1 Tokka & Rahzar, Unsupervised | ramp | Retained from the Turtle Power precon for ramp.
- 1 Turtle Lair | land | Retained from the Turtle Power precon as a land.
- 1 Vanquish the Horde | wipe | Retained from the Turtle Power precon to preserve its original sweepers.
- 1 Vernal Fen | land | Retained from the Turtle Power precon as a land.
- 1 Vibrant Cityscape | land | Retained from the Turtle Power precon as a land.
- 1 Vigor | other | Retained from the Turtle Power precon to preserve its original creature package.
- 1 Voracious Hydra | other | Retained from the Turtle Power precon to preserve its original creature package.
- 1 Wave Goodbye | wipe | Retained from the Turtle Power precon to preserve its original sweepers.
- 1 Cultivate | ramp | Added to make the deck's mana development more dependable.
- 1 Farseek | ramp | Added as efficient early ramp.
- 1 Nature's Lore | ramp | Added as efficient early ramp.
- 1 Three Visits | ramp | Added as efficient early ramp.
- 1 Fellwar Stone | ramp | Added as compact artifact ramp.
- 1 Rhystic Study | draw | Added for ongoing card draw.
- 1 Mystic Remora | draw | Added for early card draw.
- 1 Skullclamp | draw | Added as a dedicated card-draw tool.
- 1 Garruk's Uprising | draw | Added for card draw alongside the creature plan.
- 1 Swords to Plowshares | removal | Added as efficient removal.
- 1 Beast Within | removal | Added as flexible removal.
- 1 Chaos Warp | removal | Added as flexible removal.
- 1 Generous Gift | removal | Added as flexible removal.
- 1 Anguished Unmaking | removal | Added as flexible removal.
- 1 Swan Song | interaction | Added as low-cost interaction.
- 1 Heroic Intervention | interaction | Added as protective interaction.
- 1 Lightning Greaves | interaction | Added as protective interaction for key creatures.
- 1 Turtle Power! | synergy | Added to strengthen the Turtle-focused theme.

</details>

