# Questions only the owner can answer

This file is the decision queue. `docs/open-questions.md` is the roadmap log, and it holds the questions that wait for a stage. This file holds the questions that wait for a person.

The tuning loop reads this file (D-133). A fixer agent must not decide any question listed here. When a fix needs one of these answers, the agent skips that fix and says so.

Every row came out of the sessions of 2026-08-25 and 2026-08-26. Their sources are the owner's scoring of items 1 to 32, the correction session, and the batch sweep of all 66 conversations.

## Blocks the automation

No owner question blocks it. The owner answered OQ-24 to OQ-27 on 2026-08-26 (D-135 to D-138), and `AUTOTUNE_FIXER_CMD` names the fixer (D-159). The loop started seven times and kept nothing, and D-171, D-177, and D-178 record what it got wrong. D-230 and D-258 set the noise margin.

## Waits on a decision

| # | Question | Why only you | What it blocks |
|---|---|---|---|
| OQ-77 | An account off the invite list: is the form check enough, or do you want a blocking function of Identity Platform that refuses the account server-side? **The cost objection is gone**: Firebase reads "Authentication with Identity Platform" as no-cost up to 50,000 monthly active users (firebase.google.com/pricing, read 2026-09-07). The app has three. What is left is the work and one more moving part: a `beforeCreate` function to deploy, and a seven-second answer limit that fails the sign-up when it passes. | It changes the sign-in path of every reader. | The last hole of the invite gate. A caller who drives the Firebase API directly still makes an account today, and that account reads nothing: the API refuses every call and the Firestore rules deny all. |
| OQ-67 | Stage B channels: web push through Cloud Messaging, an email digest, or both? And which events: a legality change on a deck, new cards for a deck, a finished build? | Each channel asks a user for a permission in your name. | The Stage B PR after PR-23. |

## The two numbers M-5 exists to set

| # | Question | Why only you | What it blocks |
|---|---|---|---|

## Product questions the data raised

| # | Question | Why only you | What it blocks |
|---|---|---|---|
| OQ-80 | **Does a card the reader marks as a proxy count as a card they own?** A Moxfield export carries a `Proxy` column, and no other format this app reads carries one. A proxy sits in the binder and plays at a table that allows one, and the reader never bought the card. So owned-only either builds with it or refuses it, and owned-first either leaves it off the buy list or prices it. The same card is owned for play and unowned for money, and the app holds one meaning of owned for both. Three answers: a proxy counts as owned, a proxy never counts, or the reader decides on the build. | It sets what "only cards I own" promises, the way OQ-79 does, and it changes what a buy list costs. | Nothing today. The parser drops the column, and every row of the one export on record reads False. A later answer needs the flag on the entry, and no reader has uploaded one. |
| OQ-82 | **Must a bracket request build a deck with the power of that bracket?** Deck 3 asked for a high-power Urza, Lord High Artificer deck at bracket 4, and the owner grades the build baseline (F-110). The owner's review says the fix belongs in the builder, if a bracket 4 request must yield a bracket 4 deck. A weaker deck inside the bracket is the other answer, and the build then says so. Session `OFMnk7Tv2zkK8xfAwXxB` asked for bracket 5 and warned that its interaction of 7 and its turn-four mana of 4.58 sit under the bracket 5 bands. | It sets what a bracket promises a reader. | A builder item for the power of a bracket. Nothing today. |

## How to answer

Write the answer into `docs/decisions.md` as a new row, with the date. Then delete the row here. A row that stays here is a question the loop still refuses to decide.
