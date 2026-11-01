# Questions only the owner can answer

This file is the decision queue. `docs/open-questions.md` is the roadmap log, and it holds the questions that wait for a stage. This file holds the questions that wait for a person.

The tuning loop reads this file (D-133). A fixer agent must not decide any question listed here. When a fix needs one of these answers, the agent skips that fix and says so.

Every row came out of the sessions of 2026-08-25 and 2026-08-26. Their sources are the owner's scoring of items 1 to 32, the correction session, and the batch sweep of all 66 conversations.

## Blocks the automation

No owner question blocks it. The owner answered OQ-24 to OQ-27 on 2026-08-26 (D-135 to D-138), and `AUTOTUNE_FIXER_CMD` names the fixer (D-159). The loop started seven times and kept nothing, and D-171, D-177, and D-178 record what it got wrong. D-230 and D-258 set the noise margin.

## Waits on a decision

OQ-120 waits on a decision. D-1192 answered OQ-95.

| # | Question | Why only you | What it blocks |
|---|---|---|---|
| OQ-120 | GO-2026-6809 names `crypto/x509` of Go 1.27.2, and Go 1.27.3 fixes it. It fails `verify:go` and `verify:vuln` on `main`. `pnpm audit --audit-level high` also reads 7 high advisories in the lockfile of `main`, and `verify:web` fails. Where do the two fixes go? The options: (a) one pull request of their own, which keeps each change apart, but each open pull request waits for its merge; (b) the pull request that the failure stops, as D-1172 did, which ends the wait, but a live eval can not change CI or the Dockerfiles (D-1157, D-1158); (c) the Dependabot pull requests #307 and #259, which carry the bumps, but #307 fails the same checks. | The Go bump changes CI and both Dockerfiles (D-1216), and D-1218 made it an owner decision. Each earlier audit fix was an owner decision (D-1006, D-1022, D-1172). | The Codex record of #308 reads `Blocked` on the red checks alone, with no finding. #308 and each other pull request wait for green checks. |

## The two numbers M-5 exists to set

| # | Question | Why only you | What it blocks |
|---|---|---|---|

## Product questions the data raised

| # | Question | Why only you | What it blocks |
|---|---|---|---|
| OQ-97 | A card newer than the top lists reads as "unseen", and the grade falls. Should the grade read such a card as unknown? The options: (a) leave it out of the unseen share and the card rate, which gives a fair grade to a new set, but a deck of new cards then reads on fewer cards; (b) give it the mean rate of its rarity, which keeps each card in the count, but the mean is a guess; (c) keep the grade and add a note that names the new cards, which needs no new fit, but the grade stays low. | Deck `dDD9Iav9aGhgGF9dDLNQ` of Reality Fracture, four days after the release, read "many of the cards appear in no top list, and that lowers the grade". Each option changes the features of the fitted model, or the text of the grade, and the owner chose that model. | A fair grade for a deck of a new set. F-226 and D-1190 do not touch the grade. |
| OQ-112 | D-1226 plays an owned twin in place of a card to buy under each pool rule with a collection, any-card included. Keep that, or run it under owned-first alone? The options: (a) keep it, which gives an any-card deck the same play and a shorter buy list, but the deck list then follows the collection; (b) owned-first alone, which keeps any-card as D-37 wrote it, "information only", but the replays of F-232 then buy Llanowar Elves beside an owned Fyndhorn Elves. | Each replay of the eval read the gap answer as any-card. A twin has the same cost, type line, and rules text, so no option changes how the deck plays. | The scope of D-1226. |
| OQ-113 | The mana row of D-382 counts the ramp cards and the nonbasic lands of a set family against the ramp target plus the land target. Basic lands fill the land target, and the count leaves them out. Since D-1227 reads the colors of the commander, a mono-red Hobbit commander holds 22 such cards against 46, and the row asks. Keep that, or count the ramp cards against the ramp target alone? The options: (a) keep it, which asks one question for most set-limited decks of one or two colors; (b) ramp against the ramp target, beside the fixing floor of D-1227, which asks only when the sets can not make the mana, and a deck then gets fewer nonbasic lands of other sets. | The rule of D-382 is an owner decision, and D-1227 changes how often it fires. | No code. D-1227 keeps the rule of D-382 as it stands. |

## How to answer

Write the answer into `docs/decisions.md` as a new row, with the date. Then delete the row here. A row that stays here is a question the loop still refuses to decide.
