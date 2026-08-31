import { FormatId, SixtyStep } from "@mtg/api-client/mtg/v1/format_pb";
import { screen, waitFor, within } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { axe } from "jest-axe";
import { beforeEach, describe, expect, it, vi } from "vitest";

import { fakeUser, state } from "../../test-auth-state";
import { renderAt } from "../../test-utils";
import { powerFrom } from "./decks-page";

vi.mock("firebase/app");
vi.mock("firebase/auth");

const listDecks = vi.fn();
const updateDeck = vi.fn();
const getCards = vi.fn();
vi.mock("../../lib/api", () => ({
  healthClient: { check: () => Promise.resolve({ status: "ok", version: "test", cardSnapshot: "none" }) },
  collectionClient: { listCollections: () => Promise.resolve({ collections: [] }) },
  agentClient: { getSession: vi.fn() },
  cardClient: { getCards: (...a: unknown[]) => getCards(...a) },
  deckClient: {
    listDecks: (...a: unknown[]) => listDecks(...a),
    updateDeck: (...a: unknown[]) => updateDeck(...a),
    deleteDeck: vi.fn(),
    getDeck: vi.fn(),
  },
}));

function deck(over: Record<string, unknown> = {}) {
  return {
    id: "d1",
    name: "Elf Ball",
    format: { id: FormatId.COMMANDER },
    cardCount: 99,
    buyCostUsd: 12.5,
    favorite: false,
    commanderOracleIds: ["o-marwyn"],
    cards: [],
    sideboard: [],
    upgrades: [],
    ...over,
  };
}

beforeEach(() => {
  state.user = fakeUser;
  localStorage.clear();
  listDecks.mockReset();
  updateDeck.mockReset();
  getCards.mockReset();
  listDecks.mockResolvedValue({ decks: [deck()], nextPageToken: "" });
  updateDeck.mockResolvedValue({ deck: deck({ favorite: true }) });
  getCards.mockResolvedValue({
    cards: [{ oracleId: "o-marwyn", name: "Marwyn, the Nurturer", colorIdentity: [5], faces: [], defaultPrinting: { artist: "A", imageUris: { artCrop: "https://x/a.jpg", normal: "https://x/n.jpg" } } }],
    missingOracleIds: [],
  });
});

describe("powerFrom", () => {
  it("reads one select value into the two arms of the power filter", () => {
    expect(powerFrom("b3")).toEqual({ powerBracket: 3, sixtyStep: SixtyStep.UNSPECIFIED });
    expect(powerFrom("s2")).toEqual({ powerBracket: 0, sixtyStep: SixtyStep.FNM });
    expect(powerFrom("")).toEqual({ powerBracket: 0, sixtyStep: SixtyStep.UNSPECIFIED });
  });
});

describe("DecksPage", () => {
  it("shows a deck with its name, format, count, and cost", async () => {
    await renderAt("/decks");
    expect(await screen.findByRole("link", { name: "Elf Ball" })).toHaveAttribute("href", "/decks/d1");
    const card = screen.getByRole("listitem");
    expect(within(card).getByText(/Commander/)).toBeInTheDocument();
    expect(within(card).getByText(/99 cards/)).toBeInTheDocument();
    expect(within(card).getByText(/to buy/)).toBeInTheDocument();
  });

  it("names the deck's color identity for a reader who sees no color", async () => {
    await renderAt("/decks");
    await screen.findByRole("link", { name: "Elf Ball" });
    expect(await screen.findByText("Green")).toBeInTheDocument();
  });

  it("sends the search text after a pause, and only the trimmed text", async () => {
    await renderAt("/decks");
    await screen.findByRole("link", { name: "Elf Ball" });
    await userEvent.setup().type(screen.getByRole("textbox", { name: /Search decks/ }), "  elf  ");
    await waitFor(() => expect(listDecks).toHaveBeenCalledWith(expect.objectContaining({ query: "elf" })));
  });

  it("sends the format, the power, and the favorites filter", async () => {
    await renderAt("/decks");
    await screen.findByRole("link", { name: "Elf Ball" });
    const user = userEvent.setup();
    await user.selectOptions(screen.getByRole("combobox", { name: "Filter by format" }), String(FormatId.MODERN));
    await waitFor(() => expect(listDecks).toHaveBeenCalledWith(expect.objectContaining({ format: FormatId.MODERN })));
    await user.selectOptions(screen.getByRole("combobox", { name: "Filter by power" }), "b3");
    await waitFor(() => expect(listDecks).toHaveBeenCalledWith(expect.objectContaining({ powerBracket: 3 })));
    await user.click(screen.getByRole("button", { name: "Favorites" }));
    await waitFor(() => expect(listDecks).toHaveBeenCalledWith(expect.objectContaining({ favorite: true })));
  });

  it("the star writes the favorite and the grid refreshes with no reload", async () => {
    await renderAt("/decks");
    const star = await screen.findByRole("button", { name: "Add Elf Ball to your favorites" });
    listDecks.mockResolvedValue({ decks: [deck({ favorite: true })], nextPageToken: "" });
    await userEvent.setup().click(star);
    await waitFor(() => expect(updateDeck).toHaveBeenCalledWith({ deckId: "d1", favorite: true }));
    expect(await screen.findByRole("button", { name: "Remove Elf Ball from your favorites" })).toBeInTheDocument();
  });

  it("shows the empty state with no decks, and the no-match state under a filter", async () => {
    listDecks.mockResolvedValue({ decks: [], nextPageToken: "" });
    await renderAt("/decks");
    expect(await screen.findByText("No decks yet.")).toBeInTheDocument();
    await userEvent.setup().click(screen.getByRole("button", { name: "Favorites" }));
    expect(await screen.findByText("No deck matches.")).toBeInTheDocument();
  });

  it("reads the next page with the token the server gave", async () => {
    listDecks.mockResolvedValueOnce({ decks: [deck()], nextPageToken: "t2" }).mockResolvedValueOnce({ decks: [deck({ id: "d2", name: "Goblin storm" })], nextPageToken: "" });
    await renderAt("/decks");
    await screen.findByRole("link", { name: "Elf Ball" });
    await userEvent.setup().click(screen.getByRole("button", { name: "Show more decks" }));
    expect(await screen.findByRole("link", { name: "Goblin storm" })).toBeInTheDocument();
    expect(listDecks).toHaveBeenLastCalledWith(expect.objectContaining({ pageToken: "t2" }));
  });

  it("has no axe violations", async () => {
    const { container } = await renderAt("/decks");
    await screen.findByRole("link", { name: "Elf Ball" });
    expect(await axe(container)).toHaveNoViolations();
  });
});

// The whole tile opens the deck (D-365). The stretched link used to
// resolve against the header row, so only the name was clickable.
describe("a deck tile", () => {
  it("covers the whole card with one link", async () => {
    await renderAt("/decks");
    await waitFor(() => expect(document.querySelector("li.card-hover")).not.toBeNull());
    const card = document.querySelector("li.card-hover");
    expect(card).not.toBeNull();
    // The tile is the one positioned ancestor, so the stretched link of
    // the title spans the whole card.
    expect(card).toHaveClass("relative");
    // A positioned row in between would take the link's ::after with it.
    // The favorite button is the one exception: it sits above the link.
    const positioned = [...(card as HTMLElement).children].filter(
      (el) => el.className.includes("relative") && !el.className.includes("z-10"),
    );
    expect(positioned.map((el) => el.tagName)).toEqual([]);
  });
});
