import { screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { signOut } from "firebase/auth";
import { beforeEach, describe, expect, it, vi } from "vitest";

import { fakeUser, state } from "../test-auth-state";
import { renderAt } from "../test-utils";

vi.mock("firebase/app");
vi.mock("firebase/auth");
vi.mock("../lib/api", () => ({
  healthClient: { check: () => Promise.resolve({ status: "ok", version: "test", cardSnapshot: "none" }) },
}));

function Boom(): never {
  throw new Error("the page broke");
}

beforeEach(() => {
  state.user = fakeUser;
  vi.mocked(signOut).mockReset();
  vi.mocked(signOut).mockResolvedValue();
});

describe("route error element", () => {
  it("shows a render error with a link to the collection and a sign-out", async () => {
    const assign = vi.fn();
    vi.stubGlobal("location", { ...window.location, assign });
    vi.spyOn(console, "error").mockImplementation(() => {});
    await renderAt("/boom", [{ path: "/boom", element: <Boom /> }]);
    expect(await screen.findByRole("alert")).toHaveTextContent("the page broke");
    expect(screen.getByRole("link", { name: "Go to your collection" })).toHaveAttribute("href", "/collection");
    await userEvent.setup().click(screen.getByRole("button", { name: "Sign out" }));
    expect(signOut).toHaveBeenCalled();
    vi.unstubAllGlobals();
    vi.restoreAllMocks();
  });
});
