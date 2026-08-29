import { type ClassValue, clsx } from "clsx";
import { twMerge } from "tailwind-merge";

// cn joins class names and lets a later Tailwind class win over an earlier
// one of the same kind. Every primitive takes a className this way (D-311).
export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}
