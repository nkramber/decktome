import { render, screen } from "@testing-library/react";
import { describe, expect, it } from "vitest";

import { Button, buttonVariants } from "./button";

// The phone gate of PR-25 asks for a touch target of 44 pixels. Every
// size keeps its desktop height and grows on a coarse pointer, so a
// phone gets the target and a desktop keeps the layout.
describe("the button sizes", () => {
  it("every size reaches 44 pixels on a coarse pointer", () => {
    for (const size of ["default", "sm", "lg", "icon"] as const) {
      const classes = buttonVariants({ size });
      // min-h-11 and size-11 are 2.75rem, which is 44 pixels.
      expect(classes, `size ${size}`).toMatch(/pointer-coarse:(min-h-11|size-11)/);
    }
  });

  it("keeps the desktop height, so no layout moves", () => {
    render(<Button size="sm">Copy</Button>);
    expect(screen.getByRole("button", { name: "Copy" })).toHaveClass("h-8");
  });
});
