package models

// TelemetrySendingResponse represents a telemetry_sending_response
type TelemetrySendingResponse struct {
	Accepted    bool   `json:"accepted,omitempty"`
	Error       *Error `json:"error,omitempty"`
	TelemetryId string `json:"telemetry_id,omitempty"`
}
