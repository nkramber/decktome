# Session hand-off

`CLAUDE.md` is the entry point. It sends you here. Then read `docs/owner-questions.md` and the decisions that the change touches, by D-id (D-749).

This file holds the current state, the resume steps, the facts that expire, the next steps, and the three most recent sessions. `make context-budget` fails when this file passes 24,000 bytes, or its resume section passes 6,000 (D-749). Every older record sits in `docs/reference/session-handoff-archive.md`, word for word. The first narratives sit in `docs/reference/session-log-2026-08-23-to-26.md` and `docs/reference/session-log-2026-08-27-to-28.md`. The audit of 2026-08-29 is `docs/audit-2026-08-29.md`.

CAUTION: the web tests need Node 22.23.2 (`.nvmrc`). Under Node 20 every test file fails at start with `ERR_REQUIRE_ESM` from jsdom 30. Put `~/.nvm/versions/node/v22.23.2/bin` on the PATH before `make verify`. On this machine `nvm use` reports the change and does not make it, so prepend the path yourself.

## RESUME HERE (2026-09-19b)

**Pull request #191 carries PR-56, the repair-turn gate of a finding on the commander, and it waits for the owner's merge** (D-762, D-763, F-157). The fix and its measurement cost nothing.

**The next step.** Start a new clean session, and load the `one-pr-one-session` skill. Run `make where`, and read the state of #191 with `gh pr view 191`. Then do next step 1.

**The base.** `main` is `d7a5825`, from #190, which carried PR-55. #189 merged M-18 as `c2d5957`.

**What this pull request holds.** It changes the repair gate of the build, its tests, and the documents. No prompt and no model call changes.

- `go/internal/generate/generate.go` gains `fixable`. It drops every finding that names a commander of the deck, before the build decides on a repair turn.
- The build writes one log line for the findings it drops, and the deck keeps each one (D-226, D-300).
- `go/internal/generate/ownedcommander_test.go` holds two tests, and both fail on the old code.
- **F-157 closes as fixed.** Its register row carries a dated correction of its premise.
- `docs/reference/f157-owned-commander-2026-09-19.md` holds every count.

**The numbers.** A free replay of the M-18 request names the card of the `not_owned` finding. It is Gríma, Saruman's Footman, the commander the reader named, and the export of 2026-09-02 holds no copy. The model named no card outside the collection: the owned-only shortlist drops every card at 0 copies, and `Normalize` reads the pool alone. A control run with the commander covered reads profile findings alone, and no repair turn. The uncapped owned pool holds 475 cards, with 1 Game Changer against a bracket 5 floor of 8, and 1 tutor against 4. The replay deck reads 6 `profile_off_band` findings, and the three model decks of M-18 read 3.

**The paid run.** None. Every count comes from the local snapshot of 2026-09-04 and the local export of 2026-09-02.

**The checks.** `make verify` passed on this machine, exit 0. `make ste-check`, `make ref-check`, `make lifecycle-check`, and `make context-budget` each read 0 findings. Both new tests fail on the old code. The first reads 4 findings where it wants 3, and the second counts 2 provider calls where it wants 1. Every job of the verify workflow and `pr-contract` passed on `91934bb`. The job `verify:changes` skipped, because it runs on a manual start alone.

**The review.** `gitar-bot` approved `91934bb`, and it reads 0 findings and no open thread. The review is current: the head of the pull request matches, and the dashboard comment reads an edit time after the push. CAUTION: the Gitar trial ends about 2026-09-23, from the dashboard of 2026-09-20.

**What waits on the owner.**

- The merge of this pull request, after the review of `gitar-bot`.
- The next item of the roadmap (next step 1).
- Five sessions on the new files, before a decision on the checkpoint rule (next step 2, D-750).
- A deployed session with a theme that matches no card, such as "anime" (next step 3).
- A look at the first commander question after a load of the app (next step 4).
- OQ-67, OQ-77, OQ-80, and OQ-83 to OQ-86.

## How to resume

1. Load the `one-pr-one-session` skill, and do its start gate. A session works on one pull request (D-746).
2. Run `make where`. It prints the branch, the tree, and the state of the branch's pull request.
3. Run `ps aux | grep autotune` before any write. The loop resets the tree when it rejects an iteration.
4. Make a branch from `main`. Never commit on `main`, and never push to it (D-583). The owner merges (D-585). Run `make hooks` one time in a fresh checkout.
5. Load the skills. Load `ste-writing` before you write any `.md`. Load `design-doc-style` before you edit the roadmap. Load `mtg-corpus` before you reason about a format, a legality, or a card term.
6. Run `make verify`. It runs every check the verify workflow runs, on this machine, for nothing. Put Node 22.23.2 on the PATH first. The pull request runs the same jobs on Actions (D-639).
7. Do "Next steps, in order" below. Ask questions as they come up. Record each owner answer in `docs/decisions.md`, and delete the row from `docs/owner-questions.md`.
8. Open the pull request with the sections of `.github/pull_request_template.md`, and run `make pr-check`. Load the `gitar-review` skill, and follow it after each push (D-637, D-745). Gitar is the only review this repo asks for. A pull request of documents alone waits for the review too (D-679).
9. Update this file inside the pull request, before you call it ready (D-747). Read only the section that you change.
10. Keep three session records at most (D-749). Move each older record to the archive, word for word.
11. When the pull request is ready, end the session. The next pull request starts in a new clean session.

A second Mac: `docs/setup-second-mac.md` holds what to carry, what to install, and how to prove the machine.

Nineteen things a fresh session gets wrong without this file.

- A test card index with no Oracle text and no tag matches no theme. The theme row of D-725 then asks, and the build never starts. Give each fixture card its real text.
- A singular creature-type word keeps the generic rule, and its plural reads the type row (D-731). So "zombie" and "zombies" read two different lists, and F-144 records why.
- A stored deck records the size of its shortlist and no card of it. So a card that never reached the shortlist and a card that the model dropped look the same. Replay the shortlist for free before a prompt fix (M-17). `.local/m17/zz_scratch_m17_test.go` holds the method, and `list.Theme` names the theme words that matched no card.
- `make bracket-gate` runs its tool in `go/`, through `go -C go`. A relative `-rejudge` path then points inside `go/`, so pass an absolute path.
- Twelve targets and two loop scripts spend money: `make questions-gate`, `make questions-eval`, `make eval-calibrate`, `make deck-gate`, `make bracket-gate`, `make revise-gate`, `make chat-probe`, `make generate-probe`, `make summary-judge`, `make quality-judge`, `make test-smoke`, `make feedback-triage`, `scripts/autotune.sh`, and `scripts/feedback-loop.sh`. Ask the owner before each run. `docs/reference/paid-targets.md` holds the cost and the guard of each. `make autotune`, `make feedback-loop`, `make feedback-loop-dry`, `make feedback-triage-dry`, `eval sweep -dry`, and `go run ./cmd/bracket-gate -sweep` are free.
- A rerun writes to a new file. Every `*_OUT` variable refuses a document that holds a result (D-65).
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
- `Pool.Names` sorts the pool by the alphabet, and the shortlist groups its cards by role. Neither order ranks a card. `Pool.Score` holds the shortlist score (D-702).

## Facts that expire

- Comprehensive Rules: the current file is 2026-08-19 (D-272). The rule citations in the corpus match the 2026-08-07 text. The session of 2026-09-12 read rules 709.4c, 710.2, 712.8a, 712.12, 715.2, and 722.2a in that file for PR-44.
- Ban-list snapshot: 2026-08-24. Next announcement 2026-10-12, in `internal/cards/announcement_dates.json`. A test fails when that file holds no future date. This is by design.
- Commander brackets: the 2025-10-21 revision. Game Changers: 53 cards, list of 2026-02-09. Lutri is banned as a companion only, per the 2026-02-09 announcement (`companion_bans.json` holds the link). The content rules per bracket in `brackets.json` and the Spellbook thresholds were read 2026-09-02, and the Karsten tables are the 2022 articles, read 2026-09-02 (`docs/reference/bracket-profile-2026-09-02.md`).
- Standard: 18 sets, Wilds of Eldraine to The Hobbit. Six sets leave at the first 2027 set. Verified 2026-08-24.
- Card snapshot on disk: `.local/gcs/mtg-local-cards/scryfall/20260904T210157`, the newest of 13. Read 2026-09-13, when it flagged 53 Game Changers. Check every card fact against it.
- The theme table `themes.json` reads `verified_at` 2026-09-14. `make themes-check` read every slug against the snapshot of 2026-09-04 on 2026-09-14. The question gate set holds 109 conversations: 78 counted and 31 probes (D-730).
- LLM model ids and prices: `roles.json` and `prices.json`, verified 2026-08-24. The max output per provider in `llm/client.go`, verified 2026-08-29. The OpenAI rows are unverified by anyone but the owner. The judge role runs on Opus 5 (D-430).
- CAUTION: the application default credentials of this Mac failed on 2026-09-13 with `invalid_rapt`. `make feedback-list` fails, and so do the harvest and the backfill, which read the same credentials. `scripts/read-session.sh` and `gcloud logging read` still work with `CLOUDSDK_CORE_ACCOUNT`. The owner runs `gcloud auth application-default login` to repair them.
- CAUTION: the `decktome` gcloud configuration named the Wallabee account and project on 2026-09-10, so `use_decktome` put the shell on the wrong project. `scripts/read-session.sh` reads `SESSION_PROJECT` and the account of `CLOUDSDK_CORE_ACCOUNT`, so an environment override reads `decktome-prod` with no change to the configuration. The owner has the commands to repair the configuration. A read on 2026-09-12 at 19:24 UTC found the configuration unchanged.
- The backfill of 2026-09-09 read one user with a record to seed: 1 collection, 1 thumbs up, and 2 thumbs down. It counted no deck and no chat, because the reader deleted both (D-635).
- The feedback store holds 3 verdicts on 2026-09-09, and every one predates the snapshot of D-635. `make feedback-list VERDICT=` reads both verdicts now (F-87). A collection group query over it needs an index for its shape, and the store holds "verdict ascending, created_at descending" alone.
- The deployed schedules, read 2026-09-09 and again on 2026-09-11: `mtg-snapshot-schedule` at `0 * * * *` (D-634) and `mtg-meta-schedule` at `0 6 * * *`. Both read ENABLED. The API service holds minScale 0, so it scales to zero. No billing export exists, so no command reads the billed spend.
- The deployed API, read 2026-09-15 at 04:14 UTC: revision `mtg-api-00050-w9z` on image `api:4dbbc20`, from #179. Both jobs run `worker:4dbbc20`, and `/readyz` answered ok with a card snapshot of 2026-09-14 21:01 UTC. `deploy-api` ran from 04:09 to 04:13 UTC. No web build ran, because #179 changed no file under `web/**`.
- The deployed web app, read 2026-09-13 at 18:54 UTC: the release of #158, from `deploy-web` at 18:52 UTC. `index.html` loads `assets/index-DYfo4m8I.js` and no `registerSW.js`, and `sw.js` precaches the `workbox-window` chunk.
- The deployed quality model, read 2026-09-14: `20260914T070904Z`, from the meta job that started at 06:02 UTC and ended at 07:13 UTC. It fits 43,182 lists and 1,517 commanders. Its Commander fit reads `immaterial` 239 and accuracy 0.554, and its Standard fit reads a cross share of 0.667. The job read the weekly EDHREC pass, 2,077 pages and 4 lists. The mtgo source read 3,094 lists with 58 fetch errors. The mtggoldfish source read 155 lists with no failure. The mtgjson source read no list, because its deck list version differs from the stored table, as on 2026-09-12 and 2026-09-13.
- The Karsten land article of 2022-07-29, read 2026-09-11 through `infinite-api.tcgplayer.com/content/article/<id>/`, because the page draws its text in the browser. `docs/reference/m12-rules-diagnostic-2026-09-11.md` holds the formula, the error, and the cheap rules.
- Baselines, in `docs/reference/eval/baselines.json`: questions is run 42, decks is run 19, revise is run 9. The generate prompt reads version 13 since PR-45b, and run 19 read version 12. So the next whole deck gate run reads a new prompt against run 19. Run 52 is the newest whole questions run, and it reads PASS. Runs 45 to 47 read FAIL (D-669, F-112), and run 43 records the regression of F-84. `make eval-check` compares the newest whole run of a suite against its baseline. The quality gate has no baseline row, and run 23 is its newest run.
- Toolchain on this Mac, read 2026-09-11: Go 1.27.1, Node 22.23.2, pnpm 9.2.0, firebase-tools 14.14.0, and Java 17.0.20.1. Playwright is 1.63.0 with its Chromium headless shell (`playwright install chromium`). `go.mod` asks Go 1.27.0 or newer, and vitest is 5.0.0 since #133.
- The local meta store, read 2026-09-14: 17,686 TopDeck good and 6,431 top-cut Commander lists from 2026-06-04 to 2026-09-14. It also holds 1,045 EDHREC average decks and 192 MTGJSON precons. M-17 read the TopDeck lists from 2026-08-01 alone.

## Next steps, in order

1. **Ask the owner for the next item, in a new clean session** (D-746). The sequence of the roadmap ends at step 46. OQ-83 to OQ-86 hold the other suggestions of the review of the Gríma deck. M-18 gives them evidence. **F-157** records the thin owned pool and the repair turn of every build, and it carries no open question. The power pass after the build still waits (D-704). F-137 stays a record (D-718). PR-55 closes F-148 and the land half of F-144. The nonland half of F-144 stays open: the tag `typal-choose` holds 88 nonlands that no type row reads (D-760).
2. **Measure five sessions on the new files, then ask the owner about the checkpoint rule** (D-750). The method sits in `docs/reference/context-budget-2026-09-16.md`. The rule moves a session past about 300K tokens of context to a new clean session. That session continues the same pull request. It waits, because the ten sessions of the audit ran before #184 and before D-749.
3. **Read one deployed session whose theme matches no card, such as "anime"** (PR-54). The theme row must ask before the build. The owner builds it, and a session reads it with `scripts/read-session.sh`. A new session holds no copy of the collection export. The collection sits at `users/<uid>/collections/<id>` in `decktome-prod`. `scripts/read-session.sh z1hshyY6Npig1FN2NuV7` prints the user id and the collection id. Its field `entries_gz` holds gzip JSON of the entries. Keep the exported collection out of git.
4. **Read the first commander question on the app after one more load** (D-690). It waits for the owner. The server half holds: session `vY1lCRtl64uwFObznCZ9` stores the flag. The question must show "Suggest one" and no "You decide", and the pick row must still show "You decide". Session `z1hshyY6Npig1FN2NuV7` named its commander, so it showed no such row.
5. **Watch the first self-reload on the next web deploy** (D-692). The live release of #158 passed its check on 2026-09-13 at 18:54 UTC. An installed app that loaded that release must reload by itself when the next release activates. Read it on the next merge that changes `web/**`.
6. **The next collection platform, when the owner names one** (F-91). Five are left: Archidekt, Deckbox, Delver Lens, TCGplayer, and Helvault. Each one takes a real export and never a column list. `docs/reference/pr34-collection-formats-2026-09-10.md` holds the shape.
7. **The live half of the feedback loop has no run yet.** Three things want a measurement: the judge lane of the triage, one live fix cycle, and one review round. All three need the owner's word, and the cycle also needs `AUTOTUNE_FIXER_CMD` and a harvest whose verdicts carry a snapshot.
8. **Deck gate run 19 is the decks baseline** (D-686). Its grades read the local stored model `20260910T012734Z`, and the deployed app reads a newer one. The next whole run compares against run 19, and it costs about $2.75, so ask the owner first.
9. **The owner parked PR-35 and dropped PR-30** (D-655, D-656). M-9, F-97, and F-99 found no case for either one.
10. **PR-36, the reader's verdict as a quality signal** (D-651). It waits for verdicts.
11. **PR-26, the return channels**, waits on OQ-67.
12. **PR-42 is merged as #148** (D-671). Question gate run 49 missed no conversation, so it ran no rerun. Any miss still fails the run, and each miss joins the finding register.

CAUTION: `make revise-gate | tee` hides the exit code. Read the verdict line of the document, never the exit code of a pipe.

CAUTION: `make verify` runs `eval-check`, and `eval-check` reads the newest whole run of each suite. A committed FAIL run turns verify red until a newer whole run passes.

The CI step "fake gcs tests" ran no test until 2026-09-10 (D-658). Its filter matched no test name, so it passed as a no-op. It runs `scripts/gcs-check.sh` now: the script seeds the fake GCS from the trimmed snapshot and fails unless `TestLiveFakeGCS` passes by name. `make gcs-check` runs the same script.

A ruleset that requires the `verify` check on `main` is not possible. The repo is private on the free plan, and the rulesets API answers 403 (checked 2026-08-28). The owner reads the checks before a merge.

## The three most recent sessions

### 2026-09-19b: the repair-turn gate of PR-56

**The owner chose the repair gate alone, and the record of the thin pool** (D-762, D-763). A free replay of the M-18 request names the card of its `not_owned` finding. It is the commander the reader named, and the export of 2026-09-02 holds no copy. No answer of the model changes a commander, so each M-18 build spent a repair call on a block that survives it. `fixable` drops such a finding before the build decides. The uncapped owned pool reaches 1 Game Changer and 1 tutor, against the bracket 5 floors of 8 and 4. So F-157 closes with its measurement. A second session wrote this checkout during the work, and the owner stopped it. The session ran no paid target.

### 2026-09-19: the typal land signal of PR-55

**The owner chose the wider signal, kept it to lands, and extended it to a singular type word** (D-759 to D-761). A type row of `themes.json` read no land, so a typal shortlist lost every land that makes mana for its type (F-148). A new `typal_land` block holds the tag `typal-choose` and one needle for Path of Ancestry. Each row adds its own subtype word as a land needle. The five lands of F-148 are on the dinosaur shortlist again, in the plural and in the singular. The session ran no paid target, and every count comes from the local snapshot of 2026-09-04.

### 2026-09-18: the Gríma replay of M-18

**The owner chose the replay that D-727 ordered, over a typal land signal and the four open questions of the Gríma review** (D-755). The session `z1hshyY6Npig1FN2NuV7` no longer exists, and the sandbox refuses every read under `users/`, so the replay read a local ManaBox export of 2026-09-02 (D-756). The owner approved three real builds, and they cost $0.2350 (D-757). Every condition line holds, so F-140 closes as fixed (D-758). F-157 opens on the thin owned pool and the repair turn of every build.

## The archive

`docs/reference/session-handoff-archive.md` holds every record this file no longer carries. It holds the resume sections of 2026-09-08, and of 2026-09-16 to 2026-09-19, the records of 2026-08-31 to 2026-09-17, and 100 more sections, word for word. Read it for the detail behind a decision.
