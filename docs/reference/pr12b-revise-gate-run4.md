# PR-12B revise gate

Run date: 2026-09-01. Snapshot: 2026-08-31.

Verdict: PASS. 8 of 8 revisions met their bar. The bars: an unclear request gets a question or a decline with a reason, a clear request gets a revised deck with no block finding, the deck holds every cap and every removal the message names, and it keeps at least 80 percent of the untouched cards.

Cost: 19 calls, $0.5448. Time: 5m53s.

| Base | Revision | Message | Outcome | Kept | Result |
|---|---|---|---|---|---|
| 1 | 1 | Replace some lands with better options if possible. Also tune the mana curve lower - no 6 or 7 mana cards needed | question |  | PASS, no build |
| 1 | 2 | Take out Angel of Serenity. | 60 cards | 100% | PASS, kept 100% |
| 1 | 3 | Better lands please. | question |  | PASS, no build |
| 2 | 4 | No nonland cards over 4 mana. | 99 cards | 100% | PASS, kept 100% |
| 2 | 5 | Swap Defiant Bloodlord for something cheaper that does the same job. | 99 cards | 100% | PASS, kept 100% |
| 2 | 6 | Add two more cards that draw cards, and cut two of the weakest creatures. | 99 cards | 98% | PASS, kept 98% |
| 3 | 7 | Lower the mana curve. Nothing over 5 mana. | 99 cards | 100% | PASS, kept 100% |
| 3 | 8 | Take out The Lord of the Eagles. | 99 cards | 100% | PASS, kept 100% |

## Revision 1, base 1: Modern Angels, casual, mono-white, any card

**The user:** Replace some lands with better options if possible. Also tune the mana curve lower - no 6 or 7 mana cards needed

The brief:

```json
{
  "changes": [
    "Remove every nonland card with mana value 6 or more and replace them with lower-cost Angel and +1/+1 counter support."
  ],
  "remove": null,
  "keep": null,
  "max_mana_value": 5,
  "question": "What kind of land upgrades do you want: faster mana, utility lands, or lands that support additional colors?",
  "declined": []
}
```

**The reply:** Question: What kind of land upgrades do you want: faster mana, utility lands, or lands that support additional colors?


## Revision 2, base 1: Modern Angels, casual, mono-white, any card

**The user:** Take out Angel of Serenity.

The brief:

```json
{
  "changes": [],
  "remove": [
    "Angel of Serenity"
  ],
  "keep": null,
  "max_mana_value": 0,
  "question": "",
  "declined": []
}
```

**The reply:** I removed 1 Angel of Serenity. I added 1 Angel of Destiny.

Findings: curve_summary (INFO), cards_trimmed (INFO)

The deck:

- 1 Abandoned Air Temple
- 1 Abzan Falconer
- 1 Admonition Angel
- 1 Aetherstorm Roc
- 1 Ainok Bond-Kin
- 1 Angel of Condemnation
- 1 Angel of Destiny
- 1 Angel of Invention
- 1 Angel of Sanctions
- 1 Angel of the God-Pharaoh
- 1 Angel of the Ruins
- 1 Arcane Signet
- 1 Archangel of Thune
- 1 Bruna, the Fading Light
- 1 Cathars' Crusade
- 1 Cavern of Souls
- 1 Command Tower
- 1 Exemplar of Light
- 1 Felidar Retreat
- 1 Giada, Font of Hope
- 1 Guildmages' Forum
- 1 Hall of Oracles
- 1 Herald of War
- 1 Herald of the Sun
- 1 Idyllic Grange
- 1 Inspiring Overseer
- 1 Karn's Bastion
- 1 Light of the Legion
- 1 Luminarch Aspirant
- 1 Lyra Dawnbringer
- 1 Mikaeus, the Lunarch
- 1 Paragon of Modernity
- 1 Patron of the Valiant
- 1 Righteous Valkyrie
- 1 Rogue's Passage
- 1 Ruins of Oran-Rief
- 1 Sanctuary Warden
- 1 Sephara, Sky's Blade
- 1 Seraph Sanctuary
- 1 Serra Redeemer
- 1 Sunblast Angel
- 1 Tyrite Sanctum
- 1 Urza's Saga
- 1 Winged Shepherd
- 1 Wojek Investigator
- 1 Youthful Valkyrie
- 14 Plains

## Revision 3, base 1: Modern Angels, casual, mono-white, any card

**The user:** Better lands please.

The brief:

```json
{
  "changes": [],
  "remove": null,
  "keep": null,
  "max_mana_value": 0,
  "question": "What kind of land upgrade do you want: more reliable white mana, lands that help cast larger Angels faster, or utility lands with useful effects?",
  "declined": []
}
```

**The reply:** Question: What kind of land upgrade do you want: more reliable white mana, lands that help cast larger Angels faster, or utility lands with useful effects?


## Revision 4, base 2: lifegain Commander, Karlov, bracket 3, any card

**The user:** No nonland cards over 4 mana.

The brief:

```json
{
  "changes": [
    "Remove every nonland card with mana value 5 or more."
  ],
  "remove": null,
  "keep": null,
  "max_mana_value": 4,
  "question": "",
  "declined": []
}
```

**The reply:** I removed 1 Angel of Indemnity, 1 Archangel of Thune, 1 Astarion, the Decadent, 1 Carmen, Cruel Skymarcher, 1 Cavalier of Night, 1 Defiant Bloodlord, 1 Divinity of Pride, 1 Fumigate, 1 Liesa, Forgotten Archangel, 1 Lyra Dawnbringer, 1 Nykthos Paragon, 1 Sanguine Bond, 1 Solitude, 1 Vein Ripper, and 1 Vona, Butcher of Magan. I added 1 Archon of Sun's Grace, 1 Battle Angels of Tyr, 1 Enduring Tenacity, 1 Foolish Fate, 1 Gideon's Company, 1 Gisela, the Broken Blade, 1 Gumdrop Poisoner // Tempt with Treats, 1 Henrika Domnathi // Henrika, Infernal Seer, 1 Marauding Blight-Priest, 1 Nightmare's Thirst, 1 Phial of Galadriel, 1 Qala, Ajani's Pridemate, 1 The Meathook Massacre, 1 Twinblade Paladin, and 1 Wax-Wane Witness.

Findings: bracket_prose_rules (INFO), curve_summary (INFO)

The deck:

- 1 Aclazotz, Deepest Betrayal // Temple of the Dead
- 1 Adventurer's Inn
- 1 Aerith Gainsborough
- 1 Aetherflux Reservoir
- 1 Ajani's Pridemate
- 1 Ajani, Strength of the Pride
- 1 Alseid of Life's Bounty
- 1 Altar of the Pantheon
- 1 Angel of Vitality
- 1 Archivist of Oghma
- 1 Archon of Sun's Grace
- 1 Arguel's Blood Fast // Temple of Aclazotz
- 1 Attended Healer
- 1 Ayli, Eternal Pilgrim
- 1 Azor's Gateway // Sanctum of the Sun
- 1 Battle Angels of Tyr
- 1 Beza, the Bounding Spring
- 1 Blighted Steppe
- 1 Bloodbond Vampire
- 1 Bounty Board
- 1 Caduceus, Staff of Hermes
- 1 Cleric Class
- 1 Cleric of Life's Bond
- 1 Cliffhaven Vampire
- 1 Consuming Corruption
- 1 Cosmos Elixir
- 1 Crypt Ghast
- 1 Dawn of Hope
- 1 Diamond Valley
- 1 Elenda, Saint of Dusk
- 1 Enduring Innocence
- 1 Enduring Tenacity
- 1 Exemplar of Light
- 1 Faith's Shield
- 1 Fisk Tower
- 1 Foolish Fate
- 1 Foot Headquarters
- 1 Gideon's Company
- 1 Gisela, the Broken Blade
- 1 Glasswing Grace // Age-Graced Chapel
- 1 Glimmerpost
- 1 Gumdrop Poisoner // Tempt with Treats
- 1 Heliod, Sun-Crowned
- 1 Henrika Domnathi // Henrika, Infernal Seer
- 1 Hidden Hideout
- 1 Hierophant's Chalice
- 1 High Market
- 1 Hostile Hostel // Creeping Inn
- 1 Interplanar Beacon
- 1 Inventors' Fair
- 1 Kabira Crossroads
- 1 Kaya's Wrath
- 1 Legion's Landing // Adanto, the First Fort
- 1 Leyline of Hope
- 1 Mangara, the Diplomat
- 1 Marauding Blight-Priest
- 1 Markov Purifier
- 1 Metropolis Reformer
- 1 Miren, the Moaning Well
- 1 Murderous Rider // Swift End
- 1 Neglected Manor
- 1 Nightmare's Thirst
- 1 Nomad Stadium
- 1 Nuka-Cola Vending Machine
- 1 Obscura Storefront
- 1 Phial of Galadriel
- 1 Phyrexia's Core
- 1 Plaza of Harmony
- 1 Pristine Talisman
- 1 Qala, Ajani's Pridemate
- 1 Radiant Fountain
- 1 Resplendent Angel
- 1 Restless Fortress
- 1 Rhox Faithmender
- 1 Righteous Valkyrie
- 1 Rune-Tail, Kitsune Ascendant // Rune-Tail's Essence
- 1 Scoured Barrens
- 1 Seraph Sanctuary
- 1 Serra Ascendant
- 1 Shambling Vent
- 1 Sigarda's Splendor
- 1 Skyclave Cleric // Skyclave Basilica
- 1 Springjack Pasture
- 1 Starlit Sanctum
- 1 Sword of Light and Shadow
- 1 The Celestus
- 1 The Gaffer
- 1 The Meathook Massacre
- 1 Tomb of the Spirit Dragon
- 1 Twinblade Paladin
- 1 Vault of the Archangel
- 1 Vito, Thorn of the Dusk Rose
- 1 Vizkopa Guildmage
- 1 Voice of the Blessed
- 1 Wax-Wane Witness
- 1 Well of Lost Dreams
- 1 Westvale Abbey // Ormendahl, Profane Prince
- 1 Witch's Clinic
- 1 Zof Consumption // Zof Bloodbog

## Revision 5, base 2: lifegain Commander, Karlov, bracket 3, any card

**The user:** Swap Defiant Bloodlord for something cheaper that does the same job.

The brief:

```json
{
  "changes": [
    "Replace Defiant Bloodlord with a lower-mana lifegain payoff that makes opponents lose life when you gain life."
  ],
  "remove": [
    "Defiant Bloodlord"
  ],
  "keep": null,
  "max_mana_value": 0,
  "question": "",
  "declined": []
}
```

**The reply:** I removed 1 Defiant Bloodlord. I added 1 Marauding Blight-Priest.

Findings: bracket_prose_rules (INFO), curve_summary (INFO)

The deck:

- 1 Aclazotz, Deepest Betrayal // Temple of the Dead
- 1 Adventurer's Inn
- 1 Aerith Gainsborough
- 1 Aetherflux Reservoir
- 1 Ajani's Pridemate
- 1 Ajani, Strength of the Pride
- 1 Alseid of Life's Bounty
- 1 Altar of the Pantheon
- 1 Angel of Indemnity
- 1 Angel of Vitality
- 1 Archangel of Thune
- 1 Archivist of Oghma
- 1 Arguel's Blood Fast // Temple of Aclazotz
- 1 Astarion, the Decadent
- 1 Attended Healer
- 1 Ayli, Eternal Pilgrim
- 1 Azor's Gateway // Sanctum of the Sun
- 1 Beza, the Bounding Spring
- 1 Blighted Steppe
- 1 Bloodbond Vampire
- 1 Bounty Board
- 1 Caduceus, Staff of Hermes
- 1 Carmen, Cruel Skymarcher
- 1 Cavalier of Night
- 1 Cleric Class
- 1 Cleric of Life's Bond
- 1 Cliffhaven Vampire
- 1 Consuming Corruption
- 1 Cosmos Elixir
- 1 Crypt Ghast
- 1 Dawn of Hope
- 1 Diamond Valley
- 1 Divinity of Pride
- 1 Elenda, Saint of Dusk
- 1 Enduring Innocence
- 1 Exemplar of Light
- 1 Faith's Shield
- 1 Fisk Tower
- 1 Foot Headquarters
- 1 Fumigate
- 1 Glasswing Grace // Age-Graced Chapel
- 1 Glimmerpost
- 1 Heliod, Sun-Crowned
- 1 Hidden Hideout
- 1 Hierophant's Chalice
- 1 High Market
- 1 Hostile Hostel // Creeping Inn
- 1 Interplanar Beacon
- 1 Inventors' Fair
- 1 Kabira Crossroads
- 1 Kaya's Wrath
- 1 Legion's Landing // Adanto, the First Fort
- 1 Leyline of Hope
- 1 Liesa, Forgotten Archangel
- 1 Lyra Dawnbringer
- 1 Mangara, the Diplomat
- 1 Marauding Blight-Priest
- 1 Markov Purifier
- 1 Metropolis Reformer
- 1 Miren, the Moaning Well
- 1 Murderous Rider // Swift End
- 1 Neglected Manor
- 1 Nomad Stadium
- 1 Nuka-Cola Vending Machine
- 1 Nykthos Paragon
- 1 Obscura Storefront
- 1 Phyrexia's Core
- 1 Plaza of Harmony
- 1 Pristine Talisman
- 1 Radiant Fountain
- 1 Resplendent Angel
- 1 Restless Fortress
- 1 Rhox Faithmender
- 1 Righteous Valkyrie
- 1 Rune-Tail, Kitsune Ascendant // Rune-Tail's Essence
- 1 Sanguine Bond
- 1 Scoured Barrens
- 1 Seraph Sanctuary
- 1 Serra Ascendant
- 1 Shambling Vent
- 1 Sigarda's Splendor
- 1 Skyclave Cleric // Skyclave Basilica
- 1 Solitude
- 1 Springjack Pasture
- 1 Starlit Sanctum
- 1 Sword of Light and Shadow
- 1 The Celestus
- 1 The Gaffer
- 1 Tomb of the Spirit Dragon
- 1 Vault of the Archangel
- 1 Vein Ripper
- 1 Vito, Thorn of the Dusk Rose
- 1 Vizkopa Guildmage
- 1 Voice of the Blessed
- 1 Vona, Butcher of Magan
- 1 Well of Lost Dreams
- 1 Westvale Abbey // Ormendahl, Profane Prince
- 1 Witch's Clinic
- 1 Zof Consumption // Zof Bloodbog

## Revision 6, base 2: lifegain Commander, Karlov, bracket 3, any card

**The user:** Add two more cards that draw cards, and cut two of the weakest creatures.

The brief:

```json
{
  "changes": [
    "Add two more cards that draw cards",
    "Cut two of the weakest creatures"
  ],
  "remove": null,
  "keep": null,
  "max_mana_value": 0,
  "question": "",
  "declined": []
}
```

**The reply:** I removed 1 Bloodbond Vampire and 1 Cliffhaven Vampire. I added 1 Survival Cache and 1 Vampiric Rites.

Findings: bracket_prose_rules (INFO), curve_summary (INFO)

The deck:

- 1 Aclazotz, Deepest Betrayal // Temple of the Dead
- 1 Adventurer's Inn
- 1 Aerith Gainsborough
- 1 Aetherflux Reservoir
- 1 Ajani's Pridemate
- 1 Ajani, Strength of the Pride
- 1 Alseid of Life's Bounty
- 1 Altar of the Pantheon
- 1 Angel of Indemnity
- 1 Angel of Vitality
- 1 Archangel of Thune
- 1 Archivist of Oghma
- 1 Arguel's Blood Fast // Temple of Aclazotz
- 1 Astarion, the Decadent
- 1 Attended Healer
- 1 Ayli, Eternal Pilgrim
- 1 Azor's Gateway // Sanctum of the Sun
- 1 Beza, the Bounding Spring
- 1 Blighted Steppe
- 1 Bounty Board
- 1 Caduceus, Staff of Hermes
- 1 Carmen, Cruel Skymarcher
- 1 Cavalier of Night
- 1 Cleric Class
- 1 Cleric of Life's Bond
- 1 Consuming Corruption
- 1 Cosmos Elixir
- 1 Crypt Ghast
- 1 Dawn of Hope
- 1 Defiant Bloodlord
- 1 Diamond Valley
- 1 Divinity of Pride
- 1 Elenda, Saint of Dusk
- 1 Enduring Innocence
- 1 Exemplar of Light
- 1 Faith's Shield
- 1 Fisk Tower
- 1 Foot Headquarters
- 1 Fumigate
- 1 Glasswing Grace // Age-Graced Chapel
- 1 Glimmerpost
- 1 Heliod, Sun-Crowned
- 1 Hidden Hideout
- 1 Hierophant's Chalice
- 1 High Market
- 1 Hostile Hostel // Creeping Inn
- 1 Interplanar Beacon
- 1 Inventors' Fair
- 1 Kabira Crossroads
- 1 Kaya's Wrath
- 1 Legion's Landing // Adanto, the First Fort
- 1 Leyline of Hope
- 1 Liesa, Forgotten Archangel
- 1 Lyra Dawnbringer
- 1 Mangara, the Diplomat
- 1 Markov Purifier
- 1 Metropolis Reformer
- 1 Miren, the Moaning Well
- 1 Murderous Rider // Swift End
- 1 Neglected Manor
- 1 Nomad Stadium
- 1 Nuka-Cola Vending Machine
- 1 Nykthos Paragon
- 1 Obscura Storefront
- 1 Phyrexia's Core
- 1 Plaza of Harmony
- 1 Pristine Talisman
- 1 Radiant Fountain
- 1 Resplendent Angel
- 1 Restless Fortress
- 1 Rhox Faithmender
- 1 Righteous Valkyrie
- 1 Rune-Tail, Kitsune Ascendant // Rune-Tail's Essence
- 1 Sanguine Bond
- 1 Scoured Barrens
- 1 Seraph Sanctuary
- 1 Serra Ascendant
- 1 Shambling Vent
- 1 Sigarda's Splendor
- 1 Skyclave Cleric // Skyclave Basilica
- 1 Solitude
- 1 Springjack Pasture
- 1 Starlit Sanctum
- 1 Survival Cache
- 1 Sword of Light and Shadow
- 1 The Celestus
- 1 The Gaffer
- 1 Tomb of the Spirit Dragon
- 1 Vampiric Rites
- 1 Vault of the Archangel
- 1 Vein Ripper
- 1 Vito, Thorn of the Dusk Rose
- 1 Vizkopa Guildmage
- 1 Voice of the Blessed
- 1 Vona, Butcher of Magan
- 1 Well of Lost Dreams
- 1 Westvale Abbey // Ormendahl, Profane Prince
- 1 Witch's Clinic
- 1 Zof Consumption // Zof Bloodbog

## Revision 7, base 3: Commander from the Hobbit set, bracket 3

**The user:** Lower the mana curve. Nothing over 5 mana.

The brief:

```json
{
  "changes": [
    "Remove every nonland card with mana value 6 or more"
  ],
  "remove": null,
  "keep": null,
  "max_mana_value": 5,
  "question": "",
  "declined": []
}
```

**The reply:** I removed 1 Bag End Banquet, 1 Bilbo's Burglaring, 1 Boughside Wanderers, 1 Cantankerous Keepers, 1 Gigantic Big Bear, 1 Old Fat Spider, 1 The Lord of the Eagles, 1 Thranduil's Decree, and 1 Through the Forest Gate. I added 1 Attercop, 1 Bejeweled Warg, 1 Chief of the Wilds, 1 Down, Down to Goblin-town, 1 Elvish Mystic, 1 Guardian of the Halls, 1 Hithlain Knots, 1 Stern Scolding, and 1 Willow-Wind.

Findings: bracket_prose_rules (INFO), curve_summary (INFO)

The deck:

- 1 Along the Crooked Way
- 1 Arcane Signet
- 1 Arwen, Weaver of Hope
- 1 Attercop
- 1 Bejeweled Warg
- 1 Beorn the Fierce
- 1 Beorn's Hospitality
- 1 Bilbo Baggins, Burglar // Take a Glance
- 1 Bilbo's Deadly Slice
- 1 Bilbo's Ring
- 1 Bilbo, Fellow Conspirator
- 1 Bilbo, Luckwearer // Burglar's Plot
- 1 Bilbo, Thief in the Night
- 1 Bitter Downfall
- 1 Chief of the Wilds
- 1 Confusticate and Bebother
- 1 Delighted Halfling
- 1 Down in the Valley
- 1 Down, Down to Goblin-town
- 1 Dragon-Cursed Halls
- 1 Elven Chorus
- 1 Elven Passage
- 1 Elvenking's Halls
- 1 Elvish Mystic
- 1 Elvish Visionary
- 1 Enchanted River's Grasp
- 1 Fateful Discovery
- 1 Galadhrim Guide
- 1 Giant's Boulder
- 1 Gnashing of Teeth
- 1 Gollum the Abandoned
- 1 Gollum, Silent Slinker // Meager Meal
- 1 Great Fierce Bee
- 1 Guardian of the Halls
- 1 Hithlain Knots
- 1 Hobbit Hole
- 1 Key to the Side-Door
- 1 Languish
- 1 Large Bear
- 1 Little Bear
- 1 Minas Morgul, Dark Fortress
- 1 Mirkwood
- 1 Mithril Coat
- 1 Mox Amber
- 1 Nasty Little Rabbit
- 1 Night's Whisper
- 1 Old Fat Spider Can't See Me
- 1 Old Thrush
- 1 Orcish Bowmasters
- 1 Ordinary Bear
- 1 Part in Friendship
- 1 Quarrel
- 1 Raise the Palisade
- 1 Riddles in the Dark
- 1 Rivendell
- 1 Rogue's Passage
- 1 Roll-Roll-Roll-Roll
- 1 Stern Scolding
- 1 Sting, Bilbo's Sword
- 1 Stir Up Trouble
- 1 Supper for Spiders
- 1 The Black Gate
- 1 The Notary Hobbits
- 1 The One Ring
- 1 The Sackville-Bagginses
- 1 The Shire
- 1 Thrór's Map
- 1 Treasure Vault
- 1 Uncover the Moon-Letters
- 1 Uneasy Partings
- 1 Wayfarer's Bauble
- 1 Well-Worn Spatula
- 1 Willow-Wind
- 1 Wood Elves
- 2 Swamp
- 20 Forest
- 3 Island

## Revision 8, base 3: Commander from the Hobbit set, bracket 3

**The user:** Take out The Lord of the Eagles.

The brief:

```json
{
  "changes": [],
  "remove": [
    "The Lord of the Eagles"
  ],
  "keep": null,
  "max_mana_value": 0,
  "question": "",
  "declined": []
}
```

**The reply:** I removed 1 The Lord of the Eagles. I added 1 Gollum, Riddle Master.

Findings: bracket_prose_rules (INFO), curve_summary (INFO)

The deck:

- 1 Along the Crooked Way
- 1 Arcane Signet
- 1 Arwen, Weaver of Hope
- 1 Bag End Banquet
- 1 Beorn the Fierce
- 1 Beorn's Hospitality
- 1 Bilbo Baggins, Burglar // Take a Glance
- 1 Bilbo's Burglaring
- 1 Bilbo's Deadly Slice
- 1 Bilbo's Ring
- 1 Bilbo, Fellow Conspirator
- 1 Bilbo, Luckwearer // Burglar's Plot
- 1 Bilbo, Thief in the Night
- 1 Bitter Downfall
- 1 Boughside Wanderers
- 1 Cantankerous Keepers
- 1 Confusticate and Bebother
- 1 Delighted Halfling
- 1 Down in the Valley
- 1 Dragon-Cursed Halls
- 1 Elven Chorus
- 1 Elven Passage
- 1 Elvenking's Halls
- 1 Elvish Visionary
- 1 Enchanted River's Grasp
- 1 Fateful Discovery
- 1 Galadhrim Guide
- 1 Giant's Boulder
- 1 Gigantic Big Bear
- 1 Gnashing of Teeth
- 1 Gollum the Abandoned
- 1 Gollum, Riddle Master
- 1 Gollum, Silent Slinker // Meager Meal
- 1 Great Fierce Bee
- 1 Hobbit Hole
- 1 Key to the Side-Door
- 1 Languish
- 1 Large Bear
- 1 Little Bear
- 1 Minas Morgul, Dark Fortress
- 1 Mirkwood
- 1 Mithril Coat
- 1 Mox Amber
- 1 Nasty Little Rabbit
- 1 Night's Whisper
- 1 Old Fat Spider
- 1 Old Fat Spider Can't See Me
- 1 Old Thrush
- 1 Orcish Bowmasters
- 1 Ordinary Bear
- 1 Part in Friendship
- 1 Quarrel
- 1 Raise the Palisade
- 1 Riddles in the Dark
- 1 Rivendell
- 1 Rogue's Passage
- 1 Roll-Roll-Roll-Roll
- 1 Sting, Bilbo's Sword
- 1 Stir Up Trouble
- 1 Supper for Spiders
- 1 The Black Gate
- 1 The Notary Hobbits
- 1 The One Ring
- 1 The Sackville-Bagginses
- 1 The Shire
- 1 Thranduil's Decree
- 1 Through the Forest Gate
- 1 Thrór's Map
- 1 Treasure Vault
- 1 Uncover the Moon-Letters
- 1 Uneasy Partings
- 1 Wayfarer's Bauble
- 1 Well-Worn Spatula
- 1 Wood Elves
- 2 Swamp
- 20 Forest
- 3 Island
