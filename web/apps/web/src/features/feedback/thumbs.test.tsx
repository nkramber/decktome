import { FeedbackKind, FeedbackVerdict } from "@mtg/api-client/mtg/v1/feedback_service_pb";
import { QueryClientProvider } from "@tanstack/react-query";
import { render, screen, waitFor, within } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { axe } from "jest-axe";
import { beforeEach, describe, expect, it, vi } from "vitest";

import { makeQueryClient } from "../../lib/query-client";
import { feedbackOf, Thumbs, thanks } from "./thumbs";

const submitFeedback = vi.fn();
vi.mock("../../lib/api", () => ({ feedbackClient: { submitFeedback: (...a: unknown[]) => submitFeedback(...a) } }));
const notify = vi.fn();
vi.mock("../../app/components/notify", () => ({ notify: (...a: unknown[]) => notify(...a) }));

function renderThumbs(props: Partial<Parameters<typeof Thumbs>[0]> = {}) {
  return render(
    <QueryClientProvider client={makeQueryClient()}>
      <Thumbs target={{ kind: FeedbackKind.DECK, deckId: "d1" }} itemName="the deck Elf test" label="Rate this deck" {...props} />
    </QueryClientProvider>,
  );
}

beforeEach(() => {
  submitFeedback.mockReset();
  submitFeedback.mockResolvedValue({ feedbackId: "fb1" });
  notify.mockReset();
  notify.mockResolvedValue(undefined);
});

describe("Thumbs", () => {
  it("submits a thumbs up at once, thanks the reader, and locks the pair (D-557)", async () => {
    const user = userEvent.setup();
    renderThumbs();
    const group = screen.getByRole("group", { name: "Rate this deck" });
    const up = within(group).getByRole("button", { name: "This helped" });
    const down = within(group).getByRole("button", { name: "This missed" });
    expect(up).toHaveAttribute("aria-pressed", "false");
    // A touch screen gets a 44 pixel target.
    expect(up).toHaveClass("pointer-coarse:size-11");
    expect(down).toHaveClass("pointer-coarse:size-11");
    await user.click(up);
    await waitFor(() => expect(notify).toHaveBeenCalledWith("success", thanks));
    expect(submitFeedback).toHaveBeenCalledTimes(1);
    expect(submitFeedback.mock.calls[0][0]).toEqual({ feedback: { kind: FeedbackKind.DECK, verdict: FeedbackVerdict.UP, reasons: [], text: "", deckId: "d1" } });
    expect(up).toHaveAttribute("aria-pressed", "true");
    expect(up).toBeDisabled();
    expect(down).toBeDisabled();
    expect(group).toHaveAttribute("data-verdict", "up");
    expect(screen.queryByRole("dialog")).not.toBeInTheDocument();
  });

  it("opens the dialog on a thumbs down, and Submit waits for a reason or a text", async () => {
    const user = userEvent.setup();
    renderThumbs({ target: { kind: FeedbackKind.CARD, deckId: "d1", oracleId: "o-sol" }, itemName: "Sol Ring", label: undefined });
    await user.click(screen.getByRole("button", { name: "This missed" }));
    const dialog = await screen.findByRole("dialog", { name: "What missed?" });
    expect(dialog).toHaveTextContent("About Sol Ring.");
    expect(await axe(dialog)).toHaveNoViolations();
    const submit = within(dialog).getByRole("button", { name: "Submit" });
    expect(submit).toBeDisabled();
    const boxes = within(dialog).getAllByRole("checkbox");
    expect(boxes.map((b) => b.getAttribute("aria-label") ?? b.getAttribute("id"))).toHaveLength(5);
    // The reasons of a card, in the order of the design note.
    const offTheme = within(dialog).getByRole("checkbox", { name: "This card does not fit the theme." });
    const illegal = within(dialog).getByRole("checkbox", { name: "This card breaks the format or my house rules." });
    await user.click(illegal);
    expect(submit).toBeEnabled();
    await user.click(illegal);
    expect(submit).toBeDisabled();
    await user.type(within(dialog).getByLabelText("Other"), "  Banned at my table. ");
    expect(submit).toBeEnabled();
    // The clicks come in reverse, and the keys go out in the list order.
    await user.click(illegal);
    await user.click(offTheme);
    await user.click(submit);
    await waitFor(() => expect(notify).toHaveBeenCalledWith("success", thanks));
    expect(submitFeedback.mock.calls[0][0]).toEqual({
      feedback: { kind: FeedbackKind.CARD, verdict: FeedbackVerdict.DOWN, reasons: ["off_theme", "illegal"], text: "Banned at my table.", deckId: "d1", oracleId: "o-sol" },
    });
    await waitFor(() => expect(screen.queryByRole("dialog")).not.toBeInTheDocument());
    const group = screen.getByRole("group", { name: "Rate Sol Ring" });
    expect(group).toHaveAttribute("data-verdict", "down");
    expect(within(group).getByRole("button", { name: "This missed" })).toHaveAttribute("aria-pressed", "true");
    expect(within(group).getByRole("button", { name: "This helped" })).toBeDisabled();
  });

  it("forgets the draft when the dialog closes without a submit", async () => {
    const user = userEvent.setup();
    renderThumbs();
    await user.click(screen.getByRole("button", { name: "This missed" }));
    let dialog = await screen.findByRole("dialog", { name: "What missed?" });
    await user.click(within(dialog).getByRole("checkbox", { name: "The mana base is off." }));
    await user.click(within(dialog).getByRole("button", { name: "Cancel" }));
    await waitFor(() => expect(screen.queryByRole("dialog")).not.toBeInTheDocument());
    expect(submitFeedback).not.toHaveBeenCalled();
    expect(screen.getByRole("button", { name: "This missed" })).toBeEnabled();
    await user.click(screen.getByRole("button", { name: "This missed" }));
    dialog = await screen.findByRole("dialog", { name: "What missed?" });
    expect(within(dialog).getByRole("checkbox", { name: "The mana base is off." })).toHaveAttribute("aria-checked", "false");
    expect(within(dialog).getByRole("button", { name: "Submit" })).toBeDisabled();
  });

  it("says so when the write fails, and leaves the pair live", async () => {
    const user = userEvent.setup();
    submitFeedback.mockRejectedValue(new Error("down"));
    renderThumbs();
    const up = screen.getByRole("button", { name: "This helped" });
    await user.click(up);
    await waitFor(() => expect(notify).toHaveBeenCalledWith("error", "Could not save your feedback", expect.stringContaining("down")));
    expect(up).toBeEnabled();
    expect(up).toHaveAttribute("aria-pressed", "false");
    expect(screen.getByRole("group", { name: "Rate this deck" })).not.toHaveAttribute("data-verdict");
  });

  it("names the ids of each kind and nothing else", () => {
    expect(feedbackOf({ kind: FeedbackKind.QUESTION, sessionId: "s1", questionId: "q1" }, FeedbackVerdict.UP)).toEqual({
      kind: FeedbackKind.QUESTION,
      verdict: FeedbackVerdict.UP,
      reasons: [],
      text: "",
      sessionId: "s1",
      questionId: "q1",
    });
    expect(feedbackOf({ kind: FeedbackKind.SUMMARY, deckId: "d1" }, FeedbackVerdict.DOWN, ["false_claim"], "x")).toEqual({
      kind: FeedbackKind.SUMMARY,
      verdict: FeedbackVerdict.DOWN,
      reasons: ["false_claim"],
      text: "x",
      deckId: "d1",
    });
  });
});
