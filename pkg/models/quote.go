package models

import (
	"time"
)

// Quote represents a quote
type Quote struct {
	PremiumCalculation  string      `json:"premium_calculation,omitempty"`
	Smoker              bool        `json:"smoker,omitempty"`
	QuoteNumber         string      `json:"quote_number"`
	TotalPremium        *Money      `json:"total_premium,omitempty"`
	SelectedRiders      string      `json:"selected_riders,omitempty"`
	ApplicantOccupation string      `json:"applicant_occupation,omitempty"`
	ValidUntil          time.Time   `json:"valid_until"`
	BeneficiaryId       string      `json:"beneficiary_id"`
	ApplicantAge        int         `json:"applicant_age"`
	ConvertedPolicyId   string      `json:"converted_policy_id,omitempty"`
	Id                  string      `json:"id"`
	Status              interface{} `json:"status"`
	PremiumPaymentMode  string      `json:"premium_payment_mode"`
	BasePremium         *Money      `json:"base_premium,omitempty"`
	RiderPremium        *Money      `json:"rider_premium,omitempty"`
	ConvertedAt         time.Time   `json:"converted_at,omitempty"`
	AuditInfo           interface{} `json:"audit_info"`
	InsurerProductId    string      `json:"insurer_product_id"`
	SumAssured          *Money      `json:"sum_assured,omitempty"`
	TermYears           int         `json:"term_years"`
	TaxAmount           *Money      `json:"tax_amount,omitempty"`
}
