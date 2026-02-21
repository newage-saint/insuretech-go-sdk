package models

// ClaimApprovalRequest represents a claim_approval_request
type ClaimApprovalRequest struct {
	Notes          string `json:"notes,omitempty"`
	ClaimId        string `json:"claim_id"`
	ApproverId     string `json:"approver_id"`
	ApprovedAmount *Money `json:"approved_amount,omitempty"`
}
