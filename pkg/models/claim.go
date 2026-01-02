package models

import (
	"time"
)

// Claim represents a claim
type Claim struct {
	ApprovedAt          time.Time         `json:"approved_at,omitempty"`
	CreatedAt           time.Time         `json:"created_at"`
	Approvals           []*ClaimApproval  `json:"approvals,omitempty"`
	ClaimId             string            `json:"claim_id"`
	PolicyId            string            `json:"policy_id"`
	Status              interface{}       `json:"status"`
	Type                *ClaimType        `json:"type"`
	ApprovedAmount      *Money            `json:"approved_amount,omitempty"`
	IncidentDate        time.Time         `json:"incident_date"`
	UpdatedAt           time.Time         `json:"updated_at"`
	RejectionReason     string            `json:"rejection_reason,omitempty"`
	CustomerId          string            `json:"customer_id"`
	SettledAt           time.Time         `json:"settled_at,omitempty"`
	Documents           []*ClaimDocument  `json:"documents,omitempty"`
	FraudCheck          *FraudCheckResult `json:"fraud_check,omitempty"`
	ClaimNumber         string            `json:"claim_number"`
	ClaimedAmount       *Money            `json:"claimed_amount"`
	DeletedAt           time.Time         `json:"deleted_at,omitempty"`
	SettledAmount       *Money            `json:"settled_amount,omitempty"`
	IncidentDescription string            `json:"incident_description"`
	SubmittedAt         time.Time         `json:"submitted_at"`
}
