import type { Card, CardFace, ImageUris } from "@mtg/api-client/mtg/v1/card_pb";
import type { DeckCard } from "@mtg/api-client/mtg/v1/deck_pb";
import { useState } from "react";

import { priceText } from "./deck-stats";

// The full card image carries the artist and the copyright line, and no
// CSS crops it. The Scryfall guidelines ask for a separate line only
// beside an art_crop, which the app never shows (D-6, D-291).

type Face = { name: string; artist: string; imageUris?: ImageUris; typeLine: string; oracleText: string; manaCost: string };

// faces returns one entry per face (F-9). A card with no face data falls
// back to the default printing, which carries the artist too.
export function facesOf(card: Card | undefined): Face[] {
  if (!card) return [];
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

function FaceImage({ face }: { face: Face }) {
  // The normal image is 488 by 680. If the CDN refuses it, the small one
  // (146 by 204) takes its place (ui plan, section 8). The width scales
  // and the aspect ratio stays, so nothing is cropped or skewed (D-6).
  const [src, setSrc] = useState(face.imageUris?.normal || face.imageUris?.small || "");
  const small = face.imageUris?.small ?? "";
  // A second load error leaves the text tile, not a broken image.
  function onError() {
    setSrc(small && src !== small ? small : "");
  }
  if (!src) {
    return (
      <div className="flex aspect-[488/680] w-full items-center justify-center rounded border border-neutral-300 bg-neutral-100 p-2 text-center text-sm">
        {face.name} (no image)
      </div>
    );
  }
  return (
    <img
      src={src}
      alt={`${face.name} (card)`}
      width={488}
      height={680}
      loading="lazy"
      className="h-auto w-full rounded"
      onError={onError}
    />
  );
}

// CardTile shows one deck entry: every face with its art and attribution,
// the count, the owned mark or the price, and the Oracle text on demand.
// hideOwnership is for a commander entry built from the card data: the
// deck carries no owned mark for it, so the tile shows none.
export function CardTile({
  entry,
  card,
  isCommander,
  hideOwnership = false,
}: {
  entry: DeckCard;
  card: Card | undefined;
  isCommander?: boolean;
  hideOwnership?: boolean;
}) {
  const faces = facesOf(card);
  const name = card?.name || entry.name;
  return (
    <li className="flex flex-col gap-1 rounded border border-neutral-200 p-2" data-testid="card-tile">
      <div className="flex items-baseline justify-between gap-2">
        <span className="min-w-0 wrap-anywhere font-medium">
          {entry.count > 1 && <span className="mr-1 text-neutral-600">{entry.count}×</span>}
          {name}
        </span>
        {isCommander && (
          <span className="rounded border border-amber-300 bg-amber-100 px-1 text-xs" data-testid="commander-mark">
            Commander
          </span>
        )}
      </div>
      {faces.length === 0 && <p className="text-sm text-neutral-600">No card data for this entry.</p>}
      {faces.map((face, i) => (
        <figure key={face.imageUris?.normal || face.name || i} className="flex flex-col gap-1">
          <FaceImage face={face} />
          {faces.length > 1 && (
            <figcaption className="text-xs text-neutral-700">
              {face.name} (face {i + 1} of {faces.length})
            </figcaption>
          )}
        </figure>
      ))}
      {!hideOwnership && (
      <p className="text-xs">
        {entry.owned ? (
          <span className="rounded border border-green-300 bg-green-100 px-1" data-testid="owned-mark">
            Owned{entry.ownedCount > 0 ? ` (${entry.ownedCount})` : ""}
          </span>
        ) : (
          <span className="rounded border border-red-300 bg-red-100 px-1" data-testid="buy-mark">
            To buy: {priceText(entry.priceUsd)}
          </span>
        )}
      </p>
      )}
      {entry.reason && <p className="text-xs text-neutral-700">{entry.reason}</p>}
      {card && (
        <details className="text-sm">
          <summary className="cursor-pointer">Oracle text</summary>
          {faces.map((face, i) => (
            <div key={i} className="mt-1">
              <p className="font-medium">
                {face.name} {face.manaCost}
              </p>
              <p className="text-neutral-700">{face.typeLine}</p>
              <p className="whitespace-pre-line">{face.oracleText || "No rules text."}</p>
            </div>
          ))}
        </details>
      )}
    </li>
  );
}
