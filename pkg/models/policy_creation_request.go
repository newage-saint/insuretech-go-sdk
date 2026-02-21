package models

// PolicyCreationRequest represents a policy_creation_request
type PolicyCreationRequest struct {
	SumInsured    *Money         `json:"sum_insured,omitempty"`
	TenureMonths  int            `json:"tenure_months,omitempty"`
	PartnerId     string         `json:"partner_id"`
	Applicant     *Applicant     `json:"applicant,omitempty"`
	Nominees      []*Nominee     `json:"nominees,omitempty"`
	ProductId     string         `json:"product_id"`
	CustomerId    string         `json:"customer_id"`
	AgentId       string         `json:"agent_id"`
	Riders        []*PolicyRider `json:"riders,omitempty"`
	PremiumAmount *Money         `json:"premium_amount,omitempty"`
}
