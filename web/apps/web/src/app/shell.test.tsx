import { screen } from "@testing-library/react";
import { axe } from "jest-axe";
import { beforeEach, describe, expect, it, vi } from "vitest";

import { fakeUser, state } from "../test-auth-state";
import { renderAt } from "../test-utils";

vi.mock("firebase/app");
vi.mock("firebase/auth");
vi.mock("../lib/api", () => ({
  healthClient: { check: () => Promise.resolve({ status: "ok", version: "test", cardSnapshot: "2026-08-24T09:01:52Z", cardSnapshotAgeHours: 2.5 }) },
  collectionClient: { listCollections: () => Promise.resolve({ collections: [] }), getCollection: vi.fn() },
  deckClient: { listDecks: () => Promise.resolve({ decks: [], nextPageToken: "" }), getDeck: vi.fn(), updateDeck: vi.fn(), deleteDeck: vi.fn() },
  agentClient: { getSession: () => Promise.resolve({ session: { id: "abc123", turns: [], deckIds: [] } }) },
  cardClient: { getCards: () => Promise.resolve({ cards: [], missingOracleIds: [] }) },
}));

beforeEach(() => {
  state.user = fakeUser;
  localStorage.clear();
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
