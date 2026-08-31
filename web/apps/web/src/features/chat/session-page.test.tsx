import { Code, ConnectError } from "@connectrpc/connect";
import { CardRole } from "@mtg/api-client/mtg/v1/deck_pb";
import { PoolRule } from "@mtg/api-client/mtg/v1/session_pb";
import { act, fireEvent, screen, waitFor, within } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { axe } from "jest-axe";
import { beforeEach, describe, expect, it, vi } from "vitest";

import { useAppStore } from "../../lib/store";
import { fakeUser, state } from "../../test-auth-state";
import { renderAt } from "../../test-utils";
import { poolLabel, pruneDrafts } from "./session-page";

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
  deckClient: { getDeck: (...args: unknown[]) => getDeck(...args), exportDeck: vi.fn() },
  cardClient: { getCards: (...args: unknown[]) => getCards(...args) },
  collectionClient: { listCollections: () => Promise.resolve({ collections: [{ id: "c1", name: "binder-july.csv", cardCount: 4317 }] }) },
}));

type Ev = { event: { case: string; value: unknown } };
async function* events(list: Ev[]) {
  for (const e of list) {
    await Promise.resolve();
    yield e;
  }
}
const ev = (c: string, value: unknown): Ev => ({ event: { case: c, value } });

const formatQuestion = { id: "q1", slot: "format", text: "Which format?", options: ["Commander", "Standard", "Modern"], optionOracleIds: [] };
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
    const { router } = await renderAt("/session/new");
    expect(await screen.findByTestId("session-id")).toHaveTextContent("No session yet.");
    const user = userEvent.setup();
    await user.type(screen.getByLabelText("Your message"), "Build me an elf deck");
    await user.click(screen.getByRole("button", { name: "Send" }));

    expect(await screen.findByRole("group", { name: "Question: Which format?" })).toBeInTheDocument();
    expect(router.state.location.pathname).toBe("/session/s1");
    expect(useAppStore.getState().sessionId).toBe("s1");
    expect(screen.getByTestId("session-id")).toHaveTextContent("Session id: s1");
    expect(screen.getByRole("button", { name: "Commander" })).toBeInTheDocument();
    expect(screen.getByTestId("usage")).toHaveTextContent("Session spend: 1 calls, 100 in, 20 out, $0.0010.");
    expect(screen.getByTestId("pool-mode")).toHaveTextContent("Pool: any card.");
    const req = chat.mock.calls[0][0] as { sessionId: string; collectionId: string; message: string; answers: unknown[] };
    expect(req).toMatchObject({ sessionId: "", collectionId: "", message: "Build me an elf deck", answers: [] });
    // The conversation shows the user line and the asked line.
    const thread = screen.getByRole("list", { name: "Conversation" });
    expect(within(thread).getAllByRole("listitem")[0]).toHaveTextContent("Build me an elf deck");
    expect(screen.queryByLabelText("Your message")).not.toBeInTheDocument();
  });

  it("an option click sends option_index, and a built deck hands over to its address (D-335)", async () => {
    chat
      .mockReturnValueOnce(events([ev("sessionStarted", "s1"), ev("question", formatQuestion)]))
      .mockReturnValueOnce(events([ev("slots", { poolRule: PoolRule.ANY_CARD }), ev("status", "building the deck"), ev("textDelta", "Here is "), ev("textDelta", "your deck."), ev("deck", deck), ev("usage", { calls: 3 })]));
    const { router } = await renderAt("/session/new");
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

    // A built deck lives at its own address, and the session hands the
    // reader over to it (D-335).
    await waitFor(() => expect(router.state.location.pathname).toBe("/decks/d1"));
  });

  it("one submit sends every answer, and waits until each question has one", async () => {
    const q2 = { id: "q2", slot: "power", text: "How strong?", options: [], optionOracleIds: [] };
    chat
      .mockReturnValueOnce(events([ev("sessionStarted", "s1"), ev("question", formatQuestion), ev("question", q2)]))
      .mockReturnValueOnce(events([ev("slots", {})]));
    await renderAt("/session/new");
    const user = userEvent.setup();
    await user.type(await screen.findByLabelText("Your message"), "elves");
    await user.click(screen.getByRole("button", { name: "Send" }));
    const first = await screen.findByRole("group", { name: "Question: Which format?" });
    const second = screen.getByRole("group", { name: "Question: How strong?" });
    await user.type(within(first).getByRole("textbox"), "Pauper");
    // One answer of two: the submit waits.
    expect(screen.getByRole("button", { name: "Submit answers" })).toBeDisabled();
    await user.type(within(second).getByRole("textbox"), "bracket 2");
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

  it("a closed question offers no free-text field (D-295)", async () => {
    chat.mockReturnValueOnce(events([ev("sessionStarted", "s1"), ev("question", { ...formatQuestion, closed: true })]));
    await renderAt("/session/new");
    const user = userEvent.setup();
    await user.type(await screen.findByLabelText("Your message"), "elves");
    await user.click(screen.getByRole("button", { name: "Send" }));
    const card = await screen.findByRole("group", { name: "Question: Which format?" });
    expect(within(card).queryByRole("textbox")).not.toBeInTheDocument();
    await user.click(within(card).getByRole("button", { name: "Modern" }));
    expect(screen.getByRole("button", { name: "Submit answers" })).toBeEnabled();
  });

  it("an option pick toggles, and a second pick replaces it", async () => {
    chat.mockReturnValueOnce(events([ev("sessionStarted", "s1"), ev("question", formatQuestion)]));
    await renderAt("/session/new");
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
    const colors = { id: "q2", slot: "colors", text: "Any color preference?", options: [], optionOracleIds: [] };
    const power = { id: "q3", slot: "power", text: "How strong?", options: ["Casual", "FNM"], optionOracleIds: [] };
    chat
      .mockReturnValueOnce(events([ev("sessionStarted", "s1"), ev("question", formatQuestion), ev("question", colors)]))
      .mockReturnValueOnce(events([ev("question", power), ev("slots", {})]));
    await renderAt("/session/new");
    const user = userEvent.setup();
    await user.type(await screen.findByLabelText("Your message"), "angels");
    await user.click(screen.getByRole("button", { name: "Send" }));
    await user.click(await screen.findByRole("button", { name: "Modern" }));
    const colorsCard = screen.getByRole("group", { name: "Question: Any color preference?" });
    await user.type(within(colorsCard).getByRole("textbox"), "green");
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
    await renderAt("/session/new");
    const user = userEvent.setup();
    await user.type(await screen.findByLabelText("Your message"), "elves");
    await user.click(screen.getByRole("button", { name: "Send" }));
    expect(await screen.findByRole("alert")).toHaveTextContent("The model did not answer. (llm_unavailable) You can try again.");

    chat.mockImplementationOnce(() => {
      throw new ConnectError("busy", Code.ResourceExhausted);
    });
    await user.type(screen.getByLabelText("Your message"), "again");
    await user.click(screen.getByRole("button", { name: "Send" }));
    expect((await screen.findAllByRole("alert"))[1]).toHaveTextContent("busy (resource_exhausted) You can try again.");
  });

  it("sends the collection id in owned mode on the first message only", async () => {
    useAppStore.setState({ collectionId: "c1", poolMode: "owned" });
    chat.mockReturnValue(events([ev("sessionStarted", "s1")]));
    await renderAt("/session/new");
    const user = userEvent.setup();
    const box = await screen.findByLabelText("Use only cards in my collection");
    expect(box).toBeChecked();
    // The line names the collection, never its id.
    expect(await screen.findByTestId("pool-mode")).toHaveTextContent("Pool: binder-july.csv. The agent asks how strict.");
    await user.click(box);
    expect(useAppStore.getState().poolMode).toBe("any");
    await user.click(box);
    await user.type(screen.getByLabelText("Your message"), "elves");
    await user.click(screen.getByRole("button", { name: "Send" }));
    await screen.findByText("Session id: s1");
    expect((chat.mock.calls[0][0] as { collectionId: string }).collectionId).toBe("c1");
    expect(screen.queryByLabelText("Use only cards in my collection")).not.toBeInTheDocument();
    await waitFor(() => expect(screen.getByLabelText("Your message")).toBeEnabled());
    await user.type(screen.getByLabelText("Your message"), "more{enter}");
    await waitFor(() => expect(chat).toHaveBeenCalledTimes(2));
    expect((chat.mock.calls[1][0] as { sessionId: string; collectionId: string }).collectionId).toBe("");
    expect((chat.mock.calls[1][0] as { sessionId: string }).sessionId).toBe("s1");
  });

  it("takes the message box away while the agent works, and focuses it again at the end", async () => {
    let release = () => {};
    const gate = new Promise<void>((r) => (release = r));
    chat.mockImplementationOnce(async function* () {
      await gate;
      yield ev("sessionStarted", "s1");
      yield ev("slots", {});
    });
    await renderAt("/session/new");
    const user = userEvent.setup();
    await user.type(await screen.findByLabelText("Your message"), "elves");
    await user.click(screen.getByRole("button", { name: "Send" }));
    // A box that stays and does nothing reads as a dead control, so the
    // working row and its Stop take its place (PR-16B).
    expect(screen.queryByLabelText("Your message")).not.toBeInTheDocument();
    expect(screen.queryByRole("button", { name: "Send" })).not.toBeInTheDocument();
    expect(screen.getByText("The agent is working...")).toBeInTheDocument();
    release();
    await waitFor(() => expect(screen.getByLabelText("Your message")).toBeInTheDocument());
    expect(screen.getByLabelText("Your message")).toHaveFocus();
  });

  it("Stop aborts the stream, clears busy, and gives the draft back", async () => {
    chat.mockImplementationOnce(async function* (_req: unknown, opts: { signal: AbortSignal }) {
      yield ev("sessionStarted", "s1");
      await new Promise((_, reject) => opts.signal.addEventListener("abort", () => reject(new ConnectError("canceled", Code.Canceled))));
    });
    await renderAt("/session/new");
    const user = userEvent.setup();
    await user.type(await screen.findByLabelText("Your message"), "elves");
    await user.click(screen.getByRole("button", { name: "Send" }));
    await user.click(await screen.findByRole("button", { name: "Stop" }));
    await waitFor(() => expect(screen.getByLabelText("Your message")).toBeEnabled());
    expect(screen.queryByText("The agent is working...")).not.toBeInTheDocument();
    expect(screen.queryByRole("alert")).not.toBeInTheDocument();
    expect(screen.getByLabelText("Your message")).toHaveValue("elves");
  });

  it("the unmount aborts the stream", async () => {
    const abort = vi.fn();
    chat.mockImplementationOnce(async function* (_req: unknown, opts: { signal: AbortSignal }) {
      opts.signal.addEventListener("abort", abort);
      yield ev("sessionStarted", "s1");
      await new Promise(() => {});
    });
    const { unmount } = await renderAt("/session/new");
    const user = userEvent.setup();
    await user.type(await screen.findByLabelText("Your message"), "elves");
    await user.click(screen.getByRole("button", { name: "Send" }));
    await screen.findByText("Session id: s1");
    unmount();
    expect(abort).toHaveBeenCalled();
  });

  it("a thrown send gives the answered questions and their drafts back", async () => {
    chat
      .mockReturnValueOnce(events([ev("sessionStarted", "s1"), ev("question", formatQuestion)]))
      .mockImplementationOnce(() => {
        throw new ConnectError("down", Code.Unavailable);
      });
    await renderAt("/session/new");
    const user = userEvent.setup();
    await user.type(await screen.findByLabelText("Your message"), "elves");
    await user.click(screen.getByRole("button", { name: "Send" }));
    await user.click(await screen.findByRole("button", { name: "Modern" }));
    await user.click(screen.getByRole("button", { name: "Submit answers" }));
    expect(await screen.findByRole("alert")).toHaveTextContent("down (unavailable)");
    const card = screen.getByRole("group", { name: "Question: Which format?" });
    expect(within(card).getByRole("button", { name: "Modern" })).toHaveAttribute("aria-pressed", "true");
    expect(screen.getByRole("button", { name: "Submit answers" })).toBeEnabled();
  });

  it("an Aborted stream names the build in progress and stays retryable (D-303)", async () => {
    chat.mockImplementationOnce(() => {
      throw new ConnectError("a build is in progress", Code.Aborted);
    });
    await renderAt("/session/new");
    const user = userEvent.setup();
    await user.type(await screen.findByLabelText("Your message"), "elves{enter}");
    expect(await screen.findByRole("alert")).toHaveTextContent("a build is in progress. The build continues on the server, and the deck shows on reload. (aborted) You can try again.");
  });

  it("a plain error is not retryable", async () => {
    chat.mockImplementationOnce(() => {
      throw new Error("boom");
    });
    await renderAt("/session/new");
    const user = userEvent.setup();
    await user.type(await screen.findByLabelText("Your message"), "elves{enter}");
    expect(await screen.findByRole("alert")).toHaveTextContent("boom (stream)");
    expect(screen.getByRole("alert")).not.toHaveTextContent("try again");
  });

  it("warns before a navigation or an unload while the agent works (D-303)", async () => {
    let release = () => {};
    const gate = new Promise<void>((r) => (release = r));
    chat.mockImplementationOnce(async function* () {
      yield ev("sessionStarted", "s1");
      await gate;
      yield ev("slots", {});
    });
    const { router } = await renderAt("/session/new");
    const user = userEvent.setup();
    await user.type(await screen.findByLabelText("Your message"), "elves");
    await user.click(screen.getByRole("button", { name: "Send" }));
    await screen.findByText("Session id: s1");
    // The move from /session/new to the session's own route is not blocked.
    expect(router.state.location.pathname).toBe("/session/s1");
    const unload = new Event("beforeunload", { cancelable: true });
    window.dispatchEvent(unload);
    expect(unload.defaultPrevented).toBe(true);

    await user.click(screen.getByRole("link", { name: "Decks" }));
    const dialog = await screen.findByRole("alertdialog");
    expect(dialog).toHaveTextContent("The build continues on the server.");
    expect(router.state.location.pathname).toBe("/session/s1");
    await user.click(within(dialog).getByRole("button", { name: "Stay" }));
    expect(screen.queryByRole("alertdialog")).not.toBeInTheDocument();
    await user.click(screen.getByRole("link", { name: "Decks" }));
    await user.click(within(await screen.findByRole("alertdialog")).getByRole("button", { name: "Leave" }));
    await waitFor(() => expect(router.state.location.pathname).toBe("/decks"));
    release();
    const after = new Event("beforeunload", { cancelable: true });
    window.dispatchEvent(after);
    expect(after.defaultPrevented).toBe(false);
  });

  it("Enter during an IME composition sends nothing", async () => {
    await renderAt("/session/new");
    const user = userEvent.setup();
    const box = await screen.findByLabelText("Your message");
    await user.type(box, "elves");
    fireEvent.keyDown(box, { key: "Enter", isComposing: true });
    expect(chat).not.toHaveBeenCalled();
    expect(box).toHaveValue("elves");
  });

  it("moves focus to a new question group", async () => {
    chat.mockReturnValueOnce(events([ev("sessionStarted", "s1"), ev("question", formatQuestion)]));
    await renderAt("/session/new");
    const user = userEvent.setup();
    await user.type(await screen.findByLabelText("Your message"), "elves{enter}");
    await screen.findByRole("group", { name: "Question: Which format?" });
    expect(screen.getByTestId("open-questions")).toHaveFocus();
  });

  it("refuses an answer over the 8 KiB cap", async () => {
    chat.mockReturnValueOnce(events([ev("sessionStarted", "s1"), ev("question", { id: "q2", slot: "power", text: "How strong?", options: [], optionOracleIds: [] })]));
    await renderAt("/session/new");
    const user = userEvent.setup();
    await user.type(await screen.findByLabelText("Your message"), "elves{enter}");
    const card = await screen.findByRole("group", { name: "Question: How strong?" });
    await user.click(within(card).getByRole("textbox"));
    await user.paste("x".repeat(8193));
    expect(screen.getByRole("alert")).toHaveTextContent("An answer is over the 8192 byte cap.");
    expect(screen.getByRole("button", { name: "Submit answers" })).toBeDisabled();
  });

  it("a text delta after a stored agent line starts a new line", async () => {
    getSession.mockResolvedValue({
      session: { id: "s1", collectionId: "", deckIds: [], turns: [{ userMessage: "elves", agentMessage: "Here is a plan.", questions: [], answers: [] }] },
    });
    chat.mockReturnValueOnce(events([ev("textDelta", "And "), ev("textDelta", "more.")]));
    await renderAt("/session/s1");
    const user = userEvent.setup();
    await user.type(await screen.findByLabelText("Your message"), "go on{enter}");
    await screen.findByText("And more.", { selector: "p.whitespace-pre-line" });
    expect(screen.getByText("Here is a plan.", { selector: "p.whitespace-pre-line" })).toBeInTheDocument();
  });

  it("scrolls the thread only while the reader is near the end", async () => {
    const scroll = vi.fn();
    Element.prototype.scrollIntoView = scroll;
    const rect = vi.spyOn(Element.prototype, "getBoundingClientRect").mockReturnValue({ top: 5000 } as DOMRect);
    chat.mockReturnValueOnce(events([ev("sessionStarted", "s1"), ev("textDelta", "hi")]));
    await renderAt("/session/new");
    const user = userEvent.setup();
    await user.type(await screen.findByLabelText("Your message"), "elves{enter}");
    await screen.findByText("hi", { selector: "p.whitespace-pre-line" });
    expect(scroll).not.toHaveBeenCalled();
    rect.mockReturnValue({ top: 0 } as DOMRect);
    chat.mockReturnValueOnce(events([ev("textDelta", "again")]));
    await waitFor(() => expect(screen.getByLabelText("Your message")).toBeEnabled());
    await user.type(screen.getByLabelText("Your message"), "more{enter}");
    await screen.findByText("again", { selector: "p.whitespace-pre-line" });
    expect(scroll).toHaveBeenCalled();
    rect.mockRestore();
  });

  it("forgets a stored session id the server does not know, and keeps it on another error", async () => {
    useAppStore.setState({ sessionId: "s1" });
    getSession.mockRejectedValueOnce(new ConnectError("down", Code.Unavailable));
    const first = await renderAt("/session/s1");
    await screen.findByRole("alert");
    expect(useAppStore.getState().sessionId).toBe("s1");
    first.unmount();
    getSession.mockRejectedValueOnce(new ConnectError("no session", Code.NotFound));
    await renderAt("/session/s1");
    await screen.findByRole("alert");
    await waitFor(() => expect(useAppStore.getState().sessionId).toBe(""));
  });

  it("a second visit to /session/new starts a fresh panel", async () => {
    chat.mockReturnValueOnce(events([ev("sessionStarted", "s1"), ev("textDelta", "hello")]));
    const { router } = await renderAt("/session/new");
    const user = userEvent.setup();
    await user.type(await screen.findByLabelText("Your message"), "elves{enter}");
    await screen.findByText("hello", { selector: "p.whitespace-pre-line" });
    await act(() => router.navigate("/session/new"));
    expect(await screen.findByTestId("session-id")).toHaveTextContent("No session yet.");
    expect(screen.queryByText("hello", { selector: "p.whitespace-pre-line" })).not.toBeInTheDocument();
  });

  it("Enter sends, and a failed send gives the draft back", async () => {
    chat.mockImplementationOnce(() => {
      throw new ConnectError("down", Code.Unavailable);
    });
    await renderAt("/session/new");
    const user = userEvent.setup();
    await user.type(await screen.findByLabelText("Your message"), "elves{enter}");
    expect(await screen.findByRole("alert")).toHaveTextContent("down (unavailable) You can try again.");
    expect(screen.getByLabelText("Your message")).toHaveValue("elves");
    expect((chat.mock.calls[0][0] as { message: string }).message).toBe("elves");
  });

  it("offers to resume the stored session on /session/new", async () => {
    useAppStore.setState({ sessionId: "s9" });
    await renderAt("/session/new");
    expect(await screen.findByTestId("resume-link")).toHaveAttribute("href", "/session/s9");
  });

  it("keeps a stored session's open question when a later turn asked another (reload merge)", async () => {
    getSession.mockResolvedValue({
      session: {
        id: "s1",
        collectionId: "",
        deckIds: [],
        turns: [
          { userMessage: "elves", agentMessage: "", questions: [formatQuestion, { id: "q2", slot: "colors", text: "Any color preference?", options: [], optionOracleIds: [] }], answers: [] },
          { userMessage: "", agentMessage: "", questions: [{ id: "q3", slot: "power", text: "How strong?", options: [], optionOracleIds: [] }], answers: [{ questionId: "q1", optionIndex: 2, text: "" }] },
        ],
      },
    });
    await renderAt("/session/s1");
    await screen.findByRole("group", { name: "Question: How strong?" });
    expect(screen.getByRole("group", { name: "Question: Any color preference?" })).toBeInTheDocument();
    expect(screen.queryByRole("group", { name: "Question: Which format?" })).not.toBeInTheDocument();
    // The clicked option shows as the user's line. The two open questions
    // sit in their cards, so the thread holds three lines: the message,
    // the answered question, and the answer.
    const lines = within(screen.getByRole("list", { name: "Conversation" })).getAllByRole("listitem");
    expect(lines).toHaveLength(3);
    expect(lines[1]).toHaveTextContent("Which format?");
    expect(lines.at(-1)).toHaveTextContent("Modern");
  });


  it("shows a commander offer as full card images (D-287)", async () => {
    getCards.mockResolvedValue({
      cards: [
        { oracleId: "o-ghalta", name: "Ghalta, Primal Hunger", typeLine: "Legendary Creature — Elder Dinosaur", manaCost: "{10}{G}{G}", oracleText: "Ghalta costs {X} less to cast.\nTrample", cardTypes: ["Creature"], faces: [], defaultPrinting: { artist: "Chase Stone", imageUris: { normal: "https://x/ghalta.jpg", small: "https://x/ghalta-s.jpg" } } },
        { oracleId: "o-reptil", name: "Reptil, Dinomorpher", typeLine: "Legendary Creature — Human Druid", manaCost: "{1}{G}", oracleText: "Whenever a Dinosaur enters, draw a card.", cardTypes: ["Creature"], faces: [], defaultPrinting: { artist: "Someone", imageUris: { normal: "https://x/reptil.jpg", small: "https://x/reptil-s.jpg" } } },
      ],
      missingOracleIds: [],
    });
    const offer = {
      id: "q5",
      slot: "commander",
      text: "Which one do you want: Ghalta, Primal Hunger, or Reptil, Dinomorpher?",
      options: ["Ghalta, Primal Hunger", "Reptil, Dinomorpher", "None, name three more"],
      optionOracleIds: ["o-ghalta", "o-reptil", ""],
    };
    chat.mockReturnValueOnce(events([ev("sessionStarted", "s1"), ev("question", offer)]));
    await renderAt("/session/new");
    const user = userEvent.setup();
    await user.type(await screen.findByLabelText("Your message"), "dinosaurs");
    await user.click(screen.getByRole("button", { name: "Send" }));
    const card = await screen.findByRole("group", { name: /Question: Which one/ });
    expect(await within(card).findByAltText("Ghalta, Primal Hunger (card)")).toHaveAttribute("src", "https://x/ghalta.jpg");
    expect(within(card).queryByText(/Illustrated by/)).not.toBeInTheDocument();
    // The image carries the rules text, so the tile repeats none of it.
    expect(within(card).queryByText("Legendary Creature — Elder Dinosaur")).not.toBeInTheDocument();
    expect(within(card).queryByText(/Ghalta costs \{X\} less to cast/)).not.toBeInTheDocument();
    expect(within(card).getAllByTestId("card-option")).toHaveLength(2);
    expect((getCards.mock.calls[0][0] as { oracleIds: string[] }).oracleIds).toEqual(["o-ghalta", "o-reptil"]);
    expect(within(card).queryAllByRole("status")).toHaveLength(0);
    // The card image falls back to the small one, as in the deck view.
    fireEvent.error(within(card).getByAltText("Ghalta, Primal Hunger (card)"));
    expect(within(card).getByAltText("Ghalta, Primal Hunger (card)")).toHaveAttribute("src", "https://x/ghalta-s.jpg");
    // The non-card option keeps a plain button, and every option still picks.
    await user.click(within(card).getByRole("button", { name: "None, name three more" }));
    expect(within(card).getByRole("button", { name: "None, name three more" })).toHaveAttribute("aria-pressed", "true");
    await user.click(within(card).getByRole("button", { name: "Reptil, Dinomorpher" }));
    expect(within(card).getByRole("button", { name: "Reptil, Dinomorpher" })).toHaveAttribute("aria-pressed", "true");
    expect(await axe(card)).toHaveNoViolations();
  });

  it("refuses a message over the 8 KiB cap", async () => {
    await renderAt("/session/new");
    const user = userEvent.setup();
    const box = await screen.findByLabelText("Your message");
    await user.click(box);
    await user.paste("x".repeat(8193));
    expect(screen.getByRole("alert")).toHaveTextContent("over the 8192 byte cap");
    expect(screen.getByRole("button", { name: "Send" })).toBeDisabled();
    expect(chat).not.toHaveBeenCalled();
  });

  it("rebuilds the thread of a stored session with no deck on reload", async () => {
    getSession.mockResolvedValue({
      session: {
        id: "s1",
        collectionId: "c1",
        deckIds: [],
        slots: { poolRule: PoolRule.OWNED_ONLY },
        usage: { calls: 4, inputTokens: 10n, outputTokens: 5n, priced: false },
        turns: [
          { userMessage: "elves", agentMessage: "", questions: [formatQuestion, { id: "q2", slot: "power", text: "How strong?", options: [], optionOracleIds: [] }], answers: [] },
          { userMessage: "", agentMessage: "", questions: [], answers: [{ questionId: "q1", optionIndex: 2, text: "" }] },
        ],
      },
    });
    getDeck.mockResolvedValue({ deck });
    await renderAt("/session/s1");
    expect(await screen.findByTestId("pool-mode")).toHaveTextContent("Pool: only cards in your collection.");
    expect(screen.getByTestId("session-id")).toHaveTextContent("Session id: s1");
    expect(screen.getByTestId("usage")).toHaveTextContent("cost unknown");
    expect(screen.queryByRole("group", { name: "Question: Which format?" })).not.toBeInTheDocument();
    expect(screen.getByRole("group", { name: "Question: How strong?" })).toBeInTheDocument();

    // The message box hides while a question waits (D-282).
    expect(screen.queryByLabelText("Your message")).not.toBeInTheDocument();
    chat.mockReturnValue(events([ev("slots", {})]));
    const user = userEvent.setup();
    const card = screen.getByRole("group", { name: "Question: How strong?" });
    await user.type(within(card).getByRole("textbox"), "bracket 2");
    await user.click(screen.getByRole("button", { name: "Submit answers" }));
    expect((chat.mock.calls[0][0] as { sessionId: string }).sessionId).toBe("s1");
    // A stored session carries its own collection, so the request sends none.
    expect((chat.mock.calls[0][0] as { collectionId: string }).collectionId).toBe("");
    // The box returns when no question is open, and the answer shows in the thread.
    expect(await screen.findByLabelText("Your message")).toBeInTheDocument();
  });


  it("a closed question with no options shows the text field", async () => {
    chat.mockReturnValueOnce(events([ev("sessionStarted", "s1"), ev("question", { id: "q2", slot: "power", text: "How strong?", options: [], optionOracleIds: [], closed: true })]));
    await renderAt("/session/new");
    const user = userEvent.setup();
    await user.type(await screen.findByLabelText("Your message"), "elves{enter}");
    const card = await screen.findByRole("group", { name: "Question: How strong?" });
    expect(within(card).getByRole("textbox")).toBeInTheDocument();
  });

  it("reports a session that does not load", async () => {
    getSession.mockRejectedValue(new Error("[not_found] no session"));
    await renderAt("/session/nope");
    expect(await screen.findByRole("alert")).toHaveTextContent("Could not load the session: [not_found] no session");
  });

  it("shows one loading status per card question while the cards load", async () => {
    getCards.mockReturnValue(new Promise(() => {}));
    const offer = { id: "q5", slot: "commander", text: "Which one?", options: ["A", "B"], optionOracleIds: ["o-a", "o-b"] };
    chat.mockReturnValueOnce(events([ev("sessionStarted", "s1"), ev("question", offer)]));
    await renderAt("/session/new");
    const user = userEvent.setup();
    await user.type(await screen.findByLabelText("Your message"), "elves{enter}");
    const card = await screen.findByRole("group", { name: "Question: Which one?" });
    expect(within(card).getAllByRole("status")).toHaveLength(1);
  });

  it("has no axe violations with a question on screen", async () => {
    chat.mockReturnValue(events([ev("sessionStarted", "s1"), ev("question", formatQuestion)]));
    const { container } = await renderAt("/session/new");
    const user = userEvent.setup();
    await user.type(await screen.findByLabelText("Your message"), "elves");
    await user.click(screen.getByRole("button", { name: "Send" }));
    await screen.findByRole("group", { name: "Question: Which format?" });
    expect(await axe(container)).toHaveNoViolations();
  });

  it("poolLabel names the pool rule, and the collection before the rule is set", async () => {
    expect(poolLabel(PoolRule.OWNED_ONLY, "c1")).toBe("Pool: only cards in your collection.");
    expect(poolLabel(PoolRule.OWNED_FIRST, "")).toBe("Pool: your collection first, with upgrades to buy.");
    expect(poolLabel(PoolRule.ANY_CARD, "binder-july.csv")).toBe("Pool: any card.");
    expect(poolLabel(undefined, "binder-july.csv")).toBe("Pool: binder-july.csv. The agent asks how strict.");
    expect(poolLabel(undefined, "")).toBe("Pool: any card.");
  });

  it("pruneDrafts drops the drafts of questions no longer open", async () => {
    expect(pruneDrafts({ q1: { text: "a" }, q2: { optionIndex: 1, text: "" } }, [{ id: "q2" }])).toEqual({ q2: { optionIndex: 1, text: "" } });
  });
});

// The picker is the first control of a new chat. Before it, the only way
// to change the pool was the header menu, and a reader did not find it.
describe("the pool picker", () => {
  it("lists any card, every collection, and the way to add one", async () => {
    await renderAt("/session/new");
    const select = (await screen.findByLabelText("Build from")) as HTMLSelectElement;
    expect([...select.options].map((o) => o.textContent)).toEqual(["Any card", "binder-july.csv (4317 cards)", "Add a collection..."]);
    expect(select.value).toBe("");
  });

  it("takes a collection and sets the owned pool", async () => {
    await renderAt("/session/new");
    const select = await screen.findByLabelText("Build from");
    await userEvent.setup().selectOptions(select, "c1");
    expect(useAppStore.getState().collectionId).toBe("c1");
    expect(useAppStore.getState().poolMode).toBe("owned");
    expect(await screen.findByLabelText("Use only cards in my collection")).toBeChecked();
  });

  it("any card clears the collection", async () => {
    useAppStore.setState({ collectionId: "c1", poolMode: "owned" });
    await renderAt("/session/new");
    const select = await screen.findByLabelText("Build from");
    await userEvent.setup().selectOptions(select, "");
    expect(useAppStore.getState().collectionId).toBe("");
    expect(useAppStore.getState().poolMode).toBe("any");
  });

  it("add a collection goes to the upload page", async () => {
    const { router } = await renderAt("/session/new");
    const select = await screen.findByLabelText("Build from");
    await userEvent.setup().selectOptions(select, "add:collection");
    expect(router.state.location.pathname).toBe("/collection");
  });

  it("leaves once the session starts", async () => {
    chat.mockReturnValue(events([ev("sessionStarted", "s1")]));
    await renderAt("/session/new");
    await screen.findByLabelText("Build from");
    const user = userEvent.setup();
    await user.type(screen.getByLabelText("Your message"), "elves");
    await user.click(screen.getByRole("button", { name: "Send" }));
    await screen.findByText("Session id: s1");
    expect(screen.queryByLabelText("Build from")).not.toBeInTheDocument();
  });
});

// A question sits in the thread and in the open list at the same time.
// The card takes the answer, so the thread waits for the answer.
describe("a question that is open", () => {
  it("shows once, and joins the thread after the answer", async () => {
    chat.mockReturnValue(events([ev("sessionStarted", "s1"), ev("question", formatQuestion)]));
    await renderAt("/session/new");
    const user = userEvent.setup();
    await user.type(screen.getByLabelText("Your message"), "elves");
    await user.click(screen.getByRole("button", { name: "Send" }));

    const card = await screen.findByRole("group", { name: "Question: Which format?" });
    expect(card).toBeInTheDocument();
    const thread = screen.getByRole("list", { name: "Conversation" });
    expect(within(thread).queryByText("Which format?")).not.toBeInTheDocument();

    chat.mockReturnValue(events([ev("slots", {})]));
    await user.click(within(card).getByRole("button", { name: "Commander" }));
    await user.click(screen.getByRole("button", { name: "Submit answers" }));

    // The answer closes the card, and the asked line takes its place.
    await waitFor(() => expect(screen.queryByRole("group", { name: "Question: Which format?" })).not.toBeInTheDocument());
    expect(within(screen.getByRole("list", { name: "Conversation" })).getByText("Which format?")).toBeInTheDocument();
  });
});

// "Add a collection" goes to the upload screen and opens the file dialog
// there. A cancelled dialog therefore leaves no collection active.
describe("add a collection from the picker", () => {
  it("clears the collection and asks the upload screen for a file", async () => {
    useAppStore.setState({ collectionId: "c1", poolMode: "owned" });
    const clicks: string[] = [];
    const realClick = HTMLInputElement.prototype.click;
    HTMLInputElement.prototype.click = function click(this: HTMLInputElement) {
      clicks.push(this.type);
    };
    try {
      const { router } = await renderAt("/session/new");
      await userEvent.setup().selectOptions(await screen.findByLabelText("Build from"), "add:collection");
      expect(router.state.location.pathname).toBe("/collection");
      expect(useAppStore.getState().collectionId).toBe("");
      expect(useAppStore.getState().poolMode).toBe("any");
      await waitFor(() => expect(clicks).toContain("file"));
      // The mark leaves the history entry, so a reload opens no dialog.
      expect(router.state.location.state).toBeNull();
    } finally {
      HTMLInputElement.prototype.click = realClick;
    }
  });
});

// The pool is a choice of one chat, not a setting of the app (D-345).
// Only the session id reaches localStorage, so a new load reads no
// collection and the picker opens at any card.
describe("the pool of a new chat", () => {
  it("keeps the pool out of the stored state", async () => {
    useAppStore.getState().setCollection("c1");
    await waitFor(() => expect(localStorage.getItem("mtg-deck-builder")).not.toBeNull());
    const stored = JSON.parse(localStorage.getItem("mtg-deck-builder") ?? "{}") as { state: Record<string, unknown> };
    expect(Object.keys(stored.state)).toEqual(["sessionId"]);
  });

  it("opens at any card when nothing chose a collection", async () => {
    await renderAt("/session/new");
    expect(await screen.findByLabelText("Build from")).toHaveValue("");
    expect(screen.queryByLabelText("Use only cards in my collection")).not.toBeInTheDocument();
  });
});
