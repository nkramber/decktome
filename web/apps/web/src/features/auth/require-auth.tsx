import { Navigate, Outlet } from "react-router";

import { useAuth } from "./auth-context";

// Every route except /sign-in needs a signed-in user (D-275).
export function RequireAuth() {
  const { user, ready } = useAuth();
  if (!ready) {
    return (
      <p className="p-6 text-neutral-600" aria-live="polite">
        Loading your session...
      </p>
    );
  }
  if (!user) {
    return <Navigate to="/sign-in" replace />;
  }
  return <Outlet />;
}

// The root path goes to the collection when signed in, else to sign-in.
export function RootRedirect() {
  const { user, ready } = useAuth();
  if (!ready) return null;
  return <Navigate to={user ? "/collection" : "/sign-in"} replace />;
}
