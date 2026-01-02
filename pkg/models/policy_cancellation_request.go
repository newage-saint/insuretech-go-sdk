package models

// PolicyCancellationRequest represents a policy_cancellation_request
type PolicyCancellationRequest struct {
	Reason   string `json:"reason,omitempty"`
	PolicyId string `json:"policy_id"`
}
