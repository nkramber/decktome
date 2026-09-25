import { Code, ConnectError } from "@connectrpc/connect";
import { act, renderHook, screen, within } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { sendEmailVerification, signOut } from "firebase/auth";
import { axe } from "jest-axe";
import { beforeEach, describe, expect, it, vi } from "vitest";

import { fakeUser, state } from "../../test-auth-state";
import { renderAt } from "../../test-utils";
import { isNotInvited, isUnverified } from "./invite-gate";
import { resetInviteState, setInviteState, useInviteState } from "./invite-state";
import { signOutAndClear } from "./sign-out";

vi.mock("firebase/app");
vi.mock("firebase/auth");

const listCollections = vi.fn();

vi.mock("../../lib/api", () => ({
  healthClient: {
    check: () => Promise.resolve({ status: "ok", version: "test", cardSnapshot: "2026-08-24T09:01:52Z", cardSnapshotAgeHours: 2.5 }),
  },
  collectionClient: { listCollections: () => listCollections() },
  deckClient: { listDecks: () => Promise.resolve({ decks: [] }) },
  agentClient: {
    listSessions: () => Promise.resolve({ sessions: [], nextPageToken: "" }),
    getSession: () => Promise.resolve({ session: { id: "abc123", turns: [], deckIds: [] } }),
  },
}));

// refusal is what the API answers a caller off the invite list (D-314).
// The state rides in the header, and never in the sentence (F-59).
function refusal() {
  return new ConnectError("this app is open to invited users alone, and your email is not on the list", Code.PermissionDenied, {
    "Deck-Tome-Refusal": "not-invited",
  });
}

beforeEach(() => {
  // The gate latches its answer in a module, so each test starts with
  // none of the last one.
  resetInviteState();
  state.user = fakeUser;
  localStorage.clear();
  listCollections.mockReset();
  listCollections.mockResolvedValue({ collections: [] });
  vi.mocked(signOut).mockReset();
  vi.mocked(signOut).mockResolvedValue();
});

describe("isNotInvited", () => {
  it("reads the code, and the header confirms it", () => {
    expect(isNotInvited(refusal())).toBe(true);
    expect(isNotInvited(new ConnectError("nope", Code.NotFound, { "Deck-Tome-Refusal": "not-invited" }))).toBe(false);
    expect(isNotInvited(new ConnectError("firestore down", Code.Unavailable))).toBe(false);
    expect(isNotInvited(new Error("network down"))).toBe(false);
  });

  // The regression of F-59. `decktome.com` and the API are two origins,
  // and a browser reads no header the API does not expose. Every refusal
  // reaches the browser bare, and the first fix read it as an invitation.
  it("reads a refusal that carries no header, as a browser gets it", () => {
    expect(isNotInvited(new ConnectError("this app is open to invited users alone", Code.PermissionDenied))).toBe(true);
  });

  // A header that names another state still discriminates.
  it("refuses nothing when the header names another state", () => {
    expect(isNotInvited(new ConnectError("no", Code.PermissionDenied, { "Deck-Tome-Refusal": "other" }))).toBe(false);
  });
});

// F-59: before the fix a reader off the list read the header, the
// navigation, the prompt box, and a Build button that always failed.
describe("the invite gate", () => {
  it("shows the refusal screen, and no navigation, to a reader off the list", async () => {
    listCollections.mockRejectedValue(refusal());
    await renderAt("/session/new");
    const heading = await screen.findByRole("heading", { level: 1, name: "You are not on the invite list" });
    // The screen names the account that is signed in. The account menu of
    // the header carries the same address, so the match is on the card.
    expect(within(heading.closest("div")?.parentElement ?? document.body).getByText(fakeUser.email ?? "")).toBeInTheDocument();
    // No page, and no way to start a build.
    expect(screen.queryByRole("navigation", { name: "Main" })).not.toBeInTheDocument();
    expect(screen.queryByLabelText("Your message")).not.toBeInTheDocument();
    expect(screen.queryByRole("heading", { level: 1, name: "New deck" })).not.toBeInTheDocument();
  });

  // The shape a browser really delivers: the code alone, no metadata.
  it("shows the refusal screen when the refusal carries no header", async () => {
    listCollections.mockRejectedValue(new ConnectError("this app is open to invited users alone", Code.PermissionDenied));
    await renderAt("/session/new");
    expect(await screen.findByRole("heading", { level: 1, name: "You are not on the invite list" })).toBeInTheDocument();
    expect(screen.queryByRole("navigation", { name: "Main" })).not.toBeInTheDocument();
    expect(screen.queryByLabelText("Your message")).not.toBeInTheDocument();
  });

  it("lets an invited reader through to the page", async () => {
    await renderAt("/session/new");
    expect(await screen.findByRole("heading", { level: 1, name: "New deck" })).toBeInTheDocument();
    expect(screen.getByRole("navigation", { name: "Main" })).toBeInTheDocument();
    expect(screen.queryByRole("heading", { level: 1, name: "You are not on the invite list" })).not.toBeInTheDocument();
  });

  // The app stays closed when the API gives no answer (D-590). A reader
  // the API never cleared must not read the app, and a reader who is on
  // the list must not read a refusal that nobody gave.
  it("shows the server screen, and no app, when the call fails for another reason", async () => {
    listCollections.mockRejectedValue(new ConnectError("firestore down", Code.Unavailable));
    await renderAt("/session/new");
    expect(await screen.findByRole("heading", { level: 1, name: "Deck Tome could not reach the server" })).toBeInTheDocument();
    expect(screen.queryByRole("heading", { level: 1, name: "New deck" })).not.toBeInTheDocument();
    expect(screen.queryByRole("heading", { level: 1, name: "You are not on the invite list" })).not.toBeInTheDocument();
    expect(screen.queryByRole("navigation", { name: "Main" })).not.toBeInTheDocument();
    expect(screen.queryByLabelText("Your message")).not.toBeInTheDocument();
  });

  it("signs the refused reader out", async () => {
    listCollections.mockRejectedValue(refusal());
    await renderAt("/session/new");
    await screen.findByRole("heading", { level: 1, name: "You are not on the invite list" });
    const user = userEvent.setup();
    await user.click(screen.getByRole("button", { name: "Sign out" }));
    expect(vi.mocked(signOut)).toHaveBeenCalled();
  });

  // Nothing of the app draws while the gate waits. A reader off the list
  // must read no navigation at any moment, and not only after the answer.
  it("draws no navigation and no page while the answer is on the way", async () => {
    listCollections.mockReturnValue(new Promise(() => {}));
    await renderAt("/session/new");
    expect(await screen.findByText("Loading your session...")).toBeInTheDocument();
    expect(screen.queryByRole("navigation", { name: "Main" })).not.toBeInTheDocument();
    expect(screen.queryByLabelText("Your message")).not.toBeInTheDocument();
    expect(screen.queryByRole("heading", { level: 1, name: "New deck" })).not.toBeInTheDocument();
  });

  // The latch lives in a module, so it must not outlast the account that
  // set it. Without this the next reader on the browser reads a refusal
  // that the API never gave.
  it("forgets the latched answer on sign-out", async () => {
    setInviteState("refused");
    const { result } = renderHook(() => useInviteState());
    expect(result.current).toBe("refused");
    await act(async () => {
      await signOutAndClear(
        () => {},
        () => {},
      );
    });
    expect(result.current).toBe("unknown");
  });

  it("the refusal screen has no axe violations", async () => {
    listCollections.mockRejectedValue(refusal());
    const { container } = await renderAt("/session/new");
    await screen.findByRole("heading", { level: 1, name: "You are not on the invite list" });
    expect(await axe(container)).toHaveNoViolations();
  });
});

// unproved is what the API answers an invited caller whose email is not
// proved (D-903).
function unproved() {
  return new ConnectError("open the link in the email that Deck Tome sent", Code.PermissionDenied, {
    "Deck-Tome-Refusal": "email-unverified",
  });
}

// D-903: anyone can make an account with an invited address, so the API
// opens the app to a proved address alone. The reader proves it here.
describe("the proof of the email", () => {
  it("reads its own header, and the invite refusal stays apart", () => {
    expect(isUnverified(unproved())).toBe(true);
    expect(isNotInvited(unproved())).toBe(false);
    expect(isUnverified(refusal())).toBe(false);
    expect(isUnverified(new ConnectError("no", Code.PermissionDenied))).toBe(false);
  });

  it("asks for the proof, with no app and no navigation", async () => {
    listCollections.mockRejectedValue(unproved());
    await renderAt("/session/new");
    expect(await screen.findByRole("heading", { level: 1, name: "Confirm your email address" })).toBeInTheDocument();
    expect(screen.queryByRole("heading", { level: 1, name: "You are not on the invite list" })).not.toBeInTheDocument();
    expect(screen.queryByRole("navigation", { name: "Main" })).not.toBeInTheDocument();
    expect(screen.queryByLabelText("Your message")).not.toBeInTheDocument();
  });

  it("opens the app after the reader proves the email and presses Continue", async () => {
    const reader = { ...fakeUser, emailVerified: false, reload: vi.fn(), getIdToken: vi.fn(() => Promise.resolve("token-2")) };
    reader.reload.mockImplementation(() => {
      reader.emailVerified = true;
      return Promise.resolve();
    });
    state.user = reader as unknown as typeof fakeUser;
    listCollections.mockRejectedValueOnce(unproved());
    await renderAt("/session/new");
    await screen.findByRole("heading", { level: 1, name: "Confirm your email address" });
    const user = userEvent.setup();
    await user.click(screen.getByRole("button", { name: "Continue" }));
    expect(await screen.findByRole("heading", { level: 1, name: "New deck" })).toBeInTheDocument();
    // A new token carries the proof to the API with no sign-out.
    expect(reader.getIdToken).toHaveBeenCalledWith(true);
  });

  it("keeps the screen, and says why, while the email is not proved", async () => {
    state.user = { ...fakeUser, emailVerified: false, reload: () => Promise.resolve() } as unknown as typeof fakeUser;
    listCollections.mockRejectedValue(unproved());
    await renderAt("/session/new");
    await screen.findByRole("heading", { level: 1, name: "Confirm your email address" });
    const user = userEvent.setup();
    await user.click(screen.getByRole("button", { name: "Continue" }));
    expect(await screen.findByText(/not proved yet/)).toBeInTheDocument();
    await user.click(screen.getByRole("button", { name: "Send the link again" }));
    expect(vi.mocked(sendEmailVerification)).toHaveBeenCalled();
    expect(await screen.findByText(/A new link is on its way/)).toBeInTheDocument();
  });

  it("has no axe violations", async () => {
    listCollections.mockRejectedValue(unproved());
    const { container } = await renderAt("/session/new");
    await screen.findByRole("heading", { level: 1, name: "Confirm your email address" });
    expect(await axe(container)).toHaveNoViolations();
  });
});
