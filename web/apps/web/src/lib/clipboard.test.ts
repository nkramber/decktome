import { afterEach, describe, expect, it, vi } from "vitest";

import { copyText } from "./clipboard";

// REV-050, D-935: Safari refuses a clipboard write that comes after an
// awaited request, so the write must start before the text arrives.
describe("copyText", () => {
  afterEach(() => {
    vi.unstubAllGlobals();
  });

  it("starts the write before the text arrives, where ClipboardItem exists", async () => {
    const items: { types: Record<string, Promise<Blob>> }[] = [];
    vi.stubGlobal(
      "ClipboardItem",
      class {
        constructor(public types: Record<string, Promise<Blob>>) {
          items.push(this);
        }
      },
    );
    const write = vi.fn().mockResolvedValue(undefined);
    const writeText = vi.fn();
    Object.defineProperty(navigator, "clipboard", { value: { write, writeText }, configurable: true });
    let answer: (t: string) => void = () => {};
    const text = new Promise<string>((r) => (answer = r));
    const done = copyText(text);
    expect(write).toHaveBeenCalledTimes(1);
    answer("1 Llanowar Elves\n");
    await done;
    expect(await (await items[0].types["text/plain"]).text()).toBe("1 Llanowar Elves\n");
    expect(writeText).not.toHaveBeenCalled();
  });

  it("waits for the text and writes it, where ClipboardItem does not exist", async () => {
    vi.stubGlobal("ClipboardItem", undefined);
    const writeText = vi.fn().mockResolvedValue(undefined);
    Object.defineProperty(navigator, "clipboard", { value: { writeText }, configurable: true });
    await copyText(Promise.resolve("Deck\n"));
    expect(writeText).toHaveBeenCalledWith("Deck\n");
  });

  it("rejects when the text fails", async () => {
    vi.stubGlobal("ClipboardItem", undefined);
    Object.defineProperty(navigator, "clipboard", { value: { writeText: vi.fn() }, configurable: true });
    await expect(copyText(Promise.reject(new Error("export failed")))).rejects.toThrow("export failed");
  });
});
