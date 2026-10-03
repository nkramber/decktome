# Response to the review of pull request 276

The author answers the Codex record of `docs/reviews/pr-276.md`. That record read `Changes required` at head `6ac5e7a`, with one finding.

## P2-1: Creation writes can move the timestamps backwards

- The result: full merit.
- The evidence: `Note` wrote `last_seen_at` and `last_creation_at` with a `Set` and no read. A creation at an earlier time that commits last wrote the earlier time. D-1093 makes each field the newest value, and `Touch` and `Seed` already compare with `After`.
- The correction: `Note` in `go/internal/users/users.go` runs in a transaction now. It reads the record, and it moves each time only to a newer value. It also does the copy of D-1094 on a version 1 record, so the second limit of D-1094 no longer applies. D-1094 now says so.
- The regression check: `TestALateNoteNeverMovesATimeBack` writes a chat at 20:00, then a deck and an upload at 19:59. It reads 20:00 in both fields. `TestNoteMovesARecordOfVersionOne` reads the copy on the creation path. Both tests fail on `6ac5e7a` and pass on the correction, against the Firestore emulator. The tests of `go/internal/agentsvc`, `go/internal/collectionsvc`, `go/internal/feedbacksvc`, and `go/cmd/users-backfill` pass under `-race`.
