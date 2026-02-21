package models

// TenantsListingResponse represents a tenants_listing_response
type TenantsListingResponse struct {
	Error      *Error    `json:"error,omitempty"`
	Tenants    []*Tenant `json:"tenants,omitempty"`
	TotalCount int       `json:"total_count,omitempty"`
}
