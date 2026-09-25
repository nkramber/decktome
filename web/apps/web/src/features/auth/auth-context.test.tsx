import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { act, render, screen } from "@testing-library/react";
import type { User } from "firebase/auth";
import { beforeEach, describe, expect, it, vi } from "vitest";

import { useAppStore } from "../../lib/store";
import { emit, fakeUser, state } from "../../test-auth-state";
import { AuthProvider, useAuth } from "./auth-context";
import { resetInviteState, setInviteState, useInviteState } from "./invite-state";

vi.mock("firebase/app");
vi.mock("firebase/auth");

const otherUser = { ...fakeUser, uid: "u2" } as unknown as User;

function Probe() {
  const { user, ready } = useAuth();
  const invite = useInviteState();
  return <p>{ready ? `user=${user?.uid ?? "none"} invite=${invite}` : "loading"}</p>;
}

// fillAccountState puts the data of one account in each place that
// sign-out clears: the query cache, the store, and the invite answer.
function fillAccountState(client: QueryClient) {
  client.setQueryData(["collections"], ["collection of u1"]);
  useAppStore.getState().setCollection("c1");
  useAppStore.getState().setSessionId("s1");
  setInviteState("invited");
}

async function renderProvider() {
  const client = new QueryClient();
  render(
    <QueryClientProvider client={client}>
      <AuthProvider>
        <Probe />
      </AuthProvider>
    </QueryClientProvider>,
  );
  return client;
}

beforeEach(() => {
  state.user = null;
  state.listeners.clear();
  localStorage.clear();
  useAppStore.getState().reset();
  resetInviteState();
});

// REV-039. A sign-out in another tab and a sign-in of a second account
// never ran signOutAndClear in this tab, so the second reader saw the
// data of the first until a reload.
describe("AuthProvider on a change of user", () => {
  it("keeps the cache on the first emission of a user", async () => {
    state.user = fakeUser;
    const client = await renderProvider();
    await screen.findByText(/user=u1/);
    fillAccountState(client);
    expect(client.getQueryData(["collections"])).toEqual(["collection of u1"]);
    expect(useAppStore.getState().collectionId).toBe("c1");
  });

  it("clears the data of user A when A signs out and B signs in", async () => {
    state.user = fakeUser;
    const client = await renderProvider();
    await screen.findByText(/user=u1/);
    fillAccountState(client);

    await act(async () => emit(null));
    await screen.findByText("user=none invite=unknown");
    await act(async () => emit(otherUser));
    await screen.findByText(/user=u2/);

    expect(client.getQueryData(["collections"])).toBeUndefined();
    expect(useAppStore.getState().collectionId).toBe("");
    expect(useAppStore.getState().sessionId).toBe("");
    expect(screen.getByText("user=u2 invite=unknown")).toBeInTheDocument();
  });

  it("clears the data of user A when B replaces A with no sign-out between", async () => {
    state.user = fakeUser;
    const client = await renderProvider();
    await screen.findByText(/user=u1/);
    fillAccountState(client);

    await act(async () => emit(otherUser));
    await screen.findByText(/user=u2/);

    expect(client.getQueryData(["collections"])).toBeUndefined();
    expect(useAppStore.getState().collectionId).toBe("");
    expect(screen.getByText("user=u2 invite=unknown")).toBeInTheDocument();
  });

  it("keeps the cache when the same user emits again", async () => {
    state.user = fakeUser;
    const client = await renderProvider();
    await screen.findByText(/user=u1/);
    fillAccountState(client);

    await act(async () => emit(fakeUser));

    expect(client.getQueryData(["collections"])).toEqual(["collection of u1"]);
    expect(useAppStore.getState().collectionId).toBe("c1");
  });
});
