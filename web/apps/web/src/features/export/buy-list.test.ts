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
  cards: [
    { oracleId: "o-sw", name: "soul warden", count: 4, owned: false, ownedCount: 3, priceUsd: 0.5 },
    { oracleId: "o-sol", name: "Sol Ring", count: 1, owned: true, ownedCount: 0, priceUsd: 2 },
  ],
  sideboard: [{ oracleId: "o-sw", name: "Soul Warden", count: 2, owned: false, ownedCount: 0, priceUsd: 0.5 }],
  upgrades: [{ oracleId: "o-up", name: "Rhystic Study", count: 1, owned: false, ownedCount: 0, priceUsd: 40 }],
} as unknown as Deck;

describe("buyRows", () => {
  it("sums the shortfall per card, commander first, and links to the printing (D-308)", () => {
    const { needed, upgrades } = buyRows(deck, byId);
    expect(needed).toEqual([
      { oracleId: "o-cmd", name: "Anikthea, Hand of Erebos", count: 1, priceUsd: 0, scryfallUrl: "https://scryfall.com/card/cmm/8" },
      { oracleId: "o-sw", name: "Soul Warden", count: 3, priceUsd: 1.5, scryfallUrl: "https://scryfall.com/card/mm3/24" },
    ]);
    expect(upgrades).toEqual([{ oracleId: "o-up", name: "Rhystic Study", count: 1, priceUsd: 40, scryfallUrl: undefined }]);
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
