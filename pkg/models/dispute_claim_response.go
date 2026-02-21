package models

// DisputeClaimResponse represents a dispute_claim_response
type DisputeClaimResponse struct {
	Message   string `json:"message,omitempty"`
	Error     *Error `json:"error,omitempty"`
	DisputeId string `json:"dispute_id,omitempty"`
}
