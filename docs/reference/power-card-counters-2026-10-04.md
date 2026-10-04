# The power card counters, 2026-10-04

This note holds every count that PR-128 rests on (F-224, D-1159, D-1160). A scratch Go test made each count, and git holds no copy of it. The test read the card snapshot of 2026-09-04 and the local meta store.

## The verdict

On 2026-10-04 at 19:26 UTC a reader wrote verdict `x5JGF0zyIE3DcID7QB5c` on the decks screen. The verdict holds no deck id. The reader made a deck at 19:21 UTC, so the verdict names that deck by its time.

The reader wrote: "Too many bad cards like barbed sextant and others like spring leaf drum that are not needed in mono-green. Hurricane is not really a wincon"

| Deck | Commander | Bracket | Pool | Fast mana marks | Finisher marks |
|---|---|---|---|---|---|
| 19:21 UTC | Radagast of Rhosgobel | 4 | owned only | Barbed Sextant, Sol Ring, Springleaf Drum | Hurricane, Time Bomb |
| 19:31 UTC | Hearthhull, the Worldseed | 4 | owned only | Barbed Sextant, Sol Ring, Springleaf Drum, Wild Growth | Dry Spell, Hurricane |

Bracket 4 holds a fast mana floor of 3 and a finisher floor of 2 (D-704, D-726). The deck shape block asks the model to build inside the floors. So the model took the cards that the counters marked.

## Fast mana

The old rule counted each nonland, noncreature card of mana value 1 or less that adds mana. The new rule also needs one mana ability that pays no mana and taps no creature (D-1159).

| Count | Cards of the snapshot, of each format |
|---|---|
| Old rule | 97 |
| New rule | 77 |
| Dropped | 20 |

The 20 dropped cards are Springleaf Drum, Chromatic Star, Chromatic Sphere, Arcum's Astrolabe, Wizard's Rockets, "Barrels of Blasting Jelly", Terrarion, Jack-o'-Lantern, Urn of Godfire, Barbed Sextant, Skycloud Egg, Giant's Boulder, Mana Cylix, Darkwater Egg, Mossfire Egg, Shadowblood Egg, Sungrass Egg, Color Pie, Orb of Origin, and Mana Screw.

The new rule keeps Sol Ring, Mana Vault, Chrome Mox, Mox Diamond, Mox Opal, Mox Amber, Lotus Petal, Lion's Eye Diamond, Dark Ritual, and Rite of Flame. The new rule adds no card.

## The tag burn-player-each

Tagger puts the tag under `burn-player` and `burn-you`. The tag holds 132 cards, and the text of 131 of them deals damage to each player.

| Group | Decks | With a card of the tag | Of them, 8 or more lifegain cards | Under the finisher floor of 2, before | After the removal |
|---|---|---|---|---|---|
| Precons since 2023 | 96 | 4 | 1 | 11 | 11 |
| EDHREC average decks | 1,039 | 54 | 6 | 275 | 278 |
| TopDeck good, since 2026-08-01 | 8,535 | 334 | 7 | 5,912 | 5,981 |
| TopDeck top cut, since 2026-08-01 | 3,104 | 128 | 4 | 2,220 | 2,247 |

A lifegain card is a card of the Tagger tree `lifegain`. The EDHREC average decks with a card of the tag hold a median of 2 lifegain cards. The EDHREC average decks with no such card hold a median of 4.

The lists play these cards most: Cave-In, Molten Disaster, Pestilence, Crypt Rats, Earthquake, Price of Progress, and Descent into Avernus. Each one acts as removal or as damage to all players, not as the one route to a win.

## The replay

The two stored decks of the reader read these counts on the new code:

| Deck | Fast mana, before | Fast mana, after | Finishers, before | Finishers, after |
|---|---|---|---|---|
| 19:21 UTC | 3 | 1 (Sol Ring) | 2 | 0 |
| 19:31 UTC | 4 | 2 (Sol Ring, Wild Growth) | 2 | 0 |

A deck under the finisher floor gets the gap note, and no repair turn (D-743). The replay read the stored lists alone. It made no new deck, and it called no model.
