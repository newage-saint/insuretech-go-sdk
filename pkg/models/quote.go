package models

import (
	"time"
)

// Quote represents a quote
type Quote struct {
	ConvertedAt         time.Time   `json:"converted_at,omitempty"`
	AuditInfo           interface{} `json:"audit_info"`
	Id                  string      `json:"id"`
	QuoteNumber         string      `json:"quote_number"`
	BeneficiaryId       string      `json:"beneficiary_id"`
	Status              interface{} `json:"status"`
	BasePremium         *Money      `json:"base_premium,omitempty"`
	TaxAmount           *Money      `json:"tax_amount,omitempty"`
	TotalPremium        *Money      `json:"total_premium,omitempty"`
	SelectedRiders      string      `json:"selected_riders,omitempty"`
	PremiumCalculation  string      `json:"premium_calculation,omitempty"`
	ApplicantAge        int         `json:"applicant_age"`
	Smoker              bool        `json:"smoker,omitempty"`
	TermYears           int         `json:"term_years"`
	PremiumPaymentMode  string      `json:"premium_payment_mode"`
	RiderPremium        *Money      `json:"rider_premium,omitempty"`
	ApplicantOccupation string      `json:"applicant_occupation,omitempty"`
	ValidUntil          time.Time   `json:"valid_until"`
	InsurerProductId    string      `json:"insurer_product_id"`
	SumAssured          *Money      `json:"sum_assured,omitempty"`
	ConvertedPolicyId   string      `json:"converted_policy_id,omitempty"`
}
