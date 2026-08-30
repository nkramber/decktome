import { BookOpenIcon, LayersIcon, SparklesIcon } from "lucide-react";

// The three entries of the shell (D-311). The sidebar and the bottom bar
// read the same list, so a phone and a desktop never disagree.
// Build opens a new chat, the path the collection page already sends to.
export const navItems = [
  { to: "/session/new", label: "Build", icon: SparklesIcon, match: "/session" },
  { to: "/decks", label: "Decks", icon: LayersIcon, match: "/decks" },
  { to: "/collection", label: "Collection", icon: BookOpenIcon, match: "/collection" },
] as const;

// isActive marks the entry that owns the current path.
export function isActive(pathname: string, match: string): boolean {
  return pathname === match || pathname.startsWith(`${match}/`);
}
