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

// The pool picker sits in the foot of the message box of a new chat
// (D-350). It names the cards the agent may use, beside the box that
// says what to build (D-37).
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
    <div className="flex flex-wrap items-center gap-x-3 gap-y-2">
      <Label htmlFor="pool-source" className="font-display shrink-0 text-[10px] tracking-[0.15em] text-muted-foreground uppercase">
        Build from
      </Label>
      {/* A native select takes the width of its widest option, so it grew
          the moment the collections arrived. A fixed width keeps the foot
          of the message box still. */}
      <div className="relative">
        <select
          id="pool-source"
          data-testid="pool-source"
          value={collectionId}
          onChange={(e) => choose(e.target.value)}
          className="w-44 appearance-none truncate rounded-card border border-border bg-secondary py-1 pr-7 pl-2.5 text-xs transition-colors hover:border-primary/60 focus:border-primary focus:outline-none"
        >
          <option value="">Any card</option>
          {collections.map((c) => (
            <option key={c.id} value={c.id}>
              {c.name} ({c.cardCount} cards)
            </option>
          ))}
          <option value={addValue}>Add a collection...</option>
        </select>
        <ChevronDownIcon aria-hidden="true" className="pointer-events-none absolute top-1/2 right-2 size-3 -translate-y-1/2 text-muted-foreground" />
      </div>

      {collectionId ? (
        <Label className="flex items-center gap-2 text-xs font-normal">
          <Checkbox checked={poolMode === "owned_only"} onCheckedChange={(v) => setPoolMode(v === true ? "owned_only" : "owned_first")} />
          Only cards I own
        </Label>
      ) : null}
    </div>
  );
}
