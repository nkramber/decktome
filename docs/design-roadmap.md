# MtG Deck Builder - Design Roadmap

Status: **exploratory design doc, draft 1.** Nothing in this file is committed policy until the owner approves it and the item ships as a PR. The doc follows the structure of `connector-syncer-docs/docs/document-summary-roadmap.md`. It is written in ASD-STE100.

External facts were verified 2026-08-23. Sources: Scryfall (API and bulk data), the Wizards of the Coast announcements of 2026-08-10 and 2026-02-09, mtgcommander.net, and the local checkouts of connector-syncer and wallabee-ui. MtG rules and ban lists change. Re-verify every dated fact before you cite it in a PR.

Owner decisions live in `docs/decisions.md` (D-#). Open questions live in `docs/open-questions.md` (OQ-#). Research notes live in `docs/reference/`. The MtG knowledge base lives in `.claude/skills/mtg-corpus/`.

2026-08-23 correction pass 2: the owner answered OQ-1 to OQ-12 (D-15 to D-25). Changes in this pass: F-4 resolved, F-14 rewritten, F-16 and F-17 added, PR-4 storage decided. Also: PR-7 gains the gap score and M-4, PR-8 gains the deck summary, I-2 has a price spec, and the GCP project ids are set.

2026-08-23 correction pass 3: OQ-13 to OQ-17 answered (D-26 to D-30). Changes: M-5 added (manual scoring lane for invented questions), I-1 rewritten as the stale-deck banner and rerun, I-2 threshold fixed, section 9 updated.

House rule (from connector-syncer): no PR, branch name, commit message, comment, or other artifact may contain AI-attribution text.

---

## 1. Thesis

A deck builder is useful only when three things are true at the same time. The deck is legal on the day the user asks. The deck fits what the user owns and what the user wants. The deck works as a plan, not as a pile of good cards. A language model can do the third thing well. It can not do the first two things reliably without deterministic checks around it.

The system is therefore built as a **thin agent over a strict engine**. The engine owns card data, legality, ownership, and validation. The agent owns the conversation, the plan, and the card choices inside the engine's limits. Every card the model names passes through the engine before the user sees it. This is the same shape connector-syncer uses for citations: the model proposes, the code verifies, and hallucinations die at the boundary.

The program is sequenced so that each layer is testable before the next one exists. **Data and legality** come first (deterministic, cheap to test). Then the **agent loop** (measurable against the engine). Then the **UI** (thin over a streaming API). Then **meta and quality** (the expensive, judgment-heavy part). Each phase is gated on the one before it.

> *In plain English:* the AI is good at the creative question, such as a fun lifegain deck. It is bad at the exact question, such as a ban check. So the code answers the boring exact questions. The AI answers the creative ones. The AI never gets the last word on a card. The code checks every card before the user sees it.

## 2. Lessons learned (carried in from connector-syncer, and from this research)

1. **One concern per PR.** The reference roadmap lost a full attempt (#724) to a bundled change.
2. **Verify platform claims before you build on them.** Lesson 7 in the reference doc. Here: Scryfall rate limits, ManaBox column names, and emulator behavior were each checked against the real thing, not the docs alone.
3. **Model churn is a standing tax.** The role-to-model layer (D-1) and a bake-off protocol are the only durable answer. Record the resolved model in every eval row.
4. **Normalize model output against ground truth after every call.** A card name that does not exist is our hallucinated citation.
5. **Evaluate on the payload you change.** A prompt tested only on Commander tells nothing about Standard.
6. **Commit the evidence.** A/B scripts and golden sets live in the repo.
7. **Instrument cost on day one.** Every cost claim in the reference doc was an estimate until M-1 landed there.
8. **The user's words are not a spec.** "Anything goes" has no fixed meaning (D-3). The agent asks. This is a product principle, not a fallback.
9. **Facts have dates.** The ban list changed four times in 2026. Store facts with a verification date and a source.

> *In plain English:* these are the mistakes the sister project already paid for. We do not pay for them twice.

## 3. System map

| Component | Language | Owns | Reads | Writes | Sensitivity |
|---|---|---|---|---|---|
| `cards` service (card database) | Go | Scryfall snapshot, legalities, Oracle tags, images URIs | Scryfall bulk daily | Firestore `cards/`, GCS snapshot | High - every legality answer comes from here |
| `collections` service | Go | ManaBox import, ownership counts per Oracle ID | User CSV upload | Firestore `users/{uid}/collections/` | High - PII-adjacent, user data |
| `rules` engine (library) | Go | Format rules, deck validation, bracket rules, color identity | `cards` | none | Total - the last gate before the user |
| `agent` service | Go | Turn-based chat, question workflow, deck generation, LLM role layer | `cards`, `collections`, `rules`, `meta` | Firestore `users/{uid}/sessions/`, `decks/` | High - the product |
| `meta` service | Go | Metagame snapshots per format | MTGO decklists, aggregators (D-5), EDHREC | Firestore `meta/`, GCS raw | Medium - advisory input to the agent |
| `worker` | Go | Scheduled jobs: Scryfall refresh, meta refresh, ban-list watch | Cloud Scheduler, Cloud Tasks | see above | Medium |
| `web` (UI) | TypeScript, React, Vite | Chat, deck view with card art, collection upload, export | Connect-RPC API | none | Medium |
| `proto` | Protobuf | The one contract between Go and TypeScript | - | generated code, committed | High - a schema change is a cross-stack change |
| `eval` | Go + fixtures | Golden decks, deterministic checks, judge runs | all services in-process | BigQuery `evals.*` | Medium |

Three structural facts drive the plan:

- **The card database is small.** About 38,600 Oracle cards and 34,500 playable ones. It fits in memory in Go. We do not need a vector store to find cards. Scryfall Oracle tags (4,522 community tags, for example `lifegain` with 3,374 cards) plus keywords and type lines give a structured theme index. A vector index is a later option (I-3), not a foundation.
- **Legality is a lookup, not a judgment.** Scryfall updates `legalities` within a day of each announcement. The engine reads them. No prompt carries a ban list. This makes "legal as of the query date" a data-freshness problem, which is measurable (M-2).
- **Ownership is the constraint that makes the product.** Without the collection, the app is a worse EDHREC. With it, every suggestion is one the user can play tonight (D-2). The import must be robust to ManaBox column variance (F-2).

> *In plain English:* the whole card list of Magic fits in memory. The rules for "is this legal?" are a table lookup. The one thing no other tool has is the user's actual binder. So we build around the binder.

## 4. Cost model (what we expect, what we do not know)

- **LLM.** One deck-build session: 2 to 4 question turns (small model, about 2k tokens each) plus 1 to 3 generation turns (strong model, about 15k input with the candidate card list, 3k output). Estimate: under $0.10 per session on 2026 list prices. Unknown until M-1 measures it. Prompt caching of the format rules and the candidate list cuts the input cost. The role layer must expose the provider's caching knob (D-1, OQ-7).
- **Card data.** Scryfall bulk: 24.5 MB compressed per day for Oracle cards, 77.5 MB for all English printings. Free. Images hotlinked (D-6), zero storage. GCS: one snapshot per day, about 100 MB, cheap lifecycle to 30 days.
- **Meta data.** Unknown. Depends on source terms (OQ-10). MTGO decklists are official and free. Aggregator scraping may be zero or may be forbidden.
- **Firestore.** Per user: one collection doc set (a 5,000-card binder is about 5,000 small docs or one 500 KB doc, decision in PR-4), sessions, decks. Low.
- **Cloud Run.** Two services plus a worker, scale to zero. Low until users exist.
- **Eval.** Deterministic checks are free. Judge runs cost per deck. Cap per run as connector-syncer does ($5 cap in its bake-off).
- **Unknowns to measure first:** tokens per session (M-1), Scryfall refresh lag after an announcement (M-2), ManaBox import failure rate on real files (M-3).

> *In plain English:* the AI is the only real cost, and one deck should cost cents. Card data is free. Images are free because Scryfall lets us link to them. We will measure instead of guess.

## 5. Defect and finding register

Status: ✅ resolved · 🔧 planned or in progress (item listed) · 🅿 parked · ⚠ constraint on other work · ❓ needs owner input.

| # | Finding | Status |
|---|---|---|
| F-1 | **Ban lists drift fast.** Four B&R announcements in 2026 so far (03-23, 05-18, 06-29, 08-10). Next 2026-10-12. Commander changed 2026-02-09 with a new category, "banned as a companion". Any cached legality older than one day can be wrong on announcement day. | 🔧 M-2 (freshness metric) + PR-3 (daily refresh with announcement-day fast path) |
| F-2 | **ManaBox CSV columns vary.** The official guide does not list the columns. The verified column set (15 columns) comes from a third-party mapping. A whole-collection export adds a binder name column. Column order and presence can change with app versions. | 🔧 PR-4 (header-driven import, Scryfall ID first, tolerant of unknown columns) |
| F-3 | **Scryfall API rate limits are hard.** 2 requests per second on `/cards/named`, `/cards/search`, `/cards/collection`. 10 per second elsewhere. A 429 blocks for 30 seconds. Repeated overload gets a ban. Bulk files have no limit. | ⚠ binds PR-2: all card lookups go to the local snapshot. The live API is for single-card fallback only, behind a client-side limiter. |
| F-4 | **Aggregator terms of use unknown.** MTGGoldfish, MTGTop8, Aetherhub, and EDHREC have no public API and their terms were unchecked. | ✅ 2026-08-23: the owner confirmed the legal check passed (D-5). All five sources may be used. PR-14 still starts with MTGO because it is the only structured source. |
| F-5 | **Oracle tags are community data.** Scryfall Tagger tags are volunteer-made. Coverage is uneven. `lifegain` is rich (3,374 cards). Niche themes may have few tags. Weights are `median` style, not scores. | ⚠ binds PR-6: tags seed the candidate list. They never gate a card. Keywords and type lines are the second signal. The model is the third. |
| F-6 | **No Cloud Tasks emulator.** Local mode can not run real Cloud Tasks. | 🔧 PR-1: a `Dispatcher` interface with a local in-process implementation. |
| F-7 | **Docker absent on the dev machine.** | 🔧 PR-0b (owner installs Docker, D-10). Native `make dev` does not need it. |
| F-8 | **The mtgcommander.net banned-list page reads "last updated September 2024."** It does not show the 2026-02-09 changes. It is not a reliable source for the current list. | ✅ Scryfall `legalities.commander` is the source of truth. The page is for philosophy text only. |
| F-9 | **Double-faced and split cards have no top-level `image_uris`.** Images live in `card_faces[]`. Layouts `transform`, `modal_dfc`, `split`, `adventure` (about 800 Oracle cards). | 🔧 PR-2: the card model normalizes faces. The UI shows both faces. |
| F-10 | **"Anything goes" and other user phrases are ambiguous.** The owner confirmed this is by design (D-3). | ✅ product principle. PR-7's question catalog handles it. |
| F-11 | **Commander brackets are "beta" and change.** The 2026-02-09 update changed the Game Changers list. The bracket rules are prose, not data. | 🔧 PR-5 encodes brackets as data with a version date. The `game_changer` flag comes from Scryfall. |
| F-12 | **Legality is per Oracle card, but ownership is per printing.** A user may own a printing that is not legal in a format where the card is legal (for example a gold-bordered or Alchemy-rebalanced version). Scryfall marks these on the printing. | 🔧 PR-4 keeps the printing id. PR-5 checks legality on the Oracle card and flags the printing exception. |
| F-13 | **Model output can name a card that exists but is not the card meant.** Example: "Ajani's Pridemate" versus "Ajani's Welcome". Fuzzy matching hides this. | ⚠ binds PR-8: exact name match only, with the model asked for exact names. Fuzzy match is a suggestion to the user, never a silent substitution. |
| F-14 | **Variance versus determinism.** The owner wants variance between decks. A first answer to OQ-4 asked for identical output on identical input. A second answer the same day withdrew that: random variance stays (D-18). Note for the record: an LLM is not deterministic even at temperature 0, so identical output was never a guarantee. | ✅ resolved by D-18. PR-9 keeps a stored seed per deck for reproduction on request, and adds plan variants. |
| F-15 | **Standard has no rotation in 2026.** Rotation moves to the first set of 2027. Any hardcoded "September rotation" logic is wrong. | ✅ no rotation logic in code. Scryfall legalities carry it. |
| F-16 | **Scryfall prices have no condition tiers.** `prices.usd`, `usd_foil`, `usd_etched` are TCGplayer near-mint market estimates, updated once per day. The owner asked for lightly-played prices (OQ-3). No free source gives them. | ✅ D-17: show the NM estimate with a 7-day rolling average and outlier rejection, labeled as such. A condition-tiered source is a later option. |
| F-18 | **Scryfall legalities can not say "banned as a companion".** The 2026-02-09 Commander update unbanned Lutri, the Spellchaser but banned it as a companion. The Scryfall commander legality reads "legal". The open legalities map (guardrail 2) inherits this blind spot. | ⚠ binds PR-5: the rules engine owns the companion check. `Deck.companion_oracle_id` exists so the check has a target. Found in the 2026-08-24 proto re-pass. |
| F-17 | **A fixed question catalog can not cover every prompt.** The owner wants catalog questions first, model-invented questions when needed, and a metric that says which case applies (D-25). Without the metric, the agent either asks nothing new or bypasses the catalog. | 🔧 PR-7 (gap score) + M-4 (catalog coverage metric) + PR-15 (catalog-change proposals from evals). |

> *In plain English:* these are the traps we found before writing code. The biggest ones: ban lists change every few weeks. The collection file format is not documented. The AI can name a card that sounds right but is not. Each one has a planned fix or a rule that prevents it.

## 6. Guardrails (the safety contract for every PR)

1. **No card reaches the user before the rules engine has checked it.** Every generated list is validated for size, copies, legality on the query date, color identity, bracket, and ownership. A failed check blocks the response or marks the card, never silently drops it.
2. **No ban list, rotation date, or Game Changers list in any prompt or code constant.** Legality comes from the card database, which comes from Scryfall daily. Prompts may say "the engine will check legality".
3. **No model id at a call site.** All models come from the role layer (D-1). CI warns on a default change, as in connector-syncer.
4. **Exact card names only.** The model returns exact Oracle names. The normalizer does an exact match. Anything else becomes a user-visible suggestion, never a substitution (F-13).
5. **Owned cards first, always visible.** Every card in a deck carries an `owned` flag with the count. Acquisitions live in a separate list (D-2).
6. **The agent asks before it assumes** on format, power level, and house rules. Max three questions per turn. Defaults are allowed only when the user says "you decide" (D-3).
7. **Attribution on every image.** Artist and copyright are shown, images are not cropped or altered, and no paywall sits in front of card data (Scryfall terms, D-6).
8. **Generated proto code is committed and CI diffs it.** No hand edits. Pinned buf and plugin versions.
9. **Local mode has no cloud dependency.** Every service starts with emulators or fakes (D-9). A new cloud dependency must ship with its local fake in the same PR.
10. **One concern per PR. Evidence committed.** Golden decks and A/B outputs live in the repo.
11. **No raw user prompts in analytics.** Hashes and ids only, as in connector-syncer's event registry.
12. **Dated facts.** Every rules or format fact in the corpus carries a source and a verification date.

> *In plain English:* twelve promises every change must keep. The most important: the code, not the AI, has the final say on every card. And the app never quietly swaps a card the AI got wrong for one it guessed.

---

## 7. Roadmap

Ids: PR-# code, M-# measurement, I-# integration, D-# decisions (in `decisions.md`). Each entry has a gate and ends with a plain-English paragraph.

### Phase 0 - Foundations (no product code)

**PR-0a: Monorepo scaffold.** 🔧 built 2026-08-23 on branch `pr-0a-scaffold`, gate verified locally, PR not yet opened. Deviations from the plan, recorded in D-35: Vite 7 instead of 8, dev port 5180, buf built into `.bin/` from a `go tool` directive. The `verify:*` workflow exists but has not run on GitHub yet.
Layout: `proto/` (buf module), `go/` (Go workspace with `cmd/api`, `cmd/worker`, `internal/cards`, `internal/collections`, `internal/rules`, `internal/agent`, `internal/meta`, `internal/llm`), `web/` (pnpm workspace: `apps/web`, `packages/api-client` for generated TypeScript), `docs/`, `.claude/`. Makefile as the single entry point: `proto`, `lint`, `test`, `test-repeat`, `cover`, `dev`, `dev-seed`. Pinned versions: Go, buf, protoc-gen-go, protoc-gen-connect-go, protoc-gen-es, pnpm, Node, golangci-lint. CI: `verify:*` matrix with a fan-in job, path filters, and a proto-diff gate. AGENTS.md with the commands and never-edit rules. Gate: `make dev` starts an empty API and an empty UI.
> *In plain English:* the empty house with plumbing. One folder for the shared contract, one for Go, one for the web app. One command to start everything. The checks that stop bad changes are wired before there is anything to check.

**PR-0b: Developer machine setup, including Docker (D-10).** ✅ gate held 2026-08-24: the owner installed Docker 29.7.2, and `make doctor` reports all ok (12 checks). Merge pending.
A `docs/setup.md` procedure: install Homebrew, Git, Go, Node and pnpm via corepack, the firebase CLI, Java 17, Docker Desktop, and gcloud. A `make doctor` target checks each tool against the pinned version and prints the fix command. buf comes from `go/go.mod`, not from a separate install.
Docker is used for the Compose file (PR-0c) and for local Cloud Run parity. Gate: `make doctor` passes on the owner's machine.
> *In plain English:* a checklist to set up a laptop, and a command that tells you which tools are absent. The owner asked for the Docker install to be a tracked step, so it is one.

**PR-0c: Local stack (D-9).** ✅ gate held 2026-08-24 on branch `pr-0c`, merge pending. Both variants verified: native (`make dev`, all services up in 8 seconds, clean teardown) and containers (`make dev-docker`, Compose).

New in this PR: the port map D-36 (the Wallabee stack owns 8080, 8181, 4000, and 5173 on this machine), `internal/dispatch` (the Cloud Tasks stand-in, F-6), and `internal/llm` with the `Fake` provider.
`firebase.json` with Firestore and Auth emulators. `fake-gcs-server` for storage. A `Dispatcher` interface with a local in-process implementation (F-6). A `fake` LLM provider with fixture responses.

`make dev` runs all of it under one process supervisor. A Compose file gives the same stack in containers once Docker exists. GCP projects are `mtg-dev` and `mtg-prod` (D-24). No domain and no hosting yet. Local testing has priority. Gate: a developer with no GCP credentials runs the full stack and the UI loads.

Container note: the firebase emulator binds 127.0.0.1 from `firebase.json`. The emulator image rewrites the host to 0.0.0.0, or the published ports stay dead.
> *In plain English:* everything runs on the laptop with no cloud account: a fake database, fake file storage, a fake AI that returns canned answers. One command starts it all. The details are in `docs/reference/local-dev-environment.md`.

### Phase 1 - Data and rules (deterministic, fully testable)

**PR-1: Proto contract, v1.** ✅ gate held 2026-08-24 on branch `pr-1`, merge pending. Nine files under `proto/mtg/v1/`. Generated Go and TypeScript compile, and the CI diff gate passes. Contract notes: `legalities` is an open map keyed by Scryfall format keys (guardrail 2). RPC names are service-scoped (`GetDeck`, `GetCollection`) because message names share one proto package. The stream message is `ChatResponse` with a oneof event. `Question` carries `invented` and `gap_score` (D-25). `Deck` carries `seed` (D-18), `stale` (D-29), and `legality_as_of`.

Re-pass 2026-08-24 (owner-requested, against the MtG corpus) added: `COLOR_C` for produced mana, parsed `supertypes`/`card_types`/`subtypes`, `any_count_in_deck` (Relentless Rats class), commander eligibility (`can_be_commander`, `PartnerKind`, `partner_with_name`, `is_background`, `is_companion`), `Deck.sideboard` and `companion_oracle_id`, `CollectionEntry.rarity` (D-16), and `Printing.image_uris` plus `digital` (F-12, D-17).

The re-pass also produced F-18.
Messages: `Card`, `CardFace`, `Legality`, `Collection`, `CollectionEntry`, `Deck`, `DeckCard` (with `owned`, `owned_count`, `role`, `reason`), `Format`, `PowerLevel` (bracket or 60-card step, D-8), `Session`, `Turn`, `Question`, `Answer`, `ValidationResult`. Services: `CardService`, `CollectionService`, `DeckService`, `AgentService` (with a server-streaming `Chat` RPC). Connect-RPC with buf (D-7). Gate: generated Go and TypeScript compile. CI diff gate is green.
> *In plain English:* one document says what a card, a deck, and a chat message look like. Both the Go code and the web app read it. Change it in one place, and both sides update.

**PR-2: Card database from Scryfall bulk.**
A worker job downloads `oracle_cards` and `default_cards` daily (F-3: bulk only). It writes a versioned snapshot to GCS and an in-memory index in the `cards` service (name, Oracle ID, printing ID, legalities, color identity, keywords, type line, MV, produced mana, Oracle tags, `game_changer`, `edhrec_rank`, image URIs per face). Faces are normalized (F-9). The `oracle_tags` file loads into a tag tree. Rulings load on demand.

A `CardService.Lookup` by exact name, by Scryfall ID, and by Oracle ID. A `CardService.Search` with structured filters (colors, types, keywords, tags, format-legal). Gate: 100% of a fixed list of 200 tricky names resolve (split, DFC, "Aether" spelling, commas, apostrophes). Snapshot age is exposed as a metric.
> *In plain English:* every night we download the whole card list, keep a copy, and load it into memory. Anyone can ask "which green cards with lifelink are legal in Pioneer?" and get a fast exact answer with no AI involved.

**PR-3: Legality freshness and announcement-day fast path (F-1).**
The worker checks the Scryfall bulk `updated_at` every hour. On a B&R announcement day (a calendar the worker reads from a config, next 2026-10-12), it refreshes every hour until the legalities change. Every deck response carries `legality_as_of` (the snapshot date). The UI shows it. Gate: M-2 shows the lag between an announcement and the snapshot that reflects it.
> *In plain English:* ban announcements come on known dates. On those days we check more often. Every deck says which day's rules it was checked against, so the user knows.

**PR-4: ManaBox import (F-2, F-12).**
CSV parser driven by the header row, not by column position. Required: `Scryfall ID`, or `Set code` plus `Collector number`, or `Name` plus `Set name`. Optional: `Quantity`, `Foil`, `Condition`, `Language`, binder name. Unknown columns are ignored and logged once. Rows that do not resolve are returned to the user as a list, not dropped silently.

The result is a `Collection` with counts per Oracle ID and per printing. Also accepts the Arena text format (`4 Lightning Bolt (STA) 42`). Storage: the full collection is stored (D-16). One document per collection holds a compressed entry array (printing id, quantity, finish, condition, language). 

A per-Oracle-ID count map sits beside it for fast ownership checks. A content hash of the upload detects an identical re-upload. Non-English rows are reported to the user and skipped (D-23). Gate: a fixture set of real exports (owner-provided, anonymized) imports with zero silent drops. M-3 counts unresolved rows.
> *In plain English:* upload the file ManaBox gives you. We match every line to a real card and count how many you own. We show you the lines we could not match. We do not hide them.

**PR-5: Rules engine.**
A pure Go library. Inputs: a deck, a format, a power level, a collection, a card snapshot. Checks:
- deck size and copy limits (4, singleton, restricted),
- legality per card on the snapshot date,
- Commander eligibility and color identity,
- Game Changer count per bracket, with the bracket prose rules encoded as data with a version date (F-11),
- ownership counts,
- land count and color-source ranges per archetype, curve summary, and role coverage. Output: a `ValidationResult` with one finding per problem, each with a severity (`block`, `warn`, `info`). Formats and brackets are data files with `verified_at` dates. Gate: table-driven tests for every rule. A golden set of 30 known-good and 30 known-bad decks.
> *In plain English:* the referee. Give it a deck and the rules. It lists every problem: too few cards, a banned card, a card outside the commander's colors, a card you do not own. It never guesses. Every rule has a test.

**M-2: Legality freshness metric.** Snapshot age, and the lag between each announcement and the first snapshot that reflects it. **M-3: Import quality metric.** Unresolved rows per import, by reason.

### Phase 2 - The agent (gated on Phase 1)

**PR-6: Candidate-list builder.**
Given a format, colors, a theme, a power level, and the collection, build a ranked candidate list from the engine. Signals: Oracle tags (theme), keywords and type lines, `edhrec_rank` (popularity), legality, ownership. Output: about 150 to 300 owned candidates by role, plus about 50 unowned upgrade candidates (D-2). This list, not the whole database, is what the model sees. Gate: for 20 theme prompts, a human confirms the top 40 candidates are on-theme in at least 18.
> *In plain English:* before we ask the AI to build, the code shortlists the cards that fit: your cards, the right colors, on theme, legal. The AI picks from that list. It can not pick a card that is not there.

**PR-7: Question workflow.**
The turn-based core. A `Session` holds filled slots (format, commander, power, colors, theme, pool rule, budget, house rules, locked cards). Each turn: a small model classifies the prompt and fills slots it can. The code decides which slots are still empty and picks up to three questions from the catalog (`mtg-corpus` skill, section 11). The model phrases them.

The user answers in free text. The small model maps answers to slots.

Slots are stored, summarized, and carried to the next turn, as connector-syncer's schema agent does. The catalog is the first source of questions (D-25). A **gap score** decides when the catalog is not enough. It is the best catalog match between the empty slot and the user's words, from a small classifier. Below a threshold (OQ-14), the model may propose a question through a `custom_question` tool with a reason and the gap score. 

Every invented question is logged with its slot and outcome. M-4 reports how often this happens. Repeated invented questions become catalog candidates (PR-15). "Anything goes" and similar phrases route to the house-rules question (D-3).

Gate: 30 scripted conversations reach a complete slot set in at most four turns, with no repeated question. At least 25 of the 30 use catalog questions only. The gap-score threshold is set by M-5, not by this PR.
> *In plain English:* the chat. "Build me a lifegain deck" fills in "theme: lifegain" and leaves format, power, and colors empty. The app asks those three, remembers the answers, and never asks twice. If the user says something vague, the app asks what they mean. It does not guess.

**PR-8: Deck generator and normalizer (F-13).**
With all slots filled, the strong model gets four inputs. They are the rules summary for the format, the candidate list with roles, the role targets, and the plan request. It returns a structured deck (D-19). First, one summary paragraph on the deck's style and purpose. Then cards with exact names, counts, roles, and one line each. 

The normalizer exact-matches every name to the candidate list. A miss is returned to the model once as a tool error. A second miss becomes a user-visible note.

The engine validates (PR-5). A `block` finding triggers one repair turn with the findings as input. Then the deck goes to the user with the `ValidationResult` attached. Gate: on the golden prompts, 100% of returned decks pass `block` checks. Zero invented names reach the user.
> *In plain English:* the AI writes the deck from the shortlist, with a plan and a reason for each card. The code checks every name and every rule. If something is wrong, the AI gets one chance to fix it. What the user sees has already passed the referee.

**PR-9: Designed variance (F-14, D-18).**
Random variance is a feature (D-18). Variance comes from three levers, not from temperature alone. Lever 1: a seeded shuffle within each role tier of the candidate list. 

Lever 2: a "plan variant" slot (for example "lifegain aristocrats" versus "lifegain go-wide"). Lever 3: a "keep these, change the rest" re-roll. The seed is stored with the deck so a build can be reproduced on request. Identical output on identical input is not a requirement (D-18). Gate: two builds of the same prompt differ in at least 30% of nonland cards and both pass validation.
> *In plain English:* ask twice, get two different but sensible decks. Each deck remembers the dice roll that made it, so you can get the same deck back.

**PR-10: LLM role layer (D-1).**
`internal/llm` with roles: `classify`, `ask`, `generate`, `repair`, `judge`. A frozen config maps each role to a provider and model. Adapters: OpenAI on day one (D-21) and the `fake` provider. The `judge` role uses a different provider from `generate` (D-22), chosen at implementation time. Structured output through JSON Schema with strict validation.

One request budget per logical call with three retry classes (truncation, transient, terminal). Usage accounting per session that reports null when not instrumented. Prompt caching where the provider supports it. CI warns on default changes. Gate: unit tests with the fake provider. One real-provider smoke test behind an env flag.
> *In plain English:* the AI plug. Every place that calls an AI calls it through one door with a named job. Swap the vendor in one file. Count every token.

**M-1: Token and cost accounting per session.** Lands with PR-10. Every cost claim in this doc is an estimate until then.

**M-5: Manual scoring lane for invented questions (D-27, F-17).**
The owner uses the product on a fixed set of prompts. For each model-invented question, a review page shows four things. The question, the gap score, the slot, and the top three catalog questions that were possible instead. The owner scores it on a fixed rubric (OQ-19): was a catalog question good enough, was the invented question better, did it fill the slot. 

Scores go to the eval store with the prompt version and model. The gap-score threshold is chosen from these scores, and re-checked after each catalog change (D-28: the owner approves changes). Gate: at least 50 scored invented questions before the threshold is set.
> *In plain English:* the app sometimes has to make up a question. The owner will use the app, see each made-up question next to the fixed questions it could have used, and grade it. Those grades decide how eager the app is to make up questions.

**M-4: Catalog coverage metric (F-17).** Per session: catalog questions asked, invented questions asked, gap scores, and whether the invented question filled its slot. A weekly report lists invented questions by frequency. This is the input for catalog changes (D-25).
> *In plain English:* we count how often the app had to invent a question. If the same invented question appears again and again, it belongs in the fixed list.

### Phase 3 - UI (gated on PR-8)

**PR-11: Web app shell.**
React 19, Vite, TypeScript, Tailwind, the wallabee-ui patterns (TanStack Query, Zustand, lint-enforced import boundaries). Firebase Auth (D-11) with the emulator in local mode. Generated Connect client in `packages/api-client`. Gate: sign-in, upload a collection, see the count.
> *In plain English:* the website skeleton: log in, upload your binder, see how many cards we recognized.

**PR-12: Chat and deck view.**
A streaming chat thread over the `Chat` RPC. The deck view groups cards by role. It shows card art from Scryfall image URIs with artist and copyright (D-6, guardrail 7). It shows both faces for DFCs (F-9). It marks owned versus to-buy.

It shows the mana curve, the color sources, the `ValidationResult` findings, and `legality_as_of`. Hover or tap shows Oracle text. Gate: a11y checks pass. Every image has attribution in the DOM.
> *In plain English:* the main screen. The conversation on one side, the deck on the other with real card pictures, grouped by what each card does, with your own cards marked.

**PR-13: Export and share.**
Export as ManaBox text first (D-15). Other formats later. A buy list with Scryfall purchase links. Gate: a round trip ManaBox export to import loses nothing.
> *In plain English:* get the deck out of the app and into ManaBox or Arena with one click, plus a shopping list.

### Phase 4 - Meta and quality (gated on Phase 3 and on OQ-10)

**PR-14: Meta ingest, MTGO first.**
A worker job pulls published MTGO decklists per format (official source, D-5). It computes archetype shares and the most-played cards per archetype for the last 30 days. Aggregator and EDHREC ingesters follow, in order of structure: MTGTop8, MTGGoldfish, Aetherhub, EDHREC (D-5, legal check passed). The meta snapshot is advisory input to PR-6 and PR-8 for competitive power levels only. Gate: the snapshot for Modern lists at least 10 archetypes with card lists.
> *In plain English:* what wins right now. We start with the official tournament lists. Other sites are added only after someone checks their rules.

**PR-15: Eval harness.**
Golden prompts with expected slot sets and expected validation outcomes. Deterministic checks are the gate (legality, ownership, size, curve, names). A judge role scores plan quality and usefulness on a fixed rubric. Long-format results table, corpus fingerprint per run (model, effort, snapshot date, prompt version), suffix rows for informational metrics, "observe-only is not pass". Tier 0 in CI ($0).

Tier 1 nightly. Label-gated full sweep on PRs. Cost cap per run. Gate: the harness runs on PR-8's output and reports named regressions.
> *In plain English:* the test bench. Fixed questions, expected answers, a score every night. Any change that makes decks worse is named, not averaged away.

**I-1: Ban-list watch, stale-deck banner, and scoped rerun (D-29).**
A job reads the Wizards announcement feed and detects the Scryfall snapshot that reflects it. It then re-validates every stored deck in the affected formats. 

A deck with a now-illegal card gets a `stale` flag with the list of affected cards. The UI shows a banner on that deck with a "rerun" button.

An impact classifier scopes the rerun. Its inputs: how many cards are affected, which roles they filled, and whether the commander or a win condition is among them. 

Low impact: a patch turn that replaces only the affected cards from the same candidate list. 

High impact (threshold OQ-18): a full rebuild with the original slots and a new seed. The banner states which case applies and why. Gate: on the golden decks, every synthetic ban produces the correct case and a legal deck.
> *In plain English:* when Wizards bans a card, every deck we built that uses it gets a warning and a rerun button. If the ban only touches one filler card, we swap that card. If it guts the deck, we rebuild it from your answers.

**I-2: Price-aware buy list** (D-17, F-16). USD. For each card: the lowest Scryfall NM market price across legal printings and finishes, as a 7-day rolling average, with outlier days rejected (a day more than 2x the 7-day median, D-26). Digital-only and gold-bordered printings excluded. The UI labels it "NM market estimate" with the price date.

**I-3: Semantic card search** over Oracle text as a fourth candidate signal, only if PR-6's gate shows tags are not enough.

### Phase 5 - Parked (product decisions required)

- Sample-hand and goldfish simulator (D-20: later, not at launch).
- Per-card explanations longer than one line (D-19 gives one line per card).
- Non-English collections (D-23: English only for now).
- Brawl, Oathbreaker, Pauper Commander, Duel Commander.
- Sideboard builder for 60-card competitive play against a named meta.
- Collection sync from ManaBox without a file (no API exists on 2026-08-23).
- A public corpus API that serves the `mtg-corpus` content to the app's own prompts.

## 8. Sequencing - strict order, single owner

1. PR-0a scaffold.
2. PR-0b machine setup, Docker install (owner executes).
3. PR-0c local stack.
4. PR-1 proto v1.
5. PR-2 card database. Then M-2.
6. PR-4 ManaBox import. Then M-3.
7. PR-5 rules engine.
8. **GATE.** Phase 2 starts only when the golden decks pass PR-5.
9. PR-10 LLM role layer, with M-1.
10. PR-6 candidates.
11. PR-7 questions.
12. PR-8 generator.
13. PR-9 variance.
14. **GATE.** Phase 3 starts only when PR-8's gate holds on the golden prompts.
15. PR-11, PR-12, PR-13.
16. PR-15 eval harness (can start after step 12, in parallel with the UI, if a second owner exists). M-5 manual scoring runs on the first UI build (after PR-12).
17. PR-14 meta, then I-1, I-2, I-3 on evidence.
18. Phase 5 stays parked.

## 9. Open questions

See `docs/open-questions.md` for the full list with "ask when" dates. The ones that gate a phase:

1. **OQ-19 scoring rubric** gates M-5.
2. **OQ-17 sample ManaBox exports** (owner will provide, D-30) gate PR-4's fixture set.
3. **OQ-18 rerun depth rule** gates I-1.
4. PR-9's 30% variance number is a placeholder until PR-15 measures it.
