# Session hand-off

`CLAUDE.md` is the entry point. It sends you here. Then read `docs/decisions.md` and `docs/open-questions.md`.

## Last updated

2026-08-24. Merged: PR-0a to PR-5, PR-10, the audit fixes, and PR-6 (#1 to #11). PR-7 is in progress, and its work is not committed.

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

## PR-6 state (2026-08-24)

`internal/candidates` is built and tested. `cmd/candidates-review` writes the gate document. Scoring follows D-62 (payoffs over enablers) and D-64 (a theme line behind a condition counts). D-63 gives PR-7 the thin-theme question.

Run 1 (`docs/reference/pr6-candidate-review.md`): 12 of 20 on theme. The bar is 18, so run 1 fails. Eight prompts failed: 2, 3, 5, 6, 14, 16, 18, and 19. The document keeps every score and Review block.

Run 2 (`docs/reference/pr6-candidate-review-run2.md`): 20 of 20 on theme, so the gate holds. Ten lists are identical to run 1. Ten changed, and each one has a new Review block. The weakest pass is prompt 19 (control) at exactly 36 of 40, with six borderline cards. A rerun writes a new file, because `REVIEW_OUT` refuses a document that already holds a verdict (D-65).

Fixes between the two runs:
- Defect A: `themes.json` named 16 tag slugs that Scryfall Tagger does not have, over 11 theme rows. The matcher dropped each one without a message. Every slug is now real. `make themes-check` runs `TestThemeSlugsExist` against the local snapshot and fails on an unknown slug.
- Defect B: two payoff needles missed the common card wording. "deals damage to each opponent" matched 11 red cards, and the real wording with a number matches 176. Needles are now short and wording-safe.
- Defect C: the role fallback in `roles.go` called every "exile target" clause removal, so blink spells got role removal. `exileIsRemoval` now skips a clause that returns the permanent. `TestExileIsRemoval` guards it.
- Theme rows: a parent tag carries its children, which was the root cause in four prompts. `death-trigger` (aristocrats), `anthem` (tokens), `flicker` (blink), and `counters-matter` (counters) are gone. Payoffs are now narrow: `blood-artist-ability`, `synergy-token`, `reanimate-creature`, `storm-like`, `pp-counters-matter`, `burn-player`, `synergy-mill`, `synergy-poison`.

Known limits: popularity is EDHREC rank for every format, so a 60-card prompt collects Commander staples. Prompt 18 measures the cost: every Modern Burn staple except Lightning Bolt ranks below the cut. PR-14 (`MetaBoost`) is the fix. Prompt 16 shows the second limit: the word "proliferate" scores under the word "counters", so no proliferate card reaches the top 40.

## PR-7 preparation (2026-08-24)

OQ-19 is answered (D-66): six fields per invented question, three-point scales, and an 80% precision floor on the threshold M-5 sets.

Eight dogfood conversations ran through the `deck-builder-dogfood` agent before any code. None was catalog-only, and only 5 of 26 catalog questions survived without a rewrite. The catalog held the right slots and the wrong wording. One cause produced three of the invented questions: the Commander row offered suggestions and held no question to close the slot.

Corpus section 11 was rewritten from those runs. It holds 26 rows, against 11 before: 11 new rows, two rows split in two, an ask order, and a word-routing rule. Section 2.2 gained the Grist class of commander, which a type-line test reads wrong (Scryfall ruling, 2021-06-18).

New decisions: D-66 (rubric), D-67 (the card-pool question waits for format, colors, and theme), D-68 (slots freeze when a run starts), D-69 (the gap score is its own call, first threshold 0.35).

## PR-7 state (2026-08-24)

PR-7 is the question workflow and the first call site of `internal/llm`. It is **in progress**. `internal/questions` holds the whole turn loop, 2,648 lines with tests. `go test ./...` is green and `go vet ./...` is clean. Nothing is committed by this session.

### What exists

| File | What it holds |
|---|---|
| `catalog.json` | The 26 rows of corpus section 11 as data: ask order, triggers, options, and a placeholder-free fallback for each of the 9 rows that use braces. |
| `catalog.go` | Load and validate. A row needs a unique id, a real proto slot, a unique order, a text, and a fallback when it holds a placeholder. |
| `plan.go` | The turn planner and the word-routing rules. |
| `state.go` | Session state, the slot map, the freeze, and `Ready`. |
| `agent.go` | The turn loop: classify, score, ask. The agent decides the source of every question. |
| `prompts.go` | The three prompts and their strict schemas. Keep them stable: a changed prompt invalidates the M-5 scores (D-66). |
| `resolve.go` | The placeholder resolver, the clause dropper, and the phrasing guard. |
| `hints_candidates.go` | PR-6 as the value source: theme colors, commander names, owned on-theme count. |

Rules the code holds:

- At most three questions in one turn, and one question per proto slot.
- Never a repeat.
- Nothing at all when the run is frozen (D-68).
- The card-pool question waits for format, colors, and theme (D-67).

A row carries two names. `slot` is the proto field the answer informs. `key` is the row's own state. A refinement question such as table tolerance therefore survives a filled power slot.

A turn costs three model calls (D-69): `classify` fills slots, a second `classify` call scores the catalog fit and may offer a replacement, and `ask` phrases the result. The provisional fit threshold is 0.35. A replacement counts only under it.

### The live run (2026-08-24)

The owner approved one live conversation. Four calls, 1,875 input tokens, 350 output tokens, **$0.000795**, 9.7 seconds, both roles on `gpt-5.6-luna`, zero cached input tokens. A second run went out by mistake in the same session, at about the same size, so the true spend was near $0.0016.

It found four defects that 33 offline tests had missed:

1. The agent trusted the classifier's list of closed slots. One over-eager list ended a session with commander, power, and card pool still empty. Fixed: a key closes only when the agent offered it that turn, and a refused key is logged.
2. The agent sent raw placeholders to the model. The model answered by turning "{theme} is strongest in {colors}" into a second question aimed at the user. Fixed: resolve first, drop a clause with no value, fall back when the first sentence does not survive, and guard the phrasing that comes back.
3. A surviving trailing clause read as a dangling question. Fixed by the first-sentence rule.
4. `Ready` called a session complete while its questions were still unanswered, because the no-repeat rule empties the plan as soon as a question goes out. Fixed: ready needs an empty plan and no outstanding ask.

Defects 3 and 4 came out of the fixes for 1 and 2, so the live call paid for itself four times.

### How to run it

- `go test ./internal/questions` runs everything offline against the fake provider.
- `make themes-check` checks the PR-6 theme slugs against the local snapshot.
- The live conversation needs approval before every run. It costs money:
  `set -a && . ./.env && set +a && QUESTIONS_LIVE=1 go test ./internal/questions -run TestLiveConversation -v -count=1`

### Open in PR-7

- `AgentService.Chat` and `GetSession`, plus the session store. The proto is ready (PR-1b).
- The M-4 report. Every question already logs its slot, source, fit, and threshold.
- 18 more gate conversations. Twelve exist. The catalog-only half of the gate needs the model in the loop, because only the model invents a question.
- Prompt caching is off: the live run read zero cached input tokens, although both prompts hold a stable prefix.
- Known rough edge: the locked-cards row fires whenever the user names a card, including a card that became the commander.

## Next steps, in order

1. Finish PR-7: `AgentService.Chat`, the session store, the M-4 report, and 18 more gate conversations.
2. PR-8 (generator), then PR-9 (variance).
3. M-5 runs on the first UI build (after PR-12) and sets the D-27 threshold from the D-66 rubric.

## Facts that expire

- Ban-list snapshot: 2026-08-23. Next announcement 2026-10-12.
- Scryfall bulk sizes and counts: 2026-08-23.
- Game Changers: 53 cards, list of 2026-02-09.
- Standard: 18 sets, Wilds of Eldraine (2023-09-08) to The Hobbit (2026-08-14). No rotation in 2026. Six sets leave at the first 2027 set: WOE, LCI, MKM, OTJ, BLB, DSK. Verified 2026-08-24 on the Scryfall sets API.
- LLM model ids and prices: 2026-08-24 (`roles.json`, `prices.json`). The Sonnet 5 intro price claim is unverified.
- Comprehensive Rules: 2026-08-07 text. Commander brackets: 2025-10-21 revision.

## How to resume

1. Load skills `ste-writing`, `design-doc-style`, and `mtg-corpus`.
2. Read the three docs above.
3. Continue from "Next steps". Ask questions as they come up.
4. Before you end, update this file.
