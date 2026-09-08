import { screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { createUserWithEmailAndPassword, signInWithEmailAndPassword } from "firebase/auth";
import { axe } from "jest-axe";
import { beforeEach, describe, expect, it, vi } from "vitest";

import { state } from "../../test-auth-state";
import { notAuthorized } from "./sign-in-page";
import { renderAt } from "../../test-utils";

vi.mock("firebase/app");
vi.mock("firebase/auth");
const checkInvite = vi.fn();
vi.mock("../../lib/api", () => ({
  healthClient: { check: () => Promise.resolve({ status: "ok", version: "test", cardSnapshot: "none" }) },
  inviteClient: { checkInvite: (...a: unknown[]) => checkInvite(...a) },
}));

const signIn = vi.mocked(signInWithEmailAndPassword);
const signUp = vi.mocked(createUserWithEmailAndPassword);

beforeEach(() => {
  state.user = null;
  signIn.mockReset();
  signUp.mockReset();
  checkInvite.mockReset();
  checkInvite.mockResolvedValue({ allowed: true });
});

describe("SignInPage", () => {
  it("renders the form with labeled inputs", async () => {
    await renderAt("/sign-in");
    expect(screen.getByRole("heading", { level: 1, name: "Sign in" })).toBeInTheDocument();
    expect(screen.getByLabelText("Email")).toBeInTheDocument();
    expect(screen.getByLabelText("Password")).toBeInTheDocument();
  });

  it("submits email and password to signInWithEmailAndPassword", async () => {
    signIn.mockResolvedValue({} as never);
    await renderAt("/sign-in");
    const user = userEvent.setup();
    await user.type(screen.getByLabelText("Email"), "nate@example.com");
    await user.type(screen.getByLabelText("Password"), "secret1");
    await user.click(screen.getByRole("button", { name: "Sign in" }));
    expect(signIn).toHaveBeenCalledWith(expect.anything(), "nate@example.com", "secret1");
    expect(signUp).not.toHaveBeenCalled();
  });

  it("toggles to create account and calls createUserWithEmailAndPassword", async () => {
    signUp.mockResolvedValue({} as never);
    await renderAt("/sign-in");
    const user = userEvent.setup();
    await user.click(screen.getByRole("button", { name: /Create account/ }));
    expect(screen.getByRole("heading", { level: 1, name: "Create account" })).toBeInTheDocument();
    await user.type(screen.getByLabelText("Email"), "new@example.com");
    await user.type(screen.getByLabelText("Password"), "secret1");
    await user.click(screen.getByRole("button", { name: "Create account" }));
    expect(signUp).toHaveBeenCalledWith(expect.anything(), "new@example.com", "secret1");
  });

  // D-592: an email off the invite list never becomes an account. The
  // form stays where it is, and it names the state in red. Before this
  // the browser made the account and the API refused every call after.
  it("refuses an email off the invite list and creates no account", async () => {
    checkInvite.mockResolvedValue({ allowed: false });
    await renderAt("/sign-in");
    const user = userEvent.setup();
    await user.click(screen.getByRole("button", { name: /Create account/ }));
    await user.type(screen.getByLabelText("Email"), "off@example.com");
    await user.type(screen.getByLabelText("Password"), "secret1");
    await user.click(screen.getByRole("button", { name: "Create account" }));

    expect(await screen.findByRole("alert")).toHaveTextContent(notAuthorized);
    expect(signUp).not.toHaveBeenCalled();
    // The form is still the form, and it still holds what was typed.
    expect(screen.getByRole("heading", { level: 1, name: "Create account" })).toBeInTheDocument();
    expect(screen.getByLabelText("Email")).toHaveValue("off@example.com");
  });

  it("asks the list before it makes an account, and never on a sign-in", async () => {
    signIn.mockResolvedValue({} as never);
    await renderAt("/sign-in");
    const user = userEvent.setup();
    await user.type(screen.getByLabelText("Email"), "nate@example.com");
    await user.type(screen.getByLabelText("Password"), "secret1");
    await user.click(screen.getByRole("button", { name: "Sign in" }));
    expect(checkInvite).not.toHaveBeenCalled();
  });

  // A check that can not answer must not stop a person on the list. The
  // API refuses the call after it in any case.
  it("makes the account when the check does not answer", async () => {
    checkInvite.mockRejectedValue(new Error("the API is down"));
    signUp.mockResolvedValue({} as never);
    await renderAt("/sign-in");
    const user = userEvent.setup();
    await user.click(screen.getByRole("button", { name: /Create account/ }));
    await user.type(screen.getByLabelText("Email"), "new@example.com");
    await user.type(screen.getByLabelText("Password"), "secret1");
    await user.click(screen.getByRole("button", { name: "Create account" }));
    expect(signUp).toHaveBeenCalled();
  });

  it("shows a readable message for a Firebase error", async () => {
    signIn.mockRejectedValue({ code: "auth/invalid-credential", message: "Firebase: Error" });
    await renderAt("/sign-in");
    const user = userEvent.setup();
    await user.type(screen.getByLabelText("Email"), "nate@example.com");
    await user.type(screen.getByLabelText("Password"), "wrong12");
    await user.click(screen.getByRole("button", { name: "Sign in" }));
    expect(await screen.findByRole("alert")).toHaveTextContent("The email or the password is wrong.");
  });

  it("has no axe violations", async () => {
    const { container } = await renderAt("/sign-in");
    await screen.findByText(/API: ok/);
    expect(await axe(container)).toHaveNoViolations();
  });
});
