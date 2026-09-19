# PR-55: the typal land signal of a type row (2026-09-19)

This document holds every count PR-55 rests on. It answers F-148, and it closes the land half of F-144. The owner decisions are D-759, D-760, and D-761.

Source of every card fact: the local card snapshot `20260904T210157`, of 2026-09-04. The theme table `themes.json` reads `verified_at` 2026-09-14.

Every count of this document comes from a free run. No paid target ran for PR-55.

## The defect

F-148 records that a typal theme loses the lands that make mana for its type. The type rows of `themes.json` read the subtype, a typal tag, and a few payoff phrases. None of them reads a land.

The reproduction is prompt 4 of the deck gate: a dinosaur Commander deck at bracket 2, led by Gishath, Sun's Avatar, with the any-card pool. The shortlist holds 266 cards and 40 lands.

| Land | In the shortlist | Theme score | Land rank |
|---|---|---|---|
| Cavern of Souls | no | 0.000 | 19 |
| Secluded Courtyard | no | 0.000 | 19 |
| Unclaimed Territory | no | 0.000 | 19 |
| Path of Ancestry | no | 0.000 | 25 |
| Three Tree City | no | 0.000 | 31 |

The land cap of `capLands` gives half its places to the mana order and half to the theme lands. No land of this build read a theme signal, so the mana order filled all 40 places. The last land it kept ranks 13, and every land of the table above ranks under it.

## The signal

Two facts make a land typal, and both count on a land alone (D-760).

1. The Tagger tag `typal-choose`, "Choose a creature type with which to synergize". The snapshot holds it on 94 cards, and 6 of them are lands.
2. The text needle "shares a creature type with your commander". Path of Ancestry names no chosen type, and this needle is the one that reads it. The snapshot holds the needle on 1 land and on no other card.

The table `themes.json` holds both in a new `typal_land` block. Every row with a subtype carries them.

Each row also adds its own subtype word as a land needle (D-759). That reaches the lands that name the type, such as Sliver Hive and Restless Ridgeline. The count per type, over the 1,194 Commander-legal lands of the snapshot:

| Type | Lands whose text names it | Of these, lands that make mana |
|---|---|---|
| Dragon | 5 | 5 |
| Elf | 6 | 5 |
| Goblin | 5 | 5 |
| Zombie | 7 | 7 |
| Vampire | 2 | 2 |
| Angel | 3 | 3 |
| Human | 3 | 3 |
| Merfolk | 2 | 2 |
| Sliver | 1 | 1 |
| Cat | 1 | 1 |
| Dinosaur | 1 | 1 |
| Wizard | 3 | 3 |
| Knight | 1 | 1 |
| Rabbit | 2 | 2 |
| Spirit | 8 | 8 |

Forbidden Orchard is the one incidental hit. It gives an opponent a Spirit token, and a spirits theme reads it as on theme. Its mana rank is 1, so it holds a place of the mana half already.

A typal land counts once, at the payoff-text weight of 1.2 against the score cap of 2.5. The tag and the needles state one fact about one card.

## The word boundary

A land needle matches on a word boundary, with an optional plural "s". A short creature type is a part of a common word of land text. The Gitar review of #190 found it, and a free run measured it over the 238 typal words and the 1,194 Commander-legal lands:

| Needle | Lands it reads as a substring, and not as a word | The word that holds it |
|---|---|---|
| bat | 115 | battlefield |
| orc | 63 | sorcery |
| mount | 40 | mountain |
| rat | 6 | (several) |
| the other 234 words | 8 | (several) |

The word-boundary rule drops 232 such hits of 850, and it drops no true hit. The optional "s" keeps a needle that reads a plural, such as "zombie" against "Zombies you control".

The rule drops one land that belongs on a typal list. Abundant Countryside makes a Shapeshifter token with changeling, and the word "changeling" holds the letters of "angel". A changeling permanent is every creature type, so the land rewards every typal deck. The `typal_land` block holds the needle "every creature type", and the land now reads every typal theme on purpose. The snapshot holds 1 such land.

CAUTION: the Gitar finding named two examples that no card holds. It named the land Cascading Cataracts, whose oracle text holds no form of "cat". It also named "combat", which no land text holds. The three words of the table above are the real ones.

## The singular type word

A singular creature type keeps the generic rule and reads no type row (D-731). The file `go/cmd/questions-gate/conversations.json` records the theme "dinosaur" in two conversations, and not "dinosaurs". So the singular is the shape a reader writes.

A free run measured the split on the dinosaur prompt of the deck gate:

| Theme | Land signals | Lands of F-148 on the shortlist | On theme |
|---|---|---|---|
| dinosaurs | 1 tag, 3 needles | 5 of 5 | 215 |
| dinosaur, before D-761 | none | 0 of 5 | 221 |
| dinosaur, after D-761 | 1 tag, 3 needles | 5 of 5 | 227 |

The generic rule carries the signal now (D-761). The trigger is a `typal-<word>` slug that the card database holds. The snapshot of 2026-09-04 holds 238 such slugs, against the 15 type rows of `themes.json`. So a soldiers deck and an elemental deck get their lands too.

CAUTION: the card database decides the reach of the generic rule. A new Scryfall tag changes it with no change of this repo. `make themes-check` reads the rows of `themes.json` and not the generic rule.

## The result

The same build, with the signal:

| Land | In the shortlist | Theme score | Land rank |
|---|---|---|---|
| Cavern of Souls | yes | 0.480 | 19 |
| Secluded Courtyard | yes | 0.480 | 19 |
| Unclaimed Territory | yes | 0.480 | 19 |
| Path of Ancestry | yes | 0.480 | 25 |
| Three Tree City | yes | 0.480 | 31 |
| Restless Ridgeline | yes | 0.480 | 25 |
| Abundant Countryside | yes | 0.480 | 19 |

The shortlist still holds 266 cards and 40 lands. The on-theme count moves from 208 to 215. Seven lands of rank 13 leave to make the room. Each one enters tapped unless a condition holds:

- Furycalm Snarl, Fortified Village, and Game Trail: the player reveals a land card of the right type from hand.
- Overgrown Farmland and Sundown Pass: the player controls two or more other lands.
- Radiant Summit: the player controls two or more basic lands.
- Starting Town: it is the first, the second, or the third turn of the player.

## The question of D-725 is not masked

The theme row asks for the theme again when no card of the format and the colors matches (D-725). It reads `Stats.OnTheme == 0`. A land signal that fires in every deck hides a real miss. So this run measured every type row against every single color, before the change and after it.

Every one of the 75 pairs read an on-theme count of 1 or more before the change. The lowest was 1, for rabbits in blue, black, and red. So the question fired for no pair before the change, and the change masks none. The change adds 6 or 7 cards to each pair. That is the 5 generic typal lands, the changeling land, and the land that names the type where one exists.

A two-color identity holds every card of each of its colors, so a single color is the worst case.

## What no run measured

No build of a model ran for PR-55. The deck gate costs about $2.75, and the owner approved no paid run. So this document states what the shortlist holds, and it states nothing about the deck a model writes from it.
