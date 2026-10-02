import { beforeEach, describe, expect, it } from "vitest";

import { clearRunningBuild, noteRunningBuild, runningBuildMaxMs, takeRunningBuild } from "./running-build";

// D-1034: a launch at "/" opens the session of a build that ran.
describe("the running build record", () => {
  beforeEach(() => localStorage.clear());

  it("opens a recent build one time", () => {
    noteRunningBuild("s1", 1_000);
    expect(takeRunningBuild(2_000)).toBe("/session/s1");
    expect(takeRunningBuild(3_000)).toBe("");
  });

  it("opens no build past the age limit", () => {
    noteRunningBuild("s1", 1_000);
    expect(takeRunningBuild(1_000 + runningBuildMaxMs + 1)).toBe("");
  });

  it("a clear of another session keeps the record", () => {
    noteRunningBuild("s1", 1_000);
    clearRunningBuild("s2");
    expect(takeRunningBuild(2_000)).toBe("/session/s1");
    noteRunningBuild("s1", 1_000);
    clearRunningBuild("s1");
    expect(takeRunningBuild(2_000)).toBe("");
  });

  it("reads a bad record as none", () => {
    localStorage.setItem("decktome.runningBuild", "{not json");
    expect(takeRunningBuild()).toBe("");
    localStorage.setItem("decktome.runningBuild", JSON.stringify({ id: 5, at: "x" }));
    expect(takeRunningBuild()).toBe("");
  });
});
