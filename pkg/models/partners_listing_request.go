package models

// PartnersListingRequest represents a partners_listing_request
type PartnersListingRequest struct {
	PageSize  int    `json:"page_size"`
	PageToken string `json:"page_token,omitempty"`
	Filter    string `json:"filter,omitempty"`
	OrderBy   string `json:"order_by,omitempty"`
}
