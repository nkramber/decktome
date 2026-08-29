// errorMessage turns a thrown value into text a user can read.
// A Firebase Auth error carries a code like auth/wrong-password. The
// code maps to a sentence when the app knows it, else the code shows.
export function errorMessage(err: unknown): string {
  const code = authCode(err);
  if (code) {
    return authErrorText(code) ?? code;
  }
  return err instanceof Error ? err.message : String(err);
}

// signInErrorMessage is errorMessage with the sign-in prefix on an
// unknown auth code. Only the sign-in page uses it.
export function signInErrorMessage(err: unknown): string {
  const code = authCode(err);
  if (code) {
    return authErrorText(code) ?? `Sign-in failed (${code}).`;
  }
  return errorMessage(err);
}

// authCode returns the Firebase Auth code of an error, or an empty string.
export function authCode(err: unknown): string {
  if (typeof err === "object" && err !== null && "code" in err && typeof err.code === "string" && err.code.startsWith("auth/")) {
    return err.code;
  }
  return "";
}

const authErrors: Record<string, string> = {
  "auth/invalid-credential": "The email or the password is wrong.",
  "auth/wrong-password": "The email or the password is wrong.",
  "auth/user-not-found": "No account has this email. Create one below.",
  "auth/email-already-in-use": "An account with this email exists. Sign in instead.",
  "auth/weak-password": "The password needs at least six characters.",
  "auth/invalid-email": "The email address is not valid.",
  "auth/network-request-failed": "The auth server did not answer. Is the emulator up?",
  "auth/user-token-expired": "Your session expired. Sign in again.",
  "auth/user-disabled": "This account is disabled.",
};

function authErrorText(code: string): string | undefined {
  return authErrors[code];
}
