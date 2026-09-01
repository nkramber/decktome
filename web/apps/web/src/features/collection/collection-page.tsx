import { type CollectionDiff, ImportSource } from "@mtg/api-client/mtg/v1/collection_pb";
import type { ImportCollectionResponse } from "@mtg/api-client/mtg/v1/collection_service_pb";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { BookOpenIcon, PackageIcon, PencilIcon, Trash2Icon } from "lucide-react";
import { type FormEvent, useCallback, useEffect, useMemo, useRef, useState } from "react";
import { useLocation, useNavigate } from "react-router";

import { EmptyState } from "../../app/components/empty-state";
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "../../components/ui/alert-dialog";
import { ErrorState } from "../../app/components/error-state";
import { notify } from "../../app/components/notify";
import { PageHeader } from "../../app/components/page-header";
import { Button } from "../../components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "../../components/ui/card";
import { Input } from "../../components/ui/input";
import { Label } from "../../components/ui/label";
import { Skeleton } from "../../components/ui/skeleton";
import { collectionClient } from "../../lib/api";
import { cn } from "../../lib/cn";
import { errorMessage } from "../../lib/errors";
import { maxUploadBytes } from "../../lib/limits";
import { useAppStore } from "../../lib/store";
import { CollectionHero } from "./collection-hero";
import { ImportResult } from "./import-result";
import { BinderGrid } from "./binder-grid";
import { CollectionDiffView } from "./collection-diff";
import { useBinderPages, useCollectionHead } from "./use-collection";

// The collection screen (ui plan, step 2). Upload a ManaBox CSV, or skip and
// build from any card (D-37). Earlier uploads come from ListCollections.
export function CollectionPage() {
  const navigate = useNavigate();
  const location = useLocation();
  const queryClient = useQueryClient();
  const collectionId = useAppStore((s) => s.collectionId);
  const setCollection = useAppStore((s) => s.setCollection);
  const clearCollection = useAppStore((s) => s.clearCollection);

  const [file, setFile] = useState<File | null>(null);
  // fileKey remounts the file input, which is the one way to empty it.
  const [fileKey, setFileKey] = useState(0);
  const fileInput = useRef<HTMLInputElement>(null);
  const [name, setName] = useState("");
  const [dragOver, setDragOver] = useState(false);
  const [result, setResult] = useState<ImportCollectionResponse | null>(null);

  const list = useQuery({
    queryKey: ["collections"],
    queryFn: () => collectionClient.listCollections({}),
  });

  // "Add a collection" in the Build menu and in the pool picker sends the
  // reader here to pick a file, so the file dialog opens on arrival. The
  // history entry drops the mark first, so a reload opens no dialog.
  const askForFile = (location.state as { pickFile?: boolean } | null)?.pickFile === true;
  useEffect(() => {
    if (!askForFile) return;
    navigate(location.pathname, { replace: true, state: null });
    fileInput.current?.click();
  }, [askForFile, navigate, location.pathname]);

  // Deleting a collection is for good (D-347). A deck built from it keeps
  // every card. Its chat builds from the whole card database from then
  // on, and the chat says so on its next turn.
  const removeCollection = useMutation({
    mutationFn: (id: string) => collectionClient.deleteCollection({ collectionId: id }),
    onSuccess: (_res, id) => {
      if (id === collectionId) clearCollection();
      setResult(null);
      void notify("success", "Collection deleted", "A deck built from it keeps its cards.");
      void queryClient.invalidateQueries({ queryKey: ["collections"] });
      void queryClient.invalidateQueries({ queryKey: ["collection", id] });
    },
    onError: (err) => void notify("error", "Could not delete the collection", errorMessage(err)),
  });

  // The rename of the collections list. It writes the name alone, so it
  // rewrites no entry and it keeps the import date.
  const [renaming, setRenaming] = useState("");
  const [renameTo, setRenameTo] = useState("");
  const renameCollection = useMutation({
    mutationFn: (v: { id: string; name: string }) => collectionClient.updateCollection({ collectionId: v.id, name: v.name }),
    onSuccess: () => {
      setRenaming("");
      void notify("success", "Collection renamed", "");
      void queryClient.invalidateQueries({ queryKey: ["collections"] });
      void queryClient.invalidateQueries({ queryKey: ["collection"] });
    },
    onError: (err) => void notify("error", "Could not rename the collection", errorMessage(err)),
  });

  // The diff of a re-upload against the active collection (D-393). It
  // stores nothing: the reader reads it and then chooses.
  const [pendingFile, setPendingFile] = useState<File | null>(null);
  const [diffResult, setDiffResult] = useState<CollectionDiff | null>(null);
  const diff = useMutation({
    mutationFn: async (f: File) => {
      const content = new Uint8Array(await f.arrayBuffer());
      return collectionClient.diffCollections({ collectionId, source: ImportSource.MANABOX_CSV, content });
    },
    onSuccess: (res, f) => {
      setPendingFile(f);
      setDiffResult(res.diff ?? null);
    },
    onError: (err) => void notify("error", "Could not read that file", errorMessage(err)),
  });

  function cancelDiff() {
    setPendingFile(null);
    setDiffResult(null);
    clearFile();
  }

  const upload = useMutation({
    mutationFn: async (f: File) => {
      const content = new Uint8Array(await f.arrayBuffer());
      return collectionClient.importCollection({
        name: name.trim() || f.name,
        source: ImportSource.MANABOX_CSV,
        content,
        // A replacement keeps the collection id, so every deck and chat
        // that names it still works (D-393).
        replaceCollectionId: pendingFile ? collectionId : "",
      });
    },
    onSuccess: (res) => {
      setResult(res);
      // The form empties, so a second click can not import the file again.
      clearFile();
      setName("");
      setPendingFile(null);
      setDiffResult(null);
      if (res.collection) {
        setCollection(res.collection.id);
      }
      void notify("success", "Collection imported", `${res.collection?.cardCount ?? 0} cards are ready.`);
      void queryClient.invalidateQueries({ queryKey: ["collections"] });
      void queryClient.invalidateQueries({ queryKey: ["collection"] });
    },
    onError: (err) => void notify("error", "Upload failed", errorMessage(err)),
  });

  // The server refuses an upload over maxUploadBytes, so the page says so
  // before the bytes go out.
  const fileTooLarge = file !== null && file.size > maxUploadBytes;

  function clearFile() {
    setFile(null);
    setFileKey((k) => k + 1);
  }

  function onSubmit(e: FormEvent) {
    e.preventDefault();
    if (!file || fileTooLarge) return;
    // A re-upload over an active collection shows what changes first
    // (D-393). A first upload has nothing to compare against.
    if (collectionId !== "") {
      diff.mutate(file);
      return;
    }
    upload.mutate(file);
  }

  function skip() {
    clearCollection();
    clearFile();
    setResult(null);
    navigate("/session/new");
  }

  const collections = list.data?.collections ?? [];
  const active = collections.find((c) => c.id === collectionId) ?? result?.collection;
  // The binder of the active collection, for the head of the screen. It
  // loads once, and no other screen needs it (D-327).
  const binder = useCollectionHead(collectionId);
  // The rows of the binder, a page at a time. The grid asks for the next
  // page as the reader scrolls (D-392).
  const pages = useBinderPages(collectionId);
  const rows = useMemo(() => (pages.data?.pages ?? []).flatMap((p) => p.collection?.entries ?? []), [pages.data]);
  const loadMore = useCallback(() => {
    if (pages.hasNextPage && !pages.isFetchingNextPage) void pages.fetchNextPage();
  }, [pages]);

  return (
    <div className="mx-auto flex w-full max-w-5xl flex-col gap-6 p-4 md:p-6">
      <PageHeader title="Your collection" description="Upload a ManaBox export, or skip it and build from any card." />


      <div className="grid gap-6 lg:grid-cols-2 lg:items-start">
      <form onSubmit={onSubmit}>
        <Card>
          <CardHeader>
            <CardTitle>Upload a ManaBox export</CardTitle>
          </CardHeader>
          <CardContent className="flex flex-col gap-3">
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
              <Input key={fileKey} ref={fileInput} id="file" type="file" name="file" accept=".csv,text/csv" onChange={(e) => setFile(e.target.files?.[0] ?? null)} className="sr-only" />
            </div>
            <div className="flex flex-col gap-1.5">
              <Label htmlFor="name">Collection name (optional, defaults to the file name)</Label>
              <Input id="name" type="text" name="name" value={name} onChange={(e) => setName(e.target.value)} />
            </div>
            <div className="flex flex-wrap gap-2">
              <Button type="submit" disabled={!file || fileTooLarge || upload.isPending}>
                Upload
              </Button>
              <Button type="button" variant="outline" onClick={skip}>
                Skip, build from any card
              </Button>
            </div>
            <div className="min-h-6 text-sm">
              {fileTooLarge && (
                <p role="alert" className="text-danger">
                  The file is {(file.size / (1 << 20)).toFixed(1)} MiB. The limit is {maxUploadBytes >> 20} MiB.
                </p>
              )}
              {upload.isPending && <p role="status">Uploading and resolving cards...</p>}
              {upload.isError && (
                <p role="alert" className="text-danger">
                  Upload failed: {errorMessage(upload.error)}
                </p>
              )}
            </div>
          </CardContent>
        </Card>
      </form>

      <section className="flex flex-col gap-3">
        <h2 className="font-display text-lg font-semibold">Earlier uploads</h2>
        {list.isPending && (
          <div role="status" className="flex flex-col gap-2">
            <span className="sr-only">Loading collections...</span>
            <Skeleton className="h-9 w-full" />
            <Skeleton className="h-9 w-2/3" />
          </div>
        )}
        {list.isError && <ErrorState title="Could not list collections" message={errorMessage(list.error)} onRetry={() => void list.refetch()} />}
        {list.isSuccess && collections.length === 0 && <EmptyState icon={BookOpenIcon} title="No uploads yet." description="Upload a ManaBox export above, or skip and build from any card." />}
        {collections.length > 0 && (
          <ul className="flex flex-col gap-2">
            {collections.map((c) => {
              const isActive = c.id === collectionId;
              return (
                <li key={c.id} className="flex flex-wrap items-center gap-3">
                  <Button
                    variant="outline"
                    size="sm"
                    aria-pressed={isActive}
                    onClick={() => {
                      // A second click on the active one clears it (D-37).
                      setResult(null);
                      clearFile();
                      if (isActive) clearCollection();
                      else setCollection(c.id);
                    }}
                    className={cn(isActive && "border-accent font-semibold ring-2 ring-ring")}
                  >
                    {isActive && <span aria-hidden="true">✓ </span>}
                    {c.name}
                  </Button>
                  <span className="text-sm text-muted-foreground">
                    {c.cardCount} cards
                    {c.importedAt?.seconds ? `, imported ${new Date(Number(c.importedAt.seconds) * 1000).toLocaleDateString()}` : null}
                  </span>
                  <Button
                    variant="ghost"
                    size="icon"
                    aria-label={`Rename ${c.name}`}
                    className="ml-auto size-7"
                    onClick={() => {
                      setRenaming(c.id);
                      setRenameTo(c.name);
                    }}
                  >
                    <PencilIcon aria-hidden="true" />
                  </Button>
                  <AlertDialog>
                    <AlertDialogTrigger asChild>
                      <Button variant="ghost" size="icon" aria-label={`Delete ${c.name}`} className="size-7 text-danger hover:text-danger">
                        <Trash2Icon aria-hidden="true" />
                      </Button>
                    </AlertDialogTrigger>
                    <AlertDialogContent>
                      <AlertDialogHeader>
                        <AlertDialogTitle>Delete {c.name}?</AlertDialogTitle>
                        <AlertDialogDescription>
                          This can not be undone. Every deck you built from it keeps all of its cards. The chat of each one builds from the whole card database from
                          now on, and it says so.
                        </AlertDialogDescription>
                      </AlertDialogHeader>
                      <AlertDialogFooter>
                        <AlertDialogCancel>Keep it</AlertDialogCancel>
                        <AlertDialogAction onClick={() => removeCollection.mutate(c.id)}>Delete the collection</AlertDialogAction>
                      </AlertDialogFooter>
                    </AlertDialogContent>
                  </AlertDialog>
                  {renaming === c.id && (
                    <form
                      className="flex w-full flex-wrap items-center gap-2"
                      onSubmit={(e) => {
                        e.preventDefault();
                        const next = renameTo.trim();
                        if (next !== "") renameCollection.mutate({ id: c.id, name: next });
                      }}
                    >
                      <Label htmlFor={`rename-${c.id}`} className="sr-only">
                        New name for {c.name}
                      </Label>
                      {/* The reader opened this form, so the cursor
                          belongs in it. A ref focuses on mount, which
                          autoFocus can not do accessibly. */}
                      <Input
                        id={`rename-${c.id}`}
                        ref={(el) => el?.focus()}
                        value={renameTo}
                        onChange={(e) => setRenameTo(e.target.value)}
                        className="h-8 max-w-64"
                      />
                      <Button type="submit" size="sm" disabled={renameTo.trim() === "" || renameCollection.isPending}>
                        Save
                      </Button>
                      <Button type="button" size="sm" variant="outline" onClick={() => setRenaming("")}>
                        Cancel
                      </Button>
                    </form>
                  )}
                </li>
              );
            })}
          </ul>
        )}
      </section>
      </div>

      {diffResult && pendingFile && (
        <CollectionDiffView
          name={active?.name ?? "this collection"}
          diff={diffResult}
          pending={upload.isPending}
          onReplace={() => upload.mutate(pendingFile)}
          onCancel={cancelDiff}
        />
      )}

      {result && <ImportResult result={result} />}

      <Card>
        <CardContent className="flex flex-wrap items-center gap-4">
          <p data-testid="active-collection" className="grow text-sm">
            {collectionId && active
              ? `Active collection: ${active.name} (${active.cardCount} cards). The agent uses only these cards.`
              : collectionId && !list.isSuccess
                ? "Active collection: loading..."
                : collectionId
                  ? "That collection is gone. Pick another, or upload one."
                  : "No active collection. The agent builds from any card."}
          </p>
          <Button onClick={() => navigate("/session/new")}>Continue to chat</Button>
        </CardContent>
      </Card>

      {/* The binder sits under the controls, not over them. A click on an
          earlier upload shows or hides it, and nothing above it moves. */}
      {binder.data?.collection && <CollectionHero collection={binder.data.collection} loading={binder.isPending} />}
      {collectionId !== "" && binder.isPending && <Skeleton className="h-56 w-full rounded-card" />}
      {/* The rows load a page at a time, and the head above reads none
          of them (D-392). */}
      {collectionId !== "" && rows.length > 0 && (
        <BinderGrid entries={rows} loading={pages.isFetchingNextPage} onReachEnd={loadMore} />
      )}
    </div>
  );
}
