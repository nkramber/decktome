import { useParams } from "react-router";

import { useAppStore } from "../../lib/store";

// Placeholder. The chat thread lands with PR-12. The id "new" means no session yet.
export function SessionPage() {
  const { id } = useParams();
  const collectionId = useAppStore((s) => s.collectionId);
  const poolMode = useAppStore((s) => s.poolMode);
  const isNew = !id || id === "new";

  return (
    <div className="mx-auto max-w-3xl p-6">
      <h1 className="text-2xl font-semibold">Chat</h1>
      <p className="mt-2">The chat lands with PR-12.</p>
      <p className="mt-2" data-testid="session-id">
        {isNew ? "No session yet." : `Session id: ${id}`}
      </p>
      <p className="mt-2 text-sm text-neutral-600">
        {poolMode === "owned" && collectionId
          ? `Pool: only cards in collection ${collectionId}.`
          : "Pool: any card (D-37)."}
      </p>
    </div>
  );
}
