package models

// FraudRulesListingRequest represents a fraud_rules_listing_request
type FraudRulesListingRequest struct {
	ActiveOnly bool   `json:"active_only,omitempty"`
	PageSize   int    `json:"page_size,omitempty"`
	PageToken  string `json:"page_token,omitempty"`
	Category   string `json:"category"`
}
