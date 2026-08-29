import { ImportSource } from "@mtg/api-client/mtg/v1/collection_pb";
import type { ImportCollectionResponse } from "@mtg/api-client/mtg/v1/collection_service_pb";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { type FormEvent, useState } from "react";
import { useNavigate } from "react-router";

import { collectionClient } from "../../lib/api";
import { errorMessage } from "../../lib/errors";
import { maxUploadBytes } from "../../lib/limits";
import { useAppStore } from "../../lib/store";
import { ImportResult } from "./import-result";

// The collection screen (ui plan, step 2). Upload a ManaBox CSV, or skip and
// build from any card (D-37). Earlier uploads come from ListCollections.
export function CollectionPage() {
  const navigate = useNavigate();
  const queryClient = useQueryClient();
  const collectionId = useAppStore((s) => s.collectionId);
  const setCollection = useAppStore((s) => s.setCollection);
  const clearCollection = useAppStore((s) => s.clearCollection);

  const [file, setFile] = useState<File | null>(null);
  // fileKey remounts the file input, which is the one way to empty it.
  const [fileKey, setFileKey] = useState(0);
  const [name, setName] = useState("");
  const [result, setResult] = useState<ImportCollectionResponse | null>(null);

  const list = useQuery({
    queryKey: ["collections"],
    queryFn: () => collectionClient.listCollections({}),
  });

  const upload = useMutation({
    mutationFn: async (f: File) => {
      const content = new Uint8Array(await f.arrayBuffer());
      return collectionClient.importCollection({
        name: name.trim() || f.name,
        source: ImportSource.MANABOX_CSV,
        content,
      });
    },
    onSuccess: (res) => {
      setResult(res);
      // The form empties, so a second click can not import the file again.
      clearFile();
      setName("");
      if (res.collection) {
        setCollection(res.collection.id);
      }
      void queryClient.invalidateQueries({ queryKey: ["collections"] });
    },
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
    if (file && !fileTooLarge) upload.mutate(file);
  }

  function skip() {
    clearCollection();
    clearFile();
    setResult(null);
    navigate("/session/new");
  }

  const collections = list.data?.collections ?? [];
  const active = collections.find((c) => c.id === collectionId) ?? result?.collection;

  return (
    <div className="mx-auto flex max-w-3xl flex-col gap-6 p-6">
      <h1 className="text-2xl font-semibold">Your collection</h1>

      <form onSubmit={onSubmit} className="flex flex-col gap-3 rounded border border-neutral-200 p-4">
        <h2 className="text-lg font-medium">Upload a ManaBox export</h2>
        <label className="flex flex-col gap-1">
          <span>ManaBox CSV file</span>
          <input
            key={fileKey}
            type="file"
            name="file"
            accept=".csv,text/csv"
            onChange={(e) => setFile(e.target.files?.[0] ?? null)}
            className="file:mr-3 file:rounded file:border file:border-neutral-400 file:bg-neutral-100 file:px-3 file:py-1"
          />
        </label>
        <label className="flex flex-col gap-1">
          <span>Collection name (optional, defaults to the file name)</span>
          <input
            type="text"
            name="name"
            value={name}
            onChange={(e) => setName(e.target.value)}
            className="rounded border border-neutral-400 bg-white px-2 py-1 text-neutral-900"
          />
        </label>
        <div className="flex gap-3">
          <button
            type="submit"
            disabled={!file || fileTooLarge || upload.isPending}
            className="rounded bg-neutral-900 px-3 py-2 text-white disabled:bg-neutral-300 disabled:text-neutral-600"
          >
            Upload
          </button>
          <button type="button" onClick={skip} className="rounded border border-neutral-400 px-3 py-2">
            Skip, build from any card
          </button>
        </div>
        <div className="min-h-6">
          {fileTooLarge && (
            <p role="alert" className="text-red-700">
              The file is {(file.size / (1 << 20)).toFixed(1)} MiB. The limit is {maxUploadBytes >> 20} MiB.
            </p>
          )}
          {upload.isPending && <p role="status">Uploading and resolving cards...</p>}
          {upload.isError && (
            <p role="alert" className="text-red-700">
              Upload failed: {errorMessage(upload.error)}
            </p>
          )}
        </div>
      </form>

      {result && <ImportResult result={result} />}

      <section className="flex flex-col gap-2">
        <h2 className="text-lg font-medium">Earlier uploads</h2>
        <div>
          {list.isPending && <p role="status">Loading collections...</p>}
          {list.isError && (
            <p role="alert" className="text-red-700">
              Could not list collections: {errorMessage(list.error)}
            </p>
          )}
        </div>
        {list.isSuccess && collections.length === 0 && <p>No uploads yet.</p>}
        {collections.length > 0 && (
          <ul className="flex flex-col gap-1">
            {collections.map((c) => {
              const isActive = c.id === collectionId;
              return (
                <li key={c.id} className="flex items-center gap-3">
                  <button
                    type="button"
                    aria-pressed={isActive}
                    onClick={() => {
                      // A second click on the active one clears it (D-37).
                      setResult(null);
                      clearFile();
                      if (isActive) clearCollection();
                      else setCollection(c.id);
                    }}
                    className={`rounded border px-2 py-1 ${isActive ? "border-neutral-900 bg-neutral-100 font-semibold ring-2 ring-neutral-900" : "border-neutral-400"}`}
                  >
                    {isActive && <span aria-hidden="true">✓ </span>}
                    {c.name}
                  </button>
                  <span className="text-sm text-neutral-600">
                    {c.cardCount} cards
                    {c.importedAt?.seconds ? `, imported ${new Date(Number(c.importedAt.seconds) * 1000).toLocaleDateString()}` : null}
                  </span>
                </li>
              );
            })}
          </ul>
        )}
      </section>

      <section className="flex items-center gap-4 rounded border border-neutral-200 p-4">
        <p data-testid="active-collection" className="grow">
          {collectionId && active
            ? `Active collection: ${active.name} (${active.cardCount} cards). The agent uses only these cards.`
            : collectionId
              ? `Active collection: ${collectionId}.`
              : "No active collection. The agent builds from any card (D-37)."}
        </p>
        <button
          type="button"
          onClick={() => navigate("/session/new")}
          className="rounded bg-neutral-900 px-3 py-2 text-white"
        >
          Continue to chat
        </button>
      </section>
    </div>
  );
}
