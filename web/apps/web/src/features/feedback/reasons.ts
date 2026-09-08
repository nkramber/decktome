import { FeedbackKind } from "@mtg/api-client/mtg/v1/feedback_service_pb";

// The prefilled reasons of a thumbs down, per kind (PR-27, section 2.4
// of the design note). The key is what the API stores and the harvest
// of PR-28 reads. go/internal/feedback holds the same keys, and a test
// on each side pins the list, so a drift fails a test before it refuses
// a verdict.
export type Reason = { key: string; label: string };

const byKind: Record<FeedbackKind, Reason[]> = {
  [FeedbackKind.UNSPECIFIED]: [],
  [FeedbackKind.QUESTION]: [
    { key: "already_answered", label: "You already had this answer." },
    { key: "not_applicable", label: "The question does not apply to my request." },
    { key: "bad_options", label: "The options are wrong, or one is missing." },
    { key: "unclear", label: "The wording is unclear." },
  ],
  [FeedbackKind.SUMMARY]: [
    { key: "false_claim", label: "It claims something the deck does not do." },
    { key: "misses_plan", label: "It misses the plan I asked for." },
    { key: "too_long_or_vague", label: "It is too long or too vague." },
  ],
  [FeedbackKind.CARD]: [
    { key: "off_theme", label: "This card does not fit the theme." },
    { key: "illegal", label: "This card breaks the format or my house rules." },
    { key: "unwanted_buy", label: "I do not own it and did not want to buy." },
    { key: "wrong_printing", label: "This is the wrong printing." },
    { key: "wrong_power", label: "It is too weak or too strong for the bracket." },
  ],
  [FeedbackKind.DECK]: [
    { key: "off_spec", label: "It does not match what I asked for." },
    { key: "bad_mana", label: "The mana base is off." },
    { key: "too_little_interaction", label: "Too little interaction, ramp, or draw." },
    { key: "wrong_power", label: "Too weak or too strong for the bracket." },
    { key: "too_many_to_buy", label: "Too many cards to buy." },
  ],
  // The conversation as a whole (D-594). A chat that stops between
  // questions belongs to no question, so the thumbs of a question reach
  // it from nowhere.
  [FeedbackKind.CHAT]: [
    { key: "stuck", label: "The chat stopped and asked nothing more." },
    { key: "ignored_request", label: "It ignored part of what I asked for." },
    { key: "wrong_questions", label: "It asked the wrong questions." },
    { key: "no_deck", label: "It never built a deck." },
    { key: "error", label: "Something failed or showed an error." },
  ],
};

export function reasonsOf(kind: FeedbackKind): Reason[] {
  return byKind[kind] ?? [];
}
