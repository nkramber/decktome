import { FormatId, SixtyStep } from "@mtg/api-client/mtg/v1/format_pb";
import { LayersIcon, SearchIcon, StarIcon } from "lucide-react";
import { useEffect, useMemo, useState } from "react";
import { Link } from "react-router";

import { EmptyState } from "../../app/components/empty-state";
import { ErrorState } from "../../app/components/error-state";
import { notify } from "../../app/components/notify";
import { PageHeader } from "../../app/components/page-header";
import { Button } from "../../components/ui/button";
import { Input } from "../../components/ui/input";
import { Select } from "../../components/ui/select";
import { Skeleton } from "../../components/ui/skeleton";
import { cn } from "../../lib/cn";
import { errorMessage } from "../../lib/errors";
import { DeckCard } from "./deck-card";
import { type DeckFilter, emptyDeckFilter, useCommanderCards, useDeckList, useDeckWrites } from "./use-decks";

// The deck library (PR-17). The grid carries the art and the color of
// every deck, and the filters run on the server, so a match on a later
// page still shows (D-324).

// The search waits for a pause in the typing, so one word is one call.
function useDebounced<T>(value: T, ms: number): T {
  const [held, setHeld] = useState(value);
  useEffect(() => {
    const id = setTimeout(() => setHeld(value), ms);
    return () => clearTimeout(id);
  }, [value, ms]);
  return held;
}

const powerOptions = [
  { value: "", label: "Any power" },
  { value: "b1", label: "Bracket 1, exhibition" },
  { value: "b2", label: "Bracket 2, core" },
  { value: "b3", label: "Bracket 3, upgraded" },
  { value: "b4", label: "Bracket 4, optimized" },
  { value: "b5", label: "Bracket 5, competitive" },
  { value: "s1", label: "60-card casual" },
  { value: "s2", label: "60-card FNM" },
  { value: "s3", label: "60-card tournament" },
];

// powerFrom reads one select value into the two arms of the power filter.
export function powerFrom(value: string): { powerBracket: number; sixtyStep: SixtyStep } {
  if (value.startsWith("b")) return { powerBracket: Number(value.slice(1)), sixtyStep: SixtyStep.UNSPECIFIED };
  if (value.startsWith("s")) return { powerBracket: 0, sixtyStep: Number(value.slice(1)) as SixtyStep };
  return { powerBracket: 0, sixtyStep: SixtyStep.UNSPECIFIED };
}

export function DecksPage() {
  const [text, setText] = useState("");
  const [format, setFormat] = useState<FormatId>(FormatId.UNSPECIFIED);
  const [power, setPower] = useState("");
  const [favoritesOnly, setFavoritesOnly] = useState(false);
  const query = useDebounced(text.trim(), 250);

  const filter: DeckFilter = useMemo(
    () => ({ ...emptyDeckFilter, query, format, favorite: favoritesOnly ? true : undefined, ...powerFrom(power) }),
    [query, format, favoritesOnly, power],
  );

  const list = useDeckList(filter);
  const decks = useMemo(() => list.data?.pages.flatMap((p) => p.decks) ?? [], [list.data]);
  const byId = useCommanderCards(decks);
  const { setFavorite, remove } = useDeckWrites();
  const filtered = query !== "" || format !== FormatId.UNSPECIFIED || power !== "" || favoritesOnly;

  // The delete of the tile (D-439). The library refetches on success,
  // because the mutation invalidates the deck lists.
  async function onDelete(deckId: string, name: string) {
    try {
      await remove.mutateAsync({ deckId });
      await notify("success", "Deck deleted", name);
    } catch (err) {
      await notify("error", "Could not delete the deck", `${name}: ${errorMessage(err)}`);
    }
  }

  async function onFavorite(deckId: string, name: string, favorite: boolean) {
    try {
      await setFavorite.mutateAsync({ deckId, favorite });
    } catch (err) {
      await notify("error", favorite ? "Could not add the favorite" : "Could not remove the favorite", `${name}: ${errorMessage(err)}`);
    }
  }

  return (
    <div className="mx-auto flex w-full max-w-[75rem] flex-col gap-6 p-4 md:p-6">
      <PageHeader title="Your decks" description="Every deck the agent built for you." />

      <div className="flex flex-wrap items-center gap-2">
        <div className="relative min-w-56 flex-1">
          <SearchIcon aria-hidden="true" className="pointer-events-none absolute top-1/2 left-3 size-4 -translate-y-1/2 text-muted-foreground" />
          <Input aria-label="Search decks by name or commander" placeholder="Search by name or commander" value={text} onChange={(e) => setText(e.target.value)} className="pl-9" />
        </div>
        <Select aria-label="Filter by format" value={String(format)} onChange={(e) => setFormat(Number(e.target.value) as FormatId)}>
          <option value={FormatId.UNSPECIFIED}>Any format</option>
          <option value={FormatId.COMMANDER}>Commander</option>
          <option value={FormatId.STANDARD}>Standard</option>
          <option value={FormatId.MODERN}>Modern</option>
          <option value={FormatId.HOUSE}>House rules</option>
        </Select>
        <Select aria-label="Filter by power" value={power} onChange={(e) => setPower(e.target.value)}>
          {powerOptions.map((o) => (
            <option key={o.value} value={o.value}>
              {o.label}
            </option>
          ))}
        </Select>
        <Button variant={favoritesOnly ? "default" : "outline"} aria-pressed={favoritesOnly} onClick={() => setFavoritesOnly((v) => !v)}>
          <StarIcon className={cn("size-4", favoritesOnly && "fill-current")} aria-hidden="true" />
          Favorites
        </Button>
      </div>

      {list.isPending && (
        <div role="status" className="grid gap-4 sm:grid-cols-2 xl:grid-cols-3">
          <span className="sr-only">Loading decks...</span>
          <Skeleton className="h-56 w-full rounded-panel" />
          <Skeleton className="h-56 w-full rounded-panel" />
          <Skeleton className="h-56 w-full rounded-panel" />
        </div>
      )}
      {list.isError && <ErrorState title="Could not list decks" message={errorMessage(list.error)} onRetry={() => void list.refetch()} />}

      {list.isSuccess && decks.length === 0 && !filtered && (
        <EmptyState
          icon={LayersIcon}
          title="No decks yet."
          description="Start a chat and the agent builds your first deck."
          action={
            <Button asChild>
              <Link to="/session/new">Build a deck</Link>
            </Button>
          }
        />
      )}
      {list.isSuccess && decks.length === 0 && filtered && (
        <EmptyState
          icon={SearchIcon}
          title="No deck matches."
          description="Change the search or the filters."
          action={
            <Button
              variant="outline"
              onClick={() => {
                setText("");
                setFormat(FormatId.UNSPECIFIED);
                setPower("");
                setFavoritesOnly(false);
              }}
            >
              Clear the filters
            </Button>
          }
        />
      )}

      {decks.length > 0 && (
        <ul className="grid gap-4 sm:grid-cols-2 xl:grid-cols-3">
          {decks.map((d) => (
            <DeckCard
              key={d.id}
              deck={d}
              byId={byId}
              onFavorite={(favorite) => void onFavorite(d.id, d.name || d.id, favorite)}
              onDelete={() => void onDelete(d.id, d.name || d.id)}
            />
          ))}
        </ul>
      )}

      {list.hasNextPage && (
        <Button variant="outline" className="self-center" disabled={list.isFetchingNextPage} onClick={() => void list.fetchNextPage()}>
          {list.isFetchingNextPage ? "Loading..." : "Show more decks"}
        </Button>
      )}
    </div>
  );
}
