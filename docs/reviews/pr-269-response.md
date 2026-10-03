# Response to the review of pull request 269

The author answers the Codex record of `docs/reviews/pr-269.md`. That record read `Changes required` at head `b57de37`, with one finding.

## P2-1: A card with no eligible priced printing keeps an unrelated price

- The result: full merit.
- The evidence: the index set a card price only from the `cheapest` map. A card whose priced printings were all digital or `NotForPlay` kept the price of its Oracle row. That price can come from a digital or a memorabilia printing, so it is not the price of a copy to play (D-1057).
- The correction: `go/internal/cards/index.go` clears the price and its date for a card whose printings hold no playable paper price. A card with no printing in the file keeps its own price, because nothing shows that price to be wrong. Scryfall `default_cards` holds every printing, so this case occurs in test data and trimmed snapshots alone. D-1057 now states the rule.
- The regression check: `TestCardPriceIsCheapestPlayablePrinting` adds a card with a $300 Oracle price and only memorabilia, digital, and unpriced printings. It expects no price. It also expects no price for a card with one unpriced printing, and its own price for a card with no printing. The test fails on `b57de37` with "price = 300" and "price = 4", and passes with the correction. `go test ./...` passes.
