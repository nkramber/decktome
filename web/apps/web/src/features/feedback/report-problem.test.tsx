import { FeedbackKind, FeedbackVerdict } from "@mtg/api-client/mtg/v1/feedback_service_pb";
import { QueryClientProvider } from "@tanstack/react-query";
import { render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { axe } from "jest-axe";
import { beforeEach, describe, expect, it, vi } from "vitest";

import { makeQueryClient } from "../../lib/query-client";
import { ReportProblem } from "./report-problem";
import { thanks } from "./thumbs";

const submitFeedback = vi.fn();
vi.mock("../../lib/api", () => ({ feedbackClient: { submitFeedback: (...a: unknown[]) => submitFeedback(...a) } }));
const notify = vi.fn();
vi.mock("../../app/components/notify", () => ({ notify: (...a: unknown[]) => notify(...a) }));

function renderReport() {
  return render(
    <QueryClientProvider client={makeQueryClient()}>
      <ReportProblem sessionId="s1" />
    </QueryClientProvider>,
  );
}

beforeEach(() => {
  submitFeedback.mockReset();
  submitFeedback.mockResolvedValue({ feedbackId: "fb1" });
  notify.mockReset();
  notify.mockResolvedValue(undefined);
});

// D-594: the thumbs read one question, one card, or one deck. A chat
// that stops between questions belongs to none of them, and session
// oUZMC0F2vHe7GGl24LIP had nothing a reader could report.
describe("ReportProblem", () => {
  it("sends a chat verdict with the session and no question", async () => {
    const user = userEvent.setup();
    renderReport();
    await user.click(screen.getByRole("button", { name: "Report a problem" }));

    await user.click(await screen.findByRole("checkbox", { name: "The chat stopped and asked nothing more." }));
    await user.click(screen.getByRole("button", { name: "Submit" }));

    expect(submitFeedback).toHaveBeenCalledTimes(1);
    const sent = submitFeedback.mock.calls[0][0].feedback;
    expect(sent.kind).toBe(FeedbackKind.CHAT);
    expect(sent.verdict).toBe(FeedbackVerdict.DOWN);
    expect(sent.sessionId).toBe("s1");
    expect(sent.reasons).toEqual(["stuck"]);
    // A chat names no question and no deck, or the API refuses it.
    expect(sent.questionId ?? "").toBe("");
    expect(sent.deckId ?? "").toBe("");
    expect(notify).toHaveBeenCalledWith("success", thanks);
  });

  it("locks after one report, so the same chat is not sent twice", async () => {
    const user = userEvent.setup();
    renderReport();
    await user.click(screen.getByRole("button", { name: "Report a problem" }));
    await user.click(await screen.findByRole("checkbox", { name: "It never built a deck." }));
    await user.click(screen.getByRole("button", { name: "Submit" }));

    expect(await screen.findByRole("button", { name: "Problem reported" })).toBeDisabled();
  });

  it("says so when the report does not send", async () => {
    submitFeedback.mockRejectedValue(new Error("the API is down"));
    const user = userEvent.setup();
    renderReport();
    await user.click(screen.getByRole("button", { name: "Report a problem" }));
    await user.click(await screen.findByRole("checkbox", { name: "Something failed or showed an error." }));
    await user.click(screen.getByRole("button", { name: "Submit" }));

    expect(notify).toHaveBeenCalledWith("error", "Could not send your report", "the API is down");
    // The dialog stays open with the draft, so the reader can send again.
    // A failed report must never read as a sent one.
    expect(screen.getByRole("button", { name: "Submit" })).toBeInTheDocument();
    expect(screen.queryByRole("button", { name: "Problem reported" })).not.toBeInTheDocument();
  });

  it("has no axe violations", async () => {
    const { container } = renderReport();
    expect(await axe(container)).toHaveNoViolations();
  });
});
