/// <reference types="node" />
import { mkdir, writeFile } from "node:fs/promises";
import path from "node:path";

import { type Browser, type BrowserContext, type Page, devices, expect, test } from "@playwright/test";

// The live sweep of D-961. The check account walks every screen of the
// deployed web app on a desktop and on a phone, and each screen reads
// the faults that only a real deploy shows: a console error, an uncaught
// error, a failed request, a page wider than the screen, and a broken
// image. The sweep writes report.json, report.md, and one screenshot for
// each screen to LIVE_OUT, and it fails when it finds a fault.
//
// Money: only a chat turn costs money. A vague first message makes the
// agent ask its question rows, about $0.001 a turn, and the one build
// costs $0.10 to $0.20. Every later screen reads that deck for free.
// LIVE_SWEEP_BUILD=0 sends no message and sweeps the newest deck of the
// account, for nothing.
//
// Writes: the sweep never clicks a thumb or "Report a problem", because
// each one writes a reader verdict (D-635). It opens each rename and
// delete dialog and cancels it. It makes one share link, reads it
// signed out, and revokes it. LIVE_SWEEP_DELETE=1 deletes the new deck
// at the end.

const env = (name: string): string => {
  const value = process.env[name] ?? "";
  if (!value) throw new Error(`${name} is empty. Run the sweep with make live-sweep.`);
  return value;
};

const desktop = { viewport: { width: 1280, height: 800 } };
const phone = { ...devices["iPhone 13"], viewport: { width: 390, height: 844 } };

// A vague message, so the agent asks its rows before it builds.
const defaultPrompt = "Build me a Commander deck from my collection.";
const maxTurns = 8;

type Fault = { kind: string; detail: string };
type Screen = { step: string; view: string; url: string; shot: string; faults: Fault[] };

class Sweep {
  screens: Screen[] = [];
  private pending = new Map<Page, Fault[]>();
  private allow4xx = new Set<Page>();
  private count = 0;

  constructor(private out: string) {}

  // watch records every fault of a page until the next look.
  watch(page: Page) {
    const buf: Fault[] = [];
    this.pending.set(page, buf);
    page.on("console", (m) => {
      if (m.type() !== "error") return;
      // The browser logs each 4xx answer, and an error state expects one.
      if (this.allow4xx.has(page) && /status of 4\d\d/.test(m.text())) return;
      buf.push({ kind: "console", detail: m.text().slice(0, 400) });
    });
    page.on("pageerror", (e) => buf.push({ kind: "pageerror", detail: String(e).slice(0, 400) }));
    page.on("requestfailed", (r) => {
      const why = r.failure()?.errorText ?? "";
      // A navigation cancels the requests of the page it leaves.
      if (why.includes("ERR_ABORTED") || why.includes("NS_BINDING_ABORTED")) return;
      buf.push({ kind: "requestfailed", detail: `${r.method()} ${short(r.url())}: ${why}` });
    });
    page.on("response", (r) => {
      const status = r.status();
      if (status < 400) return;
      // The API refuses a turn while a cold instance loads its card index,
      // and the web sends it again (D-952, D-954).
      if (status === 503 && r.headers()["deck-tome-refusal"] === "index-loading") return;
      if (status < 500 && this.allow4xx.has(page)) return;
      buf.push({ kind: "response", detail: `${status} ${r.request().method()} ${short(r.url())}` });
    });
  }

  // expect4xx lets the next look pass a 4xx answer, for a screen that
  // shows an error state on purpose.
  expect4xx(page: Page, on: boolean) {
    if (on) this.allow4xx.add(page);
    else this.allow4xx.delete(page);
  }

  // look waits for the screen to settle, reads the layout faults, and
  // takes a screenshot.
  async look(page: Page, step: string, view: string) {
    await page.waitForLoadState("networkidle", { timeout: 15_000 }).catch(() => undefined);
    await page.waitForTimeout(500);
    const layout = await page.evaluate(() => {
      const faults: { kind: string; detail: string }[] = [];
      // The page never pans sideways (D-624).
      const doc = document.scrollingElement ?? document.documentElement;
      if (doc.scrollWidth > doc.clientWidth + 1) {
        faults.push({ kind: "overflow", detail: `the page is ${doc.scrollWidth}px wide in ${doc.clientWidth}px` });
      }
      const main = document.querySelector("main");
      if (main && main.scrollWidth > main.clientWidth + 1) {
        const wide = [...main.querySelectorAll<HTMLElement>("*")]
          .filter((el) => el.getBoundingClientRect().right > main.clientWidth + 1)
          .slice(0, 3)
          .map((el) => `${el.tagName.toLowerCase()}${el.dataset.testid ? `[${el.dataset.testid}]` : ""}`);
        faults.push({ kind: "overflow", detail: `main is ${main.scrollWidth}px wide in ${main.clientWidth}px: ${wide.join(", ")}` });
      }
      for (const img of document.images) {
        if (img.complete && img.naturalWidth === 0 && img.currentSrc) {
          faults.push({ kind: "image", detail: `${img.alt || "no alt"}: ${img.currentSrc.slice(0, 160)}` });
        }
      }
      return faults;
    });
    const buf = this.pending.get(page) ?? [];
    const faults = [...buf.splice(0), ...layout];
    this.count++;
    const shot = `${String(this.count).padStart(2, "0")}-${slug(step)}-${view}.png`;
    // The shell fits the screen, and main scrolls inside it (D-364). So
    // the shot unfolds main and each parent of it, then puts them back.
    await page.evaluate(() => {
      const main = document.querySelector("main");
      const chain: HTMLElement[] = [];
      for (let el: HTMLElement | null = main; el; el = el.parentElement) chain.push(el);
      for (const el of chain) {
        el.dataset.liveStyle = el.getAttribute("style") ?? "";
        el.style.height = "auto";
        el.style.maxHeight = "none";
        el.style.overflow = "visible";
      }
    });
    await page.screenshot({ path: path.join(this.out, "shots", shot), fullPage: true });
    await page.evaluate(() => {
      for (const el of document.querySelectorAll<HTMLElement>("[data-live-style]")) {
        el.setAttribute("style", el.dataset.liveStyle ?? "");
        delete el.dataset.liveStyle;
      }
    });
    this.screens.push({ step, view, url: new URL(page.url()).pathname, shot, faults });
  }

  // step runs one part of the sweep. A part that breaks records a flow
  // fault, and the sweep goes on with the next part.
  async step(name: string, pages: Page[], fn: () => Promise<void>) {
    try {
      await fn();
    } catch (e) {
      this.note(name, "flow", String(e).split("\n")[0].slice(0, 400));
      for (const page of pages) await page.keyboard.press("Escape").catch(() => undefined);
    }
  }

  // note records a fault of the flow itself: a step that could not run.
  note(step: string, view: string, detail: string) {
    this.screens.push({ step, view, url: "", shot: "", faults: [{ kind: "flow", detail }] });
  }

  faults(): number {
    return this.screens.reduce((n, s) => n + s.faults.length, 0);
  }
}

const short = (url: string): string => {
  try {
    const u = new URL(url);
    return `${u.host}${u.pathname}`.slice(0, 160);
  } catch {
    return url.slice(0, 160);
  }
};

const slug = (s: string): string =>
  s
    .toLowerCase()
    .replace(/[^a-z0-9]+/g, "-")
    .replace(/^-|-$/g, "")
    .slice(0, 40);

async function signIn(page: Page) {
  await page.goto("/sign-in");
  await page.getByLabel("Email").fill(env("LIVE_EMAIL"));
  await page.getByLabel("Password").fill(env("LIVE_PASSWORD"));
  await page.getByRole("button", { name: "Sign in", exact: true }).click();
  await expect(page).not.toHaveURL(/\/sign-in/);
  await expect(page.getByRole("navigation", { name: "Main" })).toBeVisible();
}

async function context(browser: Browser, options: object, sweep: Sweep): Promise<{ ctx: BrowserContext; page: Page }> {
  const ctx = await browser.newContext({ ...options, acceptDownloads: true });
  const page = await ctx.newPage();
  sweep.watch(page);
  return { ctx, page };
}

// outcome waits until a turn shows a question, a deck, or a failure.
async function outcome(page: Page): Promise<string> {
  let seen = "working";
  await expect
    .poll(
      async () => {
        if (/\/decks\/[^/]+$/.test(page.url())) seen = "deck";
        else if (await page.getByTestId("recovery").count()) seen = "failure";
        else if (await page.getByTestId("open-questions").count()) seen = "question";
        return seen;
      },
      { timeout: 540_000, intervals: [1_000] },
    )
    .not.toBe("working");
  return seen;
}

// answer picks an answer for each open question: the first card tile,
// else the first plain option, else "You decide", else a short text.
async function answer(page: Page): Promise<string[]> {
  const picked: string[] = [];
  const groups = await page.getByTestId("open-questions").getByRole("group", { name: /^Question: / }).all();
  for (const group of groups) {
    // The thumbs sit in a group of their own inside the question.
    const own = await group.locator("button").evaluateAll((els, root) => {
      return els.map((el, i) => ({
        i,
        name: (el.getAttribute("aria-label") ?? el.textContent ?? "").trim(),
        mine: el.closest('[role="group"]') === root,
        option: el.hasAttribute("aria-pressed"),
        disabled: (el as HTMLButtonElement).disabled,
      }));
    }, await group.elementHandle());
    const options = own.filter((b) => b.mine && b.option && !b.disabled);
    const decline = own.find((b) => b.mine && !b.disabled && (b.name === "You decide" || b.name === "No budget"));
    if (options.length) {
      await group.locator("button").nth(options[0].i).click();
      picked.push(options[0].name);
    } else if (decline) {
      await group.locator("button").nth(decline.i).click();
      picked.push(decline.name);
    } else {
      await group.getByRole("textbox").first().fill("You pick.");
      picked.push("text");
    }
  }
  return picked;
}

test("the check account sweeps every screen of the deployed web app", async ({ browser, baseURL }) => {
  // The chat turns call the real providers, so no run starts by accident.
  test.skip(process.env.LIVE_SWEEP !== "1", "LIVE_SWEEP=1 is required, because a run can cost money");
  test.setTimeout(1_800_000);
  const out = env("LIVE_OUT");
  await mkdir(path.join(out, "shots"), { recursive: true });
  const sweep = new Sweep(out);
  const build = process.env.LIVE_SWEEP_BUILD !== "0";
  const started = Date.now();
  const turns: { turn: number; outcome: string; asked: string[]; picked: string[] }[] = [];

  const d = await context(browser, desktop, sweep);
  const p = await context(browser, phone, sweep);
  const anonD = await context(browser, desktop, sweep);
  const anonP = await context(browser, phone, sweep);
  await d.ctx.grantPermissions(["clipboard-read", "clipboard-write"], { origin: baseURL });

  let deckPath = "";
  let sessionPath = "";
  // A break in the main flow still writes the report of every screen
  // before it.
  try {
  // Signed out: the root sends a reader to the sign-in screen.
  for (const [a, view] of [
    [anonD.page, "desktop"],
    [anonP.page, "phone"],
  ] as const) {
    await a.goto("/");
    await expect(a).toHaveURL(/\/sign-in/);
    await sweep.look(a, "signed out root", view);
    await a.getByRole("button", { name: "New here? Create account" }).click();
    await sweep.look(a, "create account form", view);
  }

  await signIn(d.page);
  await sweep.look(d.page, "after sign in", "desktop");
  await signIn(p.page);
  await sweep.look(p.page, "after sign in", "phone");
  // The install hint covers the foot of each later phone shot.
  const hint = p.page.getByRole("button", { name: "Dismiss the install hint" });
  if (await hint.count()) await hint.click();

  await sweep.step("account menu", [d.page], async () => {
    // The account menu, open and closed.
    await d.page.getByRole("button", { name: "Account menu" }).click();
    await sweep.look(d.page, "account menu", "desktop");
    await d.page.keyboard.press("Escape");
  });

  // The collection: the binder, a search, the upload dialog, and the
  // rename and delete dialogs, each one cancelled.
  for (const [page, view] of [
    [d.page, "desktop"],
    [p.page, "phone"],
  ] as const) {
    await page.goto("/collection");
    await expect(page.getByRole("heading", { name: "Your collection" })).toBeVisible();
    await sweep.look(page, "collection", view);
  }
  await sweep.step("collection dialogs", [d.page], async () => {
    const collections = d.page.getByRole("button", { name: /^Rename / });
    if (await collections.count()) {
      await collections.first().click();
      await sweep.look(d.page, "collection rename form", "desktop");
      await d.page.getByRole("button", { name: "Cancel" }).click();
      await d.page.getByRole("button", { name: /^Delete / }).first().click();
      await sweep.look(d.page, "collection delete dialog", "desktop");
      await d.page.getByRole("button", { name: "Keep it" }).click();
    } else {
      sweep.note("collection rename form", "desktop", "the account holds no collection");
    }
  });
  await sweep.step("binder and upload", [d.page], async () => {
    // The binder shows the active collection, and a new browser has none.
    const row = d.page.locator("button[aria-pressed]").first();
    if (await row.count()) {
      await row.click();
      await sweep.look(d.page, "active collection", "desktop");
    }
    const search = d.page.getByRole("searchbox", { name: "Search the binder by card name" });
    if (await search.count()) {
      await search.fill("dragon");
      await sweep.look(d.page, "binder search", "desktop");
      await search.fill("");
      await d.page.getByRole("combobox", { name: "Sort the binder" }).selectOption({ index: 1 });
      await sweep.look(d.page, "binder sort", "desktop");
    }
    await d.page.getByRole("button", { name: "Upload a collection" }).click();
    await sweep.look(d.page, "upload dialog", "desktop");
    await d.page.getByRole("button", { name: "Cancel" }).click();
  });

  // The new chat, before any message.
  for (const [page, view] of [
    [d.page, "desktop"],
    [p.page, "phone"],
  ] as const) {
    await page.goto("/session/new");
    await expect(page.getByLabel("Your message")).toBeVisible();
    await sweep.look(page, "new chat", view);
  }
  await sweep.step("chat delete dialog", [d.page], async () => {
    const unfinished = d.page.getByTestId("unfinished-chats").getByRole("button", { name: /^Delete / });
    if (await unfinished.count()) {
      await unfinished.first().click();
      await sweep.look(d.page, "chat delete dialog", "desktop");
      await d.page.getByRole("button", { name: "Keep it" }).click();
    }
  });

  // The chat: one vague message, an answer to each question, one build.
  if (build) {
    const picker = d.page.getByTestId("pool-source");
    await expect(picker.locator("option")).not.toHaveCount(2);
    await picker.selectOption((await picker.locator("option").nth(1).getAttribute("value")) ?? "");
    await d.page.getByLabel("Your message").fill(process.env.LIVE_PROMPT || defaultPrompt);
    await d.page.getByRole("button", { name: "Send" }).click();
    await expect(d.page).toHaveURL(/\/session\/(?!new$)[^/]+$/);
    sessionPath = new URL(d.page.url()).pathname;
    for (let turn = 1; turn <= maxTurns; turn++) {
      const seen = await outcome(d.page);
      const asked =
        seen === "question"
          ? await d.page
              .getByTestId("open-questions")
              .getByRole("group", { name: /^Question: / })
              .evaluateAll((els) => els.map((el) => (el.getAttribute("aria-label") ?? "").slice("Question: ".length)))
          : [];
      await sweep.look(d.page, `turn ${turn} ${seen}`, "desktop");
      if (seen !== "deck") {
        // The phone reads the same chat, for nothing.
        await p.page.goto(sessionPath);
        await sweep.look(p.page, `turn ${turn} ${seen}`, "phone");
      }
      if (seen === "question") {
        const picked = await answer(d.page);
        turns.push({ turn, outcome: seen, asked, picked });
        await d.page.getByRole("button", { name: "Submit answers" }).click();
        continue;
      }
      turns.push({ turn, outcome: seen, asked, picked: [] });
      if (seen === "deck") deckPath = new URL(d.page.url()).pathname;
      if (seen === "failure") sweep.note(`turn ${turn}`, "desktop", await d.page.getByTestId("recovery").innerText());
      break;
    }
    if (!deckPath && turns.at(-1)?.outcome === "question") sweep.note("chat", "desktop", `no deck after ${maxTurns} turns`);
  }

  // The deck list. With no build, the sweep reads the newest deck.
  for (const [page, view] of [
    [d.page, "desktop"],
    [p.page, "phone"],
  ] as const) {
    await page.goto("/decks");
    await expect(page.getByRole("heading", { name: "Your decks" })).toBeVisible();
    await sweep.look(page, "deck list", view);
  }
  if (!deckPath) {
    const first = d.page.locator('a[href^="/decks/"]').first();
    if (await first.count()) deckPath = (await first.getAttribute("href")) ?? "";
  }
  await sweep.step("deck import dialog", [d.page], async () => {
    await d.page.getByRole("button", { name: "Import a deck" }).click();
    await sweep.look(d.page, "deck import dialog", "desktop");
    await d.page.keyboard.press("Escape");
  });

  if (!deckPath) {
    sweep.note("deck screen", "desktop", "the account holds no deck, and the sweep built none");
  } else {
    // The deck screen and each of its parts.
    for (const [page, view] of [
      [d.page, "desktop"],
      [p.page, "phone"],
    ] as const) {
      await page.goto(deckPath);
      await expect(page.getByTestId("legality-line")).toBeVisible({ timeout: 60_000 });
      await sweep.look(page, "deck screen", view);
    }
    const dp = d.page;
    await sweep.step("card detail", [dp, p.page], async () => {
      for (const [page, view] of [
        [dp, "desktop"],
        [p.page, "phone"],
      ] as const) {
        await page.getByTitle("Open the card detail").first().click();
        await expect(page.getByRole("dialog")).toBeVisible();
        await sweep.look(page, "card detail", view);
        await page.keyboard.press("Escape");
      }
    });
    await sweep.step("sample hand", [dp], async () => {
      await dp.getByRole("button", { name: "Draw seven" }).click();
      await dp.getByRole("button", { name: "Mulligan" }).click();
      await sweep.look(dp, "sample hand", "desktop");
    });
    await sweep.step("buy list", [dp], async () => {
      const buy = dp.getByRole("region", { name: /^Buy list/ });
      if (!(await buy.count())) return;
      await buy.locator("summary").first().click();
      await sweep.look(dp, "buy list open", "desktop");
    });
    await sweep.step("export", [dp], async () => {
      await dp.getByRole("button", { name: "Copy deck list" }).first().click();
      const downloading = dp.waitForEvent("download");
      await dp.getByRole("button", { name: "Download deck list" }).first().click();
      const file = await downloading;
      await file.saveAs(path.join(out, file.suggestedFilename()));
      await sweep.look(dp, "export", "desktop");
    });
    await sweep.step("deck dialogs", [dp], async () => {
      await dp.getByRole("button", { name: "Rename", exact: true }).click();
      await sweep.look(dp, "deck rename dialog", "desktop");
      await dp.getByRole("button", { name: "Cancel" }).click();
      await dp.getByRole("button", { name: "Delete", exact: true }).click();
      await sweep.look(dp, "deck delete dialog", "desktop");
      await dp.getByRole("button", { name: "Keep the deck" }).click();
    });

    // One share link: read it signed out, revoke it, and read it again.
    await sweep.step("share", [dp, anonD.page, anonP.page], async () => {
      await dp.getByRole("button", { name: /^Share/ }).click();
      await dp.getByRole("button", { name: /^Make (a|a new) link$/ }).click();
      const linkBox = dp.getByLabel("The link, shown once");
      await expect(linkBox).not.toHaveValue("");
      const link = new URL(await linkBox.inputValue()).pathname;
      await sweep.look(dp, "share dialog", "desktop");
      for (const [a, view] of [
        [anonD.page, "desktop"],
        [anonP.page, "phone"],
      ] as const) {
        await a.goto(link);
        await expect(a.getByText("A shared deck, read-only.")).toBeVisible({ timeout: 60_000 });
        await sweep.look(a, "shared deck", view);
      }
      await dp.getByRole("button", { name: "Revoke the link" }).click();
      await expect(dp.getByTestId("share-state")).toBeVisible();
      await dp.keyboard.press("Escape");
      sweep.expect4xx(anonD.page, true);
      await anonD.page.goto(link);
      await expect(anonD.page.getByRole("alert")).toBeVisible({ timeout: 60_000 });
      await sweep.look(anonD.page, "revoked share", "desktop");
      sweep.expect4xx(anonD.page, false);
    });
  }

  // The error states: a deck that does not exist, and a bad link.
  await sweep.step("missing deck", [d.page], async () => {
    sweep.expect4xx(d.page, true);
    await d.page.goto("/decks/no-such-deck");
    await expect(d.page.getByRole("link", { name: "Back to your decks" })).toBeVisible({ timeout: 60_000 });
    await sweep.look(d.page, "missing deck", "desktop");
    sweep.expect4xx(d.page, false);
  });

  if (deckPath && build && process.env.LIVE_SWEEP_DELETE === "1") {
    await d.page.goto(deckPath);
    await d.page.getByRole("button", { name: "Delete" }).click();
    await d.page.getByRole("button", { name: "Delete the deck" }).click();
    await expect(d.page).toHaveURL(/\/decks$/);
    await sweep.look(d.page, "after deck delete", "desktop");
  }

  // Sign out on the desktop.
  await sweep.step("sign out", [d.page], async () => {
    await d.page.getByRole("button", { name: "Account menu" }).click();
    await d.page.getByRole("menuitem", { name: "Sign out" }).click();
    await expect(d.page).toHaveURL(/\/sign-in/);
    await sweep.look(d.page, "after sign out", "desktop");
  });

  } catch (e) {
    sweep.note("main flow", "flow", String(e).split("\n")[0].slice(0, 400));
  }

  const seconds = Math.round((Date.now() - started) / 1000);
  const report = { base: baseURL ?? "", build, session: sessionPath, deck: deckPath, seconds, turns, faults: sweep.faults(), screens: sweep.screens };
  await writeFile(path.join(out, "report.json"), JSON.stringify(report, null, 2) + "\n");
  const lines = [
    `# Live sweep`,
    ``,
    `Base: ${report.base}. Build: ${build ? "yes" : "no"}. Chat: ${sessionPath || "none"}. Deck: ${deckPath || "none"}. Time: ${seconds} seconds.`,
    ``,
    `Faults: ${report.faults} on ${sweep.screens.length} screens.`,
    ``,
    `| Turn | Outcome | Questions | Picks |`,
    `|---|---|---|---|`,
    ...turns.map((t) => `| ${t.turn} | ${t.outcome} | ${t.asked.join(" / ")} | ${t.picked.join(" / ")} |`),
    ``,
    `| Screen | View | Path | Faults | Shot |`,
    `|---|---|---|---|---|`,
    ...sweep.screens.map(
      (s) => `| ${s.step} | ${s.view} | ${s.url} | ${s.faults.map((f) => `${f.kind}: ${f.detail}`).join("<br>") || "none"} | ${s.shot} |`,
    ),
  ];
  await writeFile(path.join(out, "report.md"), lines.join("\n") + "\n");
  console.log(`sweep: ${sweep.screens.length} screens, ${report.faults} fault(s), ${turns.length} turn(s), ${seconds}s, deck ${deckPath || "none"}`);
  for (const s of sweep.screens) for (const f of s.faults) console.log(`  ${s.step} (${s.view}): ${f.kind}: ${f.detail}`);

  for (const c of [d, p, anonD, anonP]) await c.ctx.close();
  expect(report.faults, "the sweep found faults: read report.md").toBe(0);
});
