import type { Card } from "@mtg/api-client/mtg/v1/card_pb";
import type { Question } from "@mtg/api-client/mtg/v1/session_pb";
import { useQuery } from "@tanstack/react-query";

import { cardClient } from "../../lib/api";
import { errorMessage } from "../../lib/errors";
import { FaceImage, facesOf } from "../deck/card-tile";

// A question whose options are cards, for example a commander offer,
// shows each full card image above the pick button. The image carries
// the rules text, the artist, and the copyright line, so the tile
// repeats none of them (D-287, D-291).
export function optionIds(q: Question): string[] {
  return q.optionOracleIds;
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
    return <p className="text-sm text-muted-foreground">No card data for {name}.</p>;
  }
  return (
    <div className="flex flex-col gap-2" data-testid="card-option">
      {faces.map((face, i) => (
        <figure key={`${card.oracleId}-${i}`} className="flex flex-col gap-1">
          <FaceImage face={face} />
          {faces.length > 1 && (
            <figcaption className="text-xs text-muted-foreground">
              {face.name} (face {i + 1} of {faces.length})
            </figcaption>
          )}
        </figure>
      ))}
    </div>
  );
}

export function CardOptionsError({ error }: { error: unknown }) {
  return (
    <p role="alert" className="text-sm text-danger">
      Could not load the card data: {errorMessage(error)}
    </p>
  );
}
