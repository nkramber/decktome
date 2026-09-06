import { defineConfig, devices } from "@playwright/test";

// The smoke flow of PR-23 (D-313). One flow, one browser, one worker:
// it clicks through the whole app once over the emulators and the fake
// provider, and it costs nothing. scripts/smoke.sh starts the stack and
// runs it, so this file names no server. SMOKE_BASE_URL points the flow
// at another web origin.
export default defineConfig({
  testDir: "./e2e",
  fullyParallel: false,
  forbidOnly: !!process.env.CI,
  retries: 0,
  workers: 1,
  reporter: process.env.CI ? [["list"], ["html", { open: "never" }]] : [["list"]],
  // A build over the trimmed snapshot takes seconds, and the whole flow
  // stays under a minute. The limit leaves room for a cold API.
  timeout: 180_000,
  expect: { timeout: 15_000 },
  outputDir: "test-results",
  use: {
    ...devices["Desktop Chrome"],
    baseURL: process.env.SMOKE_BASE_URL ?? "http://127.0.0.1:5180",
    trace: "retain-on-failure",
    screenshot: "only-on-failure",
  },
  projects: [{ name: "chromium" }],
});
