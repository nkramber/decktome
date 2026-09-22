# The fixing floors of F-33, 2026-09-22

This document holds every number that the fixing band of F-33 rests on (D-798, D-799). The session read each one on 2026-09-22, with no provider call.

## 1. The gap in the gate decks

A free count read the lands of every deck in eight whole deck gate runs: 19, 23, 24, 25, 28, 29, 30, and 31. The column is the count of nonbasic lands. The count in parentheses is the count of basic land types, which is near the count of deck colors.

| Deck | Pool | 19 | 23 | 24 | 25 | 28 | 29 | 30 | 31 |
|---|---|---|---|---|---|---|---|---|---|
| 2. aristocrats Commander | owned first | 4 (2) | 9 (2) | 0 (2) | 2 (2) | 10 (2) | 7 (2) | 12 (2) | 12 (2) |
| 9. Standard midrange, FNM | any card | 12 (2) | 16 (2) | 4 (2) | 4 (2) | 4 (2) | 8 (2) | 12 (2) | 12 (2) |
| 10. Standard aggro, tournament | any card | 4 (2) | 6 (2) | 4 (2) | 8 (2) | 0 (2) | 12 (2) | 8 (2) | 16 (2) |
| 13. commander not owned | owned first | 6 (2) | 5 (2) | 0 (2) | 8 (2) | 5 (2) | 8 (2) | 9 (2) | 0 (2) |
| 15. delegated commander | owned first | 8 (2) | 0 (2) | 0 (2) | 4 (2) | 8 (2) | 11 (2) | 4 (2) | 8 (2) |
| 19. the Hobbit family | any card | 6 (3) | 0 (3) | 7 (3) | 0 (3) | 8 (3) | 0 (3) | 0 (3) | 10 (3) |
| 22. a set family | any card | 5 (3) | 7 (3) | 0 (3) | 0 (3) | 0 (3) | 0 (3) | 0 (3) | 9 (3) |

The pool does not cause the gap. Deck 13 of run 31 is a white-black Karlov deck with 36 basic lands. The test collection of the gate holds white-black fixing lands, for example Command Tower, Exotic Orchard, Path of Ancestry, Opal Palace, and Secluded Courtyard.

Three facts of the code let the gap through:

- The color sources band counts a basic land as a full source, so a deck of basic lands passes it.
- The swap of `swapBasics` runs at Commander brackets 4 and 5 alone (D-734).
- No line of the prompt names fixing lands.

## 2. The measurement

A scratch test read each stored list of the local meta store of 2026-09-14 against the card snapshot of 2026-09-04. The test sat in `go/cmd/deck-gate` for the run, and the session then deleted it. It counts a land as fixing when `profile.LandClassOf` reads a class below `LandOther` for the deck colors. The deck colors are the color identity of the commander in Commander, and the colors of the nonland cards in a 60-card deck.

The test did not resolve 3,535 lists, because each one named a card that the snapshot does not hold. It read each other list.

The table gives the low quarter (p25) and the median of the fixing lands. The ladder of `go/internal/meta/meta.go` names each tier:

- A great list is a top-8 finish or a competitive cEDH list.
- A good list is a league finish or the rest of a challenge.
- A typical list is the EDHREC average deck.
- A baseline list is a precon.

| Format | Tier | Colors | Lists | p25 | Median |
|---|---|---|---|---|---|
| Commander | baseline | 2 | 61 | 7 | 8 |
| Commander | baseline | 3 | 99 | 13 | 16 |
| Commander | baseline | 4 | 8 | 15 | 16 |
| Commander | baseline | 5 | 10 | 21 | 23 |
| Commander | typical | 2 | 383 | 11 | 12 |
| Commander | typical | 3 | 133 | 17 | 19 |
| Commander | typical | 4 | 6 | 20 | 21 |
| Commander | typical | 5 | 26 | 21 | 22 |
| Commander | good | 2 | 5,004 | 9 | 11 |
| Commander | good | 3 | 6,288 | 12 | 14 |
| Commander | good | 4 | 2,376 | 18 | 19 |
| Commander | good | 5 | 1,734 | 21 | 23 |
| Commander | great | 2 | 1,668 | 9 | 11 |
| Commander | great | 3 | 2,197 | 12 | 14 |
| Commander | great | 4 | 1,072 | 18 | 19 |
| Commander | great | 5 | 678 | 21 | 22 |
| Modern | good | 2 | 9,017 | 7 | 9 |
| Modern | good | 3 | 4,696 | 10 | 15 |
| Modern | good | 4 | 1,718 | 12 | 16 |
| Modern | good | 5 | 959 | 12 | 18 |
| Modern | great | 2 | 1,403 | 7 | 9 |
| Modern | great | 3 | 766 | 11 | 15 |
| Modern | great | 4 | 270 | 13 | 17 |
| Modern | great | 5 | 151 | 12 | 18 |
| Standard | good | 2 | 4,013 | 12 | 12 |
| Standard | good | 3 | 1,023 | 15 | 19 |
| Standard | good | 4 | 559 | 22 | 22 |
| Standard | good | 5 | 14 | 21 | 21 |
| Standard | great | 2 | 1,042 | 12 | 12 |
| Standard | great | 3 | 254 | 15 | 19 |
| Standard | great | 4 | 123 | 22 | 22 |
| Standard | great | 5 | 5 | 20 | 21 |
| 60-card precon | baseline | 2 | 335 | 0 | 0 |
| 60-card precon | baseline | 3 | 61 | 0 | 2 |

## 3. The floors

Each floor is the low quarter of one tier (D-799). These rules map the tiers to the power of a deck:

- Commander brackets 1 and 2 read the precons, and bracket 3 reads the EDHREC average decks.
- Commander bracket 4 reads the good lists, and bracket 5 reads the great lists.
- The fnm step reads the good lists, and the tournament step reads the great lists. Each cell takes the lower floor of Modern and Standard.
- A cell of fewer than 30 lists takes the floor of one color fewer. So a small sample never raises a floor.
- The casual step reads the 60-card precons. Their low quarter is 0 at two and three colors, so the casual step holds no floor.

| Power | 2 colors | 3 colors | 4 colors | 5 colors |
|---|---|---|---|---|
| Commander bracket 1 and 2 | 7 | 13 | 13 | 13 |
| Commander bracket 3 | 11 | 17 | 17 | 17 |
| Commander bracket 4 and 5 | 9 | 12 | 18 | 21 |
| 60-card fnm | 7 | 10 | 12 | 12 |
| 60-card tournament | 7 | 11 | 13 | 12 |

The floor of bracket 3 is above the floor of bracket 4 at two and three colors. The land band of bracket 3 is 34 to 38, and the land band of bracket 5 is 27 to 33. So a list of a high bracket plays fewer lands of every kind, and fixing lands too.

`go/internal/profile/bands.json` holds the floors under `fixing_land`.

## 4. The replay of the mana pass

`make manapass-check` replays the mana pass over the stored decks of a gate document, with no provider call. The session ran it on runs 29 and 31, on `main` and on this branch.

- Deck 13 of run 31 went from 36 basic lands to 25. It met its floor of 11, and its worst color rose from 1.02 to 1.33 of its need.
- Against the pass of `main`, the worst color rose on 8 decks and fell on none.
- No deck changed a band other than `fixing_land`.
- Decks 19 and 22 are set-family builds. Their pools hold fewer fixing lands than the floor, so each stays under it. A profile finding alone buys no repair turn (D-609), so the warning costs no model call.
