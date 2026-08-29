import { Code, ConnectError } from "@connectrpc/connect";
import type { AgentError, ChatResponse } from "@mtg/api-client/mtg/v1/agent_service_pb";
import type { Deck } from "@mtg/api-client/mtg/v1/deck_pb";
import type { Answer, Question, Session, Slots, Usage } from "@mtg/api-client/mtg/v1/session_pb";
import type { Timestamp } from "@bufbuild/protobuf/wkt";
import { useCallback, useEffect, useRef, useState } from "react";

import { agentClient } from "../../lib/api";
import { errorMessage } from "../../lib/errors";
import { byteLength, maxMessageBytes } from "../../lib/limits";

export { byteLength, maxMessageBytes };

// One line of the thread. Every streamed event gets a line (ui plan,
// step 3), and the questions also sit in `openQuestions` until answered.
// The id is unique for the life of the page, so React keys stay stable.
type ItemBody =
  | { kind: "user"; text: string }
  | { kind: "agent"; text: string }
  | { kind: "status"; text: string }
  | { kind: "question"; question: Question }
  | { kind: "failure"; failure: AgentError }
  | { kind: "deck"; deck: Deck };

export type ThreadItem = ItemBody & { id: number };

let itemSeq = 0;

export function item(body: ItemBody): ThreadItem {
  itemSeq += 1;
  return { ...body, id: itemSeq };
}

export type ChatState = {
  sessionId: string;
  thread: ThreadItem[];
  openQuestions: Question[];
  slots?: Slots;
  usage?: Usage;
  deck?: Deck;
  // baseDeck is the deck the latest one revised, for the diff (PR-12B).
  baseDeck?: Deck;
  busy: boolean;
};

// mergeOpen keeps every open question the turn did not answer, replaces
// one whose slot the turn asked again, and adds the new ones (D-278). The
// server never repeats a question that is out, so this is the only place
// an earlier question survives.
export function mergeOpen(open: Question[], asked: Question[]): Question[] {
  const slots = new Set(asked.map((q) => q.slot));
  const ids = new Set(asked.map((q) => q.id));
  return [...open.filter((q) => !ids.has(q.id) && !slots.has(q.slot)), ...asked];
}

// answerText is what the user said: the text, or the option the index
// names. Empty when neither is set.
export function answerText(a: Answer, question: Question | undefined): string {
  if (a.text) return a.text;
  if (a.optionIndex !== undefined && question) return question.options[a.optionIndex] ?? "";
  return "";
}

function seconds(t: Timestamp | undefined): number | undefined {
  return t ? Number(t.seconds) + t.nanos / 1e9 : undefined;
}

// fromSession rebuilds the thread from a stored session, so a reload
// returns to the chat. The open questions follow the same merge rule as
// the live stream, minus the ones a later turn answered (D-237, D-278).
// The deck line sits before the first turn that came after the deck, by
// created_at against the turns' at, and at the end when a time is missing.
export function fromSession(session: Session, deck?: Deck, baseDeck?: Deck): ChatState {
  const thread: ThreadItem[] = [];
  const answered = new Set<string>();
  const asked = new Map<string, Question>();
  let open: Question[] = [];
  const deckAt = seconds(deck?.createdAt);
  let deckPlaced = deck === undefined;
  for (const turn of session.turns) {
    const turnAt = seconds(turn.at);
    if (!deckPlaced && deck && deckAt !== undefined && turnAt !== undefined && turnAt > deckAt) {
      thread.push(item({ kind: "deck", deck }));
      deckPlaced = true;
    }
    const replies = turn.answers.map((a) => {
      answered.add(a.questionId);
      return answerText(a, asked.get(a.questionId));
    });
    const userText = [turn.userMessage, ...replies].filter(Boolean).join("\n");
    if (userText) thread.push(item({ kind: "user", text: userText }));
    if (turn.agentMessage) thread.push(item({ kind: "agent", text: turn.agentMessage }));
    for (const q of turn.questions) {
      asked.set(q.id, q);
      thread.push(item({ kind: "question", question: q }));
    }
    if (turn.questions.length > 0) open = mergeOpen(open, turn.questions);
  }
  if (!deckPlaced && deck) thread.push(item({ kind: "deck", deck }));
  return {
    sessionId: session.id,
    thread,
    openQuestions: open.filter((q) => !answered.has(q.id)),
    slots: session.slots,
    usage: session.usage,
    deck,
    baseDeck: deck?.revisedFromDeckId && baseDeck?.id === deck.revisedFromDeckId ? baseDeck : undefined,
    busy: false,
  };
}

export const emptyState: ChatState = { sessionId: "", thread: [], openQuestions: [], busy: false };

export type SendInput = { message: string; answers: Answer[] };

// The stream error codes a user can retry at once. Aborted means a build
// is in progress (D-303): the user waits for it, then tries again.
const retryableCodes = new Set([Code.Unavailable, Code.ResourceExhausted, Code.DeadlineExceeded, Code.Aborted]);

export function codeName(code: Code): string {
  return Code[code].replace(/([a-z])([A-Z])/g, "$1_$2").toLowerCase();
}

// streamFailure maps a thrown stream error to a failure line. A
// ConnectError keeps its code, and anything else is a "stream" failure
// nobody should retry blind.
export function streamFailure(err: unknown): AgentError {
  if (err instanceof ConnectError) {
    const aborted = err.code === Code.Aborted;
    return {
      code: codeName(err.code),
      message: aborted ? `${err.rawMessage}. The build continues on the server, and the deck shows on reload.` : err.rawMessage,
      retryable: retryableCodes.has(err.code),
    } as AgentError;
  }
  return { code: "stream", message: errorMessage(err), retryable: false } as AgentError;
}

// SendResult says how the send ended. `ok` is true when the stream ended
// without a thrown error and without a stop, so the caller can restore a
// draft the user should not lose. `restored` names the questions the send
// gave back to the open list on failure, so the caller restores their drafts.
export type SendResult = { ok: boolean; restored: Question[] };

// useChat drives the Chat stream. One call to send reads the stream to
// its end, appends every event, and stops on failure or on the stream
// end. Unmount and stop() abort the stream (ui plan, section 5).
//
// The session id and the collection id live in refs, so a send that
// starts before React commits the session_started event still carries
// the right id, and the collection goes with the first message only.
export function useChat(initial: ChatState, collectionId: string, onSessionStarted?: (id: string) => void) {
  const [state, setState] = useState<ChatState>(initial);
  // latest mirrors the committed state, so a send that fails before the
  // first render still knows which questions it took off the open list.
  const latest = useRef(state);
  useEffect(() => {
    latest.current = state;
  }, [state]);
  const abort = useRef<AbortController | null>(null);
  const sessionId = useRef(initial.sessionId);
  const mounted = useRef(true);
  const started = useRef(onSessionStarted);
  useEffect(() => {
    started.current = onSessionStarted;
  }, [onSessionStarted]);

  useEffect(() => {
    mounted.current = true;
    return () => {
      mounted.current = false;
      abort.current?.abort();
    };
  }, []);

  const update = useCallback((f: (s: ChatState) => ChatState) => {
    if (mounted.current) setState(f);
  }, []);

  const stop = useCallback(() => abort.current?.abort(), []);

  const send = useCallback(
    async ({ message, answers }: SendInput): Promise<SendResult> => {
      const controller = new AbortController();
      abort.current?.abort();
      abort.current = controller;
      const answeredIds = new Set(answers.map((a) => a.questionId));
      let ok = true;
      const ctx: StreamContext = { asked: [], newAgentLine: true };
      const answeredQuestions = latest.current.openQuestions.filter((q) => answeredIds.has(q.id));
      update((s) => {
        const userText = [message, ...answers.map((a) => answerText(a, s.openQuestions.find((q) => q.id === a.questionId)))]
          .filter(Boolean)
          .join("\n");
        return {
          ...s,
          busy: true,
          thread: userText ? [...s.thread, item({ kind: "user", text: userText })] : s.thread,
          openQuestions: s.openQuestions.filter((q) => !answeredIds.has(q.id)),
        };
      });
      try {
        const stream = agentClient.chat(
          { sessionId: sessionId.current, collectionId: sessionId.current === "" ? collectionId : "", message, answers },
          { signal: controller.signal },
        );
        for await (const res of stream) {
          if (controller.signal.aborted) break;
          apply(res, ctx, update, (id) => {
            sessionId.current = id;
            started.current?.(id);
          });
        }
      } catch (err) {
        ok = false;
        if (!controller.signal.aborted) {
          update((s) => ({ ...s, thread: [...s.thread, item({ kind: "failure", failure: streamFailure(err) })] }));
        }
      } finally {
        // A superseded or stopped send leaves the state to the live one.
        if (abort.current === controller) {
          abort.current = null;
          ok = ok && !controller.signal.aborted;
          // A failed or stopped send gives the answered questions back,
          // so the user can submit again.
          const restore = ok ? [] : answeredQuestions;
          update((s) => ({ ...s, busy: false, openQuestions: mergeOpen(mergeOpen(s.openQuestions, restore), ctx.asked) }));
        }
      }
      return { ok: ok && !controller.signal.aborted, restored: ok ? [] : answeredQuestions };
    },
    [collectionId, update],
  );

  return { state, send, stop };
}

// StreamContext is what one send collects over its stream: the questions
// it asked, and whether the next text delta starts a new agent line.
type StreamContext = { asked: Question[]; newAgentLine: boolean };

function apply(res: ChatResponse, ctx: StreamContext, update: (f: (s: ChatState) => ChatState) => void, onSessionStarted: (id: string) => void) {
  const ev = res.event;
  switch (ev.case) {
    case "sessionStarted": {
      const id = ev.value;
      update((s) => ({ ...s, sessionId: id }));
      onSessionStarted(id);
      return;
    }
    case "textDelta": {
      const delta = ev.value;
      const fresh = ctx.newAgentLine;
      ctx.newAgentLine = false;
      update((s) => {
        const last = s.thread[s.thread.length - 1];
        if (!fresh && last?.kind === "agent") {
          return { ...s, thread: [...s.thread.slice(0, -1), { ...last, text: last.text + delta }] };
        }
        return { ...s, thread: [...s.thread, item({ kind: "agent", text: delta })] };
      });
      return;
    }
    case "question": {
      const question = ev.value;
      ctx.asked.push(question);
      update((s) => ({ ...s, thread: [...s.thread, item({ kind: "question", question })] }));
      return;
    }
    case "slots":
      update((s) => ({ ...s, slots: ev.value }));
      return;
    case "status":
      update((s) => ({ ...s, thread: [...s.thread, item({ kind: "status", text: ev.value })] }));
      return;
    case "deck": {
      const deck = ev.value;
      update((s) => ({ ...s, deck, baseDeck: s.deck, thread: [...s.thread, item({ kind: "deck", deck })] }));
      return;
    }
    case "failure":
      update((s) => ({ ...s, thread: [...s.thread, item({ kind: "failure", failure: ev.value })] }));
      return;
    case "error":
      update((s) => ({
        ...s,
        thread: [...s.thread, item({ kind: "failure", failure: { code: "error", message: ev.value, retryable: false } as AgentError })],
      }));
      return;
    case "usage":
      update((s) => ({ ...s, usage: ev.value }));
      return;
    default:
      return;
  }
}
