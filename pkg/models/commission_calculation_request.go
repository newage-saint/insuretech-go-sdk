package models

// CommissionCalculationRequest represents a commission_calculation_request
type CommissionCalculationRequest struct {
	RecipientId    string `json:"recipient_id"`
	PolicyId       string `json:"policy_id"`
	CommissionType string `json:"commission_type,omitempty"`
	RecipientType  string `json:"recipient_type,omitempty"`
}
