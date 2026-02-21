package models

import (
	"time"
)

// BeneficiaryKYCCompletedEvent represents a beneficiary_kyc_completed_event
type BeneficiaryKYCCompletedEvent struct {
	KycStatus     string    `json:"kyc_status,omitempty"`
	CorrelationId string    `json:"correlation_id,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
	EventId       string    `json:"event_id,omitempty"`
	BeneficiaryId string    `json:"beneficiary_id,omitempty"`
}
