package models

import (
	"time"
)

// Claim represents a claim
type Claim struct {
	SettledAmount         *Money               `json:"settled_amount,omitempty"`
	IncidentDescription   string               `json:"incident_description,omitempty"`
	FraudCheck            *FraudCheckResult    `json:"fraud_check,omitempty"`
	ClaimId               string               `json:"claim_id,omitempty"`
	CustomerId            string               `json:"customer_id,omitempty"`
	RejectionReason       string               `json:"rejection_reason,omitempty"`
	Type                  *ClaimType           `json:"type,omitempty"`
	ClaimedAmount         *Money               `json:"claimed_amount,omitempty"`
	ProcessingType        *ClaimProcessingType `json:"processing_type,omitempty"`
	InAppMessages         string               `json:"in_app_messages,omitempty"`
	SubmittedAt           time.Time            `json:"submitted_at,omitempty"`
	SettledAt             time.Time            `json:"settled_at,omitempty"`
	CreatedAt             time.Time            `json:"created_at,omitempty"`
	DeletedAt             time.Time            `json:"deleted_at,omitempty"`
	Documents             []*ClaimDocument     `json:"documents,omitempty"`
	Status                *ClaimStatus         `json:"status,omitempty"`
	ApprovedAmount        *Money               `json:"approved_amount,omitempty"`
	Approvals             []*ClaimApproval     `json:"approvals,omitempty"`
	BankDetailsForPayout  string               `json:"bank_details_for_payout,omitempty"`
	ProcessorNotes        string               `json:"processor_notes,omitempty"`
	SettledCurrency       string               `json:"settled_currency,omitempty"`
	ClaimNumber           string               `json:"claim_number,omitempty"`
	IncidentDate          time.Time            `json:"incident_date,omitempty"`
	UpdatedAt             time.Time            `json:"updated_at,omitempty"`
	DeductibleAmount      *Money               `json:"deductible_amount,omitempty"`
	CoPayAmount           *Money               `json:"co_pay_amount,omitempty"`
	PlaceOfIncident       string               `json:"place_of_incident,omitempty"`
	ClaimedCurrency       string               `json:"claimed_currency,omitempty"`
	ApprovedCurrency      string               `json:"approved_currency,omitempty"`
	PolicyId              string               `json:"policy_id,omitempty"`
	ApprovedAt            time.Time            `json:"approved_at,omitempty"`
	AppealOptionAvailable bool                 `json:"appeal_option_available,omitempty"`
}
