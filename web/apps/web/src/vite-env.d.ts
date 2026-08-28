/// <reference types="vite/client" />

// VITE_API_BASE_URL overrides the same-origin default of src/api.ts.
// Empty means same origin: the Vite proxy in dev, one host in prod.
interface ImportMetaEnv {
  readonly VITE_API_BASE_URL?: string;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
