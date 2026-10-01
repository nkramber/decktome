import type { Card } from "@mtg/api-client/mtg/v1/card_pb";
import { type Deck, RerunCase } from "@mtg/api-client/mtg/v1/deck_pb";
import { AlertTriangleIcon } from "lucide-react";
import { Button } from "../../components/ui/button";
import { useDeckCards } from "./use-cards";

// staleNames reads the names of the cards that are no longer legal, in
// the order of the deck. A commander comes first, and a companion last.
// An imported list keeps its commander and its companion as ids alone,
// so the card data names them.
export function staleNames(deck: Deck, byId?: ReadonlyMap<string, Card>): string[] {
  const ids = new Set(deck.staleOracleIds);
  const entries = [...deck.commanders, ...deck.cards, ...deck.sideboard];
  const names: string[] = [];
  const add = (id: string) => {
    if (!ids.has(id)) return;
    const name = entries.find((dc) => dc.oracleId === id)?.name ?? byId?.get(id)?.name;
    if (name && !names.includes(name)) names.push(name);
  };
  for (const id of deck.commanderOracleIds) add(id);
  for (const dc of entries) add(dc.oracleId);
  if (deck.companionOracleId) add(deck.companionOracleId);
  return names;
}

// rerunLabel names the action of the button, from the rerun the deck
// takes (D-1008). A banned commander asks for a new one first (D-1021).
export function rerunLabel(deck: Deck): string {
  if (deck.rerunCase === RerunCase.PATCH) return "Replace the banned cards";
  const commander = deck.commanderOracleIds.some((id) => deck.staleOracleIds.includes(id));
  return commander ? "Pick a new commander" : "Rebuild the deck";
}

// The banner of a deck that a rule change made illegal (D-29, I-1). It
// names the cards, says why the rerun is a patch or a rebuild, and runs
// the rerun in the chat of the deck. A screen with no chat shows the
// banner without the button.
export function StaleBanner(props: { deck: Deck; busy?: boolean; onRerun?: () => void }) {
  if (!props.deck.stale) return null;
  return <StaleBannerBody {...props} />;
}

// StaleBannerBody reads the card data of the deck. The deck view reads
// the same query, so the banner adds no call.
function StaleBannerBody({ deck, busy, onRerun }: { deck: Deck; busy?: boolean; onRerun?: () => void }) {
  const { byId } = useDeckCards(deck);
  const names = staleNames(deck, byId);
  return (
    <section
      role="alert"
      aria-labelledby="stale-title"
      className="mb-4 flex flex-col gap-2 rounded-card border border-danger/50 bg-danger/5 p-3 text-sm print:hidden"
    >
      <h2 id="stale-title" className="flex items-center gap-2 font-semibold">
        <AlertTriangleIcon aria-hidden className="size-4 text-danger" />A rule change made this deck illegal
      </h2>
      {names.length > 0 && <p>No longer legal: {names.join(", ")}.</p>}
      {deck.staleReason && <p className="text-muted-foreground">{deck.staleReason}</p>}
      {onRerun && (
        <div>
          <Button type="button" size="sm" disabled={busy} onClick={onRerun}>
            {rerunLabel(deck)}
          </Button>
        </div>
      )}
    </section>
  );
}
