import type { LucideIcon } from "lucide-react";
import type { ReactNode } from "react";

// One empty state: an icon, a sentence, and one action (D-311). A
// compact one sits inside a card beside another card, so it stays the
// height of its neighbor and not a tall frame of its own (D-458).
export function EmptyState({
  icon: Icon,
  title,
  description,
  action,
  compact = false,
}: {
  icon: LucideIcon;
  title: string;
  description?: string;
  action?: ReactNode;
  compact?: boolean;
}) {
  return (
    <div className={compact ? "flex flex-col items-center gap-3 px-2 py-6 text-center" : "flex flex-col items-center gap-3 rounded-card border border-dashed border-border px-6 py-20 text-center"}>
      <Icon className="size-9 text-primary" aria-hidden="true" />
      <p className="font-display text-base font-semibold">{title}</p>
      {description && <p className="max-w-prose text-[15px] text-muted-foreground">{description}</p>}
      {action}
    </div>
  );
}
