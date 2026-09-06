import type { Auth } from "firebase/auth";

import { authCode } from "./errors";

// Firebase Auth loads after the shell paints (D-320). Nothing in this file
// imports the SDK at module scope, so the first-paint chunk carries no auth
// code. Every caller awaits loadAuth().
type AuthModule = typeof import("firebase/auth");

let pending: Promise<{ auth: Auth; mod: AuthModule }> | null = null;

// loadAuth downloads the SDK once and returns the same instance after that.
export function loadAuth(): Promise<{ auth: Auth; mod: AuthModule }> {
  pending ??= start();
  return pending;
}

async function start() {
  const [{ initializeApp }, mod] = await Promise.all([import("firebase/app"), import("firebase/auth")]);

  // The project id must match the one the emulator and the API agree on (D-275).
  // A deployed build carries the real Firebase web configuration through the
  // four VITE_FIREBASE_ variables (PR-22). The dev defaults serve the emulator.
  const env = import.meta.env;
  const app = initializeApp({
    apiKey: env.VITE_FIREBASE_API_KEY || "demo-key",
    projectId: env.VITE_FIREBASE_PROJECT_ID || "mtg-local",
    authDomain: env.VITE_FIREBASE_AUTH_DOMAIN || "localhost",
    appId: env.VITE_FIREBASE_APP_ID || undefined,
  });

  // browserLocalPersistence keeps the session across a reload.
  const auth = mod.initializeAuth(app, { persistence: mod.browserLocalPersistence });

  // In dev the emulator host defaults to 127.0.0.1:9199. Set VITE_AUTH_EMULATOR_HOST
  // to point a build at another emulator. An empty value outside dev means real Firebase.
  const emulatorHost = import.meta.env.VITE_AUTH_EMULATOR_HOST ?? (import.meta.env.DEV ? "127.0.0.1:9199" : "");
  if (emulatorHost) {
    mod.connectAuthEmulator(auth, `http://${emulatorHost}`, { disableWarnings: true });
  }
  return { auth, mod };
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

export async function createAccount(email: string, password: string): Promise<void> {
  const { auth, mod } = await loadAuth();
  await mod.createUserWithEmailAndPassword(auth, email, password);
}
