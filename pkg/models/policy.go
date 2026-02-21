package models

import (
	"time"
)

// Policy represents a policy
type Policy struct {
	SumInsured              *Money         `json:"sum_insured,omitempty"`
	TenureMonths            int            `json:"tenure_months,omitempty"`
	UnderwritingDecisionId  string         `json:"underwriting_decision_id,omitempty"`
	Riders                  []*PolicyRider `json:"riders,omitempty"`
	VatTax                  *Money         `json:"vat_tax,omitempty"`
	OccupationRiskClass     string         `json:"occupation_risk_class,omitempty"`
	HasExistingPolicies     bool           `json:"has_existing_policies,omitempty"`
	IssuedAt                time.Time      `json:"issued_at,omitempty"`
	UpdatedAt               time.Time      `json:"updated_at,omitempty"`
	TotalPayable            *Money         `json:"total_payable,omitempty"`
	EnrollmentStartDate     time.Time      `json:"enrollment_start_date,omitempty"`
	PremiumCurrency         string         `json:"premium_currency,omitempty"`
	Nominees                []*Nominee     `json:"nominees,omitempty"`
	ProposerDetails         *Applicant     `json:"proposer_details,omitempty"`
	PolicyId                string         `json:"policy_id,omitempty"`
	ProductId               string         `json:"product_id,omitempty"`
	AgentId                 string         `json:"agent_id,omitempty"`
	CreatedAt               time.Time      `json:"created_at,omitempty"`
	DeletedAt               time.Time      `json:"deleted_at,omitempty"`
	ClaimsHistorySummary    string         `json:"claims_history_summary,omitempty"`
	PolicyDocumentUrl       string         `json:"policy_document_url,omitempty"`
	PaymentGatewayReference string         `json:"payment_gateway_reference,omitempty"`
	ReceiptNumber           string         `json:"receipt_number,omitempty"`
	CustomerId              string         `json:"customer_id,omitempty"`
	PartnerId               string         `json:"partner_id,omitempty"`
	QuoteId                 string         `json:"quote_id,omitempty"`
	ServiceFee              *Money         `json:"service_fee,omitempty"`
	PolicyNumber            string         `json:"policy_number,omitempty"`
	PaymentFrequency        string         `json:"payment_frequency,omitempty"`
	PremiumAmount           *Money         `json:"premium_amount,omitempty"`
	StartDate               time.Time      `json:"start_date,omitempty"`
	Status                  *PolicyStatus  `json:"status,omitempty"`
	EnrollmentEndDate       time.Time      `json:"enrollment_end_date,omitempty"`
	UnderwritingData        string         `json:"underwriting_data,omitempty"`
	SumInsuredCurrency      string         `json:"sum_insured_currency,omitempty"`
	EndDate                 time.Time      `json:"end_date,omitempty"`
	ProviderName            string         `json:"provider_name,omitempty"`
}
