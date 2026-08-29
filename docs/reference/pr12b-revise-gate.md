# PR-12B revise gate

Run date: 2026-08-28. Snapshot: 2026-08-28.

Verdict: FAIL. 5 of 6 revisions met their bar. The bars: an unclear request gets a question or a decline with a reason, a clear request gets a revised deck with no block finding, the deck holds every cap and every removal the message names, and it keeps at least 80 percent of the untouched cards.

Cost: 13 calls, $0.3263. Time: 3m36s.

| Base | Revision | Message | Outcome | Kept | Result |
|---|---|---|---|---|---|
| 1 | 1 | Replace some lands with better options if possible. Also tune the mana curve lower - no 6 or 7 mana cards needed | question |  | FAIL: a clear request got a question: For better lands, do you want faster white mana, utility lands, or lands that support adding more colors? |
| 1 | 2 | Take out Angel of Deliverance. | 60 cards | 100% | PASS, kept 100% |
| 1 | 3 | Better lands please. | question |  | PASS, no build |
| 2 | 4 | No nonland cards over 4 mana. | 99 cards | 97% | PASS, kept 97% |
| 2 | 5 | Swap Defiant Bloodlord for something cheaper that does the same job. | 99 cards | 100% | PASS, kept 100% |
| 2 | 6 | Add two more cards that draw cards, and cut two of the weakest creatures. | 99 cards | 97% | PASS, kept 97% |

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
  "question": "For better lands, do you want faster white mana, utility lands, or lands that support adding more colors?",
  "declined": []
}
```

**The reply:** Question: For better lands, do you want faster white mana, utility lands, or lands that support adding more colors?


Failures: a clear request got a question: For better lands, do you want faster white mana, utility lands, or lands that support adding more colors?

## Revision 2, base 1: Modern Angels, casual, mono-white, any card

**The user:** Take out Angel of Deliverance.

The brief:

```json
{
  "changes": [],
  "remove": [
    "Angel of Deliverance"
  ],
  "keep": null,
  "max_mana_value": 0,
  "question": "",
  "declined": []
}
```

**The reply:** I removed 1 Angel of Deliverance. I added 1 Aurelia's Vindicator.

Findings: curve_summary (INFO)

The deck:

- 1 Abandoned Air Temple
- 1 Abzan Falconer
- 1 Admonition Angel
- 1 Ainok Bond-Kin
- 1 Ajani, Adversary of Tyrants
- 1 Ajani, Caller of the Pride
- 1 Angel of Condemnation
- 1 Angel of Invention
- 1 Angel of Sanctions
- 1 Angel of Serenity
- 1 Angel of the God-Pharaoh
- 1 Angel of the Ruins
- 1 Animal Sanctuary
- 1 Arcane Signet
- 1 Archangel of Thune
- 1 Aurelia's Vindicator
- 1 Basri's Lieutenant
- 1 Bruna, the Fading Light
- 1 Cathars' Crusade
- 1 Cavern of Souls
- 1 Command Tower
- 1 Evolving Wilds
- 1 Exemplar of Light
- 1 Exotic Orchard
- 1 Felidar Retreat
- 1 Giada, Font of Hope
- 1 Guildmages' Forum
- 1 Herald of War
- 1 Herald of the Sun
- 1 Inspiring Overseer
- 1 Karn's Bastion
- 1 Light of the Legion
- 1 Luminarch Aspirant
- 1 Lyra Dawnbringer
- 1 Mikaeus, the Lunarch
- 1 Nykthos Paragon
- 1 Paragon of Modernity
- 1 Patron of the Valiant
- 1 Sanctuary Warden
- 1 Seraph Sanctuary
- 1 Serra Redeemer
- 1 Sigarda's Summons
- 1 Tyrite Sanctum
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
  "question": "What kind of better lands do you want: faster white mana, utility lands, or lands that support adding more colors?",
  "declined": []
}
```

**The reply:** Question: What kind of better lands do you want: faster white mana, utility lands, or lands that support adding more colors?


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

**The reply:** I removed 1 Alhammarret's Archive, 1 Angel of Indemnity, 1 Archangel of Thune, 1 Blood Baron of Vizkopa, 1 Celestine, the Living Saint, 1 Consuming Corruption, 1 Defiant Bloodlord, 1 Divinity of Pride, 1 Epicure of Blood, 1 Fumigate, 1 Nykthos Paragon, 1 Sanguine Bond, 1 Sigarda's Splendor, 1 Solitude, 1 Valkyrie Harbinger, 1 Vein Ripper, 1 Vona, Butcher of Magan, and 1 Wurmcoil Engine. I added 1 Ajani, Strength of the Pride, 1 Angelic Accord, 1 Archon of Sun's Grace, 1 Beza, the Bounding Spring, 1 Bloodbond Vampire, 1 Cliffhaven Vampire, 1 Doctor Jane Foster, 1 Elenda, Saint of Dusk, 1 Enduring Tenacity, 1 Gideon's Company, 1 Gisela, the Broken Blade, 1 Henrika Domnathi // Henrika, Infernal Seer, 1 Qala, Ajani's Pridemate, 1 Sorin of House Markov // Sorin, Ravenous Neonate, 1 Sorin, Imperious Bloodlord, 1 The Gaffer, 1 Umezawa's Jitte, and 1 Vizkopa Guildmage.

Findings: bracket_prose_rules (INFO), curve_summary (INFO)

The deck:

- 1 Aetherflux Reservoir
- 1 Ajani's Pridemate
- 1 Ajani, Strength of the Pride
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
- 1 Bloodbond Vampire
- 1 Bounty Board
- 1 Cleric Class
- 1 Cleric of Life's Bond
- 1 Cliffhaven Vampire
- 1 Cosmos Elixir
- 1 Crypt Ghast
- 1 Dawn of Hope
- 1 Doctor Jane Foster
- 1 Elenda, Saint of Dusk
- 1 Enduring Innocence
- 1 Enduring Tenacity
- 1 Faith's Shield
- 1 Gideon's Company
- 1 Gisela, the Broken Blade
- 1 Heliod, Sun-Crowned
- 1 Henrika Domnathi // Henrika, Infernal Seer
- 1 Hierophant's Chalice
- 1 High Market
- 1 Indulging Patrician
- 1 Kabira Crossroads
- 1 Kaya's Wrath
- 1 Leyline of Hope
- 1 Light of Promise
- 1 Mangara, the Diplomat
- 1 Markov Purifier
- 1 Metropolis Reformer
- 1 Murderous Rider // Swift End
- 1 Nightmare's Thirst
- 1 Nuka-Cola Vending Machine
- 1 Orazca Relic
- 1 Potioner's Trove
- 1 Pristine Talisman
- 1 Qala, Ajani's Pridemate
- 1 Radiant Fountain
- 1 Resplendent Angel
- 1 Restless Fortress
- 1 Rhox Faithmender
- 1 Righteous Valkyrie
- 1 Rune-Tail, Kitsune Ascendant // Rune-Tail's Essence
- 1 Scoured Barrens
- 1 Shambling Vent
- 1 Skyclave Cleric // Skyclave Basilica
- 1 Sorin of House Markov // Sorin, Ravenous Neonate
- 1 Sorin, Imperious Bloodlord
- 1 Survival Cache
- 1 Sword of Light and Shadow
- 1 The Celestus
- 1 The Gaffer
- 1 The Meathook Massacre
- 1 Twinblade Paladin
- 1 Umezawa's Jitte
- 1 Vampiric Rites
- 1 Vault of the Archangel
- 1 Vito, Thorn of the Dusk Rose
- 1 Vizkopa Guildmage
- 1 Voice of the Blessed
- 1 Well of Lost Dreams
- 1 Werefox Bodyguard
- 1 Zof Consumption // Zof Bloodbog
- 10 Swamp
- 16 Plains

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

- 1 Aetherflux Reservoir
- 1 Ajani's Pridemate
- 1 Alhammarret's Archive
- 1 Alseid of Life's Bounty
- 1 Altar of the Pantheon
- 1 Angel of Indemnity
- 1 Angel of Vitality
- 1 Archangel of Thune
- 1 Archivist of Oghma
- 1 Attended Healer
- 1 Ayli, Eternal Pilgrim
- 1 Blighted Steppe
- 1 Blood Baron of Vizkopa
- 1 Bounty Board
- 1 Celestine, the Living Saint
- 1 Cleric Class
- 1 Cleric of Life's Bond
- 1 Consuming Corruption
- 1 Cosmos Elixir
- 1 Crypt Ghast
- 1 Dawn of Hope
- 1 Divinity of Pride
- 1 Enduring Innocence
- 1 Epicure of Blood
- 1 Faith's Shield
- 1 Fumigate
- 1 Heliod, Sun-Crowned
- 1 Hierophant's Chalice
- 1 High Market
- 1 Indulging Patrician
- 1 Kabira Crossroads
- 1 Kaya's Wrath
- 1 Leyline of Hope
- 1 Light of Promise
- 1 Mangara, the Diplomat
- 1 Marauding Blight-Priest
- 1 Markov Purifier
- 1 Metropolis Reformer
- 1 Murderous Rider // Swift End
- 1 Nightmare's Thirst
- 1 Nuka-Cola Vending Machine
- 1 Nykthos Paragon
- 1 Orazca Relic
- 1 Potioner's Trove
- 1 Pristine Talisman
- 1 Radiant Fountain
- 1 Resplendent Angel
- 1 Restless Fortress
- 1 Rhox Faithmender
- 1 Righteous Valkyrie
- 1 Rune-Tail, Kitsune Ascendant // Rune-Tail's Essence
- 1 Sanguine Bond
- 1 Scoured Barrens
- 1 Shambling Vent
- 1 Sigarda's Splendor
- 1 Skyclave Cleric // Skyclave Basilica
- 1 Solitude
- 1 Survival Cache
- 1 Sword of Light and Shadow
- 1 The Celestus
- 1 The Meathook Massacre
- 1 Twinblade Paladin
- 1 Valkyrie Harbinger
- 1 Vampiric Rites
- 1 Vault of the Archangel
- 1 Vein Ripper
- 1 Vito, Thorn of the Dusk Rose
- 1 Voice of the Blessed
- 1 Vona, Butcher of Magan
- 1 Well of Lost Dreams
- 1 Werefox Bodyguard
- 1 Wurmcoil Engine
- 1 Zof Consumption // Zof Bloodbog
- 10 Swamp
- 16 Plains

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

**The reply:** I removed 1 Blood Baron of Vizkopa and 1 Epicure of Blood. I added 1 Convalescent Care and 1 The Gaffer.

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
- 1 Attended Healer
- 1 Ayli, Eternal Pilgrim
- 1 Blighted Steppe
- 1 Bounty Board
- 1 Celestine, the Living Saint
- 1 Cleric Class
- 1 Cleric of Life's Bond
- 1 Consuming Corruption
- 1 Convalescent Care
- 1 Cosmos Elixir
- 1 Crypt Ghast
- 1 Dawn of Hope
- 1 Defiant Bloodlord
- 1 Divinity of Pride
- 1 Enduring Innocence
- 1 Faith's Shield
- 1 Fumigate
- 1 Heliod, Sun-Crowned
- 1 Hierophant's Chalice
- 1 High Market
- 1 Indulging Patrician
- 1 Kabira Crossroads
- 1 Kaya's Wrath
- 1 Leyline of Hope
- 1 Light of Promise
- 1 Mangara, the Diplomat
- 1 Markov Purifier
- 1 Metropolis Reformer
- 1 Murderous Rider // Swift End
- 1 Nightmare's Thirst
- 1 Nuka-Cola Vending Machine
- 1 Nykthos Paragon
- 1 Orazca Relic
- 1 Potioner's Trove
- 1 Pristine Talisman
- 1 Radiant Fountain
- 1 Resplendent Angel
- 1 Restless Fortress
- 1 Rhox Faithmender
- 1 Righteous Valkyrie
- 1 Rune-Tail, Kitsune Ascendant // Rune-Tail's Essence
- 1 Sanguine Bond
- 1 Scoured Barrens
- 1 Shambling Vent
- 1 Sigarda's Splendor
- 1 Skyclave Cleric // Skyclave Basilica
- 1 Solitude
- 1 Survival Cache
- 1 Sword of Light and Shadow
- 1 The Celestus
- 1 The Gaffer
- 1 The Meathook Massacre
- 1 Twinblade Paladin
- 1 Valkyrie Harbinger
- 1 Vampiric Rites
- 1 Vault of the Archangel
- 1 Vein Ripper
- 1 Vito, Thorn of the Dusk Rose
- 1 Voice of the Blessed
- 1 Vona, Butcher of Magan
- 1 Well of Lost Dreams
- 1 Werefox Bodyguard
- 1 Wurmcoil Engine
- 1 Zof Consumption // Zof Bloodbog
- 10 Swamp
- 16 Plains
