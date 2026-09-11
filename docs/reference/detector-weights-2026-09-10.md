# The defect detector's weights, 2026-09-10

This note records the weights of the defect detector before M-11 changes it. The fit never prints them, so a throwaway patch printed them. Every number here is free, and no run called a provider.

## The method

- The code is `main` at `b8582ef`, with one patch: a print of the detector's weights, its bias, and its cut after fold 0.
- Fold 0 gives the model that a `-write` run stores.
- The fit reads the meta store of gate run 19 and the card snapshot of 2026-09-04. The run wrote its documents to a scratch directory, and the patch never merged.
- A weight applies to a standardized feature. A negative weight says a low value pushes a deck toward broken. A positive weight says a high value does.

## Commander

Bias -2.032, cut 0.46.

| Feature | Weight |
|---|---|
| `mana_turn_four` | -1.1248 |
| `hands_two_to_four_lands` | -1.1122 |
| `card_rate` | -1.0799 |
| `fast_mana` | -0.8223 |
| `tapped_share` | -0.7867 |
| `synergy` | -0.6798 |
| `interaction` | -0.3169 |
| `ramp` | -0.3071 |
| `game_changer` | -0.2836 |
| `curve_low` | -0.2074 |
| `land` | -0.1393 |
| `wipe` | -0.1079 |
| `removal` | -0.1026 |
| `cedh_signal` | -0.0564 |
| `color_sources` | -0.0405 |
| `empty_roles` | -0.0062 |
| `draw` | +0.0049 |
| `commander_decks` | +0.1733 |
| `unseen_share` | +0.2439 |
| `high_bracket_share` | +0.3347 |
| `curve_high` | +0.3354 |
| `avg_mana_value` | +0.3508 |
| `source_spread` | +0.8055 |

## Standard

Bias -2.836, cut 0.38.

| Feature | Weight |
|---|---|
| `card_rate` | -1.0587 |
| `color_sources` | -1.0338 |
| `synergy` | -0.9828 |
| `hands_two_to_four_lands` | -0.5008 |
| `source_spread` | -0.3650 |
| `mana_turn_four` | -0.3391 |
| `unseen_share` | -0.2696 |
| `tapped_share` | -0.2463 |
| `curve_low` | -0.1759 |
| `wipe` | -0.1608 |
| `fast_mana` | -0.0516 |
| `removal` | -0.0325 |
| `interaction` | +0.0049 |
| `empty_roles` | +0.0809 |
| `land` | +0.0987 |
| `draw` | +0.1186 |
| `curve_high` | +0.1842 |
| `ramp` | +0.2427 |
| `avg_mana_value` | +0.5468 |
| `playset_share` | +0.5581 |
| `singleton_share` | +1.1189 |

## Modern

Bias -2.042, cut 0.46.

| Feature | Weight |
|---|---|
| `synergy` | -1.4987 |
| `tapped_share` | -0.4195 |
| `land` | -0.4093 |
| `mana_turn_four` | -0.3960 |
| `color_sources` | -0.3844 |
| `hands_two_to_four_lands` | -0.1986 |
| `empty_roles` | -0.1979 |
| `curve_low` | -0.1923 |
| `draw` | -0.1291 |
| `card_rate` | -0.0306 |
| `fast_mana` | -0.0159 |
| `interaction` | -0.0057 |
| `ramp` | +0.0325 |
| `game_changer` | +0.1345 |
| `removal` | +0.2002 |
| `curve_high` | +0.2063 |
| `unseen_share` | +0.2242 |
| `wipe` | +0.2330 |
| `source_spread` | +0.2734 |
| `avg_mana_value` | +0.5159 |
| `singleton_share` | +1.2151 |
| `playset_share` | +1.2586 |

## What the weights say

- **The detector learns the gap between the top lists and the precons** (F-106). The fit breaks the precons and the average decks alone, and the top lists stand among the negatives. `card_rate` weighs -1.08 in Commander and -1.06 in Standard. Gate run 19 reads no break that moves `card_rate` by more than 0.18 standard deviations.
- **The Commander detector reads the colors break by its signature** (F-107). `source_spread` weighs +0.81 and `tapped_share` -0.79, and `color_sources` weighs -0.04. The Standard detector reads `color_sources` at -1.03.
