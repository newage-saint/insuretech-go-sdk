package models

// LoginRequest represents a login_request
type LoginRequest struct {
	Password     string `json:"password"`
	DeviceId     string `json:"device_id"`
	DeviceType   string `json:"device_type,omitempty"`
	DeviceName   string `json:"device_name,omitempty"`
	MobileNumber string `json:"mobile_number,omitempty"`
}
