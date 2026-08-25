# Decisions

Every owner decision, with date. This file is the source of truth. The design doc cites it by D-number. Do not re-ask a question that has a row here. When a later decision changes an earlier one, keep both rows and mark the old one "superseded".

| # | Date | Question | Decision |
|---|---|---|---|
| D-1 | 2026-08-23 | LLM provider | Provider-agnostic role-to-model layer. The provider is a config choice. Model: connector-syncer `llm_roles.py`. |
| D-2 (amended by D-37) | 2026-08-23 | Card pool | Build from the user's owned cards first. Suggest upgrades and acquisitions separately. |
| D-3 | 2026-08-23 | "60-card anything goes" | Not a fixed format. The agent asks the user what it means. Flexibility is the principle. The owner personally means "any card, no ban list". |
| D-4 | 2026-08-23 | Audience and deployment | Multi-user web app on GCP. Mirror the Wallabee stack: Cloud Run, Firestore, GCS. |
| D-5 | 2026-08-23 | Meta data sources | Allowed: official MTGO decklists, MTGGoldfish, MTGTop8, Aetherhub, EDHREC. **2026-08-23 update (OQ-10): the legal check passed. Deck data from any of these sites may be used.** |
| D-6 | 2026-08-23 | Card images | Hotlink Scryfall image URIs. Show artist and copyright. Do not crop, skew, or watermark. |
| D-7 | 2026-08-23 | RPC style | Connect-RPC with buf. One `.proto` contract. Generated Go and TypeScript clients. Server streaming for chat. |
| D-8 | 2026-08-23 | Power scale | Commander: official brackets 1-5. 60-card: three steps, casual / FNM / tournament-meta. |
| D-9 | 2026-08-23 | Local test environment | The full stack, UI included, must run locally. Firestore emulator confirmed feasible. Local testing has priority over hosting. |
| D-10 | 2026-08-23 | Docker | The owner will install Docker. The install is a roadmap item. |
| D-11 | 2026-08-23 | Auth | Firebase Auth, with the Auth emulator in local mode. |
| D-12 | 2026-08-23 | Output language | All docs, skills, and agent files in ASD-STE100. |
| D-13 (superseded by D-32) | 2026-08-23 | Write scope | This session writes only in `docs/` and `.claude/` of this repo. All other repos are read-only. |
| D-14 | 2026-08-23 | Question cadence | Ask questions as they come up, in small batches. Never batch them for the end. Never hesitate to ask, or to call out an inconsistency. |
| D-15 | 2026-08-23 | Deck export targets (OQ-1) | ManaBox text first. Other formats later. |
| D-16 | 2026-08-23 | Collection storage (OQ-2) | Store the full collection (name, printing, quantity, finish, condition, rarity, set, collector number). A hash can not recover any of this. The hash is kept only to detect an identical re-upload. |
| D-17 | 2026-08-23 | Prices (OQ-3) | USD. Source: Scryfall daily near-mint market estimate, 7-day rolling average, outliers rejected. Show the lowest price across all legal printings and finishes. Label it "NM market estimate". Scryfall has no lightly-played tier. A condition-tiered source (TCGplayer API) is a later option, not planned. |
| D-18 | 2026-08-23 | Determinism and variance (OQ-4) | First answer: identical collection plus identical prompt should ideally give the same deck, and divergence through follow-up questions is acceptable. **Second answer, same day, supersedes it: keep random variance. Identical-input determinism is not a requirement.** The seed is still stored with each deck so a build can be reproduced on request. |
| D-19 | 2026-08-23 | Card explanations (OQ-5) | Yes. One summary paragraph on the deck's style and purpose. One line per card. |
| D-20 | 2026-08-23 | Play simulator (OQ-6) | Later, not at launch. Stays in Phase 5. |
| D-21 (extended by D-38) | 2026-08-23 | Day-one provider (OQ-7) | OpenAI on day one, behind the provider-agnostic role layer. |
| D-22 | 2026-08-23 | Judge model (OQ-8) | A different provider from the generator. The exact provider is chosen at implementation time. |
| D-23 | 2026-08-23 | Languages (OQ-9) | English only. Non-English rows in an import are reported to the user, not silently dropped. |
| D-24 | 2026-08-23 | GCP projects (OQ-11) | `mtg-dev` and `mtg-prod`. No domain name and no hosting yet. Local testing first. |
| D-25 | 2026-08-23 | Question source (OQ-12) | Both. Draw from the fixed catalog whenever possible. Add a gap score that detects when no catalog question fits. Allow model-invented questions when the gap score says one is needed. The catalog is a living artifact: extend, change, or reduce it from user feedback and evals. |
| D-26 | 2026-08-23 | Price outlier rule (OQ-13) | Yes: drop a day whose price is more than 2x the 7-day median. |
| D-27 | 2026-08-23 | Gap-score threshold (OQ-14) | Needs evals, gated on manual testing by the owner. The owner uses the product and scores each model-invented question against the catalog questions that could have been asked instead. The threshold is set from those scores. This is a complex, first-class part of the deliverable. |
| D-28 | 2026-08-23 | Catalog change approval (OQ-15) | The owner approves catalog changes to start. |
| D-29 | 2026-08-23 | Stale decks on a rule change (OQ-16) | Show a banner on a stored deck that now contains a card illegal in its format. Give a "rerun" button. The rerun rebuilds according to the nature of the change: which cards were banned and how deeply they affect the deck. |
| D-30 | 2026-08-23 | Sample ManaBox exports (OQ-17) | The owner will provide them later. |
| D-31 | 2026-08-23 | Roadmap approval | The owner approved `docs/design-roadmap.md` draft 1 (passes 1-3). PR-0a may start. |
| D-32 | 2026-08-23 | Write scope (supersedes D-13) | Writes are allowed across this repo. Other repos stay read-only. |
| D-33 | 2026-08-23 | Go module and layout | Module `github.com/nkramber/mtg-deck-builder`. One Go module at `go/` with `cmd/` and `internal/`. |
| D-34 | 2026-08-23 | CI and PR flow | GitHub Actions. PRs into `main`, squash merge, branch protection. `verify:*` matrix with a fan-in job. |
| D-35 | 2026-08-23 | Scaffold toolchain pins (PR-0a) | Go 1.26.4. buf 1.72.0, protoc-gen-go 1.36.12, protoc-gen-connect-go 1.20.0 as `go tool` directives, with buf built into `.bin/`. Node 20.17.0 (installed version, moved to 22.12 LTS by D-52), pnpm 9.2.0. Vite 7.3.6 (Vite 8 needs rolldown native bindings that pnpm 9.2 did not install). React 19.2.8. Vite dev port 5180 (5173 is used by wallabee-ui). golangci-lint v2.6.2. |
| D-36 | 2026-08-24 | Local port map (PR-0c) | This machine also runs the Wallabee stack (flash :8080, Firestore emulator :8181, emulator UI :4000, vite :5173). Our stack uses its own block: API :8090 under `make dev` (:8080 default in production images), Firestore emulator :8281, Auth emulator :9199, emulator UI :4100, hub :4490, logging :4590, fake-gcs-server :4443, web :5180. |
| D-37 | 2026-08-24 | Collection is optional (amends D-2) | The app must build decks with zero user library. A user with a library can deselect it ("use only cards in my library" off) to get a truly optimized deck. Pool modes: owned-first (default when a library exists), owned-only, any-card (default when no library exists, and the "optimized" mode). When a library exists but the user picks any-card, the deck still marks the cards the user happens to own, as information only. The proto already carries all three modes (`PoolRule`, PR-1). |
| D-38 | 2026-08-24 | Judge provider (implements D-22) | Anthropic serves the `judge` role. OpenAI serves the other four. Config validation refuses a judge on the generator's provider. |
| D-39 | 2026-08-24 | Baseline models (PR-10) | `classify` and `ask` on `gpt-5.6-luna`. `generate` and `repair` on `gpt-5.6-terra`. `judge` on `claude-sonnet-5`. This is the baseline for the PR-15 bake-off. The owner first picked `sol` and `claude-opus-5`, then the same day chose the middle tier as the baseline. |
| D-40 | 2026-08-24 | Provider SDKs | The official Go SDKs: `openai-go/v3` and `anthropic-sdk-go`. Their retries are off. The role layer owns the budget. |
| D-41 | 2026-08-24 | API key storage | Keys live in `.env` at the repo root (gitignored by the `.env.*` rule). `make dev` and `make test-smoke` source it. Never commit a key. Production sets `LLM_REQUIRE_KEYS=1` so a missing key is fatal, not a silent fake. |
| D-42 | 2026-08-24 | Audit answers (docs) | The roadmap is an approved, living document. A ✅ means the code is merged on `main`. `CLAUDE.md` is the entry point and leads into `SESSION-HANDOFF.md`. Research notes under `docs/reference/` are frozen history when they are out of date, with a banner. The STE script stays as it is (no passive-voice check). |
| D-43 | 2026-08-24 | Fixture set (OQ-17, closes D-30) | The owner's own export is the PR-4 fixture. Add public anonymized ManaBox exports when found (OQ-20). |
| D-44 | 2026-08-24 | Token rows in an import | A row whose printing is a token, emblem, or art card is reported as `NOT_PLAYABLE`. It is never resolved by name. |
| D-45 | 2026-08-24 | Judge effort | `judge` on `claude-sonnet-5` keeps thinking on at effort medium. Thinking tokens are counted. |
| D-46 | 2026-08-24 | Contract amendment | One PR-1b proto change carries every field PR-6 to PR-9 need. `buf breaking` runs in CI from now on. The PR-6 candidate list stays Go-internal until the UI needs it. |
| D-47 | 2026-08-24 | Announcement coverage (F-23) | A legality diff between snapshots decides coverage. The calendar only sets the poll cadence. |
| D-48 | 2026-08-24 | Worker shape | The worker is a Cloud Run job (`-once`) under Cloud Scheduler: hourly, and every 15 minutes on announcement days. The loop mode is for `make dev` only. |
| D-49 | 2026-08-24 | Partner variants (F-22) | `Card.partner_text` holds the text after "Partner—". Two Partner commanders need equal text. No new enum value. |
| D-50 | 2026-08-24 | Vehicle and Spacecraft commanders | Supported now per CR 903.3 (2026-08-07). |
| D-51 | 2026-08-24 | LLM keys | Keys are required by default. `LLM_REQUIRE_KEYS=0` opts out for local dev and tests. A `fake` provider is refused when keys are required. |
| D-52 | 2026-08-24 | Node version (amends D-35) | The Node 22 LTS line. Pinned at 22.23.2 (the current patch on 2026-08-24). Engines `>=22.12.0 <23`. Vite 7.3.6 needs 20.19 or 22.12 and up. jsdom 29 and up need 22.13 and up. |
| D-53 | 2026-08-24 | Golden gate size | At least 30 good and 30 bad decks, enforced by a test. The audit pass landed 41 and 53. |
| D-54 | 2026-08-24 | Corpus ban bullets | Keep the dated list and refresh it after each announcement. |
| D-55 | 2026-08-24 | Owned-only basics | Basic lands stay exempt from the ownership check in owned-only mode. The exception is recorded in the corpus checklist. |
| D-56 | 2026-08-24 | CI path filters | Struck. A skipped required check blocks a merge under branch protection. Every job runs on every PR. |
| D-57 | 2026-08-24 | Re-upload name | An identical re-upload with a new name takes the new name. |
| D-58 | 2026-08-24 | Collection entry fields | `language` and `set_name` are stored per entry. |
| D-59 | 2026-08-24 | Price history | Deferred to I-2. `price_as_of` is set now. The proto comment states today's rule and the D-17 target. |
| D-60 | 2026-08-24 | Small calls | Commander land guide is 34 to 38 (corpus wins, engine aligned). Commanders count toward the Game Changer limit, marked unverified. Checklist items on curve, roles, and plan are PR-8 model-side. ASD-STE100 Issue 8 stays the citation, with an Issue 9 note. A model override with no price row logs a startup warning. A canceled or expired caller context is `ClassBudget`. |
| D-61 | 2026-08-24 | Worker schedule and M-2 record | One Cloud Scheduler cron runs the worker job with `-once` every 15 minutes. The worker exits early when the latest snapshot is under an hour old and no announcement window is open. On other ticks it makes one small catalog GET and downloads only when Scryfall has a newer file. After a legality diff the worker writes `scryfall/<version>/legality_diff.json` so the M-2 lag outlives log retention. |
| D-62 | 2026-08-24 | Candidate scoring (PR-6) | Theme signals split into payoffs and enablers. A payoff tag (`lifegain-matters`) or needle ("whenever you gain life") outweighs an enabler tag (`lifegain`) or keyword (Lifelink). One tag per kind counts. A keyword covers its own name as a text needle. The owned-first gate view merges owned cards and upgrades by score. |
| D-63 | 2026-08-24 | Thin theme (PR-6, PR-7) | When an owned mode finds under 30 on-theme owned cards, the builder sets `ThinTheme`. PR-7 then asks the pool-mode question again: build owned-first with a buy list, or switch to any-card. |
| D-64 | 2026-08-24 | On-theme rule for a conditional line (PR-6 gate) | A card counts as on theme when it holds a real theme line, even behind a condition. Examples: Cosmic Intervention in blink, Grapeshot in burn. The Review block still names each borderline card and gives the total under the strict reading. |
| D-65 | 2026-08-24 | PR-6 gate failed, fix and rerun | The first run scores 12 of 20, and the bar is 18. Fix the three engine defects, repair `themes.json` for the eight failed prompts, then rerun `make candidates-review`. The first run keeps its file `docs/reference/pr6-candidate-review.md` with every score and Review block. The second run writes a new file, `docs/reference/pr6-candidate-review-run2.md`. A rerun never overwrites a scored document (owner directive, 2026-08-24). |
