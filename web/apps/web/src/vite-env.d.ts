/// <reference types="vite/client" />

// VITE_API_BASE_URL overrides the same-origin default of src/lib/api.ts.
// Empty means same origin: the Vite proxy in dev, one host in prod.
// VITE_AUTH_EMULATOR_HOST points firebase/auth at an emulator (D-275).
// In dev it defaults to 127.0.0.1:9199. Empty outside dev means real Firebase.
interface ImportMetaEnv {
  readonly VITE_API_BASE_URL?: string;
  readonly VITE_AUTH_EMULATOR_HOST?: string;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
