# Session hand-off

`CLAUDE.md` is the entry point. It sends you here. Then read `docs/decisions.md` from D-297, `docs/owner-questions.md`, and `docs/open-questions.md`.

This file holds the ten most recent sessions. Every older record sits in `docs/reference/session-handoff-archive.md`, word for word. The first narratives sit in `docs/reference/session-log-2026-08-23-to-26.md` and `docs/reference/session-log-2026-08-27-to-28.md`. The audit of 2026-08-29 is `docs/audit-2026-08-29.md`.

CAUTION: the web tests need Node 22.23.2 (`.nvmrc`). Under Node 20 every test file fails at start with `ERR_REQUIRE_ESM` from jsdom 30. Put `~/.nvm/versions/node/v22.23.2/bin` on the PATH before `make verify`. On this machine `nvm use` reports the change and does not make it, so prepend the path yourself.

## RESUME HERE (2026-09-10)

**The checkout.** `main` is `62b9aaa`, which is pull request #135, or a later merge. No branch holds open work. Dependabot's #131 and #132 merged the same day. #133, the web bumps, took a fix on its own branch. It holds the regenerated TypeScript of protoc-gen-es 2.14.1, and a type augmentation that puts the axe matcher on vitest 5's assertion. Run `make where` before you touch anything. Never commit on `main` (D-583). Answer the review of `gitar-bot` before you ask for a merge (D-637).

**The app is live on `decktome.com`.** A merge to `main` deploys itself on Cloud Build (D-584, D-586).

**The commander offer ignored the bracket of its own turn, and #134 fixed it** (F-104, D-662). The turn built the hints from the restored slots, and the picked power option landed after that. The pick row fires on the turn the power answer arrives (D-631), so the bracket 5 signal ordered almost no deployed offer. Session `vDzEKDPnRZntlyhz8Lqg` read Lotho, Ghalta, and Peregrin Took. The hints take the bracket after the answers now, and a test holds it.

**PR-39 is merged as #135** (D-659 to D-661, F-102, F-103). The source count reads a reliable tap ability alone, and the balance phase trades basics toward the color that falls short. The measurement of D-661 is done. Over run 18 the pass moves four decks. The worst color rises on all four and falls on none, and deck 22 gains two Forests. `docs/reference/pr39-manapass-run18.md` holds the lane.

CAUTION: the `decktome` gcloud configuration named the Wallabee account and project on 2026-09-10, so `use_decktome` put the shell on the wrong project. `scripts/read-session.sh` reads `SESSION_PROJECT` and the account of `CLOUDSDK_CORE_ACCOUNT`, so an environment override reads `decktome-prod` with no change to the configuration. The owner has the commands to repair the configuration.

**PR-37 is merged as #127, and gate run 18 reads PASS** (D-653). The synergy check drops a Commander copy whose break lowered the synergy feature by less than 0.10 standard deviations. The Commander synergy axis reads 0.84 of 232, and the precon bar 0.95 of 788. No earlier run of the quality gate in the repository reads PASS.

**CAUTION: the PASS comes from the population of the bar, and not from a better model.** The built decks graded bad stay at 15 of 25, and the judge agreement stays at 8 of 25. 37 of the 232 synergy pairs still lose, and those are the true gap of the axis.

**CAUTION: the check reads Commander alone** (D-653, F-96). In Modern the synergy break does not lower the synergy feature in 825 of 845 pairs. A check there dropped 78 percent of the Modern synergy copies, and three decks the judge reads as bad graded higher.

**CAUTION: the merge changed the production model.** The deploy of `302858a` succeeded on 2026-09-10 and pointed the meta job at the new image. The meta job refits the model on every run (`go/cmd/worker/meta.go`), so its next run stores a model with the check.

**The premise of D-652 was half right.** The check drops 157 of 389 Commander pairs from every band of precon strength, and not only the 54 weak pairs M-8 named. `docs/reference/pr37-synergy-materiality-2026-09-10.md` holds all six floors.

**The gate reports the three numbers of D-648 for free.** `make quality-gate` grades the decks of deck gate run 16 with the model it fits, and it reads the judge's tiers from judge lane run 5. It stores no model.

**The owner grades ten disputed decks by hand next** (D-657, F-100). M-10's softmax scorer grades real lists far better, and the judge agreement moved from 8 to 9 only. The defect detector holds eight of those ten decks at bad. The review sheet is `.local/review/disputed-decks-2026-09-10.md`, beside a CSV, and it stays out of git. `docs/reference/m10-softmax-scorer-2026-09-10.md` holds the fits.

**Three items aimed at the reader gap failed first**: PR-29, M-9, and PR-30 (F-94, F-98, F-99). The owner parked PR-35 and dropped PR-30.

**CAUTION: two features never get built again** (F-94). `casual_unseen_share` moved the target axis by nothing and cost three other bars. `casual_pair_share` won one pair of 389, because the fit split it against `casual_synergy` with opposite signs. `docs/reference/pr29-casual-corpus-2026-09-10.md` says why for each, and PR-35 builds neither one.

**Every quality item reports three numbers** (D-648): the precon bar, the count of built decks graded bad, and the tier judge agreement. **The owner drops an item that moves none of the three.** PR-29 moved the first up and the other two down, and that is why it closed.

**The whole feedback loop stands**, as #117, #121, and #122. **No live cycle ran yet.** It needs the owner's word, `AUTOTUNE_FIXER_CMD`, and a harvest whose verdicts carry a snapshot.

**PR-34 is merged as #123.** The app reads the format of an upload out of the file, and Moxfield reads. OQ-80 asks whether a proxy counts as a card the reader owns.

**The whole PR-22 gate holds** (D-633). The measured cost at idle reads $9.54 a month gross, and the hourly snapshot tick of D-634 leaves $3.81.

**What waits on the owner.** OQ-67, the Stage B channels. OQ-77, the blocking function of Identity Platform. OQ-79, the commander of an owned-only pool. OQ-80, a proxy and the pool.

**CAUTION: the evidence OQ-79 waits for is not on the deployed project.** A fresh owned-only build on the deployed app is the way to it.

**CAUTION: a local `make verify` is not the whole story.** It read green for weeks while shellcheck failed (F-89, D-641). Every pull request runs the workflow now.

## How to resume

1. Run `make where`. It prints the branch, the tree, and the state of the branch's pull request.
2. Run `ps aux | grep autotune` before any write. The loop resets the tree when it rejects an iteration.
3. Make a branch from `main`. Never commit on `main`, and never push to it (D-583). The owner merges (D-585). Run `make hooks` one time in a fresh checkout.
4. Load the skills. Load `ste-writing` before you write any `.md`. Load `design-doc-style` before you edit the roadmap. Load `mtg-corpus` before you reason about a format, a legality, or a card term.
5. Run `make verify`. It runs every check the verify workflow runs, on this machine, for nothing. Put Node 22.23.2 on the PATH first. The pull request runs the same jobs on Actions (D-639).
6. Do "Next steps, in order" below. Ask questions as they come up. Record each owner answer in `docs/decisions.md`, and delete the row from `docs/owner-questions.md`.
7. Open the pull request, then wait for the review of `gitar-bot` and answer it (D-637). A finding with merit takes a change, a commit, a push, and a reply. A finding with no merit takes a reply that says why, and you resolve it. Tell the owner when the pull request is ready to merge. Gitar is the only review this repo asks for.
8. Before you end, update this file. Move the oldest session to the archive when the count passes ten.

A second Mac: `docs/setup-second-mac.md` holds what to carry, what to install, and how to prove the machine.

Seven things a fresh session gets wrong without this file.

- Twelve targets and two loop scripts spend money: `make questions-gate`, `make questions-eval`, `make eval-calibrate`, `make deck-gate`, `make bracket-gate`, `make revise-gate`, `make chat-probe`, `make generate-probe`, `make summary-judge`, `make quality-judge`, `make test-smoke`, `make feedback-triage`, `scripts/autotune.sh`, and `scripts/feedback-loop.sh`. Ask the owner before each run. `make autotune`, `make feedback-loop`, `make feedback-loop-dry`, `make feedback-triage-dry`, and `eval sweep -dry` are free.
- A rerun writes to a new file. Every `*_OUT` variable refuses a document that holds a result (D-65).
- A gate run takes about 20 minutes and an eval about 13. A foreground command stops at 10 minutes, so run both in the background.
- `docker compose up` needs provider keys in `.env` now, and fails fast without them (D-267). `make dev` still starts with no keys, and the fake builds the fixture deck of PR-23 for any first message (D-552).
- The eval and the agent share a model. Every ratio it reports is a floor, not a measurement (D-136).
- The judge is noisy: two runs of identical code move up to nine bad questions (D-230). One run proves nothing on its own.
- The sweep's estimate of a step is the cost of its last run file. The deck gate estimate read $2.54 from run 14, and run 16 cost $3.79 with the plan judge and 12 repair turns. A cap set from the estimate stops the sweep before its last step.

## Facts that expire

- Comprehensive Rules: the current file is 2026-08-19 (D-272). The rule citations in the corpus match the 2026-08-07 text.
- Ban-list snapshot: 2026-08-24. Next announcement 2026-10-12, in `internal/cards/announcement_dates.json`. A test fails when that file holds no future date. This is by design.
- Commander brackets: the 2025-10-21 revision. Game Changers: 53 cards, list of 2026-02-09. Lutri is banned as a companion only, per the 2026-02-09 announcement (`companion_bans.json` holds the link). The content rules per bracket in `brackets.json` and the Spellbook thresholds were read 2026-09-02, and the Karsten tables are the 2022 articles, read 2026-09-02 (`docs/reference/bracket-profile-2026-09-02.md`).
- Standard: 18 sets, Wilds of Eldraine to The Hobbit. Six sets leave at the first 2027 set. Verified 2026-08-24.
- Card snapshot on disk: `.local/gcs/mtg-local-cards/scryfall/20260904T210157`, the newest of 13. Read 2026-09-09. Check every card fact against it.
- LLM model ids and prices: `roles.json` and `prices.json`, verified 2026-08-24. The max output per provider in `llm/client.go`, verified 2026-08-29. The OpenAI rows are unverified by anyone but the owner. The judge role runs on Opus 5 (D-430).
- The deck gate upgrade probe cost $0.32 for 2 prompts, 7 calls, and 266 seconds. A partial run reads its own item bars and never stands as the gate (D-526).
- Run cost, read from the run files on 2026-09-09. The question gate cost $0.190 on run 42, $0.193 on run 43, and $0.194 on run 44, over 18 to 21 minutes. The deck gate upgrade probe cost $0.32 for 2 prompts. The eval cost $0.092 on run 34 and $0.096 on run 35, over 12 to 14 minutes. The deck gate cost $2.71 on run 18, over 37 minutes. The revise gate cost $1.07 on run 9. The bracket gate judge lane cost $0.28.
- The backfill of 2026-09-09 read one user with a record to seed: 1 collection, 1 thumbs up, and 2 thumbs down. It counted no deck and no chat, because the reader deleted both (D-635).
- The feedback store holds 3 verdicts on 2026-09-09, and every one predates the snapshot of D-635. `make feedback-list VERDICT=` reads both verdicts now (F-87). A collection group query over it needs an index for its shape, and the store holds "verdict ascending, created_at descending" alone.
- The deployed schedules, read 2026-09-09: `mtg-snapshot-schedule` at `0 * * * *` (D-634) and `mtg-meta-schedule` at `0 6 * * *`. Both read ENABLED. The API service holds minScale 0, so it scales to zero. No billing export exists, so no command reads the billed spend.
- Baselines, in `docs/reference/eval/baselines.json`: questions is run 42, decks is run 18, revise is run 9. Run 44 is the newest whole questions run, and run 43 records the regression of F-84. `make eval-check` compares the newest whole run of a suite against its baseline. The quality gate has no baseline row, and run 18 is its newest run.
- Toolchain: Go 1.27.0, Node 22.23.2, pnpm 9.2.0, firebase-tools 14.14.0, Java 17, Playwright 1.63.0 with its Chromium headless shell (`playwright install chromium`). vitest is 5.0.0 since #133. The first three were verified 2026-09-09.

## Next steps, in order

1. **Read a bracket 4 or 5 offer on the deployed app** after the deploy of #134. It shows the fix of F-104 on a real turn.
2. **The detector reads built decks wrong** (D-659). Seven of its eight flags on the disputed decks are wrong, and the one real defect moves no weight. It comes after the basic lands (D-660).
3. **Read the log of the next meta job.** Its quality fit line for Commander names `immaterial`. Check that the count reads above zero.
4. **The next collection platform, when the owner names one** (F-91). Five are left: Archidekt, Deckbox, Delver Lens, TCGplayer, and Helvault. Each one takes a real export and never a column list. `docs/reference/pr34-collection-formats-2026-09-10.md` holds the shape.
5. **The live half of the feedback loop has no run yet.** Three things want a measurement: the judge lane of the triage, one live fix cycle, and one review round. All three need the owner's word, and the cycle also needs `AUTOTUNE_FIXER_CMD` and a harvest whose verdicts carry a snapshot.
6. **The next whole deck gate run is still run 19, and it is still outstanding.** It makes the next decks baseline.
7. **OQ-79 needs runtime evidence.** A fresh owned-only build on the deployed app is the way to it.
8. **The owner parked PR-35 and dropped PR-30** (D-655, D-656). M-9, F-97, and F-99 found no case for either one.
9. **PR-36, the reader's verdict as a quality signal** (D-651). It waits for verdicts.
10. **The weekly EDHREC read** falls on 2026-09-14 (D-499, D-565). Run `make meta-refresh`, then `make quality-gate` to a new `QUALITY_GATE_OUT`.
11. **PR-26, the return channels**, waits on OQ-67.

CAUTION: `make revise-gate | tee` hides the exit code. Read the verdict line of the document, never the exit code of a pipe.

The CI step "fake gcs tests" ran no test until 2026-09-10 (D-658). Its filter matched no test name, so it passed as a no-op. It runs `scripts/gcs-check.sh` now: the script seeds the fake GCS from the trimmed snapshot and fails unless `TestLiveFakeGCS` passes by name. `make gcs-check` runs the same script.

A ruleset that requires the `verify` check on `main` is not possible. The repo is private on the free plan, and the rulesets API answers 403 (checked 2026-08-28). The owner reads the checks before a merge.

## The ten most recent sessions

### 2026-09-10d: F-104, the offer reads the bracket of its own turn

Merged as #134. PR-39 followed and merged as #135. Dependabot's #131 to #133 carried the day's dependency bumps, and #133 took a fix for vitest 5 and protoc-gen-es 2.14.1.

**The owner asked whether the commander options were the expected quality.** Session `vDzEKDPnRZntlyhz8Lqg` asked for the best possible deck at bracket 5 from an owned-only pool. It read Lotho, Corrupt Shirriff, Ghalta, Primal Hunger, and Peregrin Took. They were not. The bracket 5 signal of OQ-48 exists since PR-14B, and the turn never handed it the bracket.

**The cause is the turn order.** The turn builds the hints from the slots as restored, and the picked power option lands inside the turn. `readFacts` refreshed the colors, the pool, the sets, and the pair flag, and never the bracket. The pick row fires on the turn the power answer arrives (D-631), so the reorder ran on almost no deployed session.

**The proof cost nothing.** A local run of the pool over both fixture exports reads the same three names with no bracket. With bracket 5 it reads Vivi Ornitier, Aang, at the Crossroads, and Hapatra. The deployed model of that morning holds a signal for 964 commanders. The log of the session shows the power pick and the offer 84 milliseconds apart on turn 2.

**The fix is one interface.** `BracketAware` joins `SlotAware`, `SetAware`, and `PairAware`, and `readFacts` hands the bracket over. The turn test fails without the call and passes with it.

**The session before this one ended without a hand-off.** Its work sits on `builder-basics-by-pips`: the source count of F-103, the balance phase of F-102, and D-659 to D-661. The measurement of D-661 ran in this session, for nothing. With the fixed count the pass moves four decks of run 18, and the worst color rises on all four. Deck 22 gains two Forests, where the old count cut two. PR-39 is open.

**The `decktome` gcloud configuration pointed at the Wallabee account and project.** The session read the deployed project through an environment override and changed no configuration.

### 2026-09-10c: PR-37, the synergy check

Merged as #127.

**The owner asked what the merit is, when no reader-facing number moved.** The answer: the bar can now judge a change, and the gate prints the three numbers for free. The model grades the built decks as before, and the judge agreement of 8 of 25 is still the gap a reader feels.

**A free test then cut the Commander casual lists to a half and a quarter** (F-97). The built decks graded bad read 15, 15, and 16, and the judge agreement 8, 8, and 7. The owner chose M-9 before PR-35 (D-654).

**M-9 refuted the premise of D-650** (F-98). With the casual features of PR-29, a larger casual corpus read worse on both reader-facing numbers. The owner parked PR-35 and chose PR-30 next (D-655).

**PR-30 failed in both forms** (F-99). With no self-match no built deck moved, and with it graded bad rose from 15 to 20. The owner dropped PR-30 and chose a structural change to the model, which M-10 measures first (D-656).

**M-10 measured a softmax scorer, and the detector held the built decks** (F-100). The owner asked for the ten disputed decks as a review sheet, to grade them by hand (D-657).

**The check measures the move after the break** (D-652). In each fold the fit reads the synergy feature of every synergy copy, and it drops a copy whose fall sits under the floor. A request that passes no check makes no copy.

**The owner chose the floor on six measured runs** (D-653). At 0.10 the model won the dropped pairs 47 percent of the time before the check, a coin flip. Each higher step drops pairs it won 72 to 90 percent of the time.

**The premise was half right.** The check drops 157 of 389 Commander pairs from every band, and not the 54 weak pairs alone. **In Modern the break does not move the synergy feature in most pairs** (F-96), so the check reads Commander alone.

**Gate run 18 reads PASS, and no reader-facing number moved.** The bar reads a new population, and the model is the same model. The gate now reports the three numbers of D-648 for free, and judge lane run 5 joined `main` for them.

**The CI step "fake gcs tests" runs its test now** (F-101, D-658). Its filter matched no test, so it passed as a no-op. `scripts/gcs-check.sh` seeds the fake GCS from the trimmed snapshot and fails unless `TestLiveFakeGCS` passes by name.

### 2026-09-10b: PR-29 closed, and M-8 read the bar

Merged as #126.

**The casual corpus worked and the trade did not.** A control fitted the same day over the same data reads the numbers apart. The Commander synergy axis read 0.70 and 0.81. The built decks graded bad read 15 of 25 and 17, and the judge agreement 8 of 25 and 6. The judge comparison holds no judge noise: one run answered every deck once, and both models read those same answers.

**Two further features failed and went back.** `casual_unseen_share` moved the target axis by nothing and cost three other bars. `casual_pair_share` won one pair of 389, because the fit split it against `casual_synergy` with opposite signs.

**M-8 found the target was never reachable.** A weak precon's cards do not pair in the corpus, so the synergy break removes half of nothing. 54 of 389 Commander pairs carry no signal and read 0.43, so a perfect model reads 0.93.

**The repository already held the rule.** D-485 says a break must be material, and it checks the copies axis and the colors axis. Synergy is the fallback, and it checks nothing (F-95).

**The owner closed PR-29 and answered the three questions of PR-37.** The evidence document stays, because two findings rest on it.

### 2026-09-10: PR-34, the collection formats

Merged as #123.

**The app reads the format out of the file** (D-647). `collections.Detect` reads the header row, and the reader never names the app their file came from. The app refuses a tie between two formats and names both. It refuses a file no format claims with the list it does read.

**F-92 was live, and nobody met it.** The web sent `MANABOX_CSV` on every upload, hardcoded. The app has read an Arena list since D-15, and every one of them failed every row of itself.

**Moxfield reads, on a real export.** The owner exported their own collection, 3193 rows. The documentation of that format disagreed with itself on `Edition` and on `Condition`, and the file settled both.

**The language column was the trap.** `Resolve` refuses a row that is not `en` (D-23), and Moxfield writes `English`. A pass-through reports all 3192 English cards of the export as non-English. The reader then reads that their whole collection failed.

**The measurement.** 3193 rows parsed, none unread, 125 sets. The snapshot of 2026-09-04 resolves it into 3035 entries and 5884 cards. One row stays unresolved, a Japanese printing, and D-23 refuses that one by design.

**One CSV walker serves every format now.** A platform is one signature entry and one row builder.

**OQ-80 is open.** A Moxfield export carries a `Proxy` column, and no other format this app reads carries one. The owner parked the question rather than decide it on no evidence: every row of the one export on record reads False.

### 2026-09-09c: PR-28c, the fix cycle

Merged as #122. The feedback loop is complete: PR-27 harvests, PR-28a writes the files, PR-28b writes the cases, and PR-28c fixes them.

**The cycle.** `scripts/feedback-loop.sh` reads a harvest and writes one case per thumbs down. It proves each case fails, hands the failures to a fixer agent, and proves each case passes (D-645).

**The case and its fix ride in one pull request.** A case is a failing test by design. A case merged on its own turns the gate red on `main` until a fix lands, and the owner chose one pull request over two.

**The cycle confirms before it fixes.** A case that already passes never measured the reader's fault. It stays as a case a change must not flip, and the fixer never sees it. That is D-234 applied to a case.

**The accept rule is sharper than the tuning loop's.** No noise margin and no ratio: every case goes from fail to pass, `make eval-check` shows no flip, and the tree builds. `cmd/case-check` reads the same bars the gate's own verdict reads.

**The cycle answers the review.** `scripts/feedback-review.sh` hands every open finding of `gitar-bot` to the fixer, pushes, and replies on each thread, over three rounds (D-637).

**The frozen list drifted, and a test caught it.** `TestTheFixerPromptAndTheCycleAgreeOnTheFrozenList` found nine paths the cycle enforces that the prompt never named.

**PR-34 joined the roadmap** (D-646, F-91): the app reads two collection formats, and a reader on any other platform can not upload at all.

### 2026-09-09b: PR-28b, the triage and the cases

Merged as #121.

**The triage.** The reason keys of the dialog name the class with no model call (D-643). PR-27 already asks the reader to name the fault, and the keys map onto the classes one to one. The judge reads three cases alone. They are free text with no reason, reasons that cross two classes, and a reader who argues with a rule the owner set. Eight of the ten fixture items cost nothing.

**The classes grew from ten to 22** (F-90). The design note of 2026-09-06 held ten, and the store offers 22 reason keys. Twelve of them named no class, and a verdict of any of the twelve fell through in silence. The chat kind of D-594 alone added five, a day after the note.

**The expectations.** A "must not ask" expectation names a catalog row that must not fire, and the question gate reads it as one more miss (D-644). The row comes from the question id, which reads `q<n>-<row>`. The deck gate reads two case assertions: a card the build must not pick again, and a deck the reader owns whole.

**The store.** A verdict about a deck keeps the session that built it (D-643), because the theme, the pool rule, and the budget live on the session. Two snapshots that pass the room of one Firestore document drop the session and keep the deck.

**The repository.** D-640 is gone (D-642). The owner permits a reader's exact words in a committed file, so the harvest commits and a case joins the gate file that owns it. The pull request diff is the accept step.

### 2026-09-09: the review process, the user record, and a public repository

Seven pull requests merged, #113 to #119.

**The question workflow.** Gate run 43 read FAIL and found a regression of D-630. `Skip` fills a key and `CloseStalled` does not. So the commander offer waited for a power the net closed. A `Requires` slot reads the skipped keys now (D-631). Run 44 reads PASS.

**The decks.** The upgrade probe measured D-628 for $0.32 and refuted the estimate of D-629: the bands answered 1 of the 8 upgrade findings, not 7. The Turtle Power precon leaves 18 free slots and the role minimums need 17, so an in-band upgrade adds no card of the theme. The owner parked F-85 for want of a second precon (D-632).

**The money.** The PR-22 gate closed on a measured cost at idle of $9.54 a month gross (D-633). The snapshot cron held 80 percent of it, and every tick paid a container start to learn there was nothing to do (F-86). The tick runs hourly now, which leaves $3.81 (D-634).

**PR-28a.** PR-28 split into three (D-636). A verdict keeps the object it names, because a reader deletes the deck or the chat they complained about (D-635). `make feedback-harvest` writes a dated document and a JSONL file, and the watermark comes from those files.

**The user record.** `users/<uid>` holds the email, the dates, and six counters (D-638). The owner ran the backfill.

**The process and the repository.** `gitar-bot` reviews every pull request, and the session answers each finding before it asks for a merge (D-637). The repository went public: every pull request runs the whole workflow (D-639), and no harvest file enters the repository (D-640). The first pull request on Actions found two broken checks. No pull request ran either one before that day (D-641, F-89).

### 2026-09-08: PR-32, PR-33, PR-25, and four walks

Sixteen pull requests merged, #97 to #112. PR-32 closed F-75 and F-76, the commander a reader names and the ownership of a commander (D-606 to D-608). PR-33 landed a build in band with one model call, where the same request had read 3 minutes 59 seconds and an error (F-77, F-78, D-613). Deck gate run 18 measured it and became the decks baseline: repair turns fell from 12 of 25 to 3 of 25 (D-617). PR-25 made the app installable on a phone, and the owner walked it from the Home Screen (D-621 to D-627). The walks found F-79 to F-83, and every one is fixed.

### 2026-09-07: the deploy, the invite gate, and the dead conversation

The app went live on `decktome.com` (D-574). A merge to `main` deploys itself on Cloud Build, through two triggers that read the diff (D-584, D-586, D-588). An uninvited reader had read the whole app, and an invite gate holds every page now (D-589, D-590, D-592). The classifier wrote over the card pool the reader chose, and the reader's choice wins now (D-591). A picked option sets its slot as data, and no turn is silent (D-597 to D-599). M-7 landed the honest bars and the axis diagnostic (D-570 to D-573).

### 2026-09-06: PR-22, PR-23, and PR-27

PR-22 landed the invite list, the spend cap, and the hosting block (D-550). PR-23 landed the Playwright smoke flow, where the fake serves a build (D-551 to D-553). PR-27 landed the feedback harvest: the thumbs, the dialog, the store, and the service (D-561). The mobile and engagement proposal set Stage A as PR-25 (D-546 to D-548).

## The archive

`docs/reference/session-handoff-archive.md` holds every record this file carried before 2026-09-09. It holds the resume section of 2026-09-08, the records of 2026-08-31 to 2026-09-05, and 42 more sections, word for word. Read it for the detail behind a decision.
