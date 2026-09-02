# Questions only the owner can answer

This file is the decision queue. `docs/open-questions.md` is the roadmap log, and it holds the questions that wait for a stage. This file holds the questions that wait for a person.

The tuning loop reads this file (D-133). A fixer agent must not decide any question listed here. When a fix needs one of these answers, the agent skips that fix and says so.

Every row came out of the sessions of 2026-08-25 and 2026-08-26. Their sources are the owner's scoring of items 1 to 32, the correction session, and the batch sweep of all 66 conversations.

## Blocks the automation

No owner question blocks it. The owner answered OQ-24 to OQ-27 on 2026-08-26 (D-135 to D-138), and `AUTOTUNE_FIXER_CMD` names the fixer (D-159). The loop started seven times and kept nothing, and D-171, D-177, and D-178 record what it got wrong. D-230 and D-258 set the noise margin.

## Waits on a decision

| # | Question | Why only you | What it blocks |
|---|---|---|---|
| OQ-52 | Do D-461 to D-464 stand as written? They are the session calls of PR-14A: a profile finding warns and buys a repair, the tags cut the shortlist, the bands start as estimates, and the gate has four bars. | The bands and the thresholds set what a bracket means in this app, and that is the promise to your users (D-452). | Nothing blocks on it. The bracket gate runs on the values as they are, and a change after the run needs a new decision id. |

## The two numbers M-5 exists to set

| # | Question | Why only you | What it blocks |
|---|---|---|---|

## Product questions the data raised

| # | Question | Why only you | What it blocks |
|---|---|---|---|

## How to answer

Write the answer into `docs/decisions.md` as a new row, with the date. Then delete the row here. A row that stays here is a question the loop still refuses to decide.
