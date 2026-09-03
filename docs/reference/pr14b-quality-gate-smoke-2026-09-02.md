# PR-14B quality gate

Run date: 2026-09-02. Card snapshot: 2026-09-02. Model: 20260902T194920Z.

## Separation on the holdout

| Format | Lists | Used | Synthetic | Holdout | Great over precon | Precon over bad | Accuracy |
|---|---|---|---|---|---|---|---|
| commander | 197 | 194 | 194 | 90 | no pair | 0.90 of 1980 (bar 0.95) | 0.79 |
| standard | 71 | 71 | 71 | 26 | 0.67 of 6 (bar 0.90) | 0.92 of 13 (bar 0.95) | 0.69 |
| modern | 515 | 461 | 461 | 170 | 1.00 of 183 (bar 0.90) | 0.65 of 5185 (bar 0.95) | 0.58 |

## The bracket 5 offer

Commander reads of 2026-09-02: 81 rows, 0 with Topdeck.gg entries.

| Commander | cEDH signal | Top cuts | Entries |
|---|---|---|---|
| Niv-Mizzet, Parun | 0.50 | 0 | 0 |
| Etali, Primal Conqueror // Etali, Primal Sickness | 0.50 | 0 | 0 |
| Urza, Lord High Artificer | 0.50 | 0 | 0 |

## The commander model

Tiers, worst first: bad, baseline, typical. Train counts: bad 149, baseline 145, typical 4. Holdout counts: bad 45, baseline 44, typical 1.

Cards with a rate: 0. Pairs that lift: 0. Commanders with a signal: 78. Top lists shape: 0.0 lands at 0.00 over 0 lists.

Features with no spread, dropped: card_rate, unseen_share, synergy, playset_share.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `land` | 36.950 | 4.117 | -0.003 |
| `avg_mana_value` | 3.715 | 0.506 | -0.071 |
| `color_sources` | 1.025 | 0.311 | +0.799 |
| `tapped_share` | 0.197 | 0.114 | +0.865 |
| `mana_turn_four` | 4.211 | 0.321 | +0.853 |
| `hands_two_to_four_lands` | 0.737 | 0.056 | +1.402 |
| `curve_low` | 0.275 | 0.101 | -0.315 |
| `curve_high` | 0.294 | 0.111 | -0.433 |
| `ramp` | 11.168 | 4.438 | +0.340 |
| `draw` | 7.976 | 3.210 | +0.415 |
| `removal` | 9.379 | 3.718 | +0.548 |
| `wipe` | 3.426 | 1.729 | +0.427 |
| `interaction` | 3.403 | 2.296 | +0.209 |
| `empty_roles` | 0.070 | 0.256 | +0.042 |
| `fast_mana` | 1.071 | 0.699 | +0.592 |
| `game_changer` | 0.178 | 0.623 | +0.423 |
| `cedh_signal` | 0.013 | 0.081 | +0.223 |
| `high_bracket_share` | 0.021 | 0.081 | +0.006 |
| `commander_decks` | 0.724 | 2.699 | -0.037 |

Thresholds: 0.197, 6.440. Loss: 0.415 over 3000 iterations.

Confusion on the holdout, rows are the label and columns the grade, worst first:

- bad: [40 5 0]
- baseline: [13 31 0]
- typical: [0 1 0]

## The standard model

Tiers, worst first: bad, baseline, good, great. Train counts: bad 58, baseline 7, good 41, great 10. Holdout counts: bad 13, baseline 1, good 6, great 6.

Cards with a rate: 180. Pairs that lift: 259. Commanders with a signal: 0. Top lists shape: 22.8 lands at 2.74 over 10 lists.

Features with no spread, dropped: game_changer, cedh_signal, high_bracket_share, commander_decks.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.102 | 0.065 | -0.041 |
| `unseen_share` | 0.225 | 0.297 | -0.928 |
| `synergy` | 0.400 | 0.462 | +1.040 |
| `land` | 22.629 | 3.255 | +0.237 |
| `avg_mana_value` | 2.997 | 0.909 | -0.614 |
| `color_sources` | 0.996 | 0.469 | +0.834 |
| `tapped_share` | 0.048 | 0.069 | +0.098 |
| `mana_turn_four` | 3.814 | 0.681 | -0.249 |
| `hands_two_to_four_lands` | 0.743 | 0.055 | +0.985 |
| `curve_low` | 0.534 | 0.201 | +0.169 |
| `curve_high` | 0.199 | 0.193 | -0.151 |
| `ramp` | 3.828 | 5.608 | -0.408 |
| `draw` | 5.957 | 3.807 | -0.034 |
| `removal` | 8.897 | 5.001 | -0.296 |
| `wipe` | 1.681 | 1.919 | +0.155 |
| `interaction` | 2.897 | 2.952 | +0.229 |
| `empty_roles` | 0.914 | 0.726 | +0.420 |
| `fast_mana` | 0.069 | 0.521 | +0.244 |
| `playset_share` | 0.506 | 0.259 | -0.063 |

Thresholds: 0.062, 0.585, 3.840. Loss: 0.723 over 3000 iterations.

Confusion on the holdout, rows are the label and columns the grade, worst first:

- bad: [13 0 0 0]
- baseline: [1 0 0 0]
- good: [1 0 5 0]
- great: [3 0 3 0]

The cards the top lists hold most: Requiting Hex, Spell Pierce, Shoot the Sheriff, Spell Snare, Bitter Triumph, Burst Lightning, Eddymurk Crab, Get Out, Hearth Elemental // Stoke Genius, Opt, Prismari Charm, Sleight of Hand.

## The modern model

Tiers, worst first: bad, baseline, good, great. Train counts: bad 376, baseline 266, good 89, great 21. Holdout counts: bad 85, baseline 61, good 21, great 3.

Cards with a rate: 287. Pairs that lift: 638. Commanders with a signal: 0. Top lists shape: 21.2 lands at 2.74 over 21 lists.

Features with no spread, dropped: cedh_signal, high_bracket_share, commander_decks.

| Feature | Mean | Std | Weight |
|---|---|---|---|
| `card_rate` | 0.027 | 0.049 | +0.069 |
| `unseen_share` | 0.745 | 0.395 | -1.129 |
| `synergy` | 0.213 | 0.505 | +0.538 |
| `land` | 22.725 | 3.249 | +0.692 |
| `avg_mana_value` | 3.119 | 0.866 | -0.310 |
| `color_sources` | 0.915 | 0.390 | +0.345 |
| `tapped_share` | 0.070 | 0.113 | +0.194 |
| `mana_turn_four` | 3.745 | 0.536 | +0.340 |
| `hands_two_to_four_lands` | 0.743 | 0.066 | +0.522 |
| `curve_low` | 0.448 | 0.209 | +0.455 |
| `curve_high` | 0.211 | 0.162 | -0.182 |
| `ramp` | 2.894 | 4.615 | -0.234 |
| `draw` | 3.540 | 3.992 | +0.199 |
| `removal` | 7.653 | 4.064 | +0.031 |
| `wipe` | 0.713 | 1.488 | -0.248 |
| `interaction` | 1.874 | 2.778 | +0.067 |
| `empty_roles` | 1.225 | 0.836 | +0.403 |
| `fast_mana` | 0.194 | 0.863 | +0.122 |
| `game_changer` | 0.081 | 0.484 | -0.004 |
| `playset_share` | 0.314 | 0.302 | -0.519 |

Thresholds: -0.070, 2.752, 5.445. Loss: 0.752 over 3000 iterations.

Confusion on the holdout, rows are the label and columns the grade, worst first:

- bad: [54 29 2 0]
- baseline: [27 34 0 0]
- good: [2 6 11 2]
- great: [0 1 2 0]

The cards the top lists hold most: Solitude, Thoughtseize, Quantum Riddler, Force of Negation, Psychic Frog, Prismatic Ending, Ephemerate, Lightning Bolt, Fatal Push, Kozilek's Command, Griselbrand, Mishra's Bauble.

## Verdict

Verdict: FAIL. Time: 3 seconds, no provider call. Failed bars: commander: great over precon no pair; commander: precon over bad 0.90 of 1980 (bar 0.95); standard: great over precon 0.67 of 6 (bar 0.90); standard: precon over bad 0.92 of 13 (bar 0.95); modern: precon over bad 0.65 of 5185 (bar 0.95); offer: no Topdeck.gg read, so no commander comes from a top cut (TOPDECK_API_KEY).
