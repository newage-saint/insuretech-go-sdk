package models

// StreamTelemetryResponse represents a stream_telemetry_response
type StreamTelemetryResponse struct {
	Received    bool   `json:"received,omitempty"`
	Message     string `json:"message,omitempty"`
	Error       *Error `json:"error,omitempty"`
	TelemetryId string `json:"telemetry_id,omitempty"`
}
