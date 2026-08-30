import type { Card } from "@mtg/api-client/mtg/v1/card_pb";
import type { Deck } from "@mtg/api-client/mtg/v1/deck_pb";
import { StarIcon } from "lucide-react";
import { useState } from "react";
import { Link } from "react-router";

import { Button } from "../../components/ui/button";
import { cn } from "../../lib/cn";
import { identityOf, identityVars } from "./color-identity";
import { formatLabel, powerLabel, priceText } from "./deck-stats";
import { ManaPips } from "./mana-pips";

// One deck in the library grid (PR-17). The deck's own color identity
// washes the card, and the commander's artwork sits behind the name
// (D-327). Scryfall serves that artwork as art_crop, so the app crops
// nothing itself.
export function DeckCard({ deck, byId, onFavorite }: { deck: Deck; byId: Map<string, Card>; onFavorite: (favorite: boolean) => void }) {
  const identity = identityOf(deck.commanderOracleIds, byId);
  const commander = deck.commanderOracleIds.map((id) => byId.get(id)).find(Boolean);
  const face = commander?.faces?.[0] ?? commander?.defaultPrinting;
  const art = face?.imageUris?.artCrop ?? "";
  const artist = commander?.defaultPrinting?.artist ?? "";
  const [artFailed, setArtFailed] = useState(false);
  const title = deck.name || "Untitled deck";

  return (
    <li className="group relative isolate flex flex-col overflow-hidden rounded-panel border border-border bg-surface shadow-card transition-shadow hover:shadow-raised" style={identityVars(identity)}>
      <span aria-hidden="true" className="identity-rule absolute inset-x-0 top-0 z-10 h-1" />
      <div className="relative h-28 overflow-hidden bg-muted">
        {art && !artFailed ? (
          <img src={art} alt="" aria-hidden="true" loading="lazy" onError={() => setArtFailed(true)} className="size-full object-cover opacity-70 transition-transform duration-500 group-hover:scale-105" />
        ) : (
          <span aria-hidden="true" className="identity-wash absolute inset-0" />
        )}
        <span aria-hidden="true" className="absolute inset-0 bg-gradient-to-t from-surface via-surface/70 to-transparent" />
        <span aria-hidden="true" className="identity-wash absolute inset-0 opacity-70" />
        <Button
          variant="ghost"
          size="icon"
          aria-pressed={deck.favorite}
          aria-label={deck.favorite ? `Remove ${title} from your favorites` : `Add ${title} to your favorites`}
          onClick={() => onFavorite(!deck.favorite)}
          className="absolute top-2 right-2 z-10 bg-background/60 backdrop-blur hover:bg-background/80"
        >
          <StarIcon className={cn("size-4", deck.favorite && "fill-warning text-warning")} />
        </Button>
      </div>

      <div className="flex grow flex-col gap-2 p-4 pt-2">
        <div className="flex items-start justify-between gap-2">
          <h2 className="wrap-anywhere text-base leading-tight font-semibold tracking-tight text-balance">
            {/* The whole card is the link target, so the row needs no second control. */}
            <Link to={`/decks/${deck.id}`} className="after:absolute after:inset-0 after:content-['']">
              {title}
            </Link>
          </h2>
          <ManaPips colors={identity} className="mt-1 shrink-0" />
        </div>

        <p className="text-xs text-muted-foreground">
          {formatLabel(deck.format?.id, deck.format?.houseRules ?? "")}
          {powerLabel(deck.power) && ` · ${powerLabel(deck.power)}`}
          {deck.cardCount > 0 && ` · ${deck.cardCount} cards`}
        </p>

        <div className="mt-auto flex flex-wrap items-center gap-x-3 gap-y-1 pt-1 text-xs text-muted-foreground">
          <span>{deck.buyCostUsd > 0 ? `${priceText(deck.buyCostUsd)} to buy` : "Nothing to buy"}</span>
          {deck.createdAt?.seconds ? <span>{new Date(Number(deck.createdAt.seconds) * 1000).toLocaleDateString()}</span> : null}
          {artist && <span className="sr-only">Art by {artist}, copyright Wizards of the Coast, from Scryfall.</span>}
        </div>
      </div>
    </li>
  );
}
