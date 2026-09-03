# PR-14B quality gate

Run date: 2026-09-03. Card snapshot: 2026-09-03. Model: 20260903T183943Z.

## Separation on the holdout

| Format | Lists | Used | Synthetic | Holdout | Great over precon | Precon over bad | Accuracy |
|---|---|---|---|---|---|---|---|
| commander | 22407 | 9226 | 6130 | 3030 | 0.98 of 20000 (bar 0.90) | 0.78 of 20000 (bar 0.95) | 0.65 |
| standard | 5336 | 5136 | 40 | 1049 | 0.92 of 229 (bar 0.90) | 1.00 of 5 (bar 0.95) | 0.51 |
| modern | 12913 | 6037 | 1635 | 1547 | 0.99 of 20000 (bar 0.90) | 0.95 of 18605 (bar 0.95) | 0.61 |

## The bracket 5 offer

Commander reads of 2026-09-03: 1563 rows, 937 with Topdeck.gg entries.

| Commander | cEDH signal | Top cuts | Entries |
|---|---|---|---|
| Aesi, Tyrant of Gyre Strait | 0.52 | 7 | 11 |
| Kotis, Sibsig Champion | 0.51 | 9 | 15 |
| Niv-Mizzet, Parun | 0.50 | 7 | 36 |

## The commander model

Tiers, worst first: bad, baseline, typical, good, great. Train counts: bad 4980, baseline 145, good 3185, great 3165, typical 851. Holdout counts: bad 1150, baseline 44, good 815, great 835, typical 186.

Cards with a rate: 9632. Pairs that lift: 200000. Commanders with a signal: 1553. Top lists shape: 27.5 lands at 2.19 over 3165 lists.

Features with no spread, dropped: playset_share, singleton_share.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.154 | 0.140 | +0.941 |
| `unseen_share` | 0.078 | 0.109 | -0.298 |
| `synergy` | 0.490 | 0.585 | +0.129 |
| `land` | 30.787 | 5.627 | -0.247 |
| `avg_mana_value` | 2.803 | 0.858 | -0.459 |
| `color_sources` | 1.256 | 0.353 | -0.022 |
| `tapped_share` | 0.074 | 0.085 | +0.359 |
| `mana_turn_four` | 4.704 | 0.611 | +0.774 |
| `hands_two_to_four_lands` | 0.652 | 0.087 | +0.512 |
| `curve_low` | 0.505 | 0.188 | +0.216 |
| `curve_high` | 0.154 | 0.133 | -0.150 |
| `ramp` | 17.575 | 7.561 | +0.044 |
| `draw` | 7.450 | 4.086 | -0.243 |
| `removal` | 6.358 | 3.580 | +0.111 |
| `wipe` | 1.794 | 1.973 | -0.169 |
| `interaction` | 7.416 | 4.588 | +0.186 |
| `empty_roles` | 0.052 | 0.228 | -0.016 |
| `fast_mana` | 5.162 | 4.803 | +0.272 |
| `game_changer` | 7.274 | 7.394 | +0.116 |
| `source_spread` | 0.095 | 0.122 | -0.496 |
| `cedh_signal` | 0.279 | 0.201 | +0.215 |
| `high_bracket_share` | 0.215 | 0.250 | +0.108 |
| `commander_decks` | 5.867 | 4.123 | -0.181 |

Thresholds: -3.185, -1.471, 0.700, 3.068. Loss: 1.052 over 3000 iterations. Defect detector accuracy on the holdout: 0.93, at a cut of 0.48.

The precon bar per broken axis:

- lands: 0.96 of 10120
- curve: 0.96 of 10120
- colors: 0.92 of 6072
- synergy: 0.77 of 20000

Confusion on the holdout, rows are the label and columns the grade, worst first:

- bad: [1114 1 31 4 0]
- baseline: [28 2 14 0 0]
- typical: [78 1 100 4 3]
- good: [37 5 50 305 418]
- great: [38 4 52 282 459]

The cards the top lists hold most: Sol Ring, Lotus Petal, Chrome Mox, Mana Vault, Mox Diamond, Arcane Signet, Mystic Remora, Mental Misstep, Flusterstorm, Rhystic Study, Force of Will, Mox Amber.

## The standard model

Tiers, worst first: bad, baseline, good, great. Train counts: bad 35, baseline 7, good 3186, great 899. Holdout counts: bad 5, baseline 1, good 814, great 229.

Cards with a rate: 771. Pairs that lift: 5234. Commanders with a signal: 0. Top lists shape: 23.1 lands at 2.51 over 899 lists.

Features with no spread, dropped: game_changer, cedh_signal, high_bracket_share, commander_decks.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.124 | 0.058 | +1.878 |
| `unseen_share` | 0.005 | 0.057 | +0.104 |
| `synergy` | 1.655 | 0.657 | +1.437 |
| `land` | 23.107 | 2.256 | +0.060 |
| `avg_mana_value` | 2.503 | 0.599 | -0.206 |
| `color_sources` | 1.056 | 0.388 | +0.853 |
| `tapped_share` | 0.030 | 0.067 | -0.109 |
| `mana_turn_four` | 4.081 | 0.756 | +0.202 |
| `hands_two_to_four_lands` | 0.759 | 0.025 | +0.201 |
| `curve_low` | 0.641 | 0.170 | -0.352 |
| `curve_high` | 0.122 | 0.122 | -0.173 |
| `ramp` | 5.078 | 6.460 | -0.204 |
| `draw` | 4.802 | 3.343 | +0.081 |
| `removal` | 8.301 | 4.455 | +0.080 |
| `wipe` | 1.339 | 1.733 | +0.043 |
| `interaction` | 3.437 | 2.953 | +0.217 |
| `empty_roles` | 0.814 | 0.682 | +0.018 |
| `fast_mana` | 0.117 | 0.689 | +0.329 |
| `playset_share` | 0.606 | 0.169 | -0.191 |
| `singleton_share` | 0.080 | 0.061 | -0.262 |
| `source_spread` | 0.281 | 0.203 | +0.297 |

Thresholds: -6.235, -2.764, 0.334. Loss: 0.628 over 3000 iterations. Defect detector accuracy on the holdout: 0.93, at a cut of 0.20.

The precon bar per broken axis:

- lands: 1.00 of 1
- curve: 1.00 of 1
- colors: 1.00 of 1
- copies: 1.00 of 1
- synergy: 1.00 of 1

Confusion on the holdout, rows are the label and columns the grade, worst first:

- bad: [5 0 0 0]
- baseline: [1 0 0 0]
- good: [28 13 400 373]
- great: [3 3 90 133]

The cards the top lists hold most: Burst Lightning, Stock Up, Opt, Sleight of Hand, Llanowar Elves, Spell Pierce, Flow State, Eddymurk Crab, Boomerang Basics, Stormchaser's Talent, Badgermole Cub, Slickshot Show-Off.

## The modern model

Tiers, worst first: bad, baseline, good, great. Train counts: bad 1330, baseline 266, good 3178, great 1351. Holdout counts: bad 305, baseline 61, good 822, great 359.

Cards with a rate: 863. Pairs that lift: 6674. Commanders with a signal: 0. Top lists shape: 21.5 lands at 2.47 over 1351 lists.

Features with no spread, dropped: cedh_signal, high_bracket_share, commander_decks.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.063 | 0.048 | +0.707 |
| `unseen_share` | 0.241 | 0.408 | -1.675 |
| `synergy` | 1.466 | 1.031 | +1.231 |
| `land` | 21.799 | 2.945 | +0.318 |
| `avg_mana_value` | 2.732 | 0.928 | -0.317 |
| `color_sources` | 0.994 | 0.304 | +0.045 |
| `tapped_share` | 0.119 | 0.136 | +0.118 |
| `mana_turn_four` | 3.866 | 0.809 | +0.078 |
| `hands_two_to_four_lands` | 0.736 | 0.054 | -0.041 |
| `curve_low` | 0.585 | 0.226 | +0.139 |
| `curve_high` | 0.175 | 0.152 | -0.126 |
| `ramp` | 5.363 | 6.961 | +0.013 |
| `draw` | 4.728 | 4.866 | +0.025 |
| `removal` | 7.512 | 4.481 | -0.097 |
| `wipe` | 0.970 | 1.714 | +0.045 |
| `interaction` | 2.933 | 3.769 | +0.103 |
| `empty_roles` | 1.082 | 0.851 | +0.043 |
| `fast_mana` | 0.507 | 1.437 | -0.025 |
| `game_changer` | 0.275 | 0.892 | -0.216 |
| `playset_share` | 0.562 | 0.233 | -0.949 |
| `singleton_share` | 0.140 | 0.137 | -0.790 |
| `source_spread` | 0.227 | 0.200 | -0.113 |

Thresholds: -5.264, -1.417, 2.133. Loss: 0.646 over 3000 iterations. Defect detector accuracy on the holdout: 0.98, at a cut of 0.48.

The precon bar per broken axis:

- lands: 0.98 of 3721
- curve: 0.99 of 3721
- colors: 0.61 of 732
- copies: 0.80 of 732
- synergy: 0.95 of 9699

Confusion on the holdout, rows are the label and columns the grade, worst first:

- bad: [292 13 0 0]
- baseline: [26 35 0 0]
- good: [7 2 435 378]
- great: [2 0 168 189]

The cards the top lists hold most: Solitude, Thoughtseize, Quantum Riddler, Force of Negation, Prismatic Ending, Fatal Push, Psychic Frog, Kozilek's Command, Malevolent Rumble, Ephemerate, Lightning Bolt, Mishra's Bauble.

## Verdict

Verdict: FAIL. Time: 53 seconds, no provider call. Failed bars: commander: precon over bad 0.78 of 20000 (bar 0.95); modern: precon over bad 0.95 of 18605 (bar 0.95).
