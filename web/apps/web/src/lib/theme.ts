import { create } from "zustand";

// The theme follows the system by default, and a choice overrides it and
// persists (D-311). The resolved value is the class on the root element.
export type ThemeChoice = "light" | "dark" | "system";
export type Theme = "light" | "dark";

export const themeStorageKey = "mtg-theme";

const choices: ThemeChoice[] = ["light", "dark", "system"];

// readChoice reads the stored choice. An absent or damaged value is "system".
export function readChoice(): ThemeChoice {
  try {
    const raw = localStorage.getItem(themeStorageKey);
    return choices.find((c) => c === raw) ?? "system";
  } catch {
    return "system";
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
// the browser paints its own controls and scrollbars to match.
export function applyTheme(theme: Theme) {
  const root = document.documentElement;
  root.classList.toggle("dark", theme === "dark");
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
