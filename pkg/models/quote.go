package models

import (
	"time"
)

// Quote represents a quote
type Quote struct {
	QuoteNumber         string      `json:"quote_number"`
	Status              interface{} `json:"status"`
	RiderPremium        *Money      `json:"rider_premium,omitempty"`
	TaxAmount           *Money      `json:"tax_amount,omitempty"`
	ApplicantOccupation string      `json:"applicant_occupation,omitempty"`
	ConvertedAt         time.Time   `json:"converted_at,omitempty"`
	Id                  string      `json:"id"`
	PremiumPaymentMode  string      `json:"premium_payment_mode"`
	SelectedRiders      string      `json:"selected_riders,omitempty"`
	ApplicantAge        int         `json:"applicant_age"`
	Smoker              bool        `json:"smoker,omitempty"`
	TermYears           int         `json:"term_years"`
	BasePremium         *Money      `json:"base_premium,omitempty"`
	TotalPremium        *Money      `json:"total_premium,omitempty"`
	PremiumCalculation  string      `json:"premium_calculation,omitempty"`
	BeneficiaryId       string      `json:"beneficiary_id"`
	InsurerProductId    string      `json:"insurer_product_id"`
	SumAssured          *Money      `json:"sum_assured,omitempty"`
	ValidUntil          time.Time   `json:"valid_until"`
	ConvertedPolicyId   string      `json:"converted_policy_id,omitempty"`
	AuditInfo           *AuditInfo  `json:"audit_info,omitempty"`
}
