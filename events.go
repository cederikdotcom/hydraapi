package hydraapi

import "time"

// EventEnvelope is the common envelope all SSE event payloads should include.
type EventEnvelope struct {
	District  string `json:"district"`
	Timestamp string `json:"timestamp"` // RFC3339
}

// NewEventEnvelope creates an EventEnvelope with the given district and the current time.
func NewEventEnvelope(district string) EventEnvelope {
	return EventEnvelope{
		District:  district,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}
}
