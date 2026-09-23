# F-166: the commander rate of a bracket 5 shortlist, 2026-09-23

This record holds the measurements of F-166 and PR-67 (D-839). The base is `f00a2af`. The card snapshot is `20260904T210157`, and the local quality model before the change is `20260914T154223Z`.

## 1. The two judge reads

Bracket gate run 9 built Najeela, the Blade-Blossom at bracket 5 for the theme "warriors combat". Both judge reads gave bracket 4, and bracket gate run 10 gave bracket 4 again. Each read names the warrior filler, for example Jazal Goldmane and Pact of the Serpent. The judge reads the deck alone, and not the theme of the request. Its prompt says that a bracket 5 deck "plays the best strategy and not a theme" (`go/internal/generate/judge.go`).

The judge reads real cEDH lists well. The calibration judge of 2026-09-13 read 11 of 12 TopDeck top-cut lists as bracket 5, and the Magda list among them.

## 2. The Najeela lists of the local meta store

The local meta store holds 64 TopDeck lists of Najeela, 8 of them top cut, from 2026-06-04 to 2026-09-14. A free count compared them with the decks of runs 9 and 10.

| | Run 9 deck | TopDeck lists of Najeela |
|---|---|---|
| Warrior creatures | 21 | median 4, from 0 to 14 |
| Cards in under 10 percent of the lists | 63 of 95 names | - |
| Derevi, Empyrial Tactician | absent | 61 of 64 |
| Thassa's Oracle | absent (run 10 holds it) | 54 of 64 |
| Demonic Consultation, Tainted Pact | absent | 54 and 55 of 64 |

## 3. The shortlist replay

A free replay built the shortlist of bracket gate prompt 15 on `f00a2af`. It held 332 cards and 144 warrior creatures. It held 30 of the 91 cards that half the Najeela lists play.

The rate of the format gives Derevi 0.08, so the shortlist never held it. The rate of the format counts nonland cards alone, so each fetch land and each original dual reads 0. Thassa's Oracle was on the shortlist of run 9, and the model left it out of the deck.

## 4. The rejudge of two scratch decks

Two lanes of `make bracket-gate` with `-rejudge` read two scratch decks, for $0.0325 and $0.0276, together $0.0601.

| Deck | Lane a | Lane b |
|---|---|---|
| 1. The newest top-cut Najeela list, 2026-09-12 | 5 | 5 |
| 2. The run 9 deck, with its 58 cards under 10 percent of the lists replaced by the most played cards that it lacks | 5 | 5 |

Deck 2 is the upper bound of the lever: it holds 5 warrior creatures, as the top-cut list does. So the judge reads the line of the bracket, and the builder misses that line.

## 5. The sweep of the theme boost

The commander rate reads the TopDeck lists before the tier cap of the fit. The size check of the fit leaves 219 commanders with 10 lists or more, over 20,775 lists. The rows hold 37,275 cards at a rate of 0.1 or more.

The free sweep built the three bracket 5 prompts of the bracket gate. The core of a commander is the set of legal cards that half its lists play.

| Prompt | Lists | Core held, no rate | Core held, each boost from 0.02 to 0.5 |
|---|---|---|---|
| 13. Kinnan, Bonder Prodigy | 1,677 | 42 of 88 | 88 of 88 |
| 14. Yuriko, the Tiger's Shadow | 144 | 45 of 85 | 85 of 85 |
| 15. Najeela, the Blade-Blossom | 64 | 30 of 91 | 87 of 91 |

The boost moved the warrior creatures of the Najeela shortlist, against 144 with no rate:

| Boost | 0.02 | 0.05 | 0.1 | 0.2 | 0.3 | 0.5 |
|---|---|---|---|---|---|---|
| Warrior creatures | 40 | 47 | 67 | 88 | 100 | 105 |

The rule of the session was the largest boost that keeps each core card. Each boost kept the same core, so that rule did not separate the boosts. The owner chose 0.05 (D-839).

Najeela misses the same four core lands at each boost: Underground Sea, Savannah, Exotic Orchard, and Boseiju, Who Endures. The land cap drops them, and this lever does not change the land cap.
