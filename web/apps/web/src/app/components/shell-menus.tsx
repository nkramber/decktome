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

// A menu of the shell loads on the first open, and it opens itself once
// it arrives. Radix measures the trigger to place the panel, and it
// reaches the trigger through the ref it passes as a prop. Every trigger
// here forwards the props it is given, or the panel lands outside the
// window with no way to click it.

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

// Build asks which cards the deck may use before it opens a chat. The
// pool rule follows the choice: a collection means owned-first, and any
// card means the whole database (D-37).
export function BuildMenuContent({ trigger }: { trigger: ReactNode }) {
  const navigate = useNavigate();
  const collectionId = useAppStore((s) => s.collectionId);
  const setCollection = useAppStore((s) => s.setCollection);
  const clearCollection = useAppStore((s) => s.clearCollection);

  const list = useQuery({
    queryKey: ["collections"],
    queryFn: () => collectionClient.listCollections({}),
  });
  const collections = list.data?.collections ?? [];

  function choose(value: string) {
    if (value === "") clearCollection();
    else setCollection(value);
    void navigate("/session/new");
  }

  return (
    <DropdownMenu defaultOpen>
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
