import { type Card, Color } from "@mtg/api-client/mtg/v1/card_pb";

// The colors of the game appear in a mana pip and nowhere else (D-329).
// Gold and purple are the app's own palette, and every deck wears them.

// wubrg is the order the game prints its colors in. Every list of colors
// in the app follows it, so two decks of the same colors read the same.
const wubrg: Color[] = [Color.W, Color.U, Color.B, Color.R, Color.G];

export const manaToken: Record<Color, string> = {
  [Color.W]: "var(--mana-w)",
  [Color.U]: "var(--mana-u)",
  [Color.B]: "var(--mana-b)",
  [Color.R]: "var(--mana-r)",
  [Color.G]: "var(--mana-g)",
  [Color.C]: "var(--mana-c)",
  [Color.UNSPECIFIED]: "var(--mana-c)",
};

// The ink of a pip. White and black need their own, because the game
// prints one nearly white and one nearly black.
export const manaInk: Record<Color, string> = {
  [Color.W]: "var(--mana-w-ink)",
  [Color.U]: "#fff",
  [Color.B]: "var(--mana-b-ink)",
  [Color.R]: "#fff",
  [Color.G]: "#fff",
  [Color.C]: "#fff",
  [Color.UNSPECIFIED]: "#fff",
};

export const manaLetter: Record<Color, string> = {
  [Color.W]: "W",
  [Color.U]: "U",
  [Color.B]: "B",
  [Color.R]: "R",
  [Color.G]: "G",
  [Color.C]: "C",
  [Color.UNSPECIFIED]: "C",
};

export const colorName: Record<Color, string> = {
  [Color.W]: "White",
  [Color.U]: "Blue",
  [Color.B]: "Black",
  [Color.R]: "Red",
  [Color.G]: "Green",
  [Color.C]: "Colorless",
  [Color.UNSPECIFIED]: "Colorless",
};

// sortIdentity drops duplicates and colorless, and puts the rest in WUBRG
// order. A colorless deck comes back empty, and the caller reads that as
// the colorless case.
export function sortIdentity(colors: Color[]): Color[] {
  const held = new Set(colors);
  return wubrg.filter((c) => held.has(c));
}

// identityOfCommanders reads the identity of the named commanders alone.
// The map may hold the commanders of many decks, as the library grid
// does, so nothing but the named ids counts. A card the index has not
// answered yet gives no color, and so does a commander that truly has
// none: both read as colorless, which is correct in both cases.
export function identityOfCommanders(commanderIds: string[], byId: Map<string, Card>): Color[] {
  const colors: Color[] = [];
  for (const id of commanderIds) {
    colors.push(...(byId.get(id)?.colorIdentity ?? []));
  }
  return sortIdentity(colors);
}

// identityOfCards reads the identity of every card in the map. Only a
// view that holds one deck's own cards may call it.
export function identityOfCards(byId: Map<string, Card>): Color[] {
  const colors: Color[] = [];
  for (const card of byId.values()) colors.push(...(card.colorIdentity ?? []));
  return sortIdentity(colors);
}

// identityLabel names the identity for a screen reader and for a filter,
// for example "White and Blue" or "Colorless".
export function identityLabel(colors: Color[]): string {
  const sorted = sortIdentity(colors);
  if (sorted.length === 0) return "Colorless";
  const names = sorted.map((c) => colorName[c]);
  if (names.length === 1) return names[0];
  return `${names.slice(0, -1).join(", ")} and ${names[names.length - 1]}`;
}

