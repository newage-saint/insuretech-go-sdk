package models

// GracePeriodRetrievalResponse represents a grace_period_retrieval_response
type GracePeriodRetrievalResponse struct {
	Error       *Error       `json:"error,omitempty"`
	GracePeriod *GracePeriod `json:"grace_period,omitempty"`
}
