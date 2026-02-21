package models

// DeviceDeactivationResponse represents a device_deactivation_response
type DeviceDeactivationResponse struct {
	Error   *Error `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}
