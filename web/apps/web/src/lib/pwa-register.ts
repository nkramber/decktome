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
