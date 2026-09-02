import { type Card, Color } from "@mtg/api-client/mtg/v1/card_pb";
import { CardRole, type Deck, type DeckCard, Severity } from "@mtg/api-client/mtg/v1/deck_pb";

import { errorMessage } from "../../lib/errors";
import { ExportPanel } from "../export/export-panel";
import { CardTile } from "./card-tile";
import { identityOfCards, identityOfCommanders } from "./color-identity";
import { ManaPips } from "./mana-pips";
import {
  colorLetters,
  colorSources,
  curveSteps,
  deckColors,
  diffDecks,
  formatLabel,
  deckSize,
  groupByRole,
  manaCurve,
  powerLabel,
  priceText,
  roleLabel,
  severityLabel,
} from "./deck-stats";
import { DiffList } from "./deck-diff";
import { useDeckCards } from "./use-cards";

// The deck view (ui plan, step 4). Cards group by role, each with its art
// and attribution (D-6), both faces for a DFC (F-9), the owned mark or
// the price (D-2, D-37), the findings, legality_as_of, the curve, and
// the color sources.
// base is the deck this one revised, when the page holds it (PR-12B).
// The five colors of the game and colorless carry fixed tokens (D-311).
const manaSwatch: Record<string, string> = { W: "bg-mana-w", U: "bg-mana-u", B: "bg-mana-b", R: "bg-mana-r", G: "bg-mana-g", C: "bg-mana-c" };

export function DeckView({ deck, base }: { deck: Deck; base?: Deck }) {
  const diff = base && deck.revisedFromDeckId && base.id === deck.revisedFromDeckId ? diffDecks(base, deck) : undefined;
  const cards = useDeckCards(deck);
  const { byId, missing } = cards;
  const allEntries = [...deck.cards, ...deck.sideboard, ...deck.upgrades];
  const nameOf = (id: string) => allEntries.find((c) => c.oracleId === id)?.name ?? byId.get(id)?.name ?? id;
  const commanders = new Set(deck.commanderOracleIds);
  const main = deck.cards.filter((c) => !commanders.has(c.oracleId));
  // The command zone: one entry per commander id, from cards when the
  // list holds it and from the card data otherwise (D-289).
  const commanderEntries: DeckCard[] = deck.commanderOracleIds.map(
    (id) =>
      deck.cards.find((c) => c.oracleId === id) ??
      ({ oracleId: id, name: byId.get(id)?.name ?? "Commander", count: 1, role: CardRole.UNSPECIFIED, owned: false, ownedCount: 0, priceUsd: 0, reason: "" } as DeckCard),
  );
  const commanderCount = deck.commanderOracleIds.length;
  const groups = groupByRole(main);
  const curve = manaCurve(deck.cards, byId);
  const sources = colorSources(deck.cards, byId);
  // The table shows the colors the deck pays for, and colorless for a
  // colorless deck. A mono-green deck full of rocks that make any color
  // is not a five-color deck. The commander's identity counts too.
  const colorsOfDeck = deckColors(deck.cards, byId, deck.commanderOracleIds);
  const sourceRows = colorLetters.filter((c) => (colorsOfDeck.size === 0 ? c.color === Color.C : colorsOfDeck.has(c.color)));
  // The count is the main deck plus the command zone, so a Commander
  // deck reads 100 (D-454).
  const total = deckSize(main.reduce((n, c) => n + c.count, 0), commanderCount);
  const validation = deck.validation;
  // A not_owned warning repeats what the tile says under the card, and
  // an owned-first deck carries one per card to buy. The list drops them.
  // A not_owned block in owned-only still shows (D-300).
  const findings = (validation?.findings ?? []).filter((f) => !(f.code === "not_owned" && f.severity !== Severity.BLOCK));
  const legalityAsOf = deck.legalityAsOf || validation?.legalityAsOf || "an unknown date";
  const curveMax = Math.max(1, ...curve);
  const sourcesMax = Math.max(1, ...sourceRows.map((c) => sources.get(c.color) ?? 0));
  // The deck owns the color of its own page (D-327).
  // The commanders name the identity. A deck with none, a 60-card deck,
  // reads its own cards, and this view holds only its own cards.
  const commanderCard = deck.commanderOracleIds.map((id) => byId.get(id)).find(Boolean);
  const commanderArt = commanderCard?.faces?.[0]?.imageUris?.artCrop ?? commanderCard?.defaultPrinting?.imageUris?.artCrop ?? "";
  const fromCommanders = identityOfCommanders(deck.commanderOracleIds, byId);
  const identity = fromCommanders.length > 0 ? fromCommanders : identityOfCards(byId);

  return (
    <article aria-labelledby={`deck-title-${deck.id}`} className="flex flex-col gap-5">
      <header className="relative isolate flex flex-col gap-1.5 overflow-hidden rounded-card border border-border bg-card p-6 shadow-card">
        {/* The commander's own artwork sits behind its deck. Scryfall
            serves it as art_crop, so the app crops nothing (D-6). */}
        {commanderArt && <img src={commanderArt} alt="" aria-hidden="true" className="pointer-events-none absolute inset-0 -z-20 size-full object-cover opacity-25" />}
        <span aria-hidden="true" className="pointer-events-none absolute inset-0 -z-10 bg-gradient-to-r from-card via-card/90 to-card/60" />
        <span aria-hidden="true" className="hatch pointer-events-none absolute inset-0 -z-10 opacity-[0.03]" />
        <div className="flex flex-wrap items-center gap-3">
          <h2 id={`deck-title-${deck.id}`} className="wrap-anywhere text-2xl font-semibold tracking-tight text-balance">
            {deck.name || "Untitled deck"}
          </h2>
          <ManaPips colors={identity} />
        </div>
        <p className="text-sm text-muted-foreground">
          {formatLabel(deck.format?.id, deck.format?.houseRules ?? "")}
          {powerLabel(deck.power) && ` · ${powerLabel(deck.power)}`}
          {` · ${total} cards`}
          {deck.sideboard.length > 0 && ` · ${deck.sideboard.reduce((n, c) => n + c.count, 0)} sideboard`}
        </p>
        <p className="text-sm" data-testid="legality-line">
          {validation
            ? `${validation.passed ? "Legal" : "Not legal"}, checked against the card data of ${legalityAsOf}.`
            : "Legality not checked yet."}
          {deck.stale && " CAUTION: a rule change made this deck illegal since."}
        </p>
        <p className="text-sm" data-testid="buy-cost">
          To buy: {deck.buyCostUsd > 0 ? priceText(deck.buyCostUsd) : "nothing. Every card is owned, or no price is known."}
        </p>
        {deck.summary && <p className="mt-2 max-w-measure leading-relaxed">{deck.summary}</p>}
      </header>

      {deck.revisionNote && (
        <section aria-labelledby={`revision-title-${deck.id}`} className="rounded-card border border-accent/40 bg-accent/5 p-3 text-sm">
          <h3 id={`revision-title-${deck.id}`} className="font-medium">
            What changed
          </h3>
          <p data-testid="revision-note">{deck.revisionNote}</p>
          {diff && (
            <DiffList diff={diff} testId="revision-diff" />
          )}
        </section>
      )}

      {findings.length > 0 && (
        <section aria-labelledby={`findings-title-${deck.id}`}>
          <h3 id={`findings-title-${deck.id}`} className="font-medium">
            Findings
          </h3>
          <ul className="list-disc pl-5 text-sm">
            {findings.map((f, i) => (
              <li key={i} className={f.severity === Severity.BLOCK ? "text-danger" : ""}>
                <span className="font-medium">{severityLabel(f.severity)}</span>
                {f.code && <span className="text-muted-foreground"> ({f.code})</span>}: {f.message}
                {f.oracleId && ` — ${nameOf(f.oracleId)}`}
              </li>
            ))}
          </ul>
        </section>
      )}

      <div className="text-sm">
        {cards.isPending && <p role="status">Loading card data...</p>}
        {cards.isPlaceholderData && <p role="status">Refreshing card data...</p>}
        {cards.isError && (
          <p role="alert" className="text-danger">
            Could not load the card data: {errorMessage(cards.error)}
          </p>
        )}
        {missing.length > 0 && (
          <p role="alert" className="text-danger">
            {missing.length} card{missing.length === 1 ? "" : "s"} of this deck are not in the card database:{" "}
            {missing.map(nameOf).join(", ")}.
          </p>
        )}
      </div>

      {cards.data && (
        <div className="@container shadow-card rounded-panel border border-border bg-card p-5 backdrop-blur-sm">
          <div className="grid items-start gap-8 @2xl:grid-cols-2">
          <table className="w-full text-sm">
            <caption className="mb-3 border-b border-border pb-2 text-left text-sm font-semibold tracking-wide uppercase">Mana curve, lands excluded</caption>
            <thead className="sr-only">
              <tr>
                <th scope="col">Mana value</th>
                <th scope="col">Cards</th>
              </tr>
            </thead>
            <tbody>
              {curveSteps.map((step, i) => (
                <tr key={step}>
                  <th scope="row" className="w-8 py-1 pr-3 text-left font-normal tabular-nums text-muted-foreground">
                    {step}
                  </th>
                  <td className="py-1">
                    <span className="flex items-center gap-3">
                      <span className="block h-2 max-w-64 grow rounded-full bg-muted" aria-hidden="true">
                        <span className="block h-2 rounded-full bg-accent transition-[width] duration-500" style={{ width: `${(curve[i] / curveMax) * 100}%` }} />
                      </span>
                      <span className="w-6 text-right tabular-nums">{curve[i]}</span>
                    </span>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
          <table className="w-full text-sm">
            <caption className="mb-3 border-b border-border pb-2 text-left text-sm font-semibold tracking-wide uppercase">Mana sources, cards that make each of the deck's colors</caption>
            <thead className="sr-only">
              <tr>
                <th scope="col">Color</th>
                <th scope="col">Sources</th>
              </tr>
            </thead>
            <tbody>
              {sourceRows.map((c) => {
                const n = sources.get(c.color) ?? 0;
                return (
                  <tr key={c.letter}>
                    <th scope="row" className="py-1 pr-3 text-left font-normal whitespace-nowrap">
                      <span className="flex items-center gap-2">
                        <span className={`inline-block size-3 rounded-full ring-1 ring-black/25 ${manaSwatch[c.letter] ?? "bg-mana-c"}`} aria-hidden="true" />
                        <span className="text-muted-foreground">{c.name}</span>
                      </span>
                    </th>
                    <td className="py-1">
                      <span className="flex items-center gap-3">
                        <span className="block h-2 max-w-64 grow rounded-full bg-muted" aria-hidden="true">
                          <span className={`block h-2 rounded-full transition-[width] duration-500 ${manaSwatch[c.letter] ?? "bg-mana-c"}`} style={{ width: `${(n / sourcesMax) * 100}%` }} />
                        </span>
                        <span className="w-6 text-right tabular-nums">{n}</span>
                      </span>
                    </td>
                  </tr>
                );
              })}
            </tbody>
          </table>
          </div>
        </div>
      )}

      <ExportPanel deck={deck} byId={byId} />

      <p className="text-xs text-muted-foreground">
        Card images and card text are unofficial Fan Content permitted under the Wizards of the Coast Fan Content Policy. They are
        copyright Wizards of the Coast, LLC, and come from Scryfall.
      </p>

      {commanderEntries.length > 0 && (
        <CardGroup title="Commander" count={commanderEntries.length} entries={commanderEntries} byId={byId} commanders={commanders} hideOwnership />
      )}
      {groups.map((g) => (
        <CardGroup key={g.role} title={roleLabel(g.role)} count={g.count} entries={g.cards} byId={byId} commanders={commanders} />
      ))}
      {deck.sideboard.length > 0 && (
        <CardGroup title="Sideboard" count={deck.sideboard.reduce((n, c) => n + c.count, 0)} entries={deck.sideboard} byId={byId} commanders={commanders} />
      )}
      {deck.upgrades.length > 0 && (
        <CardGroup title="Upgrades to buy" count={deck.upgrades.reduce((n, c) => n + c.count, 0)} entries={deck.upgrades} byId={byId} commanders={commanders} />
      )}
    </article>
  );
}

function CardGroup({
  title,
  count,
  entries,
  byId,
  commanders,
  hideOwnership = false,
}: {
  title: string;
  count: number;
  entries: DeckCard[];
  byId: Map<string, Card>;
  commanders: Set<string>;
  hideOwnership?: boolean;
}) {
  return (
    <section aria-label={`${title} (${count})`} className="@container">
      <h3 className="font-display mb-2.5 flex items-center gap-2 border-b border-border pb-2 text-xs tracking-[0.15em] text-muted-foreground uppercase">
        <span aria-hidden="true" className="inline-block h-3 w-0.5 bg-primary" />
        {title} <span>({count})</span>
      </h3>
      <ul className="mt-2 grid grid-cols-2 items-start gap-3 @xl:grid-cols-3 @3xl:grid-cols-4 @5xl:grid-cols-5 @7xl:grid-cols-6">
        {entries.map((e, i) => (
          <CardTile key={`${e.oracleId}-${i}`} entry={e} card={byId.get(e.oracleId)} isCommander={commanders.has(e.oracleId)} hideOwnership={hideOwnership} />
        ))}
      </ul>
    </section>
  );
}
