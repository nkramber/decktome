import jsxA11y from "eslint-plugin-jsx-a11y";
import reactHooks from "eslint-plugin-react-hooks";
import tseslint from "typescript-eslint";

// Import boundaries are lint rules, not prose (wallabee-ui pattern).
// A feature may import src/lib, src/app/components, and the sibling features
// in its allowlist. Nothing else. src/lib imports no feature and no React.
//
// | Feature    | May import from features |
// |------------|--------------------------|
// | auth       | none                     |
// | collection | none                     |
// | chat       | deck                     |
// | deck       | none                     |
// | export     | deck                     |
const features = ["auth", "collection", "chat", "deck", "export"];
const allow = { auth: [], collection: [], chat: ["deck"], deck: [], export: ["deck"] };

// A feature reaches a sibling by a relative path: ../deck/x from features/chat,
// or ../../deck/x from features/chat/components. The regex matches both.
function featureBoundary(name) {
  const blocked = features.filter((f) => f !== name && !allow[name].includes(f));
  return {
    files: [`src/features/${name}/**/*.{ts,tsx}`],
    rules: {
      "no-restricted-imports": [
        "error",
        {
          patterns: [
            {
              regex: `^(\\.\\./)+(${blocked.join("|")})(/|$)`,
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
    files: ["src/lib/**/*.{ts,tsx}"],
    rules: {
      "no-restricted-imports": [
        "error",
        {
          patterns: [
            { regex: "^(\\.\\./)+(features|app)(/|$)", message: "src/lib imports no feature and no app code." },
            { group: ["react", "react-dom", "react-router"], message: "src/lib holds no React code. The QueryClient and the zustand store are plain objects." },
          ],
        },
      ],
    },
  },
);
