import { create } from "zustand";

// Dark leads (D-327). A new reader gets the dark theme, and a choice of
// light or system overrides it and persists. The resolved value drives
// the class on the root element, and dark is the class-free base.
export type ThemeChoice = "light" | "dark" | "system";
export type Theme = "light" | "dark";

export const themeStorageKey = "mtg-theme";

const choices: ThemeChoice[] = ["light", "dark", "system"];

// readChoice reads the stored choice. An absent or damaged value is "system".
export function readChoice(): ThemeChoice {
  try {
    const raw = localStorage.getItem(themeStorageKey);
    return choices.find((c) => c === raw) ?? "dark";
  } catch {
    return "dark";
  }
}

// systemPrefersDark reads the media query. A browser without it reads light.
export function systemPrefersDark(): boolean {
  return typeof window !== "undefined" && typeof window.matchMedia === "function" && window.matchMedia("(prefers-color-scheme: dark)").matches;
}

export function resolveTheme(choice: ThemeChoice): Theme {
  if (choice === "system") return systemPrefersDark() ? "dark" : "light";
  return choice;
}

// applyTheme puts the class on the root element and sets color-scheme, so
// the browser paints its own controls and scrollbars to match. Dark is
// the base, so only the light theme carries a class (D-327).
export function applyTheme(theme: Theme) {
  const root = document.documentElement;
  root.classList.toggle("light", theme === "light");
  root.style.colorScheme = theme;
}

export type ThemeState = {
  choice: ThemeChoice;
  theme: Theme;
  setChoice: (choice: ThemeChoice) => void;
  // syncSystem re-resolves after the system preference changes.
  syncSystem: () => void;
};

export const useThemeStore = create<ThemeState>()((set, get) => ({
  choice: readChoice(),
  theme: resolveTheme(readChoice()),
  setChoice: (choice) => {
    try {
      localStorage.setItem(themeStorageKey, choice);
    } catch {
      // A browser that refuses storage still changes the theme for this page.
    }
    const theme = resolveTheme(choice);
    applyTheme(theme);
    set({ choice, theme });
  },
  syncSystem: () => {
    if (get().choice !== "system") return;
    const theme = resolveTheme("system");
    applyTheme(theme);
    set({ theme });
  },
}));

// watchSystemTheme reports a system change while the choice is "system".
// It returns the function that ends the watch.
export function watchSystemTheme(): () => void {
  if (typeof window === "undefined" || typeof window.matchMedia !== "function") return () => {};
  const query = window.matchMedia("(prefers-color-scheme: dark)");
  const onChange = () => useThemeStore.getState().syncSystem();
  query.addEventListener("change", onChange);
  return () => query.removeEventListener("change", onChange);
}
