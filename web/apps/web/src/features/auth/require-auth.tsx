import { Navigate, Outlet } from "react-router";

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

// Every route except /sign-in needs a signed-in user (D-275). The guard
// keeps no page to return to, so a sign-in lands on the home page (D-1085).
export function RequireAuth() {
  const { user, ready } = useAuth();
  if (!ready) {
    return <LoadingSession />;
  }
  if (!user) {
    return <Navigate to="/sign-in" replace />;
  }
  return <Outlet />;
}

// The root path goes to Build when signed in, else to sign-in. Build is
// what the app is for, and the collection is a step on the way (D-334).
export function RootRedirect() {
  const { user, ready } = useAuth();
  if (!ready) return <LoadingSession />;
  // A signed-in reader lands on a new chat (D-334, D-436).
  return <Navigate to={user ? "/session/new" : "/sign-in"} replace />;
}
