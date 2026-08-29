import type { User } from "firebase/auth";
import { signOut } from "firebase/auth";
import { describe, expect, it, vi } from "vitest";

import { fakeUser, state } from "../test-auth-state";
import { bearerInterceptor } from "./api";

vi.mock("firebase/app");
vi.mock("firebase/auth");

// The interceptor is a plain function over the request, so the test calls it directly.
function fakeRequest() {
  return { header: new Headers() } as unknown as Parameters<ReturnType<typeof bearerInterceptor>>[0];
}

describe("bearerInterceptor", () => {
  it("adds the ID token of the signed-in user (D-268)", async () => {
    state.user = fakeUser;
    const next = vi.fn((req) => Promise.resolve(req));
    const req = fakeRequest();
    await bearerInterceptor(next)(req);
    expect(req.header.get("Authorization")).toBe("Bearer token-1");
    expect(next).toHaveBeenCalledWith(req);
  });

  it("sends no header when nobody is signed in", async () => {
    state.user = null;
    const next = vi.fn((req) => Promise.resolve(req));
    const req = fakeRequest();
    await bearerInterceptor(next)(req);
    expect(req.header.get("Authorization")).toBeNull();
  });

  it("signs out when the token refresh says the session is dead", async () => {
    vi.mocked(signOut).mockClear();
    state.user = { ...fakeUser, getIdToken: () => Promise.reject({ code: "auth/user-token-expired" }) } as unknown as User;
    const next = vi.fn((req) => Promise.resolve(req));
    const req = fakeRequest();
    await bearerInterceptor(next)(req);
    expect(signOut).toHaveBeenCalledTimes(1);
    expect(req.header.get("Authorization")).toBeNull();
    expect(next).toHaveBeenCalledWith(req);
  });

  it("rethrows any other token error and keeps the session", async () => {
    vi.mocked(signOut).mockClear();
    state.user = { ...fakeUser, getIdToken: () => Promise.reject({ code: "auth/network-request-failed" }) } as unknown as User;
    const next = vi.fn((req) => Promise.resolve(req));
    await expect(bearerInterceptor(next)(fakeRequest())).rejects.toMatchObject({ code: "auth/network-request-failed" });
    expect(signOut).not.toHaveBeenCalled();
    state.user = null;
  });
});
