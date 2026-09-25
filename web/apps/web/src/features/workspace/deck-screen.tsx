import { Code, ConnectError } from "@connectrpc/connect";
import type { GetSessionResponse } from "@mtg/api-client/mtg/v1/agent_service_pb";
import { useQuery, useQueryClient } from "@tanstack/react-query";
import { useCallback, useMemo, useState } from "react";
import { Link, useNavigate, useParams } from "react-router";

import { ErrorState } from "../../app/components/error-state";
import { Button } from "../../components/ui/button";
import { Skeleton } from "../../components/ui/skeleton";
import { agentClient, deckClient } from "../../lib/api";
import { errorMessage } from "../../lib/errors";
import { ChatPanel } from "../chat/session-page";
import { emptyState, fromSession } from "../chat/use-chat";
import { needsPowerRead } from "../deck/deck-view";
import { DeckActions } from "./deck-actions";
import { DeckVersions } from "./deck-versions";

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

  // An imported deck whose bracket is the floor of the rules alone, or a
  // 60-card import with no power step, asks the judge again, once per
  // open (D-854, D-864). The answer replaces the deck in the cache. The
  // key sits outside the "deck" prefix, so a write that refreshes the deck
  // never pays for the judge again (D-933).
  const queryClient = useQueryClient();
  useQuery({
    queryKey: ["deck-bracket", id],
    queryFn: async () => {
      const res = await agentClient.readImportBracket({ deckId: id });
      queryClient.setQueryData(["deck", id], (old: typeof deckQuery.data) => (old ? { ...old, deck: res.deck } : old));
      return res;
    },
    enabled: deck !== undefined && needsPowerRead(deck),
    staleTime: Infinity,
    retry: false,
  });

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

  // A revision that ended on the server, after a reload or a Stop, gives
  // the fresh session (REV-046). A new deck moves the page to its address.
  // A build with no new deck mounts the panel again on the fresh session.
  const [reloads, setReloads] = useState(0);
  const onBuildEnded = useCallback(
    (res: GetSessionResponse) => {
      queryClient.setQueryData(["session", sessionId], res);
      const newest = res.session?.deckIds.at(-1) ?? "";
      if (newest !== "" && newest !== id) {
        void navigate(`/decks/${newest}`, { replace: true });
        return;
      }
      setReloads((n) => n + 1);
    },
    [queryClient, sessionId, id, navigate],
  );


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
  // this address names, never the latest of the session. The key is the
  // deck and the reloads: the deck view reads the power note from the
  // deck, so a new note needs no mount, and a mount would stop a turn
  // (D-933). A reload comes only after a build that no stream here reads.
  return (
    <ChatPanel
      key={`${deck.id}:${reloads}`}
      initial={initial}
      session={session}
      building={sessionQuery.data?.building}
      onBuildEnded={onBuildEnded}
      deckOverride={deck}
      baseOverride={base}
      actions={
        <>
          <DeckActions deck={deck} />
          <DeckVersions deck={deck} />
        </>
      }
      onDeckBuilt={(newId) => void navigate(`/decks/${newId}`, { replace: true })}
    />
  );
}
