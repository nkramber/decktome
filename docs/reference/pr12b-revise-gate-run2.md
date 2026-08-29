# PR-12B revise gate

Run date: 2026-08-28. Snapshot: 2026-08-28.

Verdict: PASS. 6 of 6 revisions met their bar. The bars: an unclear request gets a question or a decline with a reason, a clear request gets a revised deck with no block finding, the deck holds every cap and every removal the message names, and it keeps at least 80 percent of the untouched cards.

Cost: 13 calls, $0.2941. Time: 3m22s.

| Base | Revision | Message | Outcome | Kept | Result |
|---|---|---|---|---|---|
| 1 | 1 | Replace some lands with better options if possible. Also tune the mana curve lower - no 6 or 7 mana cards needed | question |  | PASS, no build |
| 1 | 2 | Take out Angel of Deliverance. | 60 cards | 100% | PASS, kept 100% |
| 1 | 3 | Better lands please. | question |  | PASS, no build |
| 2 | 4 | No nonland cards over 4 mana. | 99 cards | 98% | PASS, kept 98% |
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
  "question": "For the land replacements, do you want faster mana, utility lands with useful effects, or lands that support adding more colors?",
  "declined": []
}
```

**The reply:** Question: For the land replacements, do you want faster mana, utility lands with useful effects, or lands that support adding more colors?


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

- 1 Abzan Falconer
- 1 Admonition Angel
- 1 Aetherstorm Roc
- 1 Ainok Bond-Kin
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
- 1 Paragon of Modernity
- 1 Patron of the Valiant
- 1 Righteous Valkyrie
- 1 Sanctuary Warden
- 1 Serra Redeemer
- 1 Sigarda's Summons
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
  "question": "What kind of land upgrade do you want: faster mana, utility lands, or lands that support adding more colors?",
  "declined": []
}
```

**The reply:** Question: What kind of land upgrade do you want: faster mana, utility lands, or lands that support adding more colors?


## Revision 4, base 2: lifegain Commander, Karlov, bracket 3, any card

**The user:** No nonland cards over 4 mana.

The brief:

```json
{
  "changes": [
    "Remove every nonland card with mana value 5 or more"
  ],
  "remove": null,
  "keep": null,
  "max_mana_value": 4,
  "question": "",
  "declined": []
}
```

**The reply:** I removed 1 Alhammarret's Archive, 1 Angel of Indemnity, 1 Archangel of Thune, 1 Astarion, the Decadent, 1 Blood Baron of Vizkopa, 1 Carmen, Cruel Skymarcher, 1 Cavalier of Night, 1 Defiant Bloodlord, 1 Divinity of Pride, 1 Fumigate, 1 Liesa, Forgotten Archangel, 1 Lyra Dawnbringer, 1 Nykthos Paragon, 1 Sanguine Bond, 1 Sigarda's Splendor, 1 Solitude, 1 Valkyrie Harbinger, and 1 Vein Ripper. I added 1 Ajani Goldmane, 1 Archon of Sun's Grace, 1 Blood Artist, 1 Bloodbond Vampire, 1 Cleric of Life's Bond, 1 Cliffhaven Vampire, 1 Elenda, Saint of Dusk, 1 Enduring Tenacity, 1 Essence Pulse, 1 Gisela, the Broken Blade, 1 Gumdrop Poisoner // Tempt with Treats, 1 Henrika Domnathi // Henrika, Infernal Seer, 1 Orazca Relic, 1 Papalymo Totolymo, 1 Qala, Ajani's Pridemate, 1 The Celestus, 1 The Gaffer, and 1 Twinblade Paladin.

Findings: bracket_prose_rules (INFO), curve_summary (INFO)

The deck:

- 1 Aetherflux Reservoir
- 1 Ajani Goldmane
- 1 Ajani's Pridemate
- 1 Ajani, Strength of the Pride
- 1 Alseid of Life's Bounty
- 1 Angel of Vitality
- 1 Angelic Accord
- 1 Archivist of Oghma
- 1 Archon of Sun's Grace
- 1 Attended Healer
- 1 Ayara, First of Locthwain
- 1 Ayli, Eternal Pilgrim
- 1 Battle Angels of Tyr
- 1 Beza, the Bounding Spring
- 1 Blood Artist
- 1 Bloodbond Vampire
- 1 Bloodthirsty Aerialist
- 1 Bounty Board
- 1 Cleric Class
- 1 Cleric of Life's Bond
- 1 Cliffhaven Vampire
- 1 Colossal Plow
- 1 Consuming Corruption
- 1 Cosmos Elixir
- 1 Crypt Ghast
- 1 Dawn of Hope
- 1 Elenda, Saint of Dusk
- 1 Enduring Innocence
- 1 Enduring Tenacity
- 1 Essence Pulse
- 1 Exemplar of Light
- 1 Faith's Shield
- 1 Gideon's Company
- 1 Gisela, the Broken Blade
- 1 Gumdrop Poisoner // Tempt with Treats
- 1 Heliod, Sun-Crowned
- 1 Henrika Domnathi // Henrika, Infernal Seer
- 1 Hierophant's Chalice
- 1 Kaya's Wrath
- 1 Leyline of Hope
- 1 Mangara, the Diplomat
- 1 Metropolis Reformer
- 1 Murderous Rider // Swift End
- 1 Nightmare's Thirst
- 1 Nuka-Cola Vending Machine
- 1 Orazca Relic
- 1 Papalymo Totolymo
- 1 Pristine Talisman
- 1 Qala, Ajani's Pridemate
- 1 Resplendent Angel
- 1 Rhox Faithmender
- 1 Righteous Valkyrie
- 1 Rune-Tail, Kitsune Ascendant // Rune-Tail's Essence
- 1 Serra Ascendant
- 1 Sword of Light and Shadow
- 1 The Celestus
- 1 The Gaffer
- 1 Twinblade Paladin
- 1 Vampiric Rites
- 1 Vito, Thorn of the Dusk Rose
- 1 Voice of the Blessed
- 1 Well of Lost Dreams
- 1 Werefox Bodyguard
- 18 Plains
- 18 Swamp

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

**The reply:** I removed 1 Defiant Bloodlord. I added 1 Enduring Tenacity.

Findings: bracket_prose_rules (INFO), curve_summary (INFO)

The deck:

- 1 Aetherflux Reservoir
- 1 Ajani's Pridemate
- 1 Ajani, Strength of the Pride
- 1 Alhammarret's Archive
- 1 Alseid of Life's Bounty
- 1 Angel of Indemnity
- 1 Angel of Vitality
- 1 Angelic Accord
- 1 Archangel of Thune
- 1 Archivist of Oghma
- 1 Astarion, the Decadent
- 1 Attended Healer
- 1 Ayara, First of Locthwain
- 1 Ayli, Eternal Pilgrim
- 1 Battle Angels of Tyr
- 1 Beza, the Bounding Spring
- 1 Blood Baron of Vizkopa
- 1 Bloodthirsty Aerialist
- 1 Bounty Board
- 1 Carmen, Cruel Skymarcher
- 1 Cavalier of Night
- 1 Cleric Class
- 1 Colossal Plow
- 1 Consuming Corruption
- 1 Cosmos Elixir
- 1 Crypt Ghast
- 1 Dawn of Hope
- 1 Divinity of Pride
- 1 Enduring Innocence
- 1 Enduring Tenacity
- 1 Exemplar of Light
- 1 Faith's Shield
- 1 Fumigate
- 1 Gideon's Company
- 1 Heliod, Sun-Crowned
- 1 Hierophant's Chalice
- 1 Kaya's Wrath
- 1 Leyline of Hope
- 1 Liesa, Forgotten Archangel
- 1 Lyra Dawnbringer
- 1 Mangara, the Diplomat
- 1 Metropolis Reformer
- 1 Murderous Rider // Swift End
- 1 Nightmare's Thirst
- 1 Nuka-Cola Vending Machine
- 1 Nykthos Paragon
- 1 Pristine Talisman
- 1 Resplendent Angel
- 1 Rhox Faithmender
- 1 Righteous Valkyrie
- 1 Rune-Tail, Kitsune Ascendant // Rune-Tail's Essence
- 1 Sanguine Bond
- 1 Serra Ascendant
- 1 Sigarda's Splendor
- 1 Solitude
- 1 Sword of Light and Shadow
- 1 Valkyrie Harbinger
- 1 Vampiric Rites
- 1 Vein Ripper
- 1 Vito, Thorn of the Dusk Rose
- 1 Voice of the Blessed
- 1 Well of Lost Dreams
- 1 Werefox Bodyguard
- 18 Plains
- 18 Swamp

## Revision 6, base 2: lifegain Commander, Karlov, bracket 3, any card

**The user:** Add two more cards that draw cards, and cut two of the weakest creatures.

The brief:

```json
{
  "changes": [
    "Add two more cards that draw cards.",
    "Cut two of the weakest creatures."
  ],
  "remove": null,
  "keep": null,
  "max_mana_value": 0,
  "question": "",
  "declined": []
}
```

**The reply:** I removed 1 Attended Healer and 1 Gideon's Company. I added 1 Survival Cache and 1 The Gaffer.

Findings: bracket_prose_rules (INFO), curve_summary (INFO)

The deck:

- 1 Aetherflux Reservoir
- 1 Ajani's Pridemate
- 1 Ajani, Strength of the Pride
- 1 Alhammarret's Archive
- 1 Alseid of Life's Bounty
- 1 Angel of Indemnity
- 1 Angel of Vitality
- 1 Angelic Accord
- 1 Archangel of Thune
- 1 Archivist of Oghma
- 1 Astarion, the Decadent
- 1 Ayara, First of Locthwain
- 1 Ayli, Eternal Pilgrim
- 1 Battle Angels of Tyr
- 1 Beza, the Bounding Spring
- 1 Blood Baron of Vizkopa
- 1 Bloodthirsty Aerialist
- 1 Bounty Board
- 1 Carmen, Cruel Skymarcher
- 1 Cavalier of Night
- 1 Cleric Class
- 1 Colossal Plow
- 1 Consuming Corruption
- 1 Cosmos Elixir
- 1 Crypt Ghast
- 1 Dawn of Hope
- 1 Defiant Bloodlord
- 1 Divinity of Pride
- 1 Enduring Innocence
- 1 Exemplar of Light
- 1 Faith's Shield
- 1 Fumigate
- 1 Heliod, Sun-Crowned
- 1 Hierophant's Chalice
- 1 Kaya's Wrath
- 1 Leyline of Hope
- 1 Liesa, Forgotten Archangel
- 1 Lyra Dawnbringer
- 1 Mangara, the Diplomat
- 1 Metropolis Reformer
- 1 Murderous Rider // Swift End
- 1 Nightmare's Thirst
- 1 Nuka-Cola Vending Machine
- 1 Nykthos Paragon
- 1 Pristine Talisman
- 1 Resplendent Angel
- 1 Rhox Faithmender
- 1 Righteous Valkyrie
- 1 Rune-Tail, Kitsune Ascendant // Rune-Tail's Essence
- 1 Sanguine Bond
- 1 Serra Ascendant
- 1 Sigarda's Splendor
- 1 Solitude
- 1 Survival Cache
- 1 Sword of Light and Shadow
- 1 The Gaffer
- 1 Valkyrie Harbinger
- 1 Vampiric Rites
- 1 Vein Ripper
- 1 Vito, Thorn of the Dusk Rose
- 1 Voice of the Blessed
- 1 Well of Lost Dreams
- 1 Werefox Bodyguard
- 18 Plains
- 18 Swamp
