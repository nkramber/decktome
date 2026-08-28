# Session log, 2026-08-27 to 2026-08-28

Frozen history. This is the text of `docs/SESSION-HANDOFF.md` as it stood at `5a1fcfe` on 2026-08-28, before the audit of that day rewrote it. Read `docs/SESSION-HANDOFF.md` for the resume point and `docs/audit-2026-08-28.md` for what changed. Nothing below is updated.

---

# Session hand-off (frozen 2026-08-28)

`CLAUDE.md` is the entry point. It sends you here. Then read `docs/decisions.md`, `docs/owner-questions.md`, and `docs/open-questions.md`. The dated session narratives of 2026-08-23 to 2026-08-26 moved to `docs/reference/session-log-2026-08-23-to-26.md`.

## Do this first

PR-8 is done and marked in the roadmap. Deck gate run 6 passed all three bars on 16 golden prompts, at prompt version 4, for $0.8776. The Phase 3 gate of the sequencing list holds.

The branch `pr-8` is open as a pull request. `main` is at 98eac91 with PR-7c merged.

The precon work of 2026-08-28 closed OQ-40 and finished D-218. `internal/precons` embeds nine decklists the owner supplied, one file per product, and adding a precon is adding a file. Six of the nine are wholly in the owner's collection.

The 85 percent share now holds on both gate prompts, with no repair turn. Three causes had to be fixed, and only one was the model's: the precon reached the pool after the pool was built, the job targets fought the share, and the model could not count its own list (D-247 to D-250).

PR-9 is out of MVP scope (D-256). It blocks nothing: Phase 3 waits on PR-8's gate, which held. The variance row is retired with it, and `Deck.seed` and `Slots.plan_variant` are removed from the contract with their numbers reserved.

Next: Phase 3. PR-11, PR-12, and PR-13 are the UI and the meta snapshot, and PR-15 is the eval harness.

Old next: `Deck.seed` is in the proto and nothing sets it, which is PR-9's first step and a prerequisite for its other two levers.

The question workflow has a valid baseline again: run 24 of 2026-08-28, at classify prompt version 10. It passes at 28 of 30 catalog-only, with no premature session, 20 bad questions of 435 on the whole set, and 12 of 140 on the holdout. No row holds more than two bad questions. Runs 19 to 22 do not compare with it (D-66).

The re-baseline earned its cost twice. Run 23 failed on a regression D-239 had introduced, and D-252 fixed it. The variance row then fired for the first time in any gate run.

NOTE: the `budget_scope` row holds two of the twenty bad questions, and both are trigger faults and not storage faults. It fires when the context already answers it: "build owned-first with a buy list" names the scope, and "I proxy anything over 20 dollars" is a proxy rule and not a budget.



PR-8 is under way on branch `pr-8`, cut from `main` at 98eac91. The generator writes a deck, the code checks every name and every rule, and the deck gate holds.

| Measure | Deck gate run 1 | Deck gate run 2 |
|---|---|---|
| Verdict | FAIL | PASS |
| Decks with no block finding | 10 of 12 | 12 of 12 |
| Invented names that reached the user | 1 | 0 |
| Decks that needed the repair turn | 3 | 0 |
| Cost | $0.9117 | $0.6166 |

Run 1 found one defect and D-225 fixed it. The shortlist leaves basic lands out on purpose, and the generator may name nothing else, so no deck held a basic land. Read `docs/reference/pr8-deck-gate-run2.md`.

Deck gate run 4 passed on 15 prompts, on all three bars: every deck passed the block checks, no invented name reached the user, and no summary stated a false rule. F-26 is a bar of the gate now and not a footnote (D-229). Cost $0.9002.

Two faults that run 4 covers and earlier runs could not. No golden prompt named no commander, so the generator had never run the path D-147 and D-208 create, and a delegated session would have failed the engine on `no_commander` (D-232). The repair turn wrote a changelog into the summary, and prompt version 3 fixed it, which prompt 13 of run 4 exercised.

Deck gate run 3 passed on 13 prompts, with the golden set corrected. The prompt cache is measured (D-227), the cost model is verified (D-228), and the judge lane answers F-26 (D-229). OQ-36 is closed by D-226.

What PR-8 still needs: the precon share waits on a data source (D-240, OQ-40). Everything else the roadmap names is done and measured.

The question layer collected more than the build consumed, six times over: D-226, D-232, D-238, D-240, D-241, and D-242. Each was found by accident. `TestEverySlotIsReadOrNamed` now checks every slot against the build, so the seventh is caught in the test and not in a gate run (D-243).

CAUTION: the deck gate has not run since D-233. The locked-card check of D-242 is a new BLOCK finding, and no gate run has exercised it. Run the deck gate before you trust the bars.

Superseded: The generator, the normalizer, the repair turn, the ownership rule, the precon share, the prompt cache, and F-26 are all done and measured. PR-9 is next.

CAUTION: the catalog-only count of the question gate carries about two conversations of variance. Runs 19 to 22 read 29, 29, 29, and 27 on question code that did not change. The bar is 25, so the gate absorbs it, and one run proves nothing on its own (D-230).

| Measure | Run 20 | Run 21 |
|---|---|---|
| Gate verdict | PASS 29 of 30 | PASS 29 of 30 |
| Bad questions, whole set | 18 of 430 | 27 of 429 |
| Bad-question ratio, holdout | 9.3% | 10.8% |

CAUTION: run 21 is not a regression. Nothing in the agent changed between the two runs, so the pair is a same-code measurement, and D-230 reads it as one.

## The question workflow

The tuning loop ran a fifth time on 2026-08-26, as run `20260826-220840`. Read the result before you trust the loop again.

- The baseline passed the gate. It measured the best numbers so far.
- The one iteration was rejected. The rejection was wrong, and D-217 records why.
- The owner recovered the five changes of that iteration onto `pr-7c`.

The numbers below are the four runs that matter. Run 18 and `191225-000` measured identical agent code, and they differ by three holdout questions.

| Measure | Run 18 | `212512-000` | `220840-000` | `220840-001` | Run 19 |
|---|---|---|---|---|---|
| Agent code | baseline | audit fixes | D-207, D-208 | one iteration | D-209 to D-216 |
| Gate verdict | PASS 28/30 | FAIL 24/30 | PASS 28/30 | PASS 29/30 | PASS 29/30 |
| Premature sessions | 0 | 1 | 0 | 0 | 0 |
| Bad questions, whole set | 32 of 369 | 42 of 453 | 24 of 434 | 24 of 430 | 26 of 423 |
| Bad-question ratio, holdout | 10.4% | 8.5% | 7.3% | 8.8% | 8.9% |
| `duplicate` faults | 16 | 15 | 10 | 3 | 1 |
| Verdict | baseline | failed | baseline | rejected | superseded |

| Measure | Run 19 | Run 20 |
|---|---|---|
| Agent code | D-209 to D-216 | D-219 and D-220 |
| Gate verdict | PASS 29/30 | PASS 29/30 |
| Bad questions, whole set | 26 of 423 | 18 of 430 |
| Bad-question ratio, whole set | 6.1% | 4.2% |
| Bad-question ratio, holdout | 8.9% | 9.3% |
| Verdict | superseded | current |

Run 19 measures D-214 to D-216 against baseline `220840-000`. The three targets landed. The `power_sixty_confirm` row went from 6 bad questions to 0, `duplicate` faults fell from 10 to 1, and conversations 27 and 33 need no invented question. The whole set rose by 2 and the holdout by 2. Both sit inside the margin of 3, and both sit on rows the change never touched (D-217).

Run 20 measures D-219 and D-220. Conversation 33 now gets the open power question, and the competitive theme row is gone. Conversation 74 no longer asks whether Sol Ring can lead a deck, and the commander row asks instead. The whole set fell by 8 bad questions, from 6.1 percent to 4.2 percent.

CAUTION: the holdout did not move. It went from 12 bad questions to 13, which is inside the margin of 3. The whole set and the holdout disagree, and the holdout is the conservative measure. Do not report 4.2 percent as the product ratio. The next run must show the holdout move before that number holds.

Every gain above came from hand analysis, and none came from the loop. D-195 to D-208 removed 18 bad questions in one session. The loop started seven times and kept nothing.

## The loop, and why it stops

D-217 measures the checker against its own noise. Two runs of identical agent code move 16 questions, and one document judged twice moves 12. The whole-run guard rejects an iteration at 3. The guard therefore rejects every iteration, and the record of seven failures follows from that one number.

The guard also reads a text change as proof of an edit. The ask role rewrites its wording every run, so it charges a change with questions on rows nobody touched. The lessons file now holds one false lesson from this (D-217).

CAUTION: `go/cmd/tune-check` and `go/internal/tune` hold the defect. Both are frozen paths, so the fixer can never correct them. Only a person can.

D-217 fixes the guard. It now reads the harm against the help, and `TestIdenticalCodeIsAccepted` pins the number. The fix is not measured against a live run, because no loop ran after it. Start the next loop with `--max 1`, and read the verdict before you trust it.

The owner chose PR-8 next. The loop stays off until a run proves the checker holds.

## State

- `main` is at fdfe15c, "Pr 7b (#13)". Merged: PR-0a to PR-7, PR-7B, and PR-10 (#1 to #13).
- Branch `pr-7c` holds the loop work. HEAD is 351250d. It carries the `v0.0` baseline as evidence, the lesson of the rejected iteration, and the five recovered commits (D-209 to D-213).
- `docs/decisions.md` reaches D-225. D-221 to D-225 are the PR-8 work. OQ-21 is answered, so no owner question blocks PR-8. D-214 to D-216 fix the two named defects of gate run `20260826-220840-000`.
- Branches `auto-tune/20260826-191225` and `auto-tune/20260826-220840` can be deleted. `pr-7c` holds everything they carry.
- The conversation set holds 104 conversations, 30 gate and 74 probe (D-145, D-155).
- Prompt versions: classify and ask 10, eval 3, M-5 rubric 2, generate 3. A score taken at an earlier version does not carry over (D-66). CAUTION: the classify prompt moved to 10 on 2026-08-27 for D-238, so question-gate runs 19 to 22 do not compare with the next run.
- `.local/tune/` holds the JSON of run14, run14b, run16, run17, run18, and every stamped loop run.

## Next steps, in order

1. Commit the tree. The owner commits and pushes. Run 19 and its eval are new files under `docs/reference/`.
2. Answer OQ-21, the precon share. It is the only owner question that blocks PR-8.
3. PR-8 (generator), then PR-9 (variance). PR-8 owns the prompt-cache lever, the weak-commander-pool bar, and OQ-21.

One defect that run 20 found, and that nobody fixed:

- Conversation 33 spends its one invented question on the `locked` row. The ask role rewrote "Must the deck keep {locked}, or may I cut a card that does not fit the plan?" and the guard refused the rewrite. The row is a reword candidate, and no rule is wrong.

## Facts that expire

- Ban-list snapshot: 2026-08-23. Next announcement 2026-10-12.
- Scryfall bulk sizes and counts: 2026-08-23.
- Game Changers: 53 cards, list of 2026-02-09.
- Standard: 18 sets, Wilds of Eldraine (2023-09-08) to The Hobbit (2026-08-14). No rotation in 2026. Six sets leave at the first 2027 set: WOE, LCI, MKM, OTJ, BLB, DSK. Verified 2026-08-24 on the Scryfall sets API.
- LLM model ids and prices: 2026-08-24 (`roles.json`, `prices.json`). The Sonnet 5 intro price claim is unverified.
- Comprehensive Rules: 2026-08-07 text. Commander brackets: 2025-10-21 revision.
- Card snapshot on disk: `.local/gcs/mtg-local-cards/scryfall/20260824T090152`. Check every card fact against it, and never against memory. It holds art-series objects that share a real card's name, so read the `layout` field.
- Run cost, measured 2026-08-26 at 104 conversations. The gate costs $0.152 to $0.165 over about 20 minutes. The eval costs $0.092 to $0.104 over about 13 minutes. One loop iteration costs about $0.25.
- Judge noise, measured 2026-08-26. One document judged twice moves 12 questions. Two runs of identical code move 16. Read D-183 and D-217 before you trust any single ratio.
- Card-pool measurement of 2026-08-26, against the snapshot of 2026-08-24 (D-146). Historic holds 15,680 legal cards and Timeless 15,753.
- Formats the app builds, from 2026-08-26: Commander, Standard, Modern (D-155). Everything else is declined by name with no substitute (D-156).

## How to resume

1. Run `git pull`, then `git status`. Work on branch `pr-7c`. The owner commits and pushes.
2. Run `ps aux | grep autotune` before any write. The loop reverts the tree when it rejects an iteration, so a write during a run is lost.
3. Load the skills. Load `ste-writing` before you write any `.md`, and `design-doc-style` before you edit the roadmap. Load `mtg-corpus` before you reason about a format, a legality, or a card term.
4. Read `docs/decisions.md` from D-207 to the end, `docs/owner-questions.md`, and `docs/open-questions.md`.
5. Check the Go tree is green: `cd go && go build ./... && go vet ./... && go test ./...`, then `make lint-go`. The module sits in `go/`, so `./...` from the repository root finds nothing.
6. Do "Do this first" at the top of this file. Then continue from "Next steps, in order".
7. Ask questions as they come up. Record each owner answer in `docs/decisions.md`, and delete the row from `docs/owner-questions.md`.
8. Before you end, update this file.

Six things a fresh session gets wrong without reading further.

- `make questions-gate`, `make questions-eval`, `make eval-calibrate`, and `scripts/autotune.sh` spend money. Ask the owner before each run.
- A rerun writes to a new file. `GATE_OUT`, `EVAL_OUT`, `EVAL_JSON`, and `M5_OUT` all refuse to overwrite a document that holds a result (D-65).
- A gate run takes about 20 minutes and an eval about 13. A foreground command stops at 10 minutes, so run both in the background.
- `docs/reference/pr7-m5-scoring.md` is the owner's hand scoring. No target writes to it, and a new sheet needs a new `M5_OUT` name.
- Check every card fact against the local snapshot. Two false rules claims reached a user in one run, and both passed the gate and the linter.
- The eval and the agent share a model, `gpt-5.6-luna`. Every ratio it reports is a floor, not a measurement (D-136).
