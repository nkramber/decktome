// copyText writes text that a request still fetches. Safari keeps the
// activation of a click only for a clipboard call inside its handler, and
// a write after an awaited request fails there with NotAllowedError. So
// the text goes in as a promise, and the call starts before any await
// (REV-050, D-935). A browser with no ClipboardItem waits for the text,
// then writes it.
export function copyText(text: Promise<string>): Promise<void> {
  if (typeof ClipboardItem !== "undefined" && typeof navigator.clipboard?.write === "function") {
    const blob = text.then((t) => new Blob([t], { type: "text/plain" }));
    return navigator.clipboard.write([new ClipboardItem({ "text/plain": blob })]);
  }
  return text.then((t) => navigator.clipboard.writeText(t));
}
