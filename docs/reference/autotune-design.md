# The automated tuning loop

Written 2026-08-26. Updated the same day, after the owner answered every blocking question (D-135 to D-138).

`docs/design-roadmap.md` holds PR-7B, which is the authority on what this lane is for, which evals run, and what its gate is. This file is the authority on how the loop runs: the guards, the accept rules, the cost, and what the owner does run by run. Read the roadmap entry first.

The three evals of the PR-7B gate started on 2026-08-26. The loop itself has not run.

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

| Item | Per iteration | Wall clock |
|---|---|---|
| Gate run, 100 conversations | about $0.14 | about 18 minutes |
| Eval, about 98 calls on the cost tier | about $0.10 | about 13 minutes |
| Total inside the budget | about $0.24 | about 31 minutes |

Measured on 2026-08-26 at 66 conversations: the gate cost $0.0964 over 737 seconds, and the eval cost $0.0645 over 494 seconds. D-145 raised the set to 100, and the figures above scale by the message count.

A budget of $3.00 buys about **12 iterations**. On `claude-sonnet-5` the same eval costs about $1.80 a run, and $3.00 buys one. The owner's choice of the cost tier is what makes the loop possible.

Wall clock now binds before the budget does. Twelve iterations take about six hours of provider time, and the fixer agent adds its own on top. One night fits eight to ten iterations, not twelve.

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

## Where the commits go

Three layers, and the middle one is the part that is easy to miss (D-142).

| Layer | What it holds | Who moves it |
|---|---|---|
| `main` | The merged work. | The owner, through a pull request. |
| `pr-7c` | Every change the loop has written and the owner has kept. | The owner, by fast-forward, after a review. |
| `auto-tune/<stamp>` | One night. | The loop. |

`--base pr-7c` names the container. The run switches to it, cuts the night's branch off it, and pushes nothing.

Without a container, every night starts from the same commit. The second night never sees the first night's accepted work, so both nights rewrite the same catalog rows from the same base, and the owner is left merging two branches that disagree about the same file. The loop warns when the base is `main` for that reason.

The container also gives the work one shape. Ten accepted iterations over three nights become one pull request from `pr-7c`, not ten branches.

A night that goes wrong costs nothing. Delete `auto-tune/<stamp>` and the container is untouched.

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

Three answers, in order of strength. The owner chose the first (D-138).

1. **The branch and no push.** ✅ chosen. The loop commits to `auto-tune/<stamp>` and pushes nothing. `--push` is the opt-in.
2. **A worktree of its own.** Run the loop in a git worktree under `.local/`, so the owner's checkout is never touched. Not built.
3. **A container.** The fixer runs with only the worktree mounted, and reaches only the two API hosts. Not built.

One step is left, and it is the owner's. `scripts/autotune-fix.sh` ships no default agent. The owner names the fixer in `AUTOTUNE_FIXER_CMD` and caps its tokens.

## What eval calibration is

The eval role runs on the model that also writes the questions. A model marks its own wording gently, and the owner accepted that for cost (D-136). Calibration measures how gently.

`make eval-calibrate` scores the same 12 conversations twice. The first pass uses the cost tier. The second pass sets `LLM_EVAL_PROVIDER` and `LLM_EVAL_MODEL`, so a stronger model reads the same questions. `cmd/tune-check -agree` then compares the two, question by question, and reports four things.

- How many questions both models scored.
- How often the two gave the same verdict.
- How many questions each one refused.
- A warning when the cost tier refused fewer.

The last line is the one that matters. A cost-tier eval that refuses four questions where the stronger model refuses twelve is not measuring the agent. It is reporting a floor, and the real number is somewhere above it.

The pass costs about 25 cents, and the stronger model is nearly all of it. Run it once against the first gate document, and again after any change to the eval prompt.

OQ-39 holds what the owner does with the number. No floor is set, and none should be guessed.

## What is still open

Every blocking question is answered. D-135 lets the loop change the catalog inside an approved run. D-136 accepts the shared eval model. D-137 sets the target at 5 percent. D-138 gives the loop a branch of its own and no push.

OQ-39 is open, and the first calibration measures it.

## The recommendation

Run the report alone first.

```
GATE_OUT=docs/reference/pr7-question-gate-run14.md make questions-gate
EVAL_RUN=docs/reference/pr7-question-gate-run14.md make questions-eval
make eval-calibrate
```

That costs about fifty cents in total, it needs no new permission, and it answers OQ-26 and OQ-27 with data. Read the report. If it finds what the batch sweep found by hand, the loop is worth turning on. If it does not, the loop would only automate a scorer that can not see.
