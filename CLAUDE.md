# mtg-deck-builder - CLAUDE.md

Read this file first. Then read `docs/SESSION-HANDOFF.md`. It tells you where the last session stopped.

## Project

This repo is a Go + Protobuf + TypeScript monorepo for an agentic MtG deck builder. The app reads a user's ManaBox collection export. The user gives a prompt. The agent asks questions, then builds a legal, useful deck.

Stage (2026-09-02): `main` is at `cf46951`. Merged: PR-0a to PR-8, PR-7B, PR-10 to PR-13, the Phase 3B roadmap (#46), PR-16 (#47), PR-16B (#48), PR-17 (#49), PR-17B (#50), PR-18 (#53), the review fixes of PR-18 (#54), PR-19 (#55), its follow-ups (#56), and PR-14A (#57). PR-9 is out of the MVP (D-256).

**PR-14A, the bracket profile, is merged** (2026-09-02, #57, D-451 to D-453, D-459 to D-469). A bracket is a set of numbers now: the content rules per bracket, and a feature vector per built deck with a band per bracket. A goldfish simulation and a check against Commander Spellbook complete it. `make bracket-gate` is its gate. Run 1 reads FAIL on the band bar and the judge bar, and the bracket 5 misses are the power signal of PR-14B. Deck gate run 12 and its rerun 12b together pass all 24 prompts with no regression. The session of 2026-09-02 built PR-14B, the deck quality model, on branch `pr-14b` (D-470 to D-488). Gate run 10 passes the top-list bar, deck gate 13b passes, and the tier judge bar stays open on a corpus finding (D-488). The owner merges it with that on record (D-491), and PR-14C gains the casual 60-card decks (D-490). `make meta-refresh` fills the meta store over the network, and `make quality-gate` is the free gate. The Topdeck.gg key is in `.env` (OQ-54, D-479), and PR-14C brings MTGTop8 and Moxfield back after PR-24 (D-482). The owner confirmed the PR-14A calls on 2026-09-02 (D-467 to D-469). PR-24 comes after PR-14B, then PR-20 to PR-23 (D-460).

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
8. **Never hesitate to ask or to push back.** Ask a question the moment you have one. When the owner's two statements conflict, say so and quote both. When a request rests on a wrong premise, say so with the evidence. The owner sees this as the key to good LLM-user interaction. Silence is the mistake, not the question.

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
- `docs/open-questions.md` - questions not yet asked or not yet answered.
- `docs/owner-questions.md` - the decision queue. Every question here waits for the owner, and the tuning loop refuses to decide one.
- `docs/reference/` - research notes with sources and dates, and every dated gate document.
- `docs/audit-2026-08-28.md` - the full audit of 2026-08-28, its owner answers, and the change plan.
- `docs/audit-2026-08-29.md` - the quality audit of 2026-08-29 and its fixes (D-302 to D-306).
- `docs/reference/set-data-2026-08-31.md` - every set number PR-17B rests on, with its source and date.
- `docs/reference/bracket-profile-2026-09-02.md` - every bracket rule, Karsten table, and Spellbook threshold PR-14A rests on.
- `docs/reference/deck-quality-model-2026-09-02.md` - every source fact, feature, and fit rule PR-14B rests on.

## Commands that cost money

`make questions-gate` calls the real providers. One run of the 104 conversations (30 gate and 74 probe) costs $0.152 to $0.165 and takes about 20 minutes, measured on 2026-08-26. Ask the owner before every run, and write to a new `GATE_OUT` file: a rerun must never overwrite a scored document (D-65).

`make questions-eval` scores a gate run with the eval role. One 104-conversation run costs $0.092 to $0.104 (runs 14 to 25) and takes about 13 minutes. `make eval-calibrate` measures the eval model against `claude-opus-5` (D-428). It cost $0.25 to $0.30 on Sonnet 5, and Opus 5 costs about 1.7 times that.

`make autotune` is free. It prints the loop instructions and starts nothing. `scripts/autotune.sh` is the paid loop, and it refuses to start without `AUTOTUNE_ALLOW_UNATTENDED=1`. One iteration costs about $0.25 and takes 33 to 35 minutes, so a $3 budget buys about 12 iterations. Read `docs/reference/autotune-readme.md` and `docs/reference/autotune-design.md` first.

Five more targets spend money, and each has an overwrite guard and an env guard. `make deck-gate` builds the PR-8 gate document. Run 11 cost $1.46 for 24 prompts, and run 12 cost $2.24 under the profile's repair passes. `make chat-probe` drives the real `Chat` RPC to a deck. `make generate-probe` builds one deck with the real generate role.

`make summary-judge` judges every deck summary of a gate document (F-26). Each probe costs a few cents. Ask the owner before every run.

`make bracket-gate` builds three commanders at each bracket, 15 decks, and asks the judge role for the bracket of each (PR-14A). It has the same two guards, `BRACKET_GATE=1` and a verdict check on `BRACKET_GATE_OUT`. Run 1 cost $2.08 for the builds and $0.26 for the judge lane. `BRACKET_GATE_ARGS="-only 7,8,9"` runs the bracket 3 prompts alone, `-rejudge <document>` judges the decks of a document for about $0.26, and `-dry` is free.

`make revise-gate` builds three base decks and runs nine revisions over them, twelve turns with the answered questions (PR-12B, D-448). It has the same two guards. Run 7 cost $0.74 for eleven turns, and runs 4 to 6 cost $0.54 to $0.81.

`DECK_GATE_ARGS` passes flags to `make deck-gate`. `DECK_GATE_ARGS="-only 19,20,21,22,23,24"` runs the six set prompts of PR-17B alone, for about $0.35.

`make quality-judge` asks the judge role for the tier of every graded deck of a deck gate document (PR-14B). It costs a few cents a deck, and it has the guard `QUALITY_JUDGE=1` and a verdict check on `QUALITY_JUDGE_OUT`. Ask the owner before every run.

`make test-smoke` runs the live LLM smoke test and reads the keys from `.env`. It spends a few cents. The paid targets are these eleven plus the script: questions-gate, questions-eval, eval-calibrate, deck-gate, bracket-gate, revise-gate, chat-probe, generate-probe, summary-judge, quality-judge, and test-smoke.

Each other target is free. `make meta-refresh` reads the deck list sources over the network, about 40 minutes on the first run, and calls no model. `make quality-gate` fits the quality model over the local meta store and writes the PR-14B gate document. `make ste-check` checks every hand-written `.md` file against the STE rules, and `make lint` runs it. `make m5-sheet` builds the scoring sheet, and `make m5-report` reads it. `make themes-check` checks the theme slugs and the commander ranking.

`make store-check` runs the session store against the local Firestore emulator. `make candidates-review` writes the PR-6 gate document from a local snapshot. `cd go && go run ./cmd/tune-check` compares an eval summary with its baseline.

`docs/reference/pr7-m5-scoring.md` is the owner's working copy. No target writes to it. A new sheet needs a new name and points at the latest gate document, for example `M5_OUT=docs/reference/pr7-m5-scoring-run18.md M5_RUNS=../docs/reference/pr7-question-gate-run18.md make m5-sheet`.
