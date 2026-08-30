import type { Card } from "@mtg/api-client/mtg/v1/card_pb";
import type { CollectionEntry } from "@mtg/api-client/mtg/v1/collection_pb";
import { useQuery } from "@tanstack/react-query";

import { cardClient, collectionClient } from "../../lib/api";

// The collection screen shows the binder, not only the form that filled
// it (D-327). The stats below come from the entries alone, so they need
// no card lookup: rarity and set name ride on every row.

export type CollectionStats = {
  total: number;
  unique: number;
  rarity: { key: string; label: string; count: number }[];
  sets: { name: string; count: number }[];
};

const rarityOrder = [
  { key: "mythic", label: "Mythic" },
  { key: "rare", label: "Rare" },
  { key: "uncommon", label: "Uncommon" },
  { key: "common", label: "Common" },
];

export function statsOf(entries: CollectionEntry[]): CollectionStats {
  const byRarity = new Map<string, number>();
  const bySet = new Map<string, number>();
  const oracle = new Set<string>();
  let total = 0;
  for (const e of entries) {
    total += e.quantity;
    if (e.oracleId) oracle.add(e.oracleId);
    byRarity.set(e.rarity, (byRarity.get(e.rarity) ?? 0) + e.quantity);
    const set = e.setName || e.setCode;
    if (set) bySet.set(set, (bySet.get(set) ?? 0) + e.quantity);
  }
  return {
    total,
    unique: oracle.size,
    rarity: rarityOrder.map((r) => ({ ...r, count: byRarity.get(r.key) ?? 0 })).filter((r) => r.count > 0),
    sets: [...bySet.entries()]
      .map(([name, count]) => ({ name, count }))
      .sort((a, b) => b.count - a.count)
      .slice(0, 4),
  };
}

// artIds picks the cards the strip shows. The rarest come first, so the
// strip shows the collection at its best, and the order is stable.
export function artIds(entries: CollectionEntry[], want: number): string[] {
  const rank: Record<string, number> = { mythic: 0, rare: 1, uncommon: 2, common: 3 };
  const seen = new Set<string>();
  const picked: CollectionEntry[] = [];
  for (const e of entries) {
    if (!e.oracleId || seen.has(e.oracleId)) continue;
    seen.add(e.oracleId);
    picked.push(e);
  }
  picked.sort((a, b) => (rank[a.rarity] ?? 9) - (rank[b.rarity] ?? 9) || a.name.localeCompare(b.name));
  return picked.slice(0, want).map((e) => e.oracleId);
}

export function useCollection(collectionId: string) {
  return useQuery({
    queryKey: ["collection", collectionId],
    queryFn: () => collectionClient.getCollection({ collectionId }),
    enabled: collectionId !== "",
    staleTime: Infinity,
  });
}

export function useCollectionArt(ids: string[]) {
  const query = useQuery({
    queryKey: ["cards", "collection-art", ids],
    queryFn: () => cardClient.getCards({ oracleIds: ids }),
    enabled: ids.length > 0,
    staleTime: Infinity,
  });
  const byId = new Map<string, Card>();
  for (const c of query.data?.cards ?? []) byId.set(c.oracleId, c);
  return ids.map((id) => byId.get(id)).filter((c): c is Card => c !== undefined);
}
