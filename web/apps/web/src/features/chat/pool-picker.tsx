import { useQuery } from "@tanstack/react-query";
import { ChevronDownIcon } from "lucide-react";
import { useNavigate } from "react-router";

import { Checkbox } from "../../components/ui/checkbox";
import { Label } from "../../components/ui/label";
import { collectionClient } from "../../lib/api";
import { useAppStore } from "../../lib/store";

// The value that leaves the chat for the upload page. No collection id
// can hold it, because the server writes an opaque id with no colon.
const addValue = "add:collection";

// useCollections reads the uploads of the user. The Build menu of the
// header reads the same key, so one of them pays for the call.
export function useCollections(enabled = true) {
  return useQuery({
    queryKey: ["collections"],
    queryFn: () => collectionClient.listCollections({}),
    enabled,
  });
}

// The pool picker is the first control of a new chat. It names the cards
// the agent may use, and it changes them without a trip to another page
// (D-37). The header menu does the same from any screen.
export function PoolPicker() {
  const navigate = useNavigate();
  const collectionId = useAppStore((s) => s.collectionId);
  const poolMode = useAppStore((s) => s.poolMode);
  const setCollection = useAppStore((s) => s.setCollection);
  const clearCollection = useAppStore((s) => s.clearCollection);
  const setPoolMode = useAppStore((s) => s.setPoolMode);

  const list = useCollections();
  const collections = list.data?.collections ?? [];

  function choose(value: string) {
    if (value === addValue) {
      // The reader asked for a new collection, so the old one is not the
      // answer. A cancelled file dialog therefore leaves no collection
      // active, and the upload screen opens its file dialog at once.
      clearCollection();
      void navigate("/collection", { state: { pickFile: true } });
      return;
    }
    if (value === "") clearCollection();
    else setCollection(value);
  }

  return (
    <div className="flex w-fit max-w-full flex-wrap items-center gap-x-4 gap-y-3 rounded-card border border-border bg-card px-4 py-3 shadow-card">
      <div className="flex items-center gap-2.5">
        <Label htmlFor="pool-source" className="font-display shrink-0 text-[10px] tracking-[0.15em] text-muted-foreground uppercase">
          Build from
        </Label>
        <div className="relative">
          <select
            id="pool-source"
            data-testid="pool-source"
            value={collectionId}
            onChange={(e) => choose(e.target.value)}
            className="appearance-none rounded-card border border-border bg-secondary py-1.5 pr-8 pl-3 text-sm transition-colors hover:border-primary/60 focus:border-primary focus:outline-none"
          >
            <option value="">Any card</option>
            {collections.map((c) => (
              <option key={c.id} value={c.id}>
                {c.name} ({c.cardCount} cards)
              </option>
            ))}
            <option value={addValue}>Add a collection...</option>
          </select>
          <ChevronDownIcon aria-hidden="true" className="pointer-events-none absolute top-1/2 right-2.5 size-3.5 -translate-y-1/2 text-muted-foreground" />
        </div>
      </div>

      {collectionId ? (
        <Label className="flex items-center gap-2 text-sm font-normal">
          <Checkbox checked={poolMode === "owned"} onCheckedChange={(v) => setPoolMode(v === true ? "owned" : "any")} />
          Use only cards in my collection
        </Label>
      ) : (
        list.isSuccess &&
        collections.length === 0 && <p className="text-sm text-muted-foreground">No collection yet. The agent builds from any card until you add one.</p>
      )}
    </div>
  );
}
