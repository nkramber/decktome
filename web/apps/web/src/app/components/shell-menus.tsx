import { BellIcon, BellOffIcon, LogOutIcon } from "lucide-react";
import { type ReactNode, useEffect, useState } from "react";

import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuLabel, DropdownMenuSeparator, DropdownMenuTrigger } from "../../components/ui/dropdown-menu";
import { disablePush, enablePush, pushConfigured, type PushState, pushState } from "../../features/push/push";
import { notify } from "./notify";

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
        {pushConfigured() ? <PushMenuItem /> : null}
        <DropdownMenuItem variant="destructive" onSelect={onSignOut}>
          <LogOutIcon aria-hidden="true" />
          Sign out
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
}

// PushMenuItem turns on the push of a finished build for this browser
// (PR-26, D-1005). It reads the state each time the menu opens. A browser
// with no push shows no item, and an iPhone shows it in the app on the
// Home Screen alone.
function PushMenuItem() {
  const [state, setState] = useState<PushState | "reading">("reading");
  useEffect(() => {
    let live = true;
    pushState()
      .then((s) => live && setState(s))
      .catch(() => live && setState("off"));
    return () => {
      live = false;
    };
  }, []);
  if (state === "reading" || state === "unavailable") return null;
  if (state === "blocked") {
    return (
      <DropdownMenuItem disabled>
        <BellOffIcon aria-hidden="true" />
        Notifications are blocked
      </DropdownMenuItem>
    );
  }
  const on = state === "on";
  const toggle = async () => {
    try {
      const next = on ? await disablePush() : await enablePush();
      if (next === "on") void notify("success", "Notifications on", "You get a notification when a deck is ready.");
      else if (next === "blocked") void notify("error", "Notifications are blocked", "Allow notifications for this site in the browser settings.");
      else if (on) void notify("info", "Notifications off");
    } catch {
      void notify("error", on ? "Notifications could not be turned off" : "Notifications could not be turned on", "Please try again.");
    }
  };
  return (
    <DropdownMenuItem onSelect={() => void toggle()}>
      {on ? <BellOffIcon aria-hidden="true" /> : <BellIcon aria-hidden="true" />}
      {on ? "Stop deck notifications" : "Notify me when a deck is ready"}
    </DropdownMenuItem>
  );
}
