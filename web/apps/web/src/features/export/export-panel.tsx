import type { Card } from "@mtg/api-client/mtg/v1/card_pb";
import { ExportFormat } from "@mtg/api-client/mtg/v1/deck_service_pb";
import type { Deck } from "@mtg/api-client/mtg/v1/deck_pb";
import { useState } from "react";

import { Button } from "../../components/ui/button";
import { deckClient } from "../../lib/api";
import { errorMessage } from "../../lib/errors";
import { type BuyRow, buyRows, downloadText } from "./buy-list";

type Action = "copy" | "download";

// The button says what the reader gets, not the name of the text shape.
// The deck list is the Arena text shape, which ManaBox and MTG Arena
// both read (D-15, D-307).
const formatLabel: Record<ExportFormat, string> = {
  [ExportFormat.UNSPECIFIED]: "deck list",
  [ExportFormat.ARENA_TEXT]: "deck list",
  [ExportFormat.BUY_LIST_TEXT]: "buy list",
};

// The export panel (ui plan, step 5): the whole deck list for ManaBox
// (D-15, D-307), the buy list as text for a shop (D-309), and the buy
// list on screen with a Scryfall link per card (D-308). The text comes
// from DeckService.ExportDeck, so the browser and the API agree on it.
export function ExportPanel({ deck, byId }: { deck: Deck; byId: Map<string, Card> }) {
  const [status, setStatus] = useState("");
  const [error, setError] = useState("");
  const [busy, setBusy] = useState(false);
  const { needed, upgrades } = buyRows(deck, byId);

  async function run(format: ExportFormat, action: Action) {
    setBusy(true);
    setError("");
    setStatus("");
    try {
      const res = await deckClient.exportDeck({ deckId: deck.id, format });
      const lines = res.text.split("\n").filter((l) => l.trim() !== "").length;
      if (action === "copy") {
        await navigator.clipboard.writeText(res.text);
        setStatus(`Copied the ${formatLabel[format]}: ${lines} lines.`);
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

  const canExport = deck.id !== "" && !busy;
  return (
    <section aria-labelledby={`export-title-${deck.id}`} className="flex flex-col gap-3 rounded-card border border-border bg-card p-4 text-sm shadow-card">
      <h3 id={`export-title-${deck.id}`} className="font-medium">
        Export
      </h3>
      <div className="flex flex-wrap gap-2">
        <Button type="button" variant="outline" size="sm" disabled={!canExport} onClick={() => run(ExportFormat.ARENA_TEXT, "copy")}>
          Copy deck list
        </Button>
        <Button type="button" variant="outline" size="sm" disabled={!canExport} onClick={() => run(ExportFormat.ARENA_TEXT, "download")}>
          Download deck list
        </Button>
        <Button type="button" variant="outline" size="sm" disabled={!canExport} onClick={() => run(ExportFormat.BUY_LIST_TEXT, "copy")}>
          Copy buy list
        </Button>
        <Button type="button" variant="outline" size="sm" disabled={!canExport} onClick={() => run(ExportFormat.BUY_LIST_TEXT, "download")}>
          Download buy list
        </Button>
      </div>
      <p className="text-xs text-muted-foreground">
        The deck list holds every card, the commander first. ManaBox and MTG Arena both read it. The buy list pastes into a shop&apos;s mass-entry form.
      </p>
      {status && <p role="status">{status}</p>}
      {error && (
        <p role="alert" className="text-danger">
          {error}
        </p>
      )}
      <BuyList title="Buy list" rows={needed} empty="Nothing to buy: every card is owned." />
      {upgrades.length > 0 && <BuyList title="Upgrades" rows={upgrades} empty="" />}
    </section>
  );
}

// The list of cards to buy is long, so it stays shut until it is asked
// for. A deck of a hundred cards would otherwise fill the panel.
function BuyList({ title, rows, empty }: { title: string; rows: BuyRow[]; empty: string }) {
  const count = rows.reduce((n, r) => n + r.count, 0);
  const heading = (
    <>
      {title} <span className="text-muted-foreground">({count})</span>
    </>
  );
  if (rows.length === 0) {
    return (
      <section aria-label={`${title} (${count})`}>
        <h4 className="font-display text-[13px] font-semibold">{heading}</h4>
        <p>{empty}</p>
      </section>
    );
  }
  return (
    <section aria-label={`${title} (${count})`}>
      <details>
        <summary className="font-display cursor-pointer text-[13px] font-semibold marker:text-primary">{heading}</summary>
        <ul className="mt-1 list-disc pl-5">
          {rows.map((r) => (
            <li key={r.oracleId}>
              {r.count} × {r.name}
              {r.priceUsd > 0 && ` · $${r.priceUsd.toFixed(2)}`}
              {r.scryfallUrl && (
                <>
                  {" · "}
                  <a href={r.scryfallUrl} target="_blank" rel="noreferrer" className="underline">
                    Scryfall
                  </a>
                </>
              )}
            </li>
          ))}
        </ul>
      </details>
    </section>
  );
}
