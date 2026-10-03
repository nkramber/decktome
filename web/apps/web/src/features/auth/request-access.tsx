import { type FormEvent, useState } from "react";

import { Button } from "../../components/ui/button";
import { Input } from "../../components/ui/input";
import { Label } from "../../components/ui/label";
import { Textarea } from "../../components/ui/textarea";
import { inviteClient } from "../../lib/api";
import { errorMessage } from "../../lib/errors";

// maxNote is the cap of the optional note, the same as the server's
// (D-1074).
export const maxNote = 500;

// requestSent is what a person reads after a request (D-1077). It reads
// the same for a new request and a repeat, so it says nothing about the
// decision of the owner.
export const requestSent = "Thanks. We will email you when your beta access is ready.";

// alreadyInvited is what a person on the invite list reads. The form
// then opens the create-account form.
export const alreadyInvited = "This email is already invited. Create your account.";

// RequestAccessForm asks for beta access (D-1074, D-1075). It takes the
// email of the sign-in form, and an optional note.
// onSent tells the page that the request went out, so the link under the
// form reads "Return to login page" (D-1079).
export function RequestAccessForm({
  email,
  onEmail,
  onInvited,
  onSent,
}: {
  email: string;
  onEmail: (email: string) => void;
  onInvited: () => void;
  onSent?: () => void;
}) {
  const [note, setNote] = useState("");
  const [busy, setBusy] = useState(false);
  const [error, setError] = useState("");
  const [sent, setSent] = useState(false);

  async function onSubmit(e: FormEvent) {
    e.preventDefault();
    setBusy(true);
    setError("");
    try {
      const res = await inviteClient.requestAccess({ email: email.trim(), note: note.trim() });
      if (res.alreadyInvited) {
        onInvited();
        return;
      }
      setSent(true);
      onSent?.();
    } catch (err) {
      setError(errorMessage(err));
    } finally {
      setBusy(false);
    }
  }

  if (sent) {
    return (
      <p role="status" className="text-sm">
        {requestSent}
      </p>
    );
  }

  return (
    <form onSubmit={onSubmit} className="flex flex-col gap-3">
      <div className="flex flex-col gap-1.5">
        <Label htmlFor="request-email">Email</Label>
        <Input id="request-email" type="email" name="email" autoComplete="email" required value={email} onChange={(e) => onEmail(e.target.value)} />
      </div>
      <div className="flex flex-col gap-1.5">
        <Label htmlFor="request-note">What do you want to build? (optional)</Label>
        <Textarea id="request-note" value={note} onChange={(e) => setNote(e.target.value)} maxLength={maxNote} rows={3} disabled={busy} />
      </div>
      <Button type="submit" disabled={busy}>
        Request beta access
      </Button>
      <div role="alert" className="min-h-6 text-sm text-danger">
        {error}
      </div>
    </form>
  );
}
