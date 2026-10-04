# Fast mana with a hidden cost, 2026-10-04

This note holds every count that PR-130 rests on (F-225, D-1162). A scratch Go test made each count, and git holds no copy of it. The test read the card snapshot of 2026-09-04, the same snapshot as `docs/reference/power-card-counters-2026-10-04.md`.

## The deck of the check

The check of PR-128 ran `make api-build` on the deploy of `a7731f5` on 2026-10-04. It sent "Build me a bracket 4 red and green Commander deck from my collection with no budget." with `-pool owned-only`.

| Field | Value |
|---|---|
| Session | `KeGnDO4SLj0cxKBuDqTg` |
| Deck | `HpWYMR9DrNN2YbHNBoWt` |
| Commander | Toph, Hardheaded Teacher |
| Bracket | 4 |
| Cost | $0.0570 |
| Fast mana marks | Astral Cornucopia, Sol Ring |
| Finisher marks | Great Train Heist, Mongoose Lizard, Smaug, the Great Calamity |

The deck held no filter, no Springleaf Drum, no Earthquake, and no Hurricane. The test collection holds no Hurricane. So the check proves that the counts are clean, and it does not test those cards.

## The fault

Scryfall gives Astral Cornucopia the cost {X}{X}{X} and a mana value of 0. Its one mana ability taps the card alone and pays no mana, so it meets D-1159. One mana costs 3, so the card is ramp and not fast mana.

## The cards of the old rule

Under D-1159 the snapshot holds 77 fast mana cards. Three groups hide a real cost:

| Group | Cards | The real cost |
|---|---|---|
| An X in the cost | Astral Cornucopia, Mana Bloom | One mana costs 3, or 2 for Mana Bloom |
| A multikicker | Everflowing Chalice | Cast for 0, it adds no mana |
| Other slow cards | Paradise Mantle, Pyramid of the Pantheon, Thran Turbine, Lotus Bloom, Mox Tantalite, Sol Talisman | Each one needs its own rule |

The owner chose the first two groups (D-1162). The test collection holds Astral Cornucopia and Mana Bloom.

## The replay

The new rule reads 74 fast mana cards. It drops Astral Cornucopia, Mana Bloom, and Everflowing Chalice, and no other card. On the new rule the deck of the check reads 1 fast mana, Sol Ring.
