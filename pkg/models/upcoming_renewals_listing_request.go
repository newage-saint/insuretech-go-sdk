package models

// UpcomingRenewalsListingRequest represents a upcoming_renewals_listing_request
type UpcomingRenewalsListingRequest struct {
	DaysAhead int    `json:"days_ahead"`
	Status    string `json:"status,omitempty"`
	Page      int    `json:"page,omitempty"`
	PageSize  int    `json:"page_size,omitempty"`
}
