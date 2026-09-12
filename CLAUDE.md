# decktome - CLAUDE.md

Read this file first. Then read `docs/SESSION-HANDOFF.md`. It tells you where the last session stopped.

## Project

This repo is a Go + Protobuf + TypeScript monorepo for an agentic MtG deck builder. The app reads a user's ManaBox collection export. The user gives a prompt. The agent asks questions, then builds a legal, useful deck.

Stage (2026-09-11): **the app is live on `decktome.com`**, and pull request #149 is the newest merge. A merge to `main` deploys itself on Cloud Build (D-584, D-586, D-588). Phase 3B, the product UI, is the current phase.

**The whole PR-22 gate holds** (D-633), so PR-28 is free to run. Every PR through #119 is merged: PR-0a to PR-8, PR-7B, PR-10 to PR-25, PR-27, PR-32, PR-33, and PR-28a. PR-9 is out of the MVP (D-256). **PR-28 split into three** (D-636). PR-28a is the harvest, and it is merged. PR-28b is the triage and the new "must not ask" expectation, and PR-28c is the fix cycle. PR-26 waits on OQ-67. Of the weak-axes plan, PR-29 and PR-30 closed on their evidence, and PR-31 parks (D-652, D-656, D-573).

**The whole feedback loop is on `main`**, as pull requests #117, #121, and #122. The triage names one of 22 classes for each thumbs down, and the reason keys answer 20 of them for nothing (D-643). A case joins the gate file that owns it, and the pull request diff is the accept step (D-642). The fix cycle carries the case and its fix in one pull request, and it answers the review of `gitar-bot` (D-645). **No live cycle ran yet.** **PR-34 is merged as #123**: the app reads the format of an upload out of the file, and Moxfield reads (D-647, F-91 to F-93).

**PR-29 closed, and it never merged** (D-652). M-8 read the precon bar and found its target of 0.95 was never reachable. 54 of 389 pairs carry no signal, because a weak precon's cards do not pair in the corpus (F-95). **Every quality item reports three numbers now**, and the owner drops an item that moves none (D-648). **PR-37 is merged as #127.** It adds the materiality check the synergy break never had, in Commander alone at a floor of 0.10 (D-653). Gate run 18 reads PASS, and no reader-facing number moved: the bar reads a new population, and the model is the same model.

**The owner graded the ten disputed decks** (D-659). The judge bar is the right target, and the defect detector reads built decks wrong. PR-39 merged first, as #135: the basic lands follow the need of each color (D-660, D-661). Gate run 19 measured the new source count on the quality model: every bar holds, and the judge agreement rose from 8 to 9 (D-663).

**M-11 measured fifteen fitted detectors, and none meets its gate** (F-115). The tier reads the detector probability of every deck, so no variant moved a disputed deck. **#144 merged PR-40, and the deployed app grades Commander decks by the rules checks of M-12** (D-675 to D-678). M-10 ran again and M-13 fitted the date cut, and the owner closed PR-38 and PR-41 (D-680, F-118, D-683). Every item that changes a feature of the model reports the three numbers (D-664). A tier change also reports the broken copies graded bad (D-674), and an item passes only when neither reader number worsens (D-681).

**The commander offer never shares a turn with a theme or colors question** (F-111, D-669). #139 merged it with two question-flow guards (F-113, F-114, D-670), and it deployed on 2026-09-11. The owner's owned-only bracket 5 session reads right on every check, and OQ-79 stays open for a harder case (D-673).

Two rules of the deployed app come from 2026-09-09. **A verdict keeps the object it names** (D-635), because a reader deletes the deck or the chat they complained about, and the verdict outlives it. **Every user has a record** at `users/<uid>` (D-638): the verified email, the dates, and six counters of what they made. No harvest reads that record, and a test refuses the import that joins them.

The look follows a reference design the owner gave on 2026-08-30 (D-328 to D-335). `docs/reference/autotune-readme.md` holds the loop commands.

**The repository is public** (D-639). Write no email and no personal address into a file, an issue, or a pull request. D-642 permits a reader's own words. Every pull request runs the whole verify workflow.

Run `make where` before you change anything. It prints the branch, the tree, and the state of the branch's pull request. The session commits on a branch, pushes it, and opens a pull request. The owner merges (hard rule 8, D-583, D-585). Never commit on `main`, and run `make hooks` once in a fresh checkout.

Read `docs/SESSION-HANDOFF.md` next. It is the resume point.

## Hard rules from the owner

1. **Write scope.** This repo permits writes (D-32, 2026-08-23). All other repos are read-only. Application code follows the roadmap order. Do not start a roadmap item before the gate of the item before it holds.
2. **Write in ASD-STE100.** Every doc, skill, and agent file must follow Simplified Technical English. Load the `ste-writing` skill before you write. Rules that apply most: max 20 words per procedural sentence, max 25 per descriptive sentence, and active voice. Also: one instruction per sentence, no semicolons, no "-ing" verb forms, one term per concept, paragraphs of max six sentences.
3. **Ask questions when you think of them.** Do not save questions for the end. Use `AskUserQuestion` in small batches. Record each answer in `docs/decisions.md`.
4. **Do the research.** Verify facts against sources (Scryfall API, Wizards announcements, the Comprehensive Rules). Record the date of each fact. MtG rules and ban lists change often.
5. **Make hand-off simple.** Before you end a session, update `docs/SESSION-HANDOFF.md`: the completed work, the open work, and the next step.
6. **No AI-attribution text** in any PR, branch name, commit message, or comment. This house rule comes from connector-syncer.
7. **No mistakes.** Check card names, rules, and dates before you write them. When you are not sure, say so and mark the item as unverified.
8. **Every change starts on a branch.** Never commit to `main`, and never push to it (D-583). Make a branch, commit there, push it, and open a pull request. The owner merges. Run `make where` before every commit, push, and deploy. It prints the branch, the tree, and whether `main` is current. It also names the state of the branch's pull request. Run `make hooks` one time, and the pre-commit hook then refuses what these rules forbid (D-585).
9. **Deploy from `main` alone.** Never deploy any other branch to production, for any reason (D-579). Check the branch and the commit before every build, not only the tree. `docs/deploy-and-rollback.md` holds the procedure.
10. **Answer the review before you ask for a merge.** `gitar-bot` reviews every pull request (D-637). Wait for that review. Read each finding on its merit, and never on its tone.
   - A finding with merit takes a change. Make it, commit, push, and reply to the comment with what you changed.
   - A finding with no merit takes a reply that says why, and you resolve it.
   - When no finding has merit, tell the owner the pull request is ready to merge. **Gitar is the only review this repo asks for.** No second harness reads it.
   - Repeat the cycle until the review holds nothing open. The owner merges.
   - A pull request of documents alone waits for the review too (D-679).
   - When Gitar posts that it paused automatic reviews, comment `Gitar review` on the pull request. Wait for the manual review. Answer it the same way (D-685).
11. **Never hesitate to ask or to push back.** Ask a question the moment you have one. When the owner's two statements conflict, say so and quote both. When a request rests on a wrong premise, say so with the evidence. The owner sees this as the key to good LLM-user interaction. Silence is the mistake, not the question.

## Reference material

- Style model for the design doc: `/Users/nate/Repos/connector-syncer-docs/docs/document-summary-roadmap.md`. Load the `design-doc-style` skill for the section template.
- AI patterns to borrow: `/Users/nate/Repos/connector-syncer` (Python RAG app). See `docs/reference/connector-syncer-ai-patterns.md`.
- Frontend patterns to borrow: `/Users/nate/Repos/wallabee-ui`. See `docs/reference/wallabee-toolchain-conventions.md`.
- MtG terminology and rules corpus: load the `mtg-corpus` skill. It is the app's future knowledge base.

## Skills in this repo

| Skill | Use when |
|---|---|
| `ste-writing` | Before you write or edit any `.md` file here. |
| `design-doc-style` | Before you edit `docs/design-roadmap.md`. |
| `mtg-corpus` | Before you reason about formats, legality, archetypes, or card terms. |

## File map

- `docs/design-roadmap.md` - the design document (the deliverable).
- `docs/decisions.md` - every owner decision, with date.
- `docs/SESSION-HANDOFF.md` - resume point for a fresh session.
- `docs/setup-second-mac.md` - what to carry and what to install to continue the work on another Mac.
- `docs/setup-gcp.md` - the deploy on Google Cloud from the domain to the running app, with the cost estimate.
- `docs/deploy-and-rollback.md` - the redeploy after a merge to `main`, and the rollback of each part.
- `docs/reference/mobile-and-engagement-2026-09-05.md` - the phone and engagement proposal in three stages.
- `docs/open-questions.md` - questions not yet asked or not yet answered.
- `docs/owner-questions.md` - the decision queue. Every question here waits for the owner, and the tuning loop refuses to decide one.
- `docs/reference/` - research notes with sources and dates, and every dated gate document.
- `docs/audit-2026-08-28.md` - the full audit of 2026-08-28, its owner answers, and the change plan.
- `docs/audit-2026-08-29.md` - the quality audit of 2026-08-29 and its fixes (D-302 to D-306).
- `docs/reference/set-data-2026-08-31.md` - every set number PR-17B rests on, with its source and date.
- `docs/reference/bracket-profile-2026-09-02.md` - every bracket rule, Karsten table, and Spellbook threshold PR-14A rests on.
- `docs/reference/deck-quality-model-2026-09-02.md` - every source fact, feature, and fit rule PR-14B rests on.

## Commands that cost money

`make questions-gate` calls the real providers. One run of the 108 conversations (77 gate and 31 probe since D-522) costs $0.18 to $0.19 and takes about 20 minutes, measured on runs 33 to 35 (2026-09-04). Ask the owner before every run, and write to a new `GATE_OUT` file: a rerun must never overwrite a scored document (D-65).

`make questions-eval` scores a gate run with the eval role. One run costs $0.09 to $0.10 (runs 33 to 35) and takes about 13 minutes. `make eval-calibrate` measures the eval model against `claude-opus-5` (D-428). It cost $0.25 to $0.30 on Sonnet 5, and Opus 5 costs about 1.7 times that.

`make autotune` is free. It prints the loop instructions and starts nothing. `scripts/autotune.sh` is the paid loop, and it refuses to start without `AUTOTUNE_ALLOW_UNATTENDED=1`. One iteration costs about $0.25 and takes 33 to 35 minutes, so a $3 budget buys about 12 iterations. Read `docs/reference/autotune-readme.md` and `docs/reference/autotune-design.md` first.

Five more targets spend money, and each has an overwrite guard and an env guard. `make deck-gate` builds the PR-8 gate document. Run 11 cost $1.46 for 24 prompts, and runs 12 to 14 cost $2.24 to $2.64 under the profile's repair passes. `make chat-probe` drives the real `Chat` RPC to a deck. `make generate-probe` builds one deck with the real generate role.

`make summary-judge` judges every deck summary of a gate document (F-26). Each probe costs a few cents. Ask the owner before every run.

`make bracket-gate` builds three commanders at each bracket, 15 decks, and asks the judge role for the bracket of each (PR-14A). It has the same two guards, `BRACKET_GATE=1` and a verdict check on `BRACKET_GATE_OUT`. Run 1 cost $2.08 for the builds and $0.26 for the judge lane, and the second judge lane of 2026-09-05 cost $0.28 (D-540). `BRACKET_GATE_ARGS="-only 7,8,9"` runs the bracket 3 prompts alone, `-rejudge <document>` judges the decks of a document for about $0.26, and `-dry` is free.

`make revise-gate` builds three base decks and runs nine revisions over them, twelve turns with the answered questions (PR-12B, D-448). It has the same two guards. Run 7 cost $0.74 for eleven turns, run 8 cost $1.23, and runs 4 to 6 cost $0.54 to $0.81.

`DECK_GATE_ARGS` passes flags to `make deck-gate`. `DECK_GATE_ARGS="-only 19,20,21,22,23,24"` runs the six set prompts of PR-17B alone, for about $0.35. `DECK_GATE_ARGS="-only 25"` runs the precon exclusion prompt of PR-24 alone, for about $0.13. `GATE_ARGS` passes flags to `make questions-gate`, and `GATE_ARGS="-only 109"` runs the group set probe of D-525 alone. Runs 36 and 37 cost about $0.001 each and took about 10 seconds. A partial run reads its item bars alone and never stands as the gate (D-526).

`make quality-judge` asks the judge role for the tier of every graded deck of a deck gate document (PR-14B). It costs a few cents a deck, and it has the guard `QUALITY_JUDGE=1` and a verdict check on `QUALITY_JUDGE_OUT`. Ask the owner before every run.

`make test-smoke` runs the live LLM smoke test and reads the keys from `.env`. It spends a few cents. The paid targets are these eleven plus the script: questions-gate, questions-eval, eval-calibrate, deck-gate, bracket-gate, revise-gate, chat-probe, generate-probe, summary-judge, quality-judge, and test-smoke. `go run ./cmd/eval sweep -cap <USD>` drives five of them in the order of the eval list under a cap, and it needs `EVAL_SWEEP=1` (PR-15). `-dry` prints the plan for nothing, and the estimate of a step reads its last run file. Since PR-15 the deck gate spends one more judge call a deck, and run 16 cost $3.79 for 25 prompts.

Each other target is free. `make meta-refresh` reads the deck list sources over the network, about 40 minutes on the first run, and calls no model. `make quality-gate` fits the quality model over the local meta store and writes the PR-14B gate document. It also reports the three numbers of D-648: the fitted model grades the decks `QUALITY_GATE_DECKS` names, against the judge's tiers in `QUALITY_GATE_JUDGED`.

`make ste-check` checks every hand-written `.md` file against the STE rules, and `make lint` runs it. `make m5-sheet` builds the scoring sheet, and `make m5-report` reads it. `make themes-check` checks the theme slugs and the commander ranking.

`make eval-check` compares every baseline of the eval harness with its newest run and names the flips (PR-15). It is free.

`make verify` runs every check the verify workflow runs, on this machine, for nothing (D-578). Run it before every pull request. `make where` prints the branch, the tree, and the state of the branch's pull request. `make hooks` installs the pre-commit hook that refuses a commit on `main` (D-585).

`make feedback-harvest` writes every verdict since the last harvest to `docs/reference/feedback/`, as a dated document and a JSONL file (PR-28a). `SINCE=2026-09-01` sets the floor by hand, and `HARVEST_ARGS=-dry` counts and writes nothing. The watermark comes from the JSONL files, so the documents are the only record. It calls no model and costs nothing. Those files commit with the repository now (D-642).

`make feedback-triage-dry` routes every verdict of the newest harvest into a class (PR-28b). It calls no model and costs nothing. `make feedback-triage TRIAGE_OUT=<document>` asks the judge for the verdicts the reason keys can not place, at a few cents each. `TRIAGE_ARGS=-apply` writes each case into the gate file that owns it. Ask the owner before every live run.

`make feedback-loop` prints the commands of the fix cycle and starts nothing (PR-28c). `make feedback-loop-dry` plans a cycle for nothing. `scripts/feedback-loop.sh` is the paid cycle, and it refuses to start without `FEEDBACK_LOOP_ALLOW=1` and `AUTOTUNE_FIXER_CMD`. One cycle stops at $2 of gate runs (D-559). It edits code, commits, pushes, opens a pull request, and answers the review, with nobody watching. Ask the owner before every run.

`make users-backfill` seeds the user record of D-638 from what each user already holds, and `BACKFILL_ARGS=-dry` counts and writes nothing. It never lowers a count. It calls no model and costs nothing.

`make read-session SESSION=<id>` reads one chat session of the deployed project, to debug it (D-596). `make feedback-list` reads the newest verdicts of every user over one collection group query, and `VERDICT=up` and `LIMIT=` change what it reads. All three print or write what a reader wrote, and none of them writes an email. `use_decktome` puts the shell on `decktome-prod`. No decktome tool reads `PROJECT_ID`: a shell that works on more than one project exports it for another one.

`make smoke` runs the Playwright smoke flow of PR-23 over the emulators, the trimmed snapshot, and the fake provider (D-553). It calls no model, and it needs the Chromium build of Playwright once.

`make deck-gate-dry` builds every deck gate shortlist over the trimmed snapshot of the repo and calls no provider (D-521). `make deck-gate-trim` rewrites that snapshot from the local store, and the fixture must stay under 10 MB. Both are free.

`make allow EMAIL=... PROJECT_ID=...` invites one email to the deployed app, and `make disallow` takes one off (D-420). Both write one Firestore document of the deployed project with the caller's own credentials, and neither calls a model.

`make store-check` runs the session store against the local Firestore emulator. `make gcs-check` runs the live fake-GCS store test against a server seeded from the trimmed snapshot, and the CI step runs the same script (D-658). `make candidates-review` writes the PR-6 gate document from a local snapshot. `cd go && go run ./cmd/tune-check` compares an eval summary with its baseline.

`docs/reference/pr7-m5-scoring.md` is the owner's working copy. No target writes to it. A new sheet needs a new name and points at the latest gate document, for example `M5_OUT=docs/reference/pr7-m5-scoring-run18.md M5_RUNS=../docs/reference/pr7-question-gate-run18.md make m5-sheet`.
