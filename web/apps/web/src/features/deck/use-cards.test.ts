import type { Deck } from "@mtg/api-client/mtg/v1/deck_pb";
import { beforeEach, describe, expect, it, vi } from "vitest";

import { deckOracleIds, fetchCards, getCardsMax } from "./use-cards";

const getCards = vi.fn();
vi.mock("../../lib/api", () => ({
  cardClient: { getCards: (...args: unknown[]) => getCards(...args) },
}));

beforeEach(() => {
  getCards.mockReset();
});

describe("use-cards", () => {
  it("deckOracleIds lists the commander first, then every zone once, and skips empty ids", () => {
    const deck = {
      commanderOracleIds: ["o-cmd"],
      cards: [{ oracleId: "o-a" }, { oracleId: "o-cmd" }, { oracleId: "" }],
      sideboard: [{ oracleId: "o-b" }],
      upgrades: [{ oracleId: "o-a" }, { oracleId: "o-c" }],
    } as unknown as Deck;
    expect(deckOracleIds(deck)).toEqual(["o-cmd", "o-a", "o-b", "o-c"]);
  });

  it("fetchCards sends the ids in chunks of 120 and merges the answers", async () => {
    const ids = Array.from({ length: 250 }, (_, i) => `o-${i}`);
    getCards.mockImplementation((req: { oracleIds: string[] }) => {
      const { oracleIds } = req;
      return Promise.resolve({ cards: oracleIds.filter((id) => id !== "o-7").map((id) => ({ oracleId: id, name: id })), missingOracleIds: oracleIds.includes("o-7") ? ["o-7"] : [] });
    });
    const { byId, missing } = await fetchCards(ids);
    expect(getCards).toHaveBeenCalledTimes(3);
    expect(getCards.mock.calls.map((c) => (c[0] as { oracleIds: string[] }).oracleIds.length)).toEqual([getCardsMax, getCardsMax, 10]);
    expect(byId.size).toBe(249);
    expect(byId.get("o-249")?.name).toBe("o-249");
    expect(missing).toEqual(["o-7"]);
  });

  it("fetchCards with no ids calls nothing", async () => {
    const { byId, missing } = await fetchCards([]);
    expect(getCards).not.toHaveBeenCalled();
    expect(byId.size).toBe(0);
    expect(missing).toEqual([]);
  });
});
