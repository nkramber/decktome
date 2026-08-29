import { type Card, Color } from "@mtg/api-client/mtg/v1/card_pb";
import { CardRole, type Deck, type DeckCard, Severity } from "@mtg/api-client/mtg/v1/deck_pb";
import { FormatId } from "@mtg/api-client/mtg/v1/format_pb";
import { describe, expect, it } from "vitest";

import { colorSources, deckColors, diffDecks, formatLabel, groupByRole, manaCurve, powerLabel, priceText, roleLabel, severityLabel } from "./deck-stats";

const card = (id: string, manaValue: number, types: string[], produced: Color[] = []) =>
  ({ oracleId: id, manaValue, cardTypes: types, producedMana: produced }) as unknown as Card;
const dc = (id: string, count: number, role: CardRole) => ({ oracleId: id, count, role }) as unknown as DeckCard;

const byId = new Map<string, Card>([
  ["forest", card("forest", 0, ["Land"], [Color.G])],
  ["dual", card("dual", 0, ["Land"], [Color.G, Color.W])],
  ["elf", card("elf", 1, ["Creature"], [Color.G])],
  ["big", card("big", 9.5, ["Creature"])],
  ["mid", card("mid", 3, ["Instant"])],
]);

describe("deck-stats", () => {
  it("groups by role in the fixed order and drops empty roles", () => {
    const groups = groupByRole([dc("big", 1, CardRole.THREAT), dc("forest", 10, CardRole.LAND), dc("mid", 2, CardRole.REMOVAL)]);
    expect(groups.map((g) => g.role)).toEqual([CardRole.LAND, CardRole.REMOVAL, CardRole.THREAT]);
    expect(groups[0].count).toBe(10);
  });

  it("builds the curve without lands and caps at 7+", () => {
    const curve = manaCurve(
      [dc("forest", 10, CardRole.LAND), dc("elf", 4, CardRole.RAMP), dc("big", 1, CardRole.THREAT), dc("mid", 2, CardRole.REMOVAL), dc("unknown", 3, CardRole.OTHER)],
      byId,
    );
    expect(curve).toEqual([0, 4, 0, 2, 0, 0, 0, 1]);
  });

  it("counts a dual land once per color, copies included", () => {
    const sources = colorSources([dc("forest", 10, CardRole.LAND), dc("dual", 2, CardRole.LAND), dc("elf", 4, CardRole.RAMP)], byId);
    expect(sources.get(Color.G)).toBe(16);
    expect(sources.get(Color.W)).toBe(2);
    expect(sources.get(Color.U)).toBe(0);
  });

  it("labels the format and the power", () => {
    expect(formatLabel(FormatId.COMMANDER, "")).toBe("Commander");
    expect(formatLabel(FormatId.HOUSE, "no bans")).toBe("House rules: no bans");
    expect(powerLabel({ level: { case: "bracket", value: 3 } } as never)).toBe("Bracket 3");
    expect(powerLabel(undefined)).toBe("");
  });

  it("deckColors unions the card colors with the commander's color identity", () => {
    const withColors = new Map<string, Card>([
      ["elf", { oracleId: "elf", colors: [Color.G], colorIdentity: [Color.G] } as unknown as Card],
      ["rock", { oracleId: "rock", colors: [], colorIdentity: [] } as unknown as Card],
      ["cmd", { oracleId: "cmd", colors: [Color.G], colorIdentity: [Color.G, Color.U] } as unknown as Card],
    ]);
    expect([...deckColors([dc("elf", 4, CardRole.RAMP), dc("rock", 1, CardRole.RAMP)], withColors)]).toEqual([Color.G]);
    expect([...deckColors([dc("rock", 1, CardRole.RAMP)], withColors, ["cmd"])].sort()).toEqual([Color.G, Color.U].sort());
    expect(deckColors([dc("rock", 1, CardRole.RAMP)], withColors, ["unknown"]).size).toBe(0);
  });

  it("diffDecks keys by oracle id and treats the sideboard as its own zone", () => {
    const entry = (oracleId: string, name: string, count: number) => ({ oracleId, name, count }) as unknown as DeckCard;
    const base = { cards: [entry("o-a", "Alpha", 4), entry("o-b", "Beta", 2), entry("o-r", "Renamed", 1)], sideboard: [entry("o-s", "Side", 2)] } as unknown as Deck;
    const revised = {
      cards: [entry("o-a", "Alpha", 2), entry("o-c", "Gamma", 1), entry("o-r", "Renamed (new)", 1), entry("o-s", "Side", 1)],
      sideboard: [entry("o-b", "Beta", 2)],
    } as unknown as Deck;
    expect(diffDecks(base, revised)).toEqual({
      added: ["1 Gamma", "1 Side", "2 Beta (sideboard)"],
      removed: ["2 Beta", "2 Side (sideboard)"],
      changed: ["Alpha: 4 to 2"],
    });
  });

  it("labels prices, severities, and roles", () => {
    expect(priceText(0)).toBe("no price");
    expect(priceText(1.5)).toBe("$1.50");
    expect(severityLabel(Severity.BLOCK)).toBe("Block");
    expect(severityLabel(Severity.WARN)).toBe("Warning");
    expect(severityLabel(Severity.INFO)).toBe("Info");
    expect(severityLabel(Severity.UNSPECIFIED)).toBe("");
    expect(roleLabel(CardRole.WINCON)).toBe("Win conditions");
    expect(roleLabel(99 as CardRole)).toBe("Unsorted");
  });
});
