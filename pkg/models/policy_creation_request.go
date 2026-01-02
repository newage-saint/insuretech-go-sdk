package models

// PolicyCreationRequest represents a policy_creation_request
type PolicyCreationRequest struct {
	Applicant     *Applicant     `json:"applicant,omitempty"`
	Nominees      []*Nominee     `json:"nominees,omitempty"`
	Riders        []*PolicyRider `json:"riders,omitempty"`
	SumInsured    *Money         `json:"sum_insured,omitempty"`
	TenureMonths  int            `json:"tenure_months,omitempty"`
	ProductId     string         `json:"product_id"`
	CustomerId    string         `json:"customer_id"`
	PartnerId     string         `json:"partner_id"`
	AgentId       string         `json:"agent_id"`
	PremiumAmount *Money         `json:"premium_amount,omitempty"`
}
