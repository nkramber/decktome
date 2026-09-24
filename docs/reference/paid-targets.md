# Paid and free targets

This file holds the cost, the flags, and the guards of each `make` target and loop script. `CLAUDE.md` names each paid target and the rule to ask first. `make context-budget` fails when the lists of paid targets in `CLAUDE.md`, `docs/SESSION-HANDOFF.md`, and this file differ (D-749).

## The paid targets

Fifteen targets and two loop scripts spend money: `make codex-review`, `make questions-gate`, `make questions-eval`, `make eval-calibrate`, `make deck-gate`, `make bracket-gate`, `make sixty-gate`, `make revise-gate`, `make chat-probe`, `make generate-probe`, `make summary-judge`, `make quality-judge`, `make test-smoke`, `make feedback-triage`, `make api-build`, `scripts/autotune.sh`, and `scripts/feedback-loop.sh`. Ask the owner before every run, except a round of `make codex-review` (D-831).

## Each target

`make codex-review PR=<n>` starts the Codex review of one pull request (D-823). It spends the usage of the owner's Codex plan, and never an API key (D-833). So one request of the owner approves each round of the review loop of a pull request (D-831). The rule of three open rounds limits the loop (D-826). No run has a measured cost or duration yet.

- The guards: an open pull request, a checkout at its head, a clean tree, and a complete Gitar pass. A failed guard spends nothing, and the target exits 5.
- The flag: `make codex-review PR=<n> -- --skip-gitar-review` reads no Gitar pass, and it refuses an open review thread or an issue on the Gitar dashboard alone (D-838). Each Codex review of the Gitar pause uses it.
- The CLI: the target runs `npm install -g @openai/codex@latest`, and it refuses a version below 0.156.1 (D-824).
- The login: `codex login status` must give "Logged in using ChatGPT". Each call runs with `OPENAI_API_KEY` and `CODEX_API_KEY` removed.
- The model: `gpt-6-luna` at the effort `medium`, in the sandbox `danger-full-access`, from the command line alone (D-823, D-825).
- The output: the transcript goes to `.local/codex-review/`, and the last line names the outcome and the exit code (D-832).

`make ruleset-check` is free. It reads the ruleset of `main` and the merge settings through `gh api`, and it compares them with `.github/rulesets/` (D-828).

`make questions-gate` calls the real providers. One run of the 109 conversations (78 gate and 31 probe since D-730) costs $0.18 to $0.19 and takes about 20 minutes, measured on runs 33 to 35 (2026-09-04). Ask the owner before every run, and write to a new `GATE_OUT` file: a rerun must never overwrite a scored document (D-65).

`make questions-eval` scores a gate run with the eval role. One run costs $0.09 to $0.10 (runs 33 to 35) and takes about 13 minutes. `make eval-calibrate` measures the eval model against `claude-opus-5` (D-428). It cost $0.25 to $0.30 on Sonnet 5, and Opus 5 costs about 1.7 times that.

`make autotune` is free. It prints the loop instructions and starts nothing. `scripts/autotune.sh` is the paid loop, and it refuses to start without `AUTOTUNE_ALLOW_UNATTENDED=1`. One iteration costs about $0.25 and takes 33 to 35 minutes, so a $3 budget buys about 12 iterations. Read `docs/reference/autotune-readme.md` and `docs/reference/autotune-design.md` first.

Five more targets spend money, and each has an overwrite guard and an env guard. `make deck-gate` builds the PR-8 gate document. Run 11 cost $1.46 for 24 prompts, and runs 12 to 14 cost $2.24 to $2.64 under the profile's repair passes. `make chat-probe` drives the real `Chat` RPC to a deck. `make generate-probe` builds one deck with the real generate role.

`make summary-judge` judges every deck summary of a gate document (F-26). Each probe costs a few cents. Ask the owner before every run.

`make bracket-gate` builds three commanders at each bracket, 15 decks, and asks the judge role for the bracket of each (PR-14A). It has the same two guards, `BRACKET_GATE=1` and a verdict check on `BRACKET_GATE_OUT`. Run 1 cost $2.08 for the builds and $0.26 for the judge lane, and the second judge lane of 2026-09-05 cost $0.28 (D-540). `BRACKET_GATE_ARGS="-only 7,8,9"` runs the bracket 3 prompts alone, `-rejudge <document>` judges the decks of a document for about $0.26, and `-dry` is free. The rejudge profiles each deck through Commander Spellbook first, for free, so the judge reads its combos (D-790). On 2026-09-21 the rejudge of run 2 cost $0.2446, and the calibration decks cost $0.3113 and $0.3107 (D-795).

`make sixty-gate` asks the step judge for the power step of each labeled 60-card list three times (PR-71, D-863 to D-871). It has the guard `SIXTY_GATE=1`, and any file at `SIXTY_GATE_OUT` stops the run. A FAIL verdict exits 2.

`SIXTY_GATE_ARGS="-split dev -reads 1"` reads the 15 dev lists once, and that document holds no verdict. Run 1 cost $4.1086 and run 2 cost $3.9736, for 180 reads each. `-exclude <documents>` leaves out the tournament lists of earlier runs (D-876). `make sixty-gate-dry` prints the split for free.

`make revise-gate` builds three base decks and runs nine revisions over them, twelve turns with the answered questions (PR-12B, D-448). It has the same two guards. Run 7 cost $0.74 for eleven turns, run 8 cost $1.23, and runs 4 to 6 cost $0.54 to $0.81.

`DECK_GATE_ARGS` passes flags to `make deck-gate`. `DECK_GATE_ARGS="-only 19,20,21,22,23,24"` runs the six set prompts of PR-17B alone, for about $0.35. `DECK_GATE_ARGS="-only 25"` runs the precon exclusion prompt of PR-24 alone, for about $0.13. `GATE_ARGS` passes flags to `make questions-gate`, and `GATE_ARGS="-only 109"` runs the group set probe of D-525 alone. Runs 36 and 37 cost about $0.001 each and took about 10 seconds. A partial run reads its item bars alone and never stands as the gate (D-526).

`DECK_GATE_ARGS="-rejudge <absolute path of a gate document>"` judges the stored decks of a whole run again, with no build (D-789). It copies the build rows of the source run and writes new judge rows. So its run stands as a whole run of the decks suite. Run 32 judged the 25 decks of run 31 for $1.2453, and a full run costs about $2.83.

Add `-keep <absolute path of an earlier rejudge run file>` to judge again only the decks that the earlier rejudge lost. Run 33 judged 6 decks again for $0.3482. Add `-summary-only` beside `-keep` to judge every summary again and keep the plan rows. Run 34 did that for $0.4653. Prove a change to a judge alone with this lane, and never with a rebuild.

`make quality-judge` asks the judge role for the tier of every graded deck of a deck gate document (PR-14B). It costs a few cents a deck, and it has the guard `QUALITY_JUDGE=1` and a verdict check on `QUALITY_JUDGE_OUT`. Ask the owner before every run.

`make api-build` builds one deck over the deployed API, with no browser and no GUI (D-778). It signs in with an email and a password, and it imports a ManaBox CSV. It then answers every question of the agent, and it reads the built deck back out of storage. It costs about $0.10 to $0.20, because the deployed API calls the real providers.

The guards are `API_BUILD=1` and a check on `API_BUILD_OUT`. `API_BUILD_EMAIL` and `API_BUILD_PASSWORD` name the check account of D-779, and `.env` holds both. The run keeps the deck, the chat, and the collection. `API_BUILD_ARGS=-cleanup` deletes all three (D-780). `API_BUILD_ANSWERS="power=#3"` asks for a bracket 3 deck, and an empty plan builds a bracket 1 deck. `docs/reference/api-deck-build-2026-09-20.md` holds the method, the result, and the limits.

`make test-smoke` runs the live LLM smoke test and reads the keys from `.env`. It spends a few cents. The list at the top of this file names every paid target. `go run ./cmd/eval sweep -cap <USD>` drives five of them in the order of the eval list under a cap, and it needs `EVAL_SWEEP=1` (PR-15). `-dry` prints the plan for nothing, and the estimate of a step reads its last run file. Since PR-15 the deck gate spends one more judge call a deck, and run 16 cost $3.79 for 25 prompts.

Each other target is free. `make meta-refresh` reads the deck list sources over the network, about 40 minutes on the first run, and calls no model. The run of 2026-09-23 took 97 minutes, and it writes a new local quality model. `make quality-gate` fits the quality model over the local meta store and writes the PR-14B gate document. It also reports the three numbers of D-648: the fitted model grades the decks `QUALITY_GATE_DECKS` names, against the judge's tiers in `QUALITY_GATE_JUDGED`.

`make ste-check` checks every hand-written `.md` file against the STE rules, and `make lint` runs it. `make ref-check` checks every cited id and every repository path of the same files, free, and `make lint` runs it (D-753). `make m5-sheet` builds the scoring sheet, and `make m5-report` reads it. `make themes-check` checks the theme slugs and the commander ranking.

`make eval-check` compares every baseline of the eval harness with its newest run and names the flips (PR-15). It is free.

`make self-reload-check` proves that an installed app reloads itself on the next web release (D-692). It builds two releases and drives Chromium, and it calls no provider. It needs Node 22.23.2 and a pnpm store, and it takes about three minutes. `SELF_RELOAD_REF=<ref>` names a different old release, and `docs/reference/self-reload-2026-09-20.md` holds the method and the first result.

`make verify` runs every check the verify workflow runs, on this machine, for nothing (D-578). Run it before every pull request. `make pr-check` reads the pull request body and diff against the contract of D-747, and `make lifecycle-check` tests the skill wiring and the session hook. Both are free. `make where` prints the branch, the tree, and the state of the branch's pull request. `make hooks` installs the pre-commit hook that refuses a commit on `main` (D-585).

`make feedback-harvest` writes every verdict since the last harvest to `docs/reference/feedback/`, as a dated document and a JSONL file (PR-28a). `SINCE=2026-09-01` sets the floor by hand, and `HARVEST_ARGS=-dry` counts and writes nothing. The watermark comes from the JSONL files, so the documents are the only record. It calls no model and costs nothing. Those files commit with the repository now (D-642).

`make feedback-triage-dry` routes every verdict of the newest harvest into a class (PR-28b). It calls no model and costs nothing. `make feedback-triage TRIAGE_OUT=<document>` asks the judge for the verdicts the reason keys can not place, at a few cents each.

The target reads no `.env`, so run `set -a && . ./.env && set +a` first. Set `CARDS_SNAPSHOT_DIR` to `.local/gcs/mtg-local-cards/scryfall`, or each card case names its gap. The run of 2026-09-23 cost $0.0122 for one judge call. `TRIAGE_ARGS=-apply` writes each case into the gate file that owns it. Ask the owner before every live run.

`make feedback-loop` prints the commands of the fix cycle and starts nothing (PR-28c). `make feedback-loop-dry` plans a cycle for nothing. `scripts/feedback-loop.sh` is the paid cycle, and it refuses to start without `FEEDBACK_LOOP_ALLOW=1` and `AUTOTUNE_FIXER_CMD`. One cycle stops at $2 of gate runs (D-559). It edits code, commits, pushes, opens a pull request, and answers the review, with nobody watching. Ask the owner before every run.

`make users-backfill` seeds the user record of D-638 from what each user already holds, and `BACKFILL_ARGS=-dry` counts and writes nothing. It counts a revision and an imported deck apart from a first build (D-861). It never lowers a count. It calls no model and costs nothing.

`make read-session SESSION=<id>` reads one chat session of the deployed project, to debug it (D-596). `make feedback-list` reads the newest verdicts of every user over one collection group query, and `VERDICT=up` and `LIMIT=` change what it reads. All three print or write what a reader wrote, and none of them writes an email. `use_decktome` puts the shell on `decktome-prod`. No decktome tool reads `PROJECT_ID`: a shell that works on more than one project exports it for another one.

`make smoke` runs the Playwright smoke flow of PR-23 over the emulators, the trimmed snapshot, and the fake provider (D-553). It calls no model, and it needs the Chromium build of Playwright once.

`make deck-gate-dry` builds every deck gate shortlist over the trimmed snapshot of the repo and calls no provider (D-521). `make deck-gate-trim` rewrites that snapshot from the local store, and the fixture must stay under 10 MB. Both are free.

`make allow EMAIL=... PROJECT_ID=...` invites one email to the deployed app, and `make disallow` takes one off (D-420). Both write one Firestore document of the deployed project with the caller's own credentials, and neither calls a model.

`make store-check` runs the session store against the local Firestore emulator. `make gcs-check` runs the live fake-GCS store test against a server seeded from the trimmed snapshot, and the CI step runs the same script (D-658). `make candidates-review` writes the PR-6 gate document from a local snapshot. `cd go && go run ./cmd/tune-check` compares an eval summary with its baseline.

`docs/reference/pr7-m5-scoring.md` is the owner's working copy. No target writes to it. A new sheet needs a new name and points at the latest gate document, for example `M5_OUT=docs/reference/pr7-m5-scoring-run18.md M5_RUNS=../docs/reference/pr7-question-gate-run18.md make m5-sheet`.

## The exit code of a target

Each recipe line of `Makefile` that pipes sets `set -o pipefail` itself (D-782). GNU Make 3.82 added `.SHELLFLAGS`, and GNU Make 3.81 ignores it, so a recipe reads pipefail from no variable. Four paid targets end with `| tee`, which shows the live output and writes the document. Without pipefail each one reported the exit code of `tee`, and `tee` succeeds after a command that fails. `make pipefail-check` holds the rule, and `make lint` and the `verify:shell` job both run it for nothing. `docs/reference/f160-make-pipefail-2026-09-20.md` holds each exit code and each proof.

CAUTION: `tee` creates its output file before the command writes a byte. So a failed run of `make chat-probe`, `make generate-probe`, `make summary-judge`, or `make api-build` leaves an empty document. The guard of D-65 then refuses that name on the next run. Delete the empty document, or name a new one.

## Measured run costs

These facts expire. Each one names the date of its read. They moved from `docs/SESSION-HANDOFF.md` word for word (D-749).

- The deck gate upgrade probe cost $0.32 for 2 prompts, 7 calls, and 266 seconds. A partial run reads its own item bars and never stands as the gate (D-526).
- Run cost, read from the run files on 2026-09-09. The question gate cost $0.190 on run 42, $0.193 on run 43, and $0.194 on run 44, over 18 to 21 minutes. Runs 45 to 48 of 2026-09-11 cost $0.1925, $0.1920, $0.1380, and $0.1923, over 20, 22, 19, and 21 minutes. Run 49 of 2026-09-12 cost $0.1935, over 18 minutes. Run 50 of 2026-09-14 cost $0.1938, over 22 minutes. Run 51 of 2026-09-15 cost $0.1958, over 21 minutes. Run 52 of 2026-09-15 cost $0.1951, over 21 minutes. The deck gate upgrade probe cost $0.32 for 2 prompts. The eval cost $0.092 on run 34 and $0.096 on run 35, over 12 to 14 minutes. The deck gate cost $2.71 on run 18, over 37 minutes, and $2.75 on run 19 of 2026-09-12, over 36 minutes. The revise gate cost $1.07 on run 9. The bracket gate judge lane cost $0.28. Bracket gate run 2 of 2026-09-13 cost $1.3720 over 925 seconds, for 15 builds and the judge. Its second judge lane cost $0.2634 over 128 seconds. The M-15 calibration lane cost $0.3283 over 149 seconds, for 21 decks. Bracket gate run 3 of 2026-09-14 cost $0.7355 over 538 seconds, for 9 builds and the judge. Bracket gate runs 4, 5, and 6 over prompts 10 to 15 cost $0.4979, $0.5247, and $0.5619, over 424, 442, and 564 seconds. The second judge lane of run 6 cost $0.0856 over 35 seconds. Deck gate runs 20, 21, and 22 over prompt 3 cost $0.1321, $0.1174, and $0.1159, over 94, 97, and 106 seconds. Deck gate run 23 of 2026-09-16 cost $2.6509 over 2,156 seconds, for 25 prompts. Bracket gate run 7 of 2026-09-16 cost $0.5235, for 6 builds over prompts 10 to 15 and the judge.
