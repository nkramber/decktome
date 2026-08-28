import type { Card } from "@mtg/api-client/mtg/v1/card_pb";
import type { Question } from "@mtg/api-client/mtg/v1/session_pb";
import { useQuery } from "@tanstack/react-query";

import { cardClient } from "../../lib/api";
import { errorMessage } from "../../lib/errors";
import { copyrightLine, facesOf } from "../deck/card-tile";

// A question whose options are cards, for example a commander offer,
// shows each card with its art, its attribution, its type line, and its
// rules text above the pick button (D-287, D-6).
export function optionIds(q: Question): string[] {
  return q.optionOracleIds ?? [];
}

export function hasCardOptions(q: Question): boolean {
  return optionIds(q).some((id) => id !== "");
}

export function useOptionCards(q: Question) {
  const ids = optionIds(q).filter((id) => id !== "");
  return useQuery({
    queryKey: ["cards", "options", ids],
    queryFn: () => cardClient.getCards({ oracleIds: ids }),
    enabled: ids.length > 0,
    staleTime: Infinity,
  });
}

export function CardOption({ card, name }: { card: Card | undefined; name: string }) {
  const faces = facesOf(card);
  if (!card || faces.length === 0) {
    return <p className="text-sm text-neutral-600">No card data for {name}.</p>;
  }
  return (
    <div className="flex flex-col gap-2" data-testid="card-option">
      {faces.map((face, i) => (
        <figure key={face.imageUris?.normal || face.name || i} className="flex flex-col gap-1">
          {face.imageUris?.normal || face.imageUris?.small ? (
            <img
              src={face.imageUris.normal || face.imageUris.small}
              alt={face.name}
              width={488}
              height={680}
              loading="lazy"
              className="h-auto w-full rounded"
            />
          ) : (
            <div className="rounded border border-neutral-300 bg-neutral-100 p-2 text-center text-sm">{face.name} (no image)</div>
          )}
          <figcaption className="text-xs text-neutral-700">
            {face.name}
            {faces.length > 1 && ` (face ${i + 1} of ${faces.length})`}
            {face.artist ? `. Illustrated by ${face.artist}. ` : ". "}
            {copyrightLine}
          </figcaption>
          <p className="text-sm">
            <span className="font-medium">{face.typeLine}</span>
            {face.manaCost && <span className="text-neutral-600"> {face.manaCost}</span>}
          </p>
          <p className="whitespace-pre-line text-sm">{face.oracleText || "No rules text."}</p>
        </figure>
      ))}
    </div>
  );
}

export function CardOptionsError({ error }: { error: unknown }) {
  return (
    <p role="alert" className="text-sm text-red-700">
      Could not load the card data: {errorMessage(error)}
    </p>
  );
}
