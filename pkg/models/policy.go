package models

import (
	"time"
)

// Policy represents a policy
type Policy struct {
	DeletedAt               time.Time      `json:"deleted_at,omitempty"`
	TotalPayable            *Money         `json:"total_payable,omitempty"`
	ProductId               string         `json:"product_id,omitempty"`
	AgentId                 string         `json:"agent_id,omitempty"`
	Status                  *PolicyStatus  `json:"status,omitempty"`
	ReceiptNumber           string         `json:"receipt_number,omitempty"`
	EnrollmentEndDate       time.Time      `json:"enrollment_end_date,omitempty"`
	PolicyId                string         `json:"policy_id,omitempty"`
	ProposerDetails         *Applicant     `json:"proposer_details,omitempty"`
	OccupationRiskClass     string         `json:"occupation_risk_class,omitempty"`
	UnderwritingData        string         `json:"underwriting_data,omitempty"`
	SumInsuredCurrency      string         `json:"sum_insured_currency,omitempty"`
	PremiumAmount           *Money         `json:"premium_amount,omitempty"`
	IssuedAt                time.Time      `json:"issued_at,omitempty"`
	CreatedAt               time.Time      `json:"created_at,omitempty"`
	Riders                  []*PolicyRider `json:"riders,omitempty"`
	PaymentGatewayReference string         `json:"payment_gateway_reference,omitempty"`
	PolicyNumber            string         `json:"policy_number,omitempty"`
	StartDate               time.Time      `json:"start_date,omitempty"`
	Nominees                []*Nominee     `json:"nominees,omitempty"`
	VatTax                  *Money         `json:"vat_tax,omitempty"`
	PremiumCurrency         string         `json:"premium_currency,omitempty"`
	PartnerId               string         `json:"partner_id,omitempty"`
	UnderwritingDecisionId  string         `json:"underwriting_decision_id,omitempty"`
	PaymentFrequency        string         `json:"payment_frequency,omitempty"`
	ServiceFee              *Money         `json:"service_fee,omitempty"`
	ClaimsHistorySummary    string         `json:"claims_history_summary,omitempty"`
	CustomerId              string         `json:"customer_id,omitempty"`
	QuoteId                 string         `json:"quote_id,omitempty"`
	SumInsured              *Money         `json:"sum_insured,omitempty"`
	UpdatedAt               time.Time      `json:"updated_at,omitempty"`
	HasExistingPolicies     bool           `json:"has_existing_policies,omitempty"`
	ProviderName            string         `json:"provider_name,omitempty"`
	TenureMonths            int            `json:"tenure_months,omitempty"`
	EnrollmentStartDate     time.Time      `json:"enrollment_start_date,omitempty"`
	EndDate                 time.Time      `json:"end_date,omitempty"`
	PolicyDocumentUrl       string         `json:"policy_document_url,omitempty"`
}
