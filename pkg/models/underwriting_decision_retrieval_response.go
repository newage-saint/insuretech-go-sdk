package models

// UnderwritingDecisionRetrievalResponse represents a underwriting_decision_retrieval_response
type UnderwritingDecisionRetrievalResponse struct {
	Error    *Error                `json:"error,omitempty"`
	Decision *UnderwritingDecision `json:"decision,omitempty"`
}
