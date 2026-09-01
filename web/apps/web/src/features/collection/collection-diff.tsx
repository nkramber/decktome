import type { CollectionDiff } from "@mtg/api-client/mtg/v1/collection_pb";
import { CheckIcon } from "lucide-react";

import { Button } from "../../components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "../../components/ui/card";

// The reader reads what a re-upload changes before anything replaces
// anything (D-393). A row is one printing in one finish and one
// condition, which is the row a ManaBox export writes, so a foil and a
// normal copy of one card are two rows.

// rowsShown bounds each list. A reader reads the counts and a sample,
// and a diff of two thousand rows is not a list anyone reads.
const rowsShown = 8;

export function CollectionDiffView({
  name,
  diff,
  pending,
  onReplace,
  onCancel,
}: {
  name: string;
  diff: CollectionDiff;
  pending: boolean;
  onReplace: () => void;
  onCancel: () => void;
}) {
  if (diff.identical) {
    return (
      <Card>
        <CardHeader>
          <CardTitle>Nothing changed</CardTitle>
        </CardHeader>
        <CardContent className="flex flex-col gap-3">
          <p className="flex items-center gap-2 text-sm">
            <CheckIcon aria-hidden="true" className="size-4 text-success" />
            This file holds the same cards as {name}, in the same numbers.
          </p>
          <div className="flex flex-wrap gap-2">
            <Button type="button" variant="outline" onClick={onCancel}>
              Keep what I have
            </Button>
          </div>
        </CardContent>
      </Card>
    );
  }

  return (
    <Card data-testid="collection-diff">
      <CardHeader>
        <CardTitle>What changes in {name}</CardTitle>
      </CardHeader>
      <CardContent className="flex flex-col gap-4">
        <dl className="flex flex-wrap gap-x-8 gap-y-2 text-sm">
          <Count label="Added" rows={diff.added.length} cards={diff.addedCards} testid="diff-added" />
          <Count label="Removed" rows={diff.removed.length} cards={diff.removedCards} testid="diff-removed" />
          <Count label="Changed" rows={diff.changed.length} cards={diff.changedCards} testid="diff-changed" />
        </dl>

        <div className="flex flex-col gap-3">
          <Sample title="Added" names={diff.added.slice(0, rowsShown).map((e) => `${e.quantity}× ${e.name}`)} more={diff.added.length - rowsShown} />
          <Sample title="Removed" names={diff.removed.slice(0, rowsShown).map((e) => `${e.quantity}× ${e.name}`)} more={diff.removed.length - rowsShown} />
          <Sample
            title="Changed"
            names={diff.changed.slice(0, rowsShown).map((c) => `${c.entry?.name ?? ""}: ${c.from} → ${c.to}`)}
            more={diff.changed.length - rowsShown}
          />
        </div>

        <div className="flex flex-wrap gap-2">
          <Button type="button" onClick={onReplace} disabled={pending}>
            {pending ? "Replacing..." : "Replace the collection"}
          </Button>
          <Button type="button" variant="outline" onClick={onCancel} disabled={pending}>
            Keep what I have
          </Button>
        </div>
        <p className="text-sm text-muted-foreground">
          Replace keeps the same collection, so every deck and chat that names it still works. Only the cards change.
        </p>
      </CardContent>
    </Card>
  );
}

function Count({ label, rows, cards, testid }: { label: string; rows: number; cards: number; testid: string }) {
  return (
    <div className="flex flex-col">
      <dt className="text-muted-foreground">{label}</dt>
      <dd className="text-lg font-medium tabular-nums" data-testid={testid}>
        {rows.toLocaleString()} {rows === 1 ? "row" : "rows"}
        <span className="ml-2 font-mono text-[11px] text-muted-foreground">{cards.toLocaleString()} cards</span>
      </dd>
    </div>
  );
}

function Sample({ title, names, more }: { title: string; names: string[]; more: number }) {
  if (names.length === 0) return null;
  return (
    <div className="flex flex-col gap-1">
      <h3 className="font-display text-sm font-semibold">{title}</h3>
      <ul className="flex flex-col gap-0.5 font-mono text-[11px] text-muted-foreground">
        {names.map((n) => (
          <li key={n}>{n}</li>
        ))}
        {more > 0 && <li className="text-muted-foreground/70">and {more.toLocaleString()} more</li>}
      </ul>
    </div>
  );
}
