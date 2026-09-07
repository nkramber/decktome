import type { Card, CardFace, ImageUris } from "@mtg/api-client/mtg/v1/card_pb";
import type { DeckCard } from "@mtg/api-client/mtg/v1/deck_pb";
import { FeedbackKind } from "@mtg/api-client/mtg/v1/feedback_service_pb";
import { useState } from "react";

import { Thumbs } from "../feedback/thumbs";
import { priceText } from "./deck-stats";

// The full card image carries the artist and the copyright line, and no
// CSS crops it. The Scryfall guidelines ask for a separate line only
// beside an art_crop, which the app never shows (D-6, D-291).

export type Face = { name: string; artist: string; imageUris?: ImageUris; typeLine: string; oracleText: string; manaCost: string };

// faces returns one entry per face (F-9). A card with no face data falls
// back to the default printing, which carries the artist too. An owned
// printing with an image replaces the default one on a single-faced card,
// so the user sees the copy they hold (D-299).
export function facesOf(card: Card | undefined, owned?: { imageUris?: ImageUris; artist: string }): Face[] {
  if (!card) return [];
  if (owned?.imageUris && card.faces.length <= 1) {
    return [
      {
        name: card.name,
        artist: owned.artist || card.defaultPrinting?.artist || "",
        imageUris: owned.imageUris,
        typeLine: card.typeLine,
        oracleText: card.oracleText,
        manaCost: card.manaCost,
      },
    ];
  }
  if (card.faces.length > 0) {
    return card.faces.map((f: CardFace) => ({
      name: f.name,
      artist: f.artist,
      imageUris: f.imageUris,
      typeLine: f.typeLine,
      oracleText: f.oracleText,
      manaCost: f.manaCost,
    }));
  }
  const p = card.defaultPrinting;
  return [
    {
      name: card.name,
      artist: p?.artist ?? "",
      imageUris: p?.imageUris,
      typeLine: card.typeLine,
      oracleText: card.oracleText,
      manaCost: card.manaCost,
    },
  ];
}

export function FaceImage({ face, size = "normal" }: { face: Face; size?: "normal" | "small" }) {
  // The normal image is 488 by 680. If the CDN refuses it, the small one
  // (146 by 204) takes its place (ui plan, section 8). The width scales
  // and the aspect ratio stays, so nothing is cropped or skewed (D-6).
  // The sample hand asks for the small one from the start (PR-20).
  const small = face.imageUris?.small ?? "";
  const [src, setSrc] = useState(size === "small" ? small : face.imageUris?.normal || small);
  // A second load error leaves the text tile, not a broken image.
  function onError() {
    setSrc(small && src !== small ? small : "");
  }
  if (!src) {
    return (
      <div className="flex aspect-[488/680] w-full items-center justify-center rounded-card border border-border bg-muted p-2 text-center text-sm">
        {face.name} (no image)
      </div>
    );
  }
  return (
    <img
      src={src}
      alt={`${face.name} (card)`}
      width={size === "small" ? 146 : 488}
      height={size === "small" ? 204 : 680}
      loading="lazy"
      className="h-auto w-full rounded"
      onError={onError}
    />
  );
}

// CardTile shows one deck entry: every face with its art and attribution,
// the count, and the owned mark or the price. The image carries the rules text.
// hideOwnership is for a commander entry built from the card data: the
// deck carries no owned mark for it, so the tile shows none. onOpen makes
// the name a button that opens the card detail (PR-20). feedbackDeckId
// puts the thumbs at the foot (PR-27, D-559): a phone shows them always,
// and a pointer shows them on hover and on focus, and after a verdict.
export function CardTile({
  entry,
  card,
  isCommander,
  hideOwnership = false,
  onOpen,
  feedbackDeckId,
}: {
  entry: DeckCard;
  card: Card | undefined;
  isCommander?: boolean;
  hideOwnership?: boolean;
  onOpen?: () => void;
  feedbackDeckId?: string;
}) {
  const faces = facesOf(card, entry.ownedPrinting);
  const name = card?.name || entry.name;
  return (
    <li className="group flex flex-col gap-1.5 rounded-card border border-border bg-card p-2 transition-shadow hover:shadow-raised" data-testid="card-tile">
      <div className="flex items-baseline justify-between gap-2">
        <span className="min-w-0 wrap-anywhere font-medium">
          {entry.count > 1 && <span className="mr-1 text-muted-foreground">{entry.count}×</span>}
          {onOpen ? (
            <button type="button" onClick={onOpen} className="text-left underline-offset-4 hover:text-primary hover:underline focus-visible:underline" title="Open the card detail">
              {name}
            </button>
          ) : (
            name
          )}
        </span>
        {isCommander && (
          <span className="rounded-md border border-warning/50 bg-warning/15 px-1.5 py-0.5 text-xs font-medium" data-testid="commander-mark">
            Commander
          </span>
        )}
      </div>
      {/* The build marks a card the reader's sets do not hold (D-383). */}
      {entry.outsideRequestedSets && (
        <p className="text-xs">
          <span
            className="rounded-md border border-danger/50 bg-danger/15 px-1.5 py-0.5 font-medium text-danger"
            data-testid="outside-set-mark"
          >
            <span aria-hidden="true">! </span>Not from requested set
          </span>
        </p>
      )}
      {faces.length === 0 && <p className="text-sm text-muted-foreground">No card data for this entry.</p>}
      {faces.map((face, i) => (
        <figure key={`${entry.oracleId}-${i}`} className="flex flex-col gap-1 print:hidden">
          <FaceImage face={face} />
          {faces.length > 1 && (
            <figcaption className="text-xs text-muted-foreground">
              {face.name} (face {i + 1} of {faces.length})
            </figcaption>
          )}
        </figure>
      ))}
      {!hideOwnership && (
      <p className="text-xs">
        {entry.owned ? (
          <span className="rounded-md border border-success/50 bg-success/15 px-1.5 py-0.5 font-medium" data-testid="owned-mark">
            Owned{entry.ownedCount > 0 ? ` (${entry.ownedCount})` : ""}
          </span>
        ) : (
          <span className="rounded-md border border-danger/50 bg-danger/15 px-1.5 py-0.5 font-medium" data-testid="buy-mark">
            To buy: {priceText(entry.priceUsd)}
          </span>
        )}
      </p>
      )}
      {entry.reason && <p className="text-xs text-muted-foreground">{entry.reason}</p>}
      {feedbackDeckId && entry.oracleId && (
        <Thumbs
          target={{ kind: FeedbackKind.CARD, deckId: feedbackDeckId, oracleId: entry.oracleId }}
          itemName={name}
          className="transition-opacity print:hidden pointer-fine:opacity-0 pointer-fine:group-hover:opacity-100 pointer-fine:group-focus-within:opacity-100 pointer-fine:data-[verdict]:opacity-100"
        />
      )}
    </li>
  );
}
