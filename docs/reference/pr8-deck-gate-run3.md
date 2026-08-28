# PR-8 deck gate

Run date: 2026-08-27. Card snapshot: 2026-08-24.

Verdict: PASS. 13 of 13 decks passed every block check, and 0 invented names reached the user. Both bars are zero tolerance.

## Summary

| Measure | Value |
|---|---|
| Prompts | 13 |
| Decks returned | 13 |
| Decks with no block finding | 13 |
| Invented names that reached the user | 0 |
| Decks that needed the repair turn | 1 |
| Errors | 0 |
| Prompt version | 2 |
| Calls | 14 |
| Cost | $0.7173 |
| Time | 567 seconds |

## Findings by code

A block stops the deck. A warning and a note are reports.

| Code | Count |
|---|---|
| `curve_summary` | 13 |
| `bracket_prose_rules` | 7 |
| `basics_added` | 1 |
| `not_owned` | 1 |

By severity: BLOCK 0. WARN 1. INFO 21. 

## Decks

### 1. lifegain Commander, any card

Format: Commander. Theme: lifegain. Pool: any_card. Shortlist: 303 names.

Cards: 99. Repair turn: false. Block findings: 0.

**Summary:** Karlov leads a lifegain-focused white-black deck that develops its mana and hand, then turns that engine into pressure through creatures and supporting permanents. It aims to win by establishing a threatening board while using focused answers and resets to keep opposing boards from taking over. The tradeoff is a slower, board-centered plan that needs time to assemble its supporting pieces and can lose momentum after a broad reset.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 3.63 over 63 nonland cards

<details><summary>The deck list</summary>

- 18 Plains | land | Provides a reliable white mana base for the deck.
- 18 Swamp | land | Provides a reliable black mana base for the deck.
- 1 Alhammarret's Archive | draw | Included as a draw card to keep resources flowing.
- 1 Archivist of Oghma | draw | Included as a draw card to maintain card flow.
- 1 Dawn of Hope | draw | Included as a draw card for sustained resources.
- 1 Enduring Innocence | draw | Included as a draw card to support continued development.
- 1 Mangara, the Diplomat | draw | Included as a draw card to replenish options.
- 1 Markov Purifier | draw | Included as a draw card for ongoing resources.
- 1 Sheoldred, the Apocalypse | draw | Included as a draw card to sustain the hand.
- 1 Sigarda's Splendor | draw | Included as a draw card for additional resources.
- 1 The Gaffer | draw | Included as a draw card to maintain momentum.
- 1 Well of Lost Dreams | draw | Included as a draw card that supports the deck's plan.
- 1 Alseid of Life's Bounty | interaction | Included as interaction to protect the deck's position.
- 1 Faith's Shield | interaction | Included as interaction for timely protection.
- 1 Metropolis Reformer | interaction | Included as interaction that supports the board.
- 1 Restoration Magic | interaction | Included as interaction for flexible responses.
- 1 Sword of Light and Shadow | interaction | Included as interaction that supports key creatures.
- 1 Werefox Bodyguard | interaction | Included as interaction to answer opposing pressure.
- 1 Altar of the Pantheon | ramp | Included as ramp to accelerate development.
- 1 Angel of Indemnity | ramp | Included as ramp to advance the mana plan.
- 1 Crypt Ghast | ramp | Included as ramp for larger turns.
- 1 Hierophant's Chalice | ramp | Included as ramp to support the deck's development.
- 1 Life Insurance | ramp | Included as ramp for additional mana resources.
- 1 Nuka-Cola Vending Machine | ramp | Included as ramp to accelerate the game plan.
- 1 Orazca Relic | ramp | Included as ramp for steady development.
- 1 Potioner's Trove | ramp | Included as ramp to enable bigger plays.
- 1 Pristine Talisman | ramp | Included as ramp that fits the deck's plan.
- 1 The Celestus | ramp | Included as ramp for mana acceleration.
- 1 Ayli, Eternal Pilgrim | removal | Included as removal for problematic permanents.
- 1 Murderous Rider // Swift End | removal | Included as removal for opposing threats.
- 1 Solitude | removal | Included as removal to clear key threats.
- 1 Tithing Blade // Consuming Sepulcher | removal | Included as removal for board control.
- 1 Umezawa's Jitte | removal | Included as removal that supports combat.
- 1 Vein Ripper | removal | Included as removal that also adds board presence.
- 1 Vona, Butcher of Magan | removal | Included as removal on a substantial threat.
- 1 Witch of the Moors | removal | Included as removal to pressure opposing boards.
- 1 Aerith Gainsborough | synergy | Included for lifegain-focused synergy.
- 1 Ajani's Pridemate | synergy | Included for lifegain-focused synergy.
- 1 Angel of Vitality | synergy | Included for lifegain-focused synergy.
- 1 Angelic Accord | synergy | Included for lifegain-focused synergy.
- 1 Blood Artist | synergy | Included for synergy with the deck's creature plan.
- 1 Cleric Class | synergy | Included for lifegain-focused synergy.
- 1 Cleric of Life's Bond | synergy | Included for lifegain-focused synergy.
- 1 Heliod, Sun-Crowned | synergy | Included for lifegain-focused synergy.
- 1 Indulging Patrician | synergy | Included for lifegain-focused synergy.
- 1 Leyline of Hope | synergy | Included for lifegain-focused synergy.
- 1 Resplendent Angel | synergy | Included for lifegain-focused synergy.
- 1 Righteous Valkyrie | synergy | Included for lifegain-focused synergy.
- 1 Sanguine Bond | synergy | Included for lifegain-focused synergy.
- 1 Vito, Thorn of the Dusk Rose | synergy | Included for lifegain-focused synergy.
- 1 Archangel of Thune | threat | Included as a powerful lifegain-focused threat.
- 1 Astarion, the Decadent | threat | Included as a threat that fits the deck's plan.
- 1 Attended Healer | threat | Included as a threat supporting the creature plan.
- 1 Celestine, the Living Saint | threat | Included as a substantial threat.
- 1 Cliffhaven Vampire | threat | Included as a lifegain-focused threat.
- 1 Defiant Bloodlord | threat | Included as a lifegain-focused threat.
- 1 Divinity of Pride | threat | Included as a substantial threat.
- 1 Exalted Sunborn | threat | Included as a lifegain-focused threat.
- 1 Felidar Sovereign | threat | Included as a major lifegain-focused threat.
- 1 Liesa, Forgotten Archangel | threat | Included as a high-impact threat.
- 1 Lyra Dawnbringer | threat | Included as a substantial threat.
- 1 Valkyrie Harbinger | threat | Included as a lifegain-focused threat.
- 1 Ajani, Strength of the Pride | wipe | Included as a board wipe when the table gets crowded.
- 1 Fumigate | wipe | Included as a board wipe to reset opposing boards.
- 1 Kaya's Wrath | wipe | Included as a board wipe for difficult board states.

</details>

### 2. aristocrats Commander, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 178 names.

Cards: 99. Repair turn: false. Block findings: 0.

**Summary:** This is a black-white aristocrats shell centered on Denethor, Ruling Steward, using a deep draw package, efficient mana development, and layered removal to keep playing through longer games. It aims to establish a board, protect its important pieces, and convert sustained pressure into a win rather than racing immediately. The tradeoff is a lighter dedicated threat package, so the deck leans on Denethor, its synergy cards, and careful use of its interaction to close games.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 2.79 over 62 nonland cards
- [INFO] `basics_added`: the list was 1 card short, so the builder added 1 basic land

<details><summary>The deck list</summary>

- 19 Plains | land | Provides reliable white mana.
- 10 Swamp | land | Provides reliable black mana.
- 1 Command Tower | land | Provides flexible commander-focused mana.
- 1 Castle Locthwain | land | Adds a black mana source.
- 1 Evolving Wilds | land | Helps stabilize the mana base.
- 1 Fabled Passage | land | Helps stabilize the mana base.
- 1 Marsh Flats | land | Finds a needed basic land.
- 1 Minas Tirith | land | Adds a white mana source.
- 1 Takenuma, Abandoned Mire | land | Adds a black mana source.
- 1 War Room | land | Rounds out the utility lands.
- 1 Arcane Signet | ramp | Efficient mana development for the commander’s colors.
- 1 Sol Ring | ramp | Accelerates early development.
- 1 Commander's Sphere | ramp | Provides flexible mana support.
- 1 Fellwar Stone | ramp | Provides inexpensive mana acceleration.
- 1 Thought Vessel | ramp | Adds colorless mana development.
- 1 Wayfarer's Bauble | ramp | Supports early mana development.
- 1 Relic of Legends | ramp | Provides additional mana support.
- 1 Lotho, Corrupt Shirriff | ramp | Adds a creature-based ramp piece.
- 1 Deadly Dispute | ramp | Contributes to the sacrifice-oriented mana plan.
- 1 Sword of the Animist | ramp | Provides equipment-based mana development.
- 1 Ahriman | draw | Adds a creature-based draw option.
- 1 Buster Sword | draw | Provides an equipment-based draw piece.
- 1 Call of the Ring | draw | Adds steady card access.
- 1 Cirith Ungol Patrol | draw | Provides a creature-based draw slot.
- 1 Exemplar of Light | draw | Adds a white creature draw option.
- 1 Foot Chopper | draw | Provides another artifact draw piece.
- 1 Idol of Oblivion | draw | Adds a compact draw artifact.
- 1 Inspiring Overseer | draw | Provides a creature-based draw option.
- 1 Lembas | draw | Adds an artifact draw piece.
- 1 Mask of Memory | draw | Provides equipment-based card access.
- 1 Massacre Girl, Known Killer | draw | Adds a black creature draw option.
- 1 Nasty End | draw | Supports the deck with additional card access.
- 1 Night's Whisper | draw | Provides efficient card access.
- 1 Puresteel Paladin | draw | Supports the deck’s equipment draw package.
- 1 Rat King, Pale Piper | draw | Adds another black creature draw piece.
- 1 Secret Rendezvous | draw | Provides a white draw option.
- 1 Skullclamp | draw | A key draw artifact for the aristocrats shell.
- 1 Stone of Erech | draw | Provides an additional artifact draw slot.
- 1 Tome of Legends | draw | Adds repeatable card access.
- 1 Vanquisher's Banner | draw | Provides a tribal-leaning draw artifact.
- 1 Wall of Omens | draw | Adds an early creature draw piece.
- 1 Angel of Serenity | removal | Provides creature-based removal.
- 1 Bitter Triumph | removal | Adds flexible black removal.
- 1 Claim the Precious | removal | Provides targeted black removal.
- 1 Crib Swap | removal | Adds an instant-speed removal option.
- 1 Fiend Hunter | removal | Provides creature-based removal.
- 1 Generous Gift | removal | Adds broad white removal.
- 1 Infernal Grasp | removal | Provides efficient black removal.
- 1 Swords to Plowshares | removal | Adds efficient white removal.
- 1 Dismember | removal | Provides another low-cost removal option.
- 1 Dispatch | removal | Adds an artifact-friendly removal spell.
- 1 Fatal Push | removal | Provides efficient targeted removal.
- 1 Heartless Act | removal | Adds another instant-speed black answer.
- 1 Bastion Protector | interaction | Helps protect the deck’s central creature.
- 1 Boromir, Warden of the Tower | interaction | Provides a creature-based interaction piece.
- 1 Clever Concealment | interaction | Adds protection for the developed board.
- 1 Gift of Immortality | interaction | Supports a resilient creature plan.
- 1 Lightning Greaves | interaction | Provides efficient protective equipment.
- 1 Swiftfoot Boots | interaction | Adds a second protective equipment piece.
- 1 Take Up the Shield | interaction | Provides a compact protection spell.
- 1 Unbreakable Formation | interaction | Protects the board at a key moment.
- 1 Austere Command | wipe | Provides a flexible board reset.
- 1 Fumigate | wipe | Adds a reliable creature-focused reset.
- 1 Vanquish the Horde | wipe | Provides another board-clearing option.
- 1 Arcade Cabinet | synergy | Adds artifact-centered aristocrats synergy.
- 1 Gollum, Patient Plotter | synergy | Provides a black synergy creature for the core plan.
- 1 Gríma Wormtongue | synergy | Adds a black synergy piece.
- 1 Heirloom Auntie | synergy | Provides another creature-based synergy card.
- 1 Nimble Hobbit | synergy | Adds a low-profile synergy creature.
- 1 Phantom Train | synergy | Provides an artifact synergy piece.
- 1 Bill the Pony | threat | Adds a white creature threat.
- 1 Vengeful Villagers | threat | Provides an additional creature threat.

</details>

### 3. artifacts Commander, bracket 4

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 299 names.

Cards: 99. Repair turn: false. Block findings: 0.

**Summary:** Urza leads a high-power artifact deck that develops an artifact-heavy board, supports it with focused mana and card flow, and presses the table with large artifact threats backed by disruption and reset tools. It wins by establishing that board and converting its artifact pressure into a decisive advantage. The tradeoff is a tightly artifact-centered build with relatively little room for cards outside that plan.

- [INFO] `curve_summary`: average mana value 3.89 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Academy Ruins | land | Utility land slot in the artifact-focused mana base.
- 1 Archway of Innovation | land | Land slot in the artifact-focused mana base.
- 1 Blinkmoth Nexus | land | Land slot in the artifact-focused mana base.
- 1 Buried Ruin | land | Utility land slot in the artifact-focused mana base.
- 1 Conqueror's Galleon // Conqueror's Foothold | land | Land slot in the artifact-focused mana base.
- 1 Fomori Vault | land | Land slot in the artifact-focused mana base.
- 1 Glimmervoid | land | Land slot in the artifact-focused mana base.
- 1 Hall of Tagsin | land | Land slot in the artifact-focused mana base.
- 1 Inkmoth Nexus | land | Land slot in the artifact-focused mana base.
- 1 Inventors' Fair | land | Utility land slot in the artifact-focused mana base.
- 1 Mirrex | land | Land slot in the artifact-focused mana base.
- 1 Mishra's Factory | land | Land slot in the artifact-focused mana base.
- 1 Mishra's Foundry | land | Land slot in the artifact-focused mana base.
- 1 Mishra's Workshop | land | Land slot in the artifact-focused mana base.
- 1 Otawara, Soaring City | land | Utility land slot in the artifact-focused mana base.
- 1 Power Depot | land | Artifact land slot in the mana base.
- 1 Roadside Reliquary | land | Land slot in the artifact-focused mana base.
- 1 Spire of Industry | land | Land slot in the artifact-focused mana base.
- 1 Thaumatic Compass // Spires of Orazca | land | Land slot in the artifact-focused mana base.
- 1 The Everflowing Well // The Myriad Pools | land | Land slot in the artifact-focused mana base.
- 1 The Mycosynth Gardens | land | Land slot in the artifact-focused mana base.
- 1 Treasure Map // Treasure Cove | land | Land slot in the artifact-focused mana base.
- 1 Urza's Factory | land | Land slot in the artifact-focused mana base.
- 1 Urza's Saga | land | Land slot in the artifact-focused mana base.
- 1 Urza's Workshop | land | Land slot in the artifact-focused mana base.
- 1 Uthros, Titanic Godcore | land | Land slot in the artifact-focused mana base.
- 10 Island | land | Basic land foundation for the mana base.
- 1 Mox Opal | ramp | Artifact-focused ramp piece.
- 1 Metalworker | ramp | Artifact-focused ramp piece.
- 1 Krark-Clan Ironworks | ramp | Artifact-focused ramp piece.
- 1 Moonsnare Prototype | ramp | Artifact-focused ramp piece.
- 1 Chief Engineer | ramp | Artifact-focused ramp piece.
- 1 Grand Architect | ramp | Artifact-focused ramp piece.
- 1 Karn, Legacy Reforged | ramp | Artifact-focused ramp piece.
- 1 Tezzeret the Seeker | ramp | Artifact-focused ramp piece.
- 1 Inspiring Statuary | ramp | Artifact-focused ramp piece.
- 1 Solar Array | ramp | Artifact-focused ramp piece.
- 1 Thoughtcast | draw | Artifact-focused draw selection.
- 1 Thought Monitor | draw | Artifact-focused draw selection.
- 1 Thirst for Knowledge | draw | Draw selection for the artifact deck.
- 1 Sai, Master Thopterist | draw | Artifact-focused draw selection.
- 1 Vedalken Archmage | draw | Artifact-focused draw selection.
- 1 Reverse Engineer | draw | Artifact-focused draw selection.
- 1 Tezzeret, Artifice Master | draw | Artifact-focused draw selection.
- 1 Tezzeret, Betrayer of Flesh | draw | Artifact-focused draw selection.
- 1 Forensic Gadgeteer | draw | Artifact-focused draw selection.
- 1 Riddlesmith | draw | Artifact-focused draw selection.
- 1 Metallic Rebuke | interaction | Interaction slot for protecting the deck's plan.
- 1 Stoic Rebuttal | interaction | Interaction slot for protecting the deck's plan.
- 1 Disruption Protocol | interaction | Interaction slot for protecting the deck's plan.
- 1 Ice Out | interaction | Interaction slot for protecting the deck's plan.
- 1 Jin-Gitaxias, Progress Tyrant | interaction | Interaction piece with a high-impact artifact-deck profile.
- 1 Welding Jar | interaction | Interaction slot for protecting the artifact plan.
- 1 Aether Spellbomb | removal | Artifact-based removal selection.
- 1 Arcum Dagsson | removal | Artifact-focused removal selection.
- 1 Canoptek Scarab Swarm | removal | Artifact-based removal selection.
- 1 Contagion Clasp | removal | Artifact-based removal selection.
- 1 Portal to Phyrexia | removal | High-impact artifact removal selection.
- 1 Resculpt | removal | Removal selection for the deck.
- 1 Skysovereign, Consul Flagship | removal | Artifact-based removal selection.
- 1 Spine of Ish Sah | removal | Artifact-based removal selection.
- 1 Mystic Forge | synergy | Central artifact synergy piece.
- 1 Transmute Artifact | synergy | Artifact synergy selection.
- 1 Whir of Invention | synergy | Artifact synergy selection.
- 1 Unwinding Clock | synergy | Artifact synergy selection.
- 1 Voltaic Key | synergy | Artifact synergy selection.
- 1 Clock of Omens | synergy | Artifact synergy selection.
- 1 Power Artifact | synergy | Artifact synergy selection.
- 1 Manifold Key | synergy | Artifact synergy selection.
- 1 Etherium Sculptor | synergy | Artifact synergy selection.
- 1 Foundry Inspector | synergy | Artifact synergy selection.
- 1 Emry, Lurker of the Loch | synergy | Artifact synergy selection.
- 1 Scrap Trawler | synergy | Artifact synergy selection.
- 1 Shimmer Myr | synergy | Artifact synergy selection.
- 1 Panharmonicon | synergy | Artifact synergy selection.
- 1 Kappa Cannoneer | threat | Artifact-focused threat.
- 1 Kuldotha Forgemaster | threat | Artifact-focused threat.
- 1 Master Transmuter | threat | Artifact-focused threat.
- 1 Cyberdrive Awakener | threat | Artifact-focused threat.
- 1 Metalwork Colossus | threat | Artifact-focused threat.
- 1 Mycosynth Golem | threat | Artifact-focused threat.
- 1 The Capitoline Triad | threat | Artifact-focused threat.
- 1 Traxos, Scourge of Kroog | threat | Artifact-focused threat.
- 1 Phyrexian Metamorph | threat | Artifact-focused threat.
- 1 Karn, Scion of Urza | threat | Artifact-focused threat.
- 1 Threefold Thunderhulk | threat | Artifact-focused threat.
- 1 Darksteel Juggernaut | threat | Artifact-focused threat.
- 1 Nevinyrral's Disk | wipe | Board-wipe selection for resetting contested boards.
- 1 Oblivion Stone | wipe | Board-wipe selection for resetting contested boards.
- 1 Engineered Explosives | wipe | Board-wipe selection for resetting contested boards.

</details>

### 4. dinosaur tribal, bracket 2

Format: Commander. Theme: dinosaurs. Pool: any_card. Shortlist: 267 names.

Cards: 99. Repair turn: false. Block findings: 0.

**Summary:** This deck develops its mana, fills the battlefield with Dinosaur creatures, and keeps cards flowing while it builds toward its larger threats. It aims to win by attacking with a wide and imposing Dinosaur board, with protection, removal, and board wipes available to keep that plan moving. It gives up speed and relies on establishing its mana and creature board before its biggest plays take over the table.

- [INFO] `bracket_prose_rules`: bracket 2 (Core, 8+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 4.00 over 63 nonland cards

<details><summary>The deck list</summary>

- 10 Forest | land | Basic Forests form the core of the green land base.
- 6 Mountain | land | Basic Mountains support the red land base.
- 5 Plains | land | Basic Plains support the white land base.
- 1 Command Tower | land | A multicolor land for the land base.
- 1 Path of Ancestry | land | A tribal-focused land for the Dinosaur deck.
- 1 Canopy Vista | land | A green-white land for the land base.
- 1 Cinder Glade | land | A red-green land for the land base.
- 1 Rootbound Crag | land | A red-green land for the land base.
- 1 Sunpetal Grove | land | A green-white land for the land base.
- 1 Clifftop Retreat | land | A red-white land for the land base.
- 1 Temple Garden | land | A green-white land for the land base.
- 1 Sacred Foundry | land | A red-white land for the land base.
- 1 Stomping Ground | land | A red-green land for the land base.
- 1 Evolving Wilds | land | A flexible land for the land base.
- 1 Terramorphic Expanse | land | A flexible land for the land base.
- 1 Myriad Landscape | land | A land-based mana option.
- 1 Exotic Orchard | land | A multicolor land for the land base.
- 1 Rogue's Passage | land | A utility land for the land base.
- 1 Arcane Signet | ramp | A straightforward multicolor ramp piece.
- 1 Sol Ring | ramp | An efficient artifact ramp piece.
- 1 Cultivate | ramp | A reliable land-focused ramp spell.
- 1 Kodama's Reach | ramp | A reliable land-focused ramp spell.
- 1 Nature's Lore | ramp | A green ramp spell for early development.
- 1 Farseek | ramp | A multicolor land-focused ramp spell.
- 1 Thunderherd Migration | ramp | A Dinosaur-themed ramp spell.
- 1 Drover of the Mighty | ramp | A creature ramp piece that fits the Dinosaur deck.
- 1 Ranging Raptors | ramp | A Dinosaur ramp creature.
- 1 Topiary Stomper | ramp | A Dinosaur ramp creature.
- 1 Beast Whisperer | draw | A creature-based draw piece.
- 1 Garruk's Uprising | draw | A draw enchantment for the large-creature plan.
- 1 Guardian Project | draw | A creature-focused draw enchantment.
- 1 Harmonize | draw | A simple dedicated draw spell.
- 1 Ripjaw Raptor | draw | A Dinosaur draw creature.
- 1 Runic Armasaur | draw | A Dinosaur draw creature.
- 1 Shamanic Revelation | draw | A board-focused draw spell.
- 1 Titanoth Rex | draw | A Dinosaur draw card that fits the theme.
- 1 Vanquisher's Banner | draw | A tribal draw artifact for Dinosaurs.
- 1 Yidaro, Wandering Monster | draw | A Dinosaur draw card that also fits the creature plan.
- 1 Akroma's Will | interaction | A flexible protective interaction spell.
- 1 Boros Charm | interaction | A flexible protective interaction spell.
- 1 Heroic Intervention | interaction | A broad protection spell for the creature board.
- 1 Lightning Greaves | interaction | A protection equipment for key creatures.
- 1 Swiftfoot Boots | interaction | A protection equipment for key creatures.
- 1 Flawless Maneuver | interaction | A protection spell for the creature board.
- 1 Bronzebeak Foragers | removal | A Dinosaur removal creature.
- 1 Itzquinth, Firstborn of Gishath | removal | A Dinosaur-themed removal card.
- 1 Kogla and Yidaro | removal | A Dinosaur-themed removal threat.
- 1 Ravenous Sailback | removal | A Dinosaur removal creature.
- 1 Savage Stomp | removal | A Dinosaur-themed removal spell.
- 1 Thrashing Brontodon | removal | A Dinosaur removal creature.
- 1 Tranquil Frillback | removal | A Dinosaur removal creature.
- 1 Trumpeting Carnosaur | removal | A Dinosaur removal creature.
- 1 Austere Command | wipe | A flexible board wipe.
- 1 Blasphemous Act | wipe | A dedicated board wipe.
- 1 Wakening Sun's Avatar | wipe | A Dinosaur-themed board wipe.
- 1 Armored Kincaller | synergy | A Dinosaur synergy creature.
- 1 Belligerent Yearling | synergy | A Dinosaur synergy creature.
- 1 Commune with Dinosaurs | synergy | A Dinosaur-focused synergy spell.
- 1 Dinosaur Stampede | synergy | A Dinosaur-focused synergy spell.
- 1 Huatli's Raptor | synergy | A Dinosaur synergy creature.
- 1 Kinjalli's Caller | synergy | A Dinosaur support creature.
- 1 Kinjalli's Sunwing | synergy | A Dinosaur synergy creature.
- 1 Marauding Raptor | synergy | A Dinosaur synergy creature.
- 1 Otepec Huntmaster | synergy | A Dinosaur support creature.
- 1 Raptor Companion | synergy | An early Dinosaur for the creature theme.
- 1 Raptor Hatchling | synergy | An early Dinosaur for the creature theme.
- 1 Regal Imperiosaur | synergy | A Dinosaur synergy creature.
- 1 Sunfrill Imitator | synergy | A Dinosaur synergy creature.
- 1 Territorial Hammerskull | synergy | A Dinosaur synergy creature.
- 1 Ancient Brontodon | threat | A large Dinosaur threat.
- 1 Ancient Imperiosaur | threat | A large Dinosaur threat.
- 1 Annoyed Altisaur | threat | A Dinosaur threat for the top end.
- 1 Bellowing Aegisaur | threat | A Dinosaur threat for the creature plan.
- 1 Carnage Tyrant | threat | A powerful Dinosaur threat.
- 1 Colossal Dreadmaw | threat | A straightforward Dinosaur threat for a new player.
- 1 Goring Ceratops | threat | A Dinosaur threat for combat-focused games.
- 1 Palani's Hatcher | threat | A Dinosaur threat that supports the board plan.
- 1 Regisaur Alpha | threat | A Dinosaur threat for the tribal plan.
- 1 Rampaging Brontodon | threat | A large Dinosaur threat.
- 1 Thundering Spineback | threat | A Dinosaur threat for the top end.
- 1 Zetalpa, Primal Dawn | threat | A large Dinosaur threat for closing games.

</details>

### 5. blink Commander, owned first

Format: Commander. Theme: blink. Pool: owned_first. Shortlist: 168 names.

Cards: 99. Repair turn: false. Block findings: 0.

**Summary:** Gilraen leads a creature-forward blink deck that develops mana, deploys creatures, and leans on its blink synergy cards to keep advancing its board. It wins by converting that sustained development into pressure from its threat suite while retaining removal and protective interaction for pivotal turns. The deck gives up explosive finishing speed for a more methodical, permanent-based game plan that benefits from keeping key creatures in play.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 2.98 over 63 nonland cards

<details><summary>The deck list</summary>

- 26 Plains | land | Provides the core white mana base.
- 1 Abandoned Air Temple | land | Adds a land slot to the mana base.
- 1 Command Tower | land | Adds a land slot to the mana base.
- 1 Evolving Wilds | land | Adds a land slot to the mana base.
- 1 Fabled Passage | land | Adds a land slot to the mana base.
- 1 Minas Tirith | land | Adds a land slot to the mana base.
- 1 Path of Ancestry | land | Adds a land slot to the mana base.
- 1 Plaza of Heroes | land | Adds a land slot to the mana base.
- 1 Reliquary Tower | land | Adds a land slot to the mana base.
- 1 Rogue's Passage | land | Adds a land slot to the mana base.
- 1 War Room | land | Adds a land slot to the mana base.
- 1 Sol Ring | ramp | Provides ramp for the deck.
- 1 Arcane Signet | ramp | Provides ramp for the deck.
- 1 Bender's Waterskin | ramp | Provides ramp for the deck.
- 1 Commander's Sphere | ramp | Provides ramp for the deck.
- 1 Fellwar Stone | ramp | Provides ramp for the deck.
- 1 Relic of Legends | ramp | Provides ramp for the deck.
- 1 Sword of the Animist | ramp | Provides ramp for the deck.
- 1 Thought Vessel | ramp | Provides ramp for the deck.
- 1 Wayfarer's Bauble | ramp | Provides ramp for the deck.
- 1 White Lotus Tile | ramp | Provides ramp for the deck.
- 1 Adventurer's Airship | draw | Provides card draw.
- 1 Champions of Minas Tirith | draw | Provides card draw.
- 1 Diary of Dreams | draw | Provides card draw.
- 1 Energybending | draw | Provides card draw.
- 1 Faramir, Field Commander | draw | Provides card draw.
- 1 Idol of Oblivion | draw | Provides card draw.
- 1 Lembas | draw | Provides card draw.
- 1 Mask of Memory | draw | Provides card draw.
- 1 Skullclamp | draw | Provides card draw.
- 1 Tome of Legends | draw | Provides card draw.
- 1 Champion's Helm | interaction | Protects key pieces of the plan.
- 1 Clever Concealment | interaction | Provides protective interaction.
- 1 Lightning Greaves | interaction | Protects key pieces of the plan.
- 1 Swiftfoot Boots | interaction | Protects key pieces of the plan.
- 1 Together Forever | interaction | Provides protective interaction.
- 1 Unbreakable Formation | interaction | Provides protective interaction.
- 1 Banishing Light | removal | Provides removal coverage.
- 1 Generous Gift | removal | Provides removal coverage.
- 1 Get Lost | removal | Provides removal coverage.
- 1 Journey to Nowhere | removal | Provides removal coverage.
- 1 March of Otherworldly Light | removal | Provides removal coverage.
- 1 Oblation | removal | Provides removal coverage.
- 1 Swords to Plowshares | removal | Provides removal coverage.
- 1 Stroke of Midnight | removal | Provides removal coverage.
- 1 Austere Command | wipe | Provides a reset option when the board gets away from the deck.
- 1 Fumigate | wipe | Provides a reset option when the board gets away from the deck.
- 1 Vanquish the Horde | wipe | Provides a reset option when the board gets away from the deck.
- 1 Aang, the Last Airbender | synergy | Creature slot chosen to deepen the blink plan.
- 1 Angel of Condemnation | synergy | Creature slot chosen to deepen the blink plan.
- 1 Angel of Serenity | synergy | Creature slot chosen to deepen the blink plan.
- 1 Ennis, Debate Moderator | synergy | Dedicated synergy piece for the blink plan.
- 1 Fiend Hunter | synergy | Creature slot chosen to deepen the blink plan.
- 1 Flickerwisp | synergy | Dedicated synergy piece for the blink plan.
- 1 Gift of Immortality | synergy | Support piece for the blink plan.
- 1 Inspiring Overseer | synergy | Creature slot chosen to deepen the blink plan.
- 1 Jocasta, Automaton Avenger | synergy | Dedicated synergy piece for the blink plan.
- 1 Palace Jailer | synergy | Creature slot chosen to deepen the blink plan.
- 1 Personify | synergy | Dedicated synergy piece for the blink plan.
- 1 Slip On the Ring | synergy | Spell slot chosen to deepen the blink plan.
- 1 Wall of Omens | synergy | Creature slot chosen to deepen the blink plan.
- 1 Westfold Rider | synergy | Creature slot chosen to deepen the blink plan.
- 1 Angel of Sanctions | threat | Creature-based threat for the deck's finishing plan.
- 1 Bastion Protector | threat | Creature-based threat for the deck's finishing plan.
- 1 Boromir, Warden of the Tower | threat | Creature-based threat for the deck's finishing plan.
- 1 Bronze Guardian | threat | Artifact-creature threat for the deck's finishing plan.
- 1 Exemplar of Light | threat | Creature-based threat for the deck's finishing plan.
- 1 Frontline Medic | threat | Creature-based threat for the deck's finishing plan.
- 1 Giada, Font of Hope | threat | Creature-based threat for the deck's finishing plan.
- 1 Kataki, War's Wage | threat | Creature-based threat for the deck's finishing plan.
- 1 Puresteel Paladin | threat | Creature-based threat for the deck's finishing plan.
- 1 The Vision | threat | Artifact-creature threat for the deck's finishing plan.
- 1 The Walls of Ba Sing Se | threat | Artifact-creature threat for the deck's finishing plan.
- 1 Zack Fair | threat | Creature-based threat for the deck's finishing plan.

</details>

### 6. Modern tempo, tournament

Format: Modern. Theme: tempo. Pool: any_card. Shortlist: 167 names.

Cards: 60. Repair turn: false. Block findings: 0.

**Summary:** This blue-red tempo deck establishes pressure with early creature threats, then keeps that pressure intact with draw, interaction, and removal. It wins by turning a small lead on the board into steady combat damage while the temporal package supports a proactive game plan. It gives up broader permanent coverage and larger standalone threats in exchange for a focused, spell-heavy approach that rewards careful timing.

- [INFO] `curve_summary`: average mana value 2.17 over 36 nonland cards

<details><summary>The deck list</summary>

- 7 Island | land | Blue basic land that supports the deck's primary color.
- 5 Mountain | land | Red basic land for removal and threat deployment.
- 4 Steam Vents | land | Blue-red dual land for consistent access to both colors.
- 4 Scalding Tarn | land | Land that supports the blue-red mana base.
- 2 Shivan Reef | land | Blue-red land that helps cast the deck's spells on time.
- 1 Stormcarved Coast | land | Blue-red land for the tempo mana base.
- 1 Otawara, Soaring City | land | Blue land slot with additional utility.
- 4 Ragavan, Nimble Pilferer | threat | Early creature threat that supports an aggressive tempo start.
- 4 Ledger Shredder | threat | Creature threat that fits the deck's spell-heavy plan.
- 2 Faerie Mastermind | threat | Evasive creature threat for maintaining pressure.
- 2 Enduring Curiosity | threat | Creature threat that adds to the deck's board presence.
- 4 Consider | draw | Efficient draw spell that helps keep the deck moving.
- 2 Preordain | draw | Draw spell that improves access to key threats and answers.
- 4 Counterspell | interaction | Core interaction for protecting pressure and stopping opposing plans.
- 2 Spell Pierce | interaction | Low-cost interaction suited to a tempo strategy.
- 4 Lightning Bolt | removal | Flexible removal that clears opposing creatures.
- 2 Into the Flood Maw | removal | Removal option that helps preserve the attack.
- 2 Untimely Malfunction | removal | Additional removal for creatures that interfere with pressure.
- 4 Temporal Mastery | synergy | Synergy card for the deck's temporal theme and proactive plan.

</details>

### 7. Modern burn, casual

Format: Modern. Theme: burn. Pool: any_card. Shortlist: 290 names.

Cards: 60. Repair turn: false. Block findings: 0.

**Summary:** This mono-red burn deck opens with removal and burn pressure, then uses spell-focused synergy creatures and a steady stream of threats to keep the opponent under strain. It wins by combining damage-oriented spells with creature attacks, while draw helps maintain momentum into later turns. The deck gives up broad defensive flexibility for a direct, proactive plan that is strongest when it can keep pressing forward.

- [INFO] `curve_summary`: average mana value 2.94 over 36 nonland cards

<details><summary>The deck list</summary>

- 24 Mountain | land | Provides the mono-red mana base.
- 4 Risk Factor | draw | Supplies draw to keep the burn plan stocked.
- 2 Sazacap's Brew | draw | Adds further draw for the burn-focused plan.
- 2 Chandra, Dressed to Kill | ramp | Fills the ramp role while staying in the deck's red core.
- 4 Lightning Bolt | removal | Primary efficient removal for clearing resistance to the attack.
- 2 Lightning Strike | removal | Additional removal that supports the burn-focused strategy.
- 4 Guttersnipe | synergy | A synergy creature for the deck's spell-heavy game plan.
- 4 Thermo-Alchemist | synergy | Provides another synergy piece alongside the burn spells.
- 4 Hazoret the Fervent | threat | A resilient threat that gives the deck creature pressure.
- 4 Ashcloud Phoenix | threat | Provides a creature threat for sustained pressure.
- 4 Tectonic Giant | threat | Adds a larger threat to pressure opponents after the early burn.
- 2 Torbran, Thane of Red Fell | threat | Rounds out the threat suite with a red creature payoff.

</details>

### 8. Modern lifegain, FNM

Format: Modern. Theme: lifegain. Pool: any_card. Shortlist: 301 names.

Cards: 60. Repair turn: false. Block findings: 0.

**Summary:** This white-black lifegain deck establishes its board with early synergy pieces, uses its life total to support card flow and creature development, and presses the opponent with a steady mix of Vampire, Angel, and other creature threats. It wins by turning that lifegain-focused board into sustained combat pressure while removal clears the way. The tradeoff is that the deck is built around assembling complementary permanents, so it can be less explosive when its synergies are disrupted or its larger threats arrive late.

- [INFO] `curve_summary`: average mana value 3.17 over 36 nonland cards

<details><summary>The deck list</summary>

- 10 Plains | land | Provides white mana for the deck’s white cards.
- 8 Swamp | land | Provides black mana for the deck’s black cards.
- 4 Scoured Barrens | land | Supports the white-black mana base and the lifegain theme.
- 2 Shambling Vent | land | Adds another white-black land option.
- 2 Inspiring Overseer | draw | Provides a creature-based draw option.
- 2 Dawn of Hope | draw | Adds repeatable draw support for the lifegain plan.
- 2 Well of Lost Dreams | draw | Uses the deck’s lifegain focus as a draw engine.
- 2 Orzhov Keyrune | ramp | Provides ramp in the deck’s colors.
- 2 Murderous Rider // Swift End | removal | Supplies flexible creature-based removal.
- 2 Nightmare's Thirst | removal | Provides inexpensive black removal.
- 2 Tithing Blade // Consuming Sepulcher | removal | Adds another removal piece with a lasting presence.
- 4 Soul Warden | synergy | A core lifegain synergy card for the early game.
- 4 Ajani's Pridemate | synergy | Rewards the deck for pursuing its lifegain plan.
- 2 Attended Healer | threat | Provides a lifegain-focused creature threat.
- 2 Bloodbond Vampire | threat | Gives the deck a black lifegain-themed threat.
- 2 Blood Baron of Vizkopa | threat | Adds a durable Vampire threat to the deck.
- 2 Nykthos Paragon | threat | Provides a white threat that fits the lifegain strategy.
- 2 Twinblade Paladin | threat | Adds another creature threat for the lifegain plan.
- 2 Valkyrie Harbinger | threat | Serves as a high-impact Angel threat.
- 2 Rhox Faithmender | threat | Provides a substantial threat for the lifegain shell.

</details>

### 9. Standard midrange, FNM

Format: Standard. Theme: midrange. Pool: any_card. Shortlist: 164 names.

Cards: 60. Repair turn: false. Block findings: 0.

**Summary:** This black-green midrange deck develops its mana, establishes a board of substantial creature threats, and supports that pressure with removal, protection, and steady card flow. It wins by keeping a threat on the table long enough to take over the game while answering opposing plans. In exchange, it is less focused on racing quickly and leans on its board presence to carry the matchup.

- [INFO] `curve_summary`: average mana value 3.08 over 36 nonland cards

<details><summary>The deck list</summary>

- 8 Forest | land | Basic land slot supporting the deck's mana base.
- 8 Swamp | land | Basic land slot supporting the deck's mana base.
- 4 Overgrown Tomb | land | Dual land slot supporting the deck's mana base.
- 2 Blooming Marsh | land | Land slot supporting the deck's mana base.
- 2 Underground Mortuary | land | Land slot supporting the deck's mana base.
- 2 Llanowar Elves | ramp | Early ramp slot for the midrange plan.
- 2 Bitter Triumph | removal | Efficient removal slot for clearing opposition.
- 2 Assassin's Trophy | removal | Flexible removal slot for answering problems.
- 2 Nowhere to Run | removal | Removal slot that supports the deck's answer suite.
- 2 Darkstar Augur | draw | Draw source that helps sustain resources.
- 2 Phyrexian Arena | draw | Draw source for keeping the midrange plan supplied.
- 2 Unholy Annex // Ritual Chamber | draw | Draw source that supports longer games.
- 2 Snakeskin Veil | synergy | Synergy piece supporting the creature-focused plan.
- 2 Royal Treatment | synergy | Synergy piece that helps keep the board plan together.
- 2 Not Dead After All | synergy | Synergy piece supporting the deck's key creatures.
- 2 Undying Malice | synergy | Synergy piece supporting the creature-based game plan.
- 4 Goldvein Hydra | threat | Creature threat for applying pressure in a midrange game.
- 3 Vein Ripper | threat | Creature threat that strengthens the board-focused plan.
- 3 Vaultborn Tyrant | threat | Large creature threat for closing longer games.
- 2 Lumra, Bellow of the Woods | threat | Creature threat for the deck's top end.
- 2 Sephiroth, Fabled SOLDIER // Sephiroth, One-Winged Angel | threat | Creature threat that provides another powerful finisher.

</details>

### 10. Standard aggro, tournament

Format: Standard. Theme: aggro. Pool: any_card. Shortlist: 294 names.

Cards: 60. Repair turn: false. Block findings: 0.

**Summary:** This red-white aggro deck aims to establish an early board of threats, reinforce that pressure with its synergy pieces, and clear away resistance with removal and interaction. It wins by keeping attackers on the table and converting sustained pressure into combat damage. In exchange for that focused, proactive approach, it has less room for slower value plans and relies on maintaining momentum against opponents that can repeatedly answer its board.

- [INFO] `curve_summary`: average mana value 2.72 over 36 nonland cards

<details><summary>The deck list</summary>

- 8 Mountain | land | Provides a basic land slot for the red half of the deck.
- 8 Plains | land | Provides a basic land slot for the white half of the deck.
- 4 Sacred Foundry | land | Provides a red-white land slot.
- 4 Inspiring Vantage | land | Provides a red-white land slot.
- 4 Fugitive Codebreaker | draw | Supplies card access while fitting the aggressive creature plan.
- 2 Reckless Lackey | draw | Adds card access in a proactive deck slot.
- 4 Boros Charm | interaction | Provides flexible opposing-play disruption.
- 2 Sheltered by Ghosts | interaction | Adds interaction while remaining part of the proactive plan.
- 4 Harsh Annotation | removal | Provides efficient answers to opposing cards.
- 4 Case of the Gateway Express | removal | Fills a dedicated removal slot.
- 4 Warleader's Call | synergy | Supports the deck's aggressive creature-focused game plan.
- 4 Dragonback Lancer | threat | Forms a core set of aggressive threats.
- 4 Redcap Gutter-Dweller | threat | Adds more pressure through a full set of threats.
- 4 Frilled Sparkshooter | threat | Rounds out the creature pressure package.

</details>

### 11. Commander with a locked card

Format: Commander. Theme: sacrifice. Pool: any_card. Shortlist: 296 names.

Cards: 99. Repair turn: true. Block findings: 0.

**Summary:** Karlov sacrifice deck with 99 listed cards. Sol Ring remains included. Skullport Merchant now appears once; Warren Soultrader fills the replaced ramp slot while preserving the sacrifice-and-Treasure plan.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 3.16 over 63 nonland cards

<details><summary>The deck list</summary>

- 18 Plains | land | White mana base.
- 18 Swamp | land | Black mana base.
- 1 Sol Ring | ramp | Fast mana retained as requested.
- 1 Ashnod's Altar | ramp | Sacrifice outlet that converts fodder into mana.
- 1 Phyrexian Altar | ramp | Sacrifice outlet and color-fixing mana engine.
- 1 Pitiless Plunderer | ramp | Turns creature deaths into Treasure.
- 1 Priest of Forgotten Gods | ramp | Sacrifice outlet with mana production.
- 1 Pawn of Ulamog | ramp | Creates mana-producing tokens from deaths.
- 1 Sifter of Skulls | ramp | Produces Eldrazi Scions from nontoken deaths.
- 1 Culling the Weak | ramp | Efficient burst mana from expendable creatures.
- 1 Skullport Merchant | ramp | Treasure-producing sacrifice outlet.
- 1 Warren Soultrader | ramp | Replacement ramp piece that converts life and fodder into Treasure.
- 1 Corrupted Conviction | draw | Low-cost sacrifice-based card advantage.
- 1 Disciple of Bolas | draw | Converts a large creature into cards and life.
- 1 Smothering Abomination | draw | Sustained cards from the sacrifice plan.
- 1 Vampiric Rites | draw | Repeatable sacrifice outlet that draws cards.
- 1 Village Rites | draw | Efficient response to a creature dying.
- 1 Baron Bertram Graywater | draw | Vampire-based card advantage and token support.
- 1 Lord Skitter's Butcher | draw | Sacrifice-enabled card advantage.
- 1 Thraxodemon | draw | Turns disposable creatures into cards.
- 1 Skyclave Shadowcat | draw | Sacrifice outlet with repeatable card draw.
- 1 Tevesh Szat, Doom of Fools | draw | Creates fodder and converts it into cards.
- 1 Cartel Aristocrat | interaction | Reliable free sacrifice outlet with self-protection.
- 1 Dark Privilege | interaction | Protects key creatures through sacrifice.
- 1 Fanatical Devotion | interaction | Protects the board while using fodder.
- 1 Flare of Fortitude | interaction | Protects the board from major disruption.
- 1 Gift of Doom | interaction | Protects Karlov or a key engine piece.
- 1 Spirit Bonds | interaction | Creates Spirit fodder and protects creatures.
- 1 Attrition | removal | Sacrifice fodder to repeatedly answer creatures.
- 1 Ayli, Eternal Pilgrim | removal | Sacrifice outlet with flexible permanent removal.
- 1 Blasting Station | removal | Sacrifice outlet that picks off creatures and players.
- 1 Bone Shards | removal | Cheap removal that rewards sacrificing a creature.
- 1 Dictate of Erebos | removal | Turns each sacrificed creature into an opposing sacrifice.
- 1 Eaten Alive | removal | Exiles a problematic creature or planeswalker.
- 1 Grave Pact | removal | Punishes opposing boards for every creature lost.
- 1 Yawgmoth, Thran Physician | removal | Sacrifice outlet, card filtering, and creature control.
- 1 Altar of Dementia | synergy | Free sacrifice outlet that can convert a large board into a win.
- 1 Bartolomé del Presidio | synergy | Efficient free sacrifice outlet that grows into a threat.
- 1 Bastion of Remembrance | synergy | Death-trigger payoff plus a creature token.
- 1 Bloodflow Connoisseur | synergy | Free sacrifice outlet that grows steadily.
- 1 Carrion Feeder | synergy | Low-cost free sacrifice outlet.
- 1 Elas il-Kor, Sadistic Pilgrim | synergy | Drains opponents and gains life from creature deaths.
- 1 Fleshtaker | synergy | Rewards creature sacrifices with life drain and growth.
- 1 Hidden Stockpile | synergy | Provides recurring sacrifice fodder.
- 1 Nantuko Husk | synergy | Simple repeatable sacrifice outlet.
- 1 Open the Graves | synergy | Replaces dying nontoken creatures with Zombie fodder.
- 1 Viscera Seer | synergy | Free sacrifice outlet that improves draw quality.
- 1 Witch's Oven | synergy | Repeatedly converts creatures into resource tokens.
- 1 Woe Strider | synergy | Sacrifice outlet with built-in fodder and recursion.
- 1 Zulaport Cutthroat | synergy | Primary death-trigger life-drain payoff.
- 1 Abhorrent Overlord | threat | Large evasive finisher that supplies a token board.
- 1 Basri's Lieutenant | threat | Builds a board and leaves replacement threats behind.
- 1 Blood Host | threat | Sacrifice-fueled attacker that gains life.
- 1 Demon of Catastrophes | threat | Efficient flying finisher for expendable fodder.
- 1 Demonlord of Ashmouth | threat | Evasive threat that fits the sacrifice plan.
- 1 Felisa, Fang of Silverquill | threat | Turns Karlov's counters into a flying token army.
- 1 Ghoulcaller Gisa | threat | Converts creatures into a major Zombie board.
- 1 Liesa, Forgotten Archangel | threat | Resilient flying finisher with creature recursion.
- 1 Liliana, Heretical Healer // Liliana, Defiant Necromancer | threat | Transforms through creature deaths and generates long-game advantage.
- 1 Mondrak, Glory Dominus | threat | Doubles the deck's creature-token production.
- 1 Razaketh, the Foulblooded | threat | Powerful sacrifice-based tutor and finisher.
- 1 Requiem Angel | threat | Replaces fallen non-Spirit creatures with evasive tokens.
- 1 Liliana, Dreadhorde General | wipe | Board reset with lasting card advantage.
- 1 The Meathook Massacre | wipe | Board wipe that also converts deaths into life swings.
- 1 Toxic Deluge | wipe | Flexible, efficient creature-board reset.

</details>

### 12. Commander on a budget

Format: Commander. Theme: tokens. Pool: any_card. Shortlist: 296 names.

Cards: 99. Repair turn: false. Block findings: 0.

**Summary:** Adeline leads an aggressive go-wide token plan: build a broad board, reinforce it with token-focused synergy cards, and attack with a growing army backed by substantial creature threats. The deck carries card flow, removal, protective interaction, and board resets to stay involved through longer casual games. It gives up multicolor options and depends on establishing and retaining a board, so repeated sweeping effects or concentrated pressure can slow it down.

- [INFO] `bracket_prose_rules`: bracket 2 (Core, 8+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 3.41 over 63 nonland cards

<details><summary>The deck list</summary>

- 36 Plains | land | Provides the deck's mono-white land base.
- 1 Battle Angels of Tyr | ramp | Fills a ramp slot while remaining part of the creature plan.
- 1 Druidic Satchel | ramp | Provides a colorless ramp slot.
- 1 Goldvein Pick | ramp | Provides a colorless ramp slot.
- 1 Karn, Living Legacy | ramp | Fills a ramp slot with a persistent noncreature card.
- 1 Keeper of the Accord | ramp | Fills a ramp slot on a white creature.
- 1 Monologue Tax | ramp | Provides a white ramp slot.
- 1 Nuka-Cola Vending Machine | ramp | Provides a colorless ramp slot.
- 1 Prying Blade | ramp | Provides a colorless ramp slot.
- 1 Tempting Contract | ramp | Provides a colorless ramp slot.
- 1 The Restoration of Eiganjo // Architect of Restoration | ramp | Fills a ramp slot while retaining a white permanent body.
- 1 Bennie Bracks, Zoologist | draw | Provides a draw card suited to a creature-heavy board.
- 1 Bygone Bishop | draw | Fills a draw slot on a white creature.
- 1 Caretaker's Talent | draw | Provides a draw slot for the token plan.
- 1 Chivalric Alliance | draw | Provides a draw slot that matches the deck's creature theme.
- 1 Court of Grace | draw | Fills a white draw slot.
- 1 Dawn of Hope | draw | Provides a white draw slot.
- 1 Idol of Oblivion | draw | Provides a colorless draw slot.
- 1 Staff of the Storyteller | draw | Provides a colorless draw slot.
- 1 Thorough Investigation | draw | Fills a white draw slot.
- 1 Wedding Announcement // Wedding Festivity | draw | Provides a draw slot that fits a go-wide board.
- 1 Aerial Assault | removal | Fills a white removal slot.
- 1 Banishing Slash | removal | Fills a white removal slot.
- 1 Citizen's Crowbar | removal | Provides a colorless removal slot.
- 1 Hanged Executioner | removal | Fills a removal slot on a white creature.
- 1 Kellan's Lightblades | removal | Provides a white removal slot.
- 1 Release to Memory | removal | Provides a white removal slot.
- 1 The Wandering Emperor | removal | Fills a white removal slot with a flexible permanent.
- 1 Trostani's Judgment | removal | Fills a white removal slot.
- 1 Basri Ket | interaction | Provides an interaction slot that supports the white creature plan.
- 1 Elspeth, Knight-Errant | interaction | Provides an interaction slot on a white planeswalker.
- 1 Lena, Selfless Champion | interaction | Fills an interaction slot on a white creature.
- 1 Rootborn Defenses | interaction | Provides a white interaction slot for a wide board.
- 1 Spirit Bonds | interaction | Provides a white interaction slot that fits the token theme.
- 1 Teyo, the Shieldmage | interaction | Fills an interaction slot on a white planeswalker.
- 1 Anointer Priest | synergy | Provides a token-plan synergy piece.
- 1 Cathar's Call | synergy | Provides a white synergy piece for the deck's token plan.
- 1 Charismatic Conqueror | synergy | Provides a token-plan synergy creature.
- 1 Clarion Spirit | synergy | Provides a white creature synergy piece.
- 1 Divine Visitation | synergy | Provides a high-impact token-plan synergy piece.
- 1 Felidar Retreat | synergy | Provides a white synergy piece for building a board.
- 1 Horn of Gondor | synergy | Provides a colorless token-plan synergy piece.
- 1 Intangible Virtue | synergy | Provides a white go-wide synergy piece.
- 1 Luminarch Ascension | synergy | Provides a white token-plan synergy piece.
- 1 Mavren Fein, Dusk Apostle | synergy | Provides a white creature synergy piece.
- 1 Oketra's Monument | synergy | Provides a colorless synergy piece for the creature plan.
- 1 Retreat to Emeria | synergy | Provides a white token-plan synergy piece.
- 1 Rosie Cotton of South Lane | synergy | Provides a white synergy creature for a growing board.
- 1 Skrelv's Hive | synergy | Provides a white token-plan synergy piece.
- 1 Ajani's Chosen | threat | Provides a token-focused white threat.
- 1 Archon of Sun's Grace | threat | Provides a white threat that fits the enchantment package.
- 1 Attended Healer | threat | Provides a white token-focused threat.
- 1 Emeria Angel | threat | Provides a white threat for a board-building plan.
- 1 God-Eternal Oketra | threat | Provides a durable white threat for the creature plan.
- 1 Hero of Bladehold | threat | Provides an aggressive white threat for attacking wide.
- 1 Mite Overseer | threat | Provides a white token-focused threat.
- 1 Oketra the True | threat | Provides a white creature threat.
- 1 Requiem Angel | threat | Provides a white flying threat for the token plan.
- 1 Silverwing Squadron | threat | Provides a white threat for a multiplayer board.
- 1 Warren Warleader | threat | Provides an aggressive white threat.
- 1 Wingmantle Chaplain | threat | Provides a white creature threat that suits a wide board.
- 1 Elspeth, Sun's Champion | wipe | Provides a white board-reset slot that also fits the go-wide plan.
- 1 Hour of Reckoning | wipe | Provides a white board-reset slot for a token deck.
- 1 The Battle of Bywater | wipe | Provides a white board-reset slot.

</details>

### 13. owned first, and the commander is not owned

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 225 names.

Cards: 99. Repair turn: false. Block findings: 0.

**Summary:** This deck keeps Karlov at the center of a lifegain-focused plan, using supporting pieces to build momentum while equipment and protection help preserve its core presence. It wins by turning that momentum into pressure from a broad suite of threats, while removal and battlefield resets keep opposing boards manageable. The tradeoff is a focused black-white shell that favors its central plan over broader multicolor flexibility.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [WARN] `not_owned`: Karlov of the Ghost Council: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 2.97 over 63 nonland cards

<details><summary>The deck list</summary>

- 20 Plains | land | Provides dependable white mana for the deck.
- 16 Swamp | land | Provides dependable black mana for the deck.
- 1 Arcane Signet | ramp | Provides mana acceleration.
- 1 Astral Cornucopia | ramp | Provides mana acceleration.
- 1 Commander's Sphere | ramp | Provides mana acceleration.
- 1 Fellwar Stone | ramp | Provides mana acceleration.
- 1 Giada, Font of Hope | ramp | Provides mana acceleration.
- 1 Lotho, Corrupt Shirriff | ramp | Provides mana acceleration.
- 1 Sol Ring | ramp | Provides mana acceleration.
- 1 Sword of the Animist | ramp | Provides mana acceleration.
- 1 Thought Vessel | ramp | Provides mana acceleration.
- 1 Wayfarer's Bauble | ramp | Provides mana acceleration.
- 1 Buster Sword | draw | Provides card draw.
- 1 Call of the Ring | draw | Provides card draw.
- 1 Exemplar of Light | draw | Provides card draw.
- 1 Idol of Oblivion | draw | Provides card draw.
- 1 Inspiring Overseer | draw | Provides card draw.
- 1 Lembas | draw | Provides card draw.
- 1 Mask of Memory | draw | Provides card draw.
- 1 Night's Whisper | draw | Provides card draw.
- 1 Skullclamp | draw | Provides card draw.
- 1 Tome of Legends | draw | Provides card draw.
- 1 Banishing Light | removal | Provides targeted removal.
- 1 Bitter Triumph | removal | Provides targeted removal.
- 1 Crib Swap | removal | Provides targeted removal.
- 1 Dismember | removal | Provides targeted removal.
- 1 Dispatch | removal | Provides targeted removal.
- 1 Generous Gift | removal | Provides targeted removal.
- 1 Get Lost | removal | Provides targeted removal.
- 1 Swords to Plowshares | removal | Provides targeted removal.
- 1 Bastion Protector | interaction | Helps protect the deck's central plan.
- 1 Boromir, Warden of the Tower | interaction | Helps protect the deck's central plan.
- 1 Champion's Helm | interaction | Helps protect the deck's central plan.
- 1 Darksteel Plate | interaction | Helps protect the deck's central plan.
- 1 Lightning Greaves | interaction | Helps protect the deck's central plan.
- 1 Swiftfoot Boots | interaction | Helps protect the deck's central plan.
- 1 Austere Command | wipe | Resets a crowded battlefield.
- 1 Fumigate | wipe | Resets a crowded battlefield.
- 1 Vanquish the Horde | wipe | Resets a crowded battlefield.
- 1 Aerith Gainsborough | synergy | Supports the lifegain-focused Karlov plan.
- 1 Aettir and Priwen | synergy | Supports the lifegain-focused Karlov plan.
- 1 Angel of Vitality | synergy | Supports the lifegain-focused Karlov plan.
- 1 Compassionate Healer | synergy | Supports the lifegain-focused Karlov plan.
- 1 Dancer's Chakrams | synergy | Supports the lifegain-focused Karlov plan.
- 1 Elixir | synergy | Supports the lifegain-focused Karlov plan.
- 1 Excalibur II | synergy | Supports the lifegain-focused Karlov plan.
- 1 Kor Firewalker | synergy | Supports the lifegain-focused Karlov plan.
- 1 Light of Promise | synergy | Supports the lifegain-focused Karlov plan.
- 1 Night Nurse, Healer of Heroes | synergy | Supports the lifegain-focused Karlov plan.
- 1 Rosie Cotton of South Lane | synergy | Supports the lifegain-focused Karlov plan.
- 1 Second Breakfast | synergy | Supports the lifegain-focused Karlov plan.
- 1 Well-Worn Spatula | synergy | Supports the lifegain-focused Karlov plan.
- 1 White Mage's Staff | synergy | Supports the lifegain-focused Karlov plan.
- 1 Angel of Invention | threat | Provides a proactive threat.
- 1 Bill the Pony | threat | Provides a proactive threat.
- 1 Canyon Crawler | threat | Provides a proactive threat.
- 1 Dawnhand Eulogist | threat | Provides a proactive threat.
- 1 Foggy Swamp Hunters | threat | Provides a proactive threat.
- 1 Gollum, Silent Slinker // Meager Meal | threat | Provides a proactive threat.
- 1 Lyra Dawnbringer | threat | Provides a proactive threat.
- 1 Minwu, White Mage | threat | Provides a proactive threat.
- 1 Rabaroo Troop | threat | Provides a proactive threat.
- 1 Shattered Angel | threat | Provides a proactive threat.
- 1 Sneering Shadewriter | threat | Provides a proactive threat.
- 1 Victory's Herald | threat | Provides a proactive threat.

</details>

