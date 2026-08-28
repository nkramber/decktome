import type { Question } from "@mtg/api-client/mtg/v1/session_pb";

import { CardOption, CardOptionsError, hasCardOptions, useOptionCards } from "./card-options";

// Draft is the user's answer to one question before the submit. An option
// pick and free text exclude each other: the last one the user touched wins.
export type Draft = { optionIndex?: number; text: string };

export const emptyDraft: Draft = { text: "" };

export function draftAnswered(d: Draft | undefined): boolean {
  return d !== undefined && (d.optionIndex !== undefined || d.text.trim() !== "");
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
  const cards = useOptionCards(question);
  const withCards = hasCardOptions(question);
  const byId = new Map((cards.data?.cards ?? []).map((c) => [c.oracleId, c]));

  const button = (opt: string, i: number) => {
    const picked = draft.optionIndex === i;
    return (
      <button
        key={i}
        type="button"
        disabled={disabled}
        aria-pressed={picked}
        onClick={() => onChange(picked ? { text: "" } : { optionIndex: i, text: "" })}
        className={`rounded border px-2 py-1 text-sm disabled:bg-neutral-200 disabled:text-neutral-500 ${
          picked ? "border-neutral-900 bg-neutral-900 text-white" : "border-neutral-400 bg-white"
        }`}
      >
        {picked && <span aria-hidden="true">✓ </span>}
        {opt}
      </button>
    );
  };

  return (
    <div className="flex flex-col gap-2 rounded border border-blue-400 bg-blue-50 p-3" role="group" aria-label={`Question: ${question.text}`}>
      <p className="font-medium">{question.text}</p>
      {withCards && cards.isError && <CardOptionsError error={cards.error} />}
      {withCards ? (
        // Each card option is a tile: the art and the rules text, then
        // the pick button. A non-card option keeps its plain button.
        <div className="@container">
          <ul className="grid grid-cols-1 gap-3 @md:grid-cols-2 @3xl:grid-cols-3">
            {question.options.map((opt, i) => {
              const id = question.optionOracleIds[i] ?? "";
              return (
                <li key={i} className="flex flex-col gap-2 rounded border border-neutral-200 bg-white p-2">
                  {id ? cards.isPending ? <p role="status" className="text-sm text-neutral-600">Loading {opt}...</p> : <CardOption card={byId.get(id)} name={opt} /> : null}
                  <div>{button(opt, i)}</div>
                </li>
              );
            })}
          </ul>
        </div>
      ) : (
        question.options.length > 0 && <div className="flex flex-wrap gap-2">{question.options.map(button)}</div>
      )}
      <label className="flex flex-col gap-1 text-sm">
        <span>{question.options.length > 0 ? "Or answer in your own words" : "Your answer"}</span>
        <input
          type="text"
          value={draft.text}
          disabled={disabled}
          onChange={(e) => onChange({ text: e.target.value })}
          className="rounded border border-neutral-400 bg-white px-2 py-1 text-neutral-900"
        />
      </label>
      {question.invented && <p className="text-xs text-neutral-600">This question is not in the catalog (D-25).</p>}
    </div>
  );
}
