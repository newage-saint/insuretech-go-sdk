package models

// ClaimEvaluationResponse represents a claim_evaluation_response
type ClaimEvaluationResponse struct {
	RequiredDocuments []string `json:"required_documents,omitempty"`
	Error             *Error   `json:"error,omitempty"`
	AnalysisId        string   `json:"analysis_id,omitempty"`
	Recommendation    string   `json:"recommendation,omitempty"`
	SuggestedAmount   *Money   `json:"suggested_amount,omitempty"`
	Confidence        float64  `json:"confidence,omitempty"`
	Findings          []string `json:"findings,omitempty"`
}
