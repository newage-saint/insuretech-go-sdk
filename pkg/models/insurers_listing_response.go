package models

// InsurersListingResponse represents a insurers_listing_response
type InsurersListingResponse struct {
	TotalCount int        `json:"total_count,omitempty"`
	Error      *Error     `json:"error,omitempty"`
	Insurers   []*Insurer `json:"insurers,omitempty"`
}
