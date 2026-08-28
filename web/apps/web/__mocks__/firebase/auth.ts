import type { User } from "firebase/auth";
import { vi } from "vitest";

import { state } from "../../src/test-auth-state";

// Manual mock. A test activates it with vi.mock("firebase/auth").
export const browserLocalPersistence = {};
export const connectAuthEmulator = vi.fn();
export const initializeAuth = () => ({
  authStateReady: () => Promise.resolve(),
  get currentUser() {
    return state.user;
  },
});
export const signInWithEmailAndPassword = vi.fn();
export const createUserWithEmailAndPassword = vi.fn();
export const signOut = vi.fn();
export const onAuthStateChanged = vi.fn((_auth: unknown, cb: (u: User | null) => void) => {
  cb(state.user);
  return () => {};
});
