# The review of the Gríma deck, 2026-09-14

The owner gave a written review of deck `sFLbEUuKzI0QyLPft0zI` of session `z1hshyY6Npig1FN2NuV7` on 2026-09-14. This document checks each claim of that review. The sources are the stored deck, the card snapshot of 2026-09-04, the stored collection of the owner, and the code of `main` at `54b07dd`. The owner answered the plan in D-719 to D-722.

A card name that holds a comma sits in quotation marks, so each card reads once in a list.

## The deck

- The commander is "Gríma, Saruman's Footman", at bracket 5, from an owned-only pool, on the theme "opponent milling cards".
- The 99 cards sit in 75 rows: 33 lands, 30 noncreature artifacts, 23 instants, 7 sorceries, 3 enchantments, and 3 nonland creatures. One of the creatures, "Iron Spider, Stark Upgrade", is also an artifact.
- The model wrote a role on each row: ramp on 19 rows and interaction on 19. Draw sits on 18 rows, removal and land on 9 each, and wipe on 1. No row carries the role wincon or threat.
- The deck holds 0 of 4 tutors, 3 of 6 fast mana, and 1 of 8 Game Changers against the floors of bracket 5.
- The lands are 13 Island, 13 Swamp, Command Tower, City of Brass, Exotic Orchard, Plaza of Heroes, Spire of Industry, Thriving Isle, and Thriving Moor.

## The claims

| Claim of the review | Verdict | Evidence |
|---|---|---|
| The deck has no win condition. | Confirmed (F-138) | No row carries the role wincon or threat. One card carries a finisher tag: Grave Venerations, tagged drain-life. |
| The mana base is weak. | Confirmed (F-139) | No land is an untapped dual land. The lands give about 19 blue sources, and the Karsten table of the code asks 33 for {1}{U}{U}{U} in 99 cards. The review said 23 to 24. |
| Several cards miss their conditions. | Confirmed (F-140) | Skullclamp, Springleaf Drum, Idol of Oblivion, and Kindred Discovery sit beside 3 nonland creatures. |
| Eight cards pump a creature for a trigger that fires once. | Mostly refuted | Blitzball, Bender's Waterskin, Hot Dog Cart, and Interdimensional Web Watch make mana. Explorer's Scope puts lands onto the battlefield. "Iron Spider, Stark Upgrade" is a creature. Buster Sword and "Orcrist, Goblin-cleaver" pump, and the free spell of Buster Sword reads the damage. |
| The deck holds two nonland creatures. | Refuted | It holds three: Ingenious Prodigy, "Iron Spider, Stark Upgrade", and Orcish Bowmasters. |
| The deck holds almost no card from before 2020. | Refuted | 36 of 75 rows first printed before 2020, and 33 rows first printed in 2023 or later. |
| The deck misses staples, finishers, and trigger copiers. | Not a choice of the builder | The owner owns 0 of 12 staples, 0 of 3 finishers, and 1 of 5 trigger copiers. The one copier is Sword of Fire and Ice, and the deck holds it. |

## The owned pool

- The collection holds 2,235 distinct names and 6,030 copies. 784 owned cards are legal in Commander and fit a blue-black identity.
- In that pool, 14 cards carry the tag mill-opponent, and the deck holds none of them (F-141).
- 19 cards carry the tag drain-life, and 1 carries alternate-win-condition: Laboratory Maniac. A match of the rules text finds 11 evasive creatures of power 5 or more. The deck holds one drain card and none of the others.
- The owner owns Sunken Hollow, Marsh Flats, Fabled Passage, Reliquary Tower, and "Takenuma, Abandoned Mire". The deck holds none of them.

## The code on `main` at `54b07dd`

### Win conditions

- `roles.go` gives the role wincon to a card with the tag alternate-win-condition alone. The role threat reads an on-theme creature of mana value 4 or more, or an on-theme planeswalker.
- `TargetsFor` holds a flat threat target of 12 and no wincon target. No band, finding, or feature of the quality model reads a win condition.
- The model writes the role of each deck card, and the profile counts that role.
- The goldfish simulation models no damage. The plan judge of the deck gate reads the ways to win, and its grade blocks nothing.
- A profile finding alone buys no repair turn (D-613).

### Mana

- `capLands` ranks lands by the deck colors in `produced_mana`, to a maximum of 2, and then by popularity.
- A fetch land reads no color, so it ranks under every dual land. A tapped land or a land with a condition can tie with a shock land.
- The shortlist score holds no term for an owned card. No step of the mana pass swaps a basic land for a dual land.
- Brackets 4 and 5 pin a Game Changer land and fast mana. They pin no other land, and a two-mana rock never pins.
- A land counts as 1 source per color, and a rock counts as 0.75. The requirement covers 80 percent of the spells of a color, so no check reads the castability of one spell.

### Tags in the full snapshot

| Tag | Cards |
|---|---|
| burn-player | 1,930 |
| mill-opponent | 481 |
| drain-life | 437 |
| gives-double-strike | 154 |
| burn-player-each | 132 |
| alternate-win-condition | 84 |
| overrun | 81 |
| damage-multiplier | 54 |
| extra-combat-phase | 54 |
| poison-opponents | 21 |

## The plan

- M-17 measures first, for no cost (D-721, D-722). It counts the finisher tags and the land classes of real lists by bracket, and it sets the floors. It replays the request of this session with the collection of the owner. It also counts the deck cards whose conditions the deck does not meet.
- PR-52 ranks lands by quality, and the mana pass swaps a basic land for a better land of the pool (D-720).
- PR-53 adds finisher roles from tags, a target per bracket, and finisher pins (D-719). It also stores a count with a floor finding, and it adds a line to the gap note.

## After M-17

M-17 measured on 2026-09-14, and #177 merged it. `docs/reference/m17-finishers-lands-2026-09-14.md` holds every count.

- The theme word "milling" matched no card, so the shortlist held staple roles alone (F-142). The shortlist dropped the mill cards, and not the model.
- PR-54 comes first: aliases, a word-form rule, and a question when no card matches (D-723 to D-725).
- PR-53 sets a finisher target of 3 and a floor of 2 at brackets 1 to 4 (D-726). Bracket 5 takes a target of 1 and a floor of 1.
- Kindred Discovery and Idol of Oblivion work in this deck, and Skullclamp and Springleaf Drum fail. F-140 stays a record until a replay after PR-54 (D-727).

## The suggestions of the review

The review names seven suggestions. The table shows where each one sits now. `docs/reference/owner-review-grima-2026-09-14.md` holds the review word for word.

| Suggestion of the review | Where it sits |
|---|---|
| 1. Classify the payoff shape of the commander before the selection. | OQ-83 |
| 2. Discount a card that grants evasion that the commander already has. | OQ-84 |
| 3. A precondition solver that cuts a card whose condition fails, lands included. | F-140. M-17 counts the cards first (D-722). |
| 4. A mana check per spell against Karsten, and dual lands before generic filler. | PR-52 ranks lands (D-720). The owner chose no castability check now (D-720). |
| 5. Require a win condition, and estimate a clock. | PR-53 adds the target (D-719). The owner chose no clock estimate now (D-719). |
| 6. Caps per effect class, weighted by what the commander cares about. | OQ-85 |
| 7. Report the rules bracket beside a power estimate. | OQ-86 |

The review also names cards. The table shows each group and the cards of it that the owner owns. The review writes two lands by a short name, and the full name follows in parentheses.

| Group | Cards the review names | Owned |
|---|---|---|
| Trigger copiers and double strike | Strionic Resonator, Fireshrieker, Grappling Hook, Sword of Feast and Famine, Sword of Fire and Ice | Sword of Fire and Ice, in the deck |
| Finishers | Torment of Hailfire, Exsanguinate, Blue Sun's Zenith | None |
| Dual lands and utility lands | Watery Grave, Darkslick Shores, Drowned Catacomb, Choked Estuary, Undercity Sewers, Morphic Pool, River of Tears, Underground River, Bojuka Bog, Otawara ("Otawara, Soaring City"), Takenuma ("Takenuma, Abandoned Mire") | "Takenuma, Abandoned Mire", not in the deck |
| Mana rocks | Dimir Signet, Talisman of Dominance | None |
| Staples | Counterspell, Swan Song, An Offer You Can't Refuse, Rhystic Study, Mystic Remora, Toxic Deluge, Reliquary Tower | Reliquary Tower, not in the deck |
