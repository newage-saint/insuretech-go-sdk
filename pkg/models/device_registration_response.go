package models

// DeviceRegistrationResponse represents a device_registration_response
type DeviceRegistrationResponse struct {
	Message         string `json:"message,omitempty"`
	Error           *Error `json:"error,omitempty"`
	DeviceId        string `json:"device_id,omitempty"`
	ActivationToken string `json:"activation_token,omitempty"`
}
