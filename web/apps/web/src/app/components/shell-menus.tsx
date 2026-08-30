import { LogOutIcon } from "lucide-react";
import type { ReactNode } from "react";

import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuLabel, DropdownMenuRadioGroup, DropdownMenuRadioItem, DropdownMenuSeparator, DropdownMenuTrigger } from "../../components/ui/dropdown-menu";
import { type ThemeChoice, useThemeStore } from "../../lib/theme";

// The two shell menus, in one module the shell loads on the first open
// (D-320). Radix and its layer code stay off the first paint. Each menu
// opens by itself once it arrives, so the first click needs no second one.

const options: { value: ThemeChoice; label: string }[] = [
  { value: "light", label: "Light" },
  { value: "dark", label: "Dark" },
  { value: "system", label: "System" },
];

export function ThemeChoiceItems() {
  const choice = useThemeStore((s) => s.choice);
  const setChoice = useThemeStore((s) => s.setChoice);
  return (
    <DropdownMenuRadioGroup value={choice} onValueChange={(v) => setChoice(v as ThemeChoice)}>
      {options.map((o) => (
        <DropdownMenuRadioItem key={o.value} value={o.value}>
          {o.label}
        </DropdownMenuRadioItem>
      ))}
    </DropdownMenuRadioGroup>
  );
}

export function ThemeMenuContent({ trigger }: { trigger: ReactNode }) {
  return (
    <DropdownMenu defaultOpen>
      <DropdownMenuTrigger asChild>{trigger}</DropdownMenuTrigger>
      <DropdownMenuContent align="start" className="w-40">
        <DropdownMenuLabel>Theme</DropdownMenuLabel>
        <ThemeChoiceItems />
      </DropdownMenuContent>
    </DropdownMenu>
  );
}

export function AccountMenuContent({ trigger, email, onSignOut, withTheme = false, align = "start", side = "top" }: { trigger: ReactNode; email: string; onSignOut: () => void; withTheme?: boolean; align?: "start" | "center" | "end"; side?: "top" | "bottom" }) {
  return (
    <DropdownMenu defaultOpen>
      <DropdownMenuTrigger asChild>{trigger}</DropdownMenuTrigger>
      <DropdownMenuContent align={align} side={side} className="w-56">
        <DropdownMenuLabel className="truncate font-normal text-muted-foreground">{email}</DropdownMenuLabel>
        {withTheme && (
          <>
            <DropdownMenuSeparator />
            <DropdownMenuLabel>Theme</DropdownMenuLabel>
            <ThemeChoiceItems />
          </>
        )}
        <DropdownMenuSeparator />
        <DropdownMenuItem variant="destructive" onSelect={onSignOut}>
          <LogOutIcon aria-hidden="true" />
          Sign out
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
}
