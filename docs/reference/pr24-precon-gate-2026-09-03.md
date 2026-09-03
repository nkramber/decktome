# PR-24 precon exclusion gate (2026-09-03)

Run date: 2026-09-03. Branch: `pr-24`, from `main` at 4b0867f. Card snapshot: 20260903T090227. Precon table: `5.3.0+20260903`, 701 products. The gate is free: every line is a Go test, and no model call runs (D-495).

## The gate lines of the roadmap

| Line | Test | Result |
|---|---|---|
| A build for a collection that holds a precon whole uses none of its cards | `agentsvc.TestBuildExcludesThePreconsCards` | PASS: the three cards with no copy to spare leave the pool and reach the generator as excluded, and the surplus Sol Ring stays |
| The shortlist and the commander pool drop an excluded card | `candidates.TestBuildDropsTheExcludedCards`, `candidates.TestCommanderPoolDropsAnExcludedCommander` | PASS |
| The rules check blocks a card that slips through | `rules.TestGoldenDecks/bad/excluded_precon_card` | PASS: one block per card, a basic land never blocks |
| A collection that holds 99 of 100 owns nothing | `precons.TestOwnedWholeReadsEveryPrinting`, `agentsvc.TestOwnedPreconsReadThePrintings` | PASS: every printing with its count, and a Sol Ring of another set does not count |
| "Not from my precons" reads the owned products | `questions.TestApplyPreconsReadsTheOwnedPrecons` | PASS: two owned products, none owned, and no collection each give their own message |
| A named product not held whole is excluded anyway (D-497) | `questions.TestApplyPreconsMarksAPartialProduct`, `agentsvc.TestPreconNotesNameTheProducts` | PASS |
| An unknown name asks, once per name (D-376 shape) | `questions.TestApplyPreconsAsksAboutAnUnknownName`, `questions.TestThePreconRowAsksOncePerName`, `questions.TestConversations/a_precon_name_that_names_two_products` | PASS |
| The nine embedded lists match the table (D-498) | `precons.TestEmbeddedListsMatchTheTable` | PASS after one correction, see below |
| The words resolve to a product | `precons.TestResolveReadsTheProductName`, `precons.TestResolveAsksAboutTwoProductsWithOneName`, `precons.TestResolveOffersNearNamesForAnUnknownPhrase` | PASS |

## The cross-check of the nine lists

The first run read one difference. The file `ff-cloud.txt` held Dancer's Chakrams (FIC 17), and the MTGJSON row of Limit Break (FINAL FANTASY VII) holds Furious Rise (FIC 294). The Wizards decklist page and mtg.wtf both list Furious Rise in Limit Break, and Dancer's Chakrams in Scions & Spellcraft, read on 2026-09-03. The file is corrected (D-501), and the second run reads nine of nine.

Two file names differed from the product names, with the same cards: `lorwyn-blight-curse` is "Blight Curse" (ECC), and `turtle-power` is "Turtle Power!" (TMC). MTGJSON lists From Cute to Brute and Goblin Storm as Commander Decks of set SLD, so the table holds them.

## The tree

Go build, vet, the `-race` suite, golangci-lint, `make ste-check`, the web typecheck, and the 237 web tests pass. `make proto` regenerated the Go and the TypeScript code for `Slots.exclude_precon_keys`.

## What the gate does not prove

- That the classify model fills `precon_names` and `facts.exclude_precons`. The unit tests script the classifier. The next paid question gate reads it (D-495, D-496).
- That a real build with the exclusion reads well. The paid deck gate has no prompt for it yet, and a prompt needs a collection fixture that holds a whole precon.

Verdict: PASS. Time: the Go suite, about two minutes. No provider call.
