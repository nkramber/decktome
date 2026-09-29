# The read of deck gate run 35, 2026-09-29

This is a dated record of PR-104 (D-995). Run 35 is the first whole deck gate run of generate prompt version 16. The gate document is `docs/reference/pr8-deck-gate-run35.md`.

## The run

- Commit `3c85ca4`. Card snapshot 2026-09-04. quality_model `20260923T202806Z`.
- Roles: generate and repair on `gpt-5.6-terra` at effort medium, judge on `claude-opus-5` at effort medium.
- Cost $3.1278, 77 calls, 1993 seconds. The cap was $4.
- Verdict PASS. 25 of 25 decks passed every block check. 0 invented names reached the user. 0 summaries stated a false rule.
- `make eval-check` reads the decks suite as PASS against run 19.

## The counts under the gate

| Measure | Run 29 | Run 31 | Run 35 |
|---|---|---|---|
| Generate prompt version | 15 | 15 | 16 |
| Mean plan score, 0 to 1 | 0.70 | 0.74 | 0.76 |
| Plan answers `theme_fit=no` | 4 | 1 | 0 |
| Plan answers `theme_fit=partly` | 5 | 7 | 10 |
| Plan answers `useful_as_built` not yes | 9 | 7 | 5 |
| Plan answers `plan_coherent=no` | 3 | 2 | 3 |
| Plan answers `summary_honest=no` | 3 | 2 | 1 |
| Findings `profile_off_band` | 11 | 10 | 6 |
| Findings `mana_pass` | 6 | 5 | 13 |
| Decks that needed the repair turn | 0 | 0 | 2 |
| Nonbasic lands over 25 decks | 257 | 261 | 275 |
| Cost | $2.7384 | $2.8348 | $3.1278 |

The grades of the three runs read three different quality models. So this record compares no grade.

## The repair turn (D-916)

Two decks needed the repair turn. Deck 15 needed it for `deck_size`. Deck 20 needed it for `copy_limit` and `profile_off_band`. Both repaired decks hold 0 block findings. Runs 29 and 31 needed no repair turn, so no earlier rate exists for a comparison.

## The lands of each deck (F-33, D-799, D-951)

Each cell reads basic lands / nonbasic lands.

| Deck | Prompt | Run 29 | Run 31 | Run 35 |
|---|---|---|---|---|
| 1 | lifegain Commander, any card | 24/13 | 22/14 | 22/14 |
| 2 | aristocrats Commander, owned first | 29/7 | 24/12 | 25/11 |
| 3 | artifacts Commander, bracket 4 | 26/7 | 20/13 | 25/8 |
| 4 | dinosaur tribal, bracket 2 | 17/20 | 19/18 | 2/35 |
| 5 | blink Commander, owned first | 26/11 | 30/8 | 36/0 |
| 6 | Modern tempo, tournament | 10/14 | 14/10 | 6/18 |
| 7 | Modern burn, casual | 20/4 | 24/0 | 24/0 |
| 8 | Modern lifegain, FNM | 8/16 | 18/6 | 12/12 |
| 9 | Standard midrange, FNM | 16/8 | 12/12 | 16/8 |
| 10 | Standard aggro, tournament | 12/12 | 8/16 | 12/12 |
| 11 | Commander with a locked card | 14/22 | 23/13 | 24/12 |
| 12 | Commander on a budget | 30/7 | 30/7 | 30/7 |
| 13 | owned first, and the commander is not owned | 28/8 | 36/0 | 25/11 |
| 14 | the user delegates the commander | 22/14 | 15/21 | 17/21 |
| 15 | delegated commander, owned first | 26/11 | 29/8 | 32/7 |
| 16 | a tight budget, owned first | 23/14 | 33/5 | 27/10 |
| 17 | upgrade a precon, owned first | 27/10 | 26/11 | 27/11 |
| 18 | upgrade a precon, any card | 11/28 | 13/27 | 14/25 |
| 19 | the Hobbit family, two colours | 36/0 | 29/10 | 28/8 |
| 20 | the Hobbit family, a delegated commander | 31/5 | 32/4 | 35/1 |
| 21 | the Hobbit family, mana from outside | 18/18 | 15/22 | 10/27 |
| 22 | a set family and a card from outside it | 36/0 | 27/9 | 27/9 |
| 23 | two set families at once | 28/8 | 29/7 | 28/8 |
| 24 | a 60-card deck from one set | 24/0 | 24/0 | 24/0 |
| 25 | use no card of an owned precon | 36/0 | 28/8 | 35/0 |

- Decks 5 and 25 hold 0 nonbasic lands. Their commanders, Gilraen, Dúnedain Protector and Captain America, Super-Soldier, are white alone in the snapshot. So the fixing floor does not apply.
- Deck 7 and deck 24 hold 0 nonbasic lands in all three runs.
- Deck 4, the Gishath deck at bracket 2, moved from 20 nonbasic lands to 35. The lands include Plateau, Savannah, Taiga, four fetch lands, Mana Confluence, and Cavern of Souls. The plan judge read `theme_fit=partly`, because the mana base is expensive for a bracket 2 deck. Run 29 read `theme_fit=yes` with the same concern. This is one sample.
