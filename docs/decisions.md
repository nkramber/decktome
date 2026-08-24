# Decisions

Every owner decision, with date. This file is the source of truth. The design doc cites it by D-number. Do not re-ask a question that has a row here. When a later decision changes an earlier one, keep both rows and mark the old one "superseded".

| # | Date | Question | Decision |
|---|---|---|---|
| D-1 | 2026-08-23 | LLM provider | Provider-agnostic role-to-model layer. The provider is a config choice. Model: connector-syncer `llm_roles.py`. |
| D-2 | 2026-08-23 | Card pool | Build from the user's owned cards first. Suggest upgrades and acquisitions separately. |
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
| D-13 | 2026-08-23 | Write scope | This session writes only in `docs/` and `.claude/` of this repo. All other repos are read-only. |
| D-14 | 2026-08-23 | Question cadence | Ask questions as they come up, in small batches. Never batch them for the end. Never hesitate to ask, or to call out an inconsistency. |
| D-15 | 2026-08-23 | Deck export targets (OQ-1) | ManaBox text first. Other formats later. |
| D-16 | 2026-08-23 | Collection storage (OQ-2) | Store the full collection (name, printing, quantity, finish, condition, rarity, set, collector number). A hash can not recover any of this. The hash is kept only to detect an identical re-upload. |
| D-17 | 2026-08-23 | Prices (OQ-3) | USD. Source: Scryfall daily near-mint market estimate, 7-day rolling average, outliers rejected. Show the lowest price across all legal printings and finishes. Label it "NM market estimate". Scryfall has no lightly-played tier. A condition-tiered source (TCGplayer API) is a later option, not planned. |
| D-18 | 2026-08-23 | Determinism and variance (OQ-4) | First answer: identical collection plus identical prompt should ideally give the same deck, and divergence through follow-up questions is acceptable. **Second answer, same day, supersedes it: keep random variance. Identical-input determinism is not a requirement.** The seed is still stored with each deck so a build can be reproduced on request. |
| D-19 | 2026-08-23 | Card explanations (OQ-5) | Yes. One summary paragraph on the deck's style and purpose. One line per card. |
| D-20 | 2026-08-23 | Play simulator (OQ-6) | Later, not at launch. Stays in Phase 5. |
| D-21 | 2026-08-23 | Day-one provider (OQ-7) | OpenAI on day one, behind the provider-agnostic role layer. |
| D-22 | 2026-08-23 | Judge model (OQ-8) | A different provider from the generator. The exact provider is chosen at implementation time. |
| D-23 | 2026-08-23 | Languages (OQ-9) | English only. Non-English rows in an import are reported to the user, not silently dropped. |
| D-24 | 2026-08-23 | GCP projects (OQ-11) | `mtg-dev` and `mtg-prod`. No domain name and no hosting yet. Local testing first. |
| D-25 | 2026-08-23 | Question source (OQ-12) | Both. Draw from the fixed catalog whenever possible. Add a gap score that detects when no catalog question fits. Allow model-invented questions when the gap score says one is needed. The catalog is a living artifact: extend, change, or reduce it from user feedback and evals. |
| D-26 | 2026-08-23 | Price outlier rule (OQ-13) | Yes: drop a day whose price is more than 2x the 7-day median. |
| D-27 | 2026-08-23 | Gap-score threshold (OQ-14) | Needs evals, gated on manual testing by the owner. The owner uses the product and scores each model-invented question against the catalog questions that could have been asked instead. The threshold is set from those scores. This is a complex, first-class part of the deliverable. |
| D-28 | 2026-08-23 | Catalog change approval (OQ-15) | The owner approves catalog changes to start. |
| D-29 | 2026-08-23 | Stale decks on a rule change (OQ-16) | Show a banner on a stored deck that now contains a card illegal in its format. Give a "rerun" button. The rerun rebuilds according to the nature of the change: which cards were banned and how deeply they affect the deck. |
| D-30 | 2026-08-23 | Sample ManaBox exports (OQ-17) | The owner will provide them later. |
