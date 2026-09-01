# PR-8 deck gate

Run date: 2026-08-31. Card snapshot: 2026-08-31.

Verdict: PASS. 6 of 6 decks passed every block check, 0 invented names reached the user, and 0 summaries stated a false rule of the game. All three bars are zero tolerance.

## Summary

| Measure | Value |
|---|---|
| Prompts | 6 |
| Decks returned | 6 |
| Decks with no block finding | 6 |
| Invented names that reached the user | 0 |
| Decks that needed the repair turn | 0 |
| Summaries judged (F-26) | 6 |
| Summaries that state a rule of the game | 1 |
| Summaries that state a FALSE rule | 0 |
| Judge errors | 0 |
| Errors | 0 |
| Prompt version | 10 |
| Calls | 12 |
| Cost | $0.3539 |
| Time | 306 seconds |

## Findings by code

A block stops the deck. A warning and a note are reports.

| Code | Count |
|---|---|
| `curve_summary` | 6 |
| `bracket_prose_rules` | 5 |
| `outside_requested_set` | 2 |
| `basics_added` | 1 |

By severity: BLOCK 0. WARN 2. INFO 12. 

## The set filter (PR-17B)

A deck asked for a set holds cards of that set family alone. Three things earn an exception, and each one is marked: a card the reader named (D-381), a mana card the reader allowed from outside (D-382), and a basic land, which no set limit filters (D-378).

| # | Prompt | Sets | In set | Fill | Deck cards outside | Marked |
|---|---|---|---|---|---|---|
| 19 | the Hobbit family, two colours | `hob,hoc` | 170 | 0 | 0 | 0 |
| 20 | the Hobbit family, a delegated commander | `hob,hoc` | 75 | 0 | 0 | 0 |
| 21 | the Hobbit family, mana from outside | `hob,hoc` | 127 | 26 | 8 | 8 |
| 22 | a set family and a card from outside it | `hob,hoc` | 170 | 0 | 1 | 1 |
| 23 | two set families at once | `blb,blc,hob,hoc,pblb` | 201 | 0 | 0 | 0 |
| 24 | a 60-card deck from one set | `blb,blc,pblb` | 95 | 0 | 0 | 0 |

## Decks

### 19. the Hobbit family, two colours

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 174 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $503.76 to buy, $503.76 the whole deck.

**Summary:** This is an Elf-centered Thranduil deck that develops its mana, builds a broad creature board, and uses its Elf contingent as the core of its table presence. It aims to pressure opponents through that creature force while maintaining card access, targeted answers, and several ways to reset crowded boards. The deck gives up some card slots to its thematic creature base, so its individual threats are supported most effectively when the board develops together.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 3.41 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Dragon-Cursed Halls | land | A land slot that supports the deck's mana base.
- 1 Elven Passage | land | A land slot that supports the deck's mana base.
- 1 Elvenking's Halls | land | A land slot that supports the deck's mana base.
- 1 Hobbit Hole | land | A land slot that supports the deck's mana base.
- 1 Minas Morgul, Dark Fortress | land | A land slot that supports the deck's mana base.
- 1 Mirkwood | land | A land slot that supports the deck's mana base.
- 1 Rivendell | land | A land slot that supports the deck's mana base.
- 1 Rogue's Passage | land | A land slot that supports the deck's mana base.
- 1 The Black Gate | land | A land slot that supports the deck's mana base.
- 1 The Shire | land | A land slot that supports the deck's mana base.
- 1 Treasure Vault | land | A land slot that supports the deck's mana base.
- 13 Forest | land | Basic lands provide a stable foundation for the mana base.
- 7 Island | land | Basic lands provide a stable foundation for the mana base.
- 5 Swamp | land | Basic lands provide a stable foundation for the mana base.
- 1 Arcane Signet | ramp | An artifact ramp piece for early development.
- 1 Elven Chorus | ramp | An Elf-focused ramp card that supports the creature plan.
- 1 Elvish Archdruid | ramp | An Elf Druid that advances the deck's mana development.
- 1 Elvish Mystic | ramp | A low-cost Elf Druid for early mana development.
- 1 Mox Amber | ramp | A legendary artifact ramp option.
- 1 Necklace of Girion | ramp | An artifact ramp piece for building resources.
- 1 Thranduil the Strategist | ramp | A Thranduil-themed ramp card that supports the commander plan.
- 1 Through the Forest Gate | ramp | A ramp spell that supports mana development.
- 1 Wayfarer's Bauble | ramp | An artifact ramp piece for early setup.
- 1 Wood Elves | ramp | An Elf Scout that advances mana while fitting the Elf theme.
- 1 Elvish Visionary | draw | An Elf Shaman that supplies card access while fitting the theme.
- 1 Fateful Discovery | draw | A dedicated card-access enchantment.
- 1 Hithlain Knots | draw | An instant-speed card-access option.
- 1 Key to the Side-Door | draw | An artifact source of card access.
- 1 Lórien Revealed | draw | A card-access spell that fits the deck's colors.
- 1 Night's Whisper | draw | An efficient card-access spell.
- 1 Palantír of Orthanc | draw | A legendary artifact source of card access.
- 1 Plunder the Trollshaws | draw | An instant card-access option.
- 1 Thrór's Map | draw | A legendary artifact source of card access.
- 1 Uncover the Moon-Letters | draw | An enchantment that supplies card access.
- 1 Bilbo's Ring | interaction | A legendary equipment piece for flexible interaction.
- 1 Dwarven Mattock | interaction | An equipment-based interaction option.
- 1 Elrond, Moon-Reader | interaction | An Elf Noble that contributes interaction to the creature plan.
- 1 Mithril Coat | interaction | A legendary equipment interaction piece.
- 1 The One Ring | interaction | A legendary artifact interaction option.
- 1 Thranduil's Decree | interaction | A Thranduil-themed instant for interaction.
- 1 Bilbo's Deadly Slice | removal | An instant removal option.
- 1 Bitter Downfall | removal | A dedicated removal spell.
- 1 Colossal Whale | removal | A creature-based removal option.
- 1 Crude Bent Blade | removal | An equipment-based removal option.
- 1 Enchanted River's Grasp | removal | An Aura removal option.
- 1 Giant's Boulder | removal | An artifact removal option.
- 1 Quarrel | removal | An instant removal option.
- 1 Witch-king of Angmar | removal | A legendary creature that adds removal to the deck.
- 1 Gnashing of Teeth | wipe | A dedicated board-clearing spell.
- 1 Languish | wipe | A dedicated board-clearing spell.
- 1 Raise the Palisade | wipe | A dedicated board-clearing spell.
- 1 Arwen, Weaver of Hope | synergy | An Elf Noble that strengthens the deck's Elf-centered identity.
- 1 Boughside Wanderers | synergy | An Elf Scout that contributes to the Elf creature base.
- 1 Cantankerous Keepers | synergy | An Elf Soldier that supports the Elf-focused board.
- 1 Celeborn the Wise | synergy | An Elf Noble that reinforces the deck's Elf theme.
- 1 Elven Raft-Steerer | synergy | An Elf Pilot that adds to the Elf creature density.
- 1 Elvenking's Harper | synergy | An Elf Bard that supports the Elvenking theme.
- 1 Galadhrim Guide | synergy | An Elf Scout that reinforces the Elf plan.
- 1 Galion, Elvenking's Butler | synergy | An Elf Advisor that supports the Elvenking theme.
- 1 Guardian of the Halls | synergy | An Elf Soldier that expands the Elf board presence.
- 1 Grey Havens Navigator | synergy | An Elf Pilot that contributes to the Elf creature base.
- 1 Lothlórien Lookout | synergy | An Elf Scout that reinforces the Elf plan.
- 1 Mirkwood Meditator | synergy | An Elf Druid that supports the deck's Elf density.
- 1 Mirkwood Nurturer | synergy | An Elf Ranger that strengthens the Elf-centered board.
- 1 Mirkwood Pathmaker | synergy | An Elf Ranger that supports the deck's central creature theme.
- 1 Attercop | threat | A creature threat that adds pressure to the board.
- 1 Gigantic Big Bear | threat | A creature threat that adds a substantial body to the board.
- 1 Great Fierce Bee | threat | A creature threat that broadens the deck's board presence.
- 1 Large Bear | threat | A creature threat that contributes to combat pressure.
- 1 Little Bear | threat | A creature threat that adds to the board early.
- 1 Mirkwood Elk | threat | A creature threat that supports the creature-forward plan.
- 1 Nasty Little Rabbit | threat | A creature threat that adds another body to the board.
- 1 Old Thrush | threat | A creature threat that broadens the creature suite.
- 1 Ordinary Bear | threat | A creature threat that contributes to board pressure.
- 1 Ravenhill Flock | threat | A creature threat that adds to the deck's board presence.
- 1 Troll of Khazad-dûm | threat | A creature threat that provides a larger standalone presence.
- 1 Willow-Wind | threat | An Elemental creature threat for the top end of the board.

</details>

### 20. the Hobbit family, a delegated commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 77 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $529.72 to buy, $529.72 the whole deck.

**Summary:** Smaug the Magnificent leads a dragon-centered deck that develops its mana, uses draw to keep its hand supplied, and clears resistance with removal and wipes. It wins by pressing its Dragon threats and Smaug-themed cards into a sustained offensive, while giving up some specialized Dragon support for a broad collection of Middle-earth equipment, creatures, and utility cards.

- JUDGE [true]: "Smaug the Magnificent leads a dragon-centered deck (implying it is eligible to be a commander)". Smaug the Magnificent is a legendary creature printed in the Lord of the Rings: Tales of Middle-earth set, making it eligible to serve as a commander.
- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 3.23 over 62 nonland cards
- [INFO] `basics_added`: the list was 1 card short, so the builder added 1 basic land

<details><summary>The deck list</summary>

- 33 Mountain | land | Provides the deck's primary land base.
- 1 Dragon-Cursed Halls | land | Provides a land slot.
- 1 Rogue's Passage | land | Provides a land slot.
- 1 The Lonely Mountain | land | Provides a land slot.
- 1 Treasure Vault | land | Provides a land slot.
- 1 Balin, Loremaster | draw | Included as a draw card.
- 1 Key to the Side-Door | draw | Included as a draw card.
- 1 Palantír of Orthanc | draw | Included as a draw card.
- 1 Ragged Short Spear | draw | Included as a draw card.
- 1 Thrór's Map | draw | Included as a draw card.
- 1 Óin the Brave | draw | Included as a draw card.
- 1 Bilbo's Ring | interaction | Included for interaction.
- 1 Dwarven Mattock | interaction | Included for interaction.
- 1 Mithril Coat | interaction | Included for interaction.
- 1 The One Ring | interaction | Included for interaction.
- 1 Arcane Signet | ramp | Included as ramp.
- 1 Bag End Banquet | ramp | Included as ramp.
- 1 Burn, Burn, Tree and Fern | ramp | Included as ramp.
- 1 Cavern-Hoard Dragon | ramp | A Dragon that also fills a ramp slot.
- 1 Dragon's Desire | ramp | Included as ramp.
- 1 Fíli and Kíli, Joyous | ramp | Included as ramp.
- 1 Mox Amber | ramp | Included as ramp.
- 1 Orcrist, Goblin-cleaver | ramp | Included as ramp.
- 1 The Misty Mountains Cold | ramp | Included as ramp.
- 1 The Reaver Cleaver | ramp | Included as ramp.
- 1 Thorin, Company's Leader | ramp | Included as ramp.
- 1 Troop of Ponies | ramp | Included as ramp.
- 1 Wayfarer's Bauble | ramp | Included as ramp.
- 1 Battle-Scarred Goblin | removal | Included as removal.
- 1 Fire of Orthanc | removal | Included as removal.
- 1 Gandalf, Spark Starter | removal | Included as removal.
- 1 Giant's Boulder | removal | Included as removal.
- 1 Goblin Cratermaker | removal | Included as removal.
- 1 Improvised Club | removal | Included as removal.
- 1 Inferno Titan | removal | Included as removal.
- 1 Smaug, the Great Calamity // Spew Flame | removal | A Smaug-themed removal card.
- 1 Last Light of Durin's Day | synergy | Included for synergy.
- 1 Desert Were-Worm | threat | Included as a threat.
- 1 Olog-hai Crusher | threat | Included as a threat.
- 1 Call Forth the Tempest | wipe | Included as a wipe.
- 1 Desolation of Smaug | wipe | A Smaug-themed wipe.
- 1 Glóin the Mighty // Easy Pickings | wipe | Included as a wipe.
- 1 Andúril, Flame of the West | other | Included as a supporting card.
- 1 Andúril, Narsil Reforged | other | Included as a supporting card.
- 1 Bombur, Gentle Dreamer | other | Included as a supporting card.
- 1 Bothersome Noisemaker | other | Included as a supporting card.
- 1 Dori, Bearer of Friends | other | Included as a supporting card.
- 1 Dwarven Mauler | other | Included as a supporting card.
- 1 Dwarven Warriors | other | Included as a supporting card.
- 1 Dáin Ironfoot | other | Included as a supporting card.
- 1 Gandalf, Goblins' Bane // Flameshape | other | Included as a supporting card.
- 1 Getaway Barrel | other | Included as a supporting card.
- 1 Glamdring | other | Included as a supporting card.
- 1 Goblin-town Flunkies | other | Included as a supporting card.
- 1 Gundabad Opportunist | other | Included as a supporting card.
- 1 Guttersnipe | other | Included as a supporting card.
- 1 Iron Hills Stalwart | other | Included as a supporting card.
- 1 Long-Lost Lances | other | Included as a supporting card.
- 1 Misty Mountains Raider | other | Included as a supporting card.
- 1 Old Thrush | other | Included as a supporting card.
- 1 Oliphaunt | other | Included as a supporting card.
- 1 Orcish Siegemaster | other | Included as a supporting card.
- 1 Smaug's Fury | other | A Smaug-themed supporting card.
- 1 Snowslope Hunter | other | Included as a supporting card.
- 1 Sting, Bilbo's Sword | other | Included as a supporting card.
- 1 Tidings of War | other | Included as a supporting card.
- 1 Well-Worn Spatula | other | Included as a supporting card.

</details>

### 21. the Hobbit family, mana from outside

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 156 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $631.58 to buy, $631.58 the whole deck.

**Summary:** This Smaug deck builds a red-black battlefield around Dragons, Goblins, Orcs, Trolls, and a dense set of legendary artifacts and Equipment. It develops mana into large themed threats, clears away resistance with focused removal and sweeping spells, then wins through sustained combat pressure. The deck gives up some flexibility for its strong Hobbit-set flavor and its commitment to creature combat, artifacts, and themed legends.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 3.56 over 63 nonland cards
- [WARN] `outside_requested_set`: the sets you named do not hold 8 cards: Blood Crypt, Bojuka Bog, Command Tower, Dragonskull Summit, Myriad Landscape, Path of Ancestry, Reliquary Tower, Temple of the False God

<details><summary>The deck list</summary>

- 1 Blood Crypt | land | Red-black source for the deck’s mana base.
- 1 Bojuka Bog | land | Black land slot with a focused utility role.
- 1 Command Tower | land | Reliable multicolor mana for the commander’s colors.
- 1 Dragon-Cursed Halls | land | Themed land supporting the Dragon-led deck.
- 1 Dragonskull Summit | land | Red-black dual land for consistent casting.
- 1 Goblin-town | land | Themed land for the Goblin contingent.
- 1 Minas Morgul, Dark Fortress | land | Black-aligned utility land.
- 1 Mount Doom | land | Themed legendary land in the mana base.
- 1 Myriad Landscape | land | Land slot that supports the basic-land base.
- 1 Path of Ancestry | land | Themed fixing land for the creature suite.
- 1 Reliquary Tower | land | Colorless utility land.
- 1 Rogue's Passage | land | Utility land that supports combat pressure.
- 1 Temple of the False God | land | Colorless land slot for the mana base.
- 1 The Black Gate | land | Themed legendary land.
- 1 The Lonely Mountain | land | Smaug-themed Mountain land.
- 1 Treasure Vault | land | Artifact land that fits the deck’s treasure-and-artifact flavor.
- 12 Mountain | land | Primary basic source for the red-heavy deck.
- 8 Swamp | land | Basic black sources for the deck’s black spells.
- 1 Arcane Signet | ramp | Efficient color fixing and acceleration.
- 1 Bag End Banquet | ramp | Artifact-based acceleration.
- 1 Burn, Burn, Tree and Fern | ramp | Themed acceleration piece.
- 1 Cavern-Hoard Dragon | ramp | A Dragon that contributes to the deck’s mana development.
- 1 Dragon's Desire | ramp | Dragon-themed acceleration.
- 1 Mox Amber | ramp | Legendary-focused artifact acceleration.
- 1 Orcrist, Goblin-cleaver | ramp | Equipment that also supports mana development.
- 1 Smaug the Magnificent | ramp | Additional Smaug-themed Dragon acceleration.
- 1 Smaug, Wicked Worm | ramp | A second Smaug that advances the mana plan.
- 1 The Reaver Cleaver | ramp | Themed Equipment acceleration.
- 1 Balin, Loremaster | draw | Dwarf legend included for card access.
- 1 Gollum, Riddle Master | draw | Themed legendary source of card access.
- 1 Key to the Side-Door | draw | Artifact card-access piece.
- 1 Night's Whisper | draw | Straightforward black card access.
- 1 Palantír of Orthanc | draw | Legendary artifact card-access engine.
- 1 Rage into the Valley | draw | Red card-access spell.
- 1 Ragged Short Spear | draw | Equipment that contributes to card access.
- 1 Reverent Howl | draw | Card-access spell in the supporting package.
- 1 Thrór's Map | draw | Themed legendary artifact for card access.
- 1 Óin the Brave | draw | Dwarf legend providing another card-access slot.
- 1 Bilbo's Ring | interaction | Legendary Equipment utility.
- 1 Dwarven Mattock | interaction | Equipment-based utility for the creature plan.
- 1 Getaway Barrel | interaction | Flexible artifact utility.
- 1 Mithril Coat | interaction | Protective legendary Equipment utility.
- 1 My Precious // Allure of Power | interaction | Legendary Equipment with an interactive Adventure option.
- 1 The One Ring | interaction | Powerful legendary artifact utility.
- 1 Azog, Moria's Ruin | removal | Themed legendary removal creature.
- 1 Battle-Scarred Goblin | removal | Goblin creature that answers opposing pieces.
- 1 Bilbo's Deadly Slice | removal | Efficient themed removal spell.
- 1 Bitter Downfall | removal | Direct black removal.
- 1 Bolg of the North | removal | Legendary Goblin removal option.
- 1 Fire of Orthanc | removal | Red removal spell.
- 1 Orcish Bowmasters | removal | Orc creature that supplies removal pressure.
- 1 Smaug, the Great Calamity // Spew Flame | removal | Smaug-themed Dragon with a removal Adventure.
- 1 Bolg, Erebor's Reckoning | wipe | Legendary Goblin battlefield reset.
- 1 Call Forth the Tempest | wipe | Themed mass-reset spell.
- 1 Desolation of Smaug | wipe | Smaug-themed mass-reset spell.
- 1 Bothersome Noisemaker | synergy | Goblin Bard supporting the Goblin creature package.
- 1 Crude Bent Blade | synergy | Equipment supporting the artifact-and-combat package.
- 1 Down, Down to Goblin-town | synergy | Goblin-themed Saga for the deck’s supporting package.
- 1 Fearsome Goblin Pair | synergy | Goblin creature supporting the tribal contingent.
- 1 Goblin Plate Mail | synergy | Goblin-themed Equipment for the combat package.
- 1 Goblin-town Flunkies | synergy | Goblin body supporting the themed creature base.
- 1 Great Ugly-Looking Goblin // Clap! Snap! | synergy | Goblin creature with a themed Adventure.
- 1 Gundabad Opportunist | synergy | Goblin Rogue for the Goblin package.
- 1 Guttersnipe | synergy | Goblin Shaman supporting the spell-heavy portions of the deck.
- 1 Misty Mountains Raider | synergy | Goblin Soldier for the themed creature suite.
- 1 Orcish Siegemaster | synergy | Orc Soldier supporting the hostile-army theme.
- 1 Snowslope Hunter | synergy | Goblin Ranger supporting the Goblin package.
- 1 Stony-Voiced Goblins | synergy | Goblin Bard adding to the themed creature density.
- 1 Well-Worn Spatula | synergy | Equipment supporting the artifact-and-combat package.
- 1 Desert Were-Worm | threat | Dragon Wurm that serves as a large themed attacker.
- 1 Dreaded Bat-Cloud | threat | Black creature threat for the battlefield plan.
- 1 Gollum the Abandoned | threat | Legendary Horror that adds pressure.
- 1 Great Goblin, Foul-Hearted | threat | Legendary Goblin threat fitting the deck’s creature theme.
- 1 Haunt of the Dead Marshes | threat | Nightmare creature that broadens the threat base.
- 1 Inferno Titan | threat | Large red Giant for direct combat pressure.
- 1 Nighthowl Pursuer | threat | Wolf threat in the black creature package.
- 1 Olog-hai Crusher | threat | Troll Soldier that provides a substantial body.
- 1 Ravening Warg | threat | Warg creature for the aggressive creature plan.
- 1 Sauron, the Lidless Eye | threat | Legendary Avatar Horror as a top-end themed threat.
- 1 Stone-Giant of High Pass | threat | Giant threat that also fits the deck’s hostile-forces theme.
- 1 Troll of Khazad-dûm | threat | Troll creature rounding out the large-creature package.

</details>

### 22. a set family and a card from outside it

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 175 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $510.54 to buy, $510.54 the whole deck.

**Summary:** This deck develops an Elf-centered board under Thranduil, using ramp to establish its mana and draw to keep the pressure coming. It wins primarily by building a wide collection of themed creature threats and turning that board into combat pressure, while removal, interaction, and wipes keep opposing boards from taking over. It gives up alternate finish lines and some card flexibility for a focused creature-heavy Middle-earth plan.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 3.38 over 63 nonland cards
- [WARN] `outside_requested_set`: the sets you named do not hold 1 card: Sol Ring

<details><summary>The deck list</summary>

- 1 Dragon-Cursed Halls | land | A named land that supports the three-color mana base.
- 1 Elven Passage | land | A named land that supports the three-color mana base.
- 1 Elvenking's Halls | land | A named land that supports the three-color mana base.
- 1 Hobbit Hole | land | A named land that supports the three-color mana base.
- 1 Minas Morgul, Dark Fortress | land | A named land that supports the three-color mana base.
- 1 Mirkwood | land | A named land that supports the three-color mana base.
- 1 Rivendell | land | A named land that supports the three-color mana base.
- 1 Rogue's Passage | land | A named land that supports the three-color mana base.
- 1 The Black Gate | land | A named land that supports the three-color mana base.
- 1 The Shire | land | A named land that supports the three-color mana base.
- 1 Treasure Vault | land | A named land that supports the three-color mana base.
- 11 Forest | land | Basic green lands form the foundation of the mana base.
- 8 Island | land | Basic blue lands support the mana base.
- 6 Swamp | land | Basic black lands support the mana base.
- 1 Sol Ring | ramp | Fast artifact ramp kept as requested.
- 1 Arcane Signet | ramp | A compact ramp piece for early development.
- 1 Elven Chorus | ramp | Ramp that fits the Elf-centered plan.
- 1 Elvish Archdruid | ramp | An Elf ramp piece that reinforces the creature theme.
- 1 Elvish Mystic | ramp | An efficient Elf ramp piece for early turns.
- 1 Mox Amber | ramp | A compact legendary-themed ramp piece.
- 1 Necklace of Girion | ramp | Artifact ramp that helps advance the board.
- 1 Silvan Reveler | ramp | An Elf ramp card that supports the main creature plan.
- 1 Wayfarer's Bauble | ramp | Reliable artifact ramp for the mana base.
- 1 Wood Elves | ramp | An Elf ramp creature that develops the board.
- 1 Elvish Visionary | draw | An Elf that supplies card draw while contributing to the theme.
- 1 Fateful Discovery | draw | Dedicated card draw for maintaining resources.
- 1 Hithlain Knots | draw | A draw spell that keeps cards flowing.
- 1 Key to the Side-Door | draw | An artifact draw option for continued resources.
- 1 Last March of the Ents | draw | A larger draw effect for refilling the hand.
- 1 Lórien Revealed | draw | A draw spell that supports consistent access to cards.
- 1 Night's Whisper | draw | A direct draw spell for efficient card access.
- 1 Palantír of Orthanc | draw | A legendary artifact draw source.
- 1 Plunder the Trollshaws | draw | Instant-speed draw support for the deck.
- 1 Uncover the Moon-Letters | draw | An enchantment-based draw piece.
- 1 Bilbo's Ring | interaction | Interaction that also fits the Middle-earth equipment suite.
- 1 Elrond, Moon-Reader | interaction | An Elf legend providing interaction within the theme.
- 1 Mithril Coat | interaction | Equipment-based interaction for protecting the board plan.
- 1 My Precious // Allure of Power | interaction | A flexible legendary equipment interaction piece.
- 1 The One Ring | interaction | A legendary artifact interaction option.
- 1 Thranduil's Decree | interaction | Thematic interaction tied to Thranduil.
- 1 Bilbo's Deadly Slice | removal | A focused removal spell for opposing problems.
- 1 Bitter Downfall | removal | Direct removal to answer key opposing cards.
- 1 Crude Bent Blade | removal | Equipment-based removal that remains on theme.
- 1 Enchanted River's Grasp | removal | Enchantment removal that answers troublesome permanents.
- 1 Giant's Boulder | removal | Artifact removal support for the deck.
- 1 Orcish Bowmasters | removal | Creature-based removal that adds to board presence.
- 1 Quarrel | removal | Efficient removal for opposing threats.
- 1 Witch-king of Angmar | removal | A removal creature that adds a substantial body.
- 1 Gnashing of Teeth | wipe | A board wipe for resetting crowded tables.
- 1 Languish | wipe | A board wipe for recovering from opposing creature boards.
- 1 Raise the Palisade | wipe | A board wipe that helps clear a path for the creature plan.
- 1 Supper for Spiders | synergy | The dedicated synergy card from the shortlist.
- 1 Arwen, Weaver of Hope | synergy | An Elf legend that strengthens the Elf-centered shell.
- 1 Boughside Wanderers | synergy | An Elf body that contributes to the tribal board.
- 1 Cantankerous Keepers | synergy | An Elf creature that reinforces the main theme.
- 1 Celeborn the Wise | synergy | An Elf legend that deepens the Middle-earth Elf focus.
- 1 Elven Raft-Steerer | synergy | An Elf creature that adds to the themed board.
- 1 Elvenking's Harper | synergy | An Elf that supports Thranduil's Elven court theme.
- 1 Galadhrim Guide | synergy | An Elf Scout that builds the tribal board.
- 1 Galion, Elvenking's Butler | synergy | An Elf legend directly tied to the Elvenking theme.
- 1 Grey Havens Navigator | synergy | An Elf creature that reinforces the Elf concentration.
- 1 Guardian of the Halls | synergy | An Elf Soldier that supports the creature-based plan.
- 1 Lothlórien Lookout | synergy | An Elf Scout that adds to the themed creature base.
- 1 Mirkwood Meditator | synergy | An Elf Druid that supports the Mirkwood Elf theme.
- 1 Mirkwood Nurturer | synergy | An Elf Ranger that helps maintain the tribal density.
- 1 Attercop | threat | A sizeable creature threat for applying pressure.
- 1 Gigantic Big Bear | threat | A large creature that serves as a combat threat.
- 1 Great Fierce Bee | threat | A creature threat that broadens the board.
- 1 Large Bear | threat | A straightforward creature threat for combat.
- 1 Mirkwood Elk | threat | A creature threat that matches the Mirkwood setting.
- 1 Mirkwood Pathmaker | threat | An Elf threat that continues the tribal plan.
- 1 Nimrodel Watcher | threat | An Elf threat that adds to creature pressure.
- 1 Old Fat Spider | threat | A creature threat that supplements the board.
- 1 Ordinary Bear | threat | A creature threat that supports combat pressure.
- 1 Thranduil, Sindarin Liege // Silvan Rally | threat | A Thranduil-themed Elf threat and a central thematic payoff.
- 1 Troll of Khazad-dûm | threat | A substantial creature threat for closing games.
- 1 Willow-Wind | threat | A creature threat that rounds out the attack plan.

</details>

### 23. two set families at once

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 203 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $208.11 to buy, $208.11 the whole deck.

**Summary:** Kíli leads a white Dwarf-and-Equipment deck that develops through lands and mana artifacts, maintains resources with a broad draw package, and turns its creature core into steady combat pressure. It supports that plan with direct answers, protection, and several ways to reset crowded boards. The deck gives up some speed for a more board-focused game that rewards building a resilient battlefield before committing to combat.

- [INFO] `bracket_prose_rules`: bracket 3 (Upgraded, 6+ turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified 2026-08-26
- [INFO] `curve_summary`: average mana value 2.86 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Castle Ardenvale | land | A land slot with a named utility option.
- 1 Command Tower | land | A land slot for the commander’s colors.
- 1 Dragon-Cursed Halls | land | A named land slot.
- 1 Elven Passage | land | A named land slot.
- 1 Evolving Wilds | land | A land slot that helps stabilize mana.
- 1 Exotic Orchard | land | A flexible land slot.
- 1 Fabled Passage | land | A land slot that helps stabilize mana.
- 1 Fountainport | land | A named utility land slot.
- 1 Hidden Grotto | land | A named land slot.
- 1 Hobbit Hole | land | A thematic named land slot.
- 1 Lupinflower Village | land | A named land slot.
- 1 Minas Tirith | land | A thematic legendary land slot.
- 1 Path of Ancestry | land | A tribal-minded land slot.
- 1 Reliquary Tower | land | A utility land slot.
- 1 Rogue's Passage | land | A utility land slot for combat pressure.
- 1 Swarmyard | land | A named utility land slot.
- 1 Terramorphic Expanse | land | A land slot that helps stabilize mana.
- 1 Three Tree City | land | A thematic legendary land slot.
- 1 Thriving Heath | land | A land slot for consistent white mana.
- 1 Treasure Vault | land | A utility artifact land slot.
- 1 Uncharted Haven | land | A flexible land slot.
- 15 Plains | land | The basic-land foundation for the deck’s mana base.
- 1 Arcane Signet | ramp | An efficient mana artifact.
- 1 Bag End Banquet | ramp | A thematic ramp artifact.
- 1 Burnished Hart | ramp | A creature-based ramp piece.
- 1 Fellwar Stone | ramp | An efficient mana artifact.
- 1 Gilded Lotus | ramp | A high-output mana artifact.
- 1 Hedron Archive | ramp | A mana artifact that remains useful later.
- 1 Mind Stone | ramp | A compact mana artifact.
- 1 Ornithopter of Paradise | ramp | A creature-based mana source.
- 1 Patchwork Banner | ramp | A tribal-minded mana artifact.
- 1 Wayfarer's Bauble | ramp | A basic-land ramp piece.
- 1 Caretaker's Talent | draw | A repeatable draw-oriented enchantment.
- 1 Circuit Mender | draw | A creature that contributes to card flow.
- 1 Cut a Deal | draw | A straightforward draw spell.
- 1 Dawn of a New Age | draw | An enchantment-based draw option.
- 1 Errand-Rider of Gondor | draw | A thematic creature that contributes card flow.
- 1 Esgaroth Garrison | draw | A thematic creature that contributes card flow.
- 1 Fountainport Bell | draw | An artifact draw option.
- 1 Heirloom Epic | draw | A thematic artifact draw piece.
- 1 Idol of Oblivion | draw | A compact artifact draw piece.
- 1 Inspiring Overseer | draw | A creature-based draw option.
- 1 Bilbo's Gambit | interaction | A thematic interactive instant.
- 1 Bofur, Reliable Guardian // Concerted Care | interaction | A thematic protection-oriented interaction piece.
- 1 Crumb and Get It | interaction | A flexible interactive instant.
- 1 Dawn's Truce | interaction | A protective interaction spell.
- 1 Dwarven Mattock | interaction | A thematic Dwarf interaction piece.
- 1 Flowering of the White Tree | interaction | A board-supporting interaction piece.
- 1 Angel of the Ruins | removal | A creature-based answer to opposing permanents.
- 1 Banishing Light | removal | A broad permanent-answer enchantment.
- 1 Celebrate the Mountain-king | removal | A thematic removal enchantment.
- 1 Fiend Hunter | removal | A creature-based removal option.
- 1 Generous Gift | removal | A flexible answer to opposing permanents.
- 1 Loran of the Third Path | removal | A thematic creature-based removal piece.
- 1 Skyclave Apparition | removal | A creature-based permanent answer.
- 1 Swords to Plowshares | removal | An efficient creature answer.
- 1 Blade Splicer | synergy | A creature that supports the deck’s artifact-oriented core.
- 1 Carrot Cake | synergy | A thematic artifact synergy piece.
- 1 Dáin, Lord of the Iron Hills | synergy | A thematic Dwarf synergy card.
- 1 Iron Hills Blacksmith | synergy | A Dwarf Artificer that supports the deck’s core.
- 1 Ori, Keeper of Songs | synergy | A thematic Dwarf synergy card.
- 1 Tangle Tumbler | synergy | An artifact synergy piece.
- 1 Fíli the Pathfinder | threat | A thematic legendary threat.
- 1 Karn, the Great Creator | threat | A resilient artifact-focused threat.
- 1 Dusk // Dawn | wipe | A flexible board-reset option.
- 1 Martial Coup | wipe | A board-reset spell that also develops the board.
- 1 Ori, Plate Stacker | wipe | A thematic Dwarf board-reset option.
- 1 An Unexpected Party // At the Door | other | A thematic Adventure card for the deck’s supporting package.
- 1 Andúril, Flame of the West | other | A thematic legendary Equipment for the combat package.
- 1 Andúril, Narsil Reforged | other | A thematic legendary Equipment for the combat package.
- 1 Bilbo, Unexpected Adventurer | other | A thematic legendary creature for the supporting package.
- 1 Dúnedain Blade | other | A thematic Equipment for creature combat.
- 1 Dwarven Provisioner | other | A thematic Dwarf creature for the deck’s core.
- 1 Dwarven Shortsword | other | A thematic Equipment for the combat package.
- 1 Glamdring | other | A thematic legendary Equipment for the combat package.
- 1 Long-Lost Lances | other | A thematic Equipment for creature combat.
- 1 Moment of Glory | other | A thematic supporting spell.
- 1 Murmuration | other | A Bloomburrow-themed supporting enchantment.
- 1 Short Bow | other | A Bloomburrow-themed Equipment.
- 1 Sting, Bilbo's Sword | other | A thematic legendary Equipment for the combat package.
- 1 Sword of the Squeak | other | A Bloomburrow-themed Equipment for creature combat.
- 1 The Eagles Are Coming! | other | A thematic supporting instant.
- 1 Vow to Erebor | other | A thematic supporting instant.
- 1 Well-Worn Spatula | other | A Bloomburrow-themed Equipment.
- 1 Whiskervale Forerunner | other | A Bloomburrow-themed creature for the supporting package.

</details>

### 24. a 60-card deck from one set

Format: Modern. Theme: aggro. Pool: any_card. Shortlist: 96 names.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $37.46 to buy, $37.46 the whole deck.

**Summary:** This mono-red Bloomburrow aggro deck applies pressure with a dense creature suite, led by its synergy creatures and backed by a broad range of threats. It aims to win through sustained combat pressure, using removal to keep attacks moving and draw spells to avoid running out of action. Its focused mana base and proactive construction give up broader answers and defensive tools in favor of speed and consistency.

- [INFO] `curve_summary`: average mana value 3.11 over 36 nonland cards

<details><summary>The deck list</summary>

- 24 Mountain | land | Provides the consistent red mana base for a mono-red deck.
- 4 Might of the Meek | draw | Supplies low-cost card access while supporting the creature-focused plan.
- 2 Sazacap's Brew | draw | Adds further card access without moving away from red.
- 2 Patchwork Banner | ramp | Provides the requested acceleration while fitting the creature-heavy build.
- 4 Agate Assault | removal | Gives the deck efficient red removal to clear opposing resistance.
- 2 Blooming Blast | removal | Rounds out the removal suite with another red answer.
- 4 Emberheart Challenger | synergy | Forms a core synergy creature for the aggressive creature plan.
- 4 Hearthborn Battler | synergy | Adds another full set of creatures carrying the deck's synergy role.
- 4 Frilled Sparkshooter | threat | Provides a full set of aggressive threats.
- 4 Reptilian Recruiter | threat | Adds more creature pressure as a primary threat.
- 4 Stormsplitter | threat | Supplies another concentrated set of red threats.
- 2 Dragonhawk, Fate's Tempest | threat | Serves as a higher-impact threat to finish games.

</details>

