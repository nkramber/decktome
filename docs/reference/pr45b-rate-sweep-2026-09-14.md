# PR-45b: the rate sweep of the bracket 4 and 5 shortlist

Run date: 2026-09-14. Commit: `f45791a`. Card snapshot: `20260904T210157`. Quality model: `20260914T154223Z`, from the local meta refresh of 2026-09-14.

## What this reads

At brackets 4 and 5 the shortlist reads the top-list rate of each card (D-704, D-707). The sweep builds the shortlist of bracket gate prompts 10 to 15 at 55 points of a grid, and it calls no model.

- **Weight** is the weight of the rate in the score. The weight before PR-45b is 0.1.
- **Keep** is the rate at which a card with no theme signal stays on the list and keeps its full score. "none" keeps no card by its rate.
- **Pin** lets a power card at the keep rate skip the cap of its role and add to the total (D-710, D-712).
- **Cards** counts the list. **Lands** counts its lands, and **Fixing** counts the lands that make two or more colors of the deck.
- **In** counts the cards that the list before PR-45b does not hold. **Theme out** counts the on-theme cards of the list before PR-45b that the point drops. **Fixing lands out** sums the fixing lands that each list loses against the list before PR-45b.
- **Kept** counts the cards with no theme signal and no staple role that carry the rate signal, the pinned cards included. **Pinned** counts the pinned cards. **Other** counts the cards of the role other, whose cap of 10 binds the unpinned cards alone.
- **Tutors**, **Fast mana**, and **Game Changers** count by the rules of the profile (`profile.PowerOf`). A floor counts as met when the list holds the floor, or every power card with a rate that the colors reach.
- **Reach with a rate** counts the power cards in the colors that the top lists play at a rate above zero.

The card table of each prompt lists the 15 power cards with the highest rate, and the lists that hold each one.

The command, from `go/`: `CARDS_SNAPSHOT_DIR=<local snapshot> go run ./cmd/bracket-gate -sweep -only 10,11,12,13,14,15`. The output below comes from commit `f45791a`, and a run before the commit gave the same output, byte for byte.

Live Scryfall agreed with the Game Changer flag of the snapshot for each Game Changer that the card tables name, read 2026-09-14 at 16:04 UTC.

## Read

- **With no pin, the role caps block the power cards (F-131).** Fast mana takes the role ramp, and it competes with on-theme ramp for 30 places. On the Korvold list, Sol Ring, Lotus Petal, Chrome Mox, Mana Vault, and Mox Diamond pass the keep rate, and no weight under 1 lists them. An off-theme tutor takes the role other, which holds 10 places.
- **With no pin, only weight 1 meets all 18 floors, and 138 on-theme cards leave.** Weight 0.5 with a keep rate of 0.3 meets 16 floors, and 73 on-theme cards leave.
- **The first pin took places under the total, and 52 fixing lands left (F-132).** At weight 0.1 and a keep rate of 0.3, 18 on-theme cards also left, and that count was the only cost the first sweep read. The Korvold list kept 2 of 22 fixing lands, the Kinnan list 1 of 17, and the Najeela list 2 of 18.
- **With pins added to the total, weight 0.1 and a keep rate of 0.3 meet all 18 floors.** No on-theme card leaves, and 7 fixing lands leave, all from the Najeela list. The owner chose this point (D-711, D-712).
- **The keep rate, and not the pin, costs those 7 fixing lands.** The Najeela list takes 7 unpinned cards that the keep rate holds and that count toward no floor. They push out the 7 lowest-scored cards of the list before PR-45b, and each one is a fixing land.
- **More weight costs on-theme cards and meets no more floors.** At a keep rate of 0.3 with the pin, weight 0.3 drops 19 on-theme cards, and weight 0.5 drops 39.
- **A lower keep rate costs fixing lands.** At weight 0.1 with the pin, a keep rate of 0.1 brings in 283 cards, and 25 fixing lands leave.
- **The lists grow.** At the chosen point the lists at the cap hold 315 to 332 cards: Urza 315, Korvold 322, Kinnan 319, and Najeela 332. The Prosper list goes from 185 to 207 cards, and the Yuriko list goes from 210 to 235.

## Prompt 10: bracket 4 Urza, Lord High Artificer (artifacts)

Floors: tutors 2, fast mana 3, Game Changers 4. Reach with a rate: tutors 116, fast mana 40, Game Changers 22.

Rates of the power cards with a rate: highest 1.000, upper quarter 0.011, median 0.000, lower quarter 0.000, over 171 cards.

| Card | Rate | Role | Theme | Counts toward | Before | 0.1, keep 0.3, pin | 0.3, keep 0.3, pin | 0.3, keep 0.3 | 0.5, keep 0.3 | 1, keep 0.3 |
|---|---|---|---|---|---|---|---|---|---|---|
| Sol Ring | 1.000 | ramp | yes | fast mana | no | yes | yes | yes | yes | yes |
| Lotus Petal | 0.913 | ramp | yes | fast mana | no | yes | yes | no | yes | yes |
| Chrome Mox | 0.897 | ramp | yes | fast mana, Game Changers | yes | yes | yes | yes | yes | yes |
| Mana Vault | 0.857 | ramp | yes | fast mana, Game Changers | yes | yes | yes | yes | yes | yes |
| Mox Diamond | 0.841 | ramp | yes | fast mana, Game Changers | yes | yes | yes | yes | yes | yes |
| Mox Amber | 0.656 | ramp | yes | fast mana | no | yes | yes | yes | yes | yes |
| Force of Will | 0.648 | interaction | no | Game Changers | no | yes | yes | no | no | yes |
| Rhystic Study | 0.648 | draw | no | Game Changers | no | yes | yes | no | no | yes |
| Fierce Guardianship | 0.630 | interaction | no | Game Changers | no | yes | yes | no | no | yes |
| Mox Opal | 0.577 | ramp | yes | fast mana | yes | yes | yes | yes | yes | yes |
| Lion's Eye Diamond | 0.489 | ramp | yes | fast mana, Game Changers | no | yes | yes | no | no | yes |
| An Offer You Can't Refuse | 0.467 | interaction | yes | fast mana | no | yes | yes | no | no | yes |
| Mystical Tutor | 0.386 | other | no | tutors, Game Changers | no | yes | yes | yes | yes | yes |
| Thassa's Oracle | 0.328 | wincon | no | Game Changers | no | yes | yes | yes | yes | yes |
| Springleaf Drum | 0.301 | ramp | yes | fast mana | no | yes | yes | no | no | yes |

| Weight | Keep | Pin | Cards | Lands | Fixing | In | Theme out | Kept | Pinned | Other | Tutors | Fast mana | Game Changers | Floors met |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| 0.1 | none | no | 297 | 40 | 0 | 0 | 0 | 0 | 0 | 0 | 13 | 5 | 4 | 3 of 3 |
| 0.1 | 0.5 | no | 297 | 40 | 0 | 0 | 0 | 0 | 0 | 0 | 13 | 5 | 4 | 3 of 3 |
| 0.1 | 0.5 | yes | 307 | 40 | 0 | 10 | 0 | 0 | 10 | 0 | 13 | 8 | 7 | 3 of 3 |
| 0.1 | 0.3 | no | 300 | 38 | 0 | 5 | 0 | 5 | 0 | 4 | 14 | 5 | 6 | 3 of 3 |
| 0.1 | 0.3 | yes | 315 | 40 | 0 | 18 | 0 | 5 | 15 | 4 | 14 | 11 | 10 | 3 of 3 |
| 0.1 | 0.2 | no | 300 | 37 | 0 | 6 | 0 | 6 | 0 | 5 | 14 | 5 | 6 | 3 of 3 |
| 0.1 | 0.2 | yes | 317 | 39 | 0 | 21 | 0 | 6 | 17 | 5 | 14 | 11 | 12 | 3 of 3 |
| 0.1 | 0.1 | no | 300 | 32 | 0 | 12 | 1 | 11 | 0 | 10 | 15 | 5 | 7 | 3 of 3 |
| 0.1 | 0.1 | yes | 322 | 33 | 0 | 32 | 0 | 14 | 22 | 13 | 16 | 12 | 14 | 3 of 3 |
| 0.1 | 0.05 | no | 300 | 32 | 0 | 12 | 1 | 11 | 0 | 10 | 15 | 5 | 7 | 3 of 3 |
| 0.1 | 0.05 | yes | 328 | 33 | 0 | 38 | 0 | 16 | 28 | 15 | 19 | 12 | 16 | 3 of 3 |
| 0.2 | none | no | 297 | 40 | 0 | 1 | 1 | 0 | 0 | 0 | 13 | 6 | 4 | 3 of 3 |
| 0.2 | 0.5 | no | 297 | 40 | 0 | 1 | 1 | 0 | 0 | 0 | 13 | 6 | 4 | 3 of 3 |
| 0.2 | 0.5 | yes | 307 | 40 | 0 | 10 | 0 | 0 | 10 | 0 | 13 | 8 | 7 | 3 of 3 |
| 0.2 | 0.3 | no | 300 | 38 | 0 | 7 | 2 | 5 | 0 | 4 | 14 | 6 | 6 | 3 of 3 |
| 0.2 | 0.3 | yes | 315 | 40 | 0 | 19 | 1 | 5 | 15 | 4 | 14 | 11 | 10 | 3 of 3 |
| 0.2 | 0.2 | no | 300 | 37 | 0 | 8 | 2 | 6 | 0 | 5 | 14 | 6 | 6 | 3 of 3 |
| 0.2 | 0.2 | yes | 317 | 39 | 0 | 22 | 1 | 6 | 17 | 5 | 14 | 11 | 12 | 3 of 3 |
| 0.2 | 0.1 | no | 300 | 32 | 0 | 14 | 3 | 11 | 0 | 10 | 16 | 6 | 8 | 3 of 3 |
| 0.2 | 0.1 | yes | 322 | 33 | 0 | 33 | 1 | 14 | 22 | 13 | 16 | 12 | 14 | 3 of 3 |
| 0.2 | 0.05 | no | 300 | 32 | 0 | 14 | 3 | 11 | 0 | 10 | 16 | 6 | 8 | 3 of 3 |
| 0.2 | 0.05 | yes | 328 | 33 | 0 | 39 | 1 | 16 | 28 | 15 | 19 | 12 | 16 | 3 of 3 |
| 0.3 | none | no | 297 | 40 | 0 | 4 | 4 | 0 | 0 | 0 | 13 | 7 | 4 | 3 of 3 |
| 0.3 | 0.5 | no | 297 | 40 | 0 | 4 | 4 | 0 | 0 | 0 | 13 | 7 | 4 | 3 of 3 |
| 0.3 | 0.5 | yes | 307 | 40 | 0 | 11 | 1 | 0 | 10 | 0 | 13 | 8 | 8 | 3 of 3 |
| 0.3 | 0.3 | no | 300 | 38 | 0 | 11 | 6 | 5 | 0 | 4 | 14 | 7 | 6 | 3 of 3 |
| 0.3 | 0.3 | yes | 315 | 40 | 0 | 21 | 3 | 5 | 15 | 4 | 14 | 11 | 11 | 3 of 3 |
| 0.3 | 0.2 | no | 300 | 37 | 0 | 12 | 6 | 6 | 0 | 5 | 14 | 7 | 6 | 3 of 3 |
| 0.3 | 0.2 | yes | 317 | 39 | 0 | 24 | 3 | 6 | 17 | 5 | 14 | 11 | 12 | 3 of 3 |
| 0.3 | 0.1 | no | 300 | 32 | 0 | 17 | 6 | 11 | 0 | 10 | 16 | 7 | 8 | 3 of 3 |
| 0.3 | 0.1 | yes | 322 | 33 | 0 | 35 | 3 | 14 | 22 | 13 | 16 | 12 | 14 | 3 of 3 |
| 0.3 | 0.05 | no | 300 | 32 | 0 | 17 | 6 | 11 | 0 | 10 | 16 | 7 | 8 | 3 of 3 |
| 0.3 | 0.05 | yes | 328 | 33 | 0 | 41 | 3 | 16 | 28 | 15 | 19 | 12 | 16 | 3 of 3 |
| 0.5 | none | no | 297 | 40 | 0 | 6 | 6 | 0 | 0 | 0 | 13 | 8 | 5 | 3 of 3 |
| 0.5 | 0.5 | no | 297 | 40 | 0 | 6 | 6 | 0 | 0 | 0 | 13 | 8 | 5 | 3 of 3 |
| 0.5 | 0.5 | yes | 307 | 40 | 0 | 11 | 1 | 0 | 10 | 0 | 13 | 8 | 8 | 3 of 3 |
| 0.5 | 0.3 | no | 300 | 38 | 0 | 13 | 8 | 5 | 0 | 4 | 14 | 8 | 7 | 3 of 3 |
| 0.5 | 0.3 | yes | 315 | 40 | 0 | 21 | 3 | 5 | 15 | 4 | 14 | 11 | 11 | 3 of 3 |
| 0.5 | 0.2 | no | 300 | 37 | 0 | 15 | 9 | 6 | 0 | 5 | 14 | 8 | 7 | 3 of 3 |
| 0.5 | 0.2 | yes | 317 | 39 | 0 | 25 | 4 | 6 | 17 | 5 | 14 | 11 | 12 | 3 of 3 |
| 0.5 | 0.1 | no | 300 | 32 | 0 | 21 | 10 | 11 | 0 | 10 | 16 | 8 | 9 | 3 of 3 |
| 0.5 | 0.1 | yes | 322 | 33 | 0 | 37 | 5 | 14 | 22 | 13 | 16 | 12 | 14 | 3 of 3 |
| 0.5 | 0.05 | no | 300 | 32 | 0 | 21 | 10 | 11 | 0 | 10 | 16 | 8 | 9 | 3 of 3 |
| 0.5 | 0.05 | yes | 328 | 33 | 0 | 43 | 5 | 16 | 28 | 15 | 19 | 12 | 16 | 3 of 3 |
| 1 | none | no | 297 | 40 | 0 | 11 | 11 | 0 | 0 | 0 | 12 | 11 | 6 | 3 of 3 |
| 1 | 0.5 | no | 297 | 40 | 0 | 20 | 20 | 0 | 0 | 0 | 12 | 11 | 9 | 3 of 3 |
| 1 | 0.5 | yes | 307 | 40 | 0 | 21 | 11 | 0 | 10 | 0 | 13 | 11 | 9 | 3 of 3 |
| 1 | 0.3 | no | 300 | 38 | 0 | 26 | 21 | 5 | 0 | 4 | 13 | 11 | 11 | 3 of 3 |
| 1 | 0.3 | yes | 315 | 40 | 0 | 30 | 12 | 5 | 15 | 4 | 14 | 11 | 11 | 3 of 3 |
| 1 | 0.2 | no | 300 | 37 | 0 | 28 | 22 | 6 | 0 | 5 | 13 | 11 | 11 | 3 of 3 |
| 1 | 0.2 | yes | 317 | 39 | 0 | 33 | 12 | 6 | 17 | 5 | 14 | 11 | 12 | 3 of 3 |
| 1 | 0.1 | no | 300 | 32 | 0 | 34 | 23 | 11 | 0 | 10 | 14 | 11 | 13 | 3 of 3 |
| 1 | 0.1 | yes | 322 | 33 | 0 | 45 | 13 | 14 | 22 | 13 | 16 | 12 | 14 | 3 of 3 |
| 1 | 0.05 | no | 300 | 32 | 0 | 35 | 24 | 11 | 0 | 10 | 14 | 11 | 13 | 3 of 3 |
| 1 | 0.05 | yes | 328 | 33 | 0 | 52 | 14 | 16 | 28 | 15 | 19 | 12 | 16 | 3 of 3 |

## Prompt 11: bracket 4 Korvold, Fae-Cursed King (treasure sacrifice)

Floors: tutors 2, fast mana 3, Game Changers 4. Reach with a rate: tutors 280, fast mana 63, Game Changers 32.

Rates of the power cards with a rate: highest 1.000, upper quarter 0.008, median 0.000, lower quarter 0.000, over 363 cards.

| Card | Rate | Role | Theme | Counts toward | Before | 0.1, keep 0.3, pin | 0.3, keep 0.3, pin | 0.3, keep 0.3 | 0.5, keep 0.3 | 1, keep 0.3 |
|---|---|---|---|---|---|---|---|---|---|---|
| Sol Ring | 1.000 | ramp | no | fast mana | no | yes | yes | no | no | yes |
| Lotus Petal | 0.913 | ramp | no | fast mana | no | yes | yes | no | no | yes |
| Chrome Mox | 0.897 | ramp | no | fast mana, Game Changers | no | yes | yes | no | no | yes |
| Mana Vault | 0.857 | ramp | no | fast mana, Game Changers | no | yes | yes | no | no | yes |
| Mox Diamond | 0.841 | ramp | no | fast mana, Game Changers | no | yes | yes | no | no | yes |
| Mox Amber | 0.656 | ramp | no | fast mana | no | yes | yes | no | no | yes |
| Mox Opal | 0.577 | ramp | no | fast mana | no | yes | yes | no | no | no |
| Vampiric Tutor | 0.510 | other | no | tutors, Game Changers | no | yes | yes | yes | yes | yes |
| Demonic Tutor | 0.496 | other | no | tutors, Game Changers | no | yes | yes | yes | yes | yes |
| Lion's Eye Diamond | 0.489 | ramp | no | fast mana, Game Changers | no | yes | yes | no | no | no |
| Gamble | 0.446 | other | no | tutors, Game Changers | no | yes | yes | yes | yes | yes |
| Dark Ritual | 0.430 | ramp | no | fast mana | no | yes | yes | no | no | no |
| Rite of Flame | 0.424 | ramp | no | fast mana | no | yes | yes | no | no | no |
| Underworld Breach | 0.419 | other | no | Game Changers | no | yes | yes | yes | yes | yes |
| Imperial Seal | 0.401 | other | no | tutors, Game Changers | no | yes | yes | yes | yes | yes |

| Weight | Keep | Pin | Cards | Lands | Fixing | In | Theme out | Kept | Pinned | Other | Tutors | Fast mana | Game Changers | Floors met |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| 0.1 | none | no | 296 | 40 | 22 | 0 | 0 | 0 | 0 | 0 | 11 | 0 | 0 | 1 of 3 |
| 0.1 | 0.5 | no | 298 | 40 | 22 | 2 | 0 | 2 | 0 | 2 | 12 | 0 | 1 | 1 of 3 |
| 0.1 | 0.5 | yes | 305 | 40 | 22 | 9 | 0 | 2 | 8 | 2 | 12 | 7 | 4 | 3 of 3 |
| 0.1 | 0.3 | no | 300 | 34 | 16 | 10 | 0 | 10 | 0 | 10 | 17 | 0 | 6 | 2 of 3 |
| 0.1 | 0.3 | yes | 322 | 40 | 22 | 26 | 0 | 13 | 23 | 13 | 19 | 11 | 12 | 3 of 3 |
| 0.1 | 0.2 | no | 300 | 34 | 16 | 10 | 0 | 10 | 0 | 10 | 17 | 0 | 6 | 2 of 3 |
| 0.1 | 0.2 | yes | 333 | 39 | 21 | 38 | 0 | 20 | 33 | 20 | 23 | 12 | 17 | 3 of 3 |
| 0.1 | 0.1 | no | 300 | 34 | 16 | 10 | 0 | 10 | 0 | 10 | 17 | 0 | 6 | 2 of 3 |
| 0.1 | 0.1 | yes | 344 | 34 | 16 | 54 | 0 | 31 | 44 | 31 | 28 | 16 | 19 | 3 of 3 |
| 0.1 | 0.05 | no | 300 | 34 | 16 | 10 | 0 | 10 | 0 | 10 | 17 | 0 | 6 | 2 of 3 |
| 0.1 | 0.05 | yes | 354 | 34 | 16 | 64 | 0 | 37 | 54 | 37 | 35 | 19 | 20 | 3 of 3 |
| 0.2 | none | no | 296 | 40 | 22 | 0 | 0 | 0 | 0 | 0 | 11 | 0 | 0 | 1 of 3 |
| 0.2 | 0.5 | no | 298 | 40 | 22 | 2 | 0 | 2 | 0 | 2 | 12 | 0 | 1 | 1 of 3 |
| 0.2 | 0.5 | yes | 305 | 40 | 22 | 9 | 0 | 2 | 8 | 2 | 12 | 7 | 4 | 3 of 3 |
| 0.2 | 0.3 | no | 300 | 34 | 16 | 10 | 0 | 10 | 0 | 10 | 17 | 0 | 6 | 2 of 3 |
| 0.2 | 0.3 | yes | 322 | 40 | 22 | 26 | 0 | 13 | 23 | 13 | 19 | 11 | 12 | 3 of 3 |
| 0.2 | 0.2 | no | 300 | 34 | 16 | 10 | 0 | 10 | 0 | 10 | 17 | 0 | 6 | 2 of 3 |
| 0.2 | 0.2 | yes | 333 | 39 | 21 | 38 | 0 | 20 | 33 | 20 | 23 | 12 | 17 | 3 of 3 |
| 0.2 | 0.1 | no | 300 | 34 | 16 | 10 | 0 | 10 | 0 | 10 | 17 | 0 | 6 | 2 of 3 |
| 0.2 | 0.1 | yes | 344 | 34 | 16 | 54 | 0 | 31 | 44 | 31 | 28 | 16 | 19 | 3 of 3 |
| 0.2 | 0.05 | no | 300 | 34 | 16 | 10 | 0 | 10 | 0 | 10 | 17 | 0 | 6 | 2 of 3 |
| 0.2 | 0.05 | yes | 354 | 34 | 16 | 64 | 0 | 37 | 54 | 37 | 35 | 19 | 20 | 3 of 3 |
| 0.3 | none | no | 296 | 40 | 22 | 1 | 1 | 0 | 0 | 0 | 11 | 0 | 0 | 1 of 3 |
| 0.3 | 0.5 | no | 298 | 40 | 22 | 3 | 1 | 2 | 0 | 2 | 12 | 0 | 1 | 1 of 3 |
| 0.3 | 0.5 | yes | 305 | 40 | 22 | 10 | 1 | 2 | 8 | 2 | 12 | 7 | 4 | 3 of 3 |
| 0.3 | 0.3 | no | 300 | 34 | 16 | 11 | 1 | 10 | 0 | 10 | 17 | 0 | 6 | 2 of 3 |
| 0.3 | 0.3 | yes | 322 | 40 | 22 | 27 | 1 | 13 | 23 | 13 | 19 | 11 | 12 | 3 of 3 |
| 0.3 | 0.2 | no | 300 | 34 | 16 | 11 | 1 | 10 | 0 | 10 | 17 | 0 | 6 | 2 of 3 |
| 0.3 | 0.2 | yes | 333 | 39 | 21 | 38 | 0 | 20 | 33 | 20 | 23 | 12 | 17 | 3 of 3 |
| 0.3 | 0.1 | no | 300 | 34 | 16 | 11 | 1 | 10 | 0 | 10 | 17 | 0 | 6 | 2 of 3 |
| 0.3 | 0.1 | yes | 344 | 34 | 16 | 54 | 0 | 31 | 44 | 31 | 28 | 16 | 19 | 3 of 3 |
| 0.3 | 0.05 | no | 300 | 34 | 16 | 11 | 1 | 10 | 0 | 10 | 17 | 0 | 6 | 2 of 3 |
| 0.3 | 0.05 | yes | 354 | 34 | 16 | 64 | 0 | 37 | 54 | 37 | 35 | 19 | 20 | 3 of 3 |
| 0.5 | none | no | 296 | 40 | 22 | 3 | 3 | 0 | 0 | 0 | 11 | 0 | 0 | 1 of 3 |
| 0.5 | 0.5 | no | 298 | 40 | 22 | 5 | 3 | 2 | 0 | 2 | 12 | 0 | 1 | 1 of 3 |
| 0.5 | 0.5 | yes | 305 | 40 | 22 | 12 | 3 | 2 | 8 | 2 | 12 | 7 | 4 | 3 of 3 |
| 0.5 | 0.3 | no | 300 | 34 | 16 | 13 | 3 | 10 | 0 | 10 | 17 | 0 | 6 | 2 of 3 |
| 0.5 | 0.3 | yes | 322 | 40 | 22 | 29 | 3 | 13 | 23 | 13 | 19 | 11 | 12 | 3 of 3 |
| 0.5 | 0.2 | no | 300 | 34 | 16 | 13 | 3 | 10 | 0 | 10 | 17 | 0 | 6 | 2 of 3 |
| 0.5 | 0.2 | yes | 333 | 39 | 21 | 40 | 2 | 20 | 33 | 20 | 23 | 12 | 17 | 3 of 3 |
| 0.5 | 0.1 | no | 300 | 34 | 16 | 13 | 3 | 10 | 0 | 10 | 17 | 0 | 6 | 2 of 3 |
| 0.5 | 0.1 | yes | 344 | 34 | 16 | 56 | 2 | 31 | 44 | 31 | 28 | 16 | 19 | 3 of 3 |
| 0.5 | 0.05 | no | 300 | 34 | 16 | 13 | 3 | 10 | 0 | 10 | 17 | 0 | 6 | 2 of 3 |
| 0.5 | 0.05 | yes | 354 | 34 | 16 | 66 | 2 | 37 | 54 | 37 | 35 | 19 | 20 | 3 of 3 |
| 1 | none | no | 296 | 40 | 22 | 4 | 4 | 0 | 0 | 0 | 11 | 1 | 0 | 1 of 3 |
| 1 | 0.5 | no | 298 | 40 | 22 | 13 | 11 | 2 | 0 | 2 | 12 | 7 | 4 | 3 of 3 |
| 1 | 0.5 | yes | 305 | 40 | 22 | 14 | 5 | 2 | 8 | 2 | 12 | 8 | 4 | 3 of 3 |
| 1 | 0.3 | no | 300 | 34 | 16 | 25 | 15 | 10 | 0 | 10 | 17 | 7 | 9 | 3 of 3 |
| 1 | 0.3 | yes | 322 | 40 | 22 | 35 | 9 | 13 | 23 | 13 | 19 | 12 | 12 | 3 of 3 |
| 1 | 0.2 | no | 300 | 34 | 16 | 26 | 16 | 10 | 0 | 10 | 17 | 7 | 9 | 3 of 3 |
| 1 | 0.2 | yes | 333 | 39 | 21 | 47 | 9 | 20 | 33 | 20 | 23 | 12 | 17 | 3 of 3 |
| 1 | 0.1 | no | 300 | 34 | 16 | 26 | 16 | 10 | 0 | 10 | 17 | 7 | 9 | 3 of 3 |
| 1 | 0.1 | yes | 344 | 34 | 16 | 62 | 8 | 31 | 44 | 31 | 28 | 16 | 19 | 3 of 3 |
| 1 | 0.05 | no | 300 | 34 | 16 | 26 | 16 | 10 | 0 | 10 | 17 | 7 | 9 | 3 of 3 |
| 1 | 0.05 | yes | 354 | 34 | 16 | 72 | 8 | 37 | 54 | 37 | 35 | 19 | 20 | 3 of 3 |

## Prompt 12: bracket 4 Prosper, Tome-Bound (treasure)

Floors: tutors 2, fast mana 3, Game Changers 4. Reach with a rate: tutors 185, fast mana 50, Game Changers 25.

Rates of the power cards with a rate: highest 1.000, upper quarter 0.008, median 0.000, lower quarter 0.000, over 251 cards.

| Card | Rate | Role | Theme | Counts toward | Before | 0.1, keep 0.3, pin | 0.3, keep 0.3, pin | 0.3, keep 0.3 | 0.5, keep 0.3 | 1, keep 0.3 |
|---|---|---|---|---|---|---|---|---|---|---|
| Sol Ring | 1.000 | ramp | no | fast mana | no | yes | yes | no | yes | yes |
| Lotus Petal | 0.913 | ramp | no | fast mana | no | yes | yes | no | yes | yes |
| Chrome Mox | 0.897 | ramp | no | fast mana, Game Changers | no | yes | yes | no | yes | yes |
| Mana Vault | 0.857 | ramp | no | fast mana, Game Changers | no | yes | yes | no | yes | yes |
| Mox Diamond | 0.841 | ramp | no | fast mana, Game Changers | no | yes | yes | no | yes | yes |
| Mox Amber | 0.656 | ramp | no | fast mana | no | yes | yes | no | no | yes |
| Mox Opal | 0.577 | ramp | no | fast mana | no | yes | yes | no | no | yes |
| Vampiric Tutor | 0.510 | other | no | tutors, Game Changers | no | yes | yes | yes | yes | yes |
| Demonic Tutor | 0.496 | other | no | tutors, Game Changers | no | yes | yes | yes | yes | yes |
| Lion's Eye Diamond | 0.489 | ramp | no | fast mana, Game Changers | no | yes | yes | no | no | no |
| Gamble | 0.446 | other | no | tutors, Game Changers | no | yes | yes | yes | yes | yes |
| Dark Ritual | 0.430 | ramp | no | fast mana | no | yes | yes | no | no | no |
| Rite of Flame | 0.424 | ramp | no | fast mana | no | yes | yes | no | no | no |
| Underworld Breach | 0.419 | other | no | Game Changers | no | yes | yes | yes | yes | yes |
| Imperial Seal | 0.401 | other | no | tutors, Game Changers | no | yes | yes | yes | yes | yes |

| Weight | Keep | Pin | Cards | Lands | Fixing | In | Theme out | Kept | Pinned | Other | Tutors | Fast mana | Game Changers | Floors met |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| 0.1 | none | no | 185 | 40 | 34 | 0 | 0 | 0 | 0 | 0 | 1 | 0 | 3 | 0 of 3 |
| 0.1 | 0.5 | no | 187 | 40 | 34 | 2 | 0 | 2 | 0 | 2 | 2 | 0 | 4 | 2 of 3 |
| 0.1 | 0.5 | yes | 194 | 40 | 34 | 9 | 0 | 2 | 8 | 2 | 2 | 7 | 7 | 3 of 3 |
| 0.1 | 0.3 | no | 194 | 40 | 34 | 9 | 0 | 9 | 0 | 9 | 6 | 0 | 8 | 2 of 3 |
| 0.1 | 0.3 | yes | 207 | 40 | 34 | 22 | 0 | 9 | 19 | 9 | 6 | 11 | 13 | 3 of 3 |
| 0.1 | 0.2 | no | 195 | 40 | 34 | 10 | 0 | 10 | 0 | 10 | 7 | 0 | 9 | 2 of 3 |
| 0.1 | 0.2 | yes | 217 | 40 | 34 | 32 | 0 | 15 | 26 | 15 | 8 | 12 | 15 | 3 of 3 |
| 0.1 | 0.1 | no | 195 | 40 | 34 | 10 | 0 | 10 | 0 | 10 | 7 | 0 | 9 | 2 of 3 |
| 0.1 | 0.1 | yes | 227 | 40 | 34 | 42 | 0 | 22 | 32 | 22 | 11 | 15 | 16 | 3 of 3 |
| 0.1 | 0.05 | no | 195 | 40 | 34 | 14 | 0 | 10 | 0 | 10 | 7 | 0 | 9 | 2 of 3 |
| 0.1 | 0.05 | yes | 235 | 40 | 34 | 52 | 0 | 28 | 40 | 28 | 17 | 17 | 16 | 3 of 3 |
| 0.2 | none | no | 185 | 40 | 34 | 5 | 0 | 0 | 0 | 0 | 1 | 0 | 3 | 0 of 3 |
| 0.2 | 0.5 | no | 187 | 40 | 34 | 7 | 0 | 2 | 0 | 2 | 2 | 0 | 4 | 2 of 3 |
| 0.2 | 0.5 | yes | 194 | 40 | 34 | 14 | 0 | 2 | 8 | 2 | 2 | 7 | 7 | 3 of 3 |
| 0.2 | 0.3 | no | 194 | 40 | 34 | 14 | 0 | 9 | 0 | 9 | 6 | 0 | 8 | 2 of 3 |
| 0.2 | 0.3 | yes | 207 | 40 | 34 | 26 | 0 | 9 | 19 | 9 | 6 | 11 | 13 | 3 of 3 |
| 0.2 | 0.2 | no | 195 | 40 | 34 | 15 | 0 | 10 | 0 | 10 | 7 | 0 | 9 | 2 of 3 |
| 0.2 | 0.2 | yes | 217 | 40 | 34 | 34 | 0 | 15 | 26 | 15 | 8 | 12 | 15 | 3 of 3 |
| 0.2 | 0.1 | no | 195 | 40 | 34 | 15 | 0 | 10 | 0 | 10 | 7 | 0 | 9 | 2 of 3 |
| 0.2 | 0.1 | yes | 227 | 40 | 34 | 44 | 0 | 22 | 32 | 22 | 11 | 15 | 16 | 3 of 3 |
| 0.2 | 0.05 | no | 195 | 40 | 34 | 17 | 0 | 10 | 0 | 10 | 7 | 0 | 9 | 2 of 3 |
| 0.2 | 0.05 | yes | 235 | 40 | 34 | 55 | 0 | 28 | 40 | 28 | 17 | 17 | 16 | 3 of 3 |
| 0.3 | none | no | 185 | 40 | 34 | 8 | 0 | 0 | 0 | 0 | 1 | 0 | 3 | 0 of 3 |
| 0.3 | 0.5 | no | 187 | 40 | 34 | 10 | 0 | 2 | 0 | 2 | 2 | 0 | 4 | 2 of 3 |
| 0.3 | 0.5 | yes | 194 | 40 | 34 | 17 | 0 | 2 | 8 | 2 | 2 | 7 | 7 | 3 of 3 |
| 0.3 | 0.3 | no | 194 | 40 | 34 | 17 | 0 | 9 | 0 | 9 | 6 | 0 | 8 | 2 of 3 |
| 0.3 | 0.3 | yes | 207 | 40 | 34 | 30 | 0 | 9 | 19 | 9 | 6 | 11 | 13 | 3 of 3 |
| 0.3 | 0.2 | no | 195 | 40 | 34 | 18 | 0 | 10 | 0 | 10 | 7 | 0 | 8 | 2 of 3 |
| 0.3 | 0.2 | yes | 217 | 40 | 34 | 39 | 0 | 15 | 26 | 15 | 8 | 12 | 15 | 3 of 3 |
| 0.3 | 0.1 | no | 195 | 40 | 34 | 18 | 0 | 10 | 0 | 10 | 7 | 0 | 8 | 2 of 3 |
| 0.3 | 0.1 | yes | 227 | 40 | 34 | 49 | 0 | 22 | 32 | 22 | 11 | 15 | 16 | 3 of 3 |
| 0.3 | 0.05 | no | 195 | 40 | 34 | 20 | 0 | 10 | 0 | 10 | 7 | 0 | 8 | 2 of 3 |
| 0.3 | 0.05 | yes | 235 | 40 | 34 | 58 | 0 | 28 | 40 | 28 | 17 | 17 | 16 | 3 of 3 |
| 0.5 | none | no | 185 | 40 | 34 | 10 | 0 | 0 | 0 | 0 | 1 | 0 | 3 | 0 of 3 |
| 0.5 | 0.5 | no | 187 | 40 | 34 | 17 | 5 | 2 | 0 | 2 | 2 | 5 | 7 | 3 of 3 |
| 0.5 | 0.5 | yes | 194 | 40 | 34 | 19 | 0 | 2 | 8 | 2 | 2 | 7 | 7 | 3 of 3 |
| 0.5 | 0.3 | no | 194 | 40 | 34 | 24 | 5 | 9 | 0 | 9 | 6 | 5 | 11 | 3 of 3 |
| 0.5 | 0.3 | yes | 207 | 40 | 34 | 31 | 0 | 9 | 19 | 9 | 6 | 11 | 13 | 3 of 3 |
| 0.5 | 0.2 | no | 195 | 40 | 34 | 25 | 5 | 10 | 0 | 10 | 6 | 5 | 11 | 3 of 3 |
| 0.5 | 0.2 | yes | 217 | 40 | 34 | 41 | 0 | 15 | 26 | 15 | 8 | 12 | 15 | 3 of 3 |
| 0.5 | 0.1 | no | 195 | 40 | 34 | 25 | 5 | 10 | 0 | 10 | 6 | 5 | 11 | 3 of 3 |
| 0.5 | 0.1 | yes | 227 | 40 | 34 | 51 | 0 | 22 | 32 | 22 | 11 | 15 | 16 | 3 of 3 |
| 0.5 | 0.05 | no | 195 | 40 | 34 | 27 | 5 | 10 | 0 | 10 | 6 | 5 | 11 | 3 of 3 |
| 0.5 | 0.05 | yes | 235 | 40 | 34 | 60 | 0 | 28 | 40 | 28 | 17 | 17 | 16 | 3 of 3 |
| 1 | none | no | 185 | 40 | 34 | 18 | 1 | 0 | 0 | 0 | 1 | 1 | 3 | 0 of 3 |
| 1 | 0.5 | no | 187 | 40 | 34 | 27 | 8 | 2 | 0 | 2 | 2 | 8 | 7 | 3 of 3 |
| 1 | 0.5 | yes | 194 | 40 | 34 | 27 | 1 | 2 | 8 | 2 | 2 | 8 | 7 | 3 of 3 |
| 1 | 0.3 | no | 194 | 40 | 34 | 34 | 8 | 9 | 0 | 9 | 6 | 8 | 11 | 3 of 3 |
| 1 | 0.3 | yes | 207 | 40 | 34 | 40 | 2 | 9 | 19 | 9 | 6 | 12 | 13 | 3 of 3 |
| 1 | 0.2 | no | 195 | 40 | 34 | 35 | 8 | 10 | 0 | 10 | 6 | 8 | 11 | 3 of 3 |
| 1 | 0.2 | yes | 217 | 40 | 34 | 48 | 2 | 15 | 26 | 15 | 8 | 13 | 15 | 3 of 3 |
| 1 | 0.1 | no | 195 | 40 | 34 | 35 | 8 | 10 | 0 | 10 | 6 | 8 | 11 | 3 of 3 |
| 1 | 0.1 | yes | 227 | 40 | 34 | 58 | 2 | 22 | 32 | 22 | 11 | 16 | 16 | 3 of 3 |
| 1 | 0.05 | no | 195 | 40 | 34 | 35 | 8 | 10 | 0 | 10 | 6 | 8 | 11 | 3 of 3 |
| 1 | 0.05 | yes | 235 | 40 | 34 | 66 | 2 | 28 | 40 | 28 | 17 | 18 | 16 | 3 of 3 |

## Prompt 13: bracket 5 Kinnan, Bonder Prodigy (combo)

Floors: tutors 4, fast mana 6, Game Changers 8. Reach with a rate: tutors 207, fast mana 52, Game Changers 29.

Rates of the power cards with a rate: highest 1.000, upper quarter 0.010, median 0.000, lower quarter 0.000, over 278 cards.

| Card | Rate | Role | Theme | Counts toward | Before | 0.1, keep 0.3, pin | 0.3, keep 0.3, pin | 0.3, keep 0.3 | 0.5, keep 0.3 | 1, keep 0.3 |
|---|---|---|---|---|---|---|---|---|---|---|
| Sol Ring | 1.000 | ramp | no | fast mana | no | yes | yes | no | yes | yes |
| Lotus Petal | 0.913 | ramp | no | fast mana | no | yes | yes | no | yes | yes |
| Chrome Mox | 0.897 | ramp | no | fast mana, Game Changers | no | yes | yes | no | yes | yes |
| Mana Vault | 0.857 | ramp | yes | fast mana, Game Changers | yes | yes | yes | yes | yes | yes |
| Mox Diamond | 0.841 | ramp | no | fast mana, Game Changers | no | yes | yes | no | yes | yes |
| Mox Amber | 0.656 | ramp | no | fast mana | no | yes | yes | no | no | yes |
| Force of Will | 0.648 | interaction | no | Game Changers | no | yes | yes | no | yes | yes |
| Rhystic Study | 0.648 | draw | no | Game Changers | no | yes | yes | no | yes | yes |
| Fierce Guardianship | 0.630 | interaction | no | Game Changers | no | yes | yes | no | yes | yes |
| Mox Opal | 0.577 | ramp | no | fast mana | no | yes | yes | no | no | yes |
| Lion's Eye Diamond | 0.489 | ramp | no | fast mana, Game Changers | no | yes | yes | no | no | yes |
| An Offer You Can't Refuse | 0.467 | interaction | no | fast mana | no | yes | yes | no | no | yes |
| Mystical Tutor | 0.386 | synergy | yes | tutors, Game Changers | yes | yes | yes | yes | yes | yes |
| Crop Rotation | 0.358 | synergy | yes | Game Changers | yes | yes | yes | yes | yes | yes |
| Chord of Calling | 0.331 | synergy | yes | tutors | yes | yes | yes | yes | yes | yes |

| Weight | Keep | Pin | Cards | Lands | Fixing | In | Theme out | Kept | Pinned | Other | Tutors | Fast mana | Game Changers | Floors met |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| 0.1 | none | no | 300 | 36 | 17 | 0 | 0 | 0 | 0 | 0 | 26 | 1 | 8 | 2 of 3 |
| 0.1 | 0.5 | no | 300 | 36 | 17 | 0 | 0 | 0 | 0 | 0 | 26 | 1 | 8 | 2 of 3 |
| 0.1 | 0.5 | yes | 310 | 36 | 17 | 10 | 0 | 0 | 10 | 0 | 26 | 7 | 13 | 3 of 3 |
| 0.1 | 0.3 | no | 300 | 36 | 17 | 0 | 0 | 0 | 0 | 0 | 26 | 1 | 8 | 2 of 3 |
| 0.1 | 0.3 | yes | 319 | 36 | 17 | 19 | 0 | 0 | 19 | 0 | 29 | 10 | 15 | 3 of 3 |
| 0.1 | 0.2 | no | 300 | 34 | 15 | 2 | 0 | 2 | 0 | 2 | 26 | 1 | 8 | 2 of 3 |
| 0.1 | 0.2 | yes | 324 | 34 | 15 | 26 | 0 | 2 | 24 | 2 | 30 | 10 | 16 | 3 of 3 |
| 0.1 | 0.1 | no | 300 | 26 | 7 | 10 | 0 | 10 | 0 | 10 | 26 | 1 | 8 | 2 of 3 |
| 0.1 | 0.1 | yes | 334 | 27 | 8 | 43 | 0 | 10 | 34 | 10 | 31 | 12 | 16 | 3 of 3 |
| 0.1 | 0.05 | no | 300 | 26 | 7 | 10 | 0 | 10 | 0 | 10 | 26 | 1 | 8 | 2 of 3 |
| 0.1 | 0.05 | yes | 343 | 27 | 8 | 52 | 0 | 10 | 43 | 10 | 38 | 13 | 19 | 3 of 3 |
| 0.2 | none | no | 300 | 36 | 17 | 5 | 5 | 0 | 0 | 0 | 29 | 1 | 9 | 2 of 3 |
| 0.2 | 0.5 | no | 300 | 36 | 17 | 5 | 5 | 0 | 0 | 0 | 29 | 1 | 9 | 2 of 3 |
| 0.2 | 0.5 | yes | 310 | 36 | 17 | 14 | 4 | 0 | 10 | 0 | 29 | 7 | 14 | 3 of 3 |
| 0.2 | 0.3 | no | 300 | 36 | 17 | 5 | 5 | 0 | 0 | 0 | 29 | 1 | 9 | 2 of 3 |
| 0.2 | 0.3 | yes | 319 | 36 | 17 | 19 | 0 | 0 | 19 | 0 | 29 | 10 | 15 | 3 of 3 |
| 0.2 | 0.2 | no | 300 | 34 | 15 | 7 | 5 | 2 | 0 | 2 | 29 | 1 | 9 | 2 of 3 |
| 0.2 | 0.2 | yes | 324 | 34 | 15 | 26 | 0 | 2 | 24 | 2 | 29 | 10 | 16 | 3 of 3 |
| 0.2 | 0.1 | no | 300 | 26 | 7 | 15 | 5 | 10 | 0 | 10 | 29 | 1 | 9 | 2 of 3 |
| 0.2 | 0.1 | yes | 334 | 27 | 8 | 43 | 0 | 10 | 34 | 10 | 32 | 12 | 17 | 3 of 3 |
| 0.2 | 0.05 | no | 300 | 26 | 7 | 15 | 5 | 10 | 0 | 10 | 29 | 1 | 9 | 2 of 3 |
| 0.2 | 0.05 | yes | 343 | 27 | 8 | 52 | 0 | 10 | 43 | 10 | 38 | 13 | 19 | 3 of 3 |
| 0.3 | none | no | 300 | 35 | 16 | 10 | 9 | 0 | 0 | 0 | 29 | 1 | 9 | 2 of 3 |
| 0.3 | 0.5 | no | 300 | 35 | 16 | 10 | 9 | 0 | 0 | 0 | 29 | 1 | 9 | 2 of 3 |
| 0.3 | 0.5 | yes | 310 | 35 | 16 | 19 | 8 | 0 | 10 | 0 | 29 | 7 | 14 | 3 of 3 |
| 0.3 | 0.3 | no | 300 | 35 | 16 | 10 | 9 | 0 | 0 | 0 | 29 | 1 | 9 | 2 of 3 |
| 0.3 | 0.3 | yes | 319 | 35 | 16 | 26 | 6 | 0 | 19 | 0 | 30 | 10 | 15 | 3 of 3 |
| 0.3 | 0.2 | no | 300 | 34 | 15 | 11 | 9 | 2 | 0 | 2 | 29 | 1 | 9 | 2 of 3 |
| 0.3 | 0.2 | yes | 324 | 34 | 15 | 29 | 3 | 2 | 24 | 2 | 31 | 10 | 17 | 3 of 3 |
| 0.3 | 0.1 | no | 300 | 26 | 7 | 19 | 9 | 10 | 0 | 10 | 29 | 1 | 9 | 2 of 3 |
| 0.3 | 0.1 | yes | 334 | 27 | 8 | 43 | 0 | 10 | 34 | 10 | 32 | 12 | 17 | 3 of 3 |
| 0.3 | 0.05 | no | 300 | 26 | 7 | 19 | 9 | 10 | 0 | 10 | 29 | 1 | 9 | 2 of 3 |
| 0.3 | 0.05 | yes | 343 | 27 | 8 | 52 | 0 | 10 | 43 | 10 | 38 | 13 | 19 | 3 of 3 |
| 0.5 | none | no | 300 | 35 | 16 | 16 | 15 | 0 | 0 | 0 | 31 | 1 | 10 | 2 of 3 |
| 0.5 | 0.5 | no | 300 | 35 | 16 | 26 | 25 | 0 | 0 | 0 | 28 | 5 | 15 | 2 of 3 |
| 0.5 | 0.5 | yes | 310 | 35 | 16 | 28 | 17 | 0 | 10 | 0 | 30 | 7 | 15 | 3 of 3 |
| 0.5 | 0.3 | no | 300 | 35 | 16 | 27 | 26 | 0 | 0 | 0 | 28 | 5 | 15 | 2 of 3 |
| 0.5 | 0.3 | yes | 319 | 35 | 16 | 36 | 16 | 0 | 19 | 0 | 33 | 10 | 17 | 3 of 3 |
| 0.5 | 0.2 | no | 300 | 33 | 14 | 29 | 26 | 2 | 0 | 2 | 28 | 5 | 15 | 2 of 3 |
| 0.5 | 0.2 | yes | 324 | 33 | 14 | 39 | 12 | 2 | 24 | 2 | 33 | 10 | 18 | 3 of 3 |
| 0.5 | 0.1 | no | 300 | 26 | 7 | 36 | 26 | 10 | 0 | 10 | 28 | 5 | 15 | 2 of 3 |
| 0.5 | 0.1 | yes | 334 | 27 | 8 | 52 | 9 | 10 | 34 | 10 | 35 | 12 | 18 | 3 of 3 |
| 0.5 | 0.05 | no | 300 | 26 | 7 | 36 | 26 | 10 | 0 | 10 | 28 | 5 | 15 | 2 of 3 |
| 0.5 | 0.05 | yes | 343 | 27 | 8 | 57 | 5 | 10 | 43 | 10 | 37 | 13 | 19 | 3 of 3 |
| 1 | none | no | 300 | 35 | 16 | 26 | 25 | 0 | 0 | 0 | 37 | 2 | 11 | 2 of 3 |
| 1 | 0.5 | no | 300 | 35 | 16 | 40 | 39 | 0 | 0 | 0 | 34 | 8 | 16 | 3 of 3 |
| 1 | 0.5 | yes | 310 | 35 | 16 | 41 | 30 | 0 | 10 | 0 | 36 | 8 | 16 | 3 of 3 |
| 1 | 0.3 | no | 300 | 35 | 16 | 49 | 48 | 0 | 0 | 0 | 34 | 10 | 17 | 3 of 3 |
| 1 | 0.3 | yes | 319 | 35 | 16 | 53 | 33 | 0 | 19 | 0 | 35 | 10 | 17 | 3 of 3 |
| 1 | 0.2 | no | 300 | 33 | 14 | 52 | 49 | 2 | 0 | 2 | 34 | 10 | 17 | 3 of 3 |
| 1 | 0.2 | yes | 324 | 33 | 14 | 58 | 31 | 2 | 24 | 2 | 36 | 10 | 19 | 3 of 3 |
| 1 | 0.1 | no | 300 | 25 | 6 | 60 | 49 | 10 | 0 | 10 | 34 | 10 | 17 | 3 of 3 |
| 1 | 0.1 | yes | 334 | 26 | 7 | 70 | 26 | 10 | 34 | 10 | 38 | 12 | 19 | 3 of 3 |
| 1 | 0.05 | no | 300 | 25 | 6 | 60 | 49 | 10 | 0 | 10 | 34 | 10 | 17 | 3 of 3 |
| 1 | 0.05 | yes | 343 | 26 | 7 | 74 | 21 | 10 | 43 | 10 | 39 | 13 | 20 | 3 of 3 |

## Prompt 14: bracket 5 Yuriko, the Tiger's Shadow (ninjas)

Floors: tutors 4, fast mana 6, Game Changers 8. Reach with a rate: tutors 239, fast mana 47, Game Changers 33.

Rates of the power cards with a rate: highest 1.000, upper quarter 0.009, median 0.000, lower quarter 0.000, over 308 cards.

| Card | Rate | Role | Theme | Counts toward | Before | 0.1, keep 0.3, pin | 0.3, keep 0.3, pin | 0.3, keep 0.3 | 0.5, keep 0.3 | 1, keep 0.3 |
|---|---|---|---|---|---|---|---|---|---|---|
| Sol Ring | 1.000 | ramp | no | fast mana | yes | yes | yes | yes | yes | yes |
| Lotus Petal | 0.913 | ramp | no | fast mana | yes | yes | yes | yes | yes | yes |
| Chrome Mox | 0.897 | ramp | no | fast mana, Game Changers | yes | yes | yes | yes | yes | yes |
| Mana Vault | 0.857 | ramp | no | fast mana, Game Changers | yes | yes | yes | yes | yes | yes |
| Mox Diamond | 0.841 | ramp | no | fast mana, Game Changers | yes | yes | yes | yes | yes | yes |
| Mox Amber | 0.656 | ramp | no | fast mana | yes | yes | yes | yes | yes | yes |
| Force of Will | 0.648 | interaction | no | Game Changers | yes | yes | yes | yes | yes | yes |
| Rhystic Study | 0.648 | draw | no | Game Changers | yes | yes | yes | yes | yes | yes |
| Fierce Guardianship | 0.630 | interaction | no | Game Changers | yes | yes | yes | yes | yes | yes |
| Mox Opal | 0.577 | ramp | no | fast mana | yes | yes | yes | yes | yes | yes |
| Vampiric Tutor | 0.510 | other | no | tutors, Game Changers | no | yes | yes | yes | yes | yes |
| Demonic Tutor | 0.496 | other | no | tutors, Game Changers | no | yes | yes | yes | yes | yes |
| Lion's Eye Diamond | 0.489 | ramp | no | fast mana, Game Changers | yes | yes | yes | yes | yes | yes |
| An Offer You Can't Refuse | 0.467 | interaction | no | fast mana | yes | yes | yes | yes | yes | yes |
| Dark Ritual | 0.430 | ramp | no | fast mana | yes | yes | yes | yes | yes | yes |

| Weight | Keep | Pin | Cards | Lands | Fixing | In | Theme out | Kept | Pinned | Other | Tutors | Fast mana | Game Changers | Floors met |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| 0.1 | none | no | 210 | 40 | 40 | 0 | 0 | 0 | 0 | 0 | 1 | 12 | 13 | 2 of 3 |
| 0.1 | 0.5 | no | 211 | 40 | 40 | 1 | 0 | 1 | 0 | 1 | 2 | 12 | 14 | 2 of 3 |
| 0.1 | 0.5 | yes | 221 | 40 | 40 | 11 | 0 | 1 | 11 | 1 | 2 | 13 | 14 | 2 of 3 |
| 0.1 | 0.3 | no | 220 | 40 | 40 | 10 | 0 | 10 | 0 | 9 | 6 | 12 | 18 | 3 of 3 |
| 0.1 | 0.3 | yes | 235 | 40 | 40 | 25 | 0 | 10 | 21 | 9 | 6 | 13 | 18 | 3 of 3 |
| 0.1 | 0.2 | no | 221 | 40 | 40 | 12 | 0 | 11 | 0 | 10 | 7 | 12 | 18 | 3 of 3 |
| 0.1 | 0.2 | yes | 245 | 40 | 40 | 35 | 0 | 16 | 28 | 15 | 8 | 13 | 19 | 3 of 3 |
| 0.1 | 0.1 | no | 221 | 40 | 40 | 16 | 0 | 11 | 0 | 10 | 7 | 14 | 18 | 3 of 3 |
| 0.1 | 0.1 | yes | 256 | 40 | 40 | 46 | 0 | 24 | 36 | 23 | 13 | 15 | 21 | 3 of 3 |
| 0.1 | 0.05 | no | 221 | 40 | 40 | 25 | 0 | 11 | 0 | 10 | 7 | 15 | 18 | 3 of 3 |
| 0.1 | 0.05 | yes | 267 | 40 | 40 | 59 | 0 | 33 | 47 | 32 | 22 | 16 | 22 | 3 of 3 |
| 0.2 | none | no | 210 | 40 | 40 | 6 | 0 | 0 | 0 | 0 | 1 | 14 | 13 | 2 of 3 |
| 0.2 | 0.5 | no | 211 | 40 | 40 | 7 | 0 | 1 | 0 | 1 | 2 | 14 | 14 | 2 of 3 |
| 0.2 | 0.5 | yes | 221 | 40 | 40 | 14 | 0 | 1 | 11 | 1 | 2 | 14 | 14 | 2 of 3 |
| 0.2 | 0.3 | no | 220 | 40 | 40 | 16 | 0 | 10 | 0 | 9 | 6 | 14 | 18 | 3 of 3 |
| 0.2 | 0.3 | yes | 235 | 40 | 40 | 27 | 0 | 10 | 21 | 9 | 6 | 15 | 18 | 3 of 3 |
| 0.2 | 0.2 | no | 221 | 40 | 40 | 17 | 0 | 11 | 0 | 10 | 7 | 14 | 18 | 3 of 3 |
| 0.2 | 0.2 | yes | 245 | 40 | 40 | 37 | 0 | 16 | 28 | 15 | 9 | 15 | 19 | 3 of 3 |
| 0.2 | 0.1 | no | 221 | 40 | 40 | 19 | 0 | 11 | 0 | 10 | 7 | 14 | 18 | 3 of 3 |
| 0.2 | 0.1 | yes | 256 | 40 | 40 | 48 | 0 | 24 | 36 | 23 | 13 | 15 | 21 | 3 of 3 |
| 0.2 | 0.05 | no | 221 | 40 | 40 | 26 | 0 | 11 | 0 | 10 | 7 | 15 | 18 | 3 of 3 |
| 0.2 | 0.05 | yes | 267 | 40 | 40 | 60 | 0 | 33 | 47 | 32 | 22 | 16 | 22 | 3 of 3 |
| 0.3 | none | no | 210 | 40 | 40 | 11 | 0 | 0 | 0 | 0 | 1 | 14 | 13 | 2 of 3 |
| 0.3 | 0.5 | no | 211 | 40 | 40 | 12 | 0 | 1 | 0 | 1 | 2 | 14 | 14 | 2 of 3 |
| 0.3 | 0.5 | yes | 221 | 40 | 40 | 17 | 0 | 1 | 11 | 1 | 2 | 15 | 14 | 2 of 3 |
| 0.3 | 0.3 | no | 220 | 40 | 40 | 21 | 0 | 10 | 0 | 9 | 6 | 14 | 18 | 3 of 3 |
| 0.3 | 0.3 | yes | 235 | 40 | 40 | 30 | 0 | 10 | 21 | 9 | 6 | 15 | 18 | 3 of 3 |
| 0.3 | 0.2 | no | 221 | 40 | 40 | 22 | 0 | 11 | 0 | 10 | 7 | 14 | 18 | 3 of 3 |
| 0.3 | 0.2 | yes | 245 | 40 | 40 | 40 | 0 | 16 | 28 | 15 | 8 | 15 | 19 | 3 of 3 |
| 0.3 | 0.1 | no | 221 | 40 | 40 | 22 | 0 | 11 | 0 | 10 | 7 | 14 | 18 | 3 of 3 |
| 0.3 | 0.1 | yes | 256 | 40 | 40 | 51 | 0 | 24 | 36 | 23 | 14 | 15 | 21 | 3 of 3 |
| 0.3 | 0.05 | no | 221 | 40 | 40 | 27 | 0 | 11 | 0 | 10 | 7 | 15 | 18 | 3 of 3 |
| 0.3 | 0.05 | yes | 267 | 40 | 40 | 62 | 0 | 33 | 47 | 32 | 22 | 16 | 22 | 3 of 3 |
| 0.5 | none | no | 210 | 40 | 40 | 18 | 0 | 0 | 0 | 0 | 2 | 14 | 13 | 2 of 3 |
| 0.5 | 0.5 | no | 211 | 40 | 40 | 19 | 0 | 1 | 0 | 1 | 3 | 14 | 14 | 2 of 3 |
| 0.5 | 0.5 | yes | 221 | 40 | 40 | 24 | 0 | 1 | 11 | 1 | 3 | 15 | 14 | 2 of 3 |
| 0.5 | 0.3 | no | 220 | 40 | 40 | 28 | 0 | 10 | 0 | 9 | 7 | 14 | 18 | 3 of 3 |
| 0.5 | 0.3 | yes | 235 | 40 | 40 | 35 | 0 | 10 | 21 | 9 | 7 | 15 | 18 | 3 of 3 |
| 0.5 | 0.2 | no | 221 | 40 | 40 | 29 | 0 | 11 | 0 | 10 | 7 | 14 | 18 | 3 of 3 |
| 0.5 | 0.2 | yes | 245 | 40 | 40 | 44 | 0 | 16 | 28 | 15 | 9 | 15 | 19 | 3 of 3 |
| 0.5 | 0.1 | no | 221 | 40 | 40 | 29 | 0 | 11 | 0 | 10 | 7 | 14 | 18 | 3 of 3 |
| 0.5 | 0.1 | yes | 256 | 40 | 40 | 55 | 0 | 24 | 36 | 23 | 14 | 15 | 21 | 3 of 3 |
| 0.5 | 0.05 | no | 221 | 40 | 40 | 30 | 0 | 11 | 0 | 10 | 6 | 15 | 18 | 3 of 3 |
| 0.5 | 0.05 | yes | 267 | 40 | 40 | 67 | 0 | 33 | 47 | 32 | 23 | 16 | 22 | 3 of 3 |
| 1 | none | no | 210 | 40 | 40 | 22 | 0 | 0 | 0 | 0 | 2 | 14 | 13 | 2 of 3 |
| 1 | 0.5 | no | 211 | 40 | 40 | 23 | 0 | 1 | 0 | 1 | 3 | 14 | 14 | 2 of 3 |
| 1 | 0.5 | yes | 221 | 40 | 40 | 31 | 0 | 1 | 11 | 1 | 4 | 15 | 14 | 3 of 3 |
| 1 | 0.3 | no | 220 | 40 | 40 | 32 | 0 | 10 | 0 | 9 | 7 | 14 | 18 | 3 of 3 |
| 1 | 0.3 | yes | 235 | 40 | 40 | 43 | 0 | 10 | 21 | 9 | 8 | 16 | 18 | 3 of 3 |
| 1 | 0.2 | no | 221 | 40 | 40 | 33 | 0 | 11 | 0 | 10 | 7 | 14 | 18 | 3 of 3 |
| 1 | 0.2 | yes | 245 | 40 | 40 | 52 | 0 | 16 | 28 | 15 | 10 | 16 | 19 | 3 of 3 |
| 1 | 0.1 | no | 221 | 40 | 40 | 33 | 0 | 11 | 0 | 10 | 7 | 14 | 18 | 3 of 3 |
| 1 | 0.1 | yes | 256 | 40 | 40 | 63 | 0 | 24 | 36 | 23 | 15 | 17 | 21 | 3 of 3 |
| 1 | 0.05 | no | 221 | 40 | 40 | 34 | 0 | 11 | 0 | 10 | 7 | 15 | 18 | 3 of 3 |
| 1 | 0.05 | yes | 267 | 40 | 40 | 75 | 0 | 33 | 47 | 32 | 24 | 17 | 22 | 3 of 3 |

## Prompt 15: bracket 5 Najeela, the Blade-Blossom (warriors combat)

Floors: tutors 4, fast mana 6, Game Changers 8. Reach with a rate: tutors 536, fast mana 73, Game Changers 53.

Rates of the power cards with a rate: highest 1.000, upper quarter 0.005, median 0.000, lower quarter 0.000, over 646 cards.

| Card | Rate | Role | Theme | Counts toward | Before | 0.1, keep 0.3, pin | 0.3, keep 0.3, pin | 0.3, keep 0.3 | 0.5, keep 0.3 | 1, keep 0.3 |
|---|---|---|---|---|---|---|---|---|---|---|
| Sol Ring | 1.000 | ramp | no | fast mana | no | yes | yes | yes | yes | yes |
| Lotus Petal | 0.913 | ramp | no | fast mana | no | yes | yes | yes | yes | yes |
| Chrome Mox | 0.897 | ramp | no | fast mana, Game Changers | no | yes | yes | yes | yes | yes |
| Mana Vault | 0.857 | ramp | no | fast mana, Game Changers | no | yes | yes | yes | yes | yes |
| Mox Diamond | 0.841 | ramp | no | fast mana, Game Changers | no | yes | yes | yes | yes | yes |
| Mox Amber | 0.656 | ramp | no | fast mana | no | yes | yes | yes | yes | yes |
| Force of Will | 0.648 | interaction | no | Game Changers | no | yes | yes | yes | yes | yes |
| Rhystic Study | 0.648 | draw | no | Game Changers | no | yes | yes | no | yes | yes |
| Fierce Guardianship | 0.630 | interaction | no | Game Changers | no | yes | yes | yes | yes | yes |
| Mox Opal | 0.577 | ramp | no | fast mana | no | yes | yes | no | yes | yes |
| Vampiric Tutor | 0.510 | other | no | tutors, Game Changers | no | yes | yes | yes | yes | yes |
| Demonic Tutor | 0.496 | other | no | tutors, Game Changers | no | yes | yes | yes | yes | yes |
| Lion's Eye Diamond | 0.489 | ramp | no | fast mana, Game Changers | no | yes | yes | no | yes | yes |
| An Offer You Can't Refuse | 0.467 | interaction | no | fast mana | no | yes | yes | no | yes | yes |
| Gamble | 0.446 | other | no | tutors, Game Changers | no | yes | yes | yes | yes | yes |

| Weight | Keep | Pin | Cards | Lands | Fixing | In | Theme out | Kept | Pinned | Other | Tutors | Fast mana | Game Changers | Floors met |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| 0.1 | none | no | 300 | 36 | 18 | 0 | 0 | 0 | 0 | 0 | 4 | 0 | 0 | 1 of 3 |
| 0.1 | 0.5 | no | 300 | 34 | 16 | 2 | 0 | 2 | 0 | 2 | 5 | 0 | 1 | 1 of 3 |
| 0.1 | 0.5 | yes | 311 | 35 | 17 | 12 | 0 | 2 | 11 | 2 | 5 | 7 | 7 | 2 of 3 |
| 0.1 | 0.3 | no | 300 | 25 | 7 | 11 | 0 | 11 | 0 | 10 | 10 | 0 | 8 | 2 of 3 |
| 0.1 | 0.3 | yes | 332 | 29 | 11 | 39 | 0 | 21 | 32 | 20 | 15 | 12 | 19 | 3 of 3 |
| 0.1 | 0.2 | no | 300 | 25 | 7 | 11 | 0 | 11 | 0 | 10 | 10 | 0 | 8 | 2 of 3 |
| 0.1 | 0.2 | yes | 342 | 26 | 8 | 52 | 0 | 30 | 42 | 29 | 20 | 13 | 24 | 3 of 3 |
| 0.1 | 0.1 | no | 300 | 25 | 7 | 11 | 0 | 11 | 0 | 10 | 10 | 0 | 8 | 2 of 3 |
| 0.1 | 0.1 | yes | 356 | 26 | 8 | 66 | 0 | 37 | 56 | 36 | 28 | 17 | 28 | 3 of 3 |
| 0.1 | 0.05 | no | 300 | 25 | 7 | 11 | 0 | 11 | 0 | 10 | 10 | 0 | 8 | 2 of 3 |
| 0.1 | 0.05 | yes | 375 | 26 | 8 | 85 | 0 | 52 | 75 | 51 | 42 | 20 | 32 | 3 of 3 |
| 0.2 | none | no | 300 | 36 | 18 | 0 | 0 | 0 | 0 | 0 | 4 | 0 | 0 | 1 of 3 |
| 0.2 | 0.5 | no | 300 | 34 | 16 | 7 | 5 | 2 | 0 | 2 | 5 | 4 | 4 | 1 of 3 |
| 0.2 | 0.5 | yes | 311 | 35 | 17 | 13 | 1 | 2 | 11 | 2 | 5 | 7 | 7 | 2 of 3 |
| 0.2 | 0.3 | no | 300 | 25 | 7 | 16 | 5 | 11 | 0 | 10 | 10 | 4 | 11 | 2 of 3 |
| 0.2 | 0.3 | yes | 332 | 29 | 11 | 40 | 1 | 21 | 32 | 20 | 15 | 12 | 19 | 3 of 3 |
| 0.2 | 0.2 | no | 300 | 25 | 7 | 16 | 5 | 11 | 0 | 10 | 10 | 4 | 11 | 2 of 3 |
| 0.2 | 0.2 | yes | 342 | 26 | 8 | 53 | 1 | 30 | 42 | 29 | 20 | 13 | 24 | 3 of 3 |
| 0.2 | 0.1 | no | 300 | 25 | 7 | 16 | 5 | 11 | 0 | 10 | 10 | 4 | 11 | 2 of 3 |
| 0.2 | 0.1 | yes | 356 | 26 | 8 | 67 | 1 | 37 | 56 | 36 | 28 | 17 | 28 | 3 of 3 |
| 0.2 | 0.05 | no | 300 | 25 | 7 | 16 | 5 | 11 | 0 | 10 | 10 | 4 | 11 | 2 of 3 |
| 0.2 | 0.05 | yes | 375 | 26 | 8 | 86 | 1 | 52 | 75 | 51 | 42 | 20 | 32 | 3 of 3 |
| 0.3 | none | no | 300 | 36 | 18 | 2 | 2 | 0 | 0 | 0 | 4 | 0 | 0 | 1 of 3 |
| 0.3 | 0.5 | no | 300 | 34 | 16 | 17 | 15 | 2 | 0 | 2 | 5 | 6 | 6 | 2 of 3 |
| 0.3 | 0.5 | yes | 311 | 35 | 17 | 21 | 9 | 2 | 11 | 2 | 5 | 7 | 7 | 2 of 3 |
| 0.3 | 0.3 | no | 300 | 25 | 7 | 26 | 15 | 11 | 0 | 10 | 10 | 6 | 13 | 3 of 3 |
| 0.3 | 0.3 | yes | 332 | 29 | 11 | 48 | 9 | 21 | 32 | 20 | 15 | 12 | 19 | 3 of 3 |
| 0.3 | 0.2 | no | 300 | 25 | 7 | 26 | 15 | 11 | 0 | 10 | 10 | 6 | 13 | 3 of 3 |
| 0.3 | 0.2 | yes | 342 | 26 | 8 | 61 | 9 | 30 | 42 | 29 | 20 | 13 | 24 | 3 of 3 |
| 0.3 | 0.1 | no | 300 | 25 | 7 | 26 | 15 | 11 | 0 | 10 | 10 | 6 | 13 | 3 of 3 |
| 0.3 | 0.1 | yes | 356 | 26 | 8 | 75 | 9 | 37 | 56 | 36 | 28 | 17 | 28 | 3 of 3 |
| 0.3 | 0.05 | no | 300 | 25 | 7 | 26 | 15 | 11 | 0 | 10 | 10 | 6 | 13 | 3 of 3 |
| 0.3 | 0.05 | yes | 375 | 26 | 8 | 94 | 9 | 52 | 75 | 51 | 42 | 20 | 32 | 3 of 3 |
| 0.5 | none | no | 300 | 36 | 18 | 2 | 2 | 0 | 0 | 0 | 4 | 0 | 0 | 1 of 3 |
| 0.5 | 0.5 | no | 300 | 34 | 16 | 21 | 19 | 2 | 0 | 2 | 5 | 7 | 7 | 2 of 3 |
| 0.5 | 0.5 | yes | 311 | 35 | 17 | 21 | 9 | 2 | 11 | 2 | 5 | 7 | 7 | 2 of 3 |
| 0.5 | 0.3 | no | 300 | 25 | 7 | 42 | 31 | 11 | 0 | 10 | 10 | 11 | 16 | 3 of 3 |
| 0.5 | 0.3 | yes | 332 | 29 | 11 | 56 | 17 | 21 | 32 | 20 | 15 | 12 | 19 | 3 of 3 |
| 0.5 | 0.2 | no | 300 | 25 | 7 | 42 | 31 | 11 | 0 | 10 | 10 | 11 | 16 | 3 of 3 |
| 0.5 | 0.2 | yes | 342 | 26 | 8 | 69 | 17 | 30 | 42 | 29 | 20 | 13 | 24 | 3 of 3 |
| 0.5 | 0.1 | no | 300 | 25 | 7 | 42 | 31 | 11 | 0 | 10 | 10 | 11 | 16 | 3 of 3 |
| 0.5 | 0.1 | yes | 356 | 26 | 8 | 83 | 17 | 37 | 56 | 36 | 28 | 17 | 28 | 3 of 3 |
| 0.5 | 0.05 | no | 300 | 25 | 7 | 42 | 31 | 11 | 0 | 10 | 10 | 11 | 16 | 3 of 3 |
| 0.5 | 0.05 | yes | 375 | 26 | 8 | 102 | 17 | 52 | 75 | 51 | 42 | 20 | 32 | 3 of 3 |
| 1 | none | no | 300 | 36 | 18 | 16 | 16 | 0 | 0 | 0 | 6 | 5 | 5 | 1 of 3 |
| 1 | 0.5 | no | 300 | 34 | 16 | 24 | 22 | 2 | 0 | 2 | 7 | 7 | 7 | 2 of 3 |
| 1 | 0.5 | yes | 311 | 35 | 17 | 25 | 13 | 2 | 11 | 2 | 7 | 7 | 7 | 2 of 3 |
| 1 | 0.3 | no | 300 | 25 | 7 | 57 | 46 | 11 | 0 | 10 | 12 | 12 | 18 | 3 of 3 |
| 1 | 0.3 | yes | 332 | 29 | 11 | 67 | 28 | 21 | 32 | 20 | 17 | 12 | 19 | 3 of 3 |
| 1 | 0.2 | no | 300 | 25 | 7 | 61 | 50 | 11 | 0 | 10 | 12 | 13 | 18 | 3 of 3 |
| 1 | 0.2 | yes | 342 | 26 | 8 | 88 | 36 | 30 | 42 | 29 | 22 | 13 | 24 | 3 of 3 |
| 1 | 0.1 | no | 300 | 25 | 7 | 64 | 53 | 11 | 0 | 10 | 12 | 13 | 19 | 3 of 3 |
| 1 | 0.1 | yes | 356 | 26 | 8 | 103 | 37 | 37 | 56 | 36 | 28 | 17 | 28 | 3 of 3 |
| 1 | 0.05 | no | 300 | 25 | 7 | 64 | 53 | 11 | 0 | 10 | 12 | 13 | 19 | 3 of 3 |
| 1 | 0.05 | yes | 375 | 26 | 8 | 122 | 37 | 52 | 75 | 51 | 42 | 20 | 32 | 3 of 3 |

## Summary

| Weight | Keep | Pin | Prompts at every floor | Floors met | Cards in | Theme out | Fixing lands out |
|---|---|---|---|---|---|---|---|
| 0.1 | none | no | 1 of 6 | 9 of 18 | 0 | 0 | 0 |
| 0.1 | 0.5 | no | 1 of 6 | 11 of 18 | 7 | 0 | 2 |
| 0.1 | 0.5 | yes | 4 of 6 | 16 of 18 | 61 | 0 | 1 |
| 0.1 | 0.3 | no | 2 of 6 | 14 of 18 | 45 | 0 | 17 |
| 0.1 | 0.3 | yes | 6 of 6 | 18 of 18 | 149 | 0 | 7 |
| 0.1 | 0.2 | no | 2 of 6 | 14 of 18 | 51 | 0 | 19 |
| 0.1 | 0.2 | yes | 6 of 6 | 18 of 18 | 204 | 0 | 13 |
| 0.1 | 0.1 | no | 2 of 6 | 14 of 18 | 69 | 1 | 27 |
| 0.1 | 0.1 | yes | 6 of 6 | 18 of 18 | 283 | 0 | 25 |
| 0.1 | 0.05 | no | 2 of 6 | 14 of 18 | 82 | 1 | 27 |
| 0.1 | 0.05 | yes | 6 of 6 | 18 of 18 | 350 | 0 | 25 |
| 0.2 | none | no | 1 of 6 | 9 of 18 | 17 | 6 | 0 |
| 0.2 | 0.5 | no | 1 of 6 | 11 of 18 | 29 | 11 | 2 |
| 0.2 | 0.5 | yes | 4 of 6 | 16 of 18 | 74 | 5 | 1 |
| 0.2 | 0.3 | no | 2 of 6 | 14 of 18 | 68 | 12 | 17 |
| 0.2 | 0.3 | yes | 6 of 6 | 18 of 18 | 157 | 2 | 7 |
| 0.2 | 0.2 | no | 2 of 6 | 14 of 18 | 73 | 12 | 19 |
| 0.2 | 0.2 | yes | 6 of 6 | 18 of 18 | 210 | 2 | 13 |
| 0.2 | 0.1 | no | 2 of 6 | 14 of 18 | 89 | 13 | 27 |
| 0.2 | 0.1 | yes | 6 of 6 | 18 of 18 | 289 | 2 | 25 |
| 0.2 | 0.05 | no | 2 of 6 | 14 of 18 | 98 | 13 | 27 |
| 0.2 | 0.05 | yes | 6 of 6 | 18 of 18 | 356 | 2 | 25 |
| 0.3 | none | no | 1 of 6 | 9 of 18 | 36 | 16 | 1 |
| 0.3 | 0.5 | no | 1 of 6 | 12 of 18 | 56 | 29 | 3 |
| 0.3 | 0.5 | yes | 4 of 6 | 16 of 18 | 95 | 19 | 2 |
| 0.3 | 0.3 | no | 3 of 6 | 15 of 18 | 96 | 31 | 18 |
| 0.3 | 0.3 | yes | 6 of 6 | 18 of 18 | 182 | 19 | 8 |
| 0.3 | 0.2 | no | 3 of 6 | 15 of 18 | 100 | 31 | 19 |
| 0.3 | 0.2 | yes | 6 of 6 | 18 of 18 | 231 | 15 | 13 |
| 0.3 | 0.1 | no | 3 of 6 | 15 of 18 | 113 | 31 | 27 |
| 0.3 | 0.1 | yes | 6 of 6 | 18 of 18 | 307 | 12 | 25 |
| 0.3 | 0.05 | no | 3 of 6 | 15 of 18 | 120 | 31 | 27 |
| 0.3 | 0.05 | yes | 6 of 6 | 18 of 18 | 371 | 12 | 25 |
| 0.5 | none | no | 1 of 6 | 9 of 18 | 55 | 26 | 1 |
| 0.5 | 0.5 | no | 2 of 6 | 13 of 18 | 94 | 58 | 3 |
| 0.5 | 0.5 | yes | 4 of 6 | 16 of 18 | 115 | 30 | 2 |
| 0.5 | 0.3 | no | 4 of 6 | 16 of 18 | 147 | 73 | 18 |
| 0.5 | 0.3 | yes | 6 of 6 | 18 of 18 | 208 | 39 | 8 |
| 0.5 | 0.2 | no | 4 of 6 | 16 of 18 | 153 | 74 | 20 |
| 0.5 | 0.2 | yes | 6 of 6 | 18 of 18 | 258 | 35 | 14 |
| 0.5 | 0.1 | no | 4 of 6 | 16 of 18 | 166 | 75 | 27 |
| 0.5 | 0.1 | yes | 6 of 6 | 18 of 18 | 334 | 33 | 25 |
| 0.5 | 0.05 | no | 4 of 6 | 16 of 18 | 169 | 75 | 27 |
| 0.5 | 0.05 | yes | 6 of 6 | 18 of 18 | 395 | 29 | 25 |
| 1 | none | no | 1 of 6 | 9 of 18 | 97 | 57 | 1 |
| 1 | 0.5 | no | 4 of 6 | 16 of 18 | 147 | 100 | 3 |
| 1 | 0.5 | yes | 5 of 6 | 17 of 18 | 159 | 60 | 2 |
| 1 | 0.3 | no | 6 of 6 | 18 of 18 | 223 | 138 | 18 |
| 1 | 0.3 | yes | 6 of 6 | 18 of 18 | 268 | 84 | 8 |
| 1 | 0.2 | no | 6 of 6 | 18 of 18 | 235 | 145 | 20 |
| 1 | 0.2 | yes | 6 of 6 | 18 of 18 | 326 | 90 | 14 |
| 1 | 0.1 | no | 6 of 6 | 18 of 18 | 252 | 149 | 28 |
| 1 | 0.1 | yes | 6 of 6 | 18 of 18 | 401 | 86 | 26 |
| 1 | 0.05 | no | 6 of 6 | 18 of 18 | 254 | 150 | 28 |
| 1 | 0.05 | yes | 6 of 6 | 18 of 18 | 461 | 82 | 26 |
