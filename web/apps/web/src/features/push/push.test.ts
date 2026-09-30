import { afterEach, beforeEach, describe, expect, it, vi } from "vitest";

const getDevice = vi.fn();
const registerDevice = vi.fn();
const unregisterDevice = vi.fn();

vi.mock("../../lib/api", () => ({
  pushClient: {
    getDevice: (req: unknown) => getDevice(req),
    registerDevice: (req: unknown) => registerDevice(req),
    unregisterDevice: (req: unknown) => unregisterDevice(req),
  },
}));
vi.mock("../../lib/firebase", () => ({ loadAuth: () => Promise.resolve({ app: {} }) }));

const messaging = vi.hoisted(() => ({
  isSupported: vi.fn(),
  getMessaging: vi.fn(() => ({})),
  register: vi.fn(),
  unregister: vi.fn(),
  onRegistered: vi.fn(),
}));
vi.mock("firebase/messaging", () => messaging);
vi.mock("firebase/installations", () => ({ getInstallations: () => ({}), getId: () => Promise.resolve("fid-1") }));

import { disablePush, enablePush, pushFlagKey, pushReleaseKey, pushState, refreshPush, releasePushOnSignOut } from "./push";

const worker = {} as ServiceWorkerRegistration;
let permission: NotificationPermission;
const requestPermission = vi.fn();

beforeEach(() => {
  vi.stubEnv("VITE_FIREBASE_MESSAGING_SENDER_ID", "492774632746");
  permission = "default";
  requestPermission.mockReset();
  requestPermission.mockImplementation(() => Promise.resolve(permission));
  vi.stubGlobal("Notification", {
    get permission() {
      return permission;
    },
    requestPermission,
  });
  Object.defineProperty(navigator, "serviceWorker", { configurable: true, value: { ready: Promise.resolve(worker) } });
  messaging.isSupported.mockResolvedValue(true);
  let handler: ((fid: string) => void) | undefined;
  messaging.onRegistered.mockImplementation((_m: unknown, h: (fid: string) => void) => {
    handler = h;
    return () => {};
  });
  messaging.register.mockImplementation(() => {
    handler?.("fid-1");
    return Promise.resolve();
  });
  messaging.unregister.mockResolvedValue(undefined);
  getDevice.mockReset();
  registerDevice.mockReset().mockResolvedValue({});
  unregisterDevice.mockReset().mockResolvedValue({});
  localStorage.clear();
});

afterEach(() => {
  vi.unstubAllEnvs();
  vi.unstubAllGlobals();
});

describe("pushState", () => {
  it("reads unavailable in a build with no sender id", async () => {
    vi.stubEnv("VITE_FIREBASE_MESSAGING_SENDER_ID", "");
    expect(await pushState()).toBe("unavailable");
  });

  it("reads unavailable where the browser takes no push", async () => {
    messaging.isSupported.mockResolvedValue(false);
    expect(await pushState()).toBe("unavailable");
  });

  it("reads blocked when the browser refused the permission", async () => {
    permission = "denied";
    expect(await pushState()).toBe("blocked");
  });

  it("reads off with no flag, and asks the API nothing", async () => {
    permission = "granted";
    expect(await pushState()).toBe("off");
    expect(getDevice).not.toHaveBeenCalled();
  });

  it("reads on when the API holds the device", async () => {
    permission = "granted";
    localStorage.setItem(pushFlagKey, "1");
    getDevice.mockResolvedValue({ registered: true });
    expect(await pushState()).toBe("on");
    expect(getDevice).toHaveBeenCalledWith({ installationId: "fid-1" });
  });

  it("reads off and clears the flag when the API removed the device", async () => {
    permission = "granted";
    localStorage.setItem(pushFlagKey, "1");
    getDevice.mockResolvedValue({ registered: false });
    expect(await pushState()).toBe("off");
    expect(localStorage.getItem(pushFlagKey)).toBeNull();
  });
});

describe("enablePush", () => {
  it("asks inside the gesture, registers, and stores the device", async () => {
    requestPermission.mockImplementation(() => {
      permission = "granted";
      return Promise.resolve(permission);
    });
    const done = enablePush();
    // Safari asks only inside the gesture, so the ask comes before any
    // await of the call.
    expect(requestPermission).toHaveBeenCalledTimes(1);
    expect(await done).toBe("on");
    expect(messaging.register).toHaveBeenCalledWith(expect.anything(), { serviceWorkerRegistration: worker });
    expect(registerDevice).toHaveBeenCalledWith({ installationId: "fid-1" });
    expect(localStorage.getItem(pushFlagKey)).toBe("1");
  });

  it("stores nothing when the reader refuses the permission", async () => {
    requestPermission.mockImplementation(() => {
      permission = "denied";
      return Promise.resolve(permission);
    });
    expect(await enablePush()).toBe("blocked");
    expect(messaging.register).not.toHaveBeenCalled();
    expect(registerDevice).not.toHaveBeenCalled();
  });

  it("stores nothing when the reader dismisses the ask", async () => {
    expect(await enablePush()).toBe("off");
    expect(registerDevice).not.toHaveBeenCalled();
  });
});

describe("disablePush", () => {
  it("removes the device from the API before the registration ends", async () => {
    localStorage.setItem(pushFlagKey, "1");
    const order: string[] = [];
    unregisterDevice.mockImplementation(() => {
      order.push("api");
      return Promise.resolve({});
    });
    messaging.unregister.mockImplementation(() => {
      order.push("fcm");
      return Promise.resolve();
    });
    expect(await disablePush()).toBe("off");
    expect(unregisterDevice).toHaveBeenCalledWith({ installationId: "fid-1" });
    expect(order).toEqual(["api", "fcm"]);
    expect(localStorage.getItem(pushFlagKey)).toBeNull();
  });
});

describe("releasePushOnSignOut", () => {
  it("does nothing on a browser that never turned push on", async () => {
    await releasePushOnSignOut();
    expect(unregisterDevice).not.toHaveBeenCalled();
  });

  it("removes the device, and a failure never stops the sign-out", async () => {
    localStorage.setItem(pushFlagKey, "1");
    unregisterDevice.mockRejectedValue(new Error("offline"));
    await expect(releasePushOnSignOut()).resolves.toBeUndefined();
    expect(unregisterDevice).toHaveBeenCalled();
    expect(localStorage.getItem(pushFlagKey)).toBeNull();
  });

  // The review finding of #258: a failed API call must still end the
  // registration, or the next account on the browser gets the pushes of
  // the last one.
  it("ends the registration when the API call fails", async () => {
    localStorage.setItem(pushFlagKey, "1");
    unregisterDevice.mockRejectedValue(new Error("token expired"));
    await releasePushOnSignOut();
    expect(messaging.unregister).toHaveBeenCalledTimes(1);
    expect(localStorage.getItem(pushReleaseKey)).toBeNull();
  });

  it("leaves a marker when the registration can not end, and the next start ends it", async () => {
    localStorage.setItem(pushFlagKey, "1");
    unregisterDevice.mockRejectedValue(new Error("offline"));
    messaging.unregister.mockRejectedValueOnce(new Error("offline"));
    await releasePushOnSignOut();
    expect(localStorage.getItem(pushReleaseKey)).toBe("1");
    await refreshPush();
    expect(messaging.unregister).toHaveBeenCalledTimes(2);
    expect(localStorage.getItem(pushReleaseKey)).toBeNull();
    expect(registerDevice).not.toHaveBeenCalled();
  });
});
