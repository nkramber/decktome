import { registerSW } from "virtual:pwa-register";

// startServiceWorker registers the worker of the build. The plugin runs
// in autoUpdate mode, so the page reloads when a new worker takes over,
// and a reader never runs an old shell against a new API (D-621, F-122).
// An unsent draft lives in memory, and the reload drops it (D-692).
export function startServiceWorker(): void {
  registerSW({ immediate: true });
}
