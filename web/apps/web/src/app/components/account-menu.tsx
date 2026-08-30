import { UserIcon } from "lucide-react";
import { type ComponentProps, lazy, Suspense, useState } from "react";

import { Button } from "../../components/ui/button";

// The menu itself loads on the first open (D-320). The trigger below is
// the same element in both states, so nothing moves when it arrives.
const AccountMenuContent = lazy(async () => ({ default: (await import("./shell-menus")).AccountMenuContent }));

export type AccountMenuProps = { email: string; onSignOut: () => void; align?: "start" | "center" | "end"; side?: "top" | "bottom" };

// The trigger passes every prop it is given to its button, including the
// ref Radix measures to place the panel.
function AccountTrigger({ email, ...props }: { email: string } & ComponentProps<"button">) {
  return (
    <Button variant="ghost" size="sm" aria-haspopup="menu" aria-label="Account menu" {...props}>
      <UserIcon aria-hidden="true" />
      <span className="hidden max-w-32 truncate lg:inline">{email}</span>
    </Button>
  );
}

// The account menu holds the email and the sign-out (D-328).
export function AccountMenu({ email, onSignOut, align = "end", side = "bottom" }: AccountMenuProps) {
  const [opened, setOpened] = useState(false);
  if (!opened) return <AccountTrigger email={email} onClick={() => setOpened(true)} />;
  return (
    <Suspense fallback={<AccountTrigger email={email} />}>
      <AccountMenuContent trigger={<AccountTrigger email={email} />} email={email} onSignOut={onSignOut} align={align} side={side} />
    </Suspense>
  );
}
