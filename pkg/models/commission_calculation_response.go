package models

// CommissionCalculationResponse represents a commission_calculation_response
type CommissionCalculationResponse struct {
	Error                *Error `json:"error,omitempty"`
	CommissionId         string `json:"commission_id,omitempty"`
	CommissionNumber     string `json:"commission_number,omitempty"`
	Amount               *Money `json:"amount,omitempty"`
	CalculationBreakdown string `json:"calculation_breakdown,omitempty"`
}
