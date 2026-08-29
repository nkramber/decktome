import { Code, ConnectError } from "@connectrpc/connect";
import { type Answer, PoolRule, type Session } from "@mtg/api-client/mtg/v1/session_pb";
import { useQuery } from "@tanstack/react-query";
import { type FormEvent, type KeyboardEvent, useCallback, useEffect, useRef, useState } from "react";
import { Link, useBlocker, useLocation, useNavigate, useParams } from "react-router";

import { agentClient, deckClient } from "../../lib/api";
import { errorMessage } from "../../lib/errors";
import { useAppStore } from "../../lib/store";
import { DeckView } from "../deck/deck-view";
import { type Draft, draftAnswered, emptyDraft, QuestionCard } from "./question-card";
import { byteLength, type ChatState, emptyState, fromSession, maxMessageBytes, type ThreadItem, useChat } from "./use-chat";

// The chat screen (ui plan, step 3). The id "new" means no session yet.
// A stored session loads through GetSession and its latest deck through
// GetDeck, so a reload returns to the thread. The deck before the latest
// loads too when the latest revised it, for the diff (PR-12B).
export function SessionPage() {
  const { id } = useParams();
  const location = useLocation();
  const isNew = !id || id === "new";
  const storedSessionId = useAppStore((s) => s.sessionId);
  const setSessionId = useAppStore((s) => s.setSessionId);

  // The panel this page started stays mounted when the route moves from
  // /session/new to /session/<id>: its key is the location key of the
  // /session/new visit, and the started id marks the route as its own.
  // A later visit to /session/new has a new location key, so it gets a
  // fresh panel, and the started id clears so a return to /session/<id>
  // loads the stored session.
  const [started, setStarted] = useState<{ id: string; key: string } | null>(null);
  const wasNew = useRef(isNew);
  useEffect(() => {
    if (isNew && !wasNew.current) setStarted(null);
    wasNew.current = isNew;
  }, [isNew]);
  const live = isNew || id === started?.id;
  const panelKey = isNew ? location.key : started && id === started.id ? started.key : id;
  const onStarted = useCallback((sid: string) => setStarted({ id: sid, key: location.key }), [location.key]);

  const session = useQuery({
    queryKey: ["session", id],
    queryFn: () => agentClient.getSession({ sessionId: id ?? "" }),
    enabled: !live,
  });
  const deckIds = session.data?.session?.deckIds ?? [];
  const deckId = deckIds.at(-1) ?? "";
  const deck = useQuery({
    queryKey: ["deck", deckId],
    queryFn: () => deckClient.getDeck({ deckId }),
    enabled: deckId !== "",
  });
  const baseId = deck.data?.deck?.revisedFromDeckId ? (deckIds.at(-2) ?? "") : "";
  const base = useQuery({
    queryKey: ["deck", baseId],
    queryFn: () => deckClient.getDeck({ deckId: baseId }),
    enabled: baseId !== "",
  });

  // A stored id that names a session the server no longer has, or one
  // of another user, is forgotten. Any other error keeps it for a retry.
  useEffect(() => {
    if (!session.isError || storedSessionId !== id) return;
    const code = ConnectError.from(session.error).code;
    if (code === Code.NotFound || code === Code.PermissionDenied) setSessionId("");
  }, [session.isError, session.error, storedSessionId, id, setSessionId]);

  if (live) {
    return <ChatPanel key={panelKey} initial={emptyState} onStarted={onStarted} resumeId={isNew && storedSessionId ? storedSessionId : ""} />;
  }
  if (session.isPending || (deckId && deck.isPending) || (baseId && base.isPending)) {
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
  const deckError = [deck.isError ? errorMessage(deck.error) : "", base.isError ? errorMessage(base.error) : ""].filter(Boolean).join(" ");
  return <ChatPanel key={panelKey} initial={fromSession(session.data.session, deck.data?.deck, base.data?.deck)} session={session.data.session} deckError={deckError} />;
}

// pruneDrafts keeps the drafts of the open questions only.
export function pruneDrafts(drafts: Record<string, Draft>, open: { id: string }[]): Record<string, Draft> {
  const ids = new Set(open.map((q) => q.id));
  return Object.fromEntries(Object.entries(drafts).filter(([id]) => ids.has(id)));
}

// The sentinel scrolls into view only while the reader is near the end
// of the thread, so new output does not pull them away from an earlier line.
const nearBottomPx = 240;

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
  const navigateRef = useRef(navigate);
  useEffect(() => {
    navigateRef.current = navigate;
  }, [navigate]);
  const collectionId = useAppStore((s) => s.collectionId);
  const poolMode = useAppStore((s) => s.poolMode);
  const setPoolMode = useAppStore((s) => s.setPoolMode);
  const setSessionId = useAppStore((s) => s.setSessionId);

  // The collection goes with the first message only. A stored session
  // holds its own collection id, and the server reads that one.
  const sendCollection = initial.sessionId === "" && poolMode === "owned" ? collectionId : (session?.collectionId ?? "");

  // ownPath is the route of this panel's session. The move from
  // /session/new to it is never blocked.
  const ownPath = useRef(initial.sessionId ? `/session/${initial.sessionId}` : "");
  const onSessionStarted = useCallback(
    (sid: string) => {
      setSessionId(sid);
      onStarted?.(sid);
      ownPath.current = `/session/${sid}`;
      navigateRef.current(`/session/${sid}`, { replace: true });
    },
    [setSessionId, onStarted],
  );
  const { state, send, stop } = useChat(initial, sendCollection, onSessionStarted);
  const [message, setMessage] = useState("");
  const [drafts, setDrafts] = useState<Record<string, Draft>>({});
  const bytes = byteLength(message);
  const tooLong = bytes > maxMessageBytes;
  const beforeFirstMessage = state.sessionId === "";
  // The message box hides while a question waits, and returns when the
  // turn ends with no question open (D-282). It stays mounted through a
  // send, disabled, so focus does not fall to the page body.
  const showComposer = state.openQuestions.length === 0;

  // A build runs on after the page leaves, and the deck shows on reload
  // (D-303). The page says so before a navigation or an unload mid-turn.
  const blocker = useBlocker(({ nextLocation }) => state.busy && nextLocation.pathname !== ownPath.current);
  useEffect(() => {
    if (!state.busy) return;
    const warn = (e: BeforeUnloadEvent) => {
      e.preventDefault();
    };
    window.addEventListener("beforeunload", warn);
    return () => window.removeEventListener("beforeunload", warn);
  }, [state.busy]);

  // Focus returns to the message box when a send ends with it enabled,
  // and moves to a question group when one appears.
  const textarea = useRef<HTMLTextAreaElement>(null);
  const questionsForm = useRef<HTMLFormElement>(null);
  const wasBusy = useRef(false);
  useEffect(() => {
    if (wasBusy.current && !state.busy && showComposer) textarea.current?.focus();
    wasBusy.current = state.busy;
  }, [state.busy, showComposer]);
  const openCount = state.openQuestions.length;
  const prevOpenCount = useRef(initial.openQuestions.length);
  useEffect(() => {
    if (openCount > prevOpenCount.current) questionsForm.current?.focus();
    prevOpenCount.current = openCount;
  }, [openCount]);

  // New output scrolls into view. The sentinel sits under the thread.
  const end = useRef<HTMLDivElement>(null);
  const last = state.thread[state.thread.length - 1];
  const lastLength = last && "text" in last ? last.text.length : 0;
  useEffect(() => {
    const el = end.current;
    if (!el) return;
    if (el.getBoundingClientRect().top - window.innerHeight > nearBottomPx) return;
    el.scrollIntoView?.({ block: "nearest" });
  }, [state.thread.length, lastLength, openCount]);

  // Every open question needs an answer before the submit, and one send
  // carries them all (D-282). Each answer text has the same cap as a message.
  const allAnswered = openCount > 0 && state.openQuestions.every((q) => draftAnswered(drafts[q.id]));
  const answerTooLong = state.openQuestions.some((q) => byteLength(drafts[q.id]?.text ?? "") > maxMessageBytes);

  function onSubmitAnswers(e: FormEvent) {
    e.preventDefault();
    if (!allAnswered || answerTooLong || state.busy) return;
    const answers = state.openQuestions.map((q) => {
      const d = drafts[q.id];
      return { questionId: q.id, optionIndex: d.text.trim() ? undefined : d.optionIndex, text: d.text.trim() } as Answer;
    });
    const sent = drafts;
    setDrafts({});
    void send({ message: "", answers }).then((r) => {
      // A failed or stopped send gives the drafts of the restored
      // questions back, so the user does not answer twice.
      if (r.restored.length > 0) setDrafts((d) => ({ ...pruneDrafts(sent, r.restored), ...d }));
    });
  }

  async function submit() {
    const text = message.trim();
    if (!text || tooLong || state.busy) return;
    setMessage("");
    const { ok } = await send({ message: text, answers: [] });
    // A failed send gives the draft back, so the user does not retype it.
    if (!ok) setMessage((m) => m || text);
  }

  function onSubmit(e: FormEvent) {
    e.preventDefault();
    void submit();
  }

  // Enter sends, and Shift+Enter makes a new line. Enter during an IME
  // composition commits the composition and sends nothing.
  function onKeyDown(e: KeyboardEvent<HTMLTextAreaElement>) {
    if (e.key === "Enter" && !e.shiftKey && !e.nativeEvent.isComposing) {
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

        {blocker.state === "blocked" && (
          <div role="alertdialog" aria-labelledby="leave-title" aria-describedby="leave-text" className="flex flex-col gap-2 rounded border border-amber-400 bg-amber-50 p-3">
            <p id="leave-title" className="font-medium">
              The agent is still working.
            </p>
            <p id="leave-text" className="text-sm">
              The build continues on the server. Open this session again, and the deck shows when it is done.
            </p>
            <div className="flex gap-2">
              <button type="button" onClick={() => blocker.reset()} className="rounded bg-neutral-900 px-3 py-1 text-white">
                Stay
              </button>
              <button type="button" onClick={() => blocker.proceed()} className="rounded border border-neutral-400 px-3 py-1">
                Leave
              </button>
            </div>
          </div>
        )}

        <ol className="flex flex-col gap-2" aria-label="Conversation">
          {state.thread.map((item) => (
            <li key={item.id}>
              <ThreadLine item={item} />
            </li>
          ))}
        </ol>

        {openCount > 0 && (
          <form ref={questionsForm} tabIndex={-1} onSubmit={onSubmitAnswers} className="flex flex-col gap-2" data-testid="open-questions">
            {state.openQuestions.map((q) => (
              <QuestionCard
                key={q.id}
                question={q}
                draft={drafts[q.id] ?? emptyDraft}
                disabled={state.busy}
                onChange={(d) => setDrafts((all) => pruneDrafts({ ...all, [q.id]: d }, state.openQuestions))}
              />
            ))}
            {answerTooLong && (
              <p role="alert" className="text-sm text-red-700">
                An answer is over the {maxMessageBytes} byte cap. Shorten it.
              </p>
            )}
            <button
              type="submit"
              disabled={!allAnswered || answerTooLong || state.busy}
              className="self-start rounded bg-neutral-900 px-3 py-2 text-white disabled:bg-neutral-300 disabled:text-neutral-600"
            >
              Submit answers
            </button>
            {!allAnswered && !state.busy && <p className="text-xs text-neutral-600">Answer every question, then submit.</p>}
          </form>
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
                ref={textarea}
                value={message}
                disabled={state.busy}
                onChange={(e) => setMessage(e.target.value)}
                onKeyDown={onKeyDown}
                rows={3}
                className="rounded border border-neutral-400 bg-white px-2 py-1 text-neutral-900 disabled:bg-neutral-100 disabled:text-neutral-500"
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
          <DeckView deck={state.deck} base={state.baseDeck} />
        ) : (
          !deckError && <p className="text-neutral-600">The deck shows here when the agent has built one.</p>
        )}
      </section>
    </div>
  );
}

export function poolLabel(rule: PoolRule | undefined, collectionId: string): string {
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
