import { PoolRule, type Session } from "@mtg/api-client/mtg/v1/session_pb";
import { useQuery } from "@tanstack/react-query";
import { type FormEvent, type KeyboardEvent, useCallback, useEffect, useRef, useState } from "react";
import { Link, useNavigate, useParams } from "react-router";

import { agentClient, deckClient } from "../../lib/api";
import { errorMessage } from "../../lib/errors";
import { useAppStore } from "../../lib/store";
import { DeckView } from "../deck/deck-view";
import { QuestionCard } from "./question-card";
import { byteLength, type ChatState, emptyState, fromSession, maxMessageBytes, type ThreadItem, useChat } from "./use-chat";

// The chat screen (ui plan, step 3). The id "new" means no session yet.
// A stored session loads through GetSession and its latest deck through
// GetDeck, so a reload returns to the thread.
export function SessionPage() {
  const { id } = useParams();
  const isNew = !id || id === "new";
  const storedSessionId = useAppStore((s) => s.sessionId);
  const setSessionId = useAppStore((s) => s.setSessionId);

  // A session this page started stays mounted when the route moves from
  // /session/new to /session/<id>. Nothing reloads it, so the stream and
  // the thread survive the navigate. The ref is written before the
  // navigate, so no render sees the new id without it. A later visit to
  // /session/new starts a fresh panel through the nonce.
  const started = useRef("");
  const [nonce, setNonce] = useState(0);
  useEffect(() => {
    if (isNew && started.current) {
      started.current = "";
      setNonce((n) => n + 1);
    }
  }, [isNew]);
  const live = isNew || id === started.current;
  const onStarted = useCallback((sid: string) => {
    started.current = sid;
  }, []);

  const session = useQuery({
    queryKey: ["session", id],
    queryFn: () => agentClient.getSession({ sessionId: id ?? "" }),
    enabled: !live,
  });
  const deckId = session.data?.session?.deckIds.at(-1) ?? "";
  const deck = useQuery({
    queryKey: ["deck", deckId],
    queryFn: () => deckClient.getDeck({ deckId }),
    enabled: deckId !== "",
  });

  if (live) {
    return (
      <ChatPanel
        key={`live-${nonce}`}
        initial={emptyState}
        onStarted={onStarted}
        resumeId={isNew && storedSessionId ? storedSessionId : ""}
      />
    );
  }
  if (session.isPending || (deckId && deck.isPending)) {
    return (
      <div className="p-6">
        <h1 className="text-2xl font-semibold">Chat</h1>
        <p className="wrap-anywhere text-sm text-neutral-600" data-testid="session-id">
          Session id: {id}
        </p>
        <p role="status" className="text-neutral-600">
          Loading the session...
        </p>
      </div>
    );
  }
  if (session.isError || !session.data.session) {
    if (storedSessionId === id) setSessionId("");
    return (
      <div className="p-6">
        <h1 className="text-2xl font-semibold">Chat</h1>
        <p className="wrap-anywhere text-sm text-neutral-600" data-testid="session-id">
          Session id: {id}
        </p>
        <p role="alert" className="text-red-700">
          Could not load the session: {session.isError ? errorMessage(session.error) : "the server returned no session"}
        </p>
        <Link to="/session/new" className="underline">
          Start a new chat
        </Link>
      </div>
    );
  }
  return (
    <ChatPanel
      key={id}
      initial={fromSession(session.data.session, deck.data?.deck)}
      session={session.data.session}
      deckError={deck.isError ? errorMessage(deck.error) : ""}
    />
  );
}

function ChatPanel({
  initial,
  session,
  onStarted,
  resumeId = "",
  deckError = "",
}: {
  initial: ChatState;
  session?: Session;
  onStarted?: (id: string) => void;
  resumeId?: string;
  deckError?: string;
}) {
  const navigate = useNavigate();
  const collectionId = useAppStore((s) => s.collectionId);
  const poolMode = useAppStore((s) => s.poolMode);
  const setPoolMode = useAppStore((s) => s.setPoolMode);
  const setSessionId = useAppStore((s) => s.setSessionId);

  // The collection goes with the first message only. A stored session
  // holds its own collection id, and the server reads that one.
  const sendCollection = initial.sessionId === "" && poolMode === "owned" ? collectionId : (session?.collectionId ?? "");

  const onSessionStarted = useCallback(
    (sid: string) => {
      setSessionId(sid);
      onStarted?.(sid);
      navigate(`/session/${sid}`, { replace: true });
    },
    [navigate, setSessionId, onStarted],
  );
  const { state, send, stop } = useChat(initial, sendCollection, onSessionStarted);
  const [message, setMessage] = useState("");
  const bytes = byteLength(message);
  const tooLong = bytes > maxMessageBytes;
  const beforeFirstMessage = state.sessionId === "";
  // The message box hides while a question waits. The user answers the
  // question, and the box returns when none is open (owner, 2026-08-28).
  const showComposer = state.openQuestions.length === 0;

  // New output scrolls into view. The sentinel sits under the thread.
  const end = useRef<HTMLDivElement>(null);
  const last = state.thread[state.thread.length - 1];
  const lastLength = last && "text" in last ? last.text.length : 0;
  useEffect(() => {
    end.current?.scrollIntoView?.({ block: "nearest" });
  }, [state.thread.length, lastLength, state.openQuestions.length]);

  async function submit() {
    const text = message.trim();
    if (!text || tooLong || state.busy) return;
    setMessage("");
    const ok = await send({ message: text, answers: [] });
    // A failed send gives the draft back, so the user does not retype it.
    if (!ok) setMessage((m) => m || text);
  }

  function onSubmit(e: FormEvent) {
    e.preventDefault();
    void submit();
  }

  // Enter sends, and Shift+Enter makes a new line.
  function onKeyDown(e: KeyboardEvent<HTMLTextAreaElement>) {
    if (e.key === "Enter" && !e.shiftKey) {
      e.preventDefault();
      void submit();
    }
  }

  const poolText = poolLabel(state.slots?.poolRule, sendCollection);

  return (
    <div className="mx-auto flex max-w-7xl flex-col gap-4 p-6 lg:flex-row">
      <section aria-labelledby="chat-title" className="flex min-w-0 flex-1 flex-col gap-3 lg:max-w-xl">
        <h1 id="chat-title" className="text-2xl font-semibold">
          Chat
        </h1>
        <p className="wrap-anywhere text-sm text-neutral-600" data-testid="session-id">
          {beforeFirstMessage ? "No session yet." : `Session id: ${state.sessionId}`}
        </p>
        {beforeFirstMessage && resumeId && (
          <p className="text-sm">
            <Link to={`/session/${resumeId}`} className="underline" data-testid="resume-link">
              Resume your last chat
            </Link>
          </p>
        )}
        {beforeFirstMessage && collectionId ? (
          <label className="flex items-center gap-2 text-sm">
            <input type="checkbox" checked={poolMode === "owned"} onChange={(e) => setPoolMode(e.target.checked ? "owned" : "any")} />
            Use only cards in my collection
          </label>
        ) : null}
        <p className="wrap-anywhere text-sm text-neutral-600" data-testid="pool-mode">
          {poolText}
        </p>

        <ol className="flex flex-col gap-2" aria-label="Conversation" role="list">
          {state.thread.map((item, i) => (
            <li key={i}>
              <ThreadLine item={item} />
            </li>
          ))}
        </ol>

        {state.openQuestions.length > 0 && (
          <div className="flex flex-col gap-2" data-testid="open-questions">
            {state.openQuestions.map((q) => (
              <QuestionCard key={q.id} question={q} disabled={state.busy} onAnswer={(a) => void send({ message: "", answers: [a] })} />
            ))}
          </div>
        )}

        <div className="flex items-center gap-3 text-sm text-neutral-600" role="status">
          {state.busy && (
            <>
              <span>The agent is working...</span>
              <button type="button" onClick={stop} className="rounded border border-neutral-400 px-2 py-0.5">
                Stop
              </button>
            </>
          )}
        </div>

        {showComposer && (
          <form onSubmit={onSubmit} className="flex flex-col gap-2">
            <label className="flex flex-col gap-1">
              <span>Your message</span>
              <textarea
                value={message}
                onChange={(e) => setMessage(e.target.value)}
                onKeyDown={onKeyDown}
                rows={3}
                className="rounded border border-neutral-400 bg-white px-2 py-1 text-neutral-900"
                placeholder={beforeFirstMessage ? "Build me a mono-green Commander deck around elves." : ""}
              />
            </label>
            <p className={`text-xs ${tooLong ? "text-red-700" : "text-neutral-600"}`}>
              Enter sends, Shift+Enter makes a new line. {bytes} of {maxMessageBytes} bytes.
            </p>
            {tooLong && (
              <p role="alert" className="text-sm text-red-700">
                The message is over the {maxMessageBytes} byte cap. Shorten it.
              </p>
            )}
            <button
              type="submit"
              disabled={state.busy || tooLong || !message.trim()}
              className="self-start rounded bg-neutral-900 px-3 py-2 text-white disabled:bg-neutral-300 disabled:text-neutral-600"
            >
              Send
            </button>
          </form>
        )}
        <div ref={end} />

        {state.usage && state.usage.calls > 0 && (
          <p className="text-xs text-neutral-600" data-testid="usage">
            Session spend: {state.usage.calls} calls, {String(state.usage.inputTokens)} in, {String(state.usage.outputTokens)} out,{" "}
            {state.usage.priced ? `$${state.usage.costUsd.toFixed(4)}` : "cost unknown"} (M-1).
          </p>
        )}
      </section>

      <section aria-label="Deck" className="min-w-0 flex-1">
        {deckError && (
          <p role="alert" className="text-red-700">
            Could not load the deck: {deckError}
          </p>
        )}
        {state.deck ? (
          <DeckView deck={state.deck} />
        ) : (
          !deckError && <p className="text-neutral-600">The deck shows here when the agent has built one.</p>
        )}
      </section>
    </div>
  );
}

function poolLabel(rule: PoolRule | undefined, collectionId: string): string {
  switch (rule) {
    case PoolRule.OWNED_ONLY:
      return "Pool: only cards in your collection.";
    case PoolRule.OWNED_FIRST:
      return "Pool: your collection first, with upgrades to buy.";
    case PoolRule.ANY_CARD:
      return "Pool: any card (D-37).";
    default:
      return collectionId ? `Pool: your collection (${collectionId}). The agent asks how strict.` : "Pool: any card (D-37).";
  }
}

function ThreadLine({ item }: { item: ThreadItem }) {
  switch (item.kind) {
    case "user":
      return (
        <p className="ml-auto max-w-[85%] whitespace-pre-line rounded bg-neutral-100 px-3 py-2">
          <span className="sr-only">You: </span>
          {item.text}
        </p>
      );
    case "agent":
      return (
        <p className="whitespace-pre-line px-3 py-2">
          <span className="sr-only">Agent: </span>
          {item.text}
        </p>
      );
    case "status":
      return <p className="px-3 text-sm italic text-neutral-600">{item.text}</p>;
    case "question":
      return <p className="px-3 text-sm text-neutral-700">Asked: {item.question.text}</p>;
    case "failure":
      return (
        <p role="alert" className="rounded bg-red-50 px-3 py-2 text-red-700">
          {item.failure.message} ({item.failure.code}){item.failure.retryable ? " You can try again." : ""}
        </p>
      );
    case "deck":
      return <p className="px-3 text-sm text-neutral-700">Deck built: {item.deck.name || item.deck.id}.</p>;
  }
}
