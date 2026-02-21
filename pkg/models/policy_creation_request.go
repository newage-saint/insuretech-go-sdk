package models

// PolicyCreationRequest represents a policy_creation_request
type PolicyCreationRequest struct {
	AgentId       string         `json:"agent_id"`
	Applicant     *Applicant     `json:"applicant,omitempty"`
	Nominees      []*Nominee     `json:"nominees,omitempty"`
	PremiumAmount *Money         `json:"premium_amount,omitempty"`
	SumInsured    *Money         `json:"sum_insured,omitempty"`
	TenureMonths  int            `json:"tenure_months,omitempty"`
	ProductId     string         `json:"product_id"`
	PartnerId     string         `json:"partner_id"`
	Riders        []*PolicyRider `json:"riders,omitempty"`
	CustomerId    string         `json:"customer_id"`
}
