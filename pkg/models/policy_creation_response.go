package models

// PolicyCreationResponse represents a policy_creation_response
type PolicyCreationResponse struct {
	Message      string `json:"message,omitempty"`
	Error        *Error `json:"error,omitempty"`
	PolicyId     string `json:"policy_id,omitempty"`
	PolicyNumber string `json:"policy_number,omitempty"`
}
