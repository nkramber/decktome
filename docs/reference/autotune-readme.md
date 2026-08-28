# Autotune, how to run it

The loop reads the eval report and the lessons file, and lets a fixer
agent edit the code. Then it runs a new gate and a new eval. It keeps each
change the fixer made only when its own rows got better (D-181). It
commits to a branch of its own and it pushes nothing.

`autotune-design.md` holds the reasons and the honest limits. This file
holds the commands.

## Set it up once

1. Add one line to `.env`:

   ```
   AUTOTUNE_FIXER_CMD="claude -p --permission-mode bypassPermissions"
   ```

   Keep the quotation marks. Bash reads `NAME=value more words` as a
   command to run with NAME set, so an unquoted value that holds a space
   breaks every target that reads `.env`.

   The marks do not reach the fixer: `autotune-fix.sh` splits the value
   into words when it runs it. A value with a quoted argument of its own
   does not work.

   No dollar cap goes here. The fixer bills against the monthly plan
   (D-159). The `--budget` flag is a separate thing, and it pays for the
   gate and the eval.

   The loop withholds `ANTHROPIC_API_KEY` and `OPENAI_API_KEY` from the fixer (D-161). Claude Code reads the first in preference to a claude.ai login. Without this the fixer bills per token to the key and not to the plan. The second key is withheld so the fixer can not pay for a gate run the loop does not count.

   Every other name in `.env` does reach the fixer, so keep other secrets
   out of that file.

   The fixer reads `.local/tune/<label>.fixer.json`, which is the eval
   summary with every holdout verdict removed (D-271). Its checkout never
   holds the holdout.

2. Keep a container branch, and name it with `--base`. The owner names
   it, and `pr-7c` was the one until it merged into `main` on 2026-08-28.
   Every night cuts a working branch off it, and you fast-forward it after
   a review. Without a container each night starts from the same commit,
   so two nights rewrite the same rows and never add up (D-142).

3. Commit every change. The loop refuses a dirty tree and an untracked
   file, because it commits with `git add -A` and an untracked file would
   ride into a loop commit under your name (D-184). It does not need a
   clean history, and a merge commit is fine. Commit an edit to
   `scripts/autotune.sh` before a run, because that path is frozen.

4. Make a baseline. Run one gate and one eval on the code you are about to
   tune, keep both outputs, and pass the JSON with `--baseline`. Without
   one the loop pays for its own first run.

   The loop needs the eval report as well as the JSON summary. The checker reads the summary and the fixer reads the report. The loop finds the report by the name of the summary, so `.local/tune/run18.json` pairs with `docs/reference/pr7-question-eval-run18.md`. Keep that pairing, or name the report with `--baseline-doc`.

   A baseline the loop can not read stops the run. It never buys a new
   one in its place (D-184).

   A baseline holds only while the code it measured holds. A later change
   to Go, to the proto, or to a JSON file the app reads makes it stale.
   The loop then credits the fixer with work it did not do. A change to a
   document does not.

   A partial eval is no baseline. The summary carries `partial: true`
   when the eval budget stopped it, and the loop refuses it (D-271).

## Run it

One iteration, to see what the fixer does:

```
AUTOTUNE_ALLOW_UNATTENDED=1 scripts/autotune.sh \
  --base <container> --baseline .local/tune/BASELINE.json --budget 3.00 --max 1
```

A full night. Drop `--max` and the default of 20 applies, which the
budget stops first:

```
AUTOTUNE_ALLOW_UNATTENDED=1 scripts/autotune.sh \
  --base <container> --baseline .local/tune/BASELINE.json --budget 3.00
```

## Run it without a baseline

Drop `--baseline` and the loop measures its own. It runs one gate and one
eval first, calls that the baseline, commits it as `v0.0`, and then starts
the iterations.

```
AUTOTUNE_ALLOW_UNATTENDED=1 scripts/autotune.sh \
  --base <container> --budget 3.00 --max 1
```

The owner keeps a shell function named `autotune` in `.zshrc`. It takes
the branch, the budget, and the iteration count, and it passes every
other flag through to the script:

```
autotune --branch <container> --budget 3.00 --max 1
autotune --branch <container> --budget 3.00 --max 20 --baseline .local/tune/BASELINE.json
```

`--max` counts tuning iterations only. The baseline is not one of them, so
this command runs two gate and eval pairs: the baseline, then iteration 1.
It costs about $0.50 and takes about an hour.

Use it when the code has changed since the last scored run. A baseline
that measured other code makes the loop credit the fixer with work it did
not do.

## What one iteration does

1. The fixer reads the last eval report, the owner questions, and the lessons file. It commits each change on its own, with a `Rows` trailer and a `Hypothesis` trailer (D-181). A commit with no `Rows` trailer is rejected (D-271).
2. The loop checks the frozen paths, strips any attribution line from the commit messages, and builds the tree.
3. The gate and the eval run. About $0.25 and 33 to 35 minutes. A partial eval stops the loop, because a partial run decides nothing.
4. `tune-check` compares the run with the best accepted run of the night, question by question (D-271). It charges every moved question to the change that declared its row.
5. Every change kept: one commit `v0.N`. Every change dropped: the tree reverts. Some kept: the loop cherry-picks the kept commits and runs the gate again on the conversations their rows touch. It folds that eval into the baseline, and commits `v0.N` if the merged run holds.
6. The lesson is written and committed, whatever the outcome (D-182).

## Nothing is ever overwritten

Every document carries the run stamp, so two runs never collide. A
document reads `pr7-question-gate-20260826-191225-001.md`, where the first
part is the run and the last is the iteration (D-177). A partial
re-measure reads `-001-kept`, and its fold reads `-001-kept-merged`.

The loop refuses to write over a file that exists, and it names the file.
In the baseline path it stops. Inside an iteration it reverts that
iteration and goes on to the next.

## Stop it

```
touch .local/tune/STOP
```

The iteration that runs finishes and commits. The loop then ends before
the next one starts, and it removes the file.

Ctrl-C also stops the loop. It abandons the iteration in flight and wastes
that spend, up to about $0.25.

## After an interrupt

Ctrl-C leaves the tree dirty, and the loop refuses to start again. Look
before you clean, because `git clean` removes every untracked file:

```
git status
git clean -n
git reset --hard && git clean -fd
git switch <container>
```

Every finished iteration is already committed on the `auto-tune` branch.

## Read the result

| What | Where |
|---|---|
| The log of the night | `.local/tune/autotune.log` |
| What the loop learned | `docs/reference/autotune-lessons.md` |
| Each gate document | `docs/reference/pr7-question-gate-<stamp>-NNN.md` |
| Each eval report | `docs/reference/pr7-question-eval-<stamp>-NNN.md` |
| Each score summary | `.local/tune/<stamp>-NNN.json` |
| The fixer's changes, one commit each | `.local/tune/<stamp>-NNN-changes.json` |
| The checker's verdict | `.local/tune/<stamp>-NNN-verdict.txt` |
| A rejected or dropped change | `.local/tune/rejected/<stamp>-NNN/` |
| The spend | `.local/tune/ledger-<stamp>.txt` |
| The night's commits | branch `auto-tune/<stamp>` |

A rejected iteration leaves no document under `docs/reference/`. Its
diff, its commits, its documents, and its verdict sit under
`.local/tune/rejected/` (D-180), and its lesson sits in the lessons file.

## Keep or drop a night

Find the branch the night wrote:

```
git branch --list 'auto-tune/*'
```

Read what it did before you decide:

```
git log --oneline <container>..auto-tune/<stamp>
git diff <container>..auto-tune/<stamp>
```

Keep it. The night was cut from the container, so the merge is a fast-forward:

```
git switch <container>
git merge --ff-only auto-tune/<stamp>
git branch -d auto-tune/<stamp>
```

`--ff-only` fails when the container moved after the night started. That failure
is a warning and not a fault. Read both branches before you merge them
any other way.

Drop it:

```
git branch -D auto-tune/<stamp>
```

Keep one commit of a night, and drop the rest, with `git cherry-pick`.
Each accepted iteration is one commit named `v0.1`, `v0.2`, and so on,
and each lesson is one commit named `lesson: <label>`. Keep the lessons
even when you drop the rest. They are what the next night learns from.

## What it costs

One iteration is about $0.25 and 33 to 35 minutes. The gate and the eval
run on `gpt-5.6-luna` and they bill per token. A partial keep adds a gate
and an eval over a few conversations, which is cents.

A budget of $3.00 buys about 12 iterations, which is about six hours. Wall
clock therefore binds before the budget does. One night fits eight to ten.

The budget is checked before each iteration, so the spend can end one
iteration past it.

The fixer is not in that number. It bills against the monthly plan.

## Every flag

| Flag | Default | What it does |
|---|---|---|
| `--base` | current branch | The container branch the night starts from. |
| `--baseline` | none | A scored JSON to compare against. |
| `--baseline-doc` | beside the JSON | The eval report the fixer reads. |
| `--budget` | 3.00 | Dollars for the gate and the eval. |
| `--eval-budget` | 0.50 | Dollars one eval may spend before it stops. |
| `--max` | 20 | Most iterations to run. |
| `--target` | 0.05 | Stop when the holdout ratio reaches this. |
| `--noise` | 9 | Bad questions two runs of the same code may differ by. Empty reads `tune.DefaultNoise` through `tune-check -print-noise` (D-230, D-258). |
| `--branch` | `auto-tune` | The prefix of the night's branch. |
| `--push` | off | Push the branch as the night goes. `--no-push` turns it off again. |
| `--dry-run` | off | Run the fixer and the build, then stop before the gate. It calls no provider and commits nothing. |

## Where it stops by itself

- The budget is spent.
- The iteration count reaches `--max`.
- Three iterations in a row change nothing that holds.
- The holdout ratio reaches the target. The stop reads the holdout, not the whole set (D-271).
- An eval comes back partial.

## What it can not do

The loop reverts any change that breaks the build, that scores worse on
its own rows, or that touches a frozen path. These paths are the
measurement and the memory, and the fixer may not edit them:

`go/cmd/questions-eval`, `go/cmd/tune-check`, `go/internal/tune`, `go/internal/questions/lint.go`, `go/internal/questions/lint_test.go`, `go/internal/questions/metrics.go`, `go/cmd/questions-gate/main.go`, `go/cmd/questions-gate/conversations.json`, `go/internal/llm/roles.json`, `go/internal/llm/prices.json`, `scripts/autotune.sh`, `scripts/autotune-fix.sh`, `docs/owner-questions.md`, `docs/reference/autotune-fixer-prompt.md`, `docs/reference/autotune-lessons.md`. The gate command, its metrics, the role table, and the price table are the scorecard, and they joined the list on 2026-08-28 (D-271).

The fixer also reads every open question of `docs/owner-questions.md`, and
it decides none of them.

The loop can not tell a change from the judge's noise when the change moves fewer bad questions than the margin, which is 9 (D-230). It can not measure a user the conversations never imagined. Read the count from `go/cmd/questions-gate/conversations.json`. It is 104 on 2026-08-28.

Nothing reaches `main` or the container. The loop commits to `auto-tune/<stamp>`
and it pushes only with `--push`.
