import { FeedbackKind, FeedbackVerdict } from "@mtg/api-client/mtg/v1/feedback_service_pb";
import { FlagIcon } from "lucide-react";
import { useState } from "react";

import { Button } from "../../components/ui/button";
import { notify } from "../../app/components/notify";
import { errorMessage } from "../../lib/errors";
import { FeedbackDialog } from "./feedback-dialog";
import { feedbackOf, thanks } from "./thumbs";
import { useSubmitFeedback } from "./use-feedback";

// ReportProblem is the verdict on the conversation as a whole (D-594).
// The thumbs of PR-27 read one question, one card, or one deck. A chat
// that stops between questions belongs to none of them, and session
// oUZMC0F2vHe7GGl24LIP had nothing a reader could report.
//
// It carries no thumbs up: a reader who opens this is not happy. The
// verdict is down, and the dialog asks for a reason as it does for every
// other down.
export function ReportProblem({ sessionId }: { sessionId: string }) {
  const submit = useSubmitFeedback();
  const [open, setOpen] = useState(false);
  const [done, setDone] = useState(false);

  async function send(reasons: string[], text: string) {
    try {
      await submit.mutateAsync(feedbackOf({ kind: FeedbackKind.CHAT, sessionId }, FeedbackVerdict.DOWN, reasons, text));
      setDone(true);
      setOpen(false);
      await notify("success", thanks);
    } catch (err) {
      await notify("error", "Could not send your report", errorMessage(err));
    }
  }

  return (
    <>
      <Button
        type="button"
        variant="ghost"
        size="sm"
        disabled={done || submit.isPending}
        onClick={() => setOpen(true)}
        data-testid="report-problem"
        className="self-start text-muted-foreground"
      >
        <FlagIcon aria-hidden="true" />
        {done ? "Problem reported" : "Report a problem"}
      </Button>
      <FeedbackDialog
        open={open}
        onOpenChange={setOpen}
        kind={FeedbackKind.CHAT}
        itemName="this chat"
        busy={submit.isPending}
        onSubmit={(reasons, text) => void send(reasons, text)}
      />
    </>
  );
}
