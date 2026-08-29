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
export const signOut = vi.fn(() => Promise.resolve());
export const onAuthStateChanged = vi.fn((_auth: unknown, next: (u: User | null) => void, error?: (e: Error) => void) => {
  const listener = { next, error };
  state.listeners.add(listener);
  next(state.user);
  return () => {
    state.listeners.delete(listener);
  };
});
