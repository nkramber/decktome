import { registerSW } from "virtual:pwa-register";

// The browser checks for a new worker on a navigation alone. An open tab
// or an installed app asks again on this interval and on each return to
// view, so it never keeps an old shell against a new API (D-621, REV-048).
export const swUpdateIntervalMs = 60 * 60 * 1000;

// startServiceWorker registers the worker of the build. The plugin runs
// in autoUpdate mode, so the page reloads when a new worker takes over,
// and a reader never runs an old shell against a new API (D-621, F-122).
// An unsent draft lives in memory, and the reload drops it (D-692).
export function startServiceWorker(): void {
  registerSW({
    immediate: true,
    onRegisteredSW(_swUrl, registration) {
      if (!registration) return;
      // An offline check fails, and the next check tries again.
      const check = () => void registration.update().catch(() => {});
      setInterval(check, swUpdateIntervalMs);
      document.addEventListener("visibilitychange", () => {
        if (document.visibilityState === "visible") check();
      });
    },
  });
}

// The longest wait for the update check of a cold start. An offline
// check fails at once, and a slow link draws the page after this.
export const updateCheckTimeoutMs = 3_000;
// The longest time the splash waits for the reload of a new worker.
export const updateReloadTimeoutMs = 15_000;

// pendingShell reports a new worker at a cold start (D-1046). A page with
// a worker in control asks for the newest one before it draws. When one
// installs, the page reloads into it a few seconds later (D-692), so the
// caller shows a splash in place of a page that goes away. A first visit
// has no worker in control, and it draws at once.
export async function pendingShell(sw: ServiceWorkerContainer | undefined = navigator.serviceWorker, timeoutMs = updateCheckTimeoutMs): Promise<boolean> {
  if (!sw?.controller) return false;
  const registration = await sw.getRegistration().catch(() => undefined);
  if (!registration) return false;
  if (registration.installing || registration.waiting) return true;
  const checked = await Promise.race([
    registration.update().then(
      () => true,
      () => false,
    ),
    new Promise<boolean>((resolve) => setTimeout(() => resolve(false), timeoutMs)),
  ]);
  return checked && Boolean(registration.installing || registration.waiting);
}
