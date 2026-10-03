import { FeedbackKind } from "@mtg/api-client/mtg/v1/feedback_service_pb";
import { type FormEvent, useId, useState } from "react";
import { matchPath, useLocation } from "react-router";

import { notify } from "../../app/components/notify";
import { Button } from "../../components/ui/button";
import { Dialog, DialogClose, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle } from "../../components/ui/dialog";
import { Label } from "../../components/ui/label";
import { Textarea } from "../../components/ui/textarea";
import { errorMessage } from "../../lib/errors";
import { byteLength, maxMessageBytes } from "../../lib/limits";
import { thanks } from "./thumbs";
import { useSubmitFeedback } from "./use-feedback";

// phoneTop moves the dialog to the top of a narrow screen. The keyboard
// of a phone takes the lower half, and a centered dialog sat under it.
export const phoneTop = "max-md:top-[max(1rem,env(safe-area-inset-top))] max-md:translate-y-0";

export type Place = { screen: string; sessionId?: string; deckId?: string };

// placeOf names the screen of a path, with the session of a chat and the
// deck of a deck screen (D-1078). The server knows the same names.
export function placeOf(pathname: string): Place {
  if (pathname === "/session/new") return { screen: "build" };
  const chat = matchPath("/session/:id", pathname);
  if (chat?.params.id) return { screen: "chat", sessionId: chat.params.id };
  if (pathname === "/decks") return { screen: "decks" };
  const deck = matchPath("/decks/:id", pathname);
  if (deck?.params.id) return { screen: "deck", deckId: deck.params.id };
  if (pathname === "/collection") return { screen: "collection" };
  if (pathname === "/admin") return { screen: "admin" };
  return { screen: "other" };
}

// FeedbackNoteDialog sends a note to the owner (D-1078). A note carries
// no thumbs, and it needs a text. The button of the top bar opens it,
// and the shell loads it on demand.
export function FeedbackNoteDialog({ open, onOpenChange }: { open: boolean; onOpenChange: (open: boolean) => void }) {
  const id = useId();
  const { pathname } = useLocation();
  const submit = useSubmitFeedback();
  const [text, setText] = useState("");
  const tooLong = byteLength(text) > maxMessageBytes;
  const canSubmit = text.trim() !== "" && !tooLong && !submit.isPending;

  async function send(e: FormEvent) {
    e.preventDefault();
    if (!canSubmit) return;
    try {
      await submit.mutateAsync({ kind: FeedbackKind.GENERAL, text: text.trim(), ...placeOf(pathname) });
      setText("");
      onOpenChange(false);
      await notify("success", thanks);
    } catch (err) {
      await notify("error", "Could not send your feedback", errorMessage(err));
    }
  }

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
        {/* On a phone the dialog sits at the top of the screen, so the
            keyboard opens under it and never covers it (D-1079). */}
        <DialogContent aria-describedby={`${id}-about`} className={phoneTop}>
          <DialogHeader>
            <DialogTitle>Leave feedback</DialogTitle>
            <DialogDescription id={`${id}-about`}>Tell us what works, what does not, or what you want next.</DialogDescription>
          </DialogHeader>
          <form onSubmit={(e) => void send(e)} className="flex flex-col gap-4">
            <div className="flex flex-col gap-1.5">
              <Label htmlFor={`${id}-text`}>Your feedback</Label>
              <Textarea id={`${id}-text`} value={text} onChange={(e) => setText(e.target.value)} disabled={submit.isPending} rows={5} />
              {tooLong && (
                <p role="alert" className="text-sm text-danger">
                  The text is too long. Shorten it.
                </p>
              )}
            </div>
            <DialogFooter>
              <DialogClose asChild>
                <Button type="button" variant="outline">
                  Cancel
                </Button>
              </DialogClose>
              <Button type="submit" disabled={!canSubmit}>
                Send
              </Button>
            </DialogFooter>
          </form>
        </DialogContent>
    </Dialog>
  );
}
