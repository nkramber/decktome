// errorMessage turns a thrown value into text a user can read.
// Firebase Auth errors carry a code like auth/wrong-password. The code says
// more than the message, so the code wins when it is present.
export function errorMessage(err: unknown): string {
  if (typeof err === "object" && err !== null && "code" in err && typeof err.code === "string") {
    return authErrorText(err.code) ?? err.code;
  }
  return err instanceof Error ? err.message : String(err);
}

const authErrors: Record<string, string> = {
  "auth/invalid-credential": "The email or the password is wrong.",
  "auth/wrong-password": "The email or the password is wrong.",
  "auth/user-not-found": "No account has this email. Create one below.",
  "auth/email-already-in-use": "An account with this email exists. Sign in instead.",
  "auth/weak-password": "The password needs at least six characters.",
  "auth/invalid-email": "The email address is not valid.",
  "auth/network-request-failed": "The auth server did not answer. Is the emulator up?",
};

function authErrorText(code: string): string | undefined {
  return authErrors[code];
}
