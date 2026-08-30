import { SparklesIcon } from "lucide-react";
import { lazy, Suspense, useState } from "react";

import { cn } from "../../lib/cn";

// The menu itself loads on the first open (D-320). The trigger below is
// the same element in both states, so nothing moves when it arrives.
const BuildMenuContent = lazy(async () => ({ default: (await import("./shell-menus")).BuildMenuContent }));

function BuildTrigger({ active, onClick }: { active: boolean; onClick?: () => void }) {
  return (
    <button
      type="button"
      aria-haspopup="menu"
      onClick={onClick}
      className={cn(
        "font-display flex items-center gap-1.5 rounded-card px-3 py-1.5 text-xs tracking-wide transition-all",
        active ? "bg-secondary text-foreground" : "text-muted-foreground hover:text-foreground",
      )}
    >
      <SparklesIcon className="size-3.5" aria-hidden="true" />
      <span className="hidden sm:inline">Build</span>
    </button>
  );
}

// Build opens a chat, and it asks which cards the deck may use first.
export function BuildMenu({ active }: { active: boolean }) {
  const [opened, setOpened] = useState(false);
  if (!opened) return <BuildTrigger active={active} onClick={() => setOpened(true)} />;
  return (
    <Suspense fallback={<BuildTrigger active={active} />}>
      <BuildMenuContent trigger={<BuildTrigger active={active} />} />
    </Suspense>
  );
}
