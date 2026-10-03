// priceText shows one price in dollars. A card with no price reads
// "Price unknown", and never as free (D-1060).
export function priceText(usd: number): string {
  return usd > 0 ? `$${usd.toFixed(2)}` : "Price unknown";
}
