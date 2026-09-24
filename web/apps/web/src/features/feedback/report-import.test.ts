import { create } from "@bufbuild/protobuf";
import { Code, ConnectError } from "@connectrpc/connect";
import { UnreadableFileSchema, UnresolvedReason, UnresolvedRowSchema } from "@mtg/api-client/mtg/v1/collection_pb";
import { describe, expect, it } from "vitest";

import { isUnreadable, parseFaults } from "./report-import";

describe("isUnreadable", () => {
  it("reads the detail of a file the app could not read, and nothing else (D-887)", () => {
    const marked = new ConnectError("no format", Code.InvalidArgument, undefined, [{ desc: UnreadableFileSchema, value: create(UnreadableFileSchema) }]);
    expect(isUnreadable(marked)).toBe(true);
    expect(isUnreadable(new ConnectError("too large", Code.InvalidArgument))).toBe(false);
    expect(isUnreadable(new Error("network"))).toBe(false);
  });
});

describe("parseFaults", () => {
  it("counts BAD_ROW and UNKNOWN_VALUE alone (D-887)", () => {
    const rows = [UnresolvedReason.BAD_ROW, UnresolvedReason.UNKNOWN_VALUE, UnresolvedReason.UNKNOWN_CARD, UnresolvedReason.NON_ENGLISH, UnresolvedReason.NOT_PLAYABLE].map((reason) =>
      create(UnresolvedRowSchema, { reason }),
    );
    expect(parseFaults(rows)).toBe(2);
  });
});
