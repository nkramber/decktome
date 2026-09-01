import { Color } from "@mtg/api-client/mtg/v1/card_pb";
import { Condition, type CollectionEntry, Finish } from "@mtg/api-client/mtg/v1/collection_pb";
import { describe, expect, it } from "vitest";

import { type BinderFilter, filterEntries, noFilter, setOptions, sortEntries, typeOptions } from "./binder-grid";

function row(over: Partial<CollectionEntry>): CollectionEntry {
  return {
    scryfallId: "p", oracleId: "o", name: "Card", setCode: "lea", setName: "Alpha",
    collectorNumber: "1", quantity: 1, finish: Finish.NORMAL, condition: Condition.NEAR_MINT,
    rarity: "common", language: "en", colors: [], cardTypes: [], priceUsd: 0, ...over,
  } as CollectionEntry;
}

const only = (over: Partial<BinderFilter>): BinderFilter => ({ ...noFilter, ...over });

const rows = [
  row({ name: "Lightning Bolt", setCode: "lea", setName: "Alpha", quantity: 4, colors: [Color.R], cardTypes: ["Instant"], priceUsd: 2 }),
  row({ name: "Jace, the Mind Sculptor", setCode: "wwk", setName: "Worldwake", quantity: 1, rarity: "mythic", colors: [Color.U], cardTypes: ["Planeswalker"], priceUsd: 90 }),
  row({ name: "Path to Exile", setCode: "lea", setName: "Alpha", quantity: 3, colors: [Color.W], cardTypes: ["Instant"], priceUsd: 5 }),
];

describe("the binder filter", () => {
  it("matches a card name, whatever the case", () => {
    expect(filterEntries(rows, only({ query: "bolt" })).map((e) => e.name)).toEqual(["Lightning Bolt"]);
    expect(filterEntries(rows, only({ query: "  BOLT  " })).map((e) => e.name)).toEqual(["Lightning Bolt"]);
  });

  it("keeps every row for an empty search", () => {
    expect(filterEntries(rows, noFilter)).toHaveLength(3);
  });

  it("filters by set, and the two run together", () => {
    expect(filterEntries(rows, only({ set: "lea" }))).toHaveLength(2);
    expect(filterEntries(rows, only({ query: "path", set: "lea" })).map((e) => e.name)).toEqual(["Path to Exile"]);
    expect(filterEntries(rows, only({ query: "path", set: "wwk" }))).toHaveLength(0);
  });

  it("filters by color, and a card of two colors answers to both", () => {
    expect(filterEntries(rows, only({ color: "R" })).map((e) => e.name)).toEqual(["Lightning Bolt"]);
    const gold = [row({ name: "Boros Charm", colors: [Color.R, Color.W] })];
    expect(filterEntries(gold, only({ color: "R" }))).toHaveLength(1);
    expect(filterEntries(gold, only({ color: "W" }))).toHaveLength(1);
    expect(filterEntries(gold, only({ color: "U" }))).toHaveLength(0);
  });

  it("reads a card of no color as colorless", () => {
    const rocks = [row({ name: "Sol Ring", colors: [], cardTypes: ["Artifact"] }), ...rows];
    expect(filterEntries(rocks, only({ color: "C" })).map((e) => e.name)).toEqual(["Sol Ring"]);
  });

  it("filters by card type", () => {
    expect(filterEntries(rows, only({ type: "Instant" })).map((e) => e.name)).toEqual(["Lightning Bolt", "Path to Exile"]);
    expect(filterEntries(rows, only({ type: "Land" }))).toHaveLength(0);
  });

  it("filters by count, and four takes every count above it", () => {
    expect(filterEntries(rows, only({ count: "1" })).map((e) => e.name)).toEqual(["Jace, the Mind Sculptor"]);
    expect(filterEntries(rows, only({ count: "3" })).map((e) => e.name)).toEqual(["Path to Exile"]);
    expect(filterEntries([...rows, row({ name: "Rat", quantity: 11 })], only({ count: "4" })).map((e) => e.quantity)).toEqual([4, 11]);
    expect(filterEntries(rows, only({ count: "2" }))).toHaveLength(0);
  });

  it("runs every choice together", () => {
    expect(filterEntries(rows, only({ set: "lea", color: "R", type: "Instant", count: "4" })).map((e) => e.name)).toEqual(["Lightning Bolt"]);
    expect(filterEntries(rows, only({ set: "lea", color: "U" }))).toHaveLength(0);
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

  it("orders by price, dearest first, and an unpriced row sorts last", () => {
    const withUnpriced = [...rows, row({ name: "Aaa Unpriced", priceUsd: 0 })];
    expect(sortEntries(withUnpriced, "price").map((e) => e.name)).toEqual([
      "Jace, the Mind Sculptor", "Path to Exile", "Lightning Bolt", "Aaa Unpriced",
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

describe("the binder type list", () => {
  it("names every type it holds, largest first, and offers a way out", () => {
    const got = typeOptions(rows);
    expect(got[0]).toEqual({ value: "", label: "Every type" });
    // Seven instants over two rows, and one planeswalker.
    expect(got[1].label).toBe("Instant (7)");
    expect(got[2].label).toBe("Planeswalker (1)");
  });

  it("counts a card of two types under each", () => {
    const got = typeOptions([row({ quantity: 2, cardTypes: ["Artifact", "Creature"] })]);
    expect(got.map((o) => o.label)).toEqual(["Every type", "Artifact (2)", "Creature (2)"]);
  });

  it("offers only the way out for rows the index does not know", () => {
    expect(typeOptions([row({ cardTypes: [] })])).toHaveLength(1);
  });
});
