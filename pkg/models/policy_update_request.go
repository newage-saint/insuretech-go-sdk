package models

// PolicyUpdateRequest represents a policy_update_request
type PolicyUpdateRequest struct {
	Nominees []*Nominee `json:"nominees,omitempty"`
	Address  string     `json:"address,omitempty"`
	PolicyId string     `json:"policy_id"`
}
