import { Navigate, Outlet, useLocation } from "react-router";

import { useAuth } from "./auth-context";

// LoadingSession shows while firebase/auth reads the persisted session.
export function LoadingSession() {
  return (
    <p className="p-6 text-neutral-600" aria-live="polite">
      Loading your session...
    </p>
  );
}

// Every route except /sign-in needs a signed-in user (D-275). The page the
// user wanted goes along in the location state, so sign-in returns there.
export function RequireAuth() {
  const { user, ready } = useAuth();
  const location = useLocation();
  if (!ready) {
    return <LoadingSession />;
  }
  if (!user) {
    return <Navigate to="/sign-in" replace state={{ from: location }} />;
  }
  return <Outlet />;
}

// The root path goes to the collection when signed in, else to sign-in.
export function RootRedirect() {
  const { user, ready } = useAuth();
  if (!ready) return <LoadingSession />;
  return <Navigate to={user ? "/collection" : "/sign-in"} replace />;
}
