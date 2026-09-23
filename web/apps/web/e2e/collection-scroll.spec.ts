/// <reference types="node" />
import path from "node:path";

import { expect, type Page, test } from "@playwright/test";

// The gate of PR-62. The collection page shows one scrollbar, the binder
// title and search stay pinned, and a button returns the reader to the
// binder top (D-806 to D-808).
//
// The spec runs in make smoke over the emulators. With SMOKE_BASE_URL,
// API_BUILD_EMAIL, and API_BUILD_PASSWORD it signs in to the deployed app
// as the check account of D-779 instead, for the live check (D-809).

const collectionCsv = path.resolve(import.meta.dirname, "../../../../go/internal/collections/testdata/manabox_collection.csv");

const widths = [
  { name: "a phone", viewport: { width: 390, height: 844 }, hasTouch: true, isMobile: true },
  { name: "a desktop", viewport: { width: 1280, height: 800 }, hasTouch: false, isMobile: false },
];

async function signIn(page: Page) {
  await page.goto("/sign-in");
  const email = process.env.API_BUILD_EMAIL;
  const password = process.env.API_BUILD_PASSWORD;
  if (email && password) {
    await page.getByLabel("Email").fill(email);
    await page.getByLabel("Password").fill(password);
    await page.getByRole("button", { name: "Sign in", exact: true }).click();
  } else {
    await page.getByRole("button", { name: "New here? Create account" }).click();
    await page.getByLabel("Email").fill(`scroll-${Date.now()}@example.com`);
    await page.getByLabel("Password").fill("scroll-pass");
    await page.getByRole("button", { name: "Create account", exact: true }).click();
  }
  await expect(page).not.toHaveURL(/\/sign-in/);
}

// scrollAreas names every element that scrolls up and down: its overflow
// allows a scroll, and its content is taller than its box.
async function scrollAreas(page: Page): Promise<string[]> {
  return page.evaluate(() =>
    [...document.querySelectorAll("*")]
      .filter((el) => {
        const s = getComputedStyle(el);
        return (s.overflowY === "auto" || s.overflowY === "scroll") && el.scrollHeight > el.clientHeight + 1;
      })
      .map((el) => el.tagName.toLowerCase() + (el.getAttribute("data-testid") ? `[${el.getAttribute("data-testid")}]` : "")),
  );
}

// binderTop is the distance from the top of the page content to the
// binder, which is where the button moves the page.
async function binderTop(page: Page): Promise<number> {
  return page.evaluate(() => {
    const main = document.querySelector("main")!;
    const binder = document.querySelector('section[aria-labelledby="binder-grid-title"]')!;
    return binder.getBoundingClientRect().top - main.getBoundingClientRect().top + main.scrollTop;
  });
}

async function scrollMain(page: Page, top: number) {
  await page.evaluate((y) => document.querySelector("main")!.scrollTo({ top: y, behavior: "instant" }), top);
}

const mainScrollTop = (page: Page) => page.evaluate(() => document.querySelector("main")!.scrollTop);

for (const width of widths) {
  test.describe(`the collection page on ${width.name}`, () => {
    test.use({ viewport: width.viewport, hasTouch: width.hasTouch, isMobile: width.isMobile });

    test("scrolls once, pins the binder top, and returns to it in one click", async ({ page }) => {
      await signIn(page);
      await page.goto("/collection");
      await page.getByRole("button", { name: "Upload a collection" }).click();
      await page.locator('input[type="file"]').setInputFiles(collectionCsv);
      await page.getByRole("button", { name: "Upload", exact: true }).click();
      await expect(page.getByRole("heading", { name: "Import result" })).toBeVisible({ timeout: 60_000 });
      await page.getByRole("button", { name: "Done" }).click();
      await expect(page.getByTestId("binder-tile").first()).toBeVisible({ timeout: 30_000 });

      const button = page.getByRole("button", { name: "Back to the binder top" });
      const bar = page.getByTestId("binder-bar");
      const search = page.getByRole("searchbox", { name: "Search the binder by card name" });
      const main = page.locator("main");

      // One scrollbar: the page, and no box inside it.
      expect(await scrollAreas(page), "more than one element scrolls up and down").toEqual(["main"]);
      await expect(button).toHaveCount(0);

      // Three screens into the binder, the title and the search sit at
      // the top of the page (D-807).
      const top = await binderTop(page);
      const screen = await main.evaluate((m) => m.clientHeight);
      await scrollMain(page, top + 3 * screen);
      const mainBox = (await main.boundingBox())!;
      const barBox = (await bar.boundingBox())!;
      expect(Math.abs(barBox.y - mainBox.y), "the binder title leaves the top of the page").toBeLessThanOrEqual(1);
      await expect(search).toBeInViewport();
      expect(await scrollAreas(page), "a scroll of the binder adds a second scroll area").toEqual(["main"]);
      if (width.isMobile) {
        expect(barBox.height, "the pinned block holds more than a quarter of the phone screen").toBeLessThanOrEqual(width.viewport.height * 0.25);
      } else {
        const filters = (await page.getByTestId("binder-filters").boundingBox())!;
        expect(Math.abs(filters.y - (barBox.y + barBox.height)), "the filters do not pin below the title").toBeLessThanOrEqual(1);
      }

      // The next page of rows loads when the reader comes near the end.
      const rows = page.getByTestId("binder-rows");
      const before = (await rows.boundingBox())!.height;
      await scrollMain(page, await main.evaluate((m) => m.scrollHeight));
      await expect.poll(async () => (await rows.boundingBox())!.height, { message: "no next page loaded", timeout: 30_000 }).toBeGreaterThan(before);

      // The button moves the page to the binder top, and the focus to
      // the binder heading (D-808).
      await expect(button).toBeVisible();
      await button.click();
      await expect.poll(async () => Math.abs((await mainScrollTop(page)) - top), { message: "the page never reached the binder top" }).toBeLessThanOrEqual(1);
      await expect(bar).toBeInViewport();
      await expect(page.locator('[data-testid="binder-rows"] [data-index="0"]')).toBeInViewport();
      await expect(page.getByRole("heading", { name: "The binder" })).toBeFocused();
      await expect(button).toHaveCount(0);

      // A reader who asks for less motion gets the move at once.
      await page.emulateMedia({ reducedMotion: "reduce" });
      await scrollMain(page, top + 3 * screen);
      await button.click();
      expect(Math.abs((await mainScrollTop(page)) - top), "the move is not instant under reduced motion").toBeLessThanOrEqual(1);
    });
  });
}
