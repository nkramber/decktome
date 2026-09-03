# PR-14B quality gate

Run date: 2026-09-03. Card snapshot: 2026-09-03. Model: 20260903T210813Z.

## Separation on the holdout

| Format | Lists | Used | Synthetic | Holdout | Great over precon | Precon over bad | Accuracy |
|---|---|---|---|---|---|---|---|
| commander | 22415 | 9226 | 6130 | 3029 | 0.98 of 20000 (bar 0.90) | 0.78 of 20000 (bar 0.95) | 0.65 |
| standard | 5360 | 5144 | 40 | 1049 | 0.92 of 231 (bar 0.90) | 1.00 of 5 (bar 0.95) | 0.51 |
| modern | 13338 | 6234 | 2380 | 1735 | 0.99 of 20000 (bar 0.90) | 0.88 of 20000 (bar 0.95) | 0.51 |

## The bracket 5 offer

Commander reads of 2026-09-03: 1563 rows, 937 with Topdeck.gg entries.

| Commander | cEDH signal | Top cuts | Entries |
|---|---|---|---|
| Aesi, Tyrant of Gyre Strait | 0.52 | 7 | 11 |
| Kotis, Sibsig Champion | 0.51 | 9 | 15 |
| Niv-Mizzet, Parun | 0.50 | 7 | 36 |

## The commander model

Tiers, worst first: bad, baseline, typical, good, great. Train counts: bad 4980, baseline 145, good 3186, great 3165, typical 851. Holdout counts: bad 1150, baseline 44, good 814, great 835, typical 186.

Cards with a rate: 9633. Pairs that lift: 200000. Commanders with a signal: 1553. Top lists shape: 27.5 lands at 2.19 over 3165 lists.

Features with no spread, dropped: playset_share, singleton_share.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.153 | 0.139 | +0.945 |
| `unseen_share` | 0.078 | 0.109 | -0.297 |
| `synergy` | 0.491 | 0.586 | +0.128 |
| `land` | 30.788 | 5.627 | -0.250 |
| `avg_mana_value` | 2.803 | 0.858 | -0.464 |
| `color_sources` | 1.256 | 0.353 | -0.023 |
| `tapped_share` | 0.074 | 0.085 | +0.359 |
| `mana_turn_four` | 4.703 | 0.611 | +0.775 |
| `hands_two_to_four_lands` | 0.652 | 0.087 | +0.514 |
| `curve_low` | 0.505 | 0.188 | +0.214 |
| `curve_high` | 0.154 | 0.133 | -0.149 |
| `ramp` | 17.572 | 7.559 | +0.047 |
| `draw` | 7.442 | 4.088 | -0.241 |
| `removal` | 6.359 | 3.582 | +0.108 |
| `wipe` | 1.788 | 1.972 | -0.160 |
| `interaction` | 7.422 | 4.587 | +0.181 |
| `empty_roles` | 0.053 | 0.228 | -0.018 |
| `fast_mana` | 5.162 | 4.802 | +0.270 |
| `game_changer` | 7.272 | 7.391 | +0.113 |
| `source_spread` | 0.095 | 0.122 | -0.497 |
| `cedh_signal` | 0.279 | 0.201 | +0.216 |
| `high_bracket_share` | 0.215 | 0.250 | +0.107 |
| `commander_decks` | 5.869 | 4.123 | -0.180 |

Thresholds: -3.181, -1.468, 0.700, 3.065. Loss: 1.052 over 3000 iterations. Defect detector accuracy on the holdout: 0.93, at a cut of 0.48.

The precon bar per broken axis:

- lands: 0.97 of 10120
- curve: 0.96 of 10120
- colors: 0.92 of 6072
- synergy: 0.77 of 20000

Confusion on the holdout, rows are the label and columns the grade, worst first:

- bad: [1113 1 31 5 0]
- baseline: [28 2 14 0 0]
- typical: [79 0 100 4 3]
- good: [37 5 49 305 418]
- great: [38 4 52 281 460]

The cards the top lists hold most: Sol Ring, Lotus Petal, Chrome Mox, Mana Vault, Mox Diamond, Arcane Signet, Mystic Remora, Mental Misstep, Flusterstorm, Rhystic Study, Force of Will, Mox Amber.

## The standard model

Tiers, worst first: bad, baseline, good, great. Train counts: bad 35, baseline 7, good 3188, great 905. Holdout counts: bad 5, baseline 1, good 812, great 231.

Cards with a rate: 775. Pairs that lift: 5281. Commanders with a signal: 0. Top lists shape: 23.1 lands at 2.51 over 905 lists.

Features with no spread, dropped: game_changer, cedh_signal, high_bracket_share, commander_decks.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.123 | 0.057 | +1.890 |
| `unseen_share` | 0.005 | 0.057 | +0.108 |
| `synergy` | 1.656 | 0.658 | +1.443 |
| `land` | 23.108 | 2.259 | +0.055 |
| `avg_mana_value` | 2.506 | 0.600 | -0.205 |
| `color_sources` | 1.057 | 0.389 | +0.821 |
| `tapped_share` | 0.030 | 0.067 | -0.102 |
| `mana_turn_four` | 4.081 | 0.757 | +0.222 |
| `hands_two_to_four_lands` | 0.759 | 0.025 | +0.198 |
| `curve_low` | 0.641 | 0.170 | -0.369 |
| `curve_high` | 0.123 | 0.122 | -0.206 |
| `ramp` | 5.082 | 6.462 | -0.208 |
| `draw` | 4.799 | 3.352 | +0.110 |
| `removal` | 8.297 | 4.454 | +0.069 |
| `wipe` | 1.342 | 1.734 | +0.044 |
| `interaction` | 3.446 | 2.947 | +0.246 |
| `empty_roles` | 0.814 | 0.681 | +0.088 |
| `fast_mana` | 0.117 | 0.688 | +0.335 |
| `playset_share` | 0.605 | 0.170 | -0.179 |
| `singleton_share` | 0.080 | 0.061 | -0.258 |
| `source_spread` | 0.281 | 0.204 | +0.284 |

Thresholds: -6.227, -2.750, 0.333. Loss: 0.627 over 3000 iterations. Defect detector accuracy on the holdout: 0.93, at a cut of 0.20.

The precon bar per broken axis:

- lands: 1.00 of 1
- curve: 1.00 of 1
- colors: 1.00 of 1
- copies: 1.00 of 1
- synergy: 1.00 of 1

Confusion on the holdout, rows are the label and columns the grade, worst first:

- bad: [5 0 0 0]
- baseline: [1 0 0 0]
- good: [28 15 398 371]
- great: [3 3 93 132]

The cards the top lists hold most: Burst Lightning, Stock Up, Opt, Sleight of Hand, Llanowar Elves, Spell Pierce, Flow State, Eddymurk Crab, Boomerang Basics, Stormchaser's Talent, Badgermole Cub, Slickshot Show-Off.

## The modern model

Tiers, worst first: bad, baseline, typical, good, great. Train counts: bad 1935, baseline 266, good 3169, great 1388, typical 121. Holdout counts: bad 445, baseline 61, good 831, great 370, typical 28.

Cards with a rate: 862. Pairs that lift: 6592. Commanders with a signal: 0. Top lists shape: 21.5 lands at 2.47 over 1388 lists.

Features with no spread, dropped: cedh_signal, high_bracket_share, commander_decks.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.064 | 0.047 | +0.439 |
| `unseen_share` | 0.245 | 0.393 | -0.803 |
| `synergy` | 1.385 | 1.022 | +1.435 |
| `land` | 21.686 | 3.032 | -0.004 |
| `avg_mana_value` | 2.764 | 0.937 | -0.357 |
| `color_sources` | 0.967 | 0.323 | +0.284 |
| `tapped_share` | 0.119 | 0.134 | +0.082 |
| `mana_turn_four` | 3.850 | 0.804 | -0.055 |
| `hands_two_to_four_lands` | 0.733 | 0.058 | +0.212 |
| `curve_low` | 0.579 | 0.225 | -0.012 |
| `curve_high` | 0.181 | 0.156 | -0.145 |
| `ramp` | 5.253 | 6.866 | +0.106 |
| `draw` | 4.673 | 4.745 | +0.048 |
| `removal` | 7.601 | 4.482 | -0.225 |
| `wipe` | 0.996 | 1.762 | -0.037 |
| `interaction` | 3.016 | 3.815 | +0.084 |
| `empty_roles` | 1.056 | 0.845 | +0.092 |
| `fast_mana` | 0.501 | 1.437 | +0.046 |
| `game_changer` | 0.282 | 0.898 | -0.136 |
| `playset_share` | 0.560 | 0.239 | -0.507 |
| `singleton_share` | 0.147 | 0.154 | -0.582 |
| `source_spread` | 0.242 | 0.214 | -0.240 |

Thresholds: -3.893, -1.500, 0.945, 2.484. Loss: 1.043 over 3000 iterations. Defect detector accuracy on the holdout: 0.95, at a cut of 0.46.

The precon bar per broken axis:

- lands: 0.94 of 5429
- curve: 0.99 of 5429
- colors: 0.53 of 2196
- copies: 0.88 of 2440
- synergy: 0.88 of 11651

Confusion on the holdout, rows are the label and columns the grade, worst first:

- bad: [413 15 15 1 1]
- baseline: [40 21 0 0 0]
- typical: [7 1 9 7 4]
- good: [35 0 176 264 356]
- great: [13 0 79 98 180]

The cards the top lists hold most: Solitude, Thoughtseize, Quantum Riddler, Force of Negation, Prismatic Ending, Fatal Push, Psychic Frog, Kozilek's Command, Malevolent Rumble, Ephemerate, Mishra's Bauble, Lightning Bolt.

## Verdict

Verdict: FAIL. Time: 56 seconds, no provider call. Failed bars: commander: precon over bad 0.78 of 20000 (bar 0.95); modern: precon over bad 0.88 of 20000 (bar 0.95).
