# Response to the review of pull request 229

The author answers the Codex record of `docs/reviews/pr-229.md`. That record read `Changes required` at head `4d51d4c`, with one finding.

## P2-1: Delete can remove a session after a build acquires its lease

- The result: full merit.
- The evidence: `DeleteSession` read the lease with `Leased`, and `Repo.Delete` then removed the session and the lease with no lease check. A scratch emulator test on `4d51d4c` took a live lease, then called `Repo.Delete`. The delete removed the session: "the delete removed the session under a live lease". The deck delete of `go/internal/decksvc/service.go` had the same gap, and it removed the chat of a running build on any instance.
- The correction: `Delete` of the session store takes the clock, reads the lease in its own transaction, and returns `sessions.ErrLeased` while a lease holds. `DeleteSession` and the deck delete map it to Aborted, and remove nothing. The first read of `Leased` in `DeleteSession` goes, because the transaction holds the check. D-922 holds the rule.
- The regression check: `TestEmulatorBuildLease` now requires `ErrLeased` from a delete under a live lease, and the session after it. It passes on the emulator. `TestADeleteRefusesALeaseTakenAfterItsRead` takes the lease after the reads of `DeleteSession` and before its delete, and requires Aborted, the session, and the lease. The new case of `TestDeleteDeck` requires Aborted, no deleted deck, and the chat. `go test ./internal/agentsvc/ ./internal/decksvc/ ./cmd/api/` passes.

## The repeat review of `a8efdf3`: Blocked on the whole gate runs of D-921

- The result: no merit.
- The evidence: D-921 is the rule of the feedback fix cycle. The cycle opens a pull request, and its body asks for a whole run of each gate that its fix touches. PR-79 changes the text of that step, and it runs no cycle. It changes no prompt, no gate, no role, and no code of a deck build or a question turn, so no gate row can move. `make verify` runs `eval-check` on the committed baselines, and it passes. The roadmap entry of PR-79 names no paid run in its gate.
- The correction: D-921 and the gate of the fix cycle now say that the rule binds the pull request of a cycle alone. `docs/design-roadmap.md` holds the gate.
- The regression check: `make ste-check` and `make ref-check` pass on the new text.

## P2-2: Failed triage marks its harvest as processed

- The result: full merit.
- The evidence: `go/cmd/feedback-triage/main.go` called `MarkTriaged` on each input path after `applyCases`, with no read of `Result.Err`. A judge failure on one verdict marked its harvest read, so no later triage read that verdict again. A harvest that stays pending as a whole writes its good cases again on the next read, so the record names each verdict.
- The correction: `triage.Settle` records each verdict that has no failure and no open judge need. It marks a harvest read only when each of its verdicts has a line. `triage.Unapplied` drops each verdict that an earlier live run applied, so a new read writes no case twice. D-924 holds the rule.
- The regression check: `TestSettleKeepsAFailedVerdictPending` gives one good verdict in one harvest, and one good and one failed verdict in another. It requires the second harvest to stay pending, and the next read to hold the failed verdict alone. `go test ./internal/triage/ ./cmd/feedback-triage/` passes.
