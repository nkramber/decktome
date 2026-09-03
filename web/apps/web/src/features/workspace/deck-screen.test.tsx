import { Code, ConnectError } from "@connectrpc/connect";
import { FormatId } from "@mtg/api-client/mtg/v1/format_pb";
import { screen, waitFor, within } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { axe } from "jest-axe";
import { beforeEach, describe, expect, it, vi } from "vitest";

import { fakeUser, state } from "../../test-auth-state";
import { renderAt } from "../../test-utils";

vi.mock("firebase/app");
vi.mock("firebase/auth");

const getDeck = vi.fn();
const getSession = vi.fn();
const chat = vi.fn();

type Ev = { event: { case: string; value: unknown } };
async function* events(list: Ev[]) {
  for (const e of list) {
    await Promise.resolve();
    yield e;
  }
}
const ev = (c: string, value: unknown): Ev => ({ event: { case: c, value } });
const updateDeck = vi.fn();
const deleteDeck = vi.fn();
const listDecks = vi.fn();
const shareDeck = vi.fn();
const revokeShare = vi.fn();
vi.mock("../../lib/api", () => ({
  healthClient: { check: () => Promise.resolve({ status: "ok", version: "test", cardSnapshot: "none" }) },
  collectionClient: { listCollections: () => Promise.resolve({ collections: [] }) },
  agentClient: {
    listSessions: () => Promise.resolve({ sessions: [], nextPageToken: "" }), getSession: (...a: unknown[]) => getSession(...a), chat: (...a: unknown[]) => chat(...a) },
  cardClient: { getCards: () => Promise.resolve({ cards: [], missingOracleIds: [] }) },
  deckClient: {
    listDecks: (...a: unknown[]) => listDecks(...a),
    getDeck: (...a: unknown[]) => getDeck(...a),
    updateDeck: (...a: unknown[]) => updateDeck(...a),
    deleteDeck: (...a: unknown[]) => deleteDeck(...a),
    shareDeck: (...a: unknown[]) => shareDeck(...a),
    revokeShare: (...a: unknown[]) => revokeShare(...a),
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
  getSession.mockReset();
  chat.mockReset();
  chat.mockReturnValue(events([]));
  getSession.mockResolvedValue({ session: { id: "s1", turns: [], deckIds: ["d1"] } });
  updateDeck.mockReset();
  deleteDeck.mockReset();
  listDecks.mockReset();
  listDecks.mockResolvedValue({ decks: [], nextPageToken: "" });
  getDeck.mockResolvedValue({ deck });
  updateDeck.mockResolvedValue({ deck });
  deleteDeck.mockResolvedValue({});
});

describe("DeckScreen", () => {
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

  it("makes a share link, shows it once, and revokes it (D-315)", async () => {
    shareDeck.mockReset();
    shareDeck.mockResolvedValue({ token: "t".repeat(43) });
    revokeShare.mockReset();
    revokeShare.mockResolvedValue({});
    await renderAt("/decks/d1");
    const user = userEvent.setup();
    await user.click(await screen.findByRole("button", { name: "Share" }));
    expect(await screen.findByTestId("share-state")).toHaveTextContent("This deck has no link yet.");
    // The share refreshes the deck, which then carries the shared mark.
    getDeck.mockResolvedValue({ deck: { ...deck, shared: true } });
    await user.click(screen.getByRole("button", { name: "Make a link" }));
    await waitFor(() => expect(shareDeck).toHaveBeenCalledWith({ deckId: "d1" }));
    const field = await screen.findByLabelText("The link, shown once");
    expect(field).toHaveValue(`${window.location.origin}/d/${"t".repeat(43)}`);
    // A shared deck offers the revoke.
    await user.click(screen.getByRole("button", { name: "Close" }));
    await user.click(await screen.findByRole("button", { name: "Shared" }));
    await user.click(await screen.findByRole("button", { name: "Revoke the link" }));
    await waitFor(() => expect(revokeShare).toHaveBeenCalledWith({ deckId: "d1" }));
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

// The revision turn of PR-12B now happens on the deck screen (D-335).
describe("a revision on the deck screen", () => {
  const revised = {
    ...deck,
    id: "d2",
    name: "Elf Ball, fewer elves",
    revisedFromDeckId: "d1",
    revisionNote: "I changed the count of Llanowar Elves: 4 to 2.",
    cards: [{ oracleId: "o-elf", name: "Llanowar Elves", count: 2 }],
  };

  it("shows the note and the diff against the deck it revised", async () => {
    getDeck.mockImplementation(({ deckId }: { deckId: string }) => Promise.resolve({ deck: deckId === "d2" ? revised : { ...deck, cards: [{ oracleId: "o-elf", name: "Llanowar Elves", count: 4 }] } }));
    await renderAt("/decks/d2");
    // The base deck arrives after the deck, so the diff comes second.
    expect(await screen.findByTestId("revision-diff")).toHaveTextContent("Count of Llanowar Elves: 4 to 2");
    expect(screen.getByTestId("revision-note")).toHaveTextContent("I changed the count of Llanowar Elves: 4 to 2.");
    // The base deck comes from the chain, not from the session's latest.
    expect(getDeck).toHaveBeenCalledWith({ deckId: "d1" });
  });

  it("moves to the address of the deck a turn builds", async () => {
    chat.mockReturnValue(events([ev("agentMessage", "done"), ev("deck", revised)]));
    const { router } = await renderAt("/decks/d1");
    const user = userEvent.setup();
    await user.type(await screen.findByLabelText("Your message"), "fewer elves");
    await user.click(screen.getByRole("button", { name: "Send" }));
    await waitFor(() => expect(router.state.location.pathname).toBe("/decks/d2"));
  });

  it("has no axe violations with the deck and its dock on screen", async () => {
    const { container } = await renderAt("/decks/d1");
    await screen.findByRole("heading", { name: "Elf Ball" });
    expect(await axe(container)).toHaveNoViolations();
  });
});

// The version history of a deck (PR-17). A revision turn writes a new
// deck in the same chat, so the decks of one chat are its versions.
const card = (name: string, count: number) => ({ oracleId: `o-${name}`, name, count, role: 0, owned: true, ownedCount: count, priceUsd: 0, reason: "" });
const v1 = { ...deck, id: "d0", name: "Elf Ball", createdAt: { seconds: 1755000000n, nanos: 0 }, cards: [card("Llanowar Elves", 4), card("Forest", 20)] };
const v2 = { ...deck, id: "d1", name: "Elf Ball, tuned", createdAt: { seconds: 1756000000n, nanos: 0 }, revisedFromDeckId: "d0", cards: [card("Llanowar Elves", 2), card("Forest", 20), card("Elvish Mystic", 3)] };

function withVersions() {
  // The listing answers newest first.
  listDecks.mockResolvedValue({ decks: [v2, v1], nextPageToken: "" });
  getDeck.mockImplementation((req: { deckId: string }) => Promise.resolve({ deck: req.deckId === "d0" ? v1 : v2 }));
}

describe("the version history", () => {
  it("stays away when the chat built one deck", async () => {
    listDecks.mockResolvedValue({ decks: [deck], nextPageToken: "" });
    await renderAt("/decks/d1");
    await screen.findByRole("heading", { name: "Elf Ball" });
    expect(screen.queryByRole("list", { name: "Versions" })).not.toBeInTheDocument();
  });

  it("asks for the decks of this chat alone", async () => {
    withVersions();
    await renderAt("/decks/d1");
    await screen.findByRole("list", { name: "Versions" });
    expect(listDecks).toHaveBeenCalledWith({ sessionId: "s1", pageSize: 100 });
  });

  it("lists every version oldest first, and marks the one on screen", async () => {
    withVersions();
    await renderAt("/decks/d1");
    const list = await screen.findByRole("list", { name: "Versions" });
    const buttons = within(list).getAllByRole("button");
    // The date reads in the runner's own zone, so the check builds it.
    const day = (seconds: bigint) => new Date(Number(seconds) * 1000).toLocaleDateString();
    expect(buttons.map((b) => b.textContent)).toEqual([`v1Elf Ball${day(1755000000n)}`, `v2Elf Ball, tuned${day(1756000000n)}now`]);
    expect(buttons[1]).toHaveAttribute("aria-current", "page");
    expect(buttons[1]).toBeDisabled();
  });

  it("opens an earlier version at its own address", async () => {
    withVersions();
    const { router } = await renderAt("/decks/d1");
    const list = await screen.findByRole("list", { name: "Versions" });
    await userEvent.setup().click(within(list).getAllByRole("button")[0]);
    expect(router.state.location.pathname).toBe("/decks/d0");
  });

  it("compares two versions and names every change", async () => {
    withVersions();
    await renderAt("/decks/d1");
    const user = userEvent.setup();
    await user.click(await screen.findByRole("button", { name: "Compare" }));
    // The default reads the version before the one on screen.
    expect((screen.getByLabelText("Compare from") as HTMLSelectElement).value).toBe("d0");
    expect((screen.getByLabelText("With") as HTMLSelectElement).value).toBe("d1");
    const diff = await screen.findByTestId("compare-diff");
    expect(within(diff).getAllByRole("listitem").map((i) => i.textContent)).toEqual([
      "Added 3 Elvish Mystic",
      "Count of Llanowar Elves: 4 to 2",
    ]);
  });

  it("says so when the two sides are the same version", async () => {
    withVersions();
    await renderAt("/decks/d1");
    const user = userEvent.setup();
    await user.click(await screen.findByRole("button", { name: "Compare" }));
    await user.selectOptions(screen.getByLabelText("Compare from"), "d1");
    expect(screen.getByText("Pick two different versions.")).toBeInTheDocument();
    expect(screen.queryByTestId("compare-diff")).not.toBeInTheDocument();
  });
});

// The docked chat has no fold: it holds one width and stays open
// (D-370, amends D-368).
describe("the docked chat", () => {
  it("is always open, and it scrolls on its own", async () => {
    withVersions();
    await renderAt("/decks/d1");
    expect(await screen.findByRole("textbox", { name: "Your message" })).toBeInTheDocument();
    expect(screen.queryByRole("button", { name: "Hide the chat" })).not.toBeInTheDocument();
    expect(screen.queryByRole("button", { name: "Show the chat" })).not.toBeInTheDocument();
    // The thread keeps its own scroll box, so the deck beside it is free
    // to scroll on its own.
    const thread = screen.getByRole("list", { name: "Conversation" }).parentElement;
    expect(thread).toHaveClass("overflow-y-auto");
  });
});
