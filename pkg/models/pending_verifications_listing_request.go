package models

// PendingVerificationsListingRequest represents a pending_verifications_listing_request
type PendingVerificationsListingRequest struct {
	PageToken string `json:"page_token,omitempty"`
	PageSize  int    `json:"page_size"`
}
