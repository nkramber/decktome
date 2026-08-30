import type { Color } from "@mtg/api-client/mtg/v1/card_pb";

import { cn } from "../../lib/cn";
import { identityLabel, manaToken, sortIdentity } from "./color-identity";

// The identity of a deck, as the pips the game itself uses. The label
// carries the words, so a reader who sees no color loses nothing (D-327).
export function ManaPips({ colors, className, size = "md" }: { colors: Color[]; className?: string; size?: "sm" | "md" }) {
  const sorted = sortIdentity(colors);
  const pips = sorted.length > 0 ? sorted : [];
  return (
    <span className={cn("inline-flex items-center gap-1", className)}>
      <span className="sr-only">{identityLabel(colors)}</span>
      {pips.length === 0 ? (
        <span
          aria-hidden="true"
          className={cn("inline-block rounded-full ring-1 ring-black/20", size === "sm" ? "size-2" : "size-2.5")}
          style={{ backgroundColor: "var(--mana-c)" }}
        />
      ) : (
        pips.map((c) => (
          <span
            key={c}
            aria-hidden="true"
            className={cn("inline-block rounded-full ring-1 ring-black/20", size === "sm" ? "size-2" : "size-2.5")}
            style={{ backgroundColor: manaToken[c] }}
          />
        ))
      )}
    </span>
  );
}
