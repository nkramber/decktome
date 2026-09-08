import type { Card } from "@mtg/api-client/mtg/v1/card_pb";
import type { Deck, DeckCard } from "@mtg/api-client/mtg/v1/deck_pb";

// BuyRow is one card the user still needs (D-308). The price is the
// build-time display price times the count, and zero when unknown.
export type BuyRow = { oracleId: string; name: string; count: number; priceUsd: number; scryfallUrl?: string };

// scryfallCardUrl opens the printing page of a card (D-308). Scryfall
// routes /card/<set>/<number> to that printing.
export function scryfallCardUrl(setCode: string, collectorNumber: string): string | undefined {
  if (!setCode || !collectorNumber) return undefined;
  return `https://scryfall.com/card/${encodeURIComponent(setCode.toLowerCase())}/${encodeURIComponent(collectorNumber)}`;
}

// shortfall is the count the collection does not cover. An entry marked
// owned needs nothing, whatever its owned count says. Same rule as the
// Go export package.
function shortfall(e: DeckCard): number {
  if (e.owned) return 0;
  return Math.max(0, e.count - e.ownedCount);
}

// buyRows mirrors export.BuyList: the shortfall of the commander, the
// main deck, and the sideboard, summed per Oracle id in order of first
// appearance, then the upgrades with their full count (D-308).
export function buyRows(deck: Deck, byId: Map<string, Card>): { needed: BuyRow[]; upgrades: BuyRow[] } {
  const commanders = new Set(deck.commanderOracleIds);
  // A commander the deck holds no entry for carries no ownership fact
  // (F-76). The fallback here wrote `owned: false`, so a deck whose
  // every card is owned still named its commander as one to buy, with
  // no price beside it. No entry is not evidence of no copy: deck
  // `u8FV7fc98qzNvRfsuJ5q` reads 77 cards, every one owned, and a buy
  // cost of zero, and the buy list held the commander all the same.
  const commanderEntries: DeckCard[] = deck.commanderOracleIds
    .map((id) => deck.cards.find((c) => c.oracleId === id))
    .filter((c): c is DeckCard => c !== undefined);
  const order: string[] = [];
  const tally = new Map<string, { entry: DeckCard; need: number }>();
  const add = (e: DeckCard, need: number) => {
    if (need <= 0) return;
    const t = tally.get(e.oracleId);
    if (t) {
      t.need += need;
      return;
    }
    tally.set(e.oracleId, { entry: e, need });
    order.push(e.oracleId);
  };
  for (const e of commanderEntries) add(e, shortfall(e));
  for (const e of deck.cards) if (!commanders.has(e.oracleId)) add(e, shortfall(e));
  for (const e of deck.sideboard) add(e, shortfall(e));
  const row = (e: DeckCard, count: number): BuyRow => {
    const card = byId.get(e.oracleId);
    const dp = card?.defaultPrinting;
    return {
      oracleId: e.oracleId,
      name: card?.name || e.name,
      count,
      priceUsd: e.priceUsd > 0 ? e.priceUsd * count : 0,
      scryfallUrl: scryfallCardUrl(dp?.setCode ?? "", dp?.collectorNumber ?? ""),
    };
  };
  const needed = order.map((id) => {
    const t = tally.get(id)!;
    return row(t.entry, t.need);
  });
  const upgrades = deck.upgrades.filter((e) => e.count > 0).map((e) => row(e, e.count));
  return { needed, upgrades };
}

// downloadText hands the browser a file to save. The object URL is
// released after the click.
export function downloadText(fileName: string, text: string): void {
  const url = URL.createObjectURL(new Blob([text], { type: "text/plain;charset=utf-8" }));
  const a = document.createElement("a");
  a.href = url;
  a.download = fileName;
  document.body.appendChild(a);
  a.click();
  a.remove();
  URL.revokeObjectURL(url);
}
