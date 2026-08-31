import { screen, waitFor } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { axe } from "jest-axe";
import { beforeEach, describe, expect, it, vi } from "vitest";

import { useAppStore } from "../lib/store";
import { fakeUser, state } from "../test-auth-state";
import { renderAt } from "../test-utils";

vi.mock("firebase/app");
vi.mock("firebase/auth");

const listCollections = vi.fn();
vi.mock("../lib/api", () => ({
  healthClient: { check: () => Promise.resolve({ status: "ok", version: "test", cardSnapshot: "2026-08-24T09:01:52Z", cardSnapshotAgeHours: 2.5 }) },
  collectionClient: {
    listCollections: (...a: unknown[]) => listCollections(...a),
    getCollection: vi.fn(),
  },
  deckClient: { listDecks: () => Promise.resolve({ decks: [], nextPageToken: "" }), getDeck: vi.fn(), updateDeck: vi.fn(), deleteDeck: vi.fn() },
  agentClient: { getSession: () => Promise.resolve({ session: { id: "abc123", turns: [], deckIds: [] } }) },
  cardClient: { getCards: () => Promise.resolve({ cards: [], missingOracleIds: [] }) },
}));

beforeEach(() => {
  state.user = fakeUser;
  localStorage.clear();
  useAppStore.setState({ collectionId: "", sessionId: "", poolMode: "any" });
  listCollections.mockReset();
  listCollections.mockResolvedValue({ collections: [{ id: "c-old", name: "binder-july.csv", cardCount: 4317 }] });
});

describe("the shell", () => {
  it("carries one navigation, in the header (D-328)", async () => {
    await renderAt("/decks");
    expect(screen.getAllByRole("navigation", { name: "Main" })).toHaveLength(1);
    expect(screen.getByRole("link", { name: /MtG Deck Builder/ })).toBeInTheDocument();
  });

  it("marks the entry that owns the path", async () => {
    await renderAt("/decks");
    expect(screen.getByRole("link", { name: "Decks" })).toHaveAttribute("aria-current", "page");
    expect(screen.getByRole("link", { name: "Collection" })).not.toHaveAttribute("aria-current");
  });

  it("shows the account menu to a signed-in reader", async () => {
    await renderAt("/decks");
    expect(screen.getByRole("button", { name: "Account menu" })).toBeInTheDocument();
  });

  it("shows no navigation to a signed-out visitor", async () => {
    state.user = null;
    await renderAt("/sign-in");
    expect(screen.queryByRole("navigation", { name: "Main" })).not.toBeInTheDocument();
    expect(screen.queryByRole("button", { name: "Account menu" })).not.toBeInTheDocument();
  });
});

// Dark is the only theme now (D-330), so axe runs once per route.
const routes: [string, string][] = [
  ["/sign-in", "sign-in"],
  ["/collection", "collection"],
  ["/decks", "decks"],
  ["/session/new", "chat"],
];

describe("axe", () => {
  it.each(routes)("passes on %s", async (path) => {
    state.user = path === "/sign-in" ? null : fakeUser;
    const { container } = await renderAt(path);
    await screen.findByRole("heading", { level: 1 });
    expect(await axe(container)).toHaveNoViolations();
  });
});

// Build asks which cards the deck may use before it opens a chat.
describe("the Build menu", () => {
  it("offers any card and every collection, and marks the active one", async () => {
    const user = userEvent.setup();
    await renderAt("/decks");
    await user.click(screen.getByRole("button", { name: "Build" }));
    // The collections arrive after the menu, so wait for the last row.
    await screen.findByRole("menuitemradio", { name: /binder-july/ });
    const items = screen.getAllByRole("menuitemradio");
    expect(items.map((i) => i.textContent)).toEqual(["Any card", "binder-july.csv4317"]);
    expect(items[0]).toHaveAttribute("aria-checked", "true");
  });

  it("takes the collection, sets the owned pool, and opens a chat (D-37)", async () => {
    const user = userEvent.setup();
    const { router } = await renderAt("/decks");
    await user.click(screen.getByRole("button", { name: "Build" }));
    await user.click(await screen.findByRole("menuitemradio", { name: /binder-july/ }));
    expect(useAppStore.getState().collectionId).toBe("c-old");
    expect(useAppStore.getState().poolMode).toBe("owned");
    expect(router.state.location.pathname).toBe("/session/new");
  });

  it("any card clears the collection", async () => {
    useAppStore.setState({ collectionId: "c-old", poolMode: "owned" });
    const user = userEvent.setup();
    await renderAt("/decks");
    await user.click(screen.getByRole("button", { name: "Build" }));
    await user.click(await screen.findByRole("menuitemradio", { name: "Any card" }));
    expect(useAppStore.getState().collectionId).toBe("");
    expect(useAppStore.getState().poolMode).toBe("any");
  });
});

describe("a menu trigger", () => {
  // Radix places its panel from the trigger's rectangle, and it reaches
  // the trigger through the ref it passes as a prop. A trigger that
  // keeps only the props it names drops that ref, and the panel lands
  // outside the window. jsdom has no layout, so the check reads the
  // state Radix writes onto the trigger instead.
  // An open menu marks the rest of the page aria-hidden, so a role query
  // no longer reaches the trigger. The check reads the DOM instead.
  const openTrigger = () => document.querySelector('[aria-haspopup="menu"][aria-expanded="true"]');

  it("takes the props Radix gives it, on Build", async () => {
    const user = userEvent.setup();
    await renderAt("/decks");
    await user.click(screen.getByRole("button", { name: "Build" }));
    await screen.findByRole("menuitemradio", { name: "Any card" });
    expect(openTrigger()).not.toBeNull();
  });

  it("takes the props Radix gives it, on the account menu", async () => {
    const user = userEvent.setup();
    await renderAt("/decks");
    await user.click(screen.getByRole("button", { name: "Account menu" }));
    await screen.findByRole("menuitem", { name: "Sign out" });
    expect(openTrigger()).not.toBeNull();
  });
});

describe("the Build menu with no collection", () => {
  it("offers the way to add one", async () => {
    listCollections.mockResolvedValueOnce({ collections: [] });
    const user = userEvent.setup();
    const { router } = await renderAt("/decks");
    await user.click(screen.getByRole("button", { name: "Build" }));
    await user.click(await screen.findByRole("menuitem", { name: "Add a collection" }));
    expect(router.state.location.pathname).toBe("/collection");
  });
});

// The menus mount closed once the idle warm-up brings their chunk in.
// A click then opens them at once, and a second click closes them.
describe("a menu after the warm-up", () => {
  it("opens and closes with no reload", async () => {
    const { warmChunks } = await import("./chunks");
    warmChunks();
    const user = userEvent.setup();
    await renderAt("/decks");
    // A mounted menu marks its trigger closed. A trigger with no such
    // mark is the plain button of the state before the chunk lands.
    await waitFor(() => expect(screen.getByRole("button", { name: "Build" })).toHaveAttribute("aria-expanded", "false"));
    await user.click(screen.getByRole("button", { name: "Build" }));
    await screen.findByRole("menuitemradio", { name: "Any card" });
    await user.keyboard("{Escape}");
    await waitFor(() => expect(screen.queryByRole("menuitemradio", { name: "Any card" })).not.toBeInTheDocument());
    await user.click(screen.getByRole("button", { name: "Build" }));
    expect(await screen.findByRole("menuitemradio", { name: "Any card" })).toBeInTheDocument();
  });
});
