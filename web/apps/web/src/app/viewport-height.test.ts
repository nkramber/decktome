import { readdirSync, readFileSync, statSync } from "node:fs";
import path from "node:path";

import { describe, expect, it } from "vitest";

// D-625: the shell reads the dynamic viewport height and never the
// static one. 100vh on a phone is the height with the browser bar
// collapsed, so a shell of that height puts its own foot below the fold.
// The shell holds overflow-hidden, so nothing scrolls to reach it, and
// the reader meets a page that will not scroll down.
//
// This walks the source, so the rule holds for a file nobody has written
// yet. `dvh` reads the height that is on the screen now.

const src = path.join(import.meta.dirname, "..");

function sourceFiles(dir: string): string[] {
  const out: string[] = [];
  for (const name of readdirSync(dir)) {
    const full = path.join(dir, name);
    if (statSync(full).isDirectory()) {
      out.push(...sourceFiles(full));
      continue;
    }
    if (/\.(tsx?|css)$/.test(name) && !/\.test\.tsx?$/.test(name)) out.push(full);
  }
  return out;
}

describe("the viewport height", () => {
  it("no source reads the static viewport height", () => {
    const bad: string[] = [];
    for (const file of sourceFiles(src)) {
      const text = readFileSync(file, "utf8");
      for (const [i, line] of text.split("\n").entries()) {
        // A print rule is not a screen rule, and a comment names the
        // unit to explain it.
        if (line.trimStart().startsWith("//") || line.trimStart().startsWith("*")) continue;
        if (/\bh-screen\b|\bmin-h-screen\b|100vh/.test(line)) {
          bad.push(`${path.relative(src, file)}:${i + 1}: ${line.trim()}`);
        }
      }
    }
    expect(bad, `these read the static viewport height, and a phone needs dvh (D-625):\n${bad.join("\n")}`).toEqual([]);
  });

  it("the shell takes the dynamic viewport height", () => {
    const layout = readFileSync(path.join(src, "app", "layout.tsx"), "utf8");
    expect(layout).toContain("h-dvh");
  });
});
