import type { LucideIcon } from "lucide-react";
import type { ReactNode } from "react";

// One empty state: an icon, a sentence, and one action (D-311).
export function EmptyState({ icon: Icon, title, description, action }: { icon: LucideIcon; title: string; description?: string; action?: ReactNode }) {
  return (
    <div className="flex flex-col items-center gap-3 rounded-card border border-dashed border-border px-6 py-20 text-center">
      <Icon className="size-9 text-primary" aria-hidden="true" />
      <p className="font-display text-base font-semibold">{title}</p>
      {description && <p className="max-w-prose text-[15px] text-muted-foreground">{description}</p>}
      {action}
    </div>
  );
}
