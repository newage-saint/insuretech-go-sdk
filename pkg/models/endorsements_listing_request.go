package models

// EndorsementsListingRequest represents a endorsements_listing_request
type EndorsementsListingRequest struct {
	Status   string `json:"status,omitempty"`
	Page     int    `json:"page,omitempty"`
	PageSize int    `json:"page_size,omitempty"`
	PolicyId string `json:"policy_id"`
}
