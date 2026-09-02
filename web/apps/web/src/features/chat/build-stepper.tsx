import { BuildPhase } from "@mtg/api-client/mtg/v1/agent_service_pb";
import { CheckIcon } from "lucide-react";

import { cn } from "../../lib/cn";

// The stepper of a turn (PR-19, D-435). It lights the step the server
// named, so a build that runs for minutes shows which minute this is.
// The repair step shows only when a repair ran: most builds skip it,
// and a step that never lights reads as a failure.

const steps: { phase: BuildPhase; label: string }[] = [
  { phase: BuildPhase.READING, label: "Understand" },
  { phase: BuildPhase.SHORTLIST, label: "Shortlist" },
  { phase: BuildPhase.BUILDING, label: "Build" },
  { phase: BuildPhase.CHECKING, label: "Check" },
  { phase: BuildPhase.REPAIRING, label: "Repair" },
];

// stepOrder is the place of a phase on the line. A check after a repair
// lights the check again, so the order reads by phase and not by time.
const stepOrder: Record<number, number> = Object.fromEntries(steps.map((s, i) => [s.phase, i]));

export function BuildStepper({ phase, repaired }: { phase: BuildPhase; repaired: boolean }) {
  if (phase === BuildPhase.UNSPECIFIED) return null;
  const current = phase === BuildPhase.DONE ? steps.length : (stepOrder[phase] ?? -1);
  const shown = steps.filter((s) => s.phase !== BuildPhase.REPAIRING || repaired);
  return (
    <ol aria-label="Build steps" className="flex flex-wrap items-center gap-x-3 gap-y-1 font-mono text-[11px] tracking-wide uppercase" data-testid="build-stepper">
      {shown.map((s, i) => {
        const at = stepOrder[s.phase];
        const done = current > at;
        const active = current === at;
        return (
          <li
            key={s.phase}
            aria-current={active ? "step" : undefined}
            className={cn("flex items-center gap-1.5", done ? "text-muted-foreground" : active ? "text-primary" : "text-muted-foreground/50")}
          >
            <span aria-hidden="true" className={cn("grid size-4 place-items-center rounded-full border", active ? "border-primary" : done ? "border-muted-foreground" : "border-border")}>
              {done ? <CheckIcon className="size-2.5" /> : <span className={cn("size-1.5 rounded-full", active && "animate-pulse bg-primary")} />}
            </span>
            {s.label}
            {i < shown.length - 1 && <span aria-hidden="true" className="ml-1.5 text-border">·</span>}
          </li>
        );
      })}
    </ol>
  );
}
