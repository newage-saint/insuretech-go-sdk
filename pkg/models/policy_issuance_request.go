package models

// PolicyIssuanceRequest represents a policy_issuance_request
type PolicyIssuanceRequest struct {
	QuoteId   string `json:"quote_id"`
	PaymentId string `json:"payment_id"`
	PolicyId  string `json:"policy_id"`
}
