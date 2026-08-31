import { AlertTriangleIcon, RefreshCwIcon } from "lucide-react";

import { Button } from "../../components/ui/button";

// One error state for a failed query, with a retry (D-311). The message
// goes in an alert, so a screen reader announces it.
export function ErrorState({ title = "That did not load", message, onRetry }: { title?: string; message: string; onRetry?: () => void }) {
  return (
    <div className="flex flex-col items-center gap-3 rounded-card border border-border bg-card px-6 py-12 text-center shadow-card">
      <AlertTriangleIcon className="size-8 text-danger" aria-hidden="true" />
      <p className="font-display text-base font-semibold">{title}</p>
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
