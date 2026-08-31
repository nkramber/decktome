# PR-8 deck gate

Run date: 2026-08-31. Card snapshot: 2026-08-31.

Verdict: PASS. 18 of 18 decks passed every block check, 0 invented names reached the user, and 0 summaries stated a false rule of the game. All three bars are zero tolerance.

## Summary

| Measure | Value |
|---|---|
| Prompts | 18 |
| Decks returned | 18 |
| Decks with no block finding | 18 |
| Invented names that reached the user | 0 |
| Decks that needed the repair turn | 1 |
| Summaries judged (F-26) | 18 |
| Summaries that state a rule of the game | 1 |
| Summaries that state a FALSE rule | 0 |
| Judge errors | 0 |
| Errors | 0 |
| Prompt version | 10 |
| Calls | 37 |
| Cost | $1.0450 |
| Time | 786 seconds |

## Findings by code

A block stops the deck. A warning and a note are reports.

| Code | Count |
|---|---|
| `not_owned` | 106 |
| `curve_summary` | 18 |
| `bracket_prose_rules` | 12 |
| `basics_added` | 2 |
| `land_count` | 1 |

By severity: BLOCK 0. WARN 107. INFO 32. 

## Decks

### 1. lifegain Commander, any card

Format: Commander. Theme: lifegain. Pool: any_card. Shortlist: 303 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $403.84 to buy, $403.84 the whole deck.

**Summary:** This is a lifegain-focused Karlov deck that develops its mana and cards, then uses lifegain synergies to build a threatening board and apply steady pressure. It wins through its collection of powerful creatures and lifegain-linked threats while using removal, protective interaction, and wipes to keep opposing boards manageable. The deck gives up dedicated alternate win conditions in favor of a consistent creature-and-synergy game plan.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 3.43 over 63 nonland cards

<details><summary>The deck list</summary>

- 18 Plains | land | Provides white mana for the deck's lifegain-focused spells.
- 18 Swamp | land | Provides black mana for the deck's lifegain-focused spells.
- 1 Archivist of Oghma | draw | Provides card draw for the deck.
- 1 Convalescent Care | draw | Provides card draw for the lifegain plan.
- 1 Cosmos Elixir | draw | Provides card draw for the deck.
- 1 Dawn of Hope | draw | Provides card draw for the lifegain plan.
- 1 Enduring Innocence | draw | Provides card draw for the deck.
- 1 Exemplar of Light | draw | Provides card draw for the deck.
- 1 Mangara, the Diplomat | draw | Provides card draw for the deck.
- 1 Sigarda's Splendor | draw | Provides card draw for the lifegain plan.
- 1 The Gaffer | draw | Provides card draw for the deck.
- 1 Well of Lost Dreams | draw | Provides card draw for the lifegain plan.
- 1 Alseid of Life's Bounty | interaction | Adds interaction that helps protect the deck's plan.
- 1 Faith's Shield | interaction | Adds interaction that helps protect the deck's plan.
- 1 Metropolis Reformer | interaction | Adds interaction to support the board.
- 1 Restoration Magic | interaction | Adds interaction to support the board.
- 1 Sword of Light and Shadow | interaction | Adds interaction to support the board.
- 1 Werefox Bodyguard | interaction | Adds interaction to support the board.
- 1 Altar of the Pantheon | ramp | Provides ramp for the deck.
- 1 Battle Angels of Tyr | ramp | Provides ramp while contributing to the board.
- 1 Crypt Ghast | ramp | Provides ramp for the deck.
- 1 Hierophant's Chalice | ramp | Provides ramp for the deck.
- 1 Nuka-Cola Vending Machine | ramp | Provides ramp for the deck.
- 1 Orazca Relic | ramp | Provides ramp for the deck.
- 1 Phial of Galadriel | ramp | Provides ramp for the deck.
- 1 Pristine Talisman | ramp | Provides ramp for the lifegain plan.
- 1 Redemption Choir | ramp | Provides ramp for the deck.
- 1 The Celestus | ramp | Provides ramp for the deck.
- 1 Aetherflux Reservoir | removal | Provides removal for problematic opposing pieces.
- 1 Ayli, Eternal Pilgrim | removal | Provides removal for problematic opposing pieces.
- 1 Murderous Rider // Swift End | removal | Provides removal for problematic opposing pieces.
- 1 Nightmare's Thirst | removal | Provides removal for problematic opposing pieces.
- 1 Solitude | removal | Provides removal for problematic opposing pieces.
- 1 Umezawa's Jitte | removal | Provides removal for problematic opposing pieces.
- 1 Vona, Butcher of Magan | removal | Provides removal while contributing to the board.
- 1 Witch of the Moors | removal | Provides removal while supporting the lifegain plan.
- 1 Aerith Gainsborough | synergy | Supports the deck's lifegain synergies.
- 1 Ajani's Pridemate | synergy | Supports the deck's lifegain synergies.
- 1 Angelic Accord | synergy | Supports the deck's lifegain synergies.
- 1 Angel of Vitality | synergy | Supports the deck's lifegain synergies.
- 1 Blood Artist | synergy | Supports the deck's lifegain synergies.
- 1 Bloodthirsty Aerialist | synergy | Supports the deck's lifegain synergies.
- 1 Cleric Class | synergy | Supports the deck's lifegain synergies.
- 1 Heliod, Sun-Crowned | synergy | Supports the deck's lifegain synergies.
- 1 Resplendent Angel | synergy | Supports the deck's lifegain synergies.
- 1 Righteous Valkyrie | synergy | Supports the deck's lifegain synergies.
- 1 Serra Ascendant | synergy | Supports the deck's lifegain synergies.
- 1 Vito, Thorn of the Dusk Rose | synergy | Supports the deck's lifegain synergies.
- 1 Vizkopa Guildmage | synergy | Supports the deck's lifegain synergies.
- 1 Voice of the Blessed | synergy | Supports the deck's lifegain synergies.
- 1 Archangel of Thune | threat | A major threat that benefits from the lifegain plan.
- 1 Astarion, the Decadent | threat | A major threat for closing games.
- 1 Celestine, the Living Saint | threat | A major threat for the lifegain-focused board.
- 1 Cliffhaven Vampire | threat | A threat that supports the lifegain plan.
- 1 Defiant Bloodlord | threat | A threat that supports the lifegain plan.
- 1 Divinity of Pride | threat | A major threat for closing games.
- 1 Elenda, Saint of Dusk | threat | A major threat for the lifegain-focused board.
- 1 Lyra Dawnbringer | threat | A major threat for the lifegain-focused board.
- 1 Nykthos Paragon | threat | A threat that supports the lifegain plan.
- 1 Rhox Faithmender | threat | A major threat for the lifegain-focused board.
- 1 Valkyrie Harbinger | threat | A major threat for the lifegain-focused board.
- 1 Wurmcoil Engine | threat | A major threat for closing games.
- 1 Ajani, Strength of the Pride | wipe | Provides a board wipe when the board gets out of hand.
- 1 Fumigate | wipe | Provides a board wipe when the board gets out of hand.
- 1 Kaya's Wrath | wipe | Provides a board wipe when the board gets out of hand.

</details>

### 2. aristocrats Commander, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 303 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $39.81 to buy, $216.44 the whole deck.

**Summary:** This is a board-focused Orzhov aristocrats deck that develops mana, deploys sacrifice outlets and death-trigger payoffs, then turns expendable creatures into steady pressure and resource flow. It wins by compounding those sacrifice and death synergies, with large black creatures providing a secondary way to close a stalled game. The deck gives up some immediate speed for an incremental plan that is strongest when it can keep a creature board in play and grind through attrition.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [WARN] `not_owned`: Blood Artist: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Carrion Feeder: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Cruel Celebrant: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Elas il-Kor, Sadistic Pilgrim: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Zulaport Cutthroat: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Viscera Seer: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Woe Strider: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Yahenni, Undying Partisan: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bastion of Remembrance: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Fleshtaker: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bartolomé del Presidio: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Pitiless Pontiff: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Pious Evangel // Wayward Disciple: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Victimize: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Apocalypse Demon: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Blood Host: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Demonlord of Ashmouth: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Falkenrath Noble: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ghoulcaller Gisa: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Morkrut Necropod: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Novice Dissector: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Old Flitterfang: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Razaketh, the Foulblooded: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Venerated Stormsinger: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vindictive Vampire: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vito's Inquisitor: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 2.89 over 63 nonland cards

<details><summary>The deck list</summary>

- 12 Plains | land | Provides a reliable white mana base.
- 12 Swamp | land | Provides a reliable black mana base.
- 1 Command Tower | land | Provides flexible commander-color mana.
- 1 Marsh Flats | land | Helps assemble the deck's two-color mana base.
- 1 Path of Ancestry | land | Provides colored mana for the creature-heavy plan.
- 1 Secluded Courtyard | land | Supports the deck's many creature types with colored mana.
- 1 Plaza of Heroes | land | Adds flexible mana for the legendary commander and creatures.
- 1 Castle Locthwain | land | Adds a black mana source with late-game utility.
- 1 Takenuma, Abandoned Mire | land | Adds a black mana source with graveyard utility.
- 1 War Room | land | Adds a mana source with card-access utility.
- 1 Evolving Wilds | land | Fixes the basic-land mana base.
- 1 Fabled Passage | land | Fixes the basic-land mana base.
- 1 Ash Barrens | land | Helps find the needed basic land colors.
- 1 Field of Ruin | land | Provides a utility land slot in the mana base.
- 1 Skullclamp | draw | Turns expendable creatures into a strong card-flow engine.
- 1 Idol of Oblivion | draw | Provides repeatable card access alongside token-making and sacrifice plans.
- 1 Lembas | draw | Supplies a compact source of card access.
- 1 Night's Whisper | draw | Provides efficient early card access.
- 1 Painful Truths | draw | Refills the hand efficiently.
- 1 Tome of Legends | draw | Provides steady card access over a long game.
- 1 Wall of Omens | draw | Adds an early blocker that replaces itself.
- 1 Call of the Ring | draw | Provides ongoing card access.
- 1 Buster Sword | draw | Adds card access from an equipment slot.
- 1 Stone of Erech | draw | Supplies additional card access while filling a utility slot.
- 1 Boromir, Warden of the Tower | interaction | Protects the deck's board-based plan from opposing disruption.
- 1 Clever Concealment | interaction | Helps preserve the board through opposing answers.
- 1 Darksteel Plate | interaction | Protects an important creature or commander.
- 1 Lightning Greaves | interaction | Protects a key creature while enabling immediate use.
- 1 Swiftfoot Boots | interaction | Provides another flexible protection piece.
- 1 Together Forever | interaction | Helps preserve important creatures through attrition.
- 1 Arcane Signet | ramp | Provides efficient color fixing.
- 1 Sol Ring | ramp | Accelerates the deck's early development.
- 1 Fellwar Stone | ramp | Adds inexpensive mana acceleration.
- 1 Wayfarer's Bauble | ramp | Develops the basic-land mana base.
- 1 Commander's Sphere | ramp | Provides flexible mana acceleration.
- 1 Thought Vessel | ramp | Adds colorless mana acceleration.
- 1 Relic of Legends | ramp | Turns the deck's legendary creatures into mana support.
- 1 Astral Cornucopia | ramp | Provides scalable mana acceleration.
- 1 Bender's Waterskin | ramp | Fills a compact mana-acceleration slot.
- 1 Deadly Dispute | ramp | Converts a disposable permanent into resources while accelerating.
- 1 Bitter Triumph | removal | Offers flexible creature or planeswalker removal.
- 1 Banishing Light | removal | Answers a problematic opposing permanent.
- 1 Generous Gift | removal | Provides broad permanent removal.
- 1 Get Lost | removal | Efficiently answers key opposing permanents.
- 1 Infernal Grasp | removal | Provides dependable creature removal.
- 1 Swords to Plowshares | removal | Efficiently removes a dangerous creature.
- 1 Path to Exile | removal | Provides another efficient creature answer.
- 1 Fiend Hunter | removal | Puts creature removal on a body that can join the sacrifice plan.
- 1 Blood Artist | synergy | Rewards the deck for creatures dying.
- 1 Carrion Feeder | synergy | Provides a low-cost sacrifice outlet.
- 1 Cruel Celebrant | synergy | Reinforces the deck's death-trigger plan.
- 1 Elas il-Kor, Sadistic Pilgrim | synergy | Supports the core sacrifice-and-death strategy.
- 1 Zulaport Cutthroat | synergy | Turns creature deaths into incremental pressure.
- 1 Viscera Seer | synergy | Provides a cheap sacrifice outlet and card selection.
- 1 Woe Strider | synergy | Adds another sacrifice outlet for the creature package.
- 1 Yahenni, Undying Partisan | synergy | Provides a resilient sacrifice outlet.
- 1 Bastion of Remembrance | synergy | Adds a durable death-trigger payoff.
- 1 Fleshtaker | synergy | Rewards sacrificing creatures while supporting the central plan.
- 1 Bartolomé del Presidio | synergy | Adds a reliable sacrifice outlet to the board.
- 1 Pitiless Pontiff | synergy | Turns spare creatures into protection for a key attacker or blocker.
- 1 Pious Evangel // Wayward Disciple | synergy | Contributes another creature-death payoff.
- 1 Victimize | synergy | Reuses important creatures while working with sacrifice fodder.
- 1 Apocalypse Demon | threat | Provides a large black finisher for the deck's upper curve.
- 1 Blood Host | threat | Grows into a meaningful threat as creatures are sacrificed.
- 1 Demonlord of Ashmouth | threat | Provides a substantial creature threat that fits the sacrifice plan.
- 1 Falkenrath Noble | threat | Pressures opponents as creatures die.
- 1 Ghoulcaller Gisa | threat | Provides a high-impact creature threat for a sacrifice-focused board.
- 1 Morkrut Necropod | threat | Adds a large black creature to close games.
- 1 Novice Dissector | threat | Adds another creature threat to the sacrifice shell.
- 1 Old Flitterfang | threat | Provides a black creature threat for the deck's creature package.
- 1 Razaketh, the Foulblooded | threat | Acts as a powerful top-end threat for long games.
- 1 Venerated Stormsinger | threat | Adds pressure while fitting the deck's black creature base.
- 1 Vindictive Vampire | threat | Turns creature deaths into a threatening clock.
- 1 Vito's Inquisitor | threat | Adds another vampire threat to pressure opponents.
- 1 Austere Command | wipe | Provides a flexible reset when opponents get ahead.
- 1 Fumigate | wipe | Resets creature-heavy opposing boards.
- 1 Dusk // Dawn | wipe | Offers a board reset with useful follow-up value.

</details>

### 3. artifacts Commander, bracket 4

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 299 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $4758.33 to buy, $4758.33 the whole deck.

**Summary:** This deck centers its game on artifact development, using ramp, draw, and dedicated synergy pieces to establish a powerful board quickly. It wins by turning that artifact infrastructure into sustained pressure from a dense threat package, while removal, interaction, and sweepers keep opposing plans in check. Its tradeoff is a strongly artifact-focused construction rather than a broad nonartifact game plan.

- [INFO] `curve_summary`: average mana value 3.90 over 63 nonland cards

<details><summary>The deck list</summary>

- 20 Island | land | Provides the deck’s basic blue mana base.
- 1 Academy Ruins | land | Adds a utility land slot to the artifact-focused mana base.
- 1 Archway of Innovation | land | Adds a land slot to support the artifact-focused plan.
- 1 Blinkmoth Nexus | land | Provides a utility land slot in the mana base.
- 1 Buried Ruin | land | Adds artifact-oriented utility to a land slot.
- 1 Conqueror's Galleon // Conqueror's Foothold | land | Fills a flexible land slot for the deck.
- 1 Hall of Tagsin | land | Adds an artifact-oriented land slot.
- 1 Inventors' Fair | land | Provides a utility land for the artifact shell.
- 1 Mishra's Workshop | land | Supplies a powerful artifact-oriented land slot.
- 1 Otawara, Soaring City | land | Adds a blue utility land to the mana base.
- 1 Power Depot | land | Provides an artifact land slot.
- 1 Roadside Reliquary | land | Adds a utility land slot for the deck’s permanent-heavy construction.
- 1 Spire of Industry | land | Supports the artifact-focused mana base.
- 1 The Mycosynth Gardens | land | Adds an artifact-oriented utility land.
- 1 Treasure Map // Treasure Cove | land | Provides a flexible land slot.
- 1 Urza's Saga | land | Adds an Urza-themed artifact land slot.
- 1 Urza's Workshop | land | Provides an artifact-oriented land slot.
- 1 Chief Engineer | ramp | Supports the deck’s artifact acceleration package.
- 1 Grand Architect | ramp | Provides artifact-focused mana acceleration.
- 1 Inspiring Statuary | ramp | Contributes to the deck’s artifact ramp plan.
- 1 Karn, Legacy Reforged | ramp | Adds a ramp piece that fits the artifact shell.
- 1 Krark-Clan Ironworks | ramp | Provides high-power artifact ramp.
- 1 Metalworker | ramp | Supplies artifact-focused mana acceleration.
- 1 Moonsnare Prototype | ramp | Adds a compact artifact ramp piece.
- 1 Mox Opal | ramp | Provides efficient artifact ramp.
- 1 Tezzeret the Seeker | ramp | Adds a ramp planeswalker that supports artifacts.
- 1 The Mightstone and Weakstone | ramp | Provides an artifact ramp slot.
- 1 Forensic Gadgeteer | draw | Adds a card-advantage piece to the artifact shell.
- 1 One with the Machine | draw | Provides a draw spell for the deck’s artifact plan.
- 1 Reverse Engineer | draw | Adds artifact-oriented card draw.
- 1 Riddlesmith | draw | Provides repeatable draw support for the artifact shell.
- 1 Sai, Master Thopterist | draw | Adds an artifact-focused draw engine.
- 1 Tezzeret, Artifice Master | draw | Provides a draw planeswalker for the deck.
- 1 Thirst for Knowledge | draw | Adds efficient card draw for an artifact-heavy build.
- 1 Thought Monitor | draw | Provides an artifact creature draw slot.
- 1 Thoughtcast | draw | Adds efficient artifact-oriented card draw.
- 1 Vedalken Archmage | draw | Provides draw support for the artifact plan.
- 1 Assert Authority | interaction | Adds stack interaction to protect the deck’s plan.
- 1 Disruption Protocol | interaction | Provides efficient interaction for an artifact deck.
- 1 Ice Out | interaction | Adds a flexible interaction spell.
- 1 Metallic Rebuke | interaction | Provides artifact-oriented stack interaction.
- 1 Padeem, Consul of Innovation | interaction | Adds an artifact-focused interaction piece.
- 1 Welding Jar | interaction | Provides compact interaction for the artifact suite.
- 1 Aetherflux Reservoir | removal | Adds artifact-based removal to the deck.
- 1 Arcum Dagsson | removal | Provides artifact-focused removal.
- 1 Portal to Phyrexia | removal | Adds a high-impact artifact removal piece.
- 1 Ravenform | removal | Provides a flexible removal spell.
- 1 Resculpt | removal | Adds efficient removal to the interaction suite.
- 1 Skysovereign, Consul Flagship | removal | Provides artifact-based removal while fitting the deck’s theme.
- 1 Spine of Ish Sah | removal | Adds a versatile artifact removal slot.
- 1 Transmogrifying Wand | removal | Provides repeatable artifact-based removal.
- 1 Emry, Lurker of the Loch | synergy | Provides core artifact synergy.
- 1 Etherium Sculptor | synergy | Supports the deck’s artifact synergy plan.
- 1 Foundry Inspector | synergy | Adds another artifact-focused synergy piece.
- 1 Manifold Key | synergy | Provides compact artifact synergy.
- 1 Mystic Forge | synergy | Adds a high-power artifact synergy engine.
- 1 Power Artifact | synergy | Provides a focused artifact synergy piece.
- 1 Scrap Trawler | synergy | Adds artifact recursion synergy.
- 1 Shimmer Myr | synergy | Provides artifact-focused synergy.
- 1 Simulacrum Synthesizer | synergy | Adds an artifact synergy engine.
- 1 Transmute Artifact | synergy | Provides a powerful artifact synergy spell.
- 1 Unwinding Clock | synergy | Supports the deck’s artifact-focused board development.
- 1 Voltaic Key | synergy | Adds a compact artifact synergy piece.
- 1 Whir of Invention | synergy | Provides artifact-focused synergy and consistency.
- 1 Clock of Omens | synergy | Adds a dedicated artifact synergy engine.
- 1 Broodstar | threat | Provides a large artifact-themed threat.
- 1 Cyberdrive Awakener | threat | Adds a powerful artifact-focused threat.
- 1 Darksteel Juggernaut | threat | Provides a resilient artifact creature threat.
- 1 Kappa Cannoneer | threat | Adds a high-power artifact creature threat.
- 1 Karn, Scion of Urza | threat | Provides a threat planeswalker for the artifact shell.
- 1 Kuldotha Forgemaster | threat | Adds a major artifact creature threat.
- 1 Lodestone Golem | threat | Provides an artifact creature threat.
- 1 Master Transmuter | threat | Adds an artifact-focused creature threat.
- 1 Mycosynth Golem | threat | Provides a large artifact creature threat.
- 1 Phyrexian Metamorph | threat | Adds a flexible artifact creature threat.
- 1 The Capitoline Triad | threat | Provides a top-end artifact threat.
- 1 Traxos, Scourge of Kroog | threat | Adds an efficient legendary artifact threat.
- 1 Engineered Explosives | wipe | Provides an artifact-based board wipe.
- 1 Nevinyrral's Disk | wipe | Adds a broad artifact board wipe.
- 1 Oblivion Stone | wipe | Provides a third artifact-based board wipe.

</details>

### 4. dinosaur tribal, bracket 2

Format: Commander. Theme: dinosaurs. Pool: any_card. Shortlist: 269 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $271.97 to buy, $271.97 the whole deck.

**Summary:** This is a straightforward Dinosaur creature deck that develops its mana, builds a board of tribal support creatures and large threats, then wins by applying sustained combat pressure. It has card draw to keep the creatures coming, removal and board wipes to clear obstacles, and protection for pivotal turns. It gives up some flexibility for a focused, creature-heavy Dinosaur plan.

- [INFO] `bracket_prose_rules`: bracket 2 (Core, 8+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 3.92 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Beast Whisperer | draw | Draw support for the creature-heavy Dinosaur plan.
- 1 Garruk's Uprising | draw | Draw support that fits a deck built around substantial creatures.
- 1 Guardian Project | draw | Reliable draw support as the deck develops its creature board.
- 1 Harmonize | draw | Straightforward card draw for refilling the hand.
- 1 Kutzil, Malamet Exemplar | draw | Creature-based draw support for the combat-focused plan.
- 1 Return of the Wildspeaker | draw | Flexible draw support for a board of Dinosaur creatures.
- 1 Ripjaw Raptor | draw | Dinosaur draw support that contributes to the board.
- 1 Rishkar's Expertise | draw | High-impact draw spell for rebuilding a hand.
- 1 Shamanic Revelation | draw | Draw support that rewards developing a creature board.
- 1 Vanquisher's Banner | draw | Tribal draw support for a Dinosaur-focused deck.
- 1 Akroma's Will | interaction | Combat-focused protection and interaction for a key turn.
- 1 Boros Charm | interaction | Versatile protection for the creature board.
- 1 Flawless Maneuver | interaction | Protection to help preserve the Dinosaur board.
- 1 Heroic Intervention | interaction | Broad protection against opposing answers.
- 1 Lightning Greaves | interaction | Protects an important creature while supporting the attack plan.
- 1 Swiftfoot Boots | interaction | Protects a key creature and supports immediate pressure.
- 1 Battlefield Forge | land | Color-producing land for the three-color mana base.
- 1 Canopy Vista | land | Green and white source for the mana base.
- 1 Cinder Glade | land | Red and green source for the mana base.
- 1 Clifftop Retreat | land | Red and white source for the mana base.
- 1 Command Tower | land | Reliable three-color land for the commander deck.
- 1 Evolving Wilds | land | Flexible land fixing for the three-color mana base.
- 1 Exotic Orchard | land | Flexible color source for multiplayer games.
- 9 Forest | land | Basic green sources for the mana base.
- 1 Myriad Landscape | land | Land slot that supports the deck's mana development.
- 6 Mountain | land | Basic red sources for the mana base.
- 1 Path of Ancestry | land | Tribal land that supports the Dinosaur theme.
- 5 Plains | land | Basic white sources for the mana base.
- 1 Rootbound Crag | land | Red and green source for the mana base.
- 1 Sacred Foundry | land | Red and white source for the mana base.
- 1 Spectator Seating | land | Red and white source for multiplayer games.
- 1 Stomping Ground | land | Red and green source for the mana base.
- 1 Sunpetal Grove | land | Green and white source for the mana base.
- 1 Temple Garden | land | Green and white source for the mana base.
- 1 Terramorphic Expanse | land | Flexible land fixing for the three-color mana base.
- 1 Arcane Signet | ramp | Simple color fixing and mana acceleration.
- 1 Cultivate | ramp | Reliable land-based ramp and color fixing.
- 1 Drover of the Mighty | ramp | Creature-based mana support for the Dinosaur deck.
- 1 Farseek | ramp | Early land ramp that helps fix colors.
- 1 Kodama's Reach | ramp | Reliable land-based ramp and color fixing.
- 1 Nature's Lore | ramp | Efficient land ramp for reaching larger Dinosaurs.
- 1 Ranging Raptors | ramp | Dinosaur ramp support that contributes to the board.
- 1 Sol Ring | ramp | Straightforward early mana acceleration.
- 1 Thunderherd Migration | ramp | Dinosaur-themed land ramp and color fixing.
- 1 Topiary Stomper | ramp | Creature-based land ramp for the larger creature plan.
- 1 Apex Altisaur | removal | Dinosaur removal that also serves as a sizable creature.
- 1 Burning Sun's Avatar | removal | Dinosaur-based removal for opposing threats.
- 1 Itzquinth, Firstborn of Gishath | removal | Low-cost Dinosaur removal support.
- 1 Ravenous Sailback | removal | Flexible Dinosaur removal attached to a creature.
- 1 Savage Stomp | removal | Dinosaur-themed targeted removal.
- 1 Thrashing Brontodon | removal | Creature-based removal that fits the theme.
- 1 Tranquil Frillback | removal | Flexible Dinosaur removal for troublesome permanents.
- 1 Trumpeting Carnosaur | removal | Dinosaur removal that remains a meaningful threat.
- 1 Armored Kincaller | synergy | Early Dinosaur creature that supports the tribal plan.
- 1 Belligerent Yearling | synergy | Low-cost Dinosaur support for building the board.
- 1 Commune with Dinosaurs | synergy | Dinosaur-themed consistency support.
- 1 Deathgorge Scavenger | synergy | Dinosaur utility creature for the tribal plan.
- 1 Dinosaur Stampede | synergy | Combat support for a board of Dinosaurs.
- 1 Huatli's Raptor | synergy | Low-cost Dinosaur creature that supports the theme.
- 1 Hunting Velociraptor | synergy | Dinosaur support creature for the attack plan.
- 1 Kinjalli's Caller | synergy | Dinosaur-focused support for deploying larger creatures.
- 1 Kinjalli's Sunwing | synergy | Dinosaur utility creature that supports combat pressure.
- 1 Marauding Raptor | synergy | Dinosaur support creature for an aggressive board.
- 1 Otepec Huntmaster | synergy | Dinosaur-focused support for the creature plan.
- 1 Raptor Companion | synergy | Simple early Dinosaur creature for the tribal board.
- 1 Sunfrill Imitator | synergy | Dinosaur support creature that rewards the tribal theme.
- 1 Territorial Hammerskull | synergy | Dinosaur utility creature that supports attacking.
- 1 Carnage Tyrant | threat | A major Dinosaur threat for closing games through combat.
- 1 Etali, Primal Storm | threat | High-impact legendary Dinosaur threat.
- 1 Ghalta and Mavren | threat | Large Dinosaur threat that advances the combat plan.
- 1 Ghalta, Stampede Tyrant | threat | Powerful Dinosaur threat for a developed board.
- 1 Pantlaza, Sun-Favored | threat | Legendary Dinosaur threat that reinforces the theme.
- 1 Quartzwood Crasher | threat | Combat-oriented Dinosaur threat.
- 1 Regisaur Alpha | threat | Dinosaur threat that supports a wide creature board.
- 1 Sawhorn Nemesis | threat | Dinosaur threat for applying pressure.
- 1 Tyrranax Rex | threat | Large Dinosaur threat for finishing games.
- 1 Verdant Sun's Avatar | threat | Large Dinosaur threat for stabilizing the board presence.
- 1 Zetalpa, Primal Dawn | threat | High-end Dinosaur threat for combat finishes.
- 1 Zilortha, Strength Incarnate | threat | Legendary Dinosaur threat for the creature-heavy plan.
- 1 Austere Command | wipe | Flexible board wipe for recovering from difficult boards.
- 1 Blasphemous Act | wipe | Efficient reset button when creature boards get out of hand.
- 1 Wakening Sun's Avatar | wipe | Dinosaur-themed board wipe that fits the deck's creature plan.

</details>

### 5. blink Commander, owned first

Format: Commander. Theme: blink. Pool: owned_first. Shortlist: 302 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $80.24 to buy, $219.40 the whole deck.

**Summary:** This blink deck develops a creature-heavy board, repeatedly leans on its blink synergy, and uses creature-based draw and removal to keep momentum through the middle turns. It wins by turning that established board and its larger threats into sustained pressure while holding interaction and wipes for pivotal opposing plays. The deck gives up color flexibility and a broader range of threat options in exchange for a focused, library-first white blink shell.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [WARN] `not_owned`: Akroma's Will: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Aethergeode Miner: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Against All Odds: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Another Round: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Charming Prince: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Conjurer's Closet: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Daydream: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Far Traveler: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Flicker: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Golden Argosy: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Guardian of Ghirapur: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Panharmonicon: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Phelia, Exuberant Shepherd: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Teleportation Circle: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vizier of Deferment: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Wingrattle Scarecrow: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Befriending the Moths // Imperial Moth: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Elesh Norn, Mother of Machines: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Felidar Guardian: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Glimmerpoint Stag: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Icewind Stalwart: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Photon, Lady of Light: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Rattleblaze Scarecrow: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Twilight Shepherd: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Wispweaver Angel: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.19 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Abandoned Air Temple | land | Land for the deck’s mana base.
- 1 Adventurer's Inn | land | Land for the deck’s mana base.
- 1 Ash Barrens | land | Land for the deck’s mana base.
- 1 Bonders' Enclave | land | Land for the deck’s mana base.
- 1 City of Brass | land | Land for the deck’s mana base.
- 1 Crossroads Village | land | Land for the deck’s mana base.
- 1 Eclipsed Realms | land | Land for the deck’s mana base.
- 1 Eden, Seat of the Sanctum | land | Land for the deck’s mana base.
- 1 Escape Tunnel | land | Land for the deck’s mana base.
- 1 Evolving Wilds | land | Land for the deck’s mana base.
- 1 Exotic Orchard | land | Land for the deck’s mana base.
- 1 Fabled Passage | land | Land for the deck’s mana base.
- 1 Field of Ruin | land | Land for the deck’s mana base.
- 1 Fountainport | land | Land for the deck’s mana base.
- 1 Ghost Quarter | land | Land for the deck’s mana base.
- 1 Great Hall of the Citadel | land | Land for the deck’s mana base.
- 1 Minas Tirith | land | Land for the deck’s mana base.
- 1 Nesting Grounds | land | Land for the deck’s mana base.
- 1 Opal Palace | land | Land for the deck’s mana base.
- 1 Path of Ancestry | land | Land for the deck’s mana base.
- 1 Plaza of Heroes | land | Land for the deck’s mana base.
- 1 Reliquary Tower | land | Land for the deck’s mana base.
- 1 Riveteers Overlook | land | Land for the deck’s mana base.
- 1 Rogue's Passage | land | Land for the deck’s mana base.
- 1 Rumble Arena | land | Land for the deck’s mana base.
- 1 Scavenger Grounds | land | Land for the deck’s mana base.
- 1 Secluded Courtyard | land | Land for the deck’s mana base.
- 1 Shire Terrace | land | Land for the deck’s mana base.
- 1 Spire of Industry | land | Land for the deck’s mana base.
- 1 Study Hall | land | Land for the deck’s mana base.
- 1 Temple of the False God | land | Land for the deck’s mana base.
- 1 Terramorphic Expanse | land | Land for the deck’s mana base.
- 1 The Gold Saucer | land | Land for the deck’s mana base.
- 1 The Grey Havens | land | Land for the deck’s mana base.
- 1 Vibrant Cityscape | land | Land for the deck’s mana base.
- 1 Windbrisk Heights | land | Land for the deck’s mana base.
- 1 Adventurer's Airship | draw | Draw support for keeping the blink plan supplied.
- 1 Buster Sword | draw | Draw support for keeping the blink plan supplied.
- 1 Champions of Minas Tirith | draw | Draw support that also fits the creature-focused deck.
- 1 Crown of Gondor | draw | Draw support for keeping the blink plan supplied.
- 1 Diary of Dreams | draw | Draw support for keeping the blink plan supplied.
- 1 Exemplar of Light | draw | Creature-based draw that fits the deck’s board plan.
- 1 Faramir, Field Commander | draw | Creature-based draw that fits the deck’s board plan.
- 1 Idol of Oblivion | draw | Draw support for keeping the blink plan supplied.
- 1 Inspiring Overseer | draw | Creature-based draw that fits the deck’s board plan.
- 1 Lembas | draw | Draw support for keeping the blink plan supplied.
- 1 Arcane Signet | ramp | Reliable ramp for developing the board early.
- 1 Astral Cornucopia | ramp | Ramp for advancing the deck’s development.
- 1 Chromatic Lantern | ramp | Ramp for advancing the deck’s development.
- 1 Commander's Sphere | ramp | Ramp for advancing the deck’s development.
- 1 Fellwar Stone | ramp | Ramp for advancing the deck’s development.
- 1 Giada, Font of Hope | ramp | Creature-based ramp that fits the deck’s creature plan.
- 1 Relic of Legends | ramp | Ramp for advancing the deck’s development.
- 1 Sol Ring | ramp | Ramp for advancing the deck’s development.
- 1 Thought Vessel | ramp | Ramp for advancing the deck’s development.
- 1 Wayfarer's Bauble | ramp | Ramp for advancing the deck’s development.
- 1 Angel of Sanctions | removal | Creature-based removal that fits the blink plan.
- 1 Banishing Light | removal | Flexible removal for opposing permanents.
- 1 Crib Swap | removal | Focused removal for opposing creatures.
- 1 Destroy Evil | removal | Flexible removal for opposing permanents.
- 1 Fiend Hunter | removal | Creature-based removal that fits the blink plan.
- 1 Generous Gift | removal | Flexible removal for opposing permanents.
- 1 Get Lost | removal | Flexible removal for opposing permanents.
- 1 Swords to Plowshares | removal | Focused removal for opposing creatures.
- 1 Austere Command | wipe | Board-reset option for difficult positions.
- 1 Fumigate | wipe | Board-reset option for difficult positions.
- 1 Vanquish the Horde | wipe | Board-reset option for difficult positions.
- 1 Akroma's Will | interaction | Interaction for protecting the deck’s board plan.
- 1 Boromir, Warden of the Tower | interaction | Creature-based interaction that supports the board plan.
- 1 Clever Concealment | interaction | Interaction for protecting the deck’s board plan.
- 1 Crystal Fragments // Summon: Alexander | interaction | Flexible interaction for the deck’s game plan.
- 1 Lightning Greaves | interaction | Interaction that helps safeguard key creatures.
- 1 Slip On the Ring | interaction | Interaction that aligns with the blink plan.
- 1 Aethergeode Miner | synergy | Blink synergy piece for the deck’s core plan.
- 1 Against All Odds | synergy | Blink synergy piece for the deck’s core plan.
- 1 Another Round | synergy | Blink synergy piece for the deck’s core plan.
- 1 Charming Prince | synergy | Creature-based blink synergy for the deck.
- 1 Conjurer's Closet | synergy | Dedicated blink synergy for the deck’s core plan.
- 1 Daydream | synergy | Blink synergy piece for the deck’s core plan.
- 1 Ennis, Debate Moderator | synergy | Synergy support for the deck’s central plan.
- 1 Far Traveler | synergy | Dedicated blink synergy for the deck’s core plan.
- 1 Flicker | synergy | Blink synergy piece for the deck’s core plan.
- 1 Flickerwisp | synergy | Creature-based blink synergy for the deck.
- 1 Golden Argosy | synergy | Dedicated blink synergy for the deck’s core plan.
- 1 Guardian of Ghirapur | synergy | Creature-based blink synergy for the deck.
- 1 Panharmonicon | synergy | Synergy support for the deck’s creature plan.
- 1 Phelia, Exuberant Shepherd | synergy | Creature-based blink synergy for the deck.
- 1 Teleportation Circle | synergy | Dedicated blink synergy for the deck’s core plan.
- 1 Vizier of Deferment | synergy | Creature-based blink synergy for the deck.
- 1 Wingrattle Scarecrow | synergy | Synergy support for the deck’s central plan.
- 1 Befriending the Moths // Imperial Moth | threat | Threat that helps close games through the deck’s board presence.
- 1 Elesh Norn, Mother of Machines | threat | High-impact threat for the deck’s creature plan.
- 1 Felidar Guardian | threat | Threat that also fits the blink-focused game plan.
- 1 Glimmerpoint Stag | threat | Threat that also fits the blink-focused game plan.
- 1 Icewind Stalwart | threat | Threat for converting the established board into pressure.
- 1 Photon, Lady of Light | threat | High-impact threat for the deck’s creature plan.
- 1 Rattleblaze Scarecrow | threat | Threat for converting the established board into pressure.
- 1 Twilight Shepherd | threat | High-impact threat for the deck’s creature plan.
- 1 Wispweaver Angel | threat | Threat that also fits the blink-focused game plan.

</details>

### 6. Modern tempo, tournament

Format: Modern. Theme: tempo. Pool: any_card. Shortlist: 167 names.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $316.16 to buy, $316.16 the whole deck.

**Summary:** This blue-red tempo deck establishes pressure with Ledger Shredder, Faerie Mastermind, and Ragavan, Nimble Pilferer, then uses efficient removal and interaction to keep that pressure ahead of opposing development. Card-selection spells help it find the right threat or answer, while Temporal Mastery supplies the deck's temporal-spell angle. It wins by maintaining a quick board presence and protecting it rather than by committing to a large, expensive endgame.

- [INFO] `curve_summary`: average mana value 2.00 over 36 nonland cards

<details><summary>The deck list</summary>

- 8 Island | land | Provides the primary blue mana base.
- 4 Mountain | land | Provides red mana for the deck's red spells.
- 4 Steam Vents | land | Provides both deck colors.
- 4 Shivan Reef | land | Provides both deck colors.
- 4 Stormcarved Coast | land | Provides both deck colors.
- 4 Consider | draw | Efficient blue card selection and draw.
- 2 Preordain | draw | Adds early card selection and draw.
- 4 Counterspell | interaction | Core protection and disruption for a tempo plan.
- 2 Spell Pierce | interaction | Low-cost interaction to maintain tempo.
- 4 Lightning Bolt | removal | Efficient removal that supports early pressure.
- 4 Into the Flood Maw | removal | Low-cost removal for clearing a path.
- 4 Temporal Mastery | synergy | Supports the deck's temporal-spell synergy package.
- 4 Ledger Shredder | threat | An early creature that applies pressure while supporting the spell-heavy plan.
- 4 Faerie Mastermind | threat | A low-cost evasive creature for the tempo offense.
- 4 Ragavan, Nimble Pilferer | threat | An early creature that adds pressure and supports the deck's resource plan.

</details>

### 7. Modern burn, casual

Format: Modern. Theme: burn. Pool: any_card. Shortlist: 290 names.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $47.12 to buy, $47.12 the whole deck.

**Summary:** This mono-red burn deck applies pressure with direct removal and burn-oriented synergy creatures, then closes through a concentrated creature threat package. Card draw helps it keep pressuring after the first wave, while a small amount of ramp supports its heavier threats. It gives up sideboard flexibility and broad answers for a focused, all-red aggressive plan.

- [INFO] `curve_summary`: average mana value 2.94 over 36 nonland cards

<details><summary>The deck list</summary>

- 24 Mountain | land | Provides the mono-red mana base.
- 4 Browbeat | draw | Provides a draw option for the burn plan.
- 2 Risk Factor | draw | Adds further draw support.
- 2 Ancestors' Aid | ramp | Fills the requested ramp slot.
- 4 Lightning Bolt | removal | Efficiently fills a removal role in a burn-focused deck.
- 2 Shock | removal | Adds inexpensive removal to support early pressure.
- 4 Eidolon of the Great Revel | synergy | Supports the deck's burn-focused synergy plan.
- 4 Guttersnipe | synergy | Rewards the deck for leaning on burn spells.
- 4 Ashcloud Phoenix | threat | Provides a creature threat for the attacking plan.
- 4 Hazoret the Fervent | threat | Adds a resilient top-end threat.
- 4 Torbran, Thane of Red Fell | threat | Provides a powerful mono-red threat.
- 2 Stormbreath Dragon | threat | Rounds out the threat package with evasive creature pressure.

</details>

### 8. Modern lifegain, FNM

Format: Modern. Theme: lifegain. Pool: any_card. Shortlist: 301 names.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $158.56 to buy, $158.56 the whole deck.

**Summary:** This white-black lifegain deck starts with Soul Warden and Ajani's Pridemate, then builds a board of creatures that benefit from repeated life gain. It wins by turning those incremental gains into increasingly dangerous creature threats, backed by efficient removal to clear the way. The deck gives up some speed and relies on keeping its lifegain engine on the table, but its draw package helps it continue applying pressure into longer FNM games.

- [INFO] `curve_summary`: average mana value 2.76 over 34 nonland cards
- [INFO] `basics_added`: the list was 2 cards short, so the builder added 2 basic lands

<details><summary>The deck list</summary>

- 11 Plains | land | Provides reliable white mana for the deck’s early lifegain core.
- 7 Swamp | land | Provides black mana for the deck’s black spells and threats.
- 4 Scoured Barrens | land | Supports both colors while fitting the lifegain plan.
- 2 Shambling Vent | land | Provides white-black mana and remains a useful late-game land.
- 2 Restless Fortress | land | Provides white-black mana and gives the deck an additional land-based threat.
- 2 Arcane Signet | ramp | Accelerates the deck into its more expensive threats and payoffs.
- 4 Path to Exile | removal | Efficiently clears opposing creatures that would stop the lifegain board from attacking.
- 2 Murderous Rider // Swift End | removal | Adds flexible removal while leaving behind a creature threat.
- 4 Soul Warden | synergy | Forms the low-cost lifegain engine that turns the deck’s other cards on.
- 4 Ajani's Pridemate | synergy | Rewards the deck for repeatedly gaining life and becomes a major early payoff.
- 2 Dawn of Hope | draw | Turns the deck’s lifegain focus into a steady source of cards.
- 2 Well of Lost Dreams | draw | Provides additional card flow alongside repeated life gain.
- 2 Inspiring Overseer | draw | Supplies a flying body while helping refill the hand.
- 4 Attended Healer | threat | Provides a lifegain-focused creature payoff that helps build the board.
- 4 Bloodbond Vampire | threat | Serves as a scalable creature threat in a deck built to gain life often.
- 2 Archangel of Thune | threat | Acts as a powerful top-end threat for the lifegain strategy.
- 2 Cliffhaven Vampire | threat | Gives the deck another threatening payoff for its recurring lifegain plan.

</details>

### 9. Standard midrange, FNM

Format: Standard. Theme: midrange. Pool: any_card. Shortlist: 164 names.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $155.52 to buy, $155.52 the whole deck.

**Summary:** This black-green midrange deck develops an Elf-supported creature board, uses removal to keep opposing threats from taking over, and keeps resources flowing with several draw sources. It wins by applying sustained pressure with its larger creatures after the early board is stabilized. The tradeoff is a creature-heavy plan that can be less explosive than a dedicated aggressive strategy and relies on its removal to maintain favorable board states.

- [INFO] `curve_summary`: average mana value 3.22 over 36 nonland cards

<details><summary>The deck list</summary>

- 9 Forest | land | Basic Forest for the green half of the mana base.
- 9 Swamp | land | Basic Swamp for the black half of the mana base.
- 4 Overgrown Tomb | land | Swamp Forest land that supports both deck colors.
- 2 Deathcap Glade | land | Additional land slot for the black-green mana base.
- 2 Llanowar Elves | ramp | Early Elf Druid acceleration for deploying the deck’s larger creatures.
- 2 Bitter Triumph | removal | Dedicated removal for opposing threats.
- 2 Hero's Downfall | removal | Dedicated removal that supports the midrange plan.
- 2 Pick Your Poison | removal | Flexible removal slot for problematic opposing cards.
- 2 Phyrexian Arena | draw | Persistent draw source for longer games.
- 2 Darkstar Augur | draw | Creature-based draw that contributes to the board.
- 2 Midnight Reaper | draw | Creature draw source for a creature-focused strategy.
- 4 Elvish Archdruid | synergy | Elf Druid that supports the deck’s Elf creature package.
- 4 Wood Elves | synergy | Elf Scout that reinforces the deck’s Elf-focused creature plan.
- 4 Goldvein Hydra | threat | Hydra creature that serves as a major board threat.
- 2 Vein Ripper | threat | Vampire Assassin creature that pressures the opponent.
- 2 Vaultborn Tyrant | threat | Dinosaur creature for the deck’s powerful top end.
- 2 Massacre Girl, Known Killer | threat | Legendary Human Assassin that provides a resilient midrange threat.
- 2 Steel Hellkite | threat | Artifact Dragon creature that adds evasive pressure.
- 2 Zodiark, Umbral God | threat | God creature that rounds out the deck’s large threats.

</details>

### 10. Standard aggro, tournament

Format: Standard. Theme: aggro. Pool: any_card. Shortlist: 294 names.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $68.54 to buy, $68.54 the whole deck.

**Summary:** This deck aims to establish an aggressive board early and convert that pressure into a quick win through its threats and red-white synergy. Removal clears opposing obstacles while interaction helps preserve the attack, and draw cards help the deck continue applying pressure. It gives up broader late-game power for a focused, proactive game plan.

- [INFO] `curve_summary`: average mana value 2.83 over 36 nonland cards

<details><summary>The deck list</summary>

- 8 Plains | land | A basic land slot for the red-white mana base.
- 8 Mountain | land | A basic land slot for the red-white mana base.
- 4 Sacred Foundry | land | A land slot for the red-white mana base.
- 4 Inspiring Vantage | land | A land slot for the red-white mana base.
- 4 Fugitive Codebreaker | draw | A draw card that supports keeping the aggressive plan supplied.
- 2 Reckless Lackey | draw | A draw card for maintaining pressure after the opening turns.
- 4 Boros Charm | interaction | An interaction card for protecting the aggressive plan or disrupting opposition.
- 2 Sheltered by Ghosts | interaction | An interaction card that supports the board-focused plan.
- 4 Harsh Annotation | removal | A removal card for clearing obstacles to attacking.
- 4 Emeritus of Truce // Swords to Plowshares | removal | A removal card for answering opposing threats.
- 4 Warleader's Call | synergy | A synergy card that reinforces the red-white aggressive theme.
- 4 Redcap Gutter-Dweller | threat | A threat that contributes to the deck's aggressive pressure.
- 4 Teapot Slinger | threat | A threat for developing the board and attacking.
- 4 Diversion Specialist | threat | A threat that adds to the deck's attack-focused creature base.

</details>

### 11. Commander with a locked card

Format: Commander. Theme: sacrifice. Pool: any_card. Shortlist: 296 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $415.87 to buy, $415.87 the whole deck.

**Summary:** This Karlov sacrifice deck develops a creature-based engine with sacrifice outlets, payoffs, ramp, and steady card flow, then turns disposable bodies into pressure while keeping opposing boards contained. It wins by building a threatening board and leveraging sacrifice-focused payoffs alongside its larger finishers. The tradeoff is that the deck is most effective when its creatures and engine pieces remain available, so it can be slower to recover after repeated disruption.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 3.21 over 63 nonland cards

<details><summary>The deck list</summary>

- 12 Plains | land | Basic white mana for the two-color mana base.
- 12 Swamp | land | Basic black mana for the two-color mana base.
- 1 Bojuka Bog | land | A land slot with utility for the mana base.
- 1 Castle Doom | land | A land slot supporting the mana base.
- 1 Command Tower | land | A flexible land for the commander's colors.
- 1 Evolving Wilds | land | A land slot that helps assemble basic mana.
- 1 Exotic Orchard | land | A flexible land for the mana base.
- 1 High Market | land | A sacrifice-oriented utility land.
- 1 Myriad Landscape | land | A land slot that supports the mana base.
- 1 Phyrexian Tower | land | A sacrifice-oriented utility land.
- 1 Reliquary Tower | land | A utility land in the mana base.
- 1 Starlit Sanctum | land | A land slot supporting the mana base.
- 1 Terramorphic Expanse | land | A land slot that helps assemble basic mana.
- 1 Westvale Abbey // Ormendahl, Profane Prince | land | A sacrifice-themed land slot.
- 1 Sol Ring | ramp | Required card that fills a ramp slot.
- 1 Ashnod's Altar | ramp | Sacrifice-focused ramp for the deck's core plan.
- 1 Culling the Weak | ramp | A sacrifice-compatible burst of ramp.
- 1 Crowded Crypt | ramp | Ramp that fits the deck's creature-focused plan.
- 1 Pawn of Ulamog | ramp | A sacrifice-themed ramp piece.
- 1 Phyrexian Altar | ramp | Sacrifice-focused ramp for the core engine.
- 1 Pitiless Plunderer | ramp | Ramp that rewards the deck's sacrifice plan.
- 1 Priest of Forgotten Gods | ramp | A ramp creature aligned with sacrificing creatures.
- 1 Sifter of Skulls | ramp | A creature-based ramp piece for the sacrifice shell.
- 1 Skullport Merchant | ramp | A creature-based ramp piece that suits the deck's plan.
- 1 Baron Bertram Graywater | draw | Creature-based card draw for the deck.
- 1 Corrupted Conviction | draw | Sacrifice-compatible card draw.
- 1 Disciple of Bolas | draw | A sacrifice-themed creature that supplies card draw.
- 1 Ecstatic Awakener // Awoken Demon | draw | A sacrifice-oriented draw card that also adds a creature body.
- 1 Relic Vial | draw | An artifact draw piece for the deck's creature plan.
- 1 Shadowheart, Dark Justiciar | draw | A creature-based draw option for a sacrifice deck.
- 1 Smothering Abomination | draw | A sacrifice-focused creature that supplies card draw.
- 1 Tevesh Szat, Doom of Fools | draw | A draw-focused permanent that fits a creature-heavy plan.
- 1 Vampiric Rites | draw | A sacrifice-compatible source of card draw.
- 1 Village Rites | draw | Efficient card draw that fits sacrificing creatures.
- 1 Cartel Aristocrat | interaction | A sacrifice-oriented interaction creature.
- 1 Dark Privilege | interaction | Interaction that fits the deck's sacrifice theme.
- 1 Fanatical Devotion | interaction | A sacrifice-compatible protection piece.
- 1 Flare of Fortitude | interaction | A flexible interaction spell for protecting the board.
- 1 Gift of Doom | interaction | A protection-oriented interaction card for a creature deck.
- 1 Nightmare Shepherd | interaction | A creature-based interaction piece for the sacrifice shell.
- 1 Attrition | removal | Sacrifice-themed repeatable removal.
- 1 Ayli, Eternal Pilgrim | removal | A sacrifice-oriented creature that provides removal.
- 1 Bone Shards | removal | Low-cost removal suited to a sacrifice plan.
- 1 Dictate of Erebos | removal | A sacrifice payoff that pressures opposing boards.
- 1 Eaten Alive | removal | Removal that works with sacrificing creatures.
- 1 Grave Pact | removal | A sacrifice payoff that supplies removal pressure.
- 1 Teysa, Orzhov Scion | removal | A creature-based removal piece in the deck's colors.
- 1 Yawgmoth, Thran Physician | removal | A sacrifice-themed removal creature.
- 1 Altar of Dementia | synergy | A sacrifice outlet and synergy piece for the central plan.
- 1 Bartolomé del Presidio | synergy | A creature-based sacrifice synergy piece.
- 1 Bastion of Remembrance | synergy | A sacrifice payoff supporting the deck's main theme.
- 1 Bloodflow Connoisseur | synergy | A creature-based sacrifice outlet.
- 1 Carrion Feeder | synergy | An efficient creature-based sacrifice outlet.
- 1 Chthonian Nightmare | synergy | A synergy card for the sacrifice-focused shell.
- 1 Elas il-Kor, Sadistic Pilgrim | synergy | A sacrifice payoff creature for the deck's core plan.
- 1 Fleshtaker | synergy | A creature-based sacrifice payoff.
- 1 Open the Graves | synergy | A sacrifice-themed permanent that supports creature density.
- 1 Spawning Pit | synergy | An artifact sacrifice outlet for the deck's engine.
- 1 Victimize | synergy | A sacrifice-compatible synergy spell.
- 1 Viscera Seer | synergy | An efficient creature-based sacrifice outlet.
- 1 Woe Strider | synergy | A sacrifice-themed creature and synergy piece.
- 1 Zulaport Cutthroat | synergy | A sacrifice payoff that supports the deck's closing plan.
- 1 Abhorrent Overlord | threat | A large creature threat for closing games.
- 1 Felisa, Fang of Silverquill | threat | A threat that fits the creature-focused sacrifice plan.
- 1 Ghoulcaller Gisa | threat | A sacrifice-themed creature threat.
- 1 Liesa, Forgotten Archangel | threat | A resilient top-end creature threat.
- 1 Liliana, Heretical Healer // Liliana, Defiant Necromancer | threat | A sacrifice-compatible threat that adds a lasting board presence.
- 1 Mondrak, Glory Dominus | threat | A powerful creature threat for the token-oriented elements of the plan.
- 1 Ratadrabik of Urborg | threat | A creature threat supporting the deck's sacrifice theme.
- 1 Razaketh, the Foulblooded | threat | A major top-end threat for closing games.
- 1 Requiem Angel | threat | A creature threat that fits a creature-heavy sacrifice shell.
- 1 Sidisi, Undead Vizier | threat | A sacrifice-themed creature threat.
- 1 Threefold Thunderhulk | threat | A substantial artifact creature threat.
- 1 Vindictive Vampire | threat | A sacrifice-themed threat supporting the deck's finishing plan.
- 1 Elesh Norn // The Argent Etchings | wipe | A board-resetting permanent that can turn the game in your favor.
- 1 The Meathook Massacre | wipe | A board wipe that fits the sacrifice-focused game plan.
- 1 Toxic Deluge | wipe | An efficient reset button against developed boards.

</details>

### 12. Commander on a budget

Format: Commander. Theme: tokens. Pool: any_card. Shortlist: 296 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $43.80 to buy, $43.80 the whole deck.

**Summary:** This white token deck uses Adeline, Resplendent Cathar to establish a wide battlefield, then turns that board into pressure with token-focused synergy cards and a deep creature threat suite. Ramp and draw help it keep developing after its early turns, while targeted removal, protective interaction, and board wipes keep opposing plans in check. It wins by maintaining a larger board and attacking consistently, but it gives up premium mana, expensive token doublers, and the raw card quality of higher-budget builds.

- [INFO] `bracket_prose_rules`: bracket 2 (Core, 8+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 3.22 over 63 nonland cards

<details><summary>The deck list</summary>

- 36 Plains | land | Provides the deck's white mana base without using budget.
- 1 Angelic Sell-Sword | draw | Fills a low-cost draw slot for the token plan.
- 1 Bygone Bishop | draw | Adds an inexpensive draw card that supports the creature-heavy build.
- 1 Dawn of Hope | draw | Provides a budget draw option for the deck.
- 1 Faramir, Field Commander | draw | Adds draw support while fitting the deck's creature plan.
- 1 Idol of Oblivion | draw | Provides a colorless draw slot suited to a token deck.
- 1 Search the Premises | draw | Adds a low-cost draw engine to sustain the deck.
- 1 Staff of the Storyteller | draw | Provides affordable draw support.
- 1 Sunpearl Kirin | draw | Fills an inexpensive draw slot in the creature suite.
- 1 Wedding Announcement // Wedding Festivity | draw | Adds draw support that fits the deck's token theme.
- 1 Wojek Investigator | draw | Rounds out the draw package with a budget creature.
- 1 Ainok Strike Leader | interaction | Provides a creature-based interaction slot.
- 1 Basri Ket | interaction | Adds versatile interaction while staying on the token plan.
- 1 Basri, Tomorrow's Champion | interaction | Provides an inexpensive interaction card for the creature deck.
- 1 Lena, Selfless Champion | interaction | Adds protective interaction for the board-focused strategy.
- 1 Rootborn Defenses | interaction | Provides efficient interaction to protect the token board.
- 1 Spirit Bonds | interaction | Adds low-cost interaction that fits the creature-heavy shell.
- 1 Charisma Bobblehead | ramp | Provides colorless ramp at a modest cost.
- 1 Coin of Mastery | ramp | Fills a budget colorless ramp slot.
- 1 Collector's Vault | ramp | Adds affordable artifact ramp.
- 1 Currency Converter | ramp | Provides a low-cost colorless ramp option.
- 1 Druidic Satchel | ramp | Fills an inexpensive ramp slot.
- 1 Goldvein Pick | ramp | Adds cheap artifact ramp to speed up the deck.
- 1 Idol of False Gods | ramp | Provides a budget ramp card.
- 1 Karn, Living Legacy | ramp | Adds ramp while remaining within the deck's color identity.
- 1 Keeper of the Accord | ramp | Provides creature-based ramp for the white deck.
- 1 Monologue Tax | ramp | Adds efficient white ramp to support the curve.
- 1 Aerial Assault | removal | Provides low-cost targeted removal.
- 1 Banishing Slash | removal | Adds efficient sorcery-speed removal.
- 1 Citizen's Crowbar | removal | Provides a flexible artifact-based removal slot.
- 1 Generous Gift | removal | Adds broad, efficient removal.
- 1 Kellan's Lightblades | removal | Provides inexpensive instant-speed removal.
- 1 Righteous Confluence | removal | Adds flexible removal for troublesome permanents.
- 1 Skyclave Apparition | removal | Provides creature-based removal that advances the board.
- 1 Stroke of Midnight | removal | Rounds out the targeted removal with a cheap instant.
- 1 Aligned Heart | synergy | Adds a low-cost synergy piece for the token strategy.
- 1 Anafenza, Unyielding Lineage | synergy | Provides creature-based token synergy.
- 1 Animation Module | synergy | Adds artifact-based synergy for building a wider board.
- 1 Automated Assembly Line | synergy | Provides a budget token-synergy artifact.
- 1 Cat Collector | synergy | Adds an inexpensive creature that supports the token plan.
- 1 Clarion Spirit | synergy | Provides a low-cost synergy creature for going wide.
- 1 Divine Visitation | synergy | Adds a powerful but budget-friendly token synergy piece.
- 1 Eyes in the Skies | synergy | Contributes directly to the deck's token-focused plan.
- 1 Felidar Retreat | synergy | Adds flexible token synergy to the deck.
- 1 Hero of Precinct One | synergy | Provides a cheap synergy creature for the board-building plan.
- 1 Intangible Virtue | synergy | Supports the deck's wide token battlefield.
- 1 March of the Canonized | synergy | Adds an inexpensive synergy card for the token shell.
- 1 Retrofitter Foundry | synergy | Provides repeatable artifact-based token synergy.
- 1 Rosie Cotton of South Lane | synergy | Adds a low-cost legendary synergy creature.
- 1 Ajani's Chosen | threat | Adds a budget threat that fits the token theme.
- 1 Archangel Elspeth | threat | Provides a resilient high-impact threat.
- 1 Archon of Sun's Grace | threat | Adds an efficient flying threat for the deck.
- 1 Attended Healer | threat | Provides a low-cost creature threat.
- 1 Basri's Lieutenant | threat | Adds a creature threat that fits the white token shell.
- 1 Cemetery Protector | threat | Provides an efficient standalone threat.
- 1 Defiler of Faith | threat | Adds a budget creature threat to pressure opponents.
- 1 Dragonback Lancer | threat | Fills a low-cost threat slot.
- 1 Drogskol Cavalry | threat | Adds a token-oriented creature threat.
- 1 Emeria Angel | threat | Provides an inexpensive flying threat for the deck.
- 1 Gideon, Ally of Zendikar | threat | Adds a versatile white threat.
- 1 God-Eternal Oketra | threat | Provides a durable top-end threat.
- 1 Hour of Reckoning | wipe | Provides a board reset suited to a token-focused deck.
- 1 Martial Coup | wipe | Adds a budget wipe that remains aligned with the deck's plan.
- 1 Phyrexian Rebirth | wipe | Rounds out the board-wipe package with an inexpensive option.

</details>

### 13. owned first, and the commander is not owned

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 303 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $60.77 to buy, $268.98 the whole deck.

**Summary:** Karlov leads a creature-focused lifegain deck that develops its mana, establishes supporting permanents, and turns steady life gain into a growing battlefield presence. It wins by protecting Karlov and its larger threats, using targeted answers and board resets to clear a path for combat pressure. The deck gives up some speed for a resilient, resource-oriented plan with room to rebuild after opposing disruption.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [WARN] `not_owned`: Oathsworn Vampire: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Archangel of Thune: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Astarion, the Decadent: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Karlov of the Ghost Council: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.03 over 63 nonland cards

<details><summary>The deck list</summary>

- 20 Plains | land | Provides dependable white mana for the deck's core plan.
- 10 Swamp | land | Provides dependable black mana for the deck's core plan.
- 1 Command Tower | land | Provides flexible mana for Karlov's colors.
- 1 Marsh Flats | land | Helps smooth access to the deck's basic lands.
- 1 Fabled Passage | land | Helps smooth access to the deck's basic lands.
- 1 Takenuma, Abandoned Mire | land | Adds a black source with utility value.
- 1 War Room | land | Provides a utility land slot.
- 1 Reliquary Tower | land | Provides a colorless utility land slot.
- 1 Sol Ring | ramp | Efficiently accelerates the deck's development.
- 1 Arcane Signet | ramp | Provides reliable color fixing and acceleration.
- 1 Commander's Sphere | ramp | Fixes mana while supporting later resources.
- 1 Fellwar Stone | ramp | Provides low-cost mana acceleration.
- 1 Wayfarer's Bauble | ramp | Helps develop the mana base early.
- 1 Sword of the Animist | ramp | Turns attacks into mana development.
- 1 Relic of Legends | ramp | Adds flexible mana acceleration.
- 1 Thought Vessel | ramp | Provides mana acceleration and utility.
- 1 Lotho, Corrupt Shirriff | ramp | Supports the deck's mana development.
- 1 Deadly Dispute | ramp | Converts a spare permanent into resources and mana.
- 1 Buster Sword | draw | Supplies repeatable card advantage from an equipment slot.
- 1 Call of the Ring | draw | Provides a continuing source of cards.
- 1 Gollum, Riddle Master | draw | Adds a creature-based card advantage option.
- 1 Grave Venerations | draw | Provides card advantage for the midgame.
- 1 Idol of Oblivion | draw | Offers efficient artifact-based card advantage.
- 1 Inspiring Overseer | draw | Provides a creature that replaces itself with a card.
- 1 Lembas | draw | Adds a compact card advantage piece.
- 1 Mask of Memory | draw | Rewards combat with additional cards.
- 1 Night's Whisper | draw | Provides efficient early card selection.
- 1 Skullclamp | draw | Turns expendable creatures into cards.
- 1 Bastion Protector | interaction | Helps keep Karlov protected on the battlefield.
- 1 Boromir, Warden of the Tower | interaction | Provides a disruptive creature-based answer.
- 1 Clever Concealment | interaction | Protects the board from opposing pressure.
- 1 Darksteel Plate | interaction | Protects an important creature over multiple turns.
- 1 Lightning Greaves | interaction | Protects key creatures while supporting immediate attacks.
- 1 Swiftfoot Boots | interaction | Provides repeatable protection for key creatures.
- 1 Banishing Light | removal | Provides a flexible answer to a troublesome permanent.
- 1 Bitter Triumph | removal | Offers efficient spot removal.
- 1 Crib Swap | removal | Answers a problematic creature cleanly.
- 1 Dispatch | removal | Provides a low-cost targeted answer.
- 1 Generous Gift | removal | Answers a broad range of opposing permanents.
- 1 Get Lost | removal | Provides versatile targeted removal.
- 1 Infernal Grasp | removal | Provides straightforward creature removal.
- 1 Swords to Plowshares | removal | Provides efficient creature removal.
- 1 Austere Command | wipe | Provides a flexible reset when the board gets crowded.
- 1 Fumigate | wipe | Resets creature-heavy boards while fitting the deck's theme.
- 1 Vanquish the Horde | wipe | Provides an efficient board reset.
- 1 Aerith Gainsborough | synergy | Supports the deck's lifegain-focused engine.
- 1 Aettir and Priwen | synergy | Adds an equipment-based support piece for the deck's plan.
- 1 Angel of Vitality | synergy | Reinforces the lifegain theme on a flying creature.
- 1 Compassionate Healer | synergy | Provides a creature-based lifegain support piece.
- 1 Crowd of True Believers | synergy | Adds another permanent that supports the lifegain plan.
- 1 Denethor, Ruling Steward | synergy | Contributes to the deck's interconnected creature plan.
- 1 Eastfarthing Farmer | synergy | Provides another thematic support creature.
- 1 Excalibur II | synergy | Adds an equipment support piece for the creature plan.
- 1 Light of Promise | synergy | Turns the deck's central theme into added pressure.
- 1 Momo, Playful Pet | synergy | Provides a low-cost thematic support permanent.
- 1 Night Nurse, Healer of Heroes | synergy | Strengthens the deck's lifegain-focused game plan.
- 1 Oathsworn Vampire | synergy | Adds a recurring creature aligned with the deck's theme.
- 1 Rosie Cotton of South Lane | synergy | Builds pressure alongside the deck's lifegain plan.
- 1 Second Breakfast | synergy | Provides a compact spell that supports the theme.
- 1 Angel of Invention | threat | Provides a substantial evasive threat.
- 1 Bill the Pony | threat | Adds a creature threat that fits the deck's colors.
- 1 Dawnhand Eulogist | threat | Provides another meaningful creature threat.
- 1 Gollum, Silent Slinker // Meager Meal | threat | Offers a threatening creature with added flexibility.
- 1 Lyra Dawnbringer | threat | Adds a powerful flying threat.
- 1 Minwu, White Mage | threat | Provides a legendary creature threat for the deck.
- 1 Rooftop Percher | threat | Adds another creature that can pressure opponents.
- 1 Shattered Angel | threat | Provides an evasive threat aligned with the theme.
- 1 Sneering Shadewriter | threat | Adds a black creature threat to the curve.
- 1 Victory's Herald | threat | Provides a high-impact flying threat.
- 1 Archangel of Thune | threat | Serves as a premier lifegain-oriented threat.
- 1 Astarion, the Decadent | threat | Provides a powerful black-white finisher.

</details>

### 14. the user delegates the commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 244 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $564.85 to buy, $564.85 the whole deck.

**Summary:** This deck builds around Atarka, World Render with a focused Dragon roster, using ramp to bring its larger threats forward while draw keeps the pressure flowing. It wins by establishing a board of Dragons and turning that board into combat pressure, with removal, wipes, and protective interaction to clear or preserve the way. It gives up a broad mixed-creature strategy for a committed Dragon plan that relies on its major threats to carry the game.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 4.21 over 63 nonland cards

<details><summary>The deck list</summary>

- 10 Mountain | land | Forms a core part of the red land base.
- 10 Forest | land | Forms a core part of the green land base.
- 1 Command Tower | land | Helps form the deck's land base.
- 1 Stomping Ground | land | Helps form the deck's land base.
- 1 Rootbound Crag | land | Helps form the deck's land base.
- 1 Cinder Glade | land | Helps form the deck's land base.
- 1 Path of Ancestry | land | Helps form the Dragon-focused land base.
- 1 Temple of the Dragon Queen | land | Helps form the Dragon-focused land base.
- 1 Haven of the Spirit Dragon | land | Helps form the Dragon-focused land base.
- 1 Crucible of the Spirit Dragon | land | Helps form the Dragon-focused land base.
- 1 Three Tree City | land | Helps form the Dragon-focused land base.
- 1 Boseiju, Who Endures | land | Adds a utility land to the mana base.
- 1 Yavimaya, Cradle of Growth | land | Adds a utility land to the mana base.
- 1 Cavern of Souls | land | Helps form the Dragon-focused land base.
- 1 Evolving Wilds | land | Adds flexibility to the land base.
- 1 Fabled Passage | land | Adds flexibility to the land base.
- 1 Myriad Landscape | land | Adds flexibility to the land base.
- 1 Maelstrom of the Spirit Dragon | land | Helps form the Dragon-focused land base.
- 1 Scaled Nurturer | ramp | Provides early ramp for the Dragon plan.
- 1 Dragon's Hoard | ramp | Provides ramp for the Dragon plan.
- 1 Orb of Dragonkind | ramp | Provides ramp for the Dragon plan.
- 1 Carnelian Orb of Dragonkind | ramp | Provides ramp for the Dragon plan.
- 1 Jade Orb of Dragonkind | ramp | Provides ramp for the Dragon plan.
- 1 Ganax, Astral Hunter | ramp | Provides Dragon-themed ramp.
- 1 Ancient Copper Dragon | ramp | Adds a powerful Dragon ramp piece.
- 1 Klauth, Unrivaled Ancient | ramp | Adds a powerful Dragon ramp piece.
- 1 Savage Ventmaw | ramp | Adds a Dragon ramp piece.
- 1 Goldspan Dragon | ramp | Adds a Dragon ramp piece.
- 1 Elemental Bond | draw | Provides card draw alongside large creatures.
- 1 Garruk's Uprising | draw | Provides card draw for the creature plan.
- 1 Guardian Project | draw | Provides sustained card draw.
- 1 Return of the Wildspeaker | draw | Provides card draw for the creature plan.
- 1 Rishkar's Expertise | draw | Provides card draw for the creature plan.
- 1 Dragonborn Champion | draw | Adds Dragon-themed card draw.
- 1 Knollspine Dragon | draw | Adds a Dragon card-draw piece.
- 1 Dragon Mage | draw | Adds a Dragon card-draw piece.
- 1 Sylvan Library | draw | Provides card draw support.
- 1 Vanquisher's Banner | draw | Provides card draw for the Dragon plan.
- 1 Heroic Intervention | interaction | Provides protective interaction.
- 1 Lightning Greaves | interaction | Provides protective interaction for key creatures.
- 1 Swiftfoot Boots | interaction | Provides protective interaction for key creatures.
- 1 Tamiyo's Safekeeping | interaction | Provides protective interaction.
- 1 Snakeskin Veil | interaction | Provides protective interaction.
- 1 Brotherhood Regalia | interaction | Provides equipment-based interaction.
- 1 Draconic Roar | removal | Provides Dragon-themed removal.
- 1 Dragon's Fire | removal | Provides Dragon-themed removal.
- 1 Dragonlord Atarka | removal | Adds a high-impact Dragon removal piece.
- 1 Drakuseth, Maw of Flames | removal | Adds a high-impact Dragon removal piece.
- 1 Scourge of Valkas | removal | Adds Dragon-themed removal.
- 1 Terror of the Peaks | removal | Adds Dragon-themed removal.
- 1 Wrathful Red Dragon | removal | Adds Dragon-themed removal.
- 1 Invasion of Tarkir // Defiant Thundermaw | removal | Provides removal while fitting the Dragon plan.
- 1 Acolyte of Bahamut | synergy | Supports the Dragon-focused strategy.
- 1 Breaching Dragonstorm | synergy | Supports the Dragon-focused strategy.
- 1 Chaos Dragon | synergy | Adds a Dragon synergy piece.
- 1 Dragon Egg | synergy | Adds a Dragon synergy piece.
- 1 Dragon Hatchling | synergy | Adds a Dragon synergy piece.
- 1 Dragon's Approach | synergy | Supports the Dragon-focused strategy.
- 1 Dragonkin Berserker | synergy | Supports the Dragon-focused strategy.
- 1 Dragonlord's Servant | synergy | Supports the Dragon-focused strategy.
- 1 Dragonspeaker Shaman | synergy | Supports the Dragon-focused strategy.
- 1 Dragonstorm | synergy | Supports the Dragon-focused strategy.
- 1 Firespitter Whelp | synergy | Adds a Dragon synergy piece.
- 1 Kargan Dragonrider | synergy | Supports the Dragon-focused strategy.
- 1 Minion of the Mighty | synergy | Supports the Dragon-focused strategy.
- 1 Sarkhan's Triumph | synergy | Supports the Dragon-focused strategy.
- 1 Ancient Bronze Dragon | threat | Adds a powerful Dragon threat.
- 1 Backdraft Hellkite | threat | Adds a Dragon threat.
- 1 Blast-Furnace Hellkite | threat | Adds a Dragon threat.
- 1 Hellkite Charger | threat | Adds a Dragon threat.
- 1 Lathliss, Dragon Queen | threat | Adds a Dragon threat.
- 1 Scourge of the Throne | threat | Adds a Dragon threat.
- 1 Terror of Mount Velus | threat | Adds a Dragon threat.
- 1 Thrakkus the Butcher | threat | Adds a Dragon threat.
- 1 Twinflame Tyrant | threat | Adds a Dragon threat.
- 1 Utvara Hellkite | threat | Adds a powerful Dragon threat.
- 1 Thunderbreak Regent | threat | Adds a Dragon threat.
- 1 Dragonhawk, Fate's Tempest | threat | Adds a Dragon threat.
- 1 Balefire Dragon | wipe | Provides a Dragon-themed board wipe.
- 1 Draconic Intervention | wipe | Provides a board wipe.
- 1 Harbinger of the Hunt | wipe | Provides a Dragon-themed board wipe.

</details>

### 15. delegated commander, owned first

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 303 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for copy_limit. Block findings: 0.

Cost: $167.98 to buy, $294.88 the whole deck.

**Summary:** A black-white lifegain deck that steadily develops mana and card advantage, then turns recurring life gain into growing creatures, Angel armies, and draining pressure. It combines resilient threats with flexible answers and broad reset tools, allowing it to stabilize comfortably before closing through combat or life-total payoffs.

- [INFO] `bracket_prose_rules`: bracket 2 (Core, 8+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [WARN] `not_owned`: Dawn of Hope: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Sigarda's Splendor: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Heliod, Sun-Crowned: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vito, Thorn of the Dusk Rose: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Sanguine Bond: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Cleric Class: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ajani's Pridemate: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bloodthirsty Aerialist: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Voice of the Blessed: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Resplendent Angel: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Righteous Valkyrie: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Essence Channeler: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Marauding Blight-Priest: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Indulging Patrician: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Starscape Cleric: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Archangel of Thune: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Celestine, the Living Saint: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Cliffhaven Vampire: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Defiant Bloodlord: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Divinity of Pride: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Elenda, Saint of Dusk: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Enduring Tenacity: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Nykthos Paragon: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Rodolf Duskbringer: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Tivash, Gloom Summoner: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Valkyrie Harbinger: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Astarion, the Decadent: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.14 over 63 nonland cards

<details><summary>The deck list</summary>

- 12 Plains | land | Core white mana base for lifegain creatures, protection, and sweepers.
- 10 Swamp | land | Core black mana base for drain effects and removal.
- 1 Command Tower | land | Reliable access to both commander colors.
- 1 Exotic Orchard | land | Flexible dual-color fixing in multiplayer games.
- 1 Marsh Flats | land | Finds either basic color efficiently.
- 1 Fabled Passage | land | Fixes mana while supporting the basic-land base.
- 1 Evolving Wilds | land | Budget color fixing for both basic land types.
- 1 Terramorphic Expanse | land | Additional basic-land fixing.
- 1 Ash Barrens | land | Basic-land cycling smooths early colors.
- 1 Path of Ancestry | land | Produces either color and supports the Vampire commander theme.
- 1 Secluded Courtyard | land | Names Vampire to cast the commander and key creatures.
- 1 Thriving Moor | land | Black source that can also provide white.
- 1 Takenuma, Abandoned Mire | land | Black source with late-game creature recursion.
- 1 Castle Locthwain | land | Land-based card advantage when resources run low.
- 1 War Room | land | Repeatable card draw from a land slot.
- 1 Scavenger Grounds | land | Graveyard hate without using a spell slot.
- 1 Sol Ring | ramp | Efficient universal acceleration.
- 1 Arcane Signet | ramp | Reliable two-color mana acceleration.
- 1 Fellwar Stone | ramp | Early fixing and acceleration in multiplayer.
- 1 Commander's Sphere | ramp | Color fixing that can become a card later.
- 1 Wayfarer's Bauble | ramp | Finds a basic land and improves mana development.
- 1 Sword of the Animist | ramp | Turns attacking creatures into repeatable land ramp.
- 1 Thought Vessel | ramp | Mana acceleration with a useful hand-size benefit.
- 1 Relic of Legends | ramp | Converts legendary creatures into mana sources.
- 1 Lotho, Corrupt Shirriff | ramp | Rewards opposing noncreature spells with Treasure.
- 1 Oath of the Grey Host | ramp | Provides thematic mana development and board presence.
- 1 Dawn of Hope | draw | Converts life gained into a steady flow of cards.
- 1 Night's Whisper | draw | Efficient early card selection.
- 1 Painful Truths | draw | Low-cost card refuel supported by a high life total.
- 1 Skullclamp | draw | Turns disposable creatures into substantial card advantage.
- 1 Tome of Legends | draw | Builds counters through commander activity for repeatable draw.
- 1 Wall of Omens | draw | Early defense that replaces itself.
- 1 Mask of Memory | draw | Combat-based card filtering for evasive attackers.
- 1 Lembas | draw | Immediate card selection with later card draw.
- 1 Idol of Oblivion | draw | Reliable draw alongside creature-making effects.
- 1 Sigarda's Splendor | draw | Lifegain payoff that keeps cards flowing.
- 1 Swords to Plowshares | removal | Premium answer to troublesome creatures.
- 1 Generous Gift | removal | Versatile answer to any problematic permanent.
- 1 Get Lost | removal | Efficient answer to creatures, enchantments, and planeswalkers.
- 1 Stroke of Midnight | removal | Flexible instant-speed permanent removal.
- 1 Infernal Grasp | removal | Unconditional creature removal with a manageable life cost.
- 1 Bitter Triumph | removal | Efficient removal that can leverage surplus life.
- 1 Crib Swap | removal | Exiles creatures and bypasses death triggers.
- 1 Banishing Light | removal | Broad permanent-based answer for difficult threats.
- 1 Lightning Greaves | interaction | Protects key creatures while granting immediate utility.
- 1 Swiftfoot Boots | interaction | Provides repeatable protection and haste.
- 1 Darksteel Plate | interaction | Keeps an important creature alive through most damage and destruction.
- 1 Clever Concealment | interaction | Protects the entire board from sweeping removal.
- 1 Unbreakable Formation | interaction | Safeguards the creature board and can strengthen an attack.
- 1 Reprieve | interaction | Temporarily stops a pivotal spell while replacing itself.
- 1 Austere Command | wipe | Flexible sweeper that can preserve the most useful board pieces.
- 1 Fumigate | wipe | Creature reset that restores life afterward.
- 1 Vanquish the Horde | wipe | Often becomes an efficient creature-board reset.
- 1 Heliod, Sun-Crowned | synergy | Places counters from lifegain and enables powerful life-based combinations.
- 1 Vito, Thorn of the Dusk Rose | synergy | Converts each meaningful life gain into opponent life loss.
- 1 Sanguine Bond | synergy | Provides a durable lifegain-to-drain payoff.
- 1 Cleric Class | synergy | Enhances life gain and rewards building a large life total.
- 1 Ajani's Pridemate | synergy | Grows rapidly from repeated life-gain triggers.
- 1 Bloodthirsty Aerialist | synergy | Flying lifegain payoff that accumulates counters quickly.
- 1 Voice of the Blessed | synergy | Scales from life gain into a resilient evasive attacker.
- 1 Resplendent Angel | synergy | Rewards large life-gain bursts with Angel tokens.
- 1 Righteous Valkyrie | synergy | Provides life gain and a major payoff for reaching a high life total.
- 1 Angel of Vitality | synergy | Improves lifegain bursts and becomes a larger flier at high life.
- 1 Essence Channeler | synergy | Creates repeatable lifegain triggers from creature spells.
- 1 Marauding Blight-Priest | synergy | Turns life gain into incremental pressure on each opponent.
- 1 Indulging Patrician | synergy | Rewards substantial life gain with table-wide life loss.
- 1 Starscape Cleric | synergy | Pairs creature development with consistent drain pressure.
- 1 Archangel of Thune | threat | Makes every life-gain event a team-wide permanent upgrade.
- 1 Celestine, the Living Saint | threat | A lifelinking threat that recovers valuable permanents.
- 1 Cliffhaven Vampire | threat | Turns lifegain into repeated multiplayer damage.
- 1 Defiant Bloodlord | threat | Large evasive finisher that weaponizes life gain.
- 1 Divinity of Pride | threat | Efficient high-power lifelinking attacker for a life-total deck.
- 1 Elenda, Saint of Dusk | threat | Growing lifelink attacker that leaves behind a threatening board.
- 1 Enduring Tenacity | threat | Drains opponents while offering a resilient permanent body.
- 1 Nykthos Paragon | threat | Transforms a major lifegain burst into a decisive combat swing.
- 1 Rodolf Duskbringer | threat | Lifelink threat that can return an Angel or Vampire after combat.
- 1 Tivash, Gloom Summoner | threat | Creates a large lifelinking Demon from the deck's life total.
- 1 Valkyrie Harbinger | threat | Produces Angels after recurring large life-gain turns.
- 1 Lyra Dawnbringer | threat | Powerful lifelinking flier that strengthens other Angels.

</details>

### 16. a tight budget, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 353 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $6.83 to buy, $31.24 the whole deck.

**Summary:** Denethor leads a black-white aristocrats deck that develops a creature board, uses sacrifice-oriented synergy to turn that board into pressure, and backs it with ample card flow, mana development, and answers. It wins by layering its sacrifice pieces with its threats until opponents cannot comfortably stabilize, while wipes can reset crowded boards. The deck gives up premium standalone finishers and an elaborate land base in favor of library-owned cards, inexpensive additions, and a focused creature-driven plan.

- JUDGE [true]: "Denethor leads a black-white aristocrats deck". This asserts Denethor can serve as a commander, which is a rules claim about deck legality. Denethor, Ruler of Gondor is a legendary creature and thus eligible to be a commander, so the claim is true.
- [INFO] `bracket_prose_rules`: bracket 2 (Core, 8+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [WARN] `not_owned`: Corrupted Conviction: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vampiric Rites: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Village Rites: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Acolyte of Aclazotz: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bastion of Remembrance: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Blood Bairn: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Catapult Fodder // Catapult Captain: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Fleshtaker: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Legion Vanguard: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Pious Evangel // Wayward Disciple: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Pitiless Pontiff: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Viscera Seer: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Zulaport Cutthroat: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Apocalypse Demon: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Blood Host: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Demonlord of Ashmouth: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Falkenrath Noble: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Morkrut Necropod: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Novice Dissector: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Old Flitterfang: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Sivriss, Nightmare Speaker: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vindictive Vampire: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vito's Inquisitor: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vampire Warlord: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 2.92 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Grave Venerations | draw | An owned draw card for maintaining resources through the game.
- 1 Idol of Oblivion | draw | An owned draw option that fits the deck's sacrifice-focused plan.
- 1 Mask of Memory | draw | An owned draw card that helps keep cards flowing.
- 1 Skullclamp | draw | An owned draw card for a creature-focused deck.
- 1 Tome of Legends | draw | An owned draw card that provides steady support.
- 1 Wall of Omens | draw | An owned draw creature that also helps the early board.
- 1 Nasty End | draw | An owned draw spell for replenishing the hand.
- 1 Corrupted Conviction | draw | A low-cost draw spell that suits an aristocrats shell.
- 1 Vampiric Rites | draw | A low-cost draw card for the sacrifice-oriented plan.
- 1 Village Rites | draw | A low-cost draw spell that complements creature sacrifices.
- 1 Bastion Protector | interaction | An owned interaction card for protecting the deck's key pieces.
- 1 Frontline Medic | interaction | An owned interaction creature that supports the board.
- 1 Swiftfoot Boots | interaction | An owned interaction card for safeguarding an important creature.
- 1 Together Forever | interaction | An owned interaction card that supports a creature-based strategy.
- 1 Ultimate Magic: Holy | interaction | An owned interaction spell for answering opposing plays.
- 1 Unbreakable Formation | interaction | An owned interaction spell that helps preserve the board.
- 1 Arcane Signet | ramp | An owned ramp staple for early mana development.
- 1 Astral Cornucopia | ramp | An owned ramp card that helps build mana.
- 1 Chromatic Lantern | ramp | An owned ramp card for dependable mana development.
- 1 Commander's Sphere | ramp | An owned ramp card that supports the mana base.
- 1 Deadly Dispute | ramp | An owned ramp spell that fits a sacrifice-oriented deck.
- 1 Fellwar Stone | ramp | An owned ramp artifact for early development.
- 1 Inherited Envelope | ramp | An owned ramp card that keeps the budget focused on upgrades.
- 1 Relic of Legends | ramp | An owned ramp artifact for improving mana access.
- 1 Sol Ring | ramp | An owned ramp staple for accelerating the deck.
- 1 Thought Vessel | ramp | An owned ramp artifact that supports sustained development.
- 1 Angel of Serenity | removal | An owned removal creature that also contributes to the board.
- 1 Banishing Light | removal | An owned removal option for troublesome permanents.
- 1 Bitter Triumph | removal | An owned removal spell for efficient answers.
- 1 Claim the Precious | removal | An owned removal spell for handling opposing threats.
- 1 Contagion Clasp | removal | An owned removal artifact that broadens the deck's answers.
- 1 Crib Swap | removal | An owned removal spell that answers creatures.
- 1 Destroy Evil | removal | An owned removal spell with useful flexibility.
- 1 Fatal Push | removal | An owned removal spell for efficient interaction.
- 1 Austere Command | wipe | An owned wipe that provides a broad reset option.
- 1 Black Sun's Zenith | wipe | An owned wipe for clearing a developed board.
- 1 Dusk // Dawn | wipe | An owned wipe that fits the creature-focused deck.
- 1 Arcade Cabinet | synergy | An owned synergy piece for the aristocrats core.
- 1 Gollum the Abandoned | synergy | An owned synergy creature that supports the central plan.
- 1 Gollum, Patient Plotter | synergy | An owned synergy creature for the sacrifice-focused strategy.
- 1 Gríma Wormtongue | synergy | An owned synergy card that preserves the library-first approach.
- 1 Acolyte of Aclazotz | synergy | A low-cost synergy creature for the aristocrats plan.
- 1 Bastion of Remembrance | synergy | A low-cost synergy enchantment for the deck's sacrifice theme.
- 1 Blood Bairn | synergy | A low-cost synergy creature for a sacrifice-oriented board.
- 1 Catapult Fodder // Catapult Captain | synergy | A low-cost synergy creature that contributes to the main plan.
- 1 Fleshtaker | synergy | A low-cost synergy creature for the aristocrats core.
- 1 Legion Vanguard | synergy | A low-cost synergy creature that fits the deck's focus.
- 1 Pious Evangel // Wayward Disciple | synergy | A low-cost synergy creature for the sacrifice strategy.
- 1 Pitiless Pontiff | synergy | A low-cost synergy creature that supports the central plan.
- 1 Viscera Seer | synergy | A low-cost synergy creature for the aristocrats shell.
- 1 Zulaport Cutthroat | synergy | A synergy creature that gives the deck a focused aristocrats payoff.
- 1 Bill the Pony | threat | An owned threat that keeps the build rooted in the library.
- 1 Apocalypse Demon | threat | A low-cost threat for applying pressure.
- 1 Blood Host | threat | A low-cost threat that fits the creature-heavy plan.
- 1 Demonlord of Ashmouth | threat | A low-cost threat for the deck's battlefield plan.
- 1 Falkenrath Noble | threat | A low-cost threat that complements the aristocrats focus.
- 1 Morkrut Necropod | threat | A very inexpensive threat for filling out the creature suite.
- 1 Novice Dissector | threat | A low-cost threat for maintaining board presence.
- 1 Old Flitterfang | threat | A low-cost threat that supports the deck's overall pressure plan.
- 1 Sivriss, Nightmare Speaker | threat | A low-cost threat for the creature-based strategy.
- 1 Vindictive Vampire | threat | A low-cost threat that fits the aristocrats shell.
- 1 Vito's Inquisitor | threat | A low-cost threat for adding battlefield pressure.
- 1 Vampire Warlord | threat | A very inexpensive threat for the deck's creature suite.
- 1 Ash Barrens | land | An owned land that supports the mana base.
- 1 Command Tower | land | An owned land for consistent access to the deck's colors.
- 1 Evolving Wilds | land | An owned land that supports the basic-land mana base.
- 1 Exotic Orchard | land | An owned land for flexible color access.
- 1 Path of Ancestry | land | An owned land that keeps the mana base affordable.
- 1 Scavenger Grounds | land | An owned utility land for the mana base.
- 1 Secluded Courtyard | land | An owned land that supports the creature suite.
- 1 Terramorphic Expanse | land | An owned land that complements the basic-land base.
- 12 Plains | land | Basic lands that provide white mana without using budget.
- 16 Swamp | land | Basic lands that provide black mana for the deck's core plan.

</details>

### 17. upgrade a precon, owned first

Format: Commander. Theme: goblins. Pool: owned_first. Shortlist: 318 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $77.12 to buy, $271.45 the whole deck.

**Summary:** This Goblin Storm deck develops through Goblin synergy, mana acceleration, and a steady stream of draw spells before turning its board into a decisive offensive push. It wins by bringing together its Goblin threats, spell-focused payoffs, and dedicated finishers, while retaining removal and sweepers to keep opposing boards manageable. Its tradeoff is a focused Goblin Storm plan that gives up broader flexibility for stronger themed turns.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 2.57 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Ancestors' Aid | synergy | Retained precon synergy piece for the Goblin Storm plan.
- 1 Arena of Glory | land | Retained precon land.
- 1 Battle Hymn | ramp | Retained precon ramp for explosive turns.
- 1 Blasphemous Act | wipe | Retained precon board-clearing option.
- 1 Boggart Shenanigans | removal | Retained precon removal-oriented Goblin payoff.
- 1 Brightstone Ritual | ramp | Retained precon ramp for Goblin-heavy turns.
- 1 Broadside Bombardiers | removal | Retained precon Goblin removal piece.
- 1 Castle Embereth | land | Retained precon land.
- 1 Chaos Warp | removal | Retained precon flexible removal.
- 1 Conspicuous Snoop | synergy | Retained precon Goblin synergy piece.
- 1 Crimson Wisps | draw | Retained precon card-draw spell.
- 1 Den of the Bugbear | land | Retained precon land.
- 1 Dragon Fodder | synergy | Retained precon Goblin Storm synergy card.
- 1 Dwarven Mine | land | Retained precon land.
- 1 Empty the Warrens | synergy | Retained precon Goblin Storm synergy card.
- 1 Expedite | draw | Retained precon card-draw spell.
- 1 Faithless Looting | draw | Retained precon card-draw spell.
- 1 Fists of Flame | draw | Retained precon card-draw spell.
- 1 Forgotten Cave | land | Retained precon land.
- 1 Fountainport | land | Retained precon land.
- 1 Gempalm Incinerator | removal | Retained precon Goblin removal piece.
- 1 Glimpse the Impossible | ramp | Retained precon ramp spell.
- 1 Goblin Bombardment | removal | Retained precon Goblin removal payoff.
- 1 Goblin Burrows | land | Retained precon land.
- 1 Goblin Chieftain | synergy | Retained precon Goblin synergy piece.
- 1 Goblin Dark-Dwellers | threat | Retained precon Goblin threat.
- 1 Goblin Lackey | synergy | Retained precon Goblin synergy piece.
- 1 Goblin Matron | synergy | Retained precon Goblin synergy piece.
- 1 Goblin Negotiation | removal | Retained precon removal spell.
- 1 Goblin Trashmaster | removal | Retained precon Goblin removal piece.
- 1 Goblin Warchief | synergy | Retained precon Goblin synergy piece.
- 1 Grapeshot | removal | Retained precon removal spell for the storm theme.
- 1 Great Train Heist | ramp | Retained precon ramp spell.
- 1 Haze of Rage | synergy | Retained precon Goblin Storm synergy card.
- 1 Hidden Volcano | land | Retained precon land.
- 1 Howlsquad Heavy | ramp | Retained precon ramp creature.
- 1 Idol of Oblivion | draw | Retained precon draw source.
- 1 Impact Tremors | wincon | Retained precon Goblin Storm payoff.
- 1 Kher Keep | land | Retained precon land.
- 1 Krenko's Command | synergy | Retained precon Goblin Storm synergy card.
- 1 Mana Geyser | ramp | Retained precon ramp spell.
- 1 Mogg War Marshal | synergy | Retained precon Goblin synergy piece.
- 1 Pashalik Mons | removal | Retained precon Goblin removal payoff.
- 1 Past in Flames | synergy | Retained precon spell-focused synergy card.
- 1 Quest for the Goblin Lord | synergy | Retained precon Goblin synergy piece.
- 1 Redcap Gutter-Dweller | threat | Retained precon Goblin threat.
- 1 Reliquary Tower | land | Retained precon land.
- 1 Roaming Throne | threat | Retained precon threat.
- 1 Ruby Medallion | ramp | Retained precon ramp artifact.
- 1 Rundvelt Hordemaster | synergy | Retained precon Goblin synergy piece.
- 1 Sazacap's Brew | draw | Retained precon card-draw spell.
- 1 Searslicer Goblin | synergy | Retained precon Goblin synergy piece.
- 1 Seething Song | ramp | Retained precon ramp spell.
- 1 Shinka, the Bloodsoaked Keep | land | Retained precon land.
- 1 Siege-Gang Commander | removal | Retained precon Goblin removal piece.
- 1 Siege-Gang Lieutenant | removal | Retained precon Goblin removal piece.
- 1 Skirk Prospector | ramp | Retained precon Goblin ramp piece.
- 1 Skullclamp | draw | Retained precon draw artifact.
- 1 Smoldering Crater | land | Retained precon land.
- 1 Sol Ring | ramp | Retained precon ramp artifact.
- 1 Spinerock Knoll | land | Retained precon land.
- 1 Storm-Kiln Artist | ramp | Retained precon ramp creature.
- 1 Swiftfoot Boots | interaction | Retained precon interaction piece.
- 1 Throne of Eldraine | ramp | Retained precon ramp artifact.
- 1 Vandalblast | wipe | Retained precon board-clearing option.
- 1 War Room | land | Retained precon land.
- 1 Witch's Mark | draw | Retained precon card-draw spell.
- 1 Ancestral Anger | draw | Adds efficient card draw for the spell-focused plan.
- 1 Arcane Signet | ramp | Adds dependable ramp.
- 1 Cunning Maneuver | draw | Adds card draw to sustain spell-heavy turns.
- 1 Fable of the Mirror-Breaker // Reflection of Kiki-Jiki | ramp | Adds a versatile ramp piece.
- 1 Impulsive Pilferer | ramp | Adds Goblin-themed ramp.
- 1 Krenko, Mob Boss | threat | Adds a Goblin threat.
- 1 Lightning Greaves | interaction | Adds an interaction piece for the commander-focused plan.
- 1 Mask of Memory | draw | Adds a repeatable draw option.
- 22 Mountain | land | Provides the deck's core land base.
- 1 Rush the Room | synergy | Adds another spell-focused synergy card.
- 1 Vanquisher's Banner | draw | Adds a Goblin-focused draw source.

</details>

### 18. upgrade a precon, any card

Format: Commander. Theme: turtles. Pool: any_card. Shortlist: 303 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $163.04 to buy, $163.04 the whole deck.

**Summary:** This Turtle Power deck keeps the precon’s character-driven Turtle theme at center stage, developing its mana before deploying Turtle cards and other themed permanents. It wins by building pressure with its threat cards while using removal and board wipes to clear a path. The expanded land base and added draw and ramp make the deck steadier, while the deck still gives priority to its themed cards rather than pursuing a tightly streamlined finish.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [WARN] `land_count`: 39 lands: the guide range for this format is 34 to 38
- [INFO] `curve_summary`: average mana value 3.30 over 60 nonland cards
- [INFO] `basics_added`: the list was 3 cards short, so the builder added 3 basic lands

<details><summary>The deck list</summary>

- 1 Acidic Slime | synergy | A retained precon creature that supports the existing Turtle Power shell.
- 1 April O'Neil, Live on the Scene | draw | A retained precon draw card.
- 1 Arcane Signet | ramp | A retained precon ramp artifact.
- 1 Ash Barrens | land | A retained precon land.
- 1 Assassin's Trophy | removal | A retained precon removal spell.
- 1 Baxter, Fly in the Ointment | synergy | A retained precon character card for the deck's theme.
- 1 Bebop, Skull & Crossbones | synergy | A retained precon character card for the deck's theme.
- 1 Big Apple, 3 a.m. | land | A retained precon land.
- 1 Big Mother Mouser | synergy | A retained precon artifact creature for the deck's theme.
- 1 Biogenic Ooze | threat | A retained precon creature that preserves the original deck's shape.
- 1 Blasphemous Act | wipe | A retained precon board wipe.
- 1 Casey Jones, Back Alley Brute | threat | A retained precon character card for the deck's theme.
- 1 Chromatic Lantern | ramp | A retained precon ramp artifact.
- 1 Cinder Glade | land | A retained precon land.
- 1 City of Brass | land | A retained precon land.
- 1 Command Tower | land | A retained precon land.
- 1 Continue? | other | A retained precon instant that preserves the original shell.
- 1 Corpsejack Menace | synergy | A retained precon creature that preserves the original deck's shape.
- 1 Cultivate | ramp | A retained precon ramp spell.
- 1 Dimension X Pizzasaur | synergy | A retained precon artifact creature for the deck's theme.
- 1 Donatello, the Brains | synergy | A retained precon Turtle character card.
- 1 Double Jump // Flying Kick | other | A retained precon instant that preserves the original shell.
- 1 Electric Seaweed | synergy | A retained precon creature that preserves the original deck's shape.
- 1 Endless Foot Assault | synergy | A retained precon enchantment that supports the deck's theme.
- 1 Escape Tunnel | land | A retained precon land.
- 1 Everything Pizza | synergy | A retained precon artifact for the deck's theme.
- 1 Fast Forward | other | A retained precon sorcery that preserves the original shell.
- 1 Game Over | other | A retained precon sorcery that preserves the original shell.
- 1 Grand Coliseum | land | A retained precon land.
- 1 Harmonize | draw | A retained precon card that supports the deck's resources.
- 1 Here Comes a New Hero! | synergy | A retained precon sorcery for the deck's theme.
- 1 Hidden Hideout | land | A retained precon land.
- 1 High Score | synergy | A retained precon enchantment for the deck's theme.
- 1 Hinterland Harbor | land | A retained precon land.
- 1 Irma, Part-Time Mutant | synergy | A retained precon character card for the deck's theme.
- 1 Krang, the All-Powerful | threat | A retained precon artifact creature for the deck's theme.
- 1 Leatherhead, Iron Gator | threat | A retained precon character card for the deck's theme.
- 1 Leonardo, the Balance | threat | A retained precon Turtle character card.
- 1 Lessons from Life | other | A retained precon sorcery that preserves the original shell.
- 1 Level Up | synergy | A retained precon enchantment that supports the deck's theme.
- 1 Lita, Little Orphan Amphibian | synergy | A retained precon Turtle character card.
- 1 Michelangelo, the Heart | synergy | A retained precon Turtle character card.
- 1 Mona Lisa, Science Geek | synergy | A retained precon character card for the deck's theme.
- 1 Ninja Pizza | synergy | A retained precon enchantment for the deck's theme.
- 1 Path of Ancestry | land | A retained precon land.
- 1 Rain-Slicked Copse | land | A retained precon land.
- 1 Raphael, the Muscle | threat | A retained precon Turtle character card.
- 1 Rat King, Pale Piper | synergy | A retained precon character card for the deck's theme.
- 1 Ray Fillet, Wave Warrior | synergy | A retained precon character card for the deck's theme.
- 1 Rocksteady, Mutant Marauder | threat | A retained precon character card for the deck's theme.
- 1 Rootbound Crag | land | A retained precon land.
- 1 Shellshock | other | A retained precon instant that preserves the original shell.
- 1 Shredder, Shadow Master | threat | A retained precon character card for the deck's theme.
- 1 Smoldering Marsh | land | A retained precon land.
- 1 Sodden Verdure | land | A retained precon land.
- 1 Sol Ring | ramp | A retained precon ramp artifact.
- 1 Special Move | other | A retained precon instant that preserves the original shell.
- 1 Spire Garden | land | A retained precon land.
- 1 Splinter & Leo, Father & Son | synergy | A retained precon Turtle character card.
- 1 Splinter, the Mentor | synergy | A retained precon character card for the deck's theme.
- 1 Steelbane Hydra | removal | A retained precon removal creature.
- 1 Super Combo | synergy | A retained precon sorcery for the deck's theme.
- 1 Swift Demise | other | A retained precon instant that preserves the original shell.
- 1 Tempestra, Dame of Games | synergy | A retained precon character card for the deck's theme.
- 1 Together Forever | synergy | A retained precon enchantment for the deck's theme.
- 1 Tokka & Rahzar, Unsupervised | ramp | A retained precon Turtle character card that provides ramp.
- 1 Turtle Lair | land | A retained precon land.
- 1 Undergrowth Stadium | land | A retained precon land.
- 1 Vanquish the Horde | wipe | A retained precon board wipe.
- 1 Vernal Fen | land | A retained precon land.
- 1 Vibrant Cityscape | land | A retained precon land.
- 1 Voracious Hydra | threat | A retained precon creature that preserves the original deck's shape.
- 1 Wave Goodbye | other | A retained precon sorcery that preserves the original shell.
- 1 Breeding Pool | land | A land addition supporting the deck's colors.
- 1 Exotic Orchard | land | A land addition supporting the deck's colors.
- 1 Fabled Passage | land | A land addition supporting the deck's colors.
- 1 Hallowed Fountain | land | A land addition supporting the deck's colors.
- 1 Overgrown Tomb | land | A land addition supporting the deck's colors.
- 1 Steam Vents | land | A land addition supporting the deck's colors.
- 1 Beast Within | removal | A flexible removal addition.
- 1 Farseek | ramp | A ramp addition.
- 1 Nature's Lore | ramp | A ramp addition.
- 1 Rhystic Study | draw | A draw addition.
- 1 Swords to Plowshares | removal | An efficient removal addition.
- 1 Turtle Power! | synergy | A Turtle-focused synergy addition.
- 4 Forest | land | Basic lands retained to support the deck's colors.
- 4 Island | land | Basic lands retained to support the deck's colors.
- 3 Mountain | land | Basic lands retained to support the deck's colors.
- 1 Plains | land | A basic land retained to support the deck's colors.
- 2 Swamp | land | Basic lands retained to support the deck's colors.

</details>

