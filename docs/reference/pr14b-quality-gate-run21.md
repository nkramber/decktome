# PR-14B quality gate

Run date: 2026-09-12. Card snapshot: 2026-09-04. Model: 20260912T153537Z.

## Separation on the holdout

The bars read 5 folds, every list holdout once, and the pairs of a bar sample evenly under the cap of 20000 (M-7). The precon bar reads each precon against its own broken copies (D-573). The cross pairs, every precon against every copy, stand as information.

The synergy check drops a copy whose break lowered the `synergy` feature by less than 0.10 standard deviations of the real training lists (D-652). It reads commander, and every other format keeps each copy (D-653). Synthetic counts the copies the engine made, and Immaterial the ones the check dropped, each copy once over the folds.

| Format | Lists | Used | Synthetic | Immaterial | Holdout | Great over precon | Precon over own copy | Precon over bad, cross | Accuracy |
|---|---|---|---|---|---|---|---|---|---|
| commander | 23628 | 9229 | 6145 | 249 | 15125 | 0.96 of 97450 (bar 0.90) | 0.95 of 790 (bar 0.95) | 0.87 of 99933 | 0.54 |
| standard | 8228 | 5768 | 370 | 0 | 6138 | 0.92 of 2750 (bar 0.90) | 1.00 of 40 (bar 0.95) | 0.91 of 730 | 0.45 |
| modern | 20292 | 7099 | 2745 | 0 | 9844 | 0.99 of 98691 (bar 0.90) | 0.99 of 1635 (bar 0.95) | 0.85 of 99754 | 0.56 |

## The three numbers (D-648)

The precon bar is a proxy, so every quality item reads two more numbers beside it. The model this run fitted grades the built decks of `pr8-deck-gate-run16.md`. The judge's tiers come from `pr14b-quality-judge-run5.md`. One judge run serves every model, so the agreement holds no judge noise. No bar reads the last two, and an item that moves none of the three is dropped. A change that moves a tier also reads the broken copies graded bad (D-674).

| Number | Read |
|---|---|
| Commander synergy axis, precon over own copy | 0.84 of 234 |
| Commander broken copies graded bad (D-674) | 4130 of 5896 |
| Built decks graded bad | 8 of 25 |
| Judge agreement | 9 of 25 |

| # | Deck | Format | Grade | Judge | Agree |
|---|---|---|---|---|---|
| 1 | lifegain Commander, any card | Commander | bad | typical | no |
| 2 | aristocrats Commander, owned first | Commander | typical | good | no |
| 3 | artifacts Commander, bracket 4 | Commander | baseline | typical | no |
| 4 | dinosaur tribal, bracket 2 | Commander | baseline | typical | no |
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
| 20 | the Hobbit family, a delegated commander | Commander | baseline | baseline | yes |
| 21 | the Hobbit family, mana from outside | Commander | baseline | typical | no |
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

Tiers, worst first: bad, baseline, typical, good, great. Train counts: bad 4807, baseline 145, good 3223, great 3160, typical 854. Holdout counts: bad 1114, baseline 44, good 777, great 840, typical 186.

Cards with a rate: 9962. Pairs that lift: 200000. Commanders with a signal: 1552. Top lists shape: 27.5 lands at 2.19 over 3160 lists.

Features with no spread, dropped: playset_share, singleton_share.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.154 | 0.139 | +0.966 |
| `unseen_share` | 0.069 | 0.099 | -0.217 |
| `synergy` | 0.497 | 0.536 | +0.168 |
| `land` | 30.613 | 5.558 | -0.271 |
| `avg_mana_value` | 2.790 | 0.859 | -0.451 |
| `color_sources` | 1.080 | 0.308 | -0.000 |
| `tapped_share` | 0.072 | 0.083 | +0.350 |
| `mana_turn_four` | 4.712 | 0.614 | +0.758 |
| `hands_two_to_four_lands` | 0.650 | 0.086 | +0.526 |
| `curve_low` | 0.508 | 0.188 | +0.175 |
| `curve_high` | 0.152 | 0.133 | -0.167 |
| `ramp` | 17.751 | 7.582 | +0.058 |
| `draw` | 7.424 | 4.092 | -0.231 |
| `removal` | 6.381 | 3.635 | +0.083 |
| `wipe` | 1.780 | 1.970 | -0.158 |
| `interaction` | 7.445 | 4.574 | +0.193 |
| `empty_roles` | 0.049 | 0.221 | +0.002 |
| `fast_mana` | 5.214 | 4.821 | +0.258 |
| `game_changer` | 7.339 | 7.404 | +0.127 |
| `source_spread` | 0.095 | 0.122 | -0.491 |
| `cedh_signal` | 0.281 | 0.200 | +0.221 |
| `high_bracket_share` | 0.217 | 0.250 | +0.089 |
| `commander_decks` | 5.894 | 4.120 | -0.140 |

Thresholds: -3.171, -1.468, 0.651, 2.976. Loss: 1.062 over 3000 iterations. Defect detector accuracy on the holdout: 0.92, at a cut of 0.46.

The precon bar per broken axis, each precon over its own copies, with the cross pairs after it:

- lands: 1.00 of 189, cross 0.97 of 46364
- curve: 1.00 of 189, cross 0.95 of 46364
- colors: 1.00 of 178, cross 0.92 of 27525
- synergy: 0.84 of 234, cross 0.78 of 97281

The synergy check dropped 249 of 2958 synergy copies over the folds, at a floor of 0.10 and a unit of 0.5838 (D-652). By the axis the fit asked for: colors 21, copies 113, synergy 115.

The misses per axis over the folds (M-7). Both passed: the detector passed the precon and the copy, and the ladder put the copy at or above the precon. Precon flagged: the detector flagged the precon and passed the copy. Both flagged: it flagged both and read the precon as the more broken. Own copy: each precon against its own copy on the axis, with its misses by the same kinds. Moved: the mean absolute standardized delta between a precon and its own copy, for the corpus features and the three features that move most.

| Axis | Share | Both passed | Precon flagged | Both flagged | Own copy | Own misses | Corpus moved | Most moved |
|---|---|---|---|---|---|---|---|---|
| lands | 0.97 of 46364 | 605 | 406 | 545 | 189 of 189 | 0, 0, 0 | `card_rate` 0.03, `synergy` 0.38, `unseen_share` 0.33 | `land` 2.22, `hands_two_to_four_lands` 2.08, `color_sources` 1.28 |
| curve | 0.95 of 46364 | 260 | 450 | 1407 | 189 of 189 | 0, 0, 0 | `card_rate` 0.18, `synergy` 0.50, `unseen_share` 0.54 | `curve_high` 1.86, `avg_mana_value` 1.49, `curve_low` 1.24 |
| colors | 0.92 of 27525 | 760 | 656 | 770 | 178 of 178 | 0, 0, 0 | `card_rate` 0.00, `synergy` 0.00, `unseen_share` 0.00 | `source_spread` 3.72, `tapped_share` 2.77, `color_sources` 2.01 |
| synergy | 0.78 of 97281 | 7108 | 7154 | 7476 | 196 of 234 | 23, 7, 8 | `card_rate` 0.09, `synergy` 1.48, `unseen_share` 1.04 | `synergy` 1.48, `unseen_share` 1.04, `wipe` 0.71 |

The precons that lose most over the sampled pairs of the bar:

| Precon | Grade | Detector | Misses |
|---|---|---|---|
| Eternal Bargain (2013-11-01) | bad | 0.91 | 309 of 454 |
| Political Puppets (2011-06-17) | bad | 0.88 | 313 of 500 |
| Peer Through Time (2014-11-07) | bad | 0.85 | 284 of 500 |
| Heavenly Inferno (2011-06-17) | bad | 0.85 | 301 of 588 |
| Evasive Maneuvers (2013-11-01) | bad | 0.80 | 242 of 500 |
| Heavenly Inferno (2017-06-09) | bad | 0.81 | 217 of 454 |
| Evasive Maneuvers (2017-06-09) | baseline | 0.76 | 267 of 571 |
| Swell the Host (2015-11-13) | bad | 0.80 | 267 of 588 |
| Plunder the Graves (2015-11-13) | bad | 0.78 | 251 of 555 |
| Plunder the Graves (2017-06-09) | bad | 0.80 | 262 of 588 |

The audit of the bar (M-8). The bar asks the ladder to score a precon above its own broken copy. Rank is the precon's place among the holdout precons by score, 0 the weakest, and the column reads the mean rank of the pairs that win and of the pairs that lose. A break that hurts every precon alike moves the two means together. A lower mean rank among the losers says the bar fails on the weak precons, which is where a broken copy is most likely the better deck. Card rate and synergy read the signed move, the copy less the precon, over the pairs that lose. A positive card rate says the copy holds cards the top lists play more than the precon does.

| Axis | Own pairs | Lost | Mean rank, won | Mean rank, lost | `card_rate` move | `synergy` move | Copy graded above |
|---|---|---|---|---|---|---|---|
| lands | 189 | 0 | 18.58 | - | - | - | 0 of 0 |
| curve | 189 | 0 | 18.58 | - | - | - | 0 of 0 |
| colors | 178 | 0 | 18.85 | - | - | - | 0 of 0 |
| synergy | 234 | 38 | 24.31 | 13.24 | 0.01 | -0.26 | 20 of 38 |

Confusion on the holdout over the folds, rows are the label and columns the grade, worst first:

- bad: [4130 1341 419 6 0]
- baseline: [61 70 57 1 0]
- typical: [72 286 620 52 10]
- good: [248 102 290 1310 2050]
- great: [244 87 257 1301 2111]

The cards the top lists hold most: Sol Ring, Lotus Petal, Chrome Mox, Mana Vault, Mox Diamond, Arcane Signet, Mystic Remora, Mental Misstep, Flusterstorm, Rhystic Study, Force of Will, Mox Amber.

## The standard model

Tiers, worst first: bad, baseline, typical, good, great. Train counts: bad 320, baseline 7, good 3199, great 1358, typical 57. Holdout counts: bad 50, baseline 1, good 801, great 336, typical 9.

Cards with a rate: 817. Pairs that lift: 6076. Commanders with a signal: 0. Top lists shape: 23.1 lands at 2.49 over 1358 lists.

Features with no spread, dropped: game_changer, cedh_signal, high_bracket_share, commander_decks.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.117 | 0.061 | +1.384 |
| `unseen_share` | 0.016 | 0.081 | +0.104 |
| `synergy` | 1.591 | 0.738 | +0.958 |
| `land` | 22.975 | 2.380 | -0.149 |
| `avg_mana_value` | 2.545 | 0.657 | -0.958 |
| `color_sources` | 1.009 | 0.311 | +0.473 |
| `tapped_share` | 0.033 | 0.072 | -0.064 |
| `mana_turn_four` | 4.066 | 0.772 | +0.452 |
| `hands_two_to_four_lands` | 0.756 | 0.032 | +0.565 |
| `curve_low` | 0.634 | 0.174 | -0.374 |
| `curve_high` | 0.130 | 0.130 | +0.085 |
| `ramp` | 5.057 | 6.414 | -0.357 |
| `draw` | 4.890 | 3.447 | -0.071 |
| `removal` | 8.394 | 4.454 | +0.106 |
| `wipe` | 1.350 | 1.774 | +0.137 |
| `interaction` | 3.393 | 2.962 | +0.091 |
| `empty_roles` | 0.814 | 0.681 | -0.042 |
| `fast_mana` | 0.148 | 0.797 | +0.113 |
| `playset_share` | 0.599 | 0.182 | -0.004 |
| `singleton_share` | 0.088 | 0.089 | -0.475 |
| `source_spread` | 0.280 | 0.215 | +0.363 |

Thresholds: -4.493, -2.545, -0.629, 1.032. Loss: 1.075 over 3000 iterations. Defect detector accuracy on the holdout: 0.96, at a cut of 0.42.

The precon bar per broken axis, each precon over its own copies, with the cross pairs after it:

- lands: 1.00 of 8, cross 0.95 of 146
- curve: 1.00 of 8, cross 1.00 of 146
- colors: 1.00 of 4, cross 0.74 of 105
- copies: 1.00 of 3, cross 0.99 of 111
- synergy: 1.00 of 17, cross 0.86 of 222

The misses per axis over the folds (M-7). Both passed: the detector passed the precon and the copy, and the ladder put the copy at or above the precon. Precon flagged: the detector flagged the precon and passed the copy. Both flagged: it flagged both and read the precon as the more broken. Own copy: each precon against its own copy on the axis, with its misses by the same kinds. Moved: the mean absolute standardized delta between a precon and its own copy, for the corpus features and the three features that move most.

| Axis | Share | Both passed | Precon flagged | Both flagged | Own copy | Own misses | Corpus moved | Most moved |
|---|---|---|---|---|---|---|---|---|
| lands | 0.95 of 146 | 7 | 0 | 0 | 8 of 8 | 0, 0, 0 | `card_rate` 0.06, `synergy` 0.13, `unseen_share` 1.50 | `hands_two_to_four_lands` 4.20, `land` 3.30, `color_sources` 2.35 |
| curve | 1.00 of 146 | 0 | 0 | 0 | 8 of 8 | 0, 0, 0 | `card_rate` 0.22, `synergy` 0.19, `unseen_share` 2.70 | `avg_mana_value` 3.44, `curve_high` 3.02, `unseen_share` 2.70 |
| colors | 0.74 of 105 | 27 | 0 | 0 | 4 of 4 | 0, 0, 0 | `card_rate` 0.00, `synergy` 0.00, `unseen_share` 0.00 | `tapped_share` 3.12, `color_sources` 1.73, `source_spread` 1.56 |
| copies | 0.99 of 111 | 1 | 0 | 0 | 3 of 3 | 0, 0, 0 | `card_rate` 0.30, `synergy` 0.03, `unseen_share` 0.77 | `singleton_share` 5.00, `playset_share` 2.47, `empty_roles` 0.98 |
| synergy | 0.86 of 222 | 31 | 0 | 0 | 17 of 17 | 0, 0, 0 | `card_rate` 0.12, `synergy` 0.22, `unseen_share` 4.50 | `unseen_share` 4.50, `color_sources` 2.23, `playset_share` 1.78 |

The precons that lose most over the sampled pairs of the bar:

| Precon | Grade | Detector | Misses |
|---|---|---|---|
| Angels (2026-01-23) | bad | 0.26 | 18 of 105 |
| Raphael (2026-03-06) | bad | 0.15 | 8 of 65 |
| Donatello (2026-03-06) | bad | 0.28 | 11 of 100 |
| Leonardo (2026-03-06) | bad | 0.28 | 11 of 100 |
| Michealangelo (2026-03-06) | bad | 0.19 | 11 of 100 |
| Pirates (2026-01-23) | typical | 0.06 | 6 of 105 |
| Eerie (2026-04-24) | baseline | 0.04 | 1 of 50 |
| Lifegain (2026-04-24) | good | 0.01 | 0 of 105 |

| Axis | Own pairs | Lost | Mean rank, won | Mean rank, lost | `card_rate` move | `synergy` move | Copy graded above |
|---|---|---|---|---|---|---|---|
| lands | 8 | 0 | 0.75 | - | - | - | 0 of 0 |
| curve | 8 | 0 | 0.75 | - | - | - | 0 of 0 |
| colors | 4 | 0 | 0.75 | - | - | - | 0 of 0 |
| copies | 3 | 0 | 0.33 | - | - | - | 0 of 0 |
| synergy | 17 | 0 | 0.82 | - | - | - | 0 of 0 |

Confusion on the holdout over the folds, rows are the label and columns the grade, worst first:

- bad: [336 3 21 8 2]
- baseline: [5 1 1 1 0]
- typical: [17 4 20 15 10]
- good: [251 3 492 1625 1629]
- great: [111 2 177 606 798]

The cards the top lists hold most: Burst Lightning, Stock Up, Opt, Sleight of Hand, Llanowar Elves, Eddymurk Crab, Spell Pierce, Boomerang Basics, Stormchaser's Talent, Flow State, Badgermole Cub, Slickshot Show-Off.

## The modern model

Tiers, worst first: bad, baseline, typical, good, great. Train counts: bad 2205, baseline 266, good 3149, great 2021, typical 175. Holdout counts: bad 540, baseline 61, good 851, great 529, typical 47.

Cards with a rate: 891. Pairs that lift: 6786. Commanders with a signal: 0. Top lists shape: 21.7 lands at 2.44 over 2021 lists.

Features with no spread, dropped: cedh_signal, high_bracket_share, commander_decks.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.065 | 0.046 | +0.468 |
| `unseen_share` | 0.228 | 0.377 | -0.725 |
| `synergy` | 1.404 | 1.010 | +1.354 |
| `land` | 21.673 | 3.073 | -0.061 |
| `avg_mana_value` | 2.755 | 0.947 | -0.362 |
| `color_sources` | 0.935 | 0.316 | +0.317 |
| `tapped_share` | 0.124 | 0.141 | +0.146 |
| `mana_turn_four` | 3.844 | 0.797 | +0.018 |
| `hands_two_to_four_lands` | 0.732 | 0.059 | +0.274 |
| `curve_low` | 0.580 | 0.222 | -0.065 |
| `curve_high` | 0.180 | 0.158 | -0.198 |
| `ramp` | 5.346 | 6.809 | +0.103 |
| `draw` | 4.658 | 4.595 | +0.043 |
| `removal` | 7.770 | 4.566 | -0.200 |
| `wipe` | 1.000 | 1.742 | +0.022 |
| `interaction` | 3.031 | 3.815 | +0.018 |
| `empty_roles` | 1.030 | 0.845 | +0.021 |
| `fast_mana` | 0.556 | 1.509 | +0.007 |
| `game_changer` | 0.272 | 0.887 | -0.113 |
| `playset_share` | 0.569 | 0.236 | -0.550 |
| `singleton_share` | 0.145 | 0.158 | -0.654 |
| `source_spread` | 0.245 | 0.218 | -0.245 |

Thresholds: -3.808, -1.514, 0.854, 2.406. Loss: 1.055 over 3000 iterations. Defect detector accuracy on the holdout: 0.93, at a cut of 0.46.

The precon bar per broken axis, each precon over its own copies, with the cross pairs after it:

- lands: 1.00 of 327, cross 0.94 of 36476
- curve: 1.00 of 327, cross 0.98 of 36476
- colors: 0.99 of 75, cross 0.47 of 17066
- copies: 1.00 of 61, cross 0.83 of 18460
- synergy: 0.98 of 845, cross 0.85 of 70333

The misses per axis over the folds (M-7). Both passed: the detector passed the precon and the copy, and the ladder put the copy at or above the precon. Precon flagged: the detector flagged the precon and passed the copy. Both flagged: it flagged both and read the precon as the more broken. Own copy: each precon against its own copy on the axis, with its misses by the same kinds. Moved: the mean absolute standardized delta between a precon and its own copy, for the corpus features and the three features that move most.

| Axis | Share | Both passed | Precon flagged | Both flagged | Own copy | Own misses | Corpus moved | Most moved |
|---|---|---|---|---|---|---|---|---|
| lands | 0.94 of 36476 | 1344 | 328 | 648 | 327 of 327 | 0, 0, 0 | `card_rate` 0.01, `synergy` 0.00, `unseen_share` 0.08 | `land` 2.57, `hands_two_to_four_lands` 2.47, `color_sources` 1.54 |
| curve | 0.98 of 36476 | 332 | 131 | 345 | 327 of 327 | 0, 0, 0 | `card_rate` 0.04, `synergy` 0.00, `unseen_share` 0.21 | `curve_high` 2.68, `avg_mana_value` 2.22, `curve_low` 1.70 |
| colors | 0.47 of 17066 | 7016 | 1772 | 312 | 74 of 75 | 1, 0, 0 | `card_rate` 0.00, `synergy` 0.00, `unseen_share` 0.00 | `tapped_share` 1.25, `color_sources` 1.13, `source_spread` 0.96 |
| copies | 0.83 of 18460 | 2270 | 616 | 329 | 61 of 61 | 0, 0, 0 | `card_rate` 0.04, `synergy` 0.00, `unseen_share` 0.18 | `singleton_share` 3.56, `playset_share` 2.38, `empty_roles` 0.89 |
| synergy | 0.85 of 70333 | 6253 | 2399 | 1954 | 831 of 845 | 10, 0, 4 | `card_rate` 0.03, `synergy` 0.00, `unseen_share` 0.21 | `playset_share` 1.31, `empty_roles` 0.65, `curve_high` 0.60 |

The precons that lose most over the sampled pairs of the bar:

| Precon | Grade | Detector | Misses |
|---|---|---|---|
| Sultai Schemers (2014-09-26) | bad | 0.87 | 233 of 344 |
| Beasts (2005-08-01) | bad | 0.86 | 130 of 240 |
| Jeskai Monks (2014-09-26) | bad | 0.79 | 165 of 333 |
| Illusionary Might (2011-08-12) | bad | 0.83 | 118 of 240 |
| Thallids (2005-08-01) | bad | 0.83 | 117 of 240 |
| Second Sun Control (2018-04-06) | bad | 0.76 | 148 of 333 |
| Stampede (2004-06-04) | bad | 0.78 | 145 of 327 |
| Wizards (2005-08-01) | bad | 0.80 | 135 of 307 |
| Grave Advantage (2015-01-23) | bad | 0.74 | 97 of 240 |
| Stampede of Beasts (2010-07-16) | bad | 0.72 | 130 of 333 |

| Axis | Own pairs | Lost | Mean rank, won | Mean rank, lost | `card_rate` move | `synergy` move | Copy graded above |
|---|---|---|---|---|---|---|---|
| lands | 327 | 0 | 32.83 | - | - | - | 0 of 0 |
| curve | 327 | 0 | 32.83 | - | - | - | 0 of 0 |
| colors | 75 | 1 | 37.73 | 16.00 | 0.00 | 0.00 | 0 of 1 |
| copies | 61 | 0 | 37.89 | - | - | - | 0 of 0 |
| synergy | 845 | 14 | 31.98 | 36.43 | 0.06 | -0.00 | 1 of 14 |

Confusion on the holdout over the folds, rows are the label and columns the grade, worst first:

- bad: [2555 63 124 0 3]
- baseline: [207 114 6 0 0]
- typical: [34 3 74 74 37]
- good: [171 0 768 1528 1533]
- great: [84 0 513 757 1196]

The cards the top lists hold most: Solitude, Quantum Riddler, Thoughtseize, Force of Negation, Prismatic Ending, Fatal Push, Ephemerate, Mishra's Bauble, Galvanic Discharge, Kozilek's Command, Malevolent Rumble, Ragavan, Nimble Pilferer.

## Run

- Suite `quality`, run `pr14b-quality-gate-run21`, on 2026-09-12, commit `5ef19ec`.
- Roles: none, no provider call.
- Versions: card snapshot 2026-09-04, commanders_day `2026-09-07`, quality_model `20260912T153537Z`.
- Calls: 0. Cost: unpriced. Time: 105 seconds.

## Verdict

Verdict: PASS. Time: 105 seconds, no provider call.
