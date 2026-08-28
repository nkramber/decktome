# PR-8 deck gate

Run date: 2026-08-28. Card snapshot: 2026-08-24.

Verdict: PASS. 16 of 16 decks passed every block check, 0 invented names reached the user, and 0 summaries stated a false rule of the game. All three bars are zero tolerance.

## Summary

| Measure | Value |
|---|---|
| Prompts | 16 |
| Decks returned | 16 |
| Decks with no block finding | 16 |
| Invented names that reached the user | 0 |
| Decks that needed the repair turn | 0 |
| Summaries judged (F-26) | 16 |
| Summaries that state a rule of the game | 1 |
| Summaries that state a FALSE rule | 0 |
| Errors | 0 |
| Prompt version | 4 |
| Calls | 32 |
| Cost | $0.8776 |
| Time | 733 seconds |

## Findings by code

A block stops the deck. A warning and a note are reports.

| Code | Count |
|---|---|
| `not_owned` | 26 |
| `curve_summary` | 16 |
| `bracket_prose_rules` | 10 |
| `land_count` | 1 |

By severity: BLOCK 0. WARN 27. INFO 26. 

## Decks

### 1. lifegain Commander, any card

Format: Commander. Theme: lifegain. Pool: any_card. Shortlist: 303 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $394.23 to buy, $394.23 the whole deck.

**Summary:** This deck centers on repeated lifegain, developing a board of synergistic permanents while drawing into larger threats and keeping key opposing pieces in check. It wins by turning that established board into sustained pressure and attacking with its high-impact creatures. The tradeoff is that the deck is built to develop over time rather than race quickly, so it can be vulnerable when its board is repeatedly disrupted.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 3.57 over 63 nonland cards

<details><summary>The deck list</summary>

- 18 Plains | land | Provides a basic white land slot for the mana base.
- 10 Swamp | land | Provides a basic black land slot for the mana base.
- 1 Scoured Barrens | land | Provides a land slot for the mana base.
- 1 Shambling Vent | land | Provides a land slot for the mana base.
- 1 Vault of the Archangel | land | Provides a land slot for the mana base.
- 1 Restless Fortress | land | Provides a land slot for the mana base.
- 1 Radiant Fountain | land | Provides a land slot for the mana base.
- 1 Kabira Crossroads | land | Provides a land slot for the mana base.
- 1 High Market | land | Provides a land slot for the mana base.
- 1 Westvale Abbey // Ormendahl, Profane Prince | land | Provides a land slot for the mana base.
- 1 Altar of the Pantheon | ramp | Supports the deck's ramp package.
- 1 Angel of Indemnity | ramp | Supports the deck's ramp package.
- 1 Battle Angels of Tyr | ramp | Supports the deck's ramp package.
- 1 Beza, the Bounding Spring | ramp | Supports the deck's ramp package.
- 1 Colossal Plow | ramp | Supports the deck's ramp package.
- 1 Crypt Ghast | ramp | Supports the deck's ramp package.
- 1 Hierophant's Chalice | ramp | Supports the deck's ramp package.
- 1 Nuka-Cola Vending Machine | ramp | Supports the deck's ramp package.
- 1 Pristine Talisman | ramp | Supports the deck's ramp package.
- 1 The Celestus | ramp | Supports the deck's ramp package.
- 1 Alhammarret's Archive | draw | Supports the deck's card-draw package.
- 1 Archivist of Oghma | draw | Supports the deck's card-draw package.
- 1 Ayara, First of Locthwain | draw | Supports the deck's card-draw package.
- 1 Convalescent Care | draw | Supports the deck's card-draw package.
- 1 Cosmos Elixir | draw | Supports the deck's card-draw package.
- 1 Dawn of Hope | draw | Supports the deck's card-draw package.
- 1 Enduring Innocence | draw | Supports the deck's card-draw package.
- 1 Mangara, the Diplomat | draw | Supports the deck's card-draw package.
- 1 Sigarda's Splendor | draw | Supports the deck's card-draw package.
- 1 Well of Lost Dreams | draw | Supports the deck's card-draw package.
- 1 Alseid of Life's Bounty | interaction | Provides a flexible interaction piece.
- 1 Faith's Shield | interaction | Provides a flexible interaction piece.
- 1 Restoration Magic | interaction | Provides a flexible interaction piece.
- 1 Sword of Light and Shadow | interaction | Provides a flexible interaction piece.
- 1 Werefox Bodyguard | interaction | Provides a flexible interaction piece.
- 1 Enduring Angel // Angelic Enforcer | interaction | Provides a flexible interaction piece.
- 1 Aetherflux Reservoir | removal | Adds a removal option to the deck.
- 1 Ayli, Eternal Pilgrim | removal | Adds a removal option to the deck.
- 1 Cavalier of Night | removal | Adds a removal option to the deck.
- 1 Murderous Rider // Swift End | removal | Adds a removal option to the deck.
- 1 Nightmare's Thirst | removal | Adds a removal option to the deck.
- 1 Solitude | removal | Adds a removal option to the deck.
- 1 Umezawa's Jitte | removal | Adds a removal option to the deck.
- 1 Vona, Butcher of Magan | removal | Adds a removal option to the deck.
- 1 Ajani's Pridemate | synergy | Supports the lifegain-centered synergy plan.
- 1 Angel of Vitality | synergy | Supports the lifegain-centered synergy plan.
- 1 Angelic Accord | synergy | Supports the lifegain-centered synergy plan.
- 1 Blood Artist | synergy | Supports the lifegain-centered synergy plan.
- 1 Cleric Class | synergy | Supports the lifegain-centered synergy plan.
- 1 Cleric of Life's Bond | synergy | Supports the lifegain-centered synergy plan.
- 1 Cradle of Vitality | synergy | Supports the lifegain-centered synergy plan.
- 1 Heliod, Sun-Crowned | synergy | Supports the lifegain-centered synergy plan.
- 1 Indulging Patrician | synergy | Supports the lifegain-centered synergy plan.
- 1 Resplendent Angel | synergy | Supports the lifegain-centered synergy plan.
- 1 Righteous Valkyrie | synergy | Supports the lifegain-centered synergy plan.
- 1 Sanguine Bond | synergy | Supports the lifegain-centered synergy plan.
- 1 Vito, Thorn of the Dusk Rose | synergy | Supports the lifegain-centered synergy plan.
- 1 Voice of the Blessed | synergy | Supports the lifegain-centered synergy plan.
- 1 Archangel of Thune | threat | Provides a major threat for closing games.
- 1 Astarion, the Decadent | threat | Provides a major threat for closing games.
- 1 Attended Healer | threat | Provides a major threat for closing games.
- 1 Celestine, the Living Saint | threat | Provides a major threat for closing games.
- 1 Cliffhaven Vampire | threat | Provides a major threat for closing games.
- 1 Defiant Bloodlord | threat | Provides a major threat for closing games.
- 1 Divinity of Pride | threat | Provides a major threat for closing games.
- 1 Elenda, Saint of Dusk | threat | Provides a major threat for closing games.
- 1 Liesa, Forgotten Archangel | threat | Provides a major threat for closing games.
- 1 Lyra Dawnbringer | threat | Provides a major threat for closing games.
- 1 Nykthos Paragon | threat | Provides a major threat for closing games.
- 1 Valkyrie Harbinger | threat | Provides a major threat for closing games.
- 1 Ajani, Strength of the Pride | wipe | Provides a board-wipe option when the table gets ahead.
- 1 Fumigate | wipe | Provides a board-wipe option when the table gets ahead.
- 1 Kaya's Wrath | wipe | Provides a board-wipe option when the table gets ahead.

</details>

### 2. aristocrats Commander, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 178 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $0.00 to buy, $194.97 the whole deck.

**Summary:** The deck builds its board with creatures and artifact support, uses its synergy pieces to sustain an aristocrats-oriented game, and clears away opposing pieces with focused removal or broad sweepers. It wins by turning that developed board and its threats into steady pressure, while giving up some dedicated aristocrats depth for a larger package of card advantage and general-purpose answers.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 2.87 over 63 nonland cards

<details><summary>The deck list</summary>

- 12 Plains | land | Basic land for the mana base.
- 12 Swamp | land | Basic land for the mana base.
- 1 Ash Barrens | land | Flexible land slot.
- 1 Castle Locthwain | land | Black mana-base land.
- 1 Command Tower | land | Reliable commander-deck land.
- 1 Evolving Wilds | land | Fixing land for the mana base.
- 1 Fabled Passage | land | Fixing land for the mana base.
- 1 Marsh Flats | land | Fetch land for the mana base.
- 1 Minas Tirith | land | White mana-base land.
- 1 Path of Ancestry | land | Creature-focused mana-base land.
- 1 Plaza of Heroes | land | Legendary-focused mana-base land.
- 1 Scavenger Grounds | land | Utility land slot.
- 1 Takenuma, Abandoned Mire | land | Black utility land.
- 1 Terramorphic Expanse | land | Fixing land for the mana base.
- 1 Arcane Signet | ramp | Efficient mana support.
- 1 Astral Cornucopia | ramp | Artifact mana support.
- 1 Bender's Waterskin | ramp | Artifact mana support.
- 1 Commander's Sphere | ramp | Mana support for the deck.
- 1 Deadly Dispute | ramp | Supports the mana plan.
- 1 Fellwar Stone | ramp | Efficient artifact mana support.
- 1 Lotho, Corrupt Shirriff | ramp | Creature-based mana support.
- 1 Relic of Legends | ramp | Artifact mana support.
- 1 Sol Ring | ramp | Efficient artifact mana support.
- 1 Thought Vessel | ramp | Artifact mana support.
- 1 Call of the Ring | draw | Ongoing card advantage.
- 1 Idol of Oblivion | draw | Artifact-based card advantage.
- 1 Inspiring Overseer | draw | Creature-based card advantage.
- 1 Lembas | draw | Artifact-based card advantage.
- 1 Mask of Memory | draw | Equipment-based card advantage.
- 1 Night's Whisper | draw | Efficient card advantage.
- 1 Puresteel Paladin | draw | Creature-based card advantage.
- 1 Skullclamp | draw | Equipment-based card advantage.
- 1 Tome of Legends | draw | Repeatable artifact card advantage.
- 1 Wall of Omens | draw | Creature-based card advantage.
- 1 Ahriman | draw | Creature-based card advantage.
- 1 Buster Sword | draw | Equipment-based card advantage.
- 1 Cirith Ungol Patrol | draw | Creature-based card advantage.
- 1 Exemplar of Light | draw | Creature-based card advantage.
- 1 Folk Hero | draw | Ongoing card advantage.
- 1 Foot Chopper | draw | Equipment-based card advantage.
- 1 Grave Venerations | draw | Ongoing card advantage.
- 1 June, Bounty Hunter | draw | Creature-based card advantage.
- 1 Kingpin's Enforcers | draw | Creature-based card advantage.
- 1 Massacre Girl, Known Killer | draw | Creature-based card advantage.
- 1 Nasty End | draw | Efficient card advantage.
- 1 Rat King, Pale Piper | draw | Creature-based card advantage.
- 1 Secret Rendezvous | draw | Additional card advantage.
- 1 Stone of Erech | draw | Artifact-based card advantage.
- 1 The Sackville-Bagginses | draw | Creature-based card advantage.
- 1 Vanquisher's Banner | draw | Ongoing artifact card advantage.
- 1 Bastion Protector | interaction | Protective interaction for the commander plan.
- 1 Boromir, Warden of the Tower | interaction | Creature-based interaction.
- 1 Clever Concealment | interaction | Protective interaction.
- 1 Lightning Greaves | interaction | Protective equipment interaction.
- 1 Swiftfoot Boots | interaction | Protective equipment interaction.
- 1 Take Up the Shield | interaction | Protective interaction.
- 1 Angel of Serenity | removal | Creature-based removal.
- 1 Bitter Triumph | removal | Efficient removal.
- 1 Claim the Precious | removal | Focused removal.
- 1 Crib Swap | removal | Flexible removal.
- 1 Generous Gift | removal | Flexible removal.
- 1 Infernal Grasp | removal | Efficient removal.
- 1 Swords to Plowshares | removal | Efficient removal.
- 1 Witch-king of Angmar | removal | Creature-based removal.
- 1 Fiend Hunter | removal | Additional creature-based removal.
- 1 Austere Command | wipe | Flexible board reset.
- 1 Dusk // Dawn | wipe | Board-reset option.
- 1 Fumigate | wipe | Broad board reset.
- 1 Arcade Cabinet | synergy | Supports the aristocrats synergy plan.
- 1 Gollum, Patient Plotter | synergy | Supports the aristocrats synergy plan.
- 1 Gríma Wormtongue | synergy | Supports the aristocrats synergy plan.
- 1 Heirloom Auntie | synergy | Supports the aristocrats synergy plan.
- 1 Joo Dee, One of Many | synergy | Supports the aristocrats synergy plan.
- 1 Nimble Hobbit | synergy | Supports the aristocrats synergy plan.
- 1 Phantom Train | synergy | Supports the aristocrats synergy plan.
- 1 Bill the Pony | threat | Creature threat for the win plan.
- 1 Vengeful Villagers | threat | Creature threat for the win plan.

</details>

### 3. artifacts Commander, bracket 4

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 299 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $1348.90 to buy, $1348.90 the whole deck.

**Summary:** This deck builds an artifact-heavy board quickly, then converts that development into pressure with large artifact threats while maintaining cards and disruptive answers. It wins by establishing a commanding battlefield and pushing through its artifact finishers, with broad reset buttons available when opponents get ahead. The tradeoff is a strong dependence on keeping its artifact infrastructure in place, so opposing pressure on that board can slow its momentum.

- [INFO] `curve_summary`: average mana value 4.03 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Academy Ruins | land | Land slot for the artifact-focused mana base.
- 1 Archway of Innovation | land | Land slot for the artifact-focused mana base.
- 1 Blinkmoth Nexus | land | Land slot for the artifact-focused mana base.
- 1 Buried Ruin | land | Land slot for the artifact-focused mana base.
- 1 Conqueror's Galleon // Conqueror's Foothold | land | Land slot for the artifact-focused mana base.
- 1 Fomori Vault | land | Land slot for the artifact-focused mana base.
- 1 Glimmervoid | land | Land slot for the artifact-focused mana base.
- 1 Hall of Tagsin | land | Land slot for the artifact-focused mana base.
- 1 Inventors' Fair | land | Land slot for the artifact-focused mana base.
- 1 Mishra's Factory | land | Land slot for the artifact-focused mana base.
- 1 Mishra's Foundry | land | Land slot for the artifact-focused mana base.
- 1 Mishra's Workshop | land | Land slot for the artifact-focused mana base.
- 1 Otawara, Soaring City | land | Land slot for the artifact-focused mana base.
- 1 Power Depot | land | Land slot for the artifact-focused mana base.
- 1 Roadside Reliquary | land | Land slot for the artifact-focused mana base.
- 1 Spire of Industry | land | Land slot for the artifact-focused mana base.
- 1 The Monumental Facade | land | Land slot for the artifact-focused mana base.
- 1 The Mycosynth Gardens | land | Land slot for the artifact-focused mana base.
- 1 Urza's Factory | land | Land slot for the artifact-focused mana base.
- 1 Urza's Saga | land | Land slot for the artifact-focused mana base.
- 1 Urza's Workshop | land | Land slot for the artifact-focused mana base.
- 1 Uthros, Titanic Godcore | land | Land slot for the artifact-focused mana base.
- 14 Island | land | Basic land choice for the mana base.
- 1 Mox Opal | ramp | Efficient ramp for an artifact-heavy start.
- 1 Metalworker | ramp | Powerful artifact-focused ramp piece.
- 1 Krark-Clan Ironworks | ramp | Artifact-based ramp for explosive turns.
- 1 Moonsnare Prototype | ramp | Early ramp that fits the artifact shell.
- 1 Chief Engineer | ramp | Ramp support for the artifact plan.
- 1 Grand Architect | ramp | Ramp support for the artifact plan.
- 1 Karn, Legacy Reforged | ramp | Ramp that also fits the artifact theme.
- 1 Tezzeret the Seeker | ramp | Ramp support for high-impact artifact turns.
- 1 Inspiring Statuary | ramp | Artifact-focused ramp support.
- 1 The Mightstone and Weakstone | ramp | Ramp piece for the deck's larger plays.
- 1 Braided Net // Braided Quipu | draw | Artifact-based draw for maintaining resources.
- 1 Thoughtcast | draw | Efficient draw for the artifact-heavy shell.
- 1 Thought Monitor | draw | Artifact draw piece that advances the board.
- 1 Sai, Master Thopterist | draw | Draw support within the artifact plan.
- 1 Thopter Spy Network | draw | Ongoing draw support for the artifact strategy.
- 1 Vedalken Archmage | draw | Draw support for an artifact-dense deck.
- 1 Thirst for Knowledge | draw | Flexible draw to keep the deck moving.
- 1 Reverse Engineer | draw | Artifact-focused draw support.
- 1 Tezzeret, Artifice Master | draw | Draw engine that fits the artifact shell.
- 1 Forensic Gadgeteer | draw | Draw support for the artifact plan.
- 1 Assert Authority | interaction | Interaction for protecting the deck's game plan.
- 1 Metallic Rebuke | interaction | Efficient interaction in an artifact-heavy deck.
- 1 Stoic Rebuttal | interaction | Reliable blue interaction.
- 1 Disruption Protocol | interaction | Interaction that fits the artifact shell.
- 1 Darksteel Forge | interaction | Protective interaction for the artifact board.
- 1 Jin-Gitaxias, Progress Tyrant | interaction | High-impact interaction for the late game.
- 1 Aether Spellbomb | removal | Artifact-based removal that supports the shell.
- 1 Arcum Dagsson | removal | Repeatable removal option in the artifact plan.
- 1 Contagion Clasp | removal | Artifact removal that contributes to board control.
- 1 Lux Cannon | removal | Artifact-based removal for persistent control.
- 1 Portal to Phyrexia | removal | High-impact artifact removal.
- 1 Spine of Ish Sah | removal | Versatile artifact removal piece.
- 1 Resculpt | removal | Flexible blue removal.
- 1 Ravenform | removal | Additional removal for difficult permanents.
- 1 Contagion Engine | wipe | Artifact board wipe for resetting crowded boards.
- 1 Nevinyrral's Disk | wipe | Broad artifact-based board wipe.
- 1 Oblivion Stone | wipe | Flexible board wipe for emergency resets.
- 1 Clock of Omens | synergy | Supports artifact synergies across the deck.
- 1 Foundry Inspector | synergy | Core artifact synergy support.
- 1 Etherium Sculptor | synergy | Core artifact synergy support.
- 1 Emry, Lurker of the Loch | synergy | Artifact synergy piece for value-focused turns.
- 1 Mystic Forge | synergy | Artifact synergy engine for sustained development.
- 1 Unwinding Clock | synergy | Supports the deck's artifact-centric plan.
- 1 Voltaic Key | synergy | Low-cost artifact synergy support.
- 1 Manifold Key | synergy | Low-cost artifact synergy support.
- 1 Scrap Trawler | synergy | Artifact synergy piece for long games.
- 1 Coretapper | synergy | Artifact synergy support for the deck's engines.
- 1 Panharmonicon | synergy | Synergy piece that rewards artifact development.
- 1 Mirrorworks | synergy | Artifact synergy engine for building advantage.
- 1 Simulacrum Synthesizer | synergy | Artifact synergy payoff for developing the board.
- 1 Transmute Artifact | synergy | Artifact synergy tool for assembling the plan.
- 1 Arcbound Crusher | threat | Artifact threat that pressures opponents.
- 1 Cyberdrive Awakener | threat | Artifact threat that turns board development into pressure.
- 1 Darksteel Juggernaut | threat | Large artifact threat for closing games.
- 1 Kappa Cannoneer | threat | Efficient artifact threat for applying pressure.
- 1 Kuldotha Forgemaster | threat | High-impact artifact threat.
- 1 Master Transmuter | threat | Artifact threat that supports the deck's endgame.
- 1 Metalwork Colossus | threat | Large artifact threat for closing games.
- 1 Phyrexian Metamorph | threat | Flexible artifact threat for the board state.
- 1 Karn, Scion of Urza | threat | Threat that contributes to the artifact plan.
- 1 Threefold Thunderhulk | threat | Large artifact threat for closing games.
- 1 Traxos, Scourge of Kroog | threat | Efficient artifact threat.
- 1 Lodestone Golem | threat | Artifact threat that adds battlefield pressure.

</details>

### 4. dinosaur tribal, bracket 2

Format: Commander. Theme: dinosaurs. Pool: any_card. Shortlist: 267 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $246.63 to buy, $246.63 the whole deck.

**Summary:** This deck builds mana early, develops a Dinosaur board, and uses Gishath alongside large Dinosaur threats to pressure the table through combat. Tribal support and creature-based draw help keep the main plan moving, while removal, protective interaction, and a few reset buttons provide backup when opponents get ahead. It gives up some flexibility for its heavy creature focus and can be slower when its larger threats do not arrive on schedule.

- [INFO] `bracket_prose_rules`: bracket 2 (Core, 8+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 4.05 over 63 nonland cards

<details><summary>The deck list</summary>

- 11 Forest | land | Basic land that supports the deck's mana base.
- 7 Plains | land | Basic land that supports the deck's mana base.
- 6 Mountain | land | Basic land that supports the deck's mana base.
- 1 Command Tower | land | Land included for the deck's mana base.
- 1 Path of Ancestry | land | Land included for the deck's mana base.
- 1 Evolving Wilds | land | Land included for the deck's mana base.
- 1 Terramorphic Expanse | land | Land included for the deck's mana base.
- 1 Canopy Vista | land | Land included for the deck's mana base.
- 1 Cinder Glade | land | Land included for the deck's mana base.
- 1 Clifftop Retreat | land | Land included for the deck's mana base.
- 1 Rootbound Crag | land | Land included for the deck's mana base.
- 1 Sacred Foundry | land | Land included for the deck's mana base.
- 1 Stomping Ground | land | Land included for the deck's mana base.
- 1 Sunpetal Grove | land | Land included for the deck's mana base.
- 1 Temple Garden | land | Land included for the deck's mana base.
- 1 Arcane Signet | ramp | Early ramp to help deploy the deck's larger cards.
- 1 Cultivate | ramp | Ramp that supports reaching the deck's expensive Dinosaur cards.
- 1 Farseek | ramp | Ramp included to build mana early.
- 1 Drover of the Mighty | ramp | Creature-based ramp that fits the Dinosaur-focused plan.
- 1 Nature's Lore | ramp | Ramp included to build mana early.
- 1 Kodama's Reach | ramp | Ramp that supports reaching the deck's larger cards.
- 1 Rampant Growth | ramp | Ramp included to build mana early.
- 1 Sol Ring | ramp | Early ramp to help deploy the deck's larger cards.
- 1 Thunderherd Migration | ramp | Dinosaur-themed ramp for the main plan.
- 1 Topiary Stomper | ramp | Dinosaur ramp that contributes to the creature plan.
- 1 Beast Whisperer | draw | Draw support attached to a creature for the creature-heavy plan.
- 1 Garruk's Uprising | draw | Draw support for a deck built around large creatures.
- 1 Guardian Project | draw | Draw support for the creature-focused plan.
- 1 Harmonize | draw | Straightforward draw support to keep cards coming.
- 1 Kutzil, Malamet Exemplar | draw | Creature-based draw support for the attacking plan.
- 1 Ripjaw Raptor | draw | Dinosaur draw support that fits the tribe.
- 1 Rishkar's Expertise | draw | Draw support that rewards the deck's large threats.
- 1 Shamanic Revelation | draw | Draw support for a board built around creatures.
- 1 Toski, Bearer of Secrets | draw | Creature-based draw support for the combat plan.
- 1 Vanquisher's Banner | draw | Tribal draw support for the Dinosaur theme.
- 1 Akroma's Will | interaction | Interaction that helps support the creature-based game plan.
- 1 Boros Charm | interaction | Interaction included to back up the board presence.
- 1 Dawn's Truce | interaction | Interaction included to protect the main plan.
- 1 Heroic Intervention | interaction | Interaction included to protect the creature board.
- 1 Lightning Greaves | interaction | Equipment interaction for supporting an important creature.
- 1 Swiftfoot Boots | interaction | Equipment interaction for supporting an important creature.
- 1 Apex Altisaur | removal | Dinosaur removal that also contributes a creature to the plan.
- 1 Bronzebeak Foragers | removal | Dinosaur removal that fits the tribal theme.
- 1 Burning Sun's Avatar | removal | Dinosaur removal that also serves the big-creature plan.
- 1 Deathgorge Scavenger | removal | Dinosaur removal for a flexible creature slot.
- 1 Itzquinth, Firstborn of Gishath | removal | Dinosaur removal tied directly to the deck's theme.
- 1 Ravenous Sailback | removal | Dinosaur removal that keeps the deck creature-focused.
- 1 Thrashing Brontodon | removal | Dinosaur removal for a flexible answer slot.
- 1 Trumpeting Carnosaur | removal | Dinosaur removal that also adds a substantial creature.
- 1 Austere Command | wipe | Board wipe for resetting difficult board states.
- 1 Blasphemous Act | wipe | Board wipe for handling crowded battlefields.
- 1 Wakening Sun's Avatar | wipe | Dinosaur-themed board wipe that fits the deck's creature plan.
- 1 Commune with Dinosaurs | synergy | Dinosaur synergy that supports the tribal plan.
- 1 Dinosaur Stampede | synergy | Dinosaur synergy for the combat-focused plan.
- 1 Hunting Velociraptor | synergy | Dinosaur synergy attached to a tribal creature.
- 1 Kinjalli's Caller | synergy | Tribal synergy that supports deploying Dinosaur creatures.
- 1 Marauding Raptor | synergy | Dinosaur synergy for the creature-heavy plan.
- 1 Otepec Huntmaster | synergy | Tribal synergy that supports the Dinosaur plan.
- 1 Raptor Companion | synergy | Low-commitment Dinosaur creature for tribal synergy.
- 1 Raptor Hatchling | synergy | Dinosaur synergy that builds the creature theme.
- 1 Regal Imperiosaur | synergy | Dinosaur synergy for the tribe's creature plan.
- 1 Relentless Raptor | synergy | Dinosaur synergy attached to an on-theme creature.
- 1 Siegehorn Ceratops | synergy | Dinosaur synergy that keeps the deck centered on the tribe.
- 1 Sky Terror | synergy | Dinosaur synergy for the combat plan.
- 1 Sun-Collared Raptor | synergy | Dinosaur synergy attached to a tribal creature.
- 1 Territorial Hammerskull | synergy | Dinosaur synergy that supports the creature-focused strategy.
- 1 Ancient Brontodon | threat | Large Dinosaur threat for closing games through combat.
- 1 Carnage Tyrant | threat | Dinosaur threat that advances the large-creature plan.
- 1 Etali, Primal Storm | threat | Legendary Dinosaur threat for the deck's top end.
- 1 Ghalta, Primal Hunger | threat | Large legendary Dinosaur threat for combat finishes.
- 1 Ghalta, Stampede Tyrant | threat | Legendary Dinosaur threat for the deck's top end.
- 1 Goring Ceratops | threat | Dinosaur threat that reinforces the attacking plan.
- 1 Quartzwood Crasher | threat | Dinosaur threat for the combat-focused strategy.
- 1 Regisaur Alpha | threat | Dinosaur threat that supports the tribal creature plan.
- 1 Thundering Spineback | threat | Dinosaur threat for building a powerful board.
- 1 Tyrranax Rex | threat | Large Dinosaur threat for finishing through combat.
- 1 Verdant Sun's Avatar | threat | Dinosaur threat that adds a substantial tribal creature.
- 1 Zetalpa, Primal Dawn | threat | Legendary Dinosaur threat for the deck's late game.

</details>

### 5. blink Commander, owned first

Format: Commander. Theme: blink. Pool: owned_first. Shortlist: 168 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $0.00 to buy, $152.07 the whole deck.

**Summary:** This deck develops a white creature board, leans on blink-focused pieces alongside creature-based draw and removal, and protects its key creatures while controlling the table. It wins by turning repeated value and a steady board of Angels, Humans, artifacts, and other creatures into sustained pressure. It gives up a dedicated combo finish for a slower, board-centered game that depends on keeping its creatures in play.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 3.00 over 63 nonland cards

<details><summary>The deck list</summary>

- 22 Plains | land | Basic white land for the mana base.
- 1 Abandoned Air Temple | land | Land slot that supports the mana base.
- 1 Adventurer's Inn | land | Land slot that supports the mana base.
- 1 Ash Barrens | land | Land slot that supports the mana base.
- 1 Command Tower | land | Land slot that supports the mana base.
- 1 Evolving Wilds | land | Land slot that supports the mana base.
- 1 Fabled Passage | land | Land slot that supports the mana base.
- 1 Minas Tirith | land | Land slot that supports the mana base.
- 1 Path of Ancestry | land | Land slot that supports the mana base.
- 1 Plaza of Heroes | land | Land slot that supports the mana base.
- 1 Secluded Courtyard | land | Land slot that supports the mana base.
- 1 Shire Terrace | land | Land slot that supports the mana base.
- 1 Temple of the False God | land | Land slot that supports the mana base.
- 1 Terramorphic Expanse | land | Land slot that supports the mana base.
- 1 Windbrisk Heights | land | Land slot that supports the mana base.
- 1 Arcane Signet | ramp | Efficient artifact ramp for early development.
- 1 Bender's Waterskin | ramp | Artifact ramp that helps develop the board.
- 1 Commander's Sphere | ramp | Artifact ramp for steady development.
- 1 Fellwar Stone | ramp | Low-cost artifact ramp.
- 1 Relic of Legends | ramp | Artifact ramp for the mana base.
- 1 Sol Ring | ramp | Fast artifact ramp.
- 1 Sword of the Animist | ramp | Equipment-based ramp for a creature-heavy deck.
- 1 Thought Vessel | ramp | Artifact ramp for steady development.
- 1 Wayfarer's Bauble | ramp | Early ramp that helps establish mana.
- 1 White Lotus Tile | ramp | Artifact ramp for the deck's development.
- 1 Adventurer's Airship | draw | Artifact draw option for sustained resources.
- 1 Buster Sword | draw | Equipment draw option for the creature plan.
- 1 Champions of Minas Tirith | draw | Creature-based draw that fits the board-focused plan.
- 1 Crown of Gondor | draw | Artifact draw option for sustained resources.
- 1 Diary of Dreams | draw | Artifact draw option for sustained resources.
- 1 Exemplar of Light | draw | Creature-based draw that fits the blink shell.
- 1 Faramir, Field Commander | draw | Creature-based draw for the board-focused plan.
- 1 Idol of Oblivion | draw | Artifact draw option for sustained resources.
- 1 Inspiring Overseer | draw | Creature-based draw that works well in a blink shell.
- 1 Lembas | draw | Artifact draw option for sustained resources.
- 1 Mask of Memory | draw | Equipment draw option for the creature plan.
- 1 Mirror of Galadriel | draw | Artifact draw option for sustained resources.
- 1 Puresteel Paladin | draw | Creature-based draw that supports the equipment package.
- 1 Skullclamp | draw | Equipment draw option for the creature-heavy build.
- 1 Wall of Omens | draw | Creature-based draw that fits the blink shell.
- 1 Angel of Condemnation | removal | Creature-based removal that suits the blink plan.
- 1 Angel of Sanctions | removal | Creature-based removal that suits the blink plan.
- 1 Angel of Serenity | removal | Creature-based removal that suits the blink plan.
- 1 Banishing Light | removal | Flexible removal for troublesome permanents.
- 1 Contagion Clasp | removal | Artifact-based removal option.
- 1 Crib Swap | removal | Instant-speed removal option.
- 1 Destroy Evil | removal | Flexible instant-speed removal.
- 1 Dispatch | removal | Low-cost instant-speed removal.
- 1 Fiend Hunter | removal | Creature-based removal that suits the blink plan.
- 1 Generous Gift | removal | Flexible instant-speed removal.
- 1 Get Lost | removal | Flexible instant-speed removal.
- 1 Journey to Nowhere | removal | Enchantment-based removal option.
- 1 Oblation | removal | Flexible instant-speed removal.
- 1 Palace Jailer | removal | Creature-based removal that suits the blink plan.
- 1 Swords to Plowshares | removal | Efficient instant-speed removal.
- 1 Bastion Protector | interaction | Creature-based protection for the commander and board.
- 1 Boromir, Warden of the Tower | interaction | Creature-based interaction for protecting the plan.
- 1 Bronze Guardian | interaction | Artifact creature that adds protective interaction.
- 1 Champion's Helm | interaction | Equipment-based protection for an important creature.
- 1 Clever Concealment | interaction | Protective interaction for preserving the board.
- 1 Darksteel Plate | interaction | Equipment-based protection for a key creature.
- 1 Frontline Medic | interaction | Creature-based protective interaction.
- 1 Gift of Immortality | interaction | Protective enchantment for a key creature.
- 1 Lightning Greaves | interaction | Equipment-based protection for important creatures.
- 1 Reprieve | interaction | Instant interaction that buys time.
- 1 Slip On the Ring | interaction | Blink-focused instant interaction.
- 1 Swiftfoot Boots | interaction | Equipment-based protection for important creatures.
- 1 Together Forever | interaction | Protective enchantment for the creature plan.
- 1 Austere Command | wipe | Flexible board wipe for resetting difficult boards.
- 1 Dismantling Wave | wipe | Board wipe for handling opposing permanents.
- 1 Dusk // Dawn | wipe | Board wipe that also supports a creature-heavy plan.
- 1 Fumigate | wipe | Board wipe for recovering from wide boards.
- 1 Martial Coup | wipe | Board wipe that also contributes to board presence.
- 1 Vanquish the Horde | wipe | Board wipe for catching up on the table.
- 1 Ennis, Debate Moderator | synergy | Blink-focused synergy piece from the available pool.
- 1 Flickerwisp | synergy | Core blink-focused synergy creature.
- 1 Jocasta, Automaton Avenger | synergy | Blink-focused synergy piece from the available pool.
- 1 Personify | synergy | Blink-focused synergy option from the available pool.

</details>

### 6. Modern tempo, tournament

Format: Modern. Theme: tempo. Pool: any_card. Shortlist: 167 names.

Cards: 60 main, 15 sideboard. Repair turn: false. Block findings: 0.

Cost: $502.62 to buy, $502.62 the whole deck.

**Summary:** This blue-red tempo deck establishes pressure with early creatures, then uses cheap draw, removal, and interaction to keep that pressure ahead. It wins by maintaining a threat on the table while disrupting the opponent’s attempts to stabilize, with the temporal package offering another way to press an advantage. The tradeoff is that the deck is built for pace and efficiency rather than a large late-game board presence, so careful sequencing and holding up the right answers matter.

- [INFO] `curve_summary`: average mana value 2.22 over 36 nonland cards

<details><summary>The deck list</summary>

- 8 Island | land | Basic blue land for the deck's mana base.
- 4 Mountain | land | Basic red land for the deck's mana base.
- 4 Steam Vents | land | Blue-red land for the two-color mana base.
- 4 Scalding Tarn | land | Land slot supporting the mana base.
- 2 Shivan Reef | land | Blue-red land for the two-color mana base.
- 2 Stormcarved Coast | land | Blue-red land for the two-color mana base.
- 4 Consider | draw | Low-cost draw spell that supports a spell-heavy tempo plan.
- 2 Preordain | draw | Early draw spell for finding the needed part of the plan.
- 4 Counterspell | interaction | Core interaction for protecting the tempo plan.
- 2 Spell Pierce | interaction | Low-cost interaction that fits the tempo shell.
- 4 Lightning Bolt | removal | Efficient removal for clearing opposing threats.
- 2 Pongify | removal | Removal option in blue for problematic creatures.
- 2 Into the Flood Maw | removal | Removal spell that helps keep opponents off balance.
- 2 Temporal Mastery | synergy | Synergy card for the deck's temporal package.
- 2 Temporal Trespass | synergy | Synergy card that complements Temporal Mastery.
- 4 Ragavan, Nimble Pilferer | threat | Creature threat that gives the deck an early board presence.
- 4 Ledger Shredder | threat | Creature threat that supports the deck's pressure plan.
- 4 Faerie Mastermind | threat | Creature threat that contributes to the deck's evasive pressure.

</details>

### 7. Modern burn, casual

Format: Modern. Theme: burn. Pool: any_card. Shortlist: 290 names.

Cards: 60 main, 15 sideboard. Repair turn: false. Block findings: 0.

Cost: $69.95 to buy, $69.95 the whole deck.

**Summary:** This deck takes a direct mono-red burn approach: apply early pressure with damage-focused spells, use removal to clear a path, and keep opponents under strain with a substantial threat package. It aims to finish games before the opponent can stabilize, while the sideboard adds more damage, disruption, and a reset option. The trade-off is a streamlined proactive strategy with less flexibility than a broader multicolor deck.

- [INFO] `curve_summary`: average mana value 2.78 over 36 nonland cards

<details><summary>The deck list</summary>

- 24 Mountain | land | The basic red land base for a mono-red burn deck.
- 4 Ancestral Anger | draw | Draw support that helps keep the burn plan supplied.
- 2 Risk Factor | draw | Additional draw support for maintaining pressure.
- 2 Chandra, Dressed to Kill | ramp | Ramp support for deploying the deck's red spells and threats.
- 4 Lightning Bolt | removal | A primary removal spell for the direct-damage plan.
- 2 Lightning Strike | removal | Additional removal that supports the burn-focused strategy.
- 4 Eidolon of the Great Revel | synergy | A synergy piece for the deck's aggressive red game plan.
- 4 Guttersnipe | synergy | A synergy threat that rewards the deck's spell-heavy approach.
- 4 Ashcloud Phoenix | threat | A threat that adds board pressure alongside the burn package.
- 4 Hazoret the Fervent | threat | A resilient threat for closing games under pressure.
- 4 Sunspine Lynx | threat | A threat that advances the deck's aggressive pressure plan.
- 2 Torbran, Thane of Red Fell | threat | A threat that supports the deck's red damage-focused theme.

</details>

### 8. Modern lifegain, FNM

Format: Modern. Theme: lifegain. Pool: any_card. Shortlist: 301 names.

Cards: 60 main, 15 sideboard. Repair turn: false. Block findings: 0.

Cost: $96.80 to buy, $96.80 the whole deck.

**Summary:** This deck establishes a white-black lifegain board with early synergy pieces, then turns that setup into pressure through a steady creature threat package while keeping opposing threats in check with removal. It wins primarily by building a board that makes the lifegain plan matter and attacking through the resulting threats. The tradeoff is that the deck spends early turns assembling its pieces, so it can be less explosive when its lifegain synergies do not come together.

- [INFO] `curve_summary`: average mana value 3.28 over 36 nonland cards

<details><summary>The deck list</summary>

- 10 Plains | land | Basic land slot in the white-black mana base.
- 8 Swamp | land | Basic land slot in the white-black mana base.
- 4 Scoured Barrens | land | Land slot supporting the deck's two-color plan.
- 2 Shambling Vent | land | Land slot supporting the deck's two-color plan.
- 4 Inspiring Overseer | draw | Dedicated draw piece that also fits the white creature base.
- 2 Dawn of Hope | draw | Dedicated draw support for longer games.
- 2 Orzhov Keyrune | ramp | Ramp support for the deck's white-black plan.
- 4 Murderous Rider // Swift End | removal | Main-deck removal that keeps the deck interactive.
- 2 Solitude | removal | Additional main-deck removal for difficult threats.
- 4 Soul Warden | synergy | Core synergy piece for the lifegain plan.
- 4 Ajani's Pridemate | synergy | Core synergy piece for the lifegain plan.
- 4 Attended Healer | threat | A primary threat that fits the lifegain strategy.
- 4 Bloodbond Vampire | threat | A primary threat for the deck's lifegain plan.
- 2 Nykthos Paragon | threat | A larger threat for closing games after setting up.
- 2 Cliffhaven Vampire | threat | A threat that supports the white-black lifegain plan.
- 2 Angel of Invention | threat | A top-end threat for finishing games.

</details>

### 9. Standard midrange, FNM

Format: Standard. Theme: midrange. Pool: any_card. Shortlist: 164 names.

Cards: 60 main, 15 sideboard. Repair turn: false. Block findings: 0.

Cost: $300.40 to buy, $300.40 the whole deck.

**Summary:** This black-green midrange deck develops its mana, uses a steady draw package and removal suite to keep the table manageable, then turns a dense creature package into pressure. It wins through its larger threats and board-focused creatures after trading resources early. It gives up some speed and keeps several specialized answers, protection effects, and sweepers in the sideboard rather than the main deck.

- [WARN] `land_count`: 28 lands: the guide range for this format is 20 to 27
- [INFO] `curve_summary`: average mana value 3.47 over 32 nonland cards

<details><summary>The deck list</summary>

- 4 Forest | land | Basic land allocation for the mana base.
- 4 Swamp | land | Basic land allocation for the mana base.
- 4 Blooming Marsh | land | Black-green land allocation for the mana base.
- 4 Deathcap Glade | land | Black-green land allocation for the mana base.
- 4 Underground Mortuary | land | Black-green land allocation for the mana base.
- 4 Overgrown Tomb | land | Black-green land allocation for the mana base.
- 2 Llanowar Elves | ramp | Early ramp for the midrange plan.
- 2 Darkstar Augur | draw | Creature-based draw for sustained resources.
- 2 Phyrexian Arena | draw | Dedicated draw support for longer games.
- 2 Unholy Annex // Ritual Chamber | draw | Additional draw support in the main plan.
- 2 Bitter Triumph | removal | Efficient removal in the main deck.
- 2 Nowhere to Run | removal | Removal that supports the midrange plan.
- 2 Assassin's Trophy | removal | Flexible removal allocation.
- 4 Lord Skitter, Sewer King | synergy | Creature-based synergy piece for the board-focused plan.
- 2 Midnight Reaper | synergy | Creature-based synergy support for extended games.
- 2 Massacre Girl, Known Killer | synergy | Synergy creature that reinforces the midrange shell.
- 4 Goldvein Hydra | threat | Large creature threat for closing games.
- 3 Vein Ripper | threat | Creature threat for applying pressure.
- 3 Vaultborn Tyrant | threat | Top-end creature threat for the midrange plan.
- 2 Ojer Kaslem, Deepest Growth // Temple of Cultivation | threat | Legendary creature threat with a land face.
- 2 Aclazotz, Deepest Betrayal // Temple of the Dead | threat | Legendary creature threat with a land face.

</details>

### 10. Standard aggro, tournament

Format: Standard. Theme: aggro. Pool: any_card. Shortlist: 294 names.

Cards: 60 main, 15 sideboard. Repair turn: false. Block findings: 0.

Cost: $125.56 to buy, $125.56 the whole deck.

**Summary:** This red-white aggro deck leads with a steady stream of threats and uses its synergy cards to make each attack more punishing. Removal and interaction keep opposing resistance manageable, while the draw package helps maintain pressure after the first wave. It wins by staying proactive and forcing the opponent to answer the board, giving up the slower, heavier game plans and reset tools of more controlling builds.

- [INFO] `curve_summary`: average mana value 3.06 over 36 nonland cards

<details><summary>The deck list</summary>

- 4 Mountain | land | Forms part of the red portion of the land base.
- 4 Plains | land | Forms part of the white portion of the land base.
- 4 Sacred Foundry | land | Supports the deck's red-white land base.
- 4 Inspiring Vantage | land | Supports the deck's red-white land base.
- 4 Elegant Parlor | land | Supports the deck's red-white land base.
- 4 Sundown Pass | land | Supports the deck's red-white land base.
- 4 Fugitive Codebreaker | draw | Provides draw support while fitting the aggressive plan.
- 2 Reckless Lackey | draw | Adds more draw support for games that go longer.
- 4 Boros Charm | interaction | Provides efficient interaction for protecting the attack plan.
- 2 Sheltered by Ghosts | interaction | Adds interaction that supports the proactive game plan.
- 4 Emeritus of Truce // Swords to Plowshares | removal | Fills a removal role while staying aligned with the deck's colors.
- 4 Harsh Annotation | removal | Provides dedicated removal to clear a path for threats.
- 4 Warleader's Call | synergy | Supplies the core synergy element for a creature-focused attack plan.
- 4 Bedhead Beastie | threat | Provides a proactive threat for the aggro plan.
- 4 Edgewall Pack | threat | Adds another full set of threats to pressure opponents.
- 2 Frilled Sparkshooter | threat | Rounds out the threat base with more aggressive creatures.
- 2 Teapot Slinger | threat | Completes the proactive threat package.

</details>

### 11. Commander with a locked card

Format: Commander. Theme: sacrifice. Pool: any_card. Shortlist: 296 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $397.85 to buy, $397.85 the whole deck.

**Summary:** This Karlov deck plays a sacrifice-focused game, developing resources, keeping cards flowing, and using its answers to maintain a favorable table. It wins by building pressure around its sacrifice plan and deploying its larger threats, with wipes available when the board becomes crowded. The tradeoff is that much of the deck is committed to that central plan, so it is strongest when its sacrifice pieces and supporting cards are working together.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 3.17 over 63 nonland cards

<details><summary>The deck list</summary>

- 15 Swamp | land | Basic land for the deck’s black mana base.
- 11 Plains | land | Basic land for the deck’s white mana base.
- 1 Bojuka Bog | land | Listed land supporting the mana base.
- 1 Command Tower | land | Listed land supporting the mana base.
- 1 Evolving Wilds | land | Listed land supporting the mana base.
- 1 Exotic Orchard | land | Listed land supporting the mana base.
- 1 High Market | land | Listed land supporting the sacrifice plan.
- 1 Myriad Landscape | land | Listed land supporting the mana base.
- 1 Phyrexian Tower | land | Listed land supporting the sacrifice plan.
- 1 Reliquary Tower | land | Listed land supporting the mana base.
- 1 Terramorphic Expanse | land | Listed land supporting the mana base.
- 1 Westvale Abbey // Ormendahl, Profane Prince | land | Listed land supporting the sacrifice-focused build.
- 1 Ahriman | draw | Listed draw card that supports the deck’s card flow.
- 1 Baron Bertram Graywater | draw | Listed draw card that supports the deck’s card flow.
- 1 Bushmeat Poacher | draw | Listed draw card for the sacrifice-focused plan.
- 1 Corrupted Conviction | draw | Listed draw card for the sacrifice-focused plan.
- 1 Disciple of Bolas | draw | Listed draw card for the sacrifice-focused plan.
- 1 Relic Vial | draw | Listed draw card that supports the deck’s card flow.
- 1 Shadowheart, Dark Justiciar | draw | Listed draw card that supports the deck’s card flow.
- 1 Smothering Abomination | draw | Listed draw card for the sacrifice-focused plan.
- 1 Vampiric Rites | draw | Listed draw card for the sacrifice-focused plan.
- 1 Village Rites | draw | Listed draw card for the sacrifice-focused plan.
- 1 Sol Ring | ramp | Required card retained in the deck as a ramp slot.
- 1 Ashnod's Altar | ramp | Listed ramp piece for the sacrifice-focused plan.
- 1 Ashnod, Flesh Mechanist | ramp | Listed ramp piece supporting resource development.
- 1 Crowded Crypt | ramp | Listed ramp piece supporting resource development.
- 1 Pawn of Ulamog | ramp | Listed ramp piece for the sacrifice-focused plan.
- 1 Phyrexian Altar | ramp | Listed ramp piece for the sacrifice-focused plan.
- 1 Pitiless Plunderer | ramp | Listed ramp piece for the sacrifice-focused plan.
- 1 Priest of Forgotten Gods | ramp | Listed ramp piece for the sacrifice-focused plan.
- 1 Sifter of Skulls | ramp | Listed ramp piece for the sacrifice-focused plan.
- 1 Skullport Merchant | ramp | Listed ramp piece supporting resource development.
- 1 Attrition | removal | Listed removal piece for answering opposing cards.
- 1 Ayli, Eternal Pilgrim | removal | Listed removal piece for answering opposing cards.
- 1 Bone Shards | removal | Listed removal piece for answering opposing cards.
- 1 Dictate of Erebos | removal | Listed removal piece for the sacrifice-focused plan.
- 1 Eaten Alive | removal | Listed removal piece for answering opposing cards.
- 1 Grave Pact | removal | Listed removal piece for the sacrifice-focused plan.
- 1 Teysa, Orzhov Scion | removal | Listed removal piece for answering opposing cards.
- 1 Yawgmoth, Thran Physician | removal | Listed removal piece for answering opposing cards.
- 1 Cartel Aristocrat | interaction | Listed interaction supporting the sacrifice-focused plan.
- 1 Fanatical Devotion | interaction | Listed interaction for protecting the deck’s plan.
- 1 Flare of Fortitude | interaction | Listed interaction for protecting the deck’s plan.
- 1 Gift of Doom | interaction | Listed interaction for protecting the deck’s plan.
- 1 Nightmare Shepherd | interaction | Listed interaction supporting the sacrifice-focused plan.
- 1 Promise of Tomorrow | interaction | Listed interaction supporting the sacrifice-focused plan.
- 1 Liliana, Dreadhorde General | wipe | Listed wipe for resetting crowded boards.
- 1 The Meathook Massacre | wipe | Listed wipe for resetting crowded boards.
- 1 Toxic Deluge | wipe | Listed wipe for resetting crowded boards.
- 1 Blood Host | threat | Listed threat that advances the deck’s battlefield pressure.
- 1 Felisa, Fang of Silverquill | threat | Listed threat that advances the deck’s battlefield pressure.
- 1 Ghoulcaller Gisa | threat | Listed threat that advances the deck’s battlefield pressure.
- 1 Liesa, Forgotten Archangel | threat | Listed threat that advances the deck’s battlefield pressure.
- 1 Liliana, Heretical Healer // Liliana, Defiant Necromancer | threat | Listed threat that advances the deck’s battlefield pressure.
- 1 Mondrak, Glory Dominus | threat | Listed threat that advances the deck’s battlefield pressure.
- 1 Ratadrabik of Urborg | threat | Listed threat that advances the deck’s battlefield pressure.
- 1 Razaketh, the Foulblooded | threat | Listed threat that advances the deck’s battlefield pressure.
- 1 Requiem Angel | threat | Listed threat that advances the deck’s battlefield pressure.
- 1 Sidisi, Undead Vizier | threat | Listed threat that advances the deck’s battlefield pressure.
- 1 Thallid Omnivore | threat | Listed threat that advances the deck’s battlefield pressure.
- 1 Vindictive Vampire | threat | Listed threat that advances the deck’s battlefield pressure.
- 1 Altar of Dementia | synergy | Listed synergy piece for the sacrifice-focused plan.
- 1 Bartolomé del Presidio | synergy | Listed synergy piece for the sacrifice-focused plan.
- 1 Bastion of Remembrance | synergy | Listed synergy piece for the sacrifice-focused plan.
- 1 Carrion Feeder | synergy | Listed synergy piece for the sacrifice-focused plan.
- 1 Elas il-Kor, Sadistic Pilgrim | synergy | Listed synergy piece for the sacrifice-focused plan.
- 1 Fleshtaker | synergy | Listed synergy piece for the sacrifice-focused plan.
- 1 Hidden Stockpile | synergy | Listed synergy piece for the sacrifice-focused plan.
- 1 Open the Graves | synergy | Listed synergy piece for the sacrifice-focused plan.
- 1 Pitiless Pontiff | synergy | Listed synergy piece for the sacrifice-focused plan.
- 1 Spawning Pit | synergy | Listed synergy piece for the sacrifice-focused plan.
- 1 Viscera Seer | synergy | Listed synergy piece for the sacrifice-focused plan.
- 1 Woe Strider | synergy | Listed synergy piece for the sacrifice-focused plan.
- 1 Yahenni, Undying Partisan | synergy | Listed synergy piece for the sacrifice-focused plan.
- 1 Zulaport Cutthroat | synergy | Listed synergy piece for the sacrifice-focused plan.

</details>

### 12. Commander on a budget

Format: Commander. Theme: tokens. Pool: any_card. Shortlist: 296 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $43.60 to buy, $43.60 the whole deck.

**Summary:** This deck develops a wide token-focused board around Adeline, then uses its creature threats and token synergies to pressure opponents through combat. It has ample inexpensive support to keep cards and mana flowing, plus targeted answers and board resets when the table gets ahead. It gives up premium standalone power for a focused, board-centered plan, so its best games come from establishing and maintaining a large token presence.

- [INFO] `bracket_prose_rules`: bracket 2 (Core, 8+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 3.27 over 63 nonland cards

<details><summary>The deck list</summary>

- 36 Plains | land | Provides a no-cost basic-land foundation for the deck.
- 1 Angelic Sell-Sword | draw | Fills a low-cost draw slot.
- 1 Bygone Bishop | draw | Adds another inexpensive draw piece.
- 1 Dawn of Hope | draw | Supports the deck's card-flow package at low cost.
- 1 Faramir, Field Commander | draw | Contributes to the token-focused draw package.
- 1 Glimmer Seeker | draw | Provides a budget draw slot.
- 1 Hoofprints of the Stag | draw | Adds inexpensive card-flow support.
- 1 Idol of Oblivion | draw | Supplies draw support for the board-building plan.
- 1 Platoon Dispenser | draw | Adds a draw card that also fits the artifact-heavy support suite.
- 1 Staff of the Storyteller | draw | Provides another low-cost draw option.
- 1 Valiant Rescuer | draw | Rounds out the budget draw package.
- 1 Ainok Strike Leader | interaction | Adds a creature-based interaction piece.
- 1 Basri Ket | interaction | Provides flexible interaction in a low-cost slot.
- 1 Basri, Tomorrow's Champion | interaction | Adds another budget interaction option.
- 1 Lena, Selfless Champion | interaction | Supports the creature-based interaction suite.
- 1 Rootborn Defenses | interaction | Provides efficient interaction for the wide-board plan.
- 1 Teyo, the Shieldmage | interaction | Rounds out the interaction package at minimal cost.
- 1 Aerial Assault | removal | Provides efficient targeted removal.
- 1 Banishing Slash | removal | Adds a low-cost removal spell.
- 1 Battle Menu | removal | Fills a flexible budget removal slot.
- 1 Citizen's Crowbar | removal | Adds artifact-based removal support.
- 1 Conversion Chamber | removal | Provides a very inexpensive removal option.
- 1 Kellan's Lightblades | removal | Adds efficient instant-speed removal.
- 1 Path to Redemption | removal | Supplies another low-cost targeted answer.
- 1 Trostani's Judgment | removal | Rounds out the removal package with a budget spell.
- 1 Ceaseless Conflict | wipe | Provides a low-cost board-reset option.
- 1 Crisis of Conscience | wipe | Adds another inexpensive wipe for crowded boards.
- 1 Descend upon the Sinful | wipe | Completes the board-wipe suite.
- 1 Bucknard's Everfull Purse | ramp | Adds a budget artifact ramp piece.
- 1 Charisma Bobblehead | ramp | Provides inexpensive artifact-based ramp.
- 1 Coin of Mastery | ramp | Fills a low-cost ramp slot.
- 1 Collector's Vault | ramp | Adds artifact ramp support.
- 1 Currency Converter | ramp | Provides another budget ramp option.
- 1 Druidic Satchel | ramp | Supports the mana-development package at low cost.
- 1 Goldvein Pick | ramp | Adds inexpensive equipment-based ramp.
- 1 Idol of False Gods | ramp | Provides a budget artifact ramp slot.
- 1 Karn, Living Legacy | ramp | Adds ramp from a permanent that fits the deck's support plan.
- 1 Keeper of the Accord | ramp | Rounds out the ramp package with a creature-based option.
- 1 Aligned Heart | synergy | Supports the token-focused synergy package.
- 1 Anafenza, Unyielding Lineage | synergy | Adds a low-cost synergy creature.
- 1 Animation Module | synergy | Provides artifact-based support for the token plan.
- 1 Anointer Priest | synergy | Adds an inexpensive creature synergy piece.
- 1 Automated Assembly Line | synergy | Supports the deck's artifact and token themes.
- 1 Cat Collector | synergy | Adds a creature that fits the token-focused shell.
- 1 Cathar's Call | synergy | Provides a low-cost enchantment synergy piece.
- 1 Clarion Spirit | synergy | Adds a budget creature synergy card.
- 1 Divine Visitation | synergy | Provides a powerful token-focused synergy piece at a low listed cost.
- 1 Eyes in the Skies | synergy | Adds inexpensive token-plan support.
- 1 Felidar Retreat | synergy | Provides enchantment-based token synergy.
- 1 Hero of Precinct One | synergy | Adds a creature that supports the deck's broad-board plan.
- 1 Intangible Virtue | synergy | Provides efficient support for a token-focused board.
- 1 Mavren Fein, Dusk Apostle | synergy | Rounds out the creature-based synergy package.
- 1 Ajani's Chosen | threat | Adds a budget threat that fits the enchantment-heavy support suite.
- 1 Archangel Elspeth | threat | Provides a planeswalker threat for the board-focused plan.
- 1 Archon of Sun's Grace | threat | Adds an inexpensive threat for the enchantment package.
- 1 Attended Healer | threat | Provides a low-cost creature threat.
- 1 Basri's Lieutenant | threat | Adds a creature threat that complements the wide-board strategy.
- 1 Cemetery Protector | threat | Provides a standalone creature threat.
- 1 Defiler of Faith | threat | Adds a low-cost creature threat.
- 1 Dragonback Lancer | threat | Provides another budget creature threat.
- 1 Drogskol Cavalry | threat | Adds a higher-impact creature threat at low cost.
- 1 Emeria Angel | threat | Provides a creature threat that fits the land-heavy build.
- 1 Gideon, Ally of Zendikar | threat | Adds a planeswalker threat for the board-building plan.
- 1 God-Eternal Oketra | threat | Rounds out the threat package with a resilient-looking centerpiece.

</details>

### 13. owned first, and the commander is not owned

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 225 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $0.00 to buy, $199.66 the whole deck.

**Summary:** Karlov leads a white-black lifegain deck that develops its mana, establishes a resilient creature and equipment board, and uses lifegain-focused synergy to keep pressure on the table. It wins by protecting Karlov and its major threats long enough to turn a stable board into a closing attack, while its removal and sweepers prevent opponents from running away with the game. The tradeoff is that the deck is committed to creatures and themed support pieces, so it is less focused on a single fast finish.

- JUDGE [true]: "Karlov leads a white-black lifegain deck". Karlov of the White Orchard is a legendary creature and a legal commander for a white-black deck, so this assertion about which card can lead the deck is accurate.
- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [WARN] `not_owned`: Karlov of the Ghost Council: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.11 over 63 nonland cards

<details><summary>The deck list</summary>

- 15 Plains | land | Forms the primary white portion of the mana base.
- 12 Swamp | land | Forms the primary black portion of the mana base.
- 1 Command Tower | land | Provides flexible commander-focused mana fixing.
- 1 Marsh Flats | land | Supports reliable access to the deck's two colors.
- 1 Fabled Passage | land | Adds flexible mana fixing to the land base.
- 1 Evolving Wilds | land | Adds another flexible fixing land.
- 1 Ash Barrens | land | Helps smooth early land access.
- 1 Castle Locthwain | land | A black land option with added utility.
- 1 Takenuma, Abandoned Mire | land | A black land option with added utility.
- 1 War Room | land | A utility land for the deck's resource plan.
- 1 Windbrisk Heights | land | A white utility land that rounds out the mana base.
- 1 Sol Ring | ramp | Efficient early mana acceleration.
- 1 Arcane Signet | ramp | Reliable color fixing and acceleration.
- 1 Bender's Waterskin | ramp | An artifact ramp piece for developing the board.
- 1 Blitzball | ramp | Provides artifact-based acceleration.
- 1 Commander's Sphere | ramp | Fixes mana while supporting the deck's resource plan.
- 1 Fellwar Stone | ramp | A compact mana rock for early development.
- 1 Inherited Envelope | ramp | Adds another artifact acceleration piece.
- 1 Lotho, Corrupt Shirriff | ramp | A black-aligned ramp option for the deck.
- 1 Relic of Legends | ramp | Provides additional mana acceleration.
- 1 Wayfarer's Bauble | ramp | Supports early mana development.
- 1 Buster Sword | draw | Equipment-based card advantage for the deck.
- 1 Call of the Ring | draw | A black card-advantage piece.
- 1 Exemplar of Light | draw | A white creature that contributes to card flow.
- 1 Idol of Oblivion | draw | Artifact-based card advantage.
- 1 Inspiring Overseer | draw | A white creature that helps keep resources flowing.
- 1 Lembas | draw | A compact artifact draw option.
- 1 Mask of Memory | draw | Equipment-based card selection and advantage.
- 1 Night's Whisper | draw | Efficient black card draw.
- 1 Skullclamp | draw | A strong equipment-based draw engine.
- 1 Tome of Legends | draw | A persistent artifact draw source.
- 1 Bastion Protector | interaction | Helps protect the commander and key creatures.
- 1 Champion's Helm | interaction | Equipment protection for Karlov or a key threat.
- 1 Clever Concealment | interaction | A protective instant for preserving the board.
- 1 Darksteel Plate | interaction | Durable equipment protection for an important creature.
- 1 Lightning Greaves | interaction | Efficient equipment protection for Karlov.
- 1 Swiftfoot Boots | interaction | Additional equipment protection for key creatures.
- 1 Banishing Light | removal | A flexible white answer to opposing permanents.
- 1 Bitter Triumph | removal | A versatile black removal spell.
- 1 Crib Swap | removal | An instant-speed white removal option.
- 1 Generous Gift | removal | A broad instant-speed answer.
- 1 Infernal Grasp | removal | Efficient black creature removal.
- 1 Path to Exile | removal | A premium white removal spell.
- 1 Swords to Plowshares | removal | An efficient white answer to creatures.
- 1 Stroke of Midnight | removal | Flexible white permanent removal.
- 1 Adventurous Eater // Have a Bite | synergy | A lifegain-oriented synergy piece.
- 1 Aerith Gainsborough | synergy | Supports the deck's lifegain-focused creature plan.
- 1 Aettir and Priwen | synergy | Equipment synergy that supports Karlov's plan.
- 1 Angel of Vitality | synergy | A lifegain-focused Angel for the creature core.
- 1 Compassionate Healer | synergy | A dedicated lifegain synergy creature.
- 1 Dancer's Chakrams | synergy | An equipment synergy piece for the board.
- 1 Elixir | synergy | An artifact that supports the deck's synergy package.
- 1 Kor Firewalker | synergy | A lifegain-oriented creature for the deck's core plan.
- 1 Light of Promise | synergy | Builds on the deck's lifegain theme.
- 1 Scarblade's Malice | synergy | A black synergy spell for the deck's central plan.
- 1 Second Breakfast | synergy | A lifegain-themed instant for the synergy package.
- 1 The Darkness Crystal | synergy | An artifact synergy piece that supports the deck's theme.
- 1 Well-Worn Spatula | synergy | Equipment support for the creature-heavy strategy.
- 1 White Mage's Staff | synergy | An equipment synergy piece for the lifegain shell.
- 1 Angel of Invention | threat | A white Angel that serves as a meaningful battlefield threat.
- 1 Canyon Crawler | threat | A creature threat that helps apply board pressure.
- 1 Dawnhand Eulogist | threat | A black creature threat for the deck's combat plan.
- 1 Foggy Swamp Hunters | threat | A creature threat that broadens the board presence.
- 1 Gollum, Silent Slinker // Meager Meal | threat | A black threat with an attached adventure option.
- 1 Lyra Dawnbringer | threat | A high-impact Angel threat.
- 1 Minwu, White Mage | threat | A white legendary threat for the creature suite.
- 1 Rabaroo Troop | threat | A creature threat for advancing the board.
- 1 Reaping Willow | threat | A larger creature threat for the midgame.
- 1 Shattered Angel | threat | An Angel threat that fits the deck's theme.
- 1 Sneering Shadewriter | threat | A black threat that adds pressure.
- 1 Victory's Herald | threat | An Angel threat that strengthens the top end.
- 1 Austere Command | wipe | A flexible white board reset.
- 1 Fumigate | wipe | A white board wipe for catching up on the table.
- 1 Vanquish the Horde | wipe | An additional white reset when the board gets crowded.

</details>

### 14. the user delegates the commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 244 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $947.46 to buy, $947.46 the whole deck.

**Summary:** Atarka, World Render leads a focused Dragon board that builds resources early, then turns a sequence of large threats into sustained combat pressure. The deck wins by establishing a dense Dragon presence and closing through Atarka-backed attacks, with Dragon-themed answers and reset tools clearing resistance when needed. It gives up some flexibility and low-cost board development in exchange for committing heavily to expensive creatures and a direct combat plan.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 4.38 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Ancient Tomb | land | Land slot for the mana base.
- 1 Boseiju, Who Endures | land | Land slot for the mana base.
- 1 Cavern of Souls | land | Land slot for the mana base.
- 1 Cinder Glade | land | Land slot for the mana base.
- 1 Command Tower | land | Land slot for the mana base.
- 1 Crucible of the Spirit Dragon | land | Land slot for the mana base.
- 1 Evolving Wilds | land | Land slot for the mana base.
- 1 Exotic Orchard | land | Land slot for the mana base.
- 1 Fabled Passage | land | Land slot for the mana base.
- 6 Forest | land | Basic land foundation for the mana base.
- 1 Haven of the Spirit Dragon | land | Land slot for the mana base.
- 1 Mana Confluence | land | Land slot for the mana base.
- 1 Myriad Landscape | land | Land slot for the mana base.
- 5 Mountain | land | Basic land foundation for the mana base.
- 1 Nykthos, Shrine to Nyx | land | Land slot for the mana base.
- 1 Path of Ancestry | land | Land slot for the mana base.
- 1 Reflecting Pool | land | Land slot for the mana base.
- 1 Rootbound Crag | land | Land slot for the mana base.
- 1 Rogue's Passage | land | Land slot for the mana base.
- 1 Stomping Ground | land | Land slot for the mana base.
- 1 Temple of the Dragon Queen | land | Land slot for the mana base.
- 1 Temple of the False God | land | Land slot for the mana base.
- 1 Terramorphic Expanse | land | Land slot for the mana base.
- 1 Three Tree City | land | Land slot for the mana base.
- 1 War Room | land | Land slot for the mana base.
- 1 Wooded Foothills | land | Land slot for the mana base.
- 1 Yavimaya, Cradle of Growth | land | Land slot for the mana base.
- 1 Ancient Copper Dragon | ramp | Dragon-based resource acceleration.
- 1 Atsushi, the Blazing Sky | ramp | Dragon-based resource acceleration.
- 1 Carnelian Orb of Dragonkind | ramp | Resource acceleration for the deck's larger cards.
- 1 Dragon's Hoard | ramp | Dragon-themed resource acceleration.
- 1 Ganax, Astral Hunter | ramp | Dragon-based resource acceleration.
- 1 Goldspan Dragon | ramp | Dragon-based resource acceleration.
- 1 Jade Orb of Dragonkind | ramp | Resource acceleration for the deck's larger cards.
- 1 Klauth, Unrivaled Ancient | ramp | Dragon-based resource acceleration.
- 1 Old Gnawbone | ramp | Dragon-based resource acceleration.
- 1 Savage Ventmaw | ramp | Dragon-based resource acceleration.
- 1 Avaricious Dragon | draw | Dragon-based card advantage support.
- 1 Beast Whisperer | draw | Card advantage support for the creature-heavy plan.
- 1 Dragon Mage | draw | Dragon-based card advantage support.
- 1 Dragonborn Champion | draw | Card advantage support for the combat plan.
- 1 Elemental Bond | draw | Card advantage support for the creature plan.
- 1 Garruk's Uprising | draw | Card advantage support for the creature plan.
- 1 Guardian Project | draw | Card advantage support for the creature-heavy plan.
- 1 Return of the Wildspeaker | draw | Card advantage support that fits the large-creature plan.
- 1 Rishkar's Expertise | draw | Card advantage support that fits the large-creature plan.
- 1 Sylvan Library | draw | Card advantage support.
- 1 Dragon Tempest | removal | Dragon-themed answer for opposing problems.
- 1 Dragonlord Atarka | removal | Dragon-based answer for opposing problems.
- 1 Drakuseth, Maw of Flames | removal | Dragon-based answer for opposing problems.
- 1 Glorybringer | removal | Dragon-based answer for opposing problems.
- 1 Invasion of Tarkir // Defiant Thundermaw | removal | Dragon-themed answer for opposing problems.
- 1 Scourge of Valkas | removal | Dragon-based answer for opposing problems.
- 1 Terror of the Peaks | removal | Dragon-based answer for opposing problems.
- 1 Wrathful Red Dragon | removal | Dragon-based answer for opposing problems.
- 1 Balefire Dragon | wipe | Dragon-based board-reset option.
- 1 Draconic Intervention | wipe | Board-reset option that fits the dragon theme.
- 1 Incinerator of the Guilty | wipe | Dragon-based board-reset option.
- 1 Heroic Intervention | interaction | Flexible interaction for safeguarding the board plan.
- 1 Lightning Greaves | interaction | Flexible interaction for supporting a key creature.
- 1 Swiftfoot Boots | interaction | Flexible interaction for supporting a key creature.
- 1 The One Ring | interaction | Flexible interaction for the broader game plan.
- 1 Veil of Summer | interaction | Flexible interaction for safeguarding the plan.
- 1 Whispersilk Cloak | interaction | Flexible interaction for supporting a key creature.
- 1 Acolyte of Bahamut | synergy | Dragon-focused support piece.
- 1 Breaching Dragonstorm | synergy | Dragon-focused support piece.
- 1 Crucible of Fire | synergy | Dragon-focused support piece.
- 1 Dracogenesis | synergy | Dragon-focused support piece.
- 1 Dragon Egg | synergy | Dragon-themed support creature.
- 1 Dragon Hatchling | synergy | Dragon-themed support creature.
- 1 Dragonkin Berserker | synergy | Dragon-focused support piece.
- 1 Dragonlord's Servant | synergy | Dragon-focused support creature.
- 1 Dragonspeaker Shaman | synergy | Dragon-focused support creature.
- 1 Firespitter Whelp | synergy | Dragon-themed support creature.
- 1 Kargan Dragonrider | synergy | Dragon-focused support creature.
- 1 Minion of the Mighty | synergy | Dragon-focused support creature.
- 1 Sarkhan's Triumph | synergy | Dragon-focused support piece.
- 1 Shivan Devastator | synergy | Dragon-themed support creature.
- 1 Ancient Bronze Dragon | threat | High-impact Dragon threat.
- 1 Ambitious Dragonborn | threat | Dragon threat for the creature plan.
- 1 Backdraft Hellkite | threat | Dragon threat for the creature plan.
- 1 Blast-Furnace Hellkite | threat | Dragon threat for the creature plan.
- 1 Caldera Pyremaw | threat | Dragon threat for the creature plan.
- 1 Canopy Gargantuan | threat | Dragon threat for the creature plan.
- 1 Hellkite Charger | threat | Dragon threat for the creature plan.
- 1 Lathliss, Dragon Queen | threat | High-impact Dragon threat.
- 1 Scourge of the Throne | threat | Dragon threat for the creature plan.
- 1 Terror of Mount Velus | threat | Dragon threat for the creature plan.
- 1 Thrakkus the Butcher | threat | Dragon threat for the creature plan.
- 1 Utvara Hellkite | threat | High-impact Dragon threat.

</details>

### 15. delegated commander, owned first

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 225 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $0.00 to buy, $200.80 the whole deck.

**Summary:** This Cecil deck develops a lifegain-centered board, using supportive creatures and equipment to make its key creatures matter while Angels and other threats apply steady pressure. It wins by maintaining that creature presence through removal and protection, then turning a stable board into combat damage. The tradeoff is a focused, board-based plan that can be slower to rebuild after repeated disruption than a deck built around a single immediate finish.

- [INFO] `bracket_prose_rules`: bracket 2 (Core, 8+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 3.14 over 63 nonland cards

<details><summary>The deck list</summary>

- 12 Plains | land | Provides the deck’s white mana base.
- 12 Swamp | land | Provides the deck’s black mana base.
- 1 Command Tower | land | A flexible land for Cecil’s colors.
- 1 Exotic Orchard | land | Adds another flexible mana source.
- 1 Ash Barrens | land | Helps stabilize the basic-land mana base.
- 1 Evolving Wilds | land | Finds a needed basic land.
- 1 Fabled Passage | land | Finds a needed basic land.
- 1 Marsh Flats | land | Supports the Plains-and-Swamp mana base.
- 1 Path of Ancestry | land | Provides a colored land slot for the creature-focused deck.
- 1 Scavenger Grounds | land | Adds utility while occupying a land slot.
- 1 Takenuma, Abandoned Mire | land | Adds black mana and utility.
- 1 War Room | land | Provides a utility land slot.
- 1 Windbrisk Heights | land | Provides a white-producing utility land.
- 1 Sidequest: Catch a Fish // Cooking Campsite | land | Provides a land slot that fits the deck’s food-adjacent theme.
- 1 Sol Ring | ramp | A compact early mana accelerant.
- 1 Arcane Signet | ramp | Provides reliable color fixing and acceleration.
- 1 Bender's Waterskin | ramp | Adds an artifact-based mana source.
- 1 Commander's Sphere | ramp | Provides colored acceleration.
- 1 Fellwar Stone | ramp | Adds efficient artifact acceleration.
- 1 Chromatic Lantern | ramp | Helps smooth the mana base while accelerating.
- 1 Giada, Font of Hope | ramp | Supports the deck’s Angel contingent while accelerating it.
- 1 Lotho, Corrupt Shirriff | ramp | Adds a creature-based ramp piece.
- 1 Sword of the Animist | ramp | Pairs mana development with the equipment package.
- 1 Wayfarer's Bauble | ramp | Adds basic-land mana development.
- 1 Buster Sword | draw | Provides card flow through the equipment package.
- 1 Call of the Ring | draw | A dedicated draw piece for the deck.
- 1 Exemplar of Light | draw | Combines a creature body with card flow.
- 1 Idol of Oblivion | draw | An artifact draw option for the deck.
- 1 Inspiring Overseer | draw | Provides a flying creature alongside card flow.
- 1 Lembas | draw | A compact artifact draw piece.
- 1 Mask of Memory | draw | Supports combat-oriented card flow.
- 1 Night's Whisper | draw | A direct card-flow spell.
- 1 Puresteel Paladin | draw | Supports the equipment package while providing card flow.
- 1 Skullclamp | draw | An equipment-based draw option.
- 1 Banishing Light | removal | A flexible permanent answer.
- 1 Bitter Triumph | removal | A low-cost removal spell.
- 1 Crib Swap | removal | A creature-focused removal option.
- 1 Generous Gift | removal | Answers a wide range of opposing permanents.
- 1 Get Lost | removal | Provides a flexible removal spell.
- 1 Infernal Grasp | removal | A direct creature answer.
- 1 Swords to Plowshares | removal | An efficient creature answer.
- 1 Stroke of Midnight | removal | Adds another broad permanent answer.
- 1 Austere Command | wipe | Provides a flexible reset when the board gets out of hand.
- 1 Fumigate | wipe | A full-board reset that suits the lifegain theme.
- 1 Dusk // Dawn | wipe | Offers a board reset with later creature value.
- 1 Bastion Protector | interaction | Helps keep Cecil on the table.
- 1 Champion's Helm | interaction | Protects a key legendary creature through the equipment package.
- 1 Clever Concealment | interaction | Protects the developed board from disruption.
- 1 Darksteel Plate | interaction | A durable protection equipment for Cecil or a major threat.
- 1 Lightning Greaves | interaction | Provides efficient protection for important creatures.
- 1 Swiftfoot Boots | interaction | Adds another protection equipment piece.
- 1 Adventurous Eater // Have a Bite | synergy | Supports the deck’s lifegain-centered plan.
- 1 Aerith Gainsborough | synergy | A centerpiece creature for the lifegain theme.
- 1 Aettir and Priwen | synergy | Reinforces the deck’s equipment subtheme.
- 1 Angel of Vitality | synergy | A creature that fits the lifegain package.
- 1 Compassionate Healer | synergy | Directly supports the lifegain plan.
- 1 Crowd of True Believers | synergy | Adds a creature that supports the central theme.
- 1 Dancer's Chakrams | synergy | Builds on the deck’s equipment support.
- 1 Eastfarthing Farmer | synergy | Supports the deck’s food-adjacent lifegain theme.
- 1 Elixir | synergy | A compact artifact piece for the lifegain plan.
- 1 Light of Promise | synergy | Rewards the deck’s focus on gaining life.
- 1 Night Nurse, Healer of Heroes | synergy | A thematic creature for the lifegain package.
- 1 Prideful Feastling | synergy | Supports the food- and lifegain-adjacent portions of the deck.
- 1 Rosie Cotton of South Lane | synergy | A creature that complements the deck’s lifegain strategy.
- 1 Second Breakfast | synergy | Adds another spell aligned with the food-adjacent theme.
- 1 Angel of Invention | threat | A substantial creature threat that fits the white creature core.
- 1 Bill the Pony | threat | Adds another creature threat to pressure opponents.
- 1 Canyon Crawler | threat | A creature threat that helps fill out the board.
- 1 Dawnhand Eulogist | threat | Adds a creature threat to the deck’s finishing plan.
- 1 Foggy Swamp Hunters | threat | Provides another creature threat for sustained pressure.
- 1 Gollum, Silent Slinker // Meager Meal | threat | A threatening creature that also fits the deck’s food-adjacent texture.
- 1 Lyra Dawnbringer | threat | A powerful Angel threat for the deck’s creature plan.
- 1 Minwu, White Mage | threat | Adds a legendary creature threat to the board.
- 1 Rabaroo Troop | threat | Provides another creature-based avenue of pressure.
- 1 Shattered Angel | threat | A flying creature threat that suits the lifegain focus.
- 1 Sneering Shadewriter | threat | Adds a black creature threat to the deck’s curve.
- 1 Victory's Herald | threat | A top-end flying threat for closing games.

</details>

### 16. a tight budget, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 228 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $14.14 to buy, $40.76 the whole deck.

**Summary:** This is a black-white aristocrats deck that develops a creature-heavy board, uses sacrifice-focused synergy pieces to turn that board into pressure and card advantage, and closes by making creature losses matter to every opponent. It plays a broad spread of inexpensive creatures alongside removal, protection, and reset buttons, giving it a practical library-first build; in return, it is more dependent on establishing a board and can be slowed when its creatures are repeatedly answered.

- [INFO] `bracket_prose_rules`: bracket 2 (Core, 8+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [WARN] `not_owned`: Aron, Benalia's Ruin: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ayli, Eternal Pilgrim: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bartolomé del Presidio: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bastion of Remembrance: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Blood Artist: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Elas il-Kor, Sadistic Pilgrim: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Falkenrath Noble: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Skullport Merchant: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Smothering Abomination: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vindictive Vampire: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Woe Strider: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Yahenni, Undying Partisan: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Zulaport Cutthroat: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Embalmed Ascendant: the deck needs 1, the collection has 0
- [WARN] `not_owned`: High-Society Hunter: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Lord of the Forsaken: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Lord Skitter's Butcher: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Old Flitterfang: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ruthless Lawbringer: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Shilgengar, Sire of Famine: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Sivriss, Nightmare Speaker: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Umbral Collar Zealot: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vampire Gourmand: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Venerated Stormsinger: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vengeful Bloodwitch: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 2.79 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Ash Barrens | land | Flexible colorless land support for the mana base.
- 1 Castle Doom | land | A black-producing land for the deck's core color.
- 1 Command Tower | land | Reliable commander-focused color fixing.
- 1 Escape Tunnel | land | Utility land that supports mana consistency.
- 1 Evolving Wilds | land | Fixes basic-land access while filling a land slot.
- 1 Exotic Orchard | land | Owned multicolor fixing for the mana base.
- 1 Field of Ruin | land | Utility land slot with colorless mana access.
- 1 Ghost Quarter | land | Owned utility land for the mana base.
- 1 Ifnir Deadlands | land | Black mana source with useful utility.
- 1 Opal Palace | land | Commander-oriented utility land.
- 1 Path of Ancestry | land | Owned fixing land for the creature-heavy deck.
- 8 Plains | land | Basic white sources for dependable early mana.
- 1 Rogue's Passage | land | Utility land that helps a threat apply pressure.
- 1 Scavenger Grounds | land | Utility land that also provides colorless mana.
- 1 Secluded Courtyard | land | Creature-focused fixing for the deck's many bodies.
- 9 Swamp | land | Basic black sources for the deck's primary sacrifice shell.
- 1 Temple of the False God | land | Owned land that helps support larger turns.
- 1 Terramorphic Expanse | land | Basic-land fixing from the existing collection.
- 1 Thriving Moor | land | Black source that contributes to color fixing.
- 1 Unclaimed Territory | land | Creature-oriented fixing land.
- 1 Windbrisk Heights | land | White-producing utility land.
- 1 Arcane Signet | ramp | Efficient owned mana acceleration.
- 1 Astral Cornucopia | ramp | Owned artifact acceleration for longer games.
- 1 Chromatic Lantern | ramp | Owned fixing and ramp in one card.
- 1 Commander's Sphere | ramp | Reliable artifact mana source.
- 1 Deadly Dispute | ramp | Owned sacrifice-compatible resource acceleration.
- 1 Fellwar Stone | ramp | Low-cost owned mana acceleration.
- 1 Inherited Envelope | ramp | Owned artifact ramp for the deck's curve.
- 1 Sol Ring | ramp | Fast, owned artifact mana.
- 1 Thought Vessel | ramp | Owned mana rock that supports a full hand.
- 1 White Auracite | ramp | Low-cost owned artifact acceleration.
- 1 Ahriman | draw | Owned creature-based card advantage.
- 1 Beetle-Headed Merchants | draw | Low-cost owned draw support.
- 1 Circle of Power | draw | Owned card-advantage spell.
- 1 Grave Venerations | draw | Draw support that fits the deck's creature focus.
- 1 Idol of Oblivion | draw | Owned artifact source of repeatable card advantage.
- 1 Mask of Memory | draw | Equipment-based draw support for attacking creatures.
- 1 Nasty End | draw | Owned instant-speed card advantage.
- 1 Painful Truths | draw | Efficient owned draw spell.
- 1 Skullclamp | draw | Powerful owned card-advantage equipment for small creatures.
- 1 Tome of Legends | draw | Owned artifact draw that supports the commander.
- 1 Bitter Triumph | removal | Flexible owned spot removal.
- 1 Claim the Precious | removal | Reliable owned creature removal.
- 1 Crib Swap | removal | Owned instant-speed removal.
- 1 Deadly Precision | removal | Low-cost owned removal spell.
- 1 Destroy Evil | removal | Flexible owned answer to key permanents.
- 1 Fatal Push | removal | Efficient owned creature interaction.
- 1 Infernal Grasp | removal | Straightforward owned creature removal.
- 1 Swords to Plowshares | removal | Efficient owned exile-based answer.
- 1 Archfiend of Ifnir | wipe | Creature-based board-control option.
- 1 Austere Command | wipe | Flexible owned reset for troublesome boards.
- 1 Black Sun's Zenith | wipe | Scalable owned board wipe.
- 1 Bastion Protector | interaction | Owned protection for the commander.
- 1 Gift of Immortality | interaction | Owned protection that supports keeping a key creature available.
- 1 Swiftfoot Boots | interaction | Efficient owned protection equipment.
- 1 Take Up the Shield | interaction | Low-cost owned protection spell.
- 1 Together Forever | interaction | Protection-oriented enchantment for important creatures.
- 1 Ultimate Magic: Holy | interaction | Owned instant-speed defensive interaction.
- 1 Aron, Benalia's Ruin | synergy | Low-cost aristocrats-support creature.
- 1 Ayli, Eternal Pilgrim | synergy | Low-cost creature that supports the sacrifice-focused plan.
- 1 Bartolomé del Presidio | synergy | Low-cost sacrifice-shell creature.
- 1 Bastion of Remembrance | synergy | Dedicated aristocrats payoff for the deck's core plan.
- 1 Blood Artist | synergy | Central creature-death payoff for an aristocrats build.
- 1 Elas il-Kor, Sadistic Pilgrim | synergy | Creature-based payoff that supports the aristocrats shell.
- 1 Falkenrath Noble | synergy | Affordable creature-death payoff.
- 1 Gollum, Patient Plotter | synergy | Owned synergy creature for the sacrifice plan.
- 1 Skullport Merchant | synergy | Affordable sacrifice-focused value creature.
- 1 Smothering Abomination | synergy | Creature-based value piece for the aristocrats shell.
- 1 Vindictive Vampire | synergy | Affordable creature-death payoff.
- 1 Woe Strider | synergy | Affordable sacrifice-focused creature.
- 1 Yahenni, Undying Partisan | synergy | Low-cost sacrifice-shell creature.
- 1 Zulaport Cutthroat | synergy | Creature-death payoff that advances the main plan.
- 1 Embalmed Ascendant | threat | Affordable creature threat for building the board.
- 1 High-Society Hunter | threat | Low-cost creature that adds pressure.
- 1 Lord of the Forsaken | threat | Affordable Demon threat for the deck's top end.
- 1 Lord Skitter's Butcher | threat | Affordable black creature that adds board pressure.
- 1 Old Flitterfang | threat | Affordable creature threat that fits the black creature base.
- 1 Ruthless Lawbringer | threat | Low-cost white-black creature threat.
- 1 Shilgengar, Sire of Famine | threat | Affordable legendary threat for the creature plan.
- 1 Sivriss, Nightmare Speaker | threat | Cheap black creature threat.
- 1 Umbral Collar Zealot | threat | Affordable creature that builds board presence.
- 1 Vampire Gourmand | threat | Low-cost Vampire threat.
- 1 Venerated Stormsinger | threat | Affordable creature threat for the deck's board plan.
- 1 Vengeful Bloodwitch | threat | Affordable Vampire threat that adds pressure.

</details>

