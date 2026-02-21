package models

// CommissionRetrievalResponse represents a commission_retrieval_response
type CommissionRetrievalResponse struct {
	Error      *Error      `json:"error,omitempty"`
	Commission *Commission `json:"commission,omitempty"`
}
