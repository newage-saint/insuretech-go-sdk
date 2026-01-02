package models

// SessionsListingRequest represents a sessions_listing_request
type SessionsListingRequest struct {
	UserId    string `json:"user_id"`
	PageSize  int    `json:"page_size,omitempty"`
	PageToken string `json:"page_token,omitempty"`
}
