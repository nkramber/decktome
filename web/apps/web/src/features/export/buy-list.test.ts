import type { Card } from "@mtg/api-client/mtg/v1/card_pb";
import type { Deck } from "@mtg/api-client/mtg/v1/deck_pb";
import { describe, expect, it, vi } from "vitest";

import { buyRows, downloadText, scryfallCardUrl } from "./buy-list";

const byId = new Map<string, Card>([
  ["o-sw", { oracleId: "o-sw", name: "Soul Warden", defaultPrinting: { setCode: "MM3", collectorNumber: "24" } } as Card],
  ["o-cmd", { oracleId: "o-cmd", name: "Anikthea, Hand of Erebos", defaultPrinting: { setCode: "cmm", collectorNumber: "8" } } as Card],
]);

const deck = {
  id: "d1",
  commanderOracleIds: ["o-cmd"],
  commanders: [],
  cards: [
    { oracleId: "o-sw", name: "soul warden", count: 4, owned: false, ownedCount: 3, priceUsd: 0.5 },
    { oracleId: "o-sol", name: "Sol Ring", count: 1, owned: true, ownedCount: 0, priceUsd: 2 },
  ],
  sideboard: [{ oracleId: "o-sw", name: "Soul Warden", count: 2, owned: false, ownedCount: 0, priceUsd: 0.5 }],
  upgrades: [{ oracleId: "o-up", name: "Rhystic Study", count: 1, owned: false, ownedCount: 0, priceUsd: 40 }],
} as unknown as Deck;

describe("buyRows", () => {
  it("sums the shortfall per card and links to the printing (D-308)", () => {
    const { needed, upgrades } = buyRows(deck, byId);
    expect(needed).toEqual([{ oracleId: "o-sw", name: "Soul Warden", count: 3, priceUsd: 1.5, scryfallUrl: "https://scryfall.com/card/mm3/24" }]);
    expect(upgrades).toEqual([{ oracleId: "o-up", name: "Rhystic Study", count: 1, priceUsd: 40, scryfallUrl: undefined }]);
  });

  // F-76: a deck carries `commanderOracleIds` and no ownership fact for
  // the commander. The old code wrote one, `owned: false`, so every deck
  // named its commander as a card to buy, at no price. Deck
  // `u8FV7fc98qzNvRfsuJ5q` reads 77 cards, every one owned, and a buy
  // cost of zero, and the buy list held the commander all the same.
  it("names no commander the deck holds no entry for", () => {
    const { needed } = buyRows(deck, byId);
    expect(needed.map((r) => r.oracleId)).not.toContain("o-cmd");
  });

  // D-608: the deck carries the ownership of its commander now, so a
  // commander the reader does not own reaches the buy list with its
  // price, and an owned one does not.
  it("buys a commander the deck marks unowned", () => {
    const d = { ...deck, commanders: [{ oracleId: "o-cmd", name: "Anikthea, Hand of Erebos", count: 1, owned: false, ownedCount: 0, priceUsd: 3.4 }] } as unknown as Deck;
    const { needed } = buyRows(d, byId);
    expect(needed[0]).toEqual({ oracleId: "o-cmd", name: "Anikthea, Hand of Erebos", count: 1, priceUsd: 3.4, scryfallUrl: "https://scryfall.com/card/cmm/8" });
  });

  it("buys no commander the deck marks owned", () => {
    const d = { ...deck, commanders: [{ oracleId: "o-cmd", name: "Anikthea, Hand of Erebos", count: 1, owned: true, ownedCount: 1, priceUsd: 3.4 }] } as unknown as Deck;
    const { needed } = buyRows(d, byId);
    expect(needed.map((r) => r.oracleId)).not.toContain("o-cmd");
  });

  it("gives an owned entry no row, whatever its owned count says", () => {
    const { needed } = buyRows({ ...deck, commanderOracleIds: [], sideboard: [], upgrades: [] } as Deck, byId);
    expect(needed.map((r) => r.name)).toEqual(["Soul Warden"]);
  });

  it("a commander in the card list counts once, with its own entry", () => {
    const d = { ...deck, cards: [...deck.cards, { oracleId: "o-cmd", name: "Anikthea", count: 1, owned: true, ownedCount: 1, priceUsd: 0 }] } as Deck;
    const { needed } = buyRows(d, byId);
    expect(needed.map((r) => r.oracleId)).toEqual(["o-sw"]);
  });
});

describe("scryfallCardUrl", () => {
  it("lowercases the set and escapes the number", () => {
    expect(scryfallCardUrl("MM3", "24")).toBe("https://scryfall.com/card/mm3/24");
    expect(scryfallCardUrl("TSR", "3★")).toBe("https://scryfall.com/card/tsr/3%E2%98%85");
    expect(scryfallCardUrl("", "24")).toBeUndefined();
  });
});

describe("downloadText", () => {
  it("clicks a temporary anchor with the file name", () => {
    const create = vi.fn(() => "blob:x");
    const revoke = vi.fn();
    vi.stubGlobal("URL", { ...URL, createObjectURL: create, revokeObjectURL: revoke });
    const click = vi.spyOn(HTMLAnchorElement.prototype, "click").mockImplementation(() => {});
    downloadText("deck.txt", "Deck\n1 Sol Ring\n");
    expect(create).toHaveBeenCalledTimes(1);
    expect(click).toHaveBeenCalledTimes(1);
    expect(revoke).toHaveBeenCalledWith("blob:x");
    expect(document.querySelector("a[download]")).toBeNull();
    vi.unstubAllGlobals();
    click.mockRestore();
  });
});
