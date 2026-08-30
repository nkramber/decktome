import { Code, ConnectError } from "@connectrpc/connect";
import { useQuery } from "@tanstack/react-query";
import { useMemo } from "react";
import { Link, useNavigate, useParams } from "react-router";

import { ErrorState } from "../../app/components/error-state";
import { Button } from "../../components/ui/button";
import { Skeleton } from "../../components/ui/skeleton";
import { agentClient, deckClient } from "../../lib/api";
import { errorMessage } from "../../lib/errors";
import { ChatPanel } from "../chat/session-page";
import { emptyState, fromSession } from "../chat/use-chat";
import { DeckActions } from "./deck-actions";

// The one screen of a deck (D-335). A deck's address shows the deck, the
// actions the user owns, and the conversation that built it. Nothing of a
// deck lives at a second address.
export function DeckScreen() {
  const { id = "" } = useParams();
  const navigate = useNavigate();

  const deckQuery = useQuery({
    queryKey: ["deck", id],
    queryFn: () => deckClient.getDeck({ deckId: id }),
    enabled: id !== "",
  });
  const deck = deckQuery.data?.deck;
  const sessionId = deck?.sessionId ?? "";

  // The conversation that built this deck. A deck whose session went
  // shows the deck alone, and it says so.
  const sessionQuery = useQuery({
    queryKey: ["session", sessionId],
    queryFn: () => agentClient.getSession({ sessionId }),
    enabled: sessionId !== "",
  });
  const session = sessionQuery.data?.session;

  // The deck this one revised, for the diff (PR-12B).
  const baseId = deck?.revisedFromDeckId ?? "";
  const baseQuery = useQuery({
    queryKey: ["deck", baseId],
    queryFn: () => deckClient.getDeck({ deckId: baseId }),
    enabled: baseId !== "",
  });
  const base = baseQuery.data?.deck;

  // The panel takes its opening state once. A fresh object on every
  // render would reset the conversation under the reader's hands.
  const initial = useMemo(() => (session && deck ? fromSession(session, deck, base) : emptyState), [session, deck, base]);


  // The screen mounts once, when the deck and its conversation are both
  // in hand. Two mounts would tear the deck down and build it again in
  // front of the reader.
  if (deckQuery.isPending || (sessionId !== "" && sessionQuery.isPending)) {
    return (
      <div className="mx-auto flex w-full max-w-7xl flex-col gap-4 p-4 md:p-6" role="status">
        <span className="sr-only">Loading the deck...</span>
        <Skeleton className="h-9 w-40" />
        <Skeleton className="h-32 w-full rounded-panel" />
        <Skeleton className="h-64 w-full rounded-panel" />
      </div>
    );
  }

  if (deckQuery.isError || !deck) {
    const gone = deckQuery.isError && ConnectError.from(deckQuery.error).code === Code.NotFound;
    return (
      <div className="mx-auto flex w-full max-w-3xl flex-col gap-4 p-4 md:p-6">
        <ErrorState
          title={gone ? "That deck is gone" : "Could not load the deck"}
          message={gone ? "Someone deleted this deck, or the link is wrong." : deckQuery.isError ? errorMessage(deckQuery.error) : "The server returned no deck."}
          onRetry={gone ? undefined : () => void deckQuery.refetch()}
        />
        <Button asChild variant="outline" className="self-center">
          <Link to="/decks">Back to your decks</Link>
        </Button>
      </div>
    );
  }






  // The panel carries the conversation and its dock. It shows the deck
  // this address names, never the latest of the session.
  return (
    <ChatPanel
      key={deck.id}
      initial={initial}
      session={session}
      deckOverride={deck}
      baseOverride={base}
      actions={<DeckActions deck={deck} />}
      onDeckBuilt={(newId) => void navigate(`/decks/${newId}`, { replace: true })}
    />
  );
}
