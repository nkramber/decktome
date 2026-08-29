# Session hand-off

`CLAUDE.md` is the entry point. It sends you here. Then read `docs/decisions.md` from D-297, `docs/owner-questions.md`, and `docs/open-questions.md`. The session narratives of 2026-08-23 to 2026-08-28 sit in `docs/reference/session-log-2026-08-23-to-26.md` and `docs/reference/session-log-2026-08-27-to-28.md`. The audit of 2026-08-29 is `docs/audit-2026-08-29.md`.

CAUTION: the web tests need Node 22.23.2 (`.nvmrc`). Under Node 20 every test file fails at start with `ERR_REQUIRE_ESM` from jsdom 30. On the owner's machine `~/.nvm/versions/node/v22.23.2/bin` on the PATH fixes it.

## Where things stand (2026-08-29)

- `main` is at `7924658`, PR-12B merged (#41, 2026-08-29). Merged: PR-0a to PR-8, PR-7B, PR-10, PR-11 (#38), PR-12 (#40), PR-12B (#41).
- The quality audit of 2026-08-29 is merged (#43, D-302 to D-306, `docs/audit-2026-08-29.md`).
- PR-13 is merged (#45, D-307 to D-309). The owner did not run the export in the browser yet.
- Phase 3B, the product UI, has a plan and no code yet (D-310 to D-319). The roadmap holds PR-16 to PR-23, and `docs/reference/ui-phase-plan-2026-08-29.md` holds the detail. The first slice is PR-16, the design system and the shell (D-317). The owner asked for the roadmap first and no code (D-319).
- The tree is green on the branch: Go build, vet, `-race` tests, golangci-lint, `buf breaking`, web lint, typecheck, 118 web tests, and the web build. `make lint` runs the extended STE check and reports zero findings.
- The revise gate held on run 2 (D-296). The paid gates did not run on 2026-08-29.
- PR-9 is out of the MVP (D-256). Phase 3B comes before Phase 4 (D-316).
- Pull request #42 (Dependabot, anthropic-sdk-go 1.66.0 to 1.67.0) is open and waits for the owner.

## The numbers, and why none of them compare with `main` now

The audit changed two prompts (D-302). The generate prompt is at version 10 and the classify prompt at version 14. The shortlist omits the commander, the house format has a sideboard sentence, and the dead `two_plans` and `owned_mode` facts left the classify schema. So the baselines below measure the code before the audit, and the next run re-baselines (D-66).

| Measure | Baseline | Where measured |
|---|---|---|
| Question gate | PASS 25 of 27 counted, 2 invented, 0 premature, 0 lint, $0.16 | run 27, 2026-08-28, ask prompt 13, catalog of D-290 and D-294 |
| Question eval | 19 bad of 382, holdout 8 of 111 (7.2 percent) | eval of run 25, 5 conversations unjudged |
| Deck gate | PASS 18 of 18, 2 repairs, $1.09 | run 8, 2026-08-28, generate prompt 9 |
| Revise gate | PASS 6 of 6, $0.29 | run 2, 2026-08-28 (D-296) |
| Loop | off since 2026-08-26 | seven starts, nothing kept |

CAUTION: `tune-check` paired zero questions between run 24 and run 25, because the catalog and the prompt changed. The paired guard says nothing across that line, and the whole-run margins carry the verdict. The eval leaves a conversation unjudged when the judge returns fewer verdicts than questions. So 382 is the honest count, not a drop from 435.

## The audit of 2026-08-29, what changed

`docs/audit-2026-08-29.md` holds the report. The points a session needs first:

- A turn during a build gets `CodeAborted` with "a build is in progress". The build and its store writes run detached from the client, so a disconnect keeps the deck (D-303). The web app warns before it leaves a page mid-build.
- Every session, deck, and collection id passes `gzstore.ValidID`. An id with a slash is `CodeInvalidArgument`.
- `go/internal/gzstore` holds the gzip and JSON helpers of the three Firestore repos, with one inflate limit.
- The buy cost and the deck cost count the commanders and sum copies per Oracle id. A locked, kept, or commander card above the mana cap stays in the pool (`Revision.Exempt`).
- The house format offers paper cards only (D-306).
- `deck-gate` and `revise-gate` exit 1 on FAIL. A nil cost prints `unpriced` in every gate. `revise-gate` has `-only` and `-dry`. `questions-eval` and `candidates-review` refuse an existing output (D-65).
- The Makefile runs bash with `pipefail`. `candidates-review` guards on a filled score cell.
- `autotune.sh` builds `tune-check` after the branch switch, never switches in a dry run, and restores the start branch on exit.
- The go job checks out with `fetch-depth: 0`, so `make llm-defaults-check` finds a merge base. A weekly `vuln` job runs govulncheck alone on Monday 06:00 UTC (D-305).
- The STE checker flags passive voice, modals and perfect tenses, -ing forms, and the 20-word step limit (D-304). Dated records are exempt: `docs/reference/pr[0-9]*`, the session logs, `docs/audit-*`, and testdata.
- Every code comment states a rule and cites a decision id. No comment carries a date, a session id, a run number, or "the owner".

## PR-13, what it holds (2026-08-29)

- `DeckService.ExportDeck(deck_id, format)` returns the text and a file name. `EXPORT_FORMAT_ARENA_TEXT` is the ManaBox shape, and `EXPORT_FORMAT_BUY_LIST_TEXT` is one "count name" line per card to buy (D-309).
- `go/internal/export`: `ArenaText`, `BuyList`, `BuyListText`, `FileName`, `Render`. The Arena line names the owned printing when the card is owned, else the default paper printing (D-307). The line carries the full card name, and the index resolves it without ambiguity. `TestArenaTextRoundTrip` is the gate.
- The buy list is the shortfall of the commander, the main deck, and the sideboard, summed per Oracle id, and the upgrades apart (D-308). A commander with no entry in the card list counts as one card to buy.
- `web/apps/web/src/features/export`: `export-panel.tsx` (four buttons, the status line, the buy list with a Scryfall link per card) and `buy-list.ts` (the same rules as the Go package, for the screen). The deck view mounts the panel under its header. `export` is a leaf feature, and `deck` imports it.
- The web app reads the text from the API, so the copy and the file match what the API renders.

## Next steps, in order

1. The owner reads the Phase 3B roadmap and the plan, and says when PR-16 starts. Work on a branch named `pr-16`.
2. PR-16: the design system and the shell, with no new feature. Then PR-17 to PR-23 in order, one gate each.
3. The owner runs the question gate and the deck gate to re-baseline (D-302), in parallel. Ask before each run. Write each to a new `GATE_OUT` file (D-65). Record the numbers here and in the roadmap.
4. Before PR-22, ask OQ-45 (the allowlist store) and OQ-46 (the spend cap).
5. After Phase 3B: PR-15, then PR-14.

A ruleset that requires the `verify` check on `main` is not possible. The repo is private on the free plan, and the rulesets API answers 403 (checked 2026-08-28). The owner reads the checks before a merge.

Seven owner rows wait in `docs/owner-questions.md`: OQ-23, OQ-28 to OQ-31, OQ-37, and OQ-39. OQ-20, OQ-44, OQ-45, and OQ-46 wait in `docs/open-questions.md`.

## Facts that expire

- Comprehensive Rules: the current file is 2026-08-19 (D-272). The rule citations in the corpus match the 2026-08-07 text.
- Ban-list snapshot: 2026-08-24. Next announcement 2026-10-12, in `internal/cards/announcement_dates.json`. A test fails when that file holds no future date. This is by design.
- Commander brackets: the 2025-10-21 revision. Game Changers: 53 cards, list of 2026-02-09. Lutri is banned as a companion only, per the 2026-02-09 announcement (`companion_bans.json` holds the link).
- Standard: 18 sets, Wilds of Eldraine to The Hobbit. Six sets leave at the first 2027 set. Verified 2026-08-24.
- Card snapshot on disk: `.local/gcs/mtg-local-cards/scryfall/20260824T090152`. Check every card fact against it.
- LLM model ids and prices: `roles.json` and `prices.json`, verified 2026-08-24. The max output per provider in `llm/client.go`, verified 2026-08-29. The OpenAI rows are unverified by anyone but the owner.
- Run cost: the question gate costs $0.152 to $0.165 over about 20 minutes. The eval costs $0.092 to $0.104 over about 13 minutes. The deck gate cost $1.09 on run 8. The revise gate cost $0.29 on run 2. The next runs re-measure all four.
- Toolchain: Go 1.27.0, Node 22.23.2, pnpm 9.2.0, firebase-tools 14.14.0, Java 17.

## How to resume

1. Run `git pull`, then `git status`. Work on `main` or a branch the owner names. The owner commits and pushes.
2. Run `ps aux | grep autotune` before any write. The loop resets the tree when it rejects an iteration.
3. Load the skills. Load `ste-writing` before you write any `.md`. Load `design-doc-style` before you edit the roadmap. Load `mtg-corpus` before you reason about a format, a legality, or a card term.
4. Check the tree: `cd go && go build ./... && go vet ./... && go test -race ./...`, then `make lint` from the root. `make lint` runs the STE check too.
5. Do "Next steps, in order" above. Ask questions as they come up. Record each owner answer in `docs/decisions.md`, and delete the row from `docs/owner-questions.md`.
6. Before you end, update this file.

Six things a fresh session gets wrong without this file.

- Nine targets and the loop script spend money: `make questions-gate`, `make questions-eval`, `make eval-calibrate`, `make deck-gate`, `make revise-gate`, `make chat-probe`, `make generate-probe`, `make summary-judge`, `make test-smoke`, and `scripts/autotune.sh`. Ask the owner before each run. `make autotune` is free.
- A rerun writes to a new file. Every `*_OUT` variable refuses a document that holds a result (D-65).
- A gate run takes about 20 minutes and an eval about 13. A foreground command stops at 10 minutes, so run both in the background.
- `docker compose up` needs provider keys in `.env` now, and fails fast without them (D-267). `make dev` still starts with no keys, and the fake serves the health role only.
- The eval and the agent share a model. Every ratio it reports is a floor, not a measurement (D-136).
- The judge is noisy: two runs of identical code move up to nine bad questions (D-230). One run proves nothing on its own.
