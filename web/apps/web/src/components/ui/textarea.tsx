import type { ComponentProps } from "react";

import { cn } from "../../lib/cn";

export function Textarea({ className, ...props }: ComponentProps<"textarea">) {
  return (
    <textarea
      data-slot="textarea"
      className={cn(
        "flex field-sizing-content min-h-16 w-full rounded-lg border border-border bg-background px-3 py-2 text-base transition-colors placeholder:text-muted-foreground disabled:cursor-not-allowed disabled:bg-muted disabled:opacity-70 md:text-sm",
        className,
      )}
      {...props}
    />
  );
}
