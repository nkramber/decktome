# Session hand-off

`CLAUDE.md` is the entry point. It sends you here. Then read `docs/decisions.md`, `docs/owner-questions.md`, and `docs/open-questions.md`.

## Do this first

The baseline for the tuning loop is ready: `.local/tune/run18.json`.
`docs/reference/autotune-readme.md` holds every command for the loop.

Gate run 18 and its eval measure the tree as it stands, with D-155 to
D-159 in it. Nothing costs money until the owner approves the next run.

| Measure | Run 16 | Run 17 | Run 18 |
|---|---|---|---|
| Catalog-only, bar 25 | 29 | 27 | **28** |
| Linter findings | 0 | 0 | **0** |
| Premature sessions | 0 | 0 | **0** |
| Questions asked | 124 | 125 | **124** |
| Catalog questions that closed a slot | 98 | 97 | **97** |
| Bad-question ratio, holdout | 11.3% | 13.3% | **10.4%** |

Run 17 rose because of D-157, which this session wrote and then corrected
with D-158. The row that D-157 made repeat, `format_unsupported_open`,
fell from 14 bad questions to 4. The ratio is now under the level that
came before the format change.

Two notes on the eval-18 numbers. The tune split reads 7.8 percent and
the holdout 10.4 percent. No fixer has run, so that gap is sampling and
not the overfit D-134 warns about. The target of D-137 is 5 percent on
the holdout, so the loop has room.

## Last updated

2026-08-26. Merged: PR-0a to PR-7 and PR-10 (#1 to #12), the audit fixes included. PR-7 merged as #12.

Branch `pr-7b` holds PR-7B, the automated eval lane. The last commit is 0771756. `go build`, `go vet`, `go test`, and `golangci-lint` all pass. The tree carries the uncommitted fixes for D-146 to D-158. The owner commits and pushes. Do not commit unless the owner asks.

Versions that invalidate earlier scores: the classify and ask prompts are at 5, the eval prompt is at 2, and the M-5 rubric is at 2. A score taken under an earlier version does not carry over (D-66).

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

## PR-7B, the automated eval lane (2026-08-26)

The owner asked for a lane that needs no hand scoring, and for a loop that can run overnight. The roadmap holds it as PR-7B, on branch `pr-7b`. `docs/reference/autotune-design.md` holds the design, the cost, and the honest limits.

- A new `eval` role scores every question of a gate run. The owner set it on `gpt-5.6-luna`, so a 66-conversation run costs about eleven cents (D-133).
- `cmd/questions-eval` writes a report a person reads and a summary a script reads. `make questions-eval` runs it.
- `cmd/tune-check` decides whether one iteration may be kept. It costs nothing.
- `scripts/autotune.sh` is the loop. `make autotune` prints how to start it and starts nothing.
- Every third conversation is a holdout. The fixer never reads its failures, and the loop reads its ratio (D-134).

A budget of $3.00 buys about 14 iterations of gate and eval. The fixer agent's own tokens are not in that number, and they are the larger cost.

`docs/owner-questions.md` is new. It holds every question that waits for a person, and the loop refuses to decide any of them. The owner answered all four blocking questions on 2026-08-26: D-135 to D-138. One step is left before an unattended run, and the owner owns it: name the fixer in `AUTOTUNE_FIXER_CMD`. No dollar cap applies, because the owner runs Claude Code on a monthly plan (D-159).

### The three evals of PR-7B

The owner started these on 2026-08-26. They cost about fifty cents in total.

1. `GATE_OUT=docs/reference/pr7-question-gate-run14.md make questions-gate`
2. `EVAL_RUN=docs/reference/pr7-question-gate-run14.md make questions-eval`
3. `make eval-calibrate`

The first writes the transcript of all 66 conversations. The second scores every question and writes the report. The third scores 12 conversations twice, once on the cost tier and once on `claude-sonnet-5`, and reports how far the two agree. OQ-39 holds the floor the owner sets from that number.

The PR-7B gate: the report must name every defect class the batch sweep found by hand, and the owner must accept the calibration number.

### What the first three evals found

Gate run 14 passed: 27 of 30 catalog-only, and the linter found no defective question. The eval then found two false rules claims that the gate and the linter both missed.

1. "Grist, the Hunger Tide can not lead a deck." Grist is a Legendary Planeswalker by type line, and it is a legal commander (Scryfall ruling, 2021-06-18). D-129 turned a documented engine limit into a false statement, and D-140 silences the row for any legendary card it can not confirm.
2. "Do you want to use any colors beyond Grist's color identity?" The rules allow no answer. The ask role added the clause, and D-144 puts the rule in the ask prompt and the shape in the linter.

The eval instrument had two defects of its own. It read a question against answers the user gave later (D-141), which is what drove the first ratio to 39.5 percent. It also could not see that a collection was attached (D-143). The ratio after the first fix is 17.4 percent.

The calibration is the lane that pays. The agreement number is 80 percent and it is not the useful output. Reading the ten disagreements is: it found both rules claims, and it showed that `claude-sonnet-5` invents card facts of its own. Neither model is reliable alone, and OQ-39 holds what the owner does about that.

### The conversation set is 100

D-145 added conversations 67 to 100. The set is 30 gate and 70 probe, 312 messages, and every conversation stays inside four turns.

## The session of 2026-08-26, part two

Gate run 15 started at 00:23. `go run` compiles before it runs, so run 15
measures the tree at commit 0771756. Every fix below landed after that
compile, and run 15 therefore does not hold any of them. Read run 15 as
the last measurement of the old code, and as the tuning baseline the
owner asked for.

The owner answered four questions this session. D-146 to D-149 record
them, and every fix carries a test. Gate run 15 then found a fifth
defect, which D-150 records.

### OQ-22 is closed (D-146)

`words.go` mapped Historic to Modern and Timeless to Legacy. Both rows
were marked unverified since 2026-08-25. This session measured the card
pools against the snapshot instead of an argument from memory.

| Format | Pioneer | Modern | Legacy | Standard |
|---|---|---|---|---|
| Historic (15,680 legal) | **0.680** | 0.568 | 0.454 | 0.311 |
| Timeless (15,753 legal) | **0.694** | 0.581 | 0.465 | 0.310 |

The numbers are Jaccard similarity of the legal-card sets. Pioneer is
nearest to both, and the old mapping matched neither. Historic and
Timeless are 0.968 similar to each other, so no measurement separates
them. The owner chose to name no substitute at all. A new row,
`format_unsupported_open`, names the format and asks which format to
build. The catalog holds 31 rows.

### The eval invents card facts (D-149)

Eval run 14 refused three questions on two false claims. It called Ran
and Shaw "not a real Magic card", and Quina, Qu Gourmet "not a valid
Magic card option". The snapshot holds both. Ran and Shaw is a mono-red
Legendary Creature - Dragon, and the user had asked for a red dragon
deck. Quina, Qu Gourmet is a mono-green Legendary Creature - Qu.

Three of the ten `inaccurate` faults of that run were therefore the
eval's own inventions. `cmd/questions-eval` now reads the snapshot and
drops such a refusal. The check costs nothing. The dual-judge proposal
stays open under OQ-39 for the faults this check can not reach.

### A delegation never closed the commander pick (D-147)

`commander_pick` was the worst row of eval run 14, at 18 of 43 bad
questions. That is more than the next four rows together. The cause is
one gap: no rule read "You pick the commander". Eighteen of the 100
conversations hold such a phrase, the row carries `"repeat": true`, and
nothing closed the key. The row therefore asked again every turn until
the messages ran out.

A delegation is now a decline (D-93). D-123 is amended: its test asserted
that the pick row asks again after that exact message, which is the
behavior this decision removes.

### A commander did not hold the colors (D-148)

`CommanderPool` used a subset test, so a blue-red request accepted a
mono-red commander and a colorless one. Conversation 22 was offered
Birgi (mono-red), Emrakul, the Promised End (colorless), and Vnwxt
(mono-blue). Not one is blue-red. The commander's identity is the deck's
identity, so a colorless commander gives a deck that can play no colored
card. `identityCovers` now runs in `CommanderPool` alone, and the 99
keeps the subset test.

### Gate run 15 failed, and it found one more defect (D-150)

Run 15 cost $0.1445 over 1181 seconds, and it asked 135 questions across
the 30 gate conversations. It passed the catalog bar at 26 of 30, against
a bar of 25, and no session called itself complete with a slot
unanswered. It failed on the linter, which is the bar D-144 added.

Two questions named a format the app had just declined.

- Conversation 94: "I do not build Oathbreaker. The nearest format I
  build is Commander." Then, in the same turn: "What should the
  Oathbreaker deck focus on: a creature type, a mechanic, or a play
  style?"
- Conversation 95 did the same with Historic.

The catalog row is innocent. The `theme` row reads "What should the deck
do", and the ask role fitted it to the user's words. D-150 puts the rule
in the ask prompt at version 6, adds it to the phrasing guard, and keeps
the linter rule that caught it. A row that declines a format stays exempt.

The other run-15 counters, for comparison with run 16: 4 invented
questions, 4 refused as rewords, 90 catalog questions that closed a slot,
and a median gap score of 0.90. The 70 probes asked 242 questions and
drew 16 replacements.

### Gate run 16 and its eval (the fixed tree)

Run 16 holds D-146 to D-150. It passes.

| Measure | Run 15 | Run 16 |
|---|---|---|
| Verdict | FAIL | **PASS** |
| Catalog-only, bar 25 | 26 | 29 |
| Linter findings | 2 | 0 |
| Premature sessions | 0 | 0 |
| Questions asked | 135 | 124 |
| Catalog questions that closed a slot | 90 | **98** |
| `commander_pick` questions | 42 | 24 |
| Invented | 4 | 1 |
| Cost | $0.1445 | $0.1192 |

Read the last four rows and not the first. Run 16 asked 11 fewer
questions and closed 8 more slots. Run 7 of 2026-08-25 passed both bars
by the opposite route: it asked less and closed less. A drop in the
question count is only good news beside a rise in the slots closed.

The eval scored 356 questions and refused 40.

| Measure | Eval 14 | Eval 16 |
|---|---|---|
| Bad-question ratio, holdout | 17.4% | **11.3%** |
| `commander_pick` bad questions | 18 | 10 |
| `inaccurate` faults | 10 | 8 |
| Cost | - | $0.0889 |

The eval cost $0.0889 over 698 seconds. `.local/tune/run16.json` is the
tuning loop's baseline.

### The card check found one on its first run (D-149)

The eval called Cloak and Dagger, Entwined "not a real card". The
snapshot holds it: a Legendary Creature - Human Hero with deathtouch and
lifelink, white and black, legal in Commander. The check dropped the
`inaccurate` fault, and the question left the refusal list.

The same card exposed the limit of the check. In conversation 1 the eval
wrote that the card "is not an applicable commander option". That is not
a claim about existence, so the check did not fire, and the fault stands
in the report. The card is white and black with lifelink, and the user
asked for a white-black lifegain deck, so the claim is wrong. Existence
is the wrong refutation for an applicability claim. The right one reads
the color identity against the colors the user named, which is what
D-148 now enforces on the side that writes the question. OQ-39 holds
whether the eval gets that check as well.

### One defect run 16 did not fix

Conversation 3 turn 1 reads "Build around Grist, the Hunger Tide". The
agent asks the format in that turn, and it asks "What color preferences
do you have within Grist's color identity?" beside it.

The clause presumes two facts the agent does not hold. The format is not
known, and color identity is a Commander term. The role of the named card
is not known either, which is the whole subject of conversation 3
(D-118). The catalog row asserts nothing (D-108), so the ask role added
the clause. It is the D-144 class, and it appears once in about 65 color
questions.

### The three fixes after eval run 16 (D-151 to D-153)

The owner answered OQ-41 and OQ-42, and the work found a third defect.

**D-151, the color-identity clause.** It is not the one-off the first
read called it. The same conversation produced one in three runs, and the
model changed one word each time D-144 caught the old one.

| Run | Text |
|---|---|
| 14 | "any colors **beyond** Grist's color identity?" |
| 15 | "should I use **within** Grist's color identity?" |
| 16 | "do you have **within** Grist's color identity?" |

D-144 listed prepositions. The rule now reads the shape. No exception is
needed: every row that asks about the colors carries
`commander_set: false`, so a card's color identity settles nothing there.
Ask prompt version 7, the phrasing guard, and the linter all refuse it.

**D-152, the eval's card facts.** Three of the four card claims of eval
run 16 were false, and every miss came from a crossover set.

| Eval claim | The snapshot |
|---|---|
| Cloak and Dagger, Entwined "not applicable" | White-black, deathtouch and lifelink, legal. The user asked for white-black lifegain. |
| Vincent, Vengeful Atoner "not valid" | Legendary Creature - Assassin, mono-red, legal. |
| Shadow the Hedgehog "not valid" | Legendary Creature - Hedgehog Mercenary, black-red, legal. |
| Jaheira "not a legal red-white commander" | Correct. Every Jaheira holds green. |

The model reasons well from facts it holds and invents the facts it
lacks. The eval input now carries `cards_named_in_questions` for every
card a question offers, with the type line, the color identity, and the
commander legality. Eval prompt version 3.

**D-153, the stale offer.** The one true claim above is a defect in this
code, not in the eval. Probe 73 offered mono-green Jaheira on turn 1 with
no colors named. The user answered "Red and white" on turn 2, and the
same three names went out on turns 2 and 3. D-148 filters the pool, and
these names were already on the table. The agent now drops an offered
name the colors exclude, keeps every name that still fits, and tops the
list up. D-80 and D-123 still hold.

### Pairs: OQ-43 answered (D-154)

OQ-43 held two defects under one label. The color arithmetic was wrong,
and the pick row was single-commander shaped.

The measurement sets the scope. Of 3,384 commander-legal leaders, 177 can
pair, and 31 Backgrounds exist. The exact-match pool by colors named:

| Colors | min | median | max |
|---|---|---|---|
| 1 | 280 | 322 | 351 |
| 2 | 108 | 124 | 145 |
| 3 | 24 | 47 | 58 |
| **4** | **1** | **1** | 11 |
| 5 | 59 | 59 | 59 |

D-148 costs almost nothing at one, two, or three colors. At four it
leaves one commander: WUBR is Breya, WBRG is Saskia, and UBRG is Yidris.
A pair is how those decks are built.

`CommanderPool` now scores a pair as one candidate when the union of the
two identities matches. It offers pairs when the user asks, and when too
few singles fit to fill the three names of the pick row. The pairing
rules come from `rules.ValidPair`, which the deck validator already uses,
so a pair the agent offers is a pair that passes validation. A pair reads
"A + B" in the question.

Two findings from the work. A Background carries no theme signal, so a
pair holding one loses on score to two themed legends, and probe 73's
user got no Background at all. `WantsBackgroundPair` narrows the request
when the user names one. The pair walk also read the wrong enum constant
at first, `PARTNER_KIND_UNSPECIFIED` where the index holds
`PARTNER_KIND_NONE`, which let all 3,384 leaders into a loop meant for
208 cards. `canPair` holds that cut now.

Three snapshot-backed tests cover it: four colors find pairs, a
Background request returns Backgrounds, and a deep two-color pool still
offers singles alone.

## Three formats (D-155 to D-157, 2026-08-26)

The owner narrowed the app to **Commander, Standard, and Modern**. Read
D-155 and not the earlier drafts: the scope changed twice inside the
session. Pioneer, Legacy, Vintage, and Pauper are gone. "Anything goes"
is not a format a user may select, and it stays what D-3 made it, a
house-rules layer on one of the three.

### What changed

| Layer | Change |
|---|---|
| Proto | Four ids deleted and reserved. HOUSE kept for a user who names no format. `buf.yaml` FILE to WIRE_JSON. |
| Engine | `formats.json` holds four entries. `legalKeys` holds three. The restricted-card check is gone. |
| Questions | `formatNames`, `formatIDs`, four catalog rows, the `unsupported` list, classify prompt version 8. |
| Corpus | Sections 2.1, 2.4, 2.5, and the section-11 catalog. |
| Roadmap | Correction pass 21. |
| Conversations | 7, 18, 21, and 26 are probes now. Four replacements joined the gate 30. The set is 104. |

### The proto cost, measured and not assumed

`buf breaking` under FILE refused all four deletions, and it refused them
**although the numbers and the names are reserved**. FILE holds
`ENUM_VALUE_NO_DELETE`, which accepts no reservation. Under WIRE_JSON both
`buf lint` and `buf breaking --against main` pass. WIRE_JSON still refuses
every change that breaks the wire format or the JSON encoding, and this
app has no external consumer that FILE protects.

`make proto-check` fails while the work is uncommitted. It compares the
generated code against the committed tree, so it passes after a commit.

### Three defects the work found

1. **`sixtyCard` never held HOUSE.** No power row could fire for a house
   format: the 60-card rows need that test, and the Commander rows need
   the Commander format. Such a session would have ended with no power
   level. Same class as D-77.
2. **The restricted-card check died with Vintage.** Only Vintage carries
   a restricted list among the formats the app builds, so the branch was
   unreachable. It is removed with its four tests.
3. **Two dead ends the tests could not see (D-157).** A one-cent live
   check on three conversations found both. Probe 18 said "I do not build
   Pauper", the user answered "Pauper." again, and the agent sent nothing
   for two turns. Probe 21 asked the format on turn 1, the user wrote
   "Call it Vintage" on turn 2, and D-126 blocks every row on a key whose
   question is out, so the agent could never decline Vintage. Both are
   the D-130 shape, where a reasonable message gets silence.

### Gate run 17

| Measure | Run 16 | Run 17 |
|---|---|---|
| Verdict | PASS | **PASS** |
| Catalog-only, bar 25 | 29 | 27 |
| Linter findings | 0 | 0 |
| Premature sessions | 0 | 0 |
| Questions asked | 124 | 125 |
| Catalog questions that closed a slot | 98 | 97 |
| Cost | $0.1192 | $0.1540 |

The two-conversation fall is not a regression signal on its own. The gate
set changed, four conversations of it are new, and the classify and ask
prompts are at version 8. The honest counters stand still: the agent asks
the same number of questions and closes the same number of slots, and no
session finished early.

Two invariants held live. Every mention of a removed format in the whole
run is the row that declines it, and no other question names one. The
possessive color-identity clause of D-151 is gone: run 16 held one and
run 17 holds none.

### Eval run 17, and the fix it forced (D-158)

| Measure | Eval 16 | Eval 17 |
|---|---|---|
| Bad-question ratio, holdout | 11.3% | 13.3% |
| Worst row | `commander_pick` 10 | `format_unsupported_open` 14 |
| `duplicate` faults | 22 | 31 |
| Cost | $0.0889 | about $0.09 |

The rise is one row, and that row is D-157 from the same session. Making
the decline repeat every turn fixed real silence and went too far. Probe
18 said "Casual power, and 25 dollars is the cap" on turn 3 and got the
same format question a third time.

The cause is the one D-125 found. The trigger reads `Ctx.Words`, which is
the whole conversation, so a format named once keeps the fact true
forever. D-158 makes the repeat read the current message. The row states
the limit once, and asks again only while the user names the format
again.

**Read eval run 17 as a measurement of the tree before D-158.** The 13.3
percent holdout number does not describe the tree as it stands.

### The pre-flight check, and what it costs

A full gate and eval costs about $0.24 and 31 minutes. A targeted run of
the conversations one change touches costs cents, and it answered this
question twice.

`cmd/questions-gate -only` takes conversation ids. The eleven that fire a
decline row are 7, 18, 21, 26, 50, 60, 61, 69, 97, 98, and 99. Two runs
of those eleven cost $0.0287 in total.

| Measure | Run 17 | First fix | Second fix |
|---|---|---|---|
| Decline questions asked | 23 | 15 | **13** |

The first fix left probes 61 and 97 asking a third time, and the check
found it. The second fix leaves two conversations doubled, and both are
the ones where the user writes "Pauper." or "Legacy." again on turn 2.

Holding every other verdict fixed, the overall ratio projects from
49 of 378 (13.0 percent) to 39 of 368 (**10.6 percent**), against eval
16's 11.3 percent. That is a projection and not a measurement: it assumes
the ten questions that no longer go out are among the fifteen the eval
refused, and it re-judges nothing. A full run would measure it.

Run this check before every full gate. Three defects of this session
survived a green test suite and reached a live run, and two of them cost
under two cents to find.

### State of the tree

`go build`, `go vet`, `go test ./...`, and `make lint-go` are all green.
The work is not committed. The owner commits and pushes.

## Next steps, in order

1. Run gate run 16 and its eval. "Do this first" at the top of this file holds both commands and what to read in the answer. The owner approved both runs.
2. Fix whatever the eval finds, with a test for each fix and a decision row. The last two runs each hid one false rules claim behind a passing gate.
3. Answer the rest of OQ-39. D-149 answers the card-fact half of it with a deterministic snapshot check that costs nothing. The calibration of 2026-08-26 agreed 80 percent, and the ten disagreements held two real defects and two card facts `claude-sonnet-5` invented. The dual-judge proposal stays open for the faults the snapshot can not reach, at about ten cents more per run.
4. Merge `pr-7b`. Then open `pr-7c` from `main` as the container for everything the loop writes (D-142).
5. Name the fixer in `AUTOTUNE_FIXER_CMD`. That agent bills apart from the loop budget, against the owner's monthly plan, so it needs no dollar cap (D-159).
6. Start the loop against `pr-7c`, with `run15.json` as the baseline. One iteration takes about 31 minutes, so a night fits eight to ten:
   `AUTOTUNE_ALLOW_UNATTENDED=1 scripts/autotune.sh --base pr-7c --baseline .local/tune/run15.json --budget 3.00`
   Run it once with `--max 1` before a full night.
7. Build the version-2 M-5 sheet when the owner wants a hand-scored sample beside the automated one:
   `M5_OUT=docs/reference/pr7-m5-scoring-run15.md M5_RUNS=../docs/reference/pr7-question-gate-run15.md make m5-sheet`
8. Set `DefaultFitThreshold` (OQ-28) and, if the report says so, `MaxRewordOverlap` (OQ-29).
9. PR-8 (generator), then PR-9 (variance). PR-8 owns the prompt-cache lever, the weak-commander-pool bar, and OQ-21.

## Facts that expire

- Ban-list snapshot: 2026-08-23. Next announcement 2026-10-12.
- Scryfall bulk sizes and counts: 2026-08-23.
- Game Changers: 53 cards, list of 2026-02-09.
- Standard: 18 sets, Wilds of Eldraine (2023-09-08) to The Hobbit (2026-08-14). No rotation in 2026. Six sets leave at the first 2027 set: WOE, LCI, MKM, OTJ, BLB, DSK. Verified 2026-08-24 on the Scryfall sets API.
- LLM model ids and prices: 2026-08-24 (`roles.json`, `prices.json`). The Sonnet 5 intro price claim is unverified.
- Comprehensive Rules: 2026-08-07 text. Commander brackets: 2025-10-21 revision.
- Card snapshot on disk: `.local/gcs/mtg-local-cards/scryfall/20260824T090152`. Check every card fact against it, and never against memory. It holds art-series objects that share a real card's name, so read the `layout` field.
- Run cost, measured 2026-08-26 at 66 conversations: the gate cost $0.0964 over 737 seconds, and the eval cost $0.0645 over 494 seconds. At 100 conversations that scales to about $0.14 and 18 minutes, and about $0.10 and 13 minutes.
- Card-pool measurement of 2026-08-26, against the snapshot of 2026-08-24 (D-146). Historic holds 15,680 legal cards and Timeless 15,753. Jaccard against Pioneer: 0.680 and 0.694. Against Modern: 0.568 and 0.581. Historic against Timeless: 0.968. Historic bans 77 cards. Timeless restricts 4 and bans none.
- Prompt versions: classify and ask 8, eval 3, M-5 rubric 2. A score taken at an earlier version does not carry over (D-66).
- Formats the app builds, from 2026-08-26: Commander, Standard, Modern (D-155). Everything else is declined by name with no substitute (D-156).

## How to resume

1. Run `git pull`, then `git status`. The last commit is 0771756 on branch `pr-7b`. The tree is not clean: this session's fixes for D-146 to D-158 are uncommitted. The owner commits and pushes.
2. Load the skills: `ste-writing` before you write any `.md`, `design-doc-style` before you edit the roadmap, and `mtg-corpus` before you reason about a format, a legality, or a card term.
3. Read `docs/decisions.md` (D-1 to D-145), `docs/owner-questions.md`, and `docs/open-questions.md`. The decision log is the source of truth, and this file is the summary.
4. Check the Go tree is green: `cd go && go build ./... && go vet ./... && go test ./...`, then `make lint-go`. The module sits in `go/`, so `./...` from the repository root finds nothing.
5. Do "Do this first" at the top of this file. Then continue from "Next steps, in order".
6. Ask questions as they come up. Record each owner answer in `docs/decisions.md`, and delete the row from `docs/owner-questions.md`.
7. Before you end, update this file.

Six things a fresh session gets wrong without reading further.

- `make questions-gate` and `make questions-eval` spend money. The owner approved gate run 15 and its eval, and nothing after that.
- A rerun writes to a new file. `GATE_OUT`, `EVAL_OUT`, `EVAL_JSON`, and `M5_OUT` all refuse to overwrite a document that holds a result (D-65).
- A gate run takes about 18 minutes and an eval about 13. A foreground command stops at 10 minutes, so run both in the background.
- `docs/reference/pr7-m5-scoring.md` is the owner's hand scoring. No target writes to it, and the version-2 sheet needs a new `M5_OUT` name.
- Check every card fact against the local snapshot. Two false rules claims reached a user in one run, and both passed the gate and the linter.
- The eval and the agent share a model, `gpt-5.6-luna`. Every ratio it reports is a floor, not a measurement (D-136).
