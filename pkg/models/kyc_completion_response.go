package models

// KYCCompletionResponse represents a kyc_completion_response
type KYCCompletionResponse struct {
	Error     *Error `json:"error,omitempty"`
	KycStatus string `json:"kyc_status,omitempty"`
	Message   string `json:"message,omitempty"`
}
