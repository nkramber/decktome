# PR-14B quality gate

Run date: 2026-10-04. Card snapshot: 2026-09-04. Model: 20261004T051948Z.

## Separation on the holdout

The bars read 5 folds, every list holdout once, and the pairs of a bar sample evenly under the cap of 20000 (M-7). The precon bar reads each precon against its own broken copies (D-573), and the synergy copies read a bar of their own (D-957). The cross pairs, every precon against every copy, stand as information.

The synergy check drops a copy whose break lowered the `synergy` feature by less than 0.10 standard deviations of the real training lists (D-652). It reads commander, and every other format keeps each copy (D-653). Synthetic counts the copies the engine made, and Immaterial the ones the check dropped, each copy once over the folds.

| Format | Lists | Used | Synthetic | Immaterial | Holdout | Great over precon | Precon over own copy, synergy aside | Precon over own synergy copy | Precon over bad, cross | Accuracy |
|---|---|---|---|---|---|---|---|---|---|---|
| commander | 27463 | 9234 | 6170 | 219 | 15185 | 0.97 of 98351 (bar 0.90) | 1.00 of 569 (bar 0.95) | 0.78 of 255 (bar 0.75) | 0.86 of 99937 | 0.56 |
| standard | 13016 | 6722 | 980 | 0 | 7702 | 0.97 of 4007 (bar 0.90) | 1.00 of 23 (bar 0.95) | 0.94 of 17 (bar 0.75) | 0.82 of 1720 | 0.51 |
| modern | 27744 | 8435 | 3435 | 0 | 11870 | 0.98 of 98054 (bar 0.90) | 1.00 of 835 (bar 0.95) | 0.97 of 860 (bar 0.75) | 0.85 of 99853 | 0.57 |

## The three numbers (D-648)

The precon bar is a proxy, so every quality item reads two more numbers beside it. The model this run fitted grades the built decks of `pr8-deck-gate-run16.md`. The judge's tiers come from `pr14b-quality-judge-run5.md`. One judge run serves every model, so the agreement holds no judge noise. No bar reads the last two, and an item that moves none of the three is dropped. A change that moves a tier also reads the broken copies graded bad (D-674).

| Number | Read |
|---|---|
| Commander synergy axis, precon over own copy | 0.78 of 255 |
| Commander broken copies graded bad (D-674) | 4125 of 5951 |
| Built decks graded bad | 11 of 25 |
| Judge agreement | 9 of 25 |

| # | Deck | Format | Grade | Judge | Agree |
|---|---|---|---|---|---|
| 1 | lifegain Commander, any card | Commander | bad | typical | no |
| 2 | aristocrats Commander, owned first | Commander | typical | good | no |
| 3 | artifacts Commander, bracket 4 | Commander | baseline | typical | no |
| 4 | dinosaur tribal, bracket 2 | Commander | bad | typical | no |
| 5 | blink Commander, owned first | Commander | typical | typical | yes |
| 6 | Modern tempo, tournament | Modern | typical | typical | yes |
| 7 | Modern burn, casual | Modern | baseline | bad | no |
| 8 | Modern lifegain, FNM | Modern | bad | bad | yes |
| 9 | Standard midrange, FNM | Standard | bad | bad | yes |
| 10 | Standard aggro, tournament | Standard | baseline | baseline | yes |
| 11 | Commander with a locked card | Commander | baseline | typical | no |
| 12 | Commander on a budget | Commander | bad | baseline | no |
| 13 | owned first, and the commander is not owned | Commander | baseline | typical | no |
| 14 | the user delegates the commander | Commander | bad | typical | no |
| 15 | delegated commander, owned first | Commander | baseline | typical | no |
| 16 | a tight budget, owned first | Commander | bad | typical | no |
| 17 | upgrade a precon, owned first | Commander | typical | typical | yes |
| 18 | upgrade a precon, any card | Commander | baseline | baseline | yes |
| 19 | the Hobbit family, two colours | Commander | bad | baseline | no |
| 20 | the Hobbit family, a delegated commander | Commander | bad | baseline | no |
| 21 | the Hobbit family, mana from outside | Commander | bad | typical | no |
| 22 | a set family and a card from outside it | Commander | bad | typical | no |
| 23 | two set families at once | Commander | typical | typical | yes |
| 24 | a 60-card deck from one set | Modern | baseline | bad | no |
| 25 | use no card of an owned precon | Commander | typical | typical | yes |

## The bracket 5 offer

Commander reads of 2026-09-23: 1545 rows, 923 with Topdeck.gg entries.

| Commander | cEDH signal | Top cuts | Entries |
|---|---|---|---|
| Abaddon the Despoiler | 0.58 | 4 | 4 |
| The Queen of Dale | 0.53 | 4 | 5 |
| The Jolly Balloon Man | 0.52 | 5 | 7 |

## The commander model

Tiers, worst first: bad, baseline, typical, good, great. Train counts: bad 4847, baseline 151, good 3218, great 3174, typical 853. Holdout counts: bad 1112, baseline 44, good 782, great 826, typical 186.

Cards with a rate: 9453. Pairs that lift: 200000. Commanders with a signal: 1521. Top lists shape: 27.4 lands at 2.18 over 3174 lists. Typical lists shape: 35.6 lands at 3.10 over 853 lists.

Features with no spread, dropped: playset_share, singleton_share.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.154 | 0.137 | +1.099 |
| `unseen_share` | 0.078 | 0.109 | -0.215 |
| `synergy` | 0.464 | 0.447 | +0.329 |
| `land_need` | -5.866 | 4.109 | -0.142 |
| `avg_mana_value` | 2.788 | 0.856 | -0.464 |
| `color_sources` | 1.079 | 0.309 | +0.032 |
| `tapped_share` | 0.071 | 0.083 | +0.356 |
| `mana_turn_four` | 4.719 | 0.618 | +0.747 |
| `hands_two_to_four_lands` | 0.649 | 0.086 | +0.468 |
| `curve_low` | 0.509 | 0.187 | +0.224 |
| `curve_high` | 0.152 | 0.133 | -0.224 |
| `ramp` | 17.950 | 7.697 | +0.047 |
| `draw` | 7.373 | 4.153 | -0.222 |
| `removal` | 6.552 | 3.819 | +0.097 |
| `wipe` | 1.496 | 1.694 | -0.161 |
| `interaction` | 5.536 | 4.586 | +0.284 |
| `empty_roles` | 0.154 | 0.367 | -0.054 |
| `fast_mana` | 5.195 | 4.750 | +0.253 |
| `game_changer` | 7.280 | 7.249 | +0.181 |
| `source_spread` | 0.095 | 0.123 | -0.484 |
| `cedh_signal` | 0.280 | 0.200 | +0.199 |
| `high_bracket_share` | 0.225 | 0.254 | +0.041 |
| `commander_decks` | 6.003 | 4.111 | -0.151 |

Thresholds: -3.420, -1.664, 0.662, 3.149. Loss: 1.022 over 3000 iterations. Defect detector accuracy on the holdout: 0.91, at a cut of 0.52.

The precon bar per broken axis, each precon over its own copies, with the cross pairs after it:

- lands: 1.00 of 195, cross 0.97 of 48067
- curve: 1.00 of 195, cross 0.95 of 48067
- colors: 1.00 of 179, cross 0.91 of 28452
- synergy: 0.78 of 255, cross 0.76 of 98397

The synergy check dropped 219 of 2972 synergy copies over the folds, at a floor of 0.10 and a unit of 0.4870 (D-652). By the axis the fit asked for: colors 16, copies 101, synergy 102.

The misses per axis over the folds (M-7). Both passed: the detector passed the precon and the copy, and the ladder put the copy at or above the precon. Precon flagged: the detector flagged the precon and passed the copy. Both flagged: it flagged both and read the precon as the more broken. Own copy: each precon against its own copy on the axis, with its misses by the same kinds. Moved: the mean absolute standardized delta between a precon and its own copy, for the corpus features and the three features that move most.

| Axis | Share | Both passed | Precon flagged | Both flagged | Own copy | Own misses | Corpus moved | Most moved |
|---|---|---|---|---|---|---|---|---|
| lands | 0.97 of 48067 | 742 | 399 | 400 | 195 of 195 | 0, 0, 0 | `card_rate` 0.03, `synergy` 0.27, `unseen_share` 0.30 | `land_need` 2.91, `hands_two_to_four_lands` 2.07, `color_sources` 1.28 |
| curve | 0.95 of 48067 | 551 | 612 | 1192 | 195 of 195 | 0, 0, 0 | `card_rate` 0.18, `synergy` 0.37, `unseen_share` 0.53 | `curve_high` 1.87, `avg_mana_value` 1.49, `land_need` 1.34 |
| colors | 0.91 of 28452 | 1225 | 780 | 697 | 179 of 179 | 0, 0, 0 | `card_rate` 0.00, `synergy` 0.00, `unseen_share` 0.00 | `source_spread` 3.71, `tapped_share` 2.78, `color_sources` 1.99 |
| synergy | 0.76 of 98397 | 9751 | 7780 | 5961 | 198 of 255 | 46, 8, 3 | `card_rate` 0.09, `synergy` 0.97, `unseen_share` 0.90 | `synergy` 0.97, `unseen_share` 0.90, `wipe` 0.71 |

The precons that lose most over the sampled pairs of the bar:

| Precon | Grade | Detector | Misses |
|---|---|---|---|
| Eternal Bargain (2013-11-01) | bad | 0.89 | 305 of 454 |
| Peer Through Time (2014-11-07) | bad | 0.88 | 303 of 500 |
| Political Puppets (2011-06-17) | bad | 0.85 | 287 of 500 |
| Heavenly Inferno (2011-06-17) | bad | 0.87 | 316 of 571 |
| The Fantastic Four Collector's Edition (2026-06-26) | bad | 0.80 | 279 of 526 |
| Evasive Maneuvers (2013-11-01) | bad | 0.82 | 248 of 500 |
| Heavenly Inferno (2017-06-09) | bad | 0.82 | 217 of 454 |
| Evasive Maneuvers (2017-06-09) | bad | 0.75 | 244 of 526 |
| Plunder the Graves (2017-06-09) | bad | 0.79 | 258 of 571 |
| Elven Empire (2021-02-05) | bad | 0.79 | 225 of 526 |

The audit of the bar (M-8). The bar asks the ladder to score a precon above its own broken copy. Rank is the precon's place among the holdout precons by score, 0 the weakest, and the column reads the mean rank of the pairs that win and of the pairs that lose. A break that hurts every precon alike moves the two means together. A lower mean rank among the losers says the bar fails on the weak precons, which is where a broken copy is most likely the better deck. Card rate and synergy read the signed move, the copy less the precon, over the pairs that lose. A positive card rate says the copy holds cards the top lists play more than the precon does.

| Axis | Own pairs | Lost | Mean rank, won | Mean rank, lost | `card_rate` move | `synergy` move | Copy graded above |
|---|---|---|---|---|---|---|---|
| lands | 195 | 0 | 19.11 | - | - | - | 0 of 0 |
| curve | 195 | 0 | 19.11 | - | - | - | 0 of 0 |
| colors | 179 | 0 | 19.53 | - | - | - | 0 of 0 |
| synergy | 255 | 57 | 24.98 | 16.74 | 0.00 | -0.23 | 26 of 57 |

Confusion on the holdout over the folds, rows are the label and columns the grade, worst first:

- bad: [4125 1408 413 5 0]
- baseline: [69 78 42 6 0]
- typical: [63 278 638 51 9]
- good: [265 67 252 1472 1944]
- great: [250 82 193 1269 2206]

The cards the top lists hold most: Sol Ring, Lotus Petal, Chrome Mox, Mana Vault, Mox Diamond, Mystic Remora, Mental Misstep, Mox Amber, Arcane Signet, Flusterstorm, Force of Will, Pact of Negation.

## The standard model

Tiers, worst first: bad, baseline, typical, good, great. Train counts: bad 845, baseline 7, good 3188, great 2013, typical 162. Holdout counts: bad 135, baseline 1, good 812, great 513, typical 26.

Cards with a rate: 908. Pairs that lift: 7373. Commanders with a signal: 0. Top lists shape: 23.3 lands at 2.49 over 2013 lists. Typical lists shape: 23.1 lands at 2.53 over 162 lists.

Features with no spread, dropped: game_changer, cedh_signal, high_bracket_share, commander_decks.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.094 | 0.053 | +1.363 |
| `unseen_share` | 0.031 | 0.102 | +0.107 |
| `synergy` | 1.599 | 0.831 | +1.048 |
| `land_need` | 0.360 | 2.283 | +0.214 |
| `avg_mana_value` | 2.597 | 0.713 | -0.752 |
| `color_sources` | 1.010 | 0.332 | +0.486 |
| `tapped_share` | 0.036 | 0.076 | -0.129 |
| `mana_turn_four` | 4.071 | 0.784 | +0.202 |
| `hands_two_to_four_lands` | 0.755 | 0.037 | +0.540 |
| `curve_low` | 0.624 | 0.182 | -0.306 |
| `curve_high` | 0.134 | 0.140 | +0.032 |
| `ramp` | 5.323 | 6.479 | -0.260 |
| `draw` | 5.094 | 3.667 | +0.022 |
| `removal` | 8.822 | 5.059 | +0.066 |
| `wipe` | 0.826 | 1.456 | +0.142 |
| `interaction` | 2.342 | 2.780 | +0.275 |
| `empty_roles` | 1.024 | 0.671 | +0.256 |
| `fast_mana` | 0.108 | 0.679 | +0.194 |
| `playset_share` | 0.595 | 0.196 | +0.081 |
| `singleton_share` | 0.098 | 0.119 | -0.482 |
| `source_spread` | 0.265 | 0.225 | +0.268 |

Thresholds: -4.219, -2.287, -0.297, 1.404. Loss: 1.075 over 3000 iterations. Defect detector accuracy on the holdout: 0.96, at a cut of 0.58.

The precon bar per broken axis, each precon over its own copies, with the cross pairs after it:

- lands: 1.00 of 8, cross 0.92 of 344
- curve: 1.00 of 8, cross 0.98 of 344
- colors: 1.00 of 4, cross 0.59 of 273
- copies: 1.00 of 3, cross 0.95 of 283
- synergy: 0.94 of 17, cross 0.70 of 476

The misses per axis over the folds (M-7). Both passed: the detector passed the precon and the copy, and the ladder put the copy at or above the precon. Precon flagged: the detector flagged the precon and passed the copy. Both flagged: it flagged both and read the precon as the more broken. Own copy: each precon against its own copy on the axis, with its misses by the same kinds. Moved: the mean absolute standardized delta between a precon and its own copy, for the corpus features and the three features that move most.

| Axis | Share | Both passed | Precon flagged | Both flagged | Own copy | Own misses | Corpus moved | Most moved |
|---|---|---|---|---|---|---|---|---|
| lands | 0.92 of 344 | 2 | 3 | 22 | 8 of 8 | 0, 0, 0 | `card_rate` 0.07, `synergy` 0.12, `unseen_share` 0.89 | `hands_two_to_four_lands` 3.87, `land_need` 3.38, `color_sources` 1.92 |
| curve | 0.98 of 344 | 0 | 3 | 4 | 8 of 8 | 0, 0, 0 | `card_rate` 0.25, `synergy` 0.21, `unseen_share` 1.82 | `avg_mana_value` 3.32, `curve_high` 3.01, `land_need` 2.22 |
| colors | 0.59 of 273 | 46 | 41 | 25 | 4 of 4 | 0, 0, 0 | `card_rate` 0.00, `synergy` 0.00, `unseen_share` 0.00 | `tapped_share` 3.02, `color_sources` 1.63, `source_spread` 1.49 |
| copies | 0.95 of 283 | 2 | 4 | 9 | 3 of 3 | 0, 0, 0 | `card_rate` 0.28, `synergy` 0.02, `unseen_share` 0.60 | `singleton_share` 3.84, `playset_share` 2.30, `empty_roles` 0.99 |
| synergy | 0.70 of 476 | 40 | 66 | 39 | 16 of 17 | 1, 0, 0 | `card_rate` 0.11, `synergy` 0.18, `unseen_share` 2.50 | `unseen_share` 2.50, `color_sources` 1.98, `playset_share` 1.59 |

The precons that lose most over the sampled pairs of the bar:

| Precon | Grade | Detector | Misses |
|---|---|---|---|
| Leonardo (2026-03-06) | bad | 0.73 | 75 of 210 |
| Donatello (2026-03-06) | bad | 0.70 | 71 of 210 |
| Michealangelo (2026-03-06) | bad | 0.53 | 42 of 210 |
| Raphael (2026-03-06) | bad | 0.49 | 28 of 160 |
| Angels (2026-01-23) | baseline | 0.19 | 41 of 265 |
| Eerie (2026-04-24) | baseline | 0.03 | 15 of 135 |
| Pirates (2026-01-23) | typical | 0.09 | 29 of 265 |
| Lifegain (2026-04-24) | typical | 0.01 | 5 of 265 |

| Axis | Own pairs | Lost | Mean rank, won | Mean rank, lost | `card_rate` move | `synergy` move | Copy graded above |
|---|---|---|---|---|---|---|---|
| lands | 8 | 0 | 0.75 | - | - | - | 0 of 0 |
| curve | 8 | 0 | 0.75 | - | - | - | 0 of 0 |
| colors | 4 | 0 | 0.75 | - | - | - | 0 of 0 |
| copies | 3 | 0 | 0.33 | - | - | - | 0 of 0 |
| synergy | 17 | 1 | 0.81 | 1.00 | 0.11 | 0.01 | 0 of 1 |

Confusion on the holdout over the folds, rows are the label and columns the grade, worst first:

- bad: [900 8 49 21 2]
- baseline: [4 2 2 0 0]
- typical: [56 9 65 36 22]
- good: [156 8 415 1729 1692]
- great: [123 3 287 881 1232]

The cards the top lists hold most: Burst Lightning, Stock Up, Llanowar Elves, Opt, Sleight of Hand, Spell Pierce, Eddymurk Crab, Badgermole Cub, Boomerang Basics, Stormchaser's Talent, Requiting Hex, Meltstrider's Resolve.

## The modern model

Tiers, worst first: bad, baseline, typical, good, great. Train counts: bad 2750, baseline 269, good 3184, great 2968, typical 281. Holdout counts: bad 685, baseline 62, good 816, great 780, typical 75.

Cards with a rate: 949. Pairs that lift: 7120. Commanders with a signal: 0. Top lists shape: 21.7 lands at 2.47 over 2968 lists. Typical lists shape: 21.9 lands at 2.52 over 281 lists.

Features with no spread, dropped: cedh_signal, high_bracket_share, commander_decks.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.066 | 0.044 | +0.572 |
| `unseen_share` | 0.208 | 0.355 | -0.516 |
| `synergy` | 1.413 | 0.993 | +1.503 |
| `land_need` | -1.365 | 3.337 | -0.055 |
| `avg_mana_value` | 2.752 | 0.956 | -0.341 |
| `color_sources` | 0.884 | 0.364 | +0.236 |
| `tapped_share` | 0.124 | 0.142 | +0.081 |
| `mana_turn_four` | 3.858 | 0.809 | +0.024 |
| `hands_two_to_four_lands` | 0.731 | 0.061 | +0.382 |
| `curve_low` | 0.585 | 0.218 | -0.165 |
| `curve_high` | 0.180 | 0.158 | -0.329 |
| `ramp` | 5.593 | 6.792 | +0.076 |
| `draw` | 4.503 | 4.356 | +0.080 |
| `removal` | 8.245 | 4.929 | -0.212 |
| `wipe` | 0.719 | 1.560 | -0.021 |
| `interaction` | 2.679 | 3.619 | +0.005 |
| `empty_roles` | 1.046 | 0.850 | -0.023 |
| `fast_mana` | 0.616 | 1.538 | +0.023 |
| `game_changer` | 0.265 | 0.885 | -0.152 |
| `playset_share` | 0.577 | 0.232 | -0.590 |
| `singleton_share` | 0.144 | 0.162 | -0.770 |
| `source_spread` | 0.255 | 0.225 | -0.290 |

Thresholds: -3.647, -1.439, 0.815, 2.363. Loss: 1.067 over 3000 iterations. Defect detector accuracy on the holdout: 0.93, at a cut of 0.44.

The precon bar per broken axis, each precon over its own copies, with the cross pairs after it:

- lands: 1.00 of 339, cross 0.93 of 46003
- curve: 1.00 of 339, cross 0.98 of 46003
- colors: 0.96 of 84, cross 0.47 of 24232
- copies: 1.00 of 73, cross 0.89 of 27705
- synergy: 0.97 of 860, cross 0.84 of 79352

The misses per axis over the folds (M-7). Both passed: the detector passed the precon and the copy, and the ladder put the copy at or above the precon. Precon flagged: the detector flagged the precon and passed the copy. Both flagged: it flagged both and read the precon as the more broken. Own copy: each precon against its own copy on the axis, with its misses by the same kinds. Moved: the mean absolute standardized delta between a precon and its own copy, for the corpus features and the three features that move most.

| Axis | Share | Both passed | Precon flagged | Both flagged | Own copy | Own misses | Corpus moved | Most moved |
|---|---|---|---|---|---|---|---|---|
| lands | 0.93 of 46003 | 1481 | 529 | 1070 | 339 of 339 | 0, 0, 0 | `card_rate` 0.01, `synergy` 0.00, `unseen_share` 0.09 | `hands_two_to_four_lands` 2.40, `land_need` 2.34, `color_sources` 1.33 |
| curve | 0.98 of 46003 | 131 | 102 | 758 | 339 of 339 | 0, 0, 0 | `card_rate` 0.04, `synergy` 0.00, `unseen_share` 0.21 | `curve_high` 2.69, `avg_mana_value` 2.21, `curve_low` 1.71 |
| colors | 0.47 of 24232 | 8551 | 3366 | 868 | 81 of 84 | 2, 1, 0 | `card_rate` 0.00, `synergy` 0.00, `unseen_share` 0.00 | `tapped_share` 1.33, `color_sources` 1.04, `source_spread` 1.00 |
| copies | 0.89 of 27705 | 1708 | 785 | 609 | 73 of 73 | 0, 0, 0 | `card_rate` 0.04, `synergy` 0.00, `unseen_share` 0.19 | `singleton_share` 3.55, `playset_share` 2.48, `empty_roles` 0.85 |
| synergy | 0.84 of 79352 | 5768 | 3485 | 3340 | 837 of 860 | 17, 0, 6 | `card_rate` 0.03, `synergy` 0.00, `unseen_share` 0.22 | `playset_share` 1.32, `empty_roles` 0.62, `removal` 0.57 |

The precons that lose most over the sampled pairs of the bar:

| Precon | Grade | Detector | Misses |
|---|---|---|---|
| Sultai Schemers (2014-09-26) | bad | 0.89 | 236 of 344 |
| Jeskai Monks (2014-09-26) | bad | 0.85 | 204 of 327 |
| Wizards (2005-08-01) | bad | 0.83 | 167 of 303 |
| Beasts (2005-08-01) | bad | 0.83 | 124 of 238 |
| Illusionary Might (2011-08-12) | bad | 0.82 | 120 of 238 |
| Thallids (2005-08-01) | bad | 0.81 | 114 of 238 |
| Spiraling Doom (2012-02-24) | bad | 0.78 | 131 of 303 |
| Underworld Herald (2014-02-28) | bad | 0.74 | 127 of 303 |
| Stampede (2004-06-04) | bad | 0.76 | 130 of 322 |
| Grave Advantage (2015-01-23) | bad | 0.79 | 94 of 238 |

| Axis | Own pairs | Lost | Mean rank, won | Mean rank, lost | `card_rate` move | `synergy` move | Copy graded above |
|---|---|---|---|---|---|---|---|
| lands | 339 | 0 | 33.25 | - | - | - | 0 of 0 |
| curve | 339 | 0 | 33.25 | - | - | - | 0 of 0 |
| colors | 84 | 3 | 37.35 | 46.00 | 0.00 | 0.00 | 0 of 3 |
| copies | 73 | 0 | 36.71 | - | - | - | 0 of 0 |
| synergy | 860 | 23 | 32.58 | 30.87 | 0.06 | 0.00 | 1 of 23 |

Confusion on the holdout over the folds, rows are the label and columns the grade, worst first:

- bad: [3213 67 137 15 3]
- baseline: [215 110 6 0 0]
- typical: [60 6 107 116 67]
- good: [189 2 870 1524 1415]
- great: [143 0 727 1094 1784]

The cards the top lists hold most: Solitude, Quantum Riddler, Force of Negation, Thoughtseize, Galvanic Discharge, Prismatic Ending, Kozilek's Command, Mishra's Bauble, Ragavan, Nimble Pilferer, Ephemerate, Malevolent Rumble, Fatal Push.

## Run

- Suite `quality`, run `pr14b-quality-gate-run26`, on 2026-10-04, commit `6503306`.
- Roles: none, no provider call.
- Versions: card snapshot 2026-09-04, commanders_day `2026-09-23`, quality_model `20261004T051948Z`.
- Calls: 0. Cost: unpriced. Time: 128 seconds.

## Verdict

Verdict: PASS. Time: 128 seconds, no provider call.
