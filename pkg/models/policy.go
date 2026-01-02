package models

import (
	"time"
)

// Policy represents a policy
type Policy struct {
	CustomerId             string         `json:"customer_id"`
	UnderwritingDecisionId string         `json:"underwriting_decision_id,omitempty"`
	StartDate              time.Time      `json:"start_date"`
	IssuedAt               time.Time      `json:"issued_at,omitempty"`
	PolicyDocumentUrl      string         `json:"policy_document_url,omitempty"`
	DeletedAt              time.Time      `json:"deleted_at,omitempty"`
	Riders                 []*PolicyRider `json:"riders,omitempty"`
	PartnerId              string         `json:"partner_id,omitempty"`
	PremiumAmount          *Money         `json:"premium_amount"`
	CreatedAt              time.Time      `json:"created_at"`
	PolicyId               string         `json:"policy_id"`
	PolicyNumber           string         `json:"policy_number"`
	QuoteId                string         `json:"quote_id,omitempty"`
	Status                 interface{}    `json:"status"`
	TenureMonths           int            `json:"tenure_months"`
	EndDate                time.Time      `json:"end_date"`
	UpdatedAt              time.Time      `json:"updated_at"`
	ProductId              string         `json:"product_id"`
	AgentId                string         `json:"agent_id,omitempty"`
	SumInsured             *Money         `json:"sum_insured"`
	Nominees               []*Nominee     `json:"nominees,omitempty"`
}
