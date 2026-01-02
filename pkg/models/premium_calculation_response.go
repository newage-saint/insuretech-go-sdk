package models

// PremiumCalculationResponse represents a premium_calculation_response
type PremiumCalculationResponse struct {
	Breakdown    []*PremiumBreakdown `json:"breakdown,omitempty"`
	Error        *Error              `json:"error,omitempty"`
	BasePremium  *Money              `json:"base_premium,omitempty"`
	RiderPremium *Money              `json:"rider_premium,omitempty"`
	TotalPremium *Money              `json:"total_premium,omitempty"`
}
