// Renders the app icons from the tome mark (PR-25). It runs by hand, and
// its output is committed, so no build step needs a browser.
//
//   node scripts/make-icons.mjs
//
// Chromium comes from Playwright, which the smoke flow of PR-23 already
// installs. No image library joins the toolchain for four files.
import { readFileSync, writeFileSync } from "node:fs";
import path from "node:path";

import { chromium } from "@playwright/test";

const here = path.dirname(new URL(import.meta.url).pathname);
const pub = path.join(here, "..", "public");
const mark = readFileSync(path.join(pub, "favicon.svg"), "utf8");

// A maskable icon is cropped to a circle on Android, so the mark sits in
// the middle 80 percent and the ground fills the rest (the W3C maskable
// safe zone). The plain icon keeps the rounded square of the favicon.
const maskable = mark
  .replace('viewBox="0 0 32 32"', 'viewBox="-4 -4 40 40"')
  .replace('<rect width="32" height="32" rx="7"', '<rect x="-4" y="-4" width="40" height="40" rx="0"');

const icons = [
  { file: "icon-192.png", size: 192, svg: mark },
  { file: "icon-512.png", size: 512, svg: mark },
  { file: "icon-maskable-512.png", size: 512, svg: maskable },
  // iOS reads no SVG for a Home Screen icon, and 180 is its size.
  { file: "apple-touch-icon.png", size: 180, svg: mark },
];

const browser = await chromium.launch();
for (const { file, size, svg } of icons) {
  const page = await browser.newPage({ viewport: { width: size, height: size }, deviceScaleFactor: 1 });
  const data = `data:image/svg+xml;base64,${Buffer.from(svg).toString("base64")}`;
  await page.setContent(
    `<style>html,body{margin:0;padding:0;background:transparent}img{display:block;width:${size}px;height:${size}px}</style>` +
      `<img src="${data}">`,
  );
  await page.locator("img").waitFor();
  const png = await page.screenshot({ omitBackground: true });
  writeFileSync(path.join(pub, file), png);
  console.log(`wrote ${file} at ${size}px, ${png.length} bytes`);
  await page.close();
}
await browser.close();
