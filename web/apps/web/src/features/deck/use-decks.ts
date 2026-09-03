import type { Card } from "@mtg/api-client/mtg/v1/card_pb";
import type { Deck } from "@mtg/api-client/mtg/v1/deck_pb";
import type { FormatId, SixtyStep } from "@mtg/api-client/mtg/v1/format_pb";
import { useInfiniteQuery, useMutation, useQuery, useQueryClient } from "@tanstack/react-query";

import { deckClient } from "../../lib/api";
import { fetchCards } from "./use-cards";

// The library reads one page of decks and the cards of that page's
// commanders (PR-17). The commander card gives the grid its art and its
// color identity, and one GetCards call serves the whole page.

export type DeckFilter = {
  query: string;
  format: FormatId;
  favorite?: boolean;
  powerBracket: number;
  sixtyStep: SixtyStep;
};

export const emptyDeckFilter: DeckFilter = {
  query: "",
  format: 0 as FormatId,
  favorite: undefined,
  powerBracket: 0,
  sixtyStep: 0 as SixtyStep,
};

// The key holds every field of the filter, so a change refetches and a
// page token never crosses filters. The server refuses a token of another
// filter, and this key makes that impossible to reach.
function listKey(filter: DeckFilter) {
  return ["decks", filter.query, filter.format, filter.favorite ?? "any", filter.powerBracket, filter.sixtyStep] as const;
}

export function useDeckList(filter: DeckFilter) {
  return useInfiniteQuery({
    queryKey: listKey(filter),
    initialPageParam: "",
    queryFn: ({ pageParam }) =>
      deckClient.listDecks({
        query: filter.query,
        format: filter.format,
        favorite: filter.favorite,
        powerBracket: filter.powerBracket,
        powerSixtyStep: filter.sixtyStep,
        pageToken: pageParam,
      }),
    getNextPageParam: (last) => last.nextPageToken || undefined,
    placeholderData: (prev) => prev,
  });
}

// useDeckVersions reads every deck one chat built, oldest first. That is
// the version history of a deck (PR-17): a revision turn writes a new
// deck in the same session, and `revised_from_deck_id` names the one it
// revised. A deck with no session has one version, itself.
export function useDeckVersions(sessionId: string) {
  const query = useQuery({
    queryKey: ["decks", "session", sessionId],
    queryFn: () => deckClient.listDecks({ sessionId, pageSize: maxVersions }),
    enabled: sessionId !== "",
  });
  const decks = query.data?.decks ?? [];
  // The listing answers newest first, and a history reads oldest first.
  return { ...query, versions: [...decks].reverse() };
}

// maxVersions is the server's page cap. A chat with more revisions than
// this shows the newest of them.
const maxVersions = 100;

// useCommanderCards loads the commander of every deck on the page. A deck
// with no commander needs nothing, and a page with none makes no call.
export function useCommanderCards(decks: Deck[]) {
  const ids = [...new Set(decks.flatMap((d) => d.commanderOracleIds))].sort();
  const query = useQuery({
    queryKey: ["cards", "commanders", ids],
    queryFn: () => fetchCards(ids),
    enabled: ids.length > 0,
    staleTime: Infinity,
    placeholderData: (prev) => prev,
  });
  return query.data?.byId ?? new Map<string, Card>();
}

// The three writes of the library (PR-17). Each one refreshes every deck
// listing, so the grid shows the result with no reload.
export function useDeckWrites() {
  const client = useQueryClient();
  const refresh = () => {
    void client.invalidateQueries({ queryKey: ["decks"] });
    void client.invalidateQueries({ queryKey: ["deck"] });
    // A deck delete takes its chat (D-456), so the chat list reads again.
    void client.invalidateQueries({ queryKey: ["sessions"] });
  };

  const rename = useMutation({
    mutationFn: (v: { deckId: string; name: string }) => deckClient.updateDeck({ deckId: v.deckId, name: v.name }),
    onSuccess: refresh,
  });
  const setFavorite = useMutation({
    mutationFn: (v: { deckId: string; favorite: boolean }) => deckClient.updateDeck({ deckId: v.deckId, favorite: v.favorite }),
    onSuccess: refresh,
  });
  const remove = useMutation({
    mutationFn: (v: { deckId: string }) => deckClient.deleteDeck({ deckId: v.deckId }),
    onSuccess: refresh,
  });
  // The share link (D-315): a share answers the token once, and a revoke
  // ends the link. Both refresh the deck, which carries the shared mark.
  const share = useMutation({
    mutationFn: (v: { deckId: string }) => deckClient.shareDeck({ deckId: v.deckId }),
    onSuccess: refresh,
  });
  const revokeShare = useMutation({
    mutationFn: (v: { deckId: string }) => deckClient.revokeShare({ deckId: v.deckId }),
    onSuccess: refresh,
  });
  return { rename, setFavorite, remove, share, revokeShare };
}
