import { ImportSource, UnresolvedReason } from "@mtg/api-client/mtg/v1/collection_pb";
import { describe, expect, it } from "vitest";

import { reasonLabel, sourceLabel } from "./import-result";

describe("reasonLabel", () => {
  it("reads the full enum name, the short name, and the number", () => {
    expect(reasonLabel("UNRESOLVED_REASON_NON_ENGLISH")).toBe("Non-English printing: the app reads English cards only");
    expect(reasonLabel("BAD_ROW")).toBe("Bad row: the line does not parse in the format of the file");
    expect(reasonLabel(UnresolvedReason.NOT_PLAYABLE)).toBe("Not a playable card: a token, emblem, or art card");
  });

  it("returns an unknown name as is", () => {
    expect(reasonLabel("UNRESOLVED_REASON_SOMETHING_NEW")).toBe("SOMETHING_NEW");
    expect(reasonLabel(999)).toBe("999");
  });
});

// D-796: the server stores the format it read, and an older collection
// holds none, so it names nothing.
describe("sourceLabel", () => {
  it("names each format the server reads, and nothing for an unknown one", () => {
    expect(sourceLabel(ImportSource.MANABOX_CSV)).toBe("ManaBox export");
    expect(sourceLabel(ImportSource.MOXFIELD_CSV)).toBe("Moxfield export");
    expect(sourceLabel(ImportSource.ARENA_TEXT)).toBe("Arena list");
    expect(sourceLabel(ImportSource.UNSPECIFIED)).toBeUndefined();
    expect(sourceLabel(undefined)).toBeUndefined();
  });
});
