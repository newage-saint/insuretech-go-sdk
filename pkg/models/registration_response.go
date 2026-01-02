package models

// RegistrationResponse represents a registration_response
type RegistrationResponse struct {
	UserId  string `json:"user_id,omitempty"`
	Message string `json:"message,omitempty"`
	OtpSent bool   `json:"otp_sent,omitempty"`
	Error   *Error `json:"error,omitempty"`
}
