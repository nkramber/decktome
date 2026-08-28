import { UnresolvedReason } from "@mtg/api-client/mtg/v1/collection_pb";
import type { ImportCollectionResponse } from "@mtg/api-client/mtg/v1/collection_service_pb";

// reasonLabel turns an UnresolvedReason into words a user can act on. The
// server keys unresolvedByReason by the full enum name, and the row carries
// the number, so both shapes land here.
const reasonLabels: Record<string, string> = {
  UNKNOWN_CARD: "Unknown card: the name and printing match nothing in the card database",
  NON_ENGLISH: "Non-English printing: the app reads English cards only (D-23)",
  BAD_ROW: "Bad row: the line does not parse as a ManaBox row",
  NOT_PLAYABLE: "Not a playable card: a token, emblem, or art card",
  UNKNOWN_VALUE: "Unknown value: a finish or condition the app does not know",
};

export function reasonLabel(reason: string | number): string {
  const name = typeof reason === "number" ? (UnresolvedReason[reason] ?? String(reason)) : reason;
  const short = name.replace(/^UNRESOLVED_REASON_/, "");
  return reasonLabels[short] ?? short;
}

// ImportResult shows the counts and the unresolved rows of one upload (F-2, M-3).
export function ImportResult({ result }: { result: ImportCollectionResponse }) {
  const collection = result.collection;
  const report = result.report;
  const unresolved = report?.unresolved ?? [];
  const byReason = Object.entries(report?.unresolvedByReason ?? {});
  // ManaBox exports tokens as collection rows. The importer refuses them
  // on purpose (F-21), so the note says why before the user asks.
  const hasTokens = byReason.some(([reason, count]) => reason.endsWith("NOT_PLAYABLE") && count > 0);

  return (
    <section aria-live="polite" className="flex flex-col gap-3 rounded border border-neutral-200 p-4">
      <h2 className="text-lg font-medium">Import result</h2>
      <p data-testid="card-count">
        {collection
          ? `${collection.name}: ${collection.cardCount} cards, ${report?.resolvedCount ?? 0} rows resolved, ${unresolved.length} unresolved.`
          : "The import returned no collection."}
      </p>
      {byReason.length > 0 && (
        <ul className="text-sm">
          {byReason.map(([reason, count]) => (
            <li key={reason}>
              {reasonLabel(reason)}: {count}
            </li>
          ))}
        </ul>
      )}
      {hasTokens && (
        <p className="text-sm text-neutral-600" data-testid="token-note">
          ManaBox lists tokens in a collection. A token is not a card a deck can use, so the importer
          leaves it out of the count and shows the row here.
        </p>
      )}
      {unresolved.length > 0 && (
        <div className="overflow-x-auto">
          <table className="w-full text-left text-sm">
            <caption className="text-left font-medium">Unresolved rows</caption>
            <thead>
              <tr>
                <th scope="col" className="pr-3">
                  Line
                </th>
                <th scope="col" className="pr-3">
                  Row
                </th>
                <th scope="col">Reason</th>
              </tr>
            </thead>
            <tbody>
              {unresolved.map((row) => (
                <tr key={row.line}>
                  <td className="pr-3">{row.line}</td>
                  <td className="pr-3 font-mono">{row.raw}</td>
                  <td>{reasonLabel(row.reason)}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </section>
  );
}
