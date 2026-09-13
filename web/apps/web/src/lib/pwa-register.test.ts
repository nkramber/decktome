import { describe, expect, it, vi } from "vitest";

import { startServiceWorker } from "./pwa-register";

const { registerSW } = vi.hoisted(() => ({ registerSW: vi.fn() }));
vi.mock("virtual:pwa-register", () => ({ registerSW }));

// F-122. The script the plugin injects registers the worker and never
// reloads the page, so the old shell ran for one load after a deploy.
// The plugin's register module reloads the page when a new worker takes
// over (D-692).
describe("startServiceWorker", () => {
  it("registers through the plugin's register module at once", () => {
    startServiceWorker();
    expect(registerSW).toHaveBeenCalledWith({ immediate: true });
  });
});
