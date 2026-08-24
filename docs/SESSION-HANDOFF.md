# Session hand-off

Read this file first in a fresh session. Then read `CLAUDE.md`, `docs/decisions.md`, and `docs/open-questions.md`.

## Last updated

2026-08-24. Merged: PR-0a to PR-3 (#1 to #6). PR-4 built on branch `pr-4`, gate held, merge pending.

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

PR-4 is done on branch `pr-4`. Packages: `internal/collections` (header-driven ManaBox CSV parser, Arena text parser, resolver with the join order Scryfall ID then set+collector then exact name, Firestore repo with gzip payloads) and `internal/collectionsvc` (the Connect handler). The cards index gained `BySetCollector`. The owner's real export is the committed gate fixture.

Gate held end to end: 2,548/2,548 rows resolved against the full snapshot through the running API, with zero silent drops. Re-upload dedup by content hash, Get, and List are green. The corpus records the verified 18-column header.

Auth debt: requests act as one debug user (`DEBUG_USER_ID`, default `local-dev`) until Firebase Auth lands in PR-11. The UI upload screen also waits for PR-11.

## Owner directive 2026-08-24 (D-37)

The collection is optional. A user with zero library gets a fully optimized deck from the whole legal pool. A user with a library can turn the library off. The proto needed no change. The roadmap, guardrail 5, the corpus question catalog, and the PR-6/7/8/11/12 entries were updated. PR-3 content is unchanged and still awaits merge.

## Next steps, in order

1. The owner reviews and merges `pr-4`.
2. Phase 1 ends with PR-5 (the rules engine). It needs no owner input. Its golden decks (30 good, 30 bad) are the main work.
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
