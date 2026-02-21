package models

import (
	"time"
)

// Quote represents a quote
type Quote struct {
	SelectedRiders      string      `json:"selected_riders,omitempty"`
	QuoteNumber         string      `json:"quote_number"`
	InsurerProductId    string      `json:"insurer_product_id"`
	TermYears           int         `json:"term_years"`
	BasePremium         *Money      `json:"base_premium,omitempty"`
	RiderPremium        *Money      `json:"rider_premium,omitempty"`
	TaxAmount           *Money      `json:"tax_amount,omitempty"`
	AuditInfo           interface{} `json:"audit_info"`
	SumAssured          *Money      `json:"sum_assured,omitempty"`
	PremiumPaymentMode  string      `json:"premium_payment_mode"`
	PremiumCalculation  string      `json:"premium_calculation,omitempty"`
	ApplicantAge        int         `json:"applicant_age"`
	ValidUntil          time.Time   `json:"valid_until"`
	Id                  string      `json:"id"`
	Status              interface{} `json:"status"`
	ApplicantOccupation string      `json:"applicant_occupation,omitempty"`
	Smoker              bool        `json:"smoker,omitempty"`
	BeneficiaryId       string      `json:"beneficiary_id"`
	TotalPremium        *Money      `json:"total_premium,omitempty"`
	ConvertedPolicyId   string      `json:"converted_policy_id,omitempty"`
	ConvertedAt         time.Time   `json:"converted_at,omitempty"`
}
