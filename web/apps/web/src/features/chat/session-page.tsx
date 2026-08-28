import { PoolRule, type Session } from "@mtg/api-client/mtg/v1/session_pb";
import { useQuery } from "@tanstack/react-query";
import { type FormEvent, useCallback, useState } from "react";
import { useNavigate, useParams } from "react-router";

import { agentClient, deckClient } from "../../lib/api";
import { errorMessage } from "../../lib/errors";
import { useAppStore } from "../../lib/store";
import { DeckView } from "../deck/deck-view";
import { QuestionCard } from "./question-card";
import { byteLength, type ChatState, emptyState, fromSession, maxMessageBytes, useChat } from "./use-chat";

// The chat screen (ui plan, step 3). The id "new" means no session yet.
// A stored session loads through GetSession and its latest deck through
// GetDeck, so a reload returns to the thread.
export function SessionPage() {
  const { id } = useParams();
  const isNew = !id || id === "new";
  // A session this panel started stays mounted when the route moves from
  // /session/new to /session/<id>. Nothing reloads it, so the stream and
  // the thread survive the navigate.
  const [started, setStarted] = useState("");
  const live = isNew || id === started;

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
    return <ChatPanel key="live" initial={emptyState} onStarted={setStarted} />;
  }
  if (session.isPending || (deckId && deck.isPending)) {
    return (
      <div className="p-6 text-neutral-600" aria-live="polite">
        <p data-testid="session-id">Session id: {id}</p>
        <p>Loading the session...</p>
      </div>
    );
  }
  if (session.isError || !session.data.session) {
    return (
      <p role="alert" className="p-6 text-red-700">
        Could not load the session: {session.isError ? errorMessage(session.error) : "the server returned no session"}
      </p>
    );
  }
  return <ChatPanel key={id} initial={fromSession(session.data.session, deck.data?.deck)} session={session.data.session} />;
}

function ChatPanel({ initial, session, onStarted }: { initial: ChatState; session?: Session; onStarted?: (id: string) => void }) {
  const navigate = useNavigate();
  const collectionId = useAppStore((s) => s.collectionId);
  const poolMode = useAppStore((s) => s.poolMode);
  const setPoolMode = useAppStore((s) => s.setPoolMode);
  const setSessionId = useAppStore((s) => s.setSessionId);
  const isNew = initial.sessionId === "";

  // The collection goes with the first message only. A stored session
  // holds its own collection id, and the server reads that one.
  const sendCollection = isNew && poolMode === "owned" ? collectionId : (session?.collectionId ?? "");
  const [message, setMessage] = useState("");

  const onSessionStarted = useCallback(
    (sid: string) => {
      setSessionId(sid);
      onStarted?.(sid);
      navigate(`/session/${sid}`, { replace: true });
    },
    [navigate, setSessionId, onStarted],
  );
  const { state, send } = useChat(initial, sendCollection, onSessionStarted);
  const beforeFirstMessage = state.sessionId === "";
  const tooLong = byteLength(message) > maxMessageBytes;

  function onSubmit(e: FormEvent) {
    e.preventDefault();
    const text = message.trim();
    if (!text || tooLong || state.busy) return;
    setMessage("");
    void send({ message: text, answers: [] });
  }

  const poolText = poolLabel(state.slots?.poolRule, sendCollection);

  return (
    <div className="mx-auto flex max-w-7xl flex-col gap-4 p-6 lg:flex-row">
      <section aria-labelledby="chat-title" className="flex min-w-0 flex-1 flex-col gap-3 lg:max-w-xl">
        <h1 id="chat-title" className="text-2xl font-semibold">
          Chat
        </h1>
        <p className="text-sm text-neutral-600" data-testid="session-id">
          {beforeFirstMessage ? "No session yet." : `Session id: ${state.sessionId}`}
        </p>
        {beforeFirstMessage && collectionId ? (
          <label className="flex items-center gap-2 text-sm">
            <input type="checkbox" checked={poolMode === "owned"} onChange={(e) => setPoolMode(e.target.checked ? "owned" : "any")} />
            Use only cards in my collection
          </label>
        ) : null}
        <p className="text-sm text-neutral-600" data-testid="pool-mode">
          {poolText}
        </p>

        <ol className="flex flex-col gap-2" aria-label="Conversation">
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

        <div aria-live="polite" className="text-sm text-neutral-600">
          {state.busy && <p>The agent is working...</p>}
        </div>

        <form onSubmit={onSubmit} className="flex flex-col gap-2">
          <label className="flex flex-col gap-1">
            <span>Your message</span>
            <textarea
              value={message}
              onChange={(e) => setMessage(e.target.value)}
              rows={3}
              className="rounded border border-neutral-400 px-2 py-1"
              placeholder={beforeFirstMessage ? "Build me a mono-green Commander deck around elves." : ""}
            />
          </label>
          {tooLong && (
            <p role="alert" className="text-sm text-red-700">
              The message is over the {maxMessageBytes} byte cap. Shorten it.
            </p>
          )}
          <button type="submit" disabled={state.busy || tooLong || !message.trim()} className="self-start rounded bg-neutral-900 px-3 py-2 text-white disabled:opacity-50">
            Send
          </button>
        </form>

        {state.usage && state.usage.calls > 0 && (
          <p className="text-xs text-neutral-600" data-testid="usage">
            Session spend: {state.usage.calls} calls, {String(state.usage.inputTokens)} in, {String(state.usage.outputTokens)} out,{" "}
            {state.usage.priced ? `$${state.usage.costUsd.toFixed(4)}` : "cost unknown"} (M-1).
          </p>
        )}
      </section>

      <section aria-label="Deck" className="min-w-0 flex-1">
        {state.deck ? <DeckView deck={state.deck} /> : <p className="text-neutral-600">The deck shows here when the agent has built one.</p>}
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

function ThreadLine({ item }: { item: ReturnType<typeof fromSession>["thread"][number] }) {
  switch (item.kind) {
    case "user":
      return <p className="whitespace-pre-line rounded bg-neutral-100 px-3 py-2">{item.text}</p>;
    case "agent":
      return <p className="whitespace-pre-line px-3 py-2">{item.text}</p>;
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

