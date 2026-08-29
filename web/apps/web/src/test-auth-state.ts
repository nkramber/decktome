import type { User } from "firebase/auth";

// Shared state for the firebase/auth mock in __mocks__/firebase/auth.ts.
// A test sets `state.user` before it renders. No unit test talks to the emulator.
// emit and emitError drive the onAuthStateChanged listeners after the render.
type Listener = { next: (u: User | null) => void; error?: (e: Error) => void };

export const state: { user: User | null; listeners: Set<Listener> } = { user: null, listeners: new Set() };

export function emit(user: User | null) {
  state.user = user;
  for (const l of state.listeners) l.next(user);
}

export function emitError(err: Error) {
  for (const l of state.listeners) l.error?.(err);
}

export const fakeUser = {
  uid: "u1",
  email: "nate@example.com",
  getIdToken: () => Promise.resolve("token-1"),
} as unknown as User;
