import tailwindcss from "@tailwindcss/vite";
import react from "@vitejs/plugin-react";
import { VitePWA } from "vite-plugin-pwa";
import { configDefaults, defineConfig } from "vitest/config";

import { webManifest } from "./src/lib/pwa";

// The Go API listens on :8090 in local dev, and Connect-RPC routes live
// under /mtg.v1.*. The dev server proxies them, so the browser talks to
// one origin.
export default defineConfig({
  plugins: [
    react(),
    tailwindcss(),
    // The installable app of PR-25. The plugin writes the manifest and
    // the service worker from the build, so the precache list carries
    // the hashed asset names of that build.
    VitePWA({
      // The app updates itself. A reader who added it to the Home
      // Screen must never hold an old shell against a new API, and
      // F-57 is what a stale cache costs.
      registerType: "autoUpdate",
      includeAssets: ["favicon.svg", "apple-touch-icon.png"],
      manifest: { ...webManifest, icons: [...webManifest.icons] },
      workbox: {
        globPatterns: ["**/*.{js,css,html,svg,png,woff2}"],
        // The API answers a POST per user, so no response of it is ever
        // cached, and the navigation fallback never swallows its routes.
        navigateFallbackDenylist: [/^\/mtg\.v1\./, /^\/healthz/, /^\/readyz/],
      },
      // The service worker belongs to a build. `make dev` and the smoke
      // flow of PR-23 run the dev server, and neither registers one.
      devOptions: { enabled: false },
    }),
  ],
  server: {
    port: 5180,
    proxy: {
      "/mtg.v1.": { target: "http://localhost:8090", changeOrigin: true },
      "/healthz": { target: "http://localhost:8090", changeOrigin: true },
      "/readyz": { target: "http://localhost:8090", changeOrigin: true },
    },
  },
  test: {
    environment: "jsdom",
    setupFiles: ["./src/test-setup.ts"],
    // The Playwright flow under e2e is not a unit test (PR-23).
    exclude: [...configDefaults.exclude, "e2e/**"],
  },
});
