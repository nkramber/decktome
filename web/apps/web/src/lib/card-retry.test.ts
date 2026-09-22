import { Code, ConnectError } from "@connectrpc/connect";
import { describe, expect, it } from "vitest";

import { cardRetry, cardRetryDelay, cardRetryMax } from "./card-retry";

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
