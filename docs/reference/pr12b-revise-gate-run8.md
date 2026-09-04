# PR-12B revise gate

Run date: 2026-09-04. Snapshot: 2026-09-03.

Verdict: PASS. 11 of 11 turns met their bar. The bars: an unclear request gets a question or a decline with a reason, an answer to that question gets a rebuild, a clear request gets a revised deck with no block finding, the deck holds every cap, every removal, and every land swap the message names, and it keeps at least 80 percent of the untouched cards.

Cost: 33 calls, $1.2328. Time: 15m46s.

| Base | Revision | Message | Outcome | Kept | Result |
|---|---|---|---|---|---|
| 1 | 1 | Replace some lands with better options if possible. Also tune the mana curve lower - no 6 or 7 mana cards needed | question |  | PASS, no build |
| 1 | 1 | (the answer) Q: For the land replacements, do you want faster mana, utility lands, or to add access to more colors? A: A mix. | 60 cards | 100% | PASS, kept 100% |
| 1 | 2 | Take out Angel of the Ruins. | 60 cards | 100% | PASS, kept 100% |
| 1 | 3 | Better lands please. | question |  | PASS, no build |
| 1 | 3 | (the answer) Q: When you say better lands, do you want faster mana, utility lands, or lands that add more colors? A: Utility lands. | 60 cards | 100% | PASS, kept 100% |
| 2 | 4 | No nonland cards over 4 mana. | 99 cards | 97% | PASS, kept 97% |
| 2 | 5 | Swap Sephara, Sky's Blade for something cheaper that does the same job. | 99 cards | 99% | PASS, kept 99% |
| 2 | 6 | Add two more cards that draw cards, and cut two of the weakest creatures. | 99 cards | 97% | PASS, kept 97% |
| 2 | 9 | Add better lands instead of the basics. I mean dual lands that enter untapped where you can, and a couple of utility lands. | 99 cards | 100% | PASS, kept 100% |
| 3 | 7 | Lower the mana curve. Nothing over 5 mana. | 99 cards | 100% | PASS, kept 100% |
| 3 | 8 | Take out Bag End Banquet. | 99 cards | 100% | PASS, kept 100% |

## Revision 1, base 1: Modern Angels, casual, mono-white, any card

**The user:** Replace some lands with better options if possible. Also tune the mana curve lower - no 6 or 7 mana cards needed

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
  "question": "For the land replacements, do you want faster mana, utility lands, or to add access to more colors?",
  "declined": []
}
```

**The reply:** Question: For the land replacements, do you want faster mana, utility lands, or to add access to more colors?


## Revision 1, the answer, base 1: Modern Angels, casual, mono-white, any card

**The user:** Q: For the land replacements, do you want faster mana, utility lands, or to add access to more colors? A: A mix.

The brief:

```json
{
  "changes": [
    "Replace 4 Plains with a mix of faster mana lands and utility lands.",
    "Remove every nonland card with mana value 6 or more."
  ],
  "remove": null,
  "keep": null,
  "max_mana_value": 5,
  "swap_basics": 4,
  "land_kinds": "a mix of faster mana lands and utility lands",
  "question": "",
  "declined": []
}
```

**The reply:** I removed 1 Admonition Angel, 1 Angel of the Ruins, 1 Sanctuary Warden, 1 Winged Shepherd, and 2 Light of the Legion. I added 1 Eiganjo, Seat of the Empire, 1 Gemstone Caverns, 1 Karn's Bastion, 1 Seraph Sanctuary, 2 Paragon of Modernity, and 2 Wojek Investigator. I changed the count of Angel of Condemnation: 1 to 3 and Plains: 24 to 20.

Findings: curve_summary (INFO)

The deck:

- 1 Aurelia's Vindicator
- 1 Bishop of Wings
- 1 Cathars' Crusade
- 1 Eiganjo, Seat of the Empire
- 1 Gemstone Caverns
- 1 Karn's Bastion
- 1 Seraph Sanctuary
- 2 Angel of Invention
- 2 Angel of Sanctions
- 2 Archangel of Thune
- 2 Exemplar of Light
- 2 Giada, Font of Hope
- 2 Herald of War
- 2 Inspiring Overseer
- 2 Lyra Dawnbringer
- 2 Paragon of Modernity
- 2 Patron of the Valiant
- 2 Righteous Valkyrie
- 2 Serra Redeemer
- 2 Wojek Investigator
- 20 Plains
- 3 Angel of Condemnation
- 4 Youthful Valkyrie

## Revision 2, base 1: Modern Angels, casual, mono-white, any card

**The user:** Take out Angel of the Ruins.

The brief:

```json
{
  "changes": [
    "Remove Angel of the Ruins."
  ],
  "remove": [
    "Angel of the Ruins"
  ],
  "keep": null,
  "max_mana_value": 0,
  "swap_basics": 0,
  "land_kinds": "",
  "question": "",
  "declined": []
}
```

**The reply:** I removed 1 Angel of the Ruins. I added 1 Angel of Deliverance.

Findings: curve_summary (INFO)

The deck:

- 1 Admonition Angel
- 1 Angel of Condemnation
- 1 Angel of Deliverance
- 1 Aurelia's Vindicator
- 1 Bishop of Wings
- 1 Cathars' Crusade
- 1 Sanctuary Warden
- 1 Winged Shepherd
- 2 Angel of Invention
- 2 Angel of Sanctions
- 2 Archangel of Thune
- 2 Exemplar of Light
- 2 Giada, Font of Hope
- 2 Herald of War
- 2 Inspiring Overseer
- 2 Light of the Legion
- 2 Lyra Dawnbringer
- 2 Patron of the Valiant
- 2 Righteous Valkyrie
- 2 Serra Redeemer
- 24 Plains
- 4 Youthful Valkyrie

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
  "question": "When you say better lands, do you want faster mana, utility lands, or lands that add more colors?",
  "declined": []
}
```

**The reply:** Question: When you say better lands, do you want faster mana, utility lands, or lands that add more colors?


## Revision 3, the answer, base 1: Modern Angels, casual, mono-white, any card

**The user:** Q: When you say better lands, do you want faster mana, utility lands, or lands that add more colors? A: Utility lands.

The brief:

```json
{
  "changes": [
    "Replace a few basic Plains with utility lands."
  ],
  "remove": null,
  "keep": null,
  "max_mana_value": 0,
  "swap_basics": 3,
  "land_kinds": "Utility lands",
  "question": "",
  "declined": []
}
```

**The reply:** I added 1 Castle Ardenvale, 1 Eiganjo, Seat of the Empire, and 1 Seraph Sanctuary. I changed the count of Plains: 24 to 21.

Findings: curve_summary (INFO)

The deck:

- 1 Admonition Angel
- 1 Angel of Condemnation
- 1 Angel of the Ruins
- 1 Aurelia's Vindicator
- 1 Bishop of Wings
- 1 Castle Ardenvale
- 1 Cathars' Crusade
- 1 Eiganjo, Seat of the Empire
- 1 Sanctuary Warden
- 1 Seraph Sanctuary
- 1 Winged Shepherd
- 2 Angel of Invention
- 2 Angel of Sanctions
- 2 Archangel of Thune
- 2 Exemplar of Light
- 2 Giada, Font of Hope
- 2 Herald of War
- 2 Inspiring Overseer
- 2 Light of the Legion
- 2 Lyra Dawnbringer
- 2 Patron of the Valiant
- 2 Righteous Valkyrie
- 2 Serra Redeemer
- 21 Plains
- 4 Youthful Valkyrie

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

**The reply:** I removed 1 Archangel of Thune, 1 Celestine, the Living Saint, 1 Fumigate, 1 Gideon's Company, 1 Nykthos Paragon, 1 Sephara, Sky's Blade, 1 Starseer Mentor, 1 Summon: Ixion, 1 Valkyrie Harbinger, and 1 Vampire Scrivener. I added 1 Elenda, Saint of Dusk, 1 Gisela, the Broken Blade, 1 Papalymo Totolymo, 1 Qala, Ajani's Pridemate, 1 Rhox Faithmender, 1 Sorin of House Markov // Sorin, Ravenous Neonate, 1 Sword of Light and Shadow, 1 Vanguard Seraph, and 1 White Sun's Twilight. I changed the count of Plains: 13 to 14.

Findings: curve_summary (INFO), basics_added (INFO)

Repair turn: 2 missed names, profile_off_band.

The deck:

- 1 Aerith Gainsborough
- 1 Aetherflux Reservoir
- 1 Ajani's Pridemate
- 1 Alseid of Life's Bounty
- 1 Altar of the Pantheon
- 1 Angel of Vitality
- 1 Archivist of Oghma
- 1 Attended Healer
- 1 Ayli, Eternal Pilgrim
- 1 Bloodbond Vampire
- 1 Bloodthirsty Aerialist
- 1 Bounty Board
- 1 Caves of Koilos
- 1 City of Brass
- 1 Cleric Class
- 1 Cleric of Life's Bond
- 1 Cliffhaven Vampire
- 1 Command Tower
- 1 Courageous Resolve
- 1 Cryptolith Fragment // Aurora of Emrakul
- 1 Dawn of Hope
- 1 Elenda, Saint of Dusk
- 1 Enduring Innocence
- 1 Enduring Tenacity
- 1 Exemplar of Light
- 1 Exotic Orchard
- 1 Faith's Shield
- 1 Frodo, Adventurous Hobbit
- 1 Gemstone Caverns
- 1 Gisela, the Broken Blade
- 1 Godless Shrine
- 1 Gumdrop Poisoner // Tempt with Treats
- 1 Heliod, Sun-Crowned
- 1 High Market
- 1 Hot Dog Cart
- 1 Kaya's Wrath
- 1 Mana Confluence
- 1 Mangara, the Diplomat
- 1 Markov Purifier
- 1 Metropolis Reformer
- 1 Murderous Rider // Swift End
- 1 Nightmare's Thirst
- 1 Nuka-Cola Vending Machine
- 1 Ocelot Pride
- 1 Orazca Relic
- 1 Papalymo Totolymo
- 1 Phial of Galadriel
- 1 Poisoner's Apprentice
- 1 Potioner's Trove
- 1 Pristine Talisman
- 1 Qala, Ajani's Pridemate
- 1 Reflecting Pool
- 1 Resplendent Angel
- 1 Restless Fortress
- 1 Restoration Magic
- 1 Rhox Faithmender
- 1 Rune-Tail, Kitsune Ascendant // Rune-Tail's Essence
- 1 Shattered Sanctum
- 1 Sigarda's Splendor
- 1 Sorin of House Markov // Sorin, Ravenous Neonate
- 1 Sorin, Imperious Bloodlord
- 1 Sword of Light and Shadow
- 1 The Celestus
- 1 The Gaffer
- 1 The Meathook Massacre
- 1 Twinblade Paladin
- 1 Tymna the Weaver
- 1 Vanguard Seraph
- 1 Vault of Champions
- 1 Vito, Thorn of the Dusk Rose
- 1 Wax-Wane Witness
- 1 Well of Lost Dreams
- 1 Werefox Bodyguard
- 1 White Sun's Twilight
- 11 Swamp
- 14 Plains

## Revision 5, base 2: lifegain Commander, Karlov, bracket 3, any card

**The user:** Swap Sephara, Sky's Blade for something cheaper that does the same job.

The brief:

```json
{
  "changes": [
    "Replace Sephara, Sky's Blade with a lower-mana-value card that protects the creature board in a similar way."
  ],
  "remove": [
    "Sephara, Sky's Blade"
  ],
  "keep": null,
  "max_mana_value": 0,
  "swap_basics": 0,
  "land_kinds": "",
  "question": "",
  "declined": []
}
```

**The reply:** I removed 1 Enduring Tenacity and 1 Sephara, Sky's Blade. I added 1 Archon of Sun's Grace and 1 Bofur, Reliable Guardian // Concerted Care.

Findings: curve_summary (INFO)

Repair turn: two_card_combo.

The deck:

- 1 Aerith Gainsborough
- 1 Aetherflux Reservoir
- 1 Ajani's Pridemate
- 1 Alseid of Life's Bounty
- 1 Altar of the Pantheon
- 1 Angel of Vitality
- 1 Archangel of Thune
- 1 Archivist of Oghma
- 1 Archon of Sun's Grace
- 1 Attended Healer
- 1 Ayli, Eternal Pilgrim
- 1 Bloodbond Vampire
- 1 Bloodthirsty Aerialist
- 1 Bofur, Reliable Guardian // Concerted Care
- 1 Bounty Board
- 1 Caves of Koilos
- 1 Celestine, the Living Saint
- 1 City of Brass
- 1 Cleric Class
- 1 Cleric of Life's Bond
- 1 Cliffhaven Vampire
- 1 Command Tower
- 1 Courageous Resolve
- 1 Cryptolith Fragment // Aurora of Emrakul
- 1 Dawn of Hope
- 1 Enduring Innocence
- 1 Exemplar of Light
- 1 Exotic Orchard
- 1 Faith's Shield
- 1 Frodo, Adventurous Hobbit
- 1 Fumigate
- 1 Gemstone Caverns
- 1 Gideon's Company
- 1 Godless Shrine
- 1 Gumdrop Poisoner // Tempt with Treats
- 1 Heliod, Sun-Crowned
- 1 High Market
- 1 Hot Dog Cart
- 1 Kaya's Wrath
- 1 Mana Confluence
- 1 Mangara, the Diplomat
- 1 Markov Purifier
- 1 Metropolis Reformer
- 1 Murderous Rider // Swift End
- 1 Nightmare's Thirst
- 1 Nuka-Cola Vending Machine
- 1 Nykthos Paragon
- 1 Ocelot Pride
- 1 Orazca Relic
- 1 Phial of Galadriel
- 1 Poisoner's Apprentice
- 1 Potioner's Trove
- 1 Pristine Talisman
- 1 Reflecting Pool
- 1 Resplendent Angel
- 1 Restless Fortress
- 1 Restoration Magic
- 1 Rune-Tail, Kitsune Ascendant // Rune-Tail's Essence
- 1 Shattered Sanctum
- 1 Sigarda's Splendor
- 1 Sorin, Imperious Bloodlord
- 1 Starseer Mentor
- 1 Summon: Ixion
- 1 The Celestus
- 1 The Gaffer
- 1 The Meathook Massacre
- 1 Twinblade Paladin
- 1 Tymna the Weaver
- 1 Valkyrie Harbinger
- 1 Vampire Scrivener
- 1 Vault of Champions
- 1 Vito, Thorn of the Dusk Rose
- 1 Wax-Wane Witness
- 1 Well of Lost Dreams
- 1 Werefox Bodyguard
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

**The reply:** I removed 1 Bloodbond Vampire and 1 Gideon's Company. I added 1 Cosmos Elixir and 1 Survival Cache.

Findings: curve_summary (INFO)

The deck:

- 1 Aerith Gainsborough
- 1 Aetherflux Reservoir
- 1 Ajani's Pridemate
- 1 Alseid of Life's Bounty
- 1 Altar of the Pantheon
- 1 Angel of Vitality
- 1 Archangel of Thune
- 1 Archivist of Oghma
- 1 Attended Healer
- 1 Ayli, Eternal Pilgrim
- 1 Bloodthirsty Aerialist
- 1 Bounty Board
- 1 Caves of Koilos
- 1 Celestine, the Living Saint
- 1 City of Brass
- 1 Cleric Class
- 1 Cleric of Life's Bond
- 1 Cliffhaven Vampire
- 1 Command Tower
- 1 Cosmos Elixir
- 1 Courageous Resolve
- 1 Cryptolith Fragment // Aurora of Emrakul
- 1 Dawn of Hope
- 1 Enduring Innocence
- 1 Enduring Tenacity
- 1 Exemplar of Light
- 1 Exotic Orchard
- 1 Faith's Shield
- 1 Frodo, Adventurous Hobbit
- 1 Fumigate
- 1 Gemstone Caverns
- 1 Godless Shrine
- 1 Gumdrop Poisoner // Tempt with Treats
- 1 Heliod, Sun-Crowned
- 1 High Market
- 1 Hot Dog Cart
- 1 Kaya's Wrath
- 1 Mana Confluence
- 1 Mangara, the Diplomat
- 1 Markov Purifier
- 1 Metropolis Reformer
- 1 Murderous Rider // Swift End
- 1 Nightmare's Thirst
- 1 Nuka-Cola Vending Machine
- 1 Nykthos Paragon
- 1 Ocelot Pride
- 1 Orazca Relic
- 1 Phial of Galadriel
- 1 Poisoner's Apprentice
- 1 Potioner's Trove
- 1 Pristine Talisman
- 1 Reflecting Pool
- 1 Resplendent Angel
- 1 Restless Fortress
- 1 Restoration Magic
- 1 Rune-Tail, Kitsune Ascendant // Rune-Tail's Essence
- 1 Sephara, Sky's Blade
- 1 Shattered Sanctum
- 1 Sigarda's Splendor
- 1 Sorin, Imperious Bloodlord
- 1 Starseer Mentor
- 1 Summon: Ixion
- 1 Survival Cache
- 1 The Celestus
- 1 The Gaffer
- 1 The Meathook Massacre
- 1 Twinblade Paladin
- 1 Tymna the Weaver
- 1 Valkyrie Harbinger
- 1 Vampire Scrivener
- 1 Vault of Champions
- 1 Vito, Thorn of the Dusk Rose
- 1 Wax-Wane Witness
- 1 Well of Lost Dreams
- 1 Werefox Bodyguard
- 11 Swamp
- 13 Plains

## Revision 9, base 2: lifegain Commander, Karlov, bracket 3, any card

**The user:** Add better lands instead of the basics. I mean dual lands that enter untapped where you can, and a couple of utility lands.

The brief:

```json
{
  "changes": [
    "Replace basic lands with untapped Orzhov dual lands where possible and a couple of utility lands."
  ],
  "remove": null,
  "keep": null,
  "max_mana_value": 0,
  "swap_basics": 22,
  "land_kinds": "untapped white-black dual lands and a couple of utility lands",
  "question": "",
  "declined": []
}
```

**The reply:** I added 1 Aclazotz, Deepest Betrayal // Temple of the Dead, 1 Adventurer's Inn, 1 Arguel's Blood Fast // Temple of Aclazotz, 1 Azor's Gateway // Sanctum of the Sun, 1 Cavern of Souls, 1 Interplanar Beacon, 1 Inventors' Fair, 1 Isolated Chapel, 1 Kabira Crossroads, 1 Neglected Manor, 1 Nykthos, Shrine to Nyx, 1 Obscura Storefront, 1 Path of Ancestry, 1 Radiant Fountain, 1 Scoured Barrens, 1 Secluded Courtyard, 1 Seraph Sanctuary, 1 Temple of Silence, 1 Three Tree City, 1 Tomb of the Spirit Dragon, 1 Unclaimed Territory, and 1 Westvale Abbey // Ormendahl, Profane Prince. I changed the count of Plains: 13 to 1 and Swamp: 11 to 1.

Findings: curve_summary (INFO), profile_off_band (WARN), profile_off_band (WARN), repair_kept_earlier (INFO)

Repair turn: profile_off_band, profile_off_band.

The deck:

- 1 Aclazotz, Deepest Betrayal // Temple of the Dead
- 1 Adventurer's Inn
- 1 Aerith Gainsborough
- 1 Aetherflux Reservoir
- 1 Ajani's Pridemate
- 1 Alseid of Life's Bounty
- 1 Altar of the Pantheon
- 1 Angel of Vitality
- 1 Archangel of Thune
- 1 Archivist of Oghma
- 1 Arguel's Blood Fast // Temple of Aclazotz
- 1 Attended Healer
- 1 Ayli, Eternal Pilgrim
- 1 Azor's Gateway // Sanctum of the Sun
- 1 Bloodbond Vampire
- 1 Bloodthirsty Aerialist
- 1 Bounty Board
- 1 Cavern of Souls
- 1 Caves of Koilos
- 1 Celestine, the Living Saint
- 1 City of Brass
- 1 Cleric Class
- 1 Cleric of Life's Bond
- 1 Cliffhaven Vampire
- 1 Command Tower
- 1 Courageous Resolve
- 1 Cryptolith Fragment // Aurora of Emrakul
- 1 Dawn of Hope
- 1 Enduring Innocence
- 1 Enduring Tenacity
- 1 Exemplar of Light
- 1 Exotic Orchard
- 1 Faith's Shield
- 1 Frodo, Adventurous Hobbit
- 1 Fumigate
- 1 Gemstone Caverns
- 1 Gideon's Company
- 1 Godless Shrine
- 1 Gumdrop Poisoner // Tempt with Treats
- 1 Heliod, Sun-Crowned
- 1 High Market
- 1 Hot Dog Cart
- 1 Interplanar Beacon
- 1 Inventors' Fair
- 1 Isolated Chapel
- 1 Kabira Crossroads
- 1 Kaya's Wrath
- 1 Mana Confluence
- 1 Mangara, the Diplomat
- 1 Markov Purifier
- 1 Metropolis Reformer
- 1 Murderous Rider // Swift End
- 1 Neglected Manor
- 1 Nightmare's Thirst
- 1 Nuka-Cola Vending Machine
- 1 Nykthos Paragon
- 1 Nykthos, Shrine to Nyx
- 1 Obscura Storefront
- 1 Ocelot Pride
- 1 Orazca Relic
- 1 Path of Ancestry
- 1 Phial of Galadriel
- 1 Plains
- 1 Poisoner's Apprentice
- 1 Potioner's Trove
- 1 Pristine Talisman
- 1 Radiant Fountain
- 1 Reflecting Pool
- 1 Resplendent Angel
- 1 Restless Fortress
- 1 Restoration Magic
- 1 Rune-Tail, Kitsune Ascendant // Rune-Tail's Essence
- 1 Scoured Barrens
- 1 Secluded Courtyard
- 1 Sephara, Sky's Blade
- 1 Seraph Sanctuary
- 1 Shattered Sanctum
- 1 Sigarda's Splendor
- 1 Sorin, Imperious Bloodlord
- 1 Starseer Mentor
- 1 Summon: Ixion
- 1 Swamp
- 1 Temple of Silence
- 1 The Celestus
- 1 The Gaffer
- 1 The Meathook Massacre
- 1 Three Tree City
- 1 Tomb of the Spirit Dragon
- 1 Twinblade Paladin
- 1 Tymna the Weaver
- 1 Unclaimed Territory
- 1 Valkyrie Harbinger
- 1 Vampire Scrivener
- 1 Vault of Champions
- 1 Vito, Thorn of the Dusk Rose
- 1 Wax-Wane Witness
- 1 Well of Lost Dreams
- 1 Werefox Bodyguard
- 1 Westvale Abbey // Ormendahl, Profane Prince

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

**The reply:** I removed 1 Bag End Banquet, 1 Boughside Wanderers, 1 Cantankerous Keepers, 1 Mirkwood Elk, and 1 Thranduil's Decree. I added 1 Attercop, 1 Gollum, Riddle Master, 1 Great Fierce Bee, 1 The One Ring, and 1 Woodland Weavemaster. I changed the count of Forest: 6 to 9 and Island: 12 to 9.

Findings: curve_summary (INFO)

Repair turn: 1 missed names; then profile_off_band.

The deck:

- 1 Along the Crooked Way
- 1 Arcane Signet
- 1 Arwen, Weaver of Hope
- 1 Attercop
- 1 Beorn the Fierce
- 1 Beorn's Hospitality
- 1 Bilbo Baggins, Burglar // Take a Glance
- 1 Bilbo's Deadly Slice
- 1 Bilbo's Ring
- 1 Bilbo, Fellow Conspirator
- 1 Bilbo, Luckwearer // Burglar's Plot
- 1 Bilbo, Thief in the Night
- 1 Bitter Downfall
- 1 Confusticate and Bebother
- 1 Crude Bent Blade
- 1 Delighted Halfling
- 1 Down in the Valley
- 1 Dwarven Mattock
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
- 1 Galion, Elvenking's Butler
- 1 Giant's Boulder
- 1 Gnashing of Teeth
- 1 Gollum the Abandoned
- 1 Gollum, Riddle Master
- 1 Gollum, Silent Slinker // Meager Meal
- 1 Great Fierce Bee
- 1 Haunt of the Dead Marshes
- 1 Hithlain Knots
- 1 Hobbit Hole
- 1 Languish
- 1 Lórien Revealed
- 1 Merciless Executioner
- 1 Minas Morgul, Dark Fortress
- 1 Mirkwood
- 1 Mirkwood Pathmaker
- 1 Mithril Coat
- 1 Necklace of Girion
- 1 Night's Whisper
- 1 Nimrodel Watcher
- 1 Old Fat Spider Can't See Me
- 1 Orcish Bowmasters
- 1 Palantír of Orthanc
- 1 Part in Friendship
- 1 Pelargir Survivor
- 1 Plunder the Trollshaws
- 1 Quarrel
- 1 Raise the Palisade
- 1 Rivendell
- 1 Silvan Reveler
- 1 Stern Scolding
- 1 Supper for Spiders
- 1 The Black Gate
- 1 The Notary Hobbits
- 1 The One Ring
- 1 The Sackville-Bagginses
- 1 The Shire
- 1 Thranduil, Sindarin Liege // Silvan Rally
- 1 Treasure Vault
- 1 Uncover the Moon-Letters
- 1 Uneasy Partings
- 1 Willow-Wind
- 1 Wood Elves
- 1 Woodland Weavemaster
- 9 Forest
- 9 Island
- 9 Swamp

## Revision 8, base 3: Commander from the Hobbit set, bracket 3

**The user:** Take out Bag End Banquet.

The brief:

```json
{
  "changes": [
    "Remove Bag End Banquet."
  ],
  "remove": [
    "Bag End Banquet"
  ],
  "keep": null,
  "max_mana_value": 0,
  "swap_basics": 0,
  "land_kinds": "",
  "question": "",
  "declined": []
}
```

**The reply:** I removed 1 Bag End Banquet. I changed the count of Forest: 6 to 8 and Island: 12 to 11.

Findings: curve_summary (INFO)

Repair turn: profile_off_band.

The deck:

- 1 Along the Crooked Way
- 1 Arcane Signet
- 1 Arwen, Weaver of Hope
- 1 Beorn the Fierce
- 1 Beorn's Hospitality
- 1 Bilbo Baggins, Burglar // Take a Glance
- 1 Bilbo's Deadly Slice
- 1 Bilbo's Ring
- 1 Bilbo, Fellow Conspirator
- 1 Bilbo, Luckwearer // Burglar's Plot
- 1 Bilbo, Thief in the Night
- 1 Bitter Downfall
- 1 Boughside Wanderers
- 1 Cantankerous Keepers
- 1 Confusticate and Bebother
- 1 Crude Bent Blade
- 1 Delighted Halfling
- 1 Down in the Valley
- 1 Dwarven Mattock
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
- 1 Galion, Elvenking's Butler
- 1 Giant's Boulder
- 1 Gnashing of Teeth
- 1 Gollum the Abandoned
- 1 Gollum, Silent Slinker // Meager Meal
- 1 Haunt of the Dead Marshes
- 1 Hithlain Knots
- 1 Hobbit Hole
- 1 Languish
- 1 Lórien Revealed
- 1 Merciless Executioner
- 1 Minas Morgul, Dark Fortress
- 1 Mirkwood
- 1 Mirkwood Elk
- 1 Mirkwood Pathmaker
- 1 Mithril Coat
- 1 Necklace of Girion
- 1 Night's Whisper
- 1 Nimrodel Watcher
- 1 Old Fat Spider Can't See Me
- 1 Orcish Bowmasters
- 1 Palantír of Orthanc
- 1 Part in Friendship
- 1 Pelargir Survivor
- 1 Plunder the Trollshaws
- 1 Quarrel
- 1 Raise the Palisade
- 1 Rivendell
- 1 Silvan Reveler
- 1 Stern Scolding
- 1 Supper for Spiders
- 1 The Black Gate
- 1 The Notary Hobbits
- 1 The Sackville-Bagginses
- 1 The Shire
- 1 Thranduil's Decree
- 1 Thranduil, Sindarin Liege // Silvan Rally
- 1 Treasure Vault
- 1 Uncover the Moon-Letters
- 1 Uneasy Partings
- 1 Willow-Wind
- 1 Wood Elves
- 11 Island
- 8 Forest
- 9 Swamp
