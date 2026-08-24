---
name: design-doc-style
description: Section template and rules for docs/design-roadmap.md, modeled on connector-syncer's document-summary-roadmap.md. Load before you edit the design doc.
---

# Design-doc style skill

The owner wants the design doc in the style of `/Users/nate/Repos/connector-syncer-docs/docs/document-summary-roadmap.md`. This skill gives the template. Write in ASD-STE100 (load `ste-writing` first).

## Section template

1. **Status header.** State the doc status, what it supersedes, and the date you verified each external fact. Add a dated line for each correction pass. Never delete a refuted claim. Mark it refuted and keep it.
2. **Thesis.** One paragraph. What the system is for and why the plan has this order.
3. **Lessons learned.** Numbered. Each lesson names the event that taught it. Carry lessons from connector-syncer when they apply.
4. **System map.** A table of components, what each reads, and its sensitivity.
5. **Cost model.** What we pay, what we do not know, and which measurement will answer it.
6. **Defect and finding register.** A numbered table. Status legend: ✅ done (code merged), 🔧 planned (item listed), ⚠ constraint (binds a PR), ❓ needs owner input, parked. Findings carry evidence and dates. Findings bind to plan items ("binds PR-3").
7. **Guardrails.** Numbered invariants that every PR must keep.
8. **Roadmap.** Phases. Each entry has: an id (PR-#, M-#, I-#), a technical paragraph, a gate, and a plain-English paragraph in a block quote that starts with "*In plain English:*".
9. **Sequencing.** A strict ordered list with a single owner. Mark the gate.
10. **Open questions.** Numbered. Record the date and the answer when one arrives.

## Rules

- Every entry ends with a plain-English paragraph. The paragraph explains the item to a reader who does not know the code.
- Every external fact has a source and a date.
- Numbering continues across revisions. Never renumber.
- "One concern per PR" applies to the plan items.
- A refuted premise stays in the doc with a dated correction (lesson 7 in the model doc).
- Ids: F-# findings, Q-# quality findings, E-# constraints, PR-# code changes, M-# measurement, I-# integration, D-# owner decisions.

## Plain-English paragraph rules

- Max 25 words per sentence (STE rule 6.3).
- No code identifiers unless the reader needs them.
- Say what breaks today, what the change does, and why it is safe.
