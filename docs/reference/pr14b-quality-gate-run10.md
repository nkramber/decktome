# PR-14B quality gate

Run date: 2026-09-03. Card snapshot: 2026-09-02. Model: 20260903T010707Z.

## Separation on the holdout

| Format | Lists | Used | Synthetic | Holdout | Great over precon | Precon over bad | Accuracy |
|---|---|---|---|---|---|---|---|
| commander | 21538 | 8388 | 1940 | 2150 | 0.97 of 20000 (bar 0.90) | 0.88 of 18480 (bar 0.95) | 0.55 |
| standard | 3316 | 3313 | 40 | 666 | 0.97 of 147 (bar 0.90) | 1.00 of 5 (bar 0.95) | 0.51 |
| modern | 8745 | 5407 | 1635 | 1420 | 1.00 of 13908 (bar 0.90) | 0.95 of 18605 (bar 0.95) | 0.61 |

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
| `card_rate` | 0.214 | 0.134 | +0.858 |
| `unseen_share` | 0.034 | 0.089 | -0.171 |
| `synergy` | 0.589 | 0.676 | +0.206 |
| `land` | 29.198 | 5.405 | -0.331 |
| `avg_mana_value` | 2.507 | 0.789 | -0.389 |
| `color_sources` | 1.254 | 0.304 | +0.018 |
| `tapped_share` | 0.074 | 0.097 | +0.303 |
| `mana_turn_four` | 4.918 | 0.589 | +0.782 |
| `hands_two_to_four_lands` | 0.627 | 0.081 | +0.419 |
| `curve_low` | 0.570 | 0.179 | +0.200 |
| `curve_high` | 0.116 | 0.121 | -0.158 |
| `ramp` | 19.940 | 6.986 | -0.009 |
| `draw` | 6.954 | 3.543 | -0.262 |
| `removal` | 5.707 | 3.283 | +0.104 |
| `wipe` | 1.256 | 1.694 | -0.215 |
| `interaction` | 8.407 | 4.626 | +0.285 |
| `empty_roles` | 0.044 | 0.214 | -0.002 |
| `fast_mana` | 7.138 | 4.737 | +0.174 |
| `game_changer` | 10.365 | 7.235 | +0.152 |
| `source_spread` | 0.102 | 0.111 | -0.505 |
| `cedh_signal` | 0.375 | 0.152 | +0.111 |
| `high_bracket_share` | 0.169 | 0.272 | -0.148 |
| `commander_decks` | 3.704 | 4.743 | -0.080 |

Thresholds: -4.632, -2.852, -0.627, 1.691. Loss: 1.040 over 3000 iterations. Defect detector accuracy on the holdout: 0.94, at a cut of 0.54.

The precon bar per broken axis:

- lands: 0.98 of 3696
- curve: 0.99 of 3696
- colors: 0.98 of 3036
- synergy: 0.75 of 8052

Confusion on the holdout, rows are the label and columns the grade, worst first:

- bad: [412 0 7 1 0]
- baseline: [33 2 9 0 0]
- typical: [15 0 24 0 1]
- good: [45 5 45 299 420]
- great: [48 3 45 300 436]

The cards the top lists hold most: Sol Ring, Lotus Petal, Chrome Mox, Mana Vault, Mox Diamond, Arcane Signet, Mystic Remora, Mental Misstep, Flusterstorm, Rhystic Study, Force of Will, Mox Amber.

## The standard model

Tiers, worst first: bad, baseline, good, great. Train counts: bad 35, baseline 7, good 2064, great 581. Holdout counts: bad 5, baseline 1, good 513, great 147.

Cards with a rate: 678. Pairs that lift: 4413. Commanders with a signal: 0. Top lists shape: 23.2 lands at 2.55 over 581 lists.

Features with no spread, dropped: game_changer, cedh_signal, high_bracket_share, commander_decks.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.118 | 0.052 | +1.822 |
| `unseen_share` | 0.009 | 0.072 | +0.165 |
| `synergy` | 1.720 | 0.668 | +1.488 |
| `land` | 23.159 | 2.249 | -0.039 |
| `avg_mana_value` | 2.550 | 0.577 | -0.296 |
| `color_sources` | 1.079 | 0.420 | +0.935 |
| `tapped_share` | 0.029 | 0.064 | -0.145 |
| `mana_turn_four` | 4.109 | 0.766 | +0.007 |
| `hands_two_to_four_lands` | 0.760 | 0.026 | +0.284 |
| `curve_low` | 0.629 | 0.165 | -0.325 |
| `curve_high` | 0.122 | 0.121 | -0.113 |
| `ramp` | 5.252 | 6.467 | -0.364 |
| `draw` | 4.819 | 3.312 | +0.025 |
| `removal` | 8.059 | 4.417 | -0.016 |
| `wipe` | 1.381 | 1.729 | +0.075 |
| `interaction` | 3.717 | 2.939 | +0.293 |
| `empty_roles` | 0.765 | 0.695 | +0.090 |
| `fast_mana` | 0.066 | 0.521 | +0.152 |
| `playset_share` | 0.603 | 0.169 | -0.136 |
| `singleton_share` | 0.079 | 0.061 | -0.234 |
| `source_spread` | 0.280 | 0.219 | +0.292 |

Thresholds: -6.305, -2.801, 0.413. Loss: 0.616 over 3000 iterations. Defect detector accuracy on the holdout: 0.95, at a cut of 0.20.

The precon bar per broken axis:

- lands: 1.00 of 1
- curve: 1.00 of 1
- colors: 1.00 of 1
- copies: 1.00 of 1
- synergy: 1.00 of 1

Confusion on the holdout, rows are the label and columns the grade, worst first:

- bad: [5 0 0 0]
- baseline: [0 1 0 0]
- good: [17 11 252 233]
- great: [3 0 64 80]

The cards the top lists hold most: Stock Up, Burst Lightning, Llanowar Elves, Opt, Sleight of Hand, Spell Pierce, Eddymurk Crab, Surrak, Elusive Hunter, Meltstrider's Resolve, Flow State, Tablet of Discovery, Badgermole Cub.

## The modern model

Tiers, worst first: bad, baseline, good, great. Train counts: bad 1330, baseline 266, good 3174, great 852. Holdout counts: bad 305, baseline 61, good 826, great 228.

Cards with a rate: 820. Pairs that lift: 6257. Commanders with a signal: 0. Top lists shape: 21.5 lands at 2.52 over 852 lists.

Features with no spread, dropped: cedh_signal, high_bracket_share, commander_decks.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.063 | 0.050 | +0.851 |
| `unseen_share` | 0.263 | 0.420 | -1.573 |
| `synergy` | 1.417 | 1.045 | +1.347 |
| `land` | 21.822 | 2.935 | +0.318 |
| `avg_mana_value` | 2.779 | 0.940 | -0.330 |
| `color_sources` | 0.987 | 0.309 | +0.049 |
| `tapped_share` | 0.117 | 0.134 | +0.071 |
| `mana_turn_four` | 3.867 | 0.809 | +0.148 |
| `hands_two_to_four_lands` | 0.737 | 0.054 | -0.068 |
| `curve_low` | 0.575 | 0.228 | +0.164 |
| `curve_high` | 0.181 | 0.154 | -0.083 |
| `ramp` | 5.259 | 6.962 | +0.022 |
| `draw` | 4.734 | 4.940 | +0.058 |
| `removal` | 7.431 | 4.438 | -0.090 |
| `wipe` | 0.943 | 1.710 | +0.007 |
| `interaction` | 2.894 | 3.769 | +0.125 |
| `empty_roles` | 1.105 | 0.850 | +0.044 |
| `fast_mana` | 0.478 | 1.403 | -0.058 |
| `game_changer` | 0.275 | 0.896 | -0.221 |
| `playset_share` | 0.554 | 0.237 | -0.934 |
| `singleton_share` | 0.145 | 0.140 | -0.769 |
| `source_spread` | 0.228 | 0.202 | -0.120 |

Thresholds: -5.139, -1.272, 2.345. Loss: 0.644 over 3000 iterations. Defect detector accuracy on the holdout: 0.98, at a cut of 0.46.

The precon bar per broken axis:

- lands: 0.97 of 3721
- curve: 0.99 of 3721
- colors: 0.61 of 732
- copies: 0.79 of 732
- synergy: 0.96 of 9699

Confusion on the holdout, rows are the label and columns the grade, worst first:

- bad: [294 11 0 0]
- baseline: [26 35 0 0]
- good: [8 2 412 404]
- great: [1 0 100 127]

The cards the top lists hold most: Solitude, Thoughtseize, Quantum Riddler, Force of Negation, Prismatic Ending, Fatal Push, Kozilek's Command, Psychic Frog, Malevolent Rumble, Ephemerate, Lightning Bolt, Emrakul, the Promised End.

## Verdict

Verdict: FAIL. Time: 40 seconds, no provider call. Failed bars: commander: precon over bad 0.88 of 18480 (bar 0.95).
