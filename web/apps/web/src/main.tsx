import "./index.css";

import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { RouterProvider } from "react-router";

import { Providers } from "./app/providers";
import { createAppRouter } from "./app/router";
import { takeRunningBuild } from "./features/chat/running-build";
import { pendingShell, startServiceWorker, updateReloadTimeoutMs } from "./lib/pwa-register";

// A launch at "/" opens the session of a build that ran when the app
// last closed (D-1034). The router reads the address it starts on.
if (window.location.pathname === "/") {
  const to = takeRunningBuild();
  if (to) window.history.replaceState(null, "", to);
}

const root = document.getElementById("root");
if (!root) {
  throw new Error("root element not found");
}
startServiceWorker();
void boot(root);

// boot draws the app after the update check of a cold start (D-1046). A
// new worker reloads the page into the new shell (D-692, F-206), so the
// page shows a splash and not a home page that goes away seconds later.
// The time limit draws the app while the old worker still controls the
// page and still serves the old chunks. The reload follows when the new
// worker takes control.
async function boot(el: HTMLElement) {
  splash(el, "Deck Tome");
  if (await pendingShell()) {
    splash(el, "Updating Deck Tome…");
    await new Promise((resolve) => setTimeout(resolve, updateReloadTimeoutMs));
  }
  el.replaceChildren();
  createRoot(el).render(
    <StrictMode>
      <Providers>
        <RouterProvider router={createAppRouter()} />
      </Providers>
    </StrictMode>,
  );
}

function splash(el: HTMLElement, text: string) {
  const line = document.createElement("p");
  line.setAttribute("role", "status");
  line.className = "grid h-dvh place-items-center bg-background font-display text-sm text-muted-foreground";
  line.textContent = text;
  el.replaceChildren(line);
}
// A browser that turned push on registers its device again after the
// first paint (D-1005). The module and the SDK load then, and not before.
setTimeout(() => void import("./features/push/push").then((m) => m.refreshPush()), 5_000);
