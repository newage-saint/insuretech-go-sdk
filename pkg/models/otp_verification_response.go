package models

// OTPVerificationResponse represents a otp_verification_response
type OTPVerificationResponse struct {
	Verified bool   `json:"verified,omitempty"`
	UserId   string `json:"user_id,omitempty"`
	Message  string `json:"message,omitempty"`
	Error    *Error `json:"error,omitempty"`
}
