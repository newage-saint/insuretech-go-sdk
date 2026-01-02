package models

// PartnersListingResponse represents a partners_listing_response
type PartnersListingResponse struct {
	TotalCount    int        `json:"total_count,omitempty"`
	Error         *Error     `json:"error,omitempty"`
	Partners      []*Partner `json:"partners,omitempty"`
	NextPageToken string     `json:"next_page_token,omitempty"`
}
