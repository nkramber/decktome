# PR-14B quality gate

Run date: 2026-09-02. Card snapshot: 2026-09-02. Model: 20260902T220606Z.

## Separation on the holdout

| Format | Lists | Used | Synthetic | Holdout | Great over precon | Precon over bad | Accuracy |
|---|---|---|---|---|---|---|---|
| commander | 21538 | 8388 | 1940 | 2150 | 0.98 of 20000 (bar 0.90) | 0.83 of 18480 (bar 0.95) | 0.57 |
| standard | 3316 | 3313 | 40 | 666 | 1.00 of 147 (bar 0.90) | 1.00 of 5 (bar 0.95) | 0.77 |
| modern | 8745 | 5407 | 1635 | 1420 | 1.00 of 13908 (bar 0.90) | 0.76 of 18605 (bar 0.95) | 0.80 |

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

Features with no spread, dropped: playset_share.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.214 | 0.135 | +0.331 |
| `unseen_share` | 0.061 | 0.139 | -1.649 |
| `synergy` | 0.588 | 0.677 | +0.210 |
| `land` | 29.198 | 5.405 | -0.056 |
| `avg_mana_value` | 2.505 | 0.783 | -0.153 |
| `color_sources` | 1.252 | 0.305 | +0.250 |
| `tapped_share` | 0.074 | 0.097 | +0.298 |
| `mana_turn_four` | 4.911 | 0.597 | +0.267 |
| `hands_two_to_four_lands` | 0.627 | 0.081 | +0.046 |
| `curve_low` | 0.569 | 0.180 | +0.047 |
| `curve_high` | 0.116 | 0.121 | -0.155 |
| `ramp` | 19.746 | 7.188 | -0.134 |
| `draw` | 6.853 | 3.490 | -0.194 |
| `removal` | 5.790 | 3.391 | +0.016 |
| `wipe` | 1.218 | 1.631 | -0.189 |
| `interaction` | 8.362 | 4.658 | +0.177 |
| `empty_roles` | 0.043 | 0.211 | -0.034 |
| `fast_mana` | 7.124 | 4.753 | -0.056 |
| `game_changer` | 10.357 | 7.245 | +0.058 |
| `cedh_signal` | 0.375 | 0.152 | +0.193 |
| `high_bracket_share` | 0.169 | 0.272 | +0.201 |
| `commander_decks` | 3.704 | 4.743 | -0.442 |

Thresholds: -2.761, -2.284, -1.781, 1.417. Loss: 0.774 over 3000 iterations. Defect detector accuracy on the holdout: 0.96.

The precon bar per broken axis:

- lands: 0.81 of 3696
- curve: 0.84 of 3696
- colors: 0.72 of 3036
- synergy: 0.87 of 8052

Confusion on the holdout, rows are the label and columns the grade, worst first:

- bad: [418 0 0 2 0]
- baseline: [32 0 0 12 0]
- typical: [25 0 0 14 1]
- good: [43 0 0 282 489]
- great: [40 0 0 262 530]

The cards the top lists hold most: Sol Ring, Lotus Petal, Chrome Mox, Mana Vault, Mox Diamond, Arcane Signet, Mystic Remora, Mental Misstep, Flusterstorm, Rhystic Study, Force of Will, Mox Amber.

## The standard model

Tiers, worst first: bad, baseline, good, great. Train counts: bad 35, baseline 7, good 2064, great 581. Holdout counts: bad 5, baseline 1, good 513, great 147.

Cards with a rate: 678. Pairs that lift: 4413. Commanders with a signal: 0. Top lists shape: 23.2 lands at 2.55 over 581 lists.

Features with no spread, dropped: game_changer, cedh_signal, high_bracket_share, commander_decks.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.118 | 0.052 | +0.504 |
| `unseen_share` | 0.012 | 0.102 | -0.615 |
| `synergy` | 1.720 | 0.668 | +0.310 |
| `land` | 23.159 | 2.249 | +0.002 |
| `avg_mana_value` | 2.551 | 0.574 | -0.056 |
| `color_sources` | 1.080 | 0.419 | +0.105 |
| `tapped_share` | 0.029 | 0.064 | +0.043 |
| `mana_turn_four` | 4.109 | 0.766 | -0.122 |
| `hands_two_to_four_lands` | 0.760 | 0.026 | +0.116 |
| `curve_low` | 0.628 | 0.165 | -0.134 |
| `curve_high` | 0.122 | 0.121 | -0.001 |
| `ramp` | 5.248 | 6.465 | +0.002 |
| `draw` | 4.823 | 3.327 | +0.044 |
| `removal` | 8.057 | 4.416 | +0.039 |
| `wipe` | 1.384 | 1.729 | -0.072 |
| `interaction` | 3.709 | 2.944 | +0.068 |
| `empty_roles` | 0.766 | 0.695 | -0.012 |
| `fast_mana` | 0.065 | 0.520 | -0.005 |
| `playset_share` | 0.602 | 0.171 | +0.079 |

Thresholds: -5.086, -4.813, 1.419. Loss: 0.528 over 3000 iterations. Defect detector accuracy on the holdout: 1.00.

The precon bar per broken axis:

- lands: 1.00 of 1
- curve: 1.00 of 1
- colors: 1.00 of 1
- copies: 1.00 of 1
- synergy: 1.00 of 1

Confusion on the holdout, rows are the label and columns the grade, worst first:

- bad: [5 0 0 0]
- baseline: [1 0 0 0]
- good: [2 0 511 0]
- great: [0 0 147 0]

The cards the top lists hold most: Stock Up, Burst Lightning, Llanowar Elves, Opt, Sleight of Hand, Spell Pierce, Eddymurk Crab, Surrak, Elusive Hunter, Meltstrider's Resolve, Flow State, Tablet of Discovery, Badgermole Cub.

## The modern model

Tiers, worst first: bad, baseline, good, great. Train counts: bad 1330, baseline 266, good 3174, great 852. Holdout counts: bad 305, baseline 61, good 826, great 228.

Cards with a rate: 820. Pairs that lift: 6257. Commanders with a signal: 0. Top lists shape: 21.5 lands at 2.52 over 852 lists.

Features with no spread, dropped: cedh_signal, high_bracket_share, commander_decks.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.063 | 0.050 | +0.761 |
| `unseen_share` | 0.274 | 0.436 | -1.717 |
| `synergy` | 1.417 | 1.045 | +0.940 |
| `land` | 21.822 | 2.935 | +0.127 |
| `avg_mana_value` | 2.779 | 0.940 | -0.109 |
| `color_sources` | 0.989 | 0.310 | +0.045 |
| `tapped_share` | 0.116 | 0.134 | +0.091 |
| `mana_turn_four` | 3.869 | 0.808 | +0.047 |
| `hands_two_to_four_lands` | 0.736 | 0.054 | +0.023 |
| `curve_low` | 0.574 | 0.229 | +0.161 |
| `curve_high` | 0.180 | 0.153 | -0.093 |
| `ramp` | 5.292 | 6.954 | +0.080 |
| `draw` | 4.740 | 4.918 | +0.041 |
| `removal` | 7.289 | 4.341 | +0.032 |
| `wipe` | 0.926 | 1.676 | +0.124 |
| `interaction` | 2.921 | 3.776 | +0.156 |
| `empty_roles` | 1.079 | 0.848 | +0.021 |
| `fast_mana` | 0.471 | 1.401 | +0.003 |
| `game_changer` | 0.273 | 0.893 | -0.105 |
| `playset_share` | 0.536 | 0.263 | -0.079 |

Thresholds: -3.684, -2.053, 3.482. Loss: 0.522 over 3000 iterations. Defect detector accuracy on the holdout: 0.96.

The precon bar per broken axis:

- lands: 0.92 of 3721
- curve: 0.99 of 3721
- colors: 0.59 of 1586
- copies: 0.54 of 3721
- synergy: 0.71 of 5856

Confusion on the holdout, rows are the label and columns the grade, worst first:

- bad: [305 0 0 0]
- baseline: [61 0 0 0]
- good: [2 0 824 0]
- great: [0 0 228 0]

The cards the top lists hold most: Solitude, Thoughtseize, Quantum Riddler, Force of Negation, Prismatic Ending, Fatal Push, Kozilek's Command, Psychic Frog, Malevolent Rumble, Ephemerate, Lightning Bolt, Emrakul, the Promised End.

## Verdict

Verdict: FAIL. Time: 40 seconds, no provider call. Failed bars: commander: precon over bad 0.83 of 18480 (bar 0.95); modern: precon over bad 0.76 of 18605 (bar 0.95).
