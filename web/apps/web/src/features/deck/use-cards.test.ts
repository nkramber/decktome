import { Code, ConnectError } from "@connectrpc/connect";
import type { Deck } from "@mtg/api-client/mtg/v1/deck_pb";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { renderHook, waitFor } from "@testing-library/react";
import { createElement, type ReactNode } from "react";
import { beforeEach, describe, expect, it, vi } from "vitest";

import { deckOracleIds, fetchCards, getCardsMax, useDeckCards } from "./use-cards";

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

  // The API reads Unavailable until its card snapshot loads, for about 90
  // seconds after a cold start (F-164). useDeckCards keeps asking through
  // that window, so the art arrives with no reload.
  it("useDeckCards retries while the card database is Unavailable", async () => {
    getCards
      .mockRejectedValueOnce(new ConnectError("card database not loaded yet", Code.Unavailable))
      .mockResolvedValue({ cards: [{ oracleId: "o-a", name: "Sol Ring" }], missingOracleIds: [] });
    const deck = { id: "d1", commanderOracleIds: [], cards: [{ oracleId: "o-a" }], sideboard: [], upgrades: [] } as unknown as Deck;
    const client = new QueryClient();
    const wrapper = ({ children }: { children: ReactNode }) => createElement(QueryClientProvider, { client }, children);
    const { result } = renderHook(() => useDeckCards(deck), { wrapper });
    await waitFor(() => expect(result.current.byId.get("o-a")?.name).toBe("Sol Ring"), { timeout: 5000 });
    expect(getCards).toHaveBeenCalledTimes(2);
  });

  // Any other code fails at once, so a signed-out reader reads the error
  // instead of a page that waits.
  it("useDeckCards takes an Unauthenticated answer at once", async () => {
    getCards.mockRejectedValue(new ConnectError("no token", Code.Unauthenticated));
    const deck = { id: "d2", commanderOracleIds: [], cards: [{ oracleId: "o-a" }], sideboard: [], upgrades: [] } as unknown as Deck;
    const client = new QueryClient();
    const wrapper = ({ children }: { children: ReactNode }) => createElement(QueryClientProvider, { client }, children);
    const { result } = renderHook(() => useDeckCards(deck), { wrapper });
    await waitFor(() => expect(result.current.isError).toBe(true));
    expect(getCards).toHaveBeenCalledTimes(1);
  });
});
