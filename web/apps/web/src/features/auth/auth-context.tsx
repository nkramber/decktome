import { onAuthStateChanged, type User } from "firebase/auth";
import { createContext, type ReactNode, useContext, useEffect, useState } from "react";

import { auth } from "../../lib/firebase";

// The auth user lives in React context (wallabee tier 3). `ready` is false
// until firebase/auth has read the persisted session, so the route guard
// does not bounce a signed-in user to /sign-in on a reload.
export type AuthState = { user: User | null; ready: boolean };

const AuthContext = createContext<AuthState>({ user: null, ready: false });

export function AuthProvider({ children }: { children: ReactNode }) {
  const [state, setState] = useState<AuthState>({ user: null, ready: false });

  useEffect(() => {
    return onAuthStateChanged(auth, (user) => setState({ user, ready: true }));
  }, []);

  return <AuthContext.Provider value={state}>{children}</AuthContext.Provider>;
}

export function useAuth(): AuthState {
  return useContext(AuthContext);
}
