import type { DeckCard } from "@mtg/api-client/mtg/v1/deck_pb";
import type { ImportDeckResponse } from "@mtg/api-client/mtg/v1/agent_service_pb";
import { FormatId } from "@mtg/api-client/mtg/v1/format_pb";
import type { UnresolvedRow } from "@mtg/api-client/mtg/v1/collection_pb";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { FileTextIcon } from "lucide-react";
import { type FormEvent, useState } from "react";
import { useNavigate } from "react-router";

import { Button } from "../../components/ui/button";
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle } from "../../components/ui/dialog";
import { Input } from "../../components/ui/input";
import { Label } from "../../components/ui/label";
import { agentClient, collectionClient } from "../../lib/api";
import { cn } from "../../lib/cn";
import { errorMessage } from "../../lib/errors";

// The deck import (PR-70). A reader brings a deck list: a text file that
// Archidekt exports, or a pasted Arena list (D-845). The app stores it as
// a deck it built, so the deck page shows it and the chat revises it.
//
// A list can need one answer before it stores. A list that is not
// Commander asks Standard, Modern, or neither (D-857), and a Commander
// list with no commander mark asks which card leads (D-847). The dialog
// sends the same list again with the answer.

// maxImportBytes matches the cap of the server.
export const maxImportBytes = 128 * 1024;

type Step = "pick" | "format" | "commander" | "done";

// nameFromFile turns "living_weapon.txt" into "living weapon".
export function nameFromFile(file: string): string {
  return file
    .replace(/\.[^.]+$/, "")
    .replace(/[_-]+/g, " ")
    .trim();
}

export function ImportDialog({ open, onOpenChange }: { open: boolean; onOpenChange: (open: boolean) => void }) {
  // A new key on each open remounts the body, so a second open starts
  // clean, as the collection upload does.
  const [wasOpen, setWasOpen] = useState(open);
  const [openCount, setOpenCount] = useState(0);
  if (open !== wasOpen) {
    setWasOpen(open);
    if (open) setOpenCount((n) => n + 1);
  }
  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <ImportBody key={openCount} onClose={() => onOpenChange(false)} />
    </Dialog>
  );
}

function ImportBody({ onClose }: { onClose: () => void }) {
  const navigate = useNavigate();
  const queryClient = useQueryClient();
  // The key is the one of the pool picker of the chat, so one call serves
  // both.
  const collections = useQuery({ queryKey: ["collections"], queryFn: () => collectionClient.listCollections({}) }).data?.collections ?? [];
  const [step, setStep] = useState<Step>("pick");
  const [text, setText] = useState("");
  const [fileName, setFileName] = useState("");
  const [name, setName] = useState("");
  const [collectionId, setCollectionId] = useState("");
  const [format, setFormat] = useState<FormatId>(FormatId.UNSPECIFIED);
  const [options, setOptions] = useState<DeckCard[]>([]);
  const [picked, setPicked] = useState("");
  const [result, setResult] = useState<ImportDeckResponse | null>(null);

  const tooLarge = new TextEncoder().encode(text).length > maxImportBytes;

  const send = useMutation({
    mutationFn: (answer: { format?: FormatId; commander?: string }) =>
      agentClient.importDeck({
        text,
        name: name.trim() || nameFromFile(fileName),
        collectionId,
        format: answer.format ?? format,
        commanderOracleIds: answer.commander ? [answer.commander] : [],
      }),
    onSuccess: async (res) => {
      if (res.needsFormat) {
        setStep("format");
        return;
      }
      if (res.commanderOptions.length > 0) {
        setOptions(res.commanderOptions);
        setPicked(res.commanderOptions[0]?.oracleId ?? "");
        setStep("commander");
        return;
      }
      await queryClient.invalidateQueries({ queryKey: ["decks"] });
      // A list with a line that matched no card says so before the deck
      // opens (D-846).
      if (res.unresolved.length > 0) {
        setResult(res);
        setStep("done");
        return;
      }
      onClose();
      void navigate(`/decks/${res.deck?.id ?? ""}`);
    },
  });

  async function readFile(f: File | undefined) {
    if (!f) return;
    setFileName(f.name);
    setText(await f.text());
  }

  function onSubmit(e: FormEvent) {
    e.preventDefault();
    if (text.trim() === "" || tooLarge || send.isPending) return;
    send.mutate({});
  }

  const title = step === "format" ? "Which format?" : step === "commander" ? "Which card leads the deck?" : step === "done" ? "Lines the import skipped" : "Import a deck";
  const note =
    step === "format"
      ? "This list is not a Commander deck. Pick its format, or neither for a list with no format check."
      : step === "commander"
        ? "The list marks no commander. Pick the card that leads it."
        : step === "done"
          ? "Each line below matched no card, so the deck holds the rest."
          : "A text file from Archidekt, or a pasted Arena list. The deck opens like a deck the agent built.";

  return (
    <DialogContent className="max-w-xl" aria-describedby="import-dialog-note">
      <DialogHeader>
        <DialogTitle asChild>
          <h2 className="font-display">{title}</h2>
        </DialogTitle>
        <DialogDescription id="import-dialog-note">{note}</DialogDescription>
      </DialogHeader>

      {step === "pick" && (
        <form onSubmit={onSubmit} className="flex flex-col gap-3">
          <label
            htmlFor="deck-file"
            className="flex cursor-pointer flex-col items-center gap-1.5 rounded-card border-2 border-dashed border-border px-6 py-6 text-center transition-colors hover:border-accent hover:bg-muted"
          >
            <FileTextIcon className="size-6 text-primary" aria-hidden="true" />
            <span className="font-display text-[15px]">{fileName || "Choose an Archidekt text file"}</span>
            <span className="font-mono text-[11px] text-muted-foreground">.txt — or paste the list below</span>
          </label>
          <Input id="deck-file" type="file" accept=".txt,text/plain" className="sr-only" onChange={(e) => void readFile(e.target.files?.[0])} />
          <div className="flex flex-col gap-1.5">
            <Label htmlFor="deck-text">Deck list</Label>
            <textarea
              id="deck-text"
              rows={8}
              value={text}
              onChange={(e) => setText(e.target.value)}
              placeholder={"Commander\n1 Heroes in a Half Shell (TMC) 6\n\nDeck\n1 Acidic Slime (TMC) 48"}
              className="min-h-32 rounded-card border border-border bg-background p-2 font-mono text-xs"
            />
          </div>
          <div className="flex flex-col gap-1.5">
            <Label htmlFor="deck-name">Deck name (optional, defaults to the file name)</Label>
            <Input id="deck-name" value={name} onChange={(e) => setName(e.target.value)} />
          </div>
          {collections.length > 0 && (
            <div className="flex flex-col gap-1.5">
              <Label htmlFor="deck-collection">Show owned cards from this collection</Label>
              <select
                id="deck-collection"
                value={collectionId}
                onChange={(e) => setCollectionId(e.target.value)}
                className="rounded-card border border-border bg-secondary px-2.5 py-1.5 text-sm"
              >
                <option value="">None</option>
                {collections.map((c) => (
                  <option key={c.id} value={c.id}>
                    {c.name} ({c.cardCount} cards)
                  </option>
                ))}
              </select>
            </div>
          )}
          <div className="min-h-6 text-sm">
            {tooLarge && (
              <p role="alert" className="text-danger">
                The list is over {maxImportBytes >> 10} KiB. A deck list is much shorter.
              </p>
            )}
            {send.isError && (
              <p role="alert" className="text-danger">
                Import failed: {errorMessage(send.error)}
              </p>
            )}
            {send.isPending && <p role="status">Reading the list and its bracket...</p>}
          </div>
          <DialogFooter>
            <Button type="button" variant="outline" onClick={onClose} disabled={send.isPending}>
              Cancel
            </Button>
            <Button type="submit" disabled={text.trim() === "" || tooLarge || send.isPending}>
              {send.isPending ? "Working..." : "Import"}
            </Button>
          </DialogFooter>
        </form>
      )}

      {step === "format" && (
        <Choice
          name="deck-format"
          options={[
            { value: String(FormatId.STANDARD), label: "Standard" },
            { value: String(FormatId.MODERN), label: "Modern" },
            { value: String(FormatId.HOUSE), label: "Neither: no format check" },
          ]}
          value={format === FormatId.UNSPECIFIED ? "" : String(format)}
          onChange={(v) => setFormat(Number(v) as FormatId)}
          busy={send.isPending}
          error={send.isError ? errorMessage(send.error) : ""}
          onCancel={onClose}
          onSubmit={() => send.mutate({ format })}
        />
      )}

      {step === "commander" && (
        <Choice
          name="deck-commander"
          options={options.map((o) => ({ value: o.oracleId, label: o.name }))}
          value={picked}
          onChange={setPicked}
          busy={send.isPending}
          error={send.isError ? errorMessage(send.error) : ""}
          onCancel={onClose}
          onSubmit={() => send.mutate({ commander: picked })}
        />
      )}

      {step === "done" && result && (
        <div className="flex flex-col gap-4">
          <SkippedLines rows={result.unresolved} />
          <DialogFooter>
            <Button
              type="button"
              onClick={() => {
                onClose();
                void navigate(`/decks/${result.deck?.id ?? ""}`);
              }}
            >
              Open the deck
            </Button>
          </DialogFooter>
        </div>
      )}
    </DialogContent>
  );
}

// Choice is one question with one answer: the format, or the commander.
function Choice({
  name,
  options,
  value,
  onChange,
  busy,
  error,
  onCancel,
  onSubmit,
}: {
  name: string;
  options: { value: string; label: string }[];
  value: string;
  onChange: (v: string) => void;
  busy: boolean;
  error: string;
  onCancel: () => void;
  onSubmit: () => void;
}) {
  return (
    <form
      className="flex flex-col gap-3"
      onSubmit={(e) => {
        e.preventDefault();
        if (value !== "" && !busy) onSubmit();
      }}
    >
      <fieldset className="flex max-h-72 flex-col gap-1.5 overflow-y-auto">
        {options.map((o) => (
          <label key={o.value} className={cn("flex items-center gap-2 text-sm")}>
            <input type="radio" name={name} value={o.value} checked={value === o.value} onChange={() => onChange(o.value)} className="accent-primary" />
            {o.label}
          </label>
        ))}
      </fieldset>
      <div className="min-h-6 text-sm">
        {error !== "" && (
          <p role="alert" className="text-danger">
            Import failed: {error}
          </p>
        )}
        {busy && <p role="status">Reading the list and its bracket...</p>}
      </div>
      <DialogFooter>
        <Button type="button" variant="outline" onClick={onCancel} disabled={busy}>
          Cancel
        </Button>
        <Button type="submit" disabled={value === "" || busy}>
          {busy ? "Working..." : "Import"}
        </Button>
      </DialogFooter>
    </form>
  );
}

function SkippedLines({ rows }: { rows: UnresolvedRow[] }) {
  return (
    <ul className="flex max-h-60 flex-col gap-1 overflow-y-auto font-mono text-xs">
      {rows.map((r) => (
        <li key={r.line}>
          Line {r.line}: {r.raw}
        </li>
      ))}
    </ul>
  );
}
