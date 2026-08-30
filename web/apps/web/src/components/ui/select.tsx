import { ChevronDownIcon } from "lucide-react";
import type { ComponentProps } from "react";

import { cn } from "../../lib/cn";

// A styled native select. The root element sets color-scheme, so the
// browser paints the open list in the app's own theme (D-327). A native
// list needs no layer code, and it works on a phone by itself.
export function Select({ className, children, ...props }: ComponentProps<"select">) {
  return (
    <span className="relative inline-flex">
      <select
        data-slot="select"
        className={cn(
          "h-9 w-full appearance-none rounded-lg border border-border bg-surface py-1 pr-8 pl-3 text-sm shadow-card transition-colors hover:border-foreground/20 focus-visible:border-accent/60 disabled:cursor-not-allowed disabled:opacity-50",
          className,
        )}
        {...props}
      >
        {children}
      </select>
      <ChevronDownIcon aria-hidden="true" className="pointer-events-none absolute top-1/2 right-2.5 size-4 -translate-y-1/2 text-muted-foreground" />
    </span>
  );
}
