package models

// ClaimRejectionRequest represents a claim_rejection_request
type ClaimRejectionRequest struct {
	ApproverId string `json:"approver_id"`
	Reason     string `json:"reason,omitempty"`
	ClaimId    string `json:"claim_id"`
}
