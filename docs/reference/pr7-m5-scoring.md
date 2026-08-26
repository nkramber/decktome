# PR-7 M-5 scoring sheet

Rubric version 1 (D-66). Items: 60, from 793 questions across the gate runs.

A question the model left alone is not scored. Only a question it offered to replace is.

26 went out in place of a catalog question. 34 were refused as rewords and never reached a user (D-88). Both kinds are scored the same way. The refused ones decide whether the reword guard is set right.

## What this sheet decides

The agent scores every catalog question from 0 to 1 for how well it fits the user. Under a threshold, the model may throw the row away and write its own question. That threshold is a guess until you score these items. Your answers replace the guess (D-27).

## How to read an item

Each item holds two questions. **The catalog question** is the fixed row the agent planned to ask. **What the model asked instead** is the replacement it wrote. Every field below reads one of the two, and the list says which.

## The six fields

### `catalog_enough` (reads the catalog question)

Values: yes / no / unsure. Ignore the replacement. Read the catalog question alone. Had the agent simply asked that, would this conversation have been fine?

Use `unsure` freely. An unsure item is reported and left out of the calculation (D-66).

### `invented_better` (compares the two)

Values: worse / same / better. The only field that compares.

- `worse`: the replacement loses something the row had.
- `same`: no real difference. A replacement that only adds the user's colors or format is `same`.
- `better`: it serves this user better.

### `right_slot` (reads the replacement)

Values: yes / no. Does the replacement ask about the slot named in the item header?

This is a break, not a matter of taste. A question filed under `theme` that asks about money sends the answer to the wrong field.

### `filled_slot` (already filled in)

Values: yes / partly / no. The run recorded it. Change it only if you disagree.

It means one of two things. On an item the agent sent, it says whether the replacement closed the slot. On a refused item the replacement never went out, so it says whether the catalog question closed the slot.

### `faults` (reads the replacement)

A list, not one value. Leave it blank only when nothing is wrong. Mandatory otherwise.

| Fault | Means |
|---|---|
| duplicate | Asks what the user already answered |
| two questions in one | The user can not answer it in one reply |
| jargon | Uses words a player may not know |
| assumes an answer | Presumes something the user never said |
| unanswerable | Asks for a thing the user can not know |
| out of scope | Asks about a thing the app can not act on |

### `catalog_action` (reads our catalog)

Values: none / add / reword. Not about the replacement. This field is how a scored item becomes a catalog change (D-25).

- `none`: leave the row alone.
- `reword`: the right row exists, and its wording caused the problem.
- `add`: no row covers this need at all.

## Free text

`catalog_action` and `faults` accept free text after the keyword, and a proposed wording is the useful part of `reword`. Put the keyword first, then a dash, then whatever you want:

    | catalog_action | reword - presumes a list, try "Is there a card the deck must keep?" |

Keep `catalog_enough` and `invented_better` to their listed values. Those two are counted, and free text there would measure the reader rather than you (D-66).

## How your answers are used

An invention is warranted when `catalog_enough` is no and `invented_better` is not worse. The D-27 threshold is the gap score that best separates warranted from unwarranted. At least 80% of the inventions a threshold allows must be warranted.

Every item here scored 0.05 or 0.20, because the scorer counts faults rather than shades of fit. A 0.05 means it found two faults or more, and a 0.20 means one clear fault. So the threshold question is nearly a yes or no: does the 0.20 group deserve to invent?

Some items repeat. That is on purpose, and it measures self-consistency. Do not look for them.

---

## Item 1

Source: pr7-question-gate-run10, conversation "29. two hundred dollars", turn 2. Row `budget_scope`, slot `budget`, gap score 0.05.

**The catalog question:** Is your budget a cap on the cards you buy, or on the whole deck's total value?

**What the model offered instead:** What is the $200 budget a cap on: the cards you buy or the whole deck value?

The agent refused this one as a reword and sent the catalog question (D-88). Score it as if it had gone out. A `better` here says the refusal was wrong.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | yes |
| invented_better | both | same |
| right_slot | the replacement | yes |
| filled_slot | the run | no |
| faults | the replacement | |
| catalog_action | our catalog | none |

## Item 2

Source: pr7-question-gate-run11, conversation "44. sideboard help only (probe)", turn 1. Row `colors`, slot `colors`, gap score 0.20.

**The catalog question:** Any color preference? A burn deck is strongest in red.

**What the model asked instead:** Should the sideboard use any colors, or red only?

The agent sent this one in place of the catalog question.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | yes |
| invented_better | both | same |
| right_slot | the replacement | yes |
| filled_slot | the run | yes |
| faults | the replacement | |
| catalog_action | our catalog | none |

## Item 3

Source: pr7-question-gate-run11, conversation "23. land destruction", turn 2. Row `commander`, slot `commander`, gap score 0.20.

**The catalog question:** Do you have a red-green commander in mind, or should I suggest one?

**What the model offered instead:** Do you have a commander in mind, or should I suggest one?

The agent refused this one as a reword and sent the catalog question (D-88). Score it as if it had gone out. A `better` here says the refusal was wrong.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | yes |
| invented_better | both | same |
| right_slot | the replacement | yes |
| filled_slot | the run | no |
| faults | the replacement | |
| catalog_action | our catalog | none |

## Item 4

Source: pr7-question-gate-run10, conversation "23. land destruction", turn 4. Row `commander_pick`, slot `commander`, gap score 0.20.

**The catalog question:** Which commander do you want: The Earth King, Toph, Hardheaded Teacher, or Toph, the Blind Bandit? Say “none” and I’ll name three more.

**What the model offered instead:** Which one do you want: The Earth King, Toph, Hardheaded Teacher, or Toph, the Blind Bandit? Say "none" and I name three more.

The agent refused this one as a reword and sent the catalog question (D-88). Score it as if it had gone out. A `better` here says the refusal was wrong.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | yes |
| invented_better | both | same |
| right_slot | the replacement | yes |
| filled_slot | the run | no |
| faults | the replacement | |
| catalog_action | our catalog | none |

## Item 5

Source: pr7-question-gate-run11, conversation "37. a request we can not serve (probe)", turn 1. Row `format`, slot `format`, gap score 0.05.

**The catalog question:** Which format: Commander, Standard, Modern, or something else?

**What the model asked instead:** Which Yu-Gi-Oh format would you like?

The agent sent this one in place of the catalog question.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | no |
| invented_better | both | worse |
| right_slot | the replacement | no |
| filled_slot | the run | yes |
| faults | the replacement | out of scope |
| catalog_action | our catalog | none - the question is valid, but the user's query should not lead to this catalog question. this has been fixed since run 11 ran. |

## Item 6

Source: pr7-question-gate-run11, conversation "21. vintage with proxies", turn 1. Row `house_rules`, slot `house_rules`, gap score 0.20.

**The catalog question:** When you say you proxy everything, do you mean any card with no ban list, or using Vintage rules?

**What the model offered instead:** What does your table's anything-goes rule mean: any card with no ban list, or Vintage rules?

The agent refused this one as a reword and sent the catalog question (D-88). Score it as if it had gone out. A `better` here says the refusal was wrong.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | no |
| invented_better | both | same |
| right_slot | the replacement | no |
| filled_slot | the run | yes |
| faults | the replacement | |
| catalog_action | our catalog | reword - 'proxy' means the user has no budget. they will use fake versions of each card, so budget is of no concern. we need to handle 'proxy' better. |

## Item 7

Source: pr7-question-gate-run10, conversation "13. a card to keep that is not the commander", turn 1. Row `locked`, slot `locked`, gap score 0.20.

**The catalog question:** Must the deck keep Sanguine Bond, or may I cut a card that does not fit the plan?

**What the model asked instead:** Besides keeping Sanguine Bond, should the deck keep every other card you named, or may I cut cards that do not fit?

The agent sent this one in place of the catalog question.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | no |
| invented_better | both | better |
| right_slot | the replacement | yes |
| filled_slot | the run | no |
| faults | the replacement | |
| catalog_action | our catalog | reword - the question we asked did not match the user prompt at all. the user wants to specifically keep sanguine bond, but is okay with reworking everything else except the commander. |

## Item 8

Source: pr7-question-gate-run12, conversation "33. the user contradicts the budget (probe)", turn 1. Row `meta`, slot `meta`, gap score 0.20.

**The catalog question:** What do people play at your event? I tune the 15 sideboard cards to it.

**What the model asked instead:** What do people play at your Modern event?

The agent sent this one in place of the catalog question.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | yes |
| invented_better | both | worse |
| right_slot | the replacement | yes |
| filled_slot | the run | no |
| faults | the replacement | vague |
| catalog_action | our catalog | none |

## Item 9

Source: pr7-question-gate-run11, conversation "27. a card for the 99", turn 2. Row `named_card_role`, slot `commander`, gap score 0.20.

**The catalog question:** Would you like Grist, the Hunger Tide as your commander, or as one card in the 99?

**What the model offered instead:** Do you want Grist, the Hunger Tide as one card in the 99?

The agent refused this one as a reword and sent the catalog question (D-88). Score it as if it had gone out. A `better` here says the refusal was wrong.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | no |
| invented_better | both | same |
| right_slot | the replacement | yes |
| filled_slot | the run | yes |
| faults | the replacement | |
| catalog_action | our catalog | reword - the user clearly specified that they do NOT want Grist as their commander |

## Item 10

Source: pr7-question-gate-run10, conversation "29. two hundred dollars", turn 2. Row `budget_scope`, slot `budget`, gap score 0.05.

**The catalog question:** Is your budget a cap on the cards you buy, or on the whole deck's total value?

**What the model offered instead:** What is the $200 budget a cap on: the cards you buy or the whole deck value?

The agent refused this one as a reword and sent the catalog question (D-88). Score it as if it had gone out. A `better` here says the refusal was wrong.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | yes |
| invented_better | both | same |
| right_slot | the replacement | yes |
| filled_slot | the run | no |
| faults | the replacement | |
| catalog_action | our catalog | none |

## Item 11

Source: pr7-question-gate-run10, conversation "4. fun and janky", turn 3. Row `pool`, slot `pool_rule`, gap score 0.20.

**The catalog question:** For your sacrifice deck, should I build from your library first, use only cards from your library, or ignore your library for a fully optimized deck?

**What the model offered instead:** Build from your library first, only your library, or use any cards for a fully optimized deck?

The agent refused this one as a reword and sent the catalog question (D-88). Score it as if it had gone out. A `better` here says the refusal was wrong.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | yes |
| invented_better | both | better |
| right_slot | the replacement | yes |
| filled_slot | the run | yes |
| faults | the replacement | |
| catalog_action | our catalog | reword - the LLM answer is less verbose and of higher quality |

## Item 12

Source: pr7-question-gate-run10, conversation "24. stax, owned only", turn 2. Row `power_commander`, slot `power`, gap score 0.05.

**The catalog question:** Which power bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.

**What the model offered instead:** Which bracket does your table play: 2 (precon level), 3 (upgraded), or 4 (high power)?

The agent refused this one as a reword and sent the catalog question (D-88). Score it as if it had gone out. A `better` here says the refusal was wrong.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | yes |
| invented_better | both | same |
| right_slot | the replacement | yes |
| filled_slot | the run | yes |
| faults | the replacement | |
| catalog_action | our catalog | none |

## Item 13

Source: pr7-question-gate-run12, conversation "33. the user contradicts the budget (probe)", turn 1. Row `power_sixty`, slot `power`, gap score 0.05.

**The catalog question:** How strong should your competitive Modern deck be: casual, FNM level, or tournament-meta?

**What the model offered instead:** How strong should this be: casual, FNM level, or tournament-meta?

The agent refused this one as a reword and sent the catalog question (D-88). Score it as if it had gone out. A `better` here says the refusal was wrong.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | yes |
| invented_better | both | better |
| right_slot | the replacement | yes |
| filled_slot | the run | yes |
| faults | the replacement | |
| catalog_action | our catalog | none |

## Item 14

Source: pr7-question-gate-run10, conversation "21. vintage with proxies", turn 3. Row `table_tolerance`, slot `power`, gap score 0.20.

**The catalog question:** Does your table accept mill, land destruction, extra turns, or stax?

**What the model asked instead:** What power level should the deck target: casual or high power?

The agent sent this one in place of the catalog question.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | n/a |
| invented_better | both | n/a |
| right_slot | the replacement | n/a |
| filled_slot | the run | no |
| faults | the replacement | |
| catalog_action | our catalog | 'table_tolerance' needs to be removed entirely. we should not ask the user if their table is okay with things like mill, land destruction, etc. 'power_level' is enough. |

## Item 15

Source: pr7-question-gate-run11, conversation "38. shorthand and slang (probe)", turn 3. Row `budget_scope`, slot `budget`, gap score 0.20.

**The catalog question:** Is that a cap on the cards you buy, or on the whole deck value?

**What the model asked instead:** Should the $50 cap apply to the cards you buy or to the whole deck’s value?

The agent sent this one in place of the catalog question.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | yes |
| invented_better | both | better |
| right_slot | the replacement | yes |
| filled_slot | the run | no |
| faults | the replacement | |
| catalog_action | our catalog | reword |

## Item 16

Source: pr7-question-gate-run12, conversation "44. sideboard help only (probe)", turn 1. Row `colors`, slot `colors`, gap score 0.20.

**The catalog question:** Any color preference? A burn deck is strongest in red.

**What the model asked instead:** Should the sideboard use red only, or are other colors acceptable?

The agent sent this one in place of the catalog question.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | no |
| invented_better | both | better |
| right_slot | the replacement | yes |
| filled_slot | the run | yes |
| faults | the replacement | |
| catalog_action | our catalog | reword |

## Item 17

Source: pr7-question-gate-run12, conversation "1. lifegain with a collection", turn 2. Row `commander`, slot `commander`, gap score 0.05.

**The catalog question:** Do you have a commander in mind, or should I suggest one for your white-black lifegain deck?

**What the model offered instead:** Do you have a commander in mind, or should I suggest one?

The agent refused this one as a reword and sent the catalog question (D-88). Score it as if it had gone out. A `better` here says the refusal was wrong.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | yes |
| invented_better | both | better |
| right_slot | the replacement | yes |
| filled_slot | the run | yes |
| faults | the replacement | |
| catalog_action | our catalog | reword |

## Item 18

Source: pr7-question-gate-run11, conversation "22. extra turns commander", turn 3. Row `commander_pick`, slot `commander`, gap score 0.05.

**The catalog question:** Which commander do you want: Birgi, God of Storytelling // Harnfel, Horn of Bounty, Emrakul, the Promised End, or Vnwxt, Verbose Host? Choose “None—name three more” and I’ll suggest three additional options.

**What the model offered instead:** Which commander do you want: Birgi, God of Storytelling // Harnfel, Horn of Bounty, Emrakul, the Promised End, or Vnwxt, Verbose Host?

The agent refused this one as a reword and sent the catalog question (D-88). Score it as if it had gone out. A `better` here says the refusal was wrong.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | yes |
| invented_better | both | same |
| right_slot | the replacement | yes |
| filled_slot | the run | no |
| faults | the replacement | |
| catalog_action | our catalog | none |

## Item 19

Source: pr7-question-gate-run11, conversation "50. two decks at once (probe)", turn 1. Row `format`, slot `format`, gap score 0.20.

**The catalog question:** For which deck are you choosing a format: Commander, Standard, Modern, Pioneer, or Pauper?

**What the model offered instead:** Which format will the first deck use: Commander, Standard, Modern, or something else?

The agent refused this one as a reword and sent the catalog question (D-88). Score it as if it had gone out. A `better` here says the refusal was wrong.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | no |
| invented_better | both | better |
| right_slot | the replacement | yes |
| filled_slot | the run | yes |
| faults | the replacement | |
| catalog_action | our catalog | reword - the catalog question is just wrong. the replacement is better, but still not perfect. we need to specify that the tool is for ONE deck at a time. |

## Item 20

Source: pr7-question-gate-run11, conversation "44. sideboard help only (probe)", turn 1. Row `colors`, slot `colors`, gap score 0.20.

**The catalog question:** Any color preference? A burn deck is strongest in red.

**What the model asked instead:** Should the sideboard use any colors, or red only?

The agent sent this one in place of the catalog question.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | no |
| invented_better | both | better |
| right_slot | the replacement | yes |
| filled_slot | the run | yes |
| faults | the replacement | |
| catalog_action | our catalog | reword |

## Item 21

Source: pr7-question-gate-run11, conversation "38. shorthand and slang (probe)", turn 1. Row `house_rules`, slot `house_rules`, gap score 0.05.

**The catalog question:** What does anything-goes mean at your table: any card with no ban list, or Vintage rules?

**What the model asked instead:** What house rules does your table use for your Gruul dinosaur EDH deck, if any?

The agent sent this one in place of the catalog question.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | no |
| invented_better | both | better |
| right_slot | the replacement | yes |
| filled_slot | the run | no |
| faults | the replacement | |
| catalog_action | our catalog | reword - the catalog question is irrelevant here. the invented question is relevant. |

## Item 22

Source: pr7-question-gate-run10, conversation "27. a card for the 99", turn 3. Row `locked`, slot `locked`, gap score 0.20.

**The catalog question:** Must the deck keep Grist, the Hunger Tide and Grist, or may I cut a card that does not fit the sacrifice plan?

**What the model offered instead:** Must the deck keep Grist, the Hunger Tide, or may I cut a card that does not fit the plan?

The agent refused this one as a reword and sent the catalog question (D-88). Score it as if it had gone out. A `better` here says the refusal was wrong.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | no |
| invented_better | both | better |
| right_slot | the replacement | yes |
| filled_slot | the run | no |
| faults | the replacement | |
| catalog_action | our catalog | reword - the catalog question has a misplaced work - 'Grist' appears twice. the invented question fixed this. |

## Item 23

Source: pr7-question-gate-run13, conversation "27. a card for the 99", turn 2. Row `named_card_role`, slot `commander`, gap score 0.20.

**The catalog question:** Do you want Grist, the Hunger Tide as your commander, or as one card in the 99?

**What the model asked instead:** Do you want Grist, the Hunger Tide as your commander, or in the 99 with a different commander?

The agent sent this one in place of the catalog question.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | no |
| invented_better | both | same |
| right_slot | the replacement | yes |
| filled_slot | the run | no |
| faults | the replacement | duplicate |
| catalog_action | our catalog | reword - both the catalog and invented questions are wrong. the user's initial prompt is "Build around Grist, the Hunger Tide, but not as my commander" - this implies two things: we are building a commander deck, Grist, the Hunger Tide MUST be included, and CANNOT be the commander. our first follow up is "Which format would you like: Commander, Standard, Modern, Pioneer, or something else?" - this is a redundant question as we already know that we are building a commander deck. the user then says "Commander. A sacrifice deck, black and green." - at this point the user must be slightly frustrated, since 'commander' was already obvious. our next question is "Do you want Grist, the Hunger Tide as your commander, or in the 99 with a different commander?" which is also redundant - the user already said they do NOT want Grist, the Hunger Tide as their commander, and it is evident and obvious that they want this card in their deck. this needs a significant fix. |

## Item 24

Source: pr7-question-gate-run10, conversation "17. precon upgrade at bracket 2", turn 1. Row `pool`, slot `pool_rule`, gap score 0.20.

**The catalog question:** Should I build from your precon cards first, use only cards from the precon, or ignore the precon for a fully optimized deck?

**What the model offered instead:** Build from your library first, use only cards from your library, or ignore it for a fully optimized deck?

The agent refused this one as a reword and sent the catalog question (D-88). Score it as if it had gone out. A `better` here says the refusal was wrong.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | no |
| invented_better | both | better |
| right_slot | the replacement | yes |
| filled_slot | the run | yes |
| faults | the replacement | |
| catalog_action | our catalog | reword - it needs to be obvious that the user wants most of their precon cards to remain the same. "Upgrade my precon" means utilize most of the existing cards. the invented question is more on-track, asking if they should utilize just the user's library or not. either way, we need to ensure that "upgrade my precon" leads to reusing the same theme, colors, etc as the precon. |

## Item 25

Source: pr7-question-gate-run12, conversation "43. a deck as a gift (probe)", turn 2. Row `power_commander`, slot `power`, gap score 0.20.

**The catalog question:** Which bracket does your table play? 2 is precon level, 3 is upgraded, 4 is high power.

**What the model asked instead:** What power level should the black Zombie deck target: 2 (precon level), 3 (upgraded), or 4 (high power)?

The agent sent this one in place of the catalog question.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | no |
| invented_better | both | better |
| right_slot | the replacement | yes |
| filled_slot | the run | yes |
| faults | the replacement | |
| catalog_action | our catalog | reword - the catalog question is incorrect as the user never mentioned a 'table' at all. |

## Item 26

Source: pr7-question-gate-run13, conversation "16. the strongest modern deck", turn 1. Row `power_sixty`, slot `power`, gap score 0.05.

**The catalog question:** How strong should your Modern deck be: casual, FNM level, or tournament-meta?

**What the model offered instead:** How strong should this be: casual, FNM level, or tournament-meta?

The agent refused this one as a reword and sent the catalog question (D-88). Score it as if it had gone out. A `better` here says the refusal was wrong.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | no |
| invented_better | both | same |
| right_slot | the replacement | no |
| filled_slot | the run | yes |
| faults | the replacement | duplicate |
| catalog_action | our catalog | reword - the user says in their initial prompt "I want the strongest Modern deck, money is no object" - it is redundant and not useful for us to ask them any variation of 'how strong should this be' as we already know that information. |

## Item 27

Source: pr7-question-gate-run10, conversation "22. extra turns commander", turn 3. Row `table_tolerance`, slot `power`, gap score 0.20.

**The catalog question:** Does your table accept mill, land destruction, extra turns, or stax?

**What the model offered instead:** What power level does your table accept: mill, land destruction, extra turns, or stax?

The agent refused this one as a reword and sent the catalog question (D-88). Score it as if it had gone out. A `better` here says the refusal was wrong.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | n/a |
| invented_better | both | n/a |
| right_slot | the replacement | n/a |
| filled_slot | the run | no |
| faults | the replacement | |
| catalog_action | our catalog | 'table_tolerance' needs to be removed entirely. we should not ask the user if their table is okay with things like mill, land destruction, etc. 'power_level' is enough. |

## Item 28

Source: pr7-question-gate-run13, conversation "5. 60-card anything goes", turn 2. Row `colors`, slot `colors`, gap score 0.20.

**The catalog question:** Any color preference? A a dragon deck deck is strongest in blue and red.

**What the model asked instead:** Which colors would you like for your dragon deck?

The agent sent this one in place of the catalog question.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | yes |
| invented_better | both | better |
| right_slot | the replacement | yes |
| filled_slot | the run | yes |
| faults | the replacement | |
| catalog_action | our catalog | reword - there is a type in the original question (duplicated 'deck' word), and presuming which colors are strongest is a slipperly slope. getting that wrong loses user trust. |

## Item 29

Source: pr7-question-gate-run13, conversation "27. a card for the 99", turn 3. Row `commander`, slot `commander`, gap score 0.20.

**The catalog question:** Do you have a commander in mind, or should I suggest one for this Grist, the Hunger Tide deck?

**What the model offered instead:** Do you have a different commander in mind, or should I suggest one?

The agent refused this one as a reword and sent the catalog question (D-88). Score it as if it had gone out. A `better` here says the refusal was wrong.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | yes |
| invented_better | both | worse |
| right_slot | the replacement | yes |
| filled_slot | the run | yes |
| faults | the replacement | inaccurate - 'different commander' assumes they suggested ANY commander, which they did not. |
| catalog_action | our catalog | none - while this question was fine, turns 1 and 2 of this conversation were a disaster - see 'Item 23' above. |

## Item 30

Source: pr7-question-gate-run11, conversation "23. land destruction", turn 2. Row `commander`, slot `commander`, gap score 0.20.

**The catalog question:** Do you have a red-green commander in mind, or should I suggest one?

**What the model offered instead:** Do you have a commander in mind, or should I suggest one?

The agent refused this one as a reword and sent the catalog question (D-88). Score it as if it had gone out. A `better` here says the refusal was wrong.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | yes |
| invented_better | both | worse |
| right_slot | the replacement | yes |
| filled_slot | the run | no |
| faults | the replacement | omits information from catalog question |
| catalog_action | our catalog | none - however, in turn 1 we ask what format when the user CLEARLY stated 'Commander' in their initial prompt. |

## Item 31

Source: pr7-question-gate-run12, conversation "24. stax, owned only", turn 3. Row `commander_pick`, slot `commander`, gap score 0.20.

**The catalog question:** Which commander do you want: Thalia, Guardian of Thraben, Pippin, Guard of the Citadel, or Kataki, War's Wage? You can also say “none” and I’ll name three more.

**What the model offered instead:** Which one do you want: Thalia, Guardian of Thraben, Pippin, Guard of the Citadel, or Kataki, War's Wage? Say "none" and I name three more.

The agent refused this one as a reword and sent the catalog question (D-88). Score it as if it had gone out. A `better` here says the refusal was wrong.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | yes |
| invented_better | both | worse |
| right_slot | the replacement | yes |
| filled_slot | the run | no |
| faults | the replacement | omits information from catalog question |
| catalog_action | our catalog | none |

## Item 32

Source: pr7-question-gate-run12, conversation "5. 60-card anything goes", turn 1. Row `format`, slot `format`, gap score 0.20.

**The catalog question:** Which format: Commander, Standard, Modern, or something else?

**What the model asked instead:** Which format should the deck use, such as Commander, Standard, Modern, Pioneer, or Pauper?

The agent sent this one in place of the catalog question.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | yes |
| invented_better | both | same |
| right_slot | the replacement | yes |
| filled_slot | the run | yes |
| faults | the replacement | |
| catalog_action | our catalog | none - however, in turn 1 we ask "When you say anything goes, do you mean any card with no ban list or Vintage rules?" which is a good question, and turn away the invented question "What house rules should apply at your table: any card with no ban list, or Vintage rules?" as a reword. we should turn this away for a different reason - it's just an inaccurate question. they said "anything goes", so there is no reason to ask "What house rules should apply" - that verbage is wrong. |

## Item 33

Source: pr7-question-gate-run12, conversation "5. 60-card anything goes", turn 1. Row `house_rules`, slot `house_rules`, gap score 0.20.

**The catalog question:** When you say anything goes, do you mean any card with no ban list or Vintage rules?

**What the model offered instead:** What house rules should apply at your table: any card with no ban list, or Vintage rules?

The agent refused this one as a reword and sent the catalog question (D-88). Score it as if it had gone out. A `better` here says the refusal was wrong.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | |
| invented_better | both | |
| right_slot | the replacement | |
| filled_slot | the run | yes |
| faults | the replacement | |
| catalog_action | our catalog | |

## Item 34

Source: pr7-question-gate-run11, conversation "27. a card for the 99", turn 3. Row `locked`, slot `locked`, gap score 0.20.

**The catalog question:** Must the deck keep Grist, the Hunger Tide and Grist, or may I cut a card that does not fit the plan?

**What the model asked instead:** Should Grist, the Hunger Tide stay in the 99, or may I cut another card that does not fit the sacrifice plan?

The agent sent this one in place of the catalog question.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | |
| invented_better | both | |
| right_slot | the replacement | |
| filled_slot | the run | no |
| faults | the replacement | |
| catalog_action | our catalog | |

## Item 35

Source: pr7-question-gate-run11, conversation "17. precon upgrade at bracket 2", turn 1. Row `pool`, slot `pool_rule`, gap score 0.20.

**The catalog question:** Build from your library first, only your library, or ignore it for a fully optimized deck?

**What the model asked instead:** Would you like the deck built using your existing library first, only from your existing library, or with no library restriction for a fully optimized deck?

The agent sent this one in place of the catalog question.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | |
| invented_better | both | |
| right_slot | the replacement | |
| filled_slot | the run | yes |
| faults | the replacement | |
| catalog_action | our catalog | |

## Item 36

Source: pr7-question-gate-run13, conversation "4. fun and janky", turn 2. Row `power_commander`, slot `power`, gap score 0.20.

**The catalog question:** Which bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.

**What the model offered instead:** What power level does your table play: bracket 2 (precon level), bracket 3 (upgraded), or bracket 4 (high power)?

The agent refused this one as a reword and sent the catalog question (D-88). Score it as if it had gone out. A `better` here says the refusal was wrong.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | |
| invented_better | both | |
| right_slot | the replacement | |
| filled_slot | the run | yes |
| faults | the replacement | |
| catalog_action | our catalog | |

## Item 37

Source: pr7-question-gate-run13, conversation "44. sideboard help only (probe)", turn 1. Row `power_sixty`, slot `power`, gap score 0.20.

**The catalog question:** How strong should your Modern Burn sideboard be: casual, FNM-level, or tournament meta?

**What the model offered instead:** How strong should your sideboard be: casual, FNM level, or tournament-meta?

The agent refused this one as a reword and sent the catalog question (D-88). Score it as if it had gone out. A `better` here says the refusal was wrong.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | |
| invented_better | both | |
| right_slot | the replacement | |
| filled_slot | the run | yes |
| faults | the replacement | |
| catalog_action | our catalog | |

## Item 38

Source: pr7-question-gate-run11, conversation "24. stax, owned only", turn 3. Row `table_tolerance`, slot `power`, gap score 0.20.

**The catalog question:** Does your table accept mill, land destruction, extra turns, or stax?

**What the model offered instead:** Does your table accept mill, land destruction, or extra turns, in addition to stax?

The agent refused this one as a reword and sent the catalog question (D-88). Score it as if it had gone out. A `better` here says the refusal was wrong.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | n/a |
| invented_better | both | n/a |
| right_slot | the replacement | n/a |
| filled_slot | the run | yes |
| faults | the replacement | |
| catalog_action | our catalog | 'table_tolerance' needs to be removed entirely. we should not ask the user if their table is okay with things like mill, land destruction, etc. 'power_level' is enough. |

## Item 39

Source: pr7-question-gate-run13, conversation "44. sideboard help only (probe)", turn 1. Row `colors`, slot `colors`, gap score 0.20.

**The catalog question:** Any color preference? A burn deck is strongest in red.

**What the model asked instead:** Do you have any color preference for your Modern Burn sideboard?

The agent sent this one in place of the catalog question.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | |
| invented_better | both | |
| right_slot | the replacement | |
| filled_slot | the run | yes |
| faults | the replacement | |
| catalog_action | our catalog | |

## Item 40

Source: pr7-question-gate-run10, conversation "23. land destruction", turn 4. Row `commander_pick`, slot `commander`, gap score 0.20.

**The catalog question:** Which commander do you want: The Earth King, Toph, Hardheaded Teacher, or Toph, the Blind Bandit? Say “none” and I’ll name three more.

**What the model offered instead:** Which one do you want: The Earth King, Toph, Hardheaded Teacher, or Toph, the Blind Bandit? Say "none" and I name three more.

The agent refused this one as a reword and sent the catalog question (D-88). Score it as if it had gone out. A `better` here says the refusal was wrong.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | |
| invented_better | both | |
| right_slot | the replacement | |
| filled_slot | the run | no |
| faults | the replacement | |
| catalog_action | our catalog | |

## Item 41

Source: pr7-question-gate-run13, conversation "4. fun and janky", turn 4. Row `commander_pick`, slot `commander`, gap score 0.20.

**The catalog question:** Which commander do you want: The Sackville-Bagginses, June, Bounty Hunter, or Hei Bai, Spirit of Balance? Choose “None—name three more” and I’ll name three more.

**What the model offered instead:** Which commander do you want: The Sackville-Bagginses, June, Bounty Hunter, or Hei Bai, Spirit of Balance? Say "none" and I’ll name three more.

The agent refused this one as a reword and sent the catalog question (D-88). Score it as if it had gone out. A `better` here says the refusal was wrong.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | |
| invented_better | both | |
| right_slot | the replacement | |
| filled_slot | the run | no |
| faults | the replacement | |
| catalog_action | our catalog | |

## Item 42

Source: pr7-question-gate-run12, conversation "45. copy a known deck (probe)", turn 1. Row `house_rules`, slot `house_rules`, gap score 0.05.

**The catalog question:** What does anything-goes mean at your table: any card with no ban list, or Vintage rules?

**What the model asked instead:** Are there any house rules you want applied to this Standard deck?

The agent sent this one in place of the catalog question.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | |
| invented_better | both | |
| right_slot | the replacement | |
| filled_slot | the run | no |
| faults | the replacement | |
| catalog_action | our catalog | |

## Item 43

Source: pr7-question-gate-run12, conversation "13. a card to keep that is not the commander", turn 2. Row `locked`, slot `locked`, gap score 0.05.

**The catalog question:** Must the deck keep Sanguine Bond, or may I cut cards that don’t fit the lifegain plan?

**What the model offered instead:** Do you want to keep Sanguine Bond in the deck, or may I cut a card that does not fit the plan?

The agent refused this one as a reword and sent the catalog question (D-88). Score it as if it had gone out. A `better` here says the refusal was wrong.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | |
| invented_better | both | |
| right_slot | the replacement | |
| filled_slot | the run | no |
| faults | the replacement | |
| catalog_action | our catalog | |

## Item 44

Source: pr7-question-gate-run11, conversation "20. dinosaur tribal for a child", turn 3. Row `pool`, slot `pool_rule`, gap score 0.20.

**The catalog question:** Build from your library first, only your library, or ignore it for a fully optimized deck?

**What the model asked instead:** Should I build the Gishath Dinosaur deck using your existing card library first, only cards from your library, or choose cards freely?

The agent sent this one in place of the catalog question.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | |
| invented_better | both | |
| right_slot | the replacement | |
| filled_slot | the run | yes |
| faults | the replacement | |
| catalog_action | our catalog | |

## Item 45

Source: pr7-question-gate-run13, conversation "43. a deck as a gift (probe)", turn 2. Row `power_commander`, slot `power`, gap score 0.20.

**The catalog question:** Which bracket does your table play? 2 is precon level, 3 is upgraded, 4 is high power.

**What the model asked instead:** What power level should the black Zombie deck target?

The agent sent this one in place of the catalog question.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | |
| invented_better | both | |
| right_slot | the replacement | |
| filled_slot | the run | yes |
| faults | the replacement | |
| catalog_action | our catalog | |

## Item 46

Source: pr7-question-gate-run11, conversation "52. the user answers a different question (probe)", turn 3. Row `table_tolerance`, slot `power`, gap score 0.05.

**The catalog question:** Does your table accept mill, land destruction, extra turns, or stax?

**What the model offered instead:** Does your table accept land destruction, extra turns, or stax?

The agent refused this one as a reword and sent the catalog question (D-88). Score it as if it had gone out. A `better` here says the refusal was wrong.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | n/a |
| invented_better | both | n/a |
| right_slot | the replacement | n/a |
| filled_slot | the run | no |
| faults | the replacement | |
| catalog_action | our catalog | 'table_tolerance' needs to be removed entirely. we should not ask the user if their table is okay with things like mill, land destruction, etc. 'power_level' is enough. |

## Item 47

Source: pr7-question-gate-run13, conversation "27. a card for the 99", turn 4. Row `commander_pick`, slot `commander`, gap score 0.05.

**The catalog question:** Which commander do you want: Grist, Voracious Larva // Grist, the Plague Swarm; Baron Sengir; or Galactus, Devourer of Worlds? Choose “None—name three more” and I’ll name three additional options.

**What the model offered instead:** Which commander do you want: Baron Sengir or Galactus, Devourer of Worlds? Say "none" and I’ll name three more.

The agent refused this one as a reword and sent the catalog question (D-88). Score it as if it had gone out. A `better` here says the refusal was wrong.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | |
| invented_better | both | |
| right_slot | the replacement | |
| filled_slot | the run | no |
| faults | the replacement | |
| catalog_action | our catalog | |

## Item 48

Source: pr7-question-gate-run13, conversation "5. 60-card anything goes", turn 1. Row `house_rules`, slot `house_rules`, gap score 0.20.

**The catalog question:** What does anything-goes mean at your table: any card with no ban list, or Vintage rules?

**What the model asked instead:** Which house rules should apply: any card with no ban list, or Vintage rules?

The agent sent this one in place of the catalog question.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | |
| invented_better | both | |
| right_slot | the replacement | |
| filled_slot | the run | no |
| faults | the replacement | |
| catalog_action | our catalog | |

## Item 49

Source: pr7-question-gate-run12, conversation "26. legacy for an event", turn 2. Row `locked`, slot `locked`, gap score 0.05.

**The catalog question:** Must the deck keep Delver of Secrets, or may I cut a card that does not fit the plan?

**What the model asked instead:** In this Legacy blue-red Delver of Secrets tempo deck, must Delver of Secrets stay?

The agent sent this one in place of the catalog question.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | |
| invented_better | both | |
| right_slot | the replacement | |
| filled_slot | the run | no |
| faults | the replacement | |
| catalog_action | our catalog | |

## Item 50

Source: pr7-question-gate-run11, conversation "37. a request we can not serve (probe)", turn 1. Row `format`, slot `format`, gap score 0.05.

**The catalog question:** Which format: Commander, Standard, Modern, or something else?

**What the model asked instead:** Which Yu-Gi-Oh format would you like?

The agent sent this one in place of the catalog question.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | |
| invented_better | both | |
| right_slot | the replacement | |
| filled_slot | the run | yes |
| faults | the replacement | |
| catalog_action | our catalog | |

## Item 51

Source: pr7-question-gate-run12, conversation "20. dinosaur tribal for a child", turn 3. Row `pool`, slot `pool_rule`, gap score 0.05.

**The catalog question:** Build from your library first, only your library, or ignore it for a fully optimized deck?

**What the model asked instead:** Should I build the Gishath Dinosaur deck from your library first, only from your library, or ignore your library?

The agent sent this one in place of the catalog question.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | |
| invented_better | both | |
| right_slot | the replacement | |
| filled_slot | the run | yes |
| faults | the replacement | |
| catalog_action | our catalog | |

## Item 52

Source: pr7-question-gate-run13, conversation "51. a returning player (probe)", turn 2. Row `power_commander`, slot `power`, gap score 0.05.

**The catalog question:** Which bracket does your table play? 2 is precon level, 3 is upgraded, 4 is high power.

**What the model asked instead:** What power level or bracket is your Commander table aiming for?

The agent sent this one in place of the catalog question.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | |
| invented_better | both | |
| right_slot | the replacement | |
| filled_slot | the run | yes |
| faults | the replacement | |
| catalog_action | our catalog | |

## Item 53

Source: pr7-question-gate-run12, conversation "52. the user answers a different question (probe)", turn 3. Row `table_tolerance`, slot `power`, gap score 0.20.

**The catalog question:** Does your Bracket 3 table accept mill, land destruction, extra turns, or stax?

**What the model offered instead:** What power level does your table accept for mill, land destruction, extra turns, or stax?

The agent refused this one as a reword and sent the catalog question (D-88). Score it as if it had gone out. A `better` here says the refusal was wrong.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | n/a |
| invented_better | both | n/a |
| right_slot | the replacement | n/a |
| filled_slot | the run | no |
| faults | the replacement | |
| catalog_action | our catalog | 'table_tolerance' needs to be removed entirely. we should not ask the user if their table is okay with things like mill, land destruction, etc. 'power_level' is enough. |

## Item 54

Source: pr7-question-gate-run13, conversation "38. shorthand and slang (probe)", turn 1. Row `house_rules`, slot `house_rules`, gap score 0.05.

**The catalog question:** What does anything-goes mean at your table: any card with no ban list, or Vintage rules?

**What the model asked instead:** Do you have any house rules for your Gruul dinosaur stompy deck?

The agent sent this one in place of the catalog question.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | |
| invented_better | both | |
| right_slot | the replacement | |
| filled_slot | the run | no |
| faults | the replacement | |
| catalog_action | our catalog | |

## Item 55

Source: pr7-question-gate-run12, conversation "27. a card for the 99", turn 3. Row `locked`, slot `locked`, gap score 0.20.

**The catalog question:** Must the deck keep Grist, the Hunger Tide and Grist, or may I cut a card that does not fit the sacrifice plan?

**What the model offered instead:** Should the deck keep Grist, the Hunger Tide in the 99, or may I cut a card that does not fit the plan?

The agent refused this one as a reword and sent the catalog question (D-88). Score it as if it had gone out. A `better` here says the refusal was wrong.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | |
| invented_better | both | |
| right_slot | the replacement | |
| filled_slot | the run | no |
| faults | the replacement | |
| catalog_action | our catalog | |

## Item 56

Source: pr7-question-gate-run13, conversation "17. precon upgrade at bracket 2", turn 1. Row `pool`, slot `pool_rule`, gap score 0.20.

**The catalog question:** Build from your library first, only your library, or ignore it for a fully optimized deck?

**What the model asked instead:** Do you want to start with your current Atraxa precon, use only cards already in it, or ignore it for a fully optimized deck?

The agent sent this one in place of the catalog question.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | |
| invented_better | both | |
| right_slot | the replacement | |
| filled_slot | the run | yes |
| faults | the replacement | |
| catalog_action | our catalog | |

## Item 57

Source: pr7-question-gate-run13, conversation "52. the user answers a different question (probe)", turn 3. Row `table_tolerance`, slot `power`, gap score 0.20.

**The catalog question:** Does your table accept mill, land destruction, extra turns, or stax?

**What the model offered instead:** What power level does your table accept for mill, land destruction, extra turns, or stax?

The agent refused this one as a reword and sent the catalog question (D-88). Score it as if it had gone out. A `better` here says the refusal was wrong.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | n/a |
| invented_better | both | n/a |
| right_slot | the replacement | n/a |
| filled_slot | the run | no |
| faults | the replacement | |
| catalog_action | our catalog | 'table_tolerance' needs to be removed entirely. we should not ask the user if their table is okay with things like mill, land destruction, etc. 'power_level' is enough. |

## Item 58

Source: pr7-question-gate-run13, conversation "13. a card to keep that is not the commander", turn 2. Row `locked`, slot `locked`, gap score 0.20.

**The catalog question:** Should the deck keep Sanguine Bond, or may I cut it if it does not fit the lifegain plan?

**What the model offered instead:** Must the deck keep Sanguine Bond, or may I cut a card that does not fit the plan?

The agent refused this one as a reword and sent the catalog question (D-88). Score it as if it had gone out. A `better` here says the refusal was wrong.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | |
| invented_better | both | |
| right_slot | the replacement | |
| filled_slot | the run | no |
| faults | the replacement | |
| catalog_action | our catalog | |

## Item 59

Source: pr7-question-gate-run13, conversation "20. dinosaur tribal for a child", turn 3. Row `pool`, slot `pool_rule`, gap score 0.05.

**The catalog question:** Should I build from your library first, use only cards from your library, or ignore your library for a fully optimized deck?

**What the model offered instead:** Build from your library first, or ignore it for a fully optimized deck?

The agent refused this one as a reword and sent the catalog question (D-88). Score it as if it had gone out. A `better` here says the refusal was wrong.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | |
| invented_better | both | |
| right_slot | the replacement | |
| filled_slot | the run | yes |
| faults | the replacement | |
| catalog_action | our catalog | |

## Item 60

Source: pr7-question-gate-run11, conversation "21. vintage with proxies", turn 1. Row `house_rules`, slot `house_rules`, gap score 0.20.

**The catalog question:** When you say you proxy everything, do you mean any card with no ban list, or using Vintage rules?

**What the model offered instead:** What does your table's anything-goes rule mean: any card with no ban list, or Vintage rules?

The agent refused this one as a reword and sent the catalog question (D-88). Score it as if it had gone out. A `better` here says the refusal was wrong.

| Field | Reads | Value |
|---|---|---|
| catalog_enough | the catalog question | |
| invented_better | both | |
| right_slot | the replacement | |
| filled_slot | the run | yes |
| faults | the replacement | |
| catalog_action | our catalog | |

