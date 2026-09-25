# Response to the review of pull request 228

The author answers the Codex record of `docs/reviews/pr-228.md`. That record read `Changes required` at head `44d2fb0`, with one finding.

## P2-1: A slow failed snapshot run skips its alert

- The result: full merit.
- The evidence: `snapshotNotice` read the alert window from the clock at the end of the run. A run that starts at 18:00:05 and fails at 18:40:05 is past minute 15, so it sent no notice. The worker gives a refresh up to 90 minutes.
- The correction: `run` records the start of the run, and `snapshotNotice` reads the window from it. The snapshot age still reads the end of the run. `go/cmd/worker/main.go` and D-911 hold the rule.
- The regression check: `TestASlowFailedRunStillAlerts` starts a run at 18:00:05 and fails it 40 minutes later. It requires one notice with the age and the error, and no notice for a run that started outside the window. With the old rule the test failed: `sent false`. It passes now, and `TestSnapshotNotice` still passes.
