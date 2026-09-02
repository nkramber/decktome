import { type Card, Color } from "@mtg/api-client/mtg/v1/card_pb";
import { BinderSort, type CollectionEntry, type CollectionSummary } from "@mtg/api-client/mtg/v1/collection_pb";
import { keepPreviousData, useInfiniteQuery, useQueries, useQuery } from "@tanstack/react-query";
import { useEffect, useMemo, useState } from "react";

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
    sets: (summary?.sets ?? []).map((s) => ({ name: s.setName || s.setCode, count: s.count })),
  };
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

// BinderChoice is what the binder controls hold. Every string is the
// value of its control, and "" is no choice.
export type BinderChoice = { query: string; set: string; color: string; type: string; count: string };

export const noChoice: BinderChoice = { query: "", set: "", color: "", type: "", count: "" };

export type BinderSortKey = "name" | "count" | "set" | "price";

const colorOf: Record<string, Color> = { W: Color.W, U: Color.U, B: Color.B, R: Color.R, G: Color.G };

const sortOf: Record<BinderSortKey, BinderSort> = {
  name: BinderSort.NAME,
  count: BinderSort.COUNT,
  set: BinderSort.SET,
  price: BinderSort.PRICE,
};

export type BinderRequestFilter = {
  query: string;
  setCode: string;
  color: Color;
  colorless: boolean;
  cardType: string;
  quantity: number;
  quantityOrMore: boolean;
};

// binderFilter turns the controls into the filter the server runs over
// every row (D-398). "C" is colorless, and "4" is four or more, which
// is a playset.
export function binderFilter(c: BinderChoice): BinderRequestFilter {
  const n = Number(c.count) || 0;
  return {
    query: c.query.trim(),
    setCode: c.set,
    color: colorOf[c.color] ?? Color.UNSPECIFIED,
    colorless: c.color === "C",
    cardType: c.type,
    quantity: n,
    quantityOrMore: n >= 4,
  };
}

export function binderSort(key: BinderSortKey): BinderSort {
  return sortOf[key];
}

// binderPageSize is how many rows one page of the binder holds. The grid
// asks for the next page as the reader scrolls.
export const binderPageSize = 200;

// useBinderPages reads the binder a page at a time (D-392). The filter
// and the sort go to the server, so a match on a later page still shows
// and the set list needs no row (D-398).
export function useBinderPages(collectionId: string, choice: BinderChoice, sort: BinderSortKey) {
  const filter = binderFilter(choice);
  return useInfiniteQuery({
    queryKey: ["collection", "binder", collectionId, filter, sort],
    queryFn: ({ pageParam }) =>
      collectionClient.getCollection({ collectionId, pageSize: binderPageSize, pageToken: pageParam, filter, sort: binderSort(sort) }),
    initialPageParam: "",
    // An empty token is the last page.
    getNextPageParam: (last) => last.nextPageToken || undefined,
    enabled: collectionId !== "",
    staleTime: Infinity,
    // The last answer stays on screen while a new choice loads, so the
    // grid never blanks between two filters.
    placeholderData: keepPreviousData,
  });
}

// useDebounced hands back a value once it has stood still for ms. The
// search sends one request per pause, not one per keystroke.
export function useDebounced<T>(value: T, ms: number): T {
  const [settled, setSettled] = useState(value);
  useEffect(() => {
    const timer = setTimeout(() => setSettled(value), ms);
    return () => clearTimeout(timer);
  }, [value, ms]);
  return settled;
}

// artBucket is how many rows share one art request (D-401). A bucket
// holds at most 100 Oracle ids, under the 120 one GetCards call takes.
export const artBucket = 100;

// useBinderArt loads the art of the rows in view, one bucket at a time
// (D-401). A bucket is keyed by its ids, so a scroll inside it asks for
// nothing, and a scroll back finds it in the cache. Before this, every
// row that entered the window made a new request.
export function useBinderArt(rows: CollectionEntry[], first: number, last: number): Map<string, Card> {
  const buckets = useMemo(() => {
    const out: string[][] = [];
    if (rows.length === 0 || last < first) return out;
    const from = Math.floor(Math.max(first, 0) / artBucket);
    const to = Math.floor(Math.min(last, rows.length - 1) / artBucket);
    for (let b = from; b <= to; b++) {
      const ids = [...new Set(rows.slice(b * artBucket, (b + 1) * artBucket).map((e) => e.oracleId))].filter(Boolean);
      if (ids.length > 0) out.push(ids);
    }
    return out;
  }, [rows, first, last]);
  const results = useQueries({
    queries: buckets.map((ids) => ({
      queryKey: ["cards", "binder-art", ids],
      queryFn: () => cardClient.getCards({ oracleIds: ids }),
      staleTime: Infinity,
    })),
  });
  const byId = new Map<string, Card>();
  for (const r of results) {
    for (const c of r.data?.cards ?? []) byId.set(c.oracleId, c);
  }
  return byId;
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
