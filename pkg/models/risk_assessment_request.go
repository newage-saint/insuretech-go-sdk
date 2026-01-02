package models

// RiskAssessmentRequest represents a risk_assessment_request
type RiskAssessmentRequest struct {
	ApplicantId   string                 `json:"applicant_id"`
	ProductId     string                 `json:"product_id"`
	ApplicantData map[string]interface{} `json:"applicant_data,omitempty"`
}
