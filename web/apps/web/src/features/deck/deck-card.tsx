import type { Card } from "@mtg/api-client/mtg/v1/card_pb";
import type { Deck } from "@mtg/api-client/mtg/v1/deck_pb";
import { StarIcon, Trash2Icon } from "lucide-react";
import { Link } from "react-router";

import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "../../components/ui/alert-dialog";
import { Button } from "../../components/ui/button";
import { cn } from "../../lib/cn";
import { identityOfCommanders } from "./color-identity";
import { formatLabel, powerLabel, priceText } from "./deck-stats";
import { ManaPips } from "./mana-pips";

// One deck in the library grid (PR-17). Every deck wears the same gold
// frame over a hatched ground (D-329), and the colors of the game appear
// in its mana pips alone.
// The li is the one positioned ancestor of the card, so the stretched
// link covers the whole tile. A positioned row in between took the
// link's ::after with it, and only that row opened the deck (D-365).
export function DeckCard({
  deck,
  byId,
  onFavorite,
  onDelete,
}: {
  deck: Deck;
  byId: Map<string, Card>;
  onFavorite: (favorite: boolean) => void;
  // onDelete removes the deck for good, after the question (D-439).
  onDelete?: () => void;
}) {
  const identity = identityOfCommanders(deck.commanderOracleIds, byId);
  const commander = deck.commanderOracleIds.map((id) => byId.get(id)).find(Boolean);
  const title = deck.name || "Untitled deck";

  return (
    <li className="card-hover relative isolate flex flex-col gap-3 overflow-hidden rounded-card border border-border bg-card p-4">
      <span aria-hidden="true" className="hatch pointer-events-none absolute inset-0 -z-10 opacity-[0.03]" />

      <div className="flex items-start justify-between gap-2">
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
        {onDelete && (
          <AlertDialog>
            <AlertDialogTrigger asChild>
              <Button variant="ghost" size="icon" aria-label={`Delete ${title}`} className="relative z-10 -mt-1 -mr-1 size-7 text-danger hover:text-danger">
                <Trash2Icon className="size-4" />
              </Button>
            </AlertDialogTrigger>
            <AlertDialogContent>
              <AlertDialogHeader>
                <AlertDialogTitle>Delete {title}?</AlertDialogTitle>
                <AlertDialogDescription>The deck goes for good. The chat that built it stays, and you can build again from it.</AlertDialogDescription>
              </AlertDialogHeader>
              <AlertDialogFooter>
                <AlertDialogCancel>Keep it</AlertDialogCancel>
                <AlertDialogAction onClick={onDelete}>Delete the deck</AlertDialogAction>
              </AlertDialogFooter>
            </AlertDialogContent>
          </AlertDialog>
        )}
      </div>

      <div className="flex items-center gap-2">
        <ManaPips colors={identity} size="sm" />
        {commander?.name && <span className="truncate text-[13px] text-secondary-foreground">{commander.name}</span>}
      </div>

      <div>
        <div className="mb-1 flex justify-between font-mono text-[11px] text-muted-foreground">
          <span>{deck.cardCount > 0 ? `${deck.cardCount} cards` : "No cards"}</span>
          <span>{deck.buyCostUsd > 0 ? `${priceText(deck.buyCostUsd)} to buy` : "Nothing to buy"}</span>
        </div>
        <div className="h-0.5 rounded-full bg-secondary" aria-hidden="true">
          <span className="block h-full rounded-full bg-primary transition-all" style={{ width: `${Math.min(100, (deck.cardCount / 100) * 100)}%` }} />
        </div>
      </div>

      <div className="mt-auto pt-1">
        <span className="font-mono text-[10px] text-muted-foreground">{deck.createdAt?.seconds ? new Date(Number(deck.createdAt.seconds) * 1000).toLocaleDateString() : ""}</span>
      </div>
    </li>
  );
}
