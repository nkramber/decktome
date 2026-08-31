import { UserIcon } from "lucide-react";
import { type ComponentProps, useState } from "react";

import { Button } from "../../components/ui/button";
import { accountMenuChunk } from "../chunks";

// The menu itself stays off the first paint (D-320). It mounts closed as
// soon as its chunk lands, so the first click opens it and waits for
// nothing. The trigger below is the same element in every state.

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
const Menu = accountMenuChunk.Mount;

export function AccountMenu({ email, onSignOut, align = "end", side = "bottom" }: AccountMenuProps) {
  const [open, setOpen] = useState(false);
  const waiting = (
    <AccountTrigger
      email={email}
      onClick={() => {
        accountMenuChunk.preload();
        setOpen(true);
      }}
    />
  );
  return <Menu fallback={waiting} open={open} onOpenChange={setOpen} trigger={<AccountTrigger email={email} />} email={email} onSignOut={onSignOut} align={align} side={side} />;
}
