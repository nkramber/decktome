# MtG Deck Builder - Design Roadmap

Status: **approved (D-31), living document.** The owner approved draft 1 on 2026-08-23. Each PR entry carries its merge status. A ✅ means the code is merged on `main` (D-42). The doc follows the structure of `connector-syncer-docs/docs/document-summary-roadmap.md`. It is written in ASD-STE100.

External facts were verified 2026-08-23, with 2026-08-24 re-passes noted inline. Sources: Scryfall (API and bulk data), the Wizards of the Coast announcements of 2026-08-10 and 2026-02-09, mtgcommander.net, and the local checkouts of connector-syncer and wallabee-ui. MtG rules and ban lists change. Re-verify every dated fact before you cite it in a PR.

Owner decisions live in `docs/decisions.md` (D-#). Open questions live in `docs/open-questions.md` (OQ-#). The decision queue lives in `docs/owner-questions.md`. Research notes live in `docs/reference/`. The MtG knowledge base lives in `.claude/skills/mtg-corpus/`.

2026-08-26 correction pass 20: PR-7B added, the automated eval lane. The owner's hand scoring does not scale past 32 items, and a run asks about 250 questions. A new `eval` role scores every one for about eleven cents (D-133). Every third conversation is a holdout (D-134). The owner answered OQ-24, OQ-26, and OQ-27 the same day (D-135 to D-137). PR-7 gained sixteen fixes from the batch sweep of all 66 conversations (D-117 to D-132).

2026-08-25 correction pass 19: the owner scored items 1 to 32 of the M-5 sheet, and the scores asked for 16 rewords and one deletion. The correction session that followed recorded D-104 to D-116. The sheet could hold only 60 of 793 questions, which is where most defects hid (D-104).

2026-08-23 correction pass 2: the owner answered OQ-1 to OQ-12 (D-15 to D-25). Changes in this pass: F-4 resolved, F-14 rewritten, F-16 and F-17 added, PR-4 storage decided. Also: PR-7 gains the gap score and M-4, PR-8 gains the deck summary, I-2 has a price spec, and the GCP project ids are set.

2026-08-23 correction pass 3: OQ-13 to OQ-17 answered (D-26 to D-30). Changes: M-5 added (manual scoring lane for invented questions), I-1 rewritten as the stale-deck banner and rerun, I-2 threshold fixed, section 9 updated.

2026-08-24 correction pass 4 (owner directive): the collection is optional (D-37, amends D-2). Changes: thesis fact 3, guardrail 5, PR-6, PR-7, PR-8, PR-11, PR-12. No proto change: `PoolRule` and an optional `collection_id` existed since PR-1.

2026-08-24 correction pass 5 (full audit, `docs/audit-2026-08-24.md`): every PR status set to its merge state (#1 to #9). Register rows F-21 to F-25 added. PR-1b (contract amendment) added before PR-6. PR-3 added to section 8. Decisions D-42 to D-60 recorded. The fixes ship on branch `audit-fixes`.

2026-08-24 correction pass 6: PR-6 merged (#11) after a second gate run. Changes: PR-6 status and text, M-5 gains the OQ-19 rubric (D-66), section 9 open questions, and three engine defects recorded in the PR-6 entry. PR-7 gains the pool-question timing (D-67) and the slot freeze (D-68), both from eight dogfood conversations. The corpus question catalog was revised the same day.

2026-08-25 correction pass 19: the owner's scoring found a seventh fault and a hole in the reword guard (D-103). A truncation is now refused.

2026-08-25 correction pass 18: two gaps closed before PR-7 is committed. The session store now runs against the Firestore emulator (D-101), and `cmd/m5-report` reads the scored sheet and computes the thresholds (D-102).

2026-08-25 correction pass 17: runs 12 and 13 confirm D-98 and D-99 under live conditions. The M-5 sheet holds 60 items and explains its own fields (D-100). PR-7 has no open item but the owner's scoring.

2026-08-25 correction pass 16: 22 probe conversations joined the gate file (D-97). They found two defects on their first run, D-98 and D-99, and three catalog gaps. The M-5 sheet reads engine-current runs alone (D-96).

2026-08-25 correction pass 15: D-94 corrects the PR-6 commander helper, which never ranked commanders by score. Run 10 passes both bars. The weak-commander-pool trigger is wired, and PR-7 has no open item left but the owner's scoring.

2026-08-25 correction pass 14: run 9 adds the decline (D-93) and passes both bars. `State.Skip` was the last dead method in the package. The M-5 sheet holds 52 items, which passes the 50 the rubric asks for.

2026-08-25 correction pass 13: runs 7 and 8 found the root cause under D-83. A schema enum suppressed the format field, and the model answered "unknown" for a message that named the format outright (D-92). Both enum fields are free strings now.

2026-08-25 correction pass 12: gate runs 5 and 6. The catalog bar is met at 30 of 30, and a second bar now applies. D-91 fails a session that calls itself complete with a needed slot unanswered, which run 6 did five times. D-83 reached its third and final form, and D-90 makes a field the only record of an answer.

2026-08-25 correction pass 11: gate run 3 measured the fault checklist and found a regression in D-83. Changes: the PR-7 entry. D-86 corrects a side effect of D-79. D-87 deletes the acquisition row on the owner's challenge: the app can not act on a delivery date.

2026-08-25 correction pass 10: the PR-7 gate ran twice more. Run 2 scored 23 of 30, and it proved two fixes and exposed two deeper defects. Decisions D-83 to D-85 record them. The proto gained `Question.catalog_text`, which M-5 needs in the UI.

2026-08-25 correction pass 9: the PR-7 live gate ran and failed at 21 of 30. Changes: the PR-7 entry. Decisions D-78 to D-82 record the five defects the run exposed and their fixes. The worst one offered commanders outside the deck color identity.

2026-08-24 correction pass 8: PR-7 completed except for the live gate run. Changes: the PR-7 entry. Decisions D-70 to D-74 record the three items of the second live run, the "none" rule, and the session store. Four more dead rows were found while the service was wired.

2026-08-24 correction pass 7: PR-7 built in two halves. Changes: the PR-7 entry, PR-6 gains two helpers for PR-7, and D-69 records the gap-score call shape and its first threshold. One live conversation measured the cost of a turn.

House rule (from connector-syncer): no PR, branch name, commit message, comment, or other artifact may contain AI-attribution text.

---

## 1. Thesis

A deck builder is useful only when three things are true at the same time. The deck is legal on the day the user asks. The deck fits what the user owns and what the user wants. The deck works as a plan, not as a pile of good cards. A language model can do the third thing well. It can not do the first two things reliably without deterministic checks around it.

The system is therefore built as a **thin agent over a strict engine**. The engine owns card data, legality, ownership, and validation. The agent owns the conversation, the plan, and the card choices inside the engine's limits. Every card the model names passes through the engine before the user sees it. This is the same shape connector-syncer uses for citations: the model proposes, the code verifies, and hallucinations die at the boundary.

We sequence the program so that each layer is testable before the next one exists. **Data and legality** come first (deterministic, cheap to test). Then the **agent loop** (measurable against the engine). Then the **UI** (thin over a streaming API). Then **meta and quality** (the expensive, judgment-heavy part). Each phase waits for the one before it.

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
- **Ownership is a differentiator, not a requirement (D-37, 2026-08-24).** With a collection, every suggestion is one the user can play tonight. Without one, or with the library toggle off, the app builds a truly optimized deck from the whole legal pool. The pool mode is a first-class slot (`PoolRule`): owned-first, owned-only, any-card. The import must be robust to ManaBox column variance (F-2).

> *In plain English:* the whole card list of Magic fits in memory. The rules for "is this legal?" are a table lookup. The one thing no other tool has is the user's actual binder. So we build around the binder.

## 4. Cost model (what we expect, what we do not know)

- **LLM.** One deck-build session runs 2 to 4 question turns on the small model, at about 2k tokens each. It then runs 1 to 3 generation turns on the strong model, at about 15k input with the candidate card list and 3k output. Estimate: under $0.10 per session on 2026 list prices. Unknown until M-1 measures it. Prompt caching of the format rules and the candidate list cuts the input cost. The role layer must expose the provider's caching knob (D-1, D-21).
- **Card data.** Scryfall bulk: 24.5 MB compressed per day for Oracle cards, 77.5 MB for all English printings. Free. Images hotlinked (D-6), zero storage. GCS: one snapshot per day, about 100 MB, cheap lifecycle to 30 days.
- **Meta data.** Unknown. The source terms passed the legal check (D-5). MTGO decklists are official and free. Aggregator use is allowed.
- **Firestore.** Per user: one collection doc set (PR-4 decided one gzip document per collection, about 500 KB for a 5,000-card binder, D-16), sessions, decks. Low.
- **Cloud Run.** Two services plus a worker, scale to zero. Low until users exist.
- **Eval.** Deterministic checks are free. Judge runs cost per deck. Cap per run as connector-syncer does ($5 cap in its bake-off).
- **Unknowns to measure first:** tokens per session (M-1), Scryfall refresh lag after an announcement (M-2), ManaBox import failure rate on real files (M-3).

> *In plain English:* the AI is the only real cost, and one deck should cost cents. Card data is free. Images are free because Scryfall lets us link to them. We will measure instead of guess.

## 5. Defect and finding register

Status: ✅ resolved · 🔧 planned or in progress (item listed) · 🅿 parked · ⚠ constraint on other work · ❓ needs owner input.

| # | Finding | Status |
|---|---|---|
| F-1 | **Ban lists drift fast.** Four B&R announcements in 2026 so far (03-23, 05-18, 06-29, 08-10). Next 2026-10-12. Commander changed 2026-02-09 with a new category, "banned as a companion". Any cached legality older than one day can be wrong on announcement day. | ✅ PR-3 (#6). M-2 logs the lag. The audit replaced the calendar test with a legality diff (F-23, D-47). |
| F-2 | **ManaBox CSV columns vary.** The official guide does not list the columns. The verified set is 18 columns, from the owner's real export of 2026-08-24. A single-list export can drop the binder columns. Column order and presence can change with app versions. | ✅ PR-4 (#7): header-driven import, Scryfall ID first, unknown columns logged once. |
| F-3 | **Scryfall API rate limits are hard.** 2 requests per second on `/cards/named`, `/cards/search`, `/cards/random`, `/cards/collection`. 10 per minute on `/cards/manifest`. 10 per second elsewhere (verified 2026-08-24). A 429 blocks for 30 seconds. Repeated overload gets a ban. Bulk files have no limit. | ⚠ binds PR-2: all card lookups go to the local snapshot. The live API is for single-card fallback only, behind a client-side limiter. |
| F-4 | **Aggregator terms of use unknown.** MTGGoldfish, MTGTop8, Aetherhub, and EDHREC have no public API and their terms were unchecked. | ✅ 2026-08-23: the owner confirmed the legal check passed (D-5). All five sources may be used. PR-14 still starts with MTGO because it is the only structured source. |
| F-5 | **Oracle tags are community data.** Scryfall Tagger tags are volunteer-made. Coverage is uneven. `lifegain` is rich (3,374 cards). Niche themes may have few tags. Weights are `median` style, not scores. | ⚠ binds PR-6: tags seed the candidate list. They never gate a card. Keywords and type lines are the second signal. The model is the third. |
| F-6 | **No Cloud Tasks emulator.** Local mode can not run real Cloud Tasks. | ✅ PR-0c (#3): a `Dispatcher` interface with a local in-process implementation. |
| F-7 | **Docker absent on the dev machine.** | ✅ PR-0b (#2): Docker 29.7.2 installed. Native `make dev` does not need it. |
| F-8 | **The mtgcommander.net banned-list page reads "last updated September 2024."** It does not show the 2026-02-09 changes. It is not a reliable source for the current list. | ✅ Scryfall `legalities.commander` is the source of truth. The page is for philosophy text only. |
| F-9 | **Double-faced and split cards have no top-level `image_uris`.** Images live in `card_faces[]`. Layouts `transform`, `modal_dfc`, `split`, `adventure` (about 800 Oracle cards). | ✅ PR-2 (#5): the card model normalizes faces. The UI shows both faces (PR-12). |
| F-10 | **"Anything goes" and other user phrases are ambiguous.** The owner confirmed this is by design (D-3). | ✅ product principle. PR-7's question catalog handles it. |
| F-11 | **Commander brackets are "beta" and change.** The 2026-02-09 update changed the Game Changers list. The bracket rules are prose, not data. | ✅ PR-5 (#8) encodes brackets as data with a version date. The `game_changer` flag comes from Scryfall. Open: the prose rules (mass land denial, extra turns, combos) are an info finding, not a check. |
| F-12 | **Legality is per Oracle card, but ownership is per printing.** A user may own a printing that is not legal in a format where the card is legal (for example a gold-bordered or Alchemy-rebalanced version). Scryfall marks these on the printing. | ✅ PR-4 (#7) keeps the printing id. PR-5 checks legality on the Oracle card. The printing exception is an info finding when the printing data carries it (audit fix, 2026-08-24). |
| F-13 | **Model output can name a card that exists but is not the card meant.** Example: "Ajani's Pridemate" versus "Ajani's Welcome". Fuzzy matching hides this. | ⚠ binds PR-8: exact name match only, with the model asked for exact names. Fuzzy match is a suggestion to the user, never a silent substitution. |
| F-14 | **Variance versus determinism.** The owner wants variance between decks. A first answer to OQ-4 asked for identical output on identical input. A second answer the same day withdrew that: random variance stays (D-18). Note for the record: an LLM is not deterministic even at temperature 0, so identical output was never a guarantee. | ✅ resolved by D-18. PR-9 keeps a stored seed per deck for reproduction on request, and adds plan variants. |
| F-15 | **Standard has no rotation in 2026.** Rotation moves to the first set of 2027. Any hardcoded "September rotation" logic is wrong. | ✅ no rotation logic in code. Scryfall legalities carry it. |
| F-16 | **Scryfall prices have no condition tiers.** `prices.usd`, `usd_foil`, `usd_etched` are TCGplayer near-mint market estimates, updated once per day. The owner asked for lightly-played prices (OQ-3). No free source gives them. | ✅ D-17: show the NM estimate with a 7-day rolling average and outlier rejection, labeled as such. A condition-tiered source is a later option. |

> *In plain English:* these are the traps we found before writing code. The biggest ones: ban lists change every few weeks. The collection file format is not documented. The AI can name a card that sounds right but is not. Each one has a planned fix or a rule that prevents it.

| F-17 | **A fixed question catalog can not cover every prompt.** The owner wants catalog questions first, model-invented questions when needed, and a metric that says which case applies (D-25). Without the metric, the agent either asks nothing new or bypasses the catalog. | 🔧 PR-7 (gap score) + M-4 (catalog coverage metric) + PR-15 (catalog-change proposals from evals). |
| F-18 | **Scryfall legalities can not say "banned as a companion".** The 2026-02-09 Commander update unbanned Lutri, the Spellchaser but banned it as a companion. The Scryfall commander legality reads "legal". The open legalities map (guardrail 2) inherits this blind spot. | ⚠ binds PR-5: the rules engine owns the companion check. `Deck.companion_oracle_id` exists so the check has a target. Found in the 2026-08-24 proto re-pass. |
| F-19 | **The Go storage SDK's default download path 404s on fake-gcs-server.** The SDK reads objects through the XML API with percent-encoded names. The fake-gcs filesystem backend serves only the JSON paths for names with slashes. Listing works, reads fail. Found 2026-08-24 in the PR-2 smoke test. | ✅ fixed: every storage client passes `storage.WithJSONReads()`. JSON reads work on fake-gcs and on real GCS. |
| F-20 | **A snapshot version was listable before its files finished uploading.** The API loaded a mid-download snapshot and failed with "object doesn't exist". This is the reference project's F15 lesson: completion must imply artifacts. | ✅ fixed: the store writes a `complete` marker last. `LatestVersion` returns only marked versions. A test pins it. |
| F-21 | **A ManaBox token row resolved to the real card.** Row 186 of the owner's export is the Bloomburrow token `Pawpatch Recruit`. The resolver dropped tokens from the index, then fell through to the name lookup, which found the real creature. The PR-4 "2,548 of 2,548" gate counted it as resolved. Found in the 2026-08-24 audit. | ✅ audit fix: the index remembers dropped printings. A token, emblem, or art-card row reports `NOT_PLAYABLE` (D-44). |
| F-22 | **Five commander-eligibility gaps.** `Partner—[text]` variants collapsed into plain Partner. A lone Background passed. "Up to seven" and "up to nine" cards were blocked. A double-faced card qualified on its back face. Vehicles and Spacecraft (CR 903.3, 2026-08-07) were refused. Found in the 2026-08-24 audit against the Comprehensive Rules. | ✅ audit fix: `partner_text` and `max_copies_override` on the card, front-face rule, Vehicle and Spacecraft support (D-49, D-50). The golden gate grew to 41 good and 53 bad. |
| F-23 | **The announcement-day model used the calendar, not the data.** A snapshot from 09:00 UTC on announcement day counted as covered, hours before Wizards posted. The fast path never ran. | ✅ audit fix: coverage is a legality diff between snapshots (D-47). The calendar only sets the poll cadence. |
| F-24 | **The LLM layer under-counted the judge.** Anthropic thinking tokens were dropped, cache writes were priced at 1x instead of 1.25x, and `LLM_JUDGE_PROVIDER=fake` passed under `LLM_REQUIRE_KEYS=1`. Found in the 2026-08-24 audit of PR-10. | ✅ audit fix: thinking tokens counted, `cache_write` price column, keys required by default (D-51). |
| F-25 | **The proto lacked fields PR-6 to PR-9 need.** No upgrade list, no slot state, no question id, no structured answer, no seed override, no `Usage`, no per-face artist. | ✅ PR-1b (audit branch): all fields added in one contract amendment (D-46). `buf breaking` guards it from now on. |
## 6. Guardrails (the safety contract for every PR)

1. **No card reaches the user before the rules engine has checked it.** The engine validates every generated list for size, copies, legality on the query date, color identity, bracket, and ownership. A failed check blocks the response or marks the card, never silently drops it.
2. **No ban list, rotation date, or Game Changers list in any prompt or code constant.** Legality comes from the card database, which comes from Scryfall daily. Prompts may say "the engine will check legality".
3. **No model id at a call site.** All models come from the role layer (D-1). CI warns on a default change, as in connector-syncer.
4. **Exact card names only.** The model returns exact Oracle names. The normalizer does an exact match. Anything else becomes a user-visible suggestion, never a substitution (F-13).
5. **The pool mode is explicit, and ownership is always visible (D-2, D-37).** Every deck records its pool mode. When a collection is attached, every card carries an `owned` flag with the count, in every mode. Owned-first is the default with a library, any-card without one. The engine never silently narrows or widens the pool.
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

**PR-0a: Monorepo scaffold.** ✅ merged 2026-08-24 (#1, branch `pr-0a`). Deviations from the plan, recorded in D-35: Vite 7 instead of 8, dev port 5180, buf built into `.bin/` from a `go tool` directive. Node moved to 22.12 LTS in the audit (D-52).

Layout: `proto/` (buf module), `go/` (Go workspace with `cmd/api`, `cmd/worker`, `internal/cards`, `internal/collections`, `internal/rules`, `internal/agent`, `internal/meta`), `web/` (pnpm workspace: `apps/web`, `packages/api-client` for generated TypeScript), `docs/`, `.claude/`. Makefile as the single entry point: `proto`, `lint`, `test`, `test-repeat`, `cover`, `dev`, `dev-seed`. Pinned versions: Go, buf, protoc-gen-go, protoc-gen-connect-go, protoc-gen-es, pnpm, Node, golangci-lint. CI: `verify:*` matrix with a fan-in job, a proto-diff gate, and a `buf breaking` gate (audit). Path filters were planned and struck (D-56): a skipped required check blocks a merge. AGENTS.md with the commands and never-edit rules.

Gate: `make dev` starts an empty API and an empty UI.

> *In plain English:* the empty house with plumbing. One folder for the shared contract, one for Go, one for the web app. One command starts everything, and the checks that stop bad changes are wired before there is anything to check.

**PR-0b: Developer machine setup, including Docker (D-10).** ✅ merged 2026-08-24 (#2). The owner installed Docker 29.7.2, and `make doctor` reports all ok (12 checks).
A `docs/setup.md` procedure: install Homebrew, Git, Go, Node and pnpm via corepack, the firebase CLI, Java 17, Docker Desktop, and gcloud. A `make doctor` target checks each tool against the pinned version and prints the fix command. buf comes from `go/go.mod`, not from a separate install.
Docker is used for the Compose file (PR-0c) and for local Cloud Run parity. Gate: `make doctor` passes on the owner's machine.
> *In plain English:* a checklist to set up a laptop, and a command that tells you which tools are absent. The owner asked for the Docker install to be a tracked step, so it is one.

**PR-0c: Local stack (D-9).** ✅ merged 2026-08-24 (#3). Both variants verified: native (`make dev`, all services up in 8 seconds, clean teardown) and containers (`make dev-docker`, Compose).

New in this PR is the port map D-36. The Wallabee stack owns 8080, 8181, 4000, and 5173 on this machine. This PR also adds `internal/dispatch` (the Cloud Tasks stand-in, F-6), and `internal/llm` with the `Fake` provider.
`firebase.json` with Firestore and Auth emulators. `fake-gcs-server` for storage. A `Dispatcher` interface with a local in-process implementation (F-6). A `fake` LLM provider with fixture responses.

`make dev` runs all of it under one process supervisor. A Compose file gives the same stack in containers once Docker exists. GCP projects are `mtg-dev` and `mtg-prod` (D-24). No domain and no hosting yet. Local testing has priority. Gate: a developer with no GCP credentials runs the full stack and the UI loads.

Container note: the firebase emulator binds 127.0.0.1 from `firebase.json`. The emulator image rewrites the host to 0.0.0.0, or the published ports stay dead.
> *In plain English:* everything runs on the laptop with no cloud account: a fake database, fake file storage, a fake AI that returns canned answers. One command starts it all. The details are in `docs/reference/local-dev-environment.md`.

### Phase 1 - Data and rules (deterministic, fully testable)

**PR-1: Proto contract, v1.** ✅ merged 2026-08-24 (#4). Nine files under `proto/mtg/v1/`. Generated Go and TypeScript compile, and the CI diff gate passes. Contract notes: `legalities` is an open map keyed by Scryfall format keys (guardrail 2). RPC names are service-scoped (`GetDeck`, `GetCollection`) because message names share one proto package. The stream message is `ChatResponse` with a oneof event. `Question` carries `invented` and `gap_score` (D-25). `Deck` carries `seed` (D-18), `stale` (D-29), and `legality_as_of`.

The owner asked for a re-pass on 2026-08-24, against the MtG corpus. It added these fields: `COLOR_C` for produced mana, parsed `supertypes`/`card_types`/`subtypes`, `any_count_in_deck` (Relentless Rats class), commander eligibility (`can_be_commander`, `PartnerKind`, `partner_with_name`, `is_background`, `is_companion`), `Deck.sideboard` and `companion_oracle_id`, `CollectionEntry.rarity` (D-16), and `Printing.image_uris` plus `digital` (F-12, D-17).

The re-pass also produced F-18.
Messages: `Card`, `CardFace`, `Legality`, `Collection`, `CollectionEntry`, `Deck`, `DeckCard` (with `owned`, `owned_count`, `role`, `reason`), `Format`, `PowerLevel` (bracket or 60-card step, D-8), `Session`, `Turn`, `Question`, `Answer`, `ValidationResult`. Services: `CardService`, `CollectionService`, `DeckService`, `AgentService` (with a server-streaming `Chat` RPC). Connect-RPC with buf (D-7). Gate: generated Go and TypeScript compile. CI diff gate is green.
> *In plain English:* one document says what a card, a deck, and a chat message look like. Both the Go code and the web app read it. Change it in one place, and both sides update.

**PR-1b: Contract amendment (audit, D-46).** ✅ built 2026-08-24 on branch `audit-fixes`, merge pending. One proto change carries every field PR-6 to PR-9 need (F-25). The fields:
- Deck: `upgrades`, `buy_cost_usd`, and `DeckCard.price_usd`.
- Validation: `legality_as_of`, `pool_rule`, and `format` on the result. `pool_rule` and `collection_id` on the request.
- Session: `slot_states` with a `SlotState` enum, `Question.id`, `Answer`, `Turn.answers`, `status`, and `usage` with a `Usage` message (M-1).
- Chat: `answers`, `seed`, `keep_oracle_ids`, and an `AgentError` `failure` event. The string `error` event stays, deprecated, so `buf breaking` holds.
- Card: `CardFace.artist` (D-6), `partner_text`, and `max_copies_override` (F-22).
- Collection: `language`, `set_name`, `unresolved_by_reason` (M-3), and two new `UnresolvedReason` values.
- Health: `card_snapshot` and its age.

`buf breaking` now runs in CI against `main`. Gate: `buf lint`, generated code committed, every service builds.
> *In plain English:* the shared contract gains every field the next four steps need. One change now, so each later step touches only code.

**PR-2: Card database from Scryfall bulk.** ✅ merged 2026-08-24 (#5). The 200 tricky names resolve 200/200. The committed fixture covers split cards, DFCs, face names, apostrophes, and Aether spellings. End-to-end verified on the local stack. The worker downloads the three bulk files through fake-gcs and writes the completion marker (F-20). The API loads the index (about 34,000 cards) and answers Lookup and Search. Derivation facts learned from the data: the "Choose a Background" keyword has a lowercase b, and "Doctor's companion" sits on the companion card, not on the Doctor. Both are pinned by tests. This PR found and fixed F-19 and F-20.
A worker job downloads `oracle_cards`, `default_cards`, and `oracle_tags` daily (F-3: bulk only). It writes a versioned snapshot to GCS and an in-memory index in the `cards` service. The index holds the name, Oracle ID, printing ID, legalities, color identity, keywords, type line, MV, produced mana, Oracle tags, `game_changer`, `edhrec_rank`, and the image URIs per face. Faces are normalized (F-9). The `oracle_tags` file loads into a tag tree. Rulings load on demand.

A `CardService.Lookup` by exact name, by Scryfall ID, and by Oracle ID. A `CardService.Search` with structured filters (colors, types, keywords, tags, format-legal). Gate: 100% of a fixed list of 200 tricky names resolve (split, DFC, "Aether" spelling, commas, apostrophes). Snapshot age is exposed as a metric.
> *In plain English:* every night we download the whole card list, keep a copy, and load it into memory. Anyone can ask "which green cards with lifelink are legal in Pioneer?" and get a fast exact answer with no AI involved.

**PR-3: Legality freshness and announcement-day fast path (F-1).** ✅ merged 2026-08-24 (#6). The calendar is `announcement_dates.json`, embedded, with a verification date (next date: 2026-10-12). The worker checks hourly, and every 15 minutes from an announcement date until a snapshot with a legality change lands (F-23, D-47). On that snapshot, the worker logs `legality_lag` with the hours (M-2). The previous version comes from the store, not from process memory, so a restart keeps the metric. In production the worker is a Cloud Run job under Cloud Scheduler (D-48). The UI shows "Card data as of" from `/healthz`. The real M-2 number arrives with the 2026-10-12 announcement.
The worker checks the Scryfall bulk `updated_at` every hour. On and after a B&R announcement day (a committed calendar, next 2026-10-12), it checks every 15 minutes until a snapshot from that day or later lands. Every deck response carries `legality_as_of` (the snapshot date). The UI shows it. Gate: M-2 shows the lag between an announcement and the snapshot that reflects it.
> *In plain English:* ban announcements come on known dates. On those days we check more often. Every deck says which day's rules it was checked against, so the user knows.

**PR-4: ManaBox import (F-2, F-12).** ✅ merged 2026-08-24 (#7). Audit note: the gate count included one token row (F-21). The fix reports such rows as `NOT_PLAYABLE`. The owner's real export (2,548 rows, 4,317 cards, 18 columns) is the committed gate fixture at `go/internal/collections/testdata/`. End-to-end through the API against the full snapshot: 2,548 of 2,548 rows resolve, with zero unresolved. An identical re-upload updates the same document. Get and List work. Storage per D-16: one Firestore document with gzip entry and count payloads. Auth debt: a debug user id stands in until PR-11. M-3 rides the `ImportReport` counts until the analytics phase adds events.
CSV parser driven by the header row, not by column position. Required: `Scryfall ID`, or `Set code` plus `Collector number`, or `Name` plus `Set name`. Optional: `Quantity`, `Foil`, `Condition`, `Language`, binder name. Unknown columns are ignored and logged once. Rows that do not resolve are returned to the user as a list, not dropped silently.

The result is a `Collection` with counts per Oracle ID and per printing. Also accepts the Arena text format (`4 Lightning Bolt (STA) 42`). Storage: the full collection is stored (D-16). One document per collection holds a compressed entry array (printing id, quantity, finish, condition, language). 

A per-Oracle-ID count map sits beside it for fast ownership checks. A content hash of the upload detects an identical re-upload. Non-English rows are reported to the user and skipped (D-23). Gate: a fixture set of real exports (owner-provided, anonymized) imports with zero silent drops. M-3 counts unresolved rows.
> *In plain English:* upload the file ManaBox gives you. We match every line to a real card and count how many you own. We show you the lines we could not match. We do not hide them.

**PR-5: Rules engine.** ✅ merged 2026-08-24 (#8). Audit 2026-08-24: five eligibility gaps fixed (F-22). The companion is now checked for legality, color identity, and singleton. `banned_as_companion` is Commander-only. Ownership aggregates per Oracle id. The golden gate is 41 good and 53 bad decks, with a test that enforces at least 30 of each. `DeckService.Validate` accepts `pool_rule` and `collection_id`, reads the owned counts from the stored collection, and returns `legality_as_of`. A pure library in `internal/rules` with embedded, dated data files: `formats.json`, `brackets.json`, and `companion_bans.json` (the F-18 list Scryfall can not express). Checks: size, copies (basics and any-count exempt, restricted capped at 1), legality, commander eligibility, and all five partner mechanics. Also: color identity, Game Changers per bracket, companion (Lutri blocked as companion, legal in the 99), ownership per pool mode (D-37), and land-count and curve advisories. The golden gate ran 27 good and 31 bad decks at merge. The audit padded it to 41 and 53. `DeckService.Validate` is wired and smoke-tested end to end. Still open from F-11: the bracket prose rules (mass land denial, extra turns, combos) emit an info finding, not a check. Fixture lesson: Scryfall Oracle data contains token objects that share a real card's name, and the fixture builder now prefers real layouts.
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

**PR-6: Candidate-list builder.** ✅ merged 2026-08-24 (#11). The human gate failed on run 1 (12 of 20) and held on run 2 (20 of 20, bar 18). Both documents stay: `docs/reference/pr6-candidate-review.md` is run 1, and `pr6-candidate-review-run2.md` is run 2 (D-65). `internal/candidates` filters the index by legality, color identity, and the commander, then scores each card. Signals come in two kinds (D-62). Payoffs reward the theme: a payoff tag such as `lifegain-matters` (1.5) or a payoff needle such as "whenever you gain life" (1.2). Enablers do the thing: a tag such as `lifegain` (1.0), a subtype (0.8), a keyword such as Lifelink (0.5), a text needle (0.4). A kind counts once, so overlapping tags do not stack. EDHREC rank adds 0.3. A staple with no theme signal keeps half its score, so theme leads. `themes.json` maps 55 theme words to Tagger slugs, payoff and enabler apart. Run 1 exposed 16 slugs that Tagger does not have, over 11 rows, which the matcher dropped without a message. `make themes-check` now fails on an unknown slug. Run 1 also showed that a parent tag carries its children. `death-trigger`, `anthem`, `flicker`, and `counters-matter` each pulled in the wrong half of a theme. Payoffs are narrow from run 2 on. An unknown word falls back to a generic rule. Roles come from the tags first (`ramp`, `draw`, `removal`, `sweeper`, `counterspell`, `protection`, `alternate-win-condition`), then from the type line and text. Brackets 1 and 2 drop Game Changers from the list. Pool modes per D-37: any-card returns about 300 by role. Owned-first returns the owned cards plus up to 50 upgrades. An upgrade must beat the weakest owned card of the same role. Owned-only returns the owned cards. Basic lands are not candidates: the generator adds them. `cmd/candidates-review` writes the gate document from a local snapshot and a ManaBox export. Two helpers serve PR-7: `ThemeColors` names the colors a theme is strongest in, and `Commanders` names commander-eligible candidates. PR-7 puts both into its questions, so the agent states a fact instead of asking the user for it. The owner scores `docs/reference/pr6-candidate-review.md` (20 prompts, 10 with the owner's collection). Meta input (PR-14) has a hook and no data. `Stats.ThinTheme` marks an owned mode with under 30 on-theme owned cards. PR-7 asks the pool-mode question again on that flag (D-63).

Given a format, colors, a theme, a power level, the pool mode, and the collection (optional, D-37), build a ranked candidate list from the engine. 

Owned-first: owned candidates plus a bounded unowned-upgrade list. Owned-only: owned candidates alone. Any-card: the whole legal pool, ranked by theme fit and `edhrec_rank`, with meta input at competitive power (PR-14). Signals: Oracle tags (theme), keywords and type lines, `edhrec_rank` (popularity), legality, ownership. Output in owned-first mode: about 150 to 300 owned candidates by role, plus about 50 unowned upgrade candidates. Output in any-card mode: about 300 candidates by role from the full pool. 

This list, not the whole database, is what the model sees. Gate: for 20 theme prompts, a human confirms the top 40 candidates are on-theme in at least 18. Ten of the 20 run with no collection.
> *In plain English:* before we ask the AI to build, the code shortlists the cards that fit: your cards, the right colors, on theme, legal. The AI picks from that list. It can not pick a card that is not there.

**PR-7: Question workflow.** 🔧 in progress 2026-08-24. `internal/questions` holds the whole turn loop. The service and the session store are not built yet.
The turn-based core. A `Session` holds filled slots (format, commander, power, colors, theme, pool rule, budget, house rules, locked cards). Each turn: a small model classifies the prompt and fills slots it can. The code decides which slots are still empty and picks up to three questions from the catalog (`mtg-corpus` skill, section 11). The model phrases them.

The user answers in free text. The small model maps answers to slots.

Slots are stored, summarized, and carried to the next turn, as connector-syncer's schema agent does. A slot stays open through the question phase, and a later answer replaces an earlier one. Every slot freezes when a build run starts (D-68). The card-pool question waits for the format, the colors, and the theme, because PR-6 needs those three before it can count on-theme owned cards (D-67). The catalog is the first source of questions (D-25). A **gap score** decides when the catalog is not enough. It is the best catalog match between the empty slot and the user's words, from a small classifier. Below a threshold (D-27, set by M-5 with the OQ-19 rubric), the model may propose a question of its own. It gives a reason and the gap score with it. 

Every invented question is logged with its slot and outcome. M-4 reports how often this happens. Repeated invented questions become catalog candidates (PR-15). "Anything goes" and similar phrases route to the house-rules question (D-3). The pool-mode slot: with a library, the agent asks or defaults to owned-first. Without one, it defaults to any-card and does not ask (D-37).

Gate: 30 scripted conversations reach a complete slot set in at most four turns, with no repeated question. At least 25 of the 30 use catalog questions only. The gap-score threshold is set by M-5, not by this PR.

Built on 2026-08-24, in two halves. The deterministic half holds the catalog. `catalog.json` carries the rows of corpus section 11 as data. The planner picks the questions for one turn by ask order. It asks at most three, one per proto slot, never a repeat, and nothing at all when the run is frozen (D-68). A row carries a `slot` and a `key`. The slot is the proto field the answer informs, and the key is the row's own state. A refinement question such as table tolerance therefore survives a filled power slot. The model half: three calls per turn (D-69). `classify` fills slots from free text. A second `classify` call scores the catalog fit and may offer a replacement. `ask` then phrases what the agent chose. The agent decides, never the model: a replacement counts only when the fit is under 0.35.

Two guards sit between the model and the user. A placeholder is resolved before any model sees a row, and a clause with no value is dropped. A placeholder-free fallback stands in when the first sentence does not survive. A phrasing goes back to the resolved catalog text when it comes back wrong. The faults are a brace, two question marks, none at all, or a length far over the row.

Eight dogfood conversations ran on 2026-08-24, before any code. None of the eight was catalog-only, and only 5 of 26 catalog questions survived without a rewrite. One cause gave three of the invented questions: the catalog offered commander suggestions and held no question to close the slot. Section 11 of the corpus went from 11 rows to 26 from those runs. It gained 11 new rows, two rows split in two, an ask order, and a word-routing rule.

One live conversation ran on 2026-08-24 with the owner's approval, to measure a turn. Four calls, 1,875 input and 350 output tokens, $0.000795, both roles on `gpt-5.6-luna`. It found four defects that every offline test had missed. Two were mine and are fixed. The agent trusted the classifier's list of closed slots, which ended a session with three slots empty. The agent also shipped raw placeholders to the model. The model answered by turning the agent's own statement into a second question for the user. Two more came out of the fixes. A surviving trailing clause read as a dangling question. `Ready` also called a session complete while its questions were still unanswered. The no-repeat rule empties the plan as soon as a question goes out.

A second live run on 2026-08-24 confirmed every fix. It also named three items that come before the rest of PR-7. First, the locked-cards row asks the user to keep or cut a list they never gave. It fires whenever a card is named, the commander included. The scorer rated it 0.15 and the model replaced it, which is the first real M-5 row. Second, the commander row scored 0.05, because it asks whether the user has a commander and names three in the same breath. Its follow-up row can not fire at all: `Context.Suggested` is declared and read, and nothing ever sets it. Third, PR-7 passes no cache key, which costs one line when the session id exists.

The three items are fixed. The locked-cards row now fires only for a named card that is not the commander, and its text no longer presumes a list (D-70). The commander row is split. The base row asks whether the user has a commander, and the pick row carries the three names. The classify schema gained `facts.wants_suggestion`, which sets the fact the pick row needs (D-71). A name the user gives as the commander closes all three commander rows. Every model call carries the session id as the provider cache key (D-72). On the owner's call, the pick row repeats after a "none" answer, and each round names three commanders the agent did not offer before (D-73).

`AgentService.Chat` and `GetSession` are built. `internal/sessions` stores one conversation as two Firestore documents in one transaction (D-74). The first is the proto session, which `GetSession` returns. The second is a private state document with the asked rows, the user's words, the card names, and the planner triggers. Chat streams the session id, one event per question, the slots, and the usage total. A model failure ends the turn with a failure event, and the slots that the turn already filled are stored.

The M-4 report is data, not a log line. Each question leaves a record with its row, slot, source, gap score, threshold, and the fact that its key closed. `Coverage` sums the records over one session or over many. `cmd/questions-gate` runs the 30 gate conversations against the real providers and writes the gate document.

The service wiring found four more defects, each one a row that could never fire. First, the classifier was never offered the key of the question it had just asked. An advisory key could then never close, and the session was never ready. Second, a commander the user named left the color slot open, which blocked the card-pool question forever (D-67). The commander now fills that slot, because its color identity is the deck's color identity. Third, `Context.AfterBuild` had no source, so the variance row was dead. The service reads it from the stored deck ids. Fourth, `ThinTheme` and `CommanderNotOwned` had no source. The service answers both from PR-6 and from the collection.

Gate status: the offline half holds. Thirty scripted conversations reach a complete slot set in at most four turns, with no repeated question. Every catalog row fires in at least one of them.

Run 1 of the live half failed on 2026-08-25 (`docs/reference/pr7-question-gate.md`). It scored 21 of 30 catalog-only, and the bar is 25. The run cost $0.0439 over 212 calls and 365 seconds. It asked 108 questions, and the model replaced 11 of them. Two replacements filled a slot other than the one they were given, which is the `right_slot` field of the D-66 rubric. Two more lost the owned-only pool mode, so they were worse than the row they replaced (D-37).

The run exposed five defects. The worst one is D-82: the resolver asked the PR-6 hint source for every row, and it cached the answer under the theme alone. A commander list built before the user named their colors therefore survived the whole conversation. Conversation 22 asked for a blue-red deck and got Lotho, Corrupt Shirriff (white and black), Peregrin Took (green), and Massacre Girl, Known Killer (black). None of the three is legal in that deck. Conversation 24 proves the cause. Turn 3 named three commanders outside the colors, and turn 4 named three inside them. A longer skip list missed the stale cache entry.

The other four are smaller. "Casual" alone triggered the house-rules row (D-78). The competitive theme row never closed its key (D-79). The pick row named three other commanders on every turn (D-80). Two rows fired before their context existed (D-81). The owner approved all five fixes on 2026-08-25.

The gap-score threshold stays at 0.35. The 11 replacements scored 0.02, 0.05 four times, 0.10, 0.18, 0.20 twice, 0.22, and 0.30. A threshold near 0.15 would block five of them and probably pass the gate. M-5 sets that number from the D-66 rubric, and a change made to pass a gate would make the number meaningless.

Run 2 scored 23 of 30 (`docs/reference/pr7-question-gate-run2.md`). It cost $0.0411 over 192 calls, and it asked 91 questions with 7 replacements. Two fixes hold under live conditions. Conversation 24 named three commanders inside the color identity, against three outside it in run 1. The pick row kept its three names when the user answered another question, against three new names in run 1.

Run 2 also exposed two deeper defects. The first is D-83. The classifier could retire a slot by name, and it retired power and the pool rule from "Brago blink deck from my library". That session called itself complete after one question, and the deck would have had a power level nobody chose. A slot with a typed value now closes only on that value.

The second is D-84, and it binds the gate itself. The gap score was not reproducible. One question scored 0.02 and 0.98 in two turns of one conversation, and 0.98 then 0.02 across the two runs. Six conversations improved between the runs and four regressed. The bar of 25 sits inside that noise band, so the number measured the scorer more than the catalog. The score prompt asked a question of taste. It now names four faults, and the score follows the fault count.

Five of the seven run-2 replacements were not improvements. One restated the catalog row almost word for word. One dropped the bracket definitions, and one dropped the owned-only pool mode again. One asked for a budget the user gave a turn earlier. The last replaced three named commanders with an open question, after the user had asked the agent to choose.

Run 3 scored 23 of 30 again (`docs/reference/pr7-question-gate-run3.md`), and it changed the instrument. Every fit landed on the fault scale: 0.90 for 67 questions, 0.20 for 7, and 0.05 for 12. Runs 1 and 2 scattered over nine values between 0.02 and 1.00. The score now reports a fault count, so the threshold has a meaning it did not have before. At 0.35 the agent invents on any clear fault. At 0.10 it would invent only on two faults or more, which in run 3 was 5 replacements instead of 10. M-5 still owns the number (D-27).

Run 3 also found a regression in D-83. The first rule blocked a typed slot from closing by name, on the theory that the classifier always returns a typed value. It does not. It reports the answer in the free-text list and leaves the field unknown. The format slot therefore stayed open after a user answered "Pioneer", and the agent asked for the format again. The rule is now the question, not the field type: a key closes by name only while its question is out. That still blocks every case the rule was built for, because none of those keys had a question out.

Three of the run-3 replacements hit the color row. Each one replaced a statement that read as nonsense: "the best deck under budget is strongest in white, blue, black, and green". D-79 made that phrase a theme value, and the color clause used it as a subject. D-86 drops the clause when the theme names no archetype, and when the answer holds more than two colors.

The owner deleted the acquisition row on 2026-08-25 (D-87). It asked where the user buys and by what date they need the cards. The app holds no store stock and no delivery times. Scryfall gives a price estimate rather than availability (D-17). No reader could act on the answer. The row also asked two things in one sentence, and every run replaced it. The catalog holds 25 rows.

Run 5 scored 29 of 30 and run 6 scored 30 of 30 on the catalog bar. Neither number stands on its own. The reword guard of D-88 refused 15 of the 16 replacements the model offered in run 5. The guard therefore decides the count, and its 0.6 threshold came from run-4 data rather than from the rubric. The M-5 sheet now carries refused rewords for that reason.

Run 6 also carried a second bar for the first time (D-91). It scored 30 of 30 and still failed, because five sessions called themselves complete with no power level. The old bar alone called that run perfect, which is the point: a conversation that stops asking looks the same as one that finished.

The cause was the second form of D-83. The classifier closed the format by name, the format value stayed empty, and every row that triggers on the format stopped firing. D-83 is now settled in its third form. A typed slot closes on its value alone, because a name says "answered" and never says what the answer was. Asking twice is the safe failure, and building a deck with no format is not. D-90 puts the same rule in the classify prompt.

Run 7 confirmed the D-83 fix and exposed what the fix had been hiding. It passed both bars, at 28 of 30 with no premature session, and the pass was hollow. Twenty-six of the thirty conversations ended with a slot unanswered, and twenty-five of those sat on "format (asked, no answer)". The agent asked the format, the user answered it, and the answer never landed. Fewer slots filled means fewer rows fire, fewer questions go out, and fewer chances to invent one. Both bars improved while the product got worse.

D-92 is the cause, and it sits under the whole D-83 history. The classify schema constrained `format` and `pool_rule` to an enum that held "unknown". The model answered "unknown" for a message that named the format outright, while it filled every free-text field in the same reply. It even guessed colors and a power step from "I want a Modern burn deck" and still left the format empty. Four samples per variant measured it. The enum extracted 1 of 8, an enum with an empty member 3 of 8, and a free string 7 of 8. With the field freed the same probe reads 16 of 16, including the negative case where no format is named.

That explains the earlier symptom rather than excusing it. The model reported the format through `closed_keys` because the field itself was suppressed. Every version of D-83 argued about which channel to trust, instead of asking why the field was empty. A live probe of about ninety calls, for roughly two cents, answered in minutes what four gate runs could not.

Run 8 is the first honest pass, at 26 of 30 with no premature session. It asked 135 questions, the most of any run, and the format-stuck count fell from 25 to 3.

Run 9 closes the last dead method. `State.Skip` was declared and never called, so nothing let a user hand a choice back. Run 8 measured the cost: conversation 4 answered "any colors are fine" and still ended with the color slot open. D-93 adds `declined_keys`. A decline closes any key and names no value. The generator applies the default the corpus lists, and the skipped state is what tells it to. Run 9 scores 29 of 30 with no premature session. The count of conversations holding an unanswered slot fell from 26 in run 7 to 14.

A live check caught one precision fault before the run. The first wording let "any colors are fine" decline the bracket as well, which would have skipped a slot the user never mentioned. Two probes and one prompt sentence fixed it.

Run 10 adds D-94, a PR-6 correction found through PR-7. `Commanders` read the 99-card shortlist and took the first legends it met. That list ends in `capByRole`, which emits one role bucket after another with lands first, so it carries no score order at all. A blink request answered with three Ojer modal double-faced cards, which the theme scorer rates 0.16. Eighty-five on-theme blink commanders existed, and Emiel the Blessed rated 0.56. `CommanderPool` now walks the index itself. It requires a theme signal, drops the staple-role fallback, and applies no role cap, because one card fills no role quota. The 99-card pipeline is untouched, so the PR-6 gate holds.

The same pass closes the last open PR-7 item. The count of on-theme commanders in an owned pool is the weak-commander-pool signal (D-63), so no score bar had to be invented. Run 10 asked 138 questions and closed 92, both the best of any run. A lifegain user with no library is offered Vito, Thorn of the Dusk Rose and Heliod, Sun-Crowned. The same user with the owner's library is offered the best lifegain commanders that library holds, because owned-first orders them first (D-37).

Run 11 added 22 probe conversations (D-97). Every one of the 30 gate conversations holds a cooperative user who answers what the agent asks. A probe does not. It changes its mind, contradicts itself, asks a question back, or wants something the app can not build. A probe runs beside the gate and feeds the M-5 sheet. It does not count toward the catalog-only bar, because adding conversations to a bar moves the bar.

The gate set passed run 11 at 27 of 30, with no premature session. The probes found two defects at once. A declined format left the planner with nothing to route on, so no power row could fire (D-98). A Yu-Gi-Oh request drew the answer "Which Yu-Gi-Oh format would you like?", because no catalog row could decline (D-99). They also surfaced three catalog gaps with no row at all: an out-of-scope request, a sideboard-only request, and a request for two decks at once.

The M-5 sheet now reads engine-current runs alone (D-96). A run made before a defect was fixed measures the defect. Run 3 is the clearest case. The format enum threw the user's answer away, the format row fired again, and the model invented a repair question. Scoring that item would push the D-27 threshold up on evidence about a bug that no longer exists.

Runs 12 and 13 confirm both fixes. Neither holds a premature session, across the 30 gate conversations and the 22 probes, and each passes the catalog bar at 27 of 30. The probes asked 88 and 84 questions and drew 6 replacements each, against 6 from 78 in run 11. The new scope row absorbed the request that had forced the model to improvise.

Two gaps closed before the commit. The session store had never executed. Every Firestore path in `internal/sessions` sat at zero coverage. `Put` ran only far enough to reject a session with no id. The code that persists every conversation was unverified. Seven emulator tests now drive it, and coverage went from 31.9% to 78.0% (D-101). A round trip alone was not enough. `Put` and `GetState` agree with each other whatever path they use, and a rename of the private path passed. Two tests read the literal document path and hold the D-74 layout: the private state stays out of the document `GetSession` returns.

The second gap was the arithmetic. Nothing read a scored sheet, so the threshold had to be counted by hand. `cmd/m5-report` reads it and reports the fit threshold, the reword guard, and what the catalog needs (D-102). It reads the first word of a field, so the owner's free text survives. A refused reword is left out of the fit threshold, because it never reached a user.

The scoring pays for itself before it finishes. Item 8 named a fault the six of D-66 did not hold. A replacement dropped the sentence that says what the answer is for, so the user could not tell whether to name decks, colors, or archetypes. `vague` is the seventh fault, and the fit is unaffected, because `faults` never enters the threshold calculation.

The same item exposed the reword guard. Word overlap is symmetric, so a replacement that deletes half the row scores low and passes, although it says strictly less. Item 8 scored 0.44 against a 0.60 bar. The guard now refuses a truncation as well: a replacement that borrows 0.8 or more of its words from the row and keeps under 0.6 of it. Measured against all 60 items, the rule refuses item 8 and nothing else (D-103).

Open in PR-7: the M-5 scoring, then the threshold and one confirming run.

> *In plain English:* the chat. "Build me a lifegain deck" fills in "theme: lifegain" and leaves format, power, and colors empty. The app asks those three, remembers the answers, and never asks twice. If the user says something vague, the app asks what they mean. It does not guess.

**PR-7B: Automated eval lane.** 🔧 in progress 2026-08-26, branch `pr-7b`. PR-7 proved that reading transcripts finds defects and that scoring replacements does not. The version-1 M-5 sheet could hold 60 of 793 questions, because it held only a question the model offered to replace (D-104). The batch sweep of 2026-08-26 read all 66 conversations by hand and found sixteen more defects, none of which could have reached that sheet. PR-7B makes that reading automatic.

A sixth role, `eval`, scores every question of a gate run against the rubric the owner applied by hand (D-133). It runs on `gpt-5.6-luna`, the cost tier, so a 66-conversation run costs about eleven cents. It is not the judge role: D-4 gives the judge a deck, and D-22 keeps the judge off the generator's provider. The eval role shares a model with the classify and ask roles, which the owner accepted with the risk named (D-136). Every ratio it reports is a floor.

`internal/tune` reads a gate document back and holds the accept rules. `cmd/questions-eval` writes a report a person reads and a summary a script reads. `cmd/tune-check` decides whether one iteration may be kept, and it costs nothing. `scripts/autotune.sh` is the loop, and it refuses to start without `AUTOTUNE_ALLOW_UNATTENDED=1`.

Three evals run first, in this order.

| Eval | What it measures | Cost |
|---|---|---|
| Gate run | The transcript. 100 conversations, 30 gate and 70 probe (D-145). | about $0.14 |
| Question eval | Every question, scored for whether it deserved to be asked. | about $0.10 |
| Eval calibration | The cost-tier eval against `claude-sonnet-5` on 12 conversations. | about $0.25 |

The calibration answers the one question the cost tier raises: how gently does a model score work its own model produced? It scores the same 12 conversations twice, once on the cost tier and once on `claude-sonnet-5`, and `cmd/tune-check -agree` compares the two question by question. It reports how often they agree, and how many questions each one refused. A cost-tier eval that refuses four where the stronger model refuses twelve is not measuring the agent. It reports a floor, and the real number sits above it. OQ-39 holds what the owner does with that gap.

Four counters guard the ratio, because a run that asks less scores better and serves the user worse. Gate run 7 of 2026-08-25 passed both bars with 26 of 30 conversations unanswered. The counters are the questions asked, the questions that closed a slot, the premature sessions, and the linter findings. They come from the transcript and not from the M-4 table, because the table counts the 30 gate conversations alone and the terse set of D-105 is where the hard cases live.

Every third conversation is a holdout (D-134). The eval scores it, the report never names its failures, and the loop reads its ratio. A ratio that falls on the two thirds the fixer saw, and stands still on the holdout, is a reworded test set.

The loop that consumes these evals comes second, and it starts only when the gate below holds. The owner settled its four terms on 2026-08-26: it may change the catalog inside an approved run (D-135), it may share a model with the agent it scores (D-136), it stops at a 5 percent holdout ratio (D-137), and it works on a branch of its own and pushes nothing (D-138). `docs/reference/autotune-design.md` is the authority on how it runs: the guards, the accept rules, the owner's duties run by run, and the one cost the loop can not see, which is the fixer agent's own tokens.

PR-7B is the instrument and not the repair. Its scope is the eval role, `internal/tune`, the three commands, the loop, and the guards. A change the loop makes is a change to PR-7: the catalog rows, the planner, the word rules, and the prompts all belong to the question workflow. The loop writes PR-7 corrections, and the PR-7 entry records them. This keeps one concern per PR, and it keeps the measuring device apart from the thing it measures.

Each accepted iteration is one commit on a branch of its own, with the ratio and the counters in its message, and one appended row in `docs/decisions.md`. A rejected iteration leaves nothing, so every commit on that branch passed the accept rules.

The branches work in three layers (D-142). `main` holds the merged work. A branch the owner keeps, `pr-7c`, is the container for everything the loop writes. Each night cuts `auto-tune/<stamp>` off that container, and it pushes nothing. The owner reads the night in the morning and fast-forwards the container, or deletes the night.

The container matters more than it looks. A loop that always starts from `main` gives the second night none of the first night's accepted work. Two nights then change the same catalog rows from the same starting point, and the owner merges two branches that disagree. A container makes the nights add up, and it makes one reviewable pull request out of many nights.

Gate: the three evals run. The report names every defect class the batch sweep of 2026-08-26 found by hand. The calibration reports the agreement between the two eval models, and the owner accepts that number or names a stronger model (OQ-39).

> *In plain English:* the owner spent hours scoring 32 questions by hand, and one test run asks 250. This adds a second model that reads every question and says whether it deserved to be asked. It writes a short report instead of a spreadsheet. A third of the conversations are hidden from anything that tries to fix the code, so we can tell a real gain from a reworded test. The scorer is cheap, and cheap scorers are kind. We measure how kind before we trust the number.

**PR-8: Deck generator and normalizer (F-13).**
Prompt caching pays here, not in PR-7. The generate role runs on the middle model, and its prompt carries a candidate list of about 300 cards. That list is the same across a repair retry and across PR-9's re-rolls, and a cache read costs a tenth of a fresh read. Two rules protect the lever: keep the stable text first and the session text last, and hold one cache key per session. The PR-7 measurements of 2026-08-24 read zero cached tokens, because each call sat near 470 tokens and the OpenAI cache starts above 1,024. Anthropic caching is opt-in, and it charges 1.25 times input to write. The judge role therefore pays for a cache only when it reads the same prefix more than twice. The adapter has no `cache_control` wiring today.
With all slots filled, the strong model gets four inputs. They are the rules summary for the format, the candidate list with roles, the role targets, and the plan request. It returns a structured deck (D-19). First, one summary paragraph on the deck's style and purpose. Then cards with exact names, counts, roles, and one line each. 

The normalizer exact-matches every name to the candidate list. A miss is returned to the model once as a tool error. A second miss becomes a user-visible note.

The engine validates (PR-5). 

The ownership check runs only in the owned modes. In any-card mode, ownership marks are information, never findings (D-37). A `block` finding triggers one repair turn with the findings as input. Then the deck goes to the user with the `ValidationResult` attached. Gate: on the golden prompts, 100% of returned decks pass `block` checks. Zero invented names reach the user.
> *In plain English:* the AI writes the deck from the shortlist, with a plan and a reason for each card. The code checks every name and every rule. If something is wrong, the AI gets one chance to fix it. What the user sees has already passed the referee.

**PR-9: Designed variance (F-14, D-18).**
Random variance is a feature (D-18). Variance comes from three levers, not from temperature alone. Lever 1: a seeded shuffle within each role tier of the candidate list. 

Lever 2: a "plan variant" slot (for example "lifegain aristocrats" versus "lifegain go-wide"). Lever 3: a "keep these, change the rest" re-roll. The seed is stored with the deck so a build can be reproduced on request. Identical output on identical input is not a requirement (D-18). Gate: two builds of the same prompt differ in at least 30% of nonland cards and both pass validation.
> *In plain English:* ask twice, get two different but sensible decks. Each deck remembers the dice roll that made it, so you can get the same deck back.

**PR-10: LLM role layer (D-1).** ✅ merged 2026-08-24 (#9). The live smoke passed on both adapters. Classify ran on `gpt-5.6-luna` (69 in, 21 out, 2.7 s), and judge on `claude-sonnet-5` (316 in, 20 out, 3.0 s). The pair cost $0.0013 at list price.

`internal/llm` is the one door. `roles.json` is the frozen map, dated, with an owner note per default (D-38, D-39). The baseline: `classify` and `ask` on `gpt-5.6-luna`, `generate` and `repair` on `gpt-5.6-terra`, `judge` on `claude-sonnet-5` at effort medium with thinking on (D-45). Config validation refuses a judge on the generator's provider (D-22).

Adapters: OpenAI Responses API and Anthropic Messages API through the official Go SDKs (D-40), plus the fixture `Fake` and a scripted fake for tests. Both adapters send a strict JSON Schema. The client validates the output again locally. One `Budget` per logical call: four attempts, three minutes. Truncation retries once at a higher cap, bounded to min(65,536, max(8 x cap, 8,192)). Transient errors back off with a 30-second cap and jitter.

Refusal, schema, and terminal errors return at once.

`Accumulator` reports tokens and USD per session from a dated `prices.json`. It reports null when a call gave no usage or an unpriced model.

Thinking tokens count on both providers, and Anthropic cache writes are priced at 1.25x (F-24). `make llm-defaults-check` warns in CI when `roles.json` or `prices.json` changes. Keys live in `.env` (D-41). Keys are required by default. `LLM_REQUIRE_KEYS=0` (set by `make dev`) lets the fixture fake stand in (D-51).

Gate: unit tests over the fakes and `httptest` adapters pass in CI. The live smoke (`make test-smoke`) proves both adapters against the real APIs.

> *In plain English:* the AI plug. Every place that calls an AI calls it through one door with a named job. Swap the vendor in one file. Count every token.

**M-1: Token and cost accounting per session.** Lands with PR-10. Every cost claim in this doc is an estimate until then.

**M-5: Manual scoring lane for invented questions (D-27, F-17).**
The owner uses the product on a fixed set of prompts. For each model-invented question, a review page shows four things. The question, the gap score, the slot, and the top three catalog questions that were possible instead. The owner then scores six fields (D-66, closes OQ-19):

- `catalog_enough`: yes, no, or unsure. Was one of the three catalog questions good enough?
- `invented_better`: worse, same, or better than the best catalog question.
- `right_slot`: yes or no. Did the question target the correct empty slot?
- `filled_slot`: yes, partly, or no. The session fills this field. The owner only corrects it.
- `faults`: mandatory on every row that is not clean. One or more of duplicate, two questions in one, jargon, assumes an answer, unanswerable, out of scope.
- `catalog_action`: none, add, or reword an existing entry. This feeds D-25 and PR-15.

Scales hold three points, because a five-point scale drifts between sessions and makes the rows hard to compare. An invention is warranted when `catalog_enough` is no and `invented_better` is not worse. An `unsure` row is reported and left out of the fit. 

Scores go to the eval store with the prompt version, the model id, and the rubric version. The threshold is the gap score that best separates a warranted invention from an unwarranted one. A precision floor binds the choice: of the inventions the threshold allows, at least 80% must be warranted (D-66). Every tenth item repeats an earlier one, which measures self-consistency across sessions. A change to the rubric invalidates the fit. The threshold is re-checked after each catalog change (D-28: the owner approves changes). Gate: at least 50 scored invented questions before the threshold is set.
> *In plain English:* the app sometimes has to make up a question. The owner will use the app, see each made-up question next to the fixed questions it could have used, and grade it on six fixed fields. Those grades decide how eager the app is to make up questions. The rule is strict: at least four of every five made-up questions must be ones the fixed list could not cover.

**M-4: Catalog coverage metric (F-17).** Per session: catalog questions asked, invented questions asked, gap scores, and whether the invented question filled its slot. A weekly report lists invented questions by frequency. This is the input for catalog changes (D-25).
> *In plain English:* we count how often the app had to invent a question. If the same invented question appears again and again, it belongs in the fixed list.

### Phase 3 - UI (gated on PR-8)

**PR-11: Web app shell.**
React 19, Vite, TypeScript, Tailwind, the wallabee-ui patterns (TanStack Query, Zustand, lint-enforced import boundaries). Firebase Auth (D-11) with the emulator in local mode. Generated Connect client in `packages/api-client`. Gate: sign-in, then either upload a collection and see the count, or skip the upload and still reach the chat (D-37).
> *In plain English:* the website skeleton: log in, upload your binder, see how many cards we recognized.

**PR-12: Chat and deck view.**
A streaming chat thread over the `Chat` RPC. The deck view groups cards by role. It shows card art from Scryfall image URIs with artist and copyright (D-6, guardrail 7). It shows both faces for DFCs (F-9). 

It marks owned versus to-buy when a collection is attached. It shows the pool-mode toggle ("use only cards in my library") with the session's mode (D-37). In any-card mode, the buy list can be the whole deck.

It shows the mana curve, the color sources, the `ValidationResult` findings, and `legality_as_of`. Hover or tap shows Oracle text. Gate: a11y checks pass. Every image has attribution in the DOM.
> *In plain English:* the main screen. The conversation on one side, the deck on the other with real card pictures, grouped by what each card does, with your own cards marked.

**PR-13: Export and share.**
Export as ManaBox text first (D-15). Other formats later. A buy list with Scryfall purchase links. Gate: a round trip ManaBox export to import loses nothing.
> *In plain English:* get the deck out of the app and into ManaBox or Arena with one click, plus a shopping list.

### Phase 4 - Meta and quality (gated on Phase 3)

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

**I-2: Price-aware buy list** (D-17, F-16). USD. Each card carries the lowest Scryfall NM market price across legal printings and finishes. It is a 7-day rolling average, and an outlier day is rejected. An outlier is a day more than 2x the 7-day median (D-26). Digital-only and gold-bordered printings excluded. The UI labels it "NM market estimate" with the price date.

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
5. PR-2 card database.
6. PR-3 legality freshness, with M-2.
7. PR-4 ManaBox import. Then M-3.
8. PR-5 rules engine.
9. **GATE.** Phase 2 starts only when the golden decks pass PR-5. Held 2026-08-24.
10. PR-10 LLM role layer, with M-1.
11. PR-1b contract amendment (audit branch, D-46).
12. PR-6 candidates.
13. PR-7 questions.
14. PR-7B automated eval lane. It runs beside PR-8 once its three evals hold.
15. PR-8 generator.
16. PR-9 variance.
17. **GATE.** Phase 3 starts only when PR-8's gate holds on the golden prompts.
18. PR-11, PR-12, PR-13.
19. PR-15 eval harness (can start after step 15, in parallel with the UI, if a second owner exists). M-5 manual scoring runs on the first UI build (after PR-12).
20. PR-14 meta, then I-1, I-2, I-3 on evidence.
21. Phase 5 stays parked.

## 9. Open questions

See `docs/open-questions.md` for the full list with "ask when" dates. The ones that gate a phase:

1. **OQ-19 scoring rubric** answered 2026-08-24 (D-66). M-5 is no longer gated on it.
2. **OQ-18 rerun depth rule** gates I-1.
3. **OQ-20 public anonymized ManaBox exports** widen the PR-4 fixture set when found (D-43).
4. PR-9's 30% variance number is a placeholder until PR-15 measures it.
