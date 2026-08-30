import { type Card, Color } from "@mtg/api-client/mtg/v1/card_pb";
import { CardRole } from "@mtg/api-client/mtg/v1/deck_pb";
import type { CSSProperties } from "react";

// The five colors of the game are the palette (D-327). A deck carries its
// own identity, and the page it owns takes that identity as its accent.

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

// identityVars gives one element the two stops of its identity wash. One
// color washes in its own hue, two or more wash from the first to the
// last, and a colorless deck washes in steel.
export function identityVars(colors: Color[]): CSSProperties {
  const sorted = sortIdentity(colors);
  if (sorted.length === 0) {
    return { "--identity-a": manaToken[Color.C], "--identity-b": manaToken[Color.C] } as CSSProperties;
  }
  const first = manaToken[sorted[0]];
  const last = manaToken[sorted[sorted.length - 1]];
  return { "--identity-a": first, "--identity-b": last } as CSSProperties;
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

// roleToken gives one card role its hue, so the eye finds the shape of a
// deck before it reads one word.
export const roleToken: Record<CardRole, string> = {
  [CardRole.LAND]: "var(--role-land)",
  [CardRole.RAMP]: "var(--role-ramp)",
  [CardRole.DRAW]: "var(--role-draw)",
  [CardRole.REMOVAL]: "var(--role-removal)",
  [CardRole.WIPE]: "var(--role-wipe)",
  [CardRole.THREAT]: "var(--role-threat)",
  [CardRole.INTERACTION]: "var(--role-interaction)",
  [CardRole.SYNERGY]: "var(--role-synergy)",
  [CardRole.WINCON]: "var(--role-wincon)",
  [CardRole.OTHER]: "var(--role-other)",
  [CardRole.UNSPECIFIED]: "var(--role-other)",
};
