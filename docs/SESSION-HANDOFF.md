# Session hand-off

Read this file first in a fresh session. Then read `CLAUDE.md`, `docs/decisions.md`, and `docs/open-questions.md`.

## Last updated

2026-08-24. Merged: PR-0a to PR-5 (#1 to #8). Phase 1 is complete. PR-10 built on branch `pr-10`, gate held (unit tests and live smoke), merge pending.

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

Done in the PR-10 session (2026-08-24):
- `internal/llm` is the role layer (D-1, guardrail 3). See the roadmap PR-10 entry for the full shape.
- Decisions D-38 to D-41: Anthropic judge, baseline models (luna, terra, sonnet-5), official Go SDKs, keys in `.env`.
- New files: `roles.json`, `prices.json`, `client.go`, `openai.go`, `anthropic.go`, `fake.go`, `usage.go`, `env.go`, `errors.go`, `config.go`, tests, and `smoke_test.go`.
- `cmd/api` builds the client at startup and logs each role's provider and model.
- `make test-smoke` runs the live smoke. `make llm-defaults-check` runs in CI and warns on a default change.
- `.env.example` documents the keys. `scripts/dev.sh` sources `.env`.

## Where we stopped

PR-10 is complete. Lint is clean, 15 unit tests pass over the fakes, and the live smoke passed against both vendors (2026-08-24, $0.0013). The owner's keys live in `.env`.

Known limits, recorded in the roadmap entry: Anthropic reports no reasoning-token count. No call site uses the layer yet. PR-7 and PR-8 add the first ones, with real role fixtures under `internal/llm/fixtures/`.

## Owner directive 2026-08-24 (D-37)

The collection is optional. A user with zero library gets a fully optimized deck from the whole legal pool. A user with a library can turn the library off. The proto needed no change. The roadmap, guardrail 5, the corpus question catalog, and the PR-6/7/8/11/12 entries were updated.

## Next steps, in order

1. The owner reviews and merges `pr-10`.
2. PR-6 (candidates), PR-7 (questions), PR-8 (generator), PR-9 (variance). PR-7 is the first call site of `internal/llm`.

## Facts that expire

- Ban-list snapshot: 2026-08-23. Next announcement 2026-10-12.
- Scryfall bulk sizes and counts: 2026-08-23.
- Game Changers: 53 cards, list of 2026-02-09.
- Standard: 12 legal sets, no rotation in 2026.
- LLM model ids and prices: 2026-08-24 (`roles.json`, `prices.json`). Sonnet 5 intro price ends 2026-08-31.

## How to resume

1. Load skills `ste-writing`, `design-doc-style`, and `mtg-corpus`.
2. Read the three docs above.
3. Continue from "Next steps". Ask questions as they come up.
4. Before you end, update this file.
