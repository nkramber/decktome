import type { ComponentProps } from "react";

import { cn } from "../../lib/cn";

export function Input({ className, type, ...props }: ComponentProps<"input">) {
  return (
    <input
      type={type}
      data-slot="input"
      className={cn(
        "flex h-9 w-full min-w-0 rounded-card border border-border bg-secondary px-3 py-1 text-[15px] transition-colors outline-none placeholder:text-muted-foreground focus:border-primary disabled:cursor-not-allowed disabled:opacity-50",
        "file:mr-3 file:h-7 file:rounded-card file:border file:border-border file:bg-muted file:px-3 file:text-xs file:tracking-widest file:text-secondary-foreground file:uppercase",
        "aria-invalid:border-danger",
        className,
      )}
      {...props}
    />
  );
}
