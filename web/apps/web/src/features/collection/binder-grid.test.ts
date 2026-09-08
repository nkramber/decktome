import type { Card } from "@mtg/api-client/mtg/v1/card_pb";
import { Color } from "@mtg/api-client/mtg/v1/card_pb";
import { BinderSort, type SetCount } from "@mtg/api-client/mtg/v1/collection_pb";
import { describe, expect, it } from "vitest";

import { binderArt, setOptions, typeOptions } from "./binder-grid";
import { binderFilter, binderSort, noChoice } from "./use-collection";

// The binder controls become one request the server filters on (D-398).
describe("the binder request", () => {
  it("sends no filter for no choice", () => {
    expect(binderFilter(noChoice)).toEqual({ query: "", setCode: "", color: Color.UNSPECIFIED, colorless: false, cardType: "", quantity: 0, quantityOrMore: false });
  });

  it("trims the search and names the set and the type as chosen", () => {
    const got = binderFilter({ ...noChoice, query: "  bolt ", set: "lea", type: "Instant" });
    expect(got.query).toBe("bolt");
    expect(got.setCode).toBe("lea");
    expect(got.cardType).toBe("Instant");
  });

  it("reads a color letter as the Color, and C as colorless", () => {
    expect(binderFilter({ ...noChoice, color: "R" })).toMatchObject({ color: Color.R, colorless: false });
    expect(binderFilter({ ...noChoice, color: "C" })).toMatchObject({ color: Color.UNSPECIFIED, colorless: true });
  });

  it("reads a count as exact, and four as four or more", () => {
    expect(binderFilter({ ...noChoice, count: "3" })).toMatchObject({ quantity: 3, quantityOrMore: false });
    expect(binderFilter({ ...noChoice, count: "4" })).toMatchObject({ quantity: 4, quantityOrMore: true });
  });

  it("names each sort", () => {
    expect(binderSort("name")).toBe(BinderSort.NAME);
    expect(binderSort("count")).toBe(BinderSort.COUNT);
    expect(binderSort("set")).toBe(BinderSort.SET);
    expect(binderSort("price")).toBe(BinderSort.PRICE);
  });
});

describe("the binder set list", () => {
  const sets = [
    { setCode: "lea", setName: "Alpha", count: 7 },
    { setCode: "wwk", setName: "", count: 1 },
    { setCode: "", setName: "Nameless", count: 2 },
  ] as SetCount[];

  it("offers a way out, then every set of the summary in its order", () => {
    const got = setOptions(sets);
    expect(got[0]).toEqual({ value: "", label: "Every set" });
    expect(got[1]).toEqual({ value: "lea", label: "Alpha (7)" });
    // A set with no name reads by its code.
    expect(got[2]).toEqual({ value: "wwk", label: "wwk (1)" });
  });

  it("skips a set with no code, because the filter needs one", () => {
    expect(setOptions(sets)).toHaveLength(3);
  });
});

describe("the binder type list", () => {
  it("names every type it holds, largest first, and offers a way out", () => {
    const got = typeOptions({ Planeswalker: 1, Instant: 7, Creature: 7 });
    expect(got[0]).toEqual({ value: "", label: "Every type" });
    // Seven and seven tie, and the name breaks the tie.
    expect(got.map((o) => o.label)).toEqual(["Every type", "Creature (7)", "Instant (7)", "Planeswalker (1)"]);
  });

  it("offers only the way out for a summary with no types", () => {
    expect(typeOptions({})).toHaveLength(1);
  });
});

// F-60: a collection holds a set code and a collector number per copy,
// and the binder drew the default printing. A card the reader owns in
// one set showed the art of another.
describe("binderArt", () => {
  const entry = (imageUris?: { normal: string }) => ({ name: "Lightning Bolt", imageUris }) as never;
  const card = { faces: [], defaultPrinting: { imageUris: { normal: "https://img/default.jpg" } } } as unknown as Card;

  it("draws the printing the reader owns", () => {
    expect(binderArt(entry({ normal: "https://img/lea-161.jpg" }), card)).toBe("https://img/lea-161.jpg");
  });

  it("falls back to the face when the entry holds no art", () => {
    const faced = { faces: [{ imageUris: { normal: "https://img/face.jpg" } }], defaultPrinting: { imageUris: { normal: "https://img/default.jpg" } } } as unknown as Card;
    expect(binderArt(entry(), faced)).toBe("https://img/face.jpg");
  });

  it("falls back to the default printing when there is no face", () => {
    expect(binderArt(entry(), card)).toBe("https://img/default.jpg");
  });

  it("reads an empty string when nothing holds art", () => {
    expect(binderArt(entry(), undefined)).toBe("");
  });
});
