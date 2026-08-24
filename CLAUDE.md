# mtg-deck-builder - CLAUDE.md

Read this file first. Then read `docs/SESSION-HANDOFF.md`. It tells you where the last session stopped.

## Project

This repo is a Go + Protobuf + TypeScript monorepo for an agentic MtG deck builder. The app reads a user's ManaBox collection export. The user gives a prompt. The agent asks questions, then builds a legal, useful deck.

Stage (2026-08-23): **exploratory design only**. No application code exists yet. The deliverable is `docs/design-roadmap.md`.

## Hard rules from the owner

1. **Read-only outside two paths.** You may write only in `docs/` and in `.claude/` (skills, agent files) of this repo. Do not write to any other repo. Do not write application code until the owner says so.
2. **Write in ASD-STE100.** Every doc, skill, and agent file must follow Simplified Technical English. Load the `ste-writing` skill before you write. Rules that apply most: max 20 words per procedural sentence, max 25 per descriptive sentence, active voice, one instruction per sentence. Also: no semicolons, no "-ing" verb forms, one term per concept, paragraphs of max six sentences.
3. **Ask questions when you think of them.** Do not save questions for the end. Use `AskUserQuestion` in small batches. Record each answer in `docs/decisions.md`.
4. **Do the research.** Verify facts against sources (Scryfall API, Wizards announcements, the Comprehensive Rules). Record the date of each fact. MtG rules and ban lists change often.
5. **Make hand-off simple.** Before you end a session, update `docs/SESSION-HANDOFF.md`: what is done, what is open, and the next step.
6. **No AI-attribution text** in any PR, branch name, commit message, or comment. This house rule comes from connector-syncer.
7. **No mistakes.** Check card names, rules, and dates before you write them. When you are not sure, say so and mark the item as unverified.

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
- `docs/reference/` - research notes with sources and dates.
