import { BookOpenIcon, LayersIcon, SparklesIcon } from "lucide-react";

// The three entries of the shell (D-311, D-433). The top bar reads this
// list on every screen.
// Build opens a new chat (D-436). The chat is the whole start, and the
// pool picker in it is the one control outside the conversation.
export const navItems = [
  { to: "/session/new", label: "Build", icon: SparklesIcon, match: ["/session"] },
  { to: "/decks", label: "Decks", icon: LayersIcon, match: ["/decks"] },
  { to: "/collection", label: "Collection", icon: BookOpenIcon, match: ["/collection"] },
] as const;

// isActive marks the entry that owns the current path.
export function isActive(pathname: string, match: readonly string[]): boolean {
  return match.some((m) => pathname === m || pathname.startsWith(`${m}/`));
}
