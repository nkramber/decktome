import { initializeApp } from "firebase/app";
import { browserLocalPersistence, connectAuthEmulator, initializeAuth, signOut } from "firebase/auth";

import { authCode } from "./errors";

// Firebase Auth over the local emulator (D-275). The API checks the token
// through the emulator path (D-268). The project id must match the one the
// emulator and the API agree on: mtg-local.
const app = initializeApp({ apiKey: "demo-key", projectId: "mtg-local", authDomain: "localhost" });

// browserLocalPersistence keeps the session across a reload.
export const auth = initializeAuth(app, { persistence: browserLocalPersistence });

// In dev the emulator host defaults to 127.0.0.1:9199. Set VITE_AUTH_EMULATOR_HOST
// to point a build at another emulator. An empty value outside dev means real Firebase.
const emulatorHost = import.meta.env.VITE_AUTH_EMULATOR_HOST ?? (import.meta.env.DEV ? "127.0.0.1:9199" : "");
if (emulatorHost) {
  connectAuthEmulator(auth, `http://${emulatorHost}`, { disableWarnings: true });
}

// A token refresh that fails with one of these codes can never succeed
// again. The user signs out, and the route guard sends them to /sign-in.
const deadSessionCodes = new Set(["auth/user-token-expired", "auth/user-disabled"]);

// currentIdToken waits for the persisted session to load, then returns the
// token of the signed-in user. Empty when nobody is signed in.
export async function currentIdToken(): Promise<string> {
  await auth.authStateReady();
  const user = auth.currentUser;
  if (!user) return "";
  try {
    return await user.getIdToken();
  } catch (err) {
    if (deadSessionCodes.has(authCode(err))) {
      await signOut(auth);
      return "";
    }
    throw err;
  }
}
