import type { Card } from "@mtg/api-client/mtg/v1/card_pb";
import { type Deck, RerunCase } from "@mtg/api-client/mtg/v1/deck_pb";
import { QueryClientProvider } from "@tanstack/react-query";
import { render as baseRender, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { axe } from "jest-axe";
import type { ReactElement } from "react";
import { beforeEach, describe, expect, it, vi } from "vitest";
import { makeQueryClient } from "../../lib/query-client";
import { rerunLabel, StaleBanner, staleNames } from "./stale-banner";

const getCards = vi.fn();
vi.mock("../../lib/api", () => ({
  cardClient: { getCards: (...args: unknown[]) => getCards(...args) },
}));

beforeEach(() => {
  getCards.mockReset();
  getCards.mockResolvedValue({ cards: [], missingOracleIds: [] });
});

// The banner reads the card data of the deck, so each render needs a
// query client.
function render(ui: ReactElement) {
  const client = makeQueryClient();
  const wrap = (node: ReactElement) => <QueryClientProvider client={client}>{node}</QueryClientProvider>;
  const out = baseRender(wrap(ui));
  return { ...out, rerender: (next: ReactElement) => out.rerender(wrap(next)) };
}

const deck = {
  id: "d1",
  stale: true,
  staleOracleIds: ["o-welcome", "o-karlov"],
  staleReason: "The rerun replaces Ajani's Welcome alone, and keeps the rest of the deck.",
  rerunCase: RerunCase.PATCH,
  commanderOracleIds: ["o-karlov"],
  commanders: [{ oracleId: "o-karlov", name: "Karlov of the Ghost Council" }],
  cards: [
    { oracleId: "o-plains", name: "Plains" },
    { oracleId: "o-welcome", name: "Ajani's Welcome" },
  ],
  sideboard: [],
  upgrades: [],
} as unknown as Deck;

describe("StaleBanner", () => {
  it("names the cards, the reason, and the rerun (I-1)", async () => {
    const onRerun = vi.fn();
    const { container } = render(<StaleBanner deck={deck} onRerun={onRerun} />);
    expect(screen.getByRole("alert")).toHaveTextContent("A rule change made this deck illegal");
    expect(screen.getByText("No longer legal: Karlov of the Ghost Council, Ajani's Welcome.")).toBeInTheDocument();
    expect(screen.getByText(deck.staleReason)).toBeInTheDocument();
    await userEvent.click(screen.getByRole("button", { name: "Replace the banned cards" }));
    expect(onRerun).toHaveBeenCalledOnce();
    expect(await axe(container)).toHaveNoViolations();
  });

  it("shows nothing on a legal deck", () => {
    const { container } = render(<StaleBanner deck={{ ...deck, stale: false } as Deck} onRerun={() => {}} />);
    expect(container).toBeEmptyDOMElement();
  });

  // An imported Commander list keeps its companion as an id alone, outside
  // each list, and the card data names it.
  it("names a companion outside the lists from the card data", async () => {
    getCards.mockResolvedValue({
      cards: [{ oracleId: "o-lurrus", name: "Lurrus of the Dream-Den" }],
      missingOracleIds: [],
    });
    render(<StaleBanner deck={{ ...deck, staleOracleIds: ["o-lurrus"], companionOracleId: "o-lurrus" } as Deck} />);
    expect(await screen.findByText("No longer legal: Lurrus of the Dream-Den.")).toBeInTheDocument();
  });

  it("has no button without a chat, and none while a turn runs", () => {
    const { rerender } = render(<StaleBanner deck={deck} />);
    expect(screen.queryByRole("button")).toBeNull();
    rerender(<StaleBanner deck={deck} busy onRerun={() => {}} />);
    expect(screen.getByRole("button")).toBeDisabled();
  });
});

describe("rerunLabel", () => {
  it("names the action of each case (D-1008, D-1021)", () => {
    expect(rerunLabel(deck)).toBe("Replace the banned cards");
    expect(rerunLabel({ ...deck, rerunCase: RerunCase.REBUILD } as Deck)).toBe("Pick a new commander");
    expect(rerunLabel({ ...deck, rerunCase: RerunCase.REBUILD, staleOracleIds: ["o-welcome"] } as Deck)).toBe("Rebuild the deck");
  });

  it("reads names in the order of the deck, the commander first", () => {
    expect(staleNames(deck)).toEqual(["Karlov of the Ghost Council", "Ajani's Welcome"]);
  });

  it("names an imported commander from the card data, still first", () => {
    const imported = { ...deck, commanders: [] } as unknown as Deck;
    const byId = new Map([["o-karlov", { oracleId: "o-karlov", name: "Karlov of the Ghost Council" } as Card]]);
    expect(staleNames(imported, byId)).toEqual(["Karlov of the Ghost Council", "Ajani's Welcome"]);
    expect(staleNames(imported)).toEqual(["Ajani's Welcome"]);
  });
});
