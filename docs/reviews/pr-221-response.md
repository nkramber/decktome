# Pull request 221: the answer to the review

The author answers each finding of `docs/reviews/pr-221.md` here, with the procedure of `.claude/skills/pr-review/references/answer-review.md`.

## P2-1: The session sync can erase concurrent usage

- Result: full merit.
- Evidence: at `04d6373`, `storeImportPower` of `go/internal/agentsvc/import.go` set `current.Usage` to the copy that `ReadImportBracket` read before the judge call. A turn that wrote the session after that read lost its usage. The loss also occurs with no version conflict, when the turn writes before the first `GetState`. D-447 says that the session spend holds each call.
- Correction: `ReadImportBracket` passes the report of the judge to `storeImportPower`. On each try, the report adds to the usage that `GetState` reads (D-447, D-870).
- Regression check: `TestReadImportBracketKeepsConcurrentUsage` of `go/internal/agentsvc/import_test.go` adds 7 calls and moves the version between the read and the write. On the old logic it read `calls = 2, want 9`, and it passes now. The session also holds the judged bracket.
