// The byte caps of the API. The server refuses a larger request, so the
// web app checks first and names the limit.

// The message and answer cap of the Chat RPC (agentsvc.MaxMessageBytes).
export const maxMessageBytes = 8 << 10;

// The upload cap of ImportCollection (collectionsvc.maxUpload).
export const maxUploadBytes = 5 << 20;

export function byteLength(s: string): number {
  return new TextEncoder().encode(s).length;
}
