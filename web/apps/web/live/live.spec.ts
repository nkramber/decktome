/// <reference types="node" />
import { mkdir, writeFile } from "node:fs/promises";
import path from "node:path";

import { expect, test } from "@playwright/test";

// The live web lane of D-960. The test account of .env signs in on the
// deployed web app, sends one message on a new chat, and waits for the
// turn to end: a question, a deck, or a failure. The run writes what a
// reader sees to LIVE_OUT: result.json with each question and the name
// of each of its buttons, and turn.png. The lane asserts no question
// text, because each live check reads its own rule in the result.

const env = (name: string): string => {
  const value = process.env[name] ?? "";
  if (!value) throw new Error(`${name} is empty. Run the lane with make live-web.`);
  return value;
};

type Question = { text: string; buttons: string[] };
type Result = {
  base: string;
  url: string;
  collection: string;
  session: string;
  outcome: string;
  questions: Question[];
  failure: string;
  seconds: number;
};

test("the test account sends one message on the deployed web app", async ({ page, baseURL }) => {
  // The deployed API calls the real providers, so no run starts by accident.
  test.skip(process.env.LIVE_WEB !== "1", "LIVE_WEB=1 is required, because a run costs money");
  const out = env("LIVE_OUT");
  const started = Date.now();

  await page.goto("/sign-in");
  await page.getByLabel("Email").fill(env("LIVE_EMAIL"));
  await page.getByLabel("Password").fill(env("LIVE_PASSWORD"));
  await page.getByRole("button", { name: "Sign in", exact: true }).click();
  await expect(page).not.toHaveURL(/\/sign-in/);

  // The browser store holds the active collection, and a new browser
  // starts on "Any card". So the lane picks the first collection of the
  // account, the one that make api-build imports, as a reader would.
  await page.goto("/session/new");
  const picker = page.getByTestId("pool-source");
  const first = picker.locator("option").nth(1);
  await expect(picker.locator("option")).not.toHaveCount(2);
  const collection = (await first.textContent())?.trim() ?? "";
  await picker.selectOption((await first.getAttribute("value")) ?? "");
  await expect(picker).not.toHaveValue("");
  await page.getByLabel("Your message").fill(env("LIVE_PROMPT"));
  await page.getByRole("button", { name: "Send" }).click();
  await expect(page).toHaveURL(/\/session\/(?!new$)[^/]+$/);
  const session = new URL(page.url()).pathname.split("/").pop() ?? "";

  // The web sends a turn again by itself after a cold-start refusal
  // (D-952), so the poll waits past it.
  let outcome = "working";
  await expect
    .poll(
      async () => {
        if (/\/decks\/[^/]+$/.test(page.url())) outcome = "deck";
        else if (await page.getByTestId("open-questions").count()) outcome = "question";
        else if (await page.getByTestId("recovery").count()) outcome = "failure";
        return outcome;
      },
      { timeout: 540_000, intervals: [1_000] },
    )
    .not.toBe("working");

  const questions: Question[] = [];
  if (outcome === "question") {
    for (const group of await page.getByTestId("open-questions").getByRole("group").all()) {
      const label = (await group.getAttribute("aria-label")) ?? "";
      if (!label.startsWith("Question: ")) continue;
      const buttons = await group
        .locator("button")
        .evaluateAll((els) => els.map((el) => (el.getAttribute("aria-label") ?? el.textContent ?? "").trim()));
      questions.push({ text: label.slice("Question: ".length), buttons });
    }
  }
  const failure = outcome === "failure" ? await page.getByTestId("recovery").innerText() : "";

  const result: Result = {
    base: baseURL ?? "",
    url: page.url(),
    collection,
    session,
    outcome,
    questions,
    failure,
    seconds: Math.round((Date.now() - started) / 1000),
  };
  await mkdir(out, { recursive: true });
  await writeFile(path.join(out, "result.json"), JSON.stringify(result, null, 2) + "\n");
  await page.screenshot({ path: path.join(out, "turn.png"), fullPage: true });
  console.log(`session ${session}, outcome ${outcome}, ${questions.length} question(s), ${result.seconds}s`);
  for (const q of questions) console.log(`  Q ${q.text}\n    buttons: ${q.buttons.join(" | ")}`);
  if (failure) console.log(`  failure: ${failure}`);

  expect(outcome, failure).not.toBe("failure");
});
