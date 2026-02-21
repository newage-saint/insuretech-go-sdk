package models

// PremiumCalculationRequest represents a premium_calculation_request
type PremiumCalculationRequest struct {
	TenureMonths  int                    `json:"tenure_months,omitempty"`
	RiderIds      []string               `json:"rider_ids,omitempty"`
	ApplicantData map[string]interface{} `json:"applicant_data,omitempty"`
	ProductId     string                 `json:"product_id"`
	SumInsured    *Money                 `json:"sum_insured,omitempty"`
}
