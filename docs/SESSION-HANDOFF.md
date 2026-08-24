# Session hand-off

Read this file first in a fresh session. Then read `CLAUDE.md`, `docs/decisions.md`, and `docs/open-questions.md`.

## Last updated

2026-08-24. PR-0a merged (#1), PR-0b merged (#2). PR-0c built on branch `pr-0c`, gate held, commit and merge pending.

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

PR-0b is done: `docs/setup.md` (ten-step procedure) and the version-aware `scripts/doctor.sh`. The gate held 2026-08-24 with Docker 29.7.2 installed and the daemon running. The owner commits and merges `pr-0b`. Next: PR-0c.

## Next steps, in order

1. The owner commits `pr-0c` and merges.
2. Phase 1 starts with PR-1 (proto contract v1). Ask the owner to review the roadmap Phase 1 entries before you write the proto.
3. Collect roadmap corrections as dated entries.
2. Ask OQ-18 (rerun depth rule) and OQ-19 (scoring rubric) when the design reaches I-1 and M-5. OQ-17 files come from the owner later.
3. Run the STE checker (`docs/tools/ste-check.py`) on every changed `.md` file and fix findings.
4. Expand the `mtg-corpus` skill: archetype-to-card examples per format, the Karsten color-source table, and a sample ManaBox export fixture (owner-provided).
5. When the owner approves the plan, PR-0a starts. Until then, no application code.

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
