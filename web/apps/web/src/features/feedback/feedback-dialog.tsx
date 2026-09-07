import type { FeedbackKind } from "@mtg/api-client/mtg/v1/feedback_service_pb";
import { type FormEvent, useId, useState } from "react";

import { Button } from "../../components/ui/button";
import { Checkbox } from "../../components/ui/checkbox";
import { Dialog, DialogClose, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle } from "../../components/ui/dialog";
import { Label } from "../../components/ui/label";
import { Textarea } from "../../components/ui/textarea";
import { byteLength, maxMessageBytes } from "../../lib/limits";
import { reasonsOf } from "./reasons";

// The thumbs-down dialog (PR-27, D-557). It dims the screen, names the
// item, offers the reasons of the kind as check boxes, then an "Other"
// box. Submit opens when the user checks one reason or types in the
// box. The dialog holds no verdict of its own: the thumbs submit it.
export function FeedbackDialog({
  open,
  onOpenChange,
  kind,
  itemName,
  busy,
  onSubmit,
}: {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  kind: FeedbackKind;
  itemName: string;
  busy: boolean;
  onSubmit: (reasons: string[], text: string) => void;
}) {
  const id = useId();
  const [checked, setChecked] = useState<string[]>([]);
  const [text, setText] = useState("");
  const reasons = reasonsOf(kind);
  const tooLong = byteLength(text) > maxMessageBytes;
  const canSubmit = (checked.length > 0 || text.trim() !== "") && !tooLong && !busy;

  function toggle(key: string, on: boolean) {
    setChecked((all) => (on ? [...all.filter((k) => k !== key), key] : all.filter((k) => k !== key)));
  }

  // A closed dialog forgets its draft, so the next thumbs down on the
  // same item starts clean.
  function onChange(next: boolean) {
    if (!next) {
      setChecked([]);
      setText("");
    }
    onOpenChange(next);
  }

  function submit(e: FormEvent) {
    e.preventDefault();
    if (!canSubmit) return;
    // The reasons go in the order the dialog shows them, whatever the
    // order of the clicks.
    const ordered = reasons.map((r) => r.key).filter((k) => checked.includes(k));
    onSubmit(ordered, text.trim());
  }

  return (
    <Dialog open={open} onOpenChange={onChange}>
      <DialogContent aria-describedby={`${id}-item`}>
        <DialogHeader>
          <DialogTitle>What missed?</DialogTitle>
          <DialogDescription id={`${id}-item`}>About {itemName}.</DialogDescription>
        </DialogHeader>
        <form onSubmit={submit} className="flex flex-col gap-4">
          <fieldset className="flex flex-col gap-2">
            <legend className="mb-1 text-sm text-muted-foreground">Check what applies.</legend>
            {reasons.map((r) => (
              <div key={r.key} className="flex items-center gap-2 pointer-coarse:min-h-11">
                <Checkbox id={`${id}-${r.key}`} checked={checked.includes(r.key)} onCheckedChange={(on) => toggle(r.key, on === true)} disabled={busy} />
                <Label htmlFor={`${id}-${r.key}`} className="font-normal">
                  {r.label}
                </Label>
              </div>
            ))}
          </fieldset>
          <div className="flex flex-col gap-1.5">
            <Label htmlFor={`${id}-other`}>Other</Label>
            <Textarea id={`${id}-other`} value={text} onChange={(e) => setText(e.target.value)} disabled={busy} rows={3} placeholder="Say what went wrong in your own words." />
            {tooLong && (
              <p role="alert" className="text-sm text-danger">
                The text is too long. Shorten it.
              </p>
            )}
          </div>
          <DialogFooter>
            <DialogClose asChild>
              <Button type="button" variant="outline" disabled={busy}>
                Cancel
              </Button>
            </DialogClose>
            <Button type="submit" disabled={!canSubmit}>
              Submit
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
}
