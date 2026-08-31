import type { Deck } from "@mtg/api-client/mtg/v1/deck_pb";
import { useQuery } from "@tanstack/react-query";
import { Link } from "react-router";

import { deckClient } from "../../lib/api";
import { Skeleton } from "../../components/ui/skeleton";

// The three newest decks, at the head of a new chat (D-350). A reader who
// came back for a deck they already built finds it here, instead of one
// link to the last conversation.
const recentCount = 3;

function dateOf(deck: Deck): string {
  const seconds = deck.createdAt?.seconds;
  return seconds ? new Date(Number(seconds) * 1000).toLocaleDateString() : "";
}

export function RecentDecks() {
  const list = useQuery({
    queryKey: ["decks", "recent", recentCount],
    queryFn: () => deckClient.listDecks({ pageSize: recentCount }),
  });
  const decks = list.data?.decks ?? [];

  if (list.isPending) {
    return (
      <div role="status" className="flex flex-wrap gap-3">
        <span className="sr-only">Loading your decks...</span>
        <Skeleton className="h-16 w-56 rounded-card" />
        <Skeleton className="h-16 w-56 rounded-card" />
      </div>
    );
  }
  // A reader with no deck yet needs the empty page, not an empty heading.
  if (decks.length === 0) return null;

  return (
    <section aria-labelledby="recent-title" className="flex flex-col gap-2">
      <div className="flex items-baseline justify-between gap-3">
        <h2 id="recent-title" className="font-display text-[10px] tracking-[0.15em] text-muted-foreground uppercase">
          Pick up where you left off
        </h2>
        <Link to="/decks" className="font-mono text-[11px] text-muted-foreground underline-offset-4 hover:text-foreground hover:underline">
          All decks
        </Link>
      </div>
      <ul className="grid gap-3 sm:grid-cols-3">
        {decks.map((deck) => (
          <li key={deck.id}>
            <Link
              to={`/decks/${deck.id}`}
              className="card-hover flex h-full flex-col gap-1 rounded-card border border-border bg-card p-3 transition-colors hover:border-primary/60"
            >
              <span className="font-display truncate text-[13px] font-semibold">{deck.name || "Untitled deck"}</span>
              <span className="font-mono text-[10px] text-muted-foreground">
                {deck.cardCount > 0 ? `${deck.cardCount} cards` : "No cards"}
                {dateOf(deck) && ` · ${dateOf(deck)}`}
              </span>
            </Link>
          </li>
        ))}
      </ul>
    </section>
  );
}
