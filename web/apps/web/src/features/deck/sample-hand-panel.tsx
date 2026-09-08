import type { Card } from "@mtg/api-client/mtg/v1/card_pb";
import type { Deck } from "@mtg/api-client/mtg/v1/deck_pb";
import { useRef, useState } from "react";

import { Button } from "../../components/ui/button";
import { FaceImage, facesOf } from "./card-tile";
import { bottom, draw, type Hand, libraryOf, maxMulligans, mulligan, openingHand, seededRandom } from "./sample-hand";

// The sample hand panel (D-318): seven from the exact main deck, a
// mulligan to six and to five by the London rule, and one draw at a
// time. The reader picks the cards that go to the bottom.
export function SampleHand({ deck, byId, seed }: { deck: Deck; byId: Map<string, Card>; seed?: number }) {
  // The generator starts on the first click, so the render stays pure
  // and a test can hand over a seed.
  const random = useRef<(() => number) | null>(null);
  const rng = () => (random.current ??= seededRandom(seed ?? Date.now()));
  const [state, setState] = useState<Hand | null>(null);
  const library = libraryOf(deck.cards, deck.commanderOracleIds);
  const empty = library.length === 0;

  function status(): string {
    if (!state) return `${library.length} cards in the library. Draw seven to start.`;
    if (state.toBottom > 0) {
      return `Mulligan ${state.mulligans}: choose ${state.toBottom} card${state.toBottom === 1 ? "" : "s"} to put on the bottom.`;
    }
    const parts = [`${state.hand.length} cards in hand`, `${state.library.length} in the library`];
    if (state.mulligans > 0) parts.push(`${state.mulligans} mulligan${state.mulligans === 1 ? "" : "s"}`);
    if (state.library.length === 0) parts.push("the library is empty");
    return parts.join(", ") + ".";
  }

  return (
    <section aria-labelledby={`hand-title-${deck.id}`} className="rounded-panel border border-border bg-card p-5 shadow-card">
      <h3 id={`hand-title-${deck.id}`} className="font-display mb-3 text-xs tracking-[0.15em] text-muted-foreground uppercase">
        Sample hand
      </h3>
      <div className="flex flex-wrap items-center gap-2">
        <Button variant="outline" size="sm" disabled={empty} onClick={() => setState(openingHand(library, rng()))}>
          Draw seven
        </Button>
        <Button
          variant="outline"
          size="sm"
          disabled={!state || state.toBottom > 0 || state.mulligans >= maxMulligans}
          onClick={() => setState((s) => (s ? mulligan(s, rng()) : s))}
        >
          Mulligan
        </Button>
        <Button variant="outline" size="sm" disabled={!state || state.toBottom > 0 || state.library.length === 0} onClick={() => setState((s) => (s ? draw(s) : s))}>
          Draw one
        </Button>
      </div>
      <p aria-live="polite" className="mt-2 text-sm text-muted-foreground" data-testid="hand-status">
        {status()}
      </p>
      {state && (
        <ul aria-label="Hand" className="mt-3 grid grid-cols-4 gap-2 sm:grid-cols-7">
          {state.hand.map((c, i) => {
            const face = facesOf(byId.get(c.oracleId))[0];
            const picture = face ? <FaceImage face={face} size="normal" /> : <span className="block aspect-[146/204] rounded-card border border-border bg-muted p-1 text-center text-xs">{c.name}</span>;
            return (
              <li key={`${c.oracleId}-${i}`} className="flex flex-col gap-1 text-xs" data-testid="hand-card">
                {state.toBottom > 0 ? (
                  <button type="button" onClick={() => setState((s) => (s ? bottom(s, i) : s))} className="rounded-card ring-offset-2 hover:ring-2 hover:ring-primary focus-visible:ring-2 focus-visible:ring-primary" aria-label={`Put ${c.name} on the bottom`}>
                    {picture}
                  </button>
                ) : (
                  picture
                )}
                <span className="wrap-anywhere">{c.name}</span>
              </li>
            );
          })}
        </ul>
      )}
    </section>
  );
}
