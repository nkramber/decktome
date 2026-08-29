import { useQueryClient } from "@tanstack/react-query";
import { lazy, Suspense } from "react";
import { Link, Outlet } from "react-router";

import { useAuth } from "../features/auth/auth-context";
import { errorMessage } from "../lib/errors";
import { signOutOfApp } from "../lib/firebase";
import { useAppStore } from "../lib/store";
import { AccountMenu } from "./components/account-menu";
import { BottomNav, SidebarNav } from "./components/app-nav";
import { ThemeMenu } from "./components/theme-menu";
import { useIsPhone } from "./components/use-viewport";

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

// The layout is the shell (D-311, D-317): a sidebar on a desktop and a
// bottom tab bar on a phone. A signed-out visitor gets the page alone.
export function Layout() {
  const { user } = useAuth();
  const queryClient = useQueryClient();
  const reset = useAppStore((s) => s.reset);
  const phone = useIsPhone();

  async function onSignOut() {
    try {
      await signOutAndClear(reset, () => queryClient.clear());
    } catch (err) {
      const { toast } = await import("../components/ui/toaster");
      toast.error("Sign-out failed", { description: errorMessage(err) });
    }
  }

  const email = user?.email ?? "";

  return (
    <div className="flex min-h-screen bg-background text-foreground">
      {user && !phone && (
        <aside className="sticky top-0 flex h-screen w-60 shrink-0 flex-col gap-4 border-r border-border bg-surface p-3">
          <Link to="/" className="px-3 py-2 text-lg font-semibold tracking-tight">
            MtG Deck Builder
          </Link>
          <SidebarNav />
          <div className="grow" />
          <div className="flex flex-col gap-1">
            <ThemeMenu />
            <AccountMenu email={email} onSignOut={() => void onSignOut()} />
          </div>
        </aside>
      )}
      <div className="flex min-w-0 grow flex-col">
        <main className={user && phone ? "grow pb-16" : "grow"}>
          <Outlet />
        </main>
        <Suspense fallback={null}>
          <HealthFooter className={user && phone ? "pb-16" : undefined} />
        </Suspense>
      </div>
      {user && phone && (
        <BottomNav>
          <AccountMenu email={email} onSignOut={() => void onSignOut()} withTheme trigger="tab" side="top" align="end" />
        </BottomNav>
      )}
      <Suspense fallback={null}>
        <Toaster />
      </Suspense>
    </div>
  );
}
