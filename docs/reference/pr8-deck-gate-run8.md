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
| Decks that needed the repair turn | 2 |
| Summaries judged (F-26) | 18 |
| Summaries that state a rule of the game | 0 |
| Summaries that state a FALSE rule | 0 |
| Judge errors | 0 |
| Errors | 0 |
| Prompt version | 9 |
| Calls | 38 |
| Cost | $1.0898 |
| Time | 912 seconds |

## Findings by code

A block stops the deck. A warning and a note are reports.

| Code | Count |
|---|---|
| `curve_summary` | 18 |
| `not_owned` | 13 |
| `bracket_prose_rules` | 12 |
| `land_count` | 3 |
| `basics_added` | 1 |
| `precon_cards_restored` | 1 |

By severity: BLOCK 0. WARN 16. INFO 32. 

## Decks

### 1. lifegain Commander, any card

Format: Commander. Theme: lifegain. Pool: any_card. Shortlist: 303 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $403.38 to buy, $403.38 the whole deck.

**Summary:** This Karlov deck builds around a lifegain-focused creature and enchantment core, using ramp and draw to establish its board before applying pressure with Angels, Vampires, and other large threats. It keeps opposing plans in check with a solid spread of interaction, removal, and a few board resets, then looks to turn its established board into a decisive finish. The tradeoff is that the deck is committed to its permanent-based synergy plan, so it is less focused on standalone alternate-win cards.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 3.44 over 63 nonland cards

<details><summary>The deck list</summary>

- 18 Plains | land | Basic land slot for the white portion of the mana base.
- 10 Swamp | land | Basic land slot for the black portion of the mana base.
- 1 Kabira Crossroads | land | Land slot that supports the deck’s lifegain plan.
- 1 Scoured Barrens | land | Land slot for the two-color mana base.
- 1 Shambling Vent | land | Land slot for the two-color mana base.
- 1 Restless Fortress | land | Land slot for the two-color mana base.
- 1 Vault of the Archangel | land | Land slot that fits the deck’s colors.
- 1 Radiant Fountain | land | Land slot that supports the lifegain theme.
- 1 High Market | land | Utility land slot for the mana base.
- 1 Westvale Abbey // Ormendahl, Profane Prince | land | Land slot with a late-game option.
- 1 Angel of Indemnity | ramp | Ramp piece that also fits the deck’s creature theme.
- 1 Battle Angels of Tyr | ramp | Ramp piece that contributes to the deck’s board presence.
- 1 Beza, the Bounding Spring | ramp | Ramp piece for developing the deck’s resources.
- 1 Carmen, Cruel Skymarcher | ramp | Ramp piece in the deck’s colors.
- 1 Crypt Ghast | ramp | Ramp piece for accelerating the deck’s mana.
- 1 Hierophant's Chalice | ramp | Ramp artifact that supports the lifegain plan.
- 1 Nuka-Cola Vending Machine | ramp | Ramp artifact for resource development.
- 1 Orazca Relic | ramp | Ramp artifact for the mana base.
- 1 Pristine Talisman | ramp | Ramp artifact that fits the lifegain theme.
- 1 The Celestus | ramp | Ramp artifact for smoothing development.
- 1 Archivist of Oghma | draw | Draw creature that helps keep cards flowing.
- 1 Ayara, First of Locthwain | draw | Draw creature in the deck’s black core.
- 1 Convalescent Care | draw | Draw enchantment that fits the lifegain strategy.
- 1 Cosmos Elixir | draw | Draw artifact for sustained card access.
- 1 Dawn of Hope | draw | Draw enchantment aligned with the lifegain plan.
- 1 Enduring Innocence | draw | Draw creature that supports the deck’s board plan.
- 1 Mangara, the Diplomat | draw | Draw creature that provides steady card access.
- 1 Markov Purifier | draw | Draw creature that fits the Vampire subtheme.
- 1 Sigarda's Splendor | draw | Draw enchantment for the lifegain strategy.
- 1 Well of Lost Dreams | draw | Draw artifact that complements gaining life.
- 1 Alseid of Life's Bounty | interaction | Interaction piece for protecting the deck’s key permanents.
- 1 Faith's Shield | interaction | Interaction spell for safeguarding an important threat.
- 1 Metropolis Reformer | interaction | Interaction creature that supports the deck’s board.
- 1 Restoration Magic | interaction | Interaction spell for protecting the deck’s plan.
- 1 Sword of Light and Shadow | interaction | Interaction equipment that supports creature-based play.
- 1 Werefox Bodyguard | interaction | Interaction creature that contributes to board presence.
- 1 Aetherflux Reservoir | removal | Removal artifact that rewards the lifegain plan.
- 1 Ayli, Eternal Pilgrim | removal | Removal creature that fits the deck’s colors and theme.
- 1 Consuming Corruption | removal | Removal spell for answering opposing threats.
- 1 Murderous Rider // Swift End | removal | Removal option that also supplies a creature.
- 1 Solitude | removal | Removal creature for efficient answers.
- 1 Umezawa's Jitte | removal | Removal equipment that supports creature combat.
- 1 Vona, Butcher of Magan | removal | Removal creature that fits the Vampire theme.
- 1 Witch of the Moors | removal | Removal creature that supports the lifegain strategy.
- 1 Aerith Gainsborough | synergy | Synergy creature for the deck’s lifegain plan.
- 1 Ajani's Pridemate | synergy | Synergy creature that rewards the deck’s central theme.
- 1 Angel of Vitality | synergy | Synergy creature that supports gaining life.
- 1 Angelic Accord | synergy | Synergy enchantment for the lifegain game plan.
- 1 Blood Artist | synergy | Synergy creature that supports the deck’s attrition plan.
- 1 Cleric Class | synergy | Synergy enchantment for the lifegain strategy.
- 1 Cleric of Life's Bond | synergy | Synergy creature that fits the Cleric and lifegain themes.
- 1 Heliod, Sun-Crowned | synergy | Synergy permanent for building around lifegain.
- 1 Resplendent Angel | synergy | Synergy creature that supports the lifegain plan.
- 1 Righteous Valkyrie | synergy | Synergy creature for the Angel and lifegain themes.
- 1 Serra Ascendant | synergy | Synergy creature that rewards the deck’s life total plan.
- 1 Vito, Thorn of the Dusk Rose | synergy | Synergy creature that supports the lifegain strategy.
- 1 Vizkopa Guildmage | synergy | Synergy creature in the deck’s colors.
- 1 Voice of the Blessed | synergy | Synergy creature that rewards the deck’s central plan.
- 1 Archangel of Thune | threat | Threat that pairs well with the deck’s lifegain theme.
- 1 Astarion, the Decadent | threat | Threat that fits the deck’s black-white strategy.
- 1 Attended Healer | threat | Threat that supports the deck’s lifegain plan.
- 1 Blood Baron of Vizkopa | threat | Threat that fits the Vampire subtheme.
- 1 Bloodbond Vampire | threat | Threat that rewards the deck’s lifegain plan.
- 1 Celestine, the Living Saint | threat | Threat that supports the deck’s creature plan.
- 1 Cliffhaven Vampire | threat | Threat that fits the lifegain and Vampire themes.
- 1 Defiant Bloodlord | threat | Threat that supports the lifegain strategy.
- 1 Divinity of Pride | threat | Threat that fits the deck’s life-total plan.
- 1 Epicure of Blood | threat | Threat that supports the deck’s attrition plan.
- 1 Nykthos Paragon | threat | Threat that rewards the lifegain strategy.
- 1 Rhox Faithmender | threat | Threat that supports the deck’s lifegain theme.
- 1 Ajani, Strength of the Pride | wipe | Board wipe that fits the deck’s lifegain plan.
- 1 Fumigate | wipe | Board wipe for resetting difficult boards.
- 1 Kaya's Wrath | wipe | Board wipe in the deck’s colors.

</details>

### 2. aristocrats Commander, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 179 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $0.00 to buy, $149.86 the whole deck.

**Summary:** Denethor, Ruling Steward anchors an aristocrats-focused black-white shell that leans on its dedicated synergy pieces, creature-based value, and a deep supply of cards to keep the game moving. It aims to win by building lasting board advantage through the sacrifice-and-death plan, while plentiful removal and a few sweepers keep opposing boards from pulling too far ahead. The tradeoff is a slower, more support-heavy approach rather than a fast threat-driven finish.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 2.78 over 63 nonland cards

<details><summary>The deck list</summary>

- 12 Plains | land | Provides reliable white mana.
- 12 Swamp | land | Provides reliable black mana.
- 1 Ash Barrens | land | Flexible mana support for the two-color base.
- 1 Castle Locthwain | land | A black source with late-game utility.
- 1 Command Tower | land | Reliable commander mana.
- 1 Evolving Wilds | land | Fixes the deck's mana base.
- 1 Fabled Passage | land | Fixes the deck's mana base.
- 1 Marsh Flats | land | Finds the deck's basic mana sources.
- 1 Path of Ancestry | land | Color fixing for the creature-focused plan.
- 1 Plaza of Heroes | land | Utility land for the legendary commander shell.
- 1 Scavenger Grounds | land | Utility land that still supplies mana.
- 1 Takenuma, Abandoned Mire | land | Black source with useful late-game utility.
- 1 War Room | land | Utility land for longer games.
- 1 Windbrisk Heights | land | A white source with additional utility.
- 1 Arcane Signet | ramp | Efficient color fixing and mana acceleration.
- 1 Astral Cornucopia | ramp | Artifact-based mana acceleration.
- 1 Commander's Sphere | ramp | Fixes mana while retaining later utility.
- 1 Fellwar Stone | ramp | Low-cost mana acceleration.
- 1 Inherited Envelope | ramp | Artifact-based mana support.
- 1 Relic of Legends | ramp | Mana acceleration for the legendary shell.
- 1 Sol Ring | ramp | Fast colorless mana acceleration.
- 1 Thought Vessel | ramp | Mana acceleration for longer games.
- 1 Wayfarer's Bauble | ramp | Early mana development.
- 1 Sword of the Animist | ramp | Equipment-based mana development.
- 1 Ahriman | draw | Creature-based card advantage.
- 1 Call of the Ring | draw | Ongoing card advantage.
- 1 Cirith Ungol Patrol | draw | Creature-based card advantage.
- 1 Exemplar of Light | draw | Creature-based card advantage.
- 1 Idol of Oblivion | draw | Artifact card advantage for the long game.
- 1 Inspiring Overseer | draw | A creature that contributes card advantage.
- 1 Lembas | draw | Low-cost artifact card advantage.
- 1 Mask of Memory | draw | Equipment-based card advantage.
- 1 Massacre Girl, Known Killer | draw | Card advantage attached to a creature.
- 1 Night's Whisper | draw | Efficient black card advantage.
- 1 Puresteel Paladin | draw | Card advantage that also fits the equipment package.
- 1 Rat King, Pale Piper | draw | Creature-based card advantage.
- 1 Secret Rendezvous | draw | Additional card advantage for grindy games.
- 1 Skullclamp | draw | Efficient card advantage for the creature plan.
- 1 Stone of Erech | draw | Artifact card advantage.
- 1 The Sackville-Bagginses | draw | Creature-based card advantage.
- 1 Tome of Legends | draw | Reliable incremental card advantage.
- 1 Wall of Omens | draw | Early defense that replaces itself.
- 1 Bastion Protector | interaction | Protects the commander-centered plan.
- 1 Boromir, Warden of the Tower | interaction | Creature-based disruption for key moments.
- 1 Clever Concealment | interaction | Protects the board from opposing answers.
- 1 Darksteel Plate | interaction | Equipment protection for an important creature.
- 1 Frontline Medic | interaction | Creature-based protection and disruption.
- 1 Reprieve | interaction | Flexible disruption against a key spell.
- 1 Champion's Helm | interaction | Additional commander protection.
- 1 Gift of Immortality | interaction | Protection that supports recurring creature value.
- 1 Lightning Greaves | interaction | Efficient protection for key creatures.
- 1 Swiftfoot Boots | interaction | Backup protection for key creatures.
- 1 Together Forever | interaction | Protects the creature-focused game plan.
- 1 Angel of Serenity | removal | Creature-based answer to opposing permanents.
- 1 Banishing Light | removal | Versatile permanent answer.
- 1 Bitter Triumph | removal | Efficient black spot removal.
- 1 Blowfly Infestation | removal | Incremental removal support.
- 1 Claim the Precious | removal | Focused black removal.
- 1 Contagion Clasp | removal | Artifact-based removal support.
- 1 Crib Swap | removal | Flexible creature answer.
- 1 Destroy Evil | removal | Flexible answer to troublesome permanents.
- 1 Fiend Hunter | removal | Creature-based removal for the board plan.
- 1 Generous Gift | removal | Broad answer to opposing permanents.
- 1 Get Lost | removal | Efficient white removal.
- 1 Heartless Act | removal | Low-cost black creature removal.
- 1 Infernal Grasp | removal | Reliable black spot removal.
- 1 Palace Jailer | removal | Creature-based removal and board presence.
- 1 Swords to Plowshares | removal | Efficient creature removal.
- 1 Witch-king of Angmar | removal | A removal piece attached to a creature.
- 1 Austere Command | wipe | Flexible reset for difficult boards.
- 1 Dusk // Dawn | wipe | Board reset that suits the creature plan.
- 1 Fumigate | wipe | Reliable full-board reset.
- 1 Arcade Cabinet | synergy | Supports the deck's aristocrats core.
- 1 Gollum the Abandoned | synergy | Supports the sacrifice-and-death plan.
- 1 Gollum, Patient Plotter | synergy | Supports the sacrifice-and-death plan.
- 1 Joo Dee, One of Many | synergy | Adds another dedicated synergy piece.
- 1 Phantom Train | synergy | Artifact synergy for the deck's central plan.

</details>

### 3. artifacts Commander, bracket 4

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 299 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $1369.03 to buy, $1369.03 the whole deck.

**Summary:** This deck accelerates into a dense artifact board, keeps its hand supplied, and uses a compact package of interaction and broad answers to protect that position. It wins through its artifact threats and a developed board, with Urza at the center of the plan. It gives up broader card variety for a concentrated artifact shell and an Island-heavy mana base.

- [INFO] `curve_summary`: average mana value 3.98 over 63 nonland cards

<details><summary>The deck list</summary>

- 17 Island | land | Basic blue land for the mana base.
- 1 Academy Ruins | land | Specialty land slot for the artifact-focused mana base.
- 1 Archway of Innovation | land | Specialty land slot for the artifact-focused mana base.
- 1 Blinkmoth Nexus | land | Specialty land slot for the artifact-focused mana base.
- 1 Buried Ruin | land | Specialty land slot for the artifact-focused mana base.
- 1 Conqueror's Galleon // Conqueror's Foothold | land | Land slot that also keeps the deck’s artifact theme present.
- 1 Fomori Vault | land | Specialty land slot for the artifact-focused mana base.
- 1 Glimmervoid | land | Specialty land slot for the artifact-focused mana base.
- 1 Hall of Tagsin | land | Specialty land slot for the artifact-focused mana base.
- 1 Inventors' Fair | land | Specialty land slot for the artifact-focused mana base.
- 1 Mech Hangar | land | Specialty land slot for the artifact-focused mana base.
- 1 Mishra's Factory | land | Specialty land slot for the artifact-focused mana base.
- 1 Mishra's Foundry | land | Specialty land slot for the artifact-focused mana base.
- 1 Mishra's Workshop | land | Specialty land slot for the artifact-focused mana base.
- 1 Otawara, Soaring City | land | Blue specialty land for the mana base.
- 1 Power Depot | land | Artifact land slot that reinforces the deck’s theme.
- 1 Spire of Industry | land | Specialty land slot for the artifact-focused mana base.
- 1 The Mycosynth Gardens | land | Specialty land slot for the artifact-focused mana base.
- 1 Urza's Saga | land | Artifact-themed land slot for an Urza deck.
- 1 Urza's Workshop | land | Artifact-themed land slot for an Urza deck.
- 1 Thoughtcast | draw | Efficient draw selection for an artifact-focused deck.
- 1 Thought Monitor | draw | Artifact creature that fills a draw slot while supporting the theme.
- 1 Thirst for Knowledge | draw | Draw selection suited to the artifact-focused plan.
- 1 Reverse Engineer | draw | Draw selection for keeping the deck supplied.
- 1 Sai, Master Thopterist | draw | Artifact-focused legendary creature filling a draw slot.
- 1 Vedalken Archmage | draw | Artifact-focused creature draw option.
- 1 Riddlesmith | draw | Artificer draw option that fits the central theme.
- 1 Forensic Gadgeteer | draw | Artificer draw option for the deck’s focused plan.
- 1 Tezzeret, Artifice Master | draw | Artifact-themed planeswalker draw option.
- 1 Tezzeret, Betrayer of Flesh | draw | Artifact-themed planeswalker draw option.
- 1 Assert Authority | interaction | Interaction slot for protecting the deck’s plan.
- 1 Disruption Protocol | interaction | Interaction slot for protecting the deck’s plan.
- 1 Metallic Rebuke | interaction | Artifact-aligned interaction for the focused strategy.
- 1 Stoic Rebuttal | interaction | Interaction slot for protecting the deck’s plan.
- 1 Darksteel Forge | interaction | Artifact interaction piece for safeguarding the central theme.
- 1 Welding Jar | interaction | Low-profile artifact interaction for the focused strategy.
- 1 Mox Opal | ramp | Compact artifact ramp for an accelerated start.
- 1 Metalworker | ramp | Artifact creature ramp that supports rapid development.
- 1 Krark-Clan Ironworks | ramp | Artifact ramp piece for powering the deck’s plan.
- 1 Chief Engineer | ramp | Artificer ramp option for the artifact-heavy build.
- 1 Grand Architect | ramp | Artifact-focused creature ramp option.
- 1 Tezzeret the Seeker | ramp | Artifact-themed planeswalker ramp option.
- 1 Moonsnare Prototype | ramp | Compact artifact ramp for the deck’s acceleration package.
- 1 Inspiring Statuary | ramp | Artifact ramp piece for developing the board quickly.
- 1 Karn, Legacy Reforged | ramp | Artifact-themed legendary ramp option.
- 1 The Mightstone and Weakstone | ramp | Artifact ramp slot that stays within the deck’s theme.
- 1 Aether Spellbomb | removal | Artifact removal that fits the central theme.
- 1 Arcum Dagsson | removal | Artifact-focused legendary creature removal option.
- 1 Resculpt | removal | Flexible removal slot for handling opposing pieces.
- 1 Ravenform | removal | Removal slot for answering opposing pieces.
- 1 Portal to Phyrexia | removal | Artifact removal piece with a high-impact profile.
- 1 Spine of Ish Sah | removal | Artifact removal piece that reinforces the theme.
- 1 Skysovereign, Consul Flagship | removal | Artifact Vehicle removal option for the deck.
- 1 Transmogrifying Wand | removal | Artifact removal slot for the focused strategy.
- 1 Engineered Explosives | wipe | Artifact board-reset option for difficult boards.
- 1 Nevinyrral's Disk | wipe | Artifact board-reset option for difficult boards.
- 1 Oblivion Stone | wipe | Artifact board-reset option for difficult boards.
- 1 Emry, Lurker of the Loch | synergy | Artifact-focused synergy creature for the deck’s core plan.
- 1 Etherium Sculptor | synergy | Artifact creature synergy piece for the focused build.
- 1 Foundry Inspector | synergy | Artifact creature synergy piece for the focused build.
- 1 Mystic Forge | synergy | Central artifact synergy piece for the deck’s plan.
- 1 Unwinding Clock | synergy | Artifact synergy piece for maintaining momentum.
- 1 Voltaic Key | synergy | Compact artifact synergy piece for the core plan.
- 1 Clock of Omens | synergy | Artifact synergy piece for the focused strategy.
- 1 Manifold Key | synergy | Compact artifact synergy piece for the core plan.
- 1 Whir of Invention | synergy | Artifact-focused synergy spell for the deck’s plan.
- 1 Transmute Artifact | synergy | Artifact-focused synergy spell for the deck’s plan.
- 1 Reshape | synergy | Artifact-focused synergy spell for the deck’s plan.
- 1 Scrap Trawler | synergy | Artifact creature synergy piece for the focused build.
- 1 Shimmer Myr | synergy | Artifact creature synergy piece for the focused build.
- 1 Liberator, Urza's Battlethopter | synergy | Urza-linked artifact creature synergy piece.
- 1 Kuldotha Forgemaster | threat | Artifact creature threat for the deck’s finishing pressure.
- 1 Cyberdrive Awakener | threat | Artifact creature threat for the deck’s finishing pressure.
- 1 Kappa Cannoneer | threat | Artifact creature threat for the deck’s finishing pressure.
- 1 Master Transmuter | threat | Artifact creature threat that supports the central theme.
- 1 Metalwork Colossus | threat | Large artifact creature threat for closing games.
- 1 Mycosynth Golem | threat | Artifact creature threat for the deck’s finishing pressure.
- 1 Broodstar | threat | Artifact-focused creature threat for closing games.
- 1 Darksteel Juggernaut | threat | Artifact creature threat for the deck’s finishing pressure.
- 1 Karn, Scion of Urza | threat | Artifact-themed planeswalker threat for the core plan.
- 1 Tezzeret, Cruel Captain | threat | Artifact-themed planeswalker threat for the core plan.
- 1 Threefold Thunderhulk | threat | Artifact creature threat for the deck’s finishing pressure.
- 1 Traxos, Scourge of Kroog | threat | Artifact creature threat for the deck’s finishing pressure.

</details>

### 4. dinosaur tribal, bracket 2

Format: Commander. Theme: dinosaurs. Pool: any_card. Shortlist: 269 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $249.86 to buy, $249.86 the whole deck.

**Summary:** This deck builds its mana, develops a board of Dinosaurs, and uses Gishath alongside large tribal threats to win through combat. It has card draw to keep the creatures coming, Dinosaur-themed answers for opposing problems, and a few broad reset buttons when the table gets ahead. It gives up some speed and stack-based flexibility in exchange for a clear, creature-focused game plan and a high number of expensive threats.

- [INFO] `bracket_prose_rules`: bracket 2 (Core, 8+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 4.13 over 63 nonland cards

<details><summary>The deck list</summary>

- 11 Forest | land | Basic green mana supports the deck's Dinosaur-heavy core.
- 8 Mountain | land | Basic red mana supports the deck's red Dinosaurs and spells.
- 8 Plains | land | Basic white mana supports the deck's white Dinosaurs and spells.
- 1 Command Tower | land | Reliable three-color fixing for the commander deck.
- 1 Path of Ancestry | land | A tribal land that supports the Dinosaur theme.
- 1 Canopy Vista | land | Green-white mana fixing for the deck's color requirements.
- 1 Cinder Glade | land | Red-green mana fixing for the deck's color requirements.
- 1 Sacred Foundry | land | Red-white mana fixing for the deck's color requirements.
- 1 Temple Garden | land | Green-white mana fixing for the deck's color requirements.
- 1 Stomping Ground | land | Red-green mana fixing for the deck's color requirements.
- 1 Sunpetal Grove | land | Additional green-white mana fixing.
- 1 Rootbound Crag | land | Additional red-green mana fixing.
- 1 Sol Ring | ramp | Fast mana helps bring out the deck's larger creatures sooner.
- 1 Arcane Signet | ramp | Simple color fixing and acceleration.
- 1 Cultivate | ramp | Provides dependable land-based mana development.
- 1 Farseek | ramp | Finds mana fixing early in the game.
- 1 Kodama's Reach | ramp | Builds the land base while ensuring another land drop.
- 1 Nature's Lore | ramp | Efficiently improves green-based mana access.
- 1 Thunderherd Migration | ramp | Dinosaur-themed mana development.
- 1 Drover of the Mighty | ramp | A tribal-support creature that helps produce mana.
- 1 Hulking Raptor | ramp | A Dinosaur that contributes to the deck's mana plan.
- 1 Ranging Raptors | ramp | A Dinosaur that supports continued land development.
- 1 Beast Whisperer | draw | Keeps cards flowing alongside the creature-focused plan.
- 1 Garruk's Uprising | draw | Card advantage that fits a deck full of large creatures.
- 1 Guardian Project | draw | Long-game card advantage for a creature-heavy deck.
- 1 Harmonize | draw | Straightforward card draw for a new player.
- 1 Return of the Wildspeaker | draw | Flexible card draw that fits the deck's large threats.
- 1 Ripjaw Raptor | draw | A Dinosaur that supplies card advantage.
- 1 Rishkar's Expertise | draw | A powerful draw spell for a board of large creatures.
- 1 Shamanic Revelation | draw | Refills the hand from a developed creature board.
- 1 Toski, Bearer of Secrets | draw | Supports ongoing card advantage during combat.
- 1 Vanquisher's Banner | draw | Tribal card advantage that rewards the Dinosaur plan.
- 1 Akroma's Will | interaction | A flexible protection spell for important attacks and creatures.
- 1 Ephemerate | interaction | A low-cost trick that protects a key creature.
- 1 Heroic Intervention | interaction | Protects the developed Dinosaur board from disruption.
- 1 Lightning Greaves | interaction | Protects an important creature while keeping the plan moving.
- 1 Swiftfoot Boots | interaction | Provides repeatable protection for key creatures.
- 1 Temple Altisaur | interaction | A Dinosaur that helps the deck withstand opposing pressure.
- 1 Apex Altisaur | removal | A Dinosaur-based answer to opposing creatures.
- 1 Bronzebeak Foragers | removal | A Dinosaur that provides removal utility.
- 1 Burning Sun's Avatar | removal | A large Dinosaur that also contributes removal.
- 1 Itzquinth, Firstborn of Gishath | removal | A low-cost Dinosaur with removal utility.
- 1 Kogla and Yidaro | removal | A Dinosaur-themed threat that also fills a removal role.
- 1 Ravenous Sailback | removal | A Dinosaur-based answer to troublesome permanents.
- 1 Thrashing Brontodon | removal | A useful Dinosaur with removal utility.
- 1 Tranquil Frillback | removal | Flexible Dinosaur removal for multiple situations.
- 1 Austere Command | wipe | A flexible reset option when the board gets out of hand.
- 1 Blasphemous Act | wipe | A simple broad answer to a crowded battlefield.
- 1 Wakening Sun's Avatar | wipe | A Dinosaur-themed board reset attached to a large creature.
- 1 Amped Raptor | synergy | An on-theme Dinosaur that supports the tribal plan.
- 1 Commune with Dinosaurs | synergy | Helps find Dinosaur-themed resources early.
- 1 Dinosaur Stampede | synergy | A tribal combat spell for the Dinosaur board.
- 1 Gert and Old Lace, Runaways | synergy | An on-theme legendary Dinosaur support card.
- 1 Hunting Velociraptor | synergy | A Dinosaur that supports the deck's tribal game plan.
- 1 Huatli's Raptor | synergy | An inexpensive Dinosaur that reinforces the theme.
- 1 Invasion of Ixalan // Belligerent Regisaur | synergy | A Dinosaur-themed support card that remains on plan.
- 1 Kinjalli's Caller | synergy | Tribal support for deploying the deck's Dinosaur creatures.
- 1 Kinjalli's Sunwing | synergy | A Dinosaur that supports the deck's combat-focused plan.
- 1 Marauding Raptor | synergy | A tribal Dinosaur support piece.
- 1 Otepec Huntmaster | synergy | Supports the deck's Dinosaur-focused creature plan.
- 1 Orazca Frillback | synergy | An on-theme Dinosaur that supports the tribal strategy.
- 1 Raptor Companion | synergy | An efficient Dinosaur body for tribal support.
- 1 Raptor Hatchling | synergy | An early Dinosaur that stays aligned with the theme.
- 1 Ancient Brontodon | threat | A large Dinosaur that provides a straightforward finishing threat.
- 1 Carnage Tyrant | threat | A resilient-looking Dinosaur threat for closing games through combat.
- 1 Etali, Primal Storm | threat | A marquee Elder Dinosaur threat for the top end.
- 1 Ghalta, Primal Hunger | threat | A huge Dinosaur that rewards building a creature board.
- 1 Ghalta, Stampede Tyrant | threat | A major Elder Dinosaur threat for the late game.
- 1 Goring Ceratops | threat | A combat-focused Dinosaur threat.
- 1 Pantlaza, Sun-Favored | threat | A legendary Dinosaur threat that keeps the deck on theme.
- 1 Quartzwood Crasher | threat | A Dinosaur threat suited to a combat-heavy strategy.
- 1 Regisaur Alpha | threat | A powerful Dinosaur threat that reinforces the tribal board.
- 1 Thundering Spineback | threat | A Dinosaur threat for building a substantial battlefield.
- 1 Tyrranax Rex | threat | A large Dinosaur threat that pressures opponents.
- 1 Zetalpa, Primal Dawn | threat | A high-impact Elder Dinosaur for the deck's late game.

</details>

### 5. blink Commander, owned first

Format: Commander. Theme: blink. Pool: owned_first. Shortlist: 168 names.

Cards: 99 main, 0 sideboard. Repair turn: true. Block findings: 0.

Cost: $0.00 to buy, $164.41 the whole deck.

**Summary:** A mono-white blink deck built around repeatedly revisiting creature arrival abilities for cards, removal, and board presence. Protective equipment and instant-speed safeguards keep the centerpiece creatures available, while Angels, legendary creatures, and artifact threats provide a strong closing battlefield.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 3.06 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Abandoned Air Temple | land | Flexible white mana source.
- 1 Adventurer's Inn | land | White-producing utility land.
- 1 Ash Barrens | land | Fixes for a basic Plains when needed.
- 1 Command Tower | land | Reliable commander mana.
- 1 Evolving Wilds | land | Finds Plains and supports consistent mana.
- 1 Exotic Orchard | land | Reliable fixing in multiplayer games.
- 1 Fabled Passage | land | Fetches a Plains.
- 1 Great Hall of the Citadel | land | White-producing utility land.
- 1 Minas Tirith | land | Utility land that supports the Human-heavy creature base.
- 1 Path of Ancestry | land | Creature-focused fixing.
- 1 Plaza of Heroes | land | Protects and casts legendary creatures.
- 1 Secluded Courtyard | land | Names a prevalent creature type for fixing.
- 1 Terramorphic Expanse | land | Additional Plains-fetching land.
- 23 Plains | land | Stable base for a mono-white mana base.
- 1 Sol Ring | ramp | Efficient early acceleration.
- 1 Arcane Signet | ramp | Reliable mana acceleration.
- 1 Fellwar Stone | ramp | Low-cost multiplayer mana rock.
- 1 Wayfarer's Bauble | ramp | Finds a Plains and permanently advances mana.
- 1 Sword of the Animist | ramp | Turns attacks into repeatable land ramp.
- 1 Thought Vessel | ramp | Mana rock with hand-size utility.
- 1 Commander's Sphere | ramp | Mana fixing that can later cycle.
- 1 Bender's Waterskin | ramp | Artifact acceleration.
- 1 White Auracite | ramp | Additional mana acceleration.
- 1 Relic of Legends | ramp | Uses legendary creatures to produce mana.
- 1 Adventurer's Airship | draw | Repeatable card advantage source.
- 1 Buster Sword | draw | Equipment-based card advantage.
- 1 Crown of Gondor | draw | Supports the creature plan while providing cards.
- 1 Diary of Dreams | draw | Dedicated artifact draw engine.
- 1 Idol of Oblivion | draw | Efficient ongoing card advantage.
- 1 Instant Ramen | draw | Low-investment card draw.
- 1 Lembas | draw | Immediate card selection and later card draw.
- 1 Mask of Memory | draw | Combat-based filtering and card advantage.
- 1 Mirror of Galadriel | draw | Repeatable card selection.
- 1 Tome of Legends | draw | Steady card advantage alongside the commander.
- 1 Swords to Plowshares | removal | Premium answer to a creature.
- 1 Generous Gift | removal | Answers any problematic permanent.
- 1 Get Lost | removal | Efficient answer to creatures and nonland permanents.
- 1 Stroke of Midnight | removal | Versatile instant-speed permanent removal.
- 1 Destroy Evil | removal | Flexible removal for major threats.
- 1 Crib Swap | removal | Exiles a creature and works well with tribal creature themes.
- 1 Banishing Light | removal | Broad permanent-based answer.
- 1 Journey to Nowhere | removal | Efficient creature exile.
- 1 Lightning Greaves | interaction | Protects the commander and key creatures.
- 1 Swiftfoot Boots | interaction | Protection with flexible equipping.
- 1 Clever Concealment | interaction | Shields the board from opposing removal.
- 1 Reprieve | interaction | Temporarily disrupts a key spell while replacing itself.
- 1 Slip On the Ring | interaction | Protects a creature while retriggering its arrival ability.
- 1 Gift of Immortality | interaction | Keeps an important creature available through removal.
- 1 Flickerwisp | synergy | Core blink effect that reuses creature arrival abilities.
- 1 Personify | synergy | Supports the blink-focused game plan.
- 1 Angel of Condemnation | synergy | Repeatable creature blink attached to a substantial body.
- 1 Angel of Sanctions | synergy | Exile effect that becomes especially valuable when blinked.
- 1 Angel of Serenity | synergy | High-impact arrival ability for repeated blinking.
- 1 Fiend Hunter | synergy | Creature-based exile effect to reuse with blink.
- 1 Palace Jailer | synergy | Blinkable removal that adds monarch card advantage.
- 1 Inspiring Overseer | synergy | Arrival-trigger card advantage and life gain.
- 1 Wall of Omens | synergy | Efficient blinkable cantrip creature.
- 1 Exemplar of Light | synergy | Creature-based value that rewards repeated entries.
- 1 South Pole Voyager | synergy | Arrival value creature for the blink shell.
- 1 Joined Researchers // Secret Rendezvous | synergy | Provides a creature-based value piece to revisit.
- 1 Stiltzkin, Moogle Merchant | synergy | Blinkable value creature.
- 1 Weapons Vendor | synergy | Creature-based value that supports recurring arrival triggers.
- 1 Giada, Font of Hope | threat | Builds a growing Angel board.
- 1 Bronze Guardian | threat | Large artifact payoff that also protects key permanents.
- 1 Grim Poppet | threat | Scalable artifact creature that pressures creature boards.
- 1 The Vision | threat | Artifact creature threat with value potential.
- 1 The Walls of Ba Sing Se | threat | Resilient artifact creature that reinforces the board.
- 1 Zack Fair | threat | Legendary creature threat that contributes to combat pressure.
- 1 Bastion Protector | threat | Protective creature that strengthens the commander on board.
- 1 Frontline Medic | threat | Creature threat with combat utility.
- 1 Puresteel Paladin | threat | Equipment payoff creature that can generate substantial value.
- 1 Boromir, Warden of the Tower | threat | Legendary creature that disrupts opponents while applying pressure.
- 1 Vanquisher's Banner | threat | Creature-type payoff that turns a developed board into a win condition.
- 1 Crystal Fragments // Summon: Alexander | threat | Flexible artifact threat with a powerful back face.
- 1 Austere Command | wipe | Customizable sweeper that can preserve the preferred board.
- 1 Vanquish the Horde | wipe | Efficient creature-board reset.
- 1 Dusk // Dawn | wipe | Sweeps larger creatures while offering late-game creature recovery.

</details>

### 6. Modern tempo, tournament

Format: Modern. Theme: tempo. Pool: any_card. Shortlist: 167 names.

Cards: 60 main, 15 sideboard. Repair turn: false. Block findings: 0.

Cost: $508.46 to buy, $508.46 the whole deck.

**Summary:** This blue-red tempo deck starts with compact creature pressure from Ragavan, Nimble Pilferer, Ledger Shredder, and Faerie Mastermind, then uses draw, removal, and interaction to keep opponents from stabilizing. It wins by maintaining that pressure through key turns and using Temporal Mastery or Temporal Trespass to press an advantage. In exchange, it favors speed and efficient exchanges over broad answers to every kind of permanent.

- [INFO] `curve_summary`: average mana value 2.28 over 36 nonland cards

<details><summary>The deck list</summary>

- 4 Ragavan, Nimble Pilferer | threat | A low-cost creature threat for the blue-red tempo plan.
- 4 Ledger Shredder | threat | A creature threat that keeps the deck’s pressure focused on the air.
- 4 Faerie Mastermind | threat | A compact creature threat that supports a reactive tempo posture.
- 4 Consider | draw | Efficient card selection for finding pressure, answers, and lands.
- 2 Preordain | draw | Additional early card selection to smooth the game plan.
- 4 Counterspell | interaction | Broad reactive coverage to protect the tempo plan.
- 2 Spell Pierce | interaction | Cheap interaction for winning key early exchanges.
- 4 Lightning Bolt | removal | Efficient removal that clears the way for creature pressure.
- 2 Into the Flood Maw | removal | Flexible removal for preserving momentum.
- 2 Untimely Malfunction | removal | Extra removal that fits the deck’s reactive game plan.
- 2 Temporal Mastery | synergy | A temporal payoff that supports the deck’s pressure turns.
- 2 Temporal Trespass | synergy | A second temporal payoff for converting pressure into a finish.
- 4 Steam Vents | land | Primary blue-red source.
- 4 Scalding Tarn | land | Blue-red mana fixing land.
- 4 Shivan Reef | land | Fast blue-red source.
- 4 Stormcarved Coast | land | Additional blue-red source.
- 2 Sulfur Falls | land | Reliable dual land for the two-color base.
- 4 Island | land | Basic blue source.
- 2 Mountain | land | Basic red source.

</details>

### 7. Modern burn, casual

Format: Modern. Theme: burn. Pool: any_card. Shortlist: 290 names.

Cards: 60 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $41.04 to buy, $41.04 the whole deck.

**Summary:** This mono-red burn deck uses its threats and synergy pieces to keep pressure on the opponent, with draw helping it maintain resources and removal helping it stay on the front foot. It aims to finish games through sustained damage from its focused proactive plan, giving up broad flexibility and specialized answers in exchange for that direct approach.

- [INFO] `curve_summary`: average mana value 2.11 over 36 nonland cards

<details><summary>The deck list</summary>

- 24 Mountain | land | Fills the basic land base for the mono-red deck.
- 4 Risk Factor | draw | Provides draw support for the burn plan.
- 2 Voldaren Epicure | draw | Adds draw support while contributing to the deck's proactive plan.
- 2 Chandra, Dressed to Kill | ramp | Supplies the requested ramp support.
- 4 Lightning Bolt | removal | Provides efficient removal for opposing threats.
- 2 Lightning Strike | removal | Adds further removal to support the damage-focused plan.
- 4 Eidolon of the Great Revel | synergy | Supports the deck's burn-focused synergy package.
- 4 Kessig Flamebreather | synergy | Supports the deck's spell-heavy burn synergy.
- 4 Heartfire Hero | threat | Provides an aggressive threat for the deck.
- 4 Sawblade Scamp | threat | Adds more threats to maintain pressure.
- 4 Akki Lavarunner // Tok-Tok, Volcano Born | threat | Provides another threat for the proactive game plan.
- 2 Hazoret the Fervent | threat | Rounds out the threat package with a powerful finisher.

</details>

### 8. Modern lifegain, FNM

Format: Modern. Theme: lifegain. Pool: any_card. Shortlist: 301 names.

Cards: 60 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $277.04 to buy, $277.04 the whole deck.

**Summary:** This white-black deck develops a lifegain-focused board, turns that theme into pressure with its creature threats, and uses draw and removal to keep the plan moving. It wins by building enough momentum from its synergy pieces and finishing with its larger threats. The tradeoff is that the deck is geared toward its main-board lifegain plan rather than a broad range of specialized answers, and it has no dedicated sideboard.

- [INFO] `curve_summary`: average mana value 3.61 over 36 nonland cards

<details><summary>The deck list</summary>

- 8 Plains | land | Provides a basic white land base for the deck.
- 8 Swamp | land | Provides a basic black land base for the deck.
- 4 Scoured Barrens | land | Fills out the white-black land package.
- 4 Shambling Vent | land | Rounds out the deck's land package.
- 2 Altar of the Pantheon | ramp | Provides the deck's dedicated ramp slot.
- 4 Inspiring Overseer | draw | Supplies repeatable access to the deck's draw package.
- 2 Exemplar of Light | draw | Adds more draw support while staying in the creature plan.
- 4 Solitude | removal | Provides efficient removal alongside the creature-heavy plan.
- 2 Murderous Rider // Swift End | removal | Adds flexible removal to protect the lifegain board.
- 4 Soul Warden | synergy | A core lifegain synergy piece for the early game.
- 4 Ajani's Pridemate | synergy | A central payoff for the deck's lifegain synergies.
- 4 Attended Healer | threat | A lifegain-focused threat that supports the deck's main plan.
- 4 Bloodbond Vampire | threat | A threat that fits naturally beside the lifegain package.
- 2 Angel of Invention | threat | Provides a higher-impact creature threat for the middle of the curve.
- 2 Archangel of Thune | threat | Serves as a powerful top-end threat in the lifegain shell.
- 2 Defiant Bloodlord | threat | Adds additional finishing pressure at the top of the curve.

</details>

### 9. Standard midrange, FNM

Format: Standard. Theme: midrange. Pool: any_card. Shortlist: 164 names.

Cards: 60 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $176.66 to buy, $176.66 the whole deck.

**Summary:** This black-green midrange deck uses a creature-focused threat suite alongside removal and draw to keep pressure on the table. It aims to win by maintaining that pressure through the game while its support cards reinforce the main plan. In return, it favors a proactive, focused main deck over a dedicated sideboard or a sweep-heavy package.

- [WARN] `land_count`: 28 lands: the guide range for this format is 20 to 27
- [INFO] `curve_summary`: average mana value 2.88 over 32 nonland cards

<details><summary>The deck list</summary>

- 10 Forest | land | Basic green land for the black-green mana base.
- 10 Swamp | land | Basic black land for the black-green mana base.
- 2 Blooming Marsh | land | Black-green land support for the mana base.
- 2 Deathcap Glade | land | Black-green land support for the mana base.
- 2 Llanowar Elves | ramp | Early ramp for advancing the midrange plan.
- 4 Bitter Triumph | removal | Focused removal for clearing opposing threats.
- 2 Nowhere to Run | removal | Additional removal that supports the main plan.
- 2 Phyrexian Arena | draw | Reliable draw support for sustained midrange play.
- 2 Unholy Annex // Ritual Chamber | draw | Draw support that helps maintain resources.
- 2 Darkstar Augur | draw | Creature-based draw that fits the creature-focused plan.
- 4 Snakeskin Veil | synergy | Support card for protecting the deck's threat-focused plan.
- 4 Undying Malice | synergy | Support card that reinforces the creature-heavy strategy.
- 4 Goldvein Hydra | threat | A creature threat that gives the deck proactive pressure.
- 4 Vein Ripper | threat | A creature threat for the deck's top end.
- 2 Aclazotz, Deepest Betrayal // Temple of the Dead | threat | A powerful double-faced threat that also fits the mana plan.
- 2 Ojer Kaslem, Deepest Growth // Temple of Cultivation | threat | A double-faced creature threat for closing games.
- 2 Vaultborn Tyrant | threat | A large creature threat for the late game.

</details>

### 10. Standard aggro, tournament

Format: Standard. Theme: aggro. Pool: any_card. Shortlist: 294 names.

Cards: 64 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $44.50 to buy, $44.50 the whole deck.

**Summary:** This deck presses the table early with a steady stream of creature threats, using removal to keep attackers moving and interaction to preserve its momentum. It wins by maintaining pressure before an opponent can stabilize, with Warleader's Call reinforcing the red-white aggressive theme. In exchange for that speed, it gives up much of the staying power and broad answers of a slower strategy.

- [WARN] `land_count`: 28 lands: the guide range for this format is 20 to 27
- [INFO] `curve_summary`: average mana value 3.17 over 36 nonland cards

<details><summary>The deck list</summary>

- 4 Fugitive Codebreaker | draw | Provides draw while fitting the aggressive red plan.
- 2 Reckless Lackey | draw | Adds draw support for an attacking deck.
- 4 Sheltered by Ghosts | interaction | Provides interaction to protect the deck's aggressive plan.
- 2 Parting Gust | interaction | Adds flexible interaction for opposing plays.
- 4 Aang, the Last Airbender | removal | Supplies removal while remaining aligned with the red-white shell.
- 4 Aurelia's Vindicator | removal | Provides additional removal for clearing the way to attack.
- 4 Warleader's Call | synergy | Supports the deck's red-white aggro synergy.
- 4 Dragonback Lancer | threat | Serves as an aggressive creature threat.
- 4 Redcap Gutter-Dweller | threat | Adds pressure as a creature threat.
- 4 Teapot Slinger | threat | Rounds out the attacking creature suite.
- 12 Mountain | land | Provides the red mana base for the aggressive shell.
- 8 Plains | land | Provides the white mana base for the aggressive shell.
- 4 Sacred Foundry | land | Supports both colors of the red-white mana base.
- 4 Inspiring Vantage | land | Adds red-white mana support for an aggressive deck.

</details>

### 11. Commander with a locked card

Format: Commander. Theme: sacrifice. Pool: any_card. Shortlist: 296 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $408.08 to buy, $408.08 the whole deck.

**Summary:** This Karlov deck plays a black-white sacrifice game, building around creatures, artifacts, and enchantments that share that focus while using mana acceleration and card draw to keep moving. It aims to win by turning its sacrifice-focused threats and synergy pieces into a decisive board position, with removal and board wipes keeping opposing plans from stabilizing. It gives up some flexibility for a concentrated, black-heavy theme and relies on its sacrifice pieces working together.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 3.22 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Baron Bertram Graywater | draw | Included as a draw option for the sacrifice plan.
- 1 Corrupted Conviction | draw | Included as a draw option for the sacrifice plan.
- 1 Disciple of Bolas | draw | Included as a draw option for the sacrifice plan.
- 1 Ecstatic Awakener // Awoken Demon | draw | Included as a draw option that fits the deck's theme.
- 1 Relic Vial | draw | Included as a draw option for the sacrifice plan.
- 1 Shadowheart, Dark Justiciar | draw | Included as a draw option for the sacrifice plan.
- 1 Smothering Abomination | draw | Included as a draw option that fits the deck's theme.
- 1 Skyclave Shadowcat | draw | Included as a draw option for the sacrifice plan.
- 1 Vampiric Rites | draw | Included as a draw option for the sacrifice plan.
- 1 Village Rites | draw | Included as a draw option for the sacrifice plan.
- 1 Cartel Aristocrat | interaction | A sacrifice-themed interaction piece.
- 1 Dark Privilege | interaction | An interaction piece that fits the sacrifice plan.
- 1 Fanatical Devotion | interaction | An interaction piece that fits the sacrifice plan.
- 1 Flare of Fortitude | interaction | A white interaction option for protecting the deck's position.
- 1 Gift of Doom | interaction | An interaction piece that fits the sacrifice plan.
- 1 Nightmare Shepherd | interaction | An interaction piece that fits the sacrifice theme.
- 20 Swamp | land | Core black mana for the deck.
- 12 Plains | land | Core white mana for the deck.
- 1 Bojuka Bog | land | A black land slot with added utility.
- 1 Command Tower | land | A reliable land for Karlov's colors.
- 1 Evolving Wilds | land | A mana-fixing land slot.
- 1 Exotic Orchard | land | A flexible mana-fixing land slot.
- 1 Sol Ring | ramp | The requested fast mana piece.
- 1 Ashnod's Altar | ramp | A ramp piece tailored to a sacrifice deck.
- 1 Phyrexian Altar | ramp | A ramp piece tailored to a sacrifice deck.
- 1 Pitiless Plunderer | ramp | A ramp piece that supports the sacrifice plan.
- 1 Pawn of Ulamog | ramp | A ramp piece that supports the sacrifice plan.
- 1 Priest of Forgotten Gods | ramp | A sacrifice-themed ramp option.
- 1 Culling the Weak | ramp | A ramp option for the sacrifice plan.
- 1 Deadly Dispute | ramp | A ramp option that fits the sacrifice theme.
- 1 Warren Soultrader | ramp | A ramp piece that supports the sacrifice plan.
- 1 Skullport Merchant | ramp | A ramp piece that supports the sacrifice plan.
- 1 Attrition | removal | A sacrifice-themed removal piece.
- 1 Ayli, Eternal Pilgrim | removal | A removal option that fits Karlov's colors and plan.
- 1 Bone Shards | removal | An efficient removal option for a sacrifice deck.
- 1 Eaten Alive | removal | A removal option that fits the sacrifice plan.
- 1 Dictate of Erebos | removal | A removal piece that rewards the deck's theme.
- 1 Grave Pact | removal | A removal piece that rewards the deck's theme.
- 1 Teysa, Orzhov Scion | removal | A removal option aligned with the sacrifice strategy.
- 1 Yawgmoth, Thran Physician | removal | A removal piece that fits the deck's theme.
- 1 Altar of Dementia | synergy | A central sacrifice synergy piece.
- 1 Bartolomé del Presidio | synergy | A sacrifice synergy creature for the deck.
- 1 Bastion of Remembrance | synergy | A sacrifice synergy enchantment.
- 1 Carrion Feeder | synergy | A sacrifice synergy creature for the deck.
- 1 Elas il-Kor, Sadistic Pilgrim | synergy | A sacrifice synergy creature in Karlov's colors.
- 1 Fleshtaker | synergy | A sacrifice synergy creature in Karlov's colors.
- 1 Hidden Stockpile | synergy | An enchantment that supports the sacrifice theme.
- 1 Martyr's Cause | synergy | A sacrifice synergy enchantment.
- 1 Nantuko Husk | synergy | A sacrifice synergy creature for the deck.
- 1 Spawning Pit | synergy | An artifact devoted to sacrifice synergy.
- 1 Viscera Seer | synergy | A sacrifice synergy creature for the deck.
- 1 Witch's Oven | synergy | An artifact devoted to sacrifice synergy.
- 1 Woe Strider | synergy | A sacrifice synergy creature for the deck.
- 1 Zulaport Cutthroat | synergy | A sacrifice synergy creature for the deck.
- 1 Abhorrent Overlord | threat | A black threat that fits the deck's theme.
- 1 Basri's Lieutenant | threat | A white threat for the deck's board presence.
- 1 Corpse Harvester | threat | A black threat that fits the sacrifice plan.
- 1 Demon of Catastrophes | threat | A black threat aligned with the sacrifice theme.
- 1 Demonlord of Ashmouth | threat | A black threat aligned with the sacrifice theme.
- 1 Ghoulcaller Gisa | threat | A black threat for the deck's board presence.
- 1 Kuldotha Forgemaster | threat | A colorless threat for the deck.
- 1 Marrow-Gnawer | threat | A black threat for the deck's board presence.
- 1 Metalwork Colossus | threat | A colorless threat for the deck.
- 1 Mondrak, Glory Dominus | threat | A white threat for the deck's board presence.
- 1 Razaketh, the Foulblooded | threat | A major black threat for the sacrifice plan.
- 1 Requiem Angel | threat | A white threat that fits the deck's theme.
- 1 Toxic Deluge | wipe | A black board wipe for resetting difficult boards.
- 1 The Meathook Massacre | wipe | A board wipe that fits the deck's black core.
- 1 Liliana, Dreadhorde General | wipe | A board wipe option for clearing the way.

</details>

### 12. Commander on a budget

Format: Commander. Theme: tokens. Pool: any_card. Shortlist: 296 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $46.24 to buy, $46.24 the whole deck.

**Summary:** This deck builds a wide token board around Adeline, then reinforces that board with token-focused synergy pieces and follows up with steady combat pressure. It has enough card flow, ramp, removal, and reset buttons to keep participating through a typical multiplayer game, and it wins by making its battlefield presence overwhelming rather than relying on a single expensive finisher. In exchange, it gives up premium lands, explosive mana, and the raw power of higher-priced individual threats.

- [INFO] `bracket_prose_rules`: bracket 2 (Core, 8+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 3.32 over 63 nonland cards

<details><summary>The deck list</summary>

- 36 Plains | land | Primary land slot that keeps the mana base inexpensive and reliable.
- 1 Angelic Sell-Sword | draw | Budget draw slot supporting the deck's consistency.
- 1 Bygone Bishop | draw | Budget draw slot supporting the deck's consistency.
- 1 Dawn of Hope | draw | Budget draw slot supporting the deck's consistency.
- 1 Faramir, Field Commander | draw | Budget draw slot supporting the deck's consistency.
- 1 Idol of Oblivion | draw | Budget draw slot supporting the deck's consistency.
- 1 Glimmer Seeker | draw | Budget draw slot supporting the deck's consistency.
- 1 Platoon Dispenser | draw | Budget draw slot supporting the deck's consistency.
- 1 Sarah Jane Smith | draw | No-cost draw slot supporting the deck's consistency.
- 1 Staff of the Storyteller | draw | Budget draw slot supporting the deck's consistency.
- 1 Wedding Announcement // Wedding Festivity | draw | Budget draw slot that also suits the token-focused plan.
- 1 Charisma Bobblehead | ramp | Budget ramp slot supporting early development.
- 1 Coin of Mastery | ramp | Budget ramp slot supporting early development.
- 1 Collector's Vault | ramp | Budget ramp slot supporting early development.
- 1 Currency Converter | ramp | Budget ramp slot supporting early development.
- 1 Discerning Financier | ramp | Budget ramp slot supporting early development.
- 1 Goldvein Pick | ramp | Low-cost ramp slot for the deck's development.
- 1 Karn, Living Legacy | ramp | Ramp slot that supports the deck's development.
- 1 Keeper of the Accord | ramp | Low-cost ramp slot for the deck's development.
- 1 Noble's Purse | ramp | Low-cost ramp slot for the deck's development.
- 1 Prying Blade | ramp | Low-cost ramp slot for the deck's development.
- 1 Aerial Assault | removal | Low-cost removal slot.
- 1 Banishing Slash | removal | Low-cost removal slot.
- 1 Battle Menu | removal | Low-cost removal slot.
- 1 Citizen's Crowbar | removal | Low-cost removal slot.
- 1 Generous Gift | removal | Flexible budget removal slot.
- 1 Kellan's Lightblades | removal | Low-cost removal slot.
- 1 Righteous Confluence | removal | Flexible budget removal slot.
- 1 Stroke of Midnight | removal | Low-cost removal slot.
- 1 Moogles' Valor | interaction | Low-cost interaction slot for the token-focused plan.
- 1 Paladin's Arms | interaction | Low-cost interaction slot for the token-focused plan.
- 1 Pegasus Guardian // Rescue the Foal | interaction | Low-cost interaction slot for the token-focused plan.
- 1 Rootborn Defenses | interaction | Low-cost interaction slot for the token-focused plan.
- 1 Spirit Bonds | interaction | Budget interaction slot for the token-focused plan.
- 1 Squad Commander | interaction | Low-cost interaction slot for the token-focused plan.
- 1 Hour of Reckoning | wipe | Inexpensive board-wipe slot.
- 1 Martial Coup | wipe | Inexpensive board-wipe slot that fits the deck's plan.
- 1 Phyrexian Rebirth | wipe | Inexpensive board-wipe slot.
- 1 Anafenza, Unyielding Lineage | synergy | Inexpensive synergy piece for the token-focused plan.
- 1 Automated Assembly Line | synergy | Inexpensive synergy piece for the token-focused plan.
- 1 Cat Collector | synergy | Low-cost synergy piece for the token-focused plan.
- 1 Clarion Spirit | synergy | Low-cost synergy piece for the token-focused plan.
- 1 Divine Visitation | synergy | High-impact, low-cost synergy piece for the token-focused plan.
- 1 Felidar Retreat | synergy | Budget synergy piece for the token-focused plan.
- 1 Horn of Gondor | synergy | Token-focused synergy piece at a reasonable price.
- 1 Intangible Virtue | synergy | Low-cost synergy piece for the token-focused plan.
- 1 Mavren Fein, Dusk Apostle | synergy | Low-cost synergy piece for the token-focused plan.
- 1 Oketra's Monument | synergy | Synergy piece that suits the token-focused plan.
- 1 Retrofitter Foundry | synergy | Budget synergy piece for the token-focused plan.
- 1 Rosie Cotton of South Lane | synergy | Low-cost synergy piece for the token-focused plan.
- 1 Siege Veteran | synergy | Low-cost synergy piece for the token-focused plan.
- 1 Worthy Knight | synergy | Low-cost synergy piece for the token-focused plan.
- 1 Ajani's Chosen | threat | Inexpensive threat for the deck's proactive plan.
- 1 Archangel Elspeth | threat | Threat that adds battlefield pressure without straining the budget.
- 1 Attended Healer | threat | Low-cost threat for the deck's proactive plan.
- 1 Basri's Lieutenant | threat | Low-cost threat for the deck's proactive plan.
- 1 Cemetery Protector | threat | Inexpensive threat for the deck's proactive plan.
- 1 Defiler of Faith | threat | Low-cost threat for the deck's proactive plan.
- 1 Dragonback Lancer | threat | Low-cost threat for the deck's proactive plan.
- 1 Drogskol Cavalry | threat | Low-cost threat for the deck's proactive plan.
- 1 Emeria Angel | threat | Low-cost threat that suits the token-focused plan.
- 1 Gideon, Ally of Zendikar | threat | Affordable threat for the deck's proactive plan.
- 1 God-Eternal Oketra | threat | Affordable threat for the deck's proactive plan.
- 1 Hero of Bladehold | threat | Low-cost threat that fits the token-focused plan.

</details>

### 13. owned first, and the commander is not owned

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 229 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $0.00 to buy, $199.35 the whole deck.

**Summary:** This lifegain deck is built around Karlov of the Ghost Council, developing its mana and card access while assembling a board of synergistic permanents, equipment, and threats. It aims to turn that steady development into pressure, clear away important opposing pieces when needed, and close through its threat suite. In exchange for this broad, value-oriented plan, it is less focused on a single dedicated finishing line.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [WARN] `not_owned`: Karlov of the Ghost Council: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.03 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Command Tower | land | Included as a land in the mana base.
- 1 Marsh Flats | land | Included as a land in the mana base.
- 1 Fabled Passage | land | Included as a land in the mana base.
- 1 Evolving Wilds | land | Included as a land in the mana base.
- 1 Terramorphic Expanse | land | Included as a land in the mana base.
- 1 Ash Barrens | land | Included as a land in the mana base.
- 1 Takenuma, Abandoned Mire | land | Included as a land in the mana base.
- 1 Castle Locthwain | land | Included as a land in the mana base.
- 1 War Room | land | Included as a land in the mana base.
- 1 Bonders' Enclave | land | Included as a land in the mana base.
- 1 Reliquary Tower | land | Included as a land in the mana base.
- 1 Rogue's Passage | land | Included as a land in the mana base.
- 1 Windbrisk Heights | land | Included as a land in the mana base.
- 1 Scavenger Grounds | land | Included as a land in the mana base.
- 1 Field of Ruin | land | Included as a land in the mana base.
- 1 Plaza of Heroes | land | Included as a land in the mana base.
- 10 Plains | land | Basic land supporting the mana base.
- 10 Swamp | land | Basic land supporting the mana base.
- 1 Sol Ring | ramp | Included to accelerate the deck's development.
- 1 Arcane Signet | ramp | Included to accelerate the deck's development.
- 1 Fellwar Stone | ramp | Included to accelerate the deck's development.
- 1 Commander's Sphere | ramp | Included to accelerate the deck's development.
- 1 Thought Vessel | ramp | Included to accelerate the deck's development.
- 1 Wayfarer's Bauble | ramp | Included to accelerate the deck's development.
- 1 Sword of the Animist | ramp | Included to accelerate the deck's development.
- 1 Springleaf Drum | ramp | Included to accelerate the deck's development.
- 1 Inherited Envelope | ramp | Included to accelerate the deck's development.
- 1 White Lotus Tile | ramp | Included to accelerate the deck's development.
- 1 Buster Sword | draw | Included to maintain access to cards.
- 1 Call of the Ring | draw | Included to maintain access to cards.
- 1 Exemplar of Light | draw | Included to maintain access to cards.
- 1 Idol of Oblivion | draw | Included to maintain access to cards.
- 1 Inspiring Overseer | draw | Included to maintain access to cards.
- 1 Lembas | draw | Included to maintain access to cards.
- 1 Mask of Memory | draw | Included to maintain access to cards.
- 1 Night's Whisper | draw | Included to maintain access to cards.
- 1 Skullclamp | draw | Included to maintain access to cards.
- 1 Wall of Omens | draw | Included to maintain access to cards.
- 1 Lightning Greaves | interaction | Included to support and protect the central plan.
- 1 Swiftfoot Boots | interaction | Included to support and protect the central plan.
- 1 Champion's Helm | interaction | Included to support and protect the central plan.
- 1 Clever Concealment | interaction | Included to support and protect the central plan.
- 1 Unbreakable Formation | interaction | Included to support and protect the central plan.
- 1 Take Up the Shield | interaction | Included to support and protect the central plan.
- 1 Banishing Light | removal | Included as focused opposing-board control.
- 1 Bitter Triumph | removal | Included as focused opposing-board control.
- 1 Crib Swap | removal | Included as focused opposing-board control.
- 1 Dispatch | removal | Included as focused opposing-board control.
- 1 Fatal Push | removal | Included as focused opposing-board control.
- 1 Generous Gift | removal | Included as focused opposing-board control.
- 1 Get Lost | removal | Included as focused opposing-board control.
- 1 Infernal Grasp | removal | Included as focused opposing-board control.
- 1 Austere Command | wipe | Included as a reset for crowded boards.
- 1 Fumigate | wipe | Included as a reset for crowded boards.
- 1 Vanquish the Horde | wipe | Included as a reset for crowded boards.
- 1 Aerith Gainsborough | synergy | Included as a synergy piece for the deck's central plan.
- 1 Aettir and Priwen | synergy | Included as a synergy piece for the deck's central plan.
- 1 Angel of Vitality | synergy | Included as a synergy piece for the deck's central plan.
- 1 Compassionate Healer | synergy | Included as a synergy piece for the deck's central plan.
- 1 Dancer's Chakrams | synergy | Included as a synergy piece for the deck's central plan.
- 1 Elixir | synergy | Included as a synergy piece for the deck's central plan.
- 1 Excalibur II | synergy | Included as a synergy piece for the deck's central plan.
- 1 Graveyard Trespasser // Graveyard Glutton | synergy | Included as a synergy piece for the deck's central plan.
- 1 Kor Firewalker | synergy | Included as a synergy piece for the deck's central plan.
- 1 Light of Promise | synergy | Included as a synergy piece for the deck's central plan.
- 1 Northern Air Temple | synergy | Included as a synergy piece for the deck's central plan.
- 1 Pull from the Grave | synergy | Included as a synergy piece for the deck's central plan.
- 1 Second Breakfast | synergy | Included as a synergy piece for the deck's central plan.
- 1 White Mage's Staff | synergy | Included as a synergy piece for the deck's central plan.
- 1 Angel of Invention | threat | Included as a threat that advances the win plan.
- 1 Bill the Pony | threat | Included as a threat that advances the win plan.
- 1 Dawnhand Eulogist | threat | Included as a threat that advances the win plan.
- 1 Gollum, Silent Slinker // Meager Meal | threat | Included as a threat that advances the win plan.
- 1 Lo and Li, Twin Tutors | threat | Included as a threat that advances the win plan.
- 1 Lyra Dawnbringer | threat | Included as a threat that advances the win plan.
- 1 Minwu, White Mage | threat | Included as a threat that advances the win plan.
- 1 Rabaroo Troop | threat | Included as a threat that advances the win plan.
- 1 Rooftop Percher | threat | Included as a threat that advances the win plan.
- 1 Shattered Angel | threat | Included as a threat that advances the win plan.
- 1 Sneering Shadewriter | threat | Included as a threat that advances the win plan.
- 1 Victory's Herald | threat | Included as a threat that advances the win plan.

</details>

### 14. the user delegates the commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 244 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $641.51 to buy, $641.51 the whole deck.

**Summary:** This deck ramps into a board of Dragons, backs that board with Dragon-focused support, and uses removal and protective interaction to keep attacking. It wins by turning its large Dragon threats into overwhelming combat pressure, with Atarka, World Render at the head of the plan. It gives up broad utility and a low curve in exchange for a concentrated, creature-heavy Dragon strategy.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 4.40 over 63 nonland cards

<details><summary>The deck list</summary>

- 12 Mountain | land | Basic red land for the Dragon-focused mana base.
- 12 Forest | land | Basic green land for the Dragon-focused mana base.
- 1 Ancient Tomb | land | Colorless land that supports the mana base.
- 1 Boseiju, Who Endures | land | Green legendary land for the mana base.
- 1 Cavern of Souls | land | Tribal land that supports the Dragon plan.
- 1 Cinder Glade | land | Red-green land for the mana base.
- 1 Command Tower | land | Reliable multicolor land for the commander deck.
- 1 Crucible of the Spirit Dragon | land | Dragon-themed land for the mana base.
- 1 Evolving Wilds | land | Flexible land for the mana base.
- 1 Exotic Orchard | land | Flexible multicolor land for the mana base.
- 1 Fabled Passage | land | Flexible land for the mana base.
- 1 Haven of the Spirit Dragon | land | Dragon-themed land for the mana base.
- 1 Maelstrom of the Spirit Dragon | land | Dragon-themed land for the mana base.
- 1 Path of Ancestry | land | Tribal land that supports the Dragon plan.
- 1 Ancient Copper Dragon | ramp | Dragon ramp piece that advances the main creature plan.
- 1 Carnelian Orb of Dragonkind | ramp | Dragon-themed artifact ramp.
- 1 Dragon's Hoard | ramp | Dragon-themed artifact ramp.
- 1 Ganax, Astral Hunter | ramp | Dragon ramp piece that fits the creature base.
- 1 Goldspan Dragon | ramp | Dragon ramp piece that adds pressure.
- 1 Jade Orb of Dragonkind | ramp | Dragon-themed artifact ramp.
- 1 Klauth, Unrivaled Ancient | ramp | Dragon ramp piece for the top end.
- 1 Old Gnawbone | ramp | Dragon ramp piece that supports expensive threats.
- 1 Orb of Dragonkind | ramp | Dragon-themed artifact ramp.
- 1 Scaled Nurturer | ramp | Early Dragon-themed ramp.
- 1 Avaricious Dragon | draw | Dragon-based card draw for the tribal plan.
- 1 Beast Whisperer | draw | Creature-focused draw that supports a Dragon-heavy build.
- 1 Dragon Mage | draw | Dragon-based draw for the threat suite.
- 1 Dragonborn Champion | draw | Dragon-themed draw support.
- 1 Elemental Bond | draw | Creature-focused draw for large Dragons.
- 1 Garruk's Uprising | draw | Creature-focused draw for the Dragon plan.
- 1 Guardian Project | draw | Creature-based draw for a varied Dragon roster.
- 1 Harmonize | draw | Straightforward green card draw.
- 1 Return of the Wildspeaker | draw | Creature-focused draw that fits the large-threat plan.
- 1 Rishkar's Expertise | draw | High-impact draw for a deck built around large creatures.
- 1 Fog | interaction | Defensive interaction for protecting the board position.
- 1 Heroic Intervention | interaction | Protective interaction for the Dragon board.
- 1 Lightning Greaves | interaction | Equipment interaction that supports key creatures.
- 1 Snakeskin Veil | interaction | Low-cost protective interaction.
- 1 Swiftfoot Boots | interaction | Equipment interaction that supports key creatures.
- 1 Tamiyo's Safekeeping | interaction | Protective interaction for important permanents.
- 1 Bogardan Hellkite | removal | Dragon removal that remains a substantial body.
- 1 Dragon's Fire | removal | Efficient Dragon-themed removal.
- 1 Dragon Tempest | removal | Dragon-themed removal support.
- 1 Dragonlord Atarka | removal | Dragon removal that fits the top end.
- 1 Drakuseth, Maw of Flames | removal | Large Dragon that supplies removal.
- 1 Invasion of Tarkir // Defiant Thundermaw | removal | Dragon-themed removal with a Dragon back face.
- 1 Scourge of Valkas | removal | Dragon removal that rewards the tribal plan.
- 1 Terror of the Peaks | removal | Dragon removal attached to a major threat.
- 1 Balefire Dragon | wipe | Dragon board wipe for creature-heavy tables.
- 1 Draconic Intervention | wipe | Dragon-themed board wipe.
- 1 Ryusei, the Falling Star | wipe | Dragon board wipe that fits the creature plan.
- 1 Acolyte of Bahamut | synergy | Dragon-focused synergy piece.
- 1 Breaching Dragonstorm | synergy | Dragon-themed synergy for the core plan.
- 1 Crucible of Fire | synergy | Tribal Dragon synergy.
- 1 Dracogenesis | synergy | Dragon-themed synergy for the top-end plan.
- 1 Dragon Egg | synergy | Low-cost Dragon synergy piece.
- 1 Dragon Hatchling | synergy | Low-cost Dragon synergy piece.
- 1 Dragonkin Berserker | synergy | Dragon-themed synergy support.
- 1 Dragonlord's Servant | synergy | Dragon-focused synergy support.
- 1 Dragonspeaker Shaman | synergy | Dragon-focused synergy support.
- 1 Dragonstorm | synergy | Dragon-themed payoff for the creature plan.
- 1 Firespitter Whelp | synergy | Dragon synergy that contributes to the creature base.
- 1 Sarkhan's Triumph | synergy | Dragon-themed synergy and deck support.
- 1 Shivan Devastator | synergy | Flexible Dragon that supports the tribal plan.
- 1 The Dragon-Kami Reborn // Dragon-Kami's Egg | synergy | Dragon-themed saga synergy.
- 1 Ancient Bronze Dragon | threat | Premium Dragon threat for closing games through combat.
- 1 Backdraft Hellkite | threat | Dragon threat that strengthens the attacking roster.
- 1 Blast-Furnace Hellkite | threat | High-impact Dragon threat.
- 1 Canopy Gargantuan | threat | Large Dragon threat for the top end.
- 1 Dragon Broodmother | threat | Dragon threat that reinforces the tribal board.
- 1 Hellkite Charger | threat | Aggressive Dragon threat.
- 1 Lathliss, Dragon Queen | threat | Dragon threat with strong tribal presence.
- 1 Scourge of the Throne | threat | Combat-oriented Dragon threat.
- 1 Terror of Mount Velus | threat | Large Dragon threat for decisive combat turns.
- 1 Thrakkus the Butcher | threat | Dragon threat that fits the combat plan.
- 1 Thunderbreak Regent | threat | Efficient Dragon threat for the main roster.
- 1 Utvara Hellkite | threat | Top-end Dragon threat for finishing games.

</details>

### 15. delegated commander, owned first

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 229 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $0.00 to buy, $143.21 the whole deck.

**Summary:** This is a steady black-white lifegain deck built around Cecil and a creature-heavy board, with Angels, supportive permanents, and equipment helping it establish pressure. It wins by developing durable threats while using targeted answers and resets to keep opposing boards contained. In exchange for that focused creature plan, it leans on its selected draw pieces to keep resources flowing and can be slower to recover when several key permanents are answered at once.

- [INFO] `bracket_prose_rules`: bracket 2 (Core, 8+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 3.05 over 63 nonland cards

<details><summary>The deck list</summary>

- 20 Plains | land | Basic white land for the deck’s mana base.
- 16 Swamp | land | Basic black land for the deck’s mana base.
- 1 Sol Ring | ramp | A shortlist ramp option for accelerating the deck.
- 1 Arcane Signet | ramp | A shortlist ramp option for accelerating the deck.
- 1 Fellwar Stone | ramp | A shortlist ramp option for accelerating the deck.
- 1 Wayfarer's Bauble | ramp | A shortlist ramp option for accelerating the deck.
- 1 Bender's Waterskin | ramp | A shortlist ramp option for accelerating the deck.
- 1 Commander's Sphere | ramp | A shortlist ramp option for accelerating the deck.
- 1 Thought Vessel | ramp | A shortlist ramp option for accelerating the deck.
- 1 Sword of the Animist | ramp | A shortlist ramp option that also fits the equipment package.
- 1 Giada, Font of Hope | ramp | A shortlist ramp option that complements the Angel threats.
- 1 Lotho, Corrupt Shirriff | ramp | A shortlist ramp option for the black-white shell.
- 1 Exemplar of Light | draw | A shortlist draw option that fits the deck’s lifegain focus.
- 1 Inspiring Overseer | draw | A shortlist draw option that also fits the Angel package.
- 1 Lembas | draw | A shortlist draw option for the deck.
- 1 Night's Whisper | draw | A shortlist draw option for the deck.
- 1 Painful Truths | draw | A shortlist draw option for the deck.
- 1 Wall of Omens | draw | A shortlist draw option for the deck.
- 1 Skullclamp | draw | A shortlist draw option that fits the deck’s equipment suite.
- 1 Tome of Legends | draw | A shortlist draw option for the commander deck.
- 1 Buster Sword | draw | A shortlist draw option that fits the equipment suite.
- 1 Call of the Ring | draw | A shortlist draw option for the deck.
- 1 Lightning Greaves | interaction | A shortlist interaction option for protecting a key creature.
- 1 Swiftfoot Boots | interaction | A shortlist interaction option for protecting a key creature.
- 1 Bastion Protector | interaction | A shortlist interaction option for the commander-focused plan.
- 1 Clever Concealment | interaction | A shortlist interaction option for safeguarding the board.
- 1 Darksteel Plate | interaction | A shortlist interaction option that fits the equipment suite.
- 1 Take Up the Shield | interaction | A shortlist interaction option for a key creature.
- 1 Swords to Plowshares | removal | A shortlist removal option for opposing threats.
- 1 Generous Gift | removal | A shortlist removal option with broad coverage.
- 1 Get Lost | removal | A shortlist removal option for problematic permanents.
- 1 Banishing Light | removal | A shortlist removal option for problematic permanents.
- 1 Infernal Grasp | removal | A shortlist removal option for opposing creatures.
- 1 Bitter Triumph | removal | A shortlist removal option for opposing threats.
- 1 Stroke of Midnight | removal | A shortlist removal option for problematic permanents.
- 1 Crib Swap | removal | A shortlist removal option for opposing creatures.
- 1 Aerith Gainsborough | synergy | A shortlist synergy piece for the deck’s lifegain focus.
- 1 Angel of Vitality | synergy | A shortlist synergy piece for the deck’s lifegain focus.
- 1 Compassionate Healer | synergy | A shortlist synergy piece for the deck’s lifegain focus.
- 1 Kor Firewalker | synergy | A shortlist synergy piece for the deck’s lifegain focus.
- 1 Light of Promise | synergy | A shortlist synergy piece for the deck’s lifegain focus.
- 1 Night Nurse, Healer of Heroes | synergy | A shortlist synergy piece for the deck’s lifegain focus.
- 1 Rosie Cotton of South Lane | synergy | A shortlist synergy piece for the deck’s lifegain focus.
- 1 Second Breakfast | synergy | A shortlist synergy piece for the deck’s lifegain focus.
- 1 Aettir and Priwen | synergy | A shortlist synergy piece that fits the equipment package.
- 1 White Mage's Staff | synergy | A shortlist synergy piece that fits the equipment package.
- 1 Well-Worn Spatula | synergy | A shortlist synergy piece that fits the equipment package.
- 1 Prideful Feastling | synergy | A shortlist synergy piece for the deck’s lifegain focus.
- 1 Adventurous Eater // Have a Bite | synergy | A shortlist synergy piece for the deck’s lifegain focus.
- 1 Elixir | synergy | A shortlist synergy piece for the deck’s lifegain focus.
- 1 Angel of Invention | threat | A shortlist threat that fits the Angel package.
- 1 Bill the Pony | threat | A shortlist threat for advancing the board.
- 1 Dawnhand Eulogist | threat | A shortlist threat for advancing the board.
- 1 Foggy Swamp Hunters | threat | A shortlist threat for advancing the board.
- 1 Gollum, Silent Slinker // Meager Meal | threat | A shortlist threat for advancing the board.
- 1 Lyra Dawnbringer | threat | A shortlist threat that fits the Angel package.
- 1 Minwu, White Mage | threat | A shortlist threat that fits the deck’s lifegain focus.
- 1 Rabaroo Troop | threat | A shortlist threat for advancing the board.
- 1 Rooftop Percher | threat | A shortlist threat for advancing the board.
- 1 Shattered Angel | threat | A shortlist threat that fits the Angel package.
- 1 Sneering Shadewriter | threat | A shortlist threat for advancing the board.
- 1 Victory's Herald | threat | A shortlist threat that fits the Angel package.
- 1 Austere Command | wipe | A shortlist board wipe for resetting difficult boards.
- 1 Fumigate | wipe | A shortlist board wipe for resetting difficult boards.
- 1 Vanquish the Horde | wipe | A shortlist board wipe for resetting difficult boards.

</details>

### 16. a tight budget, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 229 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $6.17 to buy, $35.94 the whole deck.

**Summary:** This Denethor deck develops its mana, keeps cards moving, and builds around a dense aristocrats shell of synergy creatures and supporting permanents. It aims to take over creature-heavy games by applying pressure with its threats while using targeted answers, protection, and broad resets to keep opposing boards manageable. It gives up premium mana fixing and expensive standalone finishers in exchange for a library-first, low-cost build with a focused creature plan.

- [INFO] `bracket_prose_rules`: bracket 2 (Core, 8+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [WARN] `not_owned`: Bastion of Remembrance: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Falkenrath Noble: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Woe Strider: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Zulaport Cutthroat: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Aron, Benalia's Ruin: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ayli, Eternal Pilgrim: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Baron Bertram Graywater: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bartolomé del Presidio: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Disciple of Bolas: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Erebos, Bleak-Hearted: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Lord Skitter's Butcher: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Skullport Merchant: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.00 over 63 nonland cards

<details><summary>The deck list</summary>

- 18 Plains | land | Basic land for the white half of the mana base.
- 18 Swamp | land | Basic land for the black half of the mana base.
- 1 Arcane Signet | ramp | Owned mana acceleration.
- 1 Astral Cornucopia | ramp | Owned mana acceleration.
- 1 Bender's Waterskin | ramp | Owned mana acceleration.
- 1 Blitzball | ramp | Owned mana acceleration.
- 1 Chromatic Lantern | ramp | Owned mana acceleration.
- 1 Commander's Sphere | ramp | Owned mana acceleration.
- 1 Deadly Dispute | ramp | Owned ramp card that fits the deck's overall plan.
- 1 Fellwar Stone | ramp | Owned mana acceleration.
- 1 Inherited Envelope | ramp | Owned mana acceleration.
- 1 Sol Ring | ramp | Owned mana acceleration.
- 1 Call of the Ring | draw | Owned card flow for the deck.
- 1 Grave Venerations | draw | Owned card flow for the deck.
- 1 Idol of Oblivion | draw | Owned card flow for the deck.
- 1 Inspiring Overseer | draw | Owned creature that supplies card flow.
- 1 Mask of Memory | draw | Owned card flow for the deck.
- 1 Nasty End | draw | Owned card flow for the deck.
- 1 Painful Truths | draw | Owned card flow for the deck.
- 1 Skullclamp | draw | Owned card flow for the deck.
- 1 Tome of Legends | draw | Owned card flow for the deck.
- 1 Wall of Omens | draw | Owned creature that supplies card flow.
- 1 Bastion Protector | interaction | Owned protection for key permanents.
- 1 Frontline Medic | interaction | Owned creature that supports the defensive package.
- 1 Gift of Immortality | interaction | Owned protection piece.
- 1 Swiftfoot Boots | interaction | Owned protection equipment.
- 1 Together Forever | interaction | Owned interaction piece for the creature-focused plan.
- 1 Unbreakable Formation | interaction | Owned defensive interaction.
- 1 Angel of Serenity | removal | Owned creature-based answer.
- 1 Banishing Light | removal | Owned removal spell.
- 1 Bitter Triumph | removal | Owned removal spell.
- 1 Claim the Precious | removal | Owned removal spell.
- 1 Crib Swap | removal | Owned removal spell.
- 1 Fiend Hunter | removal | Owned creature-based answer.
- 1 Gaius van Baelsar | removal | Owned creature-based answer.
- 1 Generous Gift | removal | Owned removal spell.
- 1 Austere Command | wipe | Owned broad reset option.
- 1 Dusk // Dawn | wipe | Owned reset option that suits a creature deck.
- 1 Fumigate | wipe | Owned broad reset option.
- 1 Al Bhed Salvagers | synergy | Owned synergy creature for the aristocrats shell.
- 1 Arcade Cabinet | synergy | Owned synergy artifact for the deck's plan.
- 1 Gollum the Abandoned | synergy | Owned synergy creature for the aristocrats shell.
- 1 Gollum, Patient Plotter | synergy | Owned synergy creature for the aristocrats shell.
- 1 Gríma Wormtongue | synergy | Owned synergy creature for the deck's plan.
- 1 Heirloom Auntie | synergy | Owned synergy creature for the deck's plan.
- 1 Joo Dee, One of Many | synergy | Owned synergy creature for the deck's plan.
- 1 Nimble Hobbit | synergy | Owned synergy creature for the deck's plan.
- 1 Phantom Train | synergy | Owned synergy artifact for the deck's plan.
- 1 Rhovanion Rampager | synergy | Owned synergy creature for the aristocrats shell.
- 1 Bastion of Remembrance | synergy | Low-cost addition that supports the aristocrats shell.
- 1 Falkenrath Noble | synergy | Low-cost creature for the aristocrats shell.
- 1 Woe Strider | synergy | Low-cost creature for the aristocrats shell.
- 1 Zulaport Cutthroat | synergy | Low-cost creature for the aristocrats shell.
- 1 Bill the Pony | threat | Owned creature that contributes to board pressure.
- 1 Hei Bai, Spirit of Balance | threat | Owned creature that contributes to board pressure.
- 1 Namazu Trader | threat | Owned creature that contributes to board pressure.
- 1 Vengeful Villagers | threat | Owned creature that contributes to board pressure.
- 1 Aron, Benalia's Ruin | threat | Low-cost creature that supports the creature-heavy plan.
- 1 Ayli, Eternal Pilgrim | threat | Low-cost creature for the deck's main plan.
- 1 Baron Bertram Graywater | threat | Low-cost creature for the deck's main plan.
- 1 Bartolomé del Presidio | threat | Low-cost creature for the deck's main plan.
- 1 Disciple of Bolas | threat | Low-cost creature for the deck's main plan.
- 1 Erebos, Bleak-Hearted | threat | Low-cost permanent that adds pressure to the deck.
- 1 Lord Skitter's Butcher | threat | Low-cost creature for the deck's main plan.
- 1 Skullport Merchant | threat | Low-cost creature for the deck's main plan.

</details>

### 17. upgrade a precon, owned first

Format: Commander. Theme: goblins. Pool: owned_first. Shortlist: 196 names.

Cards: 99 main, 0 sideboard. Repair turn: true. Block findings: 0.

Cost: $73.55 to buy, $249.18 the whole deck.

**Summary:** A mono-red Goblin Storm deck that builds a broad token board, turns cheap targeted spells into sweeping card-flow and combat bursts, and closes through sacrifice damage, storm payoffs, and hasty tribal attacks. Its mana engines support explosive turns while flexible removal protects the board’s momentum.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 2.48 over 61 nonland cards
- [INFO] `precon_cards_restored`: the deck was 1 card short of the Goblin Storm precon share, so the builder put 1 card back
- [INFO] `basics_added`: the list was 1 card short, so the builder added 1 basic land

<details><summary>The deck list</summary>

- 1 Ancestors' Aid | synergy | Targeted combat trick that scales across the Goblin board with the commander.
- 1 Arena of Glory | land | Haste access for a key Goblin threat.
- 1 Battle Hymn | ramp | Converts a wide Goblin board into a burst of mana.
- 1 Blasphemous Act | wipe | Efficient reset against larger creature boards.
- 1 Boggart Shenanigans | removal | Turns Goblin deaths into direct damage.
- 1 Castle Embereth | land | Land that can turn a token board into meaningful pressure.
- 1 Chaos Warp | removal | Flexible answer to troublesome permanents.
- 1 Conspicuous Snoop | synergy | Goblin value engine that can play creatures from the library top.
- 1 Crimson Wisps | draw | Cheap cantrip that gives a creature haste.
- 1 Daring Discovery | synergy | Flexible spell slot supporting the deck's proactive plan.
- 1 Den of the Bugbear | land | Creature-land that adds pressure without using a spell slot.
- 1 Dragon Fodder | synergy | Produces multiple Goblins for go-wide and spell-copy turns.
- 1 Dwarven Mine | land | Mountain land that can add a Goblin body later.
- 1 Empty the Warrens | wincon | Storm payoff that establishes a large Goblin board.
- 1 Expedite | draw | One-mana cantrip that grants haste to a key creature.
- 1 Faithless Looting | draw | Efficient card selection and graveyard setup.
- 1 Fists of Flame | draw | Card-drawing combat payoff for a wide spell turn.
- 1 Forgotten Cave | land | Cycling land that can become a fresh card late.
- 1 Fountainport | land | Utility land for a board that makes expendable tokens.
- 1 Frontline Heroism | synergy | Supports combat-focused token pressure.
- 1 Gempalm Incinerator | removal | Goblin-scaled creature removal with cycling utility.
- 1 General Kreat, the Boltbringer | synergy | Goblin-focused value piece for the go-wide plan.
- 1 Glimpse the Impossible | ramp | Impulse card flow that can also fuel mana development.
- 1 Goblin Bombardment | removal | Converts tokens into repeatable direct damage and removal.
- 1 Goblin Burrows | land | Land-based pump for an attacking Goblin force.
- 1 Goblin Bushwhacker | synergy | Provides haste and a team-wide burst of damage.
- 1 Goblin Chieftain | synergy | Lord effect that gives Goblins haste and greater pressure.
- 1 Goblin Dark-Dwellers | threat | Reuses an impactful instant or sorcery from the graveyard.
- 1 Goblin Lackey | synergy | Accelerates large Goblins onto the battlefield through combat.
- 1 Goblin Matron | synergy | Finds the Goblin best suited to the current board.
- 1 Goblin Negotiation | removal | Goblin-themed interaction for opposing threats.
- 1 Goblin Trashmaster | removal | Goblin lord that also answers artifacts.
- 1 Goblin Warchief | synergy | Reduces Goblin costs and grants haste.
- 1 Grapeshot | removal | Storm payoff that can clear small creatures or finish weakened opponents.
- 1 Haze of Rage | wincon | Repeatable storm combat finisher for a token board.
- 1 Hidden Volcano | land | Utility land that remains within the mana base.
- 1 Idol of Oblivion | draw | Reliable card draw alongside frequent token creation.
- 1 Impact Tremors | wincon | Turns every Goblin and token burst into table-wide damage.
- 1 Kher Keep | land | Creates cheap bodies for sacrifice and go-wide synergies.
- 1 Krenko's Command | synergy | Efficiently creates multiple Goblins.
- 1 Krenko, Mob Boss | threat | Major token engine and a must-answer Goblin threat.
- 1 Mana Geyser | ramp | Explosive mana ritual for a decisive spell chain.
- 1 Mogg War Marshal | synergy | Makes multiple Goblin bodies from a single card.
- 1 Past in Flames | wincon | Lets a stocked graveyard become another storm turn.
- 1 Quest for the Goblin Lord | synergy | Rewards repeated Goblin deployment with a powerful anthem.
- 1 Reliquary Tower | land | Utility land that preserves a large drawn hand.
- 1 Roaming Throne | synergy | Doubles relevant Goblin triggered abilities.
- 1 Ruby Medallion | ramp | Reduces the cost of the deck's dense red spell suite.
- 1 Rundvelt Hordemaster | synergy | Goblin lord that provides card advantage when Goblins die.
- 1 Skullclamp | draw | Exceptional card draw when equipped to expendable tokens.
- 1 Smoldering Crater | land | Cycling land that offers late-game reach.
- 1 Sol Ring | ramp | Fast, efficient mana acceleration.
- 1 Storm-Kiln Artist | ramp | Creates Treasure throughout a spell-heavy turn.
- 1 Swiftfoot Boots | interaction | Protects a key engine creature while enabling haste.
- 1 Throne of Eldraine | ramp | Mana rock that supports a long-game board presence.
- 1 Vandalblast | wipe | Efficient artifact interaction with a powerful overload mode.
- 1 War Room | land | Land-based card draw when resources run low.
- 1 Wild Ride | synergy | Combat-focused card flow for an aggressive creature plan.
- 1 Witch's Mark | draw | Filters cards while helping a creature attack effectively.
- 1 Ancestral Anger | draw | Cheap targeting cantrip that becomes stronger with each cast.
- 1 Arcane Signet | ramp | Reliable early mana acceleration.
- 1 Brightstone Ritual | ramp | Turns a Goblin board into a large surge of red mana.
- 1 Fable of the Mirror-Breaker // Reflection of Kiki-Jiki | ramp | Provides Treasure, card filtering, and a potent creature-copy engine.
- 1 Broadside Bombardiers | removal | Turns spare Goblins or artifacts into high-impact removal.
- 1 Goblin Fireleaper | removal | Goblin-based removal that benefits from a wide board.
- 1 Goblin-town Flunkies | synergy | Efficient Goblin body for tribal pressure.
- 1 Gristle Glutton | draw | Goblin body that supplies additional card flow.
- 1 Humble Defector | draw | Low-cost creature that draws a large burst of cards.
- 1 Impulsive Pilferer | ramp | Early Goblin that leaves behind Treasure when it dies.
- 1 Lightning Bolt | removal | Efficient answer to creatures, planeswalkers, or damaged opponents.
- 1 Pashalik Mons | removal | Goblin payoff that converts tokens and sacrifices into damage.
- 1 Seething Song | ramp | Powerful ritual for explosive multi-spell turns.
- 1 Skirk Prospector | ramp | Sacrifices Goblins for mana to extend explosive turns.
- 1 Shinka, the Bloodsoaked Keep | interaction | the deck upgrades a precon, and this card is one the precon holds
- 25 Mountain | land | Stable red mana base for a mono-red Goblin strategy.

</details>

### 18. upgrade a precon, any card

Format: Commander. Theme: turtles. Pool: any_card. Shortlist: 303 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $278.09 to buy, $278.09 the whole deck.

**Summary:** This keeps the Turtle Power core and its large themed cast intact, using the commander’s Turtle focus alongside the retained synergy pieces and creature threats to pressure the table. The added draw, ramp, removal, interaction, and wipe cards make the deck more reliable at a bracket 3 table while leaving the precon’s character and broad Turtle-centered plan in place. It gives up a handful of lower-impact original slots rather than trying to turn the deck into a fully optimized generic good-stuff list.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [WARN] `land_count`: 39 lands: the guide range for this format is 34 to 38
- [INFO] `curve_summary`: average mana value 3.35 over 60 nonland cards

<details><summary>The deck list</summary>

- 1 Acidic Slime | other | Retained from the precon as part of its existing creature package.
- 1 April O'Neil, Live on the Scene | draw | Retained from the precon for card draw.
- 1 Arcade Cabinet | other | Retained from the precon to preserve its original theme.
- 1 Arcane Signet | ramp | Retained from the precon for ramp.
- 1 Ash Barrens | land | Retained from the precon as a land.
- 1 Assassin's Trophy | removal | Retained from the precon for removal.
- 1 Baxter, Fly in the Ointment | other | Retained from the precon to preserve its original theme.
- 1 Bebop, Skull & Crossbones | other | Retained from the precon to preserve its original theme.
- 1 Big Apple, 3 a.m. | land | Retained from the precon as a land.
- 1 Big Mother Mouser | other | Retained from the precon as part of its existing creature package.
- 1 Biogenic Ooze | other | Retained from the precon as part of its existing creature package.
- 1 Blasphemous Act | wipe | Retained from the precon as a board wipe.
- 1 Casey Jones, Back Alley Brute | other | Retained from the precon to preserve its original theme.
- 1 Chromatic Lantern | ramp | Retained from the precon for ramp.
- 1 Cinder Glade | land | Retained from the precon as a land.
- 1 City of Brass | land | Retained from the precon as a land.
- 1 Coin of Mastery | other | Retained from the precon to preserve its original theme.
- 1 Command Tower | land | Retained from the precon as a land.
- 1 Continue? | other | Retained from the precon to preserve its original theme.
- 1 Corpsejack Menace | other | Retained from the precon as part of its existing creature package.
- 1 Cultivate | ramp | Retained from the precon for ramp.
- 1 Donatello, the Brains | synergy | Retained from the precon for Turtle-themed synergy.
- 1 Dimension X Pizzasaur | other | Retained from the precon to preserve its original theme.
- 1 Double Jump // Flying Kick | other | Retained from the precon to preserve its original theme.
- 1 Electric Seaweed | other | Retained from the precon to preserve its original theme.
- 1 Endless Foot Assault | other | Retained from the precon to preserve its original theme.
- 1 Escape Tunnel | land | Retained from the precon as a land.
- 1 Everything Pizza | other | Retained from the precon to preserve its original theme.
- 1 Fast Forward | other | Retained from the precon to preserve its original theme.
- 1 Foot Chopper | other | Retained from the precon to preserve its original theme.
- 4 Forest | land | Basic land retained for the precon's mana base.
- 1 Game Over | other | Retained from the precon to preserve its original theme.
- 1 Grand Coliseum | land | Retained from the precon as a land.
- 1 Harmonize | other | Retained from the precon as part of its existing support package.
- 1 Here Comes a New Hero! | other | Retained from the precon to preserve its original theme.
- 1 Hidden Hideout | land | Retained from the precon as a land.
- 1 High Score | other | Retained from the precon to preserve its original theme.
- 1 Hinterland Harbor | land | Retained from the precon as a land.
- 4 Island | land | Basic land retained for the precon's mana base.
- 1 Krang, the All-Powerful | other | Retained from the precon to preserve its original theme.
- 1 Leatherhead, Iron Gator | other | Retained from the precon to preserve its original theme.
- 1 Lessons from Life | other | Retained from the precon to preserve its original theme.
- 1 Level Up | other | Retained from the precon to preserve its original theme.
- 1 Lita, Little Orphan Amphibian | synergy | Retained from the precon for Turtle-themed synergy.
- 1 Mole Module | other | Retained from the precon to preserve its original theme.
- 3 Mountain | land | Basic land retained for the precon's mana base.
- 1 Mona Lisa, Science Geek | other | Retained from the precon to preserve its original theme.
- 1 Ninja Pizza | other | Retained from the precon to preserve its original theme.
- 1 Path of Ancestry | land | Retained from the precon as a land.
- 3 Plains | land | Basic land retained for the precon's mana base.
- 1 Rain-Slicked Copse | land | Retained from the precon as a land.
- 1 Rat King, Pale Piper | other | Retained from the precon to preserve its original theme.
- 1 Ray Fillet, Wave Warrior | other | Retained from the precon to preserve its original theme.
- 1 Roadkill Rodney | other | Retained from the precon to preserve its original theme.
- 1 Rocksteady, Mutant Marauder | other | Retained from the precon to preserve its original theme.
- 1 Rootbound Crag | land | Retained from the precon as a land.
- 1 Shellshock | other | Retained from the precon to preserve its original theme.
- 1 Shredder, Shadow Master | other | Retained from the precon to preserve its original theme.
- 1 Smoldering Marsh | land | Retained from the precon as a land.
- 1 Sodden Verdure | land | Retained from the precon as a land.
- 1 Sol Ring | ramp | Retained from the precon for ramp.
- 1 Special Move | other | Retained from the precon to preserve its original theme.
- 1 Spire Garden | land | Retained from the precon as a land.
- 1 Splinter, the Mentor | other | Retained from the precon to preserve its original theme.
- 1 Steelbane Hydra | removal | Retained from the precon for removal.
- 1 Super Combo | other | Retained from the precon to preserve its original theme.
- 3 Swamp | land | Basic land retained for the precon's mana base.
- 1 Swift Demise | other | Retained from the precon to preserve its original theme.
- 1 Tempestra, Dame of Games | other | Retained from the precon to preserve its original theme.
- 1 Thriving Grove | land | Retained from the precon as a land.
- 1 Thriving Isle | land | Retained from the precon as a land.
- 1 Thriving Moor | land | Retained from the precon as a land.
- 1 Together Forever | other | Retained from the precon to preserve its original theme.
- 1 Tokka & Rahzar, Unsupervised | ramp | Retained from the precon for Turtle-themed ramp.
- 1 Turtle Lair | land | Retained from the precon as a land.
- 1 Undergrowth Stadium | land | Retained from the precon as a land.
- 1 Vanquish the Horde | wipe | Retained from the precon as a board wipe.
- 1 Vernal Fen | land | Retained from the precon as a land.
- 1 Vibrant Cityscape | land | Retained from the precon as a land.
- 1 Vigor | other | Retained from the precon as part of its existing creature package.
- 1 Voracious Hydra | other | Retained from the precon as part of its existing creature package.
- 1 Wave Goodbye | other | Retained from the precon to preserve its original theme.
- 1 Rhystic Study | draw | Added as a stronger dedicated draw card.
- 1 Smothering Tithe | ramp | Added as a stronger dedicated ramp card.
- 1 Swords to Plowshares | removal | Added as efficient removal.
- 1 Heroic Intervention | interaction | Added to strengthen interaction.
- 1 Cyclonic Rift | wipe | Added as a flexible board wipe.

</details>

