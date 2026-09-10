# PR-14B quality gate

Run date: 2026-09-10. Card snapshot: 2026-09-04. Model: 20260910T231306Z.

## Separation on the holdout

The bars read 5 folds, every list holdout once, and the pairs of a bar sample evenly under the cap of 20000 (M-7). The precon bar reads each precon against its own broken copies (D-573). The cross pairs, every precon against every copy, stand as information.

The synergy check drops a copy whose break lowered the `synergy` feature by less than 0.10 standard deviations of the real training lists (D-652). It reads commander, and every other format keeps each copy (D-653). Synthetic counts the copies the engine made, and Immaterial the ones the check dropped, each copy once over the folds.

| Format | Lists | Used | Synthetic | Immaterial | Holdout | Great over precon | Precon over own copy | Precon over bad, cross | Accuracy |
|---|---|---|---|---|---|---|---|---|---|
| commander | 23628 | 9229 | 6145 | 240 | 15134 | 0.96 of 97450 (bar 0.90) | 0.95 of 788 (bar 0.95) | 0.87 of 99933 | 0.65 |
| standard | 8228 | 5768 | 370 | 0 | 6138 | 0.92 of 2750 (bar 0.90) | 1.00 of 40 (bar 0.95) | 0.92 of 730 | 0.45 |
| modern | 20292 | 7099 | 2745 | 0 | 9844 | 0.98 of 98691 (bar 0.90) | 0.99 of 1635 (bar 0.95) | 0.86 of 99754 | 0.55 |

## The three numbers (D-648)

The precon bar is a proxy, so every quality item reads two more numbers beside it. The model this run fitted grades the built decks of `pr8-deck-gate-run16.md`. The judge's tiers come from `pr14b-quality-judge-run5.md`. One judge run serves every model, so the agreement holds no judge noise. No bar reads the last two, and an item that moves none of the three is dropped.

| Number | Read |
|---|---|
| Commander synergy axis, precon over own copy | 0.84 of 232 |
| Built decks graded bad | 16 of 25 |
| Judge agreement | 9 of 25 |

| # | Deck | Format | Grade | Judge | Agree |
|---|---|---|---|---|---|
| 1 | lifegain Commander, any card | Commander | bad | typical | no |
| 2 | aristocrats Commander, owned first | Commander | typical | good | no |
| 3 | artifacts Commander, bracket 4 | Commander | bad | typical | no |
| 4 | dinosaur tribal, bracket 2 | Commander | bad | typical | no |
| 5 | blink Commander, owned first | Commander | typical | typical | yes |
| 6 | Modern tempo, tournament | Modern | typical | typical | yes |
| 7 | Modern burn, casual | Modern | bad | bad | yes |
| 8 | Modern lifegain, FNM | Modern | bad | bad | yes |
| 9 | Standard midrange, FNM | Standard | bad | bad | yes |
| 10 | Standard aggro, tournament | Standard | baseline | baseline | yes |
| 11 | Commander with a locked card | Commander | bad | typical | no |
| 12 | Commander on a budget | Commander | bad | baseline | no |
| 13 | owned first, and the commander is not owned | Commander | bad | typical | no |
| 14 | the user delegates the commander | Commander | bad | typical | no |
| 15 | delegated commander, owned first | Commander | bad | typical | no |
| 16 | a tight budget, owned first | Commander | bad | typical | no |
| 17 | upgrade a precon, owned first | Commander | typical | typical | yes |
| 18 | upgrade a precon, any card | Commander | baseline | baseline | yes |
| 19 | the Hobbit family, two colours | Commander | bad | baseline | no |
| 20 | the Hobbit family, a delegated commander | Commander | bad | baseline | no |
| 21 | the Hobbit family, mana from outside | Commander | bad | typical | no |
| 22 | a set family and a card from outside it | Commander | bad | typical | no |
| 23 | two set families at once | Commander | typical | typical | yes |
| 24 | a 60-card deck from one set | Modern | baseline | bad | no |
| 25 | use no card of an owned precon | Commander | baseline | typical | no |

## The bracket 5 offer

Commander reads of 2026-09-07: 1562 rows, 945 with Topdeck.gg entries.

| Commander | cEDH signal | Top cuts | Entries |
|---|---|---|---|
| The Jolly Balloon Man | 0.57 | 5 | 6 |
| Niv-Mizzet, Parun | 0.50 | 5 | 31 |
| Etali, Primal Conqueror // Etali, Primal Sickness | 0.50 | 76 | 411 |

## The commander model

Tiers, worst first: bad, baseline, typical, good, great. Train counts: bad 4804, baseline 145, good 3223, great 3160, typical 854. Holdout counts: bad 1115, baseline 44, good 777, great 840, typical 186.

Cards with a rate: 9942. Pairs that lift: 200000. Commanders with a signal: 1552. Top lists shape: 27.6 lands at 2.19 over 3160 lists.

Features with no spread, dropped: playset_share, singleton_share.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.154 | 0.139 | +0.974 |
| `unseen_share` | 0.069 | 0.099 | -0.215 |
| `synergy` | 0.497 | 0.536 | +0.167 |
| `land` | 30.688 | 5.581 | -0.243 |
| `avg_mana_value` | 2.792 | 0.862 | -0.446 |
| `color_sources` | 1.081 | 0.308 | +0.000 |
| `tapped_share` | 0.071 | 0.083 | +0.352 |
| `mana_turn_four` | 4.712 | 0.614 | +0.761 |
| `hands_two_to_four_lands` | 0.651 | 0.087 | +0.529 |
| `curve_low` | 0.508 | 0.188 | +0.186 |
| `curve_high` | 0.152 | 0.133 | -0.183 |
| `ramp` | 17.664 | 7.582 | +0.057 |
| `draw` | 7.417 | 4.084 | -0.230 |
| `removal` | 6.380 | 3.637 | +0.088 |
| `wipe` | 1.776 | 1.974 | -0.156 |
| `interaction` | 7.439 | 4.578 | +0.196 |
| `empty_roles` | 0.049 | 0.220 | +0.001 |
| `fast_mana` | 5.210 | 4.825 | +0.271 |
| `game_changer` | 7.344 | 7.402 | +0.123 |
| `source_spread` | 0.095 | 0.122 | -0.478 |
| `cedh_signal` | 0.281 | 0.200 | +0.220 |
| `high_bracket_share` | 0.217 | 0.250 | +0.094 |
| `commander_decks` | 5.896 | 4.120 | -0.139 |

Thresholds: -3.170, -1.461, 0.660, 2.983. Loss: 1.061 over 3000 iterations. Defect detector accuracy on the holdout: 0.92, at a cut of 0.46.

The precon bar per broken axis, each precon over its own copies, with the cross pairs after it:

- lands: 1.00 of 189, cross 0.97 of 46364
- curve: 1.00 of 189, cross 0.96 of 46364
- colors: 1.00 of 178, cross 0.92 of 27525
- synergy: 0.84 of 232, cross 0.78 of 97145

The synergy check dropped 240 of 2958 synergy copies over the folds, at a floor of 0.10 and a unit of 0.5841 (D-652). By the axis the fit asked for: colors 17, copies 116, synergy 107.

The misses per axis over the folds (M-7). Both passed: the detector passed the precon and the copy, and the ladder put the copy at or above the precon. Precon flagged: the detector flagged the precon and passed the copy. Both flagged: it flagged both and read the precon as the more broken. Own copy: each precon against its own copy on the axis, with its misses by the same kinds. Moved: the mean absolute standardized delta between a precon and its own copy, for the corpus features and the three features that move most.

| Axis | Share | Both passed | Precon flagged | Both flagged | Own copy | Own misses | Corpus moved | Most moved |
|---|---|---|---|---|---|---|---|---|
| lands | 0.97 of 46364 | 604 | 394 | 520 | 189 of 189 | 0, 0, 0 | `card_rate` 0.03, `synergy` 0.39, `unseen_share` 0.34 | `land` 2.21, `hands_two_to_four_lands` 2.06, `color_sources` 1.27 |
| curve | 0.96 of 46364 | 266 | 471 | 1275 | 189 of 189 | 0, 0, 0 | `card_rate` 0.18, `synergy` 0.50, `unseen_share` 0.54 | `curve_high` 1.86, `avg_mana_value` 1.48, `curve_low` 1.24 |
| colors | 0.92 of 27525 | 798 | 662 | 750 | 178 of 178 | 0, 0, 0 | `card_rate` 0.00, `synergy` 0.00, `unseen_share` 0.00 | `source_spread` 3.72, `tapped_share` 2.77, `color_sources` 2.01 |
| synergy | 0.78 of 97145 | 6631 | 6769 | 7740 | 196 of 232 | 21, 9, 6 | `card_rate` 0.09, `synergy` 1.49, `unseen_share` 1.05 | `synergy` 1.49, `unseen_share` 1.05, `wipe` 0.72 |

The precons that lose most over the sampled pairs of the bar:

| Precon | Grade | Detector | Misses |
|---|---|---|---|
| Eternal Bargain (2013-11-01) | bad | 0.91 | 314 of 454 |
| Political Puppets (2011-06-17) | bad | 0.88 | 306 of 500 |
| Peer Through Time (2014-11-07) | bad | 0.85 | 304 of 500 |
| Heavenly Inferno (2011-06-17) | bad | 0.85 | 304 of 588 |
| Evasive Maneuvers (2017-06-09) | bad | 0.75 | 266 of 571 |
| Heavenly Inferno (2017-06-09) | bad | 0.81 | 211 of 454 |
| Evasive Maneuvers (2013-11-01) | bad | 0.79 | 222 of 500 |
| Plunder the Graves (2017-06-09) | bad | 0.81 | 259 of 588 |
| Swell the Host (2015-11-13) | bad | 0.80 | 257 of 588 |
| Quick Draw (2024-04-19) | bad | 0.77 | 218 of 500 |

The audit of the bar (M-8). The bar asks the ladder to score a precon above its own broken copy. Rank is the precon's place among the holdout precons by score, 0 the weakest, and the column reads the mean rank of the pairs that win and of the pairs that lose. A break that hurts every precon alike moves the two means together. A lower mean rank among the losers says the bar fails on the weak precons, which is where a broken copy is most likely the better deck. Card rate and synergy read the signed move, the copy less the precon, over the pairs that lose. A positive card rate says the copy holds cards the top lists play more than the precon does.

| Axis | Own pairs | Lost | Mean rank, won | Mean rank, lost | `card_rate` move | `synergy` move | Copy graded above |
|---|---|---|---|---|---|---|---|
| lands | 189 | 0 | 18.58 | - | - | - | 0 of 0 |
| curve | 189 | 0 | 18.58 | - | - | - | 0 of 0 |
| colors | 178 | 0 | 18.87 | - | - | - | 0 of 0 |
| synergy | 232 | 36 | 24.93 | 12.61 | 0.00 | -0.29 | 7 of 36 |

Confusion on the holdout over the folds, rows are the label and columns the grade, worst first:

- bad: [5730 12 146 17 0]
- baseline: [119 13 56 1 0]
- typical: [419 17 543 49 12]
- good: [257 11 262 1403 2067]
- great: [208 12 249 1391 2140]

The cards the top lists hold most: Sol Ring, Lotus Petal, Chrome Mox, Mana Vault, Mox Diamond, Arcane Signet, Mystic Remora, Mental Misstep, Flusterstorm, Rhystic Study, Force of Will, Mox Amber.

## The standard model

Tiers, worst first: bad, baseline, typical, good, great. Train counts: bad 320, baseline 7, good 3199, great 1358, typical 57. Holdout counts: bad 50, baseline 1, good 801, great 336, typical 9.

Cards with a rate: 815. Pairs that lift: 6032. Commanders with a signal: 0. Top lists shape: 23.1 lands at 2.49 over 1358 lists.

Features with no spread, dropped: game_changer, cedh_signal, high_bracket_share, commander_decks.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.117 | 0.061 | +1.379 |
| `unseen_share` | 0.016 | 0.081 | +0.104 |
| `synergy` | 1.590 | 0.737 | +0.969 |
| `land` | 22.999 | 2.396 | -0.144 |
| `avg_mana_value` | 2.544 | 0.659 | -0.966 |
| `color_sources` | 1.009 | 0.311 | +0.468 |
| `tapped_share` | 0.033 | 0.072 | -0.047 |
| `mana_turn_four` | 4.062 | 0.764 | +0.497 |
| `hands_two_to_four_lands` | 0.756 | 0.032 | +0.534 |
| `curve_low` | 0.635 | 0.175 | -0.387 |
| `curve_high` | 0.130 | 0.130 | +0.111 |
| `ramp` | 5.030 | 6.428 | -0.391 |
| `draw` | 4.892 | 3.441 | -0.080 |
| `removal` | 8.389 | 4.452 | +0.126 |
| `wipe` | 1.353 | 1.782 | +0.098 |
| `interaction` | 3.390 | 2.968 | +0.083 |
| `empty_roles` | 0.822 | 0.684 | -0.065 |
| `fast_mana` | 0.150 | 0.802 | +0.082 |
| `playset_share` | 0.600 | 0.182 | -0.010 |
| `singleton_share` | 0.088 | 0.088 | -0.476 |
| `source_spread` | 0.280 | 0.215 | +0.361 |

Thresholds: -4.483, -2.546, -0.634, 1.028. Loss: 1.076 over 3000 iterations. Defect detector accuracy on the holdout: 0.96, at a cut of 0.38.

The precon bar per broken axis, each precon over its own copies, with the cross pairs after it:

- lands: 1.00 of 8, cross 0.97 of 146
- curve: 1.00 of 8, cross 1.00 of 146
- colors: 1.00 of 4, cross 0.80 of 105
- copies: 1.00 of 3, cross 0.99 of 111
- synergy: 1.00 of 17, cross 0.86 of 222

The misses per axis over the folds (M-7). Both passed: the detector passed the precon and the copy, and the ladder put the copy at or above the precon. Precon flagged: the detector flagged the precon and passed the copy. Both flagged: it flagged both and read the precon as the more broken. Own copy: each precon against its own copy on the axis, with its misses by the same kinds. Moved: the mean absolute standardized delta between a precon and its own copy, for the corpus features and the three features that move most.

| Axis | Share | Both passed | Precon flagged | Both flagged | Own copy | Own misses | Corpus moved | Most moved |
|---|---|---|---|---|---|---|---|---|
| lands | 0.97 of 146 | 4 | 0 | 0 | 8 of 8 | 0, 0, 0 | `card_rate` 0.06, `synergy` 0.13, `unseen_share` 1.31 | `hands_two_to_four_lands` 4.19, `land` 3.28, `color_sources` 2.35 |
| curve | 1.00 of 146 | 0 | 0 | 0 | 8 of 8 | 0, 0, 0 | `card_rate` 0.22, `synergy` 0.22, `unseen_share` 2.78 | `avg_mana_value` 3.62, `curve_high` 3.25, `unseen_share` 2.78 |
| colors | 0.80 of 105 | 21 | 0 | 0 | 4 of 4 | 0, 0, 0 | `card_rate` 0.00, `synergy` 0.00, `unseen_share` 0.00 | `tapped_share` 3.12, `color_sources` 1.73, `source_spread` 1.56 |
| copies | 0.99 of 111 | 1 | 0 | 0 | 3 of 3 | 0, 0, 0 | `card_rate` 0.30, `synergy` 0.03, `unseen_share` 0.77 | `singleton_share` 5.00, `playset_share` 2.47, `empty_roles` 0.98 |
| synergy | 0.86 of 222 | 30 | 0 | 0 | 17 of 17 | 0, 0, 0 | `card_rate` 0.13, `synergy` 0.21, `unseen_share` 4.40 | `unseen_share` 4.40, `color_sources` 2.18, `playset_share` 1.82 |

The precons that lose most over the sampled pairs of the bar:

| Precon | Grade | Detector | Misses |
|---|---|---|---|
| Angels (2026-01-23) | bad | 0.27 | 18 of 105 |
| Raphael (2026-03-06) | bad | 0.17 | 8 of 65 |
| Donatello (2026-03-06) | bad | 0.27 | 8 of 100 |
| Leonardo (2026-03-06) | bad | 0.26 | 8 of 100 |
| Michealangelo (2026-03-06) | bad | 0.17 | 7 of 100 |
| Pirates (2026-01-23) | typical | 0.07 | 7 of 105 |
| Eerie (2026-04-24) | baseline | 0.04 | 0 of 50 |
| Lifegain (2026-04-24) | good | 0.01 | 0 of 105 |

| Axis | Own pairs | Lost | Mean rank, won | Mean rank, lost | `card_rate` move | `synergy` move | Copy graded above |
|---|---|---|---|---|---|---|---|
| lands | 8 | 0 | 0.75 | - | - | - | 0 of 0 |
| curve | 8 | 0 | 0.75 | - | - | - | 0 of 0 |
| colors | 4 | 0 | 0.75 | - | - | - | 0 of 0 |
| copies | 3 | 0 | 0.33 | - | - | - | 0 of 0 |
| synergy | 17 | 0 | 0.82 | - | - | - | 0 of 0 |

Confusion on the holdout over the folds, rows are the label and columns the grade, worst first:

- bad: [341 1 18 9 1]
- baseline: [5 1 1 1 0]
- typical: [18 6 17 16 9]
- good: [248 3 490 1604 1655]
- great: [107 2 181 599 805]

The cards the top lists hold most: Burst Lightning, Stock Up, Opt, Sleight of Hand, Llanowar Elves, Eddymurk Crab, Spell Pierce, Boomerang Basics, Stormchaser's Talent, Flow State, Badgermole Cub, Slickshot Show-Off.

## The modern model

Tiers, worst first: bad, baseline, typical, good, great. Train counts: bad 2205, baseline 266, good 3149, great 2021, typical 175. Holdout counts: bad 540, baseline 61, good 851, great 529, typical 47.

Cards with a rate: 891. Pairs that lift: 6786. Commanders with a signal: 0. Top lists shape: 21.7 lands at 2.44 over 2021 lists.

Features with no spread, dropped: cedh_signal, high_bracket_share, commander_decks.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.065 | 0.046 | +0.463 |
| `unseen_share` | 0.228 | 0.377 | -0.732 |
| `synergy` | 1.404 | 1.010 | +1.355 |
| `land` | 21.679 | 3.076 | -0.052 |
| `avg_mana_value` | 2.753 | 0.945 | -0.364 |
| `color_sources` | 0.935 | 0.316 | +0.320 |
| `tapped_share` | 0.124 | 0.141 | +0.144 |
| `mana_turn_four` | 3.844 | 0.797 | +0.027 |
| `hands_two_to_four_lands` | 0.732 | 0.059 | +0.260 |
| `curve_low` | 0.580 | 0.221 | -0.064 |
| `curve_high` | 0.180 | 0.158 | -0.190 |
| `ramp` | 5.348 | 6.806 | +0.090 |
| `draw` | 4.663 | 4.601 | +0.036 |
| `removal` | 7.758 | 4.572 | -0.197 |
| `wipe` | 1.008 | 1.746 | +0.014 |
| `interaction` | 3.030 | 3.800 | +0.015 |
| `empty_roles` | 1.031 | 0.844 | +0.015 |
| `fast_mana` | 0.553 | 1.506 | +0.010 |
| `game_changer` | 0.271 | 0.885 | -0.110 |
| `playset_share` | 0.569 | 0.236 | -0.550 |
| `singleton_share` | 0.145 | 0.157 | -0.652 |
| `source_spread` | 0.245 | 0.218 | -0.246 |

Thresholds: -3.805, -1.515, 0.851, 2.402. Loss: 1.056 over 3000 iterations. Defect detector accuracy on the holdout: 0.93, at a cut of 0.46.

The precon bar per broken axis, each precon over its own copies, with the cross pairs after it:

- lands: 1.00 of 327, cross 0.94 of 36476
- curve: 1.00 of 327, cross 0.98 of 36476
- colors: 0.99 of 75, cross 0.47 of 17066
- copies: 1.00 of 61, cross 0.83 of 18460
- synergy: 0.98 of 845, cross 0.85 of 70333

The misses per axis over the folds (M-7). Both passed: the detector passed the precon and the copy, and the ladder put the copy at or above the precon. Precon flagged: the detector flagged the precon and passed the copy. Both flagged: it flagged both and read the precon as the more broken. Own copy: each precon against its own copy on the axis, with its misses by the same kinds. Moved: the mean absolute standardized delta between a precon and its own copy, for the corpus features and the three features that move most.

| Axis | Share | Both passed | Precon flagged | Both flagged | Own copy | Own misses | Corpus moved | Most moved |
|---|---|---|---|---|---|---|---|---|
| lands | 0.94 of 36476 | 1310 | 327 | 634 | 327 of 327 | 0, 0, 0 | `card_rate` 0.01, `synergy` 0.00, `unseen_share` 0.09 | `land` 2.57, `hands_two_to_four_lands` 2.47, `color_sources` 1.54 |
| curve | 0.98 of 36476 | 365 | 132 | 355 | 327 of 327 | 0, 0, 0 | `card_rate` 0.04, `synergy` 0.00, `unseen_share` 0.20 | `curve_high` 2.68, `avg_mana_value` 2.22, `curve_low` 1.70 |
| colors | 0.47 of 17066 | 7033 | 1643 | 344 | 74 of 75 | 1, 0, 0 | `card_rate` 0.00, `synergy` 0.00, `unseen_share` 0.00 | `tapped_share` 1.25, `color_sources` 1.14, `source_spread` 0.96 |
| copies | 0.83 of 18460 | 2331 | 566 | 325 | 61 of 61 | 0, 0, 0 | `card_rate` 0.04, `synergy` 0.00, `unseen_share` 0.19 | `singleton_share` 3.56, `playset_share` 2.38, `empty_roles` 0.91 |
| synergy | 0.85 of 70333 | 6365 | 2326 | 2018 | 832 of 845 | 10, 0, 3 | `card_rate` 0.03, `synergy` 0.00, `unseen_share` 0.21 | `playset_share` 1.31, `empty_roles` 0.64, `removal` 0.61 |

The precons that lose most over the sampled pairs of the bar:

| Precon | Grade | Detector | Misses |
|---|---|---|---|
| Sultai Schemers (2014-09-26) | bad | 0.87 | 232 of 344 |
| Beasts (2005-08-01) | bad | 0.86 | 129 of 240 |
| Illusionary Might (2011-08-12) | bad | 0.83 | 122 of 240 |
| Jeskai Monks (2014-09-26) | bad | 0.79 | 165 of 333 |
| Thallids (2005-08-01) | bad | 0.82 | 114 of 240 |
| Second Sun Control (2018-04-06) | bad | 0.77 | 155 of 333 |
| Stampede (2004-06-04) | bad | 0.77 | 147 of 327 |
| Wizards (2005-08-01) | bad | 0.80 | 130 of 307 |
| Grave Advantage (2015-01-23) | bad | 0.74 | 97 of 240 |
| Stampede of Beasts (2010-07-16) | bad | 0.71 | 131 of 333 |

| Axis | Own pairs | Lost | Mean rank, won | Mean rank, lost | `card_rate` move | `synergy` move | Copy graded above |
|---|---|---|---|---|---|---|---|
| lands | 327 | 0 | 32.83 | - | - | - | 0 of 0 |
| curve | 327 | 0 | 32.83 | - | - | - | 0 of 0 |
| colors | 75 | 1 | 37.51 | 16.00 | 0.00 | 0.00 | 0 of 1 |
| copies | 61 | 0 | 37.54 | - | - | - | 0 of 0 |
| synergy | 845 | 13 | 32.07 | 34.15 | 0.07 | 0.00 | 0 of 13 |

Confusion on the holdout over the folds, rows are the label and columns the grade, worst first:

- bad: [2548 64 130 0 3]
- baseline: [207 114 6 0 0]
- typical: [34 3 73 74 38]
- good: [172 0 772 1518 1538]
- great: [83 0 515 760 1192]

The cards the top lists hold most: Solitude, Quantum Riddler, Thoughtseize, Force of Negation, Prismatic Ending, Fatal Push, Ephemerate, Mishra's Bauble, Galvanic Discharge, Kozilek's Command, Malevolent Rumble, Ragavan, Nimble Pilferer.

## Run

- Suite `quality`, run `pr14b-quality-gate-run19`, on 2026-09-10, commit `d692861`.
- Roles: none, no provider call.
- Versions: card snapshot 2026-09-04, commanders_day `2026-09-07`, quality_model `20260910T231306Z`.
- Calls: 0. Cost: unpriced. Time: 106 seconds.

## Verdict

Verdict: PASS. Time: 106 seconds, no provider call.
