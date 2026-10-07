# Questions only the owner can answer

This file is the decision queue. `docs/open-questions.md` is the roadmap log, and it holds the questions that wait for a stage. This file holds the questions that wait for a person.

The tuning loop reads this file (D-133). A fixer agent must not decide any question listed here. When a fix needs one of these answers, the agent skips that fix and says so.

Every row came out of the sessions of 2026-08-25 and 2026-08-26. Their sources are the owner's scoring of items 1 to 32, the correction session, and the batch sweep of all 66 conversations.

## Blocks the automation

No owner question blocks it. The owner answered OQ-24 to OQ-27 on 2026-08-26 (D-135 to D-138), and `AUTOTUNE_FIXER_CMD` names the fixer (D-159). The loop started seven times and kept nothing, and D-171, D-177, and D-178 record what it got wrong. D-230 and D-258 set the noise margin.

## Waits on a decision

No question waits on a decision now. D-1192 answered OQ-95.

## The two numbers M-5 exists to set

| # | Question | Why only you | What it blocks |
|---|---|---|---|

## Product questions the data raised

| # | Question | Why only you | What it blocks |
|---|---|---|---|
| OQ-97 | A card newer than the top lists reads as "unseen", and the grade falls. Should the grade read such a card as unknown? The options: (a) leave it out of the unseen share and the card rate, which gives a fair grade to a new set, but a deck of new cards then reads on fewer cards; (b) give it the mean rate of its rarity, which keeps each card in the count, but the mean is a guess; (c) keep the grade and add a note that names the new cards, which needs no new fit, but the grade stays low. | Deck `dDD9Iav9aGhgGF9dDLNQ` of Reality Fracture, four days after the release, read "many of the cards appear in no top list, and that lowers the grade". Each option changes the features of the fitted model, or the text of the grade, and the owner chose that model. | A fair grade for a deck of a new set. F-226 and D-1190 do not touch the grade. |
| OQ-98 | The question phase of chat `dDD9Iav9aGhgGF9dDLNQ` fails on its own answers. Of 38 replays of its four turns, 21 built no deck. Of the 17 decks, 10 lost the answer "1 Exhibition" and built at bracket 3. On turn 4, "Skip that question.. Jace, Multiverse Architect" often read Jace as a named card and asked its role, or said "I did not read your last answer". The live user got a bracket 1 deck. The options: (a) a new item fixes the classifier for a bracket answer and for a commander pick, which helps each user but needs paid question gates; (b) the replay tool answers each extra question with "Skip that question", which is cheap but hides the fault; (c) both. | The live chat built its deck, so the fault reaches the replays more than the user. The choice sets the next item and its cost. | Each replay of a chat with these answers. #294 judged its bar on pinned replays (D-1201). |

## How to answer

Write the answer into `docs/decisions.md` as a new row, with the date. Then delete the row here. A row that stays here is a question the loop still refuses to decide.
