import { screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { createUserWithEmailAndPassword, signInWithEmailAndPassword } from "firebase/auth";
import { axe } from "jest-axe";
import { beforeEach, describe, expect, it, vi } from "vitest";

import { state } from "../../test-auth-state";
import { renderAt } from "../../test-utils";

vi.mock("firebase/app");
vi.mock("firebase/auth");
vi.mock("../../lib/api", () => ({
  healthClient: { check: () => Promise.resolve({ status: "ok", version: "test", cardSnapshot: "none" }) },
}));

const signIn = vi.mocked(signInWithEmailAndPassword);
const signUp = vi.mocked(createUserWithEmailAndPassword);

beforeEach(() => {
  state.user = null;
  signIn.mockReset();
  signUp.mockReset();
});

describe("SignInPage", () => {
  it("renders the form with labeled inputs", () => {
    renderAt("/sign-in");
    expect(screen.getByRole("heading", { level: 1, name: "Sign in" })).toBeInTheDocument();
    expect(screen.getByLabelText("Email")).toBeInTheDocument();
    expect(screen.getByLabelText("Password")).toBeInTheDocument();
  });

  it("submits email and password to signInWithEmailAndPassword", async () => {
    signIn.mockResolvedValue({} as never);
    renderAt("/sign-in");
    const user = userEvent.setup();
    await user.type(screen.getByLabelText("Email"), "nate@example.com");
    await user.type(screen.getByLabelText("Password"), "secret1");
    await user.click(screen.getByRole("button", { name: "Sign in" }));
    expect(signIn).toHaveBeenCalledWith(expect.anything(), "nate@example.com", "secret1");
    expect(signUp).not.toHaveBeenCalled();
  });

  it("toggles to create account and calls createUserWithEmailAndPassword", async () => {
    signUp.mockResolvedValue({} as never);
    renderAt("/sign-in");
    const user = userEvent.setup();
    await user.click(screen.getByRole("button", { name: /Create account/ }));
    expect(screen.getByRole("heading", { level: 1, name: "Create account" })).toBeInTheDocument();
    await user.type(screen.getByLabelText("Email"), "new@example.com");
    await user.type(screen.getByLabelText("Password"), "secret1");
    await user.click(screen.getByRole("button", { name: "Create account" }));
    expect(signUp).toHaveBeenCalledWith(expect.anything(), "new@example.com", "secret1");
  });

  it("shows a readable message for a Firebase error", async () => {
    signIn.mockRejectedValue({ code: "auth/invalid-credential", message: "Firebase: Error" });
    renderAt("/sign-in");
    const user = userEvent.setup();
    await user.type(screen.getByLabelText("Email"), "nate@example.com");
    await user.type(screen.getByLabelText("Password"), "wrong12");
    await user.click(screen.getByRole("button", { name: "Sign in" }));
    expect(await screen.findByRole("alert")).toHaveTextContent("The email or the password is wrong.");
  });

  it("has no axe violations", async () => {
    const { container } = renderAt("/sign-in");
    await screen.findByText(/API: ok/);
    expect(await axe(container)).toHaveNoViolations();
  });
});
