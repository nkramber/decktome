import type { Deck } from "@mtg/api-client/mtg/v1/deck_pb";
import { QueryClientProvider } from "@tanstack/react-query";
import { render as baseRender, screen, waitFor } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { axe } from "jest-axe";
import type { ReactElement } from "react";
import { beforeEach, describe, expect, it, vi } from "vitest";
import { makeQueryClient } from "../../lib/query-client";
import { NewCardsPanel, newCardIds, reviseMessage } from "./new-cards-panel";

const getCards = vi.fn();
const updateDeck = vi.fn();
vi.mock("../../lib/api", () => ({
  cardClient: { getCards: (...args: unknown[]) => getCards(...args) },
  deckClient: { updateDeck: (...args: unknown[]) => updateDeck(...args) },
}));

beforeEach(() => {
  getCards.mockReset();
  updateDeck.mockReset();
  getCards.mockResolvedValue({
    cards: [
      { oracleId: "o-rex", name: "Tyrant Rex" },
      { oracleId: "o-egg", name: "Hatchling Egg" },
    ],
    missingOracleIds: [],
  });
  updateDeck.mockResolvedValue({ deck: {} });
});

function render(ui: ReactElement) {
  const client = makeQueryClient();
  return baseRender(<QueryClientProvider client={client}>{ui}</QueryClientProvider>);
}

const deck = {
  id: "d1",
  newOracleIds: ["o-rex", "o-egg", "o-held"],
  commanderOracleIds: ["o-cmd"],
  cards: [{ oracleId: "o-held", name: "Held Card" }],
  sideboard: [],
  upgrades: [],
} as unknown as Deck;

describe("newCardIds", () => {
  it("leaves out a card that the deck holds (D-1091)", () => {
    expect(newCardIds(deck)).toEqual(["o-rex", "o-egg"]);
  });

  it("reads a deck with no field as none", () => {
    expect(newCardIds({ ...deck, newOracleIds: undefined } as unknown as Deck)).toEqual([]);
  });
});

describe("NewCardsPanel", () => {
  it("names the new cards and sends the revise (D-1091)", async () => {
    const onRevise = vi.fn();
    const { container } = render(<NewCardsPanel deck={deck} onRevise={onRevise} />);
    expect(screen.getByRole("heading", { name: "New cards for this deck" })).toBeInTheDocument();
    expect(await screen.findByText("Tyrant Rex, Hatchling Egg")).toBeInTheDocument();
    await userEvent.click(screen.getByRole("button", { name: "Revise with these cards" }));
    expect(onRevise).toHaveBeenCalledWith(reviseMessage(["Tyrant Rex", "Hatchling Egg"]));
    expect(await axe(container)).toHaveNoViolations();
  });

  it("dismisses the panel through UpdateDeck (D-1095)", async () => {
    render(<NewCardsPanel deck={deck} />);
    await screen.findByText("Tyrant Rex, Hatchling Egg");
    expect(screen.queryByRole("button", { name: "Revise with these cards" })).not.toBeInTheDocument();
    await userEvent.click(screen.getByRole("button", { name: "Dismiss" }));
    expect(updateDeck).toHaveBeenCalledWith({ deckId: "d1", dismissNewCards: true });
    await waitFor(() => expect(screen.queryByRole("heading", { name: "New cards for this deck" })).not.toBeInTheDocument());
  });

  it("keeps the panel and says so when the dismiss fails", async () => {
    updateDeck.mockRejectedValue(new Error("down"));
    render(<NewCardsPanel deck={deck} />);
    await userEvent.click(await screen.findByRole("button", { name: "Dismiss" }));
    expect(await screen.findByRole("alert")).toHaveTextContent("Could not dismiss the new cards");
    expect(screen.getByRole("heading", { name: "New cards for this deck" })).toBeInTheDocument();
  });

  it("shows nothing on a deck with no new cards", () => {
    const { container } = render(<NewCardsPanel deck={{ ...deck, newOracleIds: ["o-held"] } as Deck} />);
    expect(container).toBeEmptyDOMElement();
  });
});
