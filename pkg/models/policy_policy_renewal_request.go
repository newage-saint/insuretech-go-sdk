package models

// PolicyPolicyRenewalRequest represents a policy_policy_renewal_request
type PolicyPolicyRenewalRequest struct {
	TenureMonths   int        `json:"tenure_months,omitempty"`
	UpdateNominees bool       `json:"update_nominees,omitempty"`
	Nominees       []*Nominee `json:"nominees,omitempty"`
	PolicyId       string     `json:"policy_id"`
}
