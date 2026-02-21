package models

import (
	"time"
)

// Claim represents a claim
type Claim struct {
	PolicyId              string               `json:"policy_id,omitempty"`
	IncidentDescription   string               `json:"incident_description,omitempty"`
	BankDetailsForPayout  string               `json:"bank_details_for_payout,omitempty"`
	ApprovedAmount        *Money               `json:"approved_amount,omitempty"`
	IncidentDate          time.Time            `json:"incident_date,omitempty"`
	DeletedAt             time.Time            `json:"deleted_at,omitempty"`
	CreatedAt             time.Time            `json:"created_at,omitempty"`
	RejectionReason       string               `json:"rejection_reason,omitempty"`
	FraudCheck            *FraudCheckResult    `json:"fraud_check,omitempty"`
	AppealOptionAvailable bool                 `json:"appeal_option_available,omitempty"`
	ApprovedCurrency      string               `json:"approved_currency,omitempty"`
	CoPayAmount           *Money               `json:"co_pay_amount,omitempty"`
	ClaimedAmount         *Money               `json:"claimed_amount,omitempty"`
	SettledAmount         *Money               `json:"settled_amount,omitempty"`
	SubmittedAt           time.Time            `json:"submitted_at,omitempty"`
	Documents             []*ClaimDocument     `json:"documents,omitempty"`
	InAppMessages         string               `json:"in_app_messages,omitempty"`
	DeductibleAmount      *Money               `json:"deductible_amount,omitempty"`
	ClaimNumber           string               `json:"claim_number,omitempty"`
	ProcessingType        *ClaimProcessingType `json:"processing_type,omitempty"`
	ClaimedCurrency       string               `json:"claimed_currency,omitempty"`
	Type                  *ClaimType           `json:"type,omitempty"`
	ApprovedAt            time.Time            `json:"approved_at,omitempty"`
	UpdatedAt             time.Time            `json:"updated_at,omitempty"`
	Approvals             []*ClaimApproval     `json:"approvals,omitempty"`
	PlaceOfIncident       string               `json:"place_of_incident,omitempty"`
	ProcessorNotes        string               `json:"processor_notes,omitempty"`
	SettledCurrency       string               `json:"settled_currency,omitempty"`
	ClaimId               string               `json:"claim_id,omitempty"`
	CustomerId            string               `json:"customer_id,omitempty"`
	Status                *ClaimStatus         `json:"status,omitempty"`
	SettledAt             time.Time            `json:"settled_at,omitempty"`
}
