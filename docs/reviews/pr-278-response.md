# Response to the review of pull request 278

The author answers the Codex record of `docs/reviews/pr-278.md`. That record read `Changes required` at head `1ccaf42`, with one finding.

## P2-1: A later marker can strand an earlier failed marker

- The result: full merit.
- The evidence: the refresh writes the cards of `cards.CompareVersions(previous, current)` into the marker of the new version. A newer marker thus holds only the cards that are new since the version before it. `LatestNewCards` returned the newest marker alone, and the pass returned when that marker had an end. An older marker that a stopped pass left never ran again.
- The correction: `cards.PendingNewCards` in `go/internal/cards/newcards.go` replaces `LatestNewCards`. It returns each marker with no ended pass, oldest first. `newCardsPass` in `go/cmd/worker/newcardspass.go` joins their cards in one pass under the newest version, so each user gets one push. A failure leaves each marker without an end. The roadmap entry of PR-120 says so.
- The Gitar review of the first correction `82b9089` found two faults. An older marker overwrote a deck that a newer marker wrote. Two pending markers sent two pushes, and the second pass replaced the cards of the first. So the join replaced the run in order. `needsWrite` in `go/internal/newcards/newcards.go` now skips a deck of the same or a newer version. A deck takes only the cards newer than its own version, and the cards it still shows, so a dismiss stays (D-1095).
- The regression check: `TestNewCardsPassOlderMarker` writes a pending marker and a newer ended marker. The pass runs the older one, and the deck takes its card. Two pending markers then run as one pass with one push. A mutation back to the newest marker alone failed the test. `TestPassJoinedMarkers` reads a deck of a newer marker, a dismissed deck, and a deck that shows the older cards. A mutation back to the equal version, and a mutation that opens every card, each failed it. `TestNewCardsMarker` reads the order of `PendingNewCards`. `make verify` passed.
