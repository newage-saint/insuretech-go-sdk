package models

// ApiKeysListingRequest represents a api_keys_listing_request
type ApiKeysListingRequest struct {
	Page      int    `json:"page,omitempty"`
	PageSize  int    `json:"page_size,omitempty"`
	OwnerType string `json:"owner_type,omitempty"`
	OwnerId   string `json:"owner_id"`
	Status    string `json:"status,omitempty"`
}
