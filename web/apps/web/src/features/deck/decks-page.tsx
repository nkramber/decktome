import { useQuery } from "@tanstack/react-query";
import { Link } from "react-router";

import { deckClient } from "../../lib/api";
import { errorMessage } from "../../lib/errors";

// Placeholder listing. The deck view lands with PR-12.
export function DecksPage() {
  const decks = useQuery({
    queryKey: ["decks"],
    queryFn: () => deckClient.listDecks({}),
  });

  return (
    <div className="mx-auto max-w-3xl p-6">
      <h1 className="text-2xl font-semibold">Your decks</h1>
      <div aria-live="polite" className="mt-2">
        {decks.isPending && <p>Loading decks...</p>}
        {decks.isError && (
          <p role="alert" className="text-red-700">
            Could not list decks: {errorMessage(decks.error)}
          </p>
        )}
      </div>
      {decks.isSuccess && decks.data.decks.length === 0 && <p className="mt-2">No decks yet.</p>}
      {decks.isSuccess && decks.data.decks.length > 0 && (
        <ul className="mt-2 flex flex-col gap-1">
          {decks.data.decks.map((d) => (
            <li key={d.id}>
              <span className="font-medium">{d.name || d.id}</span>
              {" · "}
              {d.cards.length} entries, legality as of {d.legalityAsOf || "unknown"}
              {d.sessionId && (
                <>
                  {" · "}
                  <Link to={`/session/${d.sessionId}`} className="underline">
                    session
                  </Link>
                </>
              )}
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}
