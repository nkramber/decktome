import { UnresolvedReason } from "@mtg/api-client/mtg/v1/collection_pb";
import { describe, expect, it } from "vitest";

import { reasonLabel } from "./import-result";

describe("reasonLabel", () => {
  it("reads the full enum name, the short name, and the number", () => {
    expect(reasonLabel("UNRESOLVED_REASON_NON_ENGLISH")).toBe("Non-English printing: the app reads English cards only");
    expect(reasonLabel("BAD_ROW")).toBe("Bad row: the line does not parse as a ManaBox row");
    expect(reasonLabel(UnresolvedReason.NOT_PLAYABLE)).toBe("Not a playable card: a token, emblem, or art card");
  });

  it("returns an unknown name as is", () => {
    expect(reasonLabel("UNRESOLVED_REASON_SOMETHING_NEW")).toBe("SOMETHING_NEW");
    expect(reasonLabel(999)).toBe("999");
  });
});
