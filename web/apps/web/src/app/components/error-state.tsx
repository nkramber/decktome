import { AlertTriangleIcon, RefreshCwIcon } from "lucide-react";

import { Button } from "../../components/ui/button";

// One error state for a failed query, with a retry (D-311). The message
// goes in an alert, so a screen reader announces it.
export function ErrorState({ title = "That did not load", message, onRetry }: { title?: string; message: string; onRetry?: () => void }) {
  return (
    <div className="flex flex-col items-center gap-3 rounded-panel border border-border bg-surface px-6 py-12 text-center shadow-card">
      <span aria-hidden="true" className="grid size-11 place-items-center rounded-full bg-danger/10 text-danger">
        <AlertTriangleIcon className="size-5" />
      </span>
      <p className="font-medium">{title}</p>
      <p role="alert" className="max-w-prose wrap-anywhere text-sm text-danger">
        {message}
      </p>
      {onRetry && (
        <Button variant="outline" onClick={onRetry}>
          <RefreshCwIcon aria-hidden="true" />
          Try again
        </Button>
      )}
    </div>
  );
}
