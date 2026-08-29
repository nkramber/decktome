import { useQueryClient } from "@tanstack/react-query";
import { signOut } from "firebase/auth";
import { useState } from "react";
import { Link, Outlet } from "react-router";

import { useAuth } from "../features/auth/auth-context";
import { errorMessage } from "../lib/errors";
import { auth } from "../lib/firebase";
import { useAppStore } from "../lib/store";
import { HealthFooter } from "./components/health-footer";

// signOutAndClear clears the persisted ids and the query cache, so the
// next account on this browser starts with nothing of the last one.
export async function signOutAndClear(reset: () => void, clear: () => void) {
  await signOut(auth);
  reset();
  clear();
}

// The layout wraps every page: a header with the nav and sign-out, the page, and the health footer.
export function Layout() {
  const { user } = useAuth();
  const queryClient = useQueryClient();
  const reset = useAppStore((s) => s.reset);
  const [signOutError, setSignOutError] = useState("");

  async function onSignOut() {
    try {
      await signOutAndClear(reset, () => queryClient.clear());
    } catch (err) {
      setSignOutError(errorMessage(err));
    }
  }

  return (
    <div className="flex min-h-screen flex-col bg-white text-neutral-900">
      <header className="flex flex-wrap items-center gap-x-4 gap-y-2 border-b border-neutral-200 px-6 py-3">
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
            <span className="min-w-0 truncate text-sm text-neutral-600">{user.email}</span>
            <button type="button" onClick={() => void onSignOut()} className="rounded border px-2 py-1 text-sm">
              Sign out
            </button>
          </>
        )}
        {signOutError && (
          <p role="alert" className="w-full text-sm text-red-700">
            Sign-out failed: {signOutError}
          </p>
        )}
      </header>
      <main className="grow">
        <Outlet />
      </main>
      <HealthFooter />
    </div>
  );
}
