import { act, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { onAuthStateChanged, signOut, type User } from "firebase/auth";
import { axe } from "jest-axe";
import { beforeEach, describe, expect, it, vi } from "vitest";

import { emit, emitError, fakeUser, state } from "../../test-auth-state";
import { renderAt } from "../../test-utils";

vi.mock("firebase/app");
vi.mock("firebase/auth");
vi.mock("../../lib/api", () => ({
  healthClient: {
    check: () =>
      Promise.resolve({ status: "ok", version: "test", cardSnapshot: "2026-08-24T09:01:52Z", cardSnapshotAgeHours: 2.5 }),
  },
  collectionClient: { listCollections: () => Promise.resolve({ collections: [] }) },
  deckClient: { listDecks: () => Promise.resolve({ decks: [] }) },
  agentClient: { getSession: () => Promise.resolve({ session: { id: "abc123", turns: [], deckIds: [] } }) },
}));

beforeEach(() => {
  state.user = null;
  localStorage.clear();
  vi.mocked(signOut).mockReset();
  vi.mocked(signOut).mockResolvedValue();
});

describe("route guard", () => {
  it("sends an unauthenticated visit to /collection to /sign-in", async () => {
    const { router } = renderAt("/collection");
    expect(await screen.findByRole("heading", { level: 1, name: "Sign in" })).toBeInTheDocument();
    expect(router.state.location.pathname).toBe("/sign-in");
  });

  it("sends / to /sign-in when signed out and to /collection when signed in", async () => {
    const out = renderAt("/");
    expect(await screen.findByRole("heading", { level: 1, name: "Sign in" })).toBeInTheDocument();
    expect(out.router.state.location.pathname).toBe("/sign-in");
    out.unmount();

    state.user = fakeUser;
    const inn = renderAt("/");
    expect(await screen.findByRole("heading", { level: 1, name: "Your collection" })).toBeInTheDocument();
    expect(inn.router.state.location.pathname).toBe("/collection");
  });

  it("sends a signed-in visit to /sign-in on to /collection", async () => {
    state.user = fakeUser;
    const { router } = renderAt("/sign-in");
    await screen.findByRole("heading", { level: 1, name: "Your collection" });
    expect(router.state.location.pathname).toBe("/collection");
  });

  it("returns to the page the guard redirected after the sign-in", async () => {
    const { router } = renderAt("/session/abc123");
    await screen.findByRole("heading", { level: 1, name: "Sign in" });
    expect(router.state.location.pathname).toBe("/sign-in");
    await act(async () => emit(fakeUser));
    expect(await screen.findByTestId("session-id")).toHaveTextContent("Session id: abc123");
    expect(router.state.location.pathname).toBe("/session/abc123");
  });

  it("shows the loading line on / and under the guard until auth is ready", () => {
    vi.mocked(onAuthStateChanged).mockImplementationOnce(() => () => {});
    renderAt("/");
    expect(screen.getByText("Loading your session...")).toBeInTheDocument();
  });

  it("an auth listener error ends the wait and shows the message", async () => {
    vi.mocked(onAuthStateChanged).mockImplementationOnce((_auth, next, error) => {
      state.listeners.add({ next: next as (u: User | null) => void, error: error as (e: Error) => void });
      return () => {};
    });
    renderAt("/collection");
    expect(screen.getByText("Loading your session...")).toBeInTheDocument();
    await act(async () => emitError(new Error("auth is down")));
    expect(await screen.findByRole("heading", { level: 1, name: "Sign in" })).toBeInTheDocument();
    expect(screen.getByRole("alert")).toHaveTextContent("auth is down");
  });

  it("shows a sign-out failure in the header", async () => {
    state.user = fakeUser;
    vi.mocked(signOut).mockRejectedValue(new Error("network down"));
    renderAt("/decks");
    await screen.findByText("No decks yet.");
    await userEvent.setup().click(screen.getByRole("button", { name: "Sign out" }));
    expect(await screen.findByRole("alert")).toHaveTextContent("Sign-out failed: network down");
  });

  it("shows the stored session's id while it loads, and signs out from the header", async () => {
    state.user = fakeUser;
    renderAt("/session/abc123");
    expect(await screen.findByTestId("session-id")).toHaveTextContent("Session id: abc123");
    await userEvent.setup().click(screen.getByRole("button", { name: "Sign out" }));
    expect(vi.mocked(signOut)).toHaveBeenCalled();
  });

  it("treats /session/new as no session yet", async () => {
    state.user = fakeUser;
    renderAt("/session/new");
    expect(await screen.findByTestId("session-id")).toHaveTextContent("No session yet.");
  });

  it("lists decks, empty, and shows the health footer", async () => {
    state.user = fakeUser;
    const { container } = renderAt("/decks");
    expect(await screen.findByText("No decks yet.")).toBeInTheDocument();
    expect(await screen.findByTestId("health")).toHaveTextContent("API: ok, version test");
    expect(screen.getByTestId("freshness")).toHaveTextContent("Card data as of 2026-08-24T09:01:52Z (2.5 h old)");
    expect(await axe(container)).toHaveNoViolations();
  });

  it("session page has no axe violations", async () => {
    state.user = fakeUser;
    const { container } = renderAt("/session/new");
    await screen.findByTestId("session-id");
    await screen.findByText(/API: ok/);
    expect(await axe(container)).toHaveNoViolations();
  });
});
