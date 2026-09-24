# Pull request 220: the answer to the review

The author answers each finding of `docs/reviews/pr-220.md` here, with the procedure of `.claude/skills/pr-review/references/answer-review.md`.

## P2-1: The parser silently drops lines after line 1000

- Result: full merit.
- Evidence: at `a32ac85`, `Parse` of `go/internal/decklist/decklist.go` stopped at line 1001 with no error. A list of 1001 lines of 9 bytes each sits far under the cap of 128 KiB. A test on the old logic read the first 1000 lines back and no error, so the trigger reproduces. D-846 says that the import reports each line that it skips.
- Correction: `Parse` returns `ErrTooManyLines` at the first line past `maxLines`, and the whole list fails. `ImportDeck` answers InvalidArgument, and it stores no deck and no session (D-846).
- Regression check: `TestParseRefusesALongList` of `go/internal/decklist/decklist_test.go` fails on the old logic and passes now. A list of exactly 1000 lines still parses. `TestImportRefusesALongList` of `go/internal/agentsvc/import_test.go` proves that the import stores no deck, no session, and no read. `make verify` passes.
