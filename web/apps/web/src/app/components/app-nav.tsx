import type { ReactNode } from "react";
import { NavLink, useLocation } from "react-router";

import { cn } from "../../lib/cn";
import { isActive, navItems } from "./nav-items";

// The sidebar navigation of a desktop (D-311).
export function SidebarNav() {
  const { pathname } = useLocation();
  return (
    <nav aria-label="Main" className="flex flex-col gap-1">
      {navItems.map((item) => {
        const active = isActive(pathname, item.match);
        return (
          <NavLink
            key={item.to}
            to={item.to}
            aria-current={active ? "page" : undefined}
            className={cn(
              "relative flex items-center gap-2.5 rounded-lg px-3 py-2 text-sm font-medium transition-colors",
              // The active entry carries a bar, not a color of text. An
              // accent that reads as a fill reads poorly as small text.
              active ? "bg-muted text-foreground before:absolute before:top-1.5 before:bottom-1.5 before:-left-px before:w-0.5 before:rounded-full before:bg-accent" : "text-muted-foreground hover:bg-muted/70 hover:text-foreground",
            )}
          >
            <item.icon className="size-4" aria-hidden="true" />
            {item.label}
          </NavLink>
        );
      })}
    </nav>
  );
}

// The bottom tab bar of a phone. It carries the same three entries.
export function BottomNav({ children }: { children?: ReactNode }) {
  const { pathname } = useLocation();
  return (
    <div className="fixed inset-x-0 bottom-0 z-40 border-t border-border bg-surface pb-[env(safe-area-inset-bottom)] shadow-overlay">
      <nav aria-label="Main" className="flex items-stretch">
        {navItems.map((item) => {
          const active = isActive(pathname, item.match);
          return (
            <NavLink
              key={item.to}
              to={item.to}
              aria-current={active ? "page" : undefined}
              className={cn(
                "relative flex flex-1 flex-col items-center gap-1 px-1 py-2.5 text-xs font-medium transition-colors",
                active ? "text-foreground before:absolute before:inset-x-5 before:top-0 before:h-0.5 before:rounded-full before:bg-accent" : "text-muted-foreground",
              )}
            >
              <item.icon className="size-5" aria-hidden="true" />
              {item.label}
            </NavLink>
          );
        })}
        {children}
      </nav>
    </div>
  );
}
