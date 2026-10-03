# Response to the review of pull request 277

The author answers the Codex record of `docs/reviews/pr-277.md`. That record read `Changes required` at head `a5583df`, with two findings. The owner chose the answer of each finding.

## P2-1: The request cap is independent on each API instance

- The result: full merit for the fact, and the owner amended the contract.
- The evidence: `New` in `go/internal/agentsvc/service.go` makes a `ratelimit.Limiter` in memory. The API runs at most 3 instances (`docs/setup-gcp.md`). So one user can make up to 30 reads a minute. Each other limiter of the app counts on each instance too, such as the public reads of D-315 and the feedback notices.
- The correction: the owner chose to keep the limiter and amend D-1100 (D-1106). The cap is 10 reads a minute on each API instance. `FetchesPerMinute`, the field comment, and the roadmap entry now say so.
- The regression check: none, because the code does not change. `TestFetchDeckListLimitsEachUser` still proves the cap of one instance.

## P2-2: An import can claim a deck that it never read

- The result: full merit.
- The evidence: `sourceURL` checked only the form of the link. A direct `ImportDeck` call with a valid Archidekt link and other text stored the link.
- The correction: the owner chose a record of each read (D-1107). `go/internal/deckreads` keeps the reads of each user for one hour in `deck_reads/<uid>`: the deck id, a hash of the text, and the time. `FetchDeckList` records each read. `ImportDeck` keeps a link only after a read of the same deck and text by the same user. A fault of the store answers `Unavailable`, for the read and for the import. The API wires the Firestore store.
- The regression check: `TestImportKeepsOnlyTheRemadeLink` refuses a link with no read, a text other than the read, and a deck other than the read. It fails when the check of the read is off. `TestAReadLogFaultFailsClosed` reads the fault path. `TestMatchesNeedsTheDeckAndTheText` reads the hour limit. `TestReadRoundTrip` passes against the Firestore emulator through `make store-check`.

## The other changes of this round

- The branch merged `main` after #276, and the owner renamed the item PR-121 (D-1108).
