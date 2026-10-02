import { afterEach, describe, expect, it, vi } from "vitest";

import { pendingShell, startServiceWorker, swUpdateIntervalMs } from "./pwa-register";

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

// D-1046. A cold start after a deploy drew the old home page, and the
// plugin reloaded it 2 to 5 seconds later. The page now asks first.
describe("pendingShell", () => {
  type Fake = { installing: unknown; waiting: unknown; update: ReturnType<typeof vi.fn> };
  function container(controller: boolean, registration: Fake | undefined): ServiceWorkerContainer {
    return { controller: controller ? {} : null, getRegistration: () => Promise.resolve(registration) } as unknown as ServiceWorkerContainer;
  }
  function fake(found: boolean, fail = false): Fake {
    const r: Fake = { installing: null, waiting: null, update: vi.fn() };
    r.update.mockImplementation(() => {
      if (fail) return Promise.reject(new Error("offline"));
      if (found) r.installing = {};
      return Promise.resolve();
    });
    return r;
  }

  it("draws at once on a first visit, with no worker in control", async () => {
    const r = fake(true);
    expect(await pendingShell(container(false, r))).toBe(false);
    expect(r.update).not.toHaveBeenCalled();
  });

  it("reports a new worker that the check finds", async () => {
    expect(await pendingShell(container(true, fake(true)))).toBe(true);
  });

  it("reports a worker that installs already, with no second check", async () => {
    const r = fake(false);
    r.installing = {};
    expect(await pendingShell(container(true, r))).toBe(true);
    expect(r.update).not.toHaveBeenCalled();
  });

  it("draws when the shell is current, and when the check fails offline", async () => {
    expect(await pendingShell(container(true, fake(false)))).toBe(false);
    expect(await pendingShell(container(true, fake(false, true)))).toBe(false);
  });

  it("draws after the time limit when the check hangs", async () => {
    vi.useFakeTimers();
    const r = fake(false);
    r.update.mockImplementation(() => new Promise(() => {}));
    const answer = pendingShell(container(true, r), 3_000);
    await vi.advanceTimersByTimeAsync(3_000);
    expect(await answer).toBe(false);
  });
});
