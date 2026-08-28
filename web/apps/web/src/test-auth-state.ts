import type { User } from "firebase/auth";

// Shared state for the firebase/auth mock in __mocks__/firebase/auth.ts.
// A test sets `state.user` before it renders. No unit test talks to the emulator.
export const state: { user: User | null } = { user: null };

export const fakeUser = {
  uid: "u1",
  email: "nate@example.com",
  getIdToken: () => Promise.resolve("token-1"),
} as unknown as User;
