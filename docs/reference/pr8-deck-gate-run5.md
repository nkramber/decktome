# PR-8 deck gate

Run date: 2026-08-28. Card snapshot: 2026-08-24.

Verdict: PASS. 15 of 15 decks passed every block check, 0 invented names reached the user, and 0 summaries stated a false rule of the game. All three bars are zero tolerance.

## Summary

| Measure | Value |
|---|---|
| Prompts | 15 |
| Decks returned | 15 |
| Decks with no block finding | 15 |
| Invented names that reached the user | 0 |
| Decks that needed the repair turn | 0 |
| Summaries judged (F-26) | 15 |
| Summaries that state a rule of the game | 0 |
| Summaries that state a FALSE rule | 0 |
| Errors | 0 |
| Prompt version | 3 |
| Calls | 30 |
| Cost | $0.8102 |
| Time | 671 seconds |

## Findings by code

A block stops the deck. A warning and a note are reports.

| Code | Count |
|---|---|
| `curve_summary` | 15 |
| `bracket_prose_rules` | 9 |
| `basics_added` | 1 |
| `over_budget` | 1 |
| `not_owned` | 1 |

By severity: BLOCK 0. WARN 2. INFO 25. 

## Decks

### 1. lifegain Commander, any card

Format: Commander. Theme: lifegain. Pool: any_card. Shortlist: 303 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $365.51 to buy, $365.51 the whole deck.

**Summary:** This Karlov deck builds around lifegain, using a steady white-black mana base to deploy Clerics, Vampires, Angels, enchantments, and artifacts that reinforce the theme. It wins by turning those lifegain synergies and its larger threats into sustained pressure while keeping opposing boards in check with focused removal and wipes. It gives up the dedicated alternate-win package in favor of a more interactive creature-and-enchantment game plan.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 3.65 over 63 nonland cards

<details><summary>The deck list</summary>

- 20 Plains | land | Provides the deck's basic white land base.
- 10 Swamp | land | Provides the deck's basic black land base.
- 1 Scoured Barrens | land | Adds a lifegain-themed land to the mana base.
- 1 Shambling Vent | land | Adds a white-black land to the mana base.
- 1 Restless Fortress | land | Adds a white-black land to the mana base.
- 1 Kabira Crossroads | land | Adds a lifegain-themed land to the mana base.
- 1 Radiant Fountain | land | Adds a lifegain-themed land to the mana base.
- 1 High Market | land | Adds a utility land to the mana base.
- 1 Alhammarret's Archive | draw | Supports the deck's card-draw package.
- 1 Archivist of Oghma | draw | Provides a creature-based draw option.
- 1 Convalescent Care | draw | Fits the deck's lifegain-focused draw plan.
- 1 Cosmos Elixir | draw | Provides a resilient artifact draw option.
- 1 Dawn of Hope | draw | Fits the deck's lifegain-focused draw plan.
- 1 Enduring Innocence | draw | Provides an enchantment-based draw option.
- 1 Mangara, the Diplomat | draw | Provides a creature-based draw option.
- 1 Markov Purifier | draw | Provides a Vampire draw option.
- 1 Sigarda's Splendor | draw | Fits the deck's lifegain-focused draw plan.
- 1 Well of Lost Dreams | draw | Fits the deck's lifegain-focused draw plan.
- 1 Altar of the Pantheon | ramp | Provides artifact-based ramp.
- 1 Angel of Indemnity | ramp | Provides creature-based ramp.
- 1 Crypt Ghast | ramp | Provides black ramp for the deck.
- 1 Hierophant's Chalice | ramp | Provides artifact-based ramp.
- 1 Nuka-Cola Vending Machine | ramp | Provides artifact-based ramp.
- 1 Orazca Relic | ramp | Provides artifact-based ramp.
- 1 Phial of Galadriel | ramp | Provides legendary artifact ramp.
- 1 Pristine Talisman | ramp | Provides artifact-based ramp for a lifegain deck.
- 1 Redemption Choir | ramp | Provides creature-based ramp.
- 1 The Celestus | ramp | Provides legendary artifact ramp.
- 1 Aetherflux Reservoir | removal | Provides a lifegain-themed removal option.
- 1 Ayli, Eternal Pilgrim | removal | Provides a lifegain-themed removal option.
- 1 Cavalier of Night | removal | Provides creature-based removal.
- 1 Murderous Rider // Swift End | removal | Provides flexible removal.
- 1 Solitude | removal | Provides creature-based removal.
- 1 Umezawa's Jitte | removal | Provides equipment-based removal.
- 1 Vein Ripper | removal | Provides a Vampire removal option.
- 1 Vona, Butcher of Magan | removal | Provides a lifegain-themed removal option.
- 1 Ajani, Strength of the Pride | wipe | Provides a planeswalker-based wipe.
- 1 Fumigate | wipe | Provides a lifegain-themed board wipe.
- 1 Kaya's Wrath | wipe | Provides a white-black board wipe.
- 1 Alseid of Life's Bounty | interaction | Provides creature-based interaction.
- 1 Enduring Angel // Angelic Enforcer | interaction | Provides an Angel interaction option.
- 1 Faith's Shield | interaction | Provides instant-speed interaction.
- 1 Restoration Magic | interaction | Provides instant-speed interaction.
- 1 Sword of Light and Shadow | interaction | Provides equipment-based interaction.
- 1 Werefox Bodyguard | interaction | Provides creature-based interaction.
- 1 Aerith Gainsborough | synergy | Supports the deck's lifegain synergy plan.
- 1 Ajani's Pridemate | synergy | Supports the deck's lifegain synergy plan.
- 1 Angel of Vitality | synergy | Supports the deck's Angel and lifegain themes.
- 1 Angelic Accord | synergy | Supports the deck's lifegain synergy plan.
- 1 Blood Artist | synergy | Supports the deck's black lifegain synergy plan.
- 1 Bloodthirsty Aerialist | synergy | Supports the deck's lifegain synergy plan.
- 1 Cleric Class | synergy | Supports the deck's Cleric and lifegain themes.
- 1 Cleric of Life's Bond | synergy | Supports the deck's Cleric and lifegain themes.
- 1 Cradle of Vitality | synergy | Supports the deck's lifegain synergy plan.
- 1 Heliod, Sun-Crowned | synergy | Supports the deck's lifegain synergy plan.
- 1 Resplendent Angel | synergy | Supports the deck's Angel and lifegain themes.
- 1 Sanguine Bond | synergy | Supports the deck's black lifegain synergy plan.
- 1 Vito, Thorn of the Dusk Rose | synergy | Supports the deck's black lifegain synergy plan.
- 1 Voice of the Blessed | synergy | Supports the deck's lifegain synergy plan.
- 1 Archangel of Thune | threat | Provides an Angel threat for the lifegain plan.
- 1 Archon of Sun's Grace | threat | Provides an enchantment-themed threat.
- 1 Astarion, the Decadent | threat | Provides a black lifegain-themed threat.
- 1 Celestine, the Living Saint | threat | Provides a legendary lifegain-themed threat.
- 1 Defiant Bloodlord | threat | Provides a black lifegain-themed threat.
- 1 Divinity of Pride | threat | Provides a white-black lifegain-themed threat.
- 1 Elenda, Saint of Dusk | threat | Provides a legendary Vampire threat.
- 1 Liesa, Forgotten Archangel | threat | Provides a legendary Angel threat.
- 1 Lyra Dawnbringer | threat | Provides a legendary Angel threat.
- 1 Rhox Faithmender | threat | Provides a lifegain-themed creature threat.
- 1 Valkyrie Harbinger | threat | Provides an Angel Cleric threat.
- 1 Wurmcoil Engine | threat | Provides a colorless artifact creature threat.

</details>

### 2. aristocrats Commander, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 178 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $0.00 to buy, $179.30 the whole deck.

**Summary:** Denethor, Ruling Steward leads a library-first aristocrats build that develops its mana, establishes a board of creatures and synergy pieces, and keeps cards flowing while protecting its important permanents. It aims to win through sustained creature pressure and attrition, using targeted removal and board wipes to keep opposing boards from taking over. The deck gives up some of the density of a tightly specialized aristocrats package in exchange for using a broad selection of available cards and maintaining answers across a long game.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 2.65 over 63 nonland cards

<details><summary>The deck list</summary>

- 15 Plains | land | Provides reliable white mana for the deck’s mana base.
- 14 Swamp | land | Provides reliable black mana for the deck’s mana base.
- 1 Castle Locthwain | land | Adds a black land to the mana base.
- 1 Command Tower | land | Provides flexible mana for the commander deck.
- 1 Evolving Wilds | land | Helps stabilize the mana base.
- 1 Fabled Passage | land | Helps stabilize the mana base.
- 1 Marsh Flats | land | Helps find the deck’s basic lands.
- 1 Path of Ancestry | land | Provides another flexible land for the creature-focused plan.
- 1 Terramorphic Expanse | land | Helps stabilize the mana base.
- 1 Arcane Signet | ramp | Provides early mana acceleration.
- 1 Astral Cornucopia | ramp | Adds a mana-producing artifact to the deck.
- 1 Commander's Sphere | ramp | Provides mana acceleration.
- 1 Deadly Dispute | ramp | Supports the deck’s resource development.
- 1 Fellwar Stone | ramp | Provides efficient mana acceleration.
- 1 Inherited Envelope | ramp | Adds another mana resource.
- 1 Lotho, Corrupt Shirriff | ramp | Contributes to the deck’s mana development.
- 1 Relic of Legends | ramp | Provides mana acceleration for the creature-heavy build.
- 1 Ring of the Lucii | ramp | Adds another ramp piece from the library.
- 1 Sol Ring | ramp | Provides efficient mana acceleration.
- 1 Springleaf Drum | ramp | Supports early mana development alongside creatures.
- 1 Thought Vessel | ramp | Adds a mana-producing artifact.
- 1 Wayfarer's Bauble | ramp | Helps develop the mana base.
- 1 White Lotus Tile | ramp | Adds another available mana resource.
- 1 Ahriman | draw | Provides an additional card-flow creature.
- 1 Buster Sword | draw | Adds card flow while fitting the artifact package.
- 1 Call of the Ring | draw | Provides a continuing source of card flow.
- 1 Foot Chopper | draw | Adds card flow from an available Equipment.
- 1 Grave Venerations | draw | Supports the deck with additional card flow.
- 1 Idol of Oblivion | draw | Provides repeatable card flow for the token- and sacrifice-oriented plan.
- 1 Lembas | draw | Adds a compact card-flow artifact.
- 1 Mask of Memory | draw | Provides card flow through an Equipment slot.
- 1 Massacre Girl, Known Killer | draw | Adds card flow on a creature body.
- 1 Nasty End | draw | Provides additional card flow for the deck.
- 1 Night's Whisper | draw | Provides efficient card flow.
- 1 Puresteel Paladin | draw | Adds card flow while supporting the Equipment presence.
- 1 Skullclamp | draw | Provides strong card flow for a creature-focused aristocrats plan.
- 1 Stone of Erech | draw | Adds another card-flow artifact.
- 1 Wall of Omens | draw | Provides card flow on an early defensive creature.
- 1 Bastion Protector | interaction | Helps protect the commander and important creatures.
- 1 Boromir, Warden of the Tower | interaction | Provides a creature-based interaction piece.
- 1 Darksteel Plate | interaction | Protects a key creature through an Equipment slot.
- 1 Gift of Immortality | interaction | Helps preserve an important creature.
- 1 Lightning Greaves | interaction | Provides efficient protection for key creatures.
- 1 Sheltered by Ghosts | interaction | Adds protection from the available library cards.
- 1 Swiftfoot Boots | interaction | Protects the commander or a major threat.
- 1 Take Up the Shield | interaction | Provides a flexible protective response.
- 1 Together Forever | interaction | Adds a resilient interaction piece for the creature plan.
- 1 Bitter Triumph | removal | Provides flexible creature removal.
- 1 Claim the Precious | removal | Adds targeted removal from the library.
- 1 Destroy Evil | removal | Provides flexible targeted removal.
- 1 Fatal Push | removal | Provides efficient targeted removal.
- 1 Fiend Hunter | removal | Adds creature-based removal to support the aristocrats shell.
- 1 Generous Gift | removal | Provides broad permanent removal.
- 1 Get Lost | removal | Adds flexible targeted removal.
- 1 Gollum the Abandoned | removal | Provides removal on a creature body.
- 1 Infernal Grasp | removal | Provides reliable targeted removal.
- 1 Stroke of Midnight | removal | Adds another flexible removal spell.
- 1 Swords to Plowshares | removal | Provides efficient creature removal.
- 1 Austere Command | wipe | Provides a flexible battlefield reset.
- 1 Dusk // Dawn | wipe | Adds a creature-focused reset for contested boards.
- 1 Fumigate | wipe | Provides a broad creature reset.
- 1 Split Up | wipe | Adds another board-clearing option.
- 1 The Battle of Bywater | wipe | Provides an additional creature wipe.
- 1 Vanquish the Horde | wipe | Adds a dependable board reset.
- 1 Arcade Cabinet | synergy | Provides a colorless synergy piece for the deck’s main plan.
- 1 Gollum, Patient Plotter | synergy | Supports the aristocrats-oriented synergy package.
- 1 Gríma Wormtongue | synergy | Adds a black synergy creature from the library.
- 1 Heirloom Auntie | synergy | Contributes another creature-based synergy piece.
- 1 Joo Dee, One of Many | synergy | Adds to the deck’s available synergy package.
- 1 Nimble Hobbit | synergy | Provides a low-cost synergy creature.
- 1 Phantom Train | synergy | Adds a colorless synergy permanent.
- 1 Vengeful Villagers | threat | Provides a creature threat that complements the deck’s attrition plan.

</details>

### 3. artifacts Commander, bracket 4

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 299 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $1485.10 to buy, $1485.10 the whole deck.

**Summary:** This deck builds a dense artifact board, using its mana acceleration and card flow to keep deploying key pieces while holding up protection and answers. It wins by converting that artifact foundation into pressure through its large artifact threats, with powerful utility artifacts helping it maintain momentum. It gives up some flexibility for a focused, permanent-heavy plan that is most effective when its artifact engine stays established.

- [INFO] `curve_summary`: average mana value 3.83 over 63 nonland cards
- [INFO] `basics_added`: the list was 1 card short, so the builder added 1 basic land

<details><summary>The deck list</summary>

- 1 Academy Ruins | land | Provides a utility land slot for the artifact-focused plan.
- 1 Archway of Innovation | land | Provides a land slot for the artifact-focused plan.
- 1 Blinkmoth Nexus | land | Provides a utility land slot for the artifact-focused plan.
- 1 Buried Ruin | land | Provides a utility land slot for the artifact-focused plan.
- 1 Conqueror's Galleon // Conqueror's Foothold | land | Provides a land slot for the artifact-focused plan.
- 1 Hall of Tagsin | land | Provides a land slot for the artifact-focused plan.
- 1 Inkmoth Nexus | land | Provides a utility land slot for the artifact-focused plan.
- 1 Inventors' Fair | land | Provides a utility land slot for the artifact-focused plan.
- 12 Island | land | Provides the deck's basic blue land base.
- 1 Mech Hangar | land | Provides a land slot for the artifact-focused plan.
- 1 Mirrex | land | Provides a utility land slot for the artifact-focused plan.
- 1 Mishra's Factory | land | Provides a utility land slot for the artifact-focused plan.
- 1 Mishra's Foundry | land | Provides a utility land slot for the artifact-focused plan.
- 1 Mishra's Workshop | land | Provides a land slot for the artifact-focused plan.
- 1 Otawara, Soaring City | land | Provides a utility land slot for the artifact-focused plan.
- 1 Phyrexia's Core | land | Provides a utility land slot for the artifact-focused plan.
- 1 Power Depot | land | Provides an artifact land slot for the plan.
- 1 Primal Amulet // Primal Wellspring | land | Provides a land slot for the artifact-focused plan.
- 1 Roadside Reliquary | land | Provides a utility land slot for the artifact-focused plan.
- 1 The Monumental Facade | land | Provides a utility land slot for the artifact-focused plan.
- 1 The Mycosynth Gardens | land | Provides a utility land slot for the artifact-focused plan.
- 1 Treasure Map // Treasure Cove | land | Provides a land slot for the artifact-focused plan.
- 1 Urza's Factory | land | Provides a land slot for the artifact-focused plan.
- 1 Urza's Saga | land | Provides an artifact-oriented land slot.
- 1 Urza's Workshop | land | Provides a land slot for the artifact-focused plan.
- 1 Forensic Gadgeteer | draw | Supplies draw for the artifact-focused plan.
- 1 One with the Machine | draw | Supplies draw for the artifact-focused plan.
- 1 Reverse Engineer | draw | Supplies draw for the artifact-focused plan.
- 1 Riddlesmith | draw | Supplies draw for the artifact-focused plan.
- 1 Sai, Master Thopterist | draw | Supplies draw for the artifact-focused plan.
- 1 Tezzeret, Artifice Master | draw | Supplies draw for the artifact-focused plan.
- 1 Thirst for Knowledge | draw | Supplies draw for the artifact-focused plan.
- 1 Thought Monitor | draw | Supplies draw for the artifact-focused plan.
- 1 Thoughtcast | draw | Supplies draw for the artifact-focused plan.
- 1 Vedalken Archmage | draw | Supplies draw for the artifact-focused plan.
- 1 Assert Authority | interaction | Provides interaction to protect the artifact plan.
- 1 Disruption Protocol | interaction | Provides interaction to protect the artifact plan.
- 1 Ice Out | interaction | Provides interaction to protect the artifact plan.
- 1 Metallic Rebuke | interaction | Provides interaction to protect the artifact plan.
- 1 Stoic Rebuttal | interaction | Provides interaction to protect the artifact plan.
- 1 Welding Jar | interaction | Provides interaction for the artifact-focused plan.
- 1 Blinkmoth Urn | ramp | Provides ramp for the artifact-focused plan.
- 1 Chief Engineer | ramp | Provides ramp for the artifact-focused plan.
- 1 Grand Architect | ramp | Provides ramp for the artifact-focused plan.
- 1 Inspiring Statuary | ramp | Provides ramp for the artifact-focused plan.
- 1 Karn, Legacy Reforged | ramp | Provides ramp for the artifact-focused plan.
- 1 Krark-Clan Ironworks | ramp | Provides ramp for the artifact-focused plan.
- 1 Metalworker | ramp | Provides ramp for the artifact-focused plan.
- 1 Moonsnare Prototype | ramp | Provides ramp for the artifact-focused plan.
- 1 Mox Opal | ramp | Provides ramp for the artifact-focused plan.
- 1 Tezzeret the Seeker | ramp | Provides ramp for the artifact-focused plan.
- 1 Aether Spellbomb | removal | Provides removal for the artifact-focused plan.
- 1 Aetherflux Reservoir | removal | Provides removal for the artifact-focused plan.
- 1 Portal to Phyrexia | removal | Provides removal for the artifact-focused plan.
- 1 Ravenform | removal | Provides removal for the artifact-focused plan.
- 1 Resculpt | removal | Provides removal for the artifact-focused plan.
- 1 Skysovereign, Consul Flagship | removal | Provides removal for the artifact-focused plan.
- 1 Spine of Ish Sah | removal | Provides removal for the artifact-focused plan.
- 1 Tormod's Crypt | removal | Provides removal for the artifact-focused plan.
- 1 Clock of Omens | synergy | Supports artifact synergies throughout the deck.
- 1 Emry, Lurker of the Loch | synergy | Supports artifact synergies throughout the deck.
- 1 Etherium Sculptor | synergy | Supports artifact synergies throughout the deck.
- 1 Foundry Inspector | synergy | Supports artifact synergies throughout the deck.
- 1 Manifold Key | synergy | Supports artifact synergies throughout the deck.
- 1 Mirran Spy | synergy | Supports artifact synergies throughout the deck.
- 1 Mystic Forge | synergy | Supports artifact synergies throughout the deck.
- 1 Panharmonicon | synergy | Supports artifact synergies throughout the deck.
- 1 Power Artifact | synergy | Supports artifact synergies throughout the deck.
- 1 Scrap Trawler | synergy | Supports artifact synergies throughout the deck.
- 1 Transmute Artifact | synergy | Supports artifact synergies throughout the deck.
- 1 Unwinding Clock | synergy | Supports artifact synergies throughout the deck.
- 1 Voltaic Key | synergy | Supports artifact synergies throughout the deck.
- 1 Whir of Invention | synergy | Supports artifact synergies throughout the deck.
- 1 Argent Sphinx | threat | Provides a threat for closing games with the artifact plan.
- 1 Broodstar | threat | Provides a threat for closing games with the artifact plan.
- 1 Cyberdrive Awakener | threat | Provides a threat for closing games with the artifact plan.
- 1 Darksteel Juggernaut | threat | Provides a threat for closing games with the artifact plan.
- 1 Kappa Cannoneer | threat | Provides a threat for closing games with the artifact plan.
- 1 Karn, Scion of Urza | threat | Provides a threat for closing games with the artifact plan.
- 1 Kuldotha Forgemaster | threat | Provides a threat for closing games with the artifact plan.
- 1 Lodestone Golem | threat | Provides a threat for closing games with the artifact plan.
- 1 Master Transmuter | threat | Provides a threat for closing games with the artifact plan.
- 1 Metalwork Colossus | threat | Provides a threat for closing games with the artifact plan.
- 1 Mycosynth Golem | threat | Provides a threat for closing games with the artifact plan.
- 1 Traxos, Scourge of Kroog | threat | Provides a threat for closing games with the artifact plan.
- 1 Engineered Explosives | wipe | Provides a board wipe for difficult positions.
- 1 Nevinyrral's Disk | wipe | Provides a board wipe for difficult positions.
- 1 Oblivion Stone | wipe | Provides a board wipe for difficult positions.

</details>

### 4. dinosaur tribal, bracket 2

Format: Commander. Theme: dinosaurs. Pool: any_card. Shortlist: 267 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $262.72 to buy, $262.72 the whole deck.

**Summary:** This deck builds toward a table full of Dinosaurs, using ramp and draw to keep the larger creatures coming while Gishath leads the creature-heavy plan. It wins through a steady stream of Dinosaur threats backed by removal, protective interaction, and a few board resets when the table gets crowded. It gives up a broad, highly flexible toolbox in favor of straightforward Dinosaur synergy and combat-focused pressure.

- [INFO] `bracket_prose_rules`: bracket 2 (Core, 8+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 4.13 over 63 nonland cards

<details><summary>The deck list</summary>

- 10 Forest | land | A basic land for the mana base.
- 8 Mountain | land | A basic land for the mana base.
- 7 Plains | land | A basic land for the mana base.
- 1 Command Tower | land | A land for the mana base.
- 1 Path of Ancestry | land | A land for the Dinosaur-focused mana base.
- 1 Canopy Vista | land | A land for the mana base.
- 1 Cinder Glade | land | A land for the mana base.
- 1 Clifftop Retreat | land | A land for the mana base.
- 1 Exotic Orchard | land | A land for the mana base.
- 1 Rootbound Crag | land | A land for the mana base.
- 1 Stomping Ground | land | A land for the mana base.
- 1 Sunpetal Grove | land | A land for the mana base.
- 1 Temple Garden | land | A land for the mana base.
- 1 Sacred Foundry | land | A land for the mana base.
- 1 Sol Ring | ramp | A straightforward ramp piece.
- 1 Arcane Signet | ramp | A straightforward ramp piece.
- 1 Cultivate | ramp | A ramp spell that supports the deck's larger cards.
- 1 Farseek | ramp | A ramp spell that supports the deck's larger cards.
- 1 Nature's Lore | ramp | A ramp spell that supports the deck's larger cards.
- 1 Kodama's Reach | ramp | A ramp spell that supports the deck's larger cards.
- 1 Rampant Growth | ramp | A ramp spell that supports the deck's larger cards.
- 1 Thunderherd Migration | ramp | A Dinosaur-themed ramp card.
- 1 Drover of the Mighty | ramp | A ramp creature for the Dinosaur deck.
- 1 Topiary Stomper | ramp | A Dinosaur ramp creature.
- 1 Beast Whisperer | draw | A draw card to keep resources flowing.
- 1 Garruk's Uprising | draw | A draw card for the creature-focused plan.
- 1 Guardian Project | draw | A draw card for the creature-focused plan.
- 1 Harmonize | draw | A simple draw spell.
- 1 Return of the Wildspeaker | draw | A draw card for the creature-focused plan.
- 1 Ripjaw Raptor | draw | A Dinosaur draw card.
- 1 Rishkar's Expertise | draw | A draw spell for the deck's threat-heavy plan.
- 1 Shamanic Revelation | draw | A draw spell for the creature-focused plan.
- 1 Toski, Bearer of Secrets | draw | A draw creature for the deck's creature-focused plan.
- 1 Vanquisher's Banner | draw | A draw card for the Dinosaur theme.
- 1 Akroma's Will | interaction | An interaction spell for protecting the deck's plan.
- 1 Boros Charm | interaction | An interaction spell for protecting the deck's plan.
- 1 Ephemerate | interaction | An interaction spell for the creature-focused deck.
- 1 Heroic Intervention | interaction | An interaction spell for protecting the deck's plan.
- 1 Lightning Greaves | interaction | An interaction card for supporting a key creature.
- 1 Swiftfoot Boots | interaction | An interaction card for supporting a key creature.
- 1 Apex Altisaur | removal | A Dinosaur removal card.
- 1 Bronzebeak Foragers | removal | A Dinosaur removal card.
- 1 Burning Sun's Avatar | removal | A Dinosaur removal card.
- 1 Ravenous Sailback | removal | A Dinosaur removal card.
- 1 Savage Stomp | removal | A Dinosaur-themed removal spell.
- 1 Thrashing Brontodon | removal | A Dinosaur removal card.
- 1 Tranquil Frillback | removal | A Dinosaur removal card.
- 1 Zacama, Primal Calamity | removal | A Dinosaur removal threat.
- 1 Commune with Dinosaurs | synergy | A synergy card for the Dinosaur theme.
- 1 Dinosaur Stampede | synergy | A synergy card for the Dinosaur theme.
- 1 Huatli's Raptor | synergy | A Dinosaur synergy creature.
- 1 Hunting Velociraptor | synergy | A Dinosaur synergy creature.
- 1 Invasion of Ixalan // Belligerent Regisaur | synergy | A synergy card that fits the Dinosaur theme.
- 1 Kaheera, the Orphanguard | synergy | A synergy creature for the Dinosaur-focused deck.
- 1 Kinjalli's Caller | synergy | A synergy creature for the Dinosaur theme.
- 1 Kinjalli's Sunwing | synergy | A Dinosaur synergy creature.
- 1 Marauding Raptor | synergy | A Dinosaur synergy creature.
- 1 Otepec Huntmaster | synergy | A synergy creature for the Dinosaur theme.
- 1 Raptor Companion | synergy | A Dinosaur synergy creature.
- 1 Raptor Hatchling | synergy | A Dinosaur synergy creature.
- 1 Regal Imperiosaur | synergy | A Dinosaur synergy creature.
- 1 Territorial Hammerskull | synergy | A Dinosaur synergy creature.
- 1 Carnage Tyrant | threat | A Dinosaur threat for the deck's primary plan.
- 1 Etali, Primal Storm | threat | A legendary Dinosaur threat.
- 1 Ghalta and Mavren | threat | A legendary Dinosaur threat.
- 1 Ghalta, Primal Hunger | threat | A legendary Dinosaur threat.
- 1 Ghalta, Stampede Tyrant | threat | A legendary Dinosaur threat.
- 1 Pantlaza, Sun-Favored | threat | A legendary Dinosaur threat.
- 1 Polyraptor | threat | A Dinosaur threat for the creature-focused plan.
- 1 Quartzwood Crasher | threat | A Dinosaur threat for the creature-focused plan.
- 1 Regisaur Alpha | threat | A Dinosaur threat for the creature-focused plan.
- 1 Thundering Spineback | threat | A Dinosaur threat for the creature-focused plan.
- 1 Verdant Sun's Avatar | threat | A Dinosaur Avatar threat.
- 1 Zetalpa, Primal Dawn | threat | A legendary Dinosaur threat.
- 1 Austere Command | wipe | A wipe for resetting difficult boards.
- 1 Blasphemous Act | wipe | A wipe for resetting difficult boards.
- 1 Wakening Sun's Avatar | wipe | A Dinosaur wipe that fits the deck's theme.

</details>

### 5. blink Commander, owned first

Format: Commander. Theme: blink. Pool: owned_first. Shortlist: 168 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $0.00 to buy, $117.93 the whole deck.

**Summary:** This blink deck builds a creature-centered board while using its synergy pieces, draw cards, and removal suite to maintain momentum. It wins by turning that broad creature suite into steady board pressure, protected by interaction and reset buttons when needed. It gives up a more concentrated finishing package for a flexible, value-oriented game plan.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 2.67 over 63 nonland cards

<details><summary>The deck list</summary>

- 36 Plains | land | Basic land for the mana base.
- 1 Adventurer's Airship | draw | Included as a draw card.
- 1 Buster Sword | draw | Included as a draw card.
- 1 Champions of Minas Tirith | draw | Included as a draw creature.
- 1 Diary of Dreams | draw | Included as a draw card.
- 1 Energybending | draw | Included as a draw card.
- 1 Exemplar of Light | draw | Included as a draw creature.
- 1 Faramir, Field Commander | draw | Included as a draw creature.
- 1 Folk Hero | draw | Included as a draw card.
- 1 Idol of Oblivion | draw | Included as a draw card.
- 1 Inspiring Overseer | draw | Included as a draw creature.
- 1 Instant Ramen | draw | Included as a draw card.
- 1 Joined Researchers // Secret Rendezvous | draw | Included as a draw creature.
- 1 Lembas | draw | Included as a draw card.
- 1 Mask of Memory | draw | Included as a draw card.
- 1 Mirror of Galadriel | draw | Included as a draw card.
- 1 Puresteel Paladin | draw | Included as a draw creature.
- 1 Skullclamp | draw | Included as a draw card.
- 1 South Pole Voyager | draw | Included as a draw creature.
- 1 Stone of Erech | draw | Included as a draw card.
- 1 Tome of Legends | draw | Included as a draw card.
- 1 Weapons Vendor | draw | Included as a draw creature.
- 1 Arcane Signet | ramp | Included as a ramp card.
- 1 Bender's Waterskin | ramp | Included as a ramp card.
- 1 Commander's Sphere | ramp | Included as a ramp card.
- 1 Fellwar Stone | ramp | Included as a ramp card.
- 1 Relic of Legends | ramp | Included as a ramp card.
- 1 Sol Ring | ramp | Included as a ramp card.
- 1 Sword of the Animist | ramp | Included as a ramp card.
- 1 Thought Vessel | ramp | Included as a ramp card.
- 1 Wayfarer's Bauble | ramp | Included as a ramp card.
- 1 White Lotus Tile | ramp | Included as a ramp card.
- 1 Bastion Protector | interaction | Included as an interaction creature.
- 1 Boromir, Warden of the Tower | interaction | Included as an interaction creature.
- 1 Frontline Medic | interaction | Included as an interaction creature.
- 1 Lightning Greaves | interaction | Included as an interaction card.
- 1 Slip On the Ring | interaction | Included as an interaction card.
- 1 Swiftfoot Boots | interaction | Included as an interaction card.
- 1 Angel of Condemnation | removal | Included as a removal creature.
- 1 Angel of Sanctions | removal | Included as a removal creature.
- 1 Angel of Serenity | removal | Included as a removal creature.
- 1 Banishing Light | removal | Included as a removal card.
- 1 Crib Swap | removal | Included as a removal card.
- 1 Destroy Evil | removal | Included as a removal card.
- 1 Dispatch | removal | Included as a removal card.
- 1 Fiend Hunter | removal | Included as a removal creature.
- 1 Generous Gift | removal | Included as a removal card.
- 1 Get Lost | removal | Included as a removal card.
- 1 Journey to Nowhere | removal | Included as a removal card.
- 1 Kataki, War's Wage | removal | Included as a removal creature.
- 1 Make Your Move | removal | Included as a removal card.
- 1 March of Otherworldly Light | removal | Included as a removal card.
- 1 Oblation | removal | Included as a removal card.
- 1 Palace Jailer | removal | Included as a removal creature.
- 1 Swords to Plowshares | removal | Included as a removal card.
- 1 Stroke of Midnight | removal | Included as a removal card.
- 1 Westfold Rider | removal | Included as a removal creature.
- 1 Austere Command | wipe | Included as a board wipe.
- 1 Fumigate | wipe | Included as a board wipe.
- 1 Martial Coup | wipe | Included as a board wipe.
- 1 Ennis, Debate Moderator | synergy | Included for blink-focused synergy.
- 1 Flickerwisp | synergy | Included for blink-focused synergy.
- 1 Jocasta, Automaton Avenger | synergy | Included for blink-focused synergy.
- 1 Personify | synergy | Included for blink-focused synergy.

</details>

### 6. Modern tempo, tournament

Format: Modern. Theme: tempo. Pool: any_card. Shortlist: 167 names.

Cards: 60 main, 15 sideboard. Repair turn: false. Block findings: 0.

Cost: $613.98 to buy, $613.98 the whole deck.

**Summary:** This deck leads with inexpensive creature pressure, then uses draw spells, removal, and interaction to keep the opponent from stabilizing. It wins by maintaining that pressure while the time-focused spell package adds another angle to the game. The tradeoff is a lean, proactive construction that relies on its early threats and careful timing rather than a large late-game board.

- [INFO] `curve_summary`: average mana value 2.28 over 36 nonland cards

<details><summary>The deck list</summary>

- 4 Ragavan, Nimble Pilferer | threat | Early creature threat for the tempo plan.
- 4 Ledger Shredder | threat | Creature threat that supports a spell-heavy tempo shell.
- 4 Faerie Mastermind | threat | Creature threat that helps apply pressure in the air.
- 4 Consider | draw | Cheap draw spell that keeps the deck moving.
- 2 Preordain | draw | Early draw and selection for smoother tempo turns.
- 4 Counterspell | interaction | Broad interaction for protecting pressure and disrupting opponents.
- 2 Spell Pierce | interaction | Cheap interaction that fits a proactive tempo plan.
- 4 Lightning Bolt | removal | Efficient instant removal for clearing blockers or pressuring life totals.
- 2 Into the Flood Maw | removal | Low-cost removal that preserves tempo.
- 2 Untimely Malfunction | removal | Flexible instant removal for problematic permanents.
- 2 Temporal Mastery | synergy | Time-focused spell for the deck's synergy package.
- 2 Temporal Trespass | synergy | Time-focused spell that complements Temporal Mastery.
- 6 Island | land | Basic blue land for the deck's core spells.
- 4 Mountain | land | Basic red land for removal and threats.
- 4 Steam Vents | land | Blue-red land that supports both halves of the deck.
- 4 Scalding Tarn | land | Land that helps access the deck's basic and blue-red lands.
- 2 Shivan Reef | land | Blue-red land for reliable early development.
- 2 Stormcarved Coast | land | Blue-red land that supports the tempo curve.
- 2 Sink into Stupor // Soporific Springs | land | Land slot with flexibility for the spell-heavy plan.

</details>

### 7. Modern burn, casual

Format: Modern. Theme: burn. Pool: any_card. Shortlist: 290 names.

Cards: 60 main, 15 sideboard. Repair turn: false. Block findings: 0.

Cost: $48.40 to buy, $48.40 the whole deck.

**Summary:** This deck applies pressure with its threats, supports them with a focused synergy package, and uses removal to keep its burn-oriented plan moving forward. It wins by sustaining pressure over successive turns rather than shifting into a broad control role. The tradeoff is narrower interaction and less late-game flexibility in exchange for a direct, aggressive game plan.

- [INFO] `curve_summary`: average mana value 2.83 over 36 nonland cards

<details><summary>The deck list</summary>

- 24 Mountain | land | Provides the consistent mono-red mana base.
- 4 Risk Factor | draw | Fills a draw role while maintaining the deck's aggressive focus.
- 2 Grab the Prize | draw | Rounds out the draw package for a burn-oriented plan.
- 2 Chandra, Dressed to Kill | ramp | Supplies the deck's ramp component.
- 4 Lightning Bolt | removal | Core efficient removal for clearing the way for pressure.
- 2 Lightning Strike | removal | Adds reliable removal alongside the primary burn suite.
- 4 Eidolon of the Great Revel | synergy | Provides a synergy piece for the deck's aggressive plan.
- 4 Thermo-Alchemist | synergy | Supports the deck's burn-focused synergy package.
- 4 Akki Lavarunner // Tok-Tok, Volcano Born | threat | Adds an early threat to start applying pressure.
- 4 Keral Keep Disciples | threat | Provides a substantial share of the creature threat package.
- 4 Fuming Effigy | threat | Adds another threat that keeps the deck's pressure dense.
- 2 Hazoret the Fervent | threat | Serves as a high-impact threat near the top of the curve.

</details>

### 8. Modern lifegain, FNM

Format: Modern. Theme: lifegain. Pool: any_card. Shortlist: 301 names.

Cards: 60 main, 15 sideboard. Repair turn: false. Block findings: 0.

Cost: $343.47 to buy, $343.47 the whole deck.

**Summary:** This white-black lifegain deck builds a board around its synergy pieces, keeps cards flowing, and turns that setup into pressure from Angels, Vampires, and other creature threats. It aims to win by establishing a durable battlefield and forcing opponents to answer multiple threats, with removal clearing the way when needed. The tradeoff is that the deck leans on developing its pieces together, so it is less focused on a single fast finish and can be pressured before its board is established.

- [INFO] `curve_summary`: average mana value 3.33 over 36 nonland cards

<details><summary>The deck list</summary>

- 8 Plains | land | Land slot for consistent development.
- 8 Swamp | land | Land slot for consistent development.
- 4 Scoured Barrens | land | Land slot in the white-black base.
- 2 Shambling Vent | land | Land slot in the white-black base.
- 2 Kabira Crossroads | land | Land slot in the white-black base.
- 2 Orzhov Keyrune | ramp | Fills a ramp slot while staying in the deck's colors.
- 2 Dawn of Hope | draw | Provides a draw piece for the lifegain plan.
- 2 Inspiring Overseer | draw | Provides a draw creature for the deck.
- 2 Well of Lost Dreams | draw | Provides another draw piece for the lifegain plan.
- 2 Murderous Rider // Swift End | removal | Flexible removal for opposing threats.
- 2 Ayli, Eternal Pilgrim | removal | Creature-based removal for the main deck.
- 2 Nightmare's Thirst | removal | Additional focused removal.
- 2 Soul Warden | synergy | Core creature for the lifegain synergy plan.
- 2 Ajani's Pridemate | synergy | Core payoff in the lifegain synergy package.
- 2 Cleric of Life's Bond | synergy | Supports the deck's lifegain synergy plan.
- 2 Vito, Thorn of the Dusk Rose | synergy | Key black-white lifegain synergy piece.
- 2 Archangel of Thune | threat | A top-end threat for closing games.
- 2 Blood Baron of Vizkopa | threat | A sturdy creature threat for the deck.
- 2 Angel of Invention | threat | A creature threat that advances the board plan.
- 2 Attended Healer | threat | A creature threat suited to the deck's plan.
- 2 Bloodthirsty Conqueror | threat | A black creature threat for finishing games.
- 2 Rhox Faithmender | threat | A substantial threat for the lifegain deck.
- 2 Valkyrie Harbinger | threat | A high-impact creature threat.

</details>

### 9. Standard midrange, FNM

Format: Standard. Theme: midrange. Pool: any_card. Shortlist: 164 names.

Cards: 60 main, 15 sideboard. Repair turn: false. Block findings: 0.

Cost: $74.47 to buy, $74.47 the whole deck.

**Summary:** This black-green midrange deck builds a steady mana base, uses ramp and card draw to keep resources flowing, and controls the battlefield with targeted removal and sweepers. It wins by clearing space for its creature cards to take over the board, while protection spells help preserve the permanents that matter. The tradeoff is that the list is built to grind through exchanges rather than race quickly, so it can be less proactive when an opponent does not commit much to the battlefield.

- [INFO] `curve_summary`: average mana value 2.56 over 36 nonland cards

<details><summary>The deck list</summary>

- 8 Forest | land | Basic land forming part of the green mana base.
- 8 Swamp | land | Basic land forming part of the black mana base.
- 4 Overgrown Tomb | land | Dual land supporting the black-green mana base.
- 4 Blooming Marsh | land | Dual land supporting the black-green mana base.
- 2 Llanowar Elves | ramp | Early ramp for the midrange plan.
- 2 Phyrexian Arena | draw | Core source of ongoing card draw.
- 2 Darkstar Augur | draw | Creature-based card draw for the board-focused plan.
- 2 Midnight Reaper | draw | Creature-based card draw that supports trading resources.
- 2 Bitter Triumph | removal | Efficient removal in the main interaction package.
- 2 Nowhere to Run | removal | Additional removal for opposing permanents.
- 2 Scavenging Ooze | removal | Creature-based removal that contributes to board presence.
- 4 Snakeskin Veil | interaction | Protection for important creatures and permanents.
- 4 Royal Treatment | interaction | Additional protection to preserve the board.
- 4 Undying Malice | interaction | Protective interaction for creature-heavy exchanges.
- 2 Swiftfoot Boots | interaction | Repeatable protection equipment for key creatures.
- 4 Massacre Wurm | wipe | Creature-based board wipe that also supplies a substantial body.
- 4 Season of Loss | wipe | Main-deck sweeper coverage for crowded boards.

</details>

### 10. Standard aggro, tournament

Format: Standard. Theme: aggro. Pool: any_card. Shortlist: 294 names.

Cards: 60 main, 15 sideboard. Repair turn: false. Block findings: 0.

Cost: $78.66 to buy, $78.66 the whole deck.

**Summary:** This deck aims to establish an aggressive red-white board early and win by turning its threats sideways while Warleader's Call supports the creature-focused pressure. Draw cards help keep the attack supplied, and removal plus interaction clear a path or protect the plan. It gives up main-deck reset effects and a slower, higher-end game in favor of speed and consistency.

- [INFO] `curve_summary`: average mana value 2.94 over 36 nonland cards

<details><summary>The deck list</summary>

- 8 Mountain | land | Basic land for the red side of the mana base.
- 8 Plains | land | Basic land for the white side of the mana base.
- 4 Sacred Foundry | land | Red-white land slot in the mana base.
- 4 Inspiring Vantage | land | Land slot supporting the aggressive mana base.
- 4 Fugitive Codebreaker | draw | Early red draw option that supports the aggressive plan.
- 2 Reckless Lackey | draw | Additional red draw option for keeping pressure supplied.
- 4 Boros Charm | interaction | Flexible red-white interaction for protecting the attack plan.
- 2 Dawn's Truce | interaction | Additional interaction for contested games.
- 4 Harsh Annotation | removal | Efficient removal slot for clearing resistance to attacks.
- 4 Case of the Gateway Express | removal | Additional removal to keep opposing threats in check.
- 4 Warleader's Call | synergy | Red-white synergy piece for the deck's creature-focused pressure.
- 4 Bedhead Beastie | threat | Low-curve aggressive threat.
- 4 Frilled Sparkshooter | threat | Red threat that keeps the deck focused on pressure.
- 4 Teapot Slinger | threat | Additional aggressive threat for a consistent creature base.

</details>

### 11. Commander with a locked card

Format: Commander. Theme: sacrifice. Pool: any_card. Shortlist: 296 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $379.76 to buy, $379.76 the whole deck.

**Summary:** This deck builds around a focused sacrifice plan, using creatures and synergy pieces to keep the engine moving while draw and ramp maintain momentum. It wins by developing its threat cards and clearing resistance with removal and battlefield resets, then pressing its advantage through the sacrifice-oriented board. It gives up broader utility for a concentrated creature-based plan that is strongest when its pieces work together.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 3.11 over 63 nonland cards

<details><summary>The deck list</summary>

- 18 Plains | land | Basic Plains for the white portion of the mana base.
- 18 Swamp | land | Basic Swamps for the black portion of the mana base.
- 1 Baron Bertram Graywater | draw | A designated draw piece for maintaining resources.
- 1 Bushmeat Poacher | draw | A designated draw piece that fits the sacrifice plan.
- 1 Corrupted Conviction | draw | A designated draw piece for turning sacrifice material into resources.
- 1 Disciple of Bolas | draw | A designated draw piece for the creature-focused shell.
- 1 Gas Guzzler | draw | A designated draw piece for the deck's core plan.
- 1 Lord Skitter's Butcher | draw | A designated draw piece for sustaining pressure.
- 1 Relic Vial | draw | A designated draw piece in the artifact suite.
- 1 Smothering Abomination | draw | A designated draw piece that supports the sacrifice theme.
- 1 Vampiric Rites | draw | A designated draw piece built around sacrifice.
- 1 Village Rites | draw | A designated draw piece that works with expendable creatures.
- 1 Cartel Aristocrat | interaction | A designated interaction piece that fits the sacrifice shell.
- 1 Dark Privilege | interaction | A designated interaction piece for protecting the plan.
- 1 Fanatical Devotion | interaction | A designated interaction piece for the creature base.
- 1 Flare of Fortitude | interaction | A designated interaction piece for defending the board.
- 1 Gift of Doom | interaction | A designated interaction piece for protecting a key permanent.
- 1 Nightmare Shepherd | interaction | A designated interaction piece that complements the creature plan.
- 1 Sol Ring | ramp | A required ramp piece for accelerating the deck.
- 1 Ashnod's Altar | ramp | A designated ramp piece that fits sacrifice-focused play.
- 1 Crowded Crypt | ramp | A designated ramp piece for the deck's resource plan.
- 1 Culling the Weak | ramp | A designated ramp piece that uses sacrifice material.
- 1 Deadly Dispute | ramp | A designated ramp piece that works with expendable creatures.
- 1 Phyrexian Altar | ramp | A designated ramp piece central to the sacrifice shell.
- 1 Pitiless Plunderer | ramp | A designated ramp piece that supports the creature plan.
- 1 Priest of Forgotten Gods | ramp | A designated ramp piece for the sacrifice strategy.
- 1 Sifter of Skulls | ramp | A designated ramp piece that complements creature sacrifices.
- 1 Warren Soultrader | ramp | A designated ramp piece in the sacrifice-focused build.
- 1 Attrition | removal | A designated removal piece for controlling opposing threats.
- 1 Ayli, Eternal Pilgrim | removal | A designated removal piece that fits the commander’s colors.
- 1 Bone Shards | removal | A designated removal piece that works with sacrifice material.
- 1 Dictate of Erebos | removal | A designated removal piece for punishing opposing boards.
- 1 Eaten Alive | removal | A designated removal piece for dealing with key threats.
- 1 Grave Pact | removal | A designated removal piece that rewards the sacrifice plan.
- 1 Rite of Consumption | removal | A designated removal piece that complements the creature shell.
- 1 Yawgmoth, Thran Physician | removal | A designated removal piece for the sacrifice-focused core.
- 1 Altar of Dementia | synergy | A designated synergy piece for converting creatures into value.
- 1 Bastion of Remembrance | synergy | A designated synergy piece for the sacrifice plan.
- 1 Carrion Feeder | synergy | A designated synergy piece and sacrifice-focused creature.
- 1 Elas il-Kor, Sadistic Pilgrim | synergy | A designated synergy piece for the creature-sacrifice shell.
- 1 Fleshtaker | synergy | A designated synergy piece that rewards the core plan.
- 1 Hidden Stockpile | synergy | A designated synergy piece for maintaining sacrifice material.
- 1 Nantuko Husk | synergy | A designated synergy piece and sacrifice-focused creature.
- 1 Open the Graves | synergy | A designated synergy piece for the creature-heavy plan.
- 1 Pious Evangel // Wayward Disciple | synergy | A designated synergy piece for the sacrifice theme.
- 1 Pitiless Pontiff | synergy | A designated synergy piece in the sacrifice creature suite.
- 1 Viscera Seer | synergy | A designated synergy piece and sacrifice-focused creature.
- 1 Witch's Oven | synergy | A designated synergy piece for expendable creatures.
- 1 Woe Strider | synergy | A designated synergy piece for the sacrifice-focused core.
- 1 Zulaport Cutthroat | synergy | A designated synergy piece for the deck’s central plan.
- 1 Abhorrent Overlord | threat | A designated threat that adds a powerful top end.
- 1 Basri's Lieutenant | threat | A designated threat for applying board pressure.
- 1 Felisa, Fang of Silverquill | threat | A designated threat that fits the commander’s colors.
- 1 Ghoulcaller Gisa | threat | A designated threat for the creature-focused strategy.
- 1 Liesa, Forgotten Archangel | threat | A designated threat for closing games.
- 1 Liliana, Heretical Healer // Liliana, Defiant Necromancer | threat | A designated threat that complements the sacrifice shell.
- 1 Mondrak, Glory Dominus | threat | A designated threat for increasing battlefield pressure.
- 1 Ratadrabik of Urborg | threat | A designated threat in the legendary creature suite.
- 1 Razaketh, the Foulblooded | threat | A designated threat at the top of the curve.
- 1 Requiem Angel | threat | A designated threat for the creature-heavy plan.
- 1 Sidisi, Undead Vizier | threat | A designated threat that fits the sacrifice strategy.
- 1 Vindictive Vampire | threat | A designated threat that supports the deck’s theme.
- 1 Liliana, Dreadhorde General | wipe | A designated wipe for resetting crowded boards.
- 1 The Meathook Massacre | wipe | A designated wipe for stabilizing the battlefield.
- 1 Toxic Deluge | wipe | A designated wipe for clearing opposing threats.

</details>

### 12. Commander on a budget

Format: Commander. Theme: tokens. Pool: any_card. Shortlist: 296 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $268.37 to buy, $268.37 the whole deck.

**Summary:** The deck establishes a wide token board, uses token-focused synergies to turn that board into sustained combat pressure, and finishes by overwhelming opponents with its threats. It supports that plan with mana development, card access, answers, and board-reset options, but its choices are concentrated on maintaining a board rather than pursuing a separate combo finish.

- [INFO] `bracket_prose_rules`: bracket 2 (Core, 8+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 3.44 over 63 nonland cards
- [WARN] `over_budget`: the cards you must buy cost about $268.37, and the budget is $100.00

<details><summary>The deck list</summary>

- 36 Plains | land | Provides the deck's land base.
- 1 Battle Angels of Tyr | ramp | Included as ramp support for the token plan.
- 1 Charisma Bobblehead | ramp | Included as ramp support for the token plan.
- 1 Coin of Mastery | ramp | Included as ramp support for the token plan.
- 1 Druidic Satchel | ramp | Included as ramp support for the token plan.
- 1 Fishing Gear | ramp | Included as ramp support for the token plan.
- 1 Karn, Living Legacy | ramp | Included as ramp support for the token plan.
- 1 Keeper of the Accord | ramp | Included as ramp support for the token plan.
- 1 Monologue Tax | ramp | Included as ramp support for the token plan.
- 1 Noble's Purse | ramp | Included as ramp support for the token plan.
- 1 Tempting Contract | ramp | Included as ramp support for the token plan.
- 1 Bygone Bishop | draw | Included as card-access support for the token plan.
- 1 Caretaker's Talent | draw | Included as card-access support for the token plan.
- 1 Chivalric Alliance | draw | Included as card-access support for the token plan.
- 1 Court of Grace | draw | Included as card-access support for the token plan.
- 1 Dawn of Hope | draw | Included as card-access support for the token plan.
- 1 Idol of Oblivion | draw | Included as card-access support for the token plan.
- 1 Sanctuary Warden | draw | Included as card-access support for the token plan.
- 1 Staff of the Storyteller | draw | Included as card-access support for the token plan.
- 1 Thorough Investigation | draw | Included as card-access support for the token plan.
- 1 Wedding Announcement // Wedding Festivity | draw | Included as card-access support for the token plan.
- 1 Blessed Sanctuary | interaction | Included to interact while preserving the token plan.
- 1 Lena, Selfless Champion | interaction | Included to interact while preserving the token plan.
- 1 Moogles' Valor | interaction | Included to interact while preserving the token plan.
- 1 Parting Gust | interaction | Included to interact while preserving the token plan.
- 1 Rootborn Defenses | interaction | Included to interact while preserving the token plan.
- 1 Spirit Bonds | interaction | Included to interact while preserving the token plan.
- 1 Aerial Assault | removal | Included as a targeted answer for opposing cards.
- 1 Banishing Slash | removal | Included as a targeted answer for opposing cards.
- 1 Citizen's Crowbar | removal | Included as a targeted answer for opposing cards.
- 1 Hanged Executioner | removal | Included as a targeted answer for opposing cards.
- 1 Kellan's Lightblades | removal | Included as a targeted answer for opposing cards.
- 1 Path to Redemption | removal | Included as a targeted answer for opposing cards.
- 1 Release to Memory | removal | Included as a targeted answer for opposing cards.
- 1 The Wandering Emperor | removal | Included as a targeted answer for opposing cards.
- 1 Elspeth, Sun's Champion | wipe | Included as a board-reset option.
- 1 Hour of Reckoning | wipe | Included as a board-reset option.
- 1 The Battle of Bywater | wipe | Included as a board-reset option.
- 1 Aligned Heart | synergy | Included for its assigned token synergy.
- 1 Animation Module | synergy | Included for its assigned token synergy.
- 1 Anointer Priest | synergy | Included for its assigned token synergy.
- 1 Automated Assembly Line | synergy | Included for its assigned token synergy.
- 1 Charismatic Conqueror | synergy | Included for its assigned token synergy.
- 1 Clarion Spirit | synergy | Included for its assigned token synergy.
- 1 Divine Visitation | synergy | Included for its assigned token synergy.
- 1 Felidar Retreat | synergy | Included for its assigned token synergy.
- 1 Horn of Gondor | synergy | Included for its assigned token synergy.
- 1 Intangible Virtue | synergy | Included for its assigned token synergy.
- 1 Ocelot Pride | synergy | Included for its assigned token synergy.
- 1 Retrofitter Foundry | synergy | Included for its assigned token synergy.
- 1 Rosie Cotton of South Lane | synergy | Included for its assigned token synergy.
- 1 Skrelv's Hive | synergy | Included for its assigned token synergy.
- 1 Archon of Sun's Grace | threat | Included as a substantial token-plan threat.
- 1 God-Eternal Oketra | threat | Included as a substantial token-plan threat.
- 1 Hero of Bladehold | threat | Included as a substantial token-plan threat.
- 1 Mondrak, Glory Dominus | threat | Included as a substantial token-plan threat.
- 1 Myrel, Shield of Argive | threat | Included as a substantial token-plan threat.
- 1 Oketra the True | threat | Included as a substantial token-plan threat.
- 1 Pentavus | threat | Included as a substantial token-plan threat.
- 1 Phantom General | threat | Included as a substantial token-plan threat.
- 1 Requiem Angel | threat | Included as a substantial token-plan threat.
- 1 Serra the Benevolent | threat | Included as a substantial token-plan threat.
- 1 Silverwing Squadron | threat | Included as a substantial token-plan threat.
- 1 Threefold Thunderhulk | threat | Included as a substantial token-plan threat.

</details>

### 13. owned first, and the commander is not owned

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 225 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $0.00 to buy, $209.85 the whole deck.

**Summary:** Karlov leads a lifegain-focused board deck that develops mana, assembles synergistic creatures and equipment, and protects its key pieces while applying steady combat pressure. It wins by building a threatening battlefield and using focused answers or broad resets to clear the way for its threats. The deck gives up some explosive speed in exchange for a stable mana base, sustained resources, and tools to play through longer games.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [WARN] `not_owned`: Karlov of the Ghost Council: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.00 over 63 nonland cards

<details><summary>The deck list</summary>

- 13 Plains | land | Provides reliable white mana for the deck.
- 11 Swamp | land | Provides reliable black mana for the deck.
- 1 Ash Barrens | land | A land slot that helps keep the mana base flexible.
- 1 Bonders' Enclave | land | A utility land for the mana base.
- 1 Castle Locthwain | land | A black utility land for the mana base.
- 1 Command Tower | land | A reliable multicolor land for the commander deck.
- 1 Evolving Wilds | land | A land slot that helps find needed basic mana.
- 1 Fabled Passage | land | A land slot that helps find needed basic mana.
- 1 Marsh Flats | land | A land slot that helps access the deck's core colors.
- 1 Minas Tirith | land | A white utility land for the mana base.
- 1 Plaza of Heroes | land | A utility land that supports a commander-focused deck.
- 1 Rogue's Passage | land | A utility land that supports combat pressure.
- 1 Takenuma, Abandoned Mire | land | A black utility land for the mana base.
- 1 War Room | land | A utility land that supports longer games.
- 1 Arcane Signet | ramp | Efficient mana acceleration for the deck's colors.
- 1 Sol Ring | ramp | Fast mana acceleration for developing the board.
- 1 Fellwar Stone | ramp | A flexible mana rock for early development.
- 1 Commander's Sphere | ramp | Mana acceleration that supports the commander deck.
- 1 Thought Vessel | ramp | A mana rock that helps sustain resources.
- 1 Wayfarer's Bauble | ramp | Early mana development for the deck.
- 1 Sword of the Animist | ramp | An equipment-based ramp piece for the combat plan.
- 1 Giada, Font of Hope | ramp | A ramp creature that fits the deck's creature plan.
- 1 Lotho, Corrupt Shirriff | ramp | A ramp creature that supports resource development.
- 1 Oath of the Grey Host | ramp | A ramp card for building toward the deck's larger turns.
- 1 Buster Sword | draw | An equipment that supplies card advantage.
- 1 Call of the Ring | draw | A draw piece for maintaining resources.
- 1 Exemplar of Light | draw | A creature that contributes card advantage.
- 1 Idol of Oblivion | draw | A compact artifact source of card advantage.
- 1 Inspiring Overseer | draw | A creature that helps refill resources.
- 1 Lembas | draw | An artifact draw piece that fits the deck's plan.
- 1 Mask of Memory | draw | An equipment that supports combat-based card advantage.
- 1 Puresteel Paladin | draw | A creature that supports the equipment and draw package.
- 1 Skullclamp | draw | An efficient equipment source of card advantage.
- 1 Tome of Legends | draw | A commander-friendly artifact draw piece.
- 1 Lightning Greaves | interaction | Protects an important creature in the deck's plan.
- 1 Swiftfoot Boots | interaction | Provides protection for key creatures.
- 1 Darksteel Plate | interaction | An equipment protection piece for a major threat.
- 1 Champion's Helm | interaction | Helps protect the commander-focused plan.
- 1 Clever Concealment | interaction | A protective interaction spell for preserving the board.
- 1 Unbreakable Formation | interaction | Protective interaction for key combat turns.
- 1 Swords to Plowshares | removal | A focused answer to an opposing problem.
- 1 Generous Gift | removal | A flexible answer for troublesome permanents.
- 1 Get Lost | removal | A flexible removal spell for opposing threats.
- 1 Infernal Grasp | removal | An efficient answer to an opposing creature.
- 1 Bitter Triumph | removal | A flexible removal option for key threats.
- 1 Dispatch | removal | A low-cost removal spell for protecting the board plan.
- 1 Crib Swap | removal | A creature answer that fits the deck's colors.
- 1 Stroke of Midnight | removal | A broad answer to troublesome permanents.
- 1 Austere Command | wipe | A flexible reset button when the board gets out of hand.
- 1 Fumigate | wipe | A board wipe for recovering from crowded boards.
- 1 Vanquish the Horde | wipe | A clean reset for opposing creature boards.
- 1 Aettir and Priwen | synergy | An equipment synergy piece for the deck's main plan.
- 1 Angel of Vitality | synergy | A synergy creature for the lifegain-focused plan.
- 1 Compassionate Healer | synergy | A synergy creature for the deck's main plan.
- 1 Crowd of True Believers | synergy | A synergy creature that supports board development.
- 1 Dancer's Chakrams | synergy | An equipment synergy piece for the creature plan.
- 1 Eastfarthing Farmer | synergy | A synergy creature for the deck's main plan.
- 1 Elixir | synergy | An artifact synergy piece for the lifegain-focused plan.
- 1 Excalibur II | synergy | An equipment synergy piece for combat development.
- 1 Kor Firewalker | synergy | A synergy creature for the lifegain-focused plan.
- 1 Light of Promise | synergy | A synergy enchantment for building a key creature.
- 1 Night Nurse, Healer of Heroes | synergy | A synergy creature for the deck's main plan.
- 1 Rosie Cotton of South Lane | synergy | A synergy creature for the lifegain-focused plan.
- 1 Second Breakfast | synergy | A synergy spell for advancing the deck's plan.
- 1 Well-Worn Spatula | synergy | An equipment synergy piece for the creature plan.
- 1 Angel of Invention | threat | A substantial creature for applying pressure.
- 1 Bill the Pony | threat | A creature threat that advances the board plan.
- 1 Dawnhand Eulogist | threat | A creature threat for the deck's combat plan.
- 1 Foggy Swamp Hunters | threat | A creature threat that adds battlefield pressure.
- 1 Gollum, Silent Slinker // Meager Meal | threat | A threat that supports the deck's pressure plan.
- 1 Invisible Woman, Sue Storm | threat | A legendary threat for the deck's board plan.
- 1 Lo and Li, Twin Tutors | threat | A legendary threat for applying pressure.
- 1 Lyra Dawnbringer | threat | A powerful creature threat for closing games.
- 1 Minwu, White Mage | threat | A legendary creature threat for the main plan.
- 1 Shattered Angel | threat | A creature threat that supports the lifegain-focused plan.
- 1 Sneering Shadewriter | threat | A creature threat that adds pressure to the board.
- 1 Victory's Herald | threat | A creature threat for finishing through combat.

</details>

### 14. the user delegates the commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 244 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $744.96 to buy, $744.96 the whole deck.

**Summary:** This is a Dragon-focused Atarka, World Render deck that builds its mana, develops Dragon synergy, and turns a steady stream of large threats into combat pressure. It wins by committing Dragons to the table and forcing opponents to answer repeated attacks, while its removal, wipes, and protective interaction help clear or preserve the way. The deck gives up some flexibility to remain focused on Dragons and its combat plan.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 4.51 over 63 nonland cards

<details><summary>The deck list</summary>

- 12 Forest | land | Basic green land for the deck’s mana base.
- 10 Mountain | land | Basic red land for the deck’s mana base.
- 1 Command Tower | land | Multicolor land for the deck’s mana base.
- 1 Stomping Ground | land | Red-green land for the deck’s mana base.
- 1 Cinder Glade | land | Red-green land for the deck’s mana base.
- 1 Rootbound Crag | land | Red-green land for the deck’s mana base.
- 1 Path of Ancestry | land | Dragon-focused land for the deck’s mana base.
- 1 Temple of the Dragon Queen | land | Dragon-focused land for the deck’s mana base.
- 1 Cavern of Souls | land | Creature-focused land for the deck’s mana base.
- 1 Haven of the Spirit Dragon | land | Dragon-focused land for the deck’s mana base.
- 1 Boseiju, Who Endures | land | Utility land for the deck’s mana base.
- 1 Yavimaya, Cradle of Growth | land | Green utility land for the deck’s mana base.
- 1 Wooded Foothills | land | Land-finding option for the deck’s mana base.
- 1 Windswept Heath | land | Land-finding option for the deck’s mana base.
- 1 Arid Mesa | land | Land-finding option for the deck’s mana base.
- 1 Fabled Passage | land | Land-finding option for the deck’s mana base.
- 1 Ancient Copper Dragon | ramp | Dragon ramp piece for the deck’s mana development.
- 1 Atsushi, the Blazing Sky | ramp | Dragon ramp piece for the deck’s mana development.
- 1 Carnelian Orb of Dragonkind | ramp | Dragon-focused artifact ramp for the deck.
- 1 Dragon's Hoard | ramp | Dragon-focused artifact ramp for the deck.
- 1 Ganax, Astral Hunter | ramp | Dragon ramp piece for the deck’s mana development.
- 1 Goldspan Dragon | ramp | Dragon ramp piece for the deck’s mana development.
- 1 Jade Orb of Dragonkind | ramp | Dragon-focused artifact ramp for the deck.
- 1 Klauth, Unrivaled Ancient | ramp | Dragon ramp piece for the deck’s mana development.
- 1 Old Gnawbone | ramp | Dragon ramp piece for the deck’s mana development.
- 1 Scaled Nurturer | ramp | Early Dragon ramp for the deck.
- 1 Elemental Bond | draw | Draw support for keeping the Dragon plan supplied.
- 1 Garruk's Uprising | draw | Draw support for the deck’s creature plan.
- 1 Guardian Project | draw | Draw support for the deck’s creature plan.
- 1 Dragonborn Champion | draw | Dragon-themed draw support.
- 1 Dragon Mage | draw | Dragon-themed draw support.
- 1 Return of the Wildspeaker | draw | Draw support for the deck’s creature plan.
- 1 Rishkar's Expertise | draw | Draw support for the deck’s creature plan.
- 1 Sylvan Library | draw | Reliable draw support for the deck.
- 1 Toski, Bearer of Secrets | draw | Draw support for the deck’s combat plan.
- 1 Vanquisher's Banner | draw | Dragon-focused draw support.
- 1 Heroic Intervention | interaction | Protective interaction for the Dragon board.
- 1 Lightning Greaves | interaction | Equipment interaction that supports key creatures.
- 1 Swiftfoot Boots | interaction | Equipment interaction that supports key creatures.
- 1 Tamiyo's Safekeeping | interaction | Protective interaction for key permanents.
- 1 Veil of Summer | interaction | Low-cost interaction for the deck.
- 1 The One Ring | interaction | Versatile interaction for the deck.
- 1 Bogardan Hellkite | removal | Dragon removal for opposing problems.
- 1 Dragon Tempest | removal | Dragon-focused removal support.
- 1 Dragonlord Atarka | removal | Dragon removal for opposing problems.
- 1 Drakuseth, Maw of Flames | removal | Dragon removal for opposing problems.
- 1 Glorybringer | removal | Dragon removal for opposing problems.
- 1 Scourge of Valkas | removal | Dragon-focused removal support.
- 1 Terror of the Peaks | removal | Dragon-focused removal support.
- 1 Wrathful Red Dragon | removal | Dragon removal for opposing problems.
- 1 Balefire Dragon | wipe | Dragon board-wipe option for crowded boards.
- 1 Draconic Intervention | wipe | Board-wipe option for crowded boards.
- 1 Incinerator of the Guilty | wipe | Dragon board-wipe option for crowded boards.
- 1 Acolyte of Bahamut | synergy | Dragon-focused synergy for the deck’s main plan.
- 1 Breaching Dragonstorm | synergy | Dragon-focused synergy for the deck’s main plan.
- 1 Crucible of Fire | synergy | Dragon-focused synergy for the deck’s main plan.
- 1 Dracogenesis | synergy | Dragon-focused synergy for the deck’s main plan.
- 1 Dragon Egg | synergy | Dragon-themed synergy piece.
- 1 Dragon Hatchling | synergy | Dragon-themed synergy piece.
- 1 Dragonkin Berserker | synergy | Dragon-focused synergy for the deck’s main plan.
- 1 Dragonlord's Servant | synergy | Dragon-focused synergy for the deck’s main plan.
- 1 Dragonspeaker Shaman | synergy | Dragon-focused synergy for the deck’s main plan.
- 1 Dragonstorm | synergy | Dragon-focused synergy for the deck’s main plan.
- 1 Firespitter Whelp | synergy | Dragon-themed synergy piece.
- 1 Kargan Dragonrider | synergy | Dragon-focused synergy for the deck’s main plan.
- 1 Minion of the Mighty | synergy | Dragon-focused synergy for the deck’s main plan.
- 1 Sarkhan's Triumph | synergy | Dragon-focused synergy for the deck’s main plan.
- 1 Ancient Bronze Dragon | threat | Large Dragon threat for the combat plan.
- 1 Blast-Furnace Hellkite | threat | Dragon threat that advances the combat plan.
- 1 Hellkite Charger | threat | Dragon threat that advances the combat plan.
- 1 Hellkite Courser | threat | Dragon threat that advances the combat plan.
- 1 Lathliss, Dragon Queen | threat | Dragon threat for the deck’s main plan.
- 1 Scourge of the Throne | threat | Dragon threat that advances the combat plan.
- 1 Terror of Mount Velus | threat | Dragon threat that advances the combat plan.
- 1 Thrakkus the Butcher | threat | Dragon threat for the deck’s main plan.
- 1 Thunderbreak Regent | threat | Dragon threat for the deck’s main plan.
- 1 Twinflame Tyrant | threat | Dragon threat that advances the combat plan.
- 1 Utvara Hellkite | threat | Dragon threat for the deck’s main plan.
- 1 Stormbreath Dragon | threat | Dragon threat that advances the combat plan.

</details>

### 15. delegated commander, owned first

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 225 names.

Cards: 99 main, 0 sideboard. Repair turn: false. Block findings: 0.

Cost: $0.00 to buy, $201.50 the whole deck.

**Summary:** This deck builds around Cecil with a lifegain-focused creature and equipment shell, using its synergy cards to establish a durable board while draw and ramp keep the plan moving. It wins by protecting its key creatures, deploying its larger threats, and clearing away resistance with focused removal or board wipes. The tradeoff is that the deck is centered on developing a board, so repeated disruption can slow its pressure and make rebuilding important.

- [INFO] `bracket_prose_rules`: bracket 2 (Core, 8+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 3.06 over 63 nonland cards

<details><summary>The deck list</summary>

- 13 Plains | land | A basic land that supplies white mana.
- 11 Swamp | land | A basic land that supplies black mana.
- 1 Command Tower | land | A land slot for the commander deck’s mana base.
- 1 Path of Ancestry | land | A land slot for the creature-focused build.
- 1 Secluded Courtyard | land | A land slot supporting the deck’s creature base.
- 1 Marsh Flats | land | A land slot for mana-base consistency.
- 1 Fabled Passage | land | A land slot for mana-base consistency.
- 1 Evolving Wilds | land | A land slot for mana-base consistency.
- 1 Terramorphic Expanse | land | A land slot for mana-base consistency.
- 1 Ash Barrens | land | A land slot for mana-base consistency.
- 1 Plaza of Heroes | land | A land slot for the commander-focused build.
- 1 Castle Locthwain | land | A land slot for the black-white mana base.
- 1 Takenuma, Abandoned Mire | land | A land slot for the black-white mana base.
- 1 War Room | land | A utility land slot.
- 1 Sol Ring | ramp | An efficient shortlisted ramp piece.
- 1 Arcane Signet | ramp | A ramp piece for developing the deck’s mana.
- 1 Commander's Sphere | ramp | A ramp piece for developing the deck’s mana.
- 1 Fellwar Stone | ramp | A ramp piece for developing the deck’s mana.
- 1 Thought Vessel | ramp | A ramp piece for developing the deck’s mana.
- 1 Wayfarer's Bauble | ramp | A ramp piece for developing the deck’s mana.
- 1 Sword of the Animist | ramp | A ramp equipment for the creature-focused plan.
- 1 Relic of Legends | ramp | A ramp piece for developing the deck’s mana.
- 1 Springleaf Drum | ramp | A ramp piece that works with the deck’s creatures.
- 1 Bender's Waterskin | ramp | A shortlisted ramp slot.
- 1 Buster Sword | draw | A draw equipment for the deck’s creature plan.
- 1 Call of the Ring | draw | A shortlisted draw source.
- 1 Exemplar of Light | draw | A draw creature that fits the deck’s white creature base.
- 1 Hero in Training | draw | A draw creature for the creature-focused build.
- 1 Idol of Oblivion | draw | A compact draw artifact.
- 1 Inspiring Overseer | draw | A draw creature that fits the deck’s white creature base.
- 1 Lembas | draw | A draw artifact slot.
- 1 Mask of Memory | draw | A draw equipment for attacking creatures.
- 1 Night's Whisper | draw | A direct draw spell.
- 1 Skullclamp | draw | A draw equipment for the creature-heavy build.
- 1 Banishing Light | removal | A flexible removal slot.
- 1 Bitter Triumph | removal | A direct removal spell.
- 1 Crib Swap | removal | A creature-removal option.
- 1 Generous Gift | removal | A flexible removal spell.
- 1 Get Lost | removal | A flexible removal spell.
- 1 Infernal Grasp | removal | A direct removal spell.
- 1 Stroke of Midnight | removal | A flexible removal spell.
- 1 Swords to Plowshares | removal | An efficient creature-removal option.
- 1 Champion's Helm | interaction | An interaction equipment protecting a key creature.
- 1 Clever Concealment | interaction | An interaction spell for protecting the board.
- 1 Darksteel Plate | interaction | An interaction equipment protecting a key creature.
- 1 Lightning Greaves | interaction | An interaction equipment for protecting a key creature.
- 1 Swiftfoot Boots | interaction | An interaction equipment for protecting a key creature.
- 1 Take Up the Shield | interaction | An interaction spell for protecting a creature.
- 1 Aerith Gainsborough | synergy | A shortlisted synergy creature for the lifegain-focused build.
- 1 Adventurous Eater // Have a Bite | synergy | A shortlisted synergy card for the lifegain-focused build.
- 1 Aettir and Priwen | synergy | A synergy equipment for the creature-focused plan.
- 1 Angel of Vitality | synergy | A shortlisted synergy creature for the lifegain-focused build.
- 1 Compassionate Healer | synergy | A shortlisted synergy creature for the lifegain-focused build.
- 1 Dancer's Chakrams | synergy | A synergy equipment for the creature-focused plan.
- 1 Elixir | synergy | A shortlisted synergy artifact.
- 1 Excalibur II | synergy | A synergy equipment for the creature-focused plan.
- 1 Light of Promise | synergy | A shortlisted synergy enchantment.
- 1 Night Nurse, Healer of Heroes | synergy | A shortlisted synergy creature for the lifegain-focused build.
- 1 Rosie Cotton of South Lane | synergy | A shortlisted synergy creature for the lifegain-focused build.
- 1 The Darkness Crystal | synergy | A shortlisted synergy artifact.
- 1 Well-Worn Spatula | synergy | A synergy equipment for the creature-focused plan.
- 1 White Mage's Staff | synergy | A synergy equipment for the lifegain-focused build.
- 1 Angel of Invention | threat | A creature threat for closing games.
- 1 Dawnhand Eulogist | threat | A creature threat for closing games.
- 1 Foggy Swamp Hunters | threat | A creature threat for closing games.
- 1 Gollum, Silent Slinker // Meager Meal | threat | A creature threat for closing games.
- 1 Lyra Dawnbringer | threat | A creature threat for closing games.
- 1 Minwu, White Mage | threat | A creature threat for closing games.
- 1 Rabaroo Troop | threat | A creature threat for closing games.
- 1 Reaping Willow | threat | A creature threat for closing games.
- 1 Shattered Angel | threat | A creature threat for closing games.
- 1 Sneering Shadewriter | threat | A creature threat for closing games.
- 1 Victory's Herald | threat | A creature threat for closing games.
- 1 Lo and Li, Twin Tutors | threat | A creature threat for closing games.
- 1 Austere Command | wipe | A flexible board-wipe slot.
- 1 Fumigate | wipe | A board-wipe slot that fits the lifegain-focused plan.
- 1 Vanquish the Horde | wipe | A board-wipe slot for resetting crowded boards.

</details>

