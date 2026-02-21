package models

import (
	"time"
)

// ClaimSubmittedEvent represents a claim_submitted_event
type ClaimSubmittedEvent struct {
	ClaimedAmount *Money    `json:"claimed_amount,omitempty"`
	EventId       string    `json:"event_id,omitempty"`
	PolicyId      string    `json:"policy_id,omitempty"`
	ClaimType     string    `json:"claim_type,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
	CorrelationId string    `json:"correlation_id,omitempty"`
	ClaimId       string    `json:"claim_id,omitempty"`
	ClaimNumber   string    `json:"claim_number,omitempty"`
	CustomerId    string    `json:"customer_id,omitempty"`
}
