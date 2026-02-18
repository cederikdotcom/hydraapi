package hydraapi

import (
	"math"
	"net/http"
	"time"
)

// HealthResponse is the standard shape all hydra services return from /api/v1/health.
type HealthResponse struct {
	Service       string         `json:"service"`
	Version       string         `json:"version"`
	District      string         `json:"district"`
	UptimeSeconds int            `json:"uptime_seconds"`
	Status        string         `json:"status"`          // "ok" or "degraded"
	Extra         map[string]any `json:"extra,omitempty"` // tool-specific fields
}

// NewHealthHandler returns an http.HandlerFunc that serves a standard health response.
// The extraFn callback is called on each request to populate tool-specific fields.
// Pass nil for extraFn if no extra fields are needed.
func NewHealthHandler(service, version, district string, startTime time.Time, extraFn func() map[string]any) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uptime := int(math.Round(time.Since(startTime).Seconds()))

		extra := map[string]any(nil)
		status := "ok"
		if extraFn != nil {
			extra = extraFn()
			// Allow extraFn to signal degraded status via a special key.
			if v, ok := extra["_status"]; ok {
				if s, ok := v.(string); ok {
					status = s
				}
				delete(extra, "_status")
			}
			if len(extra) == 0 {
				extra = nil
			}
		}

		WriteJSON(w, http.StatusOK, HealthResponse{
			Service:       service,
			Version:       version,
			District:      district,
			UptimeSeconds: uptime,
			Status:        status,
			Extra:         extra,
		})
	}
}
