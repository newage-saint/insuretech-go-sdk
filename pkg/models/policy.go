package models

import (
	"time"
)

// Policy represents a policy
type Policy struct {
	Status                  *PolicyStatus  `json:"status,omitempty"`
	IssuedAt                time.Time      `json:"issued_at,omitempty"`
	DeletedAt               time.Time      `json:"deleted_at,omitempty"`
	Riders                  []*PolicyRider `json:"riders,omitempty"`
	ServiceFee              *Money         `json:"service_fee,omitempty"`
	QuoteId                 string         `json:"quote_id,omitempty"`
	VatTax                  *Money         `json:"vat_tax,omitempty"`
	ReceiptNumber           string         `json:"receipt_number,omitempty"`
	OccupationRiskClass     string         `json:"occupation_risk_class,omitempty"`
	ClaimsHistorySummary    string         `json:"claims_history_summary,omitempty"`
	CustomerId              string         `json:"customer_id,omitempty"`
	StartDate               time.Time      `json:"start_date,omitempty"`
	CreatedAt               time.Time      `json:"created_at,omitempty"`
	TotalPayable            *Money         `json:"total_payable,omitempty"`
	AgentId                 string         `json:"agent_id,omitempty"`
	SumInsured              *Money         `json:"sum_insured,omitempty"`
	PolicyDocumentUrl       string         `json:"policy_document_url,omitempty"`
	EnrollmentStartDate     time.Time      `json:"enrollment_start_date,omitempty"`
	ProductId               string         `json:"product_id,omitempty"`
	PartnerId               string         `json:"partner_id,omitempty"`
	PremiumAmount           *Money         `json:"premium_amount,omitempty"`
	EndDate                 time.Time      `json:"end_date,omitempty"`
	ProposerDetails         *Applicant     `json:"proposer_details,omitempty"`
	HasExistingPolicies     bool           `json:"has_existing_policies,omitempty"`
	UnderwritingDecisionId  string         `json:"underwriting_decision_id,omitempty"`
	PaymentFrequency        string         `json:"payment_frequency,omitempty"`
	EnrollmentEndDate       time.Time      `json:"enrollment_end_date,omitempty"`
	UnderwritingData        string         `json:"underwriting_data,omitempty"`
	TenureMonths            int            `json:"tenure_months,omitempty"`
	Nominees                []*Nominee     `json:"nominees,omitempty"`
	SumInsuredCurrency      string         `json:"sum_insured_currency,omitempty"`
	PolicyNumber            string         `json:"policy_number,omitempty"`
	UpdatedAt               time.Time      `json:"updated_at,omitempty"`
	PaymentGatewayReference string         `json:"payment_gateway_reference,omitempty"`
	ProviderName            string         `json:"provider_name,omitempty"`
	PremiumCurrency         string         `json:"premium_currency,omitempty"`
	PolicyId                string         `json:"policy_id,omitempty"`
}
