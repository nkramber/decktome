# PR-14B quality gate

Run date: 2026-09-10. Card snapshot: 2026-09-04. Model: 20260910T012206Z.

## Separation on the holdout

The bars read 5 folds, every list holdout once, and the pairs of a bar sample evenly under the cap of 20000 (M-7). The precon bar reads each precon against its own broken copies (D-573). The cross pairs, every precon against every copy, stand as information.

| Format | Lists | Used | Synthetic | Holdout | Great over precon | Precon over own copy | Precon over bad, cross | Accuracy |
|---|---|---|---|---|---|---|---|---|
| commander | 23628 | 9229 | 6145 | 15374 | 0.95 of 97450 (bar 0.90) | 0.92 of 945 (bar 0.95) | 0.91 of 99933 | 0.66 |
| standard | 8228 | 5768 | 370 | 6138 | 0.91 of 2750 (bar 0.90) | 1.00 of 40 (bar 0.95) | 0.92 of 730 | 0.44 |
| modern | 20292 | 7099 | 2745 | 9844 | 0.99 of 98691 (bar 0.90) | 0.99 of 1635 (bar 0.95) | 0.85 of 99754 | 0.55 |

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
| `card_rate` | 0.152 | 0.139 | +0.984 |
| `unseen_share` | 0.073 | 0.103 | -0.158 |
| `synergy` | 0.490 | 0.535 | +0.178 |
| `casual_rate` | 0.049 | 0.019 | +0.224 |
| `casual_synergy` | 0.365 | 0.257 | +0.074 |
| `land` | 30.795 | 5.609 | -0.258 |
| `avg_mana_value` | 2.802 | 0.859 | -0.484 |
| `color_sources` | 1.256 | 0.353 | -0.001 |
| `tapped_share` | 0.073 | 0.084 | +0.374 |
| `mana_turn_four` | 4.704 | 0.614 | +0.726 |
| `hands_two_to_four_lands` | 0.653 | 0.087 | +0.441 |
| `curve_low` | 0.505 | 0.188 | +0.166 |
| `curve_high` | 0.153 | 0.133 | -0.125 |
| `ramp` | 17.545 | 7.595 | -0.025 |
| `draw` | 7.425 | 4.073 | -0.267 |
| `removal` | 6.407 | 3.637 | +0.078 |
| `wipe` | 1.791 | 1.977 | -0.184 |
| `interaction` | 7.389 | 4.569 | +0.158 |
| `empty_roles` | 0.049 | 0.220 | +0.009 |
| `fast_mana` | 5.144 | 4.818 | +0.286 |
| `game_changer` | 7.236 | 7.396 | +0.150 |
| `source_spread` | 0.095 | 0.122 | -0.474 |
| `cedh_signal` | 0.278 | 0.200 | +0.306 |
| `high_bracket_share` | 0.215 | 0.250 | +0.131 |
| `commander_decks` | 5.867 | 4.128 | -0.200 |

Thresholds: -3.096, -1.306, 0.824, 3.076. Loss: 1.054 over 3000 iterations. Defect detector accuracy on the holdout: 0.94, at a cut of 0.42.

The precon bar per broken axis, each precon over its own copies, with the cross pairs after it:

- lands: 1.00 of 189, cross 0.93 of 46364
- curve: 1.00 of 189, cross 0.97 of 46364
- colors: 1.00 of 178, cross 0.89 of 27525
- synergy: 0.81 of 389, cross 0.89 of 99355

The misses per axis over the folds (M-7). Both passed: the detector passed the precon and the copy, and the ladder put the copy at or above the precon. Precon flagged: the detector flagged the precon and passed the copy. Both flagged: it flagged both and read the precon as the more broken. Own copy: each precon against its own copy on the axis, with its misses by the same kinds. Moved: the mean absolute standardized delta between a precon and its own copy, for the corpus features and the three features that move most.

| Axis | Share | Both passed | Precon flagged | Both flagged | Own copy | Own misses | Corpus moved | Most moved |
|---|---|---|---|---|---|---|---|---|
| lands | 0.93 of 46364 | 977 | 648 | 1486 | 189 of 189 | 0, 0, 0 | `card_rate` 0.03, `synergy` 0.39, `unseen_share` 0.33 | `land` 2.20, `hands_two_to_four_lands` 2.06, `color_sources` 1.06 |
| curve | 0.97 of 46364 | 321 | 340 | 936 | 189 of 189 | 0, 0, 0 | `card_rate` 0.18, `synergy` 0.50, `unseen_share` 0.52 | `curve_high` 1.87, `casual_rate` 1.75, `avg_mana_value` 1.48 |
| colors | 0.89 of 27525 | 1122 | 873 | 1157 | 178 of 178 | 0, 0, 0 | `card_rate` 0.00, `synergy` 0.00, `unseen_share` 0.00 | `source_spread` 3.74, `tapped_share` 2.73, `color_sources` 1.73 |
| synergy | 0.89 of 99355 | 3344 | 2978 | 5037 | 315 of 389 | 40, 6, 28 | `card_rate` 0.08, `synergy` 0.91, `unseen_share` 0.93 | `casual_rate` 1.03, `unseen_share` 0.93, `synergy` 0.91 |

The precons that lose most over the sampled pairs of the bar:

| Precon | Grade | Detector | Misses |
|---|---|---|---|
| Eternal Bargain (2013-11-01) | bad | 0.90 | 240 of 454 |
| Mirror Mastery (2011-06-17) | bad | 0.87 | 266 of 571 |
| Political Puppets (2011-06-17) | bad | 0.86 | 221 of 500 |
| Peer Through Time (2014-11-07) | bad | 0.86 | 200 of 500 |
| Heavenly Inferno (2011-06-17) | bad | 0.87 | 222 of 588 |
| Swell the Host (2015-11-13) | bad | 0.82 | 210 of 588 |
| Heavenly Inferno (2017-06-09) | bad | 0.84 | 162 of 454 |
| Deathdancer Xira (2009-08-26) | bad | 0.81 | 191 of 571 |
| Undead Unleashed (2021-09-24) | bad | 0.79 | 188 of 588 |
| Elven Empire (2021-02-05) | bad | 0.80 | 176 of 555 |

Confusion on the holdout over the folds, rows are the label and columns the grade, worst first:

- bad: [5999 8 118 20 0]
- baseline: [122 17 50 0 0]
- typical: [250 90 625 61 14]
- good: [320 29 230 1373 2048]
- great: [261 27 236 1402 2074]

The cards the top lists hold most: Sol Ring, Lotus Petal, Chrome Mox, Mana Vault, Mox Diamond, Arcane Signet, Mystic Remora, Mental Misstep, Flusterstorm, Rhystic Study, Force of Will, Mox Amber.

## The standard model

Tiers, worst first: bad, baseline, typical, good, great. Train counts: bad 320, baseline 7, good 3199, great 1358, typical 57. Holdout counts: bad 50, baseline 1, good 801, great 336, typical 9.

Cards with a rate: 815. Pairs that lift: 6032. Commanders with a signal: 0. Top lists shape: 23.1 lands at 2.49 over 1358 lists.

Features with no spread, dropped: game_changer, cedh_signal, high_bracket_share, commander_decks.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.117 | 0.061 | +1.707 |
| `unseen_share` | 0.016 | 0.081 | +0.074 |
| `synergy` | 1.590 | 0.737 | +0.980 |
| `casual_rate` | 0.055 | 0.030 | -0.411 |
| `casual_synergy` | 0.292 | 0.415 | -0.335 |
| `land` | 22.999 | 2.396 | -0.069 |
| `avg_mana_value` | 2.544 | 0.659 | -0.741 |
| `color_sources` | 1.065 | 0.413 | +0.584 |
| `tapped_share` | 0.033 | 0.072 | -0.029 |
| `mana_turn_four` | 4.062 | 0.764 | +0.508 |
| `hands_two_to_four_lands` | 0.756 | 0.032 | +0.479 |
| `curve_low` | 0.635 | 0.175 | -0.288 |
| `curve_high` | 0.130 | 0.130 | +0.071 |
| `ramp` | 5.030 | 6.428 | -0.215 |
| `draw` | 4.892 | 3.441 | +0.073 |
| `removal` | 8.389 | 4.452 | +0.195 |
| `wipe` | 1.353 | 1.782 | +0.057 |
| `interaction` | 3.390 | 2.968 | +0.159 |
| `empty_roles` | 0.822 | 0.684 | +0.111 |
| `fast_mana` | 0.150 | 0.802 | +0.056 |
| `playset_share` | 0.600 | 0.182 | -0.106 |
| `singleton_share` | 0.088 | 0.088 | -0.534 |
| `source_spread` | 0.280 | 0.215 | +0.348 |

Thresholds: -4.608, -2.704, -0.783, 0.984. Loss: 1.058 over 3000 iterations. Defect detector accuracy on the holdout: 0.96, at a cut of 0.36.

The precon bar per broken axis, each precon over its own copies, with the cross pairs after it:

- lands: 1.00 of 8, cross 0.97 of 146
- curve: 1.00 of 8, cross 1.00 of 146
- colors: 1.00 of 4, cross 0.76 of 105
- copies: 1.00 of 3, cross 1.00 of 111
- synergy: 1.00 of 17, cross 0.89 of 222

The misses per axis over the folds (M-7). Both passed: the detector passed the precon and the copy, and the ladder put the copy at or above the precon. Precon flagged: the detector flagged the precon and passed the copy. Both flagged: it flagged both and read the precon as the more broken. Own copy: each precon against its own copy on the axis, with its misses by the same kinds. Moved: the mean absolute standardized delta between a precon and its own copy, for the corpus features and the three features that move most.

| Axis | Share | Both passed | Precon flagged | Both flagged | Own copy | Own misses | Corpus moved | Most moved |
|---|---|---|---|---|---|---|---|---|
| lands | 0.97 of 146 | 5 | 0 | 0 | 8 of 8 | 0, 0, 0 | `card_rate` 0.06, `synergy` 0.13, `unseen_share` 1.31 | `hands_two_to_four_lands` 4.19, `land` 3.28, `color_sources` 1.69 |
| curve | 1.00 of 146 | 0 | 0 | 0 | 8 of 8 | 0, 0, 0 | `card_rate` 0.22, `synergy` 0.22, `unseen_share` 2.78 | `avg_mana_value` 3.62, `curve_high` 3.25, `unseen_share` 2.78 |
| colors | 0.76 of 105 | 25 | 0 | 0 | 4 of 4 | 0, 0, 0 | `card_rate` 0.00, `synergy` 0.00, `unseen_share` 0.00 | `tapped_share` 3.12, `source_spread` 1.56, `color_sources` 1.29 |
| copies | 1.00 of 111 | 0 | 0 | 0 | 3 of 3 | 0, 0, 0 | `card_rate` 0.30, `synergy` 0.03, `unseen_share` 0.77 | `singleton_share` 5.00, `playset_share` 2.47, `empty_roles` 0.98 |
| synergy | 0.89 of 222 | 25 | 0 | 0 | 17 of 17 | 0, 0, 0 | `card_rate` 0.13, `synergy` 0.21, `unseen_share` 4.40 | `unseen_share` 4.40, `playset_share` 1.82, `color_sources` 1.64 |

The precons that lose most over the sampled pairs of the bar:

| Precon | Grade | Detector | Misses |
|---|---|---|---|
| Raphael (2026-03-06) | bad | 0.19 | 9 of 65 |
| Angels (2026-01-23) | typical | 0.18 | 13 of 105 |
| Leonardo (2026-03-06) | bad | 0.31 | 9 of 100 |
| Donatello (2026-03-06) | bad | 0.32 | 8 of 100 |
| Michealangelo (2026-03-06) | bad | 0.23 | 8 of 100 |
| Pirates (2026-01-23) | typical | 0.07 | 6 of 105 |
| Eerie (2026-04-24) | baseline | 0.03 | 1 of 50 |
| Lifegain (2026-04-24) | good | 0.01 | 1 of 105 |

Confusion on the holdout over the folds, rows are the label and columns the grade, worst first:

- bad: [338 0 20 11 1]
- baseline: [4 1 2 1 0]
- typical: [23 4 17 16 6]
- good: [192 1 569 1560 1678]
- great: [99 3 165 612 815]

The cards the top lists hold most: Burst Lightning, Stock Up, Opt, Sleight of Hand, Llanowar Elves, Eddymurk Crab, Spell Pierce, Boomerang Basics, Stormchaser's Talent, Flow State, Badgermole Cub, Slickshot Show-Off.

## The modern model

Tiers, worst first: bad, baseline, typical, good, great. Train counts: bad 2205, baseline 266, good 3149, great 2021, typical 175. Holdout counts: bad 540, baseline 61, good 851, great 529, typical 47.

Cards with a rate: 891. Pairs that lift: 6786. Commanders with a signal: 0. Top lists shape: 21.7 lands at 2.44 over 2021 lists.

Features with no spread, dropped: cedh_signal, high_bracket_share, commander_decks.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.065 | 0.046 | +1.111 |
| `unseen_share` | 0.228 | 0.377 | -0.770 |
| `synergy` | 1.404 | 1.010 | +1.367 |
| `casual_rate` | 0.026 | 0.022 | -0.558 |
| `casual_synergy` | 0.753 | 0.850 | -0.306 |
| `land` | 21.679 | 3.076 | +0.089 |
| `avg_mana_value` | 2.753 | 0.945 | -0.450 |
| `color_sources` | 0.963 | 0.327 | +0.226 |
| `tapped_share` | 0.124 | 0.141 | +0.107 |
| `mana_turn_four` | 3.844 | 0.797 | +0.113 |
| `hands_two_to_four_lands` | 0.732 | 0.059 | +0.264 |
| `curve_low` | 0.580 | 0.221 | -0.039 |
| `curve_high` | 0.180 | 0.158 | -0.084 |
| `ramp` | 5.348 | 6.806 | -0.049 |
| `draw` | 4.663 | 4.601 | +0.090 |
| `removal` | 7.758 | 4.572 | -0.172 |
| `wipe` | 1.008 | 1.746 | -0.082 |
| `interaction` | 3.030 | 3.800 | +0.010 |
| `empty_roles` | 1.031 | 0.844 | +0.013 |
| `fast_mana` | 0.553 | 1.506 | -0.033 |
| `game_changer` | 0.271 | 0.885 | -0.033 |
| `playset_share` | 0.569 | 0.236 | -0.552 |
| `singleton_share` | 0.145 | 0.157 | -0.714 |
| `source_spread` | 0.245 | 0.218 | -0.224 |

Thresholds: -3.889, -1.607, 0.784, 2.404. Loss: 1.030 over 3000 iterations. Defect detector accuracy on the holdout: 0.93, at a cut of 0.48.

The precon bar per broken axis, each precon over its own copies, with the cross pairs after it:

- lands: 1.00 of 327, cross 0.94 of 36476
- curve: 1.00 of 327, cross 0.98 of 36476
- colors: 0.97 of 75, cross 0.45 of 17066
- copies: 1.00 of 61, cross 0.83 of 18460
- synergy: 0.99 of 845, cross 0.85 of 70333

The misses per axis over the folds (M-7). Both passed: the detector passed the precon and the copy, and the ladder put the copy at or above the precon. Precon flagged: the detector flagged the precon and passed the copy. Both flagged: it flagged both and read the precon as the more broken. Own copy: each precon against its own copy on the axis, with its misses by the same kinds. Moved: the mean absolute standardized delta between a precon and its own copy, for the corpus features and the three features that move most.

| Axis | Share | Both passed | Precon flagged | Both flagged | Own copy | Own misses | Corpus moved | Most moved |
|---|---|---|---|---|---|---|---|---|
| lands | 0.94 of 36476 | 1378 | 303 | 565 | 327 of 327 | 0, 0, 0 | `card_rate` 0.01, `synergy` 0.00, `unseen_share` 0.09 | `land` 2.57, `hands_two_to_four_lands` 2.47, `color_sources` 1.46 |
| curve | 0.98 of 36476 | 451 | 131 | 319 | 327 of 327 | 0, 0, 0 | `card_rate` 0.04, `synergy` 0.00, `unseen_share` 0.20 | `curve_high` 2.68, `avg_mana_value` 2.22, `curve_low` 1.70 |
| colors | 0.45 of 17066 | 7463 | 1595 | 305 | 73 of 75 | 2, 0, 0 | `card_rate` 0.00, `synergy` 0.00, `unseen_share` 0.00 | `tapped_share` 1.25, `color_sources` 1.03, `source_spread` 0.96 |
| copies | 0.83 of 18460 | 2244 | 565 | 278 | 61 of 61 | 0, 0, 0 | `card_rate` 0.04, `synergy` 0.00, `unseen_share` 0.19 | `singleton_share` 3.56, `playset_share` 2.38, `empty_roles` 0.91 |
| synergy | 0.85 of 70333 | 6406 | 2184 | 1929 | 835 of 845 | 7, 0, 3 | `card_rate` 0.03, `synergy` 0.00, `unseen_share` 0.21 | `playset_share` 1.31, `empty_roles` 0.64, `removal` 0.61 |

The precons that lose most over the sampled pairs of the bar:

| Precon | Grade | Detector | Misses |
|---|---|---|---|
| Sultai Schemers (2014-09-26) | bad | 0.86 | 232 of 344 |
| Beasts (2005-08-01) | bad | 0.85 | 126 of 240 |
| Illusionary Might (2011-08-12) | bad | 0.83 | 126 of 240 |
| Jeskai Monks (2014-09-26) | bad | 0.80 | 167 of 333 |
| Second Sun Control (2018-04-06) | bad | 0.78 | 160 of 333 |
| Thallids (2005-08-01) | bad | 0.81 | 108 of 240 |
| Stampede (2004-06-04) | bad | 0.77 | 147 of 327 |
| Wizards (2005-08-01) | bad | 0.80 | 131 of 307 |
| Grave Advantage (2015-01-23) | bad | 0.75 | 96 of 240 |
| Repeat Performance (2012-08-03) | bad | 0.75 | 129 of 327 |

Confusion on the holdout over the folds, rows are the label and columns the grade, worst first:

- bad: [2552 61 124 4 4]
- baseline: [204 118 5 0 0]
- typical: [34 4 91 55 38]
- good: [174 0 849 1514 1463]
- great: [84 0 401 920 1145]

The cards the top lists hold most: Solitude, Quantum Riddler, Thoughtseize, Force of Negation, Prismatic Ending, Fatal Push, Ephemerate, Mishra's Bauble, Galvanic Discharge, Kozilek's Command, Malevolent Rumble, Ragavan, Nimble Pilferer.

## Run

- Suite `quality`, run `pr14b-quality-gate-run15`, on 2026-09-10, commit `d7f7621`.
- Roles: none, no provider call.
- Versions: card snapshot 2026-09-04, commanders_day `2026-09-07`, quality_model `20260910T012206Z`.
- Calls: 0. Cost: unpriced. Time: 116 seconds.

## Verdict

Verdict: FAIL. Time: 116 seconds, no provider call. Failed bars: commander: precon over own copy 0.92 of 945 (bar 0.95).
