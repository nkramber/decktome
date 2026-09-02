# PR-12B revise gate

Run date: 2026-09-02. Snapshot: 2026-09-01.

Verdict: FAIL. 9 of 10 turns met their bar. The bars: an unclear request gets a question or a decline with a reason, an answer to that question gets a rebuild, a clear request gets a revised deck with no block finding, the deck holds every cap, every removal, and every land swap the message names, and it keeps at least 80 percent of the untouched cards.

Cost: 24 calls, $0.6688. Time: 9m31s.

| Base | Revision | Message | Outcome | Kept | Result |
|---|---|---|---|---|---|
| 1 | 1 | Replace some lands with better options if possible. Also tune the mana curve lower - no 6 or 7 mana cards needed | no build |  | FAIL: an unclear request got no question |
| 1 | 2 | Take out Angel of Serenity. | 60 cards | 100% | PASS, kept 100% |
| 1 | 3 | Better lands please. | question |  | PASS, no build |
| 1 | 3 | (the answer) Q: For this one-color deck, do you want faster mana, utility lands, or lands that let you add other colors? A: Utility lands. | 60 cards | 100% | PASS, kept 100% |
| 2 | 4 | No nonland cards over 4 mana. | 99 cards | 97% | PASS, kept 97% |
| 2 | 5 | Swap Defiant Bloodlord for something cheaper that does the same job. | 99 cards | 100% | PASS, kept 100% |
| 2 | 6 | Add two more cards that draw cards, and cut two of the weakest creatures. | 99 cards | 97% | PASS, kept 97% |
| 2 | 9 | Add better lands instead of the basics. I mean dual lands that enter untapped where you can, and a couple of utility lands. | 99 cards | 100% | PASS, kept 100% |
| 3 | 7 | Lower the mana curve. Nothing over 5 mana. | 99 cards | 94% | PASS, kept 94% |
| 3 | 8 | Take out Gigantic Big Bear. | 99 cards | 100% | PASS, kept 100% |

## Revision 1, base 1: Modern Angels, casual, mono-white, any card

**The user:** Replace some lands with better options if possible. Also tune the mana curve lower - no 6 or 7 mana cards needed

The brief:

```json
{
  "changes": [
    "Replace 3 Plains with white utility lands that support a casual mono-white +1/+1 counter Angel deck.",
    "Remove every nonland card with mana value 6 or more."
  ],
  "remove": null,
  "keep": null,
  "max_mana_value": 5,
  "swap_basics": 3,
  "land_kinds": "Mono-white utility lands that provide useful effects while reliably producing white mana",
  "question": "",
  "declined": []
}
```

**The reply:** I read your message, and I found no change to make to the deck. Tell me what you want changed, and I will do it.


Failures: an unclear request got no question

## Revision 2, base 1: Modern Angels, casual, mono-white, any card

**The user:** Take out Angel of Serenity.

The brief:

```json
{
  "changes": [
    "Remove Angel of Serenity."
  ],
  "remove": [
    "Angel of Serenity"
  ],
  "keep": null,
  "max_mana_value": 0,
  "swap_basics": 0,
  "land_kinds": "",
  "question": "",
  "declined": []
}
```

**The reply:** I removed 1 Angel of Serenity. I changed the count of Plains: 20 to 21.

Findings: curve_summary (INFO), basics_added (INFO)

The deck:

- 1 Abzan Falconer
- 1 Admonition Angel
- 1 Aetherstorm Roc
- 1 Angel of Condemnation
- 1 Angel of Invention
- 1 Angel of Sanctions
- 1 Angel of the Ruins
- 1 Arcane Signet
- 1 Archangel of Thune
- 1 Aurelia's Vindicator
- 1 Basri's Lieutenant
- 1 Bruna, the Fading Light
- 1 Cathars' Crusade
- 1 Elite Scaleguard
- 1 Exemplar of Light
- 1 Felidar Retreat
- 1 Giada, Font of Hope
- 1 Hall of Oracles
- 1 Herald of War
- 1 Herald of the Sun
- 1 High Sentinels of Arashin
- 1 Inspiring Overseer
- 1 Karn's Bastion
- 1 Light of the Legion
- 1 Luminarch Aspirant
- 1 Lyra Dawnbringer
- 1 Mikaeus, the Lunarch
- 1 Paragon of Modernity
- 1 Patron of the Valiant
- 1 Righteous Valkyrie
- 1 Sanctuary Warden
- 1 Secluded Courtyard
- 1 Seraph Sanctuary
- 1 Serra Redeemer
- 1 Sigarda's Summons
- 1 Sparring Regimen
- 1 Winged Shepherd
- 1 Wojek Investigator
- 1 Youthful Valkyrie
- 21 Plains

## Revision 3, base 1: Modern Angels, casual, mono-white, any card

**The user:** Better lands please.

The brief:

```json
{
  "changes": [],
  "remove": null,
  "keep": null,
  "max_mana_value": 0,
  "swap_basics": 0,
  "land_kinds": "",
  "question": "For this one-color deck, do you want faster mana, utility lands, or lands that let you add other colors?",
  "declined": []
}
```

**The reply:** Question: For this one-color deck, do you want faster mana, utility lands, or lands that let you add other colors?


## Revision 3, the answer, base 1: Modern Angels, casual, mono-white, any card

**The user:** Q: For this one-color deck, do you want faster mana, utility lands, or lands that let you add other colors? A: Utility lands.

The brief:

```json
{
  "changes": [
    "Replace a few Plains with utility lands that support the deck without adding colors."
  ],
  "remove": null,
  "keep": null,
  "max_mana_value": 0,
  "swap_basics": 3,
  "land_kinds": "Utility lands for a mono-white +1/+1 counter Angel deck",
  "question": "",
  "declined": []
}
```

**The reply:** I added 1 Drannith Ruins, 1 Ruins of Oran-Rief, and 1 Tyrite Sanctum. I changed the count of Plains: 20 to 17.

Findings: curve_summary (INFO)

The deck:

- 1 Abzan Falconer
- 1 Admonition Angel
- 1 Aetherstorm Roc
- 1 Angel of Condemnation
- 1 Angel of Invention
- 1 Angel of Sanctions
- 1 Angel of Serenity
- 1 Angel of the Ruins
- 1 Arcane Signet
- 1 Archangel of Thune
- 1 Aurelia's Vindicator
- 1 Basri's Lieutenant
- 1 Bruna, the Fading Light
- 1 Cathars' Crusade
- 1 Drannith Ruins
- 1 Elite Scaleguard
- 1 Exemplar of Light
- 1 Felidar Retreat
- 1 Giada, Font of Hope
- 1 Hall of Oracles
- 1 Herald of War
- 1 Herald of the Sun
- 1 High Sentinels of Arashin
- 1 Inspiring Overseer
- 1 Karn's Bastion
- 1 Light of the Legion
- 1 Luminarch Aspirant
- 1 Lyra Dawnbringer
- 1 Mikaeus, the Lunarch
- 1 Paragon of Modernity
- 1 Patron of the Valiant
- 1 Righteous Valkyrie
- 1 Ruins of Oran-Rief
- 1 Sanctuary Warden
- 1 Secluded Courtyard
- 1 Seraph Sanctuary
- 1 Serra Redeemer
- 1 Sigarda's Summons
- 1 Sparring Regimen
- 1 Tyrite Sanctum
- 1 Winged Shepherd
- 1 Wojek Investigator
- 1 Youthful Valkyrie
- 17 Plains

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
  "swap_basics": 0,
  "land_kinds": "",
  "question": "",
  "declined": []
}
```

**The reply:** I removed 1 Archangel of Thune, 1 Astarion, the Decadent, 1 Blood Baron of Vizkopa, 1 Cavalier of Night, 1 Celestine, the Living Saint, 1 Convalescent Care, 1 Defiant Bloodlord, 1 Divinity of Pride, 1 Fumigate, 1 Nykthos Paragon, 1 Sanguine Bond, 1 Sigarda's Splendor, 1 Solitude, 1 Valkyrie Harbinger, 1 Vein Ripper, 1 Vona, Butcher of Magan, and 1 Wurmcoil Engine. I added 1 Ajani Goldmane, 1 Archon of Sun's Grace, 1 Bloodbond Vampire, 1 Bronzebeak Foragers, 1 Elenda, Saint of Dusk, 1 Elenda, the Dusk Rose, 1 Enduring Tenacity, 1 Exemplar of Light, 1 Henrika Domnathi // Henrika, Infernal Seer, 1 Qala, Ajani's Pridemate, 1 Sorin, Imperious Bloodlord, 1 Star Charter, 1 Summon: Ixion, 1 Survival Cache, 1 Twinblade Paladin, and 1 White Sun's Twilight. I changed the count of Plains: 20 to 21.

Findings: bracket_prose_rules (INFO), curve_summary (INFO), basics_added (INFO)

Repair turn: 2 missed names.

The deck:

- 1 Aetherflux Reservoir
- 1 Ajani Goldmane
- 1 Ajani's Pridemate
- 1 Ajani, Strength of the Pride
- 1 Alseid of Life's Bounty
- 1 Altar of the Pantheon
- 1 Angel of Vitality
- 1 Archivist of Oghma
- 1 Archon of Sun's Grace
- 1 Attended Healer
- 1 Ayara, First of Locthwain
- 1 Ayli, Eternal Pilgrim
- 1 Bloodbond Vampire
- 1 Bloodthirsty Aerialist
- 1 Bounty Board
- 1 Bronzebeak Foragers
- 1 Caves of Koilos
- 1 Cleric Class
- 1 Cliffhaven Vampire
- 1 Command Tower
- 1 Cosmos Elixir
- 1 Courageous Resolve
- 1 Crypt Ghast
- 1 Dawn of Hope
- 1 Druidic Satchel
- 1 Elenda, Saint of Dusk
- 1 Elenda, the Dusk Rose
- 1 Enduring Innocence
- 1 Enduring Tenacity
- 1 Exemplar of Light
- 1 Faith's Shield
- 1 Godless Shrine
- 1 Heliod, Sun-Crowned
- 1 Henrika Domnathi // Henrika, Infernal Seer
- 1 Hierophant's Chalice
- 1 Isolated Chapel
- 1 Kaya's Wrath
- 1 Light of Promise
- 1 Mangara, the Diplomat
- 1 Markov Purifier
- 1 Metropolis Reformer
- 1 Murderous Rider // Swift End
- 1 Nightmare's Thirst
- 1 Nuka-Cola Vending Machine
- 1 Potioner's Trove
- 1 Pristine Talisman
- 1 Qala, Ajani's Pridemate
- 1 Resplendent Angel
- 1 Restoration Magic
- 1 Rhox Faithmender
- 1 Righteous Valkyrie
- 1 Scoured Barrens
- 1 Serra Ascendant
- 1 Shattered Sanctum
- 1 Sorin, Imperious Bloodlord
- 1 Speaker of the Heavens
- 1 Star Charter
- 1 Summon: Ixion
- 1 Survival Cache
- 1 Temple of Silence
- 1 The Celestus
- 1 Treasure Chest
- 1 Twinblade Paladin
- 1 Vault of Champions
- 1 Vito, Thorn of the Dusk Rose
- 1 Vizkopa Guildmage
- 1 Voice of the Blessed
- 1 Well of Lost Dreams
- 1 Werefox Bodyguard
- 1 White Sun's Twilight
- 21 Plains
- 8 Swamp

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
  "swap_basics": 0,
  "land_kinds": "",
  "question": "",
  "declined": []
}
```

**The reply:** I removed 1 Defiant Bloodlord. I added 1 Epicure of Blood.

Findings: bracket_prose_rules (INFO), curve_summary (INFO)

The deck:

- 1 Aetherflux Reservoir
- 1 Ajani's Pridemate
- 1 Ajani, Strength of the Pride
- 1 Alseid of Life's Bounty
- 1 Altar of the Pantheon
- 1 Angel of Vitality
- 1 Archangel of Thune
- 1 Archivist of Oghma
- 1 Astarion, the Decadent
- 1 Attended Healer
- 1 Ayara, First of Locthwain
- 1 Ayli, Eternal Pilgrim
- 1 Blood Baron of Vizkopa
- 1 Bloodthirsty Aerialist
- 1 Bounty Board
- 1 Cavalier of Night
- 1 Caves of Koilos
- 1 Celestine, the Living Saint
- 1 Cleric Class
- 1 Cliffhaven Vampire
- 1 Command Tower
- 1 Convalescent Care
- 1 Cosmos Elixir
- 1 Courageous Resolve
- 1 Crypt Ghast
- 1 Dawn of Hope
- 1 Divinity of Pride
- 1 Druidic Satchel
- 1 Enduring Innocence
- 1 Epicure of Blood
- 1 Faith's Shield
- 1 Fumigate
- 1 Godless Shrine
- 1 Heliod, Sun-Crowned
- 1 Hierophant's Chalice
- 1 Isolated Chapel
- 1 Kaya's Wrath
- 1 Light of Promise
- 1 Mangara, the Diplomat
- 1 Markov Purifier
- 1 Metropolis Reformer
- 1 Murderous Rider // Swift End
- 1 Nightmare's Thirst
- 1 Nuka-Cola Vending Machine
- 1 Nykthos Paragon
- 1 Potioner's Trove
- 1 Pristine Talisman
- 1 Resplendent Angel
- 1 Restoration Magic
- 1 Rhox Faithmender
- 1 Righteous Valkyrie
- 1 Sanguine Bond
- 1 Scoured Barrens
- 1 Serra Ascendant
- 1 Shattered Sanctum
- 1 Sigarda's Splendor
- 1 Solitude
- 1 Speaker of the Heavens
- 1 Temple of Silence
- 1 The Celestus
- 1 Treasure Chest
- 1 Valkyrie Harbinger
- 1 Vault of Champions
- 1 Vein Ripper
- 1 Vito, Thorn of the Dusk Rose
- 1 Vizkopa Guildmage
- 1 Voice of the Blessed
- 1 Vona, Butcher of Magan
- 1 Well of Lost Dreams
- 1 Werefox Bodyguard
- 1 Wurmcoil Engine
- 20 Plains
- 8 Swamp

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
  "swap_basics": 0,
  "land_kinds": "",
  "question": "",
  "declined": []
}
```

**The reply:** I removed 1 Attended Healer and 1 Cliffhaven Vampire. I added 1 Survival Cache and 1 The Gaffer.

Findings: bracket_prose_rules (INFO), curve_summary (INFO)

The deck:

- 1 Aetherflux Reservoir
- 1 Ajani's Pridemate
- 1 Ajani, Strength of the Pride
- 1 Alseid of Life's Bounty
- 1 Altar of the Pantheon
- 1 Angel of Vitality
- 1 Archangel of Thune
- 1 Archivist of Oghma
- 1 Astarion, the Decadent
- 1 Ayara, First of Locthwain
- 1 Ayli, Eternal Pilgrim
- 1 Blood Baron of Vizkopa
- 1 Bloodthirsty Aerialist
- 1 Bounty Board
- 1 Cavalier of Night
- 1 Caves of Koilos
- 1 Celestine, the Living Saint
- 1 Cleric Class
- 1 Command Tower
- 1 Convalescent Care
- 1 Cosmos Elixir
- 1 Courageous Resolve
- 1 Crypt Ghast
- 1 Dawn of Hope
- 1 Defiant Bloodlord
- 1 Divinity of Pride
- 1 Druidic Satchel
- 1 Enduring Innocence
- 1 Faith's Shield
- 1 Fumigate
- 1 Godless Shrine
- 1 Heliod, Sun-Crowned
- 1 Hierophant's Chalice
- 1 Isolated Chapel
- 1 Kaya's Wrath
- 1 Light of Promise
- 1 Mangara, the Diplomat
- 1 Markov Purifier
- 1 Metropolis Reformer
- 1 Murderous Rider // Swift End
- 1 Nightmare's Thirst
- 1 Nuka-Cola Vending Machine
- 1 Nykthos Paragon
- 1 Potioner's Trove
- 1 Pristine Talisman
- 1 Resplendent Angel
- 1 Restoration Magic
- 1 Rhox Faithmender
- 1 Righteous Valkyrie
- 1 Sanguine Bond
- 1 Scoured Barrens
- 1 Serra Ascendant
- 1 Shattered Sanctum
- 1 Sigarda's Splendor
- 1 Solitude
- 1 Speaker of the Heavens
- 1 Survival Cache
- 1 Temple of Silence
- 1 The Celestus
- 1 The Gaffer
- 1 Treasure Chest
- 1 Valkyrie Harbinger
- 1 Vault of Champions
- 1 Vein Ripper
- 1 Vito, Thorn of the Dusk Rose
- 1 Vizkopa Guildmage
- 1 Voice of the Blessed
- 1 Vona, Butcher of Magan
- 1 Well of Lost Dreams
- 1 Werefox Bodyguard
- 1 Wurmcoil Engine
- 20 Plains
- 8 Swamp

## Revision 9, base 2: lifegain Commander, Karlov, bracket 3, any card

**The user:** Add better lands instead of the basics. I mean dual lands that enter untapped where you can, and a couple of utility lands.

The brief:

```json
{
  "changes": [
    "Replace basic lands with untapped white-black dual lands where possible and add a couple of utility lands."
  ],
  "remove": null,
  "keep": null,
  "max_mana_value": 0,
  "swap_basics": 14,
  "land_kinds": "Untapped white-black dual lands and a couple of utility lands",
  "question": "",
  "declined": []
}
```

**The reply:** I added 1 Aclazotz, Deepest Betrayal // Temple of the Dead, 1 Cavern of Souls, 1 City of Brass, 1 Exotic Orchard, 1 Gemstone Caverns, 1 High Market, 1 Mana Confluence, 1 Neglected Manor, 1 Path of Ancestry, 1 Reflecting Pool, 1 Restless Fortress, 1 Secluded Courtyard, 1 Tomb of the Spirit Dragon, and 1 Unclaimed Territory. I changed the count of Plains: 20 to 10 and Swamp: 8 to 4.

Findings: bracket_prose_rules (INFO), curve_summary (INFO)

The deck:

- 1 Aclazotz, Deepest Betrayal // Temple of the Dead
- 1 Aetherflux Reservoir
- 1 Ajani's Pridemate
- 1 Ajani, Strength of the Pride
- 1 Alseid of Life's Bounty
- 1 Altar of the Pantheon
- 1 Angel of Vitality
- 1 Archangel of Thune
- 1 Archivist of Oghma
- 1 Astarion, the Decadent
- 1 Attended Healer
- 1 Ayara, First of Locthwain
- 1 Ayli, Eternal Pilgrim
- 1 Blood Baron of Vizkopa
- 1 Bloodthirsty Aerialist
- 1 Bounty Board
- 1 Cavalier of Night
- 1 Cavern of Souls
- 1 Caves of Koilos
- 1 Celestine, the Living Saint
- 1 City of Brass
- 1 Cleric Class
- 1 Cliffhaven Vampire
- 1 Command Tower
- 1 Convalescent Care
- 1 Cosmos Elixir
- 1 Courageous Resolve
- 1 Crypt Ghast
- 1 Dawn of Hope
- 1 Defiant Bloodlord
- 1 Divinity of Pride
- 1 Druidic Satchel
- 1 Enduring Innocence
- 1 Exotic Orchard
- 1 Faith's Shield
- 1 Fumigate
- 1 Gemstone Caverns
- 1 Godless Shrine
- 1 Heliod, Sun-Crowned
- 1 Hierophant's Chalice
- 1 High Market
- 1 Isolated Chapel
- 1 Kaya's Wrath
- 1 Light of Promise
- 1 Mana Confluence
- 1 Mangara, the Diplomat
- 1 Markov Purifier
- 1 Metropolis Reformer
- 1 Murderous Rider // Swift End
- 1 Neglected Manor
- 1 Nightmare's Thirst
- 1 Nuka-Cola Vending Machine
- 1 Nykthos Paragon
- 1 Path of Ancestry
- 1 Potioner's Trove
- 1 Pristine Talisman
- 1 Reflecting Pool
- 1 Resplendent Angel
- 1 Restless Fortress
- 1 Restoration Magic
- 1 Rhox Faithmender
- 1 Righteous Valkyrie
- 1 Sanguine Bond
- 1 Scoured Barrens
- 1 Secluded Courtyard
- 1 Serra Ascendant
- 1 Shattered Sanctum
- 1 Sigarda's Splendor
- 1 Solitude
- 1 Speaker of the Heavens
- 1 Temple of Silence
- 1 The Celestus
- 1 Tomb of the Spirit Dragon
- 1 Treasure Chest
- 1 Unclaimed Territory
- 1 Valkyrie Harbinger
- 1 Vault of Champions
- 1 Vein Ripper
- 1 Vito, Thorn of the Dusk Rose
- 1 Vizkopa Guildmage
- 1 Voice of the Blessed
- 1 Vona, Butcher of Magan
- 1 Well of Lost Dreams
- 1 Werefox Bodyguard
- 1 Wurmcoil Engine
- 10 Plains
- 4 Swamp

## Revision 7, base 3: Commander from the Hobbit set, bracket 3

**The user:** Lower the mana curve. Nothing over 5 mana.

The brief:

```json
{
  "changes": [
    "Remove every nonland card with mana value 6 or more."
  ],
  "remove": null,
  "keep": null,
  "max_mana_value": 5,
  "swap_basics": 0,
  "land_kinds": "",
  "question": "",
  "declined": []
}
```

**The reply:** I removed 1 Attercop, 1 Bag End Banquet, 1 Beorn the Fierce, 1 Boughside Wanderers, 1 Cantankerous Keepers, 1 Celeborn the Wise, 1 Galion, Elvenking's Butler, 1 Gigantic Big Bear, and 1 Mirkwood Elk. I added 1 Arwen, Weaver of Hope, 1 Galadhrim Guide, 1 Little Bear, 1 Mirkwood Nurturer, 1 Necklace of Girion, and 1 Ordinary Bear. I changed the count of Forest: 13 to 14, Island: 7 to 8, and Swamp: 5 to 6.

Findings: bracket_prose_rules (INFO), land_count (WARN), curve_summary (INFO), basics_added (INFO)

Repair turn: 2 missed names.

The deck:

- 1 Along the Crooked Way
- 1 Arcane Signet
- 1 Arwen, Weaver of Hope
- 1 Beorn's Hospitality
- 1 Beorn, Reluctant Host // Till and Tend
- 1 Bilbo Baggins, Burglar // Take a Glance
- 1 Bilbo's Deadly Slice
- 1 Bilbo's Ring
- 1 Bilbo, Fellow Conspirator
- 1 Bilbo, Luckwearer // Burglar's Plot
- 1 Bilbo, Thief in the Night
- 1 Bitter Downfall
- 1 Confusticate and Bebother
- 1 Delighted Halfling
- 1 Down in the Valley
- 1 Dragon-Cursed Halls
- 1 Elrond, Moon-Reader
- 1 Elven Chorus
- 1 Elven Passage
- 1 Elvenking's Halls
- 1 Elvish Archdruid
- 1 Elvish Mystic
- 1 Elvish Visionary
- 1 Enchanted River's Grasp
- 1 Fateful Discovery
- 1 Galadhrim Guide
- 1 Gnashing of Teeth
- 1 Gollum, Riddle Master
- 1 Gollum, Silent Slinker // Meager Meal
- 1 Guardian of the Halls
- 1 Hithlain Knots
- 1 Hobbit Hole
- 1 Inside Information
- 1 Key to the Side-Door
- 1 Languish
- 1 Large Bear
- 1 Little Bear
- 1 Merciless Executioner
- 1 Minas Morgul, Dark Fortress
- 1 Mirkwood
- 1 Mirkwood Nurturer
- 1 Mirkwood Pathmaker
- 1 Mithril Coat
- 1 Mox Amber
- 1 My Precious // Allure of Power
- 1 Nasty Little Rabbit
- 1 Necklace of Girion
- 1 Night's Whisper
- 1 Nimrodel Watcher
- 1 Orcish Bowmasters
- 1 Ordinary Bear
- 1 Palantír of Orthanc
- 1 Part in Friendship
- 1 Quarrel
- 1 Raise the Palisade
- 1 Riddles in the Dark
- 1 Rivendell
- 1 Rogue's Passage
- 1 Roll-Roll-Roll-Roll
- 1 Stern Scolding
- 1 Supper for Spiders
- 1 The Black Arrow
- 1 The Black Gate
- 1 The Notary Hobbits
- 1 The Sackville-Bagginses
- 1 The Shire
- 1 Treasure Vault
- 1 Uncover the Moon-Letters
- 1 Uneasy Partings
- 1 Wayfarer's Bauble
- 1 Well-Worn Spatula
- 14 Forest
- 6 Swamp
- 8 Island

## Revision 8, base 3: Commander from the Hobbit set, bracket 3

**The user:** Take out Gigantic Big Bear.

The brief:

```json
{
  "changes": [
    "Remove Gigantic Big Bear."
  ],
  "remove": [
    "Gigantic Big Bear"
  ],
  "keep": null,
  "max_mana_value": 0,
  "swap_basics": 0,
  "land_kinds": "",
  "question": "",
  "declined": []
}
```

**The reply:** I removed 1 Gigantic Big Bear. I added 1 Gollum the Abandoned.

Findings: bracket_prose_rules (INFO), curve_summary (INFO)

The deck:

- 1 Along the Crooked Way
- 1 Arcane Signet
- 1 Attercop
- 1 Bag End Banquet
- 1 Beorn the Fierce
- 1 Beorn's Hospitality
- 1 Beorn, Reluctant Host // Till and Tend
- 1 Bilbo Baggins, Burglar // Take a Glance
- 1 Bilbo's Deadly Slice
- 1 Bilbo's Ring
- 1 Bilbo, Fellow Conspirator
- 1 Bilbo, Luckwearer // Burglar's Plot
- 1 Bilbo, Thief in the Night
- 1 Bitter Downfall
- 1 Boughside Wanderers
- 1 Cantankerous Keepers
- 1 Celeborn the Wise
- 1 Confusticate and Bebother
- 1 Delighted Halfling
- 1 Down in the Valley
- 1 Dragon-Cursed Halls
- 1 Elrond, Moon-Reader
- 1 Elven Chorus
- 1 Elven Passage
- 1 Elvenking's Halls
- 1 Elvish Archdruid
- 1 Elvish Mystic
- 1 Elvish Visionary
- 1 Enchanted River's Grasp
- 1 Fateful Discovery
- 1 Galion, Elvenking's Butler
- 1 Gnashing of Teeth
- 1 Gollum the Abandoned
- 1 Gollum, Riddle Master
- 1 Gollum, Silent Slinker // Meager Meal
- 1 Guardian of the Halls
- 1 Hithlain Knots
- 1 Hobbit Hole
- 1 Inside Information
- 1 Key to the Side-Door
- 1 Languish
- 1 Large Bear
- 1 Merciless Executioner
- 1 Minas Morgul, Dark Fortress
- 1 Mirkwood
- 1 Mirkwood Elk
- 1 Mirkwood Pathmaker
- 1 Mithril Coat
- 1 Mox Amber
- 1 My Precious // Allure of Power
- 1 Nasty Little Rabbit
- 1 Night's Whisper
- 1 Nimrodel Watcher
- 1 Orcish Bowmasters
- 1 Palantír of Orthanc
- 1 Part in Friendship
- 1 Quarrel
- 1 Raise the Palisade
- 1 Riddles in the Dark
- 1 Rivendell
- 1 Rogue's Passage
- 1 Roll-Roll-Roll-Roll
- 1 Stern Scolding
- 1 Supper for Spiders
- 1 The Black Arrow
- 1 The Black Gate
- 1 The Notary Hobbits
- 1 The Sackville-Bagginses
- 1 The Shire
- 1 Treasure Vault
- 1 Uncover the Moon-Letters
- 1 Uneasy Partings
- 1 Wayfarer's Bauble
- 1 Well-Worn Spatula
- 13 Forest
- 5 Swamp
- 7 Island
