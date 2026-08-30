import type { LucideIcon } from "lucide-react";
import type { ReactNode } from "react";

// One empty state: an icon, a sentence, and one action (D-311).
export function EmptyState({ icon: Icon, title, description, action }: { icon: LucideIcon; title: string; description?: string; action?: ReactNode }) {
  return (
    <div className="flex flex-col items-center gap-3 rounded-panel border border-dashed border-border bg-surface/50 px-6 py-14 text-center">
      <span aria-hidden="true" className="grid size-11 place-items-center rounded-full bg-muted text-muted-foreground">
        <Icon className="size-5" />
      </span>
      <p className="text-base font-medium">{title}</p>
      {description && <p className="max-w-prose text-sm text-muted-foreground">{description}</p>}
      {action}
    </div>
  );
}
