# Session hand-off

`CLAUDE.md` is the entry point. It sends you here. Then read `docs/decisions.md` and `docs/open-questions.md`.

## Last updated

2026-08-24. Merged: PR-0a to PR-5 and PR-10 (#1 to #9). Phase 1 is complete. The full audit ran and every finding is fixed on branch `audit-fixes`, merge pending.

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

Done in the audit session (2026-08-24), branch `audit-fixes`:
- `docs/audit-2026-08-24.md` records every finding and its resolution. Six review passes, then five fix streams.
- PR-1b: the proto contract gained every field PR-6 to PR-9 need (F-25). `buf breaking` runs in CI. The string `ChatResponse.error` is deprecated in favor of `failure`.
- PR-4b: token rows report `NOT_PLAYABLE` (F-21). BOM, rune-safe truncation, quantity cap, unknown finish and condition values, duplicate merge, per-reason counts, `language` and `set_name`, physical line numbers. `collectionsvc` has tests.
- PR-5b: five eligibility gaps fixed (F-22). Companion fully checked. Ownership per Oracle id. Golden gate 41 good and 53 bad. `DeckService.Validate` does the ownership check from the stored collection.
- Cards and ops: legality-diff coverage (F-23) with a marker object per version. The worker is a Cloud Run job (D-61). The API listens first and refuses to start on Cloud Run without `ALLOW_DEBUG_USER=1`. Also: body limits, snapshot pruning, 429 handling, collision counts, per-face artist, `price_as_of`.
- PR-10b: thinking tokens counted, cache writes priced, keys required by default (`LLM_REQUIRE_KEYS=0` opts out), adapter tests with `httptest` (F-24).
- Scaffold: Node 22.23.2, `.dockerignore`, CI with SHA-pinned actions, `go build`, `buf breaking`, buf cache. `dev.sh` hardened. Compose stack verified end to end with a worker seed service. Web tests render `App` under jsdom with an axe check.
- Docs and corpus: every status current, D-42 to D-61 recorded, OQ-17 closed, OQ-20 opened, Standard set list and bracket table corrected, ban bullets refreshed.

Owner actions before merge:
1. `nvm install && nvm use` (Node 22.23.2), then `cd web && pnpm install`. Then `make test`.
2. Review branch `audit-fixes` and merge it as one PR.

Known limits: Grist, the Hunger Tide can not be a commander in the engine (no Scryfall signal). GCS `ListVersions` and `DeleteVersion` have no unit test. The ManaBox condition vocabulary beyond `near_mint` is unverified.

## Owner directive 2026-08-24 (D-37)

The collection is optional. A user with zero library gets a fully optimized deck from the whole legal pool. A user with a library can turn the library off. The proto needed no change. The roadmap, guardrail 5, the corpus question catalog, and the PR-6/7/8/11/12 entries were updated.

## Next steps, in order

1. The owner runs the two actions above and merges `audit-fixes`.
2. PR-6 (candidates), PR-7 (questions), PR-8 (generator), PR-9 (variance). PR-7 is the first call site of `internal/llm`.

## Facts that expire

- Ban-list snapshot: 2026-08-23. Next announcement 2026-10-12.
- Scryfall bulk sizes and counts: 2026-08-23.
- Game Changers: 53 cards, list of 2026-02-09.
- Standard: 12 legal sets, no rotation in 2026.
- LLM model ids and prices: 2026-08-24 (`roles.json`, `prices.json`). The Sonnet 5 intro price claim is unverified.
- Comprehensive Rules: 2026-08-07 text. Commander brackets: 2025-10-21 revision. Standard: 18 sets, six leave in 2027.

## How to resume

1. Load skills `ste-writing`, `design-doc-style`, and `mtg-corpus`.
2. Read the three docs above.
3. Continue from "Next steps". Ask questions as they come up.
4. Before you end, update this file.
