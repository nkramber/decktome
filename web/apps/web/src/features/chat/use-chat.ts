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
  // baseDeck is the deck the latest one revised, for the diff (PR-12B).
  baseDeck?: Deck;
  busy: boolean;
};

// The message cap of the API (agentsvc.MaxMessageBytes).
export const maxMessageBytes = 8 << 10;

export function byteLength(s: string): number {
  return new TextEncoder().encode(s).length;
}

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

// fromSession rebuilds the thread from a stored session, so a reload
// returns to the chat. The open questions follow the same merge rule as
// the live stream, minus the ones a later turn answered (D-237, D-278).
export function fromSession(session: Session, deck?: Deck): ChatState {
  const thread: ThreadItem[] = [];
  const answered = new Set<string>();
  const asked = new Map<string, Question>();
  let open: Question[] = [];
  for (const turn of session.turns) {
    const replies = turn.answers.map((a) => {
      answered.add(a.questionId);
      return answerText(a, asked.get(a.questionId));
    });
    const userText = [turn.userMessage, ...replies].filter(Boolean).join("\n");
    if (userText) thread.push({ kind: "user", text: userText });
    if (turn.agentMessage) thread.push({ kind: "agent", text: turn.agentMessage });
    for (const q of turn.questions) {
      asked.set(q.id, q);
      thread.push({ kind: "question", question: q });
    }
    if (turn.questions.length > 0) open = mergeOpen(open, turn.questions);
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
// end. Unmount and stop() abort the stream (ui plan, section 5).
//
// The session id and the collection id live in refs, so a send that
// starts before React commits the session_started event still carries
// the right id, and the collection goes with the first message only.
export function useChat(initial: ChatState, collectionId: string, onSessionStarted?: (id: string) => void) {
  const [state, setState] = useState<ChatState>(initial);
  const abort = useRef<AbortController | null>(null);
  const sessionId = useRef(initial.sessionId);
  const mounted = useRef(true);

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

  // send returns true when the stream ended without a thrown error, so
  // the caller can restore a draft the user should not lose.
  const send = useCallback(
    async ({ message, answers }: SendInput): Promise<boolean> => {
      const controller = new AbortController();
      abort.current?.abort();
      abort.current = controller;
      const answeredIds = new Set(answers.map((a) => a.questionId));
      let ok = true;
      const asked: Question[] = [];
      update((s) => {
        const userText = [message, ...answers.map((a) => answerText(a, s.openQuestions.find((q) => q.id === a.questionId)))]
          .filter(Boolean)
          .join("\n");
        return {
          ...s,
          busy: true,
          thread: userText ? [...s.thread, { kind: "user", text: userText }] : s.thread,
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
          apply(res, asked, update, (id) => {
            sessionId.current = id;
            onSessionStarted?.(id);
          });
        }
      } catch (err) {
        ok = false;
        if (!controller.signal.aborted) {
          update((s) => ({
            ...s,
            thread: [...s.thread, { kind: "failure", failure: { code: "stream", message: errorMessage(err), retryable: true } as AgentError }],
          }));
        }
      } finally {
        // A superseded or stopped send leaves the state to the live one.
        if (abort.current === controller) {
          abort.current = null;
          update((s) => ({ ...s, busy: false, openQuestions: mergeOpen(s.openQuestions, asked) }));
        }
      }
      return ok && !controller.signal.aborted;
    },
    [collectionId, onSessionStarted, update],
  );

  return { state, send, stop };
}

function apply(res: ChatResponse, asked: Question[], update: (f: (s: ChatState) => ChatState) => void, onSessionStarted: (id: string) => void) {
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
      update((s) => {
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
      update((s) => ({ ...s, thread: [...s.thread, { kind: "question", question }] }));
      return;
    }
    case "slots":
      update((s) => ({ ...s, slots: ev.value }));
      return;
    case "status":
      update((s) => ({ ...s, thread: [...s.thread, { kind: "status", text: ev.value }] }));
      return;
    case "deck": {
      const deck = ev.value;
      update((s) => ({ ...s, deck, baseDeck: s.deck, thread: [...s.thread, { kind: "deck", deck }] }));
      return;
    }
    case "failure":
      update((s) => ({ ...s, thread: [...s.thread, { kind: "failure", failure: ev.value }] }));
      return;
    case "error":
      update((s) => ({
        ...s,
        thread: [...s.thread, { kind: "failure", failure: { code: "error", message: ev.value, retryable: false } as AgentError }],
      }));
      return;
    case "usage":
      update((s) => ({ ...s, usage: ev.value }));
      return;
    default:
      return;
  }
}
