import { type Card, Color } from "@mtg/api-client/mtg/v1/card_pb";
import { type CollectionEntry, Condition, Finish } from "@mtg/api-client/mtg/v1/collection_pb";
import { useVirtualizer } from "@tanstack/react-virtual";
import { SearchIcon } from "lucide-react";
import { useEffect, useMemo, useRef, useState } from "react";

import { EmptyState } from "../../app/components/empty-state";
import { Input } from "../../components/ui/input";
import { Select } from "../../components/ui/select";
import { Skeleton } from "../../components/ui/skeleton";
import { cn } from "../../lib/cn";
import { useCollectionArt } from "./use-collection";

// The binder: the reader's cards, with the art, the count, the finish,
// and the condition. It virtualizes the rows, because the owner's export
// holds 2,657 of them and every one carries an image (D-394).

// The rarity of a printing carries the color the game prints it in. The
// hero uses the same four tokens.
const rarityToken: Record<string, string> = {
  mythic: "var(--rarity-mythic)",
  rare: "var(--rarity-rare)",
  uncommon: "var(--rarity-uncommon)",
  common: "var(--rarity-common)",
};

const finishLabel: Record<number, string> = {
  [Finish.FOIL]: "Foil",
  [Finish.ETCHED]: "Etched",
};

// conditionLabel reads a Condition as a person writes it. The proto
// names it NEAR_MINT, and a reader reads "Near mint". An unset
// condition reads as nothing, because a ManaBox export may omit it.
const conditionLabel: Record<number, string> = {
  [Condition.MINT]: "Mint",
  [Condition.NEAR_MINT]: "Near mint",
  [Condition.EXCELLENT]: "Excellent",
  [Condition.GOOD]: "Good",
  [Condition.LIGHT_PLAYED]: "Light played",
  [Condition.PLAYED]: "Played",
  [Condition.POOR]: "Poor",
};

export type BinderSort = "name" | "count" | "set" | "price";

const sortOptions: { value: BinderSort; label: string }[] = [
  { value: "name", label: "By name" },
  { value: "count", label: "By count" },
  { value: "set", label: "By set" },
  { value: "price", label: "By price" },
];

// The color filter reads the colors of the card, not its color identity.
// A reader who asks for red cards means the red cards, and "Colorless"
// is a color a reader picks like any other.
export const colorOptions: { value: string; label: string }[] = [
  { value: "", label: "Every color" },
  { value: "W", label: "White" },
  { value: "U", label: "Blue" },
  { value: "B", label: "Black" },
  { value: "R", label: "Red" },
  { value: "G", label: "Green" },
  { value: "C", label: "Colorless" },
];

const colorKeys: Record<number, string> = {
  [Color.W]: "W",
  [Color.U]: "U",
  [Color.B]: "B",
  [Color.R]: "R",
  [Color.G]: "G",
};

// The count filter groups the copies a reader thinks in. Four is a
// playset, and every count above it reads the same way.
export const countOptions: { value: string; label: string }[] = [
  { value: "", label: "Every count" },
  { value: "1", label: "1 copy" },
  { value: "2", label: "2 copies" },
  { value: "3", label: "3 copies" },
  { value: "4", label: "4 or more" },
];

// BinderFilter is what the reader chose. An empty string is no choice,
// which keeps every row.
export type BinderFilter = {
  query: string;
  set: string;
  color: string;
  type: string;
  count: string;
};

export const noFilter: BinderFilter = { query: "", set: "", color: "", type: "", count: "" };

// filterEntries keeps the rows that match every choice. The search reads
// the card name alone: a reader types a card, not a code. The colors and
// the card types ride on the entry, and the server fills them from the
// card index of the day (D-396).
export function filterEntries(entries: CollectionEntry[], f: BinderFilter): CollectionEntry[] {
  const q = f.query.trim().toLowerCase();
  const want = Number(f.count);
  return entries.filter((e) => {
    if (f.set !== "" && e.setCode !== f.set) return false;
    if (f.color !== "" && !hasColor(e, f.color)) return false;
    if (f.type !== "" && !e.cardTypes.includes(f.type)) return false;
    if (f.count !== "" && !(want === 4 ? e.quantity >= 4 : e.quantity === want)) return false;
    if (q === "") return true;
    return e.name.toLowerCase().includes(q);
  });
}

// hasColor answers the color filter. A card of two colors answers to
// both. A card of none answers to "Colorless" alone.
function hasColor(e: CollectionEntry, want: string): boolean {
  if (want === "C") return e.colors.length === 0;
  return e.colors.some((c) => colorKeys[c] === want);
}

// sortEntries orders the rows. The name breaks every tie, so one binder
// reads the same on every visit.
export function sortEntries(entries: CollectionEntry[], by: BinderSort): CollectionEntry[] {
  const out = [...entries];
  out.sort((a, b) => {
    if (by === "count" && a.quantity !== b.quantity) return b.quantity - a.quantity;
    if (by === "set" && a.setCode !== b.setCode) return a.setCode.localeCompare(b.setCode);
    // Dearest first, because a reader who sorts by price looks for the
    // cards that carry the value. An unpriced row sorts last.
    if (by === "price" && a.priceUsd !== b.priceUsd) return b.priceUsd - a.priceUsd;
    return a.name.localeCompare(b.name);
  });
  return out;
}

// typeOptions lists the card types the rows hold, largest first. The
// binder filters on what it holds, so a binder with no planeswalker
// offers no planeswalker.
export function typeOptions(entries: CollectionEntry[]): { value: string; label: string }[] {
  const byType = new Map<string, number>();
  for (const e of entries) {
    for (const t of e.cardTypes) {
      byType.set(t, (byType.get(t) ?? 0) + e.quantity);
    }
  }
  return [
    { value: "", label: "Every type" },
    ...[...byType.entries()]
      .sort((a, b) => b[1] - a[1] || a[0].localeCompare(b[0]))
      .map(([t, n]) => ({ value: t, label: `${t} (${n})` })),
  ];
}

// setOptions lists the sets the rows hold, largest first. The binder
// filters on what it holds, so it needs no set table (PR-17B).
export function setOptions(entries: CollectionEntry[]): { value: string; label: string }[] {
  const byCode = new Map<string, { name: string; count: number }>();
  for (const e of entries) {
    if (!e.setCode) continue;
    const seen = byCode.get(e.setCode);
    byCode.set(e.setCode, { name: e.setName || e.setCode, count: (seen?.count ?? 0) + e.quantity });
  }
  return [
    { value: "", label: "Every set" },
    ...[...byCode.entries()]
      .sort((a, b) => b[1].count - a[1].count || a[0].localeCompare(b[0]))
      .map(([code, s]) => ({ value: code, label: `${s.name} (${s.count})` })),
  ];
}

// rowsPerPage is how many tiles one virtual row holds at each width. The
// grid is three columns wide on a desktop, and the virtualizer measures
// rows, so the columns are counted here.
function columnsFor(width: number): number {
  if (width >= 1280) return 4;
  if (width >= 768) return 3;
  if (width >= 480) return 2;
  return 1;
}

export function BinderGrid({
  entries,
  loading,
  onReachEnd,
}: {
  entries: CollectionEntry[];
  loading: boolean;
  onReachEnd?: () => void;
}) {
  const [filter, setFilter] = useState<BinderFilter>(noFilter);
  const [by, setBy] = useState<BinderSort>("name");
  const scroller = useRef<HTMLDivElement>(null);
  const [columns, setColumns] = useState(3);

  const sets = useMemo(() => setOptions(entries), [entries]);
  const types = useMemo(() => typeOptions(entries), [entries]);
  const shown = useMemo(() => sortEntries(filterEntries(entries, filter), by), [entries, filter, by]);

  // The column count follows the width of the scroller, so the tiles
  // keep their shape from a phone to a wide desktop.
  useEffect(() => {
    const el = scroller.current;
    if (!el || typeof ResizeObserver === "undefined") return;
    const measure = () => setColumns(columnsFor(el.clientWidth));
    measure();
    const observer = new ResizeObserver(measure);
    observer.observe(el);
    return () => observer.disconnect();
  }, []);

  // A new filter or sort is a new list, so the reader reads it from the
  // top. Without this they search and land in the middle of the answer.
  useEffect(() => {
    scroller.current?.scrollTo(0, 0);
  }, [filter, by]);

  const rows = Math.ceil(shown.length / columns);
  // The React Compiler skips this component, because the virtualizer
  // returns functions it can not memoize safely. The grid holds no
  // memoized child that reads them, so the skip costs nothing here.
  // eslint-disable-next-line react-hooks/incompatible-library
  const virtual = useVirtualizer({
    count: rows,
    getScrollElement: () => scroller.current,
    // A tile is the art plus two lines of text. The virtualizer measures
    // the real height after the first paint, so this is the start value.
    estimateSize: () => 312,
    overscan: 3,
    // The library flushes a scroll render synchronously by default, and
    // React refuses a flush from inside a lifecycle. Every scroll wrote
    // console errors that hide real ones. A frame test over 120 frames
    // measured the same times either way, so the flush buys nothing
    // here (D-397).
    useFlushSync: false,
  });

  // The next page loads before the reader reaches the end, so the scroll
  // does not stop while it waits.
  const items = virtual.getVirtualItems();
  const lastRow = items[items.length - 1]?.index ?? 0;
  useEffect(() => {
    if (onReachEnd && rows > 0 && lastRow >= rows - 2) onReachEnd();
  }, [lastRow, rows, onReachEnd]);

  const artIds = useMemo(
    () => [...new Set(items.flatMap((v) => shown.slice(v.index * columns, v.index * columns + columns).map((e) => e.oracleId)))].filter(Boolean),
    [items, shown, columns],
  );
  const cards = useCollectionArt(artIds);
  const byOracle = useMemo(() => new Map(cards.map((c) => [c.oracleId, c])), [cards]);

  return (
    <section aria-labelledby="binder-grid-title" className="flex flex-col gap-3">
      <div className="flex flex-wrap items-center gap-3">
        <h3 id="binder-grid-title" className="font-display text-lg font-semibold">
          The binder
        </h3>
        <span className="font-mono text-[11px] text-muted-foreground tabular-nums" data-testid="binder-count">
          {shown.length.toLocaleString()} of {entries.length.toLocaleString()} rows
        </span>
      </div>

      <div className="flex flex-wrap gap-2">
        <div className="relative min-w-48 grow">
          <SearchIcon aria-hidden="true" className="pointer-events-none absolute top-1/2 left-2.5 size-4 -translate-y-1/2 text-muted-foreground" />
          <Input
            type="search"
            aria-label="Search the binder by card name"
            placeholder="Search by name"
            value={filter.query}
            onChange={(e) => setFilter({ ...filter, query: e.target.value })}
            className="pl-8"
          />
        </div>
        <Select aria-label="Filter by set" value={filter.set} onChange={(e) => setFilter({ ...filter, set: e.target.value })}>
          {sets.map((o) => (
            <option key={o.value} value={o.value}>
              {o.label}
            </option>
          ))}
        </Select>
        <Select aria-label="Filter by color" value={filter.color} onChange={(e) => setFilter({ ...filter, color: e.target.value })}>
          {colorOptions.map((o) => (
            <option key={o.value} value={o.value}>
              {o.label}
            </option>
          ))}
        </Select>
        <Select aria-label="Filter by card type" value={filter.type} onChange={(e) => setFilter({ ...filter, type: e.target.value })}>
          {types.map((o) => (
            <option key={o.value} value={o.value}>
              {o.label}
            </option>
          ))}
        </Select>
        <Select aria-label="Filter by count" value={filter.count} onChange={(e) => setFilter({ ...filter, count: e.target.value })}>
          {countOptions.map((o) => (
            <option key={o.value} value={o.value}>
              {o.label}
            </option>
          ))}
        </Select>
        <Select aria-label="Sort the binder" value={by} onChange={(e) => setBy(e.target.value as BinderSort)}>
          {sortOptions.map((o) => (
            <option key={o.value} value={o.value}>
              {o.label}
            </option>
          ))}
        </Select>
      </div>

      {shown.length === 0 && !loading && (
        <EmptyState icon={SearchIcon} title="No card matches." description="Clear the search, or choose another filter." />
      )}

      {shown.length > 0 && (
        <div ref={scroller} data-testid="binder-scroller" className="max-h-[75vh] overflow-y-auto rounded-panel border border-border bg-card p-2">
          <div style={{ height: virtual.getTotalSize(), position: "relative", width: "100%" }}>
            {items.map((row) => (
              <div
                key={row.key}
                data-index={row.index}
                ref={virtual.measureElement}
                className={cn("absolute top-0 left-0 grid w-full gap-2 pb-2", gridCols[columns])}
                style={{ transform: `translateY(${row.start}px)` }}
              >
                {shown.slice(row.index * columns, row.index * columns + columns).map((e) => (
                  <BinderTile key={`${e.scryfallId}-${e.finish}-${e.condition}`} entry={e} card={byOracle.get(e.oracleId)} />
                ))}
              </div>
            ))}
          </div>
        </div>
      )}

      {loading && <Skeleton className="h-24 w-full rounded-card" />}
    </section>
  );
}

// gridCols names the column count as a class, because Tailwind reads the
// class names it compiles and never a computed string.
const gridCols: Record<number, string> = {
  1: "grid-cols-1",
  2: "grid-cols-2",
  3: "grid-cols-3",
  4: "grid-cols-4",
};

// BinderTile is one row of the binder: the art, the name, the count, and
// the marks the reader sorted their cards by. It wears the deck card
// tile's shape, so the two screens read alike.
function BinderTile({ entry, card }: { entry: CollectionEntry; card: Card | undefined }) {
  const art = card?.faces?.[0]?.imageUris?.normal ?? card?.defaultPrinting?.imageUris?.normal ?? "";
  const finish = finishLabel[entry.finish];
  const condition = conditionLabel[entry.condition];
  return (
    <article className="flex flex-col gap-1.5 rounded-card border border-border bg-card p-2 transition-shadow hover:shadow-raised" data-testid="binder-tile">
      <div className="flex items-baseline justify-between gap-2">
        <span className="min-w-0 wrap-anywhere font-medium">
          {entry.quantity > 1 && <span className="mr-1 text-muted-foreground">{entry.quantity}×</span>}
          {entry.name}
        </span>
        <span
          aria-hidden="true"
          className="mt-1 inline-block size-2 shrink-0 rounded-full"
          style={{ backgroundColor: rarityToken[entry.rarity] ?? "var(--rarity-common)" }}
        />
      </div>
      {art ? (
        <img src={art} alt={`${entry.name} (card)`} loading="lazy" width={488} height={680} className="h-auto w-full rounded-card" />
      ) : (
        <div className="flex aspect-[488/680] items-center justify-center rounded-card bg-muted text-xs text-muted-foreground">No card data</div>
      )}
      <p className="flex flex-wrap items-center gap-x-2 gap-y-1 font-mono text-[11px] text-muted-foreground">
        <span className="uppercase">{entry.setCode}</span>
        <span aria-hidden="true">·</span>
        <span>{entry.collectorNumber}</span>
        {finish && (
          <span className="rounded-md border border-primary/50 bg-primary/15 px-1.5 py-0.5 font-medium text-primary" data-testid="finish-mark">
            {finish}
          </span>
        )}
        {condition && <span data-testid="condition-mark">{condition}</span>}
        {/* The reader sorts by price, so the price reads on the tile.
            An unpriced row shows nothing (D-396). */}
        {entry.priceUsd > 0 && (
          <span className="ml-auto tabular-nums" data-testid="binder-price">
            ${entry.priceUsd.toFixed(2)}
          </span>
        )}
      </p>
    </article>
  );
}
