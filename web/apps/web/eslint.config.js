import jsxA11y from "eslint-plugin-jsx-a11y";
import reactHooks from "eslint-plugin-react-hooks";
import tseslint from "typescript-eslint";

// Import boundaries are lint rules, not prose (wallabee-ui pattern).
// A feature may import src/lib, src/components/ui, src/app/components, and
// the sibling features in its allowlist. Nothing else. src/lib imports no
// feature and no React. src/components/ui holds the design-system
// primitives (D-311), so it imports src/lib and nothing of the app.
//
// | Feature    | May import from features     |
// |------------|------------------------------|
// | auth       | none                         |
// | collection | none                         |
// | chat       | deck, feedback               |
// | deck       | export, feedback             |
// | export     | none                         |
// | feedback   | none                         |
// | share      | deck, export                 |
// | workspace  | chat, deck, export, feedback |
//
// export is a leaf: the deck view mounts the export panel, and the panel
// reads the deck and the card data it is given (PR-13).
//
// feedback is a leaf too (PR-27, D-558): the thumbs and the dialog take
// the ids of the thing they judge and call the feedback client. The
// chat, the deck, and the workspace mount them.
//
// share is the public page of a share link (D-315). It renders the card
// groups of the deck feature and nothing that needs a sign-in.
//
// workspace is the one screen of a deck (D-335). It holds the deck, the
// actions the user owns, and the conversation that built it, so it is
// the only feature that may reach both chat and deck.
const features = ["auth", "collection", "chat", "deck", "export", "feedback", "share", "workspace"];
const allow = {
  auth: [],
  collection: [],
  chat: ["deck", "feedback"],
  deck: ["export", "feedback"],
  export: [],
  feedback: [],
  share: ["deck", "export"],
  workspace: ["chat", "deck", "export", "feedback"],
};

// A feature reaches a sibling by a relative path: ../deck/x from
// features/chat, ../../deck/x from features/chat/components, or
// ../../features/deck/x. The regex matches every form.
export function featureBoundary(name) {
  const blocked = features.filter((f) => f !== name && !allow[name].includes(f));
  return {
    files: [`src/features/${name}/**/*.{ts,tsx}`],
    rules: {
      "no-restricted-imports": [
        "error",
        {
          patterns: [
            {
              regex: `^(\\.\\./)+(features/)?(${blocked.join("|")})(/|$)`,
              message: `features/${name} may import only: lib, app/components, ${allow[name].join(", ") || "no sibling feature"}.`,
            },
            {
              regex: "^(\\.\\./)+app/(?!components/)",
              message: "A feature imports only src/app/components, never the router, the layout, or the providers.",
            },
          ],
        },
      ],
    },
  };
}

export default tseslint.config(
  { ignores: ["dist", "node_modules"] },
  ...tseslint.configs.recommended,
  reactHooks.configs.flat.recommended,
  jsxA11y.flatConfigs.recommended,
  ...features.map(featureBoundary),
  {
    files: ["src/components/**/*.{ts,tsx}"],
    rules: {
      "no-restricted-imports": [
        "error",
        {
          patterns: [{ regex: "^(\\.\\./)+(features|app)(/|$)", message: "A primitive imports no feature and no app code. It takes what it needs as props." }],
        },
      ],
    },
  },
  {
    files: ["src/lib/**/*.{ts,tsx}"],
    rules: {
      "no-restricted-imports": [
        "error",
        {
          patterns: [
            { regex: "^(\\.\\./)+(features|app)(/|$)", message: "src/lib imports no feature and no app code." },
            // zustand's create stays allowed: useAppStore is a plain store object
            // with a hook signature, and no feature could import it from app/.
            { group: ["react", "react-dom", "react-router"], message: "src/lib holds no React code. The QueryClient and the zustand store are plain objects." },
          ],
        },
      ],
    },
  },
);
