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
