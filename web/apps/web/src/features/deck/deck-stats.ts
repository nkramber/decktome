import { type Card, Color } from "@mtg/api-client/mtg/v1/card_pb";
import { CardRole, type Deck, type DeckCard, Severity } from "@mtg/api-client/mtg/v1/deck_pb";
import { FormatId, type PowerLevel, SixtyStep } from "@mtg/api-client/mtg/v1/format_pb";

// Pure helpers for the deck view. The mana curve and the color sources
// come from the card data on the client (ui plan, section 4).

export const roleOrder: CardRole[] = [
  CardRole.LAND,
  CardRole.RAMP,
  CardRole.DRAW,
  CardRole.REMOVAL,
  CardRole.WIPE,
  CardRole.THREAT,
  CardRole.INTERACTION,
  CardRole.SYNERGY,
  CardRole.WINCON,
  CardRole.OTHER,
  CardRole.UNSPECIFIED,
];

const roleLabels: Record<CardRole, string> = {
  [CardRole.LAND]: "Lands",
  [CardRole.RAMP]: "Ramp",
  [CardRole.DRAW]: "Card draw",
  [CardRole.REMOVAL]: "Removal",
  [CardRole.WIPE]: "Board wipes",
  [CardRole.THREAT]: "Threats",
  [CardRole.INTERACTION]: "Interaction",
  [CardRole.SYNERGY]: "Synergy",
  [CardRole.WINCON]: "Win conditions",
  [CardRole.OTHER]: "Other",
  [CardRole.UNSPECIFIED]: "Unsorted",
};

export function roleLabel(role: CardRole): string {
  return roleLabels[role] ?? "Unsorted";
}

export type RoleGroup = { role: CardRole; cards: DeckCard[]; count: number };

// groupByRole keeps the roleOrder and drops empty roles.
export function groupByRole(cards: DeckCard[]): RoleGroup[] {
  return roleOrder
    .map((role) => {
      const inRole = cards.filter((c) => c.role === role);
      return { role, cards: inRole, count: inRole.reduce((n, c) => n + c.count, 0) };
    })
    .filter((g) => g.cards.length > 0);
}

export function isLand(card: Card | undefined): boolean {
  return card?.cardTypes.includes("Land") ?? false;
}

// The curve has eight steps: mana value 0 to 6, then "7+". Lands stay out.
export const curveSteps = ["0", "1", "2", "3", "4", "5", "6", "7+"];

export function manaCurve(cards: DeckCard[], byId: Map<string, Card>): number[] {
  const curve = new Array<number>(curveSteps.length).fill(0);
  for (const dc of cards) {
    const card = byId.get(dc.oracleId);
    if (!card || isLand(card)) continue;
    const step = Math.min(Math.floor(card.manaValue), curveSteps.length - 1);
    curve[step] += dc.count;
  }
  return curve;
}

export const colorLetters: { color: Color; letter: string; name: string }[] = [
  { color: Color.W, letter: "W", name: "White" },
  { color: Color.U, letter: "U", name: "Blue" },
  { color: Color.B, letter: "B", name: "Black" },
  { color: Color.R, letter: "R", name: "Red" },
  { color: Color.G, letter: "G", name: "Green" },
  { color: Color.C, letter: "C", name: "Colorless" },
];

// deckColors is the set of colors the deck's cards are, plus the color
// identity of each commander. A colorless deck gives an empty set, and
// the sources table then shows colorless only.
export function deckColors(cards: DeckCard[], byId: Map<string, Card>, commanderIds: string[] = []): Set<Color> {
  const out = new Set<Color>();
  for (const dc of cards) {
    for (const c of byId.get(dc.oracleId)?.colors ?? []) out.add(c);
  }
  for (const id of commanderIds) {
    for (const c of byId.get(id)?.colorIdentity ?? []) out.add(c);
  }
  return out;
}

// colorSources counts the copies of every card that can produce each color.
// A dual land counts once for each of its colors. A rock that makes any
// color counts for every color, so the view shows the deck's own colors.
export function colorSources(cards: DeckCard[], byId: Map<string, Card>): Map<Color, number> {
  const sources = new Map<Color, number>(colorLetters.map((c) => [c.color, 0]));
  for (const dc of cards) {
    const card = byId.get(dc.oracleId);
    if (!card) continue;
    for (const color of new Set(card.producedMana)) {
      if (sources.has(color)) sources.set(color, (sources.get(color) ?? 0) + dc.count);
    }
  }
  return sources;
}

export function formatLabel(id: FormatId | undefined, houseRules: string): string {
  switch (id) {
    case FormatId.COMMANDER:
      return "Commander";
    case FormatId.STANDARD:
      return "Standard";
    case FormatId.MODERN:
      return "Modern";
    case FormatId.HOUSE:
      return houseRules ? `House rules: ${houseRules}` : "House rules";
    default:
      return "Unknown format";
  }
}

export function powerLabel(power: PowerLevel | undefined): string {
  if (!power) return "";
  if (power.level.case === "bracket") return `Bracket ${power.level.value}`;
  if (power.level.case === "sixtyStep") {
    const labels: Record<SixtyStep, string> = {
      [SixtyStep.CASUAL]: "Casual",
      [SixtyStep.FNM]: "FNM",
      [SixtyStep.TOURNAMENT]: "Tournament",
      [SixtyStep.UNSPECIFIED]: "",
    };
    return labels[power.level.value] ?? "";
  }
  return "";
}

export function severityLabel(s: Severity): string {
  switch (s) {
    case Severity.BLOCK:
      return "Block";
    case Severity.WARN:
      return "Warning";
    case Severity.INFO:
      return "Info";
    default:
      return "";
  }
}

export function priceText(usd: number): string {
  return usd > 0 ? `$${usd.toFixed(2)}` : "no price";
}

// DeckDiff is what changed between a deck and the one it revised
// (PR-12B). The deck view shows it under the header.
export type DeckDiff = { added: string[]; removed: string[]; changed: string[] };

// The key is the oracle id (the name when the entry has none), and the
// sideboard is its own zone, so a card that moves between zones shows
// as one removal and one addition.
type Entry = { label: string; count: number };

function countByKey(deck: Deck): Map<string, Entry> {
  const out = new Map<string, Entry>();
  const add = (cards: DeckCard[], zone: string) => {
    for (const c of cards) {
      const key = `${c.oracleId || c.name}${zone}`;
      const prev = out.get(key);
      out.set(key, { label: `${c.name}${zone}`, count: (prev?.count ?? 0) + c.count });
    }
  };
  add(deck.cards, "");
  add(deck.sideboard, " (sideboard)");
  return out;
}

export function diffDecks(base: Deck, revised: Deck): DeckDiff {
  const before = countByKey(base);
  const after = countByKey(revised);
  const diff: DeckDiff = { added: [], removed: [], changed: [] };
  for (const [key, { label, count: n }] of after) {
    const m = before.get(key)?.count;
    if (m === undefined) diff.added.push(`${n} ${label}`);
    else if (m !== n) diff.changed.push(`${label}: ${m} to ${n}`);
  }
  for (const [key, { label, count: m }] of before) {
    if (!after.has(key)) diff.removed.push(`${m} ${label}`);
  }
  diff.added.sort();
  diff.removed.sort();
  diff.changed.sort();
  return diff;
}

// deckSize is the count a reader expects: the main deck plus the command
// zone, so a Commander deck reads 100 and not 99 (D-454).
export function deckSize(mainCount: number, commanders: number): number {
  return mainCount + commanders;
}

// The deck view of PR-20: the type counts, the average mana value, the
// filters, and the sorts, all from the card data on the client.

// mainTypes are the card types the type table counts, in the order the
// game lists them. A card counts once per type it holds, so an artifact
// creature counts twice.
export const mainTypes = ["Creature", "Instant", "Sorcery", "Artifact", "Enchantment", "Planeswalker", "Battle", "Land"];

export function typeCounts(cards: DeckCard[], byId: Map<string, Card>): Map<string, number> {
  const out = new Map<string, number>(mainTypes.map((t) => [t, 0]));
  for (const dc of cards) {
    const card = byId.get(dc.oracleId);
    if (!card) continue;
    for (const t of new Set(card.cardTypes)) {
      if (out.has(t)) out.set(t, (out.get(t) ?? 0) + dc.count);
    }
  }
  return out;
}

// averageManaValue is the mean mana value of the nonland cards, copies
// counted. A deck with no nonland card known reads zero.
export function averageManaValue(cards: DeckCard[], byId: Map<string, Card>): number {
  let sum = 0;
  let n = 0;
  for (const dc of cards) {
    const card = byId.get(dc.oracleId);
    if (!card || isLand(card)) continue;
    sum += card.manaValue * dc.count;
    n += dc.count;
  }
  return n === 0 ? 0 : sum / n;
}

// Filters narrow the card list. An empty field means no limit on it.
export type Filters = {
  role?: CardRole;
  color?: Color;
  // manaValue is a curve step: 0 to 6, or 7 for "7+".
  manaValue?: number;
  type?: string;
  owned?: "owned" | "to-buy";
};

export function filterEntries(entries: DeckCard[], byId: Map<string, Card>, f: Filters): DeckCard[] {
  return entries.filter((dc) => {
    const card = byId.get(dc.oracleId);
    if (f.role !== undefined && dc.role !== f.role) return false;
    if (f.owned === "owned" && !dc.owned) return false;
    if (f.owned === "to-buy" && dc.owned) return false;
    if (f.color !== undefined) {
      if (!card) return false;
      const colors = card.colors.length > 0 ? card.colors : [Color.C];
      if (!colors.includes(f.color)) return false;
    }
    if (f.manaValue !== undefined) {
      if (!card || isLand(card)) return false;
      const step = Math.min(Math.floor(card.manaValue), curveSteps.length - 1);
      if (step !== f.manaValue) return false;
    }
    if (f.type !== undefined && f.type !== "") {
      if (!card || !card.cardTypes.includes(f.type)) return false;
    }
    return true;
  });
}

export type SortKey = "role" | "mana-value" | "name" | "price";

// sortEntries orders the cards. "role" keeps the deck's own order, and
// the other keys sort inside the whole list, with the name as the tie
// breaker. A card with no data sorts last on mana value.
export function sortEntries(entries: DeckCard[], byId: Map<string, Card>, key: SortKey): DeckCard[] {
  if (key === "role") return entries;
  const out = [...entries];
  const byName = (a: DeckCard, b: DeckCard) => a.name.localeCompare(b.name);
  out.sort((a, b) => {
    switch (key) {
      case "name":
        return byName(a, b);
      case "price":
        return b.priceUsd - a.priceUsd || byName(a, b);
      case "mana-value": {
        const ma = byId.get(a.oracleId)?.manaValue ?? Number.POSITIVE_INFINITY;
        const mb = byId.get(b.oracleId)?.manaValue ?? Number.POSITIVE_INFINITY;
        return ma - mb || byName(a, b);
      }
      default:
        return 0;
    }
  });
  return out;
}
