package models

// ClaimRejectionResponse represents a claim_rejection_response
type ClaimRejectionResponse struct {
	Message string `json:"message,omitempty"`
	Error   *Error `json:"error,omitempty"`
}
