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

// bottomBarIsOnScreen reads the foot of the shell against the viewport.
// The shell holds overflow-hidden, so a foot below the fold is a foot no
// reader reaches, and the page then refuses to scroll down (D-625).
async function bottomBarIsOnScreen(page: import("@playwright/test").Page, where: string) {
  const foot = page.locator("footer").last();
  await expect(foot, `${where} has no bottom bar`).toBeVisible();
  const box = await foot.boundingBox();
  const height = page.viewportSize()?.height ?? 0;
  const bottom = (box?.y ?? Number.MAX_SAFE_INTEGER) + (box?.height ?? 0);
  // A sub-pixel layout puts the last row a fraction past the edge, and
  // one pixel is not a bar below the fold.
  expect(bottom, `${where} keeps its bottom bar below the fold: it ends at ${bottom} of ${height}`).toBeLessThanOrEqual(height + 1);
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
  await bottomBarIsOnScreen(page, "the deck list");

  // The main region is the one that scrolls, and the shell never grows
  // past the screen. A flex item of min-height auto grows instead of
  // scrolling, and the shell then clips what nobody can reach (D-625).
  const shell = await page.evaluate(() => {
    const main = document.querySelector("main");
    return {
      docHeight: document.documentElement.scrollHeight,
      viewport: window.innerHeight,
      mainScrolls: main ? main.scrollHeight > main.clientHeight : false,
      mainOverflow: main ? getComputedStyle(main).overflowY : "",
    };
  });
  expect(shell.docHeight, "the document grows past the screen").toBeLessThanOrEqual(shell.viewport + 1);
  expect(shell.mainOverflow, "the main region does not scroll on its own").toBe("auto");
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
  await bottomBarIsOnScreen(page, "the deck screen");

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
