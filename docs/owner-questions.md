# Questions only the owner can answer

This file is the decision queue. `docs/open-questions.md` is the roadmap log, and it holds the questions that wait for a stage. This file holds the questions that wait for a person.

The tuning loop reads this file (D-133). A fixer agent must not decide any question listed here. When a fix needs one of these answers, the agent skips that fix and says so.

Every row came out of the sessions of 2026-08-25 and 2026-08-26. Their sources are the owner's scoring of items 1 to 32, the correction session, and the batch sweep of all 66 conversations.

## Blocks the automation

No owner question blocks it. The owner answered OQ-24 to OQ-27 on 2026-08-26 (D-135 to D-138), and `AUTOTUNE_FIXER_CMD` names the fixer (D-159). The loop started seven times and kept nothing, and D-171, D-177, and D-178 record what it got wrong. D-230 and D-258 set the noise margin.

## Waits on a decision

| # | Question | Why only you | What it blocks |
|---|---|---|---|
| OQ-78 | A slot that stays asked after the reader answered it stops the chat for good (F-70, D-351). Do you want the agent to ask again, or to take the answer as a decline and go on? | It changes what the agent does with an answer it can not read, and both readings lose something. | The dead conversation class. Two sessions have stopped this way. |
| OQ-77 | An account off the invite list: is the form check enough, or do you want a blocking function of Identity Platform that refuses the account server-side? | It upgrades Firebase Authentication to Identity Platform, and that has a price and a quota of its own. | The last hole of the invite gate. A caller who drives the Firebase API directly still makes an account today. |
| OQ-67 | Stage B channels: web push through Cloud Messaging, an email digest, or both? And which events: a legality change on a deck, new cards for a deck, a finished build? | Each channel asks a user for a permission in your name. | The Stage B PR after PR-23. |

## The two numbers M-5 exists to set

| # | Question | Why only you | What it blocks |
|---|---|---|---|

## Product questions the data raised

| # | Question | Why only you | What it blocks |
|---|---|---|---|

## How to answer

Write the answer into `docs/decisions.md` as a new row, with the date. Then delete the row here. A row that stays here is a question the loop still refuses to decide.
