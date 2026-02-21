package models

// UserClaimsListingResponse represents a user_claims_listing_response
type UserClaimsListingResponse struct {
	TotalCount int      `json:"total_count,omitempty"`
	Error      *Error   `json:"error,omitempty"`
	Claims     []*Claim `json:"claims,omitempty"`
}
