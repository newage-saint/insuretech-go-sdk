package models

// PolicyUpdateResponse represents a policy_update_response
type PolicyUpdateResponse struct {
	Message string `json:"message,omitempty"`
	Error   *Error `json:"error,omitempty"`
}
