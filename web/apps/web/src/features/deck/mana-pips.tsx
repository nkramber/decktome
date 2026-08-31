import type { Color } from "@mtg/api-client/mtg/v1/card_pb";

import { cn } from "../../lib/cn";
import { identityLabel, manaInk, manaLetter, manaToken, sortIdentity } from "./color-identity";

// The identity of a deck, as the lettered pips the game itself prints
// (D-329). The label carries the words, so a reader who sees no color
// loses nothing.
export function ManaPips({ colors, className, size = "md" }: { colors: Color[]; className?: string; size?: "sm" | "md" }) {
  const pips = sortIdentity(colors);
  return (
    <span className={cn("inline-flex items-center gap-0.5", className)}>
      <span className="sr-only">{identityLabel(colors)}</span>
      {pips.length === 0 ? (
        <span className="font-mono text-[11px] text-muted-foreground" aria-hidden="true">
          Colorless
        </span>
      ) : (
        pips.map((c) => (
          <span
            key={c}
            aria-hidden="true"
            className={cn(
              "font-display inline-flex items-center justify-center rounded-full leading-none font-bold",
              size === "sm" ? "size-4 text-[9px]" : "size-5 text-[10px]",
            )}
            style={{ backgroundColor: manaToken[c], color: manaInk[c] }}
          >
            {manaLetter[c]}
          </span>
        ))
      )}
    </span>
  );
}
