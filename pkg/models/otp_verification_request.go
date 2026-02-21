package models

// OTPVerificationRequest represents a otp_verification_request
type OTPVerificationRequest struct {
	Code  string `json:"code,omitempty"`
	OtpId string `json:"otp_id"`
}
