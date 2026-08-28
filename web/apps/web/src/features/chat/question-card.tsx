import type { Answer, Question } from "@mtg/api-client/mtg/v1/session_pb";
import { type FormEvent, useState } from "react";

// One open question: the options as buttons and a free-text field (ui
// plan, step 3). A button click sends option_index. Free text sends text.
export function QuestionCard({ question, disabled, onAnswer }: { question: Question; disabled: boolean; onAnswer: (a: Answer) => void }) {
  const [text, setText] = useState("");

  function submitText(e: FormEvent) {
    e.preventDefault();
    const trimmed = text.trim();
    if (!trimmed) return;
    onAnswer({ questionId: question.id, text: trimmed } as Answer);
    setText("");
  }

  return (
    <div className="flex flex-col gap-2 rounded border border-blue-400 bg-blue-50 p-3" role="group" aria-label={`Question: ${question.text}`}>
      <p className="font-medium">{question.text}</p>
      {question.options.length > 0 && (
        <div className="flex flex-wrap gap-2">
          {question.options.map((opt, i) => (
            <button
              key={i}
              type="button"
              disabled={disabled}
              onClick={() => onAnswer({ questionId: question.id, optionIndex: i, text: "" } as Answer)}
              className="rounded border border-neutral-400 bg-white px-2 py-1 text-sm disabled:bg-neutral-200 disabled:text-neutral-500"
            >
              {opt}
            </button>
          ))}
        </div>
      )}
      <form onSubmit={submitText} className="flex gap-2">
        <label className="flex grow flex-col gap-1 text-sm">
          <span>Or answer in your own words</span>
          <input
            type="text"
            value={text}
            disabled={disabled}
            onChange={(e) => setText(e.target.value)}
            className="rounded border border-neutral-400 bg-white px-2 py-1 text-neutral-900"
          />
        </label>
        <button
          type="submit"
          disabled={disabled || !text.trim()}
          className="self-end rounded border border-neutral-400 bg-white px-2 py-1 text-sm disabled:bg-neutral-200 disabled:text-neutral-500"
        >
          Answer
        </button>
      </form>
      {question.invented && <p className="text-xs text-neutral-600">This question is not in the catalog (D-25).</p>}
    </div>
  );
}
