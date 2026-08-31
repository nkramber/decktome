import type { Question } from "@mtg/api-client/mtg/v1/session_pb";
import { useId } from "react";

import { Button } from "../../components/ui/button";
import { Input } from "../../components/ui/input";
import { Label } from "../../components/ui/label";
import { cn } from "../../lib/cn";
import { CardOption, CardOptionsError, hasCardOptions, useOptionCards } from "./card-options";

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

  const button = (opt: string, i: number) => {
    const picked = draft.optionIndex === i;
    return (
      <Button
        key={i}
        variant={picked ? "default" : "outline"}
        size="sm"
        disabled={disabled}
        aria-pressed={picked}
        onClick={() => onChange(picked ? { text: "" } : { optionIndex: i, text: "" })}
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
              return (
                <li key={i} className="flex flex-col gap-2 rounded-card border border-border bg-card p-2">
                  {id && !cards.isPending ? <CardOption card={byId.get(id)} name={opt} /> : null}
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
          closes for good. */}
      <div>
        <Button
          variant={draft.declined ? "default" : "outline"}
          size="sm"
          disabled={disabled}
          aria-pressed={draft.declined === true}
          onClick={() => onChange(draft.declined ? { text: "" } : { text: "", declined: true })}
        >
          {draft.declined && <span aria-hidden="true">✓ </span>}
          You decide
        </Button>
      </div>
    </div>
  );
}
