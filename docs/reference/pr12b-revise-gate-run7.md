# PR-12B revise gate

Run date: 2026-09-02. Snapshot: 2026-09-01.

Verdict: PASS. 11 of 11 turns met their bar. The bars: an unclear request gets a question or a decline with a reason, an answer to that question gets a rebuild, a clear request gets a revised deck with no block finding, the deck holds every cap, every removal, and every land swap the message names, and it keeps at least 80 percent of the untouched cards.

Cost: 25 calls, $0.7443. Time: 12m31s.

| Base | Revision | Message | Outcome | Kept | Result |
|---|---|---|---|---|---|
| 1 | 1 | Replace some lands with better options if possible. Also tune the mana curve lower - no 6 or 7 mana cards needed | question |  | PASS, no build |
| 1 | 1 | (the answer) Q: For the land upgrades, do you want faster mana, utility lands, or support for additional colors? This is currently a one-color deck, so I need to know what kind of improvement you want before replacing basics. A: A mix. | 60 cards | 100% | PASS, kept 100% |
| 1 | 2 | Take out Angel of Serenity. | 60 cards | 100% | PASS, kept 100% |
| 1 | 3 | Better lands please. | question |  | PASS, no build |
| 1 | 3 | (the answer) Q: By “better lands,” do you want faster mana, utility lands, or lands that add more colors? This is a one-color deck, so I need the kind of upgrade before replacing any Plains. A: Utility lands. | 60 cards | 100% | PASS, kept 100% |
| 2 | 4 | No nonland cards over 4 mana. | 99 cards | 96% | PASS, kept 96% |
| 2 | 5 | Swap Defiant Bloodlord for something cheaper that does the same job. | 99 cards | 99% | PASS, kept 99% |
| 2 | 6 | Add two more cards that draw cards, and cut two of the weakest creatures. | 99 cards | 98% | PASS, kept 98% |
| 2 | 9 | Add better lands instead of the basics. I mean dual lands that enter untapped where you can, and a couple of utility lands. | 99 cards | 98% | PASS, kept 98% |
| 3 | 7 | Lower the mana curve. Nothing over 5 mana. | 99 cards | 99% | PASS, kept 99% |
| 3 | 8 | Take out The Lord of the Eagles. | 99 cards | 100% | PASS, kept 100% |

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
  "question": "For the land upgrades, do you want faster mana, utility lands, or support for additional colors? This is currently a one-color deck, so I need to know what kind of improvement you want before replacing basics.",
  "declined": []
}
```

**The reply:** Question: For the land upgrades, do you want faster mana, utility lands, or support for additional colors? This is currently a one-color deck, so I need to know what kind of improvement you want before replacing basics.


## Revision 1, the answer, base 1: Modern Angels, casual, mono-white, any card

**The user:** Q: For the land upgrades, do you want faster mana, utility lands, or support for additional colors? This is currently a one-color deck, so I need to know what kind of improvement you want before replacing basics. A: A mix.

The brief:

```json
{
  "changes": [
    "Replace a few basic Plains with a mix of faster mana and utility lands.",
    "Remove every nonland card with mana value 6 or more."
  ],
  "remove": null,
  "keep": null,
  "max_mana_value": 5,
  "swap_basics": 3,
  "land_kinds": "A mix of faster mana lands and utility lands",
  "question": "",
  "declined": []
}
```

**The reply:** I removed 1 Admonition Angel, 1 Angel of Serenity, 1 Angel of the God-Pharaoh, 1 Angel of the Ruins, 1 Bruna, the Fading Light, 1 Herald of the Sun, 1 Light of the Legion, 1 Sanctuary Warden, 1 Sigarda's Summons, and 1 Winged Shepherd. I added 1 Angel of Destiny, 1 Dusk Legion Duelist, 1 Gemstone Caverns, 1 Invasion of Dominaria // Serra Faithkeeper, 1 Ruins of Oran-Rief, 1 Sparring Regimen, and 1 Tyrite Sanctum. I changed the count of Abzan Falconer: 1 to 2, Angel of Condemnation: 1 to 2, Angel of Sanctions: 1 to 2, Archangel of Thune: 1 to 2, Aurelia's Vindicator: 1 to 2, and Plains: 15 to 13.

Findings: curve_summary (INFO), basics_added (INFO)

Repair turn: 1 missed names.

The deck:

- 1 Angel of Destiny
- 1 Angel of Invention
- 1 Arcane Signet
- 1 Basri's Lieutenant
- 1 Bishop of Wings
- 1 Castle Ardenvale
- 1 Cathars' Crusade
- 1 Dueling Coach
- 1 Dusk Legion Duelist
- 1 Eiganjo, Seat of the Empire
- 1 Elite Scaleguard
- 1 Exemplar of Light
- 1 Felidar Retreat
- 1 Gemstone Caverns
- 1 Giada, Font of Hope
- 1 Herald of War
- 1 High Sentinels of Arashin
- 1 Inspiring Overseer
- 1 Invasion of Dominaria // Serra Faithkeeper
- 1 Karn's Bastion
- 1 Luminarch Aspirant
- 1 Lyra Dawnbringer
- 1 Minas Tirith
- 1 Nykthos, Shrine to Nyx
- 1 Paragon of Modernity
- 1 Patron of the Valiant
- 1 Righteous Valkyrie
- 1 Ruins of Oran-Rief
- 1 Secluded Courtyard
- 1 Seraph Sanctuary
- 1 Serra Redeemer
- 1 Sparring Regimen
- 1 Three Tree City
- 1 Tyrite Sanctum
- 1 Windbrisk Heights
- 1 Wojek Investigator
- 1 Youthful Valkyrie
- 13 Plains
- 2 Abzan Falconer
- 2 Angel of Condemnation
- 2 Angel of Sanctions
- 2 Archangel of Thune
- 2 Aurelia's Vindicator

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
  "swap_basics": 0,
  "land_kinds": "",
  "question": "",
  "declined": []
}
```

**The reply:** I removed 1 Angel of Serenity. I added 1 Angel of Deliverance.

Findings: curve_summary (INFO)

The deck:

- 1 Abzan Falconer
- 1 Admonition Angel
- 1 Angel of Condemnation
- 1 Angel of Deliverance
- 1 Angel of Invention
- 1 Angel of Sanctions
- 1 Angel of the God-Pharaoh
- 1 Angel of the Ruins
- 1 Arcane Signet
- 1 Archangel of Thune
- 1 Aurelia's Vindicator
- 1 Basri's Lieutenant
- 1 Bishop of Wings
- 1 Bruna, the Fading Light
- 1 Castle Ardenvale
- 1 Cathars' Crusade
- 1 Dueling Coach
- 1 Eiganjo, Seat of the Empire
- 1 Elite Scaleguard
- 1 Exemplar of Light
- 1 Felidar Retreat
- 1 Giada, Font of Hope
- 1 Herald of War
- 1 Herald of the Sun
- 1 High Sentinels of Arashin
- 1 Inspiring Overseer
- 1 Karn's Bastion
- 1 Light of the Legion
- 1 Luminarch Aspirant
- 1 Lyra Dawnbringer
- 1 Minas Tirith
- 1 Nykthos, Shrine to Nyx
- 1 Paragon of Modernity
- 1 Patron of the Valiant
- 1 Righteous Valkyrie
- 1 Sanctuary Warden
- 1 Secluded Courtyard
- 1 Seraph Sanctuary
- 1 Serra Redeemer
- 1 Sigarda's Summons
- 1 Three Tree City
- 1 Windbrisk Heights
- 1 Winged Shepherd
- 1 Wojek Investigator
- 1 Youthful Valkyrie
- 15 Plains

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
  "question": "By “better lands,” do you want faster mana, utility lands, or lands that add more colors? This is a one-color deck, so I need the kind of upgrade before replacing any Plains.",
  "declined": []
}
```

**The reply:** Question: By “better lands,” do you want faster mana, utility lands, or lands that add more colors? This is a one-color deck, so I need the kind of upgrade before replacing any Plains.


## Revision 3, the answer, base 1: Modern Angels, casual, mono-white, any card

**The user:** Q: By “better lands,” do you want faster mana, utility lands, or lands that add more colors? This is a one-color deck, so I need the kind of upgrade before replacing any Plains. A: Utility lands.

The brief:

```json
{
  "changes": [
    "Replace a few Plains with utility lands that support the Angel and +1/+1 counter strategy."
  ],
  "remove": null,
  "keep": null,
  "max_mana_value": 0,
  "swap_basics": 3,
  "land_kinds": "Utility lands that support Angels and +1/+1 counters",
  "question": "",
  "declined": []
}
```

**The reply:** I added 1 Abandoned Air Temple, 1 Ruins of Oran-Rief, and 1 Tyrite Sanctum. I changed the count of Plains: 15 to 12.

Findings: curve_summary (INFO)

The deck:

- 1 Abandoned Air Temple
- 1 Abzan Falconer
- 1 Admonition Angel
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
- 1 Castle Ardenvale
- 1 Cathars' Crusade
- 1 Dueling Coach
- 1 Eiganjo, Seat of the Empire
- 1 Elite Scaleguard
- 1 Exemplar of Light
- 1 Felidar Retreat
- 1 Giada, Font of Hope
- 1 Herald of War
- 1 Herald of the Sun
- 1 High Sentinels of Arashin
- 1 Inspiring Overseer
- 1 Karn's Bastion
- 1 Light of the Legion
- 1 Luminarch Aspirant
- 1 Lyra Dawnbringer
- 1 Minas Tirith
- 1 Nykthos, Shrine to Nyx
- 1 Paragon of Modernity
- 1 Patron of the Valiant
- 1 Righteous Valkyrie
- 1 Ruins of Oran-Rief
- 1 Sanctuary Warden
- 1 Secluded Courtyard
- 1 Seraph Sanctuary
- 1 Serra Redeemer
- 1 Sigarda's Summons
- 1 Three Tree City
- 1 Tyrite Sanctum
- 1 Windbrisk Heights
- 1 Winged Shepherd
- 1 Wojek Investigator
- 1 Youthful Valkyrie
- 12 Plains

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

**The reply:** I removed 1 Angel of Indemnity, 1 Archangel of Thune, 1 Astarion, the Decadent, 1 Cavalier of Night, 1 Celestine, the Living Saint, 1 Cliffhaven Vampire, 1 Defiant Bloodlord, 1 Divinity of Pride, 1 Epicure of Blood, 1 Exalted Sunborn, 1 Fumigate, 1 Sanguine Bond, 1 Sigarda's Splendor, 1 Solitude, 1 The Gaffer, 1 Valkyrie Harbinger, 1 Vein Ripper, 1 Vona, Butcher of Magan, and 1 Wurmcoil Engine. I added 1 Archon of Sun's Grace, 1 Battle Angels of Tyr, 1 Bloodbond Vampire, 1 Convalescent Care, 1 Enduring Tenacity, 1 Essence Pulse, 1 Henrika Domnathi // Henrika, Infernal Seer, 1 Papalymo Totolymo, 1 Qala, Ajani's Pridemate, 1 Sorin of House Markov // Sorin, Ravenous Neonate, 1 Summon: Ixion, 1 Survival Cache, 1 Twinblade Paladin, 1 Vanguard Seraph, 1 Vizkopa Guildmage, and 1 Wax-Wane Witness. I changed the count of Plains: 2 to 4 and Swamp: 1 to 2.

Findings: bracket_prose_rules (INFO), land_count (WARN), curve_summary (INFO), basics_added (INFO)

Repair turn: 1 missed names.

The deck:

- 1 Aclazotz, Deepest Betrayal // Temple of the Dead
- 1 Adventurer's Inn
- 1 Aetherflux Reservoir
- 1 Ajani's Pridemate
- 1 Ajani, Strength of the Pride
- 1 Alseid of Life's Bounty
- 1 Altar of the Pantheon
- 1 Angel of Vitality
- 1 Angelic Accord
- 1 Archivist of Oghma
- 1 Archon of Sun's Grace
- 1 Arguel's Blood Fast // Temple of Aclazotz
- 1 Attended Healer
- 1 Ayli, Eternal Pilgrim
- 1 Azor's Gateway // Sanctum of the Sun
- 1 Battle Angels of Tyr
- 1 Blood Artist
- 1 Bloodbond Vampire
- 1 Bloodthirsty Aerialist
- 1 Bounty Board
- 1 Cavern of Souls
- 1 Caves of Koilos
- 1 City of Brass
- 1 Cleric Class
- 1 Cleric of Life's Bond
- 1 Command Tower
- 1 Convalescent Care
- 1 Cosmos Elixir
- 1 Courageous Resolve
- 1 Dawn of Hope
- 1 Enduring Innocence
- 1 Enduring Tenacity
- 1 Essence Pulse
- 1 Exemplar of Light
- 1 Exotic Orchard
- 1 Faith's Shield
- 1 Gemstone Caverns
- 1 Godless Shrine
- 1 Heliod, Sun-Crowned
- 1 Henrika Domnathi // Henrika, Infernal Seer
- 1 Hierophant's Chalice
- 1 High Market
- 1 Indulging Patrician
- 1 Interplanar Beacon
- 1 Inventors' Fair
- 1 Isolated Chapel
- 1 Kabira Crossroads
- 1 Kaya's Wrath
- 1 Mana Confluence
- 1 Mangara, the Diplomat
- 1 Marauding Blight-Priest
- 1 Metropolis Reformer
- 1 Murderous Rider // Swift End
- 1 Neglected Manor
- 1 Nightmare's Thirst
- 1 Nuka-Cola Vending Machine
- 1 Nykthos, Shrine to Nyx
- 1 Orazca Relic
- 1 Papalymo Totolymo
- 1 Path of Ancestry
- 1 Potioner's Trove
- 1 Pristine Talisman
- 1 Qala, Ajani's Pridemate
- 1 Radiant Fountain
- 1 Reflecting Pool
- 1 Resplendent Angel
- 1 Restless Fortress
- 1 Restoration Magic
- 1 Rhox Faithmender
- 1 Rune-Tail, Kitsune Ascendant // Rune-Tail's Essence
- 1 Scoured Barrens
- 1 Secluded Courtyard
- 1 Seraph Sanctuary
- 1 Shattered Sanctum
- 1 Sorin of House Markov // Sorin, Ravenous Neonate
- 1 Summon: Ixion
- 1 Survival Cache
- 1 Temple of Silence
- 1 The Celestus
- 1 Three Tree City
- 1 Tomb of the Spirit Dragon
- 1 Treasure Chest
- 1 Twinblade Paladin
- 1 Tymna the Weaver
- 1 Unclaimed Territory
- 1 Vanguard Seraph
- 1 Vault of Champions
- 1 Vito, Thorn of the Dusk Rose
- 1 Vizkopa Guildmage
- 1 Voice of the Blessed
- 1 Wax-Wane Witness
- 1 Well of Lost Dreams
- 1 Westvale Abbey // Ormendahl, Profane Prince
- 2 Swamp
- 4 Plains

## Revision 5, base 2: lifegain Commander, Karlov, bracket 3, any card

**The user:** Swap Defiant Bloodlord for something cheaper that does the same job.

The brief:

```json
{
  "changes": [
    "Replace Defiant Bloodlord with a lower-mana-value lifegain payoff that makes opponents lose life when you gain life."
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

**The reply:** I removed 1 Angel of Indemnity and 1 Defiant Bloodlord. I added 1 Angel of Destiny and 1 Enduring Tenacity.

Findings: bracket_prose_rules (INFO), curve_summary (INFO), cards_trimmed (INFO)

The deck:

- 1 Aclazotz, Deepest Betrayal // Temple of the Dead
- 1 Adventurer's Inn
- 1 Aetherflux Reservoir
- 1 Ajani's Pridemate
- 1 Ajani, Strength of the Pride
- 1 Alseid of Life's Bounty
- 1 Altar of the Pantheon
- 1 Angel of Destiny
- 1 Angel of Vitality
- 1 Angelic Accord
- 1 Archangel of Thune
- 1 Archivist of Oghma
- 1 Arguel's Blood Fast // Temple of Aclazotz
- 1 Astarion, the Decadent
- 1 Attended Healer
- 1 Ayli, Eternal Pilgrim
- 1 Azor's Gateway // Sanctum of the Sun
- 1 Blood Artist
- 1 Bloodthirsty Aerialist
- 1 Bounty Board
- 1 Cavalier of Night
- 1 Cavern of Souls
- 1 Caves of Koilos
- 1 Celestine, the Living Saint
- 1 City of Brass
- 1 Cleric Class
- 1 Cleric of Life's Bond
- 1 Cliffhaven Vampire
- 1 Command Tower
- 1 Cosmos Elixir
- 1 Courageous Resolve
- 1 Dawn of Hope
- 1 Divinity of Pride
- 1 Enduring Innocence
- 1 Enduring Tenacity
- 1 Epicure of Blood
- 1 Exalted Sunborn
- 1 Exemplar of Light
- 1 Exotic Orchard
- 1 Faith's Shield
- 1 Fumigate
- 1 Gemstone Caverns
- 1 Godless Shrine
- 1 Heliod, Sun-Crowned
- 1 Hierophant's Chalice
- 1 High Market
- 1 Indulging Patrician
- 1 Interplanar Beacon
- 1 Inventors' Fair
- 1 Isolated Chapel
- 1 Kabira Crossroads
- 1 Kaya's Wrath
- 1 Mana Confluence
- 1 Mangara, the Diplomat
- 1 Marauding Blight-Priest
- 1 Metropolis Reformer
- 1 Murderous Rider // Swift End
- 1 Neglected Manor
- 1 Nightmare's Thirst
- 1 Nuka-Cola Vending Machine
- 1 Nykthos, Shrine to Nyx
- 1 Orazca Relic
- 1 Path of Ancestry
- 1 Potioner's Trove
- 1 Pristine Talisman
- 1 Radiant Fountain
- 1 Reflecting Pool
- 1 Resplendent Angel
- 1 Restless Fortress
- 1 Restoration Magic
- 1 Rhox Faithmender
- 1 Rune-Tail, Kitsune Ascendant // Rune-Tail's Essence
- 1 Sanguine Bond
- 1 Scoured Barrens
- 1 Secluded Courtyard
- 1 Seraph Sanctuary
- 1 Shattered Sanctum
- 1 Sigarda's Splendor
- 1 Solitude
- 1 Swamp
- 1 Temple of Silence
- 1 The Celestus
- 1 The Gaffer
- 1 Three Tree City
- 1 Tomb of the Spirit Dragon
- 1 Treasure Chest
- 1 Tymna the Weaver
- 1 Unclaimed Territory
- 1 Valkyrie Harbinger
- 1 Vault of Champions
- 1 Vein Ripper
- 1 Vito, Thorn of the Dusk Rose
- 1 Voice of the Blessed
- 1 Vona, Butcher of Magan
- 1 Well of Lost Dreams
- 1 Westvale Abbey // Ormendahl, Profane Prince
- 1 Wurmcoil Engine
- 2 Plains

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

**The reply:** I removed 1 Cliffhaven Vampire and 1 Epicure of Blood. I added 1 Convalescent Care and 1 Survival Cache.

Findings: bracket_prose_rules (INFO), curve_summary (INFO)

The deck:

- 1 Aclazotz, Deepest Betrayal // Temple of the Dead
- 1 Adventurer's Inn
- 1 Aetherflux Reservoir
- 1 Ajani's Pridemate
- 1 Ajani, Strength of the Pride
- 1 Alseid of Life's Bounty
- 1 Altar of the Pantheon
- 1 Angel of Indemnity
- 1 Angel of Vitality
- 1 Angelic Accord
- 1 Archangel of Thune
- 1 Archivist of Oghma
- 1 Arguel's Blood Fast // Temple of Aclazotz
- 1 Astarion, the Decadent
- 1 Attended Healer
- 1 Ayli, Eternal Pilgrim
- 1 Azor's Gateway // Sanctum of the Sun
- 1 Blood Artist
- 1 Bloodthirsty Aerialist
- 1 Bounty Board
- 1 Cavalier of Night
- 1 Cavern of Souls
- 1 Caves of Koilos
- 1 Celestine, the Living Saint
- 1 City of Brass
- 1 Cleric Class
- 1 Cleric of Life's Bond
- 1 Command Tower
- 1 Convalescent Care
- 1 Cosmos Elixir
- 1 Courageous Resolve
- 1 Dawn of Hope
- 1 Defiant Bloodlord
- 1 Divinity of Pride
- 1 Enduring Innocence
- 1 Exalted Sunborn
- 1 Exemplar of Light
- 1 Exotic Orchard
- 1 Faith's Shield
- 1 Fumigate
- 1 Gemstone Caverns
- 1 Godless Shrine
- 1 Heliod, Sun-Crowned
- 1 Hierophant's Chalice
- 1 High Market
- 1 Indulging Patrician
- 1 Interplanar Beacon
- 1 Inventors' Fair
- 1 Isolated Chapel
- 1 Kabira Crossroads
- 1 Kaya's Wrath
- 1 Mana Confluence
- 1 Mangara, the Diplomat
- 1 Marauding Blight-Priest
- 1 Metropolis Reformer
- 1 Murderous Rider // Swift End
- 1 Neglected Manor
- 1 Nightmare's Thirst
- 1 Nuka-Cola Vending Machine
- 1 Nykthos, Shrine to Nyx
- 1 Orazca Relic
- 1 Path of Ancestry
- 1 Potioner's Trove
- 1 Pristine Talisman
- 1 Radiant Fountain
- 1 Reflecting Pool
- 1 Resplendent Angel
- 1 Restless Fortress
- 1 Restoration Magic
- 1 Rhox Faithmender
- 1 Rune-Tail, Kitsune Ascendant // Rune-Tail's Essence
- 1 Sanguine Bond
- 1 Scoured Barrens
- 1 Secluded Courtyard
- 1 Seraph Sanctuary
- 1 Shattered Sanctum
- 1 Sigarda's Splendor
- 1 Solitude
- 1 Survival Cache
- 1 Swamp
- 1 Temple of Silence
- 1 The Celestus
- 1 The Gaffer
- 1 Three Tree City
- 1 Tomb of the Spirit Dragon
- 1 Treasure Chest
- 1 Tymna the Weaver
- 1 Unclaimed Territory
- 1 Valkyrie Harbinger
- 1 Vault of Champions
- 1 Vein Ripper
- 1 Vito, Thorn of the Dusk Rose
- 1 Voice of the Blessed
- 1 Vona, Butcher of Magan
- 1 Well of Lost Dreams
- 1 Westvale Abbey // Ormendahl, Profane Prince
- 1 Wurmcoil Engine
- 2 Plains

## Revision 9, base 2: lifegain Commander, Karlov, bracket 3, any card

**The user:** Add better lands instead of the basics. I mean dual lands that enter untapped where you can, and a couple of utility lands.

The brief:

```json
{
  "changes": [
    "Replace the three basic lands with untapped black-white dual lands where possible and a couple of utility lands."
  ],
  "remove": null,
  "keep": null,
  "max_mana_value": 0,
  "swap_basics": 3,
  "land_kinds": "Untapped black-white dual lands and utility lands",
  "question": "",
  "declined": []
}
```

**The reply:** I removed 1 Swamp and 2 Plains. I added 1 Brokers Hideout, 1 Cabaretti Courtyard, and 1 Obscura Storefront.

Findings: bracket_prose_rules (INFO), curve_summary (INFO)

The deck:

- 1 Aclazotz, Deepest Betrayal // Temple of the Dead
- 1 Adventurer's Inn
- 1 Aetherflux Reservoir
- 1 Ajani's Pridemate
- 1 Ajani, Strength of the Pride
- 1 Alseid of Life's Bounty
- 1 Altar of the Pantheon
- 1 Angel of Indemnity
- 1 Angel of Vitality
- 1 Angelic Accord
- 1 Archangel of Thune
- 1 Archivist of Oghma
- 1 Arguel's Blood Fast // Temple of Aclazotz
- 1 Astarion, the Decadent
- 1 Attended Healer
- 1 Ayli, Eternal Pilgrim
- 1 Azor's Gateway // Sanctum of the Sun
- 1 Blood Artist
- 1 Bloodthirsty Aerialist
- 1 Bounty Board
- 1 Brokers Hideout
- 1 Cabaretti Courtyard
- 1 Cavalier of Night
- 1 Cavern of Souls
- 1 Caves of Koilos
- 1 Celestine, the Living Saint
- 1 City of Brass
- 1 Cleric Class
- 1 Cleric of Life's Bond
- 1 Cliffhaven Vampire
- 1 Command Tower
- 1 Cosmos Elixir
- 1 Courageous Resolve
- 1 Dawn of Hope
- 1 Defiant Bloodlord
- 1 Divinity of Pride
- 1 Enduring Innocence
- 1 Epicure of Blood
- 1 Exalted Sunborn
- 1 Exemplar of Light
- 1 Exotic Orchard
- 1 Faith's Shield
- 1 Fumigate
- 1 Gemstone Caverns
- 1 Godless Shrine
- 1 Heliod, Sun-Crowned
- 1 Hierophant's Chalice
- 1 High Market
- 1 Indulging Patrician
- 1 Interplanar Beacon
- 1 Inventors' Fair
- 1 Isolated Chapel
- 1 Kabira Crossroads
- 1 Kaya's Wrath
- 1 Mana Confluence
- 1 Mangara, the Diplomat
- 1 Marauding Blight-Priest
- 1 Metropolis Reformer
- 1 Murderous Rider // Swift End
- 1 Neglected Manor
- 1 Nightmare's Thirst
- 1 Nuka-Cola Vending Machine
- 1 Nykthos, Shrine to Nyx
- 1 Obscura Storefront
- 1 Orazca Relic
- 1 Path of Ancestry
- 1 Potioner's Trove
- 1 Pristine Talisman
- 1 Radiant Fountain
- 1 Reflecting Pool
- 1 Resplendent Angel
- 1 Restless Fortress
- 1 Restoration Magic
- 1 Rhox Faithmender
- 1 Rune-Tail, Kitsune Ascendant // Rune-Tail's Essence
- 1 Sanguine Bond
- 1 Scoured Barrens
- 1 Secluded Courtyard
- 1 Seraph Sanctuary
- 1 Shattered Sanctum
- 1 Sigarda's Splendor
- 1 Solitude
- 1 Temple of Silence
- 1 The Celestus
- 1 The Gaffer
- 1 Three Tree City
- 1 Tomb of the Spirit Dragon
- 1 Treasure Chest
- 1 Tymna the Weaver
- 1 Unclaimed Territory
- 1 Valkyrie Harbinger
- 1 Vault of Champions
- 1 Vein Ripper
- 1 Vito, Thorn of the Dusk Rose
- 1 Voice of the Blessed
- 1 Vona, Butcher of Magan
- 1 Well of Lost Dreams
- 1 Westvale Abbey // Ormendahl, Profane Prince
- 1 Wurmcoil Engine

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
  "swap_basics": 0,
  "land_kinds": "",
  "question": "",
  "declined": []
}
```

**The reply:** I removed 1 Bag End Banquet, 1 Bilbo's Burglaring, 1 Cantankerous Keepers, 1 Gigantic Big Bear, 1 Mirkwood Elk, 1 The Lord of the Eagles, 1 Thranduil's Decree, 1 Troll of Khazad-dûm, and 1 Willow-Wind. I added 1 Bilbo, Fellow Conspirator, 1 Chief of the Wilds, 1 Desolation Prowler, 1 Dreaded Bat-Cloud, 1 Duskwatch Hunter, 1 Elrond, Moon-Reader, 1 Galion, Elvenking's Butler, 1 Gollum the Abandoned, and 1 Guardian of the Halls.

Findings: bracket_prose_rules (INFO), curve_summary (INFO)

The deck:

- 1 Along the Crooked Way
- 1 Arcane Signet
- 1 Arwen, Weaver of Hope
- 1 Attercop
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
- 1 Crude Bent Blade
- 1 Delighted Halfling
- 1 Desolation Prowler
- 1 Down in the Valley
- 1 Dragon-Cursed Halls
- 1 Dreaded Bat-Cloud
- 1 Duskwatch Hunter
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
- 1 Giant's Boulder
- 1 Gnashing of Teeth
- 1 Gollum the Abandoned
- 1 Gollum, Silent Slinker // Meager Meal
- 1 Great Fierce Bee
- 1 Guardian of the Halls
- 1 Hithlain Knots
- 1 Hobbit Hole
- 1 Languish
- 1 Large Bear
- 1 Minas Morgul, Dark Fortress
- 1 Mirkwood
- 1 Mirkwood Pathmaker
- 1 Mithril Coat
- 1 Mox Amber
- 1 My Precious // Allure of Power
- 1 Nasty Little Rabbit
- 1 Necklace of Girion
- 1 Night's Whisper
- 1 Old Fat Spider Can't See Me
- 1 Orcish Bowmasters
- 1 Ordinary Bear
- 1 Palantír of Orthanc
- 1 Part in Friendship
- 1 Plunder the Trollshaws
- 1 Quarrel
- 1 Raise the Palisade
- 1 Rivendell
- 1 Rogue's Passage
- 1 Roll-Roll-Roll-Roll
- 1 Sting, Bilbo's Sword
- 1 Supper for Spiders
- 1 The Black Gate
- 1 The Notary Hobbits
- 1 The Sackville-Bagginses
- 1 The Shire
- 1 Thrór's Map
- 1 Treasure Vault
- 1 Uncover the Moon-Letters
- 1 Uneasy Partings
- 1 Wayfarer's Bauble
- 1 Well-Worn Spatula
- 1 Wood Elves
- 15 Forest
- 4 Swamp
- 6 Island

## Revision 8, base 3: Commander from the Hobbit set, bracket 3

**The user:** Take out The Lord of the Eagles.

The brief:

```json
{
  "changes": [
    "Remove The Lord of the Eagles."
  ],
  "remove": [
    "The Lord of the Eagles"
  ],
  "keep": null,
  "max_mana_value": 0,
  "swap_basics": 0,
  "land_kinds": "",
  "question": "",
  "declined": []
}
```

**The reply:** I removed 1 The Lord of the Eagles. I changed the count of Forest: 15 to 16.

Findings: bracket_prose_rules (INFO), curve_summary (INFO), basics_added (INFO)

The deck:

- 1 Along the Crooked Way
- 1 Arcane Signet
- 1 Arwen, Weaver of Hope
- 1 Attercop
- 1 Bag End Banquet
- 1 Beorn's Hospitality
- 1 Bilbo Baggins, Burglar // Take a Glance
- 1 Bilbo's Burglaring
- 1 Bilbo's Deadly Slice
- 1 Bilbo's Ring
- 1 Bilbo, Luckwearer // Burglar's Plot
- 1 Bilbo, Thief in the Night
- 1 Bitter Downfall
- 1 Cantankerous Keepers
- 1 Confusticate and Bebother
- 1 Crude Bent Blade
- 1 Delighted Halfling
- 1 Down in the Valley
- 1 Dragon-Cursed Halls
- 1 Elven Chorus
- 1 Elven Passage
- 1 Elvenking's Halls
- 1 Elvish Archdruid
- 1 Elvish Mystic
- 1 Elvish Visionary
- 1 Enchanted River's Grasp
- 1 Fateful Discovery
- 1 Giant's Boulder
- 1 Gigantic Big Bear
- 1 Gnashing of Teeth
- 1 Gollum, Silent Slinker // Meager Meal
- 1 Great Fierce Bee
- 1 Hithlain Knots
- 1 Hobbit Hole
- 1 Languish
- 1 Large Bear
- 1 Minas Morgul, Dark Fortress
- 1 Mirkwood
- 1 Mirkwood Elk
- 1 Mirkwood Pathmaker
- 1 Mithril Coat
- 1 Mox Amber
- 1 My Precious // Allure of Power
- 1 Nasty Little Rabbit
- 1 Necklace of Girion
- 1 Night's Whisper
- 1 Old Fat Spider Can't See Me
- 1 Orcish Bowmasters
- 1 Ordinary Bear
- 1 Palantír of Orthanc
- 1 Part in Friendship
- 1 Plunder the Trollshaws
- 1 Quarrel
- 1 Raise the Palisade
- 1 Rivendell
- 1 Rogue's Passage
- 1 Roll-Roll-Roll-Roll
- 1 Sting, Bilbo's Sword
- 1 Supper for Spiders
- 1 The Black Gate
- 1 The Notary Hobbits
- 1 The Sackville-Bagginses
- 1 The Shire
- 1 Thranduil's Decree
- 1 Thrór's Map
- 1 Treasure Vault
- 1 Troll of Khazad-dûm
- 1 Uncover the Moon-Letters
- 1 Uneasy Partings
- 1 Wayfarer's Bauble
- 1 Well-Worn Spatula
- 1 Willow-Wind
- 1 Wood Elves
- 16 Forest
- 4 Swamp
- 6 Island
