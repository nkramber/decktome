import type { Deck } from "@mtg/api-client/mtg/v1/deck_pb";
import { useQuery } from "@tanstack/react-query";
import { SparklesIcon } from "lucide-react";
import { useState } from "react";
import { Button } from "../../components/ui/button";
import { cardQueryRetry } from "../../lib/card-retry";
import { deckOracleIds, fetchCards } from "./use-cards";
import { useDeckWrites } from "./use-decks";

// newCardIds lists the new cards of a new set that fit the deck, and that
// the deck does not hold yet (D-1091). A revise that added one leaves it
// off the panel. A deck stored before the field reads as none.
export function newCardIds(deck: Deck): string[] {
  const fresh = deck.newOracleIds ?? [];
  if (fresh.length === 0) return [];
  const held = new Set(deckOracleIds(deck));
  return fresh.filter((id) => !held.has(id));
}

// reviseMessage is the message of the Revise button. The revise turn of
// the chat reads it, and it costs one turn only on the tap (D-1091).
export function reviseMessage(names: string[]): string {
  return `A new set has cards that may fit this deck: ${names.join(", ")}. Add the ones that make the deck better.`;
}

// The panel of the new cards of a new set that fit the deck (D-1091). It
// names the cards and runs a revise in the chat of the deck. A screen
// with no chat shows the panel without the Revise button. A dismiss
// clears the panel, and the next pass of the snapshot job replaces it
// (D-1092).
export function NewCardsPanel(props: { deck: Deck; busy?: boolean; onRevise?: (message: string) => void }) {
  const [dismissed, setDismissed] = useState("");
  const ids = newCardIds(props.deck);
  if (ids.length === 0 || dismissed === props.deck.id) return null;
  return <NewCardsBody {...props} ids={ids} onDismissed={() => setDismissed(props.deck.id)} />;
}

function NewCardsBody({
  deck,
  ids,
  busy,
  onRevise,
  onDismissed,
}: {
  deck: Deck;
  ids: string[];
  busy?: boolean;
  onRevise?: (message: string) => void;
  onDismissed: () => void;
}) {
  const query = useQuery({
    queryKey: ["cards", "new", ids],
    queryFn: () => fetchCards(ids),
    ...cardQueryRetry,
    staleTime: Infinity,
  });
  const { dismissNewCards } = useDeckWrites();
  const names = ids.map((id) => query.data?.byId.get(id)?.name).filter((n): n is string => !!n);
  return (
    <section
      aria-labelledby="new-cards-title"
      className="mb-4 flex flex-col gap-2 rounded-card border border-primary/40 bg-primary/5 p-3 text-sm print:hidden"
    >
      <h2 id="new-cards-title" className="flex items-center gap-2 font-semibold">
        <SparklesIcon aria-hidden className="size-4 text-primary" />
        New cards for this deck
      </h2>
      {names.length > 0 ? <p>{names.join(", ")}</p> : <p className="text-muted-foreground">Loading the new cards...</p>}
      {dismissNewCards.isError && (
        <p role="alert" className="text-danger">
          Could not dismiss the new cards. Try again.
        </p>
      )}
      <div className="flex flex-wrap gap-2">
        {onRevise && (
          <Button type="button" size="sm" disabled={busy || names.length === 0} onClick={() => onRevise(reviseMessage(names))}>
            Revise with these cards
          </Button>
        )}
        <Button
          type="button"
          size="sm"
          variant="outline"
          disabled={dismissNewCards.isPending}
          onClick={() => dismissNewCards.mutate({ deckId: deck.id }, { onSuccess: onDismissed })}
        >
          Dismiss
        </Button>
      </div>
    </section>
  );
}
