package models

// TenantsListingRequest represents a tenants_listing_request
type TenantsListingRequest struct {
	Page     int    `json:"page,omitempty"`
	PageSize int    `json:"page_size,omitempty"`
	Type     string `json:"type"`
	Status   string `json:"status,omitempty"`
}
