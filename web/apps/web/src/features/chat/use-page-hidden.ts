import { useEffect } from "react";

import { pageClient } from "../../lib/api";

// usePageHidden tells the API when the page of a session goes to the
// background, and when it comes back, while a turn of it runs. A phone
// can suspend the page and keep its stream open, so the API can not see
// the page leave. A build that stores its deck while the page is hidden
// sends the push of a finished build (D-1033, F-192).
export function usePageHidden(sessionId: string, active: boolean) {
  useEffect(() => {
    if (!active || sessionId === "") return;
    // A return to view goes out only after a report of a leave.
    let reported = false;
    const onChange = () => {
      const hidden = document.visibilityState === "hidden";
      if (hidden === reported) return;
      reported = hidden;
      // A lost report sends no push, and the stream still shows the deck.
      void pageClient.setPageHidden({ sessionId, hidden }).catch(() => {});
    };
    document.addEventListener("visibilitychange", onChange);
    return () => document.removeEventListener("visibilitychange", onChange);
  }, [sessionId, active]);
}
