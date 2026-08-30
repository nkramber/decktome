import { LogOutIcon } from "lucide-react";
import type { ReactNode } from "react";

import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuLabel, DropdownMenuSeparator, DropdownMenuTrigger } from "../../components/ui/dropdown-menu";

// The account menu, in one module the shell loads on the first open
// (D-320). Radix and its layer code stay off the first paint. The menu
// opens by itself once it arrives, so the first click needs no second
// one. The theme choice left with the light theme (D-330).
export function AccountMenuContent({
  trigger,
  email,
  onSignOut,
  align = "start",
  side = "top",
}: {
  trigger: ReactNode;
  email: string;
  onSignOut: () => void;
  align?: "start" | "center" | "end";
  side?: "top" | "bottom";
}) {
  return (
    <DropdownMenu defaultOpen>
      <DropdownMenuTrigger asChild>{trigger}</DropdownMenuTrigger>
      <DropdownMenuContent align={align} side={side} className="w-56">
        <DropdownMenuLabel className="truncate font-normal text-muted-foreground">{email}</DropdownMenuLabel>
        <DropdownMenuSeparator />
        <DropdownMenuItem variant="destructive" onSelect={onSignOut}>
          <LogOutIcon aria-hidden="true" />
          Sign out
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
}
