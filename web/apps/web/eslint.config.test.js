import { Linter } from "eslint";
import { describe, expect, it } from "vitest";

import { featureBoundary } from "./eslint.config.js";

// The boundary rule runs on two import strings through ESLint's Linter.
// No file matches here: the test feeds the rule of one feature directly.
function messages(feature, source) {
  const linter = new Linter();
  const config = {
    files: ["**/*.ts"],
    languageOptions: { ecmaVersion: 2022, sourceType: "module" },
    rules: featureBoundary(feature).rules,
  };
  return linter.verify(source, [config], { filename: `src/features/${feature}/x.ts` }).map((m) => m.message);
}

describe("featureBoundary", () => {
  it("blocks a sibling feature by every relative path form", () => {
    expect(messages("chat", 'import x from "../auth/x";')).toHaveLength(1);
    expect(messages("chat", 'import x from "../../auth/x";')).toHaveLength(1);
    expect(messages("chat", 'import x from "../../features/auth/x";')).toHaveLength(1);
    expect(messages("chat", 'import x from "../../features/auth";')).toHaveLength(1);
  });

  it("allows the listed sibling, lib, and app/components", () => {
    expect(messages("chat", 'import x from "../deck/deck-view";')).toEqual([]);
    expect(messages("chat", 'import x from "../../features/deck/deck-view";')).toEqual([]);
    expect(messages("chat", 'import x from "../../lib/api";')).toEqual([]);
    expect(messages("chat", 'import x from "../../app/components/health-footer";')).toEqual([]);
    expect(messages("export", 'import x from "../deck/deck-stats";')).toEqual([]);
    expect(messages("export", "export {};")).toEqual([]);
  });

  it("blocks the router, the layout, and the providers", () => {
    expect(messages("deck", 'import x from "../../app/router";')).toHaveLength(1);
    expect(messages("deck", 'import x from "../chat/use-chat";')).toHaveLength(1);
  });
});
