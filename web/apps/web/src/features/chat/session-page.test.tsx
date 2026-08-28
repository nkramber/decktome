import { CardRole } from "@mtg/api-client/mtg/v1/deck_pb";
import { PoolRule } from "@mtg/api-client/mtg/v1/session_pb";
import { screen, waitFor, within } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { axe } from "jest-axe";
import { beforeEach, describe, expect, it, vi } from "vitest";

import { useAppStore } from "../../lib/store";
import { fakeUser, state } from "../../test-auth-state";
import { renderAt } from "../../test-utils";

vi.mock("firebase/app");
vi.mock("firebase/auth");

const chat = vi.fn();
const getSession = vi.fn();
const getDeck = vi.fn();
const getCards = vi.fn();
vi.mock("../../lib/api", () => ({
  healthClient: { check: () => Promise.resolve({ status: "ok", version: "test", cardSnapshot: "none" }) },
  agentClient: {
    chat: (...args: unknown[]) => chat(...args),
    getSession: (...args: unknown[]) => getSession(...args),
  },
  deckClient: { getDeck: (...args: unknown[]) => getDeck(...args) },
  cardClient: { getCards: (...args: unknown[]) => getCards(...args) },
}));

type Ev = { event: { case: string; value: unknown } };
async function* events(list: Ev[]) {
  for (const e of list) {
    await Promise.resolve();
    yield e;
  }
}
const ev = (c: string, value: unknown): Ev => ({ event: { case: c, value } });

const formatQuestion = { id: "q1", slot: "format", text: "Which format?", options: ["Commander", "Standard", "Modern"] };
const deck = {
  id: "d1",
  name: "Elves",
  sessionId: "s1",
  cards: [{ oracleId: "o-elf", name: "Llanowar Elves", count: 4, role: CardRole.RAMP, owned: true, ownedCount: 4 }],
  sideboard: [],
  upgrades: [],
  commanderOracleIds: [],
  legalityAsOf: "2026-08-24",
  validation: { passed: true, findings: [] },
};

beforeEach(() => {
  state.user = fakeUser;
  localStorage.clear();
  useAppStore.setState({ collectionId: "", sessionId: "", poolMode: "any" });
  chat.mockReset();
  getSession.mockReset();
  getDeck.mockReset();
  getCards.mockReset();
  getCards.mockResolvedValue({ cards: [{ oracleId: "o-elf", name: "Llanowar Elves", cardTypes: ["Creature"], manaValue: 1, faces: [], defaultPrinting: { artist: "A", imageUris: { normal: "https://x/n.jpg", small: "https://x/s.jpg" } } }], missingOracleIds: [] });
});

describe("SessionPage", () => {
  it("sends the first message, takes the session id, and shows the question with its options", async () => {
    chat.mockReturnValue(events([ev("sessionStarted", "s1"), ev("question", formatQuestion), ev("slots", { poolRule: PoolRule.ANY_CARD }), ev("usage", { calls: 1, inputTokens: 100n, outputTokens: 20n, costUsd: 0.001, priced: true })]));
    const { router } = renderAt("/session/new");
    expect(await screen.findByTestId("session-id")).toHaveTextContent("No session yet.");
    const user = userEvent.setup();
    await user.type(screen.getByLabelText("Your message"), "Build me an elf deck");
    await user.click(screen.getByRole("button", { name: "Send" }));

    expect(await screen.findByRole("group", { name: "Question: Which format?" })).toBeInTheDocument();
    expect(router.state.location.pathname).toBe("/session/s1");
    expect(useAppStore.getState().sessionId).toBe("s1");
    expect(screen.getByTestId("session-id")).toHaveTextContent("Session id: s1");
    expect(screen.getByRole("button", { name: "Commander" })).toBeInTheDocument();
    expect(screen.getByTestId("usage")).toHaveTextContent("Session spend: 1 calls, 100 in, 20 out, $0.0010");
    expect(screen.getByTestId("pool-mode")).toHaveTextContent("Pool: any card (D-37).");
    const req = chat.mock.calls[0][0] as { sessionId: string; collectionId: string; message: string; answers: unknown[] };
    expect(req).toMatchObject({ sessionId: "", collectionId: "", message: "Build me an elf deck", answers: [] });
    // The conversation shows the user line and the asked line.
    const thread = screen.getByRole("list", { name: "Conversation" });
    expect(within(thread).getAllByRole("listitem")[0]).toHaveTextContent("Build me an elf deck");
  });

  it("an option click sends option_index, and the deck event opens the deck view", async () => {
    chat
      .mockReturnValueOnce(events([ev("sessionStarted", "s1"), ev("question", formatQuestion)]))
      .mockReturnValueOnce(events([ev("slots", { poolRule: PoolRule.ANY_CARD }), ev("status", "building the deck"), ev("textDelta", "Here is "), ev("textDelta", "your deck."), ev("deck", deck), ev("usage", { calls: 3 })]));
    renderAt("/session/new");
    const user = userEvent.setup();
    await user.type(await screen.findByLabelText("Your message"), "elves");
    await user.click(screen.getByRole("button", { name: "Send" }));
    await user.click(await screen.findByRole("button", { name: "Modern" }));

    const req = chat.mock.calls[1][0] as { sessionId: string; message: string; answers: { questionId: string; optionIndex?: number; text: string }[] };
    expect(req.sessionId).toBe("s1");
    expect(req.message).toBe("");
    expect(req.answers).toEqual([{ questionId: "q1", optionIndex: 2, text: "" }]);

    expect(await screen.findByRole("heading", { name: "Elves" })).toBeInTheDocument();
    expect(screen.queryByRole("group", { name: /Question:/ })).not.toBeInTheDocument();
    expect(screen.getByText("building the deck")).toBeInTheDocument();
    expect(screen.getByText("Here is your deck.")).toBeInTheDocument();
    expect(await screen.findByAltText("Llanowar Elves")).toBeInTheDocument();
  });

  it("free text sends text, and a turn with no question keeps the other questions open (D-237)", async () => {
    const q2 = { id: "q2", slot: "power", text: "How strong?", options: [] };
    chat
      .mockReturnValueOnce(events([ev("sessionStarted", "s1"), ev("question", formatQuestion), ev("question", q2)]))
      .mockReturnValueOnce(events([ev("slots", {})]));
    renderAt("/session/new");
    const user = userEvent.setup();
    await user.type(await screen.findByLabelText("Your message"), "elves");
    await user.click(screen.getByRole("button", { name: "Send" }));
    const first = await screen.findByRole("group", { name: "Question: Which format?" });
    await user.type(within(first).getByLabelText("Or answer in your own words"), "Pauper");
    await user.click(within(first).getByRole("button", { name: "Answer" }));

    const req = chat.mock.calls[1][0] as { answers: { questionId: string; text: string }[] };
    expect(req.answers).toEqual([{ questionId: "q1", text: "Pauper" }]);
    await waitFor(() => expect(screen.queryByText("The agent is working...")).not.toBeInTheDocument());
    expect(screen.queryByRole("group", { name: "Question: Which format?" })).not.toBeInTheDocument();
    expect(screen.getByRole("group", { name: "Question: How strong?" })).toBeInTheDocument();
  });

  it("keeps an open question when the next turn asks a different one (the browser gate fault of 2026-08-28)", async () => {
    const colors = { id: "q2", slot: "colors", text: "Any color preference?", options: [] };
    const power = { id: "q3", slot: "power", text: "How strong?", options: ["Casual", "FNM"] };
    chat
      .mockReturnValueOnce(events([ev("sessionStarted", "s1"), ev("question", formatQuestion), ev("question", colors)]))
      .mockReturnValueOnce(events([ev("question", power), ev("slots", {})]));
    renderAt("/session/new");
    const user = userEvent.setup();
    await user.type(await screen.findByLabelText("Your message"), "angels");
    await user.click(screen.getByRole("button", { name: "Send" }));
    await user.click(await screen.findByRole("button", { name: "Modern" }));
    await screen.findByRole("group", { name: "Question: How strong?" });
    expect(screen.getByRole("group", { name: "Question: Any color preference?" })).toBeInTheDocument();
    expect(screen.queryByRole("group", { name: "Question: Which format?" })).not.toBeInTheDocument();
  });

  it("shows a failure event and a stream error as alerts", async () => {
    chat.mockReturnValueOnce(events([ev("sessionStarted", "s1"), ev("failure", { code: "llm_unavailable", message: "The model did not answer.", retryable: true })]));
    renderAt("/session/new");
    const user = userEvent.setup();
    await user.type(await screen.findByLabelText("Your message"), "elves");
    await user.click(screen.getByRole("button", { name: "Send" }));
    expect(await screen.findByRole("alert")).toHaveTextContent("The model did not answer. (llm_unavailable) You can try again.");

    chat.mockImplementationOnce(() => {
      throw new Error("[resource_exhausted] busy");
    });
    await user.type(screen.getByLabelText("Your message"), "again");
    await user.click(screen.getByRole("button", { name: "Send" }));
    expect((await screen.findAllByRole("alert"))[1]).toHaveTextContent("[resource_exhausted] busy (stream)");
  });

  it("sends the collection id in owned mode on the first message only", async () => {
    useAppStore.setState({ collectionId: "c1", poolMode: "owned" });
    chat.mockReturnValue(events([ev("sessionStarted", "s1")]));
    renderAt("/session/new");
    const user = userEvent.setup();
    const box = await screen.findByLabelText("Use only cards in my collection");
    expect(box).toBeChecked();
    expect(screen.getByTestId("pool-mode")).toHaveTextContent("Pool: your collection (c1). The agent asks how strict.");
    await user.click(box);
    expect(useAppStore.getState().poolMode).toBe("any");
    await user.click(box);
    await user.type(screen.getByLabelText("Your message"), "elves");
    await user.click(screen.getByRole("button", { name: "Send" }));
    await screen.findByText("Session id: s1");
    expect((chat.mock.calls[0][0] as { collectionId: string }).collectionId).toBe("c1");
    expect(screen.queryByLabelText("Use only cards in my collection")).not.toBeInTheDocument();
  });

  it("refuses a message over the 8 KiB cap", async () => {
    renderAt("/session/new");
    const user = userEvent.setup();
    const box = await screen.findByLabelText("Your message");
    await user.click(box);
    await user.paste("x".repeat(8193));
    expect(screen.getByRole("alert")).toHaveTextContent("over the 8192 byte cap");
    expect(screen.getByRole("button", { name: "Send" })).toBeDisabled();
    expect(chat).not.toHaveBeenCalled();
  });

  it("rebuilds the thread and the deck from a stored session on reload", async () => {
    getSession.mockResolvedValue({
      session: {
        id: "s1",
        collectionId: "c1",
        deckIds: ["d0", "d1"],
        slots: { poolRule: PoolRule.OWNED_ONLY },
        usage: { calls: 4, inputTokens: 10n, outputTokens: 5n, priced: false },
        turns: [
          { userMessage: "elves", agentMessage: "", questions: [formatQuestion, { id: "q2", slot: "power", text: "How strong?", options: [] }], answers: [] },
          { userMessage: "", agentMessage: "", questions: [], answers: [{ questionId: "q1", optionIndex: 2, text: "" }] },
        ],
      },
    });
    getDeck.mockResolvedValue({ deck });
    renderAt("/session/s1");
    expect(await screen.findByRole("heading", { name: "Elves" })).toBeInTheDocument();
    expect(getDeck).toHaveBeenCalledWith({ deckId: "d1" });
    expect(screen.getByTestId("session-id")).toHaveTextContent("Session id: s1");
    expect(screen.getByTestId("pool-mode")).toHaveTextContent("Pool: only cards in your collection.");
    expect(screen.getByTestId("usage")).toHaveTextContent("cost unknown");
    expect(screen.queryByRole("group", { name: "Question: Which format?" })).not.toBeInTheDocument();
    expect(screen.getByRole("group", { name: "Question: How strong?" })).toBeInTheDocument();

    chat.mockReturnValue(events([ev("slots", {})]));
    await userEvent.setup().type(screen.getByLabelText("Your message"), "bracket 2{enter}");
    await userEvent.setup().click(screen.getByRole("button", { name: "Send" }));
    expect((chat.mock.calls[0][0] as { sessionId: string; collectionId: string }).sessionId).toBe("s1");
    expect((chat.mock.calls[0][0] as { collectionId: string }).collectionId).toBe("c1");
  });

  it("reports a session that does not load", async () => {
    getSession.mockRejectedValue(new Error("[not_found] no session"));
    renderAt("/session/nope");
    expect(await screen.findByRole("alert")).toHaveTextContent("Could not load the session: [not_found] no session");
  });

  it("has no axe violations with a question and a deck on screen", async () => {
    chat.mockReturnValue(events([ev("sessionStarted", "s1"), ev("question", formatQuestion), ev("deck", deck)]));
    const { container } = renderAt("/session/new");
    const user = userEvent.setup();
    await user.type(await screen.findByLabelText("Your message"), "elves");
    await user.click(screen.getByRole("button", { name: "Send" }));
    await screen.findByAltText("Llanowar Elves");
    expect(await axe(container)).toHaveNoViolations();
  });
});
