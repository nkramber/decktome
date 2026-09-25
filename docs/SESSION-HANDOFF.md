# Session hand-off

`CLAUDE.md` is the entry point. It sends you here. Then read `docs/owner-questions.md` and the decisions that the change touches, by D-id (D-749).

This file holds the current state, the resume steps, the facts that expire, the next steps, and the three most recent sessions. `make context-budget` fails when this file passes 24,000 bytes, or its resume section passes 6,000 (D-749). Every older record sits in `docs/reference/session-handoff-archive.md`, word for word. The first narratives sit in `docs/reference/session-log-2026-08-23-to-26.md` and `docs/reference/session-log-2026-08-27-to-28.md`. The audit of 2026-08-29 is `docs/audit-2026-08-29.md`.

CAUTION: the web tests need Node 22.23.2 (`.nvmrc`). Under Node 20 every test file fails at start with `ERR_REQUIRE_ESM` from jsdom 30. Put `~/.nvm/versions/node/v22.23.2/bin` on the PATH before `make verify`. On this machine `nvm use` reports the change and does not make it, so prepend the path yourself.

## RESUME HERE (2026-09-25b)

**Pull request #229, PR-79, corrects twenty-one findings of the repository review of 2026-09-24 (D-920 to D-924).**

Author provider: Claude Code

**The next step.** Next step 2: the next findings of the review report. The P3 findings that remain come next.

**The base.** `main` is `8aec379`, from #228.

**The change.** The owner asked for as many corrections as one session can finish (D-920). Each code finding has a regression test that fails on the base.

- REV-008: a review record on a head from a fork fails RG 3. The owner chose the fork check alone, and no signed record (D-920).
- REV-029: step 6 of the fix cycle says what `make eval-check` proves, and asks for a whole run before the merge (D-921).
- REV-030, REV-043, REV-044: a build holds a lease on its chat in the store. The cap refusal comes before the chat exists. The cap and the size refusals use FailedPrecondition, and a conflict uses Unavailable (D-922).
- P3: REV-038, REV-039, REV-048, REV-056, REV-066, REV-074, REV-075, REV-084, and REV-086 (D-923).
- Seven more P3, on the owner's word: REV-065, REV-076, REV-077, REV-082, REV-087, REV-089, and REV-094 (D-924).
- No paid target ran. No production write.

**The checks.** See the pull request body. `TestEmulatorListRenameDelete` fails on a reused emulator on the base too: it never deletes one of its sessions.

**The review.** Gitar found one issue on `32c6267`: the lease ran on the client context, so a client that left lost the paid turn. The next commit takes it detached, and `TestTheLeaseOutlivesTheClient` failed before it. Gitar approved `4d51d4c` and confirmed the fix at `a8efdf3`. The Codex review of `4d51d4c` found P2-1: a delete read the lease outside its transaction. The next commit reads it inside, for the chat delete and the deck delete. The repeat review of `a8efdf3` closed P2-1 and read `Blocked` on D-921. That rule applies to a fix cycle alone, so the earlier verdict had no basis. The repeat review found P2-2 at `da63923`: a live triage can mark a harvest processed when a judge call fails. `c427583` records each verdict it applies, and a failed verdict keeps its harvest pending. The current review closes P2-2 and finds P2-3: a later write failure can cause a duplicate case on retry. The review record holds the checks and the current verdict.

**What waits on the owner.**

- The merge of this pull request.
- After the deploy, eleven calls of `CheckInvite` with a new first address each. The eleventh must fail (D-907).
- A check of a frame rule and a script policy on a Hosting preview channel (D-923). It is a deploy, so ask first.
- Next step 3 (D-750). M-19 comes after it.
- A whole deck gate run of prompt version 16 and the fixing floor of F-33. Ask first.
- A questions gate run for the rules of D-913, and a deck gate run for the repair input of D-916. Ask first.
- UNVERIFIED: the Moxfield import of the deck list that the app exports.
- Next steps 5 and 6, and OQ-67 and OQ-77.

## How to resume

1. Load the `one-pr-one-session` skill, and do its start gate. A session works on one pull request (D-746).
2. Run `make where`. It prints the branch, the tree, and the state of the branch's pull request.
3. Run `ps aux | grep autotune` before any write. The loop resets the tree when it rejects an iteration.
4. Make a branch from `main`. Never commit on `main`, and never push to it (D-583). The pull request merges itself on the green light (D-828). Run `make hooks` one time in a fresh checkout.
5. Load the skills. Load `ste-writing` before you write any `.md`. Load `design-doc-style` before you edit the roadmap. Load `mtg-corpus` before you reason about a format, a legality, or a card term.
6. Run `make verify`. It runs every check the verify workflow runs, on this machine, for nothing. Put Node 22.23.2 on the PATH first. The pull request runs the same jobs on Actions, and a change of documents alone skips six of them (D-818).
7. Do "Next steps, in order" below. Ask questions as they come up. Record each owner answer in `docs/decisions.md`, and delete the row from `docs/owner-questions.md`.
8. Open the pull request with the sections of `.github/pull_request_template.md`, and run `make pr-check`. Load the `gitar-review` skill, and follow it after each push (D-637, D-745). After Gitar, run `make codex-review PR=<n>` in the background (D-823). After the approval, ask the owner. Put the summary of four sections inside the question: What, How, CI, and Codex review. Turn on the auto-merge after the confirmation (D-828, D-834, D-836). A pull request of documents alone takes the `review-override` label in place of that review (D-812).
9. Update this file inside the pull request, before you call it ready (D-747). Read only the section that you change.
10. Keep three session records at most (D-749). Move each older record to the archive, word for word.
11. When the pull request merges, write the transitional prompt of the `one-pr-one-session` skill, and end the session.

A second Mac: `docs/setup-second-mac.md` holds what to carry, what to install, and how to prove the machine.

Twenty-two things a fresh session gets wrong without this file.

- A test card index with no Oracle text and no tag matches no theme. The theme row of D-725 then asks, and the build never starts. Give each fixture card its real text.
- A singular creature-type word keeps the generic rule, and its plural reads the type row (D-731). So "zombie" and "zombies" read two different lists, and F-144 records why.
- A stored deck records the size of its shortlist and no card of it. So a card that never reached the shortlist and a card that the model dropped look the same. Replay the shortlist for free before a prompt fix (M-17). `.local/m17/zz_scratch_m17_test.go` holds the method, and `list.Theme` names the theme words that matched no card.
- `make bracket-gate` runs its tool in `go/`, through `go -C go`. A relative `-rejudge` path then points inside `go/`, so pass an absolute path.
- Fifteen targets and two loop scripts spend money: `make codex-review`, `make questions-gate`, `make questions-eval`, `make eval-calibrate`, `make deck-gate`, `make bracket-gate`, `make sixty-gate`, `make revise-gate`, `make chat-probe`, `make generate-probe`, `make summary-judge`, `make quality-judge`, `make test-smoke`, `make feedback-triage`, `make api-build`, `scripts/autotune.sh`, and `scripts/feedback-loop.sh`. Ask the owner before each run. `docs/reference/paid-targets.md` holds the cost and the guard of each. `make autotune`, `make feedback-loop`, `make feedback-loop-dry`, `make feedback-triage-dry`, `eval sweep -dry`, and `go run ./cmd/bracket-gate -sweep` are free.
- A rerun writes to a new file. Every `*_OUT` variable refuses a document that holds a result (D-65).
- A gate document names the commit of `HEAD`, and never the tree. Deck gate run 29 ran over uncommitted work, so its header names the parent commit `5fd8085`. Commit the change before a paid run.
- A gate run takes about 20 minutes and an eval about 13. A foreground command stops at 10 minutes, so run both in the background.
- `docker compose up` needs provider keys in `.env` now, and fails fast without them (D-267). `make dev` still starts with no keys, and the fake builds the fixture deck of PR-23 for any first message (D-552).
- The eval and the agent share a model. Every ratio it reports is a floor, not a measurement (D-136).
- The judge is noisy: two runs of identical code move up to nine bad questions (D-230). One run proves nothing on its own.
- The sweep's estimate of a step is the cost of its last run file. The deck gate estimate read $2.54 from run 14, and run 16 cost $3.79 with the plan judge and 12 repair turns. A cap set from the estimate stops the sweep before its last step.
- Many local fits and verify runs fill the disk. On 2026-09-11 `make verify` failed at link time with 258 MiB free, when the Go and Docker build caches held about 50 GB. Run `df -h /System/Volumes/Data` before a long run, and ask the owner before you clear a cache.
- A Gitar review can be stale: a paused Gitar keeps the review of an older commit, and the comment names no commit. The `gitar-review` skill proves that a review is current, and it holds every Gitar trap (D-745).
- A background shell command starts in the directory the session left. On 2026-09-12 a `make` target ran in `go/` and found no rule, so head every command with an absolute `cd`.
- A live check of a web change reads the stored session and the deployed chunk, and not the screen alone. The service worker served the old shell for one load after a deploy (F-122). The protobuf-es code holds each field name in base64, so search a chunk for a property name such as `noDecline`.
- The hook `.claude/hooks/session_bind.py` binds a session to the first branch that it creates, pushes, or opens a pull request for. A command on a second branch exits with "Blocked". Start a new clean session. The owner alone removes a binding under `.git/decktome-session-bind/`.
- A read under `users/` fails in the sandbox of a session, because that path holds the verified email (D-638). A field mask that fetches no field fails too. A replay reads a local collection export instead (D-756), and the export never enters git.
- A stored session can be gone. `z1hshyY6Npig1FN2NuV7` no longer exists in `decktome-prod`, so `scripts/read-session.sh` finds nothing. M-17 recorded the request of that session, so the replay needed no session read.
- The Edit tool and a Bash heredoc can write the rune in place of the six characters `\u2014`. Write the backslash as `chr(92)` in a Python script, and read the bytes with `ascii()`.
- The `review-gate` check and the job `verify:skip` read their rules from `main`. A pull request that changes `review_gate.py` or `ci_skip.py` runs the old rule on its own head, so prove a change in the tests (D-816, D-820).
- `Pool.Names` sorts the pool by the alphabet, and the shortlist groups its cards by role. Neither order ranks a card. `Pool.Score` holds the shortlist score (D-702).

## Facts that expire

- Comprehensive Rules: the current file is 2026-08-19 (D-272). The rule citations in the corpus match the 2026-08-07 text. The session of 2026-09-12 read rules 709.4c, 710.2, 712.8a, 712.12, 715.2, and 722.2a in that file for PR-44.
- Ban-list snapshot: 2026-08-24. Next announcement 2026-10-12, in `internal/cards/announcement_dates.json`. A test fails when that file holds no future date. This is by design.
- Commander brackets: the 2025-10-21 revision. Game Changers: 53 cards, list of 2026-02-09. Lutri is banned as a companion only, per the 2026-02-09 announcement (`companion_bans.json` holds the link). The content rules per bracket in `brackets.json` and the Spellbook thresholds were read 2026-09-02, and the Karsten tables are the 2022 articles, read 2026-09-02 (`docs/reference/bracket-profile-2026-09-02.md`).
- Standard: 18 sets, Wilds of Eldraine to The Hobbit. Six sets leave at the first 2027 set. Verified 2026-08-24.
- Card snapshot on disk: `.local/gcs/mtg-local-cards/scryfall/20260904T210157`, the newest of 13. Read 2026-09-13, when it flagged 53 Game Changers. Check every card fact against it.
- The theme table `themes.json` reads `verified_at` 2026-09-20. `make themes-check` read every slug against the snapshot of 2026-09-04 on 2026-09-20. The question gate set holds 109 conversations: 78 counted and 31 probes (D-730).
- LLM model ids and prices: `roles.json` and `prices.json`, verified 2026-08-24. The max output per provider in `llm/client.go`, verified 2026-08-29. The OpenAI rows are unverified by anyone but the owner. The judge role runs on Opus 5 (D-430).
- The owner repaired the application default credentials of this Mac on 2026-09-23, after the failure of 2026-09-13. The harvest then read `decktome-prod`, and the sandbox let it read the feedback of each user.
- CAUTION: the `decktome` gcloud configuration named the Wallabee account and project on 2026-09-10, so `use_decktome` put the shell on the wrong project. `scripts/read-session.sh` reads `SESSION_PROJECT` and the account of `CLOUDSDK_CORE_ACCOUNT`, so an environment override reads `decktome-prod` with no change to the configuration. The owner has the commands to repair the configuration. A read on 2026-09-12 at 19:24 UTC found the configuration unchanged.
- The backfill of 2026-09-09 read one user with a record to seed: 1 collection, 1 thumbs up, and 2 thumbs down. It counted no deck and no chat, because the reader deleted both (D-635).
- The harvest of 2026-09-24 at 15:28 UTC read 5 verdicts, all of one user. Two carry the snapshot of D-635, of 2026-09-09 and of 2026-09-24. `docs/reference/f49-harvest-2026-09-23.md` holds the first harvest. A collection group query over it needs an index for its shape, and the store holds "verdict ascending, created_at descending" alone.
- The deployed schedules, read 2026-09-09 and again on 2026-09-11: `mtg-snapshot-schedule` at `0 * * * *` (D-634) and `mtg-meta-schedule` at `0 6 * * *`. Both read ENABLED. The API service holds minScale 0, so it scales to zero. No billing export exists, so no command reads the billed spend.
- The deployed API, read 2026-09-24: Cloud Build `58dfbace` of `deploy-api` built `930d1a6`, from #224, and ended SUCCESS at 19:53:44 UTC. Build `86152b16` of `deploy-web` built `930d1a6` and ended SUCCESS at 19:56:27 UTC. The session then made revisions `mtg-api-00077-vwp` and `mtg-api-00078-hv2` on the same image, to mount the Pushover secrets (D-893). The service holds no minimum instance, so a cold start reads Unavailable for about 90 seconds (F-164).
- The deployed web app, read 2026-09-20 at 22:45 UTC: the release of #196, Hosting version `60c80b8c4ad2a689` of 22:44:09 UTC. `index.html` loads `assets/index-Fi2SZrwg.js`. The deck page chunk `deck-view-BQeBlu0i.js` holds the power counts, and `use-cards-DNxWsKmD.js` holds every label. The release before it, `b786ceb7f89bc4fc` of 2026-09-13, came from #158.
- The local quality model is `20260923T202806Z`, from `make meta-refresh` on `1089438` on 2026-09-23. It holds the commander rates of 231 commanders (D-839).
- The deployed quality model, read 2026-09-24: `20260924T061156Z`, from the meta job that ran from 06:00:50 to 06:19:52 UTC on `worker:fa05e3c`. Its Commander fit read 25,851 lists and used 9,244, with `immaterial` 224 and accuracy 0.5595. It holds the rates of 225 commanders, and Najeela reads 64 lists and 169 cards (D-839). The EDHREC pass read nothing, because it ran on 2026-09-21. The mtgjson source read no list, because its deck list version differs from the stored table.
- The Karsten land article of 2022-07-29, read 2026-09-11 through `infinite-api.tcgplayer.com/content/article/<id>/`, because the page draws its text in the browser. `docs/reference/m12-rules-diagnostic-2026-09-11.md` holds the formula, the error, and the cheap rules.
- Baselines, in `docs/reference/eval/baselines.json`: questions is run 42, decks is run 19, revise is run 9. The generate prompt reads version 16 now, and run 19 read version 12. Run 34 of 2026-09-21 is the newest whole deck gate run. It rejudges the summaries of run 31 (D-789), and `make eval-check` reads it as PASS against run 19. Run 52 is the newest whole questions run, and it reads PASS. Runs 45 to 47 read FAIL (D-669, F-112), and run 43 records the regression of F-84. `make eval-check` compares the newest whole run of a suite against its baseline. The quality gate has no baseline row, and run 23 is its newest run.
- Toolchain on this Mac, read 2026-09-11: Go 1.27.1, Node 22.23.2, pnpm 9.2.0, firebase-tools 14.14.0, and Java 17.0.20.1. Playwright is 1.63.0 with its Chromium headless shell (`playwright install chromium`). `go.mod` asks Go 1.27.0 or newer, and vitest is 5.0.0 since #133.
- The local meta store, read 2026-09-14: 17,686 TopDeck good and 6,431 top-cut Commander lists from 2026-06-04 to 2026-09-14. It also holds 1,045 EDHREC average decks and 192 MTGJSON precons. M-17 read the TopDeck lists from 2026-08-01 alone.

## Next steps, in order

1. **PR-79: twenty-one findings of the repository review of 2026-09-24** (D-920 to D-924). This pull request is #229.
2. **The next finding of the review report** (D-901). The report sits at `.local/reviews/repository-review-2026-09-24.md`, outside git. Its section 8.2 gives the order. PR-77 to PR-79 corrected each P1 and P2 finding. The P3 findings without the mark `COMPLETE` come next. Mark each finding `COMPLETE - PR #N` after the Codex approval. The owner answers of its section 9 hold for each finding.
3. **Measure five sessions on the new files, then ask the owner about the checkpoint rule** (D-750). The method sits in `docs/reference/context-budget-2026-09-16.md`. The rule moves a session past about 300K tokens of context to a new clean session. That session continues the same pull request. It comes after PR-75 (D-890). It waits, because the ten sessions of the audit ran before #184 and before D-749.
4. **The open items of the roadmap.** Two register rows read 🔧: F-49 and F-174. PR-73 ran the live cycle of F-49 (D-880). F-166 reads ✅ (#217, D-839). F-168 and F-169 read ✅ (#219, D-843, D-844). F-170 and F-172 read ✅ (#221, D-870). F-171 reads ✅ (#220, D-861). OQ-85 waits for a shape score of the shortlist (D-773). F-157 reads ✅ (D-762, D-763). D-805 retires the power pass after the build (D-704). F-137 stays a record (D-718). A wider theme guard than D-535 waits for evidence (D-783). Bracket gate runs 9 to 11 read the fixing floor and prompt version 16 on prompts 10 to 15 alone. Run 11 read prompts 13 to 15 with the commander rate. No whole deck gate run measured them yet.
5. **Read one deployed session whose theme matches no card, such as "anime"** (PR-54). The theme row must ask before the build. The owner builds it, and a session reads it with `scripts/read-session.sh`. A new session holds no copy of the collection export. The collection sits at `users/<uid>/collections/<id>` in `decktome-prod`. `scripts/read-session.sh z1hshyY6Npig1FN2NuV7` prints the user id and the collection id. Its field `entries_gz` holds gzip JSON of the entries. Keep the exported collection out of git.
6. **Read the first commander question on the app after one more load** (D-690). It waits for the owner. The server half holds: session `vY1lCRtl64uwFObznCZ9` stores the flag. The question must show "Suggest one" and no "You decide", and the pick row must still show "You decide". Session `z1hshyY6Npig1FN2NuV7` named its commander, so it showed no such row.
7. **The next collection platform, when the owner names one** (F-91). Five are left: Archidekt, Deckbox, Delver Lens, TCGplayer, and Helvault. Each one takes a real export and never a column list. `docs/reference/pr34-collection-formats-2026-09-10.md` holds the shape.
8. **M-19: replay the owned-only shortlist of the thumbs down of 2026-09-24** (F-174, D-881). It is free. Read the session and the deck of the snapshot, and the local collection export. Keep the export out of git.
9. **Deck gate run 19 is the decks baseline** (D-686). Its grades read the local stored model `20260910T012734Z`, and the deployed app reads a newer one. Run 29 of 2026-09-20 cost $2.7384 in 1874 seconds, and it reads PASS. Ask the owner before the next whole run.
10. **The owner parked PR-35 and dropped PR-30** (D-655, D-656). M-9, F-97, and F-99 found no case for either one.
11. **PR-36, the reader's verdict as a quality signal** (D-651). It waits for verdicts.
12. **PR-26, the return channels**, waits on OQ-67.
13. **PR-42 is merged as #148** (D-671). Question gate run 49 missed no conversation, so it ran no rerun. Any miss still fails the run, and each miss joins the finding register.

CAUTION: a command such as `make revise-gate 2>&1 | tee log` hides the exit code of make, because the shell of a session sets no pipefail. Read the verdict line of the document. Each recipe of `Makefile` sets pipefail itself (D-782).

CAUTION: `make verify` runs `eval-check`, and `eval-check` reads the newest whole run of each suite. A committed FAIL run turns verify red until a newer whole run passes.

The CI step "fake gcs tests" ran no test until 2026-09-10 (D-658). Its filter matched no test name, so it passed as a no-op. It runs `scripts/gcs-check.sh` now: the script seeds the fake GCS from the trimmed snapshot and fails unless `TestLiveFakeGCS` passes by name. `make gcs-check` runs the same script.

The repository is public (D-639). The rulesets API answers, and the ruleset of `main` can require each job of `verify` (D-828). `docs/reference/merge-rules.md` holds the rules, and `make ruleset-check` compares them with GitHub.

## The three most recent sessions

### 2026-09-24f: the five P1 findings of the review, PR-77

**The owner asked for the corrections of the review report of 2026-09-24, as many as one session can finish.** The session read the owner answers in the report first. The owner picked all five P1 findings (D-901). The two answers on REV-001 conflicted, and the owner chose the author filter and the raw reply (D-902). The owner chose an Admin script for the current accounts, run before the merge (D-903), and Modern for a declined 60-card format (D-904).

### 2026-09-25a: forty-four findings of the review, PR-78

**The owner asked for the next corrections of the review report of 2026-09-24, and five times for more.** The exact trigger of REV-006 met the line cap on the base, so the test holds 990 lines (D-908). The owner approved the provider condition of REV-009 and the Pushover secrets of the jobs for REV-011 (D-906, D-911). No Google page states the shape of the forwarded header, so the key of REV-010 serves both shapes (D-907).

### 2026-09-25b: twenty-one findings of the review, PR-79

**The owner asked for the next corrections of the review report of 2026-09-24, with no rule of one concern.** The owner picked the fork check alone for REV-008, and no signed record (D-920). The base run of the new two-server test started a second paid build and wrote the chat, as the report said. A frame rule of the web host waits, because the auth domain frames its helper page from another origin (D-923).

## The archive

`docs/reference/session-handoff-archive.md` holds every record this file no longer carries. It holds the resume sections of 2026-09-08, and of 2026-09-16 to 2026-09-25a, the records of 2026-08-31 to 2026-09-24e, and 104 more sections, word for word. Read it for the detail behind a decision.
