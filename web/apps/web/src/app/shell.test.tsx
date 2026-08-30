import { screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { axe } from "jest-axe";
import { afterEach, beforeEach, describe, expect, it, vi } from "vitest";

import { applyTheme, readChoice, resolveTheme, themeStorageKey, useThemeStore } from "../lib/theme";
import { fakeUser, state } from "../test-auth-state";
import { renderAt } from "../test-utils";

vi.mock("firebase/app");
vi.mock("firebase/auth");
vi.mock("../lib/api", () => ({
  healthClient: { check: () => Promise.resolve({ status: "ok", version: "test", cardSnapshot: "2026-08-24T09:01:52Z", cardSnapshotAgeHours: 2.5 }) },
  collectionClient: { listCollections: () => Promise.resolve({ collections: [] }) },
  deckClient: { listDecks: () => Promise.resolve({ decks: [] }) },
  agentClient: { getSession: () => Promise.resolve({ session: { id: "abc123", turns: [], deckIds: [] } }) },
  cardClient: { getCards: () => Promise.resolve({ cards: [] }) },
}));

// matchMediaStub answers one query with a fixed result, so a test can play
// a phone or a dark-theme reader.
function matchMediaStub(match: (query: string) => boolean) {
  vi.stubGlobal("matchMedia", (query: string) => ({
    matches: match(query),
    media: query,
    addEventListener: () => {},
    removeEventListener: () => {},
    addListener: () => {},
    removeListener: () => {},
    dispatchEvent: () => false,
    onchange: null,
  }));
}

beforeEach(() => {
  state.user = fakeUser;
  localStorage.clear();
  document.documentElement.className = "";
  useThemeStore.setState({ choice: "system", theme: "light" });
});

afterEach(() => {
  vi.unstubAllGlobals();
  document.documentElement.className = "";
});

describe("the theme", () => {
  it("reads system when nothing is stored, and the stored choice after that", () => {
    expect(readChoice()).toBe("system");
    localStorage.setItem(themeStorageKey, "dark");
    expect(readChoice()).toBe("dark");
  });

  it("treats a damaged stored value as system", () => {
    localStorage.setItem(themeStorageKey, "purple");
    expect(readChoice()).toBe("system");
  });

  it("resolves system from the media query", () => {
    matchMediaStub((q) => q.includes("dark"));
    expect(resolveTheme("system")).toBe("dark");
    matchMediaStub(() => false);
    expect(resolveTheme("system")).toBe("light");
    expect(resolveTheme("dark")).toBe("dark");
  });

  it("puts the class and the color scheme on the root element", () => {
    applyTheme("dark");
    expect(document.documentElement).toHaveClass("dark");
    expect(document.documentElement.style.colorScheme).toBe("dark");
    applyTheme("light");
    expect(document.documentElement).not.toHaveClass("dark");
  });

  it("stores the choice from the sidebar menu and paints the root element", async () => {
    await renderAt("/decks");
    const user = userEvent.setup();
    await user.click(screen.getByRole("button", { name: /^Theme/ }));
    await user.click(await screen.findByRole("menuitemradio", { name: "Dark" }));
    expect(localStorage.getItem(themeStorageKey)).toBe("dark");
    expect(document.documentElement).toHaveClass("dark");
  });
});

describe("the shell", () => {
  it("shows the sidebar on a desktop and the bottom bar on a phone", async () => {
    matchMediaStub(() => false);
    const desktop = await renderAt("/decks");
    expect(screen.getByRole("navigation", { name: "Main" })).toBeInTheDocument();
    expect(screen.getByRole("link", { name: "MtG Deck Builder" })).toBeInTheDocument();
    desktop.unmount();

    matchMediaStub((q) => q.includes("max-width"));
    await renderAt("/decks");
    expect(screen.getByRole("navigation", { name: "Main" })).toBeInTheDocument();
    expect(screen.queryByRole("link", { name: "MtG Deck Builder" })).not.toBeInTheDocument();
    expect(screen.getByRole("button", { name: "Account menu" })).toBeInTheDocument();
  });

  it("has one Main landmark, never two", async () => {
    await renderAt("/decks");
    expect(screen.getAllByRole("navigation", { name: "Main" })).toHaveLength(1);
  });

  it("marks the entry that owns the path", async () => {
    await renderAt("/decks");
    expect(screen.getByRole("link", { name: "Decks" })).toHaveAttribute("aria-current", "page");
    expect(screen.getByRole("link", { name: "Collection" })).not.toHaveAttribute("aria-current");
  });

  it("shows no navigation to a signed-out visitor", async () => {
    state.user = null;
    await renderAt("/sign-in");
    expect(screen.queryByRole("navigation", { name: "Main" })).not.toBeInTheDocument();
  });
});

// The gate asks for axe on every route in both themes (PR-16).
const routes: [string, string][] = [
  ["/sign-in", "sign-in"],
  ["/collection", "collection"],
  ["/decks", "decks"],
  ["/session/new", "chat"],
];

describe.each(["light", "dark"] as const)("axe in the %s theme", (theme) => {
  it.each(routes)("passes on %s", async (path) => {
    state.user = path === "/sign-in" ? null : fakeUser;
    applyTheme(theme);
    const { container } = await renderAt(path);
    await screen.findByRole("heading", { level: 1 });
    expect(await axe(container)).toHaveNoViolations();
  });
});
