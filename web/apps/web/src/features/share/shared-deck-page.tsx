import { Code, ConnectError } from "@connectrpc/connect";
import type { Card } from "@mtg/api-client/mtg/v1/card_pb";
import { CardRole, type DeckCard, type SharedCard, type SharedDeck } from "@mtg/api-client/mtg/v1/deck_pb";
import { useQuery } from "@tanstack/react-query";
import { useState } from "react";
import { useParams } from "react-router";

import { Button } from "../../components/ui/button";
import { deckClient } from "../../lib/api";
import { errorMessage } from "../../lib/errors";
import { CardGroup } from "../deck/deck-view";
import { formatLabel, groupByRole, roleLabel } from "../deck/deck-stats";
import { powerLabel } from "../deck/deck-stats";
import { downloadText } from "../export/buy-list";
import { identityOfCards, identityOfCommanders } from "../deck/color-identity";
import { ManaPips } from "../deck/mana-pips";

// The public deck page of a share link (D-315): the name, the format,
// the power, the summary, the cards by role with art, and the export.
// It shows no owner, no collection, no session, and no owned printing.
// The card data rides in the answer, so the page makes one call and
// needs no sign-in. A link that does not open says so and nothing more.

// entriesOf turns the shared cards into deck entries for the tiles, with
// no owned mark and no price of the owner's copies.
export function entriesOf(cards: SharedCard[]): DeckCard[] {
  return cards.map(
    (c) => ({ oracleId: c.oracleId, name: c.name, count: c.count, role: c.role, reason: c.reason, owned: false, ownedCount: 0, priceUsd: 0, outsideRequestedSets: false }) as DeckCard,
  );
}

export function cardsOf(deck: SharedDeck): Map<string, Card> {
  const byId = new Map<string, Card>();
  for (const c of [...deck.cards, ...deck.sideboard]) {
    if (c.card) byId.set(c.oracleId, c.card);
  }
  return byId;
}

export function SharedDeckPage() {
  const { token = "" } = useParams();
  const query = useQuery({
    queryKey: ["shared", token],
    queryFn: () => deckClient.getSharedDeck({ token }),
    enabled: token !== "",
    staleTime: Infinity,
    retry: false,
  });
  if (query.isPending) {
    return (
      <div className="mx-auto max-w-6xl px-4 py-8 md:px-6">
        <p role="status">Loading the deck...</p>
      </div>
    );
  }
  if (query.isError) {
    const gone = query.error instanceof ConnectError && query.error.code === Code.NotFound;
    return (
      <div className="mx-auto max-w-6xl px-4 py-8 md:px-6">
        <p role="alert" className="text-danger">
          {gone ? "This link does not open a deck. It was revoked, or it never existed." : `Could not load the deck: ${errorMessage(query.error)}`}
        </p>
      </div>
    );
  }
  const deck = query.data.deck;
  if (!deck) {
    return (
      <div className="mx-auto max-w-6xl px-4 py-8 md:px-6">
        <p role="alert" className="text-danger">
          This link does not open a deck.
        </p>
      </div>
    );
  }
  return (
    <div className="mx-auto max-w-6xl px-4 py-6 md:px-6">
      <SharedDeckView deck={deck} token={token} />
    </div>
  );
}

export function SharedDeckView({ deck, token }: { deck: SharedDeck; token: string }) {
  const byId = cardsOf(deck);
  const commanders = new Set(deck.commanderOracleIds);
  const main = entriesOf(deck.cards.filter((c) => !commanders.has(c.oracleId)));
  const commanderEntries = entriesOf(deck.cards.filter((c) => commanders.has(c.oracleId)));
  const side = entriesOf(deck.sideboard);
  const groups = groupByRole(main);
  const fromCommanders = identityOfCommanders(deck.commanderOracleIds, byId);
  const identity = fromCommanders.length > 0 ? fromCommanders : identityOfCards(byId);
  const total = deck.cardCount > 0 ? deck.cardCount : main.reduce((n, c) => n + c.count, 0) + commanderEntries.length;
  return (
    <article aria-labelledby="shared-deck-title" className="flex flex-col gap-5">
      <header className="flex flex-col gap-1.5 rounded-card border border-border bg-card p-6 shadow-card">
        <div className="flex flex-wrap items-center gap-3">
          <h2 id="shared-deck-title" className="wrap-anywhere text-2xl font-semibold tracking-tight text-balance">
            {deck.name || "Untitled deck"}
          </h2>
          <ManaPips colors={identity} />
        </div>
        <p className="text-sm text-muted-foreground">
          {formatLabel(deck.format?.id, deck.format?.houseRules ?? "")}
          {powerLabel(deck.power) && ` · ${powerLabel(deck.power)}`}
          {` · ${total} cards`}
          {side.length > 0 && ` · ${side.reduce((n, c) => n + c.count, 0)} sideboard`}
        </p>
        {deck.legalityAsOf && <p className="text-sm">Built against the card data of {deck.legalityAsOf}.</p>}
        {deck.summary && <p className="mt-2 max-w-measure leading-relaxed">{deck.summary}</p>}
        <p className="text-xs text-muted-foreground">A shared deck, read-only.</p>
      </header>

      <div className="print:hidden">
        <SharedExport token={token} />
      </div>

      <p className="text-xs text-muted-foreground">
        Card images and card text are unofficial Fan Content permitted under the Wizards of the Coast Fan Content Policy. They are
        copyright Wizards of the Coast, LLC, and come from Scryfall.
      </p>

      {commanderEntries.length > 0 && <CardGroup title="Commander" count={commanderEntries.length} entries={commanderEntries} byId={byId} commanders={commanders} hideOwnership />}
      {groups.map((g) => (
        <CardGroup key={g.role} title={roleLabel(g.role)} count={g.count} entries={g.cards} byId={byId} commanders={commanders} hideOwnership />
      ))}
      {groups.length === 0 && main.length > 0 && (
        <CardGroup title={roleLabel(CardRole.UNSPECIFIED)} count={main.length} entries={main} byId={byId} commanders={commanders} hideOwnership />
      )}
      {side.length > 0 && <CardGroup title="Sideboard" count={side.reduce((n, c) => n + c.count, 0)} entries={side} byId={byId} commanders={commanders} hideOwnership />}
    </article>
  );
}

// SharedExport copies or saves the deck list of a shared deck, through
// the public export call.
function SharedExport({ token }: { token: string }) {
  const [status, setStatus] = useState("");
  const [error, setError] = useState("");
  const [busy, setBusy] = useState(false);
  async function run(action: "copy" | "download") {
    setBusy(true);
    setError("");
    setStatus("");
    try {
      const res = await deckClient.exportSharedDeck({ token });
      const lines = res.text.split("\n").filter((l) => l.trim() !== "").length;
      if (action === "copy") {
        await navigator.clipboard.writeText(res.text);
        setStatus(`Copied the deck list: ${lines} lines.`);
      } else {
        downloadText(res.fileName, res.text);
        setStatus(`Saved ${res.fileName}: ${lines} lines.`);
      }
    } catch (e) {
      setError(`Export failed: ${errorMessage(e)}`);
    } finally {
      setBusy(false);
    }
  }
  return (
    <section aria-labelledby="shared-export-title" className="flex flex-col gap-3 rounded-card border border-border bg-card p-4 text-sm shadow-card">
      <h3 id="shared-export-title" className="font-medium">
        Export
      </h3>
      <div className="flex flex-wrap gap-2">
        <Button type="button" variant="outline" size="sm" disabled={busy} onClick={() => run("copy")}>
          Copy deck list
        </Button>
        <Button type="button" variant="outline" size="sm" disabled={busy} onClick={() => run("download")}>
          Download deck list
        </Button>
      </div>
      <p className="text-xs text-muted-foreground">The deck list holds every card, the commander first. ManaBox and MTG Arena both read it.</p>
      {status && <p role="status">{status}</p>}
      {error && (
        <p role="alert" className="text-danger">
          {error}
        </p>
      )}
    </section>
  );
}
