/// <reference types="node" />
import path from "node:path";

import { expect, test } from "@playwright/test";

// The phone gate of PR-25. A page on a phone scrolls up and down, and
// never sideways. The owner read the deployed app on 2026-09-08 and
// found it scrolled left and right.
//
// The check runs at 390 by 844, which is the iPhone the roadmap names,
// with a coarse pointer. It reads the document width against the
// viewport, and it names the widest elements when the two differ, so a
// failure says which element to fix.

const phone = { width: 390, height: 844 };
const collectionCsv = path.resolve(import.meta.dirname, "../../../../go/internal/collections/testdata/manabox_collection.csv");

test.use({ viewport: phone, hasTouch: true, isMobile: true });

// widerThanViewport names every element that reaches past the right
// edge. The message of a failure is the list, so nobody hunts by eye.
async function overflow(page: import("@playwright/test").Page) {
  return page.evaluate(() => {
    const doc = document.documentElement;
    const over = [];
    for (const el of Array.from(document.querySelectorAll("*"))) {
      const r = el.getBoundingClientRect();
      if (r.width === 0 || r.height === 0) continue;
      if (Math.ceil(r.right) > doc.clientWidth + 1) {
        over.push(`${el.tagName.toLowerCase()}.${String(el.className).slice(0, 60)} right=${Math.round(r.right)}`);
      }
    }
    return { scrollWidth: doc.scrollWidth, clientWidth: doc.clientWidth, over: over.slice(0, 8) };
  });
}

async function readsUpAndDownAlone(page: import("@playwright/test").Page, where: string) {
  const got = await overflow(page);
  expect(got.scrollWidth, `${where} scrolls sideways. The widest elements: ${got.over.join(" | ")}`).toBeLessThanOrEqual(got.clientWidth);
}

// theShellFitsTheScreen reads the document against the viewport. The
// app shell is the whole window and the main region is the one thing
// that scrolls, so a document taller than the screen is a page that
// bounces and shows empty ground under the app (D-625, D-626).
async function theShellFitsTheScreen(page: import("@playwright/test").Page, where: string) {
  const got = await page.evaluate(() => ({
    docHeight: document.documentElement.scrollHeight,
    bodyHeight: document.body.scrollHeight,
    viewport: window.innerHeight,
    bodyOverflow: getComputedStyle(document.body).overflowY,
  }));
  expect(got.docHeight, `${where}: the document is taller than the screen, so it scrolls under the app`).toBeLessThanOrEqual(got.viewport + 1);
  expect(got.bodyOverflow, `${where}: the body scrolls, and the main region should be the one that does`).toBe("hidden");
}

test("every screen of a phone reads up and down alone", async ({ page }) => {
  const email = `phone-${Date.now()}@example.com`;
  await page.goto("/sign-in");
  await readsUpAndDownAlone(page, "the sign-in page");

  await page.getByRole("button", { name: "New here? Create account" }).click();
  await page.getByLabel("Email").fill(email);
  await page.getByLabel("Password").fill("phone-pass");
  await page.getByRole("button", { name: "Create account", exact: true }).click();
  await expect(page).toHaveURL(/\/session\/new$/);
  await readsUpAndDownAlone(page, "the new chat");

  await page.goto("/collection");
  await readsUpAndDownAlone(page, "the collection page with no collection");
  await page.getByRole("button", { name: "Upload a collection" }).click();
  await readsUpAndDownAlone(page, "the upload dialog");
  await page.locator('input[type="file"]').setInputFiles(collectionCsv);
  await page.getByRole("button", { name: "Upload", exact: true }).click();
  await expect(page.getByRole("heading", { name: "Import result" })).toBeVisible({ timeout: 60_000 });
  await readsUpAndDownAlone(page, "the import result");
  await page.getByRole("button", { name: "Done" }).click();
  await readsUpAndDownAlone(page, "the collection page with a collection");

  await page.goto("/decks");
  await readsUpAndDownAlone(page, "the deck list");
  await theShellFitsTheScreen(page, "the deck list");

  // The main region is the one that scrolls.
  const mainOverflow = await page.evaluate(() => {
    const main = document.querySelector("main");
    return main ? getComputedStyle(main).overflowY : "";
  });
  expect(mainOverflow, "the main region does not scroll on its own").toBe("auto");
});

// The deck screen holds the widest content of the app: the card grid,
// the profile table, the buy list, and the card sheet. It is the screen
// the owner read when the app scrolled sideways.
test("the deck screen of a phone reads up and down alone", async ({ page }) => {
  const email = `phone-deck-${Date.now()}@example.com`;
  await page.goto("/sign-in");
  await page.getByRole("button", { name: "New here? Create account" }).click();
  await page.getByLabel("Email").fill(email);
  await page.getByLabel("Password").fill("phone-pass");
  await page.getByRole("button", { name: "Create account", exact: true }).click();
  await expect(page).toHaveURL(/\/session\/new$/);

  await page.goto("/collection");
  await page.getByRole("button", { name: "Upload a collection" }).click();
  await page.locator('input[type="file"]').setInputFiles(collectionCsv);
  await page.getByRole("button", { name: "Upload", exact: true }).click();
  await expect(page.getByRole("heading", { name: "Import result" })).toBeVisible({ timeout: 60_000 });
  await page.getByRole("button", { name: "Done" }).click();
  await page.getByRole("button", { name: "Continue to chat" }).click();

  await page.getByLabel("Your message").fill(
    "Build a lifegain Commander deck led by Karlov of the Ghost Council at bracket 3, with no spending limit.",
  );
  await page.getByRole("button", { name: "Send" }).click();
  await expect(page.getByTestId("legality-line")).toBeVisible({ timeout: 120_000 });
  await readsUpAndDownAlone(page, "the deck screen");
  await theShellFitsTheScreen(page, "the deck screen");

  // The card sheet opens over the deck, and it is its own width. The
  // tile carries the card name, and a card image needs a network this
  // stack does not promise, so the name is what opens the sheet.
  await page.getByRole("button", { name: "Sol Ring" }).first().click();
  await expect(page.getByRole("dialog")).toBeVisible();
  await readsUpAndDownAlone(page, "the card sheet");
  await page.keyboard.press("Escape");

  // The export panel holds the buy list, which is a table of prices.
  await page.getByRole("region", { name: /^Buy list/ }).scrollIntoViewIfNeeded();
  await readsUpAndDownAlone(page, "the export panel");
});
