import { UserIcon } from "lucide-react";
import { lazy, Suspense, useState } from "react";

import { Button } from "../../components/ui/button";

// The menu itself loads on the first open (D-320). The trigger below is
// the same element in both states, so nothing moves when it arrives.
const AccountMenuContent = lazy(async () => ({ default: (await import("./shell-menus")).AccountMenuContent }));

export type AccountMenuProps = { email: string; onSignOut: () => void; withTheme?: boolean; trigger?: "sidebar" | "tab"; align?: "start" | "center" | "end"; side?: "top" | "bottom" };

function AccountTrigger({ email, trigger = "sidebar", onClick }: { email: string; trigger?: "sidebar" | "tab"; onClick?: () => void }) {
  if (trigger === "tab") {
    return (
      <button type="button" aria-haspopup="menu" aria-label="Account menu" onClick={onClick} className="flex flex-1 flex-col items-center gap-1 px-1 py-2 text-xs font-medium text-muted-foreground">
        <UserIcon className="size-5" aria-hidden="true" />
        Account
      </button>
    );
  }
  return (
    <Button variant="ghost" className="w-full justify-start" aria-haspopup="menu" aria-label="Account menu" onClick={onClick}>
      <UserIcon aria-hidden="true" />
      <span className="min-w-0 truncate">{email}</span>
    </Button>
  );
}

// The account menu holds the email and the sign-out (D-311). On a phone it
// also holds the theme choice, because the bottom bar has no room for a
// second control.
export function AccountMenu({ email, onSignOut, withTheme = false, trigger = "sidebar", align = "start", side = "top" }: AccountMenuProps) {
  const [opened, setOpened] = useState(false);
  if (!opened) return <AccountTrigger email={email} trigger={trigger} onClick={() => setOpened(true)} />;
  return (
    <Suspense fallback={<AccountTrigger email={email} trigger={trigger} />}>
      <AccountMenuContent trigger={<AccountTrigger email={email} trigger={trigger} />} email={email} onSignOut={onSignOut} withTheme={withTheme} align={align} side={side} />
    </Suspense>
  );
}
