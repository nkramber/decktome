import tailwindcss from "@tailwindcss/vite";
import react from "@vitejs/plugin-react";
import { defineConfig } from "vitest/config";

// The Go API listens on :8090 in local dev, and Connect-RPC routes live
// under /mtg.v1.*. The dev server proxies them, so the browser talks to
// one origin.
export default defineConfig({
  plugins: [react(), tailwindcss()],
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
  },
});
