/// <reference types="node" />
import { readFile } from "node:fs/promises";
import path from "node:path";

import { expect, test } from "@playwright/test";

// The smoke flow of PR-23 (D-313). A reader creates an account on the
// Auth emulator, uploads the fixture export, asks for a deck, reads it,
// and downloads the deck list. The fake provider answers every model
// call from its fixtures (D-552): the classify fixture fills every slot
// on the first message, so the turn asks nothing and builds at once, and
// the generate fixture names the deck. The trimmed snapshot of the repo
// serves the card index, and go/cmd/deck-gate/smoke_test.go proves the
// fixtures against it.

// The export the collection screen reads. The same file feeds the deck
// gate and the trimmed snapshot (D-521).
const collectionCsv = path.resolve(import.meta.dirname, "../../../../go/internal/collections/testdata/manabox_collection.csv");

// The message the reader types. The fake reads none of it, and the word
// rules of the questions agent do: it names no store and no set, or a
// catalog row fires and the turn asks a question. It refuses a spending
// limit on purpose, because the budget row fires on an any-card pool
// and the word rule of D-168 is what closes it without a number.
const request = "Build a lifegain Commander deck led by Karlov of the Ghost Council at bracket 3, with no spending limit.";

test("a reader signs in, uploads a collection, builds a deck, and exports it", async ({ page }) => {
  // Sign in. The emulator takes any email and any password of six or
  // more characters (D-275), and a new account lands on the chat.
  const email = `smoke-${Date.now()}@example.com`;
  await page.goto("/sign-in");
  await page.getByRole("button", { name: "New here? Create account" }).click();
  await page.getByLabel("Email").fill(email);
  await page.getByLabel("Password").fill("smoke-pass");
  await page.getByRole("button", { name: "Create account", exact: true }).click();
  await expect(page).toHaveURL(/\/session\/new$/);

  // Upload the export. The import resolves every row against the card
  // index, and the collection becomes the active one (D-342).
  await page.goto("/collection");
  await page.getByRole("button", { name: "Upload a collection" }).click();
  await page.locator('input[type="file"]').setInputFiles(collectionCsv);
  await page.getByRole("button", { name: "Upload", exact: true }).click();
  await expect(page.getByRole("heading", { name: "Import result" })).toBeVisible({ timeout: 60_000 });
  await page.getByRole("button", { name: "Done" }).click();
  await expect(page.getByTestId("active-collection")).toContainText("Active collection");
  await page.getByRole("button", { name: "Continue to chat" }).click();
  await expect(page).toHaveURL(/\/session\/new$/);
  await expect(page.getByTestId("pool-source")).not.toHaveValue("");

  // Ask for the deck. The turn fills every slot, asks nothing, and
  // builds. The build ends on the deck's own address (D-335), and a
  // question or a failure on the way is the finding.
  await page.getByLabel("Your message").fill(request);
  await page.getByRole("button", { name: "Send" }).click();
  await expect(page).toHaveURL(/\/session\/[^/]+$/);
  await expect
    .poll(
      async () => {
        if (/\/decks\/[^/]+$/.test(page.url())) return "deck";
        const questions = page.getByTestId("open-questions");
        if (await questions.count()) return `question: ${await questions.innerText()}`;
        const recovery = page.getByTestId("recovery");
        if (await recovery.count()) return `failure: ${await recovery.innerText()}`;
        return "working";
      },
      { timeout: 120_000, message: "the turn must end on the deck screen" },
    )
    .toBe("deck");
  await expect(page.getByRole("heading", { name: "lifegain", exact: true })).toBeVisible();
  // The deck is the one the generate fixture names, whole and legal. A
  // miss or a block would buy a repair turn the fake answers with the
  // same list, and the line would read "Not legal".
  await expect(page.getByTestId("legality-line")).toHaveText(/^Legal, /);
  await expect(page.getByText("100 cards")).toBeVisible();

  // Rate the deck. The thumbs up writes the verdict under the reader and
  // thanks them (PR-27, D-557).
  await page.getByRole("group", { name: "Rate this deck" }).getByRole("button", { name: "This helped" }).click();
  await expect(page.getByText("Thank you for your feedback!")).toBeVisible();

  // Export. The download carries the commander and the deck.
  const downloading = page.waitForEvent("download");
  await page.getByRole("button", { name: "Download deck list" }).click();
  const download = await downloading;
  expect(download.suggestedFilename()).toMatch(/\.txt$/);
  const text = await readFile(await download.path(), "utf8");
  expect(text).toContain("Karlov of the Ghost Council");
  expect(text).toMatch(/^Deck$/m);
  await expect(page.getByText(/^Saved .+: \d+ lines\.$/)).toBeVisible();
});
