import { Slot } from "@radix-ui/react-slot";
import type { ComponentProps } from "react";

import { cn } from "../../lib/cn";

// asChild puts the styles on a child, so a card can be a section and a
// title can be the heading the page needs.
type WithChild = { asChild?: boolean };

export function Card({ className, asChild = false, ...props }: ComponentProps<"div"> & WithChild) {
  const Comp = asChild ? Slot : "div";
  return <Comp data-slot="card" className={cn("flex flex-col gap-4 rounded-card border border-border bg-surface py-4 text-surface-foreground", className)} {...props} />;
}

export function CardHeader({ className, ...props }: ComponentProps<"div">) {
  return <div data-slot="card-header" className={cn("flex flex-col gap-1 px-4", className)} {...props} />;
}

export function CardTitle({ className, asChild = false, ...props }: ComponentProps<"div"> & WithChild) {
  const Comp = asChild ? Slot : "div";
  return <Comp data-slot="card-title" className={cn("font-semibold", className)} {...props} />;
}

export function CardDescription({ className, ...props }: ComponentProps<"div">) {
  return <div data-slot="card-description" className={cn("text-sm text-muted-foreground", className)} {...props} />;
}

export function CardContent({ className, ...props }: ComponentProps<"div">) {
  return <div data-slot="card-content" className={cn("px-4", className)} {...props} />;
}

export function CardFooter({ className, ...props }: ComponentProps<"div">) {
  return <div data-slot="card-footer" className={cn("flex items-center px-4", className)} {...props} />;
}
