package models

// RiskScoreUpdateRequest represents a risk_score_update_request
type RiskScoreUpdateRequest struct {
	RiskScore     string `json:"risk_score,omitempty"`
	Reason        string `json:"reason,omitempty"`
	BeneficiaryId string `json:"beneficiary_id"`
}
