package models

// KYCVerificationStartResponse represents a kyc_verification_start_response
type KYCVerificationStartResponse struct {
	Error             *Error `json:"error,omitempty"`
	KycVerificationId string `json:"kyc_verification_id,omitempty"`
	Message           string `json:"message,omitempty"`
}
