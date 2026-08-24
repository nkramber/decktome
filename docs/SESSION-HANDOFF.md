# Session hand-off

Read this file first in a fresh session. Then read `CLAUDE.md`, `docs/decisions.md`, and `docs/open-questions.md`.

## Last updated

2026-08-24. Merged: PR-0a (#1), PR-0b (#2), PR-0c (#3). PR-1 built on branch `pr-1`, gate held, commit and merge pending.

## State of the work

Done in session 1:
- Read the reference doc `connector-syncer-docs/docs/document-summary-roadmap.md` and captured its structure in the `design-doc-style` skill.
- Surveyed connector-syncer's AI, eval, and ops patterns. Notes in `docs/reference/`.
- Surveyed wallabee-ui and CI conventions. Notes in `docs/reference/wallabee-toolchain-conventions.md`.
- Researched Scryfall (API, rate limits, bulk files, image terms, catalogs, Oracle tags), ManaBox CSV columns, the 2026 ban-list changes, Standard legal sets, Commander rules and brackets, and ASD-STE100 Issue 8.
- Wrote the `mtg-corpus` skill (terminology, formats, ban snapshot, archetypes, roles, slang, ManaBox format, question catalog, validation checklist) with two generated reference files.
- Wrote the `ste-writing` skill (all 53 rules).
- Answered the owner's local-Firestore question in `docs/reference/local-dev-environment.md`.
- Wrote `docs/design-roadmap.md` draft 1 with 15 findings, 12 guardrails, 5 phases, and strict sequencing.
- Recorded 14 decisions (D-1 to D-14) and 12 open questions (OQ-1 to OQ-12).
- Pass 2: the owner answered OQ-1 to OQ-12. Recorded D-15 to D-25. Updated the roadmap (correction pass 2), the corpus skill, and CLAUDE.md (rule 8: never hesitate to ask or push back). New open questions OQ-13 to OQ-17.
- Pass 3: OQ-13 to OQ-17 answered (D-26 to D-30). Added M-5 (manual scoring lane) and rewrote I-1 (stale-deck banner and scoped rerun). New OQ-18, OQ-19. OQ-17 waits on the owner's files.

## Where we stopped

PR-1 is done on branch `pr-1`. The contract lives in nine files under `proto/mtg/v1/`: `card`, `format`, `collection`, `deck`, `session`, and four service files (`CardService`, `CollectionService`, `DeckService`, `AgentService` with the streaming `Chat`). Generated code is committed for both stacks. The gate held: both compile, `make proto-check` passes, lint and tests are green.

Design notes for the reviewer: legalities are an open map with Scryfall keys. `PowerLevel` is a oneof (bracket, or the 60-card step, D-8). `Format` has `FORMAT_ID_HOUSE` plus `house_rules` text (D-3). `Slots` mirrors the PR-7 question catalog. `ImportReport` returns unresolved rows so nothing is dropped in silence (D-23). `Question` carries `invented` and `gap_score` (D-25).

A domain re-pass on 2026-08-24 (owner-requested) added colorless produced mana, parsed types, commander eligibility fields, the sideboard, the companion, entry rarity, and printing images. It also found F-18: Scryfall can not express "banned as a companion", so the rules engine owns that check. Reviewed and deliberately skipped: a `target_meta` slot (the free-text `theme` holds it for now) and name-search on `CardService.Search` (additive later).

## Next steps, in order

1. The owner reviews the `.proto` files, commits `pr-1`, and merges. A field rename after merge is a breaking change, so review now.
2. PR-2: the card database from Scryfall bulk. It needs no owner input to start.
3. Collect contract corrections as dated register entries.

## Facts that expire

- Ban-list snapshot: 2026-08-23. Next announcement 2026-10-12.
- Scryfall bulk sizes and counts: 2026-08-23.
- Game Changers: 53 cards, list of 2026-02-09.
- Standard: 12 legal sets, no rotation in 2026.

## How to resume

1. Load skills `ste-writing`, `design-doc-style`, and `mtg-corpus`.
2. Read the three docs above.
3. Continue from "Next steps". Ask questions as they come up.
4. Before you end, update this file.
