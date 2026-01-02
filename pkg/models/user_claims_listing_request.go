package models

// UserClaimsListingRequest represents a user_claims_listing_request
type UserClaimsListingRequest struct {
	Status     *ClaimStatus `json:"status,omitempty"`
	Page       int          `json:"page,omitempty"`
	PageSize   int          `json:"page_size,omitempty"`
	CustomerId string       `json:"customer_id"`
}
