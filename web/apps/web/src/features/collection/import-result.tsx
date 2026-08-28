import { UnresolvedReason } from "@mtg/api-client/mtg/v1/collection_pb";
import type { ImportCollectionResponse } from "@mtg/api-client/mtg/v1/collection_service_pb";

// ImportResult shows the counts and the unresolved rows of one upload (F-2, M-3).
export function ImportResult({ result }: { result: ImportCollectionResponse }) {
  const collection = result.collection;
  const report = result.report;
  const unresolved = report?.unresolved ?? [];
  const byReason = Object.entries(report?.unresolvedByReason ?? {});

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
              {reason}: {count}
            </li>
          ))}
        </ul>
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
                  <td>{UnresolvedReason[row.reason] ?? String(row.reason)}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </section>
  );
}
