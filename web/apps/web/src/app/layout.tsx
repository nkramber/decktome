import { useQueryClient } from "@tanstack/react-query";
import { useEffect } from "react";
import { Link, Outlet } from "react-router";

import { useAuth } from "../features/auth/auth-context";
import { errorMessage } from "../lib/errors";
import { signOutOfApp } from "../lib/firebase";
import { useAppStore } from "../lib/store";
import { healthFooterChunk, scheduleWarm, toasterChunk } from "./chunks";
import { AccountMenu } from "./components/account-menu";
import { TopNav } from "./components/top-nav";

// The footer holds the one call that pulls the Connect client, and the
// toast host waits for a mutation. Both arrive after the first paint
// (D-320), and neither has a Suspense boundary of its own.
const HealthFooter = healthFooterChunk.Mount;
const Toaster = toasterChunk.Mount;

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


  // The deferred chunks arrive in the idle time after the first paint,
  // so the first click on a menu or a nav entry opens at once.
  useEffect(scheduleWarm, []);

  async function onSignOut() {
    try {
      await signOutAndClear(reset, () => queryClient.clear());
    } catch (err) {
      const { toast } = await import("../components/ui/toaster");
      toast.error("Sign-out failed", { description: errorMessage(err) });
    }
  }

  return (
    <div className="flex h-screen flex-col overflow-hidden bg-background text-foreground">
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

      {/* The header and the card-data line hold their place, and the
          page scrolls between them (D-364). A docked chat can then fill
          the frame and never run past it. */}
      <main className="grow overflow-y-auto">
        <Outlet />
      </main>

      <HealthFooter />
      <Toaster />
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
