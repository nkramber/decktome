// The axe matcher's types augment Jest's global matchers alone, and
// vitest 5 no longer inherits them. This puts the matcher on vitest's own
// assertion, as the jest-dom types do for theirs.
import "vitest";

declare module "vitest" {
  // The merged declaration keeps vitest's own parameter names, so T stays.
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  interface Assertion<R, T> {
    toHaveNoViolations(): R;
  }
  interface AsymmetricMatchersContaining {
    toHaveNoViolations(): void;
  }
}
