package models

// FraudAlertsListingResponse represents a fraud_alerts_listing_response
type FraudAlertsListingResponse struct {
	TotalCount  int           `json:"total_count,omitempty"`
	Error       *Error        `json:"error,omitempty"`
	FraudAlerts []*FraudAlert `json:"fraud_alerts,omitempty"`
}
