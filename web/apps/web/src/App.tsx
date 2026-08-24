import { useEffect, useState } from "react";

import { healthClient } from "./api";

type Health = { status: string; version: string } | { error: string } | null;

export function App() {
  const [health, setHealth] = useState<Health>(null);

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
    </main>
  );
}
