import { FormatId } from "@mtg/api-client/mtg/v1/format_pb";
import { screen, waitFor } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { beforeEach, describe, expect, it, vi } from "vitest";

import { fakeUser, state } from "../../test-auth-state";
import { renderAt } from "../../test-utils";
import { nameFromFile } from "./import-dialog";

vi.mock("firebase/app");
vi.mock("firebase/auth");

const importDeck = vi.fn();
const getDeck = vi.fn();
const listCollections = vi.fn();
vi.mock("../../lib/api", () => ({
  healthClient: { check: () => Promise.resolve({ status: "ok", version: "test", cardSnapshot: "none" }) },
  collectionClient: { listCollections: (...a: unknown[]) => listCollections(...a) },
  agentClient: { getSession: vi.fn(), importDeck: (...a: unknown[]) => importDeck(...a), readImportBracket: vi.fn() },
  cardClient: { getCards: () => Promise.resolve({ cards: [], missingOracleIds: [] }) },
  deckClient: {
    listDecks: () => Promise.resolve({ decks: [], nextPageToken: "" }),
    getDeck: (...a: unknown[]) => getDeck(...a),
  },
}));

const stored = { id: "d9", sessionId: "", name: "Rats", format: { id: FormatId.MODERN }, cards: [], sideboard: [], upgrades: [], commanderOracleIds: [], commanders: [], imported: true };

beforeEach(() => {
  state.user = fakeUser;
  localStorage.clear();
  importDeck.mockReset();
  getDeck.mockReset();
  listCollections.mockReset();
  listCollections.mockResolvedValue({ collections: [] });
  getDeck.mockResolvedValue({ deck: stored });
});

async function openAndPaste(list: string) {
  await renderAt("/decks");
  const user = userEvent.setup();
  await user.click(await screen.findByRole("button", { name: "Import a deck" }));
  await user.type(screen.getByRole("textbox", { name: "Deck list" }), list);
  await user.click(screen.getByRole("button", { name: "Import" }));
  return user;
}

describe("nameFromFile", () => {
  it("drops the extension and reads underscores as spaces", () => {
    expect(nameFromFile("living_weapon.txt")).toBe("living weapon");
    expect(nameFromFile("Turtle_power.txt")).toBe("Turtle power");
  });
});

describe("ImportDialog", () => {
  it("asks for the format of a list that is not Commander, then opens the deck (D-857)", async () => {
    importDeck.mockResolvedValueOnce({ needsFormat: true, commanderOptions: [], unresolved: [] });
    importDeck.mockResolvedValueOnce({ deck: stored, sessionId: "s9", commanderOptions: [], unresolved: [] });
    const user = await openAndPaste("4 Lightning Bolt");
    await user.click(await screen.findByRole("radio", { name: "Modern" }));
    await user.click(screen.getByRole("button", { name: "Import" }));
    await waitFor(() => expect(importDeck).toHaveBeenLastCalledWith(expect.objectContaining({ text: "4 Lightning Bolt", format: FormatId.MODERN })));
    await waitFor(() => expect(getDeck).toHaveBeenCalledWith({ deckId: "d9" }));
  });

  it("offers neither as the house format", async () => {
    importDeck.mockResolvedValueOnce({ needsFormat: true, commanderOptions: [], unresolved: [] });
    importDeck.mockResolvedValueOnce({ deck: stored, sessionId: "s9", commanderOptions: [], unresolved: [] });
    const user = await openAndPaste("4 Lightning Bolt");
    await user.click(await screen.findByRole("radio", { name: /Neither/ }));
    await user.click(screen.getByRole("button", { name: "Import" }));
    await waitFor(() => expect(importDeck).toHaveBeenLastCalledWith(expect.objectContaining({ format: FormatId.HOUSE })));
  });

  it("asks which card leads a list with no commander mark (D-847)", async () => {
    importDeck.mockResolvedValueOnce({
      needsFormat: false,
      commanderOptions: [
        { oracleId: "o-a", name: "Karlov of the Ghost Council" },
        { oracleId: "o-b", name: "Marrow-Gnawer" },
      ],
      unresolved: [],
    });
    importDeck.mockResolvedValueOnce({ deck: stored, sessionId: "s9", commanderOptions: [], unresolved: [] });
    const user = await openAndPaste("1 Karlov of the Ghost Council");
    await user.click(await screen.findByRole("radio", { name: "Marrow-Gnawer" }));
    await user.click(screen.getByRole("button", { name: "Import" }));
    await waitFor(() => expect(importDeck).toHaveBeenLastCalledWith(expect.objectContaining({ commanderOracleIds: ["o-b"] })));
  });

  it("lists each line that matched no card before the deck opens (D-846)", async () => {
    importDeck.mockResolvedValueOnce({ deck: stored, sessionId: "s9", commanderOptions: [], unresolved: [{ line: 4, raw: "1x No Such Card" }] });
    const user = await openAndPaste("1x No Such Card");
    expect(await screen.findByText("Line 4: 1x No Such Card")).toBeInTheDocument();
    expect(getDeck).not.toHaveBeenCalled();
    await user.click(screen.getByRole("button", { name: "Open the deck" }));
    await waitFor(() => expect(getDeck).toHaveBeenCalledWith({ deckId: "d9" }));
  });

  it("sends the collection the reader picked to show owned cards (D-849)", async () => {
    listCollections.mockResolvedValue({ collections: [{ id: "c1", name: "Binder", cardCount: 900 }] });
    importDeck.mockResolvedValueOnce({ deck: stored, sessionId: "s9", commanderOptions: [], unresolved: [] });
    await renderAt("/decks");
    const user = userEvent.setup();
    await user.click(await screen.findByRole("button", { name: "Import a deck" }));
    await user.type(screen.getByRole("textbox", { name: "Deck list" }), "4 Lightning Bolt");
    await user.selectOptions(await screen.findByRole("combobox", { name: "Show owned cards from this collection" }), "c1");
    await user.click(screen.getByRole("button", { name: "Import" }));
    await waitFor(() => expect(importDeck).toHaveBeenCalledWith(expect.objectContaining({ collectionId: "c1" })));
  });
});
