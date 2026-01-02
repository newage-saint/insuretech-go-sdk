package models

// KYCVerificationStartRequest represents a kyc_verification_start_request
type KYCVerificationStartRequest struct {
	EntityId   string `json:"entity_id"`
	Method     string `json:"method,omitempty"`
	Type       string `json:"type"`
	EntityType string `json:"entity_type"`
}
