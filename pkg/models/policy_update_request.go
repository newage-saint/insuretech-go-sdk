package models

// PolicyUpdateRequest represents a policy_update_request
type PolicyUpdateRequest struct {
	PolicyId string     `json:"policy_id"`
	Nominees []*Nominee `json:"nominees,omitempty"`
	Address  string     `json:"address,omitempty"`
}
