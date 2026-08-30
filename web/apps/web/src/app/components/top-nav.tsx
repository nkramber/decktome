import { NavLink, useLocation } from "react-router";

import { cn } from "../../lib/cn";
import { BuildMenu } from "./build-menu";
import { isActive, navItems } from "./nav-items";

// The one navigation of the app (D-328). It sits in the header on every
// screen, and it carries the same three entries on a phone.
export function TopNav() {
  const { pathname } = useLocation();
  return (
    <nav aria-label="Main" className="flex items-center gap-1">
      {navItems.map((item) => {
        const active = isActive(pathname, item.match);
        // Build asks which cards to use before it opens a chat, so it is
        // a menu and not a link.
        if (item.to === "/session/new") return <BuildMenu key={item.to} active={active} />;
        return (
          <NavLink
            key={item.to}
            to={item.to}
            aria-current={active ? "page" : undefined}
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
