import { useQueryClient } from "@tanstack/react-query";
import { useEffect } from "react";
import { Link, Outlet } from "react-router";

import { useAuth } from "../features/auth/auth-context";
import { useInviteState } from "../features/auth/invite-state";
import { signOutAndClear } from "../features/auth/sign-out";
import { InstallHint } from "../features/shell/install-hint";
import { errorMessage } from "../lib/errors";
import { useAppStore } from "../lib/store";
import { healthFooterChunk, scheduleWarm, toasterChunk } from "./chunks";
import { AccountMenu } from "./components/account-menu";
import { TopNav } from "./components/top-nav";

// The footer holds the one call that pulls the Connect client, and the
// toast host waits for a mutation. Both arrive after the first paint
// (D-320), and neither has a Suspense boundary of its own.
const HealthFooter = healthFooterChunk.Mount;
const Toaster = toasterChunk.Mount;

// The layout is the shell (D-328): one header over the whole width, the
// page under it, and the card-data line at the foot.
export function Layout() {
  const { user } = useAuth();
  const invite = useInviteState();
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
    <div className="flex h-screen flex-col overflow-hidden bg-background text-foreground print:h-auto print:overflow-visible">
      <header className="flex shrink-0 items-center justify-between gap-4 border-b border-border bg-muted px-4 py-3 md:px-6 print:hidden">
        <Link to="/" className="flex items-center gap-3 transition-opacity hover:opacity-80">
          <span aria-hidden="true" className="grid size-8 shrink-0 place-items-center rounded-card bg-accent text-accent-foreground">
            <TomeMark />
          </span>
          <span className="flex flex-col">
            <span className="font-display gold-shimmer text-[15px] font-semibold">Deck Tome</span>
            <span className="font-mono text-[10px] tracking-wide text-muted-foreground">AGENTIC DECK BUILDER</span>
          </span>
        </Link>

        {user && (
          <div className="flex items-center gap-1">
            {/* The navigation draws for a reader the API cleared, and
                for a page that asks nothing, such as a share link
                (F-59, D-590, D-315). It stays away while the gate waits,
                so a reader off the list reads no navigation at all. The
                account menu stays, because sign-out is the way out. */}
            {(invite === "invited" || invite === "unknown") && <TopNav />}
            <div className="ml-2 flex items-center gap-2 border-l border-border pl-3">
              <AccountMenu email={user.email ?? ""} onSignOut={() => void onSignOut()} side="bottom" align="end" />
            </div>
          </div>
        )}
      </header>

      {/* The header and the card-data line hold their place, and the
          page scrolls between them (D-364). A docked chat can then fill
          the frame and never run past it. */}
      <main className="grow overflow-y-auto print:overflow-visible">
        <Outlet />
      </main>

      <InstallHint />
      <HealthFooter />
      <Toaster />
    </div>
  );
}

// The mark of the app: a four-point star, cut rather than drawn.
// TomeMark is the brand mark: a closed tome with a spine groove, a gem on
// the cover, and the page block at its edge. The chat keeps SparkMark,
// because that mark stands for the agent and not for the product.
function TomeMark() {
  return (
    <svg width="16" height="16" viewBox="0 0 24 24" fill="currentColor" aria-hidden="true">
      <path
        fillRule="evenodd"
        clipRule="evenodd"
        d="M4 5.4A3.4 3.4 0 0 1 7.4 2h9.2A1.4 1.4 0 0 1 18 3.4v17.2A1.4 1.4 0 0 1 16.6 22H7.4A3.4 3.4 0 0 1 4 18.6V5.4Zm3.6-1.5h1.3v16.2H7.6V3.9Zm5.7 5.7 1.7 2.4-1.7 2.4-1.7-2.4 1.7-2.4Z"
      />
      <path d="M18.9 5.2c.9 0 1.6.7 1.6 1.6v10.4c0 .9-.7 1.6-1.6 1.6V5.2Z" opacity="0.45" />
    </svg>
  );
}
