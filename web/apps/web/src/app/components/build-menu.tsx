import { SparklesIcon } from "lucide-react";
import { type ComponentProps, useState } from "react";

import { cn } from "../../lib/cn";
import { buildMenuChunk } from "../chunks";

// The menu itself stays off the first paint (D-320). It mounts closed as
// soon as its chunk lands, so the first click opens it and waits for
// nothing. The trigger below is the same element in every state, so
// nothing moves when the menu arrives.

// The trigger passes every prop it is given to its button, including the
// ref. Radix measures the trigger through that ref to place the panel,
// and a trigger that drops it leaves the panel outside the window.
function BuildTrigger({ active, className, ...props }: { active: boolean } & ComponentProps<"button">) {
  return (
    <button
      type="button"
      aria-haspopup="menu"
      {...props}
      className={cn(
        "font-display flex items-center gap-1.5 rounded-card px-3 py-1.5 text-xs tracking-wide transition-all",
        active ? "bg-secondary text-foreground" : "text-muted-foreground hover:text-foreground",
        className,
      )}
    >
      <SparklesIcon className="size-3.5" aria-hidden="true" />
      <span className="hidden sm:inline">Build</span>
    </button>
  );
}

// Build opens a chat, and it asks which cards the deck may use first.
const Menu = buildMenuChunk.Mount;

export function BuildMenu({ active }: { active: boolean }) {
  const [open, setOpen] = useState(false);
  // A click before the chunk lands asks for it, and the menu opens the
  // moment it arrives.
  const waiting = (
    <BuildTrigger
      active={active}
      onClick={() => {
        buildMenuChunk.preload();
        setOpen(true);
      }}
    />
  );
  return <Menu fallback={waiting} open={open} onOpenChange={setOpen} trigger={<BuildTrigger active={active} />} />;
}
