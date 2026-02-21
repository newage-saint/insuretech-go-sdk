package models

// DisputeClaimResponse represents a dispute_claim_response
type DisputeClaimResponse struct {
	DisputeId string `json:"dispute_id,omitempty"`
	Message   string `json:"message,omitempty"`
	Error     *Error `json:"error,omitempty"`
}
