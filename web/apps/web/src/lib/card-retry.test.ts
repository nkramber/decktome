import { Code, ConnectError } from "@connectrpc/connect";
import { describe, expect, it } from "vitest";

import { cardRetry, cardRetryDelay, cardRetryMax, indexLoading } from "./card-retry";

describe("cardRetry", () => {
  it("retries while the card database is Unavailable", () => {
    const err = new ConnectError("card database not loaded yet", Code.Unavailable);
    expect(cardRetry(0, err)).toBe(true);
    expect(cardRetry(cardRetryMax - 1, err)).toBe(true);
  });

  it("gives up at the end of the cold start window", () => {
    const err = new ConnectError("card database not loaded yet", Code.Unavailable);
    expect(cardRetry(cardRetryMax, err)).toBe(false);
  });

  it("takes any other code at once", () => {
    for (const c of [Code.Unauthenticated, Code.InvalidArgument, Code.NotFound]) {
      expect(cardRetry(0, new ConnectError("no", c))).toBe(false);
    }
    expect(cardRetry(0, new Error("network"))).toBe(false);
  });

  it("waits under 90 seconds over every retry", () => {
    let total = 0;
    for (let i = 0; i < cardRetryMax; i++) total += cardRetryDelay(i);
    expect(total).toBeLessThanOrEqual(90_000);
    expect(total).toBeGreaterThanOrEqual(60_000);
  });

  it("backs off to 8 seconds and stays there", () => {
    expect(cardRetryDelay(0)).toBe(1000);
    expect(cardRetryDelay(3)).toBe(8000);
    expect(cardRetryDelay(20)).toBe(8000);
  });
});

// D-954: the chat sends a turn again on the refusal of a read that met no
// card index, and on no other Unavailable. A turn that conflicts with
// another turn reads Unavailable too, after a paid call.
describe("indexLoading", () => {
  it("reads the refusal header of a read with no card index", () => {
    const loading = new ConnectError("the card database is not loaded yet, so the chat waits", Code.Unavailable, { "deck-tome-refusal": "index-loading" });
    expect(indexLoading(loading)).toBe(true);
  });

  it("takes no plain Unavailable, and no other refusal", () => {
    expect(indexLoading(new ConnectError("the session is busy with another turn, send the message again", Code.Unavailable))).toBe(false);
    expect(indexLoading(new ConnectError("no", Code.PermissionDenied, { "deck-tome-refusal": "index-loading" }))).toBe(false);
    expect(indexLoading(new ConnectError("no", Code.Unavailable, { "deck-tome-refusal": "not-invited" }))).toBe(false);
    expect(indexLoading(new Error("network"))).toBe(false);
  });
});
