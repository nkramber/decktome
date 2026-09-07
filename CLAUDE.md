# decktome - CLAUDE.md

Read this file first. Then read `docs/SESSION-HANDOFF.md`. It tells you where the last session stopped.

## Project

This repo is a Go + Protobuf + TypeScript monorepo for an agentic MtG deck builder. The app reads a user's ManaBox collection export. The user gives a prompt. The agent asks questions, then builds a legal, useful deck.

Stage (2026-09-04): PR-15, the eval harness, merged as #65 after its paid gate passed on every lane (D-514 to D-517). PR-21 merged as #62. Merged before it: PR-0a to PR-8, PR-7B, PR-10 to PR-13, the Phase 3B roadmap (#46), PR-16 (#47), PR-16B (#48), PR-17 (#49), PR-17B (#50), PR-18 (#53), the review fixes of PR-18 (#54), PR-19 (#55), its follow-ups (#56), PR-14A (#57), PR-14B (#58), PR-24 (#59), PR-14C (#60), and PR-20 (#61). PR-9 is out of the MVP (D-256).

**PR-14A, the bracket profile, is merged** (2026-09-02, #57, D-451 to D-453, D-459 to D-469). A bracket is a set of numbers now: the content rules per bracket, and a feature vector per built deck with a band per bracket. A goldfish simulation and a check against Commander Spellbook complete it. `make bracket-gate` is its gate. Run 1 reads FAIL on the band bar and the judge bar, and the bracket 5 misses are the power signal of PR-14B. Deck gate run 12 and its rerun 12b together pass all 24 prompts with no regression. **PR-14B, the deck quality model, is merged** (2026-09-03, #58, D-470 to D-493). Gate run 10 passes the top-list bar, deck gate 13b passes, and the tier judge bar stays open on a corpus finding (D-488). The owner merged with that on record (D-491). **The corpus step of 2026-09-03 ran** (D-494 to D-499). Quality gate run 11 reads FAIL on the Commander and the Modern precon bars, and 19 of 24 built decks grade bad. **PR-24, the precon exclusion, is merged** (2026-09-03, #59, D-496 to D-501). A reader asks for a deck that uses no card of a precon, by name or as "not from my precons", and the build leaves those cards out. Its free gate passes, `docs/reference/pr24-precon-gate-2026-09-03.md`. **PR-14C, MTGTop8 and the casual 60-card decks, is merged** (2026-09-03, #60, D-502 to D-506). Aetherhub is out behind Cloudflare, both lanes read with zero failures, and the Modern typical rung exists. Its gate, `docs/reference/pr14c-gate-2026-09-03.md`, reads FAIL on the two precon bars of quality gate run 12, on record. **PR-20, the deck view and the card detail, is merged** (2026-09-03, #61, D-507). It holds the card detail as a Sheet with rulings and printings, the filters and the sorts, five stats, and the sample hand. Its free gate passes, `docs/reference/pr20-gate-2026-09-03.md`. **PR-21, the share link and the print view, is merged** (2026-09-03, #62, D-508). Its free gate passes, `docs/reference/pr21-gate-2026-09-03.md`. The paid sweep of 2026-09-03 and 2026-09-04 passed every gate but the open judge bar, and its one defect, F-34, merged fixed as #63 (D-509). The corpus items of 2026-09-04 merged as #64 (D-510): the MTGJSON skip and deck gate prompt 25 for the precon exclusion. **PR-15, the eval harness, is merged** (2026-09-04, #65, D-511 to D-517). Every gate writes a run file beside its document. `make eval-check` compares each baseline with its newest run, and the sweep runs the paid suites under a cap. Its paid gate passed on every lane, and runs 35, 16, and 9 are the baselines of the questions, decks, and revise suites. `docs/reference/pr15-paid-gate-2026-09-04.md` holds the read with F-36 to F-39. **The set fixes merged as #66** (2026-09-04, D-518, D-525 to D-528, F-40 to F-42): the Marvel set question, the group set request, and the partial-run fix of the harness. Probe 109 passed as question gate run 37. **The M-5 sheet is complete** (2026-09-04, D-529, D-530): 60 of 60 items scored. The fit threshold stays at 0.35, because no candidate meets the 80 percent floor. The owner accepted every row on 2026-09-05. **The precon check ignores basic lands, merged as #67** (2026-09-05, D-532, D-533, F-35), the first of the three items of D-531. **The terse conversations merged as #68** (2026-09-05, D-541): the 47 terse conversations join the bar with expectations (D-522, D-534). Run 38 read FAIL on 13 misses, nine rows and four defects, all fixed (D-535, F-43 to F-46). Run 39 read one miss, the same colorless slot from a later turn. A closed color slot changes only on a color word now (D-536). Run 40 read one miss, a budget of 2 on a message with no number. A budget applies only when the message names it now (D-537). Run 41 read one miss, both spellings of a corrected commander in one turn, and an unknown name never joins a known commander now (D-538). Run 42 passes every bar and is the questions baseline (D-539). `docs/setup-gcp.md` is the deploy guide with the cost estimate, and OQ-65 asks the owner for the API path of PR-22. **The trimmed snapshot merged as #69** (2026-09-05, D-542, D-543): 3.9 MB, and the free dry lane of the deck gate in CI. The three items of D-531 are done. The owner answered OQ-65 on 2026-09-05 (D-544): PR-22 builds the direct Cloud Run origin. The reword guard stays at 0.60 after the M-5 measurement (D-545). PR #70 merged those records (D-546). The mobile and engagement proposal is `docs/reference/mobile-and-engagement-2026-09-05.md`. Stage A is PR-25, its own PR after PR-23 (D-547). Stage B is PR-26 and waits on OQ-67, and Stage C waits on request (D-548). PR #71 merged the proposal (D-549). PR #72 merged the four code parts of PR-22 on 2026-09-06 (D-550, D-551). They are the invite list, the spend cap, the Firebase web variables, and the hosting block. The deploy runs on the owner's account by `docs/setup-gcp.md`, and the deploy half of the PR-22 gate waits for it. **PR-23, the Playwright smoke flow, is merged** (2026-09-06, #73, D-552 to D-554). The fake provider serves a whole build from schema-keyed fixtures. `make smoke` runs the flow over the emulators for nothing, and the workflow `smoke` runs it on request. Its free gate passes, `docs/reference/pr23-gate-2026-09-06.md`, and the first run on GitHub passed in 4 minutes 42 seconds (D-554). PR-25, the installable web app, waits for the deploy of PR-22 (D-555). **The feedback system is the immediate item** (2026-09-06, D-557 to D-559). PR-27 harvests a thumbs up or down on questions, deck descriptions, cards, and decks. PR-28 addresses it with the fixer agent, by `docs/reference/feedback-2026-09-06.md`. **PR-27, the feedback harvest, is merged** (2026-09-07, #75, D-561, D-562). Its free gate passes, `docs/reference/pr27-gate-2026-09-06.md`, and the smoke run on the branch passed in 2 minutes 32 seconds. PR-28 waits for the first real feedback, and that waits for the deploy. **The repository is `decktome`** on GitHub and on disk, and the Go module is `github.com/nkramber/decktome/go` (2026-09-07, D-563). PR #76 merged the records and the rename the same day (D-564). **The corpus step of 2026-09-07 ran** (D-565, D-566): F-50, the EDHREC weekly skip that never held, is fixed with a test, merged as #77 (D-569). Quality gate run 13 reads FAIL on the precon bar of every format, and Standard joins the two on its new typical rung, `docs/reference/pr14b-quality-gate-run13.md`. `docs/reference/weak-axes-2026-09-07.md` is the plan for the precon bar: M-7, then PR-29 to PR-31 (D-567). The owner answered OQ-73 to OQ-75 the same day (D-568). **M-7 is done** on branch `m7-honest-bars` (D-571). Run 14 reads the honest bars (D-572). Each precon over its own copies reads 0.88, 0.97, and 0.99, and the cross pairs 0.86, 0.87, and 0.85. The owner chose the own copies as the bar (D-573), so Commander fails on synergy alone, and PR-31 parks. **The product is Decktome, and its domain is `decktome.com`** (2026-09-06, D-556). The owner bought it at GoDaddy, the DNS stays there, and `docs/setup-gcp.md` section 14 adds the Firebase records. No paid run goes before the PR-14C code lands (D-495). `make meta-refresh` fills the meta store over the network, and `make quality-gate` is the free gate. The Topdeck.gg key is in `.env` (OQ-54, D-479), and PR-14C brings MTGTop8 and the casual 60-card decks after PR-24 (D-482, D-490, D-493). The owner confirmed the PR-14A calls on 2026-09-02 (D-467 to D-469). PR-24 came after PR-14B, PR-20 to PR-23 follow PR-14C (D-460, D-494), and PR-15 came before PR-22 (D-511).

PR-19, the chat and build experience, is merged (D-432 to D-458, #55 and #56). Every gate passes: question gate 32, deck gate 11, revise gate 7, and the PR-17B set gate run 1.

Phase 3B, the product UI, is the current phase. The look follows a reference design the owner gave on 2026-08-30 (D-328 to D-335). `docs/SESSION-HANDOFF.md` holds the moving parts, and `docs/reference/autotune-readme.md` holds the loop commands.

Run `git pull`, then `git status`, before you change anything. The owner commits and pushes. Do not commit unless the owner asks.

Read `docs/SESSION-HANDOFF.md` next. It is the resume point.

## Hard rules from the owner

1. **Write scope.** This repo permits writes (D-32, 2026-08-23). All other repos are read-only. Application code follows the roadmap order. Do not start a roadmap item before the gate of the item before it holds.
2. **Write in ASD-STE100.** Every doc, skill, and agent file must follow Simplified Technical English. Load the `ste-writing` skill before you write. Rules that apply most: max 20 words per procedural sentence, max 25 per descriptive sentence, and active voice. Also: one instruction per sentence, no semicolons, no "-ing" verb forms, one term per concept, paragraphs of max six sentences.
3. **Ask questions when you think of them.** Do not save questions for the end. Use `AskUserQuestion` in small batches. Record each answer in `docs/decisions.md`.
4. **Do the research.** Verify facts against sources (Scryfall API, Wizards announcements, the Comprehensive Rules). Record the date of each fact. MtG rules and ban lists change often.
5. **Make hand-off simple.** Before you end a session, update `docs/SESSION-HANDOFF.md`: the completed work, the open work, and the next step.
6. **No AI-attribution text** in any PR, branch name, commit message, or comment. This house rule comes from connector-syncer.
7. **No mistakes.** Check card names, rules, and dates before you write them. When you are not sure, say so and mark the item as unverified.
8. **Every change starts on a branch.** Never commit to `main`, and never push to it (D-583). Make a branch, commit there, push it, and open a pull request. The owner merges.
9. **Deploy from `main` alone.** Never deploy any other branch to production, for any reason (D-579). Check the branch and the commit before every build, not only the tree. `docs/deploy-and-rollback.md` holds the procedure.
10. **Never hesitate to ask or to push back.** Ask a question the moment you have one. When the owner's two statements conflict, say so and quote both. When a request rests on a wrong premise, say so with the evidence. The owner sees this as the key to good LLM-user interaction. Silence is the mistake, not the question.

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

Each other target is free. `make meta-refresh` reads the deck list sources over the network, about 40 minutes on the first run, and calls no model. `make quality-gate` fits the quality model over the local meta store and writes the PR-14B gate document. `make ste-check` checks every hand-written `.md` file against the STE rules, and `make lint` runs it. `make m5-sheet` builds the scoring sheet, and `make m5-report` reads it. `make themes-check` checks the theme slugs and the commander ranking.

`make eval-check` compares every baseline of the eval harness with its newest run and names the flips (PR-15). It is free.

`make smoke` runs the Playwright smoke flow of PR-23 over the emulators, the trimmed snapshot, and the fake provider (D-553). It calls no model, and it needs the Chromium build of Playwright once.

`make deck-gate-dry` builds every deck gate shortlist over the trimmed snapshot of the repo and calls no provider (D-521). `make deck-gate-trim` rewrites that snapshot from the local store, and the fixture must stay under 10 MB. Both are free.

`make allow EMAIL=... PROJECT_ID=...` invites one email to the deployed app, and `make disallow` takes one off (D-420). Both write one Firestore document of the deployed project with the caller's own credentials, and neither calls a model.

`make store-check` runs the session store against the local Firestore emulator. `make candidates-review` writes the PR-6 gate document from a local snapshot. `cd go && go run ./cmd/tune-check` compares an eval summary with its baseline.

`docs/reference/pr7-m5-scoring.md` is the owner's working copy. No target writes to it. A new sheet needs a new name and points at the latest gate document, for example `M5_OUT=docs/reference/pr7-m5-scoring-run18.md M5_RUNS=../docs/reference/pr7-question-gate-run18.md make m5-sheet`.
