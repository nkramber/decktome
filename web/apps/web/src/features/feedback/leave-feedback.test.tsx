import { FeedbackKind } from "@mtg/api-client/mtg/v1/feedback_service_pb";
import { QueryClientProvider } from "@tanstack/react-query";
import { render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { axe } from "jest-axe";
import { useState } from "react";
import { MemoryRouter } from "react-router";
import { beforeEach, describe, expect, it, vi } from "vitest";

import { makeQueryClient } from "../../lib/query-client";
import { FeedbackNoteDialog, placeOf } from "./leave-feedback";
import { thanks } from "./thumbs";

const submitFeedback = vi.fn();
vi.mock("../../lib/api", () => ({ feedbackClient: { submitFeedback: (...a: unknown[]) => submitFeedback(...a) } }));
const notify = vi.fn();
vi.mock("../../app/components/notify", () => ({ notify: (...a: unknown[]) => notify(...a) }));

function Harness() {
  const [open, setOpen] = useState(true);
  return <FeedbackNoteDialog open={open} onOpenChange={setOpen} />;
}

function renderAt(path: string) {
  return render(
    <QueryClientProvider client={makeQueryClient()}>
      <MemoryRouter initialEntries={[path]}>
        <Harness />
      </MemoryRouter>
    </QueryClientProvider>,
  );
}

beforeEach(() => {
  submitFeedback.mockReset();
  submitFeedback.mockResolvedValue({ feedbackId: "fb1" });
  notify.mockReset();
  notify.mockResolvedValue(undefined);
});

// D-1078: the button of the top bar sends a note with the screen, and
// the deck of a deck screen.
describe("FeedbackNoteDialog", () => {
  it("names the screen of each path", () => {
    expect(placeOf("/session/new")).toEqual({ screen: "build" });
    expect(placeOf("/session/abc")).toEqual({ screen: "chat", sessionId: "abc" });
    expect(placeOf("/decks")).toEqual({ screen: "decks" });
    expect(placeOf("/decks/d1")).toEqual({ screen: "deck", deckId: "d1" });
    expect(placeOf("/collection")).toEqual({ screen: "collection" });
    expect(placeOf("/admin")).toEqual({ screen: "admin" });
    expect(placeOf("/elsewhere")).toEqual({ screen: "other" });
  });

  it("sends a general note with the text and the deck screen", async () => {
    const user = userEvent.setup();
    renderAt("/decks/d1");
    const send = screen.getByRole("button", { name: "Send" });
    expect(send).toBeDisabled();
    await user.type(screen.getByLabelText("Your feedback"), "  The mana curve chart is great.  ");
    await user.click(send);
    expect(submitFeedback).toHaveBeenCalledWith({
      feedback: { kind: FeedbackKind.GENERAL, text: "The mana curve chart is great.", screen: "deck", deckId: "d1" },
    });
    expect(notify).toHaveBeenCalledWith("success", thanks);
    expect(screen.queryByRole("dialog")).not.toBeInTheDocument();
  });

  it("keeps the note when the send fails", async () => {
    submitFeedback.mockRejectedValue(new Error("down"));
    const user = userEvent.setup();
    renderAt("/collection");
    await user.type(screen.getByLabelText("Your feedback"), "Hello");
    await user.click(screen.getByRole("button", { name: "Send" }));
    expect(notify).toHaveBeenCalledWith("error", "Could not send your feedback", expect.any(String));
    expect(screen.getByLabelText("Your feedback")).toHaveValue("Hello");
  });

  it("has no axe violations", async () => {
    const { baseElement } = renderAt("/session/new");
    await screen.findByRole("dialog");
    expect(await axe(baseElement)).toHaveNoViolations();
  });
});
