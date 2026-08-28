import type { Card } from "@mtg/api-client/mtg/v1/card_pb";
import type { Deck } from "@mtg/api-client/mtg/v1/deck_pb";
import { useQuery } from "@tanstack/react-query";

import { cardClient } from "../../lib/api";

// One GetCards call carries at most 120 ids. A deck with a sideboard and
// upgrades can pass that, so the ids go out in chunks.
export const getCardsMax = 120;

export function deckOracleIds(deck: Deck): string[] {
  const ids = new Set<string>();
  for (const dc of [...deck.cards, ...deck.sideboard, ...deck.upgrades]) {
    if (dc.oracleId) ids.add(dc.oracleId);
  }
  return [...ids];
}

export async function fetchCards(ids: string[]): Promise<{ byId: Map<string, Card>; missing: string[] }> {
  const chunks: string[][] = [];
  for (let i = 0; i < ids.length; i += getCardsMax) {
    chunks.push(ids.slice(i, i + getCardsMax));
  }
  const responses = await Promise.all(chunks.map((oracleIds) => cardClient.getCards({ oracleIds })));
  const byId = new Map<string, Card>();
  const missing: string[] = [];
  for (const res of responses) {
    for (const c of res.cards) byId.set(c.oracleId, c);
    missing.push(...res.missingOracleIds);
  }
  return { byId, missing };
}

// useDeckCards loads the card data of one deck: art, faces, type line,
// mana value, and produced mana. Keyed by the deck id and the id list, so
// a deck that changes refetches.
export function useDeckCards(deck: Deck) {
  const ids = deckOracleIds(deck);
  return useQuery({
    queryKey: ["cards", deck.id, ids],
    queryFn: () => fetchCards(ids),
    staleTime: Infinity,
  });
}
