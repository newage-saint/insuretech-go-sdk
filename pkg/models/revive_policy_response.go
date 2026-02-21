package models

// RevivePolicyResponse represents a revive_policy_response
type RevivePolicyResponse struct {
	Error   *Error `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}
