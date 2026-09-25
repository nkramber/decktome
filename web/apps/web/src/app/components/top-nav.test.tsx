import { render, screen } from "@testing-library/react";
import { MemoryRouter } from "react-router";
import { describe, expect, it } from "vitest";

import { navItems } from "./nav-items";
import { TopNav } from "./top-nav";

// REV-040: on a phone the label had the class `hidden`, which hides it
// from a screen reader too, and the icon is aria-hidden. So each link had
// no name.
describe("TopNav", () => {
  it("keeps a name on each link at every width", () => {
    render(
      <MemoryRouter>
        <TopNav />
      </MemoryRouter>,
    );
    for (const item of navItems) {
      const label = screen.getByText(item.label);
      expect(label.className).not.toMatch(/(^|\s)hidden(\s|$)/);
      expect(label.className).toContain("sr-only");
      expect(screen.getByRole("link", { name: item.label })).toBeInTheDocument();
    }
  });
});
