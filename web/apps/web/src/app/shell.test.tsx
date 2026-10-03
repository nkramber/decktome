import { readFileSync } from "node:fs";
import path from "node:path";

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
  // The navigation waits for the API to clear the reader (F-59, D-590),
  // so it arrives one hop after the render.
  it("carries one navigation, in the header (D-328)", async () => {
    await renderAt("/decks");
    await screen.findByRole("navigation", { name: "Main" });
    expect(screen.getAllByRole("navigation", { name: "Main" })).toHaveLength(1);
    expect(screen.getByRole("link", { name: /Deck Tome/ })).toBeInTheDocument();
  });

  it("marks the entry that owns the path", async () => {
    await renderAt("/decks");
    expect(await screen.findByRole("link", { name: "Decks" })).toHaveAttribute("aria-current", "page");
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

// iOS 26 and later blur the top of an installed app unless a fixed or
// sticky opaque box covers the top edge, and the buffer sits inside that
// box (F-187, D-1010). jsdom reads no media query, so the second test
// reads the stylesheet.
describe("the edges of the installed app", () => {
  it("has a sticky, opaque header that holds the top buffer", async () => {
    await renderAt("/decks");
    const header = (await screen.findByRole("navigation", { name: "Main" })).closest("header");
    expect(header).toHaveClass("sticky", "top-0", "bg-muted", "pt-[calc(0.75rem+var(--edge-top))]");
    expect(header?.parentElement).toHaveClass("pb-(--edge-bottom)");
  });

  it("sets a 16px buffer in the installed app alone", () => {
    const css = readFileSync(path.join(import.meta.dirname, "..", "index.css"), "utf8");
    expect(css).toMatch(/:root \{\s*--edge-top: 0px;\s*--edge-bottom: 0px;\s*\}/);
    expect(css).toMatch(/@media \(display-mode: standalone\) \{\s*:root \{\s*--edge-top: 16px;\s*--edge-bottom: 16px;\s*\}/);
  });
  // A toast on a phone shows under the header, at the height the layout
  // writes (F-194, D-1030).
  it("writes the height of the header for the toaster", async () => {
    const height = vi.spyOn(HTMLElement.prototype, "offsetHeight", "get").mockReturnValue(97);
    try {
      await renderAt("/decks");
      await screen.findByRole("navigation", { name: "Main" });
      // The header and the feedback row under it, 97 pixels each in
      // this test, so a toast never covers the row (D-1078).
      expect(screen.getByRole("button", { name: "Leave feedback" })).toBeInTheDocument();
      expect(document.documentElement.style.getPropertyValue("--header-height")).toBe("194px");
    } finally {
      height.mockRestore();
    }
  });
});

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
