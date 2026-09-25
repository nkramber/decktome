import { afterEach, describe, expect, it, vi } from "vitest";

import { startServiceWorker, swUpdateIntervalMs } from "./pwa-register";

const { registerSW } = vi.hoisted(() => ({ registerSW: vi.fn() }));
vi.mock("virtual:pwa-register", () => ({ registerSW }));

type RegisterOptions = { immediate?: boolean; onRegisteredSW?: (url: string, r?: ServiceWorkerRegistration) => void };

// registeredWith starts the worker and hands a fake registration to the
// plugin's callback. It returns the spy of update().
function registeredWith(): ReturnType<typeof vi.fn> {
  registerSW.mockClear();
  startServiceWorker();
  const options = registerSW.mock.calls[0][0] as RegisterOptions;
  const update = vi.fn(() => Promise.resolve());
  options.onRegisteredSW?.("/sw.js", { update } as unknown as ServiceWorkerRegistration);
  return update;
}

function setVisibility(value: DocumentVisibilityState) {
  Object.defineProperty(document, "visibilityState", { configurable: true, value });
  document.dispatchEvent(new Event("visibilitychange"));
}

afterEach(() => {
  vi.useRealTimers();
  setVisibility("visible");
});

// F-122. The script the plugin injects registers the worker and never
// reloads the page, so the old shell ran for one load after a deploy.
// The plugin's register module reloads the page when a new worker takes
// over (D-692).
describe("startServiceWorker", () => {
  it("registers through the plugin's register module at once", () => {
    registerSW.mockClear();
    startServiceWorker();
    expect(registerSW).toHaveBeenCalledWith(expect.objectContaining({ immediate: true }));
  });

  // REV-048. The browser checks for a new worker on a navigation alone,
  // so an open tab or an installed app kept the old shell (D-621).
  it("checks for a new worker on each interval", () => {
    vi.useFakeTimers();
    const update = registeredWith();
    expect(update).not.toHaveBeenCalled();
    vi.advanceTimersByTime(swUpdateIntervalMs);
    expect(update).toHaveBeenCalledTimes(1);
    vi.advanceTimersByTime(swUpdateIntervalMs);
    expect(update).toHaveBeenCalledTimes(2);
  });

  it("checks for a new worker when the page becomes visible", () => {
    const update = registeredWith();
    setVisibility("hidden");
    expect(update).not.toHaveBeenCalled();
    setVisibility("visible");
    expect(update).toHaveBeenCalledTimes(1);
  });

  it("does nothing when the browser gives no registration", () => {
    registerSW.mockClear();
    startServiceWorker();
    const options = registerSW.mock.calls[0][0] as RegisterOptions;
    expect(() => options.onRegisteredSW?.("/sw.js", undefined)).not.toThrow();
  });
});
