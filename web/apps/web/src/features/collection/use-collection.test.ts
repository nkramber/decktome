import type { CollectionEntry } from "@mtg/api-client/mtg/v1/collection_pb";
import { describe, expect, it } from "vitest";

import { artIds, statsOf } from "./use-collection";

function entry(over: Partial<CollectionEntry>): CollectionEntry {
  return { oracleId: "o", name: "n", quantity: 1, rarity: "common", setName: "Set", setCode: "st", ...over } as CollectionEntry;
}

const rows = [
  entry({ oracleId: "o-bolt", name: "Lightning Bolt", quantity: 4, rarity: "common", setName: "Alpha" }),
  entry({ oracleId: "o-bolt", name: "Lightning Bolt", quantity: 2, rarity: "common", setName: "Beta" }),
  entry({ oracleId: "o-jace", name: "Jace", quantity: 1, rarity: "mythic", setName: "Worldwake" }),
  entry({ oracleId: "o-path", name: "Path to Exile", quantity: 3, rarity: "uncommon", setName: "Alpha" }),
];

describe("statsOf", () => {
  it("sums the quantity, counts the Oracle ids, and groups the rarity", () => {
    const s = statsOf(rows);
    expect(s.total).toBe(10);
    expect(s.unique).toBe(3);
    expect(s.rarity).toEqual([
      { key: "mythic", label: "Mythic", count: 1 },
      { key: "uncommon", label: "Uncommon", count: 3 },
      { key: "common", label: "Common", count: 6 },
    ]);
  });

  it("orders the sets by the cards they hold", () => {
    expect(statsOf(rows).sets[0]).toEqual({ name: "Alpha", count: 7 });
  });

  it("reads an empty collection as zero, not as an error", () => {
    expect(statsOf([])).toEqual({ total: 0, unique: 0, rarity: [], sets: [] });
  });
});

describe("artIds", () => {
  it("takes the rarest cards first, and each Oracle id once", () => {
    expect(artIds(rows, 3)).toEqual(["o-jace", "o-path", "o-bolt"]);
  });

  it("stops at the count the caller asks for", () => {
    expect(artIds(rows, 1)).toEqual(["o-jace"]);
  });
});
