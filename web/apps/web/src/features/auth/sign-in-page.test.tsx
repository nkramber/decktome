import { Code, ConnectError } from "@connectrpc/connect";
import { screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { createUserWithEmailAndPassword, sendEmailVerification, sendPasswordResetEmail, signInWithEmailAndPassword } from "firebase/auth";
import { axe } from "jest-axe";
import { beforeEach, describe, expect, it, vi } from "vitest";

import { state } from "../../test-auth-state";
import { alreadyInvited, requestSent } from "./request-access";
import { notAuthorized, resetSent } from "./sign-in-page";
import { renderAt } from "../../test-utils";

vi.mock("firebase/app");
vi.mock("firebase/auth");
const checkInvite = vi.fn();
const requestAccess = vi.fn();
const sendLink = vi.fn();
vi.mock("../../lib/api", () => ({
  healthClient: { check: () => Promise.resolve({ status: "ok", version: "test", cardSnapshot: "none" }) },
  inviteClient: { checkInvite: (...a: unknown[]) => checkInvite(...a), requestAccess: (...a: unknown[]) => requestAccess(...a) },
  proofClient: { sendLink: (...a: unknown[]) => sendLink(...a) },
}));

const signIn = vi.mocked(signInWithEmailAndPassword);
const signUp = vi.mocked(createUserWithEmailAndPassword);

beforeEach(() => {
  state.user = null;
  signIn.mockReset();
  signUp.mockReset();
  checkInvite.mockReset();
  checkInvite.mockResolvedValue({ allowed: true });
  requestAccess.mockReset();
  requestAccess.mockResolvedValue({ alreadyInvited: false });
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

  async function createNewAccount() {
    // The new account is the signed-in user once Firebase makes it.
    signUp.mockImplementation(() => {
      state.user = { uid: "u9" } as never;
      return Promise.resolve({ user: { uid: "u9" } } as never);
    });
    vi.mocked(sendEmailVerification).mockClear();
    await renderAt("/sign-in");
    const user = userEvent.setup();
    await user.click(screen.getByRole("button", { name: "Create account" }));
    await user.type(screen.getByLabelText("Email"), "new@example.com");
    await user.type(screen.getByLabelText("Password"), "secret1");
    await user.click(screen.getByRole("button", { name: "Create account" }));
  }

  // D-903 and D-1081: the API trusts a proved email alone, so a new
  // account gets the link that proves it, from the API.
  it("asks the API for the link that proves the email after it makes the account", async () => {
    sendLink.mockReset();
    sendLink.mockResolvedValue({ alreadyProved: false });
    await createNewAccount();
    await vi.waitFor(() => expect(sendLink).toHaveBeenCalled());
    expect(vi.mocked(sendEmailVerification)).not.toHaveBeenCalled();
  });

  // D-1081: a server with no email of its own answers Unimplemented, and
  // the email of Firebase goes out in its place. A failed send does the
  // same.
  it.each([Code.Unimplemented, Code.Internal])("sends the email of Firebase when the API answers %s", async (code) => {
    sendLink.mockReset();
    sendLink.mockRejectedValue(new ConnectError("no email", code));
    await createNewAccount();
    await vi.waitFor(() => expect(vi.mocked(sendEmailVerification)).toHaveBeenCalled());
  });

  // D-903: an invited person whose address another person took first sets
  // a new password through the email, so the account comes back to them.
  it("sends a reset link for the typed email, and names no account", async () => {
    vi.mocked(sendPasswordResetEmail).mockClear();
    await renderAt("/sign-in");
    const user = userEvent.setup();
    await user.click(screen.getByRole("button", { name: "Forgot your password?" }));
    expect(await screen.findByRole("alert")).toHaveTextContent("Enter your email above");
    expect(vi.mocked(sendPasswordResetEmail)).not.toHaveBeenCalled();
    await user.type(screen.getByLabelText("Email"), "ann@example.com");
    await user.click(screen.getByRole("button", { name: "Forgot your password?" }));
    expect(vi.mocked(sendPasswordResetEmail)).toHaveBeenCalledWith(expect.anything(), "ann@example.com");
    expect(await screen.findByRole("status")).toHaveTextContent(resetSent);
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

  // D-990: when the check fails open, the blocking function still
  // refuses the account, and the form names the same state.
  it("shows the invite sentence when the blocking function refuses", async () => {
    checkInvite.mockRejectedValue(new Error("unavailable"));
    const refused = Object.assign(
      new Error('Firebase: ((HTTP request returned HTTP error 403: {"error":{"message":"not-invited"}})) (auth/internal-error).'),
      { code: "auth/internal-error" },
    );
    signUp.mockRejectedValueOnce(refused);
    await renderAt("/sign-in");
    const user = userEvent.setup();
    await user.click(screen.getByRole("button", { name: /Create account/ }));
    await user.type(screen.getByLabelText("Email"), "off@example.com");
    await user.type(screen.getByLabelText("Password"), "secret1");
    await user.click(screen.getByRole("button", { name: "Create account" }));

    expect(await screen.findByRole("alert")).toHaveTextContent(notAuthorized);
    expect(signUp).toHaveBeenCalled();
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

  // D-1074: the sign-in page holds a link to the request form, and the
  // request takes the email and the optional note.
  it("sends a request for beta access with the email and the note", async () => {
    await renderAt("/sign-in");
    const user = userEvent.setup();
    await user.click(screen.getByRole("button", { name: /Request beta access/ }));
    expect(screen.getByRole("heading", { level: 1, name: "Request beta access" })).toBeInTheDocument();
    await user.type(screen.getByLabelText("Email"), "new@example.com");
    await user.type(screen.getByLabelText(/What do you want to build/), "Elves");
    await user.click(screen.getByRole("button", { name: "Request beta access" }));
    expect(requestAccess).toHaveBeenCalledWith({ email: "new@example.com", note: "Elves" });
    expect(await screen.findByRole("status")).toHaveTextContent(requestSent);
    // D-1079: after a request, the link reads "Return to login page".
    expect(screen.queryByRole("button", { name: "I have an account. Sign in." })).not.toBeInTheDocument();
    await user.click(screen.getByRole("button", { name: "Return to login page" }));
    expect(screen.getByRole("heading", { level: 1, name: "Sign in" })).toBeInTheDocument();
  });

  // D-1079: the links under the form read short names, in one column
  // with one gap.
  it("shows the short links in one column", async () => {
    await renderAt("/sign-in");
    const create = screen.getByRole("button", { name: "Create account" });
    const request = screen.getByRole("button", { name: "Request beta access" });
    const forgot = screen.getByRole("button", { name: "Forgot your password?" });
    expect(create.parentElement).toBe(request.parentElement);
    expect(forgot.parentElement).toBe(create.parentElement);
    expect(create.parentElement).toHaveClass("flex-col", "gap-1");
    expect(screen.queryByText(/New here\?|No invite yet\?/)).not.toBeInTheDocument();
  });

  // D-1074: the red refusal line offers the request, with the email of
  // the form.
  it("offers the request under the refusal, with the typed email", async () => {
    checkInvite.mockResolvedValue({ allowed: false });
    await renderAt("/sign-in");
    const user = userEvent.setup();
    await user.click(screen.getByRole("button", { name: /Create account/ }));
    await user.type(screen.getByLabelText("Email"), "off@example.com");
    await user.type(screen.getByLabelText("Password"), "secret1");
    await user.click(screen.getByRole("button", { name: "Create account" }));
    await screen.findByText(notAuthorized);
    await user.click(screen.getByRole("button", { name: "Request beta access" }));
    expect(screen.getByLabelText("Email")).toHaveValue("off@example.com");
    await user.click(screen.getByRole("button", { name: "Request beta access" }));
    expect(requestAccess).toHaveBeenCalledWith({ email: "off@example.com", note: "" });
  });

  it("opens the create-account form for an email that is already invited", async () => {
    requestAccess.mockResolvedValue({ alreadyInvited: true });
    await renderAt("/sign-in");
    const user = userEvent.setup();
    await user.click(screen.getByRole("button", { name: /Request beta access/ }));
    await user.type(screen.getByLabelText("Email"), "ann@example.com");
    await user.click(screen.getByRole("button", { name: "Request beta access" }));
    expect(await screen.findByRole("heading", { level: 1, name: "Create account" })).toBeInTheDocument();
    expect(screen.getByRole("status")).toHaveTextContent(alreadyInvited);
    expect(screen.getByLabelText("Email")).toHaveValue("ann@example.com");
  });

  it("has no axe violations on the request form", async () => {
    const { container } = await renderAt("/sign-in");
    const user = userEvent.setup();
    await user.click(await screen.findByRole("button", { name: /Request beta access/ }));
    expect(await axe(container)).toHaveNoViolations();
  });

  it("has no axe violations", async () => {
    const { container } = await renderAt("/sign-in");
    // The form is the settled page. It was the health bar until D-626,
    // and the bar is gone.
    await screen.findByRole("button", { name: "Sign in" });
    expect(await axe(container)).toHaveNoViolations();
  });
});
