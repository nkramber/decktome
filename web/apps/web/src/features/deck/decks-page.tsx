import { useQuery } from "@tanstack/react-query";
import { useState } from "react";
import { Link } from "react-router";

import { deckClient } from "../../lib/api";
import { errorMessage } from "../../lib/errors";
import { DeckView } from "./deck-view";

// The deck list. One deck opens in place with the full deck view, so the
// M-5 scoring can read every deck without its session.
export function DecksPage() {
  const [openId, setOpenId] = useState("");
  const decks = useQuery({
    queryKey: ["decks"],
    queryFn: () => deckClient.listDecks({}),
  });

  return (
    <div className="mx-auto max-w-7xl p-6">
      <h1 className="text-2xl font-semibold">Your decks</h1>
      <div className="mt-2">
        {decks.isPending && <p role="status">Loading decks...</p>}
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
              {" · "}
              <button
                type="button"
                aria-expanded={openId === d.id}
                aria-controls={`deck-panel-${d.id}`}
                onClick={() => setOpenId(openId === d.id ? "" : d.id)}
                className="rounded border border-neutral-400 px-2 py-0.5 text-sm"
              >
                {openId === d.id ? "Hide" : "View"}
              </button>
              {openId === d.id && (
                <div id={`deck-panel-${d.id}`} className="mt-2 rounded border border-neutral-200 p-4">
                  <DeckView deck={d} />
                </div>
              )}
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}
