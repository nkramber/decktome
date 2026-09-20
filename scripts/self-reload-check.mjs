// The installed app reloads itself when a new service worker takes over
// (D-692). This check proves that rule against two releases, for nothing.
//
// Usage:
//   node scripts/self-reload-check.mjs [old-ref]
//
// The old ref defaults to the commit before the newest commit that
// changed `web/`. The new release is the build of the current tree.
//
// Firebase Hosting serves the current release alone, so the assets of an
// older release never answer on decktome.com. The check builds both
// releases, serves the old one, and swaps the server to the new one. The
// page then gets one reload, which is the reader who opens the app.
//
// CAUTION: the check needs Node 22.23.2 and a pnpm store on this machine.
// It calls no provider and costs nothing.

import { execFileSync, execSync } from "node:child_process";
import { createReadStream, existsSync, rmSync, statSync } from "node:fs";
import { createServer } from "node:http";
import { createRequire } from "node:module";
import { extname, join, resolve } from "node:path";

const REPO = resolve(import.meta.dirname, "..");
const WEB = join(REPO, "web");
const DIST = join(WEB, "apps/web/dist");
const WORK = join(REPO, ".local/self-reload");
const PORT = Number(process.env.SELF_RELOAD_PORT ?? 8123);

const TYPES = {
  ".css": "text/css",
  ".html": "text/html; charset=utf-8",
  ".js": "text/javascript",
  ".json": "application/json",
  ".png": "image/png",
  ".svg": "image/svg+xml",
  ".webmanifest": "application/manifest+json",
  ".woff2": "font/woff2",
};

const git = (...args) => execFileSync("git", args, { cwd: REPO }).toString().trim();
const run = (cmd, cwd) => execSync(cmd, { cwd, stdio: "inherit" });

// The entry chunk of a build carries a content hash, so its name names
// the release.
function entryChunk(dir) {
  const html = execFileSync("grep", ["-o", "assets/index-[A-Za-z0-9_-]*\\.js", join(dir, "index.html")]);
  return "/" + html.toString().trim().split("\n")[0];
}

function buildOldRelease(ref) {
  const sha = git("rev-parse", "--short", ref);
  const tree = join(WORK, sha);
  rmSync(tree, { recursive: true, force: true });
  git("worktree", "prune");
  run(`git worktree add --detach ${tree} ${sha}`, REPO);
  run("pnpm install --offline --frozen-lockfile", join(tree, "web"));
  run("pnpm --filter @mtg/web build", join(tree, "web"));
  return { dist: join(tree, "web/apps/web/dist"), tree, sha };
}

// The server answers as Firebase Hosting does: a hashed asset is
// immutable, every other file takes no cache, and an unknown route falls
// back to `index.html`.
function serve(state) {
  const server = createServer((req, res) => {
    const path = req.url.split("?")[0];
    let file = join(state.root, path === "/" ? "index.html" : path.slice(1));
    if (!existsSync(file) || statSync(file).isDirectory()) file = join(state.root, "index.html");
    const ext = extname(file);
    res.setHeader("Content-Type", TYPES[ext] ?? "application/octet-stream");
    res.setHeader(
      "Cache-Control",
      path.startsWith("/assets/") ? "public, max-age=31536000, immutable" : "no-cache",
    );
    createReadStream(file).pipe(res);
  });
  return new Promise((ok) => server.listen(PORT, "127.0.0.1", () => ok(server)));
}

const oldRef = process.argv[2] ?? git("log", "-2", "--format=%H", "--", "web/").split("\n")[1];
if (!oldRef) {
  console.error("no old release. Name a ref: node scripts/self-reload-check.mjs <ref>");
  process.exit(2);
}

const old = buildOldRelease(oldRef);
run("pnpm --filter @mtg/web build", WEB);

const oldChunk = entryChunk(old.dist);
const newChunk = entryChunk(DIST);
console.log(`old release ${old.sha} ${oldChunk}`);
console.log(`new release ${git("rev-parse", "--short", "HEAD")} ${newChunk}`);
if (oldChunk === newChunk) {
  console.error("FAIL: the two releases name one chunk, so no worker updates");
  process.exit(1);
}

const state = { root: old.dist };
const server = await serve(state);
const require = createRequire(import.meta.url);
const { chromium } = require(join(WEB, "apps/web/node_modules/@playwright/test"));
const browser = await chromium.launch();
const page = await (await browser.newContext()).newPage();

const navs = [];
page.on("framenavigated", (f) => {
  if (f === page.mainFrame()) navs.push(Date.now());
});
const shell = () =>
  page.evaluate(() => document.querySelector('script[type="module"]')?.getAttribute("src") ?? "none");

const base = `http://127.0.0.1:${PORT}`;
await page.goto(base, { waitUntil: "load" });
await page.waitForFunction(() => !!navigator.serviceWorker.controller, null, { timeout: 30000 });
console.log(`1. the app holds the old release, shell ${await shell()}`);

state.root = DIST;
console.log("2. the deploy lands, and the server holds the new release");

const before = navs.length;
await page.reload({ waitUntil: "load" });
console.log(`3. the reader opens the app, shell ${await shell()}`);

let pass = true;
try {
  await page.waitForFunction((want) => document.querySelector('script[type="module"]')?.getAttribute("src") === want, newChunk, { timeout: 30000 });
} catch {
  pass = false;
}
console.log(`4. the shell after the wait is ${await shell()}`);
console.log(`5. the navigations after the deploy are ${navs.length - before}`);
console.log(pass ? "PASS: the app reloaded itself" : "FAIL: the app kept the old shell");

await browser.close();
server.close();
rmSync(old.tree, { recursive: true, force: true });
git("worktree", "prune");
process.exit(pass ? 0 : 1);
