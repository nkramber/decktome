import { describe, expect, it } from "vitest";

import { authCode, errorMessage, isInviteRefusal, signInErrorMessage } from "./errors";

describe("errorMessage", () => {
  it("returns the Error message and the string of anything else", () => {
    expect(errorMessage(new Error("boom"))).toBe("boom");
    expect(errorMessage("text")).toBe("text");
    expect(errorMessage(42)).toBe("42");
  });

  it("takes the auth branch only for an auth/ code", () => {
    expect(errorMessage({ code: "auth/invalid-credential", message: "x" })).toBe("The email or the password is wrong.");
    expect(errorMessage({ code: "auth/too-many-requests", message: "x" })).toBe("auth/too-many-requests");
    expect(errorMessage(Object.assign(new Error("[not_found] gone"), { code: "not_found" }))).toBe("[not_found] gone");
    expect(authCode({ code: "auth/x" })).toBe("auth/x");
    expect(authCode({ code: 5 })).toBe("");
  });

  it("signInErrorMessage adds the sign-in prefix to an unknown auth code only", () => {
    expect(signInErrorMessage({ code: "auth/too-many-requests" })).toBe("Sign-in failed (auth/too-many-requests).");
    expect(signInErrorMessage({ code: "auth/weak-password" })).toBe("The password needs at least six characters.");
    expect(signInErrorMessage(new Error("boom"))).toBe("boom");
  });

  // D-990: the SDK reports the answer of the blocking function as an
  // internal error, and the message carries the refusal.
  it("isInviteRefusal reads the refusal of the blocking function alone", () => {
    const refused = Object.assign(new Error(blockingMessage("not-invited")), { code: "auth/internal-error" });
    expect(isInviteRefusal(refused)).toBe(true);
    expect(isInviteRefusal(Object.assign(new Error(blockingMessage("down")), { code: "auth/internal-error" }))).toBe(false);
    expect(isInviteRefusal(Object.assign(new Error("not-invited"), { code: "auth/weak-password" }))).toBe(false);
    expect(isInviteRefusal({ code: "auth/internal-error", message: "not-invited" })).toBe(false);
  });
});

function blockingMessage(refusal: string): string {
  return `Firebase: ((HTTP request to https://api.example.run.app/auth/before-create returned HTTP error 403: {"error":{"status":"PERMISSION_DENIED","message":"${refusal}"}})) (auth/internal-error).`;
}
