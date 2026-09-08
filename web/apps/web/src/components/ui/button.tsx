import { Slot } from "@radix-ui/react-slot";
import { cva, type VariantProps } from "class-variance-authority";
import type { ComponentProps } from "react";

import { cn } from "../../lib/cn";

// The button of the design system (D-311, D-321). asChild puts the styles
// on a child, so a router Link reads as a button without a nested control.
const buttonVariants = cva(
  "font-display inline-flex shrink-0 items-center justify-center gap-2 rounded-card text-xs tracking-widest uppercase transition-all duration-150 active:scale-95 disabled:pointer-events-none disabled:opacity-40 disabled:active:scale-100 [&_svg]:pointer-events-none [&_svg]:shrink-0 [&_svg:not([class*='size-'])]:size-4",
  {
    variants: {
      variant: {
        default: "bg-primary text-primary-foreground hover:brightness-110",
        outline: "border border-border text-secondary-foreground hover:border-primary hover:text-primary",
        ghost: "text-muted-foreground hover:text-foreground",
        destructive: "bg-danger text-danger-foreground hover:brightness-110",
        link: "text-primary tracking-normal normal-case underline underline-offset-4 hover:no-underline",
      },
      // A finger needs 44 pixels, and a mouse does not. Every size holds
      // its desktop height and grows on a coarse pointer, so the phone
      // gate of PR-25 holds without a wider layout anywhere else.
      size: {
        default: "h-9 px-4 py-2 pointer-coarse:min-h-11",
        sm: "h-8 px-3 pointer-coarse:min-h-11",
        lg: "h-10 px-6 pointer-coarse:min-h-11",
        icon: "size-9 pointer-coarse:size-11",
      },
    },
    defaultVariants: { variant: "default", size: "default" },
  },
);

export type ButtonProps = ComponentProps<"button"> & VariantProps<typeof buttonVariants> & { asChild?: boolean };

export function Button({ className, variant, size, asChild = false, type, ...props }: ButtonProps) {
  const Comp = asChild ? Slot : "button";
  return <Comp data-slot="button" type={asChild ? undefined : (type ?? "button")} className={cn(buttonVariants({ variant, size }), className)} {...props} />;
}

export { buttonVariants };
