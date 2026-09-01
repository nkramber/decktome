import { Condition, type CollectionEntry, Finish } from "@mtg/api-client/mtg/v1/collection_pb";
import { describe, expect, it } from "vitest";

import { filterEntries, setOptions, sortEntries } from "./binder-grid";

function row(over: Partial<CollectionEntry>): CollectionEntry {
  return {
    scryfallId: "p", oracleId: "o", name: "Card", setCode: "lea", setName: "Alpha",
    collectorNumber: "1", quantity: 1, finish: Finish.NORMAL, condition: Condition.NEAR_MINT,
    rarity: "common", language: "en", ...over,
  } as CollectionEntry;
}

const rows = [
  row({ name: "Lightning Bolt", setCode: "lea", setName: "Alpha", quantity: 4 }),
  row({ name: "Jace, the Mind Sculptor", setCode: "wwk", setName: "Worldwake", quantity: 1, rarity: "mythic" }),
  row({ name: "Path to Exile", setCode: "lea", setName: "Alpha", quantity: 3 }),
];

describe("the binder filter", () => {
  it("matches a card name, whatever the case", () => {
    expect(filterEntries(rows, "bolt", "").map((e) => e.name)).toEqual(["Lightning Bolt"]);
    expect(filterEntries(rows, "  BOLT  ", "").map((e) => e.name)).toEqual(["Lightning Bolt"]);
  });

  it("keeps every row for an empty search", () => {
    expect(filterEntries(rows, "", "")).toHaveLength(3);
  });

  it("filters by set, and the two run together", () => {
    expect(filterEntries(rows, "", "lea")).toHaveLength(2);
    expect(filterEntries(rows, "path", "lea").map((e) => e.name)).toEqual(["Path to Exile"]);
    expect(filterEntries(rows, "path", "wwk")).toHaveLength(0);
  });
});

describe("the binder sort", () => {
  it("orders by name, and the name breaks every tie", () => {
    expect(sortEntries(rows, "name").map((e) => e.name)).toEqual([
      "Jace, the Mind Sculptor", "Lightning Bolt", "Path to Exile",
    ]);
  });

  it("orders by count, largest first", () => {
    expect(sortEntries(rows, "count").map((e) => e.quantity)).toEqual([4, 3, 1]);
  });

  it("orders by set, and the name breaks the tie inside one set", () => {
    expect(sortEntries(rows, "set").map((e) => e.name)).toEqual([
      "Lightning Bolt", "Path to Exile", "Jace, the Mind Sculptor",
    ]);
  });

  it("changes no input", () => {
    const before = rows.map((e) => e.name);
    sortEntries(rows, "count");
    expect(rows.map((e) => e.name)).toEqual(before);
  });
});

describe("the binder set list", () => {
  it("names every set it holds, largest first, and offers a way out", () => {
    const got = setOptions(rows);
    expect(got[0]).toEqual({ value: "", label: "Every set" });
    // Alpha holds 7 cards and Worldwake holds 1.
    expect(got[1].label).toBe("Alpha (7)");
    expect(got[2].label).toBe("Worldwake (1)");
  });

  it("skips a row with no set code", () => {
    expect(setOptions([row({ setCode: "" })])).toHaveLength(1);
  });
});
