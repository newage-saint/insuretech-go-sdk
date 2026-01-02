package models

// ApiKeysListingResponse represents a api_keys_listing_response
type ApiKeysListingResponse struct {
	ApiKeys    []*ApiKey `json:"api_keys,omitempty"`
	TotalCount int       `json:"total_count,omitempty"`
	Error      *Error    `json:"error,omitempty"`
}
