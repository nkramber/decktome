import type { Card, CardFace, ImageUris } from "@mtg/api-client/mtg/v1/card_pb";
import type { DeckCard } from "@mtg/api-client/mtg/v1/deck_pb";
import { useState } from "react";

import { priceText } from "./deck-stats";

// The copyright line every image carries (Scryfall API docs, read
// 2026-08-28: "Do not cover, crop, or clip off the copyright or artist
// name", and the footer wording "copyright Wizards of the Coast, LLC").
export const copyrightLine = "© Wizards of the Coast, LLC";

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
      alt={face.name}
      title={face.oracleText || face.name}
      width={488}
      height={680}
      loading="lazy"
      className="h-auto w-full rounded"
      onError={() => {
        if (small && src !== small) setSrc(small);
      }}
    />
  );
}

// CardTile shows one deck entry: every face with its art and attribution,
// the count, the owned mark or the price, and the Oracle text on demand.
export function CardTile({ entry, card, isCommander }: { entry: DeckCard; card: Card | undefined; isCommander?: boolean }) {
  const faces = facesOf(card);
  const name = card?.name || entry.name;
  return (
    <li className="flex flex-col gap-1 rounded border border-neutral-200 p-2" data-testid="card-tile">
      <div className="flex items-baseline justify-between gap-2">
        <span className="font-medium">
          {entry.count > 1 && <span className="mr-1 text-neutral-600">{entry.count}×</span>}
          {name}
        </span>
        {isCommander && <span className="rounded bg-amber-100 px-1 text-xs" data-testid="commander-mark">Commander</span>}
      </div>
      {faces.length === 0 && <p className="text-sm text-neutral-600">No card data for this entry.</p>}
      {faces.map((face, i) => (
        <figure key={i} className="flex flex-col gap-1">
          <FaceImage face={face} />
          <figcaption className="text-xs text-neutral-700">
            {face.name}
            {faces.length > 1 && ` (face ${i + 1} of ${faces.length})`}
            {face.artist ? `. Illustrated by ${face.artist}. ` : ". "}
            {copyrightLine}
          </figcaption>
        </figure>
      ))}
      <p className="text-xs">
        {entry.owned ? (
          <span className="rounded bg-green-100 px-1" data-testid="owned-mark">
            Owned{entry.ownedCount > 0 ? ` (${entry.ownedCount})` : ""}
          </span>
        ) : (
          <span className="rounded bg-red-100 px-1" data-testid="buy-mark">
            To buy: {priceText(entry.priceUsd)}
          </span>
        )}
      </p>
      {entry.reason && <p className="text-xs text-neutral-700">{entry.reason}</p>}
      {card && (
        <details className="text-xs">
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
