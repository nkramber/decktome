# PR-14B quality gate

Run date: 2026-09-07. Card snapshot: 2026-09-04. Model: 20260907T160711Z.

## Separation on the holdout

The bars read 5 folds, every list holdout once, and the pairs of a bar sample evenly under the cap of 20000 (M-7). The precon bar reads each precon against its own broken copies (D-573). The cross pairs, every precon against every copy, stand as information.

| Format | Lists | Used | Synthetic | Holdout | Great over precon | Precon over own copy | Precon over bad, cross | Accuracy |
|---|---|---|---|---|---|---|---|---|
| commander | 23628 | 9229 | 6145 | 15374 | 0.96 of 97450 (bar 0.90) | 0.88 of 945 (bar 0.95) | 0.86 of 99933 | 0.65 |
| standard | 7356 | 5617 | 370 | 5987 | 0.93 of 2499 (bar 0.90) | 0.97 of 40 (bar 0.95) | 0.87 of 730 | 0.47 |
| modern | 17666 | 6835 | 2745 | 9580 | 0.99 of 98458 (bar 0.90) | 0.99 of 1635 (bar 0.95) | 0.85 of 99754 | 0.55 |

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

Tiers, worst first: bad, baseline, typical, good, great. Train counts: bad 320, baseline 7, good 3202, great 1235, typical 57. Holdout counts: bad 50, baseline 1, good 798, great 308, typical 9.

Cards with a rate: 806. Pairs that lift: 5923. Commanders with a signal: 0. Top lists shape: 23.1 lands at 2.50 over 1235 lists.

Features with no spread, dropped: game_changer, cedh_signal, high_bracket_share, commander_decks.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.117 | 0.061 | +1.381 |
| `unseen_share` | 0.016 | 0.082 | +0.106 |
| `synergy` | 1.587 | 0.736 | +0.969 |
| `land` | 22.991 | 2.398 | -0.040 |
| `avg_mana_value` | 2.549 | 0.660 | -0.935 |
| `color_sources` | 1.060 | 0.407 | +0.329 |
| `tapped_share` | 0.032 | 0.072 | -0.084 |
| `mana_turn_four` | 4.057 | 0.757 | +0.400 |
| `hands_two_to_four_lands` | 0.756 | 0.032 | +0.544 |
| `curve_low` | 0.633 | 0.175 | -0.304 |
| `curve_high` | 0.130 | 0.131 | +0.148 |
| `ramp` | 4.996 | 6.373 | -0.327 |
| `draw` | 4.879 | 3.417 | -0.053 |
| `removal` | 8.395 | 4.470 | +0.090 |
| `wipe` | 1.361 | 1.778 | +0.057 |
| `interaction` | 3.421 | 2.973 | +0.102 |
| `empty_roles` | 0.814 | 0.683 | +0.005 |
| `fast_mana` | 0.143 | 0.789 | +0.099 |
| `playset_share` | 0.598 | 0.182 | -0.042 |
| `singleton_share` | 0.089 | 0.089 | -0.485 |
| `source_spread` | 0.280 | 0.214 | +0.300 |

Thresholds: -4.466, -2.555, -0.647, 1.021. Loss: 1.085 over 3000 iterations. Defect detector accuracy on the holdout: 0.97, at a cut of 0.44.

The precon bar per broken axis, each precon over its own copies, with the cross pairs after it:

- lands: 1.00 of 8, cross 0.90 of 146
- curve: 1.00 of 8, cross 1.00 of 146
- colors: 0.75 of 4, cross 0.60 of 105
- copies: 1.00 of 3, cross 0.98 of 111
- synergy: 1.00 of 17, cross 0.83 of 222

The misses per axis over the folds (M-7). Both passed: the detector passed the precon and the copy, and the ladder put the copy at or above the precon. Precon flagged: the detector flagged the precon and passed the copy. Both flagged: it flagged both and read the precon as the more broken. Own copy: each precon against its own copy on the axis, with its misses by the same kinds. Moved: the mean absolute standardized delta between a precon and its own copy, for the corpus features and the three features that move most.

| Axis | Share | Both passed | Precon flagged | Both flagged | Own copy | Own misses | Corpus moved | Most moved |
|---|---|---|---|---|---|---|---|---|
| lands | 0.90 of 146 | 5 | 6 | 4 | 8 of 8 | 0, 0, 0 | `card_rate` 0.08, `synergy` 0.13, `unseen_share` 0.98 | `hands_two_to_four_lands` 4.17, `land` 3.27, `color_sources` 1.79 |
| curve | 1.00 of 146 | 0 | 0 | 0 | 8 of 8 | 0, 0, 0 | `card_rate` 0.25, `synergy` 0.19, `unseen_share` 2.84 | `avg_mana_value` 3.67, `curve_high` 3.24, `unseen_share` 2.84 |
| colors | 0.60 of 105 | 29 | 8 | 5 | 3 of 4 | 1, 0, 0 | `card_rate` 0.00, `synergy` 0.00, `unseen_share` 0.00 | `tapped_share` 3.12, `source_spread` 1.56, `color_sources` 1.31 |
| copies | 0.98 of 111 | 0 | 2 | 0 | 3 of 3 | 0, 0, 0 | `card_rate` 0.26, `synergy` 0.04, `unseen_share` 0.85 | `singleton_share` 4.96, `playset_share` 2.47, `empty_roles` 0.98 |
| synergy | 0.83 of 222 | 19 | 18 | 1 | 17 of 17 | 0, 0, 0 | `card_rate` 0.19, `synergy` 0.22, `unseen_share` 3.95 | `unseen_share` 3.95, `playset_share` 1.78, `color_sources` 1.60 |

The precons that lose most over the sampled pairs of the bar:

| Precon | Grade | Detector | Misses |
|---|---|---|---|
| Leonardo (2026-03-06) | bad | 0.47 | 25 of 100 |
| Donatello (2026-03-06) | bad | 0.41 | 19 of 100 |
| Angels (2026-01-23) | typical | 0.18 | 15 of 105 |
| Raphael (2026-03-06) | bad | 0.19 | 9 of 65 |
| Eerie (2026-04-24) | baseline | 0.04 | 6 of 50 |
| Michealangelo (2026-03-06) | bad | 0.26 | 12 of 100 |
| Pirates (2026-01-23) | typical | 0.07 | 9 of 105 |
| Lifegain (2026-04-24) | good | 0.01 | 2 of 105 |

Confusion on the holdout over the folds, rows are the label and columns the grade, worst first:

- bad: [337 1 18 9 5]
- baseline: [4 1 2 1 0]
- typical: [21 4 16 15 10]
- good: [215 1 437 1702 1645]
- great: [70 3 136 593 741]

The cards the top lists hold most: Burst Lightning, Stock Up, Opt, Sleight of Hand, Eddymurk Crab, Llanowar Elves, Spell Pierce, Flow State, Boomerang Basics, Stormchaser's Talent, Slickshot Show-Off, Badgermole Cub.

## The modern model

Tiers, worst first: bad, baseline, typical, good, great. Train counts: bad 2205, baseline 266, good 3155, great 1805, typical 175. Holdout counts: bad 540, baseline 61, good 845, great 481, typical 47.

Cards with a rate: 879. Pairs that lift: 6743. Commanders with a signal: 0. Top lists shape: 21.6 lands at 2.45 over 1805 lists.

Features with no spread, dropped: cedh_signal, high_bracket_share, commander_decks.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.064 | 0.046 | +0.451 |
| `unseen_share` | 0.235 | 0.381 | -0.740 |
| `synergy` | 1.389 | 1.014 | +1.372 |
| `land` | 21.658 | 3.066 | -0.069 |
| `avg_mana_value` | 2.761 | 0.943 | -0.332 |
| `color_sources` | 0.959 | 0.328 | +0.332 |
| `tapped_share` | 0.122 | 0.139 | +0.136 |
| `mana_turn_four` | 3.841 | 0.796 | +0.033 |
| `hands_two_to_four_lands` | 0.732 | 0.059 | +0.276 |
| `curve_low` | 0.579 | 0.222 | -0.032 |
| `curve_high` | 0.181 | 0.159 | -0.167 |
| `ramp` | 5.261 | 6.804 | +0.064 |
| `draw` | 4.708 | 4.673 | +0.038 |
| `removal` | 7.733 | 4.573 | -0.213 |
| `wipe` | 0.984 | 1.730 | +0.039 |
| `interaction` | 3.051 | 3.806 | +0.003 |
| `empty_roles` | 1.033 | 0.841 | +0.015 |
| `fast_mana` | 0.540 | 1.493 | -0.012 |
| `game_changer` | 0.276 | 0.890 | -0.108 |
| `playset_share` | 0.566 | 0.238 | -0.557 |
| `singleton_share` | 0.147 | 0.159 | -0.658 |
| `source_spread` | 0.246 | 0.218 | -0.225 |

Thresholds: -3.755, -1.470, 0.896, 2.447. Loss: 1.057 over 3000 iterations. Defect detector accuracy on the holdout: 0.93, at a cut of 0.48.

The precon bar per broken axis, each precon over its own copies, with the cross pairs after it:

- lands: 1.00 of 327, cross 0.93 of 36476
- curve: 1.00 of 327, cross 0.98 of 36476
- colors: 0.99 of 75, cross 0.46 of 17066
- copies: 1.00 of 61, cross 0.82 of 18460
- synergy: 0.98 of 845, cross 0.85 of 70333

The misses per axis over the folds (M-7). Both passed: the detector passed the precon and the copy, and the ladder put the copy at or above the precon. Precon flagged: the detector flagged the precon and passed the copy. Both flagged: it flagged both and read the precon as the more broken. Own copy: each precon against its own copy on the axis, with its misses by the same kinds. Moved: the mean absolute standardized delta between a precon and its own copy, for the corpus features and the three features that move most.

| Axis | Share | Both passed | Precon flagged | Both flagged | Own copy | Own misses | Corpus moved | Most moved |
|---|---|---|---|---|---|---|---|---|
| lands | 0.93 of 36476 | 1494 | 307 | 643 | 327 of 327 | 0, 0, 0 | `card_rate` 0.01, `synergy` 0.00, `unseen_share` 0.08 | `land` 2.58, `hands_two_to_four_lands` 2.47, `color_sources` 1.44 |
| curve | 0.98 of 36476 | 389 | 109 | 359 | 327 of 327 | 0, 0, 0 | `card_rate` 0.03, `synergy` 0.00, `unseen_share` 0.18 | `curve_high` 2.68, `avg_mana_value` 2.22, `curve_low` 1.70 |
| colors | 0.46 of 17066 | 7264 | 1637 | 325 | 74 of 75 | 1, 0, 0 | `card_rate` 0.00, `synergy` 0.00, `unseen_share` 0.00 | `tapped_share` 1.27, `color_sources` 1.03, `source_spread` 0.96 |
| copies | 0.82 of 18460 | 2362 | 582 | 287 | 61 of 61 | 0, 0, 0 | `card_rate` 0.04, `synergy` 0.00, `unseen_share` 0.19 | `singleton_share` 3.52, `playset_share` 2.36, `empty_roles` 0.90 |
| synergy | 0.85 of 70333 | 6525 | 2282 | 2017 | 828 of 845 | 13, 0, 4 | `card_rate` 0.03, `synergy` 0.00, `unseen_share` 0.20 | `playset_share` 1.30, `empty_roles` 0.66, `removal` 0.61 |

The precons that lose most over the sampled pairs of the bar:

| Precon | Grade | Detector | Misses |
|---|---|---|---|
| Sultai Schemers (2014-09-26) | bad | 0.87 | 239 of 344 |
| Beasts (2005-08-01) | bad | 0.84 | 127 of 240 |
| Illusionary Might (2011-08-12) | bad | 0.84 | 122 of 240 |
| Jeskai Monks (2014-09-26) | bad | 0.80 | 168 of 333 |
| Second Sun Control (2018-04-06) | bad | 0.77 | 156 of 333 |
| Thallids (2005-08-01) | bad | 0.81 | 112 of 240 |
| Wizards (2005-08-01) | bad | 0.81 | 142 of 307 |
| Stampede (2004-06-04) | bad | 0.77 | 144 of 327 |
| Spiraling Doom (2012-02-24) | bad | 0.77 | 130 of 307 |
| Stampede of Beasts (2010-07-16) | bad | 0.72 | 135 of 333 |

Confusion on the holdout over the folds, rows are the label and columns the grade, worst first:

- bad: [2542 68 131 1 3]
- baseline: [207 115 5 0 0]
- typical: [35 4 73 71 39]
- good: [176 0 783 1468 1573]
- great: [75 0 470 665 1076]

The cards the top lists hold most: Solitude, Quantum Riddler, Thoughtseize, Force of Negation, Prismatic Ending, Fatal Push, Ephemerate, Mishra's Bauble, Kozilek's Command, Psychic Frog, Malevolent Rumble, Galvanic Discharge.

## Run

- Suite `quality`, run `pr14b-quality-gate-run14`, on 2026-09-07, commit `d2e19f5`.
- Roles: none, no provider call.
- Versions: card snapshot 2026-09-04, commanders_day `2026-09-07`, quality_model `20260907T160711Z`.
- Calls: 0. Cost: unpriced. Time: 110 seconds.

## Verdict

Verdict: FAIL. Time: 110 seconds, no provider call. Failed bars: commander: precon over own copy 0.88 of 945 (bar 0.95).
