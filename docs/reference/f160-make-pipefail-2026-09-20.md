# F-160: a failed paid target must fail its make target

Date: 2026-09-20. Sources: `Makefile` of this repository, the GNU Make of this Mac, and the GNU Make manual.

## The fault

`Makefile` sets `.SHELLFLAGS := -o pipefail -c`. GNU Make 3.82 added that variable. GNU Make 3.81 ignores it, and this Mac runs GNU Make 3.81. So a recipe of this machine runs under `bash -c` with no pipefail. A pipeline then reports the status of its last command alone.

Three paid targets ended with `| tee`, so `tee` gave each one its exit code. `tee` succeeds after a command that fails. The first live run of `make api-build` on 2026-09-20 failed with "card database not loaded yet", and the target exited 0. PR-61 put `set -o pipefail` in the recipe of `make api-build`. This pull request fixes the other targets, and it adds a check that holds the rule.

## The measurement

A scratch makefile carries the two `SHELL` lines of this repository. Each target below ran under `make` of this Mac, GNU Make 3.81.

| Target | Recipe | Exit code |
|---|---|---|
| `a` | `@false \| tee out.txt` | 0 |
| `b` | `@set -o pipefail; false \| tee out.txt` | 2 |

Target `a` reads the fault, and target `b` reads the fix. A recipe that sets pipefail itself needs no `.SHELLFLAGS`, so it reads the same on GNU Make 3.81 and on GNU Make 4. The `verify:shell` job runs on `ubuntu-latest`, which carries GNU Make 4. That job now runs `make pipefail-check`, so each pull request proves the rule on both make versions.

## Every pipeline of the Makefile

`grep -n "| tee" Makefile` read four recipes. A wider scan of every recipe line read two more pipelines.

| Recipe line | Target | State before | State now |
|---|---|---|---|
| `\| tee $(CHAT_PROBE_OUT)` | `make chat-probe` | the pipe hid the exit code | `set -o pipefail` |
| `\| tee $(GENERATE_PROBE_OUT)` | `make generate-probe` | the pipe hid the exit code | `set -o pipefail` |
| `\| tee $(SUMMARY_JUDGE_OUT)` | `make summary-judge` | the pipe hid the exit code | `set -o pipefail` |
| `\| tee $(API_BUILD_OUT)` | `make api-build` | correct through PR-61 | `set -o pipefail` |
| `\| tail -1` | `make cover` | the pipe hid the exit code | `set -o pipefail` |
| `\| awk` | `make help` | no fault | `pipefail-ok` marker |

CAUTION: the F-160 entry of `docs/design-roadmap.md` named `make deck-gate` and `make questions-gate` as two targets with the fault. That claim is wrong, and this pull request corrects the entry. Each gate target writes its document with `> $(OUT)` and pipes to nothing. A redirect keeps the exit code of the command, so both gate targets always failed correctly. The three probe targets and the judge target held the fault.

## The proof of each fix

Each run below overrode `GO` with `false`, so the command before the `tee` failed at once. No run called a provider, and no run cost money. Each command also wrote its output to a scratch path outside this repository.

| Command | Exit code |
|---|---|
| `make chat-probe GO=false` | 2 |
| `make generate-probe GO=false` | 2 |
| `make summary-judge GO=false` | 2 |

Each target reported `make: *** [<target>] Error 1`. Before this pull request each one exited 0.

## The rule and the check

The rule: each recipe line that holds a pipeline sets pipefail itself. The words `set -o pipefail;` come first, before the pipeline.

A recipe whose exit code carries no fault takes an exemption. It carries a make comment line directly above it, and the comment starts with `pipefail-ok:` and gives the reason. `make help` is the one exemption today. Its `grep` reads the target list, and an empty list is no fault.

`make pipefail-check` holds the rule. It runs `docs/tools/pipefail_check.py`, and `make lint` and the `verify:shell` job both run it. The check reads four faults:

- `Makefile` sets no `SHELL`, so a recipe runs under a shell with no pipefail.
- A recipe line holds a pipeline, and it neither sets pipefail nor carries the marker.
- A marker names a recipe line that holds no pipeline.
- The `make` of this machine does not fail a pipeline that sets pipefail.

The last fault is a live proof. The check writes a fixture makefile with the `SHELL` lines of the real `Makefile`. Then it runs three targets of that fixture through the `make` of this machine. One target must fail, one target must pass, and the third target reads whether this `make` honors `.SHELLFLAGS`. The report names that answer, so a reader sees which make version the machine runs.

The check reads a quoted pipe as data, and never as a pipeline. So the `-run 'TestOne|TestTwo'` argument of a `go test` call takes no marker. `docs/tools/test_pipefail_check.py` holds 16 tests of the check, and `make lifecycle-check` runs them.

## The limits

- The check reads `Makefile` alone. It reads no script of `scripts/`.
- A read of 2026-09-20 found `-o pipefail` in `scripts/autotune.sh`, `scripts/feedback-loop.sh`, `scripts/feedback-review.sh`, `scripts/gcs-check.sh`, `scripts/read-session.sh`, and `scripts/where.sh`. The other five scripts set `-u` alone. No pipe of those five carries a fault. Two kill a port, one reads a version, one prints a diff, and one holds a quoted pipe. `shellcheck` reads every script on each pull request.
- CAUTION: `tee` creates its output file before the command writes a byte. So a failed run of a probe target leaves an empty document, and the guard of D-65 then refuses the same output name. Delete the empty document, or name a new one.
- The three paid loop targets print instructions and run no script, so the pipelines of the loop scripts stay outside this rule.
