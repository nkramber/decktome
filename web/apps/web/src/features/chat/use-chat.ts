import type { AgentError, ChatResponse } from "@mtg/api-client/mtg/v1/agent_service_pb";
import type { Deck } from "@mtg/api-client/mtg/v1/deck_pb";
import type { Answer, Question, Session, Slots, Usage } from "@mtg/api-client/mtg/v1/session_pb";
import { useCallback, useEffect, useRef, useState } from "react";

import { agentClient } from "../../lib/api";
import { errorMessage } from "../../lib/errors";

// One line of the thread. Every streamed event gets a line (ui plan,
// step 3), and the questions also sit in `openQuestions` until answered.
export type ThreadItem =
  | { kind: "user"; text: string }
  | { kind: "agent"; text: string }
  | { kind: "status"; text: string }
  | { kind: "question"; question: Question }
  | { kind: "failure"; failure: AgentError }
  | { kind: "deck"; deck: Deck };

export type ChatState = {
  sessionId: string;
  thread: ThreadItem[];
  openQuestions: Question[];
  slots?: Slots;
  usage?: Usage;
  deck?: Deck;
  busy: boolean;
};

// The message cap of the API (agentsvc.MaxMessageBytes).
export const maxMessageBytes = 8 << 10;

export function byteLength(s: string): number {
  return new TextEncoder().encode(s).length;
}

// fromSession rebuilds the thread from a stored session, so a reload
// returns to the chat. The open questions are those of the last turn that
// asked any, minus the ones a later turn answered. A turn that fills some
// slots and asks nothing sends no question, so the UI holds the earlier
// ones (D-237).
export function fromSession(session: Session, deck?: Deck): ChatState {
  const thread: ThreadItem[] = [];
  const answered = new Set<string>();
  let open: Question[] = [];
  for (const turn of session.turns) {
    for (const a of turn.answers) answered.add(a.questionId);
    const userText = [turn.userMessage, ...turn.answers.map((a) => a.text).filter(Boolean)].filter(Boolean).join("\n");
    if (userText) thread.push({ kind: "user", text: userText });
    if (turn.agentMessage) thread.push({ kind: "agent", text: turn.agentMessage });
    for (const q of turn.questions) thread.push({ kind: "question", question: q });
    if (turn.questions.length > 0) open = turn.questions;
  }
  if (deck) thread.push({ kind: "deck", deck });
  return {
    sessionId: session.id,
    thread,
    openQuestions: open.filter((q) => !answered.has(q.id)),
    slots: session.slots,
    usage: session.usage,
    deck,
    busy: false,
  };
}

export const emptyState: ChatState = { sessionId: "", thread: [], openQuestions: [], busy: false };

export type SendInput = { message: string; answers: Answer[] };

// useChat drives the Chat stream. One call to send reads the stream to
// its end, appends every event, and stops on failure or on the stream
// end. Unmount aborts the stream (ui plan, section 5).
export function useChat(initial: ChatState, collectionId: string, onSessionStarted?: (id: string) => void) {
  const [state, setState] = useState<ChatState>(initial);
  const abort = useRef<AbortController | null>(null);

  useEffect(() => () => abort.current?.abort(), []);

  const send = useCallback(
    async ({ message, answers }: SendInput) => {
      const controller = new AbortController();
      abort.current?.abort();
      abort.current = controller;
      const answeredIds = new Set(answers.map((a) => a.questionId));
      const userText = [message, ...answers.map((a) => a.text)].filter(Boolean).join("\n");
      setState((s) => ({
        ...s,
        busy: true,
        thread: userText ? [...s.thread, { kind: "user", text: userText }] : s.thread,
        openQuestions: s.openQuestions.filter((q) => !answeredIds.has(q.id)),
      }));
      const asked: Question[] = [];
      try {
        const stream = agentClient.chat(
          { sessionId: state.sessionId, collectionId, message, answers },
          { signal: controller.signal },
        );
        for await (const res of stream) {
          apply(res, asked, setState, onSessionStarted);
        }
      } catch (err) {
        if (!controller.signal.aborted) {
          setState((s) => ({
            ...s,
            thread: [...s.thread, { kind: "failure", failure: { code: "stream", message: errorMessage(err), retryable: true } as AgentError }],
          }));
        }
      } finally {
        if (abort.current === controller) abort.current = null;
        // A question the server did not answer stays open. A new question
        // for the same slot replaces the old one, and the rest join the
        // list (D-278). The server never repeats a question that is out.
        setState((s) => ({ ...s, busy: false, openQuestions: mergeOpen(s.openQuestions, asked) }));
      }
    },
    [state.sessionId, collectionId, onSessionStarted],
  );

  return { state, send };
}

export function mergeOpen(open: Question[], asked: Question[]): Question[] {
  const slots = new Set(asked.map((q) => q.slot));
  const ids = new Set(asked.map((q) => q.id));
  return [...open.filter((q) => !ids.has(q.id) && !slots.has(q.slot)), ...asked];
}

function apply(
  res: ChatResponse,
  asked: Question[],
  setState: (f: (s: ChatState) => ChatState) => void,
  onSessionStarted?: (id: string) => void,
) {
  const ev = res.event;
  switch (ev.case) {
    case "sessionStarted": {
      const id = ev.value;
      setState((s) => ({ ...s, sessionId: id }));
      onSessionStarted?.(id);
      return;
    }
    case "textDelta": {
      const delta = ev.value;
      setState((s) => {
        const last = s.thread[s.thread.length - 1];
        if (last?.kind === "agent") {
          return { ...s, thread: [...s.thread.slice(0, -1), { kind: "agent", text: last.text + delta }] };
        }
        return { ...s, thread: [...s.thread, { kind: "agent", text: delta }] };
      });
      return;
    }
    case "question": {
      const question = ev.value;
      asked.push(question);
      setState((s) => ({ ...s, thread: [...s.thread, { kind: "question", question }] }));
      return;
    }
    case "slots":
      setState((s) => ({ ...s, slots: ev.value }));
      return;
    case "status":
      setState((s) => ({ ...s, thread: [...s.thread, { kind: "status", text: ev.value }] }));
      return;
    case "deck": {
      const deck = ev.value;
      setState((s) => ({ ...s, deck, thread: [...s.thread, { kind: "deck", deck }] }));
      return;
    }
    case "failure":
      setState((s) => ({ ...s, thread: [...s.thread, { kind: "failure", failure: ev.value }] }));
      return;
    case "error":
      setState((s) => ({
        ...s,
        thread: [...s.thread, { kind: "failure", failure: { code: "error", message: ev.value, retryable: false } as AgentError }],
      }));
      return;
    case "usage":
      setState((s) => ({ ...s, usage: ev.value }));
      return;
    default:
      return;
  }
}
