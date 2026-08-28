import { type Card, Color } from "@mtg/api-client/mtg/v1/card_pb";
import { CardRole, type DeckCard } from "@mtg/api-client/mtg/v1/deck_pb";
import { FormatId } from "@mtg/api-client/mtg/v1/format_pb";
import { describe, expect, it } from "vitest";

import { colorSources, formatLabel, groupByRole, manaCurve, powerLabel } from "./deck-stats";

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
});
