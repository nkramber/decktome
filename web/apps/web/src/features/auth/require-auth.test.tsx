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
  agentClient: {
    listSessions: () => Promise.resolve({ sessions: [], nextPageToken: "" }), getSession: () => Promise.resolve({ session: { id: "abc123", turns: [], deckIds: [] } }) },
}));

beforeEach(() => {
  state.user = null;
  localStorage.clear();
  vi.mocked(signOut).mockReset();
  vi.mocked(signOut).mockResolvedValue();
});

describe("route guard", () => {
  it("sends an unauthenticated visit to /collection to /sign-in", async () => {
    const { router } = await renderAt("/collection");
    expect(await screen.findByRole("heading", { level: 1, name: "Sign in" })).toBeInTheDocument();
    expect(router.state.location.pathname).toBe("/sign-in");
  });

  it("sends / to /sign-in when signed out and to Build when signed in (D-334)", async () => {
    const out = await renderAt("/");
    expect(await screen.findByRole("heading", { level: 1, name: "Sign in" })).toBeInTheDocument();
    expect(out.router.state.location.pathname).toBe("/sign-in");
    out.unmount();

    state.user = fakeUser;
    const inn = await renderAt("/");
    expect(await screen.findByRole("heading", { level: 1, name: "New deck" })).toBeInTheDocument();
    expect(inn.router.state.location.pathname).toBe("/session/new");
  });

  it("sends a signed-in visit to /sign-in on to Build (D-334)", async () => {
    state.user = fakeUser;
    const { router } = await renderAt("/sign-in");
    await screen.findByRole("heading", { level: 1, name: "New deck" });
    expect(router.state.location.pathname).toBe("/session/new");
  });

  it("returns to the page the guard redirected after the sign-in", async () => {
    const { router } = await renderAt("/session/abc123");
    await screen.findByRole("heading", { level: 1, name: "Sign in" });
    expect(router.state.location.pathname).toBe("/sign-in");
    await act(async () => emit(fakeUser));
    expect(await screen.findByTestId("session-id")).toHaveTextContent("Session id: abc123");
    expect(router.state.location.pathname).toBe("/session/abc123");
  });

  it("shows the loading line on / and under the guard until auth is ready", async () => {
    vi.mocked(onAuthStateChanged).mockImplementationOnce(() => () => {});
    await renderAt("/");
    expect(screen.getByText("Loading your session...")).toBeInTheDocument();
  });

  it("an auth listener error ends the wait and shows the message", async () => {
    vi.mocked(onAuthStateChanged).mockImplementationOnce((_auth, next, error) => {
      state.listeners.add({ next: next as (u: User | null) => void, error: error as (e: Error) => void });
      return () => {};
    });
    await renderAt("/collection");
    expect(screen.getByText("Loading your session...")).toBeInTheDocument();
    await act(async () => emitError(new Error("auth is down")));
    expect(await screen.findByRole("heading", { level: 1, name: "Sign in" })).toBeInTheDocument();
    expect(screen.getByRole("alert")).toHaveTextContent("auth is down");
  });

  it("reports a sign-out failure as a toast", async () => {
    state.user = fakeUser;
    vi.mocked(signOut).mockRejectedValue(new Error("network down"));
    await renderAt("/decks");
    await screen.findByText("No decks yet.");
    const user = userEvent.setup();
    await user.click(screen.getByRole("button", { name: "Account menu" }));
    await user.click(await screen.findByRole("menuitem", { name: "Sign out" }));
    expect(await screen.findByText("Sign-out failed")).toBeInTheDocument();
    expect(await screen.findByText("network down")).toBeInTheDocument();
  });

  it("shows the stored session's id while it loads, and signs out from the account menu", async () => {
    state.user = fakeUser;
    await renderAt("/session/abc123");
    expect(await screen.findByTestId("session-id")).toHaveTextContent("Session id: abc123");
    const user = userEvent.setup();
    await user.click(screen.getByRole("button", { name: "Account menu" }));
    await user.click(await screen.findByRole("menuitem", { name: "Sign out" }));
    expect(vi.mocked(signOut)).toHaveBeenCalled();
  });

  // A new chat says nothing of its session, because it has none (D-356).
  it("opens a new chat with no session line", async () => {
    state.user = fakeUser;
    await renderAt("/session/new");
    expect(await screen.findByRole("heading", { level: 1, name: "New deck" })).toBeInTheDocument();
    expect(screen.queryByTestId("session-id")).not.toBeInTheDocument();
  });

  // D-626: no bar of the app carries the API status or the snapshot age.
  // A deck names the card data it rests on, on the deck itself.
  it("lists decks, empty, and carries no health bar", async () => {
    state.user = fakeUser;
    const { container } = await renderAt("/decks");
    expect(await screen.findByText("No decks yet.")).toBeInTheDocument();
    expect(screen.queryByTestId("health")).not.toBeInTheDocument();
    expect(screen.queryByTestId("freshness")).not.toBeInTheDocument();
    expect(await axe(container)).toHaveNoViolations();
  });

  it("session page has no axe violations", async () => {
    state.user = fakeUser;
    const { container } = await renderAt("/session/new");
    // The heading is the settled page. The health bar was the second
    // signal until D-626, and the bar is gone.
    await screen.findByRole("heading", { level: 1 });
    expect(await axe(container)).toHaveNoViolations();
  });
});
