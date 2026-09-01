// Vitest setup: DOM matchers and the axe accessibility matcher.
import "@testing-library/jest-dom/vitest";

import { cleanup } from "@testing-library/react";
import { toHaveNoViolations } from "jest-axe";
import { afterEach, expect } from "vitest";

expect.extend(toHaveNoViolations);
afterEach(cleanup);

// The Radix menus of the shell, and the binder grid, need browser APIs
// that jsdom omits. Without them a menu never opens and the grid throws,
// and the test reads as a product bug.
class ResizeObserverStub {
  observe() {}
  unobserve() {}
  disconnect() {}
}
globalThis.ResizeObserver ??= ResizeObserverStub as unknown as typeof ResizeObserver;
Element.prototype.hasPointerCapture ??= () => false;
Element.prototype.setPointerCapture ??= () => {};
Element.prototype.releasePointerCapture ??= () => {};
Element.prototype.scrollIntoView ??= () => {};
// The binder returns its grid to the top when the filter changes, and
// jsdom gives an element no scrollTo (PR-18).
Element.prototype.scrollTo ??= () => {};
