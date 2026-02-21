package models

import (
	"time"
)

// Claim represents a claim
type Claim struct {
	IncidentDescription   string               `json:"incident_description,omitempty"`
	ProcessingType        *ClaimProcessingType `json:"processing_type,omitempty"`
	CoPayAmount           *Money               `json:"co_pay_amount,omitempty"`
	PolicyId              string               `json:"policy_id,omitempty"`
	ClaimedAmount         *Money               `json:"claimed_amount,omitempty"`
	IncidentDate          time.Time            `json:"incident_date,omitempty"`
	DeletedAt             time.Time            `json:"deleted_at,omitempty"`
	AppealOptionAvailable bool                 `json:"appeal_option_available,omitempty"`
	InAppMessages         string               `json:"in_app_messages,omitempty"`
	ProcessorNotes        string               `json:"processor_notes,omitempty"`
	ClaimedCurrency       string               `json:"claimed_currency,omitempty"`
	SettledCurrency       string               `json:"settled_currency,omitempty"`
	SettledAt             time.Time            `json:"settled_at,omitempty"`
	DeductibleAmount      *Money               `json:"deductible_amount,omitempty"`
	ApprovedAmount        *Money               `json:"approved_amount,omitempty"`
	CreatedAt             time.Time            `json:"created_at,omitempty"`
	RejectionReason       string               `json:"rejection_reason,omitempty"`
	BankDetailsForPayout  string               `json:"bank_details_for_payout,omitempty"`
	ApprovedCurrency      string               `json:"approved_currency,omitempty"`
	Status                *ClaimStatus         `json:"status,omitempty"`
	ApprovedAt            time.Time            `json:"approved_at,omitempty"`
	ClaimId               string               `json:"claim_id,omitempty"`
	UpdatedAt             time.Time            `json:"updated_at,omitempty"`
	CustomerId            string               `json:"customer_id,omitempty"`
	Type                  *ClaimType           `json:"type,omitempty"`
	SettledAmount         *Money               `json:"settled_amount,omitempty"`
	SubmittedAt           time.Time            `json:"submitted_at,omitempty"`
	Documents             []*ClaimDocument     `json:"documents,omitempty"`
	Approvals             []*ClaimApproval     `json:"approvals,omitempty"`
	FraudCheck            *FraudCheckResult    `json:"fraud_check,omitempty"`
	PlaceOfIncident       string               `json:"place_of_incident,omitempty"`
	ClaimNumber           string               `json:"claim_number,omitempty"`
}
