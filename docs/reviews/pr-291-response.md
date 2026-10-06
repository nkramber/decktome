# The author response to the review of #291

Author provider: Claude Code.

## P1-1: A session can erase the ledger and restart the budget

Result: full merit.

Evidence: `Start` read the spend file of the bundle alone, and `SumFile` read an absent file as no spend. The profile lets a session delete any file of its run folder. The log check of D-1170 ran only after the session ended.

The owner chose the correction from two options (D-1171). Option A gave a paid command read access to the session logs. A background run reaches the log only when the session reads its output, so that source stays incomplete. Option B is an append-only ledger outside the run folder.

Correction:

- `scripts/live-evals.sh`: `make_ledger` makes `$RUNS/<deck>/ledger/spend.jsonl` before each session and sets its `uappnd` flag. The session gets the path in `LIVE_EVAL_LEDGER`.
- `scripts/live-evals/sandbox.sb`: the new `LEDGER` parameter. The profile allows a read and `file-write-data` on that one file, and no other write.
- `go/internal/livespend/livespend.go`: `Start` refuses a run with no ledger. It sums the ledger and the spend file of the bundle, and it appends each line to both.
- `go/cmd/live-evals/spend.go` and `spent()`: the end sum reads the ledger beside the logs and the bundle.
- Each tick probes the rule under the profile before any session. An append must pass, and a truncate and a change of the flag must fail. A miss stops the tick with one notice.

Regression checks:

- `TestStartReadsTheLedgerAfterTheBundleIsGone`: with no bundle file and a ledger at $3.10, `Start` refuses. With a ledger at $1.10, `Start` starts.
- `TestStartRefusesWithNoLedger`: no variable, or no file, refuses the run.
- Under the real profile, with the ledger on the volume, an append and a read passed. A truncate, `rm`, a rename of the file or its folder, `chflags nouappnd`, `chmod`, and a new file beside it failed. A write in the run folder passed.
- `make_ledger` set `uappnd`, and a second call kept the content.

## P1-2: A forged end line can let another paid command start

Result: full merit.

Evidence: `Start` called `SumFiles`, which read the ledger and the spend file with no log. So a forged end line in the ledger closed a killed run, and the next start read no spend. A run whose stderr did not reach a log had the same gap in the end sum. Such a run got only the note "absent from the logs".

The owner chose the correction from three options (D-1173). Option A read the logs at the start alone. Option B also counts each run with no end line in a log as unmeasured. Option C accepted the risk.

Correction:

- `go/internal/livespend/livespend.go`: `Start` refuses a run with no `LIVE_EVAL_LOGS`. It reads the session logs, the ledger, and the spend file with `Sum(budget, true)`.
- The same file: with the logs read, the sum marks a run with no end line in a log unmeasured. A measured end in a file does not change that.
- `AddLogs` lists the folder with `os.ReadDir`. A folder that does not list is an error, so a refusal of the profile starts no paid run.
- `scripts/live-evals/sandbox.sb`: the new `LOGS` and `LOGS_RX` parameters. The profile allows a list of the logs folder and the read of `session-<n>.log`, and no other file.
- `scripts/live-evals.sh`: the session gets `LIVE_EVAL_LOGS`. Each tick probes the rule: the list and the log read pass, and the read of `context.json` fails.
- `scripts/live-evals/eval-prompt.md`: run each paid target in the foreground, and never redirect its stderr. The total command reads the logs too.

Regression checks:

- `TestStartRefusesAForgedLedgerEnd`: a forged ledger end with the start in a log, or with no log line, refuses the run. The same run with its end in a log starts. With `Start` set back to the sum without logs, the test fails.
- `TestStartRefusesWithNoLogs`: no variable, or no folder, refuses the run.
- `TestNewClientWritesTheSpendOfALiveEval`: the second run refuses until the log holds the end line of the first run.
- Under the real profile, with the logs on the volume, `live-evals spend` read a logged run as measured at $0.004. With the folder list refused, it stopped with an error.
- Under the real profile, `session-12.log` read, and `context.json` and `session-1.logx` did not. An escaped dot in the path matched only a dot.

Residual risk: a session can still echo a marked end line into a tool result, and the log sum takes it. No file inside the sandbox can prove an end line real. `docs/reference/paid-targets.md` names this risk.
