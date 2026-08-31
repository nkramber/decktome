import type { ReactNode } from "react";

import { cn } from "../../lib/cn";

// One header for every page: the title, one line under it, and the
// actions on the right (D-311).
export function PageHeader({ title, description, actions, className }: { title: string; description?: string; actions?: ReactNode; className?: string }) {
  return (
    <div className={cn("flex flex-wrap items-end justify-between gap-4", className)}>
      <div className="min-w-0">
        <h1 className="font-display text-2xl font-semibold text-balance">{title}</h1>
        {description && <p className="mt-1 max-w-measure text-[15px] text-muted-foreground">{description}</p>}
      </div>
      {actions && <div className="flex shrink-0 flex-wrap items-center gap-2">{actions}</div>}
    </div>
  );
}
