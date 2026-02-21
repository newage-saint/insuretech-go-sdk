package models

// RegistrationRequest represents a registration_request
type RegistrationRequest struct {
	DeviceId     string `json:"device_id"`
	DeviceType   string `json:"device_type,omitempty"`
	MobileNumber string `json:"mobile_number,omitempty"`
	Email        string `json:"email"`
	Password     string `json:"password"`
}
