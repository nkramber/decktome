# PR-14B quality gate

Run date: 2026-09-10. Card snapshot: 2026-09-04. Model: 20260910T012433Z.

## Separation on the holdout

The bars read 5 folds, every list holdout once, and the pairs of a bar sample evenly under the cap of 20000 (M-7). The precon bar reads each precon against its own broken copies (D-573). The cross pairs, every precon against every copy, stand as information.

| Format | Lists | Used | Synthetic | Holdout | Great over precon | Precon over own copy | Precon over bad, cross | Accuracy |
|---|---|---|---|---|---|---|---|---|
| commander | 23628 | 9229 | 6145 | 15374 | 0.96 of 97450 (bar 0.90) | 0.88 of 945 (bar 0.95) | 0.86 of 99933 | 0.65 |
| standard | 8228 | 5768 | 370 | 6138 | 0.91 of 2750 (bar 0.90) | 1.00 of 40 (bar 0.95) | 0.91 of 730 | 0.46 |
| modern | 20292 | 7099 | 2745 | 9844 | 0.98 of 98691 (bar 0.90) | 0.99 of 1635 (bar 0.95) | 0.86 of 99754 | 0.55 |

## The bracket 5 offer

Commander reads of 2026-09-07: 1562 rows, 945 with Topdeck.gg entries.

| Commander | cEDH signal | Top cuts | Entries |
|---|---|---|---|
| The Jolly Balloon Man | 0.57 | 5 | 6 |
| Niv-Mizzet, Parun | 0.50 | 5 | 31 |
| Etali, Primal Conqueror // Etali, Primal Sickness | 0.50 | 76 | 411 |

## The commander model

Tiers, worst first: bad, baseline, typical, good, great. Train counts: bad 4995, baseline 145, good 3223, great 3160, typical 854. Holdout counts: bad 1150, baseline 44, good 777, great 840, typical 186.

Cards with a rate: 9942. Pairs that lift: 200000. Commanders with a signal: 1552. Top lists shape: 27.6 lands at 2.19 over 3160 lists.

Features with no spread, dropped: playset_share, singleton_share.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.152 | 0.139 | +0.968 |
| `unseen_share` | 0.073 | 0.103 | -0.233 |
| `synergy` | 0.490 | 0.535 | +0.167 |
| `land` | 30.795 | 5.609 | -0.256 |
| `avg_mana_value` | 2.802 | 0.859 | -0.445 |
| `color_sources` | 1.256 | 0.353 | -0.004 |
| `tapped_share` | 0.073 | 0.084 | +0.345 |
| `mana_turn_four` | 4.704 | 0.614 | +0.755 |
| `hands_two_to_four_lands` | 0.653 | 0.087 | +0.524 |
| `curve_low` | 0.505 | 0.188 | +0.188 |
| `curve_high` | 0.153 | 0.133 | -0.175 |
| `ramp` | 17.545 | 7.595 | +0.064 |
| `draw` | 7.425 | 4.073 | -0.227 |
| `removal` | 6.407 | 3.637 | +0.087 |
| `wipe` | 1.791 | 1.977 | -0.146 |
| `interaction` | 7.389 | 4.569 | +0.196 |
| `empty_roles` | 0.049 | 0.220 | +0.002 |
| `fast_mana` | 5.144 | 4.818 | +0.266 |
| `game_changer` | 7.236 | 7.396 | +0.119 |
| `source_spread` | 0.095 | 0.122 | -0.472 |
| `cedh_signal` | 0.278 | 0.200 | +0.219 |
| `high_bracket_share` | 0.215 | 0.250 | +0.091 |
| `commander_decks` | 5.867 | 4.128 | -0.129 |

Thresholds: -3.139, -1.434, 0.695, 3.021. Loss: 1.062 over 3000 iterations. Defect detector accuracy on the holdout: 0.92, at a cut of 0.46.

The precon bar per broken axis, each precon over its own copies, with the cross pairs after it:

- lands: 1.00 of 189, cross 0.96 of 46364
- curve: 1.00 of 189, cross 0.95 of 46364
- colors: 1.00 of 178, cross 0.91 of 27525
- synergy: 0.70 of 389, cross 0.77 of 99355

The misses per axis over the folds (M-7). Both passed: the detector passed the precon and the copy, and the ladder put the copy at or above the precon. Precon flagged: the detector flagged the precon and passed the copy. Both flagged: it flagged both and read the precon as the more broken. Own copy: each precon against its own copy on the axis, with its misses by the same kinds. Moved: the mean absolute standardized delta between a precon and its own copy, for the corpus features and the three features that move most.

| Axis | Share | Both passed | Precon flagged | Both flagged | Own copy | Own misses | Corpus moved | Most moved |
|---|---|---|---|---|---|---|---|---|
| lands | 0.96 of 46364 | 580 | 437 | 737 | 189 of 189 | 0, 0, 0 | `card_rate` 0.03, `synergy` 0.39, `unseen_share` 0.33 | `land` 2.20, `hands_two_to_four_lands` 2.06, `color_sources` 1.06 |
| curve | 0.95 of 46364 | 259 | 559 | 1607 | 189 of 189 | 0, 0, 0 | `card_rate` 0.18, `synergy` 0.50, `unseen_share` 0.52 | `curve_high` 1.87, `avg_mana_value` 1.48, `curve_low` 1.24 |
| colors | 0.91 of 27525 | 796 | 849 | 934 | 178 of 178 | 0, 0, 0 | `card_rate` 0.00, `synergy` 0.00, `unseen_share` 0.00 | `source_spread` 3.74, `tapped_share` 2.73, `color_sources` 1.73 |
| synergy | 0.77 of 99355 | 6068 | 7656 | 9478 | 271 of 389 | 54, 24, 40 | `card_rate` 0.08, `synergy` 0.91, `unseen_share` 0.93 | `unseen_share` 0.93, `synergy` 0.91, `wipe` 0.73 |

The precons that lose most over the sampled pairs of the bar:

| Precon | Grade | Detector | Misses |
|---|---|---|---|
| Eternal Bargain (2013-11-01) | bad | 0.92 | 344 of 454 |
| Political Puppets (2011-06-17) | bad | 0.90 | 351 of 500 |
| Peer Through Time (2014-11-07) | bad | 0.87 | 308 of 500 |
| Heavenly Inferno (2011-06-17) | bad | 0.86 | 332 of 588 |
| Heavenly Inferno (2017-06-09) | bad | 0.82 | 243 of 454 |
| Evasive Maneuvers (2013-11-01) | bad | 0.81 | 255 of 500 |
| Swell the Host (2015-11-13) | bad | 0.81 | 290 of 588 |
| Evasive Maneuvers (2017-06-09) | bad | 0.78 | 279 of 571 |
| Nature of the Beast (2013-11-01) | bad | 0.79 | 244 of 500 |
| Mirror Mastery (2011-06-17) | bad | 0.77 | 276 of 571 |

Confusion on the holdout over the folds, rows are the label and columns the grade, worst first:

- bad: [5971 10 147 17 0]
- baseline: [124 9 55 1 0]
- typical: [410 21 549 50 10]
- good: [250 11 266 1399 2074]
- great: [208 11 252 1391 2138]

The cards the top lists hold most: Sol Ring, Lotus Petal, Chrome Mox, Mana Vault, Mox Diamond, Arcane Signet, Mystic Remora, Mental Misstep, Flusterstorm, Rhystic Study, Force of Will, Mox Amber.

## The standard model

Tiers, worst first: bad, baseline, typical, good, great. Train counts: bad 320, baseline 7, good 3199, great 1358, typical 57. Holdout counts: bad 50, baseline 1, good 801, great 336, typical 9.

Cards with a rate: 815. Pairs that lift: 6032. Commanders with a signal: 0. Top lists shape: 23.1 lands at 2.49 over 1358 lists.

Features with no spread, dropped: game_changer, cedh_signal, high_bracket_share, commander_decks.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.117 | 0.061 | +1.389 |
| `unseen_share` | 0.016 | 0.081 | +0.114 |
| `synergy` | 1.590 | 0.737 | +0.982 |
| `land` | 22.999 | 2.396 | -0.079 |
| `avg_mana_value` | 2.544 | 0.659 | -0.951 |
| `color_sources` | 1.065 | 0.413 | +0.344 |
| `tapped_share` | 0.033 | 0.072 | -0.064 |
| `mana_turn_four` | 4.062 | 0.764 | +0.430 |
| `hands_two_to_four_lands` | 0.756 | 0.032 | +0.535 |
| `curve_low` | 0.635 | 0.175 | -0.372 |
| `curve_high` | 0.130 | 0.130 | +0.137 |
| `ramp` | 5.030 | 6.428 | -0.358 |
| `draw` | 4.892 | 3.441 | -0.068 |
| `removal` | 8.389 | 4.452 | +0.136 |
| `wipe` | 1.353 | 1.782 | +0.071 |
| `interaction` | 3.390 | 2.968 | +0.088 |
| `empty_roles` | 0.822 | 0.684 | -0.050 |
| `fast_mana` | 0.150 | 0.802 | +0.091 |
| `playset_share` | 0.600 | 0.182 | -0.038 |
| `singleton_share` | 0.088 | 0.088 | -0.484 |
| `source_spread` | 0.280 | 0.215 | +0.309 |

Thresholds: -4.453, -2.545, -0.649, 1.012. Loss: 1.086 over 3000 iterations. Defect detector accuracy on the holdout: 0.96, at a cut of 0.36.

The precon bar per broken axis, each precon over its own copies, with the cross pairs after it:

- lands: 1.00 of 8, cross 0.98 of 146
- curve: 1.00 of 8, cross 1.00 of 146
- colors: 1.00 of 4, cross 0.70 of 105
- copies: 1.00 of 3, cross 1.00 of 111
- synergy: 1.00 of 17, cross 0.84 of 222

The misses per axis over the folds (M-7). Both passed: the detector passed the precon and the copy, and the ladder put the copy at or above the precon. Precon flagged: the detector flagged the precon and passed the copy. Both flagged: it flagged both and read the precon as the more broken. Own copy: each precon against its own copy on the axis, with its misses by the same kinds. Moved: the mean absolute standardized delta between a precon and its own copy, for the corpus features and the three features that move most.

| Axis | Share | Both passed | Precon flagged | Both flagged | Own copy | Own misses | Corpus moved | Most moved |
|---|---|---|---|---|---|---|---|---|
| lands | 0.98 of 146 | 3 | 0 | 0 | 8 of 8 | 0, 0, 0 | `card_rate` 0.06, `synergy` 0.13, `unseen_share` 1.31 | `hands_two_to_four_lands` 4.19, `land` 3.28, `color_sources` 1.69 |
| curve | 1.00 of 146 | 0 | 0 | 0 | 8 of 8 | 0, 0, 0 | `card_rate` 0.22, `synergy` 0.22, `unseen_share` 2.78 | `avg_mana_value` 3.62, `curve_high` 3.25, `unseen_share` 2.78 |
| colors | 0.70 of 105 | 31 | 0 | 0 | 4 of 4 | 0, 0, 0 | `card_rate` 0.00, `synergy` 0.00, `unseen_share` 0.00 | `tapped_share` 3.12, `source_spread` 1.56, `color_sources` 1.29 |
| copies | 1.00 of 111 | 0 | 0 | 0 | 3 of 3 | 0, 0, 0 | `card_rate` 0.30, `synergy` 0.03, `unseen_share` 0.77 | `singleton_share` 5.00, `playset_share` 2.47, `empty_roles` 0.98 |
| synergy | 0.84 of 222 | 35 | 0 | 0 | 17 of 17 | 0, 0, 0 | `card_rate` 0.13, `synergy` 0.21, `unseen_share` 4.40 | `unseen_share` 4.40, `playset_share` 1.82, `color_sources` 1.64 |

The precons that lose most over the sampled pairs of the bar:

| Precon | Grade | Detector | Misses |
|---|---|---|---|
| Angels (2026-01-23) | baseline | 0.18 | 15 of 105 |
| Donatello (2026-03-06) | bad | 0.35 | 13 of 100 |
| Leonardo (2026-03-06) | bad | 0.35 | 13 of 100 |
| Michealangelo (2026-03-06) | bad | 0.26 | 12 of 100 |
| Raphael (2026-03-06) | bad | 0.23 | 7 of 65 |
| Pirates (2026-01-23) | typical | 0.07 | 6 of 105 |
| Eerie (2026-04-24) | baseline | 0.05 | 1 of 50 |
| Lifegain (2026-04-24) | good | 0.01 | 2 of 105 |

Confusion on the holdout over the folds, rows are the label and columns the grade, worst first:

- bad: [339 1 16 12 2]
- baseline: [4 2 1 1 0]
- typical: [20 3 18 17 8]
- good: [233 1 462 1676 1628]
- great: [108 3 167 620 796]

The cards the top lists hold most: Burst Lightning, Stock Up, Opt, Sleight of Hand, Llanowar Elves, Eddymurk Crab, Spell Pierce, Boomerang Basics, Stormchaser's Talent, Flow State, Badgermole Cub, Slickshot Show-Off.

## The modern model

Tiers, worst first: bad, baseline, typical, good, great. Train counts: bad 2205, baseline 266, good 3149, great 2021, typical 175. Holdout counts: bad 540, baseline 61, good 851, great 529, typical 47.

Cards with a rate: 891. Pairs that lift: 6786. Commanders with a signal: 0. Top lists shape: 21.7 lands at 2.44 over 2021 lists.

Features with no spread, dropped: cedh_signal, high_bracket_share, commander_decks.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.065 | 0.046 | +0.464 |
| `unseen_share` | 0.228 | 0.377 | -0.742 |
| `synergy` | 1.404 | 1.010 | +1.367 |
| `land` | 21.679 | 3.076 | -0.040 |
| `avg_mana_value` | 2.753 | 0.945 | -0.367 |
| `color_sources` | 0.963 | 0.327 | +0.326 |
| `tapped_share` | 0.124 | 0.141 | +0.124 |
| `mana_turn_four` | 3.844 | 0.797 | +0.014 |
| `hands_two_to_four_lands` | 0.732 | 0.059 | +0.260 |
| `curve_low` | 0.580 | 0.221 | -0.065 |
| `curve_high` | 0.180 | 0.158 | -0.186 |
| `ramp` | 5.348 | 6.806 | +0.071 |
| `draw` | 4.663 | 4.601 | +0.030 |
| `removal` | 7.758 | 4.572 | -0.204 |
| `wipe` | 1.008 | 1.746 | +0.011 |
| `interaction` | 3.030 | 3.800 | +0.010 |
| `empty_roles` | 1.031 | 0.844 | +0.013 |
| `fast_mana` | 0.553 | 1.506 | -0.004 |
| `game_changer` | 0.271 | 0.885 | -0.116 |
| `playset_share` | 0.569 | 0.236 | -0.552 |
| `singleton_share` | 0.145 | 0.157 | -0.652 |
| `source_spread` | 0.245 | 0.218 | -0.241 |

Thresholds: -3.808, -1.518, 0.849, 2.402. Loss: 1.056 over 3000 iterations. Defect detector accuracy on the holdout: 0.93, at a cut of 0.48.

The precon bar per broken axis, each precon over its own copies, with the cross pairs after it:

- lands: 1.00 of 327, cross 0.94 of 36476
- curve: 1.00 of 327, cross 0.98 of 36476
- colors: 0.97 of 75, cross 0.46 of 17066
- copies: 1.00 of 61, cross 0.83 of 18460
- synergy: 0.98 of 845, cross 0.85 of 70333

The misses per axis over the folds (M-7). Both passed: the detector passed the precon and the copy, and the ladder put the copy at or above the precon. Precon flagged: the detector flagged the precon and passed the copy. Both flagged: it flagged both and read the precon as the more broken. Own copy: each precon against its own copy on the axis, with its misses by the same kinds. Moved: the mean absolute standardized delta between a precon and its own copy, for the corpus features and the three features that move most.

| Axis | Share | Both passed | Precon flagged | Both flagged | Own copy | Own misses | Corpus moved | Most moved |
|---|---|---|---|---|---|---|---|---|
| lands | 0.94 of 36476 | 1330 | 347 | 613 | 327 of 327 | 0, 0, 0 | `card_rate` 0.01, `synergy` 0.00, `unseen_share` 0.09 | `land` 2.57, `hands_two_to_four_lands` 2.47, `color_sources` 1.46 |
| curve | 0.98 of 36476 | 350 | 142 | 349 | 327 of 327 | 0, 0, 0 | `card_rate` 0.04, `synergy` 0.00, `unseen_share` 0.20 | `curve_high` 2.68, `avg_mana_value` 2.22, `curve_low` 1.70 |
| colors | 0.46 of 17066 | 7007 | 1784 | 358 | 73 of 75 | 2, 0, 0 | `card_rate` 0.00, `synergy` 0.00, `unseen_share` 0.00 | `tapped_share` 1.25, `color_sources` 1.03, `source_spread` 0.96 |
| copies | 0.83 of 18460 | 2257 | 615 | 315 | 61 of 61 | 0, 0, 0 | `card_rate` 0.04, `synergy` 0.00, `unseen_share` 0.19 | `singleton_share` 3.56, `playset_share` 2.38, `empty_roles` 0.91 |
| synergy | 0.85 of 70333 | 5985 | 2411 | 2105 | 832 of 845 | 10, 0, 3 | `card_rate` 0.03, `synergy` 0.00, `unseen_share` 0.21 | `playset_share` 1.31, `empty_roles` 0.64, `removal` 0.61 |

The precons that lose most over the sampled pairs of the bar:

| Precon | Grade | Detector | Misses |
|---|---|---|---|
| Sultai Schemers (2014-09-26) | bad | 0.87 | 233 of 344 |
| Illusionary Might (2011-08-12) | bad | 0.84 | 128 of 240 |
| Beasts (2005-08-01) | bad | 0.85 | 126 of 240 |
| Jeskai Monks (2014-09-26) | bad | 0.80 | 166 of 333 |
| Second Sun Control (2018-04-06) | bad | 0.78 | 158 of 333 |
| Thallids (2005-08-01) | bad | 0.82 | 110 of 240 |
| Stampede (2004-06-04) | bad | 0.77 | 146 of 327 |
| Wizards (2005-08-01) | bad | 0.81 | 134 of 307 |
| Grave Advantage (2015-01-23) | bad | 0.75 | 98 of 240 |
| Stampede of Beasts (2010-07-16) | bad | 0.72 | 131 of 333 |

Confusion on the holdout over the folds, rows are the label and columns the grade, worst first:

- bad: [2554 58 130 0 3]
- baseline: [207 115 5 0 0]
- typical: [34 4 74 70 40]
- good: [175 0 783 1503 1539]
- great: [83 0 522 739 1206]

The cards the top lists hold most: Solitude, Quantum Riddler, Thoughtseize, Force of Negation, Prismatic Ending, Fatal Push, Ephemerate, Mishra's Bauble, Galvanic Discharge, Kozilek's Command, Malevolent Rumble, Ragavan, Nimble Pilferer.

## Run

- Suite `quality`, run `pr14b-quality-gate-run15-control`, on 2026-09-10, commit `d7f7621`.
- Roles: none, no provider call.
- Versions: card snapshot 2026-09-04, commanders_day `2026-09-07`, quality_model `20260910T012433Z`.
- Calls: 0. Cost: unpriced. Time: 108 seconds.

## Verdict

Verdict: FAIL. Time: 108 seconds, no provider call. Failed bars: commander: precon over own copy 0.88 of 945 (bar 0.95).
