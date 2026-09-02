import type { Card } from "@mtg/api-client/mtg/v1/card_pb";
import type { Question } from "@mtg/api-client/mtg/v1/session_pb";
import { useQuery } from "@tanstack/react-query";

import { cardClient } from "../../lib/api";
import { cn } from "../../lib/cn";
import { errorMessage } from "../../lib/errors";
import { FaceImage, facesOf } from "../deck/card-tile";

// A question whose options are cards, for example a commander offer,
// shows each full card image above the pick button. The image carries
// the rules text, the artist, and the copyright line, so the tile
// repeats none of them (D-287, D-291).
export function optionIds(q: Question): string[] {
  return q.optionOracleIds ?? [];
}

// partnerIds is parallel to the options as well. A commander pair reads
// as "A + B" in one option, and the second card lands here (D-361).
export function partnerIds(q: Question): string[] {
  return q.optionPartnerOracleIds ?? [];
}

export function hasCardOptions(q: Question): boolean {
  return optionIds(q).some((id) => id !== "");
}

export function useOptionCards(q: Question) {
  const ids = [...optionIds(q), ...partnerIds(q)].filter((id) => id !== "");
  return useQuery({
    queryKey: ["cards", "options", ids],
    queryFn: () => cardClient.getCards({ oracleIds: ids }),
    enabled: ids.length > 0,
    staleTime: Infinity,
  });
}

// zoom says which half of a commander pair a card is. A pair draws two
// cards at half width, and the pointer lifts one to the size of a single
// card, from its outer corner, over its partner (D-372, D-443). A single
// card is already that size, so it never zooms.
// onPick makes the art itself pick the option (D-444). The name button
// stays the control a keyboard and a screen reader use, so the art is a
// pointer shortcut: out of the tab order and hidden from assistive
// technology, which already has the button.
export function CardOption({ card, name, zoom, onPick }: { card: Card | undefined; name: string; zoom?: "left" | "right"; onPick?: () => void }) {
  const faces = facesOf(card);
  if (!card || faces.length === 0) {
    return <p className="text-sm text-muted-foreground">No card data for {name}.</p>;
  }
  const art = (
    <div className="flex flex-col gap-2" data-testid="card-option">
      {faces.map((face, i) => (
        <figure key={`${card.oracleId}-${i}`} className="group/card flex flex-col gap-1">
          <span
            data-zoom={zoom}
            className={cn(
              "relative z-0 block",
              zoom && "transition-transform duration-150 group-hover/card:z-50 group-hover/card:scale-[2] motion-reduce:transition-none motion-reduce:group-hover/card:scale-100",
              zoom === "left" && "origin-top-left",
              zoom === "right" && "origin-top-right",
            )}
          >
            <FaceImage face={face} />
          </span>
          {faces.length > 1 && (
            <figcaption className="text-xs text-muted-foreground">
              {face.name} (face {i + 1} of {faces.length})
            </figcaption>
          )}
        </figure>
      ))}
    </div>
  );
  if (!onPick) return art;
  // A plain box takes the click, so the image under it stays an image
  // for assistive technology, with the card name as its text. A button
  // here would turn its children presentational and hide that name.
  return (
    // eslint-disable-next-line jsx-a11y/click-events-have-key-events, jsx-a11y/no-static-element-interactions
    <div onClick={onPick} className="cursor-pointer" data-testid="card-art-pick">
      {art}
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
