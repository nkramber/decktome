# Response to the review of pull request 278

The author answers the Codex record of `docs/reviews/pr-278.md`. That record read `Changes required` at head `1ccaf42`, with one finding.

## P2-1: A later marker can strand an earlier failed marker

- The result: full merit.
- The evidence: the refresh writes the cards of `cards.CompareVersions(previous, current)` into the marker of the new version. A newer marker thus holds only the cards that are new since the version before it. `LatestNewCards` returned the newest marker alone, and the pass returned when that marker had an end. An older marker that a stopped pass left never ran again.
- The correction: `cards.PendingNewCards` in `go/internal/cards/newcards.go` replaces `LatestNewCards`. It returns each marker with no ended pass, oldest first. `newCardsPass` in `go/cmd/worker/newcardspass.go` runs a pass for each one in that order, and it stops at the first failure. The next run starts again at that marker. The roadmap entry of PR-120 says so.
- The regression check: `TestNewCardsPassOlderMarker` writes a pending marker and a newer ended marker. The pass runs the older one, and the deck takes its card. Two pending markers then run in order, and the newer cards replace the older ones (D-1095). A mutation back to the newest marker alone failed the test. `TestNewCardsMarker` reads the order of `PendingNewCards`. `make verify` passed.
