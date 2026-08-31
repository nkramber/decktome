import type { Deck } from "@mtg/api-client/mtg/v1/deck_pb";
import { useQuery } from "@tanstack/react-query";
import { GitCompareIcon } from "lucide-react";
import { useState } from "react";
import { useNavigate } from "react-router";

import { Button } from "../../components/ui/button";
import { Label } from "../../components/ui/label";
import { Skeleton } from "../../components/ui/skeleton";
import { deckClient } from "../../lib/api";
import { cn } from "../../lib/cn";
import { errorMessage } from "../../lib/errors";
import { DiffList, diffDecks } from "../deck/deck-diff";
import { useDeckVersions } from "../deck/use-decks";

// dateOf reads the day a deck was built. A deck with no stamp shows
// nothing rather than the epoch.
function dateOf(deck: Deck): string {
  const seconds = deck.createdAt?.seconds;
  return seconds ? new Date(Number(seconds) * 1000).toLocaleDateString() : "";
}

// The version history of a deck (PR-17). A revision turn writes a new
// deck in the same chat, so the decks of one chat are the versions of
// one deck, oldest first. Every version keeps its own address.
export function DeckVersions({ deck }: { deck: Deck }) {
  const navigate = useNavigate();
  const { versions, isPending, isError, error } = useDeckVersions(deck.sessionId);
  const [compare, setCompare] = useState(false);

  if (deck.sessionId === "") return null;
  if (isPending) return <Skeleton className="h-9 w-52" />;
  if (isError) {
    return (
      <p role="alert" className="text-sm text-danger">
        Could not read the versions: {errorMessage(error)}
      </p>
    );
  }
  // One version is no history. The deck is the only thing to read.
  if (versions.length < 2) return null;

  return (
    <section aria-labelledby="versions-title" className="flex flex-col gap-3 rounded-card border border-border bg-card p-4 shadow-card">
      <div className="flex flex-wrap items-center justify-between gap-3">
        <h2 id="versions-title" className="font-display text-[10px] tracking-[0.15em] text-muted-foreground uppercase">
          Versions ({versions.length})
        </h2>
        <Button variant="outline" size="sm" aria-expanded={compare} onClick={() => setCompare((v) => !v)}>
          <GitCompareIcon aria-hidden="true" />
          Compare
        </Button>
      </div>

      <ol className="flex flex-wrap gap-2" aria-label="Versions">
        {versions.map((v, i) => {
          const current = v.id === deck.id;
          return (
            <li key={v.id}>
              <button
                type="button"
                aria-current={current ? "page" : undefined}
                disabled={current}
                onClick={() => void navigate(`/decks/${v.id}`)}
                className={cn(
                  "flex items-baseline gap-2 rounded-card border px-3 py-1.5 text-left text-sm transition-colors",
                  current ? "border-primary bg-secondary" : "border-border hover:border-primary/60 hover:bg-secondary",
                )}
              >
                <span className="font-display text-xs tracking-wide">v{i + 1}</span>
                <span className="min-w-0 truncate">{v.name || v.id}</span>
                <span className="font-mono text-[10px] text-muted-foreground">{dateOf(v)}</span>
                {current && <span className="font-mono text-[10px] text-primary">now</span>}
              </button>
            </li>
          );
        })}
      </ol>

      {compare && <ComparePanel versions={versions} current={deck} />}
    </section>
  );
}

// The compare of two versions. The listing carries no cards, so each
// side reads its deck whole. The key is the one the deck screen uses,
// so the version on screen costs no call.
function ComparePanel({ versions, current }: { versions: Deck[]; current: Deck }) {
  const currentIndex = versions.findIndex((v) => v.id === current.id);
  const [fromId, setFromId] = useState(versions[Math.max(0, currentIndex - 1)]?.id ?? versions[0].id);
  const [toId, setToId] = useState(current.id);

  const from = useDeckWhole(fromId);
  const to = useDeckWhole(toId);
  const same = fromId === toId;

  return (
    <div className="flex flex-col gap-3 border-t border-border pt-3">
      <div className="flex flex-wrap items-end gap-3">
        <VersionSelect id="compare-from" label="Compare from" versions={versions} value={fromId} onChange={setFromId} />
        <VersionSelect id="compare-to" label="With" versions={versions} value={toId} onChange={setToId} />
      </div>

      {same ? (
        <p className="text-sm text-muted-foreground">Pick two different versions.</p>
      ) : from.isPending || to.isPending ? (
        <p role="status" className="text-sm text-muted-foreground">
          Reading both versions...
        </p>
      ) : from.isError || to.isError ? (
        <p role="alert" className="text-sm text-danger">
          Could not read both versions: {errorMessage(from.error ?? to.error)}
        </p>
      ) : from.data?.deck && to.data?.deck ? (
        <DiffList diff={diffDecks(from.data.deck, to.data.deck)} testId="compare-diff" />
      ) : null}
    </div>
  );
}

function useDeckWhole(id: string) {
  return useQuery({
    queryKey: ["deck", id],
    queryFn: () => deckClient.getDeck({ deckId: id }),
    enabled: id !== "",
  });
}

function VersionSelect({
  id,
  label,
  versions,
  value,
  onChange,
}: {
  id: string;
  label: string;
  versions: Deck[];
  value: string;
  onChange: (id: string) => void;
}) {
  return (
    <div className="flex items-center gap-2">
      <Label htmlFor={id} className="font-display text-[10px] tracking-[0.15em] text-muted-foreground uppercase">
        {label}
      </Label>
      <select
        id={id}
        value={value}
        onChange={(e) => onChange(e.target.value)}
        className="rounded-card border border-border bg-secondary px-3 py-1.5 text-sm transition-colors hover:border-primary/60 focus:border-primary focus:outline-none"
      >
        {versions.map((v, i) => (
          <option key={v.id} value={v.id}>
            v{i + 1} — {v.name || v.id}
          </option>
        ))}
      </select>
    </div>
  );
}
