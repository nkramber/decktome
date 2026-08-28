import { useQuery } from "@tanstack/react-query";

import { healthClient } from "../../lib/api";

// legality_as_of surface (roadmap PR-3): every screen that shows deck data
// must show which day's card data it rests on. The Check RPC carries the
// snapshot date and age, so one call fills the footer.
export function HealthFooter() {
  const health = useQuery({
    queryKey: ["health"],
    queryFn: () => healthClient.check({}),
    refetchInterval: 60_000,
  });

  let api: string;
  let cards: string;
  if (health.isPending) {
    api = "API: checking...";
    cards = "Card data: not loaded yet.";
  } else if (health.isError) {
    api = `API: error (${health.error.message})`;
    cards = "Card data: unknown. The API did not answer.";
  } else {
    const res = health.data;
    api = `API: ${res.status}, version ${res.version}`;
    cards =
      res.cardSnapshot && res.cardSnapshot !== "none"
        ? `Card data as of ${res.cardSnapshot} (${res.cardSnapshotAgeHours.toFixed(1)} h old). Legality checks use this snapshot.`
        : "Card data: not loaded yet.";
  }

  return (
    <footer className="border-t border-neutral-200 px-6 py-2 text-sm text-neutral-600" aria-live="polite">
      <span data-testid="health">{api}</span>
      {" · "}
      <span data-testid="freshness">{cards}</span>
    </footer>
  );
}
