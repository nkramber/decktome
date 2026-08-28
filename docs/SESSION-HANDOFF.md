# Session hand-off

`CLAUDE.md` is the entry point. It sends you here. Then read `docs/decisions.md` from D-257, `docs/owner-questions.md`, and `docs/open-questions.md`. The session narratives of 2026-08-23 to 2026-08-28 sit in `docs/reference/session-log-2026-08-23-to-26.md` and `docs/reference/session-log-2026-08-27-to-28.md`.

## Where things stand (2026-08-28)

- `main` is at `29dee3e`, PR-11 merged (#38). Branch `pr-12` sits on top of it with the PR-12 slice, uncommitted at the end of the session of 2026-08-28. The audit of 2026-08-28 (#17), the baselines (#34), Go 1.27.0 (#37), and every Dependabot pull request of the day are in.
- `docs/audit-2026-08-28.md` is the audit report. Every step of its plan is merged, and the history purge ran (section 10).
- The tree is green: build, vet, `-race` tests, golangci-lint, staticcheck, `buf lint`, `buf breaking`, govulncheck (zero reachable), web lint, typecheck, tests, and build. `make ste-check` is part of `make lint` now.
- `make test` takes seconds again. The four snapshot tests of `candidates` gate on `CARDS_SNAPSHOT_DIR` and run under `make themes-check`.
- The deployable API builds decks now (D-257). Before today only `chat-probe` could.
- PR-8 is done and merged (#15). PR-9 is out of the MVP (D-256). Phase 3 started on 2026-08-28: PR-11 (#38) and PR-12 (#40) are merged. PR-12B is built on branch `pr-12b` and not merged (D-283 to D-285). Its paid gate has not run. Then PR-13 and PR-15.

## The numbers, and why none of them compare with the last run

Every measured number below moved on 2026-08-28. Question gate run 25 and deck gate run 8 ran on the new code the same day, and they are the baselines now. Runs 1 to 24 and deck gate runs 1 to 7 do not compare with them (D-66, D-263):

- The classify prompt is at version 12 (D-260, D-265), and the ask prompt at version 13 (D-290). The generate prompt is at version 10 (D-285). Runs 25 and 26 do not compare with the next run.
- The catalog holds 22 rows. Five retired (D-260), and `house_rules` stores its answer (D-265).
- The conversation set holds 104 conversations, 30 gate and 74 probe, with ids 1 to 105 and no id 67 (D-263). A `has_deck` conversation leaves the catalog-only count.
- The question agent changed in seven places a gate could not see before (audit Q-1 to Q-14). The snapshot carries every field. The nearest-format acceptance fills the slot. A retired row may ask again. The word rules read the current message.
- The generator shows the repair turn the `over_budget` and `precon_share` findings, recomputes `Passed`, and prices owned cards as owned (audit G-1 to G-3). The precon share reads the nonbasic names as a ceiling (D-259).
- The card index drops 291 front cards and refuses an ambiguous face name (D-270).

| Measure | Baseline | Where measured |
|---|---|---|
| Question gate | PASS 25 of 27 counted, 2 invented, 0 premature, 0 lint | run 26, 2026-08-28, prompt 12, after D-280 |
| Question eval | 19 bad of 382, holdout 8 of 111 (7.2 percent) | eval of run 25, 5 conversations unjudged |
| Deck gate | PASS 18 of 18, 2 repairs, $1.09 | run 8, 2026-08-28, generate prompt 9 |
| Loop | off since 2026-08-26 | seven starts, nothing kept |

CAUTION: `tune-check` paired zero questions between run 24 and run 25, because the catalog and the prompt changed. The paired guard says nothing across that line, and the whole-run margins carry the verdict. The eval leaves a conversation unjudged when the judge returns fewer verdicts than questions. So 382 is the honest count, not a drop from 435.

## PR-12B, what it holds (2026-08-28)

Branch `pr-12b` holds the revision turn (F-27, D-283 to D-285). The tree is green on the branch: Go build, vet, `-race` tests, golangci-lint, `buf breaking`, web lint, typecheck, 55 web tests, and the web build. Nothing paid has run.

- `internal/revise`: the `revise` role call, the brief, `DiffDecks`, `Note`, and `DeclineNote`. The reply comes from the diff and the declines, never from the model.
- `internal/generate`: `Request.Revision`, the revision block of the generate prompt, `CheckRevision`, `AllowedByRevision`, and `Pool.Filter`. The pool drops the removed cards and the cards over the cap.
- `internal/agentsvc`: `sendRevision` runs after a build when no slot changed, and `slotsChanged` decides. A slot change rebuilds from the start with a status line. `DeckStore` gained `Get`. `Turn.agent_message` is written now.
- `Deck.revised_from_deck_id` and `Deck.revision_note` are additive proto fields. The deck view shows the note and the diff against the deck before it.
- A commander offer before the pool question ranks on quality alone (D-293). The conversation words hold the user words only (D-292). Both came from the browser sessions of the evening.
- A card tile shows the full image and no caption (D-291). The commander offer tile shows the image and the pick button only.
- The commander row asks for a name or a suggestion in one step (D-290). The catalog text changed, so the next question gate run re-baselines (D-66).
- The shortlist follows the commander identity, and the deck view shows the commander from `commander_oracle_ids` (D-289). Before this, a mono-green commander got nine off-color Dinosaurs and no commander tile.
- A commander offer shows each card with its art and its rules text, through the new `Question.option_oracle_ids` (D-287). The branch holds it too.
- `cmd/revise-gate` and `make revise-gate` are the paid gate: two bases, six revisions, a verdict per revision. Ask the owner before the run, then record the numbers here and in the roadmap.

CAUTION: the revise prompt has never met a real model. The first gate run is the first evidence, and the prompt may need a version bump after it.

## PR-12, what it holds (2026-08-28)

Branch `pr-12` holds the slice of the ui plan, section 6. The tree is green on the branch: Go build, vet, `-race` tests, golangci-lint, `buf breaking`, web lint, typecheck, 46 web tests, and the web build. The owner has not run it in the browser yet.

- `CardService.GetCards` returns up to 120 cards by Oracle id, in request order, with a `missing_oracle_ids` list (D-277). `go/internal/cardsvc` holds it and its tests.
- `web/apps/web/src/features/chat`: `use-chat.ts` reads the `Chat` stream and holds the open questions (D-278). `session-page.tsx` is the chat beside the deck. `question-card.tsx` shows the options as buttons and a free-text field. A reload rebuilds the thread from `GetSession` and the latest deck from `GetDeck`.
- `web/apps/web/src/features/deck`: `deck-view.tsx` groups the cards by role and shows the findings, the legality date, the curve, and the color sources. `card-tile.tsx` shows every face with "Illustrated by <artist>. © Wizards of the Coast, LLC" (D-279). `deck-stats.ts` holds the pure helpers. The decks page opens a deck in place.
- Gate run 26 ran after D-280: PASS, 25 of 27 on a bar of 25, against 26 of 27 in run 25. The two invented gate questions are `pool` in conversation 4 and `colors` in one other, and run 25 invented `budget` once. The gate sends plain messages, and `UserWords` returns a plain message unchanged. So the gate ran the code of run 25, and the drop is model noise (D-230). The `Q:`/`A:` shape is still unmeasured, because only the browser sends it.
- The owner's first live session (`eIrL12hRY2YNTTCo3iS4`) showed that a message after a build is dropped and the deck is rebuilt from turn 1 (F-27). PR-12B is the fix, and it is planned and not started (D-283, D-284). Collection delete is not on the roadmap (owner, 2026-08-28).
- The answers go out in one request through "Submit answers", and the message box hides on Send (D-282).
- The UI audit of 2026-08-28 (D-281) fixed 41 defects across the chat, the deck view, the collection screen, the layout, and the styles. 50 web tests pass. The owner has not seen the audited build in the browser yet.
- The owner ran the first browser turns on 2026-08-28 and found two faults, both fixed (D-280). A structured answer echoed the question text into the word rules, and the chat dropped an open question when a new turn asked another. The classify input for a structured answer is `Q: <question>` and `A: <answer>` on two lines now, so the next question gate run measures that shape.
- The pool toggle "Use only cards in my collection" shows before the first message when a collection is active. It decides whether `collection_id` goes with the first message. The agent asks how strict the pool is, and the session's `pool_rule` shows after that.

CAUTION: the web tests need Node 22.23.2 (`.nvmrc`). Under Node 20 every test file fails at start with `ERR_REQUIRE_ESM` from jsdom 30. The shell of 2026-08-28 had Node 20 first on the PATH, and `~/.nvm/versions/node/v22.23.2/bin` fixed it.

## Next steps, in order

1. Run `make revise-gate` with the owner's go-ahead (about $0.60, unmeasured), read every brief and every reply, and fix the revise prompt where it misread. Then the owner tests a revision in the browser and merges `pr-12b`.
2. Watch the first pull request under the new `verify` workflow (D-286). The `changes` job prints the diff and its answers, so a job that skipped when it should have run is visible in that log. The workflow file is one of the inputs of every job, so this pull request runs them all.
3. PR-13: `DeckService.ExportDeck`, the export button, and the buy list with Scryfall links. Ui plan section 4 gives the text shape, and the gate is the round trip through `ParseArenaText`.
4. A ruleset that requires the `verify` check on `main` is not possible. The repo is private on the free plan, and the rulesets API answers 403 (checked 2026-08-28). The fan-in job left the workflow for that reason (D-286). The owner reads the checks before a merge.
5. After PR-12 merges, run the M-5 manual scoring on the first UI build (sequencing step 19), and ask the owner before any paid run.

Done on 2026-08-28: the audit merged (#17), the baselines merged (#34), Go moved to 1.27.0 (#37), and PR-11 merged (#38). Every Dependabot pull request of the day is merged or closed. PR-11 holds the stack, the router, the boundary lint, sign-in and sign-up over the Auth emulator, the token interceptor, and the collection screen. The owner ran the gate in the browser on the real export: 4,952 cards, 2,657 rows, one token row reported as not playable. README section 6 is the browser procedure.

Done on 2026-08-28, after the audit: the merged branches are deleted on the clone and on origin, and the history purge ran. Every hash after PR-6 changed. The repo went from 101 MB to 2 MB, and the tree at each tip is byte for byte the same. CAUTION: a clone made before 2026-08-28 holds the old history. Re-clone it, and do not merge from it.

Two owner questions stay open in `docs/owner-questions.md`: OQ-23 (a user who asks a question back) and OQ-39 (the eval tolerance). OQ-44 (the ManaBox condition vocabulary) waits in `docs/open-questions.md`.

## Facts that expire

- Comprehensive Rules: the current file is 2026-08-19 (D-272). The rule citations in the corpus were checked against the 2026-08-07 text.
- Ban-list snapshot: 2026-08-24. Next announcement 2026-10-12, in `internal/cards/announcement_dates.json`. A test fails when no future date is in that file.
- Commander brackets: the 2025-10-21 revision. Game Changers: 53 cards, list of 2026-02-09. Lutri is banned as a companion only, per the 2026-02-09 announcement (`companion_bans.json` holds the link).
- Standard: 18 sets, Wilds of Eldraine to The Hobbit. Six sets leave at the first 2027 set. Verified 2026-08-24.
- Card snapshot on disk: `.local/gcs/mtg-local-cards/scryfall/20260824T090152`. Check every card fact against it.
- LLM model ids and prices: `roles.json` and `prices.json`, verified 2026-08-24. The OpenAI rows are unverified by anyone but the owner.
- Run cost, measured before today's changes. The question gate costs $0.152 to $0.165 over about 20 minutes. The eval costs $0.092 to $0.104 over about 13 minutes. The deck gate costs about $0.90. The next runs re-measure all three.
- Toolchain: Go 1.27.0, Node 22.23.2, pnpm 9.2.0, firebase-tools 14.14.0, Java 17.

## How to resume

1. Run `git pull`, then `git status`. Work on `main` or a branch the owner names. The owner commits and pushes.
2. Run `ps aux | grep autotune` before any write. The loop resets the tree when it rejects an iteration.
3. Load the skills. Load `ste-writing` before you write any `.md`. Load `design-doc-style` before you edit the roadmap. Load `mtg-corpus` before you reason about a format, a legality, or a card term.
4. Check the tree: `cd go && go build ./... && go vet ./... && go test -race ./...`, then `make lint` from the root. `make lint` runs the STE check too.
5. Do "Next steps, in order" above. Ask questions as they come up. Record each owner answer in `docs/decisions.md`, and delete the row from `docs/owner-questions.md`.
6. Before you end, update this file.

Six things a fresh session gets wrong without this file.

- `make questions-gate`, `make questions-eval`, `make eval-calibrate`, `make deck-gate`, `make chat-probe`, `make generate-probe`, `make summary-judge`, and `scripts/autotune.sh` spend money. Ask the owner before each run.
- A rerun writes to a new file. Every `*_OUT` variable refuses a document that holds a result (D-65).
- A gate run takes about 20 minutes and an eval about 13. A foreground command stops at 10 minutes, so run both in the background.
- `docker compose up` needs provider keys in `.env` now, and fails fast without them (D-267). `make dev` still starts with no keys, and the fake serves the health role only.
- The eval and the agent share a model. Every ratio it reports is a floor, not a measurement (D-136).
- The judge is noisy: two runs of identical code move up to nine bad questions (D-230). One run proves nothing on its own.
