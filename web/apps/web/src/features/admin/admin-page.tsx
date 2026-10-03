import { timestampDate } from "@bufbuild/protobuf/wkt";
import { type AccessRequest, AccessStatus } from "@mtg/api-client/mtg/v1/admin_service_pb";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { InboxIcon } from "lucide-react";
import { useState } from "react";

import { EmptyState } from "../../app/components/empty-state";
import { ErrorState } from "../../app/components/error-state";
import { notify } from "../../app/components/notify";
import { PageHeader } from "../../app/components/page-header";
import { Button } from "../../components/ui/button";
import { adminClient } from "../../lib/api";
import { errorMessage } from "../../lib/errors";

const tabs = [
  { status: AccessStatus.PENDING, label: "Pending" },
  { status: AccessStatus.APPROVED, label: "Approved" },
  { status: AccessStatus.DISMISSED, label: "Dismissed" },
] as const;

// approvedSent and approvedUnsent are the toasts of an approval
// (D-1077). A failed send keeps the approval.
export const approvedSent = "Approved. The approval email went out.";
export const approvedUnsent = "Approved, but the email did not send. Write to the person yourself.";

// AdminPage is the admin screen of the owner (D-1076). It lists the
// requests for beta access and decides each one. The API refuses every
// call without the admin claim, so a reader who opens /admin without it
// reads the refusal.
export function AdminPage() {
  const [status, setStatus] = useState<AccessStatus>(AccessStatus.PENDING);
  const queryClient = useQueryClient();
  const list = useQuery({
    queryKey: ["access-requests", status],
    queryFn: async () => (await adminClient.listAccessRequests({ status })).requests,
  });
  const refresh = () => queryClient.invalidateQueries({ queryKey: ["access-requests"] });
  const approve = useMutation({
    mutationFn: (email: string) => adminClient.approveAccessRequest({ email }),
    onSuccess: async (res) => {
      await notify("success", res.emailSent ? approvedSent : approvedUnsent);
      await refresh();
    },
    onError: (err) => void notify("error", "Could not approve the request", errorMessage(err)),
  });
  const dismiss = useMutation({
    mutationFn: (email: string) => adminClient.dismissAccessRequest({ email }),
    onSuccess: async () => {
      await notify("success", "Dismissed.");
      await refresh();
    },
    onError: (err) => void notify("error", "Could not dismiss the request", errorMessage(err)),
  });
  const busy = approve.isPending || dismiss.isPending;

  return (
    <div className="mx-auto flex w-full max-w-[60rem] flex-col gap-6 p-4 md:p-6">
      <PageHeader title="Access requests" description="Approve a request to put the email on the invite list and send the approval email." />
      <div role="group" aria-label="Status" className="flex flex-wrap gap-2">
        {tabs.map((t) => (
          <Button key={t.status} type="button" size="sm" variant={status === t.status ? "default" : "outline"} aria-pressed={status === t.status} onClick={() => setStatus(t.status)}>
            {t.label}
          </Button>
        ))}
      </div>
      {list.isPending ? (
        <p className="text-sm text-muted-foreground">Loading the requests.</p>
      ) : list.isError ? (
        <ErrorState message={errorMessage(list.error)} onRetry={() => void list.refetch()} />
      ) : list.data.length === 0 ? (
        <EmptyState icon={InboxIcon} title="No requests here" />
      ) : (
        <ul className="flex flex-col gap-3">
          {list.data.map((r) => (
            <RequestRow key={r.email} request={r} busy={busy} onApprove={() => approve.mutate(r.email)} onDismiss={() => dismiss.mutate(r.email)} />
          ))}
        </ul>
      )}
    </div>
  );
}

function RequestRow({ request, busy, onApprove, onDismiss }: { request: AccessRequest; busy: boolean; onApprove: () => void; onDismiss: () => void }) {
  const created = request.createdAt ? timestampDate(request.createdAt).toLocaleString() : "";
  const pending = request.status === AccessStatus.PENDING;
  return (
    <li className="flex flex-col gap-2 rounded-card border border-border bg-card p-4 shadow-card">
      <div className="flex flex-wrap items-baseline justify-between gap-2">
        <span className="font-medium break-all">{request.email}</span>
        <span className="text-xs text-muted-foreground">
          {created}
          {request.count > 1 ? `, ${request.count} requests` : ""}
        </span>
      </div>
      {request.note && <p className="text-sm whitespace-pre-wrap text-muted-foreground">{request.note}</p>}
      {/* An approved request is done. A dismissed one can still be
          approved. */}
      {request.status !== AccessStatus.APPROVED && (
        <div className="flex flex-wrap gap-2">
          <Button type="button" size="sm" disabled={busy} onClick={onApprove}>
            Approve
          </Button>
          {pending && (
            <Button type="button" size="sm" variant="outline" disabled={busy} onClick={onDismiss}>
              Dismiss
            </Button>
          )}
        </div>
      )}
    </li>
  );
}
