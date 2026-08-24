# Decisions

Every owner decision, with date. This file is the source of truth. The design doc cites it by D-number. Do not re-ask a question that has a row here.

| # | Date | Question | Decision |
|---|---|---|---|
| D-1 | 2026-08-23 | LLM provider | Provider-agnostic role-to-model layer. The provider is a config choice. Model: connector-syncer `llm_roles.py`. |
| D-2 | 2026-08-23 | Card pool | Build from the user's owned cards first. Suggest upgrades and acquisitions separately. |
| D-3 | 2026-08-23 | "60-card anything goes" | Not a fixed format. The agent asks the user what it means. Flexibility is the principle. The owner personally means "any card, no ban list". |
| D-4 | 2026-08-23 | Audience and deployment | Multi-user web app on GCP. Mirror the Wallabee stack: Cloud Run, Firestore, GCS. |
| D-5 | 2026-08-23 | Meta data sources | Allowed: official MTGO decklists, MTGGoldfish / MTGTop8 / Aetherhub, EDHREC for Commander. The aggregator sites need a terms-of-use check before any scraper ships. |
| D-6 | 2026-08-23 | Card images | Hotlink Scryfall image URIs. Show artist and copyright. Do not crop, skew, or watermark. |
| D-7 | 2026-08-23 | RPC style | Connect-RPC with buf. One `.proto` contract. Generated Go and TypeScript clients. Server streaming for chat. |
| D-8 | 2026-08-23 | Power scale | Commander: official brackets 1-5. 60-card: three steps, casual / FNM / tournament-meta. |
| D-9 | 2026-08-23 | Local test environment | The full stack, UI included, must run locally. Firestore emulator confirmed feasible. |
| D-10 | 2026-08-23 | Docker | The owner will install Docker. The install is a roadmap item. |
| D-11 | 2026-08-23 | Auth | Firebase Auth, with the Auth emulator in local mode. |
| D-12 | 2026-08-23 | Output language | All docs, skills, and agent files in ASD-STE100. |
| D-13 | 2026-08-23 | Write scope | This session writes only in `docs/` and `.claude/` of this repo. All other repos are read-only. |
| D-14 | 2026-08-23 | Question cadence | Ask questions as they come up, in small batches. Never batch them for the end. |
