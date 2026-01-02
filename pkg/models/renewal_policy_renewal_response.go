package models

// RenewalPolicyRenewalResponse represents a renewal_policy_renewal_response
type RenewalPolicyRenewalResponse struct {
	Error       *Error `json:"error,omitempty"`
	NewPolicyId string `json:"new_policy_id,omitempty"`
	Message     string `json:"message,omitempty"`
}
