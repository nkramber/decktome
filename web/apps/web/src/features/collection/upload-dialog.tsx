import { type CollectionDiff, ImportSource } from "@mtg/api-client/mtg/v1/collection_pb";
import type { ImportCollectionResponse } from "@mtg/api-client/mtg/v1/collection_service_pb";
import { useMutation } from "@tanstack/react-query";
import { PackageIcon } from "lucide-react";
import { type FormEvent, useRef, useState } from "react";

import { Button } from "../../components/ui/button";
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle } from "../../components/ui/dialog";
import { Input } from "../../components/ui/input";
import { Label } from "../../components/ui/label";
import { collectionClient } from "../../lib/api";
import { cn } from "../../lib/cn";
import { errorMessage } from "../../lib/errors";
import { maxUploadBytes } from "../../lib/limits";
import { CollectionDiffBody } from "./collection-diff";
import { ImportReportBody } from "./import-result";

// The upload is a dialog: it holds the file, the progress, and the
// import report, so the page behind it never changes shape while a
// reader uploads (roadmap PR-18).
//
// A re-upload over an active collection shows what changes before
// anything replaces anything (D-393). A first upload has nothing to
// compare against, and it imports.

// step names where the dialog stands. The reader reads one thing at a
// time, and the buttons follow the step.
type Step = "pick" | "diff" | "done";

type UploadProps = {
  // activeCollectionId is the collection a re-upload compares against
  // and replaces. Empty means the reader has none, and the upload goes
  // straight in.
  activeCollectionId: string;
  activeCollectionName: string;
  // askForFile opens the file picker as the dialog opens. "Add a
  // collection" sends the reader here to pick a file (D-342).
  askForFile: boolean;
  onImported: (result: ImportCollectionResponse) => void;
};

// UploadDialog is the shell. Every state of the upload lives in the body
// below, and a new key on each open remounts the body, so a second open
// starts clean. This shell stays mounted while the dialog fades out, so
// the body can not reset itself on close without a flash of step "pick"
// over the fade. React supports this write during a render, and an
// effect that reset the body would cascade a render for no gain.
export function UploadDialog({
  open,
  onOpenChange,
  ...props
}: UploadProps & { open: boolean; onOpenChange: (open: boolean) => void }) {
  const [wasOpen, setWasOpen] = useState(open);
  const [openCount, setOpenCount] = useState(0);
  if (open !== wasOpen) {
    setWasOpen(open);
    if (open) setOpenCount((n) => n + 1);
  }
  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <UploadBody key={openCount} {...props} onClose={() => onOpenChange(false)} />
    </Dialog>
  );
}

function UploadBody({ activeCollectionId, activeCollectionName, askForFile, onImported, onClose }: UploadProps & { onClose: () => void }) {
  const [file, setFile] = useState<File | null>(null);
  const [name, setName] = useState("");
  const [dragOver, setDragOver] = useState(false);
  const [step, setStep] = useState<Step>("pick");
  const [diffResult, setDiffResult] = useState<CollectionDiff | null>(null);
  const [result, setResult] = useState<ImportCollectionResponse | null>(null);
  // The file picker opens once, when the input appears. See the ref of
  // the input below.
  const pickerOpened = useRef(false);

  // The server refuses an upload over maxUploadBytes, so the dialog says
  // so before the bytes go out.
  const fileTooLarge = file !== null && file.size > maxUploadBytes;

  const diff = useMutation({
    mutationFn: async (f: File) => {
      const content = new Uint8Array(await f.arrayBuffer());
      return collectionClient.diffCollections({ collectionId: activeCollectionId, source: ImportSource.MANABOX_CSV, content });
    },
    onSuccess: (res) => {
      setDiffResult(res.diff ?? null);
      setStep("diff");
    },
  });

  const upload = useMutation({
    mutationFn: async (f: File) => {
      const content = new Uint8Array(await f.arrayBuffer());
      return collectionClient.importCollection({
        name: name.trim() || f.name,
        source: ImportSource.MANABOX_CSV,
        content,
        // A replacement keeps the collection id, so every deck and chat
        // that names it still works (D-393).
        replaceCollectionId: step === "diff" ? activeCollectionId : "",
      });
    },
    onSuccess: (res) => {
      setResult(res);
      setStep("done");
      onImported(res);
    },
  });

  const busy = diff.isPending || upload.isPending;

  function onSubmit(e: FormEvent) {
    e.preventDefault();
    if (!file || fileTooLarge || busy) return;
    // A re-upload over an active collection reads first (D-393).
    if (activeCollectionId !== "") {
      diff.mutate(file);
      return;
    }
    upload.mutate(file);
  }

  const collectionName = activeCollectionName || "this collection";
  const identical = step === "diff" && diffResult?.identical === true;
  // The head of the dialog names the step, so the reader always reads
  // where the upload stands.
  const title = identical ? "Nothing changed" : step === "done" ? "Import result" : step === "diff" ? "What changes" : "Upload a ManaBox export";
  const note = identical
    ? `This file holds the same cards as ${collectionName}, in the same numbers.`
    : step === "done"
      ? "Every row the import could not read is listed below."
      : step === "diff"
        ? `Read what a replacement of ${collectionName} changes.`
        : "A ManaBox CSV export of your collection. Nothing is stored until you upload.";

  return (
    <DialogContent className="max-w-xl" aria-describedby="upload-dialog-note">
      <DialogHeader>
        <DialogTitle asChild>
          <h2 className="font-display">{title}</h2>
        </DialogTitle>
        <DialogDescription id="upload-dialog-note">{note}</DialogDescription>
      </DialogHeader>

      {step === "pick" && (
        <form onSubmit={onSubmit} className="flex flex-col gap-3">
          <div className="flex flex-col gap-1.5">
            <Label htmlFor="file" className="sr-only">
              ManaBox CSV file
            </Label>
            {/* The drop zone is the label of the file input, so a click
                and a drop both reach the one control. */}
            <label
              htmlFor="file"
              onDragOver={(e) => {
                e.preventDefault();
                setDragOver(true);
              }}
              onDragLeave={() => setDragOver(false)}
              onDrop={(e) => {
                e.preventDefault();
                setDragOver(false);
                setFile(e.dataTransfer.files?.[0] ?? null);
              }}
              className={cn(
                "flex cursor-pointer flex-col items-center gap-1.5 rounded-card border-2 border-dashed px-6 py-10 text-center transition-colors",
                dragOver ? "border-primary bg-secondary" : "border-border hover:border-accent hover:bg-muted",
              )}
            >
              <PackageIcon className="size-7 text-primary" aria-hidden="true" />
              <span className="font-display text-[15px]">{file ? file.name : "Drop your ManaBox export here"}</span>
              <span className="font-mono text-[11px] text-muted-foreground">.csv — or click to browse</span>
            </label>
            <Input
              // "Add a collection" sends the reader here to pick a file,
              // so the file picker opens with the dialog (D-342). The
              // dialog portal mounts this input a render after the body,
              // so a mount effect would find no input to click.
              ref={(el) => {
                if (!el || !askForFile || pickerOpened.current) return;
                pickerOpened.current = true;
                el.click();
              }}
              id="file"
              type="file"
              name="file"
              accept=".csv,text/csv"
              onChange={(e) => setFile(e.target.files?.[0] ?? null)}
              className="sr-only"
            />
          </div>
          <div className="flex flex-col gap-1.5">
            <Label htmlFor="name">Collection name (optional, defaults to the file name)</Label>
            <Input id="name" type="text" name="name" value={name} onChange={(e) => setName(e.target.value)} />
          </div>

          <UploadProgress busy={busy} label={diff.isPending ? "Reading what changed..." : "Uploading and resolving cards..."} />

          <div className="min-h-6 text-sm">
            {fileTooLarge && (
              <p role="alert" className="text-danger">
                The file is {(file.size / (1 << 20)).toFixed(1)} MiB. The limit is {maxUploadBytes >> 20} MiB.
              </p>
            )}
            {diff.isError && (
              <p role="alert" className="text-danger">
                Could not read that file: {errorMessage(diff.error)}
              </p>
            )}
            {upload.isError && (
              <p role="alert" className="text-danger">
                Upload failed: {errorMessage(upload.error)}
              </p>
            )}
          </div>

          <DialogFooter>
            <Button type="button" variant="outline" onClick={() => onClose()} disabled={busy}>
              Cancel
            </Button>
            <Button type="submit" disabled={!file || fileTooLarge || busy}>
              {busy ? "Working..." : "Upload"}
            </Button>
          </DialogFooter>
        </form>
      )}

      {step === "diff" && diffResult && file && (
        <div className="flex flex-col gap-4">
          <CollectionDiffBody diff={diffResult} />
          <UploadProgress busy={upload.isPending} label="Replacing the collection..." />
          {upload.isError && (
            <p role="alert" className="text-sm text-danger">
              Upload failed: {errorMessage(upload.error)}
            </p>
          )}
          <DialogFooter>
            <Button type="button" variant="outline" onClick={() => onClose()} disabled={upload.isPending}>
              Keep what I have
            </Button>
            {/* A file that changes nothing gets no replace button. */}
            {!diffResult.identical && (
              <Button type="button" onClick={() => upload.mutate(file)} disabled={upload.isPending}>
                {upload.isPending ? "Replacing..." : "Replace the collection"}
              </Button>
            )}
          </DialogFooter>
        </div>
      )}

      {step === "done" && result && (
        <div className="flex flex-col gap-4">
          <ImportReportBody result={result} />
          <DialogFooter>
            <Button type="button" onClick={() => onClose()}>
              Done
            </Button>
          </DialogFooter>
        </div>
      )}
    </DialogContent>
  );
}

// UploadProgress is an indeterminate bar. The transport reports no byte
// count, so a bar that filled to a number would be a guess. It says the
// work runs, and the label says which work.
function UploadProgress({ busy, label }: { busy: boolean; label: string }) {
  if (!busy) return null;
  return (
    <div className="flex flex-col gap-1.5" data-testid="upload-progress">
      <p role="status" className="text-sm">
        {label}
      </p>
      <div
        role="progressbar"
        aria-label={label}
        aria-busy="true"
        className="h-1 overflow-hidden rounded-full bg-muted"
      >
        <div className="h-full w-1/3 animate-pulse rounded-full bg-primary" />
      </div>
    </div>
  );
}
