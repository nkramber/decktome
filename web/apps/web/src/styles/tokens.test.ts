import { readFileSync } from "node:fs";
import path from "node:path";

import { describe, expect, it } from "vitest";

// WCAG AA asks 4.5:1 for text (REV-049, D-934). jsdom computes no
// contrast, so this reads the tokens and computes each pair of text and
// ground that the product shows.
const css = readFileSync(path.join(import.meta.dirname, "tokens.css"), "utf8");

function token(name: string): number[] {
  const m = new RegExp(`--${name}:\\s*#([0-9a-fA-F]{6});`).exec(css);
  if (!m) throw new Error(`no token --${name}`);
  return [0, 2, 4].map((i) => parseInt(m[1].slice(i, i + 2), 16) / 255);
}

function luminance(c: number[]): number {
  const [r, g, b] = c.map((x) => (x <= 0.03928 ? x / 12.92 : ((x + 0.055) / 1.055) ** 2.4));
  return 0.2126 * r + 0.7152 * g + 0.0722 * b;
}

function ratio(a: number[], b: number[]): number {
  const [hi, lo] = [luminance(a), luminance(b)].sort((x, y) => y - x);
  return (hi + 0.05) / (lo + 0.05);
}

function over(fg: number[], bg: number[], alpha: number): number[] {
  return fg.map((f, i) => alpha * f + (1 - alpha) * bg[i]);
}

describe("the contrast of the tokens", () => {
  const pairs: [string, () => number[], () => number[]][] = [
    ["the user bubble", () => token("accent-foreground"), () => token("accent")],
    ["the danger line on the page", () => token("danger"), () => token("background")],
    ["the danger line on a card", () => token("danger"), () => token("card")],
    ["the danger line on its own tint", () => token("danger"), () => over(token("danger"), token("card"), 0.1)],
    ["the text of a danger button", () => token("danger-foreground"), () => token("danger")],
    ["muted text on a card", () => token("muted-foreground"), () => token("card")],
    ["the text of a primary button", () => token("primary-foreground"), () => token("primary")],
    ["body text on a card", () => token("foreground"), () => token("card")],
  ];
  it.each(pairs)("%s reads 4.5:1 or more", (_, fg, bg) => {
    expect(ratio(fg(), bg())).toBeGreaterThanOrEqual(4.5);
  });
});
