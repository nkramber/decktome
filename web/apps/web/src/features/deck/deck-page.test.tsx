import { Code, ConnectError } from "@connectrpc/connect";
import { FormatId } from "@mtg/api-client/mtg/v1/format_pb";
import { screen, waitFor } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { axe } from "jest-axe";
import { beforeEach, describe, expect, it, vi } from "vitest";

import { fakeUser, state } from "../../test-auth-state";
import { renderAt } from "../../test-utils";

vi.mock("firebase/app");
vi.mock("firebase/auth");

const getDeck = vi.fn();
const updateDeck = vi.fn();
const deleteDeck = vi.fn();
vi.mock("../../lib/api", () => ({
  healthClient: { check: () => Promise.resolve({ status: "ok", version: "test", cardSnapshot: "none" }) },
  collectionClient: { listCollections: () => Promise.resolve({ collections: [] }) },
  agentClient: { getSession: vi.fn() },
  cardClient: { getCards: () => Promise.resolve({ cards: [], missingOracleIds: [] }) },
  deckClient: {
    listDecks: () => Promise.resolve({ decks: [], nextPageToken: "" }),
    getDeck: (...a: unknown[]) => getDeck(...a),
    updateDeck: (...a: unknown[]) => updateDeck(...a),
    deleteDeck: (...a: unknown[]) => deleteDeck(...a),
    exportDeck: vi.fn(),
  },
}));

const deck = {
  id: "d1",
  name: "Elf Ball",
  sessionId: "s1",
  format: { id: FormatId.COMMANDER },
  cards: [],
  sideboard: [],
  upgrades: [],
  commanderOracleIds: [],
  favorite: false,
  legalityAsOf: "2026-08-24",
  validation: { passed: true, findings: [] },
};

beforeEach(() => {
  state.user = fakeUser;
  localStorage.clear();
  getDeck.mockReset();
  updateDeck.mockReset();
  deleteDeck.mockReset();
  getDeck.mockResolvedValue({ deck });
  updateDeck.mockResolvedValue({ deck });
  deleteDeck.mockResolvedValue({});
});

describe("DeckPage", () => {
  it("reads the deck of the path and shows it", async () => {
    await renderAt("/decks/d1");
    expect(await screen.findByRole("heading", { name: "Elf Ball" })).toBeInTheDocument();
    expect(getDeck).toHaveBeenCalledWith({ deckId: "d1" });
  });

  it("renames the deck and reports it", async () => {
    await renderAt("/decks/d1");
    const user = userEvent.setup();
    await user.click(await screen.findByRole("button", { name: "Rename" }));
    const field = await screen.findByLabelText("Deck name");
    await user.clear(field);
    await user.type(field, "  Marwyn ramp  ");
    await user.click(screen.getByRole("button", { name: "Save the name" }));
    await waitFor(() => expect(updateDeck).toHaveBeenCalledWith({ deckId: "d1", name: "Marwyn ramp" }));
  });

  it("asks before it deletes, and goes back to the library after", async () => {
    const { router } = await renderAt("/decks/d1");
    const user = userEvent.setup();
    await user.click(await screen.findByRole("button", { name: "Delete" }));
    // Nothing goes out until the question is answered.
    expect(deleteDeck).not.toHaveBeenCalled();
    await user.click(await screen.findByRole("button", { name: "Delete the deck" }));
    await waitFor(() => expect(deleteDeck).toHaveBeenCalledWith({ deckId: "d1" }));
    await waitFor(() => expect(router.state.location.pathname).toBe("/decks"));
  });

  it("keeps the deck when the question is refused", async () => {
    await renderAt("/decks/d1");
    const user = userEvent.setup();
    await user.click(await screen.findByRole("button", { name: "Delete" }));
    await user.click(await screen.findByRole("button", { name: "Keep the deck" }));
    expect(deleteDeck).not.toHaveBeenCalled();
  });

  it("writes the favorite mark", async () => {
    await renderAt("/decks/d1");
    await userEvent.setup().click(await screen.findByRole("button", { name: "Add to your favorites" }));
    await waitFor(() => expect(updateDeck).toHaveBeenCalledWith({ deckId: "d1", favorite: true }));
  });

  it("says a deleted deck is gone, and offers the way back", async () => {
    getDeck.mockRejectedValue(new ConnectError("deck not found", Code.NotFound));
    await renderAt("/decks/d9");
    expect(await screen.findByText("That deck is gone")).toBeInTheDocument();
    expect(screen.getByRole("link", { name: "Back to your decks" })).toHaveAttribute("href", "/decks");
    // A gone deck offers no retry: the answer will not change.
    expect(screen.queryByRole("button", { name: "Try again" })).not.toBeInTheDocument();
  });

  it("offers a retry when the read failed for another reason", async () => {
    getDeck.mockRejectedValue(new ConnectError("firestore down", Code.Internal));
    await renderAt("/decks/d1");
    expect(await screen.findByText("Could not load the deck")).toBeInTheDocument();
    expect(screen.getByRole("button", { name: "Try again" })).toBeInTheDocument();
  });

  it("has no axe violations", async () => {
    const { container } = await renderAt("/decks/d1");
    await screen.findByRole("heading", { name: "Elf Ball" });
    expect(await axe(container)).toHaveNoViolations();
  });
});
