import type { Card } from "@mtg/api-client/mtg/v1/card_pb";
import type { Deck } from "@mtg/api-client/mtg/v1/deck_pb";
import { StarIcon } from "lucide-react";
import { Link } from "react-router";

import { Button } from "../../components/ui/button";
import { cn } from "../../lib/cn";
import { identityOfCommanders } from "./color-identity";
import { formatLabel, powerLabel, priceText } from "./deck-stats";
import { ManaPips } from "./mana-pips";

// One deck in the library grid (PR-17). Every deck wears the same gold
// frame over a hatched ground (D-329), and the colors of the game appear
// in its mana pips alone.
export function DeckCard({ deck, byId, onFavorite }: { deck: Deck; byId: Map<string, Card>; onFavorite: (favorite: boolean) => void }) {
  const identity = identityOfCommanders(deck.commanderOracleIds, byId);
  const commander = deck.commanderOracleIds.map((id) => byId.get(id)).find(Boolean);
  const title = deck.name || "Untitled deck";

  return (
    <li className="card-hover relative flex flex-col gap-3 overflow-hidden rounded-card border border-border bg-card p-4">
      <span aria-hidden="true" className="hatch pointer-events-none absolute inset-0 opacity-[0.03]" />

      <div className="relative flex items-start justify-between gap-2">
        <div className="min-w-0">
          <h2 className="font-display wrap-anywhere text-[15px] leading-tight font-semibold">
            {/* The whole card is the link target, so the row needs no second control. */}
            <Link to={`/decks/${deck.id}`} className="after:absolute after:inset-0 after:content-['']">
              {title}
            </Link>
          </h2>
          <p className="mt-0.5 font-mono text-[11px] text-muted-foreground">
            {formatLabel(deck.format?.id, deck.format?.houseRules ?? "")}
            {powerLabel(deck.power) && ` · ${powerLabel(deck.power)}`}
          </p>
        </div>
        <Button
          variant="ghost"
          size="icon"
          aria-pressed={deck.favorite}
          aria-label={deck.favorite ? `Remove ${title} from your favorites` : `Add ${title} to your favorites`}
          onClick={() => onFavorite(!deck.favorite)}
          className="relative z-10 -mt-1 -mr-1 size-7"
        >
          <StarIcon className={cn("size-4", deck.favorite && "fill-primary text-primary")} />
        </Button>
      </div>

      <div className="relative flex items-center gap-2">
        <ManaPips colors={identity} size="sm" />
        {commander?.name && <span className="truncate text-[13px] text-secondary-foreground">{commander.name}</span>}
      </div>

      <div className="relative">
        <div className="mb-1 flex justify-between font-mono text-[11px] text-muted-foreground">
          <span>{deck.cardCount > 0 ? `${deck.cardCount} cards` : "No cards"}</span>
          <span>{deck.buyCostUsd > 0 ? `${priceText(deck.buyCostUsd)} to buy` : "Nothing to buy"}</span>
        </div>
        <div className="h-0.5 rounded-full bg-secondary" aria-hidden="true">
          <span className="block h-full rounded-full bg-primary transition-all" style={{ width: `${Math.min(100, (deck.cardCount / 100) * 100)}%` }} />
        </div>
      </div>

      <div className="relative mt-auto flex items-center justify-between pt-1">
        <span className="font-mono text-[10px] text-muted-foreground">{deck.createdAt?.seconds ? new Date(Number(deck.createdAt.seconds) * 1000).toLocaleDateString() : ""}</span>
        <span className="font-display text-[11px] tracking-widest text-muted-foreground uppercase">Open →</span>
      </div>
    </li>
  );
}
