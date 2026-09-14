# PR-49: the shortlist with the top-list rate of the app

Run date: 2026-09-14. Card snapshot: `20260904T210157`. Quality model: `20260910T012734Z`, the newest local model, and the model that graded deck gate run 19.

## What this reads

The app passes the top-list rate of each card from the quality model into the shortlist (`MetaBoost`, a weight of 0.1 in `candidates.go`). Before PR-49 the bracket gate and the deck gate built their shortlists with no rate (F-129). The dry run of each gate now builds each shortlist twice, with the rate and with no rate, and prints what the rate brings in (D-706). No model call ran, and the run cost nothing.

- **Cards** is the shortlist with the rate, the upgrades included.
- **In** counts the cards that the list with the rate holds and the list with no rate does not.
- **Tutors**, **fast mana**, and **Game Changers** count by the rules of the profile (`profile.PowerCards`), each as the count with no rate to the count with the rate. A tutor is under the Tagger slug `tutor` and not `tutor-land`. Fast mana is a nonland, noncreature mana source of mana value 1 or less. A Game Changer reads its flag.

The commands: `CARDS_SNAPSHOT_DIR=<local snapshot> go run ./cmd/bracket-gate -dry`, and `go run ./cmd/deck-gate -dry -collection internal/collections/testdata/manabox_collection.csv`, both from `go/`.

## Bracket gate, 15 prompts

| # | Prompt | Pool | Cards | In | Tutors | Fast mana | Game Changers |
|---|---|---|---|---|---|---|---|
| 1 | bracket 1 Gishath, Sun's Avatar | 267 | 263 | 19 | 4 | 1 to 6 | 0 |
| 2 | bracket 1 Karlov of the Ghost Council | 303 | 300 | 1 | 4 | 1 | 0 |
| 3 | bracket 1 Adeline, Resplendent Cathar | 296 | 294 | 0 | 3 | 1 | 0 |
| 4 | bracket 2 Gishath, Sun's Avatar | 267 | 263 | 19 | 4 | 1 to 6 | 0 |
| 5 | bracket 2 Karlov of the Ghost Council | 303 | 300 | 1 | 4 | 1 | 0 |
| 6 | bracket 2 Adeline, Resplendent Cathar | 296 | 294 | 0 | 3 | 1 | 0 |
| 7 | bracket 3 Karlov of the Ghost Council | 303 | 300 | 1 | 4 | 1 | 0 |
| 8 | bracket 3 Denethor, Ruling Steward | 295 | 292 | 0 | 6 | 2 | 0 |
| 9 | bracket 3 Zada, Hedron Grinder | 294 | 292 | 10 | 4 | 3 to 9 | 2 to 5 |
| 10 | bracket 4 Urza, Lord High Artificer | 299 | 297 | 3 | 13 | 2 to 5 | 1 to 4 |
| 11 | bracket 4 Korvold, Fae-Cursed King | 300 | 296 | 0 | 11 | 0 | 0 |
| 12 | bracket 4 Prosper, Tome-Bound | 188 | 185 | 5 | 1 | 0 | 2 to 3 |
| 13 | bracket 5 Kinnan, Bonder Prodigy | 303 | 300 | 7 | 24 to 26 | 1 | 8 |
| 14 | bracket 5 Yuriko, the Tiger's Shadow | 213 | 210 | 21 | 1 | 10 to 12 | 9 to 13 |
| 15 | bracket 5 Najeela, the Blade-Blossom | 306 | 300 | 1 | 4 | 0 | 0 |

## Deck gate, 25 prompts

| # | Prompt | Pool | Cards | In | Tutors | Fast mana | Game Changers |
|---|---|---|---|---|---|---|---|
| 1 | lifegain Commander, any card | 303 | 300 | 1 | 4 | 1 | 0 |
| 2 | aristocrats Commander, owned first | 179 | 226 | 0 | 1 | 3 | 1 |
| 3 | artifacts Commander, bracket 4 | 299 | 297 | 3 | 13 | 2 to 5 | 1 to 4 |
| 4 | dinosaur tribal, bracket 2 | 267 | 263 | 19 | 4 | 1 to 6 | 0 |
| 5 | blink Commander, owned first | 168 | 216 | 0 | 0 | 4 | 0 |
| 6 | Modern tempo, tournament | 167 | 165 | 20 | 1 | 5 | 2 |
| 7 | Modern burn, casual | 290 | 289 | 1 | 4 | 0 | 0 |
| 8 | Modern lifegain, FNM | 301 | 299 | 2 | 3 | 1 | 0 |
| 9 | Standard midrange, FNM | 164 | 162 | 6 | 2 | 1 | 0 |
| 10 | Standard aggro, tournament | 294 | 292 | 5 | 6 | 1 | 0 |
| 11 | Commander with a locked card | 296 | 292 | 0 | 7 | 2 | 0 |
| 12 | Commander on a budget | 296 | 294 | 0 | 3 | 1 | 0 |
| 13 | owned first, and the commander is not owned | 229 | 276 | 1 | 1 | 3 | 1 |
| 14 | the user delegates the commander | 244 | 241 | 4 | 9 | 1 | 1 |
| 15 | delegated commander, owned first | 229 | 276 | 2 | 1 | 3 | 0 |
| 16 | a tight budget, owned first | 229 | 226 | 0 | 1 | 3 | 0 |
| 17 | upgrade a precon, owned first | 196 | 228 | 1 | 4 | 5 | 0 |
| 18 | upgrade a precon, any card | 302 | 232 | 45 | 0 | 3 to 12 | 10 to 15 |
| 19 | the Hobbit family, two colours (sets hob,hoc in 170 out 0) | 174 | 170 | 0 | 0 | 2 | 2 |
| 20 | the Hobbit family, a delegated commander (sets hob,hoc in 75 out 0) | 77 | 75 | 0 | 0 | 2 | 1 |
| 21 | the Hobbit family, mana from outside (sets hob,hoc in 127 out 26) | 156 | 153 | 0 | 0 | 2 | 3 |
| 22 | a set family and a card from outside it (sets hob,hoc in 170 out 0) | 175 | 170 | 0 | 0 | 2 | 2 |
| 23 | two set families at once (sets blb,blc,hob,hoc,pblb in 201 out 0) | 203 | 201 | 0 | 1 | 3 | 1 |
| 24 | a 60-card deck from one set (sets blb,blc,pblb in 95 out 0) | 96 | 95 | 0 | 0 | 0 | 0 |
| 25 | use no card of an owned precon (excludes 57 cards of Avengers Assemble, 27 spare) | 152 | 200 | 15 | 0 | 9 to 14 | 5 to 9 |

## Read

- **The rate moves few cards on most shortlists.** It brings in 0 to 21 cards of a bracket gate shortlist and 0 to 45 of a deck gate shortlist. 15 of the 40 shortlists do not move.
- **Where it moves cards, it brings in power.** Fast mana rises on bracket gate prompts 1, 4, 9, 10, 14, and on deck gate prompts 3, 4, 18, 25. Game Changers rise on bracket gate prompts 9, 10, 12, 14, and on deck gate prompts 3, 18, 25. Tutors rise on bracket gate prompt 13 alone.
- **The rate reads no bracket.** It brings fast mana into the Gishath shortlist at brackets 1 and 2, from 1 to 6. The prompt and the profile still hold the band of the bracket.
- **Every earlier gate run built with no rate.** Deck gate run 19 and bracket gate runs 2 and 3 measured shortlists that the app does not build. The next run of each gate reads parity and any other change together.
