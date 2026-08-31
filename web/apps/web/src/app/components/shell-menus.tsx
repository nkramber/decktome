import { LogOutIcon } from "lucide-react";
import type { ReactNode } from "react";

import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuLabel, DropdownMenuSeparator, DropdownMenuTrigger } from "../../components/ui/dropdown-menu";

// A menu of the shell mounts closed once its chunk lands, and the caller
// owns the open state. Radix measures the trigger to place the panel, and
// it reaches the trigger through the ref it passes as a prop. Every
// trigger here forwards the props it is given, or the panel lands
// outside the window with no way to click it.

// The account menu, in one module the shell keeps off the first paint
// (D-320). Radix and its layer code arrive in the idle time after it.
// The theme choice left with the light theme (D-330).
export function AccountMenuContent({
  trigger,
  email,
  onSignOut,
  align = "start",
  side = "top",
  open,
  onOpenChange,
}: {
  trigger: ReactNode;
  email: string;
  onSignOut: () => void;
  align?: "start" | "center" | "end";
  side?: "top" | "bottom";
  open: boolean;
  onOpenChange: (open: boolean) => void;
}) {
  return (
    <DropdownMenu open={open} onOpenChange={onOpenChange}>
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
