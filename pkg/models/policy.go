package models

import (
	"time"
)

// Policy represents a policy
type Policy struct {
	DeletedAt               time.Time      `json:"deleted_at,omitempty"`
	PremiumCurrency         string         `json:"premium_currency,omitempty"`
	TenureMonths            int            `json:"tenure_months,omitempty"`
	EndDate                 time.Time      `json:"end_date,omitempty"`
	PolicyDocumentUrl       string         `json:"policy_document_url,omitempty"`
	ReceiptNumber           string         `json:"receipt_number,omitempty"`
	ProviderName            string         `json:"provider_name,omitempty"`
	EnrollmentStartDate     time.Time      `json:"enrollment_start_date,omitempty"`
	PolicyId                string         `json:"policy_id,omitempty"`
	PolicyNumber            string         `json:"policy_number,omitempty"`
	SumInsured              *Money         `json:"sum_insured,omitempty"`
	IssuedAt                time.Time      `json:"issued_at,omitempty"`
	ServiceFee              *Money         `json:"service_fee,omitempty"`
	UnderwritingDecisionId  string         `json:"underwriting_decision_id,omitempty"`
	StartDate               time.Time      `json:"start_date,omitempty"`
	TotalPayable            *Money         `json:"total_payable,omitempty"`
	CustomerId              string         `json:"customer_id,omitempty"`
	PartnerId               string         `json:"partner_id,omitempty"`
	Nominees                []*Nominee     `json:"nominees,omitempty"`
	VatTax                  *Money         `json:"vat_tax,omitempty"`
	SumInsuredCurrency      string         `json:"sum_insured_currency,omitempty"`
	ProductId               string         `json:"product_id,omitempty"`
	Status                  *PolicyStatus  `json:"status,omitempty"`
	ProposerDetails         *Applicant     `json:"proposer_details,omitempty"`
	ClaimsHistorySummary    string         `json:"claims_history_summary,omitempty"`
	QuoteId                 string         `json:"quote_id,omitempty"`
	PaymentFrequency        string         `json:"payment_frequency,omitempty"`
	OccupationRiskClass     string         `json:"occupation_risk_class,omitempty"`
	PremiumAmount           *Money         `json:"premium_amount,omitempty"`
	Riders                  []*PolicyRider `json:"riders,omitempty"`
	PaymentGatewayReference string         `json:"payment_gateway_reference,omitempty"`
	HasExistingPolicies     bool           `json:"has_existing_policies,omitempty"`
	EnrollmentEndDate       time.Time      `json:"enrollment_end_date,omitempty"`
	UnderwritingData        string         `json:"underwriting_data,omitempty"`
	AgentId                 string         `json:"agent_id,omitempty"`
	CreatedAt               time.Time      `json:"created_at,omitempty"`
	UpdatedAt               time.Time      `json:"updated_at,omitempty"`
}
