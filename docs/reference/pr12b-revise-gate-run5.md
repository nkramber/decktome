# PR-12B revise gate

Run date: 2026-09-02. Snapshot: 2026-09-01.

Verdict: FAIL. 10 of 11 turns met their bar. The bars: an unclear request gets a question or a decline with a reason, an answer to that question gets a rebuild, a clear request gets a revised deck with no block finding, the deck holds every cap, every removal, and every land swap the message names, and it keeps at least 80 percent of the untouched cards.

Cost: 27 calls, $0.8115. Time: 9m26s.

| Base | Revision | Message | Outcome | Kept | Result |
|---|---|---|---|---|---|
| 1 | 1 | Replace some lands with better options if possible. Also tune the mana curve lower - no 6 or 7 mana cards needed | question |  | PASS, no build |
| 1 | 1 | (the answer) Q: For the Plains upgrades, do you want faster mana, utility lands, or lands that add other colors? A: A mix. | 60 cards | 100% | PASS, kept 100% |
| 1 | 2 | Take out Angel of Deliverance. | 60 cards | 100% | PASS, kept 100% |
| 1 | 3 | Better lands please. | question |  | PASS, no build |
| 1 | 3 | (the answer) Q: This is a one-color deck. By “better lands,” do you want faster mana, utility lands with useful effects, or lands that add other colors? A: Utility lands. | 60 cards | 100% | PASS, kept 100% |
| 2 | 4 | No nonland cards over 4 mana. | 99 cards | 100% | PASS, kept 100% |
| 2 | 5 | Swap Defiant Bloodlord for something cheaper that does the same job. | 99 cards | 100% | PASS, kept 100% |
| 2 | 6 | Add two more cards that draw cards, and cut two of the weakest creatures. | 99 cards | 97% | PASS, kept 97% |
| 2 | 9 | Add better lands instead of the basics. I mean dual lands that enter untapped where you can, and a couple of utility lands. | 99 cards | 100% | FAIL: block findings: revision_lands_kept; the deck holds 1 more nonbasic lands, and the brief asked for 12 |
| 3 | 7 | Lower the mana curve. Nothing over 5 mana. | 99 cards | 100% | PASS, kept 100% |
| 3 | 8 | Take out Colossal Whale. | 99 cards | 100% | PASS, kept 100% |

## Revision 1, base 1: Modern Angels, casual, mono-white, any card

**The user:** Replace some lands with better options if possible. Also tune the mana curve lower - no 6 or 7 mana cards needed

The brief:

```json
{
  "changes": [
    "Remove every nonland card with mana value 6 or more to lower the mana curve."
  ],
  "remove": null,
  "keep": null,
  "max_mana_value": 5,
  "swap_basics": 0,
  "land_kinds": "",
  "question": "For the Plains upgrades, do you want faster mana, utility lands, or lands that add other colors?",
  "declined": []
}
```

**The reply:** Question: For the Plains upgrades, do you want faster mana, utility lands, or lands that add other colors?


## Revision 1, the answer, base 1: Modern Angels, casual, mono-white, any card

**The user:** Q: For the Plains upgrades, do you want faster mana, utility lands, or lands that add other colors? A: A mix.

The brief:

```json
{
  "changes": [
    "Replace some Plains with a mix of faster-mana lands, utility lands, and lands that can add other colors.",
    "Remove every nonland card with mana value 6 or more."
  ],
  "remove": null,
  "keep": null,
  "max_mana_value": 5,
  "swap_basics": 4,
  "land_kinds": "A mix of faster-mana lands, utility lands, and lands that add other colors",
  "question": "",
  "declined": []
}
```

**The reply:** I removed 1 Admonition Angel, 1 Angel of Deliverance, 1 Angel of Serenity, 1 Angel of the God-Pharaoh, 1 Angel of the Ruins, 1 Bruna, the Fading Light, 1 Herald of the Sun, 1 Light of the Legion, 1 Nykthos Paragon, 1 Sanctuary Warden, and 1 Winged Shepherd. I added 1 Angel of Destiny, 1 Arid Mesa, 1 Aurelia's Vindicator, 1 Cavern of Souls, 1 Constable of the Realm, 1 Flooded Strand, 1 Karn's Bastion, 1 Mana Confluence, and 2 Dusk Legion Duelist. I changed the count of Angel of Condemnation: 1 to 2, Angel of Sanctions: 1 to 2, Inspiring Overseer: 1 to 2, Patron of the Valiant: 1 to 2, and Plains: 24 to 21.

Findings: curve_summary (INFO), basics_added (INFO)

The deck:

- 1 Anafenza, Kin-Tree Spirit
- 1 Angel of Destiny
- 1 Angel of Invention
- 1 Arcane Signet
- 1 Archangel of Thune
- 1 Arid Mesa
- 1 Aurelia's Vindicator
- 1 Basri's Lieutenant
- 1 Bishop of Wings
- 1 Cathars' Crusade
- 1 Cavern of Souls
- 1 Constable of the Realm
- 1 Dueling Coach
- 1 Exemplar of Light
- 1 Felidar Retreat
- 1 Flooded Strand
- 1 Giada, Font of Hope
- 1 Herald of War
- 1 High Sentinels of Arashin
- 1 Karn's Bastion
- 1 Luminarch Aspirant
- 1 Lyra Dawnbringer
- 1 Mana Confluence
- 1 Mikaeus, the Lunarch
- 1 Paragon of Modernity
- 1 Righteous Valkyrie
- 1 Serra Redeemer
- 1 Wojek Investigator
- 1 Youthful Valkyrie
- 2 Angel of Condemnation
- 2 Angel of Sanctions
- 2 Dusk Legion Duelist
- 2 Inspiring Overseer
- 2 Patron of the Valiant
- 21 Plains

## Revision 2, base 1: Modern Angels, casual, mono-white, any card

**The user:** Take out Angel of Deliverance.

The brief:

```json
{
  "changes": [
    "Remove Angel of Deliverance."
  ],
  "remove": [
    "Angel of Deliverance"
  ],
  "keep": null,
  "max_mana_value": 0,
  "swap_basics": 0,
  "land_kinds": "",
  "question": "",
  "declined": []
}
```

**The reply:** I removed 1 Angel of Deliverance. I added 1 Aurelia's Vindicator.

Findings: curve_summary (INFO)

The deck:

- 1 Admonition Angel
- 1 Anafenza, Kin-Tree Spirit
- 1 Angel of Condemnation
- 1 Angel of Invention
- 1 Angel of Sanctions
- 1 Angel of Serenity
- 1 Angel of the God-Pharaoh
- 1 Angel of the Ruins
- 1 Arcane Signet
- 1 Archangel of Thune
- 1 Aurelia's Vindicator
- 1 Basri's Lieutenant
- 1 Bishop of Wings
- 1 Bruna, the Fading Light
- 1 Cathars' Crusade
- 1 Dueling Coach
- 1 Exemplar of Light
- 1 Felidar Retreat
- 1 Giada, Font of Hope
- 1 Herald of War
- 1 Herald of the Sun
- 1 High Sentinels of Arashin
- 1 Inspiring Overseer
- 1 Light of the Legion
- 1 Luminarch Aspirant
- 1 Lyra Dawnbringer
- 1 Mikaeus, the Lunarch
- 1 Nykthos Paragon
- 1 Paragon of Modernity
- 1 Patron of the Valiant
- 1 Righteous Valkyrie
- 1 Sanctuary Warden
- 1 Serra Redeemer
- 1 Winged Shepherd
- 1 Wojek Investigator
- 1 Youthful Valkyrie
- 24 Plains

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
  "question": "This is a one-color deck. By “better lands,” do you want faster mana, utility lands with useful effects, or lands that add other colors?",
  "declined": []
}
```

**The reply:** Question: This is a one-color deck. By “better lands,” do you want faster mana, utility lands with useful effects, or lands that add other colors?


## Revision 3, the answer, base 1: Modern Angels, casual, mono-white, any card

**The user:** Q: This is a one-color deck. By “better lands,” do you want faster mana, utility lands with useful effects, or lands that add other colors? A: Utility lands.

The brief:

```json
{
  "changes": [
    "Replace a few Plains with utility lands that still support a mono-white Angel deck."
  ],
  "remove": null,
  "keep": null,
  "max_mana_value": 0,
  "swap_basics": 3,
  "land_kinds": "Mono-white utility lands that produce white mana or provide useful effects",
  "question": "",
  "declined": []
}
```

**The reply:** I added 1 Hall of Oracles, 1 Idyllic Grange, 1 Karn's Bastion, and 1 Seraph Sanctuary. I changed the count of Plains: 24 to 20.

Findings: curve_summary (INFO)

The deck:

- 1 Admonition Angel
- 1 Anafenza, Kin-Tree Spirit
- 1 Angel of Condemnation
- 1 Angel of Deliverance
- 1 Angel of Invention
- 1 Angel of Sanctions
- 1 Angel of Serenity
- 1 Angel of the God-Pharaoh
- 1 Angel of the Ruins
- 1 Arcane Signet
- 1 Archangel of Thune
- 1 Basri's Lieutenant
- 1 Bishop of Wings
- 1 Bruna, the Fading Light
- 1 Cathars' Crusade
- 1 Dueling Coach
- 1 Exemplar of Light
- 1 Felidar Retreat
- 1 Giada, Font of Hope
- 1 Hall of Oracles
- 1 Herald of War
- 1 Herald of the Sun
- 1 High Sentinels of Arashin
- 1 Idyllic Grange
- 1 Inspiring Overseer
- 1 Karn's Bastion
- 1 Light of the Legion
- 1 Luminarch Aspirant
- 1 Lyra Dawnbringer
- 1 Mikaeus, the Lunarch
- 1 Nykthos Paragon
- 1 Paragon of Modernity
- 1 Patron of the Valiant
- 1 Righteous Valkyrie
- 1 Sanctuary Warden
- 1 Seraph Sanctuary
- 1 Serra Redeemer
- 1 Winged Shepherd
- 1 Wojek Investigator
- 1 Youthful Valkyrie
- 20 Plains

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

**The reply:** I removed 1 Alhammarret's Archive, 1 Angel of Indemnity, 1 Archangel of Thune, 1 Astarion, the Decadent, 1 Boon Reflection, 1 Cavalier of Night, 1 Celestine, the Living Saint, 1 Defiant Bloodlord, 1 Divinity of Pride, 1 Fumigate, 1 Nykthos Paragon, 1 Regal Bloodlord, 1 Sanguine Bond, 1 Solitude, 1 Valkyrie Harbinger, 1 Vona, Butcher of Magan, and 1 Witch of the Moors. I added 1 Angelic Accord, 1 Archon of Sun's Grace, 1 Beza, the Bounding Spring, 1 Consuming Corruption, 1 Cosmos Elixir, 1 Enduring Tenacity, 1 Gideon's Company, 1 Gisela, the Broken Blade, 1 Gumdrop Poisoner // Tempt with Treats, 1 Henrika Domnathi // Henrika, Infernal Seer, 1 Leyline of Hope, 1 Lurrus of the Dream-Den, 1 Marauding Blight-Priest, 1 Qala, Ajani's Pridemate, 1 Starscape Cleric, and 1 The Battle of Bywater. I changed the count of Plains: 13 to 14.

Findings: bracket_prose_rules (INFO), curve_summary (INFO), basics_added (INFO)

The deck:

- 1 Aetherflux Reservoir
- 1 Ajani's Pridemate
- 1 Alseid of Life's Bounty
- 1 Altar of the Pantheon
- 1 Angel of Vitality
- 1 Angelic Accord
- 1 Archivist of Oghma
- 1 Archon of Sun's Grace
- 1 Attended Healer
- 1 Ayli, Eternal Pilgrim
- 1 Beza, the Bounding Spring
- 1 Blighted Steppe
- 1 Bounty Board
- 1 Cleric Class
- 1 Cliffhaven Vampire
- 1 Consuming Corruption
- 1 Cosmos Elixir
- 1 Crypt Ghast
- 1 Dawn of Hope
- 1 Diamond Valley
- 1 Enduring Innocence
- 1 Enduring Tenacity
- 1 Exemplar of Light
- 1 Faith's Shield
- 1 Gideon's Company
- 1 Gisela, the Broken Blade
- 1 Griffin Aerie
- 1 Gumdrop Poisoner // Tempt with Treats
- 1 Heliod, Sun-Crowned
- 1 Henrika Domnathi // Henrika, Infernal Seer
- 1 High Market
- 1 Kabira Crossroads
- 1 Kaya's Wrath
- 1 Legion's Landing // Adanto, the First Fort
- 1 Leyline of Hope
- 1 Lurrus of the Dream-Den
- 1 Mangara, the Diplomat
- 1 Marauding Blight-Priest
- 1 Markov Purifier
- 1 Metropolis Reformer
- 1 Murderous Rider // Swift End
- 1 Nuka-Cola Vending Machine
- 1 Orazca Relic
- 1 Potioner's Trove
- 1 Pristine Talisman
- 1 Qala, Ajani's Pridemate
- 1 Radiant Fountain
- 1 Redemption Choir
- 1 Resplendent Angel
- 1 Restless Fortress
- 1 Restoration Magic
- 1 Rhox Faithmender
- 1 Righteous Valkyrie
- 1 Rune-Tail, Kitsune Ascendant // Rune-Tail's Essence
- 1 Scoured Barrens
- 1 Shambling Vent
- 1 Sigarda's Splendor
- 1 Skyclave Cleric // Skyclave Basilica
- 1 Speaker of the Heavens
- 1 Starscape Cleric
- 1 Sword of Light and Shadow
- 1 The Battle of Bywater
- 1 The Celestus
- 1 The Gaffer
- 1 Twinblade Paladin
- 1 Umezawa's Jitte
- 1 Vampiric Rites
- 1 Vault of the Archangel
- 1 Vito, Thorn of the Dusk Rose
- 1 Vizkopa Guildmage
- 1 Voice of the Blessed
- 1 Well of Lost Dreams
- 1 White Sun's Twilight
- 1 Zof Consumption // Zof Bloodbog
- 11 Swamp
- 14 Plains

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
- 1 Alhammarret's Archive
- 1 Alseid of Life's Bounty
- 1 Altar of the Pantheon
- 1 Angel of Indemnity
- 1 Angel of Vitality
- 1 Archangel of Thune
- 1 Archivist of Oghma
- 1 Astarion, the Decadent
- 1 Attended Healer
- 1 Ayli, Eternal Pilgrim
- 1 Blighted Steppe
- 1 Boon Reflection
- 1 Bounty Board
- 1 Cavalier of Night
- 1 Celestine, the Living Saint
- 1 Cleric Class
- 1 Cliffhaven Vampire
- 1 Crypt Ghast
- 1 Dawn of Hope
- 1 Diamond Valley
- 1 Divinity of Pride
- 1 Enduring Innocence
- 1 Epicure of Blood
- 1 Exemplar of Light
- 1 Faith's Shield
- 1 Fumigate
- 1 Griffin Aerie
- 1 Heliod, Sun-Crowned
- 1 High Market
- 1 Kabira Crossroads
- 1 Kaya's Wrath
- 1 Legion's Landing // Adanto, the First Fort
- 1 Mangara, the Diplomat
- 1 Markov Purifier
- 1 Metropolis Reformer
- 1 Murderous Rider // Swift End
- 1 Nuka-Cola Vending Machine
- 1 Nykthos Paragon
- 1 Orazca Relic
- 1 Potioner's Trove
- 1 Pristine Talisman
- 1 Radiant Fountain
- 1 Redemption Choir
- 1 Regal Bloodlord
- 1 Resplendent Angel
- 1 Restless Fortress
- 1 Restoration Magic
- 1 Rhox Faithmender
- 1 Righteous Valkyrie
- 1 Rune-Tail, Kitsune Ascendant // Rune-Tail's Essence
- 1 Sanguine Bond
- 1 Scoured Barrens
- 1 Shambling Vent
- 1 Sigarda's Splendor
- 1 Skyclave Cleric // Skyclave Basilica
- 1 Solitude
- 1 Speaker of the Heavens
- 1 Sword of Light and Shadow
- 1 The Celestus
- 1 The Gaffer
- 1 Twinblade Paladin
- 1 Umezawa's Jitte
- 1 Valkyrie Harbinger
- 1 Vampiric Rites
- 1 Vault of the Archangel
- 1 Vito, Thorn of the Dusk Rose
- 1 Vizkopa Guildmage
- 1 Voice of the Blessed
- 1 Vona, Butcher of Magan
- 1 Well of Lost Dreams
- 1 White Sun's Twilight
- 1 Witch of the Moors
- 1 Zof Consumption // Zof Bloodbog
- 11 Swamp
- 13 Plains

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

**The reply:** I removed 1 Attended Healer and 1 Twinblade Paladin. I added 1 Cosmos Elixir and 1 Survival Cache.

Findings: bracket_prose_rules (INFO), curve_summary (INFO)

The deck:

- 1 Aetherflux Reservoir
- 1 Ajani's Pridemate
- 1 Alhammarret's Archive
- 1 Alseid of Life's Bounty
- 1 Altar of the Pantheon
- 1 Angel of Indemnity
- 1 Angel of Vitality
- 1 Archangel of Thune
- 1 Archivist of Oghma
- 1 Astarion, the Decadent
- 1 Ayli, Eternal Pilgrim
- 1 Blighted Steppe
- 1 Boon Reflection
- 1 Bounty Board
- 1 Cavalier of Night
- 1 Celestine, the Living Saint
- 1 Cleric Class
- 1 Cliffhaven Vampire
- 1 Cosmos Elixir
- 1 Crypt Ghast
- 1 Dawn of Hope
- 1 Defiant Bloodlord
- 1 Diamond Valley
- 1 Divinity of Pride
- 1 Enduring Innocence
- 1 Exemplar of Light
- 1 Faith's Shield
- 1 Fumigate
- 1 Griffin Aerie
- 1 Heliod, Sun-Crowned
- 1 High Market
- 1 Kabira Crossroads
- 1 Kaya's Wrath
- 1 Legion's Landing // Adanto, the First Fort
- 1 Mangara, the Diplomat
- 1 Markov Purifier
- 1 Metropolis Reformer
- 1 Murderous Rider // Swift End
- 1 Nuka-Cola Vending Machine
- 1 Nykthos Paragon
- 1 Orazca Relic
- 1 Potioner's Trove
- 1 Pristine Talisman
- 1 Radiant Fountain
- 1 Redemption Choir
- 1 Regal Bloodlord
- 1 Resplendent Angel
- 1 Restless Fortress
- 1 Restoration Magic
- 1 Rhox Faithmender
- 1 Righteous Valkyrie
- 1 Rune-Tail, Kitsune Ascendant // Rune-Tail's Essence
- 1 Sanguine Bond
- 1 Scoured Barrens
- 1 Shambling Vent
- 1 Sigarda's Splendor
- 1 Skyclave Cleric // Skyclave Basilica
- 1 Solitude
- 1 Speaker of the Heavens
- 1 Survival Cache
- 1 Sword of Light and Shadow
- 1 The Celestus
- 1 The Gaffer
- 1 Umezawa's Jitte
- 1 Valkyrie Harbinger
- 1 Vampiric Rites
- 1 Vault of the Archangel
- 1 Vito, Thorn of the Dusk Rose
- 1 Vizkopa Guildmage
- 1 Voice of the Blessed
- 1 Vona, Butcher of Magan
- 1 Well of Lost Dreams
- 1 White Sun's Twilight
- 1 Witch of the Moors
- 1 Zof Consumption // Zof Bloodbog
- 11 Swamp
- 13 Plains

## Revision 9, base 2: lifegain Commander, Karlov, bracket 3, any card

**The user:** Add better lands instead of the basics. I mean dual lands that enter untapped where you can, and a couple of utility lands.

The brief:

```json
{
  "changes": [
    "Replace 12 basic lands with untapped black-white dual lands where possible and a couple of utility lands."
  ],
  "remove": null,
  "keep": null,
  "max_mana_value": 0,
  "swap_basics": 12,
  "land_kinds": "Untapped black-white dual lands and a couple of utility lands",
  "question": "",
  "declined": []
}
```

**The reply:** I added 1 Starlit Sanctum. I changed the count of Plains: 13 to 12.

Findings: bracket_prose_rules (INFO), curve_summary (INFO), revision_lands_kept (BLOCK)

The deck:

- 1 Aetherflux Reservoir
- 1 Ajani's Pridemate
- 1 Alhammarret's Archive
- 1 Alseid of Life's Bounty
- 1 Altar of the Pantheon
- 1 Angel of Indemnity
- 1 Angel of Vitality
- 1 Archangel of Thune
- 1 Archivist of Oghma
- 1 Astarion, the Decadent
- 1 Attended Healer
- 1 Ayli, Eternal Pilgrim
- 1 Blighted Steppe
- 1 Boon Reflection
- 1 Bounty Board
- 1 Cavalier of Night
- 1 Celestine, the Living Saint
- 1 Cleric Class
- 1 Cliffhaven Vampire
- 1 Crypt Ghast
- 1 Dawn of Hope
- 1 Defiant Bloodlord
- 1 Diamond Valley
- 1 Divinity of Pride
- 1 Enduring Innocence
- 1 Exemplar of Light
- 1 Faith's Shield
- 1 Fumigate
- 1 Griffin Aerie
- 1 Heliod, Sun-Crowned
- 1 High Market
- 1 Kabira Crossroads
- 1 Kaya's Wrath
- 1 Legion's Landing // Adanto, the First Fort
- 1 Mangara, the Diplomat
- 1 Markov Purifier
- 1 Metropolis Reformer
- 1 Murderous Rider // Swift End
- 1 Nuka-Cola Vending Machine
- 1 Nykthos Paragon
- 1 Orazca Relic
- 1 Potioner's Trove
- 1 Pristine Talisman
- 1 Radiant Fountain
- 1 Redemption Choir
- 1 Regal Bloodlord
- 1 Resplendent Angel
- 1 Restless Fortress
- 1 Restoration Magic
- 1 Rhox Faithmender
- 1 Righteous Valkyrie
- 1 Rune-Tail, Kitsune Ascendant // Rune-Tail's Essence
- 1 Sanguine Bond
- 1 Scoured Barrens
- 1 Shambling Vent
- 1 Sigarda's Splendor
- 1 Skyclave Cleric // Skyclave Basilica
- 1 Solitude
- 1 Speaker of the Heavens
- 1 Starlit Sanctum
- 1 Sword of Light and Shadow
- 1 The Celestus
- 1 The Gaffer
- 1 Twinblade Paladin
- 1 Umezawa's Jitte
- 1 Valkyrie Harbinger
- 1 Vampiric Rites
- 1 Vault of the Archangel
- 1 Vito, Thorn of the Dusk Rose
- 1 Vizkopa Guildmage
- 1 Voice of the Blessed
- 1 Vona, Butcher of Magan
- 1 Well of Lost Dreams
- 1 White Sun's Twilight
- 1 Witch of the Moors
- 1 Zof Consumption // Zof Bloodbog
- 11 Swamp
- 12 Plains

Failures: block findings: revision_lands_kept; the deck holds 1 more nonbasic lands, and the brief asked for 12

## Revision 7, base 3: Commander from the Hobbit set, bracket 3

**The user:** Lower the mana curve. Nothing over 5 mana.

The brief:

```json
{
  "changes": [
    "Lower the mana curve by removing every nonland card with mana value 6 or more."
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

**The reply:** I removed 1 Bag End Banquet, 1 Bilbo's Burglaring, 1 Colossal Whale, 1 Lake-town Mariners // Gone Fishing, 1 Old Fat Spider, and 1 Thranduil's Decree. I added 1 Azog, Moria's Ruin, 1 Lórien Revealed, 1 Stern Scolding, and 1 Thranduil, Sindarin Liege // Silvan Rally. I changed the count of Forest: 9 to 10 and Island: 8 to 9.

Findings: bracket_prose_rules (INFO), curve_summary (INFO), basics_added (INFO)

The deck:

- 1 Along the Crooked Way
- 1 Arcane Signet
- 1 Azog, Moria's Ruin
- 1 Beorn's Hospitality
- 1 Bilbo Baggins, Burglar // Take a Glance
- 1 Bilbo's Deadly Slice
- 1 Bilbo's Ring
- 1 Bilbo, Luckwearer // Burglar's Plot
- 1 Bilbo, Thief in the Night
- 1 Bitter Downfall
- 1 Confusticate and Bebother
- 1 Dancing from Dark to Dawn
- 1 Delighted Halfling
- 1 Down in the Valley
- 1 Dragon-Cursed Halls
- 1 Elven Chorus
- 1 Elven Passage
- 1 Elvenking's Halls
- 1 Elvish Mystic
- 1 Elvish Visionary
- 1 Enchanted River's Grasp
- 1 Giant's Boulder
- 1 Gift of Strands
- 1 Gnashing of Teeth
- 1 Gollum the Abandoned
- 1 Gollum, Silent Slinker // Meager Meal
- 1 Hithlain Knots
- 1 Hobbit Hole
- 1 Inside Information
- 1 Key to the Side-Door
- 1 Lakeshore Apothecary
- 1 Languish
- 1 Long-Lost Lances
- 1 Lórien Revealed
- 1 Merciless Executioner
- 1 Minas Morgul, Dark Fortress
- 1 Mirkwood
- 1 Mithril Coat
- 1 Most Decrepit Old Bird // Speak Secrets
- 1 Mox Amber
- 1 Nasty Little Rabbit
- 1 Necklace of Girion
- 1 Night's Whisper
- 1 Old Fat Spider Can't See Me
- 1 Old Thrush
- 1 Orcish Bowmasters
- 1 Palantír of Orthanc
- 1 Part in Friendship
- 1 Raise the Palisade
- 1 Ravenhill Flock
- 1 Riddles in the Dark
- 1 Rivendell
- 1 Rogue's Passage
- 1 Roll-Roll-Roll-Roll
- 1 Stern Scolding
- 1 Sting, Bilbo's Sword
- 1 Supper for Spiders
- 1 The Black Gate
- 1 The Notary Hobbits
- 1 The One Ring
- 1 The Sackville-Bagginses
- 1 The Shire
- 1 Thranduil, Sindarin Liege // Silvan Rally
- 1 Treasure Vault
- 1 Troop of Ponies
- 1 Uncover the Moon-Letters
- 1 Uneasy Partings
- 1 Wayfarer's Bauble
- 1 Well-Worn Spatula
- 1 Willow-Wind
- 1 Wizard's Staff
- 1 Wood Elves
- 10 Forest
- 8 Swamp
- 9 Island

## Revision 8, base 3: Commander from the Hobbit set, bracket 3

**The user:** Take out Colossal Whale.

The brief:

```json
{
  "changes": [],
  "remove": [
    "Colossal Whale"
  ],
  "keep": null,
  "max_mana_value": 0,
  "swap_basics": 0,
  "land_kinds": "",
  "question": "",
  "declined": []
}
```

**The reply:** I removed 1 Colossal Whale. I added 1 The Black Arrow.

Findings: bracket_prose_rules (INFO), curve_summary (INFO)

The deck:

- 1 Along the Crooked Way
- 1 Arcane Signet
- 1 Bag End Banquet
- 1 Beorn's Hospitality
- 1 Bilbo Baggins, Burglar // Take a Glance
- 1 Bilbo's Burglaring
- 1 Bilbo's Deadly Slice
- 1 Bilbo's Ring
- 1 Bilbo, Luckwearer // Burglar's Plot
- 1 Bilbo, Thief in the Night
- 1 Bitter Downfall
- 1 Confusticate and Bebother
- 1 Dancing from Dark to Dawn
- 1 Delighted Halfling
- 1 Down in the Valley
- 1 Dragon-Cursed Halls
- 1 Elven Chorus
- 1 Elven Passage
- 1 Elvenking's Halls
- 1 Elvish Mystic
- 1 Elvish Visionary
- 1 Enchanted River's Grasp
- 1 Giant's Boulder
- 1 Gift of Strands
- 1 Gnashing of Teeth
- 1 Gollum the Abandoned
- 1 Gollum, Silent Slinker // Meager Meal
- 1 Hithlain Knots
- 1 Hobbit Hole
- 1 Inside Information
- 1 Key to the Side-Door
- 1 Lake-town Mariners // Gone Fishing
- 1 Lakeshore Apothecary
- 1 Languish
- 1 Long-Lost Lances
- 1 Merciless Executioner
- 1 Minas Morgul, Dark Fortress
- 1 Mirkwood
- 1 Mithril Coat
- 1 Most Decrepit Old Bird // Speak Secrets
- 1 Mox Amber
- 1 Nasty Little Rabbit
- 1 Necklace of Girion
- 1 Night's Whisper
- 1 Old Fat Spider
- 1 Old Fat Spider Can't See Me
- 1 Old Thrush
- 1 Orcish Bowmasters
- 1 Palantír of Orthanc
- 1 Part in Friendship
- 1 Raise the Palisade
- 1 Ravenhill Flock
- 1 Riddles in the Dark
- 1 Rivendell
- 1 Rogue's Passage
- 1 Roll-Roll-Roll-Roll
- 1 Sting, Bilbo's Sword
- 1 Supper for Spiders
- 1 The Black Arrow
- 1 The Black Gate
- 1 The Notary Hobbits
- 1 The One Ring
- 1 The Sackville-Bagginses
- 1 The Shire
- 1 Thranduil's Decree
- 1 Treasure Vault
- 1 Troop of Ponies
- 1 Uncover the Moon-Letters
- 1 Uneasy Partings
- 1 Wayfarer's Bauble
- 1 Well-Worn Spatula
- 1 Willow-Wind
- 1 Wizard's Staff
- 1 Wood Elves
- 8 Island
- 8 Swamp
- 9 Forest
