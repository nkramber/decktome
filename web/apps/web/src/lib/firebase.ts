import type { FirebaseApp } from "firebase/app";
import type { Auth } from "firebase/auth";

import { authCode } from "./errors";

// Firebase Auth loads after the shell paints (D-320). Nothing in this file
// imports the SDK at module scope, so the first-paint chunk carries no auth
// code. Every caller awaits loadAuth().
type AuthModule = typeof import("firebase/auth");

let pending: Promise<{ app: FirebaseApp; auth: Auth; mod: AuthModule }> | null = null;

// loadAuth downloads the SDK once and returns the same instance after that.
// The push module of PR-26 takes the app from it too.
export function loadAuth(): Promise<{ app: FirebaseApp; auth: Auth; mod: AuthModule }> {
  pending ??= start();
  return pending;
}

async function start() {
  const [{ initializeApp }, mod] = await Promise.all([import("firebase/app"), import("firebase/auth")]);

  // The project id must match the one the emulator and the API agree on (D-275).
  // A deployed build carries the real Firebase web configuration through the
  // VITE_FIREBASE_ variables (PR-22). The dev defaults serve the emulator.
  // Cloud Messaging needs the sender id, and a build with none shows no
  // push toggle (D-1005).
  const env = import.meta.env;
  const app = initializeApp({
    apiKey: env.VITE_FIREBASE_API_KEY || "demo-key",
    projectId: env.VITE_FIREBASE_PROJECT_ID || "mtg-local",
    authDomain: env.VITE_FIREBASE_AUTH_DOMAIN || "localhost",
    appId: env.VITE_FIREBASE_APP_ID || undefined,
    messagingSenderId: env.VITE_FIREBASE_MESSAGING_SENDER_ID || undefined,
  });

  // browserLocalPersistence keeps the session across a reload.
  const auth = mod.initializeAuth(app, { persistence: mod.browserLocalPersistence });

  // In dev the emulator host defaults to 127.0.0.1:9199. Set VITE_AUTH_EMULATOR_HOST
  // to point a build at another emulator. An empty value outside dev means real Firebase.
  const emulatorHost = import.meta.env.VITE_AUTH_EMULATOR_HOST ?? (import.meta.env.DEV ? "127.0.0.1:9199" : "");
  if (emulatorHost) {
    mod.connectAuthEmulator(auth, `http://${emulatorHost}`, { disableWarnings: true });
  }
  return { app, auth, mod };
}

// A token refresh that fails with one of these codes can never succeed
// again. The user signs out, and the route guard sends them to /sign-in.
const deadSessionCodes = new Set(["auth/user-token-expired", "auth/user-disabled"]);

// currentIdToken waits for the persisted session to load, then returns the
// token of the signed-in user. Empty when nobody is signed in.
export async function currentIdToken(): Promise<string> {
  const { auth, mod } = await loadAuth();
  await auth.authStateReady();
  const user = auth.currentUser;
  if (!user) return "";
  try {
    return await user.getIdToken();
  } catch (err) {
    if (deadSessionCodes.has(authCode(err))) {
      await mod.signOut(auth);
      return "";
    }
    throw err;
  }
}

// isAdmin reads the custom claim admin: true of the signed-in token
// (D-1076). It answers false on any failure, and the API checks the
// claim again on each admin call.
export async function isAdmin(): Promise<boolean> {
  try {
    const { auth, mod } = await loadAuth();
    await auth.authStateReady();
    const user = auth.currentUser;
    if (!user) return false;
    const result = await mod.getIdTokenResult(user);
    return result?.claims?.admin === true;
  } catch {
    return false;
  }
}

// signOutOfApp ends the session. The caller clears the local state.
export async function signOutOfApp(): Promise<void> {
  const { auth, mod } = await loadAuth();
  await mod.signOut(auth);
}

// signIn and createAccount serve the sign-in form only.
export async function signIn(email: string, password: string): Promise<void> {
  const { auth, mod } = await loadAuth();
  await mod.signInWithEmailAndPassword(auth, email, password);
}

// createAccount makes the account. The caller then sends the link that
// proves the email, because the invite gate trusts a proved email alone
// (D-903, D-1081).
export async function createAccount(email: string, password: string): Promise<void> {
  const { auth, mod } = await loadAuth();
  await mod.createUserWithEmailAndPassword(auth, email, password);
}

// sendFirebaseProof sends the email of Firebase that proves the address.
// It serves a server with no email of its own: local mode, or a deploy
// with no Resend secret (D-1081).
export async function sendFirebaseProof(): Promise<void> {
  const { auth, mod } = await loadAuth();
  if (auth.currentUser) await mod.sendEmailVerification(auth.currentUser);
}

// signInWithToken signs in with the custom token of a proof link
// (D-1082).
export async function signInWithToken(token: string): Promise<void> {
  const { auth, mod } = await loadAuth();
  await mod.signInWithCustomToken(auth, token);
}

// refreshProof reads the account again and takes a new token, so a proved
// email reaches the API at once. It answers whether the email is proved.
export async function refreshProof(): Promise<boolean> {
  const { auth } = await loadAuth();
  const user = auth.currentUser;
  if (!user) return false;
  await user.reload();
  await user.getIdToken(true);
  return user.emailVerified;
}

// resetPassword sends the link that sets a new password. An invited
// person whose address another person took first gets the account back
// this way (D-903).
export async function resetPassword(email: string): Promise<void> {
  const { auth, mod } = await loadAuth();
  await mod.sendPasswordResetEmail(auth, email);
}
