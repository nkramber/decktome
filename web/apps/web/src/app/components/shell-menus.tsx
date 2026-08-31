import { useQuery } from "@tanstack/react-query";
import { LogOutIcon, PlusIcon } from "lucide-react";
import type { ReactNode } from "react";
import { useNavigate } from "react-router";

import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuRadioGroup,
  DropdownMenuRadioItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "../../components/ui/dropdown-menu";
import { collectionClient } from "../../lib/api";
import { useAppStore } from "../../lib/store";

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

// Build asks which cards the deck may use before it opens a chat. The
// pool rule follows the choice: a collection means owned-first, and any
// card means the whole database (D-37).
export function BuildMenuContent({ trigger, open, onOpenChange }: { trigger: ReactNode; open: boolean; onOpenChange: (open: boolean) => void }) {
  const navigate = useNavigate();
  const collectionId = useAppStore((s) => s.collectionId);
  const setCollection = useAppStore((s) => s.setCollection);
  const clearCollection = useAppStore((s) => s.clearCollection);

  // The list arrives with the first open, never before it.
  const list = useQuery({
    queryKey: ["collections"],
    queryFn: () => collectionClient.listCollections({}),
    enabled: open,
  });
  const collections = list.data?.collections ?? [];

  function choose(value: string) {
    if (value === "") clearCollection();
    else setCollection(value);
    void navigate("/session/new");
  }

  return (
    <DropdownMenu open={open} onOpenChange={onOpenChange}>
      <DropdownMenuTrigger asChild>{trigger}</DropdownMenuTrigger>
      <DropdownMenuContent align="start" className="w-72">
        <DropdownMenuLabel>Build from</DropdownMenuLabel>
        <DropdownMenuRadioGroup value={collectionId} onValueChange={choose}>
          <DropdownMenuRadioItem value="">Any card</DropdownMenuRadioItem>
          {collections.map((c) => (
            <DropdownMenuRadioItem key={c.id} value={c.id}>
              <span className="truncate">{c.name}</span>
              <span className="ml-auto pl-2 font-mono text-[11px] text-muted-foreground">{c.cardCount}</span>
            </DropdownMenuRadioItem>
          ))}
        </DropdownMenuRadioGroup>
        {list.isSuccess && collections.length === 0 && (
          <>
            <DropdownMenuLabel className="font-normal text-muted-foreground">No collection yet. The agent builds from any card until you add one.</DropdownMenuLabel>
            <DropdownMenuSeparator />
            <DropdownMenuItem onSelect={() => void navigate("/collection")}>
              <PlusIcon aria-hidden="true" />
              Add a collection
            </DropdownMenuItem>
          </>
        )}
      </DropdownMenuContent>
    </DropdownMenu>
  );
}
