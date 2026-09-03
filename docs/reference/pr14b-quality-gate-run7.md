# PR-14B quality gate

Run date: 2026-09-02. Card snapshot: 2026-09-02. Model: 20260902T220916Z.

## Separation on the holdout

| Format | Lists | Used | Synthetic | Holdout | Great over precon | Precon over bad | Accuracy |
|---|---|---|---|---|---|---|---|
| commander | 21538 | 8388 | 1940 | 2150 | 0.96 of 20000 (bar 0.90) | 0.94 of 18480 (bar 0.95) | 0.57 |
| standard | 3316 | 3313 | 40 | 666 | 1.00 of 147 (bar 0.90) | 1.00 of 5 (bar 0.95) | 0.77 |
| modern | 8745 | 5407 | 1635 | 1420 | 1.00 of 13908 (bar 0.90) | 0.94 of 18605 (bar 0.95) | 0.79 |

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
| `card_rate` | 0.214 | 0.135 | +0.383 |
| `unseen_share` | 0.061 | 0.139 | -1.675 |
| `synergy` | 0.588 | 0.677 | +0.214 |
| `land` | 29.198 | 5.405 | -0.038 |
| `avg_mana_value` | 2.505 | 0.783 | -0.250 |
| `color_sources` | 1.252 | 0.305 | +0.042 |
| `tapped_share` | 0.074 | 0.097 | +0.255 |
| `mana_turn_four` | 4.911 | 0.597 | +0.313 |
| `hands_two_to_four_lands` | 0.627 | 0.081 | +0.120 |
| `curve_low` | 0.569 | 0.180 | +0.042 |
| `curve_high` | 0.116 | 0.121 | -0.181 |
| `ramp` | 19.746 | 7.188 | -0.025 |
| `draw` | 6.853 | 3.490 | -0.167 |
| `removal` | 5.790 | 3.391 | +0.037 |
| `wipe` | 1.218 | 1.631 | -0.164 |
| `interaction` | 8.362 | 4.658 | +0.193 |
| `empty_roles` | 0.043 | 0.211 | -0.039 |
| `fast_mana` | 7.124 | 4.753 | -0.093 |
| `game_changer` | 10.357 | 7.245 | -0.047 |
| `source_spread` | 0.102 | 0.111 | -0.431 |
| `cedh_signal` | 0.375 | 0.152 | +0.229 |
| `high_bracket_share` | 0.169 | 0.272 | +0.094 |
| `commander_decks` | 3.704 | 4.743 | -0.384 |

Thresholds: -2.837, -2.317, -1.780, 1.459. Loss: 0.761 over 3000 iterations. Defect detector accuracy on the holdout: 0.97.

The precon bar per broken axis:

- lands: 0.97 of 3696
- curve: 0.93 of 3696
- colors: 0.95 of 3036
- synergy: 0.93 of 8052

Confusion on the holdout, rows are the label and columns the grade, worst first:

- bad: [420 0 0 0 0]
- baseline: [22 0 0 22 0]
- typical: [10 0 0 30 0]
- good: [26 0 0 299 489]
- great: [23 0 0 310 499]

The cards the top lists hold most: Sol Ring, Lotus Petal, Chrome Mox, Mana Vault, Mox Diamond, Arcane Signet, Mystic Remora, Mental Misstep, Flusterstorm, Rhystic Study, Force of Will, Mox Amber.

## The standard model

Tiers, worst first: bad, baseline, good, great. Train counts: bad 35, baseline 7, good 2064, great 581. Holdout counts: bad 5, baseline 1, good 513, great 147.

Cards with a rate: 678. Pairs that lift: 4413. Commanders with a signal: 0. Top lists shape: 23.2 lands at 2.55 over 581 lists.

Features with no spread, dropped: game_changer, cedh_signal, high_bracket_share, commander_decks.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.118 | 0.052 | +0.509 |
| `unseen_share` | 0.012 | 0.101 | -0.619 |
| `synergy` | 1.720 | 0.668 | +0.313 |
| `land` | 23.159 | 2.249 | -0.003 |
| `avg_mana_value` | 2.552 | 0.575 | -0.059 |
| `color_sources` | 1.079 | 0.421 | +0.118 |
| `tapped_share` | 0.029 | 0.064 | +0.041 |
| `mana_turn_four` | 4.109 | 0.766 | -0.128 |
| `hands_two_to_four_lands` | 0.760 | 0.026 | +0.118 |
| `curve_low` | 0.629 | 0.165 | -0.142 |
| `curve_high` | 0.122 | 0.121 | -0.015 |
| `ramp` | 5.252 | 6.464 | +0.006 |
| `draw` | 4.821 | 3.326 | +0.049 |
| `removal` | 8.055 | 4.418 | +0.039 |
| `wipe` | 1.382 | 1.729 | -0.065 |
| `interaction` | 3.711 | 2.942 | +0.067 |
| `empty_roles` | 0.765 | 0.693 | -0.005 |
| `fast_mana` | 0.065 | 0.520 | -0.004 |
| `playset_share` | 0.603 | 0.170 | +0.071 |
| `singleton_share` | 0.079 | 0.062 | +0.009 |
| `source_spread` | 0.280 | 0.219 | +0.015 |

Thresholds: -5.086, -4.814, 1.418. Loss: 0.528 over 3000 iterations. Defect detector accuracy on the holdout: 0.99.

The precon bar per broken axis:

- lands: 1.00 of 1
- curve: 1.00 of 1
- colors: 1.00 of 1
- copies: 1.00 of 1
- synergy: 1.00 of 1

Confusion on the holdout, rows are the label and columns the grade, worst first:

- bad: [5 0 0 0]
- baseline: [1 0 0 0]
- good: [3 0 510 0]
- great: [0 0 147 0]

The cards the top lists hold most: Stock Up, Burst Lightning, Llanowar Elves, Opt, Sleight of Hand, Spell Pierce, Eddymurk Crab, Surrak, Elusive Hunter, Meltstrider's Resolve, Flow State, Tablet of Discovery, Badgermole Cub.

## The modern model

Tiers, worst first: bad, baseline, good, great. Train counts: bad 1330, baseline 266, good 3174, great 852. Holdout counts: bad 305, baseline 61, good 826, great 228.

Cards with a rate: 820. Pairs that lift: 6257. Commanders with a signal: 0. Top lists shape: 21.5 lands at 2.52 over 852 lists.

Features with no spread, dropped: cedh_signal, high_bracket_share, commander_decks.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.063 | 0.050 | +0.700 |
| `unseen_share` | 0.273 | 0.435 | -1.694 |
| `synergy` | 1.417 | 1.045 | +0.983 |
| `land` | 21.822 | 2.935 | +0.082 |
| `avg_mana_value` | 2.787 | 0.946 | -0.206 |
| `color_sources` | 0.988 | 0.312 | +0.004 |
| `tapped_share` | 0.116 | 0.134 | +0.062 |
| `mana_turn_four` | 3.867 | 0.808 | +0.112 |
| `hands_two_to_four_lands` | 0.737 | 0.054 | -0.001 |
| `curve_low` | 0.573 | 0.230 | +0.160 |
| `curve_high` | 0.181 | 0.154 | -0.021 |
| `ramp` | 5.299 | 6.947 | +0.058 |
| `draw` | 4.792 | 4.911 | +0.004 |
| `removal` | 7.324 | 4.367 | -0.015 |
| `wipe` | 0.936 | 1.682 | +0.075 |
| `interaction` | 2.943 | 3.778 | +0.110 |
| `empty_roles` | 1.067 | 0.849 | +0.034 |
| `fast_mana` | 0.474 | 1.403 | -0.031 |
| `game_changer` | 0.273 | 0.893 | -0.136 |
| `playset_share` | 0.554 | 0.237 | -0.466 |
| `singleton_share` | 0.145 | 0.140 | -0.544 |
| `source_spread` | 0.228 | 0.202 | -0.129 |

Thresholds: -3.712, -2.014, 3.549. Loss: 0.508 over 3000 iterations. Defect detector accuracy on the holdout: 0.97.

The precon bar per broken axis:

- lands: 0.98 of 3721
- curve: 0.99 of 3721
- colors: 0.57 of 793
- copies: 0.80 of 732
- synergy: 0.95 of 9638

Confusion on the holdout, rows are the label and columns the grade, worst first:

- bad: [304 1 0 0]
- baseline: [60 0 1 0]
- good: [8 0 817 1]
- great: [0 0 228 0]

The cards the top lists hold most: Solitude, Thoughtseize, Quantum Riddler, Force of Negation, Prismatic Ending, Fatal Push, Kozilek's Command, Psychic Frog, Malevolent Rumble, Ephemerate, Lightning Bolt, Emrakul, the Promised End.

## Verdict

Verdict: FAIL. Time: 41 seconds, no provider call. Failed bars: commander: precon over bad 0.94 of 18480 (bar 0.95); modern: precon over bad 0.94 of 18605 (bar 0.95).
