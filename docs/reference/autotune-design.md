# The automated tuning loop

Written 2026-08-26. Nothing here has run yet. The parts exist and the tests pass. The owner answered OQ-24, OQ-26, and OQ-27 the same day. OQ-25 is the last one, and it decides whether an agent may hold the shell unattended.

## The problem

The owner scored 32 questions by hand and it took hours. A gate run asks about 250 questions. Hand scoring does not scale, and the batch sweep of 2026-08-26 proved that most defects are found by reading transcripts rather than by scoring replacements.

## The shape

Six steps, in a loop.

1. Run the question gate. It asks 66 conversations on `gpt-5.6-luna`, and it writes a document.
2. Score every question with the eval role, also on `gpt-5.6-luna` (D-133).
3. Compare the run with the one before it.
4. Hand the report to a fixer agent, which changes the product.
5. Keep the change when the run got better. Revert it when it did not.
6. Stop at the budget, at the target ratio, or after three iterations that changed nothing.

`scripts/autotune.sh` is the loop. It calls `cmd/questions-gate`, `cmd/questions-eval`, `scripts/autotune-fix.sh`, and `cmd/tune-check`.

## Is it feasible

Yes, and the cost is the easy part.

| Item | Per iteration |
|---|---|
| Gate run, 66 conversations | about $0.12 |
| Eval, 66 calls on the cost tier | about $0.09 |
| Total inside the budget | about $0.21 |

A budget of $3.00 buys about **14 iterations**. On `claude-sonnet-5` the same eval costs about $1.19 a run, and $3.00 buys two. The owner's choice of the cost tier is what makes the loop possible.

## What the $3 does not cover

The fixer agent is not in that table. It reads a report and edits Go, and its own tokens are billed to the owner somewhere else. A frontier model doing real work over fourteen iterations plausibly costs **ten to forty dollars**, which is more than the loop budget by a wide margin.

Set a second budget for the fixer before the first unattended run. Nothing in this repository can measure it.

## Will it give useful information

The report will. The loop might.

The report is worth running on its own. It names every question the eval refused, the row it came from, the fault, and the evidence. That is the two hours of scoring, in about eleven cents. Run it after every gate run, whether or not the loop ever runs.

The loop is a stretch, for four reasons.

- The eval runs on the model that writes the questions. A model marks its own wording gently. `make eval-calibrate` measures how gently, for about thirty cents (OQ-26).
- The batch sweep already fixed sixteen defects that deterministic rules can find. What is left is judgment, which is where an automated score is weakest.
- A loop can only measure the 66 conversations we wrote. It can not tell us about a user we never imagined.
- A fixer that reads its own scorecard will tune the scorecard. The holdout split is the guard, and it is not a proof.

## The guards

Each one exists because something in this repository's history needed it.

| Guard | What it stops | Where |
|---|---|---|
| A branch of its own | A bad night reaching `main` or `pr-7` | `autotune.sh` |
| Frozen paths | The fixer editing its own scorer, the linter, or the test set | `autotune.sh` |
| Append-only decisions | The fixer rewriting the record | `autotune.sh` |
| Green build | A broken tree surviving an iteration | `autotune.sh` |
| The question counters | A run that scores well by asking less (gate run 7) | `internal/tune` |
| The premature check | A session that calls itself complete with a slot open (D-91) | `internal/tune` |
| The linter | A redundant or presumptuous question (D-115) | `internal/questions/lint.go` |
| The holdout split | The fixer learning the test set (D-134) | `internal/tune` |
| The owner questions | The fixer deciding what is not its to decide | `docs/owner-questions.md` |

## The holdout split

Every third conversation is held back. The fixer reads the failures of the other two thirds and never sees the holdout. The loop then reads the holdout ratio, and not the ratio it showed the fixer.

A ratio that falls on the tune split and stands still on the holdout means the fixer reworded what it was shown. The loop prints a warning when that happens, and it keeps reading the holdout.

## Accept and reject

An iteration is kept only when every one of these holds.

1. The linter finds nothing.
2. No session calls itself complete with a needed slot unanswered.
3. The agent asks at least 90 percent of the questions it asked before.
4. At least 90 percent as many questions close a slot.
5. The holdout ratio did not rise.

Anything else reverts the whole iteration with `git reset --hard`.

## Stop conditions

- The spend reaches the budget. The ledger counts both the gate and the eval.
- The holdout ratio reaches the target, which defaults to 5 percent (OQ-27).
- Three iterations in a row change nothing that holds.
- The iteration cap, which defaults to 20.

## What the owner does, run by run

The setup happens once.

1. Name the fixer in `AUTOTUNE_FIXER_CMD`. A headless agent needs its permission prompts turned off, because a prompt at two in the morning stops the loop until morning.
2. Set a token budget for that agent, wherever it bills. This repository can not see it.

Each run then costs the owner two things.

- **Before.** One command, and a look at the budget. The loop refuses to start on a dirty tree, on `main`, or without `AUTOTUNE_ALLOW_UNATTENDED=1`.
- **After.** Read `.local/tune/autotune.log`, read the branch, and merge it or delete it. Every accepted iteration is one commit with the ratio and the counts in its message.

Nothing is needed while it runs. That is the point of it.

## What unattended really means

The guards in the table above bound what a bad night does **inside this repository**. They do not sandbox the agent. A headless agent with its prompts turned off runs shell commands as the owner, and no rule in this repository stops that.

That is the real content of OQ-25. The question is not whether the loop can break the code, because a revert fixes that. The question is whether an unsupervised process may hold the owner's shell all night.

Three answers, in order of strength.

1. **The branch and no push.** The default. The loop commits to `auto-tune/<stamp>` and pushes nothing. Add `--push` to change that.
2. **A worktree of its own.** Run the loop in a git worktree under `.local/`, so the owner's checkout is never touched. Not built yet.
3. **A container.** The fixer runs with only the worktree mounted, and reaches only the two API hosts. Not built yet.

## What is still open

Three of the four answers came on 2026-08-26. D-135 lets the loop change the catalog inside an approved run. D-136 accepts the shared eval model, and every ratio it reports is a floor. D-137 sets the target at 5 percent.

OQ-25 is open. `scripts/autotune-fix.sh` ships no default agent until it closes.

## The recommendation

Run the report alone first.

```
GATE_OUT=docs/reference/pr7-question-gate-run14.md make questions-gate
EVAL_RUN=docs/reference/pr7-question-gate-run14.md make questions-eval
make eval-calibrate
```

That costs about fifty cents in total, it needs no new permission, and it answers OQ-26 and OQ-27 with data. Read the report. If it finds what the batch sweep found by hand, the loop is worth turning on. If it does not, the loop would only automate a scorer that can not see.
