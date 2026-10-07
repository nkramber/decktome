// The Moxfield bookmarklet (D-1192, D-1193). Moxfield refuses each read
// of the server (D-1102), and the deck page of Moxfield reads its own
// API. So the reader runs this code on that page, in the browser where
// Moxfield already answers. It reads the one open deck, writes an Arena
// list with the section headers of the parser, and opens the deck list
// of decktome with the list in the fragment. A fragment never reaches a
// server. The import dialog opens with the list, and the reader selects
// Import. The deck keeps no Moxfield link, because the server never read
// the deck (D-1107).
//
// The body is plain JavaScript in a string, so the bookmark holds the
// same code that the test runs. It reads O, the origin of decktome.

const body = `
const m = /(^|\\.)moxfield\\.com$/.test(location.hostname) && location.pathname.match(/^\\/decks\\/([^/?#]+)/);
if (!m) { alert("Open a deck on moxfield.com, then select the decktome bookmark again."); return; }
let d;
try {
  const r = await fetch("https://api2.moxfield.com/v3/decks/all/" + encodeURIComponent(m[1]));
  if (!r.ok) throw new Error("Moxfield answered " + r.status);
  d = await r.json();
} catch (e) {
  alert("decktome could not read this deck: " + e.message + ". Copy the list by hand: More, Export, Copy for Arena.");
  return;
}
const parts = [];
for (const [board, head] of [["commanders", "Commander"], ["companions", "Companion"], ["mainboard", "Deck"], ["sideboard", "Sideboard"]]) {
  const cards = (d && d.boards && d.boards[board] && d.boards[board].cards) || {};
  const lines = Object.values(cards)
    .filter((e) => e && e.card && typeof e.card.name === "string" && e.quantity > 0)
    .map((e) => e.quantity + " " + e.card.name);
  if (lines.length > 0) parts.push(head + "\\n" + lines.join("\\n"));
}
if (parts.length === 0) { alert("This Moxfield deck holds no cards."); return; }
const q = new URLSearchParams({ list: parts.join("\\n\\n"), name: (d && typeof d.name === "string") ? d.name : "" });
const u = O + "/decks#" + q.toString();
const w = window.open(u, "_blank");
if (w) w.opener = null; else location.href = u;
`;

// bookmarkletSource is the code of the bookmark for the decktome origin.
export function bookmarkletSource(origin: string): string {
  return `(async () => {\nconst O = ${JSON.stringify(origin)};\n${body}})();`;
}

// bookmarkletHref is the javascript: link that the reader drags to the
// bookmarks bar. React 19 replaces a javascript: href with an error, so
// the caller sets it on the element itself.
export function bookmarkletHref(origin: string): string {
  return `javascript:${encodeURIComponent(bookmarkletSource(origin))}`;
}

// maxFragmentBytes caps the list that a fragment can hand to the dialog.
// It matches the cap of the import (maxImportBytes).
const maxFragmentBytes = 128 * 1024;

// readImportFragment reads the list and the name that the bookmarklet put
// in the fragment. Any other fragment reads null.
export function readImportFragment(hash: string): { list: string; name: string } | null {
  const raw = hash.replace(/^#/, "");
  if (raw === "" || raw.length > maxFragmentBytes * 3) return null;
  let q: URLSearchParams;
  try {
    q = new URLSearchParams(raw);
  } catch {
    return null;
  }
  const list = q.get("list");
  if (list === null || list.trim() === "") return null;
  return { list, name: (q.get("name") ?? "").slice(0, 200) };
}
