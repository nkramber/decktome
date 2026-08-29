import { useEffect, useState } from "react";

// The shell renders one navigation, not two (D-311). Two landmarks with
// the same name fail the axe landmark-unique rule, and a reader hears the
// same three entries twice. The query matches Tailwind's md breakpoint.
const phoneQuery = "(max-width: 767px)";

function matches(): boolean {
  return typeof window !== "undefined" && typeof window.matchMedia === "function" && window.matchMedia(phoneQuery).matches;
}

export function useIsPhone(): boolean {
  const [phone, setPhone] = useState(matches);
  useEffect(() => {
    if (typeof window === "undefined" || typeof window.matchMedia !== "function") return;
    const query = window.matchMedia(phoneQuery);
    const onChange = () => setPhone(query.matches);
    query.addEventListener("change", onChange);
    onChange();
    return () => query.removeEventListener("change", onChange);
  }, []);
  return phone;
}
