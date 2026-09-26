import { defineConfig, devices } from "@playwright/test";

// The live web lane of D-960. One message on the deployed web app, as the
// test account of .env, in headless Chromium. `make live-web` runs it,
// and the deployed API calls the real providers, so a run costs money.
// The smoke flow of playwright.config.ts never reads this folder.
export default defineConfig({
  testDir: "./live",
  fullyParallel: false,
  forbidOnly: true,
  retries: 0,
  workers: 1,
  reporter: [["list"]],
  // A cold API waits for its card index for 10 to 21 seconds (F-176),
  // and a turn that builds a deck takes minutes.
  timeout: 600_000,
  expect: { timeout: 30_000 },
  outputDir: process.env.LIVE_OUT ? `${process.env.LIVE_OUT}/playwright` : "test-results-live",
  use: {
    ...devices["Desktop Chrome"],
    baseURL: process.env.LIVE_BASE_URL ?? "https://decktome.com",
    // A trace records the typed password, so the lane keeps none.
    trace: "off",
    screenshot: "off",
  },
  projects: [{ name: "chromium" }],
});
