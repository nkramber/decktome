# PR-8 deck gate

Run date: 2026-08-27. Card snapshot: 2026-08-24.

Verdict: PASS. 15 of 15 decks passed every block check, 0 invented names reached the user, and 0 summaries stated a false rule of the game. All three bars are zero tolerance.

## Summary

| Measure | Value |
|---|---|
| Prompts | 15 |
| Decks returned | 15 |
| Decks with no block finding | 15 |
| Invented names that reached the user | 0 |
| Decks that needed the repair turn | 1 |
| Summaries judged (F-26) | 15 |
| Summaries that state a rule of the game | 2 |
| Summaries that state a FALSE rule | 0 |
| Errors | 0 |
| Prompt version | 3 |
| Calls | 31 |
| Cost | $0.9002 |
| Time | 720 seconds |

## Findings by code

A block stops the deck. A warning and a note are reports.

| Code | Count |
|---|---|
| `curve_summary` | 15 |
| `bracket_prose_rules` | 9 |
| `not_owned` | 1 |
| `basics_added` | 1 |

By severity: BLOCK 0. WARN 1. INFO 25. 

## Decks

### 1. lifegain Commander, any card

Format: Commander. Theme: lifegain. Pool: any_card. Shortlist: 303 names.

Cards: 99. Repair turn: false. Block findings: 0.

**Summary:** Karlov leads a black-white lifegain deck that builds momentum through lifegain synergies, develops its resources, and keeps opponents in check with focused removal and board resets. It wins by turning its themed creatures and larger threats into sustained pressure while its lifegain payoffs reinforce that pressure. The deck trades raw speed for a board-focused, incremental game and can be more vulnerable when its key permanents do not stay in play.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 3.63 over 63 nonland cards

<details><summary>The deck list</summary>

- 18 Plains | land | Provides white mana for the deck's lifegain-focused cards.
- 18 Swamp | land | Provides black mana for the deck's lifegain-focused cards.
- 1 Alhammarret's Archive | draw | Adds card advantage support to the lifegain plan.
- 1 Archivist of Oghma | draw | Provides card advantage while fitting the deck's white core.
- 1 Cosmos Elixir | draw | Adds a resilient source of card advantage.
- 1 Dawn of Hope | draw | Supports the deck with card advantage tied to its main theme.
- 1 Enduring Innocence | draw | Provides card advantage from a lifegain-friendly permanent.
- 1 Mangara, the Diplomat | draw | Adds card advantage from a useful white creature.
- 1 Markov Purifier | draw | Provides card advantage within the deck's Vampire overlap.
- 1 Sigarda's Splendor | draw | Adds card advantage that suits the lifegain strategy.
- 1 The Gaffer | draw | Provides another lifegain-aligned card advantage piece.
- 1 Well of Lost Dreams | draw | Turns the deck's central theme into sustained card advantage.
- 1 Alseid of Life's Bounty | interaction | Helps protect the deck's important permanents.
- 1 Enduring Angel // Angelic Enforcer | interaction | Provides a protective option that fits the deck's theme.
- 1 Faith's Shield | interaction | Adds a flexible protective response.
- 1 Metropolis Reformer | interaction | Provides a defensive creature for the board.
- 1 Rune-Tail, Kitsune Ascendant // Rune-Tail's Essence | interaction | Supports the deck's defensive game plan.
- 1 Werefox Bodyguard | interaction | Adds a creature-based answer to opposing pressure.
- 1 Altar of the Pantheon | ramp | Provides mana development from an artifact.
- 1 Angel of Indemnity | ramp | Adds mana development on a lifegain-themed creature.
- 1 Crypt Ghast | ramp | Provides powerful black mana development.
- 1 Hierophant's Chalice | ramp | Adds mana development while matching the deck's theme.
- 1 Nuka-Cola Vending Machine | ramp | Provides artifact-based mana development.
- 1 Orazca Relic | ramp | Adds reliable artifact-based mana development.
- 1 Potioner's Trove | ramp | Provides additional mana development from an artifact.
- 1 Pristine Talisman | ramp | Develops mana while fitting the lifegain focus.
- 1 The Celestus | ramp | Adds mana development through a versatile artifact.
- 1 Treasure Chest | ramp | Provides another artifact-based mana source.
- 1 Ayli, Eternal Pilgrim | removal | Provides a lifegain-aligned answer to troublesome permanents.
- 1 Cavalier of Night | removal | Adds creature-based removal to the deck.
- 1 Murderous Rider // Swift End | removal | Provides a flexible black removal option.
- 1 Nightmare's Thirst | removal | Adds efficient black removal.
- 1 Noxious Gearhulk | removal | Provides a substantial creature-based removal option.
- 1 Solitude | removal | Adds a white removal option for key opposing creatures.
- 1 Vona, Butcher of Magan | removal | Provides repeatable removal on a lifegain-compatible creature.
- 1 Witch of the Moors | removal | Adds a theme-aligned removal threat.
- 1 Ajani's Pridemate | synergy | Rewards the deck for pursuing its lifegain plan.
- 1 Angel of Vitality | synergy | Supports the deck's core lifegain theme.
- 1 Angelic Accord | synergy | Provides a dedicated payoff for lifegain.
- 1 Blood Artist | synergy | Adds a black payoff that complements the deck's creature plan.
- 1 Bloodthirsty Aerialist | synergy | Rewards ongoing lifegain with a growing threat.
- 1 Cleric Class | synergy | Strengthens the deck's central lifegain theme.
- 1 Cleric of Life's Bond | synergy | Combines the deck's Cleric and lifegain themes.
- 1 Heliod, Sun-Crowned | synergy | Provides a powerful permanent for the lifegain plan.
- 1 Indulging Patrician | synergy | Adds a lifegain-focused Vampire payoff.
- 1 Righteous Valkyrie | synergy | Supports both the Angel and lifegain elements of the deck.
- 1 Sanguine Bond | synergy | Converts the deck's central theme into additional pressure.
- 1 Serra Ascendant | synergy | Provides an early creature that fits the lifegain strategy.
- 1 Vito, Thorn of the Dusk Rose | synergy | Adds a focused black lifegain payoff.
- 1 Voice of the Blessed | synergy | Rewards repeated lifegain with a scalable creature.
- 1 Archangel of Thune | threat | Provides a premier lifegain-themed battlefield threat.
- 1 Astarion, the Decadent | threat | Adds a strong black-white threat to close games.
- 1 Blood Baron of Vizkopa | threat | Provides a resilient Vampire threat.
- 1 Celestine, the Living Saint | threat | Adds a substantial white threat for the midgame.
- 1 Defiant Bloodlord | threat | Provides a large black lifegain-themed threat.
- 1 Divinity of Pride | threat | Adds a powerful threat suited to a high-life strategy.
- 1 Liesa, Forgotten Archangel | threat | Provides a major black-white Angel threat.
- 1 Lyra Dawnbringer | threat | Adds a potent Angel threat to the board.
- 1 Nykthos Paragon | threat | Provides a lifegain-linked threat for larger turns.
- 1 Rhox Faithmender | threat | Supports the deck's central theme while presenting a sturdy threat.
- 1 Tivash, Gloom Summoner | threat | Adds a black threat that fits the lifegain shell.
- 1 Valkyrie Harbinger | threat | Provides an Angel threat aligned with the deck's theme.
- 1 Fumigate | wipe | Resets crowded boards while fitting the deck's lifegain focus.
- 1 Kaya's Wrath | wipe | Provides a dependable black-white board reset.
- 1 The Battle of Bywater | wipe | Adds another white board reset for developed opposing boards.

</details>

### 2. aristocrats Commander, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 178 names.

Cards: 99. Repair turn: false. Block findings: 0.

**Summary:** This is a black-white aristocrats shell that develops a creature board, supports it with equipment and artifact pieces, and keeps resources flowing while controlling opposing threats. It wins by turning a sustained board and its sacrifice-focused synergies into steady pressure, with removal and resets available to reopen the game when needed. It gives up a concentrated finisher package and deeper sacrifice redundancy in exchange for a broad, library-first collection of mana, protection, draw, and answers.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 2.73 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Ash Barrens | land | A flexible land slot for the mana base.
- 1 Bonders' Enclave | land | A land slot that supports the deck's mana base.
- 1 Castle Locthwain | land | A black-producing land for the mana base.
- 1 City of Brass | land | Flexible color access for the mana base.
- 1 Command Tower | land | Reliable color access for the commander deck.
- 1 Escape Tunnel | land | A utility land slot in the mana base.
- 1 Evolving Wilds | land | A flexible land that helps assemble colors.
- 1 Exotic Orchard | land | Flexible color access for the mana base.
- 1 Fabled Passage | land | A flexible land that helps assemble colors.
- 1 Field of Ruin | land | A utility land slot in the mana base.
- 1 Ghost Quarter | land | A utility land slot in the mana base.
- 1 Ifnir Deadlands | land | A black-producing utility land.
- 1 Marsh Flats | land | A flexible land that helps assemble colors.
- 1 Minas Tirith | land | A white-producing land for the mana base.
- 1 Opal Palace | land | A utility land slot in the mana base.
- 1 Path of Ancestry | land | A flexible land for the creature-heavy shell.
- 1 Plaza of Heroes | land | A utility land for the legendary commander deck.
- 1 Reliquary Tower | land | A utility land slot in the mana base.
- 1 Rogue's Passage | land | A utility land that helps apply pressure.
- 1 Scavenger Grounds | land | A utility land slot in the mana base.
- 1 Secluded Courtyard | land | Flexible color access for the creature base.
- 1 Takenuma, Abandoned Mire | land | A black-producing utility land.
- 1 Terramorphic Expanse | land | A flexible land that helps assemble colors.
- 7 Plains | land | Core white mana for the deck.
- 6 Swamp | land | Core black mana for the deck.
- 1 Arcane Signet | ramp | Efficient mana acceleration for the deck's setup turns.
- 1 Astral Cornucopia | ramp | An artifact ramp piece for the mana base.
- 1 Bender's Waterskin | ramp | An artifact ramp piece for the deck.
- 1 Chromatic Lantern | ramp | Mana acceleration with flexible color support.
- 1 Commander's Sphere | ramp | A ramp piece that supports the commander deck's colors.
- 1 Deadly Dispute | ramp | Ramp that fits the aristocrats-focused game plan.
- 1 Explorer's Scope | ramp | A ramp option for the creature-based shell.
- 1 Fellwar Stone | ramp | Efficient artifact mana acceleration.
- 1 Lotho, Corrupt Shirriff | ramp | A creature-based ramp piece for the deck.
- 1 Sol Ring | ramp | Fast artifact mana acceleration.
- 1 Inherited Envelope | ramp | An artifact ramp piece for the deck.
- 1 Relic of Legends | ramp | A ramp piece that works alongside the legendary cards.
- 1 Thought Vessel | ramp | An artifact ramp piece for the deck.
- 1 Wayfarer's Bauble | ramp | Early mana acceleration for the mana base.
- 1 Ahriman | draw | A draw creature that supports ongoing resources.
- 1 Buster Sword | draw | A draw equipment for the creature-focused shell.
- 1 Call of the Ring | draw | A draw engine for sustained resources.
- 1 Cirith Ungol Patrol | draw | A draw creature that supports the board plan.
- 1 Grave Venerations | draw | A draw piece for sustained resources.
- 1 Idol of Oblivion | draw | A draw artifact that fits the aristocrats shell.
- 1 Lembas | draw | A draw artifact that supports the deck's resources.
- 1 Mask of Memory | draw | A draw equipment for creature-based pressure.
- 1 Night's Whisper | draw | Straightforward draw to keep resources moving.
- 1 Skullclamp | draw | A draw equipment that strongly fits the aristocrats plan.
- 1 Puresteel Paladin | draw | A draw creature that supports the equipment package.
- 1 Wall of Omens | draw | A draw creature that helps establish the board.
- 1 Tome of Legends | draw | A draw artifact for sustained resources.
- 1 Stone of Erech | draw | A draw artifact that supports the deck's resources.
- 1 Bastion Protector | interaction | Creature-based interaction that supports the commander.
- 1 Boromir, Warden of the Tower | interaction | A legendary interaction creature for the board.
- 1 Clever Concealment | interaction | Protective interaction for a developed board.
- 1 Gift of Immortality | interaction | Protective interaction for a key creature.
- 1 Lightning Greaves | interaction | Protective equipment for important creatures.
- 1 Swiftfoot Boots | interaction | Protective equipment for important creatures.
- 1 Darksteel Plate | interaction | Protective equipment for a key permanent.
- 1 Reprieve | interaction | Flexible instant-speed interaction.
- 1 Take Up the Shield | interaction | Protective interaction for the creature board.
- 1 Angel of Serenity | removal | A creature-based removal option.
- 1 Bitter Triumph | removal | Efficient targeted removal.
- 1 Claim the Precious | removal | Targeted removal for opposing threats.
- 1 Crib Swap | removal | Flexible creature removal.
- 1 Fiend Hunter | removal | Creature-based removal for the board plan.
- 1 Generous Gift | removal | Flexible removal for problematic permanents.
- 1 Infernal Grasp | removal | Efficient targeted removal.
- 1 Swords to Plowshares | removal | Efficient creature removal.
- 1 Dispatch | removal | Additional targeted removal.
- 1 Destroy Evil | removal | Flexible targeted removal.
- 1 Fatal Push | removal | Efficient targeted removal.
- 1 Austere Command | wipe | A flexible reset when the board gets out of hand.
- 1 Dusk // Dawn | wipe | A board-reset option for creature-heavy games.
- 1 Fumigate | wipe | A reset button against developed boards.
- 1 Black Sun's Zenith | wipe | An additional scalable board reset.
- 1 Martial Coup | wipe | A board reset that also supports pressure afterward.
- 1 Vanquish the Horde | wipe | An efficient board-reset option.
- 1 Arcade Cabinet | synergy | A colorless piece supporting the aristocrats synergy plan.
- 1 Gollum, Patient Plotter | synergy | A black synergy creature for the aristocrats shell.
- 1 Gríma Wormtongue | synergy | A black synergy creature for the deck's plan.
- 1 Heirloom Auntie | synergy | A synergy creature that supports the creature-based shell.
- 1 Joo Dee, One of Many | synergy | A white synergy creature for the deck.
- 1 Nimble Hobbit | synergy | A white synergy creature for the aristocrats plan.
- 1 Phantom Train | synergy | A colorless synergy piece for the deck's plan.
- 1 Bill the Pony | threat | A creature threat that adds pressure to the board.
- 1 Vengeful Villagers | threat | A creature threat that supports board pressure.

</details>

### 3. artifacts Commander, bracket 4

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 299 names.

Cards: 99. Repair turn: false. Block findings: 0.

**Summary:** Urza leads a dense artifact shell that develops mana, keeps cards flowing, and uses compact interaction and removal to protect its position. The deck wins by building toward its artifact threats and pressing that advantage after the table has been checked by board wipes. It gives up broad color access and leans heavily on artifacts remaining the center of its game plan.

- JUDGE [true]: "Urza leads a dense artifact shell". Urza, Lord High Artificer is a legendary creature and can legally serve as a commander, so this claim about deck construction is accurate.
- [INFO] `curve_summary`: average mana value 3.81 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Academy Ruins | land | Artifact-focused land support.
- 1 Archway of Innovation | land | Land slot for the artifact-focused mana base.
- 1 Blinkmoth Nexus | land | Artifact-oriented land support.
- 1 Buried Ruin | land | Artifact-oriented land support.
- 1 Fomori Vault | land | Land slot for the artifact-focused mana base.
- 1 Glimmervoid | land | Land slot for the artifact-focused mana base.
- 1 Hall of Tagsin | land | Land slot for the artifact-focused mana base.
- 1 Inventors' Fair | land | Artifact-oriented land support.
- 1 Mirrex | land | Land slot for the artifact-focused mana base.
- 1 Mishra's Factory | land | Artifact-oriented land support.
- 1 Mishra's Foundry | land | Artifact-oriented land support.
- 1 Mishra's Workshop | land | Artifact-oriented land support.
- 1 Otawara, Soaring City | land | Blue land support.
- 1 Phyrexia's Core | land | Artifact-oriented land support.
- 1 Power Depot | land | Artifact land support.
- 1 Roadside Reliquary | land | Land slot for the artifact-focused mana base.
- 1 Secluded Starforge | land | Land slot for the artifact-focused mana base.
- 1 Spire of Industry | land | Artifact-oriented land support.
- 1 The Monumental Facade | land | Land slot for the artifact-focused mana base.
- 1 The Mycosynth Gardens | land | Artifact-oriented land support.
- 1 Treasure Map // Treasure Cove | land | Artifact and land support in one slot.
- 1 Urza's Saga | land | Artifact-oriented land support.
- 1 Urza's Workshop | land | Artifact-oriented land support.
- 1 Uthros, Titanic Godcore | land | Land slot for the artifact-focused mana base.
- 1 Matzalantli, the Great Door // The Core | land | Artifact and land support in one slot.
- 1 The Everflowing Well // The Myriad Pools | land | Artifact and land support in one slot.
- 10 Island | land | Reliable blue basic land base.
- 1 Mox Opal | ramp | Efficient artifact ramp.
- 1 Metalworker | ramp | Artifact-focused ramp creature.
- 1 Krark-Clan Ironworks | ramp | Artifact-focused ramp piece.
- 1 Moonsnare Prototype | ramp | Early artifact ramp.
- 1 Chief Engineer | ramp | Artifact-focused ramp creature.
- 1 Grand Architect | ramp | Artifact-focused ramp creature.
- 1 Inspiring Statuary | ramp | Artifact-focused ramp support.
- 1 Tezzeret the Seeker | ramp | Planeswalker ramp support.
- 1 Karn, Living Legacy | ramp | Planeswalker ramp support.
- 1 Urza's Command | ramp | Blue ramp spell.
- 1 Thoughtcast | draw | Efficient artifact-plan card draw.
- 1 Thought Monitor | draw | Artifact creature that provides card draw.
- 1 Thirst for Knowledge | draw | Blue card draw for the artifact shell.
- 1 Sai, Master Thopterist | draw | Artifact-focused card draw creature.
- 1 Vedalken Archmage | draw | Artifact-focused card draw creature.
- 1 Reverse Engineer | draw | Artifact-plan card draw.
- 1 Riddlesmith | draw | Artifact-focused card draw creature.
- 1 One with the Machine | draw | Artifact-plan card draw.
- 1 Tezzeret, Artifice Master | draw | Planeswalker card draw support.
- 1 Transmutation Font | draw | Artifact card draw support.
- 1 Assert Authority | interaction | Blue interaction for protecting the plan.
- 1 Disruption Protocol | interaction | Blue interaction for protecting the plan.
- 1 Metallic Rebuke | interaction | Artifact-focused interaction.
- 1 Stoic Rebuttal | interaction | Blue interaction for protecting the plan.
- 1 Padeem, Consul of Innovation | interaction | Artifact-focused interaction creature.
- 1 Welding Jar | interaction | Low-cost artifact interaction.
- 1 Aether Spellbomb | removal | Flexible artifact removal.
- 1 Arcum Dagsson | removal | Artifact-focused removal creature.
- 1 Resculpt | removal | Blue removal spell.
- 1 Ravenform | removal | Blue removal spell.
- 1 Spine of Ish Sah | removal | Artifact removal option.
- 1 Portal to Phyrexia | removal | Large artifact removal piece.
- 1 Tormod's Crypt | removal | Compact artifact removal support.
- 1 Skysovereign, Consul Flagship | removal | Artifact removal threat.
- 1 Engineered Explosives | wipe | Compact artifact board wipe.
- 1 Nevinyrral's Disk | wipe | Artifact board wipe.
- 1 Oblivion Stone | wipe | Artifact board wipe.
- 1 Emry, Lurker of the Loch | synergy | Core blue artifact synergy creature.
- 1 Etherium Sculptor | synergy | Artifact-focused synergy creature.
- 1 Foundry Inspector | synergy | Artifact-focused synergy creature.
- 1 Mystic Forge | synergy | Artifact synergy engine.
- 1 Unwinding Clock | synergy | Artifact synergy support.
- 1 Voltaic Key | synergy | Low-cost artifact synergy piece.
- 1 Manifold Key | synergy | Low-cost artifact synergy piece.
- 1 Clock of Omens | synergy | Artifact synergy engine.
- 1 Transmute Artifact | synergy | Blue artifact synergy spell.
- 1 Whir of Invention | synergy | Blue artifact synergy spell.
- 1 Reshape | synergy | Blue artifact synergy spell.
- 1 Scrap Trawler | synergy | Artifact synergy creature.
- 1 Panharmonicon | synergy | Artifact synergy engine.
- 1 Mirrorworks | synergy | Artifact synergy engine.
- 1 Kappa Cannoneer | threat | Artifact creature threat for closing games.
- 1 Cyberdrive Awakener | threat | Artifact creature threat for closing games.
- 1 Kuldotha Forgemaster | threat | Artifact creature threat.
- 1 Master Transmuter | threat | Artifact creature threat.
- 1 Metalwork Colossus | threat | Large artifact creature threat.
- 1 Myr Enforcer | threat | Artifact creature threat.
- 1 Traxos, Scourge of Kroog | threat | Legendary artifact creature threat.
- 1 Karn, Scion of Urza | threat | Planeswalker threat.
- 1 Broodstar | threat | Artifact-focused creature threat.
- 1 Darksteel Juggernaut | threat | Artifact creature threat.
- 1 Lodestone Golem | threat | Artifact creature threat.
- 1 Phyrexian Metamorph | threat | Artifact creature threat.

</details>

### 4. dinosaur tribal, bracket 2

Format: Commander. Theme: dinosaurs. Pool: any_card. Shortlist: 267 names.

Cards: 99. Repair turn: false. Block findings: 0.

**Summary:** This deck develops its mana, builds a Dinosaur board, and uses Gishath, Sun's Avatar alongside a lineup of large Dinosaur threats to pressure the table through combat. It has support cards to keep the creature plan moving, answers for troublesome opposing pieces, and a few reset buttons when the board gets out of hand. It gives up some speed and flexibility in exchange for a direct, creature-first plan that is easy to follow: build up, protect the important creatures, and attack with dinosaurs.

- [INFO] `bracket_prose_rules`: bracket 2 (Core, 8+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 4.16 over 63 nonland cards

<details><summary>The deck list</summary>

- 8 Forest | land | Provides basic land slots in the mana base.
- 6 Plains | land | Provides basic land slots in the mana base.
- 6 Mountain | land | Provides basic land slots in the mana base.
- 1 Battlefield Forge | land | Supports the mana base.
- 1 Canopy Vista | land | Supports the mana base.
- 1 Cinder Glade | land | Supports the mana base.
- 1 Clifftop Retreat | land | Supports the mana base.
- 1 Command Tower | land | Supports the mana base.
- 1 Evolving Wilds | land | Supports the mana base.
- 1 Exotic Orchard | land | Supports the mana base.
- 1 Myriad Landscape | land | Supports the mana base.
- 1 Path of Ancestry | land | Supports the mana base.
- 1 Rootbound Crag | land | Supports the mana base.
- 1 Sacred Foundry | land | Supports the mana base.
- 1 Spectator Seating | land | Supports the mana base.
- 1 Stomping Ground | land | Supports the mana base.
- 1 Sunpetal Grove | land | Supports the mana base.
- 1 Temple Garden | land | Supports the mana base.
- 1 Terramorphic Expanse | land | Supports the mana base.
- 1 Arcane Signet | ramp | Provides early mana development.
- 1 Atzocan Seer | ramp | Provides mana development in a creature slot.
- 1 Birds of Paradise | ramp | Provides early mana development.
- 1 Cultivate | ramp | Provides reliable mana development.
- 1 Farseek | ramp | Provides reliable mana development.
- 1 Hulking Raptor | ramp | Adds mana development while fitting the Dinosaur theme.
- 1 Kodama's Reach | ramp | Provides reliable mana development.
- 1 Nature's Lore | ramp | Provides reliable mana development.
- 1 Rampant Growth | ramp | Provides reliable mana development.
- 1 Sol Ring | ramp | Provides early mana development.
- 1 Beast Whisperer | draw | Provides card flow alongside the creature plan.
- 1 Garruk's Uprising | draw | Provides card flow for the creature-focused deck.
- 1 Guardian Project | draw | Provides ongoing card flow.
- 1 Harmonize | draw | Provides straightforward card flow.
- 1 Return of the Wildspeaker | draw | Provides card flow for the creature plan.
- 1 Ripjaw Raptor | draw | Provides card flow in a Dinosaur slot.
- 1 Rishkar's Expertise | draw | Provides a strong card-flow option.
- 1 Shamanic Revelation | draw | Provides card flow for a developed board.
- 1 Vanquisher's Banner | draw | Provides card flow for the Dinosaur theme.
- 1 Vaultborn Tyrant | draw | Provides card flow in a Dinosaur threat slot.
- 1 Akroma's Will | interaction | Provides a versatile interaction option.
- 1 Boros Charm | interaction | Provides a versatile interaction option.
- 1 Heroic Intervention | interaction | Provides an interaction option for the creature plan.
- 1 Inspiring Call | interaction | Provides interaction that fits the creature plan.
- 1 Lightning Greaves | interaction | Provides an interaction option for key creatures.
- 1 Swiftfoot Boots | interaction | Provides an interaction option for key creatures.
- 1 Apex Altisaur | removal | Provides removal in a Dinosaur slot.
- 1 Bronzebeak Foragers | removal | Provides removal in a Dinosaur slot.
- 1 Burning Sun's Avatar | removal | Provides removal in a Dinosaur slot.
- 1 Itzquinth, Firstborn of Gishath | removal | Provides removal while fitting the Dinosaur theme.
- 1 Ravenous Sailback | removal | Provides removal in a Dinosaur slot.
- 1 Savage Stomp | removal | Provides a direct removal option.
- 1 Thrashing Brontodon | removal | Provides removal in a Dinosaur slot.
- 1 Trumpeting Carnosaur | removal | Provides removal in a Dinosaur slot.
- 1 Austere Command | wipe | Provides a flexible board-reset option.
- 1 Blasphemous Act | wipe | Provides a board-reset option.
- 1 Wakening Sun's Avatar | wipe | Provides a Dinosaur-themed board-reset option.
- 1 Armored Kincaller | synergy | Supports the Dinosaur theme.
- 1 Belligerent Yearling | synergy | Supports the Dinosaur theme.
- 1 Commune with Dinosaurs | synergy | Supports the Dinosaur-focused plan.
- 1 Dinosaur Stampede | synergy | Supports the Dinosaur-focused plan.
- 1 Dromosaur | synergy | Supports the Dinosaur theme.
- 1 Huatli's Raptor | synergy | Supports the Dinosaur theme.
- 1 Hunting Velociraptor | synergy | Supports the Dinosaur theme.
- 1 Invasion of Ikoria // Zilortha, Apex of Ikoria | synergy | Supports the Dinosaur-focused plan.
- 1 Invasion of Ixalan // Belligerent Regisaur | synergy | Supports the Dinosaur-focused plan.
- 1 Kinjalli's Caller | synergy | Supports the Dinosaur-focused plan.
- 1 Kinjalli's Sunwing | synergy | Supports the Dinosaur theme.
- 1 Marauding Raptor | synergy | Supports the Dinosaur theme.
- 1 Otepec Huntmaster | synergy | Supports the Dinosaur-focused plan.
- 1 Pugnacious Hammerskull | synergy | Supports the Dinosaur theme.
- 1 Carnage Tyrant | threat | Provides a major Dinosaur threat.
- 1 Etali, Primal Storm | threat | Provides a major Dinosaur threat.
- 1 Ghalta and Mavren | threat | Provides a major Dinosaur threat.
- 1 Ghalta, Primal Hunger | threat | Provides a major Dinosaur threat.
- 1 Ghalta, Stampede Tyrant | threat | Provides a major Dinosaur threat.
- 1 Goring Ceratops | threat | Provides a Dinosaur threat for creature pressure.
- 1 Pantlaza, Sun-Favored | threat | Provides a Dinosaur threat for the theme.
- 1 Polyraptor | threat | Provides a major Dinosaur threat.
- 1 Quartzwood Crasher | threat | Provides a Dinosaur threat for creature pressure.
- 1 Regisaur Alpha | threat | Provides a Dinosaur threat for the theme.
- 1 Tyrranax Rex | threat | Provides a major Dinosaur threat.
- 1 Zetalpa, Primal Dawn | threat | Provides a major Dinosaur threat.

</details>

### 5. blink Commander, owned first

Format: Commander. Theme: blink. Pool: owned_first. Shortlist: 168 names.

Cards: 99. Repair turn: false. Block findings: 0.

**Summary:** Gilraen leads a white, creature-focused blink deck that develops its mana, builds a board of Angels, Humans, and artifact creatures, and keeps returning key members of that lineup to the battlefield. It wins by maintaining steady board pressure while using removal and resets to clear a path for its creatures. The tradeoff is a focused, permanent-heavy approach that leans on establishing creatures and artifacts rather than broad multicolor options.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 2.98 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Abandoned Air Temple | land | A land slot that supports the white mana base.
- 1 Ash Barrens | land | A land slot that helps keep the mana base consistent.
- 1 Bonders' Enclave | land | A land slot in the mana base.
- 1 Command Tower | land | A land slot for dependable access to the commander's color.
- 1 Evolving Wilds | land | A land slot that works well alongside a large Plains base.
- 1 Fabled Passage | land | A land slot that works well alongside a large Plains base.
- 1 Fountainport | land | A land slot in the mana base.
- 1 Ghost Quarter | land | A utility land slot in the mana base.
- 1 Great Hall of the Citadel | land | A land slot in the mana base.
- 1 Minas Tirith | land | A land slot in the mana base.
- 1 Path of Ancestry | land | A land slot for the creature-heavy lineup.
- 18 Plains | land | The core basic land for a consistent white mana base.
- 1 Plaza of Heroes | land | A land slot that fits the legendary commander plan.
- 1 Reliquary Tower | land | A utility land slot in the mana base.
- 1 Rogue's Passage | land | A utility land slot that helps the creature plan press through.
- 1 Scavenger Grounds | land | A utility land slot in the mana base.
- 1 Secluded Courtyard | land | A land slot suited to the creature-heavy lineup.
- 1 War Room | land | A utility land slot in the mana base.
- 1 Windbrisk Heights | land | A land slot in the mana base.
- 1 Arcane Signet | ramp | A reliable early ramp piece.
- 1 Bender's Waterskin | ramp | An artifact ramp piece for developing the board.
- 1 Commander's Sphere | ramp | A ramp piece that fits the artifact package.
- 1 Fellwar Stone | ramp | An efficient artifact ramp piece.
- 1 Relic of Legends | ramp | A ramp piece for the creature-focused plan.
- 1 Sol Ring | ramp | A fast artifact ramp piece.
- 1 Sword of the Animist | ramp | A ramp equipment that gives the creature lineup another use.
- 1 Thought Vessel | ramp | An artifact ramp piece for developing the board.
- 1 Wayfarer's Bauble | ramp | A ramp piece that supports the Plains base.
- 1 White Lotus Tile | ramp | An artifact ramp piece for the deck's early development.
- 1 Adventurer's Airship | draw | An artifact included for card flow.
- 1 Buster Sword | draw | An equipment slot that supplies card flow.
- 1 Crown of Gondor | draw | A legendary artifact included for card flow.
- 1 Diary of Dreams | draw | An artifact included for card flow.
- 1 Idol of Oblivion | draw | An artifact included for card flow.
- 1 Instant Ramen | draw | An artifact included for card flow.
- 1 Lembas | draw | An artifact included for card flow.
- 1 Mask of Memory | draw | An equipment slot that supplies card flow.
- 1 Mirror of Galadriel | draw | A legendary artifact included for card flow.
- 1 Tome of Legends | draw | An artifact included for card flow.
- 1 Bastion Protector | interaction | A creature-based interaction piece for protecting the central plan.
- 1 Clever Concealment | interaction | A flexible interaction spell for safeguarding the board.
- 1 Darksteel Plate | interaction | An equipment interaction piece for the creature plan.
- 1 Lightning Greaves | interaction | An equipment interaction piece for the commander and creatures.
- 1 Swiftfoot Boots | interaction | An equipment interaction piece for the commander and creatures.
- 1 Unbreakable Formation | interaction | A board-focused interaction spell.
- 1 Banishing Light | removal | A removal piece for dealing with opposing permanents.
- 1 Crib Swap | removal | A removal spell for opposing creatures.
- 1 Dispatch | removal | An efficient removal spell that fits the artifact count.
- 1 Generous Gift | removal | A flexible removal spell for troublesome permanents.
- 1 Get Lost | removal | A removal spell for troublesome permanents.
- 1 Journey to Nowhere | removal | A removal enchantment for opposing creatures.
- 1 Stroke of Midnight | removal | A flexible removal spell for troublesome permanents.
- 1 Swords to Plowshares | removal | An efficient removal spell for opposing creatures.
- 1 Austere Command | wipe | A flexible reset for boards that get out of hand.
- 1 Fumigate | wipe | A creature-board reset when pressure needs answering.
- 1 Vanquish the Horde | wipe | A board reset for creature-heavy opposing positions.
- 1 Angel of Condemnation | synergy | A creature centerpiece for the blink-focused plan.
- 1 Angel of Sanctions | synergy | A creature piece that rewards returning to the battlefield.
- 1 Angel of Serenity | synergy | A high-impact creature for the blink-focused lineup.
- 1 Champions of Minas Tirith | synergy | A creature piece in the blink-focused lineup.
- 1 Ennis, Debate Moderator | synergy | A listed synergy piece for the deck's central plan.
- 1 Fiend Hunter | synergy | A creature piece that fits repeated battlefield returns.
- 1 Flickerwisp | synergy | A listed synergy creature for the blink plan.
- 1 Inspiring Overseer | synergy | A creature piece that fits repeated battlefield returns.
- 1 Jocasta, Automaton Avenger | synergy | A listed synergy piece for the deck's central plan.
- 1 Joined Researchers // Secret Rendezvous | synergy | A creature piece in the blink-focused lineup.
- 1 Palace Jailer | synergy | A creature piece that fits repeated battlefield returns.
- 1 Personify | synergy | A listed synergy spell for the blink plan.
- 1 Slip On the Ring | synergy | A spell that directly supports the blink-focused plan.
- 1 Wall of Omens | synergy | A creature piece that fits repeated battlefield returns.
- 1 Boromir, Warden of the Tower | threat | A legendary creature that adds meaningful board pressure.
- 1 Bronze Guardian | threat | An artifact creature that adds board pressure.
- 1 Exemplar of Light | threat | An Angel creature that adds board pressure.
- 1 Faramir, Field Commander | threat | A legendary creature that adds board pressure.
- 1 Frontline Medic | threat | A creature that adds to the deck's board presence.
- 1 Giada, Font of Hope | threat | A legendary Angel that adds board pressure.
- 1 Kataki, War's Wage | threat | A legendary creature that adds board pressure.
- 1 Puresteel Paladin | threat | A creature that adds board pressure alongside the equipment package.
- 1 The Vision | threat | A legendary artifact creature that adds board pressure.
- 1 The Walls of Ba Sing Se | threat | A legendary artifact creature that adds board presence.
- 1 Westfold Rider | threat | A creature that adds to the deck's board presence.
- 1 Zack Fair | threat | A legendary creature that adds board pressure.

</details>

### 6. Modern tempo, tournament

Format: Modern. Theme: tempo. Pool: any_card. Shortlist: 167 names.

Cards: 60. Repair turn: false. Block findings: 0.

**Summary:** This blue-red tempo shell uses inexpensive draw, removal, and interaction to keep exchanges efficient while setting up its Temporal Mastery synergy as the closing plan. It aims to win after its answers have bought enough time to turn that synergy into a decisive advantage. It gives up a dedicated creature-threat package, so it relies heavily on drawing the right spells and sequencing its interaction well.

- [INFO] `curve_summary`: average mana value 1.89 over 36 nonland cards

<details><summary>The deck list</summary>

- 10 Island | land | Basic blue land for the deck's mana base.
- 6 Mountain | land | Basic red land for the deck's mana base.
- 4 Steam Vents | land | Blue-red land supporting both main colors.
- 2 Shivan Reef | land | Blue-red land supporting the tempo shell.
- 2 Stormcarved Coast | land | Blue-red land that rounds out the mana base.
- 4 Consider | draw | Cheap card selection for finding the needed spell.
- 4 Preordain | draw | Card selection that helps shape early turns.
- 2 Opt | draw | Additional inexpensive draw for the spell-heavy plan.
- 4 Counterspell | interaction | Core interaction for protecting the deck's plan.
- 4 Spell Pierce | interaction | Low-cost interaction for tempo exchanges.
- 2 Force of Negation | interaction | Additional interaction for important opposing plays.
- 4 Lightning Bolt | removal | Efficient removal for clearing obstacles.
- 4 Pongify | removal | Cheap removal that broadens the deck's answers.
- 4 Rapid Hybridization | removal | Additional low-cost removal for creature matchups.
- 4 Temporal Mastery | synergy | The central temporal-synergy card for the deck's closing plan.

</details>

### 7. Modern burn, casual

Format: Modern. Theme: burn. Pool: any_card. Shortlist: 290 names.

Cards: 60. Repair turn: false. Block findings: 0.

**Summary:** This mono-red burn deck applies direct pressure with its removal suite while its spell-focused synergy cards and creatures keep that pressure coming. It wins by pushing damage through consistently and backing that plan with a few sturdier threats when the game lasts longer. It gives up broader answers and a flexible multicolor game plan in exchange for a focused, aggressive approach.

- [INFO] `curve_summary`: average mana value 2.72 over 36 nonland cards

<details><summary>The deck list</summary>

- 20 Mountain | land | Basic land base for a mono-red deck.
- 4 Ramunap Ruins | land | Additional land slots that stay within the mono-red plan.
- 4 Risk Factor | draw | Draw support that helps keep the burn plan supplied.
- 2 Grab the Prize | draw | Additional draw support for maintaining pressure.
- 2 Sudden Breakthrough | ramp | Provides the deck's ramp slots.
- 4 Lightning Bolt | removal | Efficient removal for the direct-damage plan.
- 2 Lava Dart | removal | Additional low-cost removal to support burn pressure.
- 4 Kessig Flamebreather | synergy | A synergy piece for a spell-heavy burn shell.
- 4 Thermo-Alchemist | synergy | A synergy piece that supports repeated burn pressure.
- 4 Ashcloud Phoenix | threat | A threat that gives the deck a resilient creature presence.
- 4 Fuming Effigy | threat | A threat that adds to the deck's pressure package.
- 4 Sunspine Lynx | threat | A threat for continuing the deck's aggressive plan.
- 2 Tectonic Giant | threat | A larger threat to round out the creature package.

</details>

### 8. Modern lifegain, FNM

Format: Modern. Theme: lifegain. Pool: any_card. Shortlist: 301 names.

Cards: 60. Repair turn: false. Block findings: 0.

**Summary:** This white-black lifegain deck develops early synergy creatures, turns repeated life gain into larger threats, and pressures the opponent with Angels and Vampires. It wins primarily through creature combat backed by removal, while giving up some speed and broad flexibility for a focused, permanent-heavy game plan.

- [INFO] `curve_summary`: average mana value 3.64 over 36 nonland cards

<details><summary>The deck list</summary>

- 8 Plains | land | Core basic land for the white half of the mana base.
- 8 Swamp | land | Core basic land for the black half of the mana base.
- 4 Scoured Barrens | land | Lifegain-themed dual land for the mana base.
- 2 Shambling Vent | land | Additional white-black land for the mana base.
- 2 Restless Fortress | land | Additional white-black land that also fits the creature-focused plan.
- 4 Inspiring Overseer | draw | Creature-based card draw that fits the Angel portion of the deck.
- 2 Well of Lost Dreams | draw | Dedicated draw support for the lifegain plan.
- 2 Orzhov Keyrune | ramp | Mana acceleration in the deck's colors.
- 4 Solitude | removal | Primary removal for opposing threats.
- 2 Murderous Rider // Swift End | removal | Flexible removal that leaves behind a creature.
- 4 Soul Warden | synergy | Core lifegain synergy creature.
- 4 Ajani's Pridemate | synergy | Primary payoff for the deck's lifegain synergies.
- 4 Archangel of Thune | threat | High-impact Angel threat that supports the lifegain theme.
- 4 Blood Baron of Vizkopa | threat | Reliable Vampire threat for combat pressure.
- 3 Bloodthirsty Conqueror | threat | Powerful Vampire threat for closing games.
- 3 Cliffhaven Vampire | threat | Lifegain-themed Vampire threat that advances the deck's plan.

</details>

### 9. Standard midrange, FNM

Format: Standard. Theme: midrange. Pool: any_card. Shortlist: 164 names.

Cards: 60. Repair turn: false. Block findings: 0.

**Summary:** This black-green midrange deck develops its mana, trades resources with removal, and uses a broad creature suite to press its advantage once the board has thinned out. It wins through sustained creature pressure backed by recurring card flow and larger top-end threats, while giving up specialized main-deck answers in favor of a more flexible, general-purpose game plan.

- [INFO] `curve_summary`: average mana value 3.47 over 34 nonland cards

<details><summary>The deck list</summary>

- 10 Forest | land | Basic green land for the black-green mana base.
- 10 Swamp | land | Basic black land for the black-green mana base.
- 1 Blooming Marsh | land | Black-green land for the mana base.
- 1 Deathcap Glade | land | Black-green land for the mana base.
- 1 Overgrown Tomb | land | Black-green land for the mana base.
- 1 Underground Mortuary | land | Black-green land for the mana base.
- 2 Llanowar Elves | ramp | Early resource development from a listed ramp card.
- 2 Bite Down | removal | Efficient listed removal for clearing opposing threats.
- 2 Bitter Triumph | removal | Listed removal that helps the deck trade resources.
- 2 Nowhere to Run | removal | Listed removal for the main-deck answer suite.
- 2 Darkstar Augur | draw | Creature-based listed draw for ongoing resources.
- 2 Unholy Annex // Ritual Chamber | draw | Listed draw that supports the deck's longer game.
- 2 Phyrexian Arena | draw | Listed draw for sustained card flow.
- 2 Elvish Archdruid | synergy | Creature-based support piece; it is listed as ramp.
- 2 Enduring Vitality | synergy | Resource-support piece; it is listed as ramp.
- 2 Not Dead After All | synergy | Interaction-based support for the creature plan.
- 2 Snakeskin Veil | synergy | Interaction-based support for the creature plan.
- 2 Goldvein Hydra | threat | Creature pressure with the added value of its listed ramp role.
- 2 Vaultborn Tyrant | threat | Large creature pressure with a listed draw role.
- 2 Vein Ripper | threat | Creature pressure that also occupies a listed removal role.
- 2 Ojer Kaslem, Deepest Growth // Temple of Cultivation | threat | Creature pressure with a land-facing secondary identity.
- 2 Zenos yae Galvus // Shinryu, Transcendent Rival | threat | Creature pressure with a listed wipe role.
- 2 Zodiark, Umbral God | threat | Creature pressure with a listed wipe role.
- 2 Massacre Wurm | threat | Creature pressure with a listed wipe role.

</details>

### 10. Standard aggro, tournament

Format: Standard. Theme: aggro. Pool: any_card. Shortlist: 294 names.

Cards: 60. Repair turn: false. Block findings: 0.

**Summary:** This deck takes a proactive red-white combat plan, using a dense threat base backed by synergistic support to pressure opponents early and keep attacking. It wins by maintaining that pressure while clearing obstacles and protecting its momentum with interaction. The tradeoff is a narrow, board-focused approach: it is built to be assertive rather than to cover every kind of opposing strategy in the first game.

- [INFO] `curve_summary`: average mana value 3.06 over 36 nonland cards

<details><summary>The deck list</summary>

- 8 Plains | land | Provides a white land base for the deck.
- 8 Mountain | land | Provides a red land base for the deck.
- 4 Sacred Foundry | land | Provides red-white land support.
- 2 Inspiring Vantage | land | Provides red-white land support.
- 2 Elegant Parlor | land | Provides red-white land support.
- 4 Fugitive Codebreaker | draw | Provides the deck's red draw package.
- 2 Inspiring Overseer | draw | Adds white draw support while contributing to the creature plan.
- 4 Boros Charm | interaction | Provides red-white interaction for a proactive deck.
- 2 Sheltered by Ghosts | interaction | Adds interaction that fits the white portion of the deck.
- 4 Harsh Annotation | removal | Provides efficient-looking red removal support.
- 4 Case of the Gateway Express | removal | Adds removal options in the white portion of the deck.
- 4 Warleader's Call | synergy | Supports the deck's red-white aggressive synergy plan.
- 4 Bedhead Beastie | threat | Supplies a full set of threats for the combat plan.
- 4 Dragonback Lancer | threat | Adds creature threats that suit a red-white attack-focused deck.
- 4 Redcap Gutter-Dweller | threat | Rounds out the proactive threat base.

</details>

### 11. Commander with a locked card

Format: Commander. Theme: sacrifice. Pool: any_card. Shortlist: 296 names.

Cards: 99. Repair turn: false. Block findings: 0.

**Summary:** Karlov leads a creature-heavy sacrifice deck that turns expendable bodies into mana, cards, disruption, and steady payoff pressure. The deck aims to build a sacrifice engine, use removal to keep opposing boards contained, and close games with its larger creature threats and death-based pressure. It gives up some flexibility for a concentrated plan that wants creatures, sacrifice outlets, and payoff pieces working together.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 3.29 over 63 nonland cards

<details><summary>The deck list</summary>

- 12 Plains | land | Basic white mana for the deck.
- 16 Swamp | land | Basic black mana for the deck.
- 1 Command Tower | land | Color-fixing land for Karlov's deck.
- 1 Bojuka Bog | land | Utility land slot.
- 1 Castle Doom | land | Black mana land slot.
- 1 Diamond Valley | land | Sacrifice-focused utility land.
- 1 High Market | land | Sacrifice-focused utility land.
- 1 Miren, the Moaning Well | land | Sacrifice-focused utility land.
- 1 Phyrexian Tower | land | Sacrifice-focused utility land.
- 1 Westvale Abbey // Ormendahl, Profane Prince | land | Land slot that supports the sacrifice plan.
- 1 Sol Ring | ramp | Required inclusion that supplies ramp.
- 1 Ashnod's Altar | ramp | Sacrifice-plan ramp piece.
- 1 Phyrexian Altar | ramp | Sacrifice-plan ramp piece.
- 1 Pitiless Plunderer | ramp | Sacrifice-plan ramp piece.
- 1 Pawn of Ulamog | ramp | Sacrifice-plan ramp piece.
- 1 Sifter of Skulls | ramp | Sacrifice-plan ramp piece.
- 1 Priest of Forgotten Gods | ramp | Sacrifice-plan ramp piece.
- 1 Culling the Weak | ramp | Fast ramp for the sacrifice plan.
- 1 Crowded Crypt | ramp | Sacrifice-plan ramp piece.
- 1 Warren Soultrader | ramp | Sacrifice-plan ramp piece.
- 1 Corrupted Conviction | draw | Sacrifice-plan card draw.
- 1 Disciple of Bolas | draw | Sacrifice-plan card draw.
- 1 Village Rites | draw | Sacrifice-plan card draw.
- 1 Vampiric Rites | draw | Repeatable sacrifice-plan card draw.
- 1 Smothering Abomination | draw | Sacrifice-plan card draw.
- 1 Relic Vial | draw | Card draw for the creature-heavy plan.
- 1 Shadowheart, Dark Justiciar | draw | Sacrifice-plan card draw.
- 1 Tevesh Szat, Doom of Fools | draw | Card draw for the sacrifice plan.
- 1 Baron Bertram Graywater | draw | Card draw for the creature plan.
- 1 Bushmeat Poacher | draw | Sacrifice-plan card draw.
- 1 Cartel Aristocrat | interaction | Interaction that fits the sacrifice plan.
- 1 Fanatical Devotion | interaction | Protective interaction for key creatures.
- 1 Dark Privilege | interaction | Protective interaction for the sacrifice plan.
- 1 Promise of Tomorrow | interaction | Interaction that supports sacrificing creatures.
- 1 Gift of Doom | interaction | Protective interaction for a key permanent.
- 1 Nightmare Shepherd | interaction | Creature-based interaction for the plan.
- 1 Attrition | removal | Sacrifice-based removal.
- 1 Ayli, Eternal Pilgrim | removal | Removal that fits the sacrifice plan.
- 1 Bone Shards | removal | Low-cost removal for the sacrifice plan.
- 1 Eaten Alive | removal | Removal that fits the sacrifice plan.
- 1 Grave Pact | removal | Sacrifice-based removal pressure.
- 1 Dictate of Erebos | removal | Sacrifice-based removal pressure.
- 1 Yawgmoth, Thran Physician | removal | Creature-based removal for the plan.
- 1 Teysa, Orzhov Scion | removal | Removal that supports sacrificing creatures.
- 1 Altar of Dementia | synergy | Sacrifice outlet and plan synergy.
- 1 Bartolomé del Presidio | synergy | Creature-based sacrifice synergy.
- 1 Carrion Feeder | synergy | Creature-based sacrifice synergy.
- 1 Bastion of Remembrance | synergy | Payoff for the sacrifice plan.
- 1 Elas il-Kor, Sadistic Pilgrim | synergy | Payoff for the sacrifice plan.
- 1 Viscera Seer | synergy | Creature-based sacrifice synergy.
- 1 Woe Strider | synergy | Creature-based sacrifice synergy.
- 1 Fleshtaker | synergy | Payoff for sacrificing creatures.
- 1 Bloodflow Connoisseur | synergy | Creature-based sacrifice synergy.
- 1 Nantuko Husk | synergy | Creature-based sacrifice synergy.
- 1 Open the Graves | synergy | Payoff for the sacrifice plan.
- 1 Hidden Stockpile | synergy | Sacrifice-plan synergy piece.
- 1 Victimize | synergy | Sacrifice-plan recursion synergy.
- 1 Zulaport Cutthroat | synergy | Payoff for the sacrifice plan.
- 1 Razaketh, the Foulblooded | threat | Powerful top-end threat for the sacrifice plan.
- 1 Mondrak, Glory Dominus | threat | Threat that supports the creature plan.
- 1 Ratadrabik of Urborg | threat | Threat for the creature-heavy plan.
- 1 Liesa, Forgotten Archangel | threat | Resilient threat for the deck.
- 1 Requiem Angel | threat | Creature-focused threat.
- 1 Ghoulcaller Gisa | threat | Sacrifice-focused threat.
- 1 Sidisi, Undead Vizier | threat | Threat that fits the sacrifice plan.
- 1 Felisa, Fang of Silverquill | threat | Threat for the creature plan.
- 1 Vindictive Vampire | threat | Threat that supports creature sacrifices.
- 1 Titan Hunter | threat | Creature threat for the deck.
- 1 Corpse Harvester | threat | Creature threat for the sacrifice plan.
- 1 Threefold Thunderhulk | threat | Artifact creature threat.
- 1 Toxic Deluge | wipe | Flexible reset for crowded boards.
- 1 The Meathook Massacre | wipe | Board reset that suits the sacrifice plan.
- 1 Liliana, Dreadhorde General | wipe | Board reset attached to a lasting threat.

</details>

### 12. Commander on a budget

Format: Commander. Theme: tokens. Pool: any_card. Shortlist: 296 names.

Cards: 99. Repair turn: false. Block findings: 0.

**Summary:** This Adeline-led deck builds a token-focused board, backs that board with synergistic permanents and threats, and aims to turn sustained combat pressure into a win. It has draw and ramp to keep deploying its plan, alongside removal and broad answers to clear resistance. It gives up access to multicolor tools in exchange for a focused white token strategy.

- [INFO] `bracket_prose_rules`: bracket 2 (Core, 8+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 3.27 over 63 nonland cards

<details><summary>The deck list</summary>

- 36 Plains | land | Provides the deck's white mana base.
- 1 Bennie Bracks, Zoologist | draw | Included as a draw option.
- 1 Bygone Bishop | draw | Included as a draw option.
- 1 Chivalric Alliance | draw | Included as a draw option.
- 1 Court of Grace | draw | Included as a draw option.
- 1 Dawn of Hope | draw | Included as a draw option.
- 1 Faramir, Field Commander | draw | Included as a draw option.
- 1 Hoofprints of the Stag | draw | Included as a draw option.
- 1 Idol of Oblivion | draw | Included as a draw option.
- 1 Staff of the Storyteller | draw | Included as a draw option.
- 1 Wedding Announcement // Wedding Festivity | draw | Included as a draw option.
- 1 Ainok Strike Leader | interaction | Included to interact with opposing plans.
- 1 Blessed Sanctuary | interaction | Included to interact with opposing plans.
- 1 Rootborn Defenses | interaction | Included to interact with opposing plans.
- 1 Spirit Bonds | interaction | Included to interact with opposing plans.
- 1 Squad Commander | interaction | Included to interact with opposing plans.
- 1 Teyo, the Shieldmage | interaction | Included to interact with opposing plans.
- 1 Battle Angels of Tyr | ramp | Included as a ramp option.
- 1 Charisma Bobblehead | ramp | Included as a ramp option.
- 1 Collector's Vault | ramp | Included as a ramp option.
- 1 Currency Converter | ramp | Included as a ramp option.
- 1 Druidic Satchel | ramp | Included as a ramp option.
- 1 Goldvein Pick | ramp | Included as a ramp option.
- 1 Keeper of the Accord | ramp | Included as a ramp option.
- 1 Monologue Tax | ramp | Included as a ramp option.
- 1 Nuka-Cola Vending Machine | ramp | Included as a ramp option.
- 1 The Restoration of Eiganjo // Architect of Restoration | ramp | Included as a ramp option.
- 1 Aerial Assault | removal | Included as targeted removal.
- 1 Banishing Slash | removal | Included as targeted removal.
- 1 Citizen's Crowbar | removal | Included as targeted removal.
- 1 Hanged Executioner | removal | Included as targeted removal.
- 1 Kellan's Lightblades | removal | Included as targeted removal.
- 1 Release to Memory | removal | Included as targeted removal.
- 1 The Wandering Emperor | removal | Included as targeted removal.
- 1 Trostani's Judgment | removal | Included as targeted removal.
- 1 Animation Module | synergy | Supports the token-focused plan.
- 1 Anointer Priest | synergy | Supports the token-focused plan.
- 1 Clarion Spirit | synergy | Supports the token-focused plan.
- 1 Divine Visitation | synergy | Supports the token-focused plan.
- 1 Felidar Retreat | synergy | Supports the token-focused plan.
- 1 Horn of Gondor | synergy | Supports the token-focused plan.
- 1 Intangible Virtue | synergy | Supports the token-focused plan.
- 1 Oketra's Monument | synergy | Supports the token-focused plan.
- 1 Prava of the Steel Legion | synergy | Supports the token-focused plan.
- 1 Retrofitter Foundry | synergy | Supports the token-focused plan.
- 1 Sigiled Sword of Valeron | synergy | Supports the token-focused plan.
- 1 Skrelv's Hive | synergy | Supports the token-focused plan.
- 1 Spawning Pit | synergy | Supports the token-focused plan.
- 1 Worthy Knight | synergy | Supports the token-focused plan.
- 1 Archangel Elspeth | threat | Provides a substantial threat.
- 1 Archon of Sun's Grace | threat | Provides a substantial threat.
- 1 Basri's Lieutenant | threat | Provides a substantial threat.
- 1 Cemetery Protector | threat | Provides a substantial threat.
- 1 Emeria Angel | threat | Provides a substantial threat.
- 1 Gideon, Ally of Zendikar | threat | Provides a substantial threat.
- 1 Hero of Bladehold | threat | Provides a substantial threat.
- 1 Illustrious Wanderglyph | threat | Provides a substantial threat.
- 1 Oketra the True | threat | Provides a substantial threat.
- 1 Phantom General | threat | Provides a substantial threat.
- 1 Requiem Angel | threat | Provides a substantial threat.
- 1 Threefold Thunderhulk | threat | Provides a substantial threat.
- 1 Hour of Reckoning | wipe | Included to reset crowded boards.
- 1 Phyrexian Rebirth | wipe | Included to reset crowded boards.
- 1 The Battle of Bywater | wipe | Included to reset crowded boards.

</details>

### 13. owned first, and the commander is not owned

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 225 names.

Cards: 99. Repair turn: true. Block findings: 0.

**Summary:** Karlov of the Ghost Council leads a resilient Orzhov lifegain strategy that builds a protected creature board, converts steady life gain into overwhelming combat pressure, and uses efficient answers to keep opposing engines contained. Angels, equipped attackers, and scalable legendary creatures provide multiple paths to a decisive finish, while broad sweepers offer a strong recovery tool when the battlefield gets out of hand.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [WARN] `not_owned`: Karlov of the Ghost Council: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 2.84 over 63 nonland cards

<details><summary>The deck list</summary>

- 18 Plains | land | Provides dependable white mana for the commander, lifegain pieces, Angels, and white interaction.
- 18 Swamp | land | Provides dependable black mana for draw, removal, and black lifegain support.
- 1 Sol Ring | ramp | Efficient early acceleration.
- 1 Arcane Signet | ramp | Reliable color fixing and acceleration.
- 1 Fellwar Stone | ramp | Low-cost mana acceleration.
- 1 Commander's Sphere | ramp | Fixes mana early and can become a card later.
- 1 Wayfarer's Bauble | ramp | Finds a basic land and improves permanent mana.
- 1 Thought Vessel | ramp | Produces mana while supporting a full hand.
- 1 Astral Cornucopia | ramp | Flexible mana production that scales into later turns.
- 1 Lotho, Corrupt Shirriff | ramp | Converts opposing development into Treasures.
- 1 Giada, Font of Hope | ramp | Accelerates the deck's Angel threats.
- 1 Sword of the Animist | ramp | Turns attacks into ongoing basic-land ramp.
- 1 Skullclamp | draw | Efficient card flow alongside small creatures and expendable bodies.
- 1 Mask of Memory | draw | Rewards creature attacks with card selection.
- 1 Tome of Legends | draw | Keeps cards flowing around the commander.
- 1 Wall of Omens | draw | Early defense that replaces itself.
- 1 Inspiring Overseer | draw | Aerial body that provides immediate card advantage.
- 1 Idol of Oblivion | draw | Repeatable card advantage for a token-capable creature deck.
- 1 Lembas | draw | Smooths draws and offers later recursion.
- 1 Exemplar of Light | draw | Lifegain-oriented creature that sustains the hand.
- 1 Night's Whisper | draw | Straightforward, inexpensive card draw.
- 1 Painful Truths | draw | Efficient refill in a two-color deck.
- 1 Swords to Plowshares | removal | Premium answer to a troublesome creature.
- 1 Generous Gift | removal | Answers any problematic permanent.
- 1 Stroke of Midnight | removal | Flexible instant-speed permanent removal.
- 1 Get Lost | removal | Efficient answer to creatures, enchantments, and planeswalkers.
- 1 Infernal Grasp | removal | Clean, unconditional creature removal.
- 1 Bitter Triumph | removal | Flexible removal that can answer creatures or planeswalkers.
- 1 Dispatch | removal | Low-cost creature answer with artifact support.
- 1 Crib Swap | removal | Exiles creatures while bypassing death triggers.
- 1 Lightning Greaves | interaction | Protects Karlov and key creatures while enabling immediate attacks.
- 1 Swiftfoot Boots | interaction | Protects vital creatures from targeted disruption.
- 1 Take Up the Shield | interaction | Protects a creature while adding a useful lifegain moment.
- 1 Clever Concealment | interaction | Shields the developed board from opposing answers.
- 1 Unbreakable Formation | interaction | Preserves the creature board through a sweep and supports aggression.
- 1 Bastion Protector | interaction | Keeps the commander safer while adding a body to the board.
- 1 Angel of Vitality | synergy | Amplifies life gains and becomes a substantial lifelink attacker.
- 1 Aerith Gainsborough | synergy | Supports the deck's life-focused creature plan.
- 1 Compassionate Healer | synergy | Supplies repeatable lifegain support for Karlov.
- 1 Kor Firewalker | synergy | A resilient lifegain source with relevant protection.
- 1 Light of Promise | synergy | Turns recurring lifegain into a rapidly growing threat.
- 1 Rosie Cotton of South Lane | synergy | Converts lifegain into permanent growth across the board.
- 1 Second Breakfast | synergy | Provides a compact lifegain-focused value effect.
- 1 Prideful Feastling | synergy | Contributes to the deck's food and lifegain theme.
- 1 Adventurous Eater // Have a Bite | synergy | Offers a themed body and a useful lifegain-adjacent spell option.
- 1 City Pigeon | synergy | A low-cost creature that supports the deck's lifegain plan.
- 1 Crowd of True Believers | synergy | Builds the board while reinforcing the deck's core theme.
- 1 Dancer's Chakrams | synergy | Equipment support for a creature-centric lifegain strategy.
- 1 Elixir | synergy | Provides a compact source of themed value.
- 1 Well-Worn Spatula | synergy | Supports equipped attackers and the deck's incremental-value plan.
- 1 Angel of Invention | threat | Produces a board presence and closes games in the air.
- 1 Lyra Dawnbringer | threat | Powerful lifelink threat that dominates combat.
- 1 Shattered Angel | threat | A flying threat that rewards opposing land development with life.
- 1 Victory's Herald | threat | Turns an attacking creature force into a powerful aerial assault.
- 1 Frodo, Sauron's Bane | threat | A scalable legendary threat with a high-impact finishing ceiling.
- 1 Minwu, White Mage | threat | A durable, life-focused legendary threat.
- 1 Gollum, Silent Slinker // Meager Meal | threat | Provides an evasive black threat with additional utility.
- 1 Sneering Shadewriter | threat | A black creature that pressures opponents while advancing the board.
- 1 Dawnhand Eulogist | threat | A substantial creature threat for closing contested games.
- 1 Lo and Li, Twin Tutors | threat | A legendary threat that provides meaningful board presence.
- 1 Invisible Woman, Sue Storm | threat | A resilient legendary threat for sustained pressure.
- 1 Reaping Willow | threat | A large creature that helps finish games after the board is stabilized.
- 1 Fumigate | wipe | Resets creature-heavy boards while restoring life.
- 1 Austere Command | wipe | Flexible sweeper that can spare the most important parts of the board.
- 1 Vanquish the Horde | wipe | Efficiently clears crowded creature boards.

</details>

### 14. the user delegates the commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 244 names.

Cards: 99. Repair turn: false. Block findings: 0.

**Summary:** Atarka, World Render leads a green-red Dragon deck that builds mana, establishes a board of Dragons, and turns that board into forceful attacks. It wins by keeping combat pressure high with large threats while its Dragon-focused support, card flow, and removal keep the plan moving. The deck gives up broad alternate plans for a focused creature strategy, so it is strongest when it can maintain its battlefield presence.

- JUDGE [true]: "Atarka, World Render leads a green-red Dragon deck". Atarka, World Render is a legendary creature and is Gruul (red-green) colored, making it a valid commander choice for a green-red Dragon deck.
- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 4.48 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Ancient Tomb | land | A land slot in the mana base.
- 1 Arid Mesa | land | A land slot in the mana base.
- 1 Bloodstained Mire | land | A land slot in the mana base.
- 1 Boseiju, Who Endures | land | A land slot in the mana base.
- 1 Cavern of Souls | land | A land slot in the mana base.
- 1 Cinder Glade | land | A land slot in the mana base.
- 1 City of Brass | land | A land slot in the mana base.
- 1 Command Tower | land | A land slot in the mana base.
- 1 Crucible of the Spirit Dragon | land | A land slot in the mana base.
- 1 Evolving Wilds | land | A land slot in the mana base.
- 1 Exotic Orchard | land | A land slot in the mana base.
- 1 Fabled Passage | land | A land slot in the mana base.
- 1 Forest | land | A basic land slot in the mana base.
- 1 Gemstone Caverns | land | A land slot in the mana base.
- 1 Haven of the Spirit Dragon | land | A land slot in the mana base.
- 1 Maelstrom of the Spirit Dragon | land | A land slot in the mana base.
- 1 Mana Confluence | land | A land slot in the mana base.
- 1 Misty Rainforest | land | A land slot in the mana base.
- 1 Mountain | land | A basic land slot in the mana base.
- 1 Myriad Landscape | land | A land slot in the mana base.
- 1 Nykthos, Shrine to Nyx | land | A land slot in the mana base.
- 1 Path of Ancestry | land | A land slot in the mana base.
- 1 Reflecting Pool | land | A land slot in the mana base.
- 1 Rogue's Passage | land | A land slot in the mana base.
- 1 Rootbound Crag | land | A land slot in the mana base.
- 1 Scalding Tarn | land | A land slot in the mana base.
- 1 Stomping Ground | land | A land slot in the mana base.
- 1 Temple of the Dragon Queen | land | A land slot in the mana base.
- 1 Temple of the False God | land | A land slot in the mana base.
- 1 Terramorphic Expanse | land | A land slot in the mana base.
- 1 Three Tree City | land | A land slot in the mana base.
- 1 Verdant Catacombs | land | A land slot in the mana base.
- 1 War Room | land | A land slot in the mana base.
- 1 Windswept Heath | land | A land slot in the mana base.
- 1 Wooded Foothills | land | A land slot in the mana base.
- 1 Yavimaya, Cradle of Growth | land | A land slot in the mana base.
- 1 Ancient Copper Dragon | ramp | A Dragon ramp card that advances the deck's mana plan.
- 1 Atsushi, the Blazing Sky | ramp | A Dragon ramp card for reaching the deck's larger plays.
- 1 Carnelian Orb of Dragonkind | ramp | A ramp card supporting the Dragon plan.
- 1 Dragon's Hoard | ramp | A ramp card supporting the Dragon plan.
- 1 Dragonstorm Globe | ramp | A ramp card supporting the Dragon plan.
- 1 Ganax, Astral Hunter | ramp | A Dragon ramp card that supports the creature-heavy plan.
- 1 Goldspan Dragon | ramp | A Dragon ramp card for accelerating the deck.
- 1 Jade Orb of Dragonkind | ramp | A ramp card supporting the Dragon plan.
- 1 Klauth, Unrivaled Ancient | ramp | A Dragon ramp card for powering the deck's top end.
- 1 Old Gnawbone | ramp | A Dragon ramp card for sustaining large turns.
- 1 Avaricious Dragon | draw | A Dragon draw card that keeps the deck supplied.
- 1 Beast Whisperer | draw | A draw card for the creature-focused plan.
- 1 Dragon Mage | draw | A Dragon draw card that supports continued pressure.
- 1 Dragonborn Champion | draw | A draw card that rewards the deck's combat focus.
- 1 Elemental Bond | draw | A draw card for a deck built around large creatures.
- 1 Garruk's Uprising | draw | A draw card for the large-creature strategy.
- 1 Guardian Project | draw | A draw card that supports the creature-heavy deck.
- 1 Return of the Wildspeaker | draw | A draw card that fits the large-creature plan.
- 1 Rishkar's Expertise | draw | A draw card for refilling during the Dragon plan.
- 1 Sylvan Library | draw | A draw card that improves the deck's access to resources.
- 1 Heroic Intervention | interaction | An interaction card for protecting the deck's board presence.
- 1 Lightning Greaves | interaction | An interaction card supporting key creatures.
- 1 Swiftfoot Boots | interaction | An interaction card supporting key creatures.
- 1 Tamiyo's Safekeeping | interaction | An interaction card for defending an important permanent.
- 1 Veil of Summer | interaction | An interaction card for protecting the deck's plan.
- 1 Whispersilk Cloak | interaction | An interaction card supporting an important threat.
- 1 Dragon Tempest | removal | A Dragon-focused removal card.
- 1 Dragonlord Atarka | removal | A Dragon removal card that also adds to the creature plan.
- 1 Drakuseth, Maw of Flames | removal | A Dragon removal card for the deck's top end.
- 1 Glorybringer | removal | A Dragon removal card that maintains pressure.
- 1 Scourge of Valkas | removal | A Dragon removal card that fits the tribal plan.
- 1 Terror of the Peaks | removal | A Dragon removal card that supports creature deployment.
- 1 Wrathful Red Dragon | removal | A Dragon removal card for the combat-focused deck.
- 1 Foe-Razer Regent | removal | A Dragon removal card that contributes to the board.
- 1 Balefire Dragon | wipe | A Dragon board-wipe card for resetting crowded boards.
- 1 Draconic Intervention | wipe | A board-wipe card for handling opposing boards.
- 1 Incinerator of the Guilty | wipe | A Dragon board-wipe card that fits the main plan.
- 1 Breaching Dragonstorm | synergy | A Dragon synergy card for the tribal strategy.
- 1 Crucible of Fire | synergy | A Dragon synergy card that reinforces the creature plan.
- 1 Dracogenesis | synergy | A Dragon synergy card for the deck's central theme.
- 1 Dragon Egg | synergy | A Dragon synergy card for the tribal shell.
- 1 Dragon Hatchling | synergy | A Dragon synergy card for the tribal shell.
- 1 Dragonkin Berserker | synergy | A Dragon synergy card for the combat plan.
- 1 Dragonlord's Servant | synergy | A Dragon synergy card supporting the deck's creature plan.
- 1 Dragonspeaker Shaman | synergy | A Dragon synergy card supporting the deck's creature plan.
- 1 Dragonstorm | synergy | A Dragon synergy card for the tribal strategy.
- 1 Firespitter Whelp | synergy | A Dragon synergy card for the creature-focused plan.
- 1 Kargan Dragonrider | synergy | A Dragon synergy card for the tribal strategy.
- 1 Last Light of Durin's Day | synergy | A synergy card supporting the Dragon plan.
- 1 Minion of the Mighty | synergy | A synergy card that supports the deck's Dragon focus.
- 1 Sarkhan's Triumph | synergy | A Dragon synergy card that supports the focused creature plan.
- 1 Ancient Bronze Dragon | threat | A major Dragon threat for closing games through combat.
- 1 Backdraft Hellkite | threat | A Dragon threat that adds to the attacking force.
- 1 Blast-Furnace Hellkite | threat | A Dragon threat for applying combat pressure.
- 1 Hellkite Charger | threat | A Dragon threat for the combat-focused strategy.
- 1 Hellkite Courser | threat | A Dragon threat that strengthens the creature plan.
- 1 Lathliss, Dragon Queen | threat | A Dragon threat for the deck's tribal board presence.
- 1 Scourge of the Throne | threat | A Dragon threat for pressing opponents in combat.
- 1 Terror of Mount Velus | threat | A Dragon threat for the deck's finishing turns.
- 1 Thrakkus the Butcher | threat | A Dragon threat for the combat-oriented plan.
- 1 Twinflame Tyrant | threat | A Dragon threat for delivering decisive pressure.
- 1 Utvara Hellkite | threat | A Dragon threat that supports the tribal attack plan.
- 1 Thunderbreak Regent | threat | A Dragon threat that reinforces the deck's battlefield presence.

</details>

### 15. delegated commander, owned first

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 225 names.

Cards: 99. Repair turn: false. Block findings: 0.

**Summary:** Cecil leads a white-black lifegain deck that develops its mana, builds a creature board, and layers lifegain and equipment synergy around its threats. It aims to win by sustaining pressure with Angels, legends, and other creatures while protecting its key pieces and clearing away opposing problems when necessary. The deck trades explosive speed and a dedicated alternate finish for a steadier, board-focused game that can rebuild with card advantage and selective protection.

- [INFO] `bracket_prose_rules`: bracket 2 (Core, 8+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 3.05 over 63 nonland cards
- [INFO] `basics_added`: the list was 1 card short, so the builder added 1 basic land

<details><summary>The deck list</summary>

- 14 Plains | land | Basic white mana for the deck’s lifegain-focused spells.
- 12 Swamp | land | Basic black mana for the deck’s support and threat suite.
- 1 Command Tower | land | Flexible mana fixing for Cecil’s colors.
- 1 Path of Ancestry | land | Color fixing within the creature-focused mana base.
- 1 Evolving Wilds | land | Flexible land fixing.
- 1 Fabled Passage | land | Flexible land fixing.
- 1 Marsh Flats | land | Fetch-land fixing for the two-color mana base.
- 1 Exotic Orchard | land | Additional flexible color fixing.
- 1 Secluded Courtyard | land | Creature-oriented color fixing.
- 1 City of Brass | land | Provides broad color fixing.
- 1 Vibrant Cityscape | land | Additional fixing land for the mana base.
- 1 Thriving Moor | land | Black source with flexible fixing support.
- 1 Sol Ring | ramp | Efficient early mana acceleration.
- 1 Arcane Signet | ramp | Reliable color-fixing mana acceleration.
- 1 Bender's Waterskin | ramp | Artifact-based mana acceleration.
- 1 Commander's Sphere | ramp | Color-fixing artifact acceleration.
- 1 Fellwar Stone | ramp | Low-cost artifact mana acceleration.
- 1 Wayfarer's Bauble | ramp | Early mana development.
- 1 Sword of the Animist | ramp | Equipment-based mana development that supports attacking creatures.
- 1 Relic of Legends | ramp | Artifact mana acceleration for the creature-heavy build.
- 1 Thought Vessel | ramp | Artifact mana acceleration.
- 1 Giada, Font of Hope | ramp | Creature-based mana acceleration that fits the white creature core.
- 1 Buster Sword | draw | Equipment-based card advantage for the creature plan.
- 1 Call of the Ring | draw | Ongoing card advantage support.
- 1 Exemplar of Light | draw | Creature-based card advantage in the white core.
- 1 Idol of Oblivion | draw | Artifact card advantage for the deck’s board plan.
- 1 Inspiring Overseer | draw | Creature-based card advantage.
- 1 Lembas | draw | Low-cost artifact card advantage.
- 1 Mask of Memory | draw | Equipment-based card advantage for attacking creatures.
- 1 Night's Whisper | draw | Direct black card advantage.
- 1 Puresteel Paladin | draw | Card advantage that complements the equipment package.
- 1 Skullclamp | draw | Equipment-based card advantage.
- 1 Bastion Protector | interaction | Creature-based protection for the deck’s key legend.
- 1 Champion's Helm | interaction | Equipment protection for Cecil or another important creature.
- 1 Lightning Greaves | interaction | Efficient equipment protection.
- 1 Swiftfoot Boots | interaction | Additional equipment protection.
- 1 Unbreakable Formation | interaction | Board-protection interaction.
- 1 Take Up the Shield | interaction | Compact creature-protection interaction.
- 1 Banishing Light | removal | Versatile white answer to problematic permanents.
- 1 Bitter Triumph | removal | Flexible black removal.
- 1 Crib Swap | removal | Creature-focused removal.
- 1 Generous Gift | removal | Broad permanent removal.
- 1 Get Lost | removal | Efficient white removal.
- 1 Infernal Grasp | removal | Direct black creature removal.
- 1 Swords to Plowshares | removal | Efficient white creature removal.
- 1 Stroke of Midnight | removal | Flexible white removal.
- 1 Austere Command | wipe | Flexible board-reset option.
- 1 Fumigate | wipe | Board-reset option for stabilizing a crowded table.
- 1 Vanquish the Horde | wipe | Efficient creature-board reset.
- 1 Aerith Gainsborough | synergy | A lifegain synergy creature for Cecil’s plan.
- 1 Aettir and Priwen | synergy | Equipment synergy for the creature-centered build.
- 1 Angel of Vitality | synergy | Lifegain synergy within the white creature suite.
- 1 Compassionate Healer | synergy | Lifegain-focused creature synergy.
- 1 Dancer's Chakrams | synergy | Equipment synergy for the deck’s creature plan.
- 1 Elixir | synergy | Artifact synergy supporting the lifegain theme.
- 1 Excalibur II | synergy | Equipment synergy for the deck’s threats.
- 1 Light of Promise | synergy | Lifegain-focused enchantment synergy.
- 1 Night Nurse, Healer of Heroes | synergy | Lifegain synergy in a legendary creature slot.
- 1 Prideful Feastling | synergy | Creature synergy for the lifegain plan.
- 1 Rosie Cotton of South Lane | synergy | Lifegain synergy in the creature core.
- 1 Second Breakfast | synergy | Lifegain-themed instant synergy.
- 1 Well-Worn Spatula | synergy | Equipment synergy for the board-focused plan.
- 1 White Mage's Staff | synergy | Equipment synergy that fits the white lifegain shell.
- 1 Angel of Invention | threat | White creature threat for closing games on the board.
- 1 Bill the Pony | threat | Legendary creature threat that adds board presence.
- 1 Canyon Crawler | threat | Creature threat for applying pressure.
- 1 Dawnhand Eulogist | threat | Creature threat for the deck’s board plan.
- 1 Foggy Swamp Hunters | threat | Creature threat that expands the black creature suite.
- 1 Gollum, Silent Slinker // Meager Meal | threat | Black legendary threat with an attached adventure.
- 1 Invisible Woman, Sue Storm | threat | Legendary creature threat for board presence.
- 1 Lyra Dawnbringer | threat | Powerful Angel threat for the white creature core.
- 1 Minwu, White Mage | threat | Legendary white creature threat.
- 1 Shattered Angel | threat | Angel threat supporting the white creature plan.
- 1 Sneering Shadewriter | threat | Black creature threat for applying pressure.
- 1 Victory's Herald | threat | Angel threat for finishing through combat.

</details>

