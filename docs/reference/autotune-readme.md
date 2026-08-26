# Autotune, how to run it

The loop reads the eval report, lets a fixer agent edit the code, runs a
new gate and a new eval, and keeps the change only when the score
improves. It commits to a branch of its own and it pushes nothing.

`autotune-design.md` holds the reasons and the honest limits. This file
holds the commands.

## Set it up once

1. Add one line to `.env`:

   ```
   AUTOTUNE_FIXER_CMD=claude -p --permission-mode bypassPermissions
   ```

   No dollar cap goes here. The fixer bills against the monthly plan
   (D-159). The `--budget` flag below is a separate thing.

2. Keep a container branch, and name it with `--base`. `pr-7c` is the one
   today. Every night cuts a working branch off it, and you fast-forward
   it after a review. Without a container each night starts from the same
   commit, so two nights rewrite the same rows instead of adding up
   (D-142).

3. Commit every change. The loop refuses a dirty working tree. It does not
   need a clean history, and a merge commit is fine. Commit an edit to
   `scripts/autotune.sh` before a run, because that path is frozen.

4. Make a baseline. Run one gate and one eval on the code you are about to
   tune, keep the JSON summary, and pass it with `--baseline`. Without one
   the loop pays for its own first run.

   A baseline holds only while the code it measured holds. A later change
   to Go, to the proto, or to a JSON file the app reads makes it stale,
   and the loop then credits the fixer with work it did not do. A change
   to a document does not.

## Run it

One iteration, to see what the fixer does:

```
AUTOTUNE_ALLOW_UNATTENDED=1 scripts/autotune.sh \
  --base pr-7c --baseline .local/tune/BASELINE.json --budget 3.00 --max 1
```

A full night. Drop `--max` and the default of 20 applies, which the
budget stops first:

```
AUTOTUNE_ALLOW_UNATTENDED=1 scripts/autotune.sh \
  --base pr-7c --baseline .local/tune/BASELINE.json --budget 3.00
```

## Stop it

```
touch .local/tune/STOP
```

The iteration that runs finishes and commits. The loop then ends before
the next one starts, and it removes the file.

Ctrl-C also stops the loop. It abandons the iteration in flight and wastes
that spend, up to about $0.24.

## After an interrupt

Ctrl-C leaves the tree dirty, and the loop refuses to start again:

```
git reset --hard && git clean -fd
git switch pr-7c
```

Every finished iteration is already committed on the `auto-tune` branch.

## Read the result

| What | Where |
|---|---|
| The log of the night | `.local/tune/autotune.log` |
| Each gate document | `docs/reference/pr7-question-gate-auto-NNN.md` |
| Each eval report | `docs/reference/pr7-question-eval-auto-NNN.md` |
| Each score summary | `.local/tune/auto-NNN.json` |
| The night's commits | branch `auto-tune/<stamp>` |

A rejected iteration leaves no document. The loop cleans the tree, so read
the log for what it tried.

Keep a night by a fast-forward of `pr-7c`. Delete the branch to drop it.

## What it costs

One iteration is about $0.24 and about 33 minutes. The gate and the eval
run on `gpt-5.6-luna` and they bill per token.

A budget of $3.00 buys about 12 iterations, which is about six hours. Wall
clock therefore binds before the budget does. One night fits eight to ten.

The fixer is not in that number. It bills against the monthly plan.

## Every flag

| Flag | Default | What it does |
|---|---|---|
| `--base` | current branch | The container branch the night starts from. |
| `--baseline` | none | A scored JSON to compare against. |
| `--budget` | 3.00 | Dollars for the gate and the eval. |
| `--max` | 20 | Most iterations to run. |
| `--target` | 0.05 | Stop when the holdout ratio reaches this. |
| `--push` | off | Push the branch as the night goes. |
| `--dry-run` | off | Print the plan and call no provider. |

## Where it stops by itself

- The budget is spent.
- The iteration count reaches `--max`.
- Three iterations in a row change nothing that holds.
- The holdout ratio reaches the target.

## What it can not do

The loop reverts any iteration that breaks the build, that scores worse,
or that touches a frozen path. These paths are the measurement, and the
fixer may not edit them:

`go/cmd/questions-eval`, `go/cmd/tune-check`, `go/internal/tune`,
`go/internal/questions/lint.go`, `go/cmd/questions-gate/conversations.json`,
`scripts/autotune.sh`, `scripts/autotune-fix.sh`, `docs/owner-questions.md`.

The fixer also reads every open question of `docs/owner-questions.md`, and
it decides none of them.

Nothing reaches `main` or `pr-7c`. The loop commits to `auto-tune/<stamp>`
and it pushes only with `--push`.
