import { useQuery } from "@tanstack/react-query";
import { LayersIcon } from "lucide-react";
import { useState } from "react";
import { Link } from "react-router";

import { EmptyState } from "../../app/components/empty-state";
import { ErrorState } from "../../app/components/error-state";
import { PageHeader } from "../../app/components/page-header";
import { Button } from "../../components/ui/button";
import { Card, CardContent } from "../../components/ui/card";
import { Skeleton } from "../../components/ui/skeleton";
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
    <div className="mx-auto flex max-w-7xl flex-col gap-6 p-4 md:p-6">
      <PageHeader title="Your decks" description="Every deck the agent built for you." />
      {decks.isPending && (
        <div role="status" className="flex flex-col gap-2">
          <span className="sr-only">Loading decks...</span>
          <Skeleton className="h-16 w-full" />
          <Skeleton className="h-16 w-full" />
        </div>
      )}
      {decks.isError && <ErrorState title="Could not list decks" message={errorMessage(decks.error)} onRetry={() => void decks.refetch()} />}
      {decks.isSuccess && decks.data.decks.length === 0 && <EmptyState icon={LayersIcon} title="No decks yet." description="Start a chat and the agent builds your first deck." action={<Button asChild><Link to="/session/new">Build a deck</Link></Button>} />}
      {decks.isSuccess && decks.data.decks.length > 0 && (
        <ul className="flex flex-col gap-3">
          {decks.data.decks.map((d) => (
            <li key={d.id}>
              <Card>
                <CardContent className="flex flex-wrap items-center gap-x-3 gap-y-2">
                  <span className="font-medium">{d.name || d.id}</span>
                  <span className="text-sm text-muted-foreground">
                    {d.cards.length} entries, legality as of {d.legalityAsOf || "unknown"}
                  </span>
                  {d.sessionId && (
                    <Button asChild variant="link" size="sm" className="px-0">
                      <Link to={`/session/${d.sessionId}`}>session</Link>
                    </Button>
                  )}
                  <span className="grow" />
                  <Button variant="outline" size="sm" aria-expanded={openId === d.id} aria-controls={`deck-panel-${d.id}`} onClick={() => setOpenId(openId === d.id ? "" : d.id)}>
                    {openId === d.id ? "Hide" : "View"}
                  </Button>
                </CardContent>
                {openId === d.id && (
                  <CardContent id={`deck-panel-${d.id}`} className="border-t border-border pt-4">
                    <DeckView deck={d} />
                  </CardContent>
                )}
              </Card>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}
