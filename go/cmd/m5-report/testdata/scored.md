# PR-7 M-5 scoring sheet

Rubric version 1 (D-66). Items: 5.

---

## Item 1

Source: run10, conversation "1. lifegain", turn 2. Row `pool`, slot `pool_rule`, gap score 0.20.

**The catalog question:** Build from your library first, only your library, or ignore it?

**What the model asked instead:** Should I use your collection, or any card?

The agent sent this one in place of the catalog question.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | yes |
| invented_better | both | worse |
| right_slot | the replacement | yes |
| filled_slot | the run | yes |
| faults | the replacement | assumes an answer - drops owned-only |
| catalog_action | our catalog | none |

## Item 2

Source: run10, conversation "2. blink", turn 1. Row `locked`, slot `locked`, gap score 0.05.

**The catalog question:** Must the deck keep that card?

**What the model asked instead:** Which cards should I treat as locked in?

The agent sent this one in place of the catalog question.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | no |
| invented_better | both | better |
| right_slot | the replacement | yes |
| filled_slot | the run | yes |
| faults | the replacement | |
| catalog_action | our catalog | reword - presumes a list, try "Is there a card the deck must keep?" |

## Item 3

Source: run11, conversation "5. anything goes", turn 1. Row `house_rules`, slot `house_rules`, gap score 0.05.

**The catalog question:** What does anything-goes mean at your table?

**What the model asked instead:** Which house rules apply?

The agent sent this one in place of the catalog question.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | no |
| invented_better | both | same |
| right_slot | the replacement | yes |
| filled_slot | the run | no |
| faults | the replacement | |
| catalog_action | our catalog | add |

## Item 4

Source: run11, conversation "7. FNM", turn 2. Row `colors`, slot `colors`, gap score 0.20.

**The catalog question:** Any color preference?

**What the model offered instead:** Do you have any color preference?

The agent refused this one as a reword and sent the catalog question (D-88). Score it as if it had gone out.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | yes |
| invented_better | both | same |
| right_slot | the replacement | yes |
| filled_slot | the run | yes |
| faults | the replacement | duplicate |
| catalog_action | our catalog | none |

## Item 5

Source: run11, conversation "9. brago", turn 3. Row `commander_pick`, slot `commander`, gap score 0.20.

**The catalog question:** Which commander do you want?

**What the model asked instead:** What is your budget?

The agent sent this one in place of the catalog question.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | unsure |
| invented_better | both | worse |
| right_slot | the replacement | no |
| filled_slot | the run | no |
| faults | the replacement | wrong slot, asks about money |
| catalog_action | our catalog | none |
