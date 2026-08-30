import { Color } from "@mtg/api-client/mtg/v1/card_pb";
import type { Card } from "@mtg/api-client/mtg/v1/card_pb";
import { describe, expect, it } from "vitest";

import { identityLabel, identityOfCards, identityOfCommanders, sortIdentity } from "./color-identity";

const card = (id: string, identity: Color[]) => ({ oracleId: id, colorIdentity: identity }) as Card;

// The grid holds the commanders of every deck on the page in one map. A
// deck must read its own commanders alone, or a colorless deck takes the
// colors of the deck beside it.
const pageMap = new Map<string, Card>([
  ["o-g", card("o-g", [Color.G])],
  ["o-ub", card("o-ub", [Color.U, Color.B])],
  ["o-c", card("o-c", [])],
]);

describe("identityOfCommanders", () => {
  it("reads the named commanders alone", () => {
    expect(identityOfCommanders(["o-g"], pageMap)).toEqual([Color.G]);
    expect(identityOfCommanders(["o-ub"], pageMap)).toEqual([Color.U, Color.B]);
  });

  it("reads a colorless commander as colorless, not as the whole page", () => {
    expect(identityOfCommanders(["o-c"], pageMap)).toEqual([]);
  });

  it("reads a commander the index has not answered as colorless", () => {
    expect(identityOfCommanders(["o-missing"], pageMap)).toEqual([]);
  });

  it("puts two commanders in the order the game prints them", () => {
    expect(identityOfCommanders(["o-ub", "o-g"], pageMap)).toEqual([Color.U, Color.B, Color.G]);
  });
});

describe("identityOfCards", () => {
  it("reads every card of the map", () => {
    expect(identityOfCards(pageMap)).toEqual([Color.U, Color.B, Color.G]);
  });
});

describe("sortIdentity and identityLabel", () => {
  it("drops duplicates and colorless, and keeps the WUBRG order", () => {
    expect(sortIdentity([Color.G, Color.W, Color.G, Color.C])).toEqual([Color.W, Color.G]);
  });

  it("names the colors in words", () => {
    expect(identityLabel([Color.G])).toBe("Green");
    expect(identityLabel([Color.U, Color.B])).toBe("Blue and Black");
    expect(identityLabel([])).toBe("Colorless");
  });
});
