import type { MessageInitShape } from "@bufbuild/protobuf";
import type { FeedbackSchema } from "@mtg/api-client/mtg/v1/feedback_service_pb";
import { useMutation } from "@tanstack/react-query";

import { feedbackClient } from "../../lib/api";

// One verdict goes to the API as one call (PR-27). The caller shows the
// toast, so a failure and a success read the same way everywhere.
export function useSubmitFeedback() {
  return useMutation({
    mutationFn: (feedback: MessageInitShape<typeof FeedbackSchema>) => feedbackClient.submitFeedback({ feedback }),
  });
}
