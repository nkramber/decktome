import { ConnectError } from "@connectrpc/connect";
import { UnreadableFileSchema, UnresolvedReason, type UnresolvedRow } from "@mtg/api-client/mtg/v1/collection_pb";
import { FeedbackKind, FeedbackVerdict, type ImportPage } from "@mtg/api-client/mtg/v1/feedback_service_pb";
import { type FormEvent, useState } from "react";

import { Button } from "../../components/ui/button";
import { Input } from "../../components/ui/input";
import { Label } from "../../components/ui/label";
import { errorMessage } from "../../lib/errors";
import { useSubmitFeedback } from "./use-feedback";

// isUnreadable reports whether an error of the upload or the import names
// a file the app could not read (D-887). The server marks such an error
// with the UnreadableFile detail. A file that is too large, or a bad
// commander pick, carries no detail and shows its own message.
export function isUnreadable(err: unknown): boolean {
  return ConnectError.from(err).findDetails(UnreadableFileSchema).length > 0;
}

// parseFaults counts the rows the parser could not read. A row of an
// unknown card, of another language, or of a card no format plays is no
// parse fault, and it shows no form (D-887).
export function parseFaults(rows: readonly UnresolvedRow[]): number {
  return rows.filter((r) => r.reason === UnresolvedReason.BAD_ROW || r.reason === UnresolvedReason.UNKNOWN_VALUE).length;
}

type ReportImportProps = {
  page: ImportPage;
  // file is what the page could not read. The form sends it again, and
  // the server reads it once more and keeps the fault, never the file
  // (D-884, D-885).
  file: Blob;
  // badRows counts the rows that did not read. Zero means the whole file
  // did not read.
  badRows: number;
};

// ReportImport is the short form of D-882. It says that an error
// occurred, and it asks which app or site the file came from, in the
// user's own words (D-883). The answer files a thumbs down of kind
// IMPORT, so the feedback review can fix the parser. It names no
// service itself (D-889).
export function ReportImport({ page, file, badRows }: ReportImportProps) {
  const submit = useSubmitFeedback();
  const [service, setService] = useState("");
  const [done, setDone] = useState(false);

  async function onSubmit(e: FormEvent) {
    e.preventDefault();
    if (submit.isPending || done) return;
    const importContent = new Uint8Array(await file.arrayBuffer());
    await submit.mutateAsync({
      kind: FeedbackKind.IMPORT,
      verdict: FeedbackVerdict.DOWN,
      text: service.trim(),
      importPage: page,
      importContent,
    });
    setDone(true);
  }

  const what = badRows > 0 ? `The app could not read ${badRows} ${badRows === 1 ? "row" : "rows"} of this file.` : "Something went wrong. The app could not read this file.";

  return (
    <div className="flex flex-col gap-2 rounded-card border border-border p-3 text-sm" data-testid="report-import">
      <p role="alert" className="text-danger">
        {what}
      </p>
      {done ? (
        <p role="status">Thank you! Your report will be reviewed so that we can resolve the issue.</p>
      ) : (
        <form onSubmit={(e) => void onSubmit(e).catch(() => {})} className="flex flex-col gap-2">
          <Label htmlFor="import-service">Which app or site did the file come from?</Label>
          <Input id="import-service" name="service" type="text" maxLength={200} value={service} onChange={(e) => setService(e.target.value)} />
          <Button type="submit" size="sm" className="self-start" disabled={submit.isPending}>
            {submit.isPending ? "Sending..." : "Send a report"}
          </Button>
          {submit.isError && (
            <p role="alert" className="text-danger">
              Could not send the report: {errorMessage(submit.error)}
            </p>
          )}
        </form>
      )}
    </div>
  );
}
