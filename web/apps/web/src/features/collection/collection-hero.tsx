import type { Collection } from "@mtg/api-client/mtg/v1/collection_pb";

import { Skeleton } from "../../components/ui/skeleton";
import { statsFrom, useCollectionArt } from "./use-collection";

// The rarity of a printing carries the color the game prints it in.
const rarityToken: Record<string, string> = {
  mythic: "var(--rarity-mythic)",
  rare: "var(--rarity-rare)",
  uncommon: "var(--rarity-uncommon)",
  common: "var(--rarity-common)",
};

// The head of the collection screen: what the binder holds, and the look
// of it. A collection of five thousand cards must show cards (D-327).
//
// It reads the summary the import stored, and never an entry (D-392).
// The art ids ride on the summary for the same reason.
export function CollectionHero({ collection, loading }: { collection: Collection; loading: boolean }) {
  const stats = statsFrom(collection.summary, collection.cardCount);
  const cards = useCollectionArt(collection.summary?.artOracleIds ?? []);
  const total = stats.total || collection.cardCount;

  return (
    <section aria-labelledby="binder-title" className="shadow-card relative isolate overflow-hidden rounded-panel border border-border bg-card backdrop-blur-sm">
      <div aria-hidden="true" className="pointer-events-none absolute inset-0 flex opacity-55">
        {cards.map((card) => {
          const art = card.faces?.[0]?.imageUris?.artCrop ?? card.defaultPrinting?.imageUris?.artCrop ?? "";
          if (!art) return null;
          return <img key={card.oracleId} src={art} alt="" loading="lazy" className="h-full min-w-0 flex-1 object-cover" />;
        })}
      </div>
      <div aria-hidden="true" className="pointer-events-none absolute inset-0 bg-gradient-to-r from-card via-card/90 to-card/60" />

      <div className="relative flex flex-col gap-5 p-6">
        <div className="flex flex-wrap items-end justify-between gap-4">
          <div className="flex flex-col gap-1">
            <h2 id="binder-title" className="text-sm font-medium tracking-wide text-muted-foreground uppercase">
              {collection.name}
            </h2>
            <p className="flex items-baseline gap-2">
              <span className="text-4xl font-semibold tracking-tight tabular-nums">{total.toLocaleString()}</span>
              <span className="text-muted-foreground">cards</span>
            </p>
          </div>
          <dl className="flex flex-wrap gap-x-8 gap-y-2 text-sm">
            <div className="flex flex-col">
              <dt className="text-muted-foreground">Unique cards</dt>
              <dd className="text-lg font-medium tabular-nums">{loading ? <Skeleton className="h-6 w-16" /> : stats.unique.toLocaleString()}</dd>
            </div>
            {stats.sets[0] && (
              <div className="flex flex-col">
                <dt className="text-muted-foreground">Biggest set</dt>
                <dd className="max-w-48 truncate text-lg font-medium">{stats.sets[0].name}</dd>
              </div>
            )}
            {collection.importedAt?.seconds ? (
              <div className="flex flex-col">
                <dt className="text-muted-foreground">Imported</dt>
                <dd className="text-lg font-medium">{new Date(Number(collection.importedAt.seconds) * 1000).toLocaleDateString()}</dd>
              </div>
            ) : null}
          </dl>
        </div>

        {stats.rarity.length > 0 && (
          <div className="flex flex-col gap-2">
            <div className="flex h-2 overflow-hidden rounded-full bg-muted" aria-hidden="true">
              {stats.rarity.map((r) => (
                <span key={r.key} style={{ width: `${(r.count / stats.total) * 100}%`, backgroundColor: rarityToken[r.key] }} />
              ))}
            </div>
            <ul className="flex flex-wrap gap-x-4 gap-y-1 text-xs text-muted-foreground">
              {stats.rarity.map((r) => (
                <li key={r.key} className="flex items-center gap-1.5">
                  <span aria-hidden="true" className="inline-block size-2 rounded-full" style={{ backgroundColor: rarityToken[r.key] }} />
                  {r.label} <span className="tabular-nums">{r.count.toLocaleString()}</span>
                </li>
              ))}
            </ul>
          </div>
        )}
      </div>
    </section>
  );
}
