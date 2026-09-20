# PR-57: the typal signal on a card that is no land (2026-09-20)

This document holds every count PR-57 rests on. It closes the typal-payoff half of F-144. The owner decisions are D-765 to D-768.

Source of every card fact: the local card snapshot `20260904T210157`, of 2026-09-04. The theme table `themes.json` reads `verified_at` 2026-09-14.

One paid target ran: `make deck-gate`, run 29, for $2.7384.

## The defect

F-144 records that a type row of `themes.json` reads no card that makes or names its creature type. PR-55 closed the land half (D-759). The nonland half stayed open.

A type row reads three things: the subtype on the card's own type line, a typal tag of that type, and a few payoff phrases. A generic typal payoff carries none of them. Door of Destinies names no creature type, its type line names no creature type, and no typal tag of a type holds it. So the dinosaur shortlist read it as a card of no theme.

## The signal

Two tags of the snapshot hold the generic typal payoffs.

1. The tag `typal-choose`, "Choose a creature type with which to synergize". It holds 94 cards: 6 lands and 88 nonlands.
2. The tag `typal-share`, which holds a card that rewards creatures of a shared type, such as Coat of Arms. It holds 77 cards: 2 lands and 75 nonlands.

The card index of the app drops 3 of the 88. Their layout is not playable: `Littjara` is a plane, `Oko, Shadowmoor Scion Emblem` is an emblem, and `Mirror Entity Avatar` is a vanguard. So a build can read 85, and 81 of the 85 are legal in Commander.

CAUTION: D-760 names Coat of Arms as a card of `typal-choose`. That is wrong. Coat of Arms carries `typal-share`. The count of 88 in D-760 is correct.

The table `themes.json` holds both tags in the `typal` block, which PR-55 named `typal_land`. The block now holds three lists:

| List | Counts on | Holds |
|---|---|---|
| `land_slugs` | a land alone | `typal-choose` |
| `land_text` | a land alone | two needles, and the subtype word of each row |
| `card_slugs` | a card that is no land | `typal-choose` |
| `noncreature_slugs` | a card that is no land and no creature | `typal-share` |

Every row with a subtype carries all four lists, and the generic rule carries them for a word that names a creature type (D-761).

## Why `typal-share` reads no creature

The tag `typal-share` holds 70 Commander-legal nonlands. They split in two:

| Shape | Count | Examples |
|---|---|---|
| No creature and no Kindred card | 30 | Coat of Arms, Shared Animosity, Descendants' Path, Pyre of Heroes |
| A creature or a Kindred card | 40 | Kilnmouth Dragon, Hunting Velociraptor, Knowledge Exploitation |

A card of the second shape names one creature type on its own type line, and it rewards that type alone. Kilnmouth Dragon reveals Dragon cards, so a rabbits list must read no signal on it. The rule for `noncreature_slugs` drops all 40 (D-765).

A Kindred card counts as a creature for this rule. Knowledge Exploitation is a "Kindred Sorcery — Rogue", and its prowl cost needs a Rogue.

The tag `typal-choose` needs no such rule, because each card of it chooses a creature type. 22 of its 81 Commander-legal nonlands are creatures that choose one. Adaptive Automaton and Metallic Mimic are two of them.

## The weight

A typal card weighs as payoff text, 1.2 against the score cap of 2.5 (D-766). It is the weight of a typal land (D-759). The tags state one fact about one card, so a card that holds both weighs 1.2 and no more.

A real card of the type still outranks a generic payoff. A payoff tag and the subtype together score 2.3 of the cap, against 1.2.

## The false positive

`Kolvori, God of Kinship` carries `typal-choose`, and its text chooses no creature type. It reads legendary creatures. That is 1 card of the 85. Every other card of the tag rewards a creature type the player chooses.

## The reach

The prompt is prompt 4 of the deck gate: a dinosaur Commander deck at bracket 2, led by Gishath, Sun's Avatar, with the any-card pool.

| Theme | On theme, before | With `typal-choose` | With both tags | Shortlist, before | Shortlist, after |
|---|---|---|---|---|---|
| dinosaurs | 215 | 268 | 289 | 266 | 295 |
| dinosaur | 227 | 280 | 301 | 275 | 295 |

The shortlist holds 40 lands before the change and 40 lands after it.

## The question of D-725 is not masked

The theme row asks for the theme again when no card of the format and the colors matches (D-725). It reads `Stats.OnTheme == 0`. So this run measured every type row against every single color, before the change and after it.

Every one of the 75 pairs read an on-theme count of 1 or more before the change. The lowest was 7, for rabbits in blue, black, and red. So the question fired for no pair before the change, and the change masks none.

A type row is not the whole reach. The generic rule carries the signal for any word with a `typal-<word>` slug (D-761). So this run also read every such word against every single color: 190 words and 950 pairs.

| Measure | Before | After |
|---|---|---|
| Pairs with an on-theme count of 0 | 5 | 5 |
| The words of those pairs | ox | ox |

The word "ox" reads 0 in every color, before the change and after it. The snapshot holds the slug `typal-ox`, and it holds no card that the slug or the subtype reads in one color alone. So the theme row still asks about "ox", and the change masks no pair of the other 189 words.

## The paid run

Deck gate run 29 ran on 2026-09-20 for $2.7384, in 1874 seconds and 75 calls. It reads PASS: 25 of 25 decks passed every block check, 0 invented names reached the user, and 0 summaries stated a false rule.

CAUTION: the run header names the commit `5fd8085`, the base of this branch. The change was not committed when the run started, so the header names the parent and not the tree.

The dinosaur deck of prompt 4 holds eight generic typal payoffs:

- Adaptive Automaton, Metallic Mimic, and Roaming Throne.
- Herald's Horn, Icon of Ancestry, and Vanquisher's Banner.
- Descendants' Path and Shared Animosity.

The same deck of run 28 held none of them. Its `theme_fit` moved from `partly` to `yes`, and its `summary_honest` moved from `yes` to `partly`.

CAUTION: run 28 ran on 2026-09-16, before #190 and #191. So the whole-run numbers carry PR-55 and PR-56 too, and not this change alone. Both runs read the same card snapshot, the same generate prompt version 15, the same plan rubric, and the same quality model.

Over the 25 prompts, 15 plan fields improved and 13 got worse. The mean plan score moved from 0.69 to 0.70. The judge is noisy (D-230), so those numbers measure nothing on their own. The dinosaur prompt is the one this change targets.

## What stays open

F-144 stays open for its other half: a card whose text names the type, and that the row reads with no signal. A free count over the Commander-legal cards of the snapshot, with the matcher of this pull request:

| Row | Cards whose text names the type | The row reads none of them | Lands among those |
|---|---|---|---|
| spirits | 342 | 182 | 0 |
| zombies | 325 | 162 | 0 |
| humans | 306 | 121 | 0 |
| dragons | 211 | 66 | 0 |
| goblins | 213 | 51 | 0 |
| angels | 85 | 49 | 0 |
| knights | 107 | 49 | 0 |
| vampires | 127 | 29 | 0 |
| cats | 55 | 20 | 0 |
| wizards | 133 | 18 | 0 |
| dinosaurs | 86 | 16 | 0 |
| elves | 128 | 16 | 0 |
| rabbits | 18 | 6 | 0 |
| merfolk | 78 | 1 | 0 |
| slivers | 119 | 0 | 0 |

PR-55 covers every land of this gap, so the last column reads 0 for each row. Most of the rest make a token of the type, such as Grave Titan, Divine Visitation, and Ugin, the Ineffable.

CAUTION: a plain text needle for the type word is not safe. It also reads the cards that punish the type. Mikaeus, the Unhallowed gives "other non-Human creatures" +1/+1, Return of the Wildspeaker reads "non-Human creatures you control", and Walk the Plank destroys a "non-Merfolk creature". A needle with no rule for "non-" ranks all three as on theme for the row they work against. So this half needs its own rule, its own measurement, and its own pull request (D-768).
