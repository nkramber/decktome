# PR-14B quality gate

Run date: 2026-09-14. Card snapshot: 2026-09-04. Model: 20260914T184742Z.

## Separation on the holdout

The bars read 5 folds, every list holdout once, and the pairs of a bar sample evenly under the cap of 20000 (M-7). The precon bar reads each precon against its own broken copies (D-573). The cross pairs, every precon against every copy, stand as information.

The synergy check drops a copy whose break lowered the `synergy` feature by less than 0.10 standard deviations of the real training lists (D-652). It reads commander, and every other format keeps each copy (D-653). Synthetic counts the copies the engine made, and Immaterial the ones the check dropped, each copy once over the folds.

| Format | Lists | Used | Synthetic | Immaterial | Holdout | Great over precon | Precon over own copy | Precon over bad, cross | Accuracy |
|---|---|---|---|---|---|---|---|---|---|
| commander | 25420 | 9229 | 6145 | 238 | 15136 | 0.97 of 97514 (bar 0.90) | 0.96 of 784 (bar 0.95) | 0.87 of 99933 | 0.56 |
| standard | 10620 | 6239 | 740 | 0 | 6979 | 0.95 of 3367 (bar 0.90) | 1.00 of 40 (bar 0.95) | 0.87 of 1355 | 0.48 |
| modern | 24081 | 7773 | 3115 | 0 | 10888 | 0.99 of 98711 (bar 0.90) | 0.99 of 1635 (bar 0.95) | 0.85 of 99754 | 0.56 |

## The three numbers (D-648)

The precon bar is a proxy, so every quality item reads two more numbers beside it. The model this run fitted grades the built decks of `pr8-deck-gate-run16.md`. The judge's tiers come from `pr14b-quality-judge-run5.md`. One judge run serves every model, so the agreement holds no judge noise. No bar reads the last two, and an item that moves none of the three is dropped. A change that moves a tier also reads the broken copies graded bad (D-674).

| Number | Read |
|---|---|
| Commander synergy axis, precon over own copy | 0.85 of 228 |
| Commander broken copies graded bad (D-674) | 4168 of 5907 |
| Built decks graded bad | 9 of 25 |
| Judge agreement | 9 of 25 |

| # | Deck | Format | Grade | Judge | Agree |
|---|---|---|---|---|---|
| 1 | lifegain Commander, any card | Commander | bad | typical | no |
| 2 | aristocrats Commander, owned first | Commander | typical | good | no |
| 3 | artifacts Commander, bracket 4 | Commander | baseline | typical | no |
| 4 | dinosaur tribal, bracket 2 | Commander | baseline | typical | no |
| 5 | blink Commander, owned first | Commander | typical | typical | yes |
| 6 | Modern tempo, tournament | Modern | typical | typical | yes |
| 7 | Modern burn, casual | Modern | typical | bad | no |
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
| 21 | the Hobbit family, mana from outside | Commander | bad | typical | no |
| 22 | a set family and a card from outside it | Commander | bad | typical | no |
| 23 | two set families at once | Commander | typical | typical | yes |
| 24 | a 60-card deck from one set | Modern | baseline | bad | no |
| 25 | use no card of an owned precon | Commander | baseline | typical | no |

## The bracket 5 offer

Commander reads of 2026-09-14: 1521 rows, 897 with Topdeck.gg entries.

| Commander | cEDH signal | Top cuts | Entries |
|---|---|---|---|
| Abaddon the Despoiler | 0.58 | 4 | 4 |
| The Jolly Balloon Man | 0.57 | 5 | 6 |
| Niv-Mizzet, Parun | 0.50 | 5 | 30 |

## The commander model

Tiers, worst first: bad, baseline, typical, good, great. Train counts: bad 4822, baseline 145, good 3222, great 3162, typical 854. Holdout counts: bad 1112, baseline 44, good 778, great 838, typical 186.

Cards with a rate: 9574. Pairs that lift: 200000. Commanders with a signal: 1507. Top lists shape: 27.3 lands at 2.18 over 3162 lists.

Features with no spread, dropped: playset_share, singleton_share.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.156 | 0.139 | +1.073 |
| `unseen_share` | 0.075 | 0.106 | -0.193 |
| `synergy` | 0.468 | 0.469 | +0.263 |
| `land` | 30.532 | 5.542 | -0.282 |
| `avg_mana_value` | 2.785 | 0.859 | -0.452 |
| `color_sources` | 1.080 | 0.307 | +0.004 |
| `tapped_share` | 0.071 | 0.083 | +0.361 |
| `mana_turn_four` | 4.718 | 0.616 | +0.775 |
| `hands_two_to_four_lands` | 0.649 | 0.087 | +0.527 |
| `curve_low` | 0.509 | 0.187 | +0.204 |
| `curve_high` | 0.151 | 0.133 | -0.149 |
| `ramp` | 17.895 | 7.625 | +0.042 |
| `draw` | 7.377 | 4.096 | -0.225 |
| `removal` | 6.350 | 3.625 | +0.077 |
| `wipe` | 1.766 | 1.976 | -0.140 |
| `interaction` | 7.457 | 4.576 | +0.194 |
| `empty_roles` | 0.049 | 0.221 | -0.010 |
| `fast_mana` | 5.242 | 4.808 | +0.252 |
| `game_changer` | 7.361 | 7.327 | +0.159 |
| `source_spread` | 0.095 | 0.123 | -0.485 |
| `cedh_signal` | 0.281 | 0.201 | +0.204 |
| `high_bracket_share` | 0.223 | 0.253 | +0.066 |
| `commander_decks` | 5.963 | 4.114 | -0.148 |

Thresholds: -3.311, -1.577, 0.663, 3.080. Loss: 1.037 over 3000 iterations. Defect detector accuracy on the holdout: 0.92, at a cut of 0.46.

The precon bar per broken axis, each precon over its own copies, with the cross pairs after it:

- lands: 1.00 of 189, cross 0.97 of 46364
- curve: 1.00 of 189, cross 0.95 of 46364
- colors: 1.00 of 178, cross 0.92 of 27525
- synergy: 0.85 of 228, cross 0.78 of 97485

The synergy check dropped 238 of 2958 synergy copies over the folds, at a floor of 0.10 and a unit of 0.5119 (D-652). By the axis the fit asked for: colors 19, copies 109, synergy 110.

The misses per axis over the folds (M-7). Both passed: the detector passed the precon and the copy, and the ladder put the copy at or above the precon. Precon flagged: the detector flagged the precon and passed the copy. Both flagged: it flagged both and read the precon as the more broken. Own copy: each precon against its own copy on the axis, with its misses by the same kinds. Moved: the mean absolute standardized delta between a precon and its own copy, for the corpus features and the three features that move most.

| Axis | Share | Both passed | Precon flagged | Both flagged | Own copy | Own misses | Corpus moved | Most moved |
|---|---|---|---|---|---|---|---|---|
| lands | 0.97 of 46364 | 687 | 386 | 525 | 189 of 189 | 0, 0, 0 | `card_rate` 0.03, `synergy` 0.32, `unseen_share` 0.30 | `land` 2.23, `hands_two_to_four_lands` 2.07, `color_sources` 1.28 |
| curve | 0.95 of 46364 | 374 | 496 | 1285 | 189 of 189 | 0, 0, 0 | `card_rate` 0.18, `synergy` 0.43, `unseen_share` 0.51 | `curve_high` 1.87, `avg_mana_value` 1.48, `curve_low` 1.24 |
| colors | 0.92 of 27525 | 864 | 587 | 749 | 178 of 178 | 0, 0, 0 | `card_rate` 0.00, `synergy` 0.00, `unseen_share` 0.00 | `source_spread` 3.70, `tapped_share` 2.76, `color_sources` 2.01 |
| synergy | 0.78 of 97485 | 7458 | 6671 | 7121 | 193 of 228 | 25, 6, 4 | `card_rate` 0.09, `synergy` 1.23, `unseen_share` 0.95 | `synergy` 1.23, `unseen_share` 0.95, `wipe` 0.68 |

The precons that lose most over the sampled pairs of the bar:

| Precon | Grade | Detector | Misses |
|---|---|---|---|
| Eternal Bargain (2013-11-01) | bad | 0.90 | 289 of 454 |
| Peer Through Time (2014-11-07) | bad | 0.87 | 303 of 500 |
| Political Puppets (2011-06-17) | bad | 0.85 | 299 of 500 |
| Heavenly Inferno (2011-06-17) | bad | 0.85 | 310 of 588 |
| Undead Unleashed (2021-09-24) | bad | 0.83 | 292 of 588 |
| Heavenly Inferno (2017-06-09) | bad | 0.81 | 218 of 454 |
| Evasive Maneuvers (2013-11-01) | bad | 0.78 | 229 of 500 |
| Plunder the Graves (2015-11-13) | bad | 0.79 | 250 of 555 |
| Plunder the Graves (2017-06-09) | bad | 0.81 | 256 of 588 |
| Swell the Host (2015-11-13) | bad | 0.79 | 255 of 588 |

The audit of the bar (M-8). The bar asks the ladder to score a precon above its own broken copy. Rank is the precon's place among the holdout precons by score, 0 the weakest, and the column reads the mean rank of the pairs that win and of the pairs that lose. A break that hurts every precon alike moves the two means together. A lower mean rank among the losers says the bar fails on the weak precons, which is where a broken copy is most likely the better deck. Card rate and synergy read the signed move, the copy less the precon, over the pairs that lose. A positive card rate says the copy holds cards the top lists play more than the precon does.

| Axis | Own pairs | Lost | Mean rank, won | Mean rank, lost | `card_rate` move | `synergy` move | Copy graded above |
|---|---|---|---|---|---|---|---|
| lands | 189 | 0 | 18.58 | - | - | - | 0 of 0 |
| curve | 189 | 0 | 18.58 | - | - | - | 0 of 0 |
| colors | 178 | 0 | 18.89 | - | - | - | 0 of 0 |
| synergy | 228 | 35 | 24.97 | 14.37 | 0.00 | -0.28 | 12 of 35 |

Confusion on the holdout over the folds, rows are the label and columns the grade, worst first:

- bad: [4168 1338 395 6 0]
- baseline: [62 71 56 0 0]
- typical: [72 276 634 48 10]
- good: [248 76 234 1447 1995]
- great: [242 89 207 1354 2108]

The cards the top lists hold most: Sol Ring, Lotus Petal, Chrome Mox, Mana Vault, Mox Diamond, Mystic Remora, Arcane Signet, Mental Misstep, Mox Amber, Flusterstorm, Force of Will, Rhystic Study.

## The standard model

Tiers, worst first: bad, baseline, typical, good, great. Train counts: bad 630, baseline 7, good 3190, great 1673, typical 119. Holdout counts: bad 110, baseline 1, good 810, great 418, typical 21.

Cards with a rate: 842. Pairs that lift: 6588. Commanders with a signal: 0. Top lists shape: 23.2 lands at 2.50 over 1673 lists.

Features with no spread, dropped: game_changer, cedh_signal, high_bracket_share, commander_decks.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.107 | 0.059 | +1.452 |
| `unseen_share` | 0.026 | 0.098 | +0.091 |
| `synergy` | 1.564 | 0.791 | +0.944 |
| `land` | 22.934 | 2.461 | -0.097 |
| `avg_mana_value` | 2.579 | 0.699 | -0.807 |
| `color_sources` | 1.005 | 0.326 | +0.516 |
| `tapped_share` | 0.035 | 0.076 | -0.038 |
| `mana_turn_four` | 4.062 | 0.780 | +0.335 |
| `hands_two_to_four_lands` | 0.755 | 0.036 | +0.650 |
| `curve_low` | 0.627 | 0.181 | -0.319 |
| `curve_high` | 0.134 | 0.138 | +0.040 |
| `ramp` | 5.168 | 6.468 | -0.368 |
| `draw` | 5.033 | 3.531 | -0.056 |
| `removal` | 8.324 | 4.524 | +0.013 |
| `wipe` | 1.362 | 1.818 | +0.089 |
| `interaction` | 3.383 | 3.030 | +0.194 |
| `empty_roles` | 0.819 | 0.681 | +0.064 |
| `fast_mana` | 0.130 | 0.745 | +0.177 |
| `playset_share` | 0.595 | 0.192 | +0.016 |
| `singleton_share` | 0.096 | 0.110 | -0.488 |
| `source_spread` | 0.277 | 0.222 | +0.325 |

Thresholds: -4.223, -2.336, -0.410, 1.262. Loss: 1.088 over 3000 iterations. Defect detector accuracy on the holdout: 0.96, at a cut of 0.44.

The precon bar per broken axis, each precon over its own copies, with the cross pairs after it:

- lands: 1.00 of 8, cross 0.95 of 271
- curve: 1.00 of 8, cross 0.99 of 271
- colors: 1.00 of 4, cross 0.72 of 210
- copies: 1.00 of 3, cross 0.94 of 216
- synergy: 1.00 of 17, cross 0.76 of 387

The misses per axis over the folds (M-7). Both passed: the detector passed the precon and the copy, and the ladder put the copy at or above the precon. Precon flagged: the detector flagged the precon and passed the copy. Both flagged: it flagged both and read the precon as the more broken. Own copy: each precon against its own copy on the axis, with its misses by the same kinds. Moved: the mean absolute standardized delta between a precon and its own copy, for the corpus features and the three features that move most.

| Axis | Share | Both passed | Precon flagged | Both flagged | Own copy | Own misses | Corpus moved | Most moved |
|---|---|---|---|---|---|---|---|---|
| lands | 0.95 of 271 | 1 | 3 | 9 | 8 of 8 | 0, 0, 0 | `card_rate` 0.06, `synergy` 0.11, `unseen_share` 1.14 | `hands_two_to_four_lands` 3.95, `land` 3.21, `color_sources` 2.28 |
| curve | 0.99 of 271 | 0 | 0 | 4 | 8 of 8 | 0, 0, 0 | `card_rate` 0.19, `synergy` 0.19, `unseen_share` 1.76 | `avg_mana_value` 3.15, `curve_high` 3.05, `curve_low` 2.10 |
| colors | 0.72 of 210 | 35 | 12 | 12 | 4 of 4 | 0, 0, 0 | `card_rate` 0.00, `synergy` 0.00, `unseen_share` 0.00 | `tapped_share` 3.01, `color_sources` 1.66, `source_spread` 1.51 |
| copies | 0.94 of 216 | 4 | 0 | 8 | 3 of 3 | 0, 0, 0 | `card_rate` 0.29, `synergy` 0.04, `unseen_share` 0.46 | `singleton_share` 4.11, `playset_share` 2.35, `empty_roles` 0.98 |
| synergy | 0.76 of 387 | 36 | 42 | 16 | 17 of 17 | 0, 0, 0 | `card_rate` 0.12, `synergy` 0.20, `unseen_share` 3.62 | `unseen_share` 3.62, `color_sources` 2.12, `playset_share` 1.65 |

The precons that lose most over the sampled pairs of the bar:

| Precon | Grade | Detector | Misses |
|---|---|---|---|
| Donatello (2026-03-06) | bad | 0.55 | 43 of 160 |
| Leonardo (2026-03-06) | bad | 0.52 | 40 of 160 |
| Raphael (2026-03-06) | bad | 0.29 | 20 of 135 |
| Angels (2026-01-23) | bad | 0.22 | 31 of 210 |
| Michealangelo (2026-03-06) | bad | 0.39 | 23 of 160 |
| Eerie (2026-04-24) | baseline | 0.03 | 8 of 110 |
| Pirates (2026-01-23) | typical | 0.06 | 15 of 210 |
| Lifegain (2026-04-24) | typical | 0.01 | 2 of 210 |

| Axis | Own pairs | Lost | Mean rank, won | Mean rank, lost | `card_rate` move | `synergy` move | Copy graded above |
|---|---|---|---|---|---|---|---|
| lands | 8 | 0 | 0.75 | - | - | - | 0 of 0 |
| curve | 8 | 0 | 0.75 | - | - | - | 0 of 0 |
| colors | 4 | 0 | 0.75 | - | - | - | 0 of 0 |
| copies | 3 | 0 | 0.33 | - | - | - | 0 of 0 |
| synergy | 17 | 0 | 0.82 | - | - | - | 0 of 0 |

Confusion on the holdout over the folds, rows are the label and columns the grade, worst first:

- bad: [675 7 45 12 1]
- baseline: [5 1 2 0 0]
- typical: [42 11 42 28 17]
- good: [174 6 476 1631 1713]
- great: [98 1 253 757 982]

The cards the top lists hold most: Burst Lightning, Stock Up, Opt, Sleight of Hand, Llanowar Elves, Eddymurk Crab, Spell Pierce, Boomerang Basics, Stormchaser's Talent, Badgermole Cub, Flow State, Slickshot Show-Off.

## The modern model

Tiers, worst first: bad, baseline, typical, good, great. Train counts: bad 2495, baseline 266, good 3154, great 2486, typical 233. Holdout counts: bad 620, baseline 61, good 846, great 664, typical 63.

Cards with a rate: 907. Pairs that lift: 6964. Commanders with a signal: 0. Top lists shape: 21.7 lands at 2.45 over 2486 lists.

Features with no spread, dropped: cedh_signal, high_bracket_share, commander_decks.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.065 | 0.045 | +0.524 |
| `unseen_share` | 0.218 | 0.365 | -0.559 |
| `synergy` | 1.409 | 0.999 | +1.466 |
| `land` | 21.640 | 3.103 | -0.084 |
| `avg_mana_value` | 2.751 | 0.946 | -0.316 |
| `color_sources` | 0.927 | 0.322 | +0.335 |
| `tapped_share` | 0.125 | 0.141 | +0.082 |
| `mana_turn_four` | 3.843 | 0.800 | -0.016 |
| `hands_two_to_four_lands` | 0.731 | 0.061 | +0.341 |
| `curve_low` | 0.583 | 0.219 | -0.183 |
| `curve_high` | 0.181 | 0.159 | -0.314 |
| `ramp` | 5.396 | 6.781 | +0.112 |
| `draw` | 4.630 | 4.504 | +0.027 |
| `removal` | 7.817 | 4.559 | -0.185 |
| `wipe` | 1.021 | 1.761 | +0.010 |
| `interaction` | 3.082 | 3.880 | +0.037 |
| `empty_roles` | 1.020 | 0.837 | -0.050 |
| `fast_mana` | 0.581 | 1.527 | -0.009 |
| `game_changer` | 0.270 | 0.885 | -0.128 |
| `playset_share` | 0.575 | 0.235 | -0.559 |
| `singleton_share` | 0.144 | 0.161 | -0.728 |
| `source_spread` | 0.249 | 0.222 | -0.239 |

Thresholds: -3.707, -1.476, 0.825, 2.385. Loss: 1.061 over 3000 iterations. Defect detector accuracy on the holdout: 0.93, at a cut of 0.48.

The precon bar per broken axis, each precon over its own copies, with the cross pairs after it:

- lands: 1.00 of 327, cross 0.93 of 41250
- curve: 1.00 of 327, cross 0.98 of 41250
- colors: 0.97 of 75, cross 0.48 of 20632
- copies: 1.00 of 61, cross 0.86 of 23234
- synergy: 0.98 of 845, cross 0.84 of 74904

The misses per axis over the folds (M-7). Both passed: the detector passed the precon and the copy, and the ladder put the copy at or above the precon. Precon flagged: the detector flagged the precon and passed the copy. Both flagged: it flagged both and read the precon as the more broken. Own copy: each precon against its own copy on the axis, with its misses by the same kinds. Moved: the mean absolute standardized delta between a precon and its own copy, for the corpus features and the three features that move most.

| Axis | Share | Both passed | Precon flagged | Both flagged | Own copy | Own misses | Corpus moved | Most moved |
|---|---|---|---|---|---|---|---|---|
| lands | 0.93 of 41250 | 1719 | 518 | 810 | 327 of 327 | 0, 0, 0 | `card_rate` 0.01, `synergy` 0.00, `unseen_share` 0.09 | `land` 2.54, `hands_two_to_four_lands` 2.40, `color_sources` 1.50 |
| curve | 0.98 of 41250 | 317 | 162 | 507 | 327 of 327 | 0, 0, 0 | `card_rate` 0.03, `synergy` 0.00, `unseen_share` 0.20 | `curve_high` 2.67, `avg_mana_value` 2.23, `curve_low` 1.72 |
| colors | 0.48 of 20632 | 7806 | 2302 | 541 | 73 of 75 | 2, 0, 0 | `card_rate` 0.00, `synergy` 0.00, `unseen_share` 0.00 | `tapped_share` 1.25, `color_sources` 1.12, `source_spread` 0.95 |
| copies | 0.86 of 23234 | 1994 | 694 | 449 | 61 of 61 | 0, 0, 0 | `card_rate` 0.04, `synergy` 0.00, `unseen_share` 0.20 | `singleton_share` 3.49, `playset_share` 2.40, `empty_roles` 0.76 |
| synergy | 0.84 of 74904 | 6788 | 3097 | 2364 | 825 of 845 | 14, 0, 6 | `card_rate` 0.03, `synergy` 0.00, `unseen_share` 0.21 | `playset_share` 1.32, `empty_roles` 0.63, `removal` 0.61 |

The precons that lose most over the sampled pairs of the bar:

| Precon | Grade | Detector | Misses |
|---|---|---|---|
| Sultai Schemers (2014-09-26) | bad | 0.89 | 230 of 344 |
| Thallids (2005-08-01) | bad | 0.85 | 141 of 240 |
| Jeskai Monks (2014-09-26) | bad | 0.84 | 195 of 333 |
| Beasts (2005-08-01) | bad | 0.87 | 130 of 240 |
| Grave Advantage (2015-01-23) | bad | 0.77 | 116 of 240 |
| Stampede (2004-06-04) | bad | 0.78 | 147 of 327 |
| Illusionary Might (2011-08-12) | bad | 0.79 | 104 of 240 |
| Spiraling Doom (2012-02-24) | bad | 0.79 | 129 of 307 |
| Second Sun Control (2018-04-06) | bad | 0.71 | 135 of 333 |
| Stampede of Beasts (2010-07-16) | bad | 0.73 | 135 of 333 |

| Axis | Own pairs | Lost | Mean rank, won | Mean rank, lost | `card_rate` move | `synergy` move | Copy graded above |
|---|---|---|---|---|---|---|---|
| lands | 327 | 0 | 32.83 | - | - | - | 0 of 0 |
| curve | 327 | 0 | 32.83 | - | - | - | 0 of 0 |
| colors | 75 | 2 | 37.62 | 28.50 | 0.00 | 0.00 | 0 of 2 |
| copies | 61 | 0 | 38.05 | - | - | - | 0 of 0 |
| synergy | 845 | 20 | 32.08 | 30.85 | 0.05 | 0.00 | 3 of 20 |

Confusion on the holdout over the folds, rows are the label and columns the grade, worst first:

- bad: [2902 70 139 1 3]
- baseline: [207 113 7 0 0]
- typical: [46 3 99 97 51]
- good: [184 0 816 1483 1517]
- great: [105 0 638 916 1491]

The cards the top lists hold most: Solitude, Quantum Riddler, Thoughtseize, Force of Negation, Prismatic Ending, Galvanic Discharge, Fatal Push, Mishra's Bauble, Kozilek's Command, Ephemerate, Ragavan, Nimble Pilferer, Malevolent Rumble.

## Run

- Suite `quality`, run `pr14b-quality-gate-run23`, on 2026-09-14, commit `b7aa42d`.
- Roles: none, no provider call.
- Versions: card snapshot 2026-09-04, commanders_day `2026-09-14`, quality_model `20260914T184742Z`.
- Calls: 0. Cost: unpriced. Time: 117 seconds.

## Verdict

Verdict: PASS. Time: 117 seconds, no provider call.
