# Session hand-off

`CLAUDE.md` is the entry point. It sends you here. Then read `docs/decisions.md` from D-297, `docs/owner-questions.md`, and `docs/open-questions.md`.

This file holds the ten most recent sessions. Every older record sits in `docs/reference/session-handoff-archive.md`, word for word. The first narratives sit in `docs/reference/session-log-2026-08-23-to-26.md` and `docs/reference/session-log-2026-08-27-to-28.md`. The audit of 2026-08-29 is `docs/audit-2026-08-29.md`.

CAUTION: the web tests need Node 22.23.2 (`.nvmrc`). Under Node 20 every test file fails at start with `ERR_REQUIRE_ESM` from jsdom 30. Put `~/.nvm/versions/node/v22.23.2/bin` on the PATH before `make verify`. On this machine `nvm use` reports the change and does not make it, so prepend the path yourself.

## RESUME HERE (2026-09-12)

**The checkout.** `main` is `c1c40fa`, which is pull request #155, or a later merge. The branch `handoff-context-wipe` holds this hand-off, as an open pull request. Run `make where` before you touch anything. Never commit on `main` (D-583). Answer the review of `gitar-bot` before you ask for a merge, and a pull request of documents alone waits for it too (D-637, D-679).

**The app is live on `decktome.com`.** A merge to `main` deploys itself on Cloud Build (D-584, D-586). #153 is the newest deploy, at 19:00 UTC on 2026-09-12: revision `mtg-api-00039-xqg` serves `api:b53fbb6`, both jobs run `worker:b53fbb6`, and `/readyz` answered 200. #154 and #155 changed documents alone. The deploy triggers read `go/**`, `docker/**`, and `web/**` alone, so a merge of documents alone starts no build.

**Deck gate run 19 and question gate run 49 read PASS** (D-686). Run 19 passes all 25 decks with one repair turn, and it is the decks baseline now. Its grade rows fell on most decks, because the local stored model `20260910T012734Z` graded them. Every curve line matches its list, so PR-43 holds on real builds. Run 49 met every expectation, and no conversation missed, so PR-42 ran no rerun.

**#153 merged PR-44, and Cloud Build deployed it: a card's types follow the rules of its layout** (F-120, D-687 to D-689). The index merged the types of every face, so Legion's Landing counted as a land. A card with more than one face takes its front face now, and split and modal double-faced cards keep every face. Quality gate run 21 reads PASS: 8 built decks grade bad against 9, and the judge agreement reads 9 against 10. The owner reads guardrail 15 as a rule for quality tuning items, so this correctness fix reports the numbers as information (D-689).

**The owner's Commander build reads right** (D-678). Session `OFMnk7Tv2zkK8xfAwXxB` built a bracket 5 treasure deck led by Smaug the Magnificent, from an owned-only pool. The grade reads typical, and its three reasons name the ladder. No rules check flags the deck: 33 lands, an average mana value of 2.29, and red sources at 1.72 of their need. The reader owns the commander, so OQ-79 stays open for its harder case.

**The read found F-119, and #149 fixed it** (D-684). The curve line of that deck reads 2.79 over 68 nonland cards, and the stored list reads 2.29 over 66. `generate.go` ran the rules engine before the mana pass of PR-33, so every rules finding read the list before the pass. #149 merged PR-43, which runs the pass first, and Cloud Build deployed it at 06:25 UTC. The pass adds only shortlist cards, so the fault is stale text alone.

**#148 merged PR-42, and Cloud Build deployed it** (F-112, D-671). A counted conversation with a miss plays two more times inside the question gate run. The document and the run file carry its miss rate. The first play still sets the verdict. No whole run shows a rerun yet.

**M-13 measured the date cut of PR-41, and the owner closed PR-41** (F-118, D-682, D-683). A free count found that the top lists hold Hobbit cards at the rate of older sets. No cut of 30 or 90 days moves a built deck to another tier, and each cut makes one number a little worse. `docs/reference/m13-date-cut-fits-2026-09-11.md` holds the count and the fits.

**#144 merged PR-40, and Cloud Build deployed it** (D-678). The build finished at 18:44 UTC on 2026-09-11. Revision `mtg-api-00036-bkn` serves `api:38ddde6`, and both jobs run `worker:38ddde6`. `/readyz` answered OK. Gate run 20 reproduces the pick of M-12: 9 built decks graded bad against 16, and the judge agreement at 10 against 9.

**PR-40 grades a Commander deck by three rules checks** (M-12, D-676 to D-678). A deck grades bad when one check fails:

- The lands fall 10 or more short of the need of Karsten's formula.
- The spells average over 4.4 mana.
- The sources of the worst color sit under 0.75 of their need.

Any other deck takes the tier of the fitted ladder. The score keeps the fitted model, so every bar reads as gate run 19. `go/internal/quality/rules.go` holds the checks, and `docs/reference/m12-rules-detector-fits-2026-09-11.md` holds the fits.

CAUTION: a grade stored on a deck built before the deploy keeps its old tier until a rebuild or a revision. The rules checks are code, and the stored model holds none of them.

**M-10 measured the softmax scorer again, and the owner closed PR-38** (D-680). The scorer drops the Commander precon bar to 748 of 788, one pair short. It grades 14 built decks bad against 9, and the owner's grades match on 2 decks against 5. **A quality item now passes only when neither reader number gets worse** (D-681, guardrail 15).

**The deployed model.** The meta job refits the model on every run (`go/cmd/worker/meta.go`). Its run of 2026-09-11 at 06:00 UTC finished at 06:41 UTC, before the deploy of PR-40, and stored `20260911T063437Z`. The run of 2026-09-12 at 06:00 UTC ran on `worker:6d9dacf`, the first run after the deploy of PR-40. It succeeded at 06:22 UTC and stored `20260912T061545Z`, fitted on 38,395 lists and 909 commanders. The Commander accuracy fell from 0.644 to 0.542, as gate run 20 predicted. The holdout accuracy reads the tier, and the tier reads the rules checks now.

CAUTION: the deployed Standard fit reads its cross share of baseline over bad at 0.690, against 0.750 on 2026-09-11 and 0.737 on 2026-09-10. Local gate run 20 reads 0.92. The share is information and no bar, and the rules checks of PR-40 read Commander alone. The deployed store holds 1,935 synthetic Standard copies, against 370 in the local store.

**M-11 is done, and no detector variant meets its gate** (F-115, D-675). Every disputed deck grades bad in all fifteen fits. The grade adds the detector probability to the bad rung of every deck, so the flag never decides the tier. The owner chose the rules detector first, and M-12 fitted it. `docs/reference/m11-detector-variants-2026-09-11.md` holds every fit, and `.local/m11/` holds the patches.

CAUTION: the `decktome` gcloud configuration named the Wallabee account and project on 2026-09-10, so `use_decktome` put the shell on the wrong project. `scripts/read-session.sh` reads `SESSION_PROJECT` and the account of `CLOUDSDK_CORE_ACCOUNT`, so an environment override reads `decktome-prod` with no change to the configuration. The owner has the commands to repair the configuration. A read on 2026-09-12 at 19:24 UTC found the configuration unchanged.

**The commander offer never shares a turn with a theme or colors question** (F-111, D-669). One turn asked both and offered three commanders, so the offer ignored the answers it ranks on. A wait for the answers failed question gate run 45, because a reader who skips a question got no offer. #139 merged it, and session `ze0Im17gZlFIBRyn7k7Z` on the deployed app asked the theme and the colors alone on turn 1.

**Gate runs 46 and 47 tripped two latent gaps, and #139 closes both** (F-112 to F-114, D-670). A delegation that names its slot declines no other key, and an occasion does not fill the theme. The gate keeps a bar of zero misses, and PR-42 reruns a missed conversation and reports its miss rate (D-671, #148). Gate run 48 reads PASS with both guards.

**The commander offer ignored the bracket of its own turn, and #134 fixed it** (F-104, D-662). The turn built the hints from the restored slots, and the picked power option landed after that. The pick row fires on the turn the power answer arrives (D-631), so the bracket 5 signal ordered almost no deployed offer. Session `vDzEKDPnRZntlyhz8Lqg` read Lotho, Ghalta, and Peregrin Took. The hints take the bracket after the answers now, and a test holds it. Session `ze0Im17gZlFIBRyn7k7Z` of 2026-09-11 confirms it: the offer named Vivi Ornitier, Aang, and Hapatra.

**PR-39 is merged as #135** (D-659 to D-661, F-102, F-103). The source count reads a reliable tap ability alone, and the balance phase trades basics toward the color that falls short. The measurement of D-661 is done. Over run 18 the pass moves four decks. The worst color rises on all four and falls on none, and deck 22 gains two Forests. `docs/reference/pr39-manapass-run18.md` holds the lane.

**Gate run 19 measured PR-39 on the quality model, for nothing** (D-663, F-105). The feature `color_sources` reads the count PR-39 changed, and no gate read the model before. Every bar holds, and the judge agreement rises from 8 to 9 of 25. A control fit on the parent commit matches run 18 in every number. **An item that changes a feature of the quality model reports the three numbers now** (guardrail 14, D-664).

**PR-37 is merged as #127, and gate run 18 reads PASS** (D-653). The synergy check drops a Commander copy whose break lowered the synergy feature by less than 0.10 standard deviations. The Commander synergy axis reads 0.84 of 232, and the precon bar 0.95 of 788. No earlier run of the quality gate in the repository reads PASS.

**CAUTION: the PASS of run 18 came from the population of the bar, and not from a better model.** Over run 18 the built decks graded bad stayed at 15 of 25, and the judge agreement at 8 of 25. Run 19 reads 16 and 9, from the source count of PR-39 (D-663). Run 20 reads 9 and 10, from the rules checks of PR-40 (D-678). 36 of the 232 synergy pairs still lose on run 19, and those are the true gap of the axis.

**CAUTION: the check reads Commander alone** (D-653, F-96). In Modern the synergy break does not lower the synergy feature in 825 of 845 pairs. A check there dropped 78 percent of the Modern synergy copies, and three decks the judge reads as bad graded higher.

**The premise of D-652 was half right.** The check drops 157 of 389 Commander pairs from every band of precon strength, and not only the 54 weak pairs M-8 named. `docs/reference/pr37-synergy-materiality-2026-09-10.md` holds all six floors.

**The gate reports the three numbers of D-648 for free.** `make quality-gate` grades the decks of deck gate run 16 with the model it fits, and it reads the judge's tiers from judge lane run 5. It stores no model.

**The owner graded the ten disputed decks** (D-659). The judge bar is the right target, and the detector reads built decks wrong. PR-40 merged the rules detector of M-12 (D-678), and the owner closed PR-38 after M-10 ran again (D-680). The review sheet is `.local/review/disputed-decks-2026-09-10.md`, beside a CSV, and it stays out of git. Its notes on decks 11 and 15 read the fit of 2026-09-07, and the ladder of gate run 19 reads both decks baseline. `docs/reference/m10-softmax-scorer-2026-09-10.md` holds the fits.

**Three items aimed at the reader gap failed first**: PR-29, M-9, and PR-30 (F-94, F-98, F-99). The owner parked PR-35 and dropped PR-30.

**CAUTION: two features never get built again** (F-94). `casual_unseen_share` moved the target axis by nothing and cost three other bars. `casual_pair_share` won one pair of 389, because the fit split it against `casual_synergy` with opposite signs. `docs/reference/pr29-casual-corpus-2026-09-10.md` says why for each, and PR-35 builds neither one.

**Every quality item reports three numbers** (D-648): the precon bar, the count of built decks graded bad, and the tier judge agreement. **The owner drops an item that moves none of the three.** PR-29 moved the first up and the other two down, and that is why it closed.

**The whole feedback loop stands**, as #117, #121, and #122. **No live cycle ran yet.** It needs the owner's word, `AUTOTUNE_FIXER_CMD`, and a harvest whose verdicts carry a snapshot.

**PR-34 is merged as #123.** The app reads the format of an upload out of the file, and Moxfield reads. OQ-80 asks whether a proxy counts as a card the reader owns.

**The whole PR-22 gate holds** (D-633). The measured cost at idle reads $9.54 a month gross, and the hourly snapshot tick of D-634 leaves $3.81.

**What waits on the owner.**

- A Commander build on the deployed app, and its session id (next step 1).
- OQ-67, the Stage B channels.
- OQ-77, the blocking function of Identity Platform.
- OQ-79, the commander of an owned-only pool, open for a harder case (D-673).
- OQ-80, a proxy and the pool.
- OQ-82, the power of a bracket (F-110).

**OQ-79 holds one owned-only session that reads right** (D-673). The harder case, a pool whose best commander the reader does not own, has no session yet.

**CAUTION: a local `make verify` is not the whole story.** It read green for weeks while shellcheck failed (F-89, D-641). Every pull request runs the workflow now.

## How to resume

1. Run `make where`. It prints the branch, the tree, and the state of the branch's pull request.
2. Run `ps aux | grep autotune` before any write. The loop resets the tree when it rejects an iteration.
3. Make a branch from `main`. Never commit on `main`, and never push to it (D-583). The owner merges (D-585). Run `make hooks` one time in a fresh checkout.
4. Load the skills. Load `ste-writing` before you write any `.md`. Load `design-doc-style` before you edit the roadmap. Load `mtg-corpus` before you reason about a format, a legality, or a card term.
5. Run `make verify`. It runs every check the verify workflow runs, on this machine, for nothing. Put Node 22.23.2 on the PATH first. The pull request runs the same jobs on Actions (D-639).
6. Do "Next steps, in order" below. Ask questions as they come up. Record each owner answer in `docs/decisions.md`, and delete the row from `docs/owner-questions.md`.
7. Open the pull request, then wait for the review of `gitar-bot` and answer it (D-637). A finding with merit takes a change, a commit, a push, and a reply. A finding with no merit takes a reply that says why, and you resolve it. Tell the owner when the pull request is ready to merge. Gitar is the only review this repo asks for. A pull request of documents alone waits for the review too (D-679).
8. Before you end, update this file. Move the oldest session to the archive when the count passes ten.

A second Mac: `docs/setup-second-mac.md` holds what to carry, what to install, and how to prove the machine.

Twelve things a fresh session gets wrong without this file.

- Twelve targets and two loop scripts spend money: `make questions-gate`, `make questions-eval`, `make eval-calibrate`, `make deck-gate`, `make bracket-gate`, `make revise-gate`, `make chat-probe`, `make generate-probe`, `make summary-judge`, `make quality-judge`, `make test-smoke`, `make feedback-triage`, `scripts/autotune.sh`, and `scripts/feedback-loop.sh`. Ask the owner before each run. `make autotune`, `make feedback-loop`, `make feedback-loop-dry`, `make feedback-triage-dry`, and `eval sweep -dry` are free.
- A rerun writes to a new file. Every `*_OUT` variable refuses a document that holds a result (D-65).
- A gate run takes about 20 minutes and an eval about 13. A foreground command stops at 10 minutes, so run both in the background.
- `docker compose up` needs provider keys in `.env` now, and fails fast without them (D-267). `make dev` still starts with no keys, and the fake builds the fixture deck of PR-23 for any first message (D-552).
- The eval and the agent share a model. Every ratio it reports is a floor, not a measurement (D-136).
- The judge is noisy: two runs of identical code move up to nine bad questions (D-230). One run proves nothing on its own.
- The sweep's estimate of a step is the cost of its last run file. The deck gate estimate read $2.54 from run 14, and run 16 cost $3.79 with the plan judge and 12 repair turns. A cap set from the estimate stops the sweep before its last step.
- Many local fits and verify runs fill the disk. On 2026-09-11 `make verify` failed at link time with 258 MiB free, when the Go and Docker build caches held about 50 GB. Run `df -h /System/Volumes/Data` before a long run, and ask the owner before you clear a cache.
- Gitar sometimes deletes its summary comment and posts a new one with a new id. Find its newest summary by author, and read the review threads. A green Gitar check does not prove that no finding is open.
- Gitar can post that it paused automatic reviews, because the trial's processing ran out for the period. On #150 to #155 the note came beside a full review each time, and no trigger ran. When the note comes with no review, comment `Gitar review` on the pull request. Answer the manual review as usual (D-685).
- The owner can merge a pull request before its Gitar finding gets an answer (#154). The pre-commit hook then refuses a commit on that branch. Carry the fix to a new branch from `main`, and reply on the old thread with the new pull request.
- A background shell command starts in the directory the session left. On 2026-09-12 a `make` target ran in `go/` and found no rule, so head every command with an absolute `cd`.

## Facts that expire

- Comprehensive Rules: the current file is 2026-08-19 (D-272). The rule citations in the corpus match the 2026-08-07 text. The session of 2026-09-12 read rules 709.4c, 710.2, 712.8a, 712.12, 715.2, and 722.2a in that file for PR-44.
- Ban-list snapshot: 2026-08-24. Next announcement 2026-10-12, in `internal/cards/announcement_dates.json`. A test fails when that file holds no future date. This is by design.
- Commander brackets: the 2025-10-21 revision. Game Changers: 53 cards, list of 2026-02-09. Lutri is banned as a companion only, per the 2026-02-09 announcement (`companion_bans.json` holds the link). The content rules per bracket in `brackets.json` and the Spellbook thresholds were read 2026-09-02, and the Karsten tables are the 2022 articles, read 2026-09-02 (`docs/reference/bracket-profile-2026-09-02.md`).
- Standard: 18 sets, Wilds of Eldraine to The Hobbit. Six sets leave at the first 2027 set. Verified 2026-08-24.
- Card snapshot on disk: `.local/gcs/mtg-local-cards/scryfall/20260904T210157`, the newest of 13. Read 2026-09-12. Check every card fact against it.
- LLM model ids and prices: `roles.json` and `prices.json`, verified 2026-08-24. The max output per provider in `llm/client.go`, verified 2026-08-29. The OpenAI rows are unverified by anyone but the owner. The judge role runs on Opus 5 (D-430).
- The deck gate upgrade probe cost $0.32 for 2 prompts, 7 calls, and 266 seconds. A partial run reads its own item bars and never stands as the gate (D-526).
- Run cost, read from the run files on 2026-09-09. The question gate cost $0.190 on run 42, $0.193 on run 43, and $0.194 on run 44, over 18 to 21 minutes. Runs 45 to 48 of 2026-09-11 cost $0.1925, $0.1920, $0.1380, and $0.1923, over 20, 22, 19, and 21 minutes. Run 49 of 2026-09-12 cost $0.1935, over 18 minutes. The deck gate upgrade probe cost $0.32 for 2 prompts. The eval cost $0.092 on run 34 and $0.096 on run 35, over 12 to 14 minutes. The deck gate cost $2.71 on run 18, over 37 minutes, and $2.75 on run 19 of 2026-09-12, over 36 minutes. The revise gate cost $1.07 on run 9. The bracket gate judge lane cost $0.28.
- The backfill of 2026-09-09 read one user with a record to seed: 1 collection, 1 thumbs up, and 2 thumbs down. It counted no deck and no chat, because the reader deleted both (D-635).
- The feedback store holds 3 verdicts on 2026-09-09, and every one predates the snapshot of D-635. `make feedback-list VERDICT=` reads both verdicts now (F-87). A collection group query over it needs an index for its shape, and the store holds "verdict ascending, created_at descending" alone.
- The deployed schedules, read 2026-09-09 and again on 2026-09-11: `mtg-snapshot-schedule` at `0 * * * *` (D-634) and `mtg-meta-schedule` at `0 6 * * *`. Both read ENABLED. The API service holds minScale 0, so it scales to zero. No billing export exists, so no command reads the billed spend.
- The deployed API, read 2026-09-12 at 19:00 UTC: revision `mtg-api-00039-xqg` on image `api:b53fbb6`, from #153. Both jobs run `worker:b53fbb6`, and `/readyz` answered 200. The build of #153 finished at 19:00 UTC, and #154 and #155 changed documents alone.
- The deployed quality model, read 2026-09-12: `20260912T061545Z`, from the meta job of 06:00 UTC. It fits 38,395 lists and 909 commanders. Its Commander fit reads `immaterial` 243 and accuracy 0.542, and its Standard fit reads a cross share of 0.690. The job ran 21 minutes, and the run of 2026-09-11 ran 41.
- The Karsten land article of 2022-07-29, read 2026-09-11 through `infinite-api.tcgplayer.com/content/article/<id>/`, because the page draws its text in the browser. `docs/reference/m12-rules-diagnostic-2026-09-11.md` holds the formula, the error, and the cheap rules.
- Baselines, in `docs/reference/eval/baselines.json`: questions is run 42, decks is run 19, revise is run 9. Run 49 is the newest whole questions run, and it reads PASS. Runs 45 to 47 read FAIL (D-669, F-112), and run 43 records the regression of F-84. `make eval-check` compares the newest whole run of a suite against its baseline. The quality gate has no baseline row, and run 21 is its newest run.
- Toolchain on this Mac, read 2026-09-11: Go 1.27.1, Node 22.23.2, pnpm 9.2.0, firebase-tools 14.14.0, and Java 17.0.20.1. Playwright is 1.63.0 with its Chromium headless shell (`playwright install chromium`). `go.mod` asks Go 1.27.0 or newer, and vitest is 5.0.0 since #133.

## Next steps, in order

1. **Read the curve line of the next deployed build** (F-119, D-684). It waits for a build of the owner. Cloud Build deployed PR-43, so the curve line must match the stored list. Read the deck with `KIND=decks` and `RAW=1` of `scripts/read-session.sh`, and decode `deck_gz`. Cloud Build deployed PR-44 too, so a card such as Legion's Landing must count as a spell.
2. **Read the meta job run of 2026-09-13 at 06:00 UTC**, the first on `worker:b53fbb6`. It is the first fit on the types of PR-44. Check that it succeeded and stored a model. Read the Standard cross share beside 0.690, and the Commander accuracy beside 0.542.
3. **The next collection platform, when the owner names one** (F-91). Five are left: Archidekt, Deckbox, Delver Lens, TCGplayer, and Helvault. Each one takes a real export and never a column list. `docs/reference/pr34-collection-formats-2026-09-10.md` holds the shape.
4. **The live half of the feedback loop has no run yet.** Three things want a measurement: the judge lane of the triage, one live fix cycle, and one review round. All three need the owner's word, and the cycle also needs `AUTOTUNE_FIXER_CMD` and a harvest whose verdicts carry a snapshot.
5. **Deck gate run 19 is the decks baseline** (D-686). Its grades read the local stored model `20260910T012734Z`, and the deployed app reads a newer one. The next whole run compares against run 19, and it costs about $2.75, so ask the owner first.
6. **OQ-79 waits for a harder case** (D-673). Session `ze0Im17gZlFIBRyn7k7Z` read right. The next evidence is an owned-only build whose best pool commander the reader does not own.
7. **The owner parked PR-35 and dropped PR-30** (D-655, D-656). M-9, F-97, and F-99 found no case for either one.
8. **PR-36, the reader's verdict as a quality signal** (D-651). It waits for verdicts.
9. **The weekly EDHREC read** falls on 2026-09-14 (D-499, D-565). Run `make meta-refresh`, then `make quality-gate` to a new `QUALITY_GATE_OUT`. The refresh also stores a new local model, and the deck gate grades read that model.
10. **PR-26, the return channels**, waits on OQ-67.
11. **PR-42 is merged as #148** (D-671). Question gate run 49 missed no conversation, so it ran no rerun. Any miss still fails the run, and each miss joins the finding register.

CAUTION: `make revise-gate | tee` hides the exit code. Read the verdict line of the document, never the exit code of a pipe.

CAUTION: `make verify` runs `eval-check`, and `eval-check` reads the newest whole run of each suite. A committed FAIL run turns verify red until a newer whole run passes.

The CI step "fake gcs tests" ran no test until 2026-09-10 (D-658). Its filter matched no test name, so it passed as a no-op. It runs `scripts/gcs-check.sh` now: the script seeds the fake GCS from the trimmed snapshot and fails unless `TestLiveFakeGCS` passes by name. `make gcs-check` runs the same script.

A ruleset that requires the `verify` check on `main` is not possible. The repo is private on the free plan, and the rulesets API answers 403 (checked 2026-08-28). The owner reads the checks before a merge.

## The ten most recent sessions

### 2026-09-12g: the documents read the state before a context wipe

**The owner merged #155 and asked for every document to read the current state** before a context wipe. The session changed no code, and it ran no paid target.

**The deploy did not change.** #154 and #155 changed documents alone, so revision `mtg-api-00039-xqg` still serves `api:b53fbb6`. The meta job of 2026-09-13 at 06:00 UTC is the next deployed event.

**The refresh.** The resume section, the next steps, and `CLAUDE.md` read the state after #155. Each record of 2026-09-11g onward names the pull request that carried it. OQ-79 and OQ-82 carry the evidence of the Smaug session.

### 2026-09-12f: PR-44 deployed

Merged as #154. #155 carried the stage date of `CLAUDE.md` that Gitar asked for.

**The owner merged #153, and Cloud Build deployed it at 19:00 UTC.** Revision `mtg-api-00039-xqg` serves `api:b53fbb6`, both jobs run `worker:b53fbb6`, and `/readyz` answered 200. The meta job of 2026-09-13 refits the deployed model on the types of PR-44.

### 2026-09-12e: the gate runs, F-120, and PR-44

Merged as #153.

**The owner approved deck gate run 19 and question gate run 49, and both read PASS.** Run 49 cost $0.1935 over 18 minutes and met every expectation. Run 19 cost $2.7504 over 36 minutes and passes all 25 decks with one repair turn. The owner made run 19 the decks baseline (D-686).

**The check of PR-43 on run 19 found F-120.** Every curve line matches its list. Deck 12 differs from a front-face count by one card, Legion's Landing, which the index typed as a land. The index merged the types of every face, against CR 712.8a.

**The owner chose the rules of each layout** (D-687, D-688). PR-44 gives a card with more than one face its front face, and split and modal double-faced cards keep every face. A test of eight real cards fails on the old types. One shortlist of 25 lost two cards. Quality gate run 21 moves deck 7 from bad to baseline, so graded bad reads 8 and agreement 9. The owner reads guardrail 15 as a rule for tuning items (D-689).

### 2026-09-12c: the Gitar pause note (D-685)

Merged as #151, and #152 corrected the reading of the rule.

**Gitar posted on #150 that it paused automatic reviews**, because the trial's processing ran out for the period. The owner chose a `Gitar review` comment when that note appears (D-685), and #151 recorded the rule.

**The note stopped no review yet.** #150 and #151 each got a full automatic review beside the note, so no trigger ran. #151 merged while `verify:go` still ran, and the job passed at 07:00 UTC. A correction reads the rule as a trigger when the note comes with no review.

### 2026-09-12b: PR-43 deployed, and the meta run of 2026-09-12

Merged as #150.

**The owner merged #149, and Cloud Build deployed it at 06:25 UTC.** Revision `mtg-api-00038-h57` serves `api:e8c1b4a`, both jobs run `worker:e8c1b4a`, and `/readyz` answered 200.

**The meta run of 06:00 UTC succeeded at 06:22 UTC**, on `worker:6d9dacf`. It stored `20260912T061545Z`, fitted on 38,395 lists and 909 commanders. The Commander accuracy fell from 0.644 to 0.542, as gate run 20 predicted. The deployed Standard fit reads its cross share at 0.690, against 0.750 the day before, and that share is no bar.

### 2026-09-12: the deployed grade, F-119, and PR-43

Merged as #149.

**The owner merged #148, and Cloud Build deployed it.** Revision `mtg-api-00037-5np` serves `api:6d9dacf`, both jobs run `worker:6d9dacf`, and `/readyz` answered 200.

**The owner built a Commander deck and named its session.** Session `OFMnk7Tv2zkK8xfAwXxB` asked for a treasure deck, named Smaug the Magnificent, and chose bracket 5 from an owned-only pool. The grade reads typical, and its reasons name the ladder, as D-678 asks. No rules check flags the deck. The reader owns the commander, so OQ-79 keeps its harder case open.

**The read found F-119.** The curve line reads 2.79 over 68 nonland cards, and the stored list reads 2.29 over 66. The engine ran before the mana pass, so every rules finding read the list before the pass. The owner chose PR-43, which runs the pass first (D-684). A whole build through a fake model fails on the old order and passes on the new one.

### 2026-09-11i: PR-42, the gate reruns a missed conversation

Merged as #148.

**The owner merged #147.** No Cloud Build followed, because #147 changed documents alone. Next steps 1 and 2 still waited on the owner and on 06:00 UTC. Steps 3 to 10 waited on the owner, on data, or on a date. So the session built PR-42, the first free item.

**The build** (D-671). A counted conversation with an expectation or must-not-ask miss plays two more times inside the run. Each rerun keeps its own misses, and a failed rerun reads its error as a miss. The first play still sets every count and the verdict. The document writes the rate beside the miss, and the run file carries `miss_rate` as an information row, so `eval-check` never flips on it.

**The proof is free.** A test plays a missed conversation through a fake player. No paid run went out, so no whole run shows a rerun yet.

### 2026-09-11h: M-13, the date cut of PR-41, and PR-41 closed

Merged as #147.

**The owner merged #146 and asked what comes next.** Next step 1 waited on a build of the owner, and step 2 on the meta job of 06:00 UTC. So the session took PR-41, and its first design question.

**A free count refutes the cause of F-108** (F-118). The fit keeps the newest 4,000 great and 4,000 good Commander lists. Those lists hold 54 percent of the nonland cards first printed in The Hobbit. Marvel Super Heroes reads 50 percent. The colors check already grades decks 19 and 22 bad, so deck 21 alone had room to gain. The owner chose to measure the planned cut first (D-682).

**M-13 fitted the cut for free, and no cut meets the gate.** The control reproduces gate run 20 in every number. The cuts of 30 and 90 days in Commander, and of 30 days in every format, move no built deck to another tier. Each cut makes one number a little worse. `.local/m13/` holds the patch, the count scripts, and the runs.

**The owner closed PR-41 on the evidence** (D-683). The fits do not measure the first days after a release, and the owner chose a close over a park.

### 2026-09-11g: the documents read the state before a context reset

Merged as #146.

**The owner merged #145 and asked for every document to read the current state** before a context reset. The session changed no code, and it ran no paid target.

**The deploy did not change.** #145 changed documents alone. At 23:11 UTC revision `mtg-api-00036-bkn` served `api:38ddde6`, and both jobs ran `worker:38ddde6`. Both schedules read ENABLED. The meta job ran last at 06:00 UTC, before the deploy of PR-40.

**The refresh.** The resume section reads the newest work first, and the next steps name the meta run of 2026-09-12 and PR-41. Four roadmap lines read "built on branch" for work merged long ago, and they read #54, #73, and #79 now. F-109 and the gate of PR-41 read the result of M-10 and D-681. The toolchain line reads Go 1.27.1, and `.local/m10/` holds a README now.

### 2026-09-11f: PR-40 deployed, M-10 again, and PR-38 closed

Merged as #145.

**The owner merged #144, and Cloud Build deployed it** at 18:44 UTC. Revision `mtg-api-00036-bkn` serves `api:38ddde6`, both jobs run the new worker image, and `/readyz` answered OK. The read of a deployed Commander grade waits for a build of the owner.

**The session answered the review of #144 first.** Gitar found that the explain ladder showed a flagged deck's tier and not the ladder. The fix showed the ladder's own probabilities, and Gitar approved with one finding resolved.

**M-10 ran again for free, as D-665 asks** (D-680). No file kept the patch of 2026-09-10, so the session rebuilt the softmax ladder. The control reproduces gate run 20. The scorer drops the Commander precon bar to 748 of 788, one pair short, and it grades 14 built decks bad against 9. The judge agreement rises from 10 to 12, and the owner's grades match on 2 decks against 5.

**Two owner answers.** PR-38 closes on its evidence (D-680). A quality item passes only when neither reader number gets worse (D-681).

## The archive

`docs/reference/session-handoff-archive.md` holds every record this file no longer carries. It holds the resume section of 2026-09-08, the records of 2026-08-31 to 2026-09-11e, and 42 more sections, word for word. Read it for the detail behind a decision.
