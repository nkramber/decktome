import { describe, expect, it, vi } from "vitest";

import { bookmarkletHref, bookmarkletSource, readImportFragment } from "./moxfield-bookmarklet";

// entry is one card of a board of the v3 API, in the shape that the
// console test of the owner read (D-1192).
function entry(name: string, quantity: number) {
  return { quantity, boardType: "mainboard", finish: "nonFoil", isFoil: false, card: { id: name, name, set: "x", type_line: "" } };
}

function board(...cards: [string, number][]) {
  const out: Record<string, ReturnType<typeof entry>> = {};
  cards.forEach(([n, q], i) => (out[`k${String(i)}`] = entry(n, q)));
  return { count: cards.reduce((s, [, q]) => s + q, 0), cards: out };
}

const deck = {
  id: "5oFmOlOqW0uzHKGvlzzAsA",
  name: "Azula Tempo",
  format: "commander",
  boards: {
    commanders: board(["Fire Lord Azula", 1]),
    companions: board(),
    mainboard: board(["Island", 12], ["Arcane Signet", 1], ["Fire // Ice", 1]),
    sideboard: board(),
    maybeboard: board(["Counterspell", 1]),
    tokens: board(["Treasure", 1]),
  },
};

type Run = { opened: string[]; alerts: string[]; fetched: string[]; href: string };

// run executes the code of the bookmark, as the browser does, with a fake
// page, a fake fetch, and a fake window.
async function run(opts: { host?: string; path?: string; answer?: { ok: boolean; status: number; body?: unknown }; blocked?: boolean }): Promise<Run> {
  const out: Run = { opened: [], alerts: [], fetched: [], href: "" };
  const location = {
    hostname: opts.host ?? "moxfield.com",
    pathname: opts.path ?? "/decks/5oFmOlOqW0uzHKGvlzzAsA",
    set href(v: string) {
      out.href = v;
    },
  };
  const answer = opts.answer ?? { ok: true, status: 200, body: deck };
  const fetch = vi.fn((url: string) => {
    out.fetched.push(url);
    return Promise.resolve({ ok: answer.ok, status: answer.status, json: () => Promise.resolve(answer.body) });
  });
  const win = {
    open: (url: string) => {
      out.opened.push(url);
      return opts.blocked ? null : { opener: {} };
    },
  };
  const source = bookmarkletSource("https://decktome.com").replace(/;$/, "");
  const start = new Function("location", "fetch", "alert", "window", `return ${source}`) as (...a: unknown[]) => Promise<void>;
  await start(location, fetch, (m: string) => out.alerts.push(m), win);
  return out;
}

describe("the Moxfield bookmarklet (D-1193)", () => {
  it("reads the open deck and opens the deck list of decktome with the list in the fragment", async () => {
    const r = await run({});
    expect(r.fetched).toEqual(["https://api2.moxfield.com/v3/decks/all/5oFmOlOqW0uzHKGvlzzAsA"]);
    expect(r.alerts).toEqual([]);
    expect(r.opened).toHaveLength(1);
    const url = new URL(r.opened[0] ?? "");
    expect(url.origin + url.pathname).toBe("https://decktome.com/decks");
    expect(url.search).toBe("");
    expect(readImportFragment(url.hash)).toEqual({
      list: "Commander\n1 Fire Lord Azula\n\nDeck\n12 Island\n1 Arcane Signet\n1 Fire // Ice",
      name: "Azula Tempo",
    });
  });

  it("leaves out the maybeboard and the tokens", async () => {
    const r = await run({});
    const list = readImportFragment(new URL(r.opened[0] ?? "").hash)?.list ?? "";
    expect(list).not.toContain("Counterspell");
    expect(list).not.toContain("Treasure");
  });

  it("writes the companion and the sideboard under their headers", async () => {
    const withSide = { ...deck, boards: { ...deck.boards, companions: board(["Lurrus of the Dream-Den", 1]), sideboard: board(["Negate", 2]) } };
    const r = await run({ answer: { ok: true, status: 200, body: withSide } });
    const list = readImportFragment(new URL(r.opened[0] ?? "").hash)?.list ?? "";
    expect(list).toContain("Companion\n1 Lurrus of the Dream-Den");
    expect(list.endsWith("Sideboard\n2 Negate")).toBe(true);
  });

  it("opens the list in the same tab when the browser blocks the new tab", async () => {
    const r = await run({ blocked: true });
    expect(r.opened).toHaveLength(1);
    expect(r.href).toBe(r.opened[0]);
  });

  it("reads nothing on a page that is not a Moxfield deck", async () => {
    for (const page of [{ host: "archidekt.com" }, { path: "/search" }, { host: "notmoxfield.com" }]) {
      const r = await run(page);
      expect(r.fetched).toEqual([]);
      expect(r.opened).toEqual([]);
      expect(r.alerts).toEqual(["Open a deck on moxfield.com, then select the decktome bookmark again."]);
    }
  });

  it("names the answer of Moxfield when the read fails, and opens nothing", async () => {
    const r = await run({ answer: { ok: false, status: 403 } });
    expect(r.opened).toEqual([]);
    expect(r.alerts[0]).toContain("Moxfield answered 403");
  });

  it("opens nothing for a deck with no cards", async () => {
    const r = await run({ answer: { ok: true, status: 200, body: { name: "Empty", boards: {} } } });
    expect(r.opened).toEqual([]);
    expect(r.alerts).toEqual(["This Moxfield deck holds no cards."]);
  });

  it("is a javascript: link that holds the code for the origin", () => {
    const href = bookmarkletHref("https://decktome.com");
    expect(href.startsWith("javascript:")).toBe(true);
    expect(decodeURIComponent(href.slice("javascript:".length))).toBe(bookmarkletSource("https://decktome.com"));
  });
});

describe("readImportFragment", () => {
  it("reads null for a fragment with no list", () => {
    expect(readImportFragment("")).toBeNull();
    expect(readImportFragment("#")).toBeNull();
    expect(readImportFragment("#section-2")).toBeNull();
    expect(readImportFragment("#list=%20%0A")).toBeNull();
  });

  it("reads null for a fragment over the cap of the import", () => {
    expect(readImportFragment(`#list=${"a".repeat(400 * 1024)}`)).toBeNull();
  });

  it("caps the name at 200 characters", () => {
    expect(readImportFragment(`#list=1%20Island&name=${"n".repeat(300)}`)?.name).toHaveLength(200);
  });
});
