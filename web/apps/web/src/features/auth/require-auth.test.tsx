import { screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { signOut } from "firebase/auth";
import { axe } from "jest-axe";
import { beforeEach, describe, expect, it, vi } from "vitest";

import { fakeUser, state } from "../../test-auth-state";
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

  it("shows the session placeholder with the id, and signs out from the header", async () => {
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
