import { Code, ConnectError } from "@connectrpc/connect";
import type { Deck } from "@mtg/api-client/mtg/v1/deck_pb";
import type { Question, Session } from "@mtg/api-client/mtg/v1/session_pb";
import { describe, expect, it, vi } from "vitest";

import { BuildPhase } from "@mtg/api-client/mtg/v1/agent_service_pb";

import { SessionStatus, SlotState } from "@mtg/api-client/mtg/v1/session_pb";

import { answerLabel, answerText, byteLength, codeName, emptyState, fromSession, mergeOpen, stillAsked, streamFailure } from "./use-chat";

vi.mock("../../lib/api", () => ({ agentClient: {} }));

const q = (id: string, slot: string, options: string[] = []) => ({ id, slot, text: `${slot}?`, options, optionOracleIds: [] }) as unknown as Question;

describe("use-chat helpers", () => {
  it("mergeOpen keeps the unanswered, replaces a repeated slot, and adds the new (D-278)", () => {
    const open = [q("q1", "format"), q("q2", "colors")];
    const merged = mergeOpen(open, [q("q3", "format"), q("q4", "power")]);
    expect(merged.map((x) => x.id)).toEqual(["q2", "q3", "q4"]);
  });

  it("answerText is the text, else the option, else empty", () => {
    const question = q("q1", "format", ["Commander", "Modern"]);
    expect(answerText({ questionId: "q1", text: "Pauper" } as never, question)).toBe("Pauper");
    expect(answerText({ questionId: "q1", optionIndex: 1, text: "" } as never, question)).toBe("Modern");
    expect(answerText({ questionId: "q1", optionIndex: 9, text: "" } as never, question)).toBe("");
    expect(answerText({ questionId: "q1", optionIndex: 1, text: "" } as never, undefined)).toBe("");
  });

  // F-62: a decline carries no text and no option, and the thread showed
  // nothing at all for it before.
  it("answerLabel reads a decline as the choice it is", () => {
    const question = q("q1", "format", ["Pauper", "Modern"]);
    const budget = q("q2", "budget");
    expect(answerLabel({ questionId: "q1", declined: true, text: "" } as never, question)).toBe("You decide");
    expect(answerLabel({ questionId: "q2", declined: true, text: "" } as never, budget)).toBe("No budget");
    expect(answerLabel({ questionId: "q1", optionIndex: 1, text: "" } as never, question)).toBe("Modern");
    expect(answerLabel({ questionId: "q1", text: "Pauper" } as never, question)).toBe("Pauper");
    expect(answerLabel({ questionId: "q1", declined: true, text: "" } as never, undefined)).toBe("You decide");
  });

  it("fromSession shows a declined answer under its question", () => {
    const session = {
      id: "s1",
      turns: [
        { userMessage: "elves", agentMessage: "", questions: [q("q1", "format", ["A", "B"])], answers: [] },
        { userMessage: "", agentMessage: "", questions: [], answers: [{ questionId: "q1", declined: true, text: "" }] },
      ],
      deckIds: [],
    } as unknown as Session;
    const state = fromSession(session);
    const asked = state.thread.find((t) => t.kind === "question");
    expect(asked?.kind === "question" && asked.answer).toBe("You decide");
    expect(state.thread.map((t) => t.kind)).toEqual(["user", "question"]);
    expect(state.openQuestions).toHaveLength(0);
  });

  it("byteLength counts UTF-8 bytes", () => {
    expect(byteLength("abc")).toBe(3);
    expect(byteLength("é")).toBe(2);
    expect(byteLength("")).toBe(0);
  });

  it("fromSession rebuilds the thread, drops answered questions, and orders the deck line by time", () => {
    const session = {
      id: "s1",
      turns: [
        { userMessage: "elves", agentMessage: "", questions: [q("q1", "format", ["A", "B"]), q("q2", "power")], answers: [], at: { seconds: 10n, nanos: 0 } },
        { userMessage: "", agentMessage: "ok", questions: [], answers: [{ questionId: "q1", optionIndex: 1, text: "" }], at: { seconds: 30n, nanos: 0 } },
      ],
      deckIds: ["d0", "d1"],
    } as unknown as Session;
    const deck = { id: "d1", name: "Elves", revisedFromDeckId: "d0", createdAt: { seconds: 20n, nanos: 0 } } as unknown as Deck;
    const base = { id: "d0", name: "Old" } as unknown as Deck;
    const state = fromSession(session, deck, base);
    expect(state.sessionId).toBe("s1");
    // The answer reads under its own question, and the turn that carried
    // it holds no message, so it adds no user block (F-62).
    expect(state.thread.map((t) => t.kind)).toEqual(["user", "question", "question", "deck", "agent"]);
    const q1 = state.thread.find((t) => t.kind === "question" && t.question.id === "q1");
    expect(q1?.kind === "question" && q1.answer).toBe("B");
    const q2 = state.thread.find((t) => t.kind === "question" && t.question.id === "q2");
    expect(q2?.kind === "question" && q2.answer).toBe(undefined);
    expect(state.thread.map((t) => t.id)).toEqual([...new Set(state.thread.map((t) => t.id))]);
    expect(state.openQuestions.map((x) => x.id)).toEqual(["q2"]);
    expect(state.deck).toBe(deck);
    expect(state.baseDeck).toBe(base);
    expect(state.busy).toBe(false);
  });

  it("fromSession appends the deck line when a time is missing, and drops a base that is not the revised one", () => {
    const session = { id: "s1", turns: [{ userMessage: "elves", agentMessage: "", questions: [], answers: [] }], deckIds: ["d1"] } as unknown as Session;
    const deck = { id: "d1", revisedFromDeckId: "d0" } as unknown as Deck;
    const state = fromSession(session, deck, { id: "d9" } as unknown as Deck);
    expect(state.thread.map((t) => t.kind)).toEqual(["user", "deck"]);
    expect(state.baseDeck).toBeUndefined();
    expect(fromSession(session).thread.map((t) => t.kind)).toEqual(["user"]);
  });

  it("streamFailure maps the Connect codes", () => {
    expect(codeName(Code.ResourceExhausted)).toBe("resource_exhausted");
    expect(streamFailure(new ConnectError("slow", Code.DeadlineExceeded))).toMatchObject({ code: "deadline_exceeded", message: "slow", retryable: true });
    expect(streamFailure(new ConnectError("down", Code.Unavailable)).retryable).toBe(true);
    expect(streamFailure(new ConnectError("no", Code.PermissionDenied))).toMatchObject({ code: "permission_denied", retryable: false });
    expect(streamFailure(new ConnectError("a build is in progress", Code.Aborted))).toMatchObject({ code: "aborted", retryable: true });
    expect(streamFailure(new Error("boom"))).toEqual({ code: "stream", message: "boom", retryable: false });
  });
});

// The phase of a turn rides beside the status lines (D-435).
describe("the build phase", () => {
  it("starts unset, and a repair marks the state", () => {
    expect(emptyState.phase).toBe(BuildPhase.UNSPECIFIED);
    expect(emptyState.repaired).toBe(false);
    const s = fromSession({ id: "s", turns: [], deckIds: [] } as never);
    expect(s.phase).toBe(BuildPhase.UNSPECIFIED);
  });
});

// A question in a stored session closes with its slot, and every
// question closes once the session built a deck (D-446).
describe("stillAsked", () => {
  const q = { id: "q6", slot: "commander", text: "Which commander?", options: [] } as unknown as Question;
  const session = (status: SessionStatus, state?: SlotState) =>
    ({ id: "s", status, slots: { slotStates: state === undefined ? {} : { commander: state } }, turns: [], deckIds: [] }) as unknown as Session;

  it("keeps a question whose slot is still asked", () => {
    expect(stillAsked(session(SessionStatus.ASKING, SlotState.ASKED), q)).toBe(true);
    expect(stillAsked(session(SessionStatus.ASKING), q)).toBe(true);
  });

  it("closes a question whose slot is filled or skipped", () => {
    expect(stillAsked(session(SessionStatus.ASKING, SlotState.FILLED), q)).toBe(false);
    expect(stillAsked(session(SessionStatus.ASKING, SlotState.SKIPPED), q)).toBe(false);
  });

  it("closes every question of a built session, even with no answer recorded", () => {
    const built = { ...session(SessionStatus.BUILT, SlotState.ASKED), turns: [{ userMessage: "Human deck", questions: [q], answers: [], at: undefined }] } as unknown as Session;
    expect(fromSession(built).openQuestions).toEqual([]);
  });
});
