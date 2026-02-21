package models

import (
	"time"
)

// PolicyCreatedEvent represents a policy_created_event
type PolicyCreatedEvent struct {
	EventId       string    `json:"event_id,omitempty"`
	PolicyId      string    `json:"policy_id,omitempty"`
	CustomerId    string    `json:"customer_id,omitempty"`
	SumInsured    *Money    `json:"sum_insured,omitempty"`
	PolicyNumber  string    `json:"policy_number,omitempty"`
	ProductId     string    `json:"product_id,omitempty"`
	PremiumAmount *Money    `json:"premium_amount,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
	CorrelationId string    `json:"correlation_id,omitempty"`
}
