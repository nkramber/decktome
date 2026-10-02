import { Code, ConnectError } from "@connectrpc/connect";
import type { GetSessionResponse } from "@mtg/api-client/mtg/v1/agent_service_pb";
import { useEffect, useRef } from "react";

import { agentClient } from "../../lib/api";

// buildPollMs is the time between two reads of a session while its build
// runs on the server. A build runs for minutes, so a read each five
// seconds costs little and shows the deck soon after it lands.
export const buildPollMs = 5000;

// useBuildWatch reads the session again and again while a build of it
// runs on the server and this page holds no stream: after a reload, or
// after Stop (REV-046). It calls onEnded once, with the fresh session,
// when the server no longer reports a build. A session that is gone, or
// that belongs to another user, stops the reads.
export function useBuildWatch(sessionId: string, watch: boolean, onEnded: (res: GetSessionResponse) => void) {
  const ended = useRef(onEnded);
  useEffect(() => {
    ended.current = onEnded;
  }, [onEnded]);
  useEffect(() => {
    if (!watch || sessionId === "") return;
    let live = true;
    let reading = false;
    let timer: ReturnType<typeof setTimeout> | undefined;
    const tick = async () => {
      reading = true;
      try {
        const res = await agentClient.getSession({ sessionId });
        if (!live) return;
        if (!res.building) {
          live = false;
          ended.current(res);
          return;
        }
      } catch (err) {
        if (!live) return;
        const code = ConnectError.from(err).code;
        if (code === Code.NotFound || code === Code.PermissionDenied) {
          live = false;
          return;
        }
      } finally {
        reading = false;
      }
      timer = setTimeout(() => void tick(), buildPollMs);
    };
    timer = setTimeout(() => void tick(), buildPollMs);
    // A phone that comes back to the app reads the session at once, so a
    // deck that landed in the background shows without a wait (F-190).
    const onVisible = () => {
      if (!live || reading || document.visibilityState !== "visible") return;
      clearTimeout(timer);
      void tick();
    };
    document.addEventListener("visibilitychange", onVisible);
    return () => {
      live = false;
      clearTimeout(timer);
      document.removeEventListener("visibilitychange", onVisible);
    };
  }, [sessionId, watch]);
}
