import { FeedbackKind } from "@mtg/api-client/mtg/v1/feedback_service_pb";
import { describe, expect, it } from "vitest";

import { reasonsOf } from "./reasons";

// The keys are the ones go/internal/feedback holds. The API refuses a key
// it does not know, so this list and the Go list move together.
describe("reasonsOf", () => {
  it("pins the keys of every kind, in the order of the design note (PR-27)", () => {
    expect(reasonsOf(FeedbackKind.QUESTION).map((r) => r.key)).toEqual(["already_answered", "not_applicable", "bad_options", "unclear"]);
    expect(reasonsOf(FeedbackKind.SUMMARY).map((r) => r.key)).toEqual(["false_claim", "misses_plan", "too_long_or_vague"]);
    expect(reasonsOf(FeedbackKind.CARD).map((r) => r.key)).toEqual(["off_theme", "illegal", "unwanted_buy", "wrong_printing", "wrong_power"]);
    expect(reasonsOf(FeedbackKind.DECK).map((r) => r.key)).toEqual(["off_spec", "bad_mana", "too_little_interaction", "wrong_power", "too_many_to_buy"]);
    expect(reasonsOf(FeedbackKind.UNSPECIFIED)).toEqual([]);
  });

  it("gives every reason a sentence", () => {
    for (const kind of [FeedbackKind.QUESTION, FeedbackKind.SUMMARY, FeedbackKind.CARD, FeedbackKind.DECK]) {
      for (const r of reasonsOf(kind)) expect(r.label).toMatch(/^[A-Z].*\.$/);
    }
  });
});
