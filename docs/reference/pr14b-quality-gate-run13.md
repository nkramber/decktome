# PR-14B quality gate

Run date: 2026-09-07. Card snapshot: 2026-09-04. Model: 20260907T081746Z.

## Separation on the holdout

| Format | Lists | Used | Synthetic | Holdout | Great over precon | Precon over bad | Accuracy |
|---|---|---|---|---|---|---|---|
| commander | 23628 | 9229 | 6145 | 2997 | 0.97 of 20000 (bar 0.90) | 0.78 of 20000 (bar 0.95) | 0.65 |
| standard | 7356 | 5617 | 370 | 1166 | 0.98 of 308 (bar 0.90) | 0.88 of 50 (bar 0.95) | 0.44 |
| modern | 17666 | 6835 | 2745 | 1974 | 0.99 of 20000 (bar 0.90) | 0.86 of 20000 (bar 0.95) | 0.53 |

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
| `card_rate` | 0.152 | 0.139 | +0.975 |
| `unseen_share` | 0.073 | 0.103 | -0.233 |
| `synergy` | 0.490 | 0.535 | +0.166 |
| `land` | 30.795 | 5.609 | -0.261 |
| `avg_mana_value` | 2.802 | 0.859 | -0.436 |
| `color_sources` | 1.255 | 0.352 | +0.001 |
| `tapped_share` | 0.073 | 0.084 | +0.347 |
| `mana_turn_four` | 4.704 | 0.613 | +0.756 |
| `hands_two_to_four_lands` | 0.652 | 0.087 | +0.530 |
| `curve_low` | 0.505 | 0.188 | +0.194 |
| `curve_high` | 0.153 | 0.133 | -0.180 |
| `ramp` | 17.550 | 7.585 | +0.065 |
| `draw` | 7.439 | 4.084 | -0.231 |
| `removal` | 6.381 | 3.632 | +0.098 |
| `wipe` | 1.791 | 1.985 | -0.152 |
| `interaction` | 7.386 | 4.568 | +0.201 |
| `empty_roles` | 0.048 | 0.219 | +0.005 |
| `fast_mana` | 5.146 | 4.816 | +0.263 |
| `game_changer` | 7.238 | 7.395 | +0.114 |
| `source_spread` | 0.095 | 0.122 | -0.471 |
| `cedh_signal` | 0.278 | 0.200 | +0.220 |
| `high_bracket_share` | 0.215 | 0.250 | +0.089 |
| `commander_decks` | 5.867 | 4.128 | -0.129 |

Thresholds: -3.141, -1.433, 0.699, 3.027. Loss: 1.060 over 3000 iterations. Defect detector accuracy on the holdout: 0.93, at a cut of 0.44.

The precon bar per broken axis:

- lands: 0.97 of 10120
- curve: 0.96 of 10120
- colors: 0.92 of 6072
- synergy: 0.78 of 20000

Confusion on the holdout, rows are the label and columns the grade, worst first:

- bad: [1123 3 18 6 0]
- baseline: [28 2 13 1 0]
- typical: [78 1 100 4 3]
- good: [51 2 40 272 412]
- great: [40 0 61 282 457]

The cards the top lists hold most: Sol Ring, Lotus Petal, Chrome Mox, Mana Vault, Mox Diamond, Arcane Signet, Mystic Remora, Mental Misstep, Flusterstorm, Rhystic Study, Force of Will, Mox Amber.

## The standard model

Tiers, worst first: bad, baseline, typical, good, great. Train counts: bad 320, baseline 7, good 3202, great 1235, typical 57. Holdout counts: bad 50, baseline 1, good 798, great 308, typical 9.

Cards with a rate: 806. Pairs that lift: 5923. Commanders with a signal: 0. Top lists shape: 23.1 lands at 2.50 over 1235 lists.

Features with no spread, dropped: game_changer, cedh_signal, high_bracket_share, commander_decks.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.117 | 0.061 | +1.400 |
| `unseen_share` | 0.016 | 0.082 | +0.116 |
| `synergy` | 1.587 | 0.736 | +0.984 |
| `land` | 22.991 | 2.398 | -0.041 |
| `avg_mana_value` | 2.548 | 0.660 | -0.872 |
| `color_sources` | 1.060 | 0.407 | +0.347 |
| `tapped_share` | 0.032 | 0.072 | -0.081 |
| `mana_turn_four` | 4.057 | 0.756 | +0.394 |
| `hands_two_to_four_lands` | 0.756 | 0.032 | +0.552 |
| `curve_low` | 0.633 | 0.175 | -0.266 |
| `curve_high` | 0.130 | 0.131 | +0.110 |
| `ramp` | 4.997 | 6.377 | -0.333 |
| `draw` | 4.880 | 3.418 | -0.064 |
| `removal` | 8.398 | 4.480 | +0.094 |
| `wipe` | 1.362 | 1.780 | +0.059 |
| `interaction` | 3.421 | 2.972 | +0.096 |
| `empty_roles` | 0.817 | 0.684 | -0.036 |
| `fast_mana` | 0.141 | 0.781 | +0.126 |
| `playset_share` | 0.598 | 0.182 | -0.051 |
| `singleton_share` | 0.089 | 0.089 | -0.493 |
| `source_spread` | 0.280 | 0.214 | +0.304 |

Thresholds: -4.475, -2.560, -0.647, 1.023. Loss: 1.083 over 3000 iterations. Defect detector accuracy on the holdout: 0.98, at a cut of 0.52.

The precon bar per broken axis:

- lands: 1.00 of 10
- curve: 1.00 of 10
- colors: 0.38 of 8
- copies: 1.00 of 9
- synergy: 0.92 of 13

Confusion on the holdout, rows are the label and columns the grade, worst first:

- bad: [49 0 1 0 0]
- baseline: [0 1 0 0 0]
- typical: [6 0 2 1 0]
- good: [39 1 112 313 333]
- great: [14 1 28 120 145]

The cards the top lists hold most: Burst Lightning, Stock Up, Opt, Sleight of Hand, Eddymurk Crab, Llanowar Elves, Spell Pierce, Flow State, Boomerang Basics, Stormchaser's Talent, Slickshot Show-Off, Badgermole Cub.

## The modern model

Tiers, worst first: bad, baseline, typical, good, great. Train counts: bad 2205, baseline 266, good 3155, great 1805, typical 175. Holdout counts: bad 540, baseline 61, good 845, great 481, typical 47.

Cards with a rate: 879. Pairs that lift: 6743. Commanders with a signal: 0. Top lists shape: 21.6 lands at 2.45 over 1805 lists.

Features with no spread, dropped: cedh_signal, high_bracket_share, commander_decks.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.064 | 0.046 | +0.450 |
| `unseen_share` | 0.235 | 0.381 | -0.739 |
| `synergy` | 1.389 | 1.014 | +1.373 |
| `land` | 21.658 | 3.066 | -0.069 |
| `avg_mana_value` | 2.760 | 0.941 | -0.327 |
| `color_sources` | 0.959 | 0.328 | +0.335 |
| `tapped_share` | 0.122 | 0.139 | +0.135 |
| `mana_turn_four` | 3.842 | 0.797 | +0.025 |
| `hands_two_to_four_lands` | 0.732 | 0.059 | +0.279 |
| `curve_low` | 0.579 | 0.222 | -0.031 |
| `curve_high` | 0.181 | 0.159 | -0.172 |
| `ramp` | 5.259 | 6.809 | +0.070 |
| `draw` | 4.708 | 4.674 | +0.039 |
| `removal` | 7.718 | 4.563 | -0.209 |
| `wipe` | 0.984 | 1.728 | +0.035 |
| `interaction` | 3.053 | 3.808 | -0.001 |
| `empty_roles` | 1.038 | 0.842 | +0.005 |
| `fast_mana` | 0.539 | 1.490 | -0.010 |
| `game_changer` | 0.277 | 0.891 | -0.111 |
| `playset_share` | 0.566 | 0.238 | -0.557 |
| `singleton_share` | 0.147 | 0.159 | -0.659 |
| `source_spread` | 0.246 | 0.218 | -0.224 |

Thresholds: -3.756, -1.471, 0.896, 2.447. Loss: 1.057 over 3000 iterations. Defect detector accuracy on the holdout: 0.94, at a cut of 0.46.

The precon bar per broken axis:

- lands: 0.93 of 6588
- curve: 0.99 of 6588
- colors: 0.51 of 3233
- copies: 0.89 of 3599
- synergy: 0.86 of 12932

Confusion on the holdout, rows are the label and columns the grade, worst first:

- bad: [503 11 25 0 1]
- baseline: [40 20 1 0 0]
- typical: [11 3 13 13 7]
- good: [36 0 165 292 352]
- great: [14 0 90 152 225]

The cards the top lists hold most: Solitude, Quantum Riddler, Thoughtseize, Force of Negation, Prismatic Ending, Fatal Push, Ephemerate, Mishra's Bauble, Kozilek's Command, Psychic Frog, Malevolent Rumble, Galvanic Discharge.

## Run

- Suite `quality`, run `pr14b-quality-gate-run13`, on 2026-09-07, commit `402ec6f`.
- Roles: none, no provider call.
- Versions: card snapshot 2026-09-04, commanders_day `2026-09-07`, quality_model `20260907T081746Z`.
- Calls: 0. Cost: unpriced. Time: 58 seconds.

## Verdict

Verdict: FAIL. Time: 58 seconds, no provider call. Failed bars: commander: precon over bad 0.78 of 20000 (bar 0.95); standard: precon over bad 0.88 of 50 (bar 0.95); modern: precon over bad 0.86 of 20000 (bar 0.95).
