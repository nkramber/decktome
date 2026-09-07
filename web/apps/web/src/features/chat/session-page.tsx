import { Code, ConnectError } from "@connectrpc/connect";
import type { Deck } from "@mtg/api-client/mtg/v1/deck_pb";
import { FeedbackKind } from "@mtg/api-client/mtg/v1/feedback_service_pb";
import { type Answer, PoolRule, type Session } from "@mtg/api-client/mtg/v1/session_pb";
import { useQuery, useQueryClient } from "@tanstack/react-query";
import { type FormEvent, type KeyboardEvent, type ReactNode, useCallback, useEffect, useRef, useState } from "react";
import { Link, useBlocker, useLocation, useNavigate, useParams } from "react-router";

import { AlertTriangleIcon, ArrowUpIcon, CheckIcon, LayersIcon } from "lucide-react";

import { Button } from "../../components/ui/button";
import { Label } from "../../components/ui/label";
import { Textarea } from "../../components/ui/textarea";
import { agentClient, deckClient } from "../../lib/api";
import { cn } from "../../lib/cn";
import { errorMessage } from "../../lib/errors";
import { type PoolMode, useAppStore } from "../../lib/store";
import { DeckView } from "../deck/deck-view";
import { PoolPicker, useCollections } from "./pool-picker";
import { BuildStepper } from "./build-stepper";
import { RecentDecks, useRecentDecks } from "./recent-decks";
import { UnfinishedChats } from "./unfinished-chats";
import { Thumbs } from "../feedback/thumbs";
import { type Draft, draftAnswered, emptyDraft, QuestionCard } from "./question-card";
import { byteLength, type ChatState, emptyState, fromSession, maxMessageBytes, type SendInput, type ThreadItem, useChat } from "./use-chat";

// The chat screen (ui plan, step 3). The id "new" means no session yet.
// A stored session loads through GetSession and its latest deck through
// GetDeck, so a reload returns to the thread. The deck before the latest
// loads too when the latest revised it, for the diff (PR-12B).
export function SessionPage() {
  const { id } = useParams();
  const navigate = useNavigate();
  const location = useLocation();
  const queryClient = useQueryClient();
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
  // newKey is the location key of the current /session/new visit. A
  // replace of the location state keeps the visit, so the panel must
  // not remount on it: a remount aborts a send that already left.
  const [newKey, setNewKey] = useState(location.key);
  const wasNew = useRef(isNew);
  useEffect(() => {
    if (isNew && !wasNew.current) {
      setStarted(null);
      setNewKey(location.key);
    }
    wasNew.current = isNew;
  }, [isNew, location.key]);
  const live = isNew || id === started?.id;
  // A reload of the session, after a lost stream (PR-19). The panel
  // leaves the live state and reads the stored session again.
  const [reloads, setReloads] = useState(0);
  const panelKey = isNew ? newKey : started && id === started.id ? started.key : `${id}:${reloads}`;
  const onResume = useCallback(() => {
    setStarted(null);
    setReloads((n) => n + 1);
    void queryClient.invalidateQueries({ queryKey: ["session", id] });
  }, [queryClient, id]);
  const onStarted = useCallback((sid: string) => setStarted({ id: sid, key: newKey }), [newKey]);

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

  const toDeck = useCallback((deckId: string) => void navigate(`/decks/${deckId}`, { replace: true }), [navigate]);

  if (live) {
    return <ChatPanel key={panelKey} initial={emptyState} onStarted={onStarted} onDeckBuilt={toDeck} onResume={isNew ? undefined : onResume} />;
  }
  if (session.isPending || (deckId && deck.isPending) || (baseId && base.isPending)) {
    return (
      <div className="flex flex-col gap-2 p-4 md:p-6">
        <h1 className="text-2xl font-semibold tracking-tight">Chat</h1>
        <p className="wrap-anywhere text-sm text-muted-foreground" data-testid="session-id">
          Session id: {id}
        </p>
        <p role="status" className="text-muted-foreground">
          Loading the session...
        </p>
      </div>
    );
  }
  if (session.isError || !session.data.session) {
    return (
      <div className="flex flex-col items-start gap-2 p-4 md:p-6">
        <h1 className="text-2xl font-semibold tracking-tight">Chat</h1>
        <p className="wrap-anywhere text-sm text-muted-foreground" data-testid="session-id">
          Session id: {id}
        </p>
        <p role="alert" className="text-danger">
          Could not load the session: {session.isError ? errorMessage(session.error) : "the server returned no session"}
        </p>
        <Button asChild variant="outline">
          <Link to="/session/new">Start a new chat</Link>
        </Button>
      </div>
    );
  }
  const deckError = [deck.isError ? errorMessage(deck.error) : "", base.isError ? errorMessage(base.error) : ""].filter(Boolean).join(" ");
  return (
    <ChatPanel
      key={panelKey}
      initial={fromSession(session.data.session, deck.data?.deck, base.data?.deck)}
      session={session.data.session}
      deckError={deckError}
      onDeckBuilt={toDeck}
      onResume={onResume}
    />
  );
}

// pruneDrafts keeps the drafts of the open questions only.
export function pruneDrafts(drafts: Record<string, Draft>, open: { id: string }[]): Record<string, Draft> {
  const ids = new Set(open.map((q) => q.id));
  return Object.fromEntries(Object.entries(drafts).filter(([id]) => ids.has(id)));
}

// The reader's own last message is the top of what they need to read
// (D-360). Everything the turn produced lands under it: the agent's
// prose, the questions it asked, and the button that sends the answers.
// The scroll therefore puts that message at the top of the frame rather
// than chasing the foot of the thread. A turn whose output fits shows
// the message and the submit together, and one that does not keeps the
// message, which is the half the reader needs.
//
// The sentinel below still serves a thread with no message of the user
// in it yet.
const nearBottomPx = 240;

// wheelFactor slows the thread against the wheel (D-370). One notch of a
// mouse wheel moves a browser about 100 px, which is a large step in a
// column this narrow: a whole turn goes by in one notch. The thread takes
// a fraction of that instead.
const wheelFactor = 0.4;

// ChatPanel serves both screens of a deck (D-335). On the session route
// it shows the conversation alone, and it hands over to the deck's own
// address the moment a deck exists. On the deck route it shows the deck
// the address names, with the actions the user owns, and its dock.
export function ChatPanel({
  initial,
  session,
  onStarted,
  deckError = "",
  deckOverride,
  baseOverride,
  actions,
  onDeckBuilt,
  onResume,
}: {
  initial: ChatState;
  session?: Session;
  onStarted?: (id: string) => void;
  deckError?: string;
  deckOverride?: Deck;
  baseOverride?: Deck;
  actions?: ReactNode;
  onDeckBuilt?: (deckId: string) => void;
  // onResume reads the stored session again after a lost stream.
  onResume?: () => void;
}) {
  const navigate = useNavigate();
  const navigateRef = useRef(navigate);
  useEffect(() => {
    navigateRef.current = navigate;
  }, [navigate]);
  const collectionId = useAppStore((s) => s.collectionId);
  const poolMode = useAppStore((s) => s.poolMode);
  const setSessionId = useAppStore((s) => s.setSessionId);

  // The collection goes with the first message only. A stored session
  // holds its own collection id, and the server reads that one.
  // The collection goes with the first message whenever the reader named
  // one. Unchecking "Only cards I own" no longer drops it: it says the
  // collection leads and the database fills a gap (D-359). A stored
  // session holds its own collection, and the server reads that one.
  const sendCollection = initial.sessionId === "" ? collectionId : (session?.collectionId ?? "");
  const sendPoolRule = initial.sessionId === "" ? poolRuleOf(poolMode, collectionId) : PoolRule.UNSPECIFIED;

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
  const { state, send, stop } = useChat(initial, sendCollection, sendPoolRule, onSessionStarted);
  const [message, setMessage] = useState("");
  const [drafts, setDrafts] = useState<Record<string, Draft>>({});
  const bytes = byteLength(message);
  const tooLong = bytes > maxMessageBytes;
  const beforeFirstMessage = state.sessionId === "";
  // The message box hides while a question waits, and returns when the
  // turn ends with no question open (D-282). It also leaves while the
  // agent works: a disabled box beside "the agent is working" reads as a
  // dead control (PR-16B). The working row and its Stop take its place.
  const showComposer = state.openQuestions.length === 0 && !state.busy;

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

  // The thread scrolls at its own pace (D-370). The handler is native and
  // not passive, because a passive listener may not stop the browser from
  // scrolling its own distance first.
  const threadScroll = useRef<HTMLDivElement>(null);
  useEffect(() => {
    const el = threadScroll.current;
    if (!el) return;
    function onWheel(e: WheelEvent) {
      const box = threadScroll.current;
      if (!box || e.ctrlKey) return;
      // A wheel at the end of the thread belongs to whatever is under it.
      const atTop = box.scrollTop <= 0 && e.deltaY < 0;
      const atEnd = Math.ceil(box.scrollTop + box.clientHeight) >= box.scrollHeight && e.deltaY > 0;
      if (atTop || atEnd) return;
      e.preventDefault();
      // deltaMode 1 counts lines, and 2 counts pages. Both are rare, and
      // a line is about 16 px.
      const step = e.deltaMode === 1 ? e.deltaY * 16 : e.deltaY;
      box.scrollTop += step * wheelFactor;
    }
    el.addEventListener("wheel", onWheel, { passive: false });
    return () => el.removeEventListener("wheel", onWheel);
  }, []);

  // New output scrolls into view. The sentinel sits under the thread.
  const end = useRef<HTMLDivElement>(null);
  const lastSubmission = useRef<HTMLLIElement>(null);
  const last = state.thread[state.thread.length - 1];
  const lastLength = last && "text" in last ? last.text.length : 0;
  useEffect(() => {
    const mine = lastSubmission.current;
    if (mine) {
      // scrollIntoView scrolls every scrollable ancestor, and an
      // overflow-hidden box still scrolls when code asks it to. On a deck
      // screen that pushed the whole chat column off the top of the
      // frame. Inside its own box the thread scrolls itself, and nothing
      // above it moves (D-370).
      const box = threadScroll.current;
      if (box?.contains(mine)) {
        box.scrollTop += mine.getBoundingClientRect().top - box.getBoundingClientRect().top;
        return;
      }
      mine.scrollIntoView?.({ block: "start" });
      return;
    }
    const el = end.current;
    if (!el) return;
    const box = threadScroll.current;
    if (box?.contains(el)) {
      box.scrollTop = box.scrollHeight;
      return;
    }
    if (el.getBoundingClientRect().top - window.innerHeight > nearBottomPx) return;
    el.scrollIntoView?.({ block: "nearest" });
  }, [state.thread.length, lastLength, openCount, state.busy]);

  // Every open question needs an answer before the submit, and one send
  // carries them all (D-282). Each answer text has the same cap as a message.
  const allAnswered = openCount > 0 && state.openQuestions.every((q) => draftAnswered(drafts[q.id]));
  const answerTooLong = state.openQuestions.some((q) => byteLength(drafts[q.id]?.text ?? "") > maxMessageBytes);

  function onSubmitAnswers(e: FormEvent) {
    e.preventDefault();
    if (!allAnswered || answerTooLong || state.busy) return;
    const answers = state.openQuestions.map((q) => {
      const d = drafts[q.id];
      // A decline carries no value on purpose (D-353).
      if (d.declined) return { questionId: q.id, declined: true, text: "" } as Answer;
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

  // The pool line names the collection, not its id. The picker and the
  // header menu read the same list, so one call serves all three.
  const collections = useCollections(sendCollection !== "");
  const collectionName = collections.data?.collections.find((c) => c.id === sendCollection)?.name ?? sendCollection;
  const poolText = poolLabel(state.slots?.poolRule, sendCollection ? collectionName : "");
  // The newest decks head a new chat (D-350). The title names them when
  // there are any, so the page opens on the reader's own work.
  const recent = useRecentDecks(beforeFirstMessage);
  // The deck owns the page once one exists, and the conversation docks
  // at the corner (D-331). Before that, the conversation is the page.
  // The address of a deck names which deck shows. A build that ends with
  // a new deck tells the page, and the page moves to that address.
  const builtDeck = deckOverride ?? state.deck;
  const streamedDeckId = state.deck?.id;
  const queryClient = useQueryClient();
  const builtSessionId = state.sessionId;
  useEffect(() => {
    if (!streamedDeckId || state.busy) return;
    // The deck screen reads the session from the cache, and the cache
    // may hold the session from before this turn, with the last question
    // still open (D-446). The turn that built the deck stales it first.
    void queryClient.invalidateQueries({ queryKey: ["session", builtSessionId] });
    if (streamedDeckId !== deckOverride?.id) onDeckBuilt?.(streamedDeckId);
  }, [streamedDeckId, state.busy, deckOverride?.id, onDeckBuilt, queryClient, builtSessionId]);
  // A question sits in the thread and in the open list at the same time,
  // and the card below it takes the answer. The thread holds the line for
  // the history, so it shows the question only after it is answered.
  const openIds = new Set(state.openQuestions.map((q) => q.id));
  const shown = state.thread.filter((item) => item.kind !== "question" || !openIds.has(item.question.id));
  // The reader's last message, which the scroll holds at the top (D-360).
  const lastMineId = shown.reduce((id, item) => (item.kind === "user" ? item.id : id), -1);
  const thread = (
    <ol className="flex flex-col gap-5" aria-label="Conversation">
      {shown.map((item) => (
        <li key={item.id} ref={item.id === lastMineId ? lastSubmission : undefined} className="scroll-mt-4">
          <ThreadLine item={item} sessionId={state.sessionId} />
        </li>
      ))}
    </ol>
  );

  const questions = openCount > 0 && (
    <form ref={questionsForm} tabIndex={-1} onSubmit={onSubmitAnswers} className="flex flex-col gap-2" data-testid="open-questions">
      {state.openQuestions.map((q) => (
        <QuestionCard
          key={q.id}
          question={q}
          draft={drafts[q.id] ?? emptyDraft}
          disabled={state.busy}
          onChange={(d) => setDrafts((all) => pruneDrafts({ ...all, [q.id]: d }, state.openQuestions))}
          sessionId={state.sessionId}
        />
      ))}
      {answerTooLong && (
        <p role="alert" className="text-sm text-danger">
          One of your answers is too long. Shorten it.
        </p>
      )}
      <Button type="submit" className="self-start" disabled={!allAnswered || answerTooLong || state.busy}>
        Submit answers
      </Button>
      {!allAnswered && !state.busy && <p className="text-xs text-muted-foreground">Answer every question, then submit.</p>}
    </form>
  );

  // What the agent is doing right now, in its own words. A build runs for
  // minutes, and "the agent is working" says nothing about which minute
  // this is. The newest status line carries that, so the working row says
  // it rather than a line of its own (D-375). The stepper under it lights
  // the step the server named (D-435).
  const step = state.busy ? (shown.reduce((text, item) => (item.kind === "status" ? item.text : text), "") ?? "") : "";
  const working = (
    <div className="flex flex-col gap-2" role="status">
      {state.busy && (
        <>
          <div className="flex flex-wrap items-center gap-3">
            <span className="flex items-center gap-2.5">
              <span className="flex gap-1" aria-hidden="true">
                <span className="size-2 animate-bounce rounded-full bg-primary [animation-delay:-0.3s]" />
                <span className="size-2 animate-bounce rounded-full bg-primary [animation-delay:-0.15s]" />
                <span className="size-2 animate-bounce rounded-full bg-primary" />
              </span>
              <span className="font-display text-[15px] font-semibold text-foreground">{sentence(step) || "The agent is working..."}</span>
            </span>
            <Button type="button" variant="outline" size="sm" onClick={stop}>
              Stop
            </Button>
          </div>
          <BuildStepper phase={state.phase} repaired={state.repaired} />
        </>
      )}
    </div>
  );

  // A failed turn ends with the reason and a way on (PR-19). A retryable
  // failure runs the same send again. A lost stream reads the stored
  // session again, so the turn the server finished shows.
  const lastFailure = !state.busy && last?.kind === "failure" ? last.failure : undefined;
  const recovery = lastFailure && (
    <div className="flex flex-wrap gap-2" data-testid="recovery">
      {lastFailure.retryable && state.lastInput && (
        <Button type="button" size="sm" onClick={() => void send(state.lastInput as SendInput)}>
          Try again
        </Button>
      )}
      {state.sessionId !== "" && onResume && (
        <Button type="button" size="sm" variant="outline" onClick={onResume}>
          Reload the session
        </Button>
      )}
    </div>
  );

  const composer = showComposer && (
    // The cards sit close under the title they belong to, and the message
    // box stands apart as a section of its own (D-369).
    // The gap over the box equals the gap under the top bar (D-442): the
    // section gap plus this margin is the page's top padding, at each
    // width.
    <form onSubmit={onSubmit} className={cn("flex flex-col gap-2", beforeFirstMessage && "mt-4 md:mt-10")}>
      {beforeFirstMessage && (
        <h2 className="font-display text-2xl font-semibold">Build a new deck</h2>
      )}
      <div className="flex flex-col rounded-card border border-border bg-card transition-colors focus-within:border-primary">
        <Label htmlFor="message" className="sr-only">
          Your message
        </Label>
        <Textarea
          id="message"
          ref={textarea}
          value={message}
          onChange={(e) => setMessage(e.target.value)}
          onKeyDown={onKeyDown}
          rows={2}
          placeholder={beforeFirstMessage ? "Build me a mono-green Commander deck around elves." : "Ask for a change."}
          className="max-h-40 resize-none border-0 bg-transparent px-4 pt-3 pb-1"
        />
        <div className="flex flex-wrap items-end justify-between gap-x-3 gap-y-2 px-3 pb-3">
          <div className="flex min-w-0 flex-col gap-2">{beforeFirstMessage && <PoolPicker />}</div>
          <Button type="submit" size="icon" aria-label="Send" className="size-8" disabled={tooLong || !message.trim()}>
            <ArrowUpIcon aria-hidden="true" />
          </Button>
        </div>
      </div>
      {tooLong && (
        <p role="alert" className="text-sm text-danger">
          That message is too long. Shorten it.
        </p>
      )}
    </form>
  );

  const leaveWarning = blocker.state === "blocked" && (
    <div role="alertdialog" aria-labelledby="leave-title" aria-describedby="leave-text" className="flex flex-col gap-2 rounded-card border border-warning/50 bg-warning/10 p-3">
      <p id="leave-title" className="font-display font-semibold">
        The agent is still working.
      </p>
      <p id="leave-text" className="text-sm">
        The build continues on the server. Open this session again, and the deck shows when it is done.
      </p>
      <div className="flex gap-2">
        <Button size="sm" onClick={() => blocker.reset()}>
          Stay
        </Button>
        <Button size="sm" variant="outline" onClick={() => blocker.proceed()}>
          Leave
        </Button>
      </div>
    </div>
  );

  const idLine = (
    <div className="flex flex-wrap items-center gap-x-2 gap-y-1 font-mono text-[11px] text-muted-foreground">
      <span className="wrap-anywhere" data-testid="session-id">
        {beforeFirstMessage ? "No session yet." : `Session id: ${state.sessionId}`}
      </span>
      <span aria-hidden="true">·</span>
      <span className="wrap-anywhere" data-testid="pool-mode">
        {poolText}
      </span>
      {state.usage && state.usage.calls > 0 && (
        <>
          <span aria-hidden="true">·</span>
          <span data-testid="usage">
            Session spend: {state.usage.calls} calls, {String(state.usage.inputTokens)} in, {String(state.usage.outputTokens)} out,{" "}
            {state.usage.priced ? `$${state.usage.costUsd.toFixed(4)}` : "cost unknown"}
          </span>
        </>
      )}
    </div>
  );

  // The deck owns the page, and the conversation is a pinned column at
  // its left (D-346). The gap between the two equals the gap between the
  // deck and the right edge of the window, so the deck sits in an even
  // frame. The column overlaps nothing.
  if (builtDeck) {
    return (
      <div className="flex flex-col gap-6 p-4 md:p-6 lg:h-full lg:flex-row lg:overflow-hidden">
        {/* The chat is a column of its own, and it always fits the frame
            (D-364). The page scrolls under it, so it sticks to the top of
            the scrolling area and its own thread scrolls inside it. */}
        <aside
          aria-labelledby="chat-title"
          className="flex w-full shrink-0 flex-col gap-2 rounded-card border border-border bg-card p-3 lg:h-full lg:w-[26.62rem] print:hidden"
        >
          <h1 id="chat-title" className="font-display text-[10px] tracking-[0.15em] text-muted-foreground uppercase">
            Chat
          </h1>
          <div className="flex min-h-0 grow flex-col gap-2">
            <div ref={threadScroll} className="min-h-0 grow overflow-y-auto overscroll-contain">
              {thread}
              <div ref={end} />
            </div>
            {working}
            {recovery}
            {leaveWarning}
            {questions}
            {composer}
            {idLine}
          </div>
        </aside>

        {/* The deck column carries the scroll, so the chat beside it holds
            its place and always fits the frame (D-364). */}
        <div className="min-w-0 grow lg:h-full lg:overflow-y-auto">
          {deckError && (
            <p role="alert" className="mb-4 text-danger">
              Could not load the deck: {deckError}
            </p>
          )}
          {actions}
          <section aria-label="Deck">
            <DeckView deck={builtDeck} base={deckOverride ? baseOverride : state.baseDeck} />
          </section>
        </div>
      </div>
    );
  }

  // Before a deck exists the conversation is the page, so it fills the
  // frame and its composer sits at the foot (D-331). A chat with no
  // session yet is two blocks instead: the decks the reader already has,
  // and the box that builds the next one. Stretching that to the
  // viewport put a void between them (D-358).
  return (
    <div className={cn("mx-auto flex w-full max-w-4xl flex-col p-4 pt-8 md:p-6 md:pt-14", !beforeFirstMessage && "min-h-[calc(100vh-9rem)]")}>
      <section aria-labelledby="chat-title" className={cn("flex min-w-0 flex-col", beforeFirstMessage ? "gap-4" : "grow gap-4")}>
        {/* The identifiers are for support, not for reading. They sit in
            one quiet row under the title, and never in the thread. */}
        {/* A new chat says nothing of its session: it has none, and the
            picker in the message box names the pool (D-356). */}
        {/* A new chat with no deck to pick up shows no title: the message
            box says what the page is, and a heading over nothing is a
            second heading for the same box. The page keeps its first
            heading for a screen reader, out of sight. */}
        {!beforeFirstMessage || recent.decks.length > 0 ? (
          <div className="flex flex-col gap-1.5">
            <h1 id="chat-title" className="font-display text-2xl font-semibold">
              {!beforeFirstMessage ? "Chat" : "Pick up where you left off"}
            </h1>
            {!beforeFirstMessage && idLine}
          </div>
        ) : (
          <h1 id="chat-title" className="sr-only">
            New deck
          </h1>
        )}
        {beforeFirstMessage && <RecentDecks decks={recent.decks} isPending={recent.isPending} />}

        {leaveWarning}

        {/* The thread block holds nothing at all on a new chat, and an
            empty block still takes its gaps. It stays away until there
            is something to read (D-358). */}
        {(!beforeFirstMessage || state.thread.length > 0 || openCount > 0 || state.busy) && (
          <div className={cn("flex flex-col gap-5", !beforeFirstMessage && "grow")}>
            {thread}
            {questions}
            {working}
            {recovery}
            <div ref={end} />
          </div>
        )}

        {composer}
        {/* The unfinished chats sit under the message box, and only when
            there is one (D-438). */}
        {beforeFirstMessage && <UnfinishedChats />}
      </section>
      {deckError && (
        <p role="alert" className="text-danger">
          Could not load the deck: {deckError}
        </p>
      )}
    </div>
  );
}

// poolRuleOf maps the reader's choice onto the contract (D-359). With no
// collection there is nothing to prefer, so the rule stays unset and the
// agent asks nothing about a pool the user does not have.
export function poolRuleOf(mode: PoolMode, collectionId: string): PoolRule {
  if (collectionId === "") return PoolRule.UNSPECIFIED;
  return mode === "owned_only" ? PoolRule.OWNED_ONLY : PoolRule.OWNED_FIRST;
}

// sentence gives a status line a capital and a full stop of its own. The
// server writes them in lower case, for a line in a thread.
export function sentence(text: string): string {
  const t = text.trim();
  if (t === "") return "";
  return t[0].toUpperCase() + t.slice(1) + "...";
}

export function poolLabel(rule: PoolRule | undefined, collection: string): string {
  switch (rule) {
    case PoolRule.OWNED_ONLY:
      return "Pool: only cards in your collection";
    case PoolRule.OWNED_FIRST:
      return "Pool: your collection first, and the whole card database fills a gap";
    case PoolRule.ANY_CARD:
      return "Pool: any card";
    default:
      return collection ? `Pool: ${collection}, and the agent asks how strict` : "Pool: any card";
  }
}

// One line of the thread. A turn of the user reads as a block on its own
// surface, and a turn of the agent reads as plain text at a comfortable
// measure. The rest are notes about the turn, and they stay quiet
// (PR-16B). A bubble on both sides reads as a messenger, and this is a
// tool.
// The mark of the agent, the same four-point star the header carries.
function SparkMark() {
  return (
    <svg width="11" height="11" viewBox="0 0 14 14" fill="currentColor" aria-hidden="true">
      <path d="M7 0l1.7 5.3L14 7l-5.3 1.7L7 14l-1.7-5.3L0 7l5.3-1.7z" />
    </svg>
  );
}

function ThreadLine({ item, sessionId }: { item: ThreadItem; sessionId: string }) {
  switch (item.kind) {
    case "user":
      return (
        <p className="ml-auto w-fit max-w-[85%] rounded-card bg-accent px-4 py-2.5 text-accent-foreground whitespace-pre-line">
          <span className="sr-only">You: </span>
          {item.text}
        </p>
      );
    case "agent":
      return (
        <div className="flex items-start gap-2.5">
          <span aria-hidden="true" className="mt-5 grid size-7 shrink-0 place-items-center rounded-full bg-accent text-accent-foreground">
            <SparkMark />
          </span>
          <span className="min-w-0">
            <span className="font-display mb-1.5 block text-[10px] tracking-[0.15em] text-muted-foreground uppercase" aria-hidden="true">
              Agent
            </span>
            <p className="max-w-measure rounded-card border border-border bg-card px-4 py-3 leading-relaxed whitespace-pre-line">
              <span className="sr-only">Agent: </span>
              {item.text}
            </p>
          </span>
        </div>
      );
    case "status":
      return (
        <p className="flex items-center gap-2 text-sm text-muted-foreground">
          <CheckIcon className="size-3.5 shrink-0" aria-hidden="true" />
          {item.text}
        </p>
      );
    case "question":
      // An answered question keeps its thumbs, so a reader can judge it
      // after the fact (PR-27).
      return (
        <div className="flex flex-col gap-1 border-l-2 border-accent/40 pl-3 text-sm text-muted-foreground">
          <p>
            <span className="sr-only">Asked: </span>
            {item.question.text}
          </p>
          {sessionId && <Thumbs target={{ kind: FeedbackKind.QUESTION, sessionId, questionId: item.question.id }} itemName={`the question "${item.question.text}"`} name="Rate this question" />}
        </div>
      );
    case "failure":
      return (
        <p role="alert" className="flex items-start gap-2 rounded-card border border-danger/30 bg-danger/10 px-3 py-2.5 text-sm text-danger">
          <AlertTriangleIcon className="mt-0.5 size-4 shrink-0" aria-hidden="true" />
          <span>
            {item.failure.message} ({item.failure.code}){item.failure.retryable ? " You can try again." : ""}
          </span>
        </p>
      );
    case "deck":
      return (
        <p className="flex items-center gap-2 text-sm text-muted-foreground">
          <LayersIcon className="size-3.5 shrink-0" aria-hidden="true" />
          Deck built: {item.deck.name || item.deck.id}.
        </p>
      );
  }
}
