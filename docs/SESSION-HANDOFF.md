# Session hand-off

Read this file first in a fresh session. Then read `CLAUDE.md`, `docs/decisions.md`, and `docs/open-questions.md`.

## Last updated

2026-08-24. Merged: PR-0a (#1), PR-0b (#2), PR-0c (#3), PR-1 (#4). PR-2 built on branch `pr-2`, gate held, commit and merge pending.

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

PR-2 is done on branch `pr-2`. Packages: `internal/scryfall` (bulk client, real User-Agent, F-3), `internal/cards` (parse, derive, index, tags, snapshot store with completion marker), `internal/cardsvc` (CardService handler). The worker downloads the three bulk files on start and every six hours (`-once` mode serves `make dev-seed`). The API loads the newest complete snapshot at start and reloads on a timer (`CARDS_RELOAD_SECONDS`, 15 in dev). `/healthz` reports the snapshot date and age (M-2 seed).

The gate held: 200/200 tricky names resolve, and the full local stack serves real lookups and a lifegain search. Two findings were fixed on the way: F-19 (fake-gcs needs `storage.WithJSONReads()`) and F-20 (snapshot completion marker). Test fixtures live in `go/internal/cards/testdata/` (200 cards, 946 printings, 124 tags, built from the 2026-08-24 bulk).

## Next steps, in order

1. The owner reviews and merges `pr-2`.
2. PR-3 (legality freshness fast path) or PR-4 (ManaBox import) is next. PR-4 needs the owner's sample exports (OQ-17, D-30). PR-3 needs nothing.
3. Collect corrections as dated register entries.

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
