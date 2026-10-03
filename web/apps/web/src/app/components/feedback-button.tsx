import { MessageSquarePlusIcon } from "lucide-react";
import { useState } from "react";

import { Button } from "../../components/ui/button";
import { feedbackNoteChunk } from "../chunks";

const Note = feedbackNoteChunk.Mount;

// FeedbackButton is the "Leave feedback" button of the row under the top
// bar (D-1078). Its label shows on every screen width.
// The dialog loads on the first press, so the shell stays small.
export function FeedbackButton() {
  const [open, setOpen] = useState(false);
  return (
    <>
      <Button
        type="button"
        variant="outline"
        size="sm"
        onClick={() => {
          feedbackNoteChunk.preload();
          setOpen(true);
        }}
      >
        <MessageSquarePlusIcon aria-hidden="true" />
        Leave feedback
      </Button>
      <Note open={open} onOpenChange={setOpen} />
    </>
  );
}
