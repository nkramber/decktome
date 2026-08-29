import type { User } from "firebase/auth";
import { createContext, type ReactNode, useContext, useEffect, useState } from "react";

import { errorMessage } from "../../lib/errors";
import { loadAuth } from "../../lib/firebase";

// The auth user lives in React context (wallabee tier 3). `ready` is false
// until firebase/auth has read the persisted session, so the route guard
// does not bounce a signed-in user to /sign-in on a reload. `error` holds
// the message when the listener fails, so the app does not wait forever.
export type AuthState = { user: User | null; ready: boolean; error: string };

const AuthContext = createContext<AuthState>({ user: null, ready: false, error: "" });

export function AuthProvider({ children }: { children: ReactNode }) {
  const [state, setState] = useState<AuthState>({ user: null, ready: false, error: "" });

  // The SDK loads after the first paint (D-320), so the listener starts in
  // a second turn. `live` drops a result that arrives after the unmount.
  useEffect(() => {
    let live = true;
    let stop = () => {};
    loadAuth()
      .then(({ auth, mod }) => {
        if (!live) return;
        stop = mod.onAuthStateChanged(
          auth,
          (user) => setState({ user, ready: true, error: "" }),
          (err) => setState({ user: null, ready: true, error: errorMessage(err) }),
        );
      })
      .catch((err: unknown) => {
        if (live) setState({ user: null, ready: true, error: errorMessage(err) });
      });
    return () => {
      live = false;
      stop();
    };
  }, []);

  return <AuthContext.Provider value={state}>{children}</AuthContext.Provider>;
}

export function useAuth(): AuthState {
  return useContext(AuthContext);
}
