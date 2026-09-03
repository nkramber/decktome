# PR-14B quality gate

Run date: 2026-09-03. Card snapshot: 2026-09-02. Model: 20260903T010127Z.

## Separation on the holdout

| Format | Lists | Used | Synthetic | Holdout | Great over precon | Precon over bad | Accuracy |
|---|---|---|---|---|---|---|---|
| commander | 21538 | 8388 | 1940 | 2150 | 0.96 of 20000 (bar 0.90) | 0.88 of 18480 (bar 0.95) | 0.55 |
| standard | 3316 | 3313 | 40 | 666 | 0.97 of 147 (bar 0.90) | 1.00 of 5 (bar 0.95) | 0.76 |
| modern | 8745 | 5407 | 1635 | 1420 | 1.00 of 13908 (bar 0.90) | 0.95 of 18605 (bar 0.95) | 0.79 |

## The bracket 5 offer

Commander reads of 2026-09-02: 1007 rows, 937 with Topdeck.gg entries.

| Commander | cEDH signal | Top cuts | Entries |
|---|---|---|---|
| Aesi, Tyrant of Gyre Strait | 0.52 | 7 | 11 |
| Kotis, Sibsig Champion | 0.51 | 9 | 15 |
| Niv-Mizzet, Parun | 0.50 | 7 | 36 |

## The commander model

Tiers, worst first: bad, baseline, typical, good, great. Train counts: bad 1520, baseline 145, good 3186, great 3168, typical 159. Holdout counts: bad 420, baseline 44, good 814, great 832, typical 40.

Cards with a rate: 9633. Pairs that lift: 200000. Commanders with a signal: 996. Top lists shape: 27.5 lands at 2.19 over 3168 lists.

Features with no spread, dropped: playset_share, singleton_share.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.214 | 0.134 | +0.490 |
| `unseen_share` | 0.034 | 0.089 | -0.905 |
| `synergy` | 0.589 | 0.676 | +0.402 |
| `land` | 29.198 | 5.405 | -0.079 |
| `avg_mana_value` | 2.507 | 0.789 | -0.297 |
| `color_sources` | 1.254 | 0.304 | -0.016 |
| `tapped_share` | 0.074 | 0.097 | +0.239 |
| `mana_turn_four` | 4.918 | 0.589 | +0.411 |
| `hands_two_to_four_lands` | 0.627 | 0.081 | +0.127 |
| `curve_low` | 0.570 | 0.179 | +0.063 |
| `curve_high` | 0.116 | 0.121 | -0.217 |
| `ramp` | 19.940 | 6.986 | -0.009 |
| `draw` | 6.954 | 3.543 | -0.191 |
| `removal` | 5.707 | 3.283 | +0.100 |
| `wipe` | 1.256 | 1.694 | -0.196 |
| `interaction` | 8.407 | 4.626 | +0.253 |
| `empty_roles` | 0.044 | 0.214 | -0.010 |
| `fast_mana` | 7.138 | 4.737 | -0.077 |
| `game_changer` | 10.365 | 7.235 | -0.009 |
| `source_spread` | 0.102 | 0.111 | -0.370 |
| `cedh_signal` | 0.375 | 0.152 | +0.212 |
| `high_bracket_share` | 0.169 | 0.272 | +0.118 |
| `commander_decks` | 3.704 | 4.743 | -0.415 |

Thresholds: -2.616, -2.249, -1.835, 1.274. Loss: 0.823 over 3000 iterations. Defect detector accuracy on the holdout: 0.94, at a cut of 0.54.

The precon bar per broken axis:

- lands: 0.98 of 3696
- curve: 0.97 of 3696
- colors: 0.96 of 3036
- synergy: 0.76 of 8052

Confusion on the holdout, rows are the label and columns the grade, worst first:

- bad: [417 0 0 3 0]
- baseline: [32 0 0 12 0]
- typical: [22 0 0 18 0]
- good: [49 0 0 261 504]
- great: [42 0 0 287 503]

The cards the top lists hold most: Sol Ring, Lotus Petal, Chrome Mox, Mana Vault, Mox Diamond, Arcane Signet, Mystic Remora, Mental Misstep, Flusterstorm, Rhystic Study, Force of Will, Mox Amber.

## The standard model

Tiers, worst first: bad, baseline, good, great. Train counts: bad 35, baseline 7, good 2064, great 581. Holdout counts: bad 5, baseline 1, good 513, great 147.

Cards with a rate: 678. Pairs that lift: 4413. Commanders with a signal: 0. Top lists shape: 23.2 lands at 2.55 over 581 lists.

Features with no spread, dropped: game_changer, cedh_signal, high_bracket_share, commander_decks.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.118 | 0.052 | +0.572 |
| `unseen_share` | 0.009 | 0.072 | -0.522 |
| `synergy` | 1.720 | 0.668 | +0.364 |
| `land` | 23.159 | 2.249 | +0.003 |
| `avg_mana_value` | 2.550 | 0.577 | -0.095 |
| `color_sources` | 1.079 | 0.420 | +0.129 |
| `tapped_share` | 0.029 | 0.064 | +0.043 |
| `mana_turn_four` | 4.109 | 0.766 | -0.107 |
| `hands_two_to_four_lands` | 0.760 | 0.026 | +0.108 |
| `curve_low` | 0.629 | 0.165 | -0.154 |
| `curve_high` | 0.122 | 0.121 | -0.033 |
| `ramp` | 5.252 | 6.467 | +0.003 |
| `draw` | 4.819 | 3.312 | +0.057 |
| `removal` | 8.059 | 4.417 | +0.057 |
| `wipe` | 1.381 | 1.729 | -0.035 |
| `interaction` | 3.717 | 2.939 | +0.073 |
| `empty_roles` | 0.765 | 0.695 | -0.005 |
| `fast_mana` | 0.066 | 0.521 | +0.001 |
| `playset_share` | 0.603 | 0.169 | +0.059 |
| `singleton_share` | 0.079 | 0.061 | -0.027 |
| `source_spread` | 0.280 | 0.219 | +0.006 |

Thresholds: -5.019, -4.754, 1.419. Loss: 0.534 over 3000 iterations. Defect detector accuracy on the holdout: 0.95, at a cut of 0.20.

The precon bar per broken axis:

- lands: 1.00 of 1
- curve: 1.00 of 1
- colors: 1.00 of 1
- copies: 1.00 of 1
- synergy: 1.00 of 1

Confusion on the holdout, rows are the label and columns the grade, worst first:

- bad: [5 0 0 0]
- baseline: [1 0 0 0]
- good: [12 0 501 0]
- great: [2 0 145 0]

The cards the top lists hold most: Stock Up, Burst Lightning, Llanowar Elves, Opt, Sleight of Hand, Spell Pierce, Eddymurk Crab, Surrak, Elusive Hunter, Meltstrider's Resolve, Flow State, Tablet of Discovery, Badgermole Cub.

## The modern model

Tiers, worst first: bad, baseline, good, great. Train counts: bad 1330, baseline 266, good 3174, great 852. Holdout counts: bad 305, baseline 61, good 826, great 228.

Cards with a rate: 820. Pairs that lift: 6257. Commanders with a signal: 0. Top lists shape: 21.5 lands at 2.52 over 852 lists.

Features with no spread, dropped: cedh_signal, high_bracket_share, commander_decks.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.063 | 0.050 | +0.735 |
| `unseen_share` | 0.263 | 0.420 | -1.524 |
| `synergy` | 1.417 | 1.045 | +1.043 |
| `land` | 21.822 | 2.935 | +0.086 |
| `avg_mana_value` | 2.779 | 0.940 | -0.212 |
| `color_sources` | 0.987 | 0.309 | +0.020 |
| `tapped_share` | 0.117 | 0.134 | +0.062 |
| `mana_turn_four` | 3.867 | 0.809 | +0.076 |
| `hands_two_to_four_lands` | 0.737 | 0.054 | -0.016 |
| `curve_low` | 0.575 | 0.228 | +0.152 |
| `curve_high` | 0.181 | 0.154 | -0.027 |
| `ramp` | 5.259 | 6.962 | +0.105 |
| `draw` | 4.734 | 4.940 | +0.017 |
| `removal` | 7.431 | 4.438 | -0.046 |
| `wipe` | 0.943 | 1.710 | +0.072 |
| `interaction` | 2.894 | 3.769 | +0.101 |
| `empty_roles` | 1.105 | 0.850 | -0.024 |
| `fast_mana` | 0.478 | 1.403 | -0.051 |
| `game_changer` | 0.275 | 0.896 | -0.130 |
| `playset_share` | 0.554 | 0.237 | -0.481 |
| `singleton_share` | 0.145 | 0.140 | -0.593 |
| `source_spread` | 0.228 | 0.202 | -0.122 |

Thresholds: -3.661, -2.013, 3.528. Loss: 0.516 over 3000 iterations. Defect detector accuracy on the holdout: 0.98, at a cut of 0.46.

The precon bar per broken axis:

- lands: 0.97 of 3721
- curve: 0.98 of 3721
- colors: 0.60 of 732
- copies: 0.80 of 732
- synergy: 0.96 of 9699

Confusion on the holdout, rows are the label and columns the grade, worst first:

- bad: [304 1 0 0]
- baseline: [60 0 1 0]
- good: [6 0 817 3]
- great: [1 0 227 0]

The cards the top lists hold most: Solitude, Thoughtseize, Quantum Riddler, Force of Negation, Prismatic Ending, Fatal Push, Kozilek's Command, Psychic Frog, Malevolent Rumble, Ephemerate, Lightning Bolt, Emrakul, the Promised End.

## Verdict

Verdict: FAIL. Time: 40 seconds, no provider call. Failed bars: commander: precon over bad 0.88 of 18480 (bar 0.95); modern: precon over bad 0.95 of 18605 (bar 0.95).
