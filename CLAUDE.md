# mtg-deck-builder - CLAUDE.md

Read this file first. Then read `docs/SESSION-HANDOFF.md`. It tells you where the last session stopped.

## Project

This repo is a Go + Protobuf + TypeScript monorepo for an agentic MtG deck builder. The app reads a user's ManaBox collection export. The user gives a prompt. The agent asks questions, then builds a legal, useful deck.

Stage (2026-08-25): PR-0a to PR-6 and PR-10 are merged (#1 to #11), the audit fixes included.

PR-7 (the question workflow) is on branch `pr-7`. The owner scored items 1 to 32 of `docs/reference/pr7-m5-scoring.md`. Those scores asked for 16 rewords and one deletion, and they exposed a class of defect the sheet could not hold. A correction session of 2026-08-25 fixed it, and a batch sweep of all 66 conversations followed. Together they record D-104 to D-132. Gate run 14 is the next step, and it costs money.

The correction is not committed. Run `git pull`, then `git status`, before you change anything: the owner reads and scores the M-5 sheet on a phone, so the remote can be ahead. The owner commits and pushes. Do not commit unless the owner asks.

Read `docs/SESSION-HANDOFF.md` next. It is the resume point.

## Hard rules from the owner

1. **Write scope.** Writes are allowed in this repo (D-32, 2026-08-23). All other repos are read-only. Application code follows the roadmap order. Do not start a roadmap item before its predecessor's gate holds.
2. **Write in ASD-STE100.** Every doc, skill, and agent file must follow Simplified Technical English. Load the `ste-writing` skill before you write. Rules that apply most: max 20 words per procedural sentence, max 25 per descriptive sentence, active voice, one instruction per sentence. Also: no semicolons, no "-ing" verb forms, one term per concept, paragraphs of max six sentences.
3. **Ask questions when you think of them.** Do not save questions for the end. Use `AskUserQuestion` in small batches. Record each answer in `docs/decisions.md`.
4. **Do the research.** Verify facts against sources (Scryfall API, Wizards announcements, the Comprehensive Rules). Record the date of each fact. MtG rules and ban lists change often.
5. **Make hand-off simple.** Before you end a session, update `docs/SESSION-HANDOFF.md`: what is done, what is open, and the next step.
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

## Commands that cost money

`make questions-gate` calls the real providers. One run of the 66 conversations costs about $0.12 and takes about 14 minutes. The 52-conversation run cost $0.09 and took 11 minutes, and D-105 added 14. Ask the owner before every run, and write to a new `GATE_OUT` file: a rerun must never overwrite a scored document (D-65).

`make questions-eval` scores a gate run with the eval role. One 66-conversation run costs about eleven cents. `make eval-calibrate` measures the eval model against a stronger one for about thirty cents. `scripts/autotune.sh` is the overnight tuning loop, and it refuses to start without `AUTOTUNE_ALLOW_UNATTENDED=1`. Read `docs/reference/autotune-design.md` first.

Everything else is free. `make m5-sheet` builds the scoring sheet, `make m5-report` reads it, `make themes-check` checks the theme slugs and the commander ranking, and `make store-check` runs the session store against the local Firestore emulator.

`docs/reference/pr7-m5-scoring.md` is the owner's working copy. No target writes to it. The version-2 sheet needs a new name: `M5_OUT=docs/reference/pr7-m5-scoring-run14.md make m5-sheet`.
