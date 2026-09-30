import { AlertTriangleIcon } from "lucide-react";
import { type Deck, RerunCase } from "@mtg/api-client/mtg/v1/deck_pb";
import { Button } from "../../components/ui/button";

// staleNames reads the names of the cards that are no longer legal, in
// the order of the deck. A commander comes first.
export function staleNames(deck: Deck): string[] {
  const ids = new Set(deck.staleOracleIds);
  const names: string[] = [];
  for (const dc of [...deck.commanders, ...deck.cards, ...deck.sideboard]) {
    if (ids.has(dc.oracleId) && !names.includes(dc.name)) names.push(dc.name);
  }
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
export function StaleBanner({ deck, busy, onRerun }: { deck: Deck; busy?: boolean; onRerun?: () => void }) {
  if (!deck.stale) return null;
  const names = staleNames(deck);
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
