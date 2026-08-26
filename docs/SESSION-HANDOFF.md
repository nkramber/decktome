# Session hand-off

`CLAUDE.md` is the entry point. It sends you here. Then read `docs/decisions.md`, `docs/owner-questions.md`, and `docs/open-questions.md`. The dated session narratives of 2026-08-23 to 2026-08-26 moved to `docs/reference/session-log-2026-08-23-to-26.md`.

## Do this first

The tuning loop ran on 2026-08-26, and the same day rebuilt it (D-179 to D-184). `docs/reference/autotune-readme.md` holds every command, and `docs/reference/autotune-lessons.md` holds what the loop learned. The table below holds the three measurements that matter. Run 18 and run -000 measured identical agent code, and they differ by three holdout questions. The same document scored twice differs on 25 of 365 verdicts. The checker now carries that margin (D-183), and it judges each change on paired evidence (D-181).

| Measure | Run 18 | Run 20260826-191225-000 | Run 20260826-191225-001 |
|---|---|---|---|
| Agent code | baseline | identical to run 18 | one fixer iteration |
| Bad-question ratio, holdout | 10.4% | 7.9% | 9.5% |
| Bad-question ratio, whole judged set | 8.7% | 9.4% | 10.3% |
| Verdict | baseline | same code | rejected |

Read the D-171 row before you trust any earlier acceptance. The iteration that branch `auto-tune/20260826-164931` accepted was judged wrong, because the checker compared against a baseline it could not read. Nothing costs money until the owner approves the next run.

CAUTION: every baseline on disk is stale. D-163 to D-169, D-195 to D-202, and the candidates fixes changed the agent code after run -000. The next loop start must measure its own baseline: drop `--baseline`.

## State

- `main` is at fdfe15c, "Pr 7b (#13)". Merged: PR-0a to PR-7, PR-7B, and PR-10 (#1 to #13), the audit fixes included.
- Branch `pr-7c` is the container for everything the loop writes (D-142). HEAD is 3391748, "v0.0: baseline run for the tuning loop". The tree is green and committed.
- Branch `auto-tune/20260826-164931` was dropped (D-204). D-172 to D-176 are retired numbers.
- `docs/decisions.md` reaches D-206 on `pr-7c`. D-179 to D-184 are the loop rebuild, D-185 to D-194 the audit fixes, D-195 to D-203 the question-workflow fixes and the D-162 to D-170 verdicts.
- The work of 2026-08-26 after the loop run is uncommitted: about 70 files. The owner commits. `git status` lists them, and every one is on `pr-7c`.
- Gate documents run 2 to run 18 and `20260826-191225-000` sit under `docs/reference/`. Eval documents exist for run 14, 14b, 16, 17, 18, `20260826-191225-000`, and `-000b`. The `-000b` report is the same gate document scored a second time, and it measures the judge alone (D-183).
- `.local/tune/` holds the JSON for run14, run14b, run16, run17, run18, and `20260826-191225-000` and `-001`. No JSON for run 15 exists.
- The conversation set holds 104 conversations, 30 gate and 74 probe (D-145, D-155).
- Prompt versions: classify and ask 8, eval 3, M-5 rubric 2. A score taken at an earlier version does not carry over (D-66).
- The fixer is named in `AUTOTUNE_FIXER_CMD` and bills apart from the loop budget (D-159).

## Next steps, in order

1. Commit the tree on `pr-7c`. Read `git status` first: it holds the loop rebuild, the audit fixes, and the rebuilt D-163 to D-169.
2. `make store-check` passed under the emulator, and the Compose stack came up healthy with the API as nonroot, both on 2026-08-26. Nothing is left to verify there.
3. Both questions the work raised are answered: D-205 sets the color rule, and D-206 keeps D-168 for the next run to measure.
4. Start the loop with `--max 1` and no `--baseline`, so it measures the changed code first. Read the lessons file after it.
5. PR-8 (generator), then PR-9 (variance). PR-8 owns the prompt-cache lever, the weak-commander-pool bar, and OQ-21.

## Facts that expire

- Ban-list snapshot: 2026-08-23. Next announcement 2026-10-12.
- Scryfall bulk sizes and counts: 2026-08-23.
- Game Changers: 53 cards, list of 2026-02-09.
- Standard: 18 sets, Wilds of Eldraine (2023-09-08) to The Hobbit (2026-08-14). No rotation in 2026. Six sets leave at the first 2027 set: WOE, LCI, MKM, OTJ, BLB, DSK. Verified 2026-08-24 on the Scryfall sets API.
- LLM model ids and prices: 2026-08-24 (`roles.json`, `prices.json`). The Sonnet 5 intro price claim is unverified.
- Comprehensive Rules: 2026-08-07 text. Commander brackets: 2025-10-21 revision.
- Card snapshot on disk: `.local/gcs/mtg-local-cards/scryfall/20260824T090152`. Check every card fact against it, and never against memory. It holds art-series objects that share a real card's name, so read the `layout` field.
- Run cost, measured 2026-08-26 at 104 conversations: the gate costs $0.152 to $0.165 over about 20 minutes, and the eval costs $0.092 to $0.099 over about 13 minutes. One loop iteration costs about $0.25 and takes 33 to 35 minutes, so a $3 budget buys about 12 iterations. A Sonnet 5 calibration costs $0.25 to $0.30.
- Card-pool measurement of 2026-08-26, against the snapshot of 2026-08-24 (D-146). Historic holds 15,680 legal cards and Timeless 15,753. Jaccard against Pioneer: 0.680 and 0.694. Against Modern: 0.568 and 0.581. Historic against Timeless: 0.968. Historic bans 77 cards. Timeless restricts 4 and bans none.
- Prompt versions: classify and ask 8, eval 3, M-5 rubric 2. A score taken at an earlier version does not carry over (D-66).
- Formats the app builds, from 2026-08-26: Commander, Standard, Modern (D-155). Everything else is declined by name with no substitute (D-156).

## How to resume

1. Run `git pull`, then `git status`. Work on branch `pr-7c`. The tree is green. The owner commits and pushes.
2. Run `ps aux | grep autotune` before any write. The loop reverts the tree when it rejects an iteration, so a write during a run is lost.
3. Load the skills: `ste-writing` before you write any `.md`, `design-doc-style` before you edit the roadmap, and `mtg-corpus` before you reason about a format, a legality, or a card term.
4. Read `docs/decisions.md` from D-179 to the end, `docs/owner-questions.md`, and `docs/open-questions.md`. The decision log is the source of truth, and this file is the summary.
5. Check the Go tree is green: `cd go && go build ./... && go vet ./... && go test ./...`, then `make lint-go`. The module sits in `go/`, so `./...` from the repository root finds nothing.
6. Do "Do this first" at the top of this file. Then continue from "Next steps, in order".
7. Ask questions as they come up. Record each owner answer in `docs/decisions.md`, and delete the row from `docs/owner-questions.md`.
8. Before you end, update this file.

Six things a fresh session gets wrong without reading further.

- `make questions-gate`, `make questions-eval`, `make eval-calibrate`, and `scripts/autotune.sh` spend money. Ask the owner before each run.
- A rerun writes to a new file. `GATE_OUT`, `EVAL_OUT`, `EVAL_JSON`, and `M5_OUT` all refuse to overwrite a document that holds a result (D-65). The loop labels each document with its run stamp for the same reason (D-177).
- A gate run takes about 20 minutes and an eval about 13. A foreground command stops at 10 minutes, so run both in the background.
- `docs/reference/pr7-m5-scoring.md` is the owner's hand scoring. No target writes to it, and a new sheet needs a new `M5_OUT` name.
- Check every card fact against the local snapshot. Two false rules claims reached a user in one run, and both passed the gate and the linter.
- The eval and the agent share a model, `gpt-5.6-luna`. Every ratio it reports is a floor, not a measurement (D-136).
