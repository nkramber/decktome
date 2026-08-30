import { Navigate, Outlet, useLocation } from "react-router";

import { Skeleton } from "../../components/ui/skeleton";
import { useAuth } from "./auth-context";

// LoadingSession shows while firebase/auth loads and reads the persisted
// session. The SDK arrives after the first paint (D-320), so the shell
// paints this first.
export function LoadingSession() {
  return (
    <div className="mx-auto flex w-full max-w-5xl flex-col gap-4 p-4 md:p-6" aria-busy="true">
      <p className="sr-only" aria-live="polite">
        Loading your session...
      </p>
      <Skeleton className="h-8 w-56" />
      <Skeleton className="h-4 w-80" />
      <Skeleton className="h-40 w-full" />
    </div>
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
