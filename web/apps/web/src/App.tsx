import { useEffect, useState } from "react";

import { healthClient } from "./api";

type Health = { status: string; version: string } | { error: string } | null;

// legality_as_of surface (roadmap PR-3): every screen that shows deck data
// must show which day's card data it rests on.
type Freshness = { cardSnapshot: string; ageHours: number } | null;

export function App() {
  const [health, setHealth] = useState<Health>(null);
  const [freshness, setFreshness] = useState<Freshness>(null);

  useEffect(() => {
    let cancelled = false;
    fetch("/healthz")
      .then((res) => res.json())
      .then((body: { card_snapshot?: string; card_snapshot_age_hours?: number }) => {
        if (!cancelled && body.card_snapshot && body.card_snapshot !== "none") {
          setFreshness({ cardSnapshot: body.card_snapshot, ageHours: body.card_snapshot_age_hours ?? -1 });
        }
      })
      .catch(() => {
        /* the health RPC below reports connectivity problems */
      });
    return () => {
      cancelled = true;
    };
  }, []);

  useEffect(() => {
    let cancelled = false;
    healthClient
      .check({})
      .then((res) => {
        if (!cancelled) setHealth({ status: res.status, version: res.version });
      })
      .catch((err: unknown) => {
        if (!cancelled) setHealth({ error: err instanceof Error ? err.message : String(err) });
      });
    return () => {
      cancelled = true;
    };
  }, []);

  return (
    <main style={{ fontFamily: "system-ui, sans-serif", padding: "2rem" }}>
      <h1>MtG Deck Builder</h1>
      <p>PR-0a scaffold. The API health check below proves the proto pipeline.</p>
      <pre data-testid="health">{health ? JSON.stringify(health, null, 2) : "checking API..."}</pre>
      <p data-testid="freshness">
        {freshness
          ? `Card data as of ${freshness.cardSnapshot} (${freshness.ageHours.toFixed(1)} h old). Legality checks use this snapshot.`
          : "Card data: not loaded yet."}
      </p>
    </main>
  );
}
