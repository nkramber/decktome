import { NavLink, useLocation } from "react-router";

import { cn } from "../../lib/cn";
import { useAppStore } from "../../lib/store";
import { isActive, navItems } from "./nav-items";

// The one navigation of the app (D-328). It sits in the header on every
// screen, and it carries the same three entries on a phone.
export function TopNav() {
  const { pathname } = useLocation();
  const clearCollection = useAppStore((s) => s.clearCollection);
  return (
    <nav aria-label="Main" className="flex items-center gap-1">
      {navItems.map((item) => {
        const active = isActive(pathname, item.match);
        return (
          <NavLink
            key={item.to}
            to={item.to}
            aria-current={active ? "page" : undefined}
            // Build always opens a chat over the whole card database
            // (D-349). The pool picker of the chat names a collection.
            onClick={item.to === "/session/new" ? () => clearCollection() : undefined}
            className={cn(
              "font-display flex items-center gap-1.5 rounded-card px-3 py-1.5 text-xs tracking-wide transition-all",
              active ? "bg-secondary text-foreground" : "text-muted-foreground hover:text-foreground",
            )}
          >
            <item.icon className="size-3.5" aria-hidden="true" />
            <span className="hidden sm:inline">{item.label}</span>
          </NavLink>
        );
      })}
    </nav>
  );
}
