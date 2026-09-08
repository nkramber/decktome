import { readFileSync } from "node:fs";
import path from "node:path";

import { describe, expect, it } from "vitest";

import { pageBackground, webManifest } from "./pwa";

// The installable gate of PR-25. A browser installs a web app when the
// manifest names the app, a start url, a display mode away from
// `browser`, and an icon of 192 pixels or more. This reads every one of
// those, and it reads the icon files themselves, so a manifest that
// names a file nobody generated fails here and not on a phone.

const publicDir = path.join(import.meta.dirname, "..", "..", "public");

// pngSize reads the width and the height out of the IHDR chunk, which
// is the first chunk of every PNG. No image library joins the toolchain
// to check four files.
function pngSize(file: string): { width: number; height: number } {
  const buf = readFileSync(path.join(publicDir, file));
  expect(buf.subarray(1, 4).toString("ascii"), `${file} is a PNG`).toBe("PNG");
  return { width: buf.readUInt32BE(16), height: buf.readUInt32BE(20) };
}

describe("the web manifest", () => {
  it("names the app, a start url, and a standalone display", () => {
    expect(webManifest.name).toBe("Deck Tome");
    expect(webManifest.short_name).toBe("Deck Tome");
    expect(webManifest.start_url).toBe("/");
    expect(webManifest.scope).toBe("/");
    // A browser installs nothing that displays as a plain browser tab.
    expect(webManifest.display).not.toBe("browser");
    expect(webManifest.display).toBe("standalone");
  });

  it("paints the dark ground of the app, and never a default white", () => {
    const tokens = readFileSync(path.join(import.meta.dirname, "..", "styles", "tokens.css"), "utf8");
    // The splash screen must match the page the reader opens, so the
    // colour comes from the token and is not typed twice.
    expect(tokens).toContain(`--background: ${pageBackground};`);
    expect(webManifest.theme_color).toBe(pageBackground);
    expect(webManifest.background_color).toBe(pageBackground);
  });

  it("holds an icon of 192 pixels, one of 512, and a maskable one", () => {
    const sizes = webManifest.icons.map((i) => i.sizes);
    expect(sizes).toContain("192x192");
    expect(sizes).toContain("512x512");
    expect(webManifest.icons.some((i) => "purpose" in i && i.purpose === "maskable")).toBe(true);
  });

  it("every icon it names is a PNG of the size it claims", () => {
    for (const icon of webManifest.icons) {
      const [w, h] = icon.sizes.split("x").map(Number);
      expect(pngSize(icon.src.replace(/^\//, ""))).toEqual({ width: w, height: h });
    }
  });

  // D-624: the owner asked for a page that never zooms and never pans
  // sideways. The rule costs the axe meta-viewport check, and the PR-25
  // gate carries the exception. This test makes the choice deliberate:
  // nobody restores the zoom by accident, and nobody removes it either.
  it("forbids the pinch zoom the owner asked to stop", () => {
    const html = readFileSync(path.join(import.meta.dirname, "..", "..", "index.html"), "utf8");
    const viewport = /<meta name="viewport" content="([^"]+)"/.exec(html)?.[1] ?? "";
    expect(viewport).toContain("width=device-width");
    expect(viewport).toContain("user-scalable=no");
    expect(viewport).toContain("maximum-scale=1");
  });

  it("holds the 180 pixel icon iOS reads for a Home Screen", () => {
    // iOS reads no manifest icon and no SVG, so index.html links this
    // one by name. A missing file gives the Home Screen a screenshot.
    expect(pngSize("apple-touch-icon.png")).toEqual({ width: 180, height: 180 });
    const html = readFileSync(path.join(import.meta.dirname, "..", "..", "index.html"), "utf8");
    expect(html).toContain('rel="apple-touch-icon" href="/apple-touch-icon.png"');
    expect(html).toContain(`content="${pageBackground}"`);
  });
});
