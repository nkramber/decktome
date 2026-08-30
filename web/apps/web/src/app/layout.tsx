import { useQueryClient } from "@tanstack/react-query";
import { lazy, Suspense } from "react";
import { Link, Outlet } from "react-router";

import { useAuth } from "../features/auth/auth-context";
import { errorMessage } from "../lib/errors";
import { signOutOfApp } from "../lib/firebase";
import { useAppStore } from "../lib/store";
import { AccountMenu } from "./components/account-menu";
import { TopNav } from "./components/top-nav";

// The footer holds the one call that pulls the Connect client, so it loads
// after the shell paints (D-320).
const HealthFooter = lazy(async () => ({ default: (await import("./components/health-footer")).HealthFooter }));

// The toast host loads after the first paint (D-320). Nothing shows a
// toast before a mutation, and a mutation needs a click first.
const Toaster = lazy(async () => ({ default: (await import("../components/ui/toaster")).Toaster }));

// signOutAndClear clears the persisted ids and the query cache, so the
// next account on this browser starts with nothing of the last one.
export async function signOutAndClear(reset: () => void, clear: () => void) {
  await signOutOfApp();
  reset();
  clear();
}

// The layout is the shell (D-328): one header over the whole width, the
// page under it, and the card-data line at the foot.
export function Layout() {
  const { user } = useAuth();
  const queryClient = useQueryClient();
  const reset = useAppStore((s) => s.reset);

  async function onSignOut() {
    try {
      await signOutAndClear(reset, () => queryClient.clear());
    } catch (err) {
      const { toast } = await import("../components/ui/toaster");
      toast.error("Sign-out failed", { description: errorMessage(err) });
    }
  }

  return (
    <div className="flex min-h-screen flex-col bg-background text-foreground">
      <header className="flex shrink-0 items-center justify-between gap-4 border-b border-border bg-muted px-4 py-3 md:px-6">
        <Link to="/" className="flex items-center gap-3 transition-opacity hover:opacity-80">
          <span aria-hidden="true" className="grid size-8 shrink-0 place-items-center rounded-card bg-accent text-accent-foreground">
            <SparkMark />
          </span>
          <span className="flex flex-col">
            <span className="font-display gold-shimmer text-[15px] font-semibold">MtG Deck Builder</span>
            <span className="font-mono text-[10px] tracking-wide text-muted-foreground">DECK FORGE · AGENTIC</span>
          </span>
        </Link>

        {user && (
          <div className="flex items-center gap-1">
            <TopNav />
            <div className="ml-2 flex items-center gap-2 border-l border-border pl-3">
              <AccountMenu email={user.email ?? ""} onSignOut={() => void onSignOut()} side="bottom" align="end" />
            </div>
          </div>
        )}
      </header>

      <main className="grow">
        <Outlet />
      </main>

      <Suspense fallback={null}>
        <HealthFooter />
      </Suspense>
      <Suspense fallback={null}>
        <Toaster />
      </Suspense>
    </div>
  );
}

// The mark of the app: a four-point star, cut rather than drawn.
function SparkMark() {
  return (
    <svg width="14" height="14" viewBox="0 0 14 14" fill="currentColor" aria-hidden="true">
      <path d="M7 0l1.7 5.3L14 7l-5.3 1.7L7 14l-1.7-5.3L0 7l5.3-1.7z" />
    </svg>
  );
}
