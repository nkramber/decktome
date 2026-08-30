import type { ComponentProps } from "react";

import { cn } from "../../lib/cn";

export function Input({ className, type, ...props }: ComponentProps<"input">) {
  return (
    <input
      type={type}
      data-slot="input"
      className={cn(
        "flex h-9 w-full min-w-0 rounded-lg border border-border bg-background px-3 py-1 text-base shadow-card transition-colors placeholder:text-muted-foreground hover:border-foreground/20 focus-visible:border-accent/60 disabled:cursor-not-allowed disabled:opacity-50 md:text-sm",
        "file:mr-3 file:h-7 file:rounded-md file:border-0 file:bg-muted file:px-3 file:text-sm file:font-medium file:text-foreground",
        "aria-invalid:border-danger",
        className,
      )}
      {...props}
    />
  );
}
