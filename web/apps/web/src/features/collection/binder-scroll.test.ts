import { afterEach, describe, expect, it } from "vitest";

import { moveBehavior, offsetIn, resetTop, showTopButton } from "./binder-scroll";

// The binder scrolls with the page, and these rules place it (D-806).
describe("the binder scroll", () => {
  afterEach(() => {
    delete (window as { matchMedia?: unknown }).matchMedia;
  });

  it("measures the binder from the top of the page content, not the screen", () => {
    const main = document.createElement("main");
    const el = document.createElement("section");
    main.getBoundingClientRect = () => ({ top: 100 }) as DOMRect;
    el.getBoundingClientRect = () => ({ top: 300 }) as DOMRect;
    Object.defineProperty(main, "scrollTop", { value: 50 });
    expect(offsetIn(el, main)).toBe(250);
  });

  it("returns a reader below the binder top to it, and leaves a reader above in place", () => {
    expect(resetTop(500, 200)).toBe(200);
    expect(resetTop(200, 200)).toBeUndefined();
    expect(resetTop(100, 200)).toBeUndefined();
  });

  it("shows the button one screen below the binder top, and not before (D-808)", () => {
    expect(showTopButton(1000, 200, 800)).toBe(false);
    expect(showTopButton(1001, 200, 800)).toBe(true);
  });

  it("jumps for a reader who asks for less motion, and scrolls smoothly for every other reader", () => {
    expect(moveBehavior()).toBe("smooth");
    for (const [reduced, want] of [
      [true, "auto"],
      [false, "smooth"],
    ] as const) {
      Object.defineProperty(window, "matchMedia", { configurable: true, value: (q: string) => ({ matches: reduced && q === "(prefers-reduced-motion: reduce)" }) });
      expect(moveBehavior()).toBe(want);
    }
  });
});
