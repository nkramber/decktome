# Session hand-off

`CLAUDE.md` is the entry point. It sends you here. Then read `docs/decisions.md` from D-297, `docs/owner-questions.md`, and `docs/open-questions.md`.

This file holds the ten most recent sessions. Every older record sits in `docs/reference/session-handoff-archive.md`, word for word. The first narratives sit in `docs/reference/session-log-2026-08-23-to-26.md` and `docs/reference/session-log-2026-08-27-to-28.md`. The audit of 2026-08-29 is `docs/audit-2026-08-29.md`.

CAUTION: the web tests need Node 22.23.2 (`.nvmrc`). Under Node 20 every test file fails at start with `ERR_REQUIRE_ESM` from jsdom 30. Put `~/.nvm/versions/node/v22.23.2/bin` on the PATH before `make verify`. On this machine `nvm use` reports the change and does not make it, so prepend the path yourself.

## RESUME HERE (2026-09-09)

**The checkout.** `main` is `97fbdee`, which is pull request #119. Branch `pr-28b` holds this session's work, and it waits for a review and a merge. Run `make where` before you touch anything. Never commit on `main` (D-583). Answer the review of `gitar-bot` before you ask for a merge (D-637).

**The app is live on `decktome.com`.** The deploy of `97fbdee` passed, and Cloud Run revision `mtg-api-00026-srg` serves the API. A merge to `main` deploys itself on Cloud Build (D-584, D-586).

**PR-28b waits for a merge** (D-642 to D-644, F-90). The triage names one of 22 classes for each thumbs down, and the reason keys answer 20 of them for nothing. The judge reads three cases alone. They are free text with no reason, reasons that cross two classes, and a reader who argues with a rule the owner set. `docs/reference/pr28b-triage-gate-2026-09-09.md` is the gate, and it reads PASS for nothing.

**A committed file takes a reader's own words** (D-642, supersedes D-640). The owner set that rule on 2026-09-09: a reader's exact words fit anywhere, and no file holds an email. So `.gitignore` no longer holds `docs/reference/feedback/`, and a triage case goes straight into the gate file that owns it. The pull request diff is the accept step, because the owner merges every pull request.

**A verdict about a deck now keeps the session that built it** (D-643). The theme, the pool rule, and the budget live on the session and never on the deck, and a deck gate prompt needs all three. The feedback schema is version 3.

**PR-28c is next.** It is the fix cycle on `scripts/autotune-fix.sh`, under the $2 cap of D-559.

**Every user has a record** at `users/<uid>` (D-638). It holds the verified email, the creation date, and the last active time. It holds six counters: decks, deck revisions, collections, chats, and the verdicts up and down. Every counter counts a creation, so a deleted deck does not lower one. The owner ran `make users-backfill` on 2026-09-09.

**The whole PR-22 gate holds** (D-633). The measured cost at idle reads $9.54 a month gross, and the hourly snapshot tick of D-634 leaves $3.81.

**What waits on the owner.** OQ-67, the Stage B channels. OQ-77, the blocking function of Identity Platform. OQ-79, the commander of an owned-only pool.

**CAUTION: the evidence OQ-79 waits for is not on the deployed project.** `make read-session SESSION=23rplEQAMA0mtJ3QFtKO` answers `no sessions ... over 1 user(s)`, and the deck it names is gone as well. The script reads every user of the project. So a fresh owned-only build on the deployed app is the way to the evidence.

**CAUTION: a local `make verify` is not the whole story.** It read green for weeks while shellcheck failed, because the recipe took the `||` branch on a finding (F-89, D-641). Every pull request runs the workflow now, and the two answer the same question.

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

- Twelve targets and the loop script spend money: `make questions-gate`, `make questions-eval`, `make eval-calibrate`, `make deck-gate`, `make bracket-gate`, `make revise-gate`, `make chat-probe`, `make generate-probe`, `make summary-judge`, `make quality-judge`, `make test-smoke`, `make feedback-triage`, and `scripts/autotune.sh`. Ask the owner before each run. `make autotune`, `make feedback-triage-dry`, and `eval sweep -dry` are free.
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
- Baselines, in `docs/reference/eval/baselines.json`: questions is run 42, decks is run 18, revise is run 9. Run 44 is the newest whole questions run, and run 43 records the regression of F-84. `make eval-check` compares the newest whole run of a suite against its baseline.
- Toolchain: Go 1.27.0, Node 22.23.2, pnpm 9.2.0, firebase-tools 14.14.0, Java 17, Playwright 1.62.1 with its Chromium headless shell (`playwright install chromium`). The first three were verified 2026-09-09.

## Next steps, in order

1. **PR-28c, the fix cycle** (D-557 to D-559, D-636). This is next in the sequence, and PR-28b comes before it. The fixer agent of the tuning loop runs on a branch of its own with the new cases. `scripts/autotune-fix.sh` takes them the way it takes an eval report today. One cycle stops at $2 of gate runs (D-559), and every guard of `docs/reference/autotune-design.md` holds.
2. **The live judge lane of PR-28b has no run yet.** The three verdicts on record predate the snapshot of D-635, so they carry no context. The fixture proves the dry lane alone. The first live run belongs with the first real harvest that carries a snapshot. `make feedback-triage TRIAGE_OUT=<document>` runs it, and it costs a few cents a verdict.
3. **The next whole deck gate run is still run 19, and it is still outstanding.** It makes the next decks baseline, and no deck-build change waits on one now. The upgrade probe of 2026-09-09 already measured D-628 (D-632).
4. **OQ-79 needs runtime evidence.** No fix goes in without it, and neither the session nor the deck it names is readable. A fresh owned-only build on the deployed app is the way to it.
5. **PR-29, the casual corpus**, stands on the weak-axes plan (D-567, D-568). Read `docs/reference/weak-axes-2026-09-07.md` first. No bar moves (D-486), and no weight changes by hand. PR-30 follows, and PR-31 parks (D-573).
6. **The weekly EDHREC read** falls on 2026-09-14 (D-499, D-565). Run `make meta-refresh`, then `make quality-gate` to a new `QUALITY_GATE_OUT`.
7. **The meta job runs already.** The `mtg-meta` job and the `mtg-meta-schedule` cron at 06:00 UTC both run, and the schedule reads ENABLED. A run takes about 1170 seconds and costs $1.83 a month (2026-09-09). No infra file in this repo holds either schedule (D-492).
8. **PR-26, the return channels**, waits on OQ-67.

CAUTION: `make revise-gate | tee` hides the exit code. Read the verdict line of the document, never the exit code of a pipe.

CAUTION: the CI step "fake gcs tests" filters on `LiveStore`, and the only live test is `TestLiveFakeGCS`. The step matches no test and passes as a no-op. With the right name it needs a seeded snapshot bucket, which the CI server lacks. A separate change fixes it, on the owner's word.

A ruleset that requires the `verify` check on `main` is not possible. The repo is private on the free plan, and the rulesets API answers 403 (checked 2026-08-28). The owner reads the checks before a merge.

## The ten most recent sessions

### 2026-09-09b: PR-28b, the triage and the cases

Branch `pr-28b`, open for review.

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

### 2026-09-05: the terse conversations and the trimmed snapshot

The 47 terse conversations joined the bar, with four classifier rules and the GCP deploy guide (D-522, D-534 to D-539). The precon check ignores basic lands, and the M-5 sheet is complete (F-35, D-523, D-529 to D-532). A trimmed card snapshot of at most 10 MB serves the free dry lane of the deck gate (D-521, D-542).

### 2026-09-04: PR-15, the eval harness

PR-15 landed the eval harness (#65). The paid gate ran with it (D-514, D-515). The corpus took the MTGJSON skip and the precon exclusion prompt. Three set fixes closed F-40 to F-42: the Marvel set question, the group set request, and the partial run.

### 2026-09-03: four product pull requests

PR-14B landed the deck quality model (#58). PR-24 landed the precon exclusion (#59). PR-14C landed MTGTop8 and the casual 60-card decks (#60). PR-20 landed the deck view and the card detail (#61), and PR-21 landed the share link and the print view (#62). Two deck gate fixes followed: the precon share reads the pool, and a failed repair keeps the deck.

### 2026-09-02: PR-19 and PR-14A

PR-19 landed the chat and build experience, the counted land swap, and the split land bucket (#55). Its follow-up made a deck count its commander (#56). PR-14A landed the bracket profile (#57). Bracket gate run 1 and deck gate run 12 ran against them.

### 2026-09-01: PR-17B and PR-18

PR-17B landed the set filter and the fixes three gate runs found (#50). PR-18 landed collection management (#53), and its review fixes made the chat refuse a turn with no card index (#54). The question-quality pass ran the same day (D-387 to D-389).

## The archive

`docs/reference/session-handoff-archive.md` holds every record this file carried before 2026-09-09. It holds the resume section of 2026-09-08, the record of 2026-08-31, and 42 more sections, word for word. Read it for the detail behind a decision.
