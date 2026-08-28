import { ImportSource } from "@mtg/api-client/mtg/v1/collection_pb";
import type { ImportCollectionResponse } from "@mtg/api-client/mtg/v1/collection_service_pb";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { type FormEvent, useState } from "react";
import { useNavigate } from "react-router";

import { collectionClient } from "../../lib/api";
import { errorMessage } from "../../lib/errors";
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
      if (res.collection) {
        setCollection(res.collection.id);
      }
      void queryClient.invalidateQueries({ queryKey: ["collections"] });
    },
  });

  function onSubmit(e: FormEvent) {
    e.preventDefault();
    if (file) upload.mutate(file);
  }

  function skip() {
    clearCollection();
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
            type="file"
            name="file"
            accept=".csv,text/csv"
            onChange={(e) => setFile(e.target.files?.[0] ?? null)}
          />
        </label>
        <label className="flex flex-col gap-1">
          <span>Collection name (optional, defaults to the file name)</span>
          <input
            type="text"
            name="name"
            value={name}
            onChange={(e) => setName(e.target.value)}
            className="rounded border border-neutral-400 px-2 py-1"
          />
        </label>
        <div className="flex gap-3">
          <button
            type="submit"
            disabled={!file || upload.isPending}
            className="rounded bg-neutral-900 px-3 py-2 text-white disabled:opacity-50"
          >
            Upload
          </button>
          <button type="button" onClick={skip} className="rounded border border-neutral-400 px-3 py-2">
            Skip, build from any card
          </button>
        </div>
        <div aria-live="polite" className="min-h-6">
          {upload.isPending && <p>Uploading and resolving cards...</p>}
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
        <div aria-live="polite">
          {list.isPending && <p>Loading collections...</p>}
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
                    onClick={() => setCollection(c.id)}
                    className={`rounded border px-2 py-1 ${isActive ? "border-neutral-900 bg-neutral-100" : "border-neutral-400"}`}
                  >
                    {c.name}
                  </button>
                  <span className="text-sm text-neutral-600">
                    {c.cardCount} cards
                    {c.importedAt && `, imported ${new Date(Number(c.importedAt.seconds) * 1000).toLocaleDateString()}`}
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
