import type { Card } from "@mtg/api-client/mtg/v1/card_pb";
import type { CollectionEntry, CollectionSummary } from "@mtg/api-client/mtg/v1/collection_pb";
import { useInfiniteQuery, useQuery } from "@tanstack/react-query";

import { cardClient, collectionClient } from "../../lib/api";

// The collection screen shows the binder, not only the form that filled
// it (D-327). The head reads the summary the import stored, and never an
// entry, so a binder of 2,657 rows costs one small answer (D-392).

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

// statsFrom reads the stored summary (D-392). The head shows cards and
// not rows, copies included, and the rarity bar needs a total to divide
// by, so the total is the sum of the rarity counts.
export function statsFrom(summary: CollectionSummary | undefined, cardCount: number): CollectionStats {
  const rarity = rarityOrder
    .map((r) => ({ ...r, count: summary?.byRarity[r.key] ?? 0 }))
    .filter((r) => r.count > 0);
  return {
    total: cardCount || rarity.reduce((n, r) => n + r.count, 0),
    unique: summary?.uniqueCards ?? 0,
    rarity,
    sets: (summary?.topSets ?? []).map((s) => ({ name: s.setName || s.setCode, count: s.count })),
  };
}

// statsOf counts a page of entries. The binder grid reads it for the
// rows it holds, and the head reads statsFrom instead.
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

// useCollectionHead reads the collection without its entries (D-392).
// The head draws from the summary alone, so it moves no megabyte.
export function useCollectionHead(collectionId: string) {
  return useQuery({
    queryKey: ["collection", "head", collectionId],
    queryFn: () => collectionClient.getCollection({ collectionId, entriesOmitted: true }),
    enabled: collectionId !== "",
    staleTime: Infinity,
  });
}

// binderPageSize is how many rows one page of the binder holds. The grid
// asks for the next page as the reader scrolls.
export const binderPageSize = 200;

// useBinderPages reads the binder a page at a time (D-392).
export function useBinderPages(collectionId: string) {
  return useInfiniteQuery({
    queryKey: ["collection", "binder", collectionId],
    queryFn: ({ pageParam }) =>
      collectionClient.getCollection({ collectionId, pageSize: binderPageSize, pageToken: pageParam }),
    initialPageParam: "",
    // An empty token is the last page.
    getNextPageParam: (last) => last.nextPageToken || undefined,
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
