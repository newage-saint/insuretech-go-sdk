package models

import (
	"time"
)

// Claim represents a claim
type Claim struct {
	ApprovedAmount        *Money               `json:"approved_amount,omitempty"`
	Documents             []*ClaimDocument     `json:"documents,omitempty"`
	FraudCheck            *FraudCheckResult    `json:"fraud_check,omitempty"`
	PlaceOfIncident       string               `json:"place_of_incident,omitempty"`
	ClaimedCurrency       string               `json:"claimed_currency,omitempty"`
	DeductibleAmount      *Money               `json:"deductible_amount,omitempty"`
	Type                  *ClaimType           `json:"type,omitempty"`
	IncidentDate          time.Time            `json:"incident_date,omitempty"`
	UpdatedAt             time.Time            `json:"updated_at,omitempty"`
	DeletedAt             time.Time            `json:"deleted_at,omitempty"`
	AppealOptionAvailable bool                 `json:"appeal_option_available,omitempty"`
	ApprovedCurrency      string               `json:"approved_currency,omitempty"`
	ClaimNumber           string               `json:"claim_number,omitempty"`
	IncidentDescription   string               `json:"incident_description,omitempty"`
	ApprovedAt            time.Time            `json:"approved_at,omitempty"`
	SettledAt             time.Time            `json:"settled_at,omitempty"`
	ClaimedAmount         *Money               `json:"claimed_amount,omitempty"`
	RejectionReason       string               `json:"rejection_reason,omitempty"`
	PolicyId              string               `json:"policy_id,omitempty"`
	CustomerId            string               `json:"customer_id,omitempty"`
	ProcessingType        *ClaimProcessingType `json:"processing_type,omitempty"`
	SettledCurrency       string               `json:"settled_currency,omitempty"`
	ClaimId               string               `json:"claim_id,omitempty"`
	Status                *ClaimStatus         `json:"status,omitempty"`
	SettledAmount         *Money               `json:"settled_amount,omitempty"`
	Approvals             []*ClaimApproval     `json:"approvals,omitempty"`
	BankDetailsForPayout  string               `json:"bank_details_for_payout,omitempty"`
	InAppMessages         string               `json:"in_app_messages,omitempty"`
	CoPayAmount           *Money               `json:"co_pay_amount,omitempty"`
	ProcessorNotes        string               `json:"processor_notes,omitempty"`
	SubmittedAt           time.Time            `json:"submitted_at,omitempty"`
	CreatedAt             time.Time            `json:"created_at,omitempty"`
}
