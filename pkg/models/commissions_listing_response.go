package models

// CommissionsListingResponse represents a commissions_listing_response
type CommissionsListingResponse struct {
	Commissions []*Commission `json:"commissions,omitempty"`
	TotalCount  int           `json:"total_count,omitempty"`
	TotalAmount *Money        `json:"total_amount,omitempty"`
	Error       *Error        `json:"error,omitempty"`
}
