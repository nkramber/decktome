import { Code, ConnectError } from "@connectrpc/connect";
import { screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { signInWithCustomToken } from "firebase/auth";
import { axe } from "jest-axe";
import { afterEach, beforeEach, describe, expect, it, vi } from "vitest";

import { state } from "../../test-auth-state";
import { renderAt } from "../../test-utils";

vi.mock("firebase/app");
vi.mock("firebase/auth");

const openLink = vi.fn();
vi.mock("../../lib/api", () => ({
  healthClient: { check: () => Promise.resolve({ status: "ok", version: "test", cardSnapshot: "none" }) },
  collectionClient: { listCollections: () => Promise.resolve({ collections: [] }) },
  agentClient: { listSessions: () => Promise.resolve({ sessions: [], nextPageToken: "" }) },
  deckClient: { listDecks: () => Promise.resolve({ decks: [], nextPageToken: "" }) },
  proofClient: { openLink: (...a: unknown[]) => openLink(...a) },
}));

const iPhone = "Mozilla/5.0 (iPhone; CPU iPhone OS 26_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/26.0 Mobile/15E148 Safari/604.1";

beforeEach(() => {
  state.user = null;
  openLink.mockReset();
  openLink.mockResolvedValue({ customToken: "tok" });
  vi.mocked(signInWithCustomToken).mockClear();
});

afterEach(() => {
  vi.restoreAllMocks();
});

// confirm selects the button that uses the link (D-1119).
async function confirm() {
  await userEvent.setup().click(await screen.findByRole("button", { name: "Confirm my email" }));
}

// D-1081 to D-1083: the short link of the email proves the address and
// signs in the browser that opens it. D-1119: the page uses the link on
// a click alone. D-1147: iOS Safari goes on to the app as well.
describe("ProofLinkPage", () => {
  it("uses the link on the click alone", async () => {
    await renderAt("/v/Load000001");
    expect(await screen.findByRole("button", { name: "Confirm my email" })).toBeEnabled();
    expect(openLink).not.toHaveBeenCalled();
    await confirm();
    await vi.waitFor(() => expect(openLink).toHaveBeenCalledTimes(1));
  });

  it("signs in and goes on to the app in a desktop browser", async () => {
    const { router } = await renderAt("/v/Desk000001");
    await confirm();
    await vi.waitFor(() => expect(router.state.location.pathname).not.toBe("/v/Desk000001"));
    expect(openLink).toHaveBeenCalledWith({ code: "Desk000001" });
    expect(vi.mocked(signInWithCustomToken)).toHaveBeenCalledWith(expect.anything(), "tok");
  });

  it("goes on to the app in an iPhone browser too", async () => {
    vi.spyOn(navigator, "userAgent", "get").mockReturnValue(iPhone);
    const { router } = await renderAt("/v/Ios0000001");
    await confirm();
    await vi.waitFor(() => expect(router.state.location.pathname).not.toBe("/v/Ios0000001"));
    expect(vi.mocked(signInWithCustomToken)).toHaveBeenCalledWith(expect.anything(), "tok");
    expect(screen.queryByText(/home screen/i)).not.toBeInTheDocument();
  });

  it("sends the person to sign in when the answer carries no token", async () => {
    openLink.mockResolvedValue({ customToken: "" });
    await renderAt("/v/NoTok00001");
    await confirm();
    expect(await screen.findByText("Your email is verified. Sign in to continue.")).toBeInTheDocument();
    expect(screen.getByRole("button", { name: "Return to login page" })).toBeInTheDocument();
    expect(vi.mocked(signInWithCustomToken)).not.toHaveBeenCalled();
  });

  it("says that the email is verified when the link worked before", async () => {
    openLink.mockResolvedValue({ customToken: "", alreadyProved: true });
    await renderAt("/v/Again00001");
    await confirm();
    expect(await screen.findByText("Your email is already verified. Sign in to continue.")).toBeInTheDocument();
    expect(screen.getByRole("button", { name: "Return to login page" })).toBeInTheDocument();
    expect(vi.mocked(signInWithCustomToken)).not.toHaveBeenCalled();
  });

  it("shows the sentence of a refusal and a way back to sign in", async () => {
    openLink.mockRejectedValue(new ConnectError("this link expired. Sign in to send a new link", Code.FailedPrecondition));
    await renderAt("/v/Late000001");
    await confirm();
    expect(await screen.findByRole("alert")).toHaveTextContent("This link expired. Sign in to send a new link.");
    expect(screen.getByRole("heading", { level: 1, name: "This link does not work" })).toBeInTheDocument();
    expect(screen.getByRole("button", { name: "Return to login page" })).toBeInTheDocument();
    expect(vi.mocked(signInWithCustomToken)).not.toHaveBeenCalled();
  });

  it("has no axe violations", async () => {
    openLink.mockResolvedValue({ customToken: "" });
    const { container } = await renderAt("/v/Axe0000001");
    expect(await axe(container)).toHaveNoViolations();
    await confirm();
    await screen.findByText(/Your email is verified/);
    expect(await axe(container)).toHaveNoViolations();
  });
});
