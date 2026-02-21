package models

import (
	"time"
)

// Quote represents a quote
type Quote struct {
	RiderPremium        *Money      `json:"rider_premium,omitempty"`
	TotalPremium        *Money      `json:"total_premium,omitempty"`
	ConvertedPolicyId   string      `json:"converted_policy_id,omitempty"`
	QuoteNumber         string      `json:"quote_number"`
	InsurerProductId    string      `json:"insurer_product_id"`
	ApplicantOccupation string      `json:"applicant_occupation,omitempty"`
	ValidUntil          time.Time   `json:"valid_until"`
	AuditInfo           interface{} `json:"audit_info"`
	SumAssured          *Money      `json:"sum_assured,omitempty"`
	TermYears           int         `json:"term_years"`
	BasePremium         *Money      `json:"base_premium,omitempty"`
	TaxAmount           *Money      `json:"tax_amount,omitempty"`
	SelectedRiders      string      `json:"selected_riders,omitempty"`
	ApplicantAge        int         `json:"applicant_age"`
	Smoker              bool        `json:"smoker,omitempty"`
	ConvertedAt         time.Time   `json:"converted_at,omitempty"`
	Id                  string      `json:"id"`
	BeneficiaryId       string      `json:"beneficiary_id"`
	Status              interface{} `json:"status"`
	PremiumPaymentMode  string      `json:"premium_payment_mode"`
	PremiumCalculation  string      `json:"premium_calculation,omitempty"`
}
