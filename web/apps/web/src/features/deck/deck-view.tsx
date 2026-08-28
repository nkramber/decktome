import type { Card } from "@mtg/api-client/mtg/v1/card_pb";
import type { Deck, DeckCard } from "@mtg/api-client/mtg/v1/deck_pb";
import { CardRole, Severity } from "@mtg/api-client/mtg/v1/deck_pb";

import { errorMessage } from "../../lib/errors";
import { CardTile } from "./card-tile";
import {
  colorLetters,
  colorSources,
  curveSteps,
  diffDecks,
  formatLabel,
  groupByRole,
  manaCurve,
  powerLabel,
  priceText,
  roleLabel,
  severityLabel,
} from "./deck-stats";
import { useDeckCards } from "./use-cards";

// The deck view (ui plan, step 4). Cards group by role, each with its art
// and attribution (D-6), both faces for a DFC (F-9), the owned mark or
// the price (D-2, D-37), the findings, legality_as_of, the curve, and
// the color sources.
// base is the deck this one revised, when the page holds it (PR-12B).
export function DeckView({ deck, base }: { deck: Deck; base?: Deck }) {
  const diff = base && deck.revisedFromDeckId && base.id === deck.revisedFromDeckId ? diffDecks(base, deck) : undefined;
  const cards = useDeckCards(deck);
  const byId = cards.data?.byId ?? new Map<string, Card>();
  const missing = cards.data?.missing ?? [];
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
  const total = deck.cards.reduce((n, c) => n + c.count, 0);
  const validation = deck.validation;
  const findings = validation?.findings ?? [];
  const legalityAsOf = deck.legalityAsOf || validation?.legalityAsOf || "an unknown date";
  const curveMax = Math.max(1, ...curve);

  return (
    <article aria-labelledby={`deck-title-${deck.id}`} className="flex flex-col gap-4">
      <header className="flex flex-col gap-1">
        <h2 id={`deck-title-${deck.id}`} className="wrap-anywhere text-xl font-semibold">
          {deck.name || "Untitled deck"}
        </h2>
        <p className="text-sm text-neutral-700">
          {formatLabel(deck.format?.id, deck.format?.houseRules ?? "")}
          {powerLabel(deck.power) && ` · ${powerLabel(deck.power)}`}
          {` · ${total} cards`}
          {commanderCount > 0 && ` + ${commanderCount} commander${commanderCount > 1 ? "s" : ""}`}
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
        {deck.summary && <p className="mt-1">{deck.summary}</p>}
      </header>

      {deck.revisionNote && (
        <section aria-labelledby={`revision-title-${deck.id}`} className="rounded border border-blue-400 bg-blue-50 p-3 text-sm">
          <h3 id={`revision-title-${deck.id}`} className="font-medium">
            What changed
          </h3>
          <p data-testid="revision-note">{deck.revisionNote}</p>
          {diff && (
            <ul className="mt-2 list-disc pl-5" data-testid="revision-diff">
              {diff.removed.map((x) => (
                <li key={`r-${x}`}>Removed {x}</li>
              ))}
              {diff.added.map((x) => (
                <li key={`a-${x}`}>Added {x}</li>
              ))}
              {diff.changed.map((x) => (
                <li key={`c-${x}`}>Count of {x}</li>
              ))}
            </ul>
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
              <li key={i} className={f.severity === Severity.BLOCK ? "text-red-700" : ""}>
                <span className="font-medium">{severityLabel(f.severity)}</span>
                {f.code && <span className="text-neutral-600"> ({f.code})</span>}: {f.message}
                {f.oracleId && ` — ${nameOf(f.oracleId)}`}
              </li>
            ))}
          </ul>
        </section>
      )}

      <div className="text-sm">
        {cards.isPending && <p role="status">Loading card data...</p>}
        {cards.isError && (
          <p role="alert" className="text-red-700">
            Could not load the card data: {errorMessage(cards.error)}
          </p>
        )}
        {missing.length > 0 && (
          <p role="alert" className="text-red-700">
            {missing.length} card{missing.length === 1 ? "" : "s"} of this deck are not in the card database:{" "}
            {missing.map(nameOf).join(", ")}.
          </p>
        )}
      </div>

      {cards.data && (
        <div className="@container grid gap-4 @md:grid-cols-2">
          <table className="text-sm">
            <caption className="text-left font-medium">Mana curve, lands excluded</caption>
            <thead>
              <tr>
                <th scope="col" className="pr-2 text-left">
                  Mana value
                </th>
                <th scope="col" className="text-left">
                  Cards
                </th>
              </tr>
            </thead>
            <tbody>
              {curveSteps.map((step, i) => (
                <tr key={step}>
                  <th scope="row" className="pr-2 text-left font-normal">
                    {step}
                  </th>
                  <td>
                    <span className="flex items-center gap-2">
                      <span className="block h-3 w-24 max-w-full bg-neutral-200" aria-hidden="true">
                        <span className="block h-3 bg-neutral-700" style={{ width: `${(curve[i] / curveMax) * 100}%` }} />
                      </span>
                      {curve[i]}
                    </span>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
          <table className="text-sm">
            <caption className="text-left font-medium">Color sources, copies of cards that make each color</caption>
            <thead>
              <tr>
                <th scope="col" className="pr-2 text-left">
                  Color
                </th>
                <th scope="col" className="text-left">
                  Sources
                </th>
              </tr>
            </thead>
            <tbody>
              {colorLetters.map((c) => (
                <tr key={c.letter}>
                  <th scope="row" className="pr-2 text-left font-normal">
                    {c.name} ({c.letter})
                  </th>
                  <td>{sources.get(c.color) ?? 0}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}

      <p className="text-xs text-neutral-600">
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
      <h3 className="font-medium">
        {title} <span className="text-neutral-600">({count})</span>
      </h3>
      <ul className="mt-2 grid grid-cols-1 items-start gap-2 @sm:grid-cols-2 @2xl:grid-cols-3 @4xl:grid-cols-4">
        {entries.map((e, i) => (
          <CardTile key={`${e.oracleId}-${i}`} entry={e} card={byId.get(e.oracleId)} isCommander={commanders.has(e.oracleId)} hideOwnership={hideOwnership} />
        ))}
      </ul>
    </section>
  );
}
