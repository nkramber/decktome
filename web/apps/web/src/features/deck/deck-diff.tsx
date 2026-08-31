import type { Deck } from "@mtg/api-client/mtg/v1/deck_pb";

import { type DeckDiff, diffDecks } from "./deck-stats";

// One reading of a diff between two decks (PR-12B, PR-17). The revision
// note of a deck shows it, and the compare of two versions shows it.
export function DiffList({ diff, testId }: { diff: DeckDiff; testId?: string }) {
  if (diff.removed.length === 0 && diff.added.length === 0 && diff.changed.length === 0) {
    return (
      <p className="text-sm text-muted-foreground" data-testid={testId}>
        The two decks hold the same cards.
      </p>
    );
  }
  return (
    <ul className="mt-2 list-disc pl-5" data-testid={testId}>
      {diff.removed.map((x) => (
        <li key={`r-${x}`}>Removed {x}</li>
      ))}
      {diff.added.map((x) => (
        <li key={`a-${x}`}>Added {x}</li>
      ))}
      {diff.changed.map((x) => (
        <li key={`c-${x}`}>Count of {x}</li>
      ))}
    </ul>
  );
}

export { diffDecks };
export type { Deck, DeckDiff };
