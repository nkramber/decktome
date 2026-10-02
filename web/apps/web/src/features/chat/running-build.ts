// The installed app on a phone can end in the background, and the next
// launch opens the start address, "/". The page keeps the id of a
// session while its build runs, and a launch at "/" opens that session
// one time (D-1034, F-190).
const runningBuildKey = "decktome.runningBuild";

// runningBuildMaxMs bounds the record. A build stops at four minutes on
// the server, so an older record names a build that ended.
export const runningBuildMaxMs = 10 * 60 * 1000;

type RunningBuild = { id: string; at: number };

function read(): RunningBuild | undefined {
  try {
    const raw = localStorage.getItem(runningBuildKey);
    if (!raw) return undefined;
    const v = JSON.parse(raw) as Partial<RunningBuild>;
    return typeof v.id === "string" && v.id !== "" && typeof v.at === "number" ? { id: v.id, at: v.at } : undefined;
  } catch {
    return undefined;
  }
}

// noteRunningBuild keeps the id of a session whose build runs.
export function noteRunningBuild(id: string, now = Date.now()) {
  try {
    localStorage.setItem(runningBuildKey, JSON.stringify({ id, at: now }));
  } catch {
    // A full or blocked store keeps no record, and the launch opens "/".
  }
}

// clearRunningBuild removes the record of the session, and the record
// of another session stays.
export function clearRunningBuild(id: string) {
  if (read()?.id !== id) return;
  try {
    localStorage.removeItem(runningBuildKey);
  } catch {
    // Nothing to remove.
  }
}

// takeRunningBuild answers the address of a recent running build, or
// "" for none, and removes the record, so it opens one time.
export function takeRunningBuild(now = Date.now()): string {
  const v = read();
  if (!v) return "";
  try {
    localStorage.removeItem(runningBuildKey);
  } catch {
    // The record stays, and its age ends it.
  }
  return now - v.at <= runningBuildMaxMs ? `/session/${v.id}` : "";
}
