import type { CollectionDiff } from "@mtg/api-client/mtg/v1/collection_pb";

// The reader reads what a re-upload changes before anything replaces
// anything (D-393). A row is one printing in one finish and one
// condition, which is the row a ManaBox export writes, so a foil and a
// normal copy of one card are two rows.
//
// The upload dialog holds the frame and the buttons, so this file draws
// the counts alone. A card inside the dialog would be a frame in a
// frame (roadmap PR-18).

// rowsShown bounds each list. A reader reads the counts and a sample,
// and a diff of two thousand rows is not a list anyone reads.
const rowsShown = 8;

export function CollectionDiffBody({ diff }: { diff: CollectionDiff }) {
  // A file that changes nothing has no counts to read. The head of the
  // dialog says so, and the one button keeps what the reader has.
  if (diff.identical) return null;

  return (
    <div className="flex flex-col gap-4" data-testid="collection-diff">
      <dl className="flex flex-wrap gap-x-8 gap-y-2 text-sm">
        <Count label="Added" rows={diff.added.length} cards={diff.addedCards} testid="diff-added" />
        <Count label="Removed" rows={diff.removed.length} cards={diff.removedCards} testid="diff-removed" />
        <Count label="Changed" rows={diff.changed.length} cards={diff.changedCards} testid="diff-changed" />
      </dl>

      {/* A long diff scrolls inside itself, so the dialog keeps its
          height and the buttons stay where the reader left them. */}
      <div className="flex max-h-72 flex-col gap-3 overflow-auto">
        <Sample title="Added" names={diff.added.slice(0, rowsShown).map((e) => `${e.quantity}× ${e.name}`)} more={diff.added.length - rowsShown} />
        <Sample title="Removed" names={diff.removed.slice(0, rowsShown).map((e) => `${e.quantity}× ${e.name}`)} more={diff.removed.length - rowsShown} />
        <Sample
          title="Changed"
          names={diff.changed.slice(0, rowsShown).map((c) => `${c.entry?.name ?? ""}: ${c.from} → ${c.to}`)}
          more={diff.changed.length - rowsShown}
        />
      </div>

      <p className="text-sm text-muted-foreground">
        Replace keeps the same collection, so every deck and chat that names it still works. Only the cards change.
      </p>
    </div>
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
