# Response to the review of pull request 269

The author answers the Codex record of `docs/reviews/pr-269.md`. That record read `Changes required` at head `b57de37`, with one finding.

## P2-1: A card with no eligible priced printing keeps an unrelated price

- The result: full merit.
- The evidence: the index set a card price only from the `cheapest` map. A card whose priced printings were all digital or `NotForPlay` kept the price of its Oracle row. That price can come from a digital or a memorabilia printing, so it is not the price of a copy to play (D-1057).
- The correction: `go/internal/cards/index.go` clears the price and its date for a card whose printings hold no playable paper price. A card with no printing in the file keeps its own price, because nothing shows that price to be wrong. Scryfall `default_cards` holds every printing, so this case occurs in test data and trimmed snapshots alone. D-1057 now states the rule.
- The regression check: `TestCardPriceIsCheapestPlayablePrinting` adds a card with a $300 Oracle price and only memorabilia, digital, and unpriced printings. It expects no price. It also expects no price for a card with one unpriced printing, and its own price for a card with no printing. The test fails on `b57de37` with "price = 300" and "price = 4", and passes with the correction. `go test ./...` passes.

## P2-2: Whole-deck budgets ignore unpriced owned cards

The record of round 2 read `Changes required` at head `fe25a21`, with this one open finding.

- The result: full merit.
- The evidence: a whole-deck cap counts every copy (D-238). `assemble` counted the unpriced cards with `UnpricedToBuy`, which skips a line where the reader owns each copy. So an owned unpriced card added nothing to the whole-deck cost and set no warning. D-1060 counts such a card as over the budget.
- The correction: `UnpricedInDeck` in `go/internal/generate/targets.go` counts every unpriced line of the deck, owned copies included. `assemble` uses it for a whole-deck cap, and `UnpricedToBuy` for a cap on the cards to buy. The warning names the scope: "in the deck" or "to buy".
- The regression check: `TestAnUnpricedCardIsNotFree` adds an owned unpriced card and a priced subtotal of $53.68 under a $100 cap. A cap on the cards to buy gives no warning, and a whole-deck cap gives the warning that names the card. The whole-deck case fails on `fe25a21` with an empty message, and passes with the correction.
