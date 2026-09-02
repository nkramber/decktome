import type { Question } from "@mtg/api-client/mtg/v1/session_pb";
import { useId } from "react";

import { Button } from "../../components/ui/button";
import { Input } from "../../components/ui/input";
import { Label } from "../../components/ui/label";
import { cn } from "../../lib/cn";
import { CardOption, CardOptionsError, hasCardOptions, partnerIds, useOptionCards } from "./card-options";

// Draft is the user's answer to one question before the submit. An option
// pick, free text, and a decline exclude each other: the last one the
// user touched wins.
export type Draft = { optionIndex?: number; text: string; declined?: boolean };

export const emptyDraft: Draft = { text: "" };

export function draftAnswered(d: Draft | undefined): boolean {
  return d !== undefined && (d.declined === true || d.optionIndex !== undefined || d.text.trim() !== "");
}

// One open question: the options as toggle buttons and a free-text field
// (ui plan, step 3). Nothing sends here. The page sends every answer at
// once through its "Submit answers" button (D-282).
export function QuestionCard({
  question,
  draft,
  disabled,
  onChange,
}: {
  question: Question;
  draft: Draft;
  disabled: boolean;
  onChange: (d: Draft) => void;
}) {
  const answerId = useId();
  const cards = useOptionCards(question);
  const withCards = hasCardOptions(question);
  const byId = new Map((cards.data?.cards ?? []).map((c) => [c.oracleId, c]));
  // A closed question with no options gives the user nothing to pick, so
  // the text field shows as for an open one (D-295).
  const closed = question.closed && question.options.length > 0;

  // toggle picks an option, or clears it when it is the picked one. The
  // name button and the card art both call it (D-444).
  const toggle = (i: number) => {
    const picked = draft.optionIndex === i;
    onChange(picked ? { text: "" } : { optionIndex: i, text: "" });
  };
  const button = (opt: string, i: number) => {
    const picked = draft.optionIndex === i;
    return (
      <Button
        key={i}
        variant={picked ? "default" : "outline"}
        size="sm"
        disabled={disabled}
        aria-pressed={picked}
        onClick={() => toggle(i)}
        className={cn("h-auto py-1 whitespace-normal", picked && "border-accent")}
      >
        {picked && <span aria-hidden="true">✓ </span>}
        {opt}
      </Button>
    );
  };

  return (
    <div className="flex flex-col gap-2 rounded-card border border-accent/40 bg-accent/5 p-3" role="group" aria-label={`Question: ${question.text}`}>
      <p className="font-medium">{question.text}</p>
      {withCards && cards.isError && <CardOptionsError error={cards.error} />}
      {withCards && cards.isPending && (
        <p role="status" className="text-sm text-muted-foreground">
          Loading the card data...
        </p>
      )}
      {withCards ? (
        // Each card option is a tile: the art and the rules text, then
        // the pick button. A non-card option keeps its plain button.
        <div className="@container">
          <ul className="grid grid-cols-1 gap-3 @md:grid-cols-2 @3xl:grid-cols-3">
            {question.options.map((opt, i) => {
              const id = question.optionOracleIds[i] ?? "";
              // A commander pair shows both cards, because the option is
              // both of them (D-361).
              const partner = partnerIds(question)[i] ?? "";
              return (
                // A commander tile lifts under the pointer with the gold
                // light every panel throws (D-445), so a reader sees which
                // one a click takes.
                <li key={i} className="card-hover flex flex-col gap-2 rounded-card border border-border bg-card p-2 hover:border-primary/60" data-testid="card-option-tile">
                  {id && !cards.isPending ? (
                    <div className={cn("grid gap-2", partner && "grid-cols-2")}>
                      <CardOption card={byId.get(id)} name={opt} zoom={partner ? "left" : undefined} onPick={disabled ? undefined : () => toggle(i)} />
                      {partner && <CardOption card={byId.get(partner)} name={opt} zoom="right" onPick={disabled ? undefined : () => toggle(i)} />}
                    </div>
                  ) : null}
                  <div>{button(opt, i)}</div>
                </li>
              );
            })}
          </ul>
        </div>
      ) : (
        question.options.length > 0 && <div className="flex flex-wrap gap-2">{question.options.map(button)}</div>
      )}
      {!closed && (
        <div className="flex flex-col gap-1.5">
          <Label htmlFor={answerId}>{question.options.length > 0 ? "Or answer in your own words" : "Your answer"}</Label>
          <Input id={answerId} type="text" value={draft.text} disabled={disabled} onChange={(e) => onChange({ text: e.target.value })} />
        </div>
      )}
      {/* A decline hands the choice back with no value (D-353). The
          agent applies the default its corpus names, and the question
          closes for good. A declined budget stores no cap, so its
          control says what the decline does (D-404). */}
      <div>
        <Button
          variant={draft.declined ? "default" : "outline"}
          size="sm"
          disabled={disabled}
          aria-pressed={draft.declined === true}
          onClick={() => onChange(draft.declined ? { text: "" } : { text: "", declined: true })}
        >
          {draft.declined && <span aria-hidden="true">✓ </span>}
          {declineLabel(question.slot)}
        </Button>
      </div>
    </div>
  );
}

// declineLabel names what a decline of this slot does. A declined
// budget is no cap at all (D-404), and every other slot hands the
// choice to the agent.
function declineLabel(slot: string): string {
  return slot === "budget" ? "No budget" : "You decide";
}
