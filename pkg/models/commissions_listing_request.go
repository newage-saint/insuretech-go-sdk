package models

// CommissionsListingRequest represents a commissions_listing_request
type CommissionsListingRequest struct {
	Page          int    `json:"page,omitempty"`
	PageSize      int    `json:"page_size,omitempty"`
	RecipientType string `json:"recipient_type,omitempty"`
	RecipientId   string `json:"recipient_id"`
	Status        string `json:"status,omitempty"`
	StartDate     string `json:"start_date,omitempty"`
	EndDate       string `json:"end_date,omitempty"`
}
