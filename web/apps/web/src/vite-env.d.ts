/// <reference types="vite/client" />

// VITE_API_BASE_URL overrides the same-origin default of src/lib/api.ts.
// Empty means same origin: the Vite proxy in dev, one host in prod.
// VITE_AUTH_EMULATOR_HOST points firebase/auth at an emulator (D-275).
// In dev it defaults to 127.0.0.1:9199. Empty outside dev means real Firebase.
// The four VITE_FIREBASE_ variables carry the web configuration of the
// deployed Firebase project (PR-22). Unset, the app talks to the emulator
// project mtg-local.
interface ImportMetaEnv {
  readonly VITE_API_BASE_URL?: string;
  readonly VITE_AUTH_EMULATOR_HOST?: string;
  readonly VITE_FIREBASE_API_KEY?: string;
  readonly VITE_FIREBASE_AUTH_DOMAIN?: string;
  readonly VITE_FIREBASE_PROJECT_ID?: string;
  readonly VITE_FIREBASE_APP_ID?: string;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
