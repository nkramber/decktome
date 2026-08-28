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
    expect(screen.queryByLabelText("Your message")).not.toBeInTheDocument();
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
    // Nothing goes out until the submit.
    expect(chat).toHaveBeenCalledTimes(1);
    await user.click(screen.getByRole("button", { name: "Submit answers" }));

    const req = chat.mock.calls[1][0] as { sessionId: string; message: string; answers: { questionId: string; optionIndex?: number; text: string }[] };
    expect(req.sessionId).toBe("s1");
    expect(req.message).toBe("");
    expect(req.answers).toEqual([{ questionId: "q1", optionIndex: 2, text: "" }]);

    expect(await screen.findByRole("heading", { name: "Elves" })).toBeInTheDocument();
    expect(screen.queryByRole("group", { name: /Question:/ })).not.toBeInTheDocument();
    // The option the user clicked shows as their line.
    expect(within(screen.getByRole("list", { name: "Conversation" })).getAllByRole("listitem")[2]).toHaveTextContent("Modern");
    expect(screen.getByText("building the deck")).toBeInTheDocument();
    expect(screen.getByText("Here is your deck.")).toBeInTheDocument();
    expect(await screen.findByAltText("Llanowar Elves")).toBeInTheDocument();
  });

  it("one submit sends every answer, and waits until each question has one", async () => {
    const q2 = { id: "q2", slot: "power", text: "How strong?", options: [] };
    chat
      .mockReturnValueOnce(events([ev("sessionStarted", "s1"), ev("question", formatQuestion), ev("question", q2)]))
      .mockReturnValueOnce(events([ev("slots", {})]));
    renderAt("/session/new");
    const user = userEvent.setup();
    await user.type(await screen.findByLabelText("Your message"), "elves");
    await user.click(screen.getByRole("button", { name: "Send" }));
    const first = await screen.findByRole("group", { name: "Question: Which format?" });
    const second = screen.getByRole("group", { name: "Question: How strong?" });
    await user.type(within(first).getByLabelText("Or answer in your own words"), "Pauper");
    // One answer of two: the submit waits.
    expect(screen.getByRole("button", { name: "Submit answers" })).toBeDisabled();
    await user.type(within(second).getByLabelText("Or answer in your own words"), "bracket 2");
    await user.click(screen.getByRole("button", { name: "Submit answers" }));

    const req = chat.mock.calls[1][0] as { answers: { questionId: string; text: string }[] };
    expect(req.answers).toEqual([
      { questionId: "q1", optionIndex: undefined, text: "Pauper" },
      { questionId: "q2", optionIndex: undefined, text: "bracket 2" },
    ]);
    await waitFor(() => expect(screen.queryByText("The agent is working...")).not.toBeInTheDocument());
    expect(screen.queryByRole("group", { name: /Question:/ })).not.toBeInTheDocument();
    expect(await screen.findByLabelText("Your message")).toBeInTheDocument();
  });

  it("an option pick toggles, and a second pick replaces it", async () => {
    chat.mockReturnValueOnce(events([ev("sessionStarted", "s1"), ev("question", formatQuestion)]));
    renderAt("/session/new");
    const user = userEvent.setup();
    await user.type(await screen.findByLabelText("Your message"), "elves");
    await user.click(screen.getByRole("button", { name: "Send" }));
    const modern = await screen.findByRole("button", { name: "Modern" });
    await user.click(modern);
    expect(modern).toHaveAttribute("aria-pressed", "true");
    await user.click(screen.getByRole("button", { name: "Standard" }));
    expect(modern).toHaveAttribute("aria-pressed", "false");
    expect(screen.getByRole("button", { name: "Standard" })).toHaveAttribute("aria-pressed", "true");
    await user.click(screen.getByRole("button", { name: "Standard" }));
    expect(screen.getByRole("button", { name: "Submit answers" })).toBeDisabled();
  });

  it("sends every answer in one request when the next turn asks a different one", async () => {
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
    const colorsCard = screen.getByRole("group", { name: "Question: Any color preference?" });
    await user.type(within(colorsCard).getByLabelText("Or answer in your own words"), "green");
    await user.click(screen.getByRole("button", { name: "Submit answers" }));
    await screen.findByRole("group", { name: "Question: How strong?" });
    // A later turn that leaves a question open keeps it, and drops the answered ones.
    expect(screen.queryByRole("group", { name: "Question: Which format?" })).not.toBeInTheDocument();
    expect(screen.queryByRole("group", { name: "Question: Any color preference?" })).not.toBeInTheDocument();
    const req = chat.mock.calls[1][0] as { answers: { questionId: string; optionIndex?: number; text: string }[] };
    expect(req.answers).toEqual([
      { questionId: "q1", optionIndex: 2, text: "" },
      { questionId: "q2", optionIndex: undefined, text: "green" },
    ]);
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

  it("hides the message box the moment a send starts", async () => {
    let release = () => {};
    const gate = new Promise<void>((r) => (release = r));
    chat.mockImplementationOnce(async function* () {
      await gate;
      yield ev("sessionStarted", "s1");
      yield ev("slots", {});
    });
    renderAt("/session/new");
    const user = userEvent.setup();
    await user.type(await screen.findByLabelText("Your message"), "elves");
    await user.click(screen.getByRole("button", { name: "Send" }));
    expect(screen.queryByLabelText("Your message")).not.toBeInTheDocument();
    expect(screen.getByText("The agent is working...")).toBeInTheDocument();
    release();
    expect(await screen.findByLabelText("Your message")).toBeInTheDocument();
  });

  it("Enter sends, and a failed send gives the draft back", async () => {
    chat.mockImplementationOnce(() => {
      throw new Error("[unavailable] down");
    });
    renderAt("/session/new");
    const user = userEvent.setup();
    await user.type(await screen.findByLabelText("Your message"), "elves{enter}");
    expect(await screen.findByRole("alert")).toHaveTextContent("[unavailable] down (stream)");
    expect(screen.getByLabelText("Your message")).toHaveValue("elves");
    expect((chat.mock.calls[0][0] as { message: string }).message).toBe("elves");
  });

  it("offers to resume the stored session on /session/new", async () => {
    useAppStore.setState({ sessionId: "s9" });
    renderAt("/session/new");
    expect(await screen.findByTestId("resume-link")).toHaveAttribute("href", "/session/s9");
  });

  it("keeps a stored session's open question when a later turn asked another (reload merge)", async () => {
    getSession.mockResolvedValue({
      session: {
        id: "s1",
        collectionId: "",
        deckIds: [],
        turns: [
          { userMessage: "elves", agentMessage: "", questions: [formatQuestion, { id: "q2", slot: "colors", text: "Any color preference?", options: [] }], answers: [] },
          { userMessage: "", agentMessage: "", questions: [{ id: "q3", slot: "power", text: "How strong?", options: [] }], answers: [{ questionId: "q1", optionIndex: 2, text: "" }] },
        ],
      },
    });
    renderAt("/session/s1");
    await screen.findByRole("group", { name: "Question: How strong?" });
    expect(screen.getByRole("group", { name: "Question: Any color preference?" })).toBeInTheDocument();
    expect(screen.queryByRole("group", { name: "Question: Which format?" })).not.toBeInTheDocument();
    // The clicked option shows as the user's line.
    expect(within(screen.getByRole("list", { name: "Conversation" })).getAllByRole("listitem")[3]).toHaveTextContent("Modern");
  });

  it("a message after the deck streams the reply and the revised deck with its diff (PR-12B)", async () => {
    const revised = {
      ...deck,
      id: "d2",
      revisedFromDeckId: "d1",
      revisionNote: "I changed the count of Llanowar Elves: 4 to 2.",
      cards: [{ ...deck.cards[0], count: 2 }],
    };
    chat
      .mockReturnValueOnce(events([ev("sessionStarted", "s1"), ev("deck", deck), ev("usage", { calls: 2 })]))
      .mockReturnValueOnce(
        events([ev("status", "reading your request"), ev("status", "revising the deck"), ev("textDelta", "I changed the count of Llanowar Elves: 4 to 2."), ev("deck", revised), ev("usage", { calls: 4 })]),
      );
    renderAt("/session/new");
    const user = userEvent.setup();
    await user.type(await screen.findByLabelText("Your message"), "elves");
    await user.click(screen.getByRole("button", { name: "Send" }));
    await screen.findByRole("heading", { name: "Elves" });
    await user.type(await screen.findByLabelText("Your message"), "fewer elves{enter}");
    expect(await screen.findByTestId("revision-note")).toHaveTextContent("I changed the count of Llanowar Elves: 4 to 2.");
    expect(screen.getByTestId("revision-diff")).toHaveTextContent("Count of Llanowar Elves: 4 to 2");
    expect(screen.getByText("I changed the count of Llanowar Elves: 4 to 2.", { selector: "p.whitespace-pre-line" })).toBeInTheDocument();
    expect((chat.mock.calls[1][0] as { sessionId: string; message: string }).message).toBe("fewer elves");
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

    // The message box hides while a question waits (owner, 2026-08-28).
    expect(screen.queryByLabelText("Your message")).not.toBeInTheDocument();
    chat.mockReturnValue(events([ev("slots", {})]));
    const user = userEvent.setup();
    const card = screen.getByRole("group", { name: "Question: How strong?" });
    await user.type(within(card).getByLabelText("Or answer in your own words"), "bracket 2");
    await user.click(screen.getByRole("button", { name: "Submit answers" }));
    expect((chat.mock.calls[0][0] as { sessionId: string }).sessionId).toBe("s1");
    // A stored session carries its own collection, so the request sends none.
    expect((chat.mock.calls[0][0] as { collectionId: string }).collectionId).toBe("");
    // The box returns when no question is open, and the answer shows in the thread.
    expect(await screen.findByLabelText("Your message")).toBeInTheDocument();
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
