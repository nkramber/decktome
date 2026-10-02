import "./index.css";

import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { RouterProvider } from "react-router";

import { Providers } from "./app/providers";
import { createAppRouter } from "./app/router";
import { takeRunningBuild } from "./features/chat/running-build";
import { startServiceWorker } from "./lib/pwa-register";

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
createRoot(root).render(
  <StrictMode>
    <Providers>
      <RouterProvider router={createAppRouter()} />
    </Providers>
  </StrictMode>,
);
startServiceWorker();
// A browser that turned push on registers its device again after the
// first paint (D-1005). The module and the SDK load then, and not before.
setTimeout(() => void import("./features/push/push").then((m) => m.refreshPush()), 5_000);
