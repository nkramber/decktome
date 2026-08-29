# Questions only the owner can answer

This file is the decision queue. `docs/open-questions.md` is the roadmap log, and it holds the questions that wait for a stage. This file holds the questions that wait for a person.

The tuning loop reads this file (D-133). A fixer agent must not decide any question listed here. When a fix needs one of these answers, the agent skips that fix and says so.

Every row came out of the sessions of 2026-08-25 and 2026-08-26. Their sources are the owner's scoring of items 1 to 32, the correction session, and the batch sweep of all 66 conversations.

## Blocks the automation

No owner question blocks it. The owner answered OQ-24 to OQ-27 on 2026-08-26 (D-135 to D-138), and `AUTOTUNE_FIXER_CMD` names the fixer (D-159). The loop started seven times and kept nothing, and D-171, D-177, and D-178 record what it got wrong. D-230 and D-258 set the noise margin.

## Waits on a decision

| # | Question | Why only you | What it blocks |
|---|---|---|---|
| OQ-39 | How far can the cost-tier eval fall behind a stronger judge? `make eval-calibrate` reports the agreement, and no number sets the floor. D-149 answers the card-fact half of this for nothing, and it caught one false claim on its first run. The dual-judge proposal stays open for the rest. | It is a tolerance, and tolerances are yours. A cheap judge that refuses half as many questions still reports a real floor, and it hides the other half. | How much weight the ratio of the loop carries. |

## The two numbers M-5 exists to set

| # | Question | Why only you | What it blocks |
|---|---|---|---|
| OQ-28 | What is the fit threshold, D-27? The code still holds the provisional 0.35. | D-27 reserves the number for you. The version-1 sheet did not set it: no candidate met the 80 percent floor. | The gap score, and how often the model can invent a question. |
| OQ-29 | Is the reword overlap bar of 0.6 right, D-88? | The version-1 report said the guard is too tight, on evidence that did not tell an exact copy from a new question. The version-2 sheet can tell them apart. | The reword guard. |
| OQ-30 | Do you score items 33 to 60 of `pr7-m5-scoring.md`, or stop at 32? | The sheet measures the old engine, and its threshold no longer applies (D-96). Your 32 scores are already spent. | Nothing. It is sunk work, and the answer saves you two hours. |
| OQ-31 | Items 3 and 30 are the same item and you scored `invented_better` as `same` and `worse`. Which stands? | Only you know which reading you meant. D-106 settled the other pair. | The self-consistency measure of the version-1 sheet. |

## Product questions the data raised

| # | Question | Why only you | What it blocks |
|---|---|---|---|
| OQ-23 | What does the agent do when the user asks a question back, such as "What is a bracket?" | Probe 36 asks it. The catalog holds questions and no answers, so an answer path is a product decision. | Probe 36, and the reply to a user question in the chat (a live gap since PR-12). |
| OQ-37 | Do the 14 terse conversations join the catalog-only bar, or stay probes? | D-105 froze the first 30 so the bar stays comparable. At some point the bar must measure the harder set. | The meaning of the gate verdict. |

## How to answer

Write the answer into `docs/decisions.md` as a new row, with the date. Then delete the row here. A row that stays here is a question the loop still refuses to decide.
