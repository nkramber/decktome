import { NavLink, useLocation } from "react-router";

import { cn } from "../../lib/cn";
import { isActive, navItems } from "./nav-items";

// The one navigation of the app (D-328). It sits in the header on every
// screen, and it carries the same three entries on a phone.
export function TopNav() {
  const { pathname } = useLocation();
  return (
    <nav aria-label="Main" className="flex items-center gap-1">
      {navItems.map((item) => {
        const active = isActive(pathname, item.match);
        return (
          <NavLink
            key={item.to}
            to={item.to}
            aria-current={active ? "page" : undefined}
            // Build keeps the pool the reader chose (D-580, F-63). The
            // picker of the chat shows it, so nothing hides, and a click
            // of Build never drops a collection without a word.
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
