# Questions only the owner can answer

This file is the decision queue. `docs/open-questions.md` is the roadmap log, and it holds the questions that wait for a stage. This file holds the questions that wait for a person.

The tuning loop reads this file (D-133). A fixer agent must not decide any question listed here. When a fix needs one of these answers, the agent skips that fix and says so.

Every row came out of the sessions of 2026-08-25 and 2026-08-26. Their sources are the owner's scoring of items 1 to 32, the correction session, and the batch sweep of all 66 conversations.

## Blocks the automation

No owner question blocks it. The owner answered OQ-24 to OQ-27 on 2026-08-26 (D-135 to D-138), and `AUTOTUNE_FIXER_CMD` names the fixer (D-159). The loop started seven times and kept nothing, and D-171, D-177, and D-178 record what it got wrong. D-230 and D-258 set the noise margin.

## Waits on a decision

| # | Question | Why only you | What it blocks |
|---|---|---|---|

## The two numbers M-5 exists to set

| # | Question | Why only you | What it blocks |
|---|---|---|---|

## Product questions the data raised

| # | Question | Why only you | What it blocks |
|---|---|---|---|
| OQ-57 | Does a paid sweep run in CI at all? The roadmap names "Tier 1 nightly" and a label-gated sweep on PRs. The sweep of 2026-09-03 cost $4.60 and 84 minutes of provider time. So a nightly run is about $140 a month and 42 hours of runner time, and the keys move to repo secrets. My recommendation: the sweep runs on your machine, on your word, with a cost cap, and CI runs Tier 0 alone. | Money, and the Actions quota you hit in August. | Slice 6 of PR-15 and the CI job of slice 3. |
| OQ-58 | Are these the four fields of the plan judge? `plan_coherent` (the cards serve the plan the summary states), `theme_fit`, `useful_as_built` (a player can pilot it as is), and `summary_honest` (the summary claims nothing the deck lacks). Each one reads no, partly, or yes. | D-66 set the M-5 rubric, and a rubric change invalidates a fit. | Slice 5 of PR-15 stays information only until you confirm. |
| OQ-59 | Do you score the last 30 items of the M-5 sheet before PR-15 sets the fit threshold? 30 of 60 are scored on 2026-09-04, and D-66 asks for 50. | Your hours. | The number of D-423. The harness ships without it. |
| OQ-60 | May the repo hold a trimmed card snapshot for Tier 0? Its size is unmeasured, a few megabytes at most. It gives the free dry runs of the deck gate in CI, and such a dry run found the empty exclusion of 2026-09-04. | Repo size is yours. | The dry-run lane of slice 3. The compare lane needs no snapshot. |
| OQ-56 | Does the whole-precon check of D-408 ignore basic lands? The export of 2026-08-24 holds 84 of the 90 printings of Avengers Assemble, and the six absent cards are basic lands. A ManaBox deck binder can omit them. The exclusion never removes a basic land (D-37), so the check loses nothing without them. | D-408 is your rule, and it decides who owns a precon. | Nothing in code. A reader whose binder omits the basics hears "holds no whole precon" today (F-35). |

## How to answer

Write the answer into `docs/decisions.md` as a new row, with the date. Then delete the row here. A row that stays here is a question the loop still refuses to decide.
