# Response to the review of pull request 261

The author answers the Codex record of `docs/reviews/pr-261.md`. That record read `Changes required` at head `f3ff561`, with two findings.

## P2-1: A Commander companion is absent from the stale-card banner

- The result: full merit.
- The evidence: an imported list keeps its companion as `CompanionOracleId` alone (`go/internal/agentsvc/import.go`). In Commander the companion is the 101st card, outside each list (`go/internal/rules/checks.go`, `checkCompanion`). So `staleNames` found no name for it. The same gap held for the commander of an imported list, because the import fills `CommanderOracleIds` and no `commanders` entry. A deck built before D-604 has no `commanders` entry either.
- The correction: `staleNames` of `web/apps/web/src/features/deck/stale-banner.tsx` takes the card data of the deck, and it names a stale id that sits in no list. The commander stays first, and the companion comes last. The banner reads the card data with `useDeckCards`, the same query as the deck view, so it adds no call. `deckOracleIds` of `web/apps/web/src/features/deck/use-cards.ts` adds the companion, so the card data holds it.
- The regression check: "names a companion outside the lists from the card data" of `web/apps/web/src/features/deck/stale-banner.test.tsx` fails with the companion line removed, and passes with it. "names an imported commander from the card data, still first" holds the commander case. A new test of `use-cards.test.ts` holds the companion id. The deck, workspace, and chat tests pass, 219 of 219.

## P2-2: A Commander companion is absent from patch removals

- The result: full merit.
- The evidence: `staleCardNames` of `go/internal/agentsvc/rerun.go` read three lists, and not `CompanionOracleId`. So the patch brief of D-1019 omitted the companion. The rebuild of a banned imported commander also unlocked no commander name.
- The correction: `staleCardNames` takes the card index of the server, and it names a stale commander or companion that sits in no list. `rerunCards` gives the index, or nil before the index loads. With no index, the function names the cards of the lists alone, as before. Both callers of `prepareRerun` pass the index.
- A note on the scope: only an import sets `CompanionOracleId`. A revision builds a new deck through the build path, and that path makes no companion. So a patched deck holds no companion with or without the fix. The fix puts the companion in `Remove`, and the build prompt then names it as a card to keep out.
- The regression check: `TestRerunPatchNamesACompanionOutsideTheLists` of `go/internal/agentsvc/rerun_test.go` fails with the index lookup turned off: `Remove:[]`. It passes with the lookup. `TestStaleCardNamesReadsTheIndexForAnIdAlone` holds the imported commander and the case with no index.
