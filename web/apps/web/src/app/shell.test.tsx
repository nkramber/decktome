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
  agentClient: {
    listSessions: () => Promise.resolve({ sessions: [], nextPageToken: "" }), getSession: () => Promise.resolve({ session: { id: "abc123", turns: [], deckIds: [] } }) },
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
    expect(screen.getByRole("link", { name: /Deck Tome/ })).toBeInTheDocument();
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

// Build keeps the pool the reader chose (D-580, F-63). A click of Build
// dropped the collection before, and the reader read no word of it.
describe("Build in the header", () => {
  it("is a link, and it keeps the collection the reader chose", async () => {
    useAppStore.setState({ collectionId: "c-old", poolMode: "owned_only" });
    const user = userEvent.setup();
    const { router } = await renderAt("/decks");
    const build = screen.getByRole("link", { name: "Build" });
    expect(build).not.toHaveAttribute("aria-haspopup");
    await user.click(build);
    expect(router.state.location.pathname).toBe("/session/new");
    expect(useAppStore.getState().collectionId).toBe("c-old");
    expect(useAppStore.getState().poolMode).toBe("owned_only");
  });
});

describe("a menu trigger", () => {
  // Radix places its panel from the trigger's rectangle, and it reaches
  // the trigger through the ref it passes as a prop. A trigger that
  // keeps only the props it names drops that ref, and the panel lands
  // outside the window. jsdom has no layout, so the check reads the
  // state Radix writes onto the trigger instead.
  it("takes the props Radix gives it, on the account menu", async () => {
    const user = userEvent.setup();
    await renderAt("/decks");
    await user.click(screen.getByRole("button", { name: "Account menu" }));
    await screen.findByRole("menuitem", { name: "Sign out" });
    expect(document.querySelector('[aria-haspopup="menu"][aria-expanded="true"]')).not.toBeNull();
  });
});

// The menus mount closed once the idle warm-up brings their chunk in.
describe("a menu after the warm-up", () => {
  it("opens and closes with no reload", async () => {
    const { warmChunks } = await import("./chunks");
    warmChunks();
    const user = userEvent.setup();
    await renderAt("/decks");
    await waitFor(() => expect(screen.getByRole("button", { name: "Account menu" })).toHaveAttribute("aria-expanded", "false"));
    await user.click(screen.getByRole("button", { name: "Account menu" }));
    await screen.findByRole("menuitem", { name: "Sign out" });
    await user.keyboard("{Escape}");
    await waitFor(() => expect(screen.queryByRole("menuitem", { name: "Sign out" })).not.toBeInTheDocument());
  });
});
