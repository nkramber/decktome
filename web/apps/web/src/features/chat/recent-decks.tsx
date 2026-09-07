import type { Deck } from "@mtg/api-client/mtg/v1/deck_pb";
import { useQuery } from "@tanstack/react-query";
import { useEffect } from "react";

import { Skeleton } from "../../components/ui/skeleton";
import { deckClient } from "../../lib/api";
import { useAppStore } from "../../lib/store";
import { notify } from "../../app/components/notify";
import { errorMessage } from "../../lib/errors";
import { DeckCard } from "../deck/deck-card";
import { useCommanderCards, useDeckWrites } from "../deck/use-decks";

// The three newest decks, at the head of a new chat (D-350). They wear
// the same tile as the library, art and all: a reader who came back for
// a deck reads the same thing in both places (D-358).
const recentCount = 3;

// useRecentDecks reads the newest decks. The page reads the same list
// for its title, so one call serves both.
//
// It also marks whether the reader owns a deck, so the next load knows
// whether a placeholder belongs on the page (F-67).
export function useRecentDecks(enabled: boolean) {
  const setHadDecks = useAppStore((s) => s.setHadDecks);
  const list = useQuery({
    queryKey: ["decks", "recent", recentCount],
    queryFn: () => deckClient.listDecks({ pageSize: recentCount }),
    enabled,
  });
  const decks = list.data?.decks ?? [];
  const answered = list.isSuccess;
  useEffect(() => {
    if (answered) setHadDecks(decks.length > 0);
  }, [answered, decks.length, setHadDecks]);
  return { decks, isPending: list.isPending && enabled };
}

export function RecentDecks({ decks, isPending }: { decks: Deck[]; isPending: boolean }) {
  const hadDecks = useAppStore((s) => s.hadDecks);
  const byId = useCommanderCards(decks);
  const { setFavorite } = useDeckWrites();

  async function onFavorite(deck: Deck, favorite: boolean) {
    try {
      await setFavorite.mutateAsync({ deckId: deck.id, favorite });
    } catch (err) {
      await notify("error", "Could not change the favorite mark", errorMessage(err));
    }
  }

  // A placeholder belongs on the page only when a deck follows it. A
  // reader with no deck saw three boxes appear and go, and a reader who
  // never opened the app before saw the same (F-67). The mark of the
  // last read answers it, and the page holds still for a reader who owns
  // a deck.
  if (isPending) {
    if (!hadDecks) return null;
    return (
      <div role="status" className="grid gap-4 sm:grid-cols-3">
        <span className="sr-only">Loading your decks...</span>
        <Skeleton className="h-40 rounded-card" />
        <Skeleton className="h-40 rounded-card" />
        <Skeleton className="h-40 rounded-card" />
      </div>
    );
  }
  // A reader with no deck yet needs the empty page, not an empty heading.
  if (decks.length === 0) return null;

  return (
    <section aria-label="Your newest decks">
      <ul className="grid gap-4 sm:grid-cols-3">
        {decks.map((deck) => (
          <DeckCard key={deck.id} deck={deck} byId={byId} onFavorite={(favorite) => void onFavorite(deck, favorite)} />
        ))}
      </ul>
    </section>
  );
}
