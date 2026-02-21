package models

import (
	"time"
)

// ClaimApprovedEvent represents a claim_approved_event
type ClaimApprovedEvent struct {
	ClaimId        string    `json:"claim_id,omitempty"`
	ClaimNumber    string    `json:"claim_number,omitempty"`
	ApproverId     string    `json:"approver_id,omitempty"`
	ApprovedAmount *Money    `json:"approved_amount,omitempty"`
	ApprovalLevel  string    `json:"approval_level,omitempty"`
	Timestamp      time.Time `json:"timestamp,omitempty"`
	CorrelationId  string    `json:"correlation_id,omitempty"`
	EventId        string    `json:"event_id,omitempty"`
}
