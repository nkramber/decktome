# Session hand-off

`CLAUDE.md` is the entry point. It sends you here. Then read `docs/decisions.md` and `docs/open-questions.md`.

## Last updated

2026-08-25. Merged: PR-0a to PR-5, PR-10, the audit fixes, and PR-6 (#1 to #11).

PR-7 was code-complete, and then the owner scored items 1 to 32 of the M-5 sheet. Those scores asked for 16 rewords and one deletion. They also exposed a class of defect the sheet could not hold. A correction session followed on the same day. Read "The correction session" below before anything else.

The Go tree is green after the correction: `go build`, `go vet`, `go test`, and `golangci-lint` all pass. Gate run 14 has not run yet, and it is the next step. It costs money and needs the owner's approval.

The correction is **not committed**. Branch `pr-7` sits at commit b349bb6 with the owner's scores. Run `git status` before you change anything, and run `git pull` first: the owner scores the M-5 sheet on a phone, so the remote can hold newer scores. The owner commits and pushes.

## State of the work

Done in session 1:
- Read the reference doc `connector-syncer-docs/docs/document-summary-roadmap.md` and captured its structure in the `design-doc-style` skill.
- Surveyed connector-syncer's AI, eval, and ops patterns. Notes in `docs/reference/`.
- Surveyed wallabee-ui and CI conventions. Notes in `docs/reference/wallabee-toolchain-conventions.md`.
- Researched Scryfall: the API, rate limits, bulk files, image terms, catalogs, and Oracle tags. Also ManaBox CSV columns, the 2026 ban-list changes, Standard legal sets, Commander rules and brackets, and ASD-STE100 Issue 8.
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

New decisions: D-66 gives the rubric. D-67 holds the card-pool question until format, colors, and theme are filled. D-68 freezes the slots when a run starts. D-69 makes the gap score its own call, with a first threshold of 0.35.

## PR-7 state (2026-08-24)

PR-7 is the question workflow and the first call site of `internal/llm`. Every item is built, and both gate bars pass. `go build`, `go vet`, `go test`, and `golangci-lint` are all green on the Go tree. The work is committed and pushed at 53ddbf4.

### What exists

| Package | What it holds |
|---|---|
| `internal/questions` | The turn loop. `catalog.json` (26 rows of corpus section 11 as data), `catalog.go` (load and validate), `plan.go` (the planner and the word rules), `state.go` (slots, keys, freeze, card lists), `agent.go` (classify, score, ask), `prompts.go` (three prompts and their schemas), `resolve.go` (placeholders, clause drop, phrasing guard), `snapshot.go` (the private state as data), `metrics.go` (the M-4 records and `Coverage`), `hints_candidates.go` (PR-6 as the value source). |
| `internal/sessions` | The Firestore store. Two documents per session, written in one transaction. |
| `internal/agentsvc` | `AgentService.Chat` and `GetSession`. It owns the facts the classifier can not answer. |
| `cmd/questions-gate` | The live gate runner and the M-4 document. It refuses to run without `QUESTIONS_GATE=1`. |
| `cmd/m5-sheet` | The M-5 scoring sheet, built from the gate documents. No model calls, no cost. |
| `cmd/m5-report` | Reads the scored sheet and computes the two thresholds. No model calls, no cost. |

Rules the code holds:

- At most three questions in one turn, and one question per proto slot.
- Never a repeat, with one exception: the commander pick row repeats after a "none" answer (D-73).
- Nothing at all when the run is frozen (D-68).
- The card-pool question waits for format, colors, and theme (D-67).
- A turn costs three model calls: classify, score, ask (D-69). The provisional fit threshold is 0.35.

### What this session did

The three items the last session named are fixed, and the owner answered two questions along the way.

1. **The locked-cards row (D-70).** It fires only for a named card that is not the commander. The state now keeps commander names and locked names apart. The text is "Must the deck keep {locked}, or may I cut a card that does not fit the plan?"
2. **The commander row split (D-71).** The base row asks whether the user has a commander. The pick row carries the three names and has its own key. The classify schema gained `facts.wants_suggestion`, which sets `Suggested`. A named commander closes all three commander rows.
3. **The cache key (D-72).** Every model call of one session carries the session id. It saves nothing today, because a call sits near 470 tokens and the OpenAI cache starts above 1,024. PR-8 owns the lever, and it must measure a role-scoped key against the session key.

Owner answers of this session:

- **D-73.** A "none" answer to the commander pick repeats the row until the user takes one. Each round names three commanders the agent did not offer before. The owner chose this over a reworded row and over a cap of two asks.
- **D-74.** Firestore holds the proto session and a private state document. The owner chose this over a rebuild from `Session.turns` and over memory-only state.

### Four more dead rows, found while the service was wired

Each one is a catalog row that could never fire in production. Each has a test.

1. **An asked key could never close (D-76).** `openKeys` listed only the rows the planner would ask next. The no-repeat rule drops an asked row from the plan. The classifier was therefore never offered the key of the question it had just asked. Every advisory row then stayed open forever, and `Ready` never became true. The rows: jank, table tolerance, budget scope, acquisition, house format limits, named card role, locked, meta, and plan.
2. **A commander left the color slot open (D-75).** The colors row never fires when a commander is set, so nothing closed that key. The card-pool question waits for the colors (D-67), so a user who names a commander was never asked about their library. `SetCommander` now closes the color key.
3. **`AfterBuild` had no source (D-77).** The variance row was dead. `agentsvc` reads it from `Session.deck_ids`, so the row becomes live when PR-8 stores a deck.
4. **`ThinTheme` and `CommanderNotOwned` had no source (D-77).** `agentsvc` answers both from PR-6 and the collection. `ThinTheme` runs one candidate build per theme, only while the pool key is open.

`WeakCommanderPool` is wired (D-94). `CommanderPool` counts the on-theme commanders an owned pool holds, and an empty pool is a weak pool. No threshold was invented.
### The M-4 report

Each question leaves an `Ask` record: question id, row, slot, key, source, gap score, threshold, turn, and whether the key closed. The records live in the private state, so they survive a resume. `Coverage` sums them over one session or over many. The fields: asked, catalog, invented, catalog filled, invented filled, catalog-only sessions, the gap scores, and the invented count per row. A row that repeats in that last map is a catalog change candidate (D-25, PR-15).

### The gate

The offline half holds. `TestConversations` runs 30 scripted conversations. Each reaches a complete slot set in at most four turns, with no repeated question. `TestEveryRowFires` proves that all 25 catalog rows fire in at least one conversation.

The live half ran four times on 2026-08-25. Every run has its own document, and a rerun never overwrites a scored one (D-65).

| Measure | Run 8 | Run 9 | Run 10 | Run 11 | Run 12 | Run 13 |
|---|---|---|---|---|---|---|
| Catalog-only, bar 25 | 26 | 29 | 28 | 27 | 27 | 27 |
| Premature sessions (D-91) | 0 | 0 | 0 | 1 | 0 | 0 |
| Verdict | pass | pass | pass | fail | pass | **pass** |
| Gate questions asked | 135 | 133 | 138 | 136 | 137 | 132 |
| Invented | 4 | 1 | 2 | 3 | 3 | 4 |
| Refused as rewords | 6 | 7 | 7 | 5 | 5 | 7 |
| Catalog questions that closed a slot | 78 | 85 | 92 | 89 | 93 | 85 |
| Probe conversations | - | - | - | 22 | 22 | 22 |

Runs 1 to 7 are in the earlier documents. They scored 21, 23, 23, 23, 29, 30, and 28 on the catalog bar, and the second bar did not exist before run 6.

Runs 11 to 13 each ran 52 conversations: the 30 the gate scores, and 22 probes. Run 11 failed on one probe, which found D-98. Runs 12 and 13 pass with no premature session anywhere, which is the live proof of D-98 and D-99.

Runs 1 to 4 scored 21, 23, 23, and 23 on the catalog bar, at 108, 91, 86, and 113 questions. The second bar did not exist yet.

Read the last two rows, not the first. Three runs passed a bar while the product got worse, and one failed while it got better.

Run 7 passed both bars and was hollow. Twenty-six of thirty conversations ended with a slot unanswered, and twenty-five sat on "format (asked, no answer)". Fewer slots filled means fewer rows fire, fewer questions go out, and fewer chances to invent one. Both bars improve as the agent stops working.

Run 8 is the first honest pass. It asked 135 questions, the most of any run, and the format-stuck count fell from 25 to 3.

Run 10 adds the commander ranking (D-94). It asked 138 questions and closed 92 of them, both the best of any run.

Run 9 adds the decline (D-93) and holds. Conversation 4 is the proof. In run 8 it ended with the colors and the commander open. In run 9 the user's "any colors are fine" closes the color slot. Only the commander is left, because the script runs out of messages while the pick row waits.

The reword guard changed the catalog bar, not the model. Run 5 refused 15 of the 16 replacements the model offered. The catalog-only count is therefore only as sound as the 0.6 overlap threshold. The M-5 sheet carries refused rewords, so the owner can check that number.

Total spend: about $0.40 over eight runs, one smoke, and roughly ninety probe calls.

### Twenty defects, and what proved each one

The first five came from run 1. The owner approved every fix.

1. **Commanders outside the color identity (D-82).** The resolver asked the PR-6 hint source for every row, and it cached the answer under the theme alone. A list built before the user named their colors therefore survived the conversation. Conversation 22 asked for a blue-red deck and got Lotho, Corrupt Shirriff (white and black), Peregrin Took (green), and Massacre Girl, Known Killer (black). Run 4 offers only legal commanders.
2. **"Casual" triggered house rules (D-78).** A parent asked for a child's deck and answered "casual means low power, not a house format".
3. **The competitive theme row never closed its key (D-79).** The generic theme row then fired one turn later and read as a repeat.
4. **The pick row named three others every turn (D-80).** The user answered a table question and got three new commanders.
5. **Two rows fired before their context existed (D-81).** Locked cards now requires the format and the theme. House format limits requires the house-rules answer.

Run 2 found two more.

6. **A slot closed on silence (D-83).** The classifier retired power and the pool rule from "Brago blink deck from my library", and the session called itself complete after one question. A deck would have had a power level nobody chose.
7. **The gap score was not reproducible (D-84).** One question scored 0.02 and 0.98 in two turns of one conversation. The score prompt asked a question of taste. It now names four faults, and the score follows the fault count. Run 3 proved the change: every fit landed on 0.90, 0.20, or 0.05, against nine scattered values before.

Run 3 found three more, and one of them was mine.

8. **D-83 was too strict, and run 3 caught it.** The first rule blocked every typed slot from closing by name, on the theory that the classifier always returns a typed value. It does not. It reports the answer in the free-text list and leaves the field unknown. The format slot stayed open after a user answered "Pioneer", and the agent asked again. The rule is now the question, not the field type: a key closes by name only while its question is out. Slot closure went from 27 of 76 back to 79 of 105.
9. **The color clause read as nonsense (D-86).** D-79 made "the best deck under budget" a theme value. The color row then used it as a subject: "the best deck under budget is strongest in white, blue, black, and green". The clause now drops when the theme names no archetype, or when the answer holds more than two colors.
10. **The acquisition row had no reader (D-87).** It asked where the user buys and by what date they need the cards. The app holds no store stock and no delivery times, and Scryfall gives a price estimate rather than availability (D-17). The owner deleted the row. The catalog holds 25 rows.

Run 4 found three more.

11. **A replacement dropped a pool mode (D-88).** Two replacements of the card-pool question each offered two of the three modes, so the user lost owned-only or owned-first. An invented question now keeps the options of the row it replaces.
12. **The model rewords, it does not invent (D-88).** Seven of run 4's eight replacements shared most of their words with the row they replaced, at a median word overlap of 0.67. The agent now refuses a replacement above 0.6 and counts it.
13. **The gate documents could not feed M-5.** A run recorded the invented question and not the row it replaced, and the D-66 rubric compares the two. `Question.catalog_text` now carries it (D-85), and `cmd/m5-sheet` builds the sheet (D-89).

Run 6 found the last two, and the first is the third and final form of D-83.

14. **A slot closed with no value (D-83, D-90).** Run 6 let the classifier close the format by name. The key went to filled and the format value stayed empty, so every row that triggers on the format stopped firing. Five sessions ended with no power level, and PR-8 would have had no format to build from. A typed slot now closes on its value alone. The deciding invariant: a slot the deck generator reads must end with a value, and asking twice is the safe failure. The classify prompt now states that an answer counts only in its own field.
15. **The catalog bar hid the defect (D-91).** Run 6 scored 30 of 30 and looked perfect. A conversation that stops asking looks the same as one that finished. A session that calls itself complete with a needed slot unanswered now fails the gate.

Runs 7 and 8 found the root cause under all of it.

16. **A schema enum suppressed the format field (D-92).** The classify schema constrained `format` and `pool_rule` to an enum holding "unknown". The model answered "unknown" for a message that named the format outright, while it filled every free-text field in the same reply. It guessed colors and a power step from "I want a Modern burn deck" and still left the format empty. Four samples per variant: the enum extracted 1 of 8, an enum with an empty member 3 of 8, a free string 7 of 8. Both fields are free strings now, and `slotWord` owns the vocabulary. After the fix the same probe reads 16 of 16, the negative case included. This is why the model kept reporting the format through `closed_keys`. That channel was not suppressed. Every version of D-83 argued about which channel to trust, instead of asking why the field was empty.

17. **Commanders were not ranked as commanders (D-94).** `Commanders` read the 99-card shortlist and took the first legends it met. That list ends in `capByRole`, which emits one role bucket after another with lands first, so it is not ordered by score. A blink request returned three Ojer modal double-faced cards, which the theme scorer rates 0.16, while 85 on-theme blink commanders existed. `CommanderPool` now walks the index itself and requires a theme signal. The 99-card pipeline is untouched, so the PR-6 gate still holds. The pool length is also the weak-commander-pool signal the corpus asks for, so `WeakCommanderPool` is no longer dead.

18. **A declined format killed every row that routes on it (D-98).** Probe 35 of gate run 11 declined the format, so no power row could fire, and the session finished with no power level. A declined format now takes the corpus default and keeps the `SKIPPED` state.
19. **The agent could not decline a request (D-99).** Gate run 11 answered "Can you build me a Yu-Gi-Oh deck?" with "Which Yu-Gi-Oh format would you like?" A new row on a new `scope` slot says the app builds Magic decks only. It fires alone.

Probe conversations found both, on the first run they existed (D-97). Every one of the 30 gate conversations holds a cooperative user who answers what the agent asks. No defect of this shape could reach them.

`State.Skip` was the last dead method, and D-93 fixed it. The classify schema gained `declined_keys`. A decline closes any key, typed or not, and it names no value on purpose. The generator applies the default the corpus lists, and the skipped state is what tells it to. The question must be out first. A live check measured the precision after one prompt correction. The first wording let "any colors are fine" decline the bracket as well, which would have skipped a slot the user never mentioned.

20. **A truncation slipped past the reword guard (D-103).** Word overlap is symmetric, so a replacement that deleted half a row scored low and passed, although it said strictly less. The owner found it while scoring item 8 of the M-5 sheet. The guard now refuses a truncation as well, and the rule refuses item 8 and nothing else across all 60 items.

### What M-5 was for

M-5 sets two numbers: the fit threshold (D-27) and the reword overlap bar (D-88). The session of 2026-08-25 chose both from its own data, which is exactly what M-5 exists to replace. The paragraphs below describe the version-1 sheet, which the owner scored to item 32. The correction session replaced that sheet with version 2. Read "The correction session" for what changed.

`docs/reference/pr7-m5-scoring.md` holds 60 items from runs 10 to 13, which passes the 50 the D-66 rubric asks for. Runs 1 to 9 are held back (D-96). Each ran with a defect that produced or distorted the replacements it recorded. Run 3 is the clearest case. The format enum threw the user's answer away, the format row fired again, and the model invented a repair question. Scoring that item reads as "the catalog was not enough and the invention was better". It would push the D-27 threshold up on evidence about a bug that no longer exists.

Every item scored 0.05 or 0.20, because the scorer counts faults. A 0.05 means two faults or more, and a 0.20 means one clear fault. The threshold question is therefore close to a yes or no. Nineteen items sit at 0.05 and 41 at 0.20, so the answer turns on the 0.20 group. Each one pairs a question the model offered with the catalog question it would have replaced. Score the six fields of D-66 on every item. Two items repeat an earlier one, unlabeled, and they measure self-consistency.

Twenty-six items went out to the user. Thirty-four are rewords the agent refused, and they never reached anyone. Both kinds are scored the same way. A `better` on a refused item says the reword guard is set too tight. This session chose the 0.6 threshold from run-4 data, not from the rubric. Run 6 shows why it matters: the guard, not the catalog, decides the catalog-only count.

One nuance when scoring a refused item. The guard compared the replacement against the resolved catalog row. The sheet shows the text that actually went out, which the ask role had already fitted to the user. Score what you see.

The threshold follows from the scores. The scale is discrete now: 0.90 means no fault, 0.20 one clear fault, and 0.05 two or more. At the present 0.35 the agent invents on any clear fault. At 0.10 it would invent only on two faults or more, which in run 3 was 5 replacements instead of 10. D-27 reserves the number for M-5, so the sheet sets it.

`make m5-sheet` rebuilds the sheet after a new run. It costs nothing and refuses to overwrite a sheet that holds scores.

### How to run it

- `cd go && go test ./...` runs everything offline against the fake provider.
- `make lint-go` runs vet and golangci-lint.
- `make themes-check` checks the PR-6 theme slugs against the local snapshot.
- The live format check runs against the real provider and guards D-92:
  `set -a && . ./.env && set +a && QUESTIONS_LIVE=1 go test ./internal/questions -run TestLiveFormatExtraction -v -count=1`
- The live conversation needs approval before every run:
  `set -a && . ./.env && set +a && QUESTIONS_LIVE=1 go test ./internal/questions -run TestLiveConversation -v -count=1`
- The live gate needs approval and writes the gate document. A rerun needs a new file (D-65):
  `GATE_OUT=docs/reference/pr7-question-gate-run5.md make questions-gate`. A cheap check first: `QUESTIONS_GATE=1 go run ./cmd/questions-gate -n 2 > /tmp/gate.md`.
- `make m5-sheet` builds the scoring sheet from the runs `M5_RUNS` names, which is run 14 alone. It costs nothing, and it needs a new `M5_OUT` name.
- `make m5-report` reads the scored sheet and computes the thresholds. It costs nothing, and it works on a half-scored sheet. It reads `M5_OUT`, which still points at the version-1 sheet.
- `go test ./internal/questions -run TestCatalogIsClean` runs the question linter over the catalog rows. It costs nothing, and CI runs it.
- `make store-check` runs the session store against the local Firestore emulator. Start one first:
  `firebase emulators:start --only firestore --project mtg-local`

A full gate run of the 52 conversations cost about $0.09 and took 11 minutes, measured over four runs. D-105 added 14 conversations, so run 14 should cost about $0.12 and take about 14 minutes. That is an estimate.

### Known limits

- The `collections` repo has no Firestore emulator test. `internal/sessions` now has one (D-101), and the same pattern would cover collections.
- `make test` fails on the web half in this environment. Node 20.17.0 is active and the repo pins 22.23.2 (D-52). The Go half is green. The owner action from the audit session still stands: `nvm install && nvm use`, then `cd web && pnpm install`.
- Grist, the Hunger Tide can not be a commander in the engine (no Scryfall signal).
- The ManaBox condition vocabulary beyond `near_mint` is unverified.

## The correction session (2026-08-25)

The owner scored items 1 to 32 of `docs/reference/pr7-m5-scoring.md` and reported that many questions were simply wrong, apart from anything M-5 measures. This session read every score, verified each claim against the four gate documents, and fixed what the evidence supported. Decisions D-104 to D-116 record every owner answer.

### What the 32 scores said

| Field | Values |
|---|---|
| `catalog_enough` | yes 17, no 13, n/a 6 |
| `invented_better` | better 13, same 12, worse 5, n/a 6 |
| `right_slot` | yes 27, no 3, n/a 6 |
| `catalog_action` | **reword 16**, none 14, delete the row 6 |

Half the rows the model touched needed wording work, and one row needed deletion. That is a catalog verdict, and not a threshold verdict. It is why `make m5-report` still finds no threshold that meets the 80 percent floor.

### Why the owner kept finding wrong questions

The version-1 sheet held only a question the model offered to replace. That is 60 of the 793 questions of runs 10 to 13, or 7.6 percent. A question that was wrong, and that the model never challenged, could not become an item. Deterministic rules find a defect in 61 of the 793 questions, and in 159 when the presumed-table wording counts. D-104 makes every question scorable.

### The five root causes

1. **The question contradicted what the user said.** Three mechanisms. The classifier missed a stated format (2 of 79 chances). The classifier could not infer a format from "not as my commander", and conversation 27 then asked the role question in all four runs. Word triggers fired on a negation, so "no proxies" and "whatever is winning" raised house rules nine times.
2. **The row presumed a fact.** 90 of 793 questions said "your table" to a user who named none. The colors row asserted which colors a theme is strongest in, and run 13 wrote "Black Lotus decks are strongest in black". Black Lotus is a colorless card.
3. **Template defects.** `{locked}` rendered "Grist, the Hunger Tide and Grist" in three runs. The colors row rendered "A a dragon deck deck".
4. **Rows that must not exist, and rows that were missing.** `table_tolerance` fired 20 times and the owner deleted it. No row said the app builds one deck at a time. No row named an unsupported format, and run 13 offered Brawl.
5. **Verbosity, with a rule under it.** The owner marked a shorter replacement better when it dropped a fitted clause, and worse when it dropped a constraint. D-116 puts that rule in the ask prompt.

### What changed in the code

- **Catalog**: 26 rows became 29. `table_tolerance` is gone (D-110). Four rows are new: `one_deck`, `format_unsupported`, `power_sixty_confirm`, and `pool_precon` (D-107, D-112, D-113). The colors row asserts nothing (D-108). The power and meta rows presume no table (D-109). The house-rules row lost its proxy and "whatever" triggers (D-111).
- **`internal/questions/words.go`** is new. It holds the deterministic word rules: negation, format inference, unsupported formats, a two-deck request, a precon, a proxy user, a card placed in the 99, and a competitive request. Every rule costs no model call.
- **`internal/questions/lint.go`** is new. It is the question linter of D-115. It reads text alone, it runs in CI over the catalog, and `cmd/questions-gate` runs it over every conversation. A finding fails the gate.
- **State**: `sameCard` merges a short card name into the full one, so the locked row no longer stutters. `AddLocked` closes the role question, because a card placed in the 99 has a settled role.
- **Prompts**: version 2. The classify role reads a format from an adjective and from a commander phrase. The ask role adds no clause that repeats a known value, states no fact, and presumes no table.
- **`cmd/m5-sheet`**: rubric version 2. The sheet holds plain items beside replacement items, and it carries `warranted`, three more faults, and a legal `n/a`.
- **`cmd/m5-report`**: an `n/a` item leaves the fit calculation. A refusal that copies the row word for word is counted apart, because the guard refused an exact copy and that says nothing about the guard.
- **Conversations**: 52 became 66. Conversations 1 to 30 are byte-identical (D-105). Conversations 53 to 66 are the terse set, and their users never repeat a value they already gave.

### What four smoke runs found

Four cheap runs went out before gate 14, at a total of about 1.3 cents. `cmd/questions-gate` gained an `-only` flag, so one conversation can be checked without a full run.

They proved the main fixes live. Conversation 53 asks no format question after "A land destruction Commander deck." Conversation 54 asks no role question after "not as my commander", and the locked row no longer stutters. The power row names no table, and the colors row states no fact.

They also found two defects, and both are fixed. The model replaced the one-deck row and dropped the sentence that names the limit, so D-117 makes such a row fixed. The linter then fired on the row that declines an unsupported format, which is the one row that must name it. A loose test had let that through, and the test is tighter now.

### The batch sweep (2026-08-26)

The owner asked for the gate to run three conversations at a time, from 1 to 66, with every defect fixed as it appeared. Twenty-two batches ran, plus re-runs, for about $0.17. The linter found no defective question in any batch after the first fix.

The sweep found sixteen more defects. Decisions D-117 to D-132 record them. None of them could reach the version-1 M-5 sheet, because the model offered a replacement for none of them.

| Defect | Where it showed | Decision |
|---|---|---|
| A limit row was paraphrased away | one deck at a time, Brawl | D-117, D-131 |
| "Build around X" read as "X is my commander" | conversation 3 | D-118 |
| An answer that repeats an option left the key open | conversation 5 | D-119 |
| "None of those" closed the commander pick | conversation 14 | D-120 |
| "The first of the new three" named no commander | conversation 14 | D-121 |
| A message answered the question it raised | conversation 21 | D-122 |
| The offered commanders were swapped unasked | conversation 23 | D-123 |
| The hint source could not see this turn's colors | conversation 23 | D-124 |
| A replaced format came back | probe 31 | D-125 |
| One slot asked twice in other words | probe 33, probe 39 | D-126, D-128 |
| A commander row with nothing to offer | probe 35, probe 47 | D-127 |
| Lightning Bolt accepted as a commander | probe 41 | D-129 |
| "Use a different commander" got silence | probe 49 | D-130 |
| The meta row never fired on a named step | conversation 62 | D-132 |

Two shapes run through most of them. A word the user wrote was read as an answer to a question the agent had not asked, or an answer the user gave never reached the key that needed it. Both are invisible to a sheet that only holds replaced questions.

### How the linter was checked

The linter ran over all 793 questions of runs 10 to 13. It found 6 redundant format questions, 90 presumed tables, 24 stated facts, 2 stuttered names, and 1 unsupported format. Every one of the 6, the 2, and the 1 matches a case found by hand. The historical documents are untouched.

### What is still open

1. The owner scored items 33 to 60 of the version-1 sheet, or stopped at 32. That sheet keeps its scores as a dated record. It measures the old engine, so its threshold no longer applies (D-96).
2. Two duplicate pairs disagree. Items 2 and 20 are the same item, and D-106 settles that item 20 stands. Items 3 and 30 differ on `invented_better` alone. Items 40, 50, and 60 repeat items 4, 5, and 6.
3. OQ-21: how much of a precon must survive a build. PR-8 owns it.
4. OQ-22: the nearest supported format for Historic and for Timeless. Both are unverified in `words.go`.

## The automated eval lane (2026-08-26)

The owner asked for a lane that needs no hand scoring, and for a loop that can run overnight. The parts exist and nothing has run. `docs/reference/autotune-design.md` holds the design, the cost, and the honest limits.

- A new `eval` role scores every question of a gate run. The owner set it on `gpt-5.6-luna`, so a 66-conversation run costs about eleven cents (D-133).
- `cmd/questions-eval` writes a report a person reads and a summary a script reads. `make questions-eval` runs it.
- `cmd/tune-check` decides whether one iteration may be kept. It costs nothing.
- `scripts/autotune.sh` is the loop. `make autotune` prints how to start it and starts nothing.
- Every third conversation is a holdout. The fixer never reads its failures, and the loop reads its ratio (D-134).

A budget of $3.00 buys about 14 iterations of gate and eval. The fixer agent's own tokens are not in that number, and they are the larger cost.

`docs/owner-questions.md` is new. It holds every question that waits for a person, and the loop refuses to decide any of them. Four of them block the first unattended run: OQ-24 to OQ-27.

The recommendation is to run the report alone first, for about fifty cents, and read what it finds. The loop is worth turning on only if the report finds what the batch sweep found by hand.

## Next steps, in order

1. Run gate run 14 as one document. The batch sweep proved the conversations one by one, and it wrote no single gate document, so M-5 still needs one. It costs about $0.12 and takes about 14 minutes, and it needs the owner's approval:
   `GATE_OUT=docs/reference/pr7-question-gate-run14.md make questions-gate`
   The linter is a bar now, so a redundant or presumptuous question fails the run.
2. Build the version-2 sheet from run 14 alone:
   `M5_OUT=docs/reference/pr7-m5-scoring-run14.md make m5-sheet`
3. Score that sheet. It holds 60 items: every replacement, and a sample of the plain questions spread over the rows.
4. Run `make m5-report` against the new sheet. It sets the D-27 threshold, it says whether the reword guard at 0.6 is too tight, and it says which questions should never have been asked.
5. Set `DefaultFitThreshold` and, if the report says so, `MaxRewordOverlap`. Then run the gate once more: a new threshold changes when the model may invent.
6. Review and commit PR-7. The owner commits and pushes.
7. PR-8 (generator), then PR-9 (variance). PR-8 owns the prompt-cache lever, the weak-commander-pool bar, and OQ-21.

## Facts that expire

- Ban-list snapshot: 2026-08-23. Next announcement 2026-10-12.
- Scryfall bulk sizes and counts: 2026-08-23.
- Game Changers: 53 cards, list of 2026-02-09.
- Standard: 18 sets, Wilds of Eldraine (2023-09-08) to The Hobbit (2026-08-14). No rotation in 2026. Six sets leave at the first 2027 set: WOE, LCI, MKM, OTJ, BLB, DSK. Verified 2026-08-24 on the Scryfall sets API.
- LLM model ids and prices: 2026-08-24 (`roles.json`, `prices.json`). The Sonnet 5 intro price claim is unverified.
- Comprehensive Rules: 2026-08-07 text. Commander brackets: 2025-10-21 revision.

## How to resume

1. Run `git pull`, then `git status`. The owner scores the M-5 sheet on a phone, so the remote can be ahead. The correction session of 2026-08-25 is not committed, so the tree is not clean.
2. Load the skills: `ste-writing` before you write any `.md`, `design-doc-style` before you edit the roadmap, and `mtg-corpus` before you reason about a format, a legality, or a card term.
3. Read `docs/decisions.md` (D-1 to D-134), `docs/owner-questions.md`, and `docs/open-questions.md`. The decision log is the source of truth, and this file is the summary.
4. Check the Go tree is green: `cd go && go build ./... && go vet ./... && go test ./...`, then `make lint-go`.
5. Continue from "Next steps, in order". Ask questions as they come up, and record each answer in `docs/decisions.md`.
6. Before you end, update this file.

Four things a fresh session gets wrong without reading further. `make questions-gate` spends money and needs approval each time. A rerun must write to a new `GATE_OUT` file, because a scored document is never overwritten (D-65). `docs/reference/pr7-m5-scoring.md` is the owner's working file, and no target writes to it. The version-2 sheet needs a new `M5_OUT` name.
