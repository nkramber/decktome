import { signOut } from "firebase/auth";
import { Link, Outlet } from "react-router";

import { useAuth } from "../features/auth/auth-context";
import { auth } from "../lib/firebase";
import { HealthFooter } from "./components/health-footer";

// The layout wraps every page: a header with the nav and sign-out, the page, and the health footer.
export function Layout() {
  const { user } = useAuth();
  return (
    <div className="flex min-h-screen flex-col">
      <header className="flex items-center gap-4 border-b border-neutral-200 px-6 py-3">
        <Link to="/" className="text-lg font-semibold">
          MtG Deck Builder
        </Link>
        {user && (
          <nav aria-label="Main" className="flex gap-3">
            <Link to="/collection" className="underline">
              Collection
            </Link>
            <Link to="/decks" className="underline">
              Decks
            </Link>
          </nav>
        )}
        <span className="grow" />
        {user && (
          <>
            <span className="text-sm text-neutral-600">{user.email}</span>
            <button type="button" onClick={() => void signOut(auth)} className="rounded border px-2 py-1 text-sm">
              Sign out
            </button>
          </>
        )}
      </header>
      <main className="grow">
        <Outlet />
      </main>
      <HealthFooter />
    </div>
  );
}
