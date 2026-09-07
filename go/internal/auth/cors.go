package auth

import (
	"net/http"
	"strings"
)

// CORS answers the browser's cross-origin checks for the origins in
// allowed. An empty list allows no cross-origin call, which is the
// same-origin deployment of one host. A request from another
// origin gets no CORS headers, and the browser refuses it.
func CORS(allowed []string, next http.Handler) http.Handler {
	set := map[string]bool{}
	for _, o := range allowed {
		if o = strings.TrimSpace(o); o != "" {
			set[o] = true
		}
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin != "" && set[origin] {
			h := w.Header()
			h.Set("Access-Control-Allow-Origin", origin)
			h.Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			h.Set("Access-Control-Allow-Headers", "Authorization, Content-Type, Connect-Protocol-Version, Connect-Timeout-Ms")
			// A browser reads no response header that this list omits.
			// The refusal header names why a call was refused, and the
			// web app reads it cross-origin (F-59, D-590).
			h.Set("Access-Control-Expose-Headers", "Connect-Protocol-Version, "+RefusalHeader)
			h.Set("Access-Control-Max-Age", "600")
			h.Add("Vary", "Origin")
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

// ParseOrigins splits the ALLOWED_ORIGINS value: a comma-separated list
// of origins, or empty for same origin only.
func ParseOrigins(v string) []string {
	var out []string
	for _, o := range strings.Split(v, ",") {
		if o = strings.TrimSpace(o); o != "" {
			out = append(out, o)
		}
	}
	return out
}
