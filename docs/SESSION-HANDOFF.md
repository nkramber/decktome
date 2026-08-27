# Session hand-off

`CLAUDE.md` is the entry point. It sends you here. Then read `docs/decisions.md`, `docs/owner-questions.md`, and `docs/open-questions.md`. The dated session narratives of 2026-08-23 to 2026-08-26 moved to `docs/reference/session-log-2026-08-23-to-26.md`.

## Do this first

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
| Verdict | baseline | failed | baseline | rejected | current |

Run 19 measures D-214 to D-216 against baseline `220840-000`. The three targets landed. The `power_sixty_confirm` row went from 6 bad questions to 0, `duplicate` faults fell from 10 to 1, and conversations 27 and 33 need no invented question. The whole set rose by 2 and the holdout by 2. Both sit inside the margin of 3, and both sit on rows the change never touched (D-217).

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
- `docs/decisions.md` reaches D-217. D-214 to D-216 fix the two named defects of gate run `20260826-220840-000`.
- Branches `auto-tune/20260826-191225` and `auto-tune/20260826-220840` can be deleted. `pr-7c` holds everything they carry.
- The conversation set holds 104 conversations, 30 gate and 74 probe (D-145, D-155).
- Prompt versions: classify and ask 8, eval 3, M-5 rubric 2. A score taken at an earlier version does not carry over (D-66).
- `.local/tune/` holds the JSON of run14, run14b, run16, run17, run18, and every stamped loop run.

## Next steps, in order

1. Commit the tree. The owner commits and pushes. Run 19 and its eval are new files under `docs/reference/`.
2. Answer OQ-21, the precon share. It is the only owner question that blocks PR-8.
3. PR-8 (generator), then PR-9 (variance). PR-8 owns the prompt-cache lever, the weak-commander-pool bar, and OQ-21.

Two defects that run 19 found, and that nobody fixed:

- Conversation 74 asks "Should Sol Ring be your commander or one of the 99 cards?". Sol Ring can not lead a deck, and D-129 must catch it. The `named_card_role` row does not read the card index.
- Conversation 33 still gets the `theme_competitive` row for "A Modern deck for an event". D-215 closed the `facts.power_competitive` path. The classifier can still fill the power step itself, and `agent.go` reads a filled step of FNM or tournament as a competitive request. Read that path before you trust D-215 as complete.

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
