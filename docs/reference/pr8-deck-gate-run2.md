# PR-8 deck gate

Run date: 2026-08-27. Card snapshot: 2026-08-24.

Verdict: PASS. 12 of 12 decks passed every block check, and 0 invented names reached the user. Both bars are zero tolerance.

## Summary

| Measure | Value |
|---|---|
| Prompts | 12 |
| Decks returned | 12 |
| Decks with no block finding | 12 |
| Invented names that reached the user | 0 |
| Decks that needed the repair turn | 0 |
| Errors | 0 |
| Prompt version | 2 |
| Calls | 12 |
| Cost | $0.6166 |
| Time | 489 seconds |

## Findings by code

A block stops the deck. A warning and a note are reports.

| Code | Count |
|---|---|
| `curve_summary` | 12 |
| `bracket_prose_rules` | 6 |
| `not_owned` | 2 |
| `basics_added` | 1 |

By severity: BLOCK 0. WARN 2. INFO 19. 

## Decks

### 1. lifegain Commander, any card

Format: Commander. Theme: lifegain. Pool: any_card. Shortlist: 303 names.

Cards: 99. Repair turn: false. Block findings: 0.

**Summary:** This is a lifegain-focused Karlov deck that develops its board, keeps cards flowing, and uses its synergy pieces to turn that plan into sustained pressure. It wins by establishing durable threats and clearing away resistance with focused answers or a reset when the table gets ahead. The tradeoff is a board-centered strategy that needs time to assemble its pieces and can be slowed when that board is repeatedly disrupted.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 3.65 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Aclazotz, Deepest Betrayal // Temple of the Dead | land | Provides a land slot for the deck.
- 1 Adventurer's Inn | land | Provides a land slot for the deck.
- 1 Arguel's Blood Fast // Temple of Aclazotz | land | Provides a land slot for the deck.
- 1 Azor's Gateway // Sanctum of the Sun | land | Provides a land slot for the deck.
- 1 Blighted Steppe | land | Provides a land slot for the deck.
- 1 Diamond Valley | land | Provides a land slot for the deck.
- 1 Glasswing Grace // Age-Graced Chapel | land | Provides a land slot for the deck.
- 1 Glimmerpost | land | Provides a land slot for the deck.
- 1 Hidden Hideout | land | Provides a land slot for the deck.
- 1 High Market | land | Provides a land slot for the deck.
- 1 Hostile Hostel // Creeping Inn | land | Provides a land slot for the deck.
- 1 Inventors' Fair | land | Provides a land slot for the deck.
- 1 Kabira Crossroads | land | Provides a land slot for the deck.
- 1 Legion's Landing // Adanto, the First Fort | land | Provides a land slot for the deck.
- 1 Miren, the Moaning Well | land | Provides a land slot for the deck.
- 1 Nomad Stadium | land | Provides a land slot for the deck.
- 1 Phyrexia's Core | land | Provides a land slot for the deck.
- 4 Plains | land | Provides basic land slots for the deck.
- 1 Radiant Fountain | land | Provides a land slot for the deck.
- 1 Restless Fortress | land | Provides a land slot for the deck.
- 1 Scoured Barrens | land | Provides a land slot for the deck.
- 1 Seraph Sanctuary | land | Provides a land slot for the deck.
- 1 Shambling Vent | land | Provides a land slot for the deck.
- 1 Skyclave Cleric // Skyclave Basilica | land | Provides a land slot for the deck.
- 1 Springjack Pasture | land | Provides a land slot for the deck.
- 1 Starlit Sanctum | land | Provides a land slot for the deck.
- 2 Swamp | land | Provides basic land slots for the deck.
- 1 Tomb of the Spirit Dragon | land | Provides a land slot for the deck.
- 1 Vault of the Archangel | land | Provides a land slot for the deck.
- 1 Westvale Abbey // Ormendahl, Profane Prince | land | Provides a land slot for the deck.
- 1 Witch's Clinic | land | Provides a land slot for the deck.
- 1 Zof Consumption // Zof Bloodbog | land | Provides a land slot for the deck.
- 1 Altar of the Pantheon | ramp | Provides ramp support for the deck's setup.
- 1 Angel of Indemnity | ramp | Provides ramp support for the deck's setup.
- 1 Battle Angels of Tyr | ramp | Provides ramp support for the deck's setup.
- 1 Beza, the Bounding Spring | ramp | Provides ramp support for the deck's setup.
- 1 Colossal Plow | ramp | Provides ramp support for the deck's setup.
- 1 Crypt Ghast | ramp | Provides ramp support for the deck's setup.
- 1 Hierophant's Chalice | ramp | Provides ramp support for the deck's setup.
- 1 Nuka-Cola Vending Machine | ramp | Provides ramp support for the deck's setup.
- 1 Orazca Relic | ramp | Provides ramp support for the deck's setup.
- 1 Pristine Talisman | ramp | Provides ramp support for the deck's setup.
- 1 Alhammarret's Archive | draw | Provides card-draw support for the lifegain plan.
- 1 Archivist of Oghma | draw | Provides card-draw support for the lifegain plan.
- 1 Convalescent Care | draw | Provides card-draw support for the lifegain plan.
- 1 Cosmos Elixir | draw | Provides card-draw support for the lifegain plan.
- 1 Dawn of Hope | draw | Provides card-draw support for the lifegain plan.
- 1 Enduring Innocence | draw | Provides card-draw support for the lifegain plan.
- 1 Mangara, the Diplomat | draw | Provides card-draw support for the lifegain plan.
- 1 Sigarda's Splendor | draw | Provides card-draw support for the lifegain plan.
- 1 The Gaffer | draw | Provides card-draw support for the lifegain plan.
- 1 Well of Lost Dreams | draw | Provides card-draw support for the lifegain plan.
- 1 Aetherflux Reservoir | removal | Provides removal for opposing problems.
- 1 Ayli, Eternal Pilgrim | removal | Provides removal for opposing problems.
- 1 Cavalier of Night | removal | Provides removal for opposing problems.
- 1 Murderous Rider // Swift End | removal | Provides removal for opposing problems.
- 1 Solitude | removal | Provides removal for opposing problems.
- 1 Umezawa's Jitte | removal | Provides removal for opposing problems.
- 1 Vein Ripper | removal | Provides removal for opposing problems.
- 1 Vona, Butcher of Magan | removal | Provides removal for opposing problems.
- 1 Alseid of Life's Bounty | interaction | Provides interaction to protect the deck's plan.
- 1 Faith's Shield | interaction | Provides interaction to protect the deck's plan.
- 1 Metropolis Reformer | interaction | Provides interaction to protect the deck's plan.
- 1 Rune-Tail, Kitsune Ascendant // Rune-Tail's Essence | interaction | Provides interaction to protect the deck's plan.
- 1 Sephara, Sky's Blade | interaction | Provides interaction to protect the deck's plan.
- 1 Werefox Bodyguard | interaction | Provides interaction to protect the deck's plan.
- 1 Ajani, Strength of the Pride | wipe | Provides a board-reset option.
- 1 Fumigate | wipe | Provides a board-reset option.
- 1 Kaya's Wrath | wipe | Provides a board-reset option.
- 1 Ajani's Pridemate | synergy | Supports the lifegain synergy plan.
- 1 Angel of Vitality | synergy | Supports the lifegain synergy plan.
- 1 Angelic Accord | synergy | Supports the lifegain synergy plan.
- 1 Blood Artist | synergy | Supports the lifegain synergy plan.
- 1 Bloodthirsty Aerialist | synergy | Supports the lifegain synergy plan.
- 1 Cleric Class | synergy | Supports the lifegain synergy plan.
- 1 Cleric of Life's Bond | synergy | Supports the lifegain synergy plan.
- 1 Heliod, Sun-Crowned | synergy | Supports the lifegain synergy plan.
- 1 Indulgent Aristocrat | synergy | Supports the lifegain synergy plan.
- 1 Lurrus of the Dream-Den | synergy | Supports the lifegain synergy plan.
- 1 Resplendent Angel | synergy | Supports the lifegain synergy plan.
- 1 Righteous Valkyrie | synergy | Supports the lifegain synergy plan.
- 1 Sanguine Bond | synergy | Supports the lifegain synergy plan.
- 1 Serra Ascendant | synergy | Supports the lifegain synergy plan.
- 1 Archangel of Thune | threat | Provides a major threat for the lifegain plan.
- 1 Astarion, the Decadent | threat | Provides a major threat for the lifegain plan.
- 1 Attended Healer | threat | Provides a major threat for the lifegain plan.
- 1 Celestine, the Living Saint | threat | Provides a major threat for the lifegain plan.
- 1 Defiant Bloodlord | threat | Provides a major threat for the lifegain plan.
- 1 Divinity of Pride | threat | Provides a major threat for the lifegain plan.
- 1 Elenda, Saint of Dusk | threat | Provides a major threat for the lifegain plan.
- 1 Liesa, Forgotten Archangel | threat | Provides a major threat for the lifegain plan.
- 1 Lyra Dawnbringer | threat | Provides a major threat for the lifegain plan.
- 1 Nykthos Paragon | threat | Provides a major threat for the lifegain plan.
- 1 Rhox Faithmender | threat | Provides a major threat for the lifegain plan.
- 1 Valkyrie Harbinger | threat | Provides a major threat for the lifegain plan.

</details>

### 2. aristocrats Commander, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 179 names.

Cards: 99. Repair turn: false. Block findings: 0.

**Summary:** This Karlov-led aristocrats deck establishes its mana, develops a creature-heavy board, and leans on its synergy cards to sustain pressure through the middle turns. It closes by turning that developed board and its selected threats into a decisive advantage while using removal, protective interaction, and wipes to prevent opposing boards from taking over. It gives up some dedicated finishing pressure and specialized synergy depth in exchange for a broad, library-first collection of lands, creatures, and supporting artifacts.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [WARN] `not_owned`: Karlov of the Ghost Council: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 2.94 over 63 nonland cards

<details><summary>The deck list</summary>

- 7 Plains | land | Basic white land for the mana base.
- 7 Swamp | land | Basic black land for the mana base.
- 1 Ash Barrens | land | Land slot from the shortlist.
- 1 Bonders' Enclave | land | Land slot from the shortlist.
- 1 Castle Doom | land | Black land slot for the mana base.
- 1 Castle Locthwain | land | Black land slot for the mana base.
- 1 Command Tower | land | Multicolor commander land slot.
- 1 Evolving Wilds | land | Land slot from the shortlist.
- 1 Exotic Orchard | land | Land slot from the shortlist.
- 1 Fabled Passage | land | Land slot from the shortlist.
- 1 Field of Ruin | land | Land slot from the shortlist.
- 1 Ghost Quarter | land | Land slot from the shortlist.
- 1 Ifnir Deadlands | land | Black land slot for the mana base.
- 1 Marsh Flats | land | Land slot from the shortlist.
- 1 Minas Tirith | land | White land slot for the mana base.
- 1 Opal Palace | land | Land slot from the shortlist.
- 1 Path of Ancestry | land | Land slot from the shortlist.
- 1 Plaza of Heroes | land | Land slot from the shortlist.
- 1 Scavenger Grounds | land | Land slot from the shortlist.
- 1 Takenuma, Abandoned Mire | land | Black land slot for the mana base.
- 1 Terramorphic Expanse | land | Land slot from the shortlist.
- 1 Thriving Moor | land | Black land slot for the mana base.
- 1 War Room | land | Land slot from the shortlist.
- 1 Windbrisk Heights | land | White land slot for the mana base.
- 1 Arcane Signet | ramp | Listed ramp piece for early development.
- 1 Commander's Sphere | ramp | Listed ramp piece for the mana package.
- 1 Fellwar Stone | ramp | Listed ramp piece for early development.
- 1 Lotho, Corrupt Shirriff | ramp | Listed ramp creature for the mana package.
- 1 Relic of Legends | ramp | Listed ramp piece for the mana package.
- 1 Sol Ring | ramp | Listed ramp piece for early development.
- 1 Springleaf Drum | ramp | Listed ramp piece for the mana package.
- 1 Sword of the Animist | ramp | Listed ramp equipment for the mana package.
- 1 Thought Vessel | ramp | Listed ramp piece for the mana package.
- 1 Wayfarer's Bauble | ramp | Listed ramp piece for early development.
- 1 Call of the Ring | draw | Listed draw card for ongoing resources.
- 1 Exemplar of Light | draw | Listed draw creature for the resource suite.
- 1 Idol of Oblivion | draw | Listed draw artifact for the resource suite.
- 1 Inspiring Overseer | draw | Listed draw creature for the resource suite.
- 1 Lembas | draw | Listed draw artifact for the resource suite.
- 1 Mask of Memory | draw | Listed draw equipment for the resource suite.
- 1 Night's Whisper | draw | Listed draw spell for the resource suite.
- 1 Puresteel Paladin | draw | Listed draw creature for the resource suite.
- 1 Skullclamp | draw | Listed draw equipment for the resource suite.
- 1 Tome of Legends | draw | Listed draw artifact for the resource suite.
- 1 Bitter Triumph | removal | Listed removal spell for opposing permanents.
- 1 Claim the Precious | removal | Listed removal spell for opposing permanents.
- 1 Crib Swap | removal | Listed removal spell for opposing creatures.
- 1 Fatal Push | removal | Listed removal spell for opposing creatures.
- 1 Generous Gift | removal | Listed removal spell for opposing permanents.
- 1 Get Lost | removal | Listed removal spell for opposing permanents.
- 1 Infernal Grasp | removal | Listed removal spell for opposing creatures.
- 1 Swords to Plowshares | removal | Listed removal spell for opposing creatures.
- 1 Bastion Protector | interaction | Listed interaction creature for the protection package.
- 1 Boromir, Warden of the Tower | interaction | Listed interaction creature for the protection package.
- 1 Clever Concealment | interaction | Listed interaction spell for the protection package.
- 1 Gift of Immortality | interaction | Listed interaction aura for the protection package.
- 1 Lightning Greaves | interaction | Listed interaction equipment for the protection package.
- 1 Swiftfoot Boots | interaction | Listed interaction equipment for the protection package.
- 1 Austere Command | wipe | Listed wipe for resetting the board.
- 1 Dusk // Dawn | wipe | Listed wipe for resetting the board.
- 1 Fumigate | wipe | Listed wipe for resetting the board.
- 1 Arcade Cabinet | synergy | Listed synergy artifact for the aristocrats shell.
- 1 Denethor, Ruling Steward | synergy | Listed synergy creature for the aristocrats shell.
- 1 Gollum, Patient Plotter | synergy | Listed synergy creature for the aristocrats shell.
- 1 Gríma Wormtongue | synergy | Listed synergy creature for the aristocrats shell.
- 1 Joo Dee, One of Many | synergy | Listed synergy creature for the aristocrats shell.
- 1 Nimble Hobbit | synergy | Listed synergy creature for the aristocrats shell.
- 1 Phantom Train | synergy | Listed synergy artifact for the aristocrats shell.
- 1 Bill the Pony | threat | Listed threat creature for applying pressure.
- 1 Ahriman | other | Additional creature from the shortlist for the board plan.
- 1 Angel of Serenity | other | Additional creature from the shortlist for the board plan.
- 1 Archfiend of Ifnir | other | Additional creature from the shortlist for the board plan.
- 1 Blowfly Infestation | other | Additional black enchantment from the shortlist.
- 1 Cirith Ungol Patrol | other | Additional black creature from the shortlist.
- 1 Fiend Hunter | other | Additional white creature from the shortlist.
- 1 Gollum the Abandoned | other | Additional black creature from the shortlist.
- 1 Gorbag of Minas Morgul | other | Additional black creature from the shortlist.
- 1 Grave Venerations | other | Additional black enchantment from the shortlist.
- 1 Kingpin's Enforcers | other | Additional black creature from the shortlist.
- 1 Massacre Girl, Known Killer | other | Additional black creature from the shortlist.
- 1 Orcish Bowmasters | other | Additional black creature from the shortlist.
- 1 Palace Jailer | other | Additional white creature from the shortlist.
- 1 Rat King, Pale Piper | other | Additional black creature from the shortlist.
- 1 Stone of Erech | other | Additional artifact from the shortlist.
- 1 The Sackville-Bagginses | other | Additional white creature from the shortlist.
- 1 The Walls of Ba Sing Se | other | Additional artifact creature from the shortlist.
- 1 Witch-king of Angmar | other | Additional black creature from the shortlist.

</details>

### 3. artifacts Commander, bracket 4

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 299 names.

Cards: 99. Repair turn: false. Block findings: 0.

**Summary:** This Urza deck builds a dense artifact board, accelerates into major artifact turns, and keeps resources flowing while it applies pressure with artifact threats. It wins by turning that established board into sustained attacks backed by removal, interaction, and reset buttons. It gives up some flexibility to stay tightly focused on a blue artifact shell and its permanent-heavy game plan.

- [INFO] `curve_summary`: average mana value 4.05 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Academy Ruins | land | A land slot for the artifact-focused mana base.
- 1 Archway of Innovation | land | A land slot for the artifact-focused mana base.
- 1 Blinkmoth Nexus | land | A land slot for the artifact-focused mana base.
- 1 Buried Ruin | land | A land slot for the artifact-focused mana base.
- 1 Fomori Vault | land | A land slot for the artifact-focused mana base.
- 1 Glimmervoid | land | A land slot for the artifact-focused mana base.
- 1 Hall of Tagsin | land | A land slot for the artifact-focused mana base.
- 1 Heap Gate | land | A land slot for the artifact-focused mana base.
- 1 Inkmoth Nexus | land | A land slot for the artifact-focused mana base.
- 1 Inventors' Fair | land | A land slot for the artifact-focused mana base.
- 1 Mech Hangar | land | A land slot for the artifact-focused mana base.
- 1 Mirrex | land | A land slot for the artifact-focused mana base.
- 1 Mishra's Factory | land | A land slot for the artifact-focused mana base.
- 1 Mishra's Foundry | land | A land slot for the artifact-focused mana base.
- 1 Mishra's Workshop | land | A land slot for the artifact-focused mana base.
- 1 Otawara, Soaring City | land | A land slot for the blue mana base.
- 1 Phyrexia's Core | land | A land slot for the artifact-focused mana base.
- 1 Power Depot | land | An artifact land slot in the mana base.
- 1 Roadside Reliquary | land | A land slot for the artifact-focused mana base.
- 1 Spire of Industry | land | A land slot for the artifact-focused mana base.
- 1 The Monumental Facade | land | A land slot for the artifact-focused mana base.
- 1 The Mycosynth Gardens | land | A land slot for the artifact-focused mana base.
- 1 Urza's Factory | land | A land slot for the artifact-focused mana base.
- 1 Urza's Saga | land | An enchantment land slot in the artifact-focused mana base.
- 1 Urza's Workshop | land | A land slot for the artifact-focused mana base.
- 11 Island | land | Basic blue lands provide a reliable foundation for the mana base.
- 1 Mox Opal | ramp | A compact artifact ramp piece for accelerating the deck's development.
- 1 Metalworker | ramp | An artifact creature ramp piece for accelerating the deck's development.
- 1 Krark-Clan Ironworks | ramp | An artifact ramp piece for powering larger turns.
- 1 Moonsnare Prototype | ramp | A low-cost artifact ramp piece for early development.
- 1 Chief Engineer | ramp | A ramp creature that supports the artifact-focused plan.
- 1 Grand Architect | ramp | A ramp creature that supports the artifact-focused plan.
- 1 Inspiring Statuary | ramp | An artifact ramp piece for powering larger turns.
- 1 Karn, Legacy Reforged | ramp | A ramp threat that supports ambitious artifact turns.
- 1 Tezzeret the Seeker | ramp | A ramp planeswalker that supports the artifact plan.
- 1 The Mightstone and Weakstone | ramp | A legendary artifact ramp piece for advancing the board.
- 1 Emry, Lurker of the Loch | synergy | A legendary artifact-focused synergy piece.
- 1 Etherium Sculptor | synergy | An artifact creature that reinforces the deck's central theme.
- 1 Foundry Inspector | synergy | An artifact creature that reinforces the deck's central theme.
- 1 Mystic Forge | synergy | An artifact synergy piece for sustaining the focused plan.
- 1 Unwinding Clock | synergy | An artifact synergy piece for supporting the board.
- 1 Clock of Omens | synergy | An artifact synergy piece for supporting the board.
- 1 Voltaic Key | synergy | A compact artifact synergy piece.
- 1 Manifold Key | synergy | A compact artifact synergy piece.
- 1 Scrap Trawler | synergy | An artifact creature that deepens the deck's synergy package.
- 1 Transmute Artifact | synergy | A synergy spell that supports the artifact-centric strategy.
- 1 Whir of Invention | synergy | A synergy instant that supports the artifact-centric strategy.
- 1 Reshape | synergy | A synergy sorcery that supports the artifact-centric strategy.
- 1 Power Artifact | synergy | An enchantment synergy piece for the artifact plan.
- 1 Panharmonicon | synergy | An artifact synergy piece that rewards the deck's permanent focus.
- 1 Thoughtcast | draw | An artifact-focused draw spell for maintaining resources.
- 1 Thought Monitor | draw | An artifact creature that contributes to the draw package.
- 1 Thirst for Knowledge | draw | An instant draw spell for maintaining resources.
- 1 Sai, Master Thopterist | draw | A legendary draw creature that supports the artifact strategy.
- 1 Vedalken Archmage | draw | A draw creature that fits the artifact-focused shell.
- 1 Reverse Engineer | draw | A draw sorcery for maintaining resources.
- 1 One with the Machine | draw | A draw sorcery for maintaining resources.
- 1 Tezzeret, Artifice Master | draw | A draw planeswalker that fits the artifact plan.
- 1 Tezzeret, Betrayer of Flesh | draw | A draw planeswalker that fits the artifact plan.
- 1 Forensic Gadgeteer | draw | A draw creature that supports the artifact-focused shell.
- 1 Metallic Rebuke | interaction | An artifact-focused interaction spell for protecting the plan.
- 1 Disruption Protocol | interaction | An interaction instant for protecting the plan.
- 1 Stoic Rebuttal | interaction | An interaction instant for protecting the plan.
- 1 Assert Authority | interaction | An interaction instant for protecting the plan.
- 1 Padeem, Consul of Innovation | interaction | A legendary interaction creature for supporting the artifact board.
- 1 Welding Jar | interaction | A compact artifact interaction piece.
- 1 Aether Spellbomb | removal | A compact artifact removal piece.
- 1 Arcum Dagsson | removal | A legendary removal creature that fits the artifact strategy.
- 1 Resculpt | removal | A flexible removal instant.
- 1 Ravenform | removal | A removal sorcery for answering opposing permanents.
- 1 Spine of Ish Sah | removal | An artifact removal piece for the permanent-focused plan.
- 1 Portal to Phyrexia | removal | A large artifact removal piece for stabilizing the board.
- 1 Skysovereign, Consul Flagship | removal | A legendary artifact removal piece that also fits the deck theme.
- 1 Lux Cannon | removal | An artifact removal piece for the permanent-focused plan.
- 1 Contagion Engine | wipe | An artifact board wipe for resetting difficult boards.
- 1 Nevinyrral's Disk | wipe | An artifact board wipe for resetting difficult boards.
- 1 Oblivion Stone | wipe | An artifact board wipe for resetting difficult boards.
- 1 Kappa Cannoneer | threat | An artifact creature threat that pressures opponents.
- 1 Kuldotha Forgemaster | threat | An artifact creature threat for the artifact-heavy plan.
- 1 Master Transmuter | threat | An artifact creature threat that fits the central strategy.
- 1 Metalwork Colossus | threat | A large artifact creature threat.
- 1 Karn, Scion of Urza | threat | A planeswalker threat that supports the artifact-focused shell.
- 1 Threefold Thunderhulk | threat | An artifact creature threat for closing games.
- 1 The Capitoline Triad | threat | A legendary creature threat for closing games.
- 1 Traxos, Scourge of Kroog | threat | A legendary artifact creature threat.
- 1 Phyrexian Metamorph | threat | An artifact creature threat that fits the permanent-focused deck.
- 1 Broodstar | threat | A creature threat for closing games.
- 1 Darksteel Juggernaut | threat | An artifact creature threat for closing games.
- 1 Cyberdrive Awakener | threat | An artifact creature threat for converting board presence into pressure.

</details>

### 4. dinosaur tribal, bracket 2

Format: Commander. Theme: dinosaurs. Pool: any_card. Shortlist: 267 names.

Cards: 99. Repair turn: false. Block findings: 0.

**Summary:** This deck builds mana early, develops a dinosaur board, and uses Gishath as the centerpiece for a straightforward combat plan. Its large dinosaur threats are the main way it wins, while themed ramp, draw, removal, and board resets help it keep moving through a longer game. It gives up some speed and finely tuned answers in exchange for a simple, creature-heavy plan that is easy to recognize and play.

- [INFO] `bracket_prose_rules`: bracket 2 (Core, 8+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 4.00 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Beast Whisperer | draw | Creature-based draw that supports the deck’s dinosaur-focused plan.
- 1 Garruk's Uprising | draw | A draw piece for keeping the deck supplied with cards.
- 1 Guardian Project | draw | A long-game draw source for the creature-heavy strategy.
- 1 Harmonize | draw | Straightforward card draw for rebuilding or finding threats.
- 1 Return of the Wildspeaker | draw | Flexible draw support for a board built around large creatures.
- 1 Ripjaw Raptor | draw | A dinosaur draw option that fits the tribal plan.
- 1 Rishkar's Expertise | draw | A high-impact draw spell for a deck with large threats.
- 1 Shamanic Revelation | draw | Draw support that rewards developing a creature board.
- 1 Toski, Bearer of Secrets | draw | A creature-based draw engine for the combat plan.
- 1 Vanquisher's Banner | draw | Tribal draw support that belongs in a dinosaur deck.
- 1 Boros Charm | interaction | A flexible interaction spell for protecting the deck’s position.
- 1 Heroic Intervention | interaction | Protection interaction for keeping key permanents in play.
- 1 Inspiring Call | interaction | Protective interaction that works with the creature strategy.
- 1 Lightning Greaves | interaction | Equipment interaction for protecting an important creature.
- 1 Swiftfoot Boots | interaction | A simple protection tool for the commander or a major threat.
- 1 Temple Altisaur | interaction | A dinosaur interaction piece that supports the creature board.
- 12 Forest | land | Basic green land for reliable access to green mana.
- 8 Mountain | land | Basic red land for reliable access to red mana.
- 7 Plains | land | Basic white land for reliable access to white mana.
- 1 Canopy Vista | land | A land slot that supports the deck’s three-color mana base.
- 1 Cinder Glade | land | A land slot that supports the deck’s three-color mana base.
- 1 Clifftop Retreat | land | A land slot that supports the deck’s three-color mana base.
- 1 Command Tower | land | A land slot for the commander’s three-color deck.
- 1 Evolving Wilds | land | A land slot that helps establish the mana base.
- 1 Path of Ancestry | land | A tribal land choice for the dinosaur deck.
- 1 Rootbound Crag | land | A land slot that supports the deck’s three-color mana base.
- 1 Sunpetal Grove | land | A land slot that supports the deck’s three-color mana base.
- 1 Terramorphic Expanse | land | A land slot that helps establish the mana base.
- 1 Arcane Signet | ramp | Simple mana acceleration for casting the deck’s larger creatures.
- 1 Cultivate | ramp | Reliable land-based ramp for reaching the deck’s expensive threats.
- 1 Drover of the Mighty | ramp | A creature ramp piece that fits beside the dinosaur plan.
- 1 Farseek | ramp | Early ramp to help deploy the commander and threats sooner.
- 1 Intrepid Paleontologist | ramp | A themed creature ramp option for the dinosaur strategy.
- 1 Kodama's Reach | ramp | Reliable land-based ramp for the three-color mana base.
- 1 Nature's Lore | ramp | Efficient ramp for developing mana early.
- 1 Rampant Growth | ramp | Straightforward ramp that helps fix the mana base.
- 1 Sol Ring | ramp | Fast mana acceleration for getting the deck started.
- 1 Thunderherd Migration | ramp | Dinosaur-themed ramp that supports the tribal plan.
- 1 Apex Altisaur | removal | A dinosaur removal option that also contributes to the creature plan.
- 1 Bronzebeak Foragers | removal | A dinosaur-based removal tool for handling opposing cards.
- 1 Burning Sun's Avatar | removal | A sizable dinosaur that also fills a removal role.
- 1 Deathgorge Scavenger | removal | A dinosaur removal option with a useful supporting role.
- 1 Itzquinth, Firstborn of Gishath | removal | A thematic removal piece for the dinosaur deck.
- 1 Ravenous Sailback | removal | A dinosaur removal option that remains on-theme.
- 1 Savage Stomp | removal | A focused removal spell for clearing a problem creature.
- 1 Thrashing Brontodon | removal | A dinosaur removal piece that fits naturally in the deck.
- 1 Commune with Dinosaurs | synergy | Tribal support that helps the deck stay focused on dinosaurs.
- 1 Hunting Velociraptor | synergy | A dinosaur synergy creature for advancing the main plan.
- 1 Kinjalli's Caller | synergy | A tribal support creature for deploying dinosaurs.
- 1 Marauding Raptor | synergy | A dinosaur synergy card that reinforces the creature-focused strategy.
- 1 Orazca Frillback | synergy | A dinosaur synergy option that keeps the deck’s theme dense.
- 1 Otepec Huntmaster | synergy | A tribal support creature for the dinosaur plan.
- 1 Pugnacious Hammerskull | synergy | An on-theme dinosaur that supports the creature strategy.
- 1 Raptor Companion | synergy | A low-cost dinosaur that builds the tribal board.
- 1 Raptor Hatchling | synergy | A themed dinosaur that contributes to the tribe’s board presence.
- 1 Regal Imperiosaur | synergy | A dinosaur synergy creature for the deck’s central theme.
- 1 Siegehorn Ceratops | synergy | An on-theme dinosaur that supports the tribal plan.
- 1 Sun-Collared Raptor | synergy | A dinosaur synergy piece that adds to the creature base.
- 1 Sunfrill Imitator | synergy | A tribal dinosaur card that supports the deck’s plan.
- 1 Territorial Hammerskull | synergy | A dinosaur synergy creature for the board-focused strategy.
- 1 Ancient Brontodon | threat | A large dinosaur threat for closing games through combat.
- 1 Carnage Tyrant | threat | A major dinosaur threat that puts pressure on opponents.
- 1 Etali, Primal Storm | threat | A powerful dinosaur threat at the top of the curve.
- 1 Ghalta and Mavren | threat | A legendary dinosaur threat for the deck’s combat plan.
- 1 Ghalta, Primal Hunger | threat | A huge dinosaur threat for ending games quickly.
- 1 Goring Ceratops | threat | A combat-focused dinosaur threat for the creature board.
- 1 Palani's Hatcher | threat | A dinosaur threat that adds weight to the top end.
- 1 Pantlaza, Sun-Favored | threat | A dinosaur threat that keeps the deck strongly tribal.
- 1 Quartzwood Crasher | threat | A dinosaur threat that supports aggressive combat turns.
- 1 Regisaur Alpha | threat | A key dinosaur threat for building battlefield pressure.
- 1 Thundering Spineback | threat | A large dinosaur threat for the deck’s late game.
- 1 Zetalpa, Primal Dawn | threat | A high-end dinosaur threat for finishing games.
- 1 Austere Command | wipe | A flexible board wipe for resetting difficult positions.
- 1 Blasphemous Act | wipe | A board wipe for catching up when the battlefield gets crowded.
- 1 Wakening Sun's Avatar | wipe | A dinosaur-themed board wipe that fits the deck’s plan.

</details>

### 5. blink Commander, owned first

Format: Commander. Theme: blink. Pool: owned_first. Shortlist: 173 names.

Cards: 99. Repair turn: false. Block findings: 0.

**Summary:** This Brago blink deck develops with ramp and draw, then uses its synergy pieces alongside removal, interaction, and wipes to keep the table manageable. It closes by applying pressure with its threats after the board has been controlled. The tradeoff is a patient, support-heavy plan rather than an all-in attacking strategy.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [WARN] `not_owned`: Brago, King Eternal: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 2.95 over 62 nonland cards
- [INFO] `basics_added`: the list was 1 cards short, so the builder added 1 basic lands

<details><summary>The deck list</summary>

- 1 Abandoned Air Temple | land | Provides a land slot in the mana base.
- 1 Ash Barrens | land | Provides a land slot in the mana base.
- 1 City of Brass | land | Provides a land slot in the mana base.
- 1 Command Tower | land | Provides a land slot in the mana base.
- 1 Escape Tunnel | land | Provides a land slot in the mana base.
- 1 Evolving Wilds | land | Provides a land slot in the mana base.
- 1 Exotic Orchard | land | Provides a land slot in the mana base.
- 1 Fabled Passage | land | Provides a land slot in the mana base.
- 1 Field of Ruin | land | Provides a land slot in the mana base.
- 1 Ghost Quarter | land | Provides a land slot in the mana base.
- 1 Glacial Fortress | land | Provides a land slot in the mana base.
- 1 Great Hall of the Citadel | land | Provides a land slot in the mana base.
- 1 Irrigated Farmland | land | Provides a land slot in the mana base.
- 5 Island | land | Provides basic blue mana sources.
- 1 Marsh Flats | land | Provides a land slot in the mana base.
- 1 Minas Tirith | land | Provides a land slot in the mana base.
- 1 Opal Palace | land | Provides a land slot in the mana base.
- 1 Path of Ancestry | land | Provides a land slot in the mana base.
- 4 Plains | land | Provides basic white mana sources.
- 1 Plaza of Heroes | land | Provides a land slot in the mana base.
- 1 Port Town | land | Provides a land slot in the mana base.
- 1 Prairie Stream | land | Provides a land slot in the mana base.
- 1 Reliquary Tower | land | Provides a land slot in the mana base.
- 1 Rogue's Passage | land | Provides a land slot in the mana base.
- 1 Scavenger Grounds | land | Provides a land slot in the mana base.
- 1 Secluded Courtyard | land | Provides a land slot in the mana base.
- 1 Spire of Industry | land | Provides a land slot in the mana base.
- 1 Temple of the False God | land | Provides a land slot in the mana base.
- 1 Terramorphic Expanse | land | Provides a land slot in the mana base.
- 1 Tranquil Cove | land | Provides a land slot in the mana base.
- 1 Birthday Escape | draw | Adds a dedicated draw option.
- 1 Champions of Minas Tirith | draw | Adds a dedicated draw option.
- 1 Exemplar of Light | draw | Adds a dedicated draw option.
- 1 Faramir, Steward of Gondor | draw | Adds a dedicated draw option.
- 1 Idol of Oblivion | draw | Adds a dedicated draw option.
- 1 Ingenious Prodigy | draw | Adds a dedicated draw option.
- 1 Inspiring Overseer | draw | Adds a dedicated draw option.
- 1 Lembas | draw | Adds a dedicated draw option.
- 1 Mask of Memory | draw | Adds a dedicated draw option.
- 1 Puresteel Paladin | draw | Adds a dedicated draw option.
- 1 Summon: Shiva | draw | Adds a dedicated draw option.
- 1 Tome of Legends | draw | Adds a dedicated draw option.
- 1 Wall of Omens | draw | Adds a dedicated draw option.
- 1 Waterbender Ascension | draw | Adds a dedicated draw option.
- 1 Waterbending Lesson | draw | Adds a dedicated draw option.
- 1 Arcane Denial | interaction | Supplies interaction for opposing plays.
- 1 Boromir, Warden of the Tower | interaction | Supplies interaction for opposing plays.
- 1 Bronze Guardian | interaction | Supplies interaction for opposing plays.
- 1 Champion's Helm | interaction | Supplies interaction for opposing plays.
- 1 Clever Concealment | interaction | Supplies interaction for opposing plays.
- 1 Darksteel Plate | interaction | Supplies interaction for opposing plays.
- 1 Dovin's Veto | interaction | Supplies interaction for opposing plays.
- 1 Lightning Greaves | interaction | Supplies interaction for opposing plays.
- 1 Slip On the Ring | interaction | Supplies interaction that fits the blink plan.
- 1 Swiftfoot Boots | interaction | Supplies interaction for opposing plays.
- 1 Aang's Journey | ramp | Adds mana development for the early game.
- 1 Arcane Signet | ramp | Adds mana development for the early game.
- 1 Chromatic Lantern | ramp | Adds mana development for the early game.
- 1 Commander's Sphere | ramp | Adds mana development for the early game.
- 1 Fellwar Stone | ramp | Adds mana development for the early game.
- 1 Ioreth of the Healing House | ramp | Adds mana development for the early game.
- 1 Relic of Legends | ramp | Adds mana development for the early game.
- 1 Ring of the Lucii | ramp | Adds mana development for the early game.
- 1 Sol Ring | ramp | Adds mana development for the early game.
- 1 Sword of the Animist | ramp | Adds mana development for the early game.
- 1 Talisman of Progress | ramp | Adds mana development for the early game.
- 1 Thought Vessel | ramp | Adds mana development for the early game.
- 1 Wayfarer's Bauble | ramp | Adds mana development for the early game.
- 1 White Lotus Tile | ramp | Adds mana development for the early game.
- 1 Aang, the Last Airbender | removal | Provides removal within the creature-focused plan.
- 1 Angel of Condemnation | removal | Provides removal that supports the blink theme.
- 1 Angel of Serenity | removal | Provides removal within the creature-focused plan.
- 1 Banishing Light | removal | Provides a dedicated removal answer.
- 1 Crib Swap | removal | Provides a dedicated removal answer.
- 1 Fiend Hunter | removal | Provides removal within the creature-focused plan.
- 1 Generous Gift | removal | Provides a dedicated removal answer.
- 1 Get Lost | removal | Provides a dedicated removal answer.
- 1 Palace Jailer | removal | Provides removal within the creature-focused plan.
- 1 Swords to Plowshares | removal | Provides a dedicated removal answer.
- 1 Ty Lee, Chi Blocker | removal | Provides removal within the creature-focused plan.
- 1 Westfold Rider | removal | Provides removal within the creature-focused plan.
- 1 Austere Command | wipe | Provides a broad reset when individual answers are not enough.
- 1 Perplexing Test | wipe | Provides a broad reset when individual answers are not enough.
- 1 Supreme Verdict | wipe | Provides a broad reset when individual answers are not enough.
- 1 Ennis, Debate Moderator | synergy | Adds a listed synergy piece for the blink plan.
- 1 Flickerwisp | synergy | Adds a listed synergy piece for the blink plan.
- 1 Jocasta, Automaton Avenger | synergy | Adds a listed synergy piece for the blink plan.
- 1 Personify | synergy | Adds a listed synergy piece for the blink plan.
- 1 Roll-Roll-Roll-Roll | synergy | Adds a listed synergy piece for the blink plan.
- 1 S.H.I.E.L.D. Flying Car | synergy | Adds a listed synergy piece for the blink plan.
- 1 Lake-town Mariners // Gone Fishing | threat | Supplies a threat for closing games.
- 1 Meneldor, Swift Savior | threat | Supplies a threat for closing games.

</details>

### 6. Modern tempo, tournament

Format: Modern. Theme: tempo. Pool: any_card. Shortlist: 167 names.

Cards: 60. Repair turn: false. Block findings: 0.

**Summary:** This deck plays a blue-red tempo game by developing a compact threat package while using draw, interaction, and removal to keep the opponent from stabilizing. It wins by preserving pressure through efficient exchanges and carrying that momentum into its temporal synergy cards. In return, it gives up the slower, larger battlefield plans available elsewhere in the shortlist in favor of a focused proactive game.

- [INFO] `curve_summary`: average mana value 2.00 over 36 nonland cards

<details><summary>The deck list</summary>

- 4 Island | land | Basic blue land for the mana base.
- 4 Mountain | land | Basic red land for the mana base.
- 4 Steam Vents | land | Blue-red land for consistent two-color access.
- 4 Scalding Tarn | land | Blue-red fetch land for the mana base.
- 4 Shivan Reef | land | Blue-red land that supports the tempo curve.
- 4 Stormcarved Coast | land | Blue-red land for the mana base.
- 4 Consider | draw | Efficient draw to keep spells flowing.
- 2 Preordain | draw | Additional draw to improve early consistency.
- 4 Counterspell | interaction | Core interaction for protecting the tempo plan.
- 2 Spell Pierce | interaction | Low-cost interaction for early exchanges.
- 4 Lightning Bolt | removal | Efficient removal that supports a proactive plan.
- 2 Pongify | removal | Additional creature removal.
- 2 Into the Flood Maw | removal | Flexible removal for clearing a path.
- 4 Temporal Mastery | synergy | Synergy piece that rewards maintaining momentum.
- 4 Ledger Shredder | threat | Creature threat that fits the spell-heavy tempo shell.
- 4 Faerie Mastermind | threat | Blue creature threat for applying pressure.
- 4 Ragavan, Nimble Pilferer | threat | Red creature threat for the proactive early game.

</details>

### 7. Modern burn, casual

Format: Modern. Theme: burn. Pool: any_card. Shortlist: 290 names.

Cards: 60. Repair turn: false. Block findings: 0.

**Summary:** This deck starts with cheap red pressure, backs it with direct-damage removal, and uses its spell-focused creatures and enchantment to make every low-cost play more punishing. It wins by steadily reducing the opponent's life total, then keeps the assault going with card draw and larger red threats when the game lasts longer. It gives up broad answers and a flexible mana base in exchange for a focused, straightforward mono-red plan.

- [INFO] `curve_summary`: average mana value 2.44 over 36 nonland cards

<details><summary>The deck list</summary>

- 18 Mountain | land | The basic red mana base for casting the deck's spells.
- 4 Ramunap Ruins | land | A red land slot that fits the deck's aggressive plan.
- 2 Sunscorched Desert | land | A land slot that contributes to the mono-red mana base.
- 4 Risk Factor | draw | A draw spell that helps keep pressure flowing.
- 2 Grab the Prize | draw | Additional draw to find more action.
- 2 Flick a Coin | ramp | A ramp card for supporting the deck's mana needs.
- 4 Lightning Bolt | removal | Efficient removal for clearing blockers or pressuring the opponent.
- 2 Burst Lightning | removal | Extra removal that suits a burn-focused shell.
- 4 Eidolon of the Great Revel | synergy | A synergy piece that reinforces the deck's low-cost pressure plan.
- 4 Kessig Flamebreather | synergy | A synergy creature for a spell-heavy red deck.
- 4 Heartfire Hero | threat | An early threat that helps establish pressure.
- 4 Akki Lavarunner // Tok-Tok, Volcano Born | threat | A threat that adds to the creature pressure plan.
- 2 Hazoret the Fervent | threat | A durable top-end threat for games that go longer.
- 2 Torbran, Thane of Red Fell | threat | A threat that supports the deck's damage-oriented plan.
- 2 Tectonic Giant | threat | A larger threat to maintain pressure after the early turns.

</details>

### 8. Modern lifegain, FNM

Format: Modern. Theme: lifegain. Pool: any_card. Shortlist: 301 names.

Cards: 60. Repair turn: false. Block findings: 0.

**Summary:** This white-black lifegain deck builds around Soul Warden and Ajani's Pridemate, then develops a creature board with Attended Healer, Bloodbond Vampire, and its other threats. It wins by applying steady board pressure while using removal to clear opposing obstacles, with draw and ramp helping it continue deploying threats. It gives up some flexibility for a focused creature-and-synergy plan, so the sideboard carries the extra removal, sweepers, and interaction for more demanding matchups.

- [INFO] `curve_summary`: average mana value 3.00 over 36 nonland cards

<details><summary>The deck list</summary>

- 10 Plains | land | Provides the white mana base.
- 10 Swamp | land | Provides the black mana base.
- 4 Scoured Barrens | land | Adds to the white-black land base.
- 1 Orzhov Keyrune | ramp | Provides ramp for the deck's higher-cost cards.
- 1 Pristine Talisman | ramp | Provides a second ramp piece.
- 1 Dawn of Hope | draw | Fills a draw slot in the lifegain shell.
- 2 Inspiring Overseer | draw | Supplies draw while contributing to the creature plan.
- 2 Survival Cache | draw | Adds efficient draw support.
- 1 Well of Lost Dreams | draw | Adds another draw engine for the deck.
- 4 Murderous Rider // Swift End | removal | Provides flexible removal slots.
- 2 Nightmare's Thirst | removal | Rounds out the removal package.
- 4 Soul Warden | synergy | Forms a central lifegain synergy package.
- 4 Ajani's Pridemate | synergy | Works with the deck's lifegain-focused synergy.
- 4 Attended Healer | threat | Provides a substantial creature threat package.
- 4 Bloodbond Vampire | threat | Adds pressure as part of the threat suite.
- 2 Cliffhaven Vampire | threat | Supports the white-black lifegain threat plan.
- 2 Twinblade Paladin | threat | Adds more creature-based pressure.
- 2 Angel of Invention | threat | Finishes out the creature threat package.

</details>

### 9. Standard midrange, FNM

Format: Standard. Theme: midrange. Pool: any_card. Shortlist: 164 names.

Cards: 60. Repair turn: false. Block findings: 0.

**Summary:** This black-green midrange deck builds its mana, develops a steady flow of cards, and uses removal to keep opposing plans from taking over the table. It aims to establish a resilient board presence, clear away pressure with its answer suite and wipes when needed, then press the advantage in a longer game. The tradeoff is that much of the list is dedicated to support, interaction, and broad answers rather than a dense package of explicitly designated finishers.

- [INFO] `curve_summary`: average mana value 2.83 over 36 nonland cards

<details><summary>The deck list</summary>

- 8 Forest | land | Provides a basic land slot in the mana base.
- 8 Swamp | land | Provides a basic land slot in the mana base.
- 4 Blooming Marsh | land | Provides a black-green land option.
- 2 Overgrown Tomb | land | Provides a black-green land option.
- 2 Underground Mortuary | land | Provides a black-green land option.
- 3 Phyrexian Arena | draw | Supplies a repeatable draw element for longer games.
- 3 Midnight Reaper | draw | Adds draw to the creature package.
- 2 Darkstar Augur | draw | Adds another draw-focused creature.
- 2 Unholy Annex // Ritual Chamber | draw | Expands the deck's draw package.
- 2 Stocking the Pantry | draw | Provides additional draw support.
- 2 Llanowar Elves | ramp | Gives the deck early ramp.
- 2 Archdruid's Charm | ramp | Adds flexible ramp support.
- 2 Aftermath Analyst | ramp | Provides further ramp in the creature package.
- 3 Bitter Triumph | removal | Forms a core set of removal answers.
- 2 Assassin's Trophy | removal | Adds broad removal coverage.
- 1 Scavenging Ooze | removal | Provides removal on a creature.
- 2 Royal Treatment | interaction | Supplies interaction for protecting the board plan.
- 2 Blossoming Defense | interaction | Adds low-cost interaction support.
- 2 Not Dead After All | interaction | Provides interaction for a creature-focused board.
- 2 Blasphemous Edict | wipe | Gives the deck access to a wipe when the board gets crowded.
- 2 Season of Loss | wipe | Adds more wipe coverage for difficult boards.
- 2 Villainous Wrath | wipe | Rounds out the main-deck wipe package.

</details>

### 10. Standard aggro, tournament

Format: Standard. Theme: aggro. Pool: any_card. Shortlist: 294 names.

Cards: 60. Repair turn: false. Block findings: 0.

**Summary:** This red-white aggro deck establishes its mana, commits threats to the table, and keeps pressure on the opponent with removal and interaction clearing the path. It wins by maintaining that pressure until its threats finish the game, while draw and synergy help the attack stay focused. It gives up a broader late-game plan in favor of a direct, proactive game plan.

- [INFO] `curve_summary`: average mana value 2.83 over 36 nonland cards

<details><summary>The deck list</summary>

- 10 Mountain | land | Provides red basic-land slots for the mana base.
- 8 Plains | land | Provides white basic-land slots for the mana base.
- 4 Sacred Foundry | land | Provides red-white land slots for the mana base.
- 2 Inspiring Vantage | land | Provides additional red-white land slots for the mana base.
- 4 Fugitive Codebreaker | draw | Adds draw while fitting the aggressive plan.
- 2 Reckless Lackey | draw | Adds draw support to keep the deck moving.
- 4 Boros Charm | interaction | Supplies interaction for protecting the deck's pressure.
- 2 Sheltered by Ghosts | interaction | Adds interaction alongside the creature-focused plan.
- 4 Harsh Annotation | removal | Provides efficient removal support for clearing the way.
- 4 Emeritus of Truce // Swords to Plowshares | removal | Adds removal to answer opposing permanents.
- 4 Warleader's Call | synergy | Serves as the deck's focused synergy piece.
- 4 Dragonback Lancer | threat | Provides a threat for the aggressive game plan.
- 4 Redcap Gutter-Dweller | threat | Provides another threat to sustain pressure.
- 4 Teapot Slinger | threat | Rounds out the threat base for attacking decks.

</details>

### 11. Commander with a locked card

Format: Commander. Theme: sacrifice. Pool: any_card. Shortlist: 296 names.

Cards: 99. Repair turn: false. Block findings: 0.

**Summary:** This Karlov deck builds a board around sacrifice synergies, using its creatures, engines, and value pieces to turn those sacrifices into pressure while maintaining cards and mana. It wins by pairing those synergies with its larger threats and by clearing away opposing resources through removal and wipes. The tradeoff is that the deck relies on assembling and keeping a sacrifice engine plus a board presence, so concentrated disruption can slow its momentum.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 3.19 over 63 nonland cards

<details><summary>The deck list</summary>

- 10 Plains | land | Provides white mana for the deck.
- 12 Swamp | land | Provides black mana for the deck.
- 1 Arguel's Blood Fast // Temple of Aclazotz | land | Provides a land slot from the shortlist.
- 1 Bloodstained Mire | land | Provides a land slot from the shortlist.
- 1 Bojuka Bog | land | Provides a land slot from the shortlist.
- 1 Command Tower | land | Provides a land slot from the shortlist.
- 1 Evolving Wilds | land | Provides a land slot from the shortlist.
- 1 Exotic Orchard | land | Provides a land slot from the shortlist.
- 1 High Market | land | Provides a land slot for the sacrifice-focused plan.
- 1 Myriad Landscape | land | Provides a land slot from the shortlist.
- 1 Path of Ancestry | land | Provides a land slot from the shortlist.
- 1 Phyrexian Tower | land | Provides a land slot for the sacrifice-focused plan.
- 1 Reliquary Tower | land | Provides a land slot from the shortlist.
- 1 Scavenger Grounds | land | Provides a land slot from the shortlist.
- 1 Terramorphic Expanse | land | Provides a land slot from the shortlist.
- 1 Westvale Abbey // Ormendahl, Profane Prince | land | Provides a land slot for the sacrifice-focused plan.
- 1 Sol Ring | ramp | Provides ramp and stays in the deck as requested.
- 1 Ashnod's Altar | ramp | Provides ramp for the sacrifice-focused plan.
- 1 Phyrexian Altar | ramp | Provides ramp for the sacrifice-focused plan.
- 1 Pitiless Plunderer | ramp | Provides ramp for the sacrifice-focused plan.
- 1 Pawn of Ulamog | ramp | Provides ramp for the sacrifice-focused plan.
- 1 Sifter of Skulls | ramp | Provides ramp for the sacrifice-focused plan.
- 1 Culling the Weak | ramp | Provides ramp for the sacrifice-focused plan.
- 1 Crowded Crypt | ramp | Provides ramp for the sacrifice-focused plan.
- 1 Priest of Forgotten Gods | ramp | Provides ramp for the sacrifice-focused plan.
- 1 Warren Soultrader | ramp | Provides ramp for the sacrifice-focused plan.
- 1 Baron Bertram Graywater | draw | Provides card draw for the deck.
- 1 Bushmeat Poacher | draw | Provides card draw for the sacrifice-focused plan.
- 1 Corrupted Conviction | draw | Provides card draw for the sacrifice-focused plan.
- 1 Disciple of Bolas | draw | Provides card draw for the sacrifice-focused plan.
- 1 Ecstatic Awakener // Awoken Demon | draw | Provides card draw for the sacrifice-focused plan.
- 1 Lord Skitter's Butcher | draw | Provides card draw for the sacrifice-focused plan.
- 1 Smothering Abomination | draw | Provides card draw for the sacrifice-focused plan.
- 1 Tevesh Szat, Doom of Fools | draw | Provides card draw for the deck.
- 1 Vampiric Rites | draw | Provides card draw for the sacrifice-focused plan.
- 1 Village Rites | draw | Provides card draw for the sacrifice-focused plan.
- 1 Cartel Aristocrat | interaction | Provides interaction for the sacrifice-focused plan.
- 1 Dark Privilege | interaction | Provides interaction for the sacrifice-focused plan.
- 1 Eldrazi Monument | interaction | Provides interaction for the sacrifice-focused plan.
- 1 Fanatical Devotion | interaction | Provides interaction for the sacrifice-focused plan.
- 1 Flare of Fortitude | interaction | Provides interaction for the deck.
- 1 Gift of Doom | interaction | Provides interaction for the sacrifice-focused plan.
- 1 Attrition | removal | Provides removal for the sacrifice-focused plan.
- 1 Ayli, Eternal Pilgrim | removal | Provides removal for the sacrifice-focused plan.
- 1 Bone Shards | removal | Provides removal for the sacrifice-focused plan.
- 1 Dictate of Erebos | removal | Provides removal for the sacrifice-focused plan.
- 1 Eaten Alive | removal | Provides removal for the sacrifice-focused plan.
- 1 Grave Pact | removal | Provides removal for the sacrifice-focused plan.
- 1 Teysa, Orzhov Scion | removal | Provides removal for the sacrifice-focused plan.
- 1 Yawgmoth, Thran Physician | removal | Provides removal for the sacrifice-focused plan.
- 1 Altar of Dementia | synergy | Supports the sacrifice plan.
- 1 Bartolomé del Presidio | synergy | Supports the sacrifice plan.
- 1 Bastion of Remembrance | synergy | Supports the sacrifice plan.
- 1 Bloodflow Connoisseur | synergy | Supports the sacrifice plan.
- 1 Carrion Feeder | synergy | Supports the sacrifice plan.
- 1 Elas il-Kor, Sadistic Pilgrim | synergy | Supports the sacrifice plan.
- 1 Hidden Stockpile | synergy | Supports the sacrifice plan.
- 1 Open the Graves | synergy | Supports the sacrifice plan.
- 1 Spawning Pit | synergy | Supports the sacrifice plan.
- 1 Viscera Seer | synergy | Supports the sacrifice plan.
- 1 Witch's Oven | synergy | Supports the sacrifice plan.
- 1 Woe Strider | synergy | Supports the sacrifice plan.
- 1 Yahenni, Undying Partisan | synergy | Supports the sacrifice plan.
- 1 Zulaport Cutthroat | synergy | Supports the sacrifice plan.
- 1 Abhorrent Overlord | threat | Provides a major threat for the deck.
- 1 Corpse Harvester | threat | Provides a threat for the sacrifice-focused plan.
- 1 Ghoulcaller Gisa | threat | Provides a threat for the sacrifice-focused plan.
- 1 Kuldotha Forgemaster | threat | Provides a threat for the deck.
- 1 Liesa, Forgotten Archangel | threat | Provides a threat for the deck.
- 1 Liliana, Heretical Healer // Liliana, Defiant Necromancer | threat | Provides a threat for the sacrifice-focused plan.
- 1 Mondrak, Glory Dominus | threat | Provides a threat for the deck.
- 1 Ratadrabik of Urborg | threat | Provides a threat for the deck.
- 1 Razaketh, the Foulblooded | threat | Provides a major threat for the sacrifice-focused plan.
- 1 Requiem Angel | threat | Provides a threat for the deck.
- 1 Sidisi, Undead Vizier | threat | Provides a threat for the sacrifice-focused plan.
- 1 Vindictive Vampire | threat | Provides a threat for the sacrifice-focused plan.
- 1 Elesh Norn // The Argent Etchings | wipe | Provides a board wipe for the deck.
- 1 The Meathook Massacre | wipe | Provides a board wipe for the deck.
- 1 Toxic Deluge | wipe | Provides a board wipe for the deck.

</details>

### 12. Commander on a budget

Format: Commander. Theme: tokens. Pool: any_card. Shortlist: 296 names.

Cards: 99. Repair turn: false. Block findings: 0.

**Summary:** This Adeline-led token deck aims to establish a broad board, reinforce that plan with closely aligned cards, and turn sustained attacks into its primary win. It has enough resource development and answers to keep moving through a typical table, but gives up some flexibility and individual card power in favor of staying focused on the token strategy.

- [INFO] `bracket_prose_rules`: bracket 2 (Core, 8+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-24
- [INFO] `curve_summary`: average mana value 3.21 over 63 nonland cards

<details><summary>The deck list</summary>

- 36 Plains | land | Provides the deck's basic land base.
- 1 Angelic Sell-Sword | draw | Included as a draw option that supports continued resources.
- 1 Bygone Bishop | draw | Included as a draw option that supports continued resources.
- 1 Chivalric Alliance | draw | Included as a draw option that supports continued resources.
- 1 Court of Grace | draw | Included as a draw option that supports continued resources.
- 1 Dawn of Hope | draw | Included as a draw option that supports continued resources.
- 1 Faramir, Field Commander | draw | Included as a draw option that supports continued resources.
- 1 Idol of Oblivion | draw | Included as a draw option that supports continued resources.
- 1 Platoon Dispenser | draw | Included as a draw option that supports continued resources.
- 1 Staff of the Storyteller | draw | Included as a draw option that supports continued resources.
- 1 Wedding Announcement // Wedding Festivity | draw | Included as a draw option that supports continued resources.
- 1 Ainok Strike Leader | interaction | Provides interaction that supports the board-focused plan.
- 1 Appa, Steadfast Guardian | interaction | Provides interaction that supports the board-focused plan.
- 1 Blessed Sanctuary | interaction | Provides interaction that supports the board-focused plan.
- 1 Rootborn Defenses | interaction | Provides interaction that supports the board-focused plan.
- 1 Spirit Bonds | interaction | Provides interaction that supports the board-focused plan.
- 1 Teyo, the Shieldmage | interaction | Provides interaction that supports the board-focused plan.
- 1 Charisma Bobblehead | ramp | Included to support the deck's resource development.
- 1 Coin of Mastery | ramp | Included to support the deck's resource development.
- 1 Collector's Vault | ramp | Included to support the deck's resource development.
- 1 Currency Converter | ramp | Included to support the deck's resource development.
- 1 Druidic Satchel | ramp | Included to support the deck's resource development.
- 1 Goldvein Pick | ramp | Included to support the deck's resource development.
- 1 Karn, Living Legacy | ramp | Included to support the deck's resource development.
- 1 Keeper of the Accord | ramp | Included to support the deck's resource development.
- 1 Noble's Purse | ramp | Included to support the deck's resource development.
- 1 Prying Blade | ramp | Included to support the deck's resource development.
- 1 Aerial Assault | removal | Provides focused removal for opposing problems.
- 1 Banishing Slash | removal | Provides focused removal for opposing problems.
- 1 Citizen's Crowbar | removal | Provides focused removal for opposing problems.
- 1 Hanged Executioner | removal | Provides focused removal for opposing problems.
- 1 Kellan's Lightblades | removal | Provides focused removal for opposing problems.
- 1 Release to Memory | removal | Provides focused removal for opposing problems.
- 1 Trostani's Judgment | removal | Provides focused removal for opposing problems.
- 1 The Wandering Emperor | removal | Provides focused removal for opposing problems.
- 1 Animation Module | synergy | Supports the deck's token-centered theme.
- 1 Anointer Priest | synergy | Supports the deck's token-centered theme.
- 1 Automated Assembly Line | synergy | Supports the deck's token-centered theme.
- 1 Cat Collector | synergy | Supports the deck's token-centered theme.
- 1 Clarion Spirit | synergy | Supports the deck's token-centered theme.
- 1 Felidar Retreat | synergy | Supports the deck's token-centered theme.
- 1 Intangible Virtue | synergy | Supports the deck's token-centered theme.
- 1 Oketra's Monument | synergy | Supports the deck's token-centered theme.
- 1 Retrofitter Foundry | synergy | Supports the deck's token-centered theme.
- 1 Siege Veteran | synergy | Supports the deck's token-centered theme.
- 1 Skrelv's Hive | synergy | Supports the deck's token-centered theme.
- 1 Spawning Pit | synergy | Supports the deck's token-centered theme.
- 1 Three Blind Mice | synergy | Supports the deck's token-centered theme.
- 1 Worthy Knight | synergy | Supports the deck's token-centered theme.
- 1 Attended Healer | threat | Serves as a board-building threat.
- 1 Basri's Lieutenant | threat | Serves as a board-building threat.
- 1 Cemetery Protector | threat | Serves as a board-building threat.
- 1 Custodi Soulbinders | threat | Serves as a board-building threat.
- 1 Dragonback Lancer | threat | Serves as a board-building threat.
- 1 Emeria Angel | threat | Serves as a board-building threat.
- 1 Hero of Bladehold | threat | Serves as a board-building threat.
- 1 Illustrious Wanderglyph | threat | Serves as a board-building threat.
- 1 Lossarnach Captain | threat | Serves as a board-building threat.
- 1 Mite Overseer | threat | Serves as a board-building threat.
- 1 Phantom General | threat | Serves as a board-building threat.
- 1 Silverwing Squadron | threat | Serves as a board-building threat.
- 1 Hour of Reckoning | wipe | Provides a reset when the board gets out of hand.
- 1 The Battle of Bywater | wipe | Provides a reset when the board gets out of hand.
- 1 Elspeth, Sun's Champion | wipe | Provides a reset when the board gets out of hand.

</details>

