import { FeedbackKind, type FeedbackSchema, FeedbackVerdict } from "@mtg/api-client/mtg/v1/feedback_service_pb";
import type { MessageInitShape } from "@bufbuild/protobuf";
import { ThumbsDownIcon, ThumbsUpIcon } from "lucide-react";
import { useState } from "react";

import { notify } from "../../app/components/notify";
import { Button } from "../../components/ui/button";
import { cn } from "../../lib/cn";
import { errorMessage } from "../../lib/errors";
import { FeedbackDialog } from "./feedback-dialog";
import { useSubmitFeedback } from "./use-feedback";

// Target names the thing a pair of thumbs judges (PR-27). The ids of a
// kind are the ones the API checks against the caller's own objects.
export type Target =
  | { kind: FeedbackKind.QUESTION; sessionId: string; questionId: string }
  | { kind: FeedbackKind.SUMMARY; deckId: string }
  | { kind: FeedbackKind.CARD; deckId: string; oracleId: string }
  | { kind: FeedbackKind.DECK; deckId: string };

export function feedbackOf(target: Target, verdict: FeedbackVerdict, reasons: string[] = [], text = ""): MessageInitShape<typeof FeedbackSchema> {
  const ids =
    target.kind === FeedbackKind.QUESTION
      ? { sessionId: target.sessionId, questionId: target.questionId }
      : target.kind === FeedbackKind.CARD
        ? { deckId: target.deckId, oracleId: target.oracleId }
        : { deckId: target.deckId };
  return { kind: target.kind, verdict, reasons, text, ...ids };
}

export const thanks = "Thank you for your feedback!";

// Thumbs is the pair of buttons (D-557): "This helped" submits at once,
// and "This missed" opens the dialog. After a verdict the pair shows the
// chosen state and takes no second click on this page. Each button is
// 44 pixels or more on a touch screen. label is the visible caption of
// the pair, name is the accessible name of the pair when the caption is
// not it, and itemName is what the dialog names.
export function Thumbs({
  target,
  itemName,
  label,
  name,
  className,
  buttonClassName,
}: {
  target: Target;
  itemName: string;
  label?: string;
  name?: string;
  className?: string;
  buttonClassName?: string;
}) {
  const submit = useSubmitFeedback();
  const [verdict, setVerdict] = useState<FeedbackVerdict | undefined>();
  const [open, setOpen] = useState(false);
  const done = verdict !== undefined;

  async function rate(next: FeedbackVerdict, reasons: string[] = [], text = "") {
    try {
      await submit.mutateAsync(feedbackOf(target, next, reasons, text));
      setVerdict(next);
      setOpen(false);
      await notify("success", thanks);
    } catch (err) {
      await notify("error", "Could not save your feedback", errorMessage(err));
    }
  }

  const button = (which: FeedbackVerdict, name: string, Icon: typeof ThumbsUpIcon, tone: string) => {
    const chosen = verdict === which;
    return (
      <Button
        type="button"
        variant="ghost"
        size="icon"
        aria-label={name}
        title={name}
        aria-pressed={chosen}
        disabled={done || submit.isPending}
        onClick={() => (which === FeedbackVerdict.UP ? void rate(which) : setOpen(true))}
        className={cn("pointer-coarse:size-11", chosen && `${tone} disabled:opacity-100`, buttonClassName)}
      >
        <Icon aria-hidden="true" />
      </Button>
    );
  };

  return (
    <div role="group" aria-label={name ?? label ?? `Rate ${itemName}`} data-testid="thumbs" data-verdict={verdict === FeedbackVerdict.UP ? "up" : verdict === FeedbackVerdict.DOWN ? "down" : undefined} className={cn("flex items-center gap-1", className)}>
      {label && <span className="mr-1 text-xs text-muted-foreground">{label}</span>}
      {button(FeedbackVerdict.UP, "This helped", ThumbsUpIcon, "text-success")}
      {button(FeedbackVerdict.DOWN, "This missed", ThumbsDownIcon, "text-danger")}
      <FeedbackDialog open={open} onOpenChange={setOpen} kind={target.kind} itemName={itemName} busy={submit.isPending} onSubmit={(reasons, text) => void rate(FeedbackVerdict.DOWN, reasons, text)} />
    </div>
  );
}
